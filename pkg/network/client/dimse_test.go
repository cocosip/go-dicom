// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"errors"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func TestDIMSEMethodsRejectDisconnectedClient(t *testing.T) {
	client := New()
	identifier := dataset.New()
	tests := []struct {
		name string
		run  func() error
	}{
		{"C-ECHO", func() error { return client.CEcho(context.Background()) }},
		{"C-STORE", func() error { return client.CStore(context.Background(), identifier) }},
		{"C-FIND", func() error {
			_, err := client.CFind(context.Background(), dimse.QueryRetrieveLevelStudy, identifier)
			return err
		}},
		{"C-CANCEL", func() error { return client.CCancel(context.Background(), 1, 1) }},
		{"C-MOVE", func() error {
			return client.CMove(context.Background(), dimse.QueryRetrieveLevelStudy, "DEST", identifier, nil)
		}},
		{"C-GET", func() error {
			return client.CGet(context.Background(), dimse.QueryRetrieveLevelStudy, identifier, nil)
		}},
		{"C-STORE priority", func() error {
			return client.CStoreWithPriority(context.Background(), identifier, dimse.PriorityHigh)
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.run(); !errors.Is(err, ErrClientNotConnected) {
				t.Fatalf("error = %v, want ErrClientNotConnected", err)
			}
		})
	}
}

func TestDIMSEMethodsValidateInputsBeforeSending(t *testing.T) {
	client := newActiveTestClient(t)
	tests := []struct {
		name string
		want string
		run  func() error
	}{
		{"C-STORE nil Dataset", "dataset is nil", func() error {
			return client.CStore(context.Background(), nil)
		}},
		{"C-FIND nil query", "query is nil", func() error {
			_, err := client.CFind(context.Background(), dimse.QueryRetrieveLevelStudy, nil)
			return err
		}},
		{"C-FIND nil callback", "callback is nil", func() error {
			return client.CFindWithCallback(context.Background(), dimse.QueryRetrieveLevelStudy, dataset.New(), nil)
		}},
		{"C-MOVE nil identifier", "identifier is nil", func() error {
			return client.CMove(context.Background(), dimse.QueryRetrieveLevelStudy, "DEST", nil, nil)
		}},
		{"C-MOVE empty destination", "move destination is empty", func() error {
			return client.CMove(context.Background(), dimse.QueryRetrieveLevelStudy, "", dataset.New(), nil)
		}},
		{"C-GET nil identifier", "identifier is nil", func() error {
			return client.CGet(context.Background(), dimse.QueryRetrieveLevelStudy, nil, nil)
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.run(); err == nil || err.Error() != test.want {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestNServiceMethodsRejectNilRequests(t *testing.T) {
	client := newActiveTestClient(t)
	tests := []struct {
		name string
		run  func() error
	}{
		{"N-CREATE", func() error { _, err := client.NCreate(context.Background(), nil); return err }},
		{"N-GET", func() error { _, err := client.NGet(context.Background(), nil); return err }},
		{"N-SET", func() error { _, err := client.NSet(context.Background(), nil); return err }},
		{"N-DELETE", func() error { _, err := client.NDelete(context.Background(), nil); return err }},
		{"N-ACTION", func() error { _, err := client.NAction(context.Background(), nil); return err }},
		{"N-EVENT-REPORT", func() error { _, err := client.NEventReport(context.Background(), nil); return err }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.run(); err == nil {
				t.Fatalf("%s nil request error = nil", test.name)
			}
		})
	}
}

func TestCStoreMultipleEmptyDoesNotSend(t *testing.T) {
	client := newActiveTestClient(t)
	count, err := client.CStoreMultiple(context.Background(), nil)
	if err != nil {
		t.Fatalf("CStoreMultiple() error = %v", err)
	}
	if count != 0 {
		t.Fatalf("CStoreMultiple() count = %d, want 0", count)
	}
}

func TestProgressFinalStatusRules(t *testing.T) {
	if err := cFindFinalStatusError(status.Success.Code, false); err != nil {
		t.Fatalf("C-FIND Success error = %v", err)
	}
	if err := cFindFinalStatusError(status.Cancel.Code, true); err != nil {
		t.Fatalf("cancelled C-FIND final Cancel error = %v", err)
	}
	if err := cMoveFinalStatusError(status.CMoveWarningSubOperationsComplete.Code, false); err != nil {
		t.Fatalf("C-MOVE warning error = %v", err)
	}
	if err := cGetFinalStatusError(status.CGetWarningSubOperationsComplete.Code, false); err != nil {
		t.Fatalf("C-GET warning error = %v", err)
	}
	if err := cFindFinalStatusError(status.CFindFailedUnableToProcess.Code, false); err == nil {
		t.Fatal("C-FIND failure status error = nil")
	}
	if err := cMoveFinalStatusError(status.CMoveFailedUnableToProcess.Code, false); err == nil {
		t.Fatal("C-MOVE failure status error = nil")
	}
	if err := cGetFinalStatusError(status.CGetFailedUnableToProcess.Code, false); err == nil {
		t.Fatal("C-GET failure status error = nil")
	}
}

func TestCEchoTreatsMissingServiceAsDisconnected(t *testing.T) {
	client := New()
	client.connected = true
	client.state = clientConnected

	if err := client.CEcho(context.Background()); !errors.Is(err, ErrClientNotConnected) {
		t.Fatalf("CEcho() error = %v, want ErrClientNotConnected", err)
	}
}

func newActiveTestClient(t *testing.T) *Client {
	t.Helper()
	svc := service.NewService(nil, nil)
	t.Cleanup(func() { _ = svc.Close() })
	client := New()
	client.connected = true
	client.state = clientConnected
	client.service = svc
	return client
}
