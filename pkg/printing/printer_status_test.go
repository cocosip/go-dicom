// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestParsePrinterStatus(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		expected  PrinterStatus
		expectErr bool
	}{
		{name: "normal", input: "NORMAL", expected: PrinterStatusNormal},
		{name: "warning lowercase", input: "warning", expected: PrinterStatusWarning},
		{name: "failure with spaces", input: "  FAILURE  ", expected: PrinterStatusFailure},
		{name: "empty", input: "", expected: PrinterStatusUnknown},
		{name: "invalid", input: "OFFLINE", expectErr: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := ParsePrinterStatus(tc.input)
			if tc.expectErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if status != tc.expected {
				t.Fatalf("expected %q, got %q", tc.expected, status)
			}
		})
	}
}

func TestNewPrinterStatusManager(t *testing.T) {
	manager := NewPrinterStatusManager("  DRY-PRINTER-01  ")
	if manager == nil {
		t.Fatal("NewPrinterStatusManager returned nil")
	}

	if manager.PrinterName() != "DRY-PRINTER-01" {
		t.Fatalf("unexpected printer name: %q", manager.PrinterName())
	}
	if manager.Status() != PrinterStatusUnknown {
		t.Fatalf("expected UNKNOWN status, got %q", manager.Status())
	}
	if manager.StatusInfo() != "" {
		t.Fatalf("expected empty status info, got %q", manager.StatusInfo())
	}
	if !manager.LastUpdated().IsZero() {
		t.Fatal("expected zero LastUpdated")
	}
	if len(manager.History()) != 0 {
		t.Fatalf("expected empty history, got %d", len(manager.History()))
	}
}

func TestPrinterStatusManagerUpdateAt(t *testing.T) {
	manager := NewPrinterStatusManager("PRINTER-A")

	t1 := time.Date(2026, 3, 5, 10, 0, 0, 0, time.UTC)
	if err := manager.UpdateAt(PrinterStatusNormal, "READY", t1); err != nil {
		t.Fatalf("UpdateAt returned error: %v", err)
	}

	if manager.Status() != PrinterStatusNormal {
		t.Fatalf("expected NORMAL, got %q", manager.Status())
	}
	if manager.StatusInfo() != "READY" {
		t.Fatalf("expected READY info, got %q", manager.StatusInfo())
	}
	if !manager.LastUpdated().Equal(t1) {
		t.Fatalf("expected updated time %v, got %v", t1, manager.LastUpdated())
	}

	t2 := t1.Add(5 * time.Minute)
	if err := manager.UpdateAt(PrinterStatusWarning, "LOW FILM", t2); err != nil {
		t.Fatalf("UpdateAt returned error: %v", err)
	}

	history := manager.History()
	if len(history) != 2 {
		t.Fatalf("expected 2 history entries, got %d", len(history))
	}
	if history[1].Status != PrinterStatusWarning {
		t.Fatalf("expected second status WARNING, got %q", history[1].Status)
	}
	if history[1].StatusInfo != "LOW FILM" {
		t.Fatalf("expected second info LOW FILM, got %q", history[1].StatusInfo)
	}
}

func TestPrinterStatusManagerUpdateRejectsUnknown(t *testing.T) {
	manager := NewPrinterStatusManager("PRINTER-A")
	if err := manager.UpdateAt(PrinterStatusUnknown, "", time.Time{}); err == nil {
		t.Fatal("expected error when updating with UNKNOWN status")
	}
	if len(manager.History()) != 0 {
		t.Fatalf("expected no history entries, got %d", len(manager.History()))
	}
}

