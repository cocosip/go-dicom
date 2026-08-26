// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	// ErrManagedClientClosed indicates that a job cannot be queued because the
	// managed client has started closing.
	ErrManagedClientClosed = errors.New("managed client is closed")

	// ErrManagedClientSending indicates that another Send call owns the queue.
	ErrManagedClientSending = errors.New("managed client is already sending")
)

// ManagedClientOptions controls association reuse for ManagedClient.
type ManagedClientOptions struct {
	MaximumRequestsPerAssociation         int
	AssociationLingerTimeout              time.Duration
	MaximumConsecutiveAssociationTimeouts int
}

// ManagedOption configures a ManagedClient.
type ManagedOption func(*managedConfig)

type managedConfig struct {
	ManagedClientOptions
	baseClientOptions []Option
	err               error
}

// PresentationContextSpec describes one requested presentation context.
type PresentationContextSpec struct {
	AbstractSyntax   string
	TransferSyntaxes []string
	SCURole          bool
	SCPRole          bool
}

// Job is an operation queued for ManagedClient delivery.
// Complete is called exactly once after Execute is attempted or the job is
// rejected because the managed client can no longer deliver it.
type Job interface {
	PresentationContexts() ([]PresentationContextSpec, error)
	Execute(context.Context, *Client) error
	Complete(error)
}

type managedJob struct {
	job      Job
	contexts []PresentationContextSpec
}

// ManagedClient batches compatible Jobs on automatically configured low-level
// Client associations. It is safe to Add concurrently with Send.
type ManagedClient struct {
	mu      sync.Mutex
	options managedConfig
	jobs    []managedJob
	sending bool
	closed  bool
	wakeCh  chan struct{}
	active  *Client
}

// NewManaged creates a managed association queue.
func NewManaged(opts ...ManagedOption) *ManagedClient {
	config := managedConfig{ManagedClientOptions: ManagedClientOptions{
		AssociationLingerTimeout:              50 * time.Millisecond,
		MaximumConsecutiveAssociationTimeouts: 3,
	}}
	for _, option := range opts {
		if option != nil {
			option(&config)
		}
	}
	return &ManagedClient{options: config, wakeCh: make(chan struct{}, 1)}
}

// WithBaseClientOptions supplies a copied immutable option slice for each
// low-level Client association.
func WithBaseClientOptions(options ...Option) ManagedOption {
	copied := append([]Option(nil), options...)
	return func(config *managedConfig) { config.baseClientOptions = copied }
}

// WithMaximumRequestsPerAssociation limits jobs sent over one association. A
// zero value keeps the limit unbounded (while presentation contexts remain
// limited to 128).
func WithMaximumRequestsPerAssociation(maximum int) ManagedOption {
	return func(config *managedConfig) {
		if maximum < 0 {
			config.err = fmt.Errorf("maximum requests per association must not be negative")
			return
		}
		config.MaximumRequestsPerAssociation = maximum
	}
}

// WithAssociationLingerTimeout controls how long Send waits for a newly added
// job before releasing an otherwise idle association.
func WithAssociationLingerTimeout(timeout time.Duration) ManagedOption {
	return func(config *managedConfig) {
		if timeout < 0 {
			config.err = fmt.Errorf("association linger timeout must not be negative")
			return
		}
		config.AssociationLingerTimeout = timeout
	}
}

// WithMaximumConsecutiveAssociationTimeouts limits pre-request association
// negotiation retries.
func WithMaximumConsecutiveAssociationTimeouts(maximum int) ManagedOption {
	return func(config *managedConfig) {
		if maximum <= 0 {
			config.err = fmt.Errorf("maximum consecutive association timeouts must be positive")
			return
		}
		config.MaximumConsecutiveAssociationTimeouts = maximum
	}
}

// Add validates and queues a job without invoking its completion callback.
func (c *ManagedClient) Add(job Job) error {
	if job == nil {
		return fmt.Errorf("managed job is nil")
	}
	contexts, err := job.PresentationContexts()
	if err != nil {
		return fmt.Errorf("job presentation contexts: %w", err)
	}
	cloned, err := clonePresentationContextSpecs(contexts)
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.options.err != nil {
		return c.options.err
	}
	if c.closed {
		return ErrManagedClientClosed
	}
	c.jobs = append(c.jobs, managedJob{job: job, contexts: cloned})
	c.signalLocked()
	return nil
}

