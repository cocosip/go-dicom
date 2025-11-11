// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"
	"os"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// FileMetaInformation represents the DICOM File Meta Information (Group 0x0002).
//
// This wraps a Dataset and provides convenient getter/setter methods for
// the standard File Meta Information elements defined in DICOM Part 10.
//
// According to DICOM PS3.10, the File Meta Information includes:
//   - FileMetaInformationVersion (0002,0001) - always {0x00, 0x01}
//   - MediaStorageSOPClassUID (0002,0002) - required
//   - MediaStorageSOPInstanceUID (0002,0003) - required
//   - TransferSyntaxUID (0002,0010) - required
//   - ImplementationClassUID (0002,0012) - required
//   - ImplementationVersionName (0002,0013) - optional
//   - SourceApplicationEntityTitle (0002,0016) - optional
//   - SendingApplicationEntityTitle (0002,0017) - optional
//   - ReceivingApplicationEntityTitle (0002,0018) - optional
//   - PrivateInformationCreatorUID (0002,0100) - optional
//   - PrivateInformation (0002,0102) - optional
type FileMetaInformation struct {
	// dataset is the underlying dataset containing the meta information
	dataset *Dataset
}

// NewFileMetaInformation creates a new empty FileMetaInformation.
func NewFileMetaInformation() *FileMetaInformation {
	return &FileMetaInformation{
		dataset: New(),
	}
}

// NewFileMetaInformationFromDataset creates a FileMetaInformation from an existing dataset.
// This wraps the provided dataset without copying it.
func NewFileMetaInformationFromDataset(ds *Dataset) *FileMetaInformation {
	return &FileMetaInformation{
		dataset: ds,
	}
}

// NewDefaultFileMetaInformation creates a FileMetaInformation with default values for required fields.
// This includes:
//   - Version: {0x00, 0x01}
//   - ImplementationClassUID: go-dicom implementation UID
//   - ImplementationVersionName: go-dicom version
//
// The caller should set MediaStorageSOPClassUID, MediaStorageSOPInstanceUID,
// and TransferSyntaxUID before writing the file.
func NewDefaultFileMetaInformation() *FileMetaInformation {
	fmi := NewFileMetaInformation()
	_ = fmi.SetVersion([]byte{0x00, 0x01})
	_ = fmi.SetImplementationClassUID(uid.GoDicomImplementationClassUID)
	_ = fmi.SetImplementationVersionName(uid.GoDicomImplementationVersionName)
	return fmi
}

// NewFileMetaInformationFromMainDataset creates FileMetaInformation from a main dataset.
// It extracts SOPClassUID and SOPInstanceUID from the dataset and uses
// the provided transfer syntax.
func NewFileMetaInformationFromMainDataset(ds *Dataset, ts *transfer.Syntax) (*FileMetaInformation, error) {
	fmi := NewDefaultFileMetaInformation()

	// Extract SOP Class UID
	sopClassUID, ok := ds.GetString(tag.SOPClassUID)
	if !ok {
		return nil, fmt.Errorf("SOPClassUID not found in dataset")
	}
	if err := fmi.SetMediaStorageSOPClassUID(sopClassUID); err != nil {
		return nil, fmt.Errorf("failed to set media storage SOP class UID: %w", err)
	}

	// Extract SOP Instance UID
	sopInstanceUID, ok := ds.GetString(tag.SOPInstanceUID)
	if !ok {
		return nil, fmt.Errorf("SOPInstanceUID not found in dataset")
	}
	if err := fmi.SetMediaStorageSOPInstanceUID(sopInstanceUID); err != nil {
		return nil, fmt.Errorf("failed to set media storage SOP instance UID: %w", err)
	}

	// Set transfer syntax
	if err := fmi.SetTransferSyntaxUID(ts.UID().UID()); err != nil {
		return nil, fmt.Errorf("failed to set transfer syntax UID: %w", err)
	}

	// Try to get optional AE titles
	if sourceAET, ok := ds.GetString(tag.SourceApplicationEntityTitle); ok {
		_ = fmi.SetSourceApplicationEntityTitle(sourceAET)
	}
	if sendingAET, ok := ds.GetString(tag.SendingApplicationEntityTitle); ok {
		_ = fmi.SetSendingApplicationEntityTitle(sendingAET)
	}
	if receivingAET, ok := ds.GetString(tag.ReceivingApplicationEntityTitle); ok {
		_ = fmi.SetReceivingApplicationEntityTitle(receivingAET)
	}

	return fmi, nil
}

