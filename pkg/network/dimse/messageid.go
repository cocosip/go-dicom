// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import "sync/atomic"

// MessageIDGenerator generates unique message IDs for DIMSE messages.
// It is thread-safe and can be used concurrently.
//
// MessageID is used to correlate DIMSE requests with their responses.
// Each request should have a unique MessageID within an association.
type MessageIDGenerator struct {
	counter uint32
}

// NewMessageIDGenerator creates a new MessageID generator.
func NewMessageIDGenerator() *MessageIDGenerator {
	return &MessageIDGenerator{counter: 0}
}

// Next returns the next available message ID.
// MessageIDs range from 1 to 65535 (uint16 range, excluding 0).
// This method is thread-safe.
func (g *MessageIDGenerator) Next() uint16 {
	val := atomic.AddUint32(&g.counter, 1)
	return uint16(((val - 1) % 0xFFFF) + 1) // #nosec G115 -- value is in 1..65535
}

// Reset resets the counter to 0.
// This is typically called when starting a new association.
func (g *MessageIDGenerator) Reset() {
	atomic.StoreUint32(&g.counter, 0)
}

// AssignMessageID assigns a message ID to a request if it doesn't have one.
// If the message already has a non-zero MessageID, it is left unchanged.
// Returns the assigned MessageID.
func (g *MessageIDGenerator) AssignMessageID(msg Request) (uint16, error) {
	currentID := msg.MessageID()

	// If already assigned, return it
	if currentID != 0 {
		return currentID, nil
	}

	// Generate new MessageID
	newID := g.Next()

	if err := msg.SetMessageID(newID); err != nil {
		return 0, err
	}
	return newID, nil
}
