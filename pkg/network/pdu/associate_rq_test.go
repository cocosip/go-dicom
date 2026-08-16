// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

const (
	testCTImageStorageUID          = "1.2.840.10008.5.1.4.1.1.2"
	testExtendedSOPClassUID        = "1.2.3"
	testExtendedServiceClassUID    = "4.5"
	testRelatedGeneralSOPClassUID1 = "6.7"
	testRelatedGeneralSOPClassUID2 = "8.9"
	// testServerAETitle and testClientAETitle are defined in associate_ac_test.go
)

func TestNewAAssociateRQ(t *testing.T) {
	rq := NewAAssociateRQ()

	if rq.ProtocolVersion != 0x0001 {
		t.Errorf("Expected protocol version 0x0001, got 0x%04X", rq.ProtocolVersion)
	}

	if rq.ApplicationContext != applicationContextUID {
		t.Errorf("Expected default application context, got %s", rq.ApplicationContext)
	}

	if rq.UserInformation == nil {
		t.Error("Expected UserInformation to be initialized")
	}

	if rq.UserInformation.MaximumLength != 16384 {
		t.Errorf("Expected default MaximumLength 16384, got %d", rq.UserInformation.MaximumLength)
	}
}

func TestAAssociateRQ_EncodeDecodeBasic(t *testing.T) {
	// Create a basic A-ASSOCIATE-RQ
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "CALLED_AE"
	rq.CallingAETitle = "CALLING_AE"

	// Add a presentation context
	rq.PresentationContexts = []PresentationContextRQ{
		{
			ID:             1,
			AbstractSyntax: testCTImageStorageUID, // CT Image Storage
			TransferSyntaxes: []string{
				testExplicitVRLittleLE, // Explicit VR Little Endian
				testImplicitVRLittleLE, // Implicit VR Little Endian
			},
		},
	}

	// Encode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	if pdu.Type != TypeAAssociateRQ {
		t.Errorf("Expected PDU type 0x01, got 0x%02X", pdu.Type)
	}

	// Decode
	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Compare
	if decoded.ProtocolVersion != rq.ProtocolVersion {
		t.Errorf("ProtocolVersion mismatch: expected 0x%04X, got 0x%04X", rq.ProtocolVersion, decoded.ProtocolVersion)
	}

	if decoded.CalledAETitle != rq.CalledAETitle {
		t.Errorf("CalledAETitle mismatch: expected %s, got %s", rq.CalledAETitle, decoded.CalledAETitle)
	}

	if decoded.CallingAETitle != rq.CallingAETitle {
		t.Errorf("CallingAETitle mismatch: expected %s, got %s", rq.CallingAETitle, decoded.CallingAETitle)
	}

	if decoded.ApplicationContext != rq.ApplicationContext {
		t.Errorf("ApplicationContext mismatch: expected %s, got %s", rq.ApplicationContext, decoded.ApplicationContext)
	}

	if len(decoded.PresentationContexts) != len(rq.PresentationContexts) {
		t.Fatalf("PresentationContexts count mismatch: expected %d, got %d", len(rq.PresentationContexts), len(decoded.PresentationContexts))
	}

	pc := decoded.PresentationContexts[0]
	if pc.ID != 1 {
		t.Errorf("PresentationContext ID mismatch: expected 1, got %d", pc.ID)
	}

	if pc.AbstractSyntax != testCTImageStorageUID {
		t.Errorf("AbstractSyntax mismatch: expected 1.2.840.10008.5.1.4.1.1.2, got %s", pc.AbstractSyntax)
	}

	if len(pc.TransferSyntaxes) != 2 {
		t.Fatalf("TransferSyntaxes count mismatch: expected 2, got %d", len(pc.TransferSyntaxes))
	}

	if decoded.UserInformation == nil {
		t.Fatal("UserInformation is nil")
	}

	if decoded.UserInformation.MaximumLength != rq.UserInformation.MaximumLength {
		t.Errorf("MaximumLength mismatch: expected %d, got %d", rq.UserInformation.MaximumLength, decoded.UserInformation.MaximumLength)
	}
}