// Dataset returns the underlying dataset.
// This allows direct access to all dataset operations.
func (fmi *FileMetaInformation) Dataset() *Dataset {
	return fmi.dataset
}

// Version returns the File Meta Information Version (0002,0001).
// This should always be {0x00, 0x01}.
func (fmi *FileMetaInformation) Version() ([]byte, bool) {
	elem, exists := fmi.dataset.Get(tag.FileMetaInformationVersion)
	if !exists {
		return nil, false
	}
	buf := elem.Buffer()
	if buf == nil {
		return nil, false
	}
	return buf.Data(), true
}

// SetVersion sets the File Meta Information Version (0002,0001).
func (fmi *FileMetaInformation) SetVersion(version []byte) error {
	return fmi.dataset.AddOrUpdate(element.NewOtherByte(tag.FileMetaInformationVersion, version))
}

// MediaStorageSOPClassUID returns the Media Storage SOP Class UID (0002,0002).
func (fmi *FileMetaInformation) MediaStorageSOPClassUID() (string, bool) {
	return fmi.dataset.GetString(tag.MediaStorageSOPClassUID)
}

// SetMediaStorageSOPClassUID sets the Media Storage SOP Class UID (0002,0002).
func (fmi *FileMetaInformation) SetMediaStorageSOPClassUID(uid string) error {
	return fmi.dataset.AddOrUpdate(element.NewString(tag.MediaStorageSOPClassUID, vr.UI, []string{uid}))
}

// MediaStorageSOPInstanceUID returns the Media Storage SOP Instance UID (0002,0003).
func (fmi *FileMetaInformation) MediaStorageSOPInstanceUID() (string, bool) {
	return fmi.dataset.GetString(tag.MediaStorageSOPInstanceUID)
}

// SetMediaStorageSOPInstanceUID sets the Media Storage SOP Instance UID (0002,0003).
func (fmi *FileMetaInformation) SetMediaStorageSOPInstanceUID(uid string) error {
	return fmi.dataset.AddOrUpdate(element.NewString(tag.MediaStorageSOPInstanceUID, vr.UI, []string{uid}))
}

// TransferSyntaxUID returns the Transfer Syntax UID (0002,0010).
func (fmi *FileMetaInformation) TransferSyntaxUID() (string, bool) {
	return fmi.dataset.GetString(tag.TransferSyntaxUID)
}

// SetTransferSyntaxUID sets the Transfer Syntax UID (0002,0010).
func (fmi *FileMetaInformation) SetTransferSyntaxUID(uid string) error {
	return fmi.dataset.AddOrUpdate(element.NewString(tag.TransferSyntaxUID, vr.UI, []string{uid}))
}

// TransferSyntax returns the parsed TransferSyntax object.
func (fmi *FileMetaInformation) TransferSyntax() (*transfer.Syntax, bool) {
	tsUID, ok := fmi.TransferSyntaxUID()
	if !ok {
		return nil, false
	}
	ts, err := transfer.Parse(tsUID)
	if err != nil {
		return nil, false
	}
	return ts, true
}

// SetTransferSyntax sets the Transfer Syntax UID from a TransferSyntax object.
func (fmi *FileMetaInformation) SetTransferSyntax(ts *transfer.Syntax) error {
	return fmi.SetTransferSyntaxUID(ts.UID().UID())
}

// ImplementationClassUID returns the Implementation Class UID (0002,0012).
func (fmi *FileMetaInformation) ImplementationClassUID() (string, bool) {
	return fmi.dataset.GetString(tag.ImplementationClassUID)
}

// SetImplementationClassUID sets the Implementation Class UID (0002,0012).
func (fmi *FileMetaInformation) SetImplementationClassUID(uid string) error {
	return fmi.dataset.AddOrUpdate(element.NewString(tag.ImplementationClassUID, vr.UI, []string{uid}))
}

// ImplementationVersionName returns the Implementation Version Name (0002,0013).
func (fmi *FileMetaInformation) ImplementationVersionName() (string, bool) {
	return fmi.dataset.GetString(tag.ImplementationVersionName)
}

// SetImplementationVersionName sets the Implementation Version Name (0002,0013).
func (fmi *FileMetaInformation) SetImplementationVersionName(name string) error {
	return fmi.dataset.AddOrUpdate(element.NewString(tag.ImplementationVersionName, vr.SH, []string{name}))
}

// SourceApplicationEntityTitle returns the Source Application Entity Title (0002,0016).
func (fmi *FileMetaInformation) SourceApplicationEntityTitle() (string, bool) {
	return fmi.dataset.GetString(tag.SourceApplicationEntityTitle)
}

