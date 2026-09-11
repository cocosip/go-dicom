// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"github.com/cocosip/go-dicom/pkg/io/endian"
	"golang.org/x/text/encoding"
)

// NewElementFromBuffer creates the concrete Element implementation for a VR.
// Text elements retain the complete Specific Character Set encoding list.
func NewElementFromBuffer(t *tag.Tag, valueRepresentation *vr.VR, buf buffer.ByteBuffer, encodings []encoding.Encoding) (Element, error) {
	if t == nil {
		return nil, fmt.Errorf("element tag is nil")
	}
	if valueRepresentation == nil {
		return nil, fmt.Errorf("element VR is nil")
	}

	text := func() Element {
		return NewStringFromBufferWithEncodings(t, valueRepresentation, buf, encodings)
	}
	ordered := func(elem Element) Element {
		return elem
	}

	switch valueRepresentation.Code() {
	case vr.CodeAE, vr.CodeAS, vr.CodeCS, vr.CodeLO, vr.CodeLT, vr.CodeSH,
		vr.CodeST, vr.CodeUC, vr.CodeUI, vr.CodeUR, vr.CodeUT:
		return ordered(text()), nil
	case vr.CodeDA:
		return NewDateFromBufferWithEncodings(t, buf, encodings), nil
	case vr.CodeDS:
		return NewDecimalStringFromBufferWithEncodings(t, buf, encodings), nil
	case vr.CodeDT:
		return NewDateTimeFromBufferWithEncodings(t, buf, encodings), nil
	case vr.CodeIS:
		return NewIntegerStringFromBufferWithEncodings(t, buf, encodings), nil
	case vr.CodePN:
		return NewPersonNameFromBufferWithEncodings(t, buf, encodings), nil
	case vr.CodeTM:
		return NewTimeFromBufferWithEncodings(t, buf, encodings), nil
	case vr.CodeUS:
		return NewUnsignedShortFromBuffer(t, buf), nil
	case vr.CodeUL:
		return NewUnsignedLongFromBuffer(t, buf), nil
	case vr.CodeSS:
		return NewSignedShortFromBuffer(t, buf), nil
	case vr.CodeSL:
		return NewSignedLongFromBuffer(t, buf), nil
	case vr.CodeFL:
		return NewFloatFromBuffer(t, buf), nil
	case vr.CodeFD:
		return NewDoubleFromBuffer(t, buf), nil
	case vr.CodeSV:
		return NewSignedVeryLongFromBuffer(t, buf), nil
	case vr.CodeUV:
		return NewUnsignedVeryLongFromBuffer(t, buf), nil
	case vr.CodeOB:
		return NewOtherByteFromBuffer(t, buf), nil
	case vr.CodeOW:
		return NewOtherWordFromBuffer(t, buf), nil
	case vr.CodeOD:
		return NewOtherDoubleFromBuffer(t, buf), nil
	case vr.CodeOF:
		return NewOtherFloatFromBuffer(t, buf), nil
	case vr.CodeOL:
		return NewOtherLongFromBuffer(t, buf), nil
	case vr.CodeOV:
		return NewOtherVeryLongFromBuffer(t, buf), nil
	case vr.CodeAT:
		return NewAttributeTagFromBuffer(t, buf), nil
	case vr.CodeUN:
		return NewUnknownFromBuffer(t, buf), nil
	default:
		return NewUnknownFromBuffer(t, buf), nil
	}
}

// NewElementFromValue creates a concrete Element from a scalar, slice, or
// array Go value. The supplied VR is authoritative; values are never used to
// guess or silently change the element's VR.
func NewElementFromValue(t *tag.Tag, valueRepresentation *vr.VR, value any) (Element, error) {
	return NewElementFromValueWithContext(t, valueRepresentation, value, CanonicalValueContext{
		TextEncodings: []encoding.Encoding{charset.Default},
		Endian:        endian.Native(),
	})
}

// NewElementFromValueWithContext creates a concrete Element using the supplied
// text encodings and byte order. The supplied VR remains authoritative.
func NewElementFromValueWithContext(t *tag.Tag, valueRepresentation *vr.VR, value any, context CanonicalValueContext) (Element, error) {
	if isRawBinaryVR(valueRepresentation) {
		if data, ok := value.([]byte); ok {
			return newRawBinaryElement(t, valueRepresentation, data)
		}
	}
	values, err := canonicalValueStrings(valueRepresentation, value)
	if err != nil {
		return nil, err
	}
	return ReplaceCanonicalStringsWithContext(t, valueRepresentation, values, context)
}

