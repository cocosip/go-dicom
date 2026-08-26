// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

const (
	managedTestOriginalName = "ORIGINAL"
	managedTestChangedName  = "CHANGED"
)

type managedTestJob struct {
	contexts  []PresentationContextSpec
	completed []error
}

func (j *managedTestJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return j.contexts, nil
}

func (j *managedTestJob) Execute(context.Context, *Client) error { return nil }

func (j *managedTestJob) Complete(err error) { j.completed = append(j.completed, err) }

func TestManagedAddClonesPresentationContextSpecs(t *testing.T) {
	managed := NewManaged()
	if got, want := managed.options.AssociationLingerTimeout, 50*time.Millisecond; got != want {
		t.Fatalf("AssociationLingerTimeout = %s, want %s", got, want)
	}
	if got, want := managed.options.MaximumConsecutiveAssociationTimeouts, 3; got != want {
		t.Fatalf("MaximumConsecutiveAssociationTimeouts = %d, want %d", got, want)
	}

	job := &managedTestJob{contexts: []PresentationContextSpec{{
		AbstractSyntax:   verificationSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
		SCURole:          true,
	}}}
	if err := managed.Add(job); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	job.contexts[0].TransferSyntaxes[0] = "changed-by-caller"

	managed.mu.Lock()
	defer managed.mu.Unlock()
	if got := managed.jobs[0].contexts[0].TransferSyntaxes[0]; got != testExplicitVRLittleEndianUID {
		t.Fatalf("queued transfer syntax = %q, want deep copy", got)
	}
	if len(job.completed) != 0 {
		t.Fatalf("Complete calls = %d after Add, want 0", len(job.completed))
	}
}

func TestManagedAddRejectsInvalidPresentationContextAtomically(t *testing.T) {
	managed := NewManaged()
	job := &managedTestJob{contexts: []PresentationContextSpec{{
		AbstractSyntax:   verificationSOPClassUID,
		TransferSyntaxes: []string{""},
	}}}
	if err := managed.Add(job); err == nil {
		t.Fatal("Add() error = nil for empty transfer syntax")
	}
	managed.mu.Lock()
	defer managed.mu.Unlock()
	if len(managed.jobs) != 0 {
		t.Fatalf("queued jobs = %d after rejected Add, want 0", len(managed.jobs))
	}
}

func TestManagedAddRejectsJobWithMoreThan128UniqueContexts(t *testing.T) {
	managed := NewManaged()
	contexts := make([]PresentationContextSpec, 129)
	for index := range contexts {
		contexts[index] = PresentationContextSpec{
			AbstractSyntax:   fmt.Sprintf("1.2.826.0.1.3680043.10.854.%d", index),
			TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
		}
	}
	if err := managed.Add(&managedTestJob{contexts: contexts}); !errors.Is(err, ErrTooManyPresentationContexts) {
		t.Fatalf("Add() error = %v, want ErrTooManyPresentationContexts", err)
	}
}

func TestManagedTakeBatchMergesContextsAndSplitsAt128(t *testing.T) {
	managed := NewManaged()
	if err := managed.Add(&managedTestJob{contexts: []PresentationContextSpec{{
		AbstractSyntax:   verificationSOPClassUID,
		TransferSyntaxes: []string{testImplicitVRLittleEndianUID, testExplicitVRLittleEndianUID},
		SCURole:          true,
	}}}); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if err := managed.Add(&managedTestJob{contexts: []PresentationContextSpec{{
		AbstractSyntax:   verificationSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID, "1.2.840.10008.1.2.2"},
		SCURole:          true,
	}}}); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	for index := 0; index < 128; index++ {
		if err := managed.Add(&managedTestJob{contexts: []PresentationContextSpec{{
			AbstractSyntax:   fmt.Sprintf("1.2.826.0.1.3680043.10.854.%d", index),
			TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
		}}}); err != nil {
			t.Fatalf("Add(%d) error = %v", index, err)
		}
	}

	jobs, contexts := managed.takeBatch()
	if got, want := len(jobs), 129; got != want {
		t.Fatalf("first batch jobs = %d, want %d", got, want)
	}
	if got, want := len(contexts), 128; got != want {
		t.Fatalf("first batch contexts = %d, want %d", got, want)
	}
	if got, want := contexts[0].TransferSyntaxes, []string{testImplicitVRLittleEndianUID, testExplicitVRLittleEndianUID, "1.2.840.10008.1.2.2"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("merged transfer syntaxes = %#v, want %#v", got, want)
	}
	jobs, contexts = managed.takeBatch()
	if got, want := len(jobs), 1; got != want {
		t.Fatalf("second batch jobs = %d, want %d", got, want)
	}
	if got, want := len(contexts), 1; got != want {
		t.Fatalf("second batch contexts = %d, want %d", got, want)
	}
}

