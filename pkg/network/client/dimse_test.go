// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

const errNotConnected = "client not connected"

func TestCEcho_NotConnected(t *testing.T) {
	client := New()

	ctx := context.Background()
	err := client.CEcho(ctx)
	if err == nil {
		t.Error("Expected error when not connected")
	}

	expectedMsg := errNotConnected
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCStore_NotConnected(t *testing.T) {
	client := New()

	ds := dataset.New()
	ctx := context.Background()
	err := client.CStore(ctx, ds)
	if err == nil {
		t.Error("Expected error when not connected")
	}

	expectedMsg := errNotConnected
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCStore_NilDataset(t *testing.T) {
	client, _ := setupMockClient()

	ctx := context.Background()
	err := client.CStore(ctx, nil)
	if err == nil {
		t.Error("Expected error with nil dataset")
	}

	expectedMsg := "dataset is nil"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCFind_NotConnected(t *testing.T) {
	client := New()

	query := dataset.New()
	ctx := context.Background()
	_, err := client.CFind(ctx, dimse.QueryRetrieveLevelStudy, query)
	if err == nil {
		t.Error("Expected error when not connected")
	}

	expectedMsg := errNotConnected
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCCancel_NotConnected(t *testing.T) {
	client := New()

	err := client.CCancel(context.Background(), 123, 1)
	if err == nil {
		t.Error("Expected error when not connected")
	}

	expectedMsg := errNotConnected
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCFind_NilQuery(t *testing.T) {
	client, _ := setupMockClient()

	ctx := context.Background()
	_, err := client.CFind(ctx, dimse.QueryRetrieveLevelStudy, nil)
	if err == nil {
		t.Error("Expected error with nil query")
	}

	expectedMsg := "query is nil"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCFindWithCallback_NilCallback(t *testing.T) {
	client, _ := setupMockClient()

	query := dataset.New()
	ctx := context.Background()
	err := client.CFindWithCallback(ctx, dimse.QueryRetrieveLevelStudy, query, nil)
	if err == nil {
		t.Error("Expected error with nil callback")
	}

	expectedMsg := "callback is nil"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCStoreWithPriority_NotConnected(t *testing.T) {
	client := New()

	ds := dataset.New()
	ctx := context.Background()
	err := client.CStoreWithPriority(ctx, ds, dimse.PriorityHigh)
	if err == nil {
		t.Error("Expected error when not connected")
	}
}

// Mock service for testing DIMSE operations
type mockServiceForDIMSE struct {
	echoResponse      *dimse.CEchoResponse
	storeResponse     *dimse.CStoreResponse
	storeHook         func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error)
	findResponses     []*dimse.CFindResponse
	findResponseCh    <-chan *dimse.CFindResponse
	findHook          func(*dimse.CFindRequest)
	findTerminalErr   error
	moveResponseCh    <-chan *dimse.CMoveResponse
	moveHook          func(*dimse.CMoveRequest)
	moveTerminalErr   error
	getResponseCh     <-chan *dimse.CGetResponse
	getHook           func(*dimse.CGetRequest)
	getTerminalErr    error
	cancelHook        func(messageID uint16, presentationContextID byte)
	cancelContextHook func(context.Context, uint16, byte)
	cancelErr         error
	nServiceRequests  []dimse.Request
	releaseStarted    chan struct{}
	releaseContinue   chan struct{}
}

// Association management methods (not used in these tests)
func (m *mockServiceForDIMSE) GetAssociation() *association.Association  { return nil }
func (m *mockServiceForDIMSE) SetAssociation(_ *association.Association) {}
func (m *mockServiceForDIMSE) SendAssociationRequest(_ context.Context, _ *pdu.AAssociateRQ) error {
	return nil
}
func (m *mockServiceForDIMSE) ReceiveAssociationResponse(_ context.Context) (*pdu.AAssociateAC, error) {
	return nil, nil
}
func (m *mockServiceForDIMSE) Start() error { return nil }
func (m *mockServiceForDIMSE) GracefulRelease(_ context.Context) error {
	continueCh := m.releaseContinue
	if m.releaseStarted != nil {
		close(m.releaseStarted)
	}
	if continueCh != nil {
		<-continueCh
	}
	return nil
}
func (m *mockServiceForDIMSE) Abort(_ context.Context, _ byte, _ byte) error { return nil }

func (m *mockServiceForDIMSE) SendCEcho(_ context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
	// Simulate processing delay
	time.Sleep(10 * time.Millisecond)

	if m.echoResponse != nil {
		return m.echoResponse, nil
	}

	// Default success response
	return dimse.NewCEchoResponseFromRequest(req, status.Success), nil
}

func (m *mockServiceForDIMSE) SendCStore(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
	// Simulate processing delay
	time.Sleep(10 * time.Millisecond)

	if m.storeHook != nil {
		return m.storeHook(ctx, req)
	}

	if m.storeResponse != nil {
		return m.storeResponse, nil
	}

	// Default success response
	return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
}

func (m *mockServiceForDIMSE) SendCFind(_ context.Context, req *dimse.CFindRequest) (<-chan *dimse.CFindResponse, error) {
	if req.MessageID() == 0 {
		_ = req.SetMessageID(101)
	}
	req.SetPresentationContextID(3)
	if m.findResponseCh != nil {
		if m.findHook != nil {
			m.findHook(req)
		}
		return m.findResponseCh, nil
	}
	// Create response channel
	respCh := make(chan *dimse.CFindResponse, 16)
	if m.findHook != nil {
		m.findHook(req)
	}

	// Simulate async responses
	go func() {
		defer close(respCh)
		time.Sleep(10 * time.Millisecond)

		if len(m.findResponses) > 0 {
			for _, resp := range m.findResponses {
				respCh <- resp
				time.Sleep(5 * time.Millisecond)
			}
		} else {
			// Default success response (no results)
			resp := dimse.NewCFindResponseFromRequest(req, status.Success, nil)
			respCh <- resp
		}
	}()

	return respCh, nil
}

func (m *mockServiceForDIMSE) SendCFindWithError(ctx context.Context, req *dimse.CFindRequest) (<-chan *dimse.CFindResponse, <-chan error, error) {
	responses, err := m.SendCFind(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	terminalErrors := make(chan error, 1)
	go func() {
		defer close(terminalErrors)
		if m.findTerminalErr != nil {
			terminalErrors <- m.findTerminalErr
		}
	}()
	return responses, terminalErrors, nil
}

func (m *mockServiceForDIMSE) SendCMove(_ context.Context, req *dimse.CMoveRequest) (<-chan *dimse.CMoveResponse, error) {
	if req.MessageID() == 0 {
		_ = req.SetMessageID(102)
	}
	req.SetPresentationContextID(5)
	if m.moveHook != nil {
		m.moveHook(req)
	}
	if m.moveResponseCh != nil {
		return m.moveResponseCh, nil
	}
	ch := make(chan *dimse.CMoveResponse, 16)
	go func() {
		defer close(ch)
		msgID, sopClass := req.MessageID(), req.AffectedSOPClassUID()
		time.Sleep(10 * time.Millisecond)
		ch <- dimse.NewCMoveResponsePending(msgID, sopClass, 5, 0, 0, 0) // 5 remaining
		time.Sleep(5 * time.Millisecond)
		ch <- dimse.NewCMoveResponsePending(msgID, sopClass, 3, 2, 0, 0) // 3 remaining, 2 completed
		time.Sleep(5 * time.Millisecond)
		ch <- dimse.NewCMoveResponsePending(msgID, sopClass, 0, 5, 0, 0) // All 5 completed
		time.Sleep(5 * time.Millisecond)
		ch <- dimse.NewCMoveResponseSuccess(msgID, sopClass)
	}()
	return ch, nil
}

func (m *mockServiceForDIMSE) SendCMoveWithError(ctx context.Context, req *dimse.CMoveRequest) (<-chan *dimse.CMoveResponse, <-chan error, error) {
	responses, err := m.SendCMove(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	terminalErrors := make(chan error, 1)
	if m.moveTerminalErr != nil {
		terminalErrors <- m.moveTerminalErr
	}
	close(terminalErrors)
	return responses, terminalErrors, nil
}

func (m *mockServiceForDIMSE) SendCGet(_ context.Context, req *dimse.CGetRequest) (<-chan *dimse.CGetResponse, error) {
	if req.MessageID() == 0 {
		_ = req.SetMessageID(103)
	}
	req.SetPresentationContextID(7)
	if m.getHook != nil {
		m.getHook(req)
	}
	if m.getResponseCh != nil {
		return m.getResponseCh, nil
	}
	ch := make(chan *dimse.CGetResponse, 16)
	go func() {
		defer close(ch)
		msgID, sopClass := req.MessageID(), req.AffectedSOPClassUID()
		time.Sleep(10 * time.Millisecond)
		ch <- dimse.NewCGetResponsePending(msgID, sopClass, 3, 0, 0, 0) // 3 remaining
		time.Sleep(5 * time.Millisecond)
		ch <- dimse.NewCGetResponsePending(msgID, sopClass, 1, 2, 0, 0) // 1 remaining, 2 completed
		time.Sleep(5 * time.Millisecond)
		ch <- dimse.NewCGetResponsePending(msgID, sopClass, 0, 3, 0, 0) // All 3 completed
		time.Sleep(5 * time.Millisecond)
		ch <- dimse.NewCGetResponseSuccess(msgID, sopClass)
	}()
	return ch, nil
}

func (m *mockServiceForDIMSE) SendCGetWithError(ctx context.Context, req *dimse.CGetRequest) (<-chan *dimse.CGetResponse, <-chan error, error) {
	responses, err := m.SendCGet(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	terminalErrors := make(chan error, 1)
	if m.getTerminalErr != nil {
		terminalErrors <- m.getTerminalErr
	}
	close(terminalErrors)
	return responses, terminalErrors, nil
}

func (m *mockServiceForDIMSE) SendCCancel(ctx context.Context, messageID uint16, presentationContextID byte) error {
	if m.cancelHook != nil {
		m.cancelHook(messageID, presentationContextID)
	}
	if m.cancelContextHook != nil {
		m.cancelContextHook(ctx, messageID, presentationContextID)
	}
	return m.cancelErr
}

func (m *mockServiceForDIMSE) SendNCreate(_ context.Context, req *dimse.NCreateRequest) (*dimse.NCreateResponse, error) {
	m.nServiceRequests = append(m.nServiceRequests, req)
	return dimse.NewNCreateResponseSuccess(req.MessageID(), req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID(), nil), nil
}

func (m *mockServiceForDIMSE) SendNGet(_ context.Context, req *dimse.NGetRequest) (*dimse.NGetResponse, error) {
	m.nServiceRequests = append(m.nServiceRequests, req)
	return dimse.NewNGetResponseSuccess(req.MessageID(), req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), nil), nil
}

func (m *mockServiceForDIMSE) SendNSet(_ context.Context, req *dimse.NSetRequest) (*dimse.NSetResponse, error) {
	m.nServiceRequests = append(m.nServiceRequests, req)
	return dimse.NewNSetResponseSuccess(req.MessageID(), req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), nil), nil
}

func (m *mockServiceForDIMSE) SendNDelete(_ context.Context, req *dimse.NDeleteRequest) (*dimse.NDeleteResponse, error) {
	m.nServiceRequests = append(m.nServiceRequests, req)
	return dimse.NewNDeleteResponseSuccess(req.MessageID(), req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID()), nil
}

func (m *mockServiceForDIMSE) SendNAction(_ context.Context, req *dimse.NActionRequest) (*dimse.NActionResponse, error) {
	m.nServiceRequests = append(m.nServiceRequests, req)
	return dimse.NewNActionResponseSuccess(req.MessageID(), req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), req.ActionTypeID(), nil), nil
}

func (m *mockServiceForDIMSE) SendNEventReport(_ context.Context, req *dimse.NEventReportRequest) (*dimse.NEventReportResponse, error) {
	m.nServiceRequests = append(m.nServiceRequests, req)
	return dimse.NewNEventReportResponseSuccess(req.MessageID(), req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID(), req.EventTypeID(), nil), nil
}

func setupMockClient() (*Client, *mockServiceForDIMSE) {
	client := New()
	client.connected = true
	client.assoc = association.NewAssociation("TEST_SCU", "TEST_SCP")
	mockService := &mockServiceForDIMSE{}
	client.service = mockService
	return client, mockService
}

func TestCEchoTreatsMissingServiceAsDisconnected(t *testing.T) {
	client := New()
	client.connected = true

	if err := client.CEcho(context.Background()); err == nil {
		t.Fatal("CEcho() succeeded with no active service")
	}
}

func TestNServiceMethodsReturnRawResponses(t *testing.T) {
	client, mock := setupMockClient()
	const sopClassUID = "1.2.840.10008.5.1.1.1"
	const sopInstanceUID = "2.25.701"

	tests := []struct {
		name string
		run  func() (dimse.Response, error)
	}{
		{"N-CREATE", func() (dimse.Response, error) {
			return client.NCreate(context.Background(), dimse.NewNCreateRequest(sopClassUID, sopInstanceUID, dataset.New()))
		}},
		{"N-GET", func() (dimse.Response, error) {
			return client.NGet(context.Background(), dimse.NewNGetRequest(sopClassUID, sopInstanceUID, nil))
		}},
		{"N-SET", func() (dimse.Response, error) {
			return client.NSet(context.Background(), dimse.NewNSetRequest(sopClassUID, sopInstanceUID, dataset.New()))
		}},
		{"N-DELETE", func() (dimse.Response, error) {
			return client.NDelete(context.Background(), dimse.NewNDeleteRequest(sopClassUID, sopInstanceUID))
		}},
		{"N-ACTION", func() (dimse.Response, error) {
			return client.NAction(context.Background(), dimse.NewNActionRequest(sopClassUID, sopInstanceUID, 1, dataset.New()))
		}},
		{"N-EVENT-REPORT", func() (dimse.Response, error) {
			return client.NEventReport(context.Background(), dimse.NewNEventReportRequest(sopClassUID, sopInstanceUID, 1, dataset.New()))
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := tt.run()
			if err != nil {
				t.Fatalf("%s error = %v", tt.name, err)
			}
			if response == nil || !response.Status().IsSuccess() {
				t.Fatalf("%s response = %#v, want raw success response", tt.name, response)
			}
		})
	}
	if got := len(mock.nServiceRequests); got != len(tests) {
		t.Fatalf("N-Service request count = %d, want %d", got, len(tests))
	}
}

func TestClosePublishesDisconnectedStateBeforeGracefulReleaseCompletes(t *testing.T) {
	client, mockService := setupMockClient()
	mockService.releaseStarted = make(chan struct{})
	mockService.releaseContinue = make(chan struct{})
	defer func() {
		if mockService.releaseContinue != nil {
			close(mockService.releaseContinue)
		}
	}()

	closeDone := make(chan error, 1)
	go func() { closeDone <- client.Close() }()
	select {
	case <-mockService.releaseStarted:
	case <-time.After(time.Second):
		t.Fatal("Close() did not start graceful release")
	}

	connected := make(chan bool, 1)
	go func() { connected <- client.IsConnected() }()
	select {
	case got := <-connected:
		if got {
			t.Fatal("client remained connected while graceful release was in progress")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("IsConnected() blocked behind graceful release")
	}

	close(mockService.releaseContinue)
	mockService.releaseContinue = nil
	if err := <-closeDone; err != nil {
		t.Fatalf("Close() error = %v", err)
	}
}

func TestCEcho_Success(t *testing.T) {
	client, _ := setupMockClient()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := client.CEcho(ctx)
	if err != nil {
		t.Errorf("CEcho failed: %v", err)
	}
}

func TestCStore_Success(t *testing.T) {
	client, _ := setupMockClient()

	// Create a valid dataset
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})) // CT Image Storage
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := client.CStore(ctx, ds)
	if err != nil {
		t.Errorf("CStore failed: %v", err)
	}
}

func TestCFind_Success_NoResults(t *testing.T) {
	client, _ := setupMockClient()

	query := dataset.New()
	_ = query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	results, err := client.CFind(ctx, dimse.QueryRetrieveLevelPatient, query)
	if err != nil {
		t.Errorf("CFind failed: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("Expected 0 results, got %d", len(results))
	}
}

func TestCFind_PropagatesAsynchronousRequestTimeout(t *testing.T) {
	client, mockService := setupMockClient()
	mockService.findTerminalErr = service.ErrRequestTimeout

	_, err := client.CFind(context.Background(), dimse.QueryRetrieveLevelStudy, dataset.New())
	if !errors.Is(err, service.ErrRequestTimeout) {
		t.Fatalf("CFind() error = %v, want ErrRequestTimeout", err)
	}
}

func TestCCancel_Success(t *testing.T) {
	client, mockService := setupMockClient()

	captured := make(chan struct {
		messageID             uint16
		presentationContextID byte
	}, 1)
	mockService.cancelHook = func(messageID uint16, presentationContextID byte) {
		captured <- struct {
			messageID             uint16
			presentationContextID byte
		}{messageID: messageID, presentationContextID: presentationContextID}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := client.CCancel(ctx, 123, 5); err != nil {
		t.Fatalf("CCancel() error = %v", err)
	}

	select {
	case got := <-captured:
		if got.messageID != 123 {
			t.Fatalf("messageID = %d, want 123", got.messageID)
		}
		if got.presentationContextID != 5 {
			t.Fatalf("presentationContextID = %d, want 5", got.presentationContextID)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-CANCEL send")
	}
}

func TestProgressiveContextCancellationSendsIndependentCCancel(t *testing.T) {
	tests := []struct {
		name      string
		want      [2]uint16
		configure func(*mockServiceForDIMSE, chan<- struct{}) func()
		run       func(context.Context, *Client) error
	}{
		{
			name: "C-FIND",
			want: [2]uint16{101, 3},
			configure: func(mock *mockServiceForDIMSE, started chan<- struct{}) func() {
				responses := make(chan *dimse.CFindResponse)
				mock.findResponseCh = responses
				mock.findHook = func(*dimse.CFindRequest) { close(started) }
				return func() { close(responses) }
			},
			run: func(ctx context.Context, client *Client) error {
				_, err := client.CFind(ctx, dimse.QueryRetrieveLevelStudy, dataset.New())
				return err
			},
		},
		{
			name: "C-MOVE",
			want: [2]uint16{102, 5},
			configure: func(mock *mockServiceForDIMSE, started chan<- struct{}) func() {
				responses := make(chan *dimse.CMoveResponse)
				mock.moveResponseCh = responses
				mock.moveHook = func(*dimse.CMoveRequest) { close(started) }
				return func() { close(responses) }
			},
			run: func(ctx context.Context, client *Client) error {
				return client.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST_AE", dataset.New(), nil)
			},
		},
		{
			name: "C-GET",
			want: [2]uint16{103, 7},
			configure: func(mock *mockServiceForDIMSE, started chan<- struct{}) func() {
				responses := make(chan *dimse.CGetResponse)
				mock.getResponseCh = responses
				mock.getHook = func(*dimse.CGetRequest) { close(started) }
				return func() { close(responses) }
			},
			run: func(ctx context.Context, client *Client) error {
				return client.CGet(ctx, dimse.QueryRetrieveLevelStudy, dataset.New(), nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, mockService := setupMockClient()
			started := make(chan struct{})
			cleanup := test.configure(mockService, started)
			defer cleanup()
			cancelled := make(chan struct {
				target [2]uint16
				err    error
			}, 1)
			mockService.cancelContextHook = func(ctx context.Context, messageID uint16, presentationContextID byte) {
				cancelled <- struct {
					target [2]uint16
					err    error
				}{[2]uint16{messageID, uint16(presentationContextID)}, ctx.Err()}
			}

			ctx, cancel := context.WithCancel(context.Background())
			result := make(chan error, 1)
			go func() { result <- test.run(ctx, client) }()
			<-started
			cancel()

			select {
			case got := <-cancelled:
				if got.target != test.want {
					t.Fatalf("C-CANCEL target = %v, want %v", got.target, test.want)
				}
				if got.err != nil {
					t.Fatalf("C-CANCEL context error = %v, want independent live context", got.err)
				}
			case <-time.After(250 * time.Millisecond):
				t.Fatal("context cancellation did not send C-CANCEL")
			}
			if err := <-result; !errors.Is(err, context.Canceled) {
				t.Fatalf("operation error = %v, want context.Canceled", err)
			}
		})
	}
}

func TestCFind_Success_MultipleResults(t *testing.T) {
	client, mockService := setupMockClient()

	// Setup mock responses

	// Create mock identifiers
	id1 := dataset.New()
	_ = id1.Add(element.NewString(tag.PatientID, vr.LO, []string{testPatientID1}))
	_ = id1.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))

	id2 := dataset.New()
	_ = id2.Add(element.NewString(tag.PatientID, vr.LO, []string{testPatientID2}))
	_ = id2.Add(element.NewString(tag.PatientName, vr.PN, []string{"SMITH^JANE"}))

	// Create mock request for response generation
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelPatient, dataset.New())
	_ = req.SetMessageID(1)

	mockService.findResponses = []*dimse.CFindResponse{
		dimse.NewCFindResponseFromRequest(req, status.CFindPending, id1), // Pending
		dimse.NewCFindResponseFromRequest(req, status.CFindPending, id2), // Pending
		dimse.NewCFindResponseFromRequest(req, status.Success, nil),      // Success
	}

	query := dataset.New()
	_ = query.Add(element.NewString(tag.PatientName, vr.PN, []string{"*"}))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results, err := client.CFind(ctx, dimse.QueryRetrieveLevelPatient, query)
	if err != nil {
		t.Errorf("CFind failed: %v", err)
	}

	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}

	// Verify results
	if len(results) >= 1 {
		patientID, ok := results[0].GetString(tag.PatientID)
		if !ok || patientID != testPatientID1 {
			t.Errorf("Expected PatientID 'P001', got '%s'", patientID)
		}
	}

	if len(results) >= 2 {
		patientID, ok := results[1].GetString(tag.PatientID)
		if !ok || patientID != testPatientID2 {
			t.Errorf("Expected PatientID 'P002', got '%s'", patientID)
		}
	}
}

func TestCFindAcceptsPendingWarning(t *testing.T) {
	client, mockService := setupMockClient()
	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.PatientID, vr.LO, []string{testPatientID1}))
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelPatient, dataset.New())
	mockService.findResponses = []*dimse.CFindResponse{
		dimse.NewCFindResponseFromRequest(req, status.CFindPendingWarningOptionalKeysNotSupported, identifier),
		dimse.NewCFindResponseFromRequest(req, status.Success, nil),
	}

	results, err := client.CFind(context.Background(), dimse.QueryRetrieveLevelPatient, dataset.New())
	if err != nil {
		t.Fatalf("CFind() error = %v", err)
	}
	if len(results) != 1 || results[0] != identifier {
		t.Fatalf("CFind() results = %#v, want pending-warning identifier", results)
	}
}

func TestCFindWithCallbackAcceptsPendingWarning(t *testing.T) {
	client, mockService := setupMockClient()
	identifier := dataset.New()
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelPatient, dataset.New())
	mockService.findResponses = []*dimse.CFindResponse{
		dimse.NewCFindResponseFromRequest(req, status.CFindPendingWarningOptionalKeysNotSupported, identifier),
		dimse.NewCFindResponseFromRequest(req, status.Success, nil),
	}

	called := 0
	err := client.CFindWithCallback(context.Background(), dimse.QueryRetrieveLevelPatient, dataset.New(), func(got *dataset.Dataset) bool {
		called++
		return got == identifier
	})
	if err != nil {
		t.Fatalf("CFindWithCallback() error = %v", err)
	}
	if called != 1 {
		t.Fatalf("callback count = %d, want 1", called)
	}
}

