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
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
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
	filmBoxResponse *dataset.Dataset
}

func newFakeDIMSEService() *fakeDIMSEService {
	return &fakeDIMSEService{
		failAt:          -1,
		statusFailureAt: -1,
		cancelAfter:     -1,
		filmBoxResponse: referencedImageBoxResponseDataset([]SOPReference{{
			SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: testRemoteImageBoxUID,
		}}),
	}
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
	var responseDataset *dataset.Dataset
	if req.AffectedSOPClassUID() == basicFilmBoxSOPClassUID {
		responseDataset = f.filmBoxResponse
	}
	return dimse.NewNCreateResponse(req.MessageID(), responseStatus, req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID(), responseDataset), nil
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

	filmBox := NewFilmBox("2.25.503", testStandardOneByOne)
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
		"N-SET 1.2.840.10008.5.1.1.4 2.25.900",
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
	lutSequence, err := service.calls[1].dataset.GetSequence(tag.PresentationLUTSequence)
	if err != nil || lutSequence.Count() != 1 {
		t.Fatalf("Presentation LUT sequence = %#v, %v", lutSequence, err)
	}
	if got, err := lutSequence.GetItem(0).GetUInt16s(tag.LUTDescriptor); err != nil || !reflect.DeepEqual(got, []uint16{2, 0, 12}) {
		t.Errorf("Presentation LUT descriptor = %v, %v", got, err)
	}
	if got, _ := service.calls[2].dataset.GetString(tag.ImageDisplayFormat); got != testStandardOneByOne {
		t.Errorf("Film Box payload format = %q", got)
	}
	if _, ok := service.calls[2].dataset.Get(tag.ReferencedImageBoxSequence); ok {
		t.Error("Film Box N-CREATE request contains SCP-produced Referenced Image Box Sequence")
	}
	assertReferenceSequence(t, service.calls[2].dataset, tag.ReferencedFilmSessionSequence, []SOPReference{{
		SOPClassUID: basicFilmSessionSOPClassUID, SOPInstanceUID: "2.25.501",
	}})
	if got, err := service.calls[3].dataset.GetUInt16(tag.ImageBoxPosition, 0); err != nil || got != 1 {
		t.Errorf("Image Box payload position = %d, %v", got, err)
	}
	if got, ok := service.calls[3].dataset.GetString(tag.SOPInstanceUID); !ok || got != testRemoteImageBoxUID {
		t.Errorf("Image Box payload SOP Instance UID = %q, %v; want remote UID", got, ok)
	}
}

func TestClientPrintKeepsLocalImageBoxUIDAfterRemoteMapping(t *testing.T) {
	session := printableSession(t)
	image := session.FindImageBox("2.25.504")
	service := newFakeDIMSEService()
	if err := NewClient(service).Print(context.Background(), session); err != nil {
		t.Fatalf("Print() error = %v", err)
	}
	if image.SOPInstanceUID != "2.25.504" {
		t.Fatalf("local Image Box UID = %q, want unchanged", image.SOPInstanceUID)
	}
}

func TestClientPrintRejectsMalformedImageBoxReferences(t *testing.T) {
	malformedItem := dataset.New()
	_ = malformedItem.Add(element.NewString(tag.ReferencedSOPClassUID, vr.UI, []string{SOPClassGrayscaleImageBox}))
	malformedDataset := dataset.New()
	_ = malformedDataset.Add(dataset.NewSequenceWithItems(tag.ReferencedImageBoxSequence, []*dataset.Dataset{malformedItem}))

	for _, test := range []struct {
		name     string
		response *dataset.Dataset
	}{
		{name: "missing attribute list", response: nil},
		{name: "missing sequence", response: dataset.New()},
		{name: "empty sequence", response: referencedImageBoxResponseDataset(nil)},
		{name: "malformed item", response: malformedDataset},
		{name: "wrong count", response: referencedImageBoxResponseDataset([]SOPReference{
			{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: "2.25.901"},
			{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: "2.25.902"},
		})},
		{name: "incompatible SOP class", response: referencedImageBoxResponseDataset([]SOPReference{{
			SOPClassUID: SOPClassColorImageBox, SOPInstanceUID: "2.25.903",
		}})},
	} {
		t.Run(test.name, func(t *testing.T) {
			service := newFakeDIMSEService()
			service.filmBoxResponse = test.response
			err := NewClient(service).Print(context.Background(), printableSession(t))
			if err == nil || !strings.Contains(err.Error(), "referenced Image Box") {
				t.Fatalf("Print() error = %v, want Referenced Image Box context", err)
			}
			if len(service.calls) != 3 {
				t.Fatalf("call count = %d, want stop after Film Box N-CREATE", len(service.calls))
			}
		})
	}
}