// Close rejects future jobs and completes queued jobs that have not yet been
// assigned to an association. It is safe to call repeatedly.
func (c *ManagedClient) Close() error {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return nil
	}
	c.closed = true
	jobs := append([]managedJob(nil), c.jobs...)
	c.jobs = nil
	active := c.active
	c.signalLocked()
	c.mu.Unlock()
	for _, queued := range jobs {
		queued.job.Complete(ErrManagedClientClosed)
	}
	if active != nil {
		return active.Abort(context.Background())
	}
	return nil
}

// Send drains queued jobs in FIFO association batches. Only one Send call may
// own a ManagedClient at a time. A job failure completes that job and does not
// prevent later jobs in the same healthy association from executing.
func (c *ManagedClient) Send(ctx context.Context, host string, port int) error {
	c.mu.Lock()
	if c.options.err != nil {
		err := c.options.err
		c.mu.Unlock()
		return err
	}
	if c.closed {
		c.mu.Unlock()
		return ErrManagedClientClosed
	}
	if c.sending {
		c.mu.Unlock()
		return ErrManagedClientSending
	}
	c.sending = true
	c.mu.Unlock()
	defer func() {
		c.mu.Lock()
		c.sending = false
		c.mu.Unlock()
	}()

	for {
		jobs, contexts := c.takeBatch()
		if len(jobs) == 0 {
			if !c.waitForJob(ctx) {
				return nil
			}
			continue
		}
		if err := c.sendBatch(ctx, host, port, jobs, contexts); err != nil {
			return err
		}
	}
}

func (c *ManagedClient) signalLocked() {
	select {
	case c.wakeCh <- struct{}{}:
	default:
	}
}

// waitForJob waits for an Add notification until linger expires. It returns
// true only when a queued job is available for a subsequent takeBatch call.
func (c *ManagedClient) waitForJob(ctx context.Context) bool {
	for {
		c.mu.Lock()
		available := len(c.jobs) != 0
		closed := c.closed
		linger := c.options.AssociationLingerTimeout
		c.mu.Unlock()
		if available {
			return true
		}
		if closed || linger == 0 {
			return false
		}
		timer := time.NewTimer(linger)
		select {
		case <-c.wakeCh:
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			continue
		case <-timer.C:
			return false
		case <-ctx.Done():
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			return false
		}
	}
}

func (c *ManagedClient) sendBatch(
	ctx context.Context,
	host string,
	port int,
	jobs []managedJob,
	contexts []PresentationContextSpec,
) error {
	c.mu.Lock()
	options := append([]Option(nil), c.options.baseClientOptions...)
	maximumTimeouts := c.options.MaximumConsecutiveAssociationTimeouts
	c.mu.Unlock()

	var lowLevel *Client
	consecutiveTimeouts := 0
	for {
		candidate, err := newManagedLowLevelClient(options, contexts)
		if err != nil {
			completeManagedJobs(jobs, err)
			return err
		}
		err = candidate.Connect(ctx, host, port)
		if err == nil {
			lowLevel = candidate
			break
		}
		_ = candidate.Close()
		if !isAssociationTimeout(err) || consecutiveTimeouts+1 >= maximumTimeouts || c.isClosed() {
			completeManagedJobs(jobs, err)
			return err
		}
		consecutiveTimeouts++
	}
	if !c.setActive(lowLevel) {
		completeManagedJobs(jobs, ErrManagedClientClosed)
		_ = lowLevel.Close()
		return ErrManagedClientClosed
	}
	defer func() {
		c.clearActive(lowLevel)
		_ = lowLevel.Close()
	}()

	executed := 0
	for {
		for index, queued := range jobs {
			if c.isClosed() {
				completeManagedJobs(jobs[index:], ErrManagedClientClosed)
				return ErrManagedClientClosed
			}
			err := queued.job.Execute(ctx, lowLevel)
			queued.job.Complete(err)
			executed++
		}
		jobs = nil

		next, ok := c.waitForCompatibleJob(ctx, contexts, executed)
		if !ok {
			if c.isClosed() {
				return ErrManagedClientClosed
			}
			return nil
		}
		jobs = []managedJob{next}
	}
}

