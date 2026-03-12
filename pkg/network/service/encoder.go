// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/dict"
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

	r := bytes.NewReader(data)

	// Determine byte order and VR mode from transfer syntax
	byteOrder := ts.Endian().ByteOrder()
	isExplicitVR := ts.IsExplicitVR()

	return decodeDatasetFromReader(r, byteOrder, isExplicitVR)
}

func decodeDatasetFromReader(r *bytes.Reader, byteOrder binary.ByteOrder, isExplicitVR bool) (*dataset.Dataset, error) {
	ds := dataset.New()

	for r.Len() > 0 {
		if r.Len() < 4 {
			break
		}

		t, err := readTagFromReader(r, byteOrder)
		if err != nil {
			return nil, err
		}

		el, err := decodeElementWithTag(r, byteOrder, isExplicitVR, t)
		if err != nil {
			return nil, err
		}

		if err := ds.Add(el); err != nil {
			return nil, fmt.Errorf("failed to add element %s to dataset: %w", t, err)
		}
	}

	return ds, nil
}

func readTagFromReader(r *bytes.Reader, byteOrder binary.ByteOrder) (*tag.Tag, error) {
	var group, elem uint16
	if err := binary.Read(r, byteOrder, &group); err != nil {
		return nil, fmt.Errorf("failed to read tag group: %w", err)
	}
	if err := binary.Read(r, byteOrder, &elem); err != nil {
		return nil, fmt.Errorf("failed to read tag element: %w", err)
	}
	return tag.New(group, elem), nil
}

func decodeElementWithTag(r *bytes.Reader, byteOrder binary.ByteOrder, isExplicitVR bool, t *tag.Tag) (element.Element, error) {
	elemVR, valueLength, err := readVROrInferAndLength(r, byteOrder, isExplicitVR, t)
	if err != nil {
		return nil, err
	}

	if elemVR == vr.SQ {
		return decodeSequenceElement(r, byteOrder, isExplicitVR, t, valueLength)
	}

	if valueLength == 0xFFFFFFFF {
		return nil, fmt.Errorf("undefined length not supported for non-sequence tag %s", t)
	}

	value := make([]byte, valueLength)
	if valueLength > 0 {
		if _, err := io.ReadFull(r, value); err != nil {
			return nil, fmt.Errorf("failed to read value for tag %s: %w", t, err)
		}
	}

	el, err := createElementFromBytes(t, elemVR, value, byteOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to create element for tag %s: %w", t, err)
	}
	return el, nil
}

func readVROrInferAndLength(r *bytes.Reader, byteOrder binary.ByteOrder, isExplicitVR bool, t *tag.Tag) (*vr.VR, uint32, error) {
	var elemVR *vr.VR
	var valueLength uint32

	if isExplicitVR {
		if r.Len() < 2 {
			return nil, 0, fmt.Errorf("not enough data to read VR for tag %s", t)
		}

		vrBytes := make([]byte, 2)
		if _, err := io.ReadFull(r, vrBytes); err != nil {
			return nil, 0, fmt.Errorf("failed to read VR for tag %s: %w", t, err)
		}

		var err error
		elemVR, err = vr.Parse(string(vrBytes))
		if err != nil {
			elemVR = vr.UN
		}

		if !elemVR.Is16bitLength() {
			reserved := make([]byte, 2)
			if _, err := io.ReadFull(r, reserved); err != nil {
				return nil, 0, fmt.Errorf("failed to read reserved bytes for tag %s: %w", t, err)
			}
			if err := binary.Read(r, byteOrder, &valueLength); err != nil {
				return nil, 0, fmt.Errorf("failed to read 4-byte length for tag %s: %w", t, err)
			}
		} else {
			var length16 uint16
			if err := binary.Read(r, byteOrder, &length16); err != nil {
				return nil, 0, fmt.Errorf("failed to read 2-byte length for tag %s: %w", t, err)
			}
			valueLength = uint32(length16)
		}
	} else {
		if err := binary.Read(r, byteOrder, &valueLength); err != nil {
			return nil, 0, fmt.Errorf("failed to read length for tag %s: %w", t, err)
		}
		elemVR = inferVRFromTag(t)
	}

	return elemVR, valueLength, nil
}