func TestManagedCloseCompletesQueuedJobsExactlyOnce(t *testing.T) {
	managed := NewManaged()
	job := &managedTestJob{contexts: []PresentationContextSpec{{
		AbstractSyntax:   verificationSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
	}}}
	if err := managed.Add(job); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if err := managed.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if len(job.completed) != 1 || !errors.Is(job.completed[0], ErrManagedClientClosed) {
		t.Fatalf("Complete calls = %#v, want ErrManagedClientClosed once", job.completed)
	}
	if err := managed.Close(); err != nil {
		t.Fatalf("second Close() error = %v", err)
	}
	if len(job.completed) != 1 {
		t.Fatalf("Complete calls = %d after second Close, want 1", len(job.completed))
	}
}

func TestManagedSendEmptyQueueReturnsWithoutConnection(t *testing.T) {
	managed := NewManaged(WithAssociationLingerTimeout(0))
	if err := managed.Send(context.Background(), "127.0.0.1", 1); err != nil {
		t.Fatalf("Send() error = %v", err)
	}
}

func TestManagedWaitForJobIsWokenByConcurrentAdd(t *testing.T) {
	managed := NewManaged(WithAssociationLingerTimeout(time.Second))
	ready := make(chan bool, 1)
	go func() { ready <- managed.waitForJob(context.Background()) }()
	time.Sleep(10 * time.Millisecond)
	if err := managed.Add(&managedTestJob{contexts: []PresentationContextSpec{{
		AbstractSyntax:   verificationSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
	}}}); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	select {
	case available := <-ready:
		if !available {
			t.Fatal("waitForJob() returned false after Add")
		}
	case <-time.After(time.Second):
		t.Fatal("waitForJob() was not woken by Add")
	}
}

func TestNewCEchoJobDerivesVerificationContext(t *testing.T) {
	job := NewCEchoJob(nil)
	contexts, err := job.PresentationContexts()
	if err != nil {
		t.Fatalf("PresentationContexts() error = %v", err)
	}
	if got, want := contexts, []PresentationContextSpec{{
		AbstractSyntax:   verificationSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID, testImplicitVRLittleEndianUID},
	}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("contexts = %#v, want %#v", got, want)
	}
}

func TestNewCStoreJobClonesDatasetAndPrefersSourceTransferSyntax(t *testing.T) {
	dataset := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	if err := dataset.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})); err != nil {
		t.Fatalf("add SOP Class UID: %v", err)
	}
	job, err := NewCStoreJob(dataset, nil)
	if err != nil {
		t.Fatalf("NewCStoreJob() error = %v", err)
	}
	dataset.SetInternalTransferSyntax(transfer.ImplicitVRLittleEndian)

	contexts, err := job.PresentationContexts()
	if err != nil {
		t.Fatalf("PresentationContexts() error = %v", err)
	}
	if got, want := contexts[0].TransferSyntaxes[0], transfer.ExplicitVRBigEndian.UID().UID(); got != want {
		t.Fatalf("first transfer syntax = %q, want source %q", got, want)
	}
}

func TestNewCFindJobDerivesQueryContextAndClonesQuery(t *testing.T) {
	query := dataset.New()
	if err := query.Add(element.NewString(tag.PatientName, vr.PN, []string{managedTestOriginalName})); err != nil {
		t.Fatalf("add PatientName: %v", err)
	}
	job, err := NewCFindJob(dimse.QueryRetrieveLevelStudy, query, nil)
	if err != nil {
		t.Fatalf("NewCFindJob() error = %v", err)
	}
	if err := query.AddOrUpdate(element.NewString(tag.PatientName, vr.PN, []string{managedTestChangedName})); err != nil {
		t.Fatalf("update PatientName: %v", err)
	}

	contexts, err := job.PresentationContexts()
	if err != nil {
		t.Fatalf("PresentationContexts() error = %v", err)
	}
	if got, want := contexts, []PresentationContextSpec{{
		AbstractSyntax:   "1.2.840.10008.5.1.4.1.2.2.1",
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID, testImplicitVRLittleEndianUID},
	}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("contexts = %#v, want %#v", got, want)
	}
	findJob, ok := job.(*managedCFindJob)
	if !ok {
		t.Fatalf("job type = %T, want *managedCFindJob", job)
	}
	if got, _ := findJob.query.GetString(tag.PatientName); got != managedTestOriginalName {
		t.Fatalf("queued PatientName = %q, want independent source copy", got)
	}
}

func TestNewCMoveAndCGetJobsDeriveQueryRetrieveContexts(t *testing.T) {
	identifier := dataset.New()
	moveJob, err := NewCMoveJob(dimse.QueryRetrieveLevelPatient, "MOVE_DEST", identifier, nil, nil)
	if err != nil {
		t.Fatalf("NewCMoveJob() error = %v", err)
	}
	getJob, err := NewCGetJob(dimse.QueryRetrieveLevelStudy, identifier, nil, nil)
	if err != nil {
		t.Fatalf("NewCGetJob() error = %v", err)
	}

	moveContexts, err := moveJob.PresentationContexts()
	if err != nil {
		t.Fatalf("move PresentationContexts() error = %v", err)
	}
	if got, want := moveContexts[0].AbstractSyntax, "1.2.840.10008.5.1.4.1.2.1.2"; got != want {
		t.Fatalf("C-MOVE abstract syntax = %q, want %q", got, want)
	}
	getContexts, err := getJob.PresentationContexts()
	if err != nil {
		t.Fatalf("get PresentationContexts() error = %v", err)
	}
	if got, want := getContexts[0].AbstractSyntax, "1.2.840.10008.5.1.4.1.2.2.3"; got != want {
		t.Fatalf("C-GET abstract syntax = %q, want %q", got, want)
	}
}