func TestCFindWithCallback_Success(t *testing.T) {
	client, mockService := setupMockClient()

	// Setup mock responses

	id1 := dataset.New()
	_ = id1.Add(element.NewString(tag.PatientID, vr.LO, []string{testPatientID1}))

	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelPatient, dataset.New())
	_ = req.SetMessageID(1)

	mockService.findResponses = []*dimse.CFindResponse{
		dimse.NewCFindResponseFromRequest(req, status.CFindPending, id1), // Pending
		dimse.NewCFindResponseFromRequest(req, status.Success, nil),      // Success
	}

	query := dataset.New()
	resultCount := 0

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := client.CFindWithCallback(ctx, dimse.QueryRetrieveLevelPatient, query,
		func(_ *dataset.Dataset) bool {
			resultCount++
			return true // Continue
		})

	if err != nil {
		t.Errorf("CFindWithCallback failed: %v", err)
	}

	if resultCount != 1 {
		t.Errorf("Expected 1 result callback, got %d", resultCount)
	}
}

func TestCFindWithCallback_StopEarly(t *testing.T) {
	client, mockService := setupMockClient()
	responses := make(chan *dimse.CFindResponse, 2)
	requests := make(chan *dimse.CFindRequest, 1)
	cancels := make(chan [2]uint16, 1)
	mockService.findResponseCh = responses
	mockService.findHook = func(req *dimse.CFindRequest) { requests <- req }
	mockService.cancelHook = func(messageID uint16, presentationContextID byte) {
		cancels <- [2]uint16{messageID, uint16(presentationContextID)}
	}

	result := make(chan error, 1)
	go func() {
		result <- client.CFindWithCallback(
			context.Background(), dimse.QueryRetrieveLevelPatient, dataset.New(),
			func(*dataset.Dataset) bool { return false },
		)
	}()
	req := <-requests
	responses <- dimse.NewCFindResponseFromRequest(req, status.CFindPending, dataset.New())
	if got := receiveCancelTarget(t, cancels); got != [2]uint16{101, 3} {
		t.Fatalf("C-CANCEL target = %v, want message/context [101 3]", got)
	}
	select {
	case err := <-result:
		t.Fatalf("CFindWithCallback() returned before final Cancel response: %v", err)
	case <-time.After(25 * time.Millisecond):
	}
	responses <- dimse.NewCFindResponseFromRequest(req, status.Cancel, nil)
	close(responses)
	if err := <-result; err != nil {
		t.Fatalf("CFindWithCallback() error = %v", err)
	}
	if err := client.CEcho(context.Background()); err != nil {
		t.Fatalf("CEcho() after cancelled C-FIND error = %v", err)
	}
}