func decodeSequenceElement(r *bytes.Reader, byteOrder binary.ByteOrder, isExplicitVR bool, t *tag.Tag, valueLength uint32) (*dataset.Sequence, error) {
	seq := dataset.NewSequence(t)

	if valueLength == 0xFFFFFFFF {
		for {
			if r.Len() < 4 {
				return nil, fmt.Errorf("unexpected EOF while reading undefined-length sequence %s", t)
			}

			itemTag, err := readTagFromReader(r, byteOrder)
			if err != nil {
				return nil, err
			}

			if isSequenceDelimitationTag(itemTag) {
				var delimLength uint32
				if err := binary.Read(r, byteOrder, &delimLength); err != nil {
					return nil, fmt.Errorf("failed to read sequence delimiter length for %s: %w", t, err)
				}
				break
			}

			if !isItemTag(itemTag) {
				return nil, fmt.Errorf("expected item tag in sequence %s, got %s", t, itemTag)
			}

			var itemLength uint32
			if err := binary.Read(r, byteOrder, &itemLength); err != nil {
				return nil, fmt.Errorf("failed to read item length in sequence %s: %w", t, err)
			}

			item, err := decodeItemDataset(r, byteOrder, isExplicitVR, itemLength)
			if err != nil {
				return nil, err
			}
			seq.AddItem(item)
		}
		return seq, nil
	}

	sequenceData := make([]byte, valueLength)
	if _, err := io.ReadFull(r, sequenceData); err != nil {
		return nil, fmt.Errorf("failed to read sequence data for %s: %w", t, err)
	}

	seqReader := bytes.NewReader(sequenceData)
	for seqReader.Len() > 0 {
		if seqReader.Len() < 4 {
			break
		}

		itemTag, err := readTagFromReader(seqReader, byteOrder)
		if err != nil {
			return nil, err
		}

		if isSequenceDelimitationTag(itemTag) {
			var delimLength uint32
			if err := binary.Read(seqReader, byteOrder, &delimLength); err != nil {
				return nil, fmt.Errorf("failed to read sequence delimiter length for %s: %w", t, err)
			}
			break
		}

		if !isItemTag(itemTag) {
			return nil, fmt.Errorf("expected item tag in sequence %s, got %s", t, itemTag)
		}

		var itemLength uint32
		if err := binary.Read(seqReader, byteOrder, &itemLength); err != nil {
			return nil, fmt.Errorf("failed to read item length in sequence %s: %w", t, err)
		}

		item, err := decodeItemDataset(seqReader, byteOrder, isExplicitVR, itemLength)
		if err != nil {
			return nil, err
		}
		seq.AddItem(item)
	}

	return seq, nil
}

func decodeItemDataset(r *bytes.Reader, byteOrder binary.ByteOrder, isExplicitVR bool, itemLength uint32) (*dataset.Dataset, error) {
	// Undefined-length item: read until Item Delimitation Item.
	if itemLength == 0xFFFFFFFF {
		item := dataset.New()
		for {
			if r.Len() < 4 {
				return nil, fmt.Errorf("unexpected EOF while reading undefined-length item")
			}

			t, err := readTagFromReader(r, byteOrder)
			if err != nil {
				return nil, err
			}

			if isItemDelimitationTag(t) {
				var delimLength uint32
				if err := binary.Read(r, byteOrder, &delimLength); err != nil {
					return nil, fmt.Errorf("failed to read item delimiter length: %w", err)
				}
				break
			}

			el, err := decodeElementWithTag(r, byteOrder, isExplicitVR, t)
			if err != nil {
				return nil, err
			}
			if err := item.Add(el); err != nil {
				return nil, fmt.Errorf("failed to add element %s to item: %w", t, err)
			}
		}
		return item, nil
	}

	if itemLength == 0 {
		return dataset.New(), nil
	}

	itemData := make([]byte, itemLength)
	if _, err := io.ReadFull(r, itemData); err != nil {
		return nil, fmt.Errorf("failed to read item data: %w", err)
	}

	return decodeDatasetFromReader(bytes.NewReader(itemData), byteOrder, isExplicitVR)
}

func isItemTag(t *tag.Tag) bool {
	return t.Group() == 0xFFFE && t.Element() == 0xE000
}

func isItemDelimitationTag(t *tag.Tag) bool {
	return t.Group() == 0xFFFE && t.Element() == 0xE00D
}

