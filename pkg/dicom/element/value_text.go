// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
)

// UnsupportedValueError reports an element whose value cannot be represented
// losslessly as canonical strings.
type UnsupportedValueError struct {
	Tag  *tag.Tag
	VR   *vr.VR
	Type string
}

// CanonicalValueContext supplies Dataset-level encoding and byte-order state
// when canonical values create an element that has no existing prototype.
type CanonicalValueContext struct {
	TextEncodings []encoding.Encoding
	Endian        endian.Endian
}

func (e *UnsupportedValueError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DICOM element %s with VR %s and type %s has no canonical string representation",
		formatValueTag(e.Tag), formatValueVR(e.VR), e.Type)
}

// CanonicalStrings returns individual DICOM values in a deterministic text
// representation. It rejects containers and bulk binary values.
func CanonicalStrings(elem Element) ([]string, error) {
	if elem == nil {
		return nil, unsupportedValue(elem)
	}
	if elem.ValueRepresentation() == nil {
		return nil, unsupportedValue(elem)
	}
	switch elem.ValueRepresentation().Code() {
	case vr.CodeSQ, vr.CodeOB, vr.CodeOW, vr.CodeOL, vr.CodeOD, vr.CodeOF, vr.CodeOV, vr.CodeUN:
		return nil, unsupportedValue(elem)
	}

	if str := underlyingString(elem); str != nil {
		return canonicalStringValues(str), nil
	}

	switch value := elem.(type) {
	case *UnsignedShort:
		values, err := value.GetValues()
		return formatUnsigned(values), err
	case *UnsignedLong:
		values, err := value.GetValues()
		return formatUnsigned(values), err
	case *SignedShort:
		values, err := value.GetValues()
		return formatSigned(values), err
	case *SignedLong:
		values, err := value.GetValues()
		return formatSigned(values), err
	case *SignedVeryLong:
		values, err := value.GetValues()
		return formatSigned(values), err
	case *UnsignedVeryLong:
		values, err := value.GetValues()
		return formatUnsigned(values), err
	case *Float:
		values, err := value.GetValues()
		if err != nil {
			return nil, err
		}
		result := make([]string, len(values))
		for index, number := range values {
			result[index] = strconv.FormatFloat(float64(number), 'g', -1, 32)
		}
		return result, nil
	case *Double:
		values, err := value.GetValues()
		if err != nil {
			return nil, err
		}
		result := make([]string, len(values))
		for index, number := range values {
			result[index] = strconv.FormatFloat(number, 'g', -1, 64)
		}
		return result, nil
	case *AttributeTag:
		values, err := value.GetValues()
		if err != nil {
			return nil, err
		}
		result := make([]string, len(values))
		for index, attributeTag := range values {
			if attributeTag == nil {
				return nil, fmt.Errorf("attribute tag value %d is nil", index)
			}
			result[index] = attributeTag.String()
		}
		return result, nil
	default:
		return nil, unsupportedValue(elem)
	}
}

// ReplaceCanonicalStrings constructs a VR-correct element from individual
// canonical values. A compatible prototype preserves text encodings and
// numeric byte order.
func ReplaceCanonicalStrings(prototype Element, targetTag *tag.Tag, targetVR *vr.VR, values []string) (Element, error) {
	return replaceCanonicalStrings(targetTag, targetVR, values, prototypeEncodings(prototype), prototypeEndian(prototype))
}

// ReplaceCanonicalStringsWithContext constructs a VR-correct element using
// explicit Dataset-level text encoding and numeric byte order.
func ReplaceCanonicalStringsWithContext(
	targetTag *tag.Tag,
	targetVR *vr.VR,
	values []string,
	context CanonicalValueContext,
) (Element, error) {
	encodings := append([]encoding.Encoding(nil), context.TextEncodings...)
	if len(encodings) == 0 {
		encodings = []encoding.Encoding{charset.Default}
	}
	return replaceCanonicalStrings(targetTag, targetVR, values, encodings, context.Endian)
}

