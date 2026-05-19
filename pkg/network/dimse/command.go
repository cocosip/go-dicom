// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

// CommandField represents DIMSE command field values.
type CommandField uint16

const (
	cStoreRQCommandName        = "C-STORE-RQ"
	cStoreRSPCommandName       = "C-STORE-RSP"
	cGetRQCommandName          = "C-GET-RQ"
	cGetRSPCommandName         = "C-GET-RSP"
	cFindRQCommandName         = "C-FIND-RQ"
	cFindRSPCommandName        = "C-FIND-RSP"
	cMoveRQCommandName         = "C-MOVE-RQ"
	cMoveRSPCommandName        = "C-MOVE-RSP"
	cEchoRQCommandName         = "C-ECHO-RQ"
	cEchoRSPCommandName        = "C-ECHO-RSP"
	cCancelRQCommandName       = "C-CANCEL-RQ"
	nEventReportRQCommandName  = "N-EVENT-REPORT-RQ"
	nEventReportRSPCommandName = "N-EVENT-REPORT-RSP"
	nGetRQCommandName          = "N-GET-RQ"
	nGetRSPCommandName         = "N-GET-RSP"
	nSetRQCommandName          = "N-SET-RQ"
	nSetRSPCommandName         = "N-SET-RSP"
	nActionRQCommandName       = "N-ACTION-RQ"
	nActionRSPCommandName      = "N-ACTION-RSP"
	nCreateRQCommandName       = "N-CREATE-RQ"
	nCreateRSPCommandName      = "N-CREATE-RSP"
	nDeleteRQCommandName       = "N-DELETE-RQ"
	nDeleteRSPCommandName      = "N-DELETE-RSP"
	unknownCommandName         = "Unknown"

	sopClassUIDVerification = "1.2.840.10008.1.1"
	sopClassUIDStudyRootGet = "1.2.840.10008.5.1.4.1.2.2.3"
)

// DIMSE-C command fields
const (
	// C-STORE
	CommandCStoreRQ  CommandField = 0x0001
	CommandCStoreRSP CommandField = 0x8001

	// C-GET
	CommandCGetRQ  CommandField = 0x0010
	CommandCGetRSP CommandField = 0x8010

	// C-FIND
	CommandCFindRQ  CommandField = 0x0020
	CommandCFindRSP CommandField = 0x8020

	// C-MOVE
	CommandCMoveRQ  CommandField = 0x0021
	CommandCMoveRSP CommandField = 0x8021

	// C-ECHO
	CommandCEchoRQ  CommandField = 0x0030
	CommandCEchoRSP CommandField = 0x8030

	// C-CANCEL
	CommandCCancelRQ CommandField = 0x0FFF
)

// DIMSE-N command fields
const (
	// N-EVENT-REPORT
	CommandNEventReportRQ  CommandField = 0x0100
	CommandNEventReportRSP CommandField = 0x8100

	// N-GET
	CommandNGetRQ  CommandField = 0x0110
	CommandNGetRSP CommandField = 0x8110

	// N-SET
	CommandNSetRQ  CommandField = 0x0120
	CommandNSetRSP CommandField = 0x8120

	// N-ACTION
	CommandNActionRQ  CommandField = 0x0130
	CommandNActionRSP CommandField = 0x8130

	// N-CREATE
	CommandNCreateRQ  CommandField = 0x0140
	CommandNCreateRSP CommandField = 0x8140

	// N-DELETE
	CommandNDeleteRQ  CommandField = 0x0150
	CommandNDeleteRSP CommandField = 0x8150
)

// String returns the string representation of the command field.
func (c CommandField) String() string {
	switch c {
	case CommandCStoreRQ:
		return cStoreRQCommandName
	case CommandCStoreRSP:
		return cStoreRSPCommandName
	case CommandCGetRQ:
		return cGetRQCommandName
	case CommandCGetRSP:
		return cGetRSPCommandName
	case CommandCFindRQ:
		return cFindRQCommandName
	case CommandCFindRSP:
		return cFindRSPCommandName
	case CommandCMoveRQ:
		return cMoveRQCommandName
	case CommandCMoveRSP:
		return cMoveRSPCommandName
	case CommandCEchoRQ:
		return cEchoRQCommandName
	case CommandCEchoRSP:
		return cEchoRSPCommandName
	case CommandCCancelRQ:
		return cCancelRQCommandName
	case CommandNEventReportRQ:
		return nEventReportRQCommandName
	case CommandNEventReportRSP:
		return nEventReportRSPCommandName
	case CommandNGetRQ:
		return nGetRQCommandName
	case CommandNGetRSP:
		return nGetRSPCommandName
	case CommandNSetRQ:
		return nSetRQCommandName
	case CommandNSetRSP:
		return nSetRSPCommandName
	case CommandNActionRQ:
		return nActionRQCommandName
	case CommandNActionRSP:
		return nActionRSPCommandName
	case CommandNCreateRQ:
		return nCreateRQCommandName
	case CommandNCreateRSP:
		return nCreateRSPCommandName
	case CommandNDeleteRQ:
		return nDeleteRQCommandName
	case CommandNDeleteRSP:
		return nDeleteRSPCommandName
	default:
		return unknownCommandName
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