func TestNewNActionJobDerivesManagedSOPContextAndClonesActionInformation(t *testing.T) {
	actionInformation := dataset.New()
	if err := actionInformation.Add(element.NewString(tag.PatientName, vr.PN, []string{managedTestOriginalName})); err != nil {
		t.Fatalf("add PatientName: %v", err)
	}
	job, err := NewNActionJob("1.2.826.0.1.3680043.10.854.42", "1.2.3", 7, actionInformation, nil)
	if err != nil {
		t.Fatalf("NewNActionJob() error = %v", err)
	}
	if err := actionInformation.AddOrUpdate(element.NewString(tag.PatientName, vr.PN, []string{managedTestChangedName})); err != nil {
		t.Fatalf("update PatientName: %v", err)
	}
	contexts, err := job.PresentationContexts()
	if err != nil {
		t.Fatalf("PresentationContexts() error = %v", err)
	}
	if got, want := contexts, []PresentationContextSpec{{
		AbstractSyntax:   "1.2.826.0.1.3680043.10.854.42",
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID, testImplicitVRLittleEndianUID},
	}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("contexts = %#v, want %#v", got, want)
	}
	actionJob, ok := job.(*managedNActionJob)
	if !ok {
		t.Fatalf("job type = %T, want *managedNActionJob", job)
	}
	if got, _ := actionJob.actionInformation.GetString(tag.PatientName); got != managedTestOriginalName {
		t.Fatalf("queued PatientName = %q, want independent source copy", got)
	}
}

func TestNewNServiceJobsDeriveManagedSOPContextAndCloneInputs(t *testing.T) {
	const sopClassUID = "1.2.826.0.1.3680043.10.854.43"
	attributes := dataset.New()
	if err := attributes.Add(element.NewString(tag.PatientName, vr.PN, []string{managedTestOriginalName})); err != nil {
		t.Fatalf("add PatientName: %v", err)
	}

	createJob, err := NewNCreateJob(sopClassUID, "1.2.3", attributes, nil)
	if err != nil {
		t.Fatalf("NewNCreateJob() error = %v", err)
	}
	getJob, err := NewNGetJob(sopClassUID, "1.2.3", []*tag.Tag{tag.PatientName}, nil)
	if err != nil {
		t.Fatalf("NewNGetJob() error = %v", err)
	}
	setJob, err := NewNSetJob(sopClassUID, "1.2.3", attributes, nil)
	if err != nil {
		t.Fatalf("NewNSetJob() error = %v", err)
	}
	deleteJob, err := NewNDeleteJob(sopClassUID, "1.2.3", nil)
	if err != nil {
		t.Fatalf("NewNDeleteJob() error = %v", err)
	}
	eventJob, err := NewNEventReportJob(sopClassUID, "1.2.3", 9, attributes, nil)
	if err != nil {
		t.Fatalf("NewNEventReportJob() error = %v", err)
	}
	if err := attributes.AddOrUpdate(element.NewString(tag.PatientName, vr.PN, []string{managedTestChangedName})); err != nil {
		t.Fatalf("update PatientName: %v", err)
	}

	for name, job := range map[string]Job{
		"N-CREATE":       createJob,
		"N-GET":          getJob,
		"N-SET":          setJob,
		"N-DELETE":       deleteJob,
		"N-EVENT-REPORT": eventJob,
	} {
		contexts, contextErr := job.PresentationContexts()
		if contextErr != nil {
			t.Fatalf("%s PresentationContexts() error = %v", name, contextErr)
		}
		if got, want := contexts, []PresentationContextSpec{{
			AbstractSyntax:   sopClassUID,
			TransferSyntaxes: []string{testExplicitVRLittleEndianUID, testImplicitVRLittleEndianUID},
		}}; !reflect.DeepEqual(got, want) {
			t.Fatalf("%s contexts = %#v, want %#v", name, got, want)
		}
	}

	if got, _ := createJob.(*managedNCreateJob).attributeList.GetString(tag.PatientName); got != managedTestOriginalName {
		t.Fatalf("queued N-CREATE PatientName = %q, want independent source copy", got)
	}
	if got, _ := setJob.(*managedNSetJob).modificationList.GetString(tag.PatientName); got != managedTestOriginalName {
		t.Fatalf("queued N-SET PatientName = %q, want independent source copy", got)
	}
	if got, _ := eventJob.(*managedNEventReportJob).eventInformation.GetString(tag.PatientName); got != managedTestOriginalName {
		t.Fatalf("queued N-EVENT-REPORT PatientName = %q, want independent source copy", got)
	}
	if got := getJob.(*managedNGetJob).attributeIdentifierList[0]; got == tag.PatientName || *got != *tag.PatientName {
		t.Fatalf("queued N-GET attribute tag = %p %#v, want independent equivalent copy", got, got)
	}
}
