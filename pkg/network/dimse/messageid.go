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
	// Atomically increment the counter
	val := atomic.AddUint32(&g.counter, 1)

	// Take the lower 16 bits
	msgID := uint16(val & 0xFFFF)

	// Ensure MessageID is never 0 (0 means unassigned)
	if msgID == 0 {
		msgID = 1
	}

	return msgID
}

// Reset resets the counter to 0.
// This is typically called when starting a new association.
func (g *MessageIDGenerator) Reset() {
	atomic.StoreUint32(&g.counter, 0)
}

// AssignMessageID assigns a message ID to a request if it doesn't have one.
// If the message already has a non-zero MessageID, it is left unchanged.
// Returns the assigned MessageID.
func (g *MessageIDGenerator) AssignMessageID(msg Message) (uint16, error) {
	currentID := msg.MessageID()

	// If already assigned, return it
	if currentID != 0 {
		return currentID, nil
	}

	// Generate new MessageID
	newID := g.Next()

	// Try to set it (only works on BaseMessage)
	if base, ok := msg.(*BaseMessage); ok {
		if err := base.SetMessageID(newID); err != nil {
			return 0, err
		}
		return newID, nil
	}

	// For Request messages
	if req, ok := msg.(Request); ok {
		if baseReq, ok := req.(*BaseRequest); ok {
			if err := baseReq.SetMessageID(newID); err != nil {
				return 0, err
			}
			return newID, nil
		}
	}

	// If we can't cast, try to access BaseMessage through the interface
	// This works for CEchoRequest, CStoreRequest, etc.
	switch m := msg.(type) {
	case *CEchoRequest:
		return g.assignToBaseRequest(m.BaseRequest, newID)
	case *CStoreRequest:
		return g.assignToBaseRequest(m.BaseRequest, newID)
	case *CFindRequest:
		return g.assignToBaseRequest(m.BaseRequest, newID)
	case *NEventReportRequest:
		return g.assignToBaseRequest(m.BaseRequest, newID)
	case *NGetRequest:
		return g.assignToBaseRequest(m.BaseRequest, newID)
	case *NSetRequest:
		return g.assignToBaseRequest(m.BaseRequest, newID)
	case *NActionRequest:
		return g.assignToBaseRequest(m.BaseRequest, newID)
	case *NCreateRequest:
		return g.assignToBaseRequest(m.BaseRequest, newID)
	case *NDeleteRequest:
		return g.assignToBaseRequest(m.BaseRequest, newID)
	}

	return 0, nil
}

func (g *MessageIDGenerator) assignToBaseRequest(req *BaseRequest, id uint16) (uint16, error) {
	if err := req.SetMessageID(id); err != nil {
		return 0, err
	}
	return id, nil
}