func TestAAssociateRQ_EncodeDecodeMultiplePresentationContexts(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "SCP_AE"
	rq.CallingAETitle = "SCU_AE"

	// Add multiple presentation contexts
	rq.PresentationContexts = []PresentationContextRQ{
		{
			ID:             1,
			AbstractSyntax: testCTImageStorageUID,
			TransferSyntaxes: []string{
				testExplicitVRLittleLE,
			},
		},
		{
			ID:             3,
			AbstractSyntax: "1.2.840.10008.5.1.4.1.1.4",
			TransferSyntaxes: []string{
				testImplicitVRLittleLE,
				testExplicitVRLittleLE,
			},
		},
		{
			ID:             5,
			AbstractSyntax: "1.2.840.10008.1.1", // Verification SOP Class
			TransferSyntaxes: []string{
				testImplicitVRLittleLE,
			},
		},
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if len(decoded.PresentationContexts) != 3 {
		t.Fatalf("Expected 3 presentation contexts, got %d", len(decoded.PresentationContexts))
	}

	// Check each context
	for i, expected := range rq.PresentationContexts {
		got := decoded.PresentationContexts[i]
		if got.ID != expected.ID {
			t.Errorf("Context %d: ID mismatch: expected %d, got %d", i, expected.ID, got.ID)
		}
		if got.AbstractSyntax != expected.AbstractSyntax {
			t.Errorf("Context %d: AbstractSyntax mismatch", i)
		}
		if len(got.TransferSyntaxes) != len(expected.TransferSyntaxes) {
			t.Errorf("Context %d: TransferSyntaxes count mismatch", i)
		}
	}
}

func TestAAssociateRQ_EncodeDecodeWithUserIdentity(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = testServerAETitle
	rq.CallingAETitle = testClientAETitle

	// Add user identity (username only)
	rq.UserInformation.UserIdentity = &UserIdentityNegotiation{
		UserIdentityType:          1, // Username
		PositiveResponseRequested: 1,
		PrimaryField:              []byte("testuser"),
		SecondaryField:            nil,
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify user identity
	if decoded.UserInformation.UserIdentity == nil {
		t.Fatal("UserIdentity is nil after decode")
	}

	ui := decoded.UserInformation.UserIdentity
	if ui.UserIdentityType != 1 {
		t.Errorf("UserIdentityType mismatch: expected 1, got %d", ui.UserIdentityType)
	}

	if ui.PositiveResponseRequested != 1 {
		t.Errorf("PositiveResponseRequested mismatch: expected 1, got %d", ui.PositiveResponseRequested)
	}

	if !bytes.Equal(ui.PrimaryField, []byte("testuser")) {
		t.Errorf("PrimaryField mismatch: expected 'testuser', got %s", string(ui.PrimaryField))
	}
}

func TestAAssociateRQ_EncodeDecodeWithUsernamePassword(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = testServerAETitle
	rq.CallingAETitle = testClientAETitle

	// Add user identity (username + password)
	rq.UserInformation.UserIdentity = &UserIdentityNegotiation{
		UserIdentityType:          2, // Username + Password
		PositiveResponseRequested: 0,
		PrimaryField:              []byte("admin"),
		SecondaryField:            []byte("password123"),
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	ui := decoded.UserInformation.UserIdentity
	if ui == nil {
		t.Fatal("UserIdentity is nil")
	}

	if ui.UserIdentityType != 2 {
		t.Errorf("UserIdentityType mismatch")
	}

	if !bytes.Equal(ui.PrimaryField, []byte("admin")) {
		t.Errorf("PrimaryField mismatch")
	}

	if !bytes.Equal(ui.SecondaryField, []byte("password123")) {
		t.Errorf("SecondaryField mismatch")
	}
}

func TestAAssociateRQ_EncodeDecodeWithExtendedNegotiation(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = testServerAETitle
	rq.CallingAETitle = testClientAETitle

	// Add extended negotiation
	rq.UserInformation.ExtendedNegotiations = []ExtendedNegotiation{
		{
			SOPClassUID:         testCTImageStorageUID,
			ServiceClassAppInfo: []byte{0x01, 0x02, 0x03},
		},
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if len(decoded.UserInformation.ExtendedNegotiations) != 1 {
		t.Fatalf("Expected 1 extended negotiation, got %d", len(decoded.UserInformation.ExtendedNegotiations))
	}

	ext := decoded.UserInformation.ExtendedNegotiations[0]
	if ext.SOPClassUID != testCTImageStorageUID {
		t.Errorf("SOPClassUID mismatch")
	}

	if !bytes.Equal(ext.ServiceClassAppInfo, []byte{0x01, 0x02, 0x03}) {
		t.Errorf("ServiceClassAppInfo mismatch")
	}
}

func TestCommonExtendedNegotiation_EncodeExactBytes(t *testing.T) {
	negotiation := CommonExtendedNegotiation{
		SOPClassUID:                testExtendedSOPClassUID,
		ServiceClassUID:            testExtendedServiceClassUID,
		RelatedGeneralSOPClassUIDs: []string{testRelatedGeneralSOPClassUID1, testRelatedGeneralSOPClassUID2},
	}

	var encoded bytes.Buffer
	if err := encodeCommonExtendedNegotiation(&encoded, negotiation); err != nil {
		t.Fatalf("encodeCommonExtendedNegotiation() error = %v", err)
	}

	want := []byte{
		0x57, 0x00, 0x00, 0x18,
		0x00, 0x05, '1', '.', '2', '.', '3',
		0x00, 0x03, '4', '.', '5',
		0x00, 0x0a,
		0x00, 0x03, '6', '.', '7',
		0x00, 0x03, '8', '.', '9',
	}
	if !bytes.Equal(encoded.Bytes(), want) {
		t.Fatalf("encoded common extended negotiation = % x, want % x", encoded.Bytes(), want)
	}
}

func TestCommonExtendedNegotiation_AllowsEmptyRelatedUIDList(t *testing.T) {
	negotiation := CommonExtendedNegotiation{
		SOPClassUID:     testExtendedSOPClassUID,
		ServiceClassUID: testExtendedServiceClassUID,
	}

	var encoded bytes.Buffer
	if err := encodeCommonExtendedNegotiation(&encoded, negotiation); err != nil {
		t.Fatalf("encodeCommonExtendedNegotiation() error = %v", err)
	}
	want := []byte{
		0x57, 0x00, 0x00, 0x0e,
		0x00, 0x05, '1', '.', '2', '.', '3',
		0x00, 0x03, '4', '.', '5',
		0x00, 0x00,
	}
	if !bytes.Equal(encoded.Bytes(), want) {
		t.Fatalf("encoded common extended negotiation = % x, want % x", encoded.Bytes(), want)
	}

	decoded, err := decodeCommonExtendedNegotiation(encoded.Bytes()[4:])
	if err != nil {
		t.Fatalf("decodeCommonExtendedNegotiation() error = %v", err)
	}
	if decoded.RelatedGeneralSOPClassUIDs == nil || len(decoded.RelatedGeneralSOPClassUIDs) != 0 {
		t.Fatalf("related general SOP Class UIDs = %#v, want present empty list", decoded.RelatedGeneralSOPClassUIDs)
	}
}

func TestAAssociateRQ_EncodeDecodeWithCombinedExtendedNegotiation(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.UserInformation.ExtendedNegotiations = []ExtendedNegotiation{{
		SOPClassUID:         testExtendedSOPClassUID,
		ServiceClassAppInfo: []byte{1, 0, 1},
	}}
	rq.UserInformation.CommonExtendedNegotiations = []CommonExtendedNegotiation{{
		SOPClassUID:                testExtendedSOPClassUID,
		ServiceClassUID:            testExtendedServiceClassUID,
		RelatedGeneralSOPClassUIDs: []string{testRelatedGeneralSOPClassUID1, testRelatedGeneralSOPClassUID2},
	}}

	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if got := decoded.UserInformation.ExtendedNegotiations; len(got) != 1 ||
		got[0].SOPClassUID != testExtendedSOPClassUID || !bytes.Equal(got[0].ServiceClassAppInfo, []byte{1, 0, 1}) {
		t.Fatalf("extended negotiations = %#v", got)
	}
	if got := decoded.UserInformation.CommonExtendedNegotiations; len(got) != 1 ||
		got[0].SOPClassUID != testExtendedSOPClassUID || got[0].ServiceClassUID != testExtendedServiceClassUID ||
		len(got[0].RelatedGeneralSOPClassUIDs) != 2 ||
		got[0].RelatedGeneralSOPClassUIDs[0] != testRelatedGeneralSOPClassUID1 || got[0].RelatedGeneralSOPClassUIDs[1] != testRelatedGeneralSOPClassUID2 {
		t.Fatalf("common extended negotiations = %#v", got)
	}
}

func TestDecodeCommonExtendedNegotiationRejectsMalformedLengths(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{name: "truncated SOP Class UID", data: []byte{0x00, 0x03, '1', '.'}},
		{name: "truncated Service Class UID", data: []byte{0x00, 0x01, '1', 0x00, 0x02, '2'}},
		{name: "orphan related byte", data: []byte{0x00, 0x01, '1', 0x00, 0x01, '2', 0x00, 0x01, 0x00}},
		{name: "truncated related UID", data: []byte{0x00, 0x01, '1', 0x00, 0x01, '2', 0x00, 0x04, 0x00, 0x03, '3', '.'}},
		{name: "trailing item data", data: []byte{0x00, 0x01, '1', 0x00, 0x01, '2', 0x00, 0x00, 0xff}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := decodeCommonExtendedNegotiation(tt.data); err == nil {
				t.Fatal("decodeCommonExtendedNegotiation() error = nil")
			}
		})
	}
}

func TestCommonExtendedNegotiationEncodeRejectsOversizedRelatedBlock(t *testing.T) {
	negotiation := CommonExtendedNegotiation{
		SOPClassUID:                testExtendedSOPClassUID,
		ServiceClassUID:            testExtendedServiceClassUID,
		RelatedGeneralSOPClassUIDs: []string{strings.Repeat("1", 65534)},
	}

	if err := encodeCommonExtendedNegotiation(&bytes.Buffer{}, negotiation); err == nil {
		t.Fatal("encodeCommonExtendedNegotiation() error = nil")
	}
}

func TestCommonExtendedNegotiationRejectsEmptyRequiredUIDs(t *testing.T) {
	encodeTests := []CommonExtendedNegotiation{
		{SOPClassUID: "", ServiceClassUID: "2"},
		{SOPClassUID: "1", ServiceClassUID: ""},
		{SOPClassUID: "1", ServiceClassUID: "2", RelatedGeneralSOPClassUIDs: []string{""}},
	}
	for _, negotiation := range encodeTests {
		if err := encodeCommonExtendedNegotiation(&bytes.Buffer{}, negotiation); err == nil {
			t.Fatalf("encodeCommonExtendedNegotiation(%#v) error = nil", negotiation)
		}
	}

	decodeTests := [][]byte{
		{0x00, 0x00, 0x00, 0x01, '2', 0x00, 0x00},
		{0x00, 0x01, '1', 0x00, 0x00, 0x00, 0x00},
		{0x00, 0x01, '1', 0x00, 0x01, '2', 0x00, 0x02, 0x00, 0x00},
	}
	for _, data := range decodeTests {
		if _, err := decodeCommonExtendedNegotiation(data); err == nil {
			t.Fatalf("decodeCommonExtendedNegotiation(% x) error = nil", data)
		}
	}
}

func TestReadItemRejectsTruncatedHeader(t *testing.T) {
	for length := 1; length < 4; length++ {
		_, _, err := readItem(bytes.NewReader([]byte{ItemTypeCommonExtendedNegotiation, 0x00, 0x00}[:length]))
		if err == nil || err == io.EOF {
			t.Fatalf("readItem() error for %d-byte header = %v, want truncation error", length, err)
		}
	}
}

func TestAAssociateRQ_EncodeDecodeWithRoleSelection(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = testServerAETitle
	rq.CallingAETitle = testClientAETitle

	// Add role selection
	rq.UserInformation.SCPSCURoleSelections = []SCPSCURoleSelection{
		{
			SOPClassUID: testCTImageStorageUID,
			SCURole:     1, // Support SCU
			SCPRole:     0, // No SCP support
		},
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if len(decoded.UserInformation.SCPSCURoleSelections) != 1 {
		t.Fatalf("Expected 1 role selection, got %d", len(decoded.UserInformation.SCPSCURoleSelections))
	}

	role := decoded.UserInformation.SCPSCURoleSelections[0]
	if role.SOPClassUID != testCTImageStorageUID {
		t.Errorf("SOPClassUID mismatch")
	}

	if role.SCURole != 1 {
		t.Errorf("SCURole mismatch: expected 1, got %d", role.SCURole)
	}

	if role.SCPRole != 0 {
		t.Errorf("SCPRole mismatch: expected 0, got %d", role.SCPRole)
	}
}

func TestAAssociateRQ_EncodeDecodeWithImplementationVersion(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = testServerAETitle
	rq.CallingAETitle = testClientAETitle
	rq.UserInformation.ImplementationVersionName = "GO_DICOM_1_0"

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if decoded.UserInformation.ImplementationVersionName != "GO_DICOM_1_0" {
		t.Errorf("ImplementationVersionName mismatch: expected GO_DICOM_1_0, got %s", decoded.UserInformation.ImplementationVersionName)
	}
}

func TestAAssociateRQ_DecodeInvalidPDUType(t *testing.T) {
	pdu := NewRawPDU(TypeAAssociateAC, []byte{0x00})

	rq := &AAssociateRQ{}
	err := rq.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid PDU type, got nil")
	}
}

func TestAETitleSpacePadding(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "SHORT"    // Only 5 chars
	rq.CallingAETitle = "ANOTHER" // 7 chars

	// Encode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// Decode
	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify no extra spaces
	if decoded.CalledAETitle != "SHORT" {
		t.Errorf("CalledAETitle should trim spaces: expected 'SHORT', got '%s'", decoded.CalledAETitle)
	}

	if decoded.CallingAETitle != "ANOTHER" {
		t.Errorf("CallingAETitle should trim spaces: expected 'ANOTHER', got '%s'", decoded.CallingAETitle)
	}
}

func TestAAssociateRQDecodeIsSilentAndReportsStructuredWarnings(t *testing.T) {
	raw := unknownItemsAssociateRQ()

	output := captureStdout(t, func() {
		decoded := &AAssociateRQ{}
		if err := decoded.Decode(raw); err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
	})
	if output != "" {
		t.Fatalf("Decode() wrote to stdout: %q", output)
	}

	var warnings []DecodeWarning
	decoded := &AAssociateRQ{}
	if err := decoded.DecodeWithWarnings(raw, func(warning DecodeWarning) {
		warnings = append(warnings, warning)
	}); err != nil {
		t.Fatalf("DecodeWithWarnings() error = %v", err)
	}
	wantCodes := []DecodeWarningCode{
		DecodeWarningUnknownItem,
		DecodeWarningUnknownPresentationContextSubItem,
		DecodeWarningUnknownUserInformationSubItem,
	}
	if len(warnings) != len(wantCodes) {
		t.Fatalf("warning count = %d, want %d: %#v", len(warnings), len(wantCodes), warnings)
	}
	for i, want := range wantCodes {
		if warnings[i].Code != want || warnings[i].ItemType != 0x99 {
			t.Errorf("warning[%d] = %#v, want code %q and item type 0x99", i, warnings[i], want)
		}
	}
}

func unknownItemsAssociateRQ() *RawPDU {
	fixed := make([]byte, 68)
	fixed[1] = 1

	unknown := encodeTestItem(0x99, nil)
	presentation := append([]byte{1, 0, 0, 0}, encodeTestItem(0x99, nil)...)
	userInformation := encodeTestItem(0x99, nil)

	data := append(fixed, unknown...)
	data = append(data, encodeTestItem(ItemTypePresentationContextRQ, presentation)...)
	data = append(data, encodeTestItem(ItemTypeUserInformation, userInformation)...)
	return NewRawPDU(TypeAAssociateRQ, data)
}

func encodeTestItem(itemType byte, data []byte) []byte {
	item := []byte{itemType, 0, byte(len(data) >> 8), byte(len(data))}
	return append(item, data...)
}

func captureStdout(t *testing.T, run func()) string {
	t.Helper()
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() error = %v", err)
	}
	original := os.Stdout
	os.Stdout = writer
	defer func() { os.Stdout = original }()

	run()
	if err := writer.Close(); err != nil {
		t.Fatalf("stdout writer Close() error = %v", err)
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll(stdout) error = %v", err)
	}
	if err := reader.Close(); err != nil {
		t.Fatalf("stdout reader Close() error = %v", err)
	}
	return string(data)
}