func isSequenceDelimitationTag(t *tag.Tag) bool {
	return t.Group() == 0xFFFE && t.Element() == 0xE0DD
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
		case 0x0600: // MoveDestination
			return vr.AE
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
	// For all other tags, look up the VR from the DICOM dictionary.
	// This is required for Implicit VR transfer syntaxes where the VR is
	// not present in the byte stream and must be inferred from the tag.
	default:
		if entry := dict.Default().Lookup(t); entry != nil {
			if vrs := entry.ValueRepresentations(); len(vrs) > 0 {
				return vrs[0]
			}
		}
		return vr.UN
	}
}

// createElementFromBytes creates a DICOM element from raw bytes based on VR.
func createElementFromBytes(t *tag.Tag, elemVR *vr.VR, value []byte, byteOrder binary.ByteOrder) (element.Element, error) {
	if isStringVR(elemVR) {
		strValue := strings.TrimRight(string(value), "\x00 ")
		var values []string
		if len(strValue) > 0 {
			values = strings.Split(strValue, "\\")
		}
		return element.NewString(t, elemVR, values), nil
	}

	if numericEl, ok, err := createNumericElement(t, elemVR, value, byteOrder); ok || err != nil {
		return numericEl, err
	}

	switch elemVR {
	case vr.OB:
		return element.NewOtherByte(t, value), nil
	case vr.OW:
		return element.NewOtherWord(t, value), nil
	case vr.OD:
		return element.NewOtherDouble(t, value), nil
	case vr.OF:
		return element.NewOtherFloat(t, value), nil
	case vr.OL:
		return element.NewOtherLong(t, value), nil
	case vr.OV:
		return element.NewOtherVeryLong(t, value), nil
	case vr.UN:
		return element.NewUnknown(t, value), nil
	case vr.SQ:
		// Sequence should be handled before reaching this function.
		return dataset.NewSequence(t), nil
	default:
		return element.NewUnknown(t, value), nil
	}
}

func isStringVR(elemVR *vr.VR) bool {
	switch elemVR {
	case vr.AE, vr.AS, vr.CS, vr.DA, vr.DS, vr.DT, vr.IS, vr.LO, vr.LT, vr.PN, vr.SH, vr.ST, vr.TM, vr.UC, vr.UI, vr.UR, vr.UT:
		return true
	default:
		return false
	}
}

func createNumericElement(t *tag.Tag, elemVR *vr.VR, value []byte, byteOrder binary.ByteOrder) (element.Element, bool, error) {
	switch elemVR {
	case vr.US:
		values, err := readTypedValues[uint16](value, byteOrder)
		if err != nil {
			return nil, true, err
		}
		return element.NewUnsignedShort(t, values), true, nil
	case vr.UL:
		values, err := readTypedValues[uint32](value, byteOrder)
		if err != nil {
			return nil, true, err
		}
		return element.NewUnsignedLong(t, values), true, nil
	case vr.SS:
		values, err := readTypedValues[int16](value, byteOrder)
		if err != nil {
			return nil, true, err
		}
		return element.NewSignedShort(t, values), true, nil
	case vr.SL:
		values, err := readTypedValues[int32](value, byteOrder)
		if err != nil {
			return nil, true, err
		}
		return element.NewSignedLong(t, values), true, nil
	case vr.FL:
		values, err := readTypedValues[float32](value, byteOrder)
		if err != nil {
			return nil, true, err
		}
		return element.NewFloat(t, values), true, nil
	case vr.FD:
		values, err := readTypedValues[float64](value, byteOrder)
		if err != nil {
			return nil, true, err
		}
		return element.NewDouble(t, values), true, nil
	case vr.AT:
		values, err := readAttributeTags(value, byteOrder)
		if err != nil {
			return nil, true, err
		}
		return element.NewAttributeTag(t, values), true, nil
	default:
		return nil, false, nil
	}
}

func readTypedValues[T any](value []byte, byteOrder binary.ByteOrder) ([]T, error) {
	var zero T
	elementSize := binary.Size(zero)
	if elementSize <= 0 {
		return nil, fmt.Errorf("invalid element size for typed value")
	}
	count := len(value) / elementSize
	values := make([]T, count)
	buf := bytes.NewReader(value)
	for i := 0; i < count; i++ {
		if err := binary.Read(buf, byteOrder, &values[i]); err != nil {
			return nil, err
		}
	}
	return values, nil
}

func readAttributeTags(value []byte, byteOrder binary.ByteOrder) ([]*tag.Tag, error) {
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
	return values, nil
}
