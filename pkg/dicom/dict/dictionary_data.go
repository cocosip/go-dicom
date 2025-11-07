// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DICOM Dictionary.xml. DO NOT EDIT.

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
		"Command Group Length", // CommandGroupLength
		"CommandGroupLength",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0001),
		"Command Length to End", // CommandLengthToEnd
		"CommandLengthToEnd",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0002),
		"Affected SOP Class UID", // AffectedSOPClassUID
		"AffectedSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0003),
		"Requested SOP Class UID", // RequestedSOPClassUID
		"RequestedSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0010),
		"Command Recognition Code", // CommandRecognitionCode
		"CommandRecognitionCode",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0100),
		"Command Field", // CommandField
		"CommandField",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0110),
		"Message ID", // MessageID
		"MessageID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0120),
		"Message ID Being Responded To", // MessageIDBeingRespondedTo
		"MessageIDBeingRespondedTo",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0200),
		"Initiator", // Initiator
		"Initiator",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0300),
		"Receiver", // Receiver
		"Receiver",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0400),
		"Find Location", // FindLocation
		"FindLocation",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0600),
		"Move Destination", // MoveDestination
		"MoveDestination",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0700),
		"Priority", // Priority
		"Priority",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0800),
		"Command Data Set Type", // CommandDataSetType
		"CommandDataSetType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0850),
		"Number of Matches", // NumberOfMatches
		"NumberOfMatches",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0860),
		"Response Sequence Number", // ResponseSequenceNumber
		"ResponseSequenceNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0900),
		"Status", // Status
		"Status",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0901),
		"Offending Element", // OffendingElement
		"OffendingElement",
		vm.VM1_n,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0902),
		"Error Comment", // ErrorComment
		"ErrorComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x0903),
		"Error ID", // ErrorID
		"ErrorID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1000),
		"Affected SOP Instance UID", // AffectedSOPInstanceUID
		"AffectedSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1001),
		"Requested SOP Instance UID", // RequestedSOPInstanceUID
		"RequestedSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1002),
		"Event Type ID", // EventTypeID
		"EventTypeID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1005),
		"Attribute Identifier List", // AttributeIdentifierList
		"AttributeIdentifierList",
		vm.VM1_n,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1008),
		"Action Type ID", // ActionTypeID
		"ActionTypeID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1020),
		"Number of Remaining Sub-operations", // NumberOfRemainingSuboperations
		"NumberOfRemainingSuboperations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1021),
		"Number of Completed Sub-operations", // NumberOfCompletedSuboperations
		"NumberOfCompletedSuboperations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1022),
		"Number of Failed Sub-operations", // NumberOfFailedSuboperations
		"NumberOfFailedSuboperations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1023),
		"Number of Warning Sub-operations", // NumberOfWarningSuboperations
		"NumberOfWarningSuboperations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1030),
		"Move Originator Application Entity Title", // MoveOriginatorApplicationEntityTitle
		"MoveOriginatorApplicationEntityTitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x1031),
		"Move Originator Message ID", // MoveOriginatorMessageID
		"MoveOriginatorMessageID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x4000),
		"Dialog Receiver", // DialogReceiver
		"DialogReceiver",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x4010),
		"Terminal Type", // TerminalType
		"TerminalType",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5010),
		"Message Set ID", // MessageSetID
		"MessageSetID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5020),
		"End Message ID", // EndMessageID
		"EndMessageID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5110),
		"Display Format", // DisplayFormat
		"DisplayFormat",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5120),
		"Page Position ID", // PagePositionID
		"PagePositionID",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5130),
		"Text Format ID", // TextFormatID
		"TextFormatID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5140),
		"Normal/Reverse", // NormalReverse
		"NormalReverse",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5150),
		"Add Gray Scale", // AddGrayScale
		"AddGrayScale",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5160),
		"Borders", // Borders
		"Borders",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5170),
		"Copies", // Copies
		"Copies",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5180),
		"Command Magnification Type", // CommandMagnificationType
		"CommandMagnificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x5190),
		"Erase", // Erase
		"Erase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x51A0),
		"Print", // Print
		"Print",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0000, 0x51B0),
		"Overlays", // Overlays
		"Overlays",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0000),
		"File Meta Information Group Length", // FileMetaInformationGroupLength
		"FileMetaInformationGroupLength",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0001),
		"File Meta Information Version", // FileMetaInformationVersion
		"FileMetaInformationVersion",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0002),
		"Media Storage SOP Class UID", // MediaStorageSOPClassUID
		"MediaStorageSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0003),
		"Media Storage SOP Instance UID", // MediaStorageSOPInstanceUID
		"MediaStorageSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0010),
		"Transfer Syntax UID", // TransferSyntaxUID
		"TransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0012),
		"Implementation Class UID", // ImplementationClassUID
		"ImplementationClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0013),
		"Implementation Version Name", // ImplementationVersionName
		"ImplementationVersionName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0016),
		"Source Application Entity Title", // SourceApplicationEntityTitle
		"SourceApplicationEntityTitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0017),
		"Sending Application Entity Title", // SendingApplicationEntityTitle
		"SendingApplicationEntityTitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0018),
		"Receiving Application Entity Title", // ReceivingApplicationEntityTitle
		"ReceivingApplicationEntityTitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0026),
		"Source Presentation Address", // SourcePresentationAddress
		"SourcePresentationAddress",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0027),
		"Sending Presentation Address", // SendingPresentationAddress
		"SendingPresentationAddress",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0028),
		"Receiving Presentation Address", // ReceivingPresentationAddress
		"ReceivingPresentationAddress",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0031),
		"RTV Meta Information Version", // RTVMetaInformationVersion
		"RTVMetaInformationVersion",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0032),
		"RTV Communication SOP Class UID", // RTVCommunicationSOPClassUID
		"RTVCommunicationSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0033),
		"RTV Communication SOP Instance UID", // RTVCommunicationSOPInstanceUID
		"RTVCommunicationSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0035),
		"RTV Source Identifier", // RTVSourceIdentifier
		"RTVSourceIdentifier",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0036),
		"RTV Flow Identifier", // RTVFlowIdentifier
		"RTVFlowIdentifier",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0037),
		"RTV Flow RTP Sampling Rate", // RTVFlowRTPSamplingRate
		"RTVFlowRTPSamplingRate",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0038),
		"RTV Flow Actual Frame Duration", // RTVFlowActualFrameDuration
		"RTVFlowActualFrameDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0100),
		"Private Information Creator UID", // PrivateInformationCreatorUID
		"PrivateInformationCreatorUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0002, 0x0102),
		"Private Information", // PrivateInformation
		"PrivateInformation",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1130),
		"File-set ID", // FileSetID
		"FileSetID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1141),
		"File-set Descriptor File ID", // FileSetDescriptorFileID
		"FileSetDescriptorFileID",
		vm.VM1_8,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1142),
		"Specific Character Set of File-set Descriptor File", // SpecificCharacterSetOfFileSetDescriptorFile
		"SpecificCharacterSetOfFileSetDescriptorFile",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1200),
		"Offset of the First Directory Record of the Root Directory Entity", // OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity
		"OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1202),
		"Offset of the Last Directory Record of the Root Directory Entity", // OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity
		"OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1212),
		"File-set Consistency Flag", // FileSetConsistencyFlag
		"FileSetConsistencyFlag",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1220),
		"Directory Record Sequence", // DirectoryRecordSequence
		"DirectoryRecordSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1400),
		"Offset of the Next Directory Record", // OffsetOfTheNextDirectoryRecord
		"OffsetOfTheNextDirectoryRecord",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1410),
		"Record In-use Flag", // RecordInUseFlag
		"RecordInUseFlag",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1420),
		"Offset of Referenced Lower-Level Directory Entity", // OffsetOfReferencedLowerLevelDirectoryEntity
		"OffsetOfReferencedLowerLevelDirectoryEntity",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1430),
		"Directory Record Type", // DirectoryRecordType
		"DirectoryRecordType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1432),
		"Private Record UID", // PrivateRecordUID
		"PrivateRecordUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1500),
		"Referenced File ID", // ReferencedFileID
		"ReferencedFileID",
		vm.VM1_8,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1504),
		"MRDR Directory Record Offset", // MRDRDirectoryRecordOffset
		"MRDRDirectoryRecordOffset",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1510),
		"Referenced SOP Class UID in File", // ReferencedSOPClassUIDInFile
		"ReferencedSOPClassUIDInFile",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1511),
		"Referenced SOP Instance UID in File", // ReferencedSOPInstanceUIDInFile
		"ReferencedSOPInstanceUIDInFile",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1512),
		"Referenced Transfer Syntax UID in File", // ReferencedTransferSyntaxUIDInFile
		"ReferencedTransferSyntaxUIDInFile",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x151A),
		"Referenced Related General SOP Class UID in File", // ReferencedRelatedGeneralSOPClassUIDInFile
		"ReferencedRelatedGeneralSOPClassUIDInFile",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0004, 0x1600),
		"Number of References", // NumberOfReferences
		"NumberOfReferences",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0006, 0x0001),
		"Current Frame Functional Groups Sequence", // CurrentFrameFunctionalGroupsSequence
		"CurrentFrameFunctionalGroupsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0001),
		"Length to End", // LengthToEnd
		"LengthToEnd",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0005),
		"Specific Character Set", // SpecificCharacterSet
		"SpecificCharacterSet",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0006),
		"Language Code Sequence", // LanguageCodeSequence
		"LanguageCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0008),
		"Image Type", // ImageType
		"ImageType",
		vm.VM2_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0010),
		"Recognition Code", // RecognitionCode
		"RecognitionCode",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0012),
		"Instance Creation Date", // InstanceCreationDate
		"InstanceCreationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0013),
		"Instance Creation Time", // InstanceCreationTime
		"InstanceCreationTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0014),
		"Instance Creator UID", // InstanceCreatorUID
		"InstanceCreatorUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0015),
		"Instance Coercion DateTime", // InstanceCoercionDateTime
		"InstanceCoercionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0016),
		"SOP Class UID", // SOPClassUID
		"SOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0017),
		"Acquisition UID", // AcquisitionUID
		"AcquisitionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0018),
		"SOP Instance UID", // SOPInstanceUID
		"SOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0019),
		"Pyramid UID", // PyramidUID
		"PyramidUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x001A),
		"Related General SOP Class UID", // RelatedGeneralSOPClassUID
		"RelatedGeneralSOPClassUID",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x001B),
		"Original Specialized SOP Class UID", // OriginalSpecializedSOPClassUID
		"OriginalSpecializedSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x001C),
		"Synthetic Data", // SyntheticData
		"SyntheticData",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0020),
		"Study Date", // StudyDate
		"StudyDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0021),
		"Series Date", // SeriesDate
		"SeriesDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0022),
		"Acquisition Date", // AcquisitionDate
		"AcquisitionDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0023),
		"Content Date", // ContentDate
		"ContentDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0024),
		"Overlay Date", // OverlayDate
		"OverlayDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0025),
		"Curve Date", // CurveDate
		"CurveDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x002A),
		"Acquisition DateTime", // AcquisitionDateTime
		"AcquisitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0030),
		"Study Time", // StudyTime
		"StudyTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0031),
		"Series Time", // SeriesTime
		"SeriesTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0032),
		"Acquisition Time", // AcquisitionTime
		"AcquisitionTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0033),
		"Content Time", // ContentTime
		"ContentTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0034),
		"Overlay Time", // OverlayTime
		"OverlayTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0035),
		"Curve Time", // CurveTime
		"CurveTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0040),
		"Data Set Type", // DataSetType
		"DataSetType",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0041),
		"Data Set Subtype", // DataSetSubtype
		"DataSetSubtype",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0042),
		"Nuclear Medicine Series Type", // NuclearMedicineSeriesType
		"NuclearMedicineSeriesType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0050),
		"Accession Number", // AccessionNumber
		"AccessionNumber",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0051),
		"Issuer of Accession Number Sequence", // IssuerOfAccessionNumberSequence
		"IssuerOfAccessionNumberSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0052),
		"Query/Retrieve Level", // QueryRetrieveLevel
		"QueryRetrieveLevel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0053),
		"Query/Retrieve View", // QueryRetrieveView
		"QueryRetrieveView",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0054),
		"Retrieve AE Title", // RetrieveAETitle
		"RetrieveAETitle",
		vm.VM1_n,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0055),
		"Station AE Title", // StationAETitle
		"StationAETitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0056),
		"Instance Availability", // InstanceAvailability
		"InstanceAvailability",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0058),
		"Failed SOP Instance UID List", // FailedSOPInstanceUIDList
		"FailedSOPInstanceUIDList",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0060),
		"Modality", // Modality
		"Modality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0061),
		"Modalities in Study", // ModalitiesInStudy
		"ModalitiesInStudy",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0062),
		"SOP Classes in Study", // SOPClassesInStudy
		"SOPClassesInStudy",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0063),
		"Anatomic Regions in Study Code Sequence", // AnatomicRegionsInStudyCodeSequence
		"AnatomicRegionsInStudyCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0064),
		"Conversion Type", // ConversionType
		"ConversionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0068),
		"Presentation Intent Type", // PresentationIntentType
		"PresentationIntentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0070),
		"Manufacturer", // Manufacturer
		"Manufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0080),
		"Institution Name", // InstitutionName
		"InstitutionName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0081),
		"Institution Address", // InstitutionAddress
		"InstitutionAddress",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0082),
		"Institution Code Sequence", // InstitutionCodeSequence
		"InstitutionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0090),
		"Referring Physician's Name", // ReferringPhysicianName
		"ReferringPhysicianName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0092),
		"Referring Physician's Address", // ReferringPhysicianAddress
		"ReferringPhysicianAddress",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0094),
		"Referring Physician's Telephone Numbers", // ReferringPhysicianTelephoneNumbers
		"ReferringPhysicianTelephoneNumbers",
		vm.VM1_n,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0096),
		"Referring Physician Identification Sequence", // ReferringPhysicianIdentificationSequence
		"ReferringPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x009C),
		"Consulting Physician's Name", // ConsultingPhysicianName
		"ConsultingPhysicianName",
		vm.VM1_n,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x009D),
		"Consulting Physician Identification Sequence", // ConsultingPhysicianIdentificationSequence
		"ConsultingPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0100),
		"Code Value", // CodeValue
		"CodeValue",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0101),
		"Extended Code Value", // ExtendedCodeValue
		"ExtendedCodeValue",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0102),
		"Coding Scheme Designator", // CodingSchemeDesignator
		"CodingSchemeDesignator",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0103),
		"Coding Scheme Version", // CodingSchemeVersion
		"CodingSchemeVersion",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0104),
		"Code Meaning", // CodeMeaning
		"CodeMeaning",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0105),
		"Mapping Resource", // MappingResource
		"MappingResource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0106),
		"Context Group Version", // ContextGroupVersion
		"ContextGroupVersion",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0107),
		"Context Group Local Version", // ContextGroupLocalVersion
		"ContextGroupLocalVersion",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0108),
		"Extended Code Meaning", // ExtendedCodeMeaning
		"ExtendedCodeMeaning",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0109),
		"Coding Scheme Resources Sequence", // CodingSchemeResourcesSequence
		"CodingSchemeResourcesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010A),
		"Coding Scheme URL Type", // CodingSchemeURLType
		"CodingSchemeURLType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010B),
		"Context Group Extension Flag", // ContextGroupExtensionFlag
		"ContextGroupExtensionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010C),
		"Coding Scheme UID", // CodingSchemeUID
		"CodingSchemeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010D),
		"Context Group Extension Creator UID", // ContextGroupExtensionCreatorUID
		"ContextGroupExtensionCreatorUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010E),
		"Coding Scheme URL", // CodingSchemeURL
		"CodingSchemeURL",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x010F),
		"Context Identifier", // ContextIdentifier
		"ContextIdentifier",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0110),
		"Coding Scheme Identification Sequence", // CodingSchemeIdentificationSequence
		"CodingSchemeIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0112),
		"Coding Scheme Registry", // CodingSchemeRegistry
		"CodingSchemeRegistry",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0114),
		"Coding Scheme External ID", // CodingSchemeExternalID
		"CodingSchemeExternalID",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0115),
		"Coding Scheme Name", // CodingSchemeName
		"CodingSchemeName",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0116),
		"Coding Scheme Responsible Organization", // CodingSchemeResponsibleOrganization
		"CodingSchemeResponsibleOrganization",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0117),
		"Context UID", // ContextUID
		"ContextUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0118),
		"Mapping Resource UID", // MappingResourceUID
		"MappingResourceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0119),
		"Long Code Value", // LongCodeValue
		"LongCodeValue",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0120),
		"URN Code Value", // URNCodeValue
		"URNCodeValue",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0121),
		"Equivalent Code Sequence", // EquivalentCodeSequence
		"EquivalentCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0122),
		"Mapping Resource Name", // MappingResourceName
		"MappingResourceName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0123),
		"Context Group Identification Sequence", // ContextGroupIdentificationSequence
		"ContextGroupIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0124),
		"Mapping Resource Identification Sequence", // MappingResourceIdentificationSequence
		"MappingResourceIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0201),
		"Timezone Offset From UTC", // TimezoneOffsetFromUTC
		"TimezoneOffsetFromUTC",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0220),
		"Responsible Group Code Sequence", // ResponsibleGroupCodeSequence
		"ResponsibleGroupCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0221),
		"Equipment Modality", // EquipmentModality
		"EquipmentModality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0222),
		"Manufacturer's Related Model Group", // ManufacturerRelatedModelGroup
		"ManufacturerRelatedModelGroup",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0300),
		"Private Data Element Characteristics Sequence", // PrivateDataElementCharacteristicsSequence
		"PrivateDataElementCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0301),
		"Private Group Reference", // PrivateGroupReference
		"PrivateGroupReference",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0302),
		"Private Creator Reference", // PrivateCreatorReference
		"PrivateCreatorReference",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0303),
		"Block Identifying Information Status", // BlockIdentifyingInformationStatus
		"BlockIdentifyingInformationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0304),
		"Nonidentifying Private Elements", // NonidentifyingPrivateElements
		"NonidentifyingPrivateElements",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0305),
		"Deidentification Action Sequence", // DeidentificationActionSequence
		"DeidentificationActionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0306),
		"Identifying Private Elements", // IdentifyingPrivateElements
		"IdentifyingPrivateElements",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0307),
		"Deidentification Action", // DeidentificationAction
		"DeidentificationAction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0308),
		"Private Data Element", // PrivateDataElement
		"PrivateDataElement",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0309),
		"Private Data Element Value Multiplicity", // PrivateDataElementValueMultiplicity
		"PrivateDataElementValueMultiplicity",
		vm.VM1_3,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030A),
		"Private Data Element Value Representation", // PrivateDataElementValueRepresentation
		"PrivateDataElementValueRepresentation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030B),
		"Private Data Element Number of Items", // PrivateDataElementNumberOfItems
		"PrivateDataElementNumberOfItems",
		vm.VM1_2,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030C),
		"Private Data Element Name", // PrivateDataElementName
		"PrivateDataElementName",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030D),
		"Private Data Element Keyword", // PrivateDataElementKeyword
		"PrivateDataElementKeyword",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030E),
		"Private Data Element Description", // PrivateDataElementDescription
		"PrivateDataElementDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x030F),
		"Private Data Element Encoding", // PrivateDataElementEncoding
		"PrivateDataElementEncoding",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0310),
		"Private Data Element Definition Sequence", // PrivateDataElementDefinitionSequence
		"PrivateDataElementDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0400),
		"Scope of Inventory Sequence", // ScopeOfInventorySequence
		"ScopeOfInventorySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0401),
		"Inventory Purpose", // InventoryPurpose
		"InventoryPurpose",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0402),
		"Inventory Instance Description", // InventoryInstanceDescription
		"InventoryInstanceDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0403),
		"Inventory Level", // InventoryLevel
		"InventoryLevel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0404),
		"Item Inventory DateTime", // ItemInventoryDateTime
		"ItemInventoryDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0405),
		"Removed from Operational Use", // RemovedFromOperationalUse
		"RemovedFromOperationalUse",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0406),
		"Reason for Removal Code Sequence", // ReasonForRemovalCodeSequence
		"ReasonForRemovalCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0407),
		"Stored Instance Base URI", // StoredInstanceBaseURI
		"StoredInstanceBaseURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0408),
		"Folder Access URI", // FolderAccessURI
		"FolderAccessURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0409),
		"File Access URI", // FileAccessURI
		"FileAccessURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040A),
		"Container File Type", // ContainerFileType
		"ContainerFileType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040B),
		"Filename in Container", // FilenameInContainer
		"FilenameInContainer",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040C),
		"File Offset in Container", // FileOffsetInContainer
		"FileOffsetInContainer",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040D),
		"File Length in Container", // FileLengthInContainer
		"FileLengthInContainer",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040E),
		"Stored Instance Transfer Syntax UID", // StoredInstanceTransferSyntaxUID
		"StoredInstanceTransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x040F),
		"Extended Matching Mechanisms", // ExtendedMatchingMechanisms
		"ExtendedMatchingMechanisms",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0410),
		"Range Matching Sequence", // RangeMatchingSequence
		"RangeMatchingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0411),
		"List of UID Matching Sequence", // ListOfUIDMatchingSequence
		"ListOfUIDMatchingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0412),
		"Empty Value Matching Sequence", // EmptyValueMatchingSequence
		"EmptyValueMatchingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0413),
		"General Matching Sequence", // GeneralMatchingSequence
		"GeneralMatchingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0414),
		"Requested Status Interval", // RequestedStatusInterval
		"RequestedStatusInterval",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0415),
		"Retain Instances", // RetainInstances
		"RetainInstances",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0416),
		"Expiration DateTime", // ExpirationDateTime
		"ExpirationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0417),
		"Transaction Status", // TransactionStatus
		"TransactionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0418),
		"Transaction Status Comment", // TransactionStatusComment
		"TransactionStatusComment",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0419),
		"File Set Access Sequence", // FileSetAccessSequence
		"FileSetAccessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041A),
		"File Access Sequence", // FileAccessSequence
		"FileAccessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041B),
		"Record Key", // RecordKey
		"RecordKey",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041C),
		"Prior Record Key", // PriorRecordKey
		"PriorRecordKey",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041D),
		"Metadata Sequence", // MetadataSequence
		"MetadataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041E),
		"Updated Metadata Sequence", // UpdatedMetadataSequence
		"UpdatedMetadataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x041F),
		"Study Update DateTime", // StudyUpdateDateTime
		"StudyUpdateDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0420),
		"Inventory Access End Points Sequence", // InventoryAccessEndPointsSequence
		"InventoryAccessEndPointsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0421),
		"Study Access End Points Sequence", // StudyAccessEndPointsSequence
		"StudyAccessEndPointsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0422),
		"Incorporated Inventory Instance Sequence", // IncorporatedInventoryInstanceSequence
		"IncorporatedInventoryInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0423),
		"Inventoried Studies Sequence", // InventoriedStudiesSequence
		"InventoriedStudiesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0424),
		"Inventoried Series Sequence", // InventoriedSeriesSequence
		"InventoriedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0425),
		"Inventoried Instances Sequence", // InventoriedInstancesSequence
		"InventoriedInstancesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0426),
		"Inventory Completion Status", // InventoryCompletionStatus
		"InventoryCompletionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0427),
		"Number of Study Records in Instance", // NumberOfStudyRecordsInInstance
		"NumberOfStudyRecordsInInstance",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0428),
		"Total Number of Study Records", // TotalNumberOfStudyRecords
		"TotalNumberOfStudyRecords",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x0429),
		"Maximum Number of Records", // MaximumNumberOfRecords
		"MaximumNumberOfRecords",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1000),
		"Network ID", // NetworkID
		"NetworkID",
		vm.VM1,
		true,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1010),
		"Station Name", // StationName
		"StationName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1030),
		"Study Description", // StudyDescription
		"StudyDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1032),
		"Procedure Code Sequence", // ProcedureCodeSequence
		"ProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x103E),
		"Series Description", // SeriesDescription
		"SeriesDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x103F),
		"Series Description Code Sequence", // SeriesDescriptionCodeSequence
		"SeriesDescriptionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1040),
		"Institutional Department Name", // InstitutionalDepartmentName
		"InstitutionalDepartmentName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1041),
		"Institutional Department Type Code Sequence", // InstitutionalDepartmentTypeCodeSequence
		"InstitutionalDepartmentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1048),
		"Physician(s) of Record", // PhysiciansOfRecord
		"PhysiciansOfRecord",
		vm.VM1_n,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1049),
		"Physician(s) of Record Identification Sequence", // PhysiciansOfRecordIdentificationSequence
		"PhysiciansOfRecordIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1050),
		"Performing Physician's Name", // PerformingPhysicianName
		"PerformingPhysicianName",
		vm.VM1_n,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1052),
		"Performing Physician Identification Sequence", // PerformingPhysicianIdentificationSequence
		"PerformingPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1060),
		"Name of Physician(s) Reading Study", // NameOfPhysiciansReadingStudy
		"NameOfPhysiciansReadingStudy",
		vm.VM1_n,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1062),
		"Physician(s) Reading Study Identification Sequence", // PhysiciansReadingStudyIdentificationSequence
		"PhysiciansReadingStudyIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1070),
		"Operators' Name", // OperatorsName
		"OperatorsName",
		vm.VM1_n,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1072),
		"Operator Identification Sequence", // OperatorIdentificationSequence
		"OperatorIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1080),
		"Admitting Diagnoses Description", // AdmittingDiagnosesDescription
		"AdmittingDiagnosesDescription",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1084),
		"Admitting Diagnoses Code Sequence", // AdmittingDiagnosesCodeSequence
		"AdmittingDiagnosesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1088),
		"Pyramid Description", // PyramidDescription
		"PyramidDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1090),
		"Manufacturer's Model Name", // ManufacturerModelName
		"ManufacturerModelName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1100),
		"Referenced Results Sequence", // ReferencedResultsSequence
		"ReferencedResultsSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1110),
		"Referenced Study Sequence", // ReferencedStudySequence
		"ReferencedStudySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1111),
		"Referenced Performed Procedure Step Sequence", // ReferencedPerformedProcedureStepSequence
		"ReferencedPerformedProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1112),
		"Referenced Instances by SOP Class Sequence", // ReferencedInstancesBySOPClassSequence
		"ReferencedInstancesBySOPClassSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1115),
		"Referenced Series Sequence", // ReferencedSeriesSequence
		"ReferencedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1120),
		"Referenced Patient Sequence", // ReferencedPatientSequence
		"ReferencedPatientSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1125),
		"Referenced Visit Sequence", // ReferencedVisitSequence
		"ReferencedVisitSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1130),
		"Referenced Overlay Sequence", // ReferencedOverlaySequence
		"ReferencedOverlaySequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1134),
		"Referenced Stereometric Instance Sequence", // ReferencedStereometricInstanceSequence
		"ReferencedStereometricInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x113A),
		"Referenced Waveform Sequence", // ReferencedWaveformSequence
		"ReferencedWaveformSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1140),
		"Referenced Image Sequence", // ReferencedImageSequence
		"ReferencedImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1145),
		"Referenced Curve Sequence", // ReferencedCurveSequence
		"ReferencedCurveSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x114A),
		"Referenced Instance Sequence", // ReferencedInstanceSequence
		"ReferencedInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x114B),
		"Referenced Real World Value Mapping Instance Sequence", // ReferencedRealWorldValueMappingInstanceSequence
		"ReferencedRealWorldValueMappingInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x114C),
		"Referenced Segmentation Sequence", // ReferencedSegmentationSequence
		"ReferencedSegmentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x114D),
		"Referenced Surface Segmentation Sequence", // ReferencedSurfaceSegmentationSequence
		"ReferencedSurfaceSegmentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1150),
		"Referenced SOP Class UID", // ReferencedSOPClassUID
		"ReferencedSOPClassUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1155),
		"Referenced SOP Instance UID", // ReferencedSOPInstanceUID
		"ReferencedSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1156),
		"Definition Source Sequence", // DefinitionSourceSequence
		"DefinitionSourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x115A),
		"SOP Classes Supported", // SOPClassesSupported
		"SOPClassesSupported",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1160),
		"Referenced Frame Number", // ReferencedFrameNumber
		"ReferencedFrameNumber",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1161),
		"Simple Frame List", // SimpleFrameList
		"SimpleFrameList",
		vm.VM1_n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1162),
		"Calculated Frame List", // CalculatedFrameList
		"CalculatedFrameList",
		vm.VM3_3n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1163),
		"Time Range", // TimeRange
		"TimeRange",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1164),
		"Frame Extraction Sequence", // FrameExtractionSequence
		"FrameExtractionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1167),
		"Multi-frame Source SOP Instance UID", // MultiFrameSourceSOPInstanceUID
		"MultiFrameSourceSOPInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1190),
		"Retrieve URL", // RetrieveURL
		"RetrieveURL",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1195),
		"Transaction UID", // TransactionUID
		"TransactionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1196),
		"Warning Reason", // WarningReason
		"WarningReason",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1197),
		"Failure Reason", // FailureReason
		"FailureReason",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1198),
		"Failed SOP Sequence", // FailedSOPSequence
		"FailedSOPSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1199),
		"Referenced SOP Sequence", // ReferencedSOPSequence
		"ReferencedSOPSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x119A),
		"Other Failures Sequence", // OtherFailuresSequence
		"OtherFailuresSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x119B),
		"Failed Study Sequence", // FailedStudySequence
		"FailedStudySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1200),
		"Studies Containing Other Referenced Instances Sequence", // StudiesContainingOtherReferencedInstancesSequence
		"StudiesContainingOtherReferencedInstancesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1250),
		"Related Series Sequence", // RelatedSeriesSequence
		"RelatedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1301),
		"Principal Diagnosis Code Sequence", // PrincipalDiagnosisCodeSequence
		"PrincipalDiagnosisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1302),
		"Primary Diagnosis Code Sequence", // PrimaryDiagnosisCodeSequence
		"PrimaryDiagnosisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1303),
		"Secondary Diagnoses Code Sequence", // SecondaryDiagnosesCodeSequence
		"SecondaryDiagnosesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x1304),
		"Histological Diagnoses Code Sequence", // HistologicalDiagnosesCodeSequence
		"HistologicalDiagnosesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2110),
		"Lossy Image Compression (Retired)", // LossyImageCompressionRetired
		"LossyImageCompressionRetired",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2111),
		"Derivation Description", // DerivationDescription
		"DerivationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2112),
		"Source Image Sequence", // SourceImageSequence
		"SourceImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2120),
		"Stage Name", // StageName
		"StageName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2122),
		"Stage Number", // StageNumber
		"StageNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2124),
		"Number of Stages", // NumberOfStages
		"NumberOfStages",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2127),
		"View Name", // ViewName
		"ViewName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2128),
		"View Number", // ViewNumber
		"ViewNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2129),
		"Number of Event Timers", // NumberOfEventTimers
		"NumberOfEventTimers",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x212A),
		"Number of Views in Stage", // NumberOfViewsInStage
		"NumberOfViewsInStage",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2130),
		"Event Elapsed Time(s)", // EventElapsedTimes
		"EventElapsedTimes",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2132),
		"Event Timer Name(s)", // EventTimerNames
		"EventTimerNames",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2133),
		"Event Timer Sequence", // EventTimerSequence
		"EventTimerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2134),
		"Event Time Offset", // EventTimeOffset
		"EventTimeOffset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2135),
		"Event Code Sequence", // EventCodeSequence
		"EventCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2142),
		"Start Trim", // StartTrim
		"StartTrim",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2143),
		"Stop Trim", // StopTrim
		"StopTrim",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2144),
		"Recommended Display Frame Rate", // RecommendedDisplayFrameRate
		"RecommendedDisplayFrameRate",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2200),
		"Transducer Position", // TransducerPosition
		"TransducerPosition",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2204),
		"Transducer Orientation", // TransducerOrientation
		"TransducerOrientation",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2208),
		"Anatomic Structure", // AnatomicStructure
		"AnatomicStructure",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2218),
		"Anatomic Region Sequence", // AnatomicRegionSequence
		"AnatomicRegionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2220),
		"Anatomic Region Modifier Sequence", // AnatomicRegionModifierSequence
		"AnatomicRegionModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2228),
		"Primary Anatomic Structure Sequence", // PrimaryAnatomicStructureSequence
		"PrimaryAnatomicStructureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2229),
		"Anatomic Structure, Space or Region Sequence", // AnatomicStructureSpaceOrRegionSequence
		"AnatomicStructureSpaceOrRegionSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2230),
		"Primary Anatomic Structure Modifier Sequence", // PrimaryAnatomicStructureModifierSequence
		"PrimaryAnatomicStructureModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2240),
		"Transducer Position Sequence", // TransducerPositionSequence
		"TransducerPositionSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2242),
		"Transducer Position Modifier Sequence", // TransducerPositionModifierSequence
		"TransducerPositionModifierSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2244),
		"Transducer Orientation Sequence", // TransducerOrientationSequence
		"TransducerOrientationSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2246),
		"Transducer Orientation Modifier Sequence", // TransducerOrientationModifierSequence
		"TransducerOrientationModifierSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2251),
		"Anatomic Structure Space Or Region Code Sequence (Trial)", // AnatomicStructureSpaceOrRegionCodeSequenceTrial
		"AnatomicStructureSpaceOrRegionCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2253),
		"Anatomic Portal Of Entrance Code Sequence (Trial)", // AnatomicPortalOfEntranceCodeSequenceTrial
		"AnatomicPortalOfEntranceCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2255),
		"Anatomic Approach Direction Code Sequence (Trial)", // AnatomicApproachDirectionCodeSequenceTrial
		"AnatomicApproachDirectionCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2256),
		"Anatomic Perspective Description (Trial)", // AnatomicPerspectiveDescriptionTrial
		"AnatomicPerspectiveDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2257),
		"Anatomic Perspective Code Sequence (Trial)", // AnatomicPerspectiveCodeSequenceTrial
		"AnatomicPerspectiveCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2258),
		"Anatomic Location Of Examining Instrument Description (Trial)", // AnatomicLocationOfExaminingInstrumentDescriptionTrial
		"AnatomicLocationOfExaminingInstrumentDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x2259),
		"Anatomic Location Of Examining Instrument Code Sequence (Trial)", // AnatomicLocationOfExaminingInstrumentCodeSequenceTrial
		"AnatomicLocationOfExaminingInstrumentCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x225A),
		"Anatomic Structure Space Or Region Modifier Code Sequence (Trial)", // AnatomicStructureSpaceOrRegionModifierCodeSequenceTrial
		"AnatomicStructureSpaceOrRegionModifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x225C),
		"On Axis Background Anatomic Structure Code Sequence (Trial)", // OnAxisBackgroundAnatomicStructureCodeSequenceTrial
		"OnAxisBackgroundAnatomicStructureCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3001),
		"Alternate Representation Sequence", // AlternateRepresentationSequence
		"AlternateRepresentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3002),
		"Available Transfer Syntax UID", // AvailableTransferSyntaxUID
		"AvailableTransferSyntaxUID",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3010),
		"Irradiation Event UID", // IrradiationEventUID
		"IrradiationEventUID",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3011),
		"Source Irradiation Event Sequence", // SourceIrradiationEventSequence
		"SourceIrradiationEventSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x3012),
		"Radiopharmaceutical Administration Event UID", // RadiopharmaceuticalAdministrationEventUID
		"RadiopharmaceuticalAdministrationEventUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x4000),
		"Identifying Comments", // IdentifyingComments
		"IdentifyingComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9007),
		"Frame Type", // FrameType
		"FrameType",
		vm.MustParse("4-5"),
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9092),
		"Referenced Image Evidence Sequence", // ReferencedImageEvidenceSequence
		"ReferencedImageEvidenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9121),
		"Referenced Raw Data Sequence", // ReferencedRawDataSequence
		"ReferencedRawDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9123),
		"Creator-Version UID", // CreatorVersionUID
		"CreatorVersionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9124),
		"Derivation Image Sequence", // DerivationImageSequence
		"DerivationImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9154),
		"Source Image Evidence Sequence", // SourceImageEvidenceSequence
		"SourceImageEvidenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9205),
		"Pixel Presentation", // PixelPresentation
		"PixelPresentation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9206),
		"Volumetric Properties", // VolumetricProperties
		"VolumetricProperties",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9207),
		"Volume Based Calculation Technique", // VolumeBasedCalculationTechnique
		"VolumeBasedCalculationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9208),
		"Complex Image Component", // ComplexImageComponent
		"ComplexImageComponent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9209),
		"Acquisition Contrast", // AcquisitionContrast
		"AcquisitionContrast",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9215),
		"Derivation Code Sequence", // DerivationCodeSequence
		"DerivationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9237),
		"Referenced Presentation State Sequence", // ReferencedPresentationStateSequence
		"ReferencedPresentationStateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9410),
		"Referenced Other Plane Sequence", // ReferencedOtherPlaneSequence
		"ReferencedOtherPlaneSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9458),
		"Frame Display Sequence", // FrameDisplaySequence
		"FrameDisplaySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9459),
		"Recommended Display Frame Rate in Float", // RecommendedDisplayFrameRateInFloat
		"RecommendedDisplayFrameRateInFloat",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0008, 0x9460),
		"Skip Frame Range Flag", // SkipFrameRangeFlag
		"SkipFrameRangeFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0010),
		"Patient's Name", // PatientName
		"PatientName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0011),
		"Person Names to Use Sequence", // PersonNamesToUseSequence
		"PersonNamesToUseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0012),
		"Name to Use", // NameToUse
		"NameToUse",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0013),
		"Name to Use Comment", // NameToUseComment
		"NameToUseComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0014),
		"Third Person Pronouns Sequence", // ThirdPersonPronounsSequence
		"ThirdPersonPronounsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0015),
		"Pronoun Code Sequence", // PronounCodeSequence
		"PronounCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0016),
		"Pronoun Comment", // PronounComment
		"PronounComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0020),
		"Patient ID", // PatientID
		"PatientID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0021),
		"Issuer of Patient ID", // IssuerOfPatientID
		"IssuerOfPatientID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0022),
		"Type of Patient ID", // TypeOfPatientID
		"TypeOfPatientID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0024),
		"Issuer of Patient ID Qualifiers Sequence", // IssuerOfPatientIDQualifiersSequence
		"IssuerOfPatientIDQualifiersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0026),
		"Source Patient Group Identification Sequence", // SourcePatientGroupIdentificationSequence
		"SourcePatientGroupIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0027),
		"Group of Patients Identification Sequence", // GroupOfPatientsIdentificationSequence
		"GroupOfPatientsIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0028),
		"Subject Relative Position in Image", // SubjectRelativePositionInImage
		"SubjectRelativePositionInImage",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0030),
		"Patient's Birth Date", // PatientBirthDate
		"PatientBirthDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0032),
		"Patient's Birth Time", // PatientBirthTime
		"PatientBirthTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0033),
		"Patient's Birth Date in Alternative Calendar", // PatientBirthDateInAlternativeCalendar
		"PatientBirthDateInAlternativeCalendar",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0034),
		"Patient's Death Date in Alternative Calendar", // PatientDeathDateInAlternativeCalendar
		"PatientDeathDateInAlternativeCalendar",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0035),
		"Patient's Alternative Calendar", // PatientAlternativeCalendar
		"PatientAlternativeCalendar",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0040),
		"Patient's Sex", // PatientSex
		"PatientSex",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0041),
		"Gender Identity Sequence", // GenderIdentitySequence
		"GenderIdentitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0042),
		"Sex Parameters for Clinical Use Category Comment", // SexParametersForClinicalUseCategoryComment
		"SexParametersForClinicalUseCategoryComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0043),
		"Sex Parameters for Clinical Use Category Sequence", // SexParametersForClinicalUseCategorySequence
		"SexParametersForClinicalUseCategorySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0044),
		"Gender Identity Code Sequence", // GenderIdentityCodeSequence
		"GenderIdentityCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0045),
		"Gender Identity Comment", // GenderIdentityComment
		"GenderIdentityComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0046),
		"Sex Parameters for Clinical Use Category Code Sequence", // SexParametersForClinicalUseCategoryCodeSequence
		"SexParametersForClinicalUseCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0047),
		"Sex Parameters for Clinical Use Category Reference", // SexParametersForClinicalUseCategoryReference
		"SexParametersForClinicalUseCategoryReference",
		vm.VM1_n,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0050),
		"Patient's Insurance Plan Code Sequence", // PatientInsurancePlanCodeSequence
		"PatientInsurancePlanCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0101),
		"Patient's Primary Language Code Sequence", // PatientPrimaryLanguageCodeSequence
		"PatientPrimaryLanguageCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0102),
		"Patient's Primary Language Modifier Code Sequence", // PatientPrimaryLanguageModifierCodeSequence
		"PatientPrimaryLanguageModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0200),
		"Quality Control Subject", // QualityControlSubject
		"QualityControlSubject",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0201),
		"Quality Control Subject Type Code Sequence", // QualityControlSubjectTypeCodeSequence
		"QualityControlSubjectTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0212),
		"Strain Description", // StrainDescription
		"StrainDescription",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0213),
		"Strain Nomenclature", // StrainNomenclature
		"StrainNomenclature",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0214),
		"Strain Stock Number", // StrainStockNumber
		"StrainStockNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0215),
		"Strain Source Registry Code Sequence", // StrainSourceRegistryCodeSequence
		"StrainSourceRegistryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0216),
		"Strain Stock Sequence", // StrainStockSequence
		"StrainStockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0217),
		"Strain Source", // StrainSource
		"StrainSource",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0218),
		"Strain Additional Information", // StrainAdditionalInformation
		"StrainAdditionalInformation",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0219),
		"Strain Code Sequence", // StrainCodeSequence
		"StrainCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0221),
		"Genetic Modifications Sequence", // GeneticModificationsSequence
		"GeneticModificationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0222),
		"Genetic Modifications Description", // GeneticModificationsDescription
		"GeneticModificationsDescription",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0223),
		"Genetic Modifications Nomenclature", // GeneticModificationsNomenclature
		"GeneticModificationsNomenclature",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x0229),
		"Genetic Modifications Code Sequence", // GeneticModificationsCodeSequence
		"GeneticModificationsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1000),
		"Other Patient IDs", // OtherPatientIDs
		"OtherPatientIDs",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1001),
		"Other Patient Names", // OtherPatientNames
		"OtherPatientNames",
		vm.VM1_n,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1002),
		"Other Patient IDs Sequence", // OtherPatientIDsSequence
		"OtherPatientIDsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1005),
		"Patient's Birth Name", // PatientBirthName
		"PatientBirthName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1010),
		"Patient's Age", // PatientAge
		"PatientAge",
		vm.VM1,
		false,
		vr.AS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1020),
		"Patient's Size", // PatientSize
		"PatientSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1021),
		"Patient's Size Code Sequence", // PatientSizeCodeSequence
		"PatientSizeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1022),
		"Patient's Body Mass Index", // PatientBodyMassIndex
		"PatientBodyMassIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1023),
		"Measured AP Dimension", // MeasuredAPDimension
		"MeasuredAPDimension",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1024),
		"Measured Lateral Dimension", // MeasuredLateralDimension
		"MeasuredLateralDimension",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1030),
		"Patient's Weight", // PatientWeight
		"PatientWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1040),
		"Patient's Address", // PatientAddress
		"PatientAddress",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1050),
		"Insurance Plan Identification", // InsurancePlanIdentification
		"InsurancePlanIdentification",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1060),
		"Patient's Mother's Birth Name", // PatientMotherBirthName
		"PatientMotherBirthName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1080),
		"Military Rank", // MilitaryRank
		"MilitaryRank",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1081),
		"Branch of Service", // BranchOfService
		"BranchOfService",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1090),
		"Medical Record Locator", // MedicalRecordLocator
		"MedicalRecordLocator",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x1100),
		"Referenced Patient Photo Sequence", // ReferencedPatientPhotoSequence
		"ReferencedPatientPhotoSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2000),
		"Medical Alerts", // MedicalAlerts
		"MedicalAlerts",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2110),
		"Allergies", // Allergies
		"Allergies",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2150),
		"Country of Residence", // CountryOfResidence
		"CountryOfResidence",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2152),
		"Region of Residence", // RegionOfResidence
		"RegionOfResidence",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2154),
		"Patient's Telephone Numbers", // PatientTelephoneNumbers
		"PatientTelephoneNumbers",
		vm.VM1_n,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2155),
		"Patient's Telecom Information", // PatientTelecomInformation
		"PatientTelecomInformation",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2160),
		"Ethnic Group", // EthnicGroup
		"EthnicGroup",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2161),
		"Ethnic Group Code Sequence", // EthnicGroupCodeSequence
		"EthnicGroupCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2162),
		"Ethnic Groups", // EthnicGroups
		"EthnicGroups",
		vm.VM1_n,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2180),
		"Occupation", // Occupation
		"Occupation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21A0),
		"Smoking Status", // SmokingStatus
		"SmokingStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21B0),
		"Additional Patient History", // AdditionalPatientHistory
		"AdditionalPatientHistory",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21C0),
		"Pregnancy Status", // PregnancyStatus
		"PregnancyStatus",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21D0),
		"Last Menstrual Date", // LastMenstrualDate
		"LastMenstrualDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x21F0),
		"Patient's Religious Preference", // PatientReligiousPreference
		"PatientReligiousPreference",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2201),
		"Patient Species Description", // PatientSpeciesDescription
		"PatientSpeciesDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2202),
		"Patient Species Code Sequence", // PatientSpeciesCodeSequence
		"PatientSpeciesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2203),
		"Patient's Sex Neutered", // PatientSexNeutered
		"PatientSexNeutered",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2210),
		"Anatomical Orientation Type", // AnatomicalOrientationType
		"AnatomicalOrientationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2292),
		"Patient Breed Description", // PatientBreedDescription
		"PatientBreedDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2293),
		"Patient Breed Code Sequence", // PatientBreedCodeSequence
		"PatientBreedCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2294),
		"Breed Registration Sequence", // BreedRegistrationSequence
		"BreedRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2295),
		"Breed Registration Number", // BreedRegistrationNumber
		"BreedRegistrationNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2296),
		"Breed Registry Code Sequence", // BreedRegistryCodeSequence
		"BreedRegistryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2297),
		"Responsible Person", // ResponsiblePerson
		"ResponsiblePerson",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2298),
		"Responsible Person Role", // ResponsiblePersonRole
		"ResponsiblePersonRole",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x2299),
		"Responsible Organization", // ResponsibleOrganization
		"ResponsibleOrganization",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x4000),
		"Patient Comments", // PatientComments
		"PatientComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0010, 0x9431),
		"Examined Body Thickness", // ExaminedBodyThickness
		"ExaminedBodyThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0010),
		"Clinical Trial Sponsor Name", // ClinicalTrialSponsorName
		"ClinicalTrialSponsorName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0020),
		"Clinical Trial Protocol ID", // ClinicalTrialProtocolID
		"ClinicalTrialProtocolID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0021),
		"Clinical Trial Protocol Name", // ClinicalTrialProtocolName
		"ClinicalTrialProtocolName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0022),
		"Issuer of Clinical Trial Protocol ID", // IssuerOfClinicalTrialProtocolID
		"IssuerOfClinicalTrialProtocolID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0023),
		"Other Clinical Trial Protocol IDs Sequence", // OtherClinicalTrialProtocolIDsSequence
		"OtherClinicalTrialProtocolIDsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0030),
		"Clinical Trial Site ID", // ClinicalTrialSiteID
		"ClinicalTrialSiteID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0031),
		"Clinical Trial Site Name", // ClinicalTrialSiteName
		"ClinicalTrialSiteName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0032),
		"Issuer of Clinical Trial Site ID", // IssuerOfClinicalTrialSiteID
		"IssuerOfClinicalTrialSiteID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0040),
		"Clinical Trial Subject ID", // ClinicalTrialSubjectID
		"ClinicalTrialSubjectID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0041),
		"Issuer of Clinical Trial Subject ID", // IssuerOfClinicalTrialSubjectID
		"IssuerOfClinicalTrialSubjectID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0042),
		"Clinical Trial Subject Reading ID", // ClinicalTrialSubjectReadingID
		"ClinicalTrialSubjectReadingID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0043),
		"Issuer of Clinical Trial Subject Reading ID", // IssuerOfClinicalTrialSubjectReadingID
		"IssuerOfClinicalTrialSubjectReadingID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0050),
		"Clinical Trial Time Point ID", // ClinicalTrialTimePointID
		"ClinicalTrialTimePointID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0051),
		"Clinical Trial Time Point Description", // ClinicalTrialTimePointDescription
		"ClinicalTrialTimePointDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0052),
		"Longitudinal Temporal Offset from Event", // LongitudinalTemporalOffsetFromEvent
		"LongitudinalTemporalOffsetFromEvent",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0053),
		"Longitudinal Temporal Event Type", // LongitudinalTemporalEventType
		"LongitudinalTemporalEventType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0054),
		"Clinical Trial Time Point Type Code Sequence", // ClinicalTrialTimePointTypeCodeSequence
		"ClinicalTrialTimePointTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0055),
		"Issuer of Clinical Trial Time Point ID", // IssuerOfClinicalTrialTimePointID
		"IssuerOfClinicalTrialTimePointID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0060),
		"Clinical Trial Coordinating Center Name", // ClinicalTrialCoordinatingCenterName
		"ClinicalTrialCoordinatingCenterName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0062),
		"Patient Identity Removed", // PatientIdentityRemoved
		"PatientIdentityRemoved",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0063),
		"De-identification Method", // DeidentificationMethod
		"DeidentificationMethod",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0064),
		"De-identification Method Code Sequence", // DeidentificationMethodCodeSequence
		"DeidentificationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0071),
		"Clinical Trial Series ID", // ClinicalTrialSeriesID
		"ClinicalTrialSeriesID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0072),
		"Clinical Trial Series Description", // ClinicalTrialSeriesDescription
		"ClinicalTrialSeriesDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0073),
		"Issuer of Clinical Trial Series ID", // IssuerOfClinicalTrialSeriesID
		"IssuerOfClinicalTrialSeriesID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0081),
		"Clinical Trial Protocol Ethics Committee Name", // ClinicalTrialProtocolEthicsCommitteeName
		"ClinicalTrialProtocolEthicsCommitteeName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0082),
		"Clinical Trial Protocol Ethics Committee Approval Number", // ClinicalTrialProtocolEthicsCommitteeApprovalNumber
		"ClinicalTrialProtocolEthicsCommitteeApprovalNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0083),
		"Consent for Clinical Trial Use Sequence", // ConsentForClinicalTrialUseSequence
		"ConsentForClinicalTrialUseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0084),
		"Distribution Type", // DistributionType
		"DistributionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0085),
		"Consent for Distribution Flag", // ConsentForDistributionFlag
		"ConsentForDistributionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0086),
		"Ethics Committee Approval Effectiveness Start Date", // EthicsCommitteeApprovalEffectivenessStartDate
		"EthicsCommitteeApprovalEffectivenessStartDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0012, 0x0087),
		"Ethics Committee Approval Effectiveness End Date", // EthicsCommitteeApprovalEffectivenessEndDate
		"EthicsCommitteeApprovalEffectivenessEndDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0023),
		"CAD File Format", // CADFileFormat
		"CADFileFormat",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0024),
		"Component Reference System", // ComponentReferenceSystem
		"ComponentReferenceSystem",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0025),
		"Component Manufacturing Procedure", // ComponentManufacturingProcedure
		"ComponentManufacturingProcedure",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0028),
		"Component Manufacturer", // ComponentManufacturer
		"ComponentManufacturer",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0030),
		"Material Thickness", // MaterialThickness
		"MaterialThickness",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0032),
		"Material Pipe Diameter", // MaterialPipeDiameter
		"MaterialPipeDiameter",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0034),
		"Material Isolation Diameter", // MaterialIsolationDiameter
		"MaterialIsolationDiameter",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0042),
		"Material Grade", // MaterialGrade
		"MaterialGrade",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0044),
		"Material Properties Description", // MaterialPropertiesDescription
		"MaterialPropertiesDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0045),
		"Material Properties File Format (Retired)", // MaterialPropertiesFileFormatRetired
		"MaterialPropertiesFileFormatRetired",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0046),
		"Material Notes", // MaterialNotes
		"MaterialNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0050),
		"Component Shape", // ComponentShape
		"ComponentShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0052),
		"Curvature Type", // CurvatureType
		"CurvatureType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0054),
		"Outer Diameter", // OuterDiameter
		"OuterDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0056),
		"Inner Diameter", // InnerDiameter
		"InnerDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0100),
		"Component Welder IDs", // ComponentWelderIDs
		"ComponentWelderIDs",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0101),
		"Secondary Approval Status", // SecondaryApprovalStatus
		"SecondaryApprovalStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0102),
		"Secondary Review Date", // SecondaryReviewDate
		"SecondaryReviewDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0103),
		"Secondary Review Time", // SecondaryReviewTime
		"SecondaryReviewTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0104),
		"Secondary Reviewer Name", // SecondaryReviewerName
		"SecondaryReviewerName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0105),
		"Repair ID", // RepairID
		"RepairID",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0106),
		"Multiple Component Approval Sequence", // MultipleComponentApprovalSequence
		"MultipleComponentApprovalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0107),
		"Other Approval Status", // OtherApprovalStatus
		"OtherApprovalStatus",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0108),
		"Other Secondary Approval Status", // OtherSecondaryApprovalStatus
		"OtherSecondaryApprovalStatus",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0200),
		"Data Element Label Sequence", // DataElementLabelSequence
		"DataElementLabelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0201),
		"Data Element Label Item Sequence", // DataElementLabelItemSequence
		"DataElementLabelItemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0202),
		"Data Element", // DataElement
		"DataElement",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0203),
		"Data Element Name", // DataElementName
		"DataElementName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0204),
		"Data Element Description", // DataElementDescription
		"DataElementDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0205),
		"Data Element Conditionality", // DataElementConditionality
		"DataElementConditionality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0206),
		"Data Element Minimum Characters", // DataElementMinimumCharacters
		"DataElementMinimumCharacters",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x0207),
		"Data Element Maximum Characters", // DataElementMaximumCharacters
		"DataElementMaximumCharacters",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x1010),
		"Actual Environmental Conditions", // ActualEnvironmentalConditions
		"ActualEnvironmentalConditions",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x1020),
		"Expiry Date", // ExpiryDate
		"ExpiryDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x1040),
		"Environmental Conditions", // EnvironmentalConditions
		"EnvironmentalConditions",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2002),
		"Evaluator Sequence", // EvaluatorSequence
		"EvaluatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2004),
		"Evaluator Number", // EvaluatorNumber
		"EvaluatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2006),
		"Evaluator Name", // EvaluatorName
		"EvaluatorName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2008),
		"Evaluation Attempt", // EvaluationAttempt
		"EvaluationAttempt",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2012),
		"Indication Sequence", // IndicationSequence
		"IndicationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2014),
		"Indication Number", // IndicationNumber
		"IndicationNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2016),
		"Indication Label", // IndicationLabel
		"IndicationLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2018),
		"Indication Description", // IndicationDescription
		"IndicationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x201A),
		"Indication Type", // IndicationType
		"IndicationType",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x201C),
		"Indication Disposition", // IndicationDisposition
		"IndicationDisposition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x201E),
		"Indication ROI Sequence", // IndicationROISequence
		"IndicationROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2030),
		"Indication Physical Property Sequence", // IndicationPhysicalPropertySequence
		"IndicationPhysicalPropertySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2032),
		"Property Label", // PropertyLabel
		"PropertyLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2202),
		"Coordinate System Number of Axes", // CoordinateSystemNumberOfAxes
		"CoordinateSystemNumberOfAxes",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2204),
		"Coordinate System Axes Sequence", // CoordinateSystemAxesSequence
		"CoordinateSystemAxesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2206),
		"Coordinate System Axis Description", // CoordinateSystemAxisDescription
		"CoordinateSystemAxisDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2208),
		"Coordinate System Data Set Mapping", // CoordinateSystemDataSetMapping
		"CoordinateSystemDataSetMapping",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x220A),
		"Coordinate System Axis Number", // CoordinateSystemAxisNumber
		"CoordinateSystemAxisNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x220C),
		"Coordinate System Axis Type", // CoordinateSystemAxisType
		"CoordinateSystemAxisType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x220E),
		"Coordinate System Axis Units", // CoordinateSystemAxisUnits
		"CoordinateSystemAxisUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2210),
		"Coordinate System Axis Values", // CoordinateSystemAxisValues
		"CoordinateSystemAxisValues",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2220),
		"Coordinate System Transform Sequence", // CoordinateSystemTransformSequence
		"CoordinateSystemTransformSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2222),
		"Transform Description", // TransformDescription
		"TransformDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2224),
		"Transform Number of Axes", // TransformNumberOfAxes
		"TransformNumberOfAxes",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2226),
		"Transform Order of Axes", // TransformOrderOfAxes
		"TransformOrderOfAxes",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x2228),
		"Transformed Axis Units", // TransformedAxisUnits
		"TransformedAxisUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x222A),
		"Coordinate System Transform Rotation and Scale Matrix", // CoordinateSystemTransformRotationAndScaleMatrix
		"CoordinateSystemTransformRotationAndScaleMatrix",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x222C),
		"Coordinate System Transform Translation Matrix", // CoordinateSystemTransformTranslationMatrix
		"CoordinateSystemTransformTranslationMatrix",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3011),
		"Internal Detector Frame Time", // InternalDetectorFrameTime
		"InternalDetectorFrameTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3012),
		"Number of Frames Integrated", // NumberOfFramesIntegrated
		"NumberOfFramesIntegrated",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3020),
		"Detector Temperature Sequence", // DetectorTemperatureSequence
		"DetectorTemperatureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3022),
		"Sensor Name", // SensorName
		"SensorName",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3024),
		"Horizontal Offset of Sensor", // HorizontalOffsetOfSensor
		"HorizontalOffsetOfSensor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3026),
		"Vertical Offset of Sensor", // VerticalOffsetOfSensor
		"VerticalOffsetOfSensor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3028),
		"Sensor Temperature", // SensorTemperature
		"SensorTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3040),
		"Dark Current Sequence", // DarkCurrentSequence
		"DarkCurrentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3050),
		"Dark Current Counts", // DarkCurrentCounts
		"DarkCurrentCounts",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3060),
		"Gain Correction Reference Sequence", // GainCorrectionReferenceSequence
		"GainCorrectionReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3070),
		"Air Counts", // AirCounts
		"AirCounts",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3071),
		"KV Used in Gain Calibration", // KVUsedInGainCalibration
		"KVUsedInGainCalibration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3072),
		"MA Used in Gain Calibration", // MAUsedInGainCalibration
		"MAUsedInGainCalibration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3073),
		"Number of Frames Used for Integration", // NumberOfFramesUsedForIntegration
		"NumberOfFramesUsedForIntegration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3074),
		"Filter Material Used in Gain Calibration", // FilterMaterialUsedInGainCalibration
		"FilterMaterialUsedInGainCalibration",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3075),
		"Filter Thickness Used in Gain Calibration", // FilterThicknessUsedInGainCalibration
		"FilterThicknessUsedInGainCalibration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3076),
		"Date of Gain Calibration", // DateOfGainCalibration
		"DateOfGainCalibration",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3077),
		"Time of Gain Calibration", // TimeOfGainCalibration
		"TimeOfGainCalibration",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3080),
		"Bad Pixel Image", // BadPixelImage
		"BadPixelImage",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3099),
		"Calibration Notes", // CalibrationNotes
		"CalibrationNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3100),
		"Linearity Correction Technique", // LinearityCorrectionTechnique
		"LinearityCorrectionTechnique",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x3101),
		"Beam Hardening Correction Technique", // BeamHardeningCorrectionTechnique
		"BeamHardeningCorrectionTechnique",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4002),
		"Pulser Equipment Sequence", // PulserEquipmentSequence
		"PulserEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4004),
		"Pulser Type", // PulserType
		"PulserType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4006),
		"Pulser Notes", // PulserNotes
		"PulserNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4008),
		"Receiver Equipment Sequence", // ReceiverEquipmentSequence
		"ReceiverEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x400A),
		"Amplifier Type", // AmplifierType
		"AmplifierType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x400C),
		"Receiver Notes", // ReceiverNotes
		"ReceiverNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x400E),
		"Pre-Amplifier Equipment Sequence", // PreAmplifierEquipmentSequence
		"PreAmplifierEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x400F),
		"Pre-Amplifier Notes", // PreAmplifierNotes
		"PreAmplifierNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4010),
		"Transmit Transducer Sequence", // TransmitTransducerSequence
		"TransmitTransducerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4011),
		"Receive Transducer Sequence", // ReceiveTransducerSequence
		"ReceiveTransducerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4012),
		"Number of Elements", // NumberOfElements
		"NumberOfElements",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4013),
		"Element Shape", // ElementShape
		"ElementShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4014),
		"Element Dimension A", // ElementDimensionA
		"ElementDimensionA",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4015),
		"Element Dimension B", // ElementDimensionB
		"ElementDimensionB",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4016),
		"Element Pitch A", // ElementPitchA
		"ElementPitchA",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4017),
		"Measured Beam Dimension A", // MeasuredBeamDimensionA
		"MeasuredBeamDimensionA",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4018),
		"Measured Beam Dimension B", // MeasuredBeamDimensionB
		"MeasuredBeamDimensionB",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4019),
		"Location of Measured Beam Diameter", // LocationOfMeasuredBeamDiameter
		"LocationOfMeasuredBeamDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x401A),
		"Nominal Frequency", // NominalFrequency
		"NominalFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x401B),
		"Measured Center Frequency", // MeasuredCenterFrequency
		"MeasuredCenterFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x401C),
		"Measured Bandwidth", // MeasuredBandwidth
		"MeasuredBandwidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x401D),
		"Element Pitch B", // ElementPitchB
		"ElementPitchB",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4020),
		"Pulser Settings Sequence", // PulserSettingsSequence
		"PulserSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4022),
		"Pulse Width", // PulseWidth
		"PulseWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4024),
		"Excitation Frequency", // ExcitationFrequency
		"ExcitationFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4026),
		"Modulation Type", // ModulationType
		"ModulationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4028),
		"Damping", // Damping
		"Damping",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4030),
		"Receiver Settings Sequence", // ReceiverSettingsSequence
		"ReceiverSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4031),
		"Acquired Soundpath Length", // AcquiredSoundpathLength
		"AcquiredSoundpathLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4032),
		"Acquisition Compression Type", // AcquisitionCompressionType
		"AcquisitionCompressionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4033),
		"Acquisition Sample Size", // AcquisitionSampleSize
		"AcquisitionSampleSize",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4034),
		"Rectifier Smoothing", // RectifierSmoothing
		"RectifierSmoothing",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4035),
		"DAC Sequence", // DACSequence
		"DACSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4036),
		"DAC Type", // DACType
		"DACType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4038),
		"DAC Gain Points", // DACGainPoints
		"DACGainPoints",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x403A),
		"DAC Time Points", // DACTimePoints
		"DACTimePoints",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x403C),
		"DAC Amplitude", // DACAmplitude
		"DACAmplitude",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4040),
		"Pre-Amplifier Settings Sequence", // PreAmplifierSettingsSequence
		"PreAmplifierSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4050),
		"Transmit Transducer Settings Sequence", // TransmitTransducerSettingsSequence
		"TransmitTransducerSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4051),
		"Receive Transducer Settings Sequence", // ReceiveTransducerSettingsSequence
		"ReceiveTransducerSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4052),
		"Incident Angle", // IncidentAngle
		"IncidentAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4054),
		"Coupling Technique", // CouplingTechnique
		"CouplingTechnique",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4056),
		"Coupling Medium", // CouplingMedium
		"CouplingMedium",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4057),
		"Coupling Velocity", // CouplingVelocity
		"CouplingVelocity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4058),
		"Probe Center Location X", // ProbeCenterLocationX
		"ProbeCenterLocationX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4059),
		"Probe Center Location Z", // ProbeCenterLocationZ
		"ProbeCenterLocationZ",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x405A),
		"Sound Path Length", // SoundPathLength
		"SoundPathLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x405C),
		"Delay Law Identifier", // DelayLawIdentifier
		"DelayLawIdentifier",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4060),
		"Gate Settings Sequence", // GateSettingsSequence
		"GateSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4062),
		"Gate Threshold", // GateThreshold
		"GateThreshold",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4064),
		"Velocity of Sound", // VelocityOfSound
		"VelocityOfSound",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4070),
		"Calibration Settings Sequence", // CalibrationSettingsSequence
		"CalibrationSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4072),
		"Calibration Procedure", // CalibrationProcedure
		"CalibrationProcedure",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4074),
		"Procedure Version", // ProcedureVersion
		"ProcedureVersion",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4076),
		"Procedure Creation Date", // ProcedureCreationDate
		"ProcedureCreationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4078),
		"Procedure Expiration Date", // ProcedureExpirationDate
		"ProcedureExpirationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x407A),
		"Procedure Last Modified Date", // ProcedureLastModifiedDate
		"ProcedureLastModifiedDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x407C),
		"Calibration Time", // CalibrationTime
		"CalibrationTime",
		vm.VM1_n,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x407E),
		"Calibration Date", // CalibrationDate
		"CalibrationDate",
		vm.VM1_n,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4080),
		"Probe Drive Equipment Sequence", // ProbeDriveEquipmentSequence
		"ProbeDriveEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4081),
		"Drive Type", // DriveType
		"DriveType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4082),
		"Probe Drive Notes", // ProbeDriveNotes
		"ProbeDriveNotes",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4083),
		"Drive Probe Sequence", // DriveProbeSequence
		"DriveProbeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4084),
		"Probe Inductance", // ProbeInductance
		"ProbeInductance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4085),
		"Probe Resistance", // ProbeResistance
		"ProbeResistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4086),
		"Receive Probe Sequence", // ReceiveProbeSequence
		"ReceiveProbeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4087),
		"Probe Drive Settings Sequence", // ProbeDriveSettingsSequence
		"ProbeDriveSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4088),
		"Bridge Resistors", // BridgeResistors
		"BridgeResistors",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4089),
		"Probe Orientation Angle", // ProbeOrientationAngle
		"ProbeOrientationAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x408B),
		"User Selected Gain Y", // UserSelectedGainY
		"UserSelectedGainY",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x408C),
		"User Selected Phase", // UserSelectedPhase
		"UserSelectedPhase",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x408D),
		"User Selected Offset X", // UserSelectedOffsetX
		"UserSelectedOffsetX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x408E),
		"User Selected Offset Y", // UserSelectedOffsetY
		"UserSelectedOffsetY",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4091),
		"Channel Settings Sequence", // ChannelSettingsSequence
		"ChannelSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4092),
		"Channel Threshold", // ChannelThreshold
		"ChannelThreshold",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409A),
		"Scanner Settings Sequence", // ScannerSettingsSequence
		"ScannerSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409B),
		"Scan Procedure", // ScanProcedure
		"ScanProcedure",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409C),
		"Translation Rate X", // TranslationRateX
		"TranslationRateX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409D),
		"Translation Rate Y", // TranslationRateY
		"TranslationRateY",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x409F),
		"Channel Overlap", // ChannelOverlap
		"ChannelOverlap",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x40A0),
		"Image Quality Indicator Type", // ImageQualityIndicatorType
		"ImageQualityIndicatorType",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x40A1),
		"Image Quality Indicator Material", // ImageQualityIndicatorMaterial
		"ImageQualityIndicatorMaterial",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x40A2),
		"Image Quality Indicator Size", // ImageQualityIndicatorSize
		"ImageQualityIndicatorSize",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4101),
		"Wave Dimensions Definition Sequence", // WaveDimensionsDefinitionSequence
		"WaveDimensionsDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4102),
		"Wave Dimension Number", // WaveDimensionNumber
		"WaveDimensionNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4103),
		"Wave Dimension Description", // WaveDimensionDescription
		"WaveDimensionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4104),
		"Wave Dimension Unit", // WaveDimensionUnit
		"WaveDimensionUnit",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4105),
		"Wave Dimension Value Type", // WaveDimensionValueType
		"WaveDimensionValueType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4106),
		"Wave Dimension Values Sequence", // WaveDimensionValuesSequence
		"WaveDimensionValuesSequence",
		vm.VM1_n,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4107),
		"Referenced Wave Dimension", // ReferencedWaveDimension
		"ReferencedWaveDimension",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4108),
		"Integer Numeric Value", // IntegerNumericValue
		"IntegerNumericValue",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x4109),
		"Byte Numeric Value", // ByteNumericValue
		"ByteNumericValue",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x410A),
		"Short Numeric Value", // ShortNumericValue
		"ShortNumericValue",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x410B),
		"Single Precision Floating Point Numeric Value", // SinglePrecisionFloatingPointNumericValue
		"SinglePrecisionFloatingPointNumericValue",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x410C),
		"Double Precision Floating Point Numeric Value", // DoublePrecisionFloatingPointNumericValue
		"DoublePrecisionFloatingPointNumericValue",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5002),
		"LINAC Energy", // LINACEnergy
		"LINACEnergy",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5004),
		"LINAC Output", // LINACOutput
		"LINACOutput",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5100),
		"Active Aperture", // ActiveAperture
		"ActiveAperture",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5101),
		"Total Aperture", // TotalAperture
		"TotalAperture",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5102),
		"Aperture Elevation", // ApertureElevation
		"ApertureElevation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5103),
		"Main Lobe Angle", // MainLobeAngle
		"MainLobeAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5104),
		"Main Roof Angle", // MainRoofAngle
		"MainRoofAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5105),
		"Connector Type", // ConnectorType
		"ConnectorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5106),
		"Wedge Model Number", // WedgeModelNumber
		"WedgeModelNumber",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5107),
		"Wedge Angle Float", // WedgeAngleFloat
		"WedgeAngleFloat",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5108),
		"Wedge Roof Angle", // WedgeRoofAngle
		"WedgeRoofAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5109),
		"Wedge Element 1 Position", // WedgeElement1Position
		"WedgeElement1Position",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510A),
		"Wedge Material Velocity", // WedgeMaterialVelocity
		"WedgeMaterialVelocity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510B),
		"Wedge Material", // WedgeMaterial
		"WedgeMaterial",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510C),
		"Wedge Offset Z", // WedgeOffsetZ
		"WedgeOffsetZ",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510D),
		"Wedge Origin Offset X", // WedgeOriginOffsetX
		"WedgeOriginOffsetX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510E),
		"Wedge Time Delay", // WedgeTimeDelay
		"WedgeTimeDelay",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x510F),
		"Wedge Name", // WedgeName
		"WedgeName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5110),
		"Wedge Manufacturer Name", // WedgeManufacturerName
		"WedgeManufacturerName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5111),
		"Wedge Description", // WedgeDescription
		"WedgeDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5112),
		"Nominal Beam Angle", // NominalBeamAngle
		"NominalBeamAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5113),
		"Wedge Offset X", // WedgeOffsetX
		"WedgeOffsetX",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5114),
		"Wedge Offset Y", // WedgeOffsetY
		"WedgeOffsetY",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5115),
		"Wedge Total Length", // WedgeTotalLength
		"WedgeTotalLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5116),
		"Wedge In Contact Length", // WedgeInContactLength
		"WedgeInContactLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5117),
		"Wedge Front Gap", // WedgeFrontGap
		"WedgeFrontGap",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5118),
		"Wedge Total Height", // WedgeTotalHeight
		"WedgeTotalHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x5119),
		"Wedge Front Height", // WedgeFrontHeight
		"WedgeFrontHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511A),
		"Wedge Rear Height", // WedgeRearHeight
		"WedgeRearHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511B),
		"Wedge Total Width", // WedgeTotalWidth
		"WedgeTotalWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511C),
		"Wedge In Contact Width", // WedgeInContactWidth
		"WedgeInContactWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511D),
		"Wedge Chamfer Height", // WedgeChamferHeight
		"WedgeChamferHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511E),
		"Wedge Curve", // WedgeCurve
		"WedgeCurve",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x511F),
		"Radius Along the Wedge", // RadiusAlongWedge
		"RadiusAlongWedge",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6001),
		"Thermal Camera Settings Sequence", // ThermalCameraSettingsSequence
		"ThermalCameraSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6002),
		"Acquisition Frame Rate", // AcquisitionFrameRate
		"AcquisitionFrameRate",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6003),
		"Integration Time", // IntegrationTime
		"IntegrationTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6004),
		"Number of Calibration Frames", // NumberOfCalibrationFrames
		"NumberOfCalibrationFrames",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6005),
		"Number of Rows in Full Acquisition Image", // NumberOfRowsInFullAcquisitionImage
		"NumberOfRowsInFullAcquisitionImage",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6006),
		"Number Of Columns in Full Acquisition Image", // NumberOfColumnsInFullAcquisitionImage
		"NumberOfColumnsInFullAcquisitionImage",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6007),
		"Thermal Source Settings Sequence", // ThermalSourceSettingsSequence
		"ThermalSourceSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6008),
		"Source Horizontal Pitch", // SourceHorizontalPitch
		"SourceHorizontalPitch",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6009),
		"Source Vertical Pitch", // SourceVerticalPitch
		"SourceVerticalPitch",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600A),
		"Source Horizontal Scan Speed", // SourceHorizontalScanSpeed
		"SourceHorizontalScanSpeed",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600B),
		"Thermal Source Modulation Frequency", // ThermalSourceModulationFrequency
		"ThermalSourceModulationFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600C),
		"Induction Source Setting Sequence", // InductionSourceSettingSequence
		"InductionSourceSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600D),
		"Coil Frequency", // CoilFrequency
		"CoilFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600E),
		"Current Amplitude Across Coil", // CurrentAmplitudeAcrossCoil
		"CurrentAmplitudeAcrossCoil",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x600F),
		"Flash Source Setting Sequence", // FlashSourceSettingSequence
		"FlashSourceSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6010),
		"Flash Duration", // FlashDuration
		"FlashDuration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6011),
		"Flash Frame Number", // FlashFrameNumber
		"FlashFrameNumber",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6012),
		"Laser Source Setting Sequence", // LaserSourceSettingSequence
		"LaserSourceSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6013),
		"Horizontal Laser Spot Dimension", // HorizontalLaserSpotDimension
		"HorizontalLaserSpotDimension",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6014),
		"Vertical Laser Spot Dimension", // VerticalLaserSpotDimension
		"VerticalLaserSpotDimension",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6015),
		"Laser Wavelength", // LaserWavelength
		"LaserWavelength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6016),
		"Laser Power", // LaserPower
		"LaserPower",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6017),
		"Forced Gas Setting Sequence", // ForcedGasSettingSequence
		"ForcedGasSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6018),
		"Vibration Source Setting Sequence", // VibrationSourceSettingSequence
		"VibrationSourceSettingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6019),
		"Vibration Excitation Frequency", // VibrationExcitationFrequency
		"VibrationExcitationFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601A),
		"Vibration Excitation Voltage", // VibrationExcitationVoltage
		"VibrationExcitationVoltage",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601B),
		"Thermography Data Capture Method", // ThermographyDataCaptureMethod
		"ThermographyDataCaptureMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601C),
		"Thermal Technique", // ThermalTechnique
		"ThermalTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601D),
		"Thermal Camera Core Sequence", // ThermalCameraCoreSequence
		"ThermalCameraCoreSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601E),
		"Detector Wavelength Range", // DetectorWavelengthRange
		"DetectorWavelengthRange",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x601F),
		"Thermal Camera Calibration Type", // ThermalCameraCalibrationType
		"ThermalCameraCalibrationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6020),
		"Acquisition Image Counter", // AcquisitionImageCounter
		"AcquisitionImageCounter",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6021),
		"Front Panel Temperature", // FrontPanelTemperature
		"FrontPanelTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6022),
		"Air Gap Temperature", // AirGapTemperature
		"AirGapTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6023),
		"Vertical Pixel Size", // VerticalPixelSize
		"VerticalPixelSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6024),
		"Horizontal Pixel Size", // HorizontalPixelSize
		"HorizontalPixelSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6025),
		"Data Streaming Protocol", // DataStreamingProtocol
		"DataStreamingProtocol",
		vm.VM1_n,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6026),
		"Lens Sequence", // LensSequence
		"LensSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6027),
		"Field of View", // FieldOfView
		"FieldOfView",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6028),
		"Lens Filter Manufacturer", // LensFilterManufacturer
		"LensFilterManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6029),
		"Cutoff Filter Type", // CutoffFilterType
		"CutoffFilterType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602A),
		"Lens Filter Cut-Off Wavelength", // LensFilterCutOffWavelength
		"LensFilterCutOffWavelength",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602B),
		"Thermal Source Sequence", // ThermalSourceSequence
		"ThermalSourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602C),
		"Thermal Source Motion State", // ThermalSourceMotionState
		"ThermalSourceMotionState",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602D),
		"Thermal Source Motion Type", // ThermalSourceMotionType
		"ThermalSourceMotionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602E),
		"Induction Heating Sequence", // InductionHeatingSequence
		"InductionHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x602F),
		"Coil Configuration ID", // CoilConfigurationID
		"CoilConfigurationID",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6030),
		"Number of Turns in Coil", // NumberOfTurnsInCoil
		"NumberOfTurnsInCoil",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6031),
		"Shape of Individual Turn", // ShapeOfIndividualTurn
		"ShapeOfIndividualTurn",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6032),
		"Size of Individual Turn", // SizeOfIndividualTurn
		"SizeOfIndividualTurn",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6033),
		"Distance Between Turns", // DistanceBetweenTurns
		"DistanceBetweenTurns",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6034),
		"Flash Heating Sequence", // FlashHeatingSequence
		"FlashHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6035),
		"Number of Lamps", // NumberOfLamps
		"NumberOfLamps",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6036),
		"Flash Synchronization Protocol", // FlashSynchronizationProtocol
		"FlashSynchronizationProtocol",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6037),
		"Flash Modification Status", // FlashModificationStatus
		"FlashModificationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6038),
		"Laser Heating Sequence", // LaserHeatingSequence
		"LaserHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6039),
		"Laser Manufacturer", // LaserManufacturer
		"LaserManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603A),
		"Laser Model Number", // LaserModelNumber
		"LaserModelNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603B),
		"Laser Type Description", // LaserTypeDescription
		"LaserTypeDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603C),
		"Forced Gas Heating Sequence", // ForcedGasHeatingSequence
		"ForcedGasHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603D),
		"Gas Used for Heating/Cooling Part", // GasUsedForHeatingCoolingPart
		"GasUsedForHeatingCoolingPart",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603E),
		"Vibration/Sonic Heating Sequence", // VibrationSonicHeatingSequence
		"VibrationSonicHeatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x603F),
		"Probe Manufacturer", // ProbeManufacturer
		"ProbeManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6040),
		"Probe Model Number", // ProbeModelNumber
		"ProbeModelNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6041),
		"Aperture Size", // ApertureSize
		"ApertureSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6042),
		"Probe Resonant Frequency", // ProbeResonantFrequency
		"ProbeResonantFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6043),
		"Heat Source Description", // HeatSourceDescription
		"HeatSourceDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6044),
		"Surface Preparation with Optical Coating", // SurfacePreparationWithOpticalCoating
		"SurfacePreparationWithOpticalCoating",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6045),
		"Optical Coating Type", // OpticalCoatingType
		"OpticalCoatingType",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6046),
		"Thermal Conductivity of Exposed Surface", // ThermalConductivityOfExposedSurface
		"ThermalConductivityOfExposedSurface",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6047),
		"Material Density", // MaterialDensity
		"MaterialDensity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6048),
		"Specific Heat of Inspection Surface", // SpecificHeatOfInspectionSurface
		"SpecificHeatOfInspectionSurface",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6049),
		"Emissivity of Inspection Surface", // EmissivityOfInspectionSurface
		"EmissivityOfInspectionSurface",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604A),
		"Electromagnetic Classification of Inspection Surface", // ElectromagneticClassificationOfInspectionSurface
		"ElectromagneticClassificationOfInspectionSurface",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604C),
		"Moving Window Size", // MovingWindowSize
		"MovingWindowSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604D),
		"Moving Window Type", // MovingWindowType
		"MovingWindowType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604E),
		"Moving Window Weights", // MovingWindowWeights
		"MovingWindowWeights",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x604F),
		"Moving Window Pitch", // MovingWindowPitch
		"MovingWindowPitch",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6050),
		"Moving Window Padding Scheme", // MovingWindowPaddingScheme
		"MovingWindowPaddingScheme",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6051),
		"Moving Window Padding Sength", // MovingWindowPaddingLength
		"MovingWindowPaddingLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6052),
		"Spatial Filtering Parameters Sequence", // SpatialFilteringParametersSequence
		"SpatialFilteringParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6053),
		"Spatial Filtering Scheme", // SpatialFilteringScheme
		"SpatialFilteringScheme",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6056),
		"Horizontal Moving Window Size", // HorizontalMovingWindowSize
		"HorizontalMovingWindowSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6057),
		"Vertical Moving Window Size", // VerticalMovingWindowSize
		"VerticalMovingWindowSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6059),
		"Polynomial Fitting Sequence", // PolynomialFittingSequence
		"PolynomialFittingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605A),
		"Fitting Data Type", // FittingDataType
		"FittingDataType",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605B),
		"Operation on Time Axis Before Fitting", // OperationOnTimeAxisBeforeFitting
		"OperationOnTimeAxisBeforeFitting",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605C),
		"Operation on Pixel Intensity Before Fitting", // OperationOnPixelIntensityBeforeFitting
		"OperationOnPixelIntensityBeforeFitting",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605D),
		"Order of Polynomial", // OrderOfPolynomial
		"OrderOfPolynomial",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605E),
		"Independent Variable for Polynomial Fit", // IndependentVariableForPolynomialFit
		"IndependentVariableForPolynomialFit",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x605F),
		"PolynomialCoefficients", // PolynomialCoefficients
		"PolynomialCoefficients",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0014, 0x6060),
		"Thermography Pixel Data Unit", // ThermographyPixelDataUnit
		"ThermographyPixelDataUnit",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0001),
		"White Point", // WhitePoint
		"WhitePoint",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0002),
		"Primary Chromaticities", // PrimaryChromaticities
		"PrimaryChromaticities",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0003),
		"Battery Level", // BatteryLevel
		"BatteryLevel",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0004),
		"Exposure Time in Seconds", // ExposureTimeInSeconds
		"ExposureTimeInSeconds",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0005),
		"F-Number", // FNumber
		"FNumber",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0006),
		"OECF Rows", // OECFRows
		"OECFRows",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0007),
		"OECF Columns", // OECFColumns
		"OECFColumns",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0008),
		"OECF Column Names", // OECFColumnNames
		"OECFColumnNames",
		vm.VM1_n,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0009),
		"OECF Values", // OECFValues
		"OECFValues",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000A),
		"Spatial Frequency Response Rows", // SpatialFrequencyResponseRows
		"SpatialFrequencyResponseRows",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000B),
		"Spatial Frequency Response Columns", // SpatialFrequencyResponseColumns
		"SpatialFrequencyResponseColumns",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000C),
		"Spatial Frequency Response Column Names", // SpatialFrequencyResponseColumnNames
		"SpatialFrequencyResponseColumnNames",
		vm.VM1_n,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000D),
		"Spatial Frequency Response Values", // SpatialFrequencyResponseValues
		"SpatialFrequencyResponseValues",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000E),
		"Color Filter Array Pattern Rows", // ColorFilterArrayPatternRows
		"ColorFilterArrayPatternRows",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x000F),
		"Color Filter Array Pattern Columns", // ColorFilterArrayPatternColumns
		"ColorFilterArrayPatternColumns",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0010),
		"Color Filter Array Pattern Values", // ColorFilterArrayPatternValues
		"ColorFilterArrayPatternValues",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0011),
		"Flash Firing Status", // FlashFiringStatus
		"FlashFiringStatus",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0012),
		"Flash Return Status", // FlashReturnStatus
		"FlashReturnStatus",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0013),
		"Flash Mode", // FlashMode
		"FlashMode",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0014),
		"Flash Function Present", // FlashFunctionPresent
		"FlashFunctionPresent",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0015),
		"Flash Red Eye Mode", // FlashRedEyeMode
		"FlashRedEyeMode",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0016),
		"Exposure Program", // ExposureProgram
		"ExposureProgram",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0017),
		"Spectral Sensitivity", // SpectralSensitivity
		"SpectralSensitivity",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0018),
		"Photographic Sensitivity", // PhotographicSensitivity
		"PhotographicSensitivity",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0019),
		"Self Timer Mode", // SelfTimerMode
		"SelfTimerMode",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001A),
		"Sensitivity Type", // SensitivityType
		"SensitivityType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001B),
		"Standard Output Sensitivity", // StandardOutputSensitivity
		"StandardOutputSensitivity",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001C),
		"Recommended Exposure Index", // RecommendedExposureIndex
		"RecommendedExposureIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001D),
		"ISO Speed", // ISOSpeed
		"ISOSpeed",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001E),
		"ISO Speed Latitude yyy", // ISOSpeedLatitudeyyy
		"ISOSpeedLatitudeyyy",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x001F),
		"ISO Speed Latitude zzz", // ISOSpeedLatitudezzz
		"ISOSpeedLatitudezzz",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0020),
		"EXIF Version", // EXIFVersion
		"EXIFVersion",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0021),
		"Shutter Speed Value", // ShutterSpeedValue
		"ShutterSpeedValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0022),
		"Aperture Value", // ApertureValue
		"ApertureValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0023),
		"Brightness Value", // BrightnessValue
		"BrightnessValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0024),
		"Exposure Bias Value", // ExposureBiasValue
		"ExposureBiasValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0025),
		"Max Aperture Value", // MaxApertureValue
		"MaxApertureValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0026),
		"Subject Distance", // SubjectDistance
		"SubjectDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0027),
		"Metering Mode", // MeteringMode
		"MeteringMode",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0028),
		"Light Source", // LightSource
		"LightSource",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0029),
		"Focal Length", // FocalLength
		"FocalLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x002A),
		"Subject Area", // SubjectArea
		"SubjectArea",
		vm.MustParse("2-4"),
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x002B),
		"Maker Note", // MakerNote
		"MakerNote",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0030),
		"Temperature", // Temperature
		"Temperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0031),
		"Humidity", // Humidity
		"Humidity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0032),
		"Pressure", // Pressure
		"Pressure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0033),
		"Water Depth", // WaterDepth
		"WaterDepth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0034),
		"Acceleration", // Acceleration
		"Acceleration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0035),
		"Camera Elevation Angle", // CameraElevationAngle
		"CameraElevationAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0036),
		"Flash Energy", // FlashEnergy
		"FlashEnergy",
		vm.VM1_2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0037),
		"Subject Location", // SubjectLocation
		"SubjectLocation",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0038),
		"Photographic Exposure Index", // PhotographicExposureIndex
		"PhotographicExposureIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0039),
		"Sensing Method", // SensingMethod
		"SensingMethod",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x003A),
		"File Source", // FileSource
		"FileSource",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x003B),
		"Scene Type", // SceneType
		"SceneType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0041),
		"Custom Rendered", // CustomRendered
		"CustomRendered",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0042),
		"Exposure Mode", // ExposureMode
		"ExposureMode",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0043),
		"White Balance", // WhiteBalance
		"WhiteBalance",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0044),
		"Digital Zoom Ratio", // DigitalZoomRatio
		"DigitalZoomRatio",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0045),
		"Focal Length In 35mm Film", // FocalLengthIn35mmFilm
		"FocalLengthIn35mmFilm",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0046),
		"Scene Capture Type", // SceneCaptureType
		"SceneCaptureType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0047),
		"Gain Control", // GainControl
		"GainControl",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0048),
		"Contrast", // Contrast
		"Contrast",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0049),
		"Saturation", // Saturation
		"Saturation",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004A),
		"Sharpness", // Sharpness
		"Sharpness",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004B),
		"Device Setting Description", // DeviceSettingDescription
		"DeviceSettingDescription",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004C),
		"Subject Distance Range", // SubjectDistanceRange
		"SubjectDistanceRange",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004D),
		"Camera Owner Name", // CameraOwnerName
		"CameraOwnerName",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004E),
		"Lens Specification", // LensSpecification
		"LensSpecification",
		vm.VM4,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x004F),
		"Lens Make", // LensMake
		"LensMake",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0050),
		"Lens Model", // LensModel
		"LensModel",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0051),
		"Lens Serial Number", // LensSerialNumber
		"LensSerialNumber",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0061),
		"Interoperability Index", // InteroperabilityIndex
		"InteroperabilityIndex",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0062),
		"Interoperability Version", // InteroperabilityVersion
		"InteroperabilityVersion",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0070),
		"GPS Version ID", // GPSVersionID
		"GPSVersionID",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0071),
		"GPS Latitude Ref", // GPSLatitudeRef
		"GPSLatitudeRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0072),
		"GPS Latitude", // GPSLatitude
		"GPSLatitude",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0073),
		"GPS Longitude Ref", // GPSLongitudeRef
		"GPSLongitudeRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0074),
		"GPS Longitude", // GPSLongitude
		"GPSLongitude",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0075),
		"GPS Altitude Ref", // GPSAltitudeRef
		"GPSAltitudeRef",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0076),
		"GPS Altitude", // GPSAltitude
		"GPSAltitude",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0077),
		"GPS Time Stamp", // GPSTimeStamp
		"GPSTimeStamp",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0078),
		"GPS Satellites", // GPSSatellites
		"GPSSatellites",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0079),
		"GPS Status", // GPSStatus
		"GPSStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007A),
		"GPS Measure Mode", // GPSMeasureMode
		"GPSMeasureMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007B),
		"GPS DOP", // GPSDOP
		"GPSDOP",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007C),
		"GPS Speed Ref", // GPSSpeedRef
		"GPSSpeedRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007D),
		"GPS Speed", // GPSSpeed
		"GPSSpeed",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007E),
		"GPS Track Ref", // GPSTrackRef
		"GPSTrackRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x007F),
		"GPS Track", // GPSTrack
		"GPSTrack",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0080),
		"GPS Img Direction Ref", // GPSImgDirectionRef
		"GPSImgDirectionRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0081),
		"GPS Img Direction", // GPSImgDirection
		"GPSImgDirection",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0082),
		"GPS Map Datum", // GPSMapDatum
		"GPSMapDatum",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0083),
		"GPS Dest Latitude Ref", // GPSDestLatitudeRef
		"GPSDestLatitudeRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0084),
		"GPS Dest Latitude", // GPSDestLatitude
		"GPSDestLatitude",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0085),
		"GPS Dest Longitude Ref", // GPSDestLongitudeRef
		"GPSDestLongitudeRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0086),
		"GPS Dest Longitude", // GPSDestLongitude
		"GPSDestLongitude",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0087),
		"GPS Dest Bearing Ref", // GPSDestBearingRef
		"GPSDestBearingRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0088),
		"GPS Dest Bearing", // GPSDestBearing
		"GPSDestBearing",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x0089),
		"GPS Dest Distance Ref", // GPSDestDistanceRef
		"GPSDestDistanceRef",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008A),
		"GPS Dest Distance", // GPSDestDistance
		"GPSDestDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008B),
		"GPS Processing Method", // GPSProcessingMethod
		"GPSProcessingMethod",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008C),
		"GPS Area Information", // GPSAreaInformation
		"GPSAreaInformation",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008D),
		"GPS Date Stamp", // GPSDateStamp
		"GPSDateStamp",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x008E),
		"GPS Differential", // GPSDifferential
		"GPSDifferential",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1001),
		"Light Source Polarization", // LightSourcePolarization
		"LightSourcePolarization",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1002),
		"Emitter Color Temperature", // EmitterColorTemperature
		"EmitterColorTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1003),
		"Contact Method", // ContactMethod
		"ContactMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1004),
		"Immersion Media", // ImmersionMedia
		"ImmersionMedia",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0016, 0x1005),
		"Optical Magnification Factor", // OpticalMagnificationFactor
		"OpticalMagnificationFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0010),
		"Contrast/Bolus Agent", // ContrastBolusAgent
		"ContrastBolusAgent",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0012),
		"Contrast/Bolus Agent Sequence", // ContrastBolusAgentSequence
		"ContrastBolusAgentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0013),
		"Contrast/Bolus T1 Relaxivity", // ContrastBolusT1Relaxivity
		"ContrastBolusT1Relaxivity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0014),
		"Contrast/Bolus Administration Route Sequence", // ContrastBolusAdministrationRouteSequence
		"ContrastBolusAdministrationRouteSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0015),
		"Body Part Examined", // BodyPartExamined
		"BodyPartExamined",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0020),
		"Scanning Sequence", // ScanningSequence
		"ScanningSequence",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0021),
		"Sequence Variant", // SequenceVariant
		"SequenceVariant",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0022),
		"Scan Options", // ScanOptions
		"ScanOptions",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0023),
		"MR Acquisition Type", // MRAcquisitionType
		"MRAcquisitionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0024),
		"Sequence Name", // SequenceName
		"SequenceName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0025),
		"Angio Flag", // AngioFlag
		"AngioFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0026),
		"Intervention Drug Information Sequence", // InterventionDrugInformationSequence
		"InterventionDrugInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0027),
		"Intervention Drug Stop Time", // InterventionDrugStopTime
		"InterventionDrugStopTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0028),
		"Intervention Drug Dose", // InterventionDrugDose
		"InterventionDrugDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0029),
		"Intervention Drug Code Sequence", // InterventionDrugCodeSequence
		"InterventionDrugCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x002A),
		"Additional Drug Sequence", // AdditionalDrugSequence
		"AdditionalDrugSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0030),
		"Radionuclide", // Radionuclide
		"Radionuclide",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0031),
		"Radiopharmaceutical", // Radiopharmaceutical
		"Radiopharmaceutical",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0032),
		"Energy Window Centerline", // EnergyWindowCenterline
		"EnergyWindowCenterline",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0033),
		"Energy Window Total Width", // EnergyWindowTotalWidth
		"EnergyWindowTotalWidth",
		vm.VM1_n,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0034),
		"Intervention Drug Name", // InterventionDrugName
		"InterventionDrugName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0035),
		"Intervention Drug Start Time", // InterventionDrugStartTime
		"InterventionDrugStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0036),
		"Intervention Sequence", // InterventionSequence
		"InterventionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0037),
		"Therapy Type", // TherapyType
		"TherapyType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0038),
		"Intervention Status", // InterventionStatus
		"InterventionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0039),
		"Therapy Description", // TherapyDescription
		"TherapyDescription",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x003A),
		"Intervention Description", // InterventionDescription
		"InterventionDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0040),
		"Cine Rate", // CineRate
		"CineRate",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0042),
		"Initial Cine Run State", // InitialCineRunState
		"InitialCineRunState",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0050),
		"Slice Thickness", // SliceThickness
		"SliceThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0060),
		"KVP", // KVP
		"KVP",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0070),
		"Counts Accumulated", // CountsAccumulated
		"CountsAccumulated",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0071),
		"Acquisition Termination Condition", // AcquisitionTerminationCondition
		"AcquisitionTerminationCondition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0072),
		"Effective Duration", // EffectiveDuration
		"EffectiveDuration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0073),
		"Acquisition Start Condition", // AcquisitionStartCondition
		"AcquisitionStartCondition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0074),
		"Acquisition Start Condition Data", // AcquisitionStartConditionData
		"AcquisitionStartConditionData",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0075),
		"Acquisition Termination Condition Data", // AcquisitionTerminationConditionData
		"AcquisitionTerminationConditionData",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0080),
		"Repetition Time", // RepetitionTime
		"RepetitionTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0081),
		"Echo Time", // EchoTime
		"EchoTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0082),
		"Inversion Time", // InversionTime
		"InversionTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0083),
		"Number of Averages", // NumberOfAverages
		"NumberOfAverages",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0084),
		"Imaging Frequency", // ImagingFrequency
		"ImagingFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0085),
		"Imaged Nucleus", // ImagedNucleus
		"ImagedNucleus",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0086),
		"Echo Number(s)", // EchoNumbers
		"EchoNumbers",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0087),
		"Magnetic Field Strength", // MagneticFieldStrength
		"MagneticFieldStrength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0088),
		"Spacing Between Slices", // SpacingBetweenSlices
		"SpacingBetweenSlices",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0089),
		"Number of Phase Encoding Steps", // NumberOfPhaseEncodingSteps
		"NumberOfPhaseEncodingSteps",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0090),
		"Data Collection Diameter", // DataCollectionDiameter
		"DataCollectionDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0091),
		"Echo Train Length", // EchoTrainLength
		"EchoTrainLength",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0093),
		"Percent Sampling", // PercentSampling
		"PercentSampling",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0094),
		"Percent Phase Field of View", // PercentPhaseFieldOfView
		"PercentPhaseFieldOfView",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x0095),
		"Pixel Bandwidth", // PixelBandwidth
		"PixelBandwidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1000),
		"Device Serial Number", // DeviceSerialNumber
		"DeviceSerialNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1002),
		"Device UID", // DeviceUID
		"DeviceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1003),
		"Device ID", // DeviceID
		"DeviceID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1004),
		"Plate ID", // PlateID
		"PlateID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1005),
		"Generator ID", // GeneratorID
		"GeneratorID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1006),
		"Grid ID", // GridID
		"GridID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1007),
		"Cassette ID", // CassetteID
		"CassetteID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1008),
		"Gantry ID", // GantryID
		"GantryID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1009),
		"Unique Device Identifier", // UniqueDeviceIdentifier
		"UniqueDeviceIdentifier",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x100A),
		"UDI Sequence", // UDISequence
		"UDISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x100B),
		"Manufacturer's Device Class UID", // ManufacturerDeviceClassUID
		"ManufacturerDeviceClassUID",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1010),
		"Secondary Capture Device ID", // SecondaryCaptureDeviceID
		"SecondaryCaptureDeviceID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1011),
		"Hardcopy Creation Device ID", // HardcopyCreationDeviceID
		"HardcopyCreationDeviceID",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1012),
		"Date of Secondary Capture", // DateOfSecondaryCapture
		"DateOfSecondaryCapture",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1014),
		"Time of Secondary Capture", // TimeOfSecondaryCapture
		"TimeOfSecondaryCapture",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1016),
		"Secondary Capture Device Manufacturer", // SecondaryCaptureDeviceManufacturer
		"SecondaryCaptureDeviceManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1017),
		"Hardcopy Device Manufacturer", // HardcopyDeviceManufacturer
		"HardcopyDeviceManufacturer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1018),
		"Secondary Capture Device Manufacturer's Model Name", // SecondaryCaptureDeviceManufacturerModelName
		"SecondaryCaptureDeviceManufacturerModelName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1019),
		"Secondary Capture Device Software Versions", // SecondaryCaptureDeviceSoftwareVersions
		"SecondaryCaptureDeviceSoftwareVersions",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x101A),
		"Hardcopy Device Software Version", // HardcopyDeviceSoftwareVersion
		"HardcopyDeviceSoftwareVersion",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x101B),
		"Hardcopy Device Manufacturer's Model Name", // HardcopyDeviceManufacturerModelName
		"HardcopyDeviceManufacturerModelName",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1020),
		"Software Versions", // SoftwareVersions
		"SoftwareVersions",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1022),
		"Video Image Format Acquired", // VideoImageFormatAcquired
		"VideoImageFormatAcquired",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1023),
		"Digital Image Format Acquired", // DigitalImageFormatAcquired
		"DigitalImageFormatAcquired",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1030),
		"Protocol Name", // ProtocolName
		"ProtocolName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1040),
		"Contrast/Bolus Route", // ContrastBolusRoute
		"ContrastBolusRoute",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1041),
		"Contrast/Bolus Volume", // ContrastBolusVolume
		"ContrastBolusVolume",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1042),
		"Contrast/Bolus Start Time", // ContrastBolusStartTime
		"ContrastBolusStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1043),
		"Contrast/Bolus Stop Time", // ContrastBolusStopTime
		"ContrastBolusStopTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1044),
		"Contrast/Bolus Total Dose", // ContrastBolusTotalDose
		"ContrastBolusTotalDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1045),
		"Syringe Counts", // SyringeCounts
		"SyringeCounts",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1046),
		"Contrast Flow Rate", // ContrastFlowRate
		"ContrastFlowRate",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1047),
		"Contrast Flow Duration", // ContrastFlowDuration
		"ContrastFlowDuration",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1048),
		"Contrast/Bolus Ingredient", // ContrastBolusIngredient
		"ContrastBolusIngredient",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1049),
		"Contrast/Bolus Ingredient Concentration", // ContrastBolusIngredientConcentration
		"ContrastBolusIngredientConcentration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1050),
		"Spatial Resolution", // SpatialResolution
		"SpatialResolution",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1060),
		"Trigger Time", // TriggerTime
		"TriggerTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1061),
		"Trigger Source or Type", // TriggerSourceOrType
		"TriggerSourceOrType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1062),
		"Nominal Interval", // NominalInterval
		"NominalInterval",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1063),
		"Frame Time", // FrameTime
		"FrameTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1064),
		"Cardiac Framing Type", // CardiacFramingType
		"CardiacFramingType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1065),
		"Frame Time Vector", // FrameTimeVector
		"FrameTimeVector",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1066),
		"Frame Delay", // FrameDelay
		"FrameDelay",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1067),
		"Image Trigger Delay", // ImageTriggerDelay
		"ImageTriggerDelay",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1068),
		"Multiplex Group Time Offset", // MultiplexGroupTimeOffset
		"MultiplexGroupTimeOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1069),
		"Trigger Time Offset", // TriggerTimeOffset
		"TriggerTimeOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x106A),
		"Synchronization Trigger", // SynchronizationTrigger
		"SynchronizationTrigger",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x106C),
		"Synchronization Channel", // SynchronizationChannel
		"SynchronizationChannel",
		vm.VM2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x106E),
		"Trigger Sample Position", // TriggerSamplePosition
		"TriggerSamplePosition",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1070),
		"Radiopharmaceutical Route", // RadiopharmaceuticalRoute
		"RadiopharmaceuticalRoute",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1071),
		"Radiopharmaceutical Volume", // RadiopharmaceuticalVolume
		"RadiopharmaceuticalVolume",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1072),
		"Radiopharmaceutical Start Time", // RadiopharmaceuticalStartTime
		"RadiopharmaceuticalStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1073),
		"Radiopharmaceutical Stop Time", // RadiopharmaceuticalStopTime
		"RadiopharmaceuticalStopTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1074),
		"Radionuclide Total Dose", // RadionuclideTotalDose
		"RadionuclideTotalDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1075),
		"Radionuclide Half Life", // RadionuclideHalfLife
		"RadionuclideHalfLife",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1076),
		"Radionuclide Positron Fraction", // RadionuclidePositronFraction
		"RadionuclidePositronFraction",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1077),
		"Radiopharmaceutical Specific Activity", // RadiopharmaceuticalSpecificActivity
		"RadiopharmaceuticalSpecificActivity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1078),
		"Radiopharmaceutical Start DateTime", // RadiopharmaceuticalStartDateTime
		"RadiopharmaceuticalStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1079),
		"Radiopharmaceutical Stop DateTime", // RadiopharmaceuticalStopDateTime
		"RadiopharmaceuticalStopDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1080),
		"Beat Rejection Flag", // BeatRejectionFlag
		"BeatRejectionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1081),
		"Low R-R Value", // LowRRValue
		"LowRRValue",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1082),
		"High R-R Value", // HighRRValue
		"HighRRValue",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1083),
		"Intervals Acquired", // IntervalsAcquired
		"IntervalsAcquired",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1084),
		"Intervals Rejected", // IntervalsRejected
		"IntervalsRejected",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1085),
		"PVC Rejection", // PVCRejection
		"PVCRejection",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1086),
		"Skip Beats", // SkipBeats
		"SkipBeats",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1088),
		"Heart Rate", // HeartRate
		"HeartRate",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1090),
		"Cardiac Number of Images", // CardiacNumberOfImages
		"CardiacNumberOfImages",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1094),
		"Trigger Window", // TriggerWindow
		"TriggerWindow",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1100),
		"Reconstruction Diameter", // ReconstructionDiameter
		"ReconstructionDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1110),
		"Distance Source to Detector", // DistanceSourceToDetector
		"DistanceSourceToDetector",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1111),
		"Distance Source to Patient", // DistanceSourceToPatient
		"DistanceSourceToPatient",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1114),
		"Estimated Radiographic Magnification Factor", // EstimatedRadiographicMagnificationFactor
		"EstimatedRadiographicMagnificationFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1120),
		"Gantry/Detector Tilt", // GantryDetectorTilt
		"GantryDetectorTilt",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1121),
		"Gantry/Detector Slew", // GantryDetectorSlew
		"GantryDetectorSlew",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1130),
		"Table Height", // TableHeight
		"TableHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1131),
		"Table Traverse", // TableTraverse
		"TableTraverse",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1134),
		"Table Motion", // TableMotion
		"TableMotion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1135),
		"Table Vertical Increment", // TableVerticalIncrement
		"TableVerticalIncrement",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1136),
		"Table Lateral Increment", // TableLateralIncrement
		"TableLateralIncrement",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1137),
		"Table Longitudinal Increment", // TableLongitudinalIncrement
		"TableLongitudinalIncrement",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1138),
		"Table Angle", // TableAngle
		"TableAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x113A),
		"Table Type", // TableType
		"TableType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1140),
		"Rotation Direction", // RotationDirection
		"RotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1141),
		"Angular Position", // AngularPosition
		"AngularPosition",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1142),
		"Radial Position", // RadialPosition
		"RadialPosition",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1143),
		"Scan Arc", // ScanArc
		"ScanArc",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1144),
		"Angular Step", // AngularStep
		"AngularStep",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1145),
		"Center of Rotation Offset", // CenterOfRotationOffset
		"CenterOfRotationOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1146),
		"Rotation Offset", // RotationOffset
		"RotationOffset",
		vm.VM1_n,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1147),
		"Field of View Shape", // FieldOfViewShape
		"FieldOfViewShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1149),
		"Field of View Dimension(s)", // FieldOfViewDimensions
		"FieldOfViewDimensions",
		vm.VM1_2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1150),
		"Exposure Time", // ExposureTime
		"ExposureTime",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1151),
		"X-Ray Tube Current", // XRayTubeCurrent
		"XRayTubeCurrent",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1152),
		"Exposure", // Exposure
		"Exposure",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1153),
		"Exposure in µAs", // ExposureInuAs
		"ExposureInuAs",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1154),
		"Average Pulse Width", // AveragePulseWidth
		"AveragePulseWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1155),
		"Radiation Setting", // RadiationSetting
		"RadiationSetting",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1156),
		"Rectification Type", // RectificationType
		"RectificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x115A),
		"Radiation Mode", // RadiationMode
		"RadiationMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x115E),
		"Image and Fluoroscopy Area Dose Product", // ImageAndFluoroscopyAreaDoseProduct
		"ImageAndFluoroscopyAreaDoseProduct",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1160),
		"Filter Type", // FilterType
		"FilterType",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1161),
		"Type of Filters", // TypeOfFilters
		"TypeOfFilters",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1162),
		"Intensifier Size", // IntensifierSize
		"IntensifierSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1164),
		"Imager Pixel Spacing", // ImagerPixelSpacing
		"ImagerPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1166),
		"Grid", // Grid
		"Grid",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1170),
		"Generator Power", // GeneratorPower
		"GeneratorPower",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1180),
		"Collimator/grid Name", // CollimatorGridName
		"CollimatorGridName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1181),
		"Collimator Type", // CollimatorType
		"CollimatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1182),
		"Focal Distance", // FocalDistance
		"FocalDistance",
		vm.VM1_2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1183),
		"X Focus Center", // XFocusCenter
		"XFocusCenter",
		vm.VM1_2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1184),
		"Y Focus Center", // YFocusCenter
		"YFocusCenter",
		vm.VM1_2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1190),
		"Focal Spot(s)", // FocalSpots
		"FocalSpots",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1191),
		"Anode Target Material", // AnodeTargetMaterial
		"AnodeTargetMaterial",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A0),
		"Body Part Thickness", // BodyPartThickness
		"BodyPartThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A2),
		"Compression Force", // CompressionForce
		"CompressionForce",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A3),
		"Compression Pressure", // CompressionPressure
		"CompressionPressure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A4),
		"Paddle Description", // PaddleDescription
		"PaddleDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11A5),
		"Compression Contact Area", // CompressionContactArea
		"CompressionContactArea",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B0),
		"Acquisition Mode", // AcquisitionMode
		"AcquisitionMode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B1),
		"Dose Mode Name", // DoseModeName
		"DoseModeName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B2),
		"Acquired Subtraction Mask Flag", // AcquiredSubtractionMaskFlag
		"AcquiredSubtractionMaskFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B3),
		"Fluoroscopy Persistence Flag", // FluoroscopyPersistenceFlag
		"FluoroscopyPersistenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B4),
		"Fluoroscopy Last Image Hold Persistence Flag", // FluoroscopyLastImageHoldPersistenceFlag
		"FluoroscopyLastImageHoldPersistenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B5),
		"Upper Limit Number Of Persistent Fluoroscopy Frames", // UpperLimitNumberOfPersistentFluoroscopyFrames
		"UpperLimitNumberOfPersistentFluoroscopyFrames",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B6),
		"Contrast/Bolus Auto Injection Trigger Flag", // ContrastBolusAutoInjectionTriggerFlag
		"ContrastBolusAutoInjectionTriggerFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B7),
		"Contrast/Bolus Injection Delay", // ContrastBolusInjectionDelay
		"ContrastBolusInjectionDelay",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B8),
		"XA Acquisition Phase Details Sequence", // XAAcquisitionPhaseDetailsSequence
		"XAAcquisitionPhaseDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11B9),
		"XA Acquisition Frame Rate", // XAAcquisitionFrameRate
		"XAAcquisitionFrameRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BA),
		"XA Plane Details Sequence", // XAPlaneDetailsSequence
		"XAPlaneDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BB),
		"Acquisition Field of View Label", // AcquisitionFieldOfViewLabel
		"AcquisitionFieldOfViewLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BC),
		"X-Ray Filter Details Sequence", // XRayFilterDetailsSequence
		"XRayFilterDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BD),
		"XA Acquisition Duration", // XAAcquisitionDuration
		"XAAcquisitionDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BE),
		"Reconstruction Pipeline Type", // ReconstructionPipelineType
		"ReconstructionPipelineType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11BF),
		"Image Filter Details Sequence", // ImageFilterDetailsSequence
		"ImageFilterDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11C0),
		"Applied Mask Subtraction Flag", // AppliedMaskSubtractionFlag
		"AppliedMaskSubtractionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x11C1),
		"Requested Series Description Code Sequence", // RequestedSeriesDescriptionCodeSequence
		"RequestedSeriesDescriptionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1200),
		"Date of Last Calibration", // DateOfLastCalibration
		"DateOfLastCalibration",
		vm.VM1_n,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1201),
		"Time of Last Calibration", // TimeOfLastCalibration
		"TimeOfLastCalibration",
		vm.VM1_n,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1202),
		"DateTime of Last Calibration", // DateTimeOfLastCalibration
		"DateTimeOfLastCalibration",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1203),
		"Calibration DateTime", // CalibrationDateTime
		"CalibrationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1204),
		"Date of Manufacture", // DateOfManufacture
		"DateOfManufacture",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1205),
		"Date of Installation", // DateOfInstallation
		"DateOfInstallation",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1210),
		"Convolution Kernel", // ConvolutionKernel
		"ConvolutionKernel",
		vm.VM1_n,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1240),
		"Upper/Lower Pixel Values", // UpperLowerPixelValues
		"UpperLowerPixelValues",
		vm.VM1_n,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1242),
		"Actual Frame Duration", // ActualFrameDuration
		"ActualFrameDuration",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1243),
		"Count Rate", // CountRate
		"CountRate",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1244),
		"Preferred Playback Sequencing", // PreferredPlaybackSequencing
		"PreferredPlaybackSequencing",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1250),
		"Receive Coil Name", // ReceiveCoilName
		"ReceiveCoilName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1251),
		"Transmit Coil Name", // TransmitCoilName
		"TransmitCoilName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1260),
		"Plate Type", // PlateType
		"PlateType",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1261),
		"Phosphor Type", // PhosphorType
		"PhosphorType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1271),
		"Water Equivalent Diameter", // WaterEquivalentDiameter
		"WaterEquivalentDiameter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1272),
		"Water Equivalent Diameter Calculation Method Code Sequence", // WaterEquivalentDiameterCalculationMethodCodeSequence
		"WaterEquivalentDiameterCalculationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1300),
		"Scan Velocity", // ScanVelocity
		"ScanVelocity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1301),
		"Whole Body Technique", // WholeBodyTechnique
		"WholeBodyTechnique",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1302),
		"Scan Length", // ScanLength
		"ScanLength",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1310),
		"Acquisition Matrix", // AcquisitionMatrix
		"AcquisitionMatrix",
		vm.VM4,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1312),
		"In-plane Phase Encoding Direction", // InPlanePhaseEncodingDirection
		"InPlanePhaseEncodingDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1314),
		"Flip Angle", // FlipAngle
		"FlipAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1315),
		"Variable Flip Angle Flag", // VariableFlipAngleFlag
		"VariableFlipAngleFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1316),
		"SAR", // SAR
		"SAR",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1318),
		"dB/dt", // dBdt
		"dBdt",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1320),
		"B1rms", // B1rms
		"B1rms",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1400),
		"Acquisition Device Processing Description", // AcquisitionDeviceProcessingDescription
		"AcquisitionDeviceProcessingDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1401),
		"Acquisition Device Processing Code", // AcquisitionDeviceProcessingCode
		"AcquisitionDeviceProcessingCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1402),
		"Cassette Orientation", // CassetteOrientation
		"CassetteOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1403),
		"Cassette Size", // CassetteSize
		"CassetteSize",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1404),
		"Exposures on Plate", // ExposuresOnPlate
		"ExposuresOnPlate",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1405),
		"Relative X-Ray Exposure", // RelativeXRayExposure
		"RelativeXRayExposure",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1411),
		"Exposure Index", // ExposureIndex
		"ExposureIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1412),
		"Target Exposure Index", // TargetExposureIndex
		"TargetExposureIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1413),
		"Deviation Index", // DeviationIndex
		"DeviationIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1450),
		"Column Angulation", // ColumnAngulation
		"ColumnAngulation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1460),
		"Tomo Layer Height", // TomoLayerHeight
		"TomoLayerHeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1470),
		"Tomo Angle", // TomoAngle
		"TomoAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1480),
		"Tomo Time", // TomoTime
		"TomoTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1490),
		"Tomo Type", // TomoType
		"TomoType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1491),
		"Tomo Class", // TomoClass
		"TomoClass",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1495),
		"Number of Tomosynthesis Source Images", // NumberOfTomosynthesisSourceImages
		"NumberOfTomosynthesisSourceImages",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1500),
		"Positioner Motion", // PositionerMotion
		"PositionerMotion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1508),
		"Positioner Type", // PositionerType
		"PositionerType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1510),
		"Positioner Primary Angle", // PositionerPrimaryAngle
		"PositionerPrimaryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1511),
		"Positioner Secondary Angle", // PositionerSecondaryAngle
		"PositionerSecondaryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1520),
		"Positioner Primary Angle Increment", // PositionerPrimaryAngleIncrement
		"PositionerPrimaryAngleIncrement",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1521),
		"Positioner Secondary Angle Increment", // PositionerSecondaryAngleIncrement
		"PositionerSecondaryAngleIncrement",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1530),
		"Detector Primary Angle", // DetectorPrimaryAngle
		"DetectorPrimaryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1531),
		"Detector Secondary Angle", // DetectorSecondaryAngle
		"DetectorSecondaryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1600),
		"Shutter Shape", // ShutterShape
		"ShutterShape",
		vm.VM1_3,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1602),
		"Shutter Left Vertical Edge", // ShutterLeftVerticalEdge
		"ShutterLeftVerticalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1604),
		"Shutter Right Vertical Edge", // ShutterRightVerticalEdge
		"ShutterRightVerticalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1606),
		"Shutter Upper Horizontal Edge", // ShutterUpperHorizontalEdge
		"ShutterUpperHorizontalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1608),
		"Shutter Lower Horizontal Edge", // ShutterLowerHorizontalEdge
		"ShutterLowerHorizontalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1610),
		"Center of Circular Shutter", // CenterOfCircularShutter
		"CenterOfCircularShutter",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1612),
		"Radius of Circular Shutter", // RadiusOfCircularShutter
		"RadiusOfCircularShutter",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1620),
		"Vertices of the Polygonal Shutter", // VerticesOfThePolygonalShutter
		"VerticesOfThePolygonalShutter",
		vm.VM2_2n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1622),
		"Shutter Presentation Value", // ShutterPresentationValue
		"ShutterPresentationValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1623),
		"Shutter Overlay Group", // ShutterOverlayGroup
		"ShutterOverlayGroup",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1624),
		"Shutter Presentation Color CIELab Value", // ShutterPresentationColorCIELabValue
		"ShutterPresentationColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1630),
		"Outline Shape Type", // OutlineShapeType
		"OutlineShapeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1631),
		"Outline Left Vertical Edge", // OutlineLeftVerticalEdge
		"OutlineLeftVerticalEdge",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1632),
		"Outline Right Vertical Edge", // OutlineRightVerticalEdge
		"OutlineRightVerticalEdge",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1633),
		"Outline Upper Horizontal Edge", // OutlineUpperHorizontalEdge
		"OutlineUpperHorizontalEdge",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1634),
		"Outline Lower Horizontal Edge", // OutlineLowerHorizontalEdge
		"OutlineLowerHorizontalEdge",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1635),
		"Center of Circular Outline", // CenterOfCircularOutline
		"CenterOfCircularOutline",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1636),
		"Diameter of Circular Outline", // DiameterOfCircularOutline
		"DiameterOfCircularOutline",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1637),
		"Number of Polygonal Vertices", // NumberOfPolygonalVertices
		"NumberOfPolygonalVertices",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1638),
		"Vertices of the Polygonal Outline", // VerticesOfThePolygonalOutline
		"VerticesOfThePolygonalOutline",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1700),
		"Collimator Shape", // CollimatorShape
		"CollimatorShape",
		vm.VM1_3,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1702),
		"Collimator Left Vertical Edge", // CollimatorLeftVerticalEdge
		"CollimatorLeftVerticalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1704),
		"Collimator Right Vertical Edge", // CollimatorRightVerticalEdge
		"CollimatorRightVerticalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1706),
		"Collimator Upper Horizontal Edge", // CollimatorUpperHorizontalEdge
		"CollimatorUpperHorizontalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1708),
		"Collimator Lower Horizontal Edge", // CollimatorLowerHorizontalEdge
		"CollimatorLowerHorizontalEdge",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1710),
		"Center of Circular Collimator", // CenterOfCircularCollimator
		"CenterOfCircularCollimator",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1712),
		"Radius of Circular Collimator", // RadiusOfCircularCollimator
		"RadiusOfCircularCollimator",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1720),
		"Vertices of the Polygonal Collimator", // VerticesOfThePolygonalCollimator
		"VerticesOfThePolygonalCollimator",
		vm.VM2_2n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1800),
		"Acquisition Time Synchronized", // AcquisitionTimeSynchronized
		"AcquisitionTimeSynchronized",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1801),
		"Time Source", // TimeSource
		"TimeSource",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1802),
		"Time Distribution Protocol", // TimeDistributionProtocol
		"TimeDistributionProtocol",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x1803),
		"NTP Source Address", // NTPSourceAddress
		"NTPSourceAddress",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2001),
		"Page Number Vector", // PageNumberVector
		"PageNumberVector",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2002),
		"Frame Label Vector", // FrameLabelVector
		"FrameLabelVector",
		vm.VM1_n,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2003),
		"Frame Primary Angle Vector", // FramePrimaryAngleVector
		"FramePrimaryAngleVector",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2004),
		"Frame Secondary Angle Vector", // FrameSecondaryAngleVector
		"FrameSecondaryAngleVector",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2005),
		"Slice Location Vector", // SliceLocationVector
		"SliceLocationVector",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2006),
		"Display Window Label Vector", // DisplayWindowLabelVector
		"DisplayWindowLabelVector",
		vm.VM1_n,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2010),
		"Nominal Scanned Pixel Spacing", // NominalScannedPixelSpacing
		"NominalScannedPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2020),
		"Digitizing Device Transport Direction", // DigitizingDeviceTransportDirection
		"DigitizingDeviceTransportDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2030),
		"Rotation of Scanned Film", // RotationOfScannedFilm
		"RotationOfScannedFilm",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2041),
		"Biopsy Target Sequence", // BiopsyTargetSequence
		"BiopsyTargetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2042),
		"Target UID", // TargetUID
		"TargetUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2043),
		"Localizing Cursor Position", // LocalizingCursorPosition
		"LocalizingCursorPosition",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2044),
		"Calculated Target Position", // CalculatedTargetPosition
		"CalculatedTargetPosition",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2045),
		"Target Label", // TargetLabel
		"TargetLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x2046),
		"Displayed Z Value", // DisplayedZValue
		"DisplayedZValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3100),
		"IVUS Acquisition", // IVUSAcquisition
		"IVUSAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3101),
		"IVUS Pullback Rate", // IVUSPullbackRate
		"IVUSPullbackRate",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3102),
		"IVUS Gated Rate", // IVUSGatedRate
		"IVUSGatedRate",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3103),
		"IVUS Pullback Start Frame Number", // IVUSPullbackStartFrameNumber
		"IVUSPullbackStartFrameNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3104),
		"IVUS Pullback Stop Frame Number", // IVUSPullbackStopFrameNumber
		"IVUSPullbackStopFrameNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x3105),
		"Lesion Number", // LesionNumber
		"LesionNumber",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x4000),
		"Acquisition Comments", // AcquisitionComments
		"AcquisitionComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5000),
		"Output Power", // OutputPower
		"OutputPower",
		vm.VM1_n,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5010),
		"Transducer Data", // TransducerData
		"TransducerData",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5011),
		"Transducer Identification Sequence", // TransducerIdentificationSequence
		"TransducerIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5012),
		"Focus Depth", // FocusDepth
		"FocusDepth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5020),
		"Processing Function", // ProcessingFunction
		"ProcessingFunction",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5021),
		"Postprocessing Function", // PostprocessingFunction
		"PostprocessingFunction",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5022),
		"Mechanical Index", // MechanicalIndex
		"MechanicalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5024),
		"Bone Thermal Index", // BoneThermalIndex
		"BoneThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5026),
		"Cranial Thermal Index", // CranialThermalIndex
		"CranialThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5027),
		"Soft Tissue Thermal Index", // SoftTissueThermalIndex
		"SoftTissueThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5028),
		"Soft Tissue-focus Thermal Index", // SoftTissueFocusThermalIndex
		"SoftTissueFocusThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5029),
		"Soft Tissue-surface Thermal Index", // SoftTissueSurfaceThermalIndex
		"SoftTissueSurfaceThermalIndex",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5030),
		"Dynamic Range", // DynamicRange
		"DynamicRange",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5040),
		"Total Gain", // TotalGain
		"TotalGain",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5050),
		"Depth of Scan Field", // DepthOfScanField
		"DepthOfScanField",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5100),
		"Patient Position", // PatientPosition
		"PatientPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5101),
		"View Position", // ViewPosition
		"ViewPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5104),
		"Projection Eponymous Name Code Sequence", // ProjectionEponymousNameCodeSequence
		"ProjectionEponymousNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5210),
		"Image Transformation Matrix", // ImageTransformationMatrix
		"ImageTransformationMatrix",
		vm.VM6,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x5212),
		"Image Translation Vector", // ImageTranslationVector
		"ImageTranslationVector",
		vm.VM3,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6000),
		"Sensitivity", // Sensitivity
		"Sensitivity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6011),
		"Sequence of Ultrasound Regions", // SequenceOfUltrasoundRegions
		"SequenceOfUltrasoundRegions",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6012),
		"Region Spatial Format", // RegionSpatialFormat
		"RegionSpatialFormat",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6014),
		"Region Data Type", // RegionDataType
		"RegionDataType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6016),
		"Region Flags", // RegionFlags
		"RegionFlags",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6018),
		"Region Location Min X0", // RegionLocationMinX0
		"RegionLocationMinX0",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x601A),
		"Region Location Min Y0", // RegionLocationMinY0
		"RegionLocationMinY0",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x601C),
		"Region Location Max X1", // RegionLocationMaxX1
		"RegionLocationMaxX1",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x601E),
		"Region Location Max Y1", // RegionLocationMaxY1
		"RegionLocationMaxY1",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6020),
		"Reference Pixel X0", // ReferencePixelX0
		"ReferencePixelX0",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6022),
		"Reference Pixel Y0", // ReferencePixelY0
		"ReferencePixelY0",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6024),
		"Physical Units X Direction", // PhysicalUnitsXDirection
		"PhysicalUnitsXDirection",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6026),
		"Physical Units Y Direction", // PhysicalUnitsYDirection
		"PhysicalUnitsYDirection",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6028),
		"Reference Pixel Physical Value X", // ReferencePixelPhysicalValueX
		"ReferencePixelPhysicalValueX",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x602A),
		"Reference Pixel Physical Value Y", // ReferencePixelPhysicalValueY
		"ReferencePixelPhysicalValueY",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x602C),
		"Physical Delta X", // PhysicalDeltaX
		"PhysicalDeltaX",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x602E),
		"Physical Delta Y", // PhysicalDeltaY
		"PhysicalDeltaY",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6030),
		"Transducer Frequency", // TransducerFrequency
		"TransducerFrequency",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6031),
		"Transducer Type", // TransducerType
		"TransducerType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6032),
		"Pulse Repetition Frequency", // PulseRepetitionFrequency
		"PulseRepetitionFrequency",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6034),
		"Doppler Correction Angle", // DopplerCorrectionAngle
		"DopplerCorrectionAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6036),
		"Steering Angle", // SteeringAngle
		"SteeringAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6038),
		"Doppler Sample Volume X Position (Retired)", // DopplerSampleVolumeXPositionRetired
		"DopplerSampleVolumeXPositionRetired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6039),
		"Doppler Sample Volume X Position", // DopplerSampleVolumeXPosition
		"DopplerSampleVolumeXPosition",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603A),
		"Doppler Sample Volume Y Position (Retired)", // DopplerSampleVolumeYPositionRetired
		"DopplerSampleVolumeYPositionRetired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603B),
		"Doppler Sample Volume Y Position", // DopplerSampleVolumeYPosition
		"DopplerSampleVolumeYPosition",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603C),
		"TM-Line Position X0 (Retired)", // TMLinePositionX0Retired
		"TMLinePositionX0Retired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603D),
		"TM-Line Position X0", // TMLinePositionX0
		"TMLinePositionX0",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603E),
		"TM-Line Position Y0 (Retired)", // TMLinePositionY0Retired
		"TMLinePositionY0Retired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x603F),
		"TM-Line Position Y0", // TMLinePositionY0
		"TMLinePositionY0",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6040),
		"TM-Line Position X1 (Retired)", // TMLinePositionX1Retired
		"TMLinePositionX1Retired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6041),
		"TM-Line Position X1", // TMLinePositionX1
		"TMLinePositionX1",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6042),
		"TM-Line Position Y1 (Retired)", // TMLinePositionY1Retired
		"TMLinePositionY1Retired",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6043),
		"TM-Line Position Y1", // TMLinePositionY1
		"TMLinePositionY1",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6044),
		"Pixel Component Organization", // PixelComponentOrganization
		"PixelComponentOrganization",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6046),
		"Pixel Component Mask", // PixelComponentMask
		"PixelComponentMask",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6048),
		"Pixel Component Range Start", // PixelComponentRangeStart
		"PixelComponentRangeStart",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x604A),
		"Pixel Component Range Stop", // PixelComponentRangeStop
		"PixelComponentRangeStop",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x604C),
		"Pixel Component Physical Units", // PixelComponentPhysicalUnits
		"PixelComponentPhysicalUnits",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x604E),
		"Pixel Component Data Type", // PixelComponentDataType
		"PixelComponentDataType",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6050),
		"Number of Table Break Points", // NumberOfTableBreakPoints
		"NumberOfTableBreakPoints",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6052),
		"Table of X Break Points", // TableOfXBreakPoints
		"TableOfXBreakPoints",
		vm.VM1_n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6054),
		"Table of Y Break Points", // TableOfYBreakPoints
		"TableOfYBreakPoints",
		vm.VM1_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6056),
		"Number of Table Entries", // NumberOfTableEntries
		"NumberOfTableEntries",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6058),
		"Table of Pixel Values", // TableOfPixelValues
		"TableOfPixelValues",
		vm.VM1_n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x605A),
		"Table of Parameter Values", // TableOfParameterValues
		"TableOfParameterValues",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6060),
		"R Wave Time Vector", // RWaveTimeVector
		"RWaveTimeVector",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x6070),
		"Active Image Area Overlay Group", // ActiveImageAreaOverlayGroup
		"ActiveImageAreaOverlayGroup",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7000),
		"Detector Conditions Nominal Flag", // DetectorConditionsNominalFlag
		"DetectorConditionsNominalFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7001),
		"Detector Temperature", // DetectorTemperature
		"DetectorTemperature",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7004),
		"Detector Type", // DetectorType
		"DetectorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7005),
		"Detector Configuration", // DetectorConfiguration
		"DetectorConfiguration",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7006),
		"Detector Description", // DetectorDescription
		"DetectorDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7008),
		"Detector Mode", // DetectorMode
		"DetectorMode",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x700A),
		"Detector ID", // DetectorID
		"DetectorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x700C),
		"Date of Last Detector Calibration", // DateOfLastDetectorCalibration
		"DateOfLastDetectorCalibration",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x700E),
		"Time of Last Detector Calibration", // TimeOfLastDetectorCalibration
		"TimeOfLastDetectorCalibration",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7010),
		"Exposures on Detector Since Last Calibration", // ExposuresOnDetectorSinceLastCalibration
		"ExposuresOnDetectorSinceLastCalibration",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7011),
		"Exposures on Detector Since Manufactured", // ExposuresOnDetectorSinceManufactured
		"ExposuresOnDetectorSinceManufactured",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7012),
		"Detector Time Since Last Exposure", // DetectorTimeSinceLastExposure
		"DetectorTimeSinceLastExposure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7014),
		"Detector Active Time", // DetectorActiveTime
		"DetectorActiveTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7016),
		"Detector Activation Offset From Exposure", // DetectorActivationOffsetFromExposure
		"DetectorActivationOffsetFromExposure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x701A),
		"Detector Binning", // DetectorBinning
		"DetectorBinning",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7020),
		"Detector Element Physical Size", // DetectorElementPhysicalSize
		"DetectorElementPhysicalSize",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7022),
		"Detector Element Spacing", // DetectorElementSpacing
		"DetectorElementSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7024),
		"Detector Active Shape", // DetectorActiveShape
		"DetectorActiveShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7026),
		"Detector Active Dimension(s)", // DetectorActiveDimensions
		"DetectorActiveDimensions",
		vm.VM1_2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7028),
		"Detector Active Origin", // DetectorActiveOrigin
		"DetectorActiveOrigin",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x702A),
		"Detector Manufacturer Name", // DetectorManufacturerName
		"DetectorManufacturerName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x702B),
		"Detector Manufacturer's Model Name", // DetectorManufacturerModelName
		"DetectorManufacturerModelName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7030),
		"Field of View Origin", // FieldOfViewOrigin
		"FieldOfViewOrigin",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7032),
		"Field of View Rotation", // FieldOfViewRotation
		"FieldOfViewRotation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7034),
		"Field of View Horizontal Flip", // FieldOfViewHorizontalFlip
		"FieldOfViewHorizontalFlip",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7036),
		"Pixel Data Area Origin Relative To FOV", // PixelDataAreaOriginRelativeToFOV
		"PixelDataAreaOriginRelativeToFOV",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7038),
		"Pixel Data Area Rotation Angle Relative To FOV", // PixelDataAreaRotationAngleRelativeToFOV
		"PixelDataAreaRotationAngleRelativeToFOV",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7040),
		"Grid Absorbing Material", // GridAbsorbingMaterial
		"GridAbsorbingMaterial",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7041),
		"Grid Spacing Material", // GridSpacingMaterial
		"GridSpacingMaterial",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7042),
		"Grid Thickness", // GridThickness
		"GridThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7044),
		"Grid Pitch", // GridPitch
		"GridPitch",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7046),
		"Grid Aspect Ratio", // GridAspectRatio
		"GridAspectRatio",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7048),
		"Grid Period", // GridPeriod
		"GridPeriod",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x704C),
		"Grid Focal Distance", // GridFocalDistance
		"GridFocalDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7050),
		"Filter Material", // FilterMaterial
		"FilterMaterial",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7052),
		"Filter Thickness Minimum", // FilterThicknessMinimum
		"FilterThicknessMinimum",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7054),
		"Filter Thickness Maximum", // FilterThicknessMaximum
		"FilterThicknessMaximum",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7056),
		"Filter Beam Path Length Minimum", // FilterBeamPathLengthMinimum
		"FilterBeamPathLengthMinimum",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7058),
		"Filter Beam Path Length Maximum", // FilterBeamPathLengthMaximum
		"FilterBeamPathLengthMaximum",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7060),
		"Exposure Control Mode", // ExposureControlMode
		"ExposureControlMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7062),
		"Exposure Control Mode Description", // ExposureControlModeDescription
		"ExposureControlModeDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7064),
		"Exposure Status", // ExposureStatus
		"ExposureStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x7065),
		"Phototimer Setting", // PhototimerSetting
		"PhototimerSetting",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x8150),
		"Exposure Time in µS", // ExposureTimeInuS
		"ExposureTimeInuS",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x8151),
		"X-Ray Tube Current in µA", // XRayTubeCurrentInuA
		"XRayTubeCurrentInuA",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9004),
		"Content Qualification", // ContentQualification
		"ContentQualification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9005),
		"Pulse Sequence Name", // PulseSequenceName
		"PulseSequenceName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9006),
		"MR Imaging Modifier Sequence", // MRImagingModifierSequence
		"MRImagingModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9008),
		"Echo Pulse Sequence", // EchoPulseSequence
		"EchoPulseSequence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9009),
		"Inversion Recovery", // InversionRecovery
		"InversionRecovery",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9010),
		"Flow Compensation", // FlowCompensation
		"FlowCompensation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9011),
		"Multiple Spin Echo", // MultipleSpinEcho
		"MultipleSpinEcho",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9012),
		"Multi-planar Excitation", // MultiPlanarExcitation
		"MultiPlanarExcitation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9014),
		"Phase Contrast", // PhaseContrast
		"PhaseContrast",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9015),
		"Time of Flight Contrast", // TimeOfFlightContrast
		"TimeOfFlightContrast",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9016),
		"Spoiling", // Spoiling
		"Spoiling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9017),
		"Steady State Pulse Sequence", // SteadyStatePulseSequence
		"SteadyStatePulseSequence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9018),
		"Echo Planar Pulse Sequence", // EchoPlanarPulseSequence
		"EchoPlanarPulseSequence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9019),
		"Tag Angle First Axis", // TagAngleFirstAxis
		"TagAngleFirstAxis",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9020),
		"Magnetization Transfer", // MagnetizationTransfer
		"MagnetizationTransfer",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9021),
		"T2 Preparation", // T2Preparation
		"T2Preparation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9022),
		"Blood Signal Nulling", // BloodSignalNulling
		"BloodSignalNulling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9024),
		"Saturation Recovery", // SaturationRecovery
		"SaturationRecovery",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9025),
		"Spectrally Selected Suppression", // SpectrallySelectedSuppression
		"SpectrallySelectedSuppression",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9026),
		"Spectrally Selected Excitation", // SpectrallySelectedExcitation
		"SpectrallySelectedExcitation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9027),
		"Spatial Pre-saturation", // SpatialPresaturation
		"SpatialPresaturation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9028),
		"Tagging", // Tagging
		"Tagging",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9029),
		"Oversampling Phase", // OversamplingPhase
		"OversamplingPhase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9030),
		"Tag Spacing First Dimension", // TagSpacingFirstDimension
		"TagSpacingFirstDimension",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9032),
		"Geometry of k-Space Traversal", // GeometryOfKSpaceTraversal
		"GeometryOfKSpaceTraversal",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9033),
		"Segmented k-Space Traversal", // SegmentedKSpaceTraversal
		"SegmentedKSpaceTraversal",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9034),
		"Rectilinear Phase Encode Reordering", // RectilinearPhaseEncodeReordering
		"RectilinearPhaseEncodeReordering",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9035),
		"Tag Thickness", // TagThickness
		"TagThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9036),
		"Partial Fourier Direction", // PartialFourierDirection
		"PartialFourierDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9037),
		"Cardiac Synchronization Technique", // CardiacSynchronizationTechnique
		"CardiacSynchronizationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9041),
		"Receive Coil Manufacturer Name", // ReceiveCoilManufacturerName
		"ReceiveCoilManufacturerName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9042),
		"MR Receive Coil Sequence", // MRReceiveCoilSequence
		"MRReceiveCoilSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9043),
		"Receive Coil Type", // ReceiveCoilType
		"ReceiveCoilType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9044),
		"Quadrature Receive Coil", // QuadratureReceiveCoil
		"QuadratureReceiveCoil",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9045),
		"Multi-Coil Definition Sequence", // MultiCoilDefinitionSequence
		"MultiCoilDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9046),
		"Multi-Coil Configuration", // MultiCoilConfiguration
		"MultiCoilConfiguration",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9047),
		"Multi-Coil Element Name", // MultiCoilElementName
		"MultiCoilElementName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9048),
		"Multi-Coil Element Used", // MultiCoilElementUsed
		"MultiCoilElementUsed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9049),
		"MR Transmit Coil Sequence", // MRTransmitCoilSequence
		"MRTransmitCoilSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9050),
		"Transmit Coil Manufacturer Name", // TransmitCoilManufacturerName
		"TransmitCoilManufacturerName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9051),
		"Transmit Coil Type", // TransmitCoilType
		"TransmitCoilType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9052),
		"Spectral Width", // SpectralWidth
		"SpectralWidth",
		vm.VM1_2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9053),
		"Chemical Shift Reference", // ChemicalShiftReference
		"ChemicalShiftReference",
		vm.VM1_2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9054),
		"Volume Localization Technique", // VolumeLocalizationTechnique
		"VolumeLocalizationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9058),
		"MR Acquisition Frequency Encoding Steps", // MRAcquisitionFrequencyEncodingSteps
		"MRAcquisitionFrequencyEncodingSteps",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9059),
		"De-coupling", // Decoupling
		"Decoupling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9060),
		"De-coupled Nucleus", // DecoupledNucleus
		"DecoupledNucleus",
		vm.VM1_2,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9061),
		"De-coupling Frequency", // DecouplingFrequency
		"DecouplingFrequency",
		vm.VM1_2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9062),
		"De-coupling Method", // DecouplingMethod
		"DecouplingMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9063),
		"De-coupling Chemical Shift Reference", // DecouplingChemicalShiftReference
		"DecouplingChemicalShiftReference",
		vm.VM1_2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9064),
		"k-space Filtering", // KSpaceFiltering
		"KSpaceFiltering",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9065),
		"Time Domain Filtering", // TimeDomainFiltering
		"TimeDomainFiltering",
		vm.VM1_2,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9066),
		"Number of Zero Fills", // NumberOfZeroFills
		"NumberOfZeroFills",
		vm.VM1_2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9067),
		"Baseline Correction", // BaselineCorrection
		"BaselineCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9069),
		"Parallel Reduction Factor In-plane", // ParallelReductionFactorInPlane
		"ParallelReductionFactorInPlane",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9070),
		"Cardiac R-R Interval Specified", // CardiacRRIntervalSpecified
		"CardiacRRIntervalSpecified",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9073),
		"Acquisition Duration", // AcquisitionDuration
		"AcquisitionDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9074),
		"Frame Acquisition DateTime", // FrameAcquisitionDateTime
		"FrameAcquisitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9075),
		"Diffusion Directionality", // DiffusionDirectionality
		"DiffusionDirectionality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9076),
		"Diffusion Gradient Direction Sequence", // DiffusionGradientDirectionSequence
		"DiffusionGradientDirectionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9077),
		"Parallel Acquisition", // ParallelAcquisition
		"ParallelAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9078),
		"Parallel Acquisition Technique", // ParallelAcquisitionTechnique
		"ParallelAcquisitionTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9079),
		"Inversion Times", // InversionTimes
		"InversionTimes",
		vm.VM1_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9080),
		"Metabolite Map Description", // MetaboliteMapDescription
		"MetaboliteMapDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9081),
		"Partial Fourier", // PartialFourier
		"PartialFourier",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9082),
		"Effective Echo Time", // EffectiveEchoTime
		"EffectiveEchoTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9083),
		"Metabolite Map Code Sequence", // MetaboliteMapCodeSequence
		"MetaboliteMapCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9084),
		"Chemical Shift Sequence", // ChemicalShiftSequence
		"ChemicalShiftSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9085),
		"Cardiac Signal Source", // CardiacSignalSource
		"CardiacSignalSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9087),
		"Diffusion b-value", // DiffusionBValue
		"DiffusionBValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9089),
		"Diffusion Gradient Orientation", // DiffusionGradientOrientation
		"DiffusionGradientOrientation",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9090),
		"Velocity Encoding Direction", // VelocityEncodingDirection
		"VelocityEncodingDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9091),
		"Velocity Encoding Minimum Value", // VelocityEncodingMinimumValue
		"VelocityEncodingMinimumValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9092),
		"Velocity Encoding Acquisition Sequence", // VelocityEncodingAcquisitionSequence
		"VelocityEncodingAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9093),
		"Number of k-Space Trajectories", // NumberOfKSpaceTrajectories
		"NumberOfKSpaceTrajectories",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9094),
		"Coverage of k-Space", // CoverageOfKSpace
		"CoverageOfKSpace",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9095),
		"Spectroscopy Acquisition Phase Rows", // SpectroscopyAcquisitionPhaseRows
		"SpectroscopyAcquisitionPhaseRows",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9096),
		"Parallel Reduction Factor In-plane (Retired)", // ParallelReductionFactorInPlaneRetired
		"ParallelReductionFactorInPlaneRetired",
		vm.VM1,
		true,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9098),
		"Transmitter Frequency", // TransmitterFrequency
		"TransmitterFrequency",
		vm.VM1_2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9100),
		"Resonant Nucleus", // ResonantNucleus
		"ResonantNucleus",
		vm.VM1_2,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9101),
		"Frequency Correction", // FrequencyCorrection
		"FrequencyCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9103),
		"MR Spectroscopy FOV/Geometry Sequence", // MRSpectroscopyFOVGeometrySequence
		"MRSpectroscopyFOVGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9104),
		"Slab Thickness", // SlabThickness
		"SlabThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9105),
		"Slab Orientation", // SlabOrientation
		"SlabOrientation",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9106),
		"Mid Slab Position", // MidSlabPosition
		"MidSlabPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9107),
		"MR Spatial Saturation Sequence", // MRSpatialSaturationSequence
		"MRSpatialSaturationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9112),
		"MR Timing and Related Parameters Sequence", // MRTimingAndRelatedParametersSequence
		"MRTimingAndRelatedParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9114),
		"MR Echo Sequence", // MREchoSequence
		"MREchoSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9115),
		"MR Modifier Sequence", // MRModifierSequence
		"MRModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9117),
		"MR Diffusion Sequence", // MRDiffusionSequence
		"MRDiffusionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9118),
		"Cardiac Synchronization Sequence", // CardiacSynchronizationSequence
		"CardiacSynchronizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9119),
		"MR Averages Sequence", // MRAveragesSequence
		"MRAveragesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9125),
		"MR FOV/Geometry Sequence", // MRFOVGeometrySequence
		"MRFOVGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9126),
		"Volume Localization Sequence", // VolumeLocalizationSequence
		"VolumeLocalizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9127),
		"Spectroscopy Acquisition Data Columns", // SpectroscopyAcquisitionDataColumns
		"SpectroscopyAcquisitionDataColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9147),
		"Diffusion Anisotropy Type", // DiffusionAnisotropyType
		"DiffusionAnisotropyType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9151),
		"Frame Reference DateTime", // FrameReferenceDateTime
		"FrameReferenceDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9152),
		"MR Metabolite Map Sequence", // MRMetaboliteMapSequence
		"MRMetaboliteMapSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9155),
		"Parallel Reduction Factor out-of-plane", // ParallelReductionFactorOutOfPlane
		"ParallelReductionFactorOutOfPlane",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9159),
		"Spectroscopy Acquisition Out-of-plane Phase Steps", // SpectroscopyAcquisitionOutOfPlanePhaseSteps
		"SpectroscopyAcquisitionOutOfPlanePhaseSteps",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9166),
		"Bulk Motion Status", // BulkMotionStatus
		"BulkMotionStatus",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9168),
		"Parallel Reduction Factor Second In-plane", // ParallelReductionFactorSecondInPlane
		"ParallelReductionFactorSecondInPlane",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9169),
		"Cardiac Beat Rejection Technique", // CardiacBeatRejectionTechnique
		"CardiacBeatRejectionTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9170),
		"Respiratory Motion Compensation Technique", // RespiratoryMotionCompensationTechnique
		"RespiratoryMotionCompensationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9171),
		"Respiratory Signal Source", // RespiratorySignalSource
		"RespiratorySignalSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9172),
		"Bulk Motion Compensation Technique", // BulkMotionCompensationTechnique
		"BulkMotionCompensationTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9173),
		"Bulk Motion Signal Source", // BulkMotionSignalSource
		"BulkMotionSignalSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9174),
		"Applicable Safety Standard Agency", // ApplicableSafetyStandardAgency
		"ApplicableSafetyStandardAgency",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9175),
		"Applicable Safety Standard Description", // ApplicableSafetyStandardDescription
		"ApplicableSafetyStandardDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9176),
		"Operating Mode Sequence", // OperatingModeSequence
		"OperatingModeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9177),
		"Operating Mode Type", // OperatingModeType
		"OperatingModeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9178),
		"Operating Mode", // OperatingMode
		"OperatingMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9179),
		"Specific Absorption Rate Definition", // SpecificAbsorptionRateDefinition
		"SpecificAbsorptionRateDefinition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9180),
		"Gradient Output Type", // GradientOutputType
		"GradientOutputType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9181),
		"Specific Absorption Rate Value", // SpecificAbsorptionRateValue
		"SpecificAbsorptionRateValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9182),
		"Gradient Output", // GradientOutput
		"GradientOutput",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9183),
		"Flow Compensation Direction", // FlowCompensationDirection
		"FlowCompensationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9184),
		"Tagging Delay", // TaggingDelay
		"TaggingDelay",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9185),
		"Respiratory Motion Compensation Technique Description", // RespiratoryMotionCompensationTechniqueDescription
		"RespiratoryMotionCompensationTechniqueDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9186),
		"Respiratory Signal Source ID", // RespiratorySignalSourceID
		"RespiratorySignalSourceID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9195),
		"Chemical Shift Minimum Integration Limit in Hz", // ChemicalShiftMinimumIntegrationLimitInHz
		"ChemicalShiftMinimumIntegrationLimitInHz",
		vm.VM1,
		true,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9196),
		"Chemical Shift Maximum Integration Limit in Hz", // ChemicalShiftMaximumIntegrationLimitInHz
		"ChemicalShiftMaximumIntegrationLimitInHz",
		vm.VM1,
		true,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9197),
		"MR Velocity Encoding Sequence", // MRVelocityEncodingSequence
		"MRVelocityEncodingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9198),
		"First Order Phase Correction", // FirstOrderPhaseCorrection
		"FirstOrderPhaseCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9199),
		"Water Referenced Phase Correction", // WaterReferencedPhaseCorrection
		"WaterReferencedPhaseCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9200),
		"MR Spectroscopy Acquisition Type", // MRSpectroscopyAcquisitionType
		"MRSpectroscopyAcquisitionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9214),
		"Respiratory Cycle Position", // RespiratoryCyclePosition
		"RespiratoryCyclePosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9217),
		"Velocity Encoding Maximum Value", // VelocityEncodingMaximumValue
		"VelocityEncodingMaximumValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9218),
		"Tag Spacing Second Dimension", // TagSpacingSecondDimension
		"TagSpacingSecondDimension",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9219),
		"Tag Angle Second Axis", // TagAngleSecondAxis
		"TagAngleSecondAxis",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9220),
		"Frame Acquisition Duration", // FrameAcquisitionDuration
		"FrameAcquisitionDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9226),
		"MR Image Frame Type Sequence", // MRImageFrameTypeSequence
		"MRImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9227),
		"MR Spectroscopy Frame Type Sequence", // MRSpectroscopyFrameTypeSequence
		"MRSpectroscopyFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9231),
		"MR Acquisition Phase Encoding Steps in-plane", // MRAcquisitionPhaseEncodingStepsInPlane
		"MRAcquisitionPhaseEncodingStepsInPlane",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9232),
		"MR Acquisition Phase Encoding Steps out-of-plane", // MRAcquisitionPhaseEncodingStepsOutOfPlane
		"MRAcquisitionPhaseEncodingStepsOutOfPlane",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9234),
		"Spectroscopy Acquisition Phase Columns", // SpectroscopyAcquisitionPhaseColumns
		"SpectroscopyAcquisitionPhaseColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9236),
		"Cardiac Cycle Position", // CardiacCyclePosition
		"CardiacCyclePosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9239),
		"Specific Absorption Rate Sequence", // SpecificAbsorptionRateSequence
		"SpecificAbsorptionRateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9240),
		"RF Echo Train Length", // RFEchoTrainLength
		"RFEchoTrainLength",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9241),
		"Gradient Echo Train Length", // GradientEchoTrainLength
		"GradientEchoTrainLength",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9250),
		"Arterial Spin Labeling Contrast", // ArterialSpinLabelingContrast
		"ArterialSpinLabelingContrast",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9251),
		"MR Arterial Spin Labeling Sequence", // MRArterialSpinLabelingSequence
		"MRArterialSpinLabelingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9252),
		"ASL Technique Description", // ASLTechniqueDescription
		"ASLTechniqueDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9253),
		"ASL Slab Number", // ASLSlabNumber
		"ASLSlabNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9254),
		"ASL Slab Thickness", // ASLSlabThickness
		"ASLSlabThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9255),
		"ASL Slab Orientation", // ASLSlabOrientation
		"ASLSlabOrientation",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9256),
		"ASL Mid Slab Position", // ASLMidSlabPosition
		"ASLMidSlabPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9257),
		"ASL Context", // ASLContext
		"ASLContext",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9258),
		"ASL Pulse Train Duration", // ASLPulseTrainDuration
		"ASLPulseTrainDuration",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9259),
		"ASL Crusher Flag", // ASLCrusherFlag
		"ASLCrusherFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925A),
		"ASL Crusher Flow Limit", // ASLCrusherFlowLimit
		"ASLCrusherFlowLimit",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925B),
		"ASL Crusher Description", // ASLCrusherDescription
		"ASLCrusherDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925C),
		"ASL Bolus Cut-off Flag", // ASLBolusCutoffFlag
		"ASLBolusCutoffFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925D),
		"ASL Bolus Cut-off Timing Sequence", // ASLBolusCutoffTimingSequence
		"ASLBolusCutoffTimingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925E),
		"ASL Bolus Cut-off Technique", // ASLBolusCutoffTechnique
		"ASLBolusCutoffTechnique",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x925F),
		"ASL Bolus Cut-off Delay Time", // ASLBolusCutoffDelayTime
		"ASLBolusCutoffDelayTime",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9260),
		"ASL Slab Sequence", // ASLSlabSequence
		"ASLSlabSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9295),
		"Chemical Shift Minimum Integration Limit in ppm", // ChemicalShiftMinimumIntegrationLimitInppm
		"ChemicalShiftMinimumIntegrationLimitInppm",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9296),
		"Chemical Shift Maximum Integration Limit in ppm", // ChemicalShiftMaximumIntegrationLimitInppm
		"ChemicalShiftMaximumIntegrationLimitInppm",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9297),
		"Water Reference Acquisition", // WaterReferenceAcquisition
		"WaterReferenceAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9298),
		"Echo Peak Position", // EchoPeakPosition
		"EchoPeakPosition",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9301),
		"CT Acquisition Type Sequence", // CTAcquisitionTypeSequence
		"CTAcquisitionTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9302),
		"Acquisition Type", // AcquisitionType
		"AcquisitionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9303),
		"Tube Angle", // TubeAngle
		"TubeAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9304),
		"CT Acquisition Details Sequence", // CTAcquisitionDetailsSequence
		"CTAcquisitionDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9305),
		"Revolution Time", // RevolutionTime
		"RevolutionTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9306),
		"Single Collimation Width", // SingleCollimationWidth
		"SingleCollimationWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9307),
		"Total Collimation Width", // TotalCollimationWidth
		"TotalCollimationWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9308),
		"CT Table Dynamics Sequence", // CTTableDynamicsSequence
		"CTTableDynamicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9309),
		"Table Speed", // TableSpeed
		"TableSpeed",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9310),
		"Table Feed per Rotation", // TableFeedPerRotation
		"TableFeedPerRotation",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9311),
		"Spiral Pitch Factor", // SpiralPitchFactor
		"SpiralPitchFactor",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9312),
		"CT Geometry Sequence", // CTGeometrySequence
		"CTGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9313),
		"Data Collection Center (Patient)", // DataCollectionCenterPatient
		"DataCollectionCenterPatient",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9314),
		"CT Reconstruction Sequence", // CTReconstructionSequence
		"CTReconstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9315),
		"Reconstruction Algorithm", // ReconstructionAlgorithm
		"ReconstructionAlgorithm",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9316),
		"Convolution Kernel Group", // ConvolutionKernelGroup
		"ConvolutionKernelGroup",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9317),
		"Reconstruction Field of View", // ReconstructionFieldOfView
		"ReconstructionFieldOfView",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9318),
		"Reconstruction Target Center (Patient)", // ReconstructionTargetCenterPatient
		"ReconstructionTargetCenterPatient",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9319),
		"Reconstruction Angle", // ReconstructionAngle
		"ReconstructionAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9320),
		"Image Filter", // ImageFilter
		"ImageFilter",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9321),
		"CT Exposure Sequence", // CTExposureSequence
		"CTExposureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9322),
		"Reconstruction Pixel Spacing", // ReconstructionPixelSpacing
		"ReconstructionPixelSpacing",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9323),
		"Exposure Modulation Type", // ExposureModulationType
		"ExposureModulationType",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9324),
		"Estimated Dose Saving", // EstimatedDoseSaving
		"EstimatedDoseSaving",
		vm.VM1,
		true,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9325),
		"CT X-Ray Details Sequence", // CTXRayDetailsSequence
		"CTXRayDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9326),
		"CT Position Sequence", // CTPositionSequence
		"CTPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9327),
		"Table Position", // TablePosition
		"TablePosition",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9328),
		"Exposure Time in ms", // ExposureTimeInms
		"ExposureTimeInms",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9329),
		"CT Image Frame Type Sequence", // CTImageFrameTypeSequence
		"CTImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9330),
		"X-Ray Tube Current in mA", // XRayTubeCurrentInmA
		"XRayTubeCurrentInmA",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9332),
		"Exposure in mAs", // ExposureInmAs
		"ExposureInmAs",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9333),
		"Constant Volume Flag", // ConstantVolumeFlag
		"ConstantVolumeFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9334),
		"Fluoroscopy Flag", // FluoroscopyFlag
		"FluoroscopyFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9335),
		"Distance Source to Data Collection Center", // DistanceSourceToDataCollectionCenter
		"DistanceSourceToDataCollectionCenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9337),
		"Contrast/Bolus Agent Number", // ContrastBolusAgentNumber
		"ContrastBolusAgentNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9338),
		"Contrast/Bolus Ingredient Code Sequence", // ContrastBolusIngredientCodeSequence
		"ContrastBolusIngredientCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9340),
		"Contrast Administration Profile Sequence", // ContrastAdministrationProfileSequence
		"ContrastAdministrationProfileSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9341),
		"Contrast/Bolus Usage Sequence", // ContrastBolusUsageSequence
		"ContrastBolusUsageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9342),
		"Contrast/Bolus Agent Administered", // ContrastBolusAgentAdministered
		"ContrastBolusAgentAdministered",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9343),
		"Contrast/Bolus Agent Detected", // ContrastBolusAgentDetected
		"ContrastBolusAgentDetected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9344),
		"Contrast/Bolus Agent Phase", // ContrastBolusAgentPhase
		"ContrastBolusAgentPhase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9345),
		"CTDIvol", // CTDIvol
		"CTDIvol",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9346),
		"CTDI Phantom Type Code Sequence", // CTDIPhantomTypeCodeSequence
		"CTDIPhantomTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9351),
		"Calcium Scoring Mass Factor Patient", // CalciumScoringMassFactorPatient
		"CalciumScoringMassFactorPatient",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9352),
		"Calcium Scoring Mass Factor Device", // CalciumScoringMassFactorDevice
		"CalciumScoringMassFactorDevice",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9353),
		"Energy Weighting Factor", // EnergyWeightingFactor
		"EnergyWeightingFactor",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9360),
		"CT Additional X-Ray Source Sequence", // CTAdditionalXRaySourceSequence
		"CTAdditionalXRaySourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9361),
		"Multi-energy CT Acquisition", // MultienergyCTAcquisition
		"MultienergyCTAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9362),
		"Multi-energy CT Acquisition Sequence", // MultienergyCTAcquisitionSequence
		"MultienergyCTAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9363),
		"Multi-energy CT Processing Sequence", // MultienergyCTProcessingSequence
		"MultienergyCTProcessingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9364),
		"Multi-energy CT Characteristics Sequence", // MultienergyCTCharacteristicsSequence
		"MultienergyCTCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9365),
		"Multi-energy CT X-Ray Source Sequence", // MultienergyCTXRaySourceSequence
		"MultienergyCTXRaySourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9366),
		"X-Ray Source Index", // XRaySourceIndex
		"XRaySourceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9367),
		"X-Ray Source ID", // XRaySourceID
		"XRaySourceID",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9368),
		"Multi-energy Source Technique", // MultienergySourceTechnique
		"MultienergySourceTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9369),
		"Source Start DateTime", // SourceStartDateTime
		"SourceStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936A),
		"Source End DateTime", // SourceEndDateTime
		"SourceEndDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936B),
		"Switching Phase Number", // SwitchingPhaseNumber
		"SwitchingPhaseNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936C),
		"Switching Phase Nominal Duration", // SwitchingPhaseNominalDuration
		"SwitchingPhaseNominalDuration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936D),
		"Switching Phase Transition Duration", // SwitchingPhaseTransitionDuration
		"SwitchingPhaseTransitionDuration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936E),
		"Effective Bin Energy", // EffectiveBinEnergy
		"EffectiveBinEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x936F),
		"Multi-energy CT X-Ray Detector Sequence", // MultienergyCTXRayDetectorSequence
		"MultienergyCTXRayDetectorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9370),
		"X-Ray Detector Index", // XRayDetectorIndex
		"XRayDetectorIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9371),
		"X-Ray Detector ID", // XRayDetectorID
		"XRayDetectorID",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9372),
		"Multi-energy Detector Type", // MultienergyDetectorType
		"MultienergyDetectorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9373),
		"X-Ray Detector Label", // XRayDetectorLabel
		"XRayDetectorLabel",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9374),
		"Nominal Max Energy", // NominalMaxEnergy
		"NominalMaxEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9375),
		"Nominal Min Energy", // NominalMinEnergy
		"NominalMinEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9376),
		"Referenced X-Ray Detector Index", // ReferencedXRayDetectorIndex
		"ReferencedXRayDetectorIndex",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9377),
		"Referenced X-Ray Source Index", // ReferencedXRaySourceIndex
		"ReferencedXRaySourceIndex",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9378),
		"Referenced Path Index", // ReferencedPathIndex
		"ReferencedPathIndex",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9379),
		"Multi-energy CT Path Sequence", // MultienergyCTPathSequence
		"MultienergyCTPathSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937A),
		"Multi-energy CT Path Index", // MultienergyCTPathIndex
		"MultienergyCTPathIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937B),
		"Multi-energy Acquisition Description", // MultienergyAcquisitionDescription
		"MultienergyAcquisitionDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937C),
		"Monoenergetic Energy Equivalent", // MonoenergeticEnergyEquivalent
		"MonoenergeticEnergyEquivalent",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937D),
		"Material Code Sequence", // MaterialCodeSequence
		"MaterialCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937E),
		"Decomposition Method", // DecompositionMethod
		"DecompositionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x937F),
		"Decomposition Description", // DecompositionDescription
		"DecompositionDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9380),
		"Decomposition Algorithm Identification Sequence", // DecompositionAlgorithmIdentificationSequence
		"DecompositionAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9381),
		"Decomposition Material Sequence", // DecompositionMaterialSequence
		"DecompositionMaterialSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9382),
		"Material Attenuation Sequence", // MaterialAttenuationSequence
		"MaterialAttenuationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9383),
		"Photon Energy", // PhotonEnergy
		"PhotonEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9384),
		"X-Ray Mass Attenuation Coefficient", // XRayMassAttenuationCoefficient
		"XRayMassAttenuationCoefficient",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9401),
		"Projection Pixel Calibration Sequence", // ProjectionPixelCalibrationSequence
		"ProjectionPixelCalibrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9402),
		"Distance Source to Isocenter", // DistanceSourceToIsocenter
		"DistanceSourceToIsocenter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9403),
		"Distance Object to Table Top", // DistanceObjectToTableTop
		"DistanceObjectToTableTop",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9404),
		"Object Pixel Spacing in Center of Beam", // ObjectPixelSpacingInCenterOfBeam
		"ObjectPixelSpacingInCenterOfBeam",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9405),
		"Positioner Position Sequence", // PositionerPositionSequence
		"PositionerPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9406),
		"Table Position Sequence", // TablePositionSequence
		"TablePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9407),
		"Collimator Shape Sequence", // CollimatorShapeSequence
		"CollimatorShapeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9410),
		"Planes in Acquisition", // PlanesInAcquisition
		"PlanesInAcquisition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9412),
		"XA/XRF Frame Characteristics Sequence", // XAXRFFrameCharacteristicsSequence
		"XAXRFFrameCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9417),
		"Frame Acquisition Sequence", // FrameAcquisitionSequence
		"FrameAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9420),
		"X-Ray Receptor Type", // XRayReceptorType
		"XRayReceptorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9423),
		"Acquisition Protocol Name", // AcquisitionProtocolName
		"AcquisitionProtocolName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9424),
		"Acquisition Protocol Description", // AcquisitionProtocolDescription
		"AcquisitionProtocolDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9425),
		"Contrast/Bolus Ingredient Opaque", // ContrastBolusIngredientOpaque
		"ContrastBolusIngredientOpaque",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9426),
		"Distance Receptor Plane to Detector Housing", // DistanceReceptorPlaneToDetectorHousing
		"DistanceReceptorPlaneToDetectorHousing",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9427),
		"Intensifier Active Shape", // IntensifierActiveShape
		"IntensifierActiveShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9428),
		"Intensifier Active Dimension(s)", // IntensifierActiveDimensions
		"IntensifierActiveDimensions",
		vm.VM1_2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9429),
		"Physical Detector Size", // PhysicalDetectorSize
		"PhysicalDetectorSize",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9430),
		"Position of Isocenter Projection", // PositionOfIsocenterProjection
		"PositionOfIsocenterProjection",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9432),
		"Field of View Sequence", // FieldOfViewSequence
		"FieldOfViewSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9433),
		"Field of View Description", // FieldOfViewDescription
		"FieldOfViewDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9434),
		"Exposure Control Sensing Regions Sequence", // ExposureControlSensingRegionsSequence
		"ExposureControlSensingRegionsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9435),
		"Exposure Control Sensing Region Shape", // ExposureControlSensingRegionShape
		"ExposureControlSensingRegionShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9436),
		"Exposure Control Sensing Region Left Vertical Edge", // ExposureControlSensingRegionLeftVerticalEdge
		"ExposureControlSensingRegionLeftVerticalEdge",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9437),
		"Exposure Control Sensing Region Right Vertical Edge", // ExposureControlSensingRegionRightVerticalEdge
		"ExposureControlSensingRegionRightVerticalEdge",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9438),
		"Exposure Control Sensing Region Upper Horizontal Edge", // ExposureControlSensingRegionUpperHorizontalEdge
		"ExposureControlSensingRegionUpperHorizontalEdge",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9439),
		"Exposure Control Sensing Region Lower Horizontal Edge", // ExposureControlSensingRegionLowerHorizontalEdge
		"ExposureControlSensingRegionLowerHorizontalEdge",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9440),
		"Center of Circular Exposure Control Sensing Region", // CenterOfCircularExposureControlSensingRegion
		"CenterOfCircularExposureControlSensingRegion",
		vm.VM2,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9441),
		"Radius of Circular Exposure Control Sensing Region", // RadiusOfCircularExposureControlSensingRegion
		"RadiusOfCircularExposureControlSensingRegion",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9442),
		"Vertices of the Polygonal Exposure Control Sensing Region", // VerticesOfThePolygonalExposureControlSensingRegion
		"VerticesOfThePolygonalExposureControlSensingRegion",
		vm.VM2_n,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9447),
		"Column Angulation (Patient)", // ColumnAngulationPatient
		"ColumnAngulationPatient",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9449),
		"Beam Angle", // BeamAngle
		"BeamAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9451),
		"Frame Detector Parameters Sequence", // FrameDetectorParametersSequence
		"FrameDetectorParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9452),
		"Calculated Anatomy Thickness", // CalculatedAnatomyThickness
		"CalculatedAnatomyThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9455),
		"Calibration Sequence", // CalibrationSequence
		"CalibrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9456),
		"Object Thickness Sequence", // ObjectThicknessSequence
		"ObjectThicknessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9457),
		"Plane Identification", // PlaneIdentification
		"PlaneIdentification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9461),
		"Field of View Dimension(s) in Float", // FieldOfViewDimensionsInFloat
		"FieldOfViewDimensionsInFloat",
		vm.VM1_2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9462),
		"Isocenter Reference System Sequence", // IsocenterReferenceSystemSequence
		"IsocenterReferenceSystemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9463),
		"Positioner Isocenter Primary Angle", // PositionerIsocenterPrimaryAngle
		"PositionerIsocenterPrimaryAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9464),
		"Positioner Isocenter Secondary Angle", // PositionerIsocenterSecondaryAngle
		"PositionerIsocenterSecondaryAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9465),
		"Positioner Isocenter Detector Rotation Angle", // PositionerIsocenterDetectorRotationAngle
		"PositionerIsocenterDetectorRotationAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9466),
		"Table X Position to Isocenter", // TableXPositionToIsocenter
		"TableXPositionToIsocenter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9467),
		"Table Y Position to Isocenter", // TableYPositionToIsocenter
		"TableYPositionToIsocenter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9468),
		"Table Z Position to Isocenter", // TableZPositionToIsocenter
		"TableZPositionToIsocenter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9469),
		"Table Horizontal Rotation Angle", // TableHorizontalRotationAngle
		"TableHorizontalRotationAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9470),
		"Table Head Tilt Angle", // TableHeadTiltAngle
		"TableHeadTiltAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9471),
		"Table Cradle Tilt Angle", // TableCradleTiltAngle
		"TableCradleTiltAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9472),
		"Frame Display Shutter Sequence", // FrameDisplayShutterSequence
		"FrameDisplayShutterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9473),
		"Acquired Image Area Dose Product", // AcquiredImageAreaDoseProduct
		"AcquiredImageAreaDoseProduct",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9474),
		"C-arm Positioner Tabletop Relationship", // CArmPositionerTabletopRelationship
		"CArmPositionerTabletopRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9476),
		"X-Ray Geometry Sequence", // XRayGeometrySequence
		"XRayGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9477),
		"Irradiation Event Identification Sequence", // IrradiationEventIdentificationSequence
		"IrradiationEventIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9504),
		"X-Ray 3D Frame Type Sequence", // XRay3DFrameTypeSequence
		"XRay3DFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9506),
		"Contributing Sources Sequence", // ContributingSourcesSequence
		"ContributingSourcesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9507),
		"X-Ray 3D Acquisition Sequence", // XRay3DAcquisitionSequence
		"XRay3DAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9508),
		"Primary Positioner Scan Arc", // PrimaryPositionerScanArc
		"PrimaryPositionerScanArc",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9509),
		"Secondary Positioner Scan Arc", // SecondaryPositionerScanArc
		"SecondaryPositionerScanArc",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9510),
		"Primary Positioner Scan Start Angle", // PrimaryPositionerScanStartAngle
		"PrimaryPositionerScanStartAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9511),
		"Secondary Positioner Scan Start Angle", // SecondaryPositionerScanStartAngle
		"SecondaryPositionerScanStartAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9514),
		"Primary Positioner Increment", // PrimaryPositionerIncrement
		"PrimaryPositionerIncrement",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9515),
		"Secondary Positioner Increment", // SecondaryPositionerIncrement
		"SecondaryPositionerIncrement",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9516),
		"Start Acquisition DateTime", // StartAcquisitionDateTime
		"StartAcquisitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9517),
		"End Acquisition DateTime", // EndAcquisitionDateTime
		"EndAcquisitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9518),
		"Primary Positioner Increment Sign", // PrimaryPositionerIncrementSign
		"PrimaryPositionerIncrementSign",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9519),
		"Secondary Positioner Increment Sign", // SecondaryPositionerIncrementSign
		"SecondaryPositionerIncrementSign",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9524),
		"Application Name", // ApplicationName
		"ApplicationName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9525),
		"Application Version", // ApplicationVersion
		"ApplicationVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9526),
		"Application Manufacturer", // ApplicationManufacturer
		"ApplicationManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9527),
		"Algorithm Type", // AlgorithmType
		"AlgorithmType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9528),
		"Algorithm Description", // AlgorithmDescription
		"AlgorithmDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9530),
		"X-Ray 3D Reconstruction Sequence", // XRay3DReconstructionSequence
		"XRay3DReconstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9531),
		"Reconstruction Description", // ReconstructionDescription
		"ReconstructionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9538),
		"Per Projection Acquisition Sequence", // PerProjectionAcquisitionSequence
		"PerProjectionAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9541),
		"Detector Position Sequence", // DetectorPositionSequence
		"DetectorPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9542),
		"X-Ray Acquisition Dose Sequence", // XRayAcquisitionDoseSequence
		"XRayAcquisitionDoseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9543),
		"X-Ray Source Isocenter Primary Angle", // XRaySourceIsocenterPrimaryAngle
		"XRaySourceIsocenterPrimaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9544),
		"X-Ray Source Isocenter Secondary Angle", // XRaySourceIsocenterSecondaryAngle
		"XRaySourceIsocenterSecondaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9545),
		"Breast Support Isocenter Primary Angle", // BreastSupportIsocenterPrimaryAngle
		"BreastSupportIsocenterPrimaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9546),
		"Breast Support Isocenter Secondary Angle", // BreastSupportIsocenterSecondaryAngle
		"BreastSupportIsocenterSecondaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9547),
		"Breast Support X Position to Isocenter", // BreastSupportXPositionToIsocenter
		"BreastSupportXPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9548),
		"Breast Support Y Position to Isocenter", // BreastSupportYPositionToIsocenter
		"BreastSupportYPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9549),
		"Breast Support Z Position to Isocenter", // BreastSupportZPositionToIsocenter
		"BreastSupportZPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9550),
		"Detector Isocenter Primary Angle", // DetectorIsocenterPrimaryAngle
		"DetectorIsocenterPrimaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9551),
		"Detector Isocenter Secondary Angle", // DetectorIsocenterSecondaryAngle
		"DetectorIsocenterSecondaryAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9552),
		"Detector X Position to Isocenter", // DetectorXPositionToIsocenter
		"DetectorXPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9553),
		"Detector Y Position to Isocenter", // DetectorYPositionToIsocenter
		"DetectorYPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9554),
		"Detector Z Position to Isocenter", // DetectorZPositionToIsocenter
		"DetectorZPositionToIsocenter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9555),
		"X-Ray Grid Sequence", // XRayGridSequence
		"XRayGridSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9556),
		"X-Ray Filter Sequence", // XRayFilterSequence
		"XRayFilterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9557),
		"Detector Active Area TLHC Position", // DetectorActiveAreaTLHCPosition
		"DetectorActiveAreaTLHCPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9558),
		"Detector Active Area Orientation", // DetectorActiveAreaOrientation
		"DetectorActiveAreaOrientation",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9559),
		"Positioner Primary Angle Direction", // PositionerPrimaryAngleDirection
		"PositionerPrimaryAngleDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9601),
		"Diffusion b-matrix Sequence", // DiffusionBMatrixSequence
		"DiffusionBMatrixSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9602),
		"Diffusion b-value XX", // DiffusionBValueXX
		"DiffusionBValueXX",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9603),
		"Diffusion b-value XY", // DiffusionBValueXY
		"DiffusionBValueXY",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9604),
		"Diffusion b-value XZ", // DiffusionBValueXZ
		"DiffusionBValueXZ",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9605),
		"Diffusion b-value YY", // DiffusionBValueYY
		"DiffusionBValueYY",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9606),
		"Diffusion b-value YZ", // DiffusionBValueYZ
		"DiffusionBValueYZ",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9607),
		"Diffusion b-value ZZ", // DiffusionBValueZZ
		"DiffusionBValueZZ",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9621),
		"Functional MR Sequence", // FunctionalMRSequence
		"FunctionalMRSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9622),
		"Functional Settling Phase Frames Present", // FunctionalSettlingPhaseFramesPresent
		"FunctionalSettlingPhaseFramesPresent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9623),
		"Functional Sync Pulse", // FunctionalSyncPulse
		"FunctionalSyncPulse",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9624),
		"Settling Phase Frame", // SettlingPhaseFrame
		"SettlingPhaseFrame",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9701),
		"Decay Correction DateTime", // DecayCorrectionDateTime
		"DecayCorrectionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9715),
		"Start Density Threshold", // StartDensityThreshold
		"StartDensityThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9716),
		"Start Relative Density Difference Threshold", // StartRelativeDensityDifferenceThreshold
		"StartRelativeDensityDifferenceThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9717),
		"Start Cardiac Trigger Count Threshold", // StartCardiacTriggerCountThreshold
		"StartCardiacTriggerCountThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9718),
		"Start Respiratory Trigger Count Threshold", // StartRespiratoryTriggerCountThreshold
		"StartRespiratoryTriggerCountThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9719),
		"Termination Counts Threshold", // TerminationCountsThreshold
		"TerminationCountsThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9720),
		"Termination Density Threshold", // TerminationDensityThreshold
		"TerminationDensityThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9721),
		"Termination Relative Density Threshold", // TerminationRelativeDensityThreshold
		"TerminationRelativeDensityThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9722),
		"Termination Time Threshold", // TerminationTimeThreshold
		"TerminationTimeThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9723),
		"Termination Cardiac Trigger Count Threshold", // TerminationCardiacTriggerCountThreshold
		"TerminationCardiacTriggerCountThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9724),
		"Termination Respiratory Trigger Count Threshold", // TerminationRespiratoryTriggerCountThreshold
		"TerminationRespiratoryTriggerCountThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9725),
		"Detector Geometry", // DetectorGeometry
		"DetectorGeometry",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9726),
		"Transverse Detector Separation", // TransverseDetectorSeparation
		"TransverseDetectorSeparation",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9727),
		"Axial Detector Dimension", // AxialDetectorDimension
		"AxialDetectorDimension",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9729),
		"Radiopharmaceutical Agent Number", // RadiopharmaceuticalAgentNumber
		"RadiopharmaceuticalAgentNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9732),
		"PET Frame Acquisition Sequence", // PETFrameAcquisitionSequence
		"PETFrameAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9733),
		"PET Detector Motion Details Sequence", // PETDetectorMotionDetailsSequence
		"PETDetectorMotionDetailsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9734),
		"PET Table Dynamics Sequence", // PETTableDynamicsSequence
		"PETTableDynamicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9735),
		"PET Position Sequence", // PETPositionSequence
		"PETPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9736),
		"PET Frame Correction Factors Sequence", // PETFrameCorrectionFactorsSequence
		"PETFrameCorrectionFactorsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9737),
		"Radiopharmaceutical Usage Sequence", // RadiopharmaceuticalUsageSequence
		"RadiopharmaceuticalUsageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9738),
		"Attenuation Correction Source", // AttenuationCorrectionSource
		"AttenuationCorrectionSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9739),
		"Number of Iterations", // NumberOfIterations
		"NumberOfIterations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9740),
		"Number of Subsets", // NumberOfSubsets
		"NumberOfSubsets",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9749),
		"PET Reconstruction Sequence", // PETReconstructionSequence
		"PETReconstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9751),
		"PET Frame Type Sequence", // PETFrameTypeSequence
		"PETFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9755),
		"Time of Flight Information Used", // TimeOfFlightInformationUsed
		"TimeOfFlightInformationUsed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9756),
		"Reconstruction Type", // ReconstructionType
		"ReconstructionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9758),
		"Decay Corrected", // DecayCorrected
		"DecayCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9759),
		"Attenuation Corrected", // AttenuationCorrected
		"AttenuationCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9760),
		"Scatter Corrected", // ScatterCorrected
		"ScatterCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9761),
		"Dead Time Corrected", // DeadTimeCorrected
		"DeadTimeCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9762),
		"Gantry Motion Corrected", // GantryMotionCorrected
		"GantryMotionCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9763),
		"Patient Motion Corrected", // PatientMotionCorrected
		"PatientMotionCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9764),
		"Count Loss Normalization Corrected", // CountLossNormalizationCorrected
		"CountLossNormalizationCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9765),
		"Randoms Corrected", // RandomsCorrected
		"RandomsCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9766),
		"Non-uniform Radial Sampling Corrected", // NonUniformRadialSamplingCorrected
		"NonUniformRadialSamplingCorrected",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9767),
		"Sensitivity Calibrated", // SensitivityCalibrated
		"SensitivityCalibrated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9768),
		"Detector Normalization Correction", // DetectorNormalizationCorrection
		"DetectorNormalizationCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9769),
		"Iterative Reconstruction Method", // IterativeReconstructionMethod
		"IterativeReconstructionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9770),
		"Attenuation Correction Temporal Relationship", // AttenuationCorrectionTemporalRelationship
		"AttenuationCorrectionTemporalRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9771),
		"Patient Physiological State Sequence", // PatientPhysiologicalStateSequence
		"PatientPhysiologicalStateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9772),
		"Patient Physiological State Code Sequence", // PatientPhysiologicalStateCodeSequence
		"PatientPhysiologicalStateCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9801),
		"Depth(s) of Focus", // DepthsOfFocus
		"DepthsOfFocus",
		vm.VM1_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9803),
		"Excluded Intervals Sequence", // ExcludedIntervalsSequence
		"ExcludedIntervalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9804),
		"Exclusion Start DateTime", // ExclusionStartDateTime
		"ExclusionStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9805),
		"Exclusion Duration", // ExclusionDuration
		"ExclusionDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9806),
		"US Image Description Sequence", // USImageDescriptionSequence
		"USImageDescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9807),
		"Image Data Type Sequence", // ImageDataTypeSequence
		"ImageDataTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9808),
		"Data Type", // DataType
		"DataType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9809),
		"Transducer Scan Pattern Code Sequence", // TransducerScanPatternCodeSequence
		"TransducerScanPatternCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980B),
		"Aliased Data Type", // AliasedDataType
		"AliasedDataType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980C),
		"Position Measuring Device Used", // PositionMeasuringDeviceUsed
		"PositionMeasuringDeviceUsed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980D),
		"Transducer Geometry Code Sequence", // TransducerGeometryCodeSequence
		"TransducerGeometryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980E),
		"Transducer Beam Steering Code Sequence", // TransducerBeamSteeringCodeSequence
		"TransducerBeamSteeringCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x980F),
		"Transducer Application Code Sequence", // TransducerApplicationCodeSequence
		"TransducerApplicationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9810),
		"Zero Velocity Pixel Value", // ZeroVelocityPixelValue
		"ZeroVelocityPixelValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9821),
		"Photoacoustic Excitation Characteristics Sequence", // PhotoacousticExcitationCharacteristicsSequence
		"PhotoacousticExcitationCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9822),
		"Excitation Spectral Width", // ExcitationSpectralWidth
		"ExcitationSpectralWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9823),
		"Excitation Energy", // ExcitationEnergy
		"ExcitationEnergy",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9824),
		"Excitation Pulse Duration", // ExcitationPulseDuration
		"ExcitationPulseDuration",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9825),
		"Excitation Wavelength Sequence", // ExcitationWavelengthSequence
		"ExcitationWavelengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9826),
		"Excitation Wavelength", // ExcitationWavelength
		"ExcitationWavelength",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9828),
		"Illumination Translation Flag", // IlluminationTranslationFlag
		"IlluminationTranslationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9829),
		"Acoustic Coupling Medium Flag", // AcousticCouplingMediumFlag
		"AcousticCouplingMediumFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982A),
		"Acoustic Coupling Medium Code Sequence", // AcousticCouplingMediumCodeSequence
		"AcousticCouplingMediumCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982B),
		"Acoustic Coupling Medium Temperature", // AcousticCouplingMediumTemperature
		"AcousticCouplingMediumTemperature",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982C),
		"Transducer Response Sequence", // TransducerResponseSequence
		"TransducerResponseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982D),
		"Center Frequency", // CenterFrequency
		"CenterFrequency",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982E),
		"Fractional Bandwidth", // FractionalBandwidth
		"FractionalBandwidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x982F),
		"Lower Cutoff Frequency", // LowerCutoffFrequency
		"LowerCutoffFrequency",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9830),
		"Upper Cutoff Frequency", // UpperCutoffFrequency
		"UpperCutoffFrequency",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9831),
		"Transducer Technology Sequence", // TransducerTechnologySequence
		"TransducerTechnologySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9832),
		"Sound Speed Correction Mechanism Code Sequence", // SoundSpeedCorrectionMechanismCodeSequence
		"SoundSpeedCorrectionMechanismCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9833),
		"Object Sound Speed", // ObjectSoundSpeed
		"ObjectSoundSpeed",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9834),
		"Acoustic Coupling Medium Sound Speed", // AcousticCouplingMediumSoundSpeed
		"AcousticCouplingMediumSoundSpeed",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9835),
		"Photoacoustic Image Frame Type Sequence", // PhotoacousticImageFrameTypeSequence
		"PhotoacousticImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9836),
		"Image Data Type Code Sequence", // ImageDataTypeCodeSequence
		"ImageDataTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9900),
		"Reference Location Label", // ReferenceLocationLabel
		"ReferenceLocationLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9901),
		"Reference Location Description", // ReferenceLocationDescription
		"ReferenceLocationDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9902),
		"Reference Basis Code Sequence", // ReferenceBasisCodeSequence
		"ReferenceBasisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9903),
		"Reference Geometry Code Sequence", // ReferenceGeometryCodeSequence
		"ReferenceGeometryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9904),
		"Offset Distance", // OffsetDistance
		"OffsetDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9905),
		"Offset Direction", // OffsetDirection
		"OffsetDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9906),
		"Potential Scheduled Protocol Code Sequence", // PotentialScheduledProtocolCodeSequence
		"PotentialScheduledProtocolCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9907),
		"Potential Requested Procedure Code Sequence", // PotentialRequestedProcedureCodeSequence
		"PotentialRequestedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9908),
		"Potential Reasons for Procedure", // PotentialReasonsForProcedure
		"PotentialReasonsForProcedure",
		vm.VM1_n,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9909),
		"Potential Reasons for Procedure Code Sequence", // PotentialReasonsForProcedureCodeSequence
		"PotentialReasonsForProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990A),
		"Potential Diagnostic Tasks", // PotentialDiagnosticTasks
		"PotentialDiagnosticTasks",
		vm.VM1_n,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990B),
		"Contraindications Code Sequence", // ContraindicationsCodeSequence
		"ContraindicationsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990C),
		"Referenced Defined Protocol Sequence", // ReferencedDefinedProtocolSequence
		"ReferencedDefinedProtocolSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990D),
		"Referenced Performed Protocol Sequence", // ReferencedPerformedProtocolSequence
		"ReferencedPerformedProtocolSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990E),
		"Predecessor Protocol Sequence", // PredecessorProtocolSequence
		"PredecessorProtocolSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x990F),
		"Protocol Planning Information", // ProtocolPlanningInformation
		"ProtocolPlanningInformation",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9910),
		"Protocol Design Rationale", // ProtocolDesignRationale
		"ProtocolDesignRationale",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9911),
		"Patient Specification Sequence", // PatientSpecificationSequence
		"PatientSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9912),
		"Model Specification Sequence", // ModelSpecificationSequence
		"ModelSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9913),
		"Parameters Specification Sequence", // ParametersSpecificationSequence
		"ParametersSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9914),
		"Instruction Sequence", // InstructionSequence
		"InstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9915),
		"Instruction Index", // InstructionIndex
		"InstructionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9916),
		"Instruction Text", // InstructionText
		"InstructionText",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9917),
		"Instruction Description", // InstructionDescription
		"InstructionDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9918),
		"Instruction Performed Flag", // InstructionPerformedFlag
		"InstructionPerformedFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9919),
		"Instruction Performed DateTime", // InstructionPerformedDateTime
		"InstructionPerformedDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991A),
		"Instruction Performance Comment", // InstructionPerformanceComment
		"InstructionPerformanceComment",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991B),
		"Patient Positioning Instruction Sequence", // PatientPositioningInstructionSequence
		"PatientPositioningInstructionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991C),
		"Positioning Method Code Sequence", // PositioningMethodCodeSequence
		"PositioningMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991D),
		"Positioning Landmark Sequence", // PositioningLandmarkSequence
		"PositioningLandmarkSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991E),
		"Target Frame of Reference UID", // TargetFrameOfReferenceUID
		"TargetFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x991F),
		"Acquisition Protocol Element Specification Sequence", // AcquisitionProtocolElementSpecificationSequence
		"AcquisitionProtocolElementSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9920),
		"Acquisition Protocol Element Sequence", // AcquisitionProtocolElementSequence
		"AcquisitionProtocolElementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9921),
		"Protocol Element Number", // ProtocolElementNumber
		"ProtocolElementNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9922),
		"Protocol Element Name", // ProtocolElementName
		"ProtocolElementName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9923),
		"Protocol Element Characteristics Summary", // ProtocolElementCharacteristicsSummary
		"ProtocolElementCharacteristicsSummary",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9924),
		"Protocol Element Purpose", // ProtocolElementPurpose
		"ProtocolElementPurpose",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9930),
		"Acquisition Motion", // AcquisitionMotion
		"AcquisitionMotion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9931),
		"Acquisition Start Location Sequence", // AcquisitionStartLocationSequence
		"AcquisitionStartLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9932),
		"Acquisition End Location Sequence", // AcquisitionEndLocationSequence
		"AcquisitionEndLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9933),
		"Reconstruction Protocol Element Specification Sequence", // ReconstructionProtocolElementSpecificationSequence
		"ReconstructionProtocolElementSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9934),
		"Reconstruction Protocol Element Sequence", // ReconstructionProtocolElementSequence
		"ReconstructionProtocolElementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9935),
		"Storage Protocol Element Specification Sequence", // StorageProtocolElementSpecificationSequence
		"StorageProtocolElementSpecificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9936),
		"Storage Protocol Element Sequence", // StorageProtocolElementSequence
		"StorageProtocolElementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9937),
		"Requested Series Description", // RequestedSeriesDescription
		"RequestedSeriesDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9938),
		"Source Acquisition Protocol Element Number", // SourceAcquisitionProtocolElementNumber
		"SourceAcquisitionProtocolElementNumber",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9939),
		"Source Acquisition Beam Number", // SourceAcquisitionBeamNumber
		"SourceAcquisitionBeamNumber",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993A),
		"Source Reconstruction Protocol Element Number", // SourceReconstructionProtocolElementNumber
		"SourceReconstructionProtocolElementNumber",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993B),
		"Reconstruction Start Location Sequence", // ReconstructionStartLocationSequence
		"ReconstructionStartLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993C),
		"Reconstruction End Location Sequence", // ReconstructionEndLocationSequence
		"ReconstructionEndLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993D),
		"Reconstruction Algorithm Sequence", // ReconstructionAlgorithmSequence
		"ReconstructionAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x993E),
		"Reconstruction Target Center Location Sequence", // ReconstructionTargetCenterLocationSequence
		"ReconstructionTargetCenterLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9941),
		"Image Filter Description", // ImageFilterDescription
		"ImageFilterDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9942),
		"CTDIvol Notification Trigger", // CTDIvolNotificationTrigger
		"CTDIvolNotificationTrigger",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9943),
		"DLP Notification Trigger", // DLPNotificationTrigger
		"DLPNotificationTrigger",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9944),
		"Auto KVP Selection Type", // AutoKVPSelectionType
		"AutoKVPSelectionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9945),
		"Auto KVP Upper Bound", // AutoKVPUpperBound
		"AutoKVPUpperBound",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9946),
		"Auto KVP Lower Bound", // AutoKVPLowerBound
		"AutoKVPLowerBound",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0x9947),
		"Protocol Defined Patient Position", // ProtocolDefinedPatientPosition
		"ProtocolDefinedPatientPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0xA001),
		"Contributing Equipment Sequence", // ContributingEquipmentSequence
		"ContributingEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0xA002),
		"Contribution DateTime", // ContributionDateTime
		"ContributionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0018, 0xA003),
		"Contribution Description", // ContributionDescription
		"ContributionDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x000D),
		"Study Instance UID", // StudyInstanceUID
		"StudyInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x000E),
		"Series Instance UID", // SeriesInstanceUID
		"SeriesInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0010),
		"Study ID", // StudyID
		"StudyID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0011),
		"Series Number", // SeriesNumber
		"SeriesNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0012),
		"Acquisition Number", // AcquisitionNumber
		"AcquisitionNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0013),
		"Instance Number", // InstanceNumber
		"InstanceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0014),
		"Isotope Number", // IsotopeNumber
		"IsotopeNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0015),
		"Phase Number", // PhaseNumber
		"PhaseNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0016),
		"Interval Number", // IntervalNumber
		"IntervalNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0017),
		"Time Slot Number", // TimeSlotNumber
		"TimeSlotNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0018),
		"Angle Number", // AngleNumber
		"AngleNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0019),
		"Item Number", // ItemNumber
		"ItemNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0020),
		"Patient Orientation", // PatientOrientation
		"PatientOrientation",
		vm.VM2,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0022),
		"Overlay Number", // OverlayNumber
		"OverlayNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0024),
		"Curve Number", // CurveNumber
		"CurveNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0026),
		"LUT Number", // LUTNumber
		"LUTNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0027),
		"Pyramid Label", // PyramidLabel
		"PyramidLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0030),
		"Image Position", // ImagePosition
		"ImagePosition",
		vm.VM3,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0032),
		"Image Position (Patient)", // ImagePositionPatient
		"ImagePositionPatient",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0035),
		"Image Orientation", // ImageOrientation
		"ImageOrientation",
		vm.VM6,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0037),
		"Image Orientation (Patient)", // ImageOrientationPatient
		"ImageOrientationPatient",
		vm.VM6,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0050),
		"Location", // Location
		"Location",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0052),
		"Frame of Reference UID", // FrameOfReferenceUID
		"FrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0060),
		"Laterality", // Laterality
		"Laterality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0062),
		"Image Laterality", // ImageLaterality
		"ImageLaterality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0070),
		"Image Geometry Type", // ImageGeometryType
		"ImageGeometryType",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0080),
		"Masking Image", // MaskingImage
		"MaskingImage",
		vm.VM1_n,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x00AA),
		"Report Number", // ReportNumber
		"ReportNumber",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0100),
		"Temporal Position Identifier", // TemporalPositionIdentifier
		"TemporalPositionIdentifier",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0105),
		"Number of Temporal Positions", // NumberOfTemporalPositions
		"NumberOfTemporalPositions",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0110),
		"Temporal Resolution", // TemporalResolution
		"TemporalResolution",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0200),
		"Synchronization Frame of Reference UID", // SynchronizationFrameOfReferenceUID
		"SynchronizationFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x0242),
		"SOP Instance UID of Concatenation Source", // SOPInstanceUIDOfConcatenationSource
		"SOPInstanceUIDOfConcatenationSource",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1000),
		"Series in Study", // SeriesInStudy
		"SeriesInStudy",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1001),
		"Acquisitions in Series", // AcquisitionsInSeries
		"AcquisitionsInSeries",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1002),
		"Images in Acquisition", // ImagesInAcquisition
		"ImagesInAcquisition",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1003),
		"Images in Series", // ImagesInSeries
		"ImagesInSeries",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1004),
		"Acquisitions in Study", // AcquisitionsInStudy
		"AcquisitionsInStudy",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1005),
		"Images in Study", // ImagesInStudy
		"ImagesInStudy",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1020),
		"Reference", // Reference
		"Reference",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x103F),
		"Target Position Reference Indicator", // TargetPositionReferenceIndicator
		"TargetPositionReferenceIndicator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1040),
		"Position Reference Indicator", // PositionReferenceIndicator
		"PositionReferenceIndicator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1041),
		"Slice Location", // SliceLocation
		"SliceLocation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1070),
		"Other Study Numbers", // OtherStudyNumbers
		"OtherStudyNumbers",
		vm.VM1_n,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1200),
		"Number of Patient Related Studies", // NumberOfPatientRelatedStudies
		"NumberOfPatientRelatedStudies",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1202),
		"Number of Patient Related Series", // NumberOfPatientRelatedSeries
		"NumberOfPatientRelatedSeries",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1204),
		"Number of Patient Related Instances", // NumberOfPatientRelatedInstances
		"NumberOfPatientRelatedInstances",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1206),
		"Number of Study Related Series", // NumberOfStudyRelatedSeries
		"NumberOfStudyRelatedSeries",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1208),
		"Number of Study Related Instances", // NumberOfStudyRelatedInstances
		"NumberOfStudyRelatedInstances",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x1209),
		"Number of Series Related Instances", // NumberOfSeriesRelatedInstances
		"NumberOfSeriesRelatedInstances",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0020,31xx)"),
		"Source Image IDs", // SourceImageIDs
		"SourceImageIDs",
		vm.VM1_n,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3401),
		"Modifying Device ID", // ModifyingDeviceID
		"ModifyingDeviceID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3402),
		"Modified Image ID", // ModifiedImageID
		"ModifiedImageID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3403),
		"Modified Image Date", // ModifiedImageDate
		"ModifiedImageDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3404),
		"Modifying Device Manufacturer", // ModifyingDeviceManufacturer
		"ModifyingDeviceManufacturer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3405),
		"Modified Image Time", // ModifiedImageTime
		"ModifiedImageTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x3406),
		"Modified Image Description", // ModifiedImageDescription
		"ModifiedImageDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x4000),
		"Image Comments", // ImageComments
		"ImageComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x5000),
		"Original Image Identification", // OriginalImageIdentification
		"OriginalImageIdentification",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x5002),
		"Original Image Identification Nomenclature", // OriginalImageIdentificationNomenclature
		"OriginalImageIdentificationNomenclature",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9056),
		"Stack ID", // StackID
		"StackID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9057),
		"In-Stack Position Number", // InStackPositionNumber
		"InStackPositionNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9071),
		"Frame Anatomy Sequence", // FrameAnatomySequence
		"FrameAnatomySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9072),
		"Frame Laterality", // FrameLaterality
		"FrameLaterality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9111),
		"Frame Content Sequence", // FrameContentSequence
		"FrameContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9113),
		"Plane Position Sequence", // PlanePositionSequence
		"PlanePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9116),
		"Plane Orientation Sequence", // PlaneOrientationSequence
		"PlaneOrientationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9128),
		"Temporal Position Index", // TemporalPositionIndex
		"TemporalPositionIndex",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9153),
		"Nominal Cardiac Trigger Delay Time", // NominalCardiacTriggerDelayTime
		"NominalCardiacTriggerDelayTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9154),
		"Nominal Cardiac Trigger Time Prior To R-Peak", // NominalCardiacTriggerTimePriorToRPeak
		"NominalCardiacTriggerTimePriorToRPeak",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9155),
		"Actual Cardiac Trigger Time Prior To R-Peak", // ActualCardiacTriggerTimePriorToRPeak
		"ActualCardiacTriggerTimePriorToRPeak",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9156),
		"Frame Acquisition Number", // FrameAcquisitionNumber
		"FrameAcquisitionNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9157),
		"Dimension Index Values", // DimensionIndexValues
		"DimensionIndexValues",
		vm.VM1_n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9158),
		"Frame Comments", // FrameComments
		"FrameComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9161),
		"Concatenation UID", // ConcatenationUID
		"ConcatenationUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9162),
		"In-concatenation Number", // InConcatenationNumber
		"InConcatenationNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9163),
		"In-concatenation Total Number", // InConcatenationTotalNumber
		"InConcatenationTotalNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9164),
		"Dimension Organization UID", // DimensionOrganizationUID
		"DimensionOrganizationUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9165),
		"Dimension Index Pointer", // DimensionIndexPointer
		"DimensionIndexPointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9167),
		"Functional Group Pointer", // FunctionalGroupPointer
		"FunctionalGroupPointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9170),
		"Unassigned Shared Converted Attributes Sequence", // UnassignedSharedConvertedAttributesSequence
		"UnassignedSharedConvertedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9171),
		"Unassigned Per-Frame Converted Attributes Sequence", // UnassignedPerFrameConvertedAttributesSequence
		"UnassignedPerFrameConvertedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9172),
		"Conversion Source Attributes Sequence", // ConversionSourceAttributesSequence
		"ConversionSourceAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9213),
		"Dimension Index Private Creator", // DimensionIndexPrivateCreator
		"DimensionIndexPrivateCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9221),
		"Dimension Organization Sequence", // DimensionOrganizationSequence
		"DimensionOrganizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9222),
		"Dimension Index Sequence", // DimensionIndexSequence
		"DimensionIndexSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9228),
		"Concatenation Frame Offset Number", // ConcatenationFrameOffsetNumber
		"ConcatenationFrameOffsetNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9238),
		"Functional Group Private Creator", // FunctionalGroupPrivateCreator
		"FunctionalGroupPrivateCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9241),
		"Nominal Percentage of Cardiac Phase", // NominalPercentageOfCardiacPhase
		"NominalPercentageOfCardiacPhase",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9245),
		"Nominal Percentage of Respiratory Phase", // NominalPercentageOfRespiratoryPhase
		"NominalPercentageOfRespiratoryPhase",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9246),
		"Starting Respiratory Amplitude", // StartingRespiratoryAmplitude
		"StartingRespiratoryAmplitude",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9247),
		"Starting Respiratory Phase", // StartingRespiratoryPhase
		"StartingRespiratoryPhase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9248),
		"Ending Respiratory Amplitude", // EndingRespiratoryAmplitude
		"EndingRespiratoryAmplitude",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9249),
		"Ending Respiratory Phase", // EndingRespiratoryPhase
		"EndingRespiratoryPhase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9250),
		"Respiratory Trigger Type", // RespiratoryTriggerType
		"RespiratoryTriggerType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9251),
		"R-R Interval Time Nominal", // RRIntervalTimeNominal
		"RRIntervalTimeNominal",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9252),
		"Actual Cardiac Trigger Delay Time", // ActualCardiacTriggerDelayTime
		"ActualCardiacTriggerDelayTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9253),
		"Respiratory Synchronization Sequence", // RespiratorySynchronizationSequence
		"RespiratorySynchronizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9254),
		"Respiratory Interval Time", // RespiratoryIntervalTime
		"RespiratoryIntervalTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9255),
		"Nominal Respiratory Trigger Delay Time", // NominalRespiratoryTriggerDelayTime
		"NominalRespiratoryTriggerDelayTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9256),
		"Respiratory Trigger Delay Threshold", // RespiratoryTriggerDelayThreshold
		"RespiratoryTriggerDelayThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9257),
		"Actual Respiratory Trigger Delay Time", // ActualRespiratoryTriggerDelayTime
		"ActualRespiratoryTriggerDelayTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9301),
		"Image Position (Volume)", // ImagePositionVolume
		"ImagePositionVolume",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9302),
		"Image Orientation (Volume)", // ImageOrientationVolume
		"ImageOrientationVolume",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9307),
		"Ultrasound Acquisition Geometry", // UltrasoundAcquisitionGeometry
		"UltrasoundAcquisitionGeometry",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9308),
		"Apex Position", // ApexPosition
		"ApexPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9309),
		"Volume to Transducer Mapping Matrix", // VolumeToTransducerMappingMatrix
		"VolumeToTransducerMappingMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930A),
		"Volume to Table Mapping Matrix", // VolumeToTableMappingMatrix
		"VolumeToTableMappingMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930B),
		"Volume to Transducer Relationship", // VolumeToTransducerRelationship
		"VolumeToTransducerRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930C),
		"Patient Frame of Reference Source", // PatientFrameOfReferenceSource
		"PatientFrameOfReferenceSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930D),
		"Temporal Position Time Offset", // TemporalPositionTimeOffset
		"TemporalPositionTimeOffset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930E),
		"Plane Position (Volume) Sequence", // PlanePositionVolumeSequence
		"PlanePositionVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x930F),
		"Plane Orientation (Volume) Sequence", // PlaneOrientationVolumeSequence
		"PlaneOrientationVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9310),
		"Temporal Position Sequence", // TemporalPositionSequence
		"TemporalPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9311),
		"Dimension Organization Type", // DimensionOrganizationType
		"DimensionOrganizationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9312),
		"Volume Frame of Reference UID", // VolumeFrameOfReferenceUID
		"VolumeFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9313),
		"Table Frame of Reference UID", // TableFrameOfReferenceUID
		"TableFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9421),
		"Dimension Description Label", // DimensionDescriptionLabel
		"DimensionDescriptionLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9450),
		"Patient Orientation in Frame Sequence", // PatientOrientationInFrameSequence
		"PatientOrientationInFrameSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9453),
		"Frame Label", // FrameLabel
		"FrameLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9518),
		"Acquisition Index", // AcquisitionIndex
		"AcquisitionIndex",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9529),
		"Contributing SOP Instances Reference Sequence", // ContributingSOPInstancesReferenceSequence
		"ContributingSOPInstancesReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0020, 0x9536),
		"Reconstruction Index", // ReconstructionIndex
		"ReconstructionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0001),
		"Light Path Filter Pass-Through Wavelength", // LightPathFilterPassThroughWavelength
		"LightPathFilterPassThroughWavelength",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0002),
		"Light Path Filter Pass Band", // LightPathFilterPassBand
		"LightPathFilterPassBand",
		vm.VM2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0003),
		"Image Path Filter Pass-Through Wavelength", // ImagePathFilterPassThroughWavelength
		"ImagePathFilterPassThroughWavelength",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0004),
		"Image Path Filter Pass Band", // ImagePathFilterPassBand
		"ImagePathFilterPassBand",
		vm.VM2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0005),
		"Patient Eye Movement Commanded", // PatientEyeMovementCommanded
		"PatientEyeMovementCommanded",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0006),
		"Patient Eye Movement Command Code Sequence", // PatientEyeMovementCommandCodeSequence
		"PatientEyeMovementCommandCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0007),
		"Spherical Lens Power", // SphericalLensPower
		"SphericalLensPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0008),
		"Cylinder Lens Power", // CylinderLensPower
		"CylinderLensPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0009),
		"Cylinder Axis", // CylinderAxis
		"CylinderAxis",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000A),
		"Emmetropic Magnification", // EmmetropicMagnification
		"EmmetropicMagnification",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000B),
		"Intra Ocular Pressure", // IntraOcularPressure
		"IntraOcularPressure",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000C),
		"Horizontal Field of View", // HorizontalFieldOfView
		"HorizontalFieldOfView",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000D),
		"Pupil Dilated", // PupilDilated
		"PupilDilated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000E),
		"Degree of Dilation", // DegreeOfDilation
		"DegreeOfDilation",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x000F),
		"Vertex Distance", // VertexDistance
		"VertexDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0010),
		"Stereo Baseline Angle", // StereoBaselineAngle
		"StereoBaselineAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0011),
		"Stereo Baseline Displacement", // StereoBaselineDisplacement
		"StereoBaselineDisplacement",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0012),
		"Stereo Horizontal Pixel Offset", // StereoHorizontalPixelOffset
		"StereoHorizontalPixelOffset",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0013),
		"Stereo Vertical Pixel Offset", // StereoVerticalPixelOffset
		"StereoVerticalPixelOffset",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0014),
		"Stereo Rotation", // StereoRotation
		"StereoRotation",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0015),
		"Acquisition Device Type Code Sequence", // AcquisitionDeviceTypeCodeSequence
		"AcquisitionDeviceTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0016),
		"Illumination Type Code Sequence", // IlluminationTypeCodeSequence
		"IlluminationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0017),
		"Light Path Filter Type Stack Code Sequence", // LightPathFilterTypeStackCodeSequence
		"LightPathFilterTypeStackCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0018),
		"Image Path Filter Type Stack Code Sequence", // ImagePathFilterTypeStackCodeSequence
		"ImagePathFilterTypeStackCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0019),
		"Lenses Code Sequence", // LensesCodeSequence
		"LensesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001A),
		"Channel Description Code Sequence", // ChannelDescriptionCodeSequence
		"ChannelDescriptionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001B),
		"Refractive State Sequence", // RefractiveStateSequence
		"RefractiveStateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001C),
		"Mydriatic Agent Code Sequence", // MydriaticAgentCodeSequence
		"MydriaticAgentCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001D),
		"Relative Image Position Code Sequence", // RelativeImagePositionCodeSequence
		"RelativeImagePositionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x001E),
		"Camera Angle of View", // CameraAngleOfView
		"CameraAngleOfView",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0020),
		"Stereo Pairs Sequence", // StereoPairsSequence
		"StereoPairsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0021),
		"Left Image Sequence", // LeftImageSequence
		"LeftImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0022),
		"Right Image Sequence", // RightImageSequence
		"RightImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0028),
		"Stereo Pairs Present", // StereoPairsPresent
		"StereoPairsPresent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0030),
		"Axial Length of the Eye", // AxialLengthOfTheEye
		"AxialLengthOfTheEye",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0031),
		"Ophthalmic Frame Location Sequence", // OphthalmicFrameLocationSequence
		"OphthalmicFrameLocationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0032),
		"Reference Coordinates", // ReferenceCoordinates
		"ReferenceCoordinates",
		vm.VM2_2n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0035),
		"Depth Spatial Resolution", // DepthSpatialResolution
		"DepthSpatialResolution",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0036),
		"Maximum Depth Distortion", // MaximumDepthDistortion
		"MaximumDepthDistortion",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0037),
		"Along-scan Spatial Resolution", // AlongScanSpatialResolution
		"AlongScanSpatialResolution",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0038),
		"Maximum Along-scan Distortion", // MaximumAlongScanDistortion
		"MaximumAlongScanDistortion",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0039),
		"Ophthalmic Image Orientation", // OphthalmicImageOrientation
		"OphthalmicImageOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0041),
		"Depth of Transverse Image", // DepthOfTransverseImage
		"DepthOfTransverseImage",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0042),
		"Mydriatic Agent Concentration Units Sequence", // MydriaticAgentConcentrationUnitsSequence
		"MydriaticAgentConcentrationUnitsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0048),
		"Across-scan Spatial Resolution", // AcrossScanSpatialResolution
		"AcrossScanSpatialResolution",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0049),
		"Maximum Across-scan Distortion", // MaximumAcrossScanDistortion
		"MaximumAcrossScanDistortion",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x004E),
		"Mydriatic Agent Concentration", // MydriaticAgentConcentration
		"MydriaticAgentConcentration",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0055),
		"Illumination Wave Length", // IlluminationWaveLength
		"IlluminationWaveLength",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0056),
		"Illumination Power", // IlluminationPower
		"IlluminationPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0057),
		"Illumination Bandwidth", // IlluminationBandwidth
		"IlluminationBandwidth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x0058),
		"Mydriatic Agent Sequence", // MydriaticAgentSequence
		"MydriaticAgentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1007),
		"Ophthalmic Axial Measurements Right Eye Sequence", // OphthalmicAxialMeasurementsRightEyeSequence
		"OphthalmicAxialMeasurementsRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1008),
		"Ophthalmic Axial Measurements Left Eye Sequence", // OphthalmicAxialMeasurementsLeftEyeSequence
		"OphthalmicAxialMeasurementsLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1009),
		"Ophthalmic Axial Measurements Device Type", // OphthalmicAxialMeasurementsDeviceType
		"OphthalmicAxialMeasurementsDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1010),
		"Ophthalmic Axial Length Measurements Type", // OphthalmicAxialLengthMeasurementsType
		"OphthalmicAxialLengthMeasurementsType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1012),
		"Ophthalmic Axial Length Sequence", // OphthalmicAxialLengthSequence
		"OphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1019),
		"Ophthalmic Axial Length", // OphthalmicAxialLength
		"OphthalmicAxialLength",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1024),
		"Lens Status Code Sequence", // LensStatusCodeSequence
		"LensStatusCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1025),
		"Vitreous Status Code Sequence", // VitreousStatusCodeSequence
		"VitreousStatusCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1028),
		"IOL Formula Code Sequence", // IOLFormulaCodeSequence
		"IOLFormulaCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1029),
		"IOL Formula Detail", // IOLFormulaDetail
		"IOLFormulaDetail",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1033),
		"Keratometer Index", // KeratometerIndex
		"KeratometerIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1035),
		"Source of Ophthalmic Axial Length Code Sequence", // SourceOfOphthalmicAxialLengthCodeSequence
		"SourceOfOphthalmicAxialLengthCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1036),
		"Source of Corneal Size Data Code Sequence", // SourceOfCornealSizeDataCodeSequence
		"SourceOfCornealSizeDataCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1037),
		"Target Refraction", // TargetRefraction
		"TargetRefraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1039),
		"Refractive Procedure Occurred", // RefractiveProcedureOccurred
		"RefractiveProcedureOccurred",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1040),
		"Refractive Surgery Type Code Sequence", // RefractiveSurgeryTypeCodeSequence
		"RefractiveSurgeryTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1044),
		"Ophthalmic Ultrasound Method Code Sequence", // OphthalmicUltrasoundMethodCodeSequence
		"OphthalmicUltrasoundMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1045),
		"Surgically Induced Astigmatism Sequence", // SurgicallyInducedAstigmatismSequence
		"SurgicallyInducedAstigmatismSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1046),
		"Type of Optical Correction", // TypeOfOpticalCorrection
		"TypeOfOpticalCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1047),
		"Toric IOL Power Sequence", // ToricIOLPowerSequence
		"ToricIOLPowerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1048),
		"Predicted Toric Error Sequence", // PredictedToricErrorSequence
		"PredictedToricErrorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1049),
		"Pre-Selected for Implantation", // PreSelectedForImplantation
		"PreSelectedForImplantation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x104A),
		"Toric IOL Power for Exact Emmetropia Sequence", // ToricIOLPowerForExactEmmetropiaSequence
		"ToricIOLPowerForExactEmmetropiaSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x104B),
		"Toric IOL Power for Exact Target Refraction Sequence", // ToricIOLPowerForExactTargetRefractionSequence
		"ToricIOLPowerForExactTargetRefractionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1050),
		"Ophthalmic Axial Length Measurements Sequence", // OphthalmicAxialLengthMeasurementsSequence
		"OphthalmicAxialLengthMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1053),
		"IOL Power", // IOLPower
		"IOLPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1054),
		"Predicted Refractive Error", // PredictedRefractiveError
		"PredictedRefractiveError",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1059),
		"Ophthalmic Axial Length Velocity", // OphthalmicAxialLengthVelocity
		"OphthalmicAxialLengthVelocity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1065),
		"Lens Status Description", // LensStatusDescription
		"LensStatusDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1066),
		"Vitreous Status Description", // VitreousStatusDescription
		"VitreousStatusDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1090),
		"IOL Power Sequence", // IOLPowerSequence
		"IOLPowerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1092),
		"Lens Constant Sequence", // LensConstantSequence
		"LensConstantSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1093),
		"IOL Manufacturer", // IOLManufacturer
		"IOLManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1094),
		"Lens Constant Description", // LensConstantDescription
		"LensConstantDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1095),
		"Implant Name", // ImplantName
		"ImplantName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1096),
		"Keratometry Measurement Type Code Sequence", // KeratometryMeasurementTypeCodeSequence
		"KeratometryMeasurementTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1097),
		"Implant Part Number", // ImplantPartNumber
		"ImplantPartNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1100),
		"Referenced Ophthalmic Axial Measurements Sequence", // ReferencedOphthalmicAxialMeasurementsSequence
		"ReferencedOphthalmicAxialMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1101),
		"Ophthalmic Axial Length Measurements Segment Name Code Sequence", // OphthalmicAxialLengthMeasurementsSegmentNameCodeSequence
		"OphthalmicAxialLengthMeasurementsSegmentNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1103),
		"Refractive Error Before Refractive Surgery Code Sequence", // RefractiveErrorBeforeRefractiveSurgeryCodeSequence
		"RefractiveErrorBeforeRefractiveSurgeryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1121),
		"IOL Power For Exact Emmetropia", // IOLPowerForExactEmmetropia
		"IOLPowerForExactEmmetropia",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1122),
		"IOL Power For Exact Target Refraction", // IOLPowerForExactTargetRefraction
		"IOLPowerForExactTargetRefraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1125),
		"Anterior Chamber Depth Definition Code Sequence", // AnteriorChamberDepthDefinitionCodeSequence
		"AnteriorChamberDepthDefinitionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1127),
		"Lens Thickness Sequence", // LensThicknessSequence
		"LensThicknessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1128),
		"Anterior Chamber Depth Sequence", // AnteriorChamberDepthSequence
		"AnteriorChamberDepthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x112A),
		"Calculation Comment Sequence", // CalculationCommentSequence
		"CalculationCommentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x112B),
		"Calculation Comment Type", // CalculationCommentType
		"CalculationCommentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x112C),
		"Calculation Comment", // CalculationComment
		"CalculationComment",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1130),
		"Lens Thickness", // LensThickness
		"LensThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1131),
		"Anterior Chamber Depth", // AnteriorChamberDepth
		"AnteriorChamberDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1132),
		"Source of Lens Thickness Data Code Sequence", // SourceOfLensThicknessDataCodeSequence
		"SourceOfLensThicknessDataCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1133),
		"Source of Anterior Chamber Depth Data Code Sequence", // SourceOfAnteriorChamberDepthDataCodeSequence
		"SourceOfAnteriorChamberDepthDataCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1134),
		"Source of Refractive Measurements Sequence", // SourceOfRefractiveMeasurementsSequence
		"SourceOfRefractiveMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1135),
		"Source of Refractive Measurements Code Sequence", // SourceOfRefractiveMeasurementsCodeSequence
		"SourceOfRefractiveMeasurementsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1140),
		"Ophthalmic Axial Length Measurement Modified", // OphthalmicAxialLengthMeasurementModified
		"OphthalmicAxialLengthMeasurementModified",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1150),
		"Ophthalmic Axial Length Data Source Code Sequence", // OphthalmicAxialLengthDataSourceCodeSequence
		"OphthalmicAxialLengthDataSourceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1153),
		"Ophthalmic Axial Length Acquisition Method Code Sequence", // OphthalmicAxialLengthAcquisitionMethodCodeSequence
		"OphthalmicAxialLengthAcquisitionMethodCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1155),
		"Signal to Noise Ratio", // SignalToNoiseRatio
		"SignalToNoiseRatio",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1159),
		"Ophthalmic Axial Length Data Source Description", // OphthalmicAxialLengthDataSourceDescription
		"OphthalmicAxialLengthDataSourceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1210),
		"Ophthalmic Axial Length Measurements Total Length Sequence", // OphthalmicAxialLengthMeasurementsTotalLengthSequence
		"OphthalmicAxialLengthMeasurementsTotalLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1211),
		"Ophthalmic Axial Length Measurements Segmental Length Sequence", // OphthalmicAxialLengthMeasurementsSegmentalLengthSequence
		"OphthalmicAxialLengthMeasurementsSegmentalLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1212),
		"Ophthalmic Axial Length Measurements Length Summation Sequence", // OphthalmicAxialLengthMeasurementsLengthSummationSequence
		"OphthalmicAxialLengthMeasurementsLengthSummationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1220),
		"Ultrasound Ophthalmic Axial Length Measurements Sequence", // UltrasoundOphthalmicAxialLengthMeasurementsSequence
		"UltrasoundOphthalmicAxialLengthMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1225),
		"Optical Ophthalmic Axial Length Measurements Sequence", // OpticalOphthalmicAxialLengthMeasurementsSequence
		"OpticalOphthalmicAxialLengthMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1230),
		"Ultrasound Selected Ophthalmic Axial Length Sequence", // UltrasoundSelectedOphthalmicAxialLengthSequence
		"UltrasoundSelectedOphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1250),
		"Ophthalmic Axial Length Selection Method Code Sequence", // OphthalmicAxialLengthSelectionMethodCodeSequence
		"OphthalmicAxialLengthSelectionMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1255),
		"Optical Selected Ophthalmic Axial Length Sequence", // OpticalSelectedOphthalmicAxialLengthSequence
		"OpticalSelectedOphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1257),
		"Selected Segmental Ophthalmic Axial Length Sequence", // SelectedSegmentalOphthalmicAxialLengthSequence
		"SelectedSegmentalOphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1260),
		"Selected Total Ophthalmic Axial Length Sequence", // SelectedTotalOphthalmicAxialLengthSequence
		"SelectedTotalOphthalmicAxialLengthSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1262),
		"Ophthalmic Axial Length Quality Metric Sequence", // OphthalmicAxialLengthQualityMetricSequence
		"OphthalmicAxialLengthQualityMetricSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1265),
		"Ophthalmic Axial Length Quality Metric Type Code Sequence", // OphthalmicAxialLengthQualityMetricTypeCodeSequence
		"OphthalmicAxialLengthQualityMetricTypeCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1273),
		"Ophthalmic Axial Length Quality Metric Type Description", // OphthalmicAxialLengthQualityMetricTypeDescription
		"OphthalmicAxialLengthQualityMetricTypeDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1300),
		"Intraocular Lens Calculations Right Eye Sequence", // IntraocularLensCalculationsRightEyeSequence
		"IntraocularLensCalculationsRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1310),
		"Intraocular Lens Calculations Left Eye Sequence", // IntraocularLensCalculationsLeftEyeSequence
		"IntraocularLensCalculationsLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1330),
		"Referenced Ophthalmic Axial Length Measurement QC Image Sequence", // ReferencedOphthalmicAxialLengthMeasurementQCImageSequence
		"ReferencedOphthalmicAxialLengthMeasurementQCImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1415),
		"Ophthalmic Mapping Device Type", // OphthalmicMappingDeviceType
		"OphthalmicMappingDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1420),
		"Acquisition Method Code Sequence", // AcquisitionMethodCodeSequence
		"AcquisitionMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1423),
		"Acquisition Method Algorithm Sequence", // AcquisitionMethodAlgorithmSequence
		"AcquisitionMethodAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1436),
		"Ophthalmic Thickness Map Type Code Sequence", // OphthalmicThicknessMapTypeCodeSequence
		"OphthalmicThicknessMapTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1443),
		"Ophthalmic Thickness Mapping Normals Sequence", // OphthalmicThicknessMappingNormalsSequence
		"OphthalmicThicknessMappingNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1445),
		"Retinal Thickness Definition Code Sequence", // RetinalThicknessDefinitionCodeSequence
		"RetinalThicknessDefinitionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1450),
		"Pixel Value Mapping to Coded Concept Sequence", // PixelValueMappingToCodedConceptSequence
		"PixelValueMappingToCodedConceptSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1452),
		"Mapped Pixel Value", // MappedPixelValue
		"MappedPixelValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1454),
		"Pixel Value Mapping Explanation", // PixelValueMappingExplanation
		"PixelValueMappingExplanation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1458),
		"Ophthalmic Thickness Map Quality Threshold Sequence", // OphthalmicThicknessMapQualityThresholdSequence
		"OphthalmicThicknessMapQualityThresholdSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1460),
		"Ophthalmic Thickness Map Threshold Quality Rating", // OphthalmicThicknessMapThresholdQualityRating
		"OphthalmicThicknessMapThresholdQualityRating",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1463),
		"Anatomic Structure Reference Point", // AnatomicStructureReferencePoint
		"AnatomicStructureReferencePoint",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1465),
		"Registration to Localizer Sequence", // RegistrationToLocalizerSequence
		"RegistrationToLocalizerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1466),
		"Registered Localizer Units", // RegisteredLocalizerUnits
		"RegisteredLocalizerUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1467),
		"Registered Localizer Top Left Hand Corner", // RegisteredLocalizerTopLeftHandCorner
		"RegisteredLocalizerTopLeftHandCorner",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1468),
		"Registered Localizer Bottom Right Hand Corner", // RegisteredLocalizerBottomRightHandCorner
		"RegisteredLocalizerBottomRightHandCorner",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1470),
		"Ophthalmic Thickness Map Quality Rating Sequence", // OphthalmicThicknessMapQualityRatingSequence
		"OphthalmicThicknessMapQualityRatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1472),
		"Relevant OPT Attributes Sequence", // RelevantOPTAttributesSequence
		"RelevantOPTAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1512),
		"Transformation Method Code Sequence", // TransformationMethodCodeSequence
		"TransformationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1513),
		"Transformation Algorithm Sequence", // TransformationAlgorithmSequence
		"TransformationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1515),
		"Ophthalmic Axial Length Method", // OphthalmicAxialLengthMethod
		"OphthalmicAxialLengthMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1517),
		"Ophthalmic FOV", // OphthalmicFOV
		"OphthalmicFOV",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1518),
		"Two Dimensional to Three Dimensional Map Sequence", // TwoDimensionalToThreeDimensionalMapSequence
		"TwoDimensionalToThreeDimensionalMapSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1525),
		"Wide Field Ophthalmic Photography Quality Rating Sequence", // WideFieldOphthalmicPhotographyQualityRatingSequence
		"WideFieldOphthalmicPhotographyQualityRatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1526),
		"Wide Field Ophthalmic Photography Quality Threshold Sequence", // WideFieldOphthalmicPhotographyQualityThresholdSequence
		"WideFieldOphthalmicPhotographyQualityThresholdSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1527),
		"Wide Field Ophthalmic Photography Threshold Quality Rating", // WideFieldOphthalmicPhotographyThresholdQualityRating
		"WideFieldOphthalmicPhotographyThresholdQualityRating",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1528),
		"X Coordinates Center Pixel View Angle", // XCoordinatesCenterPixelViewAngle
		"XCoordinatesCenterPixelViewAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1529),
		"Y Coordinates Center Pixel View Angle", // YCoordinatesCenterPixelViewAngle
		"YCoordinatesCenterPixelViewAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1530),
		"Number of Map Points", // NumberOfMapPoints
		"NumberOfMapPoints",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1531),
		"Two Dimensional to Three Dimensional Map Data", // TwoDimensionalToThreeDimensionalMapData
		"TwoDimensionalToThreeDimensionalMapData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1612),
		"Derivation Algorithm Sequence", // DerivationAlgorithmSequence
		"DerivationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1615),
		"Ophthalmic Image Type Code Sequence", // OphthalmicImageTypeCodeSequence
		"OphthalmicImageTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1616),
		"Ophthalmic Image Type Description", // OphthalmicImageTypeDescription
		"OphthalmicImageTypeDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1618),
		"Scan Pattern Type Code Sequence", // ScanPatternTypeCodeSequence
		"ScanPatternTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1620),
		"Referenced Surface Mesh Identification Sequence", // ReferencedSurfaceMeshIdentificationSequence
		"ReferencedSurfaceMeshIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1622),
		"Ophthalmic Volumetric Properties Flag", // OphthalmicVolumetricPropertiesFlag
		"OphthalmicVolumetricPropertiesFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1623),
		"Ophthalmic Anatomic Reference Point Frame Coordinate", // OphthalmicAnatomicReferencePointFrameCoordinate
		"OphthalmicAnatomicReferencePointFrameCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1624),
		"Ophthalmic Anatomic Reference Point X-Coordinate", // OphthalmicAnatomicReferencePointXCoordinate
		"OphthalmicAnatomicReferencePointXCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1626),
		"Ophthalmic Anatomic Reference Point Y-Coordinate", // OphthalmicAnatomicReferencePointYCoordinate
		"OphthalmicAnatomicReferencePointYCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1627),
		"Ophthalmic En Face Volume Descriptor Sequence", // OphthalmicEnFaceVolumeDescriptorSequence
		"OphthalmicEnFaceVolumeDescriptorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1628),
		"Ophthalmic En Face Image Quality Rating Sequence", // OphthalmicEnFaceImageQualityRatingSequence
		"OphthalmicEnFaceImageQualityRatingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1629),
		"Ophthalmic En Face Volume Descriptor Scope", // OphthalmicEnFaceVolumeDescriptorScope
		"OphthalmicEnFaceVolumeDescriptorScope",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1630),
		"Quality Threshold", // QualityThreshold
		"QualityThreshold",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1632),
		"Ophthalmic Anatomic Reference Point Sequence", // OphthalmicAnatomicReferencePointSequence
		"OphthalmicAnatomicReferencePointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1633),
		"Ophthalmic Anatomic Reference Point Localization Type", // OphthalmicAnatomicReferencePointLocalizationType
		"OphthalmicAnatomicReferencePointLocalizationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1634),
		"Primary Anatomic Structure Item Index", // PrimaryAnatomicStructureItemIndex
		"PrimaryAnatomicStructureItemIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1640),
		"OCT B-scan Analysis Acquisition Parameters Sequence", // OCTBscanAnalysisAcquisitionParametersSequence
		"OCTBscanAnalysisAcquisitionParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1642),
		"Number of B-scans Per Frame", // NumberOfBscansPerFrame
		"NumberOfBscansPerFrame",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1643),
		"B-scan Slab Thickness", // BscanSlabThickness
		"BscanSlabThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1644),
		"Distance Between B-scan Slabs", // DistanceBetweenBscanSlabs
		"DistanceBetweenBscanSlabs",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1645),
		"B-scan Cycle Time", // BscanCycleTime
		"BscanCycleTime",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1646),
		"B-scan Cycle Time Vector", // BscanCycleTimeVector
		"BscanCycleTimeVector",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1649),
		"A-scan Rate", // AscanRate
		"AscanRate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1650),
		"B-scan Rate", // BscanRate
		"BscanRate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0022, 0x1658),
		"Surface Mesh Z-Pixel Offset", // SurfaceMeshZPixelOffset
		"SurfaceMeshZPixelOffset",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0010),
		"Visual Field Horizontal Extent", // VisualFieldHorizontalExtent
		"VisualFieldHorizontalExtent",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0011),
		"Visual Field Vertical Extent", // VisualFieldVerticalExtent
		"VisualFieldVerticalExtent",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0012),
		"Visual Field Shape", // VisualFieldShape
		"VisualFieldShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0016),
		"Screening Test Mode Code Sequence", // ScreeningTestModeCodeSequence
		"ScreeningTestModeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0018),
		"Maximum Stimulus Luminance", // MaximumStimulusLuminance
		"MaximumStimulusLuminance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0020),
		"Background Luminance", // BackgroundLuminance
		"BackgroundLuminance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0021),
		"Stimulus Color Code Sequence", // StimulusColorCodeSequence
		"StimulusColorCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0024),
		"Background Illumination Color Code Sequence", // BackgroundIlluminationColorCodeSequence
		"BackgroundIlluminationColorCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0025),
		"Stimulus Area", // StimulusArea
		"StimulusArea",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0028),
		"Stimulus Presentation Time", // StimulusPresentationTime
		"StimulusPresentationTime",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0032),
		"Fixation Sequence", // FixationSequence
		"FixationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0033),
		"Fixation Monitoring Code Sequence", // FixationMonitoringCodeSequence
		"FixationMonitoringCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0034),
		"Visual Field Catch Trial Sequence", // VisualFieldCatchTrialSequence
		"VisualFieldCatchTrialSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0035),
		"Fixation Checked Quantity", // FixationCheckedQuantity
		"FixationCheckedQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0036),
		"Patient Not Properly Fixated Quantity", // PatientNotProperlyFixatedQuantity
		"PatientNotProperlyFixatedQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0037),
		"Presented Visual Stimuli Data Flag", // PresentedVisualStimuliDataFlag
		"PresentedVisualStimuliDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0038),
		"Number of Visual Stimuli", // NumberOfVisualStimuli
		"NumberOfVisualStimuli",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0039),
		"Excessive Fixation Losses Data Flag", // ExcessiveFixationLossesDataFlag
		"ExcessiveFixationLossesDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0040),
		"Excessive Fixation Losses", // ExcessiveFixationLosses
		"ExcessiveFixationLosses",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0042),
		"Stimuli Retesting Quantity", // StimuliRetestingQuantity
		"StimuliRetestingQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0044),
		"Comments on Patient's Performance of Visual Field", // CommentsOnPatientPerformanceOfVisualField
		"CommentsOnPatientPerformanceOfVisualField",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0045),
		"False Negatives Estimate Flag", // FalseNegativesEstimateFlag
		"FalseNegativesEstimateFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0046),
		"False Negatives Estimate", // FalseNegativesEstimate
		"FalseNegativesEstimate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0048),
		"Negative Catch Trials Quantity", // NegativeCatchTrialsQuantity
		"NegativeCatchTrialsQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0050),
		"False Negatives Quantity", // FalseNegativesQuantity
		"FalseNegativesQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0051),
		"Excessive False Negatives Data Flag", // ExcessiveFalseNegativesDataFlag
		"ExcessiveFalseNegativesDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0052),
		"Excessive False Negatives", // ExcessiveFalseNegatives
		"ExcessiveFalseNegatives",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0053),
		"False Positives Estimate Flag", // FalsePositivesEstimateFlag
		"FalsePositivesEstimateFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0054),
		"False Positives Estimate", // FalsePositivesEstimate
		"FalsePositivesEstimate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0055),
		"Catch Trials Data Flag", // CatchTrialsDataFlag
		"CatchTrialsDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0056),
		"Positive Catch Trials Quantity", // PositiveCatchTrialsQuantity
		"PositiveCatchTrialsQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0057),
		"Test Point Normals Data Flag", // TestPointNormalsDataFlag
		"TestPointNormalsDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0058),
		"Test Point Normals Sequence", // TestPointNormalsSequence
		"TestPointNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0059),
		"Global Deviation Probability Normals Flag", // GlobalDeviationProbabilityNormalsFlag
		"GlobalDeviationProbabilityNormalsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0060),
		"False Positives Quantity", // FalsePositivesQuantity
		"FalsePositivesQuantity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0061),
		"Excessive False Positives Data Flag", // ExcessiveFalsePositivesDataFlag
		"ExcessiveFalsePositivesDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0062),
		"Excessive False Positives", // ExcessiveFalsePositives
		"ExcessiveFalsePositives",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0063),
		"Visual Field Test Normals Flag", // VisualFieldTestNormalsFlag
		"VisualFieldTestNormalsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0064),
		"Results Normals Sequence", // ResultsNormalsSequence
		"ResultsNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0065),
		"Age Corrected Sensitivity Deviation Algorithm Sequence", // AgeCorrectedSensitivityDeviationAlgorithmSequence
		"AgeCorrectedSensitivityDeviationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0066),
		"Global Deviation From Normal", // GlobalDeviationFromNormal
		"GlobalDeviationFromNormal",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0067),
		"Generalized Defect Sensitivity Deviation Algorithm Sequence", // GeneralizedDefectSensitivityDeviationAlgorithmSequence
		"GeneralizedDefectSensitivityDeviationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0068),
		"Localized Deviation From Normal", // LocalizedDeviationFromNormal
		"LocalizedDeviationFromNormal",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0069),
		"Patient Reliability Indicator", // PatientReliabilityIndicator
		"PatientReliabilityIndicator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0070),
		"Visual Field Mean Sensitivity", // VisualFieldMeanSensitivity
		"VisualFieldMeanSensitivity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0071),
		"Global Deviation Probability", // GlobalDeviationProbability
		"GlobalDeviationProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0072),
		"Local Deviation Probability Normals Flag", // LocalDeviationProbabilityNormalsFlag
		"LocalDeviationProbabilityNormalsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0073),
		"Localized Deviation Probability", // LocalizedDeviationProbability
		"LocalizedDeviationProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0074),
		"Short Term Fluctuation Calculated", // ShortTermFluctuationCalculated
		"ShortTermFluctuationCalculated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0075),
		"Short Term Fluctuation", // ShortTermFluctuation
		"ShortTermFluctuation",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0076),
		"Short Term Fluctuation Probability Calculated", // ShortTermFluctuationProbabilityCalculated
		"ShortTermFluctuationProbabilityCalculated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0077),
		"Short Term Fluctuation Probability", // ShortTermFluctuationProbability
		"ShortTermFluctuationProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0078),
		"Corrected Localized Deviation From Normal Calculated", // CorrectedLocalizedDeviationFromNormalCalculated
		"CorrectedLocalizedDeviationFromNormalCalculated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0079),
		"Corrected Localized Deviation From Normal", // CorrectedLocalizedDeviationFromNormal
		"CorrectedLocalizedDeviationFromNormal",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0080),
		"Corrected Localized Deviation From Normal Probability Calculated", // CorrectedLocalizedDeviationFromNormalProbabilityCalculated
		"CorrectedLocalizedDeviationFromNormalProbabilityCalculated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0081),
		"Corrected Localized Deviation From Normal Probability", // CorrectedLocalizedDeviationFromNormalProbability
		"CorrectedLocalizedDeviationFromNormalProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0083),
		"Global Deviation Probability Sequence", // GlobalDeviationProbabilitySequence
		"GlobalDeviationProbabilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0085),
		"Localized Deviation Probability Sequence", // LocalizedDeviationProbabilitySequence
		"LocalizedDeviationProbabilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0086),
		"Foveal Sensitivity Measured", // FovealSensitivityMeasured
		"FovealSensitivityMeasured",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0087),
		"Foveal Sensitivity", // FovealSensitivity
		"FovealSensitivity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0088),
		"Visual Field Test Duration", // VisualFieldTestDuration
		"VisualFieldTestDuration",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0089),
		"Visual Field Test Point Sequence", // VisualFieldTestPointSequence
		"VisualFieldTestPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0090),
		"Visual Field Test Point X-Coordinate", // VisualFieldTestPointXCoordinate
		"VisualFieldTestPointXCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0091),
		"Visual Field Test Point Y-Coordinate", // VisualFieldTestPointYCoordinate
		"VisualFieldTestPointYCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0092),
		"Age Corrected Sensitivity Deviation Value", // AgeCorrectedSensitivityDeviationValue
		"AgeCorrectedSensitivityDeviationValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0093),
		"Stimulus Results", // StimulusResults
		"StimulusResults",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0094),
		"Sensitivity Value", // SensitivityValue
		"SensitivityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0095),
		"Retest Stimulus Seen", // RetestStimulusSeen
		"RetestStimulusSeen",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0096),
		"Retest Sensitivity Value", // RetestSensitivityValue
		"RetestSensitivityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0097),
		"Visual Field Test Point Normals Sequence", // VisualFieldTestPointNormalsSequence
		"VisualFieldTestPointNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0098),
		"Quantified Defect", // QuantifiedDefect
		"QuantifiedDefect",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0100),
		"Age Corrected Sensitivity Deviation Probability Value", // AgeCorrectedSensitivityDeviationProbabilityValue
		"AgeCorrectedSensitivityDeviationProbabilityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0102),
		"Generalized Defect Corrected Sensitivity Deviation Flag", // GeneralizedDefectCorrectedSensitivityDeviationFlag
		"GeneralizedDefectCorrectedSensitivityDeviationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0103),
		"Generalized Defect Corrected Sensitivity Deviation Value", // GeneralizedDefectCorrectedSensitivityDeviationValue
		"GeneralizedDefectCorrectedSensitivityDeviationValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0104),
		"Generalized Defect Corrected Sensitivity Deviation Probability Value", // GeneralizedDefectCorrectedSensitivityDeviationProbabilityValue
		"GeneralizedDefectCorrectedSensitivityDeviationProbabilityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0105),
		"Minimum Sensitivity Value", // MinimumSensitivityValue
		"MinimumSensitivityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0106),
		"Blind Spot Localized", // BlindSpotLocalized
		"BlindSpotLocalized",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0107),
		"Blind Spot X-Coordinate", // BlindSpotXCoordinate
		"BlindSpotXCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0108),
		"Blind Spot Y-Coordinate", // BlindSpotYCoordinate
		"BlindSpotYCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0110),
		"Visual Acuity Measurement Sequence", // VisualAcuityMeasurementSequence
		"VisualAcuityMeasurementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0112),
		"Refractive Parameters Used on Patient Sequence", // RefractiveParametersUsedOnPatientSequence
		"RefractiveParametersUsedOnPatientSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0113),
		"Measurement Laterality", // MeasurementLaterality
		"MeasurementLaterality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0114),
		"Ophthalmic Patient Clinical Information Left Eye Sequence", // OphthalmicPatientClinicalInformationLeftEyeSequence
		"OphthalmicPatientClinicalInformationLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0115),
		"Ophthalmic Patient Clinical Information Right Eye Sequence", // OphthalmicPatientClinicalInformationRightEyeSequence
		"OphthalmicPatientClinicalInformationRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0117),
		"Foveal Point Normative Data Flag", // FovealPointNormativeDataFlag
		"FovealPointNormativeDataFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0118),
		"Foveal Point Probability Value", // FovealPointProbabilityValue
		"FovealPointProbabilityValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0120),
		"Screening Baseline Measured", // ScreeningBaselineMeasured
		"ScreeningBaselineMeasured",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0122),
		"Screening Baseline Measured Sequence", // ScreeningBaselineMeasuredSequence
		"ScreeningBaselineMeasuredSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0124),
		"Screening Baseline Type", // ScreeningBaselineType
		"ScreeningBaselineType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0126),
		"Screening Baseline Value", // ScreeningBaselineValue
		"ScreeningBaselineValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0202),
		"Algorithm Source", // AlgorithmSource
		"AlgorithmSource",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0306),
		"Data Set Name", // DataSetName
		"DataSetName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0307),
		"Data Set Version", // DataSetVersion
		"DataSetVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0308),
		"Data Set Source", // DataSetSource
		"DataSetSource",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0309),
		"Data Set Description", // DataSetDescription
		"DataSetDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0317),
		"Visual Field Test Reliability Global Index Sequence", // VisualFieldTestReliabilityGlobalIndexSequence
		"VisualFieldTestReliabilityGlobalIndexSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0320),
		"Visual Field Global Results Index Sequence", // VisualFieldGlobalResultsIndexSequence
		"VisualFieldGlobalResultsIndexSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0325),
		"Data Observation Sequence", // DataObservationSequence
		"DataObservationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0338),
		"Index Normals Flag", // IndexNormalsFlag
		"IndexNormalsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0341),
		"Index Probability", // IndexProbability
		"IndexProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0024, 0x0344),
		"Index Probability Sequence", // IndexProbabilitySequence
		"IndexProbabilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0002),
		"Samples per Pixel", // SamplesPerPixel
		"SamplesPerPixel",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0003),
		"Samples per Pixel Used", // SamplesPerPixelUsed
		"SamplesPerPixelUsed",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0004),
		"Photometric Interpretation", // PhotometricInterpretation
		"PhotometricInterpretation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0005),
		"Image Dimensions", // ImageDimensions
		"ImageDimensions",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0006),
		"Planar Configuration", // PlanarConfiguration
		"PlanarConfiguration",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0008),
		"Number of Frames", // NumberOfFrames
		"NumberOfFrames",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0009),
		"Frame Increment Pointer", // FrameIncrementPointer
		"FrameIncrementPointer",
		vm.VM1_n,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x000A),
		"Frame Dimension Pointer", // FrameDimensionPointer
		"FrameDimensionPointer",
		vm.VM1_n,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0010),
		"Rows", // Rows
		"Rows",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0011),
		"Columns", // Columns
		"Columns",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0012),
		"Planes", // Planes
		"Planes",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0014),
		"Ultrasound Color Data Present", // UltrasoundColorDataPresent
		"UltrasoundColorDataPresent",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0030),
		"Pixel Spacing", // PixelSpacing
		"PixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0031),
		"Zoom Factor", // ZoomFactor
		"ZoomFactor",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0032),
		"Zoom Center", // ZoomCenter
		"ZoomCenter",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0034),
		"Pixel Aspect Ratio", // PixelAspectRatio
		"PixelAspectRatio",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0040),
		"Image Format", // ImageFormat
		"ImageFormat",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0050),
		"Manipulated Image", // ManipulatedImage
		"ManipulatedImage",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0051),
		"Corrected Image", // CorrectedImage
		"CorrectedImage",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x005F),
		"Compression Recognition Code", // CompressionRecognitionCode
		"CompressionRecognitionCode",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0060),
		"Compression Code", // CompressionCode
		"CompressionCode",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0061),
		"Compression Originator", // CompressionOriginator
		"CompressionOriginator",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0062),
		"Compression Label", // CompressionLabel
		"CompressionLabel",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0063),
		"Compression Description", // CompressionDescription
		"CompressionDescription",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0065),
		"Compression Sequence", // CompressionSequence
		"CompressionSequence",
		vm.VM1_n,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0066),
		"Compression Step Pointers", // CompressionStepPointers
		"CompressionStepPointers",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0068),
		"Repeat Interval", // RepeatInterval
		"RepeatInterval",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0069),
		"Bits Grouped", // BitsGrouped
		"BitsGrouped",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0070),
		"Perimeter Table", // PerimeterTable
		"PerimeterTable",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0071),
		"Perimeter Value", // PerimeterValue
		"PerimeterValue",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0080),
		"Predictor Rows", // PredictorRows
		"PredictorRows",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0081),
		"Predictor Columns", // PredictorColumns
		"PredictorColumns",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0082),
		"Predictor Constants", // PredictorConstants
		"PredictorConstants",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0090),
		"Blocked Pixels", // BlockedPixels
		"BlockedPixels",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0091),
		"Block Rows", // BlockRows
		"BlockRows",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0092),
		"Block Columns", // BlockColumns
		"BlockColumns",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0093),
		"Row Overlap", // RowOverlap
		"RowOverlap",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0094),
		"Column Overlap", // ColumnOverlap
		"ColumnOverlap",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0100),
		"Bits Allocated", // BitsAllocated
		"BitsAllocated",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0101),
		"Bits Stored", // BitsStored
		"BitsStored",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0102),
		"High Bit", // HighBit
		"HighBit",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0103),
		"Pixel Representation", // PixelRepresentation
		"PixelRepresentation",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0104),
		"Smallest Valid Pixel Value", // SmallestValidPixelValue
		"SmallestValidPixelValue",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0105),
		"Largest Valid Pixel Value", // LargestValidPixelValue
		"LargestValidPixelValue",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0106),
		"Smallest Image Pixel Value", // SmallestImagePixelValue
		"SmallestImagePixelValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0107),
		"Largest Image Pixel Value", // LargestImagePixelValue
		"LargestImagePixelValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0108),
		"Smallest Pixel Value in Series", // SmallestPixelValueInSeries
		"SmallestPixelValueInSeries",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0109),
		"Largest Pixel Value in Series", // LargestPixelValueInSeries
		"LargestPixelValueInSeries",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0110),
		"Smallest Image Pixel Value in Plane", // SmallestImagePixelValueInPlane
		"SmallestImagePixelValueInPlane",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0111),
		"Largest Image Pixel Value in Plane", // LargestImagePixelValueInPlane
		"LargestImagePixelValueInPlane",
		vm.VM1,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0120),
		"Pixel Padding Value", // PixelPaddingValue
		"PixelPaddingValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0121),
		"Pixel Padding Range Limit", // PixelPaddingRangeLimit
		"PixelPaddingRangeLimit",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0122),
		"Float Pixel Padding Value", // FloatPixelPaddingValue
		"FloatPixelPaddingValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0123),
		"Double Float Pixel Padding Value", // DoubleFloatPixelPaddingValue
		"DoubleFloatPixelPaddingValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0124),
		"Float Pixel Padding Range Limit", // FloatPixelPaddingRangeLimit
		"FloatPixelPaddingRangeLimit",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0125),
		"Double Float Pixel Padding Range Limit", // DoubleFloatPixelPaddingRangeLimit
		"DoubleFloatPixelPaddingRangeLimit",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0200),
		"Image Location", // ImageLocation
		"ImageLocation",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0300),
		"Quality Control Image", // QualityControlImage
		"QualityControlImage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0301),
		"Burned In Annotation", // BurnedInAnnotation
		"BurnedInAnnotation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0302),
		"Recognizable Visual Features", // RecognizableVisualFeatures
		"RecognizableVisualFeatures",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0303),
		"Longitudinal Temporal Information Modified", // LongitudinalTemporalInformationModified
		"LongitudinalTemporalInformationModified",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0304),
		"Referenced Color Palette Instance UID", // ReferencedColorPaletteInstanceUID
		"ReferencedColorPaletteInstanceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0400),
		"Transform Label", // TransformLabel
		"TransformLabel",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0401),
		"Transform Version Number", // TransformVersionNumber
		"TransformVersionNumber",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0402),
		"Number of Transform Steps", // NumberOfTransformSteps
		"NumberOfTransformSteps",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0403),
		"Sequence of Compressed Data", // SequenceOfCompressedData
		"SequenceOfCompressedData",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0404),
		"Details of Coefficients", // DetailsOfCoefficients
		"DetailsOfCoefficients",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,04x0)"),
		"Rows For Nth Order Coefficients", // RowsForNthOrderCoefficients
		"RowsForNthOrderCoefficients",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,04x1)"),
		"Columns For Nth Order Coefficients", // ColumnsForNthOrderCoefficients
		"ColumnsForNthOrderCoefficients",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,04x2)"),
		"Coefficient Coding", // CoefficientCoding
		"CoefficientCoding",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,04x3)"),
		"Coefficient Coding Pointers", // CoefficientCodingPointers
		"CoefficientCodingPointers",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0700),
		"DCT Label", // DCTLabel
		"DCTLabel",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0701),
		"Data Block Description", // DataBlockDescription
		"DataBlockDescription",
		vm.VM1_n,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0702),
		"Data Block", // DataBlock
		"DataBlock",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0710),
		"Normalization Factor Format", // NormalizationFactorFormat
		"NormalizationFactorFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0720),
		"Zonal Map Number Format", // ZonalMapNumberFormat
		"ZonalMapNumberFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0721),
		"Zonal Map Location", // ZonalMapLocation
		"ZonalMapLocation",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0722),
		"Zonal Map Format", // ZonalMapFormat
		"ZonalMapFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0730),
		"Adaptive Map Format", // AdaptiveMapFormat
		"AdaptiveMapFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0740),
		"Code Number Format", // CodeNumberFormat
		"CodeNumberFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x0)"),
		"Code Label", // CodeLabel
		"CodeLabel",
		vm.VM1_n,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x2)"),
		"Number of Tables", // NumberOfTables
		"NumberOfTables",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x3)"),
		"Code Table Location", // CodeTableLocation
		"CodeTableLocation",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x4)"),
		"Bits For Code Word", // BitsForCodeWord
		"BitsForCodeWord",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(0028,08x8)"),
		"Image Data Location", // ImageDataLocation
		"ImageDataLocation",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0A02),
		"Pixel Spacing Calibration Type", // PixelSpacingCalibrationType
		"PixelSpacingCalibrationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x0A04),
		"Pixel Spacing Calibration Description", // PixelSpacingCalibrationDescription
		"PixelSpacingCalibrationDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1040),
		"Pixel Intensity Relationship", // PixelIntensityRelationship
		"PixelIntensityRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1041),
		"Pixel Intensity Relationship Sign", // PixelIntensityRelationshipSign
		"PixelIntensityRelationshipSign",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1050),
		"Window Center", // WindowCenter
		"WindowCenter",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1051),
		"Window Width", // WindowWidth
		"WindowWidth",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1052),
		"Rescale Intercept", // RescaleIntercept
		"RescaleIntercept",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1053),
		"Rescale Slope", // RescaleSlope
		"RescaleSlope",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1054),
		"Rescale Type", // RescaleType
		"RescaleType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1055),
		"Window Center & Width Explanation", // WindowCenterWidthExplanation
		"WindowCenterWidthExplanation",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1056),
		"VOI LUT Function", // VOILUTFunction
		"VOILUTFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1080),
		"Gray Scale", // GrayScale
		"GrayScale",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1090),
		"Recommended Viewing Mode", // RecommendedViewingMode
		"RecommendedViewingMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1100),
		"Gray Lookup Table Descriptor", // GrayLookupTableDescriptor
		"GrayLookupTableDescriptor",
		vm.VM3,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1101),
		"Red Palette Color Lookup Table Descriptor", // RedPaletteColorLookupTableDescriptor
		"RedPaletteColorLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1102),
		"Green Palette Color Lookup Table Descriptor", // GreenPaletteColorLookupTableDescriptor
		"GreenPaletteColorLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1103),
		"Blue Palette Color Lookup Table Descriptor", // BluePaletteColorLookupTableDescriptor
		"BluePaletteColorLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1104),
		"Alpha Palette Color Lookup Table Descriptor", // AlphaPaletteColorLookupTableDescriptor
		"AlphaPaletteColorLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1111),
		"Large Red Palette Color Lookup Table Descriptor", // LargeRedPaletteColorLookupTableDescriptor
		"LargeRedPaletteColorLookupTableDescriptor",
		vm.VM4,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1112),
		"Large Green Palette Color Lookup Table Descriptor", // LargeGreenPaletteColorLookupTableDescriptor
		"LargeGreenPaletteColorLookupTableDescriptor",
		vm.VM4,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1113),
		"Large Blue Palette Color Lookup Table Descriptor", // LargeBluePaletteColorLookupTableDescriptor
		"LargeBluePaletteColorLookupTableDescriptor",
		vm.VM4,
		true,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1199),
		"Palette Color Lookup Table UID", // PaletteColorLookupTableUID
		"PaletteColorLookupTableUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1200),
		"Gray Lookup Table Data", // GrayLookupTableData
		"GrayLookupTableData",
		vm.VM1_n,
		true,
		vr.US, vr.SS, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1201),
		"Red Palette Color Lookup Table Data", // RedPaletteColorLookupTableData
		"RedPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1202),
		"Green Palette Color Lookup Table Data", // GreenPaletteColorLookupTableData
		"GreenPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1203),
		"Blue Palette Color Lookup Table Data", // BluePaletteColorLookupTableData
		"BluePaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1204),
		"Alpha Palette Color Lookup Table Data", // AlphaPaletteColorLookupTableData
		"AlphaPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1211),
		"Large Red Palette Color Lookup Table Data", // LargeRedPaletteColorLookupTableData
		"LargeRedPaletteColorLookupTableData",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1212),
		"Large Green Palette Color Lookup Table Data", // LargeGreenPaletteColorLookupTableData
		"LargeGreenPaletteColorLookupTableData",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1213),
		"Large Blue Palette Color Lookup Table Data", // LargeBluePaletteColorLookupTableData
		"LargeBluePaletteColorLookupTableData",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1214),
		"Large Palette Color Lookup Table UID", // LargePaletteColorLookupTableUID
		"LargePaletteColorLookupTableUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1221),
		"Segmented Red Palette Color Lookup Table Data", // SegmentedRedPaletteColorLookupTableData
		"SegmentedRedPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1222),
		"Segmented Green Palette Color Lookup Table Data", // SegmentedGreenPaletteColorLookupTableData
		"SegmentedGreenPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1223),
		"Segmented Blue Palette Color Lookup Table Data", // SegmentedBluePaletteColorLookupTableData
		"SegmentedBluePaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1224),
		"Segmented Alpha Palette Color Lookup Table Data", // SegmentedAlphaPaletteColorLookupTableData
		"SegmentedAlphaPaletteColorLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1230),
		"Stored Value Color Range Sequence", // StoredValueColorRangeSequence
		"StoredValueColorRangeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1231),
		"Minimum Stored Value Mapped", // MinimumStoredValueMapped
		"MinimumStoredValueMapped",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1232),
		"Maximum Stored Value Mapped", // MaximumStoredValueMapped
		"MaximumStoredValueMapped",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1300),
		"Breast Implant Present", // BreastImplantPresent
		"BreastImplantPresent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1350),
		"Partial View", // PartialView
		"PartialView",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1351),
		"Partial View Description", // PartialViewDescription
		"PartialViewDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1352),
		"Partial View Code Sequence", // PartialViewCodeSequence
		"PartialViewCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x135A),
		"Spatial Locations Preserved", // SpatialLocationsPreserved
		"SpatialLocationsPreserved",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1401),
		"Data Frame Assignment Sequence", // DataFrameAssignmentSequence
		"DataFrameAssignmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1402),
		"Data Path Assignment", // DataPathAssignment
		"DataPathAssignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1403),
		"Bits Mapped to Color Lookup Table", // BitsMappedToColorLookupTable
		"BitsMappedToColorLookupTable",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1404),
		"Blending LUT 1 Sequence", // BlendingLUT1Sequence
		"BlendingLUT1Sequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1405),
		"Blending LUT 1 Transfer Function", // BlendingLUT1TransferFunction
		"BlendingLUT1TransferFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1406),
		"Blending Weight Constant", // BlendingWeightConstant
		"BlendingWeightConstant",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1407),
		"Blending Lookup Table Descriptor", // BlendingLookupTableDescriptor
		"BlendingLookupTableDescriptor",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1408),
		"Blending Lookup Table Data", // BlendingLookupTableData
		"BlendingLookupTableData",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140B),
		"Enhanced Palette Color Lookup Table Sequence", // EnhancedPaletteColorLookupTableSequence
		"EnhancedPaletteColorLookupTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140C),
		"Blending LUT 2 Sequence", // BlendingLUT2Sequence
		"BlendingLUT2Sequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140D),
		"Blending LUT 2 Transfer Function", // BlendingLUT2TransferFunction
		"BlendingLUT2TransferFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140E),
		"Data Path ID", // DataPathID
		"DataPathID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x140F),
		"RGB LUT Transfer Function", // RGBLUTTransferFunction
		"RGBLUTTransferFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x1410),
		"Alpha LUT Transfer Function", // AlphaLUTTransferFunction
		"AlphaLUTTransferFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2000),
		"ICC Profile", // ICCProfile
		"ICCProfile",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2002),
		"Color Space", // ColorSpace
		"ColorSpace",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2110),
		"Lossy Image Compression", // LossyImageCompression
		"LossyImageCompression",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2112),
		"Lossy Image Compression Ratio", // LossyImageCompressionRatio
		"LossyImageCompressionRatio",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x2114),
		"Lossy Image Compression Method", // LossyImageCompressionMethod
		"LossyImageCompressionMethod",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3000),
		"Modality LUT Sequence", // ModalityLUTSequence
		"ModalityLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3001),
		"Variable Modality LUT Sequence", // VariableModalityLUTSequence
		"VariableModalityLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3002),
		"LUT Descriptor", // LUTDescriptor
		"LUTDescriptor",
		vm.VM3,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3003),
		"LUT Explanation", // LUTExplanation
		"LUTExplanation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3004),
		"Modality LUT Type", // ModalityLUTType
		"ModalityLUTType",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3006),
		"LUT Data", // LUTData
		"LUTData",
		vm.VM1_n,
		false,
		vr.US, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3010),
		"VOI LUT Sequence", // VOILUTSequence
		"VOILUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x3110),
		"Softcopy VOI LUT Sequence", // SoftcopyVOILUTSequence
		"SoftcopyVOILUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x4000),
		"Image Presentation Comments", // ImagePresentationComments
		"ImagePresentationComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x5000),
		"Bi-Plane Acquisition Sequence", // BiPlaneAcquisitionSequence
		"BiPlaneAcquisitionSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6010),
		"Representative Frame Number", // RepresentativeFrameNumber
		"RepresentativeFrameNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6020),
		"Frame Numbers of Interest (FOI)", // FrameNumbersOfInterest
		"FrameNumbersOfInterest",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6022),
		"Frame of Interest Description", // FrameOfInterestDescription
		"FrameOfInterestDescription",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6023),
		"Frame of Interest Type", // FrameOfInterestType
		"FrameOfInterestType",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6030),
		"Mask Pointer(s)", // MaskPointers
		"MaskPointers",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6040),
		"R Wave Pointer", // RWavePointer
		"RWavePointer",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6100),
		"Mask Subtraction Sequence", // MaskSubtractionSequence
		"MaskSubtractionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6101),
		"Mask Operation", // MaskOperation
		"MaskOperation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6102),
		"Applicable Frame Range", // ApplicableFrameRange
		"ApplicableFrameRange",
		vm.VM2_2n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6110),
		"Mask Frame Numbers", // MaskFrameNumbers
		"MaskFrameNumbers",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6112),
		"Contrast Frame Averaging", // ContrastFrameAveraging
		"ContrastFrameAveraging",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6114),
		"Mask Sub-pixel Shift", // MaskSubPixelShift
		"MaskSubPixelShift",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6120),
		"TID Offset", // TIDOffset
		"TIDOffset",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x6190),
		"Mask Operation Explanation", // MaskOperationExplanation
		"MaskOperationExplanation",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7000),
		"Equipment Administrator Sequence", // EquipmentAdministratorSequence
		"EquipmentAdministratorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7001),
		"Number of Display Subsystems", // NumberOfDisplaySubsystems
		"NumberOfDisplaySubsystems",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7002),
		"Current Configuration ID", // CurrentConfigurationID
		"CurrentConfigurationID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7003),
		"Display Subsystem ID", // DisplaySubsystemID
		"DisplaySubsystemID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7004),
		"Display Subsystem Name", // DisplaySubsystemName
		"DisplaySubsystemName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7005),
		"Display Subsystem Description", // DisplaySubsystemDescription
		"DisplaySubsystemDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7006),
		"System Status", // SystemStatus
		"SystemStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7007),
		"System Status Comment", // SystemStatusComment
		"SystemStatusComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7008),
		"Target Luminance Characteristics Sequence", // TargetLuminanceCharacteristicsSequence
		"TargetLuminanceCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7009),
		"Luminance Characteristics ID", // LuminanceCharacteristicsID
		"LuminanceCharacteristicsID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700A),
		"Display Subsystem Configuration Sequence", // DisplaySubsystemConfigurationSequence
		"DisplaySubsystemConfigurationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700B),
		"Configuration ID", // ConfigurationID
		"ConfigurationID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700C),
		"Configuration Name", // ConfigurationName
		"ConfigurationName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700D),
		"Configuration Description", // ConfigurationDescription
		"ConfigurationDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700E),
		"Referenced Target Luminance Characteristics ID", // ReferencedTargetLuminanceCharacteristicsID
		"ReferencedTargetLuminanceCharacteristicsID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x700F),
		"QA Results Sequence", // QAResultsSequence
		"QAResultsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7010),
		"Display Subsystem QA Results Sequence", // DisplaySubsystemQAResultsSequence
		"DisplaySubsystemQAResultsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7011),
		"Configuration QA Results Sequence", // ConfigurationQAResultsSequence
		"ConfigurationQAResultsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7012),
		"Measurement Equipment Sequence", // MeasurementEquipmentSequence
		"MeasurementEquipmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7013),
		"Measurement Functions", // MeasurementFunctions
		"MeasurementFunctions",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7014),
		"Measurement Equipment Type", // MeasurementEquipmentType
		"MeasurementEquipmentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7015),
		"Visual Evaluation Result Sequence", // VisualEvaluationResultSequence
		"VisualEvaluationResultSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7016),
		"Display Calibration Result Sequence", // DisplayCalibrationResultSequence
		"DisplayCalibrationResultSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7017),
		"DDL Value", // DDLValue
		"DDLValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7018),
		"CIExy White Point", // CIExyWhitePoint
		"CIExyWhitePoint",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7019),
		"Display Function Type", // DisplayFunctionType
		"DisplayFunctionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701A),
		"Gamma Value", // GammaValue
		"GammaValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701B),
		"Number of Luminance Points", // NumberOfLuminancePoints
		"NumberOfLuminancePoints",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701C),
		"Luminance Response Sequence", // LuminanceResponseSequence
		"LuminanceResponseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701D),
		"Target Minimum Luminance", // TargetMinimumLuminance
		"TargetMinimumLuminance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701E),
		"Target Maximum Luminance", // TargetMaximumLuminance
		"TargetMaximumLuminance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x701F),
		"Luminance Value", // LuminanceValue
		"LuminanceValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7020),
		"Luminance Response Description", // LuminanceResponseDescription
		"LuminanceResponseDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7021),
		"White Point Flag", // WhitePointFlag
		"WhitePointFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7022),
		"Display Device Type Code Sequence", // DisplayDeviceTypeCodeSequence
		"DisplayDeviceTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7023),
		"Display Subsystem Sequence", // DisplaySubsystemSequence
		"DisplaySubsystemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7024),
		"Luminance Result Sequence", // LuminanceResultSequence
		"LuminanceResultSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7025),
		"Ambient Light Value Source", // AmbientLightValueSource
		"AmbientLightValueSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7026),
		"Measured Characteristics", // MeasuredCharacteristics
		"MeasuredCharacteristics",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7027),
		"Luminance Uniformity Result Sequence", // LuminanceUniformityResultSequence
		"LuminanceUniformityResultSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7028),
		"Visual Evaluation Test Sequence", // VisualEvaluationTestSequence
		"VisualEvaluationTestSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7029),
		"Test Result", // TestResult
		"TestResult",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702A),
		"Test Result Comment", // TestResultComment
		"TestResultComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702B),
		"Test Image Validation", // TestImageValidation
		"TestImageValidation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702C),
		"Test Pattern Code Sequence", // TestPatternCodeSequence
		"TestPatternCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702D),
		"Measurement Pattern Code Sequence", // MeasurementPatternCodeSequence
		"MeasurementPatternCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x702E),
		"Visual Evaluation Method Code Sequence", // VisualEvaluationMethodCodeSequence
		"VisualEvaluationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x7FE0),
		"Pixel Data Provider URL", // PixelDataProviderURL
		"PixelDataProviderURL",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9001),
		"Data Point Rows", // DataPointRows
		"DataPointRows",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9002),
		"Data Point Columns", // DataPointColumns
		"DataPointColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9003),
		"Signal Domain Columns", // SignalDomainColumns
		"SignalDomainColumns",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9099),
		"Largest Monochrome Pixel Value", // LargestMonochromePixelValue
		"LargestMonochromePixelValue",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9108),
		"Data Representation", // DataRepresentation
		"DataRepresentation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9110),
		"Pixel Measures Sequence", // PixelMeasuresSequence
		"PixelMeasuresSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9132),
		"Frame VOI LUT Sequence", // FrameVOILUTSequence
		"FrameVOILUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9145),
		"Pixel Value Transformation Sequence", // PixelValueTransformationSequence
		"PixelValueTransformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9235),
		"Signal Domain Rows", // SignalDomainRows
		"SignalDomainRows",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9411),
		"Display Filter Percentage", // DisplayFilterPercentage
		"DisplayFilterPercentage",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9415),
		"Frame Pixel Shift Sequence", // FramePixelShiftSequence
		"FramePixelShiftSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9416),
		"Subtraction Item ID", // SubtractionItemID
		"SubtractionItemID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9422),
		"Pixel Intensity Relationship LUT Sequence", // PixelIntensityRelationshipLUTSequence
		"PixelIntensityRelationshipLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9443),
		"Frame Pixel Data Properties Sequence", // FramePixelDataPropertiesSequence
		"FramePixelDataPropertiesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9444),
		"Geometrical Properties", // GeometricalProperties
		"GeometricalProperties",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9445),
		"Geometric Maximum Distortion", // GeometricMaximumDistortion
		"GeometricMaximumDistortion",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9446),
		"Image Processing Applied", // ImageProcessingApplied
		"ImageProcessingApplied",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9454),
		"Mask Selection Mode", // MaskSelectionMode
		"MaskSelectionMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9474),
		"LUT Function", // LUTFunction
		"LUTFunction",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9478),
		"Mask Visibility Percentage", // MaskVisibilityPercentage
		"MaskVisibilityPercentage",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9501),
		"Pixel Shift Sequence", // PixelShiftSequence
		"PixelShiftSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9502),
		"Region Pixel Shift Sequence", // RegionPixelShiftSequence
		"RegionPixelShiftSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9503),
		"Vertices of the Region", // VerticesOfTheRegion
		"VerticesOfTheRegion",
		vm.VM2_2n,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9505),
		"Multi-frame Presentation Sequence", // MultiFramePresentationSequence
		"MultiFramePresentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9506),
		"Pixel Shift Frame Range", // PixelShiftFrameRange
		"PixelShiftFrameRange",
		vm.VM2_2n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9507),
		"LUT Frame Range", // LUTFrameRange
		"LUTFrameRange",
		vm.VM2_2n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9520),
		"Image to Equipment Mapping Matrix", // ImageToEquipmentMappingMatrix
		"ImageToEquipmentMappingMatrix",
		vm.VM16,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0028, 0x9537),
		"Equipment Coordinate System Identification", // EquipmentCoordinateSystemIdentification
		"EquipmentCoordinateSystemIdentification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x000A),
		"Study Status ID", // StudyStatusID
		"StudyStatusID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x000C),
		"Study Priority ID", // StudyPriorityID
		"StudyPriorityID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0012),
		"Study ID Issuer", // StudyIDIssuer
		"StudyIDIssuer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0032),
		"Study Verified Date", // StudyVerifiedDate
		"StudyVerifiedDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0033),
		"Study Verified Time", // StudyVerifiedTime
		"StudyVerifiedTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0034),
		"Study Read Date", // StudyReadDate
		"StudyReadDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x0035),
		"Study Read Time", // StudyReadTime
		"StudyReadTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1000),
		"Scheduled Study Start Date", // ScheduledStudyStartDate
		"ScheduledStudyStartDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1001),
		"Scheduled Study Start Time", // ScheduledStudyStartTime
		"ScheduledStudyStartTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1010),
		"Scheduled Study Stop Date", // ScheduledStudyStopDate
		"ScheduledStudyStopDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1011),
		"Scheduled Study Stop Time", // ScheduledStudyStopTime
		"ScheduledStudyStopTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1020),
		"Scheduled Study Location", // ScheduledStudyLocation
		"ScheduledStudyLocation",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1021),
		"Scheduled Study Location AE Title", // ScheduledStudyLocationAETitle
		"ScheduledStudyLocationAETitle",
		vm.VM1_n,
		true,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1030),
		"Reason for Study", // ReasonForStudy
		"ReasonForStudy",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1031),
		"Requesting Physician Identification Sequence", // RequestingPhysicianIdentificationSequence
		"RequestingPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1032),
		"Requesting Physician", // RequestingPhysician
		"RequestingPhysician",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1033),
		"Requesting Service", // RequestingService
		"RequestingService",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1034),
		"Requesting Service Code Sequence", // RequestingServiceCodeSequence
		"RequestingServiceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1040),
		"Study Arrival Date", // StudyArrivalDate
		"StudyArrivalDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1041),
		"Study Arrival Time", // StudyArrivalTime
		"StudyArrivalTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1050),
		"Study Completion Date", // StudyCompletionDate
		"StudyCompletionDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1051),
		"Study Completion Time", // StudyCompletionTime
		"StudyCompletionTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1055),
		"Study Component Status ID", // StudyComponentStatusID
		"StudyComponentStatusID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1060),
		"Requested Procedure Description", // RequestedProcedureDescription
		"RequestedProcedureDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1064),
		"Requested Procedure Code Sequence", // RequestedProcedureCodeSequence
		"RequestedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1065),
		"Requested Laterality Code Sequence", // RequestedLateralityCodeSequence
		"RequestedLateralityCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1066),
		"Reason for Visit", // ReasonForVisit
		"ReasonForVisit",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1067),
		"Reason for Visit Code Sequence", // ReasonForVisitCodeSequence
		"ReasonForVisitCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x1070),
		"Requested Contrast Agent", // RequestedContrastAgent
		"RequestedContrastAgent",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0032, 0x4000),
		"Study Comments", // StudyComments
		"StudyComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0001),
		"Flow Identifier Sequence", // FlowIdentifierSequence
		"FlowIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0002),
		"Flow Identifier", // FlowIdentifier
		"FlowIdentifier",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0003),
		"Flow Transfer Syntax UID", // FlowTransferSyntaxUID
		"FlowTransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0004),
		"Flow RTP Sampling Rate", // FlowRTPSamplingRate
		"FlowRTPSamplingRate",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0005),
		"Source Identifier", // SourceIdentifier
		"SourceIdentifier",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0007),
		"Frame Origin Timestamp", // FrameOriginTimestamp
		"FrameOriginTimestamp",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0008),
		"Includes Imaging Subject", // IncludesImagingSubject
		"IncludesImagingSubject",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x0009),
		"Frame Usefulness Group Sequence", // FrameUsefulnessGroupSequence
		"FrameUsefulnessGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x000A),
		"Real-Time Bulk Data Flow Sequence", // RealTimeBulkDataFlowSequence
		"RealTimeBulkDataFlowSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x000B),
		"Camera Position Group Sequence", // CameraPositionGroupSequence
		"CameraPositionGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x000C),
		"Includes Information", // IncludesInformation
		"IncludesInformation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0034, 0x000D),
		"Time of Frame Group Sequence", // TimeOfFrameGroupSequence
		"TimeOfFrameGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0004),
		"Referenced Patient Alias Sequence", // ReferencedPatientAliasSequence
		"ReferencedPatientAliasSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0008),
		"Visit Status ID", // VisitStatusID
		"VisitStatusID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0010),
		"Admission ID", // AdmissionID
		"AdmissionID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0011),
		"Issuer of Admission ID", // IssuerOfAdmissionID
		"IssuerOfAdmissionID",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0014),
		"Issuer of Admission ID Sequence", // IssuerOfAdmissionIDSequence
		"IssuerOfAdmissionIDSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0016),
		"Route of Admissions", // RouteOfAdmissions
		"RouteOfAdmissions",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001A),
		"Scheduled Admission Date", // ScheduledAdmissionDate
		"ScheduledAdmissionDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001B),
		"Scheduled Admission Time", // ScheduledAdmissionTime
		"ScheduledAdmissionTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001C),
		"Scheduled Discharge Date", // ScheduledDischargeDate
		"ScheduledDischargeDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001D),
		"Scheduled Discharge Time", // ScheduledDischargeTime
		"ScheduledDischargeTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x001E),
		"Scheduled Patient Institution Residence", // ScheduledPatientInstitutionResidence
		"ScheduledPatientInstitutionResidence",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0020),
		"Admitting Date", // AdmittingDate
		"AdmittingDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0021),
		"Admitting Time", // AdmittingTime
		"AdmittingTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0030),
		"Discharge Date", // DischargeDate
		"DischargeDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0032),
		"Discharge Time", // DischargeTime
		"DischargeTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0040),
		"Discharge Diagnosis Description", // DischargeDiagnosisDescription
		"DischargeDiagnosisDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0044),
		"Discharge Diagnosis Code Sequence", // DischargeDiagnosisCodeSequence
		"DischargeDiagnosisCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0050),
		"Special Needs", // SpecialNeeds
		"SpecialNeeds",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0060),
		"Service Episode ID", // ServiceEpisodeID
		"ServiceEpisodeID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0061),
		"Issuer of Service Episode ID", // IssuerOfServiceEpisodeID
		"IssuerOfServiceEpisodeID",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0062),
		"Service Episode Description", // ServiceEpisodeDescription
		"ServiceEpisodeDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0064),
		"Issuer of Service Episode ID Sequence", // IssuerOfServiceEpisodeIDSequence
		"IssuerOfServiceEpisodeIDSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0100),
		"Pertinent Documents Sequence", // PertinentDocumentsSequence
		"PertinentDocumentsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0101),
		"Pertinent Resources Sequence", // PertinentResourcesSequence
		"PertinentResourcesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0102),
		"Resource Description", // ResourceDescription
		"ResourceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0300),
		"Current Patient Location", // CurrentPatientLocation
		"CurrentPatientLocation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0400),
		"Patient's Institution Residence", // PatientInstitutionResidence
		"PatientInstitutionResidence",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0500),
		"Patient State", // PatientState
		"PatientState",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x0502),
		"Patient Clinical Trial Participation Sequence", // PatientClinicalTrialParticipationSequence
		"PatientClinicalTrialParticipationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0038, 0x4000),
		"Visit Comments", // VisitComments
		"VisitComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0004),
		"Waveform Originality", // WaveformOriginality
		"WaveformOriginality",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0005),
		"Number of Waveform Channels", // NumberOfWaveformChannels
		"NumberOfWaveformChannels",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0010),
		"Number of Waveform Samples", // NumberOfWaveformSamples
		"NumberOfWaveformSamples",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x001A),
		"Sampling Frequency", // SamplingFrequency
		"SamplingFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0020),
		"Multiplex Group Label", // MultiplexGroupLabel
		"MultiplexGroupLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0200),
		"Channel Definition Sequence", // ChannelDefinitionSequence
		"ChannelDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0202),
		"Waveform Channel Number", // WaveformChannelNumber
		"WaveformChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0203),
		"Channel Label", // ChannelLabel
		"ChannelLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0205),
		"Channel Status", // ChannelStatus
		"ChannelStatus",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0208),
		"Channel Source Sequence", // ChannelSourceSequence
		"ChannelSourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0209),
		"Channel Source Modifiers Sequence", // ChannelSourceModifiersSequence
		"ChannelSourceModifiersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x020A),
		"Source Waveform Sequence", // SourceWaveformSequence
		"SourceWaveformSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x020C),
		"Channel Derivation Description", // ChannelDerivationDescription
		"ChannelDerivationDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0210),
		"Channel Sensitivity", // ChannelSensitivity
		"ChannelSensitivity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0211),
		"Channel Sensitivity Units Sequence", // ChannelSensitivityUnitsSequence
		"ChannelSensitivityUnitsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0212),
		"Channel Sensitivity Correction Factor", // ChannelSensitivityCorrectionFactor
		"ChannelSensitivityCorrectionFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0213),
		"Channel Baseline", // ChannelBaseline
		"ChannelBaseline",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0214),
		"Channel Time Skew", // ChannelTimeSkew
		"ChannelTimeSkew",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0215),
		"Channel Sample Skew", // ChannelSampleSkew
		"ChannelSampleSkew",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0218),
		"Channel Offset", // ChannelOffset
		"ChannelOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x021A),
		"Waveform Bits Stored", // WaveformBitsStored
		"WaveformBitsStored",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0220),
		"Filter Low Frequency", // FilterLowFrequency
		"FilterLowFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0221),
		"Filter High Frequency", // FilterHighFrequency
		"FilterHighFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0222),
		"Notch Filter Frequency", // NotchFilterFrequency
		"NotchFilterFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0223),
		"Notch Filter Bandwidth", // NotchFilterBandwidth
		"NotchFilterBandwidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0230),
		"Waveform Data Display Scale", // WaveformDataDisplayScale
		"WaveformDataDisplayScale",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0231),
		"Waveform Display Background CIELab Value", // WaveformDisplayBackgroundCIELabValue
		"WaveformDisplayBackgroundCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0240),
		"Waveform Presentation Group Sequence", // WaveformPresentationGroupSequence
		"WaveformPresentationGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0241),
		"Presentation Group Number", // PresentationGroupNumber
		"PresentationGroupNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0242),
		"Channel Display Sequence", // ChannelDisplaySequence
		"ChannelDisplaySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0244),
		"Channel Recommended Display CIELab Value", // ChannelRecommendedDisplayCIELabValue
		"ChannelRecommendedDisplayCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0245),
		"Channel Position", // ChannelPosition
		"ChannelPosition",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0246),
		"Display Shading Flag", // DisplayShadingFlag
		"DisplayShadingFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0247),
		"Fractional Channel Display Scale", // FractionalChannelDisplayScale
		"FractionalChannelDisplayScale",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0248),
		"Absolute Channel Display Scale", // AbsoluteChannelDisplayScale
		"AbsoluteChannelDisplayScale",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0300),
		"Multiplexed Audio Channels Description Code Sequence", // MultiplexedAudioChannelsDescriptionCodeSequence
		"MultiplexedAudioChannelsDescriptionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0301),
		"Channel Identification Code", // ChannelIdentificationCode
		"ChannelIdentificationCode",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0302),
		"Channel Mode", // ChannelMode
		"ChannelMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0310),
		"Multiplex Group UID", // MultiplexGroupUID
		"MultiplexGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0311),
		"Powerline Frequency", // PowerlineFrequency
		"PowerlineFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0312),
		"Channel Impedance Sequence", // ChannelImpedanceSequence
		"ChannelImpedanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0313),
		"Impedance Value", // ImpedanceValue
		"ImpedanceValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0314),
		"Impedance Measurement DateTime", // ImpedanceMeasurementDateTime
		"ImpedanceMeasurementDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0315),
		"Impedance Measurement Frequency", // ImpedanceMeasurementFrequency
		"ImpedanceMeasurementFrequency",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0316),
		"Impedance Measurement Current Type", // ImpedanceMeasurementCurrentType
		"ImpedanceMeasurementCurrentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0317),
		"Waveform Amplifier Type", // WaveformAmplifierType
		"WaveformAmplifierType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0318),
		"Filter Low Frequency Characteristics Sequence", // FilterLowFrequencyCharacteristicsSequence
		"FilterLowFrequencyCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0319),
		"Filter High Frequency Characteristics Sequence", // FilterHighFrequencyCharacteristicsSequence
		"FilterHighFrequencyCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0320),
		"Summarized Filter Lookup Table Sequence", // SummarizedFilterLookupTableSequence
		"SummarizedFilterLookupTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0321),
		"Notch Filter Characteristics Sequence", // NotchFilterCharacteristicsSequence
		"NotchFilterCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0322),
		"Waveform Filter Type", // WaveformFilterType
		"WaveformFilterType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0323),
		"Analog Filter Characteristics Sequence", // AnalogFilterCharacteristicsSequence
		"AnalogFilterCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0324),
		"Analog Filter Roll Off ", // AnalogFilterRollOff
		"AnalogFilterRollOff",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0325),
		"Analog Filter Type Code Sequence", // AnalogFilterTypeCodeSequence
		"AnalogFilterTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0326),
		"Digital Filter Characteristics Sequence", // DigitalFilterCharacteristicsSequence
		"DigitalFilterCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0327),
		"Digital Filter Order", // DigitalFilterOrder
		"DigitalFilterOrder",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0328),
		"Digital Filter Type Code Sequence", // DigitalFilterTypeCodeSequence
		"DigitalFilterTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x0329),
		"Waveform Filter Description", // WaveformFilterDescription
		"WaveformFilterDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032A),
		"Filter Lookup Table Sequence", // FilterLookupTableSequence
		"FilterLookupTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032B),
		"Filter Lookup Table Description", // FilterLookupTableDescription
		"FilterLookupTableDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032C),
		"Frequency Encoding Code Sequence", // FrequencyEncodingCodeSequence
		"FrequencyEncodingCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032D),
		"Magnitude Encoding Code Sequence", // MagnitudeEncodingCodeSequence
		"MagnitudeEncodingCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x003A, 0x032E),
		"Filter Lookup Table Data", // FilterLookupTableData
		"FilterLookupTableData",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0001),
		"Scheduled Station AE Title", // ScheduledStationAETitle
		"ScheduledStationAETitle",
		vm.VM1_n,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0002),
		"Scheduled Procedure Step Start Date", // ScheduledProcedureStepStartDate
		"ScheduledProcedureStepStartDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0003),
		"Scheduled Procedure Step Start Time", // ScheduledProcedureStepStartTime
		"ScheduledProcedureStepStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0004),
		"Scheduled Procedure Step End Date", // ScheduledProcedureStepEndDate
		"ScheduledProcedureStepEndDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0005),
		"Scheduled Procedure Step End Time", // ScheduledProcedureStepEndTime
		"ScheduledProcedureStepEndTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0006),
		"Scheduled Performing Physician's Name", // ScheduledPerformingPhysicianName
		"ScheduledPerformingPhysicianName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0007),
		"Scheduled Procedure Step Description", // ScheduledProcedureStepDescription
		"ScheduledProcedureStepDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0008),
		"Scheduled Protocol Code Sequence", // ScheduledProtocolCodeSequence
		"ScheduledProtocolCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0009),
		"Scheduled Procedure Step ID", // ScheduledProcedureStepID
		"ScheduledProcedureStepID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x000A),
		"Stage Code Sequence", // StageCodeSequence
		"StageCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x000B),
		"Scheduled Performing Physician Identification Sequence", // ScheduledPerformingPhysicianIdentificationSequence
		"ScheduledPerformingPhysicianIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0010),
		"Scheduled Station Name", // ScheduledStationName
		"ScheduledStationName",
		vm.VM1_n,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0011),
		"Scheduled Procedure Step Location", // ScheduledProcedureStepLocation
		"ScheduledProcedureStepLocation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0012),
		"Pre-Medication", // PreMedication
		"PreMedication",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0020),
		"Scheduled Procedure Step Status", // ScheduledProcedureStepStatus
		"ScheduledProcedureStepStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0026),
		"Order Placer Identifier Sequence", // OrderPlacerIdentifierSequence
		"OrderPlacerIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0027),
		"Order Filler Identifier Sequence", // OrderFillerIdentifierSequence
		"OrderFillerIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0031),
		"Local Namespace Entity ID", // LocalNamespaceEntityID
		"LocalNamespaceEntityID",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0032),
		"Universal Entity ID", // UniversalEntityID
		"UniversalEntityID",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0033),
		"Universal Entity ID Type", // UniversalEntityIDType
		"UniversalEntityIDType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0035),
		"Identifier Type Code", // IdentifierTypeCode
		"IdentifierTypeCode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0036),
		"Assigning Facility Sequence", // AssigningFacilitySequence
		"AssigningFacilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0039),
		"Assigning Jurisdiction Code Sequence", // AssigningJurisdictionCodeSequence
		"AssigningJurisdictionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x003A),
		"Assigning Agency or Department Code Sequence", // AssigningAgencyOrDepartmentCodeSequence
		"AssigningAgencyOrDepartmentCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0100),
		"Scheduled Procedure Step Sequence", // ScheduledProcedureStepSequence
		"ScheduledProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0220),
		"Referenced Non-Image Composite SOP Instance Sequence", // ReferencedNonImageCompositeSOPInstanceSequence
		"ReferencedNonImageCompositeSOPInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0241),
		"Performed Station AE Title", // PerformedStationAETitle
		"PerformedStationAETitle",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0242),
		"Performed Station Name", // PerformedStationName
		"PerformedStationName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0243),
		"Performed Location", // PerformedLocation
		"PerformedLocation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0244),
		"Performed Procedure Step Start Date", // PerformedProcedureStepStartDate
		"PerformedProcedureStepStartDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0245),
		"Performed Procedure Step Start Time", // PerformedProcedureStepStartTime
		"PerformedProcedureStepStartTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0250),
		"Performed Procedure Step End Date", // PerformedProcedureStepEndDate
		"PerformedProcedureStepEndDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0251),
		"Performed Procedure Step End Time", // PerformedProcedureStepEndTime
		"PerformedProcedureStepEndTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0252),
		"Performed Procedure Step Status", // PerformedProcedureStepStatus
		"PerformedProcedureStepStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0253),
		"Performed Procedure Step ID", // PerformedProcedureStepID
		"PerformedProcedureStepID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0254),
		"Performed Procedure Step Description", // PerformedProcedureStepDescription
		"PerformedProcedureStepDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0255),
		"Performed Procedure Type Description", // PerformedProcedureTypeDescription
		"PerformedProcedureTypeDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0260),
		"Performed Protocol Code Sequence", // PerformedProtocolCodeSequence
		"PerformedProtocolCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0261),
		"Performed Protocol Type", // PerformedProtocolType
		"PerformedProtocolType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0270),
		"Scheduled Step Attributes Sequence", // ScheduledStepAttributesSequence
		"ScheduledStepAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0275),
		"Request Attributes Sequence", // RequestAttributesSequence
		"RequestAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0280),
		"Comments on the Performed Procedure Step", // CommentsOnThePerformedProcedureStep
		"CommentsOnThePerformedProcedureStep",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0281),
		"Performed Procedure Step Discontinuation Reason Code Sequence", // PerformedProcedureStepDiscontinuationReasonCodeSequence
		"PerformedProcedureStepDiscontinuationReasonCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0293),
		"Quantity Sequence", // QuantitySequence
		"QuantitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0294),
		"Quantity", // Quantity
		"Quantity",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0295),
		"Measuring Units Sequence", // MeasuringUnitsSequence
		"MeasuringUnitsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0296),
		"Billing Item Sequence", // BillingItemSequence
		"BillingItemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0300),
		"Total Time of Fluoroscopy", // TotalTimeOfFluoroscopy
		"TotalTimeOfFluoroscopy",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0301),
		"Total Number of Exposures", // TotalNumberOfExposures
		"TotalNumberOfExposures",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0302),
		"Entrance Dose", // EntranceDose
		"EntranceDose",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0303),
		"Exposed Area", // ExposedArea
		"ExposedArea",
		vm.VM1_2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0306),
		"Distance Source to Entrance", // DistanceSourceToEntrance
		"DistanceSourceToEntrance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0307),
		"Distance Source to Support", // DistanceSourceToSupport
		"DistanceSourceToSupport",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x030E),
		"Exposure Dose Sequence", // ExposureDoseSequence
		"ExposureDoseSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0310),
		"Comments on Radiation Dose", // CommentsOnRadiationDose
		"CommentsOnRadiationDose",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0312),
		"X-Ray Output", // XRayOutput
		"XRayOutput",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0314),
		"Half Value Layer", // HalfValueLayer
		"HalfValueLayer",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0316),
		"Organ Dose", // OrganDose
		"OrganDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0318),
		"Organ Exposed", // OrganExposed
		"OrganExposed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0320),
		"Billing Procedure Step Sequence", // BillingProcedureStepSequence
		"BillingProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0321),
		"Film Consumption Sequence", // FilmConsumptionSequence
		"FilmConsumptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0324),
		"Billing Supplies and Devices Sequence", // BillingSuppliesAndDevicesSequence
		"BillingSuppliesAndDevicesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0330),
		"Referenced Procedure Step Sequence", // ReferencedProcedureStepSequence
		"ReferencedProcedureStepSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0340),
		"Performed Series Sequence", // PerformedSeriesSequence
		"PerformedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0400),
		"Comments on the Scheduled Procedure Step", // CommentsOnTheScheduledProcedureStep
		"CommentsOnTheScheduledProcedureStep",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0440),
		"Protocol Context Sequence", // ProtocolContextSequence
		"ProtocolContextSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0441),
		"Content Item Modifier Sequence", // ContentItemModifierSequence
		"ContentItemModifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0500),
		"Scheduled Specimen Sequence", // ScheduledSpecimenSequence
		"ScheduledSpecimenSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x050A),
		"Specimen Accession Number", // SpecimenAccessionNumber
		"SpecimenAccessionNumber",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0512),
		"Container Identifier", // ContainerIdentifier
		"ContainerIdentifier",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0513),
		"Issuer of the Container Identifier Sequence", // IssuerOfTheContainerIdentifierSequence
		"IssuerOfTheContainerIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0515),
		"Alternate Container Identifier Sequence", // AlternateContainerIdentifierSequence
		"AlternateContainerIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0518),
		"Container Type Code Sequence", // ContainerTypeCodeSequence
		"ContainerTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x051A),
		"Container Description", // ContainerDescription
		"ContainerDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0520),
		"Container Component Sequence", // ContainerComponentSequence
		"ContainerComponentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0550),
		"Specimen Sequence", // SpecimenSequence
		"SpecimenSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0551),
		"Specimen Identifier", // SpecimenIdentifier
		"SpecimenIdentifier",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0552),
		"Specimen Description Sequence (Trial)", // SpecimenDescriptionSequenceTrial
		"SpecimenDescriptionSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0553),
		"Specimen Description (Trial)", // SpecimenDescriptionTrial
		"SpecimenDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0554),
		"Specimen UID", // SpecimenUID
		"SpecimenUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0555),
		"Acquisition Context Sequence", // AcquisitionContextSequence
		"AcquisitionContextSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0556),
		"Acquisition Context Description", // AcquisitionContextDescription
		"AcquisitionContextDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0560),
		"Specimen Description Sequence", // SpecimenDescriptionSequence
		"SpecimenDescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0562),
		"Issuer of the Specimen Identifier Sequence", // IssuerOfTheSpecimenIdentifierSequence
		"IssuerOfTheSpecimenIdentifierSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x059A),
		"Specimen Type Code Sequence", // SpecimenTypeCodeSequence
		"SpecimenTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0600),
		"Specimen Short Description", // SpecimenShortDescription
		"SpecimenShortDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0602),
		"Specimen Detailed Description", // SpecimenDetailedDescription
		"SpecimenDetailedDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0610),
		"Specimen Preparation Sequence", // SpecimenPreparationSequence
		"SpecimenPreparationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0612),
		"Specimen Preparation Step Content Item Sequence", // SpecimenPreparationStepContentItemSequence
		"SpecimenPreparationStepContentItemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0620),
		"Specimen Localization Content Item Sequence", // SpecimenLocalizationContentItemSequence
		"SpecimenLocalizationContentItemSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x06FA),
		"Slide Identifier", // SlideIdentifier
		"SlideIdentifier",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x0710),
		"Whole Slide Microscopy Image Frame Type Sequence", // WholeSlideMicroscopyImageFrameTypeSequence
		"WholeSlideMicroscopyImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x071A),
		"Image Center Point Coordinates Sequence", // ImageCenterPointCoordinatesSequence
		"ImageCenterPointCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x072A),
		"X Offset in Slide Coordinate System", // XOffsetInSlideCoordinateSystem
		"XOffsetInSlideCoordinateSystem",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x073A),
		"Y Offset in Slide Coordinate System", // YOffsetInSlideCoordinateSystem
		"YOffsetInSlideCoordinateSystem",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x074A),
		"Z Offset in Slide Coordinate System", // ZOffsetInSlideCoordinateSystem
		"ZOffsetInSlideCoordinateSystem",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x08D8),
		"Pixel Spacing Sequence", // PixelSpacingSequence
		"PixelSpacingSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x08DA),
		"Coordinate System Axis Code Sequence", // CoordinateSystemAxisCodeSequence
		"CoordinateSystemAxisCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x08EA),
		"Measurement Units Code Sequence", // MeasurementUnitsCodeSequence
		"MeasurementUnitsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x09F8),
		"Vital Stain Code Sequence (Trial)", // VitalStainCodeSequenceTrial
		"VitalStainCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1001),
		"Requested Procedure ID", // RequestedProcedureID
		"RequestedProcedureID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1002),
		"Reason for the Requested Procedure", // ReasonForTheRequestedProcedure
		"ReasonForTheRequestedProcedure",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1003),
		"Requested Procedure Priority", // RequestedProcedurePriority
		"RequestedProcedurePriority",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1004),
		"Patient Transport Arrangements", // PatientTransportArrangements
		"PatientTransportArrangements",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1005),
		"Requested Procedure Location", // RequestedProcedureLocation
		"RequestedProcedureLocation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1006),
		"Placer Order Number / Procedure", // PlacerOrderNumberProcedure
		"PlacerOrderNumberProcedure",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1007),
		"Filler Order Number / Procedure", // FillerOrderNumberProcedure
		"FillerOrderNumberProcedure",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1008),
		"Confidentiality Code", // ConfidentialityCode
		"ConfidentialityCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1009),
		"Reporting Priority", // ReportingPriority
		"ReportingPriority",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x100A),
		"Reason for Requested Procedure Code Sequence", // ReasonForRequestedProcedureCodeSequence
		"ReasonForRequestedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1010),
		"Names of Intended Recipients of Results", // NamesOfIntendedRecipientsOfResults
		"NamesOfIntendedRecipientsOfResults",
		vm.VM1_n,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1011),
		"Intended Recipients of Results Identification Sequence", // IntendedRecipientsOfResultsIdentificationSequence
		"IntendedRecipientsOfResultsIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1012),
		"Reason For Performed Procedure Code Sequence", // ReasonForPerformedProcedureCodeSequence
		"ReasonForPerformedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1060),
		"Requested Procedure Description (Trial)", // RequestedProcedureDescriptionTrial
		"RequestedProcedureDescriptionTrial",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1101),
		"Person Identification Code Sequence", // PersonIdentificationCodeSequence
		"PersonIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1102),
		"Person's Address", // PersonAddress
		"PersonAddress",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1103),
		"Person's Telephone Numbers", // PersonTelephoneNumbers
		"PersonTelephoneNumbers",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1104),
		"Person's Telecom Information", // PersonTelecomInformation
		"PersonTelecomInformation",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x1400),
		"Requested Procedure Comments", // RequestedProcedureComments
		"RequestedProcedureComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2001),
		"Reason for the Imaging Service Request", // ReasonForTheImagingServiceRequest
		"ReasonForTheImagingServiceRequest",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2004),
		"Issue Date of Imaging Service Request", // IssueDateOfImagingServiceRequest
		"IssueDateOfImagingServiceRequest",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2005),
		"Issue Time of Imaging Service Request", // IssueTimeOfImagingServiceRequest
		"IssueTimeOfImagingServiceRequest",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2006),
		"Placer Order Number / Imaging Service Request (Retired)", // PlacerOrderNumberImagingServiceRequestRetired
		"PlacerOrderNumberImagingServiceRequestRetired",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2007),
		"Filler Order Number / Imaging Service Request (Retired)", // FillerOrderNumberImagingServiceRequestRetired
		"FillerOrderNumberImagingServiceRequestRetired",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2008),
		"Order Entered By", // OrderEnteredBy
		"OrderEnteredBy",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2009),
		"Order Enterer's Location", // OrderEntererLocation
		"OrderEntererLocation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2010),
		"Order Callback Phone Number", // OrderCallbackPhoneNumber
		"OrderCallbackPhoneNumber",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2011),
		"Order Callback Telecom Information", // OrderCallbackTelecomInformation
		"OrderCallbackTelecomInformation",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2016),
		"Placer Order Number / Imaging Service Request", // PlacerOrderNumberImagingServiceRequest
		"PlacerOrderNumberImagingServiceRequest",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2017),
		"Filler Order Number / Imaging Service Request", // FillerOrderNumberImagingServiceRequest
		"FillerOrderNumberImagingServiceRequest",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x2400),
		"Imaging Service Request Comments", // ImagingServiceRequestComments
		"ImagingServiceRequestComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x3001),
		"Confidentiality Constraint on Patient Data Description", // ConfidentialityConstraintOnPatientDataDescription
		"ConfidentialityConstraintOnPatientDataDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4001),
		"General Purpose Scheduled Procedure Step Status", // GeneralPurposeScheduledProcedureStepStatus
		"GeneralPurposeScheduledProcedureStepStatus",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4002),
		"General Purpose Performed Procedure Step Status", // GeneralPurposePerformedProcedureStepStatus
		"GeneralPurposePerformedProcedureStepStatus",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4003),
		"General Purpose Scheduled Procedure Step Priority", // GeneralPurposeScheduledProcedureStepPriority
		"GeneralPurposeScheduledProcedureStepPriority",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4004),
		"Scheduled Processing Applications Code Sequence", // ScheduledProcessingApplicationsCodeSequence
		"ScheduledProcessingApplicationsCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4005),
		"Scheduled Procedure Step Start DateTime", // ScheduledProcedureStepStartDateTime
		"ScheduledProcedureStepStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4006),
		"Multiple Copies Flag", // MultipleCopiesFlag
		"MultipleCopiesFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4007),
		"Performed Processing Applications Code Sequence", // PerformedProcessingApplicationsCodeSequence
		"PerformedProcessingApplicationsCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4008),
		"Scheduled Procedure Step Expiration DateTime", // ScheduledProcedureStepExpirationDateTime
		"ScheduledProcedureStepExpirationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4009),
		"Human Performer Code Sequence", // HumanPerformerCodeSequence
		"HumanPerformerCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4010),
		"Scheduled Procedure Step Modification DateTime", // ScheduledProcedureStepModificationDateTime
		"ScheduledProcedureStepModificationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4011),
		"Expected Completion DateTime", // ExpectedCompletionDateTime
		"ExpectedCompletionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4015),
		"Resulting General Purpose Performed Procedure Steps Sequence", // ResultingGeneralPurposePerformedProcedureStepsSequence
		"ResultingGeneralPurposePerformedProcedureStepsSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4016),
		"Referenced General Purpose Scheduled Procedure Step Sequence", // ReferencedGeneralPurposeScheduledProcedureStepSequence
		"ReferencedGeneralPurposeScheduledProcedureStepSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4018),
		"Scheduled Workitem Code Sequence", // ScheduledWorkitemCodeSequence
		"ScheduledWorkitemCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4019),
		"Performed Workitem Code Sequence", // PerformedWorkitemCodeSequence
		"PerformedWorkitemCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4020),
		"Input Availability Flag", // InputAvailabilityFlag
		"InputAvailabilityFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4021),
		"Input Information Sequence", // InputInformationSequence
		"InputInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4022),
		"Relevant Information Sequence", // RelevantInformationSequence
		"RelevantInformationSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4023),
		"Referenced General Purpose Scheduled Procedure Step Transaction UID", // ReferencedGeneralPurposeScheduledProcedureStepTransactionUID
		"ReferencedGeneralPurposeScheduledProcedureStepTransactionUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4025),
		"Scheduled Station Name Code Sequence", // ScheduledStationNameCodeSequence
		"ScheduledStationNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4026),
		"Scheduled Station Class Code Sequence", // ScheduledStationClassCodeSequence
		"ScheduledStationClassCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4027),
		"Scheduled Station Geographic Location Code Sequence", // ScheduledStationGeographicLocationCodeSequence
		"ScheduledStationGeographicLocationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4028),
		"Performed Station Name Code Sequence", // PerformedStationNameCodeSequence
		"PerformedStationNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4029),
		"Performed Station Class Code Sequence", // PerformedStationClassCodeSequence
		"PerformedStationClassCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4030),
		"Performed Station Geographic Location Code Sequence", // PerformedStationGeographicLocationCodeSequence
		"PerformedStationGeographicLocationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4031),
		"Requested Subsequent Workitem Code Sequence", // RequestedSubsequentWorkitemCodeSequence
		"RequestedSubsequentWorkitemCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4032),
		"Non-DICOM Output Code Sequence", // NonDICOMOutputCodeSequence
		"NonDICOMOutputCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4033),
		"Output Information Sequence", // OutputInformationSequence
		"OutputInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4034),
		"Scheduled Human Performers Sequence", // ScheduledHumanPerformersSequence
		"ScheduledHumanPerformersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4035),
		"Actual Human Performers Sequence", // ActualHumanPerformersSequence
		"ActualHumanPerformersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4036),
		"Human Performer's Organization", // HumanPerformerOrganization
		"HumanPerformerOrganization",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4037),
		"Human Performer's Name", // HumanPerformerName
		"HumanPerformerName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4040),
		"Raw Data Handling", // RawDataHandling
		"RawDataHandling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4041),
		"Input Readiness State", // InputReadinessState
		"InputReadinessState",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4050),
		"Performed Procedure Step Start DateTime", // PerformedProcedureStepStartDateTime
		"PerformedProcedureStepStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4051),
		"Performed Procedure Step End DateTime", // PerformedProcedureStepEndDateTime
		"PerformedProcedureStepEndDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4052),
		"Procedure Step Cancellation DateTime", // ProcedureStepCancellationDateTime
		"ProcedureStepCancellationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4070),
		"Output Destination Sequence", // OutputDestinationSequence
		"OutputDestinationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4071),
		"DICOM Storage Sequence", // DICOMStorageSequence
		"DICOMStorageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4072),
		"STOW-RS Storage Sequence", // STOWRSStorageSequence
		"STOWRSStorageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4073),
		"Storage URL", // StorageURL
		"StorageURL",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x4074),
		"XDS Storage Sequence", // XDSStorageSequence
		"XDSStorageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x8302),
		"Entrance Dose in mGy", // EntranceDoseInmGy
		"EntranceDoseInmGy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x8303),
		"Entrance Dose Derivation", // EntranceDoseDerivation
		"EntranceDoseDerivation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9092),
		"Parametric Map Frame Type Sequence", // ParametricMapFrameTypeSequence
		"ParametricMapFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9094),
		"Referenced Image Real World Value Mapping Sequence", // ReferencedImageRealWorldValueMappingSequence
		"ReferencedImageRealWorldValueMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9096),
		"Real World Value Mapping Sequence", // RealWorldValueMappingSequence
		"RealWorldValueMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9098),
		"Pixel Value Mapping Code Sequence", // PixelValueMappingCodeSequence
		"PixelValueMappingCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9210),
		"LUT Label", // LUTLabel
		"LUTLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9211),
		"Real World Value Last Value Mapped", // RealWorldValueLastValueMapped
		"RealWorldValueLastValueMapped",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9212),
		"Real World Value LUT Data", // RealWorldValueLUTData
		"RealWorldValueLUTData",
		vm.VM1_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9213),
		"Double Float Real World Value Last Value Mapped", // DoubleFloatRealWorldValueLastValueMapped
		"DoubleFloatRealWorldValueLastValueMapped",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9214),
		"Double Float Real World Value First Value Mapped", // DoubleFloatRealWorldValueFirstValueMapped
		"DoubleFloatRealWorldValueFirstValueMapped",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9216),
		"Real World Value First Value Mapped", // RealWorldValueFirstValueMapped
		"RealWorldValueFirstValueMapped",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9220),
		"Quantity Definition Sequence", // QuantityDefinitionSequence
		"QuantityDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9224),
		"Real World Value Intercept", // RealWorldValueIntercept
		"RealWorldValueIntercept",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0x9225),
		"Real World Value Slope", // RealWorldValueSlope
		"RealWorldValueSlope",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA007),
		"Findings Flag (Trial)", // FindingsFlagTrial
		"FindingsFlagTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA010),
		"Relationship Type", // RelationshipType
		"RelationshipType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA020),
		"Findings Sequence (Trial)", // FindingsSequenceTrial
		"FindingsSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA021),
		"Findings Group UID (Trial)", // FindingsGroupUIDTrial
		"FindingsGroupUIDTrial",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA022),
		"Referenced Findings Group UID (Trial)", // ReferencedFindingsGroupUIDTrial
		"ReferencedFindingsGroupUIDTrial",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA023),
		"Findings Group Recording Date (Trial)", // FindingsGroupRecordingDateTrial
		"FindingsGroupRecordingDateTrial",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA024),
		"Findings Group Recording Time (Trial)", // FindingsGroupRecordingTimeTrial
		"FindingsGroupRecordingTimeTrial",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA026),
		"Findings Source Category Code Sequence (Trial)", // FindingsSourceCategoryCodeSequenceTrial
		"FindingsSourceCategoryCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA027),
		"Verifying Organization", // VerifyingOrganization
		"VerifyingOrganization",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA028),
		"Documenting Organization Identifier Code Sequence (Trial)", // DocumentingOrganizationIdentifierCodeSequenceTrial
		"DocumentingOrganizationIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA030),
		"Verification DateTime", // VerificationDateTime
		"VerificationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA032),
		"Observation DateTime", // ObservationDateTime
		"ObservationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA033),
		"Observation Start DateTime", // ObservationStartDateTime
		"ObservationStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA034),
		"Effective Start DateTime", // EffectiveStartDateTime
		"EffectiveStartDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA035),
		"Effective Stop DateTime", // EffectiveStopDateTime
		"EffectiveStopDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA040),
		"Value Type", // ValueType
		"ValueType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA043),
		"Concept Name Code Sequence", // ConceptNameCodeSequence
		"ConceptNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA047),
		"Measurement Precision Description (Trial)", // MeasurementPrecisionDescriptionTrial
		"MeasurementPrecisionDescriptionTrial",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA050),
		"Continuity Of Content", // ContinuityOfContent
		"ContinuityOfContent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA057),
		"Urgency or Priority Alerts (Trial)", // UrgencyOrPriorityAlertsTrial
		"UrgencyOrPriorityAlertsTrial",
		vm.VM1_n,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA060),
		"Sequencing Indicator (Trial)", // SequencingIndicatorTrial
		"SequencingIndicatorTrial",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA066),
		"Document Identifier Code Sequence (Trial)", // DocumentIdentifierCodeSequenceTrial
		"DocumentIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA067),
		"Document Author (Trial)", // DocumentAuthorTrial
		"DocumentAuthorTrial",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA068),
		"Document Author Identifier Code Sequence (Trial)", // DocumentAuthorIdentifierCodeSequenceTrial
		"DocumentAuthorIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA070),
		"Identifier Code Sequence (Trial)", // IdentifierCodeSequenceTrial
		"IdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA073),
		"Verifying Observer Sequence", // VerifyingObserverSequence
		"VerifyingObserverSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA074),
		"Object Binary Identifier (Trial)", // ObjectBinaryIdentifierTrial
		"ObjectBinaryIdentifierTrial",
		vm.VM1,
		true,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA075),
		"Verifying Observer Name", // VerifyingObserverName
		"VerifyingObserverName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA076),
		"Documenting Observer Identifier Code Sequence (Trial)", // DocumentingObserverIdentifierCodeSequenceTrial
		"DocumentingObserverIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA078),
		"Author Observer Sequence", // AuthorObserverSequence
		"AuthorObserverSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA07A),
		"Participant Sequence", // ParticipantSequence
		"ParticipantSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA07C),
		"Custodial Organization Sequence", // CustodialOrganizationSequence
		"CustodialOrganizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA080),
		"Participation Type", // ParticipationType
		"ParticipationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA082),
		"Participation DateTime", // ParticipationDateTime
		"ParticipationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA084),
		"Observer Type", // ObserverType
		"ObserverType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA085),
		"Procedure Identifier Code Sequence (Trial)", // ProcedureIdentifierCodeSequenceTrial
		"ProcedureIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA088),
		"Verifying Observer Identification Code Sequence", // VerifyingObserverIdentificationCodeSequence
		"VerifyingObserverIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA089),
		"Object Directory Binary Identifier (Trial)", // ObjectDirectoryBinaryIdentifierTrial
		"ObjectDirectoryBinaryIdentifierTrial",
		vm.VM1,
		true,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA090),
		"Equivalent CDA Document Sequence", // EquivalentCDADocumentSequence
		"EquivalentCDADocumentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA0B0),
		"Referenced Waveform Channels", // ReferencedWaveformChannels
		"ReferencedWaveformChannels",
		vm.VM2_2n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA110),
		"Date of Document or Verbal Transaction (Trial)", // DateOfDocumentOrVerbalTransactionTrial
		"DateOfDocumentOrVerbalTransactionTrial",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA112),
		"Time of Document Creation or Verbal Transaction (Trial)", // TimeOfDocumentCreationOrVerbalTransactionTrial
		"TimeOfDocumentCreationOrVerbalTransactionTrial",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA120),
		"DateTime", // DateTime
		"DateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA121),
		"Date", // Date
		"Date",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA122),
		"Time", // Time
		"Time",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA123),
		"Person Name", // PersonName
		"PersonName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA124),
		"UID", // UID
		"UID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA125),
		"Report Status ID (Trial)", // ReportStatusIDTrial
		"ReportStatusIDTrial",
		vm.VM2,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA130),
		"Temporal Range Type", // TemporalRangeType
		"TemporalRangeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA132),
		"Referenced Sample Positions", // ReferencedSamplePositions
		"ReferencedSamplePositions",
		vm.VM1_n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA136),
		"Referenced Frame Numbers", // ReferencedFrameNumbers
		"ReferencedFrameNumbers",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA138),
		"Referenced Time Offsets", // ReferencedTimeOffsets
		"ReferencedTimeOffsets",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA13A),
		"Referenced DateTime", // ReferencedDateTime
		"ReferencedDateTime",
		vm.VM1_n,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA160),
		"Text Value", // TextValue
		"TextValue",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA161),
		"Floating Point Value", // FloatingPointValue
		"FloatingPointValue",
		vm.VM1_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA162),
		"Rational Numerator Value", // RationalNumeratorValue
		"RationalNumeratorValue",
		vm.VM1_n,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA163),
		"Rational Denominator Value", // RationalDenominatorValue
		"RationalDenominatorValue",
		vm.VM1_n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA167),
		"Observation Category Code Sequence (Trial)", // ObservationCategoryCodeSequenceTrial
		"ObservationCategoryCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA168),
		"Concept Code Sequence", // ConceptCodeSequence
		"ConceptCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA16A),
		"Bibliographic Citation (Trial)", // BibliographicCitationTrial
		"BibliographicCitationTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA170),
		"Purpose of Reference Code Sequence", // PurposeOfReferenceCodeSequence
		"PurposeOfReferenceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA171),
		"Observation UID", // ObservationUID
		"ObservationUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA172),
		"Referenced Observation UID (Trial)", // ReferencedObservationUIDTrial
		"ReferencedObservationUIDTrial",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA173),
		"Referenced Observation Class (Trial)", // ReferencedObservationClassTrial
		"ReferencedObservationClassTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA174),
		"Referenced Object Observation Class (Trial)", // ReferencedObjectObservationClassTrial
		"ReferencedObjectObservationClassTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA180),
		"Annotation Group Number", // AnnotationGroupNumber
		"AnnotationGroupNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA192),
		"Observation Date (Trial)", // ObservationDateTrial
		"ObservationDateTrial",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA193),
		"Observation Time (Trial)", // ObservationTimeTrial
		"ObservationTimeTrial",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA194),
		"Measurement Automation (Trial)", // MeasurementAutomationTrial
		"MeasurementAutomationTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA195),
		"Modifier Code Sequence", // ModifierCodeSequence
		"ModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA224),
		"Identification Description (Trial)", // IdentificationDescriptionTrial
		"IdentificationDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA290),
		"Coordinates Set Geometric Type (Trial)", // CoordinatesSetGeometricTypeTrial
		"CoordinatesSetGeometricTypeTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA296),
		"Algorithm Code Sequence (Trial)", // AlgorithmCodeSequenceTrial
		"AlgorithmCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA297),
		"Algorithm Description (Trial)", // AlgorithmDescriptionTrial
		"AlgorithmDescriptionTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA29A),
		"Pixel Coordinates Set (Trial)", // PixelCoordinatesSetTrial
		"PixelCoordinatesSetTrial",
		vm.VM2_2n,
		true,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA300),
		"Measured Value Sequence", // MeasuredValueSequence
		"MeasuredValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA301),
		"Numeric Value Qualifier Code Sequence", // NumericValueQualifierCodeSequence
		"NumericValueQualifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA307),
		"Current Observer (Trial)", // CurrentObserverTrial
		"CurrentObserverTrial",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA30A),
		"Numeric Value", // NumericValue
		"NumericValue",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA313),
		"Referenced Accession Sequence (Trial)", // ReferencedAccessionSequenceTrial
		"ReferencedAccessionSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA33A),
		"Report Status Comment (Trial)", // ReportStatusCommentTrial
		"ReportStatusCommentTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA340),
		"Procedure Context Sequence (Trial)", // ProcedureContextSequenceTrial
		"ProcedureContextSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA352),
		"Verbal Source (Trial)", // VerbalSourceTrial
		"VerbalSourceTrial",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA353),
		"Address (Trial)", // AddressTrial
		"AddressTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA354),
		"Telephone Number (Trial)", // TelephoneNumberTrial
		"TelephoneNumberTrial",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA358),
		"Verbal Source Identifier Code Sequence (Trial)", // VerbalSourceIdentifierCodeSequenceTrial
		"VerbalSourceIdentifierCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA360),
		"Predecessor Documents Sequence", // PredecessorDocumentsSequence
		"PredecessorDocumentsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA370),
		"Referenced Request Sequence", // ReferencedRequestSequence
		"ReferencedRequestSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA372),
		"Performed Procedure Code Sequence", // PerformedProcedureCodeSequence
		"PerformedProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA375),
		"Current Requested Procedure Evidence Sequence", // CurrentRequestedProcedureEvidenceSequence
		"CurrentRequestedProcedureEvidenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA380),
		"Report Detail Sequence (Trial)", // ReportDetailSequenceTrial
		"ReportDetailSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA385),
		"Pertinent Other Evidence Sequence", // PertinentOtherEvidenceSequence
		"PertinentOtherEvidenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA390),
		"HL7 Structured Document Reference Sequence", // HL7StructuredDocumentReferenceSequence
		"HL7StructuredDocumentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA402),
		"Observation Subject UID (Trial)", // ObservationSubjectUIDTrial
		"ObservationSubjectUIDTrial",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA403),
		"Observation Subject Class (Trial)", // ObservationSubjectClassTrial
		"ObservationSubjectClassTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA404),
		"Observation Subject Type Code Sequence (Trial)", // ObservationSubjectTypeCodeSequenceTrial
		"ObservationSubjectTypeCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA491),
		"Completion Flag", // CompletionFlag
		"CompletionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA492),
		"Completion Flag Description", // CompletionFlagDescription
		"CompletionFlagDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA493),
		"Verification Flag", // VerificationFlag
		"VerificationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA494),
		"Archive Requested", // ArchiveRequested
		"ArchiveRequested",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA496),
		"Preliminary Flag", // PreliminaryFlag
		"PreliminaryFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA504),
		"Content Template Sequence", // ContentTemplateSequence
		"ContentTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA525),
		"Identical Documents Sequence", // IdenticalDocumentsSequence
		"IdenticalDocumentsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA600),
		"Observation Subject Context Flag (Trial)", // ObservationSubjectContextFlagTrial
		"ObservationSubjectContextFlagTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA601),
		"Observer Context Flag (Trial)", // ObserverContextFlagTrial
		"ObserverContextFlagTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA603),
		"Procedure Context Flag (Trial)", // ProcedureContextFlagTrial
		"ProcedureContextFlagTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA730),
		"Content Sequence", // ContentSequence
		"ContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA731),
		"Relationship Sequence (Trial)", // RelationshipSequenceTrial
		"RelationshipSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA732),
		"Relationship Type Code Sequence (Trial)", // RelationshipTypeCodeSequenceTrial
		"RelationshipTypeCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA744),
		"Language Code Sequence (Trial)", // LanguageCodeSequenceTrial
		"LanguageCodeSequenceTrial",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA801),
		"Tabulated Values Sequence", // TabulatedValuesSequence
		"TabulatedValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA802),
		"Number of Table Rows", // NumberOfTableRows
		"NumberOfTableRows",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA803),
		"Number of Table Columns", // NumberOfTableColumns
		"NumberOfTableColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA804),
		"Table Row Number", // TableRowNumber
		"TableRowNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA805),
		"Table Column Number", // TableColumnNumber
		"TableColumnNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA806),
		"Table Row Definition Sequence", // TableRowDefinitionSequence
		"TableRowDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA807),
		"Table Column Definition Sequence", // TableColumnDefinitionSequence
		"TableColumnDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA808),
		"Cell Values Sequence", // CellValuesSequence
		"CellValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xA992),
		"Uniform Resource Locator (Trial)", // UniformResourceLocatorTrial
		"UniformResourceLocatorTrial",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB020),
		"Waveform Annotation Sequence", // WaveformAnnotationSequence
		"WaveformAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB030),
		"Structured Waveform Annotation Sequence", // StructuredWaveformAnnotationSequence
		"StructuredWaveformAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB031),
		"Waveform Annotation Display Selection Sequence", // WaveformAnnotationDisplaySelectionSequence
		"WaveformAnnotationDisplaySelectionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB032),
		"Referenced Montage Index", // ReferencedMontageIndex
		"ReferencedMontageIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB033),
		"Waveform Textual Annotation Sequence", // WaveformTextualAnnotationSequence
		"WaveformTextualAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB034),
		"Annotation DateTime", // AnnotationDateTime
		"AnnotationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB035),
		"Displayed Waveform Segment Sequence", // DisplayedWaveformSegmentSequence
		"DisplayedWaveformSegmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB036),
		"Segment Definition DateTime", // SegmentDefinitionDateTime
		"SegmentDefinitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB037),
		"Montage Activation Sequence", // MontageActivationSequence
		"MontageActivationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB038),
		"Montage Activation Time Offset", // MontageActivationTimeOffset
		"MontageActivationTimeOffset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB039),
		"Waveform Montage Sequence", // WaveformMontageSequence
		"WaveformMontageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03A),
		"Referenced Montage Channel Number", // ReferencedMontageChannelNumber
		"ReferencedMontageChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03B),
		"Montage Name", // MontageName
		"MontageName",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03C),
		"Montage Channel Sequence", // MontageChannelSequence
		"MontageChannelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03D),
		"Montage Index", // MontageIndex
		"MontageIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03E),
		"Montage Channel Number", // MontageChannelNumber
		"MontageChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB03F),
		"Montage Channel Label", // MontageChannelLabel
		"MontageChannelLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB040),
		"Montage Channel Source Code Sequence", // MontageChannelSourceCodeSequence
		"MontageChannelSourceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB041),
		"Contributing Channel Sources Sequence", // ContributingChannelSourcesSequence
		"ContributingChannelSourcesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xB042),
		"Channel Weight", // ChannelWeight
		"ChannelWeight",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB00),
		"Template Identifier", // TemplateIdentifier
		"TemplateIdentifier",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB06),
		"Template Version", // TemplateVersion
		"TemplateVersion",
		vm.VM1,
		true,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB07),
		"Template Local Version", // TemplateLocalVersion
		"TemplateLocalVersion",
		vm.VM1,
		true,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB0B),
		"Template Extension Flag", // TemplateExtensionFlag
		"TemplateExtensionFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB0C),
		"Template Extension Organization UID", // TemplateExtensionOrganizationUID
		"TemplateExtensionOrganizationUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB0D),
		"Template Extension Creator UID", // TemplateExtensionCreatorUID
		"TemplateExtensionCreatorUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xDB73),
		"Referenced Content Item Identifier", // ReferencedContentItemIdentifier
		"ReferencedContentItemIdentifier",
		vm.VM1_n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE001),
		"HL7 Instance Identifier", // HL7InstanceIdentifier
		"HL7InstanceIdentifier",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE004),
		"HL7 Document Effective Time", // HL7DocumentEffectiveTime
		"HL7DocumentEffectiveTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE006),
		"HL7 Document Type Code Sequence", // HL7DocumentTypeCodeSequence
		"HL7DocumentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE008),
		"Document Class Code Sequence", // DocumentClassCodeSequence
		"DocumentClassCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE010),
		"Retrieve URI", // RetrieveURI
		"RetrieveURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE011),
		"Retrieve Location UID", // RetrieveLocationUID
		"RetrieveLocationUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE020),
		"Type of Instances", // TypeOfInstances
		"TypeOfInstances",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE021),
		"DICOM Retrieval Sequence", // DICOMRetrievalSequence
		"DICOMRetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE022),
		"DICOM Media Retrieval Sequence", // DICOMMediaRetrievalSequence
		"DICOMMediaRetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE023),
		"WADO Retrieval Sequence", // WADORetrievalSequence
		"WADORetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE024),
		"XDS Retrieval Sequence", // XDSRetrievalSequence
		"XDSRetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE025),
		"WADO-RS Retrieval Sequence", // WADORSRetrievalSequence
		"WADORSRetrievalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE030),
		"Repository Unique ID", // RepositoryUniqueID
		"RepositoryUniqueID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0040, 0xE031),
		"Home Community ID", // HomeCommunityID
		"HomeCommunityID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0010),
		"Document Title", // DocumentTitle
		"DocumentTitle",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0011),
		"Encapsulated Document", // EncapsulatedDocument
		"EncapsulatedDocument",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0012),
		"MIME Type of Encapsulated Document", // MIMETypeOfEncapsulatedDocument
		"MIMETypeOfEncapsulatedDocument",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0013),
		"Source Instance Sequence", // SourceInstanceSequence
		"SourceInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0014),
		"List of MIME Types", // ListOfMIMETypes
		"ListOfMIMETypes",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0042, 0x0015),
		"Encapsulated Document Length", // EncapsulatedDocumentLength
		"EncapsulatedDocumentLength",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0001),
		"Product Package Identifier", // ProductPackageIdentifier
		"ProductPackageIdentifier",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0002),
		"Substance Administration Approval", // SubstanceAdministrationApproval
		"SubstanceAdministrationApproval",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0003),
		"Approval Status Further Description", // ApprovalStatusFurtherDescription
		"ApprovalStatusFurtherDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0004),
		"Approval Status DateTime", // ApprovalStatusDateTime
		"ApprovalStatusDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0007),
		"Product Type Code Sequence", // ProductTypeCodeSequence
		"ProductTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0008),
		"Product Name", // ProductName
		"ProductName",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0009),
		"Product Description", // ProductDescription
		"ProductDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x000A),
		"Product Lot Identifier", // ProductLotIdentifier
		"ProductLotIdentifier",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x000B),
		"Product Expiration DateTime", // ProductExpirationDateTime
		"ProductExpirationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0010),
		"Substance Administration DateTime", // SubstanceAdministrationDateTime
		"SubstanceAdministrationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0011),
		"Substance Administration Notes", // SubstanceAdministrationNotes
		"SubstanceAdministrationNotes",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0012),
		"Substance Administration Device ID", // SubstanceAdministrationDeviceID
		"SubstanceAdministrationDeviceID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0013),
		"Product Parameter Sequence", // ProductParameterSequence
		"ProductParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0019),
		"Substance Administration Parameter Sequence", // SubstanceAdministrationParameterSequence
		"SubstanceAdministrationParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0100),
		"Approval Sequence", // ApprovalSequence
		"ApprovalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0101),
		"Assertion Code Sequence", // AssertionCodeSequence
		"AssertionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0102),
		"Assertion UID", // AssertionUID
		"AssertionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0103),
		"Asserter Identification Sequence", // AsserterIdentificationSequence
		"AsserterIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0104),
		"Assertion DateTime", // AssertionDateTime
		"AssertionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0105),
		"Assertion Expiration DateTime", // AssertionExpirationDateTime
		"AssertionExpirationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0106),
		"Assertion Comments", // AssertionComments
		"AssertionComments",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0107),
		"Related Assertion Sequence", // RelatedAssertionSequence
		"RelatedAssertionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0108),
		"Referenced Assertion UID", // ReferencedAssertionUID
		"ReferencedAssertionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0109),
		"Approval Subject Sequence", // ApprovalSubjectSequence
		"ApprovalSubjectSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x010A),
		"Organizational Role Code Sequence", // OrganizationalRoleCodeSequence
		"OrganizationalRoleCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0044, 0x0110),
		"RT Assertions Sequence", // RTAssertionsSequence
		"RTAssertionsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0012),
		"Lens Description", // LensDescription
		"LensDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0014),
		"Right Lens Sequence", // RightLensSequence
		"RightLensSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0015),
		"Left Lens Sequence", // LeftLensSequence
		"LeftLensSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0016),
		"Unspecified Laterality Lens Sequence", // UnspecifiedLateralityLensSequence
		"UnspecifiedLateralityLensSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0018),
		"Cylinder Sequence", // CylinderSequence
		"CylinderSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0028),
		"Prism Sequence", // PrismSequence
		"PrismSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0030),
		"Horizontal Prism Power", // HorizontalPrismPower
		"HorizontalPrismPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0032),
		"Horizontal Prism Base", // HorizontalPrismBase
		"HorizontalPrismBase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0034),
		"Vertical Prism Power", // VerticalPrismPower
		"VerticalPrismPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0036),
		"Vertical Prism Base", // VerticalPrismBase
		"VerticalPrismBase",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0038),
		"Lens Segment Type", // LensSegmentType
		"LensSegmentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0040),
		"Optical Transmittance", // OpticalTransmittance
		"OpticalTransmittance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0042),
		"Channel Width", // ChannelWidth
		"ChannelWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0044),
		"Pupil Size", // PupilSize
		"PupilSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0046),
		"Corneal Size", // CornealSize
		"CornealSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0047),
		"Corneal Size Sequence", // CornealSizeSequence
		"CornealSizeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0050),
		"Autorefraction Right Eye Sequence", // AutorefractionRightEyeSequence
		"AutorefractionRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0052),
		"Autorefraction Left Eye Sequence", // AutorefractionLeftEyeSequence
		"AutorefractionLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0060),
		"Distance Pupillary Distance", // DistancePupillaryDistance
		"DistancePupillaryDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0062),
		"Near Pupillary Distance", // NearPupillaryDistance
		"NearPupillaryDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0063),
		"Intermediate Pupillary Distance", // IntermediatePupillaryDistance
		"IntermediatePupillaryDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0064),
		"Other Pupillary Distance", // OtherPupillaryDistance
		"OtherPupillaryDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0070),
		"Keratometry Right Eye Sequence", // KeratometryRightEyeSequence
		"KeratometryRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0071),
		"Keratometry Left Eye Sequence", // KeratometryLeftEyeSequence
		"KeratometryLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0074),
		"Steep Keratometric Axis Sequence", // SteepKeratometricAxisSequence
		"SteepKeratometricAxisSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0075),
		"Radius of Curvature", // RadiusOfCurvature
		"RadiusOfCurvature",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0076),
		"Keratometric Power", // KeratometricPower
		"KeratometricPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0077),
		"Keratometric Axis", // KeratometricAxis
		"KeratometricAxis",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0080),
		"Flat Keratometric Axis Sequence", // FlatKeratometricAxisSequence
		"FlatKeratometricAxisSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0092),
		"Background Color", // BackgroundColor
		"BackgroundColor",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0094),
		"Optotype", // Optotype
		"Optotype",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0095),
		"Optotype Presentation", // OptotypePresentation
		"OptotypePresentation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0097),
		"Subjective Refraction Right Eye Sequence", // SubjectiveRefractionRightEyeSequence
		"SubjectiveRefractionRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0098),
		"Subjective Refraction Left Eye Sequence", // SubjectiveRefractionLeftEyeSequence
		"SubjectiveRefractionLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0100),
		"Add Near Sequence", // AddNearSequence
		"AddNearSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0101),
		"Add Intermediate Sequence", // AddIntermediateSequence
		"AddIntermediateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0102),
		"Add Other Sequence", // AddOtherSequence
		"AddOtherSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0104),
		"Add Power", // AddPower
		"AddPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0106),
		"Viewing Distance", // ViewingDistance
		"ViewingDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0110),
		"Cornea Measurements Sequence", // CorneaMeasurementsSequence
		"CorneaMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0111),
		"Source of Cornea Measurement Data Code Sequence", // SourceOfCorneaMeasurementDataCodeSequence
		"SourceOfCorneaMeasurementDataCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0112),
		"Steep Corneal Axis Sequence", // SteepCornealAxisSequence
		"SteepCornealAxisSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0113),
		"Flat Corneal Axis Sequence", // FlatCornealAxisSequence
		"FlatCornealAxisSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0114),
		"Corneal Power", // CornealPower
		"CornealPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0115),
		"Corneal Axis", // CornealAxis
		"CornealAxis",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0116),
		"Cornea Measurement Method Code Sequence", // CorneaMeasurementMethodCodeSequence
		"CorneaMeasurementMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0117),
		"Refractive Index of Cornea", // RefractiveIndexOfCornea
		"RefractiveIndexOfCornea",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0118),
		"Refractive Index of Aqueous Humor", // RefractiveIndexOfAqueousHumor
		"RefractiveIndexOfAqueousHumor",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0121),
		"Visual Acuity Type Code Sequence", // VisualAcuityTypeCodeSequence
		"VisualAcuityTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0122),
		"Visual Acuity Right Eye Sequence", // VisualAcuityRightEyeSequence
		"VisualAcuityRightEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0123),
		"Visual Acuity Left Eye Sequence", // VisualAcuityLeftEyeSequence
		"VisualAcuityLeftEyeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0124),
		"Visual Acuity Both Eyes Open Sequence", // VisualAcuityBothEyesOpenSequence
		"VisualAcuityBothEyesOpenSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0125),
		"Viewing Distance Type", // ViewingDistanceType
		"ViewingDistanceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0135),
		"Visual Acuity Modifiers", // VisualAcuityModifiers
		"VisualAcuityModifiers",
		vm.VM2,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0137),
		"Decimal Visual Acuity", // DecimalVisualAcuity
		"DecimalVisualAcuity",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0139),
		"Optotype Detailed Definition", // OptotypeDetailedDefinition
		"OptotypeDetailedDefinition",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0145),
		"Referenced Refractive Measurements Sequence", // ReferencedRefractiveMeasurementsSequence
		"ReferencedRefractiveMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0146),
		"Sphere Power", // SpherePower
		"SpherePower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0147),
		"Cylinder Power", // CylinderPower
		"CylinderPower",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0201),
		"Corneal Topography Surface", // CornealTopographySurface
		"CornealTopographySurface",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0202),
		"Corneal Vertex Location", // CornealVertexLocation
		"CornealVertexLocation",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0203),
		"Pupil Centroid X-Coordinate", // PupilCentroidXCoordinate
		"PupilCentroidXCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0204),
		"Pupil Centroid Y-Coordinate", // PupilCentroidYCoordinate
		"PupilCentroidYCoordinate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0205),
		"Equivalent Pupil Radius", // EquivalentPupilRadius
		"EquivalentPupilRadius",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0207),
		"Corneal Topography Map Type Code Sequence", // CornealTopographyMapTypeCodeSequence
		"CornealTopographyMapTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0208),
		"Vertices of the Outline of Pupil", // VerticesOfTheOutlineOfPupil
		"VerticesOfTheOutlineOfPupil",
		vm.VM2_2n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0210),
		"Corneal Topography Mapping Normals Sequence", // CornealTopographyMappingNormalsSequence
		"CornealTopographyMappingNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0211),
		"Maximum Corneal Curvature Sequence", // MaximumCornealCurvatureSequence
		"MaximumCornealCurvatureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0212),
		"Maximum Corneal Curvature", // MaximumCornealCurvature
		"MaximumCornealCurvature",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0213),
		"Maximum Corneal Curvature Location", // MaximumCornealCurvatureLocation
		"MaximumCornealCurvatureLocation",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0215),
		"Minimum Keratometric Sequence", // MinimumKeratometricSequence
		"MinimumKeratometricSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0218),
		"Simulated Keratometric Cylinder Sequence", // SimulatedKeratometricCylinderSequence
		"SimulatedKeratometricCylinderSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0220),
		"Average Corneal Power", // AverageCornealPower
		"AverageCornealPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0224),
		"Corneal I-S Value", // CornealISValue
		"CornealISValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0227),
		"Analyzed Area", // AnalyzedArea
		"AnalyzedArea",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0230),
		"Surface Regularity Index", // SurfaceRegularityIndex
		"SurfaceRegularityIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0232),
		"Surface Asymmetry Index", // SurfaceAsymmetryIndex
		"SurfaceAsymmetryIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0234),
		"Corneal Eccentricity Index", // CornealEccentricityIndex
		"CornealEccentricityIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0236),
		"Keratoconus Prediction Index", // KeratoconusPredictionIndex
		"KeratoconusPredictionIndex",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0238),
		"Decimal Potential Visual Acuity", // DecimalPotentialVisualAcuity
		"DecimalPotentialVisualAcuity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0242),
		"Corneal Topography Map Quality Evaluation", // CornealTopographyMapQualityEvaluation
		"CornealTopographyMapQualityEvaluation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0244),
		"Source Image Corneal Processed Data Sequence", // SourceImageCornealProcessedDataSequence
		"SourceImageCornealProcessedDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0247),
		"Corneal Point Location", // CornealPointLocation
		"CornealPointLocation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0248),
		"Corneal Point Estimated", // CornealPointEstimated
		"CornealPointEstimated",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0249),
		"Axial Power", // AxialPower
		"AxialPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0250),
		"Tangential Power", // TangentialPower
		"TangentialPower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0251),
		"Refractive Power", // RefractivePower
		"RefractivePower",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0252),
		"Relative Elevation", // RelativeElevation
		"RelativeElevation",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0046, 0x0253),
		"Corneal Wavefront", // CornealWavefront
		"CornealWavefront",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0001),
		"Imaged Volume Width", // ImagedVolumeWidth
		"ImagedVolumeWidth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0002),
		"Imaged Volume Height", // ImagedVolumeHeight
		"ImagedVolumeHeight",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0003),
		"Imaged Volume Depth", // ImagedVolumeDepth
		"ImagedVolumeDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0006),
		"Total Pixel Matrix Columns", // TotalPixelMatrixColumns
		"TotalPixelMatrixColumns",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0007),
		"Total Pixel Matrix Rows", // TotalPixelMatrixRows
		"TotalPixelMatrixRows",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0008),
		"Total Pixel Matrix Origin Sequence", // TotalPixelMatrixOriginSequence
		"TotalPixelMatrixOriginSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0010),
		"Specimen Label in Image", // SpecimenLabelInImage
		"SpecimenLabelInImage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0011),
		"Focus Method", // FocusMethod
		"FocusMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0012),
		"Extended Depth of Field", // ExtendedDepthOfField
		"ExtendedDepthOfField",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0013),
		"Number of Focal Planes", // NumberOfFocalPlanes
		"NumberOfFocalPlanes",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0014),
		"Distance Between Focal Planes", // DistanceBetweenFocalPlanes
		"DistanceBetweenFocalPlanes",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0015),
		"Recommended Absent Pixel CIELab Value", // RecommendedAbsentPixelCIELabValue
		"RecommendedAbsentPixelCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0100),
		"Illuminator Type Code Sequence", // IlluminatorTypeCodeSequence
		"IlluminatorTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0102),
		"Image Orientation (Slide)", // ImageOrientationSlide
		"ImageOrientationSlide",
		vm.VM6,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0105),
		"Optical Path Sequence", // OpticalPathSequence
		"OpticalPathSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0106),
		"Optical Path Identifier", // OpticalPathIdentifier
		"OpticalPathIdentifier",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0107),
		"Optical Path Description", // OpticalPathDescription
		"OpticalPathDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0108),
		"Illumination Color Code Sequence", // IlluminationColorCodeSequence
		"IlluminationColorCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0110),
		"Specimen Reference Sequence", // SpecimenReferenceSequence
		"SpecimenReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0111),
		"Condenser Lens Power", // CondenserLensPower
		"CondenserLensPower",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0112),
		"Objective Lens Power", // ObjectiveLensPower
		"ObjectiveLensPower",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0113),
		"Objective Lens Numerical Aperture", // ObjectiveLensNumericalAperture
		"ObjectiveLensNumericalAperture",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0114),
		"Confocal Mode", // ConfocalMode
		"ConfocalMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0115),
		"Tissue Location", // TissueLocation
		"TissueLocation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0116),
		"Confocal Microscopy Image Frame Type Sequence", // ConfocalMicroscopyImageFrameTypeSequence
		"ConfocalMicroscopyImageFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0117),
		"Image Acquisition Depth", // ImageAcquisitionDepth
		"ImageAcquisitionDepth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0120),
		"Palette Color Lookup Table Sequence", // PaletteColorLookupTableSequence
		"PaletteColorLookupTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0200),
		"Referenced Image Navigation Sequence", // ReferencedImageNavigationSequence
		"ReferencedImageNavigationSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0201),
		"Top Left Hand Corner of Localizer Area", // TopLeftHandCornerOfLocalizerArea
		"TopLeftHandCornerOfLocalizerArea",
		vm.VM2,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0202),
		"Bottom Right Hand Corner of Localizer Area", // BottomRightHandCornerOfLocalizerArea
		"BottomRightHandCornerOfLocalizerArea",
		vm.VM2,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0207),
		"Optical Path Identification Sequence", // OpticalPathIdentificationSequence
		"OpticalPathIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x021A),
		"Plane Position (Slide) Sequence", // PlanePositionSlideSequence
		"PlanePositionSlideSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x021E),
		"Column Position In Total Image Pixel Matrix", // ColumnPositionInTotalImagePixelMatrix
		"ColumnPositionInTotalImagePixelMatrix",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x021F),
		"Row Position In Total Image Pixel Matrix", // RowPositionInTotalImagePixelMatrix
		"RowPositionInTotalImagePixelMatrix",
		vm.VM1,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0301),
		"Pixel Origin Interpretation", // PixelOriginInterpretation
		"PixelOriginInterpretation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0302),
		"Number of Optical Paths", // NumberOfOpticalPaths
		"NumberOfOpticalPaths",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0303),
		"Total Pixel Matrix Focal Planes", // TotalPixelMatrixFocalPlanes
		"TotalPixelMatrixFocalPlanes",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0048, 0x0304),
		"Tiles Overlap", // TilesOverlap
		"TilesOverlap",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0004),
		"Calibration Image", // CalibrationImage
		"CalibrationImage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0010),
		"Device Sequence", // DeviceSequence
		"DeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0012),
		"Container Component Type Code Sequence", // ContainerComponentTypeCodeSequence
		"ContainerComponentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0013),
		"Container Component Thickness", // ContainerComponentThickness
		"ContainerComponentThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0014),
		"Device Length", // DeviceLength
		"DeviceLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0015),
		"Container Component Width", // ContainerComponentWidth
		"ContainerComponentWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0016),
		"Device Diameter", // DeviceDiameter
		"DeviceDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0017),
		"Device Diameter Units", // DeviceDiameterUnits
		"DeviceDiameterUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0018),
		"Device Volume", // DeviceVolume
		"DeviceVolume",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0019),
		"Inter-Marker Distance", // InterMarkerDistance
		"InterMarkerDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001A),
		"Container Component Material", // ContainerComponentMaterial
		"ContainerComponentMaterial",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001B),
		"Container Component ID", // ContainerComponentID
		"ContainerComponentID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001C),
		"Container Component Length", // ContainerComponentLength
		"ContainerComponentLength",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001D),
		"Container Component Diameter", // ContainerComponentDiameter
		"ContainerComponentDiameter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x001E),
		"Container Component Description", // ContainerComponentDescription
		"ContainerComponentDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0020),
		"Device Description", // DeviceDescription
		"DeviceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0050, 0x0021),
		"Long Device Description", // LongDeviceDescription
		"LongDeviceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0001),
		"Contrast/Bolus Ingredient Percent by Volume", // ContrastBolusIngredientPercentByVolume
		"ContrastBolusIngredientPercentByVolume",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0002),
		"OCT Focal Distance", // OCTFocalDistance
		"OCTFocalDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0003),
		"Beam Spot Size", // BeamSpotSize
		"BeamSpotSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0004),
		"Effective Refractive Index", // EffectiveRefractiveIndex
		"EffectiveRefractiveIndex",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0006),
		"OCT Acquisition Domain", // OCTAcquisitionDomain
		"OCTAcquisitionDomain",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0007),
		"OCT Optical Center Wavelength", // OCTOpticalCenterWavelength
		"OCTOpticalCenterWavelength",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0008),
		"Axial Resolution", // AxialResolution
		"AxialResolution",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0009),
		"Ranging Depth", // RangingDepth
		"RangingDepth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0011),
		"A-line Rate", // ALineRate
		"ALineRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0012),
		"A-lines Per Frame", // ALinesPerFrame
		"ALinesPerFrame",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0013),
		"Catheter Rotational Rate", // CatheterRotationalRate
		"CatheterRotationalRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0014),
		"A-line Pixel Spacing", // ALinePixelSpacing
		"ALinePixelSpacing",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0016),
		"Mode of Percutaneous Access Sequence", // ModeOfPercutaneousAccessSequence
		"ModeOfPercutaneousAccessSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0025),
		"Intravascular OCT Frame Type Sequence", // IntravascularOCTFrameTypeSequence
		"IntravascularOCTFrameTypeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0026),
		"OCT Z Offset Applied", // OCTZOffsetApplied
		"OCTZOffsetApplied",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0027),
		"Intravascular Frame Content Sequence", // IntravascularFrameContentSequence
		"IntravascularFrameContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0028),
		"Intravascular Longitudinal Distance", // IntravascularLongitudinalDistance
		"IntravascularLongitudinalDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0029),
		"Intravascular OCT Frame Content Sequence", // IntravascularOCTFrameContentSequence
		"IntravascularOCTFrameContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0030),
		"OCT Z Offset Correction", // OCTZOffsetCorrection
		"OCTZOffsetCorrection",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0031),
		"Catheter Direction of Rotation", // CatheterDirectionOfRotation
		"CatheterDirectionOfRotation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0033),
		"Seam Line Location", // SeamLineLocation
		"SeamLineLocation",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0034),
		"First A-line Location", // FirstALineLocation
		"FirstALineLocation",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0036),
		"Seam Line Index", // SeamLineIndex
		"SeamLineIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0038),
		"Number of Padded A-lines", // NumberOfPaddedALines
		"NumberOfPaddedALines",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x0039),
		"Interpolation Type", // InterpolationType
		"InterpolationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0052, 0x003A),
		"Refractive Index Applied", // RefractiveIndexApplied
		"RefractiveIndexApplied",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0010),
		"Energy Window Vector", // EnergyWindowVector
		"EnergyWindowVector",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0011),
		"Number of Energy Windows", // NumberOfEnergyWindows
		"NumberOfEnergyWindows",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0012),
		"Energy Window Information Sequence", // EnergyWindowInformationSequence
		"EnergyWindowInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0013),
		"Energy Window Range Sequence", // EnergyWindowRangeSequence
		"EnergyWindowRangeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0014),
		"Energy Window Lower Limit", // EnergyWindowLowerLimit
		"EnergyWindowLowerLimit",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0015),
		"Energy Window Upper Limit", // EnergyWindowUpperLimit
		"EnergyWindowUpperLimit",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0016),
		"Radiopharmaceutical Information Sequence", // RadiopharmaceuticalInformationSequence
		"RadiopharmaceuticalInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0017),
		"Residual Syringe Counts", // ResidualSyringeCounts
		"ResidualSyringeCounts",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0018),
		"Energy Window Name", // EnergyWindowName
		"EnergyWindowName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0020),
		"Detector Vector", // DetectorVector
		"DetectorVector",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0021),
		"Number of Detectors", // NumberOfDetectors
		"NumberOfDetectors",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0022),
		"Detector Information Sequence", // DetectorInformationSequence
		"DetectorInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0030),
		"Phase Vector", // PhaseVector
		"PhaseVector",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0031),
		"Number of Phases", // NumberOfPhases
		"NumberOfPhases",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0032),
		"Phase Information Sequence", // PhaseInformationSequence
		"PhaseInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0033),
		"Number of Frames in Phase", // NumberOfFramesInPhase
		"NumberOfFramesInPhase",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0036),
		"Phase Delay", // PhaseDelay
		"PhaseDelay",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0038),
		"Pause Between Frames", // PauseBetweenFrames
		"PauseBetweenFrames",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0039),
		"Phase Description", // PhaseDescription
		"PhaseDescription",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0050),
		"Rotation Vector", // RotationVector
		"RotationVector",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0051),
		"Number of Rotations", // NumberOfRotations
		"NumberOfRotations",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0052),
		"Rotation Information Sequence", // RotationInformationSequence
		"RotationInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0053),
		"Number of Frames in Rotation", // NumberOfFramesInRotation
		"NumberOfFramesInRotation",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0060),
		"R-R Interval Vector", // RRIntervalVector
		"RRIntervalVector",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0061),
		"Number of R-R Intervals", // NumberOfRRIntervals
		"NumberOfRRIntervals",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0062),
		"Gated Information Sequence", // GatedInformationSequence
		"GatedInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0063),
		"Data Information Sequence", // DataInformationSequence
		"DataInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0070),
		"Time Slot Vector", // TimeSlotVector
		"TimeSlotVector",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0071),
		"Number of Time Slots", // NumberOfTimeSlots
		"NumberOfTimeSlots",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0072),
		"Time Slot Information Sequence", // TimeSlotInformationSequence
		"TimeSlotInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0073),
		"Time Slot Time", // TimeSlotTime
		"TimeSlotTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0080),
		"Slice Vector", // SliceVector
		"SliceVector",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0081),
		"Number of Slices", // NumberOfSlices
		"NumberOfSlices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0090),
		"Angular View Vector", // AngularViewVector
		"AngularViewVector",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0100),
		"Time Slice Vector", // TimeSliceVector
		"TimeSliceVector",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0101),
		"Number of Time Slices", // NumberOfTimeSlices
		"NumberOfTimeSlices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0200),
		"Start Angle", // StartAngle
		"StartAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0202),
		"Type of Detector Motion", // TypeOfDetectorMotion
		"TypeOfDetectorMotion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0210),
		"Trigger Vector", // TriggerVector
		"TriggerVector",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0211),
		"Number of Triggers in Phase", // NumberOfTriggersInPhase
		"NumberOfTriggersInPhase",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0220),
		"View Code Sequence", // ViewCodeSequence
		"ViewCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0222),
		"View Modifier Code Sequence", // ViewModifierCodeSequence
		"ViewModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0300),
		"Radionuclide Code Sequence", // RadionuclideCodeSequence
		"RadionuclideCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0302),
		"Administration Route Code Sequence", // AdministrationRouteCodeSequence
		"AdministrationRouteCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0304),
		"Radiopharmaceutical Code Sequence", // RadiopharmaceuticalCodeSequence
		"RadiopharmaceuticalCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0306),
		"Calibration Data Sequence", // CalibrationDataSequence
		"CalibrationDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0308),
		"Energy Window Number", // EnergyWindowNumber
		"EnergyWindowNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0400),
		"Image ID", // ImageID
		"ImageID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0410),
		"Patient Orientation Code Sequence", // PatientOrientationCodeSequence
		"PatientOrientationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0412),
		"Patient Orientation Modifier Code Sequence", // PatientOrientationModifierCodeSequence
		"PatientOrientationModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0414),
		"Patient Gantry Relationship Code Sequence", // PatientGantryRelationshipCodeSequence
		"PatientGantryRelationshipCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0500),
		"Slice Progression Direction", // SliceProgressionDirection
		"SliceProgressionDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x0501),
		"Scan Progression Direction", // ScanProgressionDirection
		"ScanProgressionDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1000),
		"Series Type", // SeriesType
		"SeriesType",
		vm.VM2,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1001),
		"Units", // Units
		"Units",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1002),
		"Counts Source", // CountsSource
		"CountsSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1004),
		"Reprojection Method", // ReprojectionMethod
		"ReprojectionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1006),
		"SUV Type", // SUVType
		"SUVType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1100),
		"Randoms Correction Method", // RandomsCorrectionMethod
		"RandomsCorrectionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1101),
		"Attenuation Correction Method", // AttenuationCorrectionMethod
		"AttenuationCorrectionMethod",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1102),
		"Decay Correction", // DecayCorrection
		"DecayCorrection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1103),
		"Reconstruction Method", // ReconstructionMethod
		"ReconstructionMethod",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1104),
		"Detector Lines of Response Used", // DetectorLinesOfResponseUsed
		"DetectorLinesOfResponseUsed",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1105),
		"Scatter Correction Method", // ScatterCorrectionMethod
		"ScatterCorrectionMethod",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1200),
		"Axial Acceptance", // AxialAcceptance
		"AxialAcceptance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1201),
		"Axial Mash", // AxialMash
		"AxialMash",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1202),
		"Transverse Mash", // TransverseMash
		"TransverseMash",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1203),
		"Detector Element Size", // DetectorElementSize
		"DetectorElementSize",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1210),
		"Coincidence Window Width", // CoincidenceWindowWidth
		"CoincidenceWindowWidth",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1220),
		"Secondary Counts Type", // SecondaryCountsType
		"SecondaryCountsType",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1300),
		"Frame Reference Time", // FrameReferenceTime
		"FrameReferenceTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1310),
		"Primary (Prompts) Counts Accumulated", // PrimaryPromptsCountsAccumulated
		"PrimaryPromptsCountsAccumulated",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1311),
		"Secondary Counts Accumulated", // SecondaryCountsAccumulated
		"SecondaryCountsAccumulated",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1320),
		"Slice Sensitivity Factor", // SliceSensitivityFactor
		"SliceSensitivityFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1321),
		"Decay Factor", // DecayFactor
		"DecayFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1322),
		"Dose Calibration Factor", // DoseCalibrationFactor
		"DoseCalibrationFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1323),
		"Scatter Fraction Factor", // ScatterFractionFactor
		"ScatterFractionFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1324),
		"Dead Time Factor", // DeadTimeFactor
		"DeadTimeFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1330),
		"Image Index", // ImageIndex
		"ImageIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1400),
		"Counts Included", // CountsIncluded
		"CountsIncluded",
		vm.VM1_n,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0054, 0x1401),
		"Dead Time Correction Flag", // DeadTimeCorrectionFlag
		"DeadTimeCorrectionFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3000),
		"Histogram Sequence", // HistogramSequence
		"HistogramSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3002),
		"Histogram Number of Bins", // HistogramNumberOfBins
		"HistogramNumberOfBins",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3004),
		"Histogram First Bin Value", // HistogramFirstBinValue
		"HistogramFirstBinValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3006),
		"Histogram Last Bin Value", // HistogramLastBinValue
		"HistogramLastBinValue",
		vm.VM1,
		false,
		vr.US, vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3008),
		"Histogram Bin Width", // HistogramBinWidth
		"HistogramBinWidth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3010),
		"Histogram Explanation", // HistogramExplanation
		"HistogramExplanation",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0060, 0x3020),
		"Histogram Data", // HistogramData
		"HistogramData",
		vm.VM1_n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0001),
		"Segmentation Type", // SegmentationType
		"SegmentationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0002),
		"Segment Sequence", // SegmentSequence
		"SegmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0003),
		"Segmented Property Category Code Sequence", // SegmentedPropertyCategoryCodeSequence
		"SegmentedPropertyCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0004),
		"Segment Number", // SegmentNumber
		"SegmentNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0005),
		"Segment Label", // SegmentLabel
		"SegmentLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0006),
		"Segment Description", // SegmentDescription
		"SegmentDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0007),
		"Segmentation Algorithm Identification Sequence", // SegmentationAlgorithmIdentificationSequence
		"SegmentationAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0008),
		"Segment Algorithm Type", // SegmentAlgorithmType
		"SegmentAlgorithmType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0009),
		"Segment Algorithm Name", // SegmentAlgorithmName
		"SegmentAlgorithmName",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000A),
		"Segment Identification Sequence", // SegmentIdentificationSequence
		"SegmentIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000B),
		"Referenced Segment Number", // ReferencedSegmentNumber
		"ReferencedSegmentNumber",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000C),
		"Recommended Display Grayscale Value", // RecommendedDisplayGrayscaleValue
		"RecommendedDisplayGrayscaleValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000D),
		"Recommended Display CIELab Value", // RecommendedDisplayCIELabValue
		"RecommendedDisplayCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000E),
		"Maximum Fractional Value", // MaximumFractionalValue
		"MaximumFractionalValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x000F),
		"Segmented Property Type Code Sequence", // SegmentedPropertyTypeCodeSequence
		"SegmentedPropertyTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0010),
		"Segmentation Fractional Type", // SegmentationFractionalType
		"SegmentationFractionalType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0011),
		"Segmented Property Type Modifier Code Sequence", // SegmentedPropertyTypeModifierCodeSequence
		"SegmentedPropertyTypeModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0012),
		"Used Segments Sequence", // UsedSegmentsSequence
		"UsedSegmentsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0013),
		"Segments Overlap", // SegmentsOverlap
		"SegmentsOverlap",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0020),
		"Tracking ID", // TrackingID
		"TrackingID",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0062, 0x0021),
		"Tracking UID", // TrackingUID
		"TrackingUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0002),
		"Deformable Registration Sequence", // DeformableRegistrationSequence
		"DeformableRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0003),
		"Source Frame of Reference UID", // SourceFrameOfReferenceUID
		"SourceFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0005),
		"Deformable Registration Grid Sequence", // DeformableRegistrationGridSequence
		"DeformableRegistrationGridSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0007),
		"Grid Dimensions", // GridDimensions
		"GridDimensions",
		vm.VM3,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0008),
		"Grid Resolution", // GridResolution
		"GridResolution",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0009),
		"Vector Grid Data", // VectorGridData
		"VectorGridData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x000F),
		"Pre Deformation Matrix Registration Sequence", // PreDeformationMatrixRegistrationSequence
		"PreDeformationMatrixRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0064, 0x0010),
		"Post Deformation Matrix Registration Sequence", // PostDeformationMatrixRegistrationSequence
		"PostDeformationMatrixRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0001),
		"Number of Surfaces", // NumberOfSurfaces
		"NumberOfSurfaces",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0002),
		"Surface Sequence", // SurfaceSequence
		"SurfaceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0003),
		"Surface Number", // SurfaceNumber
		"SurfaceNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0004),
		"Surface Comments", // SurfaceComments
		"SurfaceComments",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0005),
		"Surface Offset", // SurfaceOffset
		"SurfaceOffset",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0009),
		"Surface Processing", // SurfaceProcessing
		"SurfaceProcessing",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000A),
		"Surface Processing Ratio", // SurfaceProcessingRatio
		"SurfaceProcessingRatio",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000B),
		"Surface Processing Description", // SurfaceProcessingDescription
		"SurfaceProcessingDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000C),
		"Recommended Presentation Opacity", // RecommendedPresentationOpacity
		"RecommendedPresentationOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000D),
		"Recommended Presentation Type", // RecommendedPresentationType
		"RecommendedPresentationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x000E),
		"Finite Volume", // FiniteVolume
		"FiniteVolume",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0010),
		"Manifold", // Manifold
		"Manifold",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0011),
		"Surface Points Sequence", // SurfacePointsSequence
		"SurfacePointsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0012),
		"Surface Points Normals Sequence", // SurfacePointsNormalsSequence
		"SurfacePointsNormalsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0013),
		"Surface Mesh Primitives Sequence", // SurfaceMeshPrimitivesSequence
		"SurfaceMeshPrimitivesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0015),
		"Number of Surface Points", // NumberOfSurfacePoints
		"NumberOfSurfacePoints",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0016),
		"Point Coordinates Data", // PointCoordinatesData
		"PointCoordinatesData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0017),
		"Point Position Accuracy", // PointPositionAccuracy
		"PointPositionAccuracy",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0018),
		"Mean Point Distance", // MeanPointDistance
		"MeanPointDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0019),
		"Maximum Point Distance", // MaximumPointDistance
		"MaximumPointDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001A),
		"Points Bounding Box Coordinates", // PointsBoundingBoxCoordinates
		"PointsBoundingBoxCoordinates",
		vm.VM6,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001B),
		"Axis of Rotation", // AxisOfRotation
		"AxisOfRotation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001C),
		"Center of Rotation", // CenterOfRotation
		"CenterOfRotation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001E),
		"Number of Vectors", // NumberOfVectors
		"NumberOfVectors",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x001F),
		"Vector Dimensionality", // VectorDimensionality
		"VectorDimensionality",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0020),
		"Vector Accuracy", // VectorAccuracy
		"VectorAccuracy",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0021),
		"Vector Coordinate Data", // VectorCoordinateData
		"VectorCoordinateData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0022),
		"Double Point Coordinates Data", // DoublePointCoordinatesData
		"DoublePointCoordinatesData",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0023),
		"Triangle Point Index List", // TrianglePointIndexList
		"TrianglePointIndexList",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0024),
		"Edge Point Index List", // EdgePointIndexList
		"EdgePointIndexList",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0025),
		"Vertex Point Index List", // VertexPointIndexList
		"VertexPointIndexList",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0026),
		"Triangle Strip Sequence", // TriangleStripSequence
		"TriangleStripSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0027),
		"Triangle Fan Sequence", // TriangleFanSequence
		"TriangleFanSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0028),
		"Line Sequence", // LineSequence
		"LineSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0029),
		"Primitive Point Index List", // PrimitivePointIndexList
		"PrimitivePointIndexList",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002A),
		"Surface Count", // SurfaceCount
		"SurfaceCount",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002B),
		"Referenced Surface Sequence", // ReferencedSurfaceSequence
		"ReferencedSurfaceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002C),
		"Referenced Surface Number", // ReferencedSurfaceNumber
		"ReferencedSurfaceNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002D),
		"Segment Surface Generation Algorithm Identification Sequence", // SegmentSurfaceGenerationAlgorithmIdentificationSequence
		"SegmentSurfaceGenerationAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002E),
		"Segment Surface Source Instance Sequence", // SegmentSurfaceSourceInstanceSequence
		"SegmentSurfaceSourceInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x002F),
		"Algorithm Family Code Sequence", // AlgorithmFamilyCodeSequence
		"AlgorithmFamilyCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0030),
		"Algorithm Name Code Sequence", // AlgorithmNameCodeSequence
		"AlgorithmNameCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0031),
		"Algorithm Version", // AlgorithmVersion
		"AlgorithmVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0032),
		"Algorithm Parameters", // AlgorithmParameters
		"AlgorithmParameters",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0034),
		"Facet Sequence", // FacetSequence
		"FacetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0035),
		"Surface Processing Algorithm Identification Sequence", // SurfaceProcessingAlgorithmIdentificationSequence
		"SurfaceProcessingAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0036),
		"Algorithm Name", // AlgorithmName
		"AlgorithmName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0037),
		"Recommended Point Radius", // RecommendedPointRadius
		"RecommendedPointRadius",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0038),
		"Recommended Line Thickness", // RecommendedLineThickness
		"RecommendedLineThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0040),
		"Long Primitive Point Index List", // LongPrimitivePointIndexList
		"LongPrimitivePointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0041),
		"Long Triangle Point Index List", // LongTrianglePointIndexList
		"LongTrianglePointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0042),
		"Long Edge Point Index List", // LongEdgePointIndexList
		"LongEdgePointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0043),
		"Long Vertex Point Index List", // LongVertexPointIndexList
		"LongVertexPointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0101),
		"Track Set Sequence", // TrackSetSequence
		"TrackSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0102),
		"Track Sequence", // TrackSequence
		"TrackSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0103),
		"Recommended Display CIELab Value List", // RecommendedDisplayCIELabValueList
		"RecommendedDisplayCIELabValueList",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0104),
		"Tracking Algorithm Identification Sequence", // TrackingAlgorithmIdentificationSequence
		"TrackingAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0105),
		"Track Set Number", // TrackSetNumber
		"TrackSetNumber",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0106),
		"Track Set Label", // TrackSetLabel
		"TrackSetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0107),
		"Track Set Description", // TrackSetDescription
		"TrackSetDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0108),
		"Track Set Anatomical Type Code Sequence", // TrackSetAnatomicalTypeCodeSequence
		"TrackSetAnatomicalTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0121),
		"Measurements Sequence", // MeasurementsSequence
		"MeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0124),
		"Track Set Statistics Sequence", // TrackSetStatisticsSequence
		"TrackSetStatisticsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0125),
		"Floating Point Values", // FloatingPointValues
		"FloatingPointValues",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0129),
		"Track Point Index List", // TrackPointIndexList
		"TrackPointIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0130),
		"Track Statistics Sequence", // TrackStatisticsSequence
		"TrackStatisticsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0132),
		"Measurement Values Sequence", // MeasurementValuesSequence
		"MeasurementValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0133),
		"Diffusion Acquisition Code Sequence", // DiffusionAcquisitionCodeSequence
		"DiffusionAcquisitionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0066, 0x0134),
		"Diffusion Model Code Sequence", // DiffusionModelCodeSequence
		"DiffusionModelCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6210),
		"Implant Size", // ImplantSize
		"ImplantSize",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6221),
		"Implant Template Version", // ImplantTemplateVersion
		"ImplantTemplateVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6222),
		"Replaced Implant Template Sequence", // ReplacedImplantTemplateSequence
		"ReplacedImplantTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6223),
		"Implant Type", // ImplantType
		"ImplantType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6224),
		"Derivation Implant Template Sequence", // DerivationImplantTemplateSequence
		"DerivationImplantTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6225),
		"Original Implant Template Sequence", // OriginalImplantTemplateSequence
		"OriginalImplantTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6226),
		"Effective DateTime", // EffectiveDateTime
		"EffectiveDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6230),
		"Implant Target Anatomy Sequence", // ImplantTargetAnatomySequence
		"ImplantTargetAnatomySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6260),
		"Information From Manufacturer Sequence", // InformationFromManufacturerSequence
		"InformationFromManufacturerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6265),
		"Notification From Manufacturer Sequence", // NotificationFromManufacturerSequence
		"NotificationFromManufacturerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6270),
		"Information Issue DateTime", // InformationIssueDateTime
		"InformationIssueDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6280),
		"Information Summary", // InformationSummary
		"InformationSummary",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62A0),
		"Implant Regulatory Disapproval Code Sequence", // ImplantRegulatoryDisapprovalCodeSequence
		"ImplantRegulatoryDisapprovalCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62A5),
		"Overall Template Spatial Tolerance", // OverallTemplateSpatialTolerance
		"OverallTemplateSpatialTolerance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62C0),
		"HPGL Document Sequence", // HPGLDocumentSequence
		"HPGLDocumentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62D0),
		"HPGL Document ID", // HPGLDocumentID
		"HPGLDocumentID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62D5),
		"HPGL Document Label", // HPGLDocumentLabel
		"HPGLDocumentLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62E0),
		"View Orientation Code Sequence", // ViewOrientationCodeSequence
		"ViewOrientationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62F0),
		"View Orientation Modifier Code Sequence", // ViewOrientationModifierCodeSequence
		"ViewOrientationModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x62F2),
		"HPGL Document Scaling", // HPGLDocumentScaling
		"HPGLDocumentScaling",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6300),
		"HPGL Document", // HPGLDocument
		"HPGLDocument",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6310),
		"HPGL Contour Pen Number", // HPGLContourPenNumber
		"HPGLContourPenNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6320),
		"HPGL Pen Sequence", // HPGLPenSequence
		"HPGLPenSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6330),
		"HPGL Pen Number", // HPGLPenNumber
		"HPGLPenNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6340),
		"HPGL Pen Label", // HPGLPenLabel
		"HPGLPenLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6345),
		"HPGL Pen Description", // HPGLPenDescription
		"HPGLPenDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6346),
		"Recommended Rotation Point", // RecommendedRotationPoint
		"RecommendedRotationPoint",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6347),
		"Bounding Rectangle", // BoundingRectangle
		"BoundingRectangle",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6350),
		"Implant Template 3D Model Surface Number", // ImplantTemplate3DModelSurfaceNumber
		"ImplantTemplate3DModelSurfaceNumber",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6360),
		"Surface Model Description Sequence", // SurfaceModelDescriptionSequence
		"SurfaceModelDescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6380),
		"Surface Model Label", // SurfaceModelLabel
		"SurfaceModelLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6390),
		"Surface Model Scaling Factor", // SurfaceModelScalingFactor
		"SurfaceModelScalingFactor",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63A0),
		"Materials Code Sequence", // MaterialsCodeSequence
		"MaterialsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63A4),
		"Coating Materials Code Sequence", // CoatingMaterialsCodeSequence
		"CoatingMaterialsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63A8),
		"Implant Type Code Sequence", // ImplantTypeCodeSequence
		"ImplantTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63AC),
		"Fixation Method Code Sequence", // FixationMethodCodeSequence
		"FixationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63B0),
		"Mating Feature Sets Sequence", // MatingFeatureSetsSequence
		"MatingFeatureSetsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63C0),
		"Mating Feature Set ID", // MatingFeatureSetID
		"MatingFeatureSetID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63D0),
		"Mating Feature Set Label", // MatingFeatureSetLabel
		"MatingFeatureSetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63E0),
		"Mating Feature Sequence", // MatingFeatureSequence
		"MatingFeatureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x63F0),
		"Mating Feature ID", // MatingFeatureID
		"MatingFeatureID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6400),
		"Mating Feature Degree of Freedom Sequence", // MatingFeatureDegreeOfFreedomSequence
		"MatingFeatureDegreeOfFreedomSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6410),
		"Degree of Freedom ID", // DegreeOfFreedomID
		"DegreeOfFreedomID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6420),
		"Degree of Freedom Type", // DegreeOfFreedomType
		"DegreeOfFreedomType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6430),
		"2D Mating Feature Coordinates Sequence", // TwoDMatingFeatureCoordinatesSequence
		"TwoDMatingFeatureCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6440),
		"Referenced HPGL Document ID", // ReferencedHPGLDocumentID
		"ReferencedHPGLDocumentID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6450),
		"2D Mating Point", // TwoDMatingPoint
		"TwoDMatingPoint",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6460),
		"2D Mating Axes", // TwoDMatingAxes
		"TwoDMatingAxes",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6470),
		"2D Degree of Freedom Sequence", // TwoDDegreeOfFreedomSequence
		"TwoDDegreeOfFreedomSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6490),
		"3D Degree of Freedom Axis", // ThreeDDegreeOfFreedomAxis
		"ThreeDDegreeOfFreedomAxis",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x64A0),
		"Range of Freedom", // RangeOfFreedom
		"RangeOfFreedom",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x64C0),
		"3D Mating Point", // ThreeDMatingPoint
		"ThreeDMatingPoint",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x64D0),
		"3D Mating Axes", // ThreeDMatingAxes
		"ThreeDMatingAxes",
		vm.MustParse("9"),
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x64F0),
		"2D Degree of Freedom Axis", // TwoDDegreeOfFreedomAxis
		"TwoDDegreeOfFreedomAxis",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6500),
		"Planning Landmark Point Sequence", // PlanningLandmarkPointSequence
		"PlanningLandmarkPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6510),
		"Planning Landmark Line Sequence", // PlanningLandmarkLineSequence
		"PlanningLandmarkLineSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6520),
		"Planning Landmark Plane Sequence", // PlanningLandmarkPlaneSequence
		"PlanningLandmarkPlaneSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6530),
		"Planning Landmark ID", // PlanningLandmarkID
		"PlanningLandmarkID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6540),
		"Planning Landmark Description", // PlanningLandmarkDescription
		"PlanningLandmarkDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6545),
		"Planning Landmark Identification Code Sequence", // PlanningLandmarkIdentificationCodeSequence
		"PlanningLandmarkIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6550),
		"2D Point Coordinates Sequence", // TwoDPointCoordinatesSequence
		"TwoDPointCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6560),
		"2D Point Coordinates", // TwoDPointCoordinates
		"TwoDPointCoordinates",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6590),
		"3D Point Coordinates", // ThreeDPointCoordinates
		"ThreeDPointCoordinates",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65A0),
		"2D Line Coordinates Sequence", // TwoDLineCoordinatesSequence
		"TwoDLineCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65B0),
		"2D Line Coordinates", // TwoDLineCoordinates
		"TwoDLineCoordinates",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65D0),
		"3D Line Coordinates", // ThreeDLineCoordinates
		"ThreeDLineCoordinates",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65E0),
		"2D Plane Coordinates Sequence", // TwoDPlaneCoordinatesSequence
		"TwoDPlaneCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x65F0),
		"2D Plane Intersection", // TwoDPlaneIntersection
		"TwoDPlaneIntersection",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6610),
		"3D Plane Origin", // ThreeDPlaneOrigin
		"ThreeDPlaneOrigin",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x6620),
		"3D Plane Normal", // ThreeDPlaneNormal
		"ThreeDPlaneNormal",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7001),
		"Model Modification", // ModelModification
		"ModelModification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7002),
		"Model Mirroring", // ModelMirroring
		"ModelMirroring",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7003),
		"Model Usage Code Sequence", // ModelUsageCodeSequence
		"ModelUsageCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7004),
		"Model Group UID", // ModelGroupUID
		"ModelGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0068, 0x7005),
		"Relative URI Reference Within Encapsulated Document", // RelativeURIReferenceWithinEncapsulatedDocument
		"RelativeURIReferenceWithinEncapsulatedDocument",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0001),
		"Annotation Coordinate Type", // AnnotationCoordinateType
		"AnnotationCoordinateType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0002),
		"Annotation Group Sequence", // AnnotationGroupSequence
		"AnnotationGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0003),
		"Annotation Group UID", // AnnotationGroupUID
		"AnnotationGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0005),
		"Annotation Group Label", // AnnotationGroupLabel
		"AnnotationGroupLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0006),
		"Annotation Group Description", // AnnotationGroupDescription
		"AnnotationGroupDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0007),
		"Annotation Group Generation Type", // AnnotationGroupGenerationType
		"AnnotationGroupGenerationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0008),
		"Annotation Group Algorithm Identification Sequence", // AnnotationGroupAlgorithmIdentificationSequence
		"AnnotationGroupAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0009),
		"Annotation Property Category Code Sequence", // AnnotationPropertyCategoryCodeSequence
		"AnnotationPropertyCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000A),
		"Annotation Property Type Code Sequence", // AnnotationPropertyTypeCodeSequence
		"AnnotationPropertyTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000B),
		"Annotation Property Type Modifier Code Sequence", // AnnotationPropertyTypeModifierCodeSequence
		"AnnotationPropertyTypeModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000C),
		"Number of Annotations", // NumberOfAnnotations
		"NumberOfAnnotations",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000D),
		"Annotation Applies to All Optical Paths", // AnnotationAppliesToAllOpticalPaths
		"AnnotationAppliesToAllOpticalPaths",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000E),
		"Referenced Optical Path Identifier", // ReferencedOpticalPathIdentifier
		"ReferencedOpticalPathIdentifier",
		vm.VM1_n,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x000F),
		"Annotation Applies to All Z Planes", // AnnotationAppliesToAllZPlanes
		"AnnotationAppliesToAllZPlanes",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0010),
		"Common Z Coordinate Value", // CommonZCoordinateValue
		"CommonZCoordinateValue",
		vm.VM1_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x006A, 0x0011),
		"Annotation Index List", // AnnotationIndexList
		"AnnotationIndexList",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0001),
		"Graphic Annotation Sequence", // GraphicAnnotationSequence
		"GraphicAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0002),
		"Graphic Layer", // GraphicLayer
		"GraphicLayer",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0003),
		"Bounding Box Annotation Units", // BoundingBoxAnnotationUnits
		"BoundingBoxAnnotationUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0004),
		"Anchor Point Annotation Units", // AnchorPointAnnotationUnits
		"AnchorPointAnnotationUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0005),
		"Graphic Annotation Units", // GraphicAnnotationUnits
		"GraphicAnnotationUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0006),
		"Unformatted Text Value", // UnformattedTextValue
		"UnformattedTextValue",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0008),
		"Text Object Sequence", // TextObjectSequence
		"TextObjectSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0009),
		"Graphic Object Sequence", // GraphicObjectSequence
		"GraphicObjectSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0010),
		"Bounding Box Top Left Hand Corner", // BoundingBoxTopLeftHandCorner
		"BoundingBoxTopLeftHandCorner",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0011),
		"Bounding Box Bottom Right Hand Corner", // BoundingBoxBottomRightHandCorner
		"BoundingBoxBottomRightHandCorner",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0012),
		"Bounding Box Text Horizontal Justification", // BoundingBoxTextHorizontalJustification
		"BoundingBoxTextHorizontalJustification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0014),
		"Anchor Point", // AnchorPoint
		"AnchorPoint",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0015),
		"Anchor Point Visibility", // AnchorPointVisibility
		"AnchorPointVisibility",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0020),
		"Graphic Dimensions", // GraphicDimensions
		"GraphicDimensions",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0021),
		"Number of Graphic Points", // NumberOfGraphicPoints
		"NumberOfGraphicPoints",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0022),
		"Graphic Data", // GraphicData
		"GraphicData",
		vm.VM2_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0023),
		"Graphic Type", // GraphicType
		"GraphicType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0024),
		"Graphic Filled", // GraphicFilled
		"GraphicFilled",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0040),
		"Image Rotation (Retired)", // ImageRotationRetired
		"ImageRotationRetired",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0041),
		"Image Horizontal Flip", // ImageHorizontalFlip
		"ImageHorizontalFlip",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0042),
		"Image Rotation", // ImageRotation
		"ImageRotation",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0050),
		"Displayed Area Top Left Hand Corner (Trial)", // DisplayedAreaTopLeftHandCornerTrial
		"DisplayedAreaTopLeftHandCornerTrial",
		vm.VM2,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0051),
		"Displayed Area Bottom Right Hand Corner (Trial)", // DisplayedAreaBottomRightHandCornerTrial
		"DisplayedAreaBottomRightHandCornerTrial",
		vm.VM2,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0052),
		"Displayed Area Top Left Hand Corner", // DisplayedAreaTopLeftHandCorner
		"DisplayedAreaTopLeftHandCorner",
		vm.VM2,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0053),
		"Displayed Area Bottom Right Hand Corner", // DisplayedAreaBottomRightHandCorner
		"DisplayedAreaBottomRightHandCorner",
		vm.VM2,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x005A),
		"Displayed Area Selection Sequence", // DisplayedAreaSelectionSequence
		"DisplayedAreaSelectionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0060),
		"Graphic Layer Sequence", // GraphicLayerSequence
		"GraphicLayerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0062),
		"Graphic Layer Order", // GraphicLayerOrder
		"GraphicLayerOrder",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0066),
		"Graphic Layer Recommended Display Grayscale Value", // GraphicLayerRecommendedDisplayGrayscaleValue
		"GraphicLayerRecommendedDisplayGrayscaleValue",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0067),
		"Graphic Layer Recommended Display RGB Value", // GraphicLayerRecommendedDisplayRGBValue
		"GraphicLayerRecommendedDisplayRGBValue",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0068),
		"Graphic Layer Description", // GraphicLayerDescription
		"GraphicLayerDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0080),
		"Content Label", // ContentLabel
		"ContentLabel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0081),
		"Content Description", // ContentDescription
		"ContentDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0082),
		"Presentation Creation Date", // PresentationCreationDate
		"PresentationCreationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0083),
		"Presentation Creation Time", // PresentationCreationTime
		"PresentationCreationTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0084),
		"Content Creator's Name", // ContentCreatorName
		"ContentCreatorName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0086),
		"Content Creator's Identification Code Sequence", // ContentCreatorIdentificationCodeSequence
		"ContentCreatorIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0087),
		"Alternate Content Description Sequence", // AlternateContentDescriptionSequence
		"AlternateContentDescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0100),
		"Presentation Size Mode", // PresentationSizeMode
		"PresentationSizeMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0101),
		"Presentation Pixel Spacing", // PresentationPixelSpacing
		"PresentationPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0102),
		"Presentation Pixel Aspect Ratio", // PresentationPixelAspectRatio
		"PresentationPixelAspectRatio",
		vm.VM2,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0103),
		"Presentation Pixel Magnification Ratio", // PresentationPixelMagnificationRatio
		"PresentationPixelMagnificationRatio",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0207),
		"Graphic Group Label", // GraphicGroupLabel
		"GraphicGroupLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0208),
		"Graphic Group Description", // GraphicGroupDescription
		"GraphicGroupDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0209),
		"Compound Graphic Sequence", // CompoundGraphicSequence
		"CompoundGraphicSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0226),
		"Compound Graphic Instance ID", // CompoundGraphicInstanceID
		"CompoundGraphicInstanceID",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0227),
		"Font Name", // FontName
		"FontName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0228),
		"Font Name Type", // FontNameType
		"FontNameType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0229),
		"CSS Font Name", // CSSFontName
		"CSSFontName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0230),
		"Rotation Angle", // RotationAngle
		"RotationAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0231),
		"Text Style Sequence", // TextStyleSequence
		"TextStyleSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0232),
		"Line Style Sequence", // LineStyleSequence
		"LineStyleSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0233),
		"Fill Style Sequence", // FillStyleSequence
		"FillStyleSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0234),
		"Graphic Group Sequence", // GraphicGroupSequence
		"GraphicGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0241),
		"Text Color CIELab Value", // TextColorCIELabValue
		"TextColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0242),
		"Horizontal Alignment", // HorizontalAlignment
		"HorizontalAlignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0243),
		"Vertical Alignment", // VerticalAlignment
		"VerticalAlignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0244),
		"Shadow Style", // ShadowStyle
		"ShadowStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0245),
		"Shadow Offset X", // ShadowOffsetX
		"ShadowOffsetX",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0246),
		"Shadow Offset Y", // ShadowOffsetY
		"ShadowOffsetY",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0247),
		"Shadow Color CIELab Value", // ShadowColorCIELabValue
		"ShadowColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0248),
		"Underlined", // Underlined
		"Underlined",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0249),
		"Bold", // Bold
		"Bold",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0250),
		"Italic", // Italic
		"Italic",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0251),
		"Pattern On Color CIELab Value", // PatternOnColorCIELabValue
		"PatternOnColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0252),
		"Pattern Off Color CIELab Value", // PatternOffColorCIELabValue
		"PatternOffColorCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0253),
		"Line Thickness", // LineThickness
		"LineThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0254),
		"Line Dashing Style", // LineDashingStyle
		"LineDashingStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0255),
		"Line Pattern", // LinePattern
		"LinePattern",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0256),
		"Fill Pattern", // FillPattern
		"FillPattern",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0257),
		"Fill Mode", // FillMode
		"FillMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0258),
		"Shadow Opacity", // ShadowOpacity
		"ShadowOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0261),
		"Gap Length", // GapLength
		"GapLength",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0262),
		"Diameter of Visibility", // DiameterOfVisibility
		"DiameterOfVisibility",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0273),
		"Rotation Point", // RotationPoint
		"RotationPoint",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0274),
		"Tick Alignment", // TickAlignment
		"TickAlignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0278),
		"Show Tick Label", // ShowTickLabel
		"ShowTickLabel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0279),
		"Tick Label Alignment", // TickLabelAlignment
		"TickLabelAlignment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0282),
		"Compound Graphic Units", // CompoundGraphicUnits
		"CompoundGraphicUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0284),
		"Pattern On Opacity", // PatternOnOpacity
		"PatternOnOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0285),
		"Pattern Off Opacity", // PatternOffOpacity
		"PatternOffOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0287),
		"Major Ticks Sequence", // MajorTicksSequence
		"MajorTicksSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0288),
		"Tick Position", // TickPosition
		"TickPosition",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0289),
		"Tick Label", // TickLabel
		"TickLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0294),
		"Compound Graphic Type", // CompoundGraphicType
		"CompoundGraphicType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0295),
		"Graphic Group ID", // GraphicGroupID
		"GraphicGroupID",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0306),
		"Shape Type", // ShapeType
		"ShapeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0308),
		"Registration Sequence", // RegistrationSequence
		"RegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0309),
		"Matrix Registration Sequence", // MatrixRegistrationSequence
		"MatrixRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030A),
		"Matrix Sequence", // MatrixSequence
		"MatrixSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030B),
		"Frame of Reference to Displayed Coordinate System Transformation Matrix", // FrameOfReferenceToDisplayedCoordinateSystemTransformationMatrix
		"FrameOfReferenceToDisplayedCoordinateSystemTransformationMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030C),
		"Frame of Reference Transformation Matrix Type", // FrameOfReferenceTransformationMatrixType
		"FrameOfReferenceTransformationMatrixType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030D),
		"Registration Type Code Sequence", // RegistrationTypeCodeSequence
		"RegistrationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x030F),
		"Fiducial Description", // FiducialDescription
		"FiducialDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0310),
		"Fiducial Identifier", // FiducialIdentifier
		"FiducialIdentifier",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0311),
		"Fiducial Identifier Code Sequence", // FiducialIdentifierCodeSequence
		"FiducialIdentifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0312),
		"Contour Uncertainty Radius", // ContourUncertaintyRadius
		"ContourUncertaintyRadius",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0314),
		"Used Fiducials Sequence", // UsedFiducialsSequence
		"UsedFiducialsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0315),
		"Used RT Structure Set ROI Sequence", // UsedRTStructureSetROISequence
		"UsedRTStructureSetROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0318),
		"Graphic Coordinates Data Sequence", // GraphicCoordinatesDataSequence
		"GraphicCoordinatesDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031A),
		"Fiducial UID", // FiducialUID
		"FiducialUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031B),
		"Referenced Fiducial UID", // ReferencedFiducialUID
		"ReferencedFiducialUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031C),
		"Fiducial Set Sequence", // FiducialSetSequence
		"FiducialSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031E),
		"Fiducial Sequence", // FiducialSequence
		"FiducialSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x031F),
		"Fiducials Property Category Code Sequence", // FiducialsPropertyCategoryCodeSequence
		"FiducialsPropertyCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0401),
		"Graphic Layer Recommended Display CIELab Value", // GraphicLayerRecommendedDisplayCIELabValue
		"GraphicLayerRecommendedDisplayCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0402),
		"Blending Sequence", // BlendingSequence
		"BlendingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0403),
		"Relative Opacity", // RelativeOpacity
		"RelativeOpacity",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0404),
		"Referenced Spatial Registration Sequence", // ReferencedSpatialRegistrationSequence
		"ReferencedSpatialRegistrationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x0405),
		"Blending Position", // BlendingPosition
		"BlendingPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1101),
		"Presentation Display Collection UID", // PresentationDisplayCollectionUID
		"PresentationDisplayCollectionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1102),
		"Presentation Sequence Collection UID", // PresentationSequenceCollectionUID
		"PresentationSequenceCollectionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1103),
		"Presentation Sequence Position Index", // PresentationSequencePositionIndex
		"PresentationSequencePositionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1104),
		"Rendered Image Reference Sequence", // RenderedImageReferenceSequence
		"RenderedImageReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1201),
		"Volumetric Presentation State Input Sequence", // VolumetricPresentationStateInputSequence
		"VolumetricPresentationStateInputSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1202),
		"Presentation Input Type", // PresentationInputType
		"PresentationInputType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1203),
		"Input Sequence Position Index", // InputSequencePositionIndex
		"InputSequencePositionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1204),
		"Crop", // Crop
		"Crop",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1205),
		"Cropping Specification Index", // CroppingSpecificationIndex
		"CroppingSpecificationIndex",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1206),
		"Compositing Method", // CompositingMethod
		"CompositingMethod",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1207),
		"Volumetric Presentation Input Number", // VolumetricPresentationInputNumber
		"VolumetricPresentationInputNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1208),
		"Image Volume Geometry", // ImageVolumeGeometry
		"ImageVolumeGeometry",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1209),
		"Volumetric Presentation Input Set UID", // VolumetricPresentationInputSetUID
		"VolumetricPresentationInputSetUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x120A),
		"Volumetric Presentation Input Set Sequence", // VolumetricPresentationInputSetSequence
		"VolumetricPresentationInputSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x120B),
		"Global Crop", // GlobalCrop
		"GlobalCrop",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x120C),
		"Global Cropping Specification Index", // GlobalCroppingSpecificationIndex
		"GlobalCroppingSpecificationIndex",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x120D),
		"Rendering Method", // RenderingMethod
		"RenderingMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1301),
		"Volume Cropping Sequence", // VolumeCroppingSequence
		"VolumeCroppingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1302),
		"Volume Cropping Method", // VolumeCroppingMethod
		"VolumeCroppingMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1303),
		"Bounding Box Crop", // BoundingBoxCrop
		"BoundingBoxCrop",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1304),
		"Oblique Cropping Plane Sequence", // ObliqueCroppingPlaneSequence
		"ObliqueCroppingPlaneSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1305),
		"Plane", // Plane
		"Plane",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1306),
		"Plane Normal", // PlaneNormal
		"PlaneNormal",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1309),
		"Cropping Specification Number", // CroppingSpecificationNumber
		"CroppingSpecificationNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1501),
		"Multi-Planar Reconstruction Style", // MultiPlanarReconstructionStyle
		"MultiPlanarReconstructionStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1502),
		"MPR Thickness Type", // MPRThicknessType
		"MPRThicknessType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1503),
		"MPR Slab Thickness", // MPRSlabThickness
		"MPRSlabThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1505),
		"MPR Top Left Hand Corner", // MPRTopLeftHandCorner
		"MPRTopLeftHandCorner",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1507),
		"MPR View Width Direction", // MPRViewWidthDirection
		"MPRViewWidthDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1508),
		"MPR View Width", // MPRViewWidth
		"MPRViewWidth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x150C),
		"Number of Volumetric Curve Points", // NumberOfVolumetricCurvePoints
		"NumberOfVolumetricCurvePoints",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x150D),
		"Volumetric Curve Points", // VolumetricCurvePoints
		"VolumetricCurvePoints",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1511),
		"MPR View Height Direction", // MPRViewHeightDirection
		"MPRViewHeightDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1512),
		"MPR View Height", // MPRViewHeight
		"MPRViewHeight",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1602),
		"Render Projection", // RenderProjection
		"RenderProjection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1603),
		"Viewpoint Position", // ViewpointPosition
		"ViewpointPosition",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1604),
		"Viewpoint LookAt Point", // ViewpointLookAtPoint
		"ViewpointLookAtPoint",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1605),
		"Viewpoint Up Direction", // ViewpointUpDirection
		"ViewpointUpDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1606),
		"Render Field of View", // RenderFieldOfView
		"RenderFieldOfView",
		vm.VM6,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1607),
		"Sampling Step Size", // SamplingStepSize
		"SamplingStepSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1701),
		"Shading Style", // ShadingStyle
		"ShadingStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1702),
		"Ambient Reflection Intensity", // AmbientReflectionIntensity
		"AmbientReflectionIntensity",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1703),
		"Light Direction", // LightDirection
		"LightDirection",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1704),
		"Diffuse Reflection Intensity", // DiffuseReflectionIntensity
		"DiffuseReflectionIntensity",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1705),
		"Specular Reflection Intensity", // SpecularReflectionIntensity
		"SpecularReflectionIntensity",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1706),
		"Shininess", // Shininess
		"Shininess",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1801),
		"Presentation State Classification Component Sequence", // PresentationStateClassificationComponentSequence
		"PresentationStateClassificationComponentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1802),
		"Component Type", // ComponentType
		"ComponentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1803),
		"Component Input Sequence", // ComponentInputSequence
		"ComponentInputSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1804),
		"Volumetric Presentation Input Index", // VolumetricPresentationInputIndex
		"VolumetricPresentationInputIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1805),
		"Presentation State Compositor Component Sequence", // PresentationStateCompositorComponentSequence
		"PresentationStateCompositorComponentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1806),
		"Weighting Transfer Function Sequence", // WeightingTransferFunctionSequence
		"WeightingTransferFunctionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1807),
		"Weighting Lookup Table Descriptor", // WeightingLookupTableDescriptor
		"WeightingLookupTableDescriptor",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1808),
		"Weighting Lookup Table Data", // WeightingLookupTableData
		"WeightingLookupTableData",
		vm.VM1,
		true,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1901),
		"Volumetric Annotation Sequence", // VolumetricAnnotationSequence
		"VolumetricAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1903),
		"Referenced Structured Context Sequence", // ReferencedStructuredContextSequence
		"ReferencedStructuredContextSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1904),
		"Referenced Content Item", // ReferencedContentItem
		"ReferencedContentItem",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1905),
		"Volumetric Presentation Input Annotation Sequence", // VolumetricPresentationInputAnnotationSequence
		"VolumetricPresentationInputAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1907),
		"Annotation Clipping", // AnnotationClipping
		"AnnotationClipping",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A01),
		"Presentation Animation Style", // PresentationAnimationStyle
		"PresentationAnimationStyle",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A03),
		"Recommended Animation Rate", // RecommendedAnimationRate
		"RecommendedAnimationRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A04),
		"Animation Curve Sequence", // AnimationCurveSequence
		"AnimationCurveSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A05),
		"Animation Step Size", // AnimationStepSize
		"AnimationStepSize",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A06),
		"Swivel Range", // SwivelRange
		"SwivelRange",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A07),
		"Volumetric Curve Up Directions", // VolumetricCurveUpDirections
		"VolumetricCurveUpDirections",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A08),
		"Volume Stream Sequence", // VolumeStreamSequence
		"VolumeStreamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1A09),
		"RGBA Transfer Function Description", // RGBATransferFunctionDescription
		"RGBATransferFunctionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B01),
		"Advanced Blending Sequence", // AdvancedBlendingSequence
		"AdvancedBlendingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B02),
		"Blending Input Number", // BlendingInputNumber
		"BlendingInputNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B03),
		"Blending Display Input Sequence", // BlendingDisplayInputSequence
		"BlendingDisplayInputSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B04),
		"Blending Display Sequence", // BlendingDisplaySequence
		"BlendingDisplaySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B06),
		"Blending Mode", // BlendingMode
		"BlendingMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B07),
		"Time Series Blending", // TimeSeriesBlending
		"TimeSeriesBlending",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B08),
		"Geometry for Display", // GeometryForDisplay
		"GeometryForDisplay",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B11),
		"Threshold Sequence", // ThresholdSequence
		"ThresholdSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B12),
		"Threshold Value Sequence", // ThresholdValueSequence
		"ThresholdValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B13),
		"Threshold Type", // ThresholdType
		"ThresholdType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0070, 0x1B14),
		"Threshold Value", // ThresholdValue
		"ThresholdValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0002),
		"Hanging Protocol Name", // HangingProtocolName
		"HangingProtocolName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0004),
		"Hanging Protocol Description", // HangingProtocolDescription
		"HangingProtocolDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0006),
		"Hanging Protocol Level", // HangingProtocolLevel
		"HangingProtocolLevel",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0008),
		"Hanging Protocol Creator", // HangingProtocolCreator
		"HangingProtocolCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x000A),
		"Hanging Protocol Creation DateTime", // HangingProtocolCreationDateTime
		"HangingProtocolCreationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x000C),
		"Hanging Protocol Definition Sequence", // HangingProtocolDefinitionSequence
		"HangingProtocolDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x000E),
		"Hanging Protocol User Identification Code Sequence", // HangingProtocolUserIdentificationCodeSequence
		"HangingProtocolUserIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0010),
		"Hanging Protocol User Group Name", // HangingProtocolUserGroupName
		"HangingProtocolUserGroupName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0012),
		"Source Hanging Protocol Sequence", // SourceHangingProtocolSequence
		"SourceHangingProtocolSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0014),
		"Number of Priors Referenced", // NumberOfPriorsReferenced
		"NumberOfPriorsReferenced",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0020),
		"Image Sets Sequence", // ImageSetsSequence
		"ImageSetsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0022),
		"Image Set Selector Sequence", // ImageSetSelectorSequence
		"ImageSetSelectorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0024),
		"Image Set Selector Usage Flag", // ImageSetSelectorUsageFlag
		"ImageSetSelectorUsageFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0026),
		"Selector Attribute", // SelectorAttribute
		"SelectorAttribute",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0028),
		"Selector Value Number", // SelectorValueNumber
		"SelectorValueNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0030),
		"Time Based Image Sets Sequence", // TimeBasedImageSetsSequence
		"TimeBasedImageSetsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0032),
		"Image Set Number", // ImageSetNumber
		"ImageSetNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0034),
		"Image Set Selector Category", // ImageSetSelectorCategory
		"ImageSetSelectorCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0038),
		"Relative Time", // RelativeTime
		"RelativeTime",
		vm.VM2,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x003A),
		"Relative Time Units", // RelativeTimeUnits
		"RelativeTimeUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x003C),
		"Abstract Prior Value", // AbstractPriorValue
		"AbstractPriorValue",
		vm.VM2,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x003E),
		"Abstract Prior Code Sequence", // AbstractPriorCodeSequence
		"AbstractPriorCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0040),
		"Image Set Label", // ImageSetLabel
		"ImageSetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0050),
		"Selector Attribute VR", // SelectorAttributeVR
		"SelectorAttributeVR",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0052),
		"Selector Sequence Pointer", // SelectorSequencePointer
		"SelectorSequencePointer",
		vm.VM1_n,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0054),
		"Selector Sequence Pointer Private Creator", // SelectorSequencePointerPrivateCreator
		"SelectorSequencePointerPrivateCreator",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0056),
		"Selector Attribute Private Creator", // SelectorAttributePrivateCreator
		"SelectorAttributePrivateCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x005E),
		"Selector AE Value", // SelectorAEValue
		"SelectorAEValue",
		vm.VM1_n,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x005F),
		"Selector AS Value", // SelectorASValue
		"SelectorASValue",
		vm.VM1_n,
		false,
		vr.AS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0060),
		"Selector AT Value", // SelectorATValue
		"SelectorATValue",
		vm.VM1_n,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0061),
		"Selector DA Value", // SelectorDAValue
		"SelectorDAValue",
		vm.VM1_n,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0062),
		"Selector CS Value", // SelectorCSValue
		"SelectorCSValue",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0063),
		"Selector DT Value", // SelectorDTValue
		"SelectorDTValue",
		vm.VM1_n,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0064),
		"Selector IS Value", // SelectorISValue
		"SelectorISValue",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0065),
		"Selector OB Value", // SelectorOBValue
		"SelectorOBValue",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0066),
		"Selector LO Value", // SelectorLOValue
		"SelectorLOValue",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0067),
		"Selector OF Value", // SelectorOFValue
		"SelectorOFValue",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0068),
		"Selector LT Value", // SelectorLTValue
		"SelectorLTValue",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0069),
		"Selector OW Value", // SelectorOWValue
		"SelectorOWValue",
		vm.VM1,
		false,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006A),
		"Selector PN Value", // SelectorPNValue
		"SelectorPNValue",
		vm.VM1_n,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006B),
		"Selector TM Value", // SelectorTMValue
		"SelectorTMValue",
		vm.VM1_n,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006C),
		"Selector SH Value", // SelectorSHValue
		"SelectorSHValue",
		vm.VM1_n,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006D),
		"Selector UN Value", // SelectorUNValue
		"SelectorUNValue",
		vm.VM1,
		false,
		vr.UN,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006E),
		"Selector ST Value", // SelectorSTValue
		"SelectorSTValue",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x006F),
		"Selector UC Value", // SelectorUCValue
		"SelectorUCValue",
		vm.VM1_n,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0070),
		"Selector UT Value", // SelectorUTValue
		"SelectorUTValue",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0071),
		"Selector UR Value", // SelectorURValue
		"SelectorURValue",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0072),
		"Selector DS Value", // SelectorDSValue
		"SelectorDSValue",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0073),
		"Selector OD Value", // SelectorODValue
		"SelectorODValue",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0074),
		"Selector FD Value", // SelectorFDValue
		"SelectorFDValue",
		vm.VM1_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0075),
		"Selector OL Value", // SelectorOLValue
		"SelectorOLValue",
		vm.VM1,
		false,
		vr.OL,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0076),
		"Selector FL Value", // SelectorFLValue
		"SelectorFLValue",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0078),
		"Selector UL Value", // SelectorULValue
		"SelectorULValue",
		vm.VM1_n,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x007A),
		"Selector US Value", // SelectorUSValue
		"SelectorUSValue",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x007C),
		"Selector SL Value", // SelectorSLValue
		"SelectorSLValue",
		vm.VM1_n,
		false,
		vr.SL,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x007E),
		"Selector SS Value", // SelectorSSValue
		"SelectorSSValue",
		vm.VM1_n,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x007F),
		"Selector UI Value", // SelectorUIValue
		"SelectorUIValue",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0080),
		"Selector Code Sequence Value", // SelectorCodeSequenceValue
		"SelectorCodeSequenceValue",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0081),
		"Selector OV Value", // SelectorOVValue
		"SelectorOVValue",
		vm.VM1,
		false,
		vr.OV,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0082),
		"Selector SV Value", // SelectorSVValue
		"SelectorSVValue",
		vm.VM1_n,
		false,
		vr.SV,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0083),
		"Selector UV Value", // SelectorUVValue
		"SelectorUVValue",
		vm.VM1_n,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0100),
		"Number of Screens", // NumberOfScreens
		"NumberOfScreens",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0102),
		"Nominal Screen Definition Sequence", // NominalScreenDefinitionSequence
		"NominalScreenDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0104),
		"Number of Vertical Pixels", // NumberOfVerticalPixels
		"NumberOfVerticalPixels",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0106),
		"Number of Horizontal Pixels", // NumberOfHorizontalPixels
		"NumberOfHorizontalPixels",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0108),
		"Display Environment Spatial Position", // DisplayEnvironmentSpatialPosition
		"DisplayEnvironmentSpatialPosition",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x010A),
		"Screen Minimum Grayscale Bit Depth", // ScreenMinimumGrayscaleBitDepth
		"ScreenMinimumGrayscaleBitDepth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x010C),
		"Screen Minimum Color Bit Depth", // ScreenMinimumColorBitDepth
		"ScreenMinimumColorBitDepth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x010E),
		"Application Maximum Repaint Time", // ApplicationMaximumRepaintTime
		"ApplicationMaximumRepaintTime",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0200),
		"Display Sets Sequence", // DisplaySetsSequence
		"DisplaySetsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0202),
		"Display Set Number", // DisplaySetNumber
		"DisplaySetNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0203),
		"Display Set Label", // DisplaySetLabel
		"DisplaySetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0204),
		"Display Set Presentation Group", // DisplaySetPresentationGroup
		"DisplaySetPresentationGroup",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0206),
		"Display Set Presentation Group Description", // DisplaySetPresentationGroupDescription
		"DisplaySetPresentationGroupDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0208),
		"Partial Data Display Handling", // PartialDataDisplayHandling
		"PartialDataDisplayHandling",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0210),
		"Synchronized Scrolling Sequence", // SynchronizedScrollingSequence
		"SynchronizedScrollingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0212),
		"Display Set Scrolling Group", // DisplaySetScrollingGroup
		"DisplaySetScrollingGroup",
		vm.VM2_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0214),
		"Navigation Indicator Sequence", // NavigationIndicatorSequence
		"NavigationIndicatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0216),
		"Navigation Display Set", // NavigationDisplaySet
		"NavigationDisplaySet",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0218),
		"Reference Display Sets", // ReferenceDisplaySets
		"ReferenceDisplaySets",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0300),
		"Image Boxes Sequence", // ImageBoxesSequence
		"ImageBoxesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0302),
		"Image Box Number", // ImageBoxNumber
		"ImageBoxNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0304),
		"Image Box Layout Type", // ImageBoxLayoutType
		"ImageBoxLayoutType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0306),
		"Image Box Tile Horizontal Dimension", // ImageBoxTileHorizontalDimension
		"ImageBoxTileHorizontalDimension",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0308),
		"Image Box Tile Vertical Dimension", // ImageBoxTileVerticalDimension
		"ImageBoxTileVerticalDimension",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0310),
		"Image Box Scroll Direction", // ImageBoxScrollDirection
		"ImageBoxScrollDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0312),
		"Image Box Small Scroll Type", // ImageBoxSmallScrollType
		"ImageBoxSmallScrollType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0314),
		"Image Box Small Scroll Amount", // ImageBoxSmallScrollAmount
		"ImageBoxSmallScrollAmount",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0316),
		"Image Box Large Scroll Type", // ImageBoxLargeScrollType
		"ImageBoxLargeScrollType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0318),
		"Image Box Large Scroll Amount", // ImageBoxLargeScrollAmount
		"ImageBoxLargeScrollAmount",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0320),
		"Image Box Overlap Priority", // ImageBoxOverlapPriority
		"ImageBoxOverlapPriority",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0330),
		"Cine Relative to Real-Time", // CineRelativeToRealTime
		"CineRelativeToRealTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0400),
		"Filter Operations Sequence", // FilterOperationsSequence
		"FilterOperationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0402),
		"Filter-by Category", // FilterByCategory
		"FilterByCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0404),
		"Filter-by Attribute Presence", // FilterByAttributePresence
		"FilterByAttributePresence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0406),
		"Filter-by Operator", // FilterByOperator
		"FilterByOperator",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0420),
		"Structured Display Background CIELab Value", // StructuredDisplayBackgroundCIELabValue
		"StructuredDisplayBackgroundCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0421),
		"Empty Image Box CIELab Value", // EmptyImageBoxCIELabValue
		"EmptyImageBoxCIELabValue",
		vm.VM3,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0422),
		"Structured Display Image Box Sequence", // StructuredDisplayImageBoxSequence
		"StructuredDisplayImageBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0424),
		"Structured Display Text Box Sequence", // StructuredDisplayTextBoxSequence
		"StructuredDisplayTextBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0427),
		"Referenced First Frame Sequence", // ReferencedFirstFrameSequence
		"ReferencedFirstFrameSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0430),
		"Image Box Synchronization Sequence", // ImageBoxSynchronizationSequence
		"ImageBoxSynchronizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0432),
		"Synchronized Image Box List", // SynchronizedImageBoxList
		"SynchronizedImageBoxList",
		vm.VM2_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0434),
		"Type of Synchronization", // TypeOfSynchronization
		"TypeOfSynchronization",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0500),
		"Blending Operation Type", // BlendingOperationType
		"BlendingOperationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0510),
		"Reformatting Operation Type", // ReformattingOperationType
		"ReformattingOperationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0512),
		"Reformatting Thickness", // ReformattingThickness
		"ReformattingThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0514),
		"Reformatting Interval", // ReformattingInterval
		"ReformattingInterval",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0516),
		"Reformatting Operation Initial View Direction", // ReformattingOperationInitialViewDirection
		"ReformattingOperationInitialViewDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0520),
		"3D Rendering Type", // ThreeDRenderingType
		"ThreeDRenderingType",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0600),
		"Sorting Operations Sequence", // SortingOperationsSequence
		"SortingOperationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0602),
		"Sort-by Category", // SortByCategory
		"SortByCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0604),
		"Sorting Direction", // SortingDirection
		"SortingDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0700),
		"Display Set Patient Orientation", // DisplaySetPatientOrientation
		"DisplaySetPatientOrientation",
		vm.VM2,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0702),
		"VOI Type", // VOIType
		"VOIType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0704),
		"Pseudo-Color Type", // PseudoColorType
		"PseudoColorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0705),
		"Pseudo-Color Palette Instance Reference Sequence", // PseudoColorPaletteInstanceReferenceSequence
		"PseudoColorPaletteInstanceReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0706),
		"Show Grayscale Inverted", // ShowGrayscaleInverted
		"ShowGrayscaleInverted",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0710),
		"Show Image True Size Flag", // ShowImageTrueSizeFlag
		"ShowImageTrueSizeFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0712),
		"Show Graphic Annotation Flag", // ShowGraphicAnnotationFlag
		"ShowGraphicAnnotationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0714),
		"Show Patient Demographics Flag", // ShowPatientDemographicsFlag
		"ShowPatientDemographicsFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0716),
		"Show Acquisition Techniques Flag", // ShowAcquisitionTechniquesFlag
		"ShowAcquisitionTechniquesFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0717),
		"Display Set Horizontal Justification", // DisplaySetHorizontalJustification
		"DisplaySetHorizontalJustification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0072, 0x0718),
		"Display Set Vertical Justification", // DisplaySetVerticalJustification
		"DisplaySetVerticalJustification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x0120),
		"Continuation Start Meterset", // ContinuationStartMeterset
		"ContinuationStartMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x0121),
		"Continuation End Meterset", // ContinuationEndMeterset
		"ContinuationEndMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1000),
		"Procedure Step State", // ProcedureStepState
		"ProcedureStepState",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1002),
		"Procedure Step Progress Information Sequence", // ProcedureStepProgressInformationSequence
		"ProcedureStepProgressInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1004),
		"Procedure Step Progress", // ProcedureStepProgress
		"ProcedureStepProgress",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1006),
		"Procedure Step Progress Description", // ProcedureStepProgressDescription
		"ProcedureStepProgressDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1007),
		"Procedure Step Progress Parameters Sequence", // ProcedureStepProgressParametersSequence
		"ProcedureStepProgressParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1008),
		"Procedure Step Communications URI Sequence", // ProcedureStepCommunicationsURISequence
		"ProcedureStepCommunicationsURISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x100A),
		"Contact URI", // ContactURI
		"ContactURI",
		vm.VM1,
		false,
		vr.UR,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x100C),
		"Contact Display Name", // ContactDisplayName
		"ContactDisplayName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x100E),
		"Procedure Step Discontinuation Reason Code Sequence", // ProcedureStepDiscontinuationReasonCodeSequence
		"ProcedureStepDiscontinuationReasonCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1020),
		"Beam Task Sequence", // BeamTaskSequence
		"BeamTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1022),
		"Beam Task Type", // BeamTaskType
		"BeamTaskType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1024),
		"Beam Order Index (Trial)", // BeamOrderIndexTrial
		"BeamOrderIndexTrial",
		vm.VM1,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1025),
		"Autosequence Flag", // AutosequenceFlag
		"AutosequenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1026),
		"Table Top Vertical Adjusted Position", // TableTopVerticalAdjustedPosition
		"TableTopVerticalAdjustedPosition",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1027),
		"Table Top Longitudinal Adjusted Position", // TableTopLongitudinalAdjustedPosition
		"TableTopLongitudinalAdjustedPosition",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1028),
		"Table Top Lateral Adjusted Position", // TableTopLateralAdjustedPosition
		"TableTopLateralAdjustedPosition",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x102A),
		"Patient Support Adjusted Angle", // PatientSupportAdjustedAngle
		"PatientSupportAdjustedAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x102B),
		"Table Top Eccentric Adjusted Angle", // TableTopEccentricAdjustedAngle
		"TableTopEccentricAdjustedAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x102C),
		"Table Top Pitch Adjusted Angle", // TableTopPitchAdjustedAngle
		"TableTopPitchAdjustedAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x102D),
		"Table Top Roll Adjusted Angle", // TableTopRollAdjustedAngle
		"TableTopRollAdjustedAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1030),
		"Delivery Verification Image Sequence", // DeliveryVerificationImageSequence
		"DeliveryVerificationImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1032),
		"Verification Image Timing", // VerificationImageTiming
		"VerificationImageTiming",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1034),
		"Double Exposure Flag", // DoubleExposureFlag
		"DoubleExposureFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1036),
		"Double Exposure Ordering", // DoubleExposureOrdering
		"DoubleExposureOrdering",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1038),
		"Double Exposure Meterset (Trial)", // DoubleExposureMetersetTrial
		"DoubleExposureMetersetTrial",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x103A),
		"Double Exposure Field Delta (Trial)", // DoubleExposureFieldDeltaTrial
		"DoubleExposureFieldDeltaTrial",
		vm.VM4,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1040),
		"Related Reference RT Image Sequence", // RelatedReferenceRTImageSequence
		"RelatedReferenceRTImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1042),
		"General Machine Verification Sequence", // GeneralMachineVerificationSequence
		"GeneralMachineVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1044),
		"Conventional Machine Verification Sequence", // ConventionalMachineVerificationSequence
		"ConventionalMachineVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1046),
		"Ion Machine Verification Sequence", // IonMachineVerificationSequence
		"IonMachineVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1048),
		"Failed Attributes Sequence", // FailedAttributesSequence
		"FailedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x104A),
		"Overridden Attributes Sequence", // OverriddenAttributesSequence
		"OverriddenAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x104C),
		"Conventional Control Point Verification Sequence", // ConventionalControlPointVerificationSequence
		"ConventionalControlPointVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x104E),
		"Ion Control Point Verification Sequence", // IonControlPointVerificationSequence
		"IonControlPointVerificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1050),
		"Attribute Occurrence Sequence", // AttributeOccurrenceSequence
		"AttributeOccurrenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1052),
		"Attribute Occurrence Pointer", // AttributeOccurrencePointer
		"AttributeOccurrencePointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1054),
		"Attribute Item Selector", // AttributeItemSelector
		"AttributeItemSelector",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1056),
		"Attribute Occurrence Private Creator", // AttributeOccurrencePrivateCreator
		"AttributeOccurrencePrivateCreator",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1057),
		"Selector Sequence Pointer Items", // SelectorSequencePointerItems
		"SelectorSequencePointerItems",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1200),
		"Scheduled Procedure Step Priority", // ScheduledProcedureStepPriority
		"ScheduledProcedureStepPriority",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1202),
		"Worklist Label", // WorklistLabel
		"WorklistLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1204),
		"Procedure Step Label", // ProcedureStepLabel
		"ProcedureStepLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1210),
		"Scheduled Processing Parameters Sequence", // ScheduledProcessingParametersSequence
		"ScheduledProcessingParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1212),
		"Performed Processing Parameters Sequence", // PerformedProcessingParametersSequence
		"PerformedProcessingParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1216),
		"Unified Procedure Step Performed Procedure Sequence", // UnifiedProcedureStepPerformedProcedureSequence
		"UnifiedProcedureStepPerformedProcedureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1220),
		"Related Procedure Step Sequence", // RelatedProcedureStepSequence
		"RelatedProcedureStepSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1222),
		"Procedure Step Relationship Type", // ProcedureStepRelationshipType
		"ProcedureStepRelationshipType",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1224),
		"Replaced Procedure Step Sequence", // ReplacedProcedureStepSequence
		"ReplacedProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1230),
		"Deletion Lock", // DeletionLock
		"DeletionLock",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1234),
		"Receiving AE", // ReceivingAE
		"ReceivingAE",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1236),
		"Requesting AE", // RequestingAE
		"RequestingAE",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1238),
		"Reason for Cancellation", // ReasonForCancellation
		"ReasonForCancellation",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1242),
		"SCP Status", // SCPStatus
		"SCPStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1244),
		"Subscription List Status", // SubscriptionListStatus
		"SubscriptionListStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1246),
		"Unified Procedure Step List Status", // UnifiedProcedureStepListStatus
		"UnifiedProcedureStepListStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1324),
		"Beam Order Index", // BeamOrderIndex
		"BeamOrderIndex",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1338),
		"Double Exposure Meterset", // DoubleExposureMeterset
		"DoubleExposureMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x133A),
		"Double Exposure Field Delta", // DoubleExposureFieldDelta
		"DoubleExposureFieldDelta",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1401),
		"Brachy Task Sequence", // BrachyTaskSequence
		"BrachyTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1402),
		"Continuation Start Total Reference Air Kerma", // ContinuationStartTotalReferenceAirKerma
		"ContinuationStartTotalReferenceAirKerma",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1403),
		"Continuation End Total Reference Air Kerma", // ContinuationEndTotalReferenceAirKerma
		"ContinuationEndTotalReferenceAirKerma",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1404),
		"Continuation Pulse Number", // ContinuationPulseNumber
		"ContinuationPulseNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1405),
		"Channel Delivery Order Sequence", // ChannelDeliveryOrderSequence
		"ChannelDeliveryOrderSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1406),
		"Referenced Channel Number", // ReferencedChannelNumber
		"ReferencedChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1407),
		"Start Cumulative Time Weight", // StartCumulativeTimeWeight
		"StartCumulativeTimeWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1408),
		"End Cumulative Time Weight", // EndCumulativeTimeWeight
		"EndCumulativeTimeWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x1409),
		"Omitted Channel Sequence", // OmittedChannelSequence
		"OmittedChannelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140A),
		"Reason for Channel Omission", // ReasonForChannelOmission
		"ReasonForChannelOmission",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140B),
		"Reason for Channel Omission Description", // ReasonForChannelOmissionDescription
		"ReasonForChannelOmissionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140C),
		"Channel Delivery Order Index", // ChannelDeliveryOrderIndex
		"ChannelDeliveryOrderIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140D),
		"Channel Delivery Continuation Sequence", // ChannelDeliveryContinuationSequence
		"ChannelDeliveryContinuationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0074, 0x140E),
		"Omitted Application Setup Sequence", // OmittedApplicationSetupSequence
		"OmittedApplicationSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0001),
		"Implant Assembly Template Name", // ImplantAssemblyTemplateName
		"ImplantAssemblyTemplateName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0003),
		"Implant Assembly Template Issuer", // ImplantAssemblyTemplateIssuer
		"ImplantAssemblyTemplateIssuer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0006),
		"Implant Assembly Template Version", // ImplantAssemblyTemplateVersion
		"ImplantAssemblyTemplateVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0008),
		"Replaced Implant Assembly Template Sequence", // ReplacedImplantAssemblyTemplateSequence
		"ReplacedImplantAssemblyTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x000A),
		"Implant Assembly Template Type", // ImplantAssemblyTemplateType
		"ImplantAssemblyTemplateType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x000C),
		"Original Implant Assembly Template Sequence", // OriginalImplantAssemblyTemplateSequence
		"OriginalImplantAssemblyTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x000E),
		"Derivation Implant Assembly Template Sequence", // DerivationImplantAssemblyTemplateSequence
		"DerivationImplantAssemblyTemplateSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0010),
		"Implant Assembly Template Target Anatomy Sequence", // ImplantAssemblyTemplateTargetAnatomySequence
		"ImplantAssemblyTemplateTargetAnatomySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0020),
		"Procedure Type Code Sequence", // ProcedureTypeCodeSequence
		"ProcedureTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0030),
		"Surgical Technique", // SurgicalTechnique
		"SurgicalTechnique",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0032),
		"Component Types Sequence", // ComponentTypesSequence
		"ComponentTypesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0034),
		"Component Type Code Sequence", // ComponentTypeCodeSequence
		"ComponentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0036),
		"Exclusive Component Type", // ExclusiveComponentType
		"ExclusiveComponentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0038),
		"Mandatory Component Type", // MandatoryComponentType
		"MandatoryComponentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0040),
		"Component Sequence", // ComponentSequence
		"ComponentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0055),
		"Component ID", // ComponentID
		"ComponentID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0060),
		"Component Assembly Sequence", // ComponentAssemblySequence
		"ComponentAssemblySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0070),
		"Component 1 Referenced ID", // Component1ReferencedID
		"Component1ReferencedID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0080),
		"Component 1 Referenced Mating Feature Set ID", // Component1ReferencedMatingFeatureSetID
		"Component1ReferencedMatingFeatureSetID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x0090),
		"Component 1 Referenced Mating Feature ID", // Component1ReferencedMatingFeatureID
		"Component1ReferencedMatingFeatureID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x00A0),
		"Component 2 Referenced ID", // Component2ReferencedID
		"Component2ReferencedID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x00B0),
		"Component 2 Referenced Mating Feature Set ID", // Component2ReferencedMatingFeatureSetID
		"Component2ReferencedMatingFeatureSetID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0076, 0x00C0),
		"Component 2 Referenced Mating Feature ID", // Component2ReferencedMatingFeatureID
		"Component2ReferencedMatingFeatureID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0001),
		"Implant Template Group Name", // ImplantTemplateGroupName
		"ImplantTemplateGroupName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0010),
		"Implant Template Group Description", // ImplantTemplateGroupDescription
		"ImplantTemplateGroupDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0020),
		"Implant Template Group Issuer", // ImplantTemplateGroupIssuer
		"ImplantTemplateGroupIssuer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0024),
		"Implant Template Group Version", // ImplantTemplateGroupVersion
		"ImplantTemplateGroupVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0026),
		"Replaced Implant Template Group Sequence", // ReplacedImplantTemplateGroupSequence
		"ReplacedImplantTemplateGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0028),
		"Implant Template Group Target Anatomy Sequence", // ImplantTemplateGroupTargetAnatomySequence
		"ImplantTemplateGroupTargetAnatomySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x002A),
		"Implant Template Group Members Sequence", // ImplantTemplateGroupMembersSequence
		"ImplantTemplateGroupMembersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x002E),
		"Implant Template Group Member ID", // ImplantTemplateGroupMemberID
		"ImplantTemplateGroupMemberID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0050),
		"3D Implant Template Group Member Matching Point", // ThreeDImplantTemplateGroupMemberMatchingPoint
		"ThreeDImplantTemplateGroupMemberMatchingPoint",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0060),
		"3D Implant Template Group Member Matching Axes", // ThreeDImplantTemplateGroupMemberMatchingAxes
		"ThreeDImplantTemplateGroupMemberMatchingAxes",
		vm.MustParse("9"),
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0070),
		"Implant Template Group Member Matching 2D Coordinates Sequence", // ImplantTemplateGroupMemberMatching2DCoordinatesSequence
		"ImplantTemplateGroupMemberMatching2DCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x0090),
		"2D Implant Template Group Member Matching Point", // TwoDImplantTemplateGroupMemberMatchingPoint
		"TwoDImplantTemplateGroupMemberMatchingPoint",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00A0),
		"2D Implant Template Group Member Matching Axes", // TwoDImplantTemplateGroupMemberMatchingAxes
		"TwoDImplantTemplateGroupMemberMatchingAxes",
		vm.VM4,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B0),
		"Implant Template Group Variation Dimension Sequence", // ImplantTemplateGroupVariationDimensionSequence
		"ImplantTemplateGroupVariationDimensionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B2),
		"Implant Template Group Variation Dimension Name", // ImplantTemplateGroupVariationDimensionName
		"ImplantTemplateGroupVariationDimensionName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B4),
		"Implant Template Group Variation Dimension Rank Sequence", // ImplantTemplateGroupVariationDimensionRankSequence
		"ImplantTemplateGroupVariationDimensionRankSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B6),
		"Referenced Implant Template Group Member ID", // ReferencedImplantTemplateGroupMemberID
		"ReferencedImplantTemplateGroupMemberID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0078, 0x00B8),
		"Implant Template Group Variation Dimension Rank", // ImplantTemplateGroupVariationDimensionRank
		"ImplantTemplateGroupVariationDimensionRank",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0001),
		"Surface Scan Acquisition Type Code Sequence", // SurfaceScanAcquisitionTypeCodeSequence
		"SurfaceScanAcquisitionTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0002),
		"Surface Scan Mode Code Sequence", // SurfaceScanModeCodeSequence
		"SurfaceScanModeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0003),
		"Registration Method Code Sequence", // RegistrationMethodCodeSequence
		"RegistrationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0004),
		"Shot Duration Time", // ShotDurationTime
		"ShotDurationTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0005),
		"Shot Offset Time", // ShotOffsetTime
		"ShotOffsetTime",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0006),
		"Surface Point Presentation Value Data", // SurfacePointPresentationValueData
		"SurfacePointPresentationValueData",
		vm.VM1_n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0007),
		"Surface Point Color CIELab Value Data", // SurfacePointColorCIELabValueData
		"SurfacePointColorCIELabValueData",
		vm.VM3_3n,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0008),
		"UV Mapping Sequence", // UVMappingSequence
		"UVMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0009),
		"Texture Label", // TextureLabel
		"TextureLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0010),
		"U Value Data", // UValueData
		"UValueData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0011),
		"V Value Data", // VValueData
		"VValueData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0012),
		"Referenced Texture Sequence", // ReferencedTextureSequence
		"ReferencedTextureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0080, 0x0013),
		"Referenced Surface Data Sequence", // ReferencedSurfaceDataSequence
		"ReferencedSurfaceDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0001),
		"Assessment Summary", // AssessmentSummary
		"AssessmentSummary",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0003),
		"Assessment Summary Description", // AssessmentSummaryDescription
		"AssessmentSummaryDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0004),
		"Assessed SOP Instance Sequence", // AssessedSOPInstanceSequence
		"AssessedSOPInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0005),
		"Referenced Comparison SOP Instance Sequence", // ReferencedComparisonSOPInstanceSequence
		"ReferencedComparisonSOPInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0006),
		"Number of Assessment Observations", // NumberOfAssessmentObservations
		"NumberOfAssessmentObservations",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0007),
		"Assessment Observations Sequence", // AssessmentObservationsSequence
		"AssessmentObservationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0008),
		"Observation Significance", // ObservationSignificance
		"ObservationSignificance",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x000A),
		"Observation Description", // ObservationDescription
		"ObservationDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x000C),
		"Structured Constraint Observation Sequence", // StructuredConstraintObservationSequence
		"StructuredConstraintObservationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0010),
		"Assessed Attribute Value Sequence", // AssessedAttributeValueSequence
		"AssessedAttributeValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0016),
		"Assessment Set ID", // AssessmentSetID
		"AssessmentSetID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0017),
		"Assessment Requester Sequence", // AssessmentRequesterSequence
		"AssessmentRequesterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0018),
		"Selector Attribute Name", // SelectorAttributeName
		"SelectorAttributeName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0019),
		"Selector Attribute Keyword", // SelectorAttributeKeyword
		"SelectorAttributeKeyword",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0021),
		"Assessment Type Code Sequence", // AssessmentTypeCodeSequence
		"AssessmentTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0022),
		"Observation Basis Code Sequence", // ObservationBasisCodeSequence
		"ObservationBasisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0023),
		"Assessment Label", // AssessmentLabel
		"AssessmentLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0032),
		"Constraint Type", // ConstraintType
		"ConstraintType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0033),
		"Specification Selection Guidance", // SpecificationSelectionGuidance
		"SpecificationSelectionGuidance",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0034),
		"Constraint Value Sequence", // ConstraintValueSequence
		"ConstraintValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0035),
		"Recommended Default Value Sequence", // RecommendedDefaultValueSequence
		"RecommendedDefaultValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0036),
		"Constraint Violation Significance", // ConstraintViolationSignificance
		"ConstraintViolationSignificance",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0037),
		"Constraint Violation Condition", // ConstraintViolationCondition
		"ConstraintViolationCondition",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x0082, 0x0038),
		"Modifiable Constraint Flag", // ModifiableConstraintFlag
		"ModifiableConstraintFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0130),
		"Storage Media File-set ID", // StorageMediaFileSetID
		"StorageMediaFileSetID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0140),
		"Storage Media File-set UID", // StorageMediaFileSetUID
		"StorageMediaFileSetUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0200),
		"Icon Image Sequence", // IconImageSequence
		"IconImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0904),
		"Topic Title", // TopicTitle
		"TopicTitle",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0906),
		"Topic Subject", // TopicSubject
		"TopicSubject",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0910),
		"Topic Author", // TopicAuthor
		"TopicAuthor",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0088, 0x0912),
		"Topic Keywords", // TopicKeywords
		"TopicKeywords",
		vm.VM1_32,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0100, 0x0410),
		"SOP Instance Status", // SOPInstanceStatus
		"SOPInstanceStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0100, 0x0420),
		"SOP Authorization DateTime", // SOPAuthorizationDateTime
		"SOPAuthorizationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0100, 0x0424),
		"SOP Authorization Comment", // SOPAuthorizationComment
		"SOPAuthorizationComment",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x0100, 0x0426),
		"Authorization Equipment Certification Number", // AuthorizationEquipmentCertificationNumber
		"AuthorizationEquipmentCertificationNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0005),
		"MAC ID Number", // MACIDNumber
		"MACIDNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0010),
		"MAC Calculation Transfer Syntax UID", // MACCalculationTransferSyntaxUID
		"MACCalculationTransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0015),
		"MAC Algorithm", // MACAlgorithm
		"MACAlgorithm",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0020),
		"Data Elements Signed", // DataElementsSigned
		"DataElementsSigned",
		vm.VM1_n,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0100),
		"Digital Signature UID", // DigitalSignatureUID
		"DigitalSignatureUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0105),
		"Digital Signature DateTime", // DigitalSignatureDateTime
		"DigitalSignatureDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0110),
		"Certificate Type", // CertificateType
		"CertificateType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0115),
		"Certificate of Signer", // CertificateOfSigner
		"CertificateOfSigner",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0120),
		"Signature", // Signature
		"Signature",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0305),
		"Certified Timestamp Type", // CertifiedTimestampType
		"CertifiedTimestampType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0310),
		"Certified Timestamp", // CertifiedTimestamp
		"CertifiedTimestamp",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0401),
		"Digital Signature Purpose Code Sequence", // DigitalSignaturePurposeCodeSequence
		"DigitalSignaturePurposeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0402),
		"Referenced Digital Signature Sequence", // ReferencedDigitalSignatureSequence
		"ReferencedDigitalSignatureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0403),
		"Referenced SOP Instance MAC Sequence", // ReferencedSOPInstanceMACSequence
		"ReferencedSOPInstanceMACSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0404),
		"MAC", // MAC
		"MAC",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0500),
		"Encrypted Attributes Sequence", // EncryptedAttributesSequence
		"EncryptedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0510),
		"Encrypted Content Transfer Syntax UID", // EncryptedContentTransferSyntaxUID
		"EncryptedContentTransferSyntaxUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0520),
		"Encrypted Content", // EncryptedContent
		"EncryptedContent",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0550),
		"Modified Attributes Sequence", // ModifiedAttributesSequence
		"ModifiedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0551),
		"Nonconforming Modified Attributes Sequence", // NonconformingModifiedAttributesSequence
		"NonconformingModifiedAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0552),
		"Nonconforming Data Element Value", // NonconformingDataElementValue
		"NonconformingDataElementValue",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0561),
		"Original Attributes Sequence", // OriginalAttributesSequence
		"OriginalAttributesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0562),
		"Attribute Modification DateTime", // AttributeModificationDateTime
		"AttributeModificationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0563),
		"Modifying System", // ModifyingSystem
		"ModifyingSystem",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0564),
		"Source of Previous Values", // SourceOfPreviousValues
		"SourceOfPreviousValues",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0565),
		"Reason for the Attribute Modification", // ReasonForTheAttributeModification
		"ReasonForTheAttributeModification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x0400, 0x0600),
		"Instance Origin Status", // InstanceOriginStatus
		"InstanceOriginStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx0)"),
		"Escape Triplet", // EscapeTriplet
		"EscapeTriplet",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx1)"),
		"Run Length Triplet", // RunLengthTriplet
		"RunLengthTriplet",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx2)"),
		"Huffman Table Size", // HuffmanTableSize
		"HuffmanTableSize",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx3)"),
		"Huffman Table Triplet", // HuffmanTableTriplet
		"HuffmanTableTriplet",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx4)"),
		"Shift Table Size", // ShiftTableSize
		"ShiftTableSize",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1000,xxx5)"),
		"Shift Table Triplet", // ShiftTableTriplet
		"ShiftTableTriplet",
		vm.VM3,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(1010,xxxx)"),
		"Zonal Map", // ZonalMap
		"ZonalMap",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0010),
		"Number of Copies", // NumberOfCopies
		"NumberOfCopies",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x001E),
		"Printer Configuration Sequence", // PrinterConfigurationSequence
		"PrinterConfigurationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0020),
		"Print Priority", // PrintPriority
		"PrintPriority",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0030),
		"Medium Type", // MediumType
		"MediumType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0040),
		"Film Destination", // FilmDestination
		"FilmDestination",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0050),
		"Film Session Label", // FilmSessionLabel
		"FilmSessionLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0060),
		"Memory Allocation", // MemoryAllocation
		"MemoryAllocation",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0061),
		"Maximum Memory Allocation", // MaximumMemoryAllocation
		"MaximumMemoryAllocation",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0062),
		"Color Image Printing Flag", // ColorImagePrintingFlag
		"ColorImagePrintingFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0063),
		"Collation Flag", // CollationFlag
		"CollationFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0065),
		"Annotation Flag", // AnnotationFlag
		"AnnotationFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0067),
		"Image Overlay Flag", // ImageOverlayFlag
		"ImageOverlayFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0069),
		"Presentation LUT Flag", // PresentationLUTFlag
		"PresentationLUTFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x006A),
		"Image Box Presentation LUT Flag", // ImageBoxPresentationLUTFlag
		"ImageBoxPresentationLUTFlag",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A0),
		"Memory Bit Depth", // MemoryBitDepth
		"MemoryBitDepth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A1),
		"Printing Bit Depth", // PrintingBitDepth
		"PrintingBitDepth",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A2),
		"Media Installed Sequence", // MediaInstalledSequence
		"MediaInstalledSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A4),
		"Other Media Available Sequence", // OtherMediaAvailableSequence
		"OtherMediaAvailableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x00A8),
		"Supported Image Display Formats Sequence", // SupportedImageDisplayFormatsSequence
		"SupportedImageDisplayFormatsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0500),
		"Referenced Film Box Sequence", // ReferencedFilmBoxSequence
		"ReferencedFilmBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2000, 0x0510),
		"Referenced Stored Print Sequence", // ReferencedStoredPrintSequence
		"ReferencedStoredPrintSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0010),
		"Image Display Format", // ImageDisplayFormat
		"ImageDisplayFormat",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0030),
		"Annotation Display Format ID", // AnnotationDisplayFormatID
		"AnnotationDisplayFormatID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0040),
		"Film Orientation", // FilmOrientation
		"FilmOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0050),
		"Film Size ID", // FilmSizeID
		"FilmSizeID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0052),
		"Printer Resolution ID", // PrinterResolutionID
		"PrinterResolutionID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0054),
		"Default Printer Resolution ID", // DefaultPrinterResolutionID
		"DefaultPrinterResolutionID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0060),
		"Magnification Type", // MagnificationType
		"MagnificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0080),
		"Smoothing Type", // SmoothingType
		"SmoothingType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x00A6),
		"Default Magnification Type", // DefaultMagnificationType
		"DefaultMagnificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x00A7),
		"Other Magnification Types Available", // OtherMagnificationTypesAvailable
		"OtherMagnificationTypesAvailable",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x00A8),
		"Default Smoothing Type", // DefaultSmoothingType
		"DefaultSmoothingType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x00A9),
		"Other Smoothing Types Available", // OtherSmoothingTypesAvailable
		"OtherSmoothingTypesAvailable",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0100),
		"Border Density", // BorderDensity
		"BorderDensity",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0110),
		"Empty Image Density", // EmptyImageDensity
		"EmptyImageDensity",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0120),
		"Min Density", // MinDensity
		"MinDensity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0130),
		"Max Density", // MaxDensity
		"MaxDensity",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0140),
		"Trim", // Trim
		"Trim",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0150),
		"Configuration Information", // ConfigurationInformation
		"ConfigurationInformation",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0152),
		"Configuration Information Description", // ConfigurationInformationDescription
		"ConfigurationInformationDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0154),
		"Maximum Collated Films", // MaximumCollatedFilms
		"MaximumCollatedFilms",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x015E),
		"Illumination", // Illumination
		"Illumination",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0160),
		"Reflected Ambient Light", // ReflectedAmbientLight
		"ReflectedAmbientLight",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0376),
		"Printer Pixel Spacing", // PrinterPixelSpacing
		"PrinterPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0500),
		"Referenced Film Session Sequence", // ReferencedFilmSessionSequence
		"ReferencedFilmSessionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0510),
		"Referenced Image Box Sequence", // ReferencedImageBoxSequence
		"ReferencedImageBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2010, 0x0520),
		"Referenced Basic Annotation Box Sequence", // ReferencedBasicAnnotationBoxSequence
		"ReferencedBasicAnnotationBoxSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0010),
		"Image Box Position", // ImageBoxPosition
		"ImageBoxPosition",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0020),
		"Polarity", // Polarity
		"Polarity",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0030),
		"Requested Image Size", // RequestedImageSize
		"RequestedImageSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0040),
		"Requested Decimate/Crop Behavior", // RequestedDecimateCropBehavior
		"RequestedDecimateCropBehavior",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0050),
		"Requested Resolution ID", // RequestedResolutionID
		"RequestedResolutionID",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x00A0),
		"Requested Image Size Flag", // RequestedImageSizeFlag
		"RequestedImageSizeFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x00A2),
		"Decimate/Crop Result", // DecimateCropResult
		"DecimateCropResult",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0110),
		"Basic Grayscale Image Sequence", // BasicGrayscaleImageSequence
		"BasicGrayscaleImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0111),
		"Basic Color Image Sequence", // BasicColorImageSequence
		"BasicColorImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0130),
		"Referenced Image Overlay Box Sequence", // ReferencedImageOverlayBoxSequence
		"ReferencedImageOverlayBoxSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2020, 0x0140),
		"Referenced VOI LUT Box Sequence", // ReferencedVOILUTBoxSequence
		"ReferencedVOILUTBoxSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2030, 0x0010),
		"Annotation Position", // AnnotationPosition
		"AnnotationPosition",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2030, 0x0020),
		"Text String", // TextString
		"TextString",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0010),
		"Referenced Overlay Plane Sequence", // ReferencedOverlayPlaneSequence
		"ReferencedOverlayPlaneSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0011),
		"Referenced Overlay Plane Groups", // ReferencedOverlayPlaneGroups
		"ReferencedOverlayPlaneGroups",
		vm.VM1_99,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0020),
		"Overlay Pixel Data Sequence", // OverlayPixelDataSequence
		"OverlayPixelDataSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0060),
		"Overlay Magnification Type", // OverlayMagnificationType
		"OverlayMagnificationType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0070),
		"Overlay Smoothing Type", // OverlaySmoothingType
		"OverlaySmoothingType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0072),
		"Overlay or Image Magnification", // OverlayOrImageMagnification
		"OverlayOrImageMagnification",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0074),
		"Magnify to Number of Columns", // MagnifyToNumberOfColumns
		"MagnifyToNumberOfColumns",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0080),
		"Overlay Foreground Density", // OverlayForegroundDensity
		"OverlayForegroundDensity",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0082),
		"Overlay Background Density", // OverlayBackgroundDensity
		"OverlayBackgroundDensity",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0090),
		"Overlay Mode", // OverlayMode
		"OverlayMode",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0100),
		"Threshold Density", // ThresholdDensity
		"ThresholdDensity",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2040, 0x0500),
		"Referenced Image Box Sequence (Retired)", // ReferencedImageBoxSequenceRetired
		"ReferencedImageBoxSequenceRetired",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2050, 0x0010),
		"Presentation LUT Sequence", // PresentationLUTSequence
		"PresentationLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2050, 0x0020),
		"Presentation LUT Shape", // PresentationLUTShape
		"PresentationLUTShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2050, 0x0500),
		"Referenced Presentation LUT Sequence", // ReferencedPresentationLUTSequence
		"ReferencedPresentationLUTSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0010),
		"Print Job ID", // PrintJobID
		"PrintJobID",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0020),
		"Execution Status", // ExecutionStatus
		"ExecutionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0030),
		"Execution Status Info", // ExecutionStatusInfo
		"ExecutionStatusInfo",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0040),
		"Creation Date", // CreationDate
		"CreationDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0050),
		"Creation Time", // CreationTime
		"CreationTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0070),
		"Originator", // Originator
		"Originator",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0140),
		"Destination AE", // DestinationAE
		"DestinationAE",
		vm.VM1,
		false,
		vr.AE,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0160),
		"Owner ID", // OwnerID
		"OwnerID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0170),
		"Number of Films", // NumberOfFilms
		"NumberOfFilms",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x2100, 0x0500),
		"Referenced Print Job Sequence (Pull Stored Print)", // ReferencedPrintJobSequencePullStoredPrint
		"ReferencedPrintJobSequencePullStoredPrint",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2110, 0x0010),
		"Printer Status", // PrinterStatus
		"PrinterStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2110, 0x0020),
		"Printer Status Info", // PrinterStatusInfo
		"PrinterStatusInfo",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2110, 0x0030),
		"Printer Name", // PrinterName
		"PrinterName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x2110, 0x0099),
		"Print Queue ID", // PrintQueueID
		"PrintQueueID",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x2120, 0x0010),
		"Queue Status", // QueueStatus
		"QueueStatus",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2120, 0x0050),
		"Print Job Description Sequence", // PrintJobDescriptionSequence
		"PrintJobDescriptionSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2120, 0x0070),
		"Referenced Print Job Sequence", // ReferencedPrintJobSequence
		"ReferencedPrintJobSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0010),
		"Print Management Capabilities Sequence", // PrintManagementCapabilitiesSequence
		"PrintManagementCapabilitiesSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0015),
		"Printer Characteristics Sequence", // PrinterCharacteristicsSequence
		"PrinterCharacteristicsSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0030),
		"Film Box Content Sequence", // FilmBoxContentSequence
		"FilmBoxContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0040),
		"Image Box Content Sequence", // ImageBoxContentSequence
		"ImageBoxContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0050),
		"Annotation Content Sequence", // AnnotationContentSequence
		"AnnotationContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0060),
		"Image Overlay Box Content Sequence", // ImageOverlayBoxContentSequence
		"ImageOverlayBoxContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x0080),
		"Presentation LUT Content Sequence", // PresentationLUTContentSequence
		"PresentationLUTContentSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x00A0),
		"Proposed Study Sequence", // ProposedStudySequence
		"ProposedStudySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2130, 0x00C0),
		"Original Image Sequence", // OriginalImageSequence
		"OriginalImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0001),
		"Label Using Information Extracted From Instances", // LabelUsingInformationExtractedFromInstances
		"LabelUsingInformationExtractedFromInstances",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0002),
		"Label Text", // LabelText
		"LabelText",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0003),
		"Label Style Selection", // LabelStyleSelection
		"LabelStyleSelection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0004),
		"Media Disposition", // MediaDisposition
		"MediaDisposition",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0005),
		"Barcode Value", // BarcodeValue
		"BarcodeValue",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0006),
		"Barcode Symbology", // BarcodeSymbology
		"BarcodeSymbology",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0007),
		"Allow Media Splitting", // AllowMediaSplitting
		"AllowMediaSplitting",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0008),
		"Include Non-DICOM Objects", // IncludeNonDICOMObjects
		"IncludeNonDICOMObjects",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0009),
		"Include Display Application", // IncludeDisplayApplication
		"IncludeDisplayApplication",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000A),
		"Preserve Composite Instances After Media Creation", // PreserveCompositeInstancesAfterMediaCreation
		"PreserveCompositeInstancesAfterMediaCreation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000B),
		"Total Number of Pieces of Media Created", // TotalNumberOfPiecesOfMediaCreated
		"TotalNumberOfPiecesOfMediaCreated",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000C),
		"Requested Media Application Profile", // RequestedMediaApplicationProfile
		"RequestedMediaApplicationProfile",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000D),
		"Referenced Storage Media Sequence", // ReferencedStorageMediaSequence
		"ReferencedStorageMediaSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000E),
		"Failure Attributes", // FailureAttributes
		"FailureAttributes",
		vm.VM1_n,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x000F),
		"Allow Lossy Compression", // AllowLossyCompression
		"AllowLossyCompression",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x2200, 0x0020),
		"Request Priority", // RequestPriority
		"RequestPriority",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0002),
		"RT Image Label", // RTImageLabel
		"RTImageLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0003),
		"RT Image Name", // RTImageName
		"RTImageName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0004),
		"RT Image Description", // RTImageDescription
		"RTImageDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x000A),
		"Reported Values Origin", // ReportedValuesOrigin
		"ReportedValuesOrigin",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x000C),
		"RT Image Plane", // RTImagePlane
		"RTImagePlane",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x000D),
		"X-Ray Image Receptor Translation", // XRayImageReceptorTranslation
		"XRayImageReceptorTranslation",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x000E),
		"X-Ray Image Receptor Angle", // XRayImageReceptorAngle
		"XRayImageReceptorAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0010),
		"RT Image Orientation", // RTImageOrientation
		"RTImageOrientation",
		vm.VM6,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0011),
		"Image Plane Pixel Spacing", // ImagePlanePixelSpacing
		"ImagePlanePixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0012),
		"RT Image Position", // RTImagePosition
		"RTImagePosition",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0020),
		"Radiation Machine Name", // RadiationMachineName
		"RadiationMachineName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0022),
		"Radiation Machine SAD", // RadiationMachineSAD
		"RadiationMachineSAD",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0024),
		"Radiation Machine SSD", // RadiationMachineSSD
		"RadiationMachineSSD",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0026),
		"RT Image SID", // RTImageSID
		"RTImageSID",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0028),
		"Source to Reference Object Distance", // SourceToReferenceObjectDistance
		"SourceToReferenceObjectDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0029),
		"Fraction Number", // FractionNumber
		"FractionNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0030),
		"Exposure Sequence", // ExposureSequence
		"ExposureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0032),
		"Meterset Exposure", // MetersetExposure
		"MetersetExposure",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0034),
		"Diaphragm Position", // DiaphragmPosition
		"DiaphragmPosition",
		vm.VM4,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0040),
		"Fluence Map Sequence", // FluenceMapSequence
		"FluenceMapSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0041),
		"Fluence Data Source", // FluenceDataSource
		"FluenceDataSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0042),
		"Fluence Data Scale", // FluenceDataScale
		"FluenceDataScale",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0050),
		"Primary Fluence Mode Sequence", // PrimaryFluenceModeSequence
		"PrimaryFluenceModeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0051),
		"Fluence Mode", // FluenceMode
		"FluenceMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0052),
		"Fluence Mode ID", // FluenceModeID
		"FluenceModeID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0100),
		"Selected Frame Number", // SelectedFrameNumber
		"SelectedFrameNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0101),
		"Selected Frame Functional Groups Sequence", // SelectedFrameFunctionalGroupsSequence
		"SelectedFrameFunctionalGroupsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0102),
		"RT Image Frame General Content Sequence", // RTImageFrameGeneralContentSequence
		"RTImageFrameGeneralContentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0103),
		"RT Image Frame Context Sequence", // RTImageFrameContextSequence
		"RTImageFrameContextSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0104),
		"RT Image Scope Sequence", // RTImageScopeSequence
		"RTImageScopeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0105),
		"Beam Modifier Coordinates Presence Flag", // BeamModifierCoordinatesPresenceFlag
		"BeamModifierCoordinatesPresenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0106),
		"Start Cumulative Meterset", // StartCumulativeMeterset
		"StartCumulativeMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0107),
		"Stop Cumulative Meterset", // StopCumulativeMeterset
		"StopCumulativeMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0108),
		"RT Acquisition Patient Position Sequence", // RTAcquisitionPatientPositionSequence
		"RTAcquisitionPatientPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0109),
		"RT Image Frame Imaging Device Position Sequence", // RTImageFrameImagingDevicePositionSequence
		"RTImageFrameImagingDevicePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010A),
		"RT Image Frame kV Radiation Acquisition Sequence", // RTImageFramekVRadiationAcquisitionSequence
		"RTImageFramekVRadiationAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010B),
		"RT Image Frame MV Radiation Acquisition Sequence", // RTImageFrameMVRadiationAcquisitionSequence
		"RTImageFrameMVRadiationAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010C),
		"RT Image Frame Radiation Acquisition Sequence", // RTImageFrameRadiationAcquisitionSequence
		"RTImageFrameRadiationAcquisitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010D),
		"Imaging Source Position Sequence", // ImagingSourcePositionSequence
		"ImagingSourcePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010E),
		"Image Receptor Position Sequence", // ImageReceptorPositionSequence
		"ImageReceptorPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x010F),
		"Device Position to Equipment Mapping Matrix", // DevicePositionToEquipmentMappingMatrix
		"DevicePositionToEquipmentMappingMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0110),
		"Device Position Parameter Sequence", // DevicePositionParameterSequence
		"DevicePositionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0111),
		"Imaging Source Location Specification Type", // ImagingSourceLocationSpecificationType
		"ImagingSourceLocationSpecificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0112),
		"Imaging Device Location Matrix Sequence", // ImagingDeviceLocationMatrixSequence
		"ImagingDeviceLocationMatrixSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0113),
		"Imaging Device Location Parameter Sequence", // ImagingDeviceLocationParameterSequence
		"ImagingDeviceLocationParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0114),
		"Imaging Aperture Sequence", // ImagingApertureSequence
		"ImagingApertureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0115),
		"Imaging Aperture Specification Type", // ImagingApertureSpecificationType
		"ImagingApertureSpecificationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0116),
		"Number of Acquisition Devices", // NumberOfAcquisitionDevices
		"NumberOfAcquisitionDevices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0117),
		"Acquisition Device Sequence", // AcquisitionDeviceSequence
		"AcquisitionDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0118),
		"Acquisition Task Sequence", // AcquisitionTaskSequence
		"AcquisitionTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0119),
		"Acquisition Task Workitem Code Sequence", // AcquisitionTaskWorkitemCodeSequence
		"AcquisitionTaskWorkitemCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011A),
		"Acquisition Subtask Sequence", // AcquisitionSubtaskSequence
		"AcquisitionSubtaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011B),
		"Subtask Workitem Code Sequence", // SubtaskWorkitemCodeSequence
		"SubtaskWorkitemCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011C),
		"Acquisition Task Index", // AcquisitionTaskIndex
		"AcquisitionTaskIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011D),
		"Acquisition Subtask Index", // AcquisitionSubtaskIndex
		"AcquisitionSubtaskIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011E),
		"Referenced Baseline Parameters RT Radiation Instance Sequence", // ReferencedBaselineParametersRTRadiationInstanceSequence
		"ReferencedBaselineParametersRTRadiationInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x011F),
		"Position Acquisition Template Identification Sequence", // PositionAcquisitionTemplateIdentificationSequence
		"PositionAcquisitionTemplateIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0120),
		"Position Acquisition Template ID", // PositionAcquisitionTemplateID
		"PositionAcquisitionTemplateID",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0121),
		"Position Acquisition Template Name", // PositionAcquisitionTemplateName
		"PositionAcquisitionTemplateName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0122),
		"Position Acquisition Template Code Sequence", // PositionAcquisitionTemplateCodeSequence
		"PositionAcquisitionTemplateCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0123),
		"Position Acquisition Template Description", // PositionAcquisitionTemplateDescription
		"PositionAcquisitionTemplateDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0124),
		"Acquisition Task Applicability Sequence", // AcquisitionTaskApplicabilitySequence
		"AcquisitionTaskApplicabilitySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0125),
		"Projection Imaging Acquisition Parameter Sequence", // ProjectionImagingAcquisitionParameterSequence
		"ProjectionImagingAcquisitionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0126),
		"CT Imaging Acquisition Parameter Sequence", // CTImagingAcquisitionParameterSequence
		"CTImagingAcquisitionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0127),
		"KV Imaging Generation Parameters Sequence", // KVImagingGenerationParametersSequence
		"KVImagingGenerationParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0128),
		"MV Imaging Generation Parameters Sequence", // MVImagingGenerationParametersSequence
		"MVImagingGenerationParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0129),
		"Acquisition Signal Type", // AcquisitionSignalType
		"AcquisitionSignalType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012A),
		"Acquisition Method", // AcquisitionMethod
		"AcquisitionMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012B),
		"Scan Start Position Sequence", // ScanStartPositionSequence
		"ScanStartPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012C),
		"Scan Stop Position Sequence", // ScanStopPositionSequence
		"ScanStopPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012D),
		"Imaging Source to Beam Modifier Definition Plane Distance", // ImagingSourceToBeamModifierDefinitionPlaneDistance
		"ImagingSourceToBeamModifierDefinitionPlaneDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012E),
		"Scan Arc Type", // ScanArcType
		"ScanArcType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x012F),
		"Detector Positioning Type", // DetectorPositioningType
		"DetectorPositioningType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0130),
		"Additional RT Accessory Device Sequence", // AdditionalRTAccessoryDeviceSequence
		"AdditionalRTAccessoryDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0131),
		"Device-Specific Acquisition Parameter Sequence", // DeviceSpecificAcquisitionParameterSequence
		"DeviceSpecificAcquisitionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0132),
		"Referenced Position Reference Instance Sequence", // ReferencedPositionReferenceInstanceSequence
		"ReferencedPositionReferenceInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0133),
		"Energy Derivation Code Sequence", // EnergyDerivationCodeSequence
		"EnergyDerivationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0134),
		"Maximum Cumulative Meterset Exposure", // MaximumCumulativeMetersetExposure
		"MaximumCumulativeMetersetExposure",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0135),
		"Acquisition Initiation Sequence", // AcquisitionInitiationSequence
		"AcquisitionInitiationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3002, 0x0136),
		"RT Cone-Beam Imaging Geometry Sequence", // RTConeBeamImagingGeometrySequence
		"RTConeBeamImagingGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0001),
		"DVH Type", // DVHType
		"DVHType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0002),
		"Dose Units", // DoseUnits
		"DoseUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0004),
		"Dose Type", // DoseType
		"DoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0005),
		"Spatial Transform of Dose", // SpatialTransformOfDose
		"SpatialTransformOfDose",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0006),
		"Dose Comment", // DoseComment
		"DoseComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0008),
		"Normalization Point", // NormalizationPoint
		"NormalizationPoint",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x000A),
		"Dose Summation Type", // DoseSummationType
		"DoseSummationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x000C),
		"Grid Frame Offset Vector", // GridFrameOffsetVector
		"GridFrameOffsetVector",
		vm.VM2_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x000E),
		"Dose Grid Scaling", // DoseGridScaling
		"DoseGridScaling",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0010),
		"RT Dose ROI Sequence", // RTDoseROISequence
		"RTDoseROISequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0012),
		"Dose Value", // DoseValue
		"DoseValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0014),
		"Tissue Heterogeneity Correction", // TissueHeterogeneityCorrection
		"TissueHeterogeneityCorrection",
		vm.VM1_3,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0016),
		"Recommended Isodose Level Sequence", // RecommendedIsodoseLevelSequence
		"RecommendedIsodoseLevelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0040),
		"DVH Normalization Point", // DVHNormalizationPoint
		"DVHNormalizationPoint",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0042),
		"DVH Normalization Dose Value", // DVHNormalizationDoseValue
		"DVHNormalizationDoseValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0050),
		"DVH Sequence", // DVHSequence
		"DVHSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0052),
		"DVH Dose Scaling", // DVHDoseScaling
		"DVHDoseScaling",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0054),
		"DVH Volume Units", // DVHVolumeUnits
		"DVHVolumeUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0056),
		"DVH Number of Bins", // DVHNumberOfBins
		"DVHNumberOfBins",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0058),
		"DVH Data", // DVHData
		"DVHData",
		vm.VM2_2n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0060),
		"DVH Referenced ROI Sequence", // DVHReferencedROISequence
		"DVHReferencedROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0062),
		"DVH ROI Contribution Type", // DVHROIContributionType
		"DVHROIContributionType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0070),
		"DVH Minimum Dose", // DVHMinimumDose
		"DVHMinimumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0072),
		"DVH Maximum Dose", // DVHMaximumDose
		"DVHMaximumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0074),
		"DVH Mean Dose", // DVHMeanDose
		"DVHMeanDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0080),
		"Dose Calculation Model Sequence", // DoseCalculationModelSequence
		"DoseCalculationModelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0081),
		"Dose Calculation Algorithm Sequence", // DoseCalculationAlgorithmSequence
		"DoseCalculationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0082),
		"Commissioning Status", // CommissioningStatus
		"CommissioningStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0083),
		"Dose Calculation Model Parameter Sequence", // DoseCalculationModelParameterSequence
		"DoseCalculationModelParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3004, 0x0084),
		"Dose Deposition Calculation Medium", // DoseDepositionCalculationMedium
		"DoseDepositionCalculationMedium",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0002),
		"Structure Set Label", // StructureSetLabel
		"StructureSetLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0004),
		"Structure Set Name", // StructureSetName
		"StructureSetName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0006),
		"Structure Set Description", // StructureSetDescription
		"StructureSetDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0008),
		"Structure Set Date", // StructureSetDate
		"StructureSetDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0009),
		"Structure Set Time", // StructureSetTime
		"StructureSetTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0010),
		"Referenced Frame of Reference Sequence", // ReferencedFrameOfReferenceSequence
		"ReferencedFrameOfReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0012),
		"RT Referenced Study Sequence", // RTReferencedStudySequence
		"RTReferencedStudySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0014),
		"RT Referenced Series Sequence", // RTReferencedSeriesSequence
		"RTReferencedSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0016),
		"Contour Image Sequence", // ContourImageSequence
		"ContourImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0018),
		"Predecessor Structure Set Sequence", // PredecessorStructureSetSequence
		"PredecessorStructureSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0020),
		"Structure Set ROI Sequence", // StructureSetROISequence
		"StructureSetROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0022),
		"ROI Number", // ROINumber
		"ROINumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0024),
		"Referenced Frame of Reference UID", // ReferencedFrameOfReferenceUID
		"ReferencedFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0026),
		"ROI Name", // ROIName
		"ROIName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0028),
		"ROI Description", // ROIDescription
		"ROIDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x002A),
		"ROI Display Color", // ROIDisplayColor
		"ROIDisplayColor",
		vm.VM3,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x002C),
		"ROI Volume", // ROIVolume
		"ROIVolume",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x002D),
		"ROI DateTime", // ROIDateTime
		"ROIDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x002E),
		"ROI Observation DateTime", // ROIObservationDateTime
		"ROIObservationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0030),
		"RT Related ROI Sequence", // RTRelatedROISequence
		"RTRelatedROISequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0033),
		"RT ROI Relationship", // RTROIRelationship
		"RTROIRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0036),
		"ROI Generation Algorithm", // ROIGenerationAlgorithm
		"ROIGenerationAlgorithm",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0037),
		"ROI Derivation Algorithm Identification Sequence", // ROIDerivationAlgorithmIdentificationSequence
		"ROIDerivationAlgorithmIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0038),
		"ROI Generation Description", // ROIGenerationDescription
		"ROIGenerationDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0039),
		"ROI Contour Sequence", // ROIContourSequence
		"ROIContourSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0040),
		"Contour Sequence", // ContourSequence
		"ContourSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0042),
		"Contour Geometric Type", // ContourGeometricType
		"ContourGeometricType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0044),
		"Contour Slab Thickness", // ContourSlabThickness
		"ContourSlabThickness",
		vm.VM1,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0045),
		"Contour Offset Vector", // ContourOffsetVector
		"ContourOffsetVector",
		vm.VM3,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0046),
		"Number of Contour Points", // NumberOfContourPoints
		"NumberOfContourPoints",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0048),
		"Contour Number", // ContourNumber
		"ContourNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0049),
		"Attached Contours", // AttachedContours
		"AttachedContours",
		vm.VM1_n,
		true,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004A),
		"Source Pixel Planes Characteristics Sequence", // SourcePixelPlanesCharacteristicsSequence
		"SourcePixelPlanesCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004B),
		"Source Series Sequence", // SourceSeriesSequence
		"SourceSeriesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004C),
		"Source Series Information Sequence", // SourceSeriesInformationSequence
		"SourceSeriesInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004D),
		"ROI Creator Sequence", // ROICreatorSequence
		"ROICreatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004E),
		"ROI Interpreter Sequence", // ROIInterpreterSequence
		"ROIInterpreterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x004F),
		"ROI Observation Context Code Sequence", // ROIObservationContextCodeSequence
		"ROIObservationContextCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0050),
		"Contour Data", // ContourData
		"ContourData",
		vm.VM3_3n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0080),
		"RT ROI Observations Sequence", // RTROIObservationsSequence
		"RTROIObservationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0082),
		"Observation Number", // ObservationNumber
		"ObservationNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0084),
		"Referenced ROI Number", // ReferencedROINumber
		"ReferencedROINumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0085),
		"ROI Observation Label", // ROIObservationLabel
		"ROIObservationLabel",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0086),
		"RT ROI Identification Code Sequence", // RTROIIdentificationCodeSequence
		"RTROIIdentificationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x0088),
		"ROI Observation Description", // ROIObservationDescription
		"ROIObservationDescription",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00A0),
		"Related RT ROI Observations Sequence", // RelatedRTROIObservationsSequence
		"RelatedRTROIObservationsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00A4),
		"RT ROI Interpreted Type", // RTROIInterpretedType
		"RTROIInterpretedType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00A6),
		"ROI Interpreter", // ROIInterpreter
		"ROIInterpreter",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B0),
		"ROI Physical Properties Sequence", // ROIPhysicalPropertiesSequence
		"ROIPhysicalPropertiesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B2),
		"ROI Physical Property", // ROIPhysicalProperty
		"ROIPhysicalProperty",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B4),
		"ROI Physical Property Value", // ROIPhysicalPropertyValue
		"ROIPhysicalPropertyValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B6),
		"ROI Elemental Composition Sequence", // ROIElementalCompositionSequence
		"ROIElementalCompositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B7),
		"ROI Elemental Composition Atomic Number", // ROIElementalCompositionAtomicNumber
		"ROIElementalCompositionAtomicNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B8),
		"ROI Elemental Composition Atomic Mass Fraction", // ROIElementalCompositionAtomicMassFraction
		"ROIElementalCompositionAtomicMassFraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00B9),
		"Additional RT ROI Identification Code Sequence", // AdditionalRTROIIdentificationCodeSequence
		"AdditionalRTROIIdentificationCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C0),
		"Frame of Reference Relationship Sequence", // FrameOfReferenceRelationshipSequence
		"FrameOfReferenceRelationshipSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C2),
		"Related Frame of Reference UID", // RelatedFrameOfReferenceUID
		"RelatedFrameOfReferenceUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C4),
		"Frame of Reference Transformation Type", // FrameOfReferenceTransformationType
		"FrameOfReferenceTransformationType",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C6),
		"Frame of Reference Transformation Matrix", // FrameOfReferenceTransformationMatrix
		"FrameOfReferenceTransformationMatrix",
		vm.VM16,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C8),
		"Frame of Reference Transformation Comment", // FrameOfReferenceTransformationComment
		"FrameOfReferenceTransformationComment",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00C9),
		"Patient Location Coordinates Sequence", // PatientLocationCoordinatesSequence
		"PatientLocationCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00CA),
		"Patient Location Coordinates Code Sequence", // PatientLocationCoordinatesCodeSequence
		"PatientLocationCoordinatesCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3006, 0x00CB),
		"Patient Support Position Sequence", // PatientSupportPositionSequence
		"PatientSupportPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0010),
		"Measured Dose Reference Sequence", // MeasuredDoseReferenceSequence
		"MeasuredDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0012),
		"Measured Dose Description", // MeasuredDoseDescription
		"MeasuredDoseDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0014),
		"Measured Dose Type", // MeasuredDoseType
		"MeasuredDoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0016),
		"Measured Dose Value", // MeasuredDoseValue
		"MeasuredDoseValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0020),
		"Treatment Session Beam Sequence", // TreatmentSessionBeamSequence
		"TreatmentSessionBeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0021),
		"Treatment Session Ion Beam Sequence", // TreatmentSessionIonBeamSequence
		"TreatmentSessionIonBeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0022),
		"Current Fraction Number", // CurrentFractionNumber
		"CurrentFractionNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0024),
		"Treatment Control Point Date", // TreatmentControlPointDate
		"TreatmentControlPointDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0025),
		"Treatment Control Point Time", // TreatmentControlPointTime
		"TreatmentControlPointTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x002A),
		"Treatment Termination Status", // TreatmentTerminationStatus
		"TreatmentTerminationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x002B),
		"Treatment Termination Code", // TreatmentTerminationCode
		"TreatmentTerminationCode",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x002C),
		"Treatment Verification Status", // TreatmentVerificationStatus
		"TreatmentVerificationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0030),
		"Referenced Treatment Record Sequence", // ReferencedTreatmentRecordSequence
		"ReferencedTreatmentRecordSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0032),
		"Specified Primary Meterset", // SpecifiedPrimaryMeterset
		"SpecifiedPrimaryMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0033),
		"Specified Secondary Meterset", // SpecifiedSecondaryMeterset
		"SpecifiedSecondaryMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0036),
		"Delivered Primary Meterset", // DeliveredPrimaryMeterset
		"DeliveredPrimaryMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0037),
		"Delivered Secondary Meterset", // DeliveredSecondaryMeterset
		"DeliveredSecondaryMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x003A),
		"Specified Treatment Time", // SpecifiedTreatmentTime
		"SpecifiedTreatmentTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x003B),
		"Delivered Treatment Time", // DeliveredTreatmentTime
		"DeliveredTreatmentTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0040),
		"Control Point Delivery Sequence", // ControlPointDeliverySequence
		"ControlPointDeliverySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0041),
		"Ion Control Point Delivery Sequence", // IonControlPointDeliverySequence
		"IonControlPointDeliverySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0042),
		"Specified Meterset", // SpecifiedMeterset
		"SpecifiedMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0044),
		"Delivered Meterset", // DeliveredMeterset
		"DeliveredMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0045),
		"Meterset Rate Set", // MetersetRateSet
		"MetersetRateSet",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0046),
		"Meterset Rate Delivered", // MetersetRateDelivered
		"MetersetRateDelivered",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0047),
		"Scan Spot Metersets Delivered", // ScanSpotMetersetsDelivered
		"ScanSpotMetersetsDelivered",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0048),
		"Dose Rate Delivered", // DoseRateDelivered
		"DoseRateDelivered",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0050),
		"Treatment Summary Calculated Dose Reference Sequence", // TreatmentSummaryCalculatedDoseReferenceSequence
		"TreatmentSummaryCalculatedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0052),
		"Cumulative Dose to Dose Reference", // CumulativeDoseToDoseReference
		"CumulativeDoseToDoseReference",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0054),
		"First Treatment Date", // FirstTreatmentDate
		"FirstTreatmentDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0056),
		"Most Recent Treatment Date", // MostRecentTreatmentDate
		"MostRecentTreatmentDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x005A),
		"Number of Fractions Delivered", // NumberOfFractionsDelivered
		"NumberOfFractionsDelivered",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0060),
		"Override Sequence", // OverrideSequence
		"OverrideSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0061),
		"Parameter Sequence Pointer", // ParameterSequencePointer
		"ParameterSequencePointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0062),
		"Override Parameter Pointer", // OverrideParameterPointer
		"OverrideParameterPointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0063),
		"Parameter Item Index", // ParameterItemIndex
		"ParameterItemIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0064),
		"Measured Dose Reference Number", // MeasuredDoseReferenceNumber
		"MeasuredDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0065),
		"Parameter Pointer", // ParameterPointer
		"ParameterPointer",
		vm.VM1,
		false,
		vr.AT,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0066),
		"Override Reason", // OverrideReason
		"OverrideReason",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0067),
		"Parameter Value Number", // ParameterValueNumber
		"ParameterValueNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0068),
		"Corrected Parameter Sequence", // CorrectedParameterSequence
		"CorrectedParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x006A),
		"Correction Value", // CorrectionValue
		"CorrectionValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0070),
		"Calculated Dose Reference Sequence", // CalculatedDoseReferenceSequence
		"CalculatedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0072),
		"Calculated Dose Reference Number", // CalculatedDoseReferenceNumber
		"CalculatedDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0074),
		"Calculated Dose Reference Description", // CalculatedDoseReferenceDescription
		"CalculatedDoseReferenceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0076),
		"Calculated Dose Reference Dose Value", // CalculatedDoseReferenceDoseValue
		"CalculatedDoseReferenceDoseValue",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0078),
		"Start Meterset", // StartMeterset
		"StartMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x007A),
		"End Meterset", // EndMeterset
		"EndMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0080),
		"Referenced Measured Dose Reference Sequence", // ReferencedMeasuredDoseReferenceSequence
		"ReferencedMeasuredDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0082),
		"Referenced Measured Dose Reference Number", // ReferencedMeasuredDoseReferenceNumber
		"ReferencedMeasuredDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0090),
		"Referenced Calculated Dose Reference Sequence", // ReferencedCalculatedDoseReferenceSequence
		"ReferencedCalculatedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0092),
		"Referenced Calculated Dose Reference Number", // ReferencedCalculatedDoseReferenceNumber
		"ReferencedCalculatedDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A0),
		"Beam Limiting Device Leaf Pairs Sequence", // BeamLimitingDeviceLeafPairsSequence
		"BeamLimitingDeviceLeafPairsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A1),
		"Enhanced RT Beam Limiting Device Sequence", // EnhancedRTBeamLimitingDeviceSequence
		"EnhancedRTBeamLimitingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A2),
		"Enhanced RT Beam Limiting Opening Sequence", // EnhancedRTBeamLimitingOpeningSequence
		"EnhancedRTBeamLimitingOpeningSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A3),
		"Enhanced RT Beam Limiting Device Definition Flag", // EnhancedRTBeamLimitingDeviceDefinitionFlag
		"EnhancedRTBeamLimitingDeviceDefinitionFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00A4),
		"Parallel RT Beam Delimiter Opening Extents", // ParallelRTBeamDelimiterOpeningExtents
		"ParallelRTBeamDelimiterOpeningExtents",
		vm.VM2_2n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00B0),
		"Recorded Wedge Sequence", // RecordedWedgeSequence
		"RecordedWedgeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00C0),
		"Recorded Compensator Sequence", // RecordedCompensatorSequence
		"RecordedCompensatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00D0),
		"Recorded Block Sequence", // RecordedBlockSequence
		"RecordedBlockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00D1),
		"Recorded Block Slab Sequence", // RecordedBlockSlabSequence
		"RecordedBlockSlabSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00E0),
		"Treatment Summary Measured Dose Reference Sequence", // TreatmentSummaryMeasuredDoseReferenceSequence
		"TreatmentSummaryMeasuredDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00F0),
		"Recorded Snout Sequence", // RecordedSnoutSequence
		"RecordedSnoutSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00F2),
		"Recorded Range Shifter Sequence", // RecordedRangeShifterSequence
		"RecordedRangeShifterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00F4),
		"Recorded Lateral Spreading Device Sequence", // RecordedLateralSpreadingDeviceSequence
		"RecordedLateralSpreadingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x00F6),
		"Recorded Range Modulator Sequence", // RecordedRangeModulatorSequence
		"RecordedRangeModulatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0100),
		"Recorded Source Sequence", // RecordedSourceSequence
		"RecordedSourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0105),
		"Source Serial Number", // SourceSerialNumber
		"SourceSerialNumber",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0110),
		"Treatment Session Application Setup Sequence", // TreatmentSessionApplicationSetupSequence
		"TreatmentSessionApplicationSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0116),
		"Application Setup Check", // ApplicationSetupCheck
		"ApplicationSetupCheck",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0120),
		"Recorded Brachy Accessory Device Sequence", // RecordedBrachyAccessoryDeviceSequence
		"RecordedBrachyAccessoryDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0122),
		"Referenced Brachy Accessory Device Number", // ReferencedBrachyAccessoryDeviceNumber
		"ReferencedBrachyAccessoryDeviceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0130),
		"Recorded Channel Sequence", // RecordedChannelSequence
		"RecordedChannelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0132),
		"Specified Channel Total Time", // SpecifiedChannelTotalTime
		"SpecifiedChannelTotalTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0134),
		"Delivered Channel Total Time", // DeliveredChannelTotalTime
		"DeliveredChannelTotalTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0136),
		"Specified Number of Pulses", // SpecifiedNumberOfPulses
		"SpecifiedNumberOfPulses",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0138),
		"Delivered Number of Pulses", // DeliveredNumberOfPulses
		"DeliveredNumberOfPulses",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x013A),
		"Specified Pulse Repetition Interval", // SpecifiedPulseRepetitionInterval
		"SpecifiedPulseRepetitionInterval",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x013C),
		"Delivered Pulse Repetition Interval", // DeliveredPulseRepetitionInterval
		"DeliveredPulseRepetitionInterval",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0140),
		"Recorded Source Applicator Sequence", // RecordedSourceApplicatorSequence
		"RecordedSourceApplicatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0142),
		"Referenced Source Applicator Number", // ReferencedSourceApplicatorNumber
		"ReferencedSourceApplicatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0150),
		"Recorded Channel Shield Sequence", // RecordedChannelShieldSequence
		"RecordedChannelShieldSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0152),
		"Referenced Channel Shield Number", // ReferencedChannelShieldNumber
		"ReferencedChannelShieldNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0160),
		"Brachy Control Point Delivered Sequence", // BrachyControlPointDeliveredSequence
		"BrachyControlPointDeliveredSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0162),
		"Safe Position Exit Date", // SafePositionExitDate
		"SafePositionExitDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0164),
		"Safe Position Exit Time", // SafePositionExitTime
		"SafePositionExitTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0166),
		"Safe Position Return Date", // SafePositionReturnDate
		"SafePositionReturnDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0168),
		"Safe Position Return Time", // SafePositionReturnTime
		"SafePositionReturnTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0171),
		"Pulse Specific Brachy Control Point Delivered Sequence", // PulseSpecificBrachyControlPointDeliveredSequence
		"PulseSpecificBrachyControlPointDeliveredSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0172),
		"Pulse Number", // PulseNumber
		"PulseNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0173),
		"Brachy Pulse Control Point Delivered Sequence", // BrachyPulseControlPointDeliveredSequence
		"BrachyPulseControlPointDeliveredSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0200),
		"Current Treatment Status", // CurrentTreatmentStatus
		"CurrentTreatmentStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0202),
		"Treatment Status Comment", // TreatmentStatusComment
		"TreatmentStatusComment",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0220),
		"Fraction Group Summary Sequence", // FractionGroupSummarySequence
		"FractionGroupSummarySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0223),
		"Referenced Fraction Number", // ReferencedFractionNumber
		"ReferencedFractionNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0224),
		"Fraction Group Type", // FractionGroupType
		"FractionGroupType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0230),
		"Beam Stopper Position", // BeamStopperPosition
		"BeamStopperPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0240),
		"Fraction Status Summary Sequence", // FractionStatusSummarySequence
		"FractionStatusSummarySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0250),
		"Treatment Date", // TreatmentDate
		"TreatmentDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3008, 0x0251),
		"Treatment Time", // TreatmentTime
		"TreatmentTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0002),
		"RT Plan Label", // RTPlanLabel
		"RTPlanLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0003),
		"RT Plan Name", // RTPlanName
		"RTPlanName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0004),
		"RT Plan Description", // RTPlanDescription
		"RTPlanDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0006),
		"RT Plan Date", // RTPlanDate
		"RTPlanDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0007),
		"RT Plan Time", // RTPlanTime
		"RTPlanTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0009),
		"Treatment Protocols", // TreatmentProtocols
		"TreatmentProtocols",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x000A),
		"Plan Intent", // PlanIntent
		"PlanIntent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x000B),
		"Treatment Sites", // TreatmentSites
		"TreatmentSites",
		vm.VM1_n,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x000C),
		"RT Plan Geometry", // RTPlanGeometry
		"RTPlanGeometry",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x000E),
		"Prescription Description", // PrescriptionDescription
		"PrescriptionDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0010),
		"Dose Reference Sequence", // DoseReferenceSequence
		"DoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0012),
		"Dose Reference Number", // DoseReferenceNumber
		"DoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0013),
		"Dose Reference UID", // DoseReferenceUID
		"DoseReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0014),
		"Dose Reference Structure Type", // DoseReferenceStructureType
		"DoseReferenceStructureType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0015),
		"Nominal Beam Energy Unit", // NominalBeamEnergyUnit
		"NominalBeamEnergyUnit",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0016),
		"Dose Reference Description", // DoseReferenceDescription
		"DoseReferenceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0018),
		"Dose Reference Point Coordinates", // DoseReferencePointCoordinates
		"DoseReferencePointCoordinates",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x001A),
		"Nominal Prior Dose", // NominalPriorDose
		"NominalPriorDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0020),
		"Dose Reference Type", // DoseReferenceType
		"DoseReferenceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0021),
		"Constraint Weight", // ConstraintWeight
		"ConstraintWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0022),
		"Delivery Warning Dose", // DeliveryWarningDose
		"DeliveryWarningDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0023),
		"Delivery Maximum Dose", // DeliveryMaximumDose
		"DeliveryMaximumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0025),
		"Target Minimum Dose", // TargetMinimumDose
		"TargetMinimumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0026),
		"Target Prescription Dose", // TargetPrescriptionDose
		"TargetPrescriptionDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0027),
		"Target Maximum Dose", // TargetMaximumDose
		"TargetMaximumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0028),
		"Target Underdose Volume Fraction", // TargetUnderdoseVolumeFraction
		"TargetUnderdoseVolumeFraction",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x002A),
		"Organ at Risk Full-volume Dose", // OrganAtRiskFullVolumeDose
		"OrganAtRiskFullVolumeDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x002B),
		"Organ at Risk Limit Dose", // OrganAtRiskLimitDose
		"OrganAtRiskLimitDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x002C),
		"Organ at Risk Maximum Dose", // OrganAtRiskMaximumDose
		"OrganAtRiskMaximumDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x002D),
		"Organ at Risk Overdose Volume Fraction", // OrganAtRiskOverdoseVolumeFraction
		"OrganAtRiskOverdoseVolumeFraction",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0040),
		"Tolerance Table Sequence", // ToleranceTableSequence
		"ToleranceTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0042),
		"Tolerance Table Number", // ToleranceTableNumber
		"ToleranceTableNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0043),
		"Tolerance Table Label", // ToleranceTableLabel
		"ToleranceTableLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0044),
		"Gantry Angle Tolerance", // GantryAngleTolerance
		"GantryAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0046),
		"Beam Limiting Device Angle Tolerance", // BeamLimitingDeviceAngleTolerance
		"BeamLimitingDeviceAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0048),
		"Beam Limiting Device Tolerance Sequence", // BeamLimitingDeviceToleranceSequence
		"BeamLimitingDeviceToleranceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004A),
		"Beam Limiting Device Position Tolerance", // BeamLimitingDevicePositionTolerance
		"BeamLimitingDevicePositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004B),
		"Snout Position Tolerance", // SnoutPositionTolerance
		"SnoutPositionTolerance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004C),
		"Patient Support Angle Tolerance", // PatientSupportAngleTolerance
		"PatientSupportAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004E),
		"Table Top Eccentric Angle Tolerance", // TableTopEccentricAngleTolerance
		"TableTopEccentricAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x004F),
		"Table Top Pitch Angle Tolerance", // TableTopPitchAngleTolerance
		"TableTopPitchAngleTolerance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0050),
		"Table Top Roll Angle Tolerance", // TableTopRollAngleTolerance
		"TableTopRollAngleTolerance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0051),
		"Table Top Vertical Position Tolerance", // TableTopVerticalPositionTolerance
		"TableTopVerticalPositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0052),
		"Table Top Longitudinal Position Tolerance", // TableTopLongitudinalPositionTolerance
		"TableTopLongitudinalPositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0053),
		"Table Top Lateral Position Tolerance", // TableTopLateralPositionTolerance
		"TableTopLateralPositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0054),
		"Table Top Position Alignment UID", // TableTopPositionAlignmentUID
		"TableTopPositionAlignmentUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0055),
		"RT Plan Relationship", // RTPlanRelationship
		"RTPlanRelationship",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0070),
		"Fraction Group Sequence", // FractionGroupSequence
		"FractionGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0071),
		"Fraction Group Number", // FractionGroupNumber
		"FractionGroupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0072),
		"Fraction Group Description", // FractionGroupDescription
		"FractionGroupDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0078),
		"Number of Fractions Planned", // NumberOfFractionsPlanned
		"NumberOfFractionsPlanned",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0079),
		"Number of Fraction Pattern Digits Per Day", // NumberOfFractionPatternDigitsPerDay
		"NumberOfFractionPatternDigitsPerDay",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x007A),
		"Repeat Fraction Cycle Length", // RepeatFractionCycleLength
		"RepeatFractionCycleLength",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x007B),
		"Fraction Pattern", // FractionPattern
		"FractionPattern",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0080),
		"Number of Beams", // NumberOfBeams
		"NumberOfBeams",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0082),
		"Beam Dose Specification Point", // BeamDoseSpecificationPoint
		"BeamDoseSpecificationPoint",
		vm.VM3,
		true,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0083),
		"Referenced Dose Reference UID", // ReferencedDoseReferenceUID
		"ReferencedDoseReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0084),
		"Beam Dose", // BeamDose
		"BeamDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0086),
		"Beam Meterset", // BeamMeterset
		"BeamMeterset",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0088),
		"Beam Dose Point Depth", // BeamDosePointDepth
		"BeamDosePointDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0089),
		"Beam Dose Point Equivalent Depth", // BeamDosePointEquivalentDepth
		"BeamDosePointEquivalentDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008A),
		"Beam Dose Point SSD", // BeamDosePointSSD
		"BeamDosePointSSD",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008B),
		"Beam Dose Meaning", // BeamDoseMeaning
		"BeamDoseMeaning",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008C),
		"Beam Dose Verification Control Point Sequence", // BeamDoseVerificationControlPointSequence
		"BeamDoseVerificationControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008D),
		"Average Beam Dose Point Depth", // AverageBeamDosePointDepth
		"AverageBeamDosePointDepth",
		vm.VM1,
		true,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008E),
		"Average Beam Dose Point Equivalent Depth", // AverageBeamDosePointEquivalentDepth
		"AverageBeamDosePointEquivalentDepth",
		vm.VM1,
		true,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x008F),
		"Average Beam Dose Point SSD", // AverageBeamDosePointSSD
		"AverageBeamDosePointSSD",
		vm.VM1,
		true,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0090),
		"Beam Dose Type", // BeamDoseType
		"BeamDoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0091),
		"Alternate Beam Dose", // AlternateBeamDose
		"AlternateBeamDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0092),
		"Alternate Beam Dose Type", // AlternateBeamDoseType
		"AlternateBeamDoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0093),
		"Depth Value Averaging Flag", // DepthValueAveragingFlag
		"DepthValueAveragingFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0094),
		"Beam Dose Point Source to External Contour Distance", // BeamDosePointSourceToExternalContourDistance
		"BeamDosePointSourceToExternalContourDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00A0),
		"Number of Brachy Application Setups", // NumberOfBrachyApplicationSetups
		"NumberOfBrachyApplicationSetups",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00A2),
		"Brachy Application Setup Dose Specification Point", // BrachyApplicationSetupDoseSpecificationPoint
		"BrachyApplicationSetupDoseSpecificationPoint",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00A4),
		"Brachy Application Setup Dose", // BrachyApplicationSetupDose
		"BrachyApplicationSetupDose",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B0),
		"Beam Sequence", // BeamSequence
		"BeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B2),
		"Treatment Machine Name", // TreatmentMachineName
		"TreatmentMachineName",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B3),
		"Primary Dosimeter Unit", // PrimaryDosimeterUnit
		"PrimaryDosimeterUnit",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B4),
		"Source-Axis Distance", // SourceAxisDistance
		"SourceAxisDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B6),
		"Beam Limiting Device Sequence", // BeamLimitingDeviceSequence
		"BeamLimitingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00B8),
		"RT Beam Limiting Device Type", // RTBeamLimitingDeviceType
		"RTBeamLimitingDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00BA),
		"Source to Beam Limiting Device Distance", // SourceToBeamLimitingDeviceDistance
		"SourceToBeamLimitingDeviceDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00BB),
		"Isocenter to Beam Limiting Device Distance", // IsocenterToBeamLimitingDeviceDistance
		"IsocenterToBeamLimitingDeviceDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00BC),
		"Number of Leaf/Jaw Pairs", // NumberOfLeafJawPairs
		"NumberOfLeafJawPairs",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00BE),
		"Leaf Position Boundaries", // LeafPositionBoundaries
		"LeafPositionBoundaries",
		vm.VM3_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C0),
		"Beam Number", // BeamNumber
		"BeamNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C2),
		"Beam Name", // BeamName
		"BeamName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C3),
		"Beam Description", // BeamDescription
		"BeamDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C4),
		"Beam Type", // BeamType
		"BeamType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C5),
		"Beam Delivery Duration Limit", // BeamDeliveryDurationLimit
		"BeamDeliveryDurationLimit",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C6),
		"Radiation Type", // RadiationType
		"RadiationType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C7),
		"High-Dose Technique Type", // HighDoseTechniqueType
		"HighDoseTechniqueType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00C8),
		"Reference Image Number", // ReferenceImageNumber
		"ReferenceImageNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00CA),
		"Planned Verification Image Sequence", // PlannedVerificationImageSequence
		"PlannedVerificationImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00CC),
		"Imaging Device-Specific Acquisition Parameters", // ImagingDeviceSpecificAcquisitionParameters
		"ImagingDeviceSpecificAcquisitionParameters",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00CE),
		"Treatment Delivery Type", // TreatmentDeliveryType
		"TreatmentDeliveryType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D0),
		"Number of Wedges", // NumberOfWedges
		"NumberOfWedges",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D1),
		"Wedge Sequence", // WedgeSequence
		"WedgeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D2),
		"Wedge Number", // WedgeNumber
		"WedgeNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D3),
		"Wedge Type", // WedgeType
		"WedgeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D4),
		"Wedge ID", // WedgeID
		"WedgeID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D5),
		"Wedge Angle", // WedgeAngle
		"WedgeAngle",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D6),
		"Wedge Factor", // WedgeFactor
		"WedgeFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D7),
		"Total Wedge Tray Water-Equivalent Thickness", // TotalWedgeTrayWaterEquivalentThickness
		"TotalWedgeTrayWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D8),
		"Wedge Orientation", // WedgeOrientation
		"WedgeOrientation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00D9),
		"Isocenter to Wedge Tray Distance", // IsocenterToWedgeTrayDistance
		"IsocenterToWedgeTrayDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DA),
		"Source to Wedge Tray Distance", // SourceToWedgeTrayDistance
		"SourceToWedgeTrayDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DB),
		"Wedge Thin Edge Position", // WedgeThinEdgePosition
		"WedgeThinEdgePosition",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DC),
		"Bolus ID", // BolusID
		"BolusID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DD),
		"Bolus Description", // BolusDescription
		"BolusDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00DE),
		"Effective Wedge Angle", // EffectiveWedgeAngle
		"EffectiveWedgeAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E0),
		"Number of Compensators", // NumberOfCompensators
		"NumberOfCompensators",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E1),
		"Material ID", // MaterialID
		"MaterialID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E2),
		"Total Compensator Tray Factor", // TotalCompensatorTrayFactor
		"TotalCompensatorTrayFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E3),
		"Compensator Sequence", // CompensatorSequence
		"CompensatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E4),
		"Compensator Number", // CompensatorNumber
		"CompensatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E5),
		"Compensator ID", // CompensatorID
		"CompensatorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E6),
		"Source to Compensator Tray Distance", // SourceToCompensatorTrayDistance
		"SourceToCompensatorTrayDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E7),
		"Compensator Rows", // CompensatorRows
		"CompensatorRows",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E8),
		"Compensator Columns", // CompensatorColumns
		"CompensatorColumns",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00E9),
		"Compensator Pixel Spacing", // CompensatorPixelSpacing
		"CompensatorPixelSpacing",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EA),
		"Compensator Position", // CompensatorPosition
		"CompensatorPosition",
		vm.VM2,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EB),
		"Compensator Transmission Data", // CompensatorTransmissionData
		"CompensatorTransmissionData",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EC),
		"Compensator Thickness Data", // CompensatorThicknessData
		"CompensatorThicknessData",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00ED),
		"Number of Boli", // NumberOfBoli
		"NumberOfBoli",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EE),
		"Compensator Type", // CompensatorType
		"CompensatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00EF),
		"Compensator Tray ID", // CompensatorTrayID
		"CompensatorTrayID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F0),
		"Number of Blocks", // NumberOfBlocks
		"NumberOfBlocks",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F2),
		"Total Block Tray Factor", // TotalBlockTrayFactor
		"TotalBlockTrayFactor",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F3),
		"Total Block Tray Water-Equivalent Thickness", // TotalBlockTrayWaterEquivalentThickness
		"TotalBlockTrayWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F4),
		"Block Sequence", // BlockSequence
		"BlockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F5),
		"Block Tray ID", // BlockTrayID
		"BlockTrayID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F6),
		"Source to Block Tray Distance", // SourceToBlockTrayDistance
		"SourceToBlockTrayDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F7),
		"Isocenter to Block Tray Distance", // IsocenterToBlockTrayDistance
		"IsocenterToBlockTrayDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F8),
		"Block Type", // BlockType
		"BlockType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00F9),
		"Accessory Code", // AccessoryCode
		"AccessoryCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00FA),
		"Block Divergence", // BlockDivergence
		"BlockDivergence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00FB),
		"Block Mounting Position", // BlockMountingPosition
		"BlockMountingPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00FC),
		"Block Number", // BlockNumber
		"BlockNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x00FE),
		"Block Name", // BlockName
		"BlockName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0100),
		"Block Thickness", // BlockThickness
		"BlockThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0102),
		"Block Transmission", // BlockTransmission
		"BlockTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0104),
		"Block Number of Points", // BlockNumberOfPoints
		"BlockNumberOfPoints",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0106),
		"Block Data", // BlockData
		"BlockData",
		vm.VM2_2n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0107),
		"Applicator Sequence", // ApplicatorSequence
		"ApplicatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0108),
		"Applicator ID", // ApplicatorID
		"ApplicatorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0109),
		"Applicator Type", // ApplicatorType
		"ApplicatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x010A),
		"Applicator Description", // ApplicatorDescription
		"ApplicatorDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x010C),
		"Cumulative Dose Reference Coefficient", // CumulativeDoseReferenceCoefficient
		"CumulativeDoseReferenceCoefficient",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x010E),
		"Final Cumulative Meterset Weight", // FinalCumulativeMetersetWeight
		"FinalCumulativeMetersetWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0110),
		"Number of Control Points", // NumberOfControlPoints
		"NumberOfControlPoints",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0111),
		"Control Point Sequence", // ControlPointSequence
		"ControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0112),
		"Control Point Index", // ControlPointIndex
		"ControlPointIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0114),
		"Nominal Beam Energy", // NominalBeamEnergy
		"NominalBeamEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0115),
		"Dose Rate Set", // DoseRateSet
		"DoseRateSet",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0116),
		"Wedge Position Sequence", // WedgePositionSequence
		"WedgePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0118),
		"Wedge Position", // WedgePosition
		"WedgePosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x011A),
		"Beam Limiting Device Position Sequence", // BeamLimitingDevicePositionSequence
		"BeamLimitingDevicePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x011C),
		"Leaf/Jaw Positions", // LeafJawPositions
		"LeafJawPositions",
		vm.VM2_2n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x011E),
		"Gantry Angle", // GantryAngle
		"GantryAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x011F),
		"Gantry Rotation Direction", // GantryRotationDirection
		"GantryRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0120),
		"Beam Limiting Device Angle", // BeamLimitingDeviceAngle
		"BeamLimitingDeviceAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0121),
		"Beam Limiting Device Rotation Direction", // BeamLimitingDeviceRotationDirection
		"BeamLimitingDeviceRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0122),
		"Patient Support Angle", // PatientSupportAngle
		"PatientSupportAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0123),
		"Patient Support Rotation Direction", // PatientSupportRotationDirection
		"PatientSupportRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0124),
		"Table Top Eccentric Axis Distance", // TableTopEccentricAxisDistance
		"TableTopEccentricAxisDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0125),
		"Table Top Eccentric Angle", // TableTopEccentricAngle
		"TableTopEccentricAngle",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0126),
		"Table Top Eccentric Rotation Direction", // TableTopEccentricRotationDirection
		"TableTopEccentricRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0128),
		"Table Top Vertical Position", // TableTopVerticalPosition
		"TableTopVerticalPosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0129),
		"Table Top Longitudinal Position", // TableTopLongitudinalPosition
		"TableTopLongitudinalPosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x012A),
		"Table Top Lateral Position", // TableTopLateralPosition
		"TableTopLateralPosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x012C),
		"Isocenter Position", // IsocenterPosition
		"IsocenterPosition",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x012E),
		"Surface Entry Point", // SurfaceEntryPoint
		"SurfaceEntryPoint",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0130),
		"Source to Surface Distance", // SourceToSurfaceDistance
		"SourceToSurfaceDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0131),
		"Average Beam Dose Point Source to External Contour Distance", // AverageBeamDosePointSourceToExternalContourDistance
		"AverageBeamDosePointSourceToExternalContourDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0132),
		"Source to External Contour Distance", // SourceToExternalContourDistance
		"SourceToExternalContourDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0133),
		"External Contour Entry Point", // ExternalContourEntryPoint
		"ExternalContourEntryPoint",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0134),
		"Cumulative Meterset Weight", // CumulativeMetersetWeight
		"CumulativeMetersetWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0140),
		"Table Top Pitch Angle", // TableTopPitchAngle
		"TableTopPitchAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0142),
		"Table Top Pitch Rotation Direction", // TableTopPitchRotationDirection
		"TableTopPitchRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0144),
		"Table Top Roll Angle", // TableTopRollAngle
		"TableTopRollAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0146),
		"Table Top Roll Rotation Direction", // TableTopRollRotationDirection
		"TableTopRollRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0148),
		"Head Fixation Angle", // HeadFixationAngle
		"HeadFixationAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x014A),
		"Gantry Pitch Angle", // GantryPitchAngle
		"GantryPitchAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x014C),
		"Gantry Pitch Rotation Direction", // GantryPitchRotationDirection
		"GantryPitchRotationDirection",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x014E),
		"Gantry Pitch Angle Tolerance", // GantryPitchAngleTolerance
		"GantryPitchAngleTolerance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0150),
		"Fixation Eye", // FixationEye
		"FixationEye",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0151),
		"Chair Head Frame Position", // ChairHeadFramePosition
		"ChairHeadFramePosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0152),
		"Head Fixation Angle Tolerance", // HeadFixationAngleTolerance
		"HeadFixationAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0153),
		"Chair Head Frame Position Tolerance", // ChairHeadFramePositionTolerance
		"ChairHeadFramePositionTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0154),
		"Fixation Light Azimuthal Angle Tolerance", // FixationLightAzimuthalAngleTolerance
		"FixationLightAzimuthalAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0155),
		"Fixation Light Polar Angle Tolerance", // FixationLightPolarAngleTolerance
		"FixationLightPolarAngleTolerance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0180),
		"Patient Setup Sequence", // PatientSetupSequence
		"PatientSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0182),
		"Patient Setup Number", // PatientSetupNumber
		"PatientSetupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0183),
		"Patient Setup Label", // PatientSetupLabel
		"PatientSetupLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0184),
		"Patient Additional Position", // PatientAdditionalPosition
		"PatientAdditionalPosition",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0190),
		"Fixation Device Sequence", // FixationDeviceSequence
		"FixationDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0192),
		"Fixation Device Type", // FixationDeviceType
		"FixationDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0194),
		"Fixation Device Label", // FixationDeviceLabel
		"FixationDeviceLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0196),
		"Fixation Device Description", // FixationDeviceDescription
		"FixationDeviceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0198),
		"Fixation Device Position", // FixationDevicePosition
		"FixationDevicePosition",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0199),
		"Fixation Device Pitch Angle", // FixationDevicePitchAngle
		"FixationDevicePitchAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x019A),
		"Fixation Device Roll Angle", // FixationDeviceRollAngle
		"FixationDeviceRollAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A0),
		"Shielding Device Sequence", // ShieldingDeviceSequence
		"ShieldingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A2),
		"Shielding Device Type", // ShieldingDeviceType
		"ShieldingDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A4),
		"Shielding Device Label", // ShieldingDeviceLabel
		"ShieldingDeviceLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A6),
		"Shielding Device Description", // ShieldingDeviceDescription
		"ShieldingDeviceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01A8),
		"Shielding Device Position", // ShieldingDevicePosition
		"ShieldingDevicePosition",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B0),
		"Setup Technique", // SetupTechnique
		"SetupTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B2),
		"Setup Technique Description", // SetupTechniqueDescription
		"SetupTechniqueDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B4),
		"Setup Device Sequence", // SetupDeviceSequence
		"SetupDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B6),
		"Setup Device Type", // SetupDeviceType
		"SetupDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01B8),
		"Setup Device Label", // SetupDeviceLabel
		"SetupDeviceLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01BA),
		"Setup Device Description", // SetupDeviceDescription
		"SetupDeviceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01BC),
		"Setup Device Parameter", // SetupDeviceParameter
		"SetupDeviceParameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01D0),
		"Setup Reference Description", // SetupReferenceDescription
		"SetupReferenceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01D2),
		"Table Top Vertical Setup Displacement", // TableTopVerticalSetupDisplacement
		"TableTopVerticalSetupDisplacement",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01D4),
		"Table Top Longitudinal Setup Displacement", // TableTopLongitudinalSetupDisplacement
		"TableTopLongitudinalSetupDisplacement",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x01D6),
		"Table Top Lateral Setup Displacement", // TableTopLateralSetupDisplacement
		"TableTopLateralSetupDisplacement",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0200),
		"Brachy Treatment Technique", // BrachyTreatmentTechnique
		"BrachyTreatmentTechnique",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0202),
		"Brachy Treatment Type", // BrachyTreatmentType
		"BrachyTreatmentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0206),
		"Treatment Machine Sequence", // TreatmentMachineSequence
		"TreatmentMachineSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0210),
		"Source Sequence", // SourceSequence
		"SourceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0212),
		"Source Number", // SourceNumber
		"SourceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0214),
		"Source Type", // SourceType
		"SourceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0216),
		"Source Manufacturer", // SourceManufacturer
		"SourceManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0218),
		"Active Source Diameter", // ActiveSourceDiameter
		"ActiveSourceDiameter",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x021A),
		"Active Source Length", // ActiveSourceLength
		"ActiveSourceLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x021B),
		"Source Model ID", // SourceModelID
		"SourceModelID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x021C),
		"Source Description", // SourceDescription
		"SourceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0222),
		"Source Encapsulation Nominal Thickness", // SourceEncapsulationNominalThickness
		"SourceEncapsulationNominalThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0224),
		"Source Encapsulation Nominal Transmission", // SourceEncapsulationNominalTransmission
		"SourceEncapsulationNominalTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0226),
		"Source Isotope Name", // SourceIsotopeName
		"SourceIsotopeName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0228),
		"Source Isotope Half Life", // SourceIsotopeHalfLife
		"SourceIsotopeHalfLife",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0229),
		"Source Strength Units", // SourceStrengthUnits
		"SourceStrengthUnits",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x022A),
		"Reference Air Kerma Rate", // ReferenceAirKermaRate
		"ReferenceAirKermaRate",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x022B),
		"Source Strength", // SourceStrength
		"SourceStrength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x022C),
		"Source Strength Reference Date", // SourceStrengthReferenceDate
		"SourceStrengthReferenceDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x022E),
		"Source Strength Reference Time", // SourceStrengthReferenceTime
		"SourceStrengthReferenceTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0230),
		"Application Setup Sequence", // ApplicationSetupSequence
		"ApplicationSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0232),
		"Application Setup Type", // ApplicationSetupType
		"ApplicationSetupType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0234),
		"Application Setup Number", // ApplicationSetupNumber
		"ApplicationSetupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0236),
		"Application Setup Name", // ApplicationSetupName
		"ApplicationSetupName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0238),
		"Application Setup Manufacturer", // ApplicationSetupManufacturer
		"ApplicationSetupManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0240),
		"Template Number", // TemplateNumber
		"TemplateNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0242),
		"Template Type", // TemplateType
		"TemplateType",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0244),
		"Template Name", // TemplateName
		"TemplateName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0250),
		"Total Reference Air Kerma", // TotalReferenceAirKerma
		"TotalReferenceAirKerma",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0260),
		"Brachy Accessory Device Sequence", // BrachyAccessoryDeviceSequence
		"BrachyAccessoryDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0262),
		"Brachy Accessory Device Number", // BrachyAccessoryDeviceNumber
		"BrachyAccessoryDeviceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0263),
		"Brachy Accessory Device ID", // BrachyAccessoryDeviceID
		"BrachyAccessoryDeviceID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0264),
		"Brachy Accessory Device Type", // BrachyAccessoryDeviceType
		"BrachyAccessoryDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0266),
		"Brachy Accessory Device Name", // BrachyAccessoryDeviceName
		"BrachyAccessoryDeviceName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x026A),
		"Brachy Accessory Device Nominal Thickness", // BrachyAccessoryDeviceNominalThickness
		"BrachyAccessoryDeviceNominalThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x026C),
		"Brachy Accessory Device Nominal Transmission", // BrachyAccessoryDeviceNominalTransmission
		"BrachyAccessoryDeviceNominalTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0271),
		"Channel Effective Length", // ChannelEffectiveLength
		"ChannelEffectiveLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0272),
		"Channel Inner Length", // ChannelInnerLength
		"ChannelInnerLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0273),
		"Afterloader Channel ID", // AfterloaderChannelID
		"AfterloaderChannelID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0274),
		"Source Applicator Tip Length", // SourceApplicatorTipLength
		"SourceApplicatorTipLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0280),
		"Channel Sequence", // ChannelSequence
		"ChannelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0282),
		"Channel Number", // ChannelNumber
		"ChannelNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0284),
		"Channel Length", // ChannelLength
		"ChannelLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0286),
		"Channel Total Time", // ChannelTotalTime
		"ChannelTotalTime",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0288),
		"Source Movement Type", // SourceMovementType
		"SourceMovementType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x028A),
		"Number of Pulses", // NumberOfPulses
		"NumberOfPulses",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x028C),
		"Pulse Repetition Interval", // PulseRepetitionInterval
		"PulseRepetitionInterval",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0290),
		"Source Applicator Number", // SourceApplicatorNumber
		"SourceApplicatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0291),
		"Source Applicator ID", // SourceApplicatorID
		"SourceApplicatorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0292),
		"Source Applicator Type", // SourceApplicatorType
		"SourceApplicatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0294),
		"Source Applicator Name", // SourceApplicatorName
		"SourceApplicatorName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0296),
		"Source Applicator Length", // SourceApplicatorLength
		"SourceApplicatorLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0298),
		"Source Applicator Manufacturer", // SourceApplicatorManufacturer
		"SourceApplicatorManufacturer",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x029C),
		"Source Applicator Wall Nominal Thickness", // SourceApplicatorWallNominalThickness
		"SourceApplicatorWallNominalThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x029E),
		"Source Applicator Wall Nominal Transmission", // SourceApplicatorWallNominalTransmission
		"SourceApplicatorWallNominalTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02A0),
		"Source Applicator Step Size", // SourceApplicatorStepSize
		"SourceApplicatorStepSize",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02A1),
		"Applicator Shape Referenced ROI Number", // ApplicatorShapeReferencedROINumber
		"ApplicatorShapeReferencedROINumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02A2),
		"Transfer Tube Number", // TransferTubeNumber
		"TransferTubeNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02A4),
		"Transfer Tube Length", // TransferTubeLength
		"TransferTubeLength",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B0),
		"Channel Shield Sequence", // ChannelShieldSequence
		"ChannelShieldSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B2),
		"Channel Shield Number", // ChannelShieldNumber
		"ChannelShieldNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B3),
		"Channel Shield ID", // ChannelShieldID
		"ChannelShieldID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B4),
		"Channel Shield Name", // ChannelShieldName
		"ChannelShieldName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02B8),
		"Channel Shield Nominal Thickness", // ChannelShieldNominalThickness
		"ChannelShieldNominalThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02BA),
		"Channel Shield Nominal Transmission", // ChannelShieldNominalTransmission
		"ChannelShieldNominalTransmission",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02C8),
		"Final Cumulative Time Weight", // FinalCumulativeTimeWeight
		"FinalCumulativeTimeWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02D0),
		"Brachy Control Point Sequence", // BrachyControlPointSequence
		"BrachyControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02D2),
		"Control Point Relative Position", // ControlPointRelativePosition
		"ControlPointRelativePosition",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02D4),
		"Control Point 3D Position", // ControlPoint3DPosition
		"ControlPoint3DPosition",
		vm.VM3,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02D6),
		"Cumulative Time Weight", // CumulativeTimeWeight
		"CumulativeTimeWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E0),
		"Compensator Divergence", // CompensatorDivergence
		"CompensatorDivergence",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E1),
		"Compensator Mounting Position", // CompensatorMountingPosition
		"CompensatorMountingPosition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E2),
		"Source to Compensator Distance", // SourceToCompensatorDistance
		"SourceToCompensatorDistance",
		vm.VM1_n,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E3),
		"Total Compensator Tray Water-Equivalent Thickness", // TotalCompensatorTrayWaterEquivalentThickness
		"TotalCompensatorTrayWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E4),
		"Isocenter to Compensator Tray Distance", // IsocenterToCompensatorTrayDistance
		"IsocenterToCompensatorTrayDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E5),
		"Compensator Column Offset", // CompensatorColumnOffset
		"CompensatorColumnOffset",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E6),
		"Isocenter to Compensator Distances", // IsocenterToCompensatorDistances
		"IsocenterToCompensatorDistances",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E7),
		"Compensator Relative Stopping Power Ratio", // CompensatorRelativeStoppingPowerRatio
		"CompensatorRelativeStoppingPowerRatio",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02E8),
		"Compensator Milling Tool Diameter", // CompensatorMillingToolDiameter
		"CompensatorMillingToolDiameter",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02EA),
		"Ion Range Compensator Sequence", // IonRangeCompensatorSequence
		"IonRangeCompensatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02EB),
		"Compensator Description", // CompensatorDescription
		"CompensatorDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x02EC),
		"Compensator Surface Representation Flag", // CompensatorSurfaceRepresentationFlag
		"CompensatorSurfaceRepresentationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0302),
		"Radiation Mass Number", // RadiationMassNumber
		"RadiationMassNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0304),
		"Radiation Atomic Number", // RadiationAtomicNumber
		"RadiationAtomicNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0306),
		"Radiation Charge State", // RadiationChargeState
		"RadiationChargeState",
		vm.VM1,
		false,
		vr.SS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0308),
		"Scan Mode", // ScanMode
		"ScanMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0309),
		"Modulated Scan Mode Type", // ModulatedScanModeType
		"ModulatedScanModeType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x030A),
		"Virtual Source-Axis Distances", // VirtualSourceAxisDistances
		"VirtualSourceAxisDistances",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x030C),
		"Snout Sequence", // SnoutSequence
		"SnoutSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x030D),
		"Snout Position", // SnoutPosition
		"SnoutPosition",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x030F),
		"Snout ID", // SnoutID
		"SnoutID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0312),
		"Number of Range Shifters", // NumberOfRangeShifters
		"NumberOfRangeShifters",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0314),
		"Range Shifter Sequence", // RangeShifterSequence
		"RangeShifterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0316),
		"Range Shifter Number", // RangeShifterNumber
		"RangeShifterNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0318),
		"Range Shifter ID", // RangeShifterID
		"RangeShifterID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0320),
		"Range Shifter Type", // RangeShifterType
		"RangeShifterType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0322),
		"Range Shifter Description", // RangeShifterDescription
		"RangeShifterDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0330),
		"Number of Lateral Spreading Devices", // NumberOfLateralSpreadingDevices
		"NumberOfLateralSpreadingDevices",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0332),
		"Lateral Spreading Device Sequence", // LateralSpreadingDeviceSequence
		"LateralSpreadingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0334),
		"Lateral Spreading Device Number", // LateralSpreadingDeviceNumber
		"LateralSpreadingDeviceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0336),
		"Lateral Spreading Device ID", // LateralSpreadingDeviceID
		"LateralSpreadingDeviceID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0338),
		"Lateral Spreading Device Type", // LateralSpreadingDeviceType
		"LateralSpreadingDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x033A),
		"Lateral Spreading Device Description", // LateralSpreadingDeviceDescription
		"LateralSpreadingDeviceDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x033C),
		"Lateral Spreading Device Water Equivalent Thickness", // LateralSpreadingDeviceWaterEquivalentThickness
		"LateralSpreadingDeviceWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0340),
		"Number of Range Modulators", // NumberOfRangeModulators
		"NumberOfRangeModulators",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0342),
		"Range Modulator Sequence", // RangeModulatorSequence
		"RangeModulatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0344),
		"Range Modulator Number", // RangeModulatorNumber
		"RangeModulatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0346),
		"Range Modulator ID", // RangeModulatorID
		"RangeModulatorID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0348),
		"Range Modulator Type", // RangeModulatorType
		"RangeModulatorType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x034A),
		"Range Modulator Description", // RangeModulatorDescription
		"RangeModulatorDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x034C),
		"Beam Current Modulation ID", // BeamCurrentModulationID
		"BeamCurrentModulationID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0350),
		"Patient Support Type", // PatientSupportType
		"PatientSupportType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0352),
		"Patient Support ID", // PatientSupportID
		"PatientSupportID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0354),
		"Patient Support Accessory Code", // PatientSupportAccessoryCode
		"PatientSupportAccessoryCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0355),
		"Tray Accessory Code", // TrayAccessoryCode
		"TrayAccessoryCode",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0356),
		"Fixation Light Azimuthal Angle", // FixationLightAzimuthalAngle
		"FixationLightAzimuthalAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0358),
		"Fixation Light Polar Angle", // FixationLightPolarAngle
		"FixationLightPolarAngle",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x035A),
		"Meterset Rate", // MetersetRate
		"MetersetRate",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0360),
		"Range Shifter Settings Sequence", // RangeShifterSettingsSequence
		"RangeShifterSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0362),
		"Range Shifter Setting", // RangeShifterSetting
		"RangeShifterSetting",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0364),
		"Isocenter to Range Shifter Distance", // IsocenterToRangeShifterDistance
		"IsocenterToRangeShifterDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0366),
		"Range Shifter Water Equivalent Thickness", // RangeShifterWaterEquivalentThickness
		"RangeShifterWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0370),
		"Lateral Spreading Device Settings Sequence", // LateralSpreadingDeviceSettingsSequence
		"LateralSpreadingDeviceSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0372),
		"Lateral Spreading Device Setting", // LateralSpreadingDeviceSetting
		"LateralSpreadingDeviceSetting",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0374),
		"Isocenter to Lateral Spreading Device Distance", // IsocenterToLateralSpreadingDeviceDistance
		"IsocenterToLateralSpreadingDeviceDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0380),
		"Range Modulator Settings Sequence", // RangeModulatorSettingsSequence
		"RangeModulatorSettingsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0382),
		"Range Modulator Gating Start Value", // RangeModulatorGatingStartValue
		"RangeModulatorGatingStartValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0384),
		"Range Modulator Gating Stop Value", // RangeModulatorGatingStopValue
		"RangeModulatorGatingStopValue",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0386),
		"Range Modulator Gating Start Water Equivalent Thickness", // RangeModulatorGatingStartWaterEquivalentThickness
		"RangeModulatorGatingStartWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0388),
		"Range Modulator Gating Stop Water Equivalent Thickness", // RangeModulatorGatingStopWaterEquivalentThickness
		"RangeModulatorGatingStopWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x038A),
		"Isocenter to Range Modulator Distance", // IsocenterToRangeModulatorDistance
		"IsocenterToRangeModulatorDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x038F),
		"Scan Spot Time Offset", // ScanSpotTimeOffset
		"ScanSpotTimeOffset",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0390),
		"Scan Spot Tune ID", // ScanSpotTuneID
		"ScanSpotTuneID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0391),
		"Scan Spot Prescribed Indices", // ScanSpotPrescribedIndices
		"ScanSpotPrescribedIndices",
		vm.VM1_n,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0392),
		"Number of Scan Spot Positions", // NumberOfScanSpotPositions
		"NumberOfScanSpotPositions",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0393),
		"Scan Spot Reordered", // ScanSpotReordered
		"ScanSpotReordered",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0394),
		"Scan Spot Position Map", // ScanSpotPositionMap
		"ScanSpotPositionMap",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0395),
		"Scan Spot Reordering Allowed", // ScanSpotReorderingAllowed
		"ScanSpotReorderingAllowed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0396),
		"Scan Spot Meterset Weights", // ScanSpotMetersetWeights
		"ScanSpotMetersetWeights",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0398),
		"Scanning Spot Size", // ScanningSpotSize
		"ScanningSpotSize",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0399),
		"Scan Spot Sizes Delivered", // ScanSpotSizesDelivered
		"ScanSpotSizesDelivered",
		vm.VM2_2n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x039A),
		"Number of Paintings", // NumberOfPaintings
		"NumberOfPaintings",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x039B),
		"Scan Spot Gantry Angles", // ScanSpotGantryAngles
		"ScanSpotGantryAngles",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x039C),
		"Scan Spot Patient Support Angles", // ScanSpotPatientSupportAngles
		"ScanSpotPatientSupportAngles",
		vm.VM1_n,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A0),
		"Ion Tolerance Table Sequence", // IonToleranceTableSequence
		"IonToleranceTableSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A2),
		"Ion Beam Sequence", // IonBeamSequence
		"IonBeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A4),
		"Ion Beam Limiting Device Sequence", // IonBeamLimitingDeviceSequence
		"IonBeamLimitingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A6),
		"Ion Block Sequence", // IonBlockSequence
		"IonBlockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03A8),
		"Ion Control Point Sequence", // IonControlPointSequence
		"IonControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03AA),
		"Ion Wedge Sequence", // IonWedgeSequence
		"IonWedgeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x03AC),
		"Ion Wedge Position Sequence", // IonWedgePositionSequence
		"IonWedgePositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0401),
		"Referenced Setup Image Sequence", // ReferencedSetupImageSequence
		"ReferencedSetupImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0402),
		"Setup Image Comment", // SetupImageComment
		"SetupImageComment",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0410),
		"Motion Synchronization Sequence", // MotionSynchronizationSequence
		"MotionSynchronizationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0412),
		"Control Point Orientation", // ControlPointOrientation
		"ControlPointOrientation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0420),
		"General Accessory Sequence", // GeneralAccessorySequence
		"GeneralAccessorySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0421),
		"General Accessory ID", // GeneralAccessoryID
		"GeneralAccessoryID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0422),
		"General Accessory Description", // GeneralAccessoryDescription
		"GeneralAccessoryDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0423),
		"General Accessory Type", // GeneralAccessoryType
		"GeneralAccessoryType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0424),
		"General Accessory Number", // GeneralAccessoryNumber
		"GeneralAccessoryNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0425),
		"Source to General Accessory Distance", // SourceToGeneralAccessoryDistance
		"SourceToGeneralAccessoryDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0426),
		"Isocenter to General Accessory Distance", // IsocenterToGeneralAccessoryDistance
		"IsocenterToGeneralAccessoryDistance",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0431),
		"Applicator Geometry Sequence", // ApplicatorGeometrySequence
		"ApplicatorGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0432),
		"Applicator Aperture Shape", // ApplicatorApertureShape
		"ApplicatorApertureShape",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0433),
		"Applicator Opening", // ApplicatorOpening
		"ApplicatorOpening",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0434),
		"Applicator Opening X", // ApplicatorOpeningX
		"ApplicatorOpeningX",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0435),
		"Applicator Opening Y", // ApplicatorOpeningY
		"ApplicatorOpeningY",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0436),
		"Source to Applicator Mounting Position Distance", // SourceToApplicatorMountingPositionDistance
		"SourceToApplicatorMountingPositionDistance",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0440),
		"Number of Block Slab Items", // NumberOfBlockSlabItems
		"NumberOfBlockSlabItems",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0441),
		"Block Slab Sequence", // BlockSlabSequence
		"BlockSlabSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0442),
		"Block Slab Thickness", // BlockSlabThickness
		"BlockSlabThickness",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0443),
		"Block Slab Number", // BlockSlabNumber
		"BlockSlabNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0450),
		"Device Motion Control Sequence", // DeviceMotionControlSequence
		"DeviceMotionControlSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0451),
		"Device Motion Execution Mode", // DeviceMotionExecutionMode
		"DeviceMotionExecutionMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0452),
		"Device Motion Observation Mode", // DeviceMotionObservationMode
		"DeviceMotionObservationMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0453),
		"Device Motion Parameter Code Sequence", // DeviceMotionParameterCodeSequence
		"DeviceMotionParameterCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0501),
		"Distal Depth Fraction", // DistalDepthFraction
		"DistalDepthFraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0502),
		"Distal Depth", // DistalDepth
		"DistalDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0503),
		"Nominal Range Modulation Fractions", // NominalRangeModulationFractions
		"NominalRangeModulationFractions",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0504),
		"Nominal Range Modulated Region Depths", // NominalRangeModulatedRegionDepths
		"NominalRangeModulatedRegionDepths",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0505),
		"Depth Dose Parameters Sequence", // DepthDoseParametersSequence
		"DepthDoseParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0506),
		"Delivered Depth Dose Parameters Sequence", // DeliveredDepthDoseParametersSequence
		"DeliveredDepthDoseParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0507),
		"Delivered Distal Depth Fraction", // DeliveredDistalDepthFraction
		"DeliveredDistalDepthFraction",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0508),
		"Delivered Distal Depth", // DeliveredDistalDepth
		"DeliveredDistalDepth",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0509),
		"Delivered Nominal Range Modulation Fractions", // DeliveredNominalRangeModulationFractions
		"DeliveredNominalRangeModulationFractions",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0510),
		"Delivered Nominal Range Modulated Region Depths", // DeliveredNominalRangeModulatedRegionDepths
		"DeliveredNominalRangeModulatedRegionDepths",
		vm.VM2,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0511),
		"Delivered Reference Dose Definition", // DeliveredReferenceDoseDefinition
		"DeliveredReferenceDoseDefinition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0512),
		"Reference Dose Definition", // ReferenceDoseDefinition
		"ReferenceDoseDefinition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0600),
		"RT Control Point Index", // RTControlPointIndex
		"RTControlPointIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0601),
		"Radiation Generation Mode Index", // RadiationGenerationModeIndex
		"RadiationGenerationModeIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0602),
		"Referenced Defined Device Index", // ReferencedDefinedDeviceIndex
		"ReferencedDefinedDeviceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0603),
		"Radiation Dose Identification Index", // RadiationDoseIdentificationIndex
		"RadiationDoseIdentificationIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0604),
		"Number of RT Control Points", // NumberOfRTControlPoints
		"NumberOfRTControlPoints",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0605),
		"Referenced Radiation Generation Mode Index", // ReferencedRadiationGenerationModeIndex
		"ReferencedRadiationGenerationModeIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0606),
		"Treatment Position Index", // TreatmentPositionIndex
		"TreatmentPositionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0607),
		"Referenced Device Index", // ReferencedDeviceIndex
		"ReferencedDeviceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0608),
		"Treatment Position Group Label", // TreatmentPositionGroupLabel
		"TreatmentPositionGroupLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0609),
		"Treatment Position Group UID", // TreatmentPositionGroupUID
		"TreatmentPositionGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060A),
		"Treatment Position Group Sequence", // TreatmentPositionGroupSequence
		"TreatmentPositionGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060B),
		"Referenced Treatment Position Index", // ReferencedTreatmentPositionIndex
		"ReferencedTreatmentPositionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060C),
		"Referenced Radiation Dose Identification Index", // ReferencedRadiationDoseIdentificationIndex
		"ReferencedRadiationDoseIdentificationIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060D),
		"RT Accessory Holder Water-Equivalent Thickness", // RTAccessoryHolderWaterEquivalentThickness
		"RTAccessoryHolderWaterEquivalentThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060E),
		"Referenced RT Accessory Holder Device Index", // ReferencedRTAccessoryHolderDeviceIndex
		"ReferencedRTAccessoryHolderDeviceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x060F),
		"RT Accessory Holder Slot Existence Flag", // RTAccessoryHolderSlotExistenceFlag
		"RTAccessoryHolderSlotExistenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0610),
		"RT Accessory Holder Slot Sequence", // RTAccessoryHolderSlotSequence
		"RTAccessoryHolderSlotSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0611),
		"RT Accessory Holder Slot ID", // RTAccessoryHolderSlotID
		"RTAccessoryHolderSlotID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0612),
		"RT Accessory Holder Slot Distance", // RTAccessoryHolderSlotDistance
		"RTAccessoryHolderSlotDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0613),
		"RT Accessory Slot Distance", // RTAccessorySlotDistance
		"RTAccessorySlotDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0614),
		"RT Accessory Holder Definition Sequence", // RTAccessoryHolderDefinitionSequence
		"RTAccessoryHolderDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0615),
		"RT Accessory Device Slot ID", // RTAccessoryDeviceSlotID
		"RTAccessoryDeviceSlotID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0616),
		"RT Radiation Sequence", // RTRadiationSequence
		"RTRadiationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0617),
		"Radiation Dose Sequence", // RadiationDoseSequence
		"RadiationDoseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0618),
		"Radiation Dose Identification Sequence", // RadiationDoseIdentificationSequence
		"RadiationDoseIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0619),
		"Radiation Dose Identification Label", // RadiationDoseIdentificationLabel
		"RadiationDoseIdentificationLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061A),
		"Reference Dose Type", // ReferenceDoseType
		"ReferenceDoseType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061B),
		"Primary Dose Value Indicator", // PrimaryDoseValueIndicator
		"PrimaryDoseValueIndicator",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061C),
		"Dose Values Sequence", // DoseValuesSequence
		"DoseValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061D),
		"Dose Value Purpose", // DoseValuePurpose
		"DoseValuePurpose",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061E),
		"Reference Dose Point Coordinates", // ReferenceDosePointCoordinates
		"ReferenceDosePointCoordinates",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x061F),
		"Radiation Dose Values Parameters Sequence", // RadiationDoseValuesParametersSequence
		"RadiationDoseValuesParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0620),
		"Meterset to Dose Mapping Sequence", // MetersetToDoseMappingSequence
		"MetersetToDoseMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0621),
		"Expected In-Vivo Measurement Values Sequence", // ExpectedInVivoMeasurementValuesSequence
		"ExpectedInVivoMeasurementValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0622),
		"Expected In-Vivo Measurement Value Index", // ExpectedInVivoMeasurementValueIndex
		"ExpectedInVivoMeasurementValueIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0623),
		"Radiation Dose In-Vivo Measurement Label", // RadiationDoseInVivoMeasurementLabel
		"RadiationDoseInVivoMeasurementLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0624),
		"Radiation Dose Central Axis Displacement", // RadiationDoseCentralAxisDisplacement
		"RadiationDoseCentralAxisDisplacement",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0625),
		"Radiation Dose Value", // RadiationDoseValue
		"RadiationDoseValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0626),
		"Radiation Dose Source to Skin Distance", // RadiationDoseSourceToSkinDistance
		"RadiationDoseSourceToSkinDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0627),
		"Radiation Dose Measurement Point Coordinates", // RadiationDoseMeasurementPointCoordinates
		"RadiationDoseMeasurementPointCoordinates",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0628),
		"Radiation Dose Source to External Contour Distance", // RadiationDoseSourceToExternalContourDistance
		"RadiationDoseSourceToExternalContourDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0629),
		"RT Tolerance Set Sequence", // RTToleranceSetSequence
		"RTToleranceSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062A),
		"RT Tolerance Set Label", // RTToleranceSetLabel
		"RTToleranceSetLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062B),
		"Attribute Tolerance Values Sequence", // AttributeToleranceValuesSequence
		"AttributeToleranceValuesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062C),
		"Tolerance Value", // ToleranceValue
		"ToleranceValue",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062D),
		"Patient Support Position Tolerance Sequence", // PatientSupportPositionToleranceSequence
		"PatientSupportPositionToleranceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062E),
		"Treatment Time Limit", // TreatmentTimeLimit
		"TreatmentTimeLimit",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x062F),
		"C-Arm Photon-Electron Control Point Sequence", // CArmPhotonElectronControlPointSequence
		"CArmPhotonElectronControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0630),
		"Referenced RT Radiation Sequence", // ReferencedRTRadiationSequence
		"ReferencedRTRadiationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0631),
		"Referenced RT Instance Sequence", // ReferencedRTInstanceSequence
		"ReferencedRTInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0632),
		"Referenced RT Patient Setup Sequence", // ReferencedRTPatientSetupSequence
		"ReferencedRTPatientSetupSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0634),
		"Source to Patient Surface Distance", // SourceToPatientSurfaceDistance
		"SourceToPatientSurfaceDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0635),
		"Treatment Machine Special Mode Code Sequence", // TreatmentMachineSpecialModeCodeSequence
		"TreatmentMachineSpecialModeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0636),
		"Intended Number of Fractions", // IntendedNumberOfFractions
		"IntendedNumberOfFractions",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0637),
		"RT Radiation Set Intent", // RTRadiationSetIntent
		"RTRadiationSetIntent",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0638),
		"RT Radiation Physical and Geometric Content Detail Flag", // RTRadiationPhysicalAndGeometricContentDetailFlag
		"RTRadiationPhysicalAndGeometricContentDetailFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0639),
		"RT Record Flag", // RTRecordFlag
		"RTRecordFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063A),
		"Treatment Device Identification Sequence", // TreatmentDeviceIdentificationSequence
		"TreatmentDeviceIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063B),
		"Referenced RT Physician Intent Sequence", // ReferencedRTPhysicianIntentSequence
		"ReferencedRTPhysicianIntentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063C),
		"Cumulative Meterset", // CumulativeMeterset
		"CumulativeMeterset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063D),
		"Delivery Rate", // DeliveryRate
		"DeliveryRate",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063E),
		"Delivery Rate Unit Sequence", // DeliveryRateUnitSequence
		"DeliveryRateUnitSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x063F),
		"Treatment Position Sequence", // TreatmentPositionSequence
		"TreatmentPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0640),
		"Radiation Source-Axis Distance", // RadiationSourceAxisDistance
		"RadiationSourceAxisDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0641),
		"Number of RT Beam Limiting Devices", // NumberOfRTBeamLimitingDevices
		"NumberOfRTBeamLimitingDevices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0642),
		"RT Beam Limiting Device Proximal Distance", // RTBeamLimitingDeviceProximalDistance
		"RTBeamLimitingDeviceProximalDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0643),
		"RT Beam Limiting Device Distal Distance", // RTBeamLimitingDeviceDistalDistance
		"RTBeamLimitingDeviceDistalDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0644),
		"Parallel RT Beam Delimiter Device Orientation Label Code Sequence", // ParallelRTBeamDelimiterDeviceOrientationLabelCodeSequence
		"ParallelRTBeamDelimiterDeviceOrientationLabelCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0645),
		"Beam Modifier Orientation Angle", // BeamModifierOrientationAngle
		"BeamModifierOrientationAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0646),
		"Fixed RT Beam Delimiter Device Sequence", // FixedRTBeamDelimiterDeviceSequence
		"FixedRTBeamDelimiterDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0647),
		"Parallel RT Beam Delimiter Device Sequence", // ParallelRTBeamDelimiterDeviceSequence
		"ParallelRTBeamDelimiterDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0648),
		"Number of Parallel RT Beam Delimiters", // NumberOfParallelRTBeamDelimiters
		"NumberOfParallelRTBeamDelimiters",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0649),
		"Parallel RT Beam Delimiter Boundaries", // ParallelRTBeamDelimiterBoundaries
		"ParallelRTBeamDelimiterBoundaries",
		vm.VM2_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064A),
		"Parallel RT Beam Delimiter Positions", // ParallelRTBeamDelimiterPositions
		"ParallelRTBeamDelimiterPositions",
		vm.VM2_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064B),
		"RT Beam Limiting Device Offset", // RTBeamLimitingDeviceOffset
		"RTBeamLimitingDeviceOffset",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064C),
		"RT Beam Delimiter Geometry Sequence", // RTBeamDelimiterGeometrySequence
		"RTBeamDelimiterGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064D),
		"RT Beam Limiting Device Definition Sequence", // RTBeamLimitingDeviceDefinitionSequence
		"RTBeamLimitingDeviceDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064E),
		"Parallel RT Beam Delimiter Opening Mode", // ParallelRTBeamDelimiterOpeningMode
		"ParallelRTBeamDelimiterOpeningMode",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x064F),
		"Parallel RT Beam Delimiter Leaf Mounting Side", // ParallelRTBeamDelimiterLeafMountingSide
		"ParallelRTBeamDelimiterLeafMountingSide",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0650),
		"Patient Setup UID", // PatientSetupUID
		"PatientSetupUID",
		vm.VM1,
		true,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0651),
		"Wedge Definition Sequence", // WedgeDefinitionSequence
		"WedgeDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0652),
		"Radiation Beam Wedge Angle", // RadiationBeamWedgeAngle
		"RadiationBeamWedgeAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0653),
		"Radiation Beam Wedge Thin Edge Distance", // RadiationBeamWedgeThinEdgeDistance
		"RadiationBeamWedgeThinEdgeDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0654),
		"Radiation Beam Effective Wedge Angle", // RadiationBeamEffectiveWedgeAngle
		"RadiationBeamEffectiveWedgeAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0655),
		"Number of Wedge Positions", // NumberOfWedgePositions
		"NumberOfWedgePositions",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0656),
		"RT Beam Limiting Device Opening Sequence", // RTBeamLimitingDeviceOpeningSequence
		"RTBeamLimitingDeviceOpeningSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0657),
		"Number of RT Beam Limiting Device Openings", // NumberOfRTBeamLimitingDeviceOpenings
		"NumberOfRTBeamLimitingDeviceOpenings",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0658),
		"Radiation Dosimeter Unit Sequence", // RadiationDosimeterUnitSequence
		"RadiationDosimeterUnitSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0659),
		"RT Device Distance Reference Location Code Sequence", // RTDeviceDistanceReferenceLocationCodeSequence
		"RTDeviceDistanceReferenceLocationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065A),
		"Radiation Device Configuration and Commissioning Key Sequence", // RadiationDeviceConfigurationAndCommissioningKeySequence
		"RadiationDeviceConfigurationAndCommissioningKeySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065B),
		"Patient Support Position Parameter Sequence", // PatientSupportPositionParameterSequence
		"PatientSupportPositionParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065C),
		"Patient Support Position Specification Method", // PatientSupportPositionSpecificationMethod
		"PatientSupportPositionSpecificationMethod",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065D),
		"Patient Support Position Device Parameter Sequence", // PatientSupportPositionDeviceParameterSequence
		"PatientSupportPositionDeviceParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065E),
		"Device Order Index", // DeviceOrderIndex
		"DeviceOrderIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x065F),
		"Patient Support Position Parameter Order Index", // PatientSupportPositionParameterOrderIndex
		"PatientSupportPositionParameterOrderIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0660),
		"Patient Support Position Device Tolerance Sequence", // PatientSupportPositionDeviceToleranceSequence
		"PatientSupportPositionDeviceToleranceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0661),
		"Patient Support Position Tolerance Order Index", // PatientSupportPositionToleranceOrderIndex
		"PatientSupportPositionToleranceOrderIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0662),
		"Compensator Definition Sequence", // CompensatorDefinitionSequence
		"CompensatorDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0663),
		"Compensator Map Orientation", // CompensatorMapOrientation
		"CompensatorMapOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0664),
		"Compensator Proximal Thickness Map", // CompensatorProximalThicknessMap
		"CompensatorProximalThicknessMap",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0665),
		"Compensator Distal Thickness Map", // CompensatorDistalThicknessMap
		"CompensatorDistalThicknessMap",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0666),
		"Compensator Base Plane Offset", // CompensatorBasePlaneOffset
		"CompensatorBasePlaneOffset",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0667),
		"Compensator Shape Fabrication Code Sequence", // CompensatorShapeFabricationCodeSequence
		"CompensatorShapeFabricationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0668),
		"Compensator Shape Sequence", // CompensatorShapeSequence
		"CompensatorShapeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0669),
		"Radiation Beam Compensator Milling Tool Diameter", // RadiationBeamCompensatorMillingToolDiameter
		"RadiationBeamCompensatorMillingToolDiameter",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066A),
		"Block Definition Sequence", // BlockDefinitionSequence
		"BlockDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066B),
		"Block Edge Data", // BlockEdgeData
		"BlockEdgeData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066C),
		"Block Orientation", // BlockOrientation
		"BlockOrientation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066D),
		"Radiation Beam Block Thickness", // RadiationBeamBlockThickness
		"RadiationBeamBlockThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066E),
		"Radiation Beam Block Slab Thickness", // RadiationBeamBlockSlabThickness
		"RadiationBeamBlockSlabThickness",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x066F),
		"Block Edge Data Sequence", // BlockEdgeDataSequence
		"BlockEdgeDataSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0670),
		"Number of RT Accessory Holders", // NumberOfRTAccessoryHolders
		"NumberOfRTAccessoryHolders",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0671),
		"General Accessory Definition Sequence", // GeneralAccessoryDefinitionSequence
		"GeneralAccessoryDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0672),
		"Number of General Accessories", // NumberOfGeneralAccessories
		"NumberOfGeneralAccessories",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0673),
		"Bolus Definition Sequence", // BolusDefinitionSequence
		"BolusDefinitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0674),
		"Number of Boluses", // NumberOfBoluses
		"NumberOfBoluses",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0675),
		"Equipment Frame of Reference UID", // EquipmentFrameOfReferenceUID
		"EquipmentFrameOfReferenceUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0676),
		"Equipment Frame of Reference Description", // EquipmentFrameOfReferenceDescription
		"EquipmentFrameOfReferenceDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0677),
		"Equipment Reference Point Coordinates Sequence", // EquipmentReferencePointCoordinatesSequence
		"EquipmentReferencePointCoordinatesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0678),
		"Equipment Reference Point Code Sequence", // EquipmentReferencePointCodeSequence
		"EquipmentReferencePointCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0679),
		"RT Beam Limiting Device Angle", // RTBeamLimitingDeviceAngle
		"RTBeamLimitingDeviceAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067A),
		"Source Roll Angle", // SourceRollAngle
		"SourceRollAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067B),
		"Radiation GenerationMode Sequence", // RadiationGenerationModeSequence
		"RadiationGenerationModeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067C),
		"Radiation GenerationMode Label", // RadiationGenerationModeLabel
		"RadiationGenerationModeLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067D),
		"Radiation GenerationMode Description", // RadiationGenerationModeDescription
		"RadiationGenerationModeDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067E),
		"Radiation GenerationMode Machine Code Sequence", // RadiationGenerationModeMachineCodeSequence
		"RadiationGenerationModeMachineCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x067F),
		"Radiation Type Code Sequence", // RadiationTypeCodeSequence
		"RadiationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0680),
		"Nominal Energy", // NominalEnergy
		"NominalEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0681),
		"Minimum Nominal Energy", // MinimumNominalEnergy
		"MinimumNominalEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0682),
		"Maximum Nominal Energy", // MaximumNominalEnergy
		"MaximumNominalEnergy",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0683),
		"Radiation Fluence Modifier Code Sequence", // RadiationFluenceModifierCodeSequence
		"RadiationFluenceModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0684),
		"Energy Unit Code Sequence", // EnergyUnitCodeSequence
		"EnergyUnitCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0685),
		"Number of Radiation GenerationModes", // NumberOfRadiationGenerationModes
		"NumberOfRadiationGenerationModes",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0686),
		"Patient Support Devices Sequence", // PatientSupportDevicesSequence
		"PatientSupportDevicesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0687),
		"Number of Patient Support Devices", // NumberOfPatientSupportDevices
		"NumberOfPatientSupportDevices",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0688),
		"RT Beam Modifier Definition Distance", // RTBeamModifierDefinitionDistance
		"RTBeamModifierDefinitionDistance",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0689),
		"Beam Area Limit Sequence", // BeamAreaLimitSequence
		"BeamAreaLimitSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x068A),
		"Referenced RT Prescription Sequence", // ReferencedRTPrescriptionSequence
		"ReferencedRTPrescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x068B),
		"Dose Value Interpretation", // DoseValueInterpretation
		"DoseValueInterpretation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0700),
		"Treatment Session UID", // TreatmentSessionUID
		"TreatmentSessionUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0701),
		"RT Radiation Usage", // RTRadiationUsage
		"RTRadiationUsage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0702),
		"Referenced RT Radiation Set Sequence", // ReferencedRTRadiationSetSequence
		"ReferencedRTRadiationSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0703),
		"Referenced RT Radiation Record Sequence", // ReferencedRTRadiationRecordSequence
		"ReferencedRTRadiationRecordSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0704),
		"RT Radiation Set Delivery Number", // RTRadiationSetDeliveryNumber
		"RTRadiationSetDeliveryNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0705),
		"Clinical Fraction Number", // ClinicalFractionNumber
		"ClinicalFractionNumber",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0706),
		"RT Treatment Fraction Completion Status", // RTTreatmentFractionCompletionStatus
		"RTTreatmentFractionCompletionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0707),
		"RT Radiation Set Usage", // RTRadiationSetUsage
		"RTRadiationSetUsage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0708),
		"Treatment Delivery Continuation Flag", // TreatmentDeliveryContinuationFlag
		"TreatmentDeliveryContinuationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0709),
		"Treatment Record Content Origin", // TreatmentRecordContentOrigin
		"TreatmentRecordContentOrigin",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0714),
		"RT Treatment Termination Status", // RTTreatmentTerminationStatus
		"RTTreatmentTerminationStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0715),
		"RT Treatment Termination Reason Code Sequence", // RTTreatmentTerminationReasonCodeSequence
		"RTTreatmentTerminationReasonCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0716),
		"Machine-Specific Treatment Termination Code Sequence", // MachineSpecificTreatmentTerminationCodeSequence
		"MachineSpecificTreatmentTerminationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0722),
		"RT Radiation Salvage Record Control Point Sequence", // RTRadiationSalvageRecordControlPointSequence
		"RTRadiationSalvageRecordControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0723),
		"Starting Meterset Value Known Flag", // StartingMetersetValueKnownFlag
		"StartingMetersetValueKnownFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0730),
		"Treatment Termination Description", // TreatmentTerminationDescription
		"TreatmentTerminationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0731),
		"Treatment Tolerance Violation Sequence", // TreatmentToleranceViolationSequence
		"TreatmentToleranceViolationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0732),
		"Treatment Tolerance Violation Category", // TreatmentToleranceViolationCategory
		"TreatmentToleranceViolationCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0733),
		"Treatment Tolerance Violation Attribute Sequence", // TreatmentToleranceViolationAttributeSequence
		"TreatmentToleranceViolationAttributeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0734),
		"Treatment Tolerance Violation Description", // TreatmentToleranceViolationDescription
		"TreatmentToleranceViolationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0735),
		"Treatment Tolerance Violation Identification", // TreatmentToleranceViolationIdentification
		"TreatmentToleranceViolationIdentification",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0736),
		"Treatment Tolerance Violation DateTime", // TreatmentToleranceViolationDateTime
		"TreatmentToleranceViolationDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x073A),
		"Recorded RT Control Point DateTime", // RecordedRTControlPointDateTime
		"RecordedRTControlPointDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x073B),
		"Referenced Radiation RT Control Point Index", // ReferencedRadiationRTControlPointIndex
		"ReferencedRadiationRTControlPointIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x073E),
		"Alternate Value Sequence", // AlternateValueSequence
		"AlternateValueSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x073F),
		"Confirmation Sequence", // ConfirmationSequence
		"ConfirmationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0740),
		"Interlock Sequence", // InterlockSequence
		"InterlockSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0741),
		"Interlock DateTime", // InterlockDateTime
		"InterlockDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0742),
		"Interlock Description", // InterlockDescription
		"InterlockDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0743),
		"Interlock Originating Device Sequence", // InterlockOriginatingDeviceSequence
		"InterlockOriginatingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0744),
		"Interlock Code Sequence", // InterlockCodeSequence
		"InterlockCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0745),
		"Interlock Resolution Code Sequence", // InterlockResolutionCodeSequence
		"InterlockResolutionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0746),
		"Interlock Resolution User Sequence", // InterlockResolutionUserSequence
		"InterlockResolutionUserSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0760),
		"Override DateTime", // OverrideDateTime
		"OverrideDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0761),
		"Treatment Tolerance Violation Type Code Sequence", // TreatmentToleranceViolationTypeCodeSequence
		"TreatmentToleranceViolationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0762),
		"Treatment Tolerance Violation Cause Code Sequence", // TreatmentToleranceViolationCauseCodeSequence
		"TreatmentToleranceViolationCauseCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0772),
		"Measured Meterset to Dose Mapping Sequence", // MeasuredMetersetToDoseMappingSequence
		"MeasuredMetersetToDoseMappingSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0773),
		"Referenced Expected In-Vivo Measurement Value Index", // ReferencedExpectedInVivoMeasurementValueIndex
		"ReferencedExpectedInVivoMeasurementValueIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0774),
		"Dose Measurement Device Code Sequence", // DoseMeasurementDeviceCodeSequence
		"DoseMeasurementDeviceCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0780),
		"Additional Parameter Recording Instance Sequence", // AdditionalParameterRecordingInstanceSequence
		"AdditionalParameterRecordingInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0783),
		"Interlock Origin Description", // InterlockOriginDescription
		"InterlockOriginDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0784),
		"RT Patient Position Scope Sequence", // RTPatientPositionScopeSequence
		"RTPatientPositionScopeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0785),
		"Referenced Treatment Position Group UID", // ReferencedTreatmentPositionGroupUID
		"ReferencedTreatmentPositionGroupUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0786),
		"Radiation Order Index", // RadiationOrderIndex
		"RadiationOrderIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0787),
		"Omitted Radiation Sequence", // OmittedRadiationSequence
		"OmittedRadiationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0788),
		"Reason for Omission Code Sequence", // ReasonForOmissionCodeSequence
		"ReasonForOmissionCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0789),
		"RT Delivery Start Patient Position Sequence", // RTDeliveryStartPatientPositionSequence
		"RTDeliveryStartPatientPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078A),
		"RT Treatment Preparation Patient Position Sequence", // RTTreatmentPreparationPatientPositionSequence
		"RTTreatmentPreparationPatientPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078B),
		"Referenced RT Treatment Preparation Sequence", // ReferencedRTTreatmentPreparationSequence
		"ReferencedRTTreatmentPreparationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078C),
		"Referenced Patient Setup Photo Sequence", // ReferencedPatientSetupPhotoSequence
		"ReferencedPatientSetupPhotoSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078D),
		"Patient Treatment Preparation Method Code Sequence", // PatientTreatmentPreparationMethodCodeSequence
		"PatientTreatmentPreparationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078E),
		"Patient Treatment Preparation Procedure Parameter Description", // PatientTreatmentPreparationProcedureParameterDescription
		"PatientTreatmentPreparationProcedureParameterDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x078F),
		"Patient Treatment Preparation Device Sequence", // PatientTreatmentPreparationDeviceSequence
		"PatientTreatmentPreparationDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0790),
		"Patient Treatment Preparation Procedure Sequence ", // PatientTreatmentPreparationProcedureSequence
		"PatientTreatmentPreparationProcedureSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0791),
		"Patient Treatment Preparation Procedure Code Sequence", // PatientTreatmentPreparationProcedureCodeSequence
		"PatientTreatmentPreparationProcedureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0792),
		"Patient Treatment Preparation Method Description", // PatientTreatmentPreparationMethodDescription
		"PatientTreatmentPreparationMethodDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0793),
		"Patient Treatment Preparation Procedure Parameter Sequence", // PatientTreatmentPreparationProcedureParameterSequence
		"PatientTreatmentPreparationProcedureParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0794),
		"Patient Setup Photo Description", // PatientSetupPhotoDescription
		"PatientSetupPhotoDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0795),
		"Patient Treatment Preparation Procedure Index", // PatientTreatmentPreparationProcedureIndex
		"PatientTreatmentPreparationProcedureIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0796),
		"Referenced Patient Setup Procedure Index", // ReferencedPatientSetupProcedureIndex
		"ReferencedPatientSetupProcedureIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0797),
		"RT Radiation Task Sequence", // RTRadiationTaskSequence
		"RTRadiationTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0798),
		"RT Patient Position Displacement Sequence", // RTPatientPositionDisplacementSequence
		"RTPatientPositionDisplacementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x0799),
		"RT Patient Position Sequence", // RTPatientPositionSequence
		"RTPatientPositionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079A),
		"Displacement Reference Label", // DisplacementReferenceLabel
		"DisplacementReferenceLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079B),
		"Displacement Matrix", // DisplacementMatrix
		"DisplacementMatrix",
		vm.VM16,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079C),
		"Patient Support Displacement Sequence", // PatientSupportDisplacementSequence
		"PatientSupportDisplacementSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079D),
		"Displacement Reference Location Code Sequence", // DisplacementReferenceLocationCodeSequence
		"DisplacementReferenceLocationCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079E),
		"RT Radiation Set Delivery Usage", // RTRadiationSetDeliveryUsage
		"RTRadiationSetDeliveryUsage",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x079F),
		"Patient Treatment Preparation Sequence", // PatientTreatmentPreparationSequence
		"PatientTreatmentPreparationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x07A0),
		"Patient to Equipment Relationship Sequence", // PatientToEquipmentRelationshipSequence
		"PatientToEquipmentRelationshipSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300A, 0x07A1),
		"Imaging Equipment to Treatment Delivery Device Relationship Sequence", // ImagingEquipmentToTreatmentDeliveryDeviceRelationshipSequence
		"ImagingEquipmentToTreatmentDeliveryDeviceRelationshipSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0002),
		"Referenced RT Plan Sequence", // ReferencedRTPlanSequence
		"ReferencedRTPlanSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0004),
		"Referenced Beam Sequence", // ReferencedBeamSequence
		"ReferencedBeamSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0006),
		"Referenced Beam Number", // ReferencedBeamNumber
		"ReferencedBeamNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0007),
		"Referenced Reference Image Number", // ReferencedReferenceImageNumber
		"ReferencedReferenceImageNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0008),
		"Start Cumulative Meterset Weight", // StartCumulativeMetersetWeight
		"StartCumulativeMetersetWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0009),
		"End Cumulative Meterset Weight", // EndCumulativeMetersetWeight
		"EndCumulativeMetersetWeight",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x000A),
		"Referenced Brachy Application Setup Sequence", // ReferencedBrachyApplicationSetupSequence
		"ReferencedBrachyApplicationSetupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x000C),
		"Referenced Brachy Application Setup Number", // ReferencedBrachyApplicationSetupNumber
		"ReferencedBrachyApplicationSetupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x000E),
		"Referenced Source Number", // ReferencedSourceNumber
		"ReferencedSourceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0020),
		"Referenced Fraction Group Sequence", // ReferencedFractionGroupSequence
		"ReferencedFractionGroupSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0022),
		"Referenced Fraction Group Number", // ReferencedFractionGroupNumber
		"ReferencedFractionGroupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0040),
		"Referenced Verification Image Sequence", // ReferencedVerificationImageSequence
		"ReferencedVerificationImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0042),
		"Referenced Reference Image Sequence", // ReferencedReferenceImageSequence
		"ReferencedReferenceImageSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0050),
		"Referenced Dose Reference Sequence", // ReferencedDoseReferenceSequence
		"ReferencedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0051),
		"Referenced Dose Reference Number", // ReferencedDoseReferenceNumber
		"ReferencedDoseReferenceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0055),
		"Brachy Referenced Dose Reference Sequence", // BrachyReferencedDoseReferenceSequence
		"BrachyReferencedDoseReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0060),
		"Referenced Structure Set Sequence", // ReferencedStructureSetSequence
		"ReferencedStructureSetSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x006A),
		"Referenced Patient Setup Number", // ReferencedPatientSetupNumber
		"ReferencedPatientSetupNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0080),
		"Referenced Dose Sequence", // ReferencedDoseSequence
		"ReferencedDoseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00A0),
		"Referenced Tolerance Table Number", // ReferencedToleranceTableNumber
		"ReferencedToleranceTableNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00B0),
		"Referenced Bolus Sequence", // ReferencedBolusSequence
		"ReferencedBolusSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00C0),
		"Referenced Wedge Number", // ReferencedWedgeNumber
		"ReferencedWedgeNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00D0),
		"Referenced Compensator Number", // ReferencedCompensatorNumber
		"ReferencedCompensatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00E0),
		"Referenced Block Number", // ReferencedBlockNumber
		"ReferencedBlockNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00F0),
		"Referenced Control Point Index", // ReferencedControlPointIndex
		"ReferencedControlPointIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00F2),
		"Referenced Control Point Sequence", // ReferencedControlPointSequence
		"ReferencedControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00F4),
		"Referenced Start Control Point Index", // ReferencedStartControlPointIndex
		"ReferencedStartControlPointIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x00F6),
		"Referenced Stop Control Point Index", // ReferencedStopControlPointIndex
		"ReferencedStopControlPointIndex",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0100),
		"Referenced Range Shifter Number", // ReferencedRangeShifterNumber
		"ReferencedRangeShifterNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0102),
		"Referenced Lateral Spreading Device Number", // ReferencedLateralSpreadingDeviceNumber
		"ReferencedLateralSpreadingDeviceNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0104),
		"Referenced Range Modulator Number", // ReferencedRangeModulatorNumber
		"ReferencedRangeModulatorNumber",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0111),
		"Omitted Beam Task Sequence", // OmittedBeamTaskSequence
		"OmittedBeamTaskSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0112),
		"Reason for Omission", // ReasonForOmission
		"ReasonForOmission",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0113),
		"Reason for Omission Description", // ReasonForOmissionDescription
		"ReasonForOmissionDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0114),
		"Prescription Overview Sequence", // PrescriptionOverviewSequence
		"PrescriptionOverviewSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0115),
		"Total Prescription Dose", // TotalPrescriptionDose
		"TotalPrescriptionDose",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0116),
		"Plan Overview Sequence", // PlanOverviewSequence
		"PlanOverviewSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0117),
		"Plan Overview Index", // PlanOverviewIndex
		"PlanOverviewIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0118),
		"Referenced Plan Overview Index", // ReferencedPlanOverviewIndex
		"ReferencedPlanOverviewIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0119),
		"Number of Fractions Included", // NumberOfFractionsIncluded
		"NumberOfFractionsIncluded",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0120),
		"Dose Calibration Conditions Sequence", // DoseCalibrationConditionsSequence
		"DoseCalibrationConditionsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0121),
		"Absorbed Dose to Meterset Ratio", // AbsorbedDoseToMetersetRatio
		"AbsorbedDoseToMetersetRatio",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0122),
		"Delineated Radiation Field Size", // DelineatedRadiationFieldSize
		"DelineatedRadiationFieldSize",
		vm.VM2,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0123),
		"Dose Calibration Conditions Verified Flag", // DoseCalibrationConditionsVerifiedFlag
		"DoseCalibrationConditionsVerifiedFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0124),
		"Calibration Reference Point Depth", // CalibrationReferencePointDepth
		"CalibrationReferencePointDepth",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0125),
		"Gating Beam Hold Transition Sequence", // GatingBeamHoldTransitionSequence
		"GatingBeamHoldTransitionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0126),
		"Beam Hold Transition", // BeamHoldTransition
		"BeamHoldTransition",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0127),
		"Beam Hold Transition DateTime", // BeamHoldTransitionDateTime
		"BeamHoldTransitionDateTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0128),
		"Beam Hold Originating Device Sequence", // BeamHoldOriginatingDeviceSequence
		"BeamHoldOriginatingDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x300C, 0x0129),
		"Beam Hold Transition Trigger Source", // BeamHoldTransitionTriggerSource
		"BeamHoldTransitionTriggerSource",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300E, 0x0002),
		"Approval Status", // ApprovalStatus
		"ApprovalStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x300E, 0x0004),
		"Review Date", // ReviewDate
		"ReviewDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x300E, 0x0005),
		"Review Time", // ReviewTime
		"ReviewTime",
		vm.VM1,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x300E, 0x0008),
		"Reviewer Name", // ReviewerName
		"ReviewerName",
		vm.VM1,
		false,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0001),
		"Radiobiological Dose Effect Sequence", // RadiobiologicalDoseEffectSequence
		"RadiobiologicalDoseEffectSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0002),
		"Radiobiological Dose Effect Flag", // RadiobiologicalDoseEffectFlag
		"RadiobiologicalDoseEffectFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0003),
		"Effective Dose Calculation Method Category Code Sequence", // EffectiveDoseCalculationMethodCategoryCodeSequence
		"EffectiveDoseCalculationMethodCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0004),
		"Effective Dose Calculation Method Code Sequence", // EffectiveDoseCalculationMethodCodeSequence
		"EffectiveDoseCalculationMethodCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0005),
		"Effective Dose Calculation Method Description", // EffectiveDoseCalculationMethodDescription
		"EffectiveDoseCalculationMethodDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0006),
		"Conceptual Volume UID", // ConceptualVolumeUID
		"ConceptualVolumeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0007),
		"Originating SOP Instance Reference Sequence", // OriginatingSOPInstanceReferenceSequence
		"OriginatingSOPInstanceReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0008),
		"Conceptual Volume Constituent Sequence", // ConceptualVolumeConstituentSequence
		"ConceptualVolumeConstituentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0009),
		"Equivalent Conceptual Volume Instance Reference Sequence", // EquivalentConceptualVolumeInstanceReferenceSequence
		"EquivalentConceptualVolumeInstanceReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000A),
		"Equivalent Conceptual Volumes Sequence", // EquivalentConceptualVolumesSequence
		"EquivalentConceptualVolumesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000B),
		"Referenced Conceptual Volume UID", // ReferencedConceptualVolumeUID
		"ReferencedConceptualVolumeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000C),
		"Conceptual Volume Combination Expression", // ConceptualVolumeCombinationExpression
		"ConceptualVolumeCombinationExpression",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000D),
		"Conceptual Volume Constituent Index", // ConceptualVolumeConstituentIndex
		"ConceptualVolumeConstituentIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000E),
		"Conceptual Volume Combination Flag", // ConceptualVolumeCombinationFlag
		"ConceptualVolumeCombinationFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x000F),
		"Conceptual Volume Combination Description", // ConceptualVolumeCombinationDescription
		"ConceptualVolumeCombinationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0010),
		"Conceptual Volume Segmentation Defined Flag", // ConceptualVolumeSegmentationDefinedFlag
		"ConceptualVolumeSegmentationDefinedFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0011),
		"Conceptual Volume Segmentation Reference Sequence", // ConceptualVolumeSegmentationReferenceSequence
		"ConceptualVolumeSegmentationReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0012),
		"Conceptual Volume Constituent Segmentation Reference Sequence", // ConceptualVolumeConstituentSegmentationReferenceSequence
		"ConceptualVolumeConstituentSegmentationReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0013),
		"Constituent Conceptual Volume UID", // ConstituentConceptualVolumeUID
		"ConstituentConceptualVolumeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0014),
		"Derivation Conceptual Volume Sequence", // DerivationConceptualVolumeSequence
		"DerivationConceptualVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0015),
		"Source Conceptual Volume UID", // SourceConceptualVolumeUID
		"SourceConceptualVolumeUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0016),
		"Conceptual Volume Derivation Algorithm Sequence", // ConceptualVolumeDerivationAlgorithmSequence
		"ConceptualVolumeDerivationAlgorithmSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0017),
		"Conceptual Volume Description", // ConceptualVolumeDescription
		"ConceptualVolumeDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0018),
		"Source Conceptual Volume Sequence", // SourceConceptualVolumeSequence
		"SourceConceptualVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0019),
		"Author Identification Sequence", // AuthorIdentificationSequence
		"AuthorIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001A),
		"Manufacturer's Model Version", // ManufacturerModelVersion
		"ManufacturerModelVersion",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001B),
		"Device Alternate Identifier", // DeviceAlternateIdentifier
		"DeviceAlternateIdentifier",
		vm.VM1,
		false,
		vr.UC,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001C),
		"Device Alternate Identifier Type", // DeviceAlternateIdentifierType
		"DeviceAlternateIdentifierType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001D),
		"Device Alternate Identifier Format", // DeviceAlternateIdentifierFormat
		"DeviceAlternateIdentifierFormat",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001E),
		"Segmentation Creation Template Label", // SegmentationCreationTemplateLabel
		"SegmentationCreationTemplateLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x001F),
		"Segmentation Template UID", // SegmentationTemplateUID
		"SegmentationTemplateUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0020),
		"Referenced Segment Reference Index", // ReferencedSegmentReferenceIndex
		"ReferencedSegmentReferenceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0021),
		"Segment Reference Sequence", // SegmentReferenceSequence
		"SegmentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0022),
		"Segment Reference Index", // SegmentReferenceIndex
		"SegmentReferenceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0023),
		"Direct Segment Reference Sequence", // DirectSegmentReferenceSequence
		"DirectSegmentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0024),
		"Combination Segment Reference Sequence", // CombinationSegmentReferenceSequence
		"CombinationSegmentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0025),
		"Conceptual Volume Sequence", // ConceptualVolumeSequence
		"ConceptualVolumeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0026),
		"Segmented RT Accessory Device Sequence", // SegmentedRTAccessoryDeviceSequence
		"SegmentedRTAccessoryDeviceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0027),
		"Segment Characteristics Sequence", // SegmentCharacteristicsSequence
		"SegmentCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0028),
		"Related Segment Characteristics Sequence", // RelatedSegmentCharacteristicsSequence
		"RelatedSegmentCharacteristicsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0029),
		"Segment Characteristics Precedence", // SegmentCharacteristicsPrecedence
		"SegmentCharacteristicsPrecedence",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002A),
		"RT Segment Annotation Sequence", // RTSegmentAnnotationSequence
		"RTSegmentAnnotationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002B),
		"Segment Annotation Category Code Sequence", // SegmentAnnotationCategoryCodeSequence
		"SegmentAnnotationCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002C),
		"Segment Annotation Type Code Sequence", // SegmentAnnotationTypeCodeSequence
		"SegmentAnnotationTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002D),
		"Device Label", // DeviceLabel
		"DeviceLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002E),
		"Device Type Code Sequence", // DeviceTypeCodeSequence
		"DeviceTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x002F),
		"Segment Annotation Type Modifier Code Sequence", // SegmentAnnotationTypeModifierCodeSequence
		"SegmentAnnotationTypeModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0030),
		"Patient Equipment Relationship Code Sequence", // PatientEquipmentRelationshipCodeSequence
		"PatientEquipmentRelationshipCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0031),
		"Referenced Fiducials UID", // ReferencedFiducialsUID
		"ReferencedFiducialsUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0032),
		"Patient Treatment Orientation Sequence", // PatientTreatmentOrientationSequence
		"PatientTreatmentOrientationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0033),
		"User Content Label", // UserContentLabel
		"UserContentLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0034),
		"User Content Long Label", // UserContentLongLabel
		"UserContentLongLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0035),
		"Entity Label", // EntityLabel
		"EntityLabel",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0036),
		"Entity Name", // EntityName
		"EntityName",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0037),
		"Entity Description", // EntityDescription
		"EntityDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0038),
		"Entity Long Label", // EntityLongLabel
		"EntityLongLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0039),
		"Device Index", // DeviceIndex
		"DeviceIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003A),
		"RT Treatment Phase Index", // RTTreatmentPhaseIndex
		"RTTreatmentPhaseIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003B),
		"RT Treatment Phase UID", // RTTreatmentPhaseUID
		"RTTreatmentPhaseUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003C),
		"RT Prescription Index", // RTPrescriptionIndex
		"RTPrescriptionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003D),
		"RT Segment Annotation Index", // RTSegmentAnnotationIndex
		"RTSegmentAnnotationIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003E),
		"Basis RT Treatment Phase Index", // BasisRTTreatmentPhaseIndex
		"BasisRTTreatmentPhaseIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x003F),
		"Related RT Treatment Phase Index", // RelatedRTTreatmentPhaseIndex
		"RelatedRTTreatmentPhaseIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0040),
		"Referenced RT Treatment Phase Index", // ReferencedRTTreatmentPhaseIndex
		"ReferencedRTTreatmentPhaseIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0041),
		"Referenced RT Prescription Index", // ReferencedRTPrescriptionIndex
		"ReferencedRTPrescriptionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0042),
		"Referenced Parent RT Prescription Index", // ReferencedParentRTPrescriptionIndex
		"ReferencedParentRTPrescriptionIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0043),
		"Manufacturer's Device Identifier", // ManufacturerDeviceIdentifier
		"ManufacturerDeviceIdentifier",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0044),
		"Instance-Level Referenced Performed Procedure Step Sequence", // InstanceLevelReferencedPerformedProcedureStepSequence
		"InstanceLevelReferencedPerformedProcedureStepSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0045),
		"RT Treatment Phase Intent Presence Flag", // RTTreatmentPhaseIntentPresenceFlag
		"RTTreatmentPhaseIntentPresenceFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0046),
		"Radiotherapy Treatment Type", // RadiotherapyTreatmentType
		"RadiotherapyTreatmentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0047),
		"Teletherapy Radiation Type", // TeletherapyRadiationType
		"TeletherapyRadiationType",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0048),
		"Brachytherapy Source Type", // BrachytherapySourceType
		"BrachytherapySourceType",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0049),
		"Referenced RT Treatment Phase Sequence", // ReferencedRTTreatmentPhaseSequence
		"ReferencedRTTreatmentPhaseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004A),
		"Referenced Direct Segment Instance Sequence", // ReferencedDirectSegmentInstanceSequence
		"ReferencedDirectSegmentInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004B),
		"Intended RT Treatment Phase Sequence", // IntendedRTTreatmentPhaseSequence
		"IntendedRTTreatmentPhaseSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004C),
		"Intended Phase Start Date", // IntendedPhaseStartDate
		"IntendedPhaseStartDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004D),
		"Intended Phase End Date", // IntendedPhaseEndDate
		"IntendedPhaseEndDate",
		vm.VM1,
		false,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004E),
		"RT Treatment Phase Interval Sequence", // RTTreatmentPhaseIntervalSequence
		"RTTreatmentPhaseIntervalSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x004F),
		"Temporal Relationship Interval Anchor", // TemporalRelationshipIntervalAnchor
		"TemporalRelationshipIntervalAnchor",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0050),
		"Minimum Number of Interval Days", // MinimumNumberOfIntervalDays
		"MinimumNumberOfIntervalDays",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0051),
		"Maximum Number of Interval Days", // MaximumNumberOfIntervalDays
		"MaximumNumberOfIntervalDays",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0052),
		"Pertinent SOP Classes in Study", // PertinentSOPClassesInStudy
		"PertinentSOPClassesInStudy",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0053),
		"Pertinent SOP Classes in Series", // PertinentSOPClassesInSeries
		"PertinentSOPClassesInSeries",
		vm.VM1_n,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0054),
		"RT Prescription Label", // RTPrescriptionLabel
		"RTPrescriptionLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0055),
		"RT Physician Intent Predecessor Sequence", // RTPhysicianIntentPredecessorSequence
		"RTPhysicianIntentPredecessorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0056),
		"RT Treatment Approach Label", // RTTreatmentApproachLabel
		"RTTreatmentApproachLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0057),
		"RT Physician Intent Sequence", // RTPhysicianIntentSequence
		"RTPhysicianIntentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0058),
		"RT Physician Intent Index", // RTPhysicianIntentIndex
		"RTPhysicianIntentIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0059),
		"RT Treatment Intent Type", // RTTreatmentIntentType
		"RTTreatmentIntentType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005A),
		"RT Physician Intent Narrative", // RTPhysicianIntentNarrative
		"RTPhysicianIntentNarrative",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005B),
		"RT Protocol Code Sequence", // RTProtocolCodeSequence
		"RTProtocolCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005C),
		"Reason for Superseding", // ReasonForSuperseding
		"ReasonForSuperseding",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005D),
		"RT Diagnosis Code Sequence", // RTDiagnosisCodeSequence
		"RTDiagnosisCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005E),
		"Referenced RT Physician Intent Index", // ReferencedRTPhysicianIntentIndex
		"ReferencedRTPhysicianIntentIndex",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x005F),
		"RT Physician Intent Input Instance Sequence", // RTPhysicianIntentInputInstanceSequence
		"RTPhysicianIntentInputInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0060),
		"RT Anatomic Prescription Sequence", // RTAnatomicPrescriptionSequence
		"RTAnatomicPrescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0061),
		"Prior Treatment Dose Description", // PriorTreatmentDoseDescription
		"PriorTreatmentDoseDescription",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0062),
		"Prior Treatment Reference Sequence", // PriorTreatmentReferenceSequence
		"PriorTreatmentReferenceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0063),
		"Dosimetric Objective Evaluation Scope", // DosimetricObjectiveEvaluationScope
		"DosimetricObjectiveEvaluationScope",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0064),
		"Therapeutic Role Category Code Sequence", // TherapeuticRoleCategoryCodeSequence
		"TherapeuticRoleCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0065),
		"Therapeutic Role Type Code Sequence", // TherapeuticRoleTypeCodeSequence
		"TherapeuticRoleTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0066),
		"Conceptual Volume Optimization Precedence", // ConceptualVolumeOptimizationPrecedence
		"ConceptualVolumeOptimizationPrecedence",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0067),
		"Conceptual Volume Category Code Sequence", // ConceptualVolumeCategoryCodeSequence
		"ConceptualVolumeCategoryCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0068),
		"Conceptual Volume Blocking Constraint", // ConceptualVolumeBlockingConstraint
		"ConceptualVolumeBlockingConstraint",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0069),
		"Conceptual Volume Type Code Sequence", // ConceptualVolumeTypeCodeSequence
		"ConceptualVolumeTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006A),
		"Conceptual Volume Type Modifier Code Sequence", // ConceptualVolumeTypeModifierCodeSequence
		"ConceptualVolumeTypeModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006B),
		"RT Prescription Sequence", // RTPrescriptionSequence
		"RTPrescriptionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006C),
		"Dosimetric Objective Sequence", // DosimetricObjectiveSequence
		"DosimetricObjectiveSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006D),
		"Dosimetric Objective Type Code Sequence", // DosimetricObjectiveTypeCodeSequence
		"DosimetricObjectiveTypeCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006E),
		"Dosimetric Objective UID", // DosimetricObjectiveUID
		"DosimetricObjectiveUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x006F),
		"Referenced Dosimetric Objective UID", // ReferencedDosimetricObjectiveUID
		"ReferencedDosimetricObjectiveUID",
		vm.VM1,
		false,
		vr.UI,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0070),
		"Dosimetric Objective Parameter Sequence", // DosimetricObjectiveParameterSequence
		"DosimetricObjectiveParameterSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0071),
		"Referenced Dosimetric Objectives Sequence", // ReferencedDosimetricObjectivesSequence
		"ReferencedDosimetricObjectivesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0073),
		"Absolute Dosimetric Objective Flag", // AbsoluteDosimetricObjectiveFlag
		"AbsoluteDosimetricObjectiveFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0074),
		"Dosimetric Objective Weight", // DosimetricObjectiveWeight
		"DosimetricObjectiveWeight",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0075),
		"Dosimetric Objective Purpose", // DosimetricObjectivePurpose
		"DosimetricObjectivePurpose",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0076),
		"Planning Input Information Sequence", // PlanningInputInformationSequence
		"PlanningInputInformationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0077),
		"Treatment Site", // TreatmentSite
		"TreatmentSite",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0078),
		"Treatment Site Code Sequence", // TreatmentSiteCodeSequence
		"TreatmentSiteCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0079),
		"Fraction Pattern Sequence", // FractionPatternSequence
		"FractionPatternSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007A),
		"Treatment Technique Notes", // TreatmentTechniqueNotes
		"TreatmentTechniqueNotes",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007B),
		"Prescription Notes", // PrescriptionNotes
		"PrescriptionNotes",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007C),
		"Number of Interval Fractions", // NumberOfIntervalFractions
		"NumberOfIntervalFractions",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007D),
		"Number of Fractions", // NumberOfFractions
		"NumberOfFractions",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007E),
		"Intended Delivery Duration", // IntendedDeliveryDuration
		"IntendedDeliveryDuration",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x007F),
		"Fractionation Notes", // FractionationNotes
		"FractionationNotes",
		vm.VM1,
		false,
		vr.UT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0080),
		"RT Treatment Technique Code Sequence", // RTTreatmentTechniqueCodeSequence
		"RTTreatmentTechniqueCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0081),
		"Prescription Notes Sequence", // PrescriptionNotesSequence
		"PrescriptionNotesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0082),
		"Fraction-Based Relationship Sequence", // FractionBasedRelationshipSequence
		"FractionBasedRelationshipSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0083),
		"Fraction-Based Relationship Interval Anchor", // FractionBasedRelationshipIntervalAnchor
		"FractionBasedRelationshipIntervalAnchor",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0084),
		"Minimum Hours between Fractions", // MinimumHoursBetweenFractions
		"MinimumHoursBetweenFractions",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0085),
		"Intended Fraction Start Time", // IntendedFractionStartTime
		"IntendedFractionStartTime",
		vm.VM1_n,
		false,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0086),
		"Intended Start Day of Week", // IntendedStartDayOfWeek
		"IntendedStartDayOfWeek",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0087),
		"Weekday Fraction Pattern Sequence", // WeekdayFractionPatternSequence
		"WeekdayFractionPatternSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0088),
		"Delivery Time Structure Code Sequence", // DeliveryTimeStructureCodeSequence
		"DeliveryTimeStructureCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0089),
		"Treatment Site Modifier Code Sequence", // TreatmentSiteModifierCodeSequence
		"TreatmentSiteModifierCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0090),
		"Robotic Base Location Indicator", // RoboticBaseLocationIndicator
		"RoboticBaseLocationIndicator",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0091),
		"Robotic Path Node Set Code Sequence", // RoboticPathNodeSetCodeSequence
		"RoboticPathNodeSetCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0092),
		"Robotic Node Identifier", // RoboticNodeIdentifier
		"RoboticNodeIdentifier",
		vm.VM1,
		false,
		vr.UL,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0093),
		"RT Treatment Source Coordinates", // RTTreatmentSourceCoordinates
		"RTTreatmentSourceCoordinates",
		vm.VM3,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0094),
		"Radiation Source Coordinate SystemYaw Angle", // RadiationSourceCoordinateSystemYawAngle
		"RadiationSourceCoordinateSystemYawAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0095),
		"Radiation Source Coordinate SystemRoll Angle", // RadiationSourceCoordinateSystemRollAngle
		"RadiationSourceCoordinateSystemRollAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0096),
		"Radiation Source Coordinate System Pitch Angle", // RadiationSourceCoordinateSystemPitchAngle
		"RadiationSourceCoordinateSystemPitchAngle",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0097),
		"Robotic Path Control Point Sequence", // RoboticPathControlPointSequence
		"RoboticPathControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0098),
		"Tomotherapeutic Control Point Sequence", // TomotherapeuticControlPointSequence
		"TomotherapeuticControlPointSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x0099),
		"Tomotherapeutic Leaf Open Durations", // TomotherapeuticLeafOpenDurations
		"TomotherapeuticLeafOpenDurations",
		vm.VM1_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x009A),
		"Tomotherapeutic Leaf Initial Closed Durations", // TomotherapeuticLeafInitialClosedDurations
		"TomotherapeuticLeafInitialClosedDurations",
		vm.VM1_n,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x3010, 0x00A0),
		"Conceptual Volume Identification Sequence", // ConceptualVolumeIdentificationSequence
		"ConceptualVolumeIdentificationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4000, 0x0010),
		"Arbitrary", // Arbitrary
		"Arbitrary",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4000, 0x4000),
		"Text Comments", // TextComments
		"TextComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0040),
		"Results ID", // ResultsID
		"ResultsID",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0042),
		"Results ID Issuer", // ResultsIDIssuer
		"ResultsIDIssuer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0050),
		"Referenced Interpretation Sequence", // ReferencedInterpretationSequence
		"ReferencedInterpretationSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x00FF),
		"Report Production Status (Trial)", // ReportProductionStatusTrial
		"ReportProductionStatusTrial",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0100),
		"Interpretation Recorded Date", // InterpretationRecordedDate
		"InterpretationRecordedDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0101),
		"Interpretation Recorded Time", // InterpretationRecordedTime
		"InterpretationRecordedTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0102),
		"Interpretation Recorder", // InterpretationRecorder
		"InterpretationRecorder",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0103),
		"Reference to Recorded Sound", // ReferenceToRecordedSound
		"ReferenceToRecordedSound",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0108),
		"Interpretation Transcription Date", // InterpretationTranscriptionDate
		"InterpretationTranscriptionDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0109),
		"Interpretation Transcription Time", // InterpretationTranscriptionTime
		"InterpretationTranscriptionTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x010A),
		"Interpretation Transcriber", // InterpretationTranscriber
		"InterpretationTranscriber",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x010B),
		"Interpretation Text", // InterpretationText
		"InterpretationText",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x010C),
		"Interpretation Author", // InterpretationAuthor
		"InterpretationAuthor",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0111),
		"Interpretation Approver Sequence", // InterpretationApproverSequence
		"InterpretationApproverSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0112),
		"Interpretation Approval Date", // InterpretationApprovalDate
		"InterpretationApprovalDate",
		vm.VM1,
		true,
		vr.DA,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0113),
		"Interpretation Approval Time", // InterpretationApprovalTime
		"InterpretationApprovalTime",
		vm.VM1,
		true,
		vr.TM,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0114),
		"Physician Approving Interpretation", // PhysicianApprovingInterpretation
		"PhysicianApprovingInterpretation",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0115),
		"Interpretation Diagnosis Description", // InterpretationDiagnosisDescription
		"InterpretationDiagnosisDescription",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0117),
		"Interpretation Diagnosis Code Sequence", // InterpretationDiagnosisCodeSequence
		"InterpretationDiagnosisCodeSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0118),
		"Results Distribution List Sequence", // ResultsDistributionListSequence
		"ResultsDistributionListSequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0119),
		"Distribution Name", // DistributionName
		"DistributionName",
		vm.VM1,
		true,
		vr.PN,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x011A),
		"Distribution Address", // DistributionAddress
		"DistributionAddress",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0200),
		"Interpretation ID", // InterpretationID
		"InterpretationID",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0202),
		"Interpretation ID Issuer", // InterpretationIDIssuer
		"InterpretationIDIssuer",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0210),
		"Interpretation Type ID", // InterpretationTypeID
		"InterpretationTypeID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0212),
		"Interpretation Status ID", // InterpretationStatusID
		"InterpretationStatusID",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x0300),
		"Impressions", // Impressions
		"Impressions",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x4008, 0x4000),
		"Results Comments", // ResultsComments
		"ResultsComments",
		vm.VM1,
		true,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x0001),
		"Low Energy Detectors", // LowEnergyDetectors
		"LowEnergyDetectors",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x0002),
		"High Energy Detectors", // HighEnergyDetectors
		"HighEnergyDetectors",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x0004),
		"Detector Geometry Sequence", // DetectorGeometrySequence
		"DetectorGeometrySequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1001),
		"Threat ROI Voxel Sequence", // ThreatROIVoxelSequence
		"ThreatROIVoxelSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1004),
		"Threat ROI Base", // ThreatROIBase
		"ThreatROIBase",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1005),
		"Threat ROI Extents", // ThreatROIExtents
		"ThreatROIExtents",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1006),
		"Threat ROI Bitmap", // ThreatROIBitmap
		"ThreatROIBitmap",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1007),
		"Route Segment ID", // RouteSegmentID
		"RouteSegmentID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1008),
		"Gantry Type", // GantryType
		"GantryType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1009),
		"OOI Owner Type", // OOIOwnerType
		"OOIOwnerType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x100A),
		"Route Segment Sequence", // RouteSegmentSequence
		"RouteSegmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1010),
		"Potential Threat Object ID", // PotentialThreatObjectID
		"PotentialThreatObjectID",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1011),
		"Threat Sequence", // ThreatSequence
		"ThreatSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1012),
		"Threat Category", // ThreatCategory
		"ThreatCategory",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1013),
		"Threat Category Description", // ThreatCategoryDescription
		"ThreatCategoryDescription",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1014),
		"ATD Ability Assessment", // ATDAbilityAssessment
		"ATDAbilityAssessment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1015),
		"ATD Assessment Flag", // ATDAssessmentFlag
		"ATDAssessmentFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1016),
		"ATD Assessment Probability", // ATDAssessmentProbability
		"ATDAssessmentProbability",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1017),
		"Mass", // Mass
		"Mass",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1018),
		"Density", // Density
		"Density",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1019),
		"Z Effective", // ZEffective
		"ZEffective",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101A),
		"Boarding Pass ID", // BoardingPassID
		"BoardingPassID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101B),
		"Center of Mass", // CenterOfMass
		"CenterOfMass",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101C),
		"Center of PTO", // CenterOfPTO
		"CenterOfPTO",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101D),
		"Bounding Polygon", // BoundingPolygon
		"BoundingPolygon",
		vm.MustParse("6-n"),
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101E),
		"Route Segment Start Location ID", // RouteSegmentStartLocationID
		"RouteSegmentStartLocationID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x101F),
		"Route Segment End Location ID", // RouteSegmentEndLocationID
		"RouteSegmentEndLocationID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1020),
		"Route Segment Location ID Type", // RouteSegmentLocationIDType
		"RouteSegmentLocationIDType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1021),
		"Abort Reason", // AbortReason
		"AbortReason",
		vm.VM1_n,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1023),
		"Volume of PTO", // VolumeOfPTO
		"VolumeOfPTO",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1024),
		"Abort Flag", // AbortFlag
		"AbortFlag",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1025),
		"Route Segment Start Time", // RouteSegmentStartTime
		"RouteSegmentStartTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1026),
		"Route Segment End Time", // RouteSegmentEndTime
		"RouteSegmentEndTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1027),
		"TDR Type", // TDRType
		"TDRType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1028),
		"International Route Segment", // InternationalRouteSegment
		"InternationalRouteSegment",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1029),
		"Threat Detection Algorithm and Version", // ThreatDetectionAlgorithmAndVersion
		"ThreatDetectionAlgorithmAndVersion",
		vm.VM1_n,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x102A),
		"Assigned Location", // AssignedLocation
		"AssignedLocation",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x102B),
		"Alarm Decision Time", // AlarmDecisionTime
		"AlarmDecisionTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1031),
		"Alarm Decision", // AlarmDecision
		"AlarmDecision",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1033),
		"Number of Total Objects", // NumberOfTotalObjects
		"NumberOfTotalObjects",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1034),
		"Number of Alarm Objects", // NumberOfAlarmObjects
		"NumberOfAlarmObjects",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1037),
		"PTO Representation Sequence", // PTORepresentationSequence
		"PTORepresentationSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1038),
		"ATD Assessment Sequence", // ATDAssessmentSequence
		"ATDAssessmentSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1039),
		"TIP Type", // TIPType
		"TIPType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x103A),
		"DICOS Version", // DICOSVersion
		"DICOSVersion",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1041),
		"OOI Owner Creation Time", // OOIOwnerCreationTime
		"OOIOwnerCreationTime",
		vm.VM1,
		false,
		vr.DT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1042),
		"OOI Type", // OOIType
		"OOIType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1043),
		"OOI Size", // OOISize
		"OOISize",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1044),
		"Acquisition Status", // AcquisitionStatus
		"AcquisitionStatus",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1045),
		"Basis Materials Code Sequence", // BasisMaterialsCodeSequence
		"BasisMaterialsCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1046),
		"Phantom Type", // PhantomType
		"PhantomType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1047),
		"OOI Owner Sequence", // OOIOwnerSequence
		"OOIOwnerSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1048),
		"Scan Type", // ScanType
		"ScanType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1051),
		"Itinerary ID", // ItineraryID
		"ItineraryID",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1052),
		"Itinerary ID Type", // ItineraryIDType
		"ItineraryIDType",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1053),
		"Itinerary ID Assigning Authority", // ItineraryIDAssigningAuthority
		"ItineraryIDAssigningAuthority",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1054),
		"Route ID", // RouteID
		"RouteID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1055),
		"Route ID Assigning Authority", // RouteIDAssigningAuthority
		"RouteIDAssigningAuthority",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1056),
		"Inbound Arrival Type", // InboundArrivalType
		"InboundArrivalType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1058),
		"Carrier ID", // CarrierID
		"CarrierID",
		vm.VM1,
		false,
		vr.SH,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1059),
		"Carrier ID Assigning Authority", // CarrierIDAssigningAuthority
		"CarrierIDAssigningAuthority",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1060),
		"Source Orientation", // SourceOrientation
		"SourceOrientation",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1061),
		"Source Position", // SourcePosition
		"SourcePosition",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1062),
		"Belt Height", // BeltHeight
		"BeltHeight",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1064),
		"Algorithm Routing Code Sequence", // AlgorithmRoutingCodeSequence
		"AlgorithmRoutingCodeSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1067),
		"Transport Classification", // TransportClassification
		"TransportClassification",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1068),
		"OOI Type Descriptor", // OOITypeDescriptor
		"OOITypeDescriptor",
		vm.VM1,
		false,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1069),
		"Total Processing Time", // TotalProcessingTime
		"TotalProcessingTime",
		vm.VM1,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x106C),
		"Detector Calibration Data", // DetectorCalibrationData
		"DetectorCalibrationData",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x106D),
		"Additional Screening Performed", // AdditionalScreeningPerformed
		"AdditionalScreeningPerformed",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x106E),
		"Additional Inspection Selection Criteria", // AdditionalInspectionSelectionCriteria
		"AdditionalInspectionSelectionCriteria",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x106F),
		"Additional Inspection Method Sequence", // AdditionalInspectionMethodSequence
		"AdditionalInspectionMethodSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1070),
		"AIT Device Type", // AITDeviceType
		"AITDeviceType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1071),
		"QR Measurements Sequence", // QRMeasurementsSequence
		"QRMeasurementsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1072),
		"Target Material Sequence", // TargetMaterialSequence
		"TargetMaterialSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1073),
		"SNR Threshold", // SNRThreshold
		"SNRThreshold",
		vm.VM1,
		false,
		vr.FD,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1075),
		"Image Scale Representation", // ImageScaleRepresentation
		"ImageScaleRepresentation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1076),
		"Referenced PTO Sequence", // ReferencedPTOSequence
		"ReferencedPTOSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1077),
		"Referenced TDR Instance Sequence", // ReferencedTDRInstanceSequence
		"ReferencedTDRInstanceSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1078),
		"PTO Location Description", // PTOLocationDescription
		"PTOLocationDescription",
		vm.VM1,
		false,
		vr.ST,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x1079),
		"Anomaly Locator Indicator Sequence", // AnomalyLocatorIndicatorSequence
		"AnomalyLocatorIndicatorSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107A),
		"Anomaly Locator Indicator", // AnomalyLocatorIndicator
		"AnomalyLocatorIndicator",
		vm.VM3,
		false,
		vr.FL,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107B),
		"PTO Region Sequence", // PTORegionSequence
		"PTORegionSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107C),
		"Inspection Selection Criteria", // InspectionSelectionCriteria
		"InspectionSelectionCriteria",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107D),
		"Secondary Inspection Method Sequence", // SecondaryInspectionMethodSequence
		"SecondaryInspectionMethodSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x4010, 0x107E),
		"PRCS to RCS Orientation", // PRCSToRCSOrientation
		"PRCSToRCSOrientation",
		vm.VM6,
		false,
		vr.DS,
	))
	d.Add(NewEntry(
		tag.New(0x4FFE, 0x0001),
		"MAC Parameters Sequence", // MACParametersSequence
		"MACParametersSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0005)"),
		"Curve Dimensions", // CurveDimensions
		"CurveDimensions",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0010)"),
		"Number of Points", // NumberOfPoints
		"NumberOfPoints",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0020)"),
		"Type of Data", // TypeOfData
		"TypeOfData",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0022)"),
		"Curve Description", // CurveDescription
		"CurveDescription",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0030)"),
		"Axis Units", // AxisUnits
		"AxisUnits",
		vm.VM1_n,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0040)"),
		"Axis Labels", // AxisLabels
		"AxisLabels",
		vm.VM1_n,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0103)"),
		"Data Value Representation", // DataValueRepresentation
		"DataValueRepresentation",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0104)"),
		"Minimum Coordinate Value", // MinimumCoordinateValue
		"MinimumCoordinateValue",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0105)"),
		"Maximum Coordinate Value", // MaximumCoordinateValue
		"MaximumCoordinateValue",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0106)"),
		"Curve Range", // CurveRange
		"CurveRange",
		vm.VM1_n,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0110)"),
		"Curve Data Descriptor", // CurveDataDescriptor
		"CurveDataDescriptor",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0112)"),
		"Coordinate Start Value", // CoordinateStartValue
		"CoordinateStartValue",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,0114)"),
		"Coordinate Step Value", // CoordinateStepValue
		"CoordinateStepValue",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,1001)"),
		"Curve Activation Layer", // CurveActivationLayer
		"CurveActivationLayer",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2000)"),
		"Audio Type", // AudioType
		"AudioType",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2002)"),
		"Audio Sample Format", // AudioSampleFormat
		"AudioSampleFormat",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2004)"),
		"Number of Channels", // NumberOfChannels
		"NumberOfChannels",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2006)"),
		"Number of Samples", // NumberOfSamples
		"NumberOfSamples",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2008)"),
		"Sample Rate", // SampleRate
		"SampleRate",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,200A)"),
		"Total Time", // TotalTime
		"TotalTime",
		vm.VM1,
		true,
		vr.UL,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,200C)"),
		"Audio Sample Data", // AudioSampleData
		"AudioSampleData",
		vm.VM1,
		true,
		vr.OB, vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,200E)"),
		"Audio Comments", // AudioComments
		"AudioComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2500)"),
		"Curve Label", // CurveLabel
		"CurveLabel",
		vm.VM1,
		true,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2600)"),
		"Curve Referenced Overlay Sequence", // CurveReferencedOverlaySequence
		"CurveReferencedOverlaySequence",
		vm.VM1,
		true,
		vr.SQ,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,2610)"),
		"Curve Referenced Overlay Group", // CurveReferencedOverlayGroup
		"CurveReferencedOverlayGroup",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(50xx,3000)"),
		"Curve Data", // CurveData
		"CurveData",
		vm.VM1,
		true,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5200, 0x9229),
		"Shared Functional Groups Sequence", // SharedFunctionalGroupsSequence
		"SharedFunctionalGroupsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x5200, 0x9230),
		"Per-Frame Functional Groups Sequence", // PerFrameFunctionalGroupsSequence
		"PerFrameFunctionalGroupsSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x0100),
		"Waveform Sequence", // WaveformSequence
		"WaveformSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x0110),
		"Channel Minimum Value", // ChannelMinimumValue
		"ChannelMinimumValue",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x0112),
		"Channel Maximum Value", // ChannelMaximumValue
		"ChannelMaximumValue",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x1004),
		"Waveform Bits Allocated", // WaveformBitsAllocated
		"WaveformBitsAllocated",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x1006),
		"Waveform Sample Interpretation", // WaveformSampleInterpretation
		"WaveformSampleInterpretation",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x100A),
		"Waveform Padding Value", // WaveformPaddingValue
		"WaveformPaddingValue",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5400, 0x1010),
		"Waveform Data", // WaveformData
		"WaveformData",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x5600, 0x0010),
		"First Order Phase Correction Angle", // FirstOrderPhaseCorrectionAngle
		"FirstOrderPhaseCorrectionAngle",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x5600, 0x0020),
		"Spectroscopy Data", // SpectroscopyData
		"SpectroscopyData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0010)"),
		"Overlay Rows", // OverlayRows
		"OverlayRows",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0011)"),
		"Overlay Columns", // OverlayColumns
		"OverlayColumns",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0012)"),
		"Overlay Planes", // OverlayPlanes
		"OverlayPlanes",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0015)"),
		"Number of Frames in Overlay", // NumberOfFramesInOverlay
		"NumberOfFramesInOverlay",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0022)"),
		"Overlay Description", // OverlayDescription
		"OverlayDescription",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0040)"),
		"Overlay Type", // OverlayType
		"OverlayType",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0045)"),
		"Overlay Subtype", // OverlaySubtype
		"OverlaySubtype",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0050)"),
		"Overlay Origin", // OverlayOrigin
		"OverlayOrigin",
		vm.VM2,
		false,
		vr.SS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0051)"),
		"Image Frame Origin", // ImageFrameOrigin
		"ImageFrameOrigin",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0052)"),
		"Overlay Plane Origin", // OverlayPlaneOrigin
		"OverlayPlaneOrigin",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0060)"),
		"Overlay Compression Code", // OverlayCompressionCode
		"OverlayCompressionCode",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0061)"),
		"Overlay Compression Originator", // OverlayCompressionOriginator
		"OverlayCompressionOriginator",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0062)"),
		"Overlay Compression Label", // OverlayCompressionLabel
		"OverlayCompressionLabel",
		vm.VM1,
		true,
		vr.SH,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0063)"),
		"Overlay Compression Description", // OverlayCompressionDescription
		"OverlayCompressionDescription",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0066)"),
		"Overlay Compression Step Pointers", // OverlayCompressionStepPointers
		"OverlayCompressionStepPointers",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0068)"),
		"Overlay Repeat Interval", // OverlayRepeatInterval
		"OverlayRepeatInterval",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0069)"),
		"Overlay Bits Grouped", // OverlayBitsGrouped
		"OverlayBitsGrouped",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0100)"),
		"Overlay Bits Allocated", // OverlayBitsAllocated
		"OverlayBitsAllocated",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0102)"),
		"Overlay Bit Position", // OverlayBitPosition
		"OverlayBitPosition",
		vm.VM1,
		false,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0110)"),
		"Overlay Format", // OverlayFormat
		"OverlayFormat",
		vm.VM1,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0200)"),
		"Overlay Location", // OverlayLocation
		"OverlayLocation",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0800)"),
		"Overlay Code Label", // OverlayCodeLabel
		"OverlayCodeLabel",
		vm.VM1_n,
		true,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0802)"),
		"Overlay Number of Tables", // OverlayNumberOfTables
		"OverlayNumberOfTables",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0803)"),
		"Overlay Code Table Location", // OverlayCodeTableLocation
		"OverlayCodeTableLocation",
		vm.VM1_n,
		true,
		vr.AT,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,0804)"),
		"Overlay Bits For Code Word", // OverlayBitsForCodeWord
		"OverlayBitsForCodeWord",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1001)"),
		"Overlay Activation Layer", // OverlayActivationLayer
		"OverlayActivationLayer",
		vm.VM1,
		false,
		vr.CS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1100)"),
		"Overlay Descriptor - Gray", // OverlayDescriptorGray
		"OverlayDescriptorGray",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1101)"),
		"Overlay Descriptor - Red", // OverlayDescriptorRed
		"OverlayDescriptorRed",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1102)"),
		"Overlay Descriptor - Green", // OverlayDescriptorGreen
		"OverlayDescriptorGreen",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1103)"),
		"Overlay Descriptor - Blue", // OverlayDescriptorBlue
		"OverlayDescriptorBlue",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1200)"),
		"Overlays - Gray", // OverlaysGray
		"OverlaysGray",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1201)"),
		"Overlays - Red", // OverlaysRed
		"OverlaysRed",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1202)"),
		"Overlays - Green", // OverlaysGreen
		"OverlaysGreen",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1203)"),
		"Overlays - Blue", // OverlaysBlue
		"OverlaysBlue",
		vm.VM1_n,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1301)"),
		"ROI Area", // ROIArea
		"ROIArea",
		vm.VM1,
		false,
		vr.IS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1302)"),
		"ROI Mean", // ROIMean
		"ROIMean",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1303)"),
		"ROI Standard Deviation", // ROIStandardDeviation
		"ROIStandardDeviation",
		vm.VM1,
		false,
		vr.DS,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,1500)"),
		"Overlay Label", // OverlayLabel
		"OverlayLabel",
		vm.VM1,
		false,
		vr.LO,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,3000)"),
		"Overlay Data", // OverlayData
		"OverlayData",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(60xx,4000)"),
		"Overlay Comments", // OverlayComments
		"OverlayComments",
		vm.VM1,
		true,
		vr.LT,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0001),
		"Extended Offset Table", // ExtendedOffsetTable
		"ExtendedOffsetTable",
		vm.VM1,
		false,
		vr.OV,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0002),
		"Extended Offset Table Lengths", // ExtendedOffsetTableLengths
		"ExtendedOffsetTableLengths",
		vm.VM1,
		false,
		vr.OV,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0003),
		"Encapsulated Pixel Data Value Total Length", // EncapsulatedPixelDataValueTotalLength
		"EncapsulatedPixelDataValueTotalLength",
		vm.VM1,
		false,
		vr.UV,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0008),
		"Float Pixel Data", // FloatPixelData
		"FloatPixelData",
		vm.VM1,
		false,
		vr.OF,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0009),
		"Double Float Pixel Data", // DoubleFloatPixelData
		"DoubleFloatPixelData",
		vm.VM1,
		false,
		vr.OD,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0010),
		"Pixel Data", // PixelData
		"PixelData",
		vm.VM1,
		false,
		vr.OB, vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0020),
		"Coefficients SDVN", // CoefficientsSDVN
		"CoefficientsSDVN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0030),
		"Coefficients SDHN", // CoefficientsSDHN
		"CoefficientsSDHN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0x7FE0, 0x0040),
		"Coefficients SDDN", // CoefficientsSDDN
		"CoefficientsSDDN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0010)"),
		"Variable Pixel Data", // VariablePixelData
		"VariablePixelData",
		vm.VM1,
		true,
		vr.OB, vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0011)"),
		"Variable Next Data Group", // VariableNextDataGroup
		"VariableNextDataGroup",
		vm.VM1,
		true,
		vr.US,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0020)"),
		"Variable Coefficients SDVN", // VariableCoefficientsSDVN
		"VariableCoefficientsSDVN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0030)"),
		"Variable Coefficients SDHN", // VariableCoefficientsSDHN
		"VariableCoefficientsSDHN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntryWithMask(
		tag.MustParseMaskedTag("(7Fxx,0040)"),
		"Variable Coefficients SDDN", // VariableCoefficientsSDDN
		"VariableCoefficientsSDDN",
		vm.VM1,
		true,
		vr.OW,
	))
	d.Add(NewEntry(
		tag.New(0xFFFA, 0xFFFA),
		"Digital Signatures Sequence", // DigitalSignaturesSequence
		"DigitalSignaturesSequence",
		vm.VM1,
		false,
		vr.SQ,
	))
	d.Add(NewEntry(
		tag.New(0xFFFC, 0xFFFC),
		"Data Set Trailing Padding", // DataSetTrailingPadding
		"DataSetTrailingPadding",
		vm.VM1,
		false,
		vr.OB,
	))
	d.Add(NewEntry(
		tag.New(0xFFFE, 0xE000),
		"Item", // Item
		"Item",
		vm.VM1,
		false,
		vr.None,
	))
	d.Add(NewEntry(
		tag.New(0xFFFE, 0xE00D),
		"Item Delimitation Item", // ItemDelimitationItem
		"ItemDelimitationItem",
		vm.VM1,
		false,
		vr.None,
	))
	d.Add(NewEntry(
		tag.New(0xFFFE, 0xE0DD),
		"Sequence Delimitation Item", // SequenceDelimitationItem
		"SequenceDelimitationItem",
		vm.VM1,
		false,
		vr.None,
	))
}
