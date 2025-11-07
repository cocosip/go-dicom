// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package status implements DICOM status codes and classifications.
package status

import "fmt"

// Status represents a DICOM operation status code.
type Status struct {
	Code        uint16
	State       string
	Description string
}

// String returns a string representation of the status.
func (s *Status) String() string {
	return fmt.Sprintf("0x%04X (%s): %s", s.Code, s.State, s.Description)
}

// IsSuccess returns true if the status indicates success.
func (s *Status) IsSuccess() bool {
	return s.Code == 0x0000
}

// IsPending returns true if the status indicates pending (more results to come).
func (s *Status) IsPending() bool {
	return s.Code == 0xFF00 || s.Code == 0xFF01
}

// IsWarning returns true if the status is a warning (operation completed with issues).
// Warning range: 0x0001-0x00FF and 0xB000-0xBFFF
func (s *Status) IsWarning() bool {
	return (s.Code >= 0x0001 && s.Code <= 0x00FF) || (s.Code >= 0xB000 && s.Code <= 0xBFFF)
}

// IsFailure returns true if the status indicates failure.
// Failure range: 0x0100-0xAFFF and 0xC000-0xFDFF (excluding pending/cancel ranges)
func (s *Status) IsFailure() bool {
	return (s.Code >= 0x0100 && s.Code <= 0xAFFF) || (s.Code >= 0xC000 && s.Code <= 0xFDFF)
}

// IsCancel returns true if the status indicates cancellation.
func (s *Status) IsCancel() bool {
	return s.Code == 0xFE00
}

// NewStatus creates a new Status with the given code, state, and description.
func NewStatus(code uint16, state, description string) *Status {
	return &Status{
		Code:        code,
		State:       state,
		Description: description,
	}
}

// Common status classifications
const (
	StateSuccess = "Success"
	StatePending = "Pending"
	StateCancel  = "Cancel"
	StateWarning = "Warning"
	StateFailure = "Failure"
)

// Common DICOM status codes

// Success statuses
var (
	Success = NewStatus(0x0000, StateSuccess, "Success")
)

// Pending statuses
var (
	Pending                     = NewStatus(0xFF00, StatePending, "Pending")
	PendingWarning              = NewStatus(0xFF01, StatePending, "Pending with warning")
)

// Cancel statuses
var (
	Cancel = NewStatus(0xFE00, StateCancel, "Cancel")
)

// Warning statuses (0x0001-0x00FF)
var (
	AttributeListError          = NewStatus(0x0107, StateWarning, "Attribute list error")
	AttributeValueOutOfRange    = NewStatus(0x0116, StateWarning, "Attribute value out of range")
)

// Failure statuses (0x0100-0xFEFF)

// General failure statuses
var (
	RefusedOutOfResources              = NewStatus(0x0110, StateFailure, "Refused: Out of resources")
	RefusedMoveDestinationUnknown      = NewStatus(0x0112, StateFailure, "Refused: Move destination unknown")
	ErrorCannotUnderstand              = NewStatus(0x0122, StateFailure, "Cannot understand")
	ErrorDataSetDoesNotMatchSOPClass   = NewStatus(0x0117, StateFailure, "Dataset does not match SOP class")
	ErrorDuplicateInvocation           = NewStatus(0x0210, StateFailure, "Duplicate invocation")
	ErrorDuplicateSOPInstance          = NewStatus(0x0111, StateFailure, "Duplicate SOP instance")
	ErrorInvalidArgumentValue          = NewStatus(0x0115, StateFailure, "Invalid argument value")
	ErrorInvalidAttributeValue         = NewStatus(0x0106, StateFailure, "Invalid attribute value")
	ErrorInvalidObjectInstance         = NewStatus(0x0117, StateFailure, "Invalid object instance")
	ErrorMissingAttribute              = NewStatus(0x0120, StateFailure, "Missing attribute")
	ErrorMissingAttributeValue         = NewStatus(0x0121, StateFailure, "Missing attribute value")
	ErrorMistypedArgument              = NewStatus(0x0212, StateFailure, "Mistyped argument")
	ErrorNoSuchArgument                = NewStatus(0x0114, StateFailure, "No such argument")
	ErrorNoSuchAttribute               = NewStatus(0x0105, StateFailure, "No such attribute")
	ErrorNoSuchEventType               = NewStatus(0x0113, StateFailure, "No such event type")
	ErrorNoSuchObjectInstance          = NewStatus(0x0112, StateFailure, "No such object instance")
	ErrorNoSuchSOPClass                = NewStatus(0x0118, StateFailure, "No such SOP class")
	ErrorProcessingFailure             = NewStatus(0x0110, StateFailure, "Processing failure")
	ErrorResourceLimitation            = NewStatus(0x0213, StateFailure, "Resource limitation")
	ErrorUnrecognizedOperation         = NewStatus(0x0211, StateFailure, "Unrecognized operation")
	ErrorUnexpectedRequest             = NewStatus(0x0001, StateFailure, "Unexpected request")
)

