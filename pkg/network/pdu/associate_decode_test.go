// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package pdu

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
	"testing"
)

// mustEncodeAssociation encodes an association PDU built by mutate.
func mustEncodeAssociation(t *testing.T, mutate func() (*RawPDU, error)) *RawPDU {
	t.Helper()
	raw, err := mutate()
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	return raw
}

func encodeRQForResetTest(t *testing.T, contextCount int, withIdentity bool) *RawPDU {
	t.Helper()
	return mustEncodeAssociation(t, func() (*RawPDU, error) {
		rq := NewAAssociateRQ()
		rq.CalledAETitle = testCalledAETitle
		rq.CallingAETitle = testCallingAETitle
		rq.UserInformation.MaximumLength = 1234
		rq.UserInformation.ImplementationClassUID = "1.2.3.4"
		rq.UserInformation.ImplementationVersionName = "RQ_VERSION"
		for i := 0; i < contextCount; i++ {
			rq.PresentationContexts = append(rq.PresentationContexts, PresentationContextRQ{
				ID:               byte(2*i + 1),
				AbstractSyntax:   testCTImageStorageUID,
				TransferSyntaxes: []string{testExplicitVRLittleLE},
			})
		}
		if withIdentity {
			rq.UserInformation.SCPSCURoleSelections = []SCPSCURoleSelection{{
				SOPClassUID: testCTImageStorageUID,
				SCURole:     1,
				SCPRole:     0,
			}}
			rq.UserInformation.UserIdentity = &UserIdentityNegotiation{
				UserIdentityType:          1,
				PositiveResponseRequested: 1,
				PrimaryField:              []byte("alice"),
			}
		}
		return rq.Encode()
	})
}

func encodeACForResetTest(t *testing.T, contextCount int, version string) *RawPDU {
	t.Helper()
	return mustEncodeAssociation(t, func() (*RawPDU, error) {
		ac := NewAAssociateAC()
		ac.CalledAETitle = testCalledAETitle
		ac.CallingAETitle = testCallingAETitle
		ac.UserInformation.MaximumLength = 4321
		ac.UserInformation.ImplementationClassUID = "5.6.7"
		ac.UserInformation.ImplementationVersionName = version
		for i := 0; i < contextCount; i++ {
			ac.PresentationContexts = append(ac.PresentationContexts, PresentationContextAC{
				ID:             byte(2*i + 1),
				Result:         ResultAcceptance,
				TransferSyntax: testExplicitVRLittleLE,
			})
		}
		return ac.Encode()
	})
}

// TestAAssociateRQDecodeResetsReceiverState verifies that decoding into a
// receiver that already holds a previous PDU yields the same state as decoding
// into a fresh receiver (no field or slice bleed).
func TestAAssociateRQDecodeResetsReceiverState(t *testing.T) {
	first := encodeRQForResetTest(t, 2, true)
	second := encodeRQForResetTest(t, 1, false)

	reused := &AAssociateRQ{}
	if err := reused.Decode(first); err != nil {
		t.Fatalf("Decode(first) error = %v", err)
	}
	if err := reused.Decode(second); err != nil {
		t.Fatalf("Decode(second) error = %v", err)
	}

	fresh := &AAssociateRQ{}
	if err := fresh.Decode(second); err != nil {
		t.Fatalf("Decode(second) into fresh receiver error = %v", err)
	}
	if !reflect.DeepEqual(reused, fresh) {
		t.Fatalf("reused receiver kept state from the first decode:\nreused = %#v\nfresh  = %#v", reused, fresh)
	}
	if len(reused.PresentationContexts) != 1 {
		t.Fatalf("presentation contexts = %d, want 1", len(reused.PresentationContexts))
	}
	if reused.UserInformation.UserIdentity != nil || len(reused.UserInformation.SCPSCURoleSelections) != 0 {
		t.Fatalf("user information kept state from the first decode: %#v", reused.UserInformation)
	}
}

