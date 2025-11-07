// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewFileMetaInformation(t *testing.T) {
	fmi := NewFileMetaInformation()
	if fmi == nil {
		t.Fatal("NewFileMetaInformation returned nil")
	}
	if fmi.dataset == nil {
		t.Fatal("FileMetaInformation dataset is nil")
	}
}

func TestNewDefaultFileMetaInformation(t *testing.T) {
	fmi := NewDefaultFileMetaInformation()
	if fmi == nil {
		t.Fatal("NewDefaultFileMetaInformation returned nil")
	}

	// Check version
	version, ok := fmi.Version()
	if !ok {
		t.Error("Version not set")
	} else if len(version) != 2 || version[0] != 0x00 || version[1] != 0x01 {
		t.Errorf("Expected version {0x00, 0x01}, got %v", version)
	}

	// Check implementation class UID
	implClassUID, ok := fmi.ImplementationClassUID()
	if !ok {
		t.Error("ImplementationClassUID not set")
	} else if implClassUID != uid.GoDicomImplementationClassUID {
		t.Errorf("Expected ImplementationClassUID %s, got %s", uid.GoDicomImplementationClassUID, implClassUID)
	}

	// Check implementation version name
	implVersionName, ok := fmi.ImplementationVersionName()
	if !ok {
		t.Error("ImplementationVersionName not set")
	} else if implVersionName != uid.GoDicomImplementationVersionName {
		t.Errorf("Expected ImplementationVersionName %s, got %s", uid.GoDicomImplementationVersionName, implVersionName)
	}
}

func TestFileMetaInformationSetAndGet(t *testing.T) {
	fmi := NewFileMetaInformation()

	// Test MediaStorageSOPClassUID
	testUID := "1.2.840.10008.5.1.4.1.1.2"
	fmi.SetMediaStorageSOPClassUID(testUID)
	retrievedUID, ok := fmi.MediaStorageSOPClassUID()
	if !ok {
		t.Error("MediaStorageSOPClassUID not retrieved")
	} else if retrievedUID != testUID {
		t.Errorf("Expected MediaStorageSOPClassUID %s, got %s", testUID, retrievedUID)
	}

	// Test MediaStorageSOPInstanceUID
	testInstanceUID := "1.2.3.4.5"
	fmi.SetMediaStorageSOPInstanceUID(testInstanceUID)
	retrievedInstanceUID, ok := fmi.MediaStorageSOPInstanceUID()
	if !ok {
		t.Error("MediaStorageSOPInstanceUID not retrieved")
	} else if retrievedInstanceUID != testInstanceUID {
		t.Errorf("Expected MediaStorageSOPInstanceUID %s, got %s", testInstanceUID, retrievedInstanceUID)
	}

	// Test TransferSyntaxUID
	testTSUID := transfer.ExplicitVRLittleEndian.UID().UID()
	fmi.SetTransferSyntaxUID(testTSUID)
	retrievedTSUID, ok := fmi.TransferSyntaxUID()
	if !ok {
		t.Error("TransferSyntaxUID not retrieved")
	} else if retrievedTSUID != testTSUID {
		t.Errorf("Expected TransferSyntaxUID %s, got %s", testTSUID, retrievedTSUID)
	}
}

func TestFileMetaInformationTransferSyntax(t *testing.T) {
	fmi := NewFileMetaInformation()
	fmi.SetTransferSyntax(transfer.ExplicitVRLittleEndian)

	ts, ok := fmi.TransferSyntax()
	if !ok {
		t.Fatal("TransferSyntax not retrieved")
	}
	if ts.UID().UID() != transfer.ExplicitVRLittleEndian.UID().UID() {
		t.Errorf("Expected transfer syntax %s, got %s",
			transfer.ExplicitVRLittleEndian.UID().UID(),
			ts.UID().UID())
	}
}