func TestCFindWithCallbackReturnsAutomaticCancelSendFailure(t *testing.T) {
	client, mockService := setupMockClient()
	responses := make(chan *dimse.CFindResponse, 1)
	mockService.findResponseCh = responses
	mockService.findHook = func(req *dimse.CFindRequest) {
		responses <- dimse.NewCFindResponseFromRequest(req, status.CFindPending, dataset.New())
	}
	wantErr := errors.New("cancel write failed")
	mockService.cancelErr = wantErr

	err := client.CFindWithCallback(
		context.Background(), dimse.QueryRetrieveLevelStudy, dataset.New(),
		func(*dataset.Dataset) bool { return false },
	)
	if !errors.Is(err, wantErr) {
		t.Fatalf("CFindWithCallback() error = %v, want cancel send failure", err)
	}
}

func TestCFindPatientRoot_UsesPatientRootModel(t *testing.T) {
	client, mockService := setupMockClient()

	sopClassCh := make(chan string, 1)
	mockService.findHook = func(req *dimse.CFindRequest) {
		sopClassCh <- req.AffectedSOPClassUID()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.CFindPatientRoot(ctx, dimse.QueryRetrieveLevelStudy, dataset.New())
	if err != nil {
		t.Fatalf("CFindPatientRoot failed: %v", err)
	}

	select {
	case sopClassUID := <-sopClassCh:
		const patientRootFind = "1.2.840.10008.5.1.4.1.2.1.1"
		if sopClassUID != patientRootFind {
			t.Fatalf("AffectedSOPClassUID = %q, want %q", sopClassUID, patientRootFind)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for captured C-FIND request")
	}
}

func TestCFindStudyRoot_UsesStudyRootModel(t *testing.T) {
	client, mockService := setupMockClient()

	sopClassCh := make(chan string, 1)
	mockService.findHook = func(req *dimse.CFindRequest) {
		sopClassCh <- req.AffectedSOPClassUID()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.CFindStudyRoot(ctx, dimse.QueryRetrieveLevelPatient, dataset.New())
	if err != nil {
		t.Fatalf("CFindStudyRoot failed: %v", err)
	}

	select {
	case sopClassUID := <-sopClassCh:
		const studyRootFind = "1.2.840.10008.5.1.4.1.2.2.1"
		if sopClassUID != studyRootFind {
			t.Fatalf("AffectedSOPClassUID = %q, want %q", sopClassUID, studyRootFind)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for captured C-FIND request")
	}
}

func TestCStoreWithPriority_Success(t *testing.T) {
	client, _ := setupMockClient()

	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := client.CStoreWithPriority(ctx, ds, dimse.PriorityHigh)
	if err != nil {
		t.Errorf("CStoreWithPriority failed: %v", err)
	}
}

func TestCStoreMultiple_Success(t *testing.T) {
	client, _ := setupMockClient()

	// Create multiple datasets
	datasets := make([]*dataset.Dataset, 3)
	for i := 0; i < 3; i++ {
		ds := dataset.New()
		_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
		instanceUID := fmt.Sprintf("1.2.3.4.5.6.7.8.%d", i)
		_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{instanceUID}))
		datasets[i] = ds
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	count, err := client.CStoreMultiple(ctx, datasets)
	if err != nil {
		t.Errorf("CStoreMultiple failed: %v", err)
	}

	if count != 3 {
		t.Errorf("Expected 3 datasets stored, got %d", count)
	}
}

func TestCStoreMultipleStopsAtFirstFailure(t *testing.T) {
	client, mockService := setupMockClient()
	datasets := make([]*dataset.Dataset, 3)
	for i := range datasets {
		ds := dataset.New()
		_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
		_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{fmt.Sprintf("1.2.3.%d", i)}))
		datasets[i] = ds
	}

	calls := 0
	mockService.storeHook = func(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
		calls++
		if calls == 2 {
			return nil, errors.New("store failed")
		}
		return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
	}

	count, err := client.CStoreMultiple(context.Background(), datasets)
	if err == nil {
		t.Fatal("CStoreMultiple() error = nil, want second-item failure")
	}
	if count != 1 {
		t.Fatalf("CStoreMultiple() count = %d, want 1 completed dataset", count)
	}
	if calls != 2 {
		t.Fatalf("SendCStore() calls = %d, want 2", calls)
	}
}

func TestCStoreMultipleStopsAfterCancellation(t *testing.T) {
	client, mockService := setupMockClient()
	datasets := make([]*dataset.Dataset, 3)
	for i := range datasets {
		ds := dataset.New()
		_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
		_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{fmt.Sprintf("1.2.4.%d", i)}))
		datasets[i] = ds
	}

	ctx, cancel := context.WithCancel(context.Background())
	calls := 0
	mockService.storeHook = func(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
		calls++
		cancel()
		return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
	}

	count, err := client.CStoreMultiple(ctx, datasets)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("CStoreMultiple() error = %v, want context.Canceled", err)
	}
	if count != 1 {
		t.Fatalf("CStoreMultiple() count = %d, want 1 completed dataset", count)
	}
	if calls != 1 {
		t.Fatalf("SendCStore() calls = %d, want 1", calls)
	}
}

