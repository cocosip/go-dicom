// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

// Message represents a DIMSE message (either request or response).
// A DIMSE message consists of:
//   - Command Dataset: contains command fields (CommandField, MessageID, etc.)
//   - Data Dataset: contains the actual DICOM data (optional, depending on command type)
type Message interface {
	// CommandDataset returns the command dataset
	CommandDataset() *dataset.Dataset

	// DataDataset returns the data dataset (may be nil)
	DataDataset() *dataset.Dataset

	// PresentationContextID returns the presentation context ID
	PresentationContextID() byte

	// SetPresentationContextID sets the presentation context ID
	SetPresentationContextID(id byte)

	// CommandField returns the DIMSE command field
	CommandField() uint16

	// MessageID returns the message ID
	MessageID() uint16

	// String returns a human-readable representation
	String() string
}

// Request represents a DIMSE request message.
type Request interface {
	Message

	// Priority returns the priority of the request
	Priority() uint16

	// AffectedSOPClassUID returns the affected SOP Class UID (for C-STORE, C-ECHO, etc.)
	AffectedSOPClassUID() string
}

// Response represents a DIMSE response message.
type Response interface {
	Message

	// Status returns the status of the response
	Status() *status.Status

	// MessageIDBeingRespondedTo returns the message ID being responded to
	MessageIDBeingRespondedTo() uint16

	// IsPending returns true if this is a pending response (for multi-response operations like C-FIND)
	IsPending() bool

	// IsSuccess returns true if the status indicates success
	IsSuccess() bool

	// IsFailure returns true if the status indicates failure
	IsFailure() bool
}

// BaseMessage provides common functionality for all DIMSE messages.
type BaseMessage struct {
	command               *dataset.Dataset
	data                  *dataset.Dataset
	presentationContextID byte
	messageID             uint16 // Cached messageID, 0 means not assigned yet
}

// NewBaseMessage creates a new BaseMessage.
func NewBaseMessage(command *dataset.Dataset, data *dataset.Dataset) *BaseMessage {
	return &BaseMessage{
		command: command,
		data:    data,
	}
}

// CommandDataset returns the command dataset.
func (m *BaseMessage) CommandDataset() *dataset.Dataset {
	return m.command
}

// DataDataset returns the data dataset.
func (m *BaseMessage) DataDataset() *dataset.Dataset {
	return m.data
}

// PresentationContextID returns the presentation context ID.
func (m *BaseMessage) PresentationContextID() byte {
	return m.presentationContextID
}

// SetPresentationContextID sets the presentation context ID.
func (m *BaseMessage) SetPresentationContextID(id byte) {
	m.presentationContextID = id
}

// CommandField returns the DIMSE command field.
func (m *BaseMessage) CommandField() uint16 {
	if m.command == nil {
		return 0
	}
	value, _ := m.command.GetUInt16(tag.CommandField, 0)
	return value
}

// MessageID returns the message ID.
func (m *BaseMessage) MessageID() uint16 {
	// Return cached value if available
	if m.messageID != 0 {
		return m.messageID
	}

	// Otherwise read from command dataset
	if m.command == nil {
		return 0
	}
	value, _ := m.command.GetUInt16(tag.MessageID, 0)
	m.messageID = value
	return value
}

// SetMessageID sets the message ID. This is typically called by the Client/Association
// layer before sending the message.
func (m *BaseMessage) SetMessageID(id uint16) error {
	if m.command == nil {
		return fmt.Errorf("cannot set MessageID: no command dataset")
	}

	m.messageID = id
	return m.command.AddOrUpdate(element.NewUnsignedShort(tag.MessageID, []uint16{id}))
}

// GetCommandFieldValue gets a value from the command dataset.
func (m *BaseMessage) GetCommandFieldValue(t *tag.Tag) (interface{}, error) {
	if m.command == nil {
		return nil, fmt.Errorf("no command dataset")
	}
	elem, ok := m.command.Get(t)
	if !ok {
		return nil, fmt.Errorf("tag %s not found in command dataset", t.String())
	}
	return elem, nil
}

