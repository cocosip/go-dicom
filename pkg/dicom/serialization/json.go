// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// personNameComponentGroupDelimiter is the delimiter for PersonName component groups (=)
const personNameComponentGroupDelimiter = '='

// personNameComponentGroupNames are the names of PersonName component groups
var personNameComponentGroupNames = []string{"Alphabetic", "Ideographic", "Phonetic"}

// ToJSON serializes a DICOM dataset to JSON format according to DICOM Part 18
func ToJSON(ds *dataset.Dataset, opts ...JSONOption) ([]byte, error) {
	cfg := defaultJSONConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	w := &jsonWriter{
		config: cfg,
		buf:    &bytes.Buffer{},
	}

	if err := w.writeDataset(ds); err != nil {
		return nil, err
	}

	// Apply indentation if configured
	if cfg.indent != "" {
		var indented bytes.Buffer
		if err := json.Indent(&indented, w.buf.Bytes(), "", cfg.indent); err != nil {
			return nil, fmt.Errorf("failed to indent JSON: %w", err)
		}
		return indented.Bytes(), nil
	}

	return w.buf.Bytes(), nil
}

// jsonWriter handles the JSON serialization process
type jsonWriter struct {
	config jsonConfig
	buf    *bytes.Buffer
}

// writeDataset writes a DICOM dataset as JSON
func (w *jsonWriter) writeDataset(ds *dataset.Dataset) error {
	// Start object
	w.buf.WriteString("{")
	first := true

	for _, elem := range ds.Elements() {
		// Skip group length tags (gggg,0000)
		if elem.Tag().Element() == 0 {
			continue
		}

		if !first {
			w.buf.WriteString(",")
		}
		first = false

		// Write tag as key
		if err := w.writeTagKey(elem); err != nil {
			return err
		}

		w.buf.WriteString(":")

		// Write element as value
		if err := w.writeElement(elem); err != nil {
			return err
		}
	}

	w.buf.WriteString("}")
	return nil
}

// writeTagKey writes the tag as a JSON key (either as hex tag or keyword)
func (w *jsonWriter) writeTagKey(elem element.Element) error {
	t := elem.Tag()
	entryIface := t.DictionaryEntry()

	var entry *dict.Entry
	if entryIface != nil {
		entry, _ = entryIface.(*dict.Entry)
	}

	// Check if we can write as keyword
	unknown := entry == nil ||
		entry.Keyword() == "" ||
		(entry.MaskTag() != nil && entry.MaskTag().Mask() != 0xffffffff)

	if w.config.writeTagsAsKeywords && !unknown {
		// Write as keyword
		w.buf.WriteString(`"`)
		w.buf.WriteString(entry.Keyword())
		w.buf.WriteString(`"`)
	} else {
		// Write as hex tag (GGGGEEEE)
		w.buf.WriteString(`"`)
		fmt.Fprintf(w.buf, "%04X%04X", t.Group(), t.Element())
		w.buf.WriteString(`"`)
	}

	return nil
}

// writeElement writes a DICOM element as JSON
func (w *jsonWriter) writeElement(elem element.Element) error {
	w.buf.WriteString("{")

	// Write VR
	w.buf.WriteString(`"vr":"`)
	w.buf.WriteString(elem.ValueRepresentation().Code())
	w.buf.WriteString(`"`)

	// Write value based on VR type
	if err := w.writeValueByVR(elem); err != nil {
		return err
	}

	// Write keyword and name if configured
	w.writeKeywordAndName(elem)

	w.buf.WriteString("}")
	return nil
}

// writeValueByVR writes the element value based on its VR type
func (w *jsonWriter) writeValueByVR(elem element.Element) error {
	vrCode := elem.ValueRepresentation().Code()
	switch vrCode {
	case vr.CodeSQ:
		return w.writeSequence(elem)
	case vr.CodePN:
		return w.writePersonName(elem)
	case vr.CodeOB, vr.CodeOD, vr.CodeOF, vr.CodeOL, vr.CodeOV, vr.CodeOW, vr.CodeUN:
		return w.writeOther(elem)
	case vr.CodeFL:
		return w.writeFloat32Array(elem)
	case vr.CodeFD:
		return w.writeFloat64Array(elem)
	case vr.CodeIS, vr.CodeDS:
		// IS and DS are string-based numeric types
		return w.writeNumericStringArray(elem, vrCode)
	case vr.CodeSL:
		return w.writeInt32Array(elem)
	case vr.CodeSS:
		return w.writeInt16Array(elem)
	case vr.CodeSV:
		return w.writeInt64Array(elem)
	case vr.CodeUL:
		return w.writeUint32Array(elem)
	case vr.CodeUS:
		return w.writeUint16Array(elem)
	case vr.CodeUV:
		return w.writeUint64Array(elem)
	case vr.CodeAT:
		return w.writeAttributeTag(elem)
	default:
		// All other string types
		return w.writeStringArray(elem)
	}
}

