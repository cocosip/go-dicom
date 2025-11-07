// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

// CommandField represents DIMSE command field values.
type CommandField uint16

// DIMSE-C command fields
const (
	// C-STORE
	CommandCStoreRQ CommandField = 0x0001
	CommandCStoreRSP CommandField = 0x8001

	// C-GET
	CommandCGetRQ CommandField = 0x0010
	CommandCGetRSP CommandField = 0x8010

	// C-FIND
	CommandCFindRQ CommandField = 0x0020
	CommandCFindRSP CommandField = 0x8020

	// C-MOVE
	CommandCMoveRQ CommandField = 0x0021
	CommandCMoveRSP CommandField = 0x8021

	// C-ECHO
	CommandCEchoRQ CommandField = 0x0030
	CommandCEchoRSP CommandField = 0x8030

	// C-CANCEL
	CommandCCancelRQ CommandField = 0x0FFF
)

// DIMSE-N command fields
const (
	// N-EVENT-REPORT
	CommandNEventReportRQ CommandField = 0x0100
	CommandNEventReportRSP CommandField = 0x8100

	// N-GET
	CommandNGetRQ CommandField = 0x0110
	CommandNGetRSP CommandField = 0x8110

	// N-SET
	CommandNSetRQ CommandField = 0x0120
	CommandNSetRSP CommandField = 0x8120

	// N-ACTION
	CommandNActionRQ CommandField = 0x0130
	CommandNActionRSP CommandField = 0x8130

	// N-CREATE
	CommandNCreateRQ CommandField = 0x0140
	CommandNCreateRSP CommandField = 0x8140

	// N-DELETE
	CommandNDeleteRQ CommandField = 0x0150
	CommandNDeleteRSP CommandField = 0x8150
)

// String returns the string representation of the command field.
func (c CommandField) String() string {
	switch c {
	case CommandCStoreRQ:
		return "C-STORE-RQ"
	case CommandCStoreRSP:
		return "C-STORE-RSP"
	case CommandCGetRQ:
		return "C-GET-RQ"
	case CommandCGetRSP:
		return "C-GET-RSP"
	case CommandCFindRQ:
		return "C-FIND-RQ"
	case CommandCFindRSP:
		return "C-FIND-RSP"
	case CommandCMoveRQ:
		return "C-MOVE-RQ"
	case CommandCMoveRSP:
		return "C-MOVE-RSP"
	case CommandCEchoRQ:
		return "C-ECHO-RQ"
	case CommandCEchoRSP:
		return "C-ECHO-RSP"
	case CommandCCancelRQ:
		return "C-CANCEL-RQ"
	case CommandNEventReportRQ:
		return "N-EVENT-REPORT-RQ"
	case CommandNEventReportRSP:
		return "N-EVENT-REPORT-RSP"
	case CommandNGetRQ:
		return "N-GET-RQ"
	case CommandNGetRSP:
		return "N-GET-RSP"
	case CommandNSetRQ:
		return "N-SET-RQ"
	case CommandNSetRSP:
		return "N-SET-RSP"
	case CommandNActionRQ:
		return "N-ACTION-RQ"
	case CommandNActionRSP:
		return "N-ACTION-RSP"
	case CommandNCreateRQ:
		return "N-CREATE-RQ"
	case CommandNCreateRSP:
		return "N-CREATE-RSP"
	case CommandNDeleteRQ:
		return "N-DELETE-RQ"
	case CommandNDeleteRSP:
		return "N-DELETE-RSP"
	default:
		return "Unknown"
	}
}

// IsRequest returns true if this command is a request.
func (c CommandField) IsRequest() bool {
	return (c & 0x8000) == 0
}

// IsResponse returns true if this command is a response.
func (c CommandField) IsResponse() bool {
	return (c & 0x8000) != 0
}
