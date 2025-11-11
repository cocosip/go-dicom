// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
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
	client := New()
	client.connected = true // Fake connection

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

func TestCFind_NilQuery(t *testing.T) {
	client := New()
	client.connected = true // Fake connection

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
	client := New()
	client.connected = true // Fake connection
	client.assoc = association.NewAssociation("TEST_SCU", "TEST_SCP")

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
	echoResponse  *dimse.CEchoResponse
	storeResponse *dimse.CStoreResponse
	findResponses []*dimse.CFindResponse
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
func (m *mockServiceForDIMSE) Start() error                                  { return nil }
func (m *mockServiceForDIMSE) GracefulRelease(_ context.Context) error       { return nil }
func (m *mockServiceForDIMSE) Abort(_ context.Context, _ byte, _ byte) error { return nil }

func (m *mockServiceForDIMSE) SendCEcho(_ context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
	// Simulate processing delay
	time.Sleep(10 * time.Millisecond)

	if m.echoResponse != nil {
		return m.echoResponse, nil
	}

	// Default success response
	return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
}

func (m *mockServiceForDIMSE) SendCStore(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
	// Simulate processing delay
	time.Sleep(10 * time.Millisecond)

	if m.storeResponse != nil {
		return m.storeResponse, nil
	}

	// Default success response
	return dimse.NewCStoreResponseFromRequest(req, 0x0000), nil
}

func (m *mockServiceForDIMSE) SendCFind(_ context.Context, req *dimse.CFindRequest) (<-chan *dimse.CFindResponse, error) {
	// Create response channel
	respCh := make(chan *dimse.CFindResponse, 16)

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
			resp := dimse.NewCFindResponseFromRequest(req, 0x0000, nil)
			respCh <- resp
		}
	}()

	return respCh, nil
}

func (m *mockServiceForDIMSE) SendCMove(_ context.Context, req *dimse.CMoveRequest) (<-chan *dimse.CMoveResponse, error) {
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

func (m *mockServiceForDIMSE) SendCGet(_ context.Context, req *dimse.CGetRequest) (<-chan *dimse.CGetResponse, error) {
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

func setupMockClient() (*Client, *mockServiceForDIMSE) {
	client := New()
	client.connected = true
	client.assoc = association.NewAssociation("TEST_SCU", "TEST_SCP")
	mockService := &mockServiceForDIMSE{}
	client.service = mockService
	return client, mockService
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
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"})) // CT Image Storage
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))

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
	query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))

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

func TestCFind_Success_MultipleResults(t *testing.T) {
	client, mockService := setupMockClient()

	// Setup mock responses

	// Create mock identifiers
	id1 := dataset.New()
	id1.Add(element.NewString(tag.PatientID, vr.LO, []string{"P001"}))
	id1.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))

	id2 := dataset.New()
	id2.Add(element.NewString(tag.PatientID, vr.LO, []string{"P002"}))
	id2.Add(element.NewString(tag.PatientName, vr.PN, []string{"SMITH^JANE"}))

	// Create mock request for response generation
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelPatient, dataset.New())
	req.SetMessageID(1)

	mockService.findResponses = []*dimse.CFindResponse{
		dimse.NewCFindResponseFromRequest(req, 0xFF00, id1), // Pending
		dimse.NewCFindResponseFromRequest(req, 0xFF00, id2), // Pending
		dimse.NewCFindResponseFromRequest(req, 0x0000, nil), // Success
	}

	query := dataset.New()
	query.Add(element.NewString(tag.PatientName, vr.PN, []string{"*"}))

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
		if !ok || patientID != "P001" {
			t.Errorf("Expected PatientID 'P001', got '%s'", patientID)
		}
	}

	if len(results) >= 2 {
		patientID, ok := results[1].GetString(tag.PatientID)
		if !ok || patientID != "P002" {
			t.Errorf("Expected PatientID 'P002', got '%s'", patientID)
		}
	}
}

func TestCFindWithCallback_Success(t *testing.T) {
	client, mockService := setupMockClient()

	// Setup mock responses

	id1 := dataset.New()
	id1.Add(element.NewString(tag.PatientID, vr.LO, []string{"P001"}))

	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelPatient, dataset.New())
	req.SetMessageID(1)

	mockService.findResponses = []*dimse.CFindResponse{
		dimse.NewCFindResponseFromRequest(req, 0xFF00, id1), // Pending
		dimse.NewCFindResponseFromRequest(req, 0x0000, nil), // Success
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

	// Setup mock responses with multiple results

	id1 := dataset.New()
	id1.Add(element.NewString(tag.PatientID, vr.LO, []string{"P001"}))

	id2 := dataset.New()
	id2.Add(element.NewString(tag.PatientID, vr.LO, []string{"P002"}))

	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelPatient, dataset.New())
	req.SetMessageID(1)

	mockService.findResponses = []*dimse.CFindResponse{
		dimse.NewCFindResponseFromRequest(req, 0xFF00, id1), // Pending
		dimse.NewCFindResponseFromRequest(req, 0xFF00, id2), // Pending
		dimse.NewCFindResponseFromRequest(req, 0x0000, nil), // Success
	}

	query := dataset.New()
	resultCount := 0

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := client.CFindWithCallback(ctx, dimse.QueryRetrieveLevelPatient, query,
		func(_ *dataset.Dataset) bool {
			resultCount++
			return false // Stop after first result
		})

	if err != nil {
		t.Errorf("CFindWithCallback failed: %v", err)
	}

	if resultCount != 1 {
		t.Errorf("Expected 1 result callback (stopped early), got %d", resultCount)
	}
}

func TestCStoreWithPriority_Success(t *testing.T) {
	client, _ := setupMockClient()

	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))

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
		ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
		instanceUID := fmt.Sprintf("1.2.3.4.5.6.7.8.%d", i)
		ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{instanceUID}))
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
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))

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
	identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

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

func TestCMove_NotConnected(t *testing.T) {
	client := New()

	identifier := dataset.New()
	identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

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
	identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

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
	client, _ := setupMockClient()

	identifier := dataset.New()
	identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	progressCount := 0
	err := client.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST_AE", identifier,
		func(remaining, completed, failed, warning uint16) bool {
			progressCount++
			t.Logf("C-MOVE progress: remaining=%d, completed=%d, failed=%d, warning=%d",
				remaining, completed, failed, warning)
			// Stop after first progress update
			return false
		})

	if err != nil {
		t.Errorf("CMove failed: %v", err)
	}

	if progressCount != 1 {
		t.Errorf("Expected 1 progress callback (stopped early), got %d", progressCount)
	}
}

func TestCGet_Success(t *testing.T) {
	client, _ := setupMockClient()

	identifier := dataset.New()
	identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

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

func TestCGet_NotConnected(t *testing.T) {
	client := New()

	identifier := dataset.New()
	identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

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
	client, _ := setupMockClient()

	identifier := dataset.New()
	identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	progressCount := 0
	err := client.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier,
		func(remaining, completed, failed, warning uint16) bool {
			progressCount++
			t.Logf("C-GET progress: remaining=%d, completed=%d, failed=%d, warning=%d",
				remaining, completed, failed, warning)
			// Stop after first progress update
			return false
		})

	if err != nil {
		t.Errorf("CGet failed: %v", err)
	}

	if progressCount != 1 {
		t.Errorf("Expected 1 progress callback (stopped early), got %d", progressCount)
	}
}
