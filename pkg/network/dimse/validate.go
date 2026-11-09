// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package dimse

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// ValidateRequestForSend checks that an outgoing DIMSE request carries the
// mandatory identifier fields of its operation (PS3.7 9.1.1, 9.1.2 and 10.1).
// It is meant to be called by the service send path before encoding a request.
//
// C-ECHO and C-CANCEL carry no mandatory identifier fields and always pass.
// The switch covers the built-in request types; a custom Request implementation
// is not introspected.
func ValidateRequestForSend(req Request) error {
	if req == nil {
		return fmt.Errorf("dimse: request is nil")
	}

	switch r := req.(type) {
	case *CEchoRequest, *CCancelRequest:
		return nil

	case *CStoreRequest:
		if err := requireString("C-STORE", "AffectedSOPClassUID", r.AffectedSOPClassUID()); err != nil {
			return err
		}
		return requireString("C-STORE", "AffectedSOPInstanceUID", r.AffectedSOPInstanceUID())

	case *CFindRequest:
		return requireString("C-FIND", "AffectedSOPClassUID", r.AffectedSOPClassUID())

	case *CGetRequest:
		return requireString("C-GET", "AffectedSOPClassUID", r.AffectedSOPClassUID())

	case *CMoveRequest:
		if err := requireString("C-MOVE", "AffectedSOPClassUID", r.AffectedSOPClassUID()); err != nil {
			return err
		}
		return requireString("C-MOVE", "MoveDestination", r.MoveDestination())

	case *NGetRequest:
		if err := requireString("N-GET", "RequestedSOPClassUID", r.RequestedSOPClassUID()); err != nil {
			return err
		}
		return requireString("N-GET", "RequestedSOPInstanceUID", r.RequestedSOPInstanceUID())

	case *NSetRequest:
		if err := requireString("N-SET", "RequestedSOPClassUID", r.RequestedSOPClassUID()); err != nil {
			return err
		}
		return requireString("N-SET", "RequestedSOPInstanceUID", r.RequestedSOPInstanceUID())

	case *NDeleteRequest:
		if err := requireString("N-DELETE", "RequestedSOPClassUID", r.RequestedSOPClassUID()); err != nil {
			return err
		}
		return requireString("N-DELETE", "RequestedSOPInstanceUID", r.RequestedSOPInstanceUID())

	case *NActionRequest:
		if err := requireString("N-ACTION", "RequestedSOPClassUID", r.RequestedSOPClassUID()); err != nil {
			return err
		}
		return requireString("N-ACTION", "RequestedSOPInstanceUID", r.RequestedSOPInstanceUID())

	case *NCreateRequest:
		// The instance UID may be empty: the SCP assigns it on creation.
		return requireString("N-CREATE", "AffectedSOPClassUID", r.AffectedSOPClassUID())

	case *NEventReportRequest:
		if err := requireString("N-EVENT-REPORT", "AffectedSOPClassUID", r.AffectedSOPClassUID()); err != nil {
			return err
		}
		if err := requireString("N-EVENT-REPORT", "AffectedSOPInstanceUID", r.AffectedSOPInstanceUID()); err != nil {
			return err
		}
		// Event Type ID is Type 1; 0 is a legal value, so only its presence can
		// be checked.
		if !r.CommandDataset().Contains(tag.EventTypeID) {
			return fmt.Errorf("N-EVENT-REPORT: EventTypeID is required")
		}
		return nil

	default:
		return nil
	}
}

// requireString returns an error naming the operation and the empty field.
func requireString(operation, field, value string) error {
	if value == "" {
		return fmt.Errorf("%s: %s is required", operation, field)
	}
	return nil
}