// C-STORE specific statuses
var (
	CStoreRefusedOutOfResources         = NewStatus(0xA700, StateFailure, "Refused: Out of resources")
	CStoreRefusedDataSetDoesNotMatchSOPClass = NewStatus(0xA900, StateFailure, "Error: Data set does not match SOP class")
	CStoreErrorCannotUnderstand         = NewStatus(0xC000, StateFailure, "Error: Cannot understand")
	CStoreWarningCoercionOfDataElements = NewStatus(0xB000, StateWarning, "Coercion of data elements")
	CStoreWarningDataSetDoesNotMatchSOPClass = NewStatus(0xB007, StateWarning, "Data set does not match SOP class")
	CStoreWarningElementsDiscarded      = NewStatus(0xB006, StateWarning, "Elements discarded")
)

// C-FIND specific statuses
var (
	CFindRefusedOutOfResources          = NewStatus(0xA700, StateFailure, "Refused: Out of resources")
	CFindFailedIdentifierDoesNotMatchSOPClass = NewStatus(0xA900, StateFailure, "Identifier does not match SOP class")
	CFindFailedUnableToProcess          = NewStatus(0xC000, StateFailure, "Unable to process")
	CFindCancelMatchingTerminated       = NewStatus(0xFE00, StateCancel, "Matching terminated due to Cancel request")
	CFindPending                        = NewStatus(0xFF00, StatePending, "Matches are continuing")
	CFindPendingWarningOptionalKeysNotSupported = NewStatus(0xFF01, StatePending, "Matches are continuing - Warning: one or more optional keys were not supported")
)

// C-GET specific statuses
var (
	CGetRefusedOutOfResources           = NewStatus(0xA701, StateFailure, "Refused: Out of resources - Unable to calculate number of matches")
	CGetRefusedOutOfResourcesSubOps     = NewStatus(0xA702, StateFailure, "Refused: Out of resources - Unable to perform sub-operations")
	CGetRefusedMoveDestinationUnknown   = NewStatus(0xA801, StateFailure, "Refused: Move destination unknown")
	CGetFailedIdentifierDoesNotMatchSOPClass = NewStatus(0xA900, StateFailure, "Identifier does not match SOP class")
	CGetFailedUnableToProcess           = NewStatus(0xC000, StateFailure, "Unable to process")
	CGetCancelSubOperationsTerminated   = NewStatus(0xFE00, StateCancel, "Sub-operations terminated due to Cancel indication")
	CGetWarningSubOperationsComplete    = NewStatus(0xB000, StateWarning, "Sub-operations complete - One or more Failures or Warnings")
	CGetPending                         = NewStatus(0xFF00, StatePending, "Sub-operations are continuing")
)

// C-MOVE specific statuses
var (
	CMoveRefusedOutOfResources          = NewStatus(0xA701, StateFailure, "Refused: Out of resources - Unable to calculate number of matches")
	CMoveRefusedOutOfResourcesSubOps    = NewStatus(0xA702, StateFailure, "Refused: Out of resources - Unable to perform sub-operations")
	CMoveRefusedMoveDestinationUnknown  = NewStatus(0xA801, StateFailure, "Refused: Move destination unknown")
	CMoveFailedIdentifierDoesNotMatchSOPClass = NewStatus(0xA900, StateFailure, "Identifier does not match SOP class")
	CMoveFailedUnableToProcess          = NewStatus(0xC000, StateFailure, "Unable to process")
	CMoveCancelSubOperationsTerminated  = NewStatus(0xFE00, StateCancel, "Sub-operations terminated due to Cancel indication")
	CMoveWarningSubOperationsComplete   = NewStatus(0xB000, StateWarning, "Sub-operations complete - One or more Failures or Warnings")
	CMovePending                        = NewStatus(0xFF00, StatePending, "Sub-operations are continuing")
)

// C-ECHO specific statuses
var (
	CEchoSuccess = Success
)

