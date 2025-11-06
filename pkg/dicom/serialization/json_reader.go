// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package serialization

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// FromJSON deserializes DICOM JSON to a Dataset
func FromJSON(data []byte, opts ...JSONOption) (*dataset.Dataset, error) {
	cfg := defaultJSONConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	r := &jsonReader{
		config: cfg,
	}

	return r.readDataset(data)
}

// jsonReader handles JSON deserialization
type jsonReader struct {
	config jsonConfig
}

// readDataset reads a JSON object into a Dataset
func (r *jsonReader) readDataset(data []byte) (*dataset.Dataset, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	ds := dataset.New()

	for tagStr, valueRaw := range raw {
		// Parse tag
		t, err := r.parseTag(tagStr)
		if err != nil {
			return nil, fmt.Errorf("invalid tag %q: %w", tagStr, err)
		}

		// Parse element
		elem, err := r.readElement(t, valueRaw)
		if err != nil {
			return nil, fmt.Errorf("failed to read element %s: %w", tagStr, err)
		}

		if elem != nil {
			ds.Add(elem)
		}
	}

	return ds, nil
}

// parseTag parses a tag from string (either "GGGGEEEE" or keyword)
func (r *jsonReader) parseTag(s string) (tag.Tag, error) {
	// Try to parse as hex tag first
	if len(s) == 8 {
		// Try as GGGGEEEE
		t, err := tag.Parse(fmt.Sprintf("(%s,%s)", s[0:4], s[4:8]))
		if err == nil {
			return *t, nil
		}
	}

	// Try as keyword
	t, err := tag.ParseKeyword(s)
	if err != nil {
		return tag.Tag{}, fmt.Errorf("invalid tag or keyword: %s", s)
	}
	return t, nil
}

// readElement reads a JSON element object
func (r *jsonReader) readElement(t tag.Tag, data json.RawMessage) (element.Element, error) {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, fmt.Errorf("element is not an object: %w", err)
	}

	// Get VR
	vrRaw, ok := obj["vr"]
	if !ok {
		return nil, fmt.Errorf("missing 'vr' field")
	}
	var vrStr string
	if err := json.Unmarshal(vrRaw, &vrStr); err != nil {
		return nil, fmt.Errorf("invalid 'vr' field: %w", err)
	}

	vrObj, err := vr.Parse(vrStr)
	if err != nil {
		return nil, fmt.Errorf("invalid VR %q: %w", vrStr, err)
	}

	// Check for Value, InlineBinary, or BulkDataURI
	valueRaw, hasValue := obj["Value"]
	inlineBinaryRaw, hasInlineBinary := obj["InlineBinary"]
	bulkDataURIRaw, hasBulkDataURI := obj["BulkDataURI"]

	// Handle different value types
	if hasBulkDataURI {
		// BulkDataURI
		return r.readBulkDataURI(&t, vrObj, bulkDataURIRaw)
	}

	if hasInlineBinary {
		// Binary data
		return r.readInlineBinary(&t, vrObj, inlineBinaryRaw)
	}

	if hasValue {
		// Regular value
		return r.readValue(&t, vrObj, valueRaw)
	}

	// Empty element
	return r.createEmptyElement(&t, vrObj)
}

// readValue reads the Value field based on VR
func (r *jsonReader) readValue(t *tag.Tag, vrObj *vr.VR, data json.RawMessage) (element.Element, error) {
	vrCode := vrObj.Code()

	switch vrCode {
	case vr.CodeSQ:
		return r.readSequence(t, data)
	case vr.CodePN:
		return r.readPersonName(t, vrObj, data)
	case vr.CodeAT:
		return r.readAttributeTag(t, vrObj, data)
	case vr.CodeFL:
		return r.readFloat32Array(t, data)
	case vr.CodeFD:
		return r.readFloat64Array(t, data)
	case vr.CodeSL:
		return r.readInt32Array(t, data)
	case vr.CodeSS:
		return r.readInt16Array(t, data)
	case vr.CodeSV:
		return r.readInt64Array(t, data)
	case vr.CodeUL:
		return r.readUint32Array(t, data)
	case vr.CodeUS:
		return r.readUint16Array(t, data)
	case vr.CodeUV:
		return r.readUint64Array(t, data)
	case vr.CodeIS, vr.CodeDS:
		return r.readNumericStringArray(t, vrObj, data)
	default:
		// String types
		return r.readStringArray(t, vrObj, data)
	}
}

