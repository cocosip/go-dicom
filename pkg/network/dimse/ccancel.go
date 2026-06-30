// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import (
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// CCancelRequest represents a C-CANCEL-RQ message.
// Per DICOM PS3.7 §9.3.5, C-CANCEL-RQ has no response and is used to
// cancel a pending C-FIND, C-MOVE, or C-GET operation.
type CCancelRequest struct {
	*BaseRequest
}

// NewCCancelRequest creates a new outgoing C-CANCEL-RQ request.
func NewCCancelRequest(messageIDBeingRespondedTo uint16) *CCancelRequest {
	command := CreateCommandDataset(uint16(CommandCCancelRQ), 0)
	_ = command.AddOrUpdate(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))
	return NewCCancelRequestFromCommand(command)
}

// NewCCancelRequestFromCommand creates a C-CANCEL-RQ request from a decoded command dataset.
func NewCCancelRequestFromCommand(command *dataset.Dataset) *CCancelRequest {
	return &CCancelRequest{
		BaseRequest: NewBaseRequest(command, nil),
	}
}

// MessageIDBeingRespondedTo returns the message ID of the request being cancelled.
func (r *CCancelRequest) MessageIDBeingRespondedTo() uint16 {
	value, _ := r.GetCommandUInt16(tag.MessageIDBeingRespondedTo)
	return value
}

// CommandField returns the command field for C-CANCEL-RQ.
func (r *CCancelRequest) CommandField() uint16 {
	return uint16(CommandCCancelRQ)
}

// ExecuteMode returns the execution mode (always "" for cancel).
func (r *CCancelRequest) ExecuteMode() string {
	return ""
}
