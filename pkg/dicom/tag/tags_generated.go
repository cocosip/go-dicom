// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DicomTagGenerated.cs. DO NOT EDIT.

package tag

// Standard DICOM tag constants
var (
	// CommandGroupLength (0000,0000) VR=UL VM=1 Command Group Length
	CommandGroupLength = New(0x0000, 0x0000)

	// AffectedSOPClassUID (0000,0002) VR=UI VM=1 Affected SOP Class UID
	AffectedSOPClassUID = New(0x0000, 0x0002)

	// RequestedSOPClassUID (0000,0003) VR=UI VM=1 Requested SOP Class UID
	RequestedSOPClassUID = New(0x0000, 0x0003)

	// CommandField (0000,0100) VR=US VM=1 Command Field
	CommandField = New(0x0000, 0x0100)

	// MessageID (0000,0110) VR=US VM=1 Message ID
	MessageID = New(0x0000, 0x0110)

	// MessageIDBeingRespondedTo (0000,0120) VR=US VM=1 Message ID Being Responded To
	MessageIDBeingRespondedTo = New(0x0000, 0x0120)

	// MoveDestination (0000,0600) VR=AE VM=1 Move Destination
	MoveDestination = New(0x0000, 0x0600)

	// Priority (0000,0700) VR=US VM=1 Priority
	Priority = New(0x0000, 0x0700)

	// CommandDataSetType (0000,0800) VR=US VM=1 Command Data Set Type
	CommandDataSetType = New(0x0000, 0x0800)

	// Status (0000,0900) VR=US VM=1 Status
	Status = New(0x0000, 0x0900)

	// OffendingElement (0000,0901) VR=AT VM=1-n Offending Element
	OffendingElement = New(0x0000, 0x0901)

	// ErrorComment (0000,0902) VR=LO VM=1 Error Comment
	ErrorComment = New(0x0000, 0x0902)

	// ErrorID (0000,0903) VR=US VM=1 Error ID
	ErrorID = New(0x0000, 0x0903)

	// AffectedSOPInstanceUID (0000,1000) VR=UI VM=1 Affected SOP Instance UID
	AffectedSOPInstanceUID = New(0x0000, 0x1000)

	// RequestedSOPInstanceUID (0000,1001) VR=UI VM=1 Requested SOP Instance UID
	RequestedSOPInstanceUID = New(0x0000, 0x1001)

	// EventTypeID (0000,1002) VR=US VM=1 Event Type ID
	EventTypeID = New(0x0000, 0x1002)

	// AttributeIdentifierList (0000,1005) VR=AT VM=1-n Attribute Identifier List
	AttributeIdentifierList = New(0x0000, 0x1005)

	// ActionTypeID (0000,1008) VR=US VM=1 Action Type ID
	ActionTypeID = New(0x0000, 0x1008)

	// NumberOfRemainingSuboperations (0000,1020) VR=US VM=1 Number of Remaining Sub-operations
	NumberOfRemainingSuboperations = New(0x0000, 0x1020)

	// NumberOfCompletedSuboperations (0000,1021) VR=US VM=1 Number of Completed Sub-operations
	NumberOfCompletedSuboperations = New(0x0000, 0x1021)

	// NumberOfFailedSuboperations (0000,1022) VR=US VM=1 Number of Failed Sub-operations
	NumberOfFailedSuboperations = New(0x0000, 0x1022)

	// NumberOfWarningSuboperations (0000,1023) VR=US VM=1 Number of Warning Sub-operations
	NumberOfWarningSuboperations = New(0x0000, 0x1023)

	// MoveOriginatorApplicationEntityTitle (0000,1030) VR=AE VM=1 Move Originator Application Entity Title
	MoveOriginatorApplicationEntityTitle = New(0x0000, 0x1030)

	// MoveOriginatorMessageID (0000,1031) VR=US VM=1 Move Originator Message ID
	MoveOriginatorMessageID = New(0x0000, 0x1031)

	// CommandLengthToEnd (0000,0001) VR=UL VM=1 Command Length to End
	CommandLengthToEnd = New(0x0000, 0x0001)

	// CommandRecognitionCode (0000,0010) VR=SH VM=1 Command Recognition Code
	CommandRecognitionCode = New(0x0000, 0x0010)

	// Initiator (0000,0200) VR=AE VM=1 Initiator
	Initiator = New(0x0000, 0x0200)

	// Receiver (0000,0300) VR=AE VM=1 Receiver
	Receiver = New(0x0000, 0x0300)

	// FindLocation (0000,0400) VR=AE VM=1 Find Location
	FindLocation = New(0x0000, 0x0400)

	// NumberOfMatches (0000,0850) VR=US VM=1 Number of Matches
	NumberOfMatches = New(0x0000, 0x0850)

	// ResponseSequenceNumber (0000,0860) VR=US VM=1 Response Sequence Number
	ResponseSequenceNumber = New(0x0000, 0x0860)

	// DialogReceiver (0000,4000) VR=LT VM=1 Dialog Receiver
	DialogReceiver = New(0x0000, 0x4000)

	// TerminalType (0000,4010) VR=LT VM=1 Terminal Type
	TerminalType = New(0x0000, 0x4010)

	// MessageSetID (0000,5010) VR=SH VM=1 Message Set ID
	MessageSetID = New(0x0000, 0x5010)

	// EndMessageID (0000,5020) VR=SH VM=1 End Message ID
	EndMessageID = New(0x0000, 0x5020)

	// DisplayFormat (0000,5110) VR=LT VM=1 Display Format
	DisplayFormat = New(0x0000, 0x5110)

	// PagePositionID (0000,5120) VR=LT VM=1 Page Position ID
	PagePositionID = New(0x0000, 0x5120)

	// TextFormatID (0000,5130) VR=CS VM=1 Text Format ID
	TextFormatID = New(0x0000, 0x5130)

	// NormalReverse (0000,5140) VR=CS VM=1 Normal/Reverse
	NormalReverse = New(0x0000, 0x5140)

	// AddGrayScale (0000,5150) VR=CS VM=1 Add Gray Scale
	AddGrayScale = New(0x0000, 0x5150)

	// Borders (0000,5160) VR=CS VM=1 Borders
	Borders = New(0x0000, 0x5160)

	// Copies (0000,5170) VR=IS VM=1 Copies
	Copies = New(0x0000, 0x5170)

	// CommandMagnificationType (0000,5180) VR=CS VM=1 Command Magnification Type
	CommandMagnificationType = New(0x0000, 0x5180)

	// Erase (0000,5190) VR=CS VM=1 Erase
	Erase = New(0x0000, 0x5190)

	// Print (0000,51A0) VR=CS VM=1 Print
	Print = New(0x0000, 0x51A0)

	// Overlays (0000,51B0) VR=US VM=1-n Overlays
	Overlays = New(0x0000, 0x51B0)

	// FileMetaInformationGroupLength (0002,0000) VR=UL VM=1 File Meta Information Group Length
	FileMetaInformationGroupLength = New(0x0002, 0x0000)

	// FileMetaInformationVersion (0002,0001) VR=OB VM=1 File Meta Information Version
	FileMetaInformationVersion = New(0x0002, 0x0001)

	// MediaStorageSOPClassUID (0002,0002) VR=UI VM=1 Media Storage SOP Class UID
	MediaStorageSOPClassUID = New(0x0002, 0x0002)

	// MediaStorageSOPInstanceUID (0002,0003) VR=UI VM=1 Media Storage SOP Instance UID
	MediaStorageSOPInstanceUID = New(0x0002, 0x0003)

	// TransferSyntaxUID (0002,0010) VR=UI VM=1 Transfer Syntax UID
	TransferSyntaxUID = New(0x0002, 0x0010)

	// ImplementationClassUID (0002,0012) VR=UI VM=1 Implementation Class UID
	ImplementationClassUID = New(0x0002, 0x0012)

	// ImplementationVersionName (0002,0013) VR=SH VM=1 Implementation Version Name
	ImplementationVersionName = New(0x0002, 0x0013)

	// SourceApplicationEntityTitle (0002,0016) VR=AE VM=1 Source Application Entity Title
	SourceApplicationEntityTitle = New(0x0002, 0x0016)

	// SendingApplicationEntityTitle (0002,0017) VR=AE VM=1 Sending Application Entity Title
	SendingApplicationEntityTitle = New(0x0002, 0x0017)

	// ReceivingApplicationEntityTitle (0002,0018) VR=AE VM=1 Receiving Application Entity Title
	ReceivingApplicationEntityTitle = New(0x0002, 0x0018)

	// SourcePresentationAddress (0002,0026) VR=UR VM=1 Source Presentation Address
	SourcePresentationAddress = New(0x0002, 0x0026)

	// SendingPresentationAddress (0002,0027) VR=UR VM=1 Sending Presentation Address
	SendingPresentationAddress = New(0x0002, 0x0027)

	// ReceivingPresentationAddress (0002,0028) VR=UR VM=1 Receiving Presentation Address
	ReceivingPresentationAddress = New(0x0002, 0x0028)

	// RTVMetaInformationVersion (0002,0031) VR=OB VM=1 RTV Meta Information Version
	RTVMetaInformationVersion = New(0x0002, 0x0031)

	// RTVCommunicationSOPClassUID (0002,0032) VR=UI VM=1 RTV Communication SOP Class UID
	RTVCommunicationSOPClassUID = New(0x0002, 0x0032)

	// RTVCommunicationSOPInstanceUID (0002,0033) VR=UI VM=1 RTV Communication SOP Instance UID
	RTVCommunicationSOPInstanceUID = New(0x0002, 0x0033)

	// RTVSourceIdentifier (0002,0035) VR=OB VM=1 RTV Source Identifier
	RTVSourceIdentifier = New(0x0002, 0x0035)

	// RTVFlowIdentifier (0002,0036) VR=OB VM=1 RTV Flow Identifier
	RTVFlowIdentifier = New(0x0002, 0x0036)

	// RTVFlowRTPSamplingRate (0002,0037) VR=UL VM=1 RTV Flow RTP Sampling Rate
	RTVFlowRTPSamplingRate = New(0x0002, 0x0037)

	// RTVFlowActualFrameDuration (0002,0038) VR=FD VM=1 RTV Flow Actual Frame Duration
	RTVFlowActualFrameDuration = New(0x0002, 0x0038)

	// PrivateInformationCreatorUID (0002,0100) VR=UI VM=1 Private Information Creator UID
	PrivateInformationCreatorUID = New(0x0002, 0x0100)

	// PrivateInformation (0002,0102) VR=OB VM=1 Private Information
	PrivateInformation = New(0x0002, 0x0102)

	// FileSetID (0004,1130) VR=CS VM=1 File-set ID
	FileSetID = New(0x0004, 0x1130)

	// FileSetDescriptorFileID (0004,1141) VR=CS VM=1-8 File-set Descriptor File ID
	FileSetDescriptorFileID = New(0x0004, 0x1141)

	// SpecificCharacterSetOfFileSetDescriptorFile (0004,1142) VR=CS VM=1 Specific Character Set of File-set Descriptor File
	SpecificCharacterSetOfFileSetDescriptorFile = New(0x0004, 0x1142)

	// OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity (0004,1200) VR=UL VM=1 Offset of the First Directory Record of the Root Directory Entity
	OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity = New(0x0004, 0x1200)

	// OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity (0004,1202) VR=UL VM=1 Offset of the Last Directory Record of the Root Directory Entity
	OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity = New(0x0004, 0x1202)

	// FileSetConsistencyFlag (0004,1212) VR=US VM=1 File-set Consistency Flag
	FileSetConsistencyFlag = New(0x0004, 0x1212)

	// DirectoryRecordSequence (0004,1220) VR=SQ VM=1 Directory Record Sequence
	DirectoryRecordSequence = New(0x0004, 0x1220)

	// OffsetOfTheNextDirectoryRecord (0004,1400) VR=UL VM=1 Offset of the Next Directory Record
	OffsetOfTheNextDirectoryRecord = New(0x0004, 0x1400)

	// RecordInUseFlag (0004,1410) VR=US VM=1 Record In-use Flag
	RecordInUseFlag = New(0x0004, 0x1410)

	// OffsetOfReferencedLowerLevelDirectoryEntity (0004,1420) VR=UL VM=1 Offset of Referenced Lower-Level Directory Entity
	OffsetOfReferencedLowerLevelDirectoryEntity = New(0x0004, 0x1420)

	// DirectoryRecordType (0004,1430) VR=CS VM=1 Directory Record Type
	DirectoryRecordType = New(0x0004, 0x1430)

	// PrivateRecordUID (0004,1432) VR=UI VM=1 Private Record UID
	PrivateRecordUID = New(0x0004, 0x1432)

	// ReferencedFileID (0004,1500) VR=CS VM=1-8 Referenced File ID
	ReferencedFileID = New(0x0004, 0x1500)

	// MRDRDirectoryRecordOffsetRETIRED (0004,1504) VR=UL VM=1 MRDR Directory Record Offset (RETIRED)
	MRDRDirectoryRecordOffsetRETIRED = New(0x0004, 0x1504)

	// ReferencedSOPClassUIDInFile (0004,1510) VR=UI VM=1 Referenced SOP Class UID in File
	ReferencedSOPClassUIDInFile = New(0x0004, 0x1510)

	// ReferencedSOPInstanceUIDInFile (0004,1511) VR=UI VM=1 Referenced SOP Instance UID in File
	ReferencedSOPInstanceUIDInFile = New(0x0004, 0x1511)

	// ReferencedTransferSyntaxUIDInFile (0004,1512) VR=UI VM=1 Referenced Transfer Syntax UID in File
	ReferencedTransferSyntaxUIDInFile = New(0x0004, 0x1512)

	// ReferencedRelatedGeneralSOPClassUIDInFile (0004,151A) VR=UI VM=1-n Referenced Related General SOP Class UID in File
	ReferencedRelatedGeneralSOPClassUIDInFile = New(0x0004, 0x151A)

	// NumberOfReferencesRETIRED (0004,1600) VR=UL VM=1 Number of References (RETIRED)
	NumberOfReferencesRETIRED = New(0x0004, 0x1600)

	// LengthToEndRETIRED (0008,0001) VR=UL VM=1 Length to End (RETIRED)
	LengthToEndRETIRED = New(0x0008, 0x0001)

	// SpecificCharacterSet (0008,0005) VR=CS VM=1-n Specific Character Set
	SpecificCharacterSet = New(0x0008, 0x0005)

	// LanguageCodeSequence (0008,0006) VR=SQ VM=1 Language Code Sequence
	LanguageCodeSequence = New(0x0008, 0x0006)

	// ImageType (0008,0008) VR=CS VM=2-n Image Type
	ImageType = New(0x0008, 0x0008)

	// RecognitionCodeRETIRED (0008,0010) VR=SH VM=1 Recognition Code (RETIRED)
	RecognitionCodeRETIRED = New(0x0008, 0x0010)

	// InstanceCreationDate (0008,0012) VR=DA VM=1 Instance Creation Date
	InstanceCreationDate = New(0x0008, 0x0012)

	// InstanceCreationTime (0008,0013) VR=TM VM=1 Instance Creation Time
	InstanceCreationTime = New(0x0008, 0x0013)

	// InstanceCreatorUID (0008,0014) VR=UI VM=1 Instance Creator UID
	InstanceCreatorUID = New(0x0008, 0x0014)

	// InstanceCoercionDateTime (0008,0015) VR=DT VM=1 Instance Coercion DateTime
	InstanceCoercionDateTime = New(0x0008, 0x0015)

	// SOPClassUID (0008,0016) VR=UI VM=1 SOP Class UID
	SOPClassUID = New(0x0008, 0x0016)

	// AcquisitionUID (0008,0017) VR=UI VM=1 Acquisition UID
	AcquisitionUID = New(0x0008, 0x0017)

	// SOPInstanceUID (0008,0018) VR=UI VM=1 SOP Instance UID
	SOPInstanceUID = New(0x0008, 0x0018)

	// PyramidUID (0008,0019) VR=UI VM=1 Pyramid UID
	PyramidUID = New(0x0008, 0x0019)

	// RelatedGeneralSOPClassUID (0008,001A) VR=UI VM=1-n Related General SOP Class UID
	RelatedGeneralSOPClassUID = New(0x0008, 0x001A)

	// OriginalSpecializedSOPClassUID (0008,001B) VR=UI VM=1 Original Specialized SOP Class UID
	OriginalSpecializedSOPClassUID = New(0x0008, 0x001B)

	// SyntheticData (0008,001C) VR=CS VM=1 Synthetic Data
	SyntheticData = New(0x0008, 0x001C)

	// StudyDate (0008,0020) VR=DA VM=1 Study Date
	StudyDate = New(0x0008, 0x0020)

	// SeriesDate (0008,0021) VR=DA VM=1 Series Date
	SeriesDate = New(0x0008, 0x0021)

	// AcquisitionDate (0008,0022) VR=DA VM=1 Acquisition Date
	AcquisitionDate = New(0x0008, 0x0022)

	// ContentDate (0008,0023) VR=DA VM=1 Content Date
	ContentDate = New(0x0008, 0x0023)

	// OverlayDateRETIRED (0008,0024) VR=DA VM=1 Overlay Date (RETIRED)
	OverlayDateRETIRED = New(0x0008, 0x0024)

	// CurveDateRETIRED (0008,0025) VR=DA VM=1 Curve Date (RETIRED)
	CurveDateRETIRED = New(0x0008, 0x0025)

	// AcquisitionDateTime (0008,002A) VR=DT VM=1 Acquisition DateTime
	AcquisitionDateTime = New(0x0008, 0x002A)

	// StudyTime (0008,0030) VR=TM VM=1 Study Time
	StudyTime = New(0x0008, 0x0030)

	// SeriesTime (0008,0031) VR=TM VM=1 Series Time
	SeriesTime = New(0x0008, 0x0031)

	// AcquisitionTime (0008,0032) VR=TM VM=1 Acquisition Time
	AcquisitionTime = New(0x0008, 0x0032)

	// ContentTime (0008,0033) VR=TM VM=1 Content Time
	ContentTime = New(0x0008, 0x0033)

	// OverlayTimeRETIRED (0008,0034) VR=TM VM=1 Overlay Time (RETIRED)
	OverlayTimeRETIRED = New(0x0008, 0x0034)

	// CurveTimeRETIRED (0008,0035) VR=TM VM=1 Curve Time (RETIRED)
	CurveTimeRETIRED = New(0x0008, 0x0035)

	// DataSetTypeRETIRED (0008,0040) VR=US VM=1 Data Set Type (RETIRED)
	DataSetTypeRETIRED = New(0x0008, 0x0040)

	// DataSetSubtypeRETIRED (0008,0041) VR=LO VM=1 Data Set Subtype (RETIRED)
	DataSetSubtypeRETIRED = New(0x0008, 0x0041)

	// NuclearMedicineSeriesTypeRETIRED (0008,0042) VR=CS VM=1 Nuclear Medicine Series Type (RETIRED)
	NuclearMedicineSeriesTypeRETIRED = New(0x0008, 0x0042)

	// AccessionNumber (0008,0050) VR=SH VM=1 Accession Number
	AccessionNumber = New(0x0008, 0x0050)

	// IssuerOfAccessionNumberSequence (0008,0051) VR=SQ VM=1 Issuer of Accession Number Sequence
	IssuerOfAccessionNumberSequence = New(0x0008, 0x0051)

	// QueryRetrieveLevel (0008,0052) VR=CS VM=1 Query/Retrieve Level
	QueryRetrieveLevel = New(0x0008, 0x0052)

	// QueryRetrieveView (0008,0053) VR=CS VM=1 Query/Retrieve View
	QueryRetrieveView = New(0x0008, 0x0053)

	// RetrieveAETitle (0008,0054) VR=AE VM=1-n Retrieve AE Title
	RetrieveAETitle = New(0x0008, 0x0054)

	// StationAETitle (0008,0055) VR=AE VM=1 Station AE Title
	StationAETitle = New(0x0008, 0x0055)

	// InstanceAvailability (0008,0056) VR=CS VM=1 Instance Availability
	InstanceAvailability = New(0x0008, 0x0056)

	// FailedSOPInstanceUIDList (0008,0058) VR=UI VM=1-n Failed SOP Instance UID List
	FailedSOPInstanceUIDList = New(0x0008, 0x0058)

	// Modality (0008,0060) VR=CS VM=1 Modality
	Modality = New(0x0008, 0x0060)

	// ModalitiesInStudy (0008,0061) VR=CS VM=1-n Modalities in Study
	ModalitiesInStudy = New(0x0008, 0x0061)

	// SOPClassesInStudy (0008,0062) VR=UI VM=1-n SOP Classes in Study
	SOPClassesInStudy = New(0x0008, 0x0062)

	// AnatomicRegionsInStudyCodeSequence (0008,0063) VR=SQ VM=1 Anatomic Regions in Study Code Sequence
	AnatomicRegionsInStudyCodeSequence = New(0x0008, 0x0063)

	// ConversionType (0008,0064) VR=CS VM=1 Conversion Type
	ConversionType = New(0x0008, 0x0064)

	// PresentationIntentType (0008,0068) VR=CS VM=1 Presentation Intent Type
	PresentationIntentType = New(0x0008, 0x0068)

	// Manufacturer (0008,0070) VR=LO VM=1 Manufacturer
	Manufacturer = New(0x0008, 0x0070)

	// InstitutionName (0008,0080) VR=LO VM=1 Institution Name
	InstitutionName = New(0x0008, 0x0080)

	// InstitutionAddress (0008,0081) VR=ST VM=1 Institution Address
	InstitutionAddress = New(0x0008, 0x0081)

	// InstitutionCodeSequence (0008,0082) VR=SQ VM=1 Institution Code Sequence
	InstitutionCodeSequence = New(0x0008, 0x0082)

	// ReferringPhysicianName (0008,0090) VR=PN VM=1 Referring Physician's Name
	ReferringPhysicianName = New(0x0008, 0x0090)

	// ReferringPhysicianAddress (0008,0092) VR=ST VM=1 Referring Physician's Address
	ReferringPhysicianAddress = New(0x0008, 0x0092)

	// ReferringPhysicianTelephoneNumbers (0008,0094) VR=SH VM=1-n Referring Physician's Telephone Numbers
	ReferringPhysicianTelephoneNumbers = New(0x0008, 0x0094)

	// ReferringPhysicianIdentificationSequence (0008,0096) VR=SQ VM=1 Referring Physician Identification Sequence
	ReferringPhysicianIdentificationSequence = New(0x0008, 0x0096)

	// ConsultingPhysicianName (0008,009C) VR=PN VM=1-n Consulting Physician's Name
	ConsultingPhysicianName = New(0x0008, 0x009C)

	// ConsultingPhysicianIdentificationSequence (0008,009D) VR=SQ VM=1 Consulting Physician Identification Sequence
	ConsultingPhysicianIdentificationSequence = New(0x0008, 0x009D)

	// CodeValue (0008,0100) VR=SH VM=1 Code Value
	CodeValue = New(0x0008, 0x0100)

	// ExtendedCodeValue (0008,0101) VR=LO VM=1 Extended Code Value
	ExtendedCodeValue = New(0x0008, 0x0101)

	// CodingSchemeDesignator (0008,0102) VR=SH VM=1 Coding Scheme Designator
	CodingSchemeDesignator = New(0x0008, 0x0102)

	// CodingSchemeVersion (0008,0103) VR=SH VM=1 Coding Scheme Version
	CodingSchemeVersion = New(0x0008, 0x0103)

	// CodeMeaning (0008,0104) VR=LO VM=1 Code Meaning
	CodeMeaning = New(0x0008, 0x0104)

	// MappingResource (0008,0105) VR=CS VM=1 Mapping Resource
	MappingResource = New(0x0008, 0x0105)

	// ContextGroupVersion (0008,0106) VR=DT VM=1 Context Group Version
	ContextGroupVersion = New(0x0008, 0x0106)

	// ContextGroupLocalVersion (0008,0107) VR=DT VM=1 Context Group Local Version
	ContextGroupLocalVersion = New(0x0008, 0x0107)

	// ExtendedCodeMeaning (0008,0108) VR=LT VM=1 Extended Code Meaning
	ExtendedCodeMeaning = New(0x0008, 0x0108)

	// CodingSchemeResourcesSequence (0008,0109) VR=SQ VM=1 Coding Scheme Resources Sequence
	CodingSchemeResourcesSequence = New(0x0008, 0x0109)

	// CodingSchemeURLType (0008,010A) VR=CS VM=1 Coding Scheme URL Type
	CodingSchemeURLType = New(0x0008, 0x010A)

	// ContextGroupExtensionFlag (0008,010B) VR=CS VM=1 Context Group Extension Flag
	ContextGroupExtensionFlag = New(0x0008, 0x010B)

	// CodingSchemeUID (0008,010C) VR=UI VM=1 Coding Scheme UID
	CodingSchemeUID = New(0x0008, 0x010C)

	// ContextGroupExtensionCreatorUID (0008,010D) VR=UI VM=1 Context Group Extension Creator UID
	ContextGroupExtensionCreatorUID = New(0x0008, 0x010D)

	// CodingSchemeURL (0008,010E) VR=UR VM=1 Coding Scheme URL
	CodingSchemeURL = New(0x0008, 0x010E)

	// ContextIdentifier (0008,010F) VR=CS VM=1 Context Identifier
	ContextIdentifier = New(0x0008, 0x010F)

	// CodingSchemeIdentificationSequence (0008,0110) VR=SQ VM=1 Coding Scheme Identification Sequence
	CodingSchemeIdentificationSequence = New(0x0008, 0x0110)

	// CodingSchemeRegistry (0008,0112) VR=LO VM=1 Coding Scheme Registry
	CodingSchemeRegistry = New(0x0008, 0x0112)

	// CodingSchemeExternalID (0008,0114) VR=ST VM=1 Coding Scheme External ID
	CodingSchemeExternalID = New(0x0008, 0x0114)

	// CodingSchemeName (0008,0115) VR=ST VM=1 Coding Scheme Name
	CodingSchemeName = New(0x0008, 0x0115)

	// CodingSchemeResponsibleOrganization (0008,0116) VR=ST VM=1 Coding Scheme Responsible Organization
	CodingSchemeResponsibleOrganization = New(0x0008, 0x0116)

	// ContextUID (0008,0117) VR=UI VM=1 Context UID
	ContextUID = New(0x0008, 0x0117)

	// MappingResourceUID (0008,0118) VR=UI VM=1 Mapping Resource UID
	MappingResourceUID = New(0x0008, 0x0118)

	// LongCodeValue (0008,0119) VR=UC VM=1 Long Code Value
	LongCodeValue = New(0x0008, 0x0119)

	// URNCodeValue (0008,0120) VR=UR VM=1 URN Code Value
	URNCodeValue = New(0x0008, 0x0120)

	// EquivalentCodeSequence (0008,0121) VR=SQ VM=1 Equivalent Code Sequence
	EquivalentCodeSequence = New(0x0008, 0x0121)

	// MappingResourceName (0008,0122) VR=LO VM=1 Mapping Resource Name
	MappingResourceName = New(0x0008, 0x0122)

	// ContextGroupIdentificationSequence (0008,0123) VR=SQ VM=1 Context Group Identification Sequence
	ContextGroupIdentificationSequence = New(0x0008, 0x0123)

	// MappingResourceIdentificationSequence (0008,0124) VR=SQ VM=1 Mapping Resource Identification Sequence
	MappingResourceIdentificationSequence = New(0x0008, 0x0124)

	// TimezoneOffsetFromUTC (0008,0201) VR=SH VM=1 Timezone Offset From UTC
	TimezoneOffsetFromUTC = New(0x0008, 0x0201)

	// ResponsibleGroupCodeSequence (0008,0220) VR=SQ VM=1 Responsible Group Code Sequence
	ResponsibleGroupCodeSequence = New(0x0008, 0x0220)

	// EquipmentModality (0008,0221) VR=CS VM=1 Equipment Modality
	EquipmentModality = New(0x0008, 0x0221)

	// ManufacturerRelatedModelGroup (0008,0222) VR=LO VM=1 Manufacturer's Related Model Group
	ManufacturerRelatedModelGroup = New(0x0008, 0x0222)

	// PrivateDataElementCharacteristicsSequence (0008,0300) VR=SQ VM=1 Private Data Element Characteristics Sequence
	PrivateDataElementCharacteristicsSequence = New(0x0008, 0x0300)

	// PrivateGroupReference (0008,0301) VR=US VM=1 Private Group Reference
	PrivateGroupReference = New(0x0008, 0x0301)

	// PrivateCreatorReference (0008,0302) VR=LO VM=1 Private Creator Reference
	PrivateCreatorReference = New(0x0008, 0x0302)

	// BlockIdentifyingInformationStatus (0008,0303) VR=CS VM=1 Block Identifying Information Status
	BlockIdentifyingInformationStatus = New(0x0008, 0x0303)

	// NonidentifyingPrivateElements (0008,0304) VR=US VM=1-n Nonidentifying Private Elements
	NonidentifyingPrivateElements = New(0x0008, 0x0304)

	// IdentifyingPrivateElements (0008,0306) VR=US VM=1-n Identifying Private Elements
	IdentifyingPrivateElements = New(0x0008, 0x0306)

	// DeidentificationActionSequence (0008,0305) VR=SQ VM=1 Deidentification Action Sequence
	DeidentificationActionSequence = New(0x0008, 0x0305)

	// DeidentificationAction (0008,0307) VR=CS VM=1 Deidentification Action
	DeidentificationAction = New(0x0008, 0x0307)

	// PrivateDataElement (0008,0308) VR=US VM=1 Private Data Element
	PrivateDataElement = New(0x0008, 0x0308)

	// PrivateDataElementValueMultiplicity (0008,0309) VR=UL VM=1-3 Private Data Element Value Multiplicity
	PrivateDataElementValueMultiplicity = New(0x0008, 0x0309)

	// PrivateDataElementValueRepresentation (0008,030A) VR=CS VM=1 Private Data Element Value Representation
	PrivateDataElementValueRepresentation = New(0x0008, 0x030A)

	// PrivateDataElementNumberOfItems (0008,030B) VR=UL VM=1-2 Private Data Element Number of Items
	PrivateDataElementNumberOfItems = New(0x0008, 0x030B)

	// PrivateDataElementName (0008,030C) VR=UC VM=1 Private Data Element Name
	PrivateDataElementName = New(0x0008, 0x030C)

	// PrivateDataElementKeyword (0008,030D) VR=UC VM=1 Private Data Element Keyword
	PrivateDataElementKeyword = New(0x0008, 0x030D)

	// PrivateDataElementDescription (0008,030E) VR=UT VM=1 Private Data Element Description
	PrivateDataElementDescription = New(0x0008, 0x030E)

	// PrivateDataElementEncoding (0008,030F) VR=UT VM=1 Private Data Element Encoding
	PrivateDataElementEncoding = New(0x0008, 0x030F)

	// PrivateDataElementDefinitionSequence (0008,0310) VR=SQ VM=1 Private Data Element Definition Sequence
	PrivateDataElementDefinitionSequence = New(0x0008, 0x0310)

	// ScopeOfInventorySequence (0008,0400) VR=SQ VM=1 Scope of Inventory Sequence
	ScopeOfInventorySequence = New(0x0008, 0x0400)

	// InventoryPurpose (0008,0401) VR=LT VM=1 Inventory Purpose
	InventoryPurpose = New(0x0008, 0x0401)

	// InventoryInstanceDescription (0008,0402) VR=LT VM=1 Inventory Instance Description
	InventoryInstanceDescription = New(0x0008, 0x0402)

	// InventoryLevel (0008,0403) VR=CS VM=1 Inventory Level
	InventoryLevel = New(0x0008, 0x0403)

	// ItemInventoryDateTime (0008,0404) VR=DT VM=1 Item Inventory DateTime
	ItemInventoryDateTime = New(0x0008, 0x0404)

	// RemovedFromOperationalUse (0008,0405) VR=CS VM=1 Removed from Operational Use
	RemovedFromOperationalUse = New(0x0008, 0x0405)

	// ReasonForRemovalCodeSequence (0008,0406) VR=SQ VM=1 Reason for Removal Code Sequence
	ReasonForRemovalCodeSequence = New(0x0008, 0x0406)

	// StoredInstanceBaseURI (0008,0407) VR=UR VM=1 Stored Instance Base URI
	StoredInstanceBaseURI = New(0x0008, 0x0407)

	// FolderAccessURI (0008,0408) VR=UR VM=1 Folder Access URI
	FolderAccessURI = New(0x0008, 0x0408)

	// FileAccessURI (0008,0409) VR=UR VM=1 File Access URI
	FileAccessURI = New(0x0008, 0x0409)

	// ContainerFileType (0008,040A) VR=CS VM=1 Container File Type
	ContainerFileType = New(0x0008, 0x040A)

	// FilenameInContainer (0008,040B) VR=UR VM=1 Filename in Container
	FilenameInContainer = New(0x0008, 0x040B)

	// FileOffsetInContainer (0008,040C) VR=UV VM=1 File Offset in Container
	FileOffsetInContainer = New(0x0008, 0x040C)

	// FileLengthInContainer (0008,040D) VR=UV VM=1 File Length in Container
	FileLengthInContainer = New(0x0008, 0x040D)

	// StoredInstanceTransferSyntaxUID (0008,040E) VR=UI VM=1 Stored Instance Transfer Syntax UID
	StoredInstanceTransferSyntaxUID = New(0x0008, 0x040E)

	// ExtendedMatchingMechanisms (0008,040F) VR=CS VM=1-n Extended Matching Mechanisms
	ExtendedMatchingMechanisms = New(0x0008, 0x040F)

	// RangeMatchingSequence (0008,0410) VR=SQ VM=1 Range Matching Sequence
	RangeMatchingSequence = New(0x0008, 0x0410)

	// ListOfUIDMatchingSequence (0008,0411) VR=SQ VM=1 List of UID Matching Sequence
	ListOfUIDMatchingSequence = New(0x0008, 0x0411)

	// EmptyValueMatchingSequence (0008,0412) VR=SQ VM=1 Empty Value Matching Sequence
	EmptyValueMatchingSequence = New(0x0008, 0x0412)

	// GeneralMatchingSequence (0008,0413) VR=SQ VM=1 General Matching Sequence
	GeneralMatchingSequence = New(0x0008, 0x0413)

	// RequestedStatusInterval (0008,0414) VR=US VM=1 Requested Status Interval
	RequestedStatusInterval = New(0x0008, 0x0414)

	// RetainInstances (0008,0415) VR=CS VM=1 Retain Instances
	RetainInstances = New(0x0008, 0x0415)

	// ExpirationDateTime (0008,0416) VR=DT VM=1 Expiration DateTime
	ExpirationDateTime = New(0x0008, 0x0416)

	// TransactionStatus (0008,0417) VR=CS VM=1 Transaction Status
	TransactionStatus = New(0x0008, 0x0417)

	// TransactionStatusComment (0008,0418) VR=LT VM=1 Transaction Status Comment
	TransactionStatusComment = New(0x0008, 0x0418)

	// FileSetAccessSequence (0008,0419) VR=SQ VM=1 File Set Access Sequence
	FileSetAccessSequence = New(0x0008, 0x0419)

	// FileAccessSequence (0008,041A) VR=SQ VM=1 File Access Sequence
	FileAccessSequence = New(0x0008, 0x041A)

	// RecordKey (0008,041B) VR=OB VM=1 Record Key
	RecordKey = New(0x0008, 0x041B)

	// PriorRecordKey (0008,041C) VR=OB VM=1 Prior Record Key
	PriorRecordKey = New(0x0008, 0x041C)

	// MetadataSequence (0008,041D) VR=SQ VM=1 Metadata Sequence
	MetadataSequence = New(0x0008, 0x041D)

	// UpdatedMetadataSequence (0008,041E) VR=SQ VM=1 Updated Metadata Sequence
	UpdatedMetadataSequence = New(0x0008, 0x041E)

	// StudyUpdateDateTime (0008,041F) VR=DT VM=1 Study Update DateTime
	StudyUpdateDateTime = New(0x0008, 0x041F)

	// InventoryAccessEndPointsSequence (0008,0420) VR=SQ VM=1 Inventory Access End Points Sequence
	InventoryAccessEndPointsSequence = New(0x0008, 0x0420)

	// StudyAccessEndPointsSequence (0008,0421) VR=SQ VM=1 Study Access End Points Sequence
	StudyAccessEndPointsSequence = New(0x0008, 0x0421)

	// IncorporatedInventoryInstanceSequence (0008,0422) VR=SQ VM=1 Incorporated Inventory Instance Sequence
	IncorporatedInventoryInstanceSequence = New(0x0008, 0x0422)

	// InventoriedStudiesSequence (0008,0423) VR=SQ VM=1 Inventoried Studies Sequence
	InventoriedStudiesSequence = New(0x0008, 0x0423)

	// InventoriedSeriesSequence (0008,0424) VR=SQ VM=1 Inventoried Series Sequence
	InventoriedSeriesSequence = New(0x0008, 0x0424)

	// InventoriedInstancesSequence (0008,0425) VR=SQ VM=1 Inventoried Instances Sequence
	InventoriedInstancesSequence = New(0x0008, 0x0425)

	// InventoryCompletionStatus (0008,0426) VR=CS VM=1 Inventory Completion Status
	InventoryCompletionStatus = New(0x0008, 0x0426)

	// NumberOfStudyRecordsInInstance (0008,0427) VR=UL VM=1 Number of Study Records in Instance
	NumberOfStudyRecordsInInstance = New(0x0008, 0x0427)

	// TotalNumberOfStudyRecords (0008,0428) VR=UV VM=1 Total Number of Study Records
	TotalNumberOfStudyRecords = New(0x0008, 0x0428)

	// MaximumNumberOfRecords (0008,0429) VR=UV VM=1 Maximum Number of Records
	MaximumNumberOfRecords = New(0x0008, 0x0429)

	// NetworkIDRETIRED (0008,1000) VR=AE VM=1 Network ID (RETIRED)
	NetworkIDRETIRED = New(0x0008, 0x1000)

	// StationName (0008,1010) VR=SH VM=1 Station Name
	StationName = New(0x0008, 0x1010)

	// StudyDescription (0008,1030) VR=LO VM=1 Study Description
	StudyDescription = New(0x0008, 0x1030)

	// ProcedureCodeSequence (0008,1032) VR=SQ VM=1 Procedure Code Sequence
	ProcedureCodeSequence = New(0x0008, 0x1032)

	// SeriesDescription (0008,103E) VR=LO VM=1 Series Description
	SeriesDescription = New(0x0008, 0x103E)

	// SeriesDescriptionCodeSequence (0008,103F) VR=SQ VM=1 Series Description Code Sequence
	SeriesDescriptionCodeSequence = New(0x0008, 0x103F)

	// InstitutionalDepartmentName (0008,1040) VR=LO VM=1 Institutional Department Name
	InstitutionalDepartmentName = New(0x0008, 0x1040)

	// InstitutionalDepartmentTypeCodeSequence (0008,1041) VR=SQ VM=1 Institutional Department Type Code Sequence
	InstitutionalDepartmentTypeCodeSequence = New(0x0008, 0x1041)

	// PhysiciansOfRecord (0008,1048) VR=PN VM=1-n Physician(s) of Record
	PhysiciansOfRecord = New(0x0008, 0x1048)

	// PhysiciansOfRecordIdentificationSequence (0008,1049) VR=SQ VM=1 Physician(s) of Record Identification Sequence
	PhysiciansOfRecordIdentificationSequence = New(0x0008, 0x1049)

	// PerformingPhysicianName (0008,1050) VR=PN VM=1-n Performing Physician's Name
	PerformingPhysicianName = New(0x0008, 0x1050)

	// PerformingPhysicianIdentificationSequence (0008,1052) VR=SQ VM=1 Performing Physician Identification Sequence
	PerformingPhysicianIdentificationSequence = New(0x0008, 0x1052)

	// NameOfPhysiciansReadingStudy (0008,1060) VR=PN VM=1-n Name of Physician(s) Reading Study
	NameOfPhysiciansReadingStudy = New(0x0008, 0x1060)

	// PhysiciansReadingStudyIdentificationSequence (0008,1062) VR=SQ VM=1 Physician(s) Reading Study Identification Sequence
	PhysiciansReadingStudyIdentificationSequence = New(0x0008, 0x1062)

	// OperatorsName (0008,1070) VR=PN VM=1-n Operators' Name
	OperatorsName = New(0x0008, 0x1070)

	// OperatorIdentificationSequence (0008,1072) VR=SQ VM=1 Operator Identification Sequence
	OperatorIdentificationSequence = New(0x0008, 0x1072)

	// AdmittingDiagnosesDescription (0008,1080) VR=LO VM=1-n Admitting Diagnoses Description
	AdmittingDiagnosesDescription = New(0x0008, 0x1080)

	// AdmittingDiagnosesCodeSequence (0008,1084) VR=SQ VM=1 Admitting Diagnoses Code Sequence
	AdmittingDiagnosesCodeSequence = New(0x0008, 0x1084)

	// PyramidDescription (0008,1088) VR=LO VM=1 Pyramid Description
	PyramidDescription = New(0x0008, 0x1088)

	// ManufacturerModelName (0008,1090) VR=LO VM=1 Manufacturer's Model Name
	ManufacturerModelName = New(0x0008, 0x1090)

	// ReferencedResultsSequenceRETIRED (0008,1100) VR=SQ VM=1 Referenced Results Sequence (RETIRED)
	ReferencedResultsSequenceRETIRED = New(0x0008, 0x1100)

	// ReferencedStudySequence (0008,1110) VR=SQ VM=1 Referenced Study Sequence
	ReferencedStudySequence = New(0x0008, 0x1110)

	// ReferencedPerformedProcedureStepSequence (0008,1111) VR=SQ VM=1 Referenced Performed Procedure Step Sequence
	ReferencedPerformedProcedureStepSequence = New(0x0008, 0x1111)

	// ReferencedInstancesBySOPClassSequence (0008,1112) VR=SQ VM=1 Referenced Instances by SOP Class Sequence
	ReferencedInstancesBySOPClassSequence = New(0x0008, 0x1112)

	// ReferencedSeriesSequence (0008,1115) VR=SQ VM=1 Referenced Series Sequence
	ReferencedSeriesSequence = New(0x0008, 0x1115)

	// ReferencedPatientSequence (0008,1120) VR=SQ VM=1 Referenced Patient Sequence
	ReferencedPatientSequence = New(0x0008, 0x1120)

	// ReferencedVisitSequence (0008,1125) VR=SQ VM=1 Referenced Visit Sequence
	ReferencedVisitSequence = New(0x0008, 0x1125)

	// ReferencedOverlaySequenceRETIRED (0008,1130) VR=SQ VM=1 Referenced Overlay Sequence (RETIRED)
	ReferencedOverlaySequenceRETIRED = New(0x0008, 0x1130)

	// ReferencedStereometricInstanceSequence (0008,1134) VR=SQ VM=1 Referenced Stereometric Instance Sequence
	ReferencedStereometricInstanceSequence = New(0x0008, 0x1134)

	// ReferencedWaveformSequence (0008,113A) VR=SQ VM=1 Referenced Waveform Sequence
	ReferencedWaveformSequence = New(0x0008, 0x113A)

	// ReferencedImageSequence (0008,1140) VR=SQ VM=1 Referenced Image Sequence
	ReferencedImageSequence = New(0x0008, 0x1140)

	// ReferencedCurveSequenceRETIRED (0008,1145) VR=SQ VM=1 Referenced Curve Sequence (RETIRED)
	ReferencedCurveSequenceRETIRED = New(0x0008, 0x1145)

	// ReferencedInstanceSequence (0008,114A) VR=SQ VM=1 Referenced Instance Sequence
	ReferencedInstanceSequence = New(0x0008, 0x114A)

	// ReferencedRealWorldValueMappingInstanceSequence (0008,114B) VR=SQ VM=1 Referenced Real World Value Mapping Instance Sequence
	ReferencedRealWorldValueMappingInstanceSequence = New(0x0008, 0x114B)

	// ReferencedSegmentationSequence (0008,114C) VR=SQ VM=1 Referenced Segmentation Sequence
	ReferencedSegmentationSequence = New(0x0008, 0x114C)

	// ReferencedSurfaceSegmentationSequence (0008,114D) VR=SQ VM=1 Referenced Surface Segmentation Sequence
	ReferencedSurfaceSegmentationSequence = New(0x0008, 0x114D)

	// ReferencedSOPClassUID (0008,1150) VR=UI VM=1 Referenced SOP Class UID
	ReferencedSOPClassUID = New(0x0008, 0x1150)

	// ReferencedSOPInstanceUID (0008,1155) VR=UI VM=1 Referenced SOP Instance UID
	ReferencedSOPInstanceUID = New(0x0008, 0x1155)

	// DefinitionSourceSequence (0008,1156) VR=SQ VM=1 Definition Source Sequence
	DefinitionSourceSequence = New(0x0008, 0x1156)

	// SOPClassesSupported (0008,115A) VR=UI VM=1-n SOP Classes Supported
	SOPClassesSupported = New(0x0008, 0x115A)

	// ReferencedFrameNumber (0008,1160) VR=IS VM=1-n Referenced Frame Number
	ReferencedFrameNumber = New(0x0008, 0x1160)

	// SimpleFrameList (0008,1161) VR=UL VM=1-n Simple Frame List
	SimpleFrameList = New(0x0008, 0x1161)

	// CalculatedFrameList (0008,1162) VR=UL VM=3-3n Calculated Frame List
	CalculatedFrameList = New(0x0008, 0x1162)

	// TimeRange (0008,1163) VR=FD VM=2 Time Range
	TimeRange = New(0x0008, 0x1163)

	// FrameExtractionSequence (0008,1164) VR=SQ VM=1 Frame Extraction Sequence
	FrameExtractionSequence = New(0x0008, 0x1164)

	// MultiFrameSourceSOPInstanceUID (0008,1167) VR=UI VM=1 Multi-frame Source SOP Instance UID
	MultiFrameSourceSOPInstanceUID = New(0x0008, 0x1167)

	// RetrieveURL (0008,1190) VR=UR VM=1 Retrieve URL
	RetrieveURL = New(0x0008, 0x1190)

	// TransactionUID (0008,1195) VR=UI VM=1 Transaction UID
	TransactionUID = New(0x0008, 0x1195)

	// WarningReason (0008,1196) VR=US VM=1 Warning Reason
	WarningReason = New(0x0008, 0x1196)

	// FailureReason (0008,1197) VR=US VM=1 Failure Reason
	FailureReason = New(0x0008, 0x1197)

	// FailedSOPSequence (0008,1198) VR=SQ VM=1 Failed SOP Sequence
	FailedSOPSequence = New(0x0008, 0x1198)

	// ReferencedSOPSequence (0008,1199) VR=SQ VM=1 Referenced SOP Sequence
	ReferencedSOPSequence = New(0x0008, 0x1199)

	// OtherFailuresSequence (0008,119A) VR=SQ VM=1 Other Failures Sequence
	OtherFailuresSequence = New(0x0008, 0x119A)

	// FailedStudySequence (0008,119B) VR=SQ VM=1 Failed Study Sequence
	FailedStudySequence = New(0x0008, 0x119B)

	// StudiesContainingOtherReferencedInstancesSequence (0008,1200) VR=SQ VM=1 Studies Containing Other Referenced Instances Sequence
	StudiesContainingOtherReferencedInstancesSequence = New(0x0008, 0x1200)

	// RelatedSeriesSequence (0008,1250) VR=SQ VM=1 Related Series Sequence
	RelatedSeriesSequence = New(0x0008, 0x1250)

	// PrincipalDiagnosisCodeSequence (0008,1301) VR=SQ VM=1 Principal Diagnosis Code Sequence
	PrincipalDiagnosisCodeSequence = New(0x0008, 0x1301)

	// PrimaryDiagnosisCodeSequence (0008,1302) VR=SQ VM=1 Primary Diagnosis Code Sequence
	PrimaryDiagnosisCodeSequence = New(0x0008, 0x1302)

	// SecondaryDiagnosesCodeSequence (0008,1303) VR=SQ VM=1 Secondary Diagnoses Code Sequence
	SecondaryDiagnosesCodeSequence = New(0x0008, 0x1303)

	// HistologicalDiagnosesCodeSequence (0008,1304) VR=SQ VM=1 Histological Diagnoses Code Sequence
	HistologicalDiagnosesCodeSequence = New(0x0008, 0x1304)

	// LossyImageCompressionRetiredRETIRED (0008,2110) VR=CS VM=1 Lossy Image Compression (Retired) (RETIRED)
	LossyImageCompressionRetiredRETIRED = New(0x0008, 0x2110)

	// DerivationDescription (0008,2111) VR=ST VM=1 Derivation Description
	DerivationDescription = New(0x0008, 0x2111)

	// SourceImageSequence (0008,2112) VR=SQ VM=1 Source Image Sequence
	SourceImageSequence = New(0x0008, 0x2112)

	// StageName (0008,2120) VR=SH VM=1 Stage Name
	StageName = New(0x0008, 0x2120)

	// StageNumber (0008,2122) VR=IS VM=1 Stage Number
	StageNumber = New(0x0008, 0x2122)

	// NumberOfStages (0008,2124) VR=IS VM=1 Number of Stages
	NumberOfStages = New(0x0008, 0x2124)

	// ViewName (0008,2127) VR=SH VM=1 View Name
	ViewName = New(0x0008, 0x2127)

	// ViewNumber (0008,2128) VR=IS VM=1 View Number
	ViewNumber = New(0x0008, 0x2128)

	// NumberOfEventTimers (0008,2129) VR=IS VM=1 Number of Event Timers
	NumberOfEventTimers = New(0x0008, 0x2129)

	// NumberOfViewsInStage (0008,212A) VR=IS VM=1 Number of Views in Stage
	NumberOfViewsInStage = New(0x0008, 0x212A)

	// EventElapsedTimes (0008,2130) VR=DS VM=1-n Event Elapsed Time(s)
	EventElapsedTimes = New(0x0008, 0x2130)

	// EventTimerNames (0008,2132) VR=LO VM=1-n Event Timer Name(s)
	EventTimerNames = New(0x0008, 0x2132)

	// EventTimerSequence (0008,2133) VR=SQ VM=1 Event Timer Sequence
	EventTimerSequence = New(0x0008, 0x2133)

	// EventTimeOffset (0008,2134) VR=FD VM=1 Event Time Offset
	EventTimeOffset = New(0x0008, 0x2134)

	// EventCodeSequence (0008,2135) VR=SQ VM=1 Event Code Sequence
	EventCodeSequence = New(0x0008, 0x2135)

	// StartTrim (0008,2142) VR=IS VM=1 Start Trim
	StartTrim = New(0x0008, 0x2142)

	// StopTrim (0008,2143) VR=IS VM=1 Stop Trim
	StopTrim = New(0x0008, 0x2143)

	// RecommendedDisplayFrameRate (0008,2144) VR=IS VM=1 Recommended Display Frame Rate
	RecommendedDisplayFrameRate = New(0x0008, 0x2144)

	// TransducerPositionRETIRED (0008,2200) VR=CS VM=1 Transducer Position (RETIRED)
	TransducerPositionRETIRED = New(0x0008, 0x2200)

	// TransducerOrientationRETIRED (0008,2204) VR=CS VM=1 Transducer Orientation (RETIRED)
	TransducerOrientationRETIRED = New(0x0008, 0x2204)

	// AnatomicStructureRETIRED (0008,2208) VR=CS VM=1 Anatomic Structure (RETIRED)
	AnatomicStructureRETIRED = New(0x0008, 0x2208)

	// AnatomicRegionSequence (0008,2218) VR=SQ VM=1 Anatomic Region Sequence
	AnatomicRegionSequence = New(0x0008, 0x2218)

	// AnatomicRegionModifierSequence (0008,2220) VR=SQ VM=1 Anatomic Region Modifier Sequence
	AnatomicRegionModifierSequence = New(0x0008, 0x2220)

	// PrimaryAnatomicStructureSequence (0008,2228) VR=SQ VM=1 Primary Anatomic Structure Sequence
	PrimaryAnatomicStructureSequence = New(0x0008, 0x2228)

	// AnatomicStructureSpaceOrRegionSequenceRETIRED (0008,2229) VR=SQ VM=1 Anatomic Structure, Space or Region Sequence (RETIRED)
	AnatomicStructureSpaceOrRegionSequenceRETIRED = New(0x0008, 0x2229)

	// PrimaryAnatomicStructureModifierSequence (0008,2230) VR=SQ VM=1 Primary Anatomic Structure Modifier Sequence
	PrimaryAnatomicStructureModifierSequence = New(0x0008, 0x2230)

	// TransducerPositionSequenceRETIRED (0008,2240) VR=SQ VM=1 Transducer Position Sequence (RETIRED)
	TransducerPositionSequenceRETIRED = New(0x0008, 0x2240)

	// TransducerPositionModifierSequenceRETIRED (0008,2242) VR=SQ VM=1 Transducer Position Modifier Sequence (RETIRED)
	TransducerPositionModifierSequenceRETIRED = New(0x0008, 0x2242)

	// TransducerOrientationSequenceRETIRED (0008,2244) VR=SQ VM=1 Transducer Orientation Sequence (RETIRED)
	TransducerOrientationSequenceRETIRED = New(0x0008, 0x2244)

	// TransducerOrientationModifierSequenceRETIRED (0008,2246) VR=SQ VM=1 Transducer Orientation Modifier Sequence (RETIRED)
	TransducerOrientationModifierSequenceRETIRED = New(0x0008, 0x2246)

	// AnatomicStructureSpaceOrRegionCodeSequenceTrialRETIRED (0008,2251) VR=SQ VM=1 Anatomic Structure Space Or Region Code Sequence (Trial) (RETIRED)
	AnatomicStructureSpaceOrRegionCodeSequenceTrialRETIRED = New(0x0008, 0x2251)

	// AnatomicPortalOfEntranceCodeSequenceTrialRETIRED (0008,2253) VR=SQ VM=1 Anatomic Portal Of Entrance Code Sequence (Trial) (RETIRED)
	AnatomicPortalOfEntranceCodeSequenceTrialRETIRED = New(0x0008, 0x2253)

	// AnatomicApproachDirectionCodeSequenceTrialRETIRED (0008,2255) VR=SQ VM=1 Anatomic Approach Direction Code Sequence (Trial) (RETIRED)
	AnatomicApproachDirectionCodeSequenceTrialRETIRED = New(0x0008, 0x2255)

	// AnatomicPerspectiveDescriptionTrialRETIRED (0008,2256) VR=ST VM=1 Anatomic Perspective Description (Trial) (RETIRED)
	AnatomicPerspectiveDescriptionTrialRETIRED = New(0x0008, 0x2256)

	// AnatomicPerspectiveCodeSequenceTrialRETIRED (0008,2257) VR=SQ VM=1 Anatomic Perspective Code Sequence (Trial) (RETIRED)
	AnatomicPerspectiveCodeSequenceTrialRETIRED = New(0x0008, 0x2257)

	// AnatomicLocationOfExaminingInstrumentDescriptionTrialRETIRED (0008,2258) VR=ST VM=1 Anatomic Location Of Examining Instrument Description (Trial) (RETIRED)
	AnatomicLocationOfExaminingInstrumentDescriptionTrialRETIRED = New(0x0008, 0x2258)

	// AnatomicLocationOfExaminingInstrumentCodeSequenceTrialRETIRED (0008,2259) VR=SQ VM=1 Anatomic Location Of Examining Instrument Code Sequence (Trial) (RETIRED)
	AnatomicLocationOfExaminingInstrumentCodeSequenceTrialRETIRED = New(0x0008, 0x2259)

	// AnatomicStructureSpaceOrRegionModifierCodeSequenceTrialRETIRED (0008,225A) VR=SQ VM=1 Anatomic Structure Space Or Region Modifier Code Sequence (Trial) (RETIRED)
	AnatomicStructureSpaceOrRegionModifierCodeSequenceTrialRETIRED = New(0x0008, 0x225A)

	// OnAxisBackgroundAnatomicStructureCodeSequenceTrialRETIRED (0008,225C) VR=SQ VM=1 On Axis Background Anatomic Structure Code Sequence (Trial) (RETIRED)
	OnAxisBackgroundAnatomicStructureCodeSequenceTrialRETIRED = New(0x0008, 0x225C)

	// AlternateRepresentationSequence (0008,3001) VR=SQ VM=1 Alternate Representation Sequence
	AlternateRepresentationSequence = New(0x0008, 0x3001)

	// AvailableTransferSyntaxUID (0008,3002) VR=UI VM=1-n Available Transfer Syntax UID
	AvailableTransferSyntaxUID = New(0x0008, 0x3002)

	// IrradiationEventUID (0008,3010) VR=UI VM=1-n Irradiation Event UID
	IrradiationEventUID = New(0x0008, 0x3010)

	// SourceIrradiationEventSequence (0008,3011) VR=SQ VM=1 Source Irradiation Event Sequence
	SourceIrradiationEventSequence = New(0x0008, 0x3011)

	// RadiopharmaceuticalAdministrationEventUID (0008,3012) VR=UI VM=1 Radiopharmaceutical Administration Event UID
	RadiopharmaceuticalAdministrationEventUID = New(0x0008, 0x3012)

	// IdentifyingCommentsRETIRED (0008,4000) VR=LT VM=1 Identifying Comments (RETIRED)
	IdentifyingCommentsRETIRED = New(0x0008, 0x4000)

	// FrameType (0008,9007) VR=CS VM=4-5 Frame Type
	FrameType = New(0x0008, 0x9007)

	// ReferencedImageEvidenceSequence (0008,9092) VR=SQ VM=1 Referenced Image Evidence Sequence
	ReferencedImageEvidenceSequence = New(0x0008, 0x9092)

	// ReferencedRawDataSequence (0008,9121) VR=SQ VM=1 Referenced Raw Data Sequence
	ReferencedRawDataSequence = New(0x0008, 0x9121)

	// CreatorVersionUID (0008,9123) VR=UI VM=1 Creator-Version UID
	CreatorVersionUID = New(0x0008, 0x9123)

	// DerivationImageSequence (0008,9124) VR=SQ VM=1 Derivation Image Sequence
	DerivationImageSequence = New(0x0008, 0x9124)

	// SourceImageEvidenceSequence (0008,9154) VR=SQ VM=1 Source Image Evidence Sequence
	SourceImageEvidenceSequence = New(0x0008, 0x9154)

	// PixelPresentation (0008,9205) VR=CS VM=1 Pixel Presentation
	PixelPresentation = New(0x0008, 0x9205)

	// VolumetricProperties (0008,9206) VR=CS VM=1 Volumetric Properties
	VolumetricProperties = New(0x0008, 0x9206)

	// VolumeBasedCalculationTechnique (0008,9207) VR=CS VM=1 Volume Based Calculation Technique
	VolumeBasedCalculationTechnique = New(0x0008, 0x9207)

	// ComplexImageComponent (0008,9208) VR=CS VM=1 Complex Image Component
	ComplexImageComponent = New(0x0008, 0x9208)

	// AcquisitionContrast (0008,9209) VR=CS VM=1 Acquisition Contrast
	AcquisitionContrast = New(0x0008, 0x9209)

	// DerivationCodeSequence (0008,9215) VR=SQ VM=1 Derivation Code Sequence
	DerivationCodeSequence = New(0x0008, 0x9215)

	// ReferencedPresentationStateSequence (0008,9237) VR=SQ VM=1 Referenced Presentation State Sequence
	ReferencedPresentationStateSequence = New(0x0008, 0x9237)

	// ReferencedOtherPlaneSequence (0008,9410) VR=SQ VM=1 Referenced Other Plane Sequence
	ReferencedOtherPlaneSequence = New(0x0008, 0x9410)

	// FrameDisplaySequence (0008,9458) VR=SQ VM=1 Frame Display Sequence
	FrameDisplaySequence = New(0x0008, 0x9458)

	// RecommendedDisplayFrameRateInFloat (0008,9459) VR=FL VM=1 Recommended Display Frame Rate in Float
	RecommendedDisplayFrameRateInFloat = New(0x0008, 0x9459)

	// SkipFrameRangeFlag (0008,9460) VR=CS VM=1 Skip Frame Range Flag
	SkipFrameRangeFlag = New(0x0008, 0x9460)

	// PatientName (0010,0010) VR=PN VM=1 Patient's Name
	PatientName = New(0x0010, 0x0010)

	// PersonNamesToUseSequence (0010,0011) VR=SQ VM=1 Person Names to Use Sequence
	PersonNamesToUseSequence = New(0x0010, 0x0011)

	// NameToUse (0010,0012) VR=LT VM=1 Name to Use
	NameToUse = New(0x0010, 0x0012)

	// NameToUseComment (0010,0013) VR=UT VM=1 Name to Use Comment
	NameToUseComment = New(0x0010, 0x0013)

	// ThirdPersonPronounsSequence (0010,0014) VR=SQ VM=1 Third Person Pronouns Sequence
	ThirdPersonPronounsSequence = New(0x0010, 0x0014)

	// PronounCodeSequence (0010,0015) VR=SQ VM=1 Pronoun Code Sequence
	PronounCodeSequence = New(0x0010, 0x0015)

	// PronounComment (0010,0016) VR=UT VM=1 Pronoun Comment
	PronounComment = New(0x0010, 0x0016)

	// PatientID (0010,0020) VR=LO VM=1 Patient ID
	PatientID = New(0x0010, 0x0020)

	// IssuerOfPatientID (0010,0021) VR=LO VM=1 Issuer of Patient ID
	IssuerOfPatientID = New(0x0010, 0x0021)

	// TypeOfPatientID (0010,0022) VR=CS VM=1 Type of Patient ID
	TypeOfPatientID = New(0x0010, 0x0022)

	// IssuerOfPatientIDQualifiersSequence (0010,0024) VR=SQ VM=1 Issuer of Patient ID Qualifiers Sequence
	IssuerOfPatientIDQualifiersSequence = New(0x0010, 0x0024)

	// SourcePatientGroupIdentificationSequence (0010,0026) VR=SQ VM=1 Source Patient Group Identification Sequence
	SourcePatientGroupIdentificationSequence = New(0x0010, 0x0026)

	// GroupOfPatientsIdentificationSequence (0010,0027) VR=SQ VM=1 Group of Patients Identification Sequence
	GroupOfPatientsIdentificationSequence = New(0x0010, 0x0027)

	// SubjectRelativePositionInImage (0010,0028) VR=US VM=3 Subject Relative Position in Image
	SubjectRelativePositionInImage = New(0x0010, 0x0028)

	// PatientBirthDate (0010,0030) VR=DA VM=1 Patient's Birth Date
	PatientBirthDate = New(0x0010, 0x0030)

	// PatientBirthTime (0010,0032) VR=TM VM=1 Patient's Birth Time
	PatientBirthTime = New(0x0010, 0x0032)

	// PatientBirthDateInAlternativeCalendar (0010,0033) VR=LO VM=1 Patient's Birth Date in Alternative Calendar
	PatientBirthDateInAlternativeCalendar = New(0x0010, 0x0033)

	// PatientDeathDateInAlternativeCalendar (0010,0034) VR=LO VM=1 Patient's Death Date in Alternative Calendar
	PatientDeathDateInAlternativeCalendar = New(0x0010, 0x0034)

	// PatientAlternativeCalendar (0010,0035) VR=CS VM=1 Patient's Alternative Calendar
	PatientAlternativeCalendar = New(0x0010, 0x0035)

	// PatientSex (0010,0040) VR=CS VM=1 Patient's Sex
	PatientSex = New(0x0010, 0x0040)

	// GenderIdentitySequence (0010,0041) VR=SQ VM=1 Gender Identity Sequence
	GenderIdentitySequence = New(0x0010, 0x0041)

	// SexParametersForClinicalUseCategoryComment (0010,0042) VR=UT VM=1 Sex Parameters for Clinical Use Category Comment
	SexParametersForClinicalUseCategoryComment = New(0x0010, 0x0042)

	// SexParametersForClinicalUseCategorySequence (0010,0043) VR=SQ VM=1 Sex Parameters for Clinical Use Category Sequence
	SexParametersForClinicalUseCategorySequence = New(0x0010, 0x0043)

	// GenderIdentityCodeSequence (0010,0044) VR=SQ VM=1 Gender Identity Code Sequence
	GenderIdentityCodeSequence = New(0x0010, 0x0044)

	// GenderIdentityComment (0010,0045) VR=UT VM=1 Gender Identity Comment
	GenderIdentityComment = New(0x0010, 0x0045)

	// SexParametersForClinicalUseCategoryCodeSequence (0010,0046) VR=SQ VM=1 Sex Parameters for Clinical Use Category Code Sequence
	SexParametersForClinicalUseCategoryCodeSequence = New(0x0010, 0x0046)

	// SexParametersForClinicalUseCategoryReference (0010,0047) VR=UR VM=1-n Sex Parameters for Clinical Use Category Reference
	SexParametersForClinicalUseCategoryReference = New(0x0010, 0x0047)

	// PatientInsurancePlanCodeSequence (0010,0050) VR=SQ VM=1 Patient's Insurance Plan Code Sequence
	PatientInsurancePlanCodeSequence = New(0x0010, 0x0050)

	// PatientPrimaryLanguageCodeSequence (0010,0101) VR=SQ VM=1 Patient's Primary Language Code Sequence
	PatientPrimaryLanguageCodeSequence = New(0x0010, 0x0101)

	// PatientPrimaryLanguageModifierCodeSequence (0010,0102) VR=SQ VM=1 Patient's Primary Language Modifier Code Sequence
	PatientPrimaryLanguageModifierCodeSequence = New(0x0010, 0x0102)

	// QualityControlSubject (0010,0200) VR=CS VM=1 Quality Control Subject
	QualityControlSubject = New(0x0010, 0x0200)

	// QualityControlSubjectTypeCodeSequence (0010,0201) VR=SQ VM=1 Quality Control Subject Type Code Sequence
	QualityControlSubjectTypeCodeSequence = New(0x0010, 0x0201)

	// StrainDescription (0010,0212) VR=UC VM=1 Strain Description
	StrainDescription = New(0x0010, 0x0212)

	// StrainNomenclature (0010,0213) VR=LO VM=1 Strain Nomenclature
	StrainNomenclature = New(0x0010, 0x0213)

	// StrainStockNumber (0010,0214) VR=LO VM=1 Strain Stock Number
	StrainStockNumber = New(0x0010, 0x0214)

	// StrainSourceRegistryCodeSequence (0010,0215) VR=SQ VM=1 Strain Source Registry Code Sequence
	StrainSourceRegistryCodeSequence = New(0x0010, 0x0215)

	// StrainStockSequence (0010,0216) VR=SQ VM=1 Strain Stock Sequence
	StrainStockSequence = New(0x0010, 0x0216)

	// StrainSource (0010,0217) VR=LO VM=1 Strain Source
	StrainSource = New(0x0010, 0x0217)

	// StrainAdditionalInformation (0010,0218) VR=UT VM=1 Strain Additional Information
	StrainAdditionalInformation = New(0x0010, 0x0218)

	// StrainCodeSequence (0010,0219) VR=SQ VM=1 Strain Code Sequence
	StrainCodeSequence = New(0x0010, 0x0219)

	// GeneticModificationsSequence (0010,0221) VR=SQ VM=1 Genetic Modifications Sequence
	GeneticModificationsSequence = New(0x0010, 0x0221)

	// GeneticModificationsDescription (0010,0222) VR=UC VM=1 Genetic Modifications Description
	GeneticModificationsDescription = New(0x0010, 0x0222)

	// GeneticModificationsNomenclature (0010,0223) VR=LO VM=1 Genetic Modifications Nomenclature
	GeneticModificationsNomenclature = New(0x0010, 0x0223)

	// GeneticModificationsCodeSequence (0010,0229) VR=SQ VM=1 Genetic Modifications Code Sequence
	GeneticModificationsCodeSequence = New(0x0010, 0x0229)

	// OtherPatientIDsRETIRED (0010,1000) VR=LO VM=1-n Other Patient IDs (RETIRED)
	OtherPatientIDsRETIRED = New(0x0010, 0x1000)

	// OtherPatientNames (0010,1001) VR=PN VM=1-n Other Patient Names
	OtherPatientNames = New(0x0010, 0x1001)

	// OtherPatientIDsSequence (0010,1002) VR=SQ VM=1 Other Patient IDs Sequence
	OtherPatientIDsSequence = New(0x0010, 0x1002)

	// PatientBirthName (0010,1005) VR=PN VM=1 Patient's Birth Name
	PatientBirthName = New(0x0010, 0x1005)

	// PatientAge (0010,1010) VR=AS VM=1 Patient's Age
	PatientAge = New(0x0010, 0x1010)

	// PatientSize (0010,1020) VR=DS VM=1 Patient's Size
	PatientSize = New(0x0010, 0x1020)

	// PatientSizeCodeSequence (0010,1021) VR=SQ VM=1 Patient's Size Code Sequence
	PatientSizeCodeSequence = New(0x0010, 0x1021)

	// PatientBodyMassIndex (0010,1022) VR=DS VM=1 Patient's Body Mass Index
	PatientBodyMassIndex = New(0x0010, 0x1022)

	// MeasuredAPDimension (0010,1023) VR=DS VM=1 Measured AP Dimension
	MeasuredAPDimension = New(0x0010, 0x1023)

	// MeasuredLateralDimension (0010,1024) VR=DS VM=1 Measured Lateral Dimension
	MeasuredLateralDimension = New(0x0010, 0x1024)

	// PatientWeight (0010,1030) VR=DS VM=1 Patient's Weight
	PatientWeight = New(0x0010, 0x1030)

	// PatientAddress (0010,1040) VR=LO VM=1 Patient's Address
	PatientAddress = New(0x0010, 0x1040)

	// InsurancePlanIdentificationRETIRED (0010,1050) VR=LO VM=1-n Insurance Plan Identification (RETIRED)
	InsurancePlanIdentificationRETIRED = New(0x0010, 0x1050)

	// PatientMotherBirthName (0010,1060) VR=PN VM=1 Patient's Mother's Birth Name
	PatientMotherBirthName = New(0x0010, 0x1060)

	// MilitaryRank (0010,1080) VR=LO VM=1 Military Rank
	MilitaryRank = New(0x0010, 0x1080)

	// BranchOfService (0010,1081) VR=LO VM=1 Branch of Service
	BranchOfService = New(0x0010, 0x1081)

	// MedicalRecordLocatorRETIRED (0010,1090) VR=LO VM=1 Medical Record Locator (RETIRED)
	MedicalRecordLocatorRETIRED = New(0x0010, 0x1090)

	// ReferencedPatientPhotoSequence (0010,1100) VR=SQ VM=1 Referenced Patient Photo Sequence
	ReferencedPatientPhotoSequence = New(0x0010, 0x1100)

	// MedicalAlerts (0010,2000) VR=LO VM=1-n Medical Alerts
	MedicalAlerts = New(0x0010, 0x2000)

	// Allergies (0010,2110) VR=LO VM=1-n Allergies
	Allergies = New(0x0010, 0x2110)

	// CountryOfResidence (0010,2150) VR=LO VM=1 Country of Residence
	CountryOfResidence = New(0x0010, 0x2150)

	// RegionOfResidence (0010,2152) VR=LO VM=1 Region of Residence
	RegionOfResidence = New(0x0010, 0x2152)

	// PatientTelephoneNumbers (0010,2154) VR=SH VM=1-n Patient's Telephone Numbers
	PatientTelephoneNumbers = New(0x0010, 0x2154)

	// PatientTelecomInformation (0010,2155) VR=LT VM=1 Patient's Telecom Information
	PatientTelecomInformation = New(0x0010, 0x2155)

	// EthnicGroupRETIRED (0010,2160) VR=SH VM=1 Ethnic Group (RETIRED)
	EthnicGroupRETIRED = New(0x0010, 0x2160)

	// EthnicGroupCodeSequence (0010,2161) VR=SQ VM=1 Ethnic Group Code Sequence
	EthnicGroupCodeSequence = New(0x0010, 0x2161)

	// EthnicGroups (0010,2162) VR=UC VM=1-n Ethnic Groups
	EthnicGroups = New(0x0010, 0x2162)

	// Occupation (0010,2180) VR=SH VM=1 Occupation
	Occupation = New(0x0010, 0x2180)

	// SmokingStatus (0010,21A0) VR=CS VM=1 Smoking Status
	SmokingStatus = New(0x0010, 0x21A0)

	// AdditionalPatientHistory (0010,21B0) VR=LT VM=1 Additional Patient History
	AdditionalPatientHistory = New(0x0010, 0x21B0)

	// PregnancyStatus (0010,21C0) VR=US VM=1 Pregnancy Status
	PregnancyStatus = New(0x0010, 0x21C0)

	// LastMenstrualDate (0010,21D0) VR=DA VM=1 Last Menstrual Date
	LastMenstrualDate = New(0x0010, 0x21D0)

	// PatientReligiousPreference (0010,21F0) VR=LO VM=1 Patient's Religious Preference
	PatientReligiousPreference = New(0x0010, 0x21F0)

	// PatientSpeciesDescription (0010,2201) VR=LO VM=1 Patient Species Description
	PatientSpeciesDescription = New(0x0010, 0x2201)

	// PatientSpeciesCodeSequence (0010,2202) VR=SQ VM=1 Patient Species Code Sequence
	PatientSpeciesCodeSequence = New(0x0010, 0x2202)

	// PatientSexNeutered (0010,2203) VR=CS VM=1 Patient's Sex Neutered
	PatientSexNeutered = New(0x0010, 0x2203)

	// AnatomicalOrientationType (0010,2210) VR=CS VM=1 Anatomical Orientation Type
	AnatomicalOrientationType = New(0x0010, 0x2210)

	// PatientBreedDescription (0010,2292) VR=LO VM=1 Patient Breed Description
	PatientBreedDescription = New(0x0010, 0x2292)

	// PatientBreedCodeSequence (0010,2293) VR=SQ VM=1 Patient Breed Code Sequence
	PatientBreedCodeSequence = New(0x0010, 0x2293)

	// BreedRegistrationSequence (0010,2294) VR=SQ VM=1 Breed Registration Sequence
	BreedRegistrationSequence = New(0x0010, 0x2294)

	// BreedRegistrationNumber (0010,2295) VR=LO VM=1 Breed Registration Number
	BreedRegistrationNumber = New(0x0010, 0x2295)

	// BreedRegistryCodeSequence (0010,2296) VR=SQ VM=1 Breed Registry Code Sequence
	BreedRegistryCodeSequence = New(0x0010, 0x2296)

	// ResponsiblePerson (0010,2297) VR=PN VM=1 Responsible Person
	ResponsiblePerson = New(0x0010, 0x2297)

	// ResponsiblePersonRole (0010,2298) VR=CS VM=1 Responsible Person Role
	ResponsiblePersonRole = New(0x0010, 0x2298)

	// ResponsibleOrganization (0010,2299) VR=LO VM=1 Responsible Organization
	ResponsibleOrganization = New(0x0010, 0x2299)

	// PatientComments (0010,4000) VR=LT VM=1 Patient Comments
	PatientComments = New(0x0010, 0x4000)

	// ExaminedBodyThickness (0010,9431) VR=FL VM=1 Examined Body Thickness
	ExaminedBodyThickness = New(0x0010, 0x9431)

	// ClinicalTrialSponsorName (0012,0010) VR=LO VM=1 Clinical Trial Sponsor Name
	ClinicalTrialSponsorName = New(0x0012, 0x0010)

	// ClinicalTrialProtocolID (0012,0020) VR=LO VM=1 Clinical Trial Protocol ID
	ClinicalTrialProtocolID = New(0x0012, 0x0020)

	// ClinicalTrialProtocolName (0012,0021) VR=LO VM=1 Clinical Trial Protocol Name
	ClinicalTrialProtocolName = New(0x0012, 0x0021)

	// IssuerOfClinicalTrialProtocolID (0012,0022) VR=LO VM=1 Issuer of Clinical Trial Protocol ID
	IssuerOfClinicalTrialProtocolID = New(0x0012, 0x0022)

	// OtherClinicalTrialProtocolIDsSequence (0012,0023) VR=SQ VM=1 Other Clinical Trial Protocol IDs Sequence
	OtherClinicalTrialProtocolIDsSequence = New(0x0012, 0x0023)

	// ClinicalTrialSiteID (0012,0030) VR=LO VM=1 Clinical Trial Site ID
	ClinicalTrialSiteID = New(0x0012, 0x0030)

	// ClinicalTrialSiteName (0012,0031) VR=LO VM=1 Clinical Trial Site Name
	ClinicalTrialSiteName = New(0x0012, 0x0031)

	// IssuerOfClinicalTrialSiteID (0012,0032) VR=LO VM=1 Issuer of Clinical Trial Site ID
	IssuerOfClinicalTrialSiteID = New(0x0012, 0x0032)

	// ClinicalTrialSubjectID (0012,0040) VR=LO VM=1 Clinical Trial Subject ID
	ClinicalTrialSubjectID = New(0x0012, 0x0040)

	// IssuerOfClinicalTrialSubjectID (0012,0041) VR=LO VM=1 Issuer of Clinical Trial Subject ID
	IssuerOfClinicalTrialSubjectID = New(0x0012, 0x0041)

	// ClinicalTrialSubjectReadingID (0012,0042) VR=LO VM=1 Clinical Trial Subject Reading ID
	ClinicalTrialSubjectReadingID = New(0x0012, 0x0042)

	// IssuerOfClinicalTrialSubjectReadingID (0012,0043) VR=LO VM=1 Issuer of Clinical Trial Subject Reading ID
	IssuerOfClinicalTrialSubjectReadingID = New(0x0012, 0x0043)

	// ClinicalTrialTimePointID (0012,0050) VR=LO VM=1 Clinical Trial Time Point ID
	ClinicalTrialTimePointID = New(0x0012, 0x0050)

	// ClinicalTrialTimePointDescription (0012,0051) VR=ST VM=1 Clinical Trial Time Point Description
	ClinicalTrialTimePointDescription = New(0x0012, 0x0051)

	// LongitudinalTemporalOffsetFromEvent (0012,0052) VR=FD VM=1 Longitudinal Temporal Offset from Event
	LongitudinalTemporalOffsetFromEvent = New(0x0012, 0x0052)

	// LongitudinalTemporalEventType (0012,0053) VR=CS VM=1 Longitudinal Temporal Event Type
	LongitudinalTemporalEventType = New(0x0012, 0x0053)

	// ClinicalTrialTimePointTypeCodeSequence (0012,0054) VR=SQ VM=1 Clinical Trial Time Point Type Code Sequence
	ClinicalTrialTimePointTypeCodeSequence = New(0x0012, 0x0054)

	// IssuerOfClinicalTrialTimePointID (0012,0055) VR=LO VM=1 Issuer of Clinical Trial Time Point ID
	IssuerOfClinicalTrialTimePointID = New(0x0012, 0x0055)

	// ClinicalTrialCoordinatingCenterName (0012,0060) VR=LO VM=1 Clinical Trial Coordinating Center Name
	ClinicalTrialCoordinatingCenterName = New(0x0012, 0x0060)

	// PatientIdentityRemoved (0012,0062) VR=CS VM=1 Patient Identity Removed
	PatientIdentityRemoved = New(0x0012, 0x0062)

	// DeidentificationMethod (0012,0063) VR=LO VM=1-n De-identification Method
	DeidentificationMethod = New(0x0012, 0x0063)

	// DeidentificationMethodCodeSequence (0012,0064) VR=SQ VM=1 De-identification Method Code Sequence
	DeidentificationMethodCodeSequence = New(0x0012, 0x0064)

	// ClinicalTrialSeriesID (0012,0071) VR=LO VM=1 Clinical Trial Series ID
	ClinicalTrialSeriesID = New(0x0012, 0x0071)

	// ClinicalTrialSeriesDescription (0012,0072) VR=LO VM=1 Clinical Trial Series Description
	ClinicalTrialSeriesDescription = New(0x0012, 0x0072)

	// IssuerOfClinicalTrialSeriesID (0012,0073) VR=LO VM=1 Issuer of Clinical Trial Series ID
	IssuerOfClinicalTrialSeriesID = New(0x0012, 0x0073)

	// ClinicalTrialProtocolEthicsCommitteeName (0012,0081) VR=LO VM=1 Clinical Trial Protocol Ethics Committee Name
	ClinicalTrialProtocolEthicsCommitteeName = New(0x0012, 0x0081)

	// ClinicalTrialProtocolEthicsCommitteeApprovalNumber (0012,0082) VR=LO VM=1 Clinical Trial Protocol Ethics Committee Approval Number
	ClinicalTrialProtocolEthicsCommitteeApprovalNumber = New(0x0012, 0x0082)

	// ConsentForClinicalTrialUseSequence (0012,0083) VR=SQ VM=1 Consent for Clinical Trial Use Sequence
	ConsentForClinicalTrialUseSequence = New(0x0012, 0x0083)

	// DistributionType (0012,0084) VR=CS VM=1 Distribution Type
	DistributionType = New(0x0012, 0x0084)

	// ConsentForDistributionFlag (0012,0085) VR=CS VM=1 Consent for Distribution Flag
	ConsentForDistributionFlag = New(0x0012, 0x0085)

	// EthicsCommitteeApprovalEffectivenessStartDate (0012,0086) VR=DA VM=1 Ethics Committee Approval Effectiveness Start Date
	EthicsCommitteeApprovalEffectivenessStartDate = New(0x0012, 0x0086)

	// EthicsCommitteeApprovalEffectivenessEndDate (0012,0087) VR=DA VM=1 Ethics Committee Approval Effectiveness End Date
	EthicsCommitteeApprovalEffectivenessEndDate = New(0x0012, 0x0087)

	// CADFileFormatRETIRED (0014,0023) VR=ST VM=1 CAD File Format (RETIRED)
	CADFileFormatRETIRED = New(0x0014, 0x0023)

	// ComponentReferenceSystemRETIRED (0014,0024) VR=ST VM=1 Component Reference System (RETIRED)
	ComponentReferenceSystemRETIRED = New(0x0014, 0x0024)

	// ComponentManufacturingProcedure (0014,0025) VR=ST VM=1 Component Manufacturing Procedure
	ComponentManufacturingProcedure = New(0x0014, 0x0025)

	// ComponentManufacturer (0014,0028) VR=ST VM=1 Component Manufacturer
	ComponentManufacturer = New(0x0014, 0x0028)

	// MaterialThickness (0014,0030) VR=DS VM=1-n Material Thickness
	MaterialThickness = New(0x0014, 0x0030)

	// MaterialPipeDiameter (0014,0032) VR=DS VM=1-n Material Pipe Diameter
	MaterialPipeDiameter = New(0x0014, 0x0032)

	// MaterialIsolationDiameter (0014,0034) VR=DS VM=1-n Material Isolation Diameter
	MaterialIsolationDiameter = New(0x0014, 0x0034)

	// MaterialGrade (0014,0042) VR=ST VM=1 Material Grade
	MaterialGrade = New(0x0014, 0x0042)

	// MaterialPropertiesDescription (0014,0044) VR=ST VM=1 Material Properties Description
	MaterialPropertiesDescription = New(0x0014, 0x0044)

	// MaterialPropertiesFileFormatRetiredRETIRED (0014,0045) VR=ST VM=1 Material Properties File Format (Retired) (RETIRED)
	MaterialPropertiesFileFormatRetiredRETIRED = New(0x0014, 0x0045)

	// MaterialNotes (0014,0046) VR=LT VM=1 Material Notes
	MaterialNotes = New(0x0014, 0x0046)

	// ComponentShape (0014,0050) VR=CS VM=1 Component Shape
	ComponentShape = New(0x0014, 0x0050)

	// CurvatureType (0014,0052) VR=CS VM=1 Curvature Type
	CurvatureType = New(0x0014, 0x0052)

	// OuterDiameter (0014,0054) VR=DS VM=1 Outer Diameter
	OuterDiameter = New(0x0014, 0x0054)

	// InnerDiameter (0014,0056) VR=DS VM=1 Inner Diameter
	InnerDiameter = New(0x0014, 0x0056)

	// ComponentWelderIDs (0014,0100) VR=LO VM=1-n Component Welder IDs
	ComponentWelderIDs = New(0x0014, 0x0100)

	// SecondaryApprovalStatus (0014,0101) VR=CS VM=1 Secondary Approval Status
	SecondaryApprovalStatus = New(0x0014, 0x0101)

	// SecondaryReviewDate (0014,0102) VR=DA VM=1 Secondary Review Date
	SecondaryReviewDate = New(0x0014, 0x0102)

	// SecondaryReviewTime (0014,0103) VR=TM VM=1 Secondary Review Time
	SecondaryReviewTime = New(0x0014, 0x0103)

	// SecondaryReviewerName (0014,0104) VR=PN VM=1 Secondary Reviewer Name
	SecondaryReviewerName = New(0x0014, 0x0104)

	// RepairID (0014,0105) VR=ST VM=1 Repair ID
	RepairID = New(0x0014, 0x0105)

	// MultipleComponentApprovalSequence (0014,0106) VR=SQ VM=1 Multiple Component Approval Sequence
	MultipleComponentApprovalSequence = New(0x0014, 0x0106)

	// OtherApprovalStatus (0014,0107) VR=CS VM=1-n Other Approval Status
	OtherApprovalStatus = New(0x0014, 0x0107)

	// OtherSecondaryApprovalStatus (0014,0108) VR=CS VM=1-n Other Secondary Approval Status
	OtherSecondaryApprovalStatus = New(0x0014, 0x0108)

	// DataElementLabelSequence (0014,0200) VR=SQ VM=1 Data Element Label Sequence
	DataElementLabelSequence = New(0x0014, 0x0200)

	// DataElementLabelItemSequence (0014,0201) VR=SQ VM=1 Data Element Label Item Sequence
	DataElementLabelItemSequence = New(0x0014, 0x0201)

	// DataElement (0014,0202) VR=AT VM=1 Data Element
	DataElement = New(0x0014, 0x0202)

	// DataElementName (0014,0203) VR=LO VM=1 Data Element Name
	DataElementName = New(0x0014, 0x0203)

	// DataElementDescription (0014,0204) VR=LO VM=1 Data Element Description
	DataElementDescription = New(0x0014, 0x0204)

	// DataElementConditionality (0014,0205) VR=CS VM=1 Data Element Conditionality
	DataElementConditionality = New(0x0014, 0x0205)

	// DataElementMinimumCharacters (0014,0206) VR=IS VM=1 Data Element Minimum Characters
	DataElementMinimumCharacters = New(0x0014, 0x0206)

	// DataElementMaximumCharacters (0014,0207) VR=IS VM=1 Data Element Maximum Characters
	DataElementMaximumCharacters = New(0x0014, 0x0207)

	// ActualEnvironmentalConditions (0014,1010) VR=ST VM=1 Actual Environmental Conditions
	ActualEnvironmentalConditions = New(0x0014, 0x1010)

	// ExpiryDate (0014,1020) VR=DA VM=1 Expiry Date
	ExpiryDate = New(0x0014, 0x1020)

	// EnvironmentalConditions (0014,1040) VR=ST VM=1 Environmental Conditions
	EnvironmentalConditions = New(0x0014, 0x1040)

	// EvaluatorSequence (0014,2002) VR=SQ VM=1 Evaluator Sequence
	EvaluatorSequence = New(0x0014, 0x2002)

	// EvaluatorNumber (0014,2004) VR=IS VM=1 Evaluator Number
	EvaluatorNumber = New(0x0014, 0x2004)

	// EvaluatorName (0014,2006) VR=PN VM=1 Evaluator Name
	EvaluatorName = New(0x0014, 0x2006)

	// EvaluationAttempt (0014,2008) VR=IS VM=1 Evaluation Attempt
	EvaluationAttempt = New(0x0014, 0x2008)

	// IndicationSequence (0014,2012) VR=SQ VM=1 Indication Sequence
	IndicationSequence = New(0x0014, 0x2012)

	// IndicationNumber (0014,2014) VR=IS VM=1 Indication Number
	IndicationNumber = New(0x0014, 0x2014)

	// IndicationLabel (0014,2016) VR=SH VM=1 Indication Label
	IndicationLabel = New(0x0014, 0x2016)

	// IndicationDescription (0014,2018) VR=ST VM=1 Indication Description
	IndicationDescription = New(0x0014, 0x2018)

	// IndicationType (0014,201A) VR=CS VM=1-n Indication Type
	IndicationType = New(0x0014, 0x201A)

	// IndicationDisposition (0014,201C) VR=CS VM=1 Indication Disposition
	IndicationDisposition = New(0x0014, 0x201C)

	// IndicationROISequence (0014,201E) VR=SQ VM=1 Indication ROI Sequence
	IndicationROISequence = New(0x0014, 0x201E)

	// IndicationPhysicalPropertySequence (0014,2030) VR=SQ VM=1 Indication Physical Property Sequence
	IndicationPhysicalPropertySequence = New(0x0014, 0x2030)

	// PropertyLabel (0014,2032) VR=SH VM=1 Property Label
	PropertyLabel = New(0x0014, 0x2032)

	// CoordinateSystemNumberOfAxes (0014,2202) VR=IS VM=1 Coordinate System Number of Axes
	CoordinateSystemNumberOfAxes = New(0x0014, 0x2202)

	// CoordinateSystemAxesSequence (0014,2204) VR=SQ VM=1 Coordinate System Axes Sequence
	CoordinateSystemAxesSequence = New(0x0014, 0x2204)

	// CoordinateSystemAxisDescription (0014,2206) VR=ST VM=1 Coordinate System Axis Description
	CoordinateSystemAxisDescription = New(0x0014, 0x2206)

	// CoordinateSystemDataSetMapping (0014,2208) VR=CS VM=1 Coordinate System Data Set Mapping
	CoordinateSystemDataSetMapping = New(0x0014, 0x2208)

	// CoordinateSystemAxisNumber (0014,220A) VR=IS VM=1 Coordinate System Axis Number
	CoordinateSystemAxisNumber = New(0x0014, 0x220A)

	// CoordinateSystemAxisType (0014,220C) VR=CS VM=1 Coordinate System Axis Type
	CoordinateSystemAxisType = New(0x0014, 0x220C)

	// CoordinateSystemAxisUnits (0014,220E) VR=CS VM=1 Coordinate System Axis Units
	CoordinateSystemAxisUnits = New(0x0014, 0x220E)

	// CoordinateSystemAxisValues (0014,2210) VR=OB VM=1 Coordinate System Axis Values
	CoordinateSystemAxisValues = New(0x0014, 0x2210)

	// CoordinateSystemTransformSequence (0014,2220) VR=SQ VM=1 Coordinate System Transform Sequence
	CoordinateSystemTransformSequence = New(0x0014, 0x2220)

	// TransformDescription (0014,2222) VR=ST VM=1 Transform Description
	TransformDescription = New(0x0014, 0x2222)

	// TransformNumberOfAxes (0014,2224) VR=IS VM=1 Transform Number of Axes
	TransformNumberOfAxes = New(0x0014, 0x2224)

	// TransformOrderOfAxes (0014,2226) VR=IS VM=1-n Transform Order of Axes
	TransformOrderOfAxes = New(0x0014, 0x2226)

	// TransformedAxisUnits (0014,2228) VR=CS VM=1 Transformed Axis Units
	TransformedAxisUnits = New(0x0014, 0x2228)

	// CoordinateSystemTransformRotationAndScaleMatrix (0014,222A) VR=DS VM=1-n Coordinate System Transform Rotation and Scale Matrix
	CoordinateSystemTransformRotationAndScaleMatrix = New(0x0014, 0x222A)

	// CoordinateSystemTransformTranslationMatrix (0014,222C) VR=DS VM=1-n Coordinate System Transform Translation Matrix
	CoordinateSystemTransformTranslationMatrix = New(0x0014, 0x222C)

	// InternalDetectorFrameTime (0014,3011) VR=DS VM=1 Internal Detector Frame Time
	InternalDetectorFrameTime = New(0x0014, 0x3011)

	// NumberOfFramesIntegrated (0014,3012) VR=DS VM=1 Number of Frames Integrated
	NumberOfFramesIntegrated = New(0x0014, 0x3012)

	// DetectorTemperatureSequence (0014,3020) VR=SQ VM=1 Detector Temperature Sequence
	DetectorTemperatureSequence = New(0x0014, 0x3020)

	// SensorName (0014,3022) VR=ST VM=1 Sensor Name
	SensorName = New(0x0014, 0x3022)

	// HorizontalOffsetOfSensor (0014,3024) VR=DS VM=1 Horizontal Offset of Sensor
	HorizontalOffsetOfSensor = New(0x0014, 0x3024)

	// VerticalOffsetOfSensor (0014,3026) VR=DS VM=1 Vertical Offset of Sensor
	VerticalOffsetOfSensor = New(0x0014, 0x3026)

	// SensorTemperature (0014,3028) VR=DS VM=1 Sensor Temperature
	SensorTemperature = New(0x0014, 0x3028)

	// DarkCurrentSequence (0014,3040) VR=SQ VM=1 Dark Current Sequence
	DarkCurrentSequence = New(0x0014, 0x3040)

	DarkCurrentCounts = New(0x0014, 0x3050)

	// GainCorrectionReferenceSequence (0014,3060) VR=SQ VM=1 Gain Correction Reference Sequence
	GainCorrectionReferenceSequence = New(0x0014, 0x3060)

	AirCounts = New(0x0014, 0x3070)

	// KVUsedInGainCalibration (0014,3071) VR=DS VM=1 KV Used in Gain Calibration
	KVUsedInGainCalibration = New(0x0014, 0x3071)

	// MAUsedInGainCalibration (0014,3072) VR=DS VM=1 MA Used in Gain Calibration
	MAUsedInGainCalibration = New(0x0014, 0x3072)

	// NumberOfFramesUsedForIntegration (0014,3073) VR=DS VM=1 Number of Frames Used for Integration
	NumberOfFramesUsedForIntegration = New(0x0014, 0x3073)

	// FilterMaterialUsedInGainCalibration (0014,3074) VR=LO VM=1 Filter Material Used in Gain Calibration
	FilterMaterialUsedInGainCalibration = New(0x0014, 0x3074)

	// FilterThicknessUsedInGainCalibration (0014,3075) VR=DS VM=1 Filter Thickness Used in Gain Calibration
	FilterThicknessUsedInGainCalibration = New(0x0014, 0x3075)

	// DateOfGainCalibration (0014,3076) VR=DA VM=1 Date of Gain Calibration
	DateOfGainCalibration = New(0x0014, 0x3076)

	// TimeOfGainCalibration (0014,3077) VR=TM VM=1 Time of Gain Calibration
	TimeOfGainCalibration = New(0x0014, 0x3077)

	// BadPixelImage (0014,3080) VR=OB VM=1 Bad Pixel Image
	BadPixelImage = New(0x0014, 0x3080)

	// CalibrationNotes (0014,3099) VR=LT VM=1 Calibration Notes
	CalibrationNotes = New(0x0014, 0x3099)

	// LinearityCorrectionTechnique (0014,3100) VR=LT VM=1 Linearity Correction Technique
	LinearityCorrectionTechnique = New(0x0014, 0x3100)

	// BeamHardeningCorrectionTechnique (0014,3101) VR=LT VM=1 Beam Hardening Correction Technique
	BeamHardeningCorrectionTechnique = New(0x0014, 0x3101)

	// PulserEquipmentSequence (0014,4002) VR=SQ VM=1 Pulser Equipment Sequence
	PulserEquipmentSequence = New(0x0014, 0x4002)

	// PulserType (0014,4004) VR=CS VM=1 Pulser Type
	PulserType = New(0x0014, 0x4004)

	// PulserNotes (0014,4006) VR=LT VM=1 Pulser Notes
	PulserNotes = New(0x0014, 0x4006)

	// ReceiverEquipmentSequence (0014,4008) VR=SQ VM=1 Receiver Equipment Sequence
	ReceiverEquipmentSequence = New(0x0014, 0x4008)

	// AmplifierType (0014,400A) VR=CS VM=1 Amplifier Type
	AmplifierType = New(0x0014, 0x400A)

	// ReceiverNotes (0014,400C) VR=LT VM=1 Receiver Notes
	ReceiverNotes = New(0x0014, 0x400C)

	// PreAmplifierEquipmentSequence (0014,400E) VR=SQ VM=1 Pre-Amplifier Equipment Sequence
	PreAmplifierEquipmentSequence = New(0x0014, 0x400E)

	// PreAmplifierNotes (0014,400F) VR=LT VM=1 Pre-Amplifier Notes
	PreAmplifierNotes = New(0x0014, 0x400F)

	// TransmitTransducerSequence (0014,4010) VR=SQ VM=1 Transmit Transducer Sequence
	TransmitTransducerSequence = New(0x0014, 0x4010)

	// ReceiveTransducerSequence (0014,4011) VR=SQ VM=1 Receive Transducer Sequence
	ReceiveTransducerSequence = New(0x0014, 0x4011)

	// NumberOfElements (0014,4012) VR=US VM=1 Number of Elements
	NumberOfElements = New(0x0014, 0x4012)

	// ElementShape (0014,4013) VR=CS VM=1 Element Shape
	ElementShape = New(0x0014, 0x4013)

	// ElementDimensionA (0014,4014) VR=DS VM=1 Element Dimension A
	ElementDimensionA = New(0x0014, 0x4014)

	// ElementDimensionB (0014,4015) VR=DS VM=1 Element Dimension B
	ElementDimensionB = New(0x0014, 0x4015)

	// ElementPitchA (0014,4016) VR=DS VM=1 Element Pitch A
	ElementPitchA = New(0x0014, 0x4016)

	// MeasuredBeamDimensionA (0014,4017) VR=DS VM=1 Measured Beam Dimension A
	MeasuredBeamDimensionA = New(0x0014, 0x4017)

	// MeasuredBeamDimensionB (0014,4018) VR=DS VM=1 Measured Beam Dimension B
	MeasuredBeamDimensionB = New(0x0014, 0x4018)

	// LocationOfMeasuredBeamDiameter (0014,4019) VR=DS VM=1 Location of Measured Beam Diameter
	LocationOfMeasuredBeamDiameter = New(0x0014, 0x4019)

	// NominalFrequency (0014,401A) VR=DS VM=1 Nominal Frequency
	NominalFrequency = New(0x0014, 0x401A)

	// MeasuredCenterFrequency (0014,401B) VR=DS VM=1 Measured Center Frequency
	MeasuredCenterFrequency = New(0x0014, 0x401B)

	// MeasuredBandwidth (0014,401C) VR=DS VM=1 Measured Bandwidth
	MeasuredBandwidth = New(0x0014, 0x401C)

	// ElementPitchB (0014,401D) VR=DS VM=1 Element Pitch B
	ElementPitchB = New(0x0014, 0x401D)

	// PulserSettingsSequence (0014,4020) VR=SQ VM=1 Pulser Settings Sequence
	PulserSettingsSequence = New(0x0014, 0x4020)

	// PulseWidth (0014,4022) VR=DS VM=1 Pulse Width
	PulseWidth = New(0x0014, 0x4022)

	// ExcitationFrequency (0014,4024) VR=DS VM=1 Excitation Frequency
	ExcitationFrequency = New(0x0014, 0x4024)

	// ModulationType (0014,4026) VR=CS VM=1 Modulation Type
	ModulationType = New(0x0014, 0x4026)

	// Damping (0014,4028) VR=DS VM=1 Damping
	Damping = New(0x0014, 0x4028)

	// ReceiverSettingsSequence (0014,4030) VR=SQ VM=1 Receiver Settings Sequence
	ReceiverSettingsSequence = New(0x0014, 0x4030)

	// AcquiredSoundpathLength (0014,4031) VR=DS VM=1 Acquired Soundpath Length
	AcquiredSoundpathLength = New(0x0014, 0x4031)

	// AcquisitionCompressionType (0014,4032) VR=CS VM=1 Acquisition Compression Type
	AcquisitionCompressionType = New(0x0014, 0x4032)

	// AcquisitionSampleSize (0014,4033) VR=IS VM=1 Acquisition Sample Size
	AcquisitionSampleSize = New(0x0014, 0x4033)

	// RectifierSmoothing (0014,4034) VR=DS VM=1 Rectifier Smoothing
	RectifierSmoothing = New(0x0014, 0x4034)

	// DACSequence (0014,4035) VR=SQ VM=1 DAC Sequence
	DACSequence = New(0x0014, 0x4035)

	// DACType (0014,4036) VR=CS VM=1 DAC Type
	DACType = New(0x0014, 0x4036)

	// DACGainPoints (0014,4038) VR=DS VM=1-n DAC Gain Points
	DACGainPoints = New(0x0014, 0x4038)

	// DACTimePoints (0014,403A) VR=DS VM=1-n DAC Time Points
	DACTimePoints = New(0x0014, 0x403A)

	// DACAmplitude (0014,403C) VR=DS VM=1-n DAC Amplitude
	DACAmplitude = New(0x0014, 0x403C)

	// PreAmplifierSettingsSequence (0014,4040) VR=SQ VM=1 Pre-Amplifier Settings Sequence
	PreAmplifierSettingsSequence = New(0x0014, 0x4040)

	// TransmitTransducerSettingsSequence (0014,4050) VR=SQ VM=1 Transmit Transducer Settings Sequence
	TransmitTransducerSettingsSequence = New(0x0014, 0x4050)

	// ReceiveTransducerSettingsSequence (0014,4051) VR=SQ VM=1 Receive Transducer Settings Sequence
	ReceiveTransducerSettingsSequence = New(0x0014, 0x4051)

	// IncidentAngle (0014,4052) VR=DS VM=1 Incident Angle
	IncidentAngle = New(0x0014, 0x4052)

	// CouplingTechnique (0014,4054) VR=ST VM=1 Coupling Technique
	CouplingTechnique = New(0x0014, 0x4054)

	// CouplingMedium (0014,4056) VR=ST VM=1 Coupling Medium
	CouplingMedium = New(0x0014, 0x4056)

	// CouplingVelocity (0014,4057) VR=DS VM=1 Coupling Velocity
	CouplingVelocity = New(0x0014, 0x4057)

	// ProbeCenterLocationX (0014,4058) VR=DS VM=1 Probe Center Location X
	ProbeCenterLocationX = New(0x0014, 0x4058)

	// ProbeCenterLocationZ (0014,4059) VR=DS VM=1 Probe Center Location Z
	ProbeCenterLocationZ = New(0x0014, 0x4059)

	// SoundPathLength (0014,405A) VR=DS VM=1 Sound Path Length
	SoundPathLength = New(0x0014, 0x405A)

	// DelayLawIdentifier (0014,405C) VR=ST VM=1 Delay Law Identifier
	DelayLawIdentifier = New(0x0014, 0x405C)

	// GateSettingsSequence (0014,4060) VR=SQ VM=1 Gate Settings Sequence
	GateSettingsSequence = New(0x0014, 0x4060)

	// GateThreshold (0014,4062) VR=DS VM=1 Gate Threshold
	GateThreshold = New(0x0014, 0x4062)

	// VelocityOfSound (0014,4064) VR=DS VM=1 Velocity of Sound
	VelocityOfSound = New(0x0014, 0x4064)

	// CalibrationSettingsSequence (0014,4070) VR=SQ VM=1 Calibration Settings Sequence
	CalibrationSettingsSequence = New(0x0014, 0x4070)

	// CalibrationProcedure (0014,4072) VR=ST VM=1 Calibration Procedure
	CalibrationProcedure = New(0x0014, 0x4072)

	// ProcedureVersion (0014,4074) VR=SH VM=1 Procedure Version
	ProcedureVersion = New(0x0014, 0x4074)

	// ProcedureCreationDate (0014,4076) VR=DA VM=1 Procedure Creation Date
	ProcedureCreationDate = New(0x0014, 0x4076)

	// ProcedureExpirationDate (0014,4078) VR=DA VM=1 Procedure Expiration Date
	ProcedureExpirationDate = New(0x0014, 0x4078)

	// ProcedureLastModifiedDate (0014,407A) VR=DA VM=1 Procedure Last Modified Date
	ProcedureLastModifiedDate = New(0x0014, 0x407A)

	// CalibrationTime (0014,407C) VR=TM VM=1-n Calibration Time
	CalibrationTime = New(0x0014, 0x407C)

	// CalibrationDate (0014,407E) VR=DA VM=1-n Calibration Date
	CalibrationDate = New(0x0014, 0x407E)

	// ProbeDriveEquipmentSequence (0014,4080) VR=SQ VM=1 Probe Drive Equipment Sequence
	ProbeDriveEquipmentSequence = New(0x0014, 0x4080)

	// DriveType (0014,4081) VR=CS VM=1 Drive Type
	DriveType = New(0x0014, 0x4081)

	// ProbeDriveNotes (0014,4082) VR=LT VM=1 Probe Drive Notes
	ProbeDriveNotes = New(0x0014, 0x4082)

	// DriveProbeSequence (0014,4083) VR=SQ VM=1 Drive Probe Sequence
	DriveProbeSequence = New(0x0014, 0x4083)

	// ProbeInductance (0014,4084) VR=DS VM=1 Probe Inductance
	ProbeInductance = New(0x0014, 0x4084)

	// ProbeResistance (0014,4085) VR=DS VM=1 Probe Resistance
	ProbeResistance = New(0x0014, 0x4085)

	// ReceiveProbeSequence (0014,4086) VR=SQ VM=1 Receive Probe Sequence
	ReceiveProbeSequence = New(0x0014, 0x4086)

	// ProbeDriveSettingsSequence (0014,4087) VR=SQ VM=1 Probe Drive Settings Sequence
	ProbeDriveSettingsSequence = New(0x0014, 0x4087)

	// BridgeResistors (0014,4088) VR=DS VM=1 Bridge Resistors
	BridgeResistors = New(0x0014, 0x4088)

	// ProbeOrientationAngle (0014,4089) VR=DS VM=1 Probe Orientation Angle
	ProbeOrientationAngle = New(0x0014, 0x4089)

	// UserSelectedGainY (0014,408B) VR=DS VM=1 User Selected Gain Y
	UserSelectedGainY = New(0x0014, 0x408B)

	// UserSelectedPhase (0014,408C) VR=DS VM=1 User Selected Phase
	UserSelectedPhase = New(0x0014, 0x408C)

	// UserSelectedOffsetX (0014,408D) VR=DS VM=1 User Selected Offset X
	UserSelectedOffsetX = New(0x0014, 0x408D)

	// UserSelectedOffsetY (0014,408E) VR=DS VM=1 User Selected Offset Y
	UserSelectedOffsetY = New(0x0014, 0x408E)

	// ChannelSettingsSequence (0014,4091) VR=SQ VM=1 Channel Settings Sequence
	ChannelSettingsSequence = New(0x0014, 0x4091)

	// ChannelThreshold (0014,4092) VR=DS VM=1 Channel Threshold
	ChannelThreshold = New(0x0014, 0x4092)

	// ScannerSettingsSequence (0014,409A) VR=SQ VM=1 Scanner Settings Sequence
	ScannerSettingsSequence = New(0x0014, 0x409A)

	// ScanProcedure (0014,409B) VR=ST VM=1 Scan Procedure
	ScanProcedure = New(0x0014, 0x409B)

	// TranslationRateX (0014,409C) VR=DS VM=1 Translation Rate X
	TranslationRateX = New(0x0014, 0x409C)

	// TranslationRateY (0014,409D) VR=DS VM=1 Translation Rate Y
	TranslationRateY = New(0x0014, 0x409D)

	// ChannelOverlap (0014,409F) VR=DS VM=1 Channel Overlap
	ChannelOverlap = New(0x0014, 0x409F)

	// ImageQualityIndicatorType (0014,40A0) VR=LO VM=1-n Image Quality Indicator Type
	ImageQualityIndicatorType = New(0x0014, 0x40A0)

	// ImageQualityIndicatorMaterial (0014,40A1) VR=LO VM=1-n Image Quality Indicator Material
	ImageQualityIndicatorMaterial = New(0x0014, 0x40A1)

	// ImageQualityIndicatorSize (0014,40A2) VR=LO VM=1-n Image Quality Indicator Size
	ImageQualityIndicatorSize = New(0x0014, 0x40A2)

	// WaveDimensionsDefinitionSequence (0014,4101) VR=SQ VM=1 Wave Dimensions Definition Sequence
	WaveDimensionsDefinitionSequence = New(0x0014, 0x4101)

	// WaveDimensionNumber (0014,4102) VR=US VM=1 Wave Dimension Number
	WaveDimensionNumber = New(0x0014, 0x4102)

	// WaveDimensionDescription (0014,4103) VR=LO VM=1 Wave Dimension Description
	WaveDimensionDescription = New(0x0014, 0x4103)

	// WaveDimensionUnit (0014,4104) VR=US VM=1 Wave Dimension Unit
	WaveDimensionUnit = New(0x0014, 0x4104)

	// WaveDimensionValueType (0014,4105) VR=CS VM=1 Wave Dimension Value Type
	WaveDimensionValueType = New(0x0014, 0x4105)

	// WaveDimensionValuesSequence (0014,4106) VR=SQ VM=1-n Wave Dimension Values Sequence
	WaveDimensionValuesSequence = New(0x0014, 0x4106)

	// ReferencedWaveDimension (0014,4107) VR=US VM=1 Referenced Wave Dimension
	ReferencedWaveDimension = New(0x0014, 0x4107)

	// IntegerNumericValue (0014,4108) VR=SL VM=1 Integer Numeric Value
	IntegerNumericValue = New(0x0014, 0x4108)

	// ByteNumericValue (0014,4109) VR=OB VM=1 Byte Numeric Value
	ByteNumericValue = New(0x0014, 0x4109)

	// ShortNumericValue (0014,410A) VR=OW VM=1 Short Numeric Value
	ShortNumericValue = New(0x0014, 0x410A)

	// SinglePrecisionFloatingPointNumericValue (0014,410B) VR=OF VM=1 Single Precision Floating Point Numeric Value
	SinglePrecisionFloatingPointNumericValue = New(0x0014, 0x410B)

	// DoublePrecisionFloatingPointNumericValue (0014,410C) VR=OD VM=1 Double Precision Floating Point Numeric Value
	DoublePrecisionFloatingPointNumericValue = New(0x0014, 0x410C)

	// LINACEnergy (0014,5002) VR=IS VM=1 LINAC Energy
	LINACEnergy = New(0x0014, 0x5002)

	// LINACOutput (0014,5004) VR=IS VM=1 LINAC Output
	LINACOutput = New(0x0014, 0x5004)

	// ActiveAperture (0014,5100) VR=US VM=1 Active Aperture
	ActiveAperture = New(0x0014, 0x5100)

	// TotalAperture (0014,5101) VR=DS VM=1 Total Aperture
	TotalAperture = New(0x0014, 0x5101)

	// ApertureElevation (0014,5102) VR=DS VM=1 Aperture Elevation
	ApertureElevation = New(0x0014, 0x5102)

	// MainLobeAngle (0014,5103) VR=DS VM=1 Main Lobe Angle
	MainLobeAngle = New(0x0014, 0x5103)

	// MainRoofAngle (0014,5104) VR=DS VM=1 Main Roof Angle
	MainRoofAngle = New(0x0014, 0x5104)

	// ConnectorType (0014,5105) VR=CS VM=1 Connector Type
	ConnectorType = New(0x0014, 0x5105)

	// WedgeModelNumber (0014,5106) VR=SH VM=1 Wedge Model Number
	WedgeModelNumber = New(0x0014, 0x5106)

	// WedgeAngleFloat (0014,5107) VR=DS VM=1 Wedge Angle Float
	WedgeAngleFloat = New(0x0014, 0x5107)

	// WedgeRoofAngle (0014,5108) VR=DS VM=1 Wedge Roof Angle
	WedgeRoofAngle = New(0x0014, 0x5108)

	// WedgeElement1Position (0014,5109) VR=CS VM=1 Wedge Element 1 Position
	WedgeElement1Position = New(0x0014, 0x5109)

	// WedgeMaterialVelocity (0014,510A) VR=DS VM=1 Wedge Material Velocity
	WedgeMaterialVelocity = New(0x0014, 0x510A)

	// WedgeMaterial (0014,510B) VR=SH VM=1 Wedge Material
	WedgeMaterial = New(0x0014, 0x510B)

	// WedgeOffsetZ (0014,510C) VR=DS VM=1 Wedge Offset Z
	WedgeOffsetZ = New(0x0014, 0x510C)

	// WedgeOriginOffsetX (0014,510D) VR=DS VM=1 Wedge Origin Offset X
	WedgeOriginOffsetX = New(0x0014, 0x510D)

	// WedgeTimeDelay (0014,510E) VR=DS VM=1 Wedge Time Delay
	WedgeTimeDelay = New(0x0014, 0x510E)

	// WedgeName (0014,510F) VR=SH VM=1 Wedge Name
	WedgeName = New(0x0014, 0x510F)

	// WedgeManufacturerName (0014,5110) VR=SH VM=1 Wedge Manufacturer Name
	WedgeManufacturerName = New(0x0014, 0x5110)

	// WedgeDescription (0014,5111) VR=LO VM=1 Wedge Description
	WedgeDescription = New(0x0014, 0x5111)

	// NominalBeamAngle (0014,5112) VR=DS VM=1 Nominal Beam Angle
	NominalBeamAngle = New(0x0014, 0x5112)

	// WedgeOffsetX (0014,5113) VR=DS VM=1 Wedge Offset X
	WedgeOffsetX = New(0x0014, 0x5113)

	// WedgeOffsetY (0014,5114) VR=DS VM=1 Wedge Offset Y
	WedgeOffsetY = New(0x0014, 0x5114)

	// WedgeTotalLength (0014,5115) VR=DS VM=1 Wedge Total Length
	WedgeTotalLength = New(0x0014, 0x5115)

	// WedgeInContactLength (0014,5116) VR=DS VM=1 Wedge In Contact Length
	WedgeInContactLength = New(0x0014, 0x5116)

	// WedgeFrontGap (0014,5117) VR=DS VM=1 Wedge Front Gap
	WedgeFrontGap = New(0x0014, 0x5117)

	// WedgeTotalHeight (0014,5118) VR=DS VM=1 Wedge Total Height
	WedgeTotalHeight = New(0x0014, 0x5118)

	// WedgeFrontHeight (0014,5119) VR=DS VM=1 Wedge Front Height
	WedgeFrontHeight = New(0x0014, 0x5119)

	// WedgeRearHeight (0014,511A) VR=DS VM=1 Wedge Rear Height
	WedgeRearHeight = New(0x0014, 0x511A)

	// WedgeTotalWidth (0014,511B) VR=DS VM=1 Wedge Total Width
	WedgeTotalWidth = New(0x0014, 0x511B)

	// WedgeInContactWidth (0014,511C) VR=DS VM=1 Wedge In Contact Width
	WedgeInContactWidth = New(0x0014, 0x511C)

	// WedgeChamferHeight (0014,511D) VR=DS VM=1 Wedge Chamfer Height
	WedgeChamferHeight = New(0x0014, 0x511D)

	// WedgeCurve (0014,511E) VR=CS VM=1 Wedge Curve
	WedgeCurve = New(0x0014, 0x511E)

	// RadiusAlongWedge (0014,511F) VR=DS VM=1 Radius Along the Wedge
	RadiusAlongWedge = New(0x0014, 0x511F)

	// ThermalCameraSettingsSequence (0014,6001) VR=SQ VM=1 Thermal Camera Settings Sequence
	ThermalCameraSettingsSequence = New(0x0014, 0x6001)

	// AcquisitionFrameRate (0014,6002) VR=DS VM=1 Acquisition Frame Rate
	AcquisitionFrameRate = New(0x0014, 0x6002)

	// IntegrationTime (0014,6003) VR=DS VM=1 Integration Time
	IntegrationTime = New(0x0014, 0x6003)

	// NumberOfCalibrationFrames (0014,6004) VR=DS VM=1 Number of Calibration Frames
	NumberOfCalibrationFrames = New(0x0014, 0x6004)

	// NumberOfRowsInFullAcquisitionImage (0014,6005) VR=DS VM=1 Number of Rows in Full Acquisition Image
	NumberOfRowsInFullAcquisitionImage = New(0x0014, 0x6005)

	// NumberOfColumnsInFullAcquisitionImage (0014,6006) VR=DS VM=1 Number Of Columns in Full Acquisition Image
	NumberOfColumnsInFullAcquisitionImage = New(0x0014, 0x6006)

	// ThermalSourceSettingsSequence (0014,6007) VR=SQ VM=1 Thermal Source Settings Sequence
	ThermalSourceSettingsSequence = New(0x0014, 0x6007)

	// SourceHorizontalPitch (0014,6008) VR=DS VM=1 Source Horizontal Pitch
	SourceHorizontalPitch = New(0x0014, 0x6008)

	// SourceVerticalPitch (0014,6009) VR=DS VM=1 Source Vertical Pitch
	SourceVerticalPitch = New(0x0014, 0x6009)

	// SourceHorizontalScanSpeed (0014,600A) VR=DS VM=1 Source Horizontal Scan Speed
	SourceHorizontalScanSpeed = New(0x0014, 0x600A)

	// ThermalSourceModulationFrequency (0014,600B) VR=DS VM=1 Thermal Source Modulation Frequency
	ThermalSourceModulationFrequency = New(0x0014, 0x600B)

	// InductionSourceSettingSequence (0014,600C) VR=SQ VM=1 Induction Source Setting Sequence
	InductionSourceSettingSequence = New(0x0014, 0x600C)

	// CoilFrequency (0014,600D) VR=DS VM=1 Coil Frequency
	CoilFrequency = New(0x0014, 0x600D)

	// CurrentAmplitudeAcrossCoil (0014,600E) VR=DS VM=1 Current Amplitude Across Coil
	CurrentAmplitudeAcrossCoil = New(0x0014, 0x600E)

	// FlashSourceSettingSequence (0014,600F) VR=SQ VM=1 Flash Source Setting Sequence
	FlashSourceSettingSequence = New(0x0014, 0x600F)

	// FlashDuration (0014,6010) VR=DS VM=1 Flash Duration
	FlashDuration = New(0x0014, 0x6010)

	// FlashFrameNumber (0014,6011) VR=DS VM=1-n Flash Frame Number
	FlashFrameNumber = New(0x0014, 0x6011)

	// LaserSourceSettingSequence (0014,6012) VR=SQ VM=1 Laser Source Setting Sequence
	LaserSourceSettingSequence = New(0x0014, 0x6012)

	// HorizontalLaserSpotDimension (0014,6013) VR=DS VM=1 Horizontal Laser Spot Dimension
	HorizontalLaserSpotDimension = New(0x0014, 0x6013)

	// VerticalLaserSpotDimension (0014,6014) VR=DS VM=1 Vertical Laser Spot Dimension
	VerticalLaserSpotDimension = New(0x0014, 0x6014)

	// LaserWavelength (0014,6015) VR=DS VM=1 Laser Wavelength
	LaserWavelength = New(0x0014, 0x6015)

	// LaserPower (0014,6016) VR=DS VM=1 Laser Power
	LaserPower = New(0x0014, 0x6016)

	// ForcedGasSettingSequence (0014,6017) VR=SQ VM=1 Forced Gas Setting Sequence
	ForcedGasSettingSequence = New(0x0014, 0x6017)

	// VibrationSourceSettingSequence (0014,6018) VR=SQ VM=1 Vibration Source Setting Sequence
	VibrationSourceSettingSequence = New(0x0014, 0x6018)

	// VibrationExcitationFrequency (0014,6019) VR=DS VM=1 Vibration Excitation Frequency
	VibrationExcitationFrequency = New(0x0014, 0x6019)

	// VibrationExcitationVoltage (0014,601A) VR=DS VM=1 Vibration Excitation Voltage
	VibrationExcitationVoltage = New(0x0014, 0x601A)

	// ThermographyDataCaptureMethod (0014,601B) VR=CS VM=1 Thermography Data Capture Method
	ThermographyDataCaptureMethod = New(0x0014, 0x601B)

	// ThermalTechnique (0014,601C) VR=CS VM=1 Thermal Technique
	ThermalTechnique = New(0x0014, 0x601C)

	// ThermalCameraCoreSequence (0014,601D) VR=SQ VM=1 Thermal Camera Core Sequence
	ThermalCameraCoreSequence = New(0x0014, 0x601D)

	// DetectorWavelengthRange (0014,601E) VR=CS VM=1 Detector Wavelength Range
	DetectorWavelengthRange = New(0x0014, 0x601E)

	// ThermalCameraCalibrationType (0014,601F) VR=CS VM=1 Thermal Camera Calibration Type
	ThermalCameraCalibrationType = New(0x0014, 0x601F)

	// AcquisitionImageCounter (0014,6020) VR=UV VM=1 Acquisition Image Counter
	AcquisitionImageCounter = New(0x0014, 0x6020)

	// FrontPanelTemperature (0014,6021) VR=DS VM=1 Front Panel Temperature
	FrontPanelTemperature = New(0x0014, 0x6021)

	// AirGapTemperature (0014,6022) VR=DS VM=1 Air Gap Temperature
	AirGapTemperature = New(0x0014, 0x6022)

	// VerticalPixelSize (0014,6023) VR=DS VM=1 Vertical Pixel Size
	VerticalPixelSize = New(0x0014, 0x6023)

	// HorizontalPixelSize (0014,6024) VR=DS VM=1 Horizontal Pixel Size
	HorizontalPixelSize = New(0x0014, 0x6024)

	// DataStreamingProtocol (0014,6025) VR=ST VM=1-n Data Streaming Protocol
	DataStreamingProtocol = New(0x0014, 0x6025)

	// LensSequence (0014,6026) VR=SQ VM=1 Lens Sequence
	LensSequence = New(0x0014, 0x6026)

	// FieldOfView (0014,6027) VR=DS VM=1 Field of View
	FieldOfView = New(0x0014, 0x6027)

	// LensFilterManufacturer (0014,6028) VR=LO VM=1 Lens Filter Manufacturer
	LensFilterManufacturer = New(0x0014, 0x6028)

	// CutoffFilterType (0014,6029) VR=CS VM=1 Cutoff Filter Type
	CutoffFilterType = New(0x0014, 0x6029)

	// LensFilterCutOffWavelength (0014,602A) VR=DS VM=1-n Lens Filter Cut-Off Wavelength
	LensFilterCutOffWavelength = New(0x0014, 0x602A)

	// ThermalSourceSequence (0014,602B) VR=SQ VM=1 Thermal Source Sequence
	ThermalSourceSequence = New(0x0014, 0x602B)

	// ThermalSourceMotionState (0014,602C) VR=CS VM=1 Thermal Source Motion State
	ThermalSourceMotionState = New(0x0014, 0x602C)

	// ThermalSourceMotionType (0014,602D) VR=CS VM=1 Thermal Source Motion Type
	ThermalSourceMotionType = New(0x0014, 0x602D)

	// InductionHeatingSequence (0014,602E) VR=SQ VM=1 Induction Heating Sequence
	InductionHeatingSequence = New(0x0014, 0x602E)

	// CoilConfigurationID (0014,602F) VR=ST VM=1 Coil Configuration ID
	CoilConfigurationID = New(0x0014, 0x602F)

	// NumberOfTurnsInCoil (0014,6030) VR=DS VM=1 Number of Turns in Coil
	NumberOfTurnsInCoil = New(0x0014, 0x6030)

	// ShapeOfIndividualTurn (0014,6031) VR=CS VM=1 Shape of Individual Turn
	ShapeOfIndividualTurn = New(0x0014, 0x6031)

	// SizeOfIndividualTurn (0014,6032) VR=DS VM=1-n Size of Individual Turn
	SizeOfIndividualTurn = New(0x0014, 0x6032)

	// DistanceBetweenTurns (0014,6033) VR=DS VM=1-n Distance Between Turns
	DistanceBetweenTurns = New(0x0014, 0x6033)

	// FlashHeatingSequence (0014,6034) VR=SQ VM=1 Flash Heating Sequence
	FlashHeatingSequence = New(0x0014, 0x6034)

	// NumberOfLamps (0014,6035) VR=DS VM=1 Number of Lamps
	NumberOfLamps = New(0x0014, 0x6035)

	// FlashSynchronizationProtocol (0014,6036) VR=ST VM=1 Flash Synchronization Protocol
	FlashSynchronizationProtocol = New(0x0014, 0x6036)

	// FlashModificationStatus (0014,6037) VR=CS VM=1 Flash Modification Status
	FlashModificationStatus = New(0x0014, 0x6037)

	// LaserHeatingSequence (0014,6038) VR=SQ VM=1 Laser Heating Sequence
	LaserHeatingSequence = New(0x0014, 0x6038)

	// LaserManufacturer (0014,6039) VR=LO VM=1 Laser Manufacturer
	LaserManufacturer = New(0x0014, 0x6039)

	// LaserModelNumber (0014,603A) VR=LO VM=1 Laser Model Number
	LaserModelNumber = New(0x0014, 0x603A)

	// LaserTypeDescription (0014,603B) VR=ST VM=1 Laser Type Description
	LaserTypeDescription = New(0x0014, 0x603B)

	// ForcedGasHeatingSequence (0014,603C) VR=SQ VM=1 Forced Gas Heating Sequence
	ForcedGasHeatingSequence = New(0x0014, 0x603C)

	// GasUsedForHeatingCoolingPart (0014,603D) VR=LO VM=1 Gas Used for Heating/Cooling Part
	GasUsedForHeatingCoolingPart = New(0x0014, 0x603D)

	// VibrationSonicHeatingSequence (0014,603E) VR=SQ VM=1 Vibration/Sonic Heating Sequence
	VibrationSonicHeatingSequence = New(0x0014, 0x603E)

	// ProbeManufacturer (0014,603F) VR=LO VM=1 Probe Manufacturer
	ProbeManufacturer = New(0x0014, 0x603F)

	// ProbeModelNumber (0014,6040) VR=LO VM=1 Probe Model Number
	ProbeModelNumber = New(0x0014, 0x6040)

	// ApertureSize (0014,6041) VR=DS VM=1 Aperture Size
	ApertureSize = New(0x0014, 0x6041)

	// ProbeResonantFrequency (0014,6042) VR=DS VM=1 Probe Resonant Frequency
	ProbeResonantFrequency = New(0x0014, 0x6042)

	// HeatSourceDescription (0014,6043) VR=UT VM=1 Heat Source Description
	HeatSourceDescription = New(0x0014, 0x6043)

	// SurfacePreparationWithOpticalCoating (0014,6044) VR=CS VM=1 Surface Preparation with Optical Coating
	SurfacePreparationWithOpticalCoating = New(0x0014, 0x6044)

	// OpticalCoatingType (0014,6045) VR=ST VM=1 Optical Coating Type
	OpticalCoatingType = New(0x0014, 0x6045)

	// ThermalConductivityOfExposedSurface (0014,6046) VR=DS VM=1 Thermal Conductivity of Exposed Surface
	ThermalConductivityOfExposedSurface = New(0x0014, 0x6046)

	// MaterialDensity (0014,6047) VR=DS VM=1 Material Density
	MaterialDensity = New(0x0014, 0x6047)

	// SpecificHeatOfInspectionSurface (0014,6048) VR=DS VM=1 Specific Heat of Inspection Surface
	SpecificHeatOfInspectionSurface = New(0x0014, 0x6048)

	// EmissivityOfInspectionSurface (0014,6049) VR=DS VM=1 Emissivity of Inspection Surface
	EmissivityOfInspectionSurface = New(0x0014, 0x6049)

	// ElectromagneticClassificationOfInspectionSurface (0014,604A) VR=CS VM=1-n Electromagnetic Classification of Inspection Surface
	ElectromagneticClassificationOfInspectionSurface = New(0x0014, 0x604A)

	// MovingWindowSize (0014,604C) VR=DS VM=1 Moving Window Size
	MovingWindowSize = New(0x0014, 0x604C)

	// MovingWindowType (0014,604D) VR=CS VM=1 Moving Window Type
	MovingWindowType = New(0x0014, 0x604D)

	// MovingWindowWeights (0014,604E) VR=DS VM=1-n Moving Window Weights
	MovingWindowWeights = New(0x0014, 0x604E)

	// MovingWindowPitch (0014,604F) VR=DS VM=1 Moving Window Pitch
	MovingWindowPitch = New(0x0014, 0x604F)

	// MovingWindowPaddingScheme (0014,6050) VR=CS VM=1 Moving Window Padding Scheme
	MovingWindowPaddingScheme = New(0x0014, 0x6050)

	// MovingWindowPaddingLength (0014,6051) VR=DS VM=1 Moving Window Padding Sength
	MovingWindowPaddingLength = New(0x0014, 0x6051)

	// SpatialFilteringParametersSequence (0014,6052) VR=SQ VM=1 Spatial Filtering Parameters Sequence
	SpatialFilteringParametersSequence = New(0x0014, 0x6052)

	// SpatialFilteringScheme (0014,6053) VR=CS VM=1 Spatial Filtering Scheme
	SpatialFilteringScheme = New(0x0014, 0x6053)

	// HorizontalMovingWindowSize (0014,6056) VR=DS VM=1 Horizontal Moving Window Size
	HorizontalMovingWindowSize = New(0x0014, 0x6056)

	// VerticalMovingWindowSize (0014,6057) VR=DS VM=1 Vertical Moving Window Size
	VerticalMovingWindowSize = New(0x0014, 0x6057)

	// PolynomialFittingSequence (0014,6059) VR=SQ VM=1 Polynomial Fitting Sequence
	PolynomialFittingSequence = New(0x0014, 0x6059)

	// FittingDataType (0014,605A) VR=CS VM=1-n Fitting Data Type
	FittingDataType = New(0x0014, 0x605A)

	// OperationOnTimeAxisBeforeFitting (0014,605B) VR=CS VM=1 Operation on Time Axis Before Fitting
	OperationOnTimeAxisBeforeFitting = New(0x0014, 0x605B)

	// OperationOnPixelIntensityBeforeFitting (0014,605C) VR=CS VM=1 Operation on Pixel Intensity Before Fitting
	OperationOnPixelIntensityBeforeFitting = New(0x0014, 0x605C)

	// OrderOfPolynomial (0014,605D) VR=DS VM=1 Order of Polynomial
	OrderOfPolynomial = New(0x0014, 0x605D)

	// IndependentVariableForPolynomialFit (0014,605E) VR=CS VM=1 Independent Variable for Polynomial Fit
	IndependentVariableForPolynomialFit = New(0x0014, 0x605E)

	// PolynomialCoefficients (0014,605F) VR=DS VM=1-n PolynomialCoefficients
	PolynomialCoefficients = New(0x0014, 0x605F)

	// ThermographyPixelDataUnit (0014,6060) VR=CS VM=1 Thermography Pixel Data Unit
	ThermographyPixelDataUnit = New(0x0014, 0x6060)

	// WhitePoint (0016,0001) VR=DS VM=1 White Point
	WhitePoint = New(0x0016, 0x0001)

	// PrimaryChromaticities (0016,0002) VR=DS VM=3 Primary Chromaticities
	PrimaryChromaticities = New(0x0016, 0x0002)

	// BatteryLevel (0016,0003) VR=UT VM=1 Battery Level
	BatteryLevel = New(0x0016, 0x0003)

	// ExposureTimeInSeconds (0016,0004) VR=DS VM=1 Exposure Time in Seconds
	ExposureTimeInSeconds = New(0x0016, 0x0004)

	// FNumber (0016,0005) VR=DS VM=1 F-Number
	FNumber = New(0x0016, 0x0005)

	// OECFRows (0016,0006) VR=IS VM=1 OECF Rows
	OECFRows = New(0x0016, 0x0006)

	// OECFColumns (0016,0007) VR=IS VM=1 OECF Columns
	OECFColumns = New(0x0016, 0x0007)

	// OECFColumnNames (0016,0008) VR=UC VM=1-n OECF Column Names
	OECFColumnNames = New(0x0016, 0x0008)

	// OECFValues (0016,0009) VR=DS VM=1-n OECF Values
	OECFValues = New(0x0016, 0x0009)

	// SpatialFrequencyResponseRows (0016,000A) VR=IS VM=1 Spatial Frequency Response Rows
	SpatialFrequencyResponseRows = New(0x0016, 0x000A)

	// SpatialFrequencyResponseColumns (0016,000B) VR=IS VM=1 Spatial Frequency Response Columns
	SpatialFrequencyResponseColumns = New(0x0016, 0x000B)

	// SpatialFrequencyResponseColumnNames (0016,000C) VR=UC VM=1-n Spatial Frequency Response Column Names
	SpatialFrequencyResponseColumnNames = New(0x0016, 0x000C)

	// SpatialFrequencyResponseValues (0016,000D) VR=DS VM=1-n Spatial Frequency Response Values
	SpatialFrequencyResponseValues = New(0x0016, 0x000D)

	// ColorFilterArrayPatternRows (0016,000E) VR=IS VM=1 Color Filter Array Pattern Rows
	ColorFilterArrayPatternRows = New(0x0016, 0x000E)

	// ColorFilterArrayPatternColumns (0016,000F) VR=IS VM=1 Color Filter Array Pattern Columns
	ColorFilterArrayPatternColumns = New(0x0016, 0x000F)

	// ColorFilterArrayPatternValues (0016,0010) VR=DS VM=1-n Color Filter Array Pattern Values
	ColorFilterArrayPatternValues = New(0x0016, 0x0010)

	// FlashFiringStatus (0016,0011) VR=US VM=1 Flash Firing Status
	FlashFiringStatus = New(0x0016, 0x0011)

	// FlashReturnStatus (0016,0012) VR=US VM=1 Flash Return Status
	FlashReturnStatus = New(0x0016, 0x0012)

	// FlashMode (0016,0013) VR=US VM=1 Flash Mode
	FlashMode = New(0x0016, 0x0013)

	// FlashFunctionPresent (0016,0014) VR=US VM=1 Flash Function Present
	FlashFunctionPresent = New(0x0016, 0x0014)

	// FlashRedEyeMode (0016,0015) VR=US VM=1 Flash Red Eye Mode
	FlashRedEyeMode = New(0x0016, 0x0015)

	// ExposureProgram (0016,0016) VR=US VM=1 Exposure Program
	ExposureProgram = New(0x0016, 0x0016)

	// SpectralSensitivity (0016,0017) VR=UT VM=1 Spectral Sensitivity
	SpectralSensitivity = New(0x0016, 0x0017)

	// PhotographicSensitivity (0016,0018) VR=IS VM=1 Photographic Sensitivity
	PhotographicSensitivity = New(0x0016, 0x0018)

	// SelfTimerMode (0016,0019) VR=IS VM=1 Self Timer Mode
	SelfTimerMode = New(0x0016, 0x0019)

	// SensitivityType (0016,001A) VR=US VM=1 Sensitivity Type
	SensitivityType = New(0x0016, 0x001A)

	// StandardOutputSensitivity (0016,001B) VR=IS VM=1 Standard Output Sensitivity
	StandardOutputSensitivity = New(0x0016, 0x001B)

	// RecommendedExposureIndex (0016,001C) VR=IS VM=1 Recommended Exposure Index
	RecommendedExposureIndex = New(0x0016, 0x001C)

	// ISOSpeed (0016,001D) VR=IS VM=1 ISO Speed
	ISOSpeed = New(0x0016, 0x001D)

	// ISOSpeedLatitudeyyy (0016,001E) VR=IS VM=1 ISO Speed Latitude yyy
	ISOSpeedLatitudeyyy = New(0x0016, 0x001E)

	// ISOSpeedLatitudezzz (0016,001F) VR=IS VM=1 ISO Speed Latitude zzz
	ISOSpeedLatitudezzz = New(0x0016, 0x001F)

	// EXIFVersion (0016,0020) VR=UT VM=1 EXIF Version
	EXIFVersion = New(0x0016, 0x0020)

	// ShutterSpeedValue (0016,0021) VR=DS VM=1 Shutter Speed Value
	ShutterSpeedValue = New(0x0016, 0x0021)

	// ApertureValue (0016,0022) VR=DS VM=1 Aperture Value
	ApertureValue = New(0x0016, 0x0022)

	// BrightnessValue (0016,0023) VR=DS VM=1 Brightness Value
	BrightnessValue = New(0x0016, 0x0023)

	// ExposureBiasValue (0016,0024) VR=DS VM=1 Exposure Bias Value
	ExposureBiasValue = New(0x0016, 0x0024)

	// MaxApertureValue (0016,0025) VR=DS VM=1 Max Aperture Value
	MaxApertureValue = New(0x0016, 0x0025)

	// SubjectDistance (0016,0026) VR=DS VM=1 Subject Distance
	SubjectDistance = New(0x0016, 0x0026)

	// MeteringMode (0016,0027) VR=US VM=1 Metering Mode
	MeteringMode = New(0x0016, 0x0027)

	// LightSource (0016,0028) VR=US VM=1 Light Source
	LightSource = New(0x0016, 0x0028)

	// FocalLength (0016,0029) VR=DS VM=1 Focal Length
	FocalLength = New(0x0016, 0x0029)

	// SubjectArea (0016,002A) VR=IS VM=2-4 Subject Area
	SubjectArea = New(0x0016, 0x002A)

	// MakerNote (0016,002B) VR=OB VM=1 Maker Note
	MakerNote = New(0x0016, 0x002B)

	// Temperature (0016,0030) VR=DS VM=1 Temperature
	Temperature = New(0x0016, 0x0030)

	// Humidity (0016,0031) VR=DS VM=1 Humidity
	Humidity = New(0x0016, 0x0031)

	// Pressure (0016,0032) VR=DS VM=1 Pressure
	Pressure = New(0x0016, 0x0032)

	// WaterDepth (0016,0033) VR=DS VM=1 Water Depth
	WaterDepth = New(0x0016, 0x0033)

	// Acceleration (0016,0034) VR=DS VM=1 Acceleration
	Acceleration = New(0x0016, 0x0034)

	// CameraElevationAngle (0016,0035) VR=DS VM=1 Camera Elevation Angle
	CameraElevationAngle = New(0x0016, 0x0035)

	// FlashEnergy (0016,0036) VR=DS VM=1-2 Flash Energy
	FlashEnergy = New(0x0016, 0x0036)

	// SubjectLocation (0016,0037) VR=IS VM=2 Subject Location
	SubjectLocation = New(0x0016, 0x0037)

	// PhotographicExposureIndex (0016,0038) VR=DS VM=1 Photographic Exposure Index
	PhotographicExposureIndex = New(0x0016, 0x0038)

	// SensingMethod (0016,0039) VR=US VM=1 Sensing Method
	SensingMethod = New(0x0016, 0x0039)

	// FileSource (0016,003A) VR=US VM=1 File Source
	FileSource = New(0x0016, 0x003A)

	// SceneType (0016,003B) VR=US VM=1 Scene Type
	SceneType = New(0x0016, 0x003B)

	// CustomRendered (0016,0041) VR=US VM=1 Custom Rendered
	CustomRendered = New(0x0016, 0x0041)

	// ExposureMode (0016,0042) VR=US VM=1 Exposure Mode
	ExposureMode = New(0x0016, 0x0042)

	// WhiteBalance (0016,0043) VR=US VM=1 White Balance
	WhiteBalance = New(0x0016, 0x0043)

	// DigitalZoomRatio (0016,0044) VR=DS VM=1 Digital Zoom Ratio
	DigitalZoomRatio = New(0x0016, 0x0044)

	// FocalLengthIn35mmFilm (0016,0045) VR=IS VM=1 Focal Length In 35mm Film
	FocalLengthIn35mmFilm = New(0x0016, 0x0045)

	// SceneCaptureType (0016,0046) VR=US VM=1 Scene Capture Type
	SceneCaptureType = New(0x0016, 0x0046)

	// GainControl (0016,0047) VR=US VM=1 Gain Control
	GainControl = New(0x0016, 0x0047)

	// Contrast (0016,0048) VR=US VM=1 Contrast
	Contrast = New(0x0016, 0x0048)

	// Saturation (0016,0049) VR=US VM=1 Saturation
	Saturation = New(0x0016, 0x0049)

	// Sharpness (0016,004A) VR=US VM=1 Sharpness
	Sharpness = New(0x0016, 0x004A)

	// DeviceSettingDescription (0016,004B) VR=OB VM=1 Device Setting Description
	DeviceSettingDescription = New(0x0016, 0x004B)

	// SubjectDistanceRange (0016,004C) VR=US VM=1 Subject Distance Range
	SubjectDistanceRange = New(0x0016, 0x004C)

	// CameraOwnerName (0016,004D) VR=UT VM=1 Camera Owner Name
	CameraOwnerName = New(0x0016, 0x004D)

	// LensSpecification (0016,004E) VR=DS VM=4 Lens Specification
	LensSpecification = New(0x0016, 0x004E)

	// LensMake (0016,004F) VR=UT VM=1 Lens Make
	LensMake = New(0x0016, 0x004F)

	// LensModel (0016,0050) VR=UT VM=1 Lens Model
	LensModel = New(0x0016, 0x0050)

	// LensSerialNumber (0016,0051) VR=UT VM=1 Lens Serial Number
	LensSerialNumber = New(0x0016, 0x0051)

	// InteroperabilityIndex (0016,0061) VR=CS VM=1 Interoperability Index
	InteroperabilityIndex = New(0x0016, 0x0061)

	// InteroperabilityVersion (0016,0062) VR=OB VM=1 Interoperability Version
	InteroperabilityVersion = New(0x0016, 0x0062)

	// GPSVersionID (0016,0070) VR=OB VM=1 GPS Version ID
	GPSVersionID = New(0x0016, 0x0070)

	// GPSLatitudeRef (0016,0071) VR=CS VM=1 GPS Latitude Ref
	GPSLatitudeRef = New(0x0016, 0x0071)

	// GPSLatitude (0016,0072) VR=DS VM=3 GPS Latitude
	GPSLatitude = New(0x0016, 0x0072)

	// GPSLongitudeRef (0016,0073) VR=CS VM=1 GPS Longitude Ref
	GPSLongitudeRef = New(0x0016, 0x0073)

	// GPSLongitude (0016,0074) VR=DS VM=3 GPS Longitude
	GPSLongitude = New(0x0016, 0x0074)

	// GPSAltitudeRef (0016,0075) VR=US VM=1 GPS Altitude Ref
	GPSAltitudeRef = New(0x0016, 0x0075)

	// GPSAltitude (0016,0076) VR=DS VM=1 GPS Altitude
	GPSAltitude = New(0x0016, 0x0076)

	// GPSTimeStamp (0016,0077) VR=DT VM=1 GPS Time Stamp
	GPSTimeStamp = New(0x0016, 0x0077)

	// GPSSatellites (0016,0078) VR=UT VM=1 GPS Satellites
	GPSSatellites = New(0x0016, 0x0078)

	// GPSStatus (0016,0079) VR=CS VM=1 GPS Status
	GPSStatus = New(0x0016, 0x0079)

	// GPSMeasureMode (0016,007A) VR=CS VM=1 GPS Measure Mode
	GPSMeasureMode = New(0x0016, 0x007A)

	// GPSDOP (0016,007B) VR=DS VM=1 GPS DOP
	GPSDOP = New(0x0016, 0x007B)

	// GPSSpeedRef (0016,007C) VR=CS VM=1 GPS Speed Ref
	GPSSpeedRef = New(0x0016, 0x007C)

	// GPSSpeed (0016,007D) VR=DS VM=1 GPS Speed
	GPSSpeed = New(0x0016, 0x007D)

	// GPSTrackRef (0016,007E) VR=CS VM=1 GPS Track Ref
	GPSTrackRef = New(0x0016, 0x007E)

	// GPSTrack (0016,007F) VR=DS VM=1 GPS Track
	GPSTrack = New(0x0016, 0x007F)

	// GPSImgDirectionRef (0016,0080) VR=CS VM=1 GPS Img Direction Ref
	GPSImgDirectionRef = New(0x0016, 0x0080)

	// GPSImgDirection (0016,0081) VR=DS VM=1 GPS Img Direction
	GPSImgDirection = New(0x0016, 0x0081)

	// GPSMapDatum (0016,0082) VR=UT VM=1 GPS Map Datum
	GPSMapDatum = New(0x0016, 0x0082)

	// GPSDestLatitudeRef (0016,0083) VR=CS VM=1 GPS Dest Latitude Ref
	GPSDestLatitudeRef = New(0x0016, 0x0083)

	// GPSDestLatitude (0016,0084) VR=DS VM=3 GPS Dest Latitude
	GPSDestLatitude = New(0x0016, 0x0084)

	// GPSDestLongitudeRef (0016,0085) VR=CS VM=1 GPS Dest Longitude Ref
	GPSDestLongitudeRef = New(0x0016, 0x0085)

	// GPSDestLongitude (0016,0086) VR=DS VM=3 GPS Dest Longitude
	GPSDestLongitude = New(0x0016, 0x0086)

	// GPSDestBearingRef (0016,0087) VR=CS VM=1 GPS Dest Bearing Ref
	GPSDestBearingRef = New(0x0016, 0x0087)

	// GPSDestBearing (0016,0088) VR=DS VM=1 GPS Dest Bearing
	GPSDestBearing = New(0x0016, 0x0088)

	// GPSDestDistanceRef (0016,0089) VR=CS VM=1 GPS Dest Distance Ref
	GPSDestDistanceRef = New(0x0016, 0x0089)

	// GPSDestDistance (0016,008A) VR=DS VM=1 GPS Dest Distance
	GPSDestDistance = New(0x0016, 0x008A)

	// GPSProcessingMethod (0016,008B) VR=OB VM=1 GPS Processing Method
	GPSProcessingMethod = New(0x0016, 0x008B)

	// GPSAreaInformation (0016,008C) VR=OB VM=1 GPS Area Information
	GPSAreaInformation = New(0x0016, 0x008C)

	// GPSDateStamp (0016,008D) VR=DT VM=1 GPS Date Stamp
	GPSDateStamp = New(0x0016, 0x008D)

	// GPSDifferential (0016,008E) VR=IS VM=1 GPS Differential
	GPSDifferential = New(0x0016, 0x008E)

	// LightSourcePolarization (0016,1001) VR=CS VM=1 Light Source Polarization
	LightSourcePolarization = New(0x0016, 0x1001)

	// EmitterColorTemperature (0016,1002) VR=DS VM=1 Emitter Color Temperature
	EmitterColorTemperature = New(0x0016, 0x1002)

	// ContactMethod (0016,1003) VR=CS VM=1 Contact Method
	ContactMethod = New(0x0016, 0x1003)

	// ImmersionMedia (0016,1004) VR=CS VM=1-n Immersion Media
	ImmersionMedia = New(0x0016, 0x1004)

	// OpticalMagnificationFactor (0016,1005) VR=DS VM=1 Optical Magnification Factor
	OpticalMagnificationFactor = New(0x0016, 0x1005)

	// ContrastBolusAgent (0018,0010) VR=LO VM=1 Contrast/Bolus Agent
	ContrastBolusAgent = New(0x0018, 0x0010)

	// ContrastBolusAgentSequence (0018,0012) VR=SQ VM=1 Contrast/Bolus Agent Sequence
	ContrastBolusAgentSequence = New(0x0018, 0x0012)

	// ContrastBolusT1Relaxivity (0018,0013) VR=FL VM=1 Contrast/Bolus T1 Relaxivity
	ContrastBolusT1Relaxivity = New(0x0018, 0x0013)

	// ContrastBolusAdministrationRouteSequence (0018,0014) VR=SQ VM=1 Contrast/Bolus Administration Route Sequence
	ContrastBolusAdministrationRouteSequence = New(0x0018, 0x0014)

	// BodyPartExamined (0018,0015) VR=CS VM=1 Body Part Examined
	BodyPartExamined = New(0x0018, 0x0015)

	// ScanningSequence (0018,0020) VR=CS VM=1-n Scanning Sequence
	ScanningSequence = New(0x0018, 0x0020)

	// SequenceVariant (0018,0021) VR=CS VM=1-n Sequence Variant
	SequenceVariant = New(0x0018, 0x0021)

	// ScanOptions (0018,0022) VR=CS VM=1-n Scan Options
	ScanOptions = New(0x0018, 0x0022)

	// MRAcquisitionType (0018,0023) VR=CS VM=1 MR Acquisition Type
	MRAcquisitionType = New(0x0018, 0x0023)

	// SequenceName (0018,0024) VR=SH VM=1 Sequence Name
	SequenceName = New(0x0018, 0x0024)

	// AngioFlag (0018,0025) VR=CS VM=1 Angio Flag
	AngioFlag = New(0x0018, 0x0025)

	// InterventionDrugInformationSequence (0018,0026) VR=SQ VM=1 Intervention Drug Information Sequence
	InterventionDrugInformationSequence = New(0x0018, 0x0026)

	// InterventionDrugStopTime (0018,0027) VR=TM VM=1 Intervention Drug Stop Time
	InterventionDrugStopTime = New(0x0018, 0x0027)

	// InterventionDrugDose (0018,0028) VR=DS VM=1 Intervention Drug Dose
	InterventionDrugDose = New(0x0018, 0x0028)

	// InterventionDrugCodeSequence (0018,0029) VR=SQ VM=1 Intervention Drug Code Sequence
	InterventionDrugCodeSequence = New(0x0018, 0x0029)

	// AdditionalDrugSequence (0018,002A) VR=SQ VM=1 Additional Drug Sequence
	AdditionalDrugSequence = New(0x0018, 0x002A)

	// RadionuclideRETIRED (0018,0030) VR=LO VM=1-n Radionuclide (RETIRED)
	RadionuclideRETIRED = New(0x0018, 0x0030)

	// Radiopharmaceutical (0018,0031) VR=LO VM=1 Radiopharmaceutical
	Radiopharmaceutical = New(0x0018, 0x0031)

	// EnergyWindowCenterlineRETIRED (0018,0032) VR=DS VM=1 Energy Window Centerline (RETIRED)
	EnergyWindowCenterlineRETIRED = New(0x0018, 0x0032)

	// EnergyWindowTotalWidthRETIRED (0018,0033) VR=DS VM=1-n Energy Window Total Width (RETIRED)
	EnergyWindowTotalWidthRETIRED = New(0x0018, 0x0033)

	// InterventionDrugName (0018,0034) VR=LO VM=1 Intervention Drug Name
	InterventionDrugName = New(0x0018, 0x0034)

	// InterventionDrugStartTime (0018,0035) VR=TM VM=1 Intervention Drug Start Time
	InterventionDrugStartTime = New(0x0018, 0x0035)

	// InterventionSequence (0018,0036) VR=SQ VM=1 Intervention Sequence
	InterventionSequence = New(0x0018, 0x0036)

	// TherapyTypeRETIRED (0018,0037) VR=CS VM=1 Therapy Type (RETIRED)
	TherapyTypeRETIRED = New(0x0018, 0x0037)

	// InterventionStatus (0018,0038) VR=CS VM=1 Intervention Status
	InterventionStatus = New(0x0018, 0x0038)

	// TherapyDescriptionRETIRED (0018,0039) VR=CS VM=1 Therapy Description (RETIRED)
	TherapyDescriptionRETIRED = New(0x0018, 0x0039)

	// InterventionDescription (0018,003A) VR=ST VM=1 Intervention Description
	InterventionDescription = New(0x0018, 0x003A)

	// CineRate (0018,0040) VR=IS VM=1 Cine Rate
	CineRate = New(0x0018, 0x0040)

	// InitialCineRunState (0018,0042) VR=CS VM=1 Initial Cine Run State
	InitialCineRunState = New(0x0018, 0x0042)

	// SliceThickness (0018,0050) VR=DS VM=1 Slice Thickness
	SliceThickness = New(0x0018, 0x0050)

	// KVP (0018,0060) VR=DS VM=1 KVP
	KVP = New(0x0018, 0x0060)

	// CountsAccumulated (0018,0070) VR=IS VM=1 Counts Accumulated
	CountsAccumulated = New(0x0018, 0x0070)

	// AcquisitionTerminationCondition (0018,0071) VR=CS VM=1 Acquisition Termination Condition
	AcquisitionTerminationCondition = New(0x0018, 0x0071)

	// EffectiveDuration (0018,0072) VR=DS VM=1 Effective Duration
	EffectiveDuration = New(0x0018, 0x0072)

	// AcquisitionStartCondition (0018,0073) VR=CS VM=1 Acquisition Start Condition
	AcquisitionStartCondition = New(0x0018, 0x0073)

	// AcquisitionStartConditionData (0018,0074) VR=IS VM=1 Acquisition Start Condition Data
	AcquisitionStartConditionData = New(0x0018, 0x0074)

	// AcquisitionTerminationConditionData (0018,0075) VR=IS VM=1 Acquisition Termination Condition Data
	AcquisitionTerminationConditionData = New(0x0018, 0x0075)

	// RepetitionTime (0018,0080) VR=DS VM=1 Repetition Time
	RepetitionTime = New(0x0018, 0x0080)

	// EchoTime (0018,0081) VR=DS VM=1 Echo Time
	EchoTime = New(0x0018, 0x0081)

	// InversionTime (0018,0082) VR=DS VM=1 Inversion Time
	InversionTime = New(0x0018, 0x0082)

	// NumberOfAverages (0018,0083) VR=DS VM=1 Number of Averages
	NumberOfAverages = New(0x0018, 0x0083)

	// ImagingFrequency (0018,0084) VR=DS VM=1 Imaging Frequency
	ImagingFrequency = New(0x0018, 0x0084)

	// ImagedNucleus (0018,0085) VR=SH VM=1 Imaged Nucleus
	ImagedNucleus = New(0x0018, 0x0085)

	// EchoNumbers (0018,0086) VR=IS VM=1-n Echo Number(s)
	EchoNumbers = New(0x0018, 0x0086)

	// MagneticFieldStrength (0018,0087) VR=DS VM=1 Magnetic Field Strength
	MagneticFieldStrength = New(0x0018, 0x0087)

	// SpacingBetweenSlices (0018,0088) VR=DS VM=1 Spacing Between Slices
	SpacingBetweenSlices = New(0x0018, 0x0088)

	// NumberOfPhaseEncodingSteps (0018,0089) VR=IS VM=1 Number of Phase Encoding Steps
	NumberOfPhaseEncodingSteps = New(0x0018, 0x0089)

	// DataCollectionDiameter (0018,0090) VR=DS VM=1 Data Collection Diameter
	DataCollectionDiameter = New(0x0018, 0x0090)

	// EchoTrainLength (0018,0091) VR=IS VM=1 Echo Train Length
	EchoTrainLength = New(0x0018, 0x0091)

	// PercentSampling (0018,0093) VR=DS VM=1 Percent Sampling
	PercentSampling = New(0x0018, 0x0093)

	// PercentPhaseFieldOfView (0018,0094) VR=DS VM=1 Percent Phase Field of View
	PercentPhaseFieldOfView = New(0x0018, 0x0094)

	// PixelBandwidth (0018,0095) VR=DS VM=1 Pixel Bandwidth
	PixelBandwidth = New(0x0018, 0x0095)

	// DeviceSerialNumber (0018,1000) VR=LO VM=1 Device Serial Number
	DeviceSerialNumber = New(0x0018, 0x1000)

	// DeviceUID (0018,1002) VR=UI VM=1 Device UID
	DeviceUID = New(0x0018, 0x1002)

	// DeviceID (0018,1003) VR=LO VM=1 Device ID
	DeviceID = New(0x0018, 0x1003)

	// PlateID (0018,1004) VR=LO VM=1 Plate ID
	PlateID = New(0x0018, 0x1004)

	// GeneratorID (0018,1005) VR=LO VM=1 Generator ID
	GeneratorID = New(0x0018, 0x1005)

	// GridID (0018,1006) VR=LO VM=1 Grid ID
	GridID = New(0x0018, 0x1006)

	// CassetteID (0018,1007) VR=LO VM=1 Cassette ID
	CassetteID = New(0x0018, 0x1007)

	// GantryID (0018,1008) VR=LO VM=1 Gantry ID
	GantryID = New(0x0018, 0x1008)

	// UniqueDeviceIdentifier (0018,1009) VR=UT VM=1 Unique Device Identifier
	UniqueDeviceIdentifier = New(0x0018, 0x1009)

	// UDISequence (0018,100A) VR=SQ VM=1 UDI Sequence
	UDISequence = New(0x0018, 0x100A)

	// ManufacturerDeviceClassUID (0018,100B) VR=UI VM=1-n Manufacturer's Device Class UID
	ManufacturerDeviceClassUID = New(0x0018, 0x100B)

	// SecondaryCaptureDeviceID (0018,1010) VR=LO VM=1 Secondary Capture Device ID
	SecondaryCaptureDeviceID = New(0x0018, 0x1010)

	// HardcopyCreationDeviceIDRETIRED (0018,1011) VR=LO VM=1 Hardcopy Creation Device ID (RETIRED)
	HardcopyCreationDeviceIDRETIRED = New(0x0018, 0x1011)

	// DateOfSecondaryCapture (0018,1012) VR=DA VM=1 Date of Secondary Capture
	DateOfSecondaryCapture = New(0x0018, 0x1012)

	// TimeOfSecondaryCapture (0018,1014) VR=TM VM=1 Time of Secondary Capture
	TimeOfSecondaryCapture = New(0x0018, 0x1014)

	// SecondaryCaptureDeviceManufacturer (0018,1016) VR=LO VM=1 Secondary Capture Device Manufacturer
	SecondaryCaptureDeviceManufacturer = New(0x0018, 0x1016)

	// HardcopyDeviceManufacturerRETIRED (0018,1017) VR=LO VM=1 Hardcopy Device Manufacturer (RETIRED)
	HardcopyDeviceManufacturerRETIRED = New(0x0018, 0x1017)

	// SecondaryCaptureDeviceManufacturerModelName (0018,1018) VR=LO VM=1 Secondary Capture Device Manufacturer's Model Name
	SecondaryCaptureDeviceManufacturerModelName = New(0x0018, 0x1018)

	// SecondaryCaptureDeviceSoftwareVersions (0018,1019) VR=LO VM=1-n Secondary Capture Device Software Versions
	SecondaryCaptureDeviceSoftwareVersions = New(0x0018, 0x1019)

	// HardcopyDeviceSoftwareVersionRETIRED (0018,101A) VR=LO VM=1-n Hardcopy Device Software Version (RETIRED)
	HardcopyDeviceSoftwareVersionRETIRED = New(0x0018, 0x101A)

	// HardcopyDeviceManufacturerModelNameRETIRED (0018,101B) VR=LO VM=1 Hardcopy Device Manufacturer's Model Name (RETIRED)
	HardcopyDeviceManufacturerModelNameRETIRED = New(0x0018, 0x101B)

	// SoftwareVersions (0018,1020) VR=LO VM=1-n Software Versions
	SoftwareVersions = New(0x0018, 0x1020)

	// VideoImageFormatAcquired (0018,1022) VR=SH VM=1 Video Image Format Acquired
	VideoImageFormatAcquired = New(0x0018, 0x1022)

	// DigitalImageFormatAcquired (0018,1023) VR=LO VM=1 Digital Image Format Acquired
	DigitalImageFormatAcquired = New(0x0018, 0x1023)

	// ProtocolName (0018,1030) VR=LO VM=1 Protocol Name
	ProtocolName = New(0x0018, 0x1030)

	// ContrastBolusRoute (0018,1040) VR=LO VM=1 Contrast/Bolus Route
	ContrastBolusRoute = New(0x0018, 0x1040)

	// ContrastBolusVolume (0018,1041) VR=DS VM=1 Contrast/Bolus Volume
	ContrastBolusVolume = New(0x0018, 0x1041)

	// ContrastBolusStartTime (0018,1042) VR=TM VM=1 Contrast/Bolus Start Time
	ContrastBolusStartTime = New(0x0018, 0x1042)

	// ContrastBolusStopTime (0018,1043) VR=TM VM=1 Contrast/Bolus Stop Time
	ContrastBolusStopTime = New(0x0018, 0x1043)

	// ContrastBolusTotalDose (0018,1044) VR=DS VM=1 Contrast/Bolus Total Dose
	ContrastBolusTotalDose = New(0x0018, 0x1044)

	// SyringeCounts (0018,1045) VR=IS VM=1 Syringe Counts
	SyringeCounts = New(0x0018, 0x1045)

	// ContrastFlowRate (0018,1046) VR=DS VM=1-n Contrast Flow Rate
	ContrastFlowRate = New(0x0018, 0x1046)

	// ContrastFlowDuration (0018,1047) VR=DS VM=1-n Contrast Flow Duration
	ContrastFlowDuration = New(0x0018, 0x1047)

	// ContrastBolusIngredient (0018,1048) VR=CS VM=1 Contrast/Bolus Ingredient
	ContrastBolusIngredient = New(0x0018, 0x1048)

	// ContrastBolusIngredientConcentration (0018,1049) VR=DS VM=1 Contrast/Bolus Ingredient Concentration
	ContrastBolusIngredientConcentration = New(0x0018, 0x1049)

	// SpatialResolution (0018,1050) VR=DS VM=1 Spatial Resolution
	SpatialResolution = New(0x0018, 0x1050)

	// TriggerTime (0018,1060) VR=DS VM=1 Trigger Time
	TriggerTime = New(0x0018, 0x1060)

	// TriggerSourceOrType (0018,1061) VR=LO VM=1 Trigger Source or Type
	TriggerSourceOrType = New(0x0018, 0x1061)

	// NominalInterval (0018,1062) VR=IS VM=1 Nominal Interval
	NominalInterval = New(0x0018, 0x1062)

	// FrameTime (0018,1063) VR=DS VM=1 Frame Time
	FrameTime = New(0x0018, 0x1063)

	// CardiacFramingType (0018,1064) VR=LO VM=1 Cardiac Framing Type
	CardiacFramingType = New(0x0018, 0x1064)

	// FrameTimeVector (0018,1065) VR=DS VM=1-n Frame Time Vector
	FrameTimeVector = New(0x0018, 0x1065)

	// FrameDelay (0018,1066) VR=DS VM=1 Frame Delay
	FrameDelay = New(0x0018, 0x1066)

	// ImageTriggerDelay (0018,1067) VR=DS VM=1 Image Trigger Delay
	ImageTriggerDelay = New(0x0018, 0x1067)

	// MultiplexGroupTimeOffset (0018,1068) VR=DS VM=1 Multiplex Group Time Offset
	MultiplexGroupTimeOffset = New(0x0018, 0x1068)

	// TriggerTimeOffset (0018,1069) VR=DS VM=1 Trigger Time Offset
	TriggerTimeOffset = New(0x0018, 0x1069)

	// SynchronizationTrigger (0018,106A) VR=CS VM=1 Synchronization Trigger
	SynchronizationTrigger = New(0x0018, 0x106A)

	// SynchronizationChannel (0018,106C) VR=US VM=2 Synchronization Channel
	SynchronizationChannel = New(0x0018, 0x106C)

	// TriggerSamplePosition (0018,106E) VR=UL VM=1 Trigger Sample Position
	TriggerSamplePosition = New(0x0018, 0x106E)

	// RadiopharmaceuticalRoute (0018,1070) VR=LO VM=1 Radiopharmaceutical Route
	RadiopharmaceuticalRoute = New(0x0018, 0x1070)

	// RadiopharmaceuticalVolume (0018,1071) VR=DS VM=1 Radiopharmaceutical Volume
	RadiopharmaceuticalVolume = New(0x0018, 0x1071)

	// RadiopharmaceuticalStartTime (0018,1072) VR=TM VM=1 Radiopharmaceutical Start Time
	RadiopharmaceuticalStartTime = New(0x0018, 0x1072)

	// RadiopharmaceuticalStopTime (0018,1073) VR=TM VM=1 Radiopharmaceutical Stop Time
	RadiopharmaceuticalStopTime = New(0x0018, 0x1073)

	// RadionuclideTotalDose (0018,1074) VR=DS VM=1 Radionuclide Total Dose
	RadionuclideTotalDose = New(0x0018, 0x1074)

	// RadionuclideHalfLife (0018,1075) VR=DS VM=1 Radionuclide Half Life
	RadionuclideHalfLife = New(0x0018, 0x1075)

	// RadionuclidePositronFraction (0018,1076) VR=DS VM=1 Radionuclide Positron Fraction
	RadionuclidePositronFraction = New(0x0018, 0x1076)

	// RadiopharmaceuticalSpecificActivity (0018,1077) VR=DS VM=1 Radiopharmaceutical Specific Activity
	RadiopharmaceuticalSpecificActivity = New(0x0018, 0x1077)

	// RadiopharmaceuticalStartDateTime (0018,1078) VR=DT VM=1 Radiopharmaceutical Start DateTime
	RadiopharmaceuticalStartDateTime = New(0x0018, 0x1078)

	// RadiopharmaceuticalStopDateTime (0018,1079) VR=DT VM=1 Radiopharmaceutical Stop DateTime
	RadiopharmaceuticalStopDateTime = New(0x0018, 0x1079)

	// BeatRejectionFlag (0018,1080) VR=CS VM=1 Beat Rejection Flag
	BeatRejectionFlag = New(0x0018, 0x1080)

	// LowRRValue (0018,1081) VR=IS VM=1 Low R-R Value
	LowRRValue = New(0x0018, 0x1081)

	// HighRRValue (0018,1082) VR=IS VM=1 High R-R Value
	HighRRValue = New(0x0018, 0x1082)

	// IntervalsAcquired (0018,1083) VR=IS VM=1 Intervals Acquired
	IntervalsAcquired = New(0x0018, 0x1083)

	// IntervalsRejected (0018,1084) VR=IS VM=1 Intervals Rejected
	IntervalsRejected = New(0x0018, 0x1084)

	// PVCRejection (0018,1085) VR=LO VM=1 PVC Rejection
	PVCRejection = New(0x0018, 0x1085)

	// SkipBeats (0018,1086) VR=IS VM=1 Skip Beats
	SkipBeats = New(0x0018, 0x1086)

	// HeartRate (0018,1088) VR=IS VM=1 Heart Rate
	HeartRate = New(0x0018, 0x1088)

	// CardiacNumberOfImages (0018,1090) VR=IS VM=1 Cardiac Number of Images
	CardiacNumberOfImages = New(0x0018, 0x1090)

	// TriggerWindow (0018,1094) VR=IS VM=1 Trigger Window
	TriggerWindow = New(0x0018, 0x1094)

	// ReconstructionDiameter (0018,1100) VR=DS VM=1 Reconstruction Diameter
	ReconstructionDiameter = New(0x0018, 0x1100)

	// DistanceSourceToDetector (0018,1110) VR=DS VM=1 Distance Source to Detector
	DistanceSourceToDetector = New(0x0018, 0x1110)

	// DistanceSourceToPatient (0018,1111) VR=DS VM=1 Distance Source to Patient
	DistanceSourceToPatient = New(0x0018, 0x1111)

	// EstimatedRadiographicMagnificationFactor (0018,1114) VR=DS VM=1 Estimated Radiographic Magnification Factor
	EstimatedRadiographicMagnificationFactor = New(0x0018, 0x1114)

	// GantryDetectorTilt (0018,1120) VR=DS VM=1 Gantry/Detector Tilt
	GantryDetectorTilt = New(0x0018, 0x1120)

	// GantryDetectorSlew (0018,1121) VR=DS VM=1 Gantry/Detector Slew
	GantryDetectorSlew = New(0x0018, 0x1121)

	// TableHeight (0018,1130) VR=DS VM=1 Table Height
	TableHeight = New(0x0018, 0x1130)

	// TableTraverse (0018,1131) VR=DS VM=1 Table Traverse
	TableTraverse = New(0x0018, 0x1131)

	// TableMotion (0018,1134) VR=CS VM=1 Table Motion
	TableMotion = New(0x0018, 0x1134)

	// TableVerticalIncrement (0018,1135) VR=DS VM=1-n Table Vertical Increment
	TableVerticalIncrement = New(0x0018, 0x1135)

	// TableLateralIncrement (0018,1136) VR=DS VM=1-n Table Lateral Increment
	TableLateralIncrement = New(0x0018, 0x1136)

	// TableLongitudinalIncrement (0018,1137) VR=DS VM=1-n Table Longitudinal Increment
	TableLongitudinalIncrement = New(0x0018, 0x1137)

	// TableAngle (0018,1138) VR=DS VM=1 Table Angle
	TableAngle = New(0x0018, 0x1138)

	// TableType (0018,113A) VR=CS VM=1 Table Type
	TableType = New(0x0018, 0x113A)

	// RotationDirection (0018,1140) VR=CS VM=1 Rotation Direction
	RotationDirection = New(0x0018, 0x1140)

	// AngularPositionRETIRED (0018,1141) VR=DS VM=1 Angular Position (RETIRED)
	AngularPositionRETIRED = New(0x0018, 0x1141)

	// RadialPosition (0018,1142) VR=DS VM=1-n Radial Position
	RadialPosition = New(0x0018, 0x1142)

	// ScanArc (0018,1143) VR=DS VM=1 Scan Arc
	ScanArc = New(0x0018, 0x1143)

	// AngularStep (0018,1144) VR=DS VM=1 Angular Step
	AngularStep = New(0x0018, 0x1144)

	// CenterOfRotationOffset (0018,1145) VR=DS VM=1 Center of Rotation Offset
	CenterOfRotationOffset = New(0x0018, 0x1145)

	// RotationOffsetRETIRED (0018,1146) VR=DS VM=1-n Rotation Offset (RETIRED)
	RotationOffsetRETIRED = New(0x0018, 0x1146)

	// FieldOfViewShape (0018,1147) VR=CS VM=1 Field of View Shape
	FieldOfViewShape = New(0x0018, 0x1147)

	// FieldOfViewDimensions (0018,1149) VR=IS VM=1-2 Field of View Dimension(s)
	FieldOfViewDimensions = New(0x0018, 0x1149)

	// ExposureTime (0018,1150) VR=IS VM=1 Exposure Time
	ExposureTime = New(0x0018, 0x1150)

	// XRayTubeCurrent (0018,1151) VR=IS VM=1 X-Ray Tube Current
	XRayTubeCurrent = New(0x0018, 0x1151)

	// Exposure (0018,1152) VR=IS VM=1 Exposure
	Exposure = New(0x0018, 0x1152)

	// ExposureInuAs (0018,1153) VR=IS VM=1 Exposure in µAs
	ExposureInuAs = New(0x0018, 0x1153)

	// AveragePulseWidth (0018,1154) VR=DS VM=1 Average Pulse Width
	AveragePulseWidth = New(0x0018, 0x1154)

	// RadiationSetting (0018,1155) VR=CS VM=1 Radiation Setting
	RadiationSetting = New(0x0018, 0x1155)

	// RectificationType (0018,1156) VR=CS VM=1 Rectification Type
	RectificationType = New(0x0018, 0x1156)

	// RadiationMode (0018,115A) VR=CS VM=1 Radiation Mode
	RadiationMode = New(0x0018, 0x115A)

	// ImageAndFluoroscopyAreaDoseProduct (0018,115E) VR=DS VM=1 Image and Fluoroscopy Area Dose Product
	ImageAndFluoroscopyAreaDoseProduct = New(0x0018, 0x115E)

	// FilterType (0018,1160) VR=SH VM=1 Filter Type
	FilterType = New(0x0018, 0x1160)

	// TypeOfFilters (0018,1161) VR=LO VM=1-n Type of Filters
	TypeOfFilters = New(0x0018, 0x1161)

	// IntensifierSize (0018,1162) VR=DS VM=1 Intensifier Size
	IntensifierSize = New(0x0018, 0x1162)

	// ImagerPixelSpacing (0018,1164) VR=DS VM=2 Imager Pixel Spacing
	ImagerPixelSpacing = New(0x0018, 0x1164)

	// Grid (0018,1166) VR=CS VM=1-n Grid
	Grid = New(0x0018, 0x1166)

	// GeneratorPower (0018,1170) VR=IS VM=1 Generator Power
	GeneratorPower = New(0x0018, 0x1170)

	// CollimatorGridName (0018,1180) VR=SH VM=1 Collimator/grid Name
	CollimatorGridName = New(0x0018, 0x1180)

	// CollimatorType (0018,1181) VR=CS VM=1 Collimator Type
	CollimatorType = New(0x0018, 0x1181)

	// FocalDistance (0018,1182) VR=IS VM=1-2 Focal Distance
	FocalDistance = New(0x0018, 0x1182)

	// XFocusCenter (0018,1183) VR=DS VM=1-2 X Focus Center
	XFocusCenter = New(0x0018, 0x1183)

	// YFocusCenter (0018,1184) VR=DS VM=1-2 Y Focus Center
	YFocusCenter = New(0x0018, 0x1184)

	// FocalSpots (0018,1190) VR=DS VM=1-n Focal Spot(s)
	FocalSpots = New(0x0018, 0x1190)

	// AnodeTargetMaterial (0018,1191) VR=CS VM=1 Anode Target Material
	AnodeTargetMaterial = New(0x0018, 0x1191)

	// BodyPartThickness (0018,11A0) VR=DS VM=1 Body Part Thickness
	BodyPartThickness = New(0x0018, 0x11A0)

	// CompressionForce (0018,11A2) VR=DS VM=1 Compression Force
	CompressionForce = New(0x0018, 0x11A2)

	// CompressionPressure (0018,11A3) VR=DS VM=1 Compression Pressure
	CompressionPressure = New(0x0018, 0x11A3)

	// PaddleDescription (0018,11A4) VR=LO VM=1 Paddle Description
	PaddleDescription = New(0x0018, 0x11A4)

	// CompressionContactArea (0018,11A5) VR=DS VM=1 Compression Contact Area
	CompressionContactArea = New(0x0018, 0x11A5)

	// AcquisitionMode (0018,11B0) VR=LO VM=1 Acquisition Mode
	AcquisitionMode = New(0x0018, 0x11B0)

	// DoseModeName (0018,11B1) VR=LO VM=1 Dose Mode Name
	DoseModeName = New(0x0018, 0x11B1)

	// AcquiredSubtractionMaskFlag (0018,11B2) VR=CS VM=1 Acquired Subtraction Mask Flag
	AcquiredSubtractionMaskFlag = New(0x0018, 0x11B2)

	// FluoroscopyPersistenceFlag (0018,11B3) VR=CS VM=1 Fluoroscopy Persistence Flag
	FluoroscopyPersistenceFlag = New(0x0018, 0x11B3)

	// FluoroscopyLastImageHoldPersistenceFlag (0018,11B4) VR=CS VM=1 Fluoroscopy Last Image Hold Persistence Flag
	FluoroscopyLastImageHoldPersistenceFlag = New(0x0018, 0x11B4)

	// UpperLimitNumberOfPersistentFluoroscopyFrames (0018,11B5) VR=IS VM=1 Upper Limit Number Of Persistent Fluoroscopy Frames
	UpperLimitNumberOfPersistentFluoroscopyFrames = New(0x0018, 0x11B5)

	// ContrastBolusAutoInjectionTriggerFlag (0018,11B6) VR=CS VM=1 Contrast/Bolus Auto Injection Trigger Flag
	ContrastBolusAutoInjectionTriggerFlag = New(0x0018, 0x11B6)

	// ContrastBolusInjectionDelay (0018,11B7) VR=FD VM=1 Contrast/Bolus Injection Delay
	ContrastBolusInjectionDelay = New(0x0018, 0x11B7)

	// XAAcquisitionPhaseDetailsSequence (0018,11B8) VR=SQ VM=1 XA Acquisition Phase Details Sequence
	XAAcquisitionPhaseDetailsSequence = New(0x0018, 0x11B8)

	// XAAcquisitionFrameRate (0018,11B9) VR=FD VM=1 XA Acquisition Frame Rate
	XAAcquisitionFrameRate = New(0x0018, 0x11B9)

	// XAPlaneDetailsSequence (0018,11BA) VR=SQ VM=1 XA Plane Details Sequence
	XAPlaneDetailsSequence = New(0x0018, 0x11BA)

	// AcquisitionFieldOfViewLabel (0018,11BB) VR=LO VM=1 Acquisition Field of View Label
	AcquisitionFieldOfViewLabel = New(0x0018, 0x11BB)

	// XRayFilterDetailsSequence (0018,11BC) VR=SQ VM=1 X-Ray Filter Details Sequence
	XRayFilterDetailsSequence = New(0x0018, 0x11BC)

	// XAAcquisitionDuration (0018,11BD) VR=FD VM=1 XA Acquisition Duration
	XAAcquisitionDuration = New(0x0018, 0x11BD)

	// ReconstructionPipelineType (0018,11BE) VR=CS VM=1 Reconstruction Pipeline Type
	ReconstructionPipelineType = New(0x0018, 0x11BE)

	// ImageFilterDetailsSequence (0018,11BF) VR=SQ VM=1 Image Filter Details Sequence
	ImageFilterDetailsSequence = New(0x0018, 0x11BF)

	// AppliedMaskSubtractionFlag (0018,11C0) VR=CS VM=1 Applied Mask Subtraction Flag
	AppliedMaskSubtractionFlag = New(0x0018, 0x11C0)

	// RequestedSeriesDescriptionCodeSequence (0018,11C1) VR=SQ VM=1 Requested Series Description Code Sequence
	RequestedSeriesDescriptionCodeSequence = New(0x0018, 0x11C1)

	// DateOfLastCalibration (0018,1200) VR=DA VM=1-n Date of Last Calibration
	DateOfLastCalibration = New(0x0018, 0x1200)

	// TimeOfLastCalibration (0018,1201) VR=TM VM=1-n Time of Last Calibration
	TimeOfLastCalibration = New(0x0018, 0x1201)

	// DateTimeOfLastCalibration (0018,1202) VR=DT VM=1 DateTime of Last Calibration
	DateTimeOfLastCalibration = New(0x0018, 0x1202)

	// CalibrationDateTime (0018,1203) VR=DT VM=1 Calibration DateTime
	CalibrationDateTime = New(0x0018, 0x1203)

	// DateOfManufacture (0018,1204) VR=DA VM=1 Date of Manufacture
	DateOfManufacture = New(0x0018, 0x1204)

	// DateOfInstallation (0018,1205) VR=DA VM=1 Date of Installation
	DateOfInstallation = New(0x0018, 0x1205)

	// ConvolutionKernel (0018,1210) VR=SH VM=1-n Convolution Kernel
	ConvolutionKernel = New(0x0018, 0x1210)

	// UpperLowerPixelValuesRETIRED (0018,1240) VR=IS VM=1-n Upper/Lower Pixel Values (RETIRED)
	UpperLowerPixelValuesRETIRED = New(0x0018, 0x1240)

	// ActualFrameDuration (0018,1242) VR=IS VM=1 Actual Frame Duration
	ActualFrameDuration = New(0x0018, 0x1242)

	// CountRate (0018,1243) VR=IS VM=1 Count Rate
	CountRate = New(0x0018, 0x1243)

	// PreferredPlaybackSequencing (0018,1244) VR=US VM=1 Preferred Playback Sequencing
	PreferredPlaybackSequencing = New(0x0018, 0x1244)

	// ReceiveCoilName (0018,1250) VR=SH VM=1 Receive Coil Name
	ReceiveCoilName = New(0x0018, 0x1250)

	// TransmitCoilName (0018,1251) VR=SH VM=1 Transmit Coil Name
	TransmitCoilName = New(0x0018, 0x1251)

	// PlateType (0018,1260) VR=SH VM=1 Plate Type
	PlateType = New(0x0018, 0x1260)

	// PhosphorType (0018,1261) VR=LO VM=1 Phosphor Type
	PhosphorType = New(0x0018, 0x1261)

	// WaterEquivalentDiameter (0018,1271) VR=FD VM=1 Water Equivalent Diameter
	WaterEquivalentDiameter = New(0x0018, 0x1271)

	// WaterEquivalentDiameterCalculationMethodCodeSequence (0018,1272) VR=SQ VM=1 Water Equivalent Diameter Calculation Method Code Sequence
	WaterEquivalentDiameterCalculationMethodCodeSequence = New(0x0018, 0x1272)

	// ScanVelocity (0018,1300) VR=DS VM=1 Scan Velocity
	ScanVelocity = New(0x0018, 0x1300)

	// WholeBodyTechnique (0018,1301) VR=CS VM=1-n Whole Body Technique
	WholeBodyTechnique = New(0x0018, 0x1301)

	// ScanLength (0018,1302) VR=IS VM=1 Scan Length
	ScanLength = New(0x0018, 0x1302)

	// AcquisitionMatrix (0018,1310) VR=US VM=4 Acquisition Matrix
	AcquisitionMatrix = New(0x0018, 0x1310)

	// InPlanePhaseEncodingDirection (0018,1312) VR=CS VM=1 In-plane Phase Encoding Direction
	InPlanePhaseEncodingDirection = New(0x0018, 0x1312)

	// FlipAngle (0018,1314) VR=DS VM=1 Flip Angle
	FlipAngle = New(0x0018, 0x1314)

	// VariableFlipAngleFlag (0018,1315) VR=CS VM=1 Variable Flip Angle Flag
	VariableFlipAngleFlag = New(0x0018, 0x1315)

	// SAR (0018,1316) VR=DS VM=1 SAR
	SAR = New(0x0018, 0x1316)

	// dBdt (0018,1318) VR=DS VM=1 dB/dt
	dBdt = New(0x0018, 0x1318)

	// B1rms (0018,1320) VR=FL VM=1 B1rms
	B1rms = New(0x0018, 0x1320)

	// AcquisitionDeviceProcessingDescription (0018,1400) VR=LO VM=1 Acquisition Device Processing Description
	AcquisitionDeviceProcessingDescription = New(0x0018, 0x1400)

	// AcquisitionDeviceProcessingCode (0018,1401) VR=LO VM=1 Acquisition Device Processing Code
	AcquisitionDeviceProcessingCode = New(0x0018, 0x1401)

	// CassetteOrientation (0018,1402) VR=CS VM=1 Cassette Orientation
	CassetteOrientation = New(0x0018, 0x1402)

	// CassetteSize (0018,1403) VR=CS VM=1 Cassette Size
	CassetteSize = New(0x0018, 0x1403)

	// ExposuresOnPlate (0018,1404) VR=US VM=1 Exposures on Plate
	ExposuresOnPlate = New(0x0018, 0x1404)

	// RelativeXRayExposure (0018,1405) VR=IS VM=1 Relative X-Ray Exposure
	RelativeXRayExposure = New(0x0018, 0x1405)

	// ExposureIndex (0018,1411) VR=DS VM=1 Exposure Index
	ExposureIndex = New(0x0018, 0x1411)

	// TargetExposureIndex (0018,1412) VR=DS VM=1 Target Exposure Index
	TargetExposureIndex = New(0x0018, 0x1412)

	// DeviationIndex (0018,1413) VR=DS VM=1 Deviation Index
	DeviationIndex = New(0x0018, 0x1413)

	// ColumnAngulation (0018,1450) VR=DS VM=1 Column Angulation
	ColumnAngulation = New(0x0018, 0x1450)

	// TomoLayerHeight (0018,1460) VR=DS VM=1 Tomo Layer Height
	TomoLayerHeight = New(0x0018, 0x1460)

	// TomoAngle (0018,1470) VR=DS VM=1 Tomo Angle
	TomoAngle = New(0x0018, 0x1470)

	// TomoTime (0018,1480) VR=DS VM=1 Tomo Time
	TomoTime = New(0x0018, 0x1480)

	// TomoType (0018,1490) VR=CS VM=1 Tomo Type
	TomoType = New(0x0018, 0x1490)

	// TomoClass (0018,1491) VR=CS VM=1 Tomo Class
	TomoClass = New(0x0018, 0x1491)

	// NumberOfTomosynthesisSourceImages (0018,1495) VR=IS VM=1 Number of Tomosynthesis Source Images
	NumberOfTomosynthesisSourceImages = New(0x0018, 0x1495)

	// PositionerMotion (0018,1500) VR=CS VM=1 Positioner Motion
	PositionerMotion = New(0x0018, 0x1500)

	// PositionerType (0018,1508) VR=CS VM=1 Positioner Type
	PositionerType = New(0x0018, 0x1508)

	// PositionerPrimaryAngle (0018,1510) VR=DS VM=1 Positioner Primary Angle
	PositionerPrimaryAngle = New(0x0018, 0x1510)

	// PositionerSecondaryAngle (0018,1511) VR=DS VM=1 Positioner Secondary Angle
	PositionerSecondaryAngle = New(0x0018, 0x1511)

	// PositionerPrimaryAngleIncrement (0018,1520) VR=DS VM=1-n Positioner Primary Angle Increment
	PositionerPrimaryAngleIncrement = New(0x0018, 0x1520)

	// PositionerSecondaryAngleIncrement (0018,1521) VR=DS VM=1-n Positioner Secondary Angle Increment
	PositionerSecondaryAngleIncrement = New(0x0018, 0x1521)

	// DetectorPrimaryAngle (0018,1530) VR=DS VM=1 Detector Primary Angle
	DetectorPrimaryAngle = New(0x0018, 0x1530)

	// DetectorSecondaryAngle (0018,1531) VR=DS VM=1 Detector Secondary Angle
	DetectorSecondaryAngle = New(0x0018, 0x1531)

	// ShutterShape (0018,1600) VR=CS VM=1-3 Shutter Shape
	ShutterShape = New(0x0018, 0x1600)

	// ShutterLeftVerticalEdge (0018,1602) VR=IS VM=1 Shutter Left Vertical Edge
	ShutterLeftVerticalEdge = New(0x0018, 0x1602)

	// ShutterRightVerticalEdge (0018,1604) VR=IS VM=1 Shutter Right Vertical Edge
	ShutterRightVerticalEdge = New(0x0018, 0x1604)

	// ShutterUpperHorizontalEdge (0018,1606) VR=IS VM=1 Shutter Upper Horizontal Edge
	ShutterUpperHorizontalEdge = New(0x0018, 0x1606)

	// ShutterLowerHorizontalEdge (0018,1608) VR=IS VM=1 Shutter Lower Horizontal Edge
	ShutterLowerHorizontalEdge = New(0x0018, 0x1608)

	// CenterOfCircularShutter (0018,1610) VR=IS VM=2 Center of Circular Shutter
	CenterOfCircularShutter = New(0x0018, 0x1610)

	// RadiusOfCircularShutter (0018,1612) VR=IS VM=1 Radius of Circular Shutter
	RadiusOfCircularShutter = New(0x0018, 0x1612)

	// VerticesOfThePolygonalShutter (0018,1620) VR=IS VM=2-2n Vertices of the Polygonal Shutter
	VerticesOfThePolygonalShutter = New(0x0018, 0x1620)

	// ShutterPresentationValue (0018,1622) VR=US VM=1 Shutter Presentation Value
	ShutterPresentationValue = New(0x0018, 0x1622)

	// ShutterOverlayGroup (0018,1623) VR=US VM=1 Shutter Overlay Group
	ShutterOverlayGroup = New(0x0018, 0x1623)

	// ShutterPresentationColorCIELabValue (0018,1624) VR=US VM=3 Shutter Presentation Color CIELab Value
	ShutterPresentationColorCIELabValue = New(0x0018, 0x1624)

	// OutlineShapeType (0018,1630) VR=CS VM=1 Outline Shape Type
	OutlineShapeType = New(0x0018, 0x1630)

	// OutlineLeftVerticalEdge (0018,1631) VR=FD VM=1 Outline Left Vertical Edge
	OutlineLeftVerticalEdge = New(0x0018, 0x1631)

	// OutlineRightVerticalEdge (0018,1632) VR=FD VM=1 Outline Right Vertical Edge
	OutlineRightVerticalEdge = New(0x0018, 0x1632)

	// OutlineUpperHorizontalEdge (0018,1633) VR=FD VM=1 Outline Upper Horizontal Edge
	OutlineUpperHorizontalEdge = New(0x0018, 0x1633)

	// OutlineLowerHorizontalEdge (0018,1634) VR=FD VM=1 Outline Lower Horizontal Edge
	OutlineLowerHorizontalEdge = New(0x0018, 0x1634)

	// CenterOfCircularOutline (0018,1635) VR=FD VM=2 Center of Circular Outline
	CenterOfCircularOutline = New(0x0018, 0x1635)

	// DiameterOfCircularOutline (0018,1636) VR=FD VM=1 Diameter of Circular Outline
	DiameterOfCircularOutline = New(0x0018, 0x1636)

	// NumberOfPolygonalVertices (0018,1637) VR=UL VM=1 Number of Polygonal Vertices
	NumberOfPolygonalVertices = New(0x0018, 0x1637)

	// VerticesOfThePolygonalOutline (0018,1638) VR=OF VM=1 Vertices of the Polygonal Outline
	VerticesOfThePolygonalOutline = New(0x0018, 0x1638)

	// CollimatorShape (0018,1700) VR=CS VM=1-3 Collimator Shape
	CollimatorShape = New(0x0018, 0x1700)

	// CollimatorLeftVerticalEdge (0018,1702) VR=IS VM=1 Collimator Left Vertical Edge
	CollimatorLeftVerticalEdge = New(0x0018, 0x1702)

	// CollimatorRightVerticalEdge (0018,1704) VR=IS VM=1 Collimator Right Vertical Edge
	CollimatorRightVerticalEdge = New(0x0018, 0x1704)

	// CollimatorUpperHorizontalEdge (0018,1706) VR=IS VM=1 Collimator Upper Horizontal Edge
	CollimatorUpperHorizontalEdge = New(0x0018, 0x1706)

	// CollimatorLowerHorizontalEdge (0018,1708) VR=IS VM=1 Collimator Lower Horizontal Edge
	CollimatorLowerHorizontalEdge = New(0x0018, 0x1708)

	// CenterOfCircularCollimator (0018,1710) VR=IS VM=2 Center of Circular Collimator
	CenterOfCircularCollimator = New(0x0018, 0x1710)

	// RadiusOfCircularCollimator (0018,1712) VR=IS VM=1 Radius of Circular Collimator
	RadiusOfCircularCollimator = New(0x0018, 0x1712)

	// VerticesOfThePolygonalCollimator (0018,1720) VR=IS VM=2-2n Vertices of the Polygonal Collimator
	VerticesOfThePolygonalCollimator = New(0x0018, 0x1720)

	// AcquisitionTimeSynchronized (0018,1800) VR=CS VM=1 Acquisition Time Synchronized
	AcquisitionTimeSynchronized = New(0x0018, 0x1800)

	// TimeSource (0018,1801) VR=SH VM=1 Time Source
	TimeSource = New(0x0018, 0x1801)

	// TimeDistributionProtocol (0018,1802) VR=CS VM=1 Time Distribution Protocol
	TimeDistributionProtocol = New(0x0018, 0x1802)

	// NTPSourceAddress (0018,1803) VR=LO VM=1 NTP Source Address
	NTPSourceAddress = New(0x0018, 0x1803)

	// PageNumberVector (0018,2001) VR=IS VM=1-n Page Number Vector
	PageNumberVector = New(0x0018, 0x2001)

	// FrameLabelVector (0018,2002) VR=SH VM=1-n Frame Label Vector
	FrameLabelVector = New(0x0018, 0x2002)

	// FramePrimaryAngleVector (0018,2003) VR=DS VM=1-n Frame Primary Angle Vector
	FramePrimaryAngleVector = New(0x0018, 0x2003)

	// FrameSecondaryAngleVector (0018,2004) VR=DS VM=1-n Frame Secondary Angle Vector
	FrameSecondaryAngleVector = New(0x0018, 0x2004)

	// SliceLocationVector (0018,2005) VR=DS VM=1-n Slice Location Vector
	SliceLocationVector = New(0x0018, 0x2005)

	// DisplayWindowLabelVector (0018,2006) VR=SH VM=1-n Display Window Label Vector
	DisplayWindowLabelVector = New(0x0018, 0x2006)

	// NominalScannedPixelSpacing (0018,2010) VR=DS VM=2 Nominal Scanned Pixel Spacing
	NominalScannedPixelSpacing = New(0x0018, 0x2010)

	// DigitizingDeviceTransportDirection (0018,2020) VR=CS VM=1 Digitizing Device Transport Direction
	DigitizingDeviceTransportDirection = New(0x0018, 0x2020)

	// RotationOfScannedFilm (0018,2030) VR=DS VM=1 Rotation of Scanned Film
	RotationOfScannedFilm = New(0x0018, 0x2030)

	// BiopsyTargetSequence (0018,2041) VR=SQ VM=1 Biopsy Target Sequence
	BiopsyTargetSequence = New(0x0018, 0x2041)

	// TargetUID (0018,2042) VR=UI VM=1 Target UID
	TargetUID = New(0x0018, 0x2042)

	// LocalizingCursorPosition (0018,2043) VR=FL VM=2 Localizing Cursor Position
	LocalizingCursorPosition = New(0x0018, 0x2043)

	// CalculatedTargetPosition (0018,2044) VR=FL VM=3 Calculated Target Position
	CalculatedTargetPosition = New(0x0018, 0x2044)

	// TargetLabel (0018,2045) VR=SH VM=1 Target Label
	TargetLabel = New(0x0018, 0x2045)

	// DisplayedZValue (0018,2046) VR=FL VM=1 Displayed Z Value
	DisplayedZValue = New(0x0018, 0x2046)

	// IVUSAcquisition (0018,3100) VR=CS VM=1 IVUS Acquisition
	IVUSAcquisition = New(0x0018, 0x3100)

	// IVUSPullbackRate (0018,3101) VR=DS VM=1 IVUS Pullback Rate
	IVUSPullbackRate = New(0x0018, 0x3101)

	// IVUSGatedRate (0018,3102) VR=DS VM=1 IVUS Gated Rate
	IVUSGatedRate = New(0x0018, 0x3102)

	// IVUSPullbackStartFrameNumber (0018,3103) VR=IS VM=1 IVUS Pullback Start Frame Number
	IVUSPullbackStartFrameNumber = New(0x0018, 0x3103)

	// IVUSPullbackStopFrameNumber (0018,3104) VR=IS VM=1 IVUS Pullback Stop Frame Number
	IVUSPullbackStopFrameNumber = New(0x0018, 0x3104)

	// LesionNumber (0018,3105) VR=IS VM=1-n Lesion Number
	LesionNumber = New(0x0018, 0x3105)

	// AcquisitionCommentsRETIRED (0018,4000) VR=LT VM=1 Acquisition Comments (RETIRED)
	AcquisitionCommentsRETIRED = New(0x0018, 0x4000)

	// OutputPower (0018,5000) VR=SH VM=1-n Output Power
	OutputPower = New(0x0018, 0x5000)

	// TransducerData (0018,5010) VR=LO VM=1-n Transducer Data
	TransducerData = New(0x0018, 0x5010)

	// TransducerIdentificationSequence (0018,5011) VR=SQ VM=1 Transducer Identification Sequence
	TransducerIdentificationSequence = New(0x0018, 0x5011)

	// FocusDepth (0018,5012) VR=DS VM=1 Focus Depth
	FocusDepth = New(0x0018, 0x5012)

	// ProcessingFunction (0018,5020) VR=LO VM=1 Processing Function
	ProcessingFunction = New(0x0018, 0x5020)

	// PostprocessingFunctionRETIRED (0018,5021) VR=LO VM=1 Postprocessing Function (RETIRED)
	PostprocessingFunctionRETIRED = New(0x0018, 0x5021)

	// MechanicalIndex (0018,5022) VR=DS VM=1 Mechanical Index
	MechanicalIndex = New(0x0018, 0x5022)

	// BoneThermalIndex (0018,5024) VR=DS VM=1 Bone Thermal Index
	BoneThermalIndex = New(0x0018, 0x5024)

	// CranialThermalIndex (0018,5026) VR=DS VM=1 Cranial Thermal Index
	CranialThermalIndex = New(0x0018, 0x5026)

	// SoftTissueThermalIndex (0018,5027) VR=DS VM=1 Soft Tissue Thermal Index
	SoftTissueThermalIndex = New(0x0018, 0x5027)

	// SoftTissueFocusThermalIndex (0018,5028) VR=DS VM=1 Soft Tissue-focus Thermal Index
	SoftTissueFocusThermalIndex = New(0x0018, 0x5028)

	// SoftTissueSurfaceThermalIndex (0018,5029) VR=DS VM=1 Soft Tissue-surface Thermal Index
	SoftTissueSurfaceThermalIndex = New(0x0018, 0x5029)

	// DynamicRangeRETIRED (0018,5030) VR=DS VM=1 Dynamic Range (RETIRED)
	DynamicRangeRETIRED = New(0x0018, 0x5030)

	// TotalGainRETIRED (0018,5040) VR=DS VM=1 Total Gain (RETIRED)
	TotalGainRETIRED = New(0x0018, 0x5040)

	// DepthOfScanField (0018,5050) VR=IS VM=1 Depth of Scan Field
	DepthOfScanField = New(0x0018, 0x5050)

	// PatientPosition (0018,5100) VR=CS VM=1 Patient Position
	PatientPosition = New(0x0018, 0x5100)

	// ViewPosition (0018,5101) VR=CS VM=1 View Position
	ViewPosition = New(0x0018, 0x5101)

	// ProjectionEponymousNameCodeSequence (0018,5104) VR=SQ VM=1 Projection Eponymous Name Code Sequence
	ProjectionEponymousNameCodeSequence = New(0x0018, 0x5104)

	// ImageTransformationMatrixRETIRED (0018,5210) VR=DS VM=6 Image Transformation Matrix (RETIRED)
	ImageTransformationMatrixRETIRED = New(0x0018, 0x5210)

	// ImageTranslationVectorRETIRED (0018,5212) VR=DS VM=3 Image Translation Vector (RETIRED)
	ImageTranslationVectorRETIRED = New(0x0018, 0x5212)

	// Sensitivity (0018,6000) VR=DS VM=1 Sensitivity
	Sensitivity = New(0x0018, 0x6000)

	// SequenceOfUltrasoundRegions (0018,6011) VR=SQ VM=1 Sequence of Ultrasound Regions
	SequenceOfUltrasoundRegions = New(0x0018, 0x6011)

	// RegionSpatialFormat (0018,6012) VR=US VM=1 Region Spatial Format
	RegionSpatialFormat = New(0x0018, 0x6012)

	// RegionDataType (0018,6014) VR=US VM=1 Region Data Type
	RegionDataType = New(0x0018, 0x6014)

	// RegionFlags (0018,6016) VR=UL VM=1 Region Flags
	RegionFlags = New(0x0018, 0x6016)

	// RegionLocationMinX0 (0018,6018) VR=UL VM=1 Region Location Min X0
	RegionLocationMinX0 = New(0x0018, 0x6018)

	// RegionLocationMinY0 (0018,601A) VR=UL VM=1 Region Location Min Y0
	RegionLocationMinY0 = New(0x0018, 0x601A)

	// RegionLocationMaxX1 (0018,601C) VR=UL VM=1 Region Location Max X1
	RegionLocationMaxX1 = New(0x0018, 0x601C)

	// RegionLocationMaxY1 (0018,601E) VR=UL VM=1 Region Location Max Y1
	RegionLocationMaxY1 = New(0x0018, 0x601E)

	// ReferencePixelX0 (0018,6020) VR=SL VM=1 Reference Pixel X0
	ReferencePixelX0 = New(0x0018, 0x6020)

	// ReferencePixelY0 (0018,6022) VR=SL VM=1 Reference Pixel Y0
	ReferencePixelY0 = New(0x0018, 0x6022)

	// PhysicalUnitsXDirection (0018,6024) VR=US VM=1 Physical Units X Direction
	PhysicalUnitsXDirection = New(0x0018, 0x6024)

	// PhysicalUnitsYDirection (0018,6026) VR=US VM=1 Physical Units Y Direction
	PhysicalUnitsYDirection = New(0x0018, 0x6026)

	// ReferencePixelPhysicalValueX (0018,6028) VR=FD VM=1 Reference Pixel Physical Value X
	ReferencePixelPhysicalValueX = New(0x0018, 0x6028)

	// ReferencePixelPhysicalValueY (0018,602A) VR=FD VM=1 Reference Pixel Physical Value Y
	ReferencePixelPhysicalValueY = New(0x0018, 0x602A)

	// PhysicalDeltaX (0018,602C) VR=FD VM=1 Physical Delta X
	PhysicalDeltaX = New(0x0018, 0x602C)

	// PhysicalDeltaY (0018,602E) VR=FD VM=1 Physical Delta Y
	PhysicalDeltaY = New(0x0018, 0x602E)

	// TransducerFrequency (0018,6030) VR=UL VM=1 Transducer Frequency
	TransducerFrequency = New(0x0018, 0x6030)

	// TransducerType (0018,6031) VR=CS VM=1 Transducer Type
	TransducerType = New(0x0018, 0x6031)

	// PulseRepetitionFrequency (0018,6032) VR=UL VM=1 Pulse Repetition Frequency
	PulseRepetitionFrequency = New(0x0018, 0x6032)

	// DopplerCorrectionAngle (0018,6034) VR=FD VM=1 Doppler Correction Angle
	DopplerCorrectionAngle = New(0x0018, 0x6034)

	// SteeringAngle (0018,6036) VR=FD VM=1 Steering Angle
	SteeringAngle = New(0x0018, 0x6036)

	// DopplerSampleVolumeXPositionRetiredRETIRED (0018,6038) VR=UL VM=1 Doppler Sample Volume X Position (Retired) (RETIRED)
	DopplerSampleVolumeXPositionRetiredRETIRED = New(0x0018, 0x6038)

	// DopplerSampleVolumeXPosition (0018,6039) VR=SL VM=1 Doppler Sample Volume X Position
	DopplerSampleVolumeXPosition = New(0x0018, 0x6039)

	// DopplerSampleVolumeYPositionRetiredRETIRED (0018,603A) VR=UL VM=1 Doppler Sample Volume Y Position (Retired) (RETIRED)
	DopplerSampleVolumeYPositionRetiredRETIRED = New(0x0018, 0x603A)

	// DopplerSampleVolumeYPosition (0018,603B) VR=SL VM=1 Doppler Sample Volume Y Position
	DopplerSampleVolumeYPosition = New(0x0018, 0x603B)

	// TMLinePositionX0RetiredRETIRED (0018,603C) VR=UL VM=1 TM-Line Position X0 (Retired) (RETIRED)
	TMLinePositionX0RetiredRETIRED = New(0x0018, 0x603C)

	// TMLinePositionX0 (0018,603D) VR=SL VM=1 TM-Line Position X0
	TMLinePositionX0 = New(0x0018, 0x603D)

	// TMLinePositionY0RetiredRETIRED (0018,603E) VR=UL VM=1 TM-Line Position Y0 (Retired) (RETIRED)
	TMLinePositionY0RetiredRETIRED = New(0x0018, 0x603E)

	// TMLinePositionY0 (0018,603F) VR=SL VM=1 TM-Line Position Y0
	TMLinePositionY0 = New(0x0018, 0x603F)

	// TMLinePositionX1RetiredRETIRED (0018,6040) VR=UL VM=1 TM-Line Position X1 (Retired) (RETIRED)
	TMLinePositionX1RetiredRETIRED = New(0x0018, 0x6040)

	// TMLinePositionX1 (0018,6041) VR=SL VM=1 TM-Line Position X1
	TMLinePositionX1 = New(0x0018, 0x6041)

	// TMLinePositionY1RetiredRETIRED (0018,6042) VR=UL VM=1 TM-Line Position Y1 (Retired) (RETIRED)
	TMLinePositionY1RetiredRETIRED = New(0x0018, 0x6042)

	// TMLinePositionY1 (0018,6043) VR=SL VM=1 TM-Line Position Y1
	TMLinePositionY1 = New(0x0018, 0x6043)

	// PixelComponentOrganization (0018,6044) VR=US VM=1 Pixel Component Organization
	PixelComponentOrganization = New(0x0018, 0x6044)

	// PixelComponentMask (0018,6046) VR=UL VM=1 Pixel Component Mask
	PixelComponentMask = New(0x0018, 0x6046)

	// PixelComponentRangeStart (0018,6048) VR=UL VM=1 Pixel Component Range Start
	PixelComponentRangeStart = New(0x0018, 0x6048)

	// PixelComponentRangeStop (0018,604A) VR=UL VM=1 Pixel Component Range Stop
	PixelComponentRangeStop = New(0x0018, 0x604A)

	// PixelComponentPhysicalUnits (0018,604C) VR=US VM=1 Pixel Component Physical Units
	PixelComponentPhysicalUnits = New(0x0018, 0x604C)

	// PixelComponentDataType (0018,604E) VR=US VM=1 Pixel Component Data Type
	PixelComponentDataType = New(0x0018, 0x604E)

	// NumberOfTableBreakPoints (0018,6050) VR=UL VM=1 Number of Table Break Points
	NumberOfTableBreakPoints = New(0x0018, 0x6050)

	// TableOfXBreakPoints (0018,6052) VR=UL VM=1-n Table of X Break Points
	TableOfXBreakPoints = New(0x0018, 0x6052)

	// TableOfYBreakPoints (0018,6054) VR=FD VM=1-n Table of Y Break Points
	TableOfYBreakPoints = New(0x0018, 0x6054)

	// NumberOfTableEntries (0018,6056) VR=UL VM=1 Number of Table Entries
	NumberOfTableEntries = New(0x0018, 0x6056)

	// TableOfPixelValues (0018,6058) VR=UL VM=1-n Table of Pixel Values
	TableOfPixelValues = New(0x0018, 0x6058)

	// TableOfParameterValues (0018,605A) VR=FL VM=1-n Table of Parameter Values
	TableOfParameterValues = New(0x0018, 0x605A)

	// RWaveTimeVector (0018,6060) VR=FL VM=1-n R Wave Time Vector
	RWaveTimeVector = New(0x0018, 0x6060)

	// ActiveImageAreaOverlayGroup (0018,6070) VR=US VM=1 Active Image Area Overlay Group
	ActiveImageAreaOverlayGroup = New(0x0018, 0x6070)

	// DetectorConditionsNominalFlag (0018,7000) VR=CS VM=1 Detector Conditions Nominal Flag
	DetectorConditionsNominalFlag = New(0x0018, 0x7000)

	// DetectorTemperature (0018,7001) VR=DS VM=1 Detector Temperature
	DetectorTemperature = New(0x0018, 0x7001)

	// DetectorType (0018,7004) VR=CS VM=1 Detector Type
	DetectorType = New(0x0018, 0x7004)

	// DetectorConfiguration (0018,7005) VR=CS VM=1 Detector Configuration
	DetectorConfiguration = New(0x0018, 0x7005)

	// DetectorDescription (0018,7006) VR=LT VM=1 Detector Description
	DetectorDescription = New(0x0018, 0x7006)

	// DetectorMode (0018,7008) VR=LT VM=1 Detector Mode
	DetectorMode = New(0x0018, 0x7008)

	// DetectorID (0018,700A) VR=SH VM=1 Detector ID
	DetectorID = New(0x0018, 0x700A)

	// DateOfLastDetectorCalibration (0018,700C) VR=DA VM=1 Date of Last Detector Calibration
	DateOfLastDetectorCalibration = New(0x0018, 0x700C)

	// TimeOfLastDetectorCalibration (0018,700E) VR=TM VM=1 Time of Last Detector Calibration
	TimeOfLastDetectorCalibration = New(0x0018, 0x700E)

	// ExposuresOnDetectorSinceLastCalibration (0018,7010) VR=IS VM=1 Exposures on Detector Since Last Calibration
	ExposuresOnDetectorSinceLastCalibration = New(0x0018, 0x7010)

	// ExposuresOnDetectorSinceManufactured (0018,7011) VR=IS VM=1 Exposures on Detector Since Manufactured
	ExposuresOnDetectorSinceManufactured = New(0x0018, 0x7011)

	// DetectorTimeSinceLastExposure (0018,7012) VR=DS VM=1 Detector Time Since Last Exposure
	DetectorTimeSinceLastExposure = New(0x0018, 0x7012)

	// DetectorActiveTime (0018,7014) VR=DS VM=1 Detector Active Time
	DetectorActiveTime = New(0x0018, 0x7014)

	// DetectorActivationOffsetFromExposure (0018,7016) VR=DS VM=1 Detector Activation Offset From Exposure
	DetectorActivationOffsetFromExposure = New(0x0018, 0x7016)

	// DetectorBinning (0018,701A) VR=DS VM=2 Detector Binning
	DetectorBinning = New(0x0018, 0x701A)

	// DetectorElementPhysicalSize (0018,7020) VR=DS VM=2 Detector Element Physical Size
	DetectorElementPhysicalSize = New(0x0018, 0x7020)

	// DetectorElementSpacing (0018,7022) VR=DS VM=2 Detector Element Spacing
	DetectorElementSpacing = New(0x0018, 0x7022)

	// DetectorActiveShape (0018,7024) VR=CS VM=1 Detector Active Shape
	DetectorActiveShape = New(0x0018, 0x7024)

	// DetectorActiveDimensions (0018,7026) VR=DS VM=1-2 Detector Active Dimension(s)
	DetectorActiveDimensions = New(0x0018, 0x7026)

	// DetectorActiveOrigin (0018,7028) VR=DS VM=2 Detector Active Origin
	DetectorActiveOrigin = New(0x0018, 0x7028)

	// DetectorManufacturerName (0018,702A) VR=LO VM=1 Detector Manufacturer Name
	DetectorManufacturerName = New(0x0018, 0x702A)

	// DetectorManufacturerModelName (0018,702B) VR=LO VM=1 Detector Manufacturer's Model Name
	DetectorManufacturerModelName = New(0x0018, 0x702B)

	// FieldOfViewOrigin (0018,7030) VR=DS VM=2 Field of View Origin
	FieldOfViewOrigin = New(0x0018, 0x7030)

	// FieldOfViewRotation (0018,7032) VR=DS VM=1 Field of View Rotation
	FieldOfViewRotation = New(0x0018, 0x7032)

	// FieldOfViewHorizontalFlip (0018,7034) VR=CS VM=1 Field of View Horizontal Flip
	FieldOfViewHorizontalFlip = New(0x0018, 0x7034)

	// PixelDataAreaOriginRelativeToFOV (0018,7036) VR=FL VM=2 Pixel Data Area Origin Relative To FOV
	PixelDataAreaOriginRelativeToFOV = New(0x0018, 0x7036)

	// PixelDataAreaRotationAngleRelativeToFOV (0018,7038) VR=FL VM=1 Pixel Data Area Rotation Angle Relative To FOV
	PixelDataAreaRotationAngleRelativeToFOV = New(0x0018, 0x7038)

	// GridAbsorbingMaterial (0018,7040) VR=LT VM=1 Grid Absorbing Material
	GridAbsorbingMaterial = New(0x0018, 0x7040)

	// GridSpacingMaterial (0018,7041) VR=LT VM=1 Grid Spacing Material
	GridSpacingMaterial = New(0x0018, 0x7041)

	// GridThickness (0018,7042) VR=DS VM=1 Grid Thickness
	GridThickness = New(0x0018, 0x7042)

	// GridPitch (0018,7044) VR=DS VM=1 Grid Pitch
	GridPitch = New(0x0018, 0x7044)

	// GridAspectRatio (0018,7046) VR=IS VM=2 Grid Aspect Ratio
	GridAspectRatio = New(0x0018, 0x7046)

	// GridPeriod (0018,7048) VR=DS VM=1 Grid Period
	GridPeriod = New(0x0018, 0x7048)

	// GridFocalDistance (0018,704C) VR=DS VM=1 Grid Focal Distance
	GridFocalDistance = New(0x0018, 0x704C)

	// FilterMaterial (0018,7050) VR=CS VM=1-n Filter Material
	FilterMaterial = New(0x0018, 0x7050)

	// FilterThicknessMinimum (0018,7052) VR=DS VM=1-n Filter Thickness Minimum
	FilterThicknessMinimum = New(0x0018, 0x7052)

	// FilterThicknessMaximum (0018,7054) VR=DS VM=1-n Filter Thickness Maximum
	FilterThicknessMaximum = New(0x0018, 0x7054)

	// FilterBeamPathLengthMinimum (0018,7056) VR=FL VM=1-n Filter Beam Path Length Minimum
	FilterBeamPathLengthMinimum = New(0x0018, 0x7056)

	// FilterBeamPathLengthMaximum (0018,7058) VR=FL VM=1-n Filter Beam Path Length Maximum
	FilterBeamPathLengthMaximum = New(0x0018, 0x7058)

	// ExposureControlMode (0018,7060) VR=CS VM=1 Exposure Control Mode
	ExposureControlMode = New(0x0018, 0x7060)

	// ExposureControlModeDescription (0018,7062) VR=LT VM=1 Exposure Control Mode Description
	ExposureControlModeDescription = New(0x0018, 0x7062)

	// ExposureStatus (0018,7064) VR=CS VM=1 Exposure Status
	ExposureStatus = New(0x0018, 0x7064)

	// PhototimerSetting (0018,7065) VR=DS VM=1 Phototimer Setting
	PhototimerSetting = New(0x0018, 0x7065)

	// ExposureTimeInuS (0018,8150) VR=DS VM=1 Exposure Time in µS
	ExposureTimeInuS = New(0x0018, 0x8150)

	// XRayTubeCurrentInuA (0018,8151) VR=DS VM=1 X-Ray Tube Current in µA
	XRayTubeCurrentInuA = New(0x0018, 0x8151)

	// ContentQualification (0018,9004) VR=CS VM=1 Content Qualification
	ContentQualification = New(0x0018, 0x9004)

	// PulseSequenceName (0018,9005) VR=SH VM=1 Pulse Sequence Name
	PulseSequenceName = New(0x0018, 0x9005)

	// MRImagingModifierSequence (0018,9006) VR=SQ VM=1 MR Imaging Modifier Sequence
	MRImagingModifierSequence = New(0x0018, 0x9006)

	// EchoPulseSequence (0018,9008) VR=CS VM=1 Echo Pulse Sequence
	EchoPulseSequence = New(0x0018, 0x9008)

	// InversionRecovery (0018,9009) VR=CS VM=1 Inversion Recovery
	InversionRecovery = New(0x0018, 0x9009)

	// FlowCompensation (0018,9010) VR=CS VM=1 Flow Compensation
	FlowCompensation = New(0x0018, 0x9010)

	// MultipleSpinEcho (0018,9011) VR=CS VM=1 Multiple Spin Echo
	MultipleSpinEcho = New(0x0018, 0x9011)

	// MultiPlanarExcitation (0018,9012) VR=CS VM=1 Multi-planar Excitation
	MultiPlanarExcitation = New(0x0018, 0x9012)

	// PhaseContrast (0018,9014) VR=CS VM=1 Phase Contrast
	PhaseContrast = New(0x0018, 0x9014)

	// TimeOfFlightContrast (0018,9015) VR=CS VM=1 Time of Flight Contrast
	TimeOfFlightContrast = New(0x0018, 0x9015)

	// Spoiling (0018,9016) VR=CS VM=1 Spoiling
	Spoiling = New(0x0018, 0x9016)

	// SteadyStatePulseSequence (0018,9017) VR=CS VM=1 Steady State Pulse Sequence
	SteadyStatePulseSequence = New(0x0018, 0x9017)

	// EchoPlanarPulseSequence (0018,9018) VR=CS VM=1 Echo Planar Pulse Sequence
	EchoPlanarPulseSequence = New(0x0018, 0x9018)

	// TagAngleFirstAxis (0018,9019) VR=FD VM=1 Tag Angle First Axis
	TagAngleFirstAxis = New(0x0018, 0x9019)

	// MagnetizationTransfer (0018,9020) VR=CS VM=1 Magnetization Transfer
	MagnetizationTransfer = New(0x0018, 0x9020)

	// T2Preparation (0018,9021) VR=CS VM=1 T2 Preparation
	T2Preparation = New(0x0018, 0x9021)

	// BloodSignalNulling (0018,9022) VR=CS VM=1 Blood Signal Nulling
	BloodSignalNulling = New(0x0018, 0x9022)

	// SaturationRecovery (0018,9024) VR=CS VM=1 Saturation Recovery
	SaturationRecovery = New(0x0018, 0x9024)

	// SpectrallySelectedSuppression (0018,9025) VR=CS VM=1 Spectrally Selected Suppression
	SpectrallySelectedSuppression = New(0x0018, 0x9025)

	// SpectrallySelectedExcitation (0018,9026) VR=CS VM=1 Spectrally Selected Excitation
	SpectrallySelectedExcitation = New(0x0018, 0x9026)

	// SpatialPresaturation (0018,9027) VR=CS VM=1 Spatial Pre-saturation
	SpatialPresaturation = New(0x0018, 0x9027)

	// Tagging (0018,9028) VR=CS VM=1 Tagging
	Tagging = New(0x0018, 0x9028)

	// OversamplingPhase (0018,9029) VR=CS VM=1 Oversampling Phase
	OversamplingPhase = New(0x0018, 0x9029)

	// TagSpacingFirstDimension (0018,9030) VR=FD VM=1 Tag Spacing First Dimension
	TagSpacingFirstDimension = New(0x0018, 0x9030)

	// GeometryOfKSpaceTraversal (0018,9032) VR=CS VM=1 Geometry of k-Space Traversal
	GeometryOfKSpaceTraversal = New(0x0018, 0x9032)

	// SegmentedKSpaceTraversal (0018,9033) VR=CS VM=1 Segmented k-Space Traversal
	SegmentedKSpaceTraversal = New(0x0018, 0x9033)

	// RectilinearPhaseEncodeReordering (0018,9034) VR=CS VM=1 Rectilinear Phase Encode Reordering
	RectilinearPhaseEncodeReordering = New(0x0018, 0x9034)

	// TagThickness (0018,9035) VR=FD VM=1 Tag Thickness
	TagThickness = New(0x0018, 0x9035)

	// PartialFourierDirection (0018,9036) VR=CS VM=1 Partial Fourier Direction
	PartialFourierDirection = New(0x0018, 0x9036)

	// CardiacSynchronizationTechnique (0018,9037) VR=CS VM=1 Cardiac Synchronization Technique
	CardiacSynchronizationTechnique = New(0x0018, 0x9037)

	// ReceiveCoilManufacturerName (0018,9041) VR=LO VM=1 Receive Coil Manufacturer Name
	ReceiveCoilManufacturerName = New(0x0018, 0x9041)

	// MRReceiveCoilSequence (0018,9042) VR=SQ VM=1 MR Receive Coil Sequence
	MRReceiveCoilSequence = New(0x0018, 0x9042)

	// ReceiveCoilType (0018,9043) VR=CS VM=1 Receive Coil Type
	ReceiveCoilType = New(0x0018, 0x9043)

	// QuadratureReceiveCoil (0018,9044) VR=CS VM=1 Quadrature Receive Coil
	QuadratureReceiveCoil = New(0x0018, 0x9044)

	// MultiCoilDefinitionSequence (0018,9045) VR=SQ VM=1 Multi-Coil Definition Sequence
	MultiCoilDefinitionSequence = New(0x0018, 0x9045)

	// MultiCoilConfiguration (0018,9046) VR=LO VM=1 Multi-Coil Configuration
	MultiCoilConfiguration = New(0x0018, 0x9046)

	// MultiCoilElementName (0018,9047) VR=SH VM=1 Multi-Coil Element Name
	MultiCoilElementName = New(0x0018, 0x9047)

	// MultiCoilElementUsed (0018,9048) VR=CS VM=1 Multi-Coil Element Used
	MultiCoilElementUsed = New(0x0018, 0x9048)

	// MRTransmitCoilSequence (0018,9049) VR=SQ VM=1 MR Transmit Coil Sequence
	MRTransmitCoilSequence = New(0x0018, 0x9049)

	// TransmitCoilManufacturerName (0018,9050) VR=LO VM=1 Transmit Coil Manufacturer Name
	TransmitCoilManufacturerName = New(0x0018, 0x9050)

	// TransmitCoilType (0018,9051) VR=CS VM=1 Transmit Coil Type
	TransmitCoilType = New(0x0018, 0x9051)

	// SpectralWidth (0018,9052) VR=FD VM=1-2 Spectral Width
	SpectralWidth = New(0x0018, 0x9052)

	// ChemicalShiftReference (0018,9053) VR=FD VM=1-2 Chemical Shift Reference
	ChemicalShiftReference = New(0x0018, 0x9053)

	// VolumeLocalizationTechnique (0018,9054) VR=CS VM=1 Volume Localization Technique
	VolumeLocalizationTechnique = New(0x0018, 0x9054)

	// MRAcquisitionFrequencyEncodingSteps (0018,9058) VR=US VM=1 MR Acquisition Frequency Encoding Steps
	MRAcquisitionFrequencyEncodingSteps = New(0x0018, 0x9058)

	// Decoupling (0018,9059) VR=CS VM=1 De-coupling
	Decoupling = New(0x0018, 0x9059)

	// DecoupledNucleus (0018,9060) VR=CS VM=1-2 De-coupled Nucleus
	DecoupledNucleus = New(0x0018, 0x9060)

	// DecouplingFrequency (0018,9061) VR=FD VM=1-2 De-coupling Frequency
	DecouplingFrequency = New(0x0018, 0x9061)

	// DecouplingMethod (0018,9062) VR=CS VM=1 De-coupling Method
	DecouplingMethod = New(0x0018, 0x9062)

	// DecouplingChemicalShiftReference (0018,9063) VR=FD VM=1-2 De-coupling Chemical Shift Reference
	DecouplingChemicalShiftReference = New(0x0018, 0x9063)

	// KSpaceFiltering (0018,9064) VR=CS VM=1 k-space Filtering
	KSpaceFiltering = New(0x0018, 0x9064)

	// TimeDomainFiltering (0018,9065) VR=CS VM=1-2 Time Domain Filtering
	TimeDomainFiltering = New(0x0018, 0x9065)

	// NumberOfZeroFills (0018,9066) VR=US VM=1-2 Number of Zero Fills
	NumberOfZeroFills = New(0x0018, 0x9066)

	// BaselineCorrection (0018,9067) VR=CS VM=1 Baseline Correction
	BaselineCorrection = New(0x0018, 0x9067)

	// ParallelReductionFactorInPlane (0018,9069) VR=FD VM=1 Parallel Reduction Factor In-plane
	ParallelReductionFactorInPlane = New(0x0018, 0x9069)

	// CardiacRRIntervalSpecified (0018,9070) VR=FD VM=1 Cardiac R-R Interval Specified
	CardiacRRIntervalSpecified = New(0x0018, 0x9070)

	// AcquisitionDuration (0018,9073) VR=FD VM=1 Acquisition Duration
	AcquisitionDuration = New(0x0018, 0x9073)

	// FrameAcquisitionDateTime (0018,9074) VR=DT VM=1 Frame Acquisition DateTime
	FrameAcquisitionDateTime = New(0x0018, 0x9074)

	// DiffusionDirectionality (0018,9075) VR=CS VM=1 Diffusion Directionality
	DiffusionDirectionality = New(0x0018, 0x9075)

	// DiffusionGradientDirectionSequence (0018,9076) VR=SQ VM=1 Diffusion Gradient Direction Sequence
	DiffusionGradientDirectionSequence = New(0x0018, 0x9076)

	// ParallelAcquisition (0018,9077) VR=CS VM=1 Parallel Acquisition
	ParallelAcquisition = New(0x0018, 0x9077)

	// ParallelAcquisitionTechnique (0018,9078) VR=CS VM=1 Parallel Acquisition Technique
	ParallelAcquisitionTechnique = New(0x0018, 0x9078)

	// InversionTimes (0018,9079) VR=FD VM=1-n Inversion Times
	InversionTimes = New(0x0018, 0x9079)

	// MetaboliteMapDescription (0018,9080) VR=ST VM=1 Metabolite Map Description
	MetaboliteMapDescription = New(0x0018, 0x9080)

	// PartialFourier (0018,9081) VR=CS VM=1 Partial Fourier
	PartialFourier = New(0x0018, 0x9081)

	// EffectiveEchoTime (0018,9082) VR=FD VM=1 Effective Echo Time
	EffectiveEchoTime = New(0x0018, 0x9082)

	// MetaboliteMapCodeSequence (0018,9083) VR=SQ VM=1 Metabolite Map Code Sequence
	MetaboliteMapCodeSequence = New(0x0018, 0x9083)

	// ChemicalShiftSequence (0018,9084) VR=SQ VM=1 Chemical Shift Sequence
	ChemicalShiftSequence = New(0x0018, 0x9084)

	// CardiacSignalSource (0018,9085) VR=CS VM=1 Cardiac Signal Source
	CardiacSignalSource = New(0x0018, 0x9085)

	// DiffusionBValue (0018,9087) VR=FD VM=1 Diffusion b-value
	DiffusionBValue = New(0x0018, 0x9087)

	// DiffusionGradientOrientation (0018,9089) VR=FD VM=3 Diffusion Gradient Orientation
	DiffusionGradientOrientation = New(0x0018, 0x9089)

	// VelocityEncodingDirection (0018,9090) VR=FD VM=3 Velocity Encoding Direction
	VelocityEncodingDirection = New(0x0018, 0x9090)

	// VelocityEncodingMinimumValue (0018,9091) VR=FD VM=1 Velocity Encoding Minimum Value
	VelocityEncodingMinimumValue = New(0x0018, 0x9091)

	// VelocityEncodingAcquisitionSequence (0018,9092) VR=SQ VM=1 Velocity Encoding Acquisition Sequence
	VelocityEncodingAcquisitionSequence = New(0x0018, 0x9092)

	// NumberOfKSpaceTrajectories (0018,9093) VR=US VM=1 Number of k-Space Trajectories
	NumberOfKSpaceTrajectories = New(0x0018, 0x9093)

	// CoverageOfKSpace (0018,9094) VR=CS VM=1 Coverage of k-Space
	CoverageOfKSpace = New(0x0018, 0x9094)

	// SpectroscopyAcquisitionPhaseRows (0018,9095) VR=UL VM=1 Spectroscopy Acquisition Phase Rows
	SpectroscopyAcquisitionPhaseRows = New(0x0018, 0x9095)

	// ParallelReductionFactorInPlaneRetiredRETIRED (0018,9096) VR=FD VM=1 Parallel Reduction Factor In-plane (Retired) (RETIRED)
	ParallelReductionFactorInPlaneRetiredRETIRED = New(0x0018, 0x9096)

	// TransmitterFrequency (0018,9098) VR=FD VM=1-2 Transmitter Frequency
	TransmitterFrequency = New(0x0018, 0x9098)

	// ResonantNucleus (0018,9100) VR=CS VM=1-2 Resonant Nucleus
	ResonantNucleus = New(0x0018, 0x9100)

	// FrequencyCorrection (0018,9101) VR=CS VM=1 Frequency Correction
	FrequencyCorrection = New(0x0018, 0x9101)

	// MRSpectroscopyFOVGeometrySequence (0018,9103) VR=SQ VM=1 MR Spectroscopy FOV/Geometry Sequence
	MRSpectroscopyFOVGeometrySequence = New(0x0018, 0x9103)

	// SlabThickness (0018,9104) VR=FD VM=1 Slab Thickness
	SlabThickness = New(0x0018, 0x9104)

	// SlabOrientation (0018,9105) VR=FD VM=3 Slab Orientation
	SlabOrientation = New(0x0018, 0x9105)

	// MidSlabPosition (0018,9106) VR=FD VM=3 Mid Slab Position
	MidSlabPosition = New(0x0018, 0x9106)

	// MRSpatialSaturationSequence (0018,9107) VR=SQ VM=1 MR Spatial Saturation Sequence
	MRSpatialSaturationSequence = New(0x0018, 0x9107)

	// MRTimingAndRelatedParametersSequence (0018,9112) VR=SQ VM=1 MR Timing and Related Parameters Sequence
	MRTimingAndRelatedParametersSequence = New(0x0018, 0x9112)

	// MREchoSequence (0018,9114) VR=SQ VM=1 MR Echo Sequence
	MREchoSequence = New(0x0018, 0x9114)

	// MRModifierSequence (0018,9115) VR=SQ VM=1 MR Modifier Sequence
	MRModifierSequence = New(0x0018, 0x9115)

	// MRDiffusionSequence (0018,9117) VR=SQ VM=1 MR Diffusion Sequence
	MRDiffusionSequence = New(0x0018, 0x9117)

	// CardiacSynchronizationSequence (0018,9118) VR=SQ VM=1 Cardiac Synchronization Sequence
	CardiacSynchronizationSequence = New(0x0018, 0x9118)

	// MRAveragesSequence (0018,9119) VR=SQ VM=1 MR Averages Sequence
	MRAveragesSequence = New(0x0018, 0x9119)

	// MRFOVGeometrySequence (0018,9125) VR=SQ VM=1 MR FOV/Geometry Sequence
	MRFOVGeometrySequence = New(0x0018, 0x9125)

	// VolumeLocalizationSequence (0018,9126) VR=SQ VM=1 Volume Localization Sequence
	VolumeLocalizationSequence = New(0x0018, 0x9126)

	// SpectroscopyAcquisitionDataColumns (0018,9127) VR=UL VM=1 Spectroscopy Acquisition Data Columns
	SpectroscopyAcquisitionDataColumns = New(0x0018, 0x9127)

	// DiffusionAnisotropyType (0018,9147) VR=CS VM=1 Diffusion Anisotropy Type
	DiffusionAnisotropyType = New(0x0018, 0x9147)

	// FrameReferenceDateTime (0018,9151) VR=DT VM=1 Frame Reference DateTime
	FrameReferenceDateTime = New(0x0018, 0x9151)

	// MRMetaboliteMapSequence (0018,9152) VR=SQ VM=1 MR Metabolite Map Sequence
	MRMetaboliteMapSequence = New(0x0018, 0x9152)

	// ParallelReductionFactorOutOfPlane (0018,9155) VR=FD VM=1 Parallel Reduction Factor out-of-plane
	ParallelReductionFactorOutOfPlane = New(0x0018, 0x9155)

	// SpectroscopyAcquisitionOutOfPlanePhaseSteps (0018,9159) VR=UL VM=1 Spectroscopy Acquisition Out-of-plane Phase Steps
	SpectroscopyAcquisitionOutOfPlanePhaseSteps = New(0x0018, 0x9159)

	// BulkMotionStatusRETIRED (0018,9166) VR=CS VM=1 Bulk Motion Status (RETIRED)
	BulkMotionStatusRETIRED = New(0x0018, 0x9166)

	// ParallelReductionFactorSecondInPlane (0018,9168) VR=FD VM=1 Parallel Reduction Factor Second In-plane
	ParallelReductionFactorSecondInPlane = New(0x0018, 0x9168)

	// CardiacBeatRejectionTechnique (0018,9169) VR=CS VM=1 Cardiac Beat Rejection Technique
	CardiacBeatRejectionTechnique = New(0x0018, 0x9169)

	// RespiratoryMotionCompensationTechnique (0018,9170) VR=CS VM=1 Respiratory Motion Compensation Technique
	RespiratoryMotionCompensationTechnique = New(0x0018, 0x9170)

	// RespiratorySignalSource (0018,9171) VR=CS VM=1 Respiratory Signal Source
	RespiratorySignalSource = New(0x0018, 0x9171)

	// BulkMotionCompensationTechnique (0018,9172) VR=CS VM=1 Bulk Motion Compensation Technique
	BulkMotionCompensationTechnique = New(0x0018, 0x9172)

	// BulkMotionSignalSource (0018,9173) VR=CS VM=1 Bulk Motion Signal Source
	BulkMotionSignalSource = New(0x0018, 0x9173)

	// ApplicableSafetyStandardAgency (0018,9174) VR=CS VM=1 Applicable Safety Standard Agency
	ApplicableSafetyStandardAgency = New(0x0018, 0x9174)

	// ApplicableSafetyStandardDescription (0018,9175) VR=LO VM=1 Applicable Safety Standard Description
	ApplicableSafetyStandardDescription = New(0x0018, 0x9175)

	// OperatingModeSequence (0018,9176) VR=SQ VM=1 Operating Mode Sequence
	OperatingModeSequence = New(0x0018, 0x9176)

	// OperatingModeType (0018,9177) VR=CS VM=1 Operating Mode Type
	OperatingModeType = New(0x0018, 0x9177)

	// OperatingMode (0018,9178) VR=CS VM=1 Operating Mode
	OperatingMode = New(0x0018, 0x9178)

	// SpecificAbsorptionRateDefinition (0018,9179) VR=CS VM=1 Specific Absorption Rate Definition
	SpecificAbsorptionRateDefinition = New(0x0018, 0x9179)

	// GradientOutputType (0018,9180) VR=CS VM=1 Gradient Output Type
	GradientOutputType = New(0x0018, 0x9180)

	// SpecificAbsorptionRateValue (0018,9181) VR=FD VM=1 Specific Absorption Rate Value
	SpecificAbsorptionRateValue = New(0x0018, 0x9181)

	// GradientOutput (0018,9182) VR=FD VM=1 Gradient Output
	GradientOutput = New(0x0018, 0x9182)

	// FlowCompensationDirection (0018,9183) VR=CS VM=1 Flow Compensation Direction
	FlowCompensationDirection = New(0x0018, 0x9183)

	// TaggingDelay (0018,9184) VR=FD VM=1 Tagging Delay
	TaggingDelay = New(0x0018, 0x9184)

	// RespiratoryMotionCompensationTechniqueDescription (0018,9185) VR=ST VM=1 Respiratory Motion Compensation Technique Description
	RespiratoryMotionCompensationTechniqueDescription = New(0x0018, 0x9185)

	// RespiratorySignalSourceID (0018,9186) VR=SH VM=1 Respiratory Signal Source ID
	RespiratorySignalSourceID = New(0x0018, 0x9186)

	// ChemicalShiftMinimumIntegrationLimitInHzRETIRED (0018,9195) VR=FD VM=1 Chemical Shift Minimum Integration Limit in Hz (RETIRED)
	ChemicalShiftMinimumIntegrationLimitInHzRETIRED = New(0x0018, 0x9195)

	// ChemicalShiftMaximumIntegrationLimitInHzRETIRED (0018,9196) VR=FD VM=1 Chemical Shift Maximum Integration Limit in Hz (RETIRED)
	ChemicalShiftMaximumIntegrationLimitInHzRETIRED = New(0x0018, 0x9196)

	// MRVelocityEncodingSequence (0018,9197) VR=SQ VM=1 MR Velocity Encoding Sequence
	MRVelocityEncodingSequence = New(0x0018, 0x9197)

	// FirstOrderPhaseCorrection (0018,9198) VR=CS VM=1 First Order Phase Correction
	FirstOrderPhaseCorrection = New(0x0018, 0x9198)

	// WaterReferencedPhaseCorrection (0018,9199) VR=CS VM=1 Water Referenced Phase Correction
	WaterReferencedPhaseCorrection = New(0x0018, 0x9199)

	// MRSpectroscopyAcquisitionType (0018,9200) VR=CS VM=1 MR Spectroscopy Acquisition Type
	MRSpectroscopyAcquisitionType = New(0x0018, 0x9200)

	// RespiratoryCyclePosition (0018,9214) VR=CS VM=1 Respiratory Cycle Position
	RespiratoryCyclePosition = New(0x0018, 0x9214)

	// VelocityEncodingMaximumValue (0018,9217) VR=FD VM=1 Velocity Encoding Maximum Value
	VelocityEncodingMaximumValue = New(0x0018, 0x9217)

	// TagSpacingSecondDimension (0018,9218) VR=FD VM=1 Tag Spacing Second Dimension
	TagSpacingSecondDimension = New(0x0018, 0x9218)

	// TagAngleSecondAxis (0018,9219) VR=SS VM=1 Tag Angle Second Axis
	TagAngleSecondAxis = New(0x0018, 0x9219)

	// FrameAcquisitionDuration (0018,9220) VR=FD VM=1 Frame Acquisition Duration
	FrameAcquisitionDuration = New(0x0018, 0x9220)

	// MRImageFrameTypeSequence (0018,9226) VR=SQ VM=1 MR Image Frame Type Sequence
	MRImageFrameTypeSequence = New(0x0018, 0x9226)

	// MRSpectroscopyFrameTypeSequence (0018,9227) VR=SQ VM=1 MR Spectroscopy Frame Type Sequence
	MRSpectroscopyFrameTypeSequence = New(0x0018, 0x9227)

	// MRAcquisitionPhaseEncodingStepsInPlane (0018,9231) VR=US VM=1 MR Acquisition Phase Encoding Steps in-plane
	MRAcquisitionPhaseEncodingStepsInPlane = New(0x0018, 0x9231)

	// MRAcquisitionPhaseEncodingStepsOutOfPlane (0018,9232) VR=US VM=1 MR Acquisition Phase Encoding Steps out-of-plane
	MRAcquisitionPhaseEncodingStepsOutOfPlane = New(0x0018, 0x9232)

	// SpectroscopyAcquisitionPhaseColumns (0018,9234) VR=UL VM=1 Spectroscopy Acquisition Phase Columns
	SpectroscopyAcquisitionPhaseColumns = New(0x0018, 0x9234)

	// CardiacCyclePosition (0018,9236) VR=CS VM=1 Cardiac Cycle Position
	CardiacCyclePosition = New(0x0018, 0x9236)

	// SpecificAbsorptionRateSequence (0018,9239) VR=SQ VM=1 Specific Absorption Rate Sequence
	SpecificAbsorptionRateSequence = New(0x0018, 0x9239)

	// RFEchoTrainLength (0018,9240) VR=US VM=1 RF Echo Train Length
	RFEchoTrainLength = New(0x0018, 0x9240)

	// GradientEchoTrainLength (0018,9241) VR=US VM=1 Gradient Echo Train Length
	GradientEchoTrainLength = New(0x0018, 0x9241)

	// ArterialSpinLabelingContrast (0018,9250) VR=CS VM=1 Arterial Spin Labeling Contrast
	ArterialSpinLabelingContrast = New(0x0018, 0x9250)

	// MRArterialSpinLabelingSequence (0018,9251) VR=SQ VM=1 MR Arterial Spin Labeling Sequence
	MRArterialSpinLabelingSequence = New(0x0018, 0x9251)

	// ASLTechniqueDescription (0018,9252) VR=LO VM=1 ASL Technique Description
	ASLTechniqueDescription = New(0x0018, 0x9252)

	// ASLSlabNumber (0018,9253) VR=US VM=1 ASL Slab Number
	ASLSlabNumber = New(0x0018, 0x9253)

	// ASLSlabThickness (0018,9254) VR=FD VM=1 ASL Slab Thickness
	ASLSlabThickness = New(0x0018, 0x9254)

	// ASLSlabOrientation (0018,9255) VR=FD VM=3 ASL Slab Orientation
	ASLSlabOrientation = New(0x0018, 0x9255)

	// ASLMidSlabPosition (0018,9256) VR=FD VM=3 ASL Mid Slab Position
	ASLMidSlabPosition = New(0x0018, 0x9256)

	// ASLContext (0018,9257) VR=CS VM=1 ASL Context
	ASLContext = New(0x0018, 0x9257)

	// ASLPulseTrainDuration (0018,9258) VR=UL VM=1 ASL Pulse Train Duration
	ASLPulseTrainDuration = New(0x0018, 0x9258)

	// ASLCrusherFlag (0018,9259) VR=CS VM=1 ASL Crusher Flag
	ASLCrusherFlag = New(0x0018, 0x9259)

	// ASLCrusherFlowLimit (0018,925A) VR=FD VM=1 ASL Crusher Flow Limit
	ASLCrusherFlowLimit = New(0x0018, 0x925A)

	// ASLCrusherDescription (0018,925B) VR=LO VM=1 ASL Crusher Description
	ASLCrusherDescription = New(0x0018, 0x925B)

	// ASLBolusCutoffFlag (0018,925C) VR=CS VM=1 ASL Bolus Cut-off Flag
	ASLBolusCutoffFlag = New(0x0018, 0x925C)

	// ASLBolusCutoffTimingSequence (0018,925D) VR=SQ VM=1 ASL Bolus Cut-off Timing Sequence
	ASLBolusCutoffTimingSequence = New(0x0018, 0x925D)

	// ASLBolusCutoffTechnique (0018,925E) VR=LO VM=1 ASL Bolus Cut-off Technique
	ASLBolusCutoffTechnique = New(0x0018, 0x925E)

	// ASLBolusCutoffDelayTime (0018,925F) VR=UL VM=1 ASL Bolus Cut-off Delay Time
	ASLBolusCutoffDelayTime = New(0x0018, 0x925F)

	// ASLSlabSequence (0018,9260) VR=SQ VM=1 ASL Slab Sequence
	ASLSlabSequence = New(0x0018, 0x9260)

	// ChemicalShiftMinimumIntegrationLimitInppm (0018,9295) VR=FD VM=1 Chemical Shift Minimum Integration Limit in ppm
	ChemicalShiftMinimumIntegrationLimitInppm = New(0x0018, 0x9295)

	// ChemicalShiftMaximumIntegrationLimitInppm (0018,9296) VR=FD VM=1 Chemical Shift Maximum Integration Limit in ppm
	ChemicalShiftMaximumIntegrationLimitInppm = New(0x0018, 0x9296)

	// WaterReferenceAcquisition (0018,9297) VR=CS VM=1 Water Reference Acquisition
	WaterReferenceAcquisition = New(0x0018, 0x9297)

	// EchoPeakPosition (0018,9298) VR=IS VM=1 Echo Peak Position
	EchoPeakPosition = New(0x0018, 0x9298)

	// CTAcquisitionTypeSequence (0018,9301) VR=SQ VM=1 CT Acquisition Type Sequence
	CTAcquisitionTypeSequence = New(0x0018, 0x9301)

	// AcquisitionType (0018,9302) VR=CS VM=1 Acquisition Type
	AcquisitionType = New(0x0018, 0x9302)

	// TubeAngle (0018,9303) VR=FD VM=1 Tube Angle
	TubeAngle = New(0x0018, 0x9303)

	// CTAcquisitionDetailsSequence (0018,9304) VR=SQ VM=1 CT Acquisition Details Sequence
	CTAcquisitionDetailsSequence = New(0x0018, 0x9304)

	// RevolutionTime (0018,9305) VR=FD VM=1 Revolution Time
	RevolutionTime = New(0x0018, 0x9305)

	// SingleCollimationWidth (0018,9306) VR=FD VM=1 Single Collimation Width
	SingleCollimationWidth = New(0x0018, 0x9306)

	// TotalCollimationWidth (0018,9307) VR=FD VM=1 Total Collimation Width
	TotalCollimationWidth = New(0x0018, 0x9307)

	// CTTableDynamicsSequence (0018,9308) VR=SQ VM=1 CT Table Dynamics Sequence
	CTTableDynamicsSequence = New(0x0018, 0x9308)

	// TableSpeed (0018,9309) VR=FD VM=1 Table Speed
	TableSpeed = New(0x0018, 0x9309)

	// TableFeedPerRotation (0018,9310) VR=FD VM=1 Table Feed per Rotation
	TableFeedPerRotation = New(0x0018, 0x9310)

	// SpiralPitchFactor (0018,9311) VR=FD VM=1 Spiral Pitch Factor
	SpiralPitchFactor = New(0x0018, 0x9311)

	// CTGeometrySequence (0018,9312) VR=SQ VM=1 CT Geometry Sequence
	CTGeometrySequence = New(0x0018, 0x9312)

	// DataCollectionCenterPatient (0018,9313) VR=FD VM=3 Data Collection Center (Patient)
	DataCollectionCenterPatient = New(0x0018, 0x9313)

	// CTReconstructionSequence (0018,9314) VR=SQ VM=1 CT Reconstruction Sequence
	CTReconstructionSequence = New(0x0018, 0x9314)

	// ReconstructionAlgorithm (0018,9315) VR=CS VM=1 Reconstruction Algorithm
	ReconstructionAlgorithm = New(0x0018, 0x9315)

	// ConvolutionKernelGroup (0018,9316) VR=CS VM=1 Convolution Kernel Group
	ConvolutionKernelGroup = New(0x0018, 0x9316)

	// ReconstructionFieldOfView (0018,9317) VR=FD VM=2 Reconstruction Field of View
	ReconstructionFieldOfView = New(0x0018, 0x9317)

	// ReconstructionTargetCenterPatient (0018,9318) VR=FD VM=3 Reconstruction Target Center (Patient)
	ReconstructionTargetCenterPatient = New(0x0018, 0x9318)

	// ReconstructionAngle (0018,9319) VR=FD VM=1 Reconstruction Angle
	ReconstructionAngle = New(0x0018, 0x9319)

	// ImageFilter (0018,9320) VR=SH VM=1 Image Filter
	ImageFilter = New(0x0018, 0x9320)

	// CTExposureSequence (0018,9321) VR=SQ VM=1 CT Exposure Sequence
	CTExposureSequence = New(0x0018, 0x9321)

	// ReconstructionPixelSpacing (0018,9322) VR=FD VM=2 Reconstruction Pixel Spacing
	ReconstructionPixelSpacing = New(0x0018, 0x9322)

	// ExposureModulationType (0018,9323) VR=CS VM=1-n Exposure Modulation Type
	ExposureModulationType = New(0x0018, 0x9323)

	// EstimatedDoseSavingRETIRED (0018,9324) VR=FD VM=1 Estimated Dose Saving (RETIRED)
	EstimatedDoseSavingRETIRED = New(0x0018, 0x9324)

	// CTXRayDetailsSequence (0018,9325) VR=SQ VM=1 CT X-Ray Details Sequence
	CTXRayDetailsSequence = New(0x0018, 0x9325)

	// CTPositionSequence (0018,9326) VR=SQ VM=1 CT Position Sequence
	CTPositionSequence = New(0x0018, 0x9326)

	// TablePosition (0018,9327) VR=FD VM=1 Table Position
	TablePosition = New(0x0018, 0x9327)

	// ExposureTimeInms (0018,9328) VR=FD VM=1 Exposure Time in ms
	ExposureTimeInms = New(0x0018, 0x9328)

	// CTImageFrameTypeSequence (0018,9329) VR=SQ VM=1 CT Image Frame Type Sequence
	CTImageFrameTypeSequence = New(0x0018, 0x9329)

	// XRayTubeCurrentInmA (0018,9330) VR=FD VM=1 X-Ray Tube Current in mA
	XRayTubeCurrentInmA = New(0x0018, 0x9330)

	// ExposureInmAs (0018,9332) VR=FD VM=1 Exposure in mAs
	ExposureInmAs = New(0x0018, 0x9332)

	// ConstantVolumeFlag (0018,9333) VR=CS VM=1 Constant Volume Flag
	ConstantVolumeFlag = New(0x0018, 0x9333)

	// FluoroscopyFlag (0018,9334) VR=CS VM=1 Fluoroscopy Flag
	FluoroscopyFlag = New(0x0018, 0x9334)

	// DistanceSourceToDataCollectionCenter (0018,9335) VR=FD VM=1 Distance Source to Data Collection Center
	DistanceSourceToDataCollectionCenter = New(0x0018, 0x9335)

	// ContrastBolusAgentNumber (0018,9337) VR=US VM=1 Contrast/Bolus Agent Number
	ContrastBolusAgentNumber = New(0x0018, 0x9337)

	// ContrastBolusIngredientCodeSequence (0018,9338) VR=SQ VM=1 Contrast/Bolus Ingredient Code Sequence
	ContrastBolusIngredientCodeSequence = New(0x0018, 0x9338)

	// ContrastAdministrationProfileSequence (0018,9340) VR=SQ VM=1 Contrast Administration Profile Sequence
	ContrastAdministrationProfileSequence = New(0x0018, 0x9340)

	// ContrastBolusUsageSequence (0018,9341) VR=SQ VM=1 Contrast/Bolus Usage Sequence
	ContrastBolusUsageSequence = New(0x0018, 0x9341)

	// ContrastBolusAgentAdministered (0018,9342) VR=CS VM=1 Contrast/Bolus Agent Administered
	ContrastBolusAgentAdministered = New(0x0018, 0x9342)

	// ContrastBolusAgentDetected (0018,9343) VR=CS VM=1 Contrast/Bolus Agent Detected
	ContrastBolusAgentDetected = New(0x0018, 0x9343)

	// ContrastBolusAgentPhase (0018,9344) VR=CS VM=1 Contrast/Bolus Agent Phase
	ContrastBolusAgentPhase = New(0x0018, 0x9344)

	// CTDIvol (0018,9345) VR=FD VM=1 CTDIvol
	CTDIvol = New(0x0018, 0x9345)

	// CTDIPhantomTypeCodeSequence (0018,9346) VR=SQ VM=1 CTDI Phantom Type Code Sequence
	CTDIPhantomTypeCodeSequence = New(0x0018, 0x9346)

	// CalciumScoringMassFactorPatient (0018,9351) VR=FL VM=1 Calcium Scoring Mass Factor Patient
	CalciumScoringMassFactorPatient = New(0x0018, 0x9351)

	// CalciumScoringMassFactorDevice (0018,9352) VR=FL VM=3 Calcium Scoring Mass Factor Device
	CalciumScoringMassFactorDevice = New(0x0018, 0x9352)

	// EnergyWeightingFactor (0018,9353) VR=FL VM=1 Energy Weighting Factor
	EnergyWeightingFactor = New(0x0018, 0x9353)

	// CTAdditionalXRaySourceSequence (0018,9360) VR=SQ VM=1 CT Additional X-Ray Source Sequence
	CTAdditionalXRaySourceSequence = New(0x0018, 0x9360)

	// MultienergyCTAcquisition (0018,9361) VR=CS VM=1 Multi-energy CT Acquisition
	MultienergyCTAcquisition = New(0x0018, 0x9361)

	// MultienergyCTAcquisitionSequence (0018,9362) VR=SQ VM=1 Multi-energy CT Acquisition Sequence
	MultienergyCTAcquisitionSequence = New(0x0018, 0x9362)

	// MultienergyCTProcessingSequence (0018,9363) VR=SQ VM=1 Multi-energy CT Processing Sequence
	MultienergyCTProcessingSequence = New(0x0018, 0x9363)

	// MultienergyCTCharacteristicsSequence (0018,9364) VR=SQ VM=1 Multi-energy CT Characteristics Sequence
	MultienergyCTCharacteristicsSequence = New(0x0018, 0x9364)

	// MultienergyCTXRaySourceSequence (0018,9365) VR=SQ VM=1 Multi-energy CT X-Ray Source Sequence
	MultienergyCTXRaySourceSequence = New(0x0018, 0x9365)

	// XRaySourceIndex (0018,9366) VR=US VM=1 X-Ray Source Index
	XRaySourceIndex = New(0x0018, 0x9366)

	// XRaySourceID (0018,9367) VR=UC VM=1 X-Ray Source ID
	XRaySourceID = New(0x0018, 0x9367)

	// MultienergySourceTechnique (0018,9368) VR=CS VM=1 Multi-energy Source Technique
	MultienergySourceTechnique = New(0x0018, 0x9368)

	// SourceStartDateTime (0018,9369) VR=DT VM=1 Source Start DateTime
	SourceStartDateTime = New(0x0018, 0x9369)

	// SourceEndDateTime (0018,936A) VR=DT VM=1 Source End DateTime
	SourceEndDateTime = New(0x0018, 0x936A)

	// SwitchingPhaseNumber (0018,936B) VR=US VM=1 Switching Phase Number
	SwitchingPhaseNumber = New(0x0018, 0x936B)

	// SwitchingPhaseNominalDuration (0018,936C) VR=DS VM=1 Switching Phase Nominal Duration
	SwitchingPhaseNominalDuration = New(0x0018, 0x936C)

	// SwitchingPhaseTransitionDuration (0018,936D) VR=DS VM=1 Switching Phase Transition Duration
	SwitchingPhaseTransitionDuration = New(0x0018, 0x936D)

	// EffectiveBinEnergy (0018,936E) VR=DS VM=1 Effective Bin Energy
	EffectiveBinEnergy = New(0x0018, 0x936E)

	// MultienergyCTXRayDetectorSequence (0018,936F) VR=SQ VM=1 Multi-energy CT X-Ray Detector Sequence
	MultienergyCTXRayDetectorSequence = New(0x0018, 0x936F)

	// XRayDetectorIndex (0018,9370) VR=US VM=1 X-Ray Detector Index
	XRayDetectorIndex = New(0x0018, 0x9370)

	// XRayDetectorID (0018,9371) VR=UC VM=1 X-Ray Detector ID
	XRayDetectorID = New(0x0018, 0x9371)

	// MultienergyDetectorType (0018,9372) VR=CS VM=1 Multi-energy Detector Type
	MultienergyDetectorType = New(0x0018, 0x9372)

	// XRayDetectorLabel (0018,9373) VR=ST VM=1 X-Ray Detector Label
	XRayDetectorLabel = New(0x0018, 0x9373)

	// NominalMaxEnergy (0018,9374) VR=DS VM=1 Nominal Max Energy
	NominalMaxEnergy = New(0x0018, 0x9374)

	// NominalMinEnergy (0018,9375) VR=DS VM=1 Nominal Min Energy
	NominalMinEnergy = New(0x0018, 0x9375)

	// ReferencedXRayDetectorIndex (0018,9376) VR=US VM=1-n Referenced X-Ray Detector Index
	ReferencedXRayDetectorIndex = New(0x0018, 0x9376)

	// ReferencedXRaySourceIndex (0018,9377) VR=US VM=1-n Referenced X-Ray Source Index
	ReferencedXRaySourceIndex = New(0x0018, 0x9377)

	// ReferencedPathIndex (0018,9378) VR=US VM=1-n Referenced Path Index
	ReferencedPathIndex = New(0x0018, 0x9378)

	// MultienergyCTPathSequence (0018,9379) VR=SQ VM=1 Multi-energy CT Path Sequence
	MultienergyCTPathSequence = New(0x0018, 0x9379)

	// MultienergyCTPathIndex (0018,937A) VR=US VM=1 Multi-energy CT Path Index
	MultienergyCTPathIndex = New(0x0018, 0x937A)

	// MultienergyAcquisitionDescription (0018,937B) VR=UT VM=1 Multi-energy Acquisition Description
	MultienergyAcquisitionDescription = New(0x0018, 0x937B)

	// MonoenergeticEnergyEquivalent (0018,937C) VR=FD VM=1 Monoenergetic Energy Equivalent
	MonoenergeticEnergyEquivalent = New(0x0018, 0x937C)

	// MaterialCodeSequence (0018,937D) VR=SQ VM=1 Material Code Sequence
	MaterialCodeSequence = New(0x0018, 0x937D)

	// DecompositionMethod (0018,937E) VR=CS VM=1 Decomposition Method
	DecompositionMethod = New(0x0018, 0x937E)

	// DecompositionDescription (0018,937F) VR=UT VM=1 Decomposition Description
	DecompositionDescription = New(0x0018, 0x937F)

	// DecompositionAlgorithmIdentificationSequence (0018,9380) VR=SQ VM=1 Decomposition Algorithm Identification Sequence
	DecompositionAlgorithmIdentificationSequence = New(0x0018, 0x9380)

	// DecompositionMaterialSequence (0018,9381) VR=SQ VM=1 Decomposition Material Sequence
	DecompositionMaterialSequence = New(0x0018, 0x9381)

	// MaterialAttenuationSequence (0018,9382) VR=SQ VM=1 Material Attenuation Sequence
	MaterialAttenuationSequence = New(0x0018, 0x9382)

	// PhotonEnergy (0018,9383) VR=DS VM=1 Photon Energy
	PhotonEnergy = New(0x0018, 0x9383)

	// XRayMassAttenuationCoefficient (0018,9384) VR=DS VM=1 X-Ray Mass Attenuation Coefficient
	XRayMassAttenuationCoefficient = New(0x0018, 0x9384)

	// ProjectionPixelCalibrationSequence (0018,9401) VR=SQ VM=1 Projection Pixel Calibration Sequence
	ProjectionPixelCalibrationSequence = New(0x0018, 0x9401)

	// DistanceSourceToIsocenter (0018,9402) VR=FL VM=1 Distance Source to Isocenter
	DistanceSourceToIsocenter = New(0x0018, 0x9402)

	// DistanceObjectToTableTop (0018,9403) VR=FL VM=1 Distance Object to Table Top
	DistanceObjectToTableTop = New(0x0018, 0x9403)

	// ObjectPixelSpacingInCenterOfBeam (0018,9404) VR=FL VM=2 Object Pixel Spacing in Center of Beam
	ObjectPixelSpacingInCenterOfBeam = New(0x0018, 0x9404)

	// PositionerPositionSequence (0018,9405) VR=SQ VM=1 Positioner Position Sequence
	PositionerPositionSequence = New(0x0018, 0x9405)

	// TablePositionSequence (0018,9406) VR=SQ VM=1 Table Position Sequence
	TablePositionSequence = New(0x0018, 0x9406)

	// CollimatorShapeSequence (0018,9407) VR=SQ VM=1 Collimator Shape Sequence
	CollimatorShapeSequence = New(0x0018, 0x9407)

	// PlanesInAcquisition (0018,9410) VR=CS VM=1 Planes in Acquisition
	PlanesInAcquisition = New(0x0018, 0x9410)

	// XAXRFFrameCharacteristicsSequence (0018,9412) VR=SQ VM=1 XA/XRF Frame Characteristics Sequence
	XAXRFFrameCharacteristicsSequence = New(0x0018, 0x9412)

	// FrameAcquisitionSequence (0018,9417) VR=SQ VM=1 Frame Acquisition Sequence
	FrameAcquisitionSequence = New(0x0018, 0x9417)

	// XRayReceptorType (0018,9420) VR=CS VM=1 X-Ray Receptor Type
	XRayReceptorType = New(0x0018, 0x9420)

	// AcquisitionProtocolName (0018,9423) VR=LO VM=1 Acquisition Protocol Name
	AcquisitionProtocolName = New(0x0018, 0x9423)

	// AcquisitionProtocolDescription (0018,9424) VR=LT VM=1 Acquisition Protocol Description
	AcquisitionProtocolDescription = New(0x0018, 0x9424)

	// ContrastBolusIngredientOpaque (0018,9425) VR=CS VM=1 Contrast/Bolus Ingredient Opaque
	ContrastBolusIngredientOpaque = New(0x0018, 0x9425)

	// DistanceReceptorPlaneToDetectorHousing (0018,9426) VR=FL VM=1 Distance Receptor Plane to Detector Housing
	DistanceReceptorPlaneToDetectorHousing = New(0x0018, 0x9426)

	// IntensifierActiveShape (0018,9427) VR=CS VM=1 Intensifier Active Shape
	IntensifierActiveShape = New(0x0018, 0x9427)

	// IntensifierActiveDimensions (0018,9428) VR=FL VM=1-2 Intensifier Active Dimension(s)
	IntensifierActiveDimensions = New(0x0018, 0x9428)

	// PhysicalDetectorSize (0018,9429) VR=FL VM=2 Physical Detector Size
	PhysicalDetectorSize = New(0x0018, 0x9429)

	// PositionOfIsocenterProjection (0018,9430) VR=FL VM=2 Position of Isocenter Projection
	PositionOfIsocenterProjection = New(0x0018, 0x9430)

	// FieldOfViewSequence (0018,9432) VR=SQ VM=1 Field of View Sequence
	FieldOfViewSequence = New(0x0018, 0x9432)

	// FieldOfViewDescription (0018,9433) VR=LO VM=1 Field of View Description
	FieldOfViewDescription = New(0x0018, 0x9433)

	// ExposureControlSensingRegionsSequence (0018,9434) VR=SQ VM=1 Exposure Control Sensing Regions Sequence
	ExposureControlSensingRegionsSequence = New(0x0018, 0x9434)

	// ExposureControlSensingRegionShape (0018,9435) VR=CS VM=1 Exposure Control Sensing Region Shape
	ExposureControlSensingRegionShape = New(0x0018, 0x9435)

	// ExposureControlSensingRegionLeftVerticalEdge (0018,9436) VR=SS VM=1 Exposure Control Sensing Region Left Vertical Edge
	ExposureControlSensingRegionLeftVerticalEdge = New(0x0018, 0x9436)

	// ExposureControlSensingRegionRightVerticalEdge (0018,9437) VR=SS VM=1 Exposure Control Sensing Region Right Vertical Edge
	ExposureControlSensingRegionRightVerticalEdge = New(0x0018, 0x9437)

	// ExposureControlSensingRegionUpperHorizontalEdge (0018,9438) VR=SS VM=1 Exposure Control Sensing Region Upper Horizontal Edge
	ExposureControlSensingRegionUpperHorizontalEdge = New(0x0018, 0x9438)

	// ExposureControlSensingRegionLowerHorizontalEdge (0018,9439) VR=SS VM=1 Exposure Control Sensing Region Lower Horizontal Edge
	ExposureControlSensingRegionLowerHorizontalEdge = New(0x0018, 0x9439)

	// CenterOfCircularExposureControlSensingRegion (0018,9440) VR=SS VM=2 Center of Circular Exposure Control Sensing Region
	CenterOfCircularExposureControlSensingRegion = New(0x0018, 0x9440)

	// RadiusOfCircularExposureControlSensingRegion (0018,9441) VR=US VM=1 Radius of Circular Exposure Control Sensing Region
	RadiusOfCircularExposureControlSensingRegion = New(0x0018, 0x9441)

	// VerticesOfThePolygonalExposureControlSensingRegion (0018,9442) VR=SS VM=2-n Vertices of the Polygonal Exposure Control Sensing Region
	VerticesOfThePolygonalExposureControlSensingRegion = New(0x0018, 0x9442)

	// ColumnAngulationPatient (0018,9447) VR=FL VM=1 Column Angulation (Patient)
	ColumnAngulationPatient = New(0x0018, 0x9447)

	// BeamAngle (0018,9449) VR=FL VM=1 Beam Angle
	BeamAngle = New(0x0018, 0x9449)

	// FrameDetectorParametersSequence (0018,9451) VR=SQ VM=1 Frame Detector Parameters Sequence
	FrameDetectorParametersSequence = New(0x0018, 0x9451)

	// CalculatedAnatomyThickness (0018,9452) VR=FL VM=1 Calculated Anatomy Thickness
	CalculatedAnatomyThickness = New(0x0018, 0x9452)

	// CalibrationSequence (0018,9455) VR=SQ VM=1 Calibration Sequence
	CalibrationSequence = New(0x0018, 0x9455)

	// ObjectThicknessSequence (0018,9456) VR=SQ VM=1 Object Thickness Sequence
	ObjectThicknessSequence = New(0x0018, 0x9456)

	// PlaneIdentification (0018,9457) VR=CS VM=1 Plane Identification
	PlaneIdentification = New(0x0018, 0x9457)

	// FieldOfViewDimensionsInFloat (0018,9461) VR=FL VM=1-2 Field of View Dimension(s) in Float
	FieldOfViewDimensionsInFloat = New(0x0018, 0x9461)

	// IsocenterReferenceSystemSequence (0018,9462) VR=SQ VM=1 Isocenter Reference System Sequence
	IsocenterReferenceSystemSequence = New(0x0018, 0x9462)

	// PositionerIsocenterPrimaryAngle (0018,9463) VR=FL VM=1 Positioner Isocenter Primary Angle
	PositionerIsocenterPrimaryAngle = New(0x0018, 0x9463)

	// PositionerIsocenterSecondaryAngle (0018,9464) VR=FL VM=1 Positioner Isocenter Secondary Angle
	PositionerIsocenterSecondaryAngle = New(0x0018, 0x9464)

	// PositionerIsocenterDetectorRotationAngle (0018,9465) VR=FL VM=1 Positioner Isocenter Detector Rotation Angle
	PositionerIsocenterDetectorRotationAngle = New(0x0018, 0x9465)

	// TableXPositionToIsocenter (0018,9466) VR=FL VM=1 Table X Position to Isocenter
	TableXPositionToIsocenter = New(0x0018, 0x9466)

	// TableYPositionToIsocenter (0018,9467) VR=FL VM=1 Table Y Position to Isocenter
	TableYPositionToIsocenter = New(0x0018, 0x9467)

	// TableZPositionToIsocenter (0018,9468) VR=FL VM=1 Table Z Position to Isocenter
	TableZPositionToIsocenter = New(0x0018, 0x9468)

	// TableHorizontalRotationAngle (0018,9469) VR=FL VM=1 Table Horizontal Rotation Angle
	TableHorizontalRotationAngle = New(0x0018, 0x9469)

	// TableHeadTiltAngle (0018,9470) VR=FL VM=1 Table Head Tilt Angle
	TableHeadTiltAngle = New(0x0018, 0x9470)

	// TableCradleTiltAngle (0018,9471) VR=FL VM=1 Table Cradle Tilt Angle
	TableCradleTiltAngle = New(0x0018, 0x9471)

	// FrameDisplayShutterSequence (0018,9472) VR=SQ VM=1 Frame Display Shutter Sequence
	FrameDisplayShutterSequence = New(0x0018, 0x9472)

	// AcquiredImageAreaDoseProduct (0018,9473) VR=FL VM=1 Acquired Image Area Dose Product
	AcquiredImageAreaDoseProduct = New(0x0018, 0x9473)

	// CArmPositionerTabletopRelationship (0018,9474) VR=CS VM=1 C-arm Positioner Tabletop Relationship
	CArmPositionerTabletopRelationship = New(0x0018, 0x9474)

	// XRayGeometrySequence (0018,9476) VR=SQ VM=1 X-Ray Geometry Sequence
	XRayGeometrySequence = New(0x0018, 0x9476)

	// IrradiationEventIdentificationSequence (0018,9477) VR=SQ VM=1 Irradiation Event Identification Sequence
	IrradiationEventIdentificationSequence = New(0x0018, 0x9477)

	// XRay3DFrameTypeSequence (0018,9504) VR=SQ VM=1 X-Ray 3D Frame Type Sequence
	XRay3DFrameTypeSequence = New(0x0018, 0x9504)

	// ContributingSourcesSequence (0018,9506) VR=SQ VM=1 Contributing Sources Sequence
	ContributingSourcesSequence = New(0x0018, 0x9506)

	// XRay3DAcquisitionSequence (0018,9507) VR=SQ VM=1 X-Ray 3D Acquisition Sequence
	XRay3DAcquisitionSequence = New(0x0018, 0x9507)

	// PrimaryPositionerScanArc (0018,9508) VR=FL VM=1 Primary Positioner Scan Arc
	PrimaryPositionerScanArc = New(0x0018, 0x9508)

	// SecondaryPositionerScanArc (0018,9509) VR=FL VM=1 Secondary Positioner Scan Arc
	SecondaryPositionerScanArc = New(0x0018, 0x9509)

	// PrimaryPositionerScanStartAngle (0018,9510) VR=FL VM=1 Primary Positioner Scan Start Angle
	PrimaryPositionerScanStartAngle = New(0x0018, 0x9510)

	// SecondaryPositionerScanStartAngle (0018,9511) VR=FL VM=1 Secondary Positioner Scan Start Angle
	SecondaryPositionerScanStartAngle = New(0x0018, 0x9511)

	// PrimaryPositionerIncrement (0018,9514) VR=FL VM=1 Primary Positioner Increment
	PrimaryPositionerIncrement = New(0x0018, 0x9514)

	// SecondaryPositionerIncrement (0018,9515) VR=FL VM=1 Secondary Positioner Increment
	SecondaryPositionerIncrement = New(0x0018, 0x9515)

	// StartAcquisitionDateTime (0018,9516) VR=DT VM=1 Start Acquisition DateTime
	StartAcquisitionDateTime = New(0x0018, 0x9516)

	// EndAcquisitionDateTime (0018,9517) VR=DT VM=1 End Acquisition DateTime
	EndAcquisitionDateTime = New(0x0018, 0x9517)

	// PrimaryPositionerIncrementSign (0018,9518) VR=SS VM=1 Primary Positioner Increment Sign
	PrimaryPositionerIncrementSign = New(0x0018, 0x9518)

	// SecondaryPositionerIncrementSign (0018,9519) VR=SS VM=1 Secondary Positioner Increment Sign
	SecondaryPositionerIncrementSign = New(0x0018, 0x9519)

	// ApplicationName (0018,9524) VR=LO VM=1 Application Name
	ApplicationName = New(0x0018, 0x9524)

	// ApplicationVersion (0018,9525) VR=LO VM=1 Application Version
	ApplicationVersion = New(0x0018, 0x9525)

	// ApplicationManufacturer (0018,9526) VR=LO VM=1 Application Manufacturer
	ApplicationManufacturer = New(0x0018, 0x9526)

	// AlgorithmType (0018,9527) VR=CS VM=1 Algorithm Type
	AlgorithmType = New(0x0018, 0x9527)

	// AlgorithmDescription (0018,9528) VR=LO VM=1 Algorithm Description
	AlgorithmDescription = New(0x0018, 0x9528)

	// XRay3DReconstructionSequence (0018,9530) VR=SQ VM=1 X-Ray 3D Reconstruction Sequence
	XRay3DReconstructionSequence = New(0x0018, 0x9530)

	// ReconstructionDescription (0018,9531) VR=LO VM=1 Reconstruction Description
	ReconstructionDescription = New(0x0018, 0x9531)

	// PerProjectionAcquisitionSequence (0018,9538) VR=SQ VM=1 Per Projection Acquisition Sequence
	PerProjectionAcquisitionSequence = New(0x0018, 0x9538)

	// DetectorPositionSequence (0018,9541) VR=SQ VM=1 Detector Position Sequence
	DetectorPositionSequence = New(0x0018, 0x9541)

	// XRayAcquisitionDoseSequence (0018,9542) VR=SQ VM=1 X-Ray Acquisition Dose Sequence
	XRayAcquisitionDoseSequence = New(0x0018, 0x9542)

	// XRaySourceIsocenterPrimaryAngle (0018,9543) VR=FD VM=1 X-Ray Source Isocenter Primary Angle
	XRaySourceIsocenterPrimaryAngle = New(0x0018, 0x9543)

	// XRaySourceIsocenterSecondaryAngle (0018,9544) VR=FD VM=1 X-Ray Source Isocenter Secondary Angle
	XRaySourceIsocenterSecondaryAngle = New(0x0018, 0x9544)

	// BreastSupportIsocenterPrimaryAngle (0018,9545) VR=FD VM=1 Breast Support Isocenter Primary Angle
	BreastSupportIsocenterPrimaryAngle = New(0x0018, 0x9545)

	// BreastSupportIsocenterSecondaryAngle (0018,9546) VR=FD VM=1 Breast Support Isocenter Secondary Angle
	BreastSupportIsocenterSecondaryAngle = New(0x0018, 0x9546)

	// BreastSupportXPositionToIsocenter (0018,9547) VR=FD VM=1 Breast Support X Position to Isocenter
	BreastSupportXPositionToIsocenter = New(0x0018, 0x9547)

	// BreastSupportYPositionToIsocenter (0018,9548) VR=FD VM=1 Breast Support Y Position to Isocenter
	BreastSupportYPositionToIsocenter = New(0x0018, 0x9548)

	// BreastSupportZPositionToIsocenter (0018,9549) VR=FD VM=1 Breast Support Z Position to Isocenter
	BreastSupportZPositionToIsocenter = New(0x0018, 0x9549)

	// DetectorIsocenterPrimaryAngle (0018,9550) VR=FD VM=1 Detector Isocenter Primary Angle
	DetectorIsocenterPrimaryAngle = New(0x0018, 0x9550)

	// DetectorIsocenterSecondaryAngle (0018,9551) VR=FD VM=1 Detector Isocenter Secondary Angle
	DetectorIsocenterSecondaryAngle = New(0x0018, 0x9551)

	// DetectorXPositionToIsocenter (0018,9552) VR=FD VM=1 Detector X Position to Isocenter
	DetectorXPositionToIsocenter = New(0x0018, 0x9552)

	// DetectorYPositionToIsocenter (0018,9553) VR=FD VM=1 Detector Y Position to Isocenter
	DetectorYPositionToIsocenter = New(0x0018, 0x9553)

	// DetectorZPositionToIsocenter (0018,9554) VR=FD VM=1 Detector Z Position to Isocenter
	DetectorZPositionToIsocenter = New(0x0018, 0x9554)

	// XRayGridSequence (0018,9555) VR=SQ VM=1 X-Ray Grid Sequence
	XRayGridSequence = New(0x0018, 0x9555)

	// XRayFilterSequence (0018,9556) VR=SQ VM=1 X-Ray Filter Sequence
	XRayFilterSequence = New(0x0018, 0x9556)

	// DetectorActiveAreaTLHCPosition (0018,9557) VR=FD VM=3 Detector Active Area TLHC Position
	DetectorActiveAreaTLHCPosition = New(0x0018, 0x9557)

	// DetectorActiveAreaOrientation (0018,9558) VR=FD VM=6 Detector Active Area Orientation
	DetectorActiveAreaOrientation = New(0x0018, 0x9558)

	// PositionerPrimaryAngleDirection (0018,9559) VR=CS VM=1 Positioner Primary Angle Direction
	PositionerPrimaryAngleDirection = New(0x0018, 0x9559)

	// DiffusionBMatrixSequence (0018,9601) VR=SQ VM=1 Diffusion b-matrix Sequence
	DiffusionBMatrixSequence = New(0x0018, 0x9601)

	// DiffusionBValueXX (0018,9602) VR=FD VM=1 Diffusion b-value XX
	DiffusionBValueXX = New(0x0018, 0x9602)

	// DiffusionBValueXY (0018,9603) VR=FD VM=1 Diffusion b-value XY
	DiffusionBValueXY = New(0x0018, 0x9603)

	// DiffusionBValueXZ (0018,9604) VR=FD VM=1 Diffusion b-value XZ
	DiffusionBValueXZ = New(0x0018, 0x9604)

	// DiffusionBValueYY (0018,9605) VR=FD VM=1 Diffusion b-value YY
	DiffusionBValueYY = New(0x0018, 0x9605)

	// DiffusionBValueYZ (0018,9606) VR=FD VM=1 Diffusion b-value YZ
	DiffusionBValueYZ = New(0x0018, 0x9606)

	// DiffusionBValueZZ (0018,9607) VR=FD VM=1 Diffusion b-value ZZ
	DiffusionBValueZZ = New(0x0018, 0x9607)

	// FunctionalMRSequence (0018,9621) VR=SQ VM=1 Functional MR Sequence
	FunctionalMRSequence = New(0x0018, 0x9621)

	// FunctionalSettlingPhaseFramesPresent (0018,9622) VR=CS VM=1 Functional Settling Phase Frames Present
	FunctionalSettlingPhaseFramesPresent = New(0x0018, 0x9622)

	// FunctionalSyncPulse (0018,9623) VR=DT VM=1 Functional Sync Pulse
	FunctionalSyncPulse = New(0x0018, 0x9623)

	// SettlingPhaseFrame (0018,9624) VR=CS VM=1 Settling Phase Frame
	SettlingPhaseFrame = New(0x0018, 0x9624)

	// DecayCorrectionDateTime (0018,9701) VR=DT VM=1 Decay Correction DateTime
	DecayCorrectionDateTime = New(0x0018, 0x9701)

	// StartDensityThreshold (0018,9715) VR=FD VM=1 Start Density Threshold
	StartDensityThreshold = New(0x0018, 0x9715)

	// StartRelativeDensityDifferenceThreshold (0018,9716) VR=FD VM=1 Start Relative Density Difference Threshold
	StartRelativeDensityDifferenceThreshold = New(0x0018, 0x9716)

	// StartCardiacTriggerCountThreshold (0018,9717) VR=FD VM=1 Start Cardiac Trigger Count Threshold
	StartCardiacTriggerCountThreshold = New(0x0018, 0x9717)

	// StartRespiratoryTriggerCountThreshold (0018,9718) VR=FD VM=1 Start Respiratory Trigger Count Threshold
	StartRespiratoryTriggerCountThreshold = New(0x0018, 0x9718)

	// TerminationCountsThreshold (0018,9719) VR=FD VM=1 Termination Counts Threshold
	TerminationCountsThreshold = New(0x0018, 0x9719)

	// TerminationDensityThreshold (0018,9720) VR=FD VM=1 Termination Density Threshold
	TerminationDensityThreshold = New(0x0018, 0x9720)

	// TerminationRelativeDensityThreshold (0018,9721) VR=FD VM=1 Termination Relative Density Threshold
	TerminationRelativeDensityThreshold = New(0x0018, 0x9721)

	// TerminationTimeThreshold (0018,9722) VR=FD VM=1 Termination Time Threshold
	TerminationTimeThreshold = New(0x0018, 0x9722)

	// TerminationCardiacTriggerCountThreshold (0018,9723) VR=FD VM=1 Termination Cardiac Trigger Count Threshold
	TerminationCardiacTriggerCountThreshold = New(0x0018, 0x9723)

	// TerminationRespiratoryTriggerCountThreshold (0018,9724) VR=FD VM=1 Termination Respiratory Trigger Count Threshold
	TerminationRespiratoryTriggerCountThreshold = New(0x0018, 0x9724)

	// DetectorGeometry (0018,9725) VR=CS VM=1 Detector Geometry
	DetectorGeometry = New(0x0018, 0x9725)

	// TransverseDetectorSeparation (0018,9726) VR=FD VM=1 Transverse Detector Separation
	TransverseDetectorSeparation = New(0x0018, 0x9726)

	// AxialDetectorDimension (0018,9727) VR=FD VM=1 Axial Detector Dimension
	AxialDetectorDimension = New(0x0018, 0x9727)

	// RadiopharmaceuticalAgentNumber (0018,9729) VR=US VM=1 Radiopharmaceutical Agent Number
	RadiopharmaceuticalAgentNumber = New(0x0018, 0x9729)

	// PETFrameAcquisitionSequence (0018,9732) VR=SQ VM=1 PET Frame Acquisition Sequence
	PETFrameAcquisitionSequence = New(0x0018, 0x9732)

	// PETDetectorMotionDetailsSequence (0018,9733) VR=SQ VM=1 PET Detector Motion Details Sequence
	PETDetectorMotionDetailsSequence = New(0x0018, 0x9733)

	// PETTableDynamicsSequence (0018,9734) VR=SQ VM=1 PET Table Dynamics Sequence
	PETTableDynamicsSequence = New(0x0018, 0x9734)

	// PETPositionSequence (0018,9735) VR=SQ VM=1 PET Position Sequence
	PETPositionSequence = New(0x0018, 0x9735)

	// PETFrameCorrectionFactorsSequence (0018,9736) VR=SQ VM=1 PET Frame Correction Factors Sequence
	PETFrameCorrectionFactorsSequence = New(0x0018, 0x9736)

	// RadiopharmaceuticalUsageSequence (0018,9737) VR=SQ VM=1 Radiopharmaceutical Usage Sequence
	RadiopharmaceuticalUsageSequence = New(0x0018, 0x9737)

	// AttenuationCorrectionSource (0018,9738) VR=CS VM=1 Attenuation Correction Source
	AttenuationCorrectionSource = New(0x0018, 0x9738)

	// NumberOfIterations (0018,9739) VR=US VM=1 Number of Iterations
	NumberOfIterations = New(0x0018, 0x9739)

	// NumberOfSubsets (0018,9740) VR=US VM=1 Number of Subsets
	NumberOfSubsets = New(0x0018, 0x9740)

	// PETReconstructionSequence (0018,9749) VR=SQ VM=1 PET Reconstruction Sequence
	PETReconstructionSequence = New(0x0018, 0x9749)

	// PETFrameTypeSequence (0018,9751) VR=SQ VM=1 PET Frame Type Sequence
	PETFrameTypeSequence = New(0x0018, 0x9751)

	// TimeOfFlightInformationUsed (0018,9755) VR=CS VM=1 Time of Flight Information Used
	TimeOfFlightInformationUsed = New(0x0018, 0x9755)

	// ReconstructionType (0018,9756) VR=CS VM=1 Reconstruction Type
	ReconstructionType = New(0x0018, 0x9756)

	// DecayCorrected (0018,9758) VR=CS VM=1 Decay Corrected
	DecayCorrected = New(0x0018, 0x9758)

	// AttenuationCorrected (0018,9759) VR=CS VM=1 Attenuation Corrected
	AttenuationCorrected = New(0x0018, 0x9759)

	// ScatterCorrected (0018,9760) VR=CS VM=1 Scatter Corrected
	ScatterCorrected = New(0x0018, 0x9760)

	// DeadTimeCorrected (0018,9761) VR=CS VM=1 Dead Time Corrected
	DeadTimeCorrected = New(0x0018, 0x9761)

	// GantryMotionCorrected (0018,9762) VR=CS VM=1 Gantry Motion Corrected
	GantryMotionCorrected = New(0x0018, 0x9762)

	// PatientMotionCorrected (0018,9763) VR=CS VM=1 Patient Motion Corrected
	PatientMotionCorrected = New(0x0018, 0x9763)

	// CountLossNormalizationCorrected (0018,9764) VR=CS VM=1 Count Loss Normalization Corrected
	CountLossNormalizationCorrected = New(0x0018, 0x9764)

	// RandomsCorrected (0018,9765) VR=CS VM=1 Randoms Corrected
	RandomsCorrected = New(0x0018, 0x9765)

	// NonUniformRadialSamplingCorrected (0018,9766) VR=CS VM=1 Non-uniform Radial Sampling Corrected
	NonUniformRadialSamplingCorrected = New(0x0018, 0x9766)

	// SensitivityCalibrated (0018,9767) VR=CS VM=1 Sensitivity Calibrated
	SensitivityCalibrated = New(0x0018, 0x9767)

	// DetectorNormalizationCorrection (0018,9768) VR=CS VM=1 Detector Normalization Correction
	DetectorNormalizationCorrection = New(0x0018, 0x9768)

	// IterativeReconstructionMethod (0018,9769) VR=CS VM=1 Iterative Reconstruction Method
	IterativeReconstructionMethod = New(0x0018, 0x9769)

	// AttenuationCorrectionTemporalRelationship (0018,9770) VR=CS VM=1 Attenuation Correction Temporal Relationship
	AttenuationCorrectionTemporalRelationship = New(0x0018, 0x9770)

	// PatientPhysiologicalStateSequence (0018,9771) VR=SQ VM=1 Patient Physiological State Sequence
	PatientPhysiologicalStateSequence = New(0x0018, 0x9771)

	// PatientPhysiologicalStateCodeSequence (0018,9772) VR=SQ VM=1 Patient Physiological State Code Sequence
	PatientPhysiologicalStateCodeSequence = New(0x0018, 0x9772)

	// DepthsOfFocus (0018,9801) VR=FD VM=1-n Depth(s) of Focus
	DepthsOfFocus = New(0x0018, 0x9801)

	// ExcludedIntervalsSequence (0018,9803) VR=SQ VM=1 Excluded Intervals Sequence
	ExcludedIntervalsSequence = New(0x0018, 0x9803)

	// ExclusionStartDateTime (0018,9804) VR=DT VM=1 Exclusion Start DateTime
	ExclusionStartDateTime = New(0x0018, 0x9804)

	// ExclusionDuration (0018,9805) VR=FD VM=1 Exclusion Duration
	ExclusionDuration = New(0x0018, 0x9805)

	// USImageDescriptionSequence (0018,9806) VR=SQ VM=1 US Image Description Sequence
	USImageDescriptionSequence = New(0x0018, 0x9806)

	// ImageDataTypeSequence (0018,9807) VR=SQ VM=1 Image Data Type Sequence
	ImageDataTypeSequence = New(0x0018, 0x9807)

	// DataType (0018,9808) VR=CS VM=1 Data Type
	DataType = New(0x0018, 0x9808)

	// TransducerScanPatternCodeSequence (0018,9809) VR=SQ VM=1 Transducer Scan Pattern Code Sequence
	TransducerScanPatternCodeSequence = New(0x0018, 0x9809)

	// AliasedDataType (0018,980B) VR=CS VM=1 Aliased Data Type
	AliasedDataType = New(0x0018, 0x980B)

	// PositionMeasuringDeviceUsed (0018,980C) VR=CS VM=1 Position Measuring Device Used
	PositionMeasuringDeviceUsed = New(0x0018, 0x980C)

	// TransducerGeometryCodeSequence (0018,980D) VR=SQ VM=1 Transducer Geometry Code Sequence
	TransducerGeometryCodeSequence = New(0x0018, 0x980D)

	// TransducerBeamSteeringCodeSequence (0018,980E) VR=SQ VM=1 Transducer Beam Steering Code Sequence
	TransducerBeamSteeringCodeSequence = New(0x0018, 0x980E)

	// TransducerApplicationCodeSequence (0018,980F) VR=SQ VM=1 Transducer Application Code Sequence
	TransducerApplicationCodeSequence = New(0x0018, 0x980F)

	ZeroVelocityPixelValue = New(0x0018, 0x9810)

	// PhotoacousticExcitationCharacteristicsSequence (0018,9821) VR=SQ VM=1 Photoacoustic Excitation Characteristics Sequence
	PhotoacousticExcitationCharacteristicsSequence = New(0x0018, 0x9821)

	// ExcitationSpectralWidth (0018,9822) VR=FD VM=1 Excitation Spectral Width
	ExcitationSpectralWidth = New(0x0018, 0x9822)

	// ExcitationEnergy (0018,9823) VR=FD VM=1 Excitation Energy
	ExcitationEnergy = New(0x0018, 0x9823)

	// ExcitationPulseDuration (0018,9824) VR=FD VM=1 Excitation Pulse Duration
	ExcitationPulseDuration = New(0x0018, 0x9824)

	// ExcitationWavelengthSequence (0018,9825) VR=SQ VM=1 Excitation Wavelength Sequence
	ExcitationWavelengthSequence = New(0x0018, 0x9825)

	// ExcitationWavelength (0018,9826) VR=FD VM=1 Excitation Wavelength
	ExcitationWavelength = New(0x0018, 0x9826)

	// IlluminationTranslationFlag (0018,9828) VR=CS VM=1 Illumination Translation Flag
	IlluminationTranslationFlag = New(0x0018, 0x9828)

	// AcousticCouplingMediumFlag (0018,9829) VR=CS VM=1 Acoustic Coupling Medium Flag
	AcousticCouplingMediumFlag = New(0x0018, 0x9829)

	// AcousticCouplingMediumCodeSequence (0018,982A) VR=SQ VM=1 Acoustic Coupling Medium Code Sequence
	AcousticCouplingMediumCodeSequence = New(0x0018, 0x982A)

	// AcousticCouplingMediumTemperature (0018,982B) VR=FD VM=1 Acoustic Coupling Medium Temperature
	AcousticCouplingMediumTemperature = New(0x0018, 0x982B)

	// TransducerResponseSequence (0018,982C) VR=SQ VM=1 Transducer Response Sequence
	TransducerResponseSequence = New(0x0018, 0x982C)

	// CenterFrequency (0018,982D) VR=FD VM=1 Center Frequency
	CenterFrequency = New(0x0018, 0x982D)

	// FractionalBandwidth (0018,982E) VR=FD VM=1 Fractional Bandwidth
	FractionalBandwidth = New(0x0018, 0x982E)

	// LowerCutoffFrequency (0018,982F) VR=FD VM=1 Lower Cutoff Frequency
	LowerCutoffFrequency = New(0x0018, 0x982F)

	// UpperCutoffFrequency (0018,9830) VR=FD VM=1 Upper Cutoff Frequency
	UpperCutoffFrequency = New(0x0018, 0x9830)

	// TransducerTechnologySequence (0018,9831) VR=SQ VM=1 Transducer Technology Sequence
	TransducerTechnologySequence = New(0x0018, 0x9831)

	// SoundSpeedCorrectionMechanismCodeSequence (0018,9832) VR=SQ VM=1 Sound Speed Correction Mechanism Code Sequence
	SoundSpeedCorrectionMechanismCodeSequence = New(0x0018, 0x9832)

	// ObjectSoundSpeed (0018,9833) VR=FD VM=1 Object Sound Speed
	ObjectSoundSpeed = New(0x0018, 0x9833)

	// AcousticCouplingMediumSoundSpeed (0018,9834) VR=FD VM=1 Acoustic Coupling Medium Sound Speed
	AcousticCouplingMediumSoundSpeed = New(0x0018, 0x9834)

	// PhotoacousticImageFrameTypeSequence (0018,9835) VR=SQ VM=1 Photoacoustic Image Frame Type Sequence
	PhotoacousticImageFrameTypeSequence = New(0x0018, 0x9835)

	// ImageDataTypeCodeSequence (0018,9836) VR=SQ VM=1 Image Data Type Code Sequence
	ImageDataTypeCodeSequence = New(0x0018, 0x9836)

	// ReferenceLocationLabel (0018,9900) VR=LO VM=1 Reference Location Label
	ReferenceLocationLabel = New(0x0018, 0x9900)

	// ReferenceLocationDescription (0018,9901) VR=UT VM=1 Reference Location Description
	ReferenceLocationDescription = New(0x0018, 0x9901)

	// ReferenceBasisCodeSequence (0018,9902) VR=SQ VM=1 Reference Basis Code Sequence
	ReferenceBasisCodeSequence = New(0x0018, 0x9902)

	// ReferenceGeometryCodeSequence (0018,9903) VR=SQ VM=1 Reference Geometry Code Sequence
	ReferenceGeometryCodeSequence = New(0x0018, 0x9903)

	// OffsetDistance (0018,9904) VR=DS VM=1 Offset Distance
	OffsetDistance = New(0x0018, 0x9904)

	// OffsetDirection (0018,9905) VR=CS VM=1 Offset Direction
	OffsetDirection = New(0x0018, 0x9905)

	// PotentialScheduledProtocolCodeSequence (0018,9906) VR=SQ VM=1 Potential Scheduled Protocol Code Sequence
	PotentialScheduledProtocolCodeSequence = New(0x0018, 0x9906)

	// PotentialRequestedProcedureCodeSequence (0018,9907) VR=SQ VM=1 Potential Requested Procedure Code Sequence
	PotentialRequestedProcedureCodeSequence = New(0x0018, 0x9907)

	// PotentialReasonsForProcedure (0018,9908) VR=UC VM=1-n Potential Reasons for Procedure
	PotentialReasonsForProcedure = New(0x0018, 0x9908)

	// PotentialReasonsForProcedureCodeSequence (0018,9909) VR=SQ VM=1 Potential Reasons for Procedure Code Sequence
	PotentialReasonsForProcedureCodeSequence = New(0x0018, 0x9909)

	// PotentialDiagnosticTasks (0018,990A) VR=UC VM=1-n Potential Diagnostic Tasks
	PotentialDiagnosticTasks = New(0x0018, 0x990A)

	// ContraindicationsCodeSequence (0018,990B) VR=SQ VM=1 Contraindications Code Sequence
	ContraindicationsCodeSequence = New(0x0018, 0x990B)

	// ReferencedDefinedProtocolSequence (0018,990C) VR=SQ VM=1 Referenced Defined Protocol Sequence
	ReferencedDefinedProtocolSequence = New(0x0018, 0x990C)

	// ReferencedPerformedProtocolSequence (0018,990D) VR=SQ VM=1 Referenced Performed Protocol Sequence
	ReferencedPerformedProtocolSequence = New(0x0018, 0x990D)

	// PredecessorProtocolSequence (0018,990E) VR=SQ VM=1 Predecessor Protocol Sequence
	PredecessorProtocolSequence = New(0x0018, 0x990E)

	// ProtocolPlanningInformation (0018,990F) VR=UT VM=1 Protocol Planning Information
	ProtocolPlanningInformation = New(0x0018, 0x990F)

	// ProtocolDesignRationale (0018,9910) VR=UT VM=1 Protocol Design Rationale
	ProtocolDesignRationale = New(0x0018, 0x9910)

	// PatientSpecificationSequence (0018,9911) VR=SQ VM=1 Patient Specification Sequence
	PatientSpecificationSequence = New(0x0018, 0x9911)

	// ModelSpecificationSequence (0018,9912) VR=SQ VM=1 Model Specification Sequence
	ModelSpecificationSequence = New(0x0018, 0x9912)

	// ParametersSpecificationSequence (0018,9913) VR=SQ VM=1 Parameters Specification Sequence
	ParametersSpecificationSequence = New(0x0018, 0x9913)

	// InstructionSequence (0018,9914) VR=SQ VM=1 Instruction Sequence
	InstructionSequence = New(0x0018, 0x9914)

	// InstructionIndex (0018,9915) VR=US VM=1 Instruction Index
	InstructionIndex = New(0x0018, 0x9915)

	// InstructionText (0018,9916) VR=LO VM=1 Instruction Text
	InstructionText = New(0x0018, 0x9916)

	// InstructionDescription (0018,9917) VR=UT VM=1 Instruction Description
	InstructionDescription = New(0x0018, 0x9917)

	// InstructionPerformedFlag (0018,9918) VR=CS VM=1 Instruction Performed Flag
	InstructionPerformedFlag = New(0x0018, 0x9918)

	// InstructionPerformedDateTime (0018,9919) VR=DT VM=1 Instruction Performed DateTime
	InstructionPerformedDateTime = New(0x0018, 0x9919)

	// InstructionPerformanceComment (0018,991A) VR=UT VM=1 Instruction Performance Comment
	InstructionPerformanceComment = New(0x0018, 0x991A)

	// PatientPositioningInstructionSequence (0018,991B) VR=SQ VM=1 Patient Positioning Instruction Sequence
	PatientPositioningInstructionSequence = New(0x0018, 0x991B)

	// PositioningMethodCodeSequence (0018,991C) VR=SQ VM=1 Positioning Method Code Sequence
	PositioningMethodCodeSequence = New(0x0018, 0x991C)

	// PositioningLandmarkSequence (0018,991D) VR=SQ VM=1 Positioning Landmark Sequence
	PositioningLandmarkSequence = New(0x0018, 0x991D)

	// TargetFrameOfReferenceUID (0018,991E) VR=UI VM=1 Target Frame of Reference UID
	TargetFrameOfReferenceUID = New(0x0018, 0x991E)

	// AcquisitionProtocolElementSpecificationSequence (0018,991F) VR=SQ VM=1 Acquisition Protocol Element Specification Sequence
	AcquisitionProtocolElementSpecificationSequence = New(0x0018, 0x991F)

	// AcquisitionProtocolElementSequence (0018,9920) VR=SQ VM=1 Acquisition Protocol Element Sequence
	AcquisitionProtocolElementSequence = New(0x0018, 0x9920)

	// ProtocolElementNumber (0018,9921) VR=US VM=1 Protocol Element Number
	ProtocolElementNumber = New(0x0018, 0x9921)

	// ProtocolElementName (0018,9922) VR=LO VM=1 Protocol Element Name
	ProtocolElementName = New(0x0018, 0x9922)

	// ProtocolElementCharacteristicsSummary (0018,9923) VR=UT VM=1 Protocol Element Characteristics Summary
	ProtocolElementCharacteristicsSummary = New(0x0018, 0x9923)

	// ProtocolElementPurpose (0018,9924) VR=UT VM=1 Protocol Element Purpose
	ProtocolElementPurpose = New(0x0018, 0x9924)

	// AcquisitionMotion (0018,9930) VR=CS VM=1 Acquisition Motion
	AcquisitionMotion = New(0x0018, 0x9930)

	// AcquisitionStartLocationSequence (0018,9931) VR=SQ VM=1 Acquisition Start Location Sequence
	AcquisitionStartLocationSequence = New(0x0018, 0x9931)

	// AcquisitionEndLocationSequence (0018,9932) VR=SQ VM=1 Acquisition End Location Sequence
	AcquisitionEndLocationSequence = New(0x0018, 0x9932)

	// ReconstructionProtocolElementSpecificationSequence (0018,9933) VR=SQ VM=1 Reconstruction Protocol Element Specification Sequence
	ReconstructionProtocolElementSpecificationSequence = New(0x0018, 0x9933)

	// ReconstructionProtocolElementSequence (0018,9934) VR=SQ VM=1 Reconstruction Protocol Element Sequence
	ReconstructionProtocolElementSequence = New(0x0018, 0x9934)

	// StorageProtocolElementSpecificationSequence (0018,9935) VR=SQ VM=1 Storage Protocol Element Specification Sequence
	StorageProtocolElementSpecificationSequence = New(0x0018, 0x9935)

	// StorageProtocolElementSequence (0018,9936) VR=SQ VM=1 Storage Protocol Element Sequence
	StorageProtocolElementSequence = New(0x0018, 0x9936)

	// RequestedSeriesDescription (0018,9937) VR=LO VM=1 Requested Series Description
	RequestedSeriesDescription = New(0x0018, 0x9937)

	// SourceAcquisitionProtocolElementNumber (0018,9938) VR=US VM=1-n Source Acquisition Protocol Element Number
	SourceAcquisitionProtocolElementNumber = New(0x0018, 0x9938)

	// SourceAcquisitionBeamNumber (0018,9939) VR=US VM=1-n Source Acquisition Beam Number
	SourceAcquisitionBeamNumber = New(0x0018, 0x9939)

	// SourceReconstructionProtocolElementNumber (0018,993A) VR=US VM=1-n Source Reconstruction Protocol Element Number
	SourceReconstructionProtocolElementNumber = New(0x0018, 0x993A)

	// ReconstructionStartLocationSequence (0018,993B) VR=SQ VM=1 Reconstruction Start Location Sequence
	ReconstructionStartLocationSequence = New(0x0018, 0x993B)

	// ReconstructionEndLocationSequence (0018,993C) VR=SQ VM=1 Reconstruction End Location Sequence
	ReconstructionEndLocationSequence = New(0x0018, 0x993C)

	// ReconstructionAlgorithmSequence (0018,993D) VR=SQ VM=1 Reconstruction Algorithm Sequence
	ReconstructionAlgorithmSequence = New(0x0018, 0x993D)

	// ReconstructionTargetCenterLocationSequence (0018,993E) VR=SQ VM=1 Reconstruction Target Center Location Sequence
	ReconstructionTargetCenterLocationSequence = New(0x0018, 0x993E)

	// ImageFilterDescription (0018,9941) VR=UT VM=1 Image Filter Description
	ImageFilterDescription = New(0x0018, 0x9941)

	// CTDIvolNotificationTrigger (0018,9942) VR=FD VM=1 CTDIvol Notification Trigger
	CTDIvolNotificationTrigger = New(0x0018, 0x9942)

	// DLPNotificationTrigger (0018,9943) VR=FD VM=1 DLP Notification Trigger
	DLPNotificationTrigger = New(0x0018, 0x9943)

	// AutoKVPSelectionType (0018,9944) VR=CS VM=1 Auto KVP Selection Type
	AutoKVPSelectionType = New(0x0018, 0x9944)

	// AutoKVPUpperBound (0018,9945) VR=FD VM=1 Auto KVP Upper Bound
	AutoKVPUpperBound = New(0x0018, 0x9945)

	// AutoKVPLowerBound (0018,9946) VR=FD VM=1 Auto KVP Lower Bound
	AutoKVPLowerBound = New(0x0018, 0x9946)

	// ProtocolDefinedPatientPosition (0018,9947) VR=CS VM=1 Protocol Defined Patient Position
	ProtocolDefinedPatientPosition = New(0x0018, 0x9947)

	// ContributingEquipmentSequence (0018,A001) VR=SQ VM=1 Contributing Equipment Sequence
	ContributingEquipmentSequence = New(0x0018, 0xA001)

	// ContributionDateTime (0018,A002) VR=DT VM=1 Contribution DateTime
	ContributionDateTime = New(0x0018, 0xA002)

	// ContributionDescription (0018,A003) VR=ST VM=1 Contribution Description
	ContributionDescription = New(0x0018, 0xA003)

	// StudyInstanceUID (0020,000D) VR=UI VM=1 Study Instance UID
	StudyInstanceUID = New(0x0020, 0x000D)

	// SeriesInstanceUID (0020,000E) VR=UI VM=1 Series Instance UID
	SeriesInstanceUID = New(0x0020, 0x000E)

	// StudyID (0020,0010) VR=SH VM=1 Study ID
	StudyID = New(0x0020, 0x0010)

	// SeriesNumber (0020,0011) VR=IS VM=1 Series Number
	SeriesNumber = New(0x0020, 0x0011)

	// AcquisitionNumber (0020,0012) VR=IS VM=1 Acquisition Number
	AcquisitionNumber = New(0x0020, 0x0012)

	// InstanceNumber (0020,0013) VR=IS VM=1 Instance Number
	InstanceNumber = New(0x0020, 0x0013)

	// IsotopeNumberRETIRED (0020,0014) VR=IS VM=1 Isotope Number (RETIRED)
	IsotopeNumberRETIRED = New(0x0020, 0x0014)

	// PhaseNumberRETIRED (0020,0015) VR=IS VM=1 Phase Number (RETIRED)
	PhaseNumberRETIRED = New(0x0020, 0x0015)

	// IntervalNumberRETIRED (0020,0016) VR=IS VM=1 Interval Number (RETIRED)
	IntervalNumberRETIRED = New(0x0020, 0x0016)

	// TimeSlotNumberRETIRED (0020,0017) VR=IS VM=1 Time Slot Number (RETIRED)
	TimeSlotNumberRETIRED = New(0x0020, 0x0017)

	// AngleNumberRETIRED (0020,0018) VR=IS VM=1 Angle Number (RETIRED)
	AngleNumberRETIRED = New(0x0020, 0x0018)

	// ItemNumber (0020,0019) VR=IS VM=1 Item Number
	ItemNumber = New(0x0020, 0x0019)

	// PatientOrientation (0020,0020) VR=CS VM=2 Patient Orientation
	PatientOrientation = New(0x0020, 0x0020)

	// OverlayNumberRETIRED (0020,0022) VR=IS VM=1 Overlay Number (RETIRED)
	OverlayNumberRETIRED = New(0x0020, 0x0022)

	// CurveNumberRETIRED (0020,0024) VR=IS VM=1 Curve Number (RETIRED)
	CurveNumberRETIRED = New(0x0020, 0x0024)

	// LUTNumberRETIRED (0020,0026) VR=IS VM=1 LUT Number (RETIRED)
	LUTNumberRETIRED = New(0x0020, 0x0026)

	// PyramidLabel (0020,0027) VR=LO VM=1 Pyramid Label
	PyramidLabel = New(0x0020, 0x0027)

	// ImagePositionRETIRED (0020,0030) VR=DS VM=3 Image Position (RETIRED)
	ImagePositionRETIRED = New(0x0020, 0x0030)

	// ImagePositionPatient (0020,0032) VR=DS VM=3 Image Position (Patient)
	ImagePositionPatient = New(0x0020, 0x0032)

	// ImageOrientationRETIRED (0020,0035) VR=DS VM=6 Image Orientation (RETIRED)
	ImageOrientationRETIRED = New(0x0020, 0x0035)

	// ImageOrientationPatient (0020,0037) VR=DS VM=6 Image Orientation (Patient)
	ImageOrientationPatient = New(0x0020, 0x0037)

	// LocationRETIRED (0020,0050) VR=DS VM=1 Location (RETIRED)
	LocationRETIRED = New(0x0020, 0x0050)

	// FrameOfReferenceUID (0020,0052) VR=UI VM=1 Frame of Reference UID
	FrameOfReferenceUID = New(0x0020, 0x0052)

	// Laterality (0020,0060) VR=CS VM=1 Laterality
	Laterality = New(0x0020, 0x0060)

	// ImageLaterality (0020,0062) VR=CS VM=1 Image Laterality
	ImageLaterality = New(0x0020, 0x0062)

	// ImageGeometryTypeRETIRED (0020,0070) VR=LO VM=1 Image Geometry Type (RETIRED)
	ImageGeometryTypeRETIRED = New(0x0020, 0x0070)

	// MaskingImageRETIRED (0020,0080) VR=CS VM=1-n Masking Image (RETIRED)
	MaskingImageRETIRED = New(0x0020, 0x0080)

	// ReportNumberRETIRED (0020,00AA) VR=IS VM=1 Report Number (RETIRED)
	ReportNumberRETIRED = New(0x0020, 0x00AA)

	// TemporalPositionIdentifier (0020,0100) VR=IS VM=1 Temporal Position Identifier
	TemporalPositionIdentifier = New(0x0020, 0x0100)

	// NumberOfTemporalPositions (0020,0105) VR=IS VM=1 Number of Temporal Positions
	NumberOfTemporalPositions = New(0x0020, 0x0105)

	// TemporalResolution (0020,0110) VR=DS VM=1 Temporal Resolution
	TemporalResolution = New(0x0020, 0x0110)

	// SynchronizationFrameOfReferenceUID (0020,0200) VR=UI VM=1 Synchronization Frame of Reference UID
	SynchronizationFrameOfReferenceUID = New(0x0020, 0x0200)

	// SOPInstanceUIDOfConcatenationSource (0020,0242) VR=UI VM=1 SOP Instance UID of Concatenation Source
	SOPInstanceUIDOfConcatenationSource = New(0x0020, 0x0242)

	// SeriesInStudyRETIRED (0020,1000) VR=IS VM=1 Series in Study (RETIRED)
	SeriesInStudyRETIRED = New(0x0020, 0x1000)

	// AcquisitionsInSeriesRETIRED (0020,1001) VR=IS VM=1 Acquisitions in Series (RETIRED)
	AcquisitionsInSeriesRETIRED = New(0x0020, 0x1001)

	// ImagesInAcquisition (0020,1002) VR=IS VM=1 Images in Acquisition
	ImagesInAcquisition = New(0x0020, 0x1002)

	// ImagesInSeriesRETIRED (0020,1003) VR=IS VM=1 Images in Series (RETIRED)
	ImagesInSeriesRETIRED = New(0x0020, 0x1003)

	// AcquisitionsInStudyRETIRED (0020,1004) VR=IS VM=1 Acquisitions in Study (RETIRED)
	AcquisitionsInStudyRETIRED = New(0x0020, 0x1004)

	// ImagesInStudyRETIRED (0020,1005) VR=IS VM=1 Images in Study (RETIRED)
	ImagesInStudyRETIRED = New(0x0020, 0x1005)

	// ReferenceRETIRED (0020,1020) VR=LO VM=1-n Reference (RETIRED)
	ReferenceRETIRED = New(0x0020, 0x1020)

	// TargetPositionReferenceIndicator (0020,103F) VR=LO VM=1 Target Position Reference Indicator
	TargetPositionReferenceIndicator = New(0x0020, 0x103F)

	// PositionReferenceIndicator (0020,1040) VR=LO VM=1 Position Reference Indicator
	PositionReferenceIndicator = New(0x0020, 0x1040)

	// SliceLocation (0020,1041) VR=DS VM=1 Slice Location
	SliceLocation = New(0x0020, 0x1041)

	// OtherStudyNumbersRETIRED (0020,1070) VR=IS VM=1-n Other Study Numbers (RETIRED)
	OtherStudyNumbersRETIRED = New(0x0020, 0x1070)

	// NumberOfPatientRelatedStudies (0020,1200) VR=IS VM=1 Number of Patient Related Studies
	NumberOfPatientRelatedStudies = New(0x0020, 0x1200)

	// NumberOfPatientRelatedSeries (0020,1202) VR=IS VM=1 Number of Patient Related Series
	NumberOfPatientRelatedSeries = New(0x0020, 0x1202)

	// NumberOfPatientRelatedInstances (0020,1204) VR=IS VM=1 Number of Patient Related Instances
	NumberOfPatientRelatedInstances = New(0x0020, 0x1204)

	// NumberOfStudyRelatedSeries (0020,1206) VR=IS VM=1 Number of Study Related Series
	NumberOfStudyRelatedSeries = New(0x0020, 0x1206)

	// NumberOfStudyRelatedInstances (0020,1208) VR=IS VM=1 Number of Study Related Instances
	NumberOfStudyRelatedInstances = New(0x0020, 0x1208)

	// NumberOfSeriesRelatedInstances (0020,1209) VR=IS VM=1 Number of Series Related Instances
	NumberOfSeriesRelatedInstances = New(0x0020, 0x1209)

	SourceImageIDsRETIRED = New(0x0020, 0x3100)

	// ModifyingDeviceIDRETIRED (0020,3401) VR=CS VM=1 Modifying Device ID (RETIRED)
	ModifyingDeviceIDRETIRED = New(0x0020, 0x3401)

	// ModifiedImageIDRETIRED (0020,3402) VR=CS VM=1 Modified Image ID (RETIRED)
	ModifiedImageIDRETIRED = New(0x0020, 0x3402)

	// ModifiedImageDateRETIRED (0020,3403) VR=DA VM=1 Modified Image Date (RETIRED)
	ModifiedImageDateRETIRED = New(0x0020, 0x3403)

	// ModifyingDeviceManufacturerRETIRED (0020,3404) VR=LO VM=1 Modifying Device Manufacturer (RETIRED)
	ModifyingDeviceManufacturerRETIRED = New(0x0020, 0x3404)

	// ModifiedImageTimeRETIRED (0020,3405) VR=TM VM=1 Modified Image Time (RETIRED)
	ModifiedImageTimeRETIRED = New(0x0020, 0x3405)

	// ModifiedImageDescriptionRETIRED (0020,3406) VR=LO VM=1 Modified Image Description (RETIRED)
	ModifiedImageDescriptionRETIRED = New(0x0020, 0x3406)

	// ImageComments (0020,4000) VR=LT VM=1 Image Comments
	ImageComments = New(0x0020, 0x4000)

	// OriginalImageIdentificationRETIRED (0020,5000) VR=AT VM=1-n Original Image Identification (RETIRED)
	OriginalImageIdentificationRETIRED = New(0x0020, 0x5000)

	// OriginalImageIdentificationNomenclatureRETIRED (0020,5002) VR=LO VM=1-n Original Image Identification Nomenclature (RETIRED)
	OriginalImageIdentificationNomenclatureRETIRED = New(0x0020, 0x5002)

	// StackID (0020,9056) VR=SH VM=1 Stack ID
	StackID = New(0x0020, 0x9056)

	// InStackPositionNumber (0020,9057) VR=UL VM=1 In-Stack Position Number
	InStackPositionNumber = New(0x0020, 0x9057)

	// FrameAnatomySequence (0020,9071) VR=SQ VM=1 Frame Anatomy Sequence
	FrameAnatomySequence = New(0x0020, 0x9071)

	// FrameLaterality (0020,9072) VR=CS VM=1 Frame Laterality
	FrameLaterality = New(0x0020, 0x9072)

	// FrameContentSequence (0020,9111) VR=SQ VM=1 Frame Content Sequence
	FrameContentSequence = New(0x0020, 0x9111)

	// PlanePositionSequence (0020,9113) VR=SQ VM=1 Plane Position Sequence
	PlanePositionSequence = New(0x0020, 0x9113)

	// PlaneOrientationSequence (0020,9116) VR=SQ VM=1 Plane Orientation Sequence
	PlaneOrientationSequence = New(0x0020, 0x9116)

	// TemporalPositionIndex (0020,9128) VR=UL VM=1 Temporal Position Index
	TemporalPositionIndex = New(0x0020, 0x9128)

	// NominalCardiacTriggerDelayTime (0020,9153) VR=FD VM=1 Nominal Cardiac Trigger Delay Time
	NominalCardiacTriggerDelayTime = New(0x0020, 0x9153)

	// NominalCardiacTriggerTimePriorToRPeak (0020,9154) VR=FL VM=1 Nominal Cardiac Trigger Time Prior To R-Peak
	NominalCardiacTriggerTimePriorToRPeak = New(0x0020, 0x9154)

	// ActualCardiacTriggerTimePriorToRPeak (0020,9155) VR=FL VM=1 Actual Cardiac Trigger Time Prior To R-Peak
	ActualCardiacTriggerTimePriorToRPeak = New(0x0020, 0x9155)

	// FrameAcquisitionNumber (0020,9156) VR=US VM=1 Frame Acquisition Number
	FrameAcquisitionNumber = New(0x0020, 0x9156)

	// DimensionIndexValues (0020,9157) VR=UL VM=1-n Dimension Index Values
	DimensionIndexValues = New(0x0020, 0x9157)

	// FrameComments (0020,9158) VR=LT VM=1 Frame Comments
	FrameComments = New(0x0020, 0x9158)

	// ConcatenationUID (0020,9161) VR=UI VM=1 Concatenation UID
	ConcatenationUID = New(0x0020, 0x9161)

	// InConcatenationNumber (0020,9162) VR=US VM=1 In-concatenation Number
	InConcatenationNumber = New(0x0020, 0x9162)

	// InConcatenationTotalNumber (0020,9163) VR=US VM=1 In-concatenation Total Number
	InConcatenationTotalNumber = New(0x0020, 0x9163)

	// DimensionOrganizationUID (0020,9164) VR=UI VM=1 Dimension Organization UID
	DimensionOrganizationUID = New(0x0020, 0x9164)

	// DimensionIndexPointer (0020,9165) VR=AT VM=1 Dimension Index Pointer
	DimensionIndexPointer = New(0x0020, 0x9165)

	// FunctionalGroupPointer (0020,9167) VR=AT VM=1 Functional Group Pointer
	FunctionalGroupPointer = New(0x0020, 0x9167)

	// UnassignedSharedConvertedAttributesSequence (0020,9170) VR=SQ VM=1 Unassigned Shared Converted Attributes Sequence
	UnassignedSharedConvertedAttributesSequence = New(0x0020, 0x9170)

	// UnassignedPerFrameConvertedAttributesSequence (0020,9171) VR=SQ VM=1 Unassigned Per-Frame Converted Attributes Sequence
	UnassignedPerFrameConvertedAttributesSequence = New(0x0020, 0x9171)

	// ConversionSourceAttributesSequence (0020,9172) VR=SQ VM=1 Conversion Source Attributes Sequence
	ConversionSourceAttributesSequence = New(0x0020, 0x9172)

	// DimensionIndexPrivateCreator (0020,9213) VR=LO VM=1 Dimension Index Private Creator
	DimensionIndexPrivateCreator = New(0x0020, 0x9213)

	// DimensionOrganizationSequence (0020,9221) VR=SQ VM=1 Dimension Organization Sequence
	DimensionOrganizationSequence = New(0x0020, 0x9221)

	// DimensionIndexSequence (0020,9222) VR=SQ VM=1 Dimension Index Sequence
	DimensionIndexSequence = New(0x0020, 0x9222)

	// ConcatenationFrameOffsetNumber (0020,9228) VR=UL VM=1 Concatenation Frame Offset Number
	ConcatenationFrameOffsetNumber = New(0x0020, 0x9228)

	// FunctionalGroupPrivateCreator (0020,9238) VR=LO VM=1 Functional Group Private Creator
	FunctionalGroupPrivateCreator = New(0x0020, 0x9238)

	// NominalPercentageOfCardiacPhase (0020,9241) VR=FL VM=1 Nominal Percentage of Cardiac Phase
	NominalPercentageOfCardiacPhase = New(0x0020, 0x9241)

	// NominalPercentageOfRespiratoryPhase (0020,9245) VR=FL VM=1 Nominal Percentage of Respiratory Phase
	NominalPercentageOfRespiratoryPhase = New(0x0020, 0x9245)

	// StartingRespiratoryAmplitude (0020,9246) VR=FL VM=1 Starting Respiratory Amplitude
	StartingRespiratoryAmplitude = New(0x0020, 0x9246)

	// StartingRespiratoryPhase (0020,9247) VR=CS VM=1 Starting Respiratory Phase
	StartingRespiratoryPhase = New(0x0020, 0x9247)

	// EndingRespiratoryAmplitude (0020,9248) VR=FL VM=1 Ending Respiratory Amplitude
	EndingRespiratoryAmplitude = New(0x0020, 0x9248)

	// EndingRespiratoryPhase (0020,9249) VR=CS VM=1 Ending Respiratory Phase
	EndingRespiratoryPhase = New(0x0020, 0x9249)

	// RespiratoryTriggerType (0020,9250) VR=CS VM=1 Respiratory Trigger Type
	RespiratoryTriggerType = New(0x0020, 0x9250)

	// RRIntervalTimeNominal (0020,9251) VR=FD VM=1 R-R Interval Time Nominal
	RRIntervalTimeNominal = New(0x0020, 0x9251)

	// ActualCardiacTriggerDelayTime (0020,9252) VR=FD VM=1 Actual Cardiac Trigger Delay Time
	ActualCardiacTriggerDelayTime = New(0x0020, 0x9252)

	// RespiratorySynchronizationSequence (0020,9253) VR=SQ VM=1 Respiratory Synchronization Sequence
	RespiratorySynchronizationSequence = New(0x0020, 0x9253)

	// RespiratoryIntervalTime (0020,9254) VR=FD VM=1 Respiratory Interval Time
	RespiratoryIntervalTime = New(0x0020, 0x9254)

	// NominalRespiratoryTriggerDelayTime (0020,9255) VR=FD VM=1 Nominal Respiratory Trigger Delay Time
	NominalRespiratoryTriggerDelayTime = New(0x0020, 0x9255)

	// RespiratoryTriggerDelayThreshold (0020,9256) VR=FD VM=1 Respiratory Trigger Delay Threshold
	RespiratoryTriggerDelayThreshold = New(0x0020, 0x9256)

	// ActualRespiratoryTriggerDelayTime (0020,9257) VR=FD VM=1 Actual Respiratory Trigger Delay Time
	ActualRespiratoryTriggerDelayTime = New(0x0020, 0x9257)

	// ImagePositionVolume (0020,9301) VR=FD VM=3 Image Position (Volume)
	ImagePositionVolume = New(0x0020, 0x9301)

	// ImageOrientationVolume (0020,9302) VR=FD VM=6 Image Orientation (Volume)
	ImageOrientationVolume = New(0x0020, 0x9302)

	// UltrasoundAcquisitionGeometry (0020,9307) VR=CS VM=1 Ultrasound Acquisition Geometry
	UltrasoundAcquisitionGeometry = New(0x0020, 0x9307)

	// ApexPosition (0020,9308) VR=FD VM=3 Apex Position
	ApexPosition = New(0x0020, 0x9308)

	// VolumeToTransducerMappingMatrix (0020,9309) VR=FD VM=16 Volume to Transducer Mapping Matrix
	VolumeToTransducerMappingMatrix = New(0x0020, 0x9309)

	// VolumeToTableMappingMatrix (0020,930A) VR=FD VM=16 Volume to Table Mapping Matrix
	VolumeToTableMappingMatrix = New(0x0020, 0x930A)

	// VolumeToTransducerRelationship (0020,930B) VR=CS VM=1 Volume to Transducer Relationship
	VolumeToTransducerRelationship = New(0x0020, 0x930B)

	// PatientFrameOfReferenceSource (0020,930C) VR=CS VM=1 Patient Frame of Reference Source
	PatientFrameOfReferenceSource = New(0x0020, 0x930C)

	// TemporalPositionTimeOffset (0020,930D) VR=FD VM=1 Temporal Position Time Offset
	TemporalPositionTimeOffset = New(0x0020, 0x930D)

	// PlanePositionVolumeSequence (0020,930E) VR=SQ VM=1 Plane Position (Volume) Sequence
	PlanePositionVolumeSequence = New(0x0020, 0x930E)

	// PlaneOrientationVolumeSequence (0020,930F) VR=SQ VM=1 Plane Orientation (Volume) Sequence
	PlaneOrientationVolumeSequence = New(0x0020, 0x930F)

	// TemporalPositionSequence (0020,9310) VR=SQ VM=1 Temporal Position Sequence
	TemporalPositionSequence = New(0x0020, 0x9310)

	// DimensionOrganizationType (0020,9311) VR=CS VM=1 Dimension Organization Type
	DimensionOrganizationType = New(0x0020, 0x9311)

	// VolumeFrameOfReferenceUID (0020,9312) VR=UI VM=1 Volume Frame of Reference UID
	VolumeFrameOfReferenceUID = New(0x0020, 0x9312)

	// TableFrameOfReferenceUID (0020,9313) VR=UI VM=1 Table Frame of Reference UID
	TableFrameOfReferenceUID = New(0x0020, 0x9313)

	// DimensionDescriptionLabel (0020,9421) VR=LO VM=1 Dimension Description Label
	DimensionDescriptionLabel = New(0x0020, 0x9421)

	// PatientOrientationInFrameSequence (0020,9450) VR=SQ VM=1 Patient Orientation in Frame Sequence
	PatientOrientationInFrameSequence = New(0x0020, 0x9450)

	// FrameLabel (0020,9453) VR=LO VM=1 Frame Label
	FrameLabel = New(0x0020, 0x9453)

	// AcquisitionIndex (0020,9518) VR=US VM=1-n Acquisition Index
	AcquisitionIndex = New(0x0020, 0x9518)

	// ContributingSOPInstancesReferenceSequence (0020,9529) VR=SQ VM=1 Contributing SOP Instances Reference Sequence
	ContributingSOPInstancesReferenceSequence = New(0x0020, 0x9529)

	// ReconstructionIndex (0020,9536) VR=US VM=1 Reconstruction Index
	ReconstructionIndex = New(0x0020, 0x9536)

	// LightPathFilterPassThroughWavelength (0022,0001) VR=US VM=1 Light Path Filter Pass-Through Wavelength
	LightPathFilterPassThroughWavelength = New(0x0022, 0x0001)

	// LightPathFilterPassBand (0022,0002) VR=US VM=2 Light Path Filter Pass Band
	LightPathFilterPassBand = New(0x0022, 0x0002)

	// ImagePathFilterPassThroughWavelength (0022,0003) VR=US VM=1 Image Path Filter Pass-Through Wavelength
	ImagePathFilterPassThroughWavelength = New(0x0022, 0x0003)

	// ImagePathFilterPassBand (0022,0004) VR=US VM=2 Image Path Filter Pass Band
	ImagePathFilterPassBand = New(0x0022, 0x0004)

	// PatientEyeMovementCommanded (0022,0005) VR=CS VM=1 Patient Eye Movement Commanded
	PatientEyeMovementCommanded = New(0x0022, 0x0005)

	// PatientEyeMovementCommandCodeSequence (0022,0006) VR=SQ VM=1 Patient Eye Movement Command Code Sequence
	PatientEyeMovementCommandCodeSequence = New(0x0022, 0x0006)

	// SphericalLensPower (0022,0007) VR=FL VM=1 Spherical Lens Power
	SphericalLensPower = New(0x0022, 0x0007)

	// CylinderLensPower (0022,0008) VR=FL VM=1 Cylinder Lens Power
	CylinderLensPower = New(0x0022, 0x0008)

	// CylinderAxis (0022,0009) VR=FL VM=1 Cylinder Axis
	CylinderAxis = New(0x0022, 0x0009)

	// EmmetropicMagnification (0022,000A) VR=FL VM=1 Emmetropic Magnification
	EmmetropicMagnification = New(0x0022, 0x000A)

	// IntraOcularPressure (0022,000B) VR=FL VM=1 Intra Ocular Pressure
	IntraOcularPressure = New(0x0022, 0x000B)

	// HorizontalFieldOfView (0022,000C) VR=FL VM=1 Horizontal Field of View
	HorizontalFieldOfView = New(0x0022, 0x000C)

	// PupilDilated (0022,000D) VR=CS VM=1 Pupil Dilated
	PupilDilated = New(0x0022, 0x000D)

	// DegreeOfDilation (0022,000E) VR=FL VM=1 Degree of Dilation
	DegreeOfDilation = New(0x0022, 0x000E)

	// VertexDistance (0022,000F) VR=FD VM=1 Vertex Distance
	VertexDistance = New(0x0022, 0x000F)

	// StereoBaselineAngle (0022,0010) VR=FL VM=1 Stereo Baseline Angle
	StereoBaselineAngle = New(0x0022, 0x0010)

	// StereoBaselineDisplacement (0022,0011) VR=FL VM=1 Stereo Baseline Displacement
	StereoBaselineDisplacement = New(0x0022, 0x0011)

	// StereoHorizontalPixelOffset (0022,0012) VR=FL VM=1 Stereo Horizontal Pixel Offset
	StereoHorizontalPixelOffset = New(0x0022, 0x0012)

	// StereoVerticalPixelOffset (0022,0013) VR=FL VM=1 Stereo Vertical Pixel Offset
	StereoVerticalPixelOffset = New(0x0022, 0x0013)

	// StereoRotation (0022,0014) VR=FL VM=1 Stereo Rotation
	StereoRotation = New(0x0022, 0x0014)

	// AcquisitionDeviceTypeCodeSequence (0022,0015) VR=SQ VM=1 Acquisition Device Type Code Sequence
	AcquisitionDeviceTypeCodeSequence = New(0x0022, 0x0015)

	// IlluminationTypeCodeSequence (0022,0016) VR=SQ VM=1 Illumination Type Code Sequence
	IlluminationTypeCodeSequence = New(0x0022, 0x0016)

	// LightPathFilterTypeStackCodeSequence (0022,0017) VR=SQ VM=1 Light Path Filter Type Stack Code Sequence
	LightPathFilterTypeStackCodeSequence = New(0x0022, 0x0017)

	// ImagePathFilterTypeStackCodeSequence (0022,0018) VR=SQ VM=1 Image Path Filter Type Stack Code Sequence
	ImagePathFilterTypeStackCodeSequence = New(0x0022, 0x0018)

	// LensesCodeSequence (0022,0019) VR=SQ VM=1 Lenses Code Sequence
	LensesCodeSequence = New(0x0022, 0x0019)

	// ChannelDescriptionCodeSequence (0022,001A) VR=SQ VM=1 Channel Description Code Sequence
	ChannelDescriptionCodeSequence = New(0x0022, 0x001A)

	// RefractiveStateSequence (0022,001B) VR=SQ VM=1 Refractive State Sequence
	RefractiveStateSequence = New(0x0022, 0x001B)

	// MydriaticAgentCodeSequence (0022,001C) VR=SQ VM=1 Mydriatic Agent Code Sequence
	MydriaticAgentCodeSequence = New(0x0022, 0x001C)

	// RelativeImagePositionCodeSequence (0022,001D) VR=SQ VM=1 Relative Image Position Code Sequence
	RelativeImagePositionCodeSequence = New(0x0022, 0x001D)

	// CameraAngleOfView (0022,001E) VR=FL VM=1 Camera Angle of View
	CameraAngleOfView = New(0x0022, 0x001E)

	// StereoPairsSequence (0022,0020) VR=SQ VM=1 Stereo Pairs Sequence
	StereoPairsSequence = New(0x0022, 0x0020)

	// LeftImageSequence (0022,0021) VR=SQ VM=1 Left Image Sequence
	LeftImageSequence = New(0x0022, 0x0021)

	// RightImageSequence (0022,0022) VR=SQ VM=1 Right Image Sequence
	RightImageSequence = New(0x0022, 0x0022)

	// StereoPairsPresent (0022,0028) VR=CS VM=1 Stereo Pairs Present
	StereoPairsPresent = New(0x0022, 0x0028)

	// AxialLengthOfTheEye (0022,0030) VR=FL VM=1 Axial Length of the Eye
	AxialLengthOfTheEye = New(0x0022, 0x0030)

	// OphthalmicFrameLocationSequence (0022,0031) VR=SQ VM=1 Ophthalmic Frame Location Sequence
	OphthalmicFrameLocationSequence = New(0x0022, 0x0031)

	// ReferenceCoordinates (0022,0032) VR=FL VM=2-2n Reference Coordinates
	ReferenceCoordinates = New(0x0022, 0x0032)

	// DepthSpatialResolution (0022,0035) VR=FL VM=1 Depth Spatial Resolution
	DepthSpatialResolution = New(0x0022, 0x0035)

	// MaximumDepthDistortion (0022,0036) VR=FL VM=1 Maximum Depth Distortion
	MaximumDepthDistortion = New(0x0022, 0x0036)

	// AlongScanSpatialResolution (0022,0037) VR=FL VM=1 Along-scan Spatial Resolution
	AlongScanSpatialResolution = New(0x0022, 0x0037)

	// MaximumAlongScanDistortion (0022,0038) VR=FL VM=1 Maximum Along-scan Distortion
	MaximumAlongScanDistortion = New(0x0022, 0x0038)

	// OphthalmicImageOrientation (0022,0039) VR=CS VM=1 Ophthalmic Image Orientation
	OphthalmicImageOrientation = New(0x0022, 0x0039)

	// DepthOfTransverseImage (0022,0041) VR=FL VM=1 Depth of Transverse Image
	DepthOfTransverseImage = New(0x0022, 0x0041)

	// MydriaticAgentConcentrationUnitsSequence (0022,0042) VR=SQ VM=1 Mydriatic Agent Concentration Units Sequence
	MydriaticAgentConcentrationUnitsSequence = New(0x0022, 0x0042)

	// AcrossScanSpatialResolution (0022,0048) VR=FL VM=1 Across-scan Spatial Resolution
	AcrossScanSpatialResolution = New(0x0022, 0x0048)

	// MaximumAcrossScanDistortion (0022,0049) VR=FL VM=1 Maximum Across-scan Distortion
	MaximumAcrossScanDistortion = New(0x0022, 0x0049)

	// MydriaticAgentConcentration (0022,004E) VR=DS VM=1 Mydriatic Agent Concentration
	MydriaticAgentConcentration = New(0x0022, 0x004E)

	// IlluminationWaveLength (0022,0055) VR=FL VM=1 Illumination Wave Length
	IlluminationWaveLength = New(0x0022, 0x0055)

	// IlluminationPower (0022,0056) VR=FL VM=1 Illumination Power
	IlluminationPower = New(0x0022, 0x0056)

	// IlluminationBandwidth (0022,0057) VR=FL VM=1 Illumination Bandwidth
	IlluminationBandwidth = New(0x0022, 0x0057)

	// MydriaticAgentSequence (0022,0058) VR=SQ VM=1 Mydriatic Agent Sequence
	MydriaticAgentSequence = New(0x0022, 0x0058)

	// OphthalmicAxialMeasurementsRightEyeSequence (0022,1007) VR=SQ VM=1 Ophthalmic Axial Measurements Right Eye Sequence
	OphthalmicAxialMeasurementsRightEyeSequence = New(0x0022, 0x1007)

	// OphthalmicAxialMeasurementsLeftEyeSequence (0022,1008) VR=SQ VM=1 Ophthalmic Axial Measurements Left Eye Sequence
	OphthalmicAxialMeasurementsLeftEyeSequence = New(0x0022, 0x1008)

	// OphthalmicAxialMeasurementsDeviceType (0022,1009) VR=CS VM=1 Ophthalmic Axial Measurements Device Type
	OphthalmicAxialMeasurementsDeviceType = New(0x0022, 0x1009)

	// OphthalmicAxialLengthMeasurementsType (0022,1010) VR=CS VM=1 Ophthalmic Axial Length Measurements Type
	OphthalmicAxialLengthMeasurementsType = New(0x0022, 0x1010)

	// OphthalmicAxialLengthSequence (0022,1012) VR=SQ VM=1 Ophthalmic Axial Length Sequence
	OphthalmicAxialLengthSequence = New(0x0022, 0x1012)

	// OphthalmicAxialLength (0022,1019) VR=FL VM=1 Ophthalmic Axial Length
	OphthalmicAxialLength = New(0x0022, 0x1019)

	// LensStatusCodeSequence (0022,1024) VR=SQ VM=1 Lens Status Code Sequence
	LensStatusCodeSequence = New(0x0022, 0x1024)

	// VitreousStatusCodeSequence (0022,1025) VR=SQ VM=1 Vitreous Status Code Sequence
	VitreousStatusCodeSequence = New(0x0022, 0x1025)

	// IOLFormulaCodeSequence (0022,1028) VR=SQ VM=1 IOL Formula Code Sequence
	IOLFormulaCodeSequence = New(0x0022, 0x1028)

	// IOLFormulaDetail (0022,1029) VR=LO VM=1 IOL Formula Detail
	IOLFormulaDetail = New(0x0022, 0x1029)

	// KeratometerIndex (0022,1033) VR=FL VM=1 Keratometer Index
	KeratometerIndex = New(0x0022, 0x1033)

	// SourceOfOphthalmicAxialLengthCodeSequence (0022,1035) VR=SQ VM=1 Source of Ophthalmic Axial Length Code Sequence
	SourceOfOphthalmicAxialLengthCodeSequence = New(0x0022, 0x1035)

	// SourceOfCornealSizeDataCodeSequence (0022,1036) VR=SQ VM=1 Source of Corneal Size Data Code Sequence
	SourceOfCornealSizeDataCodeSequence = New(0x0022, 0x1036)

	// TargetRefraction (0022,1037) VR=FL VM=1 Target Refraction
	TargetRefraction = New(0x0022, 0x1037)

	// RefractiveProcedureOccurred (0022,1039) VR=CS VM=1 Refractive Procedure Occurred
	RefractiveProcedureOccurred = New(0x0022, 0x1039)

	// RefractiveSurgeryTypeCodeSequence (0022,1040) VR=SQ VM=1 Refractive Surgery Type Code Sequence
	RefractiveSurgeryTypeCodeSequence = New(0x0022, 0x1040)

	// OphthalmicUltrasoundMethodCodeSequence (0022,1044) VR=SQ VM=1 Ophthalmic Ultrasound Method Code Sequence
	OphthalmicUltrasoundMethodCodeSequence = New(0x0022, 0x1044)

	// SurgicallyInducedAstigmatismSequence (0022,1045) VR=SQ VM=1 Surgically Induced Astigmatism Sequence
	SurgicallyInducedAstigmatismSequence = New(0x0022, 0x1045)

	// TypeOfOpticalCorrection (0022,1046) VR=CS VM=1 Type of Optical Correction
	TypeOfOpticalCorrection = New(0x0022, 0x1046)

	// ToricIOLPowerSequence (0022,1047) VR=SQ VM=1 Toric IOL Power Sequence
	ToricIOLPowerSequence = New(0x0022, 0x1047)

	// PredictedToricErrorSequence (0022,1048) VR=SQ VM=1 Predicted Toric Error Sequence
	PredictedToricErrorSequence = New(0x0022, 0x1048)

	// PreSelectedForImplantation (0022,1049) VR=CS VM=1 Pre-Selected for Implantation
	PreSelectedForImplantation = New(0x0022, 0x1049)

	// ToricIOLPowerForExactEmmetropiaSequence (0022,104A) VR=SQ VM=1 Toric IOL Power for Exact Emmetropia Sequence
	ToricIOLPowerForExactEmmetropiaSequence = New(0x0022, 0x104A)

	// ToricIOLPowerForExactTargetRefractionSequence (0022,104B) VR=SQ VM=1 Toric IOL Power for Exact Target Refraction Sequence
	ToricIOLPowerForExactTargetRefractionSequence = New(0x0022, 0x104B)

	// OphthalmicAxialLengthMeasurementsSequence (0022,1050) VR=SQ VM=1 Ophthalmic Axial Length Measurements Sequence
	OphthalmicAxialLengthMeasurementsSequence = New(0x0022, 0x1050)

	// IOLPower (0022,1053) VR=FL VM=1 IOL Power
	IOLPower = New(0x0022, 0x1053)

	// PredictedRefractiveError (0022,1054) VR=FL VM=1 Predicted Refractive Error
	PredictedRefractiveError = New(0x0022, 0x1054)

	// OphthalmicAxialLengthVelocity (0022,1059) VR=FL VM=1 Ophthalmic Axial Length Velocity
	OphthalmicAxialLengthVelocity = New(0x0022, 0x1059)

	// LensStatusDescription (0022,1065) VR=LO VM=1 Lens Status Description
	LensStatusDescription = New(0x0022, 0x1065)

	// VitreousStatusDescription (0022,1066) VR=LO VM=1 Vitreous Status Description
	VitreousStatusDescription = New(0x0022, 0x1066)

	// IOLPowerSequence (0022,1090) VR=SQ VM=1 IOL Power Sequence
	IOLPowerSequence = New(0x0022, 0x1090)

	// LensConstantSequence (0022,1092) VR=SQ VM=1 Lens Constant Sequence
	LensConstantSequence = New(0x0022, 0x1092)

	// IOLManufacturer (0022,1093) VR=LO VM=1 IOL Manufacturer
	IOLManufacturer = New(0x0022, 0x1093)

	// LensConstantDescriptionRETIRED (0022,1094) VR=LO VM=1 Lens Constant Description (RETIRED)
	LensConstantDescriptionRETIRED = New(0x0022, 0x1094)

	// ImplantName (0022,1095) VR=LO VM=1 Implant Name
	ImplantName = New(0x0022, 0x1095)

	// KeratometryMeasurementTypeCodeSequence (0022,1096) VR=SQ VM=1 Keratometry Measurement Type Code Sequence
	KeratometryMeasurementTypeCodeSequence = New(0x0022, 0x1096)

	// ImplantPartNumber (0022,1097) VR=LO VM=1 Implant Part Number
	ImplantPartNumber = New(0x0022, 0x1097)

	// ReferencedOphthalmicAxialMeasurementsSequence (0022,1100) VR=SQ VM=1 Referenced Ophthalmic Axial Measurements Sequence
	ReferencedOphthalmicAxialMeasurementsSequence = New(0x0022, 0x1100)

	// OphthalmicAxialLengthMeasurementsSegmentNameCodeSequence (0022,1101) VR=SQ VM=1 Ophthalmic Axial Length Measurements Segment Name Code Sequence
	OphthalmicAxialLengthMeasurementsSegmentNameCodeSequence = New(0x0022, 0x1101)

	// RefractiveErrorBeforeRefractiveSurgeryCodeSequence (0022,1103) VR=SQ VM=1 Refractive Error Before Refractive Surgery Code Sequence
	RefractiveErrorBeforeRefractiveSurgeryCodeSequence = New(0x0022, 0x1103)

	// IOLPowerForExactEmmetropia (0022,1121) VR=FL VM=1 IOL Power For Exact Emmetropia
	IOLPowerForExactEmmetropia = New(0x0022, 0x1121)

	// IOLPowerForExactTargetRefraction (0022,1122) VR=FL VM=1 IOL Power For Exact Target Refraction
	IOLPowerForExactTargetRefraction = New(0x0022, 0x1122)

	// AnteriorChamberDepthDefinitionCodeSequence (0022,1125) VR=SQ VM=1 Anterior Chamber Depth Definition Code Sequence
	AnteriorChamberDepthDefinitionCodeSequence = New(0x0022, 0x1125)

	// LensThicknessSequence (0022,1127) VR=SQ VM=1 Lens Thickness Sequence
	LensThicknessSequence = New(0x0022, 0x1127)

	// AnteriorChamberDepthSequence (0022,1128) VR=SQ VM=1 Anterior Chamber Depth Sequence
	AnteriorChamberDepthSequence = New(0x0022, 0x1128)

	// CalculationCommentSequence (0022,112A) VR=SQ VM=1 Calculation Comment Sequence
	CalculationCommentSequence = New(0x0022, 0x112A)

	// CalculationCommentType (0022,112B) VR=CS VM=1 Calculation Comment Type
	CalculationCommentType = New(0x0022, 0x112B)

	// CalculationComment (0022,112C) VR=LT VM=1 Calculation Comment
	CalculationComment = New(0x0022, 0x112C)

	// LensThickness (0022,1130) VR=FL VM=1 Lens Thickness
	LensThickness = New(0x0022, 0x1130)

	// AnteriorChamberDepth (0022,1131) VR=FL VM=1 Anterior Chamber Depth
	AnteriorChamberDepth = New(0x0022, 0x1131)

	// SourceOfLensThicknessDataCodeSequence (0022,1132) VR=SQ VM=1 Source of Lens Thickness Data Code Sequence
	SourceOfLensThicknessDataCodeSequence = New(0x0022, 0x1132)

	// SourceOfAnteriorChamberDepthDataCodeSequence (0022,1133) VR=SQ VM=1 Source of Anterior Chamber Depth Data Code Sequence
	SourceOfAnteriorChamberDepthDataCodeSequence = New(0x0022, 0x1133)

	// SourceOfRefractiveMeasurementsSequence (0022,1134) VR=SQ VM=1 Source of Refractive Measurements Sequence
	SourceOfRefractiveMeasurementsSequence = New(0x0022, 0x1134)

	// SourceOfRefractiveMeasurementsCodeSequence (0022,1135) VR=SQ VM=1 Source of Refractive Measurements Code Sequence
	SourceOfRefractiveMeasurementsCodeSequence = New(0x0022, 0x1135)

	// OphthalmicAxialLengthMeasurementModified (0022,1140) VR=CS VM=1 Ophthalmic Axial Length Measurement Modified
	OphthalmicAxialLengthMeasurementModified = New(0x0022, 0x1140)

	// OphthalmicAxialLengthDataSourceCodeSequence (0022,1150) VR=SQ VM=1 Ophthalmic Axial Length Data Source Code Sequence
	OphthalmicAxialLengthDataSourceCodeSequence = New(0x0022, 0x1150)

	// OphthalmicAxialLengthAcquisitionMethodCodeSequenceRETIRED (0022,1153) VR=SQ VM=1 Ophthalmic Axial Length Acquisition Method Code Sequence (RETIRED)
	OphthalmicAxialLengthAcquisitionMethodCodeSequenceRETIRED = New(0x0022, 0x1153)

	// SignalToNoiseRatio (0022,1155) VR=FL VM=1 Signal to Noise Ratio
	SignalToNoiseRatio = New(0x0022, 0x1155)

	// OphthalmicAxialLengthDataSourceDescription (0022,1159) VR=LO VM=1 Ophthalmic Axial Length Data Source Description
	OphthalmicAxialLengthDataSourceDescription = New(0x0022, 0x1159)

	// OphthalmicAxialLengthMeasurementsTotalLengthSequence (0022,1210) VR=SQ VM=1 Ophthalmic Axial Length Measurements Total Length Sequence
	OphthalmicAxialLengthMeasurementsTotalLengthSequence = New(0x0022, 0x1210)

	// OphthalmicAxialLengthMeasurementsSegmentalLengthSequence (0022,1211) VR=SQ VM=1 Ophthalmic Axial Length Measurements Segmental Length Sequence
	OphthalmicAxialLengthMeasurementsSegmentalLengthSequence = New(0x0022, 0x1211)

	// OphthalmicAxialLengthMeasurementsLengthSummationSequence (0022,1212) VR=SQ VM=1 Ophthalmic Axial Length Measurements Length Summation Sequence
	OphthalmicAxialLengthMeasurementsLengthSummationSequence = New(0x0022, 0x1212)

	// UltrasoundOphthalmicAxialLengthMeasurementsSequence (0022,1220) VR=SQ VM=1 Ultrasound Ophthalmic Axial Length Measurements Sequence
	UltrasoundOphthalmicAxialLengthMeasurementsSequence = New(0x0022, 0x1220)

	// OpticalOphthalmicAxialLengthMeasurementsSequence (0022,1225) VR=SQ VM=1 Optical Ophthalmic Axial Length Measurements Sequence
	OpticalOphthalmicAxialLengthMeasurementsSequence = New(0x0022, 0x1225)

	// UltrasoundSelectedOphthalmicAxialLengthSequence (0022,1230) VR=SQ VM=1 Ultrasound Selected Ophthalmic Axial Length Sequence
	UltrasoundSelectedOphthalmicAxialLengthSequence = New(0x0022, 0x1230)

	// OphthalmicAxialLengthSelectionMethodCodeSequence (0022,1250) VR=SQ VM=1 Ophthalmic Axial Length Selection Method Code Sequence
	OphthalmicAxialLengthSelectionMethodCodeSequence = New(0x0022, 0x1250)

	// OpticalSelectedOphthalmicAxialLengthSequence (0022,1255) VR=SQ VM=1 Optical Selected Ophthalmic Axial Length Sequence
	OpticalSelectedOphthalmicAxialLengthSequence = New(0x0022, 0x1255)

	// SelectedSegmentalOphthalmicAxialLengthSequence (0022,1257) VR=SQ VM=1 Selected Segmental Ophthalmic Axial Length Sequence
	SelectedSegmentalOphthalmicAxialLengthSequence = New(0x0022, 0x1257)

	// SelectedTotalOphthalmicAxialLengthSequence (0022,1260) VR=SQ VM=1 Selected Total Ophthalmic Axial Length Sequence
	SelectedTotalOphthalmicAxialLengthSequence = New(0x0022, 0x1260)

	// OphthalmicAxialLengthQualityMetricSequence (0022,1262) VR=SQ VM=1 Ophthalmic Axial Length Quality Metric Sequence
	OphthalmicAxialLengthQualityMetricSequence = New(0x0022, 0x1262)

	// OphthalmicAxialLengthQualityMetricTypeCodeSequenceRETIRED (0022,1265) VR=SQ VM=1 Ophthalmic Axial Length Quality Metric Type Code Sequence (RETIRED)
	OphthalmicAxialLengthQualityMetricTypeCodeSequenceRETIRED = New(0x0022, 0x1265)

	// OphthalmicAxialLengthQualityMetricTypeDescriptionRETIRED (0022,1273) VR=LO VM=1 Ophthalmic Axial Length Quality Metric Type Description (RETIRED)
	OphthalmicAxialLengthQualityMetricTypeDescriptionRETIRED = New(0x0022, 0x1273)

	// IntraocularLensCalculationsRightEyeSequence (0022,1300) VR=SQ VM=1 Intraocular Lens Calculations Right Eye Sequence
	IntraocularLensCalculationsRightEyeSequence = New(0x0022, 0x1300)

	// IntraocularLensCalculationsLeftEyeSequence (0022,1310) VR=SQ VM=1 Intraocular Lens Calculations Left Eye Sequence
	IntraocularLensCalculationsLeftEyeSequence = New(0x0022, 0x1310)

	// ReferencedOphthalmicAxialLengthMeasurementQCImageSequence (0022,1330) VR=SQ VM=1 Referenced Ophthalmic Axial Length Measurement QC Image Sequence
	ReferencedOphthalmicAxialLengthMeasurementQCImageSequence = New(0x0022, 0x1330)

	// OphthalmicMappingDeviceType (0022,1415) VR=CS VM=1 Ophthalmic Mapping Device Type
	OphthalmicMappingDeviceType = New(0x0022, 0x1415)

	// AcquisitionMethodCodeSequence (0022,1420) VR=SQ VM=1 Acquisition Method Code Sequence
	AcquisitionMethodCodeSequence = New(0x0022, 0x1420)

	// AcquisitionMethodAlgorithmSequence (0022,1423) VR=SQ VM=1 Acquisition Method Algorithm Sequence
	AcquisitionMethodAlgorithmSequence = New(0x0022, 0x1423)

	// OphthalmicThicknessMapTypeCodeSequence (0022,1436) VR=SQ VM=1 Ophthalmic Thickness Map Type Code Sequence
	OphthalmicThicknessMapTypeCodeSequence = New(0x0022, 0x1436)

	// OphthalmicThicknessMappingNormalsSequence (0022,1443) VR=SQ VM=1 Ophthalmic Thickness Mapping Normals Sequence
	OphthalmicThicknessMappingNormalsSequence = New(0x0022, 0x1443)

	// RetinalThicknessDefinitionCodeSequence (0022,1445) VR=SQ VM=1 Retinal Thickness Definition Code Sequence
	RetinalThicknessDefinitionCodeSequence = New(0x0022, 0x1445)

	// PixelValueMappingToCodedConceptSequence (0022,1450) VR=SQ VM=1 Pixel Value Mapping to Coded Concept Sequence
	PixelValueMappingToCodedConceptSequence = New(0x0022, 0x1450)

	MappedPixelValue = New(0x0022, 0x1452)

	// PixelValueMappingExplanation (0022,1454) VR=LO VM=1 Pixel Value Mapping Explanation
	PixelValueMappingExplanation = New(0x0022, 0x1454)

	// OphthalmicThicknessMapQualityThresholdSequence (0022,1458) VR=SQ VM=1 Ophthalmic Thickness Map Quality Threshold Sequence
	OphthalmicThicknessMapQualityThresholdSequence = New(0x0022, 0x1458)

	// OphthalmicThicknessMapThresholdQualityRating (0022,1460) VR=FL VM=1 Ophthalmic Thickness Map Threshold Quality Rating
	OphthalmicThicknessMapThresholdQualityRating = New(0x0022, 0x1460)

	// AnatomicStructureReferencePoint (0022,1463) VR=FL VM=2 Anatomic Structure Reference Point
	AnatomicStructureReferencePoint = New(0x0022, 0x1463)

	// RegistrationToLocalizerSequence (0022,1465) VR=SQ VM=1 Registration to Localizer Sequence
	RegistrationToLocalizerSequence = New(0x0022, 0x1465)

	// RegisteredLocalizerUnits (0022,1466) VR=CS VM=1 Registered Localizer Units
	RegisteredLocalizerUnits = New(0x0022, 0x1466)

	// RegisteredLocalizerTopLeftHandCorner (0022,1467) VR=FL VM=2 Registered Localizer Top Left Hand Corner
	RegisteredLocalizerTopLeftHandCorner = New(0x0022, 0x1467)

	// RegisteredLocalizerBottomRightHandCorner (0022,1468) VR=FL VM=2 Registered Localizer Bottom Right Hand Corner
	RegisteredLocalizerBottomRightHandCorner = New(0x0022, 0x1468)

	// OphthalmicThicknessMapQualityRatingSequence (0022,1470) VR=SQ VM=1 Ophthalmic Thickness Map Quality Rating Sequence
	OphthalmicThicknessMapQualityRatingSequence = New(0x0022, 0x1470)

	// RelevantOPTAttributesSequence (0022,1472) VR=SQ VM=1 Relevant OPT Attributes Sequence
	RelevantOPTAttributesSequence = New(0x0022, 0x1472)

	// TransformationMethodCodeSequence (0022,1512) VR=SQ VM=1 Transformation Method Code Sequence
	TransformationMethodCodeSequence = New(0x0022, 0x1512)

	// TransformationAlgorithmSequence (0022,1513) VR=SQ VM=1 Transformation Algorithm Sequence
	TransformationAlgorithmSequence = New(0x0022, 0x1513)

	// OphthalmicAxialLengthMethod (0022,1515) VR=CS VM=1 Ophthalmic Axial Length Method
	OphthalmicAxialLengthMethod = New(0x0022, 0x1515)

	// OphthalmicFOV (0022,1517) VR=FL VM=1 Ophthalmic FOV
	OphthalmicFOV = New(0x0022, 0x1517)

	// TwoDimensionalToThreeDimensionalMapSequence (0022,1518) VR=SQ VM=1 Two Dimensional to Three Dimensional Map Sequence
	TwoDimensionalToThreeDimensionalMapSequence = New(0x0022, 0x1518)

	// WideFieldOphthalmicPhotographyQualityRatingSequence (0022,1525) VR=SQ VM=1 Wide Field Ophthalmic Photography Quality Rating Sequence
	WideFieldOphthalmicPhotographyQualityRatingSequence = New(0x0022, 0x1525)

	// WideFieldOphthalmicPhotographyQualityThresholdSequence (0022,1526) VR=SQ VM=1 Wide Field Ophthalmic Photography Quality Threshold Sequence
	WideFieldOphthalmicPhotographyQualityThresholdSequence = New(0x0022, 0x1526)

	// WideFieldOphthalmicPhotographyThresholdQualityRating (0022,1527) VR=FL VM=1 Wide Field Ophthalmic Photography Threshold Quality Rating
	WideFieldOphthalmicPhotographyThresholdQualityRating = New(0x0022, 0x1527)

	// XCoordinatesCenterPixelViewAngle (0022,1528) VR=FL VM=1 X Coordinates Center Pixel View Angle
	XCoordinatesCenterPixelViewAngle = New(0x0022, 0x1528)

	// YCoordinatesCenterPixelViewAngle (0022,1529) VR=FL VM=1 Y Coordinates Center Pixel View Angle
	YCoordinatesCenterPixelViewAngle = New(0x0022, 0x1529)

	// NumberOfMapPoints (0022,1530) VR=UL VM=1 Number of Map Points
	NumberOfMapPoints = New(0x0022, 0x1530)

	// TwoDimensionalToThreeDimensionalMapData (0022,1531) VR=OF VM=1 Two Dimensional to Three Dimensional Map Data
	TwoDimensionalToThreeDimensionalMapData = New(0x0022, 0x1531)

	// DerivationAlgorithmSequence (0022,1612) VR=SQ VM=1 Derivation Algorithm Sequence
	DerivationAlgorithmSequence = New(0x0022, 0x1612)

	// OphthalmicImageTypeCodeSequence (0022,1615) VR=SQ VM=1 Ophthalmic Image Type Code Sequence
	OphthalmicImageTypeCodeSequence = New(0x0022, 0x1615)

	// OphthalmicImageTypeDescription (0022,1616) VR=LO VM=1 Ophthalmic Image Type Description
	OphthalmicImageTypeDescription = New(0x0022, 0x1616)

	// ScanPatternTypeCodeSequence (0022,1618) VR=SQ VM=1 Scan Pattern Type Code Sequence
	ScanPatternTypeCodeSequence = New(0x0022, 0x1618)

	// ReferencedSurfaceMeshIdentificationSequence (0022,1620) VR=SQ VM=1 Referenced Surface Mesh Identification Sequence
	ReferencedSurfaceMeshIdentificationSequence = New(0x0022, 0x1620)

	// OphthalmicVolumetricPropertiesFlag (0022,1622) VR=CS VM=1 Ophthalmic Volumetric Properties Flag
	OphthalmicVolumetricPropertiesFlag = New(0x0022, 0x1622)

	// OphthalmicAnatomicReferencePointFrameCoordinate (0022,1623) VR=FL VM=1 Ophthalmic Anatomic Reference Point Frame Coordinate
	OphthalmicAnatomicReferencePointFrameCoordinate = New(0x0022, 0x1623)

	// OphthalmicAnatomicReferencePointXCoordinate (0022,1624) VR=FL VM=1 Ophthalmic Anatomic Reference Point X-Coordinate
	OphthalmicAnatomicReferencePointXCoordinate = New(0x0022, 0x1624)

	// OphthalmicAnatomicReferencePointYCoordinate (0022,1626) VR=FL VM=1 Ophthalmic Anatomic Reference Point Y-Coordinate
	OphthalmicAnatomicReferencePointYCoordinate = New(0x0022, 0x1626)

	// OphthalmicEnFaceVolumeDescriptorSequence (0022,1627) VR=SQ VM=1 Ophthalmic En Face Volume Descriptor Sequence
	OphthalmicEnFaceVolumeDescriptorSequence = New(0x0022, 0x1627)

	// OphthalmicEnFaceImageQualityRatingSequence (0022,1628) VR=SQ VM=1 Ophthalmic En Face Image Quality Rating Sequence
	OphthalmicEnFaceImageQualityRatingSequence = New(0x0022, 0x1628)

	// OphthalmicEnFaceVolumeDescriptorScope (0022,1629) VR=CS VM=1 Ophthalmic En Face Volume Descriptor Scope
	OphthalmicEnFaceVolumeDescriptorScope = New(0x0022, 0x1629)

	// QualityThreshold (0022,1630) VR=DS VM=1 Quality Threshold
	QualityThreshold = New(0x0022, 0x1630)

	// OphthalmicAnatomicReferencePointSequence (0022,1632) VR=SQ VM=1 Ophthalmic Anatomic Reference Point Sequence
	OphthalmicAnatomicReferencePointSequence = New(0x0022, 0x1632)

	// OphthalmicAnatomicReferencePointLocalizationType (0022,1633) VR=CS VM=1 Ophthalmic Anatomic Reference Point Localization Type
	OphthalmicAnatomicReferencePointLocalizationType = New(0x0022, 0x1633)

	// PrimaryAnatomicStructureItemIndex (0022,1634) VR=IS VM=1 Primary Anatomic Structure Item Index
	PrimaryAnatomicStructureItemIndex = New(0x0022, 0x1634)

	// OCTBscanAnalysisAcquisitionParametersSequence (0022,1640) VR=SQ VM=1 OCT B-scan Analysis Acquisition Parameters Sequence
	OCTBscanAnalysisAcquisitionParametersSequence = New(0x0022, 0x1640)

	// NumberOfBscansPerFrame (0022,1642) VR=UL VM=1 Number of B-scans Per Frame
	NumberOfBscansPerFrame = New(0x0022, 0x1642)

	// BscanSlabThickness (0022,1643) VR=FL VM=1 B-scan Slab Thickness
	BscanSlabThickness = New(0x0022, 0x1643)

	// DistanceBetweenBscanSlabs (0022,1644) VR=FL VM=1 Distance Between B-scan Slabs
	DistanceBetweenBscanSlabs = New(0x0022, 0x1644)

	// BscanCycleTime (0022,1645) VR=FL VM=1 B-scan Cycle Time
	BscanCycleTime = New(0x0022, 0x1645)

	// BscanCycleTimeVector (0022,1646) VR=FL VM=1-n B-scan Cycle Time Vector
	BscanCycleTimeVector = New(0x0022, 0x1646)

	// AscanRate (0022,1649) VR=FL VM=1 A-scan Rate
	AscanRate = New(0x0022, 0x1649)

	// BscanRate (0022,1650) VR=FL VM=1 B-scan Rate
	BscanRate = New(0x0022, 0x1650)

	// SurfaceMeshZPixelOffset (0022,1658) VR=UL VM=1 Surface Mesh Z-Pixel Offset
	SurfaceMeshZPixelOffset = New(0x0022, 0x1658)

	// VisualFieldHorizontalExtent (0024,0010) VR=FL VM=1 Visual Field Horizontal Extent
	VisualFieldHorizontalExtent = New(0x0024, 0x0010)

	// VisualFieldVerticalExtent (0024,0011) VR=FL VM=1 Visual Field Vertical Extent
	VisualFieldVerticalExtent = New(0x0024, 0x0011)

	// VisualFieldShape (0024,0012) VR=CS VM=1 Visual Field Shape
	VisualFieldShape = New(0x0024, 0x0012)

	// ScreeningTestModeCodeSequence (0024,0016) VR=SQ VM=1 Screening Test Mode Code Sequence
	ScreeningTestModeCodeSequence = New(0x0024, 0x0016)

	// MaximumStimulusLuminance (0024,0018) VR=FL VM=1 Maximum Stimulus Luminance
	MaximumStimulusLuminance = New(0x0024, 0x0018)

	// BackgroundLuminance (0024,0020) VR=FL VM=1 Background Luminance
	BackgroundLuminance = New(0x0024, 0x0020)

	// StimulusColorCodeSequence (0024,0021) VR=SQ VM=1 Stimulus Color Code Sequence
	StimulusColorCodeSequence = New(0x0024, 0x0021)

	// BackgroundIlluminationColorCodeSequence (0024,0024) VR=SQ VM=1 Background Illumination Color Code Sequence
	BackgroundIlluminationColorCodeSequence = New(0x0024, 0x0024)

	// StimulusArea (0024,0025) VR=FL VM=1 Stimulus Area
	StimulusArea = New(0x0024, 0x0025)

	// StimulusPresentationTime (0024,0028) VR=FL VM=1 Stimulus Presentation Time
	StimulusPresentationTime = New(0x0024, 0x0028)

	// FixationSequence (0024,0032) VR=SQ VM=1 Fixation Sequence
	FixationSequence = New(0x0024, 0x0032)

	// FixationMonitoringCodeSequence (0024,0033) VR=SQ VM=1 Fixation Monitoring Code Sequence
	FixationMonitoringCodeSequence = New(0x0024, 0x0033)

	// VisualFieldCatchTrialSequence (0024,0034) VR=SQ VM=1 Visual Field Catch Trial Sequence
	VisualFieldCatchTrialSequence = New(0x0024, 0x0034)

	// FixationCheckedQuantity (0024,0035) VR=US VM=1 Fixation Checked Quantity
	FixationCheckedQuantity = New(0x0024, 0x0035)

	// PatientNotProperlyFixatedQuantity (0024,0036) VR=US VM=1 Patient Not Properly Fixated Quantity
	PatientNotProperlyFixatedQuantity = New(0x0024, 0x0036)

	// PresentedVisualStimuliDataFlag (0024,0037) VR=CS VM=1 Presented Visual Stimuli Data Flag
	PresentedVisualStimuliDataFlag = New(0x0024, 0x0037)

	// NumberOfVisualStimuli (0024,0038) VR=US VM=1 Number of Visual Stimuli
	NumberOfVisualStimuli = New(0x0024, 0x0038)

	// ExcessiveFixationLossesDataFlag (0024,0039) VR=CS VM=1 Excessive Fixation Losses Data Flag
	ExcessiveFixationLossesDataFlag = New(0x0024, 0x0039)

	// ExcessiveFixationLosses (0024,0040) VR=CS VM=1 Excessive Fixation Losses
	ExcessiveFixationLosses = New(0x0024, 0x0040)

	// StimuliRetestingQuantity (0024,0042) VR=US VM=1 Stimuli Retesting Quantity
	StimuliRetestingQuantity = New(0x0024, 0x0042)

	// CommentsOnPatientPerformanceOfVisualField (0024,0044) VR=LT VM=1 Comments on Patient's Performance of Visual Field
	CommentsOnPatientPerformanceOfVisualField = New(0x0024, 0x0044)

	// FalseNegativesEstimateFlag (0024,0045) VR=CS VM=1 False Negatives Estimate Flag
	FalseNegativesEstimateFlag = New(0x0024, 0x0045)

	// FalseNegativesEstimate (0024,0046) VR=FL VM=1 False Negatives Estimate
	FalseNegativesEstimate = New(0x0024, 0x0046)

	// NegativeCatchTrialsQuantity (0024,0048) VR=US VM=1 Negative Catch Trials Quantity
	NegativeCatchTrialsQuantity = New(0x0024, 0x0048)

	// FalseNegativesQuantity (0024,0050) VR=US VM=1 False Negatives Quantity
	FalseNegativesQuantity = New(0x0024, 0x0050)

	// ExcessiveFalseNegativesDataFlag (0024,0051) VR=CS VM=1 Excessive False Negatives Data Flag
	ExcessiveFalseNegativesDataFlag = New(0x0024, 0x0051)

	// ExcessiveFalseNegatives (0024,0052) VR=CS VM=1 Excessive False Negatives
	ExcessiveFalseNegatives = New(0x0024, 0x0052)

	// FalsePositivesEstimateFlag (0024,0053) VR=CS VM=1 False Positives Estimate Flag
	FalsePositivesEstimateFlag = New(0x0024, 0x0053)

	// FalsePositivesEstimate (0024,0054) VR=FL VM=1 False Positives Estimate
	FalsePositivesEstimate = New(0x0024, 0x0054)

	// CatchTrialsDataFlag (0024,0055) VR=CS VM=1 Catch Trials Data Flag
	CatchTrialsDataFlag = New(0x0024, 0x0055)

	// PositiveCatchTrialsQuantity (0024,0056) VR=US VM=1 Positive Catch Trials Quantity
	PositiveCatchTrialsQuantity = New(0x0024, 0x0056)

	// TestPointNormalsDataFlag (0024,0057) VR=CS VM=1 Test Point Normals Data Flag
	TestPointNormalsDataFlag = New(0x0024, 0x0057)

	// TestPointNormalsSequence (0024,0058) VR=SQ VM=1 Test Point Normals Sequence
	TestPointNormalsSequence = New(0x0024, 0x0058)

	// GlobalDeviationProbabilityNormalsFlag (0024,0059) VR=CS VM=1 Global Deviation Probability Normals Flag
	GlobalDeviationProbabilityNormalsFlag = New(0x0024, 0x0059)

	// FalsePositivesQuantity (0024,0060) VR=US VM=1 False Positives Quantity
	FalsePositivesQuantity = New(0x0024, 0x0060)

	// ExcessiveFalsePositivesDataFlag (0024,0061) VR=CS VM=1 Excessive False Positives Data Flag
	ExcessiveFalsePositivesDataFlag = New(0x0024, 0x0061)

	// ExcessiveFalsePositives (0024,0062) VR=CS VM=1 Excessive False Positives
	ExcessiveFalsePositives = New(0x0024, 0x0062)

	// VisualFieldTestNormalsFlag (0024,0063) VR=CS VM=1 Visual Field Test Normals Flag
	VisualFieldTestNormalsFlag = New(0x0024, 0x0063)

	// ResultsNormalsSequence (0024,0064) VR=SQ VM=1 Results Normals Sequence
	ResultsNormalsSequence = New(0x0024, 0x0064)

	// AgeCorrectedSensitivityDeviationAlgorithmSequence (0024,0065) VR=SQ VM=1 Age Corrected Sensitivity Deviation Algorithm Sequence
	AgeCorrectedSensitivityDeviationAlgorithmSequence = New(0x0024, 0x0065)

	// GlobalDeviationFromNormal (0024,0066) VR=FL VM=1 Global Deviation From Normal
	GlobalDeviationFromNormal = New(0x0024, 0x0066)

	// GeneralizedDefectSensitivityDeviationAlgorithmSequence (0024,0067) VR=SQ VM=1 Generalized Defect Sensitivity Deviation Algorithm Sequence
	GeneralizedDefectSensitivityDeviationAlgorithmSequence = New(0x0024, 0x0067)

	// LocalizedDeviationFromNormal (0024,0068) VR=FL VM=1 Localized Deviation From Normal
	LocalizedDeviationFromNormal = New(0x0024, 0x0068)

	// PatientReliabilityIndicator (0024,0069) VR=LO VM=1 Patient Reliability Indicator
	PatientReliabilityIndicator = New(0x0024, 0x0069)

	// VisualFieldMeanSensitivity (0024,0070) VR=FL VM=1 Visual Field Mean Sensitivity
	VisualFieldMeanSensitivity = New(0x0024, 0x0070)

	// GlobalDeviationProbability (0024,0071) VR=FL VM=1 Global Deviation Probability
	GlobalDeviationProbability = New(0x0024, 0x0071)

	// LocalDeviationProbabilityNormalsFlag (0024,0072) VR=CS VM=1 Local Deviation Probability Normals Flag
	LocalDeviationProbabilityNormalsFlag = New(0x0024, 0x0072)

	// LocalizedDeviationProbability (0024,0073) VR=FL VM=1 Localized Deviation Probability
	LocalizedDeviationProbability = New(0x0024, 0x0073)

	// ShortTermFluctuationCalculated (0024,0074) VR=CS VM=1 Short Term Fluctuation Calculated
	ShortTermFluctuationCalculated = New(0x0024, 0x0074)

	// ShortTermFluctuation (0024,0075) VR=FL VM=1 Short Term Fluctuation
	ShortTermFluctuation = New(0x0024, 0x0075)

	// ShortTermFluctuationProbabilityCalculated (0024,0076) VR=CS VM=1 Short Term Fluctuation Probability Calculated
	ShortTermFluctuationProbabilityCalculated = New(0x0024, 0x0076)

	// ShortTermFluctuationProbability (0024,0077) VR=FL VM=1 Short Term Fluctuation Probability
	ShortTermFluctuationProbability = New(0x0024, 0x0077)

	// CorrectedLocalizedDeviationFromNormalCalculated (0024,0078) VR=CS VM=1 Corrected Localized Deviation From Normal Calculated
	CorrectedLocalizedDeviationFromNormalCalculated = New(0x0024, 0x0078)

	// CorrectedLocalizedDeviationFromNormal (0024,0079) VR=FL VM=1 Corrected Localized Deviation From Normal
	CorrectedLocalizedDeviationFromNormal = New(0x0024, 0x0079)

	// CorrectedLocalizedDeviationFromNormalProbabilityCalculated (0024,0080) VR=CS VM=1 Corrected Localized Deviation From Normal Probability Calculated
	CorrectedLocalizedDeviationFromNormalProbabilityCalculated = New(0x0024, 0x0080)

	// CorrectedLocalizedDeviationFromNormalProbability (0024,0081) VR=FL VM=1 Corrected Localized Deviation From Normal Probability
	CorrectedLocalizedDeviationFromNormalProbability = New(0x0024, 0x0081)

	// GlobalDeviationProbabilitySequence (0024,0083) VR=SQ VM=1 Global Deviation Probability Sequence
	GlobalDeviationProbabilitySequence = New(0x0024, 0x0083)

	// LocalizedDeviationProbabilitySequence (0024,0085) VR=SQ VM=1 Localized Deviation Probability Sequence
	LocalizedDeviationProbabilitySequence = New(0x0024, 0x0085)

	// FovealSensitivityMeasured (0024,0086) VR=CS VM=1 Foveal Sensitivity Measured
	FovealSensitivityMeasured = New(0x0024, 0x0086)

	// FovealSensitivity (0024,0087) VR=FL VM=1 Foveal Sensitivity
	FovealSensitivity = New(0x0024, 0x0087)

	// VisualFieldTestDuration (0024,0088) VR=FL VM=1 Visual Field Test Duration
	VisualFieldTestDuration = New(0x0024, 0x0088)

	// VisualFieldTestPointSequence (0024,0089) VR=SQ VM=1 Visual Field Test Point Sequence
	VisualFieldTestPointSequence = New(0x0024, 0x0089)

	// VisualFieldTestPointXCoordinate (0024,0090) VR=FL VM=1 Visual Field Test Point X-Coordinate
	VisualFieldTestPointXCoordinate = New(0x0024, 0x0090)

	// VisualFieldTestPointYCoordinate (0024,0091) VR=FL VM=1 Visual Field Test Point Y-Coordinate
	VisualFieldTestPointYCoordinate = New(0x0024, 0x0091)

	// AgeCorrectedSensitivityDeviationValue (0024,0092) VR=FL VM=1 Age Corrected Sensitivity Deviation Value
	AgeCorrectedSensitivityDeviationValue = New(0x0024, 0x0092)

	// StimulusResults (0024,0093) VR=CS VM=1 Stimulus Results
	StimulusResults = New(0x0024, 0x0093)

	// SensitivityValue (0024,0094) VR=FL VM=1 Sensitivity Value
	SensitivityValue = New(0x0024, 0x0094)

	// RetestStimulusSeen (0024,0095) VR=CS VM=1 Retest Stimulus Seen
	RetestStimulusSeen = New(0x0024, 0x0095)

	// RetestSensitivityValue (0024,0096) VR=FL VM=1 Retest Sensitivity Value
	RetestSensitivityValue = New(0x0024, 0x0096)

	// VisualFieldTestPointNormalsSequence (0024,0097) VR=SQ VM=1 Visual Field Test Point Normals Sequence
	VisualFieldTestPointNormalsSequence = New(0x0024, 0x0097)

	// QuantifiedDefect (0024,0098) VR=FL VM=1 Quantified Defect
	QuantifiedDefect = New(0x0024, 0x0098)

	// AgeCorrectedSensitivityDeviationProbabilityValue (0024,0100) VR=FL VM=1 Age Corrected Sensitivity Deviation Probability Value
	AgeCorrectedSensitivityDeviationProbabilityValue = New(0x0024, 0x0100)

	// GeneralizedDefectCorrectedSensitivityDeviationFlag (0024,0102) VR=CS VM=1 Generalized Defect Corrected Sensitivity Deviation Flag
	GeneralizedDefectCorrectedSensitivityDeviationFlag = New(0x0024, 0x0102)

	// GeneralizedDefectCorrectedSensitivityDeviationValue (0024,0103) VR=FL VM=1 Generalized Defect Corrected Sensitivity Deviation Value
	GeneralizedDefectCorrectedSensitivityDeviationValue = New(0x0024, 0x0103)

	// GeneralizedDefectCorrectedSensitivityDeviationProbabilityValue (0024,0104) VR=FL VM=1 Generalized Defect Corrected Sensitivity Deviation Probability Value
	GeneralizedDefectCorrectedSensitivityDeviationProbabilityValue = New(0x0024, 0x0104)

	// MinimumSensitivityValue (0024,0105) VR=FL VM=1 Minimum Sensitivity Value
	MinimumSensitivityValue = New(0x0024, 0x0105)

	// BlindSpotLocalized (0024,0106) VR=CS VM=1 Blind Spot Localized
	BlindSpotLocalized = New(0x0024, 0x0106)

	// BlindSpotXCoordinate (0024,0107) VR=FL VM=1 Blind Spot X-Coordinate
	BlindSpotXCoordinate = New(0x0024, 0x0107)

	// BlindSpotYCoordinate (0024,0108) VR=FL VM=1 Blind Spot Y-Coordinate
	BlindSpotYCoordinate = New(0x0024, 0x0108)

	// VisualAcuityMeasurementSequence (0024,0110) VR=SQ VM=1 Visual Acuity Measurement Sequence
	VisualAcuityMeasurementSequence = New(0x0024, 0x0110)

	// RefractiveParametersUsedOnPatientSequence (0024,0112) VR=SQ VM=1 Refractive Parameters Used on Patient Sequence
	RefractiveParametersUsedOnPatientSequence = New(0x0024, 0x0112)

	// MeasurementLaterality (0024,0113) VR=CS VM=1 Measurement Laterality
	MeasurementLaterality = New(0x0024, 0x0113)

	// OphthalmicPatientClinicalInformationLeftEyeSequence (0024,0114) VR=SQ VM=1 Ophthalmic Patient Clinical Information Left Eye Sequence
	OphthalmicPatientClinicalInformationLeftEyeSequence = New(0x0024, 0x0114)

	// OphthalmicPatientClinicalInformationRightEyeSequence (0024,0115) VR=SQ VM=1 Ophthalmic Patient Clinical Information Right Eye Sequence
	OphthalmicPatientClinicalInformationRightEyeSequence = New(0x0024, 0x0115)

	// FovealPointNormativeDataFlag (0024,0117) VR=CS VM=1 Foveal Point Normative Data Flag
	FovealPointNormativeDataFlag = New(0x0024, 0x0117)

	// FovealPointProbabilityValue (0024,0118) VR=FL VM=1 Foveal Point Probability Value
	FovealPointProbabilityValue = New(0x0024, 0x0118)

	// ScreeningBaselineMeasured (0024,0120) VR=CS VM=1 Screening Baseline Measured
	ScreeningBaselineMeasured = New(0x0024, 0x0120)

	// ScreeningBaselineMeasuredSequence (0024,0122) VR=SQ VM=1 Screening Baseline Measured Sequence
	ScreeningBaselineMeasuredSequence = New(0x0024, 0x0122)

	// ScreeningBaselineType (0024,0124) VR=CS VM=1 Screening Baseline Type
	ScreeningBaselineType = New(0x0024, 0x0124)

	// ScreeningBaselineValue (0024,0126) VR=FL VM=1 Screening Baseline Value
	ScreeningBaselineValue = New(0x0024, 0x0126)

	// AlgorithmSource (0024,0202) VR=LO VM=1 Algorithm Source
	AlgorithmSource = New(0x0024, 0x0202)

	// DataSetName (0024,0306) VR=LO VM=1 Data Set Name
	DataSetName = New(0x0024, 0x0306)

	// DataSetVersion (0024,0307) VR=LO VM=1 Data Set Version
	DataSetVersion = New(0x0024, 0x0307)

	// DataSetSource (0024,0308) VR=LO VM=1 Data Set Source
	DataSetSource = New(0x0024, 0x0308)

	// DataSetDescription (0024,0309) VR=LO VM=1 Data Set Description
	DataSetDescription = New(0x0024, 0x0309)

	// VisualFieldTestReliabilityGlobalIndexSequence (0024,0317) VR=SQ VM=1 Visual Field Test Reliability Global Index Sequence
	VisualFieldTestReliabilityGlobalIndexSequence = New(0x0024, 0x0317)

	// VisualFieldGlobalResultsIndexSequence (0024,0320) VR=SQ VM=1 Visual Field Global Results Index Sequence
	VisualFieldGlobalResultsIndexSequence = New(0x0024, 0x0320)

	// DataObservationSequence (0024,0325) VR=SQ VM=1 Data Observation Sequence
	DataObservationSequence = New(0x0024, 0x0325)

	// IndexNormalsFlag (0024,0338) VR=CS VM=1 Index Normals Flag
	IndexNormalsFlag = New(0x0024, 0x0338)

	// IndexProbability (0024,0341) VR=FL VM=1 Index Probability
	IndexProbability = New(0x0024, 0x0341)

	// IndexProbabilitySequence (0024,0344) VR=SQ VM=1 Index Probability Sequence
	IndexProbabilitySequence = New(0x0024, 0x0344)

	// SamplesPerPixel (0028,0002) VR=US VM=1 Samples per Pixel
	SamplesPerPixel = New(0x0028, 0x0002)

	// SamplesPerPixelUsed (0028,0003) VR=US VM=1 Samples per Pixel Used
	SamplesPerPixelUsed = New(0x0028, 0x0003)

	// PhotometricInterpretation (0028,0004) VR=CS VM=1 Photometric Interpretation
	PhotometricInterpretation = New(0x0028, 0x0004)

	// ImageDimensionsRETIRED (0028,0005) VR=US VM=1 Image Dimensions (RETIRED)
	ImageDimensionsRETIRED = New(0x0028, 0x0005)

	// PlanarConfiguration (0028,0006) VR=US VM=1 Planar Configuration
	PlanarConfiguration = New(0x0028, 0x0006)

	// NumberOfFrames (0028,0008) VR=IS VM=1 Number of Frames
	NumberOfFrames = New(0x0028, 0x0008)

	// FrameIncrementPointer (0028,0009) VR=AT VM=1-n Frame Increment Pointer
	FrameIncrementPointer = New(0x0028, 0x0009)

	// FrameDimensionPointer (0028,000A) VR=AT VM=1-n Frame Dimension Pointer
	FrameDimensionPointer = New(0x0028, 0x000A)

	// Rows (0028,0010) VR=US VM=1 Rows
	Rows = New(0x0028, 0x0010)

	// Columns (0028,0011) VR=US VM=1 Columns
	Columns = New(0x0028, 0x0011)

	// PlanesRETIRED (0028,0012) VR=US VM=1 Planes (RETIRED)
	PlanesRETIRED = New(0x0028, 0x0012)

	// UltrasoundColorDataPresent (0028,0014) VR=US VM=1 Ultrasound Color Data Present
	UltrasoundColorDataPresent = New(0x0028, 0x0014)

	// PixelSpacing (0028,0030) VR=DS VM=2 Pixel Spacing
	PixelSpacing = New(0x0028, 0x0030)

	// ZoomFactor (0028,0031) VR=DS VM=2 Zoom Factor
	ZoomFactor = New(0x0028, 0x0031)

	// ZoomCenter (0028,0032) VR=DS VM=2 Zoom Center
	ZoomCenter = New(0x0028, 0x0032)

	// PixelAspectRatio (0028,0034) VR=IS VM=2 Pixel Aspect Ratio
	PixelAspectRatio = New(0x0028, 0x0034)

	// ImageFormatRETIRED (0028,0040) VR=CS VM=1 Image Format (RETIRED)
	ImageFormatRETIRED = New(0x0028, 0x0040)

	// ManipulatedImageRETIRED (0028,0050) VR=LO VM=1-n Manipulated Image (RETIRED)
	ManipulatedImageRETIRED = New(0x0028, 0x0050)

	// CorrectedImage (0028,0051) VR=CS VM=1-n Corrected Image
	CorrectedImage = New(0x0028, 0x0051)

	// CompressionRecognitionCodeRETIRED (0028,005F) VR=LO VM=1 Compression Recognition Code (RETIRED)
	CompressionRecognitionCodeRETIRED = New(0x0028, 0x005F)

	// CompressionCodeRETIRED (0028,0060) VR=CS VM=1 Compression Code (RETIRED)
	CompressionCodeRETIRED = New(0x0028, 0x0060)

	// CompressionOriginatorRETIRED (0028,0061) VR=SH VM=1 Compression Originator (RETIRED)
	CompressionOriginatorRETIRED = New(0x0028, 0x0061)

	// CompressionLabelRETIRED (0028,0062) VR=LO VM=1 Compression Label (RETIRED)
	CompressionLabelRETIRED = New(0x0028, 0x0062)

	// CompressionDescriptionRETIRED (0028,0063) VR=SH VM=1 Compression Description (RETIRED)
	CompressionDescriptionRETIRED = New(0x0028, 0x0063)

	// CompressionSequenceRETIRED (0028,0065) VR=CS VM=1-n Compression Sequence (RETIRED)
	CompressionSequenceRETIRED = New(0x0028, 0x0065)

	// CompressionStepPointersRETIRED (0028,0066) VR=AT VM=1-n Compression Step Pointers (RETIRED)
	CompressionStepPointersRETIRED = New(0x0028, 0x0066)

	// RepeatIntervalRETIRED (0028,0068) VR=US VM=1 Repeat Interval (RETIRED)
	RepeatIntervalRETIRED = New(0x0028, 0x0068)

	// BitsGroupedRETIRED (0028,0069) VR=US VM=1 Bits Grouped (RETIRED)
	BitsGroupedRETIRED = New(0x0028, 0x0069)

	// PerimeterTableRETIRED (0028,0070) VR=US VM=1-n Perimeter Table (RETIRED)
	PerimeterTableRETIRED = New(0x0028, 0x0070)

	PerimeterValueRETIRED = New(0x0028, 0x0071)

	// PredictorRowsRETIRED (0028,0080) VR=US VM=1 Predictor Rows (RETIRED)
	PredictorRowsRETIRED = New(0x0028, 0x0080)

	// PredictorColumnsRETIRED (0028,0081) VR=US VM=1 Predictor Columns (RETIRED)
	PredictorColumnsRETIRED = New(0x0028, 0x0081)

	// PredictorConstantsRETIRED (0028,0082) VR=US VM=1-n Predictor Constants (RETIRED)
	PredictorConstantsRETIRED = New(0x0028, 0x0082)

	// BlockedPixelsRETIRED (0028,0090) VR=CS VM=1 Blocked Pixels (RETIRED)
	BlockedPixelsRETIRED = New(0x0028, 0x0090)

	// BlockRowsRETIRED (0028,0091) VR=US VM=1 Block Rows (RETIRED)
	BlockRowsRETIRED = New(0x0028, 0x0091)

	// BlockColumnsRETIRED (0028,0092) VR=US VM=1 Block Columns (RETIRED)
	BlockColumnsRETIRED = New(0x0028, 0x0092)

	// RowOverlapRETIRED (0028,0093) VR=US VM=1 Row Overlap (RETIRED)
	RowOverlapRETIRED = New(0x0028, 0x0093)

	// ColumnOverlapRETIRED (0028,0094) VR=US VM=1 Column Overlap (RETIRED)
	ColumnOverlapRETIRED = New(0x0028, 0x0094)

	// BitsAllocated (0028,0100) VR=US VM=1 Bits Allocated
	BitsAllocated = New(0x0028, 0x0100)

	// BitsStored (0028,0101) VR=US VM=1 Bits Stored
	BitsStored = New(0x0028, 0x0101)

	// HighBit (0028,0102) VR=US VM=1 High Bit
	HighBit = New(0x0028, 0x0102)

	// PixelRepresentation (0028,0103) VR=US VM=1 Pixel Representation
	PixelRepresentation = New(0x0028, 0x0103)

	SmallestValidPixelValueRETIRED = New(0x0028, 0x0104)

	LargestValidPixelValueRETIRED = New(0x0028, 0x0105)

	SmallestImagePixelValue = New(0x0028, 0x0106)

	LargestImagePixelValue = New(0x0028, 0x0107)

	SmallestPixelValueInSeries = New(0x0028, 0x0108)

	LargestPixelValueInSeries = New(0x0028, 0x0109)

	SmallestImagePixelValueInPlaneRETIRED = New(0x0028, 0x0110)

	LargestImagePixelValueInPlaneRETIRED = New(0x0028, 0x0111)

	PixelPaddingValue = New(0x0028, 0x0120)

	PixelPaddingRangeLimit = New(0x0028, 0x0121)

	// FloatPixelPaddingValue (0028,0122) VR=FL VM=1 Float Pixel Padding Value
	FloatPixelPaddingValue = New(0x0028, 0x0122)

	// DoubleFloatPixelPaddingValue (0028,0123) VR=FD VM=1 Double Float Pixel Padding Value
	DoubleFloatPixelPaddingValue = New(0x0028, 0x0123)

	// FloatPixelPaddingRangeLimit (0028,0124) VR=FL VM=1 Float Pixel Padding Range Limit
	FloatPixelPaddingRangeLimit = New(0x0028, 0x0124)

	// DoubleFloatPixelPaddingRangeLimit (0028,0125) VR=FD VM=1 Double Float Pixel Padding Range Limit
	DoubleFloatPixelPaddingRangeLimit = New(0x0028, 0x0125)

	// ImageLocationRETIRED (0028,0200) VR=US VM=1 Image Location (RETIRED)
	ImageLocationRETIRED = New(0x0028, 0x0200)

	// QualityControlImage (0028,0300) VR=CS VM=1 Quality Control Image
	QualityControlImage = New(0x0028, 0x0300)

	// BurnedInAnnotation (0028,0301) VR=CS VM=1 Burned In Annotation
	BurnedInAnnotation = New(0x0028, 0x0301)

	// RecognizableVisualFeatures (0028,0302) VR=CS VM=1 Recognizable Visual Features
	RecognizableVisualFeatures = New(0x0028, 0x0302)

	// LongitudinalTemporalInformationModified (0028,0303) VR=CS VM=1 Longitudinal Temporal Information Modified
	LongitudinalTemporalInformationModified = New(0x0028, 0x0303)

	// ReferencedColorPaletteInstanceUID (0028,0304) VR=UI VM=1 Referenced Color Palette Instance UID
	ReferencedColorPaletteInstanceUID = New(0x0028, 0x0304)

	// TransformLabelRETIRED (0028,0400) VR=LO VM=1 Transform Label (RETIRED)
	TransformLabelRETIRED = New(0x0028, 0x0400)

	// TransformVersionNumberRETIRED (0028,0401) VR=LO VM=1 Transform Version Number (RETIRED)
	TransformVersionNumberRETIRED = New(0x0028, 0x0401)

	// NumberOfTransformStepsRETIRED (0028,0402) VR=US VM=1 Number of Transform Steps (RETIRED)
	NumberOfTransformStepsRETIRED = New(0x0028, 0x0402)

	// SequenceOfCompressedDataRETIRED (0028,0403) VR=LO VM=1-n Sequence of Compressed Data (RETIRED)
	SequenceOfCompressedDataRETIRED = New(0x0028, 0x0403)

	// DetailsOfCoefficientsRETIRED (0028,0404) VR=AT VM=1-n Details of Coefficients (RETIRED)
	DetailsOfCoefficientsRETIRED = New(0x0028, 0x0404)

	RowsForNthOrderCoefficientsRETIRED = New(0x0028, 0x0400)

	ColumnsForNthOrderCoefficientsRETIRED = New(0x0028, 0x0401)

	CoefficientCodingRETIRED = New(0x0028, 0x0402)

	CoefficientCodingPointersRETIRED = New(0x0028, 0x0403)

	// DCTLabelRETIRED (0028,0700) VR=LO VM=1 DCT Label (RETIRED)
	DCTLabelRETIRED = New(0x0028, 0x0700)

	// DataBlockDescriptionRETIRED (0028,0701) VR=CS VM=1-n Data Block Description (RETIRED)
	DataBlockDescriptionRETIRED = New(0x0028, 0x0701)

	// DataBlockRETIRED (0028,0702) VR=AT VM=1-n Data Block (RETIRED)
	DataBlockRETIRED = New(0x0028, 0x0702)

	// NormalizationFactorFormatRETIRED (0028,0710) VR=US VM=1 Normalization Factor Format (RETIRED)
	NormalizationFactorFormatRETIRED = New(0x0028, 0x0710)

	// ZonalMapNumberFormatRETIRED (0028,0720) VR=US VM=1 Zonal Map Number Format (RETIRED)
	ZonalMapNumberFormatRETIRED = New(0x0028, 0x0720)

	// ZonalMapLocationRETIRED (0028,0721) VR=AT VM=1-n Zonal Map Location (RETIRED)
	ZonalMapLocationRETIRED = New(0x0028, 0x0721)

	// ZonalMapFormatRETIRED (0028,0722) VR=US VM=1 Zonal Map Format (RETIRED)
	ZonalMapFormatRETIRED = New(0x0028, 0x0722)

	// AdaptiveMapFormatRETIRED (0028,0730) VR=US VM=1 Adaptive Map Format (RETIRED)
	AdaptiveMapFormatRETIRED = New(0x0028, 0x0730)

	// CodeNumberFormatRETIRED (0028,0740) VR=US VM=1 Code Number Format (RETIRED)
	CodeNumberFormatRETIRED = New(0x0028, 0x0740)

	CodeLabelRETIRED = New(0x0028, 0x0800)

	NumberOfTablesRETIRED = New(0x0028, 0x0802)

	CodeTableLocationRETIRED = New(0x0028, 0x0803)

	BitsForCodeWordRETIRED = New(0x0028, 0x0804)

	ImageDataLocationRETIRED = New(0x0028, 0x0808)

	// PixelSpacingCalibrationType (0028,0A02) VR=CS VM=1 Pixel Spacing Calibration Type
	PixelSpacingCalibrationType = New(0x0028, 0x0A02)

	// PixelSpacingCalibrationDescription (0028,0A04) VR=LO VM=1 Pixel Spacing Calibration Description
	PixelSpacingCalibrationDescription = New(0x0028, 0x0A04)

	// PixelIntensityRelationship (0028,1040) VR=CS VM=1 Pixel Intensity Relationship
	PixelIntensityRelationship = New(0x0028, 0x1040)

	// PixelIntensityRelationshipSign (0028,1041) VR=SS VM=1 Pixel Intensity Relationship Sign
	PixelIntensityRelationshipSign = New(0x0028, 0x1041)

	// WindowCenter (0028,1050) VR=DS VM=1-n Window Center
	WindowCenter = New(0x0028, 0x1050)

	// WindowWidth (0028,1051) VR=DS VM=1-n Window Width
	WindowWidth = New(0x0028, 0x1051)

	// RescaleIntercept (0028,1052) VR=DS VM=1 Rescale Intercept
	RescaleIntercept = New(0x0028, 0x1052)

	// RescaleSlope (0028,1053) VR=DS VM=1 Rescale Slope
	RescaleSlope = New(0x0028, 0x1053)

	// RescaleType (0028,1054) VR=LO VM=1 Rescale Type
	RescaleType = New(0x0028, 0x1054)

	// WindowCenterWidthExplanation (0028,1055) VR=LO VM=1-n Window Center &amp; Width Explanation
	WindowCenterWidthExplanation = New(0x0028, 0x1055)

	// VOILUTFunction (0028,1056) VR=CS VM=1 VOI LUT Function
	VOILUTFunction = New(0x0028, 0x1056)

	// GrayScaleRETIRED (0028,1080) VR=CS VM=1 Gray Scale (RETIRED)
	GrayScaleRETIRED = New(0x0028, 0x1080)

	// RecommendedViewingMode (0028,1090) VR=CS VM=1 Recommended Viewing Mode
	RecommendedViewingMode = New(0x0028, 0x1090)

	GrayLookupTableDescriptorRETIRED = New(0x0028, 0x1100)

	RedPaletteColorLookupTableDescriptor = New(0x0028, 0x1101)

	GreenPaletteColorLookupTableDescriptor = New(0x0028, 0x1102)

	BluePaletteColorLookupTableDescriptor = New(0x0028, 0x1103)

	// AlphaPaletteColorLookupTableDescriptor (0028,1104) VR=US VM=3 Alpha Palette Color Lookup Table Descriptor
	AlphaPaletteColorLookupTableDescriptor = New(0x0028, 0x1104)

	LargeRedPaletteColorLookupTableDescriptorRETIRED = New(0x0028, 0x1111)

	LargeGreenPaletteColorLookupTableDescriptorRETIRED = New(0x0028, 0x1112)

	LargeBluePaletteColorLookupTableDescriptorRETIRED = New(0x0028, 0x1113)

	// PaletteColorLookupTableUID (0028,1199) VR=UI VM=1 Palette Color Lookup Table UID
	PaletteColorLookupTableUID = New(0x0028, 0x1199)

	GrayLookupTableDataRETIRED = New(0x0028, 0x1200)

	// RedPaletteColorLookupTableData (0028,1201) VR=OW VM=1 Red Palette Color Lookup Table Data
	RedPaletteColorLookupTableData = New(0x0028, 0x1201)

	// GreenPaletteColorLookupTableData (0028,1202) VR=OW VM=1 Green Palette Color Lookup Table Data
	GreenPaletteColorLookupTableData = New(0x0028, 0x1202)

	// BluePaletteColorLookupTableData (0028,1203) VR=OW VM=1 Blue Palette Color Lookup Table Data
	BluePaletteColorLookupTableData = New(0x0028, 0x1203)

	// AlphaPaletteColorLookupTableData (0028,1204) VR=OW VM=1 Alpha Palette Color Lookup Table Data
	AlphaPaletteColorLookupTableData = New(0x0028, 0x1204)

	// LargeRedPaletteColorLookupTableDataRETIRED (0028,1211) VR=OW VM=1 Large Red Palette Color Lookup Table Data (RETIRED)
	LargeRedPaletteColorLookupTableDataRETIRED = New(0x0028, 0x1211)

	// LargeGreenPaletteColorLookupTableDataRETIRED (0028,1212) VR=OW VM=1 Large Green Palette Color Lookup Table Data (RETIRED)
	LargeGreenPaletteColorLookupTableDataRETIRED = New(0x0028, 0x1212)

	// LargeBluePaletteColorLookupTableDataRETIRED (0028,1213) VR=OW VM=1 Large Blue Palette Color Lookup Table Data (RETIRED)
	LargeBluePaletteColorLookupTableDataRETIRED = New(0x0028, 0x1213)

	// LargePaletteColorLookupTableUIDRETIRED (0028,1214) VR=UI VM=1 Large Palette Color Lookup Table UID (RETIRED)
	LargePaletteColorLookupTableUIDRETIRED = New(0x0028, 0x1214)

	// SegmentedRedPaletteColorLookupTableData (0028,1221) VR=OW VM=1 Segmented Red Palette Color Lookup Table Data
	SegmentedRedPaletteColorLookupTableData = New(0x0028, 0x1221)

	// SegmentedGreenPaletteColorLookupTableData (0028,1222) VR=OW VM=1 Segmented Green Palette Color Lookup Table Data
	SegmentedGreenPaletteColorLookupTableData = New(0x0028, 0x1222)

	// SegmentedBluePaletteColorLookupTableData (0028,1223) VR=OW VM=1 Segmented Blue Palette Color Lookup Table Data
	SegmentedBluePaletteColorLookupTableData = New(0x0028, 0x1223)

	// SegmentedAlphaPaletteColorLookupTableData (0028,1224) VR=OW VM=1 Segmented Alpha Palette Color Lookup Table Data
	SegmentedAlphaPaletteColorLookupTableData = New(0x0028, 0x1224)

	// StoredValueColorRangeSequence (0028,1230) VR=SQ VM=1 Stored Value Color Range Sequence
	StoredValueColorRangeSequence = New(0x0028, 0x1230)

	// MinimumStoredValueMapped (0028,1231) VR=FD VM=1 Minimum Stored Value Mapped
	MinimumStoredValueMapped = New(0x0028, 0x1231)

	// MaximumStoredValueMapped (0028,1232) VR=FD VM=1 Maximum Stored Value Mapped
	MaximumStoredValueMapped = New(0x0028, 0x1232)

	// BreastImplantPresent (0028,1300) VR=CS VM=1 Breast Implant Present
	BreastImplantPresent = New(0x0028, 0x1300)

	// PartialView (0028,1350) VR=CS VM=1 Partial View
	PartialView = New(0x0028, 0x1350)

	// PartialViewDescription (0028,1351) VR=ST VM=1 Partial View Description
	PartialViewDescription = New(0x0028, 0x1351)

	// PartialViewCodeSequence (0028,1352) VR=SQ VM=1 Partial View Code Sequence
	PartialViewCodeSequence = New(0x0028, 0x1352)

	// SpatialLocationsPreserved (0028,135A) VR=CS VM=1 Spatial Locations Preserved
	SpatialLocationsPreserved = New(0x0028, 0x135A)

	// DataFrameAssignmentSequence (0028,1401) VR=SQ VM=1 Data Frame Assignment Sequence
	DataFrameAssignmentSequence = New(0x0028, 0x1401)

	// DataPathAssignment (0028,1402) VR=CS VM=1 Data Path Assignment
	DataPathAssignment = New(0x0028, 0x1402)

	// BitsMappedToColorLookupTable (0028,1403) VR=US VM=1 Bits Mapped to Color Lookup Table
	BitsMappedToColorLookupTable = New(0x0028, 0x1403)

	// BlendingLUT1Sequence (0028,1404) VR=SQ VM=1 Blending LUT 1 Sequence
	BlendingLUT1Sequence = New(0x0028, 0x1404)

	// BlendingLUT1TransferFunction (0028,1405) VR=CS VM=1 Blending LUT 1 Transfer Function
	BlendingLUT1TransferFunction = New(0x0028, 0x1405)

	// BlendingWeightConstant (0028,1406) VR=FD VM=1 Blending Weight Constant
	BlendingWeightConstant = New(0x0028, 0x1406)

	// BlendingLookupTableDescriptor (0028,1407) VR=US VM=3 Blending Lookup Table Descriptor
	BlendingLookupTableDescriptor = New(0x0028, 0x1407)

	// BlendingLookupTableData (0028,1408) VR=OW VM=1 Blending Lookup Table Data
	BlendingLookupTableData = New(0x0028, 0x1408)

	// EnhancedPaletteColorLookupTableSequence (0028,140B) VR=SQ VM=1 Enhanced Palette Color Lookup Table Sequence
	EnhancedPaletteColorLookupTableSequence = New(0x0028, 0x140B)

	// BlendingLUT2Sequence (0028,140C) VR=SQ VM=1 Blending LUT 2 Sequence
	BlendingLUT2Sequence = New(0x0028, 0x140C)

	// BlendingLUT2TransferFunction (0028,140D) VR=CS VM=1 Blending LUT 2 Transfer Function
	BlendingLUT2TransferFunction = New(0x0028, 0x140D)

	// DataPathID (0028,140E) VR=CS VM=1 Data Path ID
	DataPathID = New(0x0028, 0x140E)

	// RGBLUTTransferFunction (0028,140F) VR=CS VM=1 RGB LUT Transfer Function
	RGBLUTTransferFunction = New(0x0028, 0x140F)

	// AlphaLUTTransferFunction (0028,1410) VR=CS VM=1 Alpha LUT Transfer Function
	AlphaLUTTransferFunction = New(0x0028, 0x1410)

	// ICCProfile (0028,2000) VR=OB VM=1 ICC Profile
	ICCProfile = New(0x0028, 0x2000)

	// ColorSpace (0028,2002) VR=CS VM=1 Color Space
	ColorSpace = New(0x0028, 0x2002)

	// LossyImageCompression (0028,2110) VR=CS VM=1 Lossy Image Compression
	LossyImageCompression = New(0x0028, 0x2110)

	// LossyImageCompressionRatio (0028,2112) VR=DS VM=1-n Lossy Image Compression Ratio
	LossyImageCompressionRatio = New(0x0028, 0x2112)

	// LossyImageCompressionMethod (0028,2114) VR=CS VM=1-n Lossy Image Compression Method
	LossyImageCompressionMethod = New(0x0028, 0x2114)

	// ModalityLUTSequence (0028,3000) VR=SQ VM=1 Modality LUT Sequence
	ModalityLUTSequence = New(0x0028, 0x3000)

	// VariableModalityLUTSequence (0028,3001) VR=SQ VM=1 Variable Modality LUT Sequence
	VariableModalityLUTSequence = New(0x0028, 0x3001)

	LUTDescriptor = New(0x0028, 0x3002)

	// LUTExplanation (0028,3003) VR=LO VM=1 LUT Explanation
	LUTExplanation = New(0x0028, 0x3003)

	// ModalityLUTType (0028,3004) VR=LO VM=1 Modality LUT Type
	ModalityLUTType = New(0x0028, 0x3004)

	LUTData = New(0x0028, 0x3006)

	// VOILUTSequence (0028,3010) VR=SQ VM=1 VOI LUT Sequence
	VOILUTSequence = New(0x0028, 0x3010)

	// SoftcopyVOILUTSequence (0028,3110) VR=SQ VM=1 Softcopy VOI LUT Sequence
	SoftcopyVOILUTSequence = New(0x0028, 0x3110)

	// ImagePresentationCommentsRETIRED (0028,4000) VR=LT VM=1 Image Presentation Comments (RETIRED)
	ImagePresentationCommentsRETIRED = New(0x0028, 0x4000)

	// BiPlaneAcquisitionSequenceRETIRED (0028,5000) VR=SQ VM=1 Bi-Plane Acquisition Sequence (RETIRED)
	BiPlaneAcquisitionSequenceRETIRED = New(0x0028, 0x5000)

	// RepresentativeFrameNumber (0028,6010) VR=US VM=1 Representative Frame Number
	RepresentativeFrameNumber = New(0x0028, 0x6010)

	// FrameNumbersOfInterest (0028,6020) VR=US VM=1-n Frame Numbers of Interest (FOI)
	FrameNumbersOfInterest = New(0x0028, 0x6020)

	// FrameOfInterestDescription (0028,6022) VR=LO VM=1-n Frame of Interest Description
	FrameOfInterestDescription = New(0x0028, 0x6022)

	// FrameOfInterestType (0028,6023) VR=CS VM=1-n Frame of Interest Type
	FrameOfInterestType = New(0x0028, 0x6023)

	// MaskPointersRETIRED (0028,6030) VR=US VM=1-n Mask Pointer(s) (RETIRED)
	MaskPointersRETIRED = New(0x0028, 0x6030)

	// RWavePointer (0028,6040) VR=US VM=1-n R Wave Pointer
	RWavePointer = New(0x0028, 0x6040)

	// MaskSubtractionSequence (0028,6100) VR=SQ VM=1 Mask Subtraction Sequence
	MaskSubtractionSequence = New(0x0028, 0x6100)

	// MaskOperation (0028,6101) VR=CS VM=1 Mask Operation
	MaskOperation = New(0x0028, 0x6101)

	// ApplicableFrameRange (0028,6102) VR=US VM=2-2n Applicable Frame Range
	ApplicableFrameRange = New(0x0028, 0x6102)

	// MaskFrameNumbers (0028,6110) VR=US VM=1-n Mask Frame Numbers
	MaskFrameNumbers = New(0x0028, 0x6110)

	// ContrastFrameAveraging (0028,6112) VR=US VM=1 Contrast Frame Averaging
	ContrastFrameAveraging = New(0x0028, 0x6112)

	// MaskSubPixelShift (0028,6114) VR=FL VM=2 Mask Sub-pixel Shift
	MaskSubPixelShift = New(0x0028, 0x6114)

	// TIDOffset (0028,6120) VR=SS VM=1 TID Offset
	TIDOffset = New(0x0028, 0x6120)

	// MaskOperationExplanation (0028,6190) VR=ST VM=1 Mask Operation Explanation
	MaskOperationExplanation = New(0x0028, 0x6190)

	// EquipmentAdministratorSequence (0028,7000) VR=SQ VM=1 Equipment Administrator Sequence
	EquipmentAdministratorSequence = New(0x0028, 0x7000)

	// NumberOfDisplaySubsystems (0028,7001) VR=US VM=1 Number of Display Subsystems
	NumberOfDisplaySubsystems = New(0x0028, 0x7001)

	// CurrentConfigurationID (0028,7002) VR=US VM=1 Current Configuration ID
	CurrentConfigurationID = New(0x0028, 0x7002)

	// DisplaySubsystemID (0028,7003) VR=US VM=1 Display Subsystem ID
	DisplaySubsystemID = New(0x0028, 0x7003)

	// DisplaySubsystemName (0028,7004) VR=SH VM=1 Display Subsystem Name
	DisplaySubsystemName = New(0x0028, 0x7004)

	// DisplaySubsystemDescription (0028,7005) VR=LO VM=1 Display Subsystem Description
	DisplaySubsystemDescription = New(0x0028, 0x7005)

	// SystemStatus (0028,7006) VR=CS VM=1 System Status
	SystemStatus = New(0x0028, 0x7006)

	// SystemStatusComment (0028,7007) VR=LO VM=1 System Status Comment
	SystemStatusComment = New(0x0028, 0x7007)

	// TargetLuminanceCharacteristicsSequence (0028,7008) VR=SQ VM=1 Target Luminance Characteristics Sequence
	TargetLuminanceCharacteristicsSequence = New(0x0028, 0x7008)

	// LuminanceCharacteristicsID (0028,7009) VR=US VM=1 Luminance Characteristics ID
	LuminanceCharacteristicsID = New(0x0028, 0x7009)

	// DisplaySubsystemConfigurationSequence (0028,700A) VR=SQ VM=1 Display Subsystem Configuration Sequence
	DisplaySubsystemConfigurationSequence = New(0x0028, 0x700A)

	// ConfigurationID (0028,700B) VR=US VM=1 Configuration ID
	ConfigurationID = New(0x0028, 0x700B)

	// ConfigurationName (0028,700C) VR=SH VM=1 Configuration Name
	ConfigurationName = New(0x0028, 0x700C)

	// ConfigurationDescription (0028,700D) VR=LO VM=1 Configuration Description
	ConfigurationDescription = New(0x0028, 0x700D)

	// ReferencedTargetLuminanceCharacteristicsID (0028,700E) VR=US VM=1 Referenced Target Luminance Characteristics ID
	ReferencedTargetLuminanceCharacteristicsID = New(0x0028, 0x700E)

	// QAResultsSequence (0028,700F) VR=SQ VM=1 QA Results Sequence
	QAResultsSequence = New(0x0028, 0x700F)

	// DisplaySubsystemQAResultsSequence (0028,7010) VR=SQ VM=1 Display Subsystem QA Results Sequence
	DisplaySubsystemQAResultsSequence = New(0x0028, 0x7010)

	// ConfigurationQAResultsSequence (0028,7011) VR=SQ VM=1 Configuration QA Results Sequence
	ConfigurationQAResultsSequence = New(0x0028, 0x7011)

	// MeasurementEquipmentSequence (0028,7012) VR=SQ VM=1 Measurement Equipment Sequence
	MeasurementEquipmentSequence = New(0x0028, 0x7012)

	// MeasurementFunctions (0028,7013) VR=CS VM=1-n Measurement Functions
	MeasurementFunctions = New(0x0028, 0x7013)

	// MeasurementEquipmentType (0028,7014) VR=CS VM=1 Measurement Equipment Type
	MeasurementEquipmentType = New(0x0028, 0x7014)

	// VisualEvaluationResultSequence (0028,7015) VR=SQ VM=1 Visual Evaluation Result Sequence
	VisualEvaluationResultSequence = New(0x0028, 0x7015)

	// DisplayCalibrationResultSequence (0028,7016) VR=SQ VM=1 Display Calibration Result Sequence
	DisplayCalibrationResultSequence = New(0x0028, 0x7016)

	// DDLValue (0028,7017) VR=US VM=1 DDL Value
	DDLValue = New(0x0028, 0x7017)

	// CIExyWhitePoint (0028,7018) VR=FL VM=2 CIExy White Point
	CIExyWhitePoint = New(0x0028, 0x7018)

	// DisplayFunctionType (0028,7019) VR=CS VM=1 Display Function Type
	DisplayFunctionType = New(0x0028, 0x7019)

	// GammaValue (0028,701A) VR=FL VM=1 Gamma Value
	GammaValue = New(0x0028, 0x701A)

	// NumberOfLuminancePoints (0028,701B) VR=US VM=1 Number of Luminance Points
	NumberOfLuminancePoints = New(0x0028, 0x701B)

	// LuminanceResponseSequence (0028,701C) VR=SQ VM=1 Luminance Response Sequence
	LuminanceResponseSequence = New(0x0028, 0x701C)

	// TargetMinimumLuminance (0028,701D) VR=FL VM=1 Target Minimum Luminance
	TargetMinimumLuminance = New(0x0028, 0x701D)

	// TargetMaximumLuminance (0028,701E) VR=FL VM=1 Target Maximum Luminance
	TargetMaximumLuminance = New(0x0028, 0x701E)

	// LuminanceValue (0028,701F) VR=FL VM=1 Luminance Value
	LuminanceValue = New(0x0028, 0x701F)

	// LuminanceResponseDescription (0028,7020) VR=LO VM=1 Luminance Response Description
	LuminanceResponseDescription = New(0x0028, 0x7020)

	// WhitePointFlag (0028,7021) VR=CS VM=1 White Point Flag
	WhitePointFlag = New(0x0028, 0x7021)

	// DisplayDeviceTypeCodeSequence (0028,7022) VR=SQ VM=1 Display Device Type Code Sequence
	DisplayDeviceTypeCodeSequence = New(0x0028, 0x7022)

	// DisplaySubsystemSequence (0028,7023) VR=SQ VM=1 Display Subsystem Sequence
	DisplaySubsystemSequence = New(0x0028, 0x7023)

	// LuminanceResultSequence (0028,7024) VR=SQ VM=1 Luminance Result Sequence
	LuminanceResultSequence = New(0x0028, 0x7024)

	// AmbientLightValueSource (0028,7025) VR=CS VM=1 Ambient Light Value Source
	AmbientLightValueSource = New(0x0028, 0x7025)

	// MeasuredCharacteristics (0028,7026) VR=CS VM=1-n Measured Characteristics
	MeasuredCharacteristics = New(0x0028, 0x7026)

	// LuminanceUniformityResultSequence (0028,7027) VR=SQ VM=1 Luminance Uniformity Result Sequence
	LuminanceUniformityResultSequence = New(0x0028, 0x7027)

	// VisualEvaluationTestSequence (0028,7028) VR=SQ VM=1 Visual Evaluation Test Sequence
	VisualEvaluationTestSequence = New(0x0028, 0x7028)

	// TestResult (0028,7029) VR=CS VM=1 Test Result
	TestResult = New(0x0028, 0x7029)

	// TestResultComment (0028,702A) VR=LO VM=1 Test Result Comment
	TestResultComment = New(0x0028, 0x702A)

	// TestImageValidation (0028,702B) VR=CS VM=1 Test Image Validation
	TestImageValidation = New(0x0028, 0x702B)

	// TestPatternCodeSequence (0028,702C) VR=SQ VM=1 Test Pattern Code Sequence
	TestPatternCodeSequence = New(0x0028, 0x702C)

	// MeasurementPatternCodeSequence (0028,702D) VR=SQ VM=1 Measurement Pattern Code Sequence
	MeasurementPatternCodeSequence = New(0x0028, 0x702D)

	// VisualEvaluationMethodCodeSequence (0028,702E) VR=SQ VM=1 Visual Evaluation Method Code Sequence
	VisualEvaluationMethodCodeSequence = New(0x0028, 0x702E)

	// PixelDataProviderURL (0028,7FE0) VR=UR VM=1 Pixel Data Provider URL
	PixelDataProviderURL = New(0x0028, 0x7FE0)

	// DataPointRows (0028,9001) VR=UL VM=1 Data Point Rows
	DataPointRows = New(0x0028, 0x9001)

	// DataPointColumns (0028,9002) VR=UL VM=1 Data Point Columns
	DataPointColumns = New(0x0028, 0x9002)

	// SignalDomainColumns (0028,9003) VR=CS VM=1 Signal Domain Columns
	SignalDomainColumns = New(0x0028, 0x9003)

	// LargestMonochromePixelValueRETIRED (0028,9099) VR=US VM=1 Largest Monochrome Pixel Value (RETIRED)
	LargestMonochromePixelValueRETIRED = New(0x0028, 0x9099)

	// DataRepresentation (0028,9108) VR=CS VM=1 Data Representation
	DataRepresentation = New(0x0028, 0x9108)

	// PixelMeasuresSequence (0028,9110) VR=SQ VM=1 Pixel Measures Sequence
	PixelMeasuresSequence = New(0x0028, 0x9110)

	// FrameVOILUTSequence (0028,9132) VR=SQ VM=1 Frame VOI LUT Sequence
	FrameVOILUTSequence = New(0x0028, 0x9132)

	// PixelValueTransformationSequence (0028,9145) VR=SQ VM=1 Pixel Value Transformation Sequence
	PixelValueTransformationSequence = New(0x0028, 0x9145)

	// SignalDomainRows (0028,9235) VR=CS VM=1 Signal Domain Rows
	SignalDomainRows = New(0x0028, 0x9235)

	// DisplayFilterPercentage (0028,9411) VR=FL VM=1 Display Filter Percentage
	DisplayFilterPercentage = New(0x0028, 0x9411)

	// FramePixelShiftSequence (0028,9415) VR=SQ VM=1 Frame Pixel Shift Sequence
	FramePixelShiftSequence = New(0x0028, 0x9415)

	// SubtractionItemID (0028,9416) VR=US VM=1 Subtraction Item ID
	SubtractionItemID = New(0x0028, 0x9416)

	// PixelIntensityRelationshipLUTSequence (0028,9422) VR=SQ VM=1 Pixel Intensity Relationship LUT Sequence
	PixelIntensityRelationshipLUTSequence = New(0x0028, 0x9422)

	// FramePixelDataPropertiesSequence (0028,9443) VR=SQ VM=1 Frame Pixel Data Properties Sequence
	FramePixelDataPropertiesSequence = New(0x0028, 0x9443)

	// GeometricalProperties (0028,9444) VR=CS VM=1 Geometrical Properties
	GeometricalProperties = New(0x0028, 0x9444)

	// GeometricMaximumDistortion (0028,9445) VR=FL VM=1 Geometric Maximum Distortion
	GeometricMaximumDistortion = New(0x0028, 0x9445)

	// ImageProcessingApplied (0028,9446) VR=CS VM=1-n Image Processing Applied
	ImageProcessingApplied = New(0x0028, 0x9446)

	// MaskSelectionMode (0028,9454) VR=CS VM=1 Mask Selection Mode
	MaskSelectionMode = New(0x0028, 0x9454)

	// LUTFunction (0028,9474) VR=CS VM=1 LUT Function
	LUTFunction = New(0x0028, 0x9474)

	// MaskVisibilityPercentage (0028,9478) VR=FL VM=1 Mask Visibility Percentage
	MaskVisibilityPercentage = New(0x0028, 0x9478)

	// PixelShiftSequence (0028,9501) VR=SQ VM=1 Pixel Shift Sequence
	PixelShiftSequence = New(0x0028, 0x9501)

	// RegionPixelShiftSequence (0028,9502) VR=SQ VM=1 Region Pixel Shift Sequence
	RegionPixelShiftSequence = New(0x0028, 0x9502)

	// VerticesOfTheRegion (0028,9503) VR=SS VM=2-2n Vertices of the Region
	VerticesOfTheRegion = New(0x0028, 0x9503)

	// MultiFramePresentationSequence (0028,9505) VR=SQ VM=1 Multi-frame Presentation Sequence
	MultiFramePresentationSequence = New(0x0028, 0x9505)

	// PixelShiftFrameRange (0028,9506) VR=US VM=2-2n Pixel Shift Frame Range
	PixelShiftFrameRange = New(0x0028, 0x9506)

	// LUTFrameRange (0028,9507) VR=US VM=2-2n LUT Frame Range
	LUTFrameRange = New(0x0028, 0x9507)

	// ImageToEquipmentMappingMatrix (0028,9520) VR=DS VM=16 Image to Equipment Mapping Matrix
	ImageToEquipmentMappingMatrix = New(0x0028, 0x9520)

	// EquipmentCoordinateSystemIdentification (0028,9537) VR=CS VM=1 Equipment Coordinate System Identification
	EquipmentCoordinateSystemIdentification = New(0x0028, 0x9537)

	// StudyStatusIDRETIRED (0032,000A) VR=CS VM=1 Study Status ID (RETIRED)
	StudyStatusIDRETIRED = New(0x0032, 0x000A)

	// StudyPriorityIDRETIRED (0032,000C) VR=CS VM=1 Study Priority ID (RETIRED)
	StudyPriorityIDRETIRED = New(0x0032, 0x000C)

	// StudyIDIssuerRETIRED (0032,0012) VR=LO VM=1 Study ID Issuer (RETIRED)
	StudyIDIssuerRETIRED = New(0x0032, 0x0012)

	// StudyVerifiedDateRETIRED (0032,0032) VR=DA VM=1 Study Verified Date (RETIRED)
	StudyVerifiedDateRETIRED = New(0x0032, 0x0032)

	// StudyVerifiedTimeRETIRED (0032,0033) VR=TM VM=1 Study Verified Time (RETIRED)
	StudyVerifiedTimeRETIRED = New(0x0032, 0x0033)

	// StudyReadDateRETIRED (0032,0034) VR=DA VM=1 Study Read Date (RETIRED)
	StudyReadDateRETIRED = New(0x0032, 0x0034)

	// StudyReadTimeRETIRED (0032,0035) VR=TM VM=1 Study Read Time (RETIRED)
	StudyReadTimeRETIRED = New(0x0032, 0x0035)

	// ScheduledStudyStartDateRETIRED (0032,1000) VR=DA VM=1 Scheduled Study Start Date (RETIRED)
	ScheduledStudyStartDateRETIRED = New(0x0032, 0x1000)

	// ScheduledStudyStartTimeRETIRED (0032,1001) VR=TM VM=1 Scheduled Study Start Time (RETIRED)
	ScheduledStudyStartTimeRETIRED = New(0x0032, 0x1001)

	// ScheduledStudyStopDateRETIRED (0032,1010) VR=DA VM=1 Scheduled Study Stop Date (RETIRED)
	ScheduledStudyStopDateRETIRED = New(0x0032, 0x1010)

	// ScheduledStudyStopTimeRETIRED (0032,1011) VR=TM VM=1 Scheduled Study Stop Time (RETIRED)
	ScheduledStudyStopTimeRETIRED = New(0x0032, 0x1011)

	// ScheduledStudyLocationRETIRED (0032,1020) VR=LO VM=1 Scheduled Study Location (RETIRED)
	ScheduledStudyLocationRETIRED = New(0x0032, 0x1020)

	// ScheduledStudyLocationAETitleRETIRED (0032,1021) VR=AE VM=1-n Scheduled Study Location AE Title (RETIRED)
	ScheduledStudyLocationAETitleRETIRED = New(0x0032, 0x1021)

	// ReasonForStudyRETIRED (0032,1030) VR=LO VM=1 Reason for Study (RETIRED)
	ReasonForStudyRETIRED = New(0x0032, 0x1030)

	// RequestingPhysicianIdentificationSequence (0032,1031) VR=SQ VM=1 Requesting Physician Identification Sequence
	RequestingPhysicianIdentificationSequence = New(0x0032, 0x1031)

	// RequestingPhysician (0032,1032) VR=PN VM=1 Requesting Physician
	RequestingPhysician = New(0x0032, 0x1032)

	// RequestingService (0032,1033) VR=LO VM=1 Requesting Service
	RequestingService = New(0x0032, 0x1033)

	// RequestingServiceCodeSequence (0032,1034) VR=SQ VM=1 Requesting Service Code Sequence
	RequestingServiceCodeSequence = New(0x0032, 0x1034)

	// StudyArrivalDateRETIRED (0032,1040) VR=DA VM=1 Study Arrival Date (RETIRED)
	StudyArrivalDateRETIRED = New(0x0032, 0x1040)

	// StudyArrivalTimeRETIRED (0032,1041) VR=TM VM=1 Study Arrival Time (RETIRED)
	StudyArrivalTimeRETIRED = New(0x0032, 0x1041)

	// StudyCompletionDateRETIRED (0032,1050) VR=DA VM=1 Study Completion Date (RETIRED)
	StudyCompletionDateRETIRED = New(0x0032, 0x1050)

	// StudyCompletionTimeRETIRED (0032,1051) VR=TM VM=1 Study Completion Time (RETIRED)
	StudyCompletionTimeRETIRED = New(0x0032, 0x1051)

	// StudyComponentStatusIDRETIRED (0032,1055) VR=CS VM=1 Study Component Status ID (RETIRED)
	StudyComponentStatusIDRETIRED = New(0x0032, 0x1055)

	// RequestedProcedureDescription (0032,1060) VR=LO VM=1 Requested Procedure Description
	RequestedProcedureDescription = New(0x0032, 0x1060)

	// RequestedProcedureCodeSequence (0032,1064) VR=SQ VM=1 Requested Procedure Code Sequence
	RequestedProcedureCodeSequence = New(0x0032, 0x1064)

	// RequestedLateralityCodeSequence (0032,1065) VR=SQ VM=1 Requested Laterality Code Sequence
	RequestedLateralityCodeSequence = New(0x0032, 0x1065)

	// ReasonForVisit (0032,1066) VR=UT VM=1 Reason for Visit
	ReasonForVisit = New(0x0032, 0x1066)

	// ReasonForVisitCodeSequence (0032,1067) VR=SQ VM=1 Reason for Visit Code Sequence
	ReasonForVisitCodeSequence = New(0x0032, 0x1067)

	// RequestedContrastAgent (0032,1070) VR=LO VM=1 Requested Contrast Agent
	RequestedContrastAgent = New(0x0032, 0x1070)

	// StudyCommentsRETIRED (0032,4000) VR=LT VM=1 Study Comments (RETIRED)
	StudyCommentsRETIRED = New(0x0032, 0x4000)

	// FlowIdentifierSequence (0034,0001) VR=SQ VM=1 Flow Identifier Sequence
	FlowIdentifierSequence = New(0x0034, 0x0001)

	// FlowIdentifier (0034,0002) VR=OB VM=1 Flow Identifier
	FlowIdentifier = New(0x0034, 0x0002)

	// FlowTransferSyntaxUID (0034,0003) VR=UI VM=1 Flow Transfer Syntax UID
	FlowTransferSyntaxUID = New(0x0034, 0x0003)

	// FlowRTPSamplingRate (0034,0004) VR=UL VM=1 Flow RTP Sampling Rate
	FlowRTPSamplingRate = New(0x0034, 0x0004)

	// SourceIdentifier (0034,0005) VR=OB VM=1 Source Identifier
	SourceIdentifier = New(0x0034, 0x0005)

	// FrameOriginTimestamp (0034,0007) VR=OB VM=1 Frame Origin Timestamp
	FrameOriginTimestamp = New(0x0034, 0x0007)

	// IncludesImagingSubject (0034,0008) VR=CS VM=1 Includes Imaging Subject
	IncludesImagingSubject = New(0x0034, 0x0008)

	// FrameUsefulnessGroupSequence (0034,0009) VR=SQ VM=1 Frame Usefulness Group Sequence
	FrameUsefulnessGroupSequence = New(0x0034, 0x0009)

	// RealTimeBulkDataFlowSequence (0034,000A) VR=SQ VM=1 Real-Time Bulk Data Flow Sequence
	RealTimeBulkDataFlowSequence = New(0x0034, 0x000A)

	// CameraPositionGroupSequence (0034,000B) VR=SQ VM=1 Camera Position Group Sequence
	CameraPositionGroupSequence = New(0x0034, 0x000B)

	// IncludesInformation (0034,000C) VR=CS VM=1 Includes Information
	IncludesInformation = New(0x0034, 0x000C)

	// TimeOfFrameGroupSequence (0034,000D) VR=SQ VM=1 Time of Frame Group Sequence
	TimeOfFrameGroupSequence = New(0x0034, 0x000D)

	// ReferencedPatientAliasSequenceRETIRED (0038,0004) VR=SQ VM=1 Referenced Patient Alias Sequence (RETIRED)
	ReferencedPatientAliasSequenceRETIRED = New(0x0038, 0x0004)

	// VisitStatusID (0038,0008) VR=CS VM=1 Visit Status ID
	VisitStatusID = New(0x0038, 0x0008)

	// AdmissionID (0038,0010) VR=LO VM=1 Admission ID
	AdmissionID = New(0x0038, 0x0010)

	// IssuerOfAdmissionIDRETIRED (0038,0011) VR=LO VM=1 Issuer of Admission ID (RETIRED)
	IssuerOfAdmissionIDRETIRED = New(0x0038, 0x0011)

	// IssuerOfAdmissionIDSequence (0038,0014) VR=SQ VM=1 Issuer of Admission ID Sequence
	IssuerOfAdmissionIDSequence = New(0x0038, 0x0014)

	// RouteOfAdmissions (0038,0016) VR=LO VM=1 Route of Admissions
	RouteOfAdmissions = New(0x0038, 0x0016)

	// ScheduledAdmissionDateRETIRED (0038,001A) VR=DA VM=1 Scheduled Admission Date (RETIRED)
	ScheduledAdmissionDateRETIRED = New(0x0038, 0x001A)

	// ScheduledAdmissionTimeRETIRED (0038,001B) VR=TM VM=1 Scheduled Admission Time (RETIRED)
	ScheduledAdmissionTimeRETIRED = New(0x0038, 0x001B)

	// ScheduledDischargeDateRETIRED (0038,001C) VR=DA VM=1 Scheduled Discharge Date (RETIRED)
	ScheduledDischargeDateRETIRED = New(0x0038, 0x001C)

	// ScheduledDischargeTimeRETIRED (0038,001D) VR=TM VM=1 Scheduled Discharge Time (RETIRED)
	ScheduledDischargeTimeRETIRED = New(0x0038, 0x001D)

	// ScheduledPatientInstitutionResidenceRETIRED (0038,001E) VR=LO VM=1 Scheduled Patient Institution Residence (RETIRED)
	ScheduledPatientInstitutionResidenceRETIRED = New(0x0038, 0x001E)

	// AdmittingDate (0038,0020) VR=DA VM=1 Admitting Date
	AdmittingDate = New(0x0038, 0x0020)

	// AdmittingTime (0038,0021) VR=TM VM=1 Admitting Time
	AdmittingTime = New(0x0038, 0x0021)

	// DischargeDateRETIRED (0038,0030) VR=DA VM=1 Discharge Date (RETIRED)
	DischargeDateRETIRED = New(0x0038, 0x0030)

	// DischargeTimeRETIRED (0038,0032) VR=TM VM=1 Discharge Time (RETIRED)
	DischargeTimeRETIRED = New(0x0038, 0x0032)

	// DischargeDiagnosisDescriptionRETIRED (0038,0040) VR=LO VM=1 Discharge Diagnosis Description (RETIRED)
	DischargeDiagnosisDescriptionRETIRED = New(0x0038, 0x0040)

	// DischargeDiagnosisCodeSequenceRETIRED (0038,0044) VR=SQ VM=1 Discharge Diagnosis Code Sequence (RETIRED)
	DischargeDiagnosisCodeSequenceRETIRED = New(0x0038, 0x0044)

	// SpecialNeeds (0038,0050) VR=LO VM=1 Special Needs
	SpecialNeeds = New(0x0038, 0x0050)

	// ServiceEpisodeID (0038,0060) VR=LO VM=1 Service Episode ID
	ServiceEpisodeID = New(0x0038, 0x0060)

	// IssuerOfServiceEpisodeIDRETIRED (0038,0061) VR=LO VM=1 Issuer of Service Episode ID (RETIRED)
	IssuerOfServiceEpisodeIDRETIRED = New(0x0038, 0x0061)

	// ServiceEpisodeDescription (0038,0062) VR=LO VM=1 Service Episode Description
	ServiceEpisodeDescription = New(0x0038, 0x0062)

	// IssuerOfServiceEpisodeIDSequence (0038,0064) VR=SQ VM=1 Issuer of Service Episode ID Sequence
	IssuerOfServiceEpisodeIDSequence = New(0x0038, 0x0064)

	// PertinentDocumentsSequence (0038,0100) VR=SQ VM=1 Pertinent Documents Sequence
	PertinentDocumentsSequence = New(0x0038, 0x0100)

	// PertinentResourcesSequence (0038,0101) VR=SQ VM=1 Pertinent Resources Sequence
	PertinentResourcesSequence = New(0x0038, 0x0101)

	// ResourceDescription (0038,0102) VR=LO VM=1 Resource Description
	ResourceDescription = New(0x0038, 0x0102)

	// CurrentPatientLocation (0038,0300) VR=LO VM=1 Current Patient Location
	CurrentPatientLocation = New(0x0038, 0x0300)

	// PatientInstitutionResidence (0038,0400) VR=LO VM=1 Patient's Institution Residence
	PatientInstitutionResidence = New(0x0038, 0x0400)

	// PatientState (0038,0500) VR=LO VM=1 Patient State
	PatientState = New(0x0038, 0x0500)

	// PatientClinicalTrialParticipationSequence (0038,0502) VR=SQ VM=1 Patient Clinical Trial Participation Sequence
	PatientClinicalTrialParticipationSequence = New(0x0038, 0x0502)

	// VisitComments (0038,4000) VR=LT VM=1 Visit Comments
	VisitComments = New(0x0038, 0x4000)

	// WaveformOriginality (003A,0004) VR=CS VM=1 Waveform Originality
	WaveformOriginality = New(0x003A, 0x0004)

	// NumberOfWaveformChannels (003A,0005) VR=US VM=1 Number of Waveform Channels
	NumberOfWaveformChannels = New(0x003A, 0x0005)

	// NumberOfWaveformSamples (003A,0010) VR=UL VM=1 Number of Waveform Samples
	NumberOfWaveformSamples = New(0x003A, 0x0010)

	// SamplingFrequency (003A,001A) VR=DS VM=1 Sampling Frequency
	SamplingFrequency = New(0x003A, 0x001A)

	// MultiplexGroupLabel (003A,0020) VR=SH VM=1 Multiplex Group Label
	MultiplexGroupLabel = New(0x003A, 0x0020)

	// ChannelDefinitionSequence (003A,0200) VR=SQ VM=1 Channel Definition Sequence
	ChannelDefinitionSequence = New(0x003A, 0x0200)

	// WaveformChannelNumber (003A,0202) VR=IS VM=1 Waveform Channel Number
	WaveformChannelNumber = New(0x003A, 0x0202)

	// ChannelLabel (003A,0203) VR=SH VM=1 Channel Label
	ChannelLabel = New(0x003A, 0x0203)

	// ChannelStatus (003A,0205) VR=CS VM=1-n Channel Status
	ChannelStatus = New(0x003A, 0x0205)

	// ChannelSourceSequence (003A,0208) VR=SQ VM=1 Channel Source Sequence
	ChannelSourceSequence = New(0x003A, 0x0208)

	// ChannelSourceModifiersSequence (003A,0209) VR=SQ VM=1 Channel Source Modifiers Sequence
	ChannelSourceModifiersSequence = New(0x003A, 0x0209)

	// SourceWaveformSequence (003A,020A) VR=SQ VM=1 Source Waveform Sequence
	SourceWaveformSequence = New(0x003A, 0x020A)

	// ChannelDerivationDescription (003A,020C) VR=LO VM=1 Channel Derivation Description
	ChannelDerivationDescription = New(0x003A, 0x020C)

	// ChannelSensitivity (003A,0210) VR=DS VM=1 Channel Sensitivity
	ChannelSensitivity = New(0x003A, 0x0210)

	// ChannelSensitivityUnitsSequence (003A,0211) VR=SQ VM=1 Channel Sensitivity Units Sequence
	ChannelSensitivityUnitsSequence = New(0x003A, 0x0211)

	// ChannelSensitivityCorrectionFactor (003A,0212) VR=DS VM=1 Channel Sensitivity Correction Factor
	ChannelSensitivityCorrectionFactor = New(0x003A, 0x0212)

	// ChannelBaseline (003A,0213) VR=DS VM=1 Channel Baseline
	ChannelBaseline = New(0x003A, 0x0213)

	// ChannelTimeSkew (003A,0214) VR=DS VM=1 Channel Time Skew
	ChannelTimeSkew = New(0x003A, 0x0214)

	// ChannelSampleSkew (003A,0215) VR=DS VM=1 Channel Sample Skew
	ChannelSampleSkew = New(0x003A, 0x0215)

	// ChannelOffset (003A,0218) VR=DS VM=1 Channel Offset
	ChannelOffset = New(0x003A, 0x0218)

	// WaveformBitsStored (003A,021A) VR=US VM=1 Waveform Bits Stored
	WaveformBitsStored = New(0x003A, 0x021A)

	// FilterLowFrequency (003A,0220) VR=DS VM=1 Filter Low Frequency
	FilterLowFrequency = New(0x003A, 0x0220)

	// FilterHighFrequency (003A,0221) VR=DS VM=1 Filter High Frequency
	FilterHighFrequency = New(0x003A, 0x0221)

	// NotchFilterFrequency (003A,0222) VR=DS VM=1 Notch Filter Frequency
	NotchFilterFrequency = New(0x003A, 0x0222)

	// NotchFilterBandwidth (003A,0223) VR=DS VM=1 Notch Filter Bandwidth
	NotchFilterBandwidth = New(0x003A, 0x0223)

	// WaveformDataDisplayScale (003A,0230) VR=FL VM=1 Waveform Data Display Scale
	WaveformDataDisplayScale = New(0x003A, 0x0230)

	// WaveformDisplayBackgroundCIELabValue (003A,0231) VR=US VM=3 Waveform Display Background CIELab Value
	WaveformDisplayBackgroundCIELabValue = New(0x003A, 0x0231)

	// WaveformPresentationGroupSequence (003A,0240) VR=SQ VM=1 Waveform Presentation Group Sequence
	WaveformPresentationGroupSequence = New(0x003A, 0x0240)

	// PresentationGroupNumber (003A,0241) VR=US VM=1 Presentation Group Number
	PresentationGroupNumber = New(0x003A, 0x0241)

	// ChannelDisplaySequence (003A,0242) VR=SQ VM=1 Channel Display Sequence
	ChannelDisplaySequence = New(0x003A, 0x0242)

	// ChannelRecommendedDisplayCIELabValue (003A,0244) VR=US VM=3 Channel Recommended Display CIELab Value
	ChannelRecommendedDisplayCIELabValue = New(0x003A, 0x0244)

	// ChannelPosition (003A,0245) VR=FL VM=1 Channel Position
	ChannelPosition = New(0x003A, 0x0245)

	// DisplayShadingFlag (003A,0246) VR=CS VM=1 Display Shading Flag
	DisplayShadingFlag = New(0x003A, 0x0246)

	// FractionalChannelDisplayScale (003A,0247) VR=FL VM=1 Fractional Channel Display Scale
	FractionalChannelDisplayScale = New(0x003A, 0x0247)

	// AbsoluteChannelDisplayScale (003A,0248) VR=FL VM=1 Absolute Channel Display Scale
	AbsoluteChannelDisplayScale = New(0x003A, 0x0248)

	// MultiplexedAudioChannelsDescriptionCodeSequence (003A,0300) VR=SQ VM=1 Multiplexed Audio Channels Description Code Sequence
	MultiplexedAudioChannelsDescriptionCodeSequence = New(0x003A, 0x0300)

	// ChannelIdentificationCode (003A,0301) VR=IS VM=1 Channel Identification Code
	ChannelIdentificationCode = New(0x003A, 0x0301)

	// ChannelMode (003A,0302) VR=CS VM=1 Channel Mode
	ChannelMode = New(0x003A, 0x0302)

	// MultiplexGroupUID (003A,0310) VR=UI VM=1 Multiplex Group UID
	MultiplexGroupUID = New(0x003A, 0x0310)

	// PowerlineFrequency (003A,0311) VR=DS VM=1 Powerline Frequency
	PowerlineFrequency = New(0x003A, 0x0311)

	// ChannelImpedanceSequence (003A,0312) VR=SQ VM=1 Channel Impedance Sequence
	ChannelImpedanceSequence = New(0x003A, 0x0312)

	// ImpedanceValue (003A,0313) VR=DS VM=1 Impedance Value
	ImpedanceValue = New(0x003A, 0x0313)

	// ImpedanceMeasurementDateTime (003A,0314) VR=DT VM=1 Impedance Measurement DateTime
	ImpedanceMeasurementDateTime = New(0x003A, 0x0314)

	// ImpedanceMeasurementFrequency (003A,0315) VR=DS VM=1 Impedance Measurement Frequency
	ImpedanceMeasurementFrequency = New(0x003A, 0x0315)

	// ImpedanceMeasurementCurrentType (003A,0316) VR=CS VM=1 Impedance Measurement Current Type
	ImpedanceMeasurementCurrentType = New(0x003A, 0x0316)

	// WaveformAmplifierType (003A,0317) VR=CS VM=1 Waveform Amplifier Type
	WaveformAmplifierType = New(0x003A, 0x0317)

	// FilterLowFrequencyCharacteristicsSequence (003A,0318) VR=SQ VM=1 Filter Low Frequency Characteristics Sequence
	FilterLowFrequencyCharacteristicsSequence = New(0x003A, 0x0318)

	// FilterHighFrequencyCharacteristicsSequence (003A,0319) VR=SQ VM=1 Filter High Frequency Characteristics Sequence
	FilterHighFrequencyCharacteristicsSequence = New(0x003A, 0x0319)

	// SummarizedFilterLookupTableSequence (003A,0320) VR=SQ VM=1 Summarized Filter Lookup Table Sequence
	SummarizedFilterLookupTableSequence = New(0x003A, 0x0320)

	// NotchFilterCharacteristicsSequence (003A,0321) VR=SQ VM=1 Notch Filter Characteristics Sequence
	NotchFilterCharacteristicsSequence = New(0x003A, 0x0321)

	// WaveformFilterType (003A,0322) VR=CS VM=1 Waveform Filter Type
	WaveformFilterType = New(0x003A, 0x0322)

	// AnalogFilterCharacteristicsSequence (003A,0323) VR=SQ VM=1 Analog Filter Characteristics Sequence
	AnalogFilterCharacteristicsSequence = New(0x003A, 0x0323)

	// AnalogFilterRollOff (003A,0324) VR=DS VM=1 Analog Filter Roll Off
	AnalogFilterRollOff = New(0x003A, 0x0324)

	// AnalogFilterTypeCodeSequence (003A,0325) VR=SQ VM=1 Analog Filter Type Code Sequence
	AnalogFilterTypeCodeSequence = New(0x003A, 0x0325)

	// DigitalFilterCharacteristicsSequence (003A,0326) VR=SQ VM=1 Digital Filter Characteristics Sequence
	DigitalFilterCharacteristicsSequence = New(0x003A, 0x0326)

	// DigitalFilterOrder (003A,0327) VR=IS VM=1 Digital Filter Order
	DigitalFilterOrder = New(0x003A, 0x0327)

	// DigitalFilterTypeCodeSequence (003A,0328) VR=SQ VM=1 Digital Filter Type Code Sequence
	DigitalFilterTypeCodeSequence = New(0x003A, 0x0328)

	// WaveformFilterDescription (003A,0329) VR=ST VM=1 Waveform Filter Description
	WaveformFilterDescription = New(0x003A, 0x0329)

	// FilterLookupTableSequence (003A,032A) VR=SQ VM=1 Filter Lookup Table Sequence
	FilterLookupTableSequence = New(0x003A, 0x032A)

	// FilterLookupTableDescription (003A,032B) VR=ST VM=1 Filter Lookup Table Description
	FilterLookupTableDescription = New(0x003A, 0x032B)

	// FrequencyEncodingCodeSequence (003A,032C) VR=SQ VM=1 Frequency Encoding Code Sequence
	FrequencyEncodingCodeSequence = New(0x003A, 0x032C)

	// MagnitudeEncodingCodeSequence (003A,032D) VR=SQ VM=1 Magnitude Encoding Code Sequence
	MagnitudeEncodingCodeSequence = New(0x003A, 0x032D)

	// FilterLookupTableData (003A,032E) VR=OD VM=1 Filter Lookup Table Data
	FilterLookupTableData = New(0x003A, 0x032E)

	// ScheduledStationAETitle (0040,0001) VR=AE VM=1-n Scheduled Station AE Title
	ScheduledStationAETitle = New(0x0040, 0x0001)

	// ScheduledProcedureStepStartDate (0040,0002) VR=DA VM=1 Scheduled Procedure Step Start Date
	ScheduledProcedureStepStartDate = New(0x0040, 0x0002)

	// ScheduledProcedureStepStartTime (0040,0003) VR=TM VM=1 Scheduled Procedure Step Start Time
	ScheduledProcedureStepStartTime = New(0x0040, 0x0003)

	// ScheduledProcedureStepEndDate (0040,0004) VR=DA VM=1 Scheduled Procedure Step End Date
	ScheduledProcedureStepEndDate = New(0x0040, 0x0004)

	// ScheduledProcedureStepEndTime (0040,0005) VR=TM VM=1 Scheduled Procedure Step End Time
	ScheduledProcedureStepEndTime = New(0x0040, 0x0005)

	// ScheduledPerformingPhysicianName (0040,0006) VR=PN VM=1 Scheduled Performing Physician's Name
	ScheduledPerformingPhysicianName = New(0x0040, 0x0006)

	// ScheduledProcedureStepDescription (0040,0007) VR=LO VM=1 Scheduled Procedure Step Description
	ScheduledProcedureStepDescription = New(0x0040, 0x0007)

	// ScheduledProtocolCodeSequence (0040,0008) VR=SQ VM=1 Scheduled Protocol Code Sequence
	ScheduledProtocolCodeSequence = New(0x0040, 0x0008)

	// ScheduledProcedureStepID (0040,0009) VR=SH VM=1 Scheduled Procedure Step ID
	ScheduledProcedureStepID = New(0x0040, 0x0009)

	// StageCodeSequence (0040,000A) VR=SQ VM=1 Stage Code Sequence
	StageCodeSequence = New(0x0040, 0x000A)

	// ScheduledPerformingPhysicianIdentificationSequence (0040,000B) VR=SQ VM=1 Scheduled Performing Physician Identification Sequence
	ScheduledPerformingPhysicianIdentificationSequence = New(0x0040, 0x000B)

	// ScheduledStationName (0040,0010) VR=SH VM=1-n Scheduled Station Name
	ScheduledStationName = New(0x0040, 0x0010)

	// ScheduledProcedureStepLocation (0040,0011) VR=SH VM=1 Scheduled Procedure Step Location
	ScheduledProcedureStepLocation = New(0x0040, 0x0011)

	// PreMedication (0040,0012) VR=LO VM=1 Pre-Medication
	PreMedication = New(0x0040, 0x0012)

	// ScheduledProcedureStepStatus (0040,0020) VR=CS VM=1 Scheduled Procedure Step Status
	ScheduledProcedureStepStatus = New(0x0040, 0x0020)

	// OrderPlacerIdentifierSequence (0040,0026) VR=SQ VM=1 Order Placer Identifier Sequence
	OrderPlacerIdentifierSequence = New(0x0040, 0x0026)

	// OrderFillerIdentifierSequence (0040,0027) VR=SQ VM=1 Order Filler Identifier Sequence
	OrderFillerIdentifierSequence = New(0x0040, 0x0027)

	// LocalNamespaceEntityID (0040,0031) VR=UT VM=1 Local Namespace Entity ID
	LocalNamespaceEntityID = New(0x0040, 0x0031)

	// UniversalEntityID (0040,0032) VR=UT VM=1 Universal Entity ID
	UniversalEntityID = New(0x0040, 0x0032)

	// UniversalEntityIDType (0040,0033) VR=CS VM=1 Universal Entity ID Type
	UniversalEntityIDType = New(0x0040, 0x0033)

	// IdentifierTypeCode (0040,0035) VR=CS VM=1 Identifier Type Code
	IdentifierTypeCode = New(0x0040, 0x0035)

	// AssigningFacilitySequence (0040,0036) VR=SQ VM=1 Assigning Facility Sequence
	AssigningFacilitySequence = New(0x0040, 0x0036)

	// AssigningJurisdictionCodeSequence (0040,0039) VR=SQ VM=1 Assigning Jurisdiction Code Sequence
	AssigningJurisdictionCodeSequence = New(0x0040, 0x0039)

	// AssigningAgencyOrDepartmentCodeSequence (0040,003A) VR=SQ VM=1 Assigning Agency or Department Code Sequence
	AssigningAgencyOrDepartmentCodeSequence = New(0x0040, 0x003A)

	// ScheduledProcedureStepSequence (0040,0100) VR=SQ VM=1 Scheduled Procedure Step Sequence
	ScheduledProcedureStepSequence = New(0x0040, 0x0100)

	// ReferencedNonImageCompositeSOPInstanceSequence (0040,0220) VR=SQ VM=1 Referenced Non-Image Composite SOP Instance Sequence
	ReferencedNonImageCompositeSOPInstanceSequence = New(0x0040, 0x0220)

	// PerformedStationAETitle (0040,0241) VR=AE VM=1 Performed Station AE Title
	PerformedStationAETitle = New(0x0040, 0x0241)

	// PerformedStationName (0040,0242) VR=SH VM=1 Performed Station Name
	PerformedStationName = New(0x0040, 0x0242)

	// PerformedLocation (0040,0243) VR=SH VM=1 Performed Location
	PerformedLocation = New(0x0040, 0x0243)

	// PerformedProcedureStepStartDate (0040,0244) VR=DA VM=1 Performed Procedure Step Start Date
	PerformedProcedureStepStartDate = New(0x0040, 0x0244)

	// PerformedProcedureStepStartTime (0040,0245) VR=TM VM=1 Performed Procedure Step Start Time
	PerformedProcedureStepStartTime = New(0x0040, 0x0245)

	// PerformedProcedureStepEndDate (0040,0250) VR=DA VM=1 Performed Procedure Step End Date
	PerformedProcedureStepEndDate = New(0x0040, 0x0250)

	// PerformedProcedureStepEndTime (0040,0251) VR=TM VM=1 Performed Procedure Step End Time
	PerformedProcedureStepEndTime = New(0x0040, 0x0251)

	// PerformedProcedureStepStatus (0040,0252) VR=CS VM=1 Performed Procedure Step Status
	PerformedProcedureStepStatus = New(0x0040, 0x0252)

	// PerformedProcedureStepID (0040,0253) VR=SH VM=1 Performed Procedure Step ID
	PerformedProcedureStepID = New(0x0040, 0x0253)

	// PerformedProcedureStepDescription (0040,0254) VR=LO VM=1 Performed Procedure Step Description
	PerformedProcedureStepDescription = New(0x0040, 0x0254)

	// PerformedProcedureTypeDescription (0040,0255) VR=LO VM=1 Performed Procedure Type Description
	PerformedProcedureTypeDescription = New(0x0040, 0x0255)

	// PerformedProtocolCodeSequence (0040,0260) VR=SQ VM=1 Performed Protocol Code Sequence
	PerformedProtocolCodeSequence = New(0x0040, 0x0260)

	// PerformedProtocolType (0040,0261) VR=CS VM=1 Performed Protocol Type
	PerformedProtocolType = New(0x0040, 0x0261)

	// ScheduledStepAttributesSequence (0040,0270) VR=SQ VM=1 Scheduled Step Attributes Sequence
	ScheduledStepAttributesSequence = New(0x0040, 0x0270)

	// RequestAttributesSequence (0040,0275) VR=SQ VM=1 Request Attributes Sequence
	RequestAttributesSequence = New(0x0040, 0x0275)

	// CommentsOnThePerformedProcedureStep (0040,0280) VR=ST VM=1 Comments on the Performed Procedure Step
	CommentsOnThePerformedProcedureStep = New(0x0040, 0x0280)

	// PerformedProcedureStepDiscontinuationReasonCodeSequence (0040,0281) VR=SQ VM=1 Performed Procedure Step Discontinuation Reason Code Sequence
	PerformedProcedureStepDiscontinuationReasonCodeSequence = New(0x0040, 0x0281)

	// QuantitySequence (0040,0293) VR=SQ VM=1 Quantity Sequence
	QuantitySequence = New(0x0040, 0x0293)

	// Quantity (0040,0294) VR=DS VM=1 Quantity
	Quantity = New(0x0040, 0x0294)

	// MeasuringUnitsSequence (0040,0295) VR=SQ VM=1 Measuring Units Sequence
	MeasuringUnitsSequence = New(0x0040, 0x0295)

	// BillingItemSequence (0040,0296) VR=SQ VM=1 Billing Item Sequence
	BillingItemSequence = New(0x0040, 0x0296)

	// TotalTimeOfFluoroscopyRETIRED (0040,0300) VR=US VM=1 Total Time of Fluoroscopy (RETIRED)
	TotalTimeOfFluoroscopyRETIRED = New(0x0040, 0x0300)

	// TotalNumberOfExposuresRETIRED (0040,0301) VR=US VM=1 Total Number of Exposures (RETIRED)
	TotalNumberOfExposuresRETIRED = New(0x0040, 0x0301)

	// EntranceDose (0040,0302) VR=US VM=1 Entrance Dose
	EntranceDose = New(0x0040, 0x0302)

	// ExposedArea (0040,0303) VR=US VM=1-2 Exposed Area
	ExposedArea = New(0x0040, 0x0303)

	// DistanceSourceToEntrance (0040,0306) VR=DS VM=1 Distance Source to Entrance
	DistanceSourceToEntrance = New(0x0040, 0x0306)

	// DistanceSourceToSupportRETIRED (0040,0307) VR=DS VM=1 Distance Source to Support (RETIRED)
	DistanceSourceToSupportRETIRED = New(0x0040, 0x0307)

	// ExposureDoseSequenceRETIRED (0040,030E) VR=SQ VM=1 Exposure Dose Sequence (RETIRED)
	ExposureDoseSequenceRETIRED = New(0x0040, 0x030E)

	// CommentsOnRadiationDose (0040,0310) VR=ST VM=1 Comments on Radiation Dose
	CommentsOnRadiationDose = New(0x0040, 0x0310)

	// XRayOutput (0040,0312) VR=DS VM=1 X-Ray Output
	XRayOutput = New(0x0040, 0x0312)

	// HalfValueLayer (0040,0314) VR=DS VM=1 Half Value Layer
	HalfValueLayer = New(0x0040, 0x0314)

	// OrganDose (0040,0316) VR=DS VM=1 Organ Dose
	OrganDose = New(0x0040, 0x0316)

	// OrganExposed (0040,0318) VR=CS VM=1 Organ Exposed
	OrganExposed = New(0x0040, 0x0318)

	// BillingProcedureStepSequence (0040,0320) VR=SQ VM=1 Billing Procedure Step Sequence
	BillingProcedureStepSequence = New(0x0040, 0x0320)

	// FilmConsumptionSequence (0040,0321) VR=SQ VM=1 Film Consumption Sequence
	FilmConsumptionSequence = New(0x0040, 0x0321)

	// BillingSuppliesAndDevicesSequence (0040,0324) VR=SQ VM=1 Billing Supplies and Devices Sequence
	BillingSuppliesAndDevicesSequence = New(0x0040, 0x0324)

	// ReferencedProcedureStepSequenceRETIRED (0040,0330) VR=SQ VM=1 Referenced Procedure Step Sequence (RETIRED)
	ReferencedProcedureStepSequenceRETIRED = New(0x0040, 0x0330)

	// PerformedSeriesSequence (0040,0340) VR=SQ VM=1 Performed Series Sequence
	PerformedSeriesSequence = New(0x0040, 0x0340)

	// CommentsOnTheScheduledProcedureStep (0040,0400) VR=LT VM=1 Comments on the Scheduled Procedure Step
	CommentsOnTheScheduledProcedureStep = New(0x0040, 0x0400)

	// ProtocolContextSequence (0040,0440) VR=SQ VM=1 Protocol Context Sequence
	ProtocolContextSequence = New(0x0040, 0x0440)

	// ContentItemModifierSequence (0040,0441) VR=SQ VM=1 Content Item Modifier Sequence
	ContentItemModifierSequence = New(0x0040, 0x0441)

	// ScheduledSpecimenSequence (0040,0500) VR=SQ VM=1 Scheduled Specimen Sequence
	ScheduledSpecimenSequence = New(0x0040, 0x0500)

	// SpecimenAccessionNumberRETIRED (0040,050A) VR=LO VM=1 Specimen Accession Number (RETIRED)
	SpecimenAccessionNumberRETIRED = New(0x0040, 0x050A)

	// ContainerIdentifier (0040,0512) VR=LO VM=1 Container Identifier
	ContainerIdentifier = New(0x0040, 0x0512)

	// IssuerOfTheContainerIdentifierSequence (0040,0513) VR=SQ VM=1 Issuer of the Container Identifier Sequence
	IssuerOfTheContainerIdentifierSequence = New(0x0040, 0x0513)

	// AlternateContainerIdentifierSequence (0040,0515) VR=SQ VM=1 Alternate Container Identifier Sequence
	AlternateContainerIdentifierSequence = New(0x0040, 0x0515)

	// ContainerTypeCodeSequence (0040,0518) VR=SQ VM=1 Container Type Code Sequence
	ContainerTypeCodeSequence = New(0x0040, 0x0518)

	// ContainerDescription (0040,051A) VR=LO VM=1 Container Description
	ContainerDescription = New(0x0040, 0x051A)

	// ContainerComponentSequence (0040,0520) VR=SQ VM=1 Container Component Sequence
	ContainerComponentSequence = New(0x0040, 0x0520)

	// SpecimenSequenceRETIRED (0040,0550) VR=SQ VM=1 Specimen Sequence (RETIRED)
	SpecimenSequenceRETIRED = New(0x0040, 0x0550)

	// SpecimenIdentifier (0040,0551) VR=LO VM=1 Specimen Identifier
	SpecimenIdentifier = New(0x0040, 0x0551)

	// SpecimenDescriptionSequenceTrialRETIRED (0040,0552) VR=SQ VM=1 Specimen Description Sequence (Trial) (RETIRED)
	SpecimenDescriptionSequenceTrialRETIRED = New(0x0040, 0x0552)

	// SpecimenDescriptionTrialRETIRED (0040,0553) VR=ST VM=1 Specimen Description (Trial) (RETIRED)
	SpecimenDescriptionTrialRETIRED = New(0x0040, 0x0553)

	// SpecimenUID (0040,0554) VR=UI VM=1 Specimen UID
	SpecimenUID = New(0x0040, 0x0554)

	// AcquisitionContextSequence (0040,0555) VR=SQ VM=1 Acquisition Context Sequence
	AcquisitionContextSequence = New(0x0040, 0x0555)

	// AcquisitionContextDescription (0040,0556) VR=ST VM=1 Acquisition Context Description
	AcquisitionContextDescription = New(0x0040, 0x0556)

	// SpecimenTypeCodeSequence (0040,059A) VR=SQ VM=1 Specimen Type Code Sequence
	SpecimenTypeCodeSequence = New(0x0040, 0x059A)

	// SpecimenDescriptionSequence (0040,0560) VR=SQ VM=1 Specimen Description Sequence
	SpecimenDescriptionSequence = New(0x0040, 0x0560)

	// IssuerOfTheSpecimenIdentifierSequence (0040,0562) VR=SQ VM=1 Issuer of the Specimen Identifier Sequence
	IssuerOfTheSpecimenIdentifierSequence = New(0x0040, 0x0562)

	// SpecimenShortDescription (0040,0600) VR=LO VM=1 Specimen Short Description
	SpecimenShortDescription = New(0x0040, 0x0600)

	// SpecimenDetailedDescription (0040,0602) VR=UT VM=1 Specimen Detailed Description
	SpecimenDetailedDescription = New(0x0040, 0x0602)

	// SpecimenPreparationSequence (0040,0610) VR=SQ VM=1 Specimen Preparation Sequence
	SpecimenPreparationSequence = New(0x0040, 0x0610)

	// SpecimenPreparationStepContentItemSequence (0040,0612) VR=SQ VM=1 Specimen Preparation Step Content Item Sequence
	SpecimenPreparationStepContentItemSequence = New(0x0040, 0x0612)

	// SpecimenLocalizationContentItemSequence (0040,0620) VR=SQ VM=1 Specimen Localization Content Item Sequence
	SpecimenLocalizationContentItemSequence = New(0x0040, 0x0620)

	// SlideIdentifierRETIRED (0040,06FA) VR=LO VM=1 Slide Identifier (RETIRED)
	SlideIdentifierRETIRED = New(0x0040, 0x06FA)

	// WholeSlideMicroscopyImageFrameTypeSequence (0040,0710) VR=SQ VM=1 Whole Slide Microscopy Image Frame Type Sequence
	WholeSlideMicroscopyImageFrameTypeSequence = New(0x0040, 0x0710)

	// ImageCenterPointCoordinatesSequence (0040,071A) VR=SQ VM=1 Image Center Point Coordinates Sequence
	ImageCenterPointCoordinatesSequence = New(0x0040, 0x071A)

	// XOffsetInSlideCoordinateSystem (0040,072A) VR=DS VM=1 X Offset in Slide Coordinate System
	XOffsetInSlideCoordinateSystem = New(0x0040, 0x072A)

	// YOffsetInSlideCoordinateSystem (0040,073A) VR=DS VM=1 Y Offset in Slide Coordinate System
	YOffsetInSlideCoordinateSystem = New(0x0040, 0x073A)

	// ZOffsetInSlideCoordinateSystem (0040,074A) VR=DS VM=1 Z Offset in Slide Coordinate System
	ZOffsetInSlideCoordinateSystem = New(0x0040, 0x074A)

	// PixelSpacingSequenceRETIRED (0040,08D8) VR=SQ VM=1 Pixel Spacing Sequence (RETIRED)
	PixelSpacingSequenceRETIRED = New(0x0040, 0x08D8)

	// CoordinateSystemAxisCodeSequenceRETIRED (0040,08DA) VR=SQ VM=1 Coordinate System Axis Code Sequence (RETIRED)
	CoordinateSystemAxisCodeSequenceRETIRED = New(0x0040, 0x08DA)

	// MeasurementUnitsCodeSequence (0040,08EA) VR=SQ VM=1 Measurement Units Code Sequence
	MeasurementUnitsCodeSequence = New(0x0040, 0x08EA)

	// VitalStainCodeSequenceTrialRETIRED (0040,09F8) VR=SQ VM=1 Vital Stain Code Sequence (Trial) (RETIRED)
	VitalStainCodeSequenceTrialRETIRED = New(0x0040, 0x09F8)

	// RequestedProcedureID (0040,1001) VR=SH VM=1 Requested Procedure ID
	RequestedProcedureID = New(0x0040, 0x1001)

	// ReasonForTheRequestedProcedure (0040,1002) VR=LO VM=1 Reason for the Requested Procedure
	ReasonForTheRequestedProcedure = New(0x0040, 0x1002)

	// RequestedProcedurePriority (0040,1003) VR=SH VM=1 Requested Procedure Priority
	RequestedProcedurePriority = New(0x0040, 0x1003)

	// PatientTransportArrangements (0040,1004) VR=LO VM=1 Patient Transport Arrangements
	PatientTransportArrangements = New(0x0040, 0x1004)

	// RequestedProcedureLocation (0040,1005) VR=LO VM=1 Requested Procedure Location
	RequestedProcedureLocation = New(0x0040, 0x1005)

	// PlacerOrderNumberProcedureRETIRED (0040,1006) VR=SH VM=1 Placer Order Number / Procedure (RETIRED)
	PlacerOrderNumberProcedureRETIRED = New(0x0040, 0x1006)

	// FillerOrderNumberProcedureRETIRED (0040,1007) VR=SH VM=1 Filler Order Number / Procedure (RETIRED)
	FillerOrderNumberProcedureRETIRED = New(0x0040, 0x1007)

	// ConfidentialityCode (0040,1008) VR=LO VM=1 Confidentiality Code
	ConfidentialityCode = New(0x0040, 0x1008)

	// ReportingPriority (0040,1009) VR=SH VM=1 Reporting Priority
	ReportingPriority = New(0x0040, 0x1009)

	// ReasonForRequestedProcedureCodeSequence (0040,100A) VR=SQ VM=1 Reason for Requested Procedure Code Sequence
	ReasonForRequestedProcedureCodeSequence = New(0x0040, 0x100A)

	// NamesOfIntendedRecipientsOfResults (0040,1010) VR=PN VM=1-n Names of Intended Recipients of Results
	NamesOfIntendedRecipientsOfResults = New(0x0040, 0x1010)

	// IntendedRecipientsOfResultsIdentificationSequence (0040,1011) VR=SQ VM=1 Intended Recipients of Results Identification Sequence
	IntendedRecipientsOfResultsIdentificationSequence = New(0x0040, 0x1011)

	// ReasonForPerformedProcedureCodeSequence (0040,1012) VR=SQ VM=1 Reason For Performed Procedure Code Sequence
	ReasonForPerformedProcedureCodeSequence = New(0x0040, 0x1012)

	// RequestedProcedureDescriptionTrialRETIRED (0040,1060) VR=LO VM=1 Requested Procedure Description (Trial) (RETIRED)
	RequestedProcedureDescriptionTrialRETIRED = New(0x0040, 0x1060)

	// PersonIdentificationCodeSequence (0040,1101) VR=SQ VM=1 Person Identification Code Sequence
	PersonIdentificationCodeSequence = New(0x0040, 0x1101)

	// PersonAddress (0040,1102) VR=ST VM=1 Person's Address
	PersonAddress = New(0x0040, 0x1102)

	// PersonTelephoneNumbers (0040,1103) VR=LO VM=1-n Person's Telephone Numbers
	PersonTelephoneNumbers = New(0x0040, 0x1103)

	// PersonTelecomInformation (0040,1104) VR=LT VM=1 Person's Telecom Information
	PersonTelecomInformation = New(0x0040, 0x1104)

	// RequestedProcedureComments (0040,1400) VR=LT VM=1 Requested Procedure Comments
	RequestedProcedureComments = New(0x0040, 0x1400)

	// ReasonForTheImagingServiceRequestRETIRED (0040,2001) VR=LO VM=1 Reason for the Imaging Service Request (RETIRED)
	ReasonForTheImagingServiceRequestRETIRED = New(0x0040, 0x2001)

	// IssueDateOfImagingServiceRequest (0040,2004) VR=DA VM=1 Issue Date of Imaging Service Request
	IssueDateOfImagingServiceRequest = New(0x0040, 0x2004)

	// IssueTimeOfImagingServiceRequest (0040,2005) VR=TM VM=1 Issue Time of Imaging Service Request
	IssueTimeOfImagingServiceRequest = New(0x0040, 0x2005)

	// PlacerOrderNumberImagingServiceRequestRetiredRETIRED (0040,2006) VR=SH VM=1 Placer Order Number / Imaging Service Request (Retired) (RETIRED)
	PlacerOrderNumberImagingServiceRequestRetiredRETIRED = New(0x0040, 0x2006)

	// FillerOrderNumberImagingServiceRequestRetiredRETIRED (0040,2007) VR=SH VM=1 Filler Order Number / Imaging Service Request (Retired) (RETIRED)
	FillerOrderNumberImagingServiceRequestRetiredRETIRED = New(0x0040, 0x2007)

	// OrderEnteredBy (0040,2008) VR=PN VM=1 Order Entered By
	OrderEnteredBy = New(0x0040, 0x2008)

	// OrderEntererLocation (0040,2009) VR=SH VM=1 Order Enterer's Location
	OrderEntererLocation = New(0x0040, 0x2009)

	// OrderCallbackPhoneNumber (0040,2010) VR=SH VM=1 Order Callback Phone Number
	OrderCallbackPhoneNumber = New(0x0040, 0x2010)

	// OrderCallbackTelecomInformation (0040,2011) VR=LT VM=1 Order Callback Telecom Information
	OrderCallbackTelecomInformation = New(0x0040, 0x2011)

	// PlacerOrderNumberImagingServiceRequest (0040,2016) VR=LO VM=1 Placer Order Number / Imaging Service Request
	PlacerOrderNumberImagingServiceRequest = New(0x0040, 0x2016)

	// FillerOrderNumberImagingServiceRequest (0040,2017) VR=LO VM=1 Filler Order Number / Imaging Service Request
	FillerOrderNumberImagingServiceRequest = New(0x0040, 0x2017)

	// ImagingServiceRequestComments (0040,2400) VR=LT VM=1 Imaging Service Request Comments
	ImagingServiceRequestComments = New(0x0040, 0x2400)

	// ConfidentialityConstraintOnPatientDataDescription (0040,3001) VR=LO VM=1 Confidentiality Constraint on Patient Data Description
	ConfidentialityConstraintOnPatientDataDescription = New(0x0040, 0x3001)

	// GeneralPurposeScheduledProcedureStepStatusRETIRED (0040,4001) VR=CS VM=1 General Purpose Scheduled Procedure Step Status (RETIRED)
	GeneralPurposeScheduledProcedureStepStatusRETIRED = New(0x0040, 0x4001)

	// GeneralPurposePerformedProcedureStepStatusRETIRED (0040,4002) VR=CS VM=1 General Purpose Performed Procedure Step Status (RETIRED)
	GeneralPurposePerformedProcedureStepStatusRETIRED = New(0x0040, 0x4002)

	// GeneralPurposeScheduledProcedureStepPriorityRETIRED (0040,4003) VR=CS VM=1 General Purpose Scheduled Procedure Step Priority (RETIRED)
	GeneralPurposeScheduledProcedureStepPriorityRETIRED = New(0x0040, 0x4003)

	// ScheduledProcessingApplicationsCodeSequenceRETIRED (0040,4004) VR=SQ VM=1 Scheduled Processing Applications Code Sequence (RETIRED)
	ScheduledProcessingApplicationsCodeSequenceRETIRED = New(0x0040, 0x4004)

	// ScheduledProcedureStepStartDateTime (0040,4005) VR=DT VM=1 Scheduled Procedure Step Start DateTime
	ScheduledProcedureStepStartDateTime = New(0x0040, 0x4005)

	// MultipleCopiesFlagRETIRED (0040,4006) VR=CS VM=1 Multiple Copies Flag (RETIRED)
	MultipleCopiesFlagRETIRED = New(0x0040, 0x4006)

	// PerformedProcessingApplicationsCodeSequenceRETIRED (0040,4007) VR=SQ VM=1 Performed Processing Applications Code Sequence (RETIRED)
	PerformedProcessingApplicationsCodeSequenceRETIRED = New(0x0040, 0x4007)

	// ScheduledProcedureStepExpirationDateTime (0040,4008) VR=DT VM=1 Scheduled Procedure Step Expiration DateTime
	ScheduledProcedureStepExpirationDateTime = New(0x0040, 0x4008)

	// HumanPerformerCodeSequence (0040,4009) VR=SQ VM=1 Human Performer Code Sequence
	HumanPerformerCodeSequence = New(0x0040, 0x4009)

	// ScheduledProcedureStepModificationDateTime (0040,4010) VR=DT VM=1 Scheduled Procedure Step Modification DateTime
	ScheduledProcedureStepModificationDateTime = New(0x0040, 0x4010)

	// ExpectedCompletionDateTime (0040,4011) VR=DT VM=1 Expected Completion DateTime
	ExpectedCompletionDateTime = New(0x0040, 0x4011)

	// ResultingGeneralPurposePerformedProcedureStepsSequenceRETIRED (0040,4015) VR=SQ VM=1 Resulting General Purpose Performed Procedure Steps Sequence (RETIRED)
	ResultingGeneralPurposePerformedProcedureStepsSequenceRETIRED = New(0x0040, 0x4015)

	// ReferencedGeneralPurposeScheduledProcedureStepSequenceRETIRED (0040,4016) VR=SQ VM=1 Referenced General Purpose Scheduled Procedure Step Sequence (RETIRED)
	ReferencedGeneralPurposeScheduledProcedureStepSequenceRETIRED = New(0x0040, 0x4016)

	// ScheduledWorkitemCodeSequence (0040,4018) VR=SQ VM=1 Scheduled Workitem Code Sequence
	ScheduledWorkitemCodeSequence = New(0x0040, 0x4018)

	// PerformedWorkitemCodeSequence (0040,4019) VR=SQ VM=1 Performed Workitem Code Sequence
	PerformedWorkitemCodeSequence = New(0x0040, 0x4019)

	// InputAvailabilityFlagRETIRED (0040,4020) VR=CS VM=1 Input Availability Flag (RETIRED)
	InputAvailabilityFlagRETIRED = New(0x0040, 0x4020)

	// InputInformationSequence (0040,4021) VR=SQ VM=1 Input Information Sequence
	InputInformationSequence = New(0x0040, 0x4021)

	// RelevantInformationSequenceRETIRED (0040,4022) VR=SQ VM=1 Relevant Information Sequence (RETIRED)
	RelevantInformationSequenceRETIRED = New(0x0040, 0x4022)

	// ReferencedGeneralPurposeScheduledProcedureStepTransactionUIDRETIRED (0040,4023) VR=UI VM=1 Referenced General Purpose Scheduled Procedure Step Transaction UID (RETIRED)
	ReferencedGeneralPurposeScheduledProcedureStepTransactionUIDRETIRED = New(0x0040, 0x4023)

	// ScheduledStationNameCodeSequence (0040,4025) VR=SQ VM=1 Scheduled Station Name Code Sequence
	ScheduledStationNameCodeSequence = New(0x0040, 0x4025)

	// ScheduledStationClassCodeSequence (0040,4026) VR=SQ VM=1 Scheduled Station Class Code Sequence
	ScheduledStationClassCodeSequence = New(0x0040, 0x4026)

	// ScheduledStationGeographicLocationCodeSequence (0040,4027) VR=SQ VM=1 Scheduled Station Geographic Location Code Sequence
	ScheduledStationGeographicLocationCodeSequence = New(0x0040, 0x4027)

	// PerformedStationNameCodeSequence (0040,4028) VR=SQ VM=1 Performed Station Name Code Sequence
	PerformedStationNameCodeSequence = New(0x0040, 0x4028)

	// PerformedStationClassCodeSequence (0040,4029) VR=SQ VM=1 Performed Station Class Code Sequence
	PerformedStationClassCodeSequence = New(0x0040, 0x4029)

	// PerformedStationGeographicLocationCodeSequence (0040,4030) VR=SQ VM=1 Performed Station Geographic Location Code Sequence
	PerformedStationGeographicLocationCodeSequence = New(0x0040, 0x4030)

	// RequestedSubsequentWorkitemCodeSequenceRETIRED (0040,4031) VR=SQ VM=1 Requested Subsequent Workitem Code Sequence (RETIRED)
	RequestedSubsequentWorkitemCodeSequenceRETIRED = New(0x0040, 0x4031)

	// NonDICOMOutputCodeSequenceRETIRED (0040,4032) VR=SQ VM=1 Non-DICOM Output Code Sequence (RETIRED)
	NonDICOMOutputCodeSequenceRETIRED = New(0x0040, 0x4032)

	// OutputInformationSequence (0040,4033) VR=SQ VM=1 Output Information Sequence
	OutputInformationSequence = New(0x0040, 0x4033)

	// ScheduledHumanPerformersSequence (0040,4034) VR=SQ VM=1 Scheduled Human Performers Sequence
	ScheduledHumanPerformersSequence = New(0x0040, 0x4034)

	// ActualHumanPerformersSequence (0040,4035) VR=SQ VM=1 Actual Human Performers Sequence
	ActualHumanPerformersSequence = New(0x0040, 0x4035)

	// HumanPerformerOrganization (0040,4036) VR=LO VM=1 Human Performer's Organization
	HumanPerformerOrganization = New(0x0040, 0x4036)

	// HumanPerformerName (0040,4037) VR=PN VM=1 Human Performer's Name
	HumanPerformerName = New(0x0040, 0x4037)

	// RawDataHandling (0040,4040) VR=CS VM=1 Raw Data Handling
	RawDataHandling = New(0x0040, 0x4040)

	// InputReadinessState (0040,4041) VR=CS VM=1 Input Readiness State
	InputReadinessState = New(0x0040, 0x4041)

	// PerformedProcedureStepStartDateTime (0040,4050) VR=DT VM=1 Performed Procedure Step Start DateTime
	PerformedProcedureStepStartDateTime = New(0x0040, 0x4050)

	// PerformedProcedureStepEndDateTime (0040,4051) VR=DT VM=1 Performed Procedure Step End DateTime
	PerformedProcedureStepEndDateTime = New(0x0040, 0x4051)

	// ProcedureStepCancellationDateTime (0040,4052) VR=DT VM=1 Procedure Step Cancellation DateTime
	ProcedureStepCancellationDateTime = New(0x0040, 0x4052)

	// OutputDestinationSequence (0040,4070) VR=SQ VM=1 Output Destination Sequence
	OutputDestinationSequence = New(0x0040, 0x4070)

	// DICOMStorageSequence (0040,4071) VR=SQ VM=1 DICOM Storage Sequence
	DICOMStorageSequence = New(0x0040, 0x4071)

	// STOWRSStorageSequence (0040,4072) VR=SQ VM=1 STOW-RS Storage Sequence
	STOWRSStorageSequence = New(0x0040, 0x4072)

	// StorageURL (0040,4073) VR=UR VM=1 Storage URL
	StorageURL = New(0x0040, 0x4073)

	// XDSStorageSequence (0040,4074) VR=SQ VM=1 XDS Storage Sequence
	XDSStorageSequence = New(0x0040, 0x4074)

	// EntranceDoseInmGy (0040,8302) VR=DS VM=1 Entrance Dose in mGy
	EntranceDoseInmGy = New(0x0040, 0x8302)

	// EntranceDoseDerivation (0040,8303) VR=CS VM=1 Entrance Dose Derivation
	EntranceDoseDerivation = New(0x0040, 0x8303)

	// ParametricMapFrameTypeSequence (0040,9092) VR=SQ VM=1 Parametric Map Frame Type Sequence
	ParametricMapFrameTypeSequence = New(0x0040, 0x9092)

	// ReferencedImageRealWorldValueMappingSequence (0040,9094) VR=SQ VM=1 Referenced Image Real World Value Mapping Sequence
	ReferencedImageRealWorldValueMappingSequence = New(0x0040, 0x9094)

	// RealWorldValueMappingSequence (0040,9096) VR=SQ VM=1 Real World Value Mapping Sequence
	RealWorldValueMappingSequence = New(0x0040, 0x9096)

	// PixelValueMappingCodeSequence (0040,9098) VR=SQ VM=1 Pixel Value Mapping Code Sequence
	PixelValueMappingCodeSequence = New(0x0040, 0x9098)

	// LUTLabel (0040,9210) VR=SH VM=1 LUT Label
	LUTLabel = New(0x0040, 0x9210)

	RealWorldValueLastValueMapped = New(0x0040, 0x9211)

	// RealWorldValueLUTData (0040,9212) VR=FD VM=1-n Real World Value LUT Data
	RealWorldValueLUTData = New(0x0040, 0x9212)

	// DoubleFloatRealWorldValueLastValueMapped (0040,9213) VR=FD VM=1 Double Float Real World Value Last Value Mapped
	DoubleFloatRealWorldValueLastValueMapped = New(0x0040, 0x9213)

	// DoubleFloatRealWorldValueFirstValueMapped (0040,9214) VR=FD VM=1 Double Float Real World Value First Value Mapped
	DoubleFloatRealWorldValueFirstValueMapped = New(0x0040, 0x9214)

	RealWorldValueFirstValueMapped = New(0x0040, 0x9216)

	// QuantityDefinitionSequence (0040,9220) VR=SQ VM=1 Quantity Definition Sequence
	QuantityDefinitionSequence = New(0x0040, 0x9220)

	// RealWorldValueIntercept (0040,9224) VR=FD VM=1 Real World Value Intercept
	RealWorldValueIntercept = New(0x0040, 0x9224)

	// RealWorldValueSlope (0040,9225) VR=FD VM=1 Real World Value Slope
	RealWorldValueSlope = New(0x0040, 0x9225)

	// FindingsFlagTrialRETIRED (0040,A007) VR=CS VM=1 Findings Flag (Trial) (RETIRED)
	FindingsFlagTrialRETIRED = New(0x0040, 0xA007)

	// RelationshipType (0040,A010) VR=CS VM=1 Relationship Type
	RelationshipType = New(0x0040, 0xA010)

	// FindingsSequenceTrialRETIRED (0040,A020) VR=SQ VM=1 Findings Sequence (Trial) (RETIRED)
	FindingsSequenceTrialRETIRED = New(0x0040, 0xA020)

	// FindingsGroupUIDTrialRETIRED (0040,A021) VR=UI VM=1 Findings Group UID (Trial) (RETIRED)
	FindingsGroupUIDTrialRETIRED = New(0x0040, 0xA021)

	// ReferencedFindingsGroupUIDTrialRETIRED (0040,A022) VR=UI VM=1 Referenced Findings Group UID (Trial) (RETIRED)
	ReferencedFindingsGroupUIDTrialRETIRED = New(0x0040, 0xA022)

	// FindingsGroupRecordingDateTrialRETIRED (0040,A023) VR=DA VM=1 Findings Group Recording Date (Trial) (RETIRED)
	FindingsGroupRecordingDateTrialRETIRED = New(0x0040, 0xA023)

	// FindingsGroupRecordingTimeTrialRETIRED (0040,A024) VR=TM VM=1 Findings Group Recording Time (Trial) (RETIRED)
	FindingsGroupRecordingTimeTrialRETIRED = New(0x0040, 0xA024)

	// FindingsSourceCategoryCodeSequenceTrialRETIRED (0040,A026) VR=SQ VM=1 Findings Source Category Code Sequence (Trial) (RETIRED)
	FindingsSourceCategoryCodeSequenceTrialRETIRED = New(0x0040, 0xA026)

	// VerifyingOrganization (0040,A027) VR=LO VM=1 Verifying Organization
	VerifyingOrganization = New(0x0040, 0xA027)

	// DocumentingOrganizationIdentifierCodeSequenceTrialRETIRED (0040,A028) VR=SQ VM=1 Documenting Organization Identifier Code Sequence (Trial) (RETIRED)
	DocumentingOrganizationIdentifierCodeSequenceTrialRETIRED = New(0x0040, 0xA028)

	// VerificationDateTime (0040,A030) VR=DT VM=1 Verification DateTime
	VerificationDateTime = New(0x0040, 0xA030)

	// ObservationDateTime (0040,A032) VR=DT VM=1 Observation DateTime
	ObservationDateTime = New(0x0040, 0xA032)

	// ObservationStartDateTime (0040,A033) VR=DT VM=1 Observation Start DateTime
	ObservationStartDateTime = New(0x0040, 0xA033)

	// EffectiveStartDateTime (0040,A034) VR=DT VM=1 Effective Start DateTime
	EffectiveStartDateTime = New(0x0040, 0xA034)

	// EffectiveStopDateTime (0040,A035) VR=DT VM=1 Effective Stop DateTime
	EffectiveStopDateTime = New(0x0040, 0xA035)

	// ValueType (0040,A040) VR=CS VM=1 Value Type
	ValueType = New(0x0040, 0xA040)

	// ConceptNameCodeSequence (0040,A043) VR=SQ VM=1 Concept Name Code Sequence
	ConceptNameCodeSequence = New(0x0040, 0xA043)

	// MeasurementPrecisionDescriptionTrialRETIRED (0040,A047) VR=LO VM=1 Measurement Precision Description (Trial) (RETIRED)
	MeasurementPrecisionDescriptionTrialRETIRED = New(0x0040, 0xA047)

	// ContinuityOfContent (0040,A050) VR=CS VM=1 Continuity Of Content
	ContinuityOfContent = New(0x0040, 0xA050)

	// UrgencyOrPriorityAlertsTrialRETIRED (0040,A057) VR=CS VM=1-n Urgency or Priority Alerts (Trial) (RETIRED)
	UrgencyOrPriorityAlertsTrialRETIRED = New(0x0040, 0xA057)

	// SequencingIndicatorTrialRETIRED (0040,A060) VR=LO VM=1 Sequencing Indicator (Trial) (RETIRED)
	SequencingIndicatorTrialRETIRED = New(0x0040, 0xA060)

	// DocumentIdentifierCodeSequenceTrialRETIRED (0040,A066) VR=SQ VM=1 Document Identifier Code Sequence (Trial) (RETIRED)
	DocumentIdentifierCodeSequenceTrialRETIRED = New(0x0040, 0xA066)

	// DocumentAuthorTrialRETIRED (0040,A067) VR=PN VM=1 Document Author (Trial) (RETIRED)
	DocumentAuthorTrialRETIRED = New(0x0040, 0xA067)

	// DocumentAuthorIdentifierCodeSequenceTrialRETIRED (0040,A068) VR=SQ VM=1 Document Author Identifier Code Sequence (Trial) (RETIRED)
	DocumentAuthorIdentifierCodeSequenceTrialRETIRED = New(0x0040, 0xA068)

	// IdentifierCodeSequenceTrialRETIRED (0040,A070) VR=SQ VM=1 Identifier Code Sequence (Trial) (RETIRED)
	IdentifierCodeSequenceTrialRETIRED = New(0x0040, 0xA070)

	// VerifyingObserverSequence (0040,A073) VR=SQ VM=1 Verifying Observer Sequence
	VerifyingObserverSequence = New(0x0040, 0xA073)

	// ObjectBinaryIdentifierTrialRETIRED (0040,A074) VR=OB VM=1 Object Binary Identifier (Trial) (RETIRED)
	ObjectBinaryIdentifierTrialRETIRED = New(0x0040, 0xA074)

	// VerifyingObserverName (0040,A075) VR=PN VM=1 Verifying Observer Name
	VerifyingObserverName = New(0x0040, 0xA075)

	// DocumentingObserverIdentifierCodeSequenceTrialRETIRED (0040,A076) VR=SQ VM=1 Documenting Observer Identifier Code Sequence (Trial) (RETIRED)
	DocumentingObserverIdentifierCodeSequenceTrialRETIRED = New(0x0040, 0xA076)

	// AuthorObserverSequence (0040,A078) VR=SQ VM=1 Author Observer Sequence
	AuthorObserverSequence = New(0x0040, 0xA078)

	// ParticipantSequence (0040,A07A) VR=SQ VM=1 Participant Sequence
	ParticipantSequence = New(0x0040, 0xA07A)

	// CustodialOrganizationSequence (0040,A07C) VR=SQ VM=1 Custodial Organization Sequence
	CustodialOrganizationSequence = New(0x0040, 0xA07C)

	// ParticipationType (0040,A080) VR=CS VM=1 Participation Type
	ParticipationType = New(0x0040, 0xA080)

	// ParticipationDateTime (0040,A082) VR=DT VM=1 Participation DateTime
	ParticipationDateTime = New(0x0040, 0xA082)

	// ObserverType (0040,A084) VR=CS VM=1 Observer Type
	ObserverType = New(0x0040, 0xA084)

	// ProcedureIdentifierCodeSequenceTrialRETIRED (0040,A085) VR=SQ VM=1 Procedure Identifier Code Sequence (Trial) (RETIRED)
	ProcedureIdentifierCodeSequenceTrialRETIRED = New(0x0040, 0xA085)

	// VerifyingObserverIdentificationCodeSequence (0040,A088) VR=SQ VM=1 Verifying Observer Identification Code Sequence
	VerifyingObserverIdentificationCodeSequence = New(0x0040, 0xA088)

	// ObjectDirectoryBinaryIdentifierTrialRETIRED (0040,A089) VR=OB VM=1 Object Directory Binary Identifier (Trial) (RETIRED)
	ObjectDirectoryBinaryIdentifierTrialRETIRED = New(0x0040, 0xA089)

	// EquivalentCDADocumentSequenceRETIRED (0040,A090) VR=SQ VM=1 Equivalent CDA Document Sequence (RETIRED)
	EquivalentCDADocumentSequenceRETIRED = New(0x0040, 0xA090)

	// ReferencedWaveformChannels (0040,A0B0) VR=US VM=2-2n Referenced Waveform Channels
	ReferencedWaveformChannels = New(0x0040, 0xA0B0)

	// DateOfDocumentOrVerbalTransactionTrialRETIRED (0040,A110) VR=DA VM=1 Date of Document or Verbal Transaction (Trial) (RETIRED)
	DateOfDocumentOrVerbalTransactionTrialRETIRED = New(0x0040, 0xA110)

	// TimeOfDocumentCreationOrVerbalTransactionTrialRETIRED (0040,A112) VR=TM VM=1 Time of Document Creation or Verbal Transaction (Trial) (RETIRED)
	TimeOfDocumentCreationOrVerbalTransactionTrialRETIRED = New(0x0040, 0xA112)

	// DateTime (0040,A120) VR=DT VM=1 DateTime
	DateTime = New(0x0040, 0xA120)

	// Date (0040,A121) VR=DA VM=1 Date
	Date = New(0x0040, 0xA121)

	// Time (0040,A122) VR=TM VM=1 Time
	Time = New(0x0040, 0xA122)

	// PersonName (0040,A123) VR=PN VM=1 Person Name
	PersonName = New(0x0040, 0xA123)

	// UID (0040,A124) VR=UI VM=1 UID
	UID = New(0x0040, 0xA124)

	// ReportStatusIDTrialRETIRED (0040,A125) VR=CS VM=2 Report Status ID (Trial) (RETIRED)
	ReportStatusIDTrialRETIRED = New(0x0040, 0xA125)

	// TemporalRangeType (0040,A130) VR=CS VM=1 Temporal Range Type
	TemporalRangeType = New(0x0040, 0xA130)

	// ReferencedSamplePositions (0040,A132) VR=UL VM=1-n Referenced Sample Positions
	ReferencedSamplePositions = New(0x0040, 0xA132)

	// ReferencedFrameNumbersRETIRED (0040,A136) VR=US VM=1-n Referenced Frame Numbers (RETIRED)
	ReferencedFrameNumbersRETIRED = New(0x0040, 0xA136)

	// ReferencedTimeOffsets (0040,A138) VR=DS VM=1-n Referenced Time Offsets
	ReferencedTimeOffsets = New(0x0040, 0xA138)

	// ReferencedDateTime (0040,A13A) VR=DT VM=1-n Referenced DateTime
	ReferencedDateTime = New(0x0040, 0xA13A)

	// TextValue (0040,A160) VR=UT VM=1 Text Value
	TextValue = New(0x0040, 0xA160)

	// FloatingPointValue (0040,A161) VR=FD VM=1-n Floating Point Value
	FloatingPointValue = New(0x0040, 0xA161)

	// RationalNumeratorValue (0040,A162) VR=SL VM=1-n Rational Numerator Value
	RationalNumeratorValue = New(0x0040, 0xA162)

	// RationalDenominatorValue (0040,A163) VR=UL VM=1-n Rational Denominator Value
	RationalDenominatorValue = New(0x0040, 0xA163)

	// ObservationCategoryCodeSequenceTrialRETIRED (0040,A167) VR=SQ VM=1 Observation Category Code Sequence (Trial) (RETIRED)
	ObservationCategoryCodeSequenceTrialRETIRED = New(0x0040, 0xA167)

	// ConceptCodeSequence (0040,A168) VR=SQ VM=1 Concept Code Sequence
	ConceptCodeSequence = New(0x0040, 0xA168)

	// BibliographicCitationTrialRETIRED (0040,A16A) VR=ST VM=1 Bibliographic Citation (Trial) (RETIRED)
	BibliographicCitationTrialRETIRED = New(0x0040, 0xA16A)

	// PurposeOfReferenceCodeSequence (0040,A170) VR=SQ VM=1 Purpose of Reference Code Sequence
	PurposeOfReferenceCodeSequence = New(0x0040, 0xA170)

	// ObservationUID (0040,A171) VR=UI VM=1 Observation UID
	ObservationUID = New(0x0040, 0xA171)

	// ReferencedObservationUIDTrialRETIRED (0040,A172) VR=UI VM=1 Referenced Observation UID (Trial) (RETIRED)
	ReferencedObservationUIDTrialRETIRED = New(0x0040, 0xA172)

	// ReferencedObservationClassTrialRETIRED (0040,A173) VR=CS VM=1 Referenced Observation Class (Trial) (RETIRED)
	ReferencedObservationClassTrialRETIRED = New(0x0040, 0xA173)

	// ReferencedObjectObservationClassTrialRETIRED (0040,A174) VR=CS VM=1 Referenced Object Observation Class (Trial) (RETIRED)
	ReferencedObjectObservationClassTrialRETIRED = New(0x0040, 0xA174)

	// AnnotationGroupNumber (0040,A180) VR=US VM=1 Annotation Group Number
	AnnotationGroupNumber = New(0x0040, 0xA180)

	// ObservationDateTrialRETIRED (0040,A192) VR=DA VM=1 Observation Date (Trial) (RETIRED)
	ObservationDateTrialRETIRED = New(0x0040, 0xA192)

	// ObservationTimeTrialRETIRED (0040,A193) VR=TM VM=1 Observation Time (Trial) (RETIRED)
	ObservationTimeTrialRETIRED = New(0x0040, 0xA193)

	// MeasurementAutomationTrialRETIRED (0040,A194) VR=CS VM=1 Measurement Automation (Trial) (RETIRED)
	MeasurementAutomationTrialRETIRED = New(0x0040, 0xA194)

	// ModifierCodeSequence (0040,A195) VR=SQ VM=1 Modifier Code Sequence
	ModifierCodeSequence = New(0x0040, 0xA195)

	// IdentificationDescriptionTrialRETIRED (0040,A224) VR=ST VM=1 Identification Description (Trial) (RETIRED)
	IdentificationDescriptionTrialRETIRED = New(0x0040, 0xA224)

	// CoordinatesSetGeometricTypeTrialRETIRED (0040,A290) VR=CS VM=1 Coordinates Set Geometric Type (Trial) (RETIRED)
	CoordinatesSetGeometricTypeTrialRETIRED = New(0x0040, 0xA290)

	// AlgorithmCodeSequenceTrialRETIRED (0040,A296) VR=SQ VM=1 Algorithm Code Sequence (Trial) (RETIRED)
	AlgorithmCodeSequenceTrialRETIRED = New(0x0040, 0xA296)

	// AlgorithmDescriptionTrialRETIRED (0040,A297) VR=ST VM=1 Algorithm Description (Trial) (RETIRED)
	AlgorithmDescriptionTrialRETIRED = New(0x0040, 0xA297)

	// PixelCoordinatesSetTrialRETIRED (0040,A29A) VR=SL VM=2-2n Pixel Coordinates Set (Trial) (RETIRED)
	PixelCoordinatesSetTrialRETIRED = New(0x0040, 0xA29A)

	// MeasuredValueSequence (0040,A300) VR=SQ VM=1 Measured Value Sequence
	MeasuredValueSequence = New(0x0040, 0xA300)

	// NumericValueQualifierCodeSequence (0040,A301) VR=SQ VM=1 Numeric Value Qualifier Code Sequence
	NumericValueQualifierCodeSequence = New(0x0040, 0xA301)

	// CurrentObserverTrialRETIRED (0040,A307) VR=PN VM=1 Current Observer (Trial) (RETIRED)
	CurrentObserverTrialRETIRED = New(0x0040, 0xA307)

	// NumericValue (0040,A30A) VR=DS VM=1-n Numeric Value
	NumericValue = New(0x0040, 0xA30A)

	// ReferencedAccessionSequenceTrialRETIRED (0040,A313) VR=SQ VM=1 Referenced Accession Sequence (Trial) (RETIRED)
	ReferencedAccessionSequenceTrialRETIRED = New(0x0040, 0xA313)

	// ReportStatusCommentTrialRETIRED (0040,A33A) VR=ST VM=1 Report Status Comment (Trial) (RETIRED)
	ReportStatusCommentTrialRETIRED = New(0x0040, 0xA33A)

	// ProcedureContextSequenceTrialRETIRED (0040,A340) VR=SQ VM=1 Procedure Context Sequence (Trial) (RETIRED)
	ProcedureContextSequenceTrialRETIRED = New(0x0040, 0xA340)

	// VerbalSourceTrialRETIRED (0040,A352) VR=PN VM=1 Verbal Source (Trial) (RETIRED)
	VerbalSourceTrialRETIRED = New(0x0040, 0xA352)

	// AddressTrialRETIRED (0040,A353) VR=ST VM=1 Address (Trial) (RETIRED)
	AddressTrialRETIRED = New(0x0040, 0xA353)

	// TelephoneNumberTrialRETIRED (0040,A354) VR=LO VM=1 Telephone Number (Trial) (RETIRED)
	TelephoneNumberTrialRETIRED = New(0x0040, 0xA354)

	// VerbalSourceIdentifierCodeSequenceTrialRETIRED (0040,A358) VR=SQ VM=1 Verbal Source Identifier Code Sequence (Trial) (RETIRED)
	VerbalSourceIdentifierCodeSequenceTrialRETIRED = New(0x0040, 0xA358)

	// PredecessorDocumentsSequence (0040,A360) VR=SQ VM=1 Predecessor Documents Sequence
	PredecessorDocumentsSequence = New(0x0040, 0xA360)

	// ReferencedRequestSequence (0040,A370) VR=SQ VM=1 Referenced Request Sequence
	ReferencedRequestSequence = New(0x0040, 0xA370)

	// PerformedProcedureCodeSequence (0040,A372) VR=SQ VM=1 Performed Procedure Code Sequence
	PerformedProcedureCodeSequence = New(0x0040, 0xA372)

	// CurrentRequestedProcedureEvidenceSequence (0040,A375) VR=SQ VM=1 Current Requested Procedure Evidence Sequence
	CurrentRequestedProcedureEvidenceSequence = New(0x0040, 0xA375)

	// ReportDetailSequenceTrialRETIRED (0040,A380) VR=SQ VM=1 Report Detail Sequence (Trial) (RETIRED)
	ReportDetailSequenceTrialRETIRED = New(0x0040, 0xA380)

	// PertinentOtherEvidenceSequence (0040,A385) VR=SQ VM=1 Pertinent Other Evidence Sequence
	PertinentOtherEvidenceSequence = New(0x0040, 0xA385)

	// HL7StructuredDocumentReferenceSequence (0040,A390) VR=SQ VM=1 HL7 Structured Document Reference Sequence
	HL7StructuredDocumentReferenceSequence = New(0x0040, 0xA390)

	// ObservationSubjectUIDTrialRETIRED (0040,A402) VR=UI VM=1 Observation Subject UID (Trial) (RETIRED)
	ObservationSubjectUIDTrialRETIRED = New(0x0040, 0xA402)

	// ObservationSubjectClassTrialRETIRED (0040,A403) VR=CS VM=1 Observation Subject Class (Trial) (RETIRED)
	ObservationSubjectClassTrialRETIRED = New(0x0040, 0xA403)

	// ObservationSubjectTypeCodeSequenceTrialRETIRED (0040,A404) VR=SQ VM=1 Observation Subject Type Code Sequence (Trial) (RETIRED)
	ObservationSubjectTypeCodeSequenceTrialRETIRED = New(0x0040, 0xA404)

	// CompletionFlag (0040,A491) VR=CS VM=1 Completion Flag
	CompletionFlag = New(0x0040, 0xA491)

	// CompletionFlagDescription (0040,A492) VR=LO VM=1 Completion Flag Description
	CompletionFlagDescription = New(0x0040, 0xA492)

	// VerificationFlag (0040,A493) VR=CS VM=1 Verification Flag
	VerificationFlag = New(0x0040, 0xA493)

	// ArchiveRequested (0040,A494) VR=CS VM=1 Archive Requested
	ArchiveRequested = New(0x0040, 0xA494)

	// PreliminaryFlag (0040,A496) VR=CS VM=1 Preliminary Flag
	PreliminaryFlag = New(0x0040, 0xA496)

	// ContentTemplateSequence (0040,A504) VR=SQ VM=1 Content Template Sequence
	ContentTemplateSequence = New(0x0040, 0xA504)

	// IdenticalDocumentsSequence (0040,A525) VR=SQ VM=1 Identical Documents Sequence
	IdenticalDocumentsSequence = New(0x0040, 0xA525)

	// ObservationSubjectContextFlagTrialRETIRED (0040,A600) VR=CS VM=1 Observation Subject Context Flag (Trial) (RETIRED)
	ObservationSubjectContextFlagTrialRETIRED = New(0x0040, 0xA600)

	// ObserverContextFlagTrialRETIRED (0040,A601) VR=CS VM=1 Observer Context Flag (Trial) (RETIRED)
	ObserverContextFlagTrialRETIRED = New(0x0040, 0xA601)

	// ProcedureContextFlagTrialRETIRED (0040,A603) VR=CS VM=1 Procedure Context Flag (Trial) (RETIRED)
	ProcedureContextFlagTrialRETIRED = New(0x0040, 0xA603)

	// ContentSequence (0040,A730) VR=SQ VM=1 Content Sequence
	ContentSequence = New(0x0040, 0xA730)

	// RelationshipSequenceTrialRETIRED (0040,A731) VR=SQ VM=1 Relationship Sequence (Trial) (RETIRED)
	RelationshipSequenceTrialRETIRED = New(0x0040, 0xA731)

	// RelationshipTypeCodeSequenceTrialRETIRED (0040,A732) VR=SQ VM=1 Relationship Type Code Sequence (Trial) (RETIRED)
	RelationshipTypeCodeSequenceTrialRETIRED = New(0x0040, 0xA732)

	// LanguageCodeSequenceTrialRETIRED (0040,A744) VR=SQ VM=1 Language Code Sequence (Trial) (RETIRED)
	LanguageCodeSequenceTrialRETIRED = New(0x0040, 0xA744)

	// TabulatedValuesSequence (0040,A801) VR=SQ VM=1 Tabulated Values Sequence
	TabulatedValuesSequence = New(0x0040, 0xA801)

	// NumberOfTableRows (0040,A802) VR=UL VM=1 Number of Table Rows
	NumberOfTableRows = New(0x0040, 0xA802)

	// NumberOfTableColumns (0040,A803) VR=UL VM=1 Number of Table Columns
	NumberOfTableColumns = New(0x0040, 0xA803)

	// TableRowNumber (0040,A804) VR=UL VM=1 Table Row Number
	TableRowNumber = New(0x0040, 0xA804)

	// TableColumnNumber (0040,A805) VR=UL VM=1 Table Column Number
	TableColumnNumber = New(0x0040, 0xA805)

	// TableRowDefinitionSequence (0040,A806) VR=SQ VM=1 Table Row Definition Sequence
	TableRowDefinitionSequence = New(0x0040, 0xA806)

	// TableColumnDefinitionSequence (0040,A807) VR=SQ VM=1 Table Column Definition Sequence
	TableColumnDefinitionSequence = New(0x0040, 0xA807)

	// CellValuesSequence (0040,A808) VR=SQ VM=1 Cell Values Sequence
	CellValuesSequence = New(0x0040, 0xA808)

	// UniformResourceLocatorTrialRETIRED (0040,A992) VR=ST VM=1 Uniform Resource Locator (Trial) (RETIRED)
	UniformResourceLocatorTrialRETIRED = New(0x0040, 0xA992)

	// WaveformAnnotationSequence (0040,B020) VR=SQ VM=1 Waveform Annotation Sequence
	WaveformAnnotationSequence = New(0x0040, 0xB020)

	// StructuredWaveformAnnotationSequence (0040,B030) VR=SQ VM=1 Structured Waveform Annotation Sequence
	StructuredWaveformAnnotationSequence = New(0x0040, 0xB030)

	// WaveformAnnotationDisplaySelectionSequence (0040,B031) VR=SQ VM=1 Waveform Annotation Display Selection Sequence
	WaveformAnnotationDisplaySelectionSequence = New(0x0040, 0xB031)

	// ReferencedMontageIndex (0040,B032) VR=US VM=1 Referenced Montage Index
	ReferencedMontageIndex = New(0x0040, 0xB032)

	// WaveformTextualAnnotationSequence (0040,B033) VR=SQ VM=1 Waveform Textual Annotation Sequence
	WaveformTextualAnnotationSequence = New(0x0040, 0xB033)

	// AnnotationDateTime (0040,B034) VR=DT VM=1 Annotation DateTime
	AnnotationDateTime = New(0x0040, 0xB034)

	// DisplayedWaveformSegmentSequence (0040,B035) VR=SQ VM=1 Displayed Waveform Segment Sequence
	DisplayedWaveformSegmentSequence = New(0x0040, 0xB035)

	// SegmentDefinitionDateTime (0040,B036) VR=DT VM=1 Segment Definition DateTime
	SegmentDefinitionDateTime = New(0x0040, 0xB036)

	// MontageActivationSequence (0040,B037) VR=SQ VM=1 Montage Activation Sequence
	MontageActivationSequence = New(0x0040, 0xB037)

	// MontageActivationTimeOffset (0040,B038) VR=DS VM=1 Montage Activation Time Offset
	MontageActivationTimeOffset = New(0x0040, 0xB038)

	// WaveformMontageSequence (0040,B039) VR=SQ VM=1 Waveform Montage Sequence
	WaveformMontageSequence = New(0x0040, 0xB039)

	// ReferencedMontageChannelNumber (0040,B03A) VR=IS VM=1 Referenced Montage Channel Number
	ReferencedMontageChannelNumber = New(0x0040, 0xB03A)

	// MontageName (0040,B03B) VR=LT VM=1 Montage Name
	MontageName = New(0x0040, 0xB03B)

	// MontageChannelSequence (0040,B03C) VR=SQ VM=1 Montage Channel Sequence
	MontageChannelSequence = New(0x0040, 0xB03C)

	// MontageIndex (0040,B03D) VR=US VM=1 Montage Index
	MontageIndex = New(0x0040, 0xB03D)

	// MontageChannelNumber (0040,B03E) VR=IS VM=1 Montage Channel Number
	MontageChannelNumber = New(0x0040, 0xB03E)

	// MontageChannelLabel (0040,B03F) VR=LO VM=1 Montage Channel Label
	MontageChannelLabel = New(0x0040, 0xB03F)

	// MontageChannelSourceCodeSequence (0040,B040) VR=SQ VM=1 Montage Channel Source Code Sequence
	MontageChannelSourceCodeSequence = New(0x0040, 0xB040)

	// ContributingChannelSourcesSequence (0040,B041) VR=SQ VM=1 Contributing Channel Sources Sequence
	ContributingChannelSourcesSequence = New(0x0040, 0xB041)

	// ChannelWeight (0040,B042) VR=FL VM=1 Channel Weight
	ChannelWeight = New(0x0040, 0xB042)

	// TemplateIdentifier (0040,DB00) VR=CS VM=1 Template Identifier
	TemplateIdentifier = New(0x0040, 0xDB00)

	// TemplateVersionRETIRED (0040,DB06) VR=DT VM=1 Template Version (RETIRED)
	TemplateVersionRETIRED = New(0x0040, 0xDB06)

	// TemplateLocalVersionRETIRED (0040,DB07) VR=DT VM=1 Template Local Version (RETIRED)
	TemplateLocalVersionRETIRED = New(0x0040, 0xDB07)

	// TemplateExtensionFlagRETIRED (0040,DB0B) VR=CS VM=1 Template Extension Flag (RETIRED)
	TemplateExtensionFlagRETIRED = New(0x0040, 0xDB0B)

	// TemplateExtensionOrganizationUIDRETIRED (0040,DB0C) VR=UI VM=1 Template Extension Organization UID (RETIRED)
	TemplateExtensionOrganizationUIDRETIRED = New(0x0040, 0xDB0C)

	// TemplateExtensionCreatorUIDRETIRED (0040,DB0D) VR=UI VM=1 Template Extension Creator UID (RETIRED)
	TemplateExtensionCreatorUIDRETIRED = New(0x0040, 0xDB0D)

	// ReferencedContentItemIdentifier (0040,DB73) VR=UL VM=1-n Referenced Content Item Identifier
	ReferencedContentItemIdentifier = New(0x0040, 0xDB73)

	// HL7InstanceIdentifier (0040,E001) VR=ST VM=1 HL7 Instance Identifier
	HL7InstanceIdentifier = New(0x0040, 0xE001)

	// HL7DocumentEffectiveTime (0040,E004) VR=DT VM=1 HL7 Document Effective Time
	HL7DocumentEffectiveTime = New(0x0040, 0xE004)

	// HL7DocumentTypeCodeSequence (0040,E006) VR=SQ VM=1 HL7 Document Type Code Sequence
	HL7DocumentTypeCodeSequence = New(0x0040, 0xE006)

	// DocumentClassCodeSequence (0040,E008) VR=SQ VM=1 Document Class Code Sequence
	DocumentClassCodeSequence = New(0x0040, 0xE008)

	// RetrieveURI (0040,E010) VR=UR VM=1 Retrieve URI
	RetrieveURI = New(0x0040, 0xE010)

	// RetrieveLocationUID (0040,E011) VR=UI VM=1 Retrieve Location UID
	RetrieveLocationUID = New(0x0040, 0xE011)

	// TypeOfInstances (0040,E020) VR=CS VM=1 Type of Instances
	TypeOfInstances = New(0x0040, 0xE020)

	// DICOMRetrievalSequence (0040,E021) VR=SQ VM=1 DICOM Retrieval Sequence
	DICOMRetrievalSequence = New(0x0040, 0xE021)

	// DICOMMediaRetrievalSequence (0040,E022) VR=SQ VM=1 DICOM Media Retrieval Sequence
	DICOMMediaRetrievalSequence = New(0x0040, 0xE022)

	// WADORetrievalSequence (0040,E023) VR=SQ VM=1 WADO Retrieval Sequence
	WADORetrievalSequence = New(0x0040, 0xE023)

	// XDSRetrievalSequence (0040,E024) VR=SQ VM=1 XDS Retrieval Sequence
	XDSRetrievalSequence = New(0x0040, 0xE024)

	// WADORSRetrievalSequence (0040,E025) VR=SQ VM=1 WADO-RS Retrieval Sequence
	WADORSRetrievalSequence = New(0x0040, 0xE025)

	// RepositoryUniqueID (0040,E030) VR=UI VM=1 Repository Unique ID
	RepositoryUniqueID = New(0x0040, 0xE030)

	// HomeCommunityID (0040,E031) VR=UI VM=1 Home Community ID
	HomeCommunityID = New(0x0040, 0xE031)

	// DocumentTitle (0042,0010) VR=ST VM=1 Document Title
	DocumentTitle = New(0x0042, 0x0010)

	// EncapsulatedDocument (0042,0011) VR=OB VM=1 Encapsulated Document
	EncapsulatedDocument = New(0x0042, 0x0011)

	// MIMETypeOfEncapsulatedDocument (0042,0012) VR=LO VM=1 MIME Type of Encapsulated Document
	MIMETypeOfEncapsulatedDocument = New(0x0042, 0x0012)

	// SourceInstanceSequence (0042,0013) VR=SQ VM=1 Source Instance Sequence
	SourceInstanceSequence = New(0x0042, 0x0013)

	// ListOfMIMETypes (0042,0014) VR=LO VM=1-n List of MIME Types
	ListOfMIMETypes = New(0x0042, 0x0014)

	// EncapsulatedDocumentLength (0042,0015) VR=UL VM=1 Encapsulated Document Length
	EncapsulatedDocumentLength = New(0x0042, 0x0015)

	// ProductPackageIdentifier (0044,0001) VR=ST VM=1 Product Package Identifier
	ProductPackageIdentifier = New(0x0044, 0x0001)

	// SubstanceAdministrationApproval (0044,0002) VR=CS VM=1 Substance Administration Approval
	SubstanceAdministrationApproval = New(0x0044, 0x0002)

	// ApprovalStatusFurtherDescription (0044,0003) VR=LT VM=1 Approval Status Further Description
	ApprovalStatusFurtherDescription = New(0x0044, 0x0003)

	// ApprovalStatusDateTime (0044,0004) VR=DT VM=1 Approval Status DateTime
	ApprovalStatusDateTime = New(0x0044, 0x0004)

	// ProductTypeCodeSequence (0044,0007) VR=SQ VM=1 Product Type Code Sequence
	ProductTypeCodeSequence = New(0x0044, 0x0007)

	// ProductName (0044,0008) VR=LO VM=1-n Product Name
	ProductName = New(0x0044, 0x0008)

	// ProductDescription (0044,0009) VR=LT VM=1 Product Description
	ProductDescription = New(0x0044, 0x0009)

	// ProductLotIdentifier (0044,000A) VR=LO VM=1 Product Lot Identifier
	ProductLotIdentifier = New(0x0044, 0x000A)

	// ProductExpirationDateTime (0044,000B) VR=DT VM=1 Product Expiration DateTime
	ProductExpirationDateTime = New(0x0044, 0x000B)

	// SubstanceAdministrationDateTime (0044,0010) VR=DT VM=1 Substance Administration DateTime
	SubstanceAdministrationDateTime = New(0x0044, 0x0010)

	// SubstanceAdministrationNotes (0044,0011) VR=LO VM=1 Substance Administration Notes
	SubstanceAdministrationNotes = New(0x0044, 0x0011)

	// SubstanceAdministrationDeviceID (0044,0012) VR=LO VM=1 Substance Administration Device ID
	SubstanceAdministrationDeviceID = New(0x0044, 0x0012)

	// ProductParameterSequence (0044,0013) VR=SQ VM=1 Product Parameter Sequence
	ProductParameterSequence = New(0x0044, 0x0013)

	// SubstanceAdministrationParameterSequence (0044,0019) VR=SQ VM=1 Substance Administration Parameter Sequence
	SubstanceAdministrationParameterSequence = New(0x0044, 0x0019)

	// ApprovalSequence (0044,0100) VR=SQ VM=1 Approval Sequence
	ApprovalSequence = New(0x0044, 0x0100)

	// AssertionCodeSequence (0044,0101) VR=SQ VM=1 Assertion Code Sequence
	AssertionCodeSequence = New(0x0044, 0x0101)

	// AssertionUID (0044,0102) VR=UI VM=1 Assertion UID
	AssertionUID = New(0x0044, 0x0102)

	// AsserterIdentificationSequence (0044,0103) VR=SQ VM=1 Asserter Identification Sequence
	AsserterIdentificationSequence = New(0x0044, 0x0103)

	// AssertionDateTime (0044,0104) VR=DT VM=1 Assertion DateTime
	AssertionDateTime = New(0x0044, 0x0104)

	// AssertionExpirationDateTime (0044,0105) VR=DT VM=1 Assertion Expiration DateTime
	AssertionExpirationDateTime = New(0x0044, 0x0105)

	// AssertionComments (0044,0106) VR=UT VM=1 Assertion Comments
	AssertionComments = New(0x0044, 0x0106)

	// RelatedAssertionSequence (0044,0107) VR=SQ VM=1 Related Assertion Sequence
	RelatedAssertionSequence = New(0x0044, 0x0107)

	// ReferencedAssertionUID (0044,0108) VR=UI VM=1 Referenced Assertion UID
	ReferencedAssertionUID = New(0x0044, 0x0108)

	// ApprovalSubjectSequence (0044,0109) VR=SQ VM=1 Approval Subject Sequence
	ApprovalSubjectSequence = New(0x0044, 0x0109)

	// OrganizationalRoleCodeSequence (0044,010A) VR=SQ VM=1 Organizational Role Code Sequence
	OrganizationalRoleCodeSequence = New(0x0044, 0x010A)

	// RTAssertionsSequence (0044,0110) VR=SQ VM=1 RT Assertions Sequence
	RTAssertionsSequence = New(0x0044, 0x0110)

	// LensDescription (0046,0012) VR=LO VM=1 Lens Description
	LensDescription = New(0x0046, 0x0012)

	// RightLensSequence (0046,0014) VR=SQ VM=1 Right Lens Sequence
	RightLensSequence = New(0x0046, 0x0014)

	// LeftLensSequence (0046,0015) VR=SQ VM=1 Left Lens Sequence
	LeftLensSequence = New(0x0046, 0x0015)

	// UnspecifiedLateralityLensSequence (0046,0016) VR=SQ VM=1 Unspecified Laterality Lens Sequence
	UnspecifiedLateralityLensSequence = New(0x0046, 0x0016)

	// CylinderSequence (0046,0018) VR=SQ VM=1 Cylinder Sequence
	CylinderSequence = New(0x0046, 0x0018)

	// PrismSequence (0046,0028) VR=SQ VM=1 Prism Sequence
	PrismSequence = New(0x0046, 0x0028)

	// HorizontalPrismPower (0046,0030) VR=FD VM=1 Horizontal Prism Power
	HorizontalPrismPower = New(0x0046, 0x0030)

	// HorizontalPrismBase (0046,0032) VR=CS VM=1 Horizontal Prism Base
	HorizontalPrismBase = New(0x0046, 0x0032)

	// VerticalPrismPower (0046,0034) VR=FD VM=1 Vertical Prism Power
	VerticalPrismPower = New(0x0046, 0x0034)

	// VerticalPrismBase (0046,0036) VR=CS VM=1 Vertical Prism Base
	VerticalPrismBase = New(0x0046, 0x0036)

	// LensSegmentType (0046,0038) VR=CS VM=1 Lens Segment Type
	LensSegmentType = New(0x0046, 0x0038)

	// OpticalTransmittance (0046,0040) VR=FD VM=1 Optical Transmittance
	OpticalTransmittance = New(0x0046, 0x0040)

	// ChannelWidth (0046,0042) VR=FD VM=1 Channel Width
	ChannelWidth = New(0x0046, 0x0042)

	// PupilSize (0046,0044) VR=FD VM=1 Pupil Size
	PupilSize = New(0x0046, 0x0044)

	// CornealSize (0046,0046) VR=FD VM=1 Corneal Size
	CornealSize = New(0x0046, 0x0046)

	// CornealSizeSequence (0046,0047) VR=SQ VM=1 Corneal Size Sequence
	CornealSizeSequence = New(0x0046, 0x0047)

	// AutorefractionRightEyeSequence (0046,0050) VR=SQ VM=1 Autorefraction Right Eye Sequence
	AutorefractionRightEyeSequence = New(0x0046, 0x0050)

	// AutorefractionLeftEyeSequence (0046,0052) VR=SQ VM=1 Autorefraction Left Eye Sequence
	AutorefractionLeftEyeSequence = New(0x0046, 0x0052)

	// DistancePupillaryDistance (0046,0060) VR=FD VM=1 Distance Pupillary Distance
	DistancePupillaryDistance = New(0x0046, 0x0060)

	// NearPupillaryDistance (0046,0062) VR=FD VM=1 Near Pupillary Distance
	NearPupillaryDistance = New(0x0046, 0x0062)

	// IntermediatePupillaryDistance (0046,0063) VR=FD VM=1 Intermediate Pupillary Distance
	IntermediatePupillaryDistance = New(0x0046, 0x0063)

	// OtherPupillaryDistance (0046,0064) VR=FD VM=1 Other Pupillary Distance
	OtherPupillaryDistance = New(0x0046, 0x0064)

	// KeratometryRightEyeSequence (0046,0070) VR=SQ VM=1 Keratometry Right Eye Sequence
	KeratometryRightEyeSequence = New(0x0046, 0x0070)

	// KeratometryLeftEyeSequence (0046,0071) VR=SQ VM=1 Keratometry Left Eye Sequence
	KeratometryLeftEyeSequence = New(0x0046, 0x0071)

	// SteepKeratometricAxisSequence (0046,0074) VR=SQ VM=1 Steep Keratometric Axis Sequence
	SteepKeratometricAxisSequence = New(0x0046, 0x0074)

	// RadiusOfCurvature (0046,0075) VR=FD VM=1 Radius of Curvature
	RadiusOfCurvature = New(0x0046, 0x0075)

	// KeratometricPower (0046,0076) VR=FD VM=1 Keratometric Power
	KeratometricPower = New(0x0046, 0x0076)

	// KeratometricAxis (0046,0077) VR=FD VM=1 Keratometric Axis
	KeratometricAxis = New(0x0046, 0x0077)

	// FlatKeratometricAxisSequence (0046,0080) VR=SQ VM=1 Flat Keratometric Axis Sequence
	FlatKeratometricAxisSequence = New(0x0046, 0x0080)

	// BackgroundColor (0046,0092) VR=CS VM=1 Background Color
	BackgroundColor = New(0x0046, 0x0092)

	// Optotype (0046,0094) VR=CS VM=1 Optotype
	Optotype = New(0x0046, 0x0094)

	// OptotypePresentation (0046,0095) VR=CS VM=1 Optotype Presentation
	OptotypePresentation = New(0x0046, 0x0095)

	// SubjectiveRefractionRightEyeSequence (0046,0097) VR=SQ VM=1 Subjective Refraction Right Eye Sequence
	SubjectiveRefractionRightEyeSequence = New(0x0046, 0x0097)

	// SubjectiveRefractionLeftEyeSequence (0046,0098) VR=SQ VM=1 Subjective Refraction Left Eye Sequence
	SubjectiveRefractionLeftEyeSequence = New(0x0046, 0x0098)

	// AddNearSequence (0046,0100) VR=SQ VM=1 Add Near Sequence
	AddNearSequence = New(0x0046, 0x0100)

	// AddIntermediateSequence (0046,0101) VR=SQ VM=1 Add Intermediate Sequence
	AddIntermediateSequence = New(0x0046, 0x0101)

	// AddOtherSequence (0046,0102) VR=SQ VM=1 Add Other Sequence
	AddOtherSequence = New(0x0046, 0x0102)

	// AddPower (0046,0104) VR=FD VM=1 Add Power
	AddPower = New(0x0046, 0x0104)

	// ViewingDistance (0046,0106) VR=FD VM=1 Viewing Distance
	ViewingDistance = New(0x0046, 0x0106)

	// CorneaMeasurementsSequence (0046,0110) VR=SQ VM=1 Cornea Measurements Sequence
	CorneaMeasurementsSequence = New(0x0046, 0x0110)

	// SourceOfCorneaMeasurementDataCodeSequence (0046,0111) VR=SQ VM=1 Source of Cornea Measurement Data Code Sequence
	SourceOfCorneaMeasurementDataCodeSequence = New(0x0046, 0x0111)

	// SteepCornealAxisSequence (0046,0112) VR=SQ VM=1 Steep Corneal Axis Sequence
	SteepCornealAxisSequence = New(0x0046, 0x0112)

	// FlatCornealAxisSequence (0046,0113) VR=SQ VM=1 Flat Corneal Axis Sequence
	FlatCornealAxisSequence = New(0x0046, 0x0113)

	// CornealPower (0046,0114) VR=FD VM=1 Corneal Power
	CornealPower = New(0x0046, 0x0114)

	// CornealAxis (0046,0115) VR=FD VM=1 Corneal Axis
	CornealAxis = New(0x0046, 0x0115)

	// CorneaMeasurementMethodCodeSequence (0046,0116) VR=SQ VM=1 Cornea Measurement Method Code Sequence
	CorneaMeasurementMethodCodeSequence = New(0x0046, 0x0116)

	// RefractiveIndexOfCornea (0046,0117) VR=FL VM=1 Refractive Index of Cornea
	RefractiveIndexOfCornea = New(0x0046, 0x0117)

	// RefractiveIndexOfAqueousHumor (0046,0118) VR=FL VM=1 Refractive Index of Aqueous Humor
	RefractiveIndexOfAqueousHumor = New(0x0046, 0x0118)

	// VisualAcuityTypeCodeSequence (0046,0121) VR=SQ VM=1 Visual Acuity Type Code Sequence
	VisualAcuityTypeCodeSequence = New(0x0046, 0x0121)

	// VisualAcuityRightEyeSequence (0046,0122) VR=SQ VM=1 Visual Acuity Right Eye Sequence
	VisualAcuityRightEyeSequence = New(0x0046, 0x0122)

	// VisualAcuityLeftEyeSequence (0046,0123) VR=SQ VM=1 Visual Acuity Left Eye Sequence
	VisualAcuityLeftEyeSequence = New(0x0046, 0x0123)

	// VisualAcuityBothEyesOpenSequence (0046,0124) VR=SQ VM=1 Visual Acuity Both Eyes Open Sequence
	VisualAcuityBothEyesOpenSequence = New(0x0046, 0x0124)

	// ViewingDistanceType (0046,0125) VR=CS VM=1 Viewing Distance Type
	ViewingDistanceType = New(0x0046, 0x0125)

	// VisualAcuityModifiers (0046,0135) VR=SS VM=2 Visual Acuity Modifiers
	VisualAcuityModifiers = New(0x0046, 0x0135)

	// DecimalVisualAcuity (0046,0137) VR=FD VM=1 Decimal Visual Acuity
	DecimalVisualAcuity = New(0x0046, 0x0137)

	// OptotypeDetailedDefinition (0046,0139) VR=LO VM=1 Optotype Detailed Definition
	OptotypeDetailedDefinition = New(0x0046, 0x0139)

	// ReferencedRefractiveMeasurementsSequence (0046,0145) VR=SQ VM=1 Referenced Refractive Measurements Sequence
	ReferencedRefractiveMeasurementsSequence = New(0x0046, 0x0145)

	// SpherePower (0046,0146) VR=FD VM=1 Sphere Power
	SpherePower = New(0x0046, 0x0146)

	// CylinderPower (0046,0147) VR=FD VM=1 Cylinder Power
	CylinderPower = New(0x0046, 0x0147)

	// CornealTopographySurface (0046,0201) VR=CS VM=1 Corneal Topography Surface
	CornealTopographySurface = New(0x0046, 0x0201)

	// CornealVertexLocation (0046,0202) VR=FL VM=2 Corneal Vertex Location
	CornealVertexLocation = New(0x0046, 0x0202)

	// PupilCentroidXCoordinate (0046,0203) VR=FL VM=1 Pupil Centroid X-Coordinate
	PupilCentroidXCoordinate = New(0x0046, 0x0203)

	// PupilCentroidYCoordinate (0046,0204) VR=FL VM=1 Pupil Centroid Y-Coordinate
	PupilCentroidYCoordinate = New(0x0046, 0x0204)

	// EquivalentPupilRadius (0046,0205) VR=FL VM=1 Equivalent Pupil Radius
	EquivalentPupilRadius = New(0x0046, 0x0205)

	// CornealTopographyMapTypeCodeSequence (0046,0207) VR=SQ VM=1 Corneal Topography Map Type Code Sequence
	CornealTopographyMapTypeCodeSequence = New(0x0046, 0x0207)

	// VerticesOfTheOutlineOfPupil (0046,0208) VR=IS VM=2-2n Vertices of the Outline of Pupil
	VerticesOfTheOutlineOfPupil = New(0x0046, 0x0208)

	// CornealTopographyMappingNormalsSequence (0046,0210) VR=SQ VM=1 Corneal Topography Mapping Normals Sequence
	CornealTopographyMappingNormalsSequence = New(0x0046, 0x0210)

	// MaximumCornealCurvatureSequence (0046,0211) VR=SQ VM=1 Maximum Corneal Curvature Sequence
	MaximumCornealCurvatureSequence = New(0x0046, 0x0211)

	// MaximumCornealCurvature (0046,0212) VR=FL VM=1 Maximum Corneal Curvature
	MaximumCornealCurvature = New(0x0046, 0x0212)

	// MaximumCornealCurvatureLocation (0046,0213) VR=FL VM=2 Maximum Corneal Curvature Location
	MaximumCornealCurvatureLocation = New(0x0046, 0x0213)

	// MinimumKeratometricSequence (0046,0215) VR=SQ VM=1 Minimum Keratometric Sequence
	MinimumKeratometricSequence = New(0x0046, 0x0215)

	// SimulatedKeratometricCylinderSequence (0046,0218) VR=SQ VM=1 Simulated Keratometric Cylinder Sequence
	SimulatedKeratometricCylinderSequence = New(0x0046, 0x0218)

	// AverageCornealPower (0046,0220) VR=FL VM=1 Average Corneal Power
	AverageCornealPower = New(0x0046, 0x0220)

	// CornealISValue (0046,0224) VR=FL VM=1 Corneal I-S Value
	CornealISValue = New(0x0046, 0x0224)

	// AnalyzedArea (0046,0227) VR=FL VM=1 Analyzed Area
	AnalyzedArea = New(0x0046, 0x0227)

	// SurfaceRegularityIndex (0046,0230) VR=FL VM=1 Surface Regularity Index
	SurfaceRegularityIndex = New(0x0046, 0x0230)

	// SurfaceAsymmetryIndex (0046,0232) VR=FL VM=1 Surface Asymmetry Index
	SurfaceAsymmetryIndex = New(0x0046, 0x0232)

	// CornealEccentricityIndex (0046,0234) VR=FL VM=1 Corneal Eccentricity Index
	CornealEccentricityIndex = New(0x0046, 0x0234)

	// KeratoconusPredictionIndex (0046,0236) VR=FL VM=1 Keratoconus Prediction Index
	KeratoconusPredictionIndex = New(0x0046, 0x0236)

	// DecimalPotentialVisualAcuity (0046,0238) VR=FL VM=1 Decimal Potential Visual Acuity
	DecimalPotentialVisualAcuity = New(0x0046, 0x0238)

	// CornealTopographyMapQualityEvaluation (0046,0242) VR=CS VM=1 Corneal Topography Map Quality Evaluation
	CornealTopographyMapQualityEvaluation = New(0x0046, 0x0242)

	// SourceImageCornealProcessedDataSequence (0046,0244) VR=SQ VM=1 Source Image Corneal Processed Data Sequence
	SourceImageCornealProcessedDataSequence = New(0x0046, 0x0244)

	// CornealPointLocation (0046,0247) VR=FL VM=3 Corneal Point Location
	CornealPointLocation = New(0x0046, 0x0247)

	// CornealPointEstimated (0046,0248) VR=CS VM=1 Corneal Point Estimated
	CornealPointEstimated = New(0x0046, 0x0248)

	// AxialPower (0046,0249) VR=FL VM=1 Axial Power
	AxialPower = New(0x0046, 0x0249)

	// TangentialPower (0046,0250) VR=FL VM=1 Tangential Power
	TangentialPower = New(0x0046, 0x0250)

	// RefractivePower (0046,0251) VR=FL VM=1 Refractive Power
	RefractivePower = New(0x0046, 0x0251)

	// RelativeElevation (0046,0252) VR=FL VM=1 Relative Elevation
	RelativeElevation = New(0x0046, 0x0252)

	// CornealWavefront (0046,0253) VR=FL VM=1 Corneal Wavefront
	CornealWavefront = New(0x0046, 0x0253)

	// ImagedVolumeWidth (0048,0001) VR=FL VM=1 Imaged Volume Width
	ImagedVolumeWidth = New(0x0048, 0x0001)

	// ImagedVolumeHeight (0048,0002) VR=FL VM=1 Imaged Volume Height
	ImagedVolumeHeight = New(0x0048, 0x0002)

	// ImagedVolumeDepth (0048,0003) VR=FL VM=1 Imaged Volume Depth
	ImagedVolumeDepth = New(0x0048, 0x0003)

	// TotalPixelMatrixColumns (0048,0006) VR=UL VM=1 Total Pixel Matrix Columns
	TotalPixelMatrixColumns = New(0x0048, 0x0006)

	// TotalPixelMatrixRows (0048,0007) VR=UL VM=1 Total Pixel Matrix Rows
	TotalPixelMatrixRows = New(0x0048, 0x0007)

	// TotalPixelMatrixOriginSequence (0048,0008) VR=SQ VM=1 Total Pixel Matrix Origin Sequence
	TotalPixelMatrixOriginSequence = New(0x0048, 0x0008)

	// SpecimenLabelInImage (0048,0010) VR=CS VM=1 Specimen Label in Image
	SpecimenLabelInImage = New(0x0048, 0x0010)

	// FocusMethod (0048,0011) VR=CS VM=1 Focus Method
	FocusMethod = New(0x0048, 0x0011)

	// ExtendedDepthOfField (0048,0012) VR=CS VM=1 Extended Depth of Field
	ExtendedDepthOfField = New(0x0048, 0x0012)

	// NumberOfFocalPlanes (0048,0013) VR=US VM=1 Number of Focal Planes
	NumberOfFocalPlanes = New(0x0048, 0x0013)

	// DistanceBetweenFocalPlanes (0048,0014) VR=FL VM=1 Distance Between Focal Planes
	DistanceBetweenFocalPlanes = New(0x0048, 0x0014)

	// RecommendedAbsentPixelCIELabValue (0048,0015) VR=US VM=3 Recommended Absent Pixel CIELab Value
	RecommendedAbsentPixelCIELabValue = New(0x0048, 0x0015)

	// IlluminatorTypeCodeSequence (0048,0100) VR=SQ VM=1 Illuminator Type Code Sequence
	IlluminatorTypeCodeSequence = New(0x0048, 0x0100)

	// ImageOrientationSlide (0048,0102) VR=DS VM=6 Image Orientation (Slide)
	ImageOrientationSlide = New(0x0048, 0x0102)

	// OpticalPathSequence (0048,0105) VR=SQ VM=1 Optical Path Sequence
	OpticalPathSequence = New(0x0048, 0x0105)

	// OpticalPathIdentifier (0048,0106) VR=SH VM=1 Optical Path Identifier
	OpticalPathIdentifier = New(0x0048, 0x0106)

	// OpticalPathDescription (0048,0107) VR=ST VM=1 Optical Path Description
	OpticalPathDescription = New(0x0048, 0x0107)

	// IlluminationColorCodeSequence (0048,0108) VR=SQ VM=1 Illumination Color Code Sequence
	IlluminationColorCodeSequence = New(0x0048, 0x0108)

	// SpecimenReferenceSequence (0048,0110) VR=SQ VM=1 Specimen Reference Sequence
	SpecimenReferenceSequence = New(0x0048, 0x0110)

	// CondenserLensPower (0048,0111) VR=DS VM=1 Condenser Lens Power
	CondenserLensPower = New(0x0048, 0x0111)

	// ObjectiveLensPower (0048,0112) VR=DS VM=1 Objective Lens Power
	ObjectiveLensPower = New(0x0048, 0x0112)

	// ObjectiveLensNumericalAperture (0048,0113) VR=DS VM=1 Objective Lens Numerical Aperture
	ObjectiveLensNumericalAperture = New(0x0048, 0x0113)

	// ConfocalMode (0048,0114) VR=CS VM=1 Confocal Mode
	ConfocalMode = New(0x0048, 0x0114)

	// TissueLocation (0048,0115) VR=CS VM=1 Tissue Location
	TissueLocation = New(0x0048, 0x0115)

	// ConfocalMicroscopyImageFrameTypeSequence (0048,0116) VR=SQ VM=1 Confocal Microscopy Image Frame Type Sequence
	ConfocalMicroscopyImageFrameTypeSequence = New(0x0048, 0x0116)

	// ImageAcquisitionDepth (0048,0117) VR=FD VM=1 Image Acquisition Depth
	ImageAcquisitionDepth = New(0x0048, 0x0117)

	// PaletteColorLookupTableSequence (0048,0120) VR=SQ VM=1 Palette Color Lookup Table Sequence
	PaletteColorLookupTableSequence = New(0x0048, 0x0120)

	// ReferencedImageNavigationSequenceRETIRED (0048,0200) VR=SQ VM=1 Referenced Image Navigation Sequence (RETIRED)
	ReferencedImageNavigationSequenceRETIRED = New(0x0048, 0x0200)

	// TopLeftHandCornerOfLocalizerAreaRETIRED (0048,0201) VR=US VM=2 Top Left Hand Corner of Localizer Area (RETIRED)
	TopLeftHandCornerOfLocalizerAreaRETIRED = New(0x0048, 0x0201)

	// BottomRightHandCornerOfLocalizerAreaRETIRED (0048,0202) VR=US VM=2 Bottom Right Hand Corner of Localizer Area (RETIRED)
	BottomRightHandCornerOfLocalizerAreaRETIRED = New(0x0048, 0x0202)

	// OpticalPathIdentificationSequence (0048,0207) VR=SQ VM=1 Optical Path Identification Sequence
	OpticalPathIdentificationSequence = New(0x0048, 0x0207)

	// PlanePositionSlideSequence (0048,021A) VR=SQ VM=1 Plane Position (Slide) Sequence
	PlanePositionSlideSequence = New(0x0048, 0x021A)

	// ColumnPositionInTotalImagePixelMatrix (0048,021E) VR=SL VM=1 Column Position In Total Image Pixel Matrix
	ColumnPositionInTotalImagePixelMatrix = New(0x0048, 0x021E)

	// RowPositionInTotalImagePixelMatrix (0048,021F) VR=SL VM=1 Row Position In Total Image Pixel Matrix
	RowPositionInTotalImagePixelMatrix = New(0x0048, 0x021F)

	// PixelOriginInterpretation (0048,0301) VR=CS VM=1 Pixel Origin Interpretation
	PixelOriginInterpretation = New(0x0048, 0x0301)

	// NumberOfOpticalPaths (0048,0302) VR=UL VM=1 Number of Optical Paths
	NumberOfOpticalPaths = New(0x0048, 0x0302)

	// TotalPixelMatrixFocalPlanes (0048,0303) VR=UL VM=1 Total Pixel Matrix Focal Planes
	TotalPixelMatrixFocalPlanes = New(0x0048, 0x0303)

	// TilesOverlap (0048,0304) VR=CS VM=1 Tiles Overlap
	TilesOverlap = New(0x0048, 0x0304)

	// CalibrationImage (0050,0004) VR=CS VM=1 Calibration Image
	CalibrationImage = New(0x0050, 0x0004)

	// DeviceSequence (0050,0010) VR=SQ VM=1 Device Sequence
	DeviceSequence = New(0x0050, 0x0010)

	// ContainerComponentTypeCodeSequence (0050,0012) VR=SQ VM=1 Container Component Type Code Sequence
	ContainerComponentTypeCodeSequence = New(0x0050, 0x0012)

	// ContainerComponentThickness (0050,0013) VR=FD VM=1 Container Component Thickness
	ContainerComponentThickness = New(0x0050, 0x0013)

	// DeviceLength (0050,0014) VR=DS VM=1 Device Length
	DeviceLength = New(0x0050, 0x0014)

	// ContainerComponentWidth (0050,0015) VR=FD VM=1 Container Component Width
	ContainerComponentWidth = New(0x0050, 0x0015)

	// DeviceDiameter (0050,0016) VR=DS VM=1 Device Diameter
	DeviceDiameter = New(0x0050, 0x0016)

	// DeviceDiameterUnits (0050,0017) VR=CS VM=1 Device Diameter Units
	DeviceDiameterUnits = New(0x0050, 0x0017)

	// DeviceVolume (0050,0018) VR=DS VM=1 Device Volume
	DeviceVolume = New(0x0050, 0x0018)

	// InterMarkerDistance (0050,0019) VR=DS VM=1 Inter-Marker Distance
	InterMarkerDistance = New(0x0050, 0x0019)

	// ContainerComponentMaterial (0050,001A) VR=CS VM=1 Container Component Material
	ContainerComponentMaterial = New(0x0050, 0x001A)

	// ContainerComponentID (0050,001B) VR=LO VM=1 Container Component ID
	ContainerComponentID = New(0x0050, 0x001B)

	// ContainerComponentLength (0050,001C) VR=FD VM=1 Container Component Length
	ContainerComponentLength = New(0x0050, 0x001C)

	// ContainerComponentDiameter (0050,001D) VR=FD VM=1 Container Component Diameter
	ContainerComponentDiameter = New(0x0050, 0x001D)

	// ContainerComponentDescription (0050,001E) VR=LO VM=1 Container Component Description
	ContainerComponentDescription = New(0x0050, 0x001E)

	// DeviceDescription (0050,0020) VR=LO VM=1 Device Description
	DeviceDescription = New(0x0050, 0x0020)

	// LongDeviceDescription (0050,0021) VR=ST VM=1 Long Device Description
	LongDeviceDescription = New(0x0050, 0x0021)

	// ContrastBolusIngredientPercentByVolume (0052,0001) VR=FL VM=1 Contrast/Bolus Ingredient Percent by Volume
	ContrastBolusIngredientPercentByVolume = New(0x0052, 0x0001)

	// OCTFocalDistance (0052,0002) VR=FD VM=1 OCT Focal Distance
	OCTFocalDistance = New(0x0052, 0x0002)

	// BeamSpotSize (0052,0003) VR=FD VM=1 Beam Spot Size
	BeamSpotSize = New(0x0052, 0x0003)

	// EffectiveRefractiveIndex (0052,0004) VR=FD VM=1 Effective Refractive Index
	EffectiveRefractiveIndex = New(0x0052, 0x0004)

	// OCTAcquisitionDomain (0052,0006) VR=CS VM=1 OCT Acquisition Domain
	OCTAcquisitionDomain = New(0x0052, 0x0006)

	// OCTOpticalCenterWavelength (0052,0007) VR=FD VM=1 OCT Optical Center Wavelength
	OCTOpticalCenterWavelength = New(0x0052, 0x0007)

	// AxialResolution (0052,0008) VR=FD VM=1 Axial Resolution
	AxialResolution = New(0x0052, 0x0008)

	// RangingDepth (0052,0009) VR=FD VM=1 Ranging Depth
	RangingDepth = New(0x0052, 0x0009)

	// ALineRate (0052,0011) VR=FD VM=1 A-line Rate
	ALineRate = New(0x0052, 0x0011)

	// ALinesPerFrame (0052,0012) VR=US VM=1 A-lines Per Frame
	ALinesPerFrame = New(0x0052, 0x0012)

	// CatheterRotationalRate (0052,0013) VR=FD VM=1 Catheter Rotational Rate
	CatheterRotationalRate = New(0x0052, 0x0013)

	// ALinePixelSpacing (0052,0014) VR=FD VM=1 A-line Pixel Spacing
	ALinePixelSpacing = New(0x0052, 0x0014)

	// ModeOfPercutaneousAccessSequence (0052,0016) VR=SQ VM=1 Mode of Percutaneous Access Sequence
	ModeOfPercutaneousAccessSequence = New(0x0052, 0x0016)

	// IntravascularOCTFrameTypeSequence (0052,0025) VR=SQ VM=1 Intravascular OCT Frame Type Sequence
	IntravascularOCTFrameTypeSequence = New(0x0052, 0x0025)

	// OCTZOffsetApplied (0052,0026) VR=CS VM=1 OCT Z Offset Applied
	OCTZOffsetApplied = New(0x0052, 0x0026)

	// IntravascularFrameContentSequence (0052,0027) VR=SQ VM=1 Intravascular Frame Content Sequence
	IntravascularFrameContentSequence = New(0x0052, 0x0027)

	// IntravascularLongitudinalDistance (0052,0028) VR=FD VM=1 Intravascular Longitudinal Distance
	IntravascularLongitudinalDistance = New(0x0052, 0x0028)

	// IntravascularOCTFrameContentSequence (0052,0029) VR=SQ VM=1 Intravascular OCT Frame Content Sequence
	IntravascularOCTFrameContentSequence = New(0x0052, 0x0029)

	// OCTZOffsetCorrection (0052,0030) VR=SS VM=1 OCT Z Offset Correction
	OCTZOffsetCorrection = New(0x0052, 0x0030)

	// CatheterDirectionOfRotation (0052,0031) VR=CS VM=1 Catheter Direction of Rotation
	CatheterDirectionOfRotation = New(0x0052, 0x0031)

	// SeamLineLocation (0052,0033) VR=FD VM=1 Seam Line Location
	SeamLineLocation = New(0x0052, 0x0033)

	// FirstALineLocation (0052,0034) VR=FD VM=1 First A-line Location
	FirstALineLocation = New(0x0052, 0x0034)

	// SeamLineIndex (0052,0036) VR=US VM=1 Seam Line Index
	SeamLineIndex = New(0x0052, 0x0036)

	// NumberOfPaddedALines (0052,0038) VR=US VM=1 Number of Padded A-lines
	NumberOfPaddedALines = New(0x0052, 0x0038)

	// InterpolationType (0052,0039) VR=CS VM=1 Interpolation Type
	InterpolationType = New(0x0052, 0x0039)

	// RefractiveIndexApplied (0052,003A) VR=CS VM=1 Refractive Index Applied
	RefractiveIndexApplied = New(0x0052, 0x003A)

	// EnergyWindowVector (0054,0010) VR=US VM=1-n Energy Window Vector
	EnergyWindowVector = New(0x0054, 0x0010)

	// NumberOfEnergyWindows (0054,0011) VR=US VM=1 Number of Energy Windows
	NumberOfEnergyWindows = New(0x0054, 0x0011)

	// EnergyWindowInformationSequence (0054,0012) VR=SQ VM=1 Energy Window Information Sequence
	EnergyWindowInformationSequence = New(0x0054, 0x0012)

	// EnergyWindowRangeSequence (0054,0013) VR=SQ VM=1 Energy Window Range Sequence
	EnergyWindowRangeSequence = New(0x0054, 0x0013)

	// EnergyWindowLowerLimit (0054,0014) VR=DS VM=1 Energy Window Lower Limit
	EnergyWindowLowerLimit = New(0x0054, 0x0014)

	// EnergyWindowUpperLimit (0054,0015) VR=DS VM=1 Energy Window Upper Limit
	EnergyWindowUpperLimit = New(0x0054, 0x0015)

	// RadiopharmaceuticalInformationSequence (0054,0016) VR=SQ VM=1 Radiopharmaceutical Information Sequence
	RadiopharmaceuticalInformationSequence = New(0x0054, 0x0016)

	// ResidualSyringeCounts (0054,0017) VR=IS VM=1 Residual Syringe Counts
	ResidualSyringeCounts = New(0x0054, 0x0017)

	// EnergyWindowName (0054,0018) VR=SH VM=1 Energy Window Name
	EnergyWindowName = New(0x0054, 0x0018)

	// DetectorVector (0054,0020) VR=US VM=1-n Detector Vector
	DetectorVector = New(0x0054, 0x0020)

	// NumberOfDetectors (0054,0021) VR=US VM=1 Number of Detectors
	NumberOfDetectors = New(0x0054, 0x0021)

	// DetectorInformationSequence (0054,0022) VR=SQ VM=1 Detector Information Sequence
	DetectorInformationSequence = New(0x0054, 0x0022)

	// PhaseVector (0054,0030) VR=US VM=1-n Phase Vector
	PhaseVector = New(0x0054, 0x0030)

	// NumberOfPhases (0054,0031) VR=US VM=1 Number of Phases
	NumberOfPhases = New(0x0054, 0x0031)

	// PhaseInformationSequence (0054,0032) VR=SQ VM=1 Phase Information Sequence
	PhaseInformationSequence = New(0x0054, 0x0032)

	// NumberOfFramesInPhase (0054,0033) VR=US VM=1 Number of Frames in Phase
	NumberOfFramesInPhase = New(0x0054, 0x0033)

	// PhaseDelay (0054,0036) VR=IS VM=1 Phase Delay
	PhaseDelay = New(0x0054, 0x0036)

	// PauseBetweenFrames (0054,0038) VR=IS VM=1 Pause Between Frames
	PauseBetweenFrames = New(0x0054, 0x0038)

	// PhaseDescription (0054,0039) VR=CS VM=1 Phase Description
	PhaseDescription = New(0x0054, 0x0039)

	// RotationVector (0054,0050) VR=US VM=1-n Rotation Vector
	RotationVector = New(0x0054, 0x0050)

	// NumberOfRotations (0054,0051) VR=US VM=1 Number of Rotations
	NumberOfRotations = New(0x0054, 0x0051)

	// RotationInformationSequence (0054,0052) VR=SQ VM=1 Rotation Information Sequence
	RotationInformationSequence = New(0x0054, 0x0052)

	// NumberOfFramesInRotation (0054,0053) VR=US VM=1 Number of Frames in Rotation
	NumberOfFramesInRotation = New(0x0054, 0x0053)

	// RRIntervalVector (0054,0060) VR=US VM=1-n R-R Interval Vector
	RRIntervalVector = New(0x0054, 0x0060)

	// NumberOfRRIntervals (0054,0061) VR=US VM=1 Number of R-R Intervals
	NumberOfRRIntervals = New(0x0054, 0x0061)

	// GatedInformationSequence (0054,0062) VR=SQ VM=1 Gated Information Sequence
	GatedInformationSequence = New(0x0054, 0x0062)

	// DataInformationSequence (0054,0063) VR=SQ VM=1 Data Information Sequence
	DataInformationSequence = New(0x0054, 0x0063)

	// TimeSlotVector (0054,0070) VR=US VM=1-n Time Slot Vector
	TimeSlotVector = New(0x0054, 0x0070)

	// NumberOfTimeSlots (0054,0071) VR=US VM=1 Number of Time Slots
	NumberOfTimeSlots = New(0x0054, 0x0071)

	// TimeSlotInformationSequence (0054,0072) VR=SQ VM=1 Time Slot Information Sequence
	TimeSlotInformationSequence = New(0x0054, 0x0072)

	// TimeSlotTime (0054,0073) VR=DS VM=1 Time Slot Time
	TimeSlotTime = New(0x0054, 0x0073)

	// SliceVector (0054,0080) VR=US VM=1-n Slice Vector
	SliceVector = New(0x0054, 0x0080)

	// NumberOfSlices (0054,0081) VR=US VM=1 Number of Slices
	NumberOfSlices = New(0x0054, 0x0081)

	// AngularViewVector (0054,0090) VR=US VM=1-n Angular View Vector
	AngularViewVector = New(0x0054, 0x0090)

	// TimeSliceVector (0054,0100) VR=US VM=1-n Time Slice Vector
	TimeSliceVector = New(0x0054, 0x0100)

	// NumberOfTimeSlices (0054,0101) VR=US VM=1 Number of Time Slices
	NumberOfTimeSlices = New(0x0054, 0x0101)

	// StartAngle (0054,0200) VR=DS VM=1 Start Angle
	StartAngle = New(0x0054, 0x0200)

	// TypeOfDetectorMotion (0054,0202) VR=CS VM=1 Type of Detector Motion
	TypeOfDetectorMotion = New(0x0054, 0x0202)

	// TriggerVector (0054,0210) VR=IS VM=1-n Trigger Vector
	TriggerVector = New(0x0054, 0x0210)

	// NumberOfTriggersInPhase (0054,0211) VR=US VM=1 Number of Triggers in Phase
	NumberOfTriggersInPhase = New(0x0054, 0x0211)

	// ViewCodeSequence (0054,0220) VR=SQ VM=1 View Code Sequence
	ViewCodeSequence = New(0x0054, 0x0220)

	// ViewModifierCodeSequence (0054,0222) VR=SQ VM=1 View Modifier Code Sequence
	ViewModifierCodeSequence = New(0x0054, 0x0222)

	// RadionuclideCodeSequence (0054,0300) VR=SQ VM=1 Radionuclide Code Sequence
	RadionuclideCodeSequence = New(0x0054, 0x0300)

	// AdministrationRouteCodeSequence (0054,0302) VR=SQ VM=1 Administration Route Code Sequence
	AdministrationRouteCodeSequence = New(0x0054, 0x0302)

	// RadiopharmaceuticalCodeSequence (0054,0304) VR=SQ VM=1 Radiopharmaceutical Code Sequence
	RadiopharmaceuticalCodeSequence = New(0x0054, 0x0304)

	// CalibrationDataSequence (0054,0306) VR=SQ VM=1 Calibration Data Sequence
	CalibrationDataSequence = New(0x0054, 0x0306)

	// EnergyWindowNumber (0054,0308) VR=US VM=1 Energy Window Number
	EnergyWindowNumber = New(0x0054, 0x0308)

	// ImageID (0054,0400) VR=SH VM=1 Image ID
	ImageID = New(0x0054, 0x0400)

	// PatientOrientationCodeSequence (0054,0410) VR=SQ VM=1 Patient Orientation Code Sequence
	PatientOrientationCodeSequence = New(0x0054, 0x0410)

	// PatientOrientationModifierCodeSequence (0054,0412) VR=SQ VM=1 Patient Orientation Modifier Code Sequence
	PatientOrientationModifierCodeSequence = New(0x0054, 0x0412)

	// PatientGantryRelationshipCodeSequence (0054,0414) VR=SQ VM=1 Patient Gantry Relationship Code Sequence
	PatientGantryRelationshipCodeSequence = New(0x0054, 0x0414)

	// SliceProgressionDirection (0054,0500) VR=CS VM=1 Slice Progression Direction
	SliceProgressionDirection = New(0x0054, 0x0500)

	// ScanProgressionDirection (0054,0501) VR=CS VM=1 Scan Progression Direction
	ScanProgressionDirection = New(0x0054, 0x0501)

	// SeriesType (0054,1000) VR=CS VM=2 Series Type
	SeriesType = New(0x0054, 0x1000)

	// Units (0054,1001) VR=CS VM=1 Units
	Units = New(0x0054, 0x1001)

	// CountsSource (0054,1002) VR=CS VM=1 Counts Source
	CountsSource = New(0x0054, 0x1002)

	// ReprojectionMethod (0054,1004) VR=CS VM=1 Reprojection Method
	ReprojectionMethod = New(0x0054, 0x1004)

	// SUVType (0054,1006) VR=CS VM=1 SUV Type
	SUVType = New(0x0054, 0x1006)

	// RandomsCorrectionMethod (0054,1100) VR=CS VM=1 Randoms Correction Method
	RandomsCorrectionMethod = New(0x0054, 0x1100)

	// AttenuationCorrectionMethod (0054,1101) VR=LO VM=1 Attenuation Correction Method
	AttenuationCorrectionMethod = New(0x0054, 0x1101)

	// DecayCorrection (0054,1102) VR=CS VM=1 Decay Correction
	DecayCorrection = New(0x0054, 0x1102)

	// ReconstructionMethod (0054,1103) VR=LO VM=1 Reconstruction Method
	ReconstructionMethod = New(0x0054, 0x1103)

	// DetectorLinesOfResponseUsed (0054,1104) VR=LO VM=1 Detector Lines of Response Used
	DetectorLinesOfResponseUsed = New(0x0054, 0x1104)

	// ScatterCorrectionMethod (0054,1105) VR=LO VM=1 Scatter Correction Method
	ScatterCorrectionMethod = New(0x0054, 0x1105)

	// AxialAcceptance (0054,1200) VR=DS VM=1 Axial Acceptance
	AxialAcceptance = New(0x0054, 0x1200)

	// AxialMash (0054,1201) VR=IS VM=2 Axial Mash
	AxialMash = New(0x0054, 0x1201)

	// TransverseMash (0054,1202) VR=IS VM=1 Transverse Mash
	TransverseMash = New(0x0054, 0x1202)

	// DetectorElementSize (0054,1203) VR=DS VM=2 Detector Element Size
	DetectorElementSize = New(0x0054, 0x1203)

	// CoincidenceWindowWidth (0054,1210) VR=DS VM=1 Coincidence Window Width
	CoincidenceWindowWidth = New(0x0054, 0x1210)

	// SecondaryCountsType (0054,1220) VR=CS VM=1-n Secondary Counts Type
	SecondaryCountsType = New(0x0054, 0x1220)

	// FrameReferenceTime (0054,1300) VR=DS VM=1 Frame Reference Time
	FrameReferenceTime = New(0x0054, 0x1300)

	// PrimaryPromptsCountsAccumulated (0054,1310) VR=IS VM=1 Primary (Prompts) Counts Accumulated
	PrimaryPromptsCountsAccumulated = New(0x0054, 0x1310)

	// SecondaryCountsAccumulated (0054,1311) VR=IS VM=1-n Secondary Counts Accumulated
	SecondaryCountsAccumulated = New(0x0054, 0x1311)

	// SliceSensitivityFactor (0054,1320) VR=DS VM=1 Slice Sensitivity Factor
	SliceSensitivityFactor = New(0x0054, 0x1320)

	// DecayFactor (0054,1321) VR=DS VM=1 Decay Factor
	DecayFactor = New(0x0054, 0x1321)

	// DoseCalibrationFactor (0054,1322) VR=DS VM=1 Dose Calibration Factor
	DoseCalibrationFactor = New(0x0054, 0x1322)

	// ScatterFractionFactor (0054,1323) VR=DS VM=1 Scatter Fraction Factor
	ScatterFractionFactor = New(0x0054, 0x1323)

	// DeadTimeFactor (0054,1324) VR=DS VM=1 Dead Time Factor
	DeadTimeFactor = New(0x0054, 0x1324)

	// ImageIndex (0054,1330) VR=US VM=1 Image Index
	ImageIndex = New(0x0054, 0x1330)

	// CountsIncludedRETIRED (0054,1400) VR=CS VM=1-n Counts Included (RETIRED)
	CountsIncludedRETIRED = New(0x0054, 0x1400)

	// DeadTimeCorrectionFlagRETIRED (0054,1401) VR=CS VM=1 Dead Time Correction Flag (RETIRED)
	DeadTimeCorrectionFlagRETIRED = New(0x0054, 0x1401)

	// HistogramSequence (0060,3000) VR=SQ VM=1 Histogram Sequence
	HistogramSequence = New(0x0060, 0x3000)

	// HistogramNumberOfBins (0060,3002) VR=US VM=1 Histogram Number of Bins
	HistogramNumberOfBins = New(0x0060, 0x3002)

	HistogramFirstBinValue = New(0x0060, 0x3004)

	HistogramLastBinValue = New(0x0060, 0x3006)

	// HistogramBinWidth (0060,3008) VR=US VM=1 Histogram Bin Width
	HistogramBinWidth = New(0x0060, 0x3008)

	// HistogramExplanation (0060,3010) VR=LO VM=1 Histogram Explanation
	HistogramExplanation = New(0x0060, 0x3010)

	// HistogramData (0060,3020) VR=UL VM=1-n Histogram Data
	HistogramData = New(0x0060, 0x3020)

	// SegmentationType (0062,0001) VR=CS VM=1 Segmentation Type
	SegmentationType = New(0x0062, 0x0001)

	// SegmentSequence (0062,0002) VR=SQ VM=1 Segment Sequence
	SegmentSequence = New(0x0062, 0x0002)

	// SegmentedPropertyCategoryCodeSequence (0062,0003) VR=SQ VM=1 Segmented Property Category Code Sequence
	SegmentedPropertyCategoryCodeSequence = New(0x0062, 0x0003)

	// SegmentNumber (0062,0004) VR=US VM=1 Segment Number
	SegmentNumber = New(0x0062, 0x0004)

	// SegmentLabel (0062,0005) VR=LO VM=1 Segment Label
	SegmentLabel = New(0x0062, 0x0005)

	// SegmentDescription (0062,0006) VR=ST VM=1 Segment Description
	SegmentDescription = New(0x0062, 0x0006)

	// SegmentationAlgorithmIdentificationSequence (0062,0007) VR=SQ VM=1 Segmentation Algorithm Identification Sequence
	SegmentationAlgorithmIdentificationSequence = New(0x0062, 0x0007)

	// SegmentAlgorithmType (0062,0008) VR=CS VM=1 Segment Algorithm Type
	SegmentAlgorithmType = New(0x0062, 0x0008)

	// SegmentAlgorithmName (0062,0009) VR=LO VM=1-n Segment Algorithm Name
	SegmentAlgorithmName = New(0x0062, 0x0009)

	// SegmentIdentificationSequence (0062,000A) VR=SQ VM=1 Segment Identification Sequence
	SegmentIdentificationSequence = New(0x0062, 0x000A)

	// ReferencedSegmentNumber (0062,000B) VR=US VM=1-n Referenced Segment Number
	ReferencedSegmentNumber = New(0x0062, 0x000B)

	// RecommendedDisplayGrayscaleValue (0062,000C) VR=US VM=1 Recommended Display Grayscale Value
	RecommendedDisplayGrayscaleValue = New(0x0062, 0x000C)

	// RecommendedDisplayCIELabValue (0062,000D) VR=US VM=3 Recommended Display CIELab Value
	RecommendedDisplayCIELabValue = New(0x0062, 0x000D)

	// MaximumFractionalValue (0062,000E) VR=US VM=1 Maximum Fractional Value
	MaximumFractionalValue = New(0x0062, 0x000E)

	// SegmentedPropertyTypeCodeSequence (0062,000F) VR=SQ VM=1 Segmented Property Type Code Sequence
	SegmentedPropertyTypeCodeSequence = New(0x0062, 0x000F)

	// SegmentationFractionalType (0062,0010) VR=CS VM=1 Segmentation Fractional Type
	SegmentationFractionalType = New(0x0062, 0x0010)

	// SegmentedPropertyTypeModifierCodeSequence (0062,0011) VR=SQ VM=1 Segmented Property Type Modifier Code Sequence
	SegmentedPropertyTypeModifierCodeSequence = New(0x0062, 0x0011)

	// UsedSegmentsSequence (0062,0012) VR=SQ VM=1 Used Segments Sequence
	UsedSegmentsSequence = New(0x0062, 0x0012)

	// SegmentsOverlap (0062,0013) VR=CS VM=1 Segments Overlap
	SegmentsOverlap = New(0x0062, 0x0013)

	// TrackingID (0062,0020) VR=UT VM=1 Tracking ID
	TrackingID = New(0x0062, 0x0020)

	// TrackingUID (0062,0021) VR=UI VM=1 Tracking UID
	TrackingUID = New(0x0062, 0x0021)

	// DeformableRegistrationSequence (0064,0002) VR=SQ VM=1 Deformable Registration Sequence
	DeformableRegistrationSequence = New(0x0064, 0x0002)

	// SourceFrameOfReferenceUID (0064,0003) VR=UI VM=1 Source Frame of Reference UID
	SourceFrameOfReferenceUID = New(0x0064, 0x0003)

	// DeformableRegistrationGridSequence (0064,0005) VR=SQ VM=1 Deformable Registration Grid Sequence
	DeformableRegistrationGridSequence = New(0x0064, 0x0005)

	// GridDimensions (0064,0007) VR=UL VM=3 Grid Dimensions
	GridDimensions = New(0x0064, 0x0007)

	// GridResolution (0064,0008) VR=FD VM=3 Grid Resolution
	GridResolution = New(0x0064, 0x0008)

	// VectorGridData (0064,0009) VR=OF VM=1 Vector Grid Data
	VectorGridData = New(0x0064, 0x0009)

	// PreDeformationMatrixRegistrationSequence (0064,000F) VR=SQ VM=1 Pre Deformation Matrix Registration Sequence
	PreDeformationMatrixRegistrationSequence = New(0x0064, 0x000F)

	// PostDeformationMatrixRegistrationSequence (0064,0010) VR=SQ VM=1 Post Deformation Matrix Registration Sequence
	PostDeformationMatrixRegistrationSequence = New(0x0064, 0x0010)

	// NumberOfSurfaces (0066,0001) VR=UL VM=1 Number of Surfaces
	NumberOfSurfaces = New(0x0066, 0x0001)

	// SurfaceSequence (0066,0002) VR=SQ VM=1 Surface Sequence
	SurfaceSequence = New(0x0066, 0x0002)

	// SurfaceNumber (0066,0003) VR=UL VM=1 Surface Number
	SurfaceNumber = New(0x0066, 0x0003)

	// SurfaceComments (0066,0004) VR=LT VM=1 Surface Comments
	SurfaceComments = New(0x0066, 0x0004)

	// SurfaceOffset (0066,0005) VR=FL VM=1 Surface Offset
	SurfaceOffset = New(0x0066, 0x0005)

	// SurfaceProcessing (0066,0009) VR=CS VM=1 Surface Processing
	SurfaceProcessing = New(0x0066, 0x0009)

	// SurfaceProcessingRatio (0066,000A) VR=FL VM=1 Surface Processing Ratio
	SurfaceProcessingRatio = New(0x0066, 0x000A)

	// SurfaceProcessingDescription (0066,000B) VR=LO VM=1 Surface Processing Description
	SurfaceProcessingDescription = New(0x0066, 0x000B)

	// RecommendedPresentationOpacity (0066,000C) VR=FL VM=1 Recommended Presentation Opacity
	RecommendedPresentationOpacity = New(0x0066, 0x000C)

	// RecommendedPresentationType (0066,000D) VR=CS VM=1 Recommended Presentation Type
	RecommendedPresentationType = New(0x0066, 0x000D)

	// FiniteVolume (0066,000E) VR=CS VM=1 Finite Volume
	FiniteVolume = New(0x0066, 0x000E)

	// Manifold (0066,0010) VR=CS VM=1 Manifold
	Manifold = New(0x0066, 0x0010)

	// SurfacePointsSequence (0066,0011) VR=SQ VM=1 Surface Points Sequence
	SurfacePointsSequence = New(0x0066, 0x0011)

	// SurfacePointsNormalsSequence (0066,0012) VR=SQ VM=1 Surface Points Normals Sequence
	SurfacePointsNormalsSequence = New(0x0066, 0x0012)

	// SurfaceMeshPrimitivesSequence (0066,0013) VR=SQ VM=1 Surface Mesh Primitives Sequence
	SurfaceMeshPrimitivesSequence = New(0x0066, 0x0013)

	// NumberOfSurfacePoints (0066,0015) VR=UL VM=1 Number of Surface Points
	NumberOfSurfacePoints = New(0x0066, 0x0015)

	// PointCoordinatesData (0066,0016) VR=OF VM=1 Point Coordinates Data
	PointCoordinatesData = New(0x0066, 0x0016)

	// PointPositionAccuracy (0066,0017) VR=FL VM=3 Point Position Accuracy
	PointPositionAccuracy = New(0x0066, 0x0017)

	// MeanPointDistance (0066,0018) VR=FL VM=1 Mean Point Distance
	MeanPointDistance = New(0x0066, 0x0018)

	// MaximumPointDistance (0066,0019) VR=FL VM=1 Maximum Point Distance
	MaximumPointDistance = New(0x0066, 0x0019)

	// PointsBoundingBoxCoordinates (0066,001A) VR=FL VM=6 Points Bounding Box Coordinates
	PointsBoundingBoxCoordinates = New(0x0066, 0x001A)

	// AxisOfRotation (0066,001B) VR=FL VM=3 Axis of Rotation
	AxisOfRotation = New(0x0066, 0x001B)

	// CenterOfRotation (0066,001C) VR=FL VM=3 Center of Rotation
	CenterOfRotation = New(0x0066, 0x001C)

	// NumberOfVectors (0066,001E) VR=UL VM=1 Number of Vectors
	NumberOfVectors = New(0x0066, 0x001E)

	// VectorDimensionality (0066,001F) VR=US VM=1 Vector Dimensionality
	VectorDimensionality = New(0x0066, 0x001F)

	// VectorAccuracy (0066,0020) VR=FL VM=1-n Vector Accuracy
	VectorAccuracy = New(0x0066, 0x0020)

	// VectorCoordinateData (0066,0021) VR=OF VM=1 Vector Coordinate Data
	VectorCoordinateData = New(0x0066, 0x0021)

	// DoublePointCoordinatesData (0066,0022) VR=OD VM=1 Double Point Coordinates Data
	DoublePointCoordinatesData = New(0x0066, 0x0022)

	// TrianglePointIndexListRETIRED (0066,0023) VR=OW VM=1 Triangle Point Index List (RETIRED)
	TrianglePointIndexListRETIRED = New(0x0066, 0x0023)

	// EdgePointIndexListRETIRED (0066,0024) VR=OW VM=1 Edge Point Index List (RETIRED)
	EdgePointIndexListRETIRED = New(0x0066, 0x0024)

	// VertexPointIndexListRETIRED (0066,0025) VR=OW VM=1 Vertex Point Index List (RETIRED)
	VertexPointIndexListRETIRED = New(0x0066, 0x0025)

	// TriangleStripSequence (0066,0026) VR=SQ VM=1 Triangle Strip Sequence
	TriangleStripSequence = New(0x0066, 0x0026)

	// TriangleFanSequence (0066,0027) VR=SQ VM=1 Triangle Fan Sequence
	TriangleFanSequence = New(0x0066, 0x0027)

	// LineSequence (0066,0028) VR=SQ VM=1 Line Sequence
	LineSequence = New(0x0066, 0x0028)

	// PrimitivePointIndexListRETIRED (0066,0029) VR=OW VM=1 Primitive Point Index List (RETIRED)
	PrimitivePointIndexListRETIRED = New(0x0066, 0x0029)

	// SurfaceCount (0066,002A) VR=UL VM=1 Surface Count
	SurfaceCount = New(0x0066, 0x002A)

	// ReferencedSurfaceSequence (0066,002B) VR=SQ VM=1 Referenced Surface Sequence
	ReferencedSurfaceSequence = New(0x0066, 0x002B)

	// ReferencedSurfaceNumber (0066,002C) VR=UL VM=1 Referenced Surface Number
	ReferencedSurfaceNumber = New(0x0066, 0x002C)

	// SegmentSurfaceGenerationAlgorithmIdentificationSequence (0066,002D) VR=SQ VM=1 Segment Surface Generation Algorithm Identification Sequence
	SegmentSurfaceGenerationAlgorithmIdentificationSequence = New(0x0066, 0x002D)

	// SegmentSurfaceSourceInstanceSequence (0066,002E) VR=SQ VM=1 Segment Surface Source Instance Sequence
	SegmentSurfaceSourceInstanceSequence = New(0x0066, 0x002E)

	// AlgorithmFamilyCodeSequence (0066,002F) VR=SQ VM=1 Algorithm Family Code Sequence
	AlgorithmFamilyCodeSequence = New(0x0066, 0x002F)

	// AlgorithmNameCodeSequence (0066,0030) VR=SQ VM=1 Algorithm Name Code Sequence
	AlgorithmNameCodeSequence = New(0x0066, 0x0030)

	// AlgorithmVersion (0066,0031) VR=LO VM=1 Algorithm Version
	AlgorithmVersion = New(0x0066, 0x0031)

	// AlgorithmParameters (0066,0032) VR=LT VM=1 Algorithm Parameters
	AlgorithmParameters = New(0x0066, 0x0032)

	// FacetSequence (0066,0034) VR=SQ VM=1 Facet Sequence
	FacetSequence = New(0x0066, 0x0034)

	// SurfaceProcessingAlgorithmIdentificationSequence (0066,0035) VR=SQ VM=1 Surface Processing Algorithm Identification Sequence
	SurfaceProcessingAlgorithmIdentificationSequence = New(0x0066, 0x0035)

	// AlgorithmName (0066,0036) VR=LO VM=1 Algorithm Name
	AlgorithmName = New(0x0066, 0x0036)

	// RecommendedPointRadius (0066,0037) VR=FL VM=1 Recommended Point Radius
	RecommendedPointRadius = New(0x0066, 0x0037)

	// RecommendedLineThickness (0066,0038) VR=FL VM=1 Recommended Line Thickness
	RecommendedLineThickness = New(0x0066, 0x0038)

	// LongPrimitivePointIndexList (0066,0040) VR=OL VM=1 Long Primitive Point Index List
	LongPrimitivePointIndexList = New(0x0066, 0x0040)

	// LongTrianglePointIndexList (0066,0041) VR=OL VM=1 Long Triangle Point Index List
	LongTrianglePointIndexList = New(0x0066, 0x0041)

	// LongEdgePointIndexList (0066,0042) VR=OL VM=1 Long Edge Point Index List
	LongEdgePointIndexList = New(0x0066, 0x0042)

	// LongVertexPointIndexList (0066,0043) VR=OL VM=1 Long Vertex Point Index List
	LongVertexPointIndexList = New(0x0066, 0x0043)

	// TrackSetSequence (0066,0101) VR=SQ VM=1 Track Set Sequence
	TrackSetSequence = New(0x0066, 0x0101)

	// TrackSequence (0066,0102) VR=SQ VM=1 Track Sequence
	TrackSequence = New(0x0066, 0x0102)

	// RecommendedDisplayCIELabValueList (0066,0103) VR=OW VM=1 Recommended Display CIELab Value List
	RecommendedDisplayCIELabValueList = New(0x0066, 0x0103)

	// TrackingAlgorithmIdentificationSequence (0066,0104) VR=SQ VM=1 Tracking Algorithm Identification Sequence
	TrackingAlgorithmIdentificationSequence = New(0x0066, 0x0104)

	// TrackSetNumber (0066,0105) VR=UL VM=1 Track Set Number
	TrackSetNumber = New(0x0066, 0x0105)

	// TrackSetLabel (0066,0106) VR=LO VM=1 Track Set Label
	TrackSetLabel = New(0x0066, 0x0106)

	// TrackSetDescription (0066,0107) VR=UT VM=1 Track Set Description
	TrackSetDescription = New(0x0066, 0x0107)

	// TrackSetAnatomicalTypeCodeSequence (0066,0108) VR=SQ VM=1 Track Set Anatomical Type Code Sequence
	TrackSetAnatomicalTypeCodeSequence = New(0x0066, 0x0108)

	// MeasurementsSequence (0066,0121) VR=SQ VM=1 Measurements Sequence
	MeasurementsSequence = New(0x0066, 0x0121)

	// TrackSetStatisticsSequence (0066,0124) VR=SQ VM=1 Track Set Statistics Sequence
	TrackSetStatisticsSequence = New(0x0066, 0x0124)

	// FloatingPointValues (0066,0125) VR=OF VM=1 Floating Point Values
	FloatingPointValues = New(0x0066, 0x0125)

	// TrackPointIndexList (0066,0129) VR=OL VM=1 Track Point Index List
	TrackPointIndexList = New(0x0066, 0x0129)

	// TrackStatisticsSequence (0066,0130) VR=SQ VM=1 Track Statistics Sequence
	TrackStatisticsSequence = New(0x0066, 0x0130)

	// MeasurementValuesSequence (0066,0132) VR=SQ VM=1 Measurement Values Sequence
	MeasurementValuesSequence = New(0x0066, 0x0132)

	// DiffusionAcquisitionCodeSequence (0066,0133) VR=SQ VM=1 Diffusion Acquisition Code Sequence
	DiffusionAcquisitionCodeSequence = New(0x0066, 0x0133)

	// DiffusionModelCodeSequence (0066,0134) VR=SQ VM=1 Diffusion Model Code Sequence
	DiffusionModelCodeSequence = New(0x0066, 0x0134)

	// ImplantSize (0068,6210) VR=LO VM=1 Implant Size
	ImplantSize = New(0x0068, 0x6210)

	// ImplantTemplateVersion (0068,6221) VR=LO VM=1 Implant Template Version
	ImplantTemplateVersion = New(0x0068, 0x6221)

	// ReplacedImplantTemplateSequence (0068,6222) VR=SQ VM=1 Replaced Implant Template Sequence
	ReplacedImplantTemplateSequence = New(0x0068, 0x6222)

	// ImplantType (0068,6223) VR=CS VM=1 Implant Type
	ImplantType = New(0x0068, 0x6223)

	// DerivationImplantTemplateSequence (0068,6224) VR=SQ VM=1 Derivation Implant Template Sequence
	DerivationImplantTemplateSequence = New(0x0068, 0x6224)

	// OriginalImplantTemplateSequence (0068,6225) VR=SQ VM=1 Original Implant Template Sequence
	OriginalImplantTemplateSequence = New(0x0068, 0x6225)

	// EffectiveDateTime (0068,6226) VR=DT VM=1 Effective DateTime
	EffectiveDateTime = New(0x0068, 0x6226)

	// ImplantTargetAnatomySequence (0068,6230) VR=SQ VM=1 Implant Target Anatomy Sequence
	ImplantTargetAnatomySequence = New(0x0068, 0x6230)

	// InformationFromManufacturerSequence (0068,6260) VR=SQ VM=1 Information From Manufacturer Sequence
	InformationFromManufacturerSequence = New(0x0068, 0x6260)

	// NotificationFromManufacturerSequence (0068,6265) VR=SQ VM=1 Notification From Manufacturer Sequence
	NotificationFromManufacturerSequence = New(0x0068, 0x6265)

	// InformationIssueDateTime (0068,6270) VR=DT VM=1 Information Issue DateTime
	InformationIssueDateTime = New(0x0068, 0x6270)

	// InformationSummary (0068,6280) VR=ST VM=1 Information Summary
	InformationSummary = New(0x0068, 0x6280)

	// ImplantRegulatoryDisapprovalCodeSequence (0068,62A0) VR=SQ VM=1 Implant Regulatory Disapproval Code Sequence
	ImplantRegulatoryDisapprovalCodeSequence = New(0x0068, 0x62A0)

	// OverallTemplateSpatialTolerance (0068,62A5) VR=FD VM=1 Overall Template Spatial Tolerance
	OverallTemplateSpatialTolerance = New(0x0068, 0x62A5)

	// HPGLDocumentSequence (0068,62C0) VR=SQ VM=1 HPGL Document Sequence
	HPGLDocumentSequence = New(0x0068, 0x62C0)

	// HPGLDocumentID (0068,62D0) VR=US VM=1 HPGL Document ID
	HPGLDocumentID = New(0x0068, 0x62D0)

	// HPGLDocumentLabel (0068,62D5) VR=LO VM=1 HPGL Document Label
	HPGLDocumentLabel = New(0x0068, 0x62D5)

	// ViewOrientationCodeSequence (0068,62E0) VR=SQ VM=1 View Orientation Code Sequence
	ViewOrientationCodeSequence = New(0x0068, 0x62E0)

	// ViewOrientationModifierCodeSequence (0068,62F0) VR=SQ VM=1 View Orientation Modifier Code Sequence
	ViewOrientationModifierCodeSequence = New(0x0068, 0x62F0)

	// HPGLDocumentScaling (0068,62F2) VR=FD VM=1 HPGL Document Scaling
	HPGLDocumentScaling = New(0x0068, 0x62F2)

	// HPGLDocument (0068,6300) VR=OB VM=1 HPGL Document
	HPGLDocument = New(0x0068, 0x6300)

	// HPGLContourPenNumber (0068,6310) VR=US VM=1 HPGL Contour Pen Number
	HPGLContourPenNumber = New(0x0068, 0x6310)

	// HPGLPenSequence (0068,6320) VR=SQ VM=1 HPGL Pen Sequence
	HPGLPenSequence = New(0x0068, 0x6320)

	// HPGLPenNumber (0068,6330) VR=US VM=1 HPGL Pen Number
	HPGLPenNumber = New(0x0068, 0x6330)

	// HPGLPenLabel (0068,6340) VR=LO VM=1 HPGL Pen Label
	HPGLPenLabel = New(0x0068, 0x6340)

	// HPGLPenDescription (0068,6345) VR=ST VM=1 HPGL Pen Description
	HPGLPenDescription = New(0x0068, 0x6345)

	// RecommendedRotationPoint (0068,6346) VR=FD VM=2 Recommended Rotation Point
	RecommendedRotationPoint = New(0x0068, 0x6346)

	// BoundingRectangle (0068,6347) VR=FD VM=4 Bounding Rectangle
	BoundingRectangle = New(0x0068, 0x6347)

	// ImplantTemplate3DModelSurfaceNumber (0068,6350) VR=US VM=1-n Implant Template 3D Model Surface Number
	ImplantTemplate3DModelSurfaceNumber = New(0x0068, 0x6350)

	// SurfaceModelDescriptionSequence (0068,6360) VR=SQ VM=1 Surface Model Description Sequence
	SurfaceModelDescriptionSequence = New(0x0068, 0x6360)

	// SurfaceModelLabel (0068,6380) VR=LO VM=1 Surface Model Label
	SurfaceModelLabel = New(0x0068, 0x6380)

	// SurfaceModelScalingFactor (0068,6390) VR=FD VM=1 Surface Model Scaling Factor
	SurfaceModelScalingFactor = New(0x0068, 0x6390)

	// MaterialsCodeSequence (0068,63A0) VR=SQ VM=1 Materials Code Sequence
	MaterialsCodeSequence = New(0x0068, 0x63A0)

	// CoatingMaterialsCodeSequence (0068,63A4) VR=SQ VM=1 Coating Materials Code Sequence
	CoatingMaterialsCodeSequence = New(0x0068, 0x63A4)

	// ImplantTypeCodeSequence (0068,63A8) VR=SQ VM=1 Implant Type Code Sequence
	ImplantTypeCodeSequence = New(0x0068, 0x63A8)

	// FixationMethodCodeSequence (0068,63AC) VR=SQ VM=1 Fixation Method Code Sequence
	FixationMethodCodeSequence = New(0x0068, 0x63AC)

	// MatingFeatureSetsSequence (0068,63B0) VR=SQ VM=1 Mating Feature Sets Sequence
	MatingFeatureSetsSequence = New(0x0068, 0x63B0)

	// MatingFeatureSetID (0068,63C0) VR=US VM=1 Mating Feature Set ID
	MatingFeatureSetID = New(0x0068, 0x63C0)

	// MatingFeatureSetLabel (0068,63D0) VR=LO VM=1 Mating Feature Set Label
	MatingFeatureSetLabel = New(0x0068, 0x63D0)

	// MatingFeatureSequence (0068,63E0) VR=SQ VM=1 Mating Feature Sequence
	MatingFeatureSequence = New(0x0068, 0x63E0)

	// MatingFeatureID (0068,63F0) VR=US VM=1 Mating Feature ID
	MatingFeatureID = New(0x0068, 0x63F0)

	// MatingFeatureDegreeOfFreedomSequence (0068,6400) VR=SQ VM=1 Mating Feature Degree of Freedom Sequence
	MatingFeatureDegreeOfFreedomSequence = New(0x0068, 0x6400)

	// DegreeOfFreedomID (0068,6410) VR=US VM=1 Degree of Freedom ID
	DegreeOfFreedomID = New(0x0068, 0x6410)

	// DegreeOfFreedomType (0068,6420) VR=CS VM=1 Degree of Freedom Type
	DegreeOfFreedomType = New(0x0068, 0x6420)

	// TwoDMatingFeatureCoordinatesSequence (0068,6430) VR=SQ VM=1 2D Mating Feature Coordinates Sequence
	TwoDMatingFeatureCoordinatesSequence = New(0x0068, 0x6430)

	// ReferencedHPGLDocumentID (0068,6440) VR=US VM=1 Referenced HPGL Document ID
	ReferencedHPGLDocumentID = New(0x0068, 0x6440)

	// TwoDMatingPoint (0068,6450) VR=FD VM=2 2D Mating Point
	TwoDMatingPoint = New(0x0068, 0x6450)

	// TwoDMatingAxes (0068,6460) VR=FD VM=4 2D Mating Axes
	TwoDMatingAxes = New(0x0068, 0x6460)

	// TwoDDegreeOfFreedomSequence (0068,6470) VR=SQ VM=1 2D Degree of Freedom Sequence
	TwoDDegreeOfFreedomSequence = New(0x0068, 0x6470)

	// ThreeDDegreeOfFreedomAxis (0068,6490) VR=FD VM=3 3D Degree of Freedom Axis
	ThreeDDegreeOfFreedomAxis = New(0x0068, 0x6490)

	// RangeOfFreedom (0068,64A0) VR=FD VM=2 Range of Freedom
	RangeOfFreedom = New(0x0068, 0x64A0)

	// ThreeDMatingPoint (0068,64C0) VR=FD VM=3 3D Mating Point
	ThreeDMatingPoint = New(0x0068, 0x64C0)

	// ThreeDMatingAxes (0068,64D0) VR=FD VM=9 3D Mating Axes
	ThreeDMatingAxes = New(0x0068, 0x64D0)

	// TwoDDegreeOfFreedomAxis (0068,64F0) VR=FD VM=3 2D Degree of Freedom Axis
	TwoDDegreeOfFreedomAxis = New(0x0068, 0x64F0)

	// PlanningLandmarkPointSequence (0068,6500) VR=SQ VM=1 Planning Landmark Point Sequence
	PlanningLandmarkPointSequence = New(0x0068, 0x6500)

	// PlanningLandmarkLineSequence (0068,6510) VR=SQ VM=1 Planning Landmark Line Sequence
	PlanningLandmarkLineSequence = New(0x0068, 0x6510)

	// PlanningLandmarkPlaneSequence (0068,6520) VR=SQ VM=1 Planning Landmark Plane Sequence
	PlanningLandmarkPlaneSequence = New(0x0068, 0x6520)

	// PlanningLandmarkID (0068,6530) VR=US VM=1 Planning Landmark ID
	PlanningLandmarkID = New(0x0068, 0x6530)

	// PlanningLandmarkDescription (0068,6540) VR=LO VM=1 Planning Landmark Description
	PlanningLandmarkDescription = New(0x0068, 0x6540)

	// PlanningLandmarkIdentificationCodeSequence (0068,6545) VR=SQ VM=1 Planning Landmark Identification Code Sequence
	PlanningLandmarkIdentificationCodeSequence = New(0x0068, 0x6545)

	// TwoDPointCoordinatesSequence (0068,6550) VR=SQ VM=1 2D Point Coordinates Sequence
	TwoDPointCoordinatesSequence = New(0x0068, 0x6550)

	// TwoDPointCoordinates (0068,6560) VR=FD VM=2 2D Point Coordinates
	TwoDPointCoordinates = New(0x0068, 0x6560)

	// ThreeDPointCoordinates (0068,6590) VR=FD VM=3 3D Point Coordinates
	ThreeDPointCoordinates = New(0x0068, 0x6590)

	// TwoDLineCoordinatesSequence (0068,65A0) VR=SQ VM=1 2D Line Coordinates Sequence
	TwoDLineCoordinatesSequence = New(0x0068, 0x65A0)

	// TwoDLineCoordinates (0068,65B0) VR=FD VM=4 2D Line Coordinates
	TwoDLineCoordinates = New(0x0068, 0x65B0)

	// ThreeDLineCoordinates (0068,65D0) VR=FD VM=6 3D Line Coordinates
	ThreeDLineCoordinates = New(0x0068, 0x65D0)

	// TwoDPlaneCoordinatesSequence (0068,65E0) VR=SQ VM=1 2D Plane Coordinates Sequence
	TwoDPlaneCoordinatesSequence = New(0x0068, 0x65E0)

	// TwoDPlaneIntersection (0068,65F0) VR=FD VM=4 2D Plane Intersection
	TwoDPlaneIntersection = New(0x0068, 0x65F0)

	// ThreeDPlaneOrigin (0068,6610) VR=FD VM=3 3D Plane Origin
	ThreeDPlaneOrigin = New(0x0068, 0x6610)

	// ThreeDPlaneNormal (0068,6620) VR=FD VM=3 3D Plane Normal
	ThreeDPlaneNormal = New(0x0068, 0x6620)

	// ModelModification (0068,7001) VR=CS VM=1 Model Modification
	ModelModification = New(0x0068, 0x7001)

	// ModelMirroring (0068,7002) VR=CS VM=1 Model Mirroring
	ModelMirroring = New(0x0068, 0x7002)

	// ModelUsageCodeSequence (0068,7003) VR=SQ VM=1 Model Usage Code Sequence
	ModelUsageCodeSequence = New(0x0068, 0x7003)

	// ModelGroupUID (0068,7004) VR=UI VM=1 Model Group UID
	ModelGroupUID = New(0x0068, 0x7004)

	// RelativeURIReferenceWithinEncapsulatedDocument (0068,7005) VR=UR VM=1 Relative URI Reference Within Encapsulated Document
	RelativeURIReferenceWithinEncapsulatedDocument = New(0x0068, 0x7005)

	// AnnotationCoordinateType (006A,0001) VR=CS VM=1 Annotation Coordinate Type
	AnnotationCoordinateType = New(0x006A, 0x0001)

	// AnnotationGroupSequence (006A,0002) VR=SQ VM=1 Annotation Group Sequence
	AnnotationGroupSequence = New(0x006A, 0x0002)

	// AnnotationGroupUID (006A,0003) VR=UI VM=1 Annotation Group UID
	AnnotationGroupUID = New(0x006A, 0x0003)

	// AnnotationGroupLabel (006A,0005) VR=LO VM=1 Annotation Group Label
	AnnotationGroupLabel = New(0x006A, 0x0005)

	// AnnotationGroupDescription (006A,0006) VR=UT VM=1 Annotation Group Description
	AnnotationGroupDescription = New(0x006A, 0x0006)

	// AnnotationGroupGenerationType (006A,0007) VR=CS VM=1 Annotation Group Generation Type
	AnnotationGroupGenerationType = New(0x006A, 0x0007)

	// AnnotationGroupAlgorithmIdentificationSequence (006A,0008) VR=SQ VM=1 Annotation Group Algorithm Identification Sequence
	AnnotationGroupAlgorithmIdentificationSequence = New(0x006A, 0x0008)

	// AnnotationPropertyCategoryCodeSequence (006A,0009) VR=SQ VM=1 Annotation Property Category Code Sequence
	AnnotationPropertyCategoryCodeSequence = New(0x006A, 0x0009)

	// AnnotationPropertyTypeCodeSequence (006A,000A) VR=SQ VM=1 Annotation Property Type Code Sequence
	AnnotationPropertyTypeCodeSequence = New(0x006A, 0x000A)

	// AnnotationPropertyTypeModifierCodeSequence (006A,000B) VR=SQ VM=1 Annotation Property Type Modifier Code Sequence
	AnnotationPropertyTypeModifierCodeSequence = New(0x006A, 0x000B)

	// NumberOfAnnotations (006A,000C) VR=UL VM=1 Number of Annotations
	NumberOfAnnotations = New(0x006A, 0x000C)

	// AnnotationAppliesToAllOpticalPaths (006A,000D) VR=CS VM=1 Annotation Applies to All Optical Paths
	AnnotationAppliesToAllOpticalPaths = New(0x006A, 0x000D)

	// ReferencedOpticalPathIdentifier (006A,000E) VR=SH VM=1-n Referenced Optical Path Identifier
	ReferencedOpticalPathIdentifier = New(0x006A, 0x000E)

	// AnnotationAppliesToAllZPlanes (006A,000F) VR=CS VM=1 Annotation Applies to All Z Planes
	AnnotationAppliesToAllZPlanes = New(0x006A, 0x000F)

	// CommonZCoordinateValue (006A,0010) VR=FD VM=1-n Common Z Coordinate Value
	CommonZCoordinateValue = New(0x006A, 0x0010)

	// AnnotationIndexList (006A,0011) VR=OL VM=1 Annotation Index List
	AnnotationIndexList = New(0x006A, 0x0011)

	// GraphicAnnotationSequence (0070,0001) VR=SQ VM=1 Graphic Annotation Sequence
	GraphicAnnotationSequence = New(0x0070, 0x0001)

	// GraphicLayer (0070,0002) VR=CS VM=1 Graphic Layer
	GraphicLayer = New(0x0070, 0x0002)

	// BoundingBoxAnnotationUnits (0070,0003) VR=CS VM=1 Bounding Box Annotation Units
	BoundingBoxAnnotationUnits = New(0x0070, 0x0003)

	// AnchorPointAnnotationUnits (0070,0004) VR=CS VM=1 Anchor Point Annotation Units
	AnchorPointAnnotationUnits = New(0x0070, 0x0004)

	// GraphicAnnotationUnits (0070,0005) VR=CS VM=1 Graphic Annotation Units
	GraphicAnnotationUnits = New(0x0070, 0x0005)

	// UnformattedTextValue (0070,0006) VR=ST VM=1 Unformatted Text Value
	UnformattedTextValue = New(0x0070, 0x0006)

	// TextObjectSequence (0070,0008) VR=SQ VM=1 Text Object Sequence
	TextObjectSequence = New(0x0070, 0x0008)

	// GraphicObjectSequence (0070,0009) VR=SQ VM=1 Graphic Object Sequence
	GraphicObjectSequence = New(0x0070, 0x0009)

	// BoundingBoxTopLeftHandCorner (0070,0010) VR=FL VM=2 Bounding Box Top Left Hand Corner
	BoundingBoxTopLeftHandCorner = New(0x0070, 0x0010)

	// BoundingBoxBottomRightHandCorner (0070,0011) VR=FL VM=2 Bounding Box Bottom Right Hand Corner
	BoundingBoxBottomRightHandCorner = New(0x0070, 0x0011)

	// BoundingBoxTextHorizontalJustification (0070,0012) VR=CS VM=1 Bounding Box Text Horizontal Justification
	BoundingBoxTextHorizontalJustification = New(0x0070, 0x0012)

	// AnchorPoint (0070,0014) VR=FL VM=2 Anchor Point
	AnchorPoint = New(0x0070, 0x0014)

	// AnchorPointVisibility (0070,0015) VR=CS VM=1 Anchor Point Visibility
	AnchorPointVisibility = New(0x0070, 0x0015)

	// GraphicDimensions (0070,0020) VR=US VM=1 Graphic Dimensions
	GraphicDimensions = New(0x0070, 0x0020)

	// NumberOfGraphicPoints (0070,0021) VR=US VM=1 Number of Graphic Points
	NumberOfGraphicPoints = New(0x0070, 0x0021)

	// GraphicData (0070,0022) VR=FL VM=2-n Graphic Data
	GraphicData = New(0x0070, 0x0022)

	// GraphicType (0070,0023) VR=CS VM=1 Graphic Type
	GraphicType = New(0x0070, 0x0023)

	// GraphicFilled (0070,0024) VR=CS VM=1 Graphic Filled
	GraphicFilled = New(0x0070, 0x0024)

	// ImageRotationRetiredRETIRED (0070,0040) VR=IS VM=1 Image Rotation (Retired) (RETIRED)
	ImageRotationRetiredRETIRED = New(0x0070, 0x0040)

	// ImageHorizontalFlip (0070,0041) VR=CS VM=1 Image Horizontal Flip
	ImageHorizontalFlip = New(0x0070, 0x0041)

	// ImageRotation (0070,0042) VR=US VM=1 Image Rotation
	ImageRotation = New(0x0070, 0x0042)

	// DisplayedAreaTopLeftHandCornerTrialRETIRED (0070,0050) VR=US VM=2 Displayed Area Top Left Hand Corner (Trial) (RETIRED)
	DisplayedAreaTopLeftHandCornerTrialRETIRED = New(0x0070, 0x0050)

	// DisplayedAreaBottomRightHandCornerTrialRETIRED (0070,0051) VR=US VM=2 Displayed Area Bottom Right Hand Corner (Trial) (RETIRED)
	DisplayedAreaBottomRightHandCornerTrialRETIRED = New(0x0070, 0x0051)

	// DisplayedAreaTopLeftHandCorner (0070,0052) VR=SL VM=2 Displayed Area Top Left Hand Corner
	DisplayedAreaTopLeftHandCorner = New(0x0070, 0x0052)

	// DisplayedAreaBottomRightHandCorner (0070,0053) VR=SL VM=2 Displayed Area Bottom Right Hand Corner
	DisplayedAreaBottomRightHandCorner = New(0x0070, 0x0053)

	// DisplayedAreaSelectionSequence (0070,005A) VR=SQ VM=1 Displayed Area Selection Sequence
	DisplayedAreaSelectionSequence = New(0x0070, 0x005A)

	// GraphicLayerSequence (0070,0060) VR=SQ VM=1 Graphic Layer Sequence
	GraphicLayerSequence = New(0x0070, 0x0060)

	// GraphicLayerOrder (0070,0062) VR=IS VM=1 Graphic Layer Order
	GraphicLayerOrder = New(0x0070, 0x0062)

	// GraphicLayerRecommendedDisplayGrayscaleValue (0070,0066) VR=US VM=1 Graphic Layer Recommended Display Grayscale Value
	GraphicLayerRecommendedDisplayGrayscaleValue = New(0x0070, 0x0066)

	// GraphicLayerRecommendedDisplayRGBValueRETIRED (0070,0067) VR=US VM=3 Graphic Layer Recommended Display RGB Value (RETIRED)
	GraphicLayerRecommendedDisplayRGBValueRETIRED = New(0x0070, 0x0067)

	// GraphicLayerDescription (0070,0068) VR=LO VM=1 Graphic Layer Description
	GraphicLayerDescription = New(0x0070, 0x0068)

	// ContentLabel (0070,0080) VR=CS VM=1 Content Label
	ContentLabel = New(0x0070, 0x0080)

	// ContentDescription (0070,0081) VR=LO VM=1 Content Description
	ContentDescription = New(0x0070, 0x0081)

	// PresentationCreationDate (0070,0082) VR=DA VM=1 Presentation Creation Date
	PresentationCreationDate = New(0x0070, 0x0082)

	// PresentationCreationTime (0070,0083) VR=TM VM=1 Presentation Creation Time
	PresentationCreationTime = New(0x0070, 0x0083)

	// ContentCreatorName (0070,0084) VR=PN VM=1 Content Creator's Name
	ContentCreatorName = New(0x0070, 0x0084)

	// ContentCreatorIdentificationCodeSequence (0070,0086) VR=SQ VM=1 Content Creator's Identification Code Sequence
	ContentCreatorIdentificationCodeSequence = New(0x0070, 0x0086)

	// AlternateContentDescriptionSequence (0070,0087) VR=SQ VM=1 Alternate Content Description Sequence
	AlternateContentDescriptionSequence = New(0x0070, 0x0087)

	// PresentationSizeMode (0070,0100) VR=CS VM=1 Presentation Size Mode
	PresentationSizeMode = New(0x0070, 0x0100)

	// PresentationPixelSpacing (0070,0101) VR=DS VM=2 Presentation Pixel Spacing
	PresentationPixelSpacing = New(0x0070, 0x0101)

	// PresentationPixelAspectRatio (0070,0102) VR=IS VM=2 Presentation Pixel Aspect Ratio
	PresentationPixelAspectRatio = New(0x0070, 0x0102)

	// PresentationPixelMagnificationRatio (0070,0103) VR=FL VM=1 Presentation Pixel Magnification Ratio
	PresentationPixelMagnificationRatio = New(0x0070, 0x0103)

	// GraphicGroupLabel (0070,0207) VR=LO VM=1 Graphic Group Label
	GraphicGroupLabel = New(0x0070, 0x0207)

	// GraphicGroupDescription (0070,0208) VR=ST VM=1 Graphic Group Description
	GraphicGroupDescription = New(0x0070, 0x0208)

	// CompoundGraphicSequence (0070,0209) VR=SQ VM=1 Compound Graphic Sequence
	CompoundGraphicSequence = New(0x0070, 0x0209)

	// CompoundGraphicInstanceID (0070,0226) VR=UL VM=1 Compound Graphic Instance ID
	CompoundGraphicInstanceID = New(0x0070, 0x0226)

	// FontName (0070,0227) VR=LO VM=1 Font Name
	FontName = New(0x0070, 0x0227)

	// FontNameType (0070,0228) VR=CS VM=1 Font Name Type
	FontNameType = New(0x0070, 0x0228)

	// CSSFontName (0070,0229) VR=LO VM=1 CSS Font Name
	CSSFontName = New(0x0070, 0x0229)

	// RotationAngle (0070,0230) VR=FD VM=1 Rotation Angle
	RotationAngle = New(0x0070, 0x0230)

	// TextStyleSequence (0070,0231) VR=SQ VM=1 Text Style Sequence
	TextStyleSequence = New(0x0070, 0x0231)

	// LineStyleSequence (0070,0232) VR=SQ VM=1 Line Style Sequence
	LineStyleSequence = New(0x0070, 0x0232)

	// FillStyleSequence (0070,0233) VR=SQ VM=1 Fill Style Sequence
	FillStyleSequence = New(0x0070, 0x0233)

	// GraphicGroupSequence (0070,0234) VR=SQ VM=1 Graphic Group Sequence
	GraphicGroupSequence = New(0x0070, 0x0234)

	// TextColorCIELabValue (0070,0241) VR=US VM=3 Text Color CIELab Value
	TextColorCIELabValue = New(0x0070, 0x0241)

	// HorizontalAlignment (0070,0242) VR=CS VM=1 Horizontal Alignment
	HorizontalAlignment = New(0x0070, 0x0242)

	// VerticalAlignment (0070,0243) VR=CS VM=1 Vertical Alignment
	VerticalAlignment = New(0x0070, 0x0243)

	// ShadowStyle (0070,0244) VR=CS VM=1 Shadow Style
	ShadowStyle = New(0x0070, 0x0244)

	// ShadowOffsetX (0070,0245) VR=FL VM=1 Shadow Offset X
	ShadowOffsetX = New(0x0070, 0x0245)

	// ShadowOffsetY (0070,0246) VR=FL VM=1 Shadow Offset Y
	ShadowOffsetY = New(0x0070, 0x0246)

	// ShadowColorCIELabValue (0070,0247) VR=US VM=3 Shadow Color CIELab Value
	ShadowColorCIELabValue = New(0x0070, 0x0247)

	// Underlined (0070,0248) VR=CS VM=1 Underlined
	Underlined = New(0x0070, 0x0248)

	// Bold (0070,0249) VR=CS VM=1 Bold
	Bold = New(0x0070, 0x0249)

	// Italic (0070,0250) VR=CS VM=1 Italic
	Italic = New(0x0070, 0x0250)

	// PatternOnColorCIELabValue (0070,0251) VR=US VM=3 Pattern On Color CIELab Value
	PatternOnColorCIELabValue = New(0x0070, 0x0251)

	// PatternOffColorCIELabValue (0070,0252) VR=US VM=3 Pattern Off Color CIELab Value
	PatternOffColorCIELabValue = New(0x0070, 0x0252)

	// LineThickness (0070,0253) VR=FL VM=1 Line Thickness
	LineThickness = New(0x0070, 0x0253)

	// LineDashingStyle (0070,0254) VR=CS VM=1 Line Dashing Style
	LineDashingStyle = New(0x0070, 0x0254)

	// LinePattern (0070,0255) VR=UL VM=1 Line Pattern
	LinePattern = New(0x0070, 0x0255)

	// FillPattern (0070,0256) VR=OB VM=1 Fill Pattern
	FillPattern = New(0x0070, 0x0256)

	// FillMode (0070,0257) VR=CS VM=1 Fill Mode
	FillMode = New(0x0070, 0x0257)

	// ShadowOpacity (0070,0258) VR=FL VM=1 Shadow Opacity
	ShadowOpacity = New(0x0070, 0x0258)

	// GapLength (0070,0261) VR=FL VM=1 Gap Length
	GapLength = New(0x0070, 0x0261)

	// DiameterOfVisibility (0070,0262) VR=FL VM=1 Diameter of Visibility
	DiameterOfVisibility = New(0x0070, 0x0262)

	// RotationPoint (0070,0273) VR=FL VM=2 Rotation Point
	RotationPoint = New(0x0070, 0x0273)

	// TickAlignment (0070,0274) VR=CS VM=1 Tick Alignment
	TickAlignment = New(0x0070, 0x0274)

	// ShowTickLabel (0070,0278) VR=CS VM=1 Show Tick Label
	ShowTickLabel = New(0x0070, 0x0278)

	// TickLabelAlignment (0070,0279) VR=CS VM=1 Tick Label Alignment
	TickLabelAlignment = New(0x0070, 0x0279)

	// CompoundGraphicUnits (0070,0282) VR=CS VM=1 Compound Graphic Units
	CompoundGraphicUnits = New(0x0070, 0x0282)

	// PatternOnOpacity (0070,0284) VR=FL VM=1 Pattern On Opacity
	PatternOnOpacity = New(0x0070, 0x0284)

	// PatternOffOpacity (0070,0285) VR=FL VM=1 Pattern Off Opacity
	PatternOffOpacity = New(0x0070, 0x0285)

	// MajorTicksSequence (0070,0287) VR=SQ VM=1 Major Ticks Sequence
	MajorTicksSequence = New(0x0070, 0x0287)

	// TickPosition (0070,0288) VR=FL VM=1 Tick Position
	TickPosition = New(0x0070, 0x0288)

	// TickLabel (0070,0289) VR=SH VM=1 Tick Label
	TickLabel = New(0x0070, 0x0289)

	// CompoundGraphicType (0070,0294) VR=CS VM=1 Compound Graphic Type
	CompoundGraphicType = New(0x0070, 0x0294)

	// GraphicGroupID (0070,0295) VR=UL VM=1 Graphic Group ID
	GraphicGroupID = New(0x0070, 0x0295)

	// ShapeType (0070,0306) VR=CS VM=1 Shape Type
	ShapeType = New(0x0070, 0x0306)

	// RegistrationSequence (0070,0308) VR=SQ VM=1 Registration Sequence
	RegistrationSequence = New(0x0070, 0x0308)

	// MatrixRegistrationSequence (0070,0309) VR=SQ VM=1 Matrix Registration Sequence
	MatrixRegistrationSequence = New(0x0070, 0x0309)

	// MatrixSequence (0070,030A) VR=SQ VM=1 Matrix Sequence
	MatrixSequence = New(0x0070, 0x030A)

	// FrameOfReferenceToDisplayedCoordinateSystemTransformationMatrix (0070,030B) VR=FD VM=16 Frame of Reference to Displayed Coordinate System Transformation Matrix
	FrameOfReferenceToDisplayedCoordinateSystemTransformationMatrix = New(0x0070, 0x030B)

	// FrameOfReferenceTransformationMatrixType (0070,030C) VR=CS VM=1 Frame of Reference Transformation Matrix Type
	FrameOfReferenceTransformationMatrixType = New(0x0070, 0x030C)

	// RegistrationTypeCodeSequence (0070,030D) VR=SQ VM=1 Registration Type Code Sequence
	RegistrationTypeCodeSequence = New(0x0070, 0x030D)

	// FiducialDescription (0070,030F) VR=ST VM=1 Fiducial Description
	FiducialDescription = New(0x0070, 0x030F)

	// FiducialIdentifier (0070,0310) VR=SH VM=1 Fiducial Identifier
	FiducialIdentifier = New(0x0070, 0x0310)

	// FiducialIdentifierCodeSequence (0070,0311) VR=SQ VM=1 Fiducial Identifier Code Sequence
	FiducialIdentifierCodeSequence = New(0x0070, 0x0311)

	// ContourUncertaintyRadius (0070,0312) VR=FD VM=1 Contour Uncertainty Radius
	ContourUncertaintyRadius = New(0x0070, 0x0312)

	// UsedFiducialsSequence (0070,0314) VR=SQ VM=1 Used Fiducials Sequence
	UsedFiducialsSequence = New(0x0070, 0x0314)

	// UsedRTStructureSetROISequence (0070,0315) VR=SQ VM=1 Used RT Structure Set ROI Sequence
	UsedRTStructureSetROISequence = New(0x0070, 0x0315)

	// GraphicCoordinatesDataSequence (0070,0318) VR=SQ VM=1 Graphic Coordinates Data Sequence
	GraphicCoordinatesDataSequence = New(0x0070, 0x0318)

	// FiducialUID (0070,031A) VR=UI VM=1 Fiducial UID
	FiducialUID = New(0x0070, 0x031A)

	// ReferencedFiducialUID (0070,031B) VR=UI VM=1 Referenced Fiducial UID
	ReferencedFiducialUID = New(0x0070, 0x031B)

	// FiducialSetSequence (0070,031C) VR=SQ VM=1 Fiducial Set Sequence
	FiducialSetSequence = New(0x0070, 0x031C)

	// FiducialSequence (0070,031E) VR=SQ VM=1 Fiducial Sequence
	FiducialSequence = New(0x0070, 0x031E)

	// FiducialsPropertyCategoryCodeSequence (0070,031F) VR=SQ VM=1 Fiducials Property Category Code Sequence
	FiducialsPropertyCategoryCodeSequence = New(0x0070, 0x031F)

	// GraphicLayerRecommendedDisplayCIELabValue (0070,0401) VR=US VM=3 Graphic Layer Recommended Display CIELab Value
	GraphicLayerRecommendedDisplayCIELabValue = New(0x0070, 0x0401)

	// BlendingSequence (0070,0402) VR=SQ VM=1 Blending Sequence
	BlendingSequence = New(0x0070, 0x0402)

	// RelativeOpacity (0070,0403) VR=FL VM=1 Relative Opacity
	RelativeOpacity = New(0x0070, 0x0403)

	// ReferencedSpatialRegistrationSequence (0070,0404) VR=SQ VM=1 Referenced Spatial Registration Sequence
	ReferencedSpatialRegistrationSequence = New(0x0070, 0x0404)

	// BlendingPosition (0070,0405) VR=CS VM=1 Blending Position
	BlendingPosition = New(0x0070, 0x0405)

	// PresentationDisplayCollectionUID (0070,1101) VR=UI VM=1 Presentation Display Collection UID
	PresentationDisplayCollectionUID = New(0x0070, 0x1101)

	// PresentationSequenceCollectionUID (0070,1102) VR=UI VM=1 Presentation Sequence Collection UID
	PresentationSequenceCollectionUID = New(0x0070, 0x1102)

	// PresentationSequencePositionIndex (0070,1103) VR=US VM=1 Presentation Sequence Position Index
	PresentationSequencePositionIndex = New(0x0070, 0x1103)

	// RenderedImageReferenceSequence (0070,1104) VR=SQ VM=1 Rendered Image Reference Sequence
	RenderedImageReferenceSequence = New(0x0070, 0x1104)

	// VolumetricPresentationStateInputSequence (0070,1201) VR=SQ VM=1 Volumetric Presentation State Input Sequence
	VolumetricPresentationStateInputSequence = New(0x0070, 0x1201)

	// PresentationInputType (0070,1202) VR=CS VM=1 Presentation Input Type
	PresentationInputType = New(0x0070, 0x1202)

	// InputSequencePositionIndex (0070,1203) VR=US VM=1 Input Sequence Position Index
	InputSequencePositionIndex = New(0x0070, 0x1203)

	// Crop (0070,1204) VR=CS VM=1 Crop
	Crop = New(0x0070, 0x1204)

	// CroppingSpecificationIndex (0070,1205) VR=US VM=1-n Cropping Specification Index
	CroppingSpecificationIndex = New(0x0070, 0x1205)

	// CompositingMethodRETIRED (0070,1206) VR=CS VM=1 Compositing Method (RETIRED)
	CompositingMethodRETIRED = New(0x0070, 0x1206)

	// VolumetricPresentationInputNumber (0070,1207) VR=US VM=1 Volumetric Presentation Input Number
	VolumetricPresentationInputNumber = New(0x0070, 0x1207)

	// ImageVolumeGeometry (0070,1208) VR=CS VM=1 Image Volume Geometry
	ImageVolumeGeometry = New(0x0070, 0x1208)

	// VolumetricPresentationInputSetUID (0070,1209) VR=UI VM=1 Volumetric Presentation Input Set UID
	VolumetricPresentationInputSetUID = New(0x0070, 0x1209)

	// VolumetricPresentationInputSetSequence (0070,120A) VR=SQ VM=1 Volumetric Presentation Input Set Sequence
	VolumetricPresentationInputSetSequence = New(0x0070, 0x120A)

	// GlobalCrop (0070,120B) VR=CS VM=1 Global Crop
	GlobalCrop = New(0x0070, 0x120B)

	// GlobalCroppingSpecificationIndex (0070,120C) VR=US VM=1-n Global Cropping Specification Index
	GlobalCroppingSpecificationIndex = New(0x0070, 0x120C)

	// RenderingMethod (0070,120D) VR=CS VM=1 Rendering Method
	RenderingMethod = New(0x0070, 0x120D)

	// VolumeCroppingSequence (0070,1301) VR=SQ VM=1 Volume Cropping Sequence
	VolumeCroppingSequence = New(0x0070, 0x1301)

	// VolumeCroppingMethod (0070,1302) VR=CS VM=1 Volume Cropping Method
	VolumeCroppingMethod = New(0x0070, 0x1302)

	// BoundingBoxCrop (0070,1303) VR=FD VM=6 Bounding Box Crop
	BoundingBoxCrop = New(0x0070, 0x1303)

	// ObliqueCroppingPlaneSequence (0070,1304) VR=SQ VM=1 Oblique Cropping Plane Sequence
	ObliqueCroppingPlaneSequence = New(0x0070, 0x1304)

	// Plane (0070,1305) VR=FD VM=4 Plane
	Plane = New(0x0070, 0x1305)

	// PlaneNormal (0070,1306) VR=FD VM=3 Plane Normal
	PlaneNormal = New(0x0070, 0x1306)

	// CroppingSpecificationNumber (0070,1309) VR=US VM=1 Cropping Specification Number
	CroppingSpecificationNumber = New(0x0070, 0x1309)

	// MultiPlanarReconstructionStyle (0070,1501) VR=CS VM=1 Multi-Planar Reconstruction Style
	MultiPlanarReconstructionStyle = New(0x0070, 0x1501)

	// MPRThicknessType (0070,1502) VR=CS VM=1 MPR Thickness Type
	MPRThicknessType = New(0x0070, 0x1502)

	// MPRSlabThickness (0070,1503) VR=FD VM=1 MPR Slab Thickness
	MPRSlabThickness = New(0x0070, 0x1503)

	// MPRTopLeftHandCorner (0070,1505) VR=FD VM=3 MPR Top Left Hand Corner
	MPRTopLeftHandCorner = New(0x0070, 0x1505)

	// MPRViewWidthDirection (0070,1507) VR=FD VM=3 MPR View Width Direction
	MPRViewWidthDirection = New(0x0070, 0x1507)

	// MPRViewWidth (0070,1508) VR=FD VM=1 MPR View Width
	MPRViewWidth = New(0x0070, 0x1508)

	// NumberOfVolumetricCurvePoints (0070,150C) VR=UL VM=1 Number of Volumetric Curve Points
	NumberOfVolumetricCurvePoints = New(0x0070, 0x150C)

	// VolumetricCurvePoints (0070,150D) VR=OD VM=1 Volumetric Curve Points
	VolumetricCurvePoints = New(0x0070, 0x150D)

	// MPRViewHeightDirection (0070,1511) VR=FD VM=3 MPR View Height Direction
	MPRViewHeightDirection = New(0x0070, 0x1511)

	// MPRViewHeight (0070,1512) VR=FD VM=1 MPR View Height
	MPRViewHeight = New(0x0070, 0x1512)

	// RenderProjection (0070,1602) VR=CS VM=1 Render Projection
	RenderProjection = New(0x0070, 0x1602)

	// ViewpointPosition (0070,1603) VR=FD VM=3 Viewpoint Position
	ViewpointPosition = New(0x0070, 0x1603)

	// ViewpointLookAtPoint (0070,1604) VR=FD VM=3 Viewpoint LookAt Point
	ViewpointLookAtPoint = New(0x0070, 0x1604)

	// ViewpointUpDirection (0070,1605) VR=FD VM=3 Viewpoint Up Direction
	ViewpointUpDirection = New(0x0070, 0x1605)

	// RenderFieldOfView (0070,1606) VR=FD VM=6 Render Field of View
	RenderFieldOfView = New(0x0070, 0x1606)

	// SamplingStepSize (0070,1607) VR=FD VM=1 Sampling Step Size
	SamplingStepSize = New(0x0070, 0x1607)

	// ShadingStyle (0070,1701) VR=CS VM=1 Shading Style
	ShadingStyle = New(0x0070, 0x1701)

	// AmbientReflectionIntensity (0070,1702) VR=FD VM=1 Ambient Reflection Intensity
	AmbientReflectionIntensity = New(0x0070, 0x1702)

	// LightDirection (0070,1703) VR=FD VM=3 Light Direction
	LightDirection = New(0x0070, 0x1703)

	// DiffuseReflectionIntensity (0070,1704) VR=FD VM=1 Diffuse Reflection Intensity
	DiffuseReflectionIntensity = New(0x0070, 0x1704)

	// SpecularReflectionIntensity (0070,1705) VR=FD VM=1 Specular Reflection Intensity
	SpecularReflectionIntensity = New(0x0070, 0x1705)

	// Shininess (0070,1706) VR=FD VM=1 Shininess
	Shininess = New(0x0070, 0x1706)

	// PresentationStateClassificationComponentSequence (0070,1801) VR=SQ VM=1 Presentation State Classification Component Sequence
	PresentationStateClassificationComponentSequence = New(0x0070, 0x1801)

	// ComponentType (0070,1802) VR=CS VM=1 Component Type
	ComponentType = New(0x0070, 0x1802)

	// ComponentInputSequence (0070,1803) VR=SQ VM=1 Component Input Sequence
	ComponentInputSequence = New(0x0070, 0x1803)

	// VolumetricPresentationInputIndex (0070,1804) VR=US VM=1 Volumetric Presentation Input Index
	VolumetricPresentationInputIndex = New(0x0070, 0x1804)

	// PresentationStateCompositorComponentSequence (0070,1805) VR=SQ VM=1 Presentation State Compositor Component Sequence
	PresentationStateCompositorComponentSequence = New(0x0070, 0x1805)

	// WeightingTransferFunctionSequence (0070,1806) VR=SQ VM=1 Weighting Transfer Function Sequence
	WeightingTransferFunctionSequence = New(0x0070, 0x1806)

	// WeightingLookupTableDescriptorRETIRED (0070,1807) VR=US VM=3 Weighting Lookup Table Descriptor (RETIRED)
	WeightingLookupTableDescriptorRETIRED = New(0x0070, 0x1807)

	// WeightingLookupTableDataRETIRED (0070,1808) VR=OB VM=1 Weighting Lookup Table Data (RETIRED)
	WeightingLookupTableDataRETIRED = New(0x0070, 0x1808)

	// VolumetricAnnotationSequence (0070,1901) VR=SQ VM=1 Volumetric Annotation Sequence
	VolumetricAnnotationSequence = New(0x0070, 0x1901)

	// ReferencedStructuredContextSequence (0070,1903) VR=SQ VM=1 Referenced Structured Context Sequence
	ReferencedStructuredContextSequence = New(0x0070, 0x1903)

	// ReferencedContentItem (0070,1904) VR=UI VM=1 Referenced Content Item
	ReferencedContentItem = New(0x0070, 0x1904)

	// VolumetricPresentationInputAnnotationSequence (0070,1905) VR=SQ VM=1 Volumetric Presentation Input Annotation Sequence
	VolumetricPresentationInputAnnotationSequence = New(0x0070, 0x1905)

	// AnnotationClipping (0070,1907) VR=CS VM=1 Annotation Clipping
	AnnotationClipping = New(0x0070, 0x1907)

	// PresentationAnimationStyle (0070,1A01) VR=CS VM=1 Presentation Animation Style
	PresentationAnimationStyle = New(0x0070, 0x1A01)

	// RecommendedAnimationRate (0070,1A03) VR=FD VM=1 Recommended Animation Rate
	RecommendedAnimationRate = New(0x0070, 0x1A03)

	// AnimationCurveSequence (0070,1A04) VR=SQ VM=1 Animation Curve Sequence
	AnimationCurveSequence = New(0x0070, 0x1A04)

	// AnimationStepSize (0070,1A05) VR=FD VM=1 Animation Step Size
	AnimationStepSize = New(0x0070, 0x1A05)

	// SwivelRange (0070,1A06) VR=FD VM=1 Swivel Range
	SwivelRange = New(0x0070, 0x1A06)

	// VolumetricCurveUpDirections (0070,1A07) VR=OD VM=1 Volumetric Curve Up Directions
	VolumetricCurveUpDirections = New(0x0070, 0x1A07)

	// VolumeStreamSequence (0070,1A08) VR=SQ VM=1 Volume Stream Sequence
	VolumeStreamSequence = New(0x0070, 0x1A08)

	// RGBATransferFunctionDescription (0070,1A09) VR=LO VM=1 RGBA Transfer Function Description
	RGBATransferFunctionDescription = New(0x0070, 0x1A09)

	// AdvancedBlendingSequence (0070,1B01) VR=SQ VM=1 Advanced Blending Sequence
	AdvancedBlendingSequence = New(0x0070, 0x1B01)

	// BlendingInputNumber (0070,1B02) VR=US VM=1 Blending Input Number
	BlendingInputNumber = New(0x0070, 0x1B02)

	// BlendingDisplayInputSequence (0070,1B03) VR=SQ VM=1 Blending Display Input Sequence
	BlendingDisplayInputSequence = New(0x0070, 0x1B03)

	// BlendingDisplaySequence (0070,1B04) VR=SQ VM=1 Blending Display Sequence
	BlendingDisplaySequence = New(0x0070, 0x1B04)

	// BlendingMode (0070,1B06) VR=CS VM=1 Blending Mode
	BlendingMode = New(0x0070, 0x1B06)

	// TimeSeriesBlending (0070,1B07) VR=CS VM=1 Time Series Blending
	TimeSeriesBlending = New(0x0070, 0x1B07)

	// GeometryForDisplay (0070,1B08) VR=CS VM=1 Geometry for Display
	GeometryForDisplay = New(0x0070, 0x1B08)

	// ThresholdSequence (0070,1B11) VR=SQ VM=1 Threshold Sequence
	ThresholdSequence = New(0x0070, 0x1B11)

	// ThresholdValueSequence (0070,1B12) VR=SQ VM=1 Threshold Value Sequence
	ThresholdValueSequence = New(0x0070, 0x1B12)

	// ThresholdType (0070,1B13) VR=CS VM=1 Threshold Type
	ThresholdType = New(0x0070, 0x1B13)

	// ThresholdValue (0070,1B14) VR=FD VM=1 Threshold Value
	ThresholdValue = New(0x0070, 0x1B14)

	// HangingProtocolName (0072,0002) VR=SH VM=1 Hanging Protocol Name
	HangingProtocolName = New(0x0072, 0x0002)

	// HangingProtocolDescription (0072,0004) VR=LO VM=1 Hanging Protocol Description
	HangingProtocolDescription = New(0x0072, 0x0004)

	// HangingProtocolLevel (0072,0006) VR=CS VM=1 Hanging Protocol Level
	HangingProtocolLevel = New(0x0072, 0x0006)

	// HangingProtocolCreator (0072,0008) VR=LO VM=1 Hanging Protocol Creator
	HangingProtocolCreator = New(0x0072, 0x0008)

	// HangingProtocolCreationDateTime (0072,000A) VR=DT VM=1 Hanging Protocol Creation DateTime
	HangingProtocolCreationDateTime = New(0x0072, 0x000A)

	// HangingProtocolDefinitionSequence (0072,000C) VR=SQ VM=1 Hanging Protocol Definition Sequence
	HangingProtocolDefinitionSequence = New(0x0072, 0x000C)

	// HangingProtocolUserIdentificationCodeSequence (0072,000E) VR=SQ VM=1 Hanging Protocol User Identification Code Sequence
	HangingProtocolUserIdentificationCodeSequence = New(0x0072, 0x000E)

	// HangingProtocolUserGroupName (0072,0010) VR=LO VM=1 Hanging Protocol User Group Name
	HangingProtocolUserGroupName = New(0x0072, 0x0010)

	// SourceHangingProtocolSequence (0072,0012) VR=SQ VM=1 Source Hanging Protocol Sequence
	SourceHangingProtocolSequence = New(0x0072, 0x0012)

	// NumberOfPriorsReferenced (0072,0014) VR=US VM=1 Number of Priors Referenced
	NumberOfPriorsReferenced = New(0x0072, 0x0014)

	// ImageSetsSequence (0072,0020) VR=SQ VM=1 Image Sets Sequence
	ImageSetsSequence = New(0x0072, 0x0020)

	// ImageSetSelectorSequence (0072,0022) VR=SQ VM=1 Image Set Selector Sequence
	ImageSetSelectorSequence = New(0x0072, 0x0022)

	// ImageSetSelectorUsageFlag (0072,0024) VR=CS VM=1 Image Set Selector Usage Flag
	ImageSetSelectorUsageFlag = New(0x0072, 0x0024)

	// SelectorAttribute (0072,0026) VR=AT VM=1 Selector Attribute
	SelectorAttribute = New(0x0072, 0x0026)

	// SelectorValueNumber (0072,0028) VR=US VM=1 Selector Value Number
	SelectorValueNumber = New(0x0072, 0x0028)

	// TimeBasedImageSetsSequence (0072,0030) VR=SQ VM=1 Time Based Image Sets Sequence
	TimeBasedImageSetsSequence = New(0x0072, 0x0030)

	// ImageSetNumber (0072,0032) VR=US VM=1 Image Set Number
	ImageSetNumber = New(0x0072, 0x0032)

	// ImageSetSelectorCategory (0072,0034) VR=CS VM=1 Image Set Selector Category
	ImageSetSelectorCategory = New(0x0072, 0x0034)

	// RelativeTime (0072,0038) VR=US VM=2 Relative Time
	RelativeTime = New(0x0072, 0x0038)

	// RelativeTimeUnits (0072,003A) VR=CS VM=1 Relative Time Units
	RelativeTimeUnits = New(0x0072, 0x003A)

	// AbstractPriorValue (0072,003C) VR=SS VM=2 Abstract Prior Value
	AbstractPriorValue = New(0x0072, 0x003C)

	// AbstractPriorCodeSequence (0072,003E) VR=SQ VM=1 Abstract Prior Code Sequence
	AbstractPriorCodeSequence = New(0x0072, 0x003E)

	// ImageSetLabel (0072,0040) VR=LO VM=1 Image Set Label
	ImageSetLabel = New(0x0072, 0x0040)

	// SelectorAttributeVR (0072,0050) VR=CS VM=1 Selector Attribute VR
	SelectorAttributeVR = New(0x0072, 0x0050)

	// SelectorSequencePointer (0072,0052) VR=AT VM=1-n Selector Sequence Pointer
	SelectorSequencePointer = New(0x0072, 0x0052)

	// SelectorSequencePointerPrivateCreator (0072,0054) VR=LO VM=1-n Selector Sequence Pointer Private Creator
	SelectorSequencePointerPrivateCreator = New(0x0072, 0x0054)

	// SelectorAttributePrivateCreator (0072,0056) VR=LO VM=1 Selector Attribute Private Creator
	SelectorAttributePrivateCreator = New(0x0072, 0x0056)

	// SelectorAEValue (0072,005E) VR=AE VM=1-n Selector AE Value
	SelectorAEValue = New(0x0072, 0x005E)

	// SelectorASValue (0072,005F) VR=AS VM=1-n Selector AS Value
	SelectorASValue = New(0x0072, 0x005F)

	// SelectorATValue (0072,0060) VR=AT VM=1-n Selector AT Value
	SelectorATValue = New(0x0072, 0x0060)

	// SelectorDAValue (0072,0061) VR=DA VM=1-n Selector DA Value
	SelectorDAValue = New(0x0072, 0x0061)

	// SelectorCSValue (0072,0062) VR=CS VM=1-n Selector CS Value
	SelectorCSValue = New(0x0072, 0x0062)

	// SelectorDTValue (0072,0063) VR=DT VM=1-n Selector DT Value
	SelectorDTValue = New(0x0072, 0x0063)

	// SelectorISValue (0072,0064) VR=IS VM=1-n Selector IS Value
	SelectorISValue = New(0x0072, 0x0064)

	// SelectorOBValue (0072,0065) VR=OB VM=1 Selector OB Value
	SelectorOBValue = New(0x0072, 0x0065)

	// SelectorLOValue (0072,0066) VR=LO VM=1-n Selector LO Value
	SelectorLOValue = New(0x0072, 0x0066)

	// SelectorOFValue (0072,0067) VR=OF VM=1 Selector OF Value
	SelectorOFValue = New(0x0072, 0x0067)

	// SelectorLTValue (0072,0068) VR=LT VM=1 Selector LT Value
	SelectorLTValue = New(0x0072, 0x0068)

	// SelectorOWValue (0072,0069) VR=OW VM=1 Selector OW Value
	SelectorOWValue = New(0x0072, 0x0069)

	// SelectorPNValue (0072,006A) VR=PN VM=1-n Selector PN Value
	SelectorPNValue = New(0x0072, 0x006A)

	// SelectorTMValue (0072,006B) VR=TM VM=1-n Selector TM Value
	SelectorTMValue = New(0x0072, 0x006B)

	// SelectorSHValue (0072,006C) VR=SH VM=1-n Selector SH Value
	SelectorSHValue = New(0x0072, 0x006C)

	// SelectorUNValue (0072,006D) VR=UN VM=1 Selector UN Value
	SelectorUNValue = New(0x0072, 0x006D)

	// SelectorSTValue (0072,006E) VR=ST VM=1 Selector ST Value
	SelectorSTValue = New(0x0072, 0x006E)

	// SelectorUCValue (0072,006F) VR=UC VM=1-n Selector UC Value
	SelectorUCValue = New(0x0072, 0x006F)

	// SelectorUTValue (0072,0070) VR=UT VM=1 Selector UT Value
	SelectorUTValue = New(0x0072, 0x0070)

	// SelectorURValue (0072,0071) VR=UR VM=1 Selector UR Value
	SelectorURValue = New(0x0072, 0x0071)

	// SelectorDSValue (0072,0072) VR=DS VM=1-n Selector DS Value
	SelectorDSValue = New(0x0072, 0x0072)

	// SelectorODValue (0072,0073) VR=OD VM=1 Selector OD Value
	SelectorODValue = New(0x0072, 0x0073)

	// SelectorFDValue (0072,0074) VR=FD VM=1-n Selector FD Value
	SelectorFDValue = New(0x0072, 0x0074)

	// SelectorOLValue (0072,0075) VR=OL VM=1 Selector OL Value
	SelectorOLValue = New(0x0072, 0x0075)

	// SelectorFLValue (0072,0076) VR=FL VM=1-n Selector FL Value
	SelectorFLValue = New(0x0072, 0x0076)

	// SelectorULValue (0072,0078) VR=UL VM=1-n Selector UL Value
	SelectorULValue = New(0x0072, 0x0078)

	// SelectorUSValue (0072,007A) VR=US VM=1-n Selector US Value
	SelectorUSValue = New(0x0072, 0x007A)

	// SelectorSLValue (0072,007C) VR=SL VM=1-n Selector SL Value
	SelectorSLValue = New(0x0072, 0x007C)

	// SelectorSSValue (0072,007E) VR=SS VM=1-n Selector SS Value
	SelectorSSValue = New(0x0072, 0x007E)

	// SelectorUIValue (0072,007F) VR=UI VM=1-n Selector UI Value
	SelectorUIValue = New(0x0072, 0x007F)

	// SelectorCodeSequenceValue (0072,0080) VR=SQ VM=1 Selector Code Sequence Value
	SelectorCodeSequenceValue = New(0x0072, 0x0080)

	// SelectorOVValue (0072,0081) VR=OV VM=1 Selector OV Value
	SelectorOVValue = New(0x0072, 0x0081)

	// SelectorSVValue (0072,0082) VR=SV VM=1-n Selector SV Value
	SelectorSVValue = New(0x0072, 0x0082)

	// SelectorUVValue (0072,0083) VR=UV VM=1-n Selector UV Value
	SelectorUVValue = New(0x0072, 0x0083)

	// NumberOfScreens (0072,0100) VR=US VM=1 Number of Screens
	NumberOfScreens = New(0x0072, 0x0100)

	// NominalScreenDefinitionSequence (0072,0102) VR=SQ VM=1 Nominal Screen Definition Sequence
	NominalScreenDefinitionSequence = New(0x0072, 0x0102)

	// NumberOfVerticalPixels (0072,0104) VR=US VM=1 Number of Vertical Pixels
	NumberOfVerticalPixels = New(0x0072, 0x0104)

	// NumberOfHorizontalPixels (0072,0106) VR=US VM=1 Number of Horizontal Pixels
	NumberOfHorizontalPixels = New(0x0072, 0x0106)

	// DisplayEnvironmentSpatialPosition (0072,0108) VR=FD VM=4 Display Environment Spatial Position
	DisplayEnvironmentSpatialPosition = New(0x0072, 0x0108)

	// ScreenMinimumGrayscaleBitDepth (0072,010A) VR=US VM=1 Screen Minimum Grayscale Bit Depth
	ScreenMinimumGrayscaleBitDepth = New(0x0072, 0x010A)

	// ScreenMinimumColorBitDepth (0072,010C) VR=US VM=1 Screen Minimum Color Bit Depth
	ScreenMinimumColorBitDepth = New(0x0072, 0x010C)

	// ApplicationMaximumRepaintTime (0072,010E) VR=US VM=1 Application Maximum Repaint Time
	ApplicationMaximumRepaintTime = New(0x0072, 0x010E)

	// DisplaySetsSequence (0072,0200) VR=SQ VM=1 Display Sets Sequence
	DisplaySetsSequence = New(0x0072, 0x0200)

	// DisplaySetNumber (0072,0202) VR=US VM=1 Display Set Number
	DisplaySetNumber = New(0x0072, 0x0202)

	// DisplaySetLabel (0072,0203) VR=LO VM=1 Display Set Label
	DisplaySetLabel = New(0x0072, 0x0203)

	// DisplaySetPresentationGroup (0072,0204) VR=US VM=1 Display Set Presentation Group
	DisplaySetPresentationGroup = New(0x0072, 0x0204)

	// DisplaySetPresentationGroupDescription (0072,0206) VR=LO VM=1 Display Set Presentation Group Description
	DisplaySetPresentationGroupDescription = New(0x0072, 0x0206)

	// PartialDataDisplayHandling (0072,0208) VR=CS VM=1 Partial Data Display Handling
	PartialDataDisplayHandling = New(0x0072, 0x0208)

	// SynchronizedScrollingSequence (0072,0210) VR=SQ VM=1 Synchronized Scrolling Sequence
	SynchronizedScrollingSequence = New(0x0072, 0x0210)

	// DisplaySetScrollingGroup (0072,0212) VR=US VM=2-n Display Set Scrolling Group
	DisplaySetScrollingGroup = New(0x0072, 0x0212)

	// NavigationIndicatorSequence (0072,0214) VR=SQ VM=1 Navigation Indicator Sequence
	NavigationIndicatorSequence = New(0x0072, 0x0214)

	// NavigationDisplaySet (0072,0216) VR=US VM=1 Navigation Display Set
	NavigationDisplaySet = New(0x0072, 0x0216)

	// ReferenceDisplaySets (0072,0218) VR=US VM=1-n Reference Display Sets
	ReferenceDisplaySets = New(0x0072, 0x0218)

	// ImageBoxesSequence (0072,0300) VR=SQ VM=1 Image Boxes Sequence
	ImageBoxesSequence = New(0x0072, 0x0300)

	// ImageBoxNumber (0072,0302) VR=US VM=1 Image Box Number
	ImageBoxNumber = New(0x0072, 0x0302)

	// ImageBoxLayoutType (0072,0304) VR=CS VM=1 Image Box Layout Type
	ImageBoxLayoutType = New(0x0072, 0x0304)

	// ImageBoxTileHorizontalDimension (0072,0306) VR=US VM=1 Image Box Tile Horizontal Dimension
	ImageBoxTileHorizontalDimension = New(0x0072, 0x0306)

	// ImageBoxTileVerticalDimension (0072,0308) VR=US VM=1 Image Box Tile Vertical Dimension
	ImageBoxTileVerticalDimension = New(0x0072, 0x0308)

	// ImageBoxScrollDirection (0072,0310) VR=CS VM=1 Image Box Scroll Direction
	ImageBoxScrollDirection = New(0x0072, 0x0310)

	// ImageBoxSmallScrollType (0072,0312) VR=CS VM=1 Image Box Small Scroll Type
	ImageBoxSmallScrollType = New(0x0072, 0x0312)

	// ImageBoxSmallScrollAmount (0072,0314) VR=US VM=1 Image Box Small Scroll Amount
	ImageBoxSmallScrollAmount = New(0x0072, 0x0314)

	// ImageBoxLargeScrollType (0072,0316) VR=CS VM=1 Image Box Large Scroll Type
	ImageBoxLargeScrollType = New(0x0072, 0x0316)

	// ImageBoxLargeScrollAmount (0072,0318) VR=US VM=1 Image Box Large Scroll Amount
	ImageBoxLargeScrollAmount = New(0x0072, 0x0318)

	// ImageBoxOverlapPriority (0072,0320) VR=US VM=1 Image Box Overlap Priority
	ImageBoxOverlapPriority = New(0x0072, 0x0320)

	// CineRelativeToRealTime (0072,0330) VR=FD VM=1 Cine Relative to Real-Time
	CineRelativeToRealTime = New(0x0072, 0x0330)

	// FilterOperationsSequence (0072,0400) VR=SQ VM=1 Filter Operations Sequence
	FilterOperationsSequence = New(0x0072, 0x0400)

	// FilterByCategory (0072,0402) VR=CS VM=1 Filter-by Category
	FilterByCategory = New(0x0072, 0x0402)

	// FilterByAttributePresence (0072,0404) VR=CS VM=1 Filter-by Attribute Presence
	FilterByAttributePresence = New(0x0072, 0x0404)

	// FilterByOperator (0072,0406) VR=CS VM=1 Filter-by Operator
	FilterByOperator = New(0x0072, 0x0406)

	// StructuredDisplayBackgroundCIELabValue (0072,0420) VR=US VM=3 Structured Display Background CIELab Value
	StructuredDisplayBackgroundCIELabValue = New(0x0072, 0x0420)

	// EmptyImageBoxCIELabValue (0072,0421) VR=US VM=3 Empty Image Box CIELab Value
	EmptyImageBoxCIELabValue = New(0x0072, 0x0421)

	// StructuredDisplayImageBoxSequence (0072,0422) VR=SQ VM=1 Structured Display Image Box Sequence
	StructuredDisplayImageBoxSequence = New(0x0072, 0x0422)

	// StructuredDisplayTextBoxSequence (0072,0424) VR=SQ VM=1 Structured Display Text Box Sequence
	StructuredDisplayTextBoxSequence = New(0x0072, 0x0424)

	// ReferencedFirstFrameSequence (0072,0427) VR=SQ VM=1 Referenced First Frame Sequence
	ReferencedFirstFrameSequence = New(0x0072, 0x0427)

	// ImageBoxSynchronizationSequence (0072,0430) VR=SQ VM=1 Image Box Synchronization Sequence
	ImageBoxSynchronizationSequence = New(0x0072, 0x0430)

	// SynchronizedImageBoxList (0072,0432) VR=US VM=2-n Synchronized Image Box List
	SynchronizedImageBoxList = New(0x0072, 0x0432)

	// TypeOfSynchronization (0072,0434) VR=CS VM=1 Type of Synchronization
	TypeOfSynchronization = New(0x0072, 0x0434)

	// BlendingOperationType (0072,0500) VR=CS VM=1 Blending Operation Type
	BlendingOperationType = New(0x0072, 0x0500)

	// ReformattingOperationType (0072,0510) VR=CS VM=1 Reformatting Operation Type
	ReformattingOperationType = New(0x0072, 0x0510)

	// ReformattingThickness (0072,0512) VR=FD VM=1 Reformatting Thickness
	ReformattingThickness = New(0x0072, 0x0512)

	// ReformattingInterval (0072,0514) VR=FD VM=1 Reformatting Interval
	ReformattingInterval = New(0x0072, 0x0514)

	// ReformattingOperationInitialViewDirection (0072,0516) VR=CS VM=1 Reformatting Operation Initial View Direction
	ReformattingOperationInitialViewDirection = New(0x0072, 0x0516)

	// ThreeDRenderingType (0072,0520) VR=CS VM=1-n 3D Rendering Type
	ThreeDRenderingType = New(0x0072, 0x0520)

	// SortingOperationsSequence (0072,0600) VR=SQ VM=1 Sorting Operations Sequence
	SortingOperationsSequence = New(0x0072, 0x0600)

	// SortByCategory (0072,0602) VR=CS VM=1 Sort-by Category
	SortByCategory = New(0x0072, 0x0602)

	// SortingDirection (0072,0604) VR=CS VM=1 Sorting Direction
	SortingDirection = New(0x0072, 0x0604)

	// DisplaySetPatientOrientation (0072,0700) VR=CS VM=2 Display Set Patient Orientation
	DisplaySetPatientOrientation = New(0x0072, 0x0700)

	// VOIType (0072,0702) VR=CS VM=1 VOI Type
	VOIType = New(0x0072, 0x0702)

	// PseudoColorType (0072,0704) VR=CS VM=1 Pseudo-Color Type
	PseudoColorType = New(0x0072, 0x0704)

	// PseudoColorPaletteInstanceReferenceSequence (0072,0705) VR=SQ VM=1 Pseudo-Color Palette Instance Reference Sequence
	PseudoColorPaletteInstanceReferenceSequence = New(0x0072, 0x0705)

	// ShowGrayscaleInverted (0072,0706) VR=CS VM=1 Show Grayscale Inverted
	ShowGrayscaleInverted = New(0x0072, 0x0706)

	// ShowImageTrueSizeFlag (0072,0710) VR=CS VM=1 Show Image True Size Flag
	ShowImageTrueSizeFlag = New(0x0072, 0x0710)

	// ShowGraphicAnnotationFlag (0072,0712) VR=CS VM=1 Show Graphic Annotation Flag
	ShowGraphicAnnotationFlag = New(0x0072, 0x0712)

	// ShowPatientDemographicsFlag (0072,0714) VR=CS VM=1 Show Patient Demographics Flag
	ShowPatientDemographicsFlag = New(0x0072, 0x0714)

	// ShowAcquisitionTechniquesFlag (0072,0716) VR=CS VM=1 Show Acquisition Techniques Flag
	ShowAcquisitionTechniquesFlag = New(0x0072, 0x0716)

	// DisplaySetHorizontalJustification (0072,0717) VR=CS VM=1 Display Set Horizontal Justification
	DisplaySetHorizontalJustification = New(0x0072, 0x0717)

	// DisplaySetVerticalJustification (0072,0718) VR=CS VM=1 Display Set Vertical Justification
	DisplaySetVerticalJustification = New(0x0072, 0x0718)

	// ContinuationStartMeterset (0074,0120) VR=FD VM=1 Continuation Start Meterset
	ContinuationStartMeterset = New(0x0074, 0x0120)

	// ContinuationEndMeterset (0074,0121) VR=FD VM=1 Continuation End Meterset
	ContinuationEndMeterset = New(0x0074, 0x0121)

	// ProcedureStepState (0074,1000) VR=CS VM=1 Procedure Step State
	ProcedureStepState = New(0x0074, 0x1000)

	// ProcedureStepProgressInformationSequence (0074,1002) VR=SQ VM=1 Procedure Step Progress Information Sequence
	ProcedureStepProgressInformationSequence = New(0x0074, 0x1002)

	// ProcedureStepProgress (0074,1004) VR=DS VM=1 Procedure Step Progress
	ProcedureStepProgress = New(0x0074, 0x1004)

	// ProcedureStepProgressDescription (0074,1006) VR=ST VM=1 Procedure Step Progress Description
	ProcedureStepProgressDescription = New(0x0074, 0x1006)

	// ProcedureStepProgressParametersSequence (0074,1007) VR=SQ VM=1 Procedure Step Progress Parameters Sequence
	ProcedureStepProgressParametersSequence = New(0x0074, 0x1007)

	// ProcedureStepCommunicationsURISequence (0074,1008) VR=SQ VM=1 Procedure Step Communications URI Sequence
	ProcedureStepCommunicationsURISequence = New(0x0074, 0x1008)

	// ContactURI (0074,100A) VR=UR VM=1 Contact URI
	ContactURI = New(0x0074, 0x100A)

	// ContactDisplayName (0074,100C) VR=LO VM=1 Contact Display Name
	ContactDisplayName = New(0x0074, 0x100C)

	// ProcedureStepDiscontinuationReasonCodeSequence (0074,100E) VR=SQ VM=1 Procedure Step Discontinuation Reason Code Sequence
	ProcedureStepDiscontinuationReasonCodeSequence = New(0x0074, 0x100E)

	// BeamTaskSequence (0074,1020) VR=SQ VM=1 Beam Task Sequence
	BeamTaskSequence = New(0x0074, 0x1020)

	// BeamTaskType (0074,1022) VR=CS VM=1 Beam Task Type
	BeamTaskType = New(0x0074, 0x1022)

	// BeamOrderIndexTrialRETIRED (0074,1024) VR=IS VM=1 Beam Order Index (Trial) (RETIRED)
	BeamOrderIndexTrialRETIRED = New(0x0074, 0x1024)

	// AutosequenceFlag (0074,1025) VR=CS VM=1 Autosequence Flag
	AutosequenceFlag = New(0x0074, 0x1025)

	// TableTopVerticalAdjustedPosition (0074,1026) VR=FD VM=1 Table Top Vertical Adjusted Position
	TableTopVerticalAdjustedPosition = New(0x0074, 0x1026)

	// TableTopLongitudinalAdjustedPosition (0074,1027) VR=FD VM=1 Table Top Longitudinal Adjusted Position
	TableTopLongitudinalAdjustedPosition = New(0x0074, 0x1027)

	// TableTopLateralAdjustedPosition (0074,1028) VR=FD VM=1 Table Top Lateral Adjusted Position
	TableTopLateralAdjustedPosition = New(0x0074, 0x1028)

	// PatientSupportAdjustedAngle (0074,102A) VR=FD VM=1 Patient Support Adjusted Angle
	PatientSupportAdjustedAngle = New(0x0074, 0x102A)

	// TableTopEccentricAdjustedAngle (0074,102B) VR=FD VM=1 Table Top Eccentric Adjusted Angle
	TableTopEccentricAdjustedAngle = New(0x0074, 0x102B)

	// TableTopPitchAdjustedAngle (0074,102C) VR=FD VM=1 Table Top Pitch Adjusted Angle
	TableTopPitchAdjustedAngle = New(0x0074, 0x102C)

	// TableTopRollAdjustedAngle (0074,102D) VR=FD VM=1 Table Top Roll Adjusted Angle
	TableTopRollAdjustedAngle = New(0x0074, 0x102D)

	// DeliveryVerificationImageSequence (0074,1030) VR=SQ VM=1 Delivery Verification Image Sequence
	DeliveryVerificationImageSequence = New(0x0074, 0x1030)

	// VerificationImageTiming (0074,1032) VR=CS VM=1 Verification Image Timing
	VerificationImageTiming = New(0x0074, 0x1032)

	// DoubleExposureFlag (0074,1034) VR=CS VM=1 Double Exposure Flag
	DoubleExposureFlag = New(0x0074, 0x1034)

	// DoubleExposureOrdering (0074,1036) VR=CS VM=1 Double Exposure Ordering
	DoubleExposureOrdering = New(0x0074, 0x1036)

	// DoubleExposureMetersetTrialRETIRED (0074,1038) VR=DS VM=1 Double Exposure Meterset (Trial) (RETIRED)
	DoubleExposureMetersetTrialRETIRED = New(0x0074, 0x1038)

	// DoubleExposureFieldDeltaTrialRETIRED (0074,103A) VR=DS VM=4 Double Exposure Field Delta (Trial) (RETIRED)
	DoubleExposureFieldDeltaTrialRETIRED = New(0x0074, 0x103A)

	// RelatedReferenceRTImageSequence (0074,1040) VR=SQ VM=1 Related Reference RT Image Sequence
	RelatedReferenceRTImageSequence = New(0x0074, 0x1040)

	// GeneralMachineVerificationSequence (0074,1042) VR=SQ VM=1 General Machine Verification Sequence
	GeneralMachineVerificationSequence = New(0x0074, 0x1042)

	// ConventionalMachineVerificationSequence (0074,1044) VR=SQ VM=1 Conventional Machine Verification Sequence
	ConventionalMachineVerificationSequence = New(0x0074, 0x1044)

	// IonMachineVerificationSequence (0074,1046) VR=SQ VM=1 Ion Machine Verification Sequence
	IonMachineVerificationSequence = New(0x0074, 0x1046)

	// FailedAttributesSequence (0074,1048) VR=SQ VM=1 Failed Attributes Sequence
	FailedAttributesSequence = New(0x0074, 0x1048)

	// OverriddenAttributesSequence (0074,104A) VR=SQ VM=1 Overridden Attributes Sequence
	OverriddenAttributesSequence = New(0x0074, 0x104A)

	// ConventionalControlPointVerificationSequence (0074,104C) VR=SQ VM=1 Conventional Control Point Verification Sequence
	ConventionalControlPointVerificationSequence = New(0x0074, 0x104C)

	// IonControlPointVerificationSequence (0074,104E) VR=SQ VM=1 Ion Control Point Verification Sequence
	IonControlPointVerificationSequence = New(0x0074, 0x104E)

	// AttributeOccurrenceSequence (0074,1050) VR=SQ VM=1 Attribute Occurrence Sequence
	AttributeOccurrenceSequence = New(0x0074, 0x1050)

	// AttributeOccurrencePointer (0074,1052) VR=AT VM=1 Attribute Occurrence Pointer
	AttributeOccurrencePointer = New(0x0074, 0x1052)

	// AttributeItemSelector (0074,1054) VR=UL VM=1 Attribute Item Selector
	AttributeItemSelector = New(0x0074, 0x1054)

	// AttributeOccurrencePrivateCreator (0074,1056) VR=LO VM=1 Attribute Occurrence Private Creator
	AttributeOccurrencePrivateCreator = New(0x0074, 0x1056)

	// SelectorSequencePointerItems (0074,1057) VR=IS VM=1-n Selector Sequence Pointer Items
	SelectorSequencePointerItems = New(0x0074, 0x1057)

	// ScheduledProcedureStepPriority (0074,1200) VR=CS VM=1 Scheduled Procedure Step Priority
	ScheduledProcedureStepPriority = New(0x0074, 0x1200)

	// WorklistLabel (0074,1202) VR=LO VM=1 Worklist Label
	WorklistLabel = New(0x0074, 0x1202)

	// ProcedureStepLabel (0074,1204) VR=LO VM=1 Procedure Step Label
	ProcedureStepLabel = New(0x0074, 0x1204)

	// ScheduledProcessingParametersSequence (0074,1210) VR=SQ VM=1 Scheduled Processing Parameters Sequence
	ScheduledProcessingParametersSequence = New(0x0074, 0x1210)

	// PerformedProcessingParametersSequence (0074,1212) VR=SQ VM=1 Performed Processing Parameters Sequence
	PerformedProcessingParametersSequence = New(0x0074, 0x1212)

	// UnifiedProcedureStepPerformedProcedureSequence (0074,1216) VR=SQ VM=1 Unified Procedure Step Performed Procedure Sequence
	UnifiedProcedureStepPerformedProcedureSequence = New(0x0074, 0x1216)

	// RelatedProcedureStepSequenceRETIRED (0074,1220) VR=SQ VM=1 Related Procedure Step Sequence (RETIRED)
	RelatedProcedureStepSequenceRETIRED = New(0x0074, 0x1220)

	// ProcedureStepRelationshipTypeRETIRED (0074,1222) VR=LO VM=1 Procedure Step Relationship Type (RETIRED)
	ProcedureStepRelationshipTypeRETIRED = New(0x0074, 0x1222)

	// ReplacedProcedureStepSequence (0074,1224) VR=SQ VM=1 Replaced Procedure Step Sequence
	ReplacedProcedureStepSequence = New(0x0074, 0x1224)

	// DeletionLock (0074,1230) VR=LO VM=1 Deletion Lock
	DeletionLock = New(0x0074, 0x1230)

	// ReceivingAE (0074,1234) VR=AE VM=1 Receiving AE
	ReceivingAE = New(0x0074, 0x1234)

	// RequestingAE (0074,1236) VR=AE VM=1 Requesting AE
	RequestingAE = New(0x0074, 0x1236)

	// ReasonForCancellation (0074,1238) VR=LT VM=1 Reason for Cancellation
	ReasonForCancellation = New(0x0074, 0x1238)

	// SCPStatus (0074,1242) VR=CS VM=1 SCP Status
	SCPStatus = New(0x0074, 0x1242)

	// SubscriptionListStatus (0074,1244) VR=CS VM=1 Subscription List Status
	SubscriptionListStatus = New(0x0074, 0x1244)

	// UnifiedProcedureStepListStatus (0074,1246) VR=CS VM=1 Unified Procedure Step List Status
	UnifiedProcedureStepListStatus = New(0x0074, 0x1246)

	// BeamOrderIndex (0074,1324) VR=UL VM=1 Beam Order Index
	BeamOrderIndex = New(0x0074, 0x1324)

	// DoubleExposureMeterset (0074,1338) VR=FD VM=1 Double Exposure Meterset
	DoubleExposureMeterset = New(0x0074, 0x1338)

	// DoubleExposureFieldDelta (0074,133A) VR=FD VM=4 Double Exposure Field Delta
	DoubleExposureFieldDelta = New(0x0074, 0x133A)

	// BrachyTaskSequence (0074,1401) VR=SQ VM=1 Brachy Task Sequence
	BrachyTaskSequence = New(0x0074, 0x1401)

	// ContinuationStartTotalReferenceAirKerma (0074,1402) VR=DS VM=1 Continuation Start Total Reference Air Kerma
	ContinuationStartTotalReferenceAirKerma = New(0x0074, 0x1402)

	// ContinuationEndTotalReferenceAirKerma (0074,1403) VR=DS VM=1 Continuation End Total Reference Air Kerma
	ContinuationEndTotalReferenceAirKerma = New(0x0074, 0x1403)

	// ContinuationPulseNumber (0074,1404) VR=IS VM=1 Continuation Pulse Number
	ContinuationPulseNumber = New(0x0074, 0x1404)

	// ChannelDeliveryOrderSequence (0074,1405) VR=SQ VM=1 Channel Delivery Order Sequence
	ChannelDeliveryOrderSequence = New(0x0074, 0x1405)

	// ReferencedChannelNumber (0074,1406) VR=IS VM=1 Referenced Channel Number
	ReferencedChannelNumber = New(0x0074, 0x1406)

	// StartCumulativeTimeWeight (0074,1407) VR=DS VM=1 Start Cumulative Time Weight
	StartCumulativeTimeWeight = New(0x0074, 0x1407)

	// EndCumulativeTimeWeight (0074,1408) VR=DS VM=1 End Cumulative Time Weight
	EndCumulativeTimeWeight = New(0x0074, 0x1408)

	// OmittedChannelSequence (0074,1409) VR=SQ VM=1 Omitted Channel Sequence
	OmittedChannelSequence = New(0x0074, 0x1409)

	// ReasonForChannelOmission (0074,140A) VR=CS VM=1 Reason for Channel Omission
	ReasonForChannelOmission = New(0x0074, 0x140A)

	// ReasonForChannelOmissionDescription (0074,140B) VR=LO VM=1 Reason for Channel Omission Description
	ReasonForChannelOmissionDescription = New(0x0074, 0x140B)

	// ChannelDeliveryOrderIndex (0074,140C) VR=IS VM=1 Channel Delivery Order Index
	ChannelDeliveryOrderIndex = New(0x0074, 0x140C)

	// ChannelDeliveryContinuationSequence (0074,140D) VR=SQ VM=1 Channel Delivery Continuation Sequence
	ChannelDeliveryContinuationSequence = New(0x0074, 0x140D)

	// OmittedApplicationSetupSequence (0074,140E) VR=SQ VM=1 Omitted Application Setup Sequence
	OmittedApplicationSetupSequence = New(0x0074, 0x140E)

	// ImplantAssemblyTemplateName (0076,0001) VR=LO VM=1 Implant Assembly Template Name
	ImplantAssemblyTemplateName = New(0x0076, 0x0001)

	// ImplantAssemblyTemplateIssuer (0076,0003) VR=LO VM=1 Implant Assembly Template Issuer
	ImplantAssemblyTemplateIssuer = New(0x0076, 0x0003)

	// ImplantAssemblyTemplateVersion (0076,0006) VR=LO VM=1 Implant Assembly Template Version
	ImplantAssemblyTemplateVersion = New(0x0076, 0x0006)

	// ReplacedImplantAssemblyTemplateSequence (0076,0008) VR=SQ VM=1 Replaced Implant Assembly Template Sequence
	ReplacedImplantAssemblyTemplateSequence = New(0x0076, 0x0008)

	// ImplantAssemblyTemplateType (0076,000A) VR=CS VM=1 Implant Assembly Template Type
	ImplantAssemblyTemplateType = New(0x0076, 0x000A)

	// OriginalImplantAssemblyTemplateSequence (0076,000C) VR=SQ VM=1 Original Implant Assembly Template Sequence
	OriginalImplantAssemblyTemplateSequence = New(0x0076, 0x000C)

	// DerivationImplantAssemblyTemplateSequence (0076,000E) VR=SQ VM=1 Derivation Implant Assembly Template Sequence
	DerivationImplantAssemblyTemplateSequence = New(0x0076, 0x000E)

	// ImplantAssemblyTemplateTargetAnatomySequence (0076,0010) VR=SQ VM=1 Implant Assembly Template Target Anatomy Sequence
	ImplantAssemblyTemplateTargetAnatomySequence = New(0x0076, 0x0010)

	// ProcedureTypeCodeSequence (0076,0020) VR=SQ VM=1 Procedure Type Code Sequence
	ProcedureTypeCodeSequence = New(0x0076, 0x0020)

	// SurgicalTechnique (0076,0030) VR=LO VM=1 Surgical Technique
	SurgicalTechnique = New(0x0076, 0x0030)

	// ComponentTypesSequence (0076,0032) VR=SQ VM=1 Component Types Sequence
	ComponentTypesSequence = New(0x0076, 0x0032)

	// ComponentTypeCodeSequence (0076,0034) VR=SQ VM=1 Component Type Code Sequence
	ComponentTypeCodeSequence = New(0x0076, 0x0034)

	// ExclusiveComponentType (0076,0036) VR=CS VM=1 Exclusive Component Type
	ExclusiveComponentType = New(0x0076, 0x0036)

	// MandatoryComponentType (0076,0038) VR=CS VM=1 Mandatory Component Type
	MandatoryComponentType = New(0x0076, 0x0038)

	// ComponentSequence (0076,0040) VR=SQ VM=1 Component Sequence
	ComponentSequence = New(0x0076, 0x0040)

	// ComponentID (0076,0055) VR=US VM=1 Component ID
	ComponentID = New(0x0076, 0x0055)

	// ComponentAssemblySequence (0076,0060) VR=SQ VM=1 Component Assembly Sequence
	ComponentAssemblySequence = New(0x0076, 0x0060)

	// Component1ReferencedID (0076,0070) VR=US VM=1 Component 1 Referenced ID
	Component1ReferencedID = New(0x0076, 0x0070)

	// Component1ReferencedMatingFeatureSetID (0076,0080) VR=US VM=1 Component 1 Referenced Mating Feature Set ID
	Component1ReferencedMatingFeatureSetID = New(0x0076, 0x0080)

	// Component1ReferencedMatingFeatureID (0076,0090) VR=US VM=1 Component 1 Referenced Mating Feature ID
	Component1ReferencedMatingFeatureID = New(0x0076, 0x0090)

	// Component2ReferencedID (0076,00A0) VR=US VM=1 Component 2 Referenced ID
	Component2ReferencedID = New(0x0076, 0x00A0)

	// Component2ReferencedMatingFeatureSetID (0076,00B0) VR=US VM=1 Component 2 Referenced Mating Feature Set ID
	Component2ReferencedMatingFeatureSetID = New(0x0076, 0x00B0)

	// Component2ReferencedMatingFeatureID (0076,00C0) VR=US VM=1 Component 2 Referenced Mating Feature ID
	Component2ReferencedMatingFeatureID = New(0x0076, 0x00C0)

	// ImplantTemplateGroupName (0078,0001) VR=LO VM=1 Implant Template Group Name
	ImplantTemplateGroupName = New(0x0078, 0x0001)

	// ImplantTemplateGroupDescription (0078,0010) VR=ST VM=1 Implant Template Group Description
	ImplantTemplateGroupDescription = New(0x0078, 0x0010)

	// ImplantTemplateGroupIssuer (0078,0020) VR=LO VM=1 Implant Template Group Issuer
	ImplantTemplateGroupIssuer = New(0x0078, 0x0020)

	// ImplantTemplateGroupVersion (0078,0024) VR=LO VM=1 Implant Template Group Version
	ImplantTemplateGroupVersion = New(0x0078, 0x0024)

	// ReplacedImplantTemplateGroupSequence (0078,0026) VR=SQ VM=1 Replaced Implant Template Group Sequence
	ReplacedImplantTemplateGroupSequence = New(0x0078, 0x0026)

	// ImplantTemplateGroupTargetAnatomySequence (0078,0028) VR=SQ VM=1 Implant Template Group Target Anatomy Sequence
	ImplantTemplateGroupTargetAnatomySequence = New(0x0078, 0x0028)

	// ImplantTemplateGroupMembersSequence (0078,002A) VR=SQ VM=1 Implant Template Group Members Sequence
	ImplantTemplateGroupMembersSequence = New(0x0078, 0x002A)

	// ImplantTemplateGroupMemberID (0078,002E) VR=US VM=1 Implant Template Group Member ID
	ImplantTemplateGroupMemberID = New(0x0078, 0x002E)

	// ThreeDImplantTemplateGroupMemberMatchingPoint (0078,0050) VR=FD VM=3 3D Implant Template Group Member Matching Point
	ThreeDImplantTemplateGroupMemberMatchingPoint = New(0x0078, 0x0050)

	// ThreeDImplantTemplateGroupMemberMatchingAxes (0078,0060) VR=FD VM=9 3D Implant Template Group Member Matching Axes
	ThreeDImplantTemplateGroupMemberMatchingAxes = New(0x0078, 0x0060)

	// ImplantTemplateGroupMemberMatching2DCoordinatesSequence (0078,0070) VR=SQ VM=1 Implant Template Group Member Matching 2D Coordinates Sequence
	ImplantTemplateGroupMemberMatching2DCoordinatesSequence = New(0x0078, 0x0070)

	// TwoDImplantTemplateGroupMemberMatchingPoint (0078,0090) VR=FD VM=2 2D Implant Template Group Member Matching Point
	TwoDImplantTemplateGroupMemberMatchingPoint = New(0x0078, 0x0090)

	// TwoDImplantTemplateGroupMemberMatchingAxes (0078,00A0) VR=FD VM=4 2D Implant Template Group Member Matching Axes
	TwoDImplantTemplateGroupMemberMatchingAxes = New(0x0078, 0x00A0)

	// ImplantTemplateGroupVariationDimensionSequence (0078,00B0) VR=SQ VM=1 Implant Template Group Variation Dimension Sequence
	ImplantTemplateGroupVariationDimensionSequence = New(0x0078, 0x00B0)

	// ImplantTemplateGroupVariationDimensionName (0078,00B2) VR=LO VM=1 Implant Template Group Variation Dimension Name
	ImplantTemplateGroupVariationDimensionName = New(0x0078, 0x00B2)

	// ImplantTemplateGroupVariationDimensionRankSequence (0078,00B4) VR=SQ VM=1 Implant Template Group Variation Dimension Rank Sequence
	ImplantTemplateGroupVariationDimensionRankSequence = New(0x0078, 0x00B4)

	// ReferencedImplantTemplateGroupMemberID (0078,00B6) VR=US VM=1 Referenced Implant Template Group Member ID
	ReferencedImplantTemplateGroupMemberID = New(0x0078, 0x00B6)

	// ImplantTemplateGroupVariationDimensionRank (0078,00B8) VR=US VM=1 Implant Template Group Variation Dimension Rank
	ImplantTemplateGroupVariationDimensionRank = New(0x0078, 0x00B8)

	// SurfaceScanAcquisitionTypeCodeSequence (0080,0001) VR=SQ VM=1 Surface Scan Acquisition Type Code Sequence
	SurfaceScanAcquisitionTypeCodeSequence = New(0x0080, 0x0001)

	// SurfaceScanModeCodeSequence (0080,0002) VR=SQ VM=1 Surface Scan Mode Code Sequence
	SurfaceScanModeCodeSequence = New(0x0080, 0x0002)

	// RegistrationMethodCodeSequence (0080,0003) VR=SQ VM=1 Registration Method Code Sequence
	RegistrationMethodCodeSequence = New(0x0080, 0x0003)

	// ShotDurationTime (0080,0004) VR=FD VM=1 Shot Duration Time
	ShotDurationTime = New(0x0080, 0x0004)

	// ShotOffsetTime (0080,0005) VR=FD VM=1 Shot Offset Time
	ShotOffsetTime = New(0x0080, 0x0005)

	// SurfacePointPresentationValueData (0080,0006) VR=US VM=1-n Surface Point Presentation Value Data
	SurfacePointPresentationValueData = New(0x0080, 0x0006)

	// SurfacePointColorCIELabValueData (0080,0007) VR=US VM=3-3n Surface Point Color CIELab Value Data
	SurfacePointColorCIELabValueData = New(0x0080, 0x0007)

	// UVMappingSequence (0080,0008) VR=SQ VM=1 UV Mapping Sequence
	UVMappingSequence = New(0x0080, 0x0008)

	// TextureLabel (0080,0009) VR=SH VM=1 Texture Label
	TextureLabel = New(0x0080, 0x0009)

	// UValueData (0080,0010) VR=OF VM=1 U Value Data
	UValueData = New(0x0080, 0x0010)

	// VValueData (0080,0011) VR=OF VM=1 V Value Data
	VValueData = New(0x0080, 0x0011)

	// ReferencedTextureSequence (0080,0012) VR=SQ VM=1 Referenced Texture Sequence
	ReferencedTextureSequence = New(0x0080, 0x0012)

	// ReferencedSurfaceDataSequence (0080,0013) VR=SQ VM=1 Referenced Surface Data Sequence
	ReferencedSurfaceDataSequence = New(0x0080, 0x0013)

	// AssessmentSummary (0082,0001) VR=CS VM=1 Assessment Summary
	AssessmentSummary = New(0x0082, 0x0001)

	// AssessmentSummaryDescription (0082,0003) VR=UT VM=1 Assessment Summary Description
	AssessmentSummaryDescription = New(0x0082, 0x0003)

	// AssessedSOPInstanceSequence (0082,0004) VR=SQ VM=1 Assessed SOP Instance Sequence
	AssessedSOPInstanceSequence = New(0x0082, 0x0004)

	// ReferencedComparisonSOPInstanceSequence (0082,0005) VR=SQ VM=1 Referenced Comparison SOP Instance Sequence
	ReferencedComparisonSOPInstanceSequence = New(0x0082, 0x0005)

	// NumberOfAssessmentObservations (0082,0006) VR=UL VM=1 Number of Assessment Observations
	NumberOfAssessmentObservations = New(0x0082, 0x0006)

	// AssessmentObservationsSequence (0082,0007) VR=SQ VM=1 Assessment Observations Sequence
	AssessmentObservationsSequence = New(0x0082, 0x0007)

	// ObservationSignificance (0082,0008) VR=CS VM=1 Observation Significance
	ObservationSignificance = New(0x0082, 0x0008)

	// ObservationDescription (0082,000A) VR=UT VM=1 Observation Description
	ObservationDescription = New(0x0082, 0x000A)

	// StructuredConstraintObservationSequence (0082,000C) VR=SQ VM=1 Structured Constraint Observation Sequence
	StructuredConstraintObservationSequence = New(0x0082, 0x000C)

	// AssessedAttributeValueSequence (0082,0010) VR=SQ VM=1 Assessed Attribute Value Sequence
	AssessedAttributeValueSequence = New(0x0082, 0x0010)

	// AssessmentSetID (0082,0016) VR=LO VM=1 Assessment Set ID
	AssessmentSetID = New(0x0082, 0x0016)

	// AssessmentRequesterSequence (0082,0017) VR=SQ VM=1 Assessment Requester Sequence
	AssessmentRequesterSequence = New(0x0082, 0x0017)

	// SelectorAttributeName (0082,0018) VR=LO VM=1 Selector Attribute Name
	SelectorAttributeName = New(0x0082, 0x0018)

	// SelectorAttributeKeyword (0082,0019) VR=LO VM=1 Selector Attribute Keyword
	SelectorAttributeKeyword = New(0x0082, 0x0019)

	// AssessmentTypeCodeSequence (0082,0021) VR=SQ VM=1 Assessment Type Code Sequence
	AssessmentTypeCodeSequence = New(0x0082, 0x0021)

	// ObservationBasisCodeSequence (0082,0022) VR=SQ VM=1 Observation Basis Code Sequence
	ObservationBasisCodeSequence = New(0x0082, 0x0022)

	// AssessmentLabel (0082,0023) VR=LO VM=1 Assessment Label
	AssessmentLabel = New(0x0082, 0x0023)

	// ConstraintType (0082,0032) VR=CS VM=1 Constraint Type
	ConstraintType = New(0x0082, 0x0032)

	// SpecificationSelectionGuidance (0082,0033) VR=UT VM=1 Specification Selection Guidance
	SpecificationSelectionGuidance = New(0x0082, 0x0033)

	// ConstraintValueSequence (0082,0034) VR=SQ VM=1 Constraint Value Sequence
	ConstraintValueSequence = New(0x0082, 0x0034)

	// RecommendedDefaultValueSequence (0082,0035) VR=SQ VM=1 Recommended Default Value Sequence
	RecommendedDefaultValueSequence = New(0x0082, 0x0035)

	// ConstraintViolationSignificance (0082,0036) VR=CS VM=1 Constraint Violation Significance
	ConstraintViolationSignificance = New(0x0082, 0x0036)

	// ConstraintViolationCondition (0082,0037) VR=UT VM=1 Constraint Violation Condition
	ConstraintViolationCondition = New(0x0082, 0x0037)

	// ModifiableConstraintFlag (0082,0038) VR=CS VM=1 Modifiable Constraint Flag
	ModifiableConstraintFlag = New(0x0082, 0x0038)

	// StorageMediaFileSetID (0088,0130) VR=SH VM=1 Storage Media File-set ID
	StorageMediaFileSetID = New(0x0088, 0x0130)

	// StorageMediaFileSetUID (0088,0140) VR=UI VM=1 Storage Media File-set UID
	StorageMediaFileSetUID = New(0x0088, 0x0140)

	// IconImageSequence (0088,0200) VR=SQ VM=1 Icon Image Sequence
	IconImageSequence = New(0x0088, 0x0200)

	// TopicTitleRETIRED (0088,0904) VR=LO VM=1 Topic Title (RETIRED)
	TopicTitleRETIRED = New(0x0088, 0x0904)

	// TopicSubjectRETIRED (0088,0906) VR=ST VM=1 Topic Subject (RETIRED)
	TopicSubjectRETIRED = New(0x0088, 0x0906)

	// TopicAuthorRETIRED (0088,0910) VR=LO VM=1 Topic Author (RETIRED)
	TopicAuthorRETIRED = New(0x0088, 0x0910)

	// TopicKeywordsRETIRED (0088,0912) VR=LO VM=1-32 Topic Keywords (RETIRED)
	TopicKeywordsRETIRED = New(0x0088, 0x0912)

	// SOPInstanceStatus (0100,0410) VR=CS VM=1 SOP Instance Status
	SOPInstanceStatus = New(0x0100, 0x0410)

	// SOPAuthorizationDateTime (0100,0420) VR=DT VM=1 SOP Authorization DateTime
	SOPAuthorizationDateTime = New(0x0100, 0x0420)

	// SOPAuthorizationComment (0100,0424) VR=LT VM=1 SOP Authorization Comment
	SOPAuthorizationComment = New(0x0100, 0x0424)

	// AuthorizationEquipmentCertificationNumber (0100,0426) VR=LO VM=1 Authorization Equipment Certification Number
	AuthorizationEquipmentCertificationNumber = New(0x0100, 0x0426)

	// MACIDNumber (0400,0005) VR=US VM=1 MAC ID Number
	MACIDNumber = New(0x0400, 0x0005)

	// MACCalculationTransferSyntaxUID (0400,0010) VR=UI VM=1 MAC Calculation Transfer Syntax UID
	MACCalculationTransferSyntaxUID = New(0x0400, 0x0010)

	// MACAlgorithm (0400,0015) VR=CS VM=1 MAC Algorithm
	MACAlgorithm = New(0x0400, 0x0015)

	// DataElementsSigned (0400,0020) VR=AT VM=1-n Data Elements Signed
	DataElementsSigned = New(0x0400, 0x0020)

	// DigitalSignatureUID (0400,0100) VR=UI VM=1 Digital Signature UID
	DigitalSignatureUID = New(0x0400, 0x0100)

	// DigitalSignatureDateTime (0400,0105) VR=DT VM=1 Digital Signature DateTime
	DigitalSignatureDateTime = New(0x0400, 0x0105)

	// CertificateType (0400,0110) VR=CS VM=1 Certificate Type
	CertificateType = New(0x0400, 0x0110)

	// CertificateOfSigner (0400,0115) VR=OB VM=1 Certificate of Signer
	CertificateOfSigner = New(0x0400, 0x0115)

	// Signature (0400,0120) VR=OB VM=1 Signature
	Signature = New(0x0400, 0x0120)

	// CertifiedTimestampType (0400,0305) VR=CS VM=1 Certified Timestamp Type
	CertifiedTimestampType = New(0x0400, 0x0305)

	// CertifiedTimestamp (0400,0310) VR=OB VM=1 Certified Timestamp
	CertifiedTimestamp = New(0x0400, 0x0310)

	// DigitalSignaturePurposeCodeSequence (0400,0401) VR=SQ VM=1 Digital Signature Purpose Code Sequence
	DigitalSignaturePurposeCodeSequence = New(0x0400, 0x0401)

	// ReferencedDigitalSignatureSequence (0400,0402) VR=SQ VM=1 Referenced Digital Signature Sequence
	ReferencedDigitalSignatureSequence = New(0x0400, 0x0402)

	// ReferencedSOPInstanceMACSequence (0400,0403) VR=SQ VM=1 Referenced SOP Instance MAC Sequence
	ReferencedSOPInstanceMACSequence = New(0x0400, 0x0403)

	// MAC (0400,0404) VR=OB VM=1 MAC
	MAC = New(0x0400, 0x0404)

	// EncryptedAttributesSequence (0400,0500) VR=SQ VM=1 Encrypted Attributes Sequence
	EncryptedAttributesSequence = New(0x0400, 0x0500)

	// EncryptedContentTransferSyntaxUID (0400,0510) VR=UI VM=1 Encrypted Content Transfer Syntax UID
	EncryptedContentTransferSyntaxUID = New(0x0400, 0x0510)

	// EncryptedContent (0400,0520) VR=OB VM=1 Encrypted Content
	EncryptedContent = New(0x0400, 0x0520)

	// ModifiedAttributesSequence (0400,0550) VR=SQ VM=1 Modified Attributes Sequence
	ModifiedAttributesSequence = New(0x0400, 0x0550)

	// NonconformingModifiedAttributesSequence (0400,0551) VR=SQ VM=1 Nonconforming Modified Attributes Sequence
	NonconformingModifiedAttributesSequence = New(0x0400, 0x0551)

	// NonconformingDataElementValue (0400,0552) VR=OB VM=1 Nonconforming Data Element Value
	NonconformingDataElementValue = New(0x0400, 0x0552)

	// OriginalAttributesSequence (0400,0561) VR=SQ VM=1 Original Attributes Sequence
	OriginalAttributesSequence = New(0x0400, 0x0561)

	// AttributeModificationDateTime (0400,0562) VR=DT VM=1 Attribute Modification DateTime
	AttributeModificationDateTime = New(0x0400, 0x0562)

	// ModifyingSystem (0400,0563) VR=LO VM=1 Modifying System
	ModifyingSystem = New(0x0400, 0x0563)

	// SourceOfPreviousValues (0400,0564) VR=LO VM=1 Source of Previous Values
	SourceOfPreviousValues = New(0x0400, 0x0564)

	// ReasonForTheAttributeModification (0400,0565) VR=CS VM=1 Reason for the Attribute Modification
	ReasonForTheAttributeModification = New(0x0400, 0x0565)

	// InstanceOriginStatus (0400,0600) VR=CS VM=1 Instance Origin Status
	InstanceOriginStatus = New(0x0400, 0x0600)

	EscapeTripletRETIRED = New(0x1000, 0x0000)

	RunLengthTripletRETIRED = New(0x1000, 0x0001)

	HuffmanTableSizeRETIRED = New(0x1000, 0x0002)

	HuffmanTableTripletRETIRED = New(0x1000, 0x0003)

	ShiftTableSizeRETIRED = New(0x1000, 0x0004)

	ShiftTableTripletRETIRED = New(0x1000, 0x0005)

	ZonalMapRETIRED = New(0x1010, 0x0000)

	// NumberOfCopies (2000,0010) VR=IS VM=1 Number of Copies
	NumberOfCopies = New(0x2000, 0x0010)

	// PrinterConfigurationSequence (2000,001E) VR=SQ VM=1 Printer Configuration Sequence
	PrinterConfigurationSequence = New(0x2000, 0x001E)

	// PrintPriority (2000,0020) VR=CS VM=1 Print Priority
	PrintPriority = New(0x2000, 0x0020)

	// MediumType (2000,0030) VR=CS VM=1 Medium Type
	MediumType = New(0x2000, 0x0030)

	// FilmDestination (2000,0040) VR=CS VM=1 Film Destination
	FilmDestination = New(0x2000, 0x0040)

	// FilmSessionLabel (2000,0050) VR=LO VM=1 Film Session Label
	FilmSessionLabel = New(0x2000, 0x0050)

	// MemoryAllocation (2000,0060) VR=IS VM=1 Memory Allocation
	MemoryAllocation = New(0x2000, 0x0060)

	// MaximumMemoryAllocation (2000,0061) VR=IS VM=1 Maximum Memory Allocation
	MaximumMemoryAllocation = New(0x2000, 0x0061)

	// ColorImagePrintingFlagRETIRED (2000,0062) VR=CS VM=1 Color Image Printing Flag (RETIRED)
	ColorImagePrintingFlagRETIRED = New(0x2000, 0x0062)

	// CollationFlagRETIRED (2000,0063) VR=CS VM=1 Collation Flag (RETIRED)
	CollationFlagRETIRED = New(0x2000, 0x0063)

	// AnnotationFlagRETIRED (2000,0065) VR=CS VM=1 Annotation Flag (RETIRED)
	AnnotationFlagRETIRED = New(0x2000, 0x0065)

	// ImageOverlayFlagRETIRED (2000,0067) VR=CS VM=1 Image Overlay Flag (RETIRED)
	ImageOverlayFlagRETIRED = New(0x2000, 0x0067)

	// PresentationLUTFlagRETIRED (2000,0069) VR=CS VM=1 Presentation LUT Flag (RETIRED)
	PresentationLUTFlagRETIRED = New(0x2000, 0x0069)

	// ImageBoxPresentationLUTFlagRETIRED (2000,006A) VR=CS VM=1 Image Box Presentation LUT Flag (RETIRED)
	ImageBoxPresentationLUTFlagRETIRED = New(0x2000, 0x006A)

	// MemoryBitDepth (2000,00A0) VR=US VM=1 Memory Bit Depth
	MemoryBitDepth = New(0x2000, 0x00A0)

	// PrintingBitDepth (2000,00A1) VR=US VM=1 Printing Bit Depth
	PrintingBitDepth = New(0x2000, 0x00A1)

	// MediaInstalledSequence (2000,00A2) VR=SQ VM=1 Media Installed Sequence
	MediaInstalledSequence = New(0x2000, 0x00A2)

	// OtherMediaAvailableSequence (2000,00A4) VR=SQ VM=1 Other Media Available Sequence
	OtherMediaAvailableSequence = New(0x2000, 0x00A4)

	// SupportedImageDisplayFormatsSequence (2000,00A8) VR=SQ VM=1 Supported Image Display Formats Sequence
	SupportedImageDisplayFormatsSequence = New(0x2000, 0x00A8)

	// ReferencedFilmBoxSequence (2000,0500) VR=SQ VM=1 Referenced Film Box Sequence
	ReferencedFilmBoxSequence = New(0x2000, 0x0500)

	// ReferencedStoredPrintSequenceRETIRED (2000,0510) VR=SQ VM=1 Referenced Stored Print Sequence (RETIRED)
	ReferencedStoredPrintSequenceRETIRED = New(0x2000, 0x0510)

	// ImageDisplayFormat (2010,0010) VR=ST VM=1 Image Display Format
	ImageDisplayFormat = New(0x2010, 0x0010)

	// AnnotationDisplayFormatID (2010,0030) VR=CS VM=1 Annotation Display Format ID
	AnnotationDisplayFormatID = New(0x2010, 0x0030)

	// FilmOrientation (2010,0040) VR=CS VM=1 Film Orientation
	FilmOrientation = New(0x2010, 0x0040)

	// FilmSizeID (2010,0050) VR=CS VM=1 Film Size ID
	FilmSizeID = New(0x2010, 0x0050)

	// PrinterResolutionID (2010,0052) VR=CS VM=1 Printer Resolution ID
	PrinterResolutionID = New(0x2010, 0x0052)

	// DefaultPrinterResolutionID (2010,0054) VR=CS VM=1 Default Printer Resolution ID
	DefaultPrinterResolutionID = New(0x2010, 0x0054)

	// MagnificationType (2010,0060) VR=CS VM=1 Magnification Type
	MagnificationType = New(0x2010, 0x0060)

	// SmoothingType (2010,0080) VR=CS VM=1 Smoothing Type
	SmoothingType = New(0x2010, 0x0080)

	// DefaultMagnificationType (2010,00A6) VR=CS VM=1 Default Magnification Type
	DefaultMagnificationType = New(0x2010, 0x00A6)

	// OtherMagnificationTypesAvailable (2010,00A7) VR=CS VM=1-n Other Magnification Types Available
	OtherMagnificationTypesAvailable = New(0x2010, 0x00A7)

	// DefaultSmoothingType (2010,00A8) VR=CS VM=1 Default Smoothing Type
	DefaultSmoothingType = New(0x2010, 0x00A8)

	// OtherSmoothingTypesAvailable (2010,00A9) VR=CS VM=1-n Other Smoothing Types Available
	OtherSmoothingTypesAvailable = New(0x2010, 0x00A9)

	// BorderDensity (2010,0100) VR=CS VM=1 Border Density
	BorderDensity = New(0x2010, 0x0100)

	// EmptyImageDensity (2010,0110) VR=CS VM=1 Empty Image Density
	EmptyImageDensity = New(0x2010, 0x0110)

	// MinDensity (2010,0120) VR=US VM=1 Min Density
	MinDensity = New(0x2010, 0x0120)

	// MaxDensity (2010,0130) VR=US VM=1 Max Density
	MaxDensity = New(0x2010, 0x0130)

	// Trim (2010,0140) VR=CS VM=1 Trim
	Trim = New(0x2010, 0x0140)

	// ConfigurationInformation (2010,0150) VR=ST VM=1 Configuration Information
	ConfigurationInformation = New(0x2010, 0x0150)

	// ConfigurationInformationDescription (2010,0152) VR=LT VM=1 Configuration Information Description
	ConfigurationInformationDescription = New(0x2010, 0x0152)

	// MaximumCollatedFilms (2010,0154) VR=IS VM=1 Maximum Collated Films
	MaximumCollatedFilms = New(0x2010, 0x0154)

	// Illumination (2010,015E) VR=US VM=1 Illumination
	Illumination = New(0x2010, 0x015E)

	// ReflectedAmbientLight (2010,0160) VR=US VM=1 Reflected Ambient Light
	ReflectedAmbientLight = New(0x2010, 0x0160)

	// PrinterPixelSpacing (2010,0376) VR=DS VM=2 Printer Pixel Spacing
	PrinterPixelSpacing = New(0x2010, 0x0376)

	// ReferencedFilmSessionSequence (2010,0500) VR=SQ VM=1 Referenced Film Session Sequence
	ReferencedFilmSessionSequence = New(0x2010, 0x0500)

	// ReferencedImageBoxSequence (2010,0510) VR=SQ VM=1 Referenced Image Box Sequence
	ReferencedImageBoxSequence = New(0x2010, 0x0510)

	// ReferencedBasicAnnotationBoxSequence (2010,0520) VR=SQ VM=1 Referenced Basic Annotation Box Sequence
	ReferencedBasicAnnotationBoxSequence = New(0x2010, 0x0520)

	// ImageBoxPosition (2020,0010) VR=US VM=1 Image Box Position
	ImageBoxPosition = New(0x2020, 0x0010)

	// Polarity (2020,0020) VR=CS VM=1 Polarity
	Polarity = New(0x2020, 0x0020)

	// RequestedImageSize (2020,0030) VR=DS VM=1 Requested Image Size
	RequestedImageSize = New(0x2020, 0x0030)

	// RequestedDecimateCropBehavior (2020,0040) VR=CS VM=1 Requested Decimate/Crop Behavior
	RequestedDecimateCropBehavior = New(0x2020, 0x0040)

	// RequestedResolutionID (2020,0050) VR=CS VM=1 Requested Resolution ID
	RequestedResolutionID = New(0x2020, 0x0050)

	// RequestedImageSizeFlag (2020,00A0) VR=CS VM=1 Requested Image Size Flag
	RequestedImageSizeFlag = New(0x2020, 0x00A0)

	// DecimateCropResult (2020,00A2) VR=CS VM=1 Decimate/Crop Result
	DecimateCropResult = New(0x2020, 0x00A2)

	// BasicGrayscaleImageSequence (2020,0110) VR=SQ VM=1 Basic Grayscale Image Sequence
	BasicGrayscaleImageSequence = New(0x2020, 0x0110)

	// BasicColorImageSequence (2020,0111) VR=SQ VM=1 Basic Color Image Sequence
	BasicColorImageSequence = New(0x2020, 0x0111)

	// ReferencedImageOverlayBoxSequenceRETIRED (2020,0130) VR=SQ VM=1 Referenced Image Overlay Box Sequence (RETIRED)
	ReferencedImageOverlayBoxSequenceRETIRED = New(0x2020, 0x0130)

	// ReferencedVOILUTBoxSequenceRETIRED (2020,0140) VR=SQ VM=1 Referenced VOI LUT Box Sequence (RETIRED)
	ReferencedVOILUTBoxSequenceRETIRED = New(0x2020, 0x0140)

	// AnnotationPosition (2030,0010) VR=US VM=1 Annotation Position
	AnnotationPosition = New(0x2030, 0x0010)

	// TextString (2030,0020) VR=LO VM=1 Text String
	TextString = New(0x2030, 0x0020)

	// ReferencedOverlayPlaneSequenceRETIRED (2040,0010) VR=SQ VM=1 Referenced Overlay Plane Sequence (RETIRED)
	ReferencedOverlayPlaneSequenceRETIRED = New(0x2040, 0x0010)

	// ReferencedOverlayPlaneGroupsRETIRED (2040,0011) VR=US VM=1-99 Referenced Overlay Plane Groups (RETIRED)
	ReferencedOverlayPlaneGroupsRETIRED = New(0x2040, 0x0011)

	// OverlayPixelDataSequenceRETIRED (2040,0020) VR=SQ VM=1 Overlay Pixel Data Sequence (RETIRED)
	OverlayPixelDataSequenceRETIRED = New(0x2040, 0x0020)

	// OverlayMagnificationTypeRETIRED (2040,0060) VR=CS VM=1 Overlay Magnification Type (RETIRED)
	OverlayMagnificationTypeRETIRED = New(0x2040, 0x0060)

	// OverlaySmoothingTypeRETIRED (2040,0070) VR=CS VM=1 Overlay Smoothing Type (RETIRED)
	OverlaySmoothingTypeRETIRED = New(0x2040, 0x0070)

	// OverlayOrImageMagnificationRETIRED (2040,0072) VR=CS VM=1 Overlay or Image Magnification (RETIRED)
	OverlayOrImageMagnificationRETIRED = New(0x2040, 0x0072)

	// MagnifyToNumberOfColumnsRETIRED (2040,0074) VR=US VM=1 Magnify to Number of Columns (RETIRED)
	MagnifyToNumberOfColumnsRETIRED = New(0x2040, 0x0074)

	// OverlayForegroundDensityRETIRED (2040,0080) VR=CS VM=1 Overlay Foreground Density (RETIRED)
	OverlayForegroundDensityRETIRED = New(0x2040, 0x0080)

	// OverlayBackgroundDensityRETIRED (2040,0082) VR=CS VM=1 Overlay Background Density (RETIRED)
	OverlayBackgroundDensityRETIRED = New(0x2040, 0x0082)

	// OverlayModeRETIRED (2040,0090) VR=CS VM=1 Overlay Mode (RETIRED)
	OverlayModeRETIRED = New(0x2040, 0x0090)

	// ThresholdDensityRETIRED (2040,0100) VR=CS VM=1 Threshold Density (RETIRED)
	ThresholdDensityRETIRED = New(0x2040, 0x0100)

	// ReferencedImageBoxSequenceRetiredRETIRED (2040,0500) VR=SQ VM=1 Referenced Image Box Sequence (Retired) (RETIRED)
	ReferencedImageBoxSequenceRetiredRETIRED = New(0x2040, 0x0500)

	// PresentationLUTSequence (2050,0010) VR=SQ VM=1 Presentation LUT Sequence
	PresentationLUTSequence = New(0x2050, 0x0010)

	// PresentationLUTShape (2050,0020) VR=CS VM=1 Presentation LUT Shape
	PresentationLUTShape = New(0x2050, 0x0020)

	// ReferencedPresentationLUTSequence (2050,0500) VR=SQ VM=1 Referenced Presentation LUT Sequence
	ReferencedPresentationLUTSequence = New(0x2050, 0x0500)

	// PrintJobIDRETIRED (2100,0010) VR=SH VM=1 Print Job ID (RETIRED)
	PrintJobIDRETIRED = New(0x2100, 0x0010)

	// ExecutionStatus (2100,0020) VR=CS VM=1 Execution Status
	ExecutionStatus = New(0x2100, 0x0020)

	// ExecutionStatusInfo (2100,0030) VR=CS VM=1 Execution Status Info
	ExecutionStatusInfo = New(0x2100, 0x0030)

	// CreationDate (2100,0040) VR=DA VM=1 Creation Date
	CreationDate = New(0x2100, 0x0040)

	// CreationTime (2100,0050) VR=TM VM=1 Creation Time
	CreationTime = New(0x2100, 0x0050)

	// Originator (2100,0070) VR=AE VM=1 Originator
	Originator = New(0x2100, 0x0070)

	// DestinationAE (2100,0140) VR=AE VM=1 Destination AE
	DestinationAE = New(0x2100, 0x0140)

	// OwnerID (2100,0160) VR=SH VM=1 Owner ID
	OwnerID = New(0x2100, 0x0160)

	// NumberOfFilms (2100,0170) VR=IS VM=1 Number of Films
	NumberOfFilms = New(0x2100, 0x0170)

	// ReferencedPrintJobSequencePullStoredPrintRETIRED (2100,0500) VR=SQ VM=1 Referenced Print Job Sequence (Pull Stored Print) (RETIRED)
	ReferencedPrintJobSequencePullStoredPrintRETIRED = New(0x2100, 0x0500)

	// PrinterStatus (2110,0010) VR=CS VM=1 Printer Status
	PrinterStatus = New(0x2110, 0x0010)

	// PrinterStatusInfo (2110,0020) VR=CS VM=1 Printer Status Info
	PrinterStatusInfo = New(0x2110, 0x0020)

	// PrinterName (2110,0030) VR=LO VM=1 Printer Name
	PrinterName = New(0x2110, 0x0030)

	// PrintQueueIDRETIRED (2110,0099) VR=SH VM=1 Print Queue ID (RETIRED)
	PrintQueueIDRETIRED = New(0x2110, 0x0099)

	// QueueStatusRETIRED (2120,0010) VR=CS VM=1 Queue Status (RETIRED)
	QueueStatusRETIRED = New(0x2120, 0x0010)

	// PrintJobDescriptionSequenceRETIRED (2120,0050) VR=SQ VM=1 Print Job Description Sequence (RETIRED)
	PrintJobDescriptionSequenceRETIRED = New(0x2120, 0x0050)

	// ReferencedPrintJobSequenceRETIRED (2120,0070) VR=SQ VM=1 Referenced Print Job Sequence (RETIRED)
	ReferencedPrintJobSequenceRETIRED = New(0x2120, 0x0070)

	// PrintManagementCapabilitiesSequenceRETIRED (2130,0010) VR=SQ VM=1 Print Management Capabilities Sequence (RETIRED)
	PrintManagementCapabilitiesSequenceRETIRED = New(0x2130, 0x0010)

	// PrinterCharacteristicsSequenceRETIRED (2130,0015) VR=SQ VM=1 Printer Characteristics Sequence (RETIRED)
	PrinterCharacteristicsSequenceRETIRED = New(0x2130, 0x0015)

	// FilmBoxContentSequenceRETIRED (2130,0030) VR=SQ VM=1 Film Box Content Sequence (RETIRED)
	FilmBoxContentSequenceRETIRED = New(0x2130, 0x0030)

	// ImageBoxContentSequenceRETIRED (2130,0040) VR=SQ VM=1 Image Box Content Sequence (RETIRED)
	ImageBoxContentSequenceRETIRED = New(0x2130, 0x0040)

	// AnnotationContentSequenceRETIRED (2130,0050) VR=SQ VM=1 Annotation Content Sequence (RETIRED)
	AnnotationContentSequenceRETIRED = New(0x2130, 0x0050)

	// ImageOverlayBoxContentSequenceRETIRED (2130,0060) VR=SQ VM=1 Image Overlay Box Content Sequence (RETIRED)
	ImageOverlayBoxContentSequenceRETIRED = New(0x2130, 0x0060)

	// PresentationLUTContentSequenceRETIRED (2130,0080) VR=SQ VM=1 Presentation LUT Content Sequence (RETIRED)
	PresentationLUTContentSequenceRETIRED = New(0x2130, 0x0080)

	// ProposedStudySequence (2130,00A0) VR=SQ VM=1 Proposed Study Sequence
	ProposedStudySequence = New(0x2130, 0x00A0)

	// OriginalImageSequence (2130,00C0) VR=SQ VM=1 Original Image Sequence
	OriginalImageSequence = New(0x2130, 0x00C0)

	// LabelUsingInformationExtractedFromInstances (2200,0001) VR=CS VM=1 Label Using Information Extracted From Instances
	LabelUsingInformationExtractedFromInstances = New(0x2200, 0x0001)

	// LabelText (2200,0002) VR=UT VM=1 Label Text
	LabelText = New(0x2200, 0x0002)

	// LabelStyleSelection (2200,0003) VR=CS VM=1 Label Style Selection
	LabelStyleSelection = New(0x2200, 0x0003)

	// MediaDisposition (2200,0004) VR=LT VM=1 Media Disposition
	MediaDisposition = New(0x2200, 0x0004)

	// BarcodeValue (2200,0005) VR=LT VM=1 Barcode Value
	BarcodeValue = New(0x2200, 0x0005)

	// BarcodeSymbology (2200,0006) VR=CS VM=1 Barcode Symbology
	BarcodeSymbology = New(0x2200, 0x0006)

	// AllowMediaSplitting (2200,0007) VR=CS VM=1 Allow Media Splitting
	AllowMediaSplitting = New(0x2200, 0x0007)

	// IncludeNonDICOMObjects (2200,0008) VR=CS VM=1 Include Non-DICOM Objects
	IncludeNonDICOMObjects = New(0x2200, 0x0008)

	// IncludeDisplayApplication (2200,0009) VR=CS VM=1 Include Display Application
	IncludeDisplayApplication = New(0x2200, 0x0009)

	// PreserveCompositeInstancesAfterMediaCreation (2200,000A) VR=CS VM=1 Preserve Composite Instances After Media Creation
	PreserveCompositeInstancesAfterMediaCreation = New(0x2200, 0x000A)

	// TotalNumberOfPiecesOfMediaCreated (2200,000B) VR=US VM=1 Total Number of Pieces of Media Created
	TotalNumberOfPiecesOfMediaCreated = New(0x2200, 0x000B)

	// RequestedMediaApplicationProfile (2200,000C) VR=LO VM=1 Requested Media Application Profile
	RequestedMediaApplicationProfile = New(0x2200, 0x000C)

	// ReferencedStorageMediaSequence (2200,000D) VR=SQ VM=1 Referenced Storage Media Sequence
	ReferencedStorageMediaSequence = New(0x2200, 0x000D)

	// FailureAttributes (2200,000E) VR=AT VM=1-n Failure Attributes
	FailureAttributes = New(0x2200, 0x000E)

	// AllowLossyCompression (2200,000F) VR=CS VM=1 Allow Lossy Compression
	AllowLossyCompression = New(0x2200, 0x000F)

	// RequestPriority (2200,0020) VR=CS VM=1 Request Priority
	RequestPriority = New(0x2200, 0x0020)

	// RTImageLabel (3002,0002) VR=SH VM=1 RT Image Label
	RTImageLabel = New(0x3002, 0x0002)

	// RTImageName (3002,0003) VR=LO VM=1 RT Image Name
	RTImageName = New(0x3002, 0x0003)

	// RTImageDescription (3002,0004) VR=ST VM=1 RT Image Description
	RTImageDescription = New(0x3002, 0x0004)

	// ReportedValuesOrigin (3002,000A) VR=CS VM=1 Reported Values Origin
	ReportedValuesOrigin = New(0x3002, 0x000A)

	// RTImagePlane (3002,000C) VR=CS VM=1 RT Image Plane
	RTImagePlane = New(0x3002, 0x000C)

	// XRayImageReceptorTranslation (3002,000D) VR=DS VM=3 X-Ray Image Receptor Translation
	XRayImageReceptorTranslation = New(0x3002, 0x000D)

	// XRayImageReceptorAngle (3002,000E) VR=DS VM=1 X-Ray Image Receptor Angle
	XRayImageReceptorAngle = New(0x3002, 0x000E)

	// RTImageOrientation (3002,0010) VR=DS VM=6 RT Image Orientation
	RTImageOrientation = New(0x3002, 0x0010)

	// ImagePlanePixelSpacing (3002,0011) VR=DS VM=2 Image Plane Pixel Spacing
	ImagePlanePixelSpacing = New(0x3002, 0x0011)

	// RTImagePosition (3002,0012) VR=DS VM=2 RT Image Position
	RTImagePosition = New(0x3002, 0x0012)

	// RadiationMachineName (3002,0020) VR=SH VM=1 Radiation Machine Name
	RadiationMachineName = New(0x3002, 0x0020)

	// RadiationMachineSAD (3002,0022) VR=DS VM=1 Radiation Machine SAD
	RadiationMachineSAD = New(0x3002, 0x0022)

	// RadiationMachineSSD (3002,0024) VR=DS VM=1 Radiation Machine SSD
	RadiationMachineSSD = New(0x3002, 0x0024)

	// RTImageSID (3002,0026) VR=DS VM=1 RT Image SID
	RTImageSID = New(0x3002, 0x0026)

	// SourceToReferenceObjectDistance (3002,0028) VR=DS VM=1 Source to Reference Object Distance
	SourceToReferenceObjectDistance = New(0x3002, 0x0028)

	// FractionNumber (3002,0029) VR=IS VM=1 Fraction Number
	FractionNumber = New(0x3002, 0x0029)

	// ExposureSequence (3002,0030) VR=SQ VM=1 Exposure Sequence
	ExposureSequence = New(0x3002, 0x0030)

	// MetersetExposure (3002,0032) VR=DS VM=1 Meterset Exposure
	MetersetExposure = New(0x3002, 0x0032)

	// DiaphragmPosition (3002,0034) VR=DS VM=4 Diaphragm Position
	DiaphragmPosition = New(0x3002, 0x0034)

	// FluenceMapSequence (3002,0040) VR=SQ VM=1 Fluence Map Sequence
	FluenceMapSequence = New(0x3002, 0x0040)

	// FluenceDataSource (3002,0041) VR=CS VM=1 Fluence Data Source
	FluenceDataSource = New(0x3002, 0x0041)

	// FluenceDataScale (3002,0042) VR=DS VM=1 Fluence Data Scale
	FluenceDataScale = New(0x3002, 0x0042)

	// PrimaryFluenceModeSequence (3002,0050) VR=SQ VM=1 Primary Fluence Mode Sequence
	PrimaryFluenceModeSequence = New(0x3002, 0x0050)

	// FluenceMode (3002,0051) VR=CS VM=1 Fluence Mode
	FluenceMode = New(0x3002, 0x0051)

	// FluenceModeID (3002,0052) VR=SH VM=1 Fluence Mode ID
	FluenceModeID = New(0x3002, 0x0052)

	// SelectedFrameNumber (3002,0100) VR=IS VM=1 Selected Frame Number
	SelectedFrameNumber = New(0x3002, 0x0100)

	// SelectedFrameFunctionalGroupsSequence (3002,0101) VR=SQ VM=1 Selected Frame Functional Groups Sequence
	SelectedFrameFunctionalGroupsSequence = New(0x3002, 0x0101)

	// RTImageFrameGeneralContentSequence (3002,0102) VR=SQ VM=1 RT Image Frame General Content Sequence
	RTImageFrameGeneralContentSequence = New(0x3002, 0x0102)

	// RTImageFrameContextSequence (3002,0103) VR=SQ VM=1 RT Image Frame Context Sequence
	RTImageFrameContextSequence = New(0x3002, 0x0103)

	// RTImageScopeSequence (3002,0104) VR=SQ VM=1 RT Image Scope Sequence
	RTImageScopeSequence = New(0x3002, 0x0104)

	// BeamModifierCoordinatesPresenceFlag (3002,0105) VR=CS VM=1 Beam Modifier Coordinates Presence Flag
	BeamModifierCoordinatesPresenceFlag = New(0x3002, 0x0105)

	// StartCumulativeMeterset (3002,0106) VR=FD VM=1 Start Cumulative Meterset
	StartCumulativeMeterset = New(0x3002, 0x0106)

	// StopCumulativeMeterset (3002,0107) VR=FD VM=1 Stop Cumulative Meterset
	StopCumulativeMeterset = New(0x3002, 0x0107)

	// RTAcquisitionPatientPositionSequence (3002,0108) VR=SQ VM=1 RT Acquisition Patient Position Sequence
	RTAcquisitionPatientPositionSequence = New(0x3002, 0x0108)

	// RTImageFrameImagingDevicePositionSequence (3002,0109) VR=SQ VM=1 RT Image Frame Imaging Device Position Sequence
	RTImageFrameImagingDevicePositionSequence = New(0x3002, 0x0109)

	// RTImageFramekVRadiationAcquisitionSequence (3002,010A) VR=SQ VM=1 RT Image Frame kV Radiation Acquisition Sequence
	RTImageFramekVRadiationAcquisitionSequence = New(0x3002, 0x010A)

	// RTImageFrameMVRadiationAcquisitionSequence (3002,010B) VR=SQ VM=1 RT Image Frame MV Radiation Acquisition Sequence
	RTImageFrameMVRadiationAcquisitionSequence = New(0x3002, 0x010B)

	// RTImageFrameRadiationAcquisitionSequence (3002,010C) VR=SQ VM=1 RT Image Frame Radiation Acquisition Sequence
	RTImageFrameRadiationAcquisitionSequence = New(0x3002, 0x010C)

	// ImagingSourcePositionSequence (3002,010D) VR=SQ VM=1 Imaging Source Position Sequence
	ImagingSourcePositionSequence = New(0x3002, 0x010D)

	// ImageReceptorPositionSequence (3002,010E) VR=SQ VM=1 Image Receptor Position Sequence
	ImageReceptorPositionSequence = New(0x3002, 0x010E)

	// DevicePositionToEquipmentMappingMatrix (3002,010F) VR=FD VM=16 Device Position to Equipment Mapping Matrix
	DevicePositionToEquipmentMappingMatrix = New(0x3002, 0x010F)

	// DevicePositionParameterSequence (3002,0110) VR=SQ VM=1 Device Position Parameter Sequence
	DevicePositionParameterSequence = New(0x3002, 0x0110)

	// ImagingSourceLocationSpecificationType (3002,0111) VR=CS VM=1 Imaging Source Location Specification Type
	ImagingSourceLocationSpecificationType = New(0x3002, 0x0111)

	// ImagingDeviceLocationMatrixSequence (3002,0112) VR=SQ VM=1 Imaging Device Location Matrix Sequence
	ImagingDeviceLocationMatrixSequence = New(0x3002, 0x0112)

	// ImagingDeviceLocationParameterSequence (3002,0113) VR=SQ VM=1 Imaging Device Location Parameter Sequence
	ImagingDeviceLocationParameterSequence = New(0x3002, 0x0113)

	// ImagingApertureSequence (3002,0114) VR=SQ VM=1 Imaging Aperture Sequence
	ImagingApertureSequence = New(0x3002, 0x0114)

	// ImagingApertureSpecificationType (3002,0115) VR=CS VM=1 Imaging Aperture Specification Type
	ImagingApertureSpecificationType = New(0x3002, 0x0115)

	// NumberOfAcquisitionDevices (3002,0116) VR=US VM=1 Number of Acquisition Devices
	NumberOfAcquisitionDevices = New(0x3002, 0x0116)

	// AcquisitionDeviceSequence (3002,0117) VR=SQ VM=1 Acquisition Device Sequence
	AcquisitionDeviceSequence = New(0x3002, 0x0117)

	// AcquisitionTaskSequence (3002,0118) VR=SQ VM=1 Acquisition Task Sequence
	AcquisitionTaskSequence = New(0x3002, 0x0118)

	// AcquisitionTaskWorkitemCodeSequence (3002,0119) VR=SQ VM=1 Acquisition Task Workitem Code Sequence
	AcquisitionTaskWorkitemCodeSequence = New(0x3002, 0x0119)

	// AcquisitionSubtaskSequence (3002,011A) VR=SQ VM=1 Acquisition Subtask Sequence
	AcquisitionSubtaskSequence = New(0x3002, 0x011A)

	// SubtaskWorkitemCodeSequence (3002,011B) VR=SQ VM=1 Subtask Workitem Code Sequence
	SubtaskWorkitemCodeSequence = New(0x3002, 0x011B)

	// AcquisitionTaskIndex (3002,011C) VR=US VM=1 Acquisition Task Index
	AcquisitionTaskIndex = New(0x3002, 0x011C)

	// AcquisitionSubtaskIndex (3002,011D) VR=US VM=1 Acquisition Subtask Index
	AcquisitionSubtaskIndex = New(0x3002, 0x011D)

	// ReferencedBaselineParametersRTRadiationInstanceSequence (3002,011E) VR=SQ VM=1 Referenced Baseline Parameters RT Radiation Instance Sequence
	ReferencedBaselineParametersRTRadiationInstanceSequence = New(0x3002, 0x011E)

	// PositionAcquisitionTemplateIdentificationSequence (3002,011F) VR=SQ VM=1 Position Acquisition Template Identification Sequence
	PositionAcquisitionTemplateIdentificationSequence = New(0x3002, 0x011F)

	// PositionAcquisitionTemplateID (3002,0120) VR=ST VM=1 Position Acquisition Template ID
	PositionAcquisitionTemplateID = New(0x3002, 0x0120)

	// PositionAcquisitionTemplateName (3002,0121) VR=LO VM=1 Position Acquisition Template Name
	PositionAcquisitionTemplateName = New(0x3002, 0x0121)

	// PositionAcquisitionTemplateCodeSequence (3002,0122) VR=SQ VM=1 Position Acquisition Template Code Sequence
	PositionAcquisitionTemplateCodeSequence = New(0x3002, 0x0122)

	// PositionAcquisitionTemplateDescription (3002,0123) VR=LT VM=1 Position Acquisition Template Description
	PositionAcquisitionTemplateDescription = New(0x3002, 0x0123)

	// AcquisitionTaskApplicabilitySequence (3002,0124) VR=SQ VM=1 Acquisition Task Applicability Sequence
	AcquisitionTaskApplicabilitySequence = New(0x3002, 0x0124)

	// ProjectionImagingAcquisitionParameterSequence (3002,0125) VR=SQ VM=1 Projection Imaging Acquisition Parameter Sequence
	ProjectionImagingAcquisitionParameterSequence = New(0x3002, 0x0125)

	// CTImagingAcquisitionParameterSequence (3002,0126) VR=SQ VM=1 CT Imaging Acquisition Parameter Sequence
	CTImagingAcquisitionParameterSequence = New(0x3002, 0x0126)

	// KVImagingGenerationParametersSequence (3002,0127) VR=SQ VM=1 KV Imaging Generation Parameters Sequence
	KVImagingGenerationParametersSequence = New(0x3002, 0x0127)

	// MVImagingGenerationParametersSequence (3002,0128) VR=SQ VM=1 MV Imaging Generation Parameters Sequence
	MVImagingGenerationParametersSequence = New(0x3002, 0x0128)

	// AcquisitionSignalType (3002,0129) VR=CS VM=1 Acquisition Signal Type
	AcquisitionSignalType = New(0x3002, 0x0129)

	// AcquisitionMethod (3002,012A) VR=CS VM=1 Acquisition Method
	AcquisitionMethod = New(0x3002, 0x012A)

	// ScanStartPositionSequence (3002,012B) VR=SQ VM=1 Scan Start Position Sequence
	ScanStartPositionSequence = New(0x3002, 0x012B)

	// ScanStopPositionSequence (3002,012C) VR=SQ VM=1 Scan Stop Position Sequence
	ScanStopPositionSequence = New(0x3002, 0x012C)

	// ImagingSourceToBeamModifierDefinitionPlaneDistance (3002,012D) VR=FD VM=1 Imaging Source to Beam Modifier Definition Plane Distance
	ImagingSourceToBeamModifierDefinitionPlaneDistance = New(0x3002, 0x012D)

	// ScanArcType (3002,012E) VR=CS VM=1 Scan Arc Type
	ScanArcType = New(0x3002, 0x012E)

	// DetectorPositioningType (3002,012F) VR=CS VM=1 Detector Positioning Type
	DetectorPositioningType = New(0x3002, 0x012F)

	// AdditionalRTAccessoryDeviceSequence (3002,0130) VR=SQ VM=1 Additional RT Accessory Device Sequence
	AdditionalRTAccessoryDeviceSequence = New(0x3002, 0x0130)

	// DeviceSpecificAcquisitionParameterSequence (3002,0131) VR=SQ VM=1 Device-Specific Acquisition Parameter Sequence
	DeviceSpecificAcquisitionParameterSequence = New(0x3002, 0x0131)

	// ReferencedPositionReferenceInstanceSequence (3002,0132) VR=SQ VM=1 Referenced Position Reference Instance Sequence
	ReferencedPositionReferenceInstanceSequence = New(0x3002, 0x0132)

	// EnergyDerivationCodeSequence (3002,0133) VR=SQ VM=1 Energy Derivation Code Sequence
	EnergyDerivationCodeSequence = New(0x3002, 0x0133)

	// MaximumCumulativeMetersetExposure (3002,0134) VR=FD VM=1 Maximum Cumulative Meterset Exposure
	MaximumCumulativeMetersetExposure = New(0x3002, 0x0134)

	// AcquisitionInitiationSequence (3002,0135) VR=SQ VM=1 Acquisition Initiation Sequence
	AcquisitionInitiationSequence = New(0x3002, 0x0135)

	// RTConeBeamImagingGeometrySequence (3002,0136) VR=SQ VM=1 RT Cone-Beam Imaging Geometry Sequence
	RTConeBeamImagingGeometrySequence = New(0x3002, 0x0136)

	// DVHType (3004,0001) VR=CS VM=1 DVH Type
	DVHType = New(0x3004, 0x0001)

	// DoseUnits (3004,0002) VR=CS VM=1 Dose Units
	DoseUnits = New(0x3004, 0x0002)

	// DoseType (3004,0004) VR=CS VM=1 Dose Type
	DoseType = New(0x3004, 0x0004)

	// SpatialTransformOfDose (3004,0005) VR=CS VM=1 Spatial Transform of Dose
	SpatialTransformOfDose = New(0x3004, 0x0005)

	// DoseComment (3004,0006) VR=LO VM=1 Dose Comment
	DoseComment = New(0x3004, 0x0006)

	// NormalizationPoint (3004,0008) VR=DS VM=3 Normalization Point
	NormalizationPoint = New(0x3004, 0x0008)

	// DoseSummationType (3004,000A) VR=CS VM=1 Dose Summation Type
	DoseSummationType = New(0x3004, 0x000A)

	// GridFrameOffsetVector (3004,000C) VR=DS VM=2-n Grid Frame Offset Vector
	GridFrameOffsetVector = New(0x3004, 0x000C)

	// DoseGridScaling (3004,000E) VR=DS VM=1 Dose Grid Scaling
	DoseGridScaling = New(0x3004, 0x000E)

	// RTDoseROISequenceRETIRED (3004,0010) VR=SQ VM=1 RT Dose ROI Sequence (RETIRED)
	RTDoseROISequenceRETIRED = New(0x3004, 0x0010)

	// DoseValue (3004,0012) VR=DS VM=1 Dose Value
	DoseValue = New(0x3004, 0x0012)

	// TissueHeterogeneityCorrection (3004,0014) VR=CS VM=1-3 Tissue Heterogeneity Correction
	TissueHeterogeneityCorrection = New(0x3004, 0x0014)

	// RecommendedIsodoseLevelSequence (3004,0016) VR=SQ VM=1 Recommended Isodose Level Sequence
	RecommendedIsodoseLevelSequence = New(0x3004, 0x0016)

	// DVHNormalizationPoint (3004,0040) VR=DS VM=3 DVH Normalization Point
	DVHNormalizationPoint = New(0x3004, 0x0040)

	// DVHNormalizationDoseValue (3004,0042) VR=DS VM=1 DVH Normalization Dose Value
	DVHNormalizationDoseValue = New(0x3004, 0x0042)

	// DVHSequence (3004,0050) VR=SQ VM=1 DVH Sequence
	DVHSequence = New(0x3004, 0x0050)

	// DVHDoseScaling (3004,0052) VR=DS VM=1 DVH Dose Scaling
	DVHDoseScaling = New(0x3004, 0x0052)

	// DVHVolumeUnits (3004,0054) VR=CS VM=1 DVH Volume Units
	DVHVolumeUnits = New(0x3004, 0x0054)

	// DVHNumberOfBins (3004,0056) VR=IS VM=1 DVH Number of Bins
	DVHNumberOfBins = New(0x3004, 0x0056)

	// DVHData (3004,0058) VR=DS VM=2-2n DVH Data
	DVHData = New(0x3004, 0x0058)

	// DVHReferencedROISequence (3004,0060) VR=SQ VM=1 DVH Referenced ROI Sequence
	DVHReferencedROISequence = New(0x3004, 0x0060)

	// DVHROIContributionType (3004,0062) VR=CS VM=1 DVH ROI Contribution Type
	DVHROIContributionType = New(0x3004, 0x0062)

	// DVHMinimumDose (3004,0070) VR=DS VM=1 DVH Minimum Dose
	DVHMinimumDose = New(0x3004, 0x0070)

	// DVHMaximumDose (3004,0072) VR=DS VM=1 DVH Maximum Dose
	DVHMaximumDose = New(0x3004, 0x0072)

	// DVHMeanDose (3004,0074) VR=DS VM=1 DVH Mean Dose
	DVHMeanDose = New(0x3004, 0x0074)

	// DoseCalculationModelSequence (3004,0080) VR=SQ VM=1 Dose Calculation Model Sequence
	DoseCalculationModelSequence = New(0x3004, 0x0080)

	// DoseCalculationAlgorithmSequence (3004,0081) VR=SQ VM=1 Dose Calculation Algorithm Sequence
	DoseCalculationAlgorithmSequence = New(0x3004, 0x0081)

	// CommissioningStatus (3004,0082) VR=CS VM=1 Commissioning Status
	CommissioningStatus = New(0x3004, 0x0082)

	// DoseCalculationModelParameterSequence (3004,0083) VR=SQ VM=1 Dose Calculation Model Parameter Sequence
	DoseCalculationModelParameterSequence = New(0x3004, 0x0083)

	// DoseDepositionCalculationMedium (3004,0084) VR=CS VM=1 Dose Deposition Calculation Medium
	DoseDepositionCalculationMedium = New(0x3004, 0x0084)

	// StructureSetLabel (3006,0002) VR=SH VM=1 Structure Set Label
	StructureSetLabel = New(0x3006, 0x0002)

	// StructureSetName (3006,0004) VR=LO VM=1 Structure Set Name
	StructureSetName = New(0x3006, 0x0004)

	// StructureSetDescription (3006,0006) VR=ST VM=1 Structure Set Description
	StructureSetDescription = New(0x3006, 0x0006)

	// StructureSetDate (3006,0008) VR=DA VM=1 Structure Set Date
	StructureSetDate = New(0x3006, 0x0008)

	// StructureSetTime (3006,0009) VR=TM VM=1 Structure Set Time
	StructureSetTime = New(0x3006, 0x0009)

	// ReferencedFrameOfReferenceSequence (3006,0010) VR=SQ VM=1 Referenced Frame of Reference Sequence
	ReferencedFrameOfReferenceSequence = New(0x3006, 0x0010)

	// RTReferencedStudySequence (3006,0012) VR=SQ VM=1 RT Referenced Study Sequence
	RTReferencedStudySequence = New(0x3006, 0x0012)

	// RTReferencedSeriesSequence (3006,0014) VR=SQ VM=1 RT Referenced Series Sequence
	RTReferencedSeriesSequence = New(0x3006, 0x0014)

	// ContourImageSequence (3006,0016) VR=SQ VM=1 Contour Image Sequence
	ContourImageSequence = New(0x3006, 0x0016)

	// PredecessorStructureSetSequence (3006,0018) VR=SQ VM=1 Predecessor Structure Set Sequence
	PredecessorStructureSetSequence = New(0x3006, 0x0018)

	// StructureSetROISequence (3006,0020) VR=SQ VM=1 Structure Set ROI Sequence
	StructureSetROISequence = New(0x3006, 0x0020)

	// ROINumber (3006,0022) VR=IS VM=1 ROI Number
	ROINumber = New(0x3006, 0x0022)

	// ReferencedFrameOfReferenceUID (3006,0024) VR=UI VM=1 Referenced Frame of Reference UID
	ReferencedFrameOfReferenceUID = New(0x3006, 0x0024)

	// ROIName (3006,0026) VR=LO VM=1 ROI Name
	ROIName = New(0x3006, 0x0026)

	// ROIDescription (3006,0028) VR=ST VM=1 ROI Description
	ROIDescription = New(0x3006, 0x0028)

	// ROIDisplayColor (3006,002A) VR=IS VM=3 ROI Display Color
	ROIDisplayColor = New(0x3006, 0x002A)

	// ROIVolume (3006,002C) VR=DS VM=1 ROI Volume
	ROIVolume = New(0x3006, 0x002C)

	// ROIDateTime (3006,002D) VR=DT VM=1 ROI DateTime
	ROIDateTime = New(0x3006, 0x002D)

	// ROIObservationDateTime (3006,002E) VR=DT VM=1 ROI Observation DateTime
	ROIObservationDateTime = New(0x3006, 0x002E)

	// RTRelatedROISequence (3006,0030) VR=SQ VM=1 RT Related ROI Sequence
	RTRelatedROISequence = New(0x3006, 0x0030)

	// RTROIRelationship (3006,0033) VR=CS VM=1 RT ROI Relationship
	RTROIRelationship = New(0x3006, 0x0033)

	// ROIGenerationAlgorithm (3006,0036) VR=CS VM=1 ROI Generation Algorithm
	ROIGenerationAlgorithm = New(0x3006, 0x0036)

	// ROIDerivationAlgorithmIdentificationSequence (3006,0037) VR=SQ VM=1 ROI Derivation Algorithm Identification Sequence
	ROIDerivationAlgorithmIdentificationSequence = New(0x3006, 0x0037)

	// ROIGenerationDescription (3006,0038) VR=LO VM=1 ROI Generation Description
	ROIGenerationDescription = New(0x3006, 0x0038)

	// ROIContourSequence (3006,0039) VR=SQ VM=1 ROI Contour Sequence
	ROIContourSequence = New(0x3006, 0x0039)

	// ContourSequence (3006,0040) VR=SQ VM=1 Contour Sequence
	ContourSequence = New(0x3006, 0x0040)

	// ContourGeometricType (3006,0042) VR=CS VM=1 Contour Geometric Type
	ContourGeometricType = New(0x3006, 0x0042)

	// ContourSlabThicknessRETIRED (3006,0044) VR=DS VM=1 Contour Slab Thickness (RETIRED)
	ContourSlabThicknessRETIRED = New(0x3006, 0x0044)

	// ContourOffsetVectorRETIRED (3006,0045) VR=DS VM=3 Contour Offset Vector (RETIRED)
	ContourOffsetVectorRETIRED = New(0x3006, 0x0045)

	// NumberOfContourPoints (3006,0046) VR=IS VM=1 Number of Contour Points
	NumberOfContourPoints = New(0x3006, 0x0046)

	// ContourNumber (3006,0048) VR=IS VM=1 Contour Number
	ContourNumber = New(0x3006, 0x0048)

	// AttachedContoursRETIRED (3006,0049) VR=IS VM=1-n Attached Contours (RETIRED)
	AttachedContoursRETIRED = New(0x3006, 0x0049)

	// SourcePixelPlanesCharacteristicsSequence (3006,004A) VR=SQ VM=1 Source Pixel Planes Characteristics Sequence
	SourcePixelPlanesCharacteristicsSequence = New(0x3006, 0x004A)

	// SourceSeriesSequence (3006,004B) VR=SQ VM=1 Source Series Sequence
	SourceSeriesSequence = New(0x3006, 0x004B)

	// SourceSeriesInformationSequence (3006,004C) VR=SQ VM=1 Source Series Information Sequence
	SourceSeriesInformationSequence = New(0x3006, 0x004C)

	// ROICreatorSequence (3006,004D) VR=SQ VM=1 ROI Creator Sequence
	ROICreatorSequence = New(0x3006, 0x004D)

	// ROIInterpreterSequence (3006,004E) VR=SQ VM=1 ROI Interpreter Sequence
	ROIInterpreterSequence = New(0x3006, 0x004E)

	// ROIObservationContextCodeSequence (3006,004F) VR=SQ VM=1 ROI Observation Context Code Sequence
	ROIObservationContextCodeSequence = New(0x3006, 0x004F)

	// ContourData (3006,0050) VR=DS VM=3-3n Contour Data
	ContourData = New(0x3006, 0x0050)

	// RTROIObservationsSequence (3006,0080) VR=SQ VM=1 RT ROI Observations Sequence
	RTROIObservationsSequence = New(0x3006, 0x0080)

	// ObservationNumber (3006,0082) VR=IS VM=1 Observation Number
	ObservationNumber = New(0x3006, 0x0082)

	// ReferencedROINumber (3006,0084) VR=IS VM=1 Referenced ROI Number
	ReferencedROINumber = New(0x3006, 0x0084)

	// ROIObservationLabelRETIRED (3006,0085) VR=SH VM=1 ROI Observation Label (RETIRED)
	ROIObservationLabelRETIRED = New(0x3006, 0x0085)

	// RTROIIdentificationCodeSequence (3006,0086) VR=SQ VM=1 RT ROI Identification Code Sequence
	RTROIIdentificationCodeSequence = New(0x3006, 0x0086)

	// ROIObservationDescriptionRETIRED (3006,0088) VR=ST VM=1 ROI Observation Description (RETIRED)
	ROIObservationDescriptionRETIRED = New(0x3006, 0x0088)

	// RelatedRTROIObservationsSequence (3006,00A0) VR=SQ VM=1 Related RT ROI Observations Sequence
	RelatedRTROIObservationsSequence = New(0x3006, 0x00A0)

	// RTROIInterpretedType (3006,00A4) VR=CS VM=1 RT ROI Interpreted Type
	RTROIInterpretedType = New(0x3006, 0x00A4)

	// ROIInterpreter (3006,00A6) VR=PN VM=1 ROI Interpreter
	ROIInterpreter = New(0x3006, 0x00A6)

	// ROIPhysicalPropertiesSequence (3006,00B0) VR=SQ VM=1 ROI Physical Properties Sequence
	ROIPhysicalPropertiesSequence = New(0x3006, 0x00B0)

	// ROIPhysicalProperty (3006,00B2) VR=CS VM=1 ROI Physical Property
	ROIPhysicalProperty = New(0x3006, 0x00B2)

	// ROIPhysicalPropertyValue (3006,00B4) VR=DS VM=1 ROI Physical Property Value
	ROIPhysicalPropertyValue = New(0x3006, 0x00B4)

	// ROIElementalCompositionSequence (3006,00B6) VR=SQ VM=1 ROI Elemental Composition Sequence
	ROIElementalCompositionSequence = New(0x3006, 0x00B6)

	// ROIElementalCompositionAtomicNumber (3006,00B7) VR=US VM=1 ROI Elemental Composition Atomic Number
	ROIElementalCompositionAtomicNumber = New(0x3006, 0x00B7)

	// ROIElementalCompositionAtomicMassFraction (3006,00B8) VR=FL VM=1 ROI Elemental Composition Atomic Mass Fraction
	ROIElementalCompositionAtomicMassFraction = New(0x3006, 0x00B8)

	// AdditionalRTROIIdentificationCodeSequenceRETIRED (3006,00B9) VR=SQ VM=1 Additional RT ROI Identification Code Sequence (RETIRED)
	AdditionalRTROIIdentificationCodeSequenceRETIRED = New(0x3006, 0x00B9)

	// FrameOfReferenceRelationshipSequenceRETIRED (3006,00C0) VR=SQ VM=1 Frame of Reference Relationship Sequence (RETIRED)
	FrameOfReferenceRelationshipSequenceRETIRED = New(0x3006, 0x00C0)

	// RelatedFrameOfReferenceUIDRETIRED (3006,00C2) VR=UI VM=1 Related Frame of Reference UID (RETIRED)
	RelatedFrameOfReferenceUIDRETIRED = New(0x3006, 0x00C2)

	// FrameOfReferenceTransformationTypeRETIRED (3006,00C4) VR=CS VM=1 Frame of Reference Transformation Type (RETIRED)
	FrameOfReferenceTransformationTypeRETIRED = New(0x3006, 0x00C4)

	// FrameOfReferenceTransformationMatrix (3006,00C6) VR=DS VM=16 Frame of Reference Transformation Matrix
	FrameOfReferenceTransformationMatrix = New(0x3006, 0x00C6)

	// FrameOfReferenceTransformationComment (3006,00C8) VR=LO VM=1 Frame of Reference Transformation Comment
	FrameOfReferenceTransformationComment = New(0x3006, 0x00C8)

	// PatientLocationCoordinatesSequence (3006,00C9) VR=SQ VM=1 Patient Location Coordinates Sequence
	PatientLocationCoordinatesSequence = New(0x3006, 0x00C9)

	// PatientLocationCoordinatesCodeSequence (3006,00CA) VR=SQ VM=1 Patient Location Coordinates Code Sequence
	PatientLocationCoordinatesCodeSequence = New(0x3006, 0x00CA)

	// PatientSupportPositionSequence (3006,00CB) VR=SQ VM=1 Patient Support Position Sequence
	PatientSupportPositionSequence = New(0x3006, 0x00CB)

	// MeasuredDoseReferenceSequence (3008,0010) VR=SQ VM=1 Measured Dose Reference Sequence
	MeasuredDoseReferenceSequence = New(0x3008, 0x0010)

	// MeasuredDoseDescription (3008,0012) VR=ST VM=1 Measured Dose Description
	MeasuredDoseDescription = New(0x3008, 0x0012)

	// MeasuredDoseType (3008,0014) VR=CS VM=1 Measured Dose Type
	MeasuredDoseType = New(0x3008, 0x0014)

	// MeasuredDoseValue (3008,0016) VR=DS VM=1 Measured Dose Value
	MeasuredDoseValue = New(0x3008, 0x0016)

	// TreatmentSessionBeamSequence (3008,0020) VR=SQ VM=1 Treatment Session Beam Sequence
	TreatmentSessionBeamSequence = New(0x3008, 0x0020)

	// TreatmentSessionIonBeamSequence (3008,0021) VR=SQ VM=1 Treatment Session Ion Beam Sequence
	TreatmentSessionIonBeamSequence = New(0x3008, 0x0021)

	// CurrentFractionNumber (3008,0022) VR=IS VM=1 Current Fraction Number
	CurrentFractionNumber = New(0x3008, 0x0022)

	// TreatmentControlPointDate (3008,0024) VR=DA VM=1 Treatment Control Point Date
	TreatmentControlPointDate = New(0x3008, 0x0024)

	// TreatmentControlPointTime (3008,0025) VR=TM VM=1 Treatment Control Point Time
	TreatmentControlPointTime = New(0x3008, 0x0025)

	// TreatmentTerminationStatus (3008,002A) VR=CS VM=1 Treatment Termination Status
	TreatmentTerminationStatus = New(0x3008, 0x002A)

	// TreatmentTerminationCodeRETIRED (3008,002B) VR=SH VM=1 Treatment Termination Code (RETIRED)
	TreatmentTerminationCodeRETIRED = New(0x3008, 0x002B)

	// TreatmentVerificationStatus (3008,002C) VR=CS VM=1 Treatment Verification Status
	TreatmentVerificationStatus = New(0x3008, 0x002C)

	// ReferencedTreatmentRecordSequence (3008,0030) VR=SQ VM=1 Referenced Treatment Record Sequence
	ReferencedTreatmentRecordSequence = New(0x3008, 0x0030)

	// SpecifiedPrimaryMeterset (3008,0032) VR=DS VM=1 Specified Primary Meterset
	SpecifiedPrimaryMeterset = New(0x3008, 0x0032)

	// SpecifiedSecondaryMeterset (3008,0033) VR=DS VM=1 Specified Secondary Meterset
	SpecifiedSecondaryMeterset = New(0x3008, 0x0033)

	// DeliveredPrimaryMeterset (3008,0036) VR=DS VM=1 Delivered Primary Meterset
	DeliveredPrimaryMeterset = New(0x3008, 0x0036)

	// DeliveredSecondaryMeterset (3008,0037) VR=DS VM=1 Delivered Secondary Meterset
	DeliveredSecondaryMeterset = New(0x3008, 0x0037)

	// SpecifiedTreatmentTime (3008,003A) VR=DS VM=1 Specified Treatment Time
	SpecifiedTreatmentTime = New(0x3008, 0x003A)

	// DeliveredTreatmentTime (3008,003B) VR=DS VM=1 Delivered Treatment Time
	DeliveredTreatmentTime = New(0x3008, 0x003B)

	// ControlPointDeliverySequence (3008,0040) VR=SQ VM=1 Control Point Delivery Sequence
	ControlPointDeliverySequence = New(0x3008, 0x0040)

	// IonControlPointDeliverySequence (3008,0041) VR=SQ VM=1 Ion Control Point Delivery Sequence
	IonControlPointDeliverySequence = New(0x3008, 0x0041)

	// SpecifiedMeterset (3008,0042) VR=DS VM=1 Specified Meterset
	SpecifiedMeterset = New(0x3008, 0x0042)

	// DeliveredMeterset (3008,0044) VR=DS VM=1 Delivered Meterset
	DeliveredMeterset = New(0x3008, 0x0044)

	// MetersetRateSet (3008,0045) VR=FL VM=1 Meterset Rate Set
	MetersetRateSet = New(0x3008, 0x0045)

	// MetersetRateDelivered (3008,0046) VR=FL VM=1 Meterset Rate Delivered
	MetersetRateDelivered = New(0x3008, 0x0046)

	// ScanSpotMetersetsDelivered (3008,0047) VR=FL VM=1-n Scan Spot Metersets Delivered
	ScanSpotMetersetsDelivered = New(0x3008, 0x0047)

	// DoseRateDelivered (3008,0048) VR=DS VM=1 Dose Rate Delivered
	DoseRateDelivered = New(0x3008, 0x0048)

	// TreatmentSummaryCalculatedDoseReferenceSequence (3008,0050) VR=SQ VM=1 Treatment Summary Calculated Dose Reference Sequence
	TreatmentSummaryCalculatedDoseReferenceSequence = New(0x3008, 0x0050)

	// CumulativeDoseToDoseReference (3008,0052) VR=DS VM=1 Cumulative Dose to Dose Reference
	CumulativeDoseToDoseReference = New(0x3008, 0x0052)

	// FirstTreatmentDate (3008,0054) VR=DA VM=1 First Treatment Date
	FirstTreatmentDate = New(0x3008, 0x0054)

	// MostRecentTreatmentDate (3008,0056) VR=DA VM=1 Most Recent Treatment Date
	MostRecentTreatmentDate = New(0x3008, 0x0056)

	// NumberOfFractionsDelivered (3008,005A) VR=IS VM=1 Number of Fractions Delivered
	NumberOfFractionsDelivered = New(0x3008, 0x005A)

	// OverrideSequence (3008,0060) VR=SQ VM=1 Override Sequence
	OverrideSequence = New(0x3008, 0x0060)

	// ParameterSequencePointer (3008,0061) VR=AT VM=1 Parameter Sequence Pointer
	ParameterSequencePointer = New(0x3008, 0x0061)

	// OverrideParameterPointer (3008,0062) VR=AT VM=1 Override Parameter Pointer
	OverrideParameterPointer = New(0x3008, 0x0062)

	// ParameterItemIndex (3008,0063) VR=IS VM=1 Parameter Item Index
	ParameterItemIndex = New(0x3008, 0x0063)

	// MeasuredDoseReferenceNumber (3008,0064) VR=IS VM=1 Measured Dose Reference Number
	MeasuredDoseReferenceNumber = New(0x3008, 0x0064)

	// ParameterPointer (3008,0065) VR=AT VM=1 Parameter Pointer
	ParameterPointer = New(0x3008, 0x0065)

	// OverrideReason (3008,0066) VR=ST VM=1 Override Reason
	OverrideReason = New(0x3008, 0x0066)

	// ParameterValueNumber (3008,0067) VR=US VM=1 Parameter Value Number
	ParameterValueNumber = New(0x3008, 0x0067)

	// CorrectedParameterSequence (3008,0068) VR=SQ VM=1 Corrected Parameter Sequence
	CorrectedParameterSequence = New(0x3008, 0x0068)

	// CorrectionValue (3008,006A) VR=FL VM=1 Correction Value
	CorrectionValue = New(0x3008, 0x006A)

	// CalculatedDoseReferenceSequence (3008,0070) VR=SQ VM=1 Calculated Dose Reference Sequence
	CalculatedDoseReferenceSequence = New(0x3008, 0x0070)

	// CalculatedDoseReferenceNumber (3008,0072) VR=IS VM=1 Calculated Dose Reference Number
	CalculatedDoseReferenceNumber = New(0x3008, 0x0072)

	// CalculatedDoseReferenceDescription (3008,0074) VR=ST VM=1 Calculated Dose Reference Description
	CalculatedDoseReferenceDescription = New(0x3008, 0x0074)

	// CalculatedDoseReferenceDoseValue (3008,0076) VR=DS VM=1 Calculated Dose Reference Dose Value
	CalculatedDoseReferenceDoseValue = New(0x3008, 0x0076)

	// StartMeterset (3008,0078) VR=DS VM=1 Start Meterset
	StartMeterset = New(0x3008, 0x0078)

	// EndMeterset (3008,007A) VR=DS VM=1 End Meterset
	EndMeterset = New(0x3008, 0x007A)

	// ReferencedMeasuredDoseReferenceSequence (3008,0080) VR=SQ VM=1 Referenced Measured Dose Reference Sequence
	ReferencedMeasuredDoseReferenceSequence = New(0x3008, 0x0080)

	// ReferencedMeasuredDoseReferenceNumber (3008,0082) VR=IS VM=1 Referenced Measured Dose Reference Number
	ReferencedMeasuredDoseReferenceNumber = New(0x3008, 0x0082)

	// ReferencedCalculatedDoseReferenceSequence (3008,0090) VR=SQ VM=1 Referenced Calculated Dose Reference Sequence
	ReferencedCalculatedDoseReferenceSequence = New(0x3008, 0x0090)

	// ReferencedCalculatedDoseReferenceNumber (3008,0092) VR=IS VM=1 Referenced Calculated Dose Reference Number
	ReferencedCalculatedDoseReferenceNumber = New(0x3008, 0x0092)

	// BeamLimitingDeviceLeafPairsSequence (3008,00A0) VR=SQ VM=1 Beam Limiting Device Leaf Pairs Sequence
	BeamLimitingDeviceLeafPairsSequence = New(0x3008, 0x00A0)

	// EnhancedRTBeamLimitingDeviceSequence (3008,00A1) VR=SQ VM=1 Enhanced RT Beam Limiting Device Sequence
	EnhancedRTBeamLimitingDeviceSequence = New(0x3008, 0x00A1)

	// EnhancedRTBeamLimitingOpeningSequence (3008,00A2) VR=SQ VM=1 Enhanced RT Beam Limiting Opening Sequence
	EnhancedRTBeamLimitingOpeningSequence = New(0x3008, 0x00A2)

	// EnhancedRTBeamLimitingDeviceDefinitionFlag (3008,00A3) VR=CS VM=1 Enhanced RT Beam Limiting Device Definition Flag
	EnhancedRTBeamLimitingDeviceDefinitionFlag = New(0x3008, 0x00A3)

	// ParallelRTBeamDelimiterOpeningExtents (3008,00A4) VR=FD VM=2-2n Parallel RT Beam Delimiter Opening Extents
	ParallelRTBeamDelimiterOpeningExtents = New(0x3008, 0x00A4)

	// RecordedWedgeSequence (3008,00B0) VR=SQ VM=1 Recorded Wedge Sequence
	RecordedWedgeSequence = New(0x3008, 0x00B0)

	// RecordedCompensatorSequence (3008,00C0) VR=SQ VM=1 Recorded Compensator Sequence
	RecordedCompensatorSequence = New(0x3008, 0x00C0)

	// RecordedBlockSequence (3008,00D0) VR=SQ VM=1 Recorded Block Sequence
	RecordedBlockSequence = New(0x3008, 0x00D0)

	// RecordedBlockSlabSequence (3008,00D1) VR=SQ VM=1 Recorded Block Slab Sequence
	RecordedBlockSlabSequence = New(0x3008, 0x00D1)

	// TreatmentSummaryMeasuredDoseReferenceSequence (3008,00E0) VR=SQ VM=1 Treatment Summary Measured Dose Reference Sequence
	TreatmentSummaryMeasuredDoseReferenceSequence = New(0x3008, 0x00E0)

	// RecordedSnoutSequence (3008,00F0) VR=SQ VM=1 Recorded Snout Sequence
	RecordedSnoutSequence = New(0x3008, 0x00F0)

	// RecordedRangeShifterSequence (3008,00F2) VR=SQ VM=1 Recorded Range Shifter Sequence
	RecordedRangeShifterSequence = New(0x3008, 0x00F2)

	// RecordedLateralSpreadingDeviceSequence (3008,00F4) VR=SQ VM=1 Recorded Lateral Spreading Device Sequence
	RecordedLateralSpreadingDeviceSequence = New(0x3008, 0x00F4)

	// RecordedRangeModulatorSequence (3008,00F6) VR=SQ VM=1 Recorded Range Modulator Sequence
	RecordedRangeModulatorSequence = New(0x3008, 0x00F6)

	// RecordedSourceSequence (3008,0100) VR=SQ VM=1 Recorded Source Sequence
	RecordedSourceSequence = New(0x3008, 0x0100)

	// SourceSerialNumber (3008,0105) VR=LO VM=1 Source Serial Number
	SourceSerialNumber = New(0x3008, 0x0105)

	// TreatmentSessionApplicationSetupSequence (3008,0110) VR=SQ VM=1 Treatment Session Application Setup Sequence
	TreatmentSessionApplicationSetupSequence = New(0x3008, 0x0110)

	// ApplicationSetupCheck (3008,0116) VR=CS VM=1 Application Setup Check
	ApplicationSetupCheck = New(0x3008, 0x0116)

	// RecordedBrachyAccessoryDeviceSequence (3008,0120) VR=SQ VM=1 Recorded Brachy Accessory Device Sequence
	RecordedBrachyAccessoryDeviceSequence = New(0x3008, 0x0120)

	// ReferencedBrachyAccessoryDeviceNumber (3008,0122) VR=IS VM=1 Referenced Brachy Accessory Device Number
	ReferencedBrachyAccessoryDeviceNumber = New(0x3008, 0x0122)

	// RecordedChannelSequence (3008,0130) VR=SQ VM=1 Recorded Channel Sequence
	RecordedChannelSequence = New(0x3008, 0x0130)

	// SpecifiedChannelTotalTime (3008,0132) VR=DS VM=1 Specified Channel Total Time
	SpecifiedChannelTotalTime = New(0x3008, 0x0132)

	// DeliveredChannelTotalTime (3008,0134) VR=DS VM=1 Delivered Channel Total Time
	DeliveredChannelTotalTime = New(0x3008, 0x0134)

	// SpecifiedNumberOfPulses (3008,0136) VR=IS VM=1 Specified Number of Pulses
	SpecifiedNumberOfPulses = New(0x3008, 0x0136)

	// DeliveredNumberOfPulses (3008,0138) VR=IS VM=1 Delivered Number of Pulses
	DeliveredNumberOfPulses = New(0x3008, 0x0138)

	// SpecifiedPulseRepetitionInterval (3008,013A) VR=DS VM=1 Specified Pulse Repetition Interval
	SpecifiedPulseRepetitionInterval = New(0x3008, 0x013A)

	// DeliveredPulseRepetitionInterval (3008,013C) VR=DS VM=1 Delivered Pulse Repetition Interval
	DeliveredPulseRepetitionInterval = New(0x3008, 0x013C)

	// RecordedSourceApplicatorSequence (3008,0140) VR=SQ VM=1 Recorded Source Applicator Sequence
	RecordedSourceApplicatorSequence = New(0x3008, 0x0140)

	// ReferencedSourceApplicatorNumber (3008,0142) VR=IS VM=1 Referenced Source Applicator Number
	ReferencedSourceApplicatorNumber = New(0x3008, 0x0142)

	// RecordedChannelShieldSequence (3008,0150) VR=SQ VM=1 Recorded Channel Shield Sequence
	RecordedChannelShieldSequence = New(0x3008, 0x0150)

	// ReferencedChannelShieldNumber (3008,0152) VR=IS VM=1 Referenced Channel Shield Number
	ReferencedChannelShieldNumber = New(0x3008, 0x0152)

	// BrachyControlPointDeliveredSequence (3008,0160) VR=SQ VM=1 Brachy Control Point Delivered Sequence
	BrachyControlPointDeliveredSequence = New(0x3008, 0x0160)

	// SafePositionExitDate (3008,0162) VR=DA VM=1 Safe Position Exit Date
	SafePositionExitDate = New(0x3008, 0x0162)

	// SafePositionExitTime (3008,0164) VR=TM VM=1 Safe Position Exit Time
	SafePositionExitTime = New(0x3008, 0x0164)

	// SafePositionReturnDate (3008,0166) VR=DA VM=1 Safe Position Return Date
	SafePositionReturnDate = New(0x3008, 0x0166)

	// SafePositionReturnTime (3008,0168) VR=TM VM=1 Safe Position Return Time
	SafePositionReturnTime = New(0x3008, 0x0168)

	// PulseSpecificBrachyControlPointDeliveredSequence (3008,0171) VR=SQ VM=1 Pulse Specific Brachy Control Point Delivered Sequence
	PulseSpecificBrachyControlPointDeliveredSequence = New(0x3008, 0x0171)

	// PulseNumber (3008,0172) VR=US VM=1 Pulse Number
	PulseNumber = New(0x3008, 0x0172)

	// BrachyPulseControlPointDeliveredSequence (3008,0173) VR=SQ VM=1 Brachy Pulse Control Point Delivered Sequence
	BrachyPulseControlPointDeliveredSequence = New(0x3008, 0x0173)

	// CurrentTreatmentStatus (3008,0200) VR=CS VM=1 Current Treatment Status
	CurrentTreatmentStatus = New(0x3008, 0x0200)

	// TreatmentStatusComment (3008,0202) VR=ST VM=1 Treatment Status Comment
	TreatmentStatusComment = New(0x3008, 0x0202)

	// FractionGroupSummarySequence (3008,0220) VR=SQ VM=1 Fraction Group Summary Sequence
	FractionGroupSummarySequence = New(0x3008, 0x0220)

	// ReferencedFractionNumber (3008,0223) VR=IS VM=1 Referenced Fraction Number
	ReferencedFractionNumber = New(0x3008, 0x0223)

	// FractionGroupType (3008,0224) VR=CS VM=1 Fraction Group Type
	FractionGroupType = New(0x3008, 0x0224)

	// BeamStopperPosition (3008,0230) VR=CS VM=1 Beam Stopper Position
	BeamStopperPosition = New(0x3008, 0x0230)

	// FractionStatusSummarySequence (3008,0240) VR=SQ VM=1 Fraction Status Summary Sequence
	FractionStatusSummarySequence = New(0x3008, 0x0240)

	// TreatmentDate (3008,0250) VR=DA VM=1 Treatment Date
	TreatmentDate = New(0x3008, 0x0250)

	// TreatmentTime (3008,0251) VR=TM VM=1 Treatment Time
	TreatmentTime = New(0x3008, 0x0251)

	// RTPlanLabel (300A,0002) VR=SH VM=1 RT Plan Label
	RTPlanLabel = New(0x300A, 0x0002)

	// RTPlanName (300A,0003) VR=LO VM=1 RT Plan Name
	RTPlanName = New(0x300A, 0x0003)

	// RTPlanDescription (300A,0004) VR=ST VM=1 RT Plan Description
	RTPlanDescription = New(0x300A, 0x0004)

	// RTPlanDate (300A,0006) VR=DA VM=1 RT Plan Date
	RTPlanDate = New(0x300A, 0x0006)

	// RTPlanTime (300A,0007) VR=TM VM=1 RT Plan Time
	RTPlanTime = New(0x300A, 0x0007)

	// TreatmentProtocols (300A,0009) VR=LO VM=1-n Treatment Protocols
	TreatmentProtocols = New(0x300A, 0x0009)

	// PlanIntent (300A,000A) VR=CS VM=1 Plan Intent
	PlanIntent = New(0x300A, 0x000A)

	// TreatmentSitesRETIRED (300A,000B) VR=LO VM=1-n Treatment Sites (RETIRED)
	TreatmentSitesRETIRED = New(0x300A, 0x000B)

	// RTPlanGeometry (300A,000C) VR=CS VM=1 RT Plan Geometry
	RTPlanGeometry = New(0x300A, 0x000C)

	// PrescriptionDescription (300A,000E) VR=ST VM=1 Prescription Description
	PrescriptionDescription = New(0x300A, 0x000E)

	// DoseReferenceSequence (300A,0010) VR=SQ VM=1 Dose Reference Sequence
	DoseReferenceSequence = New(0x300A, 0x0010)

	// DoseReferenceNumber (300A,0012) VR=IS VM=1 Dose Reference Number
	DoseReferenceNumber = New(0x300A, 0x0012)

	// DoseReferenceUID (300A,0013) VR=UI VM=1 Dose Reference UID
	DoseReferenceUID = New(0x300A, 0x0013)

	// DoseReferenceStructureType (300A,0014) VR=CS VM=1 Dose Reference Structure Type
	DoseReferenceStructureType = New(0x300A, 0x0014)

	// NominalBeamEnergyUnit (300A,0015) VR=CS VM=1 Nominal Beam Energy Unit
	NominalBeamEnergyUnit = New(0x300A, 0x0015)

	// DoseReferenceDescription (300A,0016) VR=LO VM=1 Dose Reference Description
	DoseReferenceDescription = New(0x300A, 0x0016)

	// DoseReferencePointCoordinates (300A,0018) VR=DS VM=3 Dose Reference Point Coordinates
	DoseReferencePointCoordinates = New(0x300A, 0x0018)

	// NominalPriorDose (300A,001A) VR=DS VM=1 Nominal Prior Dose
	NominalPriorDose = New(0x300A, 0x001A)

	// DoseReferenceType (300A,0020) VR=CS VM=1 Dose Reference Type
	DoseReferenceType = New(0x300A, 0x0020)

	// ConstraintWeight (300A,0021) VR=DS VM=1 Constraint Weight
	ConstraintWeight = New(0x300A, 0x0021)

	// DeliveryWarningDose (300A,0022) VR=DS VM=1 Delivery Warning Dose
	DeliveryWarningDose = New(0x300A, 0x0022)

	// DeliveryMaximumDose (300A,0023) VR=DS VM=1 Delivery Maximum Dose
	DeliveryMaximumDose = New(0x300A, 0x0023)

	// TargetMinimumDose (300A,0025) VR=DS VM=1 Target Minimum Dose
	TargetMinimumDose = New(0x300A, 0x0025)

	// TargetPrescriptionDose (300A,0026) VR=DS VM=1 Target Prescription Dose
	TargetPrescriptionDose = New(0x300A, 0x0026)

	// TargetMaximumDose (300A,0027) VR=DS VM=1 Target Maximum Dose
	TargetMaximumDose = New(0x300A, 0x0027)

	// TargetUnderdoseVolumeFraction (300A,0028) VR=DS VM=1 Target Underdose Volume Fraction
	TargetUnderdoseVolumeFraction = New(0x300A, 0x0028)

	// OrganAtRiskFullVolumeDose (300A,002A) VR=DS VM=1 Organ at Risk Full-volume Dose
	OrganAtRiskFullVolumeDose = New(0x300A, 0x002A)

	// OrganAtRiskLimitDose (300A,002B) VR=DS VM=1 Organ at Risk Limit Dose
	OrganAtRiskLimitDose = New(0x300A, 0x002B)

	// OrganAtRiskMaximumDose (300A,002C) VR=DS VM=1 Organ at Risk Maximum Dose
	OrganAtRiskMaximumDose = New(0x300A, 0x002C)

	// OrganAtRiskOverdoseVolumeFraction (300A,002D) VR=DS VM=1 Organ at Risk Overdose Volume Fraction
	OrganAtRiskOverdoseVolumeFraction = New(0x300A, 0x002D)

	// ToleranceTableSequence (300A,0040) VR=SQ VM=1 Tolerance Table Sequence
	ToleranceTableSequence = New(0x300A, 0x0040)

	// ToleranceTableNumber (300A,0042) VR=IS VM=1 Tolerance Table Number
	ToleranceTableNumber = New(0x300A, 0x0042)

	// ToleranceTableLabel (300A,0043) VR=SH VM=1 Tolerance Table Label
	ToleranceTableLabel = New(0x300A, 0x0043)

	// GantryAngleTolerance (300A,0044) VR=DS VM=1 Gantry Angle Tolerance
	GantryAngleTolerance = New(0x300A, 0x0044)

	// BeamLimitingDeviceAngleTolerance (300A,0046) VR=DS VM=1 Beam Limiting Device Angle Tolerance
	BeamLimitingDeviceAngleTolerance = New(0x300A, 0x0046)

	// BeamLimitingDeviceToleranceSequence (300A,0048) VR=SQ VM=1 Beam Limiting Device Tolerance Sequence
	BeamLimitingDeviceToleranceSequence = New(0x300A, 0x0048)

	// BeamLimitingDevicePositionTolerance (300A,004A) VR=DS VM=1 Beam Limiting Device Position Tolerance
	BeamLimitingDevicePositionTolerance = New(0x300A, 0x004A)

	// SnoutPositionTolerance (300A,004B) VR=FL VM=1 Snout Position Tolerance
	SnoutPositionTolerance = New(0x300A, 0x004B)

	// PatientSupportAngleTolerance (300A,004C) VR=DS VM=1 Patient Support Angle Tolerance
	PatientSupportAngleTolerance = New(0x300A, 0x004C)

	// TableTopEccentricAngleTolerance (300A,004E) VR=DS VM=1 Table Top Eccentric Angle Tolerance
	TableTopEccentricAngleTolerance = New(0x300A, 0x004E)

	// TableTopPitchAngleTolerance (300A,004F) VR=FL VM=1 Table Top Pitch Angle Tolerance
	TableTopPitchAngleTolerance = New(0x300A, 0x004F)

	// TableTopRollAngleTolerance (300A,0050) VR=FL VM=1 Table Top Roll Angle Tolerance
	TableTopRollAngleTolerance = New(0x300A, 0x0050)

	// TableTopVerticalPositionTolerance (300A,0051) VR=DS VM=1 Table Top Vertical Position Tolerance
	TableTopVerticalPositionTolerance = New(0x300A, 0x0051)

	// TableTopLongitudinalPositionTolerance (300A,0052) VR=DS VM=1 Table Top Longitudinal Position Tolerance
	TableTopLongitudinalPositionTolerance = New(0x300A, 0x0052)

	// TableTopLateralPositionTolerance (300A,0053) VR=DS VM=1 Table Top Lateral Position Tolerance
	TableTopLateralPositionTolerance = New(0x300A, 0x0053)

	// TableTopPositionAlignmentUID (300A,0054) VR=UI VM=1 Table Top Position Alignment UID
	TableTopPositionAlignmentUID = New(0x300A, 0x0054)

	// RTPlanRelationship (300A,0055) VR=CS VM=1 RT Plan Relationship
	RTPlanRelationship = New(0x300A, 0x0055)

	// FractionGroupSequence (300A,0070) VR=SQ VM=1 Fraction Group Sequence
	FractionGroupSequence = New(0x300A, 0x0070)

	// FractionGroupNumber (300A,0071) VR=IS VM=1 Fraction Group Number
	FractionGroupNumber = New(0x300A, 0x0071)

	// FractionGroupDescription (300A,0072) VR=LO VM=1 Fraction Group Description
	FractionGroupDescription = New(0x300A, 0x0072)

	// NumberOfFractionsPlanned (300A,0078) VR=IS VM=1 Number of Fractions Planned
	NumberOfFractionsPlanned = New(0x300A, 0x0078)

	// NumberOfFractionPatternDigitsPerDay (300A,0079) VR=IS VM=1 Number of Fraction Pattern Digits Per Day
	NumberOfFractionPatternDigitsPerDay = New(0x300A, 0x0079)

	// RepeatFractionCycleLength (300A,007A) VR=IS VM=1 Repeat Fraction Cycle Length
	RepeatFractionCycleLength = New(0x300A, 0x007A)

	// FractionPattern (300A,007B) VR=LT VM=1 Fraction Pattern
	FractionPattern = New(0x300A, 0x007B)

	// NumberOfBeams (300A,0080) VR=IS VM=1 Number of Beams
	NumberOfBeams = New(0x300A, 0x0080)

	// BeamDoseSpecificationPointRETIRED (300A,0082) VR=DS VM=3 Beam Dose Specification Point (RETIRED)
	BeamDoseSpecificationPointRETIRED = New(0x300A, 0x0082)

	// ReferencedDoseReferenceUID (300A,0083) VR=UI VM=1 Referenced Dose Reference UID
	ReferencedDoseReferenceUID = New(0x300A, 0x0083)

	// BeamDose (300A,0084) VR=DS VM=1 Beam Dose
	BeamDose = New(0x300A, 0x0084)

	// BeamMeterset (300A,0086) VR=DS VM=1 Beam Meterset
	BeamMeterset = New(0x300A, 0x0086)

	// BeamDosePointDepth (300A,0088) VR=FL VM=1 Beam Dose Point Depth
	BeamDosePointDepth = New(0x300A, 0x0088)

	// BeamDosePointEquivalentDepth (300A,0089) VR=FL VM=1 Beam Dose Point Equivalent Depth
	BeamDosePointEquivalentDepth = New(0x300A, 0x0089)

	// BeamDosePointSSD (300A,008A) VR=FL VM=1 Beam Dose Point SSD
	BeamDosePointSSD = New(0x300A, 0x008A)

	// BeamDoseMeaning (300A,008B) VR=CS VM=1 Beam Dose Meaning
	BeamDoseMeaning = New(0x300A, 0x008B)

	// BeamDoseVerificationControlPointSequence (300A,008C) VR=SQ VM=1 Beam Dose Verification Control Point Sequence
	BeamDoseVerificationControlPointSequence = New(0x300A, 0x008C)

	// AverageBeamDosePointDepthRETIRED (300A,008D) VR=FL VM=1 Average Beam Dose Point Depth (RETIRED)
	AverageBeamDosePointDepthRETIRED = New(0x300A, 0x008D)

	// AverageBeamDosePointEquivalentDepthRETIRED (300A,008E) VR=FL VM=1 Average Beam Dose Point Equivalent Depth (RETIRED)
	AverageBeamDosePointEquivalentDepthRETIRED = New(0x300A, 0x008E)

	// AverageBeamDosePointSSDRETIRED (300A,008F) VR=FL VM=1 Average Beam Dose Point SSD (RETIRED)
	AverageBeamDosePointSSDRETIRED = New(0x300A, 0x008F)

	// BeamDoseType (300A,0090) VR=CS VM=1 Beam Dose Type
	BeamDoseType = New(0x300A, 0x0090)

	// AlternateBeamDose (300A,0091) VR=DS VM=1 Alternate Beam Dose
	AlternateBeamDose = New(0x300A, 0x0091)

	// AlternateBeamDoseType (300A,0092) VR=CS VM=1 Alternate Beam Dose Type
	AlternateBeamDoseType = New(0x300A, 0x0092)

	// DepthValueAveragingFlag (300A,0093) VR=CS VM=1 Depth Value Averaging Flag
	DepthValueAveragingFlag = New(0x300A, 0x0093)

	// BeamDosePointSourceToExternalContourDistance (300A,0094) VR=DS VM=1 Beam Dose Point Source to External Contour Distance
	BeamDosePointSourceToExternalContourDistance = New(0x300A, 0x0094)

	// NumberOfBrachyApplicationSetups (300A,00A0) VR=IS VM=1 Number of Brachy Application Setups
	NumberOfBrachyApplicationSetups = New(0x300A, 0x00A0)

	// BrachyApplicationSetupDoseSpecificationPoint (300A,00A2) VR=DS VM=3 Brachy Application Setup Dose Specification Point
	BrachyApplicationSetupDoseSpecificationPoint = New(0x300A, 0x00A2)

	// BrachyApplicationSetupDose (300A,00A4) VR=DS VM=1 Brachy Application Setup Dose
	BrachyApplicationSetupDose = New(0x300A, 0x00A4)

	// BeamSequence (300A,00B0) VR=SQ VM=1 Beam Sequence
	BeamSequence = New(0x300A, 0x00B0)

	// TreatmentMachineName (300A,00B2) VR=SH VM=1 Treatment Machine Name
	TreatmentMachineName = New(0x300A, 0x00B2)

	// PrimaryDosimeterUnit (300A,00B3) VR=CS VM=1 Primary Dosimeter Unit
	PrimaryDosimeterUnit = New(0x300A, 0x00B3)

	// SourceAxisDistance (300A,00B4) VR=DS VM=1 Source-Axis Distance
	SourceAxisDistance = New(0x300A, 0x00B4)

	// BeamLimitingDeviceSequence (300A,00B6) VR=SQ VM=1 Beam Limiting Device Sequence
	BeamLimitingDeviceSequence = New(0x300A, 0x00B6)

	// RTBeamLimitingDeviceType (300A,00B8) VR=CS VM=1 RT Beam Limiting Device Type
	RTBeamLimitingDeviceType = New(0x300A, 0x00B8)

	// SourceToBeamLimitingDeviceDistance (300A,00BA) VR=DS VM=1 Source to Beam Limiting Device Distance
	SourceToBeamLimitingDeviceDistance = New(0x300A, 0x00BA)

	// IsocenterToBeamLimitingDeviceDistance (300A,00BB) VR=FL VM=1 Isocenter to Beam Limiting Device Distance
	IsocenterToBeamLimitingDeviceDistance = New(0x300A, 0x00BB)

	// NumberOfLeafJawPairs (300A,00BC) VR=IS VM=1 Number of Leaf/Jaw Pairs
	NumberOfLeafJawPairs = New(0x300A, 0x00BC)

	// LeafPositionBoundaries (300A,00BE) VR=DS VM=3-n Leaf Position Boundaries
	LeafPositionBoundaries = New(0x300A, 0x00BE)

	// BeamNumber (300A,00C0) VR=IS VM=1 Beam Number
	BeamNumber = New(0x300A, 0x00C0)

	// BeamName (300A,00C2) VR=LO VM=1 Beam Name
	BeamName = New(0x300A, 0x00C2)

	// BeamDescription (300A,00C3) VR=ST VM=1 Beam Description
	BeamDescription = New(0x300A, 0x00C3)

	// BeamType (300A,00C4) VR=CS VM=1 Beam Type
	BeamType = New(0x300A, 0x00C4)

	// BeamDeliveryDurationLimit (300A,00C5) VR=FD VM=1 Beam Delivery Duration Limit
	BeamDeliveryDurationLimit = New(0x300A, 0x00C5)

	// RadiationType (300A,00C6) VR=CS VM=1 Radiation Type
	RadiationType = New(0x300A, 0x00C6)

	// HighDoseTechniqueType (300A,00C7) VR=CS VM=1 High-Dose Technique Type
	HighDoseTechniqueType = New(0x300A, 0x00C7)

	// ReferenceImageNumber (300A,00C8) VR=IS VM=1 Reference Image Number
	ReferenceImageNumber = New(0x300A, 0x00C8)

	// PlannedVerificationImageSequence (300A,00CA) VR=SQ VM=1 Planned Verification Image Sequence
	PlannedVerificationImageSequence = New(0x300A, 0x00CA)

	// ImagingDeviceSpecificAcquisitionParameters (300A,00CC) VR=LO VM=1-n Imaging Device-Specific Acquisition Parameters
	ImagingDeviceSpecificAcquisitionParameters = New(0x300A, 0x00CC)

	// TreatmentDeliveryType (300A,00CE) VR=CS VM=1 Treatment Delivery Type
	TreatmentDeliveryType = New(0x300A, 0x00CE)

	// NumberOfWedges (300A,00D0) VR=IS VM=1 Number of Wedges
	NumberOfWedges = New(0x300A, 0x00D0)

	// WedgeSequence (300A,00D1) VR=SQ VM=1 Wedge Sequence
	WedgeSequence = New(0x300A, 0x00D1)

	// WedgeNumber (300A,00D2) VR=IS VM=1 Wedge Number
	WedgeNumber = New(0x300A, 0x00D2)

	// WedgeType (300A,00D3) VR=CS VM=1 Wedge Type
	WedgeType = New(0x300A, 0x00D3)

	// WedgeID (300A,00D4) VR=SH VM=1 Wedge ID
	WedgeID = New(0x300A, 0x00D4)

	// WedgeAngle (300A,00D5) VR=IS VM=1 Wedge Angle
	WedgeAngle = New(0x300A, 0x00D5)

	// WedgeFactor (300A,00D6) VR=DS VM=1 Wedge Factor
	WedgeFactor = New(0x300A, 0x00D6)

	// TotalWedgeTrayWaterEquivalentThickness (300A,00D7) VR=FL VM=1 Total Wedge Tray Water-Equivalent Thickness
	TotalWedgeTrayWaterEquivalentThickness = New(0x300A, 0x00D7)

	// WedgeOrientation (300A,00D8) VR=DS VM=1 Wedge Orientation
	WedgeOrientation = New(0x300A, 0x00D8)

	// IsocenterToWedgeTrayDistance (300A,00D9) VR=FL VM=1 Isocenter to Wedge Tray Distance
	IsocenterToWedgeTrayDistance = New(0x300A, 0x00D9)

	// SourceToWedgeTrayDistance (300A,00DA) VR=DS VM=1 Source to Wedge Tray Distance
	SourceToWedgeTrayDistance = New(0x300A, 0x00DA)

	// WedgeThinEdgePosition (300A,00DB) VR=FL VM=1 Wedge Thin Edge Position
	WedgeThinEdgePosition = New(0x300A, 0x00DB)

	// BolusID (300A,00DC) VR=SH VM=1 Bolus ID
	BolusID = New(0x300A, 0x00DC)

	// BolusDescription (300A,00DD) VR=ST VM=1 Bolus Description
	BolusDescription = New(0x300A, 0x00DD)

	// EffectiveWedgeAngle (300A,00DE) VR=DS VM=1 Effective Wedge Angle
	EffectiveWedgeAngle = New(0x300A, 0x00DE)

	// NumberOfCompensators (300A,00E0) VR=IS VM=1 Number of Compensators
	NumberOfCompensators = New(0x300A, 0x00E0)

	// MaterialID (300A,00E1) VR=SH VM=1 Material ID
	MaterialID = New(0x300A, 0x00E1)

	// TotalCompensatorTrayFactor (300A,00E2) VR=DS VM=1 Total Compensator Tray Factor
	TotalCompensatorTrayFactor = New(0x300A, 0x00E2)

	// CompensatorSequence (300A,00E3) VR=SQ VM=1 Compensator Sequence
	CompensatorSequence = New(0x300A, 0x00E3)

	// CompensatorNumber (300A,00E4) VR=IS VM=1 Compensator Number
	CompensatorNumber = New(0x300A, 0x00E4)

	// CompensatorID (300A,00E5) VR=SH VM=1 Compensator ID
	CompensatorID = New(0x300A, 0x00E5)

	// SourceToCompensatorTrayDistance (300A,00E6) VR=DS VM=1 Source to Compensator Tray Distance
	SourceToCompensatorTrayDistance = New(0x300A, 0x00E6)

	// CompensatorRows (300A,00E7) VR=IS VM=1 Compensator Rows
	CompensatorRows = New(0x300A, 0x00E7)

	// CompensatorColumns (300A,00E8) VR=IS VM=1 Compensator Columns
	CompensatorColumns = New(0x300A, 0x00E8)

	// CompensatorPixelSpacing (300A,00E9) VR=DS VM=2 Compensator Pixel Spacing
	CompensatorPixelSpacing = New(0x300A, 0x00E9)

	// CompensatorPosition (300A,00EA) VR=DS VM=2 Compensator Position
	CompensatorPosition = New(0x300A, 0x00EA)

	// CompensatorTransmissionData (300A,00EB) VR=DS VM=1-n Compensator Transmission Data
	CompensatorTransmissionData = New(0x300A, 0x00EB)

	// CompensatorThicknessData (300A,00EC) VR=DS VM=1-n Compensator Thickness Data
	CompensatorThicknessData = New(0x300A, 0x00EC)

	// NumberOfBoli (300A,00ED) VR=IS VM=1 Number of Boli
	NumberOfBoli = New(0x300A, 0x00ED)

	// CompensatorType (300A,00EE) VR=CS VM=1 Compensator Type
	CompensatorType = New(0x300A, 0x00EE)

	// CompensatorTrayID (300A,00EF) VR=SH VM=1 Compensator Tray ID
	CompensatorTrayID = New(0x300A, 0x00EF)

	// NumberOfBlocks (300A,00F0) VR=IS VM=1 Number of Blocks
	NumberOfBlocks = New(0x300A, 0x00F0)

	// TotalBlockTrayFactor (300A,00F2) VR=DS VM=1 Total Block Tray Factor
	TotalBlockTrayFactor = New(0x300A, 0x00F2)

	// TotalBlockTrayWaterEquivalentThickness (300A,00F3) VR=FL VM=1 Total Block Tray Water-Equivalent Thickness
	TotalBlockTrayWaterEquivalentThickness = New(0x300A, 0x00F3)

	// BlockSequence (300A,00F4) VR=SQ VM=1 Block Sequence
	BlockSequence = New(0x300A, 0x00F4)

	// BlockTrayID (300A,00F5) VR=SH VM=1 Block Tray ID
	BlockTrayID = New(0x300A, 0x00F5)

	// SourceToBlockTrayDistance (300A,00F6) VR=DS VM=1 Source to Block Tray Distance
	SourceToBlockTrayDistance = New(0x300A, 0x00F6)

	// IsocenterToBlockTrayDistance (300A,00F7) VR=FL VM=1 Isocenter to Block Tray Distance
	IsocenterToBlockTrayDistance = New(0x300A, 0x00F7)

	// BlockType (300A,00F8) VR=CS VM=1 Block Type
	BlockType = New(0x300A, 0x00F8)

	// AccessoryCode (300A,00F9) VR=LO VM=1 Accessory Code
	AccessoryCode = New(0x300A, 0x00F9)

	// BlockDivergence (300A,00FA) VR=CS VM=1 Block Divergence
	BlockDivergence = New(0x300A, 0x00FA)

	// BlockMountingPosition (300A,00FB) VR=CS VM=1 Block Mounting Position
	BlockMountingPosition = New(0x300A, 0x00FB)

	// BlockNumber (300A,00FC) VR=IS VM=1 Block Number
	BlockNumber = New(0x300A, 0x00FC)

	// BlockName (300A,00FE) VR=LO VM=1 Block Name
	BlockName = New(0x300A, 0x00FE)

	// BlockThickness (300A,0100) VR=DS VM=1 Block Thickness
	BlockThickness = New(0x300A, 0x0100)

	// BlockTransmission (300A,0102) VR=DS VM=1 Block Transmission
	BlockTransmission = New(0x300A, 0x0102)

	// BlockNumberOfPoints (300A,0104) VR=IS VM=1 Block Number of Points
	BlockNumberOfPoints = New(0x300A, 0x0104)

	// BlockData (300A,0106) VR=DS VM=2-2n Block Data
	BlockData = New(0x300A, 0x0106)

	// ApplicatorSequence (300A,0107) VR=SQ VM=1 Applicator Sequence
	ApplicatorSequence = New(0x300A, 0x0107)

	// ApplicatorID (300A,0108) VR=SH VM=1 Applicator ID
	ApplicatorID = New(0x300A, 0x0108)

	// ApplicatorType (300A,0109) VR=CS VM=1 Applicator Type
	ApplicatorType = New(0x300A, 0x0109)

	// ApplicatorDescription (300A,010A) VR=LO VM=1 Applicator Description
	ApplicatorDescription = New(0x300A, 0x010A)

	// CumulativeDoseReferenceCoefficient (300A,010C) VR=DS VM=1 Cumulative Dose Reference Coefficient
	CumulativeDoseReferenceCoefficient = New(0x300A, 0x010C)

	// FinalCumulativeMetersetWeight (300A,010E) VR=DS VM=1 Final Cumulative Meterset Weight
	FinalCumulativeMetersetWeight = New(0x300A, 0x010E)

	// NumberOfControlPoints (300A,0110) VR=IS VM=1 Number of Control Points
	NumberOfControlPoints = New(0x300A, 0x0110)

	// ControlPointSequence (300A,0111) VR=SQ VM=1 Control Point Sequence
	ControlPointSequence = New(0x300A, 0x0111)

	// ControlPointIndex (300A,0112) VR=IS VM=1 Control Point Index
	ControlPointIndex = New(0x300A, 0x0112)

	// NominalBeamEnergy (300A,0114) VR=DS VM=1 Nominal Beam Energy
	NominalBeamEnergy = New(0x300A, 0x0114)

	// DoseRateSet (300A,0115) VR=DS VM=1 Dose Rate Set
	DoseRateSet = New(0x300A, 0x0115)

	// WedgePositionSequence (300A,0116) VR=SQ VM=1 Wedge Position Sequence
	WedgePositionSequence = New(0x300A, 0x0116)

	// WedgePosition (300A,0118) VR=CS VM=1 Wedge Position
	WedgePosition = New(0x300A, 0x0118)

	// BeamLimitingDevicePositionSequence (300A,011A) VR=SQ VM=1 Beam Limiting Device Position Sequence
	BeamLimitingDevicePositionSequence = New(0x300A, 0x011A)

	// LeafJawPositions (300A,011C) VR=DS VM=2-2n Leaf/Jaw Positions
	LeafJawPositions = New(0x300A, 0x011C)

	// GantryAngle (300A,011E) VR=DS VM=1 Gantry Angle
	GantryAngle = New(0x300A, 0x011E)

	// GantryRotationDirection (300A,011F) VR=CS VM=1 Gantry Rotation Direction
	GantryRotationDirection = New(0x300A, 0x011F)

	// BeamLimitingDeviceAngle (300A,0120) VR=DS VM=1 Beam Limiting Device Angle
	BeamLimitingDeviceAngle = New(0x300A, 0x0120)

	// BeamLimitingDeviceRotationDirection (300A,0121) VR=CS VM=1 Beam Limiting Device Rotation Direction
	BeamLimitingDeviceRotationDirection = New(0x300A, 0x0121)

	// PatientSupportAngle (300A,0122) VR=DS VM=1 Patient Support Angle
	PatientSupportAngle = New(0x300A, 0x0122)

	// PatientSupportRotationDirection (300A,0123) VR=CS VM=1 Patient Support Rotation Direction
	PatientSupportRotationDirection = New(0x300A, 0x0123)

	// TableTopEccentricAxisDistance (300A,0124) VR=DS VM=1 Table Top Eccentric Axis Distance
	TableTopEccentricAxisDistance = New(0x300A, 0x0124)

	// TableTopEccentricAngle (300A,0125) VR=DS VM=1 Table Top Eccentric Angle
	TableTopEccentricAngle = New(0x300A, 0x0125)

	// TableTopEccentricRotationDirection (300A,0126) VR=CS VM=1 Table Top Eccentric Rotation Direction
	TableTopEccentricRotationDirection = New(0x300A, 0x0126)

	// TableTopVerticalPosition (300A,0128) VR=DS VM=1 Table Top Vertical Position
	TableTopVerticalPosition = New(0x300A, 0x0128)

	// TableTopLongitudinalPosition (300A,0129) VR=DS VM=1 Table Top Longitudinal Position
	TableTopLongitudinalPosition = New(0x300A, 0x0129)

	// TableTopLateralPosition (300A,012A) VR=DS VM=1 Table Top Lateral Position
	TableTopLateralPosition = New(0x300A, 0x012A)

	// IsocenterPosition (300A,012C) VR=DS VM=3 Isocenter Position
	IsocenterPosition = New(0x300A, 0x012C)

	// SurfaceEntryPoint (300A,012E) VR=DS VM=3 Surface Entry Point
	SurfaceEntryPoint = New(0x300A, 0x012E)

	// SourceToSurfaceDistance (300A,0130) VR=DS VM=1 Source to Surface Distance
	SourceToSurfaceDistance = New(0x300A, 0x0130)

	// AverageBeamDosePointSourceToExternalContourDistance (300A,0131) VR=FL VM=1 Average Beam Dose Point Source to External Contour Distance
	AverageBeamDosePointSourceToExternalContourDistance = New(0x300A, 0x0131)

	// SourceToExternalContourDistance (300A,0132) VR=FL VM=1 Source to External Contour Distance
	SourceToExternalContourDistance = New(0x300A, 0x0132)

	// ExternalContourEntryPoint (300A,0133) VR=FL VM=3 External Contour Entry Point
	ExternalContourEntryPoint = New(0x300A, 0x0133)

	// CumulativeMetersetWeight (300A,0134) VR=DS VM=1 Cumulative Meterset Weight
	CumulativeMetersetWeight = New(0x300A, 0x0134)

	// TableTopPitchAngle (300A,0140) VR=FL VM=1 Table Top Pitch Angle
	TableTopPitchAngle = New(0x300A, 0x0140)

	// TableTopPitchRotationDirection (300A,0142) VR=CS VM=1 Table Top Pitch Rotation Direction
	TableTopPitchRotationDirection = New(0x300A, 0x0142)

	// TableTopRollAngle (300A,0144) VR=FL VM=1 Table Top Roll Angle
	TableTopRollAngle = New(0x300A, 0x0144)

	// TableTopRollRotationDirection (300A,0146) VR=CS VM=1 Table Top Roll Rotation Direction
	TableTopRollRotationDirection = New(0x300A, 0x0146)

	// HeadFixationAngle (300A,0148) VR=FL VM=1 Head Fixation Angle
	HeadFixationAngle = New(0x300A, 0x0148)

	// GantryPitchAngle (300A,014A) VR=FL VM=1 Gantry Pitch Angle
	GantryPitchAngle = New(0x300A, 0x014A)

	// GantryPitchRotationDirection (300A,014C) VR=CS VM=1 Gantry Pitch Rotation Direction
	GantryPitchRotationDirection = New(0x300A, 0x014C)

	// GantryPitchAngleTolerance (300A,014E) VR=FL VM=1 Gantry Pitch Angle Tolerance
	GantryPitchAngleTolerance = New(0x300A, 0x014E)

	// FixationEye (300A,0150) VR=CS VM=1 Fixation Eye
	FixationEye = New(0x300A, 0x0150)

	// ChairHeadFramePosition (300A,0151) VR=DS VM=1 Chair Head Frame Position
	ChairHeadFramePosition = New(0x300A, 0x0151)

	// HeadFixationAngleTolerance (300A,0152) VR=DS VM=1 Head Fixation Angle Tolerance
	HeadFixationAngleTolerance = New(0x300A, 0x0152)

	// ChairHeadFramePositionTolerance (300A,0153) VR=DS VM=1 Chair Head Frame Position Tolerance
	ChairHeadFramePositionTolerance = New(0x300A, 0x0153)

	// FixationLightAzimuthalAngleTolerance (300A,0154) VR=DS VM=1 Fixation Light Azimuthal Angle Tolerance
	FixationLightAzimuthalAngleTolerance = New(0x300A, 0x0154)

	// FixationLightPolarAngleTolerance (300A,0155) VR=DS VM=1 Fixation Light Polar Angle Tolerance
	FixationLightPolarAngleTolerance = New(0x300A, 0x0155)

	// PatientSetupSequence (300A,0180) VR=SQ VM=1 Patient Setup Sequence
	PatientSetupSequence = New(0x300A, 0x0180)

	// PatientSetupNumber (300A,0182) VR=IS VM=1 Patient Setup Number
	PatientSetupNumber = New(0x300A, 0x0182)

	// PatientSetupLabel (300A,0183) VR=LO VM=1 Patient Setup Label
	PatientSetupLabel = New(0x300A, 0x0183)

	// PatientAdditionalPosition (300A,0184) VR=LO VM=1 Patient Additional Position
	PatientAdditionalPosition = New(0x300A, 0x0184)

	// FixationDeviceSequence (300A,0190) VR=SQ VM=1 Fixation Device Sequence
	FixationDeviceSequence = New(0x300A, 0x0190)

	// FixationDeviceType (300A,0192) VR=CS VM=1 Fixation Device Type
	FixationDeviceType = New(0x300A, 0x0192)

	// FixationDeviceLabel (300A,0194) VR=SH VM=1 Fixation Device Label
	FixationDeviceLabel = New(0x300A, 0x0194)

	// FixationDeviceDescription (300A,0196) VR=ST VM=1 Fixation Device Description
	FixationDeviceDescription = New(0x300A, 0x0196)

	// FixationDevicePosition (300A,0198) VR=SH VM=1 Fixation Device Position
	FixationDevicePosition = New(0x300A, 0x0198)

	// FixationDevicePitchAngle (300A,0199) VR=FL VM=1 Fixation Device Pitch Angle
	FixationDevicePitchAngle = New(0x300A, 0x0199)

	// FixationDeviceRollAngle (300A,019A) VR=FL VM=1 Fixation Device Roll Angle
	FixationDeviceRollAngle = New(0x300A, 0x019A)

	// ShieldingDeviceSequence (300A,01A0) VR=SQ VM=1 Shielding Device Sequence
	ShieldingDeviceSequence = New(0x300A, 0x01A0)

	// ShieldingDeviceType (300A,01A2) VR=CS VM=1 Shielding Device Type
	ShieldingDeviceType = New(0x300A, 0x01A2)

	// ShieldingDeviceLabel (300A,01A4) VR=SH VM=1 Shielding Device Label
	ShieldingDeviceLabel = New(0x300A, 0x01A4)

	// ShieldingDeviceDescription (300A,01A6) VR=ST VM=1 Shielding Device Description
	ShieldingDeviceDescription = New(0x300A, 0x01A6)

	// ShieldingDevicePosition (300A,01A8) VR=SH VM=1 Shielding Device Position
	ShieldingDevicePosition = New(0x300A, 0x01A8)

	// SetupTechnique (300A,01B0) VR=CS VM=1 Setup Technique
	SetupTechnique = New(0x300A, 0x01B0)

	// SetupTechniqueDescription (300A,01B2) VR=ST VM=1 Setup Technique Description
	SetupTechniqueDescription = New(0x300A, 0x01B2)

	// SetupDeviceSequence (300A,01B4) VR=SQ VM=1 Setup Device Sequence
	SetupDeviceSequence = New(0x300A, 0x01B4)

	// SetupDeviceType (300A,01B6) VR=CS VM=1 Setup Device Type
	SetupDeviceType = New(0x300A, 0x01B6)

	// SetupDeviceLabel (300A,01B8) VR=SH VM=1 Setup Device Label
	SetupDeviceLabel = New(0x300A, 0x01B8)

	// SetupDeviceDescription (300A,01BA) VR=ST VM=1 Setup Device Description
	SetupDeviceDescription = New(0x300A, 0x01BA)

	// SetupDeviceParameter (300A,01BC) VR=DS VM=1 Setup Device Parameter
	SetupDeviceParameter = New(0x300A, 0x01BC)

	// SetupReferenceDescription (300A,01D0) VR=ST VM=1 Setup Reference Description
	SetupReferenceDescription = New(0x300A, 0x01D0)

	// TableTopVerticalSetupDisplacement (300A,01D2) VR=DS VM=1 Table Top Vertical Setup Displacement
	TableTopVerticalSetupDisplacement = New(0x300A, 0x01D2)

	// TableTopLongitudinalSetupDisplacement (300A,01D4) VR=DS VM=1 Table Top Longitudinal Setup Displacement
	TableTopLongitudinalSetupDisplacement = New(0x300A, 0x01D4)

	// TableTopLateralSetupDisplacement (300A,01D6) VR=DS VM=1 Table Top Lateral Setup Displacement
	TableTopLateralSetupDisplacement = New(0x300A, 0x01D6)

	// BrachyTreatmentTechnique (300A,0200) VR=CS VM=1 Brachy Treatment Technique
	BrachyTreatmentTechnique = New(0x300A, 0x0200)

	// BrachyTreatmentType (300A,0202) VR=CS VM=1 Brachy Treatment Type
	BrachyTreatmentType = New(0x300A, 0x0202)

	// TreatmentMachineSequence (300A,0206) VR=SQ VM=1 Treatment Machine Sequence
	TreatmentMachineSequence = New(0x300A, 0x0206)

	// SourceSequence (300A,0210) VR=SQ VM=1 Source Sequence
	SourceSequence = New(0x300A, 0x0210)

	// SourceNumber (300A,0212) VR=IS VM=1 Source Number
	SourceNumber = New(0x300A, 0x0212)

	// SourceType (300A,0214) VR=CS VM=1 Source Type
	SourceType = New(0x300A, 0x0214)

	// SourceManufacturer (300A,0216) VR=LO VM=1 Source Manufacturer
	SourceManufacturer = New(0x300A, 0x0216)

	// ActiveSourceDiameter (300A,0218) VR=DS VM=1 Active Source Diameter
	ActiveSourceDiameter = New(0x300A, 0x0218)

	// ActiveSourceLength (300A,021A) VR=DS VM=1 Active Source Length
	ActiveSourceLength = New(0x300A, 0x021A)

	// SourceModelID (300A,021B) VR=SH VM=1 Source Model ID
	SourceModelID = New(0x300A, 0x021B)

	// SourceDescription (300A,021C) VR=LO VM=1 Source Description
	SourceDescription = New(0x300A, 0x021C)

	// SourceEncapsulationNominalThickness (300A,0222) VR=DS VM=1 Source Encapsulation Nominal Thickness
	SourceEncapsulationNominalThickness = New(0x300A, 0x0222)

	// SourceEncapsulationNominalTransmission (300A,0224) VR=DS VM=1 Source Encapsulation Nominal Transmission
	SourceEncapsulationNominalTransmission = New(0x300A, 0x0224)

	// SourceIsotopeName (300A,0226) VR=LO VM=1 Source Isotope Name
	SourceIsotopeName = New(0x300A, 0x0226)

	// SourceIsotopeHalfLife (300A,0228) VR=DS VM=1 Source Isotope Half Life
	SourceIsotopeHalfLife = New(0x300A, 0x0228)

	// SourceStrengthUnits (300A,0229) VR=CS VM=1 Source Strength Units
	SourceStrengthUnits = New(0x300A, 0x0229)

	// ReferenceAirKermaRate (300A,022A) VR=DS VM=1 Reference Air Kerma Rate
	ReferenceAirKermaRate = New(0x300A, 0x022A)

	// SourceStrength (300A,022B) VR=DS VM=1 Source Strength
	SourceStrength = New(0x300A, 0x022B)

	// SourceStrengthReferenceDate (300A,022C) VR=DA VM=1 Source Strength Reference Date
	SourceStrengthReferenceDate = New(0x300A, 0x022C)

	// SourceStrengthReferenceTime (300A,022E) VR=TM VM=1 Source Strength Reference Time
	SourceStrengthReferenceTime = New(0x300A, 0x022E)

	// ApplicationSetupSequence (300A,0230) VR=SQ VM=1 Application Setup Sequence
	ApplicationSetupSequence = New(0x300A, 0x0230)

	// ApplicationSetupType (300A,0232) VR=CS VM=1 Application Setup Type
	ApplicationSetupType = New(0x300A, 0x0232)

	// ApplicationSetupNumber (300A,0234) VR=IS VM=1 Application Setup Number
	ApplicationSetupNumber = New(0x300A, 0x0234)

	// ApplicationSetupName (300A,0236) VR=LO VM=1 Application Setup Name
	ApplicationSetupName = New(0x300A, 0x0236)

	// ApplicationSetupManufacturer (300A,0238) VR=LO VM=1 Application Setup Manufacturer
	ApplicationSetupManufacturer = New(0x300A, 0x0238)

	// TemplateNumber (300A,0240) VR=IS VM=1 Template Number
	TemplateNumber = New(0x300A, 0x0240)

	// TemplateType (300A,0242) VR=SH VM=1 Template Type
	TemplateType = New(0x300A, 0x0242)

	// TemplateName (300A,0244) VR=LO VM=1 Template Name
	TemplateName = New(0x300A, 0x0244)

	// TotalReferenceAirKerma (300A,0250) VR=DS VM=1 Total Reference Air Kerma
	TotalReferenceAirKerma = New(0x300A, 0x0250)

	// BrachyAccessoryDeviceSequence (300A,0260) VR=SQ VM=1 Brachy Accessory Device Sequence
	BrachyAccessoryDeviceSequence = New(0x300A, 0x0260)

	// BrachyAccessoryDeviceNumber (300A,0262) VR=IS VM=1 Brachy Accessory Device Number
	BrachyAccessoryDeviceNumber = New(0x300A, 0x0262)

	// BrachyAccessoryDeviceID (300A,0263) VR=SH VM=1 Brachy Accessory Device ID
	BrachyAccessoryDeviceID = New(0x300A, 0x0263)

	// BrachyAccessoryDeviceType (300A,0264) VR=CS VM=1 Brachy Accessory Device Type
	BrachyAccessoryDeviceType = New(0x300A, 0x0264)

	// BrachyAccessoryDeviceName (300A,0266) VR=LO VM=1 Brachy Accessory Device Name
	BrachyAccessoryDeviceName = New(0x300A, 0x0266)

	// BrachyAccessoryDeviceNominalThickness (300A,026A) VR=DS VM=1 Brachy Accessory Device Nominal Thickness
	BrachyAccessoryDeviceNominalThickness = New(0x300A, 0x026A)

	// BrachyAccessoryDeviceNominalTransmission (300A,026C) VR=DS VM=1 Brachy Accessory Device Nominal Transmission
	BrachyAccessoryDeviceNominalTransmission = New(0x300A, 0x026C)

	// ChannelEffectiveLength (300A,0271) VR=DS VM=1 Channel Effective Length
	ChannelEffectiveLength = New(0x300A, 0x0271)

	// ChannelInnerLength (300A,0272) VR=DS VM=1 Channel Inner Length
	ChannelInnerLength = New(0x300A, 0x0272)

	// AfterloaderChannelID (300A,0273) VR=SH VM=1 Afterloader Channel ID
	AfterloaderChannelID = New(0x300A, 0x0273)

	// SourceApplicatorTipLength (300A,0274) VR=DS VM=1 Source Applicator Tip Length
	SourceApplicatorTipLength = New(0x300A, 0x0274)

	// ChannelSequence (300A,0280) VR=SQ VM=1 Channel Sequence
	ChannelSequence = New(0x300A, 0x0280)

	// ChannelNumber (300A,0282) VR=IS VM=1 Channel Number
	ChannelNumber = New(0x300A, 0x0282)

	// ChannelLength (300A,0284) VR=DS VM=1 Channel Length
	ChannelLength = New(0x300A, 0x0284)

	// ChannelTotalTime (300A,0286) VR=DS VM=1 Channel Total Time
	ChannelTotalTime = New(0x300A, 0x0286)

	// SourceMovementType (300A,0288) VR=CS VM=1 Source Movement Type
	SourceMovementType = New(0x300A, 0x0288)

	// NumberOfPulses (300A,028A) VR=IS VM=1 Number of Pulses
	NumberOfPulses = New(0x300A, 0x028A)

	// PulseRepetitionInterval (300A,028C) VR=DS VM=1 Pulse Repetition Interval
	PulseRepetitionInterval = New(0x300A, 0x028C)

	// SourceApplicatorNumber (300A,0290) VR=IS VM=1 Source Applicator Number
	SourceApplicatorNumber = New(0x300A, 0x0290)

	// SourceApplicatorID (300A,0291) VR=SH VM=1 Source Applicator ID
	SourceApplicatorID = New(0x300A, 0x0291)

	// SourceApplicatorType (300A,0292) VR=CS VM=1 Source Applicator Type
	SourceApplicatorType = New(0x300A, 0x0292)

	// SourceApplicatorName (300A,0294) VR=LO VM=1 Source Applicator Name
	SourceApplicatorName = New(0x300A, 0x0294)

	// SourceApplicatorLength (300A,0296) VR=DS VM=1 Source Applicator Length
	SourceApplicatorLength = New(0x300A, 0x0296)

	// SourceApplicatorManufacturer (300A,0298) VR=LO VM=1 Source Applicator Manufacturer
	SourceApplicatorManufacturer = New(0x300A, 0x0298)

	// SourceApplicatorWallNominalThickness (300A,029C) VR=DS VM=1 Source Applicator Wall Nominal Thickness
	SourceApplicatorWallNominalThickness = New(0x300A, 0x029C)

	// SourceApplicatorWallNominalTransmission (300A,029E) VR=DS VM=1 Source Applicator Wall Nominal Transmission
	SourceApplicatorWallNominalTransmission = New(0x300A, 0x029E)

	// SourceApplicatorStepSize (300A,02A0) VR=DS VM=1 Source Applicator Step Size
	SourceApplicatorStepSize = New(0x300A, 0x02A0)

	// ApplicatorShapeReferencedROINumber (300A,02A1) VR=IS VM=1 Applicator Shape Referenced ROI Number
	ApplicatorShapeReferencedROINumber = New(0x300A, 0x02A1)

	// TransferTubeNumber (300A,02A2) VR=IS VM=1 Transfer Tube Number
	TransferTubeNumber = New(0x300A, 0x02A2)

	// TransferTubeLength (300A,02A4) VR=DS VM=1 Transfer Tube Length
	TransferTubeLength = New(0x300A, 0x02A4)

	// ChannelShieldSequence (300A,02B0) VR=SQ VM=1 Channel Shield Sequence
	ChannelShieldSequence = New(0x300A, 0x02B0)

	// ChannelShieldNumber (300A,02B2) VR=IS VM=1 Channel Shield Number
	ChannelShieldNumber = New(0x300A, 0x02B2)

	// ChannelShieldID (300A,02B3) VR=SH VM=1 Channel Shield ID
	ChannelShieldID = New(0x300A, 0x02B3)

	// ChannelShieldName (300A,02B4) VR=LO VM=1 Channel Shield Name
	ChannelShieldName = New(0x300A, 0x02B4)

	// ChannelShieldNominalThickness (300A,02B8) VR=DS VM=1 Channel Shield Nominal Thickness
	ChannelShieldNominalThickness = New(0x300A, 0x02B8)

	// ChannelShieldNominalTransmission (300A,02BA) VR=DS VM=1 Channel Shield Nominal Transmission
	ChannelShieldNominalTransmission = New(0x300A, 0x02BA)

	// FinalCumulativeTimeWeight (300A,02C8) VR=DS VM=1 Final Cumulative Time Weight
	FinalCumulativeTimeWeight = New(0x300A, 0x02C8)

	// BrachyControlPointSequence (300A,02D0) VR=SQ VM=1 Brachy Control Point Sequence
	BrachyControlPointSequence = New(0x300A, 0x02D0)

	// ControlPointRelativePosition (300A,02D2) VR=DS VM=1 Control Point Relative Position
	ControlPointRelativePosition = New(0x300A, 0x02D2)

	// ControlPoint3DPosition (300A,02D4) VR=DS VM=3 Control Point 3D Position
	ControlPoint3DPosition = New(0x300A, 0x02D4)

	// CumulativeTimeWeight (300A,02D6) VR=DS VM=1 Cumulative Time Weight
	CumulativeTimeWeight = New(0x300A, 0x02D6)

	// CompensatorDivergence (300A,02E0) VR=CS VM=1 Compensator Divergence
	CompensatorDivergence = New(0x300A, 0x02E0)

	// CompensatorMountingPosition (300A,02E1) VR=CS VM=1 Compensator Mounting Position
	CompensatorMountingPosition = New(0x300A, 0x02E1)

	// SourceToCompensatorDistance (300A,02E2) VR=DS VM=1-n Source to Compensator Distance
	SourceToCompensatorDistance = New(0x300A, 0x02E2)

	// TotalCompensatorTrayWaterEquivalentThickness (300A,02E3) VR=FL VM=1 Total Compensator Tray Water-Equivalent Thickness
	TotalCompensatorTrayWaterEquivalentThickness = New(0x300A, 0x02E3)

	// IsocenterToCompensatorTrayDistance (300A,02E4) VR=FL VM=1 Isocenter to Compensator Tray Distance
	IsocenterToCompensatorTrayDistance = New(0x300A, 0x02E4)

	// CompensatorColumnOffset (300A,02E5) VR=FL VM=1 Compensator Column Offset
	CompensatorColumnOffset = New(0x300A, 0x02E5)

	// IsocenterToCompensatorDistances (300A,02E6) VR=FL VM=1-n Isocenter to Compensator Distances
	IsocenterToCompensatorDistances = New(0x300A, 0x02E6)

	// CompensatorRelativeStoppingPowerRatio (300A,02E7) VR=FL VM=1 Compensator Relative Stopping Power Ratio
	CompensatorRelativeStoppingPowerRatio = New(0x300A, 0x02E7)

	// CompensatorMillingToolDiameter (300A,02E8) VR=FL VM=1 Compensator Milling Tool Diameter
	CompensatorMillingToolDiameter = New(0x300A, 0x02E8)

	// IonRangeCompensatorSequence (300A,02EA) VR=SQ VM=1 Ion Range Compensator Sequence
	IonRangeCompensatorSequence = New(0x300A, 0x02EA)

	// CompensatorDescription (300A,02EB) VR=LT VM=1 Compensator Description
	CompensatorDescription = New(0x300A, 0x02EB)

	// CompensatorSurfaceRepresentationFlag (300A,02EC) VR=CS VM=1 Compensator Surface Representation Flag
	CompensatorSurfaceRepresentationFlag = New(0x300A, 0x02EC)

	// RadiationMassNumber (300A,0302) VR=IS VM=1 Radiation Mass Number
	RadiationMassNumber = New(0x300A, 0x0302)

	// RadiationAtomicNumber (300A,0304) VR=IS VM=1 Radiation Atomic Number
	RadiationAtomicNumber = New(0x300A, 0x0304)

	// RadiationChargeState (300A,0306) VR=SS VM=1 Radiation Charge State
	RadiationChargeState = New(0x300A, 0x0306)

	// ScanMode (300A,0308) VR=CS VM=1 Scan Mode
	ScanMode = New(0x300A, 0x0308)

	// ModulatedScanModeType (300A,0309) VR=CS VM=1 Modulated Scan Mode Type
	ModulatedScanModeType = New(0x300A, 0x0309)

	// VirtualSourceAxisDistances (300A,030A) VR=FL VM=2 Virtual Source-Axis Distances
	VirtualSourceAxisDistances = New(0x300A, 0x030A)

	// SnoutSequence (300A,030C) VR=SQ VM=1 Snout Sequence
	SnoutSequence = New(0x300A, 0x030C)

	// SnoutPosition (300A,030D) VR=FL VM=1 Snout Position
	SnoutPosition = New(0x300A, 0x030D)

	// SnoutID (300A,030F) VR=SH VM=1 Snout ID
	SnoutID = New(0x300A, 0x030F)

	// NumberOfRangeShifters (300A,0312) VR=IS VM=1 Number of Range Shifters
	NumberOfRangeShifters = New(0x300A, 0x0312)

	// RangeShifterSequence (300A,0314) VR=SQ VM=1 Range Shifter Sequence
	RangeShifterSequence = New(0x300A, 0x0314)

	// RangeShifterNumber (300A,0316) VR=IS VM=1 Range Shifter Number
	RangeShifterNumber = New(0x300A, 0x0316)

	// RangeShifterID (300A,0318) VR=SH VM=1 Range Shifter ID
	RangeShifterID = New(0x300A, 0x0318)

	// RangeShifterType (300A,0320) VR=CS VM=1 Range Shifter Type
	RangeShifterType = New(0x300A, 0x0320)

	// RangeShifterDescription (300A,0322) VR=LO VM=1 Range Shifter Description
	RangeShifterDescription = New(0x300A, 0x0322)

	// NumberOfLateralSpreadingDevices (300A,0330) VR=IS VM=1 Number of Lateral Spreading Devices
	NumberOfLateralSpreadingDevices = New(0x300A, 0x0330)

	// LateralSpreadingDeviceSequence (300A,0332) VR=SQ VM=1 Lateral Spreading Device Sequence
	LateralSpreadingDeviceSequence = New(0x300A, 0x0332)

	// LateralSpreadingDeviceNumber (300A,0334) VR=IS VM=1 Lateral Spreading Device Number
	LateralSpreadingDeviceNumber = New(0x300A, 0x0334)

	// LateralSpreadingDeviceID (300A,0336) VR=SH VM=1 Lateral Spreading Device ID
	LateralSpreadingDeviceID = New(0x300A, 0x0336)

	// LateralSpreadingDeviceType (300A,0338) VR=CS VM=1 Lateral Spreading Device Type
	LateralSpreadingDeviceType = New(0x300A, 0x0338)

	// LateralSpreadingDeviceDescription (300A,033A) VR=LO VM=1 Lateral Spreading Device Description
	LateralSpreadingDeviceDescription = New(0x300A, 0x033A)

	// LateralSpreadingDeviceWaterEquivalentThickness (300A,033C) VR=FL VM=1 Lateral Spreading Device Water Equivalent Thickness
	LateralSpreadingDeviceWaterEquivalentThickness = New(0x300A, 0x033C)

	// NumberOfRangeModulators (300A,0340) VR=IS VM=1 Number of Range Modulators
	NumberOfRangeModulators = New(0x300A, 0x0340)

	// RangeModulatorSequence (300A,0342) VR=SQ VM=1 Range Modulator Sequence
	RangeModulatorSequence = New(0x300A, 0x0342)

	// RangeModulatorNumber (300A,0344) VR=IS VM=1 Range Modulator Number
	RangeModulatorNumber = New(0x300A, 0x0344)

	// RangeModulatorID (300A,0346) VR=SH VM=1 Range Modulator ID
	RangeModulatorID = New(0x300A, 0x0346)

	// RangeModulatorType (300A,0348) VR=CS VM=1 Range Modulator Type
	RangeModulatorType = New(0x300A, 0x0348)

	// RangeModulatorDescription (300A,034A) VR=LO VM=1 Range Modulator Description
	RangeModulatorDescription = New(0x300A, 0x034A)

	// BeamCurrentModulationID (300A,034C) VR=SH VM=1 Beam Current Modulation ID
	BeamCurrentModulationID = New(0x300A, 0x034C)

	// PatientSupportType (300A,0350) VR=CS VM=1 Patient Support Type
	PatientSupportType = New(0x300A, 0x0350)

	// PatientSupportID (300A,0352) VR=SH VM=1 Patient Support ID
	PatientSupportID = New(0x300A, 0x0352)

	// PatientSupportAccessoryCode (300A,0354) VR=LO VM=1 Patient Support Accessory Code
	PatientSupportAccessoryCode = New(0x300A, 0x0354)

	// TrayAccessoryCode (300A,0355) VR=LO VM=1 Tray Accessory Code
	TrayAccessoryCode = New(0x300A, 0x0355)

	// FixationLightAzimuthalAngle (300A,0356) VR=FL VM=1 Fixation Light Azimuthal Angle
	FixationLightAzimuthalAngle = New(0x300A, 0x0356)

	// FixationLightPolarAngle (300A,0358) VR=FL VM=1 Fixation Light Polar Angle
	FixationLightPolarAngle = New(0x300A, 0x0358)

	// MetersetRate (300A,035A) VR=FL VM=1 Meterset Rate
	MetersetRate = New(0x300A, 0x035A)

	// RangeShifterSettingsSequence (300A,0360) VR=SQ VM=1 Range Shifter Settings Sequence
	RangeShifterSettingsSequence = New(0x300A, 0x0360)

	// RangeShifterSetting (300A,0362) VR=LO VM=1 Range Shifter Setting
	RangeShifterSetting = New(0x300A, 0x0362)

	// IsocenterToRangeShifterDistance (300A,0364) VR=FL VM=1 Isocenter to Range Shifter Distance
	IsocenterToRangeShifterDistance = New(0x300A, 0x0364)

	// RangeShifterWaterEquivalentThickness (300A,0366) VR=FL VM=1 Range Shifter Water Equivalent Thickness
	RangeShifterWaterEquivalentThickness = New(0x300A, 0x0366)

	// LateralSpreadingDeviceSettingsSequence (300A,0370) VR=SQ VM=1 Lateral Spreading Device Settings Sequence
	LateralSpreadingDeviceSettingsSequence = New(0x300A, 0x0370)

	// LateralSpreadingDeviceSetting (300A,0372) VR=LO VM=1 Lateral Spreading Device Setting
	LateralSpreadingDeviceSetting = New(0x300A, 0x0372)

	// IsocenterToLateralSpreadingDeviceDistance (300A,0374) VR=FL VM=1 Isocenter to Lateral Spreading Device Distance
	IsocenterToLateralSpreadingDeviceDistance = New(0x300A, 0x0374)

	// RangeModulatorSettingsSequence (300A,0380) VR=SQ VM=1 Range Modulator Settings Sequence
	RangeModulatorSettingsSequence = New(0x300A, 0x0380)

	// RangeModulatorGatingStartValue (300A,0382) VR=FL VM=1 Range Modulator Gating Start Value
	RangeModulatorGatingStartValue = New(0x300A, 0x0382)

	// RangeModulatorGatingStopValue (300A,0384) VR=FL VM=1 Range Modulator Gating Stop Value
	RangeModulatorGatingStopValue = New(0x300A, 0x0384)

	// RangeModulatorGatingStartWaterEquivalentThickness (300A,0386) VR=FL VM=1 Range Modulator Gating Start Water Equivalent Thickness
	RangeModulatorGatingStartWaterEquivalentThickness = New(0x300A, 0x0386)

	// RangeModulatorGatingStopWaterEquivalentThickness (300A,0388) VR=FL VM=1 Range Modulator Gating Stop Water Equivalent Thickness
	RangeModulatorGatingStopWaterEquivalentThickness = New(0x300A, 0x0388)

	// IsocenterToRangeModulatorDistance (300A,038A) VR=FL VM=1 Isocenter to Range Modulator Distance
	IsocenterToRangeModulatorDistance = New(0x300A, 0x038A)

	// ScanSpotTimeOffset (300A,038F) VR=FL VM=1-n Scan Spot Time Offset
	ScanSpotTimeOffset = New(0x300A, 0x038F)

	// ScanSpotTuneID (300A,0390) VR=SH VM=1 Scan Spot Tune ID
	ScanSpotTuneID = New(0x300A, 0x0390)

	// ScanSpotPrescribedIndices (300A,0391) VR=IS VM=1-n Scan Spot Prescribed Indices
	ScanSpotPrescribedIndices = New(0x300A, 0x0391)

	// NumberOfScanSpotPositions (300A,0392) VR=IS VM=1 Number of Scan Spot Positions
	NumberOfScanSpotPositions = New(0x300A, 0x0392)

	// ScanSpotReordered (300A,0393) VR=CS VM=1 Scan Spot Reordered
	ScanSpotReordered = New(0x300A, 0x0393)

	// ScanSpotPositionMap (300A,0394) VR=FL VM=1-n Scan Spot Position Map
	ScanSpotPositionMap = New(0x300A, 0x0394)

	// ScanSpotReorderingAllowed (300A,0395) VR=CS VM=1 Scan Spot Reordering Allowed
	ScanSpotReorderingAllowed = New(0x300A, 0x0395)

	// ScanSpotMetersetWeights (300A,0396) VR=FL VM=1-n Scan Spot Meterset Weights
	ScanSpotMetersetWeights = New(0x300A, 0x0396)

	// ScanningSpotSize (300A,0398) VR=FL VM=2 Scanning Spot Size
	ScanningSpotSize = New(0x300A, 0x0398)

	// ScanSpotSizesDelivered (300A,0399) VR=FL VM=2-2n Scan Spot Sizes Delivered
	ScanSpotSizesDelivered = New(0x300A, 0x0399)

	// NumberOfPaintings (300A,039A) VR=IS VM=1 Number of Paintings
	NumberOfPaintings = New(0x300A, 0x039A)

	// ScanSpotGantryAngles (300A,039B) VR=FL VM=1-n Scan Spot Gantry Angles
	ScanSpotGantryAngles = New(0x300A, 0x039B)

	// ScanSpotPatientSupportAngles (300A,039C) VR=FL VM=1-n Scan Spot Patient Support Angles
	ScanSpotPatientSupportAngles = New(0x300A, 0x039C)

	// IonToleranceTableSequence (300A,03A0) VR=SQ VM=1 Ion Tolerance Table Sequence
	IonToleranceTableSequence = New(0x300A, 0x03A0)

	// IonBeamSequence (300A,03A2) VR=SQ VM=1 Ion Beam Sequence
	IonBeamSequence = New(0x300A, 0x03A2)

	// IonBeamLimitingDeviceSequence (300A,03A4) VR=SQ VM=1 Ion Beam Limiting Device Sequence
	IonBeamLimitingDeviceSequence = New(0x300A, 0x03A4)

	// IonBlockSequence (300A,03A6) VR=SQ VM=1 Ion Block Sequence
	IonBlockSequence = New(0x300A, 0x03A6)

	// IonControlPointSequence (300A,03A8) VR=SQ VM=1 Ion Control Point Sequence
	IonControlPointSequence = New(0x300A, 0x03A8)

	// IonWedgeSequence (300A,03AA) VR=SQ VM=1 Ion Wedge Sequence
	IonWedgeSequence = New(0x300A, 0x03AA)

	// IonWedgePositionSequence (300A,03AC) VR=SQ VM=1 Ion Wedge Position Sequence
	IonWedgePositionSequence = New(0x300A, 0x03AC)

	// ReferencedSetupImageSequence (300A,0401) VR=SQ VM=1 Referenced Setup Image Sequence
	ReferencedSetupImageSequence = New(0x300A, 0x0401)

	// SetupImageComment (300A,0402) VR=ST VM=1 Setup Image Comment
	SetupImageComment = New(0x300A, 0x0402)

	// MotionSynchronizationSequence (300A,0410) VR=SQ VM=1 Motion Synchronization Sequence
	MotionSynchronizationSequence = New(0x300A, 0x0410)

	// ControlPointOrientation (300A,0412) VR=FL VM=3 Control Point Orientation
	ControlPointOrientation = New(0x300A, 0x0412)

	// GeneralAccessorySequence (300A,0420) VR=SQ VM=1 General Accessory Sequence
	GeneralAccessorySequence = New(0x300A, 0x0420)

	// GeneralAccessoryID (300A,0421) VR=SH VM=1 General Accessory ID
	GeneralAccessoryID = New(0x300A, 0x0421)

	// GeneralAccessoryDescription (300A,0422) VR=ST VM=1 General Accessory Description
	GeneralAccessoryDescription = New(0x300A, 0x0422)

	// GeneralAccessoryType (300A,0423) VR=CS VM=1 General Accessory Type
	GeneralAccessoryType = New(0x300A, 0x0423)

	// GeneralAccessoryNumber (300A,0424) VR=IS VM=1 General Accessory Number
	GeneralAccessoryNumber = New(0x300A, 0x0424)

	// SourceToGeneralAccessoryDistance (300A,0425) VR=FL VM=1 Source to General Accessory Distance
	SourceToGeneralAccessoryDistance = New(0x300A, 0x0425)

	// IsocenterToGeneralAccessoryDistance (300A,0426) VR=DS VM=1 Isocenter to General Accessory Distance
	IsocenterToGeneralAccessoryDistance = New(0x300A, 0x0426)

	// ApplicatorGeometrySequence (300A,0431) VR=SQ VM=1 Applicator Geometry Sequence
	ApplicatorGeometrySequence = New(0x300A, 0x0431)

	// ApplicatorApertureShape (300A,0432) VR=CS VM=1 Applicator Aperture Shape
	ApplicatorApertureShape = New(0x300A, 0x0432)

	// ApplicatorOpening (300A,0433) VR=FL VM=1 Applicator Opening
	ApplicatorOpening = New(0x300A, 0x0433)

	// ApplicatorOpeningX (300A,0434) VR=FL VM=1 Applicator Opening X
	ApplicatorOpeningX = New(0x300A, 0x0434)

	// ApplicatorOpeningY (300A,0435) VR=FL VM=1 Applicator Opening Y
	ApplicatorOpeningY = New(0x300A, 0x0435)

	// SourceToApplicatorMountingPositionDistance (300A,0436) VR=FL VM=1 Source to Applicator Mounting Position Distance
	SourceToApplicatorMountingPositionDistance = New(0x300A, 0x0436)

	// NumberOfBlockSlabItems (300A,0440) VR=IS VM=1 Number of Block Slab Items
	NumberOfBlockSlabItems = New(0x300A, 0x0440)

	// BlockSlabSequence (300A,0441) VR=SQ VM=1 Block Slab Sequence
	BlockSlabSequence = New(0x300A, 0x0441)

	// BlockSlabThickness (300A,0442) VR=DS VM=1 Block Slab Thickness
	BlockSlabThickness = New(0x300A, 0x0442)

	// BlockSlabNumber (300A,0443) VR=US VM=1 Block Slab Number
	BlockSlabNumber = New(0x300A, 0x0443)

	// DeviceMotionControlSequence (300A,0450) VR=SQ VM=1 Device Motion Control Sequence
	DeviceMotionControlSequence = New(0x300A, 0x0450)

	// DeviceMotionExecutionMode (300A,0451) VR=CS VM=1 Device Motion Execution Mode
	DeviceMotionExecutionMode = New(0x300A, 0x0451)

	// DeviceMotionObservationMode (300A,0452) VR=CS VM=1 Device Motion Observation Mode
	DeviceMotionObservationMode = New(0x300A, 0x0452)

	// DeviceMotionParameterCodeSequence (300A,0453) VR=SQ VM=1 Device Motion Parameter Code Sequence
	DeviceMotionParameterCodeSequence = New(0x300A, 0x0453)

	// DistalDepthFraction (300A,0501) VR=FL VM=1 Distal Depth Fraction
	DistalDepthFraction = New(0x300A, 0x0501)

	// DistalDepth (300A,0502) VR=FL VM=1 Distal Depth
	DistalDepth = New(0x300A, 0x0502)

	// NominalRangeModulationFractions (300A,0503) VR=FL VM=2 Nominal Range Modulation Fractions
	NominalRangeModulationFractions = New(0x300A, 0x0503)

	// NominalRangeModulatedRegionDepths (300A,0504) VR=FL VM=2 Nominal Range Modulated Region Depths
	NominalRangeModulatedRegionDepths = New(0x300A, 0x0504)

	// DepthDoseParametersSequence (300A,0505) VR=SQ VM=1 Depth Dose Parameters Sequence
	DepthDoseParametersSequence = New(0x300A, 0x0505)

	// DeliveredDepthDoseParametersSequence (300A,0506) VR=SQ VM=1 Delivered Depth Dose Parameters Sequence
	DeliveredDepthDoseParametersSequence = New(0x300A, 0x0506)

	// DeliveredDistalDepthFraction (300A,0507) VR=FL VM=1 Delivered Distal Depth Fraction
	DeliveredDistalDepthFraction = New(0x300A, 0x0507)

	// DeliveredDistalDepth (300A,0508) VR=FL VM=1 Delivered Distal Depth
	DeliveredDistalDepth = New(0x300A, 0x0508)

	// DeliveredNominalRangeModulationFractions (300A,0509) VR=FL VM=2 Delivered Nominal Range Modulation Fractions
	DeliveredNominalRangeModulationFractions = New(0x300A, 0x0509)

	// DeliveredNominalRangeModulatedRegionDepths (300A,0510) VR=FL VM=2 Delivered Nominal Range Modulated Region Depths
	DeliveredNominalRangeModulatedRegionDepths = New(0x300A, 0x0510)

	// DeliveredReferenceDoseDefinition (300A,0511) VR=CS VM=1 Delivered Reference Dose Definition
	DeliveredReferenceDoseDefinition = New(0x300A, 0x0511)

	// ReferenceDoseDefinition (300A,0512) VR=CS VM=1 Reference Dose Definition
	ReferenceDoseDefinition = New(0x300A, 0x0512)

	// RTControlPointIndex (300A,0600) VR=US VM=1 RT Control Point Index
	RTControlPointIndex = New(0x300A, 0x0600)

	// RadiationGenerationModeIndex (300A,0601) VR=US VM=1 Radiation Generation Mode Index
	RadiationGenerationModeIndex = New(0x300A, 0x0601)

	// ReferencedDefinedDeviceIndex (300A,0602) VR=US VM=1 Referenced Defined Device Index
	ReferencedDefinedDeviceIndex = New(0x300A, 0x0602)

	// RadiationDoseIdentificationIndex (300A,0603) VR=US VM=1 Radiation Dose Identification Index
	RadiationDoseIdentificationIndex = New(0x300A, 0x0603)

	// NumberOfRTControlPoints (300A,0604) VR=US VM=1 Number of RT Control Points
	NumberOfRTControlPoints = New(0x300A, 0x0604)

	// ReferencedRadiationGenerationModeIndex (300A,0605) VR=US VM=1 Referenced Radiation Generation Mode Index
	ReferencedRadiationGenerationModeIndex = New(0x300A, 0x0605)

	// TreatmentPositionIndex (300A,0606) VR=US VM=1 Treatment Position Index
	TreatmentPositionIndex = New(0x300A, 0x0606)

	// ReferencedDeviceIndex (300A,0607) VR=US VM=1 Referenced Device Index
	ReferencedDeviceIndex = New(0x300A, 0x0607)

	// TreatmentPositionGroupLabel (300A,0608) VR=LO VM=1 Treatment Position Group Label
	TreatmentPositionGroupLabel = New(0x300A, 0x0608)

	// TreatmentPositionGroupUID (300A,0609) VR=UI VM=1 Treatment Position Group UID
	TreatmentPositionGroupUID = New(0x300A, 0x0609)

	// TreatmentPositionGroupSequence (300A,060A) VR=SQ VM=1 Treatment Position Group Sequence
	TreatmentPositionGroupSequence = New(0x300A, 0x060A)

	// ReferencedTreatmentPositionIndex (300A,060B) VR=US VM=1 Referenced Treatment Position Index
	ReferencedTreatmentPositionIndex = New(0x300A, 0x060B)

	// ReferencedRadiationDoseIdentificationIndex (300A,060C) VR=US VM=1 Referenced Radiation Dose Identification Index
	ReferencedRadiationDoseIdentificationIndex = New(0x300A, 0x060C)

	// RTAccessoryHolderWaterEquivalentThickness (300A,060D) VR=FD VM=1 RT Accessory Holder Water-Equivalent Thickness
	RTAccessoryHolderWaterEquivalentThickness = New(0x300A, 0x060D)

	// ReferencedRTAccessoryHolderDeviceIndex (300A,060E) VR=US VM=1 Referenced RT Accessory Holder Device Index
	ReferencedRTAccessoryHolderDeviceIndex = New(0x300A, 0x060E)

	// RTAccessoryHolderSlotExistenceFlag (300A,060F) VR=CS VM=1 RT Accessory Holder Slot Existence Flag
	RTAccessoryHolderSlotExistenceFlag = New(0x300A, 0x060F)

	// RTAccessoryHolderSlotSequence (300A,0610) VR=SQ VM=1 RT Accessory Holder Slot Sequence
	RTAccessoryHolderSlotSequence = New(0x300A, 0x0610)

	// RTAccessoryHolderSlotID (300A,0611) VR=LO VM=1 RT Accessory Holder Slot ID
	RTAccessoryHolderSlotID = New(0x300A, 0x0611)

	// RTAccessoryHolderSlotDistance (300A,0612) VR=FD VM=1 RT Accessory Holder Slot Distance
	RTAccessoryHolderSlotDistance = New(0x300A, 0x0612)

	// RTAccessorySlotDistance (300A,0613) VR=FD VM=1 RT Accessory Slot Distance
	RTAccessorySlotDistance = New(0x300A, 0x0613)

	// RTAccessoryHolderDefinitionSequence (300A,0614) VR=SQ VM=1 RT Accessory Holder Definition Sequence
	RTAccessoryHolderDefinitionSequence = New(0x300A, 0x0614)

	// RTAccessoryDeviceSlotID (300A,0615) VR=LO VM=1 RT Accessory Device Slot ID
	RTAccessoryDeviceSlotID = New(0x300A, 0x0615)

	// RTRadiationSequence (300A,0616) VR=SQ VM=1 RT Radiation Sequence
	RTRadiationSequence = New(0x300A, 0x0616)

	// RadiationDoseSequence (300A,0617) VR=SQ VM=1 Radiation Dose Sequence
	RadiationDoseSequence = New(0x300A, 0x0617)

	// RadiationDoseIdentificationSequence (300A,0618) VR=SQ VM=1 Radiation Dose Identification Sequence
	RadiationDoseIdentificationSequence = New(0x300A, 0x0618)

	// RadiationDoseIdentificationLabel (300A,0619) VR=LO VM=1 Radiation Dose Identification Label
	RadiationDoseIdentificationLabel = New(0x300A, 0x0619)

	// ReferenceDoseType (300A,061A) VR=CS VM=1 Reference Dose Type
	ReferenceDoseType = New(0x300A, 0x061A)

	// PrimaryDoseValueIndicator (300A,061B) VR=CS VM=1 Primary Dose Value Indicator
	PrimaryDoseValueIndicator = New(0x300A, 0x061B)

	// DoseValuesSequence (300A,061C) VR=SQ VM=1 Dose Values Sequence
	DoseValuesSequence = New(0x300A, 0x061C)

	// DoseValuePurpose (300A,061D) VR=CS VM=1-n Dose Value Purpose
	DoseValuePurpose = New(0x300A, 0x061D)

	// ReferenceDosePointCoordinates (300A,061E) VR=FD VM=3 Reference Dose Point Coordinates
	ReferenceDosePointCoordinates = New(0x300A, 0x061E)

	// RadiationDoseValuesParametersSequence (300A,061F) VR=SQ VM=1 Radiation Dose Values Parameters Sequence
	RadiationDoseValuesParametersSequence = New(0x300A, 0x061F)

	// MetersetToDoseMappingSequence (300A,0620) VR=SQ VM=1 Meterset to Dose Mapping Sequence
	MetersetToDoseMappingSequence = New(0x300A, 0x0620)

	// ExpectedInVivoMeasurementValuesSequence (300A,0621) VR=SQ VM=1 Expected In-Vivo Measurement Values Sequence
	ExpectedInVivoMeasurementValuesSequence = New(0x300A, 0x0621)

	// ExpectedInVivoMeasurementValueIndex (300A,0622) VR=US VM=1 Expected In-Vivo Measurement Value Index
	ExpectedInVivoMeasurementValueIndex = New(0x300A, 0x0622)

	// RadiationDoseInVivoMeasurementLabel (300A,0623) VR=LO VM=1 Radiation Dose In-Vivo Measurement Label
	RadiationDoseInVivoMeasurementLabel = New(0x300A, 0x0623)

	// RadiationDoseCentralAxisDisplacement (300A,0624) VR=FD VM=2 Radiation Dose Central Axis Displacement
	RadiationDoseCentralAxisDisplacement = New(0x300A, 0x0624)

	// RadiationDoseValue (300A,0625) VR=FD VM=1 Radiation Dose Value
	RadiationDoseValue = New(0x300A, 0x0625)

	// RadiationDoseSourceToSkinDistance (300A,0626) VR=FD VM=1 Radiation Dose Source to Skin Distance
	RadiationDoseSourceToSkinDistance = New(0x300A, 0x0626)

	// RadiationDoseMeasurementPointCoordinates (300A,0627) VR=FD VM=3 Radiation Dose Measurement Point Coordinates
	RadiationDoseMeasurementPointCoordinates = New(0x300A, 0x0627)

	// RadiationDoseSourceToExternalContourDistance (300A,0628) VR=FD VM=1 Radiation Dose Source to External Contour Distance
	RadiationDoseSourceToExternalContourDistance = New(0x300A, 0x0628)

	// RTToleranceSetSequence (300A,0629) VR=SQ VM=1 RT Tolerance Set Sequence
	RTToleranceSetSequence = New(0x300A, 0x0629)

	// RTToleranceSetLabel (300A,062A) VR=LO VM=1 RT Tolerance Set Label
	RTToleranceSetLabel = New(0x300A, 0x062A)

	// AttributeToleranceValuesSequence (300A,062B) VR=SQ VM=1 Attribute Tolerance Values Sequence
	AttributeToleranceValuesSequence = New(0x300A, 0x062B)

	// ToleranceValue (300A,062C) VR=FD VM=1 Tolerance Value
	ToleranceValue = New(0x300A, 0x062C)

	// PatientSupportPositionToleranceSequence (300A,062D) VR=SQ VM=1 Patient Support Position Tolerance Sequence
	PatientSupportPositionToleranceSequence = New(0x300A, 0x062D)

	// TreatmentTimeLimit (300A,062E) VR=FD VM=1 Treatment Time Limit
	TreatmentTimeLimit = New(0x300A, 0x062E)

	// CArmPhotonElectronControlPointSequence (300A,062F) VR=SQ VM=1 C-Arm Photon-Electron Control Point Sequence
	CArmPhotonElectronControlPointSequence = New(0x300A, 0x062F)

	// ReferencedRTRadiationSequence (300A,0630) VR=SQ VM=1 Referenced RT Radiation Sequence
	ReferencedRTRadiationSequence = New(0x300A, 0x0630)

	// ReferencedRTInstanceSequence (300A,0631) VR=SQ VM=1 Referenced RT Instance Sequence
	ReferencedRTInstanceSequence = New(0x300A, 0x0631)

	// ReferencedRTPatientSetupSequenceRETIRED (300A,0632) VR=SQ VM=1 Referenced RT Patient Setup Sequence (RETIRED)
	ReferencedRTPatientSetupSequenceRETIRED = New(0x300A, 0x0632)

	// SourceToPatientSurfaceDistance (300A,0634) VR=FD VM=1 Source to Patient Surface Distance
	SourceToPatientSurfaceDistance = New(0x300A, 0x0634)

	// TreatmentMachineSpecialModeCodeSequence (300A,0635) VR=SQ VM=1 Treatment Machine Special Mode Code Sequence
	TreatmentMachineSpecialModeCodeSequence = New(0x300A, 0x0635)

	// IntendedNumberOfFractions (300A,0636) VR=US VM=1 Intended Number of Fractions
	IntendedNumberOfFractions = New(0x300A, 0x0636)

	// RTRadiationSetIntent (300A,0637) VR=CS VM=1 RT Radiation Set Intent
	RTRadiationSetIntent = New(0x300A, 0x0637)

	// RTRadiationPhysicalAndGeometricContentDetailFlag (300A,0638) VR=CS VM=1 RT Radiation Physical and Geometric Content Detail Flag
	RTRadiationPhysicalAndGeometricContentDetailFlag = New(0x300A, 0x0638)

	// RTRecordFlag (300A,0639) VR=CS VM=1 RT Record Flag
	RTRecordFlag = New(0x300A, 0x0639)

	// TreatmentDeviceIdentificationSequence (300A,063A) VR=SQ VM=1 Treatment Device Identification Sequence
	TreatmentDeviceIdentificationSequence = New(0x300A, 0x063A)

	// ReferencedRTPhysicianIntentSequence (300A,063B) VR=SQ VM=1 Referenced RT Physician Intent Sequence
	ReferencedRTPhysicianIntentSequence = New(0x300A, 0x063B)

	// CumulativeMeterset (300A,063C) VR=FD VM=1 Cumulative Meterset
	CumulativeMeterset = New(0x300A, 0x063C)

	// DeliveryRate (300A,063D) VR=FD VM=1 Delivery Rate
	DeliveryRate = New(0x300A, 0x063D)

	// DeliveryRateUnitSequence (300A,063E) VR=SQ VM=1 Delivery Rate Unit Sequence
	DeliveryRateUnitSequence = New(0x300A, 0x063E)

	// TreatmentPositionSequence (300A,063F) VR=SQ VM=1 Treatment Position Sequence
	TreatmentPositionSequence = New(0x300A, 0x063F)

	// RadiationSourceAxisDistance (300A,0640) VR=FD VM=1 Radiation Source-Axis Distance
	RadiationSourceAxisDistance = New(0x300A, 0x0640)

	// NumberOfRTBeamLimitingDevices (300A,0641) VR=US VM=1 Number of RT Beam Limiting Devices
	NumberOfRTBeamLimitingDevices = New(0x300A, 0x0641)

	// RTBeamLimitingDeviceProximalDistance (300A,0642) VR=FD VM=1 RT Beam Limiting Device Proximal Distance
	RTBeamLimitingDeviceProximalDistance = New(0x300A, 0x0642)

	// RTBeamLimitingDeviceDistalDistance (300A,0643) VR=FD VM=1 RT Beam Limiting Device Distal Distance
	RTBeamLimitingDeviceDistalDistance = New(0x300A, 0x0643)

	// ParallelRTBeamDelimiterDeviceOrientationLabelCodeSequence (300A,0644) VR=SQ VM=1 Parallel RT Beam Delimiter Device Orientation Label Code Sequence
	ParallelRTBeamDelimiterDeviceOrientationLabelCodeSequence = New(0x300A, 0x0644)

	// BeamModifierOrientationAngle (300A,0645) VR=FD VM=1 Beam Modifier Orientation Angle
	BeamModifierOrientationAngle = New(0x300A, 0x0645)

	// FixedRTBeamDelimiterDeviceSequence (300A,0646) VR=SQ VM=1 Fixed RT Beam Delimiter Device Sequence
	FixedRTBeamDelimiterDeviceSequence = New(0x300A, 0x0646)

	// ParallelRTBeamDelimiterDeviceSequence (300A,0647) VR=SQ VM=1 Parallel RT Beam Delimiter Device Sequence
	ParallelRTBeamDelimiterDeviceSequence = New(0x300A, 0x0647)

	// NumberOfParallelRTBeamDelimiters (300A,0648) VR=US VM=1 Number of Parallel RT Beam Delimiters
	NumberOfParallelRTBeamDelimiters = New(0x300A, 0x0648)

	// ParallelRTBeamDelimiterBoundaries (300A,0649) VR=FD VM=2-n Parallel RT Beam Delimiter Boundaries
	ParallelRTBeamDelimiterBoundaries = New(0x300A, 0x0649)

	// ParallelRTBeamDelimiterPositions (300A,064A) VR=FD VM=2-n Parallel RT Beam Delimiter Positions
	ParallelRTBeamDelimiterPositions = New(0x300A, 0x064A)

	// RTBeamLimitingDeviceOffset (300A,064B) VR=FD VM=2 RT Beam Limiting Device Offset
	RTBeamLimitingDeviceOffset = New(0x300A, 0x064B)

	// RTBeamDelimiterGeometrySequence (300A,064C) VR=SQ VM=1 RT Beam Delimiter Geometry Sequence
	RTBeamDelimiterGeometrySequence = New(0x300A, 0x064C)

	// RTBeamLimitingDeviceDefinitionSequence (300A,064D) VR=SQ VM=1 RT Beam Limiting Device Definition Sequence
	RTBeamLimitingDeviceDefinitionSequence = New(0x300A, 0x064D)

	// ParallelRTBeamDelimiterOpeningMode (300A,064E) VR=CS VM=1 Parallel RT Beam Delimiter Opening Mode
	ParallelRTBeamDelimiterOpeningMode = New(0x300A, 0x064E)

	// ParallelRTBeamDelimiterLeafMountingSide (300A,064F) VR=CS VM=1-n Parallel RT Beam Delimiter Leaf Mounting Side
	ParallelRTBeamDelimiterLeafMountingSide = New(0x300A, 0x064F)

	// PatientSetupUIDRETIRED (300A,0650) VR=UI VM=1 Patient Setup UID (RETIRED)
	PatientSetupUIDRETIRED = New(0x300A, 0x0650)

	// WedgeDefinitionSequence (300A,0651) VR=SQ VM=1 Wedge Definition Sequence
	WedgeDefinitionSequence = New(0x300A, 0x0651)

	// RadiationBeamWedgeAngle (300A,0652) VR=FD VM=1 Radiation Beam Wedge Angle
	RadiationBeamWedgeAngle = New(0x300A, 0x0652)

	// RadiationBeamWedgeThinEdgeDistance (300A,0653) VR=FD VM=1 Radiation Beam Wedge Thin Edge Distance
	RadiationBeamWedgeThinEdgeDistance = New(0x300A, 0x0653)

	// RadiationBeamEffectiveWedgeAngle (300A,0654) VR=FD VM=1 Radiation Beam Effective Wedge Angle
	RadiationBeamEffectiveWedgeAngle = New(0x300A, 0x0654)

	// NumberOfWedgePositions (300A,0655) VR=US VM=1 Number of Wedge Positions
	NumberOfWedgePositions = New(0x300A, 0x0655)

	// RTBeamLimitingDeviceOpeningSequence (300A,0656) VR=SQ VM=1 RT Beam Limiting Device Opening Sequence
	RTBeamLimitingDeviceOpeningSequence = New(0x300A, 0x0656)

	// NumberOfRTBeamLimitingDeviceOpenings (300A,0657) VR=US VM=1 Number of RT Beam Limiting Device Openings
	NumberOfRTBeamLimitingDeviceOpenings = New(0x300A, 0x0657)

	// RadiationDosimeterUnitSequence (300A,0658) VR=SQ VM=1 Radiation Dosimeter Unit Sequence
	RadiationDosimeterUnitSequence = New(0x300A, 0x0658)

	// RTDeviceDistanceReferenceLocationCodeSequence (300A,0659) VR=SQ VM=1 RT Device Distance Reference Location Code Sequence
	RTDeviceDistanceReferenceLocationCodeSequence = New(0x300A, 0x0659)

	// RadiationDeviceConfigurationAndCommissioningKeySequence (300A,065A) VR=SQ VM=1 Radiation Device Configuration and Commissioning Key Sequence
	RadiationDeviceConfigurationAndCommissioningKeySequence = New(0x300A, 0x065A)

	// PatientSupportPositionParameterSequence (300A,065B) VR=SQ VM=1 Patient Support Position Parameter Sequence
	PatientSupportPositionParameterSequence = New(0x300A, 0x065B)

	// PatientSupportPositionSpecificationMethod (300A,065C) VR=CS VM=1 Patient Support Position Specification Method
	PatientSupportPositionSpecificationMethod = New(0x300A, 0x065C)

	// PatientSupportPositionDeviceParameterSequence (300A,065D) VR=SQ VM=1 Patient Support Position Device Parameter Sequence
	PatientSupportPositionDeviceParameterSequence = New(0x300A, 0x065D)

	// DeviceOrderIndex (300A,065E) VR=US VM=1 Device Order Index
	DeviceOrderIndex = New(0x300A, 0x065E)

	// PatientSupportPositionParameterOrderIndex (300A,065F) VR=US VM=1 Patient Support Position Parameter Order Index
	PatientSupportPositionParameterOrderIndex = New(0x300A, 0x065F)

	// PatientSupportPositionDeviceToleranceSequence (300A,0660) VR=SQ VM=1 Patient Support Position Device Tolerance Sequence
	PatientSupportPositionDeviceToleranceSequence = New(0x300A, 0x0660)

	// PatientSupportPositionToleranceOrderIndex (300A,0661) VR=US VM=1 Patient Support Position Tolerance Order Index
	PatientSupportPositionToleranceOrderIndex = New(0x300A, 0x0661)

	// CompensatorDefinitionSequence (300A,0662) VR=SQ VM=1 Compensator Definition Sequence
	CompensatorDefinitionSequence = New(0x300A, 0x0662)

	// CompensatorMapOrientation (300A,0663) VR=CS VM=1 Compensator Map Orientation
	CompensatorMapOrientation = New(0x300A, 0x0663)

	// CompensatorProximalThicknessMap (300A,0664) VR=OF VM=1 Compensator Proximal Thickness Map
	CompensatorProximalThicknessMap = New(0x300A, 0x0664)

	// CompensatorDistalThicknessMap (300A,0665) VR=OF VM=1 Compensator Distal Thickness Map
	CompensatorDistalThicknessMap = New(0x300A, 0x0665)

	// CompensatorBasePlaneOffset (300A,0666) VR=FD VM=1 Compensator Base Plane Offset
	CompensatorBasePlaneOffset = New(0x300A, 0x0666)

	// CompensatorShapeFabricationCodeSequence (300A,0667) VR=SQ VM=1 Compensator Shape Fabrication Code Sequence
	CompensatorShapeFabricationCodeSequence = New(0x300A, 0x0667)

	// CompensatorShapeSequence (300A,0668) VR=SQ VM=1 Compensator Shape Sequence
	CompensatorShapeSequence = New(0x300A, 0x0668)

	// RadiationBeamCompensatorMillingToolDiameter (300A,0669) VR=FD VM=1 Radiation Beam Compensator Milling Tool Diameter
	RadiationBeamCompensatorMillingToolDiameter = New(0x300A, 0x0669)

	// BlockDefinitionSequence (300A,066A) VR=SQ VM=1 Block Definition Sequence
	BlockDefinitionSequence = New(0x300A, 0x066A)

	// BlockEdgeData (300A,066B) VR=OF VM=1 Block Edge Data
	BlockEdgeData = New(0x300A, 0x066B)

	// BlockOrientation (300A,066C) VR=CS VM=1 Block Orientation
	BlockOrientation = New(0x300A, 0x066C)

	// RadiationBeamBlockThickness (300A,066D) VR=FD VM=1 Radiation Beam Block Thickness
	RadiationBeamBlockThickness = New(0x300A, 0x066D)

	// RadiationBeamBlockSlabThickness (300A,066E) VR=FD VM=1 Radiation Beam Block Slab Thickness
	RadiationBeamBlockSlabThickness = New(0x300A, 0x066E)

	// BlockEdgeDataSequence (300A,066F) VR=SQ VM=1 Block Edge Data Sequence
	BlockEdgeDataSequence = New(0x300A, 0x066F)

	// NumberOfRTAccessoryHolders (300A,0670) VR=US VM=1 Number of RT Accessory Holders
	NumberOfRTAccessoryHolders = New(0x300A, 0x0670)

	// GeneralAccessoryDefinitionSequence (300A,0671) VR=SQ VM=1 General Accessory Definition Sequence
	GeneralAccessoryDefinitionSequence = New(0x300A, 0x0671)

	// NumberOfGeneralAccessories (300A,0672) VR=US VM=1 Number of General Accessories
	NumberOfGeneralAccessories = New(0x300A, 0x0672)

	// BolusDefinitionSequence (300A,0673) VR=SQ VM=1 Bolus Definition Sequence
	BolusDefinitionSequence = New(0x300A, 0x0673)

	// NumberOfBoluses (300A,0674) VR=US VM=1 Number of Boluses
	NumberOfBoluses = New(0x300A, 0x0674)

	// EquipmentFrameOfReferenceUID (300A,0675) VR=UI VM=1 Equipment Frame of Reference UID
	EquipmentFrameOfReferenceUID = New(0x300A, 0x0675)

	// EquipmentFrameOfReferenceDescription (300A,0676) VR=ST VM=1 Equipment Frame of Reference Description
	EquipmentFrameOfReferenceDescription = New(0x300A, 0x0676)

	// EquipmentReferencePointCoordinatesSequence (300A,0677) VR=SQ VM=1 Equipment Reference Point Coordinates Sequence
	EquipmentReferencePointCoordinatesSequence = New(0x300A, 0x0677)

	// EquipmentReferencePointCodeSequence (300A,0678) VR=SQ VM=1 Equipment Reference Point Code Sequence
	EquipmentReferencePointCodeSequence = New(0x300A, 0x0678)

	// RTBeamLimitingDeviceAngle (300A,0679) VR=FD VM=1 RT Beam Limiting Device Angle
	RTBeamLimitingDeviceAngle = New(0x300A, 0x0679)

	// SourceRollAngle (300A,067A) VR=FD VM=1 Source Roll Angle
	SourceRollAngle = New(0x300A, 0x067A)

	// RadiationGenerationModeSequence (300A,067B) VR=SQ VM=1 Radiation GenerationMode Sequence
	RadiationGenerationModeSequence = New(0x300A, 0x067B)

	// RadiationGenerationModeLabel (300A,067C) VR=SH VM=1 Radiation GenerationMode Label
	RadiationGenerationModeLabel = New(0x300A, 0x067C)

	// RadiationGenerationModeDescription (300A,067D) VR=ST VM=1 Radiation GenerationMode Description
	RadiationGenerationModeDescription = New(0x300A, 0x067D)

	// RadiationGenerationModeMachineCodeSequence (300A,067E) VR=SQ VM=1 Radiation GenerationMode Machine Code Sequence
	RadiationGenerationModeMachineCodeSequence = New(0x300A, 0x067E)

	// RadiationTypeCodeSequence (300A,067F) VR=SQ VM=1 Radiation Type Code Sequence
	RadiationTypeCodeSequence = New(0x300A, 0x067F)

	// NominalEnergy (300A,0680) VR=DS VM=1 Nominal Energy
	NominalEnergy = New(0x300A, 0x0680)

	// MinimumNominalEnergy (300A,0681) VR=DS VM=1 Minimum Nominal Energy
	MinimumNominalEnergy = New(0x300A, 0x0681)

	// MaximumNominalEnergy (300A,0682) VR=DS VM=1 Maximum Nominal Energy
	MaximumNominalEnergy = New(0x300A, 0x0682)

	// RadiationFluenceModifierCodeSequence (300A,0683) VR=SQ VM=1 Radiation Fluence Modifier Code Sequence
	RadiationFluenceModifierCodeSequence = New(0x300A, 0x0683)

	// EnergyUnitCodeSequence (300A,0684) VR=SQ VM=1 Energy Unit Code Sequence
	EnergyUnitCodeSequence = New(0x300A, 0x0684)

	// NumberOfRadiationGenerationModes (300A,0685) VR=US VM=1 Number of Radiation GenerationModes
	NumberOfRadiationGenerationModes = New(0x300A, 0x0685)

	// PatientSupportDevicesSequence (300A,0686) VR=SQ VM=1 Patient Support Devices Sequence
	PatientSupportDevicesSequence = New(0x300A, 0x0686)

	// NumberOfPatientSupportDevices (300A,0687) VR=US VM=1 Number of Patient Support Devices
	NumberOfPatientSupportDevices = New(0x300A, 0x0687)

	// RTBeamModifierDefinitionDistance (300A,0688) VR=FD VM=1 RT Beam Modifier Definition Distance
	RTBeamModifierDefinitionDistance = New(0x300A, 0x0688)

	// BeamAreaLimitSequence (300A,0689) VR=SQ VM=1 Beam Area Limit Sequence
	BeamAreaLimitSequence = New(0x300A, 0x0689)

	// ReferencedRTPrescriptionSequence (300A,068A) VR=SQ VM=1 Referenced RT Prescription Sequence
	ReferencedRTPrescriptionSequence = New(0x300A, 0x068A)

	// DoseValueInterpretation (300A,068B) VR=CS VM=1 Dose Value Interpretation
	DoseValueInterpretation = New(0x300A, 0x068B)

	// TreatmentSessionUID (300A,0700) VR=UI VM=1 Treatment Session UID
	TreatmentSessionUID = New(0x300A, 0x0700)

	// RTRadiationUsage (300A,0701) VR=CS VM=1 RT Radiation Usage
	RTRadiationUsage = New(0x300A, 0x0701)

	// ReferencedRTRadiationSetSequence (300A,0702) VR=SQ VM=1 Referenced RT Radiation Set Sequence
	ReferencedRTRadiationSetSequence = New(0x300A, 0x0702)

	// ReferencedRTRadiationRecordSequence (300A,0703) VR=SQ VM=1 Referenced RT Radiation Record Sequence
	ReferencedRTRadiationRecordSequence = New(0x300A, 0x0703)

	// RTRadiationSetDeliveryNumber (300A,0704) VR=US VM=1 RT Radiation Set Delivery Number
	RTRadiationSetDeliveryNumber = New(0x300A, 0x0704)

	// ClinicalFractionNumber (300A,0705) VR=US VM=1 Clinical Fraction Number
	ClinicalFractionNumber = New(0x300A, 0x0705)

	// RTTreatmentFractionCompletionStatus (300A,0706) VR=CS VM=1 RT Treatment Fraction Completion Status
	RTTreatmentFractionCompletionStatus = New(0x300A, 0x0706)

	// RTRadiationSetUsage (300A,0707) VR=CS VM=1 RT Radiation Set Usage
	RTRadiationSetUsage = New(0x300A, 0x0707)

	// TreatmentDeliveryContinuationFlag (300A,0708) VR=CS VM=1 Treatment Delivery Continuation Flag
	TreatmentDeliveryContinuationFlag = New(0x300A, 0x0708)

	// TreatmentRecordContentOrigin (300A,0709) VR=CS VM=1 Treatment Record Content Origin
	TreatmentRecordContentOrigin = New(0x300A, 0x0709)

	// RTTreatmentTerminationStatus (300A,0714) VR=CS VM=1 RT Treatment Termination Status
	RTTreatmentTerminationStatus = New(0x300A, 0x0714)

	// RTTreatmentTerminationReasonCodeSequence (300A,0715) VR=SQ VM=1 RT Treatment Termination Reason Code Sequence
	RTTreatmentTerminationReasonCodeSequence = New(0x300A, 0x0715)

	// MachineSpecificTreatmentTerminationCodeSequence (300A,0716) VR=SQ VM=1 Machine-Specific Treatment Termination Code Sequence
	MachineSpecificTreatmentTerminationCodeSequence = New(0x300A, 0x0716)

	// RTRadiationSalvageRecordControlPointSequence (300A,0722) VR=SQ VM=1 RT Radiation Salvage Record Control Point Sequence
	RTRadiationSalvageRecordControlPointSequence = New(0x300A, 0x0722)

	// StartingMetersetValueKnownFlag (300A,0723) VR=CS VM=1 Starting Meterset Value Known Flag
	StartingMetersetValueKnownFlag = New(0x300A, 0x0723)

	// TreatmentTerminationDescription (300A,0730) VR=ST VM=1 Treatment Termination Description
	TreatmentTerminationDescription = New(0x300A, 0x0730)

	// TreatmentToleranceViolationSequence (300A,0731) VR=SQ VM=1 Treatment Tolerance Violation Sequence
	TreatmentToleranceViolationSequence = New(0x300A, 0x0731)

	// TreatmentToleranceViolationCategory (300A,0732) VR=CS VM=1 Treatment Tolerance Violation Category
	TreatmentToleranceViolationCategory = New(0x300A, 0x0732)

	// TreatmentToleranceViolationAttributeSequence (300A,0733) VR=SQ VM=1 Treatment Tolerance Violation Attribute Sequence
	TreatmentToleranceViolationAttributeSequence = New(0x300A, 0x0733)

	// TreatmentToleranceViolationDescription (300A,0734) VR=ST VM=1 Treatment Tolerance Violation Description
	TreatmentToleranceViolationDescription = New(0x300A, 0x0734)

	// TreatmentToleranceViolationIdentification (300A,0735) VR=ST VM=1 Treatment Tolerance Violation Identification
	TreatmentToleranceViolationIdentification = New(0x300A, 0x0735)

	// TreatmentToleranceViolationDateTime (300A,0736) VR=DT VM=1 Treatment Tolerance Violation DateTime
	TreatmentToleranceViolationDateTime = New(0x300A, 0x0736)

	// RecordedRTControlPointDateTime (300A,073A) VR=DT VM=1 Recorded RT Control Point DateTime
	RecordedRTControlPointDateTime = New(0x300A, 0x073A)

	// ReferencedRadiationRTControlPointIndex (300A,073B) VR=US VM=1 Referenced Radiation RT Control Point Index
	ReferencedRadiationRTControlPointIndex = New(0x300A, 0x073B)

	// AlternateValueSequence (300A,073E) VR=SQ VM=1 Alternate Value Sequence
	AlternateValueSequence = New(0x300A, 0x073E)

	// ConfirmationSequence (300A,073F) VR=SQ VM=1 Confirmation Sequence
	ConfirmationSequence = New(0x300A, 0x073F)

	// InterlockSequence (300A,0740) VR=SQ VM=1 Interlock Sequence
	InterlockSequence = New(0x300A, 0x0740)

	// InterlockDateTime (300A,0741) VR=DT VM=1 Interlock DateTime
	InterlockDateTime = New(0x300A, 0x0741)

	// InterlockDescription (300A,0742) VR=ST VM=1 Interlock Description
	InterlockDescription = New(0x300A, 0x0742)

	// InterlockOriginatingDeviceSequence (300A,0743) VR=SQ VM=1 Interlock Originating Device Sequence
	InterlockOriginatingDeviceSequence = New(0x300A, 0x0743)

	// InterlockCodeSequence (300A,0744) VR=SQ VM=1 Interlock Code Sequence
	InterlockCodeSequence = New(0x300A, 0x0744)

	// InterlockResolutionCodeSequence (300A,0745) VR=SQ VM=1 Interlock Resolution Code Sequence
	InterlockResolutionCodeSequence = New(0x300A, 0x0745)

	// InterlockResolutionUserSequence (300A,0746) VR=SQ VM=1 Interlock Resolution User Sequence
	InterlockResolutionUserSequence = New(0x300A, 0x0746)

	// OverrideDateTime (300A,0760) VR=DT VM=1 Override DateTime
	OverrideDateTime = New(0x300A, 0x0760)

	// TreatmentToleranceViolationTypeCodeSequence (300A,0761) VR=SQ VM=1 Treatment Tolerance Violation Type Code Sequence
	TreatmentToleranceViolationTypeCodeSequence = New(0x300A, 0x0761)

	// TreatmentToleranceViolationCauseCodeSequence (300A,0762) VR=SQ VM=1 Treatment Tolerance Violation Cause Code Sequence
	TreatmentToleranceViolationCauseCodeSequence = New(0x300A, 0x0762)

	// MeasuredMetersetToDoseMappingSequence (300A,0772) VR=SQ VM=1 Measured Meterset to Dose Mapping Sequence
	MeasuredMetersetToDoseMappingSequence = New(0x300A, 0x0772)

	// ReferencedExpectedInVivoMeasurementValueIndex (300A,0773) VR=US VM=1 Referenced Expected In-Vivo Measurement Value Index
	ReferencedExpectedInVivoMeasurementValueIndex = New(0x300A, 0x0773)

	// DoseMeasurementDeviceCodeSequence (300A,0774) VR=SQ VM=1 Dose Measurement Device Code Sequence
	DoseMeasurementDeviceCodeSequence = New(0x300A, 0x0774)

	// AdditionalParameterRecordingInstanceSequence (300A,0780) VR=SQ VM=1 Additional Parameter Recording Instance Sequence
	AdditionalParameterRecordingInstanceSequence = New(0x300A, 0x0780)

	// InterlockOriginDescription (300A,0783) VR=ST VM=1 Interlock Origin Description
	InterlockOriginDescription = New(0x300A, 0x0783)

	// RTPatientPositionScopeSequence (300A,0784) VR=SQ VM=1 RT Patient Position Scope Sequence
	RTPatientPositionScopeSequence = New(0x300A, 0x0784)

	// ReferencedTreatmentPositionGroupUID (300A,0785) VR=UI VM=1 Referenced Treatment Position Group UID
	ReferencedTreatmentPositionGroupUID = New(0x300A, 0x0785)

	// RadiationOrderIndex (300A,0786) VR=US VM=1 Radiation Order Index
	RadiationOrderIndex = New(0x300A, 0x0786)

	// OmittedRadiationSequence (300A,0787) VR=SQ VM=1 Omitted Radiation Sequence
	OmittedRadiationSequence = New(0x300A, 0x0787)

	// ReasonForOmissionCodeSequence (300A,0788) VR=SQ VM=1 Reason for Omission Code Sequence
	ReasonForOmissionCodeSequence = New(0x300A, 0x0788)

	// RTDeliveryStartPatientPositionSequence (300A,0789) VR=SQ VM=1 RT Delivery Start Patient Position Sequence
	RTDeliveryStartPatientPositionSequence = New(0x300A, 0x0789)

	// RTTreatmentPreparationPatientPositionSequence (300A,078A) VR=SQ VM=1 RT Treatment Preparation Patient Position Sequence
	RTTreatmentPreparationPatientPositionSequence = New(0x300A, 0x078A)

	// ReferencedRTTreatmentPreparationSequence (300A,078B) VR=SQ VM=1 Referenced RT Treatment Preparation Sequence
	ReferencedRTTreatmentPreparationSequence = New(0x300A, 0x078B)

	// ReferencedPatientSetupPhotoSequence (300A,078C) VR=SQ VM=1 Referenced Patient Setup Photo Sequence
	ReferencedPatientSetupPhotoSequence = New(0x300A, 0x078C)

	// PatientTreatmentPreparationMethodCodeSequence (300A,078D) VR=SQ VM=1 Patient Treatment Preparation Method Code Sequence
	PatientTreatmentPreparationMethodCodeSequence = New(0x300A, 0x078D)

	// PatientTreatmentPreparationProcedureParameterDescription (300A,078E) VR=LT VM=1 Patient Treatment Preparation Procedure Parameter Description
	PatientTreatmentPreparationProcedureParameterDescription = New(0x300A, 0x078E)

	// PatientTreatmentPreparationDeviceSequence (300A,078F) VR=SQ VM=1 Patient Treatment Preparation Device Sequence
	PatientTreatmentPreparationDeviceSequence = New(0x300A, 0x078F)

	// PatientTreatmentPreparationProcedureSequence (300A,0790) VR=SQ VM=1 Patient Treatment Preparation Procedure Sequence
	PatientTreatmentPreparationProcedureSequence = New(0x300A, 0x0790)

	// PatientTreatmentPreparationProcedureCodeSequence (300A,0791) VR=SQ VM=1 Patient Treatment Preparation Procedure Code Sequence
	PatientTreatmentPreparationProcedureCodeSequence = New(0x300A, 0x0791)

	// PatientTreatmentPreparationMethodDescription (300A,0792) VR=LT VM=1 Patient Treatment Preparation Method Description
	PatientTreatmentPreparationMethodDescription = New(0x300A, 0x0792)

	// PatientTreatmentPreparationProcedureParameterSequence (300A,0793) VR=SQ VM=1 Patient Treatment Preparation Procedure Parameter Sequence
	PatientTreatmentPreparationProcedureParameterSequence = New(0x300A, 0x0793)

	// PatientSetupPhotoDescription (300A,0794) VR=LT VM=1 Patient Setup Photo Description
	PatientSetupPhotoDescription = New(0x300A, 0x0794)

	// PatientTreatmentPreparationProcedureIndex (300A,0795) VR=US VM=1 Patient Treatment Preparation Procedure Index
	PatientTreatmentPreparationProcedureIndex = New(0x300A, 0x0795)

	// ReferencedPatientSetupProcedureIndex (300A,0796) VR=US VM=1 Referenced Patient Setup Procedure Index
	ReferencedPatientSetupProcedureIndex = New(0x300A, 0x0796)

	// RTRadiationTaskSequence (300A,0797) VR=SQ VM=1 RT Radiation Task Sequence
	RTRadiationTaskSequence = New(0x300A, 0x0797)

	// RTPatientPositionDisplacementSequence (300A,0798) VR=SQ VM=1 RT Patient Position Displacement Sequence
	RTPatientPositionDisplacementSequence = New(0x300A, 0x0798)

	// RTPatientPositionSequence (300A,0799) VR=SQ VM=1 RT Patient Position Sequence
	RTPatientPositionSequence = New(0x300A, 0x0799)

	// DisplacementReferenceLabel (300A,079A) VR=LO VM=1 Displacement Reference Label
	DisplacementReferenceLabel = New(0x300A, 0x079A)

	// DisplacementMatrix (300A,079B) VR=FD VM=16 Displacement Matrix
	DisplacementMatrix = New(0x300A, 0x079B)

	// PatientSupportDisplacementSequence (300A,079C) VR=SQ VM=1 Patient Support Displacement Sequence
	PatientSupportDisplacementSequence = New(0x300A, 0x079C)

	// DisplacementReferenceLocationCodeSequence (300A,079D) VR=SQ VM=1 Displacement Reference Location Code Sequence
	DisplacementReferenceLocationCodeSequence = New(0x300A, 0x079D)

	// RTRadiationSetDeliveryUsage (300A,079E) VR=CS VM=1 RT Radiation Set Delivery Usage
	RTRadiationSetDeliveryUsage = New(0x300A, 0x079E)

	// PatientTreatmentPreparationSequence (300A,079F) VR=SQ VM=1 Patient Treatment Preparation Sequence
	PatientTreatmentPreparationSequence = New(0x300A, 0x079F)

	// PatientToEquipmentRelationshipSequence (300A,07A0) VR=SQ VM=1 Patient to Equipment Relationship Sequence
	PatientToEquipmentRelationshipSequence = New(0x300A, 0x07A0)

	// ImagingEquipmentToTreatmentDeliveryDeviceRelationshipSequence (300A,07A1) VR=SQ VM=1 Imaging Equipment to Treatment Delivery Device Relationship Sequence
	ImagingEquipmentToTreatmentDeliveryDeviceRelationshipSequence = New(0x300A, 0x07A1)

	// ReferencedRTPlanSequence (300C,0002) VR=SQ VM=1 Referenced RT Plan Sequence
	ReferencedRTPlanSequence = New(0x300C, 0x0002)

	// ReferencedBeamSequence (300C,0004) VR=SQ VM=1 Referenced Beam Sequence
	ReferencedBeamSequence = New(0x300C, 0x0004)

	// ReferencedBeamNumber (300C,0006) VR=IS VM=1 Referenced Beam Number
	ReferencedBeamNumber = New(0x300C, 0x0006)

	// ReferencedReferenceImageNumber (300C,0007) VR=IS VM=1 Referenced Reference Image Number
	ReferencedReferenceImageNumber = New(0x300C, 0x0007)

	// StartCumulativeMetersetWeight (300C,0008) VR=DS VM=1 Start Cumulative Meterset Weight
	StartCumulativeMetersetWeight = New(0x300C, 0x0008)

	// EndCumulativeMetersetWeight (300C,0009) VR=DS VM=1 End Cumulative Meterset Weight
	EndCumulativeMetersetWeight = New(0x300C, 0x0009)

	// ReferencedBrachyApplicationSetupSequence (300C,000A) VR=SQ VM=1 Referenced Brachy Application Setup Sequence
	ReferencedBrachyApplicationSetupSequence = New(0x300C, 0x000A)

	// ReferencedBrachyApplicationSetupNumber (300C,000C) VR=IS VM=1 Referenced Brachy Application Setup Number
	ReferencedBrachyApplicationSetupNumber = New(0x300C, 0x000C)

	// ReferencedSourceNumber (300C,000E) VR=IS VM=1 Referenced Source Number
	ReferencedSourceNumber = New(0x300C, 0x000E)

	// ReferencedFractionGroupSequence (300C,0020) VR=SQ VM=1 Referenced Fraction Group Sequence
	ReferencedFractionGroupSequence = New(0x300C, 0x0020)

	// ReferencedFractionGroupNumber (300C,0022) VR=IS VM=1 Referenced Fraction Group Number
	ReferencedFractionGroupNumber = New(0x300C, 0x0022)

	// ReferencedVerificationImageSequence (300C,0040) VR=SQ VM=1 Referenced Verification Image Sequence
	ReferencedVerificationImageSequence = New(0x300C, 0x0040)

	// ReferencedReferenceImageSequence (300C,0042) VR=SQ VM=1 Referenced Reference Image Sequence
	ReferencedReferenceImageSequence = New(0x300C, 0x0042)

	// ReferencedDoseReferenceSequence (300C,0050) VR=SQ VM=1 Referenced Dose Reference Sequence
	ReferencedDoseReferenceSequence = New(0x300C, 0x0050)

	// ReferencedDoseReferenceNumber (300C,0051) VR=IS VM=1 Referenced Dose Reference Number
	ReferencedDoseReferenceNumber = New(0x300C, 0x0051)

	// BrachyReferencedDoseReferenceSequence (300C,0055) VR=SQ VM=1 Brachy Referenced Dose Reference Sequence
	BrachyReferencedDoseReferenceSequence = New(0x300C, 0x0055)

	// ReferencedStructureSetSequence (300C,0060) VR=SQ VM=1 Referenced Structure Set Sequence
	ReferencedStructureSetSequence = New(0x300C, 0x0060)

	// ReferencedPatientSetupNumber (300C,006A) VR=IS VM=1 Referenced Patient Setup Number
	ReferencedPatientSetupNumber = New(0x300C, 0x006A)

	// ReferencedDoseSequence (300C,0080) VR=SQ VM=1 Referenced Dose Sequence
	ReferencedDoseSequence = New(0x300C, 0x0080)

	// ReferencedToleranceTableNumber (300C,00A0) VR=IS VM=1 Referenced Tolerance Table Number
	ReferencedToleranceTableNumber = New(0x300C, 0x00A0)

	// ReferencedBolusSequence (300C,00B0) VR=SQ VM=1 Referenced Bolus Sequence
	ReferencedBolusSequence = New(0x300C, 0x00B0)

	// ReferencedWedgeNumber (300C,00C0) VR=IS VM=1 Referenced Wedge Number
	ReferencedWedgeNumber = New(0x300C, 0x00C0)

	// ReferencedCompensatorNumber (300C,00D0) VR=IS VM=1 Referenced Compensator Number
	ReferencedCompensatorNumber = New(0x300C, 0x00D0)

	// ReferencedBlockNumber (300C,00E0) VR=IS VM=1 Referenced Block Number
	ReferencedBlockNumber = New(0x300C, 0x00E0)

	// ReferencedControlPointIndex (300C,00F0) VR=IS VM=1 Referenced Control Point Index
	ReferencedControlPointIndex = New(0x300C, 0x00F0)

	// ReferencedControlPointSequence (300C,00F2) VR=SQ VM=1 Referenced Control Point Sequence
	ReferencedControlPointSequence = New(0x300C, 0x00F2)

	// ReferencedStartControlPointIndex (300C,00F4) VR=IS VM=1 Referenced Start Control Point Index
	ReferencedStartControlPointIndex = New(0x300C, 0x00F4)

	// ReferencedStopControlPointIndex (300C,00F6) VR=IS VM=1 Referenced Stop Control Point Index
	ReferencedStopControlPointIndex = New(0x300C, 0x00F6)

	// ReferencedRangeShifterNumber (300C,0100) VR=IS VM=1 Referenced Range Shifter Number
	ReferencedRangeShifterNumber = New(0x300C, 0x0100)

	// ReferencedLateralSpreadingDeviceNumber (300C,0102) VR=IS VM=1 Referenced Lateral Spreading Device Number
	ReferencedLateralSpreadingDeviceNumber = New(0x300C, 0x0102)

	// ReferencedRangeModulatorNumber (300C,0104) VR=IS VM=1 Referenced Range Modulator Number
	ReferencedRangeModulatorNumber = New(0x300C, 0x0104)

	// OmittedBeamTaskSequence (300C,0111) VR=SQ VM=1 Omitted Beam Task Sequence
	OmittedBeamTaskSequence = New(0x300C, 0x0111)

	// ReasonForOmission (300C,0112) VR=CS VM=1 Reason for Omission
	ReasonForOmission = New(0x300C, 0x0112)

	// ReasonForOmissionDescription (300C,0113) VR=LO VM=1 Reason for Omission Description
	ReasonForOmissionDescription = New(0x300C, 0x0113)

	// PrescriptionOverviewSequence (300C,0114) VR=SQ VM=1 Prescription Overview Sequence
	PrescriptionOverviewSequence = New(0x300C, 0x0114)

	// TotalPrescriptionDose (300C,0115) VR=FL VM=1 Total Prescription Dose
	TotalPrescriptionDose = New(0x300C, 0x0115)

	// PlanOverviewSequence (300C,0116) VR=SQ VM=1 Plan Overview Sequence
	PlanOverviewSequence = New(0x300C, 0x0116)

	// PlanOverviewIndex (300C,0117) VR=US VM=1 Plan Overview Index
	PlanOverviewIndex = New(0x300C, 0x0117)

	// ReferencedPlanOverviewIndex (300C,0118) VR=US VM=1 Referenced Plan Overview Index
	ReferencedPlanOverviewIndex = New(0x300C, 0x0118)

	// NumberOfFractionsIncluded (300C,0119) VR=US VM=1 Number of Fractions Included
	NumberOfFractionsIncluded = New(0x300C, 0x0119)

	// DoseCalibrationConditionsSequence (300C,0120) VR=SQ VM=1 Dose Calibration Conditions Sequence
	DoseCalibrationConditionsSequence = New(0x300C, 0x0120)

	// AbsorbedDoseToMetersetRatio (300C,0121) VR=FD VM=1 Absorbed Dose to Meterset Ratio
	AbsorbedDoseToMetersetRatio = New(0x300C, 0x0121)

	// DelineatedRadiationFieldSize (300C,0122) VR=FD VM=2 Delineated Radiation Field Size
	DelineatedRadiationFieldSize = New(0x300C, 0x0122)

	// DoseCalibrationConditionsVerifiedFlag (300C,0123) VR=CS VM=1 Dose Calibration Conditions Verified Flag
	DoseCalibrationConditionsVerifiedFlag = New(0x300C, 0x0123)

	// CalibrationReferencePointDepth (300C,0124) VR=FD VM=1 Calibration Reference Point Depth
	CalibrationReferencePointDepth = New(0x300C, 0x0124)

	// GatingBeamHoldTransitionSequence (300C,0125) VR=SQ VM=1 Gating Beam Hold Transition Sequence
	GatingBeamHoldTransitionSequence = New(0x300C, 0x0125)

	// BeamHoldTransition (300C,0126) VR=CS VM=1 Beam Hold Transition
	BeamHoldTransition = New(0x300C, 0x0126)

	// BeamHoldTransitionDateTime (300C,0127) VR=DT VM=1 Beam Hold Transition DateTime
	BeamHoldTransitionDateTime = New(0x300C, 0x0127)

	// BeamHoldOriginatingDeviceSequence (300C,0128) VR=SQ VM=1 Beam Hold Originating Device Sequence
	BeamHoldOriginatingDeviceSequence = New(0x300C, 0x0128)

	// BeamHoldTransitionTriggerSource (300C,0129) VR=CS VM=1 Beam Hold Transition Trigger Source
	BeamHoldTransitionTriggerSource = New(0x300C, 0x0129)

	// ApprovalStatus (300E,0002) VR=CS VM=1 Approval Status
	ApprovalStatus = New(0x300E, 0x0002)

	// ReviewDate (300E,0004) VR=DA VM=1 Review Date
	ReviewDate = New(0x300E, 0x0004)

	// ReviewTime (300E,0005) VR=TM VM=1 Review Time
	ReviewTime = New(0x300E, 0x0005)

	// ReviewerName (300E,0008) VR=PN VM=1 Reviewer Name
	ReviewerName = New(0x300E, 0x0008)

	// RadiobiologicalDoseEffectSequence (3010,0001) VR=SQ VM=1 Radiobiological Dose Effect Sequence
	RadiobiologicalDoseEffectSequence = New(0x3010, 0x0001)

	// RadiobiologicalDoseEffectFlag (3010,0002) VR=CS VM=1 Radiobiological Dose Effect Flag
	RadiobiologicalDoseEffectFlag = New(0x3010, 0x0002)

	// EffectiveDoseCalculationMethodCategoryCodeSequence (3010,0003) VR=SQ VM=1 Effective Dose Calculation Method Category Code Sequence
	EffectiveDoseCalculationMethodCategoryCodeSequence = New(0x3010, 0x0003)

	// EffectiveDoseCalculationMethodCodeSequence (3010,0004) VR=SQ VM=1 Effective Dose Calculation Method Code Sequence
	EffectiveDoseCalculationMethodCodeSequence = New(0x3010, 0x0004)

	// EffectiveDoseCalculationMethodDescription (3010,0005) VR=LO VM=1 Effective Dose Calculation Method Description
	EffectiveDoseCalculationMethodDescription = New(0x3010, 0x0005)

	// ConceptualVolumeUID (3010,0006) VR=UI VM=1 Conceptual Volume UID
	ConceptualVolumeUID = New(0x3010, 0x0006)

	// OriginatingSOPInstanceReferenceSequence (3010,0007) VR=SQ VM=1 Originating SOP Instance Reference Sequence
	OriginatingSOPInstanceReferenceSequence = New(0x3010, 0x0007)

	// ConceptualVolumeConstituentSequence (3010,0008) VR=SQ VM=1 Conceptual Volume Constituent Sequence
	ConceptualVolumeConstituentSequence = New(0x3010, 0x0008)

	// EquivalentConceptualVolumeInstanceReferenceSequence (3010,0009) VR=SQ VM=1 Equivalent Conceptual Volume Instance Reference Sequence
	EquivalentConceptualVolumeInstanceReferenceSequence = New(0x3010, 0x0009)

	// EquivalentConceptualVolumesSequence (3010,000A) VR=SQ VM=1 Equivalent Conceptual Volumes Sequence
	EquivalentConceptualVolumesSequence = New(0x3010, 0x000A)

	// ReferencedConceptualVolumeUID (3010,000B) VR=UI VM=1 Referenced Conceptual Volume UID
	ReferencedConceptualVolumeUID = New(0x3010, 0x000B)

	// ConceptualVolumeCombinationExpression (3010,000C) VR=UT VM=1 Conceptual Volume Combination Expression
	ConceptualVolumeCombinationExpression = New(0x3010, 0x000C)

	// ConceptualVolumeConstituentIndex (3010,000D) VR=US VM=1 Conceptual Volume Constituent Index
	ConceptualVolumeConstituentIndex = New(0x3010, 0x000D)

	// ConceptualVolumeCombinationFlag (3010,000E) VR=CS VM=1 Conceptual Volume Combination Flag
	ConceptualVolumeCombinationFlag = New(0x3010, 0x000E)

	// ConceptualVolumeCombinationDescription (3010,000F) VR=ST VM=1 Conceptual Volume Combination Description
	ConceptualVolumeCombinationDescription = New(0x3010, 0x000F)

	// ConceptualVolumeSegmentationDefinedFlag (3010,0010) VR=CS VM=1 Conceptual Volume Segmentation Defined Flag
	ConceptualVolumeSegmentationDefinedFlag = New(0x3010, 0x0010)

	// ConceptualVolumeSegmentationReferenceSequence (3010,0011) VR=SQ VM=1 Conceptual Volume Segmentation Reference Sequence
	ConceptualVolumeSegmentationReferenceSequence = New(0x3010, 0x0011)

	// ConceptualVolumeConstituentSegmentationReferenceSequence (3010,0012) VR=SQ VM=1 Conceptual Volume Constituent Segmentation Reference Sequence
	ConceptualVolumeConstituentSegmentationReferenceSequence = New(0x3010, 0x0012)

	// ConstituentConceptualVolumeUID (3010,0013) VR=UI VM=1 Constituent Conceptual Volume UID
	ConstituentConceptualVolumeUID = New(0x3010, 0x0013)

	// DerivationConceptualVolumeSequence (3010,0014) VR=SQ VM=1 Derivation Conceptual Volume Sequence
	DerivationConceptualVolumeSequence = New(0x3010, 0x0014)

	// SourceConceptualVolumeUID (3010,0015) VR=UI VM=1 Source Conceptual Volume UID
	SourceConceptualVolumeUID = New(0x3010, 0x0015)

	// ConceptualVolumeDerivationAlgorithmSequence (3010,0016) VR=SQ VM=1 Conceptual Volume Derivation Algorithm Sequence
	ConceptualVolumeDerivationAlgorithmSequence = New(0x3010, 0x0016)

	// ConceptualVolumeDescription (3010,0017) VR=ST VM=1 Conceptual Volume Description
	ConceptualVolumeDescription = New(0x3010, 0x0017)

	// SourceConceptualVolumeSequence (3010,0018) VR=SQ VM=1 Source Conceptual Volume Sequence
	SourceConceptualVolumeSequence = New(0x3010, 0x0018)

	// AuthorIdentificationSequence (3010,0019) VR=SQ VM=1 Author Identification Sequence
	AuthorIdentificationSequence = New(0x3010, 0x0019)

	// ManufacturerModelVersion (3010,001A) VR=LO VM=1 Manufacturer's Model Version
	ManufacturerModelVersion = New(0x3010, 0x001A)

	// DeviceAlternateIdentifier (3010,001B) VR=UC VM=1 Device Alternate Identifier
	DeviceAlternateIdentifier = New(0x3010, 0x001B)

	// DeviceAlternateIdentifierType (3010,001C) VR=CS VM=1 Device Alternate Identifier Type
	DeviceAlternateIdentifierType = New(0x3010, 0x001C)

	// DeviceAlternateIdentifierFormat (3010,001D) VR=LT VM=1 Device Alternate Identifier Format
	DeviceAlternateIdentifierFormat = New(0x3010, 0x001D)

	// SegmentationCreationTemplateLabel (3010,001E) VR=LO VM=1 Segmentation Creation Template Label
	SegmentationCreationTemplateLabel = New(0x3010, 0x001E)

	// SegmentationTemplateUID (3010,001F) VR=UI VM=1 Segmentation Template UID
	SegmentationTemplateUID = New(0x3010, 0x001F)

	// ReferencedSegmentReferenceIndex (3010,0020) VR=US VM=1 Referenced Segment Reference Index
	ReferencedSegmentReferenceIndex = New(0x3010, 0x0020)

	// SegmentReferenceSequence (3010,0021) VR=SQ VM=1 Segment Reference Sequence
	SegmentReferenceSequence = New(0x3010, 0x0021)

	// SegmentReferenceIndex (3010,0022) VR=US VM=1 Segment Reference Index
	SegmentReferenceIndex = New(0x3010, 0x0022)

	// DirectSegmentReferenceSequence (3010,0023) VR=SQ VM=1 Direct Segment Reference Sequence
	DirectSegmentReferenceSequence = New(0x3010, 0x0023)

	// CombinationSegmentReferenceSequence (3010,0024) VR=SQ VM=1 Combination Segment Reference Sequence
	CombinationSegmentReferenceSequence = New(0x3010, 0x0024)

	// ConceptualVolumeSequence (3010,0025) VR=SQ VM=1 Conceptual Volume Sequence
	ConceptualVolumeSequence = New(0x3010, 0x0025)

	// SegmentedRTAccessoryDeviceSequence (3010,0026) VR=SQ VM=1 Segmented RT Accessory Device Sequence
	SegmentedRTAccessoryDeviceSequence = New(0x3010, 0x0026)

	// SegmentCharacteristicsSequence (3010,0027) VR=SQ VM=1 Segment Characteristics Sequence
	SegmentCharacteristicsSequence = New(0x3010, 0x0027)

	// RelatedSegmentCharacteristicsSequence (3010,0028) VR=SQ VM=1 Related Segment Characteristics Sequence
	RelatedSegmentCharacteristicsSequence = New(0x3010, 0x0028)

	// SegmentCharacteristicsPrecedence (3010,0029) VR=US VM=1 Segment Characteristics Precedence
	SegmentCharacteristicsPrecedence = New(0x3010, 0x0029)

	// RTSegmentAnnotationSequence (3010,002A) VR=SQ VM=1 RT Segment Annotation Sequence
	RTSegmentAnnotationSequence = New(0x3010, 0x002A)

	// SegmentAnnotationCategoryCodeSequence (3010,002B) VR=SQ VM=1 Segment Annotation Category Code Sequence
	SegmentAnnotationCategoryCodeSequence = New(0x3010, 0x002B)

	// SegmentAnnotationTypeCodeSequence (3010,002C) VR=SQ VM=1 Segment Annotation Type Code Sequence
	SegmentAnnotationTypeCodeSequence = New(0x3010, 0x002C)

	// DeviceLabel (3010,002D) VR=LO VM=1 Device Label
	DeviceLabel = New(0x3010, 0x002D)

	// DeviceTypeCodeSequence (3010,002E) VR=SQ VM=1 Device Type Code Sequence
	DeviceTypeCodeSequence = New(0x3010, 0x002E)

	// SegmentAnnotationTypeModifierCodeSequence (3010,002F) VR=SQ VM=1 Segment Annotation Type Modifier Code Sequence
	SegmentAnnotationTypeModifierCodeSequence = New(0x3010, 0x002F)

	// PatientEquipmentRelationshipCodeSequence (3010,0030) VR=SQ VM=1 Patient Equipment Relationship Code Sequence
	PatientEquipmentRelationshipCodeSequence = New(0x3010, 0x0030)

	// ReferencedFiducialsUID (3010,0031) VR=UI VM=1 Referenced Fiducials UID
	ReferencedFiducialsUID = New(0x3010, 0x0031)

	// PatientTreatmentOrientationSequence (3010,0032) VR=SQ VM=1 Patient Treatment Orientation Sequence
	PatientTreatmentOrientationSequence = New(0x3010, 0x0032)

	// UserContentLabel (3010,0033) VR=SH VM=1 User Content Label
	UserContentLabel = New(0x3010, 0x0033)

	// UserContentLongLabel (3010,0034) VR=LO VM=1 User Content Long Label
	UserContentLongLabel = New(0x3010, 0x0034)

	// EntityLabel (3010,0035) VR=SH VM=1 Entity Label
	EntityLabel = New(0x3010, 0x0035)

	// EntityName (3010,0036) VR=LO VM=1 Entity Name
	EntityName = New(0x3010, 0x0036)

	// EntityDescription (3010,0037) VR=ST VM=1 Entity Description
	EntityDescription = New(0x3010, 0x0037)

	// EntityLongLabel (3010,0038) VR=LO VM=1 Entity Long Label
	EntityLongLabel = New(0x3010, 0x0038)

	// DeviceIndex (3010,0039) VR=US VM=1 Device Index
	DeviceIndex = New(0x3010, 0x0039)

	// RTTreatmentPhaseIndex (3010,003A) VR=US VM=1 RT Treatment Phase Index
	RTTreatmentPhaseIndex = New(0x3010, 0x003A)

	// RTTreatmentPhaseUID (3010,003B) VR=UI VM=1 RT Treatment Phase UID
	RTTreatmentPhaseUID = New(0x3010, 0x003B)

	// RTPrescriptionIndex (3010,003C) VR=US VM=1 RT Prescription Index
	RTPrescriptionIndex = New(0x3010, 0x003C)

	// RTSegmentAnnotationIndex (3010,003D) VR=US VM=1 RT Segment Annotation Index
	RTSegmentAnnotationIndex = New(0x3010, 0x003D)

	// BasisRTTreatmentPhaseIndex (3010,003E) VR=US VM=1 Basis RT Treatment Phase Index
	BasisRTTreatmentPhaseIndex = New(0x3010, 0x003E)

	// RelatedRTTreatmentPhaseIndex (3010,003F) VR=US VM=1 Related RT Treatment Phase Index
	RelatedRTTreatmentPhaseIndex = New(0x3010, 0x003F)

	// ReferencedRTTreatmentPhaseIndex (3010,0040) VR=US VM=1 Referenced RT Treatment Phase Index
	ReferencedRTTreatmentPhaseIndex = New(0x3010, 0x0040)

	// ReferencedRTPrescriptionIndex (3010,0041) VR=US VM=1 Referenced RT Prescription Index
	ReferencedRTPrescriptionIndex = New(0x3010, 0x0041)

	// ReferencedParentRTPrescriptionIndex (3010,0042) VR=US VM=1 Referenced Parent RT Prescription Index
	ReferencedParentRTPrescriptionIndex = New(0x3010, 0x0042)

	// ManufacturerDeviceIdentifier (3010,0043) VR=ST VM=1 Manufacturer's Device Identifier
	ManufacturerDeviceIdentifier = New(0x3010, 0x0043)

	// InstanceLevelReferencedPerformedProcedureStepSequence (3010,0044) VR=SQ VM=1 Instance-Level Referenced Performed Procedure Step Sequence
	InstanceLevelReferencedPerformedProcedureStepSequence = New(0x3010, 0x0044)

	// RTTreatmentPhaseIntentPresenceFlag (3010,0045) VR=CS VM=1 RT Treatment Phase Intent Presence Flag
	RTTreatmentPhaseIntentPresenceFlag = New(0x3010, 0x0045)

	// RadiotherapyTreatmentType (3010,0046) VR=CS VM=1 Radiotherapy Treatment Type
	RadiotherapyTreatmentType = New(0x3010, 0x0046)

	// TeletherapyRadiationType (3010,0047) VR=CS VM=1-n Teletherapy Radiation Type
	TeletherapyRadiationType = New(0x3010, 0x0047)

	// BrachytherapySourceType (3010,0048) VR=CS VM=1-n Brachytherapy Source Type
	BrachytherapySourceType = New(0x3010, 0x0048)

	// ReferencedRTTreatmentPhaseSequence (3010,0049) VR=SQ VM=1 Referenced RT Treatment Phase Sequence
	ReferencedRTTreatmentPhaseSequence = New(0x3010, 0x0049)

	// ReferencedDirectSegmentInstanceSequence (3010,004A) VR=SQ VM=1 Referenced Direct Segment Instance Sequence
	ReferencedDirectSegmentInstanceSequence = New(0x3010, 0x004A)

	// IntendedRTTreatmentPhaseSequence (3010,004B) VR=SQ VM=1 Intended RT Treatment Phase Sequence
	IntendedRTTreatmentPhaseSequence = New(0x3010, 0x004B)

	// IntendedPhaseStartDate (3010,004C) VR=DA VM=1 Intended Phase Start Date
	IntendedPhaseStartDate = New(0x3010, 0x004C)

	// IntendedPhaseEndDate (3010,004D) VR=DA VM=1 Intended Phase End Date
	IntendedPhaseEndDate = New(0x3010, 0x004D)

	// RTTreatmentPhaseIntervalSequence (3010,004E) VR=SQ VM=1 RT Treatment Phase Interval Sequence
	RTTreatmentPhaseIntervalSequence = New(0x3010, 0x004E)

	// TemporalRelationshipIntervalAnchor (3010,004F) VR=CS VM=1 Temporal Relationship Interval Anchor
	TemporalRelationshipIntervalAnchor = New(0x3010, 0x004F)

	// MinimumNumberOfIntervalDays (3010,0050) VR=FD VM=1 Minimum Number of Interval Days
	MinimumNumberOfIntervalDays = New(0x3010, 0x0050)

	// MaximumNumberOfIntervalDays (3010,0051) VR=FD VM=1 Maximum Number of Interval Days
	MaximumNumberOfIntervalDays = New(0x3010, 0x0051)

	// PertinentSOPClassesInStudy (3010,0052) VR=UI VM=1-n Pertinent SOP Classes in Study
	PertinentSOPClassesInStudy = New(0x3010, 0x0052)

	// PertinentSOPClassesInSeries (3010,0053) VR=UI VM=1-n Pertinent SOP Classes in Series
	PertinentSOPClassesInSeries = New(0x3010, 0x0053)

	// RTPrescriptionLabel (3010,0054) VR=LO VM=1 RT Prescription Label
	RTPrescriptionLabel = New(0x3010, 0x0054)

	// RTPhysicianIntentPredecessorSequence (3010,0055) VR=SQ VM=1 RT Physician Intent Predecessor Sequence
	RTPhysicianIntentPredecessorSequence = New(0x3010, 0x0055)

	// RTTreatmentApproachLabel (3010,0056) VR=LO VM=1 RT Treatment Approach Label
	RTTreatmentApproachLabel = New(0x3010, 0x0056)

	// RTPhysicianIntentSequence (3010,0057) VR=SQ VM=1 RT Physician Intent Sequence
	RTPhysicianIntentSequence = New(0x3010, 0x0057)

	// RTPhysicianIntentIndex (3010,0058) VR=US VM=1 RT Physician Intent Index
	RTPhysicianIntentIndex = New(0x3010, 0x0058)

	// RTTreatmentIntentType (3010,0059) VR=CS VM=1 RT Treatment Intent Type
	RTTreatmentIntentType = New(0x3010, 0x0059)

	// RTPhysicianIntentNarrative (3010,005A) VR=UT VM=1 RT Physician Intent Narrative
	RTPhysicianIntentNarrative = New(0x3010, 0x005A)

	// RTProtocolCodeSequence (3010,005B) VR=SQ VM=1 RT Protocol Code Sequence
	RTProtocolCodeSequence = New(0x3010, 0x005B)

	// ReasonForSuperseding (3010,005C) VR=ST VM=1 Reason for Superseding
	ReasonForSuperseding = New(0x3010, 0x005C)

	// RTDiagnosisCodeSequence (3010,005D) VR=SQ VM=1 RT Diagnosis Code Sequence
	RTDiagnosisCodeSequence = New(0x3010, 0x005D)

	// ReferencedRTPhysicianIntentIndex (3010,005E) VR=US VM=1 Referenced RT Physician Intent Index
	ReferencedRTPhysicianIntentIndex = New(0x3010, 0x005E)

	// RTPhysicianIntentInputInstanceSequence (3010,005F) VR=SQ VM=1 RT Physician Intent Input Instance Sequence
	RTPhysicianIntentInputInstanceSequence = New(0x3010, 0x005F)

	// RTAnatomicPrescriptionSequence (3010,0060) VR=SQ VM=1 RT Anatomic Prescription Sequence
	RTAnatomicPrescriptionSequence = New(0x3010, 0x0060)

	// PriorTreatmentDoseDescription (3010,0061) VR=UT VM=1 Prior Treatment Dose Description
	PriorTreatmentDoseDescription = New(0x3010, 0x0061)

	// PriorTreatmentReferenceSequence (3010,0062) VR=SQ VM=1 Prior Treatment Reference Sequence
	PriorTreatmentReferenceSequence = New(0x3010, 0x0062)

	// DosimetricObjectiveEvaluationScope (3010,0063) VR=CS VM=1 Dosimetric Objective Evaluation Scope
	DosimetricObjectiveEvaluationScope = New(0x3010, 0x0063)

	// TherapeuticRoleCategoryCodeSequence (3010,0064) VR=SQ VM=1 Therapeutic Role Category Code Sequence
	TherapeuticRoleCategoryCodeSequence = New(0x3010, 0x0064)

	// TherapeuticRoleTypeCodeSequence (3010,0065) VR=SQ VM=1 Therapeutic Role Type Code Sequence
	TherapeuticRoleTypeCodeSequence = New(0x3010, 0x0065)

	// ConceptualVolumeOptimizationPrecedence (3010,0066) VR=US VM=1 Conceptual Volume Optimization Precedence
	ConceptualVolumeOptimizationPrecedence = New(0x3010, 0x0066)

	// ConceptualVolumeCategoryCodeSequence (3010,0067) VR=SQ VM=1 Conceptual Volume Category Code Sequence
	ConceptualVolumeCategoryCodeSequence = New(0x3010, 0x0067)

	// ConceptualVolumeBlockingConstraint (3010,0068) VR=CS VM=1 Conceptual Volume Blocking Constraint
	ConceptualVolumeBlockingConstraint = New(0x3010, 0x0068)

	// ConceptualVolumeTypeCodeSequence (3010,0069) VR=SQ VM=1 Conceptual Volume Type Code Sequence
	ConceptualVolumeTypeCodeSequence = New(0x3010, 0x0069)

	// ConceptualVolumeTypeModifierCodeSequence (3010,006A) VR=SQ VM=1 Conceptual Volume Type Modifier Code Sequence
	ConceptualVolumeTypeModifierCodeSequence = New(0x3010, 0x006A)

	// RTPrescriptionSequence (3010,006B) VR=SQ VM=1 RT Prescription Sequence
	RTPrescriptionSequence = New(0x3010, 0x006B)

	// DosimetricObjectiveSequence (3010,006C) VR=SQ VM=1 Dosimetric Objective Sequence
	DosimetricObjectiveSequence = New(0x3010, 0x006C)

	// DosimetricObjectiveTypeCodeSequence (3010,006D) VR=SQ VM=1 Dosimetric Objective Type Code Sequence
	DosimetricObjectiveTypeCodeSequence = New(0x3010, 0x006D)

	// DosimetricObjectiveUID (3010,006E) VR=UI VM=1 Dosimetric Objective UID
	DosimetricObjectiveUID = New(0x3010, 0x006E)

	// ReferencedDosimetricObjectiveUID (3010,006F) VR=UI VM=1 Referenced Dosimetric Objective UID
	ReferencedDosimetricObjectiveUID = New(0x3010, 0x006F)

	// DosimetricObjectiveParameterSequence (3010,0070) VR=SQ VM=1 Dosimetric Objective Parameter Sequence
	DosimetricObjectiveParameterSequence = New(0x3010, 0x0070)

	// ReferencedDosimetricObjectivesSequence (3010,0071) VR=SQ VM=1 Referenced Dosimetric Objectives Sequence
	ReferencedDosimetricObjectivesSequence = New(0x3010, 0x0071)

	// AbsoluteDosimetricObjectiveFlag (3010,0073) VR=CS VM=1 Absolute Dosimetric Objective Flag
	AbsoluteDosimetricObjectiveFlag = New(0x3010, 0x0073)

	// DosimetricObjectiveWeight (3010,0074) VR=FD VM=1 Dosimetric Objective Weight
	DosimetricObjectiveWeight = New(0x3010, 0x0074)

	// DosimetricObjectivePurpose (3010,0075) VR=CS VM=1 Dosimetric Objective Purpose
	DosimetricObjectivePurpose = New(0x3010, 0x0075)

	// PlanningInputInformationSequence (3010,0076) VR=SQ VM=1 Planning Input Information Sequence
	PlanningInputInformationSequence = New(0x3010, 0x0076)

	// TreatmentSite (3010,0077) VR=LO VM=1 Treatment Site
	TreatmentSite = New(0x3010, 0x0077)

	// TreatmentSiteCodeSequence (3010,0078) VR=SQ VM=1 Treatment Site Code Sequence
	TreatmentSiteCodeSequence = New(0x3010, 0x0078)

	// FractionPatternSequence (3010,0079) VR=SQ VM=1 Fraction Pattern Sequence
	FractionPatternSequence = New(0x3010, 0x0079)

	// TreatmentTechniqueNotes (3010,007A) VR=UT VM=1 Treatment Technique Notes
	TreatmentTechniqueNotes = New(0x3010, 0x007A)

	// PrescriptionNotes (3010,007B) VR=UT VM=1 Prescription Notes
	PrescriptionNotes = New(0x3010, 0x007B)

	// NumberOfIntervalFractions (3010,007C) VR=IS VM=1 Number of Interval Fractions
	NumberOfIntervalFractions = New(0x3010, 0x007C)

	// NumberOfFractions (3010,007D) VR=US VM=1 Number of Fractions
	NumberOfFractions = New(0x3010, 0x007D)

	// IntendedDeliveryDuration (3010,007E) VR=US VM=1 Intended Delivery Duration
	IntendedDeliveryDuration = New(0x3010, 0x007E)

	// FractionationNotes (3010,007F) VR=UT VM=1 Fractionation Notes
	FractionationNotes = New(0x3010, 0x007F)

	// RTTreatmentTechniqueCodeSequence (3010,0080) VR=SQ VM=1 RT Treatment Technique Code Sequence
	RTTreatmentTechniqueCodeSequence = New(0x3010, 0x0080)

	// PrescriptionNotesSequence (3010,0081) VR=SQ VM=1 Prescription Notes Sequence
	PrescriptionNotesSequence = New(0x3010, 0x0081)

	// FractionBasedRelationshipSequence (3010,0082) VR=SQ VM=1 Fraction-Based Relationship Sequence
	FractionBasedRelationshipSequence = New(0x3010, 0x0082)

	// FractionBasedRelationshipIntervalAnchor (3010,0083) VR=CS VM=1 Fraction-Based Relationship Interval Anchor
	FractionBasedRelationshipIntervalAnchor = New(0x3010, 0x0083)

	// MinimumHoursBetweenFractions (3010,0084) VR=FD VM=1 Minimum Hours between Fractions
	MinimumHoursBetweenFractions = New(0x3010, 0x0084)

	// IntendedFractionStartTime (3010,0085) VR=TM VM=1-n Intended Fraction Start Time
	IntendedFractionStartTime = New(0x3010, 0x0085)

	// IntendedStartDayOfWeek (3010,0086) VR=LT VM=1 Intended Start Day of Week
	IntendedStartDayOfWeek = New(0x3010, 0x0086)

	// WeekdayFractionPatternSequence (3010,0087) VR=SQ VM=1 Weekday Fraction Pattern Sequence
	WeekdayFractionPatternSequence = New(0x3010, 0x0087)

	// DeliveryTimeStructureCodeSequence (3010,0088) VR=SQ VM=1 Delivery Time Structure Code Sequence
	DeliveryTimeStructureCodeSequence = New(0x3010, 0x0088)

	// TreatmentSiteModifierCodeSequence (3010,0089) VR=SQ VM=1 Treatment Site Modifier Code Sequence
	TreatmentSiteModifierCodeSequence = New(0x3010, 0x0089)

	// RoboticBaseLocationIndicatorRETIRED (3010,0090) VR=CS VM=1 Robotic Base Location Indicator (RETIRED)
	RoboticBaseLocationIndicatorRETIRED = New(0x3010, 0x0090)

	// RoboticPathNodeSetCodeSequence (3010,0091) VR=SQ VM=1 Robotic Path Node Set Code Sequence
	RoboticPathNodeSetCodeSequence = New(0x3010, 0x0091)

	// RoboticNodeIdentifier (3010,0092) VR=UL VM=1 Robotic Node Identifier
	RoboticNodeIdentifier = New(0x3010, 0x0092)

	// RTTreatmentSourceCoordinates (3010,0093) VR=FD VM=3 RT Treatment Source Coordinates
	RTTreatmentSourceCoordinates = New(0x3010, 0x0093)

	// RadiationSourceCoordinateSystemYawAngle (3010,0094) VR=FD VM=1 Radiation Source Coordinate SystemYaw Angle
	RadiationSourceCoordinateSystemYawAngle = New(0x3010, 0x0094)

	// RadiationSourceCoordinateSystemRollAngle (3010,0095) VR=FD VM=1 Radiation Source Coordinate SystemRoll Angle
	RadiationSourceCoordinateSystemRollAngle = New(0x3010, 0x0095)

	// RadiationSourceCoordinateSystemPitchAngle (3010,0096) VR=FD VM=1 Radiation Source Coordinate System Pitch Angle
	RadiationSourceCoordinateSystemPitchAngle = New(0x3010, 0x0096)

	// RoboticPathControlPointSequence (3010,0097) VR=SQ VM=1 Robotic Path Control Point Sequence
	RoboticPathControlPointSequence = New(0x3010, 0x0097)

	// TomotherapeuticControlPointSequence (3010,0098) VR=SQ VM=1 Tomotherapeutic Control Point Sequence
	TomotherapeuticControlPointSequence = New(0x3010, 0x0098)

	// TomotherapeuticLeafOpenDurations (3010,0099) VR=FD VM=1-n Tomotherapeutic Leaf Open Durations
	TomotherapeuticLeafOpenDurations = New(0x3010, 0x0099)

	// TomotherapeuticLeafInitialClosedDurations (3010,009A) VR=FD VM=1-n Tomotherapeutic Leaf Initial Closed Durations
	TomotherapeuticLeafInitialClosedDurations = New(0x3010, 0x009A)

	// ConceptualVolumeIdentificationSequence (3010,00A0) VR=SQ VM=1 Conceptual Volume Identification Sequence
	ConceptualVolumeIdentificationSequence = New(0x3010, 0x00A0)

	// ArbitraryRETIRED (4000,0010) VR=LT VM=1 Arbitrary (RETIRED)
	ArbitraryRETIRED = New(0x4000, 0x0010)

	// TextCommentsRETIRED (4000,4000) VR=LT VM=1 Text Comments (RETIRED)
	TextCommentsRETIRED = New(0x4000, 0x4000)

	// ResultsIDRETIRED (4008,0040) VR=SH VM=1 Results ID (RETIRED)
	ResultsIDRETIRED = New(0x4008, 0x0040)

	// ResultsIDIssuerRETIRED (4008,0042) VR=LO VM=1 Results ID Issuer (RETIRED)
	ResultsIDIssuerRETIRED = New(0x4008, 0x0042)

	// ReferencedInterpretationSequenceRETIRED (4008,0050) VR=SQ VM=1 Referenced Interpretation Sequence (RETIRED)
	ReferencedInterpretationSequenceRETIRED = New(0x4008, 0x0050)

	// ReportProductionStatusTrialRETIRED (4008,00FF) VR=CS VM=1 Report Production Status (Trial) (RETIRED)
	ReportProductionStatusTrialRETIRED = New(0x4008, 0x00FF)

	// InterpretationRecordedDateRETIRED (4008,0100) VR=DA VM=1 Interpretation Recorded Date (RETIRED)
	InterpretationRecordedDateRETIRED = New(0x4008, 0x0100)

	// InterpretationRecordedTimeRETIRED (4008,0101) VR=TM VM=1 Interpretation Recorded Time (RETIRED)
	InterpretationRecordedTimeRETIRED = New(0x4008, 0x0101)

	// InterpretationRecorderRETIRED (4008,0102) VR=PN VM=1 Interpretation Recorder (RETIRED)
	InterpretationRecorderRETIRED = New(0x4008, 0x0102)

	// ReferenceToRecordedSoundRETIRED (4008,0103) VR=LO VM=1 Reference to Recorded Sound (RETIRED)
	ReferenceToRecordedSoundRETIRED = New(0x4008, 0x0103)

	// InterpretationTranscriptionDateRETIRED (4008,0108) VR=DA VM=1 Interpretation Transcription Date (RETIRED)
	InterpretationTranscriptionDateRETIRED = New(0x4008, 0x0108)

	// InterpretationTranscriptionTimeRETIRED (4008,0109) VR=TM VM=1 Interpretation Transcription Time (RETIRED)
	InterpretationTranscriptionTimeRETIRED = New(0x4008, 0x0109)

	// InterpretationTranscriberRETIRED (4008,010A) VR=PN VM=1 Interpretation Transcriber (RETIRED)
	InterpretationTranscriberRETIRED = New(0x4008, 0x010A)

	// InterpretationTextRETIRED (4008,010B) VR=ST VM=1 Interpretation Text (RETIRED)
	InterpretationTextRETIRED = New(0x4008, 0x010B)

	// InterpretationAuthorRETIRED (4008,010C) VR=PN VM=1 Interpretation Author (RETIRED)
	InterpretationAuthorRETIRED = New(0x4008, 0x010C)

	// InterpretationApproverSequenceRETIRED (4008,0111) VR=SQ VM=1 Interpretation Approver Sequence (RETIRED)
	InterpretationApproverSequenceRETIRED = New(0x4008, 0x0111)

	// InterpretationApprovalDateRETIRED (4008,0112) VR=DA VM=1 Interpretation Approval Date (RETIRED)
	InterpretationApprovalDateRETIRED = New(0x4008, 0x0112)

	// InterpretationApprovalTimeRETIRED (4008,0113) VR=TM VM=1 Interpretation Approval Time (RETIRED)
	InterpretationApprovalTimeRETIRED = New(0x4008, 0x0113)

	// PhysicianApprovingInterpretationRETIRED (4008,0114) VR=PN VM=1 Physician Approving Interpretation (RETIRED)
	PhysicianApprovingInterpretationRETIRED = New(0x4008, 0x0114)

	// InterpretationDiagnosisDescriptionRETIRED (4008,0115) VR=LT VM=1 Interpretation Diagnosis Description (RETIRED)
	InterpretationDiagnosisDescriptionRETIRED = New(0x4008, 0x0115)

	// InterpretationDiagnosisCodeSequenceRETIRED (4008,0117) VR=SQ VM=1 Interpretation Diagnosis Code Sequence (RETIRED)
	InterpretationDiagnosisCodeSequenceRETIRED = New(0x4008, 0x0117)

	// ResultsDistributionListSequenceRETIRED (4008,0118) VR=SQ VM=1 Results Distribution List Sequence (RETIRED)
	ResultsDistributionListSequenceRETIRED = New(0x4008, 0x0118)

	// DistributionNameRETIRED (4008,0119) VR=PN VM=1 Distribution Name (RETIRED)
	DistributionNameRETIRED = New(0x4008, 0x0119)

	// DistributionAddressRETIRED (4008,011A) VR=LO VM=1 Distribution Address (RETIRED)
	DistributionAddressRETIRED = New(0x4008, 0x011A)

	// InterpretationIDRETIRED (4008,0200) VR=SH VM=1 Interpretation ID (RETIRED)
	InterpretationIDRETIRED = New(0x4008, 0x0200)

	// InterpretationIDIssuerRETIRED (4008,0202) VR=LO VM=1 Interpretation ID Issuer (RETIRED)
	InterpretationIDIssuerRETIRED = New(0x4008, 0x0202)

	// InterpretationTypeIDRETIRED (4008,0210) VR=CS VM=1 Interpretation Type ID (RETIRED)
	InterpretationTypeIDRETIRED = New(0x4008, 0x0210)

	// InterpretationStatusIDRETIRED (4008,0212) VR=CS VM=1 Interpretation Status ID (RETIRED)
	InterpretationStatusIDRETIRED = New(0x4008, 0x0212)

	// ImpressionsRETIRED (4008,0300) VR=ST VM=1 Impressions (RETIRED)
	ImpressionsRETIRED = New(0x4008, 0x0300)

	// ResultsCommentsRETIRED (4008,4000) VR=ST VM=1 Results Comments (RETIRED)
	ResultsCommentsRETIRED = New(0x4008, 0x4000)

	// LowEnergyDetectors (4010,0001) VR=CS VM=1 Low Energy Detectors
	LowEnergyDetectors = New(0x4010, 0x0001)

	// HighEnergyDetectors (4010,0002) VR=CS VM=1 High Energy Detectors
	HighEnergyDetectors = New(0x4010, 0x0002)

	// DetectorGeometrySequence (4010,0004) VR=SQ VM=1 Detector Geometry Sequence
	DetectorGeometrySequence = New(0x4010, 0x0004)

	// ThreatROIVoxelSequence (4010,1001) VR=SQ VM=1 Threat ROI Voxel Sequence
	ThreatROIVoxelSequence = New(0x4010, 0x1001)

	// ThreatROIBase (4010,1004) VR=FL VM=3 Threat ROI Base
	ThreatROIBase = New(0x4010, 0x1004)

	// ThreatROIExtents (4010,1005) VR=FL VM=3 Threat ROI Extents
	ThreatROIExtents = New(0x4010, 0x1005)

	// ThreatROIBitmap (4010,1006) VR=OB VM=1 Threat ROI Bitmap
	ThreatROIBitmap = New(0x4010, 0x1006)

	// RouteSegmentID (4010,1007) VR=SH VM=1 Route Segment ID
	RouteSegmentID = New(0x4010, 0x1007)

	// GantryType (4010,1008) VR=CS VM=1 Gantry Type
	GantryType = New(0x4010, 0x1008)

	// OOIOwnerType (4010,1009) VR=CS VM=1 OOI Owner Type
	OOIOwnerType = New(0x4010, 0x1009)

	// RouteSegmentSequence (4010,100A) VR=SQ VM=1 Route Segment Sequence
	RouteSegmentSequence = New(0x4010, 0x100A)

	// PotentialThreatObjectID (4010,1010) VR=US VM=1 Potential Threat Object ID
	PotentialThreatObjectID = New(0x4010, 0x1010)

	// ThreatSequence (4010,1011) VR=SQ VM=1 Threat Sequence
	ThreatSequence = New(0x4010, 0x1011)

	// ThreatCategory (4010,1012) VR=CS VM=1 Threat Category
	ThreatCategory = New(0x4010, 0x1012)

	// ThreatCategoryDescription (4010,1013) VR=LT VM=1 Threat Category Description
	ThreatCategoryDescription = New(0x4010, 0x1013)

	// ATDAbilityAssessment (4010,1014) VR=CS VM=1 ATD Ability Assessment
	ATDAbilityAssessment = New(0x4010, 0x1014)

	// ATDAssessmentFlag (4010,1015) VR=CS VM=1 ATD Assessment Flag
	ATDAssessmentFlag = New(0x4010, 0x1015)

	// ATDAssessmentProbability (4010,1016) VR=FL VM=1 ATD Assessment Probability
	ATDAssessmentProbability = New(0x4010, 0x1016)

	// Mass (4010,1017) VR=FL VM=1 Mass
	Mass = New(0x4010, 0x1017)

	// Density (4010,1018) VR=FL VM=1 Density
	Density = New(0x4010, 0x1018)

	// ZEffective (4010,1019) VR=FL VM=1 Z Effective
	ZEffective = New(0x4010, 0x1019)

	// BoardingPassID (4010,101A) VR=SH VM=1 Boarding Pass ID
	BoardingPassID = New(0x4010, 0x101A)

	// CenterOfMass (4010,101B) VR=FL VM=3 Center of Mass
	CenterOfMass = New(0x4010, 0x101B)

	// CenterOfPTO (4010,101C) VR=FL VM=3 Center of PTO
	CenterOfPTO = New(0x4010, 0x101C)

	// BoundingPolygon (4010,101D) VR=FL VM=6-n Bounding Polygon
	BoundingPolygon = New(0x4010, 0x101D)

	// RouteSegmentStartLocationID (4010,101E) VR=SH VM=1 Route Segment Start Location ID
	RouteSegmentStartLocationID = New(0x4010, 0x101E)

	// RouteSegmentEndLocationID (4010,101F) VR=SH VM=1 Route Segment End Location ID
	RouteSegmentEndLocationID = New(0x4010, 0x101F)

	// RouteSegmentLocationIDType (4010,1020) VR=CS VM=1 Route Segment Location ID Type
	RouteSegmentLocationIDType = New(0x4010, 0x1020)

	// AbortReason (4010,1021) VR=CS VM=1-n Abort Reason
	AbortReason = New(0x4010, 0x1021)

	// VolumeOfPTO (4010,1023) VR=FL VM=1 Volume of PTO
	VolumeOfPTO = New(0x4010, 0x1023)

	// AbortFlag (4010,1024) VR=CS VM=1 Abort Flag
	AbortFlag = New(0x4010, 0x1024)

	// RouteSegmentStartTime (4010,1025) VR=DT VM=1 Route Segment Start Time
	RouteSegmentStartTime = New(0x4010, 0x1025)

	// RouteSegmentEndTime (4010,1026) VR=DT VM=1 Route Segment End Time
	RouteSegmentEndTime = New(0x4010, 0x1026)

	// TDRType (4010,1027) VR=CS VM=1 TDR Type
	TDRType = New(0x4010, 0x1027)

	// InternationalRouteSegment (4010,1028) VR=CS VM=1 International Route Segment
	InternationalRouteSegment = New(0x4010, 0x1028)

	// ThreatDetectionAlgorithmAndVersion (4010,1029) VR=LO VM=1-n Threat Detection Algorithm and Version
	ThreatDetectionAlgorithmAndVersion = New(0x4010, 0x1029)

	// AssignedLocation (4010,102A) VR=SH VM=1 Assigned Location
	AssignedLocation = New(0x4010, 0x102A)

	// AlarmDecisionTime (4010,102B) VR=DT VM=1 Alarm Decision Time
	AlarmDecisionTime = New(0x4010, 0x102B)

	// AlarmDecision (4010,1031) VR=CS VM=1 Alarm Decision
	AlarmDecision = New(0x4010, 0x1031)

	// NumberOfTotalObjects (4010,1033) VR=US VM=1 Number of Total Objects
	NumberOfTotalObjects = New(0x4010, 0x1033)

	// NumberOfAlarmObjects (4010,1034) VR=US VM=1 Number of Alarm Objects
	NumberOfAlarmObjects = New(0x4010, 0x1034)

	// PTORepresentationSequence (4010,1037) VR=SQ VM=1 PTO Representation Sequence
	PTORepresentationSequence = New(0x4010, 0x1037)

	// ATDAssessmentSequence (4010,1038) VR=SQ VM=1 ATD Assessment Sequence
	ATDAssessmentSequence = New(0x4010, 0x1038)

	// TIPType (4010,1039) VR=CS VM=1 TIP Type
	TIPType = New(0x4010, 0x1039)

	// DICOSVersion (4010,103A) VR=CS VM=1 DICOS Version
	DICOSVersion = New(0x4010, 0x103A)

	// OOIOwnerCreationTime (4010,1041) VR=DT VM=1 OOI Owner Creation Time
	OOIOwnerCreationTime = New(0x4010, 0x1041)

	// OOIType (4010,1042) VR=CS VM=1 OOI Type
	OOIType = New(0x4010, 0x1042)

	// OOISize (4010,1043) VR=FL VM=3 OOI Size
	OOISize = New(0x4010, 0x1043)

	// AcquisitionStatus (4010,1044) VR=CS VM=1 Acquisition Status
	AcquisitionStatus = New(0x4010, 0x1044)

	// BasisMaterialsCodeSequence (4010,1045) VR=SQ VM=1 Basis Materials Code Sequence
	BasisMaterialsCodeSequence = New(0x4010, 0x1045)

	// PhantomType (4010,1046) VR=CS VM=1 Phantom Type
	PhantomType = New(0x4010, 0x1046)

	// OOIOwnerSequence (4010,1047) VR=SQ VM=1 OOI Owner Sequence
	OOIOwnerSequence = New(0x4010, 0x1047)

	// ScanType (4010,1048) VR=CS VM=1 Scan Type
	ScanType = New(0x4010, 0x1048)

	// ItineraryID (4010,1051) VR=LO VM=1 Itinerary ID
	ItineraryID = New(0x4010, 0x1051)

	// ItineraryIDType (4010,1052) VR=SH VM=1 Itinerary ID Type
	ItineraryIDType = New(0x4010, 0x1052)

	// ItineraryIDAssigningAuthority (4010,1053) VR=LO VM=1 Itinerary ID Assigning Authority
	ItineraryIDAssigningAuthority = New(0x4010, 0x1053)

	// RouteID (4010,1054) VR=SH VM=1 Route ID
	RouteID = New(0x4010, 0x1054)

	// RouteIDAssigningAuthority (4010,1055) VR=SH VM=1 Route ID Assigning Authority
	RouteIDAssigningAuthority = New(0x4010, 0x1055)

	// InboundArrivalType (4010,1056) VR=CS VM=1 Inbound Arrival Type
	InboundArrivalType = New(0x4010, 0x1056)

	// CarrierID (4010,1058) VR=SH VM=1 Carrier ID
	CarrierID = New(0x4010, 0x1058)

	// CarrierIDAssigningAuthority (4010,1059) VR=CS VM=1 Carrier ID Assigning Authority
	CarrierIDAssigningAuthority = New(0x4010, 0x1059)

	// SourceOrientation (4010,1060) VR=FL VM=3 Source Orientation
	SourceOrientation = New(0x4010, 0x1060)

	// SourcePosition (4010,1061) VR=FL VM=3 Source Position
	SourcePosition = New(0x4010, 0x1061)

	// BeltHeight (4010,1062) VR=FL VM=1 Belt Height
	BeltHeight = New(0x4010, 0x1062)

	// AlgorithmRoutingCodeSequence (4010,1064) VR=SQ VM=1 Algorithm Routing Code Sequence
	AlgorithmRoutingCodeSequence = New(0x4010, 0x1064)

	// TransportClassification (4010,1067) VR=CS VM=1 Transport Classification
	TransportClassification = New(0x4010, 0x1067)

	// OOITypeDescriptor (4010,1068) VR=LT VM=1 OOI Type Descriptor
	OOITypeDescriptor = New(0x4010, 0x1068)

	// TotalProcessingTime (4010,1069) VR=FL VM=1 Total Processing Time
	TotalProcessingTime = New(0x4010, 0x1069)

	// DetectorCalibrationData (4010,106C) VR=OB VM=1 Detector Calibration Data
	DetectorCalibrationData = New(0x4010, 0x106C)

	// AdditionalScreeningPerformed (4010,106D) VR=CS VM=1 Additional Screening Performed
	AdditionalScreeningPerformed = New(0x4010, 0x106D)

	// AdditionalInspectionSelectionCriteria (4010,106E) VR=CS VM=1 Additional Inspection Selection Criteria
	AdditionalInspectionSelectionCriteria = New(0x4010, 0x106E)

	// AdditionalInspectionMethodSequence (4010,106F) VR=SQ VM=1 Additional Inspection Method Sequence
	AdditionalInspectionMethodSequence = New(0x4010, 0x106F)

	// AITDeviceType (4010,1070) VR=CS VM=1 AIT Device Type
	AITDeviceType = New(0x4010, 0x1070)

	// QRMeasurementsSequence (4010,1071) VR=SQ VM=1 QR Measurements Sequence
	QRMeasurementsSequence = New(0x4010, 0x1071)

	// TargetMaterialSequence (4010,1072) VR=SQ VM=1 Target Material Sequence
	TargetMaterialSequence = New(0x4010, 0x1072)

	// SNRThreshold (4010,1073) VR=FD VM=1 SNR Threshold
	SNRThreshold = New(0x4010, 0x1073)

	// ImageScaleRepresentation (4010,1075) VR=DS VM=1 Image Scale Representation
	ImageScaleRepresentation = New(0x4010, 0x1075)

	// ReferencedPTOSequence (4010,1076) VR=SQ VM=1 Referenced PTO Sequence
	ReferencedPTOSequence = New(0x4010, 0x1076)

	// ReferencedTDRInstanceSequence (4010,1077) VR=SQ VM=1 Referenced TDR Instance Sequence
	ReferencedTDRInstanceSequence = New(0x4010, 0x1077)

	// PTOLocationDescription (4010,1078) VR=ST VM=1 PTO Location Description
	PTOLocationDescription = New(0x4010, 0x1078)

	// AnomalyLocatorIndicatorSequence (4010,1079) VR=SQ VM=1 Anomaly Locator Indicator Sequence
	AnomalyLocatorIndicatorSequence = New(0x4010, 0x1079)

	// AnomalyLocatorIndicator (4010,107A) VR=FL VM=3 Anomaly Locator Indicator
	AnomalyLocatorIndicator = New(0x4010, 0x107A)

	// PTORegionSequence (4010,107B) VR=SQ VM=1 PTO Region Sequence
	PTORegionSequence = New(0x4010, 0x107B)

	// InspectionSelectionCriteria (4010,107C) VR=CS VM=1 Inspection Selection Criteria
	InspectionSelectionCriteria = New(0x4010, 0x107C)

	// SecondaryInspectionMethodSequence (4010,107D) VR=SQ VM=1 Secondary Inspection Method Sequence
	SecondaryInspectionMethodSequence = New(0x4010, 0x107D)

	// PRCSToRCSOrientation (4010,107E) VR=DS VM=6 PRCS to RCS Orientation
	PRCSToRCSOrientation = New(0x4010, 0x107E)

	// MACParametersSequence (4FFE,0001) VR=SQ VM=1 MAC Parameters Sequence
	MACParametersSequence = New(0x4FFE, 0x0001)

	CurveDimensionsRETIRED = New(0x5000, 0x0005)

	NumberOfPointsRETIRED = New(0x5000, 0x0010)

	TypeOfDataRETIRED = New(0x5000, 0x0020)

	CurveDescriptionRETIRED = New(0x5000, 0x0022)

	AxisUnitsRETIRED = New(0x5000, 0x0030)

	AxisLabelsRETIRED = New(0x5000, 0x0040)

	DataValueRepresentationRETIRED = New(0x5000, 0x0103)

	MinimumCoordinateValueRETIRED = New(0x5000, 0x0104)

	MaximumCoordinateValueRETIRED = New(0x5000, 0x0105)

	CurveRangeRETIRED = New(0x5000, 0x0106)

	CurveDataDescriptorRETIRED = New(0x5000, 0x0110)

	CoordinateStartValueRETIRED = New(0x5000, 0x0112)

	CoordinateStepValueRETIRED = New(0x5000, 0x0114)

	CurveActivationLayerRETIRED = New(0x5000, 0x1001)

	AudioTypeRETIRED = New(0x5000, 0x2000)

	AudioSampleFormatRETIRED = New(0x5000, 0x2002)

	NumberOfChannelsRETIRED = New(0x5000, 0x2004)

	NumberOfSamplesRETIRED = New(0x5000, 0x2006)

	SampleRateRETIRED = New(0x5000, 0x2008)

	TotalTimeRETIRED = New(0x5000, 0x200A)

	AudioSampleDataRETIRED = New(0x5000, 0x200C)

	AudioCommentsRETIRED = New(0x5000, 0x200E)

	CurveLabelRETIRED = New(0x5000, 0x2500)

	CurveReferencedOverlaySequenceRETIRED = New(0x5000, 0x2600)

	CurveReferencedOverlayGroupRETIRED = New(0x5000, 0x2610)

	CurveDataRETIRED = New(0x5000, 0x3000)

	// SharedFunctionalGroupsSequence (5200,9229) VR=SQ VM=1 Shared Functional Groups Sequence
	SharedFunctionalGroupsSequence = New(0x5200, 0x9229)

	// PerFrameFunctionalGroupsSequence (5200,9230) VR=SQ VM=1 Per-Frame Functional Groups Sequence
	PerFrameFunctionalGroupsSequence = New(0x5200, 0x9230)

	// WaveformSequence (5400,0100) VR=SQ VM=1 Waveform Sequence
	WaveformSequence = New(0x5400, 0x0100)

	ChannelMinimumValue = New(0x5400, 0x0110)

	ChannelMaximumValue = New(0x5400, 0x0112)

	// WaveformBitsAllocated (5400,1004) VR=US VM=1 Waveform Bits Allocated
	WaveformBitsAllocated = New(0x5400, 0x1004)

	// WaveformSampleInterpretation (5400,1006) VR=CS VM=1 Waveform Sample Interpretation
	WaveformSampleInterpretation = New(0x5400, 0x1006)

	WaveformPaddingValue = New(0x5400, 0x100A)

	WaveformData = New(0x5400, 0x1010)

	// FirstOrderPhaseCorrectionAngle (5600,0010) VR=OF VM=1 First Order Phase Correction Angle
	FirstOrderPhaseCorrectionAngle = New(0x5600, 0x0010)

	// SpectroscopyData (5600,0020) VR=OF VM=1 Spectroscopy Data
	SpectroscopyData = New(0x5600, 0x0020)

	OverlayRows = New(0x6000, 0x0010)

	OverlayColumns = New(0x6000, 0x0011)

	OverlayPlanesRETIRED = New(0x6000, 0x0012)

	NumberOfFramesInOverlay = New(0x6000, 0x0015)

	OverlayDescription = New(0x6000, 0x0022)

	OverlayType = New(0x6000, 0x0040)

	OverlaySubtype = New(0x6000, 0x0045)

	OverlayOrigin = New(0x6000, 0x0050)

	ImageFrameOrigin = New(0x6000, 0x0051)

	OverlayPlaneOriginRETIRED = New(0x6000, 0x0052)

	OverlayCompressionCodeRETIRED = New(0x6000, 0x0060)

	OverlayCompressionOriginatorRETIRED = New(0x6000, 0x0061)

	OverlayCompressionLabelRETIRED = New(0x6000, 0x0062)

	OverlayCompressionDescriptionRETIRED = New(0x6000, 0x0063)

	OverlayCompressionStepPointersRETIRED = New(0x6000, 0x0066)

	OverlayRepeatIntervalRETIRED = New(0x6000, 0x0068)

	OverlayBitsGroupedRETIRED = New(0x6000, 0x0069)

	OverlayBitsAllocated = New(0x6000, 0x0100)

	OverlayBitPosition = New(0x6000, 0x0102)

	OverlayFormatRETIRED = New(0x6000, 0x0110)

	OverlayLocationRETIRED = New(0x6000, 0x0200)

	OverlayCodeLabelRETIRED = New(0x6000, 0x0800)

	OverlayNumberOfTablesRETIRED = New(0x6000, 0x0802)

	OverlayCodeTableLocationRETIRED = New(0x6000, 0x0803)

	OverlayBitsForCodeWordRETIRED = New(0x6000, 0x0804)

	OverlayActivationLayer = New(0x6000, 0x1001)

	OverlayDescriptorGrayRETIRED = New(0x6000, 0x1100)

	OverlayDescriptorRedRETIRED = New(0x6000, 0x1101)

	OverlayDescriptorGreenRETIRED = New(0x6000, 0x1102)

	OverlayDescriptorBlueRETIRED = New(0x6000, 0x1103)

	OverlaysGrayRETIRED = New(0x6000, 0x1200)

	OverlaysRedRETIRED = New(0x6000, 0x1201)

	OverlaysGreenRETIRED = New(0x6000, 0x1202)

	OverlaysBlueRETIRED = New(0x6000, 0x1203)

	ROIArea = New(0x6000, 0x1301)

	ROIMean = New(0x6000, 0x1302)

	ROIStandardDeviation = New(0x6000, 0x1303)

	OverlayLabel = New(0x6000, 0x1500)

	OverlayData = New(0x6000, 0x3000)

	OverlayCommentsRETIRED = New(0x6000, 0x4000)

	// ExtendedOffsetTable (7FE0,0001) VR=OV VM=1 Extended Offset Table
	ExtendedOffsetTable = New(0x7FE0, 0x0001)

	// ExtendedOffsetTableLengths (7FE0,0002) VR=OV VM=1 Extended Offset Table Lengths
	ExtendedOffsetTableLengths = New(0x7FE0, 0x0002)

	// EncapsulatedPixelDataValueTotalLength (7FE0,0003) VR=UV VM=1 Encapsulated Pixel Data Value Total Length
	EncapsulatedPixelDataValueTotalLength = New(0x7FE0, 0x0003)

	// FloatPixelData (7FE0,0008) VR=OF VM=1 Float Pixel Data
	FloatPixelData = New(0x7FE0, 0x0008)

	// DoubleFloatPixelData (7FE0,0009) VR=OD VM=1 Double Float Pixel Data
	DoubleFloatPixelData = New(0x7FE0, 0x0009)

	PixelData = New(0x7FE0, 0x0010)

	// CoefficientsSDVNRETIRED (7FE0,0020) VR=OW VM=1 Coefficients SDVN (RETIRED)
	CoefficientsSDVNRETIRED = New(0x7FE0, 0x0020)

	// CoefficientsSDHNRETIRED (7FE0,0030) VR=OW VM=1 Coefficients SDHN (RETIRED)
	CoefficientsSDHNRETIRED = New(0x7FE0, 0x0030)

	// CoefficientsSDDNRETIRED (7FE0,0040) VR=OW VM=1 Coefficients SDDN (RETIRED)
	CoefficientsSDDNRETIRED = New(0x7FE0, 0x0040)

	VariablePixelDataRETIRED = New(0x7F00, 0x0010)

	VariableNextDataGroupRETIRED = New(0x7F00, 0x0011)

	VariableCoefficientsSDVNRETIRED = New(0x7F00, 0x0020)

	VariableCoefficientsSDHNRETIRED = New(0x7F00, 0x0030)

	VariableCoefficientsSDDNRETIRED = New(0x7F00, 0x0040)

	// DigitalSignaturesSequence (FFFA,FFFA) VR=SQ VM=1 Digital Signatures Sequence
	DigitalSignaturesSequence = New(0xFFFA, 0xFFFA)

	// DataSetTrailingPadding (FFFC,FFFC) VR=OB VM=1 Data Set Trailing Padding
	DataSetTrailingPadding = New(0xFFFC, 0xFFFC)

	// Item (FFFE,E000) VR=NONE VM=1 Item
	Item = New(0xFFFE, 0xE000)

	// ItemDelimitationItem (FFFE,E00D) VR=NONE VM=1 Item Delimitation Item
	ItemDelimitationItem = New(0xFFFE, 0xE00D)

	// SequenceDelimitationItem (FFFE,E0DD) VR=NONE VM=1 Sequence Delimitation Item
	SequenceDelimitationItem = New(0xFFFE, 0xE0DD)

	// CurrentFrameFunctionalGroupsSequence (0006,0001) VR=SQ VM=1 Current Frame Functional Groups Sequence
	CurrentFrameFunctionalGroupsSequence = New(0x0006, 0x0001)
)