// readStringArray reads a string array
func (r *jsonReader) readStringArray(t *tag.Tag, vr *vr.VR, data json.RawMessage) (element.Element, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, err
	}

	values := make([]string, len(rawValues))
	for i, raw := range rawValues {
		// Check if null
		var str string
		if err := json.Unmarshal(raw, &str); err != nil {
			// Could be null
			values[i] = ""
		} else {
			values[i] = str
		}
	}

	return element.NewString(t, vr, values), nil
}

// readNumericStringArray reads IS or DS arrays
func (r *jsonReader) readNumericStringArray(t *tag.Tag, vrObj *vr.VR, data json.RawMessage) (element.Element, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, err
	}

	values := make([]string, len(rawValues))
	for i, raw := range rawValues {
		// Try as number first
		var num float64
		if err := json.Unmarshal(raw, &num); err == nil {
			if vrObj.Code() == vr.CodeIS {
				values[i] = strconv.FormatInt(int64(num), 10)
			} else {
				values[i] = strconv.FormatFloat(num, 'f', -1, 64)
			}
		} else {
			// Try as string
			var str string
			if err := json.Unmarshal(raw, &str); err != nil {
				values[i] = ""
			} else {
				values[i] = str
			}
		}
	}

	return element.NewString(t, vrObj, values), nil
}

// readFloat32Array reads a float32 array (FL)
func (r *jsonReader) readFloat32Array(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, err
	}

	values := make([]float32, len(rawValues))
	for i, raw := range rawValues {
		val, err := r.parseFloat32(raw)
		if err != nil {
			return nil, fmt.Errorf("invalid float at index %d: %w", i, err)
		}
		values[i] = val
	}

	return element.NewFloat(t, values), nil
}

// readFloat64Array reads a float64 array (FD)
func (r *jsonReader) readFloat64Array(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, err
	}

	values := make([]float64, len(rawValues))
	for i, raw := range rawValues {
		val, err := r.parseFloat64(raw)
		if err != nil {
			return nil, fmt.Errorf("invalid double at index %d: %w", i, err)
		}
		values[i] = val
	}

	return element.NewDouble(t, values), nil
}

// parseFloat32 parses a float32 from JSON (number or special string)
func (r *jsonReader) parseFloat32(data json.RawMessage) (float32, error) {
	// Try as number first
	var num float32
	if err := json.Unmarshal(data, &num); err == nil {
		return num, nil
	}

	// Try as string (for NaN, Infinity, etc.)
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return 0, err
	}

	switch str {
	case "NaN":
		return float32(math.NaN()), nil
	case "Infinity":
		return float32(math.Inf(1)), nil
	case "-Infinity":
		return float32(math.Inf(-1)), nil
	default:
		val, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return 0, fmt.Errorf("invalid float string: %s", str)
		}
		return float32(val), nil
	}
}

// parseFloat64 parses a float64 from JSON (number or special string)
func (r *jsonReader) parseFloat64(data json.RawMessage) (float64, error) {
	// Try as number first
	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		return num, nil
	}

	// Try as string (for NaN, Infinity, etc.)
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return 0, err
	}

	switch str {
	case "NaN":
		return math.NaN(), nil
	case "Infinity":
		return math.Inf(1), nil
	case "-Infinity":
		return math.Inf(-1), nil
	default:
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid double string: %s", str)
		}
		return val, nil
	}
}

// readInt16Array reads an int16 array (SS)
func (r *jsonReader) readInt16Array(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var values []int16
	if err := json.Unmarshal(data, &values); err != nil {
		return nil, err
	}
	return element.NewSignedShort(t, values), nil
}

