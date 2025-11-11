// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import "testing"

// reasonStringer is an interface for PDU types that can convert reason codes to strings.
type reasonStringer interface {
	ReasonString() string
}

// testReasonStrings is a helper function to test reason string conversions for PDU types.
// It takes a list of test cases and a factory function that creates the PDU instance.
func testReasonStrings(t *testing.T, tests []struct {
	name     string
	source   byte
	reason   byte
	expected string
}, factory func(source, reason byte) reasonStringer) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pdu := factory(tt.source, tt.reason)
			result := pdu.ReasonString()
			if result != tt.expected {
				t.Errorf("ReasonString: expected %s, got %s", tt.expected, result)
			}
		})
	}
}
