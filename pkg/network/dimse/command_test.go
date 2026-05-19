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
		{cStoreRQCommandName, CommandCStoreRQ, cStoreRQCommandName},
		{cStoreRSPCommandName, CommandCStoreRSP, cStoreRSPCommandName},
		{cFindRQCommandName, CommandCFindRQ, cFindRQCommandName},
		{cFindRSPCommandName, CommandCFindRSP, cFindRSPCommandName},
		{cGetRQCommandName, CommandCGetRQ, cGetRQCommandName},
		{cGetRSPCommandName, CommandCGetRSP, cGetRSPCommandName},
		{cMoveRQCommandName, CommandCMoveRQ, cMoveRQCommandName},
		{cMoveRSPCommandName, CommandCMoveRSP, cMoveRSPCommandName},
		{cEchoRQCommandName, CommandCEchoRQ, cEchoRQCommandName},
		{cEchoRSPCommandName, CommandCEchoRSP, cEchoRSPCommandName},
		{cCancelRQCommandName, CommandCCancelRQ, cCancelRQCommandName},
		{nEventReportRQCommandName, CommandNEventReportRQ, nEventReportRQCommandName},
		{nEventReportRSPCommandName, CommandNEventReportRSP, nEventReportRSPCommandName},
		{nGetRQCommandName, CommandNGetRQ, nGetRQCommandName},
		{nGetRSPCommandName, CommandNGetRSP, nGetRSPCommandName},
		{nSetRQCommandName, CommandNSetRQ, nSetRQCommandName},
		{nSetRSPCommandName, CommandNSetRSP, nSetRSPCommandName},
		{nActionRQCommandName, CommandNActionRQ, nActionRQCommandName},
		{nActionRSPCommandName, CommandNActionRSP, nActionRSPCommandName},
		{nCreateRQCommandName, CommandNCreateRQ, nCreateRQCommandName},
		{nCreateRSPCommandName, CommandNCreateRSP, nCreateRSPCommandName},
		{nDeleteRQCommandName, CommandNDeleteRQ, nDeleteRQCommandName},
		{nDeleteRSPCommandName, CommandNDeleteRSP, nDeleteRSPCommandName},
		{unknownCommandName, CommandField(0x9999), unknownCommandName},
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
		{cStoreRQCommandName, CommandCStoreRQ, true},
		{cStoreRSPCommandName, CommandCStoreRSP, false},
		{cEchoRQCommandName, CommandCEchoRQ, true},
		{cEchoRSPCommandName, CommandCEchoRSP, false},
		{nGetRQCommandName, CommandNGetRQ, true},
		{nGetRSPCommandName, CommandNGetRSP, false},
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
		{cStoreRQCommandName, CommandCStoreRQ, false},
		{cStoreRSPCommandName, CommandCStoreRSP, true},
		{cEchoRQCommandName, CommandCEchoRQ, false},
		{cEchoRSPCommandName, CommandCEchoRSP, true},
		{nGetRQCommandName, CommandNGetRQ, false},
		{nGetRSPCommandName, CommandNGetRSP, true},
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
		{cStoreRQCommandName, CommandCStoreRQ, 0x0001},
		{cStoreRSPCommandName, CommandCStoreRSP, 0x8001},
		{cGetRQCommandName, CommandCGetRQ, 0x0010},
		{cGetRSPCommandName, CommandCGetRSP, 0x8010},
		{cFindRQCommandName, CommandCFindRQ, 0x0020},
		{cFindRSPCommandName, CommandCFindRSP, 0x8020},
		{cMoveRQCommandName, CommandCMoveRQ, 0x0021},
		{cMoveRSPCommandName, CommandCMoveRSP, 0x8021},
		{cEchoRQCommandName, CommandCEchoRQ, 0x0030},
		{cEchoRSPCommandName, CommandCEchoRSP, 0x8030},
		{cCancelRQCommandName, CommandCCancelRQ, 0x0FFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if uint16(tt.command) != tt.expected {
				t.Errorf("Expected value 0x%04X for %s, got 0x%04X", tt.expected, tt.name, uint16(tt.command))
			}
		})
	}
}