// readInt32Array reads an int32 array (SL)
func (r *jsonReader) readInt32Array(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var values []int32
	if err := json.Unmarshal(data, &values); err != nil {
		return nil, err
	}
	return element.NewSignedLong(t, values), nil
}

// readInt64Array reads an int64 array (SV)
func (r *jsonReader) readInt64Array(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, err
	}

	values := make([]int64, len(rawValues))
	for i, raw := range rawValues {
		// Try as number first, then as string
		var num int64
		if err := json.Unmarshal(raw, &num); err == nil {
			values[i] = num
		} else {
			var str string
			if err := json.Unmarshal(raw, &str); err != nil {
				return nil, fmt.Errorf("invalid int64 at index %d", i)
			}
			val, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid int64 string at index %d: %s", i, str)
			}
			values[i] = val
		}
	}

	return element.NewSignedVeryLong(t, values), nil
}

// readUint16Array reads a uint16 array (US)
func (r *jsonReader) readUint16Array(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var values []uint16
	if err := json.Unmarshal(data, &values); err != nil {
		return nil, err
	}
	return element.NewUnsignedShort(t, values), nil
}

// readUint32Array reads a uint32 array (UL)
func (r *jsonReader) readUint32Array(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var values []uint32
	if err := json.Unmarshal(data, &values); err != nil {
		return nil, err
	}
	return element.NewUnsignedLong(t, values), nil
}

// readUint64Array reads a uint64 array (UV)
func (r *jsonReader) readUint64Array(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, err
	}

	values := make([]uint64, len(rawValues))
	for i, raw := range rawValues {
		// Try as number first, then as string
		var num uint64
		if err := json.Unmarshal(raw, &num); err == nil {
			values[i] = num
		} else {
			var str string
			if err := json.Unmarshal(raw, &str); err != nil {
				return nil, fmt.Errorf("invalid uint64 at index %d", i)
			}
			val, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid uint64 string at index %d: %s", i, str)
			}
			values[i] = val
		}
	}

	return element.NewUnsignedVeryLong(t, values), nil
}

// readAttributeTag reads an AT (Attribute Tag) value
func (r *jsonReader) readAttributeTag(t *tag.Tag, vr *vr.VR, data json.RawMessage) (element.Element, error) {
	var tagStrs []string
	if err := json.Unmarshal(data, &tagStrs); err != nil {
		return nil, err
	}

	values := make([]string, len(tagStrs))
	for i, tagStr := range tagStrs {
		if tagStr == "" {
			values[i] = ""
			continue
		}
		// Parse as hex tag
		if len(tagStr) == 8 {
			// Format as (GGGG,EEEE)
			values[i] = fmt.Sprintf("(%s,%s)", tagStr[0:4], tagStr[4:8])
		} else {
			values[i] = tagStr
		}
	}

	return element.NewString(t, vr, values), nil
}

// readPersonName reads a PN (Person Name) array
func (r *jsonReader) readPersonName(t *tag.Tag, vr *vr.VR, data json.RawMessage) (element.Element, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, err
	}

	values := make([]string, len(rawValues))
	for i, raw := range rawValues {
		// Check if null
		var isNull interface{}
		if err := json.Unmarshal(raw, &isNull); err == nil && isNull == nil {
			values[i] = ""
			continue
		}

		// Parse as PersonName object
		var pnObj map[string]string
		if err := json.Unmarshal(raw, &pnObj); err != nil {
			return nil, fmt.Errorf("invalid PersonName at index %d: %w", i, err)
		}

		// Build PersonName string: Alphabetic=Ideographic=Phonetic
		parts := make([]string, 3)
		if val, ok := pnObj["Alphabetic"]; ok {
			parts[0] = val
		}
		if val, ok := pnObj["Ideographic"]; ok {
			parts[1] = val
		}
		if val, ok := pnObj["Phonetic"]; ok {
			parts[2] = val
		}

		// Trim trailing empty parts
		lastNonEmpty := -1
		for j := len(parts) - 1; j >= 0; j-- {
			if parts[j] != "" {
				lastNonEmpty = j
				break
			}
		}

		if lastNonEmpty >= 0 {
			values[i] = strings.Join(parts[:lastNonEmpty+1], "=")
		} else {
			values[i] = ""
		}
	}

	return element.NewString(t, vr, values), nil
}

