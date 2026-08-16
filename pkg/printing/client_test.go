// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

type recordedPrintCall struct {
	operation string
	dataset   *dataset.Dataset
}

type fakeDIMSEService struct {
	calls           []recordedPrintCall
	failAt          int
	statusFailureAt int
	cancelAfter     int
	cancel          context.CancelFunc
}

func newFakeDIMSEService() *fakeDIMSEService {
	return &fakeDIMSEService{failAt: -1, statusFailureAt: -1, cancelAfter: -1}
}

func (f *fakeDIMSEService) record(operation string, ds *dataset.Dataset) (int, error) {
	index := len(f.calls)
	f.calls = append(f.calls, recordedPrintCall{operation: operation, dataset: ds})
	if f.cancel != nil && len(f.calls) == f.cancelAfter {
		f.cancel()
	}
	if index == f.failAt {
		return index, errors.New("operation failed")
	}
	return index, nil
}

func (f *fakeDIMSEService) SendNCreate(_ context.Context, req *dimse.NCreateRequest) (*dimse.NCreateResponse, error) {
	index, err := f.record(fmt.Sprintf("N-CREATE %s %s", req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID()), req.DataDataset())
	if err != nil {
		return nil, err
	}
	responseStatus := status.NCreateSuccess
	if index == f.statusFailureAt {
		responseStatus = status.NCreateFailureProcessingFailure
	}
	return dimse.NewNCreateResponse(req.MessageID(), responseStatus, req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID(), nil), nil
}

func (f *fakeDIMSEService) SendNSet(_ context.Context, req *dimse.NSetRequest) (*dimse.NSetResponse, error) {
	index, err := f.record(fmt.Sprintf("N-SET %s %s", req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID()), req.DataDataset())
	if err != nil {
		return nil, err
	}
	responseStatus := status.NSetSuccess
	if index == f.statusFailureAt {
		responseStatus = status.NSetFailureProcessingFailure
	}
	return dimse.NewNSetResponse(req.MessageID(), responseStatus, req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), nil), nil
}

func (f *fakeDIMSEService) SendNAction(_ context.Context, req *dimse.NActionRequest) (*dimse.NActionResponse, error) {
	index, err := f.record(fmt.Sprintf("N-ACTION %s %s %d", req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), req.ActionTypeID()), req.DataDataset())
	if err != nil {
		return nil, err
	}
	responseStatus := status.NActionSuccess
	if index == f.statusFailureAt {
		responseStatus = status.NActionFailureProcessingFailure
	}
	return dimse.NewNActionResponse(req.MessageID(), responseStatus, req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), req.ActionTypeID(), nil), nil
}

func (f *fakeDIMSEService) SendNDelete(_ context.Context, req *dimse.NDeleteRequest) (*dimse.NDeleteResponse, error) {
	index, err := f.record(fmt.Sprintf("N-DELETE %s %s", req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID()), nil)
	if err != nil {
		return nil, err
	}
	responseStatus := status.NDeleteSuccess
	if index == f.statusFailureAt {
		responseStatus = status.NDeleteFailureProcessingFailure
	}
	return dimse.NewNDeleteResponse(req.MessageID(), responseStatus, req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID()), nil
}

func printableSession(t *testing.T) *FilmSession {
	t.Helper()
	session := NewFilmSession(basicFilmSessionSOPClassUID, "2.25.501", false)
	session.FilmSessionLabel = "workflow"

	lut := NewPresentationLUT("2.25.502")
	if err := lut.SetLUT(2, 0, 12, []uint16{10, 20}); err != nil {
		t.Fatalf("SetLUT() error = %v", err)
	}
	session.AddPresentationLUT(lut)

	filmBox := NewFilmBox("2.25.503", `STANDARD\1,1`)
	imageBox := NewImageBox("2.25.504", false)
	imageBox.ImageBoxPosition = 1
	imageBox.SetImageData([]byte{1, 2, 3, 4})
	filmBox.AddImageBox(imageBox)
	session.AddFilmBox(filmBox)
	return session
}

