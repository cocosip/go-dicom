// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// EncodeDIMSEMessage encodes a DIMSE message into command and data datasets.
// Returns:
//   - commandData: Encoded command dataset (always present)
//   - datasetData: Encoded data dataset (may be nil if no dataset)
//   - error: Any encoding error
//
// The command dataset is always encoded in Implicit VR Little Endian.
// The data dataset is encoded according to the transfer syntax.
func EncodeDIMSEMessage(msg dimse.Message, transferSyntax *transfer.Syntax) (commandData, datasetData []byte, err error) {
	// Get command dataset
	commandDS := msg.CommandDataset()
	if commandDS == nil {
		return nil, nil, fmt.Errorf("message has no command dataset")
	}

	// Calculate CommandGroupLength (0000,0000)
	// This is the total length of all command group elements AFTER the CommandGroupLength element
	// We need to encode once to get the size, then update CommandGroupLength and encode again

	// First, encode to calculate size
	tempBuf := &bytes.Buffer{}
	if err := writer.Write(tempBuf, commandDS,
		writer.WithTransferSyntax(transfer.ImplicitVRLittleEndian),
		writer.WithoutPreamble(),
		writer.WithKeepGroupLengths()); err != nil {
		return nil, nil, fmt.Errorf("failed to encode command dataset for length calculation: %w", err)
	}

	// Calculate CommandGroupLength value
	// Total size minus the CommandGroupLength element itself (tag + VL + value = 4 + 4 + 4 = 12 bytes)
	totalSize := len(tempBuf.Bytes())
	commandGroupLength := uint32(totalSize - 12)

	// Update CommandGroupLength in the dataset
	if err := commandDS.AddOrUpdate(element.NewUnsignedLong(tag.CommandGroupLength, []uint32{commandGroupLength})); err != nil {
		return nil, nil, fmt.Errorf("failed to update CommandGroupLength: %w", err)
	}

	// Encode command dataset again with correct CommandGroupLength (always Implicit VR Little Endian)
	commandBuf := &bytes.Buffer{}
	if err := writer.Write(commandBuf, commandDS,
		writer.WithTransferSyntax(transfer.ImplicitVRLittleEndian),
		writer.WithoutPreamble(),
		writer.WithKeepGroupLengths()); err != nil {
		return nil, nil, fmt.Errorf("failed to encode command dataset: %w", err)
	}
	commandData = commandBuf.Bytes()

	// Get data dataset (may be nil)
	dataDS := msg.DataDataset()
	if dataDS != nil {
		// Use the negotiated transfer syntax for data dataset
		if transferSyntax == nil {
			transferSyntax = transfer.ExplicitVRLittleEndian
		}

		dataBuf := &bytes.Buffer{}
		if err := writer.Write(dataBuf, dataDS,
			writer.WithTransferSyntax(transferSyntax),
			writer.WithoutPreamble()); err != nil {
			return nil, nil, fmt.Errorf("failed to encode data dataset: %w", err)
		}
		datasetData = dataBuf.Bytes()
	}

	return commandData, datasetData, nil
}

// DecodeDIMSEMessage decodes command and data bytes into datasets.
// This function parses raw DICOM datasets without file preamble or file meta information.
func DecodeDIMSEMessage(commandData, datasetData []byte, transferSyntax *transfer.Syntax) (*dataset.Dataset, *dataset.Dataset, error) {
	// Decode command dataset (always Implicit VR Little Endian per DICOM standard)
	commandDS, err := decodeRawDataset(commandData, transfer.ImplicitVRLittleEndian)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode command dataset: %w", err)
	}

	// Decode data dataset if present
	var dataDS *dataset.Dataset
	if len(datasetData) > 0 {
		if transferSyntax == nil {
			transferSyntax = transfer.ExplicitVRLittleEndian
		}

		dataDS, err = decodeRawDataset(datasetData, transferSyntax)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to decode data dataset: %w", err)
		}
	}

	return commandDS, dataDS, nil
}

