// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"
	"reflect"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// AddValue resolves the dictionary VR for t, creates one element containing
// the supplied scalar, slice, or array values, and adds it to the Dataset.
// Values are never split into separate elements.
func (ds *Dataset) AddValue(t *tag.Tag, value any) error {
	valueRepresentation, err := resolveValueVR(t, value)
	if err != nil {
		return err
	}
	return ds.AddValueWithVR(t, valueRepresentation, value)
}

// AddOrUpdateValue resolves the dictionary VR and adds or replaces one value
// element in the Dataset.
func (ds *Dataset) AddOrUpdateValue(t *tag.Tag, value any) error {
	valueRepresentation, err := resolveValueVR(t, value)
	if err != nil {
		return err
	}
	return ds.AddOrUpdateValueWithVR(t, valueRepresentation, value)
}

// AddValueWithVR creates one element using an explicitly selected VR and adds
// it to the Dataset. This is required for private tags and ambiguous
// dictionary entries with multiple legal VRs.
func (ds *Dataset) AddValueWithVR(t *tag.Tag, valueRepresentation *vr.VR, value any) error {
	if ds == nil {
		return fmt.Errorf("cannot add to nil Dataset")
	}
	if t == nil || valueRepresentation == nil {
		return validateTagVR(t, valueRepresentation)
	}
	if ds.validationEnabled() {
		if err := validateTagVR(t, valueRepresentation); err != nil {
			return err
		}
	}
	elem, err := element.NewElementFromValueWithContext(t, valueRepresentation, value, ds.valueContext())
	if err != nil {
		return fmt.Errorf("create element %s with VR %s: %w", t, valueRepresentation.Code(), err)
	}
	return ds.Add(elem)
}

// AddOrUpdateValueWithVR creates one element using an explicitly selected VR
// and adds or replaces it in the Dataset.
func (ds *Dataset) AddOrUpdateValueWithVR(t *tag.Tag, valueRepresentation *vr.VR, value any) error {
	if ds == nil {
		return fmt.Errorf("cannot add to nil Dataset")
	}
	if t == nil || valueRepresentation == nil {
		return validateTagVR(t, valueRepresentation)
	}
	if ds.validationEnabled() {
		if err := validateTagVR(t, valueRepresentation); err != nil {
			return err
		}
	}
	elem, err := element.NewElementFromValueWithContext(t, valueRepresentation, value, ds.valueContext())
	if err != nil {
		return fmt.Errorf("create element %s with VR %s: %w", t, valueRepresentation.Code(), err)
	}
	return ds.AddOrUpdate(elem)
}

func (ds *Dataset) valueContext() element.CanonicalValueContext {
	context := element.CanonicalValueContext{
		TextEncodings: ds.effectiveTextEncodings(),
	}
	if syntax := ds.InternalTransferSyntax(); syntax != nil {
		context.Endian = syntax.Endian()
	}
	return context
}

// AddElements adds a collection of elements, preserving each element's
// existing VR and rejecting the whole operation on the first error.
func (ds *Dataset) AddElements(elements ...element.Element) error {
	for _, elem := range elements {
		if err := ds.Add(elem); err != nil {
			return err
		}
	}
	return nil
}

// AddOrUpdateElements adds or replaces a collection of elements.
func (ds *Dataset) AddOrUpdateElements(elements ...element.Element) error {
	for _, elem := range elements {
		if err := ds.AddOrUpdate(elem); err != nil {
			return err
		}
	}
	return nil
}

func resolveValueVR(t *tag.Tag, value any) (*vr.VR, error) {
	if t == nil {
		return nil, fmt.Errorf("cannot resolve VR for nil tag")
	}
	entry := dict.Default().Lookup(t)
	if entry == nil || len(entry.ValueRepresentations()) == 0 {
		return nil, fmt.Errorf("no dictionary VR for tag %s; use AddValueWithVR", t)
	}
	allowed := entry.ValueRepresentations()
	matching := make([]*vr.VR, 0, len(allowed))
	for _, candidate := range allowed {
		if valueMatchesVR(candidate, value) {
			matching = append(matching, candidate)
		}
	}
	if len(matching) == 1 {
		return matching[0], nil
	}
	if len(matching) == 0 {
		return nil, fmt.Errorf("go value type %T is not valid for any VR of tag %s (%v); use AddValueWithVR", value, t, entry.VRs())
	}
	if t.ToUint32() == tag.PixelData.ToUint32() {
		return nil, fmt.Errorf(
			"pixel data VR cannot be inferred from Go value type %T; use pixeldata.NewForDataset or AddValueWithVR",
			value,
		)
	}
	return nil, fmt.Errorf("tag %s has ambiguous VR for Go value type %T (%v); use AddValueWithVR", t, value, entry.VRs())
}