func TestCStoreMultiple_Empty(t *testing.T) {
	client, _ := setupMockClient()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	count, err := client.CStoreMultiple(ctx, []*dataset.Dataset{})
	if err != nil {
		t.Errorf("CStoreMultiple with empty slice should not fail: %v", err)
	}

	if count != 0 {
		t.Errorf("Expected 0 datasets stored, got %d", count)
	}
}

func TestCStoreMultiple_NotConnected(t *testing.T) {
	client := New()

	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))

	ctx := context.Background()
	_, err := client.CStoreMultiple(ctx, []*dataset.Dataset{ds})
	if err == nil {
		t.Error("Expected error when not connected")
	}

	expectedMsg := errNotConnected
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestPing_Success(t *testing.T) {
	client, _ := setupMockClient()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := client.Ping(ctx)
	if err != nil {
		t.Errorf("Ping failed: %v", err)
	}
}

func TestCMove_Success(t *testing.T) {
	client, _ := setupMockClient()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	progressCount := 0
	err := client.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST_AE", identifier,
		func(remaining, completed, failed, warning uint16) bool {
			progressCount++
			t.Logf("C-MOVE progress: remaining=%d, completed=%d, failed=%d, warning=%d",
				remaining, completed, failed, warning)
			return true // Continue receiving updates
		})

	if err != nil {
		t.Errorf("CMove failed: %v", err)
	}

	if progressCount != 3 {
		t.Errorf("Expected 3 progress callbacks, got %d", progressCount)
	}
}

