// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"errors"
	"reflect"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
)

func TestCanonicalStringsPreservesValueBoundaries(t *testing.T) {
	tests := []struct {
		name string
		elem Element
		want []string
	}{
		{
			name: "text",
			elem: NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY"}),
			want: []string{"ORIGINAL", "PRIMARY"},
		},
		{
			name: "empty text",
			elem: NewString(tag.PatientName, vr.PN, nil),
			want: nil,
		},
		{
			name: "person name",
			elem: NewPersonName(tag.PatientName, []string{"Doe^Jane", "Smith^John"}),
			want: []string{"Doe^Jane", "Smith^John"},
		},
		{
			name: "decimal string",
			elem: NewDecimalString(tag.PixelSpacing, []string{"0.5", "0.75"}),
			want: []string{"0.5", "0.75"},
		},
		{
			name: "unsigned short",
			elem: NewUnsignedShort(tag.Rows, []uint16{512}),
			want: []string{"512"},
		},
		{
			name: "signed very long",
			elem: NewSignedVeryLong(tag.New(0x0011, 0x1010), []int64{-9, 12}),
			want: []string{"-9", "12"},
		},
		{
			name: "float",
			elem: NewFloat(tag.New(0x0011, 0x1011), []float32{1.25, -2.5}),
			want: []string{"1.25", "-2.5"},
		},
		{
			name: "double",
			elem: NewDouble(tag.New(0x0011, 0x1012), []float64{1.25}),
			want: []string{"1.25"},
		},
		{
			name: "attribute tag",
			elem: NewAttributeTag(tag.DimensionIndexPointer, []*tag.Tag{tag.Rows, tag.Columns}),
			want: []string{"(0028,0010)", "(0028,0011)"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CanonicalStrings(tc.elem)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("CanonicalStrings() = %#v, want %#v", got, tc.want)
			}
		})
	}
}

func TestCanonicalStringsRejectsUnsupportedValues(t *testing.T) {
	tests := []Element{
		NewOtherByte(tag.PixelData, []byte{1, 2}),
		NewOtherWord(tag.PixelData, []byte{1, 2}),
		NewFragmentSequence(tag.PixelData, vr.OB),
		&unsupportedSequenceElement{},
	}
	for _, elem := range tests {
		_, err := CanonicalStrings(elem)
		var unsupported *UnsupportedValueError
		if !errors.As(err, &unsupported) {
			t.Fatalf("CanonicalStrings(%T) error = %v, want UnsupportedValueError", elem, err)
		}
		if unsupported.Tag == nil || unsupported.VR == nil {
			t.Fatalf("UnsupportedValueError = %#v", unsupported)
		}
	}
}

func TestCanonicalStringsRejectsNilElement(t *testing.T) {
	_, err := CanonicalStrings(nil)
	var unsupported *UnsupportedValueError
	if !errors.As(err, &unsupported) {
		t.Fatalf("error = %v, want UnsupportedValueError", err)
	}
}

func TestReplaceCanonicalStringsPreservesNumericPrototypeByteOrder(t *testing.T) {
	prototype := NewUnsignedShortWithEndian(tag.Rows, []uint16{512}, endian.Big)

	got, err := ReplaceCanonicalStrings(prototype, tag.Rows, vr.US, []string{"1024"})
	if err != nil {
		t.Fatal(err)
	}
	values, err := got.(*UnsignedShort).GetValues()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(values, []uint16{1024}) {
		t.Fatalf("values = %v, want [1024]", values)
	}
	if !reflect.DeepEqual(got.Buffer().Data(), []byte{0x04, 0x00}) {
		t.Fatalf("buffer = %x, want big-endian 0400", got.Buffer().Data())
	}
}

func TestReplaceCanonicalStringsPreservesTextEncodingAndConcreteType(t *testing.T) {
	encoded, err := charset.EncodeString("旧^名", []encoding.Encoding{japanese.ShiftJIS})
	if err != nil {
		t.Fatal(err)
	}
	prototype := NewPersonNameFromBuffer(tag.PatientName, buffer.NewMemory(encoded), japanese.ShiftJIS)

	got, err := ReplaceCanonicalStrings(prototype, tag.PatientName, vr.PN, []string{"新^名"})
	if err != nil {
		t.Fatal(err)
	}
	personName, ok := got.(*PersonName)
	if !ok {
		t.Fatalf("replacement type = %T, want *PersonName", got)
	}
	if personName.str.encodings[0] != japanese.ShiftJIS {
		t.Fatal("replacement did not preserve Shift-JIS encoding")
	}
	values, err := CanonicalStrings(personName)
	if err != nil || !reflect.DeepEqual(values, []string{"新^名"}) {
		t.Fatalf("values = %v, err = %v", values, err)
	}
}