// decodeRawDataset decodes a raw DICOM dataset from bytes.
// This is a simplified decoder that reads DICOM elements directly without
// expecting file preamble or file meta information.
func decodeRawDataset(data []byte, ts *transfer.Syntax) (*dataset.Dataset, error) {
	if len(data) == 0 {
		return dataset.New(), nil
	}

	ds := dataset.New()
	r := bytes.NewReader(data)

	// Determine byte order and VR mode from transfer syntax
	byteOrder := ts.Endian().ByteOrder()
	isExplicitVR := ts.IsExplicitVR()

	for r.Len() > 0 {
		// Read tag (4 bytes)
		if r.Len() < 4 {
			break // Not enough data for a tag
		}

		var group, elem uint16
		if err := binary.Read(r, byteOrder, &group); err != nil {
			return nil, fmt.Errorf("failed to read tag group: %w", err)
		}
		if err := binary.Read(r, byteOrder, &elem); err != nil {
			return nil, fmt.Errorf("failed to read tag element: %w", err)
		}

		t := tag.New(group, elem)

		var elemVR *vr.VR
		var valueLength uint32

		if isExplicitVR {
			// Read VR (2 bytes)
			if r.Len() < 2 {
				return nil, fmt.Errorf("not enough data to read VR for tag %s", t)
			}

			vrBytes := make([]byte, 2)
			if _, err := r.Read(vrBytes); err != nil {
				return nil, fmt.Errorf("failed to read VR for tag %s: %w", t, err)
			}

			var err error
			elemVR, err = vr.Parse(string(vrBytes))
			if err != nil {
				// Unknown VR, use UN (Unknown)
				elemVR = vr.UN
			}

			// Read length (2 or 6 bytes depending on VR)
			// VRs with 4-byte length: OB, OD, OF, OL, OV, OW, SQ, UC, UR, UT, UN
			if !elemVR.Is16bitLength() {
				// Reserved (2 bytes) + Length (4 bytes)
				reserved := make([]byte, 2)
				if _, err := r.Read(reserved); err != nil {
					return nil, fmt.Errorf("failed to read reserved bytes for tag %s: %w", t, err)
				}

				if err := binary.Read(r, byteOrder, &valueLength); err != nil {
					return nil, fmt.Errorf("failed to read 4-byte length for tag %s: %w", t, err)
				}
			} else {
				// Length (2 bytes)
				var length16 uint16
				if err := binary.Read(r, byteOrder, &length16); err != nil {
					return nil, fmt.Errorf("failed to read 2-byte length for tag %s: %w", t, err)
				}
				valueLength = uint32(length16)
			}
		} else {
			// Implicit VR: length is 4 bytes, VR is looked up from dictionary
			if err := binary.Read(r, byteOrder, &valueLength); err != nil {
				return nil, fmt.Errorf("failed to read length for tag %s: %w", t, err)
			}

			// For implicit VR, infer VR based on tag
			// This is a simplified implementation for common command dataset tags
			elemVR = inferVRFromTag(t)
		}

		// Handle undefined length (0xFFFFFFFF)
		if valueLength == 0xFFFFFFFF {
			// Sequences and items can have undefined length
			// For now, skip parsing sequence items properly and just create an empty sequence
			if elemVR == vr.SQ {
				seq := dataset.NewSequence(t)
				_ = ds.Add(seq)
				// TODO: Parse sequence items properly
				// For now, we'll read until we find the sequence delimiter
				continue
			}
			return nil, fmt.Errorf("undefined length not supported for non-sequence tag %s", t)
		}

		// Read value
		value := make([]byte, valueLength)
		if valueLength > 0 {
			if _, err := r.Read(value); err != nil {
				return nil, fmt.Errorf("failed to read value for tag %s: %w", t, err)
			}
		}

		// Create element from raw bytes
		el, err := createElementFromBytes(t, elemVR, value, byteOrder)
		if err != nil {
			return nil, fmt.Errorf("failed to create element for tag %s: %w", t, err)
		}

		if err := ds.Add(el); err != nil {
			return nil, fmt.Errorf("failed to add element %s to dataset: %w", t, err)
		}
	}

	return ds, nil
}

// inferVRFromTag infers the VR for a tag in implicit VR transfer syntax.
// This is primarily used for command datasets which use Implicit VR Little Endian.
func inferVRFromTag(t *tag.Tag) *vr.VR {
	// Common DIMSE command tags and their VRs
	// Reference: DICOM Part 7, Annex E - Command Dictionary
	switch {
	// Group 0000 - Command Group Elements
	case t.Group() == 0x0000:
		switch t.Element() {
		// UL (Unsigned Long) tags
		case 0x0000: // CommandGroupLength
			return vr.UL
		case 0x0001: // CommandLengthToEnd (Retired)
			return vr.UL

		// US (Unsigned Short) tags - Command Field, Message IDs, Priority, Status
		case 0x0100: // CommandField
			return vr.US
		case 0x0110: // MessageID
			return vr.US
		case 0x0120: // MessageIDBeingRespondedTo
			return vr.US
		case 0x0600: // MoveDestination (retired)
			return vr.US
		case 0x0700: // Priority
			return vr.US
		case 0x0800: // CommandDataSetType
			return vr.US
		case 0x0900: // Status
			return vr.US
		case 0x0901: // OffendingElement
			return vr.AT
		case 0x0902: // ErrorComment
			return vr.LO
		case 0x0903: // ErrorID
			return vr.US
		case 0x1000: // AffectedSOPInstanceUID
			return vr.UI
		case 0x1001: // RequestedSOPInstanceUID
			return vr.UI
		case 0x1002: // EventTypeID
			return vr.US
		case 0x1005: // AttributeIdentifierList
			return vr.AT
		case 0x1008: // ActionTypeID
			return vr.US
		case 0x1020: // NumberOfRemainingSuboperations
			return vr.US
		case 0x1021: // NumberOfCompletedSuboperations
			return vr.US
		case 0x1022: // NumberOfFailedSuboperations
			return vr.US
		case 0x1023: // NumberOfWarningSuboperations
			return vr.US
		case 0x1030: // MoveOriginatorApplicationEntityTitle
			return vr.AE
		case 0x1031: // MoveOriginatorMessageID
			return vr.US

		// UI (Unique Identifier) tags - SOP Class UIDs
		case 0x0002: // AffectedSOPClassUID
			return vr.UI
		case 0x0003: // RequestedSOPClassUID
			return vr.UI

		default:
			// Unknown command tag, use UN
			return vr.UN
		}
	// Most other tags
	default:
		return vr.UN
	}
}