// SetSourceApplicationEntityTitle sets the Source Application Entity Title (0002,0016).
func (fmi *FileMetaInformation) SetSourceApplicationEntityTitle(aet string) error {
	return fmi.dataset.AddOrUpdate(element.NewString(tag.SourceApplicationEntityTitle, vr.AE, []string{aet}))
}

// SendingApplicationEntityTitle returns the Sending Application Entity Title (0002,0017).
func (fmi *FileMetaInformation) SendingApplicationEntityTitle() (string, bool) {
	return fmi.dataset.GetString(tag.SendingApplicationEntityTitle)
}

// SetSendingApplicationEntityTitle sets the Sending Application Entity Title (0002,0017).
func (fmi *FileMetaInformation) SetSendingApplicationEntityTitle(aet string) error {
	return fmi.dataset.AddOrUpdate(element.NewString(tag.SendingApplicationEntityTitle, vr.AE, []string{aet}))
}

// ReceivingApplicationEntityTitle returns the Receiving Application Entity Title (0002,0018).
func (fmi *FileMetaInformation) ReceivingApplicationEntityTitle() (string, bool) {
	return fmi.dataset.GetString(tag.ReceivingApplicationEntityTitle)
}

// SetReceivingApplicationEntityTitle sets the Receiving Application Entity Title (0002,0018).
func (fmi *FileMetaInformation) SetReceivingApplicationEntityTitle(aet string) error {
	return fmi.dataset.AddOrUpdate(element.NewString(tag.ReceivingApplicationEntityTitle, vr.AE, []string{aet}))
}

// PrivateInformationCreatorUID returns the Private Information Creator UID (0002,0100).
func (fmi *FileMetaInformation) PrivateInformationCreatorUID() (string, bool) {
	return fmi.dataset.GetString(tag.PrivateInformationCreatorUID)
}

// SetPrivateInformationCreatorUID sets the Private Information Creator UID (0002,0100).
func (fmi *FileMetaInformation) SetPrivateInformationCreatorUID(uid string) error {
	return fmi.dataset.AddOrUpdate(element.NewString(tag.PrivateInformationCreatorUID, vr.UI, []string{uid}))
}

// PrivateInformation returns the Private Information (0002,0102).
func (fmi *FileMetaInformation) PrivateInformation() ([]byte, bool) {
	elem, exists := fmi.dataset.Get(tag.PrivateInformation)
	if !exists {
		return nil, false
	}
	buf := elem.Buffer()
	if buf == nil {
		return nil, false
	}
	return buf.Data(), true
}

// SetPrivateInformation sets the Private Information (0002,0102).
func (fmi *FileMetaInformation) SetPrivateInformation(data []byte) error {
	return fmi.dataset.AddOrUpdate(element.NewOtherByte(tag.PrivateInformation, data))
}

// Validate checks if the FileMetaInformation contains all required elements.
func (fmi *FileMetaInformation) Validate() error {
	// Check required elements
	if _, ok := fmi.Version(); !ok {
		return fmt.Errorf("missing required FileMetaInformationVersion")
	}
	if _, ok := fmi.MediaStorageSOPClassUID(); !ok {
		return fmt.Errorf("missing required MediaStorageSOPClassUID")
	}
	if _, ok := fmi.MediaStorageSOPInstanceUID(); !ok {
		return fmt.Errorf("missing required MediaStorageSOPInstanceUID")
	}
	if _, ok := fmi.TransferSyntaxUID(); !ok {
		return fmt.Errorf("missing required TransferSyntaxUID")
	}
	if _, ok := fmi.ImplementationClassUID(); !ok {
		return fmt.Errorf("missing required ImplementationClassUID")
	}

	// Validate that all tags are in group 0x0002
	for _, elem := range fmi.dataset.Elements() {
		if elem.Tag().Group() != 0x0002 {
			return fmt.Errorf("tag %s is not allowed in file meta information (must be in group 0x0002)", elem.Tag().String())
		}
	}

	return nil
}

// String returns a string representation of the FileMetaInformation.
func (fmi *FileMetaInformation) String() string {
	return "DICOM File Meta Information"
}

// CreateDefaultSourceAET creates a default Source Application Entity Title from the hostname.
// The returned value is truncated to 16 characters as per DICOM standard.
func CreateDefaultSourceAET() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "GO-DICOM"
	}
	if len(hostname) > 16 {
		return hostname[:16]
	}
	return hostname
}