func replaceCanonicalStrings(
	targetTag *tag.Tag,
	targetVR *vr.VR,
	values []string,
	encodings []encoding.Encoding,
	byteOrder endian.Endian,
) (Element, error) {
	if targetTag == nil {
		return nil, fmt.Errorf("canonical value target tag is nil")
	}
	if targetVR == nil {
		return nil, fmt.Errorf("canonical value target VR is nil")
	}
	values = append([]string(nil), values...)

	var result Element
	var err error
	switch targetVR.Code() {
	case vr.CodeAE, vr.CodeAS, vr.CodeCS, vr.CodeLO, vr.CodeLT, vr.CodeSH,
		vr.CodeST, vr.CodeUC, vr.CodeUI, vr.CodeUR, vr.CodeUT:
		result, err = newCanonicalString(targetTag, targetVR, values, encodings)
	case vr.CodeDA:
		var str *String
		str, err = newCanonicalString(targetTag, targetVR, values, encodings)
		result = &Date{str: str}
	case vr.CodeTM:
		var str *String
		str, err = newCanonicalString(targetTag, targetVR, values, encodings)
		result = &Time{str: str}
	case vr.CodeDT:
		var str *String
		str, err = newCanonicalString(targetTag, targetVR, values, encodings)
		result = &DateTime{str: str}
	case vr.CodePN:
		var str *String
		str, err = newCanonicalString(targetTag, targetVR, values, encodings)
		result = &PersonName{str: str}
	case vr.CodeDS:
		var str *String
		str, err = newCanonicalString(targetTag, targetVR, values, encodings)
		result = &DecimalString{str: str}
	case vr.CodeIS:
		var str *String
		str, err = newCanonicalString(targetTag, targetVR, values, encodings)
		result = &IntegerString{str: str}
	case vr.CodeUS:
		var parsed []uint16
		parsed, err = parseUnsigned[uint16](values, 16)
		result = NewUnsignedShortWithEndian(targetTag, parsed, byteOrder)
	case vr.CodeUL:
		var parsed []uint32
		parsed, err = parseUnsigned[uint32](values, 32)
		result = NewUnsignedLongWithEndian(targetTag, parsed, byteOrder)
	case vr.CodeUV:
		var parsed []uint64
		parsed, err = parseUnsigned[uint64](values, 64)
		result = NewUnsignedVeryLongWithEndian(targetTag, parsed, byteOrder)
	case vr.CodeSS:
		var parsed []int16
		parsed, err = parseSigned[int16](values, 16)
		result = NewSignedShortWithEndian(targetTag, parsed, byteOrder)
	case vr.CodeSL:
		var parsed []int32
		parsed, err = parseSigned[int32](values, 32)
		result = NewSignedLongWithEndian(targetTag, parsed, byteOrder)
	case vr.CodeSV:
		var parsed []int64
		parsed, err = parseSigned[int64](values, 64)
		result = NewSignedVeryLongWithEndian(targetTag, parsed, byteOrder)
	case vr.CodeFL:
		var parsed []float32
		parsed, err = parseFloat32s(values)
		result = NewFloatWithEndian(targetTag, parsed, byteOrder)
	case vr.CodeFD:
		var parsed []float64
		parsed, err = parseFloat64s(values)
		result = NewDoubleWithEndian(targetTag, parsed, byteOrder)
	case vr.CodeAT:
		var parsed []*tag.Tag
		parsed, err = parseAttributeTags(values)
		result = NewAttributeTagWithEndian(targetTag, parsed, byteOrder)
	default:
		return nil, &UnsupportedValueError{Tag: targetTag, VR: targetVR, Type: "replacement"}
	}
	if err != nil {
		return nil, fmt.Errorf("convert canonical values for %s VR %s: %w", targetTag, targetVR.Code(), err)
	}
	if err := ValidateValue(result); err != nil {
		return nil, fmt.Errorf("validate canonical values for %s VR %s: %w", targetTag, targetVR.Code(), err)
	}
	return result, nil
}

func underlyingString(elem Element) *String {
	switch value := elem.(type) {
	case *String:
		return value
	case *DecimalString:
		return value.str
	case *IntegerString:
		return value.str
	case *Date:
		return value.str
	case *Time:
		return value.str
	case *DateTime:
		return value.str
	case *PersonName:
		return value.str
	default:
		return nil
	}
}

func canonicalStringValues(value *String) []string {
	complete := value.GetString()
	if complete == "" {
		return nil
	}
	return strings.Split(complete, "\\")
}