// createElementFromBytes creates a DICOM element from raw bytes based on VR.
func createElementFromBytes(t *tag.Tag, elemVR *vr.VR, value []byte, byteOrder binary.ByteOrder) (element.Element, error) {
	// Handle string-based VRs
	switch elemVR {
	case vr.AE, vr.AS, vr.CS, vr.DA, vr.DS, vr.DT, vr.IS, vr.LO, vr.LT, vr.PN, vr.SH, vr.ST, vr.TM, vr.UC, vr.UI, vr.UR, vr.UT:
		// String VRs - parse as null-terminated or space-padded strings
		strValue := string(value)
		strValue = strings.TrimRight(strValue, "\x00 ") // Remove null bytes and trailing spaces

		// Some VRs can have multiple values separated by backslash
		var values []string
		if len(strValue) > 0 {
			values = strings.Split(strValue, "\\")
		}

		return element.NewString(t, elemVR, values), nil

	case vr.US: // Unsigned Short
		count := len(value) / 2
		values := make([]uint16, count)
		buf := bytes.NewReader(value)
		for i := 0; i < count; i++ {
			if err := binary.Read(buf, byteOrder, &values[i]); err != nil {
				return nil, err
			}
		}
		return element.NewUnsignedShort(t, values), nil

	case vr.UL: // Unsigned Long
		count := len(value) / 4
		values := make([]uint32, count)
		buf := bytes.NewReader(value)
		for i := 0; i < count; i++ {
			if err := binary.Read(buf, byteOrder, &values[i]); err != nil {
				return nil, err
			}
		}
		return element.NewUnsignedLong(t, values), nil

	case vr.SS: // Signed Short
		count := len(value) / 2
		values := make([]int16, count)
		buf := bytes.NewReader(value)
		for i := 0; i < count; i++ {
			if err := binary.Read(buf, byteOrder, &values[i]); err != nil {
				return nil, err
			}
		}
		return element.NewSignedShort(t, values), nil

	case vr.SL: // Signed Long
		count := len(value) / 4
		values := make([]int32, count)
		buf := bytes.NewReader(value)
		for i := 0; i < count; i++ {
			if err := binary.Read(buf, byteOrder, &values[i]); err != nil {
				return nil, err
			}
		}
		return element.NewSignedLong(t, values), nil

	case vr.FL: // Float
		count := len(value) / 4
		values := make([]float32, count)
		buf := bytes.NewReader(value)
		for i := 0; i < count; i++ {
			if err := binary.Read(buf, byteOrder, &values[i]); err != nil {
				return nil, err
			}
		}
		return element.NewFloat(t, values), nil

	case vr.FD: // Double
		count := len(value) / 8
		values := make([]float64, count)
		buf := bytes.NewReader(value)
		for i := 0; i < count; i++ {
			if err := binary.Read(buf, byteOrder, &values[i]); err != nil {
				return nil, err
			}
		}
		return element.NewDouble(t, values), nil

	case vr.AT: // Attribute Tag
		count := len(value) / 4
		values := make([]*tag.Tag, count)
		buf := bytes.NewReader(value)
		for i := 0; i < count; i++ {
			var group, elem uint16
			if err := binary.Read(buf, byteOrder, &group); err != nil {
				return nil, err
			}
			if err := binary.Read(buf, byteOrder, &elem); err != nil {
				return nil, err
			}
			values[i] = tag.New(group, elem)
		}
		return element.NewAttributeTag(t, values), nil

	case vr.OB, vr.OD, vr.OF, vr.OL, vr.OV, vr.OW, vr.UN:
		// Binary VRs - store as is
		return element.NewOtherByte(t, value), nil

	case vr.SQ:
		// Sequence - should have been handled earlier
		return dataset.NewSequence(t), nil

	default:
		// Unknown VR - treat as binary
		return element.NewOtherByte(t, value), nil
	}
}