func TestClientPrintSendsOrderedDatasetBackedWorkflow(t *testing.T) {
	service := newFakeDIMSEService()
	client := NewClient(service)
	if err := client.Print(context.Background(), printableSession(t)); err != nil {
		t.Fatalf("Print() error = %v", err)
	}

	wantOperations := []string{
		"N-CREATE 1.2.840.10008.5.1.1.1 2.25.501",
		"N-CREATE 1.2.840.10008.5.1.1.23 2.25.502",
		"N-CREATE 1.2.840.10008.5.1.1.2 2.25.503",
		"N-SET 1.2.840.10008.5.1.1.4 2.25.504",
		"N-ACTION 1.2.840.10008.5.1.1.1 2.25.501 1",
	}
	gotOperations := make([]string, len(service.calls))
	for index, call := range service.calls {
		gotOperations[index] = call.operation
	}
	if !reflect.DeepEqual(gotOperations, wantOperations) {
		t.Fatalf("operation order = %#v, want %#v", gotOperations, wantOperations)
	}

	if got, _ := service.calls[0].dataset.GetString(tag.FilmSessionLabel); got != "workflow" {
		t.Errorf("Film Session payload label = %q", got)
	}
	if got, err := service.calls[1].dataset.GetUInt16s(tag.LUTDescriptor); err != nil || !reflect.DeepEqual(got, []uint16{2, 0, 12}) {
		t.Errorf("Presentation LUT descriptor = %v, %v", got, err)
	}
	if got, _ := service.calls[2].dataset.GetString(tag.ImageDisplayFormat); got != `STANDARD\1,1` {
		t.Errorf("Film Box payload format = %q", got)
	}
	if got, err := service.calls[3].dataset.GetUInt16(tag.ImageBoxPosition, 0); err != nil || got != 1 {
		t.Errorf("Image Box payload position = %d, %v", got, err)
	}
}

func TestClientPrintStopsAtFirstTransportError(t *testing.T) {
	for failAt := 0; failAt < 5; failAt++ {
		t.Run(fmt.Sprintf("operation_%d", failAt), func(t *testing.T) {
			service := newFakeDIMSEService()
			service.failAt = failAt
			err := NewClient(service).Print(context.Background(), printableSession(t))
			if err == nil || !strings.Contains(err.Error(), "operation failed") {
				t.Fatalf("Print() error = %v, want wrapped transport failure", err)
			}
			if len(service.calls) != failAt+1 {
				t.Fatalf("call count = %d, want %d", len(service.calls), failAt+1)
			}
		})
	}
}

func TestClientPrintRejectsFailureStatus(t *testing.T) {
	service := newFakeDIMSEService()
	service.statusFailureAt = 3
	err := NewClient(service).Print(context.Background(), printableSession(t))
	if err == nil || !strings.Contains(err.Error(), "N-SET") || !strings.Contains(err.Error(), "2.25.504") {
		t.Fatalf("Print() error = %v, want N-SET UID context", err)
	}
	if len(service.calls) != 4 {
		t.Fatalf("call count = %d, want 4", len(service.calls))
	}
}

func TestClientPrintHonorsCancellation(t *testing.T) {
	t.Run("before first request", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		service := newFakeDIMSEService()
		err := NewClient(service).Print(ctx, printableSession(t))
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("Print() error = %v, want context.Canceled", err)
		}
		if len(service.calls) != 0 {
			t.Fatalf("call count = %d, want 0", len(service.calls))
		}
	})

	t.Run("between requests", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		service := newFakeDIMSEService()
		service.cancel = cancel
		service.cancelAfter = 1
		err := NewClient(service).Print(ctx, printableSession(t))
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("Print() error = %v, want context.Canceled", err)
		}
		if len(service.calls) != 1 {
			t.Fatalf("call count = %d, want 1", len(service.calls))
		}
	})
}