func TestReplaceCanonicalStringsRejectsLossyConversion(t *testing.T) {
	for _, value := range []string{"-1", "1.5", "65536"} {
		_, err := ReplaceCanonicalStrings(nil, tag.Rows, vr.US, []string{value})
		if err == nil {
			t.Fatalf("US conversion of %q succeeded, want error", value)
		}
	}
}

func TestReplaceCanonicalStringsBuildsSupportedVRWithoutPrototype(t *testing.T) {
	tests := []struct {
		name                string
		tag                 *tag.Tag
		valueRepresentation *vr.VR
		values              []string
		wantType            any
	}{
		{name: "string", tag: tag.PatientID, valueRepresentation: vr.LO, values: []string{"123"}, wantType: (*String)(nil)},
		{name: "person name", tag: tag.PatientName, valueRepresentation: vr.PN, values: []string{"Doe^Jane"}, wantType: (*PersonName)(nil)},
		{name: "date", tag: tag.StudyDate, valueRepresentation: vr.DA, values: []string{"20260816"}, wantType: (*Date)(nil)},
		{name: "decimal", tag: tag.PixelSpacing, valueRepresentation: vr.DS, values: []string{"0.5", "0.5"}, wantType: (*DecimalString)(nil)},
		{name: "unsigned short", tag: tag.Rows, valueRepresentation: vr.US, values: []string{"512"}, wantType: (*UnsignedShort)(nil)},
		{name: "signed long", tag: tag.New(0x0011, 0x1010), valueRepresentation: vr.SL, values: []string{"-4"}, wantType: (*SignedLong)(nil)},
		{name: "attribute tag", tag: tag.DimensionIndexPointer, valueRepresentation: vr.AT, values: []string{"(0028,0010)"}, wantType: (*AttributeTag)(nil)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ReplaceCanonicalStrings(nil, tc.tag, tc.valueRepresentation, tc.values)
			if err != nil {
				t.Fatal(err)
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tc.wantType) {
				t.Fatalf("replacement type = %T, want %T", got, tc.wantType)
			}
			values, err := CanonicalStrings(got)
			if err != nil || !reflect.DeepEqual(values, tc.values) {
				t.Fatalf("values = %v, err = %v, want %v", values, err, tc.values)
			}
		})
	}
}

func TestReplaceCanonicalStringsRejectsUnsupportedOrInvalidTargets(t *testing.T) {
	tests := []struct {
		name                string
		tag                 *tag.Tag
		valueRepresentation *vr.VR
	}{
		{name: "nil tag", tag: nil, valueRepresentation: vr.LO},
		{name: "nil VR", tag: tag.PatientID, valueRepresentation: nil},
		{name: "sequence", tag: tag.ReferencedImageSequence, valueRepresentation: vr.SQ},
		{name: "bulk", tag: tag.PixelData, valueRepresentation: vr.OB},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := ReplaceCanonicalStrings(nil, tc.tag, tc.valueRepresentation, []string{"x"}); err == nil {
				t.Fatal("ReplaceCanonicalStrings() succeeded, want error")
			}
		})
	}
}

type unsupportedSequenceElement struct{}

func (*unsupportedSequenceElement) Tag() *tag.Tag               { return tag.ReferencedImageSequence }
func (*unsupportedSequenceElement) ValueRepresentation() *vr.VR { return vr.SQ }
func (*unsupportedSequenceElement) Buffer() buffer.ByteBuffer   { return nil }
func (*unsupportedSequenceElement) Length() uint32              { return 0 }
func (*unsupportedSequenceElement) Count() int                  { return 0 }
func (*unsupportedSequenceElement) String() string              { return "sequence" }
func (*unsupportedSequenceElement) Validate() error             { return nil }