func TestCMovePropagatesAsynchronousRequestTimeout(t *testing.T) {
	client, mockService := setupMockClient()
	mockService.moveTerminalErr = service.ErrRequestTimeout

	err := client.CMove(
		context.Background(), dimse.QueryRetrieveLevelStudy, "DEST_AE", dataset.New(), nil,
	)
	if !errors.Is(err, service.ErrRequestTimeout) {
		t.Fatalf("CMove() error = %v, want ErrRequestTimeout", err)
	}
}

func TestCMove_NotConnected(t *testing.T) {
	client := New()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))

	ctx := context.Background()
	err := client.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST_AE", identifier, nil)
	if err == nil {
		t.Error("Expected error when not connected")
	}

	expectedMsg := errNotConnected
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCMove_NilIdentifier(t *testing.T) {
	client, _ := setupMockClient()

	ctx := context.Background()
	err := client.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST_AE", nil, nil)
	if err == nil {
		t.Error("Expected error when identifier is nil")
	}

	expectedMsg := "identifier is nil"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCMove_EmptyDestination(t *testing.T) {
	client, _ := setupMockClient()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))

	ctx := context.Background()
	err := client.CMove(ctx, dimse.QueryRetrieveLevelStudy, "", identifier, nil)
	if err == nil {
		t.Error("Expected error when destination is empty")
	}

	expectedMsg := "move destination is empty"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCMove_StopEarly(t *testing.T) {
	client, mockService := setupMockClient()
	responses := make(chan *dimse.CMoveResponse, 2)
	requests := make(chan *dimse.CMoveRequest, 1)
	cancels := make(chan [2]uint16, 1)
	mockService.moveResponseCh = responses
	mockService.moveHook = func(req *dimse.CMoveRequest) { requests <- req }
	mockService.cancelHook = func(messageID uint16, presentationContextID byte) {
		cancels <- [2]uint16{messageID, uint16(presentationContextID)}
	}

	result := make(chan error, 1)
	go func() {
		result <- client.CMove(
			context.Background(), dimse.QueryRetrieveLevelStudy, "DEST_AE", dataset.New(),
			func(uint16, uint16, uint16, uint16) bool { return false },
		)
	}()
	req := <-requests
	responses <- dimse.NewCMoveResponsePending(req.MessageID(), req.AffectedSOPClassUID(), 1, 0, 0, 0)
	if got := receiveCancelTarget(t, cancels); got != [2]uint16{102, 5} {
		t.Fatalf("C-CANCEL target = %v, want message/context [102 5]", got)
	}
	select {
	case err := <-result:
		t.Fatalf("CMove() returned before final Cancel response: %v", err)
	case <-time.After(25 * time.Millisecond):
	}
	responses <- dimse.NewCMoveResponseFromRequest(req, status.Cancel)
	close(responses)
	if err := <-result; err != nil {
		t.Fatalf("CMove() error = %v", err)
	}
}

