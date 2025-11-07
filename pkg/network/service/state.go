// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import "fmt"

// State represents the state of a DICOM service connection.
// These states follow the DICOM association state machine.
type State int

const (
	// StateIdle represents an idle state with no connection.
	StateIdle State = iota

	// StateAssociationRequested represents the state after an SCU has sent
	// an A-ASSOCIATE-RQ PDU and is waiting for a response.
	StateAssociationRequested

	// StateAssociationAccepted represents an established association.
	// In this state, DIMSE messages can be exchanged.
	StateAssociationAccepted

	// StateTransferring represents active DIMSE message transfer.
	// This is a sub-state of AssociationAccepted.
	StateTransferring

	// StateReleaseRequested represents the state after an A-RELEASE-RQ
	// has been sent and the service is waiting for an A-RELEASE-RP.
	StateReleaseRequested

	// StateClosed represents a closed connection.
	// No further operations are possible in this state.
	StateClosed

	// StateAborted represents an aborted connection.
	// This occurs after an A-ABORT PDU is sent or received.
	StateAborted
)

// String returns the string representation of the state.
func (s State) String() string {
	switch s {
	case StateIdle:
		return "Idle"
	case StateAssociationRequested:
		return "AssociationRequested"
	case StateAssociationAccepted:
		return "AssociationAccepted"
	case StateTransferring:
		return "Transferring"
	case StateReleaseRequested:
		return "ReleaseRequested"
	case StateClosed:
		return "Closed"
	case StateAborted:
		return "Aborted"
	default:
		return fmt.Sprintf("Unknown(%d)", s)
	}
}

// StateTransition represents a valid state transition.
type StateTransition struct {
	From  State
	To    State
	Event string // Description of the event that triggers this transition
}

// validTransitions defines all valid state transitions in the DICOM service state machine.
var validTransitions = []StateTransition{
	// From Idle
	{StateIdle, StateAssociationRequested, "Send A-ASSOCIATE-RQ"},
	{StateIdle, StateAssociationAccepted, "Receive A-ASSOCIATE-RQ and send A-ASSOCIATE-AC"},
	{StateIdle, StateClosed, "Connection error"},
	{StateIdle, StateAborted, "Send/Receive A-ABORT"},

	// From AssociationRequested
	{StateAssociationRequested, StateAssociationAccepted, "Receive A-ASSOCIATE-AC"},
	{StateAssociationRequested, StateClosed, "Receive A-ASSOCIATE-RJ"},
	{StateAssociationRequested, StateAborted, "Send/Receive A-ABORT"},
	{StateAssociationRequested, StateIdle, "Connection timeout"},

	// From AssociationAccepted
	{StateAssociationAccepted, StateTransferring, "Send/Receive DIMSE message"},
	{StateAssociationAccepted, StateReleaseRequested, "Send A-RELEASE-RQ"},
	{StateAssociationAccepted, StateClosed, "Receive A-RELEASE-RQ and send A-RELEASE-RP"},
	{StateAssociationAccepted, StateAborted, "Send/Receive A-ABORT"},

	// From Transferring
	{StateTransferring, StateAssociationAccepted, "DIMSE message complete"},
	{StateTransferring, StateReleaseRequested, "Send A-RELEASE-RQ"},
	{StateTransferring, StateAborted, "Send/Receive A-ABORT"},

	// From ReleaseRequested
	{StateReleaseRequested, StateClosed, "Receive A-RELEASE-RP"},
	{StateReleaseRequested, StateAborted, "Send/Receive A-ABORT"},
	{StateReleaseRequested, StateClosed, "Timeout waiting for A-RELEASE-RP"},

	// Terminal states (can only transition to themselves or stay)
	{StateClosed, StateClosed, "Already closed"},
	{StateAborted, StateAborted, "Already aborted"},
}

// IsValidTransition checks if a state transition is valid.
func IsValidTransition(from, to State) bool {
	// Same state is always valid
	if from == to {
		return true
	}

	// Check if transition exists in validTransitions
	for _, t := range validTransitions {
		if t.From == from && t.To == to {
			return true
		}
	}

	return false
}

// GetTransitionEvent returns the event description for a state transition.
// Returns empty string if the transition is not valid.
func GetTransitionEvent(from, to State) string {
	for _, t := range validTransitions {
		if t.From == from && t.To == to {
			return t.Event
		}
	}
	return ""
}

// CanSendDIMSE returns true if DIMSE messages can be sent in this state.
func (s State) CanSendDIMSE() bool {
	return s == StateAssociationAccepted || s == StateTransferring
}

// CanReceiveDIMSE returns true if DIMSE messages can be received in this state.
func (s State) CanReceiveDIMSE() bool {
	return s == StateAssociationAccepted || s == StateTransferring
}

// IsTerminal returns true if this is a terminal state (no further operations possible).
func (s State) IsTerminal() bool {
	return s == StateClosed || s == StateAborted
}

// IsActive returns true if the association is active (can exchange messages).
func (s State) IsActive() bool {
	return s == StateAssociationAccepted || s == StateTransferring
}