func TestClientPrintRejectsDuplicateRemoteImageBoxUID(t *testing.T) {
	session := printableSession(t)
	box := session.BasicFilmBoxes[0]
	second := NewImageBox("2.25.505", false)
	second.ImageBoxPosition = 2
	second.SetImageData([]byte{5, 6, 7, 8})
	box.AddImageBox(second)
	box.ImageDisplayFormat = testStandardTwoByOne

	service := newFakeDIMSEService()
	service.filmBoxResponse = referencedImageBoxResponseDataset([]SOPReference{
		{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: "2.25.904"},
		{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: "2.25.904"},
	})
	err := NewClient(service).Print(context.Background(), session)
	if err == nil || !strings.Contains(err.Error(), "duplicate") {
		t.Fatalf("Print() error = %v, want duplicate UID error", err)
	}
	if len(service.calls) != 3 {
		t.Fatalf("call count = %d, want stop after Film Box N-CREATE", len(service.calls))
	}
}

func TestClientPrintRejectsDuplicateLocalImageBoxPosition(t *testing.T) {
	session := printableSession(t)
	second := NewImageBox("2.25.505", false)
	second.ImageBoxPosition = 1
	session.BasicFilmBoxes[0].AddImageBox(second)
	session.BasicFilmBoxes[0].ImageDisplayFormat = testStandardTwoByOne
	service := newFakeDIMSEService()

	err := NewClient(service).Print(context.Background(), session)
	if err == nil || !strings.Contains(err.Error(), "duplicate Image Box position 1") {
		t.Fatalf("Print() error = %v, want duplicate position error", err)
	}
	if len(service.calls) != 0 {
		t.Fatalf("call count = %d, want validation before network operations", len(service.calls))
	}
}

func TestClientPrintRejectsOutOfRangeLocalImageBoxPosition(t *testing.T) {
	session := printableSession(t)
	session.BasicFilmBoxes[0].BasicImageBoxes[0].ImageBoxPosition = 2
	service := newFakeDIMSEService()

	err := NewClient(service).Print(context.Background(), session)
	if err == nil || !strings.Contains(err.Error(), "Image Box position 2 outside 1..1") {
		t.Fatalf("Print() error = %v, want position range error", err)
	}
	if len(service.calls) != 0 {
		t.Fatalf("call count = %d, want validation before network operations", len(service.calls))
	}
}

func TestClientPrintRejectsImageBoxCountMismatchBeforeNetworkOperations(t *testing.T) {
	session := printableSession(t)
	session.BasicFilmBoxes[0].ImageDisplayFormat = testStandardTwoByOne
	service := newFakeDIMSEService()

	err := NewClient(service).Print(context.Background(), session)
	if err == nil || !strings.Contains(err.Error(), "requires 2 Image Boxes, got 1") {
		t.Fatalf("Print() error = %v, want Image Display Format count error", err)
	}
	if len(service.calls) != 0 {
		t.Fatalf("call count = %d, want validation before network operations", len(service.calls))
	}
}

func TestClientPrintDefersConfigurationDependentImageBoxCounts(t *testing.T) {
	for _, displayFormat := range []string{slideImageDisplayFormat, "SUPERSLIDE", `CUSTOM\7`} {
		t.Run(displayFormat, func(t *testing.T) {
			session := printableSession(t)
			filmBox := session.BasicFilmBoxes[0]
			filmBox.ImageDisplayFormat = displayFormat
			second := NewImageBox("2.25.505", false)
			second.ImageBoxPosition = 2
			second.SetImageData([]byte{5, 6, 7, 8})
			filmBox.AddImageBox(second)

			service := newFakeDIMSEService()
			service.filmBoxResponse = referencedImageBoxResponseDataset([]SOPReference{
				{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: testRemoteImageBoxUID},
				{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: "2.25.901"},
			})

			if err := NewClient(service).Print(context.Background(), session); err != nil {
				t.Fatalf("Print() error = %v", err)
			}
			if len(service.calls) != 6 {
				t.Fatalf("call count = %d, want complete workflow", len(service.calls))
			}
		})
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
	if err == nil || !strings.Contains(err.Error(), "N-SET") || !strings.Contains(err.Error(), "2.25.900") {
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

func referencedImageBoxResponseDataset(references []SOPReference) *dataset.Dataset {
	ds := dataset.New()
	items := make([]*dataset.Dataset, 0, len(references))
	for _, reference := range references {
		item := dataset.New()
		_ = item.Add(element.NewString(tag.ReferencedSOPClassUID, vr.UI, []string{reference.SOPClassUID}))
		_ = item.Add(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{reference.SOPInstanceUID}))
		items = append(items, item)
	}
	_ = ds.Add(dataset.NewSequenceWithItems(tag.ReferencedImageBoxSequence, items))
	return ds
}