func TestCMoveStopAcceptsSuccessAlreadyInFlight(t *testing.T) {
	client, mockService := setupMockClient()
	responses := make(chan *dimse.CMoveResponse, 2)
	mockService.moveResponseCh = responses
	mockService.moveHook = func(req *dimse.CMoveRequest) {
		responses <- dimse.NewCMoveResponsePending(req.MessageID(), req.AffectedSOPClassUID(), 1, 0, 0, 0)
		responses <- dimse.NewCMoveResponseSuccess(req.MessageID(), req.AffectedSOPClassUID())
		close(responses)
	}

	err := client.CMove(
		context.Background(), dimse.QueryRetrieveLevelStudy, "DEST_AE", dataset.New(),
		func(uint16, uint16, uint16, uint16) bool { return false },
	)
	if err != nil {
		t.Fatalf("CMove() error = %v, want nil after final Success", err)
	}
}

func TestCFindStopAcceptsSuccessAlreadyInFlight(t *testing.T) {
	client, mockService := setupMockClient()
	responses := make(chan *dimse.CFindResponse, 2)
	mockService.findResponseCh = responses
	mockService.findHook = func(req *dimse.CFindRequest) {
		responses <- dimse.NewCFindResponseFromRequest(req, status.CFindPending, dataset.New())
		responses <- dimse.NewCFindResponseFromRequest(req, status.Success, nil)
		close(responses)
	}

	err := client.CFindWithCallback(context.Background(), dimse.QueryRetrieveLevelStudy, dataset.New(), func(*dataset.Dataset) bool {
		return false
	})
	if err != nil {
		t.Fatalf("CFindWithCallback() error = %v, want nil after final Success", err)
	}
}