func TestFileMetaInformationAETitles(t *testing.T) {
	fmi := NewFileMetaInformation()

	// Test SourceApplicationEntityTitle
	sourceAET := "SOURCE_AET"
	fmi.SetSourceApplicationEntityTitle(sourceAET)
	retrieved, ok := fmi.SourceApplicationEntityTitle()
	if !ok {
		t.Error("SourceApplicationEntityTitle not retrieved")
	} else if retrieved != sourceAET {
		t.Errorf("Expected SourceApplicationEntityTitle %s, got %s", sourceAET, retrieved)
	}

	// Test SendingApplicationEntityTitle
	sendingAET := "SENDING_AET"
	fmi.SetSendingApplicationEntityTitle(sendingAET)
	retrieved, ok = fmi.SendingApplicationEntityTitle()
	if !ok {
		t.Error("SendingApplicationEntityTitle not retrieved")
	} else if retrieved != sendingAET {
		t.Errorf("Expected SendingApplicationEntityTitle %s, got %s", sendingAET, retrieved)
	}

	// Test ReceivingApplicationEntityTitle
	receivingAET := "RECV_AET"
	fmi.SetReceivingApplicationEntityTitle(receivingAET)
	retrieved, ok = fmi.ReceivingApplicationEntityTitle()
	if !ok {
		t.Error("ReceivingApplicationEntityTitle not retrieved")
	} else if retrieved != receivingAET {
		t.Errorf("Expected ReceivingApplicationEntityTitle %s, got %s", receivingAET, retrieved)
	}
}

func TestFileMetaInformationPrivateInformation(t *testing.T) {
	fmi := NewFileMetaInformation()

	// Test PrivateInformationCreatorUID
	privUID := "1.2.3.4.5.6"
	fmi.SetPrivateInformationCreatorUID(privUID)
	retrieved, ok := fmi.PrivateInformationCreatorUID()
	if !ok {
		t.Error("PrivateInformationCreatorUID not retrieved")
	} else if retrieved != privUID {
		t.Errorf("Expected PrivateInformationCreatorUID %s, got %s", privUID, retrieved)
	}

	// Test PrivateInformation
	privData := []byte{0x01, 0x02, 0x03, 0x04}
	fmi.SetPrivateInformation(privData)
	retrievedData, ok := fmi.PrivateInformation()
	if !ok {
		t.Error("PrivateInformation not retrieved")
	} else {
		if len(retrievedData) != len(privData) {
			t.Errorf("Expected PrivateInformation length %d, got %d", len(privData), len(retrievedData))
		}
		for i := range privData {
			if retrievedData[i] != privData[i] {
				t.Errorf("PrivateInformation byte %d: expected 0x%02X, got 0x%02X", i, privData[i], retrievedData[i])
			}
		}
	}
}

func TestFileMetaInformationValidate(t *testing.T) {
	// Test valid FileMetaInformation
	fmi := NewDefaultFileMetaInformation()
	fmi.SetMediaStorageSOPClassUID("1.2.840.10008.5.1.4.1.1.2")
	fmi.SetMediaStorageSOPInstanceUID("1.2.3.4.5")
	fmi.SetTransferSyntaxUID(transfer.ExplicitVRLittleEndian.UID().UID())

	if err := fmi.Validate(); err != nil {
		t.Errorf("Validate failed for valid FileMetaInformation: %v", err)
	}

	// Test missing required element
	fmi2 := NewFileMetaInformation()
	fmi2.SetVersion([]byte{0x00, 0x01})
	if err := fmi2.Validate(); err == nil {
		t.Error("Validate should fail for FileMetaInformation missing required elements")
	}

	// Test invalid group
	fmi3 := NewDefaultFileMetaInformation()
	fmi3.SetMediaStorageSOPClassUID("1.2.840.10008.5.1.4.1.1.2")
	fmi3.SetMediaStorageSOPInstanceUID("1.2.3.4.5")
	fmi3.SetTransferSyntaxUID(transfer.ExplicitVRLittleEndian.UID().UID())
	// Add element from wrong group
	fmi3.dataset.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test"}))
	if err := fmi3.Validate(); err == nil {
		t.Error("Validate should fail for FileMetaInformation with non-0002 group tag")
	}
}