func isRawBinaryVR(valueRepresentation *vr.VR) bool {
	if valueRepresentation == nil {
		return false
	}
	switch valueRepresentation.Code() {
	case vr.CodeOB, vr.CodeOW, vr.CodeOD, vr.CodeOF, vr.CodeOL, vr.CodeOV, vr.CodeUN:
		return true
	default:
		return false
	}
}

func newRawBinaryElement(t *tag.Tag, valueRepresentation *vr.VR, data []byte) (Element, error) {
	buf := buffer.NewMemory(data)
	switch valueRepresentation.Code() {
	case vr.CodeOB:
		return NewOtherByteFromBuffer(t, buf), nil
	case vr.CodeOW:
		return NewOtherWordFromBuffer(t, buf), nil
	case vr.CodeOD:
		return NewOtherDoubleFromBuffer(t, buf), nil
	case vr.CodeOF:
		return NewOtherFloatFromBuffer(t, buf), nil
	case vr.CodeOL:
		return NewOtherLongFromBuffer(t, buf), nil
	case vr.CodeOV:
		return NewOtherVeryLongFromBuffer(t, buf), nil
	case vr.CodeUN:
		return NewUnknownFromBuffer(t, buf), nil
	default:
		return nil, fmt.Errorf("VR %s is not a raw binary VR", valueRepresentation.Code())
	}
}

func canonicalValueStrings(valueRepresentation *vr.VR, value any) ([]string, error) {
	if valueRepresentation == nil {
		return nil, fmt.Errorf("element VR is nil")
	}
	if value == nil {
		return nil, nil
	}
	items := valueItems(value)
	values := make([]string, len(items))
	for index, item := range items {
		var err error
		values[index], err = canonicalValue(valueRepresentation, item)
		if err != nil {
			return nil, fmt.Errorf("value[%d]: %w", index, err)
		}
	}
	return values, nil
}

func valueItems(value any) []any {
	rv := reflect.ValueOf(value)
	for rv.IsValid() && (rv.Kind() == reflect.Interface || rv.Kind() == reflect.Pointer) {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return []any{rv.Interface()}
	}
	items := make([]any, rv.Len())
	for index := 0; index < rv.Len(); index++ {
		items[index] = rv.Index(index).Interface()
	}
	return items
}

func canonicalValue(valueRepresentation *vr.VR, value any) (string, error) {
	if value == nil {
		return "", nil
	}
	if tm, ok := value.(time.Time); ok {
		switch valueRepresentation.Code() {
		case vr.CodeDA:
			return tm.Format("20060102"), nil
		case vr.CodeTM:
			return tm.Format("150405.000000"), nil
		case vr.CodeDT:
			return tm.Format("20060102150405.000000-0700"), nil
		default:
			return "", fmt.Errorf("time.Time is not valid for VR %s", valueRepresentation.Code())
		}
	}
	if t, ok := value.(*tag.Tag); ok {
		if valueRepresentation.Code() != vr.CodeAT {
			return "", fmt.Errorf("*tag.Tag is only valid for VR AT")
		}
		if t == nil {
			return "", fmt.Errorf("tag value is nil")
		}
		return t.String(), nil
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.String:
		if isBinaryVR(valueRepresentation) {
			return "", fmt.Errorf("string values are not valid for binary VR %s", valueRepresentation.Code())
		}
		return rv.String(), nil
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if valueRepresentation.Code() == vr.CodeDS {
			return strconv.FormatInt(rv.Int(), 10), nil
		}
		return strconv.FormatInt(rv.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'f', -1, rv.Type().Bits()), nil
	default:
		if stringer, ok := value.(fmt.Stringer); ok && !isBinaryVR(valueRepresentation) {
			return stringer.String(), nil
		}
		return "", fmt.Errorf("unsupported Go value type %T for VR %s", value, valueRepresentation.Code())
	}
}

func isBinaryVR(valueRepresentation *vr.VR) bool {
	return isRawBinaryVR(valueRepresentation) || valueRepresentation.Code() == vr.CodeUS ||
		valueRepresentation.Code() == vr.CodeUL || valueRepresentation.Code() == vr.CodeSS ||
		valueRepresentation.Code() == vr.CodeSL || valueRepresentation.Code() == vr.CodeSV ||
		valueRepresentation.Code() == vr.CodeUV || valueRepresentation.Code() == vr.CodeFL ||
		valueRepresentation.Code() == vr.CodeFD || valueRepresentation.Code() == vr.CodeAT
}