func validateTagVR(t *tag.Tag, valueRepresentation *vr.VR) error {
	if t == nil {
		return validationError(ValidationStructural, nil, fmt.Errorf("element tag is nil"))
	}
	if valueRepresentation == nil {
		return validationError(ValidationStructural, nil, fmt.Errorf("element VR is nil"))
	}
	entry := dict.Default().Lookup(t)
	if entry == nil || t.IsPrivate() {
		return nil
	}
	for _, allowed := range entry.ValueRepresentations() {
		if allowed.Code() == valueRepresentation.Code() {
			return nil
		}
	}
	return validationError(ValidationVR, nil, fmt.Errorf("tag %s does not allow VR %s (allowed %v)", t, valueRepresentation.Code(), entry.VRs()))
}

func valueMatchesVR(valueRepresentation *vr.VR, value any) bool {
	if valueRepresentation == nil {
		return false
	}
	code := valueRepresentation.Code()
	switch value.(type) {
	case nil:
		return true
	case string:
		return valueRepresentation.IsString()
	case bool:
		return valueRepresentation.IsString()
	case int, int8, int16, int32, int64:
		return signedKindMatchesVR(code)
	case uint, uint8, uint16, uint32, uint64:
		return unsignedKindMatchesVR(code)
	case float32:
		return float32MatchesVR(code)
	case float64:
		return float64MatchesVR(code)
	case time.Time:
		return code == vr.CodeDA || code == vr.CodeTM || code == vr.CodeDT
	}

	rv := reflect.ValueOf(value)
	for rv.IsValid() && (rv.Kind() == reflect.Interface || rv.Kind() == reflect.Pointer) {
		if rv.IsNil() {
			return true
		}
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return true
	}
	if rv.Type() == reflect.TypeFor[time.Time]() {
		return code == vr.CodeDA || code == vr.CodeTM || code == vr.CodeDT
	}
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		if rv.Type().Elem().Kind() == reflect.Uint8 && isRawVR(valueRepresentation) {
			return true
		}
		if rv.Len() == 0 {
			return true
		}
		elementType := rv.Type().Elem()
		if elementType == reflect.TypeFor[time.Time]() {
			return code == vr.CodeDA || code == vr.CodeTM || code == vr.CodeDT
		}
		if elementType.Kind() != reflect.Interface && elementType.Kind() != reflect.Pointer {
			return valueKindMatchesVR(valueRepresentation, elementType.Kind())
		}
		return valueMatchesVR(valueRepresentation, rv.Index(0).Interface())
	}

	return valueKindMatchesVR(valueRepresentation, rv.Kind())
}

func valueKindMatchesVR(valueRepresentation *vr.VR, kind reflect.Kind) bool {
	code := valueRepresentation.Code()
	switch kind {
	case reflect.String:
		return valueRepresentation.IsString()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return signedKindMatchesVR(code)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return unsignedKindMatchesVR(code)
	case reflect.Float32:
		return float32MatchesVR(code)
	case reflect.Float64:
		return float64MatchesVR(code)
	case reflect.Bool:
		return valueRepresentation.IsString()
	default:
		return false
	}
}

func signedKindMatchesVR(code string) bool {
	return code == vr.CodeDS || code == vr.CodeIS || code == vr.CodeSS || code == vr.CodeSL || code == vr.CodeSV
}

func unsignedKindMatchesVR(code string) bool {
	return code == vr.CodeDS || code == vr.CodeIS || code == vr.CodeUS || code == vr.CodeUL || code == vr.CodeUV
}

func float32MatchesVR(code string) bool {
	return code == vr.CodeDS || code == vr.CodeFL || code == vr.CodeFD
}

func float64MatchesVR(code string) bool {
	return code == vr.CodeDS || code == vr.CodeFD || code == vr.CodeFL
}

func isRawVR(valueRepresentation *vr.VR) bool {
	switch valueRepresentation.Code() {
	case vr.CodeOB, vr.CodeOW, vr.CodeOD, vr.CodeOF, vr.CodeOL, vr.CodeOV, vr.CodeUN:
		return true
	default:
		return false
	}
}