func TestNewFileMetaInformationFromMainDataset(t *testing.T) {
	// Create main dataset
	ds := New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
	ds.Add(element.NewString(tag.SourceApplicationEntityTitle, vr.AE, []string{"SOURCE"}))

	// Create FileMetaInformation from dataset
	fmi, err := NewFileMetaInformationFromMainDataset(ds, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("NewFileMetaInformationFromMainDataset failed: %v", err)
	}

	// Verify MediaStorageSOPClassUID
	sopClassUID, ok := fmi.MediaStorageSOPClassUID()
	if !ok {
		t.Error("MediaStorageSOPClassUID not set")
	} else if sopClassUID != "1.2.840.10008.5.1.4.1.1.2" {
		t.Errorf("Unexpected MediaStorageSOPClassUID: %s", sopClassUID)
	}

	// Verify MediaStorageSOPInstanceUID
	sopInstanceUID, ok := fmi.MediaStorageSOPInstanceUID()
	if !ok {
		t.Error("MediaStorageSOPInstanceUID not set")
	} else if sopInstanceUID != "1.2.3.4.5" {
		t.Errorf("Unexpected MediaStorageSOPInstanceUID: %s", sopInstanceUID)
	}

	// Verify TransferSyntaxUID
	tsUID, ok := fmi.TransferSyntaxUID()
	if !ok {
		t.Error("TransferSyntaxUID not set")
	} else if tsUID != transfer.ExplicitVRLittleEndian.UID().UID() {
		t.Errorf("Unexpected TransferSyntaxUID: %s", tsUID)
	}

	// Verify SourceApplicationEntityTitle was copied
	sourceAET, ok := fmi.SourceApplicationEntityTitle()
	if !ok {
		t.Error("SourceApplicationEntityTitle not copied")
	} else if sourceAET != "SOURCE" {
		t.Errorf("Unexpected SourceApplicationEntityTitle: %s", sourceAET)
	}
}

func TestNewFileMetaInformationFromMainDatasetMissingUID(t *testing.T) {
	// Create main dataset without SOPClassUID
	ds := New()
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

	_, err := NewFileMetaInformationFromMainDataset(ds, transfer.ExplicitVRLittleEndian)
	if err == nil {
		t.Error("NewFileMetaInformationFromMainDataset should fail without SOPClassUID")
	}

	// Create main dataset without SOPInstanceUID
	ds2 := New()
	ds2.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))

	_, err = NewFileMetaInformationFromMainDataset(ds2, transfer.ExplicitVRLittleEndian)
	if err == nil {
		t.Error("NewFileMetaInformationFromMainDataset should fail without SOPInstanceUID")
	}
}

func TestCreateDefaultSourceAET(t *testing.T) {
	aet := CreateDefaultSourceAET()
	if aet == "" {
		t.Error("CreateDefaultSourceAET returned empty string")
	}
	if len(aet) > 16 {
		t.Errorf("CreateDefaultSourceAET returned AET longer than 16 characters: %s", aet)
	}
}

func TestFileMetaInformationDataset(t *testing.T) {
	fmi := NewFileMetaInformation()
	fmi.SetMediaStorageSOPClassUID("1.2.3")

	ds := fmi.Dataset()
	if ds == nil {
		t.Fatal("Dataset() returned nil")
	}

	// Verify we can access the dataset directly
	sopClassUID, ok := ds.GetString(tag.MediaStorageSOPClassUID)
	if !ok {
		t.Error("Could not get MediaStorageSOPClassUID from dataset")
	} else if sopClassUID != "1.2.3" {
		t.Errorf("Expected 1.2.3, got %s", sopClassUID)
	}
}
