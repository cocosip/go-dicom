// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"fmt"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestCreateElementUsesVRSpecificTextTypes(t *testing.T) {
	ctx := newParseContext()
	tests := []struct {
		name string
		tag  *tag.Tag
		vr   *vr.VR
		want any
		data string
	}{
		{name: "decimal", tag: tag.RescaleSlope, vr: vr.DS, want: (*element.DecimalString)(nil), data: "1 "},
		{name: "integer", tag: tag.NumberOfFrames, vr: vr.IS, want: (*element.IntegerString)(nil), data: "2 "},
		{name: "date", tag: tag.StudyDate, vr: vr.DA, want: (*element.Date)(nil), data: "20250101"},
		{name: "time", tag: tag.StudyTime, vr: vr.TM, want: (*element.Time)(nil), data: "120000"},
		{name: "datetime", tag: tag.AcquisitionDateTime, vr: vr.DT, want: (*element.DateTime)(nil), data: "20250101120000"},
		{name: "person name", tag: tag.PatientName, vr: vr.PN, want: (*element.PersonName)(nil), data: "Doe^John"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ctx.createElement(tt.tag, tt.vr, buffer.NewMemory([]byte(tt.data)))
			if err != nil {
				t.Fatalf("createElement() error = %v", err)
			}
			if reflectType := typeOf(got); reflectType != typeOf(tt.want) {
				t.Fatalf("createElement() type = %T, want %T", got, tt.want)
			}
		})
	}
}

func typeOf(value any) string {
	return fmt.Sprintf("%T", value)
}
