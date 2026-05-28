// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// PrinterStatus represents the operational state of a DICOM printer.
type PrinterStatus string

const (
	// PrinterStatusUnknown means no status has been reported yet.
	PrinterStatusUnknown PrinterStatus = "UNKNOWN"
	// PrinterStatusNormal means printer is operating normally.
	PrinterStatusNormal PrinterStatus = "NORMAL"
	// PrinterStatusWarning means printer can operate but needs attention.
	PrinterStatusWarning PrinterStatus = "WARNING"
	// PrinterStatusFailure means printer cannot process new print jobs.
	PrinterStatusFailure PrinterStatus = "FAILURE"
)

// IsValid returns true when the status is known by this package.
func (s PrinterStatus) IsValid() bool {
	switch s {
	case PrinterStatusUnknown, PrinterStatusNormal, PrinterStatusWarning, PrinterStatusFailure:
		return true
	default:
		return false
	}
}

// IsKnown returns true for DICOM-defined operational states.
func (s PrinterStatus) IsKnown() bool {
	switch s {
	case PrinterStatusNormal, PrinterStatusWarning, PrinterStatusFailure:
		return true
	default:
		return false
	}
}

// ParsePrinterStatus parses a printer status value from dataset/string input.
func ParsePrinterStatus(value string) (PrinterStatus, error) {
	normalized := strings.ToUpper(strings.TrimSpace(value))
	if normalized == "" {
		return PrinterStatusUnknown, nil
	}

	status := PrinterStatus(normalized)
	if !status.IsValid() {
		return PrinterStatusUnknown, fmt.Errorf("invalid printer status: %q", value)
	}

	return status, nil
}

// PrinterStatusSnapshot is a point-in-time printer status record.
type PrinterStatusSnapshot struct {
	PrinterName string
	Status      PrinterStatus
	StatusInfo  string
	UpdatedAt   time.Time
}

// PrinterStatusManager tracks and updates printer status for DICOM Print Management.
type PrinterStatusManager struct {
	mu          sync.Mutex
	printerName string
	status      PrinterStatus
	statusInfo  string
	lastUpdated time.Time
	history     []PrinterStatusSnapshot
}

// NewPrinterStatusManager creates a new printer status manager.
func NewPrinterStatusManager(printerName string) *PrinterStatusManager {
	return &PrinterStatusManager{
		printerName: strings.TrimSpace(printerName),
		status:      PrinterStatusUnknown,
		history:     make([]PrinterStatusSnapshot, 0),
	}
}

// PrinterName returns the configured printer name.
func (m *PrinterStatusManager) PrinterName() string {
	return m.printerName
}

// SetPrinterName updates the printer name.
func (m *PrinterStatusManager) SetPrinterName(printerName string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.printerName = strings.TrimSpace(printerName)
}

// Status returns the current printer status.
func (m *PrinterStatusManager) Status() PrinterStatus {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.status
}

// StatusInfo returns the current status details.
func (m *PrinterStatusManager) StatusInfo() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.statusInfo
}

// LastUpdated returns the timestamp of the most recent status update.
func (m *PrinterStatusManager) LastUpdated() time.Time {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.lastUpdated
}

// Snapshot returns the current status snapshot.
func (m *PrinterStatusManager) Snapshot() PrinterStatusSnapshot {
	m.mu.Lock()
	defer m.mu.Unlock()
	return PrinterStatusSnapshot{
		PrinterName: m.printerName,
		Status:      m.status,
		StatusInfo:  m.statusInfo,
		UpdatedAt:   m.lastUpdated,
	}
}

// History returns a copy of all status snapshots.
func (m *PrinterStatusManager) History() []PrinterStatusSnapshot {
	m.mu.Lock()
	defer m.mu.Unlock()
	history := make([]PrinterStatusSnapshot, len(m.history))
	copy(history, m.history)
	return history
}

// IsAvailable returns true when the printer can still accept print jobs.
func (m *PrinterStatusManager) IsAvailable() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.status == PrinterStatusNormal || m.status == PrinterStatusWarning
}

