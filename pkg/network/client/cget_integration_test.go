// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package client

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/server"
	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func cgetIntegrationTransferSyntaxes() []string {
	return []string{testExplicitVRLittleEndianUID, testImplicitVRLittleEndianUID}
}

func cgetIntegrationDataset(sopClassUID string) *dataset.Dataset {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{sopClassUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"CGET^PROBE"}))
	return ds
}

// startCGetServer runs a QR SCP that answers a C-GET by pushing one C-STORE
// sub-operation for sopClassUID, mirroring a real retrieval SCP.
func startCGetServer(t *testing.T, sopClassUID string) int {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("reserve port: %v", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	_ = listener.Close()

	srv := server.New(server.WithPort(port))
	srv.SetCGetHandler(func(ctx context.Context, op service.CGetOperation) error {
		resp, err := op.SendCStore(ctx, cgetIntegrationDataset(sopClassUID))
		if err != nil {
			return op.SendWarning()
		}
		if !resp.Status().IsSuccess() {
			return op.SendWarning()
		}
		return op.SendSuccess()
	})

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		defer close(done)
		_ = srv.ListenAndServe(ctx)
	}()
	t.Cleanup(func() {
		cancel()
		select {
		case <-done:
		case <-time.After(2 * time.Second):
			t.Log("C-GET test server did not stop in time")
		}
	})

	deadline := time.Now().Add(3 * time.Second)
	for {
		conn, dialErr := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%d", port), 50*time.Millisecond)
		if dialErr == nil {
			_ = conn.Close()
			return port
		}
		if time.Now().After(deadline) {
			t.Fatalf("C-GET test server did not start: %v", dialErr)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func newCGetClient(cstoreHandler func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error)) *Client {
	options := []Option{
		WithCallingAE("CGET_SCU"),
		WithCalledAE("CGET_SCP"),
		WithAssociationTimeout(5 * time.Second),
	}
	if cstoreHandler != nil {
		options = append(options, WithCStoreHandler(cstoreHandler))
	}
	return New(options...)
}

func TestCGetReceivesCStoreWhenStorageSCPRoleNegotiated(t *testing.T) {
	port := startCGetServer(t, uid.CTImageStorage.UID())

	var received int32
	c := newCGetClient(func(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
		atomic.AddInt32(&received, 1)
		return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
	})
	if err := c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelGet.UID(), cgetIntegrationTransferSyntaxes()...); err != nil {
		t.Fatalf("AddPresentationContext(C-GET) error = %v", err)
	}
	if err := c.AddStoragePresentationContexts([]string{uid.CTImageStorage.UID()}, cgetIntegrationTransferSyntaxes()...); err != nil {
		t.Fatalf("AddStoragePresentationContexts() error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := c.Connect(ctx, "127.0.0.1", port); err != nil {
		t.Fatalf("Connect() error = %v", err)
	}
	defer func() { _ = c.Close() }()

	identifier := dataset.New()
	if err := identifier.AddValue(tag.StudyInstanceUID, "1.2.3.4"); err != nil {
		t.Fatalf("build identifier: %v", err)
	}
	var failed uint16
	if err := c.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier, func(_, _, fail, _ uint16) bool {
		failed = fail
		return true
	}); err != nil {
		t.Fatalf("CGet() error = %v", err)
	}
	if failed != 0 {
		t.Fatalf("CGet() reported failed sub-operations = %d, want 0", failed)
	}
	if got := atomic.LoadInt32(&received); got != 1 {
		t.Fatalf("received C-STORE sub-operations = %d, want 1", got)
	}
}

func TestCGetRejectsStorageWithoutSCPRole(t *testing.T) {
	port := startCGetServer(t, uid.CTImageStorage.UID())

	c := newCGetClient(func(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
		return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
	})
	if err := c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelGet.UID(), cgetIntegrationTransferSyntaxes()...); err != nil {
		t.Fatalf("AddPresentationContext(C-GET) error = %v", err)
	}
	// Plain storage context: no SCP role, so the SCP cannot push instances back.
	if err := c.AddPresentationContext(uid.CTImageStorage.UID(), cgetIntegrationTransferSyntaxes()...); err != nil {
		t.Fatalf("AddPresentationContext(storage) error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := c.Connect(ctx, "127.0.0.1", port); err != nil {
		t.Fatalf("Connect() error = %v", err)
	}
	defer func() { _ = c.Close() }()

	identifier := dataset.New()
	if err := identifier.AddValue(tag.StudyInstanceUID, "1.2.3.4"); err != nil {
		t.Fatalf("build identifier: %v", err)
	}
	err := c.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier, nil)
	if err == nil || !strings.Contains(err.Error(), "SCP role") {
		t.Fatalf("CGet() error = %v, want storage SCP role error", err)
	}
}

func TestCGetRejectsMissingCStoreHandler(t *testing.T) {
	port := startCGetServer(t, uid.CTImageStorage.UID())

	c := newCGetClient(nil)
	if err := c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelGet.UID(), cgetIntegrationTransferSyntaxes()...); err != nil {
		t.Fatalf("AddPresentationContext(C-GET) error = %v", err)
	}
	if err := c.AddStoragePresentationContexts([]string{uid.CTImageStorage.UID()}, cgetIntegrationTransferSyntaxes()...); err != nil {
		t.Fatalf("AddStoragePresentationContexts() error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := c.Connect(ctx, "127.0.0.1", port); err != nil {
		t.Fatalf("Connect() error = %v", err)
	}
	defer func() { _ = c.Close() }()

	identifier := dataset.New()
	if err := identifier.AddValue(tag.StudyInstanceUID, "1.2.3.4"); err != nil {
		t.Fatalf("build identifier: %v", err)
	}
	err := c.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier, nil)
	if err == nil || !strings.Contains(err.Error(), "CStoreHandler") {
		t.Fatalf("CGet() error = %v, want CStoreHandler error", err)
	}
}

func TestManagedCGetReceivesCStore(t *testing.T) {
	port := startCGetServer(t, uid.CTImageStorage.UID())

	var received int32
	managed := NewManaged(
		WithBaseClientOptions(
			WithCallingAE("MGMT_SCU"),
			WithCalledAE("MGMT_SCP"),
			WithAssociationTimeout(5*time.Second),
			WithCStoreHandler(func(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
				atomic.AddInt32(&received, 1)
				return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
			}),
		),
	)
	defer func() { _ = managed.Close() }()

	identifier := dataset.New()
	if err := identifier.AddValue(tag.StudyInstanceUID, "1.2.3.4"); err != nil {
		t.Fatalf("build identifier: %v", err)
	}
	completed := make(chan error, 1)
	job, err := NewCGetJob(dimse.QueryRetrieveLevelStudy, identifier, nil,
		func(err error) { completed <- err }, uid.CTImageStorage.UID())
	if err != nil {
		t.Fatalf("NewCGetJob() error = %v", err)
	}
	if err := managed.Add(job); err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := managed.Send(ctx, "127.0.0.1", port); err != nil {
		t.Fatalf("Send() error = %v", err)
	}
	select {
	case err := <-completed:
		if err != nil {
			t.Fatalf("managed C-GET job error = %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("managed C-GET job did not complete")
	}
	if got := atomic.LoadInt32(&received); got != 1 {
		t.Fatalf("received C-STORE sub-operations = %d, want 1", got)
	}
}
