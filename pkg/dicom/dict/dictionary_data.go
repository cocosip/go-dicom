// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DICOM Dictionary.xml (version 2026b). DO NOT EDIT.

package dict

import (
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vm"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// loadStandardEntries loads all standard DICOM dictionary entries.
func loadStandardEntries(d *Dictionary) {
	d.Add(NewEntry(
		tag.New(0x0000, 0x0000),
		"Command Group Length",
		"CommandGroupLength",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0002),
		"Affected SOP Class UID",
		"AffectedSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0003),
		"Requested SOP Class UID",
		"RequestedSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0100),
		"Command Field",
		"CommandField",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0110),
		"Message ID",
		"MessageID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0120),
		"Message ID Being Responded To",
		"MessageIDBeingRespondedTo",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0600),
		"Move Destination",
		"MoveDestination",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0700),
		"Priority",
		"Priority",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0800),
		"Command Data Set Type",
		"CommandDataSetType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0900),
		"Status",
		"Status",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0901),
		"Offending Element",
		"OffendingElement",
		vm.VM1N,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0902),
		"Error Comment",
		"ErrorComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0903),
		"Error ID",
		"ErrorID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1000),
		"Affected SOP Instance UID",
		"AffectedSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1001),
		"Requested SOP Instance UID",
		"RequestedSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1002),
		"Event Type ID",
		"EventTypeID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1005),
		"Attribute Identifier List",
		"AttributeIdentifierList",
		vm.VM1N,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1008),
		"Action Type ID",
		"ActionTypeID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1020),
		"Number of Remaining Sub-operations",
		"NumberOfRemainingSuboperations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1021),
		"Number of Completed Sub-operations",
		"NumberOfCompletedSuboperations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1022),
		"Number of Failed Sub-operations",
		"NumberOfFailedSuboperations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1023),
		"Number of Warning Sub-operations",
		"NumberOfWarningSuboperations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1030),
		"Move Originator Application Entity Title",
		"MoveOriginatorApplicationEntityTitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1031),
		"Move Originator Message ID",
		"MoveOriginatorMessageID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0001),
		"Command Length to End",
		"CommandLengthToEnd",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0010),
		"Command Recognition Code",
		"CommandRecognitionCode",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0200),
		"Initiator",
		"Initiator",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0300),
		"Receiver",
		"Receiver",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0400),
		"Find Location",
		"FindLocation",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0850),
		"Number of Matches",
		"NumberOfMatches",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0860),
		"Response Sequence Number",
		"ResponseSequenceNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x4000),
		"Dialog Receiver",
		"DialogReceiver",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x4010),
		"Terminal Type",
		"TerminalType",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5010),
		"Message Set ID",
		"MessageSetID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5020),
		"End Message ID",
		"EndMessageID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5110),
		"Display Format",
		"DisplayFormat",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5120),
		"Page Position ID",
		"PagePositionID",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5130),
		"Text Format ID",
		"TextFormatID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5140),
		"Normal/Reverse",
		"NormalReverse",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5150),
		"Add Gray Scale",
		"AddGrayScale",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5160),
		"Borders",
		"Borders",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5170),
		"Copies",
		"Copies",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5180),
		"Command Magnification Type",
		"CommandMagnificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5190),
		"Erase",
		"Erase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x51A0),
		"Print",
		"Print",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x51B0),
		"Overlays",
		"Overlays",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0000),
		"File Meta Information Group Length",
		"FileMetaInformationGroupLength",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0001),
		"File Meta Information Version",
		"FileMetaInformationVersion",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0002),
		"Media Storage SOP Class UID",
		"MediaStorageSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0003),
		"Media Storage SOP Instance UID",
		"MediaStorageSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0010),
		"Transfer Syntax UID",
		"TransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0012),
		"Implementation Class UID",
		"ImplementationClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0013),
		"Implementation Version Name",
		"ImplementationVersionName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0016),
		"Source Application Entity Title",
		"SourceApplicationEntityTitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0017),
		"Sending Application Entity Title",
		"SendingApplicationEntityTitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0018),
		"Receiving Application Entity Title",
		"ReceivingApplicationEntityTitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0026),
		"Source Presentation Address",
		"SourcePresentationAddress",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0027),
		"Sending Presentation Address",
		"SendingPresentationAddress",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0028),
		"Receiving Presentation Address",
		"ReceivingPresentationAddress",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0031),
		"RTV Meta Information Version",
		"RTVMetaInformationVersion",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0032),
		"RTV Communication SOP Class UID",
		"RTVCommunicationSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0033),
		"RTV Communication SOP Instance UID",
		"RTVCommunicationSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0035),
		"RTV Source Identifier",
		"RTVSourceIdentifier",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0036),
		"RTV Flow Identifier",
		"RTVFlowIdentifier",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0037),
		"RTV Flow RTP Sampling Rate",
		"RTVFlowRTPSamplingRate",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0038),
		"RTV Flow Actual Frame Duration",
		"RTVFlowActualFrameDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0100),
		"Private Information Creator UID",
		"PrivateInformationCreatorUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0102),
		"Private Information",
		"PrivateInformation",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1130),
		"File-set ID",
		"FileSetID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1141),
		"File-set Descriptor File ID",
		"FileSetDescriptorFileID",
		vm.VM18,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1142),
		"Specific Character Set of File-set Descriptor File",
		"SpecificCharacterSetOfFileSetDescriptorFile",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1200),
		"Offset of the First Directory Record of the Root Directory Entity",
		"OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1202),
		"Offset of the Last Directory Record of the Root Directory Entity",
		"OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1212),
		"File-set Consistency Flag",
		"FileSetConsistencyFlag",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1220),
		"Directory Record Sequence",
		"DirectoryRecordSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1400),
		"Offset of the Next Directory Record",
		"OffsetOfTheNextDirectoryRecord",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1410),
		"Record In-use Flag",
		"RecordInUseFlag",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1420),
		"Offset of Referenced Lower-Level Directory Entity",
		"OffsetOfReferencedLowerLevelDirectoryEntity",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1430),
		"Directory Record Type",
		"DirectoryRecordType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1432),
		"Private Record UID",
		"PrivateRecordUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1500),
		"Referenced File ID",
		"ReferencedFileID",
		vm.VM18,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1504),
		"MRDR Directory Record Offset",
		"MRDRDirectoryRecordOffset",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1510),
		"Referenced SOP Class UID in File",
		"ReferencedSOPClassUIDInFile",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1511),
		"Referenced SOP Instance UID in File",
		"ReferencedSOPInstanceUIDInFile",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1512),
		"Referenced Transfer Syntax UID in File",
		"ReferencedTransferSyntaxUIDInFile",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x151A),
		"Referenced Related General SOP Class UID in File",
		"ReferencedRelatedGeneralSOPClassUIDInFile",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1600),
		"Number of References",
		"NumberOfReferences",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0001),
		"Length to End",
		"LengthToEnd",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0005),
		"Specific Character Set",
		"SpecificCharacterSet",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0006),
		"Language Code Sequence",
		"LanguageCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0008),
		"Image Type",
		"ImageType",
		vm.VM2N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0010),
		"Recognition Code",
		"RecognitionCode",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0012),
		"Instance Creation Date",
		"InstanceCreationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0013),
		"Instance Creation Time",
		"InstanceCreationTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0014),
		"Instance Creator UID",
		"InstanceCreatorUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0015),
		"Instance Coercion DateTime",
		"InstanceCoercionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0016),
		"SOP Class UID",
		"SOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0017),
		"Acquisition UID",
		"AcquisitionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0018),
		"SOP Instance UID",
		"SOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0019),
		"Pyramid UID",
		"PyramidUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x001A),
		"Related General SOP Class UID",
		"RelatedGeneralSOPClassUID",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x001B),
		"Original Specialized SOP Class UID",
		"OriginalSpecializedSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x001C),
		"Synthetic Data",
		"SyntheticData",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x001D),
		"Sensitive Content Code Sequence",
		"SensitiveContentCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0020),
		"Study Date",
		"StudyDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0021),
		"Series Date",
		"SeriesDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0022),
		"Acquisition Date",
		"AcquisitionDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0023),
		"Content Date",
		"ContentDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0024),
		"Overlay Date",
		"OverlayDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0025),
		"Curve Date",
		"CurveDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x002A),
		"Acquisition DateTime",
		"AcquisitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0030),
		"Study Time",
		"StudyTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0031),
		"Series Time",
		"SeriesTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0032),
		"Acquisition Time",
		"AcquisitionTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0033),
		"Content Time",
		"ContentTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0034),
		"Overlay Time",
		"OverlayTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0035),
		"Curve Time",
		"CurveTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0040),
		"Data Set Type",
		"DataSetType",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0041),
		"Data Set Subtype",
		"DataSetSubtype",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0042),
		"Nuclear Medicine Series Type",
		"NuclearMedicineSeriesType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0050),
		"Accession Number",
		"AccessionNumber",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0051),
		"Issuer of Accession Number Sequence",
		"IssuerOfAccessionNumberSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0052),
		"Query/Retrieve Level",
		"QueryRetrieveLevel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0053),
		"Query/Retrieve View",
		"QueryRetrieveView",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0054),
		"Retrieve AE Title",
		"RetrieveAETitle",
		vm.VM1N,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0055),
		"Station AE Title",
		"StationAETitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0056),
		"Instance Availability",
		"InstanceAvailability",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0058),
		"Failed SOP Instance UID List",
		"FailedSOPInstanceUIDList",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0060),
		"Modality",
		"Modality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0061),
		"Modalities in Study",
		"ModalitiesInStudy",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0062),
		"SOP Classes in Study",
		"SOPClassesInStudy",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0063),
		"Anatomic Regions in Study Code Sequence",
		"AnatomicRegionsInStudyCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0064),
		"Conversion Type",
		"ConversionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0068),
		"Presentation Intent Type",
		"PresentationIntentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0070),
		"Manufacturer",
		"Manufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0080),
		"Institution Name",
		"InstitutionName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0081),
		"Institution Address",
		"InstitutionAddress",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0082),
		"Institution Code Sequence",
		"InstitutionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0090),
		"Referring Physician's Name",
		"ReferringPhysicianName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0092),
		"Referring Physician's Address",
		"ReferringPhysicianAddress",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0094),
		"Referring Physician's Telephone Numbers",
		"ReferringPhysicianTelephoneNumbers",
		vm.VM1N,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0096),
		"Referring Physician Identification Sequence",
		"ReferringPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x009C),
		"Consulting Physician's Name",
		"ConsultingPhysicianName",
		vm.VM1N,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x009D),
		"Consulting Physician Identification Sequence",
		"ConsultingPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0100),
		"Code Value",
		"CodeValue",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0101),
		"Extended Code Value",
		"ExtendedCodeValue",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0102),
		"Coding Scheme Designator",
		"CodingSchemeDesignator",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0103),
		"Coding Scheme Version",
		"CodingSchemeVersion",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0104),
		"Code Meaning",
		"CodeMeaning",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0105),
		"Mapping Resource",
		"MappingResource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0106),
		"Context Group Version",
		"ContextGroupVersion",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0107),
		"Context Group Local Version",
		"ContextGroupLocalVersion",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0108),
		"Extended Code Meaning",
		"ExtendedCodeMeaning",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0109),
		"Coding Scheme Resources Sequence",
		"CodingSchemeResourcesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010A),
		"Coding Scheme URL Type",
		"CodingSchemeURLType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010B),
		"Context Group Extension Flag",
		"ContextGroupExtensionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010C),
		"Coding Scheme UID",
		"CodingSchemeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010D),
		"Context Group Extension Creator UID",
		"ContextGroupExtensionCreatorUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010E),
		"Coding Scheme URL",
		"CodingSchemeURL",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010F),
		"Context Identifier",
		"ContextIdentifier",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0110),
		"Coding Scheme Identification Sequence",
		"CodingSchemeIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0112),
		"Coding Scheme Registry",
		"CodingSchemeRegistry",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0114),
		"Coding Scheme External ID",
		"CodingSchemeExternalID",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0115),
		"Coding Scheme Name",
		"CodingSchemeName",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0116),
		"Coding Scheme Responsible Organization",
		"CodingSchemeResponsibleOrganization",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0117),
		"Context UID",
		"ContextUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0118),
		"Mapping Resource UID",
		"MappingResourceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0119),
		"Long Code Value",
		"LongCodeValue",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0120),
		"URN Code Value",
		"URNCodeValue",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0121),
		"Equivalent Code Sequence",
		"EquivalentCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0122),
		"Mapping Resource Name",
		"MappingResourceName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0123),
		"Context Group Identification Sequence",
		"ContextGroupIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0124),
		"Mapping Resource Identification Sequence",
		"MappingResourceIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0201),
		"Timezone Offset From UTC",
		"TimezoneOffsetFromUTC",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0220),
		"Responsible Group Code Sequence",
		"ResponsibleGroupCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0221),
		"Equipment Modality",
		"EquipmentModality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0222),
		"Manufacturer's Related Model Group",
		"ManufacturerRelatedModelGroup",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0300),
		"Private Data Element Characteristics Sequence",
		"PrivateDataElementCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0301),
		"Private Group Reference",
		"PrivateGroupReference",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0302),
		"Private Creator Reference",
		"PrivateCreatorReference",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0303),
		"Block Identifying Information Status",
		"BlockIdentifyingInformationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0304),
		"Nonidentifying Private Elements",
		"NonidentifyingPrivateElements",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0306),
		"Identifying Private Elements",
		"IdentifyingPrivateElements",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0305),
		"Deidentification Action Sequence",
		"DeidentificationActionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0307),
		"Deidentification Action",
		"DeidentificationAction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0308),
		"Private Data Element",
		"PrivateDataElement",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0309),
		"Private Data Element Value Multiplicity",
		"PrivateDataElementValueMultiplicity",
		vm.VM13,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030A),
		"Private Data Element Value Representation",
		"PrivateDataElementValueRepresentation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030B),
		"Private Data Element Number of Items",
		"PrivateDataElementNumberOfItems",
		vm.VM12,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030C),
		"Private Data Element Name",
		"PrivateDataElementName",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030D),
		"Private Data Element Keyword",
		"PrivateDataElementKeyword",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030E),
		"Private Data Element Description",
		"PrivateDataElementDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030F),
		"Private Data Element Encoding",
		"PrivateDataElementEncoding",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0310),
		"Private Data Element Definition Sequence",
		"PrivateDataElementDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0400),
		"Scope of Inventory Sequence",
		"ScopeOfInventorySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0401),
		"Inventory Purpose",
		"InventoryPurpose",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0402),
		"Inventory Instance Description",
		"InventoryInstanceDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0403),
		"Inventory Level",
		"InventoryLevel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0404),
		"Item Inventory DateTime",
		"ItemInventoryDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0405),
		"Removed from Operational Use",
		"RemovedFromOperationalUse",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0406),
		"Reason for Removal Code Sequence",
		"ReasonForRemovalCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0407),
		"Stored Instance Base URI",
		"StoredInstanceBaseURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0408),
		"Folder Access URI",
		"FolderAccessURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0409),
		"File Access URI",
		"FileAccessURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040A),
		"Container File Type",
		"ContainerFileType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040B),
		"Filename in Container",
		"FilenameInContainer",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040C),
		"File Offset in Container",
		"FileOffsetInContainer",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040D),
		"File Length in Container",
		"FileLengthInContainer",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040E),
		"Stored Instance Transfer Syntax UID",
		"StoredInstanceTransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040F),
		"Extended Matching Mechanisms",
		"ExtendedMatchingMechanisms",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0410),
		"Range Matching Sequence",
		"RangeMatchingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0411),
		"List of UID Matching Sequence",
		"ListOfUIDMatchingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0412),
		"Empty Value Matching Sequence",
		"EmptyValueMatchingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0413),
		"General Matching Sequence",
		"GeneralMatchingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0414),
		"Requested Status Interval",
		"RequestedStatusInterval",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0415),
		"Retain Instances",
		"RetainInstances",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0416),
		"Expiration DateTime",
		"ExpirationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0417),
		"Transaction Status",
		"TransactionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0418),
		"Transaction Status Comment",
		"TransactionStatusComment",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0419),
		"File Set Access Sequence",
		"FileSetAccessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041A),
		"File Access Sequence",
		"FileAccessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041B),
		"Record Key",
		"RecordKey",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041C),
		"Prior Record Key",
		"PriorRecordKey",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041D),
		"Metadata Sequence",
		"MetadataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041E),
		"Updated Metadata Sequence",
		"UpdatedMetadataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041F),
		"Study Update DateTime",
		"StudyUpdateDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0420),
		"Inventory Access End Points Sequence",
		"InventoryAccessEndPointsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0421),
		"Study Access End Points Sequence",
		"StudyAccessEndPointsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0422),
		"Incorporated Inventory Instance Sequence",
		"IncorporatedInventoryInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0423),
		"Inventoried Studies Sequence",
		"InventoriedStudiesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0424),
		"Inventoried Series Sequence",
		"InventoriedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0425),
		"Inventoried Instances Sequence",
		"InventoriedInstancesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0426),
		"Inventory Completion Status",
		"InventoryCompletionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0427),
		"Number of Study Records in Instance",
		"NumberOfStudyRecordsInInstance",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0428),
		"Total Number of Study Records",
		"TotalNumberOfStudyRecords",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0429),
		"Maximum Number of Records",
		"MaximumNumberOfRecords",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1000),
		"Network ID",
		"NetworkID",
		vm.VM1,
		true,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1010),
		"Station Name",
		"StationName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1030),
		"Study Description",
		"StudyDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1032),
		"Procedure Code Sequence",
		"ProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x103E),
		"Series Description",
		"SeriesDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x103F),
		"Series Description Code Sequence",
		"SeriesDescriptionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1040),
		"Institutional Department Name",
		"InstitutionalDepartmentName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1041),
		"Institutional Department Type Code Sequence",
		"InstitutionalDepartmentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1048),
		"Physician(s) of Record",
		"PhysiciansOfRecord",
		vm.VM1N,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1049),
		"Physician(s) of Record Identification Sequence",
		"PhysiciansOfRecordIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1050),
		"Performing Physician's Name",
		"PerformingPhysicianName",
		vm.VM1N,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1052),
		"Performing Physician Identification Sequence",
		"PerformingPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1060),
		"Name of Physician(s) Reading Study",
		"NameOfPhysiciansReadingStudy",
		vm.VM1N,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1062),
		"Physician(s) Reading Study Identification Sequence",
		"PhysiciansReadingStudyIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1070),
		"Operators' Name",
		"OperatorsName",
		vm.VM1N,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1072),
		"Operator Identification Sequence",
		"OperatorIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1080),
		"Admitting Diagnoses Description",
		"AdmittingDiagnosesDescription",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1084),
		"Admitting Diagnoses Code Sequence",
		"AdmittingDiagnosesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1088),
		"Pyramid Description",
		"PyramidDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1090),
		"Manufacturer's Model Name",
		"ManufacturerModelName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1100),
		"Referenced Results Sequence",
		"ReferencedResultsSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1110),
		"Referenced Study Sequence",
		"ReferencedStudySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1111),
		"Referenced Performed Procedure Step Sequence",
		"ReferencedPerformedProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1112),
		"Referenced Instances by SOP Class Sequence",
		"ReferencedInstancesBySOPClassSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1115),
		"Referenced Series Sequence",
		"ReferencedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1120),
		"Referenced Patient Sequence",
		"ReferencedPatientSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1125),
		"Referenced Visit Sequence",
		"ReferencedVisitSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1130),
		"Referenced Overlay Sequence",
		"ReferencedOverlaySequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1134),
		"Referenced Stereometric Instance Sequence",
		"ReferencedStereometricInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x113A),
		"Referenced Waveform Sequence",
		"ReferencedWaveformSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1140),
		"Referenced Image Sequence",
		"ReferencedImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1145),
		"Referenced Curve Sequence",
		"ReferencedCurveSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x114A),
		"Referenced Instance Sequence",
		"ReferencedInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x114B),
		"Referenced Real World Value Mapping Instance Sequence",
		"ReferencedRealWorldValueMappingInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x114C),
		"Referenced Segmentation Sequence",
		"ReferencedSegmentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x114D),
		"Referenced Surface Segmentation Sequence",
		"ReferencedSurfaceSegmentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1150),
		"Referenced SOP Class UID",
		"ReferencedSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1155),
		"Referenced SOP Instance UID",
		"ReferencedSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1156),
		"Definition Source Sequence",
		"DefinitionSourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x115A),
		"SOP Classes Supported",
		"SOPClassesSupported",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1160),
		"Referenced Frame Number",
		"ReferencedFrameNumber",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1161),
		"Simple Frame List",
		"SimpleFrameList",
		vm.VM1N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1162),
		"Calculated Frame List",
		"CalculatedFrameList",
		vm.VM33N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1163),
		"Time Range",
		"TimeRange",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1164),
		"Frame Extraction Sequence",
		"FrameExtractionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1167),
		"Multi-frame Source SOP Instance UID",
		"MultiFrameSourceSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1190),
		"Retrieve URL",
		"RetrieveURL",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1195),
		"Transaction UID",
		"TransactionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1196),
		"Warning Reason",
		"WarningReason",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1197),
		"Failure Reason",
		"FailureReason",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1198),
		"Failed SOP Sequence",
		"FailedSOPSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1199),
		"Referenced SOP Sequence",
		"ReferencedSOPSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x119A),
		"Other Failures Sequence",
		"OtherFailuresSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x119B),
		"Failed Study Sequence",
		"FailedStudySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1200),
		"Studies Containing Other Referenced Instances Sequence",
		"StudiesContainingOtherReferencedInstancesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1250),
		"Related Series Sequence",
		"RelatedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1301),
		"Principal Diagnosis Code Sequence",
		"PrincipalDiagnosisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1302),
		"Primary Diagnosis Code Sequence",
		"PrimaryDiagnosisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1303),
		"Secondary Diagnoses Code Sequence",
		"SecondaryDiagnosesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1304),
		"Histological Diagnoses Code Sequence",
		"HistologicalDiagnosesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2110),
		"Lossy Image Compression (Retired)",
		"LossyImageCompressionRetired",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2111),
		"Derivation Description",
		"DerivationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2112),
		"Source Image Sequence",
		"SourceImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2120),
		"Stage Name",
		"StageName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2122),
		"Stage Number",
		"StageNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2124),
		"Number of Stages",
		"NumberOfStages",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2127),
		"View Name",
		"ViewName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2128),
		"View Number",
		"ViewNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2129),
		"Number of Event Timers",
		"NumberOfEventTimers",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x212A),
		"Number of Views in Stage",
		"NumberOfViewsInStage",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2130),
		"Event Elapsed Time(s)",
		"EventElapsedTimes",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2132),
		"Event Timer Name(s)",
		"EventTimerNames",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2133),
		"Event Timer Sequence",
		"EventTimerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2134),
		"Event Time Offset",
		"EventTimeOffset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2135),
		"Event Code Sequence",
		"EventCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2142),
		"Start Trim",
		"StartTrim",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2143),
		"Stop Trim",
		"StopTrim",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2144),
		"Recommended Display Frame Rate",
		"RecommendedDisplayFrameRate",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2200),
		"Transducer Position",
		"TransducerPosition",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2204),
		"Transducer Orientation",
		"TransducerOrientation",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2208),
		"Anatomic Structure",
		"AnatomicStructure",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2218),
		"Anatomic Region Sequence",
		"AnatomicRegionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2220),
		"Anatomic Region Modifier Sequence",
		"AnatomicRegionModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2228),
		"Primary Anatomic Structure Sequence",
		"PrimaryAnatomicStructureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2229),
		"Anatomic Structure, Space or Region Sequence",
		"AnatomicStructureSpaceOrRegionSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2230),
		"Primary Anatomic Structure Modifier Sequence",
		"PrimaryAnatomicStructureModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2240),
		"Transducer Position Sequence",
		"TransducerPositionSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2242),
		"Transducer Position Modifier Sequence",
		"TransducerPositionModifierSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2244),
		"Transducer Orientation Sequence",
		"TransducerOrientationSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2246),
		"Transducer Orientation Modifier Sequence",
		"TransducerOrientationModifierSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2251),
		"Anatomic Structure Space Or Region Code Sequence (Trial)",
		"AnatomicStructureSpaceOrRegionCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2253),
		"Anatomic Portal Of Entrance Code Sequence (Trial)",
		"AnatomicPortalOfEntranceCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2255),
		"Anatomic Approach Direction Code Sequence (Trial)",
		"AnatomicApproachDirectionCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2256),
		"Anatomic Perspective Description (Trial)",
		"AnatomicPerspectiveDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2257),
		"Anatomic Perspective Code Sequence (Trial)",
		"AnatomicPerspectiveCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2258),
		"Anatomic Location Of Examining Instrument Description (Trial)",
		"AnatomicLocationOfExaminingInstrumentDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2259),
		"Anatomic Location Of Examining Instrument Code Sequence (Trial)",
		"AnatomicLocationOfExaminingInstrumentCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x225A),
		"Anatomic Structure Space Or Region Modifier Code Sequence (Trial)",
		"AnatomicStructureSpaceOrRegionModifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x225C),
		"On Axis Background Anatomic Structure Code Sequence (Trial)",
		"OnAxisBackgroundAnatomicStructureCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3001),
		"Alternate Representation Sequence",
		"AlternateRepresentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3002),
		"Available Transfer Syntax UID",
		"AvailableTransferSyntaxUID",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3010),
		"Irradiation Event UID",
		"IrradiationEventUID",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3011),
		"Source Irradiation Event Sequence",
		"SourceIrradiationEventSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3012),
		"Radiopharmaceutical Administration Event UID",
		"RadiopharmaceuticalAdministrationEventUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x4000),
		"Identifying Comments",
		"IdentifyingComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9007),
		"Frame Type",
		"FrameType",
		vm.MustParse("4-5"),
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9092),
		"Referenced Image Evidence Sequence",
		"ReferencedImageEvidenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9121),
		"Referenced Raw Data Sequence",
		"ReferencedRawDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9123),
		"Creator-Version UID",
		"CreatorVersionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9124),
		"Derivation Image Sequence",
		"DerivationImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9154),
		"Source Image Evidence Sequence",
		"SourceImageEvidenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9205),
		"Pixel Presentation",
		"PixelPresentation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9206),
		"Volumetric Properties",
		"VolumetricProperties",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9207),
		"Volume Based Calculation Technique",
		"VolumeBasedCalculationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9208),
		"Complex Image Component",
		"ComplexImageComponent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9209),
		"Acquisition Contrast",
		"AcquisitionContrast",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9215),
		"Derivation Code Sequence",
		"DerivationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9237),
		"Referenced Presentation State Sequence",
		"ReferencedPresentationStateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9410),
		"Referenced Other Plane Sequence",
		"ReferencedOtherPlaneSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9458),
		"Frame Display Sequence",
		"FrameDisplaySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9459),
		"Recommended Display Frame Rate in Float",
		"RecommendedDisplayFrameRateInFloat",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9460),
		"Skip Frame Range Flag",
		"SkipFrameRangeFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0010),
		"Patient's Name",
		"PatientName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0011),
		"Person Names to Use Sequence",
		"PersonNamesToUseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0012),
		"Name to Use",
		"NameToUse",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0013),
		"Name to Use Comment",
		"NameToUseComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0014),
		"Third Person Pronouns Sequence",
		"ThirdPersonPronounsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0015),
		"Pronoun Code Sequence",
		"PronounCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0016),
		"Pronoun Comment",
		"PronounComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0020),
		"Patient ID",
		"PatientID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0021),
		"Issuer of Patient ID",
		"IssuerOfPatientID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0022),
		"Type of Patient ID",
		"TypeOfPatientID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0024),
		"Issuer of Patient ID Qualifiers Sequence",
		"IssuerOfPatientIDQualifiersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0026),
		"Source Patient Group Identification Sequence",
		"SourcePatientGroupIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0027),
		"Group of Patients Identification Sequence",
		"GroupOfPatientsIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0028),
		"Subject Relative Position in Image",
		"SubjectRelativePositionInImage",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0030),
		"Patient's Birth Date",
		"PatientBirthDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0032),
		"Patient's Birth Time",
		"PatientBirthTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0033),
		"Patient's Birth Date in Alternative Calendar",
		"PatientBirthDateInAlternativeCalendar",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0034),
		"Patient's Death Date in Alternative Calendar",
		"PatientDeathDateInAlternativeCalendar",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0035),
		"Patient's Alternative Calendar",
		"PatientAlternativeCalendar",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0040),
		"Patient's Sex",
		"PatientSex",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0041),
		"Gender Identity Sequence",
		"GenderIdentitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0042),
		"Sex Parameters for Clinical Use Category Comment",
		"SexParametersForClinicalUseCategoryComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0043),
		"Sex Parameters for Clinical Use Category Sequence",
		"SexParametersForClinicalUseCategorySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0044),
		"Gender Identity Code Sequence",
		"GenderIdentityCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0045),
		"Gender Identity Comment",
		"GenderIdentityComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0046),
		"Sex Parameters for Clinical Use Category Code Sequence",
		"SexParametersForClinicalUseCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0047),
		"Sex Parameters for Clinical Use Category Reference",
		"SexParametersForClinicalUseCategoryReference",
		vm.VM1N,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0050),
		"Patient's Insurance Plan Code Sequence",
		"PatientInsurancePlanCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0101),
		"Patient's Primary Language Code Sequence",
		"PatientPrimaryLanguageCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0102),
		"Patient's Primary Language Modifier Code Sequence",
		"PatientPrimaryLanguageModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0200),
		"Quality Control Subject",
		"QualityControlSubject",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0201),
		"Quality Control Subject Type Code Sequence",
		"QualityControlSubjectTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0212),
		"Strain Description",
		"StrainDescription",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0213),
		"Strain Nomenclature",
		"StrainNomenclature",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0214),
		"Strain Stock Number",
		"StrainStockNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0215),
		"Strain Source Registry Code Sequence",
		"StrainSourceRegistryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0216),
		"Strain Stock Sequence",
		"StrainStockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0217),
		"Strain Source",
		"StrainSource",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0218),
		"Strain Additional Information",
		"StrainAdditionalInformation",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0219),
		"Strain Code Sequence",
		"StrainCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0221),
		"Genetic Modifications Sequence",
		"GeneticModificationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0222),
		"Genetic Modifications Description",
		"GeneticModificationsDescription",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0223),
		"Genetic Modifications Nomenclature",
		"GeneticModificationsNomenclature",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0229),
		"Genetic Modifications Code Sequence",
		"GeneticModificationsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1000),
		"Other Patient IDs",
		"OtherPatientIDs",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1001),
		"Other Patient Names",
		"OtherPatientNames",
		vm.VM1N,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1002),
		"Other Patient IDs Sequence",
		"OtherPatientIDsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1005),
		"Patient's Birth Name",
		"PatientBirthName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1010),
		"Patient's Age",
		"PatientAge",
		vm.VM1,
		false,
		vr.AS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1020),
		"Patient's Size",
		"PatientSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1021),
		"Patient's Size Code Sequence",
		"PatientSizeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1022),
		"Patient's Body Mass Index",
		"PatientBodyMassIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1023),
		"Measured AP Dimension",
		"MeasuredAPDimension",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1024),
		"Measured Lateral Dimension",
		"MeasuredLateralDimension",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1030),
		"Patient's Weight",
		"PatientWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1040),
		"Patient's Address",
		"PatientAddress",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1050),
		"Insurance Plan Identification",
		"InsurancePlanIdentification",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1060),
		"Patient's Mother's Birth Name",
		"PatientMotherBirthName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1080),
		"Military Rank",
		"MilitaryRank",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1081),
		"Branch of Service",
		"BranchOfService",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1090),
		"Medical Record Locator",
		"MedicalRecordLocator",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1100),
		"Referenced Patient Photo Sequence",
		"ReferencedPatientPhotoSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2000),
		"Medical Alerts",
		"MedicalAlerts",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2110),
		"Allergies",
		"Allergies",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2150),
		"Country of Residence",
		"CountryOfResidence",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2152),
		"Region of Residence",
		"RegionOfResidence",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2154),
		"Patient's Telephone Numbers",
		"PatientTelephoneNumbers",
		vm.VM1N,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2155),
		"Patient's Telecom Information",
		"PatientTelecomInformation",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2160),
		"Ethnic Group",
		"EthnicGroup",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2161),
		"Ethnic Group Code Sequence",
		"EthnicGroupCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2162),
		"Ethnic Groups",
		"EthnicGroups",
		vm.VM1N,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2180),
		"Occupation",
		"Occupation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21A0),
		"Smoking Status",
		"SmokingStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21B0),
		"Additional Patient History",
		"AdditionalPatientHistory",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21C0),
		"Pregnancy Status",
		"PregnancyStatus",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21D0),
		"Last Menstrual Date",
		"LastMenstrualDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21F0),
		"Patient's Religious Preference",
		"PatientReligiousPreference",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2201),
		"Patient Species Description",
		"PatientSpeciesDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2202),
		"Patient Species Code Sequence",
		"PatientSpeciesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2203),
		"Patient's Sex Neutered",
		"PatientSexNeutered",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2210),
		"Anatomical Orientation Type",
		"AnatomicalOrientationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2292),
		"Patient Breed Description",
		"PatientBreedDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2293),
		"Patient Breed Code Sequence",
		"PatientBreedCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2294),
		"Breed Registration Sequence",
		"BreedRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2295),
		"Breed Registration Number",
		"BreedRegistrationNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2296),
		"Breed Registry Code Sequence",
		"BreedRegistryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2297),
		"Responsible Person",
		"ResponsiblePerson",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2298),
		"Responsible Person Role",
		"ResponsiblePersonRole",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2299),
		"Responsible Organization",
		"ResponsibleOrganization",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x4000),
		"Patient Comments",
		"PatientComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x9431),
		"Examined Body Thickness",
		"ExaminedBodyThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0010),
		"Clinical Trial Sponsor Name",
		"ClinicalTrialSponsorName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0020),
		"Clinical Trial Protocol ID",
		"ClinicalTrialProtocolID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0021),
		"Clinical Trial Protocol Name",
		"ClinicalTrialProtocolName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0022),
		"Issuer of Clinical Trial Protocol ID",
		"IssuerOfClinicalTrialProtocolID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0023),
		"Other Clinical Trial Protocol IDs Sequence",
		"OtherClinicalTrialProtocolIDsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0030),
		"Clinical Trial Site ID",
		"ClinicalTrialSiteID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0031),
		"Clinical Trial Site Name",
		"ClinicalTrialSiteName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0032),
		"Issuer of Clinical Trial Site ID",
		"IssuerOfClinicalTrialSiteID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0040),
		"Clinical Trial Subject ID",
		"ClinicalTrialSubjectID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0041),
		"Issuer of Clinical Trial Subject ID",
		"IssuerOfClinicalTrialSubjectID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0042),
		"Clinical Trial Subject Reading ID",
		"ClinicalTrialSubjectReadingID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0043),
		"Issuer of Clinical Trial Subject Reading ID",
		"IssuerOfClinicalTrialSubjectReadingID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0050),
		"Clinical Trial Time Point ID",
		"ClinicalTrialTimePointID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0051),
		"Clinical Trial Time Point Description",
		"ClinicalTrialTimePointDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0052),
		"Longitudinal Temporal Offset from Event",
		"LongitudinalTemporalOffsetFromEvent",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0053),
		"Longitudinal Temporal Event Type",
		"LongitudinalTemporalEventType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0054),
		"Clinical Trial Time Point Type Code Sequence",
		"ClinicalTrialTimePointTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0055),
		"Issuer of Clinical Trial Time Point ID",
		"IssuerOfClinicalTrialTimePointID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0060),
		"Clinical Trial Coordinating Center Name",
		"ClinicalTrialCoordinatingCenterName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0062),
		"Patient Identity Removed",
		"PatientIdentityRemoved",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0063),
		"De-identification Method",
		"DeidentificationMethod",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0064),
		"De-identification Method Code Sequence",
		"DeidentificationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0071),
		"Clinical Trial Series ID",
		"ClinicalTrialSeriesID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0072),
		"Clinical Trial Series Description",
		"ClinicalTrialSeriesDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0073),
		"Issuer of Clinical Trial Series ID",
		"IssuerOfClinicalTrialSeriesID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0081),
		"Clinical Trial Protocol Ethics Committee Name",
		"ClinicalTrialProtocolEthicsCommitteeName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0082),
		"Clinical Trial Protocol Ethics Committee Approval Number",
		"ClinicalTrialProtocolEthicsCommitteeApprovalNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0083),
		"Consent for Clinical Trial Use Sequence",
		"ConsentForClinicalTrialUseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0084),
		"Distribution Type",
		"DistributionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0085),
		"Consent for Distribution Flag",
		"ConsentForDistributionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0086),
		"Ethics Committee Approval Effectiveness Start Date",
		"EthicsCommitteeApprovalEffectivenessStartDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0087),
		"Ethics Committee Approval Effectiveness End Date",
		"EthicsCommitteeApprovalEffectivenessEndDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0023),
		"CAD File Format",
		"CADFileFormat",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0024),
		"Component Reference System",
		"ComponentReferenceSystem",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0025),
		"Component Manufacturing Procedure",
		"ComponentManufacturingProcedure",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0028),
		"Component Manufacturer",
		"ComponentManufacturer",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0030),
		"Material Thickness",
		"MaterialThickness",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0032),
		"Material Pipe Diameter",
		"MaterialPipeDiameter",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0034),
		"Material Isolation Diameter",
		"MaterialIsolationDiameter",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0042),
		"Material Grade",
		"MaterialGrade",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0044),
		"Material Properties Description",
		"MaterialPropertiesDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0045),
		"Material Properties File Format (Retired)",
		"MaterialPropertiesFileFormatRetired",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0046),
		"Material Notes",
		"MaterialNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0050),
		"Component Shape",
		"ComponentShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0052),
		"Curvature Type",
		"CurvatureType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0054),
		"Outer Diameter",
		"OuterDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0056),
		"Inner Diameter",
		"InnerDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0100),
		"Component Welder IDs",
		"ComponentWelderIDs",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0101),
		"Secondary Approval Status",
		"SecondaryApprovalStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0102),
		"Secondary Review Date",
		"SecondaryReviewDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0103),
		"Secondary Review Time",
		"SecondaryReviewTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0104),
		"Secondary Reviewer Name",
		"SecondaryReviewerName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0105),
		"Repair ID",
		"RepairID",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0106),
		"Multiple Component Approval Sequence",
		"MultipleComponentApprovalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0107),
		"Other Approval Status",
		"OtherApprovalStatus",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0108),
		"Other Secondary Approval Status",
		"OtherSecondaryApprovalStatus",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0200),
		"Data Element Label Sequence",
		"DataElementLabelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0201),
		"Data Element Label Item Sequence",
		"DataElementLabelItemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0202),
		"Data Element",
		"DataElement",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0203),
		"Data Element Name",
		"DataElementName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0204),
		"Data Element Description",
		"DataElementDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0205),
		"Data Element Conditionality",
		"DataElementConditionality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0206),
		"Data Element Minimum Characters",
		"DataElementMinimumCharacters",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0207),
		"Data Element Maximum Characters",
		"DataElementMaximumCharacters",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x1010),
		"Actual Environmental Conditions",
		"ActualEnvironmentalConditions",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x1020),
		"Expiry Date",
		"ExpiryDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x1040),
		"Environmental Conditions",
		"EnvironmentalConditions",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2002),
		"Evaluator Sequence",
		"EvaluatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2004),
		"Evaluator Number",
		"EvaluatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2006),
		"Evaluator Name",
		"EvaluatorName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2008),
		"Evaluation Attempt",
		"EvaluationAttempt",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2012),
		"Indication Sequence",
		"IndicationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2014),
		"Indication Number",
		"IndicationNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2016),
		"Indication Label",
		"IndicationLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2018),
		"Indication Description",
		"IndicationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x201A),
		"Indication Type",
		"IndicationType",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x201C),
		"Indication Disposition",
		"IndicationDisposition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x201E),
		"Indication ROI Sequence",
		"IndicationROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2030),
		"Indication Physical Property Sequence",
		"IndicationPhysicalPropertySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2032),
		"Property Label",
		"PropertyLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2202),
		"Coordinate System Number of Axes",
		"CoordinateSystemNumberOfAxes",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2204),
		"Coordinate System Axes Sequence",
		"CoordinateSystemAxesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2206),
		"Coordinate System Axis Description",
		"CoordinateSystemAxisDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2208),
		"Coordinate System Data Set Mapping",
		"CoordinateSystemDataSetMapping",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x220A),
		"Coordinate System Axis Number",
		"CoordinateSystemAxisNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x220C),
		"Coordinate System Axis Type",
		"CoordinateSystemAxisType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x220E),
		"Coordinate System Axis Units",
		"CoordinateSystemAxisUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2210),
		"Coordinate System Axis Values",
		"CoordinateSystemAxisValues",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2220),
		"Coordinate System Transform Sequence",
		"CoordinateSystemTransformSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2222),
		"Transform Description",
		"TransformDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2224),
		"Transform Number of Axes",
		"TransformNumberOfAxes",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2226),
		"Transform Order of Axes",
		"TransformOrderOfAxes",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2228),
		"Transformed Axis Units",
		"TransformedAxisUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x222A),
		"Coordinate System Transform Rotation and Scale Matrix",
		"CoordinateSystemTransformRotationAndScaleMatrix",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x222C),
		"Coordinate System Transform Translation Matrix",
		"CoordinateSystemTransformTranslationMatrix",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3011),
		"Internal Detector Frame Time",
		"InternalDetectorFrameTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3012),
		"Number of Frames Integrated",
		"NumberOfFramesIntegrated",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3020),
		"Detector Temperature Sequence",
		"DetectorTemperatureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3022),
		"Sensor Name",
		"SensorName",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3024),
		"Horizontal Offset of Sensor",
		"HorizontalOffsetOfSensor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3026),
		"Vertical Offset of Sensor",
		"VerticalOffsetOfSensor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3028),
		"Sensor Temperature",
		"SensorTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3040),
		"Dark Current Sequence",
		"DarkCurrentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3050),
		"Dark Current Counts",
		"DarkCurrentCounts",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3060),
		"Gain Correction Reference Sequence",
		"GainCorrectionReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3070),
		"Air Counts",
		"AirCounts",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3071),
		"KV Used in Gain Calibration",
		"KVUsedInGainCalibration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3072),
		"MA Used in Gain Calibration",
		"MAUsedInGainCalibration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3073),
		"Number of Frames Used for Integration",
		"NumberOfFramesUsedForIntegration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3074),
		"Filter Material Used in Gain Calibration",
		"FilterMaterialUsedInGainCalibration",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3075),
		"Filter Thickness Used in Gain Calibration",
		"FilterThicknessUsedInGainCalibration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3076),
		"Date of Gain Calibration",
		"DateOfGainCalibration",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3077),
		"Time of Gain Calibration",
		"TimeOfGainCalibration",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3080),
		"Bad Pixel Image",
		"BadPixelImage",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3099),
		"Calibration Notes",
		"CalibrationNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3100),
		"Linearity Correction Technique",
		"LinearityCorrectionTechnique",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3101),
		"Beam Hardening Correction Technique",
		"BeamHardeningCorrectionTechnique",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4002),
		"Pulser Equipment Sequence",
		"PulserEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4004),
		"Pulser Type",
		"PulserType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4006),
		"Pulser Notes",
		"PulserNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4008),
		"Receiver Equipment Sequence",
		"ReceiverEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x400A),
		"Amplifier Type",
		"AmplifierType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x400C),
		"Receiver Notes",
		"ReceiverNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x400E),
		"Pre-Amplifier Equipment Sequence",
		"PreAmplifierEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x400F),
		"Pre-Amplifier Notes",
		"PreAmplifierNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4010),
		"Transmit Transducer Sequence",
		"TransmitTransducerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4011),
		"Receive Transducer Sequence",
		"ReceiveTransducerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4012),
		"Number of Elements",
		"NumberOfElements",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4013),
		"Element Shape",
		"ElementShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4014),
		"Element Dimension A",
		"ElementDimensionA",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4015),
		"Element Dimension B",
		"ElementDimensionB",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4016),
		"Element Pitch A",
		"ElementPitchA",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4017),
		"Measured Beam Dimension A",
		"MeasuredBeamDimensionA",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4018),
		"Measured Beam Dimension B",
		"MeasuredBeamDimensionB",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4019),
		"Location of Measured Beam Diameter",
		"LocationOfMeasuredBeamDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x401A),
		"Nominal Frequency",
		"NominalFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x401B),
		"Measured Center Frequency",
		"MeasuredCenterFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x401C),
		"Measured Bandwidth",
		"MeasuredBandwidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x401D),
		"Element Pitch B",
		"ElementPitchB",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4020),
		"Pulser Settings Sequence",
		"PulserSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4022),
		"Pulse Width",
		"PulseWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4024),
		"Excitation Frequency",
		"ExcitationFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4026),
		"Modulation Type",
		"ModulationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4028),
		"Damping",
		"Damping",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4030),
		"Receiver Settings Sequence",
		"ReceiverSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4031),
		"Acquired Soundpath Length",
		"AcquiredSoundpathLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4032),
		"Acquisition Compression Type",
		"AcquisitionCompressionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4033),
		"Acquisition Sample Size",
		"AcquisitionSampleSize",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4034),
		"Rectifier Smoothing",
		"RectifierSmoothing",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4035),
		"DAC Sequence",
		"DACSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4036),
		"DAC Type",
		"DACType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4038),
		"DAC Gain Points",
		"DACGainPoints",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x403A),
		"DAC Time Points",
		"DACTimePoints",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x403C),
		"DAC Amplitude",
		"DACAmplitude",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4040),
		"Pre-Amplifier Settings Sequence",
		"PreAmplifierSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4050),
		"Transmit Transducer Settings Sequence",
		"TransmitTransducerSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4051),
		"Receive Transducer Settings Sequence",
		"ReceiveTransducerSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4052),
		"Incident Angle",
		"IncidentAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4054),
		"Coupling Technique",
		"CouplingTechnique",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4056),
		"Coupling Medium",
		"CouplingMedium",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4057),
		"Coupling Velocity",
		"CouplingVelocity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4058),
		"Probe Center Location X",
		"ProbeCenterLocationX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4059),
		"Probe Center Location Z",
		"ProbeCenterLocationZ",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x405A),
		"Sound Path Length",
		"SoundPathLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x405C),
		"Delay Law Identifier",
		"DelayLawIdentifier",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4060),
		"Gate Settings Sequence",
		"GateSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4062),
		"Gate Threshold",
		"GateThreshold",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4064),
		"Velocity of Sound",
		"VelocityOfSound",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4070),
		"Calibration Settings Sequence",
		"CalibrationSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4072),
		"Calibration Procedure",
		"CalibrationProcedure",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4074),
		"Procedure Version",
		"ProcedureVersion",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4076),
		"Procedure Creation Date",
		"ProcedureCreationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4078),
		"Procedure Expiration Date",
		"ProcedureExpirationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x407A),
		"Procedure Last Modified Date",
		"ProcedureLastModifiedDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x407C),
		"Calibration Time",
		"CalibrationTime",
		vm.VM1N,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x407E),
		"Calibration Date",
		"CalibrationDate",
		vm.VM1N,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4080),
		"Probe Drive Equipment Sequence",
		"ProbeDriveEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4081),
		"Drive Type",
		"DriveType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4082),
		"Probe Drive Notes",
		"ProbeDriveNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4083),
		"Drive Probe Sequence",
		"DriveProbeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4084),
		"Probe Inductance",
		"ProbeInductance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4085),
		"Probe Resistance",
		"ProbeResistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4086),
		"Receive Probe Sequence",
		"ReceiveProbeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4087),
		"Probe Drive Settings Sequence",
		"ProbeDriveSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4088),
		"Bridge Resistors",
		"BridgeResistors",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4089),
		"Probe Orientation Angle",
		"ProbeOrientationAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x408B),
		"User Selected Gain Y",
		"UserSelectedGainY",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x408C),
		"User Selected Phase",
		"UserSelectedPhase",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x408D),
		"User Selected Offset X",
		"UserSelectedOffsetX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x408E),
		"User Selected Offset Y",
		"UserSelectedOffsetY",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4091),
		"Channel Settings Sequence",
		"ChannelSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4092),
		"Channel Threshold",
		"ChannelThreshold",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409A),
		"Scanner Settings Sequence",
		"ScannerSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409B),
		"Scan Procedure",
		"ScanProcedure",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409C),
		"Translation Rate X",
		"TranslationRateX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409D),
		"Translation Rate Y",
		"TranslationRateY",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409F),
		"Channel Overlap",
		"ChannelOverlap",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x40A0),
		"Image Quality Indicator Type",
		"ImageQualityIndicatorType",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x40A1),
		"Image Quality Indicator Material",
		"ImageQualityIndicatorMaterial",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x40A2),
		"Image Quality Indicator Size",
		"ImageQualityIndicatorSize",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4101),
		"Wave Dimensions Definition Sequence",
		"WaveDimensionsDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4102),
		"Wave Dimension Number",
		"WaveDimensionNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4103),
		"Wave Dimension Description",
		"WaveDimensionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4104),
		"Wave Dimension Unit",
		"WaveDimensionUnit",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4105),
		"Wave Dimension Value Type",
		"WaveDimensionValueType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4106),
		"Wave Dimension Values Sequence",
		"WaveDimensionValuesSequence",
		vm.VM1N,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4107),
		"Referenced Wave Dimension",
		"ReferencedWaveDimension",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4108),
		"Integer Numeric Value",
		"IntegerNumericValue",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4109),
		"Byte Numeric Value",
		"ByteNumericValue",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x410A),
		"Short Numeric Value",
		"ShortNumericValue",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x410B),
		"Single Precision Floating Point Numeric Value",
		"SinglePrecisionFloatingPointNumericValue",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x410C),
		"Double Precision Floating Point Numeric Value",
		"DoublePrecisionFloatingPointNumericValue",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5002),
		"LINAC Energy",
		"LINACEnergy",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5004),
		"LINAC Output",
		"LINACOutput",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5100),
		"Active Aperture",
		"ActiveAperture",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5101),
		"Total Aperture",
		"TotalAperture",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5102),
		"Aperture Elevation",
		"ApertureElevation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5103),
		"Main Lobe Angle",
		"MainLobeAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5104),
		"Main Roof Angle",
		"MainRoofAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5105),
		"Connector Type",
		"ConnectorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5106),
		"Wedge Model Number",
		"WedgeModelNumber",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5107),
		"Wedge Angle Float",
		"WedgeAngleFloat",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5108),
		"Wedge Roof Angle",
		"WedgeRoofAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5109),
		"Wedge Element 1 Position",
		"WedgeElement1Position",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510A),
		"Wedge Material Velocity",
		"WedgeMaterialVelocity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510B),
		"Wedge Material",
		"WedgeMaterial",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510C),
		"Wedge Offset Z",
		"WedgeOffsetZ",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510D),
		"Wedge Origin Offset X",
		"WedgeOriginOffsetX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510E),
		"Wedge Time Delay",
		"WedgeTimeDelay",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510F),
		"Wedge Name",
		"WedgeName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5110),
		"Wedge Manufacturer Name",
		"WedgeManufacturerName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5111),
		"Wedge Description",
		"WedgeDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5112),
		"Nominal Beam Angle",
		"NominalBeamAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5113),
		"Wedge Offset X",
		"WedgeOffsetX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5114),
		"Wedge Offset Y",
		"WedgeOffsetY",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5115),
		"Wedge Total Length",
		"WedgeTotalLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5116),
		"Wedge In Contact Length",
		"WedgeInContactLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5117),
		"Wedge Front Gap",
		"WedgeFrontGap",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5118),
		"Wedge Total Height",
		"WedgeTotalHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5119),
		"Wedge Front Height",
		"WedgeFrontHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511A),
		"Wedge Rear Height",
		"WedgeRearHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511B),
		"Wedge Total Width",
		"WedgeTotalWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511C),
		"Wedge In Contact Width",
		"WedgeInContactWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511D),
		"Wedge Chamfer Height",
		"WedgeChamferHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511E),
		"Wedge Curve",
		"WedgeCurve",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511F),
		"Radius Along the Wedge",
		"RadiusAlongWedge",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6001),
		"Thermal Camera Settings Sequence",
		"ThermalCameraSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6002),
		"Acquisition Frame Rate",
		"AcquisitionFrameRate",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6003),
		"Integration Time",
		"IntegrationTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6004),
		"Number of Calibration Frames",
		"NumberOfCalibrationFrames",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6005),
		"Number of Rows in Full Acquisition Image",
		"NumberOfRowsInFullAcquisitionImage",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6006),
		"Number Of Columns in Full Acquisition Image",
		"NumberOfColumnsInFullAcquisitionImage",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6007),
		"Thermal Source Settings Sequence",
		"ThermalSourceSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6008),
		"Source Horizontal Pitch",
		"SourceHorizontalPitch",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6009),
		"Source Vertical Pitch",
		"SourceVerticalPitch",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600A),
		"Source Horizontal Scan Speed",
		"SourceHorizontalScanSpeed",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600B),
		"Thermal Source Modulation Frequency",
		"ThermalSourceModulationFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600C),
		"Induction Source Setting Sequence",
		"InductionSourceSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600D),
		"Coil Frequency",
		"CoilFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600E),
		"Current Amplitude Across Coil",
		"CurrentAmplitudeAcrossCoil",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600F),
		"Flash Source Setting Sequence",
		"FlashSourceSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6010),
		"Flash Duration",
		"FlashDuration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6011),
		"Flash Frame Number",
		"FlashFrameNumber",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6012),
		"Laser Source Setting Sequence",
		"LaserSourceSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6013),
		"Horizontal Laser Spot Dimension",
		"HorizontalLaserSpotDimension",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6014),
		"Vertical Laser Spot Dimension",
		"VerticalLaserSpotDimension",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6015),
		"Laser Wavelength",
		"LaserWavelength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6016),
		"Laser Power",
		"LaserPower",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6017),
		"Forced Gas Setting Sequence",
		"ForcedGasSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6018),
		"Vibration Source Setting Sequence",
		"VibrationSourceSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6019),
		"Vibration Excitation Frequency",
		"VibrationExcitationFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601A),
		"Vibration Excitation Voltage",
		"VibrationExcitationVoltage",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601B),
		"Thermography Data Capture Method",
		"ThermographyDataCaptureMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601C),
		"Thermal Technique",
		"ThermalTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601D),
		"Thermal Camera Core Sequence",
		"ThermalCameraCoreSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601E),
		"Detector Wavelength Range",
		"DetectorWavelengthRange",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601F),
		"Thermal Camera Calibration Type",
		"ThermalCameraCalibrationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6020),
		"Acquisition Image Counter",
		"AcquisitionImageCounter",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6021),
		"Front Panel Temperature",
		"FrontPanelTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6022),
		"Air Gap Temperature",
		"AirGapTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6023),
		"Vertical Pixel Size",
		"VerticalPixelSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6024),
		"Horizontal Pixel Size",
		"HorizontalPixelSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6025),
		"Data Streaming Protocol",
		"DataStreamingProtocol",
		vm.VM1N,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6026),
		"Lens Sequence",
		"LensSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6027),
		"Field of View",
		"FieldOfView",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6028),
		"Lens Filter Manufacturer",
		"LensFilterManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6029),
		"Cutoff Filter Type",
		"CutoffFilterType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602A),
		"Lens Filter Cut-Off Wavelength",
		"LensFilterCutOffWavelength",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602B),
		"Thermal Source Sequence",
		"ThermalSourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602C),
		"Thermal Source Motion State",
		"ThermalSourceMotionState",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602D),
		"Thermal Source Motion Type",
		"ThermalSourceMotionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602E),
		"Induction Heating Sequence",
		"InductionHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602F),
		"Coil Configuration ID",
		"CoilConfigurationID",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6030),
		"Number of Turns in Coil",
		"NumberOfTurnsInCoil",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6031),
		"Shape of Individual Turn",
		"ShapeOfIndividualTurn",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6032),
		"Size of Individual Turn",
		"SizeOfIndividualTurn",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6033),
		"Distance Between Turns",
		"DistanceBetweenTurns",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6034),
		"Flash Heating Sequence",
		"FlashHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6035),
		"Number of Lamps",
		"NumberOfLamps",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6036),
		"Flash Synchronization Protocol",
		"FlashSynchronizationProtocol",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6037),
		"Flash Modification Status",
		"FlashModificationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6038),
		"Laser Heating Sequence",
		"LaserHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6039),
		"Laser Manufacturer",
		"LaserManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603A),
		"Laser Model Number",
		"LaserModelNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603B),
		"Laser Type Description",
		"LaserTypeDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603C),
		"Forced Gas Heating Sequence",
		"ForcedGasHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603D),
		"Gas Used for Heating/Cooling Part",
		"GasUsedForHeatingCoolingPart",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603E),
		"Vibration/Sonic Heating Sequence",
		"VibrationSonicHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603F),
		"Probe Manufacturer",
		"ProbeManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6040),
		"Probe Model Number",
		"ProbeModelNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6041),
		"Aperture Size",
		"ApertureSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6042),
		"Probe Resonant Frequency",
		"ProbeResonantFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6043),
		"Heat Source Description",
		"HeatSourceDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6044),
		"Surface Preparation with Optical Coating",
		"SurfacePreparationWithOpticalCoating",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6045),
		"Optical Coating Type",
		"OpticalCoatingType",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6046),
		"Thermal Conductivity of Exposed Surface",
		"ThermalConductivityOfExposedSurface",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6047),
		"Material Density",
		"MaterialDensity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6048),
		"Specific Heat of Inspection Surface",
		"SpecificHeatOfInspectionSurface",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6049),
		"Emissivity of Inspection Surface",
		"EmissivityOfInspectionSurface",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604A),
		"Electromagnetic Classification of Inspection Surface",
		"ElectromagneticClassificationOfInspectionSurface",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604C),
		"Moving Window Size",
		"MovingWindowSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604D),
		"Moving Window Type",
		"MovingWindowType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604E),
		"Moving Window Weights",
		"MovingWindowWeights",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604F),
		"Moving Window Pitch",
		"MovingWindowPitch",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6050),
		"Moving Window Padding Scheme",
		"MovingWindowPaddingScheme",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6051),
		"Moving Window Padding Length",
		"MovingWindowPaddingLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6052),
		"Spatial Filtering Parameters Sequence",
		"SpatialFilteringParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6053),
		"Spatial Filtering Scheme",
		"SpatialFilteringScheme",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6056),
		"Horizontal Moving Window Size",
		"HorizontalMovingWindowSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6057),
		"Vertical Moving Window Size",
		"VerticalMovingWindowSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6059),
		"Polynomial Fitting Sequence",
		"PolynomialFittingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605A),
		"Fitting Data Type",
		"FittingDataType",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605B),
		"Operation on Time Axis Before Fitting",
		"OperationOnTimeAxisBeforeFitting",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605C),
		"Operation on Pixel Intensity Before Fitting",
		"OperationOnPixelIntensityBeforeFitting",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605D),
		"Order of Polynomial",
		"OrderOfPolynomial",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605E),
		"Independent Variable for Polynomial Fit",
		"IndependentVariableForPolynomialFit",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605F),
		"PolynomialCoefficients",
		"PolynomialCoefficients",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6060),
		"Thermography Pixel Data Unit",
		"ThermographyPixelDataUnit",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0001),
		"White Point",
		"WhitePoint",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0002),
		"Primary Chromaticities",
		"PrimaryChromaticities",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0003),
		"Battery Level",
		"BatteryLevel",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0004),
		"Exposure Time in Seconds",
		"ExposureTimeInSeconds",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0005),
		"F-Number",
		"FNumber",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0006),
		"OECF Rows",
		"OECFRows",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0007),
		"OECF Columns",
		"OECFColumns",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0008),
		"OECF Column Names",
		"OECFColumnNames",
		vm.VM1N,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0009),
		"OECF Values",
		"OECFValues",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000A),
		"Spatial Frequency Response Rows",
		"SpatialFrequencyResponseRows",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000B),
		"Spatial Frequency Response Columns",
		"SpatialFrequencyResponseColumns",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000C),
		"Spatial Frequency Response Column Names",
		"SpatialFrequencyResponseColumnNames",
		vm.VM1N,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000D),
		"Spatial Frequency Response Values",
		"SpatialFrequencyResponseValues",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000E),
		"Color Filter Array Pattern Rows",
		"ColorFilterArrayPatternRows",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000F),
		"Color Filter Array Pattern Columns",
		"ColorFilterArrayPatternColumns",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0010),
		"Color Filter Array Pattern Values",
		"ColorFilterArrayPatternValues",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0011),
		"Flash Firing Status",
		"FlashFiringStatus",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0012),
		"Flash Return Status",
		"FlashReturnStatus",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0013),
		"Flash Mode",
		"FlashMode",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0014),
		"Flash Function Present",
		"FlashFunctionPresent",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0015),
		"Flash Red Eye Mode",
		"FlashRedEyeMode",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0016),
		"Exposure Program",
		"ExposureProgram",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0017),
		"Spectral Sensitivity",
		"SpectralSensitivity",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0018),
		"Photographic Sensitivity",
		"PhotographicSensitivity",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0019),
		"Self Timer Mode",
		"SelfTimerMode",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001A),
		"Sensitivity Type",
		"SensitivityType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001B),
		"Standard Output Sensitivity",
		"StandardOutputSensitivity",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001C),
		"Recommended Exposure Index",
		"RecommendedExposureIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001D),
		"ISO Speed",
		"ISOSpeed",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001E),
		"ISO Speed Latitude yyy",
		"ISOSpeedLatitudeyyy",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001F),
		"ISO Speed Latitude zzz",
		"ISOSpeedLatitudezzz",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0020),
		"EXIF Version",
		"EXIFVersion",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0021),
		"Shutter Speed Value",
		"ShutterSpeedValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0022),
		"Aperture Value",
		"ApertureValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0023),
		"Brightness Value",
		"BrightnessValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0024),
		"Exposure Bias Value",
		"ExposureBiasValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0025),
		"Max Aperture Value",
		"MaxApertureValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0026),
		"Subject Distance",
		"SubjectDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0027),
		"Metering Mode",
		"MeteringMode",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0028),
		"Light Source",
		"LightSource",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0029),
		"Focal Length",
		"FocalLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x002A),
		"Subject Area",
		"SubjectArea",
		vm.MustParse("2-4"),
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x002B),
		"Maker Note",
		"MakerNote",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0030),
		"Temperature",
		"Temperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0031),
		"Humidity",
		"Humidity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0032),
		"Pressure",
		"Pressure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0033),
		"Water Depth",
		"WaterDepth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0034),
		"Acceleration",
		"Acceleration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0035),
		"Camera Elevation Angle",
		"CameraElevationAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0036),
		"Flash Energy",
		"FlashEnergy",
		vm.VM12,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0037),
		"Subject Location",
		"SubjectLocation",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0038),
		"Photographic Exposure Index",
		"PhotographicExposureIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0039),
		"Sensing Method",
		"SensingMethod",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x003A),
		"File Source",
		"FileSource",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x003B),
		"Scene Type",
		"SceneType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0041),
		"Custom Rendered",
		"CustomRendered",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0042),
		"Exposure Mode",
		"ExposureMode",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0043),
		"White Balance",
		"WhiteBalance",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0044),
		"Digital Zoom Ratio",
		"DigitalZoomRatio",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0045),
		"Focal Length In 35mm Film",
		"FocalLengthIn35mmFilm",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0046),
		"Scene Capture Type",
		"SceneCaptureType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0047),
		"Gain Control",
		"GainControl",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0048),
		"Contrast",
		"Contrast",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0049),
		"Saturation",
		"Saturation",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004A),
		"Sharpness",
		"Sharpness",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004B),
		"Device Setting Description",
		"DeviceSettingDescription",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004C),
		"Subject Distance Range",
		"SubjectDistanceRange",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004D),
		"Camera Owner Name",
		"CameraOwnerName",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004E),
		"Lens Specification",
		"LensSpecification",
		vm.VM4,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004F),
		"Lens Make",
		"LensMake",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0050),
		"Lens Model",
		"LensModel",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0051),
		"Lens Serial Number",
		"LensSerialNumber",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0061),
		"Interoperability Index",
		"InteroperabilityIndex",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0062),
		"Interoperability Version",
		"InteroperabilityVersion",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0070),
		"GPS Version ID",
		"GPSVersionID",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0071),
		"GPS Latitude Ref",
		"GPSLatitudeRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0072),
		"GPS Latitude",
		"GPSLatitude",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0073),
		"GPS Longitude Ref",
		"GPSLongitudeRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0074),
		"GPS Longitude",
		"GPSLongitude",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0075),
		"GPS Altitude Ref",
		"GPSAltitudeRef",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0076),
		"GPS Altitude",
		"GPSAltitude",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0077),
		"GPS Time Stamp",
		"GPSTimeStamp",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0078),
		"GPS Satellites",
		"GPSSatellites",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0079),
		"GPS Status",
		"GPSStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007A),
		"GPS Measure Mode",
		"GPSMeasureMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007B),
		"GPS DOP",
		"GPSDOP",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007C),
		"GPS Speed Ref",
		"GPSSpeedRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007D),
		"GPS Speed",
		"GPSSpeed",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007E),
		"GPS Track Ref",
		"GPSTrackRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007F),
		"GPS Track",
		"GPSTrack",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0080),
		"GPS Img Direction Ref",
		"GPSImgDirectionRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0081),
		"GPS Img Direction",
		"GPSImgDirection",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0082),
		"GPS Map Datum",
		"GPSMapDatum",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0083),
		"GPS Dest Latitude Ref",
		"GPSDestLatitudeRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0084),
		"GPS Dest Latitude",
		"GPSDestLatitude",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0085),
		"GPS Dest Longitude Ref",
		"GPSDestLongitudeRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0086),
		"GPS Dest Longitude",
		"GPSDestLongitude",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0087),
		"GPS Dest Bearing Ref",
		"GPSDestBearingRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0088),
		"GPS Dest Bearing",
		"GPSDestBearing",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0089),
		"GPS Dest Distance Ref",
		"GPSDestDistanceRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008A),
		"GPS Dest Distance",
		"GPSDestDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008B),
		"GPS Processing Method",
		"GPSProcessingMethod",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008C),
		"GPS Area Information",
		"GPSAreaInformation",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008D),
		"GPS Date Stamp",
		"GPSDateStamp",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008E),
		"GPS Differential",
		"GPSDifferential",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1001),
		"Light Source Polarization",
		"LightSourcePolarization",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1002),
		"Emitter Color Temperature",
		"EmitterColorTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1003),
		"Contact Method",
		"ContactMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1004),
		"Immersion Media",
		"ImmersionMedia",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1005),
		"Optical Magnification Factor",
		"OpticalMagnificationFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0010),
		"Contrast/Bolus Agent",
		"ContrastBolusAgent",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0012),
		"Contrast/Bolus Agent Sequence",
		"ContrastBolusAgentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0013),
		"Contrast/Bolus T1 Relaxivity",
		"ContrastBolusT1Relaxivity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0014),
		"Contrast/Bolus Administration Route Sequence",
		"ContrastBolusAdministrationRouteSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0015),
		"Body Part Examined",
		"BodyPartExamined",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0020),
		"Scanning Sequence",
		"ScanningSequence",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0021),
		"Sequence Variant",
		"SequenceVariant",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0022),
		"Scan Options",
		"ScanOptions",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0023),
		"MR Acquisition Type",
		"MRAcquisitionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0024),
		"Sequence Name",
		"SequenceName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0025),
		"Angio Flag",
		"AngioFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0026),
		"Intervention Drug Information Sequence",
		"InterventionDrugInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0027),
		"Intervention Drug Stop Time",
		"InterventionDrugStopTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0028),
		"Intervention Drug Dose",
		"InterventionDrugDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0029),
		"Intervention Drug Code Sequence",
		"InterventionDrugCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x002A),
		"Additional Drug Sequence",
		"AdditionalDrugSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0030),
		"Radionuclide",
		"Radionuclide",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0031),
		"Radiopharmaceutical",
		"Radiopharmaceutical",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0032),
		"Energy Window Centerline",
		"EnergyWindowCenterline",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0033),
		"Energy Window Total Width",
		"EnergyWindowTotalWidth",
		vm.VM1N,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0034),
		"Intervention Drug Name",
		"InterventionDrugName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0035),
		"Intervention Drug Start Time",
		"InterventionDrugStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0036),
		"Intervention Sequence",
		"InterventionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0037),
		"Therapy Type",
		"TherapyType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0038),
		"Intervention Status",
		"InterventionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0039),
		"Therapy Description",
		"TherapyDescription",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x003A),
		"Intervention Description",
		"InterventionDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0040),
		"Cine Rate",
		"CineRate",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0042),
		"Initial Cine Run State",
		"InitialCineRunState",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0050),
		"Slice Thickness",
		"SliceThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0060),
		"KVP",
		"KVP",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0070),
		"Counts Accumulated",
		"CountsAccumulated",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0071),
		"Acquisition Termination Condition",
		"AcquisitionTerminationCondition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0072),
		"Effective Duration",
		"EffectiveDuration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0073),
		"Acquisition Start Condition",
		"AcquisitionStartCondition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0074),
		"Acquisition Start Condition Data",
		"AcquisitionStartConditionData",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0075),
		"Acquisition Termination Condition Data",
		"AcquisitionTerminationConditionData",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0080),
		"Repetition Time",
		"RepetitionTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0081),
		"Echo Time",
		"EchoTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0082),
		"Inversion Time",
		"InversionTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0083),
		"Number of Averages",
		"NumberOfAverages",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0084),
		"Imaging Frequency",
		"ImagingFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0085),
		"Imaged Nucleus",
		"ImagedNucleus",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0086),
		"Echo Number(s)",
		"EchoNumbers",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0087),
		"Magnetic Field Strength",
		"MagneticFieldStrength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0088),
		"Spacing Between Slices",
		"SpacingBetweenSlices",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0089),
		"Number of Phase Encoding Steps",
		"NumberOfPhaseEncodingSteps",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0090),
		"Data Collection Diameter",
		"DataCollectionDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0091),
		"Echo Train Length",
		"EchoTrainLength",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0093),
		"Percent Sampling",
		"PercentSampling",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0094),
		"Percent Phase Field of View",
		"PercentPhaseFieldOfView",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0095),
		"Pixel Bandwidth",
		"PixelBandwidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1000),
		"Device Serial Number",
		"DeviceSerialNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1002),
		"Device UID",
		"DeviceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1003),
		"Device ID",
		"DeviceID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1004),
		"Plate ID",
		"PlateID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1005),
		"Generator ID",
		"GeneratorID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1006),
		"Grid ID",
		"GridID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1007),
		"Cassette ID",
		"CassetteID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1008),
		"Gantry ID",
		"GantryID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1009),
		"Unique Device Identifier",
		"UniqueDeviceIdentifier",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x100A),
		"UDI Sequence",
		"UDISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x100B),
		"Manufacturer's Device Class UID",
		"ManufacturerDeviceClassUID",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1010),
		"Secondary Capture Device ID",
		"SecondaryCaptureDeviceID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1011),
		"Hardcopy Creation Device ID",
		"HardcopyCreationDeviceID",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1012),
		"Date of Secondary Capture",
		"DateOfSecondaryCapture",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1014),
		"Time of Secondary Capture",
		"TimeOfSecondaryCapture",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1016),
		"Secondary Capture Device Manufacturer",
		"SecondaryCaptureDeviceManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1017),
		"Hardcopy Device Manufacturer",
		"HardcopyDeviceManufacturer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1018),
		"Secondary Capture Device Manufacturer's Model Name",
		"SecondaryCaptureDeviceManufacturerModelName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1019),
		"Secondary Capture Device Software Versions",
		"SecondaryCaptureDeviceSoftwareVersions",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x101A),
		"Hardcopy Device Software Version",
		"HardcopyDeviceSoftwareVersion",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x101B),
		"Hardcopy Device Manufacturer's Model Name",
		"HardcopyDeviceManufacturerModelName",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1020),
		"Software Versions",
		"SoftwareVersions",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1022),
		"Video Image Format Acquired",
		"VideoImageFormatAcquired",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1023),
		"Digital Image Format Acquired",
		"DigitalImageFormatAcquired",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1030),
		"Protocol Name",
		"ProtocolName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1040),
		"Contrast/Bolus Route",
		"ContrastBolusRoute",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1041),
		"Contrast/Bolus Volume",
		"ContrastBolusVolume",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1042),
		"Contrast/Bolus Start Time",
		"ContrastBolusStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1043),
		"Contrast/Bolus Stop Time",
		"ContrastBolusStopTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1044),
		"Contrast/Bolus Total Dose",
		"ContrastBolusTotalDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1045),
		"Syringe Counts",
		"SyringeCounts",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1046),
		"Contrast Flow Rate",
		"ContrastFlowRate",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1047),
		"Contrast Flow Duration",
		"ContrastFlowDuration",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1048),
		"Contrast/Bolus Ingredient",
		"ContrastBolusIngredient",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1049),
		"Contrast/Bolus Ingredient Concentration",
		"ContrastBolusIngredientConcentration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1050),
		"Spatial Resolution",
		"SpatialResolution",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1060),
		"Trigger Time",
		"TriggerTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1061),
		"Trigger Source or Type",
		"TriggerSourceOrType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1062),
		"Nominal Interval",
		"NominalInterval",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1063),
		"Frame Time",
		"FrameTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1064),
		"Cardiac Framing Type",
		"CardiacFramingType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1065),
		"Frame Time Vector",
		"FrameTimeVector",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1066),
		"Frame Delay",
		"FrameDelay",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1067),
		"Image Trigger Delay",
		"ImageTriggerDelay",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1068),
		"Multiplex Group Time Offset",
		"MultiplexGroupTimeOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1069),
		"Trigger Time Offset",
		"TriggerTimeOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x106A),
		"Synchronization Trigger",
		"SynchronizationTrigger",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x106C),
		"Synchronization Channel",
		"SynchronizationChannel",
		vm.VM2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x106E),
		"Trigger Sample Position",
		"TriggerSamplePosition",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1070),
		"Radiopharmaceutical Route",
		"RadiopharmaceuticalRoute",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1071),
		"Radiopharmaceutical Volume",
		"RadiopharmaceuticalVolume",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1072),
		"Radiopharmaceutical Start Time",
		"RadiopharmaceuticalStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1073),
		"Radiopharmaceutical Stop Time",
		"RadiopharmaceuticalStopTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1074),
		"Radionuclide Total Dose",
		"RadionuclideTotalDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1075),
		"Radionuclide Half Life",
		"RadionuclideHalfLife",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1076),
		"Radionuclide Positron Fraction",
		"RadionuclidePositronFraction",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1077),
		"Radiopharmaceutical Specific Activity",
		"RadiopharmaceuticalSpecificActivity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1078),
		"Radiopharmaceutical Start DateTime",
		"RadiopharmaceuticalStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1079),
		"Radiopharmaceutical Stop DateTime",
		"RadiopharmaceuticalStopDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1080),
		"Beat Rejection Flag",
		"BeatRejectionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1081),
		"Low R-R Value",
		"LowRRValue",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1082),
		"High R-R Value",
		"HighRRValue",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1083),
		"Intervals Acquired",
		"IntervalsAcquired",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1084),
		"Intervals Rejected",
		"IntervalsRejected",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1085),
		"PVC Rejection",
		"PVCRejection",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1086),
		"Skip Beats",
		"SkipBeats",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1088),
		"Heart Rate",
		"HeartRate",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1090),
		"Cardiac Number of Images",
		"CardiacNumberOfImages",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1094),
		"Trigger Window",
		"TriggerWindow",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1100),
		"Reconstruction Diameter",
		"ReconstructionDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1110),
		"Distance Source to Detector",
		"DistanceSourceToDetector",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1111),
		"Distance Source to Patient",
		"DistanceSourceToPatient",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1114),
		"Estimated Radiographic Magnification Factor",
		"EstimatedRadiographicMagnificationFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1120),
		"Gantry/Detector Tilt",
		"GantryDetectorTilt",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1121),
		"Gantry/Detector Slew",
		"GantryDetectorSlew",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1130),
		"Table Height",
		"TableHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1131),
		"Table Traverse",
		"TableTraverse",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1134),
		"Table Motion",
		"TableMotion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1135),
		"Table Vertical Increment",
		"TableVerticalIncrement",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1136),
		"Table Lateral Increment",
		"TableLateralIncrement",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1137),
		"Table Longitudinal Increment",
		"TableLongitudinalIncrement",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1138),
		"Table Angle",
		"TableAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x113A),
		"Table Type",
		"TableType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1140),
		"Rotation Direction",
		"RotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1141),
		"Angular Position",
		"AngularPosition",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1142),
		"Radial Position",
		"RadialPosition",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1143),
		"Scan Arc",
		"ScanArc",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1144),
		"Angular Step",
		"AngularStep",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1145),
		"Center of Rotation Offset",
		"CenterOfRotationOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1146),
		"Rotation Offset",
		"RotationOffset",
		vm.VM1N,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1147),
		"Field of View Shape",
		"FieldOfViewShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1149),
		"Field of View Dimension(s)",
		"FieldOfViewDimensions",
		vm.VM12,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1150),
		"Exposure Time",
		"ExposureTime",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1151),
		"X-Ray Tube Current",
		"XRayTubeCurrent",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1152),
		"Exposure",
		"Exposure",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1153),
		"Exposure in µAs",
		"ExposureInuAs",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1154),
		"Average Pulse Width",
		"AveragePulseWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1155),
		"Radiation Setting",
		"RadiationSetting",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1156),
		"Rectification Type",
		"RectificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x115A),
		"Radiation Mode",
		"RadiationMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x115E),
		"Image and Fluoroscopy Area Dose Product",
		"ImageAndFluoroscopyAreaDoseProduct",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1160),
		"Filter Type",
		"FilterType",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1161),
		"Type of Filters",
		"TypeOfFilters",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1162),
		"Intensifier Size",
		"IntensifierSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1164),
		"Imager Pixel Spacing",
		"ImagerPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1166),
		"Grid",
		"Grid",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1170),
		"Generator Power",
		"GeneratorPower",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1180),
		"Collimator/grid Name",
		"CollimatorGridName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1181),
		"Collimator Type",
		"CollimatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1182),
		"Focal Distance",
		"FocalDistance",
		vm.VM12,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1183),
		"X Focus Center",
		"XFocusCenter",
		vm.VM12,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1184),
		"Y Focus Center",
		"YFocusCenter",
		vm.VM12,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1190),
		"Focal Spot(s)",
		"FocalSpots",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1191),
		"Anode Target Material",
		"AnodeTargetMaterial",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A0),
		"Body Part Thickness",
		"BodyPartThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A2),
		"Compression Force",
		"CompressionForce",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A3),
		"Compression Pressure",
		"CompressionPressure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A4),
		"Paddle Description",
		"PaddleDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A5),
		"Compression Contact Area",
		"CompressionContactArea",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B0),
		"Acquisition Mode",
		"AcquisitionMode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B1),
		"Dose Mode Name",
		"DoseModeName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B2),
		"Acquired Subtraction Mask Flag",
		"AcquiredSubtractionMaskFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B3),
		"Fluoroscopy Persistence Flag",
		"FluoroscopyPersistenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B4),
		"Fluoroscopy Last Image Hold Persistence Flag",
		"FluoroscopyLastImageHoldPersistenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B5),
		"Upper Limit Number Of Persistent Fluoroscopy Frames",
		"UpperLimitNumberOfPersistentFluoroscopyFrames",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B6),
		"Contrast/Bolus Auto Injection Trigger Flag",
		"ContrastBolusAutoInjectionTriggerFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B7),
		"Contrast/Bolus Injection Delay",
		"ContrastBolusInjectionDelay",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B8),
		"XA Acquisition Phase Details Sequence",
		"XAAcquisitionPhaseDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B9),
		"XA Acquisition Frame Rate",
		"XAAcquisitionFrameRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BA),
		"XA Plane Details Sequence",
		"XAPlaneDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BB),
		"Acquisition Field of View Label",
		"AcquisitionFieldOfViewLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BC),
		"X-Ray Filter Details Sequence",
		"XRayFilterDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BD),
		"XA Acquisition Duration",
		"XAAcquisitionDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BE),
		"Reconstruction Pipeline Type",
		"ReconstructionPipelineType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BF),
		"Image Filter Details Sequence",
		"ImageFilterDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11C0),
		"Applied Mask Subtraction Flag",
		"AppliedMaskSubtractionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11C1),
		"Requested Series Description Code Sequence",
		"RequestedSeriesDescriptionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1200),
		"Date of Last Calibration",
		"DateOfLastCalibration",
		vm.VM1N,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1201),
		"Time of Last Calibration",
		"TimeOfLastCalibration",
		vm.VM1N,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1202),
		"DateTime of Last Calibration",
		"DateTimeOfLastCalibration",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1203),
		"Calibration DateTime",
		"CalibrationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1204),
		"Date of Manufacture",
		"DateOfManufacture",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1205),
		"Date of Installation",
		"DateOfInstallation",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1210),
		"Convolution Kernel",
		"ConvolutionKernel",
		vm.VM1N,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1240),
		"Upper/Lower Pixel Values",
		"UpperLowerPixelValues",
		vm.VM1N,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1242),
		"Actual Frame Duration",
		"ActualFrameDuration",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1243),
		"Count Rate",
		"CountRate",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1244),
		"Preferred Playback Sequencing",
		"PreferredPlaybackSequencing",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1250),
		"Receive Coil Name",
		"ReceiveCoilName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1251),
		"Transmit Coil Name",
		"TransmitCoilName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1260),
		"Plate Type",
		"PlateType",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1261),
		"Phosphor Type",
		"PhosphorType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1271),
		"Water Equivalent Diameter",
		"WaterEquivalentDiameter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1272),
		"Water Equivalent Diameter Calculation Method Code Sequence",
		"WaterEquivalentDiameterCalculationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1300),
		"Scan Velocity",
		"ScanVelocity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1301),
		"Whole Body Technique",
		"WholeBodyTechnique",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1302),
		"Scan Length",
		"ScanLength",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1310),
		"Acquisition Matrix",
		"AcquisitionMatrix",
		vm.VM4,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1312),
		"In-plane Phase Encoding Direction",
		"InPlanePhaseEncodingDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1314),
		"Flip Angle",
		"FlipAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1315),
		"Variable Flip Angle Flag",
		"VariableFlipAngleFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1316),
		"SAR",
		"SAR",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1318),
		"dB/dt",
		"dBdt",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1320),
		"B1rms",
		"B1rms",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1400),
		"Acquisition Device Processing Description",
		"AcquisitionDeviceProcessingDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1401),
		"Acquisition Device Processing Code",
		"AcquisitionDeviceProcessingCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1402),
		"Cassette Orientation",
		"CassetteOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1403),
		"Cassette Size",
		"CassetteSize",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1404),
		"Exposures on Plate",
		"ExposuresOnPlate",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1405),
		"Relative X-Ray Exposure",
		"RelativeXRayExposure",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1411),
		"Exposure Index",
		"ExposureIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1412),
		"Target Exposure Index",
		"TargetExposureIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1413),
		"Deviation Index",
		"DeviationIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1450),
		"Column Angulation",
		"ColumnAngulation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1460),
		"Tomo Layer Height",
		"TomoLayerHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1470),
		"Tomo Angle",
		"TomoAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1480),
		"Tomo Time",
		"TomoTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1490),
		"Tomo Type",
		"TomoType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1491),
		"Tomo Class",
		"TomoClass",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1495),
		"Number of Tomosynthesis Source Images",
		"NumberOfTomosynthesisSourceImages",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1500),
		"Positioner Motion",
		"PositionerMotion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1508),
		"Positioner Type",
		"PositionerType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1510),
		"Positioner Primary Angle",
		"PositionerPrimaryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1511),
		"Positioner Secondary Angle",
		"PositionerSecondaryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1520),
		"Positioner Primary Angle Increment",
		"PositionerPrimaryAngleIncrement",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1521),
		"Positioner Secondary Angle Increment",
		"PositionerSecondaryAngleIncrement",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1530),
		"Detector Primary Angle",
		"DetectorPrimaryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1531),
		"Detector Secondary Angle",
		"DetectorSecondaryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1600),
		"Shutter Shape",
		"ShutterShape",
		vm.VM13,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1602),
		"Shutter Left Vertical Edge",
		"ShutterLeftVerticalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1604),
		"Shutter Right Vertical Edge",
		"ShutterRightVerticalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1606),
		"Shutter Upper Horizontal Edge",
		"ShutterUpperHorizontalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1608),
		"Shutter Lower Horizontal Edge",
		"ShutterLowerHorizontalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1610),
		"Center of Circular Shutter",
		"CenterOfCircularShutter",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1612),
		"Radius of Circular Shutter",
		"RadiusOfCircularShutter",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1620),
		"Vertices of the Polygonal Shutter",
		"VerticesOfThePolygonalShutter",
		vm.VM22N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1622),
		"Shutter Presentation Value",
		"ShutterPresentationValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1623),
		"Shutter Overlay Group",
		"ShutterOverlayGroup",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1624),
		"Shutter Presentation Color CIELab Value",
		"ShutterPresentationColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1630),
		"Outline Shape Type",
		"OutlineShapeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1631),
		"Outline Left Vertical Edge",
		"OutlineLeftVerticalEdge",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1632),
		"Outline Right Vertical Edge",
		"OutlineRightVerticalEdge",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1633),
		"Outline Upper Horizontal Edge",
		"OutlineUpperHorizontalEdge",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1634),
		"Outline Lower Horizontal Edge",
		"OutlineLowerHorizontalEdge",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1635),
		"Center of Circular Outline",
		"CenterOfCircularOutline",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1636),
		"Diameter of Circular Outline",
		"DiameterOfCircularOutline",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1637),
		"Number of Polygonal Vertices",
		"NumberOfPolygonalVertices",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1638),
		"Vertices of the Polygonal Outline",
		"VerticesOfThePolygonalOutline",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1700),
		"Collimator Shape",
		"CollimatorShape",
		vm.VM13,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1702),
		"Collimator Left Vertical Edge",
		"CollimatorLeftVerticalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1704),
		"Collimator Right Vertical Edge",
		"CollimatorRightVerticalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1706),
		"Collimator Upper Horizontal Edge",
		"CollimatorUpperHorizontalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1708),
		"Collimator Lower Horizontal Edge",
		"CollimatorLowerHorizontalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1710),
		"Center of Circular Collimator",
		"CenterOfCircularCollimator",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1712),
		"Radius of Circular Collimator",
		"RadiusOfCircularCollimator",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1720),
		"Vertices of the Polygonal Collimator",
		"VerticesOfThePolygonalCollimator",
		vm.VM22N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1800),
		"Acquisition Time Synchronized",
		"AcquisitionTimeSynchronized",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1801),
		"Time Source",
		"TimeSource",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1802),
		"Time Distribution Protocol",
		"TimeDistributionProtocol",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1803),
		"NTP Source Address",
		"NTPSourceAddress",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2001),
		"Page Number Vector",
		"PageNumberVector",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2002),
		"Frame Label Vector",
		"FrameLabelVector",
		vm.VM1N,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2003),
		"Frame Primary Angle Vector",
		"FramePrimaryAngleVector",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2004),
		"Frame Secondary Angle Vector",
		"FrameSecondaryAngleVector",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2005),
		"Slice Location Vector",
		"SliceLocationVector",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2006),
		"Display Window Label Vector",
		"DisplayWindowLabelVector",
		vm.VM1N,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2010),
		"Nominal Scanned Pixel Spacing",
		"NominalScannedPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2020),
		"Digitizing Device Transport Direction",
		"DigitizingDeviceTransportDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2030),
		"Rotation of Scanned Film",
		"RotationOfScannedFilm",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2041),
		"Biopsy Target Sequence",
		"BiopsyTargetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2042),
		"Target UID",
		"TargetUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2043),
		"Localizing Cursor Position",
		"LocalizingCursorPosition",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2044),
		"Calculated Target Position",
		"CalculatedTargetPosition",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2045),
		"Target Label",
		"TargetLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2046),
		"Displayed Z Value",
		"DisplayedZValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3100),
		"IVUS Acquisition",
		"IVUSAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3101),
		"IVUS Pullback Rate",
		"IVUSPullbackRate",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3102),
		"IVUS Gated Rate",
		"IVUSGatedRate",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3103),
		"IVUS Pullback Start Frame Number",
		"IVUSPullbackStartFrameNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3104),
		"IVUS Pullback Stop Frame Number",
		"IVUSPullbackStopFrameNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3105),
		"Lesion Number",
		"LesionNumber",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x4000),
		"Acquisition Comments",
		"AcquisitionComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5000),
		"Output Power",
		"OutputPower",
		vm.VM1N,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5010),
		"Transducer Data",
		"TransducerData",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5011),
		"Transducer Identification Sequence",
		"TransducerIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5012),
		"Focus Depth",
		"FocusDepth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5020),
		"Processing Function",
		"ProcessingFunction",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5021),
		"Postprocessing Function",
		"PostprocessingFunction",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5022),
		"Mechanical Index",
		"MechanicalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5024),
		"Bone Thermal Index",
		"BoneThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5026),
		"Cranial Thermal Index",
		"CranialThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5027),
		"Soft Tissue Thermal Index",
		"SoftTissueThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5028),
		"Soft Tissue-focus Thermal Index",
		"SoftTissueFocusThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5029),
		"Soft Tissue-surface Thermal Index",
		"SoftTissueSurfaceThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5030),
		"Dynamic Range",
		"DynamicRange",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5040),
		"Total Gain",
		"TotalGain",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5050),
		"Depth of Scan Field",
		"DepthOfScanField",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5100),
		"Patient Position",
		"PatientPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5101),
		"View Position",
		"ViewPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5104),
		"Projection Eponymous Name Code Sequence",
		"ProjectionEponymousNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5210),
		"Image Transformation Matrix",
		"ImageTransformationMatrix",
		vm.VM6,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5212),
		"Image Translation Vector",
		"ImageTranslationVector",
		vm.VM3,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6000),
		"Sensitivity",
		"Sensitivity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6011),
		"Sequence of Ultrasound Regions",
		"SequenceOfUltrasoundRegions",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6012),
		"Region Spatial Format",
		"RegionSpatialFormat",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6014),
		"Region Data Type",
		"RegionDataType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6016),
		"Region Flags",
		"RegionFlags",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6018),
		"Region Location Min X0",
		"RegionLocationMinX0",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x601A),
		"Region Location Min Y0",
		"RegionLocationMinY0",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x601C),
		"Region Location Max X1",
		"RegionLocationMaxX1",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x601E),
		"Region Location Max Y1",
		"RegionLocationMaxY1",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6020),
		"Reference Pixel X0",
		"ReferencePixelX0",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6022),
		"Reference Pixel Y0",
		"ReferencePixelY0",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6024),
		"Physical Units X Direction",
		"PhysicalUnitsXDirection",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6026),
		"Physical Units Y Direction",
		"PhysicalUnitsYDirection",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6028),
		"Reference Pixel Physical Value X",
		"ReferencePixelPhysicalValueX",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x602A),
		"Reference Pixel Physical Value Y",
		"ReferencePixelPhysicalValueY",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x602C),
		"Physical Delta X",
		"PhysicalDeltaX",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x602E),
		"Physical Delta Y",
		"PhysicalDeltaY",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6030),
		"Transducer Frequency",
		"TransducerFrequency",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6031),
		"Transducer Type",
		"TransducerType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6032),
		"Pulse Repetition Frequency",
		"PulseRepetitionFrequency",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6034),
		"Doppler Correction Angle",
		"DopplerCorrectionAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6036),
		"Steering Angle",
		"SteeringAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6038),
		"Doppler Sample Volume X Position (Retired)",
		"DopplerSampleVolumeXPositionRetired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6039),
		"Doppler Sample Volume X Position",
		"DopplerSampleVolumeXPosition",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603A),
		"Doppler Sample Volume Y Position (Retired)",
		"DopplerSampleVolumeYPositionRetired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603B),
		"Doppler Sample Volume Y Position",
		"DopplerSampleVolumeYPosition",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603C),
		"TM-Line Position X0 (Retired)",
		"TMLinePositionX0Retired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603D),
		"TM-Line Position X0",
		"TMLinePositionX0",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603E),
		"TM-Line Position Y0 (Retired)",
		"TMLinePositionY0Retired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603F),
		"TM-Line Position Y0",
		"TMLinePositionY0",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6040),
		"TM-Line Position X1 (Retired)",
		"TMLinePositionX1Retired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6041),
		"TM-Line Position X1",
		"TMLinePositionX1",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6042),
		"TM-Line Position Y1 (Retired)",
		"TMLinePositionY1Retired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6043),
		"TM-Line Position Y1",
		"TMLinePositionY1",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6044),
		"Pixel Component Organization",
		"PixelComponentOrganization",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6046),
		"Pixel Component Mask",
		"PixelComponentMask",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6048),
		"Pixel Component Range Start",
		"PixelComponentRangeStart",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x604A),
		"Pixel Component Range Stop",
		"PixelComponentRangeStop",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x604C),
		"Pixel Component Physical Units",
		"PixelComponentPhysicalUnits",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x604E),
		"Pixel Component Data Type",
		"PixelComponentDataType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6050),
		"Number of Table Break Points",
		"NumberOfTableBreakPoints",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6052),
		"Table of X Break Points",
		"TableOfXBreakPoints",
		vm.VM1N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6054),
		"Table of Y Break Points",
		"TableOfYBreakPoints",
		vm.VM1N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6056),
		"Number of Table Entries",
		"NumberOfTableEntries",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6058),
		"Table of Pixel Values",
		"TableOfPixelValues",
		vm.VM1N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x605A),
		"Table of Parameter Values",
		"TableOfParameterValues",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6060),
		"R Wave Time Vector",
		"RWaveTimeVector",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6070),
		"Active Image Area Overlay Group",
		"ActiveImageAreaOverlayGroup",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7000),
		"Detector Conditions Nominal Flag",
		"DetectorConditionsNominalFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7001),
		"Detector Temperature",
		"DetectorTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7004),
		"Detector Type",
		"DetectorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7005),
		"Detector Configuration",
		"DetectorConfiguration",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7006),
		"Detector Description",
		"DetectorDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7008),
		"Detector Mode",
		"DetectorMode",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x700A),
		"Detector ID",
		"DetectorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x700C),
		"Date of Last Detector Calibration",
		"DateOfLastDetectorCalibration",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x700E),
		"Time of Last Detector Calibration",
		"TimeOfLastDetectorCalibration",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7010),
		"Exposures on Detector Since Last Calibration",
		"ExposuresOnDetectorSinceLastCalibration",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7011),
		"Exposures on Detector Since Manufactured",
		"ExposuresOnDetectorSinceManufactured",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7012),
		"Detector Time Since Last Exposure",
		"DetectorTimeSinceLastExposure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7014),
		"Detector Active Time",
		"DetectorActiveTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7016),
		"Detector Activation Offset From Exposure",
		"DetectorActivationOffsetFromExposure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x701A),
		"Detector Binning",
		"DetectorBinning",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7020),
		"Detector Element Physical Size",
		"DetectorElementPhysicalSize",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7022),
		"Detector Element Spacing",
		"DetectorElementSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7024),
		"Detector Active Shape",
		"DetectorActiveShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7026),
		"Detector Active Dimension(s)",
		"DetectorActiveDimensions",
		vm.VM12,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7028),
		"Detector Active Origin",
		"DetectorActiveOrigin",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x702A),
		"Detector Manufacturer Name",
		"DetectorManufacturerName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x702B),
		"Detector Manufacturer's Model Name",
		"DetectorManufacturerModelName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7030),
		"Field of View Origin",
		"FieldOfViewOrigin",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7032),
		"Field of View Rotation",
		"FieldOfViewRotation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7034),
		"Field of View Horizontal Flip",
		"FieldOfViewHorizontalFlip",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7036),
		"Pixel Data Area Origin Relative To FOV",
		"PixelDataAreaOriginRelativeToFOV",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7038),
		"Pixel Data Area Rotation Angle Relative To FOV",
		"PixelDataAreaRotationAngleRelativeToFOV",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7040),
		"Grid Absorbing Material",
		"GridAbsorbingMaterial",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7041),
		"Grid Spacing Material",
		"GridSpacingMaterial",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7042),
		"Grid Thickness",
		"GridThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7044),
		"Grid Pitch",
		"GridPitch",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7046),
		"Grid Aspect Ratio",
		"GridAspectRatio",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7048),
		"Grid Period",
		"GridPeriod",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x704C),
		"Grid Focal Distance",
		"GridFocalDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7050),
		"Filter Material",
		"FilterMaterial",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7052),
		"Filter Thickness Minimum",
		"FilterThicknessMinimum",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7054),
		"Filter Thickness Maximum",
		"FilterThicknessMaximum",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7056),
		"Filter Beam Path Length Minimum",
		"FilterBeamPathLengthMinimum",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7058),
		"Filter Beam Path Length Maximum",
		"FilterBeamPathLengthMaximum",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7060),
		"Exposure Control Mode",
		"ExposureControlMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7062),
		"Exposure Control Mode Description",
		"ExposureControlModeDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7064),
		"Exposure Status",
		"ExposureStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7065),
		"Phototimer Setting",
		"PhototimerSetting",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x8150),
		"Exposure Time in µS",
		"ExposureTimeInuS",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x8151),
		"X-Ray Tube Current in µA",
		"XRayTubeCurrentInuA",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9004),
		"Content Qualification",
		"ContentQualification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9005),
		"Pulse Sequence Name",
		"PulseSequenceName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9006),
		"MR Imaging Modifier Sequence",
		"MRImagingModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9008),
		"Echo Pulse Sequence",
		"EchoPulseSequence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9009),
		"Inversion Recovery",
		"InversionRecovery",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9010),
		"Flow Compensation",
		"FlowCompensation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9011),
		"Multiple Spin Echo",
		"MultipleSpinEcho",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9012),
		"Multi-planar Excitation",
		"MultiPlanarExcitation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9014),
		"Phase Contrast",
		"PhaseContrast",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9015),
		"Time of Flight Contrast",
		"TimeOfFlightContrast",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9016),
		"Spoiling",
		"Spoiling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9017),
		"Steady State Pulse Sequence",
		"SteadyStatePulseSequence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9018),
		"Echo Planar Pulse Sequence",
		"EchoPlanarPulseSequence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9019),
		"Tag Angle First Axis",
		"TagAngleFirstAxis",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9020),
		"Magnetization Transfer",
		"MagnetizationTransfer",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9021),
		"T2 Preparation",
		"T2Preparation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9022),
		"Blood Signal Nulling",
		"BloodSignalNulling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9024),
		"Saturation Recovery",
		"SaturationRecovery",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9025),
		"Spectrally Selected Suppression",
		"SpectrallySelectedSuppression",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9026),
		"Spectrally Selected Excitation",
		"SpectrallySelectedExcitation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9027),
		"Spatial Pre-saturation",
		"SpatialPresaturation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9028),
		"Tagging",
		"Tagging",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9029),
		"Oversampling Phase",
		"OversamplingPhase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9030),
		"Tag Spacing First Dimension",
		"TagSpacingFirstDimension",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9032),
		"Geometry of k-Space Traversal",
		"GeometryOfKSpaceTraversal",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9033),
		"Segmented k-Space Traversal",
		"SegmentedKSpaceTraversal",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9034),
		"Rectilinear Phase Encode Reordering",
		"RectilinearPhaseEncodeReordering",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9035),
		"Tag Thickness",
		"TagThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9036),
		"Partial Fourier Direction",
		"PartialFourierDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9037),
		"Cardiac Synchronization Technique",
		"CardiacSynchronizationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9041),
		"Receive Coil Manufacturer Name",
		"ReceiveCoilManufacturerName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9042),
		"MR Receive Coil Sequence",
		"MRReceiveCoilSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9043),
		"Receive Coil Type",
		"ReceiveCoilType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9044),
		"Quadrature Receive Coil",
		"QuadratureReceiveCoil",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9045),
		"Multi-Coil Definition Sequence",
		"MultiCoilDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9046),
		"Multi-Coil Configuration",
		"MultiCoilConfiguration",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9047),
		"Multi-Coil Element Name",
		"MultiCoilElementName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9048),
		"Multi-Coil Element Used",
		"MultiCoilElementUsed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9049),
		"MR Transmit Coil Sequence",
		"MRTransmitCoilSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9050),
		"Transmit Coil Manufacturer Name",
		"TransmitCoilManufacturerName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9051),
		"Transmit Coil Type",
		"TransmitCoilType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9052),
		"Spectral Width",
		"SpectralWidth",
		vm.VM12,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9053),
		"Chemical Shift Reference",
		"ChemicalShiftReference",
		vm.VM12,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9054),
		"Volume Localization Technique",
		"VolumeLocalizationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9058),
		"MR Acquisition Frequency Encoding Steps",
		"MRAcquisitionFrequencyEncodingSteps",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9059),
		"De-coupling",
		"Decoupling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9060),
		"De-coupled Nucleus",
		"DecoupledNucleus",
		vm.VM12,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9061),
		"De-coupling Frequency",
		"DecouplingFrequency",
		vm.VM12,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9062),
		"De-coupling Method",
		"DecouplingMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9063),
		"De-coupling Chemical Shift Reference",
		"DecouplingChemicalShiftReference",
		vm.VM12,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9064),
		"k-space Filtering",
		"KSpaceFiltering",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9065),
		"Time Domain Filtering",
		"TimeDomainFiltering",
		vm.VM12,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9066),
		"Number of Zero Fills",
		"NumberOfZeroFills",
		vm.VM12,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9067),
		"Baseline Correction",
		"BaselineCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9069),
		"Parallel Reduction Factor In-plane",
		"ParallelReductionFactorInPlane",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9070),
		"Cardiac R-R Interval Specified",
		"CardiacRRIntervalSpecified",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9073),
		"Acquisition Duration",
		"AcquisitionDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9074),
		"Frame Acquisition DateTime",
		"FrameAcquisitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9075),
		"Diffusion Directionality",
		"DiffusionDirectionality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9076),
		"Diffusion Gradient Direction Sequence",
		"DiffusionGradientDirectionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9077),
		"Parallel Acquisition",
		"ParallelAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9078),
		"Parallel Acquisition Technique",
		"ParallelAcquisitionTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9079),
		"Inversion Times",
		"InversionTimes",
		vm.VM1N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9080),
		"Metabolite Map Description",
		"MetaboliteMapDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9081),
		"Partial Fourier",
		"PartialFourier",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9082),
		"Effective Echo Time",
		"EffectiveEchoTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9083),
		"Metabolite Map Code Sequence",
		"MetaboliteMapCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9084),
		"Chemical Shift Sequence",
		"ChemicalShiftSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9085),
		"Cardiac Signal Source",
		"CardiacSignalSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9087),
		"Diffusion b-value",
		"DiffusionBValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9089),
		"Diffusion Gradient Orientation",
		"DiffusionGradientOrientation",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9090),
		"Velocity Encoding Direction",
		"VelocityEncodingDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9091),
		"Velocity Encoding Minimum Value",
		"VelocityEncodingMinimumValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9092),
		"Velocity Encoding Acquisition Sequence",
		"VelocityEncodingAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9093),
		"Number of k-Space Trajectories",
		"NumberOfKSpaceTrajectories",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9094),
		"Coverage of k-Space",
		"CoverageOfKSpace",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9095),
		"Spectroscopy Acquisition Phase Rows",
		"SpectroscopyAcquisitionPhaseRows",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9096),
		"Parallel Reduction Factor In-plane (Retired)",
		"ParallelReductionFactorInPlaneRetired",
		vm.VM1,
		true,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9098),
		"Transmitter Frequency",
		"TransmitterFrequency",
		vm.VM12,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9100),
		"Resonant Nucleus",
		"ResonantNucleus",
		vm.VM12,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9101),
		"Frequency Correction",
		"FrequencyCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9103),
		"MR Spectroscopy FOV/Geometry Sequence",
		"MRSpectroscopyFOVGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9104),
		"Slab Thickness",
		"SlabThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9105),
		"Slab Orientation",
		"SlabOrientation",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9106),
		"Mid Slab Position",
		"MidSlabPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9107),
		"MR Spatial Saturation Sequence",
		"MRSpatialSaturationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9112),
		"MR Timing and Related Parameters Sequence",
		"MRTimingAndRelatedParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9114),
		"MR Echo Sequence",
		"MREchoSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9115),
		"MR Modifier Sequence",
		"MRModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9117),
		"MR Diffusion Sequence",
		"MRDiffusionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9118),
		"Cardiac Synchronization Sequence",
		"CardiacSynchronizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9119),
		"MR Averages Sequence",
		"MRAveragesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9125),
		"MR FOV/Geometry Sequence",
		"MRFOVGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9126),
		"Volume Localization Sequence",
		"VolumeLocalizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9127),
		"Spectroscopy Acquisition Data Columns",
		"SpectroscopyAcquisitionDataColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9147),
		"Diffusion Anisotropy Type",
		"DiffusionAnisotropyType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9151),
		"Frame Reference DateTime",
		"FrameReferenceDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9152),
		"MR Metabolite Map Sequence",
		"MRMetaboliteMapSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9155),
		"Parallel Reduction Factor out-of-plane",
		"ParallelReductionFactorOutOfPlane",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9159),
		"Spectroscopy Acquisition Out-of-plane Phase Steps",
		"SpectroscopyAcquisitionOutOfPlanePhaseSteps",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9166),
		"Bulk Motion Status",
		"BulkMotionStatus",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9168),
		"Parallel Reduction Factor Second In-plane",
		"ParallelReductionFactorSecondInPlane",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9169),
		"Cardiac Beat Rejection Technique",
		"CardiacBeatRejectionTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9170),
		"Respiratory Motion Compensation Technique",
		"RespiratoryMotionCompensationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9171),
		"Respiratory Signal Source",
		"RespiratorySignalSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9172),
		"Bulk Motion Compensation Technique",
		"BulkMotionCompensationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9173),
		"Bulk Motion Signal Source",
		"BulkMotionSignalSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9174),
		"Applicable Safety Standard Agency",
		"ApplicableSafetyStandardAgency",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9175),
		"Applicable Safety Standard Description",
		"ApplicableSafetyStandardDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9176),
		"Operating Mode Sequence",
		"OperatingModeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9177),
		"Operating Mode Type",
		"OperatingModeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9178),
		"Operating Mode",
		"OperatingMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9179),
		"Specific Absorption Rate Definition",
		"SpecificAbsorptionRateDefinition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9180),
		"Gradient Output Type",
		"GradientOutputType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9181),
		"Specific Absorption Rate Value",
		"SpecificAbsorptionRateValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9182),
		"Gradient Output",
		"GradientOutput",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9183),
		"Flow Compensation Direction",
		"FlowCompensationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9184),
		"Tagging Delay",
		"TaggingDelay",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9185),
		"Respiratory Motion Compensation Technique Description",
		"RespiratoryMotionCompensationTechniqueDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9186),
		"Respiratory Signal Source ID",
		"RespiratorySignalSourceID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9195),
		"Chemical Shift Minimum Integration Limit in Hz",
		"ChemicalShiftMinimumIntegrationLimitInHz",
		vm.VM1,
		true,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9196),
		"Chemical Shift Maximum Integration Limit in Hz",
		"ChemicalShiftMaximumIntegrationLimitInHz",
		vm.VM1,
		true,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9197),
		"MR Velocity Encoding Sequence",
		"MRVelocityEncodingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9198),
		"First Order Phase Correction",
		"FirstOrderPhaseCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9199),
		"Water Referenced Phase Correction",
		"WaterReferencedPhaseCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9200),
		"MR Spectroscopy Acquisition Type",
		"MRSpectroscopyAcquisitionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9214),
		"Respiratory Cycle Position",
		"RespiratoryCyclePosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9217),
		"Velocity Encoding Maximum Value",
		"VelocityEncodingMaximumValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9218),
		"Tag Spacing Second Dimension",
		"TagSpacingSecondDimension",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9219),
		"Tag Angle Second Axis",
		"TagAngleSecondAxis",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9220),
		"Frame Acquisition Duration",
		"FrameAcquisitionDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9226),
		"MR Image Frame Type Sequence",
		"MRImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9227),
		"MR Spectroscopy Frame Type Sequence",
		"MRSpectroscopyFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9231),
		"MR Acquisition Phase Encoding Steps in-plane",
		"MRAcquisitionPhaseEncodingStepsInPlane",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9232),
		"MR Acquisition Phase Encoding Steps out-of-plane",
		"MRAcquisitionPhaseEncodingStepsOutOfPlane",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9234),
		"Spectroscopy Acquisition Phase Columns",
		"SpectroscopyAcquisitionPhaseColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9236),
		"Cardiac Cycle Position",
		"CardiacCyclePosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9239),
		"Specific Absorption Rate Sequence",
		"SpecificAbsorptionRateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9240),
		"RF Echo Train Length",
		"RFEchoTrainLength",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9241),
		"Gradient Echo Train Length",
		"GradientEchoTrainLength",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9250),
		"Arterial Spin Labeling Contrast",
		"ArterialSpinLabelingContrast",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9251),
		"MR Arterial Spin Labeling Sequence",
		"MRArterialSpinLabelingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9252),
		"ASL Technique Description",
		"ASLTechniqueDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9253),
		"ASL Slab Number",
		"ASLSlabNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9254),
		"ASL Slab Thickness",
		"ASLSlabThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9255),
		"ASL Slab Orientation",
		"ASLSlabOrientation",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9256),
		"ASL Mid Slab Position",
		"ASLMidSlabPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9257),
		"ASL Context",
		"ASLContext",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9258),
		"ASL Pulse Train Duration",
		"ASLPulseTrainDuration",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9259),
		"ASL Crusher Flag",
		"ASLCrusherFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925A),
		"ASL Crusher Flow Limit",
		"ASLCrusherFlowLimit",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925B),
		"ASL Crusher Description",
		"ASLCrusherDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925C),
		"ASL Bolus Cut-off Flag",
		"ASLBolusCutoffFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925D),
		"ASL Bolus Cut-off Timing Sequence",
		"ASLBolusCutoffTimingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925E),
		"ASL Bolus Cut-off Technique",
		"ASLBolusCutoffTechnique",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925F),
		"ASL Bolus Cut-off Delay Time",
		"ASLBolusCutoffDelayTime",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9260),
		"ASL Slab Sequence",
		"ASLSlabSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9295),
		"Chemical Shift Minimum Integration Limit in ppm",
		"ChemicalShiftMinimumIntegrationLimitInppm",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9296),
		"Chemical Shift Maximum Integration Limit in ppm",
		"ChemicalShiftMaximumIntegrationLimitInppm",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9297),
		"Water Reference Acquisition",
		"WaterReferenceAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9298),
		"Echo Peak Position",
		"EchoPeakPosition",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9301),
		"CT Acquisition Type Sequence",
		"CTAcquisitionTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9302),
		"Acquisition Type",
		"AcquisitionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9303),
		"Tube Angle",
		"TubeAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9304),
		"CT Acquisition Details Sequence",
		"CTAcquisitionDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9305),
		"Revolution Time",
		"RevolutionTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9306),
		"Single Collimation Width",
		"SingleCollimationWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9307),
		"Total Collimation Width",
		"TotalCollimationWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9308),
		"CT Table Dynamics Sequence",
		"CTTableDynamicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9309),
		"Table Speed",
		"TableSpeed",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9310),
		"Table Feed per Rotation",
		"TableFeedPerRotation",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9311),
		"Spiral Pitch Factor",
		"SpiralPitchFactor",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9312),
		"CT Geometry Sequence",
		"CTGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9313),
		"Data Collection Center (Patient)",
		"DataCollectionCenterPatient",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9314),
		"CT Reconstruction Sequence",
		"CTReconstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9315),
		"Reconstruction Algorithm",
		"ReconstructionAlgorithm",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9316),
		"Convolution Kernel Group",
		"ConvolutionKernelGroup",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9317),
		"Reconstruction Field of View",
		"ReconstructionFieldOfView",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9318),
		"Reconstruction Target Center (Patient)",
		"ReconstructionTargetCenterPatient",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9319),
		"Reconstruction Angle",
		"ReconstructionAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9320),
		"Image Filter",
		"ImageFilter",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9321),
		"CT Exposure Sequence",
		"CTExposureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9322),
		"Reconstruction Pixel Spacing",
		"ReconstructionPixelSpacing",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9323),
		"Exposure Modulation Type",
		"ExposureModulationType",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9324),
		"Estimated Dose Saving",
		"EstimatedDoseSaving",
		vm.VM1,
		true,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9325),
		"CT X-Ray Details Sequence",
		"CTXRayDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9326),
		"CT Position Sequence",
		"CTPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9327),
		"Table Position",
		"TablePosition",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9328),
		"Exposure Time in ms",
		"ExposureTimeInms",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9329),
		"CT Image Frame Type Sequence",
		"CTImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9330),
		"X-Ray Tube Current in mA",
		"XRayTubeCurrentInmA",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9332),
		"Exposure in mAs",
		"ExposureInmAs",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9333),
		"Constant Volume Flag",
		"ConstantVolumeFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9334),
		"Fluoroscopy Flag",
		"FluoroscopyFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9335),
		"Distance Source to Data Collection Center",
		"DistanceSourceToDataCollectionCenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9337),
		"Contrast/Bolus Agent Number",
		"ContrastBolusAgentNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9338),
		"Contrast/Bolus Ingredient Code Sequence",
		"ContrastBolusIngredientCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9340),
		"Contrast Administration Profile Sequence",
		"ContrastAdministrationProfileSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9341),
		"Contrast/Bolus Usage Sequence",
		"ContrastBolusUsageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9342),
		"Contrast/Bolus Agent Administered",
		"ContrastBolusAgentAdministered",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9343),
		"Contrast/Bolus Agent Detected",
		"ContrastBolusAgentDetected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9344),
		"Contrast/Bolus Agent Phase",
		"ContrastBolusAgentPhase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9345),
		"CTDIvol",
		"CTDIvol",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9346),
		"CTDI Phantom Type Code Sequence",
		"CTDIPhantomTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9351),
		"Calcium Scoring Mass Factor Patient",
		"CalciumScoringMassFactorPatient",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9352),
		"Calcium Scoring Mass Factor Device",
		"CalciumScoringMassFactorDevice",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9353),
		"Energy Weighting Factor",
		"EnergyWeightingFactor",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9360),
		"CT Additional X-Ray Source Sequence",
		"CTAdditionalXRaySourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9361),
		"Multi-energy CT Acquisition",
		"MultienergyCTAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9362),
		"Multi-energy CT Acquisition Sequence",
		"MultienergyCTAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9363),
		"Multi-energy CT Processing Sequence",
		"MultienergyCTProcessingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9364),
		"Multi-energy CT Characteristics Sequence",
		"MultienergyCTCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9365),
		"Multi-energy CT X-Ray Source Sequence",
		"MultienergyCTXRaySourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9366),
		"X-Ray Source Index",
		"XRaySourceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9367),
		"X-Ray Source ID",
		"XRaySourceID",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9368),
		"Multi-energy Source Technique",
		"MultienergySourceTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9369),
		"Source Start DateTime",
		"SourceStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936A),
		"Source End DateTime",
		"SourceEndDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936B),
		"Switching Phase Number",
		"SwitchingPhaseNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936C),
		"Switching Phase Nominal Duration",
		"SwitchingPhaseNominalDuration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936D),
		"Switching Phase Transition Duration",
		"SwitchingPhaseTransitionDuration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936E),
		"Effective Bin Energy",
		"EffectiveBinEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936F),
		"Multi-energy CT X-Ray Detector Sequence",
		"MultienergyCTXRayDetectorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9370),
		"X-Ray Detector Index",
		"XRayDetectorIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9371),
		"X-Ray Detector ID",
		"XRayDetectorID",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9372),
		"Multi-energy Detector Type",
		"MultienergyDetectorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9373),
		"X-Ray Detector Label",
		"XRayDetectorLabel",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9374),
		"Nominal Max Energy",
		"NominalMaxEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9375),
		"Nominal Min Energy",
		"NominalMinEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9376),
		"Referenced X-Ray Detector Index",
		"ReferencedXRayDetectorIndex",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9377),
		"Referenced X-Ray Source Index",
		"ReferencedXRaySourceIndex",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9378),
		"Referenced Path Index",
		"ReferencedPathIndex",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9379),
		"Multi-energy CT Path Sequence",
		"MultienergyCTPathSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937A),
		"Multi-energy CT Path Index",
		"MultienergyCTPathIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937B),
		"Multi-energy Acquisition Description",
		"MultienergyAcquisitionDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937C),
		"Monoenergetic Energy Equivalent",
		"MonoenergeticEnergyEquivalent",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937D),
		"Material Code Sequence",
		"MaterialCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937E),
		"Decomposition Method",
		"DecompositionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937F),
		"Decomposition Description",
		"DecompositionDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9380),
		"Decomposition Algorithm Identification Sequence",
		"DecompositionAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9381),
		"Decomposition Material Sequence",
		"DecompositionMaterialSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9382),
		"Material Attenuation Sequence",
		"MaterialAttenuationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9383),
		"Photon Energy",
		"PhotonEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9384),
		"X-Ray Mass Attenuation Coefficient",
		"XRayMassAttenuationCoefficient",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9390),
		"Metal Artifact Reduction Sequence",
		"MetalArtifactReductionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9391),
		"Metal Artifact Reduction Applied",
		"MetalArtifactReductionApplied",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9392),
		"Metal Artifact Reduction Algorithm Identification Sequence",
		"MetalArtifactReductionAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9401),
		"Projection Pixel Calibration Sequence",
		"ProjectionPixelCalibrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9402),
		"Distance Source to Isocenter",
		"DistanceSourceToIsocenter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9403),
		"Distance Object to Table Top",
		"DistanceObjectToTableTop",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9404),
		"Object Pixel Spacing in Center of Beam",
		"ObjectPixelSpacingInCenterOfBeam",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9405),
		"Positioner Position Sequence",
		"PositionerPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9406),
		"Table Position Sequence",
		"TablePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9407),
		"Collimator Shape Sequence",
		"CollimatorShapeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9410),
		"Planes in Acquisition",
		"PlanesInAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9412),
		"XA/XRF Frame Characteristics Sequence",
		"XAXRFFrameCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9417),
		"Frame Acquisition Sequence",
		"FrameAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9420),
		"X-Ray Receptor Type",
		"XRayReceptorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9423),
		"Acquisition Protocol Name",
		"AcquisitionProtocolName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9424),
		"Acquisition Protocol Description",
		"AcquisitionProtocolDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9425),
		"Contrast/Bolus Ingredient Opaque",
		"ContrastBolusIngredientOpaque",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9426),
		"Distance Receptor Plane to Detector Housing",
		"DistanceReceptorPlaneToDetectorHousing",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9427),
		"Intensifier Active Shape",
		"IntensifierActiveShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9428),
		"Intensifier Active Dimension(s)",
		"IntensifierActiveDimensions",
		vm.VM12,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9429),
		"Physical Detector Size",
		"PhysicalDetectorSize",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9430),
		"Position of Isocenter Projection",
		"PositionOfIsocenterProjection",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9432),
		"Field of View Sequence",
		"FieldOfViewSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9433),
		"Field of View Description",
		"FieldOfViewDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9434),
		"Exposure Control Sensing Regions Sequence",
		"ExposureControlSensingRegionsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9435),
		"Exposure Control Sensing Region Shape",
		"ExposureControlSensingRegionShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9436),
		"Exposure Control Sensing Region Left Vertical Edge",
		"ExposureControlSensingRegionLeftVerticalEdge",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9437),
		"Exposure Control Sensing Region Right Vertical Edge",
		"ExposureControlSensingRegionRightVerticalEdge",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9438),
		"Exposure Control Sensing Region Upper Horizontal Edge",
		"ExposureControlSensingRegionUpperHorizontalEdge",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9439),
		"Exposure Control Sensing Region Lower Horizontal Edge",
		"ExposureControlSensingRegionLowerHorizontalEdge",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9440),
		"Center of Circular Exposure Control Sensing Region",
		"CenterOfCircularExposureControlSensingRegion",
		vm.VM2,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9441),
		"Radius of Circular Exposure Control Sensing Region",
		"RadiusOfCircularExposureControlSensingRegion",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9442),
		"Vertices of the Polygonal Exposure Control Sensing Region",
		"VerticesOfThePolygonalExposureControlSensingRegion",
		vm.VM2N,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9447),
		"Column Angulation (Patient)",
		"ColumnAngulationPatient",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9449),
		"Beam Angle",
		"BeamAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9451),
		"Frame Detector Parameters Sequence",
		"FrameDetectorParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9452),
		"Calculated Anatomy Thickness",
		"CalculatedAnatomyThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9455),
		"Calibration Sequence",
		"CalibrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9456),
		"Object Thickness Sequence",
		"ObjectThicknessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9457),
		"Plane Identification",
		"PlaneIdentification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9461),
		"Field of View Dimension(s) in Float",
		"FieldOfViewDimensionsInFloat",
		vm.VM12,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9462),
		"Isocenter Reference System Sequence",
		"IsocenterReferenceSystemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9463),
		"Positioner Isocenter Primary Angle",
		"PositionerIsocenterPrimaryAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9464),
		"Positioner Isocenter Secondary Angle",
		"PositionerIsocenterSecondaryAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9465),
		"Positioner Isocenter Detector Rotation Angle",
		"PositionerIsocenterDetectorRotationAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9466),
		"Table X Position to Isocenter",
		"TableXPositionToIsocenter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9467),
		"Table Y Position to Isocenter",
		"TableYPositionToIsocenter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9468),
		"Table Z Position to Isocenter",
		"TableZPositionToIsocenter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9469),
		"Table Horizontal Rotation Angle",
		"TableHorizontalRotationAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9470),
		"Table Head Tilt Angle",
		"TableHeadTiltAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9471),
		"Table Cradle Tilt Angle",
		"TableCradleTiltAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9472),
		"Frame Display Shutter Sequence",
		"FrameDisplayShutterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9473),
		"Acquired Image Area Dose Product",
		"AcquiredImageAreaDoseProduct",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9474),
		"C-arm Positioner Tabletop Relationship",
		"CArmPositionerTabletopRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9476),
		"X-Ray Geometry Sequence",
		"XRayGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9477),
		"Irradiation Event Identification Sequence",
		"IrradiationEventIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9504),
		"X-Ray 3D Frame Type Sequence",
		"XRay3DFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9506),
		"Contributing Sources Sequence",
		"ContributingSourcesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9507),
		"X-Ray 3D Acquisition Sequence",
		"XRay3DAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9508),
		"Primary Positioner Scan Arc",
		"PrimaryPositionerScanArc",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9509),
		"Secondary Positioner Scan Arc",
		"SecondaryPositionerScanArc",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9510),
		"Primary Positioner Scan Start Angle",
		"PrimaryPositionerScanStartAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9511),
		"Secondary Positioner Scan Start Angle",
		"SecondaryPositionerScanStartAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9514),
		"Primary Positioner Increment",
		"PrimaryPositionerIncrement",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9515),
		"Secondary Positioner Increment",
		"SecondaryPositionerIncrement",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9516),
		"Start Acquisition DateTime",
		"StartAcquisitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9517),
		"End Acquisition DateTime",
		"EndAcquisitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9518),
		"Primary Positioner Increment Sign",
		"PrimaryPositionerIncrementSign",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9519),
		"Secondary Positioner Increment Sign",
		"SecondaryPositionerIncrementSign",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9524),
		"Application Name",
		"ApplicationName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9525),
		"Application Version",
		"ApplicationVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9526),
		"Application Manufacturer",
		"ApplicationManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9527),
		"Algorithm Type",
		"AlgorithmType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9528),
		"Algorithm Description",
		"AlgorithmDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9530),
		"X-Ray 3D Reconstruction Sequence",
		"XRay3DReconstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9531),
		"Reconstruction Description",
		"ReconstructionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9538),
		"Per Projection Acquisition Sequence",
		"PerProjectionAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9541),
		"Detector Position Sequence",
		"DetectorPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9542),
		"X-Ray Acquisition Dose Sequence",
		"XRayAcquisitionDoseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9543),
		"X-Ray Source Isocenter Primary Angle",
		"XRaySourceIsocenterPrimaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9544),
		"X-Ray Source Isocenter Secondary Angle",
		"XRaySourceIsocenterSecondaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9545),
		"Breast Support Isocenter Primary Angle",
		"BreastSupportIsocenterPrimaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9546),
		"Breast Support Isocenter Secondary Angle",
		"BreastSupportIsocenterSecondaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9547),
		"Breast Support X Position to Isocenter",
		"BreastSupportXPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9548),
		"Breast Support Y Position to Isocenter",
		"BreastSupportYPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9549),
		"Breast Support Z Position to Isocenter",
		"BreastSupportZPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9550),
		"Detector Isocenter Primary Angle",
		"DetectorIsocenterPrimaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9551),
		"Detector Isocenter Secondary Angle",
		"DetectorIsocenterSecondaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9552),
		"Detector X Position to Isocenter",
		"DetectorXPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9553),
		"Detector Y Position to Isocenter",
		"DetectorYPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9554),
		"Detector Z Position to Isocenter",
		"DetectorZPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9555),
		"X-Ray Grid Sequence",
		"XRayGridSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9556),
		"X-Ray Filter Sequence",
		"XRayFilterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9557),
		"Detector Active Area TLHC Position",
		"DetectorActiveAreaTLHCPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9558),
		"Detector Active Area Orientation",
		"DetectorActiveAreaOrientation",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9559),
		"Positioner Primary Angle Direction",
		"PositionerPrimaryAngleDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9601),
		"Diffusion b-matrix Sequence",
		"DiffusionBMatrixSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9602),
		"Diffusion b-value XX",
		"DiffusionBValueXX",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9603),
		"Diffusion b-value XY",
		"DiffusionBValueXY",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9604),
		"Diffusion b-value XZ",
		"DiffusionBValueXZ",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9605),
		"Diffusion b-value YY",
		"DiffusionBValueYY",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9606),
		"Diffusion b-value YZ",
		"DiffusionBValueYZ",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9607),
		"Diffusion b-value ZZ",
		"DiffusionBValueZZ",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9621),
		"Functional MR Sequence",
		"FunctionalMRSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9622),
		"Functional Settling Phase Frames Present",
		"FunctionalSettlingPhaseFramesPresent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9623),
		"Functional Sync Pulse",
		"FunctionalSyncPulse",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9624),
		"Settling Phase Frame",
		"SettlingPhaseFrame",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9701),
		"Decay Correction DateTime",
		"DecayCorrectionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9715),
		"Start Density Threshold",
		"StartDensityThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9716),
		"Start Relative Density Difference Threshold",
		"StartRelativeDensityDifferenceThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9717),
		"Start Cardiac Trigger Count Threshold",
		"StartCardiacTriggerCountThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9718),
		"Start Respiratory Trigger Count Threshold",
		"StartRespiratoryTriggerCountThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9719),
		"Termination Counts Threshold",
		"TerminationCountsThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9720),
		"Termination Density Threshold",
		"TerminationDensityThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9721),
		"Termination Relative Density Threshold",
		"TerminationRelativeDensityThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9722),
		"Termination Time Threshold",
		"TerminationTimeThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9723),
		"Termination Cardiac Trigger Count Threshold",
		"TerminationCardiacTriggerCountThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9724),
		"Termination Respiratory Trigger Count Threshold",
		"TerminationRespiratoryTriggerCountThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9725),
		"Detector Geometry",
		"DetectorGeometry",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9726),
		"Transverse Detector Separation",
		"TransverseDetectorSeparation",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9727),
		"Axial Detector Dimension",
		"AxialDetectorDimension",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9729),
		"Radiopharmaceutical Agent Number",
		"RadiopharmaceuticalAgentNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9732),
		"PET Frame Acquisition Sequence",
		"PETFrameAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9733),
		"PET Detector Motion Details Sequence",
		"PETDetectorMotionDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9734),
		"PET Table Dynamics Sequence",
		"PETTableDynamicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9735),
		"PET Position Sequence",
		"PETPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9736),
		"PET Frame Correction Factors Sequence",
		"PETFrameCorrectionFactorsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9737),
		"Radiopharmaceutical Usage Sequence",
		"RadiopharmaceuticalUsageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9738),
		"Attenuation Correction Source",
		"AttenuationCorrectionSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9739),
		"Number of Iterations",
		"NumberOfIterations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9740),
		"Number of Subsets",
		"NumberOfSubsets",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9749),
		"PET Reconstruction Sequence",
		"PETReconstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9751),
		"PET Frame Type Sequence",
		"PETFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9755),
		"Time of Flight Information Used",
		"TimeOfFlightInformationUsed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9756),
		"Reconstruction Type",
		"ReconstructionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9758),
		"Decay Corrected",
		"DecayCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9759),
		"Attenuation Corrected",
		"AttenuationCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9760),
		"Scatter Corrected",
		"ScatterCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9761),
		"Dead Time Corrected",
		"DeadTimeCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9762),
		"Gantry Motion Corrected",
		"GantryMotionCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9763),
		"Patient Motion Corrected",
		"PatientMotionCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9764),
		"Count Loss Normalization Corrected",
		"CountLossNormalizationCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9765),
		"Randoms Corrected",
		"RandomsCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9766),
		"Non-uniform Radial Sampling Corrected",
		"NonUniformRadialSamplingCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9767),
		"Sensitivity Calibrated",
		"SensitivityCalibrated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9768),
		"Detector Normalization Correction",
		"DetectorNormalizationCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9769),
		"Iterative Reconstruction Method",
		"IterativeReconstructionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9770),
		"Attenuation Correction Temporal Relationship",
		"AttenuationCorrectionTemporalRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9771),
		"Patient Physiological State Sequence",
		"PatientPhysiologicalStateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9772),
		"Patient Physiological State Code Sequence",
		"PatientPhysiologicalStateCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9801),
		"Depth(s) of Focus",
		"DepthsOfFocus",
		vm.VM1N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9803),
		"Excluded Intervals Sequence",
		"ExcludedIntervalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9804),
		"Exclusion Start DateTime",
		"ExclusionStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9805),
		"Exclusion Duration",
		"ExclusionDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9806),
		"US Image Description Sequence",
		"USImageDescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9807),
		"Image Data Type Sequence",
		"ImageDataTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9808),
		"Data Type",
		"DataType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9809),
		"Transducer Scan Pattern Code Sequence",
		"TransducerScanPatternCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980B),
		"Aliased Data Type",
		"AliasedDataType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980C),
		"Position Measuring Device Used",
		"PositionMeasuringDeviceUsed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980D),
		"Transducer Geometry Code Sequence",
		"TransducerGeometryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980E),
		"Transducer Beam Steering Code Sequence",
		"TransducerBeamSteeringCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980F),
		"Transducer Application Code Sequence",
		"TransducerApplicationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9810),
		"Zero Velocity Pixel Value",
		"ZeroVelocityPixelValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9821),
		"Photoacoustic Excitation Characteristics Sequence",
		"PhotoacousticExcitationCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9822),
		"Excitation Spectral Width",
		"ExcitationSpectralWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9823),
		"Excitation Energy",
		"ExcitationEnergy",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9824),
		"Excitation Pulse Duration",
		"ExcitationPulseDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9825),
		"Excitation Wavelength Sequence",
		"ExcitationWavelengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9826),
		"Excitation Wavelength",
		"ExcitationWavelength",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9828),
		"Illumination Translation Flag",
		"IlluminationTranslationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9829),
		"Acoustic Coupling Medium Flag",
		"AcousticCouplingMediumFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982A),
		"Acoustic Coupling Medium Code Sequence",
		"AcousticCouplingMediumCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982B),
		"Acoustic Coupling Medium Temperature",
		"AcousticCouplingMediumTemperature",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982C),
		"Transducer Response Sequence",
		"TransducerResponseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982D),
		"Center Frequency",
		"CenterFrequency",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982E),
		"Fractional Bandwidth",
		"FractionalBandwidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982F),
		"Lower Cutoff Frequency",
		"LowerCutoffFrequency",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9830),
		"Upper Cutoff Frequency",
		"UpperCutoffFrequency",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9831),
		"Transducer Technology Sequence",
		"TransducerTechnologySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9832),
		"Sound Speed Correction Mechanism Code Sequence",
		"SoundSpeedCorrectionMechanismCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9833),
		"Object Sound Speed",
		"ObjectSoundSpeed",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9834),
		"Acoustic Coupling Medium Sound Speed",
		"AcousticCouplingMediumSoundSpeed",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9835),
		"Photoacoustic Image Frame Type Sequence",
		"PhotoacousticImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9836),
		"Image Data Type Code Sequence",
		"ImageDataTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9900),
		"Reference Location Label",
		"ReferenceLocationLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9901),
		"Reference Location Description",
		"ReferenceLocationDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9902),
		"Reference Basis Code Sequence",
		"ReferenceBasisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9903),
		"Reference Geometry Code Sequence",
		"ReferenceGeometryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9904),
		"Offset Distance",
		"OffsetDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9905),
		"Offset Direction",
		"OffsetDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9906),
		"Potential Scheduled Protocol Code Sequence",
		"PotentialScheduledProtocolCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9907),
		"Potential Requested Procedure Code Sequence",
		"PotentialRequestedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9908),
		"Potential Reasons for Procedure",
		"PotentialReasonsForProcedure",
		vm.VM1N,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9909),
		"Potential Reasons for Procedure Code Sequence",
		"PotentialReasonsForProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990A),
		"Potential Diagnostic Tasks",
		"PotentialDiagnosticTasks",
		vm.VM1N,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990B),
		"Contraindications Code Sequence",
		"ContraindicationsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990C),
		"Referenced Defined Protocol Sequence",
		"ReferencedDefinedProtocolSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990D),
		"Referenced Performed Protocol Sequence",
		"ReferencedPerformedProtocolSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990E),
		"Predecessor Protocol Sequence",
		"PredecessorProtocolSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990F),
		"Protocol Planning Information",
		"ProtocolPlanningInformation",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9910),
		"Protocol Design Rationale",
		"ProtocolDesignRationale",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9911),
		"Patient Specification Sequence",
		"PatientSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9912),
		"Model Specification Sequence",
		"ModelSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9913),
		"Parameters Specification Sequence",
		"ParametersSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9914),
		"Instruction Sequence",
		"InstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9915),
		"Instruction Index",
		"InstructionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9916),
		"Instruction Text",
		"InstructionText",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9917),
		"Instruction Description",
		"InstructionDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9918),
		"Instruction Performed Flag",
		"InstructionPerformedFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9919),
		"Instruction Performed DateTime",
		"InstructionPerformedDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991A),
		"Instruction Performance Comment",
		"InstructionPerformanceComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991B),
		"Patient Positioning Instruction Sequence",
		"PatientPositioningInstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991C),
		"Positioning Method Code Sequence",
		"PositioningMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991D),
		"Positioning Landmark Sequence",
		"PositioningLandmarkSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991E),
		"Target Frame of Reference UID",
		"TargetFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991F),
		"Acquisition Protocol Element Specification Sequence",
		"AcquisitionProtocolElementSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9920),
		"Acquisition Protocol Element Sequence",
		"AcquisitionProtocolElementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9921),
		"Protocol Element Number",
		"ProtocolElementNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9922),
		"Protocol Element Name",
		"ProtocolElementName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9923),
		"Protocol Element Characteristics Summary",
		"ProtocolElementCharacteristicsSummary",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9924),
		"Protocol Element Purpose",
		"ProtocolElementPurpose",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9930),
		"Acquisition Motion",
		"AcquisitionMotion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9931),
		"Acquisition Start Location Sequence",
		"AcquisitionStartLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9932),
		"Acquisition End Location Sequence",
		"AcquisitionEndLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9933),
		"Reconstruction Protocol Element Specification Sequence",
		"ReconstructionProtocolElementSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9934),
		"Reconstruction Protocol Element Sequence",
		"ReconstructionProtocolElementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9935),
		"Storage Protocol Element Specification Sequence",
		"StorageProtocolElementSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9936),
		"Storage Protocol Element Sequence",
		"StorageProtocolElementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9937),
		"Requested Series Description",
		"RequestedSeriesDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9938),
		"Source Acquisition Protocol Element Number",
		"SourceAcquisitionProtocolElementNumber",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9939),
		"Source Acquisition Beam Number",
		"SourceAcquisitionBeamNumber",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993A),
		"Source Reconstruction Protocol Element Number",
		"SourceReconstructionProtocolElementNumber",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993B),
		"Reconstruction Start Location Sequence",
		"ReconstructionStartLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993C),
		"Reconstruction End Location Sequence",
		"ReconstructionEndLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993D),
		"Reconstruction Algorithm Sequence",
		"ReconstructionAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993E),
		"Reconstruction Target Center Location Sequence",
		"ReconstructionTargetCenterLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9941),
		"Image Filter Description",
		"ImageFilterDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9942),
		"CTDIvol Notification Trigger",
		"CTDIvolNotificationTrigger",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9943),
		"DLP Notification Trigger",
		"DLPNotificationTrigger",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9944),
		"Auto KVP Selection Type",
		"AutoKVPSelectionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9945),
		"Auto KVP Upper Bound",
		"AutoKVPUpperBound",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9946),
		"Auto KVP Lower Bound",
		"AutoKVPLowerBound",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9947),
		"Protocol Defined Patient Position",
		"ProtocolDefinedPatientPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0xA001),
		"Contributing Equipment Sequence",
		"ContributingEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0xA002),
		"Contribution DateTime",
		"ContributionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0xA003),
		"Contribution Description",
		"ContributionDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x000D),
		"Study Instance UID",
		"StudyInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x000E),
		"Series Instance UID",
		"SeriesInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0010),
		"Study ID",
		"StudyID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0011),
		"Series Number",
		"SeriesNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0012),
		"Acquisition Number",
		"AcquisitionNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0013),
		"Instance Number",
		"InstanceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0014),
		"Isotope Number",
		"IsotopeNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0015),
		"Phase Number",
		"PhaseNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0016),
		"Interval Number",
		"IntervalNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0017),
		"Time Slot Number",
		"TimeSlotNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0018),
		"Angle Number",
		"AngleNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0019),
		"Item Number",
		"ItemNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0020),
		"Patient Orientation",
		"PatientOrientation",
		vm.VM2,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0022),
		"Overlay Number",
		"OverlayNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0024),
		"Curve Number",
		"CurveNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0026),
		"LUT Number",
		"LUTNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0027),
		"Pyramid Label",
		"PyramidLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0030),
		"Image Position",
		"ImagePosition",
		vm.VM3,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0032),
		"Image Position (Patient)",
		"ImagePositionPatient",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0035),
		"Image Orientation",
		"ImageOrientation",
		vm.VM6,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0037),
		"Image Orientation (Patient)",
		"ImageOrientationPatient",
		vm.VM6,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0050),
		"Location",
		"Location",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0052),
		"Frame of Reference UID",
		"FrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0060),
		"Laterality",
		"Laterality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0062),
		"Image Laterality",
		"ImageLaterality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0070),
		"Image Geometry Type",
		"ImageGeometryType",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0080),
		"Masking Image",
		"MaskingImage",
		vm.VM1N,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x00AA),
		"Report Number",
		"ReportNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0100),
		"Temporal Position Identifier",
		"TemporalPositionIdentifier",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0105),
		"Number of Temporal Positions",
		"NumberOfTemporalPositions",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0110),
		"Temporal Resolution",
		"TemporalResolution",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0200),
		"Synchronization Frame of Reference UID",
		"SynchronizationFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0242),
		"SOP Instance UID of Concatenation Source",
		"SOPInstanceUIDOfConcatenationSource",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1000),
		"Series in Study",
		"SeriesInStudy",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1001),
		"Acquisitions in Series",
		"AcquisitionsInSeries",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1002),
		"Images in Acquisition",
		"ImagesInAcquisition",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1003),
		"Images in Series",
		"ImagesInSeries",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1004),
		"Acquisitions in Study",
		"AcquisitionsInStudy",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1005),
		"Images in Study",
		"ImagesInStudy",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1020),
		"Reference",
		"Reference",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x103F),
		"Target Position Reference Indicator",
		"TargetPositionReferenceIndicator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1040),
		"Position Reference Indicator",
		"PositionReferenceIndicator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1041),
		"Slice Location",
		"SliceLocation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1070),
		"Other Study Numbers",
		"OtherStudyNumbers",
		vm.VM1N,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1200),
		"Number of Patient Related Studies",
		"NumberOfPatientRelatedStudies",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1202),
		"Number of Patient Related Series",
		"NumberOfPatientRelatedSeries",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1204),
		"Number of Patient Related Instances",
		"NumberOfPatientRelatedInstances",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1206),
		"Number of Study Related Series",
		"NumberOfStudyRelatedSeries",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1208),
		"Number of Study Related Instances",
		"NumberOfStudyRelatedInstances",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1209),
		"Number of Series Related Instances",
		"NumberOfSeriesRelatedInstances",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0020,31xx)"),
		"Source Image IDs",
		"SourceImageIDs",
		vm.VM1N,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3401),
		"Modifying Device ID",
		"ModifyingDeviceID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3402),
		"Modified Image ID",
		"ModifiedImageID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3403),
		"Modified Image Date",
		"ModifiedImageDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3404),
		"Modifying Device Manufacturer",
		"ModifyingDeviceManufacturer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3405),
		"Modified Image Time",
		"ModifiedImageTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3406),
		"Modified Image Description",
		"ModifiedImageDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x4000),
		"Image Comments",
		"ImageComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x5000),
		"Original Image Identification",
		"OriginalImageIdentification",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x5002),
		"Original Image Identification Nomenclature",
		"OriginalImageIdentificationNomenclature",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9056),
		"Stack ID",
		"StackID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9057),
		"In-Stack Position Number",
		"InStackPositionNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9071),
		"Frame Anatomy Sequence",
		"FrameAnatomySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9072),
		"Frame Laterality",
		"FrameLaterality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9111),
		"Frame Content Sequence",
		"FrameContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9113),
		"Plane Position Sequence",
		"PlanePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9116),
		"Plane Orientation Sequence",
		"PlaneOrientationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9128),
		"Temporal Position Index",
		"TemporalPositionIndex",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9153),
		"Nominal Cardiac Trigger Delay Time",
		"NominalCardiacTriggerDelayTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9154),
		"Nominal Cardiac Trigger Time Prior To R-Peak",
		"NominalCardiacTriggerTimePriorToRPeak",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9155),
		"Actual Cardiac Trigger Time Prior To R-Peak",
		"ActualCardiacTriggerTimePriorToRPeak",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9156),
		"Frame Acquisition Number",
		"FrameAcquisitionNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9157),
		"Dimension Index Values",
		"DimensionIndexValues",
		vm.VM1N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9158),
		"Frame Comments",
		"FrameComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9161),
		"Concatenation UID",
		"ConcatenationUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9162),
		"In-concatenation Number",
		"InConcatenationNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9163),
		"In-concatenation Total Number",
		"InConcatenationTotalNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9164),
		"Dimension Organization UID",
		"DimensionOrganizationUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9165),
		"Dimension Index Pointer",
		"DimensionIndexPointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9167),
		"Functional Group Pointer",
		"FunctionalGroupPointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9170),
		"Unassigned Shared Converted Attributes Sequence",
		"UnassignedSharedConvertedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9171),
		"Unassigned Per-Frame Converted Attributes Sequence",
		"UnassignedPerFrameConvertedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9172),
		"Conversion Source Attributes Sequence",
		"ConversionSourceAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9213),
		"Dimension Index Private Creator",
		"DimensionIndexPrivateCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9221),
		"Dimension Organization Sequence",
		"DimensionOrganizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9222),
		"Dimension Index Sequence",
		"DimensionIndexSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9228),
		"Concatenation Frame Offset Number",
		"ConcatenationFrameOffsetNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9238),
		"Functional Group Private Creator",
		"FunctionalGroupPrivateCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9241),
		"Nominal Percentage of Cardiac Phase",
		"NominalPercentageOfCardiacPhase",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9245),
		"Nominal Percentage of Respiratory Phase",
		"NominalPercentageOfRespiratoryPhase",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9246),
		"Starting Respiratory Amplitude",
		"StartingRespiratoryAmplitude",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9247),
		"Starting Respiratory Phase",
		"StartingRespiratoryPhase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9248),
		"Ending Respiratory Amplitude",
		"EndingRespiratoryAmplitude",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9249),
		"Ending Respiratory Phase",
		"EndingRespiratoryPhase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9250),
		"Respiratory Trigger Type",
		"RespiratoryTriggerType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9251),
		"R-R Interval Time Nominal",
		"RRIntervalTimeNominal",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9252),
		"Actual Cardiac Trigger Delay Time",
		"ActualCardiacTriggerDelayTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9253),
		"Respiratory Synchronization Sequence",
		"RespiratorySynchronizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9254),
		"Respiratory Interval Time",
		"RespiratoryIntervalTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9255),
		"Nominal Respiratory Trigger Delay Time",
		"NominalRespiratoryTriggerDelayTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9256),
		"Respiratory Trigger Delay Threshold",
		"RespiratoryTriggerDelayThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9257),
		"Actual Respiratory Trigger Delay Time",
		"ActualRespiratoryTriggerDelayTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9301),
		"Image Position (Volume)",
		"ImagePositionVolume",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9302),
		"Image Orientation (Volume)",
		"ImageOrientationVolume",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9307),
		"Ultrasound Acquisition Geometry",
		"UltrasoundAcquisitionGeometry",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9308),
		"Apex Position",
		"ApexPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9309),
		"Volume to Transducer Mapping Matrix",
		"VolumeToTransducerMappingMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930A),
		"Volume to Table Mapping Matrix",
		"VolumeToTableMappingMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930B),
		"Volume to Transducer Relationship",
		"VolumeToTransducerRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930C),
		"Patient Frame of Reference Source",
		"PatientFrameOfReferenceSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930D),
		"Temporal Position Time Offset",
		"TemporalPositionTimeOffset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930E),
		"Plane Position (Volume) Sequence",
		"PlanePositionVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930F),
		"Plane Orientation (Volume) Sequence",
		"PlaneOrientationVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9310),
		"Temporal Position Sequence",
		"TemporalPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9311),
		"Dimension Organization Type",
		"DimensionOrganizationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9312),
		"Volume Frame of Reference UID",
		"VolumeFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9313),
		"Table Frame of Reference UID",
		"TableFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9421),
		"Dimension Description Label",
		"DimensionDescriptionLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9450),
		"Patient Orientation in Frame Sequence",
		"PatientOrientationInFrameSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9453),
		"Frame Label",
		"FrameLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9518),
		"Acquisition Index",
		"AcquisitionIndex",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9529),
		"Contributing SOP Instances Reference Sequence",
		"ContributingSOPInstancesReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9536),
		"Reconstruction Index",
		"ReconstructionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0001),
		"Light Path Filter Pass-Through Wavelength",
		"LightPathFilterPassThroughWavelength",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0002),
		"Light Path Filter Pass Band",
		"LightPathFilterPassBand",
		vm.VM2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0003),
		"Image Path Filter Pass-Through Wavelength",
		"ImagePathFilterPassThroughWavelength",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0004),
		"Image Path Filter Pass Band",
		"ImagePathFilterPassBand",
		vm.VM2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0005),
		"Patient Eye Movement Commanded",
		"PatientEyeMovementCommanded",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0006),
		"Patient Eye Movement Command Code Sequence",
		"PatientEyeMovementCommandCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0007),
		"Spherical Lens Power",
		"SphericalLensPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0008),
		"Cylinder Lens Power",
		"CylinderLensPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0009),
		"Cylinder Axis",
		"CylinderAxis",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000A),
		"Emmetropic Magnification",
		"EmmetropicMagnification",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000B),
		"Intra Ocular Pressure",
		"IntraOcularPressure",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000C),
		"Horizontal Field of View",
		"HorizontalFieldOfView",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000D),
		"Pupil Dilated",
		"PupilDilated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000E),
		"Degree of Dilation",
		"DegreeOfDilation",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000F),
		"Vertex Distance",
		"VertexDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0010),
		"Stereo Baseline Angle",
		"StereoBaselineAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0011),
		"Stereo Baseline Displacement",
		"StereoBaselineDisplacement",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0012),
		"Stereo Horizontal Pixel Offset",
		"StereoHorizontalPixelOffset",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0013),
		"Stereo Vertical Pixel Offset",
		"StereoVerticalPixelOffset",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0014),
		"Stereo Rotation",
		"StereoRotation",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0015),
		"Acquisition Device Type Code Sequence",
		"AcquisitionDeviceTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0016),
		"Illumination Type Code Sequence",
		"IlluminationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0017),
		"Light Path Filter Type Stack Code Sequence",
		"LightPathFilterTypeStackCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0018),
		"Image Path Filter Type Stack Code Sequence",
		"ImagePathFilterTypeStackCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0019),
		"Lenses Code Sequence",
		"LensesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001A),
		"Channel Description Code Sequence",
		"ChannelDescriptionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001B),
		"Refractive State Sequence",
		"RefractiveStateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001C),
		"Mydriatic Agent Code Sequence",
		"MydriaticAgentCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001D),
		"Relative Image Position Code Sequence",
		"RelativeImagePositionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001E),
		"Camera Angle of View",
		"CameraAngleOfView",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0020),
		"Stereo Pairs Sequence",
		"StereoPairsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0021),
		"Left Image Sequence",
		"LeftImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0022),
		"Right Image Sequence",
		"RightImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0028),
		"Stereo Pairs Present",
		"StereoPairsPresent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0030),
		"Axial Length of the Eye",
		"AxialLengthOfTheEye",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0031),
		"Ophthalmic Frame Location Sequence",
		"OphthalmicFrameLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0032),
		"Reference Coordinates",
		"ReferenceCoordinates",
		vm.VM22N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0035),
		"Depth Spatial Resolution",
		"DepthSpatialResolution",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0036),
		"Maximum Depth Distortion",
		"MaximumDepthDistortion",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0037),
		"Along-scan Spatial Resolution",
		"AlongScanSpatialResolution",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0038),
		"Maximum Along-scan Distortion",
		"MaximumAlongScanDistortion",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0039),
		"Ophthalmic Image Orientation",
		"OphthalmicImageOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0041),
		"Depth of Transverse Image",
		"DepthOfTransverseImage",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0042),
		"Mydriatic Agent Concentration Units Sequence",
		"MydriaticAgentConcentrationUnitsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0048),
		"Across-scan Spatial Resolution",
		"AcrossScanSpatialResolution",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0049),
		"Maximum Across-scan Distortion",
		"MaximumAcrossScanDistortion",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x004E),
		"Mydriatic Agent Concentration",
		"MydriaticAgentConcentration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0055),
		"Illumination Wave Length",
		"IlluminationWaveLength",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0056),
		"Illumination Power",
		"IlluminationPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0057),
		"Illumination Bandwidth",
		"IlluminationBandwidth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0058),
		"Mydriatic Agent Sequence",
		"MydriaticAgentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1007),
		"Ophthalmic Axial Measurements Right Eye Sequence",
		"OphthalmicAxialMeasurementsRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1008),
		"Ophthalmic Axial Measurements Left Eye Sequence",
		"OphthalmicAxialMeasurementsLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1009),
		"Ophthalmic Axial Measurements Device Type",
		"OphthalmicAxialMeasurementsDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1010),
		"Ophthalmic Axial Length Measurements Type",
		"OphthalmicAxialLengthMeasurementsType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1012),
		"Ophthalmic Axial Length Sequence",
		"OphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1019),
		"Ophthalmic Axial Length",
		"OphthalmicAxialLength",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1024),
		"Lens Status Code Sequence",
		"LensStatusCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1025),
		"Vitreous Status Code Sequence",
		"VitreousStatusCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1028),
		"IOL Formula Code Sequence",
		"IOLFormulaCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1029),
		"IOL Formula Detail",
		"IOLFormulaDetail",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1033),
		"Keratometer Index",
		"KeratometerIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1035),
		"Source of Ophthalmic Axial Length Code Sequence",
		"SourceOfOphthalmicAxialLengthCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1036),
		"Source of Corneal Size Data Code Sequence",
		"SourceOfCornealSizeDataCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1037),
		"Target Refraction",
		"TargetRefraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1039),
		"Refractive Procedure Occurred",
		"RefractiveProcedureOccurred",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1040),
		"Refractive Surgery Type Code Sequence",
		"RefractiveSurgeryTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1044),
		"Ophthalmic Ultrasound Method Code Sequence",
		"OphthalmicUltrasoundMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1045),
		"Surgically Induced Astigmatism Sequence",
		"SurgicallyInducedAstigmatismSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1046),
		"Type of Optical Correction",
		"TypeOfOpticalCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1047),
		"Toric IOL Power Sequence",
		"ToricIOLPowerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1048),
		"Predicted Toric Error Sequence",
		"PredictedToricErrorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1049),
		"Pre-Selected for Implantation",
		"PreSelectedForImplantation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x104A),
		"Toric IOL Power for Exact Emmetropia Sequence",
		"ToricIOLPowerForExactEmmetropiaSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x104B),
		"Toric IOL Power for Exact Target Refraction Sequence",
		"ToricIOLPowerForExactTargetRefractionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1050),
		"Ophthalmic Axial Length Measurements Sequence",
		"OphthalmicAxialLengthMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1053),
		"IOL Power",
		"IOLPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1054),
		"Predicted Refractive Error",
		"PredictedRefractiveError",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1059),
		"Ophthalmic Axial Length Velocity",
		"OphthalmicAxialLengthVelocity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1065),
		"Lens Status Description",
		"LensStatusDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1066),
		"Vitreous Status Description",
		"VitreousStatusDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1090),
		"IOL Power Sequence",
		"IOLPowerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1092),
		"Lens Constant Sequence",
		"LensConstantSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1093),
		"IOL Manufacturer",
		"IOLManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1094),
		"Lens Constant Description",
		"LensConstantDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1095),
		"Implant Name",
		"ImplantName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1096),
		"Keratometry Measurement Type Code Sequence",
		"KeratometryMeasurementTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1097),
		"Implant Part Number",
		"ImplantPartNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1100),
		"Referenced Ophthalmic Axial Measurements Sequence",
		"ReferencedOphthalmicAxialMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1101),
		"Ophthalmic Axial Length Measurements Segment Name Code Sequence",
		"OphthalmicAxialLengthMeasurementsSegmentNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1103),
		"Refractive Error Before Refractive Surgery Code Sequence",
		"RefractiveErrorBeforeRefractiveSurgeryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1121),
		"IOL Power For Exact Emmetropia",
		"IOLPowerForExactEmmetropia",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1122),
		"IOL Power For Exact Target Refraction",
		"IOLPowerForExactTargetRefraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1125),
		"Anterior Chamber Depth Definition Code Sequence",
		"AnteriorChamberDepthDefinitionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1127),
		"Lens Thickness Sequence",
		"LensThicknessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1128),
		"Anterior Chamber Depth Sequence",
		"AnteriorChamberDepthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x112A),
		"Calculation Comment Sequence",
		"CalculationCommentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x112B),
		"Calculation Comment Type",
		"CalculationCommentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x112C),
		"Calculation Comment",
		"CalculationComment",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1130),
		"Lens Thickness",
		"LensThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1131),
		"Anterior Chamber Depth",
		"AnteriorChamberDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1132),
		"Source of Lens Thickness Data Code Sequence",
		"SourceOfLensThicknessDataCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1133),
		"Source of Anterior Chamber Depth Data Code Sequence",
		"SourceOfAnteriorChamberDepthDataCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1134),
		"Source of Refractive Measurements Sequence",
		"SourceOfRefractiveMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1135),
		"Source of Refractive Measurements Code Sequence",
		"SourceOfRefractiveMeasurementsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1140),
		"Ophthalmic Axial Length Measurement Modified",
		"OphthalmicAxialLengthMeasurementModified",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1150),
		"Ophthalmic Axial Length Data Source Code Sequence",
		"OphthalmicAxialLengthDataSourceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1153),
		"Ophthalmic Axial Length Acquisition Method Code Sequence",
		"OphthalmicAxialLengthAcquisitionMethodCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1155),
		"Signal to Noise Ratio",
		"SignalToNoiseRatio",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1159),
		"Ophthalmic Axial Length Data Source Description",
		"OphthalmicAxialLengthDataSourceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1210),
		"Ophthalmic Axial Length Measurements Total Length Sequence",
		"OphthalmicAxialLengthMeasurementsTotalLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1211),
		"Ophthalmic Axial Length Measurements Segmental Length Sequence",
		"OphthalmicAxialLengthMeasurementsSegmentalLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1212),
		"Ophthalmic Axial Length Measurements Length Summation Sequence",
		"OphthalmicAxialLengthMeasurementsLengthSummationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1220),
		"Ultrasound Ophthalmic Axial Length Measurements Sequence",
		"UltrasoundOphthalmicAxialLengthMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1225),
		"Optical Ophthalmic Axial Length Measurements Sequence",
		"OpticalOphthalmicAxialLengthMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1230),
		"Ultrasound Selected Ophthalmic Axial Length Sequence",
		"UltrasoundSelectedOphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1250),
		"Ophthalmic Axial Length Selection Method Code Sequence",
		"OphthalmicAxialLengthSelectionMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1255),
		"Optical Selected Ophthalmic Axial Length Sequence",
		"OpticalSelectedOphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1257),
		"Selected Segmental Ophthalmic Axial Length Sequence",
		"SelectedSegmentalOphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1260),
		"Selected Total Ophthalmic Axial Length Sequence",
		"SelectedTotalOphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1262),
		"Ophthalmic Axial Length Quality Metric Sequence",
		"OphthalmicAxialLengthQualityMetricSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1265),
		"Ophthalmic Axial Length Quality Metric Type Code Sequence",
		"OphthalmicAxialLengthQualityMetricTypeCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1273),
		"Ophthalmic Axial Length Quality Metric Type Description",
		"OphthalmicAxialLengthQualityMetricTypeDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1300),
		"Intraocular Lens Calculations Right Eye Sequence",
		"IntraocularLensCalculationsRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1310),
		"Intraocular Lens Calculations Left Eye Sequence",
		"IntraocularLensCalculationsLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1330),
		"Referenced Ophthalmic Axial Length Measurement QC Image Sequence",
		"ReferencedOphthalmicAxialLengthMeasurementQCImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1415),
		"Ophthalmic Mapping Device Type",
		"OphthalmicMappingDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1420),
		"Acquisition Method Code Sequence",
		"AcquisitionMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1423),
		"Acquisition Method Algorithm Sequence",
		"AcquisitionMethodAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1436),
		"Ophthalmic Thickness Map Type Code Sequence",
		"OphthalmicThicknessMapTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1443),
		"Ophthalmic Thickness Mapping Normals Sequence",
		"OphthalmicThicknessMappingNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1445),
		"Retinal Thickness Definition Code Sequence",
		"RetinalThicknessDefinitionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1450),
		"Pixel Value Mapping to Coded Concept Sequence",
		"PixelValueMappingToCodedConceptSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1452),
		"Mapped Pixel Value",
		"MappedPixelValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1454),
		"Pixel Value Mapping Explanation",
		"PixelValueMappingExplanation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1458),
		"Ophthalmic Thickness Map Quality Threshold Sequence",
		"OphthalmicThicknessMapQualityThresholdSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1460),
		"Ophthalmic Thickness Map Threshold Quality Rating",
		"OphthalmicThicknessMapThresholdQualityRating",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1463),
		"Anatomic Structure Reference Point",
		"AnatomicStructureReferencePoint",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1465),
		"Registration to Localizer Sequence",
		"RegistrationToLocalizerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1466),
		"Registered Localizer Units",
		"RegisteredLocalizerUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1467),
		"Registered Localizer Top Left Hand Corner",
		"RegisteredLocalizerTopLeftHandCorner",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1468),
		"Registered Localizer Bottom Right Hand Corner",
		"RegisteredLocalizerBottomRightHandCorner",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1470),
		"Ophthalmic Thickness Map Quality Rating Sequence",
		"OphthalmicThicknessMapQualityRatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1472),
		"Relevant OPT Attributes Sequence",
		"RelevantOPTAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1512),
		"Transformation Method Code Sequence",
		"TransformationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1513),
		"Transformation Algorithm Sequence",
		"TransformationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1515),
		"Ophthalmic Axial Length Method",
		"OphthalmicAxialLengthMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1517),
		"Ophthalmic FOV",
		"OphthalmicFOV",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1518),
		"Two Dimensional to Three Dimensional Map Sequence",
		"TwoDimensionalToThreeDimensionalMapSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1525),
		"Wide Field Ophthalmic Photography Quality Rating Sequence",
		"WideFieldOphthalmicPhotographyQualityRatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1526),
		"Wide Field Ophthalmic Photography Quality Threshold Sequence",
		"WideFieldOphthalmicPhotographyQualityThresholdSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1527),
		"Wide Field Ophthalmic Photography Threshold Quality Rating",
		"WideFieldOphthalmicPhotographyThresholdQualityRating",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1528),
		"X Coordinates Center Pixel View Angle",
		"XCoordinatesCenterPixelViewAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1529),
		"Y Coordinates Center Pixel View Angle",
		"YCoordinatesCenterPixelViewAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1530),
		"Number of Map Points",
		"NumberOfMapPoints",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1531),
		"Two Dimensional to Three Dimensional Map Data",
		"TwoDimensionalToThreeDimensionalMapData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1612),
		"Derivation Algorithm Sequence",
		"DerivationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1615),
		"Ophthalmic Image Type Code Sequence",
		"OphthalmicImageTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1616),
		"Ophthalmic Image Type Description",
		"OphthalmicImageTypeDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1618),
		"Scan Pattern Type Code Sequence",
		"ScanPatternTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1620),
		"Referenced Surface Mesh Identification Sequence",
		"ReferencedSurfaceMeshIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1622),
		"Ophthalmic Volumetric Properties Flag",
		"OphthalmicVolumetricPropertiesFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1623),
		"Ophthalmic Anatomic Reference Point Frame Coordinate",
		"OphthalmicAnatomicReferencePointFrameCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1624),
		"Ophthalmic Anatomic Reference Point X-Coordinate",
		"OphthalmicAnatomicReferencePointXCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1626),
		"Ophthalmic Anatomic Reference Point Y-Coordinate",
		"OphthalmicAnatomicReferencePointYCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1627),
		"Ophthalmic En Face Volume Descriptor Sequence",
		"OphthalmicEnFaceVolumeDescriptorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1628),
		"Ophthalmic En Face Image Quality Rating Sequence",
		"OphthalmicEnFaceImageQualityRatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1629),
		"Ophthalmic En Face Volume Descriptor Scope",
		"OphthalmicEnFaceVolumeDescriptorScope",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1630),
		"Quality Threshold",
		"QualityThreshold",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1632),
		"Ophthalmic Anatomic Reference Point Sequence",
		"OphthalmicAnatomicReferencePointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1633),
		"Ophthalmic Anatomic Reference Point Localization Type",
		"OphthalmicAnatomicReferencePointLocalizationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1634),
		"Primary Anatomic Structure Item Index",
		"PrimaryAnatomicStructureItemIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1640),
		"OCT B-scan Analysis Acquisition Parameters Sequence",
		"OCTBscanAnalysisAcquisitionParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1642),
		"Number of B-scans Per Frame",
		"NumberOfBscansPerFrame",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1643),
		"B-scan Slab Thickness",
		"BscanSlabThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1644),
		"Distance Between B-scan Slabs",
		"DistanceBetweenBscanSlabs",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1645),
		"B-scan Cycle Time",
		"BscanCycleTime",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1646),
		"B-scan Cycle Time Vector",
		"BscanCycleTimeVector",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1649),
		"A-scan Rate",
		"AscanRate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1650),
		"B-scan Rate",
		"BscanRate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1658),
		"Surface Mesh Z-Pixel Offset",
		"SurfaceMeshZPixelOffset",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0010),
		"Visual Field Horizontal Extent",
		"VisualFieldHorizontalExtent",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0011),
		"Visual Field Vertical Extent",
		"VisualFieldVerticalExtent",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0012),
		"Visual Field Shape",
		"VisualFieldShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0016),
		"Screening Test Mode Code Sequence",
		"ScreeningTestModeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0018),
		"Maximum Stimulus Luminance",
		"MaximumStimulusLuminance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0020),
		"Background Luminance",
		"BackgroundLuminance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0021),
		"Stimulus Color Code Sequence",
		"StimulusColorCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0024),
		"Background Illumination Color Code Sequence",
		"BackgroundIlluminationColorCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0025),
		"Stimulus Area",
		"StimulusArea",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0028),
		"Stimulus Presentation Time",
		"StimulusPresentationTime",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0032),
		"Fixation Sequence",
		"FixationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0033),
		"Fixation Monitoring Code Sequence",
		"FixationMonitoringCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0034),
		"Visual Field Catch Trial Sequence",
		"VisualFieldCatchTrialSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0035),
		"Fixation Checked Quantity",
		"FixationCheckedQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0036),
		"Patient Not Properly Fixated Quantity",
		"PatientNotProperlyFixatedQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0037),
		"Presented Visual Stimuli Data Flag",
		"PresentedVisualStimuliDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0038),
		"Number of Visual Stimuli",
		"NumberOfVisualStimuli",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0039),
		"Excessive Fixation Losses Data Flag",
		"ExcessiveFixationLossesDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0040),
		"Excessive Fixation Losses",
		"ExcessiveFixationLosses",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0042),
		"Stimuli Retesting Quantity",
		"StimuliRetestingQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0044),
		"Comments on Patient's Performance of Visual Field",
		"CommentsOnPatientPerformanceOfVisualField",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0045),
		"False Negatives Estimate Flag",
		"FalseNegativesEstimateFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0046),
		"False Negatives Estimate",
		"FalseNegativesEstimate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0048),
		"Negative Catch Trials Quantity",
		"NegativeCatchTrialsQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0050),
		"False Negatives Quantity",
		"FalseNegativesQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0051),
		"Excessive False Negatives Data Flag",
		"ExcessiveFalseNegativesDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0052),
		"Excessive False Negatives",
		"ExcessiveFalseNegatives",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0053),
		"False Positives Estimate Flag",
		"FalsePositivesEstimateFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0054),
		"False Positives Estimate",
		"FalsePositivesEstimate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0055),
		"Catch Trials Data Flag",
		"CatchTrialsDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0056),
		"Positive Catch Trials Quantity",
		"PositiveCatchTrialsQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0057),
		"Test Point Normals Data Flag",
		"TestPointNormalsDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0058),
		"Test Point Normals Sequence",
		"TestPointNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0059),
		"Global Deviation Probability Normals Flag",
		"GlobalDeviationProbabilityNormalsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0060),
		"False Positives Quantity",
		"FalsePositivesQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0061),
		"Excessive False Positives Data Flag",
		"ExcessiveFalsePositivesDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0062),
		"Excessive False Positives",
		"ExcessiveFalsePositives",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0063),
		"Visual Field Test Normals Flag",
		"VisualFieldTestNormalsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0064),
		"Results Normals Sequence",
		"ResultsNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0065),
		"Age Corrected Sensitivity Deviation Algorithm Sequence",
		"AgeCorrectedSensitivityDeviationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0066),
		"Global Deviation From Normal",
		"GlobalDeviationFromNormal",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0067),
		"Generalized Defect Sensitivity Deviation Algorithm Sequence",
		"GeneralizedDefectSensitivityDeviationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0068),
		"Localized Deviation From Normal",
		"LocalizedDeviationFromNormal",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0069),
		"Patient Reliability Indicator",
		"PatientReliabilityIndicator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0070),
		"Visual Field Mean Sensitivity",
		"VisualFieldMeanSensitivity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0071),
		"Global Deviation Probability",
		"GlobalDeviationProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0072),
		"Local Deviation Probability Normals Flag",
		"LocalDeviationProbabilityNormalsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0073),
		"Localized Deviation Probability",
		"LocalizedDeviationProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0074),
		"Short Term Fluctuation Calculated",
		"ShortTermFluctuationCalculated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0075),
		"Short Term Fluctuation",
		"ShortTermFluctuation",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0076),
		"Short Term Fluctuation Probability Calculated",
		"ShortTermFluctuationProbabilityCalculated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0077),
		"Short Term Fluctuation Probability",
		"ShortTermFluctuationProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0078),
		"Corrected Localized Deviation From Normal Calculated",
		"CorrectedLocalizedDeviationFromNormalCalculated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0079),
		"Corrected Localized Deviation From Normal",
		"CorrectedLocalizedDeviationFromNormal",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0080),
		"Corrected Localized Deviation From Normal Probability Calculated",
		"CorrectedLocalizedDeviationFromNormalProbabilityCalculated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0081),
		"Corrected Localized Deviation From Normal Probability",
		"CorrectedLocalizedDeviationFromNormalProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0083),
		"Global Deviation Probability Sequence",
		"GlobalDeviationProbabilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0085),
		"Localized Deviation Probability Sequence",
		"LocalizedDeviationProbabilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0086),
		"Foveal Sensitivity Measured",
		"FovealSensitivityMeasured",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0087),
		"Foveal Sensitivity",
		"FovealSensitivity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0088),
		"Visual Field Test Duration",
		"VisualFieldTestDuration",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0089),
		"Visual Field Test Point Sequence",
		"VisualFieldTestPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0090),
		"Visual Field Test Point X-Coordinate",
		"VisualFieldTestPointXCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0091),
		"Visual Field Test Point Y-Coordinate",
		"VisualFieldTestPointYCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0092),
		"Age Corrected Sensitivity Deviation Value",
		"AgeCorrectedSensitivityDeviationValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0093),
		"Stimulus Results",
		"StimulusResults",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0094),
		"Sensitivity Value",
		"SensitivityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0095),
		"Retest Stimulus Seen",
		"RetestStimulusSeen",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0096),
		"Retest Sensitivity Value",
		"RetestSensitivityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0097),
		"Visual Field Test Point Normals Sequence",
		"VisualFieldTestPointNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0098),
		"Quantified Defect",
		"QuantifiedDefect",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0100),
		"Age Corrected Sensitivity Deviation Probability Value",
		"AgeCorrectedSensitivityDeviationProbabilityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0102),
		"Generalized Defect Corrected Sensitivity Deviation Flag",
		"GeneralizedDefectCorrectedSensitivityDeviationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0103),
		"Generalized Defect Corrected Sensitivity Deviation Value",
		"GeneralizedDefectCorrectedSensitivityDeviationValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0104),
		"Generalized Defect Corrected Sensitivity Deviation Probability Value",
		"GeneralizedDefectCorrectedSensitivityDeviationProbabilityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0105),
		"Minimum Sensitivity Value",
		"MinimumSensitivityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0106),
		"Blind Spot Localized",
		"BlindSpotLocalized",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0107),
		"Blind Spot X-Coordinate",
		"BlindSpotXCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0108),
		"Blind Spot Y-Coordinate",
		"BlindSpotYCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0110),
		"Visual Acuity Measurement Sequence",
		"VisualAcuityMeasurementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0112),
		"Refractive Parameters Used on Patient Sequence",
		"RefractiveParametersUsedOnPatientSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0113),
		"Measurement Laterality",
		"MeasurementLaterality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0114),
		"Ophthalmic Patient Clinical Information Left Eye Sequence",
		"OphthalmicPatientClinicalInformationLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0115),
		"Ophthalmic Patient Clinical Information Right Eye Sequence",
		"OphthalmicPatientClinicalInformationRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0117),
		"Foveal Point Normative Data Flag",
		"FovealPointNormativeDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0118),
		"Foveal Point Probability Value",
		"FovealPointProbabilityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0120),
		"Screening Baseline Measured",
		"ScreeningBaselineMeasured",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0122),
		"Screening Baseline Measured Sequence",
		"ScreeningBaselineMeasuredSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0124),
		"Screening Baseline Type",
		"ScreeningBaselineType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0126),
		"Screening Baseline Value",
		"ScreeningBaselineValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0202),
		"Algorithm Source",
		"AlgorithmSource",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0306),
		"Data Set Name",
		"DataSetName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0307),
		"Data Set Version",
		"DataSetVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0308),
		"Data Set Source",
		"DataSetSource",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0309),
		"Data Set Description",
		"DataSetDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0317),
		"Visual Field Test Reliability Global Index Sequence",
		"VisualFieldTestReliabilityGlobalIndexSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0320),
		"Visual Field Global Results Index Sequence",
		"VisualFieldGlobalResultsIndexSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0325),
		"Data Observation Sequence",
		"DataObservationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0338),
		"Index Normals Flag",
		"IndexNormalsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0341),
		"Index Probability",
		"IndexProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0344),
		"Index Probability Sequence",
		"IndexProbabilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0002),
		"Samples per Pixel",
		"SamplesPerPixel",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0003),
		"Samples per Pixel Used",
		"SamplesPerPixelUsed",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0004),
		"Photometric Interpretation",
		"PhotometricInterpretation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0005),
		"Image Dimensions",
		"ImageDimensions",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0006),
		"Planar Configuration",
		"PlanarConfiguration",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0008),
		"Number of Frames",
		"NumberOfFrames",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0009),
		"Frame Increment Pointer",
		"FrameIncrementPointer",
		vm.VM1N,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x000A),
		"Frame Dimension Pointer",
		"FrameDimensionPointer",
		vm.VM1N,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0010),
		"Rows",
		"Rows",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0011),
		"Columns",
		"Columns",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0012),
		"Planes",
		"Planes",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0014),
		"Ultrasound Color Data Present",
		"UltrasoundColorDataPresent",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0030),
		"Pixel Spacing",
		"PixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0031),
		"Zoom Factor",
		"ZoomFactor",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0032),
		"Zoom Center",
		"ZoomCenter",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0034),
		"Pixel Aspect Ratio",
		"PixelAspectRatio",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0040),
		"Image Format",
		"ImageFormat",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0050),
		"Manipulated Image",
		"ManipulatedImage",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0051),
		"Corrected Image",
		"CorrectedImage",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x005F),
		"Compression Recognition Code",
		"CompressionRecognitionCode",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0060),
		"Compression Code",
		"CompressionCode",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0061),
		"Compression Originator",
		"CompressionOriginator",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0062),
		"Compression Label",
		"CompressionLabel",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0063),
		"Compression Description",
		"CompressionDescription",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0065),
		"Compression Sequence",
		"CompressionSequence",
		vm.VM1N,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0066),
		"Compression Step Pointers",
		"CompressionStepPointers",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0068),
		"Repeat Interval",
		"RepeatInterval",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0069),
		"Bits Grouped",
		"BitsGrouped",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0070),
		"Perimeter Table",
		"PerimeterTable",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0071),
		"Perimeter Value",
		"PerimeterValue",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0080),
		"Predictor Rows",
		"PredictorRows",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0081),
		"Predictor Columns",
		"PredictorColumns",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0082),
		"Predictor Constants",
		"PredictorConstants",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0090),
		"Blocked Pixels",
		"BlockedPixels",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0091),
		"Block Rows",
		"BlockRows",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0092),
		"Block Columns",
		"BlockColumns",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0093),
		"Row Overlap",
		"RowOverlap",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0094),
		"Column Overlap",
		"ColumnOverlap",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0100),
		"Bits Allocated",
		"BitsAllocated",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0101),
		"Bits Stored",
		"BitsStored",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0102),
		"High Bit",
		"HighBit",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0103),
		"Pixel Representation",
		"PixelRepresentation",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0104),
		"Smallest Valid Pixel Value",
		"SmallestValidPixelValue",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0105),
		"Largest Valid Pixel Value",
		"LargestValidPixelValue",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0106),
		"Smallest Image Pixel Value",
		"SmallestImagePixelValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0107),
		"Largest Image Pixel Value",
		"LargestImagePixelValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0108),
		"Smallest Pixel Value in Series",
		"SmallestPixelValueInSeries",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0109),
		"Largest Pixel Value in Series",
		"LargestPixelValueInSeries",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0110),
		"Smallest Image Pixel Value in Plane",
		"SmallestImagePixelValueInPlane",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0111),
		"Largest Image Pixel Value in Plane",
		"LargestImagePixelValueInPlane",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0120),
		"Pixel Padding Value",
		"PixelPaddingValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0121),
		"Pixel Padding Range Limit",
		"PixelPaddingRangeLimit",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0122),
		"Float Pixel Padding Value",
		"FloatPixelPaddingValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0123),
		"Double Float Pixel Padding Value",
		"DoubleFloatPixelPaddingValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0124),
		"Float Pixel Padding Range Limit",
		"FloatPixelPaddingRangeLimit",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0125),
		"Double Float Pixel Padding Range Limit",
		"DoubleFloatPixelPaddingRangeLimit",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0200),
		"Image Location",
		"ImageLocation",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0300),
		"Quality Control Image",
		"QualityControlImage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0301),
		"Burned In Annotation",
		"BurnedInAnnotation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0302),
		"Recognizable Visual Features",
		"RecognizableVisualFeatures",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0303),
		"Longitudinal Temporal Information Modified",
		"LongitudinalTemporalInformationModified",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0304),
		"Referenced Color Palette Instance UID",
		"ReferencedColorPaletteInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0400),
		"Transform Label",
		"TransformLabel",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0401),
		"Transform Version Number",
		"TransformVersionNumber",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0402),
		"Number of Transform Steps",
		"NumberOfTransformSteps",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0403),
		"Sequence of Compressed Data",
		"SequenceOfCompressedData",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0404),
		"Details of Coefficients",
		"DetailsOfCoefficients",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,04x0)"),
		"Rows For Nth Order Coefficients",
		"RowsForNthOrderCoefficients",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,04x1)"),
		"Columns For Nth Order Coefficients",
		"ColumnsForNthOrderCoefficients",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,04x2)"),
		"Coefficient Coding",
		"CoefficientCoding",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,04x3)"),
		"Coefficient Coding Pointers",
		"CoefficientCodingPointers",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0700),
		"DCT Label",
		"DCTLabel",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0701),
		"Data Block Description",
		"DataBlockDescription",
		vm.VM1N,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0702),
		"Data Block",
		"DataBlock",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0710),
		"Normalization Factor Format",
		"NormalizationFactorFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0720),
		"Zonal Map Number Format",
		"ZonalMapNumberFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0721),
		"Zonal Map Location",
		"ZonalMapLocation",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0722),
		"Zonal Map Format",
		"ZonalMapFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0730),
		"Adaptive Map Format",
		"AdaptiveMapFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0740),
		"Code Number Format",
		"CodeNumberFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x0)"),
		"Code Label",
		"CodeLabel",
		vm.VM1N,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x2)"),
		"Number of Tables",
		"NumberOfTables",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x3)"),
		"Code Table Location",
		"CodeTableLocation",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x4)"),
		"Bits For Code Word",
		"BitsForCodeWord",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x8)"),
		"Image Data Location",
		"ImageDataLocation",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0A02),
		"Pixel Spacing Calibration Type",
		"PixelSpacingCalibrationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0A04),
		"Pixel Spacing Calibration Description",
		"PixelSpacingCalibrationDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1040),
		"Pixel Intensity Relationship",
		"PixelIntensityRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1041),
		"Pixel Intensity Relationship Sign",
		"PixelIntensityRelationshipSign",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1050),
		"Window Center",
		"WindowCenter",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1051),
		"Window Width",
		"WindowWidth",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1052),
		"Rescale Intercept",
		"RescaleIntercept",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1053),
		"Rescale Slope",
		"RescaleSlope",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1054),
		"Rescale Type",
		"RescaleType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1055),
		"Window Center & Width Explanation",
		"WindowCenterWidthExplanation",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1056),
		"VOI LUT Function",
		"VOILUTFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1080),
		"Gray Scale",
		"GrayScale",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1090),
		"Recommended Viewing Mode",
		"RecommendedViewingMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1100),
		"Gray Lookup Table Descriptor",
		"GrayLookupTableDescriptor",
		vm.VM3,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1101),
		"Red Palette Color Lookup Table Descriptor",
		"RedPaletteColorLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1102),
		"Green Palette Color Lookup Table Descriptor",
		"GreenPaletteColorLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1103),
		"Blue Palette Color Lookup Table Descriptor",
		"BluePaletteColorLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1104),
		"Alpha Palette Color Lookup Table Descriptor",
		"AlphaPaletteColorLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1111),
		"Large Red Palette Color Lookup Table Descriptor",
		"LargeRedPaletteColorLookupTableDescriptor",
		vm.VM4,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1112),
		"Large Green Palette Color Lookup Table Descriptor",
		"LargeGreenPaletteColorLookupTableDescriptor",
		vm.VM4,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1113),
		"Large Blue Palette Color Lookup Table Descriptor",
		"LargeBluePaletteColorLookupTableDescriptor",
		vm.VM4,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1199),
		"Palette Color Lookup Table UID",
		"PaletteColorLookupTableUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1200),
		"Gray Lookup Table Data",
		"GrayLookupTableData",
		vm.VM1N,
		true,
		vr.US, vr.SS, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1201),
		"Red Palette Color Lookup Table Data",
		"RedPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1202),
		"Green Palette Color Lookup Table Data",
		"GreenPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1203),
		"Blue Palette Color Lookup Table Data",
		"BluePaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1204),
		"Alpha Palette Color Lookup Table Data",
		"AlphaPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1211),
		"Large Red Palette Color Lookup Table Data",
		"LargeRedPaletteColorLookupTableData",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1212),
		"Large Green Palette Color Lookup Table Data",
		"LargeGreenPaletteColorLookupTableData",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1213),
		"Large Blue Palette Color Lookup Table Data",
		"LargeBluePaletteColorLookupTableData",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1214),
		"Large Palette Color Lookup Table UID",
		"LargePaletteColorLookupTableUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1221),
		"Segmented Red Palette Color Lookup Table Data",
		"SegmentedRedPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1222),
		"Segmented Green Palette Color Lookup Table Data",
		"SegmentedGreenPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1223),
		"Segmented Blue Palette Color Lookup Table Data",
		"SegmentedBluePaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1224),
		"Segmented Alpha Palette Color Lookup Table Data",
		"SegmentedAlphaPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1230),
		"Stored Value Color Range Sequence",
		"StoredValueColorRangeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1231),
		"Minimum Stored Value Mapped",
		"MinimumStoredValueMapped",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1232),
		"Maximum Stored Value Mapped",
		"MaximumStoredValueMapped",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1300),
		"Breast Implant Present",
		"BreastImplantPresent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1350),
		"Partial View",
		"PartialView",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1351),
		"Partial View Description",
		"PartialViewDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1352),
		"Partial View Code Sequence",
		"PartialViewCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x135A),
		"Spatial Locations Preserved",
		"SpatialLocationsPreserved",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1401),
		"Data Frame Assignment Sequence",
		"DataFrameAssignmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1402),
		"Data Path Assignment",
		"DataPathAssignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1403),
		"Bits Mapped to Color Lookup Table",
		"BitsMappedToColorLookupTable",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1404),
		"Blending LUT 1 Sequence",
		"BlendingLUT1Sequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1405),
		"Blending LUT 1 Transfer Function",
		"BlendingLUT1TransferFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1406),
		"Blending Weight Constant",
		"BlendingWeightConstant",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1407),
		"Blending Lookup Table Descriptor",
		"BlendingLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1408),
		"Blending Lookup Table Data",
		"BlendingLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140B),
		"Enhanced Palette Color Lookup Table Sequence",
		"EnhancedPaletteColorLookupTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140C),
		"Blending LUT 2 Sequence",
		"BlendingLUT2Sequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140D),
		"Blending LUT 2 Transfer Function",
		"BlendingLUT2TransferFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140E),
		"Data Path ID",
		"DataPathID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140F),
		"RGB LUT Transfer Function",
		"RGBLUTTransferFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1410),
		"Alpha LUT Transfer Function",
		"AlphaLUTTransferFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2000),
		"ICC Profile",
		"ICCProfile",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2002),
		"Color Space",
		"ColorSpace",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2110),
		"Lossy Image Compression",
		"LossyImageCompression",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2112),
		"Lossy Image Compression Ratio",
		"LossyImageCompressionRatio",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2114),
		"Lossy Image Compression Method",
		"LossyImageCompressionMethod",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3000),
		"Modality LUT Sequence",
		"ModalityLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3001),
		"Variable Modality LUT Sequence",
		"VariableModalityLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3002),
		"LUT Descriptor",
		"LUTDescriptor",
		vm.VM3,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3003),
		"LUT Explanation",
		"LUTExplanation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3004),
		"Modality LUT Type",
		"ModalityLUTType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3006),
		"LUT Data",
		"LUTData",
		vm.VM1N,
		false,
		vr.US, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3010),
		"VOI LUT Sequence",
		"VOILUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3110),
		"Softcopy VOI LUT Sequence",
		"SoftcopyVOILUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x4000),
		"Image Presentation Comments",
		"ImagePresentationComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x5000),
		"Bi-Plane Acquisition Sequence",
		"BiPlaneAcquisitionSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6010),
		"Representative Frame Number",
		"RepresentativeFrameNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6020),
		"Frame Numbers of Interest (FOI)",
		"FrameNumbersOfInterest",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6022),
		"Frame of Interest Description",
		"FrameOfInterestDescription",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6023),
		"Frame of Interest Type",
		"FrameOfInterestType",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6030),
		"Mask Pointer(s)",
		"MaskPointers",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6040),
		"R Wave Pointer",
		"RWavePointer",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6100),
		"Mask Subtraction Sequence",
		"MaskSubtractionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6101),
		"Mask Operation",
		"MaskOperation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6102),
		"Applicable Frame Range",
		"ApplicableFrameRange",
		vm.VM22N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6110),
		"Mask Frame Numbers",
		"MaskFrameNumbers",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6112),
		"Contrast Frame Averaging",
		"ContrastFrameAveraging",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6114),
		"Mask Sub-pixel Shift",
		"MaskSubPixelShift",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6120),
		"TID Offset",
		"TIDOffset",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6190),
		"Mask Operation Explanation",
		"MaskOperationExplanation",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7000),
		"Equipment Administrator Sequence",
		"EquipmentAdministratorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7001),
		"Number of Display Subsystems",
		"NumberOfDisplaySubsystems",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7002),
		"Current Configuration ID",
		"CurrentConfigurationID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7003),
		"Display Subsystem ID",
		"DisplaySubsystemID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7004),
		"Display Subsystem Name",
		"DisplaySubsystemName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7005),
		"Display Subsystem Description",
		"DisplaySubsystemDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7006),
		"System Status",
		"SystemStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7007),
		"System Status Comment",
		"SystemStatusComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7008),
		"Target Luminance Characteristics Sequence",
		"TargetLuminanceCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7009),
		"Luminance Characteristics ID",
		"LuminanceCharacteristicsID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700A),
		"Display Subsystem Configuration Sequence",
		"DisplaySubsystemConfigurationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700B),
		"Configuration ID",
		"ConfigurationID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700C),
		"Configuration Name",
		"ConfigurationName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700D),
		"Configuration Description",
		"ConfigurationDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700E),
		"Referenced Target Luminance Characteristics ID",
		"ReferencedTargetLuminanceCharacteristicsID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700F),
		"QA Results Sequence",
		"QAResultsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7010),
		"Display Subsystem QA Results Sequence",
		"DisplaySubsystemQAResultsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7011),
		"Configuration QA Results Sequence",
		"ConfigurationQAResultsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7012),
		"Measurement Equipment Sequence",
		"MeasurementEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7013),
		"Measurement Functions",
		"MeasurementFunctions",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7014),
		"Measurement Equipment Type",
		"MeasurementEquipmentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7015),
		"Visual Evaluation Result Sequence",
		"VisualEvaluationResultSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7016),
		"Display Calibration Result Sequence",
		"DisplayCalibrationResultSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7017),
		"DDL Value",
		"DDLValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7018),
		"CIExy White Point",
		"CIExyWhitePoint",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7019),
		"Display Function Type",
		"DisplayFunctionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701A),
		"Gamma Value",
		"GammaValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701B),
		"Number of Luminance Points",
		"NumberOfLuminancePoints",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701C),
		"Luminance Response Sequence",
		"LuminanceResponseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701D),
		"Target Minimum Luminance",
		"TargetMinimumLuminance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701E),
		"Target Maximum Luminance",
		"TargetMaximumLuminance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701F),
		"Luminance Value",
		"LuminanceValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7020),
		"Luminance Response Description",
		"LuminanceResponseDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7021),
		"White Point Flag",
		"WhitePointFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7022),
		"Display Device Type Code Sequence",
		"DisplayDeviceTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7023),
		"Display Subsystem Sequence",
		"DisplaySubsystemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7024),
		"Luminance Result Sequence",
		"LuminanceResultSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7025),
		"Ambient Light Value Source",
		"AmbientLightValueSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7026),
		"Measured Characteristics",
		"MeasuredCharacteristics",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7027),
		"Luminance Uniformity Result Sequence",
		"LuminanceUniformityResultSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7028),
		"Visual Evaluation Test Sequence",
		"VisualEvaluationTestSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7029),
		"Test Result",
		"TestResult",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702A),
		"Test Result Comment",
		"TestResultComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702B),
		"Test Image Validation",
		"TestImageValidation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702C),
		"Test Pattern Code Sequence",
		"TestPatternCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702D),
		"Measurement Pattern Code Sequence",
		"MeasurementPatternCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702E),
		"Visual Evaluation Method Code Sequence",
		"VisualEvaluationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7FE0),
		"Pixel Data Provider URL",
		"PixelDataProviderURL",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9001),
		"Data Point Rows",
		"DataPointRows",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9002),
		"Data Point Columns",
		"DataPointColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9003),
		"Signal Domain Columns",
		"SignalDomainColumns",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9099),
		"Largest Monochrome Pixel Value",
		"LargestMonochromePixelValue",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9108),
		"Data Representation",
		"DataRepresentation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9110),
		"Pixel Measures Sequence",
		"PixelMeasuresSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9132),
		"Frame VOI LUT Sequence",
		"FrameVOILUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9145),
		"Pixel Value Transformation Sequence",
		"PixelValueTransformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9235),
		"Signal Domain Rows",
		"SignalDomainRows",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9411),
		"Display Filter Percentage",
		"DisplayFilterPercentage",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9415),
		"Frame Pixel Shift Sequence",
		"FramePixelShiftSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9416),
		"Subtraction Item ID",
		"SubtractionItemID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9422),
		"Pixel Intensity Relationship LUT Sequence",
		"PixelIntensityRelationshipLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9443),
		"Frame Pixel Data Properties Sequence",
		"FramePixelDataPropertiesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9444),
		"Geometrical Properties",
		"GeometricalProperties",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9445),
		"Geometric Maximum Distortion",
		"GeometricMaximumDistortion",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9446),
		"Image Processing Applied",
		"ImageProcessingApplied",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9454),
		"Mask Selection Mode",
		"MaskSelectionMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9474),
		"LUT Function",
		"LUTFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9478),
		"Mask Visibility Percentage",
		"MaskVisibilityPercentage",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9501),
		"Pixel Shift Sequence",
		"PixelShiftSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9502),
		"Region Pixel Shift Sequence",
		"RegionPixelShiftSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9503),
		"Vertices of the Region",
		"VerticesOfTheRegion",
		vm.VM22N,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9505),
		"Multi-frame Presentation Sequence",
		"MultiFramePresentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9506),
		"Pixel Shift Frame Range",
		"PixelShiftFrameRange",
		vm.VM22N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9507),
		"LUT Frame Range",
		"LUTFrameRange",
		vm.VM22N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9520),
		"Image to Equipment Mapping Matrix",
		"ImageToEquipmentMappingMatrix",
		vm.VM16,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9537),
		"Equipment Coordinate System Identification",
		"EquipmentCoordinateSystemIdentification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x000A),
		"Study Status ID",
		"StudyStatusID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x000C),
		"Study Priority ID",
		"StudyPriorityID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0012),
		"Study ID Issuer",
		"StudyIDIssuer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0032),
		"Study Verified Date",
		"StudyVerifiedDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0033),
		"Study Verified Time",
		"StudyVerifiedTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0034),
		"Study Read Date",
		"StudyReadDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0035),
		"Study Read Time",
		"StudyReadTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1000),
		"Scheduled Study Start Date",
		"ScheduledStudyStartDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1001),
		"Scheduled Study Start Time",
		"ScheduledStudyStartTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1010),
		"Scheduled Study Stop Date",
		"ScheduledStudyStopDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1011),
		"Scheduled Study Stop Time",
		"ScheduledStudyStopTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1020),
		"Scheduled Study Location",
		"ScheduledStudyLocation",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1021),
		"Scheduled Study Location AE Title",
		"ScheduledStudyLocationAETitle",
		vm.VM1N,
		true,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1030),
		"Reason for Study",
		"ReasonForStudy",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1031),
		"Requesting Physician Identification Sequence",
		"RequestingPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1032),
		"Requesting Physician",
		"RequestingPhysician",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1033),
		"Requesting Service",
		"RequestingService",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1034),
		"Requesting Service Code Sequence",
		"RequestingServiceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1040),
		"Study Arrival Date",
		"StudyArrivalDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1041),
		"Study Arrival Time",
		"StudyArrivalTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1050),
		"Study Completion Date",
		"StudyCompletionDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1051),
		"Study Completion Time",
		"StudyCompletionTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1055),
		"Study Component Status ID",
		"StudyComponentStatusID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1060),
		"Requested Procedure Description",
		"RequestedProcedureDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1064),
		"Requested Procedure Code Sequence",
		"RequestedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1065),
		"Requested Laterality Code Sequence",
		"RequestedLateralityCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1066),
		"Reason for Visit",
		"ReasonForVisit",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1067),
		"Reason for Visit Code Sequence",
		"ReasonForVisitCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1070),
		"Requested Contrast Agent",
		"RequestedContrastAgent",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x4000),
		"Study Comments",
		"StudyComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0001),
		"Flow Identifier Sequence",
		"FlowIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0002),
		"Flow Identifier",
		"FlowIdentifier",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0003),
		"Flow Transfer Syntax UID",
		"FlowTransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0004),
		"Flow RTP Sampling Rate",
		"FlowRTPSamplingRate",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0005),
		"Source Identifier",
		"SourceIdentifier",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0007),
		"Frame Origin Timestamp",
		"FrameOriginTimestamp",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0008),
		"Includes Imaging Subject",
		"IncludesImagingSubject",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0009),
		"Frame Usefulness Group Sequence",
		"FrameUsefulnessGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x000A),
		"Real-Time Bulk Data Flow Sequence",
		"RealTimeBulkDataFlowSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x000B),
		"Camera Position Group Sequence",
		"CameraPositionGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x000C),
		"Includes Information",
		"IncludesInformation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x000D),
		"Time of Frame Group Sequence",
		"TimeOfFrameGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0004),
		"Referenced Patient Alias Sequence",
		"ReferencedPatientAliasSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0008),
		"Visit Status ID",
		"VisitStatusID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0010),
		"Admission ID",
		"AdmissionID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0011),
		"Issuer of Admission ID",
		"IssuerOfAdmissionID",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0014),
		"Issuer of Admission ID Sequence",
		"IssuerOfAdmissionIDSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0016),
		"Route of Admissions",
		"RouteOfAdmissions",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001A),
		"Scheduled Admission Date",
		"ScheduledAdmissionDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001B),
		"Scheduled Admission Time",
		"ScheduledAdmissionTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001C),
		"Scheduled Discharge Date",
		"ScheduledDischargeDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001D),
		"Scheduled Discharge Time",
		"ScheduledDischargeTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001E),
		"Scheduled Patient Institution Residence",
		"ScheduledPatientInstitutionResidence",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0020),
		"Admitting Date",
		"AdmittingDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0021),
		"Admitting Time",
		"AdmittingTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0030),
		"Discharge Date",
		"DischargeDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0032),
		"Discharge Time",
		"DischargeTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0040),
		"Discharge Diagnosis Description",
		"DischargeDiagnosisDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0044),
		"Discharge Diagnosis Code Sequence",
		"DischargeDiagnosisCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0050),
		"Special Needs",
		"SpecialNeeds",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0060),
		"Service Episode ID",
		"ServiceEpisodeID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0061),
		"Issuer of Service Episode ID",
		"IssuerOfServiceEpisodeID",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0062),
		"Service Episode Description",
		"ServiceEpisodeDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0064),
		"Issuer of Service Episode ID Sequence",
		"IssuerOfServiceEpisodeIDSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0100),
		"Pertinent Documents Sequence",
		"PertinentDocumentsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0101),
		"Pertinent Resources Sequence",
		"PertinentResourcesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0102),
		"Resource Description",
		"ResourceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0300),
		"Current Patient Location",
		"CurrentPatientLocation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0400),
		"Patient's Institution Residence",
		"PatientInstitutionResidence",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0500),
		"Patient State",
		"PatientState",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0502),
		"Patient Clinical Trial Participation Sequence",
		"PatientClinicalTrialParticipationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x4000),
		"Visit Comments",
		"VisitComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0004),
		"Waveform Originality",
		"WaveformOriginality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0005),
		"Number of Waveform Channels",
		"NumberOfWaveformChannels",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0010),
		"Number of Waveform Samples",
		"NumberOfWaveformSamples",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x001A),
		"Sampling Frequency",
		"SamplingFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0020),
		"Multiplex Group Label",
		"MultiplexGroupLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0200),
		"Channel Definition Sequence",
		"ChannelDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0202),
		"Waveform Channel Number",
		"WaveformChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0203),
		"Channel Label",
		"ChannelLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0205),
		"Channel Status",
		"ChannelStatus",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0208),
		"Channel Source Sequence",
		"ChannelSourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0209),
		"Channel Source Modifiers Sequence",
		"ChannelSourceModifiersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x020A),
		"Source Waveform Sequence",
		"SourceWaveformSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x020C),
		"Channel Derivation Description",
		"ChannelDerivationDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0210),
		"Channel Sensitivity",
		"ChannelSensitivity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0211),
		"Channel Sensitivity Units Sequence",
		"ChannelSensitivityUnitsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0212),
		"Channel Sensitivity Correction Factor",
		"ChannelSensitivityCorrectionFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0213),
		"Channel Baseline",
		"ChannelBaseline",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0214),
		"Channel Time Skew",
		"ChannelTimeSkew",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0215),
		"Channel Sample Skew",
		"ChannelSampleSkew",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0218),
		"Channel Offset",
		"ChannelOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x021A),
		"Waveform Bits Stored",
		"WaveformBitsStored",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0220),
		"Filter Low Frequency",
		"FilterLowFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0221),
		"Filter High Frequency",
		"FilterHighFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0222),
		"Notch Filter Frequency",
		"NotchFilterFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0223),
		"Notch Filter Bandwidth",
		"NotchFilterBandwidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0230),
		"Waveform Data Display Scale",
		"WaveformDataDisplayScale",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0231),
		"Waveform Display Background CIELab Value",
		"WaveformDisplayBackgroundCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0240),
		"Waveform Presentation Group Sequence",
		"WaveformPresentationGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0241),
		"Presentation Group Number",
		"PresentationGroupNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0242),
		"Channel Display Sequence",
		"ChannelDisplaySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0244),
		"Channel Recommended Display CIELab Value",
		"ChannelRecommendedDisplayCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0245),
		"Channel Position",
		"ChannelPosition",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0246),
		"Display Shading Flag",
		"DisplayShadingFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0247),
		"Fractional Channel Display Scale",
		"FractionalChannelDisplayScale",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0248),
		"Absolute Channel Display Scale",
		"AbsoluteChannelDisplayScale",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0300),
		"Multiplexed Audio Channels Description Code Sequence",
		"MultiplexedAudioChannelsDescriptionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0301),
		"Channel Identification Code",
		"ChannelIdentificationCode",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0302),
		"Channel Mode",
		"ChannelMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0310),
		"Multiplex Group UID",
		"MultiplexGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0311),
		"Powerline Frequency",
		"PowerlineFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0312),
		"Channel Impedance Sequence",
		"ChannelImpedanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0313),
		"Impedance Value",
		"ImpedanceValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0314),
		"Impedance Measurement DateTime",
		"ImpedanceMeasurementDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0315),
		"Impedance Measurement Frequency",
		"ImpedanceMeasurementFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0316),
		"Impedance Measurement Current Type",
		"ImpedanceMeasurementCurrentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0317),
		"Waveform Amplifier Type",
		"WaveformAmplifierType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0318),
		"Filter Low Frequency Characteristics Sequence",
		"FilterLowFrequencyCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0319),
		"Filter High Frequency Characteristics Sequence",
		"FilterHighFrequencyCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0320),
		"Summarized Filter Lookup Table Sequence",
		"SummarizedFilterLookupTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0321),
		"Notch Filter Characteristics Sequence",
		"NotchFilterCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0322),
		"Waveform Filter Type",
		"WaveformFilterType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0323),
		"Analog Filter Characteristics Sequence",
		"AnalogFilterCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0324),
		"Analog Filter Roll Off",
		"AnalogFilterRollOff",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0325),
		"Analog Filter Type Code Sequence",
		"AnalogFilterTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0326),
		"Digital Filter Characteristics Sequence",
		"DigitalFilterCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0327),
		"Digital Filter Order",
		"DigitalFilterOrder",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0328),
		"Digital Filter Type Code Sequence",
		"DigitalFilterTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0329),
		"Waveform Filter Description",
		"WaveformFilterDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032A),
		"Filter Lookup Table Sequence",
		"FilterLookupTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032B),
		"Filter Lookup Table Description",
		"FilterLookupTableDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032C),
		"Frequency Encoding Code Sequence",
		"FrequencyEncodingCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032D),
		"Magnitude Encoding Code Sequence",
		"MagnitudeEncodingCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032E),
		"Filter Lookup Table Data",
		"FilterLookupTableData",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0001),
		"Scheduled Station AE Title",
		"ScheduledStationAETitle",
		vm.VM1N,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0002),
		"Scheduled Procedure Step Start Date",
		"ScheduledProcedureStepStartDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0003),
		"Scheduled Procedure Step Start Time",
		"ScheduledProcedureStepStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0004),
		"Scheduled Procedure Step End Date",
		"ScheduledProcedureStepEndDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0005),
		"Scheduled Procedure Step End Time",
		"ScheduledProcedureStepEndTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0006),
		"Scheduled Performing Physician's Name",
		"ScheduledPerformingPhysicianName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0007),
		"Scheduled Procedure Step Description",
		"ScheduledProcedureStepDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0008),
		"Scheduled Protocol Code Sequence",
		"ScheduledProtocolCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0009),
		"Scheduled Procedure Step ID",
		"ScheduledProcedureStepID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x000A),
		"Stage Code Sequence",
		"StageCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x000B),
		"Scheduled Performing Physician Identification Sequence",
		"ScheduledPerformingPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0010),
		"Scheduled Station Name",
		"ScheduledStationName",
		vm.VM1N,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0011),
		"Scheduled Procedure Step Location",
		"ScheduledProcedureStepLocation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0012),
		"Pre-Medication",
		"PreMedication",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0020),
		"Scheduled Procedure Step Status",
		"ScheduledProcedureStepStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0026),
		"Order Placer Identifier Sequence",
		"OrderPlacerIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0027),
		"Order Filler Identifier Sequence",
		"OrderFillerIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0031),
		"Local Namespace Entity ID",
		"LocalNamespaceEntityID",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0032),
		"Universal Entity ID",
		"UniversalEntityID",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0033),
		"Universal Entity ID Type",
		"UniversalEntityIDType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0035),
		"Identifier Type Code",
		"IdentifierTypeCode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0036),
		"Assigning Facility Sequence",
		"AssigningFacilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0039),
		"Assigning Jurisdiction Code Sequence",
		"AssigningJurisdictionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x003A),
		"Assigning Agency or Department Code Sequence",
		"AssigningAgencyOrDepartmentCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0100),
		"Scheduled Procedure Step Sequence",
		"ScheduledProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0220),
		"Referenced Non-Image Composite SOP Instance Sequence",
		"ReferencedNonImageCompositeSOPInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0241),
		"Performed Station AE Title",
		"PerformedStationAETitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0242),
		"Performed Station Name",
		"PerformedStationName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0243),
		"Performed Location",
		"PerformedLocation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0244),
		"Performed Procedure Step Start Date",
		"PerformedProcedureStepStartDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0245),
		"Performed Procedure Step Start Time",
		"PerformedProcedureStepStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0250),
		"Performed Procedure Step End Date",
		"PerformedProcedureStepEndDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0251),
		"Performed Procedure Step End Time",
		"PerformedProcedureStepEndTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0252),
		"Performed Procedure Step Status",
		"PerformedProcedureStepStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0253),
		"Performed Procedure Step ID",
		"PerformedProcedureStepID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0254),
		"Performed Procedure Step Description",
		"PerformedProcedureStepDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0255),
		"Performed Procedure Type Description",
		"PerformedProcedureTypeDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0260),
		"Performed Protocol Code Sequence",
		"PerformedProtocolCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0261),
		"Performed Protocol Type",
		"PerformedProtocolType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0270),
		"Scheduled Step Attributes Sequence",
		"ScheduledStepAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0275),
		"Request Attributes Sequence",
		"RequestAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0280),
		"Comments on the Performed Procedure Step",
		"CommentsOnThePerformedProcedureStep",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0281),
		"Performed Procedure Step Discontinuation Reason Code Sequence",
		"PerformedProcedureStepDiscontinuationReasonCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0293),
		"Quantity Sequence",
		"QuantitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0294),
		"Quantity",
		"Quantity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0295),
		"Measuring Units Sequence",
		"MeasuringUnitsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0296),
		"Billing Item Sequence",
		"BillingItemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0300),
		"Total Time of Fluoroscopy",
		"TotalTimeOfFluoroscopy",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0301),
		"Total Number of Exposures",
		"TotalNumberOfExposures",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0302),
		"Entrance Dose",
		"EntranceDose",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0303),
		"Exposed Area",
		"ExposedArea",
		vm.VM12,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0306),
		"Distance Source to Entrance",
		"DistanceSourceToEntrance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0307),
		"Distance Source to Support",
		"DistanceSourceToSupport",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x030E),
		"Exposure Dose Sequence",
		"ExposureDoseSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0310),
		"Comments on Radiation Dose",
		"CommentsOnRadiationDose",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0312),
		"X-Ray Output",
		"XRayOutput",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0314),
		"Half Value Layer",
		"HalfValueLayer",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0316),
		"Organ Dose",
		"OrganDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0318),
		"Organ Exposed",
		"OrganExposed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0320),
		"Billing Procedure Step Sequence",
		"BillingProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0321),
		"Film Consumption Sequence",
		"FilmConsumptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0324),
		"Billing Supplies and Devices Sequence",
		"BillingSuppliesAndDevicesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0330),
		"Referenced Procedure Step Sequence",
		"ReferencedProcedureStepSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0340),
		"Performed Series Sequence",
		"PerformedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0400),
		"Comments on the Scheduled Procedure Step",
		"CommentsOnTheScheduledProcedureStep",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0440),
		"Protocol Context Sequence",
		"ProtocolContextSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0441),
		"Content Item Modifier Sequence",
		"ContentItemModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0500),
		"Scheduled Specimen Sequence",
		"ScheduledSpecimenSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x050A),
		"Specimen Accession Number",
		"SpecimenAccessionNumber",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0512),
		"Container Identifier",
		"ContainerIdentifier",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0513),
		"Issuer of the Container Identifier Sequence",
		"IssuerOfTheContainerIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0515),
		"Alternate Container Identifier Sequence",
		"AlternateContainerIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0518),
		"Container Type Code Sequence",
		"ContainerTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x051A),
		"Container Description",
		"ContainerDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0520),
		"Container Component Sequence",
		"ContainerComponentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0550),
		"Specimen Sequence",
		"SpecimenSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0551),
		"Specimen Identifier",
		"SpecimenIdentifier",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0552),
		"Specimen Description Sequence (Trial)",
		"SpecimenDescriptionSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0553),
		"Specimen Description (Trial)",
		"SpecimenDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0554),
		"Specimen UID",
		"SpecimenUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0555),
		"Acquisition Context Sequence",
		"AcquisitionContextSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0556),
		"Acquisition Context Description",
		"AcquisitionContextDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x059A),
		"Specimen Type Code Sequence",
		"SpecimenTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0560),
		"Specimen Description Sequence",
		"SpecimenDescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0562),
		"Issuer of the Specimen Identifier Sequence",
		"IssuerOfTheSpecimenIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0600),
		"Specimen Short Description",
		"SpecimenShortDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0602),
		"Specimen Detailed Description",
		"SpecimenDetailedDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0610),
		"Specimen Preparation Sequence",
		"SpecimenPreparationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0612),
		"Specimen Preparation Step Content Item Sequence",
		"SpecimenPreparationStepContentItemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0620),
		"Specimen Localization Content Item Sequence",
		"SpecimenLocalizationContentItemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x06FA),
		"Slide Identifier",
		"SlideIdentifier",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0710),
		"Whole Slide Microscopy Image Frame Type Sequence",
		"WholeSlideMicroscopyImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x071A),
		"Image Center Point Coordinates Sequence",
		"ImageCenterPointCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x072A),
		"X Offset in Slide Coordinate System",
		"XOffsetInSlideCoordinateSystem",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x073A),
		"Y Offset in Slide Coordinate System",
		"YOffsetInSlideCoordinateSystem",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x074A),
		"Z Offset in Slide Coordinate System",
		"ZOffsetInSlideCoordinateSystem",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x08D8),
		"Pixel Spacing Sequence",
		"PixelSpacingSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x08DA),
		"Coordinate System Axis Code Sequence",
		"CoordinateSystemAxisCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x08EA),
		"Measurement Units Code Sequence",
		"MeasurementUnitsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x09F8),
		"Vital Stain Code Sequence (Trial)",
		"VitalStainCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1001),
		"Requested Procedure ID",
		"RequestedProcedureID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1002),
		"Reason for the Requested Procedure",
		"ReasonForTheRequestedProcedure",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1003),
		"Requested Procedure Priority",
		"RequestedProcedurePriority",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1004),
		"Patient Transport Arrangements",
		"PatientTransportArrangements",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1005),
		"Requested Procedure Location",
		"RequestedProcedureLocation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1006),
		"Placer Order Number / Procedure",
		"PlacerOrderNumberProcedure",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1007),
		"Filler Order Number / Procedure",
		"FillerOrderNumberProcedure",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1008),
		"Confidentiality Code",
		"ConfidentialityCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1009),
		"Reporting Priority",
		"ReportingPriority",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x100A),
		"Reason for Requested Procedure Code Sequence",
		"ReasonForRequestedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1010),
		"Names of Intended Recipients of Results",
		"NamesOfIntendedRecipientsOfResults",
		vm.VM1N,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1011),
		"Intended Recipients of Results Identification Sequence",
		"IntendedRecipientsOfResultsIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1012),
		"Reason For Performed Procedure Code Sequence",
		"ReasonForPerformedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1060),
		"Requested Procedure Description (Trial)",
		"RequestedProcedureDescriptionTrial",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1101),
		"Person Identification Code Sequence",
		"PersonIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1102),
		"Person's Address",
		"PersonAddress",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1103),
		"Person's Telephone Numbers",
		"PersonTelephoneNumbers",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1104),
		"Person's Telecom Information",
		"PersonTelecomInformation",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1400),
		"Requested Procedure Comments",
		"RequestedProcedureComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2001),
		"Reason for the Imaging Service Request",
		"ReasonForTheImagingServiceRequest",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2004),
		"Issue Date of Imaging Service Request",
		"IssueDateOfImagingServiceRequest",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2005),
		"Issue Time of Imaging Service Request",
		"IssueTimeOfImagingServiceRequest",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2006),
		"Placer Order Number / Imaging Service Request (Retired)",
		"PlacerOrderNumberImagingServiceRequestRetired",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2007),
		"Filler Order Number / Imaging Service Request (Retired)",
		"FillerOrderNumberImagingServiceRequestRetired",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2008),
		"Order Entered By",
		"OrderEnteredBy",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2009),
		"Order Enterer's Location",
		"OrderEntererLocation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2010),
		"Order Callback Phone Number",
		"OrderCallbackPhoneNumber",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2011),
		"Order Callback Telecom Information",
		"OrderCallbackTelecomInformation",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2016),
		"Placer Order Number / Imaging Service Request",
		"PlacerOrderNumberImagingServiceRequest",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2017),
		"Filler Order Number / Imaging Service Request",
		"FillerOrderNumberImagingServiceRequest",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2400),
		"Imaging Service Request Comments",
		"ImagingServiceRequestComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x3001),
		"Confidentiality Constraint on Patient Data Description",
		"ConfidentialityConstraintOnPatientDataDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4001),
		"General Purpose Scheduled Procedure Step Status",
		"GeneralPurposeScheduledProcedureStepStatus",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4002),
		"General Purpose Performed Procedure Step Status",
		"GeneralPurposePerformedProcedureStepStatus",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4003),
		"General Purpose Scheduled Procedure Step Priority",
		"GeneralPurposeScheduledProcedureStepPriority",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4004),
		"Scheduled Processing Applications Code Sequence",
		"ScheduledProcessingApplicationsCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4005),
		"Scheduled Procedure Step Start DateTime",
		"ScheduledProcedureStepStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4006),
		"Multiple Copies Flag",
		"MultipleCopiesFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4007),
		"Performed Processing Applications Code Sequence",
		"PerformedProcessingApplicationsCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4008),
		"Scheduled Procedure Step Expiration DateTime",
		"ScheduledProcedureStepExpirationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4009),
		"Human Performer Code Sequence",
		"HumanPerformerCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4010),
		"Scheduled Procedure Step Modification DateTime",
		"ScheduledProcedureStepModificationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4011),
		"Expected Completion DateTime",
		"ExpectedCompletionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4015),
		"Resulting General Purpose Performed Procedure Steps Sequence",
		"ResultingGeneralPurposePerformedProcedureStepsSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4016),
		"Referenced General Purpose Scheduled Procedure Step Sequence",
		"ReferencedGeneralPurposeScheduledProcedureStepSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4018),
		"Scheduled Workitem Code Sequence",
		"ScheduledWorkitemCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4019),
		"Performed Workitem Code Sequence",
		"PerformedWorkitemCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4020),
		"Input Availability Flag",
		"InputAvailabilityFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4021),
		"Input Information Sequence",
		"InputInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4022),
		"Relevant Information Sequence",
		"RelevantInformationSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4023),
		"Referenced General Purpose Scheduled Procedure Step Transaction UID",
		"ReferencedGeneralPurposeScheduledProcedureStepTransactionUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4025),
		"Scheduled Station Name Code Sequence",
		"ScheduledStationNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4026),
		"Scheduled Station Class Code Sequence",
		"ScheduledStationClassCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4027),
		"Scheduled Station Geographic Location Code Sequence",
		"ScheduledStationGeographicLocationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4028),
		"Performed Station Name Code Sequence",
		"PerformedStationNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4029),
		"Performed Station Class Code Sequence",
		"PerformedStationClassCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4030),
		"Performed Station Geographic Location Code Sequence",
		"PerformedStationGeographicLocationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4031),
		"Requested Subsequent Workitem Code Sequence",
		"RequestedSubsequentWorkitemCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4032),
		"Non-DICOM Output Code Sequence",
		"NonDICOMOutputCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4033),
		"Output Information Sequence",
		"OutputInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4034),
		"Scheduled Human Performers Sequence",
		"ScheduledHumanPerformersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4035),
		"Actual Human Performers Sequence",
		"ActualHumanPerformersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4036),
		"Human Performer's Organization",
		"HumanPerformerOrganization",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4037),
		"Human Performer's Name",
		"HumanPerformerName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4040),
		"Raw Data Handling",
		"RawDataHandling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4041),
		"Input Readiness State",
		"InputReadinessState",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4050),
		"Performed Procedure Step Start DateTime",
		"PerformedProcedureStepStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4051),
		"Performed Procedure Step End DateTime",
		"PerformedProcedureStepEndDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4052),
		"Procedure Step Cancellation DateTime",
		"ProcedureStepCancellationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4070),
		"Output Destination Sequence",
		"OutputDestinationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4071),
		"DICOM Storage Sequence",
		"DICOMStorageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4072),
		"STOW-RS Storage Sequence",
		"STOWRSStorageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4073),
		"Storage URL",
		"StorageURL",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4074),
		"XDS Storage Sequence",
		"XDSStorageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x8302),
		"Entrance Dose in mGy",
		"EntranceDoseInmGy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x8303),
		"Entrance Dose Derivation",
		"EntranceDoseDerivation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9092),
		"Parametric Map Frame Type Sequence",
		"ParametricMapFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9094),
		"Referenced Image Real World Value Mapping Sequence",
		"ReferencedImageRealWorldValueMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9096),
		"Real World Value Mapping Sequence",
		"RealWorldValueMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9098),
		"Pixel Value Mapping Code Sequence",
		"PixelValueMappingCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9210),
		"LUT Label",
		"LUTLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9211),
		"Real World Value Last Value Mapped",
		"RealWorldValueLastValueMapped",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9212),
		"Real World Value LUT Data",
		"RealWorldValueLUTData",
		vm.VM1N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9213),
		"Double Float Real World Value Last Value Mapped",
		"DoubleFloatRealWorldValueLastValueMapped",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9214),
		"Double Float Real World Value First Value Mapped",
		"DoubleFloatRealWorldValueFirstValueMapped",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9216),
		"Real World Value First Value Mapped",
		"RealWorldValueFirstValueMapped",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9220),
		"Quantity Definition Sequence",
		"QuantityDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9224),
		"Real World Value Intercept",
		"RealWorldValueIntercept",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9225),
		"Real World Value Slope",
		"RealWorldValueSlope",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA007),
		"Findings Flag (Trial)",
		"FindingsFlagTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA010),
		"Relationship Type",
		"RelationshipType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA020),
		"Findings Sequence (Trial)",
		"FindingsSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA021),
		"Findings Group UID (Trial)",
		"FindingsGroupUIDTrial",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA022),
		"Referenced Findings Group UID (Trial)",
		"ReferencedFindingsGroupUIDTrial",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA023),
		"Findings Group Recording Date (Trial)",
		"FindingsGroupRecordingDateTrial",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA024),
		"Findings Group Recording Time (Trial)",
		"FindingsGroupRecordingTimeTrial",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA026),
		"Findings Source Category Code Sequence (Trial)",
		"FindingsSourceCategoryCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA027),
		"Verifying Organization",
		"VerifyingOrganization",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA028),
		"Documenting Organization Identifier Code Sequence (Trial)",
		"DocumentingOrganizationIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA030),
		"Verification DateTime",
		"VerificationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA032),
		"Observation DateTime",
		"ObservationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA033),
		"Observation Start DateTime",
		"ObservationStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA034),
		"Effective Start DateTime",
		"EffectiveStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA035),
		"Effective Stop DateTime",
		"EffectiveStopDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA040),
		"Value Type",
		"ValueType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA043),
		"Concept Name Code Sequence",
		"ConceptNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA047),
		"Measurement Precision Description (Trial)",
		"MeasurementPrecisionDescriptionTrial",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA050),
		"Continuity Of Content",
		"ContinuityOfContent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA057),
		"Urgency or Priority Alerts (Trial)",
		"UrgencyOrPriorityAlertsTrial",
		vm.VM1N,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA060),
		"Sequencing Indicator (Trial)",
		"SequencingIndicatorTrial",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA066),
		"Document Identifier Code Sequence (Trial)",
		"DocumentIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA067),
		"Document Author (Trial)",
		"DocumentAuthorTrial",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA068),
		"Document Author Identifier Code Sequence (Trial)",
		"DocumentAuthorIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA070),
		"Identifier Code Sequence (Trial)",
		"IdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA073),
		"Verifying Observer Sequence",
		"VerifyingObserverSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA074),
		"Object Binary Identifier (Trial)",
		"ObjectBinaryIdentifierTrial",
		vm.VM1,
		true,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA075),
		"Verifying Observer Name",
		"VerifyingObserverName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA076),
		"Documenting Observer Identifier Code Sequence (Trial)",
		"DocumentingObserverIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA078),
		"Author Observer Sequence",
		"AuthorObserverSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA07A),
		"Participant Sequence",
		"ParticipantSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA07C),
		"Custodial Organization Sequence",
		"CustodialOrganizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA080),
		"Participation Type",
		"ParticipationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA082),
		"Participation DateTime",
		"ParticipationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA084),
		"Observer Type",
		"ObserverType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA085),
		"Procedure Identifier Code Sequence (Trial)",
		"ProcedureIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA088),
		"Verifying Observer Identification Code Sequence",
		"VerifyingObserverIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA089),
		"Object Directory Binary Identifier (Trial)",
		"ObjectDirectoryBinaryIdentifierTrial",
		vm.VM1,
		true,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA090),
		"Equivalent CDA Document Sequence",
		"EquivalentCDADocumentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA0B0),
		"Referenced Waveform Channels",
		"ReferencedWaveformChannels",
		vm.VM22N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA110),
		"Date of Document or Verbal Transaction (Trial)",
		"DateOfDocumentOrVerbalTransactionTrial",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA112),
		"Time of Document Creation or Verbal Transaction (Trial)",
		"TimeOfDocumentCreationOrVerbalTransactionTrial",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA120),
		"DateTime",
		"DateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA121),
		"Date",
		"Date",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA122),
		"Time",
		"Time",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA123),
		"Person Name",
		"PersonName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA124),
		"UID",
		"UID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA125),
		"Report Status ID (Trial)",
		"ReportStatusIDTrial",
		vm.VM2,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA130),
		"Temporal Range Type",
		"TemporalRangeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA132),
		"Referenced Sample Positions",
		"ReferencedSamplePositions",
		vm.VM1N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA136),
		"Referenced Frame Numbers",
		"ReferencedFrameNumbers",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA138),
		"Referenced Time Offsets",
		"ReferencedTimeOffsets",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA13A),
		"Referenced DateTime",
		"ReferencedDateTime",
		vm.VM1N,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA160),
		"Text Value",
		"TextValue",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA161),
		"Floating Point Value",
		"FloatingPointValue",
		vm.VM1N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA162),
		"Rational Numerator Value",
		"RationalNumeratorValue",
		vm.VM1N,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA163),
		"Rational Denominator Value",
		"RationalDenominatorValue",
		vm.VM1N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA167),
		"Observation Category Code Sequence (Trial)",
		"ObservationCategoryCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA168),
		"Concept Code Sequence",
		"ConceptCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA16A),
		"Bibliographic Citation (Trial)",
		"BibliographicCitationTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA170),
		"Purpose of Reference Code Sequence",
		"PurposeOfReferenceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA171),
		"Observation UID",
		"ObservationUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA172),
		"Referenced Observation UID (Trial)",
		"ReferencedObservationUIDTrial",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA173),
		"Referenced Observation Class (Trial)",
		"ReferencedObservationClassTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA174),
		"Referenced Object Observation Class (Trial)",
		"ReferencedObjectObservationClassTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA180),
		"Annotation Group Number",
		"AnnotationGroupNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA192),
		"Observation Date (Trial)",
		"ObservationDateTrial",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA193),
		"Observation Time (Trial)",
		"ObservationTimeTrial",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA194),
		"Measurement Automation (Trial)",
		"MeasurementAutomationTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA195),
		"Modifier Code Sequence",
		"ModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA224),
		"Identification Description (Trial)",
		"IdentificationDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA290),
		"Coordinates Set Geometric Type (Trial)",
		"CoordinatesSetGeometricTypeTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA296),
		"Algorithm Code Sequence (Trial)",
		"AlgorithmCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA297),
		"Algorithm Description (Trial)",
		"AlgorithmDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA29A),
		"Pixel Coordinates Set (Trial)",
		"PixelCoordinatesSetTrial",
		vm.VM22N,
		true,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA300),
		"Measured Value Sequence",
		"MeasuredValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA301),
		"Numeric Value Qualifier Code Sequence",
		"NumericValueQualifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA307),
		"Current Observer (Trial)",
		"CurrentObserverTrial",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA30A),
		"Numeric Value",
		"NumericValue",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA313),
		"Referenced Accession Sequence (Trial)",
		"ReferencedAccessionSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA33A),
		"Report Status Comment (Trial)",
		"ReportStatusCommentTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA340),
		"Procedure Context Sequence (Trial)",
		"ProcedureContextSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA352),
		"Verbal Source (Trial)",
		"VerbalSourceTrial",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA353),
		"Address (Trial)",
		"AddressTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA354),
		"Telephone Number (Trial)",
		"TelephoneNumberTrial",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA358),
		"Verbal Source Identifier Code Sequence (Trial)",
		"VerbalSourceIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA360),
		"Predecessor Documents Sequence",
		"PredecessorDocumentsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA370),
		"Referenced Request Sequence",
		"ReferencedRequestSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA372),
		"Performed Procedure Code Sequence",
		"PerformedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA375),
		"Current Requested Procedure Evidence Sequence",
		"CurrentRequestedProcedureEvidenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA380),
		"Report Detail Sequence (Trial)",
		"ReportDetailSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA385),
		"Pertinent Other Evidence Sequence",
		"PertinentOtherEvidenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA390),
		"HL7 Structured Document Reference Sequence",
		"HL7StructuredDocumentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA402),
		"Observation Subject UID (Trial)",
		"ObservationSubjectUIDTrial",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA403),
		"Observation Subject Class (Trial)",
		"ObservationSubjectClassTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA404),
		"Observation Subject Type Code Sequence (Trial)",
		"ObservationSubjectTypeCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA491),
		"Completion Flag",
		"CompletionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA492),
		"Completion Flag Description",
		"CompletionFlagDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA493),
		"Verification Flag",
		"VerificationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA494),
		"Archive Requested",
		"ArchiveRequested",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA496),
		"Preliminary Flag",
		"PreliminaryFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA504),
		"Content Template Sequence",
		"ContentTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA525),
		"Identical Documents Sequence",
		"IdenticalDocumentsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA600),
		"Observation Subject Context Flag (Trial)",
		"ObservationSubjectContextFlagTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA601),
		"Observer Context Flag (Trial)",
		"ObserverContextFlagTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA603),
		"Procedure Context Flag (Trial)",
		"ProcedureContextFlagTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA730),
		"Content Sequence",
		"ContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA731),
		"Relationship Sequence (Trial)",
		"RelationshipSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA732),
		"Relationship Type Code Sequence (Trial)",
		"RelationshipTypeCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA744),
		"Language Code Sequence (Trial)",
		"LanguageCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA801),
		"Tabulated Values Sequence",
		"TabulatedValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA802),
		"Number of Table Rows",
		"NumberOfTableRows",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA803),
		"Number of Table Columns",
		"NumberOfTableColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA804),
		"Table Row Number",
		"TableRowNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA805),
		"Table Column Number",
		"TableColumnNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA806),
		"Table Row Definition Sequence",
		"TableRowDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA807),
		"Table Column Definition Sequence",
		"TableColumnDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA808),
		"Cell Values Sequence",
		"CellValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA992),
		"Uniform Resource Locator (Trial)",
		"UniformResourceLocatorTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB020),
		"Waveform Annotation Sequence",
		"WaveformAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB030),
		"Structured Waveform Annotation Sequence",
		"StructuredWaveformAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB031),
		"Waveform Annotation Display Selection Sequence",
		"WaveformAnnotationDisplaySelectionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB032),
		"Referenced Montage Index",
		"ReferencedMontageIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB033),
		"Waveform Textual Annotation Sequence",
		"WaveformTextualAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB034),
		"Annotation DateTime",
		"AnnotationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB035),
		"Displayed Waveform Segment Sequence",
		"DisplayedWaveformSegmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB036),
		"Segment Definition DateTime",
		"SegmentDefinitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB037),
		"Montage Activation Sequence",
		"MontageActivationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB038),
		"Montage Activation Time Offset",
		"MontageActivationTimeOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB039),
		"Waveform Montage Sequence",
		"WaveformMontageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03A),
		"Referenced Montage Channel Number",
		"ReferencedMontageChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03B),
		"Montage Name",
		"MontageName",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03C),
		"Montage Channel Sequence",
		"MontageChannelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03D),
		"Montage Index",
		"MontageIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03E),
		"Montage Channel Number",
		"MontageChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03F),
		"Montage Channel Label",
		"MontageChannelLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB040),
		"Montage Channel Source Code Sequence",
		"MontageChannelSourceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB041),
		"Contributing Channel Sources Sequence",
		"ContributingChannelSourcesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB042),
		"Channel Weight",
		"ChannelWeight",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB00),
		"Template Identifier",
		"TemplateIdentifier",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB06),
		"Template Version",
		"TemplateVersion",
		vm.VM1,
		true,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB07),
		"Template Local Version",
		"TemplateLocalVersion",
		vm.VM1,
		true,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB0B),
		"Template Extension Flag",
		"TemplateExtensionFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB0C),
		"Template Extension Organization UID",
		"TemplateExtensionOrganizationUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB0D),
		"Template Extension Creator UID",
		"TemplateExtensionCreatorUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB73),
		"Referenced Content Item Identifier",
		"ReferencedContentItemIdentifier",
		vm.VM1N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE001),
		"HL7 Instance Identifier",
		"HL7InstanceIdentifier",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE004),
		"HL7 Document Effective Time",
		"HL7DocumentEffectiveTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE006),
		"HL7 Document Type Code Sequence",
		"HL7DocumentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE008),
		"Document Class Code Sequence",
		"DocumentClassCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE010),
		"Retrieve URI",
		"RetrieveURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE011),
		"Retrieve Location UID",
		"RetrieveLocationUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE020),
		"Type of Instances",
		"TypeOfInstances",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE021),
		"DICOM Retrieval Sequence",
		"DICOMRetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE022),
		"DICOM Media Retrieval Sequence",
		"DICOMMediaRetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE023),
		"WADO Retrieval Sequence",
		"WADORetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE024),
		"XDS Retrieval Sequence",
		"XDSRetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE025),
		"WADO-RS Retrieval Sequence",
		"WADORSRetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE030),
		"Repository Unique ID",
		"RepositoryUniqueID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE031),
		"Home Community ID",
		"HomeCommunityID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0010),
		"Document Title",
		"DocumentTitle",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0011),
		"Encapsulated Document",
		"EncapsulatedDocument",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0012),
		"MIME Type of Encapsulated Document",
		"MIMETypeOfEncapsulatedDocument",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0013),
		"Source Instance Sequence",
		"SourceInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0014),
		"List of MIME Types",
		"ListOfMIMETypes",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0015),
		"Encapsulated Document Length",
		"EncapsulatedDocumentLength",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0001),
		"Product Package Identifier",
		"ProductPackageIdentifier",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0002),
		"Substance Administration Approval",
		"SubstanceAdministrationApproval",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0003),
		"Approval Status Further Description",
		"ApprovalStatusFurtherDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0004),
		"Approval Status DateTime",
		"ApprovalStatusDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0007),
		"Product Type Code Sequence",
		"ProductTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0008),
		"Product Name",
		"ProductName",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0009),
		"Product Description",
		"ProductDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x000A),
		"Product Lot Identifier",
		"ProductLotIdentifier",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x000B),
		"Product Expiration DateTime",
		"ProductExpirationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0010),
		"Substance Administration DateTime",
		"SubstanceAdministrationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0011),
		"Substance Administration Notes",
		"SubstanceAdministrationNotes",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0012),
		"Substance Administration Device ID",
		"SubstanceAdministrationDeviceID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0013),
		"Product Parameter Sequence",
		"ProductParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0019),
		"Substance Administration Parameter Sequence",
		"SubstanceAdministrationParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0100),
		"Approval Sequence",
		"ApprovalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0101),
		"Assertion Code Sequence",
		"AssertionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0102),
		"Assertion UID",
		"AssertionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0103),
		"Asserter Identification Sequence",
		"AsserterIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0104),
		"Assertion DateTime",
		"AssertionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0105),
		"Assertion Expiration DateTime",
		"AssertionExpirationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0106),
		"Assertion Comments",
		"AssertionComments",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0107),
		"Related Assertion Sequence",
		"RelatedAssertionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0108),
		"Referenced Assertion UID",
		"ReferencedAssertionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0109),
		"Approval Subject Sequence",
		"ApprovalSubjectSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x010A),
		"Organizational Role Code Sequence",
		"OrganizationalRoleCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0110),
		"RT Assertions Sequence",
		"RTAssertionsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0012),
		"Lens Description",
		"LensDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0014),
		"Right Lens Sequence",
		"RightLensSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0015),
		"Left Lens Sequence",
		"LeftLensSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0016),
		"Unspecified Laterality Lens Sequence",
		"UnspecifiedLateralityLensSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0018),
		"Cylinder Sequence",
		"CylinderSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0028),
		"Prism Sequence",
		"PrismSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0030),
		"Horizontal Prism Power",
		"HorizontalPrismPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0032),
		"Horizontal Prism Base",
		"HorizontalPrismBase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0034),
		"Vertical Prism Power",
		"VerticalPrismPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0036),
		"Vertical Prism Base",
		"VerticalPrismBase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0038),
		"Lens Segment Type",
		"LensSegmentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0040),
		"Optical Transmittance",
		"OpticalTransmittance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0042),
		"Channel Width",
		"ChannelWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0044),
		"Pupil Size",
		"PupilSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0046),
		"Corneal Size",
		"CornealSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0047),
		"Corneal Size Sequence",
		"CornealSizeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0050),
		"Autorefraction Right Eye Sequence",
		"AutorefractionRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0052),
		"Autorefraction Left Eye Sequence",
		"AutorefractionLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0060),
		"Distance Pupillary Distance",
		"DistancePupillaryDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0062),
		"Near Pupillary Distance",
		"NearPupillaryDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0063),
		"Intermediate Pupillary Distance",
		"IntermediatePupillaryDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0064),
		"Other Pupillary Distance",
		"OtherPupillaryDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0070),
		"Keratometry Right Eye Sequence",
		"KeratometryRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0071),
		"Keratometry Left Eye Sequence",
		"KeratometryLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0074),
		"Steep Keratometric Axis Sequence",
		"SteepKeratometricAxisSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0075),
		"Radius of Curvature",
		"RadiusOfCurvature",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0076),
		"Keratometric Power",
		"KeratometricPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0077),
		"Keratometric Axis",
		"KeratometricAxis",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0080),
		"Flat Keratometric Axis Sequence",
		"FlatKeratometricAxisSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0092),
		"Background Color",
		"BackgroundColor",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0094),
		"Optotype",
		"Optotype",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0095),
		"Optotype Presentation",
		"OptotypePresentation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0097),
		"Subjective Refraction Right Eye Sequence",
		"SubjectiveRefractionRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0098),
		"Subjective Refraction Left Eye Sequence",
		"SubjectiveRefractionLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0100),
		"Add Near Sequence",
		"AddNearSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0101),
		"Add Intermediate Sequence",
		"AddIntermediateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0102),
		"Add Other Sequence",
		"AddOtherSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0104),
		"Add Power",
		"AddPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0106),
		"Viewing Distance",
		"ViewingDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0110),
		"Cornea Measurements Sequence",
		"CorneaMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0111),
		"Source of Cornea Measurement Data Code Sequence",
		"SourceOfCorneaMeasurementDataCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0112),
		"Steep Corneal Axis Sequence",
		"SteepCornealAxisSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0113),
		"Flat Corneal Axis Sequence",
		"FlatCornealAxisSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0114),
		"Corneal Power",
		"CornealPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0115),
		"Corneal Axis",
		"CornealAxis",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0116),
		"Cornea Measurement Method Code Sequence",
		"CorneaMeasurementMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0117),
		"Refractive Index of Cornea",
		"RefractiveIndexOfCornea",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0118),
		"Refractive Index of Aqueous Humor",
		"RefractiveIndexOfAqueousHumor",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0121),
		"Visual Acuity Type Code Sequence",
		"VisualAcuityTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0122),
		"Visual Acuity Right Eye Sequence",
		"VisualAcuityRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0123),
		"Visual Acuity Left Eye Sequence",
		"VisualAcuityLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0124),
		"Visual Acuity Both Eyes Open Sequence",
		"VisualAcuityBothEyesOpenSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0125),
		"Viewing Distance Type",
		"ViewingDistanceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0135),
		"Visual Acuity Modifiers",
		"VisualAcuityModifiers",
		vm.VM2,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0137),
		"Decimal Visual Acuity",
		"DecimalVisualAcuity",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0139),
		"Optotype Detailed Definition",
		"OptotypeDetailedDefinition",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0145),
		"Referenced Refractive Measurements Sequence",
		"ReferencedRefractiveMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0146),
		"Sphere Power",
		"SpherePower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0147),
		"Cylinder Power",
		"CylinderPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0201),
		"Corneal Topography Surface",
		"CornealTopographySurface",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0202),
		"Corneal Vertex Location",
		"CornealVertexLocation",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0203),
		"Pupil Centroid X-Coordinate",
		"PupilCentroidXCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0204),
		"Pupil Centroid Y-Coordinate",
		"PupilCentroidYCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0205),
		"Equivalent Pupil Radius",
		"EquivalentPupilRadius",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0207),
		"Corneal Topography Map Type Code Sequence",
		"CornealTopographyMapTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0208),
		"Vertices of the Outline of Pupil",
		"VerticesOfTheOutlineOfPupil",
		vm.VM22N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0210),
		"Corneal Topography Mapping Normals Sequence",
		"CornealTopographyMappingNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0211),
		"Maximum Corneal Curvature Sequence",
		"MaximumCornealCurvatureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0212),
		"Maximum Corneal Curvature",
		"MaximumCornealCurvature",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0213),
		"Maximum Corneal Curvature Location",
		"MaximumCornealCurvatureLocation",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0215),
		"Minimum Keratometric Sequence",
		"MinimumKeratometricSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0218),
		"Simulated Keratometric Cylinder Sequence",
		"SimulatedKeratometricCylinderSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0220),
		"Average Corneal Power",
		"AverageCornealPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0224),
		"Corneal I-S Value",
		"CornealISValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0227),
		"Analyzed Area",
		"AnalyzedArea",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0230),
		"Surface Regularity Index",
		"SurfaceRegularityIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0232),
		"Surface Asymmetry Index",
		"SurfaceAsymmetryIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0234),
		"Corneal Eccentricity Index",
		"CornealEccentricityIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0236),
		"Keratoconus Prediction Index",
		"KeratoconusPredictionIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0238),
		"Decimal Potential Visual Acuity",
		"DecimalPotentialVisualAcuity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0242),
		"Corneal Topography Map Quality Evaluation",
		"CornealTopographyMapQualityEvaluation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0244),
		"Source Image Corneal Processed Data Sequence",
		"SourceImageCornealProcessedDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0247),
		"Corneal Point Location",
		"CornealPointLocation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0248),
		"Corneal Point Estimated",
		"CornealPointEstimated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0249),
		"Axial Power",
		"AxialPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0250),
		"Tangential Power",
		"TangentialPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0251),
		"Refractive Power",
		"RefractivePower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0252),
		"Relative Elevation",
		"RelativeElevation",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0253),
		"Corneal Wavefront",
		"CornealWavefront",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0001),
		"Imaged Volume Width",
		"ImagedVolumeWidth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0002),
		"Imaged Volume Height",
		"ImagedVolumeHeight",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0003),
		"Imaged Volume Depth",
		"ImagedVolumeDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0006),
		"Total Pixel Matrix Columns",
		"TotalPixelMatrixColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0007),
		"Total Pixel Matrix Rows",
		"TotalPixelMatrixRows",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0008),
		"Total Pixel Matrix Origin Sequence",
		"TotalPixelMatrixOriginSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0010),
		"Specimen Label in Image",
		"SpecimenLabelInImage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0011),
		"Focus Method",
		"FocusMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0012),
		"Extended Depth of Field",
		"ExtendedDepthOfField",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0013),
		"Number of Focal Planes",
		"NumberOfFocalPlanes",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0014),
		"Distance Between Focal Planes",
		"DistanceBetweenFocalPlanes",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0015),
		"Recommended Absent Pixel CIELab Value",
		"RecommendedAbsentPixelCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0100),
		"Illuminator Type Code Sequence",
		"IlluminatorTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0102),
		"Image Orientation (Slide)",
		"ImageOrientationSlide",
		vm.VM6,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0105),
		"Optical Path Sequence",
		"OpticalPathSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0106),
		"Optical Path Identifier",
		"OpticalPathIdentifier",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0107),
		"Optical Path Description",
		"OpticalPathDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0108),
		"Illumination Color Code Sequence",
		"IlluminationColorCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0110),
		"Specimen Reference Sequence",
		"SpecimenReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0111),
		"Condenser Lens Power",
		"CondenserLensPower",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0112),
		"Objective Lens Power",
		"ObjectiveLensPower",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0113),
		"Objective Lens Numerical Aperture",
		"ObjectiveLensNumericalAperture",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0114),
		"Confocal Mode",
		"ConfocalMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0115),
		"Tissue Location",
		"TissueLocation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0116),
		"Confocal Microscopy Image Frame Type Sequence",
		"ConfocalMicroscopyImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0117),
		"Image Acquisition Depth",
		"ImageAcquisitionDepth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0120),
		"Palette Color Lookup Table Sequence",
		"PaletteColorLookupTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0200),
		"Referenced Image Navigation Sequence",
		"ReferencedImageNavigationSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0201),
		"Top Left Hand Corner of Localizer Area",
		"TopLeftHandCornerOfLocalizerArea",
		vm.VM2,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0202),
		"Bottom Right Hand Corner of Localizer Area",
		"BottomRightHandCornerOfLocalizerArea",
		vm.VM2,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0207),
		"Optical Path Identification Sequence",
		"OpticalPathIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x021A),
		"Plane Position (Slide) Sequence",
		"PlanePositionSlideSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x021E),
		"Column Position In Total Image Pixel Matrix",
		"ColumnPositionInTotalImagePixelMatrix",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x021F),
		"Row Position In Total Image Pixel Matrix",
		"RowPositionInTotalImagePixelMatrix",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0301),
		"Pixel Origin Interpretation",
		"PixelOriginInterpretation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0302),
		"Number of Optical Paths",
		"NumberOfOpticalPaths",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0303),
		"Total Pixel Matrix Focal Planes",
		"TotalPixelMatrixFocalPlanes",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0304),
		"Tiles Overlap",
		"TilesOverlap",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0004),
		"Calibration Image",
		"CalibrationImage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0010),
		"Device Sequence",
		"DeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0012),
		"Container Component Type Code Sequence",
		"ContainerComponentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0013),
		"Container Component Thickness",
		"ContainerComponentThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0014),
		"Device Length",
		"DeviceLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0015),
		"Container Component Width",
		"ContainerComponentWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0016),
		"Device Diameter",
		"DeviceDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0017),
		"Device Diameter Units",
		"DeviceDiameterUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0018),
		"Device Volume",
		"DeviceVolume",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0019),
		"Inter-Marker Distance",
		"InterMarkerDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001A),
		"Container Component Material",
		"ContainerComponentMaterial",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001B),
		"Container Component ID",
		"ContainerComponentID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001C),
		"Container Component Length",
		"ContainerComponentLength",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001D),
		"Container Component Diameter",
		"ContainerComponentDiameter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001E),
		"Container Component Description",
		"ContainerComponentDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0020),
		"Device Description",
		"DeviceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0021),
		"Long Device Description",
		"LongDeviceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0001),
		"Contrast/Bolus Ingredient Percent by Volume",
		"ContrastBolusIngredientPercentByVolume",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0002),
		"OCT Focal Distance",
		"OCTFocalDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0003),
		"Beam Spot Size",
		"BeamSpotSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0004),
		"Effective Refractive Index",
		"EffectiveRefractiveIndex",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0006),
		"OCT Acquisition Domain",
		"OCTAcquisitionDomain",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0007),
		"OCT Optical Center Wavelength",
		"OCTOpticalCenterWavelength",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0008),
		"Axial Resolution",
		"AxialResolution",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0009),
		"Ranging Depth",
		"RangingDepth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0011),
		"A-line Rate",
		"ALineRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0012),
		"A-lines Per Frame",
		"ALinesPerFrame",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0013),
		"Catheter Rotational Rate",
		"CatheterRotationalRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0014),
		"A-line Pixel Spacing",
		"ALinePixelSpacing",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0016),
		"Mode of Percutaneous Access Sequence",
		"ModeOfPercutaneousAccessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0025),
		"Intravascular OCT Frame Type Sequence",
		"IntravascularOCTFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0026),
		"OCT Z Offset Applied",
		"OCTZOffsetApplied",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0027),
		"Intravascular Frame Content Sequence",
		"IntravascularFrameContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0028),
		"Intravascular Longitudinal Distance",
		"IntravascularLongitudinalDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0029),
		"Intravascular OCT Frame Content Sequence",
		"IntravascularOCTFrameContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0030),
		"OCT Z Offset Correction",
		"OCTZOffsetCorrection",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0031),
		"Catheter Direction of Rotation",
		"CatheterDirectionOfRotation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0033),
		"Seam Line Location",
		"SeamLineLocation",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0034),
		"First A-line Location",
		"FirstALineLocation",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0036),
		"Seam Line Index",
		"SeamLineIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0038),
		"Number of Padded A-lines",
		"NumberOfPaddedALines",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0039),
		"Interpolation Type",
		"InterpolationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x003A),
		"Refractive Index Applied",
		"RefractiveIndexApplied",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0010),
		"Energy Window Vector",
		"EnergyWindowVector",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0011),
		"Number of Energy Windows",
		"NumberOfEnergyWindows",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0012),
		"Energy Window Information Sequence",
		"EnergyWindowInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0013),
		"Energy Window Range Sequence",
		"EnergyWindowRangeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0014),
		"Energy Window Lower Limit",
		"EnergyWindowLowerLimit",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0015),
		"Energy Window Upper Limit",
		"EnergyWindowUpperLimit",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0016),
		"Radiopharmaceutical Information Sequence",
		"RadiopharmaceuticalInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0017),
		"Residual Syringe Counts",
		"ResidualSyringeCounts",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0018),
		"Energy Window Name",
		"EnergyWindowName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0020),
		"Detector Vector",
		"DetectorVector",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0021),
		"Number of Detectors",
		"NumberOfDetectors",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0022),
		"Detector Information Sequence",
		"DetectorInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0030),
		"Phase Vector",
		"PhaseVector",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0031),
		"Number of Phases",
		"NumberOfPhases",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0032),
		"Phase Information Sequence",
		"PhaseInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0033),
		"Number of Frames in Phase",
		"NumberOfFramesInPhase",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0036),
		"Phase Delay",
		"PhaseDelay",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0038),
		"Pause Between Frames",
		"PauseBetweenFrames",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0039),
		"Phase Description",
		"PhaseDescription",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0050),
		"Rotation Vector",
		"RotationVector",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0051),
		"Number of Rotations",
		"NumberOfRotations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0052),
		"Rotation Information Sequence",
		"RotationInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0053),
		"Number of Frames in Rotation",
		"NumberOfFramesInRotation",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0060),
		"R-R Interval Vector",
		"RRIntervalVector",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0061),
		"Number of R-R Intervals",
		"NumberOfRRIntervals",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0062),
		"Gated Information Sequence",
		"GatedInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0063),
		"Data Information Sequence",
		"DataInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0070),
		"Time Slot Vector",
		"TimeSlotVector",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0071),
		"Number of Time Slots",
		"NumberOfTimeSlots",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0072),
		"Time Slot Information Sequence",
		"TimeSlotInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0073),
		"Time Slot Time",
		"TimeSlotTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0080),
		"Slice Vector",
		"SliceVector",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0081),
		"Number of Slices",
		"NumberOfSlices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0090),
		"Angular View Vector",
		"AngularViewVector",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0100),
		"Time Slice Vector",
		"TimeSliceVector",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0101),
		"Number of Time Slices",
		"NumberOfTimeSlices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0200),
		"Start Angle",
		"StartAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0202),
		"Type of Detector Motion",
		"TypeOfDetectorMotion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0210),
		"Trigger Vector",
		"TriggerVector",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0211),
		"Number of Triggers in Phase",
		"NumberOfTriggersInPhase",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0220),
		"View Code Sequence",
		"ViewCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0222),
		"View Modifier Code Sequence",
		"ViewModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0300),
		"Radionuclide Code Sequence",
		"RadionuclideCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0302),
		"Administration Route Code Sequence",
		"AdministrationRouteCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0304),
		"Radiopharmaceutical Code Sequence",
		"RadiopharmaceuticalCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0306),
		"Calibration Data Sequence",
		"CalibrationDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0308),
		"Energy Window Number",
		"EnergyWindowNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0400),
		"Image ID",
		"ImageID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0410),
		"Patient Orientation Code Sequence",
		"PatientOrientationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0412),
		"Patient Orientation Modifier Code Sequence",
		"PatientOrientationModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0414),
		"Patient Gantry Relationship Code Sequence",
		"PatientGantryRelationshipCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0500),
		"Slice Progression Direction",
		"SliceProgressionDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0501),
		"Scan Progression Direction",
		"ScanProgressionDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1000),
		"Series Type",
		"SeriesType",
		vm.VM2,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1001),
		"Units",
		"Units",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1002),
		"Counts Source",
		"CountsSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1004),
		"Reprojection Method",
		"ReprojectionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1006),
		"SUV Type",
		"SUVType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1100),
		"Randoms Correction Method",
		"RandomsCorrectionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1101),
		"Attenuation Correction Method",
		"AttenuationCorrectionMethod",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1102),
		"Decay Correction",
		"DecayCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1103),
		"Reconstruction Method",
		"ReconstructionMethod",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1104),
		"Detector Lines of Response Used",
		"DetectorLinesOfResponseUsed",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1105),
		"Scatter Correction Method",
		"ScatterCorrectionMethod",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1200),
		"Axial Acceptance",
		"AxialAcceptance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1201),
		"Axial Mash",
		"AxialMash",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1202),
		"Transverse Mash",
		"TransverseMash",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1203),
		"Detector Element Size",
		"DetectorElementSize",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1210),
		"Coincidence Window Width",
		"CoincidenceWindowWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1220),
		"Secondary Counts Type",
		"SecondaryCountsType",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1300),
		"Frame Reference Time",
		"FrameReferenceTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1310),
		"Primary (Prompts) Counts Accumulated",
		"PrimaryPromptsCountsAccumulated",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1311),
		"Secondary Counts Accumulated",
		"SecondaryCountsAccumulated",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1320),
		"Slice Sensitivity Factor",
		"SliceSensitivityFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1321),
		"Decay Factor",
		"DecayFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1322),
		"Dose Calibration Factor",
		"DoseCalibrationFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1323),
		"Scatter Fraction Factor",
		"ScatterFractionFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1324),
		"Dead Time Factor",
		"DeadTimeFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1330),
		"Image Index",
		"ImageIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1400),
		"Counts Included",
		"CountsIncluded",
		vm.VM1N,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1401),
		"Dead Time Correction Flag",
		"DeadTimeCorrectionFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3000),
		"Histogram Sequence",
		"HistogramSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3002),
		"Histogram Number of Bins",
		"HistogramNumberOfBins",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3004),
		"Histogram First Bin Value",
		"HistogramFirstBinValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3006),
		"Histogram Last Bin Value",
		"HistogramLastBinValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3008),
		"Histogram Bin Width",
		"HistogramBinWidth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3010),
		"Histogram Explanation",
		"HistogramExplanation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3020),
		"Histogram Data",
		"HistogramData",
		vm.VM1N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0001),
		"Segmentation Type",
		"SegmentationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0002),
		"Segment Sequence",
		"SegmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0003),
		"Segmented Property Category Code Sequence",
		"SegmentedPropertyCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0004),
		"Segment Number",
		"SegmentNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0005),
		"Segment Label",
		"SegmentLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0006),
		"Segment Description",
		"SegmentDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0007),
		"Segmentation Algorithm Identification Sequence",
		"SegmentationAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0008),
		"Segment Algorithm Type",
		"SegmentAlgorithmType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0009),
		"Segment Algorithm Name",
		"SegmentAlgorithmName",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000A),
		"Segment Identification Sequence",
		"SegmentIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000B),
		"Referenced Segment Number",
		"ReferencedSegmentNumber",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000C),
		"Recommended Display Grayscale Value",
		"RecommendedDisplayGrayscaleValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000D),
		"Recommended Display CIELab Value",
		"RecommendedDisplayCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000E),
		"Maximum Fractional Value",
		"MaximumFractionalValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000F),
		"Segmented Property Type Code Sequence",
		"SegmentedPropertyTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0010),
		"Segmentation Fractional Type",
		"SegmentationFractionalType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0011),
		"Segmented Property Type Modifier Code Sequence",
		"SegmentedPropertyTypeModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0012),
		"Used Segments Sequence",
		"UsedSegmentsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0013),
		"Segments Overlap",
		"SegmentsOverlap",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0020),
		"Tracking ID",
		"TrackingID",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0021),
		"Tracking UID",
		"TrackingUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0002),
		"Deformable Registration Sequence",
		"DeformableRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0003),
		"Source Frame of Reference UID",
		"SourceFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0005),
		"Deformable Registration Grid Sequence",
		"DeformableRegistrationGridSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0007),
		"Grid Dimensions",
		"GridDimensions",
		vm.VM3,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0008),
		"Grid Resolution",
		"GridResolution",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0009),
		"Vector Grid Data",
		"VectorGridData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x000F),
		"Pre Deformation Matrix Registration Sequence",
		"PreDeformationMatrixRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0010),
		"Post Deformation Matrix Registration Sequence",
		"PostDeformationMatrixRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0001),
		"Number of Surfaces",
		"NumberOfSurfaces",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0002),
		"Surface Sequence",
		"SurfaceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0003),
		"Surface Number",
		"SurfaceNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0004),
		"Surface Comments",
		"SurfaceComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0005),
		"Surface Offset",
		"SurfaceOffset",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0009),
		"Surface Processing",
		"SurfaceProcessing",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000A),
		"Surface Processing Ratio",
		"SurfaceProcessingRatio",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000B),
		"Surface Processing Description",
		"SurfaceProcessingDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000C),
		"Recommended Presentation Opacity",
		"RecommendedPresentationOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000D),
		"Recommended Presentation Type",
		"RecommendedPresentationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000E),
		"Finite Volume",
		"FiniteVolume",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0010),
		"Manifold",
		"Manifold",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0011),
		"Surface Points Sequence",
		"SurfacePointsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0012),
		"Surface Points Normals Sequence",
		"SurfacePointsNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0013),
		"Surface Mesh Primitives Sequence",
		"SurfaceMeshPrimitivesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0015),
		"Number of Surface Points",
		"NumberOfSurfacePoints",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0016),
		"Point Coordinates Data",
		"PointCoordinatesData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0017),
		"Point Position Accuracy",
		"PointPositionAccuracy",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0018),
		"Mean Point Distance",
		"MeanPointDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0019),
		"Maximum Point Distance",
		"MaximumPointDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001A),
		"Points Bounding Box Coordinates",
		"PointsBoundingBoxCoordinates",
		vm.VM6,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001B),
		"Axis of Rotation",
		"AxisOfRotation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001C),
		"Center of Rotation",
		"CenterOfRotation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001E),
		"Number of Vectors",
		"NumberOfVectors",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001F),
		"Vector Dimensionality",
		"VectorDimensionality",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0020),
		"Vector Accuracy",
		"VectorAccuracy",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0021),
		"Vector Coordinate Data",
		"VectorCoordinateData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0022),
		"Double Point Coordinates Data",
		"DoublePointCoordinatesData",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0023),
		"Triangle Point Index List",
		"TrianglePointIndexList",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0024),
		"Edge Point Index List",
		"EdgePointIndexList",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0025),
		"Vertex Point Index List",
		"VertexPointIndexList",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0026),
		"Triangle Strip Sequence",
		"TriangleStripSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0027),
		"Triangle Fan Sequence",
		"TriangleFanSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0028),
		"Line Sequence",
		"LineSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0029),
		"Primitive Point Index List",
		"PrimitivePointIndexList",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002A),
		"Surface Count",
		"SurfaceCount",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002B),
		"Referenced Surface Sequence",
		"ReferencedSurfaceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002C),
		"Referenced Surface Number",
		"ReferencedSurfaceNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002D),
		"Segment Surface Generation Algorithm Identification Sequence",
		"SegmentSurfaceGenerationAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002E),
		"Segment Surface Source Instance Sequence",
		"SegmentSurfaceSourceInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002F),
		"Algorithm Family Code Sequence",
		"AlgorithmFamilyCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0030),
		"Algorithm Name Code Sequence",
		"AlgorithmNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0031),
		"Algorithm Version",
		"AlgorithmVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0032),
		"Algorithm Parameters",
		"AlgorithmParameters",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0034),
		"Facet Sequence",
		"FacetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0035),
		"Surface Processing Algorithm Identification Sequence",
		"SurfaceProcessingAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0036),
		"Algorithm Name",
		"AlgorithmName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0037),
		"Recommended Point Radius",
		"RecommendedPointRadius",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0038),
		"Recommended Line Thickness",
		"RecommendedLineThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0040),
		"Long Primitive Point Index List",
		"LongPrimitivePointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0041),
		"Long Triangle Point Index List",
		"LongTrianglePointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0042),
		"Long Edge Point Index List",
		"LongEdgePointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0043),
		"Long Vertex Point Index List",
		"LongVertexPointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0101),
		"Track Set Sequence",
		"TrackSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0102),
		"Track Sequence",
		"TrackSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0103),
		"Recommended Display CIELab Value List",
		"RecommendedDisplayCIELabValueList",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0104),
		"Tracking Algorithm Identification Sequence",
		"TrackingAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0105),
		"Track Set Number",
		"TrackSetNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0106),
		"Track Set Label",
		"TrackSetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0107),
		"Track Set Description",
		"TrackSetDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0108),
		"Track Set Anatomical Type Code Sequence",
		"TrackSetAnatomicalTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0121),
		"Measurements Sequence",
		"MeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0124),
		"Track Set Statistics Sequence",
		"TrackSetStatisticsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0125),
		"Floating Point Values",
		"FloatingPointValues",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0129),
		"Track Point Index List",
		"TrackPointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0130),
		"Track Statistics Sequence",
		"TrackStatisticsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0132),
		"Measurement Values Sequence",
		"MeasurementValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0133),
		"Diffusion Acquisition Code Sequence",
		"DiffusionAcquisitionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0134),
		"Diffusion Model Code Sequence",
		"DiffusionModelCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6210),
		"Implant Size",
		"ImplantSize",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6221),
		"Implant Template Version",
		"ImplantTemplateVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6222),
		"Replaced Implant Template Sequence",
		"ReplacedImplantTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6223),
		"Implant Type",
		"ImplantType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6224),
		"Derivation Implant Template Sequence",
		"DerivationImplantTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6225),
		"Original Implant Template Sequence",
		"OriginalImplantTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6226),
		"Effective DateTime",
		"EffectiveDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6230),
		"Implant Target Anatomy Sequence",
		"ImplantTargetAnatomySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6260),
		"Information From Manufacturer Sequence",
		"InformationFromManufacturerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6265),
		"Notification From Manufacturer Sequence",
		"NotificationFromManufacturerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6270),
		"Information Issue DateTime",
		"InformationIssueDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6280),
		"Information Summary",
		"InformationSummary",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62A0),
		"Implant Regulatory Disapproval Code Sequence",
		"ImplantRegulatoryDisapprovalCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62A5),
		"Overall Template Spatial Tolerance",
		"OverallTemplateSpatialTolerance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62C0),
		"HPGL Document Sequence",
		"HPGLDocumentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62D0),
		"HPGL Document ID",
		"HPGLDocumentID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62D5),
		"HPGL Document Label",
		"HPGLDocumentLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62E0),
		"View Orientation Code Sequence",
		"ViewOrientationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62F0),
		"View Orientation Modifier Code Sequence",
		"ViewOrientationModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62F2),
		"HPGL Document Scaling",
		"HPGLDocumentScaling",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6300),
		"HPGL Document",
		"HPGLDocument",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6310),
		"HPGL Contour Pen Number",
		"HPGLContourPenNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6320),
		"HPGL Pen Sequence",
		"HPGLPenSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6330),
		"HPGL Pen Number",
		"HPGLPenNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6340),
		"HPGL Pen Label",
		"HPGLPenLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6345),
		"HPGL Pen Description",
		"HPGLPenDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6346),
		"Recommended Rotation Point",
		"RecommendedRotationPoint",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6347),
		"Bounding Rectangle",
		"BoundingRectangle",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6350),
		"Implant Template 3D Model Surface Number",
		"ImplantTemplate3DModelSurfaceNumber",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6360),
		"Surface Model Description Sequence",
		"SurfaceModelDescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6380),
		"Surface Model Label",
		"SurfaceModelLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6390),
		"Surface Model Scaling Factor",
		"SurfaceModelScalingFactor",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63A0),
		"Materials Code Sequence",
		"MaterialsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63A4),
		"Coating Materials Code Sequence",
		"CoatingMaterialsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63A8),
		"Implant Type Code Sequence",
		"ImplantTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63AC),
		"Fixation Method Code Sequence",
		"FixationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63B0),
		"Mating Feature Sets Sequence",
		"MatingFeatureSetsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63C0),
		"Mating Feature Set ID",
		"MatingFeatureSetID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63D0),
		"Mating Feature Set Label",
		"MatingFeatureSetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63E0),
		"Mating Feature Sequence",
		"MatingFeatureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63F0),
		"Mating Feature ID",
		"MatingFeatureID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6400),
		"Mating Feature Degree of Freedom Sequence",
		"MatingFeatureDegreeOfFreedomSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6410),
		"Degree of Freedom ID",
		"DegreeOfFreedomID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6420),
		"Degree of Freedom Type",
		"DegreeOfFreedomType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6430),
		"2D Mating Feature Coordinates Sequence",
		"TwoDMatingFeatureCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6440),
		"Referenced HPGL Document ID",
		"ReferencedHPGLDocumentID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6450),
		"2D Mating Point",
		"TwoDMatingPoint",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6460),
		"2D Mating Axes",
		"TwoDMatingAxes",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6470),
		"2D Degree of Freedom Sequence",
		"TwoDDegreeOfFreedomSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6490),
		"3D Degree of Freedom Axis",
		"ThreeDDegreeOfFreedomAxis",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x64A0),
		"Range of Freedom",
		"RangeOfFreedom",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x64C0),
		"3D Mating Point",
		"ThreeDMatingPoint",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x64D0),
		"3D Mating Axes",
		"ThreeDMatingAxes",
		vm.MustParse("9"),
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x64F0),
		"2D Degree of Freedom Axis",
		"TwoDDegreeOfFreedomAxis",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6500),
		"Planning Landmark Point Sequence",
		"PlanningLandmarkPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6510),
		"Planning Landmark Line Sequence",
		"PlanningLandmarkLineSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6520),
		"Planning Landmark Plane Sequence",
		"PlanningLandmarkPlaneSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6530),
		"Planning Landmark ID",
		"PlanningLandmarkID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6540),
		"Planning Landmark Description",
		"PlanningLandmarkDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6545),
		"Planning Landmark Identification Code Sequence",
		"PlanningLandmarkIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6550),
		"2D Point Coordinates Sequence",
		"TwoDPointCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6560),
		"2D Point Coordinates",
		"TwoDPointCoordinates",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6590),
		"3D Point Coordinates",
		"ThreeDPointCoordinates",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65A0),
		"2D Line Coordinates Sequence",
		"TwoDLineCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65B0),
		"2D Line Coordinates",
		"TwoDLineCoordinates",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65D0),
		"3D Line Coordinates",
		"ThreeDLineCoordinates",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65E0),
		"2D Plane Coordinates Sequence",
		"TwoDPlaneCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65F0),
		"2D Plane Intersection",
		"TwoDPlaneIntersection",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6610),
		"3D Plane Origin",
		"ThreeDPlaneOrigin",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6620),
		"3D Plane Normal",
		"ThreeDPlaneNormal",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7001),
		"Model Modification",
		"ModelModification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7002),
		"Model Mirroring",
		"ModelMirroring",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7003),
		"Model Usage Code Sequence",
		"ModelUsageCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7004),
		"Model Group UID",
		"ModelGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7005),
		"Relative URI Reference Within Encapsulated Document",
		"RelativeURIReferenceWithinEncapsulatedDocument",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0001),
		"Annotation Coordinate Type",
		"AnnotationCoordinateType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0002),
		"Annotation Group Sequence",
		"AnnotationGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0003),
		"Annotation Group UID",
		"AnnotationGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0005),
		"Annotation Group Label",
		"AnnotationGroupLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0006),
		"Annotation Group Description",
		"AnnotationGroupDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0007),
		"Annotation Group Generation Type",
		"AnnotationGroupGenerationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0008),
		"Annotation Group Algorithm Identification Sequence",
		"AnnotationGroupAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0009),
		"Annotation Property Category Code Sequence",
		"AnnotationPropertyCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000A),
		"Annotation Property Type Code Sequence",
		"AnnotationPropertyTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000B),
		"Annotation Property Type Modifier Code Sequence",
		"AnnotationPropertyTypeModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000C),
		"Number of Annotations",
		"NumberOfAnnotations",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000D),
		"Annotation Applies to All Optical Paths",
		"AnnotationAppliesToAllOpticalPaths",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000E),
		"Referenced Optical Path Identifier",
		"ReferencedOpticalPathIdentifier",
		vm.VM1N,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000F),
		"Annotation Applies to All Z Planes",
		"AnnotationAppliesToAllZPlanes",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0010),
		"Common Z Coordinate Value",
		"CommonZCoordinateValue",
		vm.VM1N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0011),
		"Annotation Index List",
		"AnnotationIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0001),
		"Graphic Annotation Sequence",
		"GraphicAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0002),
		"Graphic Layer",
		"GraphicLayer",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0003),
		"Bounding Box Annotation Units",
		"BoundingBoxAnnotationUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0004),
		"Anchor Point Annotation Units",
		"AnchorPointAnnotationUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0005),
		"Graphic Annotation Units",
		"GraphicAnnotationUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0006),
		"Unformatted Text Value",
		"UnformattedTextValue",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0008),
		"Text Object Sequence",
		"TextObjectSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0009),
		"Graphic Object Sequence",
		"GraphicObjectSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0010),
		"Bounding Box Top Left Hand Corner",
		"BoundingBoxTopLeftHandCorner",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0011),
		"Bounding Box Bottom Right Hand Corner",
		"BoundingBoxBottomRightHandCorner",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0012),
		"Bounding Box Text Horizontal Justification",
		"BoundingBoxTextHorizontalJustification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0014),
		"Anchor Point",
		"AnchorPoint",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0015),
		"Anchor Point Visibility",
		"AnchorPointVisibility",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0020),
		"Graphic Dimensions",
		"GraphicDimensions",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0021),
		"Number of Graphic Points",
		"NumberOfGraphicPoints",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0022),
		"Graphic Data",
		"GraphicData",
		vm.VM2N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0023),
		"Graphic Type",
		"GraphicType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0024),
		"Graphic Filled",
		"GraphicFilled",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0040),
		"Image Rotation (Retired)",
		"ImageRotationRetired",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0041),
		"Image Horizontal Flip",
		"ImageHorizontalFlip",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0042),
		"Image Rotation",
		"ImageRotation",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0050),
		"Displayed Area Top Left Hand Corner (Trial)",
		"DisplayedAreaTopLeftHandCornerTrial",
		vm.VM2,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0051),
		"Displayed Area Bottom Right Hand Corner (Trial)",
		"DisplayedAreaBottomRightHandCornerTrial",
		vm.VM2,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0052),
		"Displayed Area Top Left Hand Corner",
		"DisplayedAreaTopLeftHandCorner",
		vm.VM2,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0053),
		"Displayed Area Bottom Right Hand Corner",
		"DisplayedAreaBottomRightHandCorner",
		vm.VM2,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x005A),
		"Displayed Area Selection Sequence",
		"DisplayedAreaSelectionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0060),
		"Graphic Layer Sequence",
		"GraphicLayerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0062),
		"Graphic Layer Order",
		"GraphicLayerOrder",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0066),
		"Graphic Layer Recommended Display Grayscale Value",
		"GraphicLayerRecommendedDisplayGrayscaleValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0067),
		"Graphic Layer Recommended Display RGB Value",
		"GraphicLayerRecommendedDisplayRGBValue",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0068),
		"Graphic Layer Description",
		"GraphicLayerDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0080),
		"Content Label",
		"ContentLabel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0081),
		"Content Description",
		"ContentDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0082),
		"Presentation Creation Date",
		"PresentationCreationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0083),
		"Presentation Creation Time",
		"PresentationCreationTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0084),
		"Content Creator's Name",
		"ContentCreatorName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0086),
		"Content Creator's Identification Code Sequence",
		"ContentCreatorIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0087),
		"Alternate Content Description Sequence",
		"AlternateContentDescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0100),
		"Presentation Size Mode",
		"PresentationSizeMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0101),
		"Presentation Pixel Spacing",
		"PresentationPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0102),
		"Presentation Pixel Aspect Ratio",
		"PresentationPixelAspectRatio",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0103),
		"Presentation Pixel Magnification Ratio",
		"PresentationPixelMagnificationRatio",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0207),
		"Graphic Group Label",
		"GraphicGroupLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0208),
		"Graphic Group Description",
		"GraphicGroupDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0209),
		"Compound Graphic Sequence",
		"CompoundGraphicSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0226),
		"Compound Graphic Instance ID",
		"CompoundGraphicInstanceID",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0227),
		"Font Name",
		"FontName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0228),
		"Font Name Type",
		"FontNameType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0229),
		"CSS Font Name",
		"CSSFontName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0230),
		"Rotation Angle",
		"RotationAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0231),
		"Text Style Sequence",
		"TextStyleSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0232),
		"Line Style Sequence",
		"LineStyleSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0233),
		"Fill Style Sequence",
		"FillStyleSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0234),
		"Graphic Group Sequence",
		"GraphicGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0241),
		"Text Color CIELab Value",
		"TextColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0242),
		"Horizontal Alignment",
		"HorizontalAlignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0243),
		"Vertical Alignment",
		"VerticalAlignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0244),
		"Shadow Style",
		"ShadowStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0245),
		"Shadow Offset X",
		"ShadowOffsetX",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0246),
		"Shadow Offset Y",
		"ShadowOffsetY",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0247),
		"Shadow Color CIELab Value",
		"ShadowColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0248),
		"Underlined",
		"Underlined",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0249),
		"Bold",
		"Bold",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0250),
		"Italic",
		"Italic",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0251),
		"Pattern On Color CIELab Value",
		"PatternOnColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0252),
		"Pattern Off Color CIELab Value",
		"PatternOffColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0253),
		"Line Thickness",
		"LineThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0254),
		"Line Dashing Style",
		"LineDashingStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0255),
		"Line Pattern",
		"LinePattern",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0256),
		"Fill Pattern",
		"FillPattern",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0257),
		"Fill Mode",
		"FillMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0258),
		"Shadow Opacity",
		"ShadowOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0261),
		"Gap Length",
		"GapLength",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0262),
		"Diameter of Visibility",
		"DiameterOfVisibility",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0273),
		"Rotation Point",
		"RotationPoint",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0274),
		"Tick Alignment",
		"TickAlignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0278),
		"Show Tick Label",
		"ShowTickLabel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0279),
		"Tick Label Alignment",
		"TickLabelAlignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0282),
		"Compound Graphic Units",
		"CompoundGraphicUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0284),
		"Pattern On Opacity",
		"PatternOnOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0285),
		"Pattern Off Opacity",
		"PatternOffOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0287),
		"Major Ticks Sequence",
		"MajorTicksSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0288),
		"Tick Position",
		"TickPosition",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0289),
		"Tick Label",
		"TickLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0294),
		"Compound Graphic Type",
		"CompoundGraphicType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0295),
		"Graphic Group ID",
		"GraphicGroupID",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0306),
		"Shape Type",
		"ShapeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0308),
		"Registration Sequence",
		"RegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0309),
		"Matrix Registration Sequence",
		"MatrixRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030A),
		"Matrix Sequence",
		"MatrixSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030B),
		"Frame of Reference to Displayed Coordinate System Transformation Matrix",
		"FrameOfReferenceToDisplayedCoordinateSystemTransformationMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030C),
		"Frame of Reference Transformation Matrix Type",
		"FrameOfReferenceTransformationMatrixType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030D),
		"Registration Type Code Sequence",
		"RegistrationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030F),
		"Fiducial Description",
		"FiducialDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0310),
		"Fiducial Identifier",
		"FiducialIdentifier",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0311),
		"Fiducial Identifier Code Sequence",
		"FiducialIdentifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0312),
		"Contour Uncertainty Radius",
		"ContourUncertaintyRadius",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0314),
		"Used Fiducials Sequence",
		"UsedFiducialsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0315),
		"Used RT Structure Set ROI Sequence",
		"UsedRTStructureSetROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0318),
		"Graphic Coordinates Data Sequence",
		"GraphicCoordinatesDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031A),
		"Fiducial UID",
		"FiducialUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031B),
		"Referenced Fiducial UID",
		"ReferencedFiducialUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031C),
		"Fiducial Set Sequence",
		"FiducialSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031E),
		"Fiducial Sequence",
		"FiducialSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031F),
		"Fiducials Property Category Code Sequence",
		"FiducialsPropertyCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0401),
		"Graphic Layer Recommended Display CIELab Value",
		"GraphicLayerRecommendedDisplayCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0402),
		"Blending Sequence",
		"BlendingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0403),
		"Relative Opacity",
		"RelativeOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0404),
		"Referenced Spatial Registration Sequence",
		"ReferencedSpatialRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0405),
		"Blending Position",
		"BlendingPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1101),
		"Presentation Display Collection UID",
		"PresentationDisplayCollectionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1102),
		"Presentation Sequence Collection UID",
		"PresentationSequenceCollectionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1103),
		"Presentation Sequence Position Index",
		"PresentationSequencePositionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1104),
		"Rendered Image Reference Sequence",
		"RenderedImageReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1201),
		"Volumetric Presentation State Input Sequence",
		"VolumetricPresentationStateInputSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1202),
		"Presentation Input Type",
		"PresentationInputType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1203),
		"Input Sequence Position Index",
		"InputSequencePositionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1204),
		"Crop",
		"Crop",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1205),
		"Cropping Specification Index",
		"CroppingSpecificationIndex",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1206),
		"Compositing Method",
		"CompositingMethod",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1207),
		"Volumetric Presentation Input Number",
		"VolumetricPresentationInputNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1208),
		"Image Volume Geometry",
		"ImageVolumeGeometry",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1209),
		"Volumetric Presentation Input Set UID",
		"VolumetricPresentationInputSetUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x120A),
		"Volumetric Presentation Input Set Sequence",
		"VolumetricPresentationInputSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x120B),
		"Global Crop",
		"GlobalCrop",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x120C),
		"Global Cropping Specification Index",
		"GlobalCroppingSpecificationIndex",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x120D),
		"Rendering Method",
		"RenderingMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1301),
		"Volume Cropping Sequence",
		"VolumeCroppingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1302),
		"Volume Cropping Method",
		"VolumeCroppingMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1303),
		"Bounding Box Crop",
		"BoundingBoxCrop",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1304),
		"Oblique Cropping Plane Sequence",
		"ObliqueCroppingPlaneSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1305),
		"Plane",
		"Plane",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1306),
		"Plane Normal",
		"PlaneNormal",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1309),
		"Cropping Specification Number",
		"CroppingSpecificationNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1501),
		"Multi-Planar Reconstruction Style",
		"MultiPlanarReconstructionStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1502),
		"MPR Thickness Type",
		"MPRThicknessType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1503),
		"MPR Slab Thickness",
		"MPRSlabThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1505),
		"MPR Top Left Hand Corner",
		"MPRTopLeftHandCorner",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1507),
		"MPR View Width Direction",
		"MPRViewWidthDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1508),
		"MPR View Width",
		"MPRViewWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x150C),
		"Number of Volumetric Curve Points",
		"NumberOfVolumetricCurvePoints",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x150D),
		"Volumetric Curve Points",
		"VolumetricCurvePoints",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1511),
		"MPR View Height Direction",
		"MPRViewHeightDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1512),
		"MPR View Height",
		"MPRViewHeight",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1602),
		"Render Projection",
		"RenderProjection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1603),
		"Viewpoint Position",
		"ViewpointPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1604),
		"Viewpoint LookAt Point",
		"ViewpointLookAtPoint",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1605),
		"Viewpoint Up Direction",
		"ViewpointUpDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1606),
		"Render Field of View",
		"RenderFieldOfView",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1607),
		"Sampling Step Size",
		"SamplingStepSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1701),
		"Shading Style",
		"ShadingStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1702),
		"Ambient Reflection Intensity",
		"AmbientReflectionIntensity",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1703),
		"Light Direction",
		"LightDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1704),
		"Diffuse Reflection Intensity",
		"DiffuseReflectionIntensity",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1705),
		"Specular Reflection Intensity",
		"SpecularReflectionIntensity",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1706),
		"Shininess",
		"Shininess",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1801),
		"Presentation State Classification Component Sequence",
		"PresentationStateClassificationComponentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1802),
		"Component Type",
		"ComponentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1803),
		"Component Input Sequence",
		"ComponentInputSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1804),
		"Volumetric Presentation Input Index",
		"VolumetricPresentationInputIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1805),
		"Presentation State Compositor Component Sequence",
		"PresentationStateCompositorComponentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1806),
		"Weighting Transfer Function Sequence",
		"WeightingTransferFunctionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1807),
		"Weighting Lookup Table Descriptor",
		"WeightingLookupTableDescriptor",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1808),
		"Weighting Lookup Table Data",
		"WeightingLookupTableData",
		vm.VM1,
		true,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1901),
		"Volumetric Annotation Sequence",
		"VolumetricAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1903),
		"Referenced Structured Context Sequence",
		"ReferencedStructuredContextSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1904),
		"Referenced Content Item",
		"ReferencedContentItem",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1905),
		"Volumetric Presentation Input Annotation Sequence",
		"VolumetricPresentationInputAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1907),
		"Annotation Clipping",
		"AnnotationClipping",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A01),
		"Presentation Animation Style",
		"PresentationAnimationStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A03),
		"Recommended Animation Rate",
		"RecommendedAnimationRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A04),
		"Animation Curve Sequence",
		"AnimationCurveSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A05),
		"Animation Step Size",
		"AnimationStepSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A06),
		"Swivel Range",
		"SwivelRange",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A07),
		"Volumetric Curve Up Directions",
		"VolumetricCurveUpDirections",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A08),
		"Volume Stream Sequence",
		"VolumeStreamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A09),
		"RGBA Transfer Function Description",
		"RGBATransferFunctionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B01),
		"Advanced Blending Sequence",
		"AdvancedBlendingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B02),
		"Blending Input Number",
		"BlendingInputNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B03),
		"Blending Display Input Sequence",
		"BlendingDisplayInputSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B04),
		"Blending Display Sequence",
		"BlendingDisplaySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B06),
		"Blending Mode",
		"BlendingMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B07),
		"Time Series Blending",
		"TimeSeriesBlending",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B08),
		"Geometry for Display",
		"GeometryForDisplay",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B11),
		"Threshold Sequence",
		"ThresholdSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B12),
		"Threshold Value Sequence",
		"ThresholdValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B13),
		"Threshold Type",
		"ThresholdType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B14),
		"Threshold Value",
		"ThresholdValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0002),
		"Hanging Protocol Name",
		"HangingProtocolName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0004),
		"Hanging Protocol Description",
		"HangingProtocolDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0006),
		"Hanging Protocol Level",
		"HangingProtocolLevel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0008),
		"Hanging Protocol Creator",
		"HangingProtocolCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x000A),
		"Hanging Protocol Creation DateTime",
		"HangingProtocolCreationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x000C),
		"Hanging Protocol Definition Sequence",
		"HangingProtocolDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x000E),
		"Hanging Protocol User Identification Code Sequence",
		"HangingProtocolUserIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0010),
		"Hanging Protocol User Group Name",
		"HangingProtocolUserGroupName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0012),
		"Source Hanging Protocol Sequence",
		"SourceHangingProtocolSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0014),
		"Number of Priors Referenced",
		"NumberOfPriorsReferenced",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0020),
		"Image Sets Sequence",
		"ImageSetsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0022),
		"Image Set Selector Sequence",
		"ImageSetSelectorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0024),
		"Image Set Selector Usage Flag",
		"ImageSetSelectorUsageFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0026),
		"Selector Attribute",
		"SelectorAttribute",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0028),
		"Selector Value Number",
		"SelectorValueNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0030),
		"Time Based Image Sets Sequence",
		"TimeBasedImageSetsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0032),
		"Image Set Number",
		"ImageSetNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0034),
		"Image Set Selector Category",
		"ImageSetSelectorCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0038),
		"Relative Time",
		"RelativeTime",
		vm.VM2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x003A),
		"Relative Time Units",
		"RelativeTimeUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x003C),
		"Abstract Prior Value",
		"AbstractPriorValue",
		vm.VM2,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x003E),
		"Abstract Prior Code Sequence",
		"AbstractPriorCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0040),
		"Image Set Label",
		"ImageSetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0050),
		"Selector Attribute VR",
		"SelectorAttributeVR",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0052),
		"Selector Sequence Pointer",
		"SelectorSequencePointer",
		vm.VM1N,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0054),
		"Selector Sequence Pointer Private Creator",
		"SelectorSequencePointerPrivateCreator",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0056),
		"Selector Attribute Private Creator",
		"SelectorAttributePrivateCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x005E),
		"Selector AE Value",
		"SelectorAEValue",
		vm.VM1N,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x005F),
		"Selector AS Value",
		"SelectorASValue",
		vm.VM1N,
		false,
		vr.AS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0060),
		"Selector AT Value",
		"SelectorATValue",
		vm.VM1N,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0061),
		"Selector DA Value",
		"SelectorDAValue",
		vm.VM1N,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0062),
		"Selector CS Value",
		"SelectorCSValue",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0063),
		"Selector DT Value",
		"SelectorDTValue",
		vm.VM1N,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0064),
		"Selector IS Value",
		"SelectorISValue",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0065),
		"Selector OB Value",
		"SelectorOBValue",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0066),
		"Selector LO Value",
		"SelectorLOValue",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0067),
		"Selector OF Value",
		"SelectorOFValue",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0068),
		"Selector LT Value",
		"SelectorLTValue",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0069),
		"Selector OW Value",
		"SelectorOWValue",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006A),
		"Selector PN Value",
		"SelectorPNValue",
		vm.VM1N,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006B),
		"Selector TM Value",
		"SelectorTMValue",
		vm.VM1N,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006C),
		"Selector SH Value",
		"SelectorSHValue",
		vm.VM1N,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006D),
		"Selector UN Value",
		"SelectorUNValue",
		vm.VM1,
		false,
		vr.UN,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006E),
		"Selector ST Value",
		"SelectorSTValue",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006F),
		"Selector UC Value",
		"SelectorUCValue",
		vm.VM1N,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0070),
		"Selector UT Value",
		"SelectorUTValue",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0071),
		"Selector UR Value",
		"SelectorURValue",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0072),
		"Selector DS Value",
		"SelectorDSValue",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0073),
		"Selector OD Value",
		"SelectorODValue",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0074),
		"Selector FD Value",
		"SelectorFDValue",
		vm.VM1N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0075),
		"Selector OL Value",
		"SelectorOLValue",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0076),
		"Selector FL Value",
		"SelectorFLValue",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0078),
		"Selector UL Value",
		"SelectorULValue",
		vm.VM1N,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x007A),
		"Selector US Value",
		"SelectorUSValue",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x007C),
		"Selector SL Value",
		"SelectorSLValue",
		vm.VM1N,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x007E),
		"Selector SS Value",
		"SelectorSSValue",
		vm.VM1N,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x007F),
		"Selector UI Value",
		"SelectorUIValue",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0080),
		"Selector Code Sequence Value",
		"SelectorCodeSequenceValue",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0081),
		"Selector OV Value",
		"SelectorOVValue",
		vm.VM1,
		false,
		vr.OV,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0082),
		"Selector SV Value",
		"SelectorSVValue",
		vm.VM1N,
		false,
		vr.SV,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0083),
		"Selector UV Value",
		"SelectorUVValue",
		vm.VM1N,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0100),
		"Number of Screens",
		"NumberOfScreens",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0102),
		"Nominal Screen Definition Sequence",
		"NominalScreenDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0104),
		"Number of Vertical Pixels",
		"NumberOfVerticalPixels",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0106),
		"Number of Horizontal Pixels",
		"NumberOfHorizontalPixels",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0108),
		"Display Environment Spatial Position",
		"DisplayEnvironmentSpatialPosition",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x010A),
		"Screen Minimum Grayscale Bit Depth",
		"ScreenMinimumGrayscaleBitDepth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x010C),
		"Screen Minimum Color Bit Depth",
		"ScreenMinimumColorBitDepth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x010E),
		"Application Maximum Repaint Time",
		"ApplicationMaximumRepaintTime",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0200),
		"Display Sets Sequence",
		"DisplaySetsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0202),
		"Display Set Number",
		"DisplaySetNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0203),
		"Display Set Label",
		"DisplaySetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0204),
		"Display Set Presentation Group",
		"DisplaySetPresentationGroup",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0206),
		"Display Set Presentation Group Description",
		"DisplaySetPresentationGroupDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0208),
		"Partial Data Display Handling",
		"PartialDataDisplayHandling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0210),
		"Synchronized Scrolling Sequence",
		"SynchronizedScrollingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0212),
		"Display Set Scrolling Group",
		"DisplaySetScrollingGroup",
		vm.VM2N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0214),
		"Navigation Indicator Sequence",
		"NavigationIndicatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0216),
		"Navigation Display Set",
		"NavigationDisplaySet",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0218),
		"Reference Display Sets",
		"ReferenceDisplaySets",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0300),
		"Image Boxes Sequence",
		"ImageBoxesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0302),
		"Image Box Number",
		"ImageBoxNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0304),
		"Image Box Layout Type",
		"ImageBoxLayoutType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0306),
		"Image Box Tile Horizontal Dimension",
		"ImageBoxTileHorizontalDimension",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0308),
		"Image Box Tile Vertical Dimension",
		"ImageBoxTileVerticalDimension",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0310),
		"Image Box Scroll Direction",
		"ImageBoxScrollDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0312),
		"Image Box Small Scroll Type",
		"ImageBoxSmallScrollType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0314),
		"Image Box Small Scroll Amount",
		"ImageBoxSmallScrollAmount",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0316),
		"Image Box Large Scroll Type",
		"ImageBoxLargeScrollType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0318),
		"Image Box Large Scroll Amount",
		"ImageBoxLargeScrollAmount",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0320),
		"Image Box Overlap Priority",
		"ImageBoxOverlapPriority",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0330),
		"Cine Relative to Real-Time",
		"CineRelativeToRealTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0400),
		"Filter Operations Sequence",
		"FilterOperationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0402),
		"Filter-by Category",
		"FilterByCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0404),
		"Filter-by Attribute Presence",
		"FilterByAttributePresence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0406),
		"Filter-by Operator",
		"FilterByOperator",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0420),
		"Structured Display Background CIELab Value",
		"StructuredDisplayBackgroundCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0421),
		"Empty Image Box CIELab Value",
		"EmptyImageBoxCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0422),
		"Structured Display Image Box Sequence",
		"StructuredDisplayImageBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0424),
		"Structured Display Text Box Sequence",
		"StructuredDisplayTextBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0427),
		"Referenced First Frame Sequence",
		"ReferencedFirstFrameSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0430),
		"Image Box Synchronization Sequence",
		"ImageBoxSynchronizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0432),
		"Synchronized Image Box List",
		"SynchronizedImageBoxList",
		vm.VM2N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0434),
		"Type of Synchronization",
		"TypeOfSynchronization",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0500),
		"Blending Operation Type",
		"BlendingOperationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0510),
		"Reformatting Operation Type",
		"ReformattingOperationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0512),
		"Reformatting Thickness",
		"ReformattingThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0514),
		"Reformatting Interval",
		"ReformattingInterval",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0516),
		"Reformatting Operation Initial View Direction",
		"ReformattingOperationInitialViewDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0520),
		"3D Rendering Type",
		"ThreeDRenderingType",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0600),
		"Sorting Operations Sequence",
		"SortingOperationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0602),
		"Sort-by Category",
		"SortByCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0604),
		"Sorting Direction",
		"SortingDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0700),
		"Display Set Patient Orientation",
		"DisplaySetPatientOrientation",
		vm.VM2,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0702),
		"VOI Type",
		"VOIType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0704),
		"Pseudo-Color Type",
		"PseudoColorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0705),
		"Pseudo-Color Palette Instance Reference Sequence",
		"PseudoColorPaletteInstanceReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0706),
		"Show Grayscale Inverted",
		"ShowGrayscaleInverted",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0710),
		"Show Image True Size Flag",
		"ShowImageTrueSizeFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0712),
		"Show Graphic Annotation Flag",
		"ShowGraphicAnnotationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0714),
		"Show Patient Demographics Flag",
		"ShowPatientDemographicsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0716),
		"Show Acquisition Techniques Flag",
		"ShowAcquisitionTechniquesFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0717),
		"Display Set Horizontal Justification",
		"DisplaySetHorizontalJustification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0718),
		"Display Set Vertical Justification",
		"DisplaySetVerticalJustification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x0120),
		"Continuation Start Meterset",
		"ContinuationStartMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x0121),
		"Continuation End Meterset",
		"ContinuationEndMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1000),
		"Procedure Step State",
		"ProcedureStepState",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1002),
		"Procedure Step Progress Information Sequence",
		"ProcedureStepProgressInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1004),
		"Procedure Step Progress",
		"ProcedureStepProgress",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1006),
		"Procedure Step Progress Description",
		"ProcedureStepProgressDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1007),
		"Procedure Step Progress Parameters Sequence",
		"ProcedureStepProgressParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1008),
		"Procedure Step Communications URI Sequence",
		"ProcedureStepCommunicationsURISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x100A),
		"Contact URI",
		"ContactURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x100C),
		"Contact Display Name",
		"ContactDisplayName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x100E),
		"Procedure Step Discontinuation Reason Code Sequence",
		"ProcedureStepDiscontinuationReasonCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1020),
		"Beam Task Sequence",
		"BeamTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1022),
		"Beam Task Type",
		"BeamTaskType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1024),
		"Beam Order Index (Trial)",
		"BeamOrderIndexTrial",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1025),
		"Autosequence Flag",
		"AutosequenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1026),
		"Table Top Vertical Adjusted Position",
		"TableTopVerticalAdjustedPosition",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1027),
		"Table Top Longitudinal Adjusted Position",
		"TableTopLongitudinalAdjustedPosition",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1028),
		"Table Top Lateral Adjusted Position",
		"TableTopLateralAdjustedPosition",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x102A),
		"Patient Support Adjusted Angle",
		"PatientSupportAdjustedAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x102B),
		"Table Top Eccentric Adjusted Angle",
		"TableTopEccentricAdjustedAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x102C),
		"Table Top Pitch Adjusted Angle",
		"TableTopPitchAdjustedAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x102D),
		"Table Top Roll Adjusted Angle",
		"TableTopRollAdjustedAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1030),
		"Delivery Verification Image Sequence",
		"DeliveryVerificationImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1032),
		"Verification Image Timing",
		"VerificationImageTiming",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1034),
		"Double Exposure Flag",
		"DoubleExposureFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1036),
		"Double Exposure Ordering",
		"DoubleExposureOrdering",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1038),
		"Double Exposure Meterset (Trial)",
		"DoubleExposureMetersetTrial",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x103A),
		"Double Exposure Field Delta (Trial)",
		"DoubleExposureFieldDeltaTrial",
		vm.VM4,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1040),
		"Related Reference RT Image Sequence",
		"RelatedReferenceRTImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1042),
		"General Machine Verification Sequence",
		"GeneralMachineVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1044),
		"Conventional Machine Verification Sequence",
		"ConventionalMachineVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1046),
		"Ion Machine Verification Sequence",
		"IonMachineVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1048),
		"Failed Attributes Sequence",
		"FailedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x104A),
		"Overridden Attributes Sequence",
		"OverriddenAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x104C),
		"Conventional Control Point Verification Sequence",
		"ConventionalControlPointVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x104E),
		"Ion Control Point Verification Sequence",
		"IonControlPointVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1050),
		"Attribute Occurrence Sequence",
		"AttributeOccurrenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1052),
		"Attribute Occurrence Pointer",
		"AttributeOccurrencePointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1054),
		"Attribute Item Selector",
		"AttributeItemSelector",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1056),
		"Attribute Occurrence Private Creator",
		"AttributeOccurrencePrivateCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1057),
		"Selector Sequence Pointer Items",
		"SelectorSequencePointerItems",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1200),
		"Scheduled Procedure Step Priority",
		"ScheduledProcedureStepPriority",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1202),
		"Worklist Label",
		"WorklistLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1204),
		"Procedure Step Label",
		"ProcedureStepLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1210),
		"Scheduled Processing Parameters Sequence",
		"ScheduledProcessingParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1212),
		"Performed Processing Parameters Sequence",
		"PerformedProcessingParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1216),
		"Unified Procedure Step Performed Procedure Sequence",
		"UnifiedProcedureStepPerformedProcedureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1220),
		"Related Procedure Step Sequence",
		"RelatedProcedureStepSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1222),
		"Procedure Step Relationship Type",
		"ProcedureStepRelationshipType",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1224),
		"Replaced Procedure Step Sequence",
		"ReplacedProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1230),
		"Deletion Lock",
		"DeletionLock",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1234),
		"Receiving AE",
		"ReceivingAE",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1236),
		"Requesting AE",
		"RequestingAE",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1238),
		"Reason for Cancellation",
		"ReasonForCancellation",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1242),
		"SCP Status",
		"SCPStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1244),
		"Subscription List Status",
		"SubscriptionListStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1246),
		"Unified Procedure Step List Status",
		"UnifiedProcedureStepListStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1324),
		"Beam Order Index",
		"BeamOrderIndex",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1338),
		"Double Exposure Meterset",
		"DoubleExposureMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x133A),
		"Double Exposure Field Delta",
		"DoubleExposureFieldDelta",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1401),
		"Brachy Task Sequence",
		"BrachyTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1402),
		"Continuation Start Total Reference Air Kerma",
		"ContinuationStartTotalReferenceAirKerma",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1403),
		"Continuation End Total Reference Air Kerma",
		"ContinuationEndTotalReferenceAirKerma",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1404),
		"Continuation Pulse Number",
		"ContinuationPulseNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1405),
		"Channel Delivery Order Sequence",
		"ChannelDeliveryOrderSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1406),
		"Referenced Channel Number",
		"ReferencedChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1407),
		"Start Cumulative Time Weight",
		"StartCumulativeTimeWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1408),
		"End Cumulative Time Weight",
		"EndCumulativeTimeWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1409),
		"Omitted Channel Sequence",
		"OmittedChannelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140A),
		"Reason for Channel Omission",
		"ReasonForChannelOmission",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140B),
		"Reason for Channel Omission Description",
		"ReasonForChannelOmissionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140C),
		"Channel Delivery Order Index",
		"ChannelDeliveryOrderIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140D),
		"Channel Delivery Continuation Sequence",
		"ChannelDeliveryContinuationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140E),
		"Omitted Application Setup Sequence",
		"OmittedApplicationSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0001),
		"Implant Assembly Template Name",
		"ImplantAssemblyTemplateName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0003),
		"Implant Assembly Template Issuer",
		"ImplantAssemblyTemplateIssuer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0006),
		"Implant Assembly Template Version",
		"ImplantAssemblyTemplateVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0008),
		"Replaced Implant Assembly Template Sequence",
		"ReplacedImplantAssemblyTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x000A),
		"Implant Assembly Template Type",
		"ImplantAssemblyTemplateType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x000C),
		"Original Implant Assembly Template Sequence",
		"OriginalImplantAssemblyTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x000E),
		"Derivation Implant Assembly Template Sequence",
		"DerivationImplantAssemblyTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0010),
		"Implant Assembly Template Target Anatomy Sequence",
		"ImplantAssemblyTemplateTargetAnatomySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0020),
		"Procedure Type Code Sequence",
		"ProcedureTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0030),
		"Surgical Technique",
		"SurgicalTechnique",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0032),
		"Component Types Sequence",
		"ComponentTypesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0034),
		"Component Type Code Sequence",
		"ComponentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0036),
		"Exclusive Component Type",
		"ExclusiveComponentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0038),
		"Mandatory Component Type",
		"MandatoryComponentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0040),
		"Component Sequence",
		"ComponentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0055),
		"Component ID",
		"ComponentID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0060),
		"Component Assembly Sequence",
		"ComponentAssemblySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0070),
		"Component 1 Referenced ID",
		"Component1ReferencedID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0080),
		"Component 1 Referenced Mating Feature Set ID",
		"Component1ReferencedMatingFeatureSetID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0090),
		"Component 1 Referenced Mating Feature ID",
		"Component1ReferencedMatingFeatureID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x00A0),
		"Component 2 Referenced ID",
		"Component2ReferencedID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x00B0),
		"Component 2 Referenced Mating Feature Set ID",
		"Component2ReferencedMatingFeatureSetID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x00C0),
		"Component 2 Referenced Mating Feature ID",
		"Component2ReferencedMatingFeatureID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0001),
		"Implant Template Group Name",
		"ImplantTemplateGroupName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0010),
		"Implant Template Group Description",
		"ImplantTemplateGroupDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0020),
		"Implant Template Group Issuer",
		"ImplantTemplateGroupIssuer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0024),
		"Implant Template Group Version",
		"ImplantTemplateGroupVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0026),
		"Replaced Implant Template Group Sequence",
		"ReplacedImplantTemplateGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0028),
		"Implant Template Group Target Anatomy Sequence",
		"ImplantTemplateGroupTargetAnatomySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x002A),
		"Implant Template Group Members Sequence",
		"ImplantTemplateGroupMembersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x002E),
		"Implant Template Group Member ID",
		"ImplantTemplateGroupMemberID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0050),
		"3D Implant Template Group Member Matching Point",
		"ThreeDImplantTemplateGroupMemberMatchingPoint",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0060),
		"3D Implant Template Group Member Matching Axes",
		"ThreeDImplantTemplateGroupMemberMatchingAxes",
		vm.MustParse("9"),
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0070),
		"Implant Template Group Member Matching 2D Coordinates Sequence",
		"ImplantTemplateGroupMemberMatching2DCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0090),
		"2D Implant Template Group Member Matching Point",
		"TwoDImplantTemplateGroupMemberMatchingPoint",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00A0),
		"2D Implant Template Group Member Matching Axes",
		"TwoDImplantTemplateGroupMemberMatchingAxes",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B0),
		"Implant Template Group Variation Dimension Sequence",
		"ImplantTemplateGroupVariationDimensionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B2),
		"Implant Template Group Variation Dimension Name",
		"ImplantTemplateGroupVariationDimensionName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B4),
		"Implant Template Group Variation Dimension Rank Sequence",
		"ImplantTemplateGroupVariationDimensionRankSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B6),
		"Referenced Implant Template Group Member ID",
		"ReferencedImplantTemplateGroupMemberID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B8),
		"Implant Template Group Variation Dimension Rank",
		"ImplantTemplateGroupVariationDimensionRank",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0001),
		"Surface Scan Acquisition Type Code Sequence",
		"SurfaceScanAcquisitionTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0002),
		"Surface Scan Mode Code Sequence",
		"SurfaceScanModeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0003),
		"Registration Method Code Sequence",
		"RegistrationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0004),
		"Shot Duration Time",
		"ShotDurationTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0005),
		"Shot Offset Time",
		"ShotOffsetTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0006),
		"Surface Point Presentation Value Data",
		"SurfacePointPresentationValueData",
		vm.VM1N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0007),
		"Surface Point Color CIELab Value Data",
		"SurfacePointColorCIELabValueData",
		vm.VM33N,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0008),
		"UV Mapping Sequence",
		"UVMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0009),
		"Texture Label",
		"TextureLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0010),
		"U Value Data",
		"UValueData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0011),
		"V Value Data",
		"VValueData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0012),
		"Referenced Texture Sequence",
		"ReferencedTextureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0013),
		"Referenced Surface Data Sequence",
		"ReferencedSurfaceDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0001),
		"Assessment Summary",
		"AssessmentSummary",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0003),
		"Assessment Summary Description",
		"AssessmentSummaryDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0004),
		"Assessed SOP Instance Sequence",
		"AssessedSOPInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0005),
		"Referenced Comparison SOP Instance Sequence",
		"ReferencedComparisonSOPInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0006),
		"Number of Assessment Observations",
		"NumberOfAssessmentObservations",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0007),
		"Assessment Observations Sequence",
		"AssessmentObservationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0008),
		"Observation Significance",
		"ObservationSignificance",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x000A),
		"Observation Description",
		"ObservationDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x000C),
		"Structured Constraint Observation Sequence",
		"StructuredConstraintObservationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0010),
		"Assessed Attribute Value Sequence",
		"AssessedAttributeValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0016),
		"Assessment Set ID",
		"AssessmentSetID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0017),
		"Assessment Requester Sequence",
		"AssessmentRequesterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0018),
		"Selector Attribute Name",
		"SelectorAttributeName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0019),
		"Selector Attribute Keyword",
		"SelectorAttributeKeyword",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0021),
		"Assessment Type Code Sequence",
		"AssessmentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0022),
		"Observation Basis Code Sequence",
		"ObservationBasisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0023),
		"Assessment Label",
		"AssessmentLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0032),
		"Constraint Type",
		"ConstraintType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0033),
		"Specification Selection Guidance",
		"SpecificationSelectionGuidance",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0034),
		"Constraint Value Sequence",
		"ConstraintValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0035),
		"Recommended Default Value Sequence",
		"RecommendedDefaultValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0036),
		"Constraint Violation Significance",
		"ConstraintViolationSignificance",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0037),
		"Constraint Violation Condition",
		"ConstraintViolationCondition",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0038),
		"Modifiable Constraint Flag",
		"ModifiableConstraintFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0130),
		"Storage Media File-set ID",
		"StorageMediaFileSetID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0140),
		"Storage Media File-set UID",
		"StorageMediaFileSetUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0200),
		"Icon Image Sequence",
		"IconImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0904),
		"Topic Title",
		"TopicTitle",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0906),
		"Topic Subject",
		"TopicSubject",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0910),
		"Topic Author",
		"TopicAuthor",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0912),
		"Topic Keywords",
		"TopicKeywords",
		vm.VM132,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0100, 0x0410),
		"SOP Instance Status",
		"SOPInstanceStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0100, 0x0420),
		"SOP Authorization DateTime",
		"SOPAuthorizationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0100, 0x0424),
		"SOP Authorization Comment",
		"SOPAuthorizationComment",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0100, 0x0426),
		"Authorization Equipment Certification Number",
		"AuthorizationEquipmentCertificationNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0005),
		"MAC ID Number",
		"MACIDNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0010),
		"MAC Calculation Transfer Syntax UID",
		"MACCalculationTransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0015),
		"MAC Algorithm",
		"MACAlgorithm",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0020),
		"Data Elements Signed",
		"DataElementsSigned",
		vm.VM1N,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0100),
		"Digital Signature UID",
		"DigitalSignatureUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0105),
		"Digital Signature DateTime",
		"DigitalSignatureDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0110),
		"Certificate Type",
		"CertificateType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0115),
		"Certificate of Signer",
		"CertificateOfSigner",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0120),
		"Signature",
		"Signature",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0305),
		"Certified Timestamp Type",
		"CertifiedTimestampType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0310),
		"Certified Timestamp",
		"CertifiedTimestamp",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0401),
		"Digital Signature Purpose Code Sequence",
		"DigitalSignaturePurposeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0402),
		"Referenced Digital Signature Sequence",
		"ReferencedDigitalSignatureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0403),
		"Referenced SOP Instance MAC Sequence",
		"ReferencedSOPInstanceMACSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0404),
		"MAC",
		"MAC",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0500),
		"Encrypted Attributes Sequence",
		"EncryptedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0510),
		"Encrypted Content Transfer Syntax UID",
		"EncryptedContentTransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0520),
		"Encrypted Content",
		"EncryptedContent",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0550),
		"Modified Attributes Sequence",
		"ModifiedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0551),
		"Nonconforming Modified Attributes Sequence",
		"NonconformingModifiedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0552),
		"Nonconforming Data Element Value",
		"NonconformingDataElementValue",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0561),
		"Original Attributes Sequence",
		"OriginalAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0562),
		"Attribute Modification DateTime",
		"AttributeModificationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0563),
		"Modifying System",
		"ModifyingSystem",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0564),
		"Source of Previous Values",
		"SourceOfPreviousValues",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0565),
		"Reason for the Attribute Modification",
		"ReasonForTheAttributeModification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0600),
		"Instance Origin Status",
		"InstanceOriginStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx0)"),
		"Escape Triplet",
		"EscapeTriplet",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx1)"),
		"Run Length Triplet",
		"RunLengthTriplet",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx2)"),
		"Huffman Table Size",
		"HuffmanTableSize",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx3)"),
		"Huffman Table Triplet",
		"HuffmanTableTriplet",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx4)"),
		"Shift Table Size",
		"ShiftTableSize",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx5)"),
		"Shift Table Triplet",
		"ShiftTableTriplet",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1010,xxxx)"),
		"Zonal Map",
		"ZonalMap",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0010),
		"Number of Copies",
		"NumberOfCopies",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x001E),
		"Printer Configuration Sequence",
		"PrinterConfigurationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0020),
		"Print Priority",
		"PrintPriority",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0030),
		"Medium Type",
		"MediumType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0040),
		"Film Destination",
		"FilmDestination",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0050),
		"Film Session Label",
		"FilmSessionLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0060),
		"Memory Allocation",
		"MemoryAllocation",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0061),
		"Maximum Memory Allocation",
		"MaximumMemoryAllocation",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0062),
		"Color Image Printing Flag",
		"ColorImagePrintingFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0063),
		"Collation Flag",
		"CollationFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0065),
		"Annotation Flag",
		"AnnotationFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0067),
		"Image Overlay Flag",
		"ImageOverlayFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0069),
		"Presentation LUT Flag",
		"PresentationLUTFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x006A),
		"Image Box Presentation LUT Flag",
		"ImageBoxPresentationLUTFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A0),
		"Memory Bit Depth",
		"MemoryBitDepth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A1),
		"Printing Bit Depth",
		"PrintingBitDepth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A2),
		"Media Installed Sequence",
		"MediaInstalledSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A4),
		"Other Media Available Sequence",
		"OtherMediaAvailableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A8),
		"Supported Image Display Formats Sequence",
		"SupportedImageDisplayFormatsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0500),
		"Referenced Film Box Sequence",
		"ReferencedFilmBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0510),
		"Referenced Stored Print Sequence",
		"ReferencedStoredPrintSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0010),
		"Image Display Format",
		"ImageDisplayFormat",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0030),
		"Annotation Display Format ID",
		"AnnotationDisplayFormatID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0040),
		"Film Orientation",
		"FilmOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0050),
		"Film Size ID",
		"FilmSizeID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0052),
		"Printer Resolution ID",
		"PrinterResolutionID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0054),
		"Default Printer Resolution ID",
		"DefaultPrinterResolutionID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0060),
		"Magnification Type",
		"MagnificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0080),
		"Smoothing Type",
		"SmoothingType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x00A6),
		"Default Magnification Type",
		"DefaultMagnificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x00A7),
		"Other Magnification Types Available",
		"OtherMagnificationTypesAvailable",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x00A8),
		"Default Smoothing Type",
		"DefaultSmoothingType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x00A9),
		"Other Smoothing Types Available",
		"OtherSmoothingTypesAvailable",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0100),
		"Border Density",
		"BorderDensity",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0110),
		"Empty Image Density",
		"EmptyImageDensity",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0120),
		"Min Density",
		"MinDensity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0130),
		"Max Density",
		"MaxDensity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0140),
		"Trim",
		"Trim",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0150),
		"Configuration Information",
		"ConfigurationInformation",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0152),
		"Configuration Information Description",
		"ConfigurationInformationDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0154),
		"Maximum Collated Films",
		"MaximumCollatedFilms",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x015E),
		"Illumination",
		"Illumination",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0160),
		"Reflected Ambient Light",
		"ReflectedAmbientLight",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0376),
		"Printer Pixel Spacing",
		"PrinterPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0500),
		"Referenced Film Session Sequence",
		"ReferencedFilmSessionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0510),
		"Referenced Image Box Sequence",
		"ReferencedImageBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0520),
		"Referenced Basic Annotation Box Sequence",
		"ReferencedBasicAnnotationBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0010),
		"Image Box Position",
		"ImageBoxPosition",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0020),
		"Polarity",
		"Polarity",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0030),
		"Requested Image Size",
		"RequestedImageSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0040),
		"Requested Decimate/Crop Behavior",
		"RequestedDecimateCropBehavior",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0050),
		"Requested Resolution ID",
		"RequestedResolutionID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x00A0),
		"Requested Image Size Flag",
		"RequestedImageSizeFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x00A2),
		"Decimate/Crop Result",
		"DecimateCropResult",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0110),
		"Basic Grayscale Image Sequence",
		"BasicGrayscaleImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0111),
		"Basic Color Image Sequence",
		"BasicColorImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0130),
		"Referenced Image Overlay Box Sequence",
		"ReferencedImageOverlayBoxSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0140),
		"Referenced VOI LUT Box Sequence",
		"ReferencedVOILUTBoxSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2030, 0x0010),
		"Annotation Position",
		"AnnotationPosition",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2030, 0x0020),
		"Text String",
		"TextString",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0010),
		"Referenced Overlay Plane Sequence",
		"ReferencedOverlayPlaneSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0011),
		"Referenced Overlay Plane Groups",
		"ReferencedOverlayPlaneGroups",
		vm.VM199,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0020),
		"Overlay Pixel Data Sequence",
		"OverlayPixelDataSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0060),
		"Overlay Magnification Type",
		"OverlayMagnificationType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0070),
		"Overlay Smoothing Type",
		"OverlaySmoothingType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0072),
		"Overlay or Image Magnification",
		"OverlayOrImageMagnification",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0074),
		"Magnify to Number of Columns",
		"MagnifyToNumberOfColumns",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0080),
		"Overlay Foreground Density",
		"OverlayForegroundDensity",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0082),
		"Overlay Background Density",
		"OverlayBackgroundDensity",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0090),
		"Overlay Mode",
		"OverlayMode",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0100),
		"Threshold Density",
		"ThresholdDensity",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0500),
		"Referenced Image Box Sequence (Retired)",
		"ReferencedImageBoxSequenceRetired",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2050, 0x0010),
		"Presentation LUT Sequence",
		"PresentationLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2050, 0x0020),
		"Presentation LUT Shape",
		"PresentationLUTShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2050, 0x0500),
		"Referenced Presentation LUT Sequence",
		"ReferencedPresentationLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0010),
		"Print Job ID",
		"PrintJobID",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0020),
		"Execution Status",
		"ExecutionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0030),
		"Execution Status Info",
		"ExecutionStatusInfo",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0040),
		"Creation Date",
		"CreationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0050),
		"Creation Time",
		"CreationTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0070),
		"Originator",
		"Originator",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0140),
		"Destination AE",
		"DestinationAE",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0160),
		"Owner ID",
		"OwnerID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0170),
		"Number of Films",
		"NumberOfFilms",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0500),
		"Referenced Print Job Sequence (Pull Stored Print)",
		"ReferencedPrintJobSequencePullStoredPrint",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2110, 0x0010),
		"Printer Status",
		"PrinterStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2110, 0x0020),
		"Printer Status Info",
		"PrinterStatusInfo",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2110, 0x0030),
		"Printer Name",
		"PrinterName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x2110, 0x0099),
		"Print Queue ID",
		"PrintQueueID",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x2120, 0x0010),
		"Queue Status",
		"QueueStatus",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2120, 0x0050),
		"Print Job Description Sequence",
		"PrintJobDescriptionSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2120, 0x0070),
		"Referenced Print Job Sequence",
		"ReferencedPrintJobSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0010),
		"Print Management Capabilities Sequence",
		"PrintManagementCapabilitiesSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0015),
		"Printer Characteristics Sequence",
		"PrinterCharacteristicsSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0030),
		"Film Box Content Sequence",
		"FilmBoxContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0040),
		"Image Box Content Sequence",
		"ImageBoxContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0050),
		"Annotation Content Sequence",
		"AnnotationContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0060),
		"Image Overlay Box Content Sequence",
		"ImageOverlayBoxContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0080),
		"Presentation LUT Content Sequence",
		"PresentationLUTContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x00A0),
		"Proposed Study Sequence",
		"ProposedStudySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x00C0),
		"Original Image Sequence",
		"OriginalImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0001),
		"Label Using Information Extracted From Instances",
		"LabelUsingInformationExtractedFromInstances",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0002),
		"Label Text",
		"LabelText",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0003),
		"Label Style Selection",
		"LabelStyleSelection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0004),
		"Media Disposition",
		"MediaDisposition",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0005),
		"Barcode Value",
		"BarcodeValue",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0006),
		"Barcode Symbology",
		"BarcodeSymbology",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0007),
		"Allow Media Splitting",
		"AllowMediaSplitting",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0008),
		"Include Non-DICOM Objects",
		"IncludeNonDICOMObjects",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0009),
		"Include Display Application",
		"IncludeDisplayApplication",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000A),
		"Preserve Composite Instances After Media Creation",
		"PreserveCompositeInstancesAfterMediaCreation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000B),
		"Total Number of Pieces of Media Created",
		"TotalNumberOfPiecesOfMediaCreated",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000C),
		"Requested Media Application Profile",
		"RequestedMediaApplicationProfile",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000D),
		"Referenced Storage Media Sequence",
		"ReferencedStorageMediaSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000E),
		"Failure Attributes",
		"FailureAttributes",
		vm.VM1N,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000F),
		"Allow Lossy Compression",
		"AllowLossyCompression",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0020),
		"Request Priority",
		"RequestPriority",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0002),
		"RT Image Label",
		"RTImageLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0003),
		"RT Image Name",
		"RTImageName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0004),
		"RT Image Description",
		"RTImageDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x000A),
		"Reported Values Origin",
		"ReportedValuesOrigin",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x000C),
		"RT Image Plane",
		"RTImagePlane",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x000D),
		"X-Ray Image Receptor Translation",
		"XRayImageReceptorTranslation",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x000E),
		"X-Ray Image Receptor Angle",
		"XRayImageReceptorAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0010),
		"RT Image Orientation",
		"RTImageOrientation",
		vm.VM6,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0011),
		"Image Plane Pixel Spacing",
		"ImagePlanePixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0012),
		"RT Image Position",
		"RTImagePosition",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0020),
		"Radiation Machine Name",
		"RadiationMachineName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0022),
		"Radiation Machine SAD",
		"RadiationMachineSAD",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0024),
		"Radiation Machine SSD",
		"RadiationMachineSSD",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0026),
		"RT Image SID",
		"RTImageSID",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0028),
		"Source to Reference Object Distance",
		"SourceToReferenceObjectDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0029),
		"Fraction Number",
		"FractionNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0030),
		"Exposure Sequence",
		"ExposureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0032),
		"Meterset Exposure",
		"MetersetExposure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0034),
		"Diaphragm Position",
		"DiaphragmPosition",
		vm.VM4,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0040),
		"Fluence Map Sequence",
		"FluenceMapSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0041),
		"Fluence Data Source",
		"FluenceDataSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0042),
		"Fluence Data Scale",
		"FluenceDataScale",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0050),
		"Primary Fluence Mode Sequence",
		"PrimaryFluenceModeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0051),
		"Fluence Mode",
		"FluenceMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0052),
		"Fluence Mode ID",
		"FluenceModeID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0100),
		"Selected Frame Number",
		"SelectedFrameNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0101),
		"Selected Frame Functional Groups Sequence",
		"SelectedFrameFunctionalGroupsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0102),
		"RT Image Frame General Content Sequence",
		"RTImageFrameGeneralContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0103),
		"RT Image Frame Context Sequence",
		"RTImageFrameContextSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0104),
		"RT Image Scope Sequence",
		"RTImageScopeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0105),
		"Beam Modifier Coordinates Presence Flag",
		"BeamModifierCoordinatesPresenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0106),
		"Start Cumulative Meterset",
		"StartCumulativeMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0107),
		"Stop Cumulative Meterset",
		"StopCumulativeMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0108),
		"RT Acquisition Patient Position Sequence",
		"RTAcquisitionPatientPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0109),
		"RT Image Frame Imaging Device Position Sequence",
		"RTImageFrameImagingDevicePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010A),
		"RT Image Frame kV Radiation Acquisition Sequence",
		"RTImageFramekVRadiationAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010B),
		"RT Image Frame MV Radiation Acquisition Sequence",
		"RTImageFrameMVRadiationAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010C),
		"RT Image Frame Radiation Acquisition Sequence",
		"RTImageFrameRadiationAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010D),
		"Imaging Source Position Sequence",
		"ImagingSourcePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010E),
		"Image Receptor Position Sequence",
		"ImageReceptorPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010F),
		"Device Position to Equipment Mapping Matrix",
		"DevicePositionToEquipmentMappingMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0110),
		"Device Position Parameter Sequence",
		"DevicePositionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0111),
		"Imaging Source Location Specification Type",
		"ImagingSourceLocationSpecificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0112),
		"Imaging Device Location Matrix Sequence",
		"ImagingDeviceLocationMatrixSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0113),
		"Imaging Device Location Parameter Sequence",
		"ImagingDeviceLocationParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0114),
		"Imaging Aperture Sequence",
		"ImagingApertureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0115),
		"Imaging Aperture Specification Type",
		"ImagingApertureSpecificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0116),
		"Number of Acquisition Devices",
		"NumberOfAcquisitionDevices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0117),
		"Acquisition Device Sequence",
		"AcquisitionDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0118),
		"Acquisition Task Sequence",
		"AcquisitionTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0119),
		"Acquisition Task Workitem Code Sequence",
		"AcquisitionTaskWorkitemCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011A),
		"Acquisition Subtask Sequence",
		"AcquisitionSubtaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011B),
		"Subtask Workitem Code Sequence",
		"SubtaskWorkitemCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011C),
		"Acquisition Task Index",
		"AcquisitionTaskIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011D),
		"Acquisition Subtask Index",
		"AcquisitionSubtaskIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011E),
		"Referenced Baseline Parameters RT Radiation Instance Sequence",
		"ReferencedBaselineParametersRTRadiationInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011F),
		"Position Acquisition Template Identification Sequence",
		"PositionAcquisitionTemplateIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0120),
		"Position Acquisition Template ID",
		"PositionAcquisitionTemplateID",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0121),
		"Position Acquisition Template Name",
		"PositionAcquisitionTemplateName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0122),
		"Position Acquisition Template Code Sequence",
		"PositionAcquisitionTemplateCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0123),
		"Position Acquisition Template Description",
		"PositionAcquisitionTemplateDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0124),
		"Acquisition Task Applicability Sequence",
		"AcquisitionTaskApplicabilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0125),
		"Projection Imaging Acquisition Parameter Sequence",
		"ProjectionImagingAcquisitionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0126),
		"CT Imaging Acquisition Parameter Sequence",
		"CTImagingAcquisitionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0127),
		"KV Imaging Generation Parameters Sequence",
		"KVImagingGenerationParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0128),
		"MV Imaging Generation Parameters Sequence",
		"MVImagingGenerationParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0129),
		"Acquisition Signal Type",
		"AcquisitionSignalType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012A),
		"Acquisition Method",
		"AcquisitionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012B),
		"Scan Start Position Sequence",
		"ScanStartPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012C),
		"Scan Stop Position Sequence",
		"ScanStopPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012D),
		"Imaging Source to Beam Modifier Definition Plane Distance",
		"ImagingSourceToBeamModifierDefinitionPlaneDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012E),
		"Scan Arc Type",
		"ScanArcType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012F),
		"Detector Positioning Type",
		"DetectorPositioningType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0130),
		"Additional RT Accessory Device Sequence",
		"AdditionalRTAccessoryDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0131),
		"Device-Specific Acquisition Parameter Sequence",
		"DeviceSpecificAcquisitionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0132),
		"Referenced Position Reference Instance Sequence",
		"ReferencedPositionReferenceInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0133),
		"Energy Derivation Code Sequence",
		"EnergyDerivationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0134),
		"Maximum Cumulative Meterset Exposure",
		"MaximumCumulativeMetersetExposure",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0135),
		"Acquisition Initiation Sequence",
		"AcquisitionInitiationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0136),
		"RT Cone-Beam Imaging Geometry Sequence",
		"RTConeBeamImagingGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0001),
		"DVH Type",
		"DVHType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0002),
		"Dose Units",
		"DoseUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0004),
		"Dose Type",
		"DoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0005),
		"Spatial Transform of Dose",
		"SpatialTransformOfDose",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0006),
		"Dose Comment",
		"DoseComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0008),
		"Normalization Point",
		"NormalizationPoint",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x000A),
		"Dose Summation Type",
		"DoseSummationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x000C),
		"Grid Frame Offset Vector",
		"GridFrameOffsetVector",
		vm.VM2N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x000E),
		"Dose Grid Scaling",
		"DoseGridScaling",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0010),
		"RT Dose ROI Sequence",
		"RTDoseROISequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0012),
		"Dose Value",
		"DoseValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0014),
		"Tissue Heterogeneity Correction",
		"TissueHeterogeneityCorrection",
		vm.VM13,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0016),
		"Recommended Isodose Level Sequence",
		"RecommendedIsodoseLevelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0020),
		"Dose Unit Code Sequence",
		"DoseUnitCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0021),
		"RT Dose Interpreted Type Code Sequence",
		"RTDoseInterpretedTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0022),
		"RT Dose Interpreted Type Code Modifier Sequence",
		"RTDoseInterpretedTypeCodeModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0023),
		"Dose Radiobiological Interpretation Sequence",
		"DoseRadiobiologicalInterpretationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0024),
		"RT Dose Intent Code Sequence",
		"RTDoseIntentCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0040),
		"DVH Normalization Point",
		"DVHNormalizationPoint",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0042),
		"DVH Normalization Dose Value",
		"DVHNormalizationDoseValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0050),
		"DVH Sequence",
		"DVHSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0052),
		"DVH Dose Scaling",
		"DVHDoseScaling",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0054),
		"DVH Volume Units",
		"DVHVolumeUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0056),
		"DVH Number of Bins",
		"DVHNumberOfBins",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0058),
		"DVH Data",
		"DVHData",
		vm.VM22N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0060),
		"DVH Referenced ROI Sequence",
		"DVHReferencedROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0062),
		"DVH ROI Contribution Type",
		"DVHROIContributionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0070),
		"DVH Minimum Dose",
		"DVHMinimumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0072),
		"DVH Maximum Dose",
		"DVHMaximumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0074),
		"DVH Mean Dose",
		"DVHMeanDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0080),
		"Dose Calculation Model Sequence",
		"DoseCalculationModelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0081),
		"Dose Calculation Algorithm Sequence",
		"DoseCalculationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0082),
		"Commissioning Status",
		"CommissioningStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0083),
		"Dose Calculation Model Parameter Sequence",
		"DoseCalculationModelParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0084),
		"Dose Deposition Calculation Medium",
		"DoseDepositionCalculationMedium",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0002),
		"Structure Set Label",
		"StructureSetLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0004),
		"Structure Set Name",
		"StructureSetName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0006),
		"Structure Set Description",
		"StructureSetDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0008),
		"Structure Set Date",
		"StructureSetDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0009),
		"Structure Set Time",
		"StructureSetTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0010),
		"Referenced Frame of Reference Sequence",
		"ReferencedFrameOfReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0012),
		"RT Referenced Study Sequence",
		"RTReferencedStudySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0014),
		"RT Referenced Series Sequence",
		"RTReferencedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0016),
		"Contour Image Sequence",
		"ContourImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0018),
		"Predecessor Structure Set Sequence",
		"PredecessorStructureSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0020),
		"Structure Set ROI Sequence",
		"StructureSetROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0022),
		"ROI Number",
		"ROINumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0024),
		"Referenced Frame of Reference UID",
		"ReferencedFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0026),
		"ROI Name",
		"ROIName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0028),
		"ROI Description",
		"ROIDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x002A),
		"ROI Display Color",
		"ROIDisplayColor",
		vm.VM3,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x002C),
		"ROI Volume",
		"ROIVolume",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x002D),
		"ROI DateTime",
		"ROIDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x002E),
		"ROI Observation DateTime",
		"ROIObservationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0030),
		"RT Related ROI Sequence",
		"RTRelatedROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0033),
		"RT ROI Relationship",
		"RTROIRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0036),
		"ROI Generation Algorithm",
		"ROIGenerationAlgorithm",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0037),
		"ROI Derivation Algorithm Identification Sequence",
		"ROIDerivationAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0038),
		"ROI Generation Description",
		"ROIGenerationDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0039),
		"ROI Contour Sequence",
		"ROIContourSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0040),
		"Contour Sequence",
		"ContourSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0042),
		"Contour Geometric Type",
		"ContourGeometricType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0044),
		"Contour Slab Thickness",
		"ContourSlabThickness",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0045),
		"Contour Offset Vector",
		"ContourOffsetVector",
		vm.VM3,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0046),
		"Number of Contour Points",
		"NumberOfContourPoints",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0048),
		"Contour Number",
		"ContourNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0049),
		"Attached Contours",
		"AttachedContours",
		vm.VM1N,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004A),
		"Source Pixel Planes Characteristics Sequence",
		"SourcePixelPlanesCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004B),
		"Source Series Sequence",
		"SourceSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004C),
		"Source Series Information Sequence",
		"SourceSeriesInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004D),
		"ROI Creator Sequence",
		"ROICreatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004E),
		"ROI Interpreter Sequence",
		"ROIInterpreterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004F),
		"ROI Observation Context Code Sequence",
		"ROIObservationContextCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0050),
		"Contour Data",
		"ContourData",
		vm.VM33N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0080),
		"RT ROI Observations Sequence",
		"RTROIObservationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0082),
		"Observation Number",
		"ObservationNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0084),
		"Referenced ROI Number",
		"ReferencedROINumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0085),
		"ROI Observation Label",
		"ROIObservationLabel",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0086),
		"RT ROI Identification Code Sequence",
		"RTROIIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0088),
		"ROI Observation Description",
		"ROIObservationDescription",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00A0),
		"Related RT ROI Observations Sequence",
		"RelatedRTROIObservationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00A4),
		"RT ROI Interpreted Type",
		"RTROIInterpretedType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00A6),
		"ROI Interpreter",
		"ROIInterpreter",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B0),
		"ROI Physical Properties Sequence",
		"ROIPhysicalPropertiesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B2),
		"ROI Physical Property",
		"ROIPhysicalProperty",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B4),
		"ROI Physical Property Value",
		"ROIPhysicalPropertyValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B6),
		"ROI Elemental Composition Sequence",
		"ROIElementalCompositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B7),
		"ROI Elemental Composition Atomic Number",
		"ROIElementalCompositionAtomicNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B8),
		"ROI Elemental Composition Atomic Mass Fraction",
		"ROIElementalCompositionAtomicMassFraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B9),
		"Additional RT ROI Identification Code Sequence",
		"AdditionalRTROIIdentificationCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C0),
		"Frame of Reference Relationship Sequence",
		"FrameOfReferenceRelationshipSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C2),
		"Related Frame of Reference UID",
		"RelatedFrameOfReferenceUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C4),
		"Frame of Reference Transformation Type",
		"FrameOfReferenceTransformationType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C6),
		"Frame of Reference Transformation Matrix",
		"FrameOfReferenceTransformationMatrix",
		vm.VM16,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C8),
		"Frame of Reference Transformation Comment",
		"FrameOfReferenceTransformationComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C9),
		"Patient Location Coordinates Sequence",
		"PatientLocationCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00CA),
		"Patient Location Coordinates Code Sequence",
		"PatientLocationCoordinatesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00CB),
		"Patient Support Position Sequence",
		"PatientSupportPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0010),
		"Measured Dose Reference Sequence",
		"MeasuredDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0012),
		"Measured Dose Description",
		"MeasuredDoseDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0014),
		"Measured Dose Type",
		"MeasuredDoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0016),
		"Measured Dose Value",
		"MeasuredDoseValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0020),
		"Treatment Session Beam Sequence",
		"TreatmentSessionBeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0021),
		"Treatment Session Ion Beam Sequence",
		"TreatmentSessionIonBeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0022),
		"Current Fraction Number",
		"CurrentFractionNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0024),
		"Treatment Control Point Date",
		"TreatmentControlPointDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0025),
		"Treatment Control Point Time",
		"TreatmentControlPointTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x002A),
		"Treatment Termination Status",
		"TreatmentTerminationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x002B),
		"Treatment Termination Code",
		"TreatmentTerminationCode",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x002C),
		"Treatment Verification Status",
		"TreatmentVerificationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0030),
		"Referenced Treatment Record Sequence",
		"ReferencedTreatmentRecordSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0032),
		"Specified Primary Meterset",
		"SpecifiedPrimaryMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0033),
		"Specified Secondary Meterset",
		"SpecifiedSecondaryMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0036),
		"Delivered Primary Meterset",
		"DeliveredPrimaryMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0037),
		"Delivered Secondary Meterset",
		"DeliveredSecondaryMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x003A),
		"Specified Treatment Time",
		"SpecifiedTreatmentTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x003B),
		"Delivered Treatment Time",
		"DeliveredTreatmentTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0040),
		"Control Point Delivery Sequence",
		"ControlPointDeliverySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0041),
		"Ion Control Point Delivery Sequence",
		"IonControlPointDeliverySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0042),
		"Specified Meterset",
		"SpecifiedMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0044),
		"Delivered Meterset",
		"DeliveredMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0045),
		"Meterset Rate Set",
		"MetersetRateSet",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0046),
		"Meterset Rate Delivered",
		"MetersetRateDelivered",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0047),
		"Scan Spot Metersets Delivered",
		"ScanSpotMetersetsDelivered",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0048),
		"Dose Rate Delivered",
		"DoseRateDelivered",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0050),
		"Treatment Summary Calculated Dose Reference Sequence",
		"TreatmentSummaryCalculatedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0052),
		"Cumulative Dose to Dose Reference",
		"CumulativeDoseToDoseReference",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0054),
		"First Treatment Date",
		"FirstTreatmentDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0056),
		"Most Recent Treatment Date",
		"MostRecentTreatmentDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x005A),
		"Number of Fractions Delivered",
		"NumberOfFractionsDelivered",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0060),
		"Override Sequence",
		"OverrideSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0061),
		"Parameter Sequence Pointer",
		"ParameterSequencePointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0062),
		"Override Parameter Pointer",
		"OverrideParameterPointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0063),
		"Parameter Item Index",
		"ParameterItemIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0064),
		"Measured Dose Reference Number",
		"MeasuredDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0065),
		"Parameter Pointer",
		"ParameterPointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0066),
		"Override Reason",
		"OverrideReason",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0067),
		"Parameter Value Number",
		"ParameterValueNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0068),
		"Corrected Parameter Sequence",
		"CorrectedParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x006A),
		"Correction Value",
		"CorrectionValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0070),
		"Calculated Dose Reference Sequence",
		"CalculatedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0072),
		"Calculated Dose Reference Number",
		"CalculatedDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0074),
		"Calculated Dose Reference Description",
		"CalculatedDoseReferenceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0076),
		"Calculated Dose Reference Dose Value",
		"CalculatedDoseReferenceDoseValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0078),
		"Start Meterset",
		"StartMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x007A),
		"End Meterset",
		"EndMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0080),
		"Referenced Measured Dose Reference Sequence",
		"ReferencedMeasuredDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0082),
		"Referenced Measured Dose Reference Number",
		"ReferencedMeasuredDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0090),
		"Referenced Calculated Dose Reference Sequence",
		"ReferencedCalculatedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0092),
		"Referenced Calculated Dose Reference Number",
		"ReferencedCalculatedDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A0),
		"Beam Limiting Device Leaf Pairs Sequence",
		"BeamLimitingDeviceLeafPairsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A1),
		"Enhanced RT Beam Limiting Device Sequence",
		"EnhancedRTBeamLimitingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A2),
		"Enhanced RT Beam Limiting Opening Sequence",
		"EnhancedRTBeamLimitingOpeningSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A3),
		"Enhanced RT Beam Limiting Device Definition Flag",
		"EnhancedRTBeamLimitingDeviceDefinitionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A4),
		"Parallel RT Beam Delimiter Opening Extents",
		"ParallelRTBeamDelimiterOpeningExtents",
		vm.VM22N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00B0),
		"Recorded Wedge Sequence",
		"RecordedWedgeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00C0),
		"Recorded Compensator Sequence",
		"RecordedCompensatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00D0),
		"Recorded Block Sequence",
		"RecordedBlockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00D1),
		"Recorded Block Slab Sequence",
		"RecordedBlockSlabSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00E0),
		"Treatment Summary Measured Dose Reference Sequence",
		"TreatmentSummaryMeasuredDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00F0),
		"Recorded Snout Sequence",
		"RecordedSnoutSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00F2),
		"Recorded Range Shifter Sequence",
		"RecordedRangeShifterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00F4),
		"Recorded Lateral Spreading Device Sequence",
		"RecordedLateralSpreadingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00F6),
		"Recorded Range Modulator Sequence",
		"RecordedRangeModulatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0100),
		"Recorded Source Sequence",
		"RecordedSourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0105),
		"Source Serial Number",
		"SourceSerialNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0110),
		"Treatment Session Application Setup Sequence",
		"TreatmentSessionApplicationSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0116),
		"Application Setup Check",
		"ApplicationSetupCheck",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0120),
		"Recorded Brachy Accessory Device Sequence",
		"RecordedBrachyAccessoryDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0122),
		"Referenced Brachy Accessory Device Number",
		"ReferencedBrachyAccessoryDeviceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0130),
		"Recorded Channel Sequence",
		"RecordedChannelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0132),
		"Specified Channel Total Time",
		"SpecifiedChannelTotalTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0134),
		"Delivered Channel Total Time",
		"DeliveredChannelTotalTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0136),
		"Specified Number of Pulses",
		"SpecifiedNumberOfPulses",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0138),
		"Delivered Number of Pulses",
		"DeliveredNumberOfPulses",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x013A),
		"Specified Pulse Repetition Interval",
		"SpecifiedPulseRepetitionInterval",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x013C),
		"Delivered Pulse Repetition Interval",
		"DeliveredPulseRepetitionInterval",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0140),
		"Recorded Source Applicator Sequence",
		"RecordedSourceApplicatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0142),
		"Referenced Source Applicator Number",
		"ReferencedSourceApplicatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0150),
		"Recorded Channel Shield Sequence",
		"RecordedChannelShieldSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0152),
		"Referenced Channel Shield Number",
		"ReferencedChannelShieldNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0160),
		"Brachy Control Point Delivered Sequence",
		"BrachyControlPointDeliveredSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0162),
		"Safe Position Exit Date",
		"SafePositionExitDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0164),
		"Safe Position Exit Time",
		"SafePositionExitTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0166),
		"Safe Position Return Date",
		"SafePositionReturnDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0168),
		"Safe Position Return Time",
		"SafePositionReturnTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0171),
		"Pulse Specific Brachy Control Point Delivered Sequence",
		"PulseSpecificBrachyControlPointDeliveredSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0172),
		"Pulse Number",
		"PulseNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0173),
		"Brachy Pulse Control Point Delivered Sequence",
		"BrachyPulseControlPointDeliveredSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0200),
		"Current Treatment Status",
		"CurrentTreatmentStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0202),
		"Treatment Status Comment",
		"TreatmentStatusComment",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0220),
		"Fraction Group Summary Sequence",
		"FractionGroupSummarySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0223),
		"Referenced Fraction Number",
		"ReferencedFractionNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0224),
		"Fraction Group Type",
		"FractionGroupType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0230),
		"Beam Stopper Position",
		"BeamStopperPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0240),
		"Fraction Status Summary Sequence",
		"FractionStatusSummarySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0250),
		"Treatment Date",
		"TreatmentDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0251),
		"Treatment Time",
		"TreatmentTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0002),
		"RT Plan Label",
		"RTPlanLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0003),
		"RT Plan Name",
		"RTPlanName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0004),
		"RT Plan Description",
		"RTPlanDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0006),
		"RT Plan Date",
		"RTPlanDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0007),
		"RT Plan Time",
		"RTPlanTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0009),
		"Treatment Protocols",
		"TreatmentProtocols",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x000A),
		"Plan Intent",
		"PlanIntent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x000B),
		"Treatment Sites",
		"TreatmentSites",
		vm.VM1N,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x000C),
		"RT Plan Geometry",
		"RTPlanGeometry",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x000E),
		"Prescription Description",
		"PrescriptionDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0010),
		"Dose Reference Sequence",
		"DoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0012),
		"Dose Reference Number",
		"DoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0013),
		"Dose Reference UID",
		"DoseReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0014),
		"Dose Reference Structure Type",
		"DoseReferenceStructureType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0015),
		"Nominal Beam Energy Unit",
		"NominalBeamEnergyUnit",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0016),
		"Dose Reference Description",
		"DoseReferenceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0018),
		"Dose Reference Point Coordinates",
		"DoseReferencePointCoordinates",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x001A),
		"Nominal Prior Dose",
		"NominalPriorDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0020),
		"Dose Reference Type",
		"DoseReferenceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0021),
		"Constraint Weight",
		"ConstraintWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0022),
		"Delivery Warning Dose",
		"DeliveryWarningDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0023),
		"Delivery Maximum Dose",
		"DeliveryMaximumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0025),
		"Target Minimum Dose",
		"TargetMinimumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0026),
		"Target Prescription Dose",
		"TargetPrescriptionDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0027),
		"Target Maximum Dose",
		"TargetMaximumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0028),
		"Target Underdose Volume Fraction",
		"TargetUnderdoseVolumeFraction",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x002A),
		"Organ at Risk Full-volume Dose",
		"OrganAtRiskFullVolumeDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x002B),
		"Organ at Risk Limit Dose",
		"OrganAtRiskLimitDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x002C),
		"Organ at Risk Maximum Dose",
		"OrganAtRiskMaximumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x002D),
		"Organ at Risk Overdose Volume Fraction",
		"OrganAtRiskOverdoseVolumeFraction",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0040),
		"Tolerance Table Sequence",
		"ToleranceTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0042),
		"Tolerance Table Number",
		"ToleranceTableNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0043),
		"Tolerance Table Label",
		"ToleranceTableLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0044),
		"Gantry Angle Tolerance",
		"GantryAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0046),
		"Beam Limiting Device Angle Tolerance",
		"BeamLimitingDeviceAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0048),
		"Beam Limiting Device Tolerance Sequence",
		"BeamLimitingDeviceToleranceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004A),
		"Beam Limiting Device Position Tolerance",
		"BeamLimitingDevicePositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004B),
		"Snout Position Tolerance",
		"SnoutPositionTolerance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004C),
		"Patient Support Angle Tolerance",
		"PatientSupportAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004E),
		"Table Top Eccentric Angle Tolerance",
		"TableTopEccentricAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004F),
		"Table Top Pitch Angle Tolerance",
		"TableTopPitchAngleTolerance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0050),
		"Table Top Roll Angle Tolerance",
		"TableTopRollAngleTolerance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0051),
		"Table Top Vertical Position Tolerance",
		"TableTopVerticalPositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0052),
		"Table Top Longitudinal Position Tolerance",
		"TableTopLongitudinalPositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0053),
		"Table Top Lateral Position Tolerance",
		"TableTopLateralPositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0054),
		"Table Top Position Alignment UID",
		"TableTopPositionAlignmentUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0055),
		"RT Plan Relationship",
		"RTPlanRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0070),
		"Fraction Group Sequence",
		"FractionGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0071),
		"Fraction Group Number",
		"FractionGroupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0072),
		"Fraction Group Description",
		"FractionGroupDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0078),
		"Number of Fractions Planned",
		"NumberOfFractionsPlanned",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0079),
		"Number of Fraction Pattern Digits Per Day",
		"NumberOfFractionPatternDigitsPerDay",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x007A),
		"Repeat Fraction Cycle Length",
		"RepeatFractionCycleLength",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x007B),
		"Fraction Pattern",
		"FractionPattern",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0080),
		"Number of Beams",
		"NumberOfBeams",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0082),
		"Beam Dose Specification Point",
		"BeamDoseSpecificationPoint",
		vm.VM3,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0083),
		"Referenced Dose Reference UID",
		"ReferencedDoseReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0084),
		"Beam Dose",
		"BeamDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0086),
		"Beam Meterset",
		"BeamMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0088),
		"Beam Dose Point Depth",
		"BeamDosePointDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0089),
		"Beam Dose Point Equivalent Depth",
		"BeamDosePointEquivalentDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008A),
		"Beam Dose Point SSD",
		"BeamDosePointSSD",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008B),
		"Beam Dose Meaning",
		"BeamDoseMeaning",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008C),
		"Beam Dose Verification Control Point Sequence",
		"BeamDoseVerificationControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008D),
		"Average Beam Dose Point Depth",
		"AverageBeamDosePointDepth",
		vm.VM1,
		true,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008E),
		"Average Beam Dose Point Equivalent Depth",
		"AverageBeamDosePointEquivalentDepth",
		vm.VM1,
		true,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008F),
		"Average Beam Dose Point SSD",
		"AverageBeamDosePointSSD",
		vm.VM1,
		true,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0090),
		"Beam Dose Type",
		"BeamDoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0091),
		"Alternate Beam Dose",
		"AlternateBeamDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0092),
		"Alternate Beam Dose Type",
		"AlternateBeamDoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0093),
		"Depth Value Averaging Flag",
		"DepthValueAveragingFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0094),
		"Beam Dose Point Source to External Contour Distance",
		"BeamDosePointSourceToExternalContourDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00A0),
		"Number of Brachy Application Setups",
		"NumberOfBrachyApplicationSetups",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00A2),
		"Brachy Application Setup Dose Specification Point",
		"BrachyApplicationSetupDoseSpecificationPoint",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00A4),
		"Brachy Application Setup Dose",
		"BrachyApplicationSetupDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B0),
		"Beam Sequence",
		"BeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B2),
		"Treatment Machine Name",
		"TreatmentMachineName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B3),
		"Primary Dosimeter Unit",
		"PrimaryDosimeterUnit",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B4),
		"Source-Axis Distance",
		"SourceAxisDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B6),
		"Beam Limiting Device Sequence",
		"BeamLimitingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B8),
		"RT Beam Limiting Device Type",
		"RTBeamLimitingDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00BA),
		"Source to Beam Limiting Device Distance",
		"SourceToBeamLimitingDeviceDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00BB),
		"Isocenter to Beam Limiting Device Distance",
		"IsocenterToBeamLimitingDeviceDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00BC),
		"Number of Leaf/Jaw Pairs",
		"NumberOfLeafJawPairs",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00BE),
		"Leaf Position Boundaries",
		"LeafPositionBoundaries",
		vm.VM3N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C0),
		"Beam Number",
		"BeamNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C2),
		"Beam Name",
		"BeamName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C3),
		"Beam Description",
		"BeamDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C4),
		"Beam Type",
		"BeamType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C5),
		"Beam Delivery Duration Limit",
		"BeamDeliveryDurationLimit",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C6),
		"Radiation Type",
		"RadiationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C7),
		"High-Dose Technique Type",
		"HighDoseTechniqueType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C8),
		"Reference Image Number",
		"ReferenceImageNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00CA),
		"Planned Verification Image Sequence",
		"PlannedVerificationImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00CC),
		"Imaging Device-Specific Acquisition Parameters",
		"ImagingDeviceSpecificAcquisitionParameters",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00CE),
		"Treatment Delivery Type",
		"TreatmentDeliveryType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D0),
		"Number of Wedges",
		"NumberOfWedges",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D1),
		"Wedge Sequence",
		"WedgeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D2),
		"Wedge Number",
		"WedgeNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D3),
		"Wedge Type",
		"WedgeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D4),
		"Wedge ID",
		"WedgeID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D5),
		"Wedge Angle",
		"WedgeAngle",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D6),
		"Wedge Factor",
		"WedgeFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D7),
		"Total Wedge Tray Water-Equivalent Thickness",
		"TotalWedgeTrayWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D8),
		"Wedge Orientation",
		"WedgeOrientation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D9),
		"Isocenter to Wedge Tray Distance",
		"IsocenterToWedgeTrayDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DA),
		"Source to Wedge Tray Distance",
		"SourceToWedgeTrayDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DB),
		"Wedge Thin Edge Position",
		"WedgeThinEdgePosition",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DC),
		"Bolus ID",
		"BolusID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DD),
		"Bolus Description",
		"BolusDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DE),
		"Effective Wedge Angle",
		"EffectiveWedgeAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E0),
		"Number of Compensators",
		"NumberOfCompensators",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E1),
		"Material ID",
		"MaterialID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E2),
		"Total Compensator Tray Factor",
		"TotalCompensatorTrayFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E3),
		"Compensator Sequence",
		"CompensatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E4),
		"Compensator Number",
		"CompensatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E5),
		"Compensator ID",
		"CompensatorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E6),
		"Source to Compensator Tray Distance",
		"SourceToCompensatorTrayDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E7),
		"Compensator Rows",
		"CompensatorRows",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E8),
		"Compensator Columns",
		"CompensatorColumns",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E9),
		"Compensator Pixel Spacing",
		"CompensatorPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EA),
		"Compensator Position",
		"CompensatorPosition",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EB),
		"Compensator Transmission Data",
		"CompensatorTransmissionData",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EC),
		"Compensator Thickness Data",
		"CompensatorThicknessData",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00ED),
		"Number of Boli",
		"NumberOfBoli",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EE),
		"Compensator Type",
		"CompensatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EF),
		"Compensator Tray ID",
		"CompensatorTrayID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F0),
		"Number of Blocks",
		"NumberOfBlocks",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F2),
		"Total Block Tray Factor",
		"TotalBlockTrayFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F3),
		"Total Block Tray Water-Equivalent Thickness",
		"TotalBlockTrayWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F4),
		"Block Sequence",
		"BlockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F5),
		"Block Tray ID",
		"BlockTrayID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F6),
		"Source to Block Tray Distance",
		"SourceToBlockTrayDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F7),
		"Isocenter to Block Tray Distance",
		"IsocenterToBlockTrayDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F8),
		"Block Type",
		"BlockType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F9),
		"Accessory Code",
		"AccessoryCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00FA),
		"Block Divergence",
		"BlockDivergence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00FB),
		"Block Mounting Position",
		"BlockMountingPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00FC),
		"Block Number",
		"BlockNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00FE),
		"Block Name",
		"BlockName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0100),
		"Block Thickness",
		"BlockThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0102),
		"Block Transmission",
		"BlockTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0104),
		"Block Number of Points",
		"BlockNumberOfPoints",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0106),
		"Block Data",
		"BlockData",
		vm.VM22N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0107),
		"Applicator Sequence",
		"ApplicatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0108),
		"Applicator ID",
		"ApplicatorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0109),
		"Applicator Type",
		"ApplicatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x010A),
		"Applicator Description",
		"ApplicatorDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x010C),
		"Cumulative Dose Reference Coefficient",
		"CumulativeDoseReferenceCoefficient",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x010E),
		"Final Cumulative Meterset Weight",
		"FinalCumulativeMetersetWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0110),
		"Number of Control Points",
		"NumberOfControlPoints",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0111),
		"Control Point Sequence",
		"ControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0112),
		"Control Point Index",
		"ControlPointIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0114),
		"Nominal Beam Energy",
		"NominalBeamEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0115),
		"Dose Rate Set",
		"DoseRateSet",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0116),
		"Wedge Position Sequence",
		"WedgePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0118),
		"Wedge Position",
		"WedgePosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x011A),
		"Beam Limiting Device Position Sequence",
		"BeamLimitingDevicePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x011C),
		"Leaf/Jaw Positions",
		"LeafJawPositions",
		vm.VM22N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x011E),
		"Gantry Angle",
		"GantryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x011F),
		"Gantry Rotation Direction",
		"GantryRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0120),
		"Beam Limiting Device Angle",
		"BeamLimitingDeviceAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0121),
		"Beam Limiting Device Rotation Direction",
		"BeamLimitingDeviceRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0122),
		"Patient Support Angle",
		"PatientSupportAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0123),
		"Patient Support Rotation Direction",
		"PatientSupportRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0124),
		"Table Top Eccentric Axis Distance",
		"TableTopEccentricAxisDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0125),
		"Table Top Eccentric Angle",
		"TableTopEccentricAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0126),
		"Table Top Eccentric Rotation Direction",
		"TableTopEccentricRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0128),
		"Table Top Vertical Position",
		"TableTopVerticalPosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0129),
		"Table Top Longitudinal Position",
		"TableTopLongitudinalPosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x012A),
		"Table Top Lateral Position",
		"TableTopLateralPosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x012C),
		"Isocenter Position",
		"IsocenterPosition",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x012E),
		"Surface Entry Point",
		"SurfaceEntryPoint",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0130),
		"Source to Surface Distance",
		"SourceToSurfaceDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0131),
		"Average Beam Dose Point Source to External Contour Distance",
		"AverageBeamDosePointSourceToExternalContourDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0132),
		"Source to External Contour Distance",
		"SourceToExternalContourDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0133),
		"External Contour Entry Point",
		"ExternalContourEntryPoint",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0134),
		"Cumulative Meterset Weight",
		"CumulativeMetersetWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0140),
		"Table Top Pitch Angle",
		"TableTopPitchAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0142),
		"Table Top Pitch Rotation Direction",
		"TableTopPitchRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0144),
		"Table Top Roll Angle",
		"TableTopRollAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0146),
		"Table Top Roll Rotation Direction",
		"TableTopRollRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0148),
		"Head Fixation Angle",
		"HeadFixationAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x014A),
		"Gantry Pitch Angle",
		"GantryPitchAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x014C),
		"Gantry Pitch Rotation Direction",
		"GantryPitchRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x014E),
		"Gantry Pitch Angle Tolerance",
		"GantryPitchAngleTolerance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0150),
		"Fixation Eye",
		"FixationEye",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0151),
		"Chair Head Frame Position",
		"ChairHeadFramePosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0152),
		"Head Fixation Angle Tolerance",
		"HeadFixationAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0153),
		"Chair Head Frame Position Tolerance",
		"ChairHeadFramePositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0154),
		"Fixation Light Azimuthal Angle Tolerance",
		"FixationLightAzimuthalAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0155),
		"Fixation Light Polar Angle Tolerance",
		"FixationLightPolarAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0180),
		"Patient Setup Sequence",
		"PatientSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0182),
		"Patient Setup Number",
		"PatientSetupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0183),
		"Patient Setup Label",
		"PatientSetupLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0184),
		"Patient Additional Position",
		"PatientAdditionalPosition",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0190),
		"Fixation Device Sequence",
		"FixationDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0192),
		"Fixation Device Type",
		"FixationDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0194),
		"Fixation Device Label",
		"FixationDeviceLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0196),
		"Fixation Device Description",
		"FixationDeviceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0198),
		"Fixation Device Position",
		"FixationDevicePosition",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0199),
		"Fixation Device Pitch Angle",
		"FixationDevicePitchAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x019A),
		"Fixation Device Roll Angle",
		"FixationDeviceRollAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A0),
		"Shielding Device Sequence",
		"ShieldingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A2),
		"Shielding Device Type",
		"ShieldingDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A4),
		"Shielding Device Label",
		"ShieldingDeviceLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A6),
		"Shielding Device Description",
		"ShieldingDeviceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A8),
		"Shielding Device Position",
		"ShieldingDevicePosition",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B0),
		"Setup Technique",
		"SetupTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B2),
		"Setup Technique Description",
		"SetupTechniqueDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B4),
		"Setup Device Sequence",
		"SetupDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B6),
		"Setup Device Type",
		"SetupDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B8),
		"Setup Device Label",
		"SetupDeviceLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01BA),
		"Setup Device Description",
		"SetupDeviceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01BC),
		"Setup Device Parameter",
		"SetupDeviceParameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01D0),
		"Setup Reference Description",
		"SetupReferenceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01D2),
		"Table Top Vertical Setup Displacement",
		"TableTopVerticalSetupDisplacement",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01D4),
		"Table Top Longitudinal Setup Displacement",
		"TableTopLongitudinalSetupDisplacement",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01D6),
		"Table Top Lateral Setup Displacement",
		"TableTopLateralSetupDisplacement",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0200),
		"Brachy Treatment Technique",
		"BrachyTreatmentTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0202),
		"Brachy Treatment Type",
		"BrachyTreatmentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0206),
		"Treatment Machine Sequence",
		"TreatmentMachineSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0210),
		"Source Sequence",
		"SourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0212),
		"Source Number",
		"SourceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0214),
		"Source Type",
		"SourceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0216),
		"Source Manufacturer",
		"SourceManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0218),
		"Active Source Diameter",
		"ActiveSourceDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x021A),
		"Active Source Length",
		"ActiveSourceLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x021B),
		"Source Model ID",
		"SourceModelID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x021C),
		"Source Description",
		"SourceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0222),
		"Source Encapsulation Nominal Thickness",
		"SourceEncapsulationNominalThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0224),
		"Source Encapsulation Nominal Transmission",
		"SourceEncapsulationNominalTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0226),
		"Source Isotope Name",
		"SourceIsotopeName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0228),
		"Source Isotope Half Life",
		"SourceIsotopeHalfLife",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0229),
		"Source Strength Units",
		"SourceStrengthUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x022A),
		"Reference Air Kerma Rate",
		"ReferenceAirKermaRate",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x022B),
		"Source Strength",
		"SourceStrength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x022C),
		"Source Strength Reference Date",
		"SourceStrengthReferenceDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x022E),
		"Source Strength Reference Time",
		"SourceStrengthReferenceTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0230),
		"Application Setup Sequence",
		"ApplicationSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0232),
		"Application Setup Type",
		"ApplicationSetupType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0234),
		"Application Setup Number",
		"ApplicationSetupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0236),
		"Application Setup Name",
		"ApplicationSetupName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0238),
		"Application Setup Manufacturer",
		"ApplicationSetupManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0240),
		"Template Number",
		"TemplateNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0242),
		"Template Type",
		"TemplateType",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0244),
		"Template Name",
		"TemplateName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0250),
		"Total Reference Air Kerma",
		"TotalReferenceAirKerma",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0260),
		"Brachy Accessory Device Sequence",
		"BrachyAccessoryDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0262),
		"Brachy Accessory Device Number",
		"BrachyAccessoryDeviceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0263),
		"Brachy Accessory Device ID",
		"BrachyAccessoryDeviceID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0264),
		"Brachy Accessory Device Type",
		"BrachyAccessoryDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0266),
		"Brachy Accessory Device Name",
		"BrachyAccessoryDeviceName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x026A),
		"Brachy Accessory Device Nominal Thickness",
		"BrachyAccessoryDeviceNominalThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x026C),
		"Brachy Accessory Device Nominal Transmission",
		"BrachyAccessoryDeviceNominalTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0271),
		"Channel Effective Length",
		"ChannelEffectiveLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0272),
		"Channel Inner Length",
		"ChannelInnerLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0273),
		"Afterloader Channel ID",
		"AfterloaderChannelID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0274),
		"Source Applicator Tip Length",
		"SourceApplicatorTipLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0280),
		"Channel Sequence",
		"ChannelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0282),
		"Channel Number",
		"ChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0284),
		"Channel Length",
		"ChannelLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0286),
		"Channel Total Time",
		"ChannelTotalTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0288),
		"Source Movement Type",
		"SourceMovementType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x028A),
		"Number of Pulses",
		"NumberOfPulses",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x028C),
		"Pulse Repetition Interval",
		"PulseRepetitionInterval",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0290),
		"Source Applicator Number",
		"SourceApplicatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0291),
		"Source Applicator ID",
		"SourceApplicatorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0292),
		"Source Applicator Type",
		"SourceApplicatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0294),
		"Source Applicator Name",
		"SourceApplicatorName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0296),
		"Source Applicator Length",
		"SourceApplicatorLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0298),
		"Source Applicator Manufacturer",
		"SourceApplicatorManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x029C),
		"Source Applicator Wall Nominal Thickness",
		"SourceApplicatorWallNominalThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x029E),
		"Source Applicator Wall Nominal Transmission",
		"SourceApplicatorWallNominalTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02A0),
		"Source Applicator Step Size",
		"SourceApplicatorStepSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02A1),
		"Applicator Shape Referenced ROI Number",
		"ApplicatorShapeReferencedROINumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02A2),
		"Transfer Tube Number",
		"TransferTubeNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02A4),
		"Transfer Tube Length",
		"TransferTubeLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B0),
		"Channel Shield Sequence",
		"ChannelShieldSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B2),
		"Channel Shield Number",
		"ChannelShieldNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B3),
		"Channel Shield ID",
		"ChannelShieldID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B4),
		"Channel Shield Name",
		"ChannelShieldName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B8),
		"Channel Shield Nominal Thickness",
		"ChannelShieldNominalThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02BA),
		"Channel Shield Nominal Transmission",
		"ChannelShieldNominalTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02C8),
		"Final Cumulative Time Weight",
		"FinalCumulativeTimeWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02D0),
		"Brachy Control Point Sequence",
		"BrachyControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02D2),
		"Control Point Relative Position",
		"ControlPointRelativePosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02D4),
		"Control Point 3D Position",
		"ControlPoint3DPosition",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02D6),
		"Cumulative Time Weight",
		"CumulativeTimeWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E0),
		"Compensator Divergence",
		"CompensatorDivergence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E1),
		"Compensator Mounting Position",
		"CompensatorMountingPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E2),
		"Source to Compensator Distance",
		"SourceToCompensatorDistance",
		vm.VM1N,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E3),
		"Total Compensator Tray Water-Equivalent Thickness",
		"TotalCompensatorTrayWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E4),
		"Isocenter to Compensator Tray Distance",
		"IsocenterToCompensatorTrayDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E5),
		"Compensator Column Offset",
		"CompensatorColumnOffset",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E6),
		"Isocenter to Compensator Distances",
		"IsocenterToCompensatorDistances",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E7),
		"Compensator Relative Stopping Power Ratio",
		"CompensatorRelativeStoppingPowerRatio",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E8),
		"Compensator Milling Tool Diameter",
		"CompensatorMillingToolDiameter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02EA),
		"Ion Range Compensator Sequence",
		"IonRangeCompensatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02EB),
		"Compensator Description",
		"CompensatorDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02EC),
		"Compensator Surface Representation Flag",
		"CompensatorSurfaceRepresentationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0302),
		"Radiation Mass Number",
		"RadiationMassNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0304),
		"Radiation Atomic Number",
		"RadiationAtomicNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0306),
		"Radiation Charge State",
		"RadiationChargeState",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0308),
		"Scan Mode",
		"ScanMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0309),
		"Modulated Scan Mode Type",
		"ModulatedScanModeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x030A),
		"Virtual Source-Axis Distances",
		"VirtualSourceAxisDistances",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x030C),
		"Snout Sequence",
		"SnoutSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x030D),
		"Snout Position",
		"SnoutPosition",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x030F),
		"Snout ID",
		"SnoutID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0312),
		"Number of Range Shifters",
		"NumberOfRangeShifters",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0314),
		"Range Shifter Sequence",
		"RangeShifterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0316),
		"Range Shifter Number",
		"RangeShifterNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0318),
		"Range Shifter ID",
		"RangeShifterID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0320),
		"Range Shifter Type",
		"RangeShifterType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0322),
		"Range Shifter Description",
		"RangeShifterDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0330),
		"Number of Lateral Spreading Devices",
		"NumberOfLateralSpreadingDevices",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0332),
		"Lateral Spreading Device Sequence",
		"LateralSpreadingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0334),
		"Lateral Spreading Device Number",
		"LateralSpreadingDeviceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0336),
		"Lateral Spreading Device ID",
		"LateralSpreadingDeviceID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0338),
		"Lateral Spreading Device Type",
		"LateralSpreadingDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x033A),
		"Lateral Spreading Device Description",
		"LateralSpreadingDeviceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x033C),
		"Lateral Spreading Device Water Equivalent Thickness",
		"LateralSpreadingDeviceWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0340),
		"Number of Range Modulators",
		"NumberOfRangeModulators",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0342),
		"Range Modulator Sequence",
		"RangeModulatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0344),
		"Range Modulator Number",
		"RangeModulatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0346),
		"Range Modulator ID",
		"RangeModulatorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0348),
		"Range Modulator Type",
		"RangeModulatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x034A),
		"Range Modulator Description",
		"RangeModulatorDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x034C),
		"Beam Current Modulation ID",
		"BeamCurrentModulationID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0350),
		"Patient Support Type",
		"PatientSupportType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0352),
		"Patient Support ID",
		"PatientSupportID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0354),
		"Patient Support Accessory Code",
		"PatientSupportAccessoryCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0355),
		"Tray Accessory Code",
		"TrayAccessoryCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0356),
		"Fixation Light Azimuthal Angle",
		"FixationLightAzimuthalAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0358),
		"Fixation Light Polar Angle",
		"FixationLightPolarAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x035A),
		"Meterset Rate",
		"MetersetRate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0360),
		"Range Shifter Settings Sequence",
		"RangeShifterSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0362),
		"Range Shifter Setting",
		"RangeShifterSetting",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0364),
		"Isocenter to Range Shifter Distance",
		"IsocenterToRangeShifterDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0366),
		"Range Shifter Water Equivalent Thickness",
		"RangeShifterWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0370),
		"Lateral Spreading Device Settings Sequence",
		"LateralSpreadingDeviceSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0372),
		"Lateral Spreading Device Setting",
		"LateralSpreadingDeviceSetting",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0374),
		"Isocenter to Lateral Spreading Device Distance",
		"IsocenterToLateralSpreadingDeviceDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0380),
		"Range Modulator Settings Sequence",
		"RangeModulatorSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0382),
		"Range Modulator Gating Start Value",
		"RangeModulatorGatingStartValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0384),
		"Range Modulator Gating Stop Value",
		"RangeModulatorGatingStopValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0386),
		"Range Modulator Gating Start Water Equivalent Thickness",
		"RangeModulatorGatingStartWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0388),
		"Range Modulator Gating Stop Water Equivalent Thickness",
		"RangeModulatorGatingStopWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x038A),
		"Isocenter to Range Modulator Distance",
		"IsocenterToRangeModulatorDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x038F),
		"Scan Spot Time Offset",
		"ScanSpotTimeOffset",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0390),
		"Scan Spot Tune ID",
		"ScanSpotTuneID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0391),
		"Scan Spot Prescribed Indices",
		"ScanSpotPrescribedIndices",
		vm.VM1N,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0392),
		"Number of Scan Spot Positions",
		"NumberOfScanSpotPositions",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0393),
		"Scan Spot Reordered",
		"ScanSpotReordered",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0394),
		"Scan Spot Position Map",
		"ScanSpotPositionMap",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0395),
		"Scan Spot Reordering Allowed",
		"ScanSpotReorderingAllowed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0396),
		"Scan Spot Meterset Weights",
		"ScanSpotMetersetWeights",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0398),
		"Scanning Spot Size",
		"ScanningSpotSize",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0399),
		"Scan Spot Sizes Delivered",
		"ScanSpotSizesDelivered",
		vm.VM22N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x039A),
		"Number of Paintings",
		"NumberOfPaintings",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x039B),
		"Scan Spot Gantry Angles",
		"ScanSpotGantryAngles",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x039C),
		"Scan Spot Patient Support Angles",
		"ScanSpotPatientSupportAngles",
		vm.VM1N,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A0),
		"Ion Tolerance Table Sequence",
		"IonToleranceTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A2),
		"Ion Beam Sequence",
		"IonBeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A4),
		"Ion Beam Limiting Device Sequence",
		"IonBeamLimitingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A6),
		"Ion Block Sequence",
		"IonBlockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A8),
		"Ion Control Point Sequence",
		"IonControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03AA),
		"Ion Wedge Sequence",
		"IonWedgeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03AC),
		"Ion Wedge Position Sequence",
		"IonWedgePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0401),
		"Referenced Setup Image Sequence",
		"ReferencedSetupImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0402),
		"Setup Image Comment",
		"SetupImageComment",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0410),
		"Motion Synchronization Sequence",
		"MotionSynchronizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0412),
		"Control Point Orientation",
		"ControlPointOrientation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0420),
		"General Accessory Sequence",
		"GeneralAccessorySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0421),
		"General Accessory ID",
		"GeneralAccessoryID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0422),
		"General Accessory Description",
		"GeneralAccessoryDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0423),
		"General Accessory Type",
		"GeneralAccessoryType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0424),
		"General Accessory Number",
		"GeneralAccessoryNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0425),
		"Source to General Accessory Distance",
		"SourceToGeneralAccessoryDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0426),
		"Isocenter to General Accessory Distance",
		"IsocenterToGeneralAccessoryDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0431),
		"Applicator Geometry Sequence",
		"ApplicatorGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0432),
		"Applicator Aperture Shape",
		"ApplicatorApertureShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0433),
		"Applicator Opening",
		"ApplicatorOpening",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0434),
		"Applicator Opening X",
		"ApplicatorOpeningX",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0435),
		"Applicator Opening Y",
		"ApplicatorOpeningY",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0436),
		"Source to Applicator Mounting Position Distance",
		"SourceToApplicatorMountingPositionDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0440),
		"Number of Block Slab Items",
		"NumberOfBlockSlabItems",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0441),
		"Block Slab Sequence",
		"BlockSlabSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0442),
		"Block Slab Thickness",
		"BlockSlabThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0443),
		"Block Slab Number",
		"BlockSlabNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0450),
		"Device Motion Control Sequence",
		"DeviceMotionControlSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0451),
		"Device Motion Execution Mode",
		"DeviceMotionExecutionMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0452),
		"Device Motion Observation Mode",
		"DeviceMotionObservationMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0453),
		"Device Motion Parameter Code Sequence",
		"DeviceMotionParameterCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0501),
		"Distal Depth Fraction",
		"DistalDepthFraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0502),
		"Distal Depth",
		"DistalDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0503),
		"Nominal Range Modulation Fractions",
		"NominalRangeModulationFractions",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0504),
		"Nominal Range Modulated Region Depths",
		"NominalRangeModulatedRegionDepths",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0505),
		"Depth Dose Parameters Sequence",
		"DepthDoseParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0506),
		"Delivered Depth Dose Parameters Sequence",
		"DeliveredDepthDoseParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0507),
		"Delivered Distal Depth Fraction",
		"DeliveredDistalDepthFraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0508),
		"Delivered Distal Depth",
		"DeliveredDistalDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0509),
		"Delivered Nominal Range Modulation Fractions",
		"DeliveredNominalRangeModulationFractions",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0510),
		"Delivered Nominal Range Modulated Region Depths",
		"DeliveredNominalRangeModulatedRegionDepths",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0511),
		"Delivered Reference Dose Definition",
		"DeliveredReferenceDoseDefinition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0512),
		"Reference Dose Definition",
		"ReferenceDoseDefinition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0600),
		"RT Control Point Index",
		"RTControlPointIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0601),
		"Radiation Generation Mode Index",
		"RadiationGenerationModeIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0602),
		"Referenced Defined Device Index",
		"ReferencedDefinedDeviceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0603),
		"Radiation Dose Identification Index",
		"RadiationDoseIdentificationIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0604),
		"Number of RT Control Points",
		"NumberOfRTControlPoints",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0605),
		"Referenced Radiation Generation Mode Index",
		"ReferencedRadiationGenerationModeIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0606),
		"Treatment Position Index",
		"TreatmentPositionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0607),
		"Referenced Device Index",
		"ReferencedDeviceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0608),
		"Treatment Position Group Label",
		"TreatmentPositionGroupLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0609),
		"Treatment Position Group UID",
		"TreatmentPositionGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060A),
		"Treatment Position Group Sequence",
		"TreatmentPositionGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060B),
		"Referenced Treatment Position Index",
		"ReferencedTreatmentPositionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060C),
		"Referenced Radiation Dose Identification Index",
		"ReferencedRadiationDoseIdentificationIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060D),
		"RT Accessory Holder Water-Equivalent Thickness",
		"RTAccessoryHolderWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060E),
		"Referenced RT Accessory Holder Device Index",
		"ReferencedRTAccessoryHolderDeviceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060F),
		"RT Accessory Holder Slot Existence Flag",
		"RTAccessoryHolderSlotExistenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0610),
		"RT Accessory Holder Slot Sequence",
		"RTAccessoryHolderSlotSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0611),
		"RT Accessory Holder Slot ID",
		"RTAccessoryHolderSlotID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0612),
		"RT Accessory Holder Slot Distance",
		"RTAccessoryHolderSlotDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0613),
		"RT Accessory Slot Distance",
		"RTAccessorySlotDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0614),
		"RT Accessory Holder Definition Sequence",
		"RTAccessoryHolderDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0615),
		"RT Accessory Device Slot ID",
		"RTAccessoryDeviceSlotID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0616),
		"RT Radiation Sequence",
		"RTRadiationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0617),
		"Radiation Dose Sequence",
		"RadiationDoseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0618),
		"Radiation Dose Identification Sequence",
		"RadiationDoseIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0619),
		"Radiation Dose Identification Label",
		"RadiationDoseIdentificationLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061A),
		"Reference Dose Type",
		"ReferenceDoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061B),
		"Primary Dose Value Indicator",
		"PrimaryDoseValueIndicator",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061C),
		"Dose Values Sequence",
		"DoseValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061D),
		"Dose Value Purpose",
		"DoseValuePurpose",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061E),
		"Reference Dose Point Coordinates",
		"ReferenceDosePointCoordinates",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061F),
		"Radiation Dose Values Parameters Sequence",
		"RadiationDoseValuesParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0620),
		"Meterset to Dose Mapping Sequence",
		"MetersetToDoseMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0621),
		"Expected In-Vivo Measurement Values Sequence",
		"ExpectedInVivoMeasurementValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0622),
		"Expected In-Vivo Measurement Value Index",
		"ExpectedInVivoMeasurementValueIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0623),
		"Radiation Dose In-Vivo Measurement Label",
		"RadiationDoseInVivoMeasurementLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0624),
		"Radiation Dose Central Axis Displacement",
		"RadiationDoseCentralAxisDisplacement",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0625),
		"Radiation Dose Value",
		"RadiationDoseValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0626),
		"Radiation Dose Source to Skin Distance",
		"RadiationDoseSourceToSkinDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0627),
		"Radiation Dose Measurement Point Coordinates",
		"RadiationDoseMeasurementPointCoordinates",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0628),
		"Radiation Dose Source to External Contour Distance",
		"RadiationDoseSourceToExternalContourDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0629),
		"RT Tolerance Set Sequence",
		"RTToleranceSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062A),
		"RT Tolerance Set Label",
		"RTToleranceSetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062B),
		"Attribute Tolerance Values Sequence",
		"AttributeToleranceValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062C),
		"Tolerance Value",
		"ToleranceValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062D),
		"Patient Support Position Tolerance Sequence",
		"PatientSupportPositionToleranceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062E),
		"Treatment Time Limit",
		"TreatmentTimeLimit",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062F),
		"C-Arm Photon-Electron Control Point Sequence",
		"CArmPhotonElectronControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0630),
		"Referenced RT Radiation Sequence",
		"ReferencedRTRadiationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0631),
		"Referenced RT Instance Sequence",
		"ReferencedRTInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0632),
		"Referenced RT Patient Setup Sequence",
		"ReferencedRTPatientSetupSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0634),
		"Source to Patient Surface Distance",
		"SourceToPatientSurfaceDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0635),
		"Treatment Machine Special Mode Code Sequence",
		"TreatmentMachineSpecialModeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0636),
		"Intended Number of Fractions",
		"IntendedNumberOfFractions",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0637),
		"RT Radiation Set Intent",
		"RTRadiationSetIntent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0638),
		"RT Radiation Physical and Geometric Content Detail Flag",
		"RTRadiationPhysicalAndGeometricContentDetailFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0639),
		"RT Record Flag",
		"RTRecordFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063A),
		"Treatment Device Identification Sequence",
		"TreatmentDeviceIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063B),
		"Referenced RT Physician Intent Sequence",
		"ReferencedRTPhysicianIntentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063C),
		"Cumulative Meterset",
		"CumulativeMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063D),
		"Delivery Rate",
		"DeliveryRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063E),
		"Delivery Rate Unit Sequence",
		"DeliveryRateUnitSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063F),
		"Treatment Position Sequence",
		"TreatmentPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0640),
		"Radiation Source-Axis Distance",
		"RadiationSourceAxisDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0641),
		"Number of RT Beam Limiting Devices",
		"NumberOfRTBeamLimitingDevices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0642),
		"RT Beam Limiting Device Proximal Distance",
		"RTBeamLimitingDeviceProximalDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0643),
		"RT Beam Limiting Device Distal Distance",
		"RTBeamLimitingDeviceDistalDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0644),
		"Parallel RT Beam Delimiter Device Orientation Label Code Sequence",
		"ParallelRTBeamDelimiterDeviceOrientationLabelCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0645),
		"Beam Modifier Orientation Angle",
		"BeamModifierOrientationAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0646),
		"Fixed RT Beam Delimiter Device Sequence",
		"FixedRTBeamDelimiterDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0647),
		"Parallel RT Beam Delimiter Device Sequence",
		"ParallelRTBeamDelimiterDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0648),
		"Number of Parallel RT Beam Delimiters",
		"NumberOfParallelRTBeamDelimiters",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0649),
		"Parallel RT Beam Delimiter Boundaries",
		"ParallelRTBeamDelimiterBoundaries",
		vm.VM2N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064A),
		"Parallel RT Beam Delimiter Positions",
		"ParallelRTBeamDelimiterPositions",
		vm.VM2N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064B),
		"RT Beam Limiting Device Offset",
		"RTBeamLimitingDeviceOffset",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064C),
		"RT Beam Delimiter Geometry Sequence",
		"RTBeamDelimiterGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064D),
		"RT Beam Limiting Device Definition Sequence",
		"RTBeamLimitingDeviceDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064E),
		"Parallel RT Beam Delimiter Opening Mode",
		"ParallelRTBeamDelimiterOpeningMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064F),
		"Parallel RT Beam Delimiter Leaf Mounting Side",
		"ParallelRTBeamDelimiterLeafMountingSide",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0650),
		"Patient Setup UID",
		"PatientSetupUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0651),
		"Wedge Definition Sequence",
		"WedgeDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0652),
		"Radiation Beam Wedge Angle",
		"RadiationBeamWedgeAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0653),
		"Radiation Beam Wedge Thin Edge Distance",
		"RadiationBeamWedgeThinEdgeDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0654),
		"Radiation Beam Effective Wedge Angle",
		"RadiationBeamEffectiveWedgeAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0655),
		"Number of Wedge Positions",
		"NumberOfWedgePositions",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0656),
		"RT Beam Limiting Device Opening Sequence",
		"RTBeamLimitingDeviceOpeningSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0657),
		"Number of RT Beam Limiting Device Openings",
		"NumberOfRTBeamLimitingDeviceOpenings",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0658),
		"Radiation Dosimeter Unit Sequence",
		"RadiationDosimeterUnitSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0659),
		"RT Device Distance Reference Location Code Sequence",
		"RTDeviceDistanceReferenceLocationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065A),
		"Radiation Device Configuration and Commissioning Key Sequence",
		"RadiationDeviceConfigurationAndCommissioningKeySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065B),
		"Patient Support Position Parameter Sequence",
		"PatientSupportPositionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065C),
		"Patient Support Position Specification Method",
		"PatientSupportPositionSpecificationMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065D),
		"Patient Support Position Device Parameter Sequence",
		"PatientSupportPositionDeviceParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065E),
		"Device Order Index",
		"DeviceOrderIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065F),
		"Patient Support Position Parameter Order Index",
		"PatientSupportPositionParameterOrderIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0660),
		"Patient Support Position Device Tolerance Sequence",
		"PatientSupportPositionDeviceToleranceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0661),
		"Patient Support Position Tolerance Order Index",
		"PatientSupportPositionToleranceOrderIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0662),
		"Compensator Definition Sequence",
		"CompensatorDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0663),
		"Compensator Map Orientation",
		"CompensatorMapOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0664),
		"Compensator Proximal Thickness Map",
		"CompensatorProximalThicknessMap",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0665),
		"Compensator Distal Thickness Map",
		"CompensatorDistalThicknessMap",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0666),
		"Compensator Base Plane Offset",
		"CompensatorBasePlaneOffset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0667),
		"Compensator Shape Fabrication Code Sequence",
		"CompensatorShapeFabricationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0668),
		"Compensator Shape Sequence",
		"CompensatorShapeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0669),
		"Radiation Beam Compensator Milling Tool Diameter",
		"RadiationBeamCompensatorMillingToolDiameter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066A),
		"Block Definition Sequence",
		"BlockDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066B),
		"Block Edge Data",
		"BlockEdgeData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066C),
		"Block Orientation",
		"BlockOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066D),
		"Radiation Beam Block Thickness",
		"RadiationBeamBlockThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066E),
		"Radiation Beam Block Slab Thickness",
		"RadiationBeamBlockSlabThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066F),
		"Block Edge Data Sequence",
		"BlockEdgeDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0670),
		"Number of RT Accessory Holders",
		"NumberOfRTAccessoryHolders",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0671),
		"General Accessory Definition Sequence",
		"GeneralAccessoryDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0672),
		"Number of General Accessories",
		"NumberOfGeneralAccessories",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0673),
		"Bolus Definition Sequence",
		"BolusDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0674),
		"Number of Boluses",
		"NumberOfBoluses",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0675),
		"Equipment Frame of Reference UID",
		"EquipmentFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0676),
		"Equipment Frame of Reference Description",
		"EquipmentFrameOfReferenceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0677),
		"Equipment Reference Point Coordinates Sequence",
		"EquipmentReferencePointCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0678),
		"Equipment Reference Point Code Sequence",
		"EquipmentReferencePointCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0679),
		"RT Beam Limiting Device Angle",
		"RTBeamLimitingDeviceAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067A),
		"Source Roll Angle",
		"SourceRollAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067B),
		"Radiation GenerationMode Sequence",
		"RadiationGenerationModeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067C),
		"Radiation GenerationMode Label",
		"RadiationGenerationModeLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067D),
		"Radiation GenerationMode Description",
		"RadiationGenerationModeDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067E),
		"Radiation GenerationMode Machine Code Sequence",
		"RadiationGenerationModeMachineCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067F),
		"Radiation Type Code Sequence",
		"RadiationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0680),
		"Nominal Energy",
		"NominalEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0681),
		"Minimum Nominal Energy",
		"MinimumNominalEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0682),
		"Maximum Nominal Energy",
		"MaximumNominalEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0683),
		"Radiation Fluence Modifier Code Sequence",
		"RadiationFluenceModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0684),
		"Energy Unit Code Sequence",
		"EnergyUnitCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0685),
		"Number of Radiation GenerationModes",
		"NumberOfRadiationGenerationModes",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0686),
		"Patient Support Devices Sequence",
		"PatientSupportDevicesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0687),
		"Number of Patient Support Devices",
		"NumberOfPatientSupportDevices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0688),
		"RT Beam Modifier Definition Distance",
		"RTBeamModifierDefinitionDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0689),
		"Beam Area Limit Sequence",
		"BeamAreaLimitSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x068A),
		"Referenced RT Prescription Sequence",
		"ReferencedRTPrescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x068B),
		"Dose Value Interpretation",
		"DoseValueInterpretation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0700),
		"Treatment Session UID",
		"TreatmentSessionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0701),
		"RT Radiation Usage",
		"RTRadiationUsage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0702),
		"Referenced RT Radiation Set Sequence",
		"ReferencedRTRadiationSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0703),
		"Referenced RT Radiation Record Sequence",
		"ReferencedRTRadiationRecordSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0704),
		"RT Radiation Set Delivery Number",
		"RTRadiationSetDeliveryNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0705),
		"Clinical Fraction Number",
		"ClinicalFractionNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0706),
		"RT Treatment Fraction Completion Status",
		"RTTreatmentFractionCompletionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0707),
		"RT Radiation Set Usage",
		"RTRadiationSetUsage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0708),
		"Treatment Delivery Continuation Flag",
		"TreatmentDeliveryContinuationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0709),
		"Treatment Record Content Origin",
		"TreatmentRecordContentOrigin",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0714),
		"RT Treatment Termination Status",
		"RTTreatmentTerminationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0715),
		"RT Treatment Termination Reason Code Sequence",
		"RTTreatmentTerminationReasonCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0716),
		"Machine-Specific Treatment Termination Code Sequence",
		"MachineSpecificTreatmentTerminationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0722),
		"RT Radiation Salvage Record Control Point Sequence",
		"RTRadiationSalvageRecordControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0723),
		"Starting Meterset Value Known Flag",
		"StartingMetersetValueKnownFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0730),
		"Treatment Termination Description",
		"TreatmentTerminationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0731),
		"Treatment Tolerance Violation Sequence",
		"TreatmentToleranceViolationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0732),
		"Treatment Tolerance Violation Category",
		"TreatmentToleranceViolationCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0733),
		"Treatment Tolerance Violation Attribute Sequence",
		"TreatmentToleranceViolationAttributeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0734),
		"Treatment Tolerance Violation Description",
		"TreatmentToleranceViolationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0735),
		"Treatment Tolerance Violation Identification",
		"TreatmentToleranceViolationIdentification",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0736),
		"Treatment Tolerance Violation DateTime",
		"TreatmentToleranceViolationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x073A),
		"Recorded RT Control Point DateTime",
		"RecordedRTControlPointDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x073B),
		"Referenced Radiation RT Control Point Index",
		"ReferencedRadiationRTControlPointIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x073E),
		"Alternate Value Sequence",
		"AlternateValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x073F),
		"Confirmation Sequence",
		"ConfirmationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0740),
		"Interlock Sequence",
		"InterlockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0741),
		"Interlock DateTime",
		"InterlockDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0742),
		"Interlock Description",
		"InterlockDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0743),
		"Interlock Originating Device Sequence",
		"InterlockOriginatingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0744),
		"Interlock Code Sequence",
		"InterlockCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0745),
		"Interlock Resolution Code Sequence",
		"InterlockResolutionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0746),
		"Interlock Resolution User Sequence",
		"InterlockResolutionUserSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0760),
		"Override DateTime",
		"OverrideDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0761),
		"Treatment Tolerance Violation Type Code Sequence",
		"TreatmentToleranceViolationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0762),
		"Treatment Tolerance Violation Cause Code Sequence",
		"TreatmentToleranceViolationCauseCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0772),
		"Measured Meterset to Dose Mapping Sequence",
		"MeasuredMetersetToDoseMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0773),
		"Referenced Expected In-Vivo Measurement Value Index",
		"ReferencedExpectedInVivoMeasurementValueIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0774),
		"Dose Measurement Device Code Sequence",
		"DoseMeasurementDeviceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0780),
		"Additional Parameter Recording Instance Sequence",
		"AdditionalParameterRecordingInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0783),
		"Interlock Origin Description",
		"InterlockOriginDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0784),
		"RT Patient Position Scope Sequence",
		"RTPatientPositionScopeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0785),
		"Referenced Treatment Position Group UID",
		"ReferencedTreatmentPositionGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0786),
		"Radiation Order Index",
		"RadiationOrderIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0787),
		"Omitted Radiation Sequence",
		"OmittedRadiationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0788),
		"Reason for Omission Code Sequence",
		"ReasonForOmissionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0789),
		"RT Delivery Start Patient Position Sequence",
		"RTDeliveryStartPatientPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078A),
		"RT Treatment Preparation Patient Position Sequence",
		"RTTreatmentPreparationPatientPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078B),
		"Referenced RT Treatment Preparation Sequence",
		"ReferencedRTTreatmentPreparationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078C),
		"Referenced Patient Setup Photo Sequence",
		"ReferencedPatientSetupPhotoSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078D),
		"Patient Treatment Preparation Method Code Sequence",
		"PatientTreatmentPreparationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078E),
		"Patient Treatment Preparation Procedure Parameter Description",
		"PatientTreatmentPreparationProcedureParameterDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078F),
		"Patient Treatment Preparation Device Sequence",
		"PatientTreatmentPreparationDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0790),
		"Patient Treatment Preparation Procedure Sequence",
		"PatientTreatmentPreparationProcedureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0791),
		"Patient Treatment Preparation Procedure Code Sequence",
		"PatientTreatmentPreparationProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0792),
		"Patient Treatment Preparation Method Description",
		"PatientTreatmentPreparationMethodDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0793),
		"Patient Treatment Preparation Procedure Parameter Sequence",
		"PatientTreatmentPreparationProcedureParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0794),
		"Patient Setup Photo Description",
		"PatientSetupPhotoDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0795),
		"Patient Treatment Preparation Procedure Index",
		"PatientTreatmentPreparationProcedureIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0796),
		"Referenced Patient Setup Procedure Index",
		"ReferencedPatientSetupProcedureIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0797),
		"RT Radiation Task Sequence",
		"RTRadiationTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0798),
		"RT Patient Position Displacement Sequence",
		"RTPatientPositionDisplacementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0799),
		"RT Patient Position Sequence",
		"RTPatientPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079A),
		"Displacement Reference Label",
		"DisplacementReferenceLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079B),
		"Displacement Matrix",
		"DisplacementMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079C),
		"Patient Support Displacement Sequence",
		"PatientSupportDisplacementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079D),
		"Displacement Reference Location Code Sequence",
		"DisplacementReferenceLocationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079E),
		"RT Radiation Set Delivery Usage",
		"RTRadiationSetDeliveryUsage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079F),
		"Patient Treatment Preparation Sequence",
		"PatientTreatmentPreparationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x07A0),
		"Patient to Equipment Relationship Sequence",
		"PatientToEquipmentRelationshipSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x07A1),
		"Imaging Equipment to Treatment Delivery Device Relationship Sequence",
		"ImagingEquipmentToTreatmentDeliveryDeviceRelationshipSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0002),
		"Referenced RT Plan Sequence",
		"ReferencedRTPlanSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0004),
		"Referenced Beam Sequence",
		"ReferencedBeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0006),
		"Referenced Beam Number",
		"ReferencedBeamNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0007),
		"Referenced Reference Image Number",
		"ReferencedReferenceImageNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0008),
		"Start Cumulative Meterset Weight",
		"StartCumulativeMetersetWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0009),
		"End Cumulative Meterset Weight",
		"EndCumulativeMetersetWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x000A),
		"Referenced Brachy Application Setup Sequence",
		"ReferencedBrachyApplicationSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x000C),
		"Referenced Brachy Application Setup Number",
		"ReferencedBrachyApplicationSetupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x000E),
		"Referenced Source Number",
		"ReferencedSourceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0020),
		"Referenced Fraction Group Sequence",
		"ReferencedFractionGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0022),
		"Referenced Fraction Group Number",
		"ReferencedFractionGroupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0040),
		"Referenced Verification Image Sequence",
		"ReferencedVerificationImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0042),
		"Referenced Reference Image Sequence",
		"ReferencedReferenceImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0050),
		"Referenced Dose Reference Sequence",
		"ReferencedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0051),
		"Referenced Dose Reference Number",
		"ReferencedDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0055),
		"Brachy Referenced Dose Reference Sequence",
		"BrachyReferencedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0060),
		"Referenced Structure Set Sequence",
		"ReferencedStructureSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x006A),
		"Referenced Patient Setup Number",
		"ReferencedPatientSetupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0080),
		"Referenced Dose Sequence",
		"ReferencedDoseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00A0),
		"Referenced Tolerance Table Number",
		"ReferencedToleranceTableNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00B0),
		"Referenced Bolus Sequence",
		"ReferencedBolusSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00C0),
		"Referenced Wedge Number",
		"ReferencedWedgeNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00D0),
		"Referenced Compensator Number",
		"ReferencedCompensatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00E0),
		"Referenced Block Number",
		"ReferencedBlockNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00F0),
		"Referenced Control Point Index",
		"ReferencedControlPointIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00F2),
		"Referenced Control Point Sequence",
		"ReferencedControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00F4),
		"Referenced Start Control Point Index",
		"ReferencedStartControlPointIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00F6),
		"Referenced Stop Control Point Index",
		"ReferencedStopControlPointIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0100),
		"Referenced Range Shifter Number",
		"ReferencedRangeShifterNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0102),
		"Referenced Lateral Spreading Device Number",
		"ReferencedLateralSpreadingDeviceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0104),
		"Referenced Range Modulator Number",
		"ReferencedRangeModulatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0111),
		"Omitted Beam Task Sequence",
		"OmittedBeamTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0112),
		"Reason for Omission",
		"ReasonForOmission",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0113),
		"Reason for Omission Description",
		"ReasonForOmissionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0114),
		"Prescription Overview Sequence",
		"PrescriptionOverviewSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0115),
		"Total Prescription Dose",
		"TotalPrescriptionDose",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0116),
		"Plan Overview Sequence",
		"PlanOverviewSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0117),
		"Plan Overview Index",
		"PlanOverviewIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0118),
		"Referenced Plan Overview Index",
		"ReferencedPlanOverviewIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0119),
		"Number of Fractions Included",
		"NumberOfFractionsIncluded",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0120),
		"Dose Calibration Conditions Sequence",
		"DoseCalibrationConditionsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0121),
		"Absorbed Dose to Meterset Ratio",
		"AbsorbedDoseToMetersetRatio",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0122),
		"Delineated Radiation Field Size",
		"DelineatedRadiationFieldSize",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0123),
		"Dose Calibration Conditions Verified Flag",
		"DoseCalibrationConditionsVerifiedFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0124),
		"Calibration Reference Point Depth",
		"CalibrationReferencePointDepth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0125),
		"Gating Beam Hold Transition Sequence",
		"GatingBeamHoldTransitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0126),
		"Beam Hold Transition",
		"BeamHoldTransition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0127),
		"Beam Hold Transition DateTime",
		"BeamHoldTransitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0128),
		"Beam Hold Originating Device Sequence",
		"BeamHoldOriginatingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0129),
		"Beam Hold Transition Trigger Source",
		"BeamHoldTransitionTriggerSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300E, 0x0002),
		"Approval Status",
		"ApprovalStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300E, 0x0004),
		"Review Date",
		"ReviewDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x300E, 0x0005),
		"Review Time",
		"ReviewTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x300E, 0x0008),
		"Reviewer Name",
		"ReviewerName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0001),
		"Radiobiological Dose Effect Sequence",
		"RadiobiologicalDoseEffectSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0002),
		"Radiobiological Dose Effect Flag",
		"RadiobiologicalDoseEffectFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0003),
		"Effective Dose Calculation Method Category Code Sequence",
		"EffectiveDoseCalculationMethodCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0004),
		"Effective Dose Calculation Method Code Sequence",
		"EffectiveDoseCalculationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0005),
		"Effective Dose Calculation Method Description",
		"EffectiveDoseCalculationMethodDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0006),
		"Conceptual Volume UID",
		"ConceptualVolumeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0007),
		"Originating SOP Instance Reference Sequence",
		"OriginatingSOPInstanceReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0008),
		"Conceptual Volume Constituent Sequence",
		"ConceptualVolumeConstituentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0009),
		"Equivalent Conceptual Volume Instance Reference Sequence",
		"EquivalentConceptualVolumeInstanceReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000A),
		"Equivalent Conceptual Volumes Sequence",
		"EquivalentConceptualVolumesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000B),
		"Referenced Conceptual Volume UID",
		"ReferencedConceptualVolumeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000C),
		"Conceptual Volume Combination Expression",
		"ConceptualVolumeCombinationExpression",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000D),
		"Conceptual Volume Constituent Index",
		"ConceptualVolumeConstituentIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000E),
		"Conceptual Volume Combination Flag",
		"ConceptualVolumeCombinationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000F),
		"Conceptual Volume Combination Description",
		"ConceptualVolumeCombinationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0010),
		"Conceptual Volume Segmentation Defined Flag",
		"ConceptualVolumeSegmentationDefinedFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0011),
		"Conceptual Volume Segmentation Reference Sequence",
		"ConceptualVolumeSegmentationReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0012),
		"Conceptual Volume Constituent Segmentation Reference Sequence",
		"ConceptualVolumeConstituentSegmentationReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0013),
		"Constituent Conceptual Volume UID",
		"ConstituentConceptualVolumeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0014),
		"Derivation Conceptual Volume Sequence",
		"DerivationConceptualVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0015),
		"Source Conceptual Volume UID",
		"SourceConceptualVolumeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0016),
		"Conceptual Volume Derivation Algorithm Sequence",
		"ConceptualVolumeDerivationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0017),
		"Conceptual Volume Description",
		"ConceptualVolumeDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0018),
		"Source Conceptual Volume Sequence",
		"SourceConceptualVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0019),
		"Author Identification Sequence",
		"AuthorIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001A),
		"Manufacturer's Model Version",
		"ManufacturerModelVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001B),
		"Device Alternate Identifier",
		"DeviceAlternateIdentifier",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001C),
		"Device Alternate Identifier Type",
		"DeviceAlternateIdentifierType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001D),
		"Device Alternate Identifier Format",
		"DeviceAlternateIdentifierFormat",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001E),
		"Segmentation Creation Template Label",
		"SegmentationCreationTemplateLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001F),
		"Segmentation Template UID",
		"SegmentationTemplateUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0020),
		"Referenced Segment Reference Index",
		"ReferencedSegmentReferenceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0021),
		"Segment Reference Sequence",
		"SegmentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0022),
		"Segment Reference Index",
		"SegmentReferenceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0023),
		"Direct Segment Reference Sequence",
		"DirectSegmentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0024),
		"Combination Segment Reference Sequence",
		"CombinationSegmentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0025),
		"Conceptual Volume Sequence",
		"ConceptualVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0026),
		"Segmented RT Accessory Device Sequence",
		"SegmentedRTAccessoryDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0027),
		"Segment Characteristics Sequence",
		"SegmentCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0028),
		"Related Segment Characteristics Sequence",
		"RelatedSegmentCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0029),
		"Segment Characteristics Precedence",
		"SegmentCharacteristicsPrecedence",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002A),
		"RT Segment Annotation Sequence",
		"RTSegmentAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002B),
		"Segment Annotation Category Code Sequence",
		"SegmentAnnotationCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002C),
		"Segment Annotation Type Code Sequence",
		"SegmentAnnotationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002D),
		"Device Label",
		"DeviceLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002E),
		"Device Type Code Sequence",
		"DeviceTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002F),
		"Segment Annotation Type Modifier Code Sequence",
		"SegmentAnnotationTypeModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0030),
		"Patient Equipment Relationship Code Sequence",
		"PatientEquipmentRelationshipCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0031),
		"Referenced Fiducials UID",
		"ReferencedFiducialsUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0032),
		"Patient Treatment Orientation Sequence",
		"PatientTreatmentOrientationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0033),
		"User Content Label",
		"UserContentLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0034),
		"User Content Long Label",
		"UserContentLongLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0035),
		"Entity Label",
		"EntityLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0036),
		"Entity Name",
		"EntityName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0037),
		"Entity Description",
		"EntityDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0038),
		"Entity Long Label",
		"EntityLongLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0039),
		"Device Index",
		"DeviceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003A),
		"RT Treatment Phase Index",
		"RTTreatmentPhaseIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003B),
		"RT Treatment Phase UID",
		"RTTreatmentPhaseUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003C),
		"RT Prescription Index",
		"RTPrescriptionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003D),
		"RT Segment Annotation Index",
		"RTSegmentAnnotationIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003E),
		"Basis RT Treatment Phase Index",
		"BasisRTTreatmentPhaseIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003F),
		"Related RT Treatment Phase Index",
		"RelatedRTTreatmentPhaseIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0040),
		"Referenced RT Treatment Phase Index",
		"ReferencedRTTreatmentPhaseIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0041),
		"Referenced RT Prescription Index",
		"ReferencedRTPrescriptionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0042),
		"Referenced Parent RT Prescription Index",
		"ReferencedParentRTPrescriptionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0043),
		"Manufacturer's Device Identifier",
		"ManufacturerDeviceIdentifier",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0044),
		"Instance-Level Referenced Performed Procedure Step Sequence",
		"InstanceLevelReferencedPerformedProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0045),
		"RT Treatment Phase Intent Presence Flag",
		"RTTreatmentPhaseIntentPresenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0046),
		"Radiotherapy Treatment Type",
		"RadiotherapyTreatmentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0047),
		"Teletherapy Radiation Type",
		"TeletherapyRadiationType",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0048),
		"Brachytherapy Source Type",
		"BrachytherapySourceType",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0049),
		"Referenced RT Treatment Phase Sequence",
		"ReferencedRTTreatmentPhaseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004A),
		"Referenced Direct Segment Instance Sequence",
		"ReferencedDirectSegmentInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004B),
		"Intended RT Treatment Phase Sequence",
		"IntendedRTTreatmentPhaseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004C),
		"Intended Phase Start Date",
		"IntendedPhaseStartDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004D),
		"Intended Phase End Date",
		"IntendedPhaseEndDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004E),
		"RT Treatment Phase Interval Sequence",
		"RTTreatmentPhaseIntervalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004F),
		"Temporal Relationship Interval Anchor",
		"TemporalRelationshipIntervalAnchor",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0050),
		"Minimum Number of Interval Days",
		"MinimumNumberOfIntervalDays",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0051),
		"Maximum Number of Interval Days",
		"MaximumNumberOfIntervalDays",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0052),
		"Pertinent SOP Classes in Study",
		"PertinentSOPClassesInStudy",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0053),
		"Pertinent SOP Classes in Series",
		"PertinentSOPClassesInSeries",
		vm.VM1N,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0054),
		"RT Prescription Label",
		"RTPrescriptionLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0055),
		"RT Physician Intent Predecessor Sequence",
		"RTPhysicianIntentPredecessorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0056),
		"RT Treatment Approach Label",
		"RTTreatmentApproachLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0057),
		"RT Physician Intent Sequence",
		"RTPhysicianIntentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0058),
		"RT Physician Intent Index",
		"RTPhysicianIntentIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0059),
		"RT Treatment Intent Type",
		"RTTreatmentIntentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005A),
		"RT Physician Intent Narrative",
		"RTPhysicianIntentNarrative",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005B),
		"RT Protocol Code Sequence",
		"RTProtocolCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005C),
		"Reason for Superseding",
		"ReasonForSuperseding",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005D),
		"RT Diagnosis Code Sequence",
		"RTDiagnosisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005E),
		"Referenced RT Physician Intent Index",
		"ReferencedRTPhysicianIntentIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005F),
		"RT Physician Intent Input Instance Sequence",
		"RTPhysicianIntentInputInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0060),
		"RT Anatomic Prescription Sequence",
		"RTAnatomicPrescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0061),
		"Prior Treatment Dose Description",
		"PriorTreatmentDoseDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0062),
		"Prior Treatment Reference Sequence",
		"PriorTreatmentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0063),
		"Dosimetric Objective Evaluation Scope",
		"DosimetricObjectiveEvaluationScope",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0064),
		"Therapeutic Role Category Code Sequence",
		"TherapeuticRoleCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0065),
		"Therapeutic Role Type Code Sequence",
		"TherapeuticRoleTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0066),
		"Conceptual Volume Optimization Precedence",
		"ConceptualVolumeOptimizationPrecedence",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0067),
		"Conceptual Volume Category Code Sequence",
		"ConceptualVolumeCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0068),
		"Conceptual Volume Blocking Constraint",
		"ConceptualVolumeBlockingConstraint",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0069),
		"Conceptual Volume Type Code Sequence",
		"ConceptualVolumeTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006A),
		"Conceptual Volume Type Modifier Code Sequence",
		"ConceptualVolumeTypeModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006B),
		"RT Prescription Sequence",
		"RTPrescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006C),
		"Dosimetric Objective Sequence",
		"DosimetricObjectiveSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006D),
		"Dosimetric Objective Type Code Sequence",
		"DosimetricObjectiveTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006E),
		"Dosimetric Objective UID",
		"DosimetricObjectiveUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006F),
		"Referenced Dosimetric Objective UID",
		"ReferencedDosimetricObjectiveUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0070),
		"Dosimetric Objective Parameter Sequence",
		"DosimetricObjectiveParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0071),
		"Referenced Dosimetric Objectives Sequence",
		"ReferencedDosimetricObjectivesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0073),
		"Absolute Dosimetric Objective Flag",
		"AbsoluteDosimetricObjectiveFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0074),
		"Dosimetric Objective Weight",
		"DosimetricObjectiveWeight",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0075),
		"Dosimetric Objective Purpose",
		"DosimetricObjectivePurpose",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0076),
		"Planning Input Information Sequence",
		"PlanningInputInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0077),
		"Treatment Site",
		"TreatmentSite",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0078),
		"Treatment Site Code Sequence",
		"TreatmentSiteCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0079),
		"Fraction Pattern Sequence",
		"FractionPatternSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007A),
		"Treatment Technique Notes",
		"TreatmentTechniqueNotes",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007B),
		"Prescription Notes",
		"PrescriptionNotes",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007C),
		"Number of Interval Fractions",
		"NumberOfIntervalFractions",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007D),
		"Number of Fractions",
		"NumberOfFractions",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007E),
		"Intended Delivery Duration",
		"IntendedDeliveryDuration",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007F),
		"Fractionation Notes",
		"FractionationNotes",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0080),
		"RT Treatment Technique Code Sequence",
		"RTTreatmentTechniqueCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0081),
		"Prescription Notes Sequence",
		"PrescriptionNotesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0082),
		"Fraction-Based Relationship Sequence",
		"FractionBasedRelationshipSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0083),
		"Fraction-Based Relationship Interval Anchor",
		"FractionBasedRelationshipIntervalAnchor",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0084),
		"Minimum Hours between Fractions",
		"MinimumHoursBetweenFractions",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0085),
		"Intended Fraction Start Time",
		"IntendedFractionStartTime",
		vm.VM1N,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0086),
		"Intended Start Day of Week",
		"IntendedStartDayOfWeek",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0087),
		"Weekday Fraction Pattern Sequence",
		"WeekdayFractionPatternSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0088),
		"Delivery Time Structure Code Sequence",
		"DeliveryTimeStructureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0089),
		"Treatment Site Modifier Code Sequence",
		"TreatmentSiteModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0090),
		"Robotic Base Location Indicator",
		"RoboticBaseLocationIndicator",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0091),
		"Robotic Path Node Set Code Sequence",
		"RoboticPathNodeSetCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0092),
		"Robotic Node Identifier",
		"RoboticNodeIdentifier",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0093),
		"RT Treatment Source Coordinates",
		"RTTreatmentSourceCoordinates",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0094),
		"Radiation Source Coordinate SystemYaw Angle",
		"RadiationSourceCoordinateSystemYawAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0095),
		"Radiation Source Coordinate SystemRoll Angle",
		"RadiationSourceCoordinateSystemRollAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0096),
		"Radiation Source Coordinate System Pitch Angle",
		"RadiationSourceCoordinateSystemPitchAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0097),
		"Robotic Path Control Point Sequence",
		"RoboticPathControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0098),
		"Tomotherapeutic Control Point Sequence",
		"TomotherapeuticControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0099),
		"Tomotherapeutic Leaf Open Durations",
		"TomotherapeuticLeafOpenDurations",
		vm.VM1N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x009A),
		"Tomotherapeutic Leaf Initial Closed Durations",
		"TomotherapeuticLeafInitialClosedDurations",
		vm.VM1N,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x00A0),
		"Conceptual Volume Identification Sequence",
		"ConceptualVolumeIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4000, 0x0010),
		"Arbitrary",
		"Arbitrary",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4000, 0x4000),
		"Text Comments",
		"TextComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0040),
		"Results ID",
		"ResultsID",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0042),
		"Results ID Issuer",
		"ResultsIDIssuer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0050),
		"Referenced Interpretation Sequence",
		"ReferencedInterpretationSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x00FF),
		"Report Production Status (Trial)",
		"ReportProductionStatusTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0100),
		"Interpretation Recorded Date",
		"InterpretationRecordedDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0101),
		"Interpretation Recorded Time",
		"InterpretationRecordedTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0102),
		"Interpretation Recorder",
		"InterpretationRecorder",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0103),
		"Reference to Recorded Sound",
		"ReferenceToRecordedSound",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0108),
		"Interpretation Transcription Date",
		"InterpretationTranscriptionDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0109),
		"Interpretation Transcription Time",
		"InterpretationTranscriptionTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x010A),
		"Interpretation Transcriber",
		"InterpretationTranscriber",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x010B),
		"Interpretation Text",
		"InterpretationText",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x010C),
		"Interpretation Author",
		"InterpretationAuthor",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0111),
		"Interpretation Approver Sequence",
		"InterpretationApproverSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0112),
		"Interpretation Approval Date",
		"InterpretationApprovalDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0113),
		"Interpretation Approval Time",
		"InterpretationApprovalTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0114),
		"Physician Approving Interpretation",
		"PhysicianApprovingInterpretation",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0115),
		"Interpretation Diagnosis Description",
		"InterpretationDiagnosisDescription",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0117),
		"Interpretation Diagnosis Code Sequence",
		"InterpretationDiagnosisCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0118),
		"Results Distribution List Sequence",
		"ResultsDistributionListSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0119),
		"Distribution Name",
		"DistributionName",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x011A),
		"Distribution Address",
		"DistributionAddress",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0200),
		"Interpretation ID",
		"InterpretationID",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0202),
		"Interpretation ID Issuer",
		"InterpretationIDIssuer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0210),
		"Interpretation Type ID",
		"InterpretationTypeID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0212),
		"Interpretation Status ID",
		"InterpretationStatusID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0300),
		"Impressions",
		"Impressions",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x4000),
		"Results Comments",
		"ResultsComments",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x0001),
		"Low Energy Detectors",
		"LowEnergyDetectors",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x0002),
		"High Energy Detectors",
		"HighEnergyDetectors",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x0004),
		"Detector Geometry Sequence",
		"DetectorGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1001),
		"Threat ROI Voxel Sequence",
		"ThreatROIVoxelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1004),
		"Threat ROI Base",
		"ThreatROIBase",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1005),
		"Threat ROI Extents",
		"ThreatROIExtents",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1006),
		"Threat ROI Bitmap",
		"ThreatROIBitmap",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1007),
		"Route Segment ID",
		"RouteSegmentID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1008),
		"Gantry Type",
		"GantryType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1009),
		"OOI Owner Type",
		"OOIOwnerType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x100A),
		"Route Segment Sequence",
		"RouteSegmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1010),
		"Potential Threat Object ID",
		"PotentialThreatObjectID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1011),
		"Threat Sequence",
		"ThreatSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1012),
		"Threat Category",
		"ThreatCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1013),
		"Threat Category Description",
		"ThreatCategoryDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1014),
		"ATD Ability Assessment",
		"ATDAbilityAssessment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1015),
		"ATD Assessment Flag",
		"ATDAssessmentFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1016),
		"ATD Assessment Probability",
		"ATDAssessmentProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1017),
		"Mass",
		"Mass",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1018),
		"Density",
		"Density",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1019),
		"Z Effective",
		"ZEffective",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101A),
		"Boarding Pass ID",
		"BoardingPassID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101B),
		"Center of Mass",
		"CenterOfMass",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101C),
		"Center of PTO",
		"CenterOfPTO",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101D),
		"Bounding Polygon",
		"BoundingPolygon",
		vm.MustParse("6-n"),
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101E),
		"Route Segment Start Location ID",
		"RouteSegmentStartLocationID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101F),
		"Route Segment End Location ID",
		"RouteSegmentEndLocationID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1020),
		"Route Segment Location ID Type",
		"RouteSegmentLocationIDType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1021),
		"Abort Reason",
		"AbortReason",
		vm.VM1N,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1023),
		"Volume of PTO",
		"VolumeOfPTO",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1024),
		"Abort Flag",
		"AbortFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1025),
		"Route Segment Start Time",
		"RouteSegmentStartTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1026),
		"Route Segment End Time",
		"RouteSegmentEndTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1027),
		"TDR Type",
		"TDRType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1028),
		"International Route Segment",
		"InternationalRouteSegment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1029),
		"Threat Detection Algorithm and Version",
		"ThreatDetectionAlgorithmAndVersion",
		vm.VM1N,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x102A),
		"Assigned Location",
		"AssignedLocation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x102B),
		"Alarm Decision Time",
		"AlarmDecisionTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1031),
		"Alarm Decision",
		"AlarmDecision",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1033),
		"Number of Total Objects",
		"NumberOfTotalObjects",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1034),
		"Number of Alarm Objects",
		"NumberOfAlarmObjects",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1037),
		"PTO Representation Sequence",
		"PTORepresentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1038),
		"ATD Assessment Sequence",
		"ATDAssessmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1039),
		"TIP Type",
		"TIPType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x103A),
		"DICOS Version",
		"DICOSVersion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1041),
		"OOI Owner Creation Time",
		"OOIOwnerCreationTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1042),
		"OOI Type",
		"OOIType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1043),
		"OOI Size",
		"OOISize",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1044),
		"Acquisition Status",
		"AcquisitionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1045),
		"Basis Materials Code Sequence",
		"BasisMaterialsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1046),
		"Phantom Type",
		"PhantomType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1047),
		"OOI Owner Sequence",
		"OOIOwnerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1048),
		"Scan Type",
		"ScanType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1051),
		"Itinerary ID",
		"ItineraryID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1052),
		"Itinerary ID Type",
		"ItineraryIDType",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1053),
		"Itinerary ID Assigning Authority",
		"ItineraryIDAssigningAuthority",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1054),
		"Route ID",
		"RouteID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1055),
		"Route ID Assigning Authority",
		"RouteIDAssigningAuthority",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1056),
		"Inbound Arrival Type",
		"InboundArrivalType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1058),
		"Carrier ID",
		"CarrierID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1059),
		"Carrier ID Assigning Authority",
		"CarrierIDAssigningAuthority",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1060),
		"Source Orientation",
		"SourceOrientation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1061),
		"Source Position",
		"SourcePosition",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1062),
		"Belt Height",
		"BeltHeight",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1064),
		"Algorithm Routing Code Sequence",
		"AlgorithmRoutingCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1067),
		"Transport Classification",
		"TransportClassification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1068),
		"OOI Type Descriptor",
		"OOITypeDescriptor",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1069),
		"Total Processing Time",
		"TotalProcessingTime",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x106C),
		"Detector Calibration Data",
		"DetectorCalibrationData",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x106D),
		"Additional Screening Performed",
		"AdditionalScreeningPerformed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x106E),
		"Additional Inspection Selection Criteria",
		"AdditionalInspectionSelectionCriteria",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x106F),
		"Additional Inspection Method Sequence",
		"AdditionalInspectionMethodSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1070),
		"AIT Device Type",
		"AITDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1071),
		"QR Measurements Sequence",
		"QRMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1072),
		"Target Material Sequence",
		"TargetMaterialSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1073),
		"SNR Threshold",
		"SNRThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1075),
		"Image Scale Representation",
		"ImageScaleRepresentation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1076),
		"Referenced PTO Sequence",
		"ReferencedPTOSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1077),
		"Referenced TDR Instance Sequence",
		"ReferencedTDRInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1078),
		"PTO Location Description",
		"PTOLocationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1079),
		"Anomaly Locator Indicator Sequence",
		"AnomalyLocatorIndicatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107A),
		"Anomaly Locator Indicator",
		"AnomalyLocatorIndicator",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107B),
		"PTO Region Sequence",
		"PTORegionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107C),
		"Inspection Selection Criteria",
		"InspectionSelectionCriteria",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107D),
		"Secondary Inspection Method Sequence",
		"SecondaryInspectionMethodSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107E),
		"PRCS to RCS Orientation",
		"PRCSToRCSOrientation",
		vm.VM6,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x4FFE, 0x0001),
		"MAC Parameters Sequence",
		"MACParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0005)"),
		"Curve Dimensions",
		"CurveDimensions",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0010)"),
		"Number of Points",
		"NumberOfPoints",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0020)"),
		"Type of Data",
		"TypeOfData",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0022)"),
		"Curve Description",
		"CurveDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0030)"),
		"Axis Units",
		"AxisUnits",
		vm.VM1N,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0040)"),
		"Axis Labels",
		"AxisLabels",
		vm.VM1N,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0103)"),
		"Data Value Representation",
		"DataValueRepresentation",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0104)"),
		"Minimum Coordinate Value",
		"MinimumCoordinateValue",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0105)"),
		"Maximum Coordinate Value",
		"MaximumCoordinateValue",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0106)"),
		"Curve Range",
		"CurveRange",
		vm.VM1N,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0110)"),
		"Curve Data Descriptor",
		"CurveDataDescriptor",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0112)"),
		"Coordinate Start Value",
		"CoordinateStartValue",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0114)"),
		"Coordinate Step Value",
		"CoordinateStepValue",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,1001)"),
		"Curve Activation Layer",
		"CurveActivationLayer",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2000)"),
		"Audio Type",
		"AudioType",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2002)"),
		"Audio Sample Format",
		"AudioSampleFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2004)"),
		"Number of Channels",
		"NumberOfChannels",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2006)"),
		"Number of Samples",
		"NumberOfSamples",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2008)"),
		"Sample Rate",
		"SampleRate",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,200A)"),
		"Total Time",
		"TotalTime",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,200C)"),
		"Audio Sample Data",
		"AudioSampleData",
		vm.VM1,
		true,
		vr.OB, vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,200E)"),
		"Audio Comments",
		"AudioComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2500)"),
		"Curve Label",
		"CurveLabel",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2600)"),
		"Curve Referenced Overlay Sequence",
		"CurveReferencedOverlaySequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2610)"),
		"Curve Referenced Overlay Group",
		"CurveReferencedOverlayGroup",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,3000)"),
		"Curve Data",
		"CurveData",
		vm.VM1,
		true,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5200, 0x9229),
		"Shared Functional Groups Sequence",
		"SharedFunctionalGroupsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x5200, 0x9230),
		"Per-Frame Functional Groups Sequence",
		"PerFrameFunctionalGroupsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x0100),
		"Waveform Sequence",
		"WaveformSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x0110),
		"Channel Minimum Value",
		"ChannelMinimumValue",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x0112),
		"Channel Maximum Value",
		"ChannelMaximumValue",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x1004),
		"Waveform Bits Allocated",
		"WaveformBitsAllocated",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x1006),
		"Waveform Sample Interpretation",
		"WaveformSampleInterpretation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x100A),
		"Waveform Padding Value",
		"WaveformPaddingValue",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x1010),
		"Waveform Data",
		"WaveformData",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5600, 0x0010),
		"First Order Phase Correction Angle",
		"FirstOrderPhaseCorrectionAngle",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x5600, 0x0020),
		"Spectroscopy Data",
		"SpectroscopyData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0010)"),
		"Overlay Rows",
		"OverlayRows",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0011)"),
		"Overlay Columns",
		"OverlayColumns",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0012)"),
		"Overlay Planes",
		"OverlayPlanes",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0015)"),
		"Number of Frames in Overlay",
		"NumberOfFramesInOverlay",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0022)"),
		"Overlay Description",
		"OverlayDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0040)"),
		"Overlay Type",
		"OverlayType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0045)"),
		"Overlay Subtype",
		"OverlaySubtype",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0050)"),
		"Overlay Origin",
		"OverlayOrigin",
		vm.VM2,
		false,
		vr.SS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0051)"),
		"Image Frame Origin",
		"ImageFrameOrigin",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0052)"),
		"Overlay Plane Origin",
		"OverlayPlaneOrigin",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0060)"),
		"Overlay Compression Code",
		"OverlayCompressionCode",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0061)"),
		"Overlay Compression Originator",
		"OverlayCompressionOriginator",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0062)"),
		"Overlay Compression Label",
		"OverlayCompressionLabel",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0063)"),
		"Overlay Compression Description",
		"OverlayCompressionDescription",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0066)"),
		"Overlay Compression Step Pointers",
		"OverlayCompressionStepPointers",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0068)"),
		"Overlay Repeat Interval",
		"OverlayRepeatInterval",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0069)"),
		"Overlay Bits Grouped",
		"OverlayBitsGrouped",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0100)"),
		"Overlay Bits Allocated",
		"OverlayBitsAllocated",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0102)"),
		"Overlay Bit Position",
		"OverlayBitPosition",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0110)"),
		"Overlay Format",
		"OverlayFormat",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0200)"),
		"Overlay Location",
		"OverlayLocation",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0800)"),
		"Overlay Code Label",
		"OverlayCodeLabel",
		vm.VM1N,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0802)"),
		"Overlay Number of Tables",
		"OverlayNumberOfTables",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0803)"),
		"Overlay Code Table Location",
		"OverlayCodeTableLocation",
		vm.VM1N,
		true,
		vr.AT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0804)"),
		"Overlay Bits For Code Word",
		"OverlayBitsForCodeWord",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1001)"),
		"Overlay Activation Layer",
		"OverlayActivationLayer",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1100)"),
		"Overlay Descriptor - Gray",
		"OverlayDescriptorGray",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1101)"),
		"Overlay Descriptor - Red",
		"OverlayDescriptorRed",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1102)"),
		"Overlay Descriptor - Green",
		"OverlayDescriptorGreen",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1103)"),
		"Overlay Descriptor - Blue",
		"OverlayDescriptorBlue",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1200)"),
		"Overlays - Gray",
		"OverlaysGray",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1201)"),
		"Overlays - Red",
		"OverlaysRed",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1202)"),
		"Overlays - Green",
		"OverlaysGreen",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1203)"),
		"Overlays - Blue",
		"OverlaysBlue",
		vm.VM1N,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1301)"),
		"ROI Area",
		"ROIArea",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1302)"),
		"ROI Mean",
		"ROIMean",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1303)"),
		"ROI Standard Deviation",
		"ROIStandardDeviation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1500)"),
		"Overlay Label",
		"OverlayLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,3000)"),
		"Overlay Data",
		"OverlayData",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,4000)"),
		"Overlay Comments",
		"OverlayComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0001),
		"Extended Offset Table",
		"ExtendedOffsetTable",
		vm.VM1,
		false,
		vr.OV,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0002),
		"Extended Offset Table Lengths",
		"ExtendedOffsetTableLengths",
		vm.VM1,
		false,
		vr.OV,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0003),
		"Encapsulated Pixel Data Value Total Length",
		"EncapsulatedPixelDataValueTotalLength",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0008),
		"Float Pixel Data",
		"FloatPixelData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0009),
		"Double Float Pixel Data",
		"DoubleFloatPixelData",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0010),
		"Pixel Data",
		"PixelData",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0020),
		"Coefficients SDVN",
		"CoefficientsSDVN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0030),
		"Coefficients SDHN",
		"CoefficientsSDHN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0040),
		"Coefficients SDDN",
		"CoefficientsSDDN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0010)"),
		"Variable Pixel Data",
		"VariablePixelData",
		vm.VM1,
		true,
		vr.OB, vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0011)"),
		"Variable Next Data Group",
		"VariableNextDataGroup",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0020)"),
		"Variable Coefficients SDVN",
		"VariableCoefficientsSDVN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0030)"),
		"Variable Coefficients SDHN",
		"VariableCoefficientsSDHN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0040)"),
		"Variable Coefficients SDDN",
		"VariableCoefficientsSDDN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0xFFFA, 0xFFFA),
		"Digital Signatures Sequence",
		"DigitalSignaturesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0xFFFC, 0xFFFC),
		"Data Set Trailing Padding",
		"DataSetTrailingPadding",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0xFFFE, 0xE000),
		"Item",
		"Item",
		vm.VM1,
		false,
		vr.None,
	))
	d.Add(NewEntry(
		tag.New(0xFFFE, 0xE00D),
		"Item Delimitation Item",
		"ItemDelimitationItem",
		vm.VM1,
		false,
		vr.None,
	))
	d.Add(NewEntry(
		tag.New(0xFFFE, 0xE0DD),
		"Sequence Delimitation Item",
		"SequenceDelimitationItem",
		vm.VM1,
		false,
		vr.None,
	))
	d.Add(NewEntry(
		tag.New(0x0006, 0x0001),
		"Current Frame Functional Groups Sequence",
		"CurrentFrameFunctionalGroupsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
}
