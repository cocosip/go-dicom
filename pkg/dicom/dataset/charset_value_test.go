// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset_test

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)

const testGB18030CharacterSet = "GB18030"

func TestAddValueUsesDatasetSpecificCharacterSet(t *testing.T) {
	tests := []struct {
		name     string
		charSet  string
		encoding func(string) ([]byte, error)
	}{
		{
			name:    "UTF-8",
			charSet: "ISO_IR 192",
			encoding: func(value string) ([]byte, error) {
				return unicode.UTF8.NewEncoder().Bytes([]byte(value))
			},
		},
		{
			name:    testGB18030CharacterSet,
			charSet: testGB18030CharacterSet,
			encoding: func(value string) ([]byte, error) {
				return simplifiedchinese.GB18030.NewEncoder().Bytes([]byte(value))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.New()
			if err := ds.SetSpecificCharacterSet(tt.charSet); err != nil {
				t.Fatalf("SetSpecificCharacterSet() error = %v", err)
			}
			if err := ds.AddValue(tag.PatientName, "张三"); err != nil {
				t.Fatalf("AddValue() error = %v", err)
			}

			got, err := ds.GetBytes(tag.PatientName)
			if err != nil {
				t.Fatalf("GetBytes() error = %v", err)
			}
			want, err := tt.encoding("张三")
			if err != nil {
				t.Fatalf("encode expected value: %v", err)
			}
			if !bytes.Equal(got, want) {
				t.Fatalf("PatientName bytes = %x, want %x", got, want)
			}
			if gotName, ok := ds.GetString(tag.PatientName); !ok || gotName != "张三" {
				t.Fatalf("GetString(PatientName) = %q, %v; want 张三, true", gotName, ok)
			}
			if gotCharsets := ds.SpecificCharacterSet(); len(gotCharsets) != 1 || gotCharsets[0] != tt.charSet {
				t.Fatalf("SpecificCharacterSet() = %v, want [%s]", gotCharsets, tt.charSet)
			}
		})
	}
}

func TestSetSpecificCharacterSetRejectsUnknownValue(t *testing.T) {
	ds := dataset.New()
	if err := ds.SetSpecificCharacterSet("NOT_A_DICOM_CHARSET"); err == nil {
		t.Fatal("SetSpecificCharacterSet() error = nil, want unsupported character set error")
	}
	if ds.Contains(tag.SpecificCharacterSet) {
		t.Fatal("SetSpecificCharacterSet() modified Dataset after rejecting the value")
	}
}

func TestGetBytesAllowsExplicitRecoveryFromIncorrectDeclaredCharacterSet(t *testing.T) {
	raw, err := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("张三"))
	if err != nil {
		t.Fatalf("encode GB18030 fixture: %v", err)
	}

	ds := dataset.New()
	if err := ds.SetSpecificCharacterSet("ISO_IR 192"); err != nil {
		t.Fatalf("SetSpecificCharacterSet() error = %v", err)
	}
	malformed := element.NewStringFromBufferWithEncodings(
		tag.PatientName,
		vr.PN,
		buffer.NewMemory(raw),
		[]encoding.Encoding{unicode.UTF8},
	)
	if err := ds.Add(malformed); err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	gotRaw, err := ds.GetBytes(tag.PatientName)
	if err != nil {
		t.Fatalf("GetBytes() error = %v", err)
	}
	if !bytes.Equal(gotRaw, raw) {
		t.Fatalf("GetBytes() = %x, want original %x", gotRaw, raw)
	}
	decoded, err := charset.DecodeString(gotRaw, charset.GetEncodings([]string{testGB18030CharacterSet}))
	if err != nil {
		t.Fatalf("DecodeString(GB18030) error = %v", err)
	}
	if decoded != "张三" {
		t.Fatalf("DecodeString(GB18030) = %q, want 张三", decoded)
	}
}

func TestAddedSequenceItemInheritsDatasetSpecificCharacterSet(t *testing.T) {
	ds := dataset.New()
	if err := ds.SetSpecificCharacterSet("ISO_IR 192"); err != nil {
		t.Fatalf("SetSpecificCharacterSet() error = %v", err)
	}
	item := dataset.New()
	if err := ds.Add(dataset.NewSequenceWithItems(tag.RequestAttributesSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatalf("Add(sequence) error = %v", err)
	}
	if err := item.AddValue(tag.StudyDescription, "中文描述"); err != nil {
		t.Fatalf("item.AddValue() error = %v", err)
	}
	if got, ok := item.GetString(tag.StudyDescription); !ok || got != "中文描述" {
		t.Fatalf("GetString(StudyDescription) = %q, %v; want 中文描述, true", got, ok)
	}
}