func TestCGetStopAcceptsWarningAlreadyInFlight(t *testing.T) {
	client, mockService := setupMockClient()
	responses := make(chan *dimse.CGetResponse, 2)
	mockService.getResponseCh = responses
	mockService.getHook = func(req *dimse.CGetRequest) {
		responses <- dimse.NewCGetResponsePending(req.MessageID(), req.AffectedSOPClassUID(), 1, 0, 0, 0)
		responses <- dimse.NewCGetResponseFromRequest(req, status.CGetWarningSubOperationsComplete)
		close(responses)
	}

	err := client.CGet(context.Background(), dimse.QueryRetrieveLevelStudy, dataset.New(), func(uint16, uint16, uint16, uint16) bool {
		return false
	})
	if err != nil {
		t.Fatalf("CGet() error = %v, want nil after final Warning", err)
	}
}

func TestCMoveBufferedFinalResponseWinsOverCanceledContext(t *testing.T) {
	client, mockService := setupMockClient()
	responses := make(chan *dimse.CMoveResponse, 1)
	mockService.moveResponseCh = responses
	cancelSent := make(chan struct{}, 1)
	mockService.cancelHook = func(uint16, byte) { cancelSent <- struct{}{} }
	ctx, cancel := context.WithCancel(context.Background())
	mockService.moveHook = func(req *dimse.CMoveRequest) {
		responses <- dimse.NewCMoveResponseSuccess(req.MessageID(), req.AffectedSOPClassUID())
		close(responses)
		cancel()
	}

	if err := client.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST_AE", dataset.New(), nil); err != nil {
		t.Fatalf("CMove() error = %v, want buffered Success", err)
	}
	select {
	case <-cancelSent:
		t.Fatal("CMove() sent C-CANCEL after receiving a final Success")
	default:
	}
}

