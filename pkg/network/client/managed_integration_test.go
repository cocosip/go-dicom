// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"errors"
	"net"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/service"
)

type managedExecutionJob struct {
	contexts []PresentationContextSpec
	execute  func(context.Context, *Client) error
	complete func(error)
}

func (j *managedExecutionJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return j.contexts, nil
}

func (j *managedExecutionJob) Execute(ctx context.Context, client *Client) error {
	if j.execute == nil {
		return nil
	}
	return j.execute(ctx, client)
}

func (j *managedExecutionJob) Complete(err error) {
	if j.complete != nil {
		j.complete(err)
	}
}

func TestManagedSendKeepsCompatibleAssociationForLingerAdd(t *testing.T) {
	host, port, associations, closeServer := startManagedAssociationServer(t)
	defer closeServer()

	managed := NewManaged(WithAssociationLingerTimeout(500 * time.Millisecond))
	firstExecuted := make(chan struct{})
	secondExecuted := make(chan struct{})
	contexts := []PresentationContextSpec{{
		AbstractSyntax:   verificationSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
	}}
	if err := managed.Add(&managedExecutionJob{
		contexts: contexts,
		execute: func(context.Context, *Client) error {
			close(firstExecuted)
			return nil
		},
	}); err != nil {
		t.Fatalf("Add(first) error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sent := make(chan error, 1)
	go func() { sent <- managed.Send(ctx, host, port) }()
	select {
	case <-firstExecuted:
	case <-ctx.Done():
		t.Fatal("first managed job was not executed")
	}
	if err := managed.Add(&managedExecutionJob{
		contexts: contexts,
		execute: func(context.Context, *Client) error {
			close(secondExecuted)
			return nil
		},
	}); err != nil {
		t.Fatalf("Add(second) error = %v", err)
	}
	select {
	case err := <-sent:
		if err != nil {
			t.Fatalf("Send() error = %v", err)
		}
	case <-ctx.Done():
		t.Fatal("Send() did not finish")
	}
	select {
	case <-secondExecuted:
	case <-ctx.Done():
		t.Fatal("second managed job was not executed")
	}

	select {
	case <-associations:
	case <-ctx.Done():
		t.Fatal("server did not observe an association")
	}
	select {
	case <-associations:
		t.Fatal("compatible linger job opened a second association")
	case <-time.After(100 * time.Millisecond):
	}
}

func TestManagedCloseDuringLingerJobCompletesDequeuedJobOnce(t *testing.T) {
	host, port, _, closeServer := startManagedAssociationServer(t)
	defer closeServer()

	managed := NewManaged(WithAssociationLingerTimeout(time.Second))
	contexts := []PresentationContextSpec{{
		AbstractSyntax:   verificationSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
	}}
	firstExecuted := make(chan struct{})
	secondCompleted := make(chan error, 1)
	if err := managed.Add(&managedExecutionJob{
		contexts: contexts,
		execute: func(context.Context, *Client) error {
			close(firstExecuted)
			return nil
		},
	}); err != nil {
		t.Fatalf("Add(first) error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sent := make(chan error, 1)
	go func() { sent <- managed.Send(ctx, host, port) }()
	select {
	case <-firstExecuted:
	case <-ctx.Done():
		t.Fatal("first managed job was not executed")
	}

	if err := managed.Add(&managedExecutionJob{
		contexts: contexts,
		complete: func(err error) { secondCompleted <- err },
	}); err != nil {
		t.Fatalf("Add(second) error = %v", err)
	}

	// Make Close contend with the sender while it consumes the linger job. The
	// observable contract is that the dequeued job still completes exactly once.
	closed := make(chan error, 1)
	go func() {
		runtime.Gosched()
		closed <- managed.Close()
	}()
	select {
	case err := <-closed:
		if err != nil {
			t.Fatalf("Close() error = %v", err)
		}
	case <-ctx.Done():
		t.Fatal("Close() did not return")
	}
	select {
	case err := <-sent:
		if !errors.Is(err, ErrManagedClientClosed) {
			t.Fatalf("Send() error = %v, want ErrManagedClientClosed", err)
		}
	case <-ctx.Done():
		t.Fatal("Send() did not return after Close")
	}
	select {
	case <-secondCompleted:
	case <-ctx.Done():
		t.Fatal("dequeued linger job was not completed")
	}
	select {
	case extra := <-secondCompleted:
		t.Fatalf("dequeued linger job completed more than once: %v", extra)
	case <-time.After(50 * time.Millisecond):
	}
}

func TestManagedCloseAbortsActiveAssociationAndCompletesJobsOnce(t *testing.T) {
	host, port, requestStarted, closeServer := startBlockingCEchoServer(t)
	defer closeServer()

	managed := NewManaged(WithAssociationLingerTimeout(0))
	firstCompleted := make(chan error, 1)
	secondCompleted := make(chan error, 1)
	if err := managed.Add(NewCEchoJob(func(err error) { firstCompleted <- err })); err != nil {
		t.Fatalf("Add(first) error = %v", err)
	}
	if err := managed.Add(NewCEchoJob(func(err error) { secondCompleted <- err })); err != nil {
		t.Fatalf("Add(second) error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sent := make(chan error, 1)
	go func() { sent <- managed.Send(ctx, host, port) }()
	select {
	case <-requestStarted:
	case <-ctx.Done():
		t.Fatal("server did not receive active C-ECHO")
	}
	if err := managed.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	select {
	case err := <-sent:
		if !errors.Is(err, ErrManagedClientClosed) {
			t.Fatalf("Send() error = %v, want ErrManagedClientClosed", err)
		}
	case <-ctx.Done():
		t.Fatal("Send() did not return after Close")
	}
	for name, complete := range map[string]<-chan error{
		"first":  firstCompleted,
		"second": secondCompleted,
	} {
		select {
		case <-complete:
		case <-ctx.Done():
			t.Fatalf("%s Job was not completed", name)
		}
		select {
		case extra := <-complete:
			t.Fatalf("%s Job completed more than once: %v", name, extra)
		case <-time.After(50 * time.Millisecond):
		}
	}
}

func TestManagedSendSplitsMoreThan128ContextsAcrossTCPAssociations(t *testing.T) {
	host, port, associations, closeServer := startManagedAssociationServer(t)
	defer closeServer()

	managed := NewManaged(WithAssociationLingerTimeout(0))
	var executedMu sync.Mutex
	executed := make([]int, 0, 129)
	for index := 0; index < 129; index++ {
		jobIndex := index
		if err := managed.Add(&managedExecutionJob{
			contexts: []PresentationContextSpec{{
				AbstractSyntax:   "1.2.826.0.1.3680043.10.854.100." + strconv.Itoa(jobIndex),
				TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
			}},
			execute: func(context.Context, *Client) error {
				executedMu.Lock()
				executed = append(executed, jobIndex)
				executedMu.Unlock()
				return nil
			},
		}); err != nil {
			t.Fatalf("Add(%d) error = %v", index, err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := managed.Send(ctx, host, port); err != nil {
		t.Fatalf("Send() error = %v", err)
	}
	for position, want := range []int{128, 1} {
		select {
		case got := <-associations:
			if got != want {
				t.Fatalf("association %d contexts = %d, want %d", position+1, got, want)
			}
		case <-ctx.Done():
			t.Fatalf("server did not observe association %d", position+1)
		}
	}
	executedMu.Lock()
	defer executedMu.Unlock()
	if len(executed) != 129 {
		t.Fatalf("executed jobs = %d, want 129", len(executed))
	}
	for index, got := range executed {
		if got != index {
			t.Fatalf("executed job %d = %d, want FIFO %d", index, got, index)
		}
	}
}

func TestManagedSendRetriesAssociationTimeoutBeforeExecutingJob(t *testing.T) {
	host, port, attempts, closeServer := startAssociationTimeoutThenAcceptServer(t)
	defer closeServer()

	managed := NewManaged(
		WithAssociationLingerTimeout(0),
		WithMaximumConsecutiveAssociationTimeouts(2),
		WithBaseClientOptions(WithAssociationTimeout(75*time.Millisecond)),
	)
	executed := make(chan struct{}, 1)
	if err := managed.Add(&managedExecutionJob{
		contexts: []PresentationContextSpec{{
			AbstractSyntax:   verificationSOPClassUID,
			TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
		}},
		execute: func(context.Context, *Client) error {
			executed <- struct{}{}
			return nil
		},
	}); err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := managed.Send(ctx, host, port); err != nil {
		t.Fatalf("Send() error = %v", err)
	}
	select {
	case <-executed:
	case <-ctx.Done():
		t.Fatal("job was not executed after retry")
	}
	select {
	case got := <-attempts:
		if got != 2 {
			t.Fatalf("association attempts = %d, want 2", got)
		}
	case <-ctx.Done():
		t.Fatal("server did not report association attempts")
	}
}

func startManagedAssociationServer(t *testing.T) (string, int, <-chan int, func()) {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	associationAccepted := make(chan int, 8)
	var servicesMu sync.Mutex
	var services []*service.Service
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		for {
			conn, acceptErr := listener.Accept()
			if acceptErr != nil {
				return
			}
			go func() {
				svc := service.NewService(conn, nil,
					service.WithAssociationNegotiator(service.FuncAssociationNegotiator(
						func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
							for _, presentationContext := range assoc.PresentationContexts {
								presentationContext.Accept(presentationContext.ProposedTransferSyntaxes[0])
							}
							return responder.SendAccept(ctx, assoc)
						},
					)),
				)
				servicesMu.Lock()
				services = append(services, svc)
				servicesMu.Unlock()
				if _, serviceErr := svc.ReceiveAssociationRequest(context.Background()); serviceErr != nil {
					return
				}
				if serviceErr := svc.Start(); serviceErr != nil {
					return
				}
				associationAccepted <- len(svc.GetAssociation().PresentationContexts)
				svc.WaitForClose()
			}()
		}
	}()
	address := listener.Addr().(*net.TCPAddr)
	return address.IP.String(), address.Port, associationAccepted, func() {
		_ = listener.Close()
		servicesMu.Lock()
		defer servicesMu.Unlock()
		for _, svc := range services {
			_ = svc.Close()
		}
		<-stopped
	}
}

func startBlockingCEchoServer(t *testing.T) (string, int, <-chan struct{}, func()) {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	requestStarted := make(chan struct{}, 1)
	serviceDone := make(chan struct{})
	go func() {
		defer close(serviceDone)
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			return
		}
		svc := service.NewService(conn, nil,
			service.WithAssociationNegotiator(service.FuncAssociationNegotiator(
				func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
					for _, presentationContext := range assoc.PresentationContexts {
						presentationContext.Accept(presentationContext.ProposedTransferSyntaxes[0])
					}
					return responder.SendAccept(ctx, assoc)
				},
			)),
			service.WithCEchoHandler(func(ctx context.Context, _ *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
				requestStarted <- struct{}{}
				<-ctx.Done()
				return nil, ctx.Err()
			}),
		)
		defer func() { _ = svc.Close() }()
		if _, serviceErr := svc.ReceiveAssociationRequest(context.Background()); serviceErr != nil {
			return
		}
		if serviceErr := svc.Start(); serviceErr != nil {
			return
		}
		svc.WaitForClose()
	}()
	address := listener.Addr().(*net.TCPAddr)
	return address.IP.String(), address.Port, requestStarted, func() {
		_ = listener.Close()
		<-serviceDone
	}
}

func startAssociationTimeoutThenAcceptServer(t *testing.T) (string, int, <-chan int, func()) {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	attempts := make(chan int, 1)
	serverDone := make(chan struct{})
	go func() {
		defer close(serverDone)
		first, acceptErr := listener.Accept()
		if acceptErr != nil {
			return
		}
		defer func() { _ = first.Close() }()
		second, acceptErr := listener.Accept()
		if acceptErr != nil {
			return
		}
		svc := service.NewService(second, nil,
			service.WithAssociationNegotiator(service.FuncAssociationNegotiator(
				func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
					for _, presentationContext := range assoc.PresentationContexts {
						presentationContext.Accept(presentationContext.ProposedTransferSyntaxes[0])
					}
					return responder.SendAccept(ctx, assoc)
				},
			)),
		)
		defer func() { _ = svc.Close() }()
		if _, serviceErr := svc.ReceiveAssociationRequest(context.Background()); serviceErr != nil {
			return
		}
		if serviceErr := svc.Start(); serviceErr != nil {
			return
		}
		attempts <- 2
		svc.WaitForClose()
	}()
	address := listener.Addr().(*net.TCPAddr)
	return address.IP.String(), address.Port, attempts, func() {
		_ = listener.Close()
		<-serverDone
	}
}