// NeedsAttention returns true when the printer requires operator attention.
func (m *PrinterStatusManager) NeedsAttention() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.status == PrinterStatusWarning || m.status == PrinterStatusFailure
}

// HasFailure returns true when the printer is in FAILURE state.
func (m *PrinterStatusManager) HasFailure() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.status == PrinterStatusFailure
}

// Update updates printer status using the current time.
func (m *PrinterStatusManager) Update(status PrinterStatus, statusInfo string) error {
	return m.UpdateAt(status, statusInfo, time.Now().UTC())
}

// UpdateAt updates printer status at a specific timestamp.
func (m *PrinterStatusManager) UpdateAt(status PrinterStatus, statusInfo string, updatedAt time.Time) error {
	if !status.IsKnown() {
		return fmt.Errorf("invalid printer status for update: %q", status)
	}

	if updatedAt.IsZero() {
		updatedAt = time.Now().UTC()
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.status = status
	m.statusInfo = strings.TrimSpace(statusInfo)
	m.lastUpdated = updatedAt
	m.history = append(m.history, PrinterStatusSnapshot{
		PrinterName: m.printerName,
		Status:      m.status,
		StatusInfo:  m.statusInfo,
		UpdatedAt:   m.lastUpdated,
	})

	return nil
}

// UpdateFromDataset updates printer status fields from a DICOM dataset.
func (m *PrinterStatusManager) UpdateFromDataset(ds *dataset.Dataset) error {
	if ds == nil {
		return fmt.Errorf("dataset is nil")
	}

	m.mu.Lock()
	// No defer unlock — the hasStatus path releases the lock early before calling Update().

	if printerName, exists := ds.GetString(tag.PrinterName); exists {
		m.printerName = strings.TrimSpace(printerName)
	}

	statusValue, hasStatus := ds.GetString(tag.PrinterStatus)
	infoValue, hasInfo := ds.GetString(tag.PrinterStatusInfo)

	if hasStatus {
		status, err := ParsePrinterStatus(statusValue)
		if err != nil {
			m.mu.Unlock()
			return err
		}
		if status == PrinterStatusUnknown {
			m.mu.Unlock()
			return fmt.Errorf("printer status is empty")
		}
		if !hasInfo {
			infoValue = m.statusInfo
		}
		// Release lock before calling Update (which acquires it)
		m.mu.Unlock()
		return m.Update(status, infoValue)
	}

	if hasInfo {
		m.statusInfo = strings.TrimSpace(infoValue)
		m.lastUpdated = time.Now().UTC()
		m.history = append(m.history, PrinterStatusSnapshot{
			PrinterName: m.printerName,
			Status:      m.status,
			StatusInfo:  m.statusInfo,
			UpdatedAt:   m.lastUpdated,
		})
	}

	m.mu.Unlock()
	return nil
}

// ApplyToDataset writes current printer status values into an existing dataset.
func (m *PrinterStatusManager) ApplyToDataset(ds *dataset.Dataset) error {
	if ds == nil {
		return fmt.Errorf("dataset is nil")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.printerName != "" {
		if err := ds.AddOrUpdate(element.NewString(tag.PrinterName, vr.LO, []string{m.printerName})); err != nil {
			return err
		}
	}

	if m.status.IsKnown() {
		if err := ds.AddOrUpdate(element.NewString(tag.PrinterStatus, vr.CS, []string{string(m.status)})); err != nil {
			return err
		}
	}

	if m.statusInfo != "" {
		if err := ds.AddOrUpdate(element.NewString(tag.PrinterStatusInfo, vr.CS, []string{m.statusInfo})); err != nil {
			return err
		}
	}

	return nil
}

// ToDataset creates a new dataset with the current printer status fields.
func (m *PrinterStatusManager) ToDataset() *dataset.Dataset {
	ds := dataset.New()
	_ = m.ApplyToDataset(ds)
	return ds
}