// TestAAssociateACDecodeResetsReceiverState is the A-ASSOCIATE-AC counterpart.
func TestAAssociateACDecodeResetsReceiverState(t *testing.T) {
	first := encodeACForResetTest(t, 2, "FIRST")
	second := encodeACForResetTest(t, 1, "SECOND")

	reused := &AAssociateAC{}
	if err := reused.Decode(first); err != nil {
		t.Fatalf("Decode(first) error = %v", err)
	}
	if err := reused.Decode(second); err != nil {
		t.Fatalf("Decode(second) error = %v", err)
	}

	fresh := &AAssociateAC{}
	if err := fresh.Decode(second); err != nil {
		t.Fatalf("Decode(second) into fresh receiver error = %v", err)
	}
	if !reflect.DeepEqual(reused, fresh) {
		t.Fatalf("reused receiver kept state from the first decode:\nreused = %#v\nfresh  = %#v", reused, fresh)
	}
	if len(reused.PresentationContexts) != 1 {
		t.Fatalf("presentation contexts = %d, want 1", len(reused.PresentationContexts))
	}
}

// padApplicationContext rebuilds raw with trailing space/NUL padding on the
// Application Context UID value, as a peer may legally emit (PS3.5 UID padding).
func padApplicationContext(t *testing.T, raw *RawPDU) *RawPDU {
	t.Helper()
	reader := bytes.NewReader(raw.Data[68:])
	var out bytes.Buffer
	out.Write(raw.Data[:68])
	for {
		itemType, data, err := readItem(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("readItem() error = %v", err)
		}
		if itemType == ItemTypeApplicationContext {
			data = append(append([]byte(nil), data...), ' ', 0)
		}
		if err := writeItem(&out, itemType, data); err != nil {
			t.Fatalf("writeItem() error = %v", err)
		}
	}
	return NewRawPDU(raw.Type, out.Bytes())
}

// TestAssociateDecodeTrimsApplicationContextPadding verifies that a padded
// Application Context UID is trimmed before the mandatory-value comparison.
func TestAssociateDecodeTrimsApplicationContextPadding(t *testing.T) {
	t.Run("A-ASSOCIATE-RQ", func(t *testing.T) {
		raw := padApplicationContext(t, encodeRQForResetTest(t, 1, false))
		decoded := &AAssociateRQ{}
		if err := decoded.Decode(raw); err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
		if decoded.ApplicationContext != applicationContextUID {
			t.Fatalf("ApplicationContext = %q, want %q", decoded.ApplicationContext, applicationContextUID)
		}
		if err := decoded.ValidateStructure(); err != nil {
			t.Fatalf("ValidateStructure() error = %v", err)
		}
	})

	t.Run("A-ASSOCIATE-AC", func(t *testing.T) {
		raw := padApplicationContext(t, encodeACForResetTest(t, 1, "VERSION"))
		decoded := &AAssociateAC{}
		if err := decoded.Decode(raw); err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
		if decoded.ApplicationContext != applicationContextUID {
			t.Fatalf("ApplicationContext = %q, want %q", decoded.ApplicationContext, applicationContextUID)
		}
		if err := decoded.ValidateStructure(); err != nil {
			t.Fatalf("ValidateStructure() error = %v", err)
		}
	})
}

// encodeUserIdentityTestValue encodes a User Identity Sub-Item body.
func encodeUserIdentityTestValue(identityType, positiveResponse byte, primary, secondary []byte) []byte {
	data := []byte{identityType, positiveResponse}
	data = binary.BigEndian.AppendUint16(data, uint16(len(primary))) // #nosec G115 -- test data
	data = append(data, primary...)
	data = binary.BigEndian.AppendUint16(data, uint16(len(secondary))) // #nosec G115 -- test data
	return append(data, secondary...)
}

