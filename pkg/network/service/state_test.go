// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import "testing"

func TestStateString(t *testing.T) {
	tests := []struct {
		state    State
		expected string
	}{
		{StateIdle, "Idle"},
		{StateAssociationRequested, "AssociationRequested"},
		{StateAssociationAccepted, "AssociationAccepted"},
		{StateTransferring, "Transferring"},
		{StateReleaseRequested, "ReleaseRequested"},
		{StateClosed, "Closed"},
		{StateAborted, "Aborted"},
		{State(999), "Unknown(999)"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.state.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestIsValidTransition(t *testing.T) {
	tests := []struct {
		from  State
		to    State
		valid bool
	}{
		// Valid transitions from Idle
		{StateIdle, StateAssociationRequested, true},
		{StateIdle, StateAssociationAccepted, true},
		{StateIdle, StateClosed, true},
		{StateIdle, StateAborted, true},

		// Invalid transitions from Idle
		{StateIdle, StateTransferring, false},
		{StateIdle, StateReleaseRequested, false},

		// Valid transitions from AssociationRequested
		{StateAssociationRequested, StateAssociationAccepted, true},
		{StateAssociationRequested, StateClosed, true},
		{StateAssociationRequested, StateAborted, true},
		{StateAssociationRequested, StateIdle, true},

		// Invalid transitions from AssociationRequested
		{StateAssociationRequested, StateTransferring, false},
		{StateAssociationRequested, StateReleaseRequested, false},

		// Valid transitions from AssociationAccepted
		{StateAssociationAccepted, StateTransferring, true},
		{StateAssociationAccepted, StateReleaseRequested, true},
		{StateAssociationAccepted, StateClosed, true},
		{StateAssociationAccepted, StateAborted, true},

		// Invalid transitions from AssociationAccepted
		{StateAssociationAccepted, StateIdle, false},
		{StateAssociationAccepted, StateAssociationRequested, false},

		// Valid transitions from Transferring
		{StateTransferring, StateAssociationAccepted, true},
		{StateTransferring, StateReleaseRequested, true},
		{StateTransferring, StateAborted, true},

		// Invalid transitions from Transferring
		{StateTransferring, StateIdle, false},
		{StateTransferring, StateAssociationRequested, false},
		{StateTransferring, StateClosed, false},

		// Valid transitions from ReleaseRequested
		{StateReleaseRequested, StateClosed, true},
		{StateReleaseRequested, StateAborted, true},

		// Invalid transitions from ReleaseRequested
		{StateReleaseRequested, StateIdle, false},
		{StateReleaseRequested, StateAssociationRequested, false},
		{StateReleaseRequested, StateAssociationAccepted, false},
		{StateReleaseRequested, StateTransferring, false},

		// Terminal states
		{StateClosed, StateClosed, true},
		{StateAborted, StateAborted, true},
		{StateClosed, StateIdle, false},
		{StateAborted, StateIdle, false},

		// Same state (always valid)
		{StateIdle, StateIdle, true},
		{StateAssociationRequested, StateAssociationRequested, true},
	}

	for _, tt := range tests {
		t.Run(tt.from.String()+"->"+tt.to.String(), func(t *testing.T) {
			result := IsValidTransition(tt.from, tt.to)
			if result != tt.valid {
				t.Errorf("IsValidTransition(%s, %s) = %v, expected %v",
					tt.from, tt.to, result, tt.valid)
			}
		})
	}
}

func TestGetTransitionEvent(t *testing.T) {
	tests := []struct {
		from     State
		to       State
		hasEvent bool
	}{
		{StateIdle, StateAssociationRequested, true},
		{StateAssociationRequested, StateAssociationAccepted, true},
		{StateAssociationAccepted, StateTransferring, true},
		{StateTransferring, StateAssociationAccepted, true},
		{StateAssociationAccepted, StateReleaseRequested, true},
		{StateReleaseRequested, StateClosed, true},

		// Invalid transitions should have no event
		{StateIdle, StateTransferring, false},
		{StateClosed, StateIdle, false},
	}

	for _, tt := range tests {
		t.Run(tt.from.String()+"->"+tt.to.String(), func(t *testing.T) {
			event := GetTransitionEvent(tt.from, tt.to)
			if tt.hasEvent {
				if event == "" {
					t.Errorf("Expected non-empty event for %s -> %s", tt.from, tt.to)
				}
			} else {
				if event != "" {
					t.Errorf("Expected empty event for invalid transition %s -> %s, got %s",
						tt.from, tt.to, event)
				}
			}
		})
	}
}

func TestStateCanSendDIMSE(t *testing.T) {
	tests := []struct {
		state State
		can   bool
	}{
		{StateIdle, false},
		{StateAssociationRequested, false},
		{StateAssociationAccepted, true},
		{StateTransferring, true},
		{StateReleaseRequested, false},
		{StateClosed, false},
		{StateAborted, false},
	}

	for _, tt := range tests {
		t.Run(tt.state.String(), func(t *testing.T) {
			result := tt.state.CanSendDIMSE()
			if result != tt.can {
				t.Errorf("State %s CanSendDIMSE() = %v, expected %v",
					tt.state, result, tt.can)
			}
		})
	}
}

func TestStateCanReceiveDIMSE(t *testing.T) {
	tests := []struct {
		state State
		can   bool
	}{
		{StateIdle, false},
		{StateAssociationRequested, false},
		{StateAssociationAccepted, true},
		{StateTransferring, true},
		{StateReleaseRequested, false},
		{StateClosed, false},
		{StateAborted, false},
	}

	for _, tt := range tests {
		t.Run(tt.state.String(), func(t *testing.T) {
			result := tt.state.CanReceiveDIMSE()
			if result != tt.can {
				t.Errorf("State %s CanReceiveDIMSE() = %v, expected %v",
					tt.state, result, tt.can)
			}
		})
	}
}

func TestStateIsTerminal(t *testing.T) {
	tests := []struct {
		state      State
		isTerminal bool
	}{
		{StateIdle, false},
		{StateAssociationRequested, false},
		{StateAssociationAccepted, false},
		{StateTransferring, false},
		{StateReleaseRequested, false},
		{StateClosed, true},
		{StateAborted, true},
	}

	for _, tt := range tests {
		t.Run(tt.state.String(), func(t *testing.T) {
			result := tt.state.IsTerminal()
			if result != tt.isTerminal {
				t.Errorf("State %s IsTerminal() = %v, expected %v",
					tt.state, result, tt.isTerminal)
			}
		})
	}
}

func TestStateIsActive(t *testing.T) {
	tests := []struct {
		state    State
		isActive bool
	}{
		{StateIdle, false},
		{StateAssociationRequested, false},
		{StateAssociationAccepted, true},
		{StateTransferring, true},
		{StateReleaseRequested, false},
		{StateClosed, false},
		{StateAborted, false},
	}

	for _, tt := range tests {
		t.Run(tt.state.String(), func(t *testing.T) {
			result := tt.state.IsActive()
			if result != tt.isActive {
				t.Errorf("State %s IsActive() = %v, expected %v",
					tt.state, result, tt.isActive)
			}
		})
	}
}

func TestStateTransitionScenario_NormalFlow(t *testing.T) {
	// Test a normal association flow: Idle -> Request -> Accept -> Transfer -> Release -> Close
	transitions := []struct {
		from State
		to   State
	}{
		{StateIdle, StateAssociationRequested},
		{StateAssociationRequested, StateAssociationAccepted},
		{StateAssociationAccepted, StateTransferring},
		{StateTransferring, StateAssociationAccepted},
		{StateAssociationAccepted, StateReleaseRequested},
		{StateReleaseRequested, StateClosed},
	}

	for i, tt := range transitions {
		if !IsValidTransition(tt.from, tt.to) {
			t.Errorf("Transition %d: %s -> %s should be valid in normal flow",
				i, tt.from, tt.to)
		}
	}
}

func TestStateTransitionScenario_AbortFlow(t *testing.T) {
	// Test abort can happen from most states
	states := []State{
		StateIdle,
		StateAssociationRequested,
		StateAssociationAccepted,
		StateTransferring,
		StateReleaseRequested,
	}

	for _, state := range states {
		if !IsValidTransition(state, StateAborted) {
			t.Errorf("Transition from %s to Aborted should be valid", state)
		}
	}
}