// writeKeywordAndName writes keyword and name metadata if configured
func (w *jsonWriter) writeKeywordAndName(elem element.Element) {
	if !w.config.writeKeyword && !w.config.writeName {
		return
	}

	entryIface := elem.Tag().DictionaryEntry()
	var entry *dict.Entry
	if entryIface != nil {
		entry, _ = entryIface.(*dict.Entry)
	}

	unknown := entry == nil ||
		entry.Keyword() == "" ||
		(entry.MaskTag() != nil && entry.MaskTag().Mask() != 0xffffffff)

	if unknown {
		return
	}

	if w.config.writeKeyword {
		w.buf.WriteString(`,"keyword":"`)
		w.buf.WriteString(entry.Keyword())
		w.buf.WriteString(`"`)
	}
	if w.config.writeName {
		w.buf.WriteString(`,"name":"`)
		w.buf.WriteString(entry.Name())
		w.buf.WriteString(`"`)
	}
}

// writeStringArray writes a string array value
func (w *jsonWriter) writeStringArray(elem element.Element) error {
	// Type assert to *element.String
	strElem, ok := elem.(*element.String)
	if !ok {
		// Not a string element, skip value
		return nil
	}

	values := strElem.GetValues()
	if len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		if val == "" {
			w.buf.WriteString("null")
		} else {
			// Escape and write string
			escaped, _ := json.Marshal(val)
			w.buf.Write(escaped)
		}
	}
	w.buf.WriteString("]")
	return nil
}

// writeFloat32Array writes a float32 array as JSON numbers
func (w *jsonWriter) writeFloat32Array(elem element.Element) error {
	// Type assert to *element.Float
	floatElem, ok := elem.(*element.Float)
	if !ok {
		return nil
	}

	values, err := floatElem.GetValues()
	if err != nil || len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		w.writeFloat32(val)
	}
	w.buf.WriteString("]")
	return nil
}

// writeFloat64Array writes a float64 array as JSON numbers
func (w *jsonWriter) writeFloat64Array(elem element.Element) error {
	// Type assert to *element.Double
	doubleElem, ok := elem.(*element.Double)
	if !ok {
		return nil
	}

	values, err := doubleElem.GetValues()
	if err != nil || len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		w.writeFloat64(val)
	}
	w.buf.WriteString("]")
	return nil
}

// writeFloat32 writes a single float32 value
func (w *jsonWriter) writeFloat32(val float32) {
	if math.IsNaN(float64(val)) {
		w.buf.WriteString(`"NaN"`)
	} else if math.IsInf(float64(val), 1) {
		w.buf.WriteString(`"Infinity"`)
	} else if math.IsInf(float64(val), -1) {
		w.buf.WriteString(`"-Infinity"`)
	} else {
		w.buf.WriteString(strconv.FormatFloat(float64(val), 'f', -1, 32))
	}
}

// writeFloat64 writes a single float64 value
func (w *jsonWriter) writeFloat64(val float64) {
	if math.IsNaN(val) {
		w.buf.WriteString(`"NaN"`)
	} else if math.IsInf(val, 1) {
		w.buf.WriteString(`"Infinity"`)
	} else if math.IsInf(val, -1) {
		w.buf.WriteString(`"-Infinity"`)
	} else {
		w.buf.WriteString(strconv.FormatFloat(val, 'f', -1, 64))
	}
}

// writeInt16Array writes an int16 array as JSON numbers
func (w *jsonWriter) writeInt16Array(elem element.Element) error {
	// Type assert to *element.SignedShort
	ssElem, ok := elem.(*element.SignedShort)
	if !ok {
		return nil
	}

	values, err := ssElem.GetValues()
	if err != nil || len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		w.buf.WriteString(strconv.FormatInt(int64(val), 10))
	}
	w.buf.WriteString("]")
	return nil
}