// TestAAssociateRQEncodeRejectsInvalidUserIdentity verifies the encoder applies
// the same Table D.3-14 shape rules as the decoder, so a locally constructed
// identity that a conformant peer would reject never reaches the wire.
func TestAAssociateRQEncodeRejectsInvalidUserIdentity(t *testing.T) {
	newRequest := func() *AAssociateRQ {
		rq := NewAAssociateRQ()
		rq.CalledAETitle = testServerAETitle
		rq.CallingAETitle = testClientAETitle
		rq.PresentationContexts = testProposedContexts()
		return rq
	}

	invalidType := newRequest()
	invalidType.UserInformation.UserIdentity = &UserIdentityNegotiation{UserIdentityType: 0}
	if _, err := invalidType.Encode(); err == nil {
		t.Fatal("Encode() accepted a reserved user identity type")
	}

	unexpectedSecondary := newRequest()
	unexpectedSecondary.UserInformation.UserIdentity = &UserIdentityNegotiation{
		UserIdentityType: 1,
		PrimaryField:     []byte("alice"),
		SecondaryField:   []byte("secret"),
	}
	if _, err := unexpectedSecondary.Encode(); err == nil {
		t.Fatal("Encode() accepted a secondary field for a username identity")
	}

	emptyPasscode := newRequest()
	emptyPasscode.UserInformation.UserIdentity = &UserIdentityNegotiation{
		UserIdentityType: 2,
		PrimaryField:     []byte("alice"),
	}
	if _, err := emptyPasscode.Encode(); err != nil {
		t.Fatalf("Encode() rejected a username/password identity without a passcode: %v", err)
	}
}

// TestDecodeUserIdentityValidTypes verifies the identity types of PS3.7
// Table D.3-14 are accepted.
func TestDecodeUserIdentityValidTypes(t *testing.T) {
	tests := []struct {
		name      string
		identity  byte
		primary   []byte
		secondary []byte
	}{
		{"username", 1, []byte("alice"), nil},
		{"username and password", 2, []byte("alice"), []byte("secret")},
		// The standard only says the secondary field shall be non-zero only
		// for type 2; an empty passcode is not forbidden (Table D.3-14).
		{"username and password without passcode", 2, []byte("alice"), nil},
		{"kerberos", 3, []byte("ticket"), nil},
		{"SAML", 4, []byte("assertion"), nil},
		{"JWT", 5, []byte("token"), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeUserIdentity(encodeUserIdentityTestValue(tt.identity, 1, tt.primary, tt.secondary))
			if err != nil {
				t.Fatalf("decodeUserIdentity() error = %v", err)
			}
			if got.UserIdentityType != tt.identity ||
				!bytes.Equal(got.PrimaryField, tt.primary) ||
				!bytes.Equal(got.SecondaryField, tt.secondary) {
				t.Fatalf("decoded identity = %#v", got)
			}
		})
	}
}

// TestDecodeUserIdentityRejectsInvalidValues verifies the PS3.7 Table D.3-14
// constraints: defined type and positive-response values, the prohibition of a
// secondary field for types other than username/password, and no trailing
// bytes.
func TestDecodeUserIdentityRejectsInvalidValues(t *testing.T) {
	valid := encodeUserIdentityTestValue(1, 1, []byte("alice"), nil)
	tests := []struct {
		name string
		data []byte
	}{
		{"reserved identity type 0", encodeUserIdentityTestValue(0, 1, []byte("alice"), nil)},
		{"reserved identity type 6", encodeUserIdentityTestValue(6, 1, []byte("alice"), nil)},
		{"reserved positive response value", encodeUserIdentityTestValue(1, 2, []byte("alice"), nil)},
		{"username with secondary", encodeUserIdentityTestValue(1, 1, []byte("alice"), []byte("secret"))},
		{"JWT with secondary", encodeUserIdentityTestValue(5, 1, []byte("token"), []byte("secret"))},
		{"trailing bytes", append(append([]byte(nil), valid...), 0x00)},
		{"truncated", valid[:len(valid)-1]},
		{"too short", []byte{1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := decodeUserIdentity(tt.data); err == nil {
				t.Fatal("decodeUserIdentity() accepted an invalid user identity")
			}
		})
	}
}

// TestAAssociateRQDecodeRejectsDuplicateUserIdentity verifies PS3.7 D.3.3.4:
// at most one User Identity Sub-Item is present.
func TestAAssociateRQDecodeRejectsDuplicateUserIdentity(t *testing.T) {
	data := append(
		encodeTestItem(ItemTypeUserIdentityNegotiation, encodeUserIdentityTestValue(1, 1, []byte("alice"), nil)),
		encodeTestItem(ItemTypeUserIdentityNegotiation, encodeUserIdentityTestValue(1, 1, []byte("bob"), nil))...,
	)

	rq := &AAssociateRQ{}
	if _, err := rq.decodeUserInformation(data, nil); err == nil {
		t.Fatal("decodeUserInformation() accepted duplicate user identity sub-items")
	}
}