func newManagedLowLevelClient(options []Option, contexts []PresentationContextSpec) (*Client, error) {
	lowLevel := New(options...)
	for _, presentationContext := range contexts {
		var err error
		if presentationContext.SCURole || presentationContext.SCPRole {
			err = lowLevel.AddPresentationContextWithRoles(
				presentationContext.AbstractSyntax,
				presentationContext.SCURole,
				presentationContext.SCPRole,
				presentationContext.TransferSyntaxes...,
			)
		} else {
			err = lowLevel.AddPresentationContext(presentationContext.AbstractSyntax, presentationContext.TransferSyntaxes...)
		}
		if err != nil {
			return nil, err
		}
	}
	return lowLevel, nil
}

func (c *ManagedClient) waitForCompatibleJob(
	ctx context.Context,
	contexts []PresentationContextSpec,
	executed int,
) (managedJob, bool) {
	for {
		if next, ok := c.takeCompatibleJob(contexts, executed); ok {
			return next, true
		}
		c.mu.Lock()
		queued := len(c.jobs) != 0
		closed := c.closed
		linger := c.options.AssociationLingerTimeout
		c.mu.Unlock()
		if queued || closed || linger == 0 {
			return managedJob{}, false
		}

		timer := time.NewTimer(linger)
		select {
		case <-c.wakeCh:
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
		case <-timer.C:
			return managedJob{}, false
		case <-ctx.Done():
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			return managedJob{}, false
		}
	}
}

func (c *ManagedClient) takeCompatibleJob(
	contexts []PresentationContextSpec,
	executed int,
) (managedJob, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed || len(c.jobs) == 0 ||
		(c.options.MaximumRequestsPerAssociation > 0 && executed >= c.options.MaximumRequestsPerAssociation) {
		return managedJob{}, false
	}
	indexes := make(map[presentationContextKey]int, len(contexts))
	for index, presentationContext := range contexts {
		indexes[presentationContextKey{
			abstractSyntax: presentationContext.AbstractSyntax,
			scuRole:        presentationContext.SCURole,
			scpRole:        presentationContext.SCPRole,
		}] = index
	}
	merged, ok := mergePresentationContextSpecs(contexts, indexes, c.jobs[0].contexts)
	if !ok || !equalPresentationContextSpecs(merged, contexts) {
		return managedJob{}, false
	}
	next := c.jobs[0]
	c.jobs = append([]managedJob(nil), c.jobs[1:]...)
	return next, true
}

func equalPresentationContextSpecs(first, second []PresentationContextSpec) bool {
	if len(first) != len(second) {
		return false
	}
	for index := range first {
		if first[index].AbstractSyntax != second[index].AbstractSyntax ||
			first[index].SCURole != second[index].SCURole ||
			first[index].SCPRole != second[index].SCPRole ||
			len(first[index].TransferSyntaxes) != len(second[index].TransferSyntaxes) {
			return false
		}
		for syntaxIndex := range first[index].TransferSyntaxes {
			if first[index].TransferSyntaxes[syntaxIndex] != second[index].TransferSyntaxes[syntaxIndex] {
				return false
			}
		}
	}
	return true
}

func (c *ManagedClient) setActive(lowLevel *Client) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return false
	}
	c.active = lowLevel
	return true
}

func (c *ManagedClient) clearActive(lowLevel *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.active == lowLevel {
		c.active = nil
	}
}

func (c *ManagedClient) isClosed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.closed
}

func isAssociationTimeout(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	var networkErr net.Error
	return errors.As(err, &networkErr) && networkErr.Timeout()
}

func completeManagedJobs(jobs []managedJob, err error) {
	for _, queued := range jobs {
		queued.job.Complete(err)
	}
}

