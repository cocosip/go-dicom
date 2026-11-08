// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package serialization

import (
	"bytes"
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
	ds.SetAutoValidate(false)
	defer ds.SetAutoValidate(true)

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
			if err := ds.Add(elem); err != nil {
				return nil, fmt.Errorf("failed to add element %s: %w", tagStr, err)
			}
		}
	}

	if r.config.autoValidate {
		if err := ds.Validate(); err != nil {
			return nil, fmt.Errorf("DICOM JSON validation failed: %w", err)
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
	t := r.config.lookup.LookupKeyword(s)
	if t == nil {
		return tag.Tag{}, fmt.Errorf("invalid tag or keyword: %s", s)
	}
	return *t, nil
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
		if isJSONNull(raw) {
			values[i] = ""
			continue
		}
		var str string
		if err := json.Unmarshal(raw, &str); err != nil {
			return nil, fmt.Errorf("value %d: expected a string or null for VR %s: %w", i, vr.Code(), err)
		}
		values[i] = str
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
		if isJSONNull(raw) {
			values[i] = ""
			continue
		}
		// Try as number first
		var num json.Number
		if err := json.Unmarshal(raw, &num); err == nil {
			if vrObj.Code() == vr.CodeIS {
				values[i] = num.String()
			} else {
				if f, err := num.Float64(); err == nil {
					values[i] = strconv.FormatFloat(f, 'f', -1, 64)
				} else {
					values[i] = num.String()
				}
			}
			continue
		}
		// Try as string
		var str string
		if err := json.Unmarshal(raw, &str); err != nil {
			return nil, fmt.Errorf("value %d: expected a number, string, or null for VR %s: %w", i, vrObj.Code(), err)
		}
		values[i] = str
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

// parseInt64FromJSON parses an int64 from JSON, handling both number and string representations.
func parseInt64FromJSON(raw json.RawMessage, index int) (int64, error) {
	// Try as number first, then as string
	var num int64
	if err := json.Unmarshal(raw, &num); err == nil {
		return num, nil
	}

	var str string
	if err := json.Unmarshal(raw, &str); err != nil {
		return 0, fmt.Errorf("invalid int64 at index %d", index)
	}
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid int64 string at index %d: %s", index, str)
	}
	return val, nil
}

// parseUint64FromJSON parses a uint64 from JSON, handling both number and string representations.
func parseUint64FromJSON(raw json.RawMessage, index int) (uint64, error) {
	// Try as number first, then as string
	var num uint64
	if err := json.Unmarshal(raw, &num); err == nil {
		return num, nil
	}

	var str string
	if err := json.Unmarshal(raw, &str); err != nil {
		return 0, fmt.Errorf("invalid uint64 at index %d", index)
	}
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid uint64 string at index %d: %s", index, str)
	}
	return val, nil
}

// readInt64Array reads an int64 array (SV)
func (r *jsonReader) readInt64Array(t *tag.Tag, data json.RawMessage) (element.Element, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, err
	}

	values := make([]int64, len(rawValues))
	for i, raw := range rawValues {
		val, err := parseInt64FromJSON(raw, i)
		if err != nil {
			return nil, err
		}
		values[i] = val
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
		val, err := parseUint64FromJSON(raw, i)
		if err != nil {
			return nil, err
		}
		values[i] = val
	}

	return element.NewUnsignedVeryLong(t, values), nil
}

// readAttributeTag reads an AT (Attribute Tag) value.
//
// PS3.18 F.2.3: AT values are JSON strings of exactly eight uppercase
// hexadecimal digits (GGGGEEEE).
func (r *jsonReader) readAttributeTag(t *tag.Tag, _ *vr.VR, data json.RawMessage) (element.Element, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, err
	}

	values := make([]*tag.Tag, 0, len(rawValues))
	for i, raw := range rawValues {
		if isJSONNull(raw) {
			// A JSON null represents an empty value (PS3.18 F.2.2); skip it.
			continue
		}
		var tagStr string
		if err := json.Unmarshal(raw, &tagStr); err != nil {
			return nil, fmt.Errorf("value %d: expected a string for VR AT: %w", i, err)
		}
		if !isUppercaseHex8(tagStr) {
			return nil, fmt.Errorf("value %d: AT must be eight uppercase hexadecimal digits, got %q", i, tagStr)
		}
		parsed, err := tag.Parse(tagStr)
		if err != nil {
			return nil, fmt.Errorf("value %d: invalid AT %q: %w", i, tagStr, err)
		}
		values = append(values, parsed)
	}

	return element.NewAttributeTag(t, values), nil
}

// isJSONNull reports whether a raw JSON value is the null literal.
func isJSONNull(raw json.RawMessage) bool {
	return bytes.Equal(bytes.TrimSpace(raw), []byte("null"))
}

// isUppercaseHex8 reports whether s is exactly eight uppercase hexadecimal
// digits, the only form PS3.18 F.2.3 permits for AT values.
func isUppercaseHex8(s string) bool {
	if len(s) != 8 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < '0' || c > '9') && (c < 'A' || c > 'F') {
			return false
		}
	}
	return true
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

// bulkDataURIVRs lists the VRs for which PS3.18 F.2.2 permits BulkDataURI.
func bulkDataURIVRs(code string) bool {
	switch code {
	case vr.CodeDS, vr.CodeFL, vr.CodeFD, vr.CodeIS, vr.CodeLT,
		vr.CodeOB, vr.CodeOD, vr.CodeOF, vr.CodeOL, vr.CodeOV, vr.CodeOW,
		vr.CodeSL, vr.CodeSS, vr.CodeST, vr.CodeSV,
		vr.CodeUC, vr.CodeUL, vr.CodeUN, vr.CodeUS, vr.CodeUT, vr.CodeUV:
		return true
	default:
		return false
	}
}

// readBulkDataURI reads a BulkDataURI reference.
//
// The reference is stored in a lazy buffer and the element keeps its VR, so
// string and numeric VRs allowed by PS3.18 F.2.2 are preserved as well.
func (r *jsonReader) readBulkDataURI(t *tag.Tag, vrObj *vr.VR, data json.RawMessage) (element.Element, error) {
	var uri string
	if err := json.Unmarshal(data, &uri); err != nil {
		return nil, fmt.Errorf("invalid BulkDataURI: %w", err)
	}

	// Empty URI is not allowed
	if uri == "" {
		return nil, fmt.Errorf("BulkDataURI cannot be empty")
	}

	if !bulkDataURIVRs(vrObj.Code()) {
		return nil, fmt.Errorf("BulkDataURI not supported for VR %s", vrObj.Code())
	}

	buf := buffer.NewBulkDataURI(uri)
	return element.NewElementFromBuffer(t, vrObj, buf, nil)
}

// createEmptyElement creates an empty element with the concrete type that
// matches its VR (for example *element.AttributeTag for AT and
// *element.UnsignedShort for US).
func (r *jsonReader) createEmptyElement(t *tag.Tag, vrObj *vr.VR) (element.Element, error) {
	if vrObj.Code() == vr.CodeSQ {
		return dataset.NewSequence(t), nil
	}
	return element.NewElementFromBuffer(t, vrObj, buffer.Empty, nil)
}