func TestCGet_Success(t *testing.T) {
	client, _ := setupMockClient()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	progressCount := 0
	err := client.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier,
		func(remaining, completed, failed, warning uint16) bool {
			progressCount++
			t.Logf("C-GET progress: remaining=%d, completed=%d, failed=%d, warning=%d",
				remaining, completed, failed, warning)
			return true // Continue receiving updates
		})

	if err != nil {
		t.Errorf("CGet failed: %v", err)
	}

	if progressCount != 3 {
		t.Errorf("Expected 3 progress callbacks, got %d", progressCount)
	}
}

func TestCGetPropagatesAsynchronousAssociationClose(t *testing.T) {
	client, mockService := setupMockClient()
	mockService.getTerminalErr = service.ErrServiceClosed

	err := client.CGet(context.Background(), dimse.QueryRetrieveLevelStudy, dataset.New(), nil)
	if !errors.Is(err, service.ErrServiceClosed) {
		t.Fatalf("CGet() error = %v, want ErrServiceClosed", err)
	}
}

func TestCGet_NotConnected(t *testing.T) {
	client := New()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))

	ctx := context.Background()
	err := client.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier, nil)
	if err == nil {
		t.Error("Expected error when not connected")
	}

	expectedMsg := errNotConnected
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCGet_NilIdentifier(t *testing.T) {
	client, _ := setupMockClient()

	ctx := context.Background()
	err := client.CGet(ctx, dimse.QueryRetrieveLevelStudy, nil, nil)
	if err == nil {
		t.Error("Expected error when identifier is nil")
	}

	expectedMsg := "identifier is nil"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestCGet_StopEarly(t *testing.T) {
	client, mockService := setupMockClient()
	responses := make(chan *dimse.CGetResponse, 2)
	requests := make(chan *dimse.CGetRequest, 1)
	cancels := make(chan [2]uint16, 1)
	mockService.getResponseCh = responses
	mockService.getHook = func(req *dimse.CGetRequest) { requests <- req }
	mockService.cancelHook = func(messageID uint16, presentationContextID byte) {
		cancels <- [2]uint16{messageID, uint16(presentationContextID)}
	}

	result := make(chan error, 1)
	go func() {
		result <- client.CGet(
			context.Background(), dimse.QueryRetrieveLevelStudy, dataset.New(),
			func(uint16, uint16, uint16, uint16) bool { return false },
		)
	}()
	req := <-requests
	responses <- dimse.NewCGetResponsePending(req.MessageID(), req.AffectedSOPClassUID(), 1, 0, 0, 0)
	if got := receiveCancelTarget(t, cancels); got != [2]uint16{103, 7} {
		t.Fatalf("C-CANCEL target = %v, want message/context [103 7]", got)
	}
	select {
	case err := <-result:
		t.Fatalf("CGet() returned before final Cancel response: %v", err)
	case <-time.After(25 * time.Millisecond):
	}
	responses <- dimse.NewCGetResponseFromRequest(req, status.Cancel)
	close(responses)
	if err := <-result; err != nil {
		t.Fatalf("CGet() error = %v", err)
	}
}

func receiveCancelTarget(t *testing.T, cancels <-chan [2]uint16) [2]uint16 {
	t.Helper()
	select {
	case target := <-cancels:
		return target
	case <-time.After(250 * time.Millisecond):
		t.Fatal("operation did not send C-CANCEL")
		return [2]uint16{}
	}
}