func clonePresentationContextSpecs(contexts []PresentationContextSpec) ([]PresentationContextSpec, error) {
	cloned := make([]PresentationContextSpec, len(contexts))
	unique := make(map[presentationContextKey]struct{}, len(contexts))
	for index, context := range contexts {
		if context.AbstractSyntax == "" {
			return nil, fmt.Errorf("%w: abstract syntax is empty", ErrInvalidPresentationContext)
		}
		if len(context.TransferSyntaxes) == 0 {
			return nil, fmt.Errorf("%w: transfer syntaxes are empty", ErrInvalidPresentationContext)
		}
		for _, syntax := range context.TransferSyntaxes {
			if syntax == "" {
				return nil, fmt.Errorf("%w: transfer syntax is empty", ErrInvalidPresentationContext)
			}
		}
		unique[presentationContextKey{context.AbstractSyntax, context.SCURole, context.SCPRole}] = struct{}{}
		if len(unique) > 128 {
			return nil, fmt.Errorf("%w: job proposes %d unique contexts", ErrTooManyPresentationContexts, len(unique))
		}
		cloned[index] = PresentationContextSpec{
			AbstractSyntax:   context.AbstractSyntax,
			TransferSyntaxes: append([]string(nil), context.TransferSyntaxes...),
			SCURole:          context.SCURole,
			SCPRole:          context.SCPRole,
		}
	}
	return cloned, nil
}

type presentationContextKey struct {
	abstractSyntax string
	scuRole        bool
	scpRole        bool
}

// takeBatch removes the next FIFO-compatible job batch. It never executes a
// Job or calls Complete while holding the managed client mutex.
func (c *ManagedClient) takeBatch() ([]managedJob, []PresentationContextSpec) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.jobs) == 0 {
		return nil, nil
	}

	contexts := make([]PresentationContextSpec, 0, 128)
	indexes := make(map[presentationContextKey]int)
	count := 0
	for count < len(c.jobs) {
		if c.options.MaximumRequestsPerAssociation > 0 && count >= c.options.MaximumRequestsPerAssociation {
			break
		}
		merged, ok := mergePresentationContextSpecs(contexts, indexes, c.jobs[count].contexts)
		if !ok {
			break
		}
		contexts = merged
		for index, context := range contexts {
			indexes[presentationContextKey{context.AbstractSyntax, context.SCURole, context.SCPRole}] = index
		}
		count++
	}
	if count == 0 {
		return nil, nil
	}
	jobs := append([]managedJob(nil), c.jobs[:count]...)
	c.jobs = append([]managedJob(nil), c.jobs[count:]...)
	return jobs, contexts
}

func mergePresentationContextSpecs(
	current []PresentationContextSpec,
	indexes map[presentationContextKey]int,
	additional []PresentationContextSpec,
) ([]PresentationContextSpec, bool) {
	merged := make([]PresentationContextSpec, len(current))
	for index, context := range current {
		merged[index] = PresentationContextSpec{
			AbstractSyntax:   context.AbstractSyntax,
			TransferSyntaxes: append([]string(nil), context.TransferSyntaxes...),
			SCURole:          context.SCURole,
			SCPRole:          context.SCPRole,
		}
	}
	mergedIndexes := make(map[presentationContextKey]int, len(indexes))
	for key, index := range indexes {
		mergedIndexes[key] = index
	}
	for _, context := range additional {
		key := presentationContextKey{context.AbstractSyntax, context.SCURole, context.SCPRole}
		if index, exists := mergedIndexes[key]; exists {
			merged[index].TransferSyntaxes = mergeTransferSyntaxes(merged[index].TransferSyntaxes, context.TransferSyntaxes)
			continue
		}
		if len(merged) == 128 {
			return nil, false
		}
		mergedIndexes[key] = len(merged)
		merged = append(merged, PresentationContextSpec{
			AbstractSyntax:   context.AbstractSyntax,
			TransferSyntaxes: append([]string(nil), context.TransferSyntaxes...),
			SCURole:          context.SCURole,
			SCPRole:          context.SCPRole,
		})
	}
	return merged, true
}

func mergeTransferSyntaxes(first, additional []string) []string {
	merged := append([]string(nil), first...)
	seen := make(map[string]struct{}, len(merged))
	for _, syntax := range merged {
		seen[syntax] = struct{}{}
	}
	for _, syntax := range additional {
		if _, exists := seen[syntax]; exists {
			continue
		}
		seen[syntax] = struct{}{}
		merged = append(merged, syntax)
	}
	return merged
}