// writeInt32Array writes an int32 array as JSON numbers
func (w *jsonWriter) writeInt32Array(elem element.Element) error {
	// Type assert to *element.SignedLong
	slElem, ok := elem.(*element.SignedLong)
	if !ok {
		return nil
	}

	values, err := slElem.GetValues()
	if err != nil || len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		w.buf.WriteString(strconv.FormatInt(int64(val), 10))
	}
	w.buf.WriteString("]")
	return nil
}

// writeInt64Array writes an int64 array (SV) as JSON numbers or strings
func (w *jsonWriter) writeInt64Array(elem element.Element) error {
	// Type assert to *element.SignedVeryLong
	svElem, ok := elem.(*element.SignedVeryLong)
	if !ok {
		return nil
	}

	values, err := svElem.GetValues()
	if err != nil || len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		w.writeNumberOrString(strconv.FormatInt(val, 10))
	}
	w.buf.WriteString("]")
	return nil
}

// writeUint16Array writes a uint16 array as JSON numbers
func (w *jsonWriter) writeUint16Array(elem element.Element) error {
	// Type assert to *element.UnsignedShort
	usElem, ok := elem.(*element.UnsignedShort)
	if !ok {
		return nil
	}

	values, err := usElem.GetValues()
	if err != nil || len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		w.buf.WriteString(strconv.FormatUint(uint64(val), 10))
	}
	w.buf.WriteString("]")
	return nil
}

// writeUint32Array writes a uint32 array as JSON numbers
func (w *jsonWriter) writeUint32Array(elem element.Element) error {
	// Type assert to *element.UnsignedLong
	ulElem, ok := elem.(*element.UnsignedLong)
	if !ok {
		return nil
	}

	values, err := ulElem.GetValues()
	if err != nil || len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		w.buf.WriteString(strconv.FormatUint(uint64(val), 10))
	}
	w.buf.WriteString("]")
	return nil
}

// writeUint64Array writes a uint64 array (UV) as JSON numbers or strings
func (w *jsonWriter) writeUint64Array(elem element.Element) error {
	// Type assert to *element.UnsignedVeryLong
	uvElem, ok := elem.(*element.UnsignedVeryLong)
	if !ok {
		return nil
	}

	values, err := uvElem.GetValues()
	if err != nil || len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		w.writeNumberOrString(strconv.FormatUint(val, 10))
	}
	w.buf.WriteString("]")
	return nil
}

// writeNumericStringArray writes an IS or DS (numeric string) array
func (w *jsonWriter) writeNumericStringArray(elem element.Element, vrCode string) error {
	// Type assert to *element.String
	strElem, ok := elem.(*element.String)
	if !ok {
		return nil
	}

	values := strElem.GetValues()
	if len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		if val == "" {
			w.buf.WriteString("null")
		} else {
			trimmed := strings.TrimSpace(val)
			if vrCode == vr.CodeDS {
				trimmed = fixDecimalString(trimmed)
			}
			w.writeNumberOrString(trimmed)
		}
	}
	w.buf.WriteString("]")
	return nil
}

// writeNumberOrString writes a value as number or string based on configuration
func (w *jsonWriter) writeNumberOrString(val string) {
	if w.config.numberSerializationMode == AsString {
		// Always write as string
		escaped, _ := json.Marshal(val)
		w.buf.Write(escaped)
		return
	}

	// Try to write as number
	switch w.config.numberSerializationMode {
	case AsNumber:
		// Must parse as number
		w.buf.WriteString(val)
	case PreferablyAsNumber:
		// Try as number, fallback to string
		if isValidNumber(val) {
			w.buf.WriteString(val)
		} else {
			escaped, _ := json.Marshal(val)
			w.buf.Write(escaped)
		}
	}
}

