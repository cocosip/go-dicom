// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package main

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func newParseResult(t *testing.T, mainDataset, metaDataset *dataset.Dataset) *parser.ParseResult {
	t.Helper()
	return &parser.ParseResult{
		Dataset:             mainDataset,
		FileMetaInformation: dataset.NewFileMetaInformationFromDataset(metaDataset),
	}
}

func TestToJSONWithMetaInfoUsesPS318TagKeys(t *testing.T) {
	mainDataset := dataset.New()
	if err := mainDataset.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3"})); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	metaDataset := dataset.New()
	if err := metaDataset.Add(element.NewString(tag.TransferSyntaxUID, vr.UI, []string{"1.2.840.10008.1.2.1"})); err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	out, err := toJSONWithMetaInfo(newParseResult(t, mainDataset, metaDataset), nil)
	if err != nil {
		t.Fatalf("toJSONWithMetaInfo() error = %v", err)
	}
	lower := strings.ToLower(string(out))
	for _, key := range []string{`"00020010"`, `"00080018"`} {
		if !strings.Contains(lower, key) {
			t.Errorf("JSON = %s, want PS3.18 tag key %s", lower, key)
		}
	}
	if strings.Contains(string(out), "TransferSyntaxUID") {
		t.Errorf("JSON = %s, want tag keys rather than keywords by default", out)
	}
}

func TestToJSONWithMetaInfoFailsInsteadOfEmittingPartialDocument(t *testing.T) {
	mainDataset := dataset.New()
	mainDataset.SetAutoValidate(false)
	// A UI value longer than 64 characters is invalid; the merge into the
	// validated meta dataset must fail the conversion.
	if err := mainDataset.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{strings.Repeat("1", 70)})); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	metaDataset := dataset.New()
	if err := metaDataset.Add(element.NewString(tag.TransferSyntaxUID, vr.UI, []string{"1.2.840.10008.1.2.1"})); err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	out, err := toJSONWithMetaInfo(newParseResult(t, mainDataset, metaDataset), nil)
	if err == nil {
		t.Fatalf("toJSONWithMetaInfo() = %s, nil; want a merge error", out)
	}
	if out != nil {
		t.Fatalf("toJSONWithMetaInfo() returned a partial document alongside the error: %s", out)
	}
}