func newCanonicalString(targetTag *tag.Tag, targetVR *vr.VR, values []string, encodings []encoding.Encoding) (*String, error) {
	joined := strings.Join(values, "\\")
	data, err := charset.EncodeString(joined, encodings)
	if err != nil {
		return nil, err
	}
	if targetVR.Code() == vr.CodeUI && len(data)%2 != 0 {
		data = append(data, 0)
	}
	return NewStringFromBufferWithEncodings(targetTag, targetVR, buffer.NewMemory(data), encodings), nil
}

func prototypeEncodings(prototype Element) []encoding.Encoding {
	str := underlyingString(prototype)
	if str == nil {
		return []encoding.Encoding{charset.Default}
	}
	return append([]encoding.Encoding(nil), str.encodings...)
}

func prototypeEndian(prototype Element) endian.Endian {
	base := underlyingBase(prototype)
	if base != nil && base.getByteOrder() == binary.BigEndian {
		return endian.Big
	}
	return endian.Little
}

func underlyingBase(elem Element) *base {
	switch value := elem.(type) {
	case *AttributeTag:
		return value.base
	case *UnsignedShort:
		return value.base
	case *UnsignedLong:
		return value.base
	case *SignedShort:
		return value.base
	case *SignedLong:
		return value.base
	case *Float:
		return value.base
	case *Double:
		return value.base
	case *SignedVeryLong:
		return value.base
	case *UnsignedVeryLong:
		return value.base
	default:
		return nil
	}
}

func parseUnsigned[T ~uint16 | ~uint32 | ~uint64](values []string, bits int) ([]T, error) {
	result := make([]T, len(values))
	for index, value := range values {
		parsed, err := strconv.ParseUint(value, 10, bits)
		if err != nil {
			return nil, fmt.Errorf("value[%d] %q: %w", index, value, err)
		}
		result[index] = T(parsed)
	}
	return result, nil
}

func parseSigned[T ~int16 | ~int32 | ~int64](values []string, bits int) ([]T, error) {
	result := make([]T, len(values))
	for index, value := range values {
		parsed, err := strconv.ParseInt(value, 10, bits)
		if err != nil {
			return nil, fmt.Errorf("value[%d] %q: %w", index, value, err)
		}
		result[index] = T(parsed)
	}
	return result, nil
}

func parseFloat32s(values []string) ([]float32, error) {
	result := make([]float32, len(values))
	for index, value := range values {
		parsed, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return nil, fmt.Errorf("value[%d] %q: %w", index, value, err)
		}
		result[index] = float32(parsed)
	}
	return result, nil
}

func parseFloat64s(values []string) ([]float64, error) {
	result := make([]float64, len(values))
	for index, value := range values {
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("value[%d] %q: %w", index, value, err)
		}
		result[index] = parsed
	}
	return result, nil
}

func parseAttributeTags(values []string) ([]*tag.Tag, error) {
	result := make([]*tag.Tag, len(values))
	for index, value := range values {
		parsed, err := tag.Parse(value)
		if err != nil {
			return nil, fmt.Errorf("value[%d] %q: %w", index, value, err)
		}
		result[index] = parsed
	}
	return result, nil
}

func formatUnsigned[T ~uint16 | ~uint32 | ~uint64](values []T) []string {
	result := make([]string, len(values))
	for index, number := range values {
		result[index] = strconv.FormatUint(uint64(number), 10)
	}
	return result
}

func formatSigned[T ~int16 | ~int32 | ~int64](values []T) []string {
	result := make([]string, len(values))
	for index, number := range values {
		result[index] = strconv.FormatInt(int64(number), 10)
	}
	return result
}

func unsupportedValue(elem Element) error {
	if elem == nil {
		return &UnsupportedValueError{Type: "<nil>"}
	}
	return &UnsupportedValueError{
		Tag:  elem.Tag(),
		VR:   elem.ValueRepresentation(),
		Type: fmt.Sprintf("%T", elem),
	}
}

func formatValueTag(value *tag.Tag) string {
	if value == nil {
		return "<nil-tag>"
	}
	return value.String()
}

func formatValueVR(value *vr.VR) string {
	if value == nil {
		return "<nil-vr>"
	}
	return value.Code()
}