// isValidNumber checks if a string is a valid JSON number
func isValidNumber(s string) bool {
	if s == "" {
		return false
	}
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// fixDecimalString fixes a DICOM DS value to be a valid JSON number
func fixDecimalString(val string) string {
	val = strings.TrimSpace(val)
	val = strings.TrimRight(val, "\x00")

	if val == "" {
		return "0"
	}

	negative := false
	// Strip leading + or - sign
	switch val[0] {
	case '+':
		val = val[1:]
	case '-':
		negative = true
		val = val[1:]
	}

	// Strip leading zeros (except before decimal point)
	if len(val) > 1 && val[0] == '0' && val[1] != '.' {
		i := 0
		for i < len(val)-1 && val[i] == '0' && val[i+1] != '.' {
			i++
		}
		val = val[i:]
	}

	// Re-add negative sign
	if negative {
		val = "-" + val
	}

	return val
}

// writeAttributeTag writes an AT (Attribute Tag) value
func (w *jsonWriter) writeAttributeTag(elem element.Element) error {
	// Type assert to *element.String (AT is stored as string)
	strElem, ok := elem.(*element.String)
	if !ok {
		return nil
	}

	values := strElem.GetValues()
	if len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		if val == "" {
			w.buf.WriteString("null")
		} else {
			// Parse tag and format as 8-character hex
			t, err := tag.Parse(val)
			if err != nil {
				w.buf.WriteString("null")
			} else {
				w.buf.WriteString(`"`)
				fmt.Fprintf(w.buf, "%08X", t.Uint32())
				w.buf.WriteString(`"`)
			}
		}
	}
	w.buf.WriteString("]")
	return nil
}

// writeOther writes binary data (OB, OW, etc.) as InlineBinary (Base64) or BulkDataURI
func (w *jsonWriter) writeOther(elem element.Element) error {
	buf := elem.Buffer()

	// Check if this is a BulkDataUri buffer
	if bulkDataBuf, ok := buf.(*buffer.BulkDataURIByteBuffer); ok {
		// Write as BulkDataURI
		w.buf.WriteString(`,"BulkDataURI":"`)
		w.buf.WriteString(bulkDataBuf.BulkDataURI())
		w.buf.WriteString(`"`)
		return nil
	}

	// Otherwise, write as InlineBinary (Base64)
	data := buf.Data()
	if len(data) == 0 {
		return nil
	}

	w.buf.WriteString(`,"InlineBinary":"`)
	encoded := base64.StdEncoding.EncodeToString(data)
	w.buf.WriteString(encoded)
	w.buf.WriteString(`"`)

	return nil
}

// writePersonName writes a PN (Person Name) value
func (w *jsonWriter) writePersonName(elem element.Element) error {
	// Type assert to *element.String
	strElem, ok := elem.(*element.String)
	if !ok {
		return nil
	}

	values := strElem.GetValues()
	if len(values) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, val := range values {
		if i > 0 {
			w.buf.WriteString(",")
		}
		if val == "" {
			w.buf.WriteString("null")
		} else {
			w.writePersonNameValue(val)
		}
	}
	w.buf.WriteString("]")
	return nil
}

// writePersonNameValue writes a single PersonName value as an object
func (w *jsonWriter) writePersonNameValue(val string) {
	// Split by '=' into component groups: Alphabetic, Ideographic, Phonetic
	groups := strings.Split(val, string(personNameComponentGroupDelimiter))

	w.buf.WriteString("{")
	first := true
	for i, group := range groups {
		if i >= len(personNameComponentGroupNames) {
			break
		}
		// Skip empty groups
		if strings.TrimSpace(group) == "" {
			continue
		}

		if !first {
			w.buf.WriteString(",")
		}
		first = false

		w.buf.WriteString(`"`)
		w.buf.WriteString(personNameComponentGroupNames[i])
		w.buf.WriteString(`":`)
		escaped, _ := json.Marshal(group)
		w.buf.Write(escaped)
	}
	w.buf.WriteString("}")
}

// writeSequence writes a SQ (Sequence) value
func (w *jsonWriter) writeSequence(elem element.Element) error {
	seq, ok := elem.(*dataset.Sequence)
	if !ok {
		return nil
	}

	items := seq.GetItems()
	if len(items) == 0 {
		return nil
	}

	w.buf.WriteString(`,"Value":[`)
	for i, item := range items {
		if i > 0 {
			w.buf.WriteString(",")
		}
		if err := w.writeDataset(item); err != nil {
			return err
		}
	}
	w.buf.WriteString("]")
	return nil
}
