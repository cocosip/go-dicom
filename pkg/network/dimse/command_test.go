// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import "testing"

func TestCommandField_String(t *testing.T) {
	tests := []struct {
		name     string
		command  CommandField
		expected string
	}{
		{"C-STORE-RQ", CommandCStoreRQ, "C-STORE-RQ"},
		{"C-STORE-RSP", CommandCStoreRSP, "C-STORE-RSP"},
		{"C-FIND-RQ", CommandCFindRQ, "C-FIND-RQ"},
		{"C-FIND-RSP", CommandCFindRSP, "C-FIND-RSP"},
		{"C-GET-RQ", CommandCGetRQ, "C-GET-RQ"},
		{"C-GET-RSP", CommandCGetRSP, "C-GET-RSP"},
		{"C-MOVE-RQ", CommandCMoveRQ, "C-MOVE-RQ"},
		{"C-MOVE-RSP", CommandCMoveRSP, "C-MOVE-RSP"},
		{"C-ECHO-RQ", CommandCEchoRQ, "C-ECHO-RQ"},
		{"C-ECHO-RSP", CommandCEchoRSP, "C-ECHO-RSP"},
		{"C-CANCEL-RQ", CommandCCancelRQ, "C-CANCEL-RQ"},
		{"N-EVENT-REPORT-RQ", CommandNEventReportRQ, "N-EVENT-REPORT-RQ"},
		{"N-EVENT-REPORT-RSP", CommandNEventReportRSP, "N-EVENT-REPORT-RSP"},
		{"N-GET-RQ", CommandNGetRQ, "N-GET-RQ"},
		{"N-GET-RSP", CommandNGetRSP, "N-GET-RSP"},
		{"N-SET-RQ", CommandNSetRQ, "N-SET-RQ"},
		{"N-SET-RSP", CommandNSetRSP, "N-SET-RSP"},
		{"N-ACTION-RQ", CommandNActionRQ, "N-ACTION-RQ"},
		{"N-ACTION-RSP", CommandNActionRSP, "N-ACTION-RSP"},
		{"N-CREATE-RQ", CommandNCreateRQ, "N-CREATE-RQ"},
		{"N-CREATE-RSP", CommandNCreateRSP, "N-CREATE-RSP"},
		{"N-DELETE-RQ", CommandNDeleteRQ, "N-DELETE-RQ"},
		{"N-DELETE-RSP", CommandNDeleteRSP, "N-DELETE-RSP"},
		{"Unknown", CommandField(0x9999), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.command.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestCommandField_IsRequest(t *testing.T) {
	tests := []struct {
		name     string
		command  CommandField
		expected bool
	}{
		{"C-STORE-RQ", CommandCStoreRQ, true},
		{"C-STORE-RSP", CommandCStoreRSP, false},
		{"C-ECHO-RQ", CommandCEchoRQ, true},
		{"C-ECHO-RSP", CommandCEchoRSP, false},
		{"N-GET-RQ", CommandNGetRQ, true},
		{"N-GET-RSP", CommandNGetRSP, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.command.IsRequest()
			if result != tt.expected {
				t.Errorf("Expected IsRequest() = %v for %s", tt.expected, tt.name)
			}
		})
	}
}

func TestCommandField_IsResponse(t *testing.T) {
	tests := []struct {
		name     string
		command  CommandField
		expected bool
	}{
		{"C-STORE-RQ", CommandCStoreRQ, false},
		{"C-STORE-RSP", CommandCStoreRSP, true},
		{"C-ECHO-RQ", CommandCEchoRQ, false},
		{"C-ECHO-RSP", CommandCEchoRSP, true},
		{"N-GET-RQ", CommandNGetRQ, false},
		{"N-GET-RSP", CommandNGetRSP, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.command.IsResponse()
			if result != tt.expected {
				t.Errorf("Expected IsResponse() = %v for %s", tt.expected, tt.name)
			}
		})
	}
}

func TestCommandField_Values(t *testing.T) {
	// Test correct values
	tests := []struct {
		name     string
		command  CommandField
		expected uint16
	}{
		{"C-STORE-RQ", CommandCStoreRQ, 0x0001},
		{"C-STORE-RSP", CommandCStoreRSP, 0x8001},
		{"C-GET-RQ", CommandCGetRQ, 0x0010},
		{"C-GET-RSP", CommandCGetRSP, 0x8010},
		{"C-FIND-RQ", CommandCFindRQ, 0x0020},
		{"C-FIND-RSP", CommandCFindRSP, 0x8020},
		{"C-MOVE-RQ", CommandCMoveRQ, 0x0021},
		{"C-MOVE-RSP", CommandCMoveRSP, 0x8021},
		{"C-ECHO-RQ", CommandCEchoRQ, 0x0030},
		{"C-ECHO-RSP", CommandCEchoRSP, 0x8030},
		{"C-CANCEL-RQ", CommandCCancelRQ, 0x0FFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if uint16(tt.command) != tt.expected {
				t.Errorf("Expected value 0x%04X for %s, got 0x%04X", tt.expected, tt.name, uint16(tt.command))
			}
		})
	}
}