// GetCommandString gets a string value from the command dataset.
func (m *BaseMessage) GetCommandString(t *tag.Tag) (string, error) {
	if m.command == nil {
		return "", fmt.Errorf("no command dataset")
	}
	value, ok := m.command.GetString(t)
	if !ok {
		return "", fmt.Errorf("tag %s not found in command dataset", t.String())
	}
	return value, nil
}

// GetCommandUInt16 gets a uint16 value from the command dataset.
func (m *BaseMessage) GetCommandUInt16(t *tag.Tag) (uint16, error) {
	if m.command == nil {
		return 0, fmt.Errorf("no command dataset")
	}
	return m.command.GetUInt16(t, 0)
}

// String returns a human-readable representation.
func (m *BaseMessage) String() string {
	cmdField := m.CommandField()
	msgID := m.MessageID()
	return fmt.Sprintf("DIMSE Message [Command=0x%04X, MessageID=%d]", cmdField, msgID)
}

// BaseRequest provides common functionality for DIMSE request messages.
type BaseRequest struct {
	*BaseMessage
}

// NewBaseRequest creates a new BaseRequest.
func NewBaseRequest(command *dataset.Dataset, data *dataset.Dataset) *BaseRequest {
	return &BaseRequest{
		BaseMessage: NewBaseMessage(command, data),
	}
}

// Priority returns the priority of the request.
func (r *BaseRequest) Priority() uint16 {
	value, _ := r.GetCommandUInt16(tag.Priority)
	if value == 0 {
		return uint16(PriorityMedium) // Default to medium
	}
	return value
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *BaseRequest) AffectedSOPClassUID() string {
	value, _ := r.GetCommandString(tag.AffectedSOPClassUID)
	return value
}

// BaseResponse provides common functionality for DIMSE response messages.
type BaseResponse struct {
	*BaseMessage
}

// NewBaseResponse creates a new BaseResponse.
func NewBaseResponse(command *dataset.Dataset, data *dataset.Dataset) *BaseResponse {
	return &BaseResponse{
		BaseMessage: NewBaseMessage(command, data),
	}
}

// Status returns the status of the response.
func (r *BaseResponse) Status() *status.Status {
	code, err := r.GetCommandUInt16(tag.Status)
	if err != nil {
		return status.NewStatus(0xFFFF, "Unknown", "Status not found")
	}
	return status.LookupStatus(code)
}

// MessageIDBeingRespondedTo returns the message ID being responded to.
func (r *BaseResponse) MessageIDBeingRespondedTo() uint16 {
	value, _ := r.GetCommandUInt16(tag.MessageIDBeingRespondedTo)
	return value
}

// IsPending returns true if this is a pending response.
func (r *BaseResponse) IsPending() bool {
	st := r.Status()
	return st.IsPending()
}

// IsSuccess returns true if the status indicates success.
func (r *BaseResponse) IsSuccess() bool {
	st := r.Status()
	return st.IsSuccess()
}

// IsFailure returns true if the status indicates failure.
func (r *BaseResponse) IsFailure() bool {
	st := r.Status()
	return st.IsFailure()
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *BaseResponse) AffectedSOPClassUID() string {
	value, _ := r.GetCommandString(tag.AffectedSOPClassUID)
	return value
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
func (r *BaseResponse) AffectedSOPInstanceUID() string {
	value, _ := r.GetCommandString(tag.AffectedSOPInstanceUID)
	return value
}

// CreateCommandDataset creates a command dataset with common fields.
func CreateCommandDataset(commandField uint16, messageID uint16) *dataset.Dataset {
	ds := dataset.New()

	// CommandGroupLength will be calculated during encoding
	// We'll add a placeholder for now
	ds.Add(element.NewUnsignedLong(tag.CommandGroupLength, []uint32{0}))

	// AffectedSOPClassUID - will be set by specific message types
	// CommandField
	ds.Add(element.NewUnsignedShort(tag.CommandField, []uint16{commandField}))

	// MessageID
	ds.Add(element.NewUnsignedShort(tag.MessageID, []uint16{messageID}))

	// CommandDataSetType - 0x0101 means no dataset present
	ds.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))

	return ds
}
