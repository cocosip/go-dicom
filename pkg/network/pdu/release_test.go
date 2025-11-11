package pdu

import (
	"testing"
)

const (
	testReleaseRQStr = "A-RELEASE-RQ"
	testReleaseRPStr = "A-RELEASE-RP"
)

func TestNewAReleaseRQ(t *testing.T) {
	rq := NewAReleaseRQ()
	if rq == nil {
		t.Error("Expected non-nil A-RELEASE-RQ")
	}
}

func TestAReleaseRQ_EncodeDecode(t *testing.T) {
	// Create
	original := NewAReleaseRQ()

	// Encode
	pdu, err := original.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	if pdu.Type != TypeAReleaseRQ {
		t.Errorf("Expected PDU type 0x05, got 0x%02X", pdu.Type)
	}

	if pdu.Length != 4 {
		t.Errorf("Expected PDU length 4, got %d", pdu.Length)
	}

	// Decode
	decoded := &AReleaseRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
}

func TestAReleaseRQ_DecodeInvalidPDUType(t *testing.T) {
	pdu := NewRawPDU(TypeAReleaseRP, make([]byte, 4))

	rq := &AReleaseRQ{}
	err := rq.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid PDU type, got nil")
	}
}

func TestAReleaseRQ_DecodeInvalidLength(t *testing.T) {
	pdu := NewRawPDU(TypeAReleaseRQ, []byte{0x00, 0x00})

	rq := &AReleaseRQ{}
	err := rq.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid data length, got nil")
	}
}

func TestAReleaseRQ_String(t *testing.T) {
	rq := NewAReleaseRQ()
	if rq.String() != testReleaseRQStr {
		t.Errorf("Expected 'A-RELEASE-RQ', got %s", rq.String())
	}
}

func TestNewAReleaseRP(t *testing.T) {
	rp := NewAReleaseRP()
	if rp == nil {
		t.Error("Expected non-nil A-RELEASE-RP")
	}
}

func TestAReleaseRP_EncodeDecode(t *testing.T) {
	// Create
	original := NewAReleaseRP()

	// Encode
	pdu, err := original.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	if pdu.Type != TypeAReleaseRP {
		t.Errorf("Expected PDU type 0x06, got 0x%02X", pdu.Type)
	}

	if pdu.Length != 4 {
		t.Errorf("Expected PDU length 4, got %d", pdu.Length)
	}

	// Decode
	decoded := &AReleaseRP{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
}

func TestAReleaseRP_DecodeInvalidPDUType(t *testing.T) {
	pdu := NewRawPDU(TypeAReleaseRQ, make([]byte, 4))

	rp := &AReleaseRP{}
	err := rp.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid PDU type, got nil")
	}
}

func TestAReleaseRP_DecodeInvalidLength(t *testing.T) {
	pdu := NewRawPDU(TypeAReleaseRP, []byte{0x00, 0x00})

	rp := &AReleaseRP{}
	err := rp.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid data length, got nil")
	}
}

func TestAReleaseRP_String(t *testing.T) {
	rp := NewAReleaseRP()
	if rp.String() != testReleaseRPStr {
		t.Errorf("Expected 'A-RELEASE-RP', got %s", rp.String())
	}
}

func TestReleaseRoundtrip(t *testing.T) {
	// Simulate a complete release handshake

	// 1. Create and encode request
	rq := NewAReleaseRQ()
	rqPDU, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encoding request failed: %v", err)
	}

	// 2. Decode request
	decodedRQ := &AReleaseRQ{}
	if err := decodedRQ.Decode(rqPDU); err != nil {
		t.Fatalf("Decoding request failed: %v", err)
	}

	// 3. Create and encode response
	rp := NewAReleaseRP()
	rpPDU, err := rp.Encode()
	if err != nil {
		t.Fatalf("Encoding response failed: %v", err)
	}

	// 4. Decode response
	decodedRP := &AReleaseRP{}
	if err := decodedRP.Decode(rpPDU); err != nil {
		t.Fatalf("Decoding response failed: %v", err)
	}

	// Verify PDU types
	if rqPDU.Type != TypeAReleaseRQ {
		t.Errorf("Request PDU type mismatch: expected 0x05, got 0x%02X", rqPDU.Type)
	}

	if rpPDU.Type != TypeAReleaseRP {
		t.Errorf("Response PDU type mismatch: expected 0x06, got 0x%02X", rpPDU.Type)
	}
}