// readSequence reads a SQ (Sequence) value
func (r *jsonReader) readSequence(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var rawItems []json.RawMessage
	if err := json.Unmarshal(data, &rawItems); err != nil {
		return nil, err
	}

	seq := dataset.NewSequence(t)
	for i, raw := range rawItems {
		item, err := r.readDataset(raw)
		if err != nil {
			return nil, fmt.Errorf("failed to read sequence item %d: %w", i, err)
		}
		seq.AddItem(item)
	}

	return seq, nil
}

// readInlineBinary reads InlineBinary (Base64-encoded binary data)
func (r *jsonReader) readInlineBinary(t *tag.Tag, vrObj *vr.VR, data json.RawMessage) (element.Element, error) {
	var encoded string
	if err := json.Unmarshal(data, &encoded); err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %w", err)
	}

	buf := buffer.NewMemory(decoded)

	// Create element based on VR
	vrCode := vrObj.Code()
	switch vrCode {
	case vr.CodeOB:
		return element.NewOtherByteFromBuffer(t, buf), nil
	case vr.CodeOW:
		return element.NewOtherWordFromBuffer(t, buf), nil
	case vr.CodeOD:
		return element.NewOtherDoubleFromBuffer(t, buf), nil
	case vr.CodeOF:
		return element.NewOtherFloatFromBuffer(t, buf), nil
	case vr.CodeOL:
		return element.NewOtherLongFromBuffer(t, buf), nil
	case vr.CodeOV:
		return element.NewOtherVeryLongFromBuffer(t, buf), nil
	case vr.CodeUN:
		return element.NewUnknownFromBuffer(t, buf), nil
	default:
		return nil, fmt.Errorf("InlineBinary not supported for VR %s", vrCode)
	}
}

// readBulkDataURI reads a BulkDataURI reference
func (r *jsonReader) readBulkDataURI(t *tag.Tag, vrObj *vr.VR, data json.RawMessage) (element.Element, error) {
	var uri string
	if err := json.Unmarshal(data, &uri); err != nil {
		return nil, fmt.Errorf("invalid BulkDataURI: %w", err)
	}

	// Empty URI is not allowed
	if uri == "" {
		return nil, fmt.Errorf("BulkDataURI cannot be empty")
	}

	// Create a BulkDataUri buffer with the URI
	// The actual data will be loaded later when needed
	buf := buffer.NewBulkDataUri(uri)

	// Create element based on VR
	vrCode := vrObj.Code()
	switch vrCode {
	case vr.CodeOB:
		return element.NewOtherByteFromBuffer(t, buf), nil
	case vr.CodeOW:
		return element.NewOtherWordFromBuffer(t, buf), nil
	case vr.CodeOD:
		return element.NewOtherDoubleFromBuffer(t, buf), nil
	case vr.CodeOF:
		return element.NewOtherFloatFromBuffer(t, buf), nil
	case vr.CodeOL:
		return element.NewOtherLongFromBuffer(t, buf), nil
	case vr.CodeOV:
		return element.NewOtherVeryLongFromBuffer(t, buf), nil
	case vr.CodeUN:
		return element.NewUnknownFromBuffer(t, buf), nil
	default:
		return nil, fmt.Errorf("BulkDataURI not supported for VR %s", vrCode)
	}
}

// createEmptyElement creates an empty element
func (r *jsonReader) createEmptyElement(t *tag.Tag, vrObj *vr.VR) (element.Element, error) {
	switch vrObj.Code() {
	case vr.CodeSQ:
		return dataset.NewSequence(t), nil
	default:
		return element.NewString(t, vrObj, []string{}), nil
	}
}