// N-EVENT-REPORT specific statuses
var (
	NEventReportSuccess                 = Success
	NEventReportFailureProcessingFailure = NewStatus(0x0110, StateFailure, "Processing failure")
	NEventReportFailureNoSuchEventType  = NewStatus(0x0113, StateFailure, "No such event type")
	NEventReportFailureNoSuchObjectInstance = NewStatus(0x0112, StateFailure, "No such object instance")
	NEventReportFailureInvalidArgumentValue = NewStatus(0x0115, StateFailure, "Invalid argument value")
)

// N-GET specific statuses
var (
	NGetSuccess                         = Success
	NGetFailureProcessingFailure        = NewStatus(0x0110, StateFailure, "Processing failure")
	NGetFailureNoSuchObjectInstance     = NewStatus(0x0112, StateFailure, "No such object instance")
	NGetFailureNoSuchAttribute          = NewStatus(0x0105, StateFailure, "No such attribute")
	NGetFailureInvalidAttributeValue    = NewStatus(0x0106, StateFailure, "Invalid attribute value")
	NGetWarningAttributeListError       = NewStatus(0x0107, StateWarning, "Attribute list error")
)

// N-SET specific statuses
var (
	NSetSuccess                         = Success
	NSetFailureProcessingFailure        = NewStatus(0x0110, StateFailure, "Processing failure")
	NSetFailureNoSuchObjectInstance     = NewStatus(0x0112, StateFailure, "No such object instance")
	NSetFailureInvalidAttributeValue    = NewStatus(0x0106, StateFailure, "Invalid attribute value")
	NSetFailureMissingAttributeValue    = NewStatus(0x0120, StateFailure, "Missing attribute value")
	NSetWarningAttributeListError       = NewStatus(0x0107, StateWarning, "Attribute list error")
	NSetWarningAttributeValueOutOfRange = NewStatus(0x0116, StateWarning, "Attribute value out of range")
)

// N-ACTION specific statuses
var (
	NActionSuccess                      = Success
	NActionFailureProcessingFailure     = NewStatus(0x0110, StateFailure, "Processing failure")
	NActionFailureNoSuchObjectInstance  = NewStatus(0x0112, StateFailure, "No such object instance")
	NActionFailureNoSuchArgument        = NewStatus(0x0114, StateFailure, "No such argument")
	NActionFailureInvalidArgumentValue  = NewStatus(0x0115, StateFailure, "Invalid argument value")
	NActionFailureNoSuchActionType      = NewStatus(0x0123, StateFailure, "No such action type")
)

// N-CREATE specific statuses
var (
	NCreateSuccess                      = Success
	NCreateFailureProcessingFailure     = NewStatus(0x0110, StateFailure, "Processing failure")
	NCreateFailureDuplicateSOPInstance  = NewStatus(0x0111, StateFailure, "Duplicate SOP instance")
	NCreateFailureInvalidAttributeValue = NewStatus(0x0106, StateFailure, "Invalid attribute value")
	NCreateFailureMissingAttribute      = NewStatus(0x0120, StateFailure, "Missing attribute")
	NCreateFailureMissingAttributeValue = NewStatus(0x0121, StateFailure, "Missing attribute value")
	NCreateWarningAttributeListError    = NewStatus(0x0107, StateWarning, "Attribute list error")
	NCreateWarningAttributeValueOutOfRange = NewStatus(0x0116, StateWarning, "Attribute value out of range")
)

// N-DELETE specific statuses
var (
	NDeleteSuccess                      = Success
	NDeleteFailureProcessingFailure     = NewStatus(0x0110, StateFailure, "Processing failure")
	NDeleteFailureNoSuchObjectInstance  = NewStatus(0x0112, StateFailure, "No such object instance")
)

// LookupStatus looks up a status by its code.
// Returns a generic status if the specific code is not found.
func LookupStatus(code uint16) *Status {
	// Check exact matches in common statuses
	statuses := []*Status{
		Success,
		Pending, PendingWarning,
		Cancel,
		AttributeListError, AttributeValueOutOfRange,
		// Add more as needed
	}

	for _, s := range statuses {
		if s.Code == code {
			return s
		}
	}

	// Return a generic status based on code range
	if code == 0x0000 {
		return NewStatus(code, StateSuccess, "Success")
	} else if code == 0xFF00 || code == 0xFF01 {
		return NewStatus(code, StatePending, "Pending")
	} else if code == 0xFE00 {
		return NewStatus(code, StateCancel, "Cancel")
	} else if code >= 0x0001 && code <= 0x00FF {
		return NewStatus(code, StateWarning, "Warning")
	} else {
		return NewStatus(code, StateFailure, "Failure")
	}
}