func TestPrinterStatusManagerUpdateFromDataset(t *testing.T) {
	manager := NewPrinterStatusManager("OLD-NAME")
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewString(tag.PrinterName, vr.LO, []string{"NEW-NAME"}),
		element.NewString(tag.PrinterStatus, vr.CS, []string{"WARNING"}),
		element.NewString(tag.PrinterStatusInfo, vr.CS, []string{"CHECK TONER"}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	if err := manager.UpdateFromDataset(ds); err != nil {
		t.Fatalf("UpdateFromDataset returned error: %v", err)
	}

	if manager.PrinterName() != "NEW-NAME" {
		t.Fatalf("expected printer name NEW-NAME, got %q", manager.PrinterName())
	}
	if manager.Status() != PrinterStatusWarning {
		t.Fatalf("expected WARNING, got %q", manager.Status())
	}
	if manager.StatusInfo() != "CHECK TONER" {
		t.Fatalf("expected status info CHECK TONER, got %q", manager.StatusInfo())
	}
	if len(manager.History()) != 1 {
		t.Fatalf("expected 1 history entry, got %d", len(manager.History()))
	}
}

func TestPrinterStatusManagerUpdateFromDatasetInfoOnly(t *testing.T) {
	manager := NewPrinterStatusManager("PRINTER-A")
	if err := manager.UpdateAt(PrinterStatusNormal, "READY", time.Date(2026, 3, 5, 9, 0, 0, 0, time.UTC)); err != nil {
		t.Fatalf("UpdateAt returned error: %v", err)
	}

	ds, err := dataset.NewWithElements([]element.Element{
		element.NewString(tag.PrinterStatusInfo, vr.CS, []string{"TRAY LOW"}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	if err := manager.UpdateFromDataset(ds); err != nil {
		t.Fatalf("UpdateFromDataset returned error: %v", err)
	}

	if manager.Status() != PrinterStatusNormal {
		t.Fatalf("status should remain NORMAL, got %q", manager.Status())
	}
	if manager.StatusInfo() != "TRAY LOW" {
		t.Fatalf("expected info TRAY LOW, got %q", manager.StatusInfo())
	}
	if len(manager.History()) != 2 {
		t.Fatalf("expected 2 history entries, got %d", len(manager.History()))
	}
}

func TestPrinterStatusManagerUpdateFromDatasetInvalidStatus(t *testing.T) {
	manager := NewPrinterStatusManager("PRINTER-A")
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewString(tag.PrinterStatus, vr.CS, []string{"OFFLINE"}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	if err := manager.UpdateFromDataset(ds); err == nil {
		t.Fatal("expected error for invalid status")
	}
}

func TestPrinterStatusManagerApplyToDataset(t *testing.T) {
	manager := NewPrinterStatusManager("PRINTER-A")
	if err := manager.UpdateAt(PrinterStatusFailure, "NO PAPER", time.Date(2026, 3, 5, 11, 0, 0, 0, time.UTC)); err != nil {
		t.Fatalf("UpdateAt returned error: %v", err)
	}

	ds, err := dataset.NewWithElements([]element.Element{
		element.NewString(tag.PrinterStatus, vr.CS, []string{"NORMAL"}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}
	if err := manager.ApplyToDataset(ds); err != nil {
		t.Fatalf("ApplyToDataset returned error: %v", err)
	}

	printerName, ok := ds.GetString(tag.PrinterName)
	if !ok || printerName != "PRINTER-A" {
		t.Fatalf("unexpected printer name in dataset: %q", printerName)
	}
	status, ok := ds.GetString(tag.PrinterStatus)
	if !ok || status != "FAILURE" {
		t.Fatalf("unexpected status in dataset: %q", status)
	}
	statusInfo, ok := ds.GetString(tag.PrinterStatusInfo)
	if !ok || statusInfo != "NO PAPER" {
		t.Fatalf("unexpected status info in dataset: %q", statusInfo)
	}

	generated := manager.ToDataset()
	if generated == nil {
		t.Fatal("ToDataset returned nil")
	}
	if got, _ := generated.GetString(tag.PrinterStatus); got != "FAILURE" {
		t.Fatalf("expected FAILURE in generated dataset, got %q", got)
	}
}

func TestPrinterStatusManagerStateHelpers(t *testing.T) {
	manager := NewPrinterStatusManager("PRINTER-A")

	if manager.IsAvailable() {
		t.Fatal("UNKNOWN state should not be available")
	}
	if manager.NeedsAttention() {
		t.Fatal("UNKNOWN state should not need attention")
	}

	if err := manager.UpdateAt(PrinterStatusNormal, "", time.Now().UTC()); err != nil {
		t.Fatalf("UpdateAt returned error: %v", err)
	}
	if !manager.IsAvailable() {
		t.Fatal("NORMAL state should be available")
	}
	if manager.NeedsAttention() {
		t.Fatal("NORMAL state should not need attention")
	}

	if err := manager.UpdateAt(PrinterStatusWarning, "", time.Now().UTC()); err != nil {
		t.Fatalf("UpdateAt returned error: %v", err)
	}
	if !manager.IsAvailable() {
		t.Fatal("WARNING state should be available")
	}
	if !manager.NeedsAttention() {
		t.Fatal("WARNING state should need attention")
	}

	if err := manager.UpdateAt(PrinterStatusFailure, "", time.Now().UTC()); err != nil {
		t.Fatalf("UpdateAt returned error: %v", err)
	}
	if manager.IsAvailable() {
		t.Fatal("FAILURE state should not be available")
	}
	if !manager.NeedsAttention() {
		t.Fatal("FAILURE state should need attention")
	}
	if !manager.HasFailure() {
		t.Fatal("FAILURE state should report failure")
	}
}
