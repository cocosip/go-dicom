// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DicomUIDGenerated.cs. DO NOT EDIT.

package uid

// Standard DICOM UID constants
var (
	// Verification Verification SOP Class
	Verification = New("1.2.840.10008.1.1", "Verification SOP Class", TypeSOPClass, false)

	// ImplicitVRLittleEndian Implicit VR Little Endian: Default Transfer Syntax for DICOM
	ImplicitVRLittleEndian = New("1.2.840.10008.1.2", "Implicit VR Little Endian: Default Transfer Syntax for DICOM", TypeTransferSyntax, false)

	// ExplicitVRLittleEndian Explicit VR Little Endian
	ExplicitVRLittleEndian = New("1.2.840.10008.1.2.1", "Explicit VR Little Endian", TypeTransferSyntax, false)

	// EncapsulatedUncompressedExplicitVRLittleEndian Encapsulated Uncompressed Explicit VR Little Endian
	EncapsulatedUncompressedExplicitVRLittleEndian = New("1.2.840.10008.1.2.1.98", "Encapsulated Uncompressed Explicit VR Little Endian", TypeTransferSyntax, false)

	// DeflatedExplicitVRLittleEndian Deflated Explicit VR Little Endian
	DeflatedExplicitVRLittleEndian = New("1.2.840.10008.1.2.1.99", "Deflated Explicit VR Little Endian", TypeTransferSyntax, false)

	// ExplicitVRBigEndianRETIRED Explicit VR Big Endian (Retired)
	ExplicitVRBigEndianRETIRED = New("1.2.840.10008.1.2.2", "Explicit VR Big Endian (Retired)", TypeTransferSyntax, true)

	// JPEGBaseline8Bit JPEG Baseline (Process 1): Default Transfer Syntax for Lossy JPEG 8 Bit Image Compression
	JPEGBaseline8Bit = New("1.2.840.10008.1.2.4.50", "JPEG Baseline (Process 1): Default Transfer Syntax for Lossy JPEG 8 Bit Image Compression", TypeTransferSyntax, false)

	// JPEGExtended12Bit JPEG Extended (Process 2 & 4): Default Transfer Syntax for Lossy JPEG 12 Bit Image Compression (Process 4 only)
	JPEGExtended12Bit = New("1.2.840.10008.1.2.4.51", "JPEG Extended (Process 2 & 4): Default Transfer Syntax for Lossy JPEG 12 Bit Image Compression (Process 4 only)", TypeTransferSyntax, false)

	// JPEGExtended35RETIRED JPEG Extended (Process 3 & 5) (Retired)
	JPEGExtended35RETIRED = New("1.2.840.10008.1.2.4.52", "JPEG Extended (Process 3 & 5) (Retired)", TypeTransferSyntax, true)

	// JPEGSpectralSelectionNonHierarchical68RETIRED JPEG Spectral Selection, Non-Hierarchical (Process 6 & 8) (Retired)
	JPEGSpectralSelectionNonHierarchical68RETIRED = New("1.2.840.10008.1.2.4.53", "JPEG Spectral Selection, Non-Hierarchical (Process 6 & 8) (Retired)", TypeTransferSyntax, true)

	// JPEGSpectralSelectionNonHierarchical79RETIRED JPEG Spectral Selection, Non-Hierarchical (Process 7 & 9) (Retired)
	JPEGSpectralSelectionNonHierarchical79RETIRED = New("1.2.840.10008.1.2.4.54", "JPEG Spectral Selection, Non-Hierarchical (Process 7 & 9) (Retired)", TypeTransferSyntax, true)

	// JPEGFullProgressionNonHierarchical1012RETIRED JPEG Full Progression, Non-Hierarchical (Process 10 & 12) (Retired)
	JPEGFullProgressionNonHierarchical1012RETIRED = New("1.2.840.10008.1.2.4.55", "JPEG Full Progression, Non-Hierarchical (Process 10 & 12) (Retired)", TypeTransferSyntax, true)

	// JPEGFullProgressionNonHierarchical1113RETIRED JPEG Full Progression, Non-Hierarchical (Process 11 & 13) (Retired)
	JPEGFullProgressionNonHierarchical1113RETIRED = New("1.2.840.10008.1.2.4.56", "JPEG Full Progression, Non-Hierarchical (Process 11 & 13) (Retired)", TypeTransferSyntax, true)

	// JPEGLossless JPEG Lossless, Non-Hierarchical (Process 14)
	JPEGLossless = New("1.2.840.10008.1.2.4.57", "JPEG Lossless, Non-Hierarchical (Process 14)", TypeTransferSyntax, false)

	// JPEGLosslessNonHierarchical15RETIRED JPEG Lossless, Non-Hierarchical (Process 15) (Retired)
	JPEGLosslessNonHierarchical15RETIRED = New("1.2.840.10008.1.2.4.58", "JPEG Lossless, Non-Hierarchical (Process 15) (Retired)", TypeTransferSyntax, true)

	// JPEGExtendedHierarchical1618RETIRED JPEG Extended, Hierarchical (Process 16 & 18) (Retired)
	JPEGExtendedHierarchical1618RETIRED = New("1.2.840.10008.1.2.4.59", "JPEG Extended, Hierarchical (Process 16 & 18) (Retired)", TypeTransferSyntax, true)

	// JPEGExtendedHierarchical1719RETIRED JPEG Extended, Hierarchical (Process 17 & 19) (Retired)
	JPEGExtendedHierarchical1719RETIRED = New("1.2.840.10008.1.2.4.60", "JPEG Extended, Hierarchical (Process 17 & 19) (Retired)", TypeTransferSyntax, true)

	// JPEGSpectralSelectionHierarchical2022RETIRED JPEG Spectral Selection, Hierarchical (Process 20 & 22) (Retired)
	JPEGSpectralSelectionHierarchical2022RETIRED = New("1.2.840.10008.1.2.4.61", "JPEG Spectral Selection, Hierarchical (Process 20 & 22) (Retired)", TypeTransferSyntax, true)

	// JPEGSpectralSelectionHierarchical2123RETIRED JPEG Spectral Selection, Hierarchical (Process 21 & 23) (Retired)
	JPEGSpectralSelectionHierarchical2123RETIRED = New("1.2.840.10008.1.2.4.62", "JPEG Spectral Selection, Hierarchical (Process 21 & 23) (Retired)", TypeTransferSyntax, true)

	// JPEGFullProgressionHierarchical2426RETIRED JPEG Full Progression, Hierarchical (Process 24 & 26) (Retired)
	JPEGFullProgressionHierarchical2426RETIRED = New("1.2.840.10008.1.2.4.63", "JPEG Full Progression, Hierarchical (Process 24 & 26) (Retired)", TypeTransferSyntax, true)

	// JPEGFullProgressionHierarchical2527RETIRED JPEG Full Progression, Hierarchical (Process 25 & 27) (Retired)
	JPEGFullProgressionHierarchical2527RETIRED = New("1.2.840.10008.1.2.4.64", "JPEG Full Progression, Hierarchical (Process 25 & 27) (Retired)", TypeTransferSyntax, true)

	// JPEGLosslessHierarchical28RETIRED JPEG Lossless, Hierarchical (Process 28) (Retired)
	JPEGLosslessHierarchical28RETIRED = New("1.2.840.10008.1.2.4.65", "JPEG Lossless, Hierarchical (Process 28) (Retired)", TypeTransferSyntax, true)

	// JPEGLosslessHierarchical29RETIRED JPEG Lossless, Hierarchical (Process 29) (Retired)
	JPEGLosslessHierarchical29RETIRED = New("1.2.840.10008.1.2.4.66", "JPEG Lossless, Hierarchical (Process 29) (Retired)", TypeTransferSyntax, true)

	// JPEGLosslessSV1 JPEG Lossless, Non-Hierarchical, First-Order Prediction (Process 14 [Selection Value 1]): Default Transfer Syntax for Lossless JPEG Image Compression
	JPEGLosslessSV1 = New("1.2.840.10008.1.2.4.70", "JPEG Lossless, Non-Hierarchical, First-Order Prediction (Process 14 [Selection Value 1]): Default Transfer Syntax for Lossless JPEG Image Compression", TypeTransferSyntax, false)

	// JPEGLSLossless JPEG-LS Lossless Image Compression
	JPEGLSLossless = New("1.2.840.10008.1.2.4.80", "JPEG-LS Lossless Image Compression", TypeTransferSyntax, false)

	// JPEGLSNearLossless JPEG-LS Lossy (Near-Lossless) Image Compression
	JPEGLSNearLossless = New("1.2.840.10008.1.2.4.81", "JPEG-LS Lossy (Near-Lossless) Image Compression", TypeTransferSyntax, false)

	// JPEG2000Lossless JPEG 2000 Image Compression (Lossless Only)
	JPEG2000Lossless = New("1.2.840.10008.1.2.4.90", "JPEG 2000 Image Compression (Lossless Only)", TypeTransferSyntax, false)

	// JPEG2000 JPEG 2000 Image Compression
	JPEG2000 = New("1.2.840.10008.1.2.4.91", "JPEG 2000 Image Compression", TypeTransferSyntax, false)

	// JPEG2000MCLossless JPEG 2000 Part 2 Multi-component Image Compression (Lossless Only)
	JPEG2000MCLossless = New("1.2.840.10008.1.2.4.92", "JPEG 2000 Part 2 Multi-component Image Compression (Lossless Only)", TypeTransferSyntax, false)

	// JPEG2000MC JPEG 2000 Part 2 Multi-component Image Compression
	JPEG2000MC = New("1.2.840.10008.1.2.4.93", "JPEG 2000 Part 2 Multi-component Image Compression", TypeTransferSyntax, false)

	// JPIPReferenced JPIP Referenced
	JPIPReferenced = New("1.2.840.10008.1.2.4.94", "JPIP Referenced", TypeTransferSyntax, false)

	// JPIPReferencedDeflate JPIP Referenced Deflate
	JPIPReferencedDeflate = New("1.2.840.10008.1.2.4.95", "JPIP Referenced Deflate", TypeTransferSyntax, false)

	// MPEG2MPML MPEG2 Main Profile / Main Level
	MPEG2MPML = New("1.2.840.10008.1.2.4.100", "MPEG2 Main Profile / Main Level", TypeTransferSyntax, false)

	// MPEG2MPMLF Fragmentable MPEG2 Main Profile / Main Level
	MPEG2MPMLF = New("1.2.840.10008.1.2.4.100.1", "Fragmentable MPEG2 Main Profile / Main Level", TypeTransferSyntax, false)

	// MPEG2MPHL MPEG2 Main Profile / High Level
	MPEG2MPHL = New("1.2.840.10008.1.2.4.101", "MPEG2 Main Profile / High Level", TypeTransferSyntax, false)

	// MPEG2MPHLF Fragmentable MPEG2 Main Profile / High Level
	MPEG2MPHLF = New("1.2.840.10008.1.2.4.101.1", "Fragmentable MPEG2 Main Profile / High Level", TypeTransferSyntax, false)

	// MPEG4HP41 MPEG-4 AVC/H.264 High Profile / Level 4.1
	MPEG4HP41 = New("1.2.840.10008.1.2.4.102", "MPEG-4 AVC/H.264 High Profile / Level 4.1", TypeTransferSyntax, false)

	// MPEG4HP41F Fragmentable MPEG-4 AVC/H.264 High Profile / Level 4.1
	MPEG4HP41F = New("1.2.840.10008.1.2.4.102.1", "Fragmentable MPEG-4 AVC/H.264 High Profile / Level 4.1", TypeTransferSyntax, false)

	// MPEG4HP41BD MPEG-4 AVC/H.264 BD-compatible High Profile / Level 4.1
	MPEG4HP41BD = New("1.2.840.10008.1.2.4.103", "MPEG-4 AVC/H.264 BD-compatible High Profile / Level 4.1", TypeTransferSyntax, false)

	// MPEG4HP41BDF Fragmentable MPEG-4 AVC/H.264 BD-compatible High Profile / Level 4.1
	MPEG4HP41BDF = New("1.2.840.10008.1.2.4.103.1", "Fragmentable MPEG-4 AVC/H.264 BD-compatible High Profile / Level 4.1", TypeTransferSyntax, false)

	// MPEG4HP422D MPEG-4 AVC/H.264 High Profile / Level 4.2 For 2D Video
	MPEG4HP422D = New("1.2.840.10008.1.2.4.104", "MPEG-4 AVC/H.264 High Profile / Level 4.2 For 2D Video", TypeTransferSyntax, false)

	// MPEG4HP422DF Fragmentable MPEG-4 AVC/H.264 High Profile / Level 4.2 For 2D Video
	MPEG4HP422DF = New("1.2.840.10008.1.2.4.104.1", "Fragmentable MPEG-4 AVC/H.264 High Profile / Level 4.2 For 2D Video", TypeTransferSyntax, false)

	// MPEG4HP423D MPEG-4 AVC/H.264 High Profile / Level 4.2 For 3D Video
	MPEG4HP423D = New("1.2.840.10008.1.2.4.105", "MPEG-4 AVC/H.264 High Profile / Level 4.2 For 3D Video", TypeTransferSyntax, false)

	// MPEG4HP423DF Fragmentable MPEG-4 AVC/H.264 High Profile / Level 4.2 For 3D Video
	MPEG4HP423DF = New("1.2.840.10008.1.2.4.105.1", "Fragmentable MPEG-4 AVC/H.264 High Profile / Level 4.2 For 3D Video", TypeTransferSyntax, false)

	// MPEG4HP42STEREO MPEG-4 AVC/H.264 Stereo High Profile / Level 4.2
	MPEG4HP42STEREO = New("1.2.840.10008.1.2.4.106", "MPEG-4 AVC/H.264 Stereo High Profile / Level 4.2", TypeTransferSyntax, false)

	// MPEG4HP42STEREOF Fragmentable MPEG-4 AVC/H.264 Stereo High Profile / Level 4.2
	MPEG4HP42STEREOF = New("1.2.840.10008.1.2.4.106.1", "Fragmentable MPEG-4 AVC/H.264 Stereo High Profile / Level 4.2", TypeTransferSyntax, false)

	// HEVCMP51 HEVC/H.265 Main Profile / Level 5.1
	HEVCMP51 = New("1.2.840.10008.1.2.4.107", "HEVC/H.265 Main Profile / Level 5.1", TypeTransferSyntax, false)

	// HEVCM10P51 HEVC/H.265 Main 10 Profile / Level 5.1
	HEVCM10P51 = New("1.2.840.10008.1.2.4.108", "HEVC/H.265 Main 10 Profile / Level 5.1", TypeTransferSyntax, false)

	// JPEGXLLossless JPEG XL Lossless
	JPEGXLLossless = New("1.2.840.10008.1.2.4.110", "JPEG XL Lossless", TypeTransferSyntax, false)

	// JPEGXLJPEGRecompression JPEG XL JPEG Recompression
	JPEGXLJPEGRecompression = New("1.2.840.10008.1.2.4.111", "JPEG XL JPEG Recompression", TypeTransferSyntax, false)

	// JPEGXL JPEG XL
	JPEGXL = New("1.2.840.10008.1.2.4.112", "JPEG XL", TypeTransferSyntax, false)

	// HTJ2KLossless High-Throughput JPEG 2000 Image Compression (Lossless Only)
	HTJ2KLossless = New("1.2.840.10008.1.2.4.201", "High-Throughput JPEG 2000 Image Compression (Lossless Only)", TypeTransferSyntax, false)

	// HTJ2KLosslessRPCL High-Throughput JPEG 2000 with RPCL Options Image Compression (Lossless Only)
	HTJ2KLosslessRPCL = New("1.2.840.10008.1.2.4.202", "High-Throughput JPEG 2000 with RPCL Options Image Compression (Lossless Only)", TypeTransferSyntax, false)

	// HTJ2K High-Throughput JPEG 2000 Image Compression
	HTJ2K = New("1.2.840.10008.1.2.4.203", "High-Throughput JPEG 2000 Image Compression", TypeTransferSyntax, false)

	// JPIPHTJ2KReferenced JPIP HTJ2K Referenced
	JPIPHTJ2KReferenced = New("1.2.840.10008.1.2.4.204", "JPIP HTJ2K Referenced", TypeTransferSyntax, false)

	// JPIPHTJ2KReferencedDeflate JPIP HTJ2K Referenced Deflate
	JPIPHTJ2KReferencedDeflate = New("1.2.840.10008.1.2.4.205", "JPIP HTJ2K Referenced Deflate", TypeTransferSyntax, false)

	// RLELossless RLE Lossless
	RLELossless = New("1.2.840.10008.1.2.5", "RLE Lossless", TypeTransferSyntax, false)

	// RFC2557MIMEEncapsulationRETIRED RFC 2557 MIME encapsulation (Retired)
	RFC2557MIMEEncapsulationRETIRED = New("1.2.840.10008.1.2.6.1", "RFC 2557 MIME encapsulation (Retired)", TypeTransferSyntax, true)

	// XMLEncodingRETIRED XML Encoding (Retired)
	XMLEncodingRETIRED = New("1.2.840.10008.1.2.6.2", "XML Encoding (Retired)", TypeTransferSyntax, true)

	// SMPTEST211020UncompressedProgressiveActiveVideo SMPTE ST 2110-20 Uncompressed Progressive Active Video
	SMPTEST211020UncompressedProgressiveActiveVideo = New("1.2.840.10008.1.2.7.1", "SMPTE ST 2110-20 Uncompressed Progressive Active Video", TypeTransferSyntax, false)

	// SMPTEST211020UncompressedInterlacedActiveVideo SMPTE ST 2110-20 Uncompressed Interlaced Active Video
	SMPTEST211020UncompressedInterlacedActiveVideo = New("1.2.840.10008.1.2.7.2", "SMPTE ST 2110-20 Uncompressed Interlaced Active Video", TypeTransferSyntax, false)

	// SMPTEST211030PCMDigitalAudio SMPTE ST 2110-30 PCM Digital Audio
	SMPTEST211030PCMDigitalAudio = New("1.2.840.10008.1.2.7.3", "SMPTE ST 2110-30 PCM Digital Audio", TypeTransferSyntax, false)

	// DeflatedImageFrameCompression Deflated Image Frame Compression
	DeflatedImageFrameCompression = New("1.2.840.10008.1.2.8.1", "Deflated Image Frame Compression", TypeTransferSyntax, false)

	// MediaStorageDirectoryStorage Media Storage Directory Storage
	MediaStorageDirectoryStorage = New("1.2.840.10008.1.3.10", "Media Storage Directory Storage", TypeSOPClass, false)

	// HotIronPalette Hot Iron Color Palette SOP Instance
	HotIronPalette = New("1.2.840.10008.1.5.1", "Hot Iron Color Palette SOP Instance", TypeSOPInstance, false)

	// PETPalette PET Color Palette SOP Instance
	PETPalette = New("1.2.840.10008.1.5.2", "PET Color Palette SOP Instance", TypeSOPInstance, false)

	// HotMetalBluePalette Hot Metal Blue Color Palette SOP Instance
	HotMetalBluePalette = New("1.2.840.10008.1.5.3", "Hot Metal Blue Color Palette SOP Instance", TypeSOPInstance, false)

	// PET20StepPalette PET 20 Step Color Palette SOP Instance
	PET20StepPalette = New("1.2.840.10008.1.5.4", "PET 20 Step Color Palette SOP Instance", TypeSOPInstance, false)

	// SpringPalette Spring Color Palette SOP Instance
	SpringPalette = New("1.2.840.10008.1.5.5", "Spring Color Palette SOP Instance", TypeSOPInstance, false)

	// SummerPalette Summer Color Palette SOP Instance
	SummerPalette = New("1.2.840.10008.1.5.6", "Summer Color Palette SOP Instance", TypeSOPInstance, false)

	// FallPalette Fall Color Palette SOP Instance
	FallPalette = New("1.2.840.10008.1.5.7", "Fall Color Palette SOP Instance", TypeSOPInstance, false)

	// WinterPalette Winter Color Palette SOP Instance
	WinterPalette = New("1.2.840.10008.1.5.8", "Winter Color Palette SOP Instance", TypeSOPInstance, false)

	// BasicStudyContentNotificationRETIRED Basic Study Content Notification SOP Class (Retired)
	BasicStudyContentNotificationRETIRED = New("1.2.840.10008.1.9", "Basic Study Content Notification SOP Class (Retired)", TypeSOPClass, true)

	// Papyrus3ImplicitVRLittleEndianRETIRED Papyrus 3 Implicit VR Little Endian (Retired)
	Papyrus3ImplicitVRLittleEndianRETIRED = New("1.2.840.10008.1.20", "Papyrus 3 Implicit VR Little Endian (Retired)", TypeTransferSyntax, true)

	// StorageCommitmentPushModel Storage Commitment Push Model SOP Class
	StorageCommitmentPushModel = New("1.2.840.10008.1.20.1", "Storage Commitment Push Model SOP Class", TypeSOPClass, false)

	// StorageCommitmentPushModelInstance Storage Commitment Push Model SOP Instance
	StorageCommitmentPushModelInstance = New("1.2.840.10008.1.20.1.1", "Storage Commitment Push Model SOP Instance", TypeSOPInstance, false)

	// StorageCommitmentPullModelRETIRED Storage Commitment Pull Model SOP Class (Retired)
	StorageCommitmentPullModelRETIRED = New("1.2.840.10008.1.20.2", "Storage Commitment Pull Model SOP Class (Retired)", TypeSOPClass, true)

	// StorageCommitmentPullModelInstanceRETIRED Storage Commitment Pull Model SOP Instance (Retired)
	StorageCommitmentPullModelInstanceRETIRED = New("1.2.840.10008.1.20.2.1", "Storage Commitment Pull Model SOP Instance (Retired)", TypeSOPInstance, true)

	// ProceduralEventLogging Procedural Event Logging SOP Class
	ProceduralEventLogging = New("1.2.840.10008.1.40", "Procedural Event Logging SOP Class", TypeSOPClass, false)

	// ProceduralEventLoggingInstance Procedural Event Logging SOP Instance
	ProceduralEventLoggingInstance = New("1.2.840.10008.1.40.1", "Procedural Event Logging SOP Instance", TypeSOPInstance, false)

	// SubstanceAdministrationLogging Substance Administration Logging SOP Class
	SubstanceAdministrationLogging = New("1.2.840.10008.1.42", "Substance Administration Logging SOP Class", TypeSOPClass, false)

	// SubstanceAdministrationLoggingInstance Substance Administration Logging SOP Instance
	SubstanceAdministrationLoggingInstance = New("1.2.840.10008.1.42.1", "Substance Administration Logging SOP Instance", TypeSOPInstance, false)

	// DCMUID DICOM UID Registry
	DCMUID = New("1.2.840.10008.2.6.1", "DICOM UID Registry", TypeCodingScheme, false)

	// DCM DICOM Controlled Terminology
	DCM = New("1.2.840.10008.2.16.4", "DICOM Controlled Terminology", TypeCodingScheme, false)

	// MA Adult Mouse Anatomy Ontology
	MA = New("1.2.840.10008.2.16.5", "Adult Mouse Anatomy Ontology", TypeCodingScheme, false)

	// UBERON Uberon Ontology
	UBERON = New("1.2.840.10008.2.16.6", "Uberon Ontology", TypeCodingScheme, false)

	// ITIS_TSN Integrated Taxonomic Information System (ITIS) Taxonomic Serial Number (TSN)
	ITIS_TSN = New("1.2.840.10008.2.16.7", "Integrated Taxonomic Information System (ITIS) Taxonomic Serial Number (TSN)", TypeCodingScheme, false)

	// MGI Mouse Genome Initiative (MGI)
	MGI = New("1.2.840.10008.2.16.8", "Mouse Genome Initiative (MGI)", TypeCodingScheme, false)

	// PUBCHEM_CID PubChem Compound CID
	PUBCHEM_CID = New("1.2.840.10008.2.16.9", "PubChem Compound CID", TypeCodingScheme, false)

	// DC Dublin Core
	DC = New("1.2.840.10008.2.16.10", "Dublin Core", TypeCodingScheme, false)

	// NYUMCCG New York University Melanoma Clinical Cooperative Group
	NYUMCCG = New("1.2.840.10008.2.16.11", "New York University Melanoma Clinical Cooperative Group", TypeCodingScheme, false)

	// MAYONRISBSASRG Mayo Clinic Non-radiological Images Specific Body Structure Anatomical Surface Region Guide
	MAYONRISBSASRG = New("1.2.840.10008.2.16.12", "Mayo Clinic Non-radiological Images Specific Body Structure Anatomical Surface Region Guide", TypeCodingScheme, false)

	// IBSI Image Biomarker Standardisation Initiative
	IBSI = New("1.2.840.10008.2.16.13", "Image Biomarker Standardisation Initiative", TypeCodingScheme, false)

	// RO Radiomics Ontology
	RO = New("1.2.840.10008.2.16.14", "Radiomics Ontology", TypeCodingScheme, false)

	// RADELEMENT RadElement
	RADELEMENT = New("1.2.840.10008.2.16.15", "RadElement", TypeCodingScheme, false)

	// I11 ICD-11
	I11 = New("1.2.840.10008.2.16.16", "ICD-11", TypeCodingScheme, false)

	// UNS Unified numbering system (UNS) for metals and alloys
	UNS = New("1.2.840.10008.2.16.17", "Unified numbering system (UNS) for metals and alloys", TypeCodingScheme, false)

	// RRID Research Resource Identification
	RRID = New("1.2.840.10008.2.16.18", "Research Resource Identification", TypeCodingScheme, false)

	// DICOMApplicationContext DICOM Application Context Name
	DICOMApplicationContext = New("1.2.840.10008.3.1.1.1", "DICOM Application Context Name", TypeApplicationContextName, false)

	// DetachedPatientManagementRETIRED Detached Patient Management SOP Class (Retired)
	DetachedPatientManagementRETIRED = New("1.2.840.10008.3.1.2.1.1", "Detached Patient Management SOP Class (Retired)", TypeSOPClass, true)

	// DetachedPatientManagementMetaRETIRED Detached Patient Management Meta SOP Class (Retired)
	DetachedPatientManagementMetaRETIRED = New("1.2.840.10008.3.1.2.1.4", "Detached Patient Management Meta SOP Class (Retired)", TypeMetaSOPClass, true)

	// DetachedVisitManagementRETIRED Detached Visit Management SOP Class (Retired)
	DetachedVisitManagementRETIRED = New("1.2.840.10008.3.1.2.2.1", "Detached Visit Management SOP Class (Retired)", TypeSOPClass, true)

	// DetachedStudyManagementRETIRED Detached Study Management SOP Class (Retired)
	DetachedStudyManagementRETIRED = New("1.2.840.10008.3.1.2.3.1", "Detached Study Management SOP Class (Retired)", TypeSOPClass, true)

	// StudyComponentManagementRETIRED Study Component Management SOP Class (Retired)
	StudyComponentManagementRETIRED = New("1.2.840.10008.3.1.2.3.2", "Study Component Management SOP Class (Retired)", TypeSOPClass, true)

	// ModalityPerformedProcedureStep Modality Performed Procedure Step SOP Class
	ModalityPerformedProcedureStep = New("1.2.840.10008.3.1.2.3.3", "Modality Performed Procedure Step SOP Class", TypeSOPClass, false)

	// ModalityPerformedProcedureStepRetrieve Modality Performed Procedure Step Retrieve SOP Class
	ModalityPerformedProcedureStepRetrieve = New("1.2.840.10008.3.1.2.3.4", "Modality Performed Procedure Step Retrieve SOP Class", TypeSOPClass, false)

	// ModalityPerformedProcedureStepNotification Modality Performed Procedure Step Notification SOP Class
	ModalityPerformedProcedureStepNotification = New("1.2.840.10008.3.1.2.3.5", "Modality Performed Procedure Step Notification SOP Class", TypeSOPClass, false)

	// DetachedResultsManagementRETIRED Detached Results Management SOP Class (Retired)
	DetachedResultsManagementRETIRED = New("1.2.840.10008.3.1.2.5.1", "Detached Results Management SOP Class (Retired)", TypeSOPClass, true)

	// DetachedResultsManagementMetaRETIRED Detached Results Management Meta SOP Class (Retired)
	DetachedResultsManagementMetaRETIRED = New("1.2.840.10008.3.1.2.5.4", "Detached Results Management Meta SOP Class (Retired)", TypeMetaSOPClass, true)

	// DetachedStudyManagementMetaRETIRED Detached Study Management Meta SOP Class (Retired)
	DetachedStudyManagementMetaRETIRED = New("1.2.840.10008.3.1.2.5.5", "Detached Study Management Meta SOP Class (Retired)", TypeMetaSOPClass, true)

	// DetachedInterpretationManagementRETIRED Detached Interpretation Management SOP Class (Retired)
	DetachedInterpretationManagementRETIRED = New("1.2.840.10008.3.1.2.6.1", "Detached Interpretation Management SOP Class (Retired)", TypeSOPClass, true)

	// Storage Storage Service Class
	Storage = New("1.2.840.10008.4.2", "Storage Service Class", TypeServiceClass, false)

	// BasicFilmSession Basic Film Session SOP Class
	BasicFilmSession = New("1.2.840.10008.5.1.1.1", "Basic Film Session SOP Class", TypeSOPClass, false)

	// BasicFilmBox Basic Film Box SOP Class
	BasicFilmBox = New("1.2.840.10008.5.1.1.2", "Basic Film Box SOP Class", TypeSOPClass, false)

	// BasicGrayscaleImageBox Basic Grayscale Image Box SOP Class
	BasicGrayscaleImageBox = New("1.2.840.10008.5.1.1.4", "Basic Grayscale Image Box SOP Class", TypeSOPClass, false)

	// BasicColorImageBox Basic Color Image Box SOP Class
	BasicColorImageBox = New("1.2.840.10008.5.1.1.4.1", "Basic Color Image Box SOP Class", TypeSOPClass, false)

	// ReferencedImageBoxRETIRED Referenced Image Box SOP Class (Retired)
	ReferencedImageBoxRETIRED = New("1.2.840.10008.5.1.1.4.2", "Referenced Image Box SOP Class (Retired)", TypeSOPClass, true)

	// BasicGrayscalePrintManagementMeta Basic Grayscale Print Management Meta SOP Class
	BasicGrayscalePrintManagementMeta = New("1.2.840.10008.5.1.1.9", "Basic Grayscale Print Management Meta SOP Class", TypeMetaSOPClass, false)

	// ReferencedGrayscalePrintManagementMetaRETIRED Referenced Grayscale Print Management Meta SOP Class (Retired)
	ReferencedGrayscalePrintManagementMetaRETIRED = New("1.2.840.10008.5.1.1.9.1", "Referenced Grayscale Print Management Meta SOP Class (Retired)", TypeMetaSOPClass, true)

	// PrintJob Print Job SOP Class
	PrintJob = New("1.2.840.10008.5.1.1.14", "Print Job SOP Class", TypeSOPClass, false)

	// BasicAnnotationBox Basic Annotation Box SOP Class
	BasicAnnotationBox = New("1.2.840.10008.5.1.1.15", "Basic Annotation Box SOP Class", TypeSOPClass, false)

	// Printer Printer SOP Class
	Printer = New("1.2.840.10008.5.1.1.16", "Printer SOP Class", TypeSOPClass, false)

	// PrinterConfigurationRetrieval Printer Configuration Retrieval SOP Class
	PrinterConfigurationRetrieval = New("1.2.840.10008.5.1.1.16.376", "Printer Configuration Retrieval SOP Class", TypeSOPClass, false)

	// PrinterInstance Printer SOP Instance
	PrinterInstance = New("1.2.840.10008.5.1.1.17", "Printer SOP Instance", TypeSOPInstance, false)

	// PrinterConfigurationRetrievalInstance Printer Configuration Retrieval SOP Instance
	PrinterConfigurationRetrievalInstance = New("1.2.840.10008.5.1.1.17.376", "Printer Configuration Retrieval SOP Instance", TypeSOPInstance, false)

	// BasicColorPrintManagementMeta Basic Color Print Management Meta SOP Class
	BasicColorPrintManagementMeta = New("1.2.840.10008.5.1.1.18", "Basic Color Print Management Meta SOP Class", TypeMetaSOPClass, false)

	// ReferencedColorPrintManagementMetaRETIRED Referenced Color Print Management Meta SOP Class (Retired)
	ReferencedColorPrintManagementMetaRETIRED = New("1.2.840.10008.5.1.1.18.1", "Referenced Color Print Management Meta SOP Class (Retired)", TypeMetaSOPClass, true)

	// VOILUTBox VOI LUT Box SOP Class
	VOILUTBox = New("1.2.840.10008.5.1.1.22", "VOI LUT Box SOP Class", TypeSOPClass, false)

	// PresentationLUT Presentation LUT SOP Class
	PresentationLUT = New("1.2.840.10008.5.1.1.23", "Presentation LUT SOP Class", TypeSOPClass, false)

	// ImageOverlayBoxRETIRED Image Overlay Box SOP Class (Retired)
	ImageOverlayBoxRETIRED = New("1.2.840.10008.5.1.1.24", "Image Overlay Box SOP Class (Retired)", TypeSOPClass, true)

	// BasicPrintImageOverlayBoxRETIRED Basic Print Image Overlay Box SOP Class (Retired)
	BasicPrintImageOverlayBoxRETIRED = New("1.2.840.10008.5.1.1.24.1", "Basic Print Image Overlay Box SOP Class (Retired)", TypeSOPClass, true)

	// PrintQueueInstanceRETIRED Print Queue SOP Instance (Retired)
	PrintQueueInstanceRETIRED = New("1.2.840.10008.5.1.1.25", "Print Queue SOP Instance (Retired)", TypeSOPInstance, true)

	// PrintQueueManagementRETIRED Print Queue Management SOP Class (Retired)
	PrintQueueManagementRETIRED = New("1.2.840.10008.5.1.1.26", "Print Queue Management SOP Class (Retired)", TypeSOPClass, true)

	// StoredPrintStorageRETIRED Stored Print Storage SOP Class (Retired)
	StoredPrintStorageRETIRED = New("1.2.840.10008.5.1.1.27", "Stored Print Storage SOP Class (Retired)", TypeSOPClass, true)

	// HardcopyGrayscaleImageStorageRETIRED Hardcopy Grayscale Image Storage SOP Class (Retired)
	HardcopyGrayscaleImageStorageRETIRED = New("1.2.840.10008.5.1.1.29", "Hardcopy Grayscale Image Storage SOP Class (Retired)", TypeSOPClass, true)

	// HardcopyColorImageStorageRETIRED Hardcopy Color Image Storage SOP Class (Retired)
	HardcopyColorImageStorageRETIRED = New("1.2.840.10008.5.1.1.30", "Hardcopy Color Image Storage SOP Class (Retired)", TypeSOPClass, true)

	// PullPrintRequestRETIRED Pull Print Request SOP Class (Retired)
	PullPrintRequestRETIRED = New("1.2.840.10008.5.1.1.31", "Pull Print Request SOP Class (Retired)", TypeSOPClass, true)

	// PullStoredPrintManagementMetaRETIRED Pull Stored Print Management Meta SOP Class (Retired)
	PullStoredPrintManagementMetaRETIRED = New("1.2.840.10008.5.1.1.32", "Pull Stored Print Management Meta SOP Class (Retired)", TypeMetaSOPClass, true)

	// MediaCreationManagement Media Creation Management SOP Class UID
	MediaCreationManagement = New("1.2.840.10008.5.1.1.33", "Media Creation Management SOP Class UID", TypeSOPClass, false)

	// DisplaySystem Display System SOP Class
	DisplaySystem = New("1.2.840.10008.5.1.1.40", "Display System SOP Class", TypeSOPClass, false)

	// DisplaySystemInstance Display System SOP Instance
	DisplaySystemInstance = New("1.2.840.10008.5.1.1.40.1", "Display System SOP Instance", TypeSOPInstance, false)

	// ComputedRadiographyImageStorage Computed Radiography Image Storage
	ComputedRadiographyImageStorage = New("1.2.840.10008.5.1.4.1.1.1", "Computed Radiography Image Storage", TypeSOPClass, false)

	// DigitalXRayImageStorageForPresentation Digital X-Ray Image Storage - For Presentation
	DigitalXRayImageStorageForPresentation = New("1.2.840.10008.5.1.4.1.1.1.1", "Digital X-Ray Image Storage - For Presentation", TypeSOPClass, false)

	// DigitalXRayImageStorageForProcessing Digital X-Ray Image Storage - For Processing
	DigitalXRayImageStorageForProcessing = New("1.2.840.10008.5.1.4.1.1.1.1.1", "Digital X-Ray Image Storage - For Processing", TypeSOPClass, false)

	// DigitalMammographyXRayImageStorageForPresentation Digital Mammography X-Ray Image Storage - For Presentation
	DigitalMammographyXRayImageStorageForPresentation = New("1.2.840.10008.5.1.4.1.1.1.2", "Digital Mammography X-Ray Image Storage - For Presentation", TypeSOPClass, false)

	// DigitalMammographyXRayImageStorageForProcessing Digital Mammography X-Ray Image Storage - For Processing
	DigitalMammographyXRayImageStorageForProcessing = New("1.2.840.10008.5.1.4.1.1.1.2.1", "Digital Mammography X-Ray Image Storage - For Processing", TypeSOPClass, false)

	// DigitalIntraOralXRayImageStorageForPresentation Digital Intra-Oral X-Ray Image Storage - For Presentation
	DigitalIntraOralXRayImageStorageForPresentation = New("1.2.840.10008.5.1.4.1.1.1.3", "Digital Intra-Oral X-Ray Image Storage - For Presentation", TypeSOPClass, false)

	// DigitalIntraOralXRayImageStorageForProcessing Digital Intra-Oral X-Ray Image Storage - For Processing
	DigitalIntraOralXRayImageStorageForProcessing = New("1.2.840.10008.5.1.4.1.1.1.3.1", "Digital Intra-Oral X-Ray Image Storage - For Processing", TypeSOPClass, false)

	// CTImageStorage CT Image Storage
	CTImageStorage = New("1.2.840.10008.5.1.4.1.1.2", "CT Image Storage", TypeSOPClass, false)

	// EnhancedCTImageStorage Enhanced CT Image Storage
	EnhancedCTImageStorage = New("1.2.840.10008.5.1.4.1.1.2.1", "Enhanced CT Image Storage", TypeSOPClass, false)

	// LegacyConvertedEnhancedCTImageStorage Legacy Converted Enhanced CT Image Storage
	LegacyConvertedEnhancedCTImageStorage = New("1.2.840.10008.5.1.4.1.1.2.2", "Legacy Converted Enhanced CT Image Storage", TypeSOPClass, false)

	// UltrasoundMultiFrameImageStorageRetiredRETIRED Ultrasound Multi-frame Image Storage (Retired)
	UltrasoundMultiFrameImageStorageRetiredRETIRED = New("1.2.840.10008.5.1.4.1.1.3", "Ultrasound Multi-frame Image Storage (Retired)", TypeSOPClass, true)

	// UltrasoundMultiFrameImageStorage Ultrasound Multi-frame Image Storage
	UltrasoundMultiFrameImageStorage = New("1.2.840.10008.5.1.4.1.1.3.1", "Ultrasound Multi-frame Image Storage", TypeSOPClass, false)

	// MRImageStorage MR Image Storage
	MRImageStorage = New("1.2.840.10008.5.1.4.1.1.4", "MR Image Storage", TypeSOPClass, false)

	// EnhancedMRImageStorage Enhanced MR Image Storage
	EnhancedMRImageStorage = New("1.2.840.10008.5.1.4.1.1.4.1", "Enhanced MR Image Storage", TypeSOPClass, false)

	// MRSpectroscopyStorage MR Spectroscopy Storage
	MRSpectroscopyStorage = New("1.2.840.10008.5.1.4.1.1.4.2", "MR Spectroscopy Storage", TypeSOPClass, false)

	// EnhancedMRColorImageStorage Enhanced MR Color Image Storage
	EnhancedMRColorImageStorage = New("1.2.840.10008.5.1.4.1.1.4.3", "Enhanced MR Color Image Storage", TypeSOPClass, false)

	// LegacyConvertedEnhancedMRImageStorage Legacy Converted Enhanced MR Image Storage
	LegacyConvertedEnhancedMRImageStorage = New("1.2.840.10008.5.1.4.1.1.4.4", "Legacy Converted Enhanced MR Image Storage", TypeSOPClass, false)

	// NuclearMedicineImageStorageRetiredRETIRED Nuclear Medicine Image Storage (Retired)
	NuclearMedicineImageStorageRetiredRETIRED = New("1.2.840.10008.5.1.4.1.1.5", "Nuclear Medicine Image Storage (Retired)", TypeSOPClass, true)

	// UltrasoundImageStorageRetiredRETIRED Ultrasound Image Storage (Retired)
	UltrasoundImageStorageRetiredRETIRED = New("1.2.840.10008.5.1.4.1.1.6", "Ultrasound Image Storage (Retired)", TypeSOPClass, true)

	// UltrasoundImageStorage Ultrasound Image Storage
	UltrasoundImageStorage = New("1.2.840.10008.5.1.4.1.1.6.1", "Ultrasound Image Storage", TypeSOPClass, false)

	// EnhancedUSVolumeStorage Enhanced US Volume Storage
	EnhancedUSVolumeStorage = New("1.2.840.10008.5.1.4.1.1.6.2", "Enhanced US Volume Storage", TypeSOPClass, false)

	// PhotoacousticImageStorage Photoacoustic Image Storage
	PhotoacousticImageStorage = New("1.2.840.10008.5.1.4.1.1.6.3", "Photoacoustic Image Storage ", TypeSOPClass, false)

	// SecondaryCaptureImageStorage Secondary Capture Image Storage
	SecondaryCaptureImageStorage = New("1.2.840.10008.5.1.4.1.1.7", "Secondary Capture Image Storage", TypeSOPClass, false)

	// MultiFrameSingleBitSecondaryCaptureImageStorage Multi-frame Single Bit Secondary Capture Image Storage
	MultiFrameSingleBitSecondaryCaptureImageStorage = New("1.2.840.10008.5.1.4.1.1.7.1", "Multi-frame Single Bit Secondary Capture Image Storage", TypeSOPClass, false)

	// MultiFrameGrayscaleByteSecondaryCaptureImageStorage Multi-frame Grayscale Byte Secondary Capture Image Storage
	MultiFrameGrayscaleByteSecondaryCaptureImageStorage = New("1.2.840.10008.5.1.4.1.1.7.2", "Multi-frame Grayscale Byte Secondary Capture Image Storage", TypeSOPClass, false)

	// MultiFrameGrayscaleWordSecondaryCaptureImageStorage Multi-frame Grayscale Word Secondary Capture Image Storage
	MultiFrameGrayscaleWordSecondaryCaptureImageStorage = New("1.2.840.10008.5.1.4.1.1.7.3", "Multi-frame Grayscale Word Secondary Capture Image Storage", TypeSOPClass, false)

	// MultiFrameTrueColorSecondaryCaptureImageStorage Multi-frame True Color Secondary Capture Image Storage
	MultiFrameTrueColorSecondaryCaptureImageStorage = New("1.2.840.10008.5.1.4.1.1.7.4", "Multi-frame True Color Secondary Capture Image Storage", TypeSOPClass, false)

	// StandaloneOverlayStorageRETIRED Standalone Overlay Storage (Retired)
	StandaloneOverlayStorageRETIRED = New("1.2.840.10008.5.1.4.1.1.8", "Standalone Overlay Storage (Retired)", TypeSOPClass, true)

	// StandaloneCurveStorageRETIRED Standalone Curve Storage (Retired)
	StandaloneCurveStorageRETIRED = New("1.2.840.10008.5.1.4.1.1.9", "Standalone Curve Storage (Retired)", TypeSOPClass, true)

	// WaveformStorageTrialRETIRED Waveform Storage - Trial (Retired)
	WaveformStorageTrialRETIRED = New("1.2.840.10008.5.1.4.1.1.9.1", "Waveform Storage - Trial (Retired)", TypeSOPClass, true)

	// TwelveLeadECGWaveformStorage 12-lead ECG Waveform Storage
	TwelveLeadECGWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.1.1", "12-lead ECG Waveform Storage", TypeSOPClass, false)

	// GeneralECGWaveformStorage General ECG Waveform Storage
	GeneralECGWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.1.2", "General ECG Waveform Storage", TypeSOPClass, false)

	// AmbulatoryECGWaveformStorage Ambulatory ECG Waveform Storage
	AmbulatoryECGWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.1.3", "Ambulatory ECG Waveform Storage", TypeSOPClass, false)

	// General32bitECGWaveformStorage General 32-bit ECG Waveform Storage
	General32bitECGWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.1.4", "General 32-bit ECG Waveform Storage", TypeSOPClass, false)

	// HemodynamicWaveformStorage Hemodynamic Waveform Storage
	HemodynamicWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.2.1", "Hemodynamic Waveform Storage", TypeSOPClass, false)

	// CardiacElectrophysiologyWaveformStorage Cardiac Electrophysiology Waveform Storage
	CardiacElectrophysiologyWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.3.1", "Cardiac Electrophysiology Waveform Storage", TypeSOPClass, false)

	// BasicVoiceAudioWaveformStorage Basic Voice Audio Waveform Storage
	BasicVoiceAudioWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.4.1", "Basic Voice Audio Waveform Storage", TypeSOPClass, false)

	// GeneralAudioWaveformStorage General Audio Waveform Storage
	GeneralAudioWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.4.2", "General Audio Waveform Storage", TypeSOPClass, false)

	// ArterialPulseWaveformStorage Arterial Pulse Waveform Storage
	ArterialPulseWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.5.1", "Arterial Pulse Waveform Storage", TypeSOPClass, false)

	// RespiratoryWaveformStorage Respiratory Waveform Storage
	RespiratoryWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.6.1", "Respiratory Waveform Storage", TypeSOPClass, false)

	// MultichannelRespiratoryWaveformStorage Multi-channel Respiratory Waveform Storage
	MultichannelRespiratoryWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.6.2", "Multi-channel Respiratory Waveform Storage", TypeSOPClass, false)

	// RoutineScalpElectroencephalogramWaveformStorage Routine Scalp Electroencephalogram Waveform Storage
	RoutineScalpElectroencephalogramWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.7.1", "Routine Scalp Electroencephalogram Waveform Storage", TypeSOPClass, false)

	// ElectromyogramWaveformStorage Electromyogram Waveform Storage
	ElectromyogramWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.7.2", "Electromyogram Waveform Storage", TypeSOPClass, false)

	// ElectrooculogramWaveformStorage Electrooculogram Waveform Storage
	ElectrooculogramWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.7.3", "Electrooculogram Waveform Storage", TypeSOPClass, false)

	// SleepElectroencephalogramWaveformStorage Sleep Electroencephalogram Waveform Storage
	SleepElectroencephalogramWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.7.4", "Sleep Electroencephalogram Waveform Storage", TypeSOPClass, false)

	// BodyPositionWaveformStorage Body Position Waveform Storage
	BodyPositionWaveformStorage = New("1.2.840.10008.5.1.4.1.1.9.8.1", "Body Position Waveform Storage", TypeSOPClass, false)

	// WaveformPresentationStateStorage Waveform Presentation State Storage
	WaveformPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.9.100.1", "Waveform Presentation State Storage", TypeSOPClass, false)

	// WaveformAcquisitionPresentationStateStorage Waveform Acquisition Presentation State Storage
	WaveformAcquisitionPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.9.100.2", "Waveform Acquisition Presentation State Storage", TypeSOPClass, false)

	// StandaloneModalityLUTStorageRETIRED Standalone Modality LUT Storage (Retired)
	StandaloneModalityLUTStorageRETIRED = New("1.2.840.10008.5.1.4.1.1.10", "Standalone Modality LUT Storage (Retired)", TypeSOPClass, true)

	// StandaloneVOILUTStorageRETIRED Standalone VOI LUT Storage (Retired)
	StandaloneVOILUTStorageRETIRED = New("1.2.840.10008.5.1.4.1.1.11", "Standalone VOI LUT Storage (Retired)", TypeSOPClass, true)

	// GrayscaleSoftcopyPresentationStateStorage Grayscale Softcopy Presentation State Storage
	GrayscaleSoftcopyPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.1", "Grayscale Softcopy Presentation State Storage", TypeSOPClass, false)

	// ColorSoftcopyPresentationStateStorage Color Softcopy Presentation State Storage
	ColorSoftcopyPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.2", "Color Softcopy Presentation State Storage", TypeSOPClass, false)

	// PseudoColorSoftcopyPresentationStateStorage Pseudo-Color Softcopy Presentation State Storage
	PseudoColorSoftcopyPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.3", "Pseudo-Color Softcopy Presentation State Storage", TypeSOPClass, false)

	// BlendingSoftcopyPresentationStateStorage Blending Softcopy Presentation State Storage
	BlendingSoftcopyPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.4", "Blending Softcopy Presentation State Storage", TypeSOPClass, false)

	// XAXRFGrayscaleSoftcopyPresentationStateStorage XA/XRF Grayscale Softcopy Presentation State Storage
	XAXRFGrayscaleSoftcopyPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.5", "XA/XRF Grayscale Softcopy Presentation State Storage", TypeSOPClass, false)

	// GrayscalePlanarMPRVolumetricPresentationStateStorage Grayscale Planar MPR Volumetric Presentation State Storage
	GrayscalePlanarMPRVolumetricPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.6", "Grayscale Planar MPR Volumetric Presentation State Storage", TypeSOPClass, false)

	// CompositingPlanarMPRVolumetricPresentationStateStorage Compositing Planar MPR Volumetric Presentation State Storage
	CompositingPlanarMPRVolumetricPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.7", "Compositing Planar MPR Volumetric Presentation State Storage", TypeSOPClass, false)

	// AdvancedBlendingPresentationStateStorage Advanced Blending Presentation State Storage
	AdvancedBlendingPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.8", "Advanced Blending Presentation State Storage", TypeSOPClass, false)

	// VolumeRenderingVolumetricPresentationStateStorage Volume Rendering Volumetric Presentation State Storage
	VolumeRenderingVolumetricPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.9", "Volume Rendering Volumetric Presentation State Storage", TypeSOPClass, false)

	// SegmentedVolumeRenderingVolumetricPresentationStateStorage Segmented Volume Rendering Volumetric Presentation State Storage
	SegmentedVolumeRenderingVolumetricPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.10", "Segmented Volume Rendering Volumetric Presentation State Storage", TypeSOPClass, false)

	// MultipleVolumeRenderingVolumetricPresentationStateStorage Multiple Volume Rendering Volumetric Presentation State Storage
	MultipleVolumeRenderingVolumetricPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.11", "Multiple Volume Rendering Volumetric Presentation State Storage", TypeSOPClass, false)

	// VariableModalityLUTSoftcopyPresentationStateStorage Variable Modality LUT Softcopy Presentation State Storage
	VariableModalityLUTSoftcopyPresentationStateStorage = New("1.2.840.10008.5.1.4.1.1.11.12", "Variable Modality LUT Softcopy Presentation State Storage", TypeSOPClass, false)

	// XRayAngiographicImageStorage X-Ray Angiographic Image Storage
	XRayAngiographicImageStorage = New("1.2.840.10008.5.1.4.1.1.12.1", "X-Ray Angiographic Image Storage", TypeSOPClass, false)

	// EnhancedXAImageStorage Enhanced XA Image Storage
	EnhancedXAImageStorage = New("1.2.840.10008.5.1.4.1.1.12.1.1", "Enhanced XA Image Storage", TypeSOPClass, false)

	// XRayRadiofluoroscopicImageStorage X-Ray Radiofluoroscopic Image Storage
	XRayRadiofluoroscopicImageStorage = New("1.2.840.10008.5.1.4.1.1.12.2", "X-Ray Radiofluoroscopic Image Storage", TypeSOPClass, false)

	// EnhancedXRFImageStorage Enhanced XRF Image Storage
	EnhancedXRFImageStorage = New("1.2.840.10008.5.1.4.1.1.12.2.1", "Enhanced XRF Image Storage", TypeSOPClass, false)

	// XRayAngiographicBiPlaneImageStorageRETIRED X-Ray Angiographic Bi-Plane Image Storage (Retired)
	XRayAngiographicBiPlaneImageStorageRETIRED = New("1.2.840.10008.5.1.4.1.1.12.3", "X-Ray Angiographic Bi-Plane Image Storage (Retired)", TypeSOPClass, true)

	// XRay3DAngiographicImageStorage X-Ray 3D Angiographic Image Storage
	XRay3DAngiographicImageStorage = New("1.2.840.10008.5.1.4.1.1.13.1.1", "X-Ray 3D Angiographic Image Storage", TypeSOPClass, false)

	// XRay3DCraniofacialImageStorage X-Ray 3D Craniofacial Image Storage
	XRay3DCraniofacialImageStorage = New("1.2.840.10008.5.1.4.1.1.13.1.2", "X-Ray 3D Craniofacial Image Storage", TypeSOPClass, false)

	// BreastTomosynthesisImageStorage Breast Tomosynthesis Image Storage
	BreastTomosynthesisImageStorage = New("1.2.840.10008.5.1.4.1.1.13.1.3", "Breast Tomosynthesis Image Storage", TypeSOPClass, false)

	// BreastProjectionXRayImageStorageForPresentation Breast Projection X-Ray Image Storage - For Presentation
	BreastProjectionXRayImageStorageForPresentation = New("1.2.840.10008.5.1.4.1.1.13.1.4", "Breast Projection X-Ray Image Storage - For Presentation", TypeSOPClass, false)

	// BreastProjectionXRayImageStorageForProcessing Breast Projection X-Ray Image Storage - For Processing
	BreastProjectionXRayImageStorageForProcessing = New("1.2.840.10008.5.1.4.1.1.13.1.5", "Breast Projection X-Ray Image Storage - For Processing", TypeSOPClass, false)

	// IntravascularOpticalCoherenceTomographyImageStorageForPresentation Intravascular Optical Coherence Tomography Image Storage - For Presentation
	IntravascularOpticalCoherenceTomographyImageStorageForPresentation = New("1.2.840.10008.5.1.4.1.1.14.1", "Intravascular Optical Coherence Tomography Image Storage - For Presentation", TypeSOPClass, false)

	// IntravascularOpticalCoherenceTomographyImageStorageForProcessing Intravascular Optical Coherence Tomography Image Storage - For Processing
	IntravascularOpticalCoherenceTomographyImageStorageForProcessing = New("1.2.840.10008.5.1.4.1.1.14.2", "Intravascular Optical Coherence Tomography Image Storage - For Processing", TypeSOPClass, false)

	// NuclearMedicineImageStorage Nuclear Medicine Image Storage
	NuclearMedicineImageStorage = New("1.2.840.10008.5.1.4.1.1.20", "Nuclear Medicine Image Storage", TypeSOPClass, false)

	// ParametricMapStorage Parametric Map Storage
	ParametricMapStorage = New("1.2.840.10008.5.1.4.1.1.30", "Parametric Map Storage", TypeSOPClass, false)

	// RawDataStorage Raw Data Storage
	RawDataStorage = New("1.2.840.10008.5.1.4.1.1.66", "Raw Data Storage", TypeSOPClass, false)

	// SpatialRegistrationStorage Spatial Registration Storage
	SpatialRegistrationStorage = New("1.2.840.10008.5.1.4.1.1.66.1", "Spatial Registration Storage", TypeSOPClass, false)

	// SpatialFiducialsStorage Spatial Fiducials Storage
	SpatialFiducialsStorage = New("1.2.840.10008.5.1.4.1.1.66.2", "Spatial Fiducials Storage", TypeSOPClass, false)

	// DeformableSpatialRegistrationStorage Deformable Spatial Registration Storage
	DeformableSpatialRegistrationStorage = New("1.2.840.10008.5.1.4.1.1.66.3", "Deformable Spatial Registration Storage", TypeSOPClass, false)

	// SegmentationStorage Segmentation Storage
	SegmentationStorage = New("1.2.840.10008.5.1.4.1.1.66.4", "Segmentation Storage", TypeSOPClass, false)

	// SurfaceSegmentationStorage Surface Segmentation Storage
	SurfaceSegmentationStorage = New("1.2.840.10008.5.1.4.1.1.66.5", "Surface Segmentation Storage", TypeSOPClass, false)

	// TractographyResultsStorage Tractography Results Storage
	TractographyResultsStorage = New("1.2.840.10008.5.1.4.1.1.66.6", "Tractography Results Storage", TypeSOPClass, false)

	// LabelMapSegmentationStorage Label Map Segmentation Storage
	LabelMapSegmentationStorage = New("1.2.840.10008.5.1.4.1.1.66.7", "Label Map Segmentation Storage", TypeSOPClass, false)

	// HeightMapSegmentationStorage Height Map Segmentation Storage
	HeightMapSegmentationStorage = New("1.2.840.10008.5.1.4.1.1.66.8", "Height Map Segmentation Storage", TypeSOPClass, false)

	// RealWorldValueMappingStorage Real World Value Mapping Storage
	RealWorldValueMappingStorage = New("1.2.840.10008.5.1.4.1.1.67", "Real World Value Mapping Storage", TypeSOPClass, false)

	// SurfaceScanMeshStorage Surface Scan Mesh Storage
	SurfaceScanMeshStorage = New("1.2.840.10008.5.1.4.1.1.68.1", "Surface Scan Mesh Storage", TypeSOPClass, false)

	// SurfaceScanPointCloudStorage Surface Scan Point Cloud Storage
	SurfaceScanPointCloudStorage = New("1.2.840.10008.5.1.4.1.1.68.2", "Surface Scan Point Cloud Storage", TypeSOPClass, false)

	// VLImageStorageTrialRETIRED VL Image Storage - Trial (Retired)
	VLImageStorageTrialRETIRED = New("1.2.840.10008.5.1.4.1.1.77.1", "VL Image Storage - Trial (Retired)", TypeSOPClass, true)

	// VLMultiFrameImageStorageTrialRETIRED VL Multi-frame Image Storage - Trial (Retired)
	VLMultiFrameImageStorageTrialRETIRED = New("1.2.840.10008.5.1.4.1.1.77.2", "VL Multi-frame Image Storage - Trial (Retired)", TypeSOPClass, true)

	// VLEndoscopicImageStorage VL Endoscopic Image Storage
	VLEndoscopicImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.1", "VL Endoscopic Image Storage", TypeSOPClass, false)

	// VideoEndoscopicImageStorage Video Endoscopic Image Storage
	VideoEndoscopicImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.1.1", "Video Endoscopic Image Storage", TypeSOPClass, false)

	// VLMicroscopicImageStorage VL Microscopic Image Storage
	VLMicroscopicImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.2", "VL Microscopic Image Storage", TypeSOPClass, false)

	// VideoMicroscopicImageStorage Video Microscopic Image Storage
	VideoMicroscopicImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.2.1", "Video Microscopic Image Storage", TypeSOPClass, false)

	// VLSlideCoordinatesMicroscopicImageStorage VL Slide-Coordinates Microscopic Image Storage
	VLSlideCoordinatesMicroscopicImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.3", "VL Slide-Coordinates Microscopic Image Storage", TypeSOPClass, false)

	// VLPhotographicImageStorage VL Photographic Image Storage
	VLPhotographicImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.4", "VL Photographic Image Storage", TypeSOPClass, false)

	// VideoPhotographicImageStorage Video Photographic Image Storage
	VideoPhotographicImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.4.1", "Video Photographic Image Storage", TypeSOPClass, false)

	// OphthalmicPhotography8BitImageStorage Ophthalmic Photography 8 Bit Image Storage
	OphthalmicPhotography8BitImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.5.1", "Ophthalmic Photography 8 Bit Image Storage", TypeSOPClass, false)

	// OphthalmicPhotography16BitImageStorage Ophthalmic Photography 16 Bit Image Storage
	OphthalmicPhotography16BitImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.5.2", "Ophthalmic Photography 16 Bit Image Storage", TypeSOPClass, false)

	// StereometricRelationshipStorage Stereometric Relationship Storage
	StereometricRelationshipStorage = New("1.2.840.10008.5.1.4.1.1.77.1.5.3", "Stereometric Relationship Storage", TypeSOPClass, false)

	// OphthalmicTomographyImageStorage Ophthalmic Tomography Image Storage
	OphthalmicTomographyImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.5.4", "Ophthalmic Tomography Image Storage", TypeSOPClass, false)

	// WideFieldOphthalmicPhotographyStereographicProjectionImageStorage Wide Field Ophthalmic Photography Stereographic Projection Image Storage
	WideFieldOphthalmicPhotographyStereographicProjectionImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.5.5", "Wide Field Ophthalmic Photography Stereographic Projection Image Storage", TypeSOPClass, false)

	// WideFieldOphthalmicPhotography3DCoordinatesImageStorage Wide Field Ophthalmic Photography 3D Coordinates Image Storage
	WideFieldOphthalmicPhotography3DCoordinatesImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.5.6", "Wide Field Ophthalmic Photography 3D Coordinates Image Storage", TypeSOPClass, false)

	// OphthalmicOpticalCoherenceTomographyEnFaceImageStorage Ophthalmic Optical Coherence Tomography En Face Image Storage
	OphthalmicOpticalCoherenceTomographyEnFaceImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.5.7", "Ophthalmic Optical Coherence Tomography En Face Image Storage", TypeSOPClass, false)

	// OphthalmicOpticalCoherenceTomographyBscanVolumeAnalysisStorage Ophthalmic Optical Coherence Tomography B-scan Volume Analysis Storage
	OphthalmicOpticalCoherenceTomographyBscanVolumeAnalysisStorage = New("1.2.840.10008.5.1.4.1.1.77.1.5.8", "Ophthalmic Optical Coherence Tomography B-scan Volume Analysis Storage", TypeSOPClass, false)

	// VLWholeSlideMicroscopyImageStorage VL Whole Slide Microscopy Image Storage
	VLWholeSlideMicroscopyImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.6", "VL Whole Slide Microscopy Image Storage", TypeSOPClass, false)

	// DermoscopicPhotographyImageStorage Dermoscopic Photography Image Storage
	DermoscopicPhotographyImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.7", "Dermoscopic Photography Image Storage", TypeSOPClass, false)

	// ConfocalMicroscopyImageStorage Confocal Microscopy Image Storage
	ConfocalMicroscopyImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.8", "Confocal Microscopy Image Storage", TypeSOPClass, false)

	// ConfocalMicroscopyTiledPyramidalImageStorage Confocal Microscopy Tiled Pyramidal Image Storage
	ConfocalMicroscopyTiledPyramidalImageStorage = New("1.2.840.10008.5.1.4.1.1.77.1.9", "Confocal Microscopy Tiled Pyramidal Image Storage", TypeSOPClass, false)

	// LensometryMeasurementsStorage Lensometry Measurements Storage
	LensometryMeasurementsStorage = New("1.2.840.10008.5.1.4.1.1.78.1", "Lensometry Measurements Storage", TypeSOPClass, false)

	// AutorefractionMeasurementsStorage Autorefraction Measurements Storage
	AutorefractionMeasurementsStorage = New("1.2.840.10008.5.1.4.1.1.78.2", "Autorefraction Measurements Storage", TypeSOPClass, false)

	// KeratometryMeasurementsStorage Keratometry Measurements Storage
	KeratometryMeasurementsStorage = New("1.2.840.10008.5.1.4.1.1.78.3", "Keratometry Measurements Storage", TypeSOPClass, false)

	// SubjectiveRefractionMeasurementsStorage Subjective Refraction Measurements Storage
	SubjectiveRefractionMeasurementsStorage = New("1.2.840.10008.5.1.4.1.1.78.4", "Subjective Refraction Measurements Storage", TypeSOPClass, false)

	// VisualAcuityMeasurementsStorage Visual Acuity Measurements Storage
	VisualAcuityMeasurementsStorage = New("1.2.840.10008.5.1.4.1.1.78.5", "Visual Acuity Measurements Storage", TypeSOPClass, false)

	// SpectaclePrescriptionReportStorage Spectacle Prescription Report Storage
	SpectaclePrescriptionReportStorage = New("1.2.840.10008.5.1.4.1.1.78.6", "Spectacle Prescription Report Storage", TypeSOPClass, false)

	// OphthalmicAxialMeasurementsStorage Ophthalmic Axial Measurements Storage
	OphthalmicAxialMeasurementsStorage = New("1.2.840.10008.5.1.4.1.1.78.7", "Ophthalmic Axial Measurements Storage", TypeSOPClass, false)

	// IntraocularLensCalculationsStorage Intraocular Lens Calculations Storage
	IntraocularLensCalculationsStorage = New("1.2.840.10008.5.1.4.1.1.78.8", "Intraocular Lens Calculations Storage", TypeSOPClass, false)

	// MacularGridThicknessAndVolumeReportStorage Macular Grid Thickness and Volume Report Storage
	MacularGridThicknessAndVolumeReportStorage = New("1.2.840.10008.5.1.4.1.1.79.1", "Macular Grid Thickness and Volume Report Storage", TypeSOPClass, false)

	// OphthalmicVisualFieldStaticPerimetryMeasurementsStorage Ophthalmic Visual Field Static Perimetry Measurements Storage
	OphthalmicVisualFieldStaticPerimetryMeasurementsStorage = New("1.2.840.10008.5.1.4.1.1.80.1", "Ophthalmic Visual Field Static Perimetry Measurements Storage", TypeSOPClass, false)

	// OphthalmicThicknessMapStorage Ophthalmic Thickness Map Storage
	OphthalmicThicknessMapStorage = New("1.2.840.10008.5.1.4.1.1.81.1", "Ophthalmic Thickness Map Storage", TypeSOPClass, false)

	// CornealTopographyMapStorage Corneal Topography Map Storage
	CornealTopographyMapStorage = New("1.2.840.10008.5.1.4.1.1.82.1", "Corneal Topography Map Storage", TypeSOPClass, false)

	// TextSRStorageTrialRETIRED Text SR Storage - Trial (Retired)
	TextSRStorageTrialRETIRED = New("1.2.840.10008.5.1.4.1.1.88.1", "Text SR Storage - Trial (Retired)", TypeSOPClass, true)

	// AudioSRStorageTrialRETIRED Audio SR Storage - Trial (Retired)
	AudioSRStorageTrialRETIRED = New("1.2.840.10008.5.1.4.1.1.88.2", "Audio SR Storage - Trial (Retired)", TypeSOPClass, true)

	// DetailSRStorageTrialRETIRED Detail SR Storage - Trial (Retired)
	DetailSRStorageTrialRETIRED = New("1.2.840.10008.5.1.4.1.1.88.3", "Detail SR Storage - Trial (Retired)", TypeSOPClass, true)

	// ComprehensiveSRStorageTrialRETIRED Comprehensive SR Storage - Trial (Retired)
	ComprehensiveSRStorageTrialRETIRED = New("1.2.840.10008.5.1.4.1.1.88.4", "Comprehensive SR Storage - Trial (Retired)", TypeSOPClass, true)

	// BasicTextSRStorage Basic Text SR Storage
	BasicTextSRStorage = New("1.2.840.10008.5.1.4.1.1.88.11", "Basic Text SR Storage", TypeSOPClass, false)

	// EnhancedSRStorage Enhanced SR Storage
	EnhancedSRStorage = New("1.2.840.10008.5.1.4.1.1.88.22", "Enhanced SR Storage", TypeSOPClass, false)

	// ComprehensiveSRStorage Comprehensive SR Storage
	ComprehensiveSRStorage = New("1.2.840.10008.5.1.4.1.1.88.33", "Comprehensive SR Storage", TypeSOPClass, false)

	// Comprehensive3DSRStorage Comprehensive 3D SR Storage
	Comprehensive3DSRStorage = New("1.2.840.10008.5.1.4.1.1.88.34", "Comprehensive 3D SR Storage", TypeSOPClass, false)

	// ExtensibleSRStorage Extensible SR Storage
	ExtensibleSRStorage = New("1.2.840.10008.5.1.4.1.1.88.35", "Extensible SR Storage", TypeSOPClass, false)

	// ProcedureLogStorage Procedure Log Storage
	ProcedureLogStorage = New("1.2.840.10008.5.1.4.1.1.88.40", "Procedure Log Storage", TypeSOPClass, false)

	// MammographyCADSRStorage Mammography CAD SR Storage
	MammographyCADSRStorage = New("1.2.840.10008.5.1.4.1.1.88.50", "Mammography CAD SR Storage", TypeSOPClass, false)

	// KeyObjectSelectionDocumentStorage Key Object Selection Document Storage
	KeyObjectSelectionDocumentStorage = New("1.2.840.10008.5.1.4.1.1.88.59", "Key Object Selection Document Storage", TypeSOPClass, false)

	// ChestCADSRStorage Chest CAD SR Storage
	ChestCADSRStorage = New("1.2.840.10008.5.1.4.1.1.88.65", "Chest CAD SR Storage", TypeSOPClass, false)

	// XRayRadiationDoseSRStorage X-Ray Radiation Dose SR Storage
	XRayRadiationDoseSRStorage = New("1.2.840.10008.5.1.4.1.1.88.67", "X-Ray Radiation Dose SR Storage", TypeSOPClass, false)

	// RadiopharmaceuticalRadiationDoseSRStorage Radiopharmaceutical Radiation Dose SR Storage
	RadiopharmaceuticalRadiationDoseSRStorage = New("1.2.840.10008.5.1.4.1.1.88.68", "Radiopharmaceutical Radiation Dose SR Storage", TypeSOPClass, false)

	// ColonCADSRStorage Colon CAD SR Storage
	ColonCADSRStorage = New("1.2.840.10008.5.1.4.1.1.88.69", "Colon CAD SR Storage", TypeSOPClass, false)

	// ImplantationPlanSRStorage Implantation Plan SR Storage
	ImplantationPlanSRStorage = New("1.2.840.10008.5.1.4.1.1.88.70", "Implantation Plan SR Storage", TypeSOPClass, false)

	// AcquisitionContextSRStorage Acquisition Context SR Storage
	AcquisitionContextSRStorage = New("1.2.840.10008.5.1.4.1.1.88.71", "Acquisition Context SR Storage", TypeSOPClass, false)

	// SimplifiedAdultEchoSRStorage Simplified Adult Echo SR Storage
	SimplifiedAdultEchoSRStorage = New("1.2.840.10008.5.1.4.1.1.88.72", "Simplified Adult Echo SR Storage", TypeSOPClass, false)

	// PatientRadiationDoseSRStorage Patient Radiation Dose SR Storage
	PatientRadiationDoseSRStorage = New("1.2.840.10008.5.1.4.1.1.88.73", "Patient Radiation Dose SR Storage", TypeSOPClass, false)

	// PlannedImagingAgentAdministrationSRStorage Planned Imaging Agent Administration SR Storage
	PlannedImagingAgentAdministrationSRStorage = New("1.2.840.10008.5.1.4.1.1.88.74", "Planned Imaging Agent Administration SR Storage", TypeSOPClass, false)

	// PerformedImagingAgentAdministrationSRStorage Performed Imaging Agent Administration SR Storage
	PerformedImagingAgentAdministrationSRStorage = New("1.2.840.10008.5.1.4.1.1.88.75", "Performed Imaging Agent Administration SR Storage", TypeSOPClass, false)

	// EnhancedXRayRadiationDoseSRStorage Enhanced X-Ray Radiation Dose SR Storage
	EnhancedXRayRadiationDoseSRStorage = New("1.2.840.10008.5.1.4.1.1.88.76", "Enhanced X-Ray Radiation Dose SR Storage", TypeSOPClass, false)

	// WaveformAnnotationSRStorage Waveform Annotation SR Storage
	WaveformAnnotationSRStorage = New("1.2.840.10008.5.1.4.1.1.88.77", "Waveform Annotation SR Storage", TypeSOPClass, false)

	// ContentAssessmentResultsStorage Content Assessment Results Storage
	ContentAssessmentResultsStorage = New("1.2.840.10008.5.1.4.1.1.90.1", "Content Assessment Results Storage", TypeSOPClass, false)

	// MicroscopyBulkSimpleAnnotationsStorage Microscopy Bulk Simple Annotations Storage
	MicroscopyBulkSimpleAnnotationsStorage = New("1.2.840.10008.5.1.4.1.1.91.1", "Microscopy Bulk Simple Annotations Storage", TypeSOPClass, false)

	// EncapsulatedPDFStorage Encapsulated PDF Storage
	EncapsulatedPDFStorage = New("1.2.840.10008.5.1.4.1.1.104.1", "Encapsulated PDF Storage", TypeSOPClass, false)

	// EncapsulatedCDAStorage Encapsulated CDA Storage
	EncapsulatedCDAStorage = New("1.2.840.10008.5.1.4.1.1.104.2", "Encapsulated CDA Storage", TypeSOPClass, false)

	// EncapsulatedSTLStorage Encapsulated STL Storage
	EncapsulatedSTLStorage = New("1.2.840.10008.5.1.4.1.1.104.3", "Encapsulated STL Storage", TypeSOPClass, false)

	// EncapsulatedOBJStorage Encapsulated OBJ Storage
	EncapsulatedOBJStorage = New("1.2.840.10008.5.1.4.1.1.104.4", "Encapsulated OBJ Storage", TypeSOPClass, false)

	// EncapsulatedMTLStorage Encapsulated MTL Storage
	EncapsulatedMTLStorage = New("1.2.840.10008.5.1.4.1.1.104.5", "Encapsulated MTL Storage", TypeSOPClass, false)

	// PositronEmissionTomographyImageStorage Positron Emission Tomography Image Storage
	PositronEmissionTomographyImageStorage = New("1.2.840.10008.5.1.4.1.1.128", "Positron Emission Tomography Image Storage", TypeSOPClass, false)

	// LegacyConvertedEnhancedPETImageStorage Legacy Converted Enhanced PET Image Storage
	LegacyConvertedEnhancedPETImageStorage = New("1.2.840.10008.5.1.4.1.1.128.1", "Legacy Converted Enhanced PET Image Storage", TypeSOPClass, false)

	// StandalonePETCurveStorageRETIRED Standalone PET Curve Storage (Retired)
	StandalonePETCurveStorageRETIRED = New("1.2.840.10008.5.1.4.1.1.129", "Standalone PET Curve Storage (Retired)", TypeSOPClass, true)

	// EnhancedPETImageStorage Enhanced PET Image Storage
	EnhancedPETImageStorage = New("1.2.840.10008.5.1.4.1.1.130", "Enhanced PET Image Storage", TypeSOPClass, false)

	// BasicStructuredDisplayStorage Basic Structured Display Storage
	BasicStructuredDisplayStorage = New("1.2.840.10008.5.1.4.1.1.131", "Basic Structured Display Storage", TypeSOPClass, false)

	// CTDefinedProcedureProtocolStorage CT Defined Procedure Protocol Storage
	CTDefinedProcedureProtocolStorage = New("1.2.840.10008.5.1.4.1.1.200.1", "CT Defined Procedure Protocol Storage", TypeSOPClass, false)

	// CTPerformedProcedureProtocolStorage CT Performed Procedure Protocol Storage
	CTPerformedProcedureProtocolStorage = New("1.2.840.10008.5.1.4.1.1.200.2", "CT Performed Procedure Protocol Storage", TypeSOPClass, false)

	// ProtocolApprovalStorage Protocol Approval Storage
	ProtocolApprovalStorage = New("1.2.840.10008.5.1.4.1.1.200.3", "Protocol Approval Storage", TypeSOPClass, false)

	// ProtocolApprovalInformationModelFind Protocol Approval Information Model - FIND
	ProtocolApprovalInformationModelFind = New("1.2.840.10008.5.1.4.1.1.200.4", "Protocol Approval Information Model - FIND", TypeSOPClass, false)

	// ProtocolApprovalInformationModelMove Protocol Approval Information Model - MOVE
	ProtocolApprovalInformationModelMove = New("1.2.840.10008.5.1.4.1.1.200.5", "Protocol Approval Information Model - MOVE", TypeSOPClass, false)

	// ProtocolApprovalInformationModelGet Protocol Approval Information Model - GET
	ProtocolApprovalInformationModelGet = New("1.2.840.10008.5.1.4.1.1.200.6", "Protocol Approval Information Model - GET", TypeSOPClass, false)

	// XADefinedProcedureProtocolStorage XA Defined Procedure Protocol Storage
	XADefinedProcedureProtocolStorage = New("1.2.840.10008.5.1.4.1.1.200.7", "XA Defined Procedure Protocol Storage", TypeSOPClass, false)

	// XAPerformedProcedureProtocolStorage XA Performed Procedure Protocol Storage
	XAPerformedProcedureProtocolStorage = New("1.2.840.10008.5.1.4.1.1.200.8", "XA Performed Procedure Protocol Storage", TypeSOPClass, false)

	// InventoryStorage Inventory Storage
	InventoryStorage = New("1.2.840.10008.5.1.4.1.1.201.1", "Inventory Storage", TypeSOPClass, false)

	// InventoryFind Inventory - FIND
	InventoryFind = New("1.2.840.10008.5.1.4.1.1.201.2", "Inventory - FIND", TypeSOPClass, false)

	// InventoryMove Inventory - MOVE
	InventoryMove = New("1.2.840.10008.5.1.4.1.1.201.3", "Inventory - MOVE", TypeSOPClass, false)

	// InventoryGet Inventory - GET
	InventoryGet = New("1.2.840.10008.5.1.4.1.1.201.4", "Inventory - GET", TypeSOPClass, false)

	// InventoryCreation Inventory Creation
	InventoryCreation = New("1.2.840.10008.5.1.4.1.1.201.5", "Inventory Creation", TypeSOPClass, false)

	// RepositoryQuery Repository Query
	RepositoryQuery = New("1.2.840.10008.5.1.4.1.1.201.6", "Repository Query", TypeSOPClass, false)

	// StorageManagementInstance Storage Management SOP Instance
	StorageManagementInstance = New("1.2.840.10008.5.1.4.1.1.201.1.1", "Storage Management SOP Instance", TypeSOPInstance, false)

	// RTImageStorage RT Image Storage
	RTImageStorage = New("1.2.840.10008.5.1.4.1.1.481.1", "RT Image Storage", TypeSOPClass, false)

	// RTDoseStorage RT Dose Storage
	RTDoseStorage = New("1.2.840.10008.5.1.4.1.1.481.2", "RT Dose Storage", TypeSOPClass, false)

	// RTStructureSetStorage RT Structure Set Storage
	RTStructureSetStorage = New("1.2.840.10008.5.1.4.1.1.481.3", "RT Structure Set Storage", TypeSOPClass, false)

	// RTBeamsTreatmentRecordStorage RT Beams Treatment Record Storage
	RTBeamsTreatmentRecordStorage = New("1.2.840.10008.5.1.4.1.1.481.4", "RT Beams Treatment Record Storage", TypeSOPClass, false)

	// RTPlanStorage RT Plan Storage
	RTPlanStorage = New("1.2.840.10008.5.1.4.1.1.481.5", "RT Plan Storage", TypeSOPClass, false)

	// RTBrachyTreatmentRecordStorage RT Brachy Treatment Record Storage
	RTBrachyTreatmentRecordStorage = New("1.2.840.10008.5.1.4.1.1.481.6", "RT Brachy Treatment Record Storage", TypeSOPClass, false)

	// RTTreatmentSummaryRecordStorage RT Treatment Summary Record Storage
	RTTreatmentSummaryRecordStorage = New("1.2.840.10008.5.1.4.1.1.481.7", "RT Treatment Summary Record Storage", TypeSOPClass, false)

	// RTIonPlanStorage RT Ion Plan Storage
	RTIonPlanStorage = New("1.2.840.10008.5.1.4.1.1.481.8", "RT Ion Plan Storage", TypeSOPClass, false)

	// RTIonBeamsTreatmentRecordStorage RT Ion Beams Treatment Record Storage
	RTIonBeamsTreatmentRecordStorage = New("1.2.840.10008.5.1.4.1.1.481.9", "RT Ion Beams Treatment Record Storage", TypeSOPClass, false)

	// RTPhysicianIntentStorage RT Physician Intent Storage
	RTPhysicianIntentStorage = New("1.2.840.10008.5.1.4.1.1.481.10", "RT Physician Intent Storage", TypeSOPClass, false)

	// RTSegmentAnnotationStorage RT Segment Annotation Storage
	RTSegmentAnnotationStorage = New("1.2.840.10008.5.1.4.1.1.481.11", "RT Segment Annotation Storage", TypeSOPClass, false)

	// RTRadiationSetStorage RT Radiation Set Storage
	RTRadiationSetStorage = New("1.2.840.10008.5.1.4.1.1.481.12", "RT Radiation Set Storage", TypeSOPClass, false)

	// CArmPhotonElectronRadiationStorage C-Arm Photon-Electron Radiation Storage
	CArmPhotonElectronRadiationStorage = New("1.2.840.10008.5.1.4.1.1.481.13", "C-Arm Photon-Electron Radiation Storage", TypeSOPClass, false)

	// TomotherapeuticRadiationStorage Tomotherapeutic Radiation Storage
	TomotherapeuticRadiationStorage = New("1.2.840.10008.5.1.4.1.1.481.14", "Tomotherapeutic Radiation Storage", TypeSOPClass, false)

	// RoboticArmRadiationStorage Robotic-Arm Radiation Storage
	RoboticArmRadiationStorage = New("1.2.840.10008.5.1.4.1.1.481.15", "Robotic-Arm Radiation Storage", TypeSOPClass, false)

	// RTRadiationRecordSetStorage RT Radiation Record Set Storage
	RTRadiationRecordSetStorage = New("1.2.840.10008.5.1.4.1.1.481.16", "RT Radiation Record Set Storage", TypeSOPClass, false)

	// RTRadiationSalvageRecordStorage RT Radiation Salvage Record Storage
	RTRadiationSalvageRecordStorage = New("1.2.840.10008.5.1.4.1.1.481.17", "RT Radiation Salvage Record Storage", TypeSOPClass, false)

	// TomotherapeuticRadiationRecordStorage Tomotherapeutic Radiation Record Storage
	TomotherapeuticRadiationRecordStorage = New("1.2.840.10008.5.1.4.1.1.481.18", "Tomotherapeutic Radiation Record Storage", TypeSOPClass, false)

	// CArmPhotonElectronRadiationRecordStorage C-Arm Photon-Electron Radiation Record Storage
	CArmPhotonElectronRadiationRecordStorage = New("1.2.840.10008.5.1.4.1.1.481.19", "C-Arm Photon-Electron Radiation Record Storage", TypeSOPClass, false)

	// RoboticRadiationRecordStorage Robotic Radiation Record Storage
	RoboticRadiationRecordStorage = New("1.2.840.10008.5.1.4.1.1.481.20", "Robotic Radiation Record Storage", TypeSOPClass, false)

	// RTRadiationSetDeliveryInstructionStorage RT Radiation Set Delivery Instruction Storage
	RTRadiationSetDeliveryInstructionStorage = New("1.2.840.10008.5.1.4.1.1.481.21", "RT Radiation Set Delivery Instruction Storage", TypeSOPClass, false)

	// RTTreatmentPreparationStorage RT Treatment Preparation Storage
	RTTreatmentPreparationStorage = New("1.2.840.10008.5.1.4.1.1.481.22", "RT Treatment Preparation Storage", TypeSOPClass, false)

	// EnhancedRTImageStorage Enhanced RT Image Storage
	EnhancedRTImageStorage = New("1.2.840.10008.5.1.4.1.1.481.23", "Enhanced RT Image Storage", TypeSOPClass, false)

	// EnhancedContinuousRTImageStorage Enhanced Continuous RT Image Storage
	EnhancedContinuousRTImageStorage = New("1.2.840.10008.5.1.4.1.1.481.24", "Enhanced Continuous RT Image Storage", TypeSOPClass, false)

	// RTPatientPositionAcquisitionInstructionStorage RT Patient Position Acquisition Instruction Storage
	RTPatientPositionAcquisitionInstructionStorage = New("1.2.840.10008.5.1.4.1.1.481.25", "RT Patient Position Acquisition Instruction Storage", TypeSOPClass, false)

	// DICOSCTImageStorage DICOS CT Image Storage
	DICOSCTImageStorage = New("1.2.840.10008.5.1.4.1.1.501.1", "DICOS CT Image Storage", TypeSOPClass, false)

	// DICOSDigitalXRayImageStorageForPresentation DICOS Digital X-Ray Image Storage - For Presentation
	DICOSDigitalXRayImageStorageForPresentation = New("1.2.840.10008.5.1.4.1.1.501.2.1", "DICOS Digital X-Ray Image Storage - For Presentation", TypeSOPClass, false)

	// DICOSDigitalXRayImageStorageForProcessing DICOS Digital X-Ray Image Storage - For Processing
	DICOSDigitalXRayImageStorageForProcessing = New("1.2.840.10008.5.1.4.1.1.501.2.2", "DICOS Digital X-Ray Image Storage - For Processing", TypeSOPClass, false)

	// DICOSThreatDetectionReportStorage DICOS Threat Detection Report Storage
	DICOSThreatDetectionReportStorage = New("1.2.840.10008.5.1.4.1.1.501.3", "DICOS Threat Detection Report Storage", TypeSOPClass, false)

	// DICOS2DAITStorage DICOS 2D AIT Storage
	DICOS2DAITStorage = New("1.2.840.10008.5.1.4.1.1.501.4", "DICOS 2D AIT Storage", TypeSOPClass, false)

	// DICOS3DAITStorage DICOS 3D AIT Storage
	DICOS3DAITStorage = New("1.2.840.10008.5.1.4.1.1.501.5", "DICOS 3D AIT Storage", TypeSOPClass, false)

	// DICOSQuadrupoleResonanceStorage DICOS Quadrupole Resonance (QR) Storage
	DICOSQuadrupoleResonanceStorage = New("1.2.840.10008.5.1.4.1.1.501.6", "DICOS Quadrupole Resonance (QR) Storage", TypeSOPClass, false)

	// EddyCurrentImageStorage Eddy Current Image Storage
	EddyCurrentImageStorage = New("1.2.840.10008.5.1.4.1.1.601.1", "Eddy Current Image Storage", TypeSOPClass, false)

	// EddyCurrentMultiFrameImageStorage Eddy Current Multi-frame Image Storage
	EddyCurrentMultiFrameImageStorage = New("1.2.840.10008.5.1.4.1.1.601.2", "Eddy Current Multi-frame Image Storage", TypeSOPClass, false)

	// ThermographyImageStorage Thermography Image Storage
	ThermographyImageStorage = New("1.2.840.10008.5.1.4.1.1.601.3", "Thermography Image Storage", TypeSOPClass, false)

	// ThermographyMultiFrameImageStorage Thermography Multi-frame Image Storage
	ThermographyMultiFrameImageStorage = New("1.2.840.10008.5.1.4.1.1.601.4", "Thermography Multi-frame Image Storage", TypeSOPClass, false)

	// UltrasoundWaveformStorage Ultrasound Waveform Storage
	UltrasoundWaveformStorage = New("1.2.840.10008.5.1.4.1.1.601.5", "Ultrasound Waveform Storage", TypeSOPClass, false)

	// PatientRootQueryRetrieveInformationModelFind Patient Root Query/Retrieve Information Model - FIND
	PatientRootQueryRetrieveInformationModelFind = New("1.2.840.10008.5.1.4.1.2.1.1", "Patient Root Query/Retrieve Information Model - FIND", TypeSOPClass, false)

	// PatientRootQueryRetrieveInformationModelMove Patient Root Query/Retrieve Information Model - MOVE
	PatientRootQueryRetrieveInformationModelMove = New("1.2.840.10008.5.1.4.1.2.1.2", "Patient Root Query/Retrieve Information Model - MOVE", TypeSOPClass, false)

	// PatientRootQueryRetrieveInformationModelGet Patient Root Query/Retrieve Information Model - GET
	PatientRootQueryRetrieveInformationModelGet = New("1.2.840.10008.5.1.4.1.2.1.3", "Patient Root Query/Retrieve Information Model - GET", TypeSOPClass, false)

	// StudyRootQueryRetrieveInformationModelFind Study Root Query/Retrieve Information Model - FIND
	StudyRootQueryRetrieveInformationModelFind = New("1.2.840.10008.5.1.4.1.2.2.1", "Study Root Query/Retrieve Information Model - FIND", TypeSOPClass, false)

	// StudyRootQueryRetrieveInformationModelMove Study Root Query/Retrieve Information Model - MOVE
	StudyRootQueryRetrieveInformationModelMove = New("1.2.840.10008.5.1.4.1.2.2.2", "Study Root Query/Retrieve Information Model - MOVE", TypeSOPClass, false)

	// StudyRootQueryRetrieveInformationModelGet Study Root Query/Retrieve Information Model - GET
	StudyRootQueryRetrieveInformationModelGet = New("1.2.840.10008.5.1.4.1.2.2.3", "Study Root Query/Retrieve Information Model - GET", TypeSOPClass, false)

	// PatientStudyOnlyQueryRetrieveInformationModelFindRETIRED Patient/Study Only Query/Retrieve Information Model - FIND (Retired)
	PatientStudyOnlyQueryRetrieveInformationModelFindRETIRED = New("1.2.840.10008.5.1.4.1.2.3.1", "Patient/Study Only Query/Retrieve Information Model - FIND (Retired)", TypeSOPClass, true)

	// PatientStudyOnlyQueryRetrieveInformationModelMoveRETIRED Patient/Study Only Query/Retrieve Information Model - MOVE (Retired)
	PatientStudyOnlyQueryRetrieveInformationModelMoveRETIRED = New("1.2.840.10008.5.1.4.1.2.3.2", "Patient/Study Only Query/Retrieve Information Model - MOVE (Retired)", TypeSOPClass, true)

	// PatientStudyOnlyQueryRetrieveInformationModelGetRETIRED Patient/Study Only Query/Retrieve Information Model - GET (Retired)
	PatientStudyOnlyQueryRetrieveInformationModelGetRETIRED = New("1.2.840.10008.5.1.4.1.2.3.3", "Patient/Study Only Query/Retrieve Information Model - GET (Retired)", TypeSOPClass, true)

	// CompositeInstanceRootRetrieveMove Composite Instance Root Retrieve - MOVE
	CompositeInstanceRootRetrieveMove = New("1.2.840.10008.5.1.4.1.2.4.2", "Composite Instance Root Retrieve - MOVE", TypeSOPClass, false)

	// CompositeInstanceRootRetrieveGet Composite Instance Root Retrieve - GET
	CompositeInstanceRootRetrieveGet = New("1.2.840.10008.5.1.4.1.2.4.3", "Composite Instance Root Retrieve - GET", TypeSOPClass, false)

	// CompositeInstanceRetrieveWithoutBulkDataGet Composite Instance Retrieve Without Bulk Data - GET
	CompositeInstanceRetrieveWithoutBulkDataGet = New("1.2.840.10008.5.1.4.1.2.5.3", "Composite Instance Retrieve Without Bulk Data - GET", TypeSOPClass, false)

	// DefinedProcedureProtocolInformationModelFind Defined Procedure Protocol Information Model - FIND
	DefinedProcedureProtocolInformationModelFind = New("1.2.840.10008.5.1.4.20.1", "Defined Procedure Protocol Information Model - FIND", TypeSOPClass, false)

	// DefinedProcedureProtocolInformationModelMove Defined Procedure Protocol Information Model - MOVE
	DefinedProcedureProtocolInformationModelMove = New("1.2.840.10008.5.1.4.20.2", "Defined Procedure Protocol Information Model - MOVE", TypeSOPClass, false)

	// DefinedProcedureProtocolInformationModelGet Defined Procedure Protocol Information Model - GET
	DefinedProcedureProtocolInformationModelGet = New("1.2.840.10008.5.1.4.20.3", "Defined Procedure Protocol Information Model - GET", TypeSOPClass, false)

	// ModalityWorklistInformationModelFind Modality Worklist Information Model - FIND
	ModalityWorklistInformationModelFind = New("1.2.840.10008.5.1.4.31", "Modality Worklist Information Model - FIND", TypeSOPClass, false)

	// GeneralPurposeWorklistManagementMetaRETIRED General Purpose Worklist Management Meta SOP Class (Retired)
	GeneralPurposeWorklistManagementMetaRETIRED = New("1.2.840.10008.5.1.4.32", "General Purpose Worklist Management Meta SOP Class (Retired)", TypeMetaSOPClass, true)

	// GeneralPurposeWorklistInformationModelFindRETIRED General Purpose Worklist Information Model - FIND (Retired)
	GeneralPurposeWorklistInformationModelFindRETIRED = New("1.2.840.10008.5.1.4.32.1", "General Purpose Worklist Information Model - FIND (Retired)", TypeSOPClass, true)

	// GeneralPurposeScheduledProcedureStepRETIRED General Purpose Scheduled Procedure Step SOP Class (Retired)
	GeneralPurposeScheduledProcedureStepRETIRED = New("1.2.840.10008.5.1.4.32.2", "General Purpose Scheduled Procedure Step SOP Class (Retired)", TypeSOPClass, true)

	// GeneralPurposePerformedProcedureStepRETIRED General Purpose Performed Procedure Step SOP Class (Retired)
	GeneralPurposePerformedProcedureStepRETIRED = New("1.2.840.10008.5.1.4.32.3", "General Purpose Performed Procedure Step SOP Class (Retired)", TypeSOPClass, true)

	// InstanceAvailabilityNotification Instance Availability Notification SOP Class
	InstanceAvailabilityNotification = New("1.2.840.10008.5.1.4.33", "Instance Availability Notification SOP Class", TypeSOPClass, false)

	// RTBeamsDeliveryInstructionStorageTrialRETIRED RT Beams Delivery Instruction Storage - Trial (Retired)
	RTBeamsDeliveryInstructionStorageTrialRETIRED = New("1.2.840.10008.5.1.4.34.1", "RT Beams Delivery Instruction Storage - Trial (Retired)", TypeSOPClass, true)

	// RTConventionalMachineVerificationTrialRETIRED RT Conventional Machine Verification - Trial (Retired)
	RTConventionalMachineVerificationTrialRETIRED = New("1.2.840.10008.5.1.4.34.2", "RT Conventional Machine Verification - Trial (Retired)", TypeSOPClass, true)

	// RTIonMachineVerificationTrialRETIRED RT Ion Machine Verification - Trial (Retired)
	RTIonMachineVerificationTrialRETIRED = New("1.2.840.10008.5.1.4.34.3", "RT Ion Machine Verification - Trial (Retired)", TypeSOPClass, true)

	// UnifiedWorklistAndProcedureStepTrialRETIRED Unified Worklist and Procedure Step Service Class - Trial (Retired)
	UnifiedWorklistAndProcedureStepTrialRETIRED = New("1.2.840.10008.5.1.4.34.4", "Unified Worklist and Procedure Step Service Class - Trial (Retired)", TypeServiceClass, true)

	// UnifiedProcedureStepPushTrialRETIRED Unified Procedure Step - Push SOP Class - Trial (Retired)
	UnifiedProcedureStepPushTrialRETIRED = New("1.2.840.10008.5.1.4.34.4.1", "Unified Procedure Step - Push SOP Class - Trial (Retired)", TypeSOPClass, true)

	// UnifiedProcedureStepWatchTrialRETIRED Unified Procedure Step - Watch SOP Class - Trial (Retired)
	UnifiedProcedureStepWatchTrialRETIRED = New("1.2.840.10008.5.1.4.34.4.2", "Unified Procedure Step - Watch SOP Class - Trial (Retired)", TypeSOPClass, true)

	// UnifiedProcedureStepPullTrialRETIRED Unified Procedure Step - Pull SOP Class - Trial (Retired)
	UnifiedProcedureStepPullTrialRETIRED = New("1.2.840.10008.5.1.4.34.4.3", "Unified Procedure Step - Pull SOP Class - Trial (Retired)", TypeSOPClass, true)

	// UnifiedProcedureStepEventTrialRETIRED Unified Procedure Step - Event SOP Class - Trial (Retired)
	UnifiedProcedureStepEventTrialRETIRED = New("1.2.840.10008.5.1.4.34.4.4", "Unified Procedure Step - Event SOP Class - Trial (Retired)", TypeSOPClass, true)

	// UPSGlobalSubscriptionInstance UPS Global Subscription SOP Instance
	UPSGlobalSubscriptionInstance = New("1.2.840.10008.5.1.4.34.5", "UPS Global Subscription SOP Instance", TypeSOPInstance, false)

	// UPSFilteredGlobalSubscriptionInstance UPS Filtered Global Subscription SOP Instance
	UPSFilteredGlobalSubscriptionInstance = New("1.2.840.10008.5.1.4.34.5.1", "UPS Filtered Global Subscription SOP Instance", TypeSOPInstance, false)

	// UnifiedWorklistAndProcedureStep Unified Worklist and Procedure Step Service Class
	UnifiedWorklistAndProcedureStep = New("1.2.840.10008.5.1.4.34.6", "Unified Worklist and Procedure Step Service Class", TypeServiceClass, false)

	// UnifiedProcedureStepPush Unified Procedure Step - Push SOP Class
	UnifiedProcedureStepPush = New("1.2.840.10008.5.1.4.34.6.1", "Unified Procedure Step - Push SOP Class", TypeSOPClass, false)

	// UnifiedProcedureStepWatch Unified Procedure Step - Watch SOP Class
	UnifiedProcedureStepWatch = New("1.2.840.10008.5.1.4.34.6.2", "Unified Procedure Step - Watch SOP Class", TypeSOPClass, false)

	// UnifiedProcedureStepPull Unified Procedure Step - Pull SOP Class
	UnifiedProcedureStepPull = New("1.2.840.10008.5.1.4.34.6.3", "Unified Procedure Step - Pull SOP Class", TypeSOPClass, false)

	// UnifiedProcedureStepEvent Unified Procedure Step - Event SOP Class
	UnifiedProcedureStepEvent = New("1.2.840.10008.5.1.4.34.6.4", "Unified Procedure Step - Event SOP Class", TypeSOPClass, false)

	// UnifiedProcedureStepQuery Unified Procedure Step - Query SOP Class
	UnifiedProcedureStepQuery = New("1.2.840.10008.5.1.4.34.6.5", "Unified Procedure Step - Query SOP Class", TypeSOPClass, false)

	// RTBeamsDeliveryInstructionStorage RT Beams Delivery Instruction Storage
	RTBeamsDeliveryInstructionStorage = New("1.2.840.10008.5.1.4.34.7", "RT Beams Delivery Instruction Storage", TypeSOPClass, false)

	// RTConventionalMachineVerification RT Conventional Machine Verification
	RTConventionalMachineVerification = New("1.2.840.10008.5.1.4.34.8", "RT Conventional Machine Verification", TypeSOPClass, false)

	// RTIonMachineVerification RT Ion Machine Verification
	RTIonMachineVerification = New("1.2.840.10008.5.1.4.34.9", "RT Ion Machine Verification", TypeSOPClass, false)

	// RTBrachyApplicationSetupDeliveryInstructionStorage RT Brachy Application Setup Delivery Instruction Storage
	RTBrachyApplicationSetupDeliveryInstructionStorage = New("1.2.840.10008.5.1.4.34.10", "RT Brachy Application Setup Delivery Instruction Storage", TypeSOPClass, false)

	// GeneralRelevantPatientInformationQuery General Relevant Patient Information Query
	GeneralRelevantPatientInformationQuery = New("1.2.840.10008.5.1.4.37.1", "General Relevant Patient Information Query", TypeSOPClass, false)

	// BreastImagingRelevantPatientInformationQuery Breast Imaging Relevant Patient Information Query
	BreastImagingRelevantPatientInformationQuery = New("1.2.840.10008.5.1.4.37.2", "Breast Imaging Relevant Patient Information Query", TypeSOPClass, false)

	// CardiacRelevantPatientInformationQuery Cardiac Relevant Patient Information Query
	CardiacRelevantPatientInformationQuery = New("1.2.840.10008.5.1.4.37.3", "Cardiac Relevant Patient Information Query", TypeSOPClass, false)

	// HangingProtocolStorage Hanging Protocol Storage
	HangingProtocolStorage = New("1.2.840.10008.5.1.4.38.1", "Hanging Protocol Storage", TypeSOPClass, false)

	// HangingProtocolInformationModelFind Hanging Protocol Information Model - FIND
	HangingProtocolInformationModelFind = New("1.2.840.10008.5.1.4.38.2", "Hanging Protocol Information Model - FIND", TypeSOPClass, false)

	// HangingProtocolInformationModelMove Hanging Protocol Information Model - MOVE
	HangingProtocolInformationModelMove = New("1.2.840.10008.5.1.4.38.3", "Hanging Protocol Information Model - MOVE", TypeSOPClass, false)

	// HangingProtocolInformationModelGet Hanging Protocol Information Model - GET
	HangingProtocolInformationModelGet = New("1.2.840.10008.5.1.4.38.4", "Hanging Protocol Information Model - GET", TypeSOPClass, false)

	// ColorPaletteStorage Color Palette Storage
	ColorPaletteStorage = New("1.2.840.10008.5.1.4.39.1", "Color Palette Storage", TypeSOPClass, false)

	// ColorPaletteQueryRetrieveInformationModelFind Color Palette Query/Retrieve Information Model - FIND
	ColorPaletteQueryRetrieveInformationModelFind = New("1.2.840.10008.5.1.4.39.2", "Color Palette Query/Retrieve Information Model - FIND", TypeSOPClass, false)

	// ColorPaletteQueryRetrieveInformationModelMove Color Palette Query/Retrieve Information Model - MOVE
	ColorPaletteQueryRetrieveInformationModelMove = New("1.2.840.10008.5.1.4.39.3", "Color Palette Query/Retrieve Information Model - MOVE", TypeSOPClass, false)

	// ColorPaletteQueryRetrieveInformationModelGet Color Palette Query/Retrieve Information Model - GET
	ColorPaletteQueryRetrieveInformationModelGet = New("1.2.840.10008.5.1.4.39.4", "Color Palette Query/Retrieve Information Model - GET", TypeSOPClass, false)

	// ProductCharacteristicsQuery Product Characteristics Query SOP Class
	ProductCharacteristicsQuery = New("1.2.840.10008.5.1.4.41", "Product Characteristics Query SOP Class", TypeSOPClass, false)

	// SubstanceApprovalQuery Substance Approval Query SOP Class
	SubstanceApprovalQuery = New("1.2.840.10008.5.1.4.42", "Substance Approval Query SOP Class", TypeSOPClass, false)

	// GenericImplantTemplateStorage Generic Implant Template Storage
	GenericImplantTemplateStorage = New("1.2.840.10008.5.1.4.43.1", "Generic Implant Template Storage", TypeSOPClass, false)

	// GenericImplantTemplateInformationModelFind Generic Implant Template Information Model - FIND
	GenericImplantTemplateInformationModelFind = New("1.2.840.10008.5.1.4.43.2", "Generic Implant Template Information Model - FIND", TypeSOPClass, false)

	// GenericImplantTemplateInformationModelMove Generic Implant Template Information Model - MOVE
	GenericImplantTemplateInformationModelMove = New("1.2.840.10008.5.1.4.43.3", "Generic Implant Template Information Model - MOVE", TypeSOPClass, false)

	// GenericImplantTemplateInformationModelGet Generic Implant Template Information Model - GET
	GenericImplantTemplateInformationModelGet = New("1.2.840.10008.5.1.4.43.4", "Generic Implant Template Information Model - GET", TypeSOPClass, false)

	// ImplantAssemblyTemplateStorage Implant Assembly Template Storage
	ImplantAssemblyTemplateStorage = New("1.2.840.10008.5.1.4.44.1", "Implant Assembly Template Storage", TypeSOPClass, false)

	// ImplantAssemblyTemplateInformationModelFind Implant Assembly Template Information Model - FIND
	ImplantAssemblyTemplateInformationModelFind = New("1.2.840.10008.5.1.4.44.2", "Implant Assembly Template Information Model - FIND", TypeSOPClass, false)

	// ImplantAssemblyTemplateInformationModelMove Implant Assembly Template Information Model - MOVE
	ImplantAssemblyTemplateInformationModelMove = New("1.2.840.10008.5.1.4.44.3", "Implant Assembly Template Information Model - MOVE", TypeSOPClass, false)

	// ImplantAssemblyTemplateInformationModelGet Implant Assembly Template Information Model - GET
	ImplantAssemblyTemplateInformationModelGet = New("1.2.840.10008.5.1.4.44.4", "Implant Assembly Template Information Model - GET", TypeSOPClass, false)

	// ImplantTemplateGroupStorage Implant Template Group Storage
	ImplantTemplateGroupStorage = New("1.2.840.10008.5.1.4.45.1", "Implant Template Group Storage", TypeSOPClass, false)

	// ImplantTemplateGroupInformationModelFind Implant Template Group Information Model - FIND
	ImplantTemplateGroupInformationModelFind = New("1.2.840.10008.5.1.4.45.2", "Implant Template Group Information Model - FIND", TypeSOPClass, false)

	// ImplantTemplateGroupInformationModelMove Implant Template Group Information Model - MOVE
	ImplantTemplateGroupInformationModelMove = New("1.2.840.10008.5.1.4.45.3", "Implant Template Group Information Model - MOVE", TypeSOPClass, false)

	// ImplantTemplateGroupInformationModelGet Implant Template Group Information Model - GET
	ImplantTemplateGroupInformationModelGet = New("1.2.840.10008.5.1.4.45.4", "Implant Template Group Information Model - GET", TypeSOPClass, false)

	// NativeDICOMModel Native DICOM Model
	NativeDICOMModel = New("1.2.840.10008.7.1.1", "Native DICOM Model", TypeApplicationHostingModel, false)

	// AbstractMultiDimensionalImageModel Abstract Multi-Dimensional Image Model
	AbstractMultiDimensionalImageModel = New("1.2.840.10008.7.1.2", "Abstract Multi-Dimensional Image Model", TypeApplicationHostingModel, false)

	// DICOMContentMappingResource DICOM Content Mapping Resource
	DICOMContentMappingResource = New("1.2.840.10008.8.1.1", "DICOM Content Mapping Resource", TypeMappingResource, false)

	// VideoEndoscopicImageRealTimeCommunication Video Endoscopic Image Real-Time Communication
	VideoEndoscopicImageRealTimeCommunication = New("1.2.840.10008.10.1", "Video Endoscopic Image Real-Time Communication", TypeSOPClass, false)

	// VideoPhotographicImageRealTimeCommunication Video Photographic Image Real-Time Communication
	VideoPhotographicImageRealTimeCommunication = New("1.2.840.10008.10.2", "Video Photographic Image Real-Time Communication", TypeSOPClass, false)

	// AudioWaveformRealTimeCommunication Audio Waveform Real-Time Communication
	AudioWaveformRealTimeCommunication = New("1.2.840.10008.10.3", "Audio Waveform Real-Time Communication", TypeSOPClass, false)

	// RenditionSelectionDocumentRealTimeCommunication Rendition Selection Document Real-Time Communication
	RenditionSelectionDocumentRealTimeCommunication = New("1.2.840.10008.10.4", "Rendition Selection Document Real-Time Communication", TypeSOPClass, false)

	// dicomDeviceName dicomDeviceName
	dicomDeviceName = New("1.2.840.10008.15.0.3.1", "dicomDeviceName", TypeLDAP, false)

	// dicomDescription dicomDescription
	dicomDescription = New("1.2.840.10008.15.0.3.2", "dicomDescription", TypeLDAP, false)

	// dicomManufacturer dicomManufacturer
	dicomManufacturer = New("1.2.840.10008.15.0.3.3", "dicomManufacturer", TypeLDAP, false)

	// dicomManufacturerModelName dicomManufacturerModelName
	dicomManufacturerModelName = New("1.2.840.10008.15.0.3.4", "dicomManufacturerModelName", TypeLDAP, false)

	// dicomSoftwareVersion dicomSoftwareVersion
	dicomSoftwareVersion = New("1.2.840.10008.15.0.3.5", "dicomSoftwareVersion", TypeLDAP, false)

	// dicomVendorData dicomVendorData
	dicomVendorData = New("1.2.840.10008.15.0.3.6", "dicomVendorData", TypeLDAP, false)

	// dicomAETitle dicomAETitle
	dicomAETitle = New("1.2.840.10008.15.0.3.7", "dicomAETitle", TypeLDAP, false)

	// dicomNetworkConnectionReference dicomNetworkConnectionReference
	dicomNetworkConnectionReference = New("1.2.840.10008.15.0.3.8", "dicomNetworkConnectionReference", TypeLDAP, false)

	// dicomApplicationCluster dicomApplicationCluster
	dicomApplicationCluster = New("1.2.840.10008.15.0.3.9", "dicomApplicationCluster", TypeLDAP, false)

	// dicomAssociationInitiator dicomAssociationInitiator
	dicomAssociationInitiator = New("1.2.840.10008.15.0.3.10", "dicomAssociationInitiator", TypeLDAP, false)

	// dicomAssociationAcceptor dicomAssociationAcceptor
	dicomAssociationAcceptor = New("1.2.840.10008.15.0.3.11", "dicomAssociationAcceptor", TypeLDAP, false)

	// dicomHostname dicomHostname
	dicomHostname = New("1.2.840.10008.15.0.3.12", "dicomHostname", TypeLDAP, false)

	// dicomPort dicomPort
	dicomPort = New("1.2.840.10008.15.0.3.13", "dicomPort", TypeLDAP, false)

	// dicomSOPClass dicomSOPClass
	dicomSOPClass = New("1.2.840.10008.15.0.3.14", "dicomSOPClass", TypeLDAP, false)

	// dicomTransferRole dicomTransferRole
	dicomTransferRole = New("1.2.840.10008.15.0.3.15", "dicomTransferRole", TypeLDAP, false)

	// dicomTransferSyntax dicomTransferSyntax
	dicomTransferSyntax = New("1.2.840.10008.15.0.3.16", "dicomTransferSyntax", TypeLDAP, false)

	// dicomPrimaryDeviceType dicomPrimaryDeviceType
	dicomPrimaryDeviceType = New("1.2.840.10008.15.0.3.17", "dicomPrimaryDeviceType", TypeLDAP, false)

	// dicomRelatedDeviceReference dicomRelatedDeviceReference
	dicomRelatedDeviceReference = New("1.2.840.10008.15.0.3.18", "dicomRelatedDeviceReference", TypeLDAP, false)

	// dicomPreferredCalledAETitle dicomPreferredCalledAETitle
	dicomPreferredCalledAETitle = New("1.2.840.10008.15.0.3.19", "dicomPreferredCalledAETitle", TypeLDAP, false)

	// dicomTLSCyphersuite dicomTLSCyphersuite
	dicomTLSCyphersuite = New("1.2.840.10008.15.0.3.20", "dicomTLSCyphersuite", TypeLDAP, false)

	// dicomAuthorizedNodeCertificateReference dicomAuthorizedNodeCertificateReference
	dicomAuthorizedNodeCertificateReference = New("1.2.840.10008.15.0.3.21", "dicomAuthorizedNodeCertificateReference", TypeLDAP, false)

	// dicomThisNodeCertificateReference dicomThisNodeCertificateReference
	dicomThisNodeCertificateReference = New("1.2.840.10008.15.0.3.22", "dicomThisNodeCertificateReference", TypeLDAP, false)

	// dicomInstalled dicomInstalled
	dicomInstalled = New("1.2.840.10008.15.0.3.23", "dicomInstalled", TypeLDAP, false)

	// dicomStationName dicomStationName
	dicomStationName = New("1.2.840.10008.15.0.3.24", "dicomStationName", TypeLDAP, false)

	// dicomDeviceSerialNumber dicomDeviceSerialNumber
	dicomDeviceSerialNumber = New("1.2.840.10008.15.0.3.25", "dicomDeviceSerialNumber", TypeLDAP, false)

	// dicomInstitutionName dicomInstitutionName
	dicomInstitutionName = New("1.2.840.10008.15.0.3.26", "dicomInstitutionName", TypeLDAP, false)

	// dicomInstitutionAddress dicomInstitutionAddress
	dicomInstitutionAddress = New("1.2.840.10008.15.0.3.27", "dicomInstitutionAddress", TypeLDAP, false)

	// dicomInstitutionDepartmentName dicomInstitutionDepartmentName
	dicomInstitutionDepartmentName = New("1.2.840.10008.15.0.3.28", "dicomInstitutionDepartmentName", TypeLDAP, false)

	// dicomIssuerOfPatientID dicomIssuerOfPatientID
	dicomIssuerOfPatientID = New("1.2.840.10008.15.0.3.29", "dicomIssuerOfPatientID", TypeLDAP, false)

	// dicomPreferredCallingAETitle dicomPreferredCallingAETitle
	dicomPreferredCallingAETitle = New("1.2.840.10008.15.0.3.30", "dicomPreferredCallingAETitle", TypeLDAP, false)

	// dicomSupportedCharacterSet dicomSupportedCharacterSet
	dicomSupportedCharacterSet = New("1.2.840.10008.15.0.3.31", "dicomSupportedCharacterSet", TypeLDAP, false)

	// dicomConfigurationRoot dicomConfigurationRoot
	dicomConfigurationRoot = New("1.2.840.10008.15.0.4.1", "dicomConfigurationRoot", TypeLDAP, false)

	// dicomDevicesRoot dicomDevicesRoot
	dicomDevicesRoot = New("1.2.840.10008.15.0.4.2", "dicomDevicesRoot", TypeLDAP, false)

	// dicomUniqueAETitlesRegistryRoot dicomUniqueAETitlesRegistryRoot
	dicomUniqueAETitlesRegistryRoot = New("1.2.840.10008.15.0.4.3", "dicomUniqueAETitlesRegistryRoot", TypeLDAP, false)

	// dicomDevice dicomDevice
	dicomDevice = New("1.2.840.10008.15.0.4.4", "dicomDevice", TypeLDAP, false)

	// dicomNetworkAE dicomNetworkAE
	dicomNetworkAE = New("1.2.840.10008.15.0.4.5", "dicomNetworkAE", TypeLDAP, false)

	// dicomNetworkConnection dicomNetworkConnection
	dicomNetworkConnection = New("1.2.840.10008.15.0.4.6", "dicomNetworkConnection", TypeLDAP, false)

	// dicomUniqueAETitle dicomUniqueAETitle
	dicomUniqueAETitle = New("1.2.840.10008.15.0.4.7", "dicomUniqueAETitle", TypeLDAP, false)

	// dicomTransferCapability dicomTransferCapability
	dicomTransferCapability = New("1.2.840.10008.15.0.4.8", "dicomTransferCapability", TypeLDAP, false)

	// UTC Universal Coordinated Time
	UTC = New("1.2.840.10008.15.1.1", "Universal Coordinated Time", TypeFrameOfReference, false)

	// AnatomicModifier2 Anatomic Modifier (2)
	AnatomicModifier2 = New("1.2.840.10008.6.1.1", "Anatomic Modifier (2)", TypeContextGroupName, false)

	// AnatomicRegion4 Anatomic Region (4)
	AnatomicRegion4 = New("1.2.840.10008.6.1.2", "Anatomic Region (4)", TypeContextGroupName, false)

	// TransducerApproach5 Transducer Approach (5)
	TransducerApproach5 = New("1.2.840.10008.6.1.3", "Transducer Approach (5)", TypeContextGroupName, false)

	// TransducerOrientation6 Transducer Orientation (6)
	TransducerOrientation6 = New("1.2.840.10008.6.1.4", "Transducer Orientation (6)", TypeContextGroupName, false)

	// UltrasoundBeamPath7 Ultrasound Beam Path (7)
	UltrasoundBeamPath7 = New("1.2.840.10008.6.1.5", "Ultrasound Beam Path (7)", TypeContextGroupName, false)

	// AngiographicInterventionalDevice8 Angiographic Interventional Device (8)
	AngiographicInterventionalDevice8 = New("1.2.840.10008.6.1.6", "Angiographic Interventional Device (8)", TypeContextGroupName, false)

	// ImageGuidedTherapeuticProcedure9 Image Guided Therapeutic Procedure (9)
	ImageGuidedTherapeuticProcedure9 = New("1.2.840.10008.6.1.7", "Image Guided Therapeutic Procedure (9)", TypeContextGroupName, false)

	// InterventionalDrug10 Interventional Drug (10)
	InterventionalDrug10 = New("1.2.840.10008.6.1.8", "Interventional Drug (10)", TypeContextGroupName, false)

	// AdministrationRoute11 Administration Route (11)
	AdministrationRoute11 = New("1.2.840.10008.6.1.9", "Administration Route (11)", TypeContextGroupName, false)

	// ImagingContrastAgent12 Imaging Contrast Agent (12)
	ImagingContrastAgent12 = New("1.2.840.10008.6.1.10", "Imaging Contrast Agent (12)", TypeContextGroupName, false)

	// ImagingContrastAgentIngredient13 Imaging Contrast Agent Ingredient (13)
	ImagingContrastAgentIngredient13 = New("1.2.840.10008.6.1.11", "Imaging Contrast Agent Ingredient (13)", TypeContextGroupName, false)

	// RadiopharmaceuticalIsotope18 Radiopharmaceutical Isotope (18)
	RadiopharmaceuticalIsotope18 = New("1.2.840.10008.6.1.12", "Radiopharmaceutical Isotope (18)", TypeContextGroupName, false)

	// PatientOrientation19 Patient Orientation (19)
	PatientOrientation19 = New("1.2.840.10008.6.1.13", "Patient Orientation (19)", TypeContextGroupName, false)

	// PatientOrientationModifier20 Patient Orientation Modifier (20)
	PatientOrientationModifier20 = New("1.2.840.10008.6.1.14", "Patient Orientation Modifier (20)", TypeContextGroupName, false)

	// PatientEquipmentRelationship21 Patient Equipment Relationship (21)
	PatientEquipmentRelationship21 = New("1.2.840.10008.6.1.15", "Patient Equipment Relationship (21)", TypeContextGroupName, false)

	// CranioCaudadAngulation23 Cranio-Caudad Angulation (23)
	CranioCaudadAngulation23 = New("1.2.840.10008.6.1.16", "Cranio-Caudad Angulation (23)", TypeContextGroupName, false)

	// Radiopharmaceutical25 Radiopharmaceutical (25)
	Radiopharmaceutical25 = New("1.2.840.10008.6.1.17", "Radiopharmaceutical (25)", TypeContextGroupName, false)

	// NuclearMedicineProjection26 Nuclear Medicine Projection (26)
	NuclearMedicineProjection26 = New("1.2.840.10008.6.1.18", "Nuclear Medicine Projection (26)", TypeContextGroupName, false)

	// AcquisitionModality29 Acquisition Modality (29)
	AcquisitionModality29 = New("1.2.840.10008.6.1.19", "Acquisition Modality (29)", TypeContextGroupName, false)

	// DICOMDevice30 DICOM Device (30)
	DICOMDevice30 = New("1.2.840.10008.6.1.20", "DICOM Device (30)", TypeContextGroupName, false)

	// AbstractPrior31 Abstract Prior (31)
	AbstractPrior31 = New("1.2.840.10008.6.1.21", "Abstract Prior (31)", TypeContextGroupName, false)

	// NumericValueQualifier42 Numeric Value Qualifier (42)
	NumericValueQualifier42 = New("1.2.840.10008.6.1.22", "Numeric Value Qualifier (42)", TypeContextGroupName, false)

	// MeasurementUnit82 Measurement Unit (82)
	MeasurementUnit82 = New("1.2.840.10008.6.1.23", "Measurement Unit (82)", TypeContextGroupName, false)

	// RealWorldValueMappingUnit83 Real World Value Mapping Unit (83)
	RealWorldValueMappingUnit83 = New("1.2.840.10008.6.1.24", "Real World Value Mapping Unit (83)", TypeContextGroupName, false)

	// SignificanceLevel220 Significance Level (220)
	SignificanceLevel220 = New("1.2.840.10008.6.1.25", "Significance Level (220)", TypeContextGroupName, false)

	// MeasurementRangeConcept221 Measurement Range Concept (221)
	MeasurementRangeConcept221 = New("1.2.840.10008.6.1.26", "Measurement Range Concept (221)", TypeContextGroupName, false)

	// Normality222 Normality (222)
	Normality222 = New("1.2.840.10008.6.1.27", "Normality (222)", TypeContextGroupName, false)

	// NormalRangeValue223 Normal Range Value (223)
	NormalRangeValue223 = New("1.2.840.10008.6.1.28", "Normal Range Value (223)", TypeContextGroupName, false)

	// SelectionMethod224 Selection Method (224)
	SelectionMethod224 = New("1.2.840.10008.6.1.29", "Selection Method (224)", TypeContextGroupName, false)

	// MeasurementUncertaintyConcept225 Measurement Uncertainty Concept (225)
	MeasurementUncertaintyConcept225 = New("1.2.840.10008.6.1.30", "Measurement Uncertainty Concept (225)", TypeContextGroupName, false)

	// PopulationStatisticalDescriptor226 Population Statistical Descriptor (226)
	PopulationStatisticalDescriptor226 = New("1.2.840.10008.6.1.31", "Population Statistical Descriptor (226)", TypeContextGroupName, false)

	// SampleStatisticalDescriptor227 Sample Statistical Descriptor (227)
	SampleStatisticalDescriptor227 = New("1.2.840.10008.6.1.32", "Sample Statistical Descriptor (227)", TypeContextGroupName, false)

	// EquationOrTable228 Equation or Table (228)
	EquationOrTable228 = New("1.2.840.10008.6.1.33", "Equation or Table (228)", TypeContextGroupName, false)

	// YesNo230 Yes-No (230)
	YesNo230 = New("1.2.840.10008.6.1.34", "Yes-No (230)", TypeContextGroupName, false)

	// PresentAbsent240 Present-Absent (240)
	PresentAbsent240 = New("1.2.840.10008.6.1.35", "Present-Absent (240)", TypeContextGroupName, false)

	// NormalAbnormal242 Normal-Abnormal (242)
	NormalAbnormal242 = New("1.2.840.10008.6.1.36", "Normal-Abnormal (242)", TypeContextGroupName, false)

	// Laterality244 Laterality (244)
	Laterality244 = New("1.2.840.10008.6.1.37", "Laterality (244)", TypeContextGroupName, false)

	// PositiveNegative250 Positive-Negative (250)
	PositiveNegative250 = New("1.2.840.10008.6.1.38", "Positive-Negative (250)", TypeContextGroupName, false)

	// ComplicationSeverity251 Complication Severity (251)
	ComplicationSeverity251 = New("1.2.840.10008.6.1.39", "Complication Severity (251)", TypeContextGroupName, false)

	// ObserverType270 Observer Type (270)
	ObserverType270 = New("1.2.840.10008.6.1.40", "Observer Type (270)", TypeContextGroupName, false)

	// ObservationSubjectClass271 Observation Subject Class (271)
	ObservationSubjectClass271 = New("1.2.840.10008.6.1.41", "Observation Subject Class (271)", TypeContextGroupName, false)

	// AudioChannelSource3000 Audio Channel Source (3000)
	AudioChannelSource3000 = New("1.2.840.10008.6.1.42", "Audio Channel Source (3000)", TypeContextGroupName, false)

	// ECGLead3001 ECG Lead (3001)
	ECGLead3001 = New("1.2.840.10008.6.1.43", "ECG Lead (3001)", TypeContextGroupName, false)

	// HemodynamicWaveformSource3003 Hemodynamic Waveform Source (3003)
	HemodynamicWaveformSource3003 = New("1.2.840.10008.6.1.44", "Hemodynamic Waveform Source (3003)", TypeContextGroupName, false)

	// CardiovascularAnatomicStructure3010 Cardiovascular Anatomic Structure (3010)
	CardiovascularAnatomicStructure3010 = New("1.2.840.10008.6.1.45", "Cardiovascular Anatomic Structure (3010)", TypeContextGroupName, false)

	// ElectrophysiologyAnatomicLocation3011 Electrophysiology Anatomic Location (3011)
	ElectrophysiologyAnatomicLocation3011 = New("1.2.840.10008.6.1.46", "Electrophysiology Anatomic Location (3011)", TypeContextGroupName, false)

	// CoronaryArterySegment3014 Coronary Artery Segment (3014)
	CoronaryArterySegment3014 = New("1.2.840.10008.6.1.47", "Coronary Artery Segment (3014)", TypeContextGroupName, false)

	// CoronaryArtery3015 Coronary Artery (3015)
	CoronaryArtery3015 = New("1.2.840.10008.6.1.48", "Coronary Artery (3015)", TypeContextGroupName, false)

	// CardiovascularAnatomicStructureModifier3019 Cardiovascular Anatomic Structure Modifier (3019)
	CardiovascularAnatomicStructureModifier3019 = New("1.2.840.10008.6.1.49", "Cardiovascular Anatomic Structure Modifier (3019)", TypeContextGroupName, false)

	// CardiologyMeasurementUnit3082RETIRED Cardiology Measurement Unit (Retired) (3082)
	CardiologyMeasurementUnit3082RETIRED = New("1.2.840.10008.6.1.50", "Cardiology Measurement Unit (Retired) (3082)", TypeContextGroupName, true)

	// TimeSynchronizationChannelType3090 Time Synchronization Channel Type (3090)
	TimeSynchronizationChannelType3090 = New("1.2.840.10008.6.1.51", "Time Synchronization Channel Type (3090)", TypeContextGroupName, false)

	// CardiacProceduralStateValue3101 Cardiac Procedural State Value (3101)
	CardiacProceduralStateValue3101 = New("1.2.840.10008.6.1.52", "Cardiac Procedural State Value (3101)", TypeContextGroupName, false)

	// ElectrophysiologyMeasurementFunctionTechnique3240 Electrophysiology Measurement Function/Technique (3240)
	ElectrophysiologyMeasurementFunctionTechnique3240 = New("1.2.840.10008.6.1.53", "Electrophysiology Measurement Function/Technique (3240)", TypeContextGroupName, false)

	// HemodynamicMeasurementTechnique3241 Hemodynamic Measurement Technique (3241)
	HemodynamicMeasurementTechnique3241 = New("1.2.840.10008.6.1.54", "Hemodynamic Measurement Technique (3241)", TypeContextGroupName, false)

	// CatheterizationProcedurePhase3250 Catheterization Procedure Phase (3250)
	CatheterizationProcedurePhase3250 = New("1.2.840.10008.6.1.55", "Catheterization Procedure Phase (3250)", TypeContextGroupName, false)

	// ElectrophysiologyProcedurePhase3254 Electrophysiology Procedure Phase (3254)
	ElectrophysiologyProcedurePhase3254 = New("1.2.840.10008.6.1.56", "Electrophysiology Procedure Phase (3254)", TypeContextGroupName, false)

	// StressProtocol3261 Stress Protocol (3261)
	StressProtocol3261 = New("1.2.840.10008.6.1.57", "Stress Protocol (3261)", TypeContextGroupName, false)

	// ECGPatientStateValue3262 ECG Patient State Value (3262)
	ECGPatientStateValue3262 = New("1.2.840.10008.6.1.58", "ECG Patient State Value (3262)", TypeContextGroupName, false)

	// ElectrodePlacementValue3263 Electrode Placement Value (3263)
	ElectrodePlacementValue3263 = New("1.2.840.10008.6.1.59", "Electrode Placement Value (3263)", TypeContextGroupName, false)

	// XYZElectrodePlacementValues3264RETIRED XYZ Electrode Placement Values (Retired) (3264)
	XYZElectrodePlacementValues3264RETIRED = New("1.2.840.10008.6.1.60", "XYZ Electrode Placement Values (Retired) (3264)", TypeContextGroupName, true)

	// HemodynamicPhysiologicalChallenge3271 Hemodynamic Physiological Challenge (3271)
	HemodynamicPhysiologicalChallenge3271 = New("1.2.840.10008.6.1.61", "Hemodynamic Physiological Challenge (3271)", TypeContextGroupName, false)

	// ECGAnnotation3335 ECG Annotation (3335)
	ECGAnnotation3335 = New("1.2.840.10008.6.1.62", "ECG Annotation (3335)", TypeContextGroupName, false)

	// HemodynamicAnnotation3337 Hemodynamic Annotation (3337)
	HemodynamicAnnotation3337 = New("1.2.840.10008.6.1.63", "Hemodynamic Annotation (3337)", TypeContextGroupName, false)

	// ElectrophysiologyAnnotation3339 Electrophysiology Annotation (3339)
	ElectrophysiologyAnnotation3339 = New("1.2.840.10008.6.1.64", "Electrophysiology Annotation (3339)", TypeContextGroupName, false)

	// ProcedureLogTitle3400 Procedure Log Title (3400)
	ProcedureLogTitle3400 = New("1.2.840.10008.6.1.65", "Procedure Log Title (3400)", TypeContextGroupName, false)

	// LogNoteType3401 Log Note Type (3401)
	LogNoteType3401 = New("1.2.840.10008.6.1.66", "Log Note Type (3401)", TypeContextGroupName, false)

	// PatientStatusAndEvent3402 Patient Status and Event (3402)
	PatientStatusAndEvent3402 = New("1.2.840.10008.6.1.67", "Patient Status and Event (3402)", TypeContextGroupName, false)

	// PercutaneousEntry3403 Percutaneous Entry (3403)
	PercutaneousEntry3403 = New("1.2.840.10008.6.1.68", "Percutaneous Entry (3403)", TypeContextGroupName, false)

	// StaffAction3404 Staff Action (3404)
	StaffAction3404 = New("1.2.840.10008.6.1.69", "Staff Action (3404)", TypeContextGroupName, false)

	// ProcedureActionValue3405 Procedure Action Value (3405)
	ProcedureActionValue3405 = New("1.2.840.10008.6.1.70", "Procedure Action Value (3405)", TypeContextGroupName, false)

	// NonCoronaryTranscatheterIntervention3406 Non-coronary Transcatheter Intervention (3406)
	NonCoronaryTranscatheterIntervention3406 = New("1.2.840.10008.6.1.71", "Non-coronary Transcatheter Intervention (3406)", TypeContextGroupName, false)

	// ObjectReferencePurpose3407 Object Reference Purpose (3407)
	ObjectReferencePurpose3407 = New("1.2.840.10008.6.1.72", "Object Reference Purpose (3407)", TypeContextGroupName, false)

	// ConsumableAction3408 Consumable Action (3408)
	ConsumableAction3408 = New("1.2.840.10008.6.1.73", "Consumable Action (3408)", TypeContextGroupName, false)

	// DrugContrastAdministration3409 Drug/Contrast Administration (3409)
	DrugContrastAdministration3409 = New("1.2.840.10008.6.1.74", "Drug/Contrast Administration (3409)", TypeContextGroupName, false)

	// DrugContrastNumericParameter3410 Drug/Contrast Numeric Parameter (3410)
	DrugContrastNumericParameter3410 = New("1.2.840.10008.6.1.75", "Drug/Contrast Numeric Parameter (3410)", TypeContextGroupName, false)

	// IntracoronaryDevice3411 Intracoronary Device (3411)
	IntracoronaryDevice3411 = New("1.2.840.10008.6.1.76", "Intracoronary Device (3411)", TypeContextGroupName, false)

	// InterventionActionStatus3412 Intervention Action/Status (3412)
	InterventionActionStatus3412 = New("1.2.840.10008.6.1.77", "Intervention Action/Status (3412)", TypeContextGroupName, false)

	// AdverseOutcome3413 Adverse Outcome (3413)
	AdverseOutcome3413 = New("1.2.840.10008.6.1.78", "Adverse Outcome (3413)", TypeContextGroupName, false)

	// ProcedureUrgency3414 Procedure Urgency (3414)
	ProcedureUrgency3414 = New("1.2.840.10008.6.1.79", "Procedure Urgency (3414)", TypeContextGroupName, false)

	// CardiacRhythm3415 Cardiac Rhythm (3415)
	CardiacRhythm3415 = New("1.2.840.10008.6.1.80", "Cardiac Rhythm (3415)", TypeContextGroupName, false)

	// RespirationRhythm3416 Respiration Rhythm (3416)
	RespirationRhythm3416 = New("1.2.840.10008.6.1.81", "Respiration Rhythm (3416)", TypeContextGroupName, false)

	// LesionRisk3418 Lesion Risk (3418)
	LesionRisk3418 = New("1.2.840.10008.6.1.82", "Lesion Risk (3418)", TypeContextGroupName, false)

	// FindingTitle3419 Finding Title (3419)
	FindingTitle3419 = New("1.2.840.10008.6.1.83", "Finding Title (3419)", TypeContextGroupName, false)

	// ProcedureAction3421 Procedure Action (3421)
	ProcedureAction3421 = New("1.2.840.10008.6.1.84", "Procedure Action (3421)", TypeContextGroupName, false)

	// DeviceUseAction3422 Device Use Action (3422)
	DeviceUseAction3422 = New("1.2.840.10008.6.1.85", "Device Use Action (3422)", TypeContextGroupName, false)

	// NumericDeviceCharacteristic3423 Numeric Device Characteristic (3423)
	NumericDeviceCharacteristic3423 = New("1.2.840.10008.6.1.86", "Numeric Device Characteristic (3423)", TypeContextGroupName, false)

	// InterventionParameter3425 Intervention Parameter (3425)
	InterventionParameter3425 = New("1.2.840.10008.6.1.87", "Intervention Parameter (3425)", TypeContextGroupName, false)

	// ConsumablesParameter3426 Consumables Parameter (3426)
	ConsumablesParameter3426 = New("1.2.840.10008.6.1.88", "Consumables Parameter (3426)", TypeContextGroupName, false)

	// EquipmentEvent3427 Equipment Event (3427)
	EquipmentEvent3427 = New("1.2.840.10008.6.1.89", "Equipment Event (3427)", TypeContextGroupName, false)

	// CardiovascularImagingProcedure3428 Cardiovascular Imaging Procedure (3428)
	CardiovascularImagingProcedure3428 = New("1.2.840.10008.6.1.90", "Cardiovascular Imaging Procedure (3428)", TypeContextGroupName, false)

	// CatheterizationDevice3429 Catheterization Device (3429)
	CatheterizationDevice3429 = New("1.2.840.10008.6.1.91", "Catheterization Device (3429)", TypeContextGroupName, false)

	// DateTimeQualifier3430 DateTime Qualifier (3430)
	DateTimeQualifier3430 = New("1.2.840.10008.6.1.92", "DateTime Qualifier (3430)", TypeContextGroupName, false)

	// PeripheralPulseLocation3440 Peripheral Pulse Location (3440)
	PeripheralPulseLocation3440 = New("1.2.840.10008.6.1.93", "Peripheral Pulse Location (3440)", TypeContextGroupName, false)

	// PatientAssessment3441 Patient Assessment (3441)
	PatientAssessment3441 = New("1.2.840.10008.6.1.94", "Patient Assessment (3441)", TypeContextGroupName, false)

	// PeripheralPulseMethod3442 Peripheral Pulse Method (3442)
	PeripheralPulseMethod3442 = New("1.2.840.10008.6.1.95", "Peripheral Pulse Method (3442)", TypeContextGroupName, false)

	// SkinCondition3446 Skin Condition (3446)
	SkinCondition3446 = New("1.2.840.10008.6.1.96", "Skin Condition (3446)", TypeContextGroupName, false)

	// AirwayAssessment3448 Airway Assessment (3448)
	AirwayAssessment3448 = New("1.2.840.10008.6.1.97", "Airway Assessment (3448)", TypeContextGroupName, false)

	// CalibrationObject3451 Calibration Object (3451)
	CalibrationObject3451 = New("1.2.840.10008.6.1.98", "Calibration Object (3451)", TypeContextGroupName, false)

	// CalibrationMethod3452 Calibration Method (3452)
	CalibrationMethod3452 = New("1.2.840.10008.6.1.99", "Calibration Method (3452)", TypeContextGroupName, false)

	// CardiacVolumeMethod3453 Cardiac Volume Method (3453)
	CardiacVolumeMethod3453 = New("1.2.840.10008.6.1.100", "Cardiac Volume Method (3453)", TypeContextGroupName, false)

	// IndexMethod3455 Index Method (3455)
	IndexMethod3455 = New("1.2.840.10008.6.1.101", "Index Method (3455)", TypeContextGroupName, false)

	// SubSegmentMethod3456 Sub-segment Method (3456)
	SubSegmentMethod3456 = New("1.2.840.10008.6.1.102", "Sub-segment Method (3456)", TypeContextGroupName, false)

	// ContourRealignment3458 Contour Realignment (3458)
	ContourRealignment3458 = New("1.2.840.10008.6.1.103", "Contour Realignment (3458)", TypeContextGroupName, false)

	// CircumferentialExtent3460 Circumferential Extent (3460)
	CircumferentialExtent3460 = New("1.2.840.10008.6.1.104", "Circumferential Extent (3460)", TypeContextGroupName, false)

	// RegionalExtent3461 Regional Extent (3461)
	RegionalExtent3461 = New("1.2.840.10008.6.1.105", "Regional Extent (3461)", TypeContextGroupName, false)

	// ChamberIdentification3462 Chamber Identification (3462)
	ChamberIdentification3462 = New("1.2.840.10008.6.1.106", "Chamber Identification (3462)", TypeContextGroupName, false)

	// QAReferenceMethod3465 QA Reference Method (3465)
	QAReferenceMethod3465 = New("1.2.840.10008.6.1.107", "QA Reference Method (3465)", TypeContextGroupName, false)

	// PlaneIdentification3466 Plane Identification (3466)
	PlaneIdentification3466 = New("1.2.840.10008.6.1.108", "Plane Identification (3466)", TypeContextGroupName, false)

	// EjectionFraction3467 Ejection Fraction (3467)
	EjectionFraction3467 = New("1.2.840.10008.6.1.109", "Ejection Fraction (3467)", TypeContextGroupName, false)

	// EDVolume3468 ED Volume (3468)
	EDVolume3468 = New("1.2.840.10008.6.1.110", "ED Volume (3468)", TypeContextGroupName, false)

	// ESVolume3469 ES Volume (3469)
	ESVolume3469 = New("1.2.840.10008.6.1.111", "ES Volume (3469)", TypeContextGroupName, false)

	// VesselLumenCrossSectionalAreaCalculationMethod3470 Vessel Lumen Cross-sectional Area Calculation Method (3470)
	VesselLumenCrossSectionalAreaCalculationMethod3470 = New("1.2.840.10008.6.1.112", "Vessel Lumen Cross-sectional Area Calculation Method (3470)", TypeContextGroupName, false)

	// EstimatedVolume3471 Estimated Volume (3471)
	EstimatedVolume3471 = New("1.2.840.10008.6.1.113", "Estimated Volume (3471)", TypeContextGroupName, false)

	// CardiacContractionPhase3472 Cardiac Contraction Phase (3472)
	CardiacContractionPhase3472 = New("1.2.840.10008.6.1.114", "Cardiac Contraction Phase (3472)", TypeContextGroupName, false)

	// IVUSProcedurePhase3480 IVUS Procedure Phase (3480)
	IVUSProcedurePhase3480 = New("1.2.840.10008.6.1.115", "IVUS Procedure Phase (3480)", TypeContextGroupName, false)

	// IVUSDistanceMeasurement3481 IVUS Distance Measurement (3481)
	IVUSDistanceMeasurement3481 = New("1.2.840.10008.6.1.116", "IVUS Distance Measurement (3481)", TypeContextGroupName, false)

	// IVUSAreaMeasurement3482 IVUS Area Measurement (3482)
	IVUSAreaMeasurement3482 = New("1.2.840.10008.6.1.117", "IVUS Area Measurement (3482)", TypeContextGroupName, false)

	// IVUSLongitudinalMeasurement3483 IVUS Longitudinal Measurement (3483)
	IVUSLongitudinalMeasurement3483 = New("1.2.840.10008.6.1.118", "IVUS Longitudinal Measurement (3483)", TypeContextGroupName, false)

	// IVUSIndexRatio3484 IVUS Index/Ratio (3484)
	IVUSIndexRatio3484 = New("1.2.840.10008.6.1.119", "IVUS Index/Ratio (3484)", TypeContextGroupName, false)

	// IVUSVolumeMeasurement3485 IVUS Volume Measurement (3485)
	IVUSVolumeMeasurement3485 = New("1.2.840.10008.6.1.120", "IVUS Volume Measurement (3485)", TypeContextGroupName, false)

	// VascularMeasurementSite3486 Vascular Measurement Site (3486)
	VascularMeasurementSite3486 = New("1.2.840.10008.6.1.121", "Vascular Measurement Site (3486)", TypeContextGroupName, false)

	// IntravascularVolumetricRegion3487 Intravascular Volumetric Region (3487)
	IntravascularVolumetricRegion3487 = New("1.2.840.10008.6.1.122", "Intravascular Volumetric Region (3487)", TypeContextGroupName, false)

	// MinMaxMean3488 Min/Max/Mean (3488)
	MinMaxMean3488 = New("1.2.840.10008.6.1.123", "Min/Max/Mean (3488)", TypeContextGroupName, false)

	// CalciumDistribution3489 Calcium Distribution (3489)
	CalciumDistribution3489 = New("1.2.840.10008.6.1.124", "Calcium Distribution (3489)", TypeContextGroupName, false)

	// IVUSLesionMorphology3491 IVUS Lesion Morphology (3491)
	IVUSLesionMorphology3491 = New("1.2.840.10008.6.1.125", "IVUS Lesion Morphology (3491)", TypeContextGroupName, false)

	// VascularDissectionClassification3492 Vascular Dissection Classification (3492)
	VascularDissectionClassification3492 = New("1.2.840.10008.6.1.126", "Vascular Dissection Classification (3492)", TypeContextGroupName, false)

	// IVUSRelativeStenosisSeverity3493 IVUS Relative Stenosis Severity (3493)
	IVUSRelativeStenosisSeverity3493 = New("1.2.840.10008.6.1.127", "IVUS Relative Stenosis Severity (3493)", TypeContextGroupName, false)

	// IVUSNonMorphologicalFinding3494 IVUS Non Morphological Finding (3494)
	IVUSNonMorphologicalFinding3494 = New("1.2.840.10008.6.1.128", "IVUS Non Morphological Finding (3494)", TypeContextGroupName, false)

	// IVUSPlaqueComposition3495 IVUS Plaque Composition (3495)
	IVUSPlaqueComposition3495 = New("1.2.840.10008.6.1.129", "IVUS Plaque Composition (3495)", TypeContextGroupName, false)

	// IVUSFiducialPoint3496 IVUS Fiducial Point (3496)
	IVUSFiducialPoint3496 = New("1.2.840.10008.6.1.130", "IVUS Fiducial Point (3496)", TypeContextGroupName, false)

	// IVUSArterialMorphology3497 IVUS Arterial Morphology (3497)
	IVUSArterialMorphology3497 = New("1.2.840.10008.6.1.131", "IVUS Arterial Morphology (3497)", TypeContextGroupName, false)

	// PressureUnit3500 Pressure Unit (3500)
	PressureUnit3500 = New("1.2.840.10008.6.1.132", "Pressure Unit (3500)", TypeContextGroupName, false)

	// HemodynamicResistanceUnit3502 Hemodynamic Resistance Unit (3502)
	HemodynamicResistanceUnit3502 = New("1.2.840.10008.6.1.133", "Hemodynamic Resistance Unit (3502)", TypeContextGroupName, false)

	// IndexedHemodynamicResistanceUnit3503 Indexed Hemodynamic Resistance Unit (3503)
	IndexedHemodynamicResistanceUnit3503 = New("1.2.840.10008.6.1.134", "Indexed Hemodynamic Resistance Unit (3503)", TypeContextGroupName, false)

	// CatheterSizeUnit3510 Catheter Size Unit (3510)
	CatheterSizeUnit3510 = New("1.2.840.10008.6.1.135", "Catheter Size Unit (3510)", TypeContextGroupName, false)

	// SpecimenCollection3515 Specimen Collection (3515)
	SpecimenCollection3515 = New("1.2.840.10008.6.1.136", "Specimen Collection (3515)", TypeContextGroupName, false)

	// BloodSourceType3520 Blood Source Type (3520)
	BloodSourceType3520 = New("1.2.840.10008.6.1.137", "Blood Source Type (3520)", TypeContextGroupName, false)

	// BloodGasPressure3524 Blood Gas Pressure (3524)
	BloodGasPressure3524 = New("1.2.840.10008.6.1.138", "Blood Gas Pressure (3524)", TypeContextGroupName, false)

	// BloodGasContent3525 Blood Gas Content (3525)
	BloodGasContent3525 = New("1.2.840.10008.6.1.139", "Blood Gas Content (3525)", TypeContextGroupName, false)

	// BloodGasSaturation3526 Blood Gas Saturation (3526)
	BloodGasSaturation3526 = New("1.2.840.10008.6.1.140", "Blood Gas Saturation (3526)", TypeContextGroupName, false)

	// BloodBaseExcess3527 Blood Base Excess (3527)
	BloodBaseExcess3527 = New("1.2.840.10008.6.1.141", "Blood Base Excess (3527)", TypeContextGroupName, false)

	// BloodPH3528 Blood pH (3528)
	BloodPH3528 = New("1.2.840.10008.6.1.142", "Blood pH (3528)", TypeContextGroupName, false)

	// ArterialVenousContent3529 Arterial / Venous Content (3529)
	ArterialVenousContent3529 = New("1.2.840.10008.6.1.143", "Arterial / Venous Content (3529)", TypeContextGroupName, false)

	// OxygenAdministrationAction3530 Oxygen Administration Action (3530)
	OxygenAdministrationAction3530 = New("1.2.840.10008.6.1.144", "Oxygen Administration Action (3530)", TypeContextGroupName, false)

	// OxygenAdministration3531 Oxygen Administration (3531)
	OxygenAdministration3531 = New("1.2.840.10008.6.1.145", "Oxygen Administration (3531)", TypeContextGroupName, false)

	// CirculatorySupportAction3550 Circulatory Support Action (3550)
	CirculatorySupportAction3550 = New("1.2.840.10008.6.1.146", "Circulatory Support Action (3550)", TypeContextGroupName, false)

	// VentilationAction3551 Ventilation Action (3551)
	VentilationAction3551 = New("1.2.840.10008.6.1.147", "Ventilation Action (3551)", TypeContextGroupName, false)

	// PacingAction3552 Pacing Action (3552)
	PacingAction3552 = New("1.2.840.10008.6.1.148", "Pacing Action (3552)", TypeContextGroupName, false)

	// CirculatorySupport3553 Circulatory Support (3553)
	CirculatorySupport3553 = New("1.2.840.10008.6.1.149", "Circulatory Support (3553)", TypeContextGroupName, false)

	// Ventilation3554 Ventilation (3554)
	Ventilation3554 = New("1.2.840.10008.6.1.150", "Ventilation (3554)", TypeContextGroupName, false)

	// Pacing3555 Pacing (3555)
	Pacing3555 = New("1.2.840.10008.6.1.151", "Pacing (3555)", TypeContextGroupName, false)

	// BloodPressureMethod3560 Blood Pressure Method (3560)
	BloodPressureMethod3560 = New("1.2.840.10008.6.1.152", "Blood Pressure Method (3560)", TypeContextGroupName, false)

	// RelativeTime3600 Relative Time (3600)
	RelativeTime3600 = New("1.2.840.10008.6.1.153", "Relative Time (3600)", TypeContextGroupName, false)

	// HemodynamicPatientState3602 Hemodynamic Patient State (3602)
	HemodynamicPatientState3602 = New("1.2.840.10008.6.1.154", "Hemodynamic Patient State (3602)", TypeContextGroupName, false)

	// ArterialLesionLocation3604 Arterial Lesion Location (3604)
	ArterialLesionLocation3604 = New("1.2.840.10008.6.1.155", "Arterial Lesion Location (3604)", TypeContextGroupName, false)

	// ArterialSourceLocation3606 Arterial Source Location (3606)
	ArterialSourceLocation3606 = New("1.2.840.10008.6.1.156", "Arterial Source Location (3606)", TypeContextGroupName, false)

	// VenousSourceLocation3607 Venous Source Location (3607)
	VenousSourceLocation3607 = New("1.2.840.10008.6.1.157", "Venous Source Location (3607)", TypeContextGroupName, false)

	// AtrialSourceLocation3608 Atrial Source Location (3608)
	AtrialSourceLocation3608 = New("1.2.840.10008.6.1.158", "Atrial Source Location (3608)", TypeContextGroupName, false)

	// VentricularSourceLocation3609 Ventricular Source Location (3609)
	VentricularSourceLocation3609 = New("1.2.840.10008.6.1.159", "Ventricular Source Location (3609)", TypeContextGroupName, false)

	// GradientSourceLocation3610 Gradient Source Location (3610)
	GradientSourceLocation3610 = New("1.2.840.10008.6.1.160", "Gradient Source Location (3610)", TypeContextGroupName, false)

	// PressureMeasurement3611 Pressure Measurement (3611)
	PressureMeasurement3611 = New("1.2.840.10008.6.1.161", "Pressure Measurement (3611)", TypeContextGroupName, false)

	// BloodVelocityMeasurement3612 Blood Velocity Measurement (3612)
	BloodVelocityMeasurement3612 = New("1.2.840.10008.6.1.162", "Blood Velocity Measurement (3612)", TypeContextGroupName, false)

	// HemodynamicTimeMeasurement3613 Hemodynamic Time Measurement (3613)
	HemodynamicTimeMeasurement3613 = New("1.2.840.10008.6.1.163", "Hemodynamic Time Measurement (3613)", TypeContextGroupName, false)

	// NonMitralValveArea3614 Non-mitral Valve Area (3614)
	NonMitralValveArea3614 = New("1.2.840.10008.6.1.164", "Non-mitral Valve Area (3614)", TypeContextGroupName, false)

	// ValveArea3615 Valve Area (3615)
	ValveArea3615 = New("1.2.840.10008.6.1.165", "Valve Area (3615)", TypeContextGroupName, false)

	// HemodynamicPeriodMeasurement3616 Hemodynamic Period Measurement (3616)
	HemodynamicPeriodMeasurement3616 = New("1.2.840.10008.6.1.166", "Hemodynamic Period Measurement (3616)", TypeContextGroupName, false)

	// ValveFlow3617 Valve Flow (3617)
	ValveFlow3617 = New("1.2.840.10008.6.1.167", "Valve Flow (3617)", TypeContextGroupName, false)

	// HemodynamicFlow3618 Hemodynamic Flow (3618)
	HemodynamicFlow3618 = New("1.2.840.10008.6.1.168", "Hemodynamic Flow (3618)", TypeContextGroupName, false)

	// HemodynamicResistanceMeasurement3619 Hemodynamic Resistance Measurement (3619)
	HemodynamicResistanceMeasurement3619 = New("1.2.840.10008.6.1.169", "Hemodynamic Resistance Measurement (3619)", TypeContextGroupName, false)

	// HemodynamicRatio3620 Hemodynamic Ratio (3620)
	HemodynamicRatio3620 = New("1.2.840.10008.6.1.170", "Hemodynamic Ratio (3620)", TypeContextGroupName, false)

	// FractionalFlowReserve3621 Fractional Flow Reserve (3621)
	FractionalFlowReserve3621 = New("1.2.840.10008.6.1.171", "Fractional Flow Reserve (3621)", TypeContextGroupName, false)

	// MeasurementType3627 Measurement Type (3627)
	MeasurementType3627 = New("1.2.840.10008.6.1.172", "Measurement Type (3627)", TypeContextGroupName, false)

	// CardiacOutputMethod3628 Cardiac Output Method (3628)
	CardiacOutputMethod3628 = New("1.2.840.10008.6.1.173", "Cardiac Output Method (3628)", TypeContextGroupName, false)

	// ProcedureIntent3629 Procedure Intent (3629)
	ProcedureIntent3629 = New("1.2.840.10008.6.1.174", "Procedure Intent (3629)", TypeContextGroupName, false)

	// CardiovascularAnatomicLocation3630 Cardiovascular Anatomic Location (3630)
	CardiovascularAnatomicLocation3630 = New("1.2.840.10008.6.1.175", "Cardiovascular Anatomic Location (3630)", TypeContextGroupName, false)

	// Hypertension3640 Hypertension (3640)
	Hypertension3640 = New("1.2.840.10008.6.1.176", "Hypertension (3640)", TypeContextGroupName, false)

	// HemodynamicAssessment3641 Hemodynamic Assessment (3641)
	HemodynamicAssessment3641 = New("1.2.840.10008.6.1.177", "Hemodynamic Assessment (3641)", TypeContextGroupName, false)

	// DegreeFinding3642 Degree Finding (3642)
	DegreeFinding3642 = New("1.2.840.10008.6.1.178", "Degree Finding (3642)", TypeContextGroupName, false)

	// HemodynamicMeasurementPhase3651 Hemodynamic Measurement Phase (3651)
	HemodynamicMeasurementPhase3651 = New("1.2.840.10008.6.1.179", "Hemodynamic Measurement Phase (3651)", TypeContextGroupName, false)

	// BodySurfaceAreaEquation3663 Body Surface Area Equation (3663)
	BodySurfaceAreaEquation3663 = New("1.2.840.10008.6.1.180", "Body Surface Area Equation (3663)", TypeContextGroupName, false)

	// OxygenConsumptionEquationTable3664 Oxygen Consumption Equation/Table (3664)
	OxygenConsumptionEquationTable3664 = New("1.2.840.10008.6.1.181", "Oxygen Consumption Equation/Table (3664)", TypeContextGroupName, false)

	// P50Equation3666 P50 Equation (3666)
	P50Equation3666 = New("1.2.840.10008.6.1.182", "P50 Equation (3666)", TypeContextGroupName, false)

	// FraminghamScore3667 Framingham Score (3667)
	FraminghamScore3667 = New("1.2.840.10008.6.1.183", "Framingham Score (3667)", TypeContextGroupName, false)

	// FraminghamTable3668 Framingham Table (3668)
	FraminghamTable3668 = New("1.2.840.10008.6.1.184", "Framingham Table (3668)", TypeContextGroupName, false)

	// ECGProcedureType3670 ECG Procedure Type (3670)
	ECGProcedureType3670 = New("1.2.840.10008.6.1.185", "ECG Procedure Type (3670)", TypeContextGroupName, false)

	// ReasonForECGStudy3671 Reason for ECG Study (3671)
	ReasonForECGStudy3671 = New("1.2.840.10008.6.1.186", "Reason for ECG Study (3671)", TypeContextGroupName, false)

	// Pacemaker3672 Pacemaker (3672)
	Pacemaker3672 = New("1.2.840.10008.6.1.187", "Pacemaker (3672)", TypeContextGroupName, false)

	// Diagnosis3673RETIRED Diagnosis (Retired) (3673)
	Diagnosis3673RETIRED = New("1.2.840.10008.6.1.188", "Diagnosis (Retired) (3673)", TypeContextGroupName, true)

	// OtherFilters3675RETIRED Other Filters (Retired) (3675)
	OtherFilters3675RETIRED = New("1.2.840.10008.6.1.189", "Other Filters (Retired) (3675)", TypeContextGroupName, true)

	// LeadMeasurementTechnique3676 Lead Measurement Technique (3676)
	LeadMeasurementTechnique3676 = New("1.2.840.10008.6.1.190", "Lead Measurement Technique (3676)", TypeContextGroupName, false)

	// SummaryCodesECG3677 Summary Codes ECG (3677)
	SummaryCodesECG3677 = New("1.2.840.10008.6.1.191", "Summary Codes ECG (3677)", TypeContextGroupName, false)

	// QTCorrectionAlgorithm3678 QT Correction Algorithm (3678)
	QTCorrectionAlgorithm3678 = New("1.2.840.10008.6.1.192", "QT Correction Algorithm (3678)", TypeContextGroupName, false)

	// ECGMorphologyDescription3679RETIRED ECG Morphology Description (Retired) (3679)
	ECGMorphologyDescription3679RETIRED = New("1.2.840.10008.6.1.193", "ECG Morphology Description (Retired) (3679)", TypeContextGroupName, true)

	// ECGLeadNoiseDescription3680 ECG Lead Noise Description (3680)
	ECGLeadNoiseDescription3680 = New("1.2.840.10008.6.1.194", "ECG Lead Noise Description (3680)", TypeContextGroupName, false)

	// ECGLeadNoiseModifier3681RETIRED ECG Lead Noise Modifier (Retired) (3681)
	ECGLeadNoiseModifier3681RETIRED = New("1.2.840.10008.6.1.195", "ECG Lead Noise Modifier (Retired) (3681)", TypeContextGroupName, true)

	// Probability3682RETIRED Probability (Retired) (3682)
	Probability3682RETIRED = New("1.2.840.10008.6.1.196", "Probability (Retired) (3682)", TypeContextGroupName, true)

	// Modifier3683RETIRED Modifier (Retired) (3683)
	Modifier3683RETIRED = New("1.2.840.10008.6.1.197", "Modifier (Retired) (3683)", TypeContextGroupName, true)

	// Trend3684RETIRED Trend (Retired) (3684)
	Trend3684RETIRED = New("1.2.840.10008.6.1.198", "Trend (Retired) (3684)", TypeContextGroupName, true)

	// ConjunctiveTerm3685RETIRED Conjunctive Term (Retired) (3685)
	ConjunctiveTerm3685RETIRED = New("1.2.840.10008.6.1.199", "Conjunctive Term (Retired) (3685)", TypeContextGroupName, true)

	// ECGInterpretiveStatement3686RETIRED ECG Interpretive Statement (Retired) (3686)
	ECGInterpretiveStatement3686RETIRED = New("1.2.840.10008.6.1.200", "ECG Interpretive Statement (Retired) (3686)", TypeContextGroupName, true)

	// ElectrophysiologyWaveformDuration3687 Electrophysiology Waveform Duration (3687)
	ElectrophysiologyWaveformDuration3687 = New("1.2.840.10008.6.1.201", "Electrophysiology Waveform Duration (3687)", TypeContextGroupName, false)

	// ElectrophysiologyWaveformVoltage3688 Electrophysiology Waveform Voltage (3688)
	ElectrophysiologyWaveformVoltage3688 = New("1.2.840.10008.6.1.202", "Electrophysiology Waveform Voltage (3688)", TypeContextGroupName, false)

	// CathDiagnosis3700 Cath Diagnosis (3700)
	CathDiagnosis3700 = New("1.2.840.10008.6.1.203", "Cath Diagnosis (3700)", TypeContextGroupName, false)

	// CardiacValveTract3701 Cardiac Valve/Tract (3701)
	CardiacValveTract3701 = New("1.2.840.10008.6.1.204", "Cardiac Valve/Tract (3701)", TypeContextGroupName, false)

	// WallMotion3703 Wall Motion (3703)
	WallMotion3703 = New("1.2.840.10008.6.1.205", "Wall Motion (3703)", TypeContextGroupName, false)

	// MyocardiumWallMorphologyFinding3704 Myocardium Wall Morphology Finding (3704)
	MyocardiumWallMorphologyFinding3704 = New("1.2.840.10008.6.1.206", "Myocardium Wall Morphology Finding (3704)", TypeContextGroupName, false)

	// ChamberSize3705 Chamber Size (3705)
	ChamberSize3705 = New("1.2.840.10008.6.1.207", "Chamber Size (3705)", TypeContextGroupName, false)

	// OverallContractility3706 Overall Contractility (3706)
	OverallContractility3706 = New("1.2.840.10008.6.1.208", "Overall Contractility (3706)", TypeContextGroupName, false)

	// VSDDescription3707 VSD Description (3707)
	VSDDescription3707 = New("1.2.840.10008.6.1.209", "VSD Description (3707)", TypeContextGroupName, false)

	// AorticRootDescription3709 Aortic Root Description (3709)
	AorticRootDescription3709 = New("1.2.840.10008.6.1.210", "Aortic Root Description (3709)", TypeContextGroupName, false)

	// CoronaryDominance3710 Coronary Dominance (3710)
	CoronaryDominance3710 = New("1.2.840.10008.6.1.211", "Coronary Dominance (3710)", TypeContextGroupName, false)

	// ValvularAbnormality3711 Valvular Abnormality (3711)
	ValvularAbnormality3711 = New("1.2.840.10008.6.1.212", "Valvular Abnormality (3711)", TypeContextGroupName, false)

	// VesselDescriptor3712 Vessel Descriptor (3712)
	VesselDescriptor3712 = New("1.2.840.10008.6.1.213", "Vessel Descriptor (3712)", TypeContextGroupName, false)

	// TIMIFlowCharacteristic3713 TIMI Flow Characteristic (3713)
	TIMIFlowCharacteristic3713 = New("1.2.840.10008.6.1.214", "TIMI Flow Characteristic (3713)", TypeContextGroupName, false)

	// Thrombus3714 Thrombus (3714)
	Thrombus3714 = New("1.2.840.10008.6.1.215", "Thrombus (3714)", TypeContextGroupName, false)

	// LesionMargin3715 Lesion Margin (3715)
	LesionMargin3715 = New("1.2.840.10008.6.1.216", "Lesion Margin (3715)", TypeContextGroupName, false)

	// Severity3716 Severity (3716)
	Severity3716 = New("1.2.840.10008.6.1.217", "Severity (3716)", TypeContextGroupName, false)

	// LeftVentricleMyocardialWall17SegmentModel3717 Left Ventricle Myocardial Wall 17 Segment Model (3717)
	LeftVentricleMyocardialWall17SegmentModel3717 = New("1.2.840.10008.6.1.218", "Left Ventricle Myocardial Wall 17 Segment Model (3717)", TypeContextGroupName, false)

	// MyocardialWallSegmentsInProjection3718 Myocardial Wall Segments in Projection (3718)
	MyocardialWallSegmentsInProjection3718 = New("1.2.840.10008.6.1.219", "Myocardial Wall Segments in Projection (3718)", TypeContextGroupName, false)

	// CanadianClinicalClassification3719 Canadian Clinical Classification (3719)
	CanadianClinicalClassification3719 = New("1.2.840.10008.6.1.220", "Canadian Clinical Classification (3719)", TypeContextGroupName, false)

	// CardiacHistoryDate3720RETIRED Cardiac History Date (Retired) (3720)
	CardiacHistoryDate3720RETIRED = New("1.2.840.10008.6.1.221", "Cardiac History Date (Retired) (3720)", TypeContextGroupName, true)

	// CardiovascularSurgery3721 Cardiovascular Surgery (3721)
	CardiovascularSurgery3721 = New("1.2.840.10008.6.1.222", "Cardiovascular Surgery (3721)", TypeContextGroupName, false)

	// DiabeticTherapy3722 Diabetic Therapy (3722)
	DiabeticTherapy3722 = New("1.2.840.10008.6.1.223", "Diabetic Therapy (3722)", TypeContextGroupName, false)

	// MIType3723 MI Type (3723)
	MIType3723 = New("1.2.840.10008.6.1.224", "MI Type (3723)", TypeContextGroupName, false)

	// SmokingHistory3724 Smoking History (3724)
	SmokingHistory3724 = New("1.2.840.10008.6.1.225", "Smoking History (3724)", TypeContextGroupName, false)

	// CoronaryInterventionIndication3726 Coronary Intervention Indication (3726)
	CoronaryInterventionIndication3726 = New("1.2.840.10008.6.1.226", "Coronary Intervention Indication (3726)", TypeContextGroupName, false)

	// CatheterizationIndication3727 Catheterization Indication (3727)
	CatheterizationIndication3727 = New("1.2.840.10008.6.1.227", "Catheterization Indication (3727)", TypeContextGroupName, false)

	// CathFinding3728 Cath Finding (3728)
	CathFinding3728 = New("1.2.840.10008.6.1.228", "Cath Finding (3728)", TypeContextGroupName, false)

	// AdmissionStatus3729 Admission Status (3729)
	AdmissionStatus3729 = New("1.2.840.10008.6.1.229", "Admission Status (3729)", TypeContextGroupName, false)

	// InsurancePayor3730 Insurance Payor (3730)
	InsurancePayor3730 = New("1.2.840.10008.6.1.230", "Insurance Payor (3730)", TypeContextGroupName, false)

	// PrimaryCauseOfDeath3733 Primary Cause of Death (3733)
	PrimaryCauseOfDeath3733 = New("1.2.840.10008.6.1.231", "Primary Cause of Death (3733)", TypeContextGroupName, false)

	// AcuteCoronarySyndromeTimePeriod3735 Acute Coronary Syndrome Time Period (3735)
	AcuteCoronarySyndromeTimePeriod3735 = New("1.2.840.10008.6.1.232", "Acute Coronary Syndrome Time Period (3735)", TypeContextGroupName, false)

	// NYHAClassification3736 NYHA Classification (3736)
	NYHAClassification3736 = New("1.2.840.10008.6.1.233", "NYHA Classification (3736)", TypeContextGroupName, false)

	// IschemiaNonInvasiveTest3737 Ischemia Non-invasive Test (3737)
	IschemiaNonInvasiveTest3737 = New("1.2.840.10008.6.1.234", "Ischemia Non-invasive Test (3737)", TypeContextGroupName, false)

	// PreCathAnginaType3738 Pre-Cath Angina Type (3738)
	PreCathAnginaType3738 = New("1.2.840.10008.6.1.235", "Pre-Cath Angina Type (3738)", TypeContextGroupName, false)

	// CathProcedureType3739 Cath Procedure Type (3739)
	CathProcedureType3739 = New("1.2.840.10008.6.1.236", "Cath Procedure Type (3739)", TypeContextGroupName, false)

	// ThrombolyticAdministration3740 Thrombolytic Administration (3740)
	ThrombolyticAdministration3740 = New("1.2.840.10008.6.1.237", "Thrombolytic Administration (3740)", TypeContextGroupName, false)

	// LabVisitMedicationAdministration3741 Lab Visit Medication Administration (3741)
	LabVisitMedicationAdministration3741 = New("1.2.840.10008.6.1.238", "Lab Visit Medication Administration (3741)", TypeContextGroupName, false)

	// PCIMedicationAdministration3742 PCI Medication Administration (3742)
	PCIMedicationAdministration3742 = New("1.2.840.10008.6.1.239", "PCI Medication Administration (3742)", TypeContextGroupName, false)

	// ClopidogrelTiclopidineAdministration3743 Clopidogrel/Ticlopidine Administration (3743)
	ClopidogrelTiclopidineAdministration3743 = New("1.2.840.10008.6.1.240", "Clopidogrel/Ticlopidine Administration (3743)", TypeContextGroupName, false)

	// EFTestingMethod3744 EF Testing Method (3744)
	EFTestingMethod3744 = New("1.2.840.10008.6.1.241", "EF Testing Method (3744)", TypeContextGroupName, false)

	// CalculationMethod3745 Calculation Method (3745)
	CalculationMethod3745 = New("1.2.840.10008.6.1.242", "Calculation Method (3745)", TypeContextGroupName, false)

	// PercutaneousEntrySite3746 Percutaneous Entry Site (3746)
	PercutaneousEntrySite3746 = New("1.2.840.10008.6.1.243", "Percutaneous Entry Site (3746)", TypeContextGroupName, false)

	// PercutaneousClosure3747 Percutaneous Closure (3747)
	PercutaneousClosure3747 = New("1.2.840.10008.6.1.244", "Percutaneous Closure (3747)", TypeContextGroupName, false)

	// AngiographicEFTestingMethod3748 Angiographic EF Testing Method (3748)
	AngiographicEFTestingMethod3748 = New("1.2.840.10008.6.1.245", "Angiographic EF Testing Method (3748)", TypeContextGroupName, false)

	// PCIProcedureResult3749 PCI Procedure Result (3749)
	PCIProcedureResult3749 = New("1.2.840.10008.6.1.246", "PCI Procedure Result (3749)", TypeContextGroupName, false)

	// PreviouslyDilatedLesion3750 Previously Dilated Lesion (3750)
	PreviouslyDilatedLesion3750 = New("1.2.840.10008.6.1.247", "Previously Dilated Lesion (3750)", TypeContextGroupName, false)

	// GuidewireCrossing3752 Guidewire Crossing (3752)
	GuidewireCrossing3752 = New("1.2.840.10008.6.1.248", "Guidewire Crossing (3752)", TypeContextGroupName, false)

	// VascularComplication3754 Vascular Complication (3754)
	VascularComplication3754 = New("1.2.840.10008.6.1.249", "Vascular Complication (3754)", TypeContextGroupName, false)

	// CathComplication3755 Cath Complication (3755)
	CathComplication3755 = New("1.2.840.10008.6.1.250", "Cath Complication (3755)", TypeContextGroupName, false)

	// CardiacPatientRiskFactor3756 Cardiac Patient Risk Factor (3756)
	CardiacPatientRiskFactor3756 = New("1.2.840.10008.6.1.251", "Cardiac Patient Risk Factor (3756)", TypeContextGroupName, false)

	// CardiacDiagnosticProcedure3757 Cardiac Diagnostic Procedure (3757)
	CardiacDiagnosticProcedure3757 = New("1.2.840.10008.6.1.252", "Cardiac Diagnostic Procedure (3757)", TypeContextGroupName, false)

	// CardiovascularFamilyHistory3758 Cardiovascular Family History (3758)
	CardiovascularFamilyHistory3758 = New("1.2.840.10008.6.1.253", "Cardiovascular Family History (3758)", TypeContextGroupName, false)

	// HypertensionTherapy3760 Hypertension Therapy (3760)
	HypertensionTherapy3760 = New("1.2.840.10008.6.1.254", "Hypertension Therapy (3760)", TypeContextGroupName, false)

	// AntilipemicAgent3761 Antilipemic Agent (3761)
	AntilipemicAgent3761 = New("1.2.840.10008.6.1.255", "Antilipemic Agent (3761)", TypeContextGroupName, false)

	// AntiarrhythmicAgent3762 Antiarrhythmic Agent (3762)
	AntiarrhythmicAgent3762 = New("1.2.840.10008.6.1.256", "Antiarrhythmic Agent (3762)", TypeContextGroupName, false)

	// MyocardialInfarctionTherapy3764 Myocardial Infarction Therapy (3764)
	MyocardialInfarctionTherapy3764 = New("1.2.840.10008.6.1.257", "Myocardial Infarction Therapy (3764)", TypeContextGroupName, false)

	// ConcernType3769 Concern Type (3769)
	ConcernType3769 = New("1.2.840.10008.6.1.258", "Concern Type (3769)", TypeContextGroupName, false)

	// ProblemStatus3770 Problem Status (3770)
	ProblemStatus3770 = New("1.2.840.10008.6.1.259", "Problem Status (3770)", TypeContextGroupName, false)

	// HealthStatus3772 Health Status (3772)
	HealthStatus3772 = New("1.2.840.10008.6.1.260", "Health Status (3772)", TypeContextGroupName, false)

	// UseStatus3773 Use Status (3773)
	UseStatus3773 = New("1.2.840.10008.6.1.261", "Use Status (3773)", TypeContextGroupName, false)

	// SocialHistory3774 Social History (3774)
	SocialHistory3774 = New("1.2.840.10008.6.1.262", "Social History (3774)", TypeContextGroupName, false)

	// CardiovascularImplant3777 Cardiovascular Implant (3777)
	CardiovascularImplant3777 = New("1.2.840.10008.6.1.263", "Cardiovascular Implant (3777)", TypeContextGroupName, false)

	// PlaqueStructure3802 Plaque Structure (3802)
	PlaqueStructure3802 = New("1.2.840.10008.6.1.264", "Plaque Structure (3802)", TypeContextGroupName, false)

	// StenosisMeasurementMethod3804 Stenosis Measurement Method (3804)
	StenosisMeasurementMethod3804 = New("1.2.840.10008.6.1.265", "Stenosis Measurement Method (3804)", TypeContextGroupName, false)

	// StenosisType3805 Stenosis Type (3805)
	StenosisType3805 = New("1.2.840.10008.6.1.266", "Stenosis Type (3805)", TypeContextGroupName, false)

	// StenosisShape3806 Stenosis Shape (3806)
	StenosisShape3806 = New("1.2.840.10008.6.1.267", "Stenosis Shape (3806)", TypeContextGroupName, false)

	// VolumeMeasurementMethod3807 Volume Measurement Method (3807)
	VolumeMeasurementMethod3807 = New("1.2.840.10008.6.1.268", "Volume Measurement Method (3807)", TypeContextGroupName, false)

	// AneurysmType3808 Aneurysm Type (3808)
	AneurysmType3808 = New("1.2.840.10008.6.1.269", "Aneurysm Type (3808)", TypeContextGroupName, false)

	// AssociatedCondition3809 Associated Condition (3809)
	AssociatedCondition3809 = New("1.2.840.10008.6.1.270", "Associated Condition (3809)", TypeContextGroupName, false)

	// VascularMorphology3810 Vascular Morphology (3810)
	VascularMorphology3810 = New("1.2.840.10008.6.1.271", "Vascular Morphology (3810)", TypeContextGroupName, false)

	// StentFinding3813 Stent Finding (3813)
	StentFinding3813 = New("1.2.840.10008.6.1.272", "Stent Finding (3813)", TypeContextGroupName, false)

	// StentComposition3814 Stent Composition (3814)
	StentComposition3814 = New("1.2.840.10008.6.1.273", "Stent Composition (3814)", TypeContextGroupName, false)

	// SourceOfVascularFinding3815 Source of Vascular Finding (3815)
	SourceOfVascularFinding3815 = New("1.2.840.10008.6.1.274", "Source of Vascular Finding (3815)", TypeContextGroupName, false)

	// VascularSclerosisType3817 Vascular Sclerosis Type (3817)
	VascularSclerosisType3817 = New("1.2.840.10008.6.1.275", "Vascular Sclerosis Type (3817)", TypeContextGroupName, false)

	// NonInvasiveVascularProcedure3820 Non-invasive Vascular Procedure (3820)
	NonInvasiveVascularProcedure3820 = New("1.2.840.10008.6.1.276", "Non-invasive Vascular Procedure (3820)", TypeContextGroupName, false)

	// PapillaryMuscleIncludedExcluded3821 Papillary Muscle Included/Excluded (3821)
	PapillaryMuscleIncludedExcluded3821 = New("1.2.840.10008.6.1.277", "Papillary Muscle Included/Excluded (3821)", TypeContextGroupName, false)

	// RespiratoryStatus3823 Respiratory Status (3823)
	RespiratoryStatus3823 = New("1.2.840.10008.6.1.278", "Respiratory Status (3823)", TypeContextGroupName, false)

	// HeartRhythm3826 Heart Rhythm (3826)
	HeartRhythm3826 = New("1.2.840.10008.6.1.279", "Heart Rhythm (3826)", TypeContextGroupName, false)

	// VesselSegment3827 Vessel Segment (3827)
	VesselSegment3827 = New("1.2.840.10008.6.1.280", "Vessel Segment (3827)", TypeContextGroupName, false)

	// PulmonaryArtery3829 Pulmonary Artery (3829)
	PulmonaryArtery3829 = New("1.2.840.10008.6.1.281", "Pulmonary Artery (3829)", TypeContextGroupName, false)

	// StenosisLength3831 Stenosis Length (3831)
	StenosisLength3831 = New("1.2.840.10008.6.1.282", "Stenosis Length (3831)", TypeContextGroupName, false)

	// StenosisGrade3832 Stenosis Grade (3832)
	StenosisGrade3832 = New("1.2.840.10008.6.1.283", "Stenosis Grade (3832)", TypeContextGroupName, false)

	// CardiacEjectionFraction3833 Cardiac Ejection Fraction (3833)
	CardiacEjectionFraction3833 = New("1.2.840.10008.6.1.284", "Cardiac Ejection Fraction (3833)", TypeContextGroupName, false)

	// CardiacVolumeMeasurement3835 Cardiac Volume Measurement (3835)
	CardiacVolumeMeasurement3835 = New("1.2.840.10008.6.1.285", "Cardiac Volume Measurement (3835)", TypeContextGroupName, false)

	// TimeBasedPerfusionMeasurement3836 Time-based Perfusion Measurement (3836)
	TimeBasedPerfusionMeasurement3836 = New("1.2.840.10008.6.1.286", "Time-based Perfusion Measurement (3836)", TypeContextGroupName, false)

	// FiducialFeature3837 Fiducial Feature (3837)
	FiducialFeature3837 = New("1.2.840.10008.6.1.287", "Fiducial Feature (3837)", TypeContextGroupName, false)

	// DiameterDerivation3838 Diameter Derivation (3838)
	DiameterDerivation3838 = New("1.2.840.10008.6.1.288", "Diameter Derivation (3838)", TypeContextGroupName, false)

	// CoronaryVein3839 Coronary Vein (3839)
	CoronaryVein3839 = New("1.2.840.10008.6.1.289", "Coronary Vein (3839)", TypeContextGroupName, false)

	// PulmonaryVein3840 Pulmonary Vein (3840)
	PulmonaryVein3840 = New("1.2.840.10008.6.1.290", "Pulmonary Vein (3840)", TypeContextGroupName, false)

	// MyocardialSubsegment3843 Myocardial Subsegment (3843)
	MyocardialSubsegment3843 = New("1.2.840.10008.6.1.291", "Myocardial Subsegment (3843)", TypeContextGroupName, false)

	// PartialViewSectionForMammography4005 Partial View Section for Mammography (4005)
	PartialViewSectionForMammography4005 = New("1.2.840.10008.6.1.292", "Partial View Section for Mammography (4005)", TypeContextGroupName, false)

	// DXAnatomyImaged4009 DX Anatomy Imaged (4009)
	DXAnatomyImaged4009 = New("1.2.840.10008.6.1.293", "DX Anatomy Imaged (4009)", TypeContextGroupName, false)

	// DXView4010 DX View (4010)
	DXView4010 = New("1.2.840.10008.6.1.294", "DX View (4010)", TypeContextGroupName, false)

	// DXViewModifier4011 DX View Modifier (4011)
	DXViewModifier4011 = New("1.2.840.10008.6.1.295", "DX View Modifier (4011)", TypeContextGroupName, false)

	// ProjectionEponymousName4012 Projection Eponymous Name (4012)
	ProjectionEponymousName4012 = New("1.2.840.10008.6.1.296", "Projection Eponymous Name (4012)", TypeContextGroupName, false)

	// AnatomicRegionForMammography4013 Anatomic Region for Mammography (4013)
	AnatomicRegionForMammography4013 = New("1.2.840.10008.6.1.297", "Anatomic Region for Mammography (4013)", TypeContextGroupName, false)

	// ViewForMammography4014 View for Mammography (4014)
	ViewForMammography4014 = New("1.2.840.10008.6.1.298", "View for Mammography (4014)", TypeContextGroupName, false)

	// ViewModifierForMammography4015 View Modifier for Mammography (4015)
	ViewModifierForMammography4015 = New("1.2.840.10008.6.1.299", "View Modifier for Mammography (4015)", TypeContextGroupName, false)

	// AnatomicRegionForIntraOralRadiography4016 Anatomic Region for Intra-oral Radiography (4016)
	AnatomicRegionForIntraOralRadiography4016 = New("1.2.840.10008.6.1.300", "Anatomic Region for Intra-oral Radiography (4016)", TypeContextGroupName, false)

	// AnatomicRegionModifierForIntraOralRadiography4017 Anatomic Region Modifier for Intra-oral Radiography (4017)
	AnatomicRegionModifierForIntraOralRadiography4017 = New("1.2.840.10008.6.1.301", "Anatomic Region Modifier for Intra-oral Radiography (4017)", TypeContextGroupName, false)

	// PrimaryAnatomicStructureForIntraOralRadiographyPermanentDentitionDesignationOfTeeth4018 Primary Anatomic Structure for Intra-oral Radiography (Permanent Dentition - Designation of Teeth) (4018)
	PrimaryAnatomicStructureForIntraOralRadiographyPermanentDentitionDesignationOfTeeth4018 = New("1.2.840.10008.6.1.302", "Primary Anatomic Structure for Intra-oral Radiography (Permanent Dentition - Designation of Teeth) (4018)", TypeContextGroupName, false)

	// PrimaryAnatomicStructureForIntraOralRadiographyDeciduousDentitionDesignationOfTeeth4019 Primary Anatomic Structure for Intra-oral Radiography (Deciduous Dentition - Designation of Teeth) (4019)
	PrimaryAnatomicStructureForIntraOralRadiographyDeciduousDentitionDesignationOfTeeth4019 = New("1.2.840.10008.6.1.303", "Primary Anatomic Structure for Intra-oral Radiography (Deciduous Dentition - Designation of Teeth) (4019)", TypeContextGroupName, false)

	// PETRadionuclide4020 PET Radionuclide (4020)
	PETRadionuclide4020 = New("1.2.840.10008.6.1.304", "PET Radionuclide (4020)", TypeContextGroupName, false)

	// PETRadiopharmaceutical4021 PET Radiopharmaceutical (4021)
	PETRadiopharmaceutical4021 = New("1.2.840.10008.6.1.305", "PET Radiopharmaceutical (4021)", TypeContextGroupName, false)

	// CraniofacialAnatomicRegion4028 Craniofacial Anatomic Region (4028)
	CraniofacialAnatomicRegion4028 = New("1.2.840.10008.6.1.306", "Craniofacial Anatomic Region (4028)", TypeContextGroupName, false)

	// CTMRAndPETAnatomyImaged4030 CT, MR and PET Anatomy Imaged (4030)
	CTMRAndPETAnatomyImaged4030 = New("1.2.840.10008.6.1.307", "CT, MR and PET Anatomy Imaged (4030)", TypeContextGroupName, false)

	// CommonAnatomicRegion4031 Common Anatomic Region (4031)
	CommonAnatomicRegion4031 = New("1.2.840.10008.6.1.308", "Common Anatomic Region (4031)", TypeContextGroupName, false)

	// MRSpectroscopyMetabolite4032 MR Spectroscopy Metabolite (4032)
	MRSpectroscopyMetabolite4032 = New("1.2.840.10008.6.1.309", "MR Spectroscopy Metabolite (4032)", TypeContextGroupName, false)

	// MRProtonSpectroscopyMetabolite4033 MR Proton Spectroscopy Metabolite (4033)
	MRProtonSpectroscopyMetabolite4033 = New("1.2.840.10008.6.1.310", "MR Proton Spectroscopy Metabolite (4033)", TypeContextGroupName, false)

	// EndoscopyAnatomicRegion4040 Endoscopy Anatomic Region (4040)
	EndoscopyAnatomicRegion4040 = New("1.2.840.10008.6.1.311", "Endoscopy Anatomic Region (4040)", TypeContextGroupName, false)

	// XAXRFAnatomyImaged4042 XA/XRF Anatomy Imaged (4042)
	XAXRFAnatomyImaged4042 = New("1.2.840.10008.6.1.312", "XA/XRF Anatomy Imaged (4042)", TypeContextGroupName, false)

	// DrugOrContrastAgentCharacteristic4050 Drug or Contrast Agent Characteristic (4050)
	DrugOrContrastAgentCharacteristic4050 = New("1.2.840.10008.6.1.313", "Drug or Contrast Agent Characteristic (4050)", TypeContextGroupName, false)

	// GeneralDevice4051 General Device (4051)
	GeneralDevice4051 = New("1.2.840.10008.6.1.314", "General Device (4051)", TypeContextGroupName, false)

	// PhantomDevice4052 Phantom Device (4052)
	PhantomDevice4052 = New("1.2.840.10008.6.1.315", "Phantom Device (4052)", TypeContextGroupName, false)

	// OphthalmicImagingAgent4200 Ophthalmic Imaging Agent (4200)
	OphthalmicImagingAgent4200 = New("1.2.840.10008.6.1.316", "Ophthalmic Imaging Agent (4200)", TypeContextGroupName, false)

	// PatientEyeMovementCommand4201 Patient Eye Movement Command (4201)
	PatientEyeMovementCommand4201 = New("1.2.840.10008.6.1.317", "Patient Eye Movement Command (4201)", TypeContextGroupName, false)

	// OphthalmicPhotographyAcquisitionDevice4202 Ophthalmic Photography Acquisition Device (4202)
	OphthalmicPhotographyAcquisitionDevice4202 = New("1.2.840.10008.6.1.318", "Ophthalmic Photography Acquisition Device (4202)", TypeContextGroupName, false)

	// OphthalmicPhotographyIllumination4203 Ophthalmic Photography Illumination (4203)
	OphthalmicPhotographyIllumination4203 = New("1.2.840.10008.6.1.319", "Ophthalmic Photography Illumination (4203)", TypeContextGroupName, false)

	// OphthalmicFilter4204 Ophthalmic Filter (4204)
	OphthalmicFilter4204 = New("1.2.840.10008.6.1.320", "Ophthalmic Filter (4204)", TypeContextGroupName, false)

	// OphthalmicLens4205 Ophthalmic Lens (4205)
	OphthalmicLens4205 = New("1.2.840.10008.6.1.321", "Ophthalmic Lens (4205)", TypeContextGroupName, false)

	// OphthalmicChannelDescription4206 Ophthalmic Channel Description (4206)
	OphthalmicChannelDescription4206 = New("1.2.840.10008.6.1.322", "Ophthalmic Channel Description (4206)", TypeContextGroupName, false)

	// OphthalmicImagePosition4207 Ophthalmic Image Position (4207)
	OphthalmicImagePosition4207 = New("1.2.840.10008.6.1.323", "Ophthalmic Image Position (4207)", TypeContextGroupName, false)

	// MydriaticAgent4208 Mydriatic Agent (4208)
	MydriaticAgent4208 = New("1.2.840.10008.6.1.324", "Mydriatic Agent (4208)", TypeContextGroupName, false)

	// OphthalmicAnatomicStructureImaged4209 Ophthalmic Anatomic Structure Imaged (4209)
	OphthalmicAnatomicStructureImaged4209 = New("1.2.840.10008.6.1.325", "Ophthalmic Anatomic Structure Imaged (4209)", TypeContextGroupName, false)

	// OphthalmicTomographyAcquisitionDevice4210 Ophthalmic Tomography Acquisition Device (4210)
	OphthalmicTomographyAcquisitionDevice4210 = New("1.2.840.10008.6.1.326", "Ophthalmic Tomography Acquisition Device (4210)", TypeContextGroupName, false)

	// OphthalmicOCTAnatomicStructureImaged4211 Ophthalmic OCT Anatomic Structure Imaged (4211)
	OphthalmicOCTAnatomicStructureImaged4211 = New("1.2.840.10008.6.1.327", "Ophthalmic OCT Anatomic Structure Imaged (4211)", TypeContextGroupName, false)

	// Language5000 Language (5000)
	Language5000 = New("1.2.840.10008.6.1.328", "Language (5000)", TypeContextGroupName, false)

	// Country5001 Country (5001)
	Country5001 = New("1.2.840.10008.6.1.329", "Country (5001)", TypeContextGroupName, false)

	// OverallBreastComposition6000 Overall Breast Composition (6000)
	OverallBreastComposition6000 = New("1.2.840.10008.6.1.330", "Overall Breast Composition (6000)", TypeContextGroupName, false)

	// OverallBreastCompositionFromBIRADS6001 Overall Breast Composition from BI-RADS® (6001)
	OverallBreastCompositionFromBIRADS6001 = New("1.2.840.10008.6.1.331", "Overall Breast Composition from BI-RADS® (6001)", TypeContextGroupName, false)

	// ChangeSinceLastMammogramOrPriorSurgery6002 Change Since Last Mammogram or Prior Surgery (6002)
	ChangeSinceLastMammogramOrPriorSurgery6002 = New("1.2.840.10008.6.1.332", "Change Since Last Mammogram or Prior Surgery (6002)", TypeContextGroupName, false)

	// ChangeSinceLastMammogramOrPriorSurgeryFromBIRADS6003 Change Since Last Mammogram or Prior Surgery from BI-RADS® (6003)
	ChangeSinceLastMammogramOrPriorSurgeryFromBIRADS6003 = New("1.2.840.10008.6.1.333", "Change Since Last Mammogram or Prior Surgery from BI-RADS® (6003)", TypeContextGroupName, false)

	// MammographyShapeCharacteristic6004 Mammography Shape Characteristic (6004)
	MammographyShapeCharacteristic6004 = New("1.2.840.10008.6.1.334", "Mammography Shape Characteristic (6004)", TypeContextGroupName, false)

	// ShapeCharacteristicFromBIRADS6005 Shape Characteristic from BI-RADS® (6005)
	ShapeCharacteristicFromBIRADS6005 = New("1.2.840.10008.6.1.335", "Shape Characteristic from BI-RADS® (6005)", TypeContextGroupName, false)

	// MammographyMarginCharacteristic6006 Mammography Margin Characteristic (6006)
	MammographyMarginCharacteristic6006 = New("1.2.840.10008.6.1.336", "Mammography Margin Characteristic (6006)", TypeContextGroupName, false)

	// MarginCharacteristicFromBIRADS6007 Margin Characteristic from BI-RADS® (6007)
	MarginCharacteristicFromBIRADS6007 = New("1.2.840.10008.6.1.337", "Margin Characteristic from BI-RADS® (6007)", TypeContextGroupName, false)

	// DensityModifier6008 Density Modifier (6008)
	DensityModifier6008 = New("1.2.840.10008.6.1.338", "Density Modifier (6008)", TypeContextGroupName, false)

	// DensityModifierFromBIRADS6009 Density Modifier from BI-RADS® (6009)
	DensityModifierFromBIRADS6009 = New("1.2.840.10008.6.1.339", "Density Modifier from BI-RADS® (6009)", TypeContextGroupName, false)

	// MammographyCalcificationType6010 Mammography Calcification Type (6010)
	MammographyCalcificationType6010 = New("1.2.840.10008.6.1.340", "Mammography Calcification Type (6010)", TypeContextGroupName, false)

	// CalcificationTypeFromBIRADS6011 Calcification Type from BI-RADS® (6011)
	CalcificationTypeFromBIRADS6011 = New("1.2.840.10008.6.1.341", "Calcification Type from BI-RADS® (6011)", TypeContextGroupName, false)

	// CalcificationDistributionModifier6012 Calcification Distribution Modifier (6012)
	CalcificationDistributionModifier6012 = New("1.2.840.10008.6.1.342", "Calcification Distribution Modifier (6012)", TypeContextGroupName, false)

	// CalcificationDistributionModifierFromBIRADS6013 Calcification Distribution Modifier from BI-RADS® (6013)
	CalcificationDistributionModifierFromBIRADS6013 = New("1.2.840.10008.6.1.343", "Calcification Distribution Modifier from BI-RADS® (6013)", TypeContextGroupName, false)

	// MammographySingleImageFinding6014 Mammography Single Image Finding (6014)
	MammographySingleImageFinding6014 = New("1.2.840.10008.6.1.344", "Mammography Single Image Finding (6014)", TypeContextGroupName, false)

	// SingleImageFindingFromBIRADS6015 Single Image Finding from BI-RADS® (6015)
	SingleImageFindingFromBIRADS6015 = New("1.2.840.10008.6.1.345", "Single Image Finding from BI-RADS® (6015)", TypeContextGroupName, false)

	// MammographyCompositeFeature6016 Mammography Composite Feature (6016)
	MammographyCompositeFeature6016 = New("1.2.840.10008.6.1.346", "Mammography Composite Feature (6016)", TypeContextGroupName, false)

	// CompositeFeatureFromBIRADS6017 Composite Feature from BI-RADS® (6017)
	CompositeFeatureFromBIRADS6017 = New("1.2.840.10008.6.1.347", "Composite Feature from BI-RADS® (6017)", TypeContextGroupName, false)

	// ClockfaceLocationOrRegion6018 Clockface Location or Region (6018)
	ClockfaceLocationOrRegion6018 = New("1.2.840.10008.6.1.348", "Clockface Location or Region (6018)", TypeContextGroupName, false)

	// ClockfaceLocationOrRegionFromBIRADS6019 Clockface Location or Region from BI-RADS® (6019)
	ClockfaceLocationOrRegionFromBIRADS6019 = New("1.2.840.10008.6.1.349", "Clockface Location or Region from BI-RADS® (6019)", TypeContextGroupName, false)

	// QuadrantLocation6020 Quadrant Location (6020)
	QuadrantLocation6020 = New("1.2.840.10008.6.1.350", "Quadrant Location (6020)", TypeContextGroupName, false)

	// QuadrantLocationFromBIRADS6021 Quadrant Location from BI-RADS® (6021)
	QuadrantLocationFromBIRADS6021 = New("1.2.840.10008.6.1.351", "Quadrant Location from BI-RADS® (6021)", TypeContextGroupName, false)

	// Side6022 Side (6022)
	Side6022 = New("1.2.840.10008.6.1.352", "Side (6022)", TypeContextGroupName, false)

	// SideFromBIRADS6023 Side from BI-RADS® (6023)
	SideFromBIRADS6023 = New("1.2.840.10008.6.1.353", "Side from BI-RADS® (6023)", TypeContextGroupName, false)

	// Depth6024 Depth (6024)
	Depth6024 = New("1.2.840.10008.6.1.354", "Depth (6024)", TypeContextGroupName, false)

	// DepthFromBIRADS6025 Depth from BI-RADS® (6025)
	DepthFromBIRADS6025 = New("1.2.840.10008.6.1.355", "Depth from BI-RADS® (6025)", TypeContextGroupName, false)

	// MammographyAssessment6026 Mammography Assessment (6026)
	MammographyAssessment6026 = New("1.2.840.10008.6.1.356", "Mammography Assessment (6026)", TypeContextGroupName, false)

	// AssessmentFromBIRADS6027 Assessment from BI-RADS® (6027)
	AssessmentFromBIRADS6027 = New("1.2.840.10008.6.1.357", "Assessment from BI-RADS® (6027)", TypeContextGroupName, false)

	// MammographyRecommendedFollowUp6028 Mammography Recommended Follow-up (6028)
	MammographyRecommendedFollowUp6028 = New("1.2.840.10008.6.1.358", "Mammography Recommended Follow-up (6028)", TypeContextGroupName, false)

	// RecommendedFollowUpFromBIRADS6029 Recommended Follow-up from BI-RADS® (6029)
	RecommendedFollowUpFromBIRADS6029 = New("1.2.840.10008.6.1.359", "Recommended Follow-up from BI-RADS® (6029)", TypeContextGroupName, false)

	// MammographyPathologyCode6030 Mammography Pathology Code (6030)
	MammographyPathologyCode6030 = New("1.2.840.10008.6.1.360", "Mammography Pathology Code (6030)", TypeContextGroupName, false)

	// BenignPathologyCodeFromBIRADS6031 Benign Pathology Code from BI-RADS® (6031)
	BenignPathologyCodeFromBIRADS6031 = New("1.2.840.10008.6.1.361", "Benign Pathology Code from BI-RADS® (6031)", TypeContextGroupName, false)

	// HighRiskLesionPathologyCodeFromBIRADS6032 High Risk Lesion Pathology Code from BI-RADS® (6032)
	HighRiskLesionPathologyCodeFromBIRADS6032 = New("1.2.840.10008.6.1.362", "High Risk Lesion Pathology Code from BI-RADS® (6032)", TypeContextGroupName, false)

	// MalignantPathologyCodeFromBIRADS6033 Malignant Pathology Code from BI-RADS® (6033)
	MalignantPathologyCodeFromBIRADS6033 = New("1.2.840.10008.6.1.363", "Malignant Pathology Code from BI-RADS® (6033)", TypeContextGroupName, false)

	// CADOutputIntendedUse6034 CAD Output Intended Use (6034)
	CADOutputIntendedUse6034 = New("1.2.840.10008.6.1.364", "CAD Output Intended Use (6034)", TypeContextGroupName, false)

	// CompositeFeatureRelation6035 Composite Feature Relation (6035)
	CompositeFeatureRelation6035 = New("1.2.840.10008.6.1.365", "Composite Feature Relation (6035)", TypeContextGroupName, false)

	// FeatureScope6036 Feature Scope (6036)
	FeatureScope6036 = New("1.2.840.10008.6.1.366", "Feature Scope (6036)", TypeContextGroupName, false)

	// MammographyQuantitativeTemporalDifferenceType6037 Mammography Quantitative Temporal Difference Type (6037)
	MammographyQuantitativeTemporalDifferenceType6037 = New("1.2.840.10008.6.1.367", "Mammography Quantitative Temporal Difference Type (6037)", TypeContextGroupName, false)

	// MammographyQualitativeTemporalDifferenceType6038 Mammography Qualitative Temporal Difference Type (6038)
	MammographyQualitativeTemporalDifferenceType6038 = New("1.2.840.10008.6.1.368", "Mammography Qualitative Temporal Difference Type (6038)", TypeContextGroupName, false)

	// NippleCharacteristic6039 Nipple Characteristic (6039)
	NippleCharacteristic6039 = New("1.2.840.10008.6.1.369", "Nipple Characteristic (6039)", TypeContextGroupName, false)

	// NonLesionObjectType6040 Non-lesion Object Type (6040)
	NonLesionObjectType6040 = New("1.2.840.10008.6.1.370", "Non-lesion Object Type (6040)", TypeContextGroupName, false)

	// MammographyImageQualityFinding6041 Mammography Image Quality Finding (6041)
	MammographyImageQualityFinding6041 = New("1.2.840.10008.6.1.371", "Mammography Image Quality Finding (6041)", TypeContextGroupName, false)

	// ResultStatus6042 Result Status (6042)
	ResultStatus6042 = New("1.2.840.10008.6.1.372", "Result Status (6042)", TypeContextGroupName, false)

	// MammographyCADAnalysisType6043 Mammography CAD Analysis Type (6043)
	MammographyCADAnalysisType6043 = New("1.2.840.10008.6.1.373", "Mammography CAD Analysis Type (6043)", TypeContextGroupName, false)

	// ImageQualityAssessmentType6044 Image Quality Assessment Type (6044)
	ImageQualityAssessmentType6044 = New("1.2.840.10008.6.1.374", "Image Quality Assessment Type (6044)", TypeContextGroupName, false)

	// MammographyQualityControlStandardType6045 Mammography Quality Control Standard Type (6045)
	MammographyQualityControlStandardType6045 = New("1.2.840.10008.6.1.375", "Mammography Quality Control Standard Type (6045)", TypeContextGroupName, false)

	// FollowUpIntervalUnit6046 Follow-up Interval Unit (6046)
	FollowUpIntervalUnit6046 = New("1.2.840.10008.6.1.376", "Follow-up Interval Unit (6046)", TypeContextGroupName, false)

	// CADProcessingAndFindingSummary6047 CAD Processing and Finding Summary (6047)
	CADProcessingAndFindingSummary6047 = New("1.2.840.10008.6.1.377", "CAD Processing and Finding Summary (6047)", TypeContextGroupName, false)

	// CADOperatingPointAxisLabel6048 CAD Operating Point Axis Label (6048)
	CADOperatingPointAxisLabel6048 = New("1.2.840.10008.6.1.378", "CAD Operating Point Axis Label (6048)", TypeContextGroupName, false)

	// BreastProcedureReported6050 Breast Procedure Reported (6050)
	BreastProcedureReported6050 = New("1.2.840.10008.6.1.379", "Breast Procedure Reported (6050)", TypeContextGroupName, false)

	// BreastProcedureReason6051 Breast Procedure Reason (6051)
	BreastProcedureReason6051 = New("1.2.840.10008.6.1.380", "Breast Procedure Reason (6051)", TypeContextGroupName, false)

	// BreastImagingReportSectionTitle6052 Breast Imaging Report Section Title (6052)
	BreastImagingReportSectionTitle6052 = New("1.2.840.10008.6.1.381", "Breast Imaging Report Section Title (6052)", TypeContextGroupName, false)

	// BreastImagingReportElement6053 Breast Imaging Report Element (6053)
	BreastImagingReportElement6053 = New("1.2.840.10008.6.1.382", "Breast Imaging Report Element (6053)", TypeContextGroupName, false)

	// BreastImagingFinding6054 Breast Imaging Finding (6054)
	BreastImagingFinding6054 = New("1.2.840.10008.6.1.383", "Breast Imaging Finding (6054)", TypeContextGroupName, false)

	// BreastClinicalFindingOrIndicatedProblem6055 Breast Clinical Finding or Indicated Problem (6055)
	BreastClinicalFindingOrIndicatedProblem6055 = New("1.2.840.10008.6.1.384", "Breast Clinical Finding or Indicated Problem (6055)", TypeContextGroupName, false)

	// AssociatedFindingForBreast6056 Associated Finding for Breast (6056)
	AssociatedFindingForBreast6056 = New("1.2.840.10008.6.1.385", "Associated Finding for Breast (6056)", TypeContextGroupName, false)

	// DuctographyFindingForBreast6057 Ductography Finding for Breast (6057)
	DuctographyFindingForBreast6057 = New("1.2.840.10008.6.1.386", "Ductography Finding for Breast (6057)", TypeContextGroupName, false)

	// ProcedureModifiersForBreast6058 Procedure Modifiers for Breast (6058)
	ProcedureModifiersForBreast6058 = New("1.2.840.10008.6.1.387", "Procedure Modifiers for Breast (6058)", TypeContextGroupName, false)

	// BreastImplantType6059 Breast Implant Type (6059)
	BreastImplantType6059 = New("1.2.840.10008.6.1.388", "Breast Implant Type (6059)", TypeContextGroupName, false)

	// BreastBiopsyTechnique6060 Breast Biopsy Technique (6060)
	BreastBiopsyTechnique6060 = New("1.2.840.10008.6.1.389", "Breast Biopsy Technique (6060)", TypeContextGroupName, false)

	// BreastImagingProcedureModifier6061 Breast Imaging Procedure Modifier (6061)
	BreastImagingProcedureModifier6061 = New("1.2.840.10008.6.1.390", "Breast Imaging Procedure Modifier (6061)", TypeContextGroupName, false)

	// InterventionalProcedureComplication6062 Interventional Procedure Complication (6062)
	InterventionalProcedureComplication6062 = New("1.2.840.10008.6.1.391", "Interventional Procedure Complication (6062)", TypeContextGroupName, false)

	// InterventionalProcedureResult6063 Interventional Procedure Result (6063)
	InterventionalProcedureResult6063 = New("1.2.840.10008.6.1.392", "Interventional Procedure Result (6063)", TypeContextGroupName, false)

	// UltrasoundFindingForBreast6064 Ultrasound Finding for Breast (6064)
	UltrasoundFindingForBreast6064 = New("1.2.840.10008.6.1.393", "Ultrasound Finding for Breast (6064)", TypeContextGroupName, false)

	// InstrumentApproach6065 Instrument Approach (6065)
	InstrumentApproach6065 = New("1.2.840.10008.6.1.394", "Instrument Approach (6065)", TypeContextGroupName, false)

	// TargetConfirmation6066 Target Confirmation (6066)
	TargetConfirmation6066 = New("1.2.840.10008.6.1.395", "Target Confirmation (6066)", TypeContextGroupName, false)

	// FluidColor6067 Fluid Color (6067)
	FluidColor6067 = New("1.2.840.10008.6.1.396", "Fluid Color (6067)", TypeContextGroupName, false)

	// TumorStagesFromAJCC6068 Tumor Stages From AJCC (6068)
	TumorStagesFromAJCC6068 = New("1.2.840.10008.6.1.397", "Tumor Stages From AJCC (6068)", TypeContextGroupName, false)

	// NottinghamCombinedHistologicGrade6069 Nottingham Combined Histologic Grade (6069)
	NottinghamCombinedHistologicGrade6069 = New("1.2.840.10008.6.1.398", "Nottingham Combined Histologic Grade (6069)", TypeContextGroupName, false)

	// BloomRichardsonHistologicGrade6070 Bloom-Richardson Histologic Grade (6070)
	BloomRichardsonHistologicGrade6070 = New("1.2.840.10008.6.1.399", "Bloom-Richardson Histologic Grade (6070)", TypeContextGroupName, false)

	// HistologicGradingMethod6071 Histologic Grading Method (6071)
	HistologicGradingMethod6071 = New("1.2.840.10008.6.1.400", "Histologic Grading Method (6071)", TypeContextGroupName, false)

	// BreastImplantFinding6072 Breast Implant Finding (6072)
	BreastImplantFinding6072 = New("1.2.840.10008.6.1.401", "Breast Implant Finding (6072)", TypeContextGroupName, false)

	// GynecologicalHormone6080 Gynecological Hormone (6080)
	GynecologicalHormone6080 = New("1.2.840.10008.6.1.402", "Gynecological Hormone (6080)", TypeContextGroupName, false)

	// BreastCancerRiskFactor6081 Breast Cancer Risk Factor (6081)
	BreastCancerRiskFactor6081 = New("1.2.840.10008.6.1.403", "Breast Cancer Risk Factor (6081)", TypeContextGroupName, false)

	// GynecologicalProcedure6082 Gynecological Procedure (6082)
	GynecologicalProcedure6082 = New("1.2.840.10008.6.1.404", "Gynecological Procedure (6082)", TypeContextGroupName, false)

	// ProceduresForBreast6083 Procedures for Breast (6083)
	ProceduresForBreast6083 = New("1.2.840.10008.6.1.405", "Procedures for Breast (6083)", TypeContextGroupName, false)

	// MammoplastyProcedure6084 Mammoplasty Procedure (6084)
	MammoplastyProcedure6084 = New("1.2.840.10008.6.1.406", "Mammoplasty Procedure (6084)", TypeContextGroupName, false)

	// TherapiesForBreast6085 Therapies for Breast (6085)
	TherapiesForBreast6085 = New("1.2.840.10008.6.1.407", "Therapies for Breast (6085)", TypeContextGroupName, false)

	// MenopausalPhase6086 Menopausal Phase (6086)
	MenopausalPhase6086 = New("1.2.840.10008.6.1.408", "Menopausal Phase (6086)", TypeContextGroupName, false)

	// GeneralRiskFactor6087 General Risk Factor (6087)
	GeneralRiskFactor6087 = New("1.2.840.10008.6.1.409", "General Risk Factor (6087)", TypeContextGroupName, false)

	// OBGYNMaternalRiskFactor6088 OB-GYN Maternal Risk Factor (6088)
	OBGYNMaternalRiskFactor6088 = New("1.2.840.10008.6.1.410", "OB-GYN Maternal Risk Factor (6088)", TypeContextGroupName, false)

	// Substance6089 Substance (6089)
	Substance6089 = New("1.2.840.10008.6.1.411", "Substance (6089)", TypeContextGroupName, false)

	// RelativeUsageExposureAmount6090 Relative Usage/Exposure Amount (6090)
	RelativeUsageExposureAmount6090 = New("1.2.840.10008.6.1.412", "Relative Usage/Exposure Amount (6090)", TypeContextGroupName, false)

	// RelativeFrequencyOfEventValue6091 Relative Frequency of Event Value (6091)
	RelativeFrequencyOfEventValue6091 = New("1.2.840.10008.6.1.413", "Relative Frequency of Event Value (6091)", TypeContextGroupName, false)

	// UsageExposureQualitativeConcept6092 Usage/Exposure Qualitative Concept (6092)
	UsageExposureQualitativeConcept6092 = New("1.2.840.10008.6.1.414", "Usage/Exposure Qualitative Concept (6092)", TypeContextGroupName, false)

	// UsageExposureAmountQualitativeConcept6093 Usage/Exposure/Amount Qualitative Concept (6093)
	UsageExposureAmountQualitativeConcept6093 = New("1.2.840.10008.6.1.415", "Usage/Exposure/Amount Qualitative Concept (6093)", TypeContextGroupName, false)

	// UsageExposureFrequencyQualitativeConcept6094 Usage/Exposure/Frequency Qualitative Concept (6094)
	UsageExposureFrequencyQualitativeConcept6094 = New("1.2.840.10008.6.1.416", "Usage/Exposure/Frequency Qualitative Concept (6094)", TypeContextGroupName, false)

	// ProcedureNumericProperty6095 Procedure Numeric Property (6095)
	ProcedureNumericProperty6095 = New("1.2.840.10008.6.1.417", "Procedure Numeric Property (6095)", TypeContextGroupName, false)

	// PregnancyStatus6096 Pregnancy Status (6096)
	PregnancyStatus6096 = New("1.2.840.10008.6.1.418", "Pregnancy Status (6096)", TypeContextGroupName, false)

	// SideOfFamily6097 Side of Family (6097)
	SideOfFamily6097 = New("1.2.840.10008.6.1.419", "Side of Family (6097)", TypeContextGroupName, false)

	// ChestComponentCategory6100 Chest Component Category (6100)
	ChestComponentCategory6100 = New("1.2.840.10008.6.1.420", "Chest Component Category (6100)", TypeContextGroupName, false)

	// ChestFindingOrFeature6101 Chest Finding or Feature (6101)
	ChestFindingOrFeature6101 = New("1.2.840.10008.6.1.421", "Chest Finding or Feature (6101)", TypeContextGroupName, false)

	// ChestFindingOrFeatureModifier6102 Chest Finding or Feature Modifier (6102)
	ChestFindingOrFeatureModifier6102 = New("1.2.840.10008.6.1.422", "Chest Finding or Feature Modifier (6102)", TypeContextGroupName, false)

	// AbnormalLinesFindingOrFeature6103 Abnormal Lines Finding or Feature (6103)
	AbnormalLinesFindingOrFeature6103 = New("1.2.840.10008.6.1.423", "Abnormal Lines Finding or Feature (6103)", TypeContextGroupName, false)

	// AbnormalOpacityFindingOrFeature6104 Abnormal Opacity Finding or Feature (6104)
	AbnormalOpacityFindingOrFeature6104 = New("1.2.840.10008.6.1.424", "Abnormal Opacity Finding or Feature (6104)", TypeContextGroupName, false)

	// AbnormalLucencyFindingOrFeature6105 Abnormal Lucency Finding or Feature (6105)
	AbnormalLucencyFindingOrFeature6105 = New("1.2.840.10008.6.1.425", "Abnormal Lucency Finding or Feature (6105)", TypeContextGroupName, false)

	// AbnormalTextureFindingOrFeature6106 Abnormal Texture Finding or Feature (6106)
	AbnormalTextureFindingOrFeature6106 = New("1.2.840.10008.6.1.426", "Abnormal Texture Finding or Feature (6106)", TypeContextGroupName, false)

	// WidthDescriptor6107 Width Descriptor (6107)
	WidthDescriptor6107 = New("1.2.840.10008.6.1.427", "Width Descriptor (6107)", TypeContextGroupName, false)

	// ChestAnatomicStructureAbnormalDistribution6108 Chest Anatomic Structure Abnormal Distribution (6108)
	ChestAnatomicStructureAbnormalDistribution6108 = New("1.2.840.10008.6.1.428", "Chest Anatomic Structure Abnormal Distribution (6108)", TypeContextGroupName, false)

	// RadiographicAnatomyFindingOrFeature6109 Radiographic Anatomy Finding or Feature (6109)
	RadiographicAnatomyFindingOrFeature6109 = New("1.2.840.10008.6.1.429", "Radiographic Anatomy Finding or Feature (6109)", TypeContextGroupName, false)

	// LungAnatomyFindingOrFeature6110 Lung Anatomy Finding or Feature (6110)
	LungAnatomyFindingOrFeature6110 = New("1.2.840.10008.6.1.430", "Lung Anatomy Finding or Feature (6110)", TypeContextGroupName, false)

	// BronchovascularAnatomyFindingOrFeature6111 Bronchovascular Anatomy Finding or Feature (6111)
	BronchovascularAnatomyFindingOrFeature6111 = New("1.2.840.10008.6.1.431", "Bronchovascular Anatomy Finding or Feature (6111)", TypeContextGroupName, false)

	// PleuraAnatomyFindingOrFeature6112 Pleura Anatomy Finding or Feature (6112)
	PleuraAnatomyFindingOrFeature6112 = New("1.2.840.10008.6.1.432", "Pleura Anatomy Finding or Feature (6112)", TypeContextGroupName, false)

	// MediastinumAnatomyFindingOrFeature6113 Mediastinum Anatomy Finding or Feature (6113)
	MediastinumAnatomyFindingOrFeature6113 = New("1.2.840.10008.6.1.433", "Mediastinum Anatomy Finding or Feature (6113)", TypeContextGroupName, false)

	// OsseousAnatomyFindingOrFeature6114 Osseous Anatomy Finding or Feature (6114)
	OsseousAnatomyFindingOrFeature6114 = New("1.2.840.10008.6.1.434", "Osseous Anatomy Finding or Feature (6114)", TypeContextGroupName, false)

	// OsseousAnatomyModifier6115 Osseous Anatomy Modifier (6115)
	OsseousAnatomyModifier6115 = New("1.2.840.10008.6.1.435", "Osseous Anatomy Modifier (6115)", TypeContextGroupName, false)

	// MuscularAnatomy6116 Muscular Anatomy (6116)
	MuscularAnatomy6116 = New("1.2.840.10008.6.1.436", "Muscular Anatomy (6116)", TypeContextGroupName, false)

	// VascularAnatomy6117 Vascular Anatomy (6117)
	VascularAnatomy6117 = New("1.2.840.10008.6.1.437", "Vascular Anatomy (6117)", TypeContextGroupName, false)

	// SizeDescriptor6118 Size Descriptor (6118)
	SizeDescriptor6118 = New("1.2.840.10008.6.1.438", "Size Descriptor (6118)", TypeContextGroupName, false)

	// ChestBorderShape6119 Chest Border Shape (6119)
	ChestBorderShape6119 = New("1.2.840.10008.6.1.439", "Chest Border Shape (6119)", TypeContextGroupName, false)

	// ChestBorderDefinition6120 Chest Border Definition (6120)
	ChestBorderDefinition6120 = New("1.2.840.10008.6.1.440", "Chest Border Definition (6120)", TypeContextGroupName, false)

	// ChestOrientationDescriptor6121 Chest Orientation Descriptor (6121)
	ChestOrientationDescriptor6121 = New("1.2.840.10008.6.1.441", "Chest Orientation Descriptor (6121)", TypeContextGroupName, false)

	// ChestContentDescriptor6122 Chest Content Descriptor (6122)
	ChestContentDescriptor6122 = New("1.2.840.10008.6.1.442", "Chest Content Descriptor (6122)", TypeContextGroupName, false)

	// ChestOpacityDescriptor6123 Chest Opacity Descriptor (6123)
	ChestOpacityDescriptor6123 = New("1.2.840.10008.6.1.443", "Chest Opacity Descriptor (6123)", TypeContextGroupName, false)

	// LocationInChest6124 Location in Chest (6124)
	LocationInChest6124 = New("1.2.840.10008.6.1.444", "Location in Chest (6124)", TypeContextGroupName, false)

	// GeneralChestLocation6125 General Chest Location (6125)
	GeneralChestLocation6125 = New("1.2.840.10008.6.1.445", "General Chest Location (6125)", TypeContextGroupName, false)

	// LocationInLung6126 Location in Lung (6126)
	LocationInLung6126 = New("1.2.840.10008.6.1.446", "Location in Lung (6126)", TypeContextGroupName, false)

	// SegmentLocationInLung6127 Segment Location in Lung (6127)
	SegmentLocationInLung6127 = New("1.2.840.10008.6.1.447", "Segment Location in Lung (6127)", TypeContextGroupName, false)

	// ChestDistributionDescriptor6128 Chest Distribution Descriptor (6128)
	ChestDistributionDescriptor6128 = New("1.2.840.10008.6.1.448", "Chest Distribution Descriptor (6128)", TypeContextGroupName, false)

	// ChestSiteInvolvement6129 Chest Site Involvement (6129)
	ChestSiteInvolvement6129 = New("1.2.840.10008.6.1.449", "Chest Site Involvement (6129)", TypeContextGroupName, false)

	// SeverityDescriptor6130 Severity Descriptor (6130)
	SeverityDescriptor6130 = New("1.2.840.10008.6.1.450", "Severity Descriptor (6130)", TypeContextGroupName, false)

	// ChestTextureDescriptor6131 Chest Texture Descriptor (6131)
	ChestTextureDescriptor6131 = New("1.2.840.10008.6.1.451", "Chest Texture Descriptor (6131)", TypeContextGroupName, false)

	// ChestCalcificationDescriptor6132 Chest Calcification Descriptor (6132)
	ChestCalcificationDescriptor6132 = New("1.2.840.10008.6.1.452", "Chest Calcification Descriptor (6132)", TypeContextGroupName, false)

	// ChestQuantitativeTemporalDifferenceType6133 Chest Quantitative Temporal Difference Type (6133)
	ChestQuantitativeTemporalDifferenceType6133 = New("1.2.840.10008.6.1.453", "Chest Quantitative Temporal Difference Type (6133)", TypeContextGroupName, false)

	// ChestQualitativeTemporalDifferenceType6134 Chest Qualitative Temporal Difference Type (6134)
	ChestQualitativeTemporalDifferenceType6134 = New("1.2.840.10008.6.1.454", "Chest Qualitative Temporal Difference Type (6134)", TypeContextGroupName, false)

	// ImageQualityFinding6135 Image Quality Finding (6135)
	ImageQualityFinding6135 = New("1.2.840.10008.6.1.455", "Image Quality Finding (6135)", TypeContextGroupName, false)

	// ChestTypesOfQualityControlStandard6136 Chest Types of Quality Control Standard (6136)
	ChestTypesOfQualityControlStandard6136 = New("1.2.840.10008.6.1.456", "Chest Types of Quality Control Standard (6136)", TypeContextGroupName, false)

	// CADAnalysisType6137 CAD Analysis Type (6137)
	CADAnalysisType6137 = New("1.2.840.10008.6.1.457", "CAD Analysis Type (6137)", TypeContextGroupName, false)

	// ChestNonLesionObjectType6138 Chest Non-lesion Object Type (6138)
	ChestNonLesionObjectType6138 = New("1.2.840.10008.6.1.458", "Chest Non-lesion Object Type (6138)", TypeContextGroupName, false)

	// NonLesionModifier6139 Non-lesion Modifier (6139)
	NonLesionModifier6139 = New("1.2.840.10008.6.1.459", "Non-lesion Modifier (6139)", TypeContextGroupName, false)

	// CalculationMethod6140 Calculation Method (6140)
	CalculationMethod6140 = New("1.2.840.10008.6.1.460", "Calculation Method (6140)", TypeContextGroupName, false)

	// AttenuationCoefficientMeasurement6141 Attenuation Coefficient Measurement (6141)
	AttenuationCoefficientMeasurement6141 = New("1.2.840.10008.6.1.461", "Attenuation Coefficient Measurement (6141)", TypeContextGroupName, false)

	// CalculatedValue6142 Calculated Value (6142)
	CalculatedValue6142 = New("1.2.840.10008.6.1.462", "Calculated Value (6142)", TypeContextGroupName, false)

	// LesionResponse6143 Lesion Response (6143)
	LesionResponse6143 = New("1.2.840.10008.6.1.463", "Lesion Response (6143)", TypeContextGroupName, false)

	// RECISTDefinedLesionResponse6144 RECIST Defined Lesion Response (6144)
	RECISTDefinedLesionResponse6144 = New("1.2.840.10008.6.1.464", "RECIST Defined Lesion Response (6144)", TypeContextGroupName, false)

	// BaselineCategory6145 Baseline Category (6145)
	BaselineCategory6145 = New("1.2.840.10008.6.1.465", "Baseline Category (6145)", TypeContextGroupName, false)

	// BackgroundEchotexture6151 Background Echotexture (6151)
	BackgroundEchotexture6151 = New("1.2.840.10008.6.1.466", "Background Echotexture (6151)", TypeContextGroupName, false)

	// Orientation6152 Orientation (6152)
	Orientation6152 = New("1.2.840.10008.6.1.467", "Orientation (6152)", TypeContextGroupName, false)

	// LesionBoundary6153 Lesion Boundary (6153)
	LesionBoundary6153 = New("1.2.840.10008.6.1.468", "Lesion Boundary (6153)", TypeContextGroupName, false)

	// EchoPattern6154 Echo Pattern (6154)
	EchoPattern6154 = New("1.2.840.10008.6.1.469", "Echo Pattern (6154)", TypeContextGroupName, false)

	// PosteriorAcousticFeature6155 Posterior Acoustic Feature (6155)
	PosteriorAcousticFeature6155 = New("1.2.840.10008.6.1.470", "Posterior Acoustic Feature (6155)", TypeContextGroupName, false)

	// Vascularity6157 Vascularity (6157)
	Vascularity6157 = New("1.2.840.10008.6.1.471", "Vascularity (6157)", TypeContextGroupName, false)

	// CorrelationToOtherFinding6158 Correlation to Other Finding (6158)
	CorrelationToOtherFinding6158 = New("1.2.840.10008.6.1.472", "Correlation to Other Finding (6158)", TypeContextGroupName, false)

	// MalignancyType6159 Malignancy Type (6159)
	MalignancyType6159 = New("1.2.840.10008.6.1.473", "Malignancy Type (6159)", TypeContextGroupName, false)

	// BreastPrimaryTumorAssessmentFromAJCC6160 Breast Primary Tumor Assessment From AJCC (6160)
	BreastPrimaryTumorAssessmentFromAJCC6160 = New("1.2.840.10008.6.1.474", "Breast Primary Tumor Assessment From AJCC (6160)", TypeContextGroupName, false)

	// PathologicalRegionalLymphNodeAssessmentForBreast6161 Pathological Regional Lymph Node Assessment for Breast (6161)
	PathologicalRegionalLymphNodeAssessmentForBreast6161 = New("1.2.840.10008.6.1.475", "Pathological Regional Lymph Node Assessment for Breast (6161)", TypeContextGroupName, false)

	// AssessmentOfMetastasisForBreast6162 Assessment of Metastasis for Breast (6162)
	AssessmentOfMetastasisForBreast6162 = New("1.2.840.10008.6.1.476", "Assessment of Metastasis for Breast (6162)", TypeContextGroupName, false)

	// MenstrualCyclePhase6163 Menstrual Cycle Phase (6163)
	MenstrualCyclePhase6163 = New("1.2.840.10008.6.1.477", "Menstrual Cycle Phase (6163)", TypeContextGroupName, false)

	// TimeInterval6164 Time Interval (6164)
	TimeInterval6164 = New("1.2.840.10008.6.1.478", "Time Interval (6164)", TypeContextGroupName, false)

	// BreastLinearMeasurement6165 Breast Linear Measurement (6165)
	BreastLinearMeasurement6165 = New("1.2.840.10008.6.1.479", "Breast Linear Measurement (6165)", TypeContextGroupName, false)

	// CADGeometrySecondaryGraphicalRepresentation6166 CAD Geometry Secondary Graphical Representation (6166)
	CADGeometrySecondaryGraphicalRepresentation6166 = New("1.2.840.10008.6.1.480", "CAD Geometry Secondary Graphical Representation (6166)", TypeContextGroupName, false)

	// DiagnosticImagingReportDocumentTitle7000 Diagnostic Imaging Report Document Title (7000)
	DiagnosticImagingReportDocumentTitle7000 = New("1.2.840.10008.6.1.481", "Diagnostic Imaging Report Document Title (7000)", TypeContextGroupName, false)

	// DiagnosticImagingReportHeading7001 Diagnostic Imaging Report Heading (7001)
	DiagnosticImagingReportHeading7001 = New("1.2.840.10008.6.1.482", "Diagnostic Imaging Report Heading (7001)", TypeContextGroupName, false)

	// DiagnosticImagingReportElement7002 Diagnostic Imaging Report Element (7002)
	DiagnosticImagingReportElement7002 = New("1.2.840.10008.6.1.483", "Diagnostic Imaging Report Element (7002)", TypeContextGroupName, false)

	// DiagnosticImagingReportPurposeOfReference7003 Diagnostic Imaging Report Purpose of Reference (7003)
	DiagnosticImagingReportPurposeOfReference7003 = New("1.2.840.10008.6.1.484", "Diagnostic Imaging Report Purpose of Reference (7003)", TypeContextGroupName, false)

	// WaveformPurposeOfReference7004 Waveform Purpose of Reference (7004)
	WaveformPurposeOfReference7004 = New("1.2.840.10008.6.1.485", "Waveform Purpose of Reference (7004)", TypeContextGroupName, false)

	// ContributingEquipmentPurposeOfReference7005 Contributing Equipment Purpose of Reference (7005)
	ContributingEquipmentPurposeOfReference7005 = New("1.2.840.10008.6.1.486", "Contributing Equipment Purpose of Reference (7005)", TypeContextGroupName, false)

	// SRDocumentPurposeOfReference7006 SR Document Purpose of Reference (7006)
	SRDocumentPurposeOfReference7006 = New("1.2.840.10008.6.1.487", "SR Document Purpose of Reference (7006)", TypeContextGroupName, false)

	// SignaturePurpose7007 Signature Purpose (7007)
	SignaturePurpose7007 = New("1.2.840.10008.6.1.488", "Signature Purpose (7007)", TypeContextGroupName, false)

	// MediaImport7008 Media Import (7008)
	MediaImport7008 = New("1.2.840.10008.6.1.489", "Media Import (7008)", TypeContextGroupName, false)

	// KeyObjectSelectionDocumentTitle7010 Key Object Selection Document Title (7010)
	KeyObjectSelectionDocumentTitle7010 = New("1.2.840.10008.6.1.490", "Key Object Selection Document Title (7010)", TypeContextGroupName, false)

	// RejectedForQualityReason7011 Rejected for Quality Reason (7011)
	RejectedForQualityReason7011 = New("1.2.840.10008.6.1.491", "Rejected for Quality Reason (7011)", TypeContextGroupName, false)

	// BestInSet7012 Best in Set (7012)
	BestInSet7012 = New("1.2.840.10008.6.1.492", "Best in Set (7012)", TypeContextGroupName, false)

	// DocumentTitle7020 Document Title (7020)
	DocumentTitle7020 = New("1.2.840.10008.6.1.493", "Document Title (7020)", TypeContextGroupName, false)

	// RCSRegistrationMethodType7100 RCS Registration Method Type (7100)
	RCSRegistrationMethodType7100 = New("1.2.840.10008.6.1.494", "RCS Registration Method Type (7100)", TypeContextGroupName, false)

	// BrainAtlasFiducial7101 Brain Atlas Fiducial (7101)
	BrainAtlasFiducial7101 = New("1.2.840.10008.6.1.495", "Brain Atlas Fiducial (7101)", TypeContextGroupName, false)

	// SegmentationPropertyCategory7150 Segmentation Property Category (7150)
	SegmentationPropertyCategory7150 = New("1.2.840.10008.6.1.496", "Segmentation Property Category (7150)", TypeContextGroupName, false)

	// SegmentationPropertyType7151 Segmentation Property Type (7151)
	SegmentationPropertyType7151 = New("1.2.840.10008.6.1.497", "Segmentation Property Type (7151)", TypeContextGroupName, false)

	// CardiacStructureSegmentationType7152 Cardiac Structure Segmentation Type (7152)
	CardiacStructureSegmentationType7152 = New("1.2.840.10008.6.1.498", "Cardiac Structure Segmentation Type (7152)", TypeContextGroupName, false)

	// CNSSegmentationType7153 CNS Segmentation Type (7153)
	CNSSegmentationType7153 = New("1.2.840.10008.6.1.499", "CNS Segmentation Type (7153)", TypeContextGroupName, false)

	// AbdominalSegmentationType7154 Abdominal Segmentation Type (7154)
	AbdominalSegmentationType7154 = New("1.2.840.10008.6.1.500", "Abdominal Segmentation Type (7154)", TypeContextGroupName, false)

	// ThoracicSegmentationType7155 Thoracic Segmentation Type (7155)
	ThoracicSegmentationType7155 = New("1.2.840.10008.6.1.501", "Thoracic Segmentation Type (7155)", TypeContextGroupName, false)

	// VascularSegmentationType7156 Vascular Segmentation Type (7156)
	VascularSegmentationType7156 = New("1.2.840.10008.6.1.502", "Vascular Segmentation Type (7156)", TypeContextGroupName, false)

	// DeviceSegmentationType7157 Device Segmentation Type (7157)
	DeviceSegmentationType7157 = New("1.2.840.10008.6.1.503", "Device Segmentation Type (7157)", TypeContextGroupName, false)

	// ArtifactSegmentationType7158 Artifact Segmentation Type (7158)
	ArtifactSegmentationType7158 = New("1.2.840.10008.6.1.504", "Artifact Segmentation Type (7158)", TypeContextGroupName, false)

	// LesionSegmentationType7159 Lesion Segmentation Type (7159)
	LesionSegmentationType7159 = New("1.2.840.10008.6.1.505", "Lesion Segmentation Type (7159)", TypeContextGroupName, false)

	// PelvicOrganSegmentationType7160 Pelvic Organ Segmentation Type (7160)
	PelvicOrganSegmentationType7160 = New("1.2.840.10008.6.1.506", "Pelvic Organ Segmentation Type (7160)", TypeContextGroupName, false)

	// PhysiologySegmentationType7161 Physiology Segmentation Type (7161)
	PhysiologySegmentationType7161 = New("1.2.840.10008.6.1.507", "Physiology Segmentation Type (7161)", TypeContextGroupName, false)

	// ReferencedImagePurposeOfReference7201 Referenced Image Purpose of Reference (7201)
	ReferencedImagePurposeOfReference7201 = New("1.2.840.10008.6.1.508", "Referenced Image Purpose of Reference (7201)", TypeContextGroupName, false)

	// SourceImagePurposeOfReference7202 Source Image Purpose of Reference (7202)
	SourceImagePurposeOfReference7202 = New("1.2.840.10008.6.1.509", "Source Image Purpose of Reference (7202)", TypeContextGroupName, false)

	// ImageDerivation7203 Image Derivation (7203)
	ImageDerivation7203 = New("1.2.840.10008.6.1.510", "Image Derivation (7203)", TypeContextGroupName, false)

	// PurposeOfReferenceToAlternateRepresentation7205 Purpose of Reference to Alternate Representation (7205)
	PurposeOfReferenceToAlternateRepresentation7205 = New("1.2.840.10008.6.1.511", "Purpose of Reference to Alternate Representation (7205)", TypeContextGroupName, false)

	// RelatedSeriesPurposeOfReference7210 Related Series Purpose of Reference (7210)
	RelatedSeriesPurposeOfReference7210 = New("1.2.840.10008.6.1.512", "Related Series Purpose of Reference (7210)", TypeContextGroupName, false)

	// MultiFrameSubsetType7250 Multi-Frame Subset Type (7250)
	MultiFrameSubsetType7250 = New("1.2.840.10008.6.1.513", "Multi-Frame Subset Type (7250)", TypeContextGroupName, false)

	// PersonRole7450 Person Role (7450)
	PersonRole7450 = New("1.2.840.10008.6.1.514", "Person Role (7450)", TypeContextGroupName, false)

	// FamilyMember7451 Family Member (7451)
	FamilyMember7451 = New("1.2.840.10008.6.1.515", "Family Member (7451)", TypeContextGroupName, false)

	// OrganizationalRole7452 Organizational Role (7452)
	OrganizationalRole7452 = New("1.2.840.10008.6.1.516", "Organizational Role (7452)", TypeContextGroupName, false)

	// PerformingRole7453 Performing Role (7453)
	PerformingRole7453 = New("1.2.840.10008.6.1.517", "Performing Role (7453)", TypeContextGroupName, false)

	// AnimalTaxonomicRankValue7454 Animal Taxonomic Rank Value (7454)
	AnimalTaxonomicRankValue7454 = New("1.2.840.10008.6.1.518", "Animal Taxonomic Rank Value (7454)", TypeContextGroupName, false)

	// Sex7455 Sex (7455)
	Sex7455 = New("1.2.840.10008.6.1.519", "Sex (7455)", TypeContextGroupName, false)

	// AgeUnit7456 Age Unit (7456)
	AgeUnit7456 = New("1.2.840.10008.6.1.520", "Age Unit (7456)", TypeContextGroupName, false)

	// LinearMeasurementUnit7460 Linear Measurement Unit (7460)
	LinearMeasurementUnit7460 = New("1.2.840.10008.6.1.521", "Linear Measurement Unit (7460)", TypeContextGroupName, false)

	// AreaMeasurementUnit7461 Area Measurement Unit (7461)
	AreaMeasurementUnit7461 = New("1.2.840.10008.6.1.522", "Area Measurement Unit (7461)", TypeContextGroupName, false)

	// VolumeMeasurementUnit7462 Volume Measurement Unit (7462)
	VolumeMeasurementUnit7462 = New("1.2.840.10008.6.1.523", "Volume Measurement Unit (7462)", TypeContextGroupName, false)

	// LinearMeasurement7470 Linear Measurement (7470)
	LinearMeasurement7470 = New("1.2.840.10008.6.1.524", "Linear Measurement (7470)", TypeContextGroupName, false)

	// AreaMeasurement7471 Area Measurement (7471)
	AreaMeasurement7471 = New("1.2.840.10008.6.1.525", "Area Measurement (7471)", TypeContextGroupName, false)

	// VolumeMeasurement7472 Volume Measurement (7472)
	VolumeMeasurement7472 = New("1.2.840.10008.6.1.526", "Volume Measurement (7472)", TypeContextGroupName, false)

	// GeneralAreaCalculationMethod7473 General Area Calculation Method (7473)
	GeneralAreaCalculationMethod7473 = New("1.2.840.10008.6.1.527", "General Area Calculation Method (7473)", TypeContextGroupName, false)

	// GeneralVolumeCalculationMethod7474 General Volume Calculation Method (7474)
	GeneralVolumeCalculationMethod7474 = New("1.2.840.10008.6.1.528", "General Volume Calculation Method (7474)", TypeContextGroupName, false)

	// Breed7480 Breed (7480)
	Breed7480 = New("1.2.840.10008.6.1.529", "Breed (7480)", TypeContextGroupName, false)

	// BreedRegistry7481 Breed Registry (7481)
	BreedRegistry7481 = New("1.2.840.10008.6.1.530", "Breed Registry (7481)", TypeContextGroupName, false)

	// WorkitemDefinition9231 Workitem Definition (9231)
	WorkitemDefinition9231 = New("1.2.840.10008.6.1.531", "Workitem Definition (9231)", TypeContextGroupName, false)

	// NonDICOMOutputTypes9232RETIRED Non-DICOM Output Types (Retired) (9232)
	NonDICOMOutputTypes9232RETIRED = New("1.2.840.10008.6.1.532", "Non-DICOM Output Types (Retired) (9232)", TypeContextGroupName, true)

	// ProcedureDiscontinuationReason9300 Procedure Discontinuation Reason (9300)
	ProcedureDiscontinuationReason9300 = New("1.2.840.10008.6.1.533", "Procedure Discontinuation Reason (9300)", TypeContextGroupName, false)

	// ScopeOfAccumulation10000 Scope of Accumulation (10000)
	ScopeOfAccumulation10000 = New("1.2.840.10008.6.1.534", "Scope of Accumulation (10000)", TypeContextGroupName, false)

	// UIDType10001 UID Type (10001)
	UIDType10001 = New("1.2.840.10008.6.1.535", "UID Type (10001)", TypeContextGroupName, false)

	// IrradiationEventType10002 Irradiation Event Type (10002)
	IrradiationEventType10002 = New("1.2.840.10008.6.1.536", "Irradiation Event Type (10002)", TypeContextGroupName, false)

	// EquipmentPlaneIdentification10003 Equipment Plane Identification (10003)
	EquipmentPlaneIdentification10003 = New("1.2.840.10008.6.1.537", "Equipment Plane Identification (10003)", TypeContextGroupName, false)

	// FluoroMode10004 Fluoro Mode (10004)
	FluoroMode10004 = New("1.2.840.10008.6.1.538", "Fluoro Mode (10004)", TypeContextGroupName, false)

	// XRayFilterMaterial10006 X-Ray Filter Material (10006)
	XRayFilterMaterial10006 = New("1.2.840.10008.6.1.539", "X-Ray Filter Material (10006)", TypeContextGroupName, false)

	// XRayFilterType10007 X-Ray Filter Type (10007)
	XRayFilterType10007 = New("1.2.840.10008.6.1.540", "X-Ray Filter Type (10007)", TypeContextGroupName, false)

	// DoseRelatedDistanceMeasurement10008 Dose Related Distance Measurement (10008)
	DoseRelatedDistanceMeasurement10008 = New("1.2.840.10008.6.1.541", "Dose Related Distance Measurement (10008)", TypeContextGroupName, false)

	// MeasuredCalculated10009 Measured/Calculated (10009)
	MeasuredCalculated10009 = New("1.2.840.10008.6.1.542", "Measured/Calculated (10009)", TypeContextGroupName, false)

	// DoseMeasurementDevice10010 Dose Measurement Device (10010)
	DoseMeasurementDevice10010 = New("1.2.840.10008.6.1.543", "Dose Measurement Device (10010)", TypeContextGroupName, false)

	// EffectiveDoseEvaluationMethod10011 Effective Dose Evaluation Method (10011)
	EffectiveDoseEvaluationMethod10011 = New("1.2.840.10008.6.1.544", "Effective Dose Evaluation Method (10011)", TypeContextGroupName, false)

	// CTAcquisitionType10013 CT Acquisition Type (10013)
	CTAcquisitionType10013 = New("1.2.840.10008.6.1.545", "CT Acquisition Type (10013)", TypeContextGroupName, false)

	// CTIVContrastImagingTechnique10014 CT IV Contrast Imaging Technique (10014)
	CTIVContrastImagingTechnique10014 = New("1.2.840.10008.6.1.546", "CT IV Contrast Imaging Technique (10014)", TypeContextGroupName, false)

	// CTDoseReferenceAuthority10015 CT Dose Reference Authority (10015)
	CTDoseReferenceAuthority10015 = New("1.2.840.10008.6.1.547", "CT Dose Reference Authority (10015)", TypeContextGroupName, false)

	// AnodeTargetMaterial10016 Anode Target Material (10016)
	AnodeTargetMaterial10016 = New("1.2.840.10008.6.1.548", "Anode Target Material (10016)", TypeContextGroupName, false)

	// XRayGrid10017 X-Ray Grid (10017)
	XRayGrid10017 = New("1.2.840.10008.6.1.549", "X-Ray Grid (10017)", TypeContextGroupName, false)

	// UltrasoundProtocolType12001 Ultrasound Protocol Type (12001)
	UltrasoundProtocolType12001 = New("1.2.840.10008.6.1.550", "Ultrasound Protocol Type (12001)", TypeContextGroupName, false)

	// UltrasoundProtocolStageType12002 Ultrasound Protocol Stage Type (12002)
	UltrasoundProtocolStageType12002 = New("1.2.840.10008.6.1.551", "Ultrasound Protocol Stage Type (12002)", TypeContextGroupName, false)

	// OBGYNDate12003 OB-GYN Date (12003)
	OBGYNDate12003 = New("1.2.840.10008.6.1.552", "OB-GYN Date (12003)", TypeContextGroupName, false)

	// FetalBiometryRatio12004 Fetal Biometry Ratio (12004)
	FetalBiometryRatio12004 = New("1.2.840.10008.6.1.553", "Fetal Biometry Ratio (12004)", TypeContextGroupName, false)

	// FetalBiometryMeasurement12005 Fetal Biometry Measurement (12005)
	FetalBiometryMeasurement12005 = New("1.2.840.10008.6.1.554", "Fetal Biometry Measurement (12005)", TypeContextGroupName, false)

	// FetalLongBonesBiometryMeasurement12006 Fetal Long Bones Biometry Measurement (12006)
	FetalLongBonesBiometryMeasurement12006 = New("1.2.840.10008.6.1.555", "Fetal Long Bones Biometry Measurement (12006)", TypeContextGroupName, false)

	// FetalCraniumMeasurement12007 Fetal Cranium Measurement (12007)
	FetalCraniumMeasurement12007 = New("1.2.840.10008.6.1.556", "Fetal Cranium Measurement (12007)", TypeContextGroupName, false)

	// OBGYNAmnioticSacMeasurement12008 OB-GYN Amniotic Sac Measurement (12008)
	OBGYNAmnioticSacMeasurement12008 = New("1.2.840.10008.6.1.557", "OB-GYN Amniotic Sac Measurement (12008)", TypeContextGroupName, false)

	// EarlyGestationBiometryMeasurement12009 Early Gestation Biometry Measurement (12009)
	EarlyGestationBiometryMeasurement12009 = New("1.2.840.10008.6.1.558", "Early Gestation Biometry Measurement (12009)", TypeContextGroupName, false)

	// UltrasoundPelvisAndUterusMeasurement12011 Ultrasound Pelvis and Uterus Measurement (12011)
	UltrasoundPelvisAndUterusMeasurement12011 = New("1.2.840.10008.6.1.559", "Ultrasound Pelvis and Uterus Measurement (12011)", TypeContextGroupName, false)

	// OBEquationTable12012 OB Equation/Table (12012)
	OBEquationTable12012 = New("1.2.840.10008.6.1.560", "OB Equation/Table (12012)", TypeContextGroupName, false)

	// GestationalAgeEquationTable12013 Gestational Age Equation/Table (12013)
	GestationalAgeEquationTable12013 = New("1.2.840.10008.6.1.561", "Gestational Age Equation/Table (12013)", TypeContextGroupName, false)

	// OBFetalBodyWeightEquationTable12014 OB Fetal Body Weight Equation/Table (12014)
	OBFetalBodyWeightEquationTable12014 = New("1.2.840.10008.6.1.562", "OB Fetal Body Weight Equation/Table (12014)", TypeContextGroupName, false)

	// FetalGrowthEquationTable12015 Fetal Growth Equation/Table (12015)
	FetalGrowthEquationTable12015 = New("1.2.840.10008.6.1.563", "Fetal Growth Equation/Table (12015)", TypeContextGroupName, false)

	// EstimatedFetalWeightPercentileEquationTable12016 Estimated Fetal Weight Percentile Equation/Table (12016)
	EstimatedFetalWeightPercentileEquationTable12016 = New("1.2.840.10008.6.1.564", "Estimated Fetal Weight Percentile Equation/Table (12016)", TypeContextGroupName, false)

	// GrowthDistributionRank12017 Growth Distribution Rank (12017)
	GrowthDistributionRank12017 = New("1.2.840.10008.6.1.565", "Growth Distribution Rank (12017)", TypeContextGroupName, false)

	// OBGYNSummary12018 OB-GYN Summary (12018)
	OBGYNSummary12018 = New("1.2.840.10008.6.1.566", "OB-GYN Summary (12018)", TypeContextGroupName, false)

	// OBGYNFetusSummary12019 OB-GYN Fetus Summary (12019)
	OBGYNFetusSummary12019 = New("1.2.840.10008.6.1.567", "OB-GYN Fetus Summary (12019)", TypeContextGroupName, false)

	// VascularSummary12101 Vascular Summary (12101)
	VascularSummary12101 = New("1.2.840.10008.6.1.568", "Vascular Summary (12101)", TypeContextGroupName, false)

	// TemporalPeriodRelatingToProcedureOrTherapy12102 Temporal Period Relating to Procedure or Therapy (12102)
	TemporalPeriodRelatingToProcedureOrTherapy12102 = New("1.2.840.10008.6.1.569", "Temporal Period Relating to Procedure or Therapy (12102)", TypeContextGroupName, false)

	// VascularUltrasoundAnatomicLocation12103 Vascular Ultrasound Anatomic Location (12103)
	VascularUltrasoundAnatomicLocation12103 = New("1.2.840.10008.6.1.570", "Vascular Ultrasound Anatomic Location (12103)", TypeContextGroupName, false)

	// ExtracranialArtery12104 Extracranial Artery (12104)
	ExtracranialArtery12104 = New("1.2.840.10008.6.1.571", "Extracranial Artery (12104)", TypeContextGroupName, false)

	// IntracranialCerebralVessel12105 Intracranial Cerebral Vessel (12105)
	IntracranialCerebralVessel12105 = New("1.2.840.10008.6.1.572", "Intracranial Cerebral Vessel (12105)", TypeContextGroupName, false)

	// IntracranialCerebralVesselUnilateral12106 Intracranial Cerebral Vessel (Unilateral) (12106)
	IntracranialCerebralVesselUnilateral12106 = New("1.2.840.10008.6.1.573", "Intracranial Cerebral Vessel (Unilateral) (12106)", TypeContextGroupName, false)

	// UpperExtremityArtery12107 Upper Extremity Artery (12107)
	UpperExtremityArtery12107 = New("1.2.840.10008.6.1.574", "Upper Extremity Artery (12107)", TypeContextGroupName, false)

	// UpperExtremityVein12108 Upper Extremity Vein (12108)
	UpperExtremityVein12108 = New("1.2.840.10008.6.1.575", "Upper Extremity Vein (12108)", TypeContextGroupName, false)

	// LowerExtremityArtery12109 Lower Extremity Artery (12109)
	LowerExtremityArtery12109 = New("1.2.840.10008.6.1.576", "Lower Extremity Artery (12109)", TypeContextGroupName, false)

	// LowerExtremityVein12110 Lower Extremity Vein (12110)
	LowerExtremityVein12110 = New("1.2.840.10008.6.1.577", "Lower Extremity Vein (12110)", TypeContextGroupName, false)

	// AbdominopelvicArteryPaired12111 Abdominopelvic Artery (Paired) (12111)
	AbdominopelvicArteryPaired12111 = New("1.2.840.10008.6.1.578", "Abdominopelvic Artery (Paired) (12111)", TypeContextGroupName, false)

	// AbdominopelvicArteryUnpaired12112 Abdominopelvic Artery (Unpaired) (12112)
	AbdominopelvicArteryUnpaired12112 = New("1.2.840.10008.6.1.579", "Abdominopelvic Artery (Unpaired) (12112)", TypeContextGroupName, false)

	// AbdominopelvicVeinPaired12113 Abdominopelvic Vein (Paired) (12113)
	AbdominopelvicVeinPaired12113 = New("1.2.840.10008.6.1.580", "Abdominopelvic Vein (Paired) (12113)", TypeContextGroupName, false)

	// AbdominopelvicVeinUnpaired12114 Abdominopelvic Vein (Unpaired) (12114)
	AbdominopelvicVeinUnpaired12114 = New("1.2.840.10008.6.1.581", "Abdominopelvic Vein (Unpaired) (12114)", TypeContextGroupName, false)

	// RenalVessel12115 Renal Vessel (12115)
	RenalVessel12115 = New("1.2.840.10008.6.1.582", "Renal Vessel (12115)", TypeContextGroupName, false)

	// VesselSegmentModifier12116 Vessel Segment Modifier (12116)
	VesselSegmentModifier12116 = New("1.2.840.10008.6.1.583", "Vessel Segment Modifier (12116)", TypeContextGroupName, false)

	// VesselBranchModifier12117 Vessel Branch Modifier (12117)
	VesselBranchModifier12117 = New("1.2.840.10008.6.1.584", "Vessel Branch Modifier (12117)", TypeContextGroupName, false)

	// VascularUltrasoundProperty12119 Vascular Ultrasound Property (12119)
	VascularUltrasoundProperty12119 = New("1.2.840.10008.6.1.585", "Vascular Ultrasound Property (12119)", TypeContextGroupName, false)

	// UltrasoundBloodVelocityMeasurement12120 Ultrasound Blood Velocity Measurement (12120)
	UltrasoundBloodVelocityMeasurement12120 = New("1.2.840.10008.6.1.586", "Ultrasound Blood Velocity Measurement (12120)", TypeContextGroupName, false)

	// VascularIndexRatio12121 Vascular Index/Ratio (12121)
	VascularIndexRatio12121 = New("1.2.840.10008.6.1.587", "Vascular Index/Ratio (12121)", TypeContextGroupName, false)

	// OtherVascularProperty12122 Other Vascular Property (12122)
	OtherVascularProperty12122 = New("1.2.840.10008.6.1.588", "Other Vascular Property (12122)", TypeContextGroupName, false)

	// CarotidRatio12123 Carotid Ratio (12123)
	CarotidRatio12123 = New("1.2.840.10008.6.1.589", "Carotid Ratio (12123)", TypeContextGroupName, false)

	// RenalRatio12124 Renal Ratio (12124)
	RenalRatio12124 = New("1.2.840.10008.6.1.590", "Renal Ratio (12124)", TypeContextGroupName, false)

	// PelvicVasculatureAnatomicalLocation12140 Pelvic Vasculature Anatomical Location (12140)
	PelvicVasculatureAnatomicalLocation12140 = New("1.2.840.10008.6.1.591", "Pelvic Vasculature Anatomical Location (12140)", TypeContextGroupName, false)

	// FetalVasculatureAnatomicalLocation12141 Fetal Vasculature Anatomical Location (12141)
	FetalVasculatureAnatomicalLocation12141 = New("1.2.840.10008.6.1.592", "Fetal Vasculature Anatomical Location (12141)", TypeContextGroupName, false)

	// EchocardiographyLeftVentricleMeasurement12200 Echocardiography Left Ventricle Measurement (12200)
	EchocardiographyLeftVentricleMeasurement12200 = New("1.2.840.10008.6.1.593", "Echocardiography Left Ventricle Measurement (12200)", TypeContextGroupName, false)

	// LeftVentricleLinearMeasurement12201 Left Ventricle Linear Measurement (12201)
	LeftVentricleLinearMeasurement12201 = New("1.2.840.10008.6.1.594", "Left Ventricle Linear Measurement (12201)", TypeContextGroupName, false)

	// LeftVentricleVolumeMeasurement12202 Left Ventricle Volume Measurement (12202)
	LeftVentricleVolumeMeasurement12202 = New("1.2.840.10008.6.1.595", "Left Ventricle Volume Measurement (12202)", TypeContextGroupName, false)

	// LeftVentricleOtherMeasurement12203 Left Ventricle Other Measurement (12203)
	LeftVentricleOtherMeasurement12203 = New("1.2.840.10008.6.1.596", "Left Ventricle Other Measurement (12203)", TypeContextGroupName, false)

	// EchocardiographyRightVentricleMeasurement12204 Echocardiography Right Ventricle Measurement (12204)
	EchocardiographyRightVentricleMeasurement12204 = New("1.2.840.10008.6.1.597", "Echocardiography Right Ventricle Measurement (12204)", TypeContextGroupName, false)

	// EchocardiographyLeftAtriumMeasurement12205 Echocardiography Left Atrium Measurement (12205)
	EchocardiographyLeftAtriumMeasurement12205 = New("1.2.840.10008.6.1.598", "Echocardiography Left Atrium Measurement (12205)", TypeContextGroupName, false)

	// EchocardiographyRightAtriumMeasurement12206 Echocardiography Right Atrium Measurement (12206)
	EchocardiographyRightAtriumMeasurement12206 = New("1.2.840.10008.6.1.599", "Echocardiography Right Atrium Measurement (12206)", TypeContextGroupName, false)

	// EchocardiographyMitralValveMeasurement12207 Echocardiography Mitral Valve Measurement (12207)
	EchocardiographyMitralValveMeasurement12207 = New("1.2.840.10008.6.1.600", "Echocardiography Mitral Valve Measurement (12207)", TypeContextGroupName, false)

	// EchocardiographyTricuspidValveMeasurement12208 Echocardiography Tricuspid Valve Measurement (12208)
	EchocardiographyTricuspidValveMeasurement12208 = New("1.2.840.10008.6.1.601", "Echocardiography Tricuspid Valve Measurement (12208)", TypeContextGroupName, false)

	// EchocardiographyPulmonicValveMeasurement12209 Echocardiography Pulmonic Valve Measurement (12209)
	EchocardiographyPulmonicValveMeasurement12209 = New("1.2.840.10008.6.1.602", "Echocardiography Pulmonic Valve Measurement (12209)", TypeContextGroupName, false)

	// EchocardiographyPulmonaryArteryMeasurement12210 Echocardiography Pulmonary Artery Measurement (12210)
	EchocardiographyPulmonaryArteryMeasurement12210 = New("1.2.840.10008.6.1.603", "Echocardiography Pulmonary Artery Measurement (12210)", TypeContextGroupName, false)

	// EchocardiographyAorticValveMeasurement12211 Echocardiography Aortic Valve Measurement (12211)
	EchocardiographyAorticValveMeasurement12211 = New("1.2.840.10008.6.1.604", "Echocardiography Aortic Valve Measurement (12211)", TypeContextGroupName, false)

	// EchocardiographyAortaMeasurement12212 Echocardiography Aorta Measurement (12212)
	EchocardiographyAortaMeasurement12212 = New("1.2.840.10008.6.1.605", "Echocardiography Aorta Measurement (12212)", TypeContextGroupName, false)

	// EchocardiographyPulmonaryVeinMeasurement12214 Echocardiography Pulmonary Vein Measurement (12214)
	EchocardiographyPulmonaryVeinMeasurement12214 = New("1.2.840.10008.6.1.606", "Echocardiography Pulmonary Vein Measurement (12214)", TypeContextGroupName, false)

	// EchocardiographyVenaCavaMeasurement12215 Echocardiography Vena Cava Measurement (12215)
	EchocardiographyVenaCavaMeasurement12215 = New("1.2.840.10008.6.1.607", "Echocardiography Vena Cava Measurement (12215)", TypeContextGroupName, false)

	// EchocardiographyHepaticVeinMeasurement12216 Echocardiography Hepatic Vein Measurement (12216)
	EchocardiographyHepaticVeinMeasurement12216 = New("1.2.840.10008.6.1.608", "Echocardiography Hepatic Vein Measurement (12216)", TypeContextGroupName, false)

	// EchocardiographyCardiacShuntMeasurement12217 Echocardiography Cardiac Shunt Measurement (12217)
	EchocardiographyCardiacShuntMeasurement12217 = New("1.2.840.10008.6.1.609", "Echocardiography Cardiac Shunt Measurement (12217)", TypeContextGroupName, false)

	// EchocardiographyCongenitalAnomalyMeasurement12218 Echocardiography Congenital Anomaly Measurement (12218)
	EchocardiographyCongenitalAnomalyMeasurement12218 = New("1.2.840.10008.6.1.610", "Echocardiography Congenital Anomaly Measurement (12218)", TypeContextGroupName, false)

	// PulmonaryVeinModifier12219 Pulmonary Vein Modifier (12219)
	PulmonaryVeinModifier12219 = New("1.2.840.10008.6.1.611", "Pulmonary Vein Modifier (12219)", TypeContextGroupName, false)

	// EchocardiographyCommonMeasurement12220 Echocardiography Common Measurement (12220)
	EchocardiographyCommonMeasurement12220 = New("1.2.840.10008.6.1.612", "Echocardiography Common Measurement (12220)", TypeContextGroupName, false)

	// FlowDirection12221 Flow Direction (12221)
	FlowDirection12221 = New("1.2.840.10008.6.1.613", "Flow Direction (12221)", TypeContextGroupName, false)

	// OrificeFlowProperty12222 Orifice Flow Property (12222)
	OrificeFlowProperty12222 = New("1.2.840.10008.6.1.614", "Orifice Flow Property (12222)", TypeContextGroupName, false)

	// EchocardiographyStrokeVolumeOrigin12223 Echocardiography Stroke Volume Origin (12223)
	EchocardiographyStrokeVolumeOrigin12223 = New("1.2.840.10008.6.1.615", "Echocardiography Stroke Volume Origin (12223)", TypeContextGroupName, false)

	// UltrasoundImageMode12224 Ultrasound Image Mode (12224)
	UltrasoundImageMode12224 = New("1.2.840.10008.6.1.616", "Ultrasound Image Mode (12224)", TypeContextGroupName, false)

	// EchocardiographyImageView12226 Echocardiography Image View (12226)
	EchocardiographyImageView12226 = New("1.2.840.10008.6.1.617", "Echocardiography Image View (12226)", TypeContextGroupName, false)

	// EchocardiographyMeasurementMethod12227 Echocardiography Measurement Method (12227)
	EchocardiographyMeasurementMethod12227 = New("1.2.840.10008.6.1.618", "Echocardiography Measurement Method (12227)", TypeContextGroupName, false)

	// EchocardiographyVolumeMethod12228 Echocardiography Volume Method (12228)
	EchocardiographyVolumeMethod12228 = New("1.2.840.10008.6.1.619", "Echocardiography Volume Method (12228)", TypeContextGroupName, false)

	// EchocardiographyAreaMethod12229 Echocardiography Area Method (12229)
	EchocardiographyAreaMethod12229 = New("1.2.840.10008.6.1.620", "Echocardiography Area Method (12229)", TypeContextGroupName, false)

	// GradientMethod12230 Gradient Method (12230)
	GradientMethod12230 = New("1.2.840.10008.6.1.621", "Gradient Method (12230)", TypeContextGroupName, false)

	// VolumeFlowMethod12231 Volume Flow Method (12231)
	VolumeFlowMethod12231 = New("1.2.840.10008.6.1.622", "Volume Flow Method (12231)", TypeContextGroupName, false)

	// MyocardiumMassMethod12232 Myocardium Mass Method (12232)
	MyocardiumMassMethod12232 = New("1.2.840.10008.6.1.623", "Myocardium Mass Method (12232)", TypeContextGroupName, false)

	// CardiacPhase12233 Cardiac Phase (12233)
	CardiacPhase12233 = New("1.2.840.10008.6.1.624", "Cardiac Phase (12233)", TypeContextGroupName, false)

	// RespirationState12234 Respiration State (12234)
	RespirationState12234 = New("1.2.840.10008.6.1.625", "Respiration State (12234)", TypeContextGroupName, false)

	// MitralValveAnatomicSite12235 Mitral Valve Anatomic Site (12235)
	MitralValveAnatomicSite12235 = New("1.2.840.10008.6.1.626", "Mitral Valve Anatomic Site (12235)", TypeContextGroupName, false)

	// EchocardiographyAnatomicSite12236 Echocardiography Anatomic Site (12236)
	EchocardiographyAnatomicSite12236 = New("1.2.840.10008.6.1.627", "Echocardiography Anatomic Site (12236)", TypeContextGroupName, false)

	// EchocardiographyAnatomicSiteModifier12237 Echocardiography Anatomic Site Modifier (12237)
	EchocardiographyAnatomicSiteModifier12237 = New("1.2.840.10008.6.1.628", "Echocardiography Anatomic Site Modifier (12237)", TypeContextGroupName, false)

	// WallMotionScoringScheme12238 Wall Motion Scoring Scheme (12238)
	WallMotionScoringScheme12238 = New("1.2.840.10008.6.1.629", "Wall Motion Scoring Scheme (12238)", TypeContextGroupName, false)

	// CardiacOutputProperty12239 Cardiac Output Property (12239)
	CardiacOutputProperty12239 = New("1.2.840.10008.6.1.630", "Cardiac Output Property (12239)", TypeContextGroupName, false)

	// LeftVentricleAreaMeasurement12240 Left Ventricle Area Measurement (12240)
	LeftVentricleAreaMeasurement12240 = New("1.2.840.10008.6.1.631", "Left Ventricle Area Measurement (12240)", TypeContextGroupName, false)

	// TricuspidValveFindingSite12241 Tricuspid Valve Finding Site (12241)
	TricuspidValveFindingSite12241 = New("1.2.840.10008.6.1.632", "Tricuspid Valve Finding Site (12241)", TypeContextGroupName, false)

	// AorticValveFindingSite12242 Aortic Valve Finding Site (12242)
	AorticValveFindingSite12242 = New("1.2.840.10008.6.1.633", "Aortic Valve Finding Site (12242)", TypeContextGroupName, false)

	// LeftVentricleFindingSite12243 Left Ventricle Finding Site (12243)
	LeftVentricleFindingSite12243 = New("1.2.840.10008.6.1.634", "Left Ventricle Finding Site (12243)", TypeContextGroupName, false)

	// CongenitalFindingSite12244 Congenital Finding Site (12244)
	CongenitalFindingSite12244 = New("1.2.840.10008.6.1.635", "Congenital Finding Site (12244)", TypeContextGroupName, false)

	// SurfaceProcessingAlgorithmFamily7162 Surface Processing Algorithm Family (7162)
	SurfaceProcessingAlgorithmFamily7162 = New("1.2.840.10008.6.1.636", "Surface Processing Algorithm Family (7162)", TypeContextGroupName, false)

	// StressTestProcedurePhase3207 Stress Test Procedure Phase (3207)
	StressTestProcedurePhase3207 = New("1.2.840.10008.6.1.637", "Stress Test Procedure Phase (3207)", TypeContextGroupName, false)

	// Stage3778 Stage (3778)
	Stage3778 = New("1.2.840.10008.6.1.638", "Stage (3778)", TypeContextGroupName, false)

	// SMLSizeDescriptor252 S-M-L Size Descriptor (252)
	SMLSizeDescriptor252 = New("1.2.840.10008.6.1.735", "S-M-L Size Descriptor (252)", TypeContextGroupName, false)

	// MajorCoronaryArtery3016 Major Coronary Artery (3016)
	MajorCoronaryArtery3016 = New("1.2.840.10008.6.1.736", "Major Coronary Artery (3016)", TypeContextGroupName, false)

	// RadioactivityUnit3083 Radioactivity Unit (3083)
	RadioactivityUnit3083 = New("1.2.840.10008.6.1.737", "Radioactivity Unit (3083)", TypeContextGroupName, false)

	// RestStressState3102 Rest/Stress State (3102)
	RestStressState3102 = New("1.2.840.10008.6.1.738", "Rest/Stress State (3102)", TypeContextGroupName, false)

	// PETCardiologyProtocol3106 PET Cardiology Protocol (3106)
	PETCardiologyProtocol3106 = New("1.2.840.10008.6.1.739", "PET Cardiology Protocol (3106)", TypeContextGroupName, false)

	// PETCardiologyRadiopharmaceutical3107 PET Cardiology Radiopharmaceutical (3107)
	PETCardiologyRadiopharmaceutical3107 = New("1.2.840.10008.6.1.740", "PET Cardiology Radiopharmaceutical (3107)", TypeContextGroupName, false)

	// NMPETProcedure3108 NM/PET Procedure (3108)
	NMPETProcedure3108 = New("1.2.840.10008.6.1.741", "NM/PET Procedure (3108)", TypeContextGroupName, false)

	// NuclearCardiologyProtocol3110 Nuclear Cardiology Protocol (3110)
	NuclearCardiologyProtocol3110 = New("1.2.840.10008.6.1.742", "Nuclear Cardiology Protocol (3110)", TypeContextGroupName, false)

	// NuclearCardiologyRadiopharmaceutical3111 Nuclear Cardiology Radiopharmaceutical (3111)
	NuclearCardiologyRadiopharmaceutical3111 = New("1.2.840.10008.6.1.743", "Nuclear Cardiology Radiopharmaceutical (3111)", TypeContextGroupName, false)

	// AttenuationCorrection3112 Attenuation Correction (3112)
	AttenuationCorrection3112 = New("1.2.840.10008.6.1.744", "Attenuation Correction (3112)", TypeContextGroupName, false)

	// PerfusionDefectType3113 Perfusion Defect Type (3113)
	PerfusionDefectType3113 = New("1.2.840.10008.6.1.745", "Perfusion Defect Type (3113)", TypeContextGroupName, false)

	// StudyQuality3114 Study Quality (3114)
	StudyQuality3114 = New("1.2.840.10008.6.1.746", "Study Quality (3114)", TypeContextGroupName, false)

	// StressImagingQualityIssue3115 Stress Imaging Quality Issue (3115)
	StressImagingQualityIssue3115 = New("1.2.840.10008.6.1.747", "Stress Imaging Quality Issue (3115)", TypeContextGroupName, false)

	// NMExtracardiacFinding3116 NM Extracardiac Finding (3116)
	NMExtracardiacFinding3116 = New("1.2.840.10008.6.1.748", "NM Extracardiac Finding (3116)", TypeContextGroupName, false)

	// AttenuationCorrectionMethod3117 Attenuation Correction Method (3117)
	AttenuationCorrectionMethod3117 = New("1.2.840.10008.6.1.749", "Attenuation Correction Method (3117)", TypeContextGroupName, false)

	// LevelOfRisk3118 Level of Risk (3118)
	LevelOfRisk3118 = New("1.2.840.10008.6.1.750", "Level of Risk (3118)", TypeContextGroupName, false)

	// LVFunction3119 LV Function (3119)
	LVFunction3119 = New("1.2.840.10008.6.1.751", "LV Function (3119)", TypeContextGroupName, false)

	// PerfusionFinding3120 Perfusion Finding (3120)
	PerfusionFinding3120 = New("1.2.840.10008.6.1.752", "Perfusion Finding (3120)", TypeContextGroupName, false)

	// PerfusionMorphology3121 Perfusion Morphology (3121)
	PerfusionMorphology3121 = New("1.2.840.10008.6.1.753", "Perfusion Morphology (3121)", TypeContextGroupName, false)

	// VentricularEnlargement3122 Ventricular Enlargement (3122)
	VentricularEnlargement3122 = New("1.2.840.10008.6.1.754", "Ventricular Enlargement (3122)", TypeContextGroupName, false)

	// StressTestProcedure3200 Stress Test Procedure (3200)
	StressTestProcedure3200 = New("1.2.840.10008.6.1.755", "Stress Test Procedure (3200)", TypeContextGroupName, false)

	// IndicationsForStressTest3201 Indications for Stress Test (3201)
	IndicationsForStressTest3201 = New("1.2.840.10008.6.1.756", "Indications for Stress Test (3201)", TypeContextGroupName, false)

	// ChestPain3202 Chest Pain (3202)
	ChestPain3202 = New("1.2.840.10008.6.1.757", "Chest Pain (3202)", TypeContextGroupName, false)

	// ExerciserDevice3203 Exerciser Device (3203)
	ExerciserDevice3203 = New("1.2.840.10008.6.1.758", "Exerciser Device (3203)", TypeContextGroupName, false)

	// StressAgent3204 Stress Agent (3204)
	StressAgent3204 = New("1.2.840.10008.6.1.759", "Stress Agent (3204)", TypeContextGroupName, false)

	// IndicationsForPharmacologicalStressTest3205 Indications for Pharmacological Stress Test (3205)
	IndicationsForPharmacologicalStressTest3205 = New("1.2.840.10008.6.1.760", "Indications for Pharmacological Stress Test (3205)", TypeContextGroupName, false)

	// NonInvasiveCardiacImagingProcedure3206 Non-invasive Cardiac Imaging Procedure (3206)
	NonInvasiveCardiacImagingProcedure3206 = New("1.2.840.10008.6.1.761", "Non-invasive Cardiac Imaging Procedure (3206)", TypeContextGroupName, false)

	// ExerciseECGSummaryCode3208 Exercise ECG Summary Code (3208)
	ExerciseECGSummaryCode3208 = New("1.2.840.10008.6.1.763", "Exercise ECG Summary Code (3208)", TypeContextGroupName, false)

	// StressImagingSummaryCode3209 Stress Imaging Summary Code (3209)
	StressImagingSummaryCode3209 = New("1.2.840.10008.6.1.764", "Stress Imaging Summary Code (3209)", TypeContextGroupName, false)

	// SpeedOfResponse3210 Speed of Response (3210)
	SpeedOfResponse3210 = New("1.2.840.10008.6.1.765", "Speed of Response (3210)", TypeContextGroupName, false)

	// BPResponse3211 BP Response (3211)
	BPResponse3211 = New("1.2.840.10008.6.1.766", "BP Response (3211)", TypeContextGroupName, false)

	// TreadmillSpeed3212 Treadmill Speed (3212)
	TreadmillSpeed3212 = New("1.2.840.10008.6.1.767", "Treadmill Speed (3212)", TypeContextGroupName, false)

	// StressHemodynamicFinding3213 Stress Hemodynamic Finding (3213)
	StressHemodynamicFinding3213 = New("1.2.840.10008.6.1.768", "Stress Hemodynamic Finding (3213)", TypeContextGroupName, false)

	// PerfusionFindingMethod3215 Perfusion Finding Method (3215)
	PerfusionFindingMethod3215 = New("1.2.840.10008.6.1.769", "Perfusion Finding Method (3215)", TypeContextGroupName, false)

	// ComparisonFinding3217 Comparison Finding (3217)
	ComparisonFinding3217 = New("1.2.840.10008.6.1.770", "Comparison Finding (3217)", TypeContextGroupName, false)

	// StressSymptom3220 Stress Symptom (3220)
	StressSymptom3220 = New("1.2.840.10008.6.1.771", "Stress Symptom (3220)", TypeContextGroupName, false)

	// StressTestTerminationReason3221 Stress Test Termination Reason (3221)
	StressTestTerminationReason3221 = New("1.2.840.10008.6.1.772", "Stress Test Termination Reason (3221)", TypeContextGroupName, false)

	// QTcMeasurement3227 QTc Measurement (3227)
	QTcMeasurement3227 = New("1.2.840.10008.6.1.773", "QTc Measurement (3227)", TypeContextGroupName, false)

	// ECGTimingMeasurement3228 ECG Timing Measurement (3228)
	ECGTimingMeasurement3228 = New("1.2.840.10008.6.1.774", "ECG Timing Measurement (3228)", TypeContextGroupName, false)

	// ECGAxisMeasurement3229 ECG Axis Measurement (3229)
	ECGAxisMeasurement3229 = New("1.2.840.10008.6.1.775", "ECG Axis Measurement (3229)", TypeContextGroupName, false)

	// ECGFinding3230 ECG Finding (3230)
	ECGFinding3230 = New("1.2.840.10008.6.1.776", "ECG Finding (3230)", TypeContextGroupName, false)

	// STSegmentFinding3231 ST Segment Finding (3231)
	STSegmentFinding3231 = New("1.2.840.10008.6.1.777", "ST Segment Finding (3231)", TypeContextGroupName, false)

	// STSegmentLocation3232 ST Segment Location (3232)
	STSegmentLocation3232 = New("1.2.840.10008.6.1.778", "ST Segment Location (3232)", TypeContextGroupName, false)

	// STSegmentMorphology3233 ST Segment Morphology (3233)
	STSegmentMorphology3233 = New("1.2.840.10008.6.1.779", "ST Segment Morphology (3233)", TypeContextGroupName, false)

	// EctopicBeatMorphology3234 Ectopic Beat Morphology (3234)
	EctopicBeatMorphology3234 = New("1.2.840.10008.6.1.780", "Ectopic Beat Morphology (3234)", TypeContextGroupName, false)

	// PerfusionComparisonFinding3235 Perfusion Comparison Finding (3235)
	PerfusionComparisonFinding3235 = New("1.2.840.10008.6.1.781", "Perfusion Comparison Finding (3235)", TypeContextGroupName, false)

	// ToleranceComparisonFinding3236 Tolerance Comparison Finding (3236)
	ToleranceComparisonFinding3236 = New("1.2.840.10008.6.1.782", "Tolerance Comparison Finding (3236)", TypeContextGroupName, false)

	// WallMotionComparisonFinding3237 Wall Motion Comparison Finding (3237)
	WallMotionComparisonFinding3237 = New("1.2.840.10008.6.1.783", "Wall Motion Comparison Finding (3237)", TypeContextGroupName, false)

	// StressScoringScale3238 Stress Scoring Scale (3238)
	StressScoringScale3238 = New("1.2.840.10008.6.1.784", "Stress Scoring Scale (3238)", TypeContextGroupName, false)

	// PerceivedExertionScale3239 Perceived Exertion Scale (3239)
	PerceivedExertionScale3239 = New("1.2.840.10008.6.1.785", "Perceived Exertion Scale (3239)", TypeContextGroupName, false)

	// VentricleIdentification3463 Ventricle Identification (3463)
	VentricleIdentification3463 = New("1.2.840.10008.6.1.786", "Ventricle Identification (3463)", TypeContextGroupName, false)

	// ColonOverallAssessment6200 Colon Overall Assessment (6200)
	ColonOverallAssessment6200 = New("1.2.840.10008.6.1.787", "Colon Overall Assessment (6200)", TypeContextGroupName, false)

	// ColonFindingOrFeature6201 Colon Finding or Feature (6201)
	ColonFindingOrFeature6201 = New("1.2.840.10008.6.1.788", "Colon Finding or Feature (6201)", TypeContextGroupName, false)

	// ColonFindingOrFeatureModifier6202 Colon Finding or Feature Modifier (6202)
	ColonFindingOrFeatureModifier6202 = New("1.2.840.10008.6.1.789", "Colon Finding or Feature Modifier (6202)", TypeContextGroupName, false)

	// ColonNonLesionObjectType6203 Colon Non-lesion Object Type (6203)
	ColonNonLesionObjectType6203 = New("1.2.840.10008.6.1.790", "Colon Non-lesion Object Type (6203)", TypeContextGroupName, false)

	// AnatomicNonColonFinding6204 Anatomic Non-colon Finding (6204)
	AnatomicNonColonFinding6204 = New("1.2.840.10008.6.1.791", "Anatomic Non-colon Finding (6204)", TypeContextGroupName, false)

	// ClockfaceLocationForColon6205 Clockface Location for Colon (6205)
	ClockfaceLocationForColon6205 = New("1.2.840.10008.6.1.792", "Clockface Location for Colon (6205)", TypeContextGroupName, false)

	// RecumbentPatientOrientationForColon6206 Recumbent Patient Orientation for Colon (6206)
	RecumbentPatientOrientationForColon6206 = New("1.2.840.10008.6.1.793", "Recumbent Patient Orientation for Colon (6206)", TypeContextGroupName, false)

	// ColonQuantitativeTemporalDifferenceType6207 Colon Quantitative Temporal Difference Type (6207)
	ColonQuantitativeTemporalDifferenceType6207 = New("1.2.840.10008.6.1.794", "Colon Quantitative Temporal Difference Type (6207)", TypeContextGroupName, false)

	// ColonTypesOfQualityControlStandard6208 Colon Types of Quality Control Standard (6208)
	ColonTypesOfQualityControlStandard6208 = New("1.2.840.10008.6.1.795", "Colon Types of Quality Control Standard (6208)", TypeContextGroupName, false)

	// ColonMorphologyDescriptor6209 Colon Morphology Descriptor (6209)
	ColonMorphologyDescriptor6209 = New("1.2.840.10008.6.1.796", "Colon Morphology Descriptor (6209)", TypeContextGroupName, false)

	// LocationInIntestinalTract6210 Location in Intestinal Tract (6210)
	LocationInIntestinalTract6210 = New("1.2.840.10008.6.1.797", "Location in Intestinal Tract (6210)", TypeContextGroupName, false)

	// ColonCADMaterialDescription6211 Colon CAD Material Description (6211)
	ColonCADMaterialDescription6211 = New("1.2.840.10008.6.1.798", "Colon CAD Material Description (6211)", TypeContextGroupName, false)

	// CalculatedValueForColonFinding6212 Calculated Value for Colon Finding (6212)
	CalculatedValueForColonFinding6212 = New("1.2.840.10008.6.1.799", "Calculated Value for Colon Finding (6212)", TypeContextGroupName, false)

	// OphthalmicHorizontalDirection4214 Ophthalmic Horizontal Direction (4214)
	OphthalmicHorizontalDirection4214 = New("1.2.840.10008.6.1.800", "Ophthalmic Horizontal Direction (4214)", TypeContextGroupName, false)

	// OphthalmicVerticalDirection4215 Ophthalmic Vertical Direction (4215)
	OphthalmicVerticalDirection4215 = New("1.2.840.10008.6.1.801", "Ophthalmic Vertical Direction (4215)", TypeContextGroupName, false)

	// OphthalmicVisualAcuityType4216 Ophthalmic Visual Acuity Type (4216)
	OphthalmicVisualAcuityType4216 = New("1.2.840.10008.6.1.802", "Ophthalmic Visual Acuity Type (4216)", TypeContextGroupName, false)

	// ArterialPulseWaveform3004 Arterial Pulse Waveform (3004)
	ArterialPulseWaveform3004 = New("1.2.840.10008.6.1.803", "Arterial Pulse Waveform (3004)", TypeContextGroupName, false)

	// RespirationWaveform3005 Respiration Waveform (3005)
	RespirationWaveform3005 = New("1.2.840.10008.6.1.804", "Respiration Waveform (3005)", TypeContextGroupName, false)

	// UltrasoundContrastBolusAgent12030 Ultrasound Contrast/Bolus Agent (12030)
	UltrasoundContrastBolusAgent12030 = New("1.2.840.10008.6.1.805", "Ultrasound Contrast/Bolus Agent (12030)", TypeContextGroupName, false)

	// ProtocolIntervalEvent12031 Protocol Interval Event (12031)
	ProtocolIntervalEvent12031 = New("1.2.840.10008.6.1.806", "Protocol Interval Event (12031)", TypeContextGroupName, false)

	// TransducerScanPattern12032 Transducer Scan Pattern (12032)
	TransducerScanPattern12032 = New("1.2.840.10008.6.1.807", "Transducer Scan Pattern (12032)", TypeContextGroupName, false)

	// UltrasoundTransducerGeometry12033 Ultrasound Transducer Geometry (12033)
	UltrasoundTransducerGeometry12033 = New("1.2.840.10008.6.1.808", "Ultrasound Transducer Geometry (12033)", TypeContextGroupName, false)

	// UltrasoundTransducerBeamSteering12034 Ultrasound Transducer Beam Steering (12034)
	UltrasoundTransducerBeamSteering12034 = New("1.2.840.10008.6.1.809", "Ultrasound Transducer Beam Steering (12034)", TypeContextGroupName, false)

	// UltrasoundTransducerApplication12035 Ultrasound Transducer Application (12035)
	UltrasoundTransducerApplication12035 = New("1.2.840.10008.6.1.810", "Ultrasound Transducer Application (12035)", TypeContextGroupName, false)

	// InstanceAvailabilityStatus50 Instance Availability Status (50)
	InstanceAvailabilityStatus50 = New("1.2.840.10008.6.1.811", "Instance Availability Status (50)", TypeContextGroupName, false)

	// ModalityPPSDiscontinuationReason9301 Modality PPS Discontinuation Reason (9301)
	ModalityPPSDiscontinuationReason9301 = New("1.2.840.10008.6.1.812", "Modality PPS Discontinuation Reason (9301)", TypeContextGroupName, false)

	// MediaImportPPSDiscontinuationReason9302 Media Import PPS Discontinuation Reason (9302)
	MediaImportPPSDiscontinuationReason9302 = New("1.2.840.10008.6.1.813", "Media Import PPS Discontinuation Reason (9302)", TypeContextGroupName, false)

	// DXAnatomyImagedForAnimal7482 DX Anatomy Imaged for Animal (7482)
	DXAnatomyImagedForAnimal7482 = New("1.2.840.10008.6.1.814", "DX Anatomy Imaged for Animal (7482)", TypeContextGroupName, false)

	// CommonAnatomicRegionsForAnimal7483 Common Anatomic Regions for Animal (7483)
	CommonAnatomicRegionsForAnimal7483 = New("1.2.840.10008.6.1.815", "Common Anatomic Regions for Animal (7483)", TypeContextGroupName, false)

	// DXViewForAnimal7484 DX View for Animal (7484)
	DXViewForAnimal7484 = New("1.2.840.10008.6.1.816", "DX View for Animal (7484)", TypeContextGroupName, false)

	// InstitutionalDepartmentUnitService7030 Institutional Department/Unit/Service (7030)
	InstitutionalDepartmentUnitService7030 = New("1.2.840.10008.6.1.817", "Institutional Department/Unit/Service (7030)", TypeContextGroupName, false)

	// PurposeOfReferenceToPredecessorReport7009 Purpose of Reference to Predecessor Report (7009)
	PurposeOfReferenceToPredecessorReport7009 = New("1.2.840.10008.6.1.818", "Purpose of Reference to Predecessor Report (7009)", TypeContextGroupName, false)

	// VisualFixationQualityDuringAcquisition4220 Visual Fixation Quality During Acquisition (4220)
	VisualFixationQualityDuringAcquisition4220 = New("1.2.840.10008.6.1.819", "Visual Fixation Quality During Acquisition (4220)", TypeContextGroupName, false)

	// VisualFixationQualityProblem4221 Visual Fixation Quality Problem (4221)
	VisualFixationQualityProblem4221 = New("1.2.840.10008.6.1.820", "Visual Fixation Quality Problem (4221)", TypeContextGroupName, false)

	// OphthalmicMacularGridProblem4222 Ophthalmic Macular Grid Problem (4222)
	OphthalmicMacularGridProblem4222 = New("1.2.840.10008.6.1.821", "Ophthalmic Macular Grid Problem (4222)", TypeContextGroupName, false)

	// Organization5002 Organization (5002)
	Organization5002 = New("1.2.840.10008.6.1.822", "Organization (5002)", TypeContextGroupName, false)

	// MixedBreed7486 Mixed Breed (7486)
	MixedBreed7486 = New("1.2.840.10008.6.1.823", "Mixed Breed (7486)", TypeContextGroupName, false)

	// BroselowLutenPediatricSizeCategory7040 Broselow-Luten Pediatric Size Category (7040)
	BroselowLutenPediatricSizeCategory7040 = New("1.2.840.10008.6.1.824", "Broselow-Luten Pediatric Size Category (7040)", TypeContextGroupName, false)

	// CMDCTECCCalciumScoringPatientSizeCategory7042 CMDCTECC Calcium Scoring Patient Size Category (7042)
	CMDCTECCCalciumScoringPatientSizeCategory7042 = New("1.2.840.10008.6.1.825", "CMDCTECC Calcium Scoring Patient Size Category (7042)", TypeContextGroupName, false)

	// CardiacUltrasoundReportTitle12245 Cardiac Ultrasound Report Title (12245)
	CardiacUltrasoundReportTitle12245 = New("1.2.840.10008.6.1.826", "Cardiac Ultrasound Report Title (12245)", TypeContextGroupName, false)

	// CardiacUltrasoundIndicationForStudy12246 Cardiac Ultrasound Indication for Study (12246)
	CardiacUltrasoundIndicationForStudy12246 = New("1.2.840.10008.6.1.827", "Cardiac Ultrasound Indication for Study (12246)", TypeContextGroupName, false)

	// PediatricFetalAndCongenitalCardiacSurgicalIntervention12247 Pediatric, Fetal and Congenital Cardiac Surgical Intervention (12247)
	PediatricFetalAndCongenitalCardiacSurgicalIntervention12247 = New("1.2.840.10008.6.1.828", "Pediatric, Fetal and Congenital Cardiac Surgical Intervention (12247)", TypeContextGroupName, false)

	// CardiacUltrasoundSummaryCode12248 Cardiac Ultrasound Summary Code (12248)
	CardiacUltrasoundSummaryCode12248 = New("1.2.840.10008.6.1.829", "Cardiac Ultrasound Summary Code (12248)", TypeContextGroupName, false)

	// CardiacUltrasoundFetalSummaryCode12249 Cardiac Ultrasound Fetal Summary Code (12249)
	CardiacUltrasoundFetalSummaryCode12249 = New("1.2.840.10008.6.1.830", "Cardiac Ultrasound Fetal Summary Code (12249)", TypeContextGroupName, false)

	// CardiacUltrasoundCommonLinearMeasurement12250 Cardiac Ultrasound Common Linear Measurement (12250)
	CardiacUltrasoundCommonLinearMeasurement12250 = New("1.2.840.10008.6.1.831", "Cardiac Ultrasound Common Linear Measurement (12250)", TypeContextGroupName, false)

	// CardiacUltrasoundLinearValveMeasurement12251 Cardiac Ultrasound Linear Valve Measurement (12251)
	CardiacUltrasoundLinearValveMeasurement12251 = New("1.2.840.10008.6.1.832", "Cardiac Ultrasound Linear Valve Measurement (12251)", TypeContextGroupName, false)

	// CardiacUltrasoundCardiacFunction12252 Cardiac Ultrasound Cardiac Function (12252)
	CardiacUltrasoundCardiacFunction12252 = New("1.2.840.10008.6.1.833", "Cardiac Ultrasound Cardiac Function (12252)", TypeContextGroupName, false)

	// CardiacUltrasoundAreaMeasurement12253 Cardiac Ultrasound Area Measurement (12253)
	CardiacUltrasoundAreaMeasurement12253 = New("1.2.840.10008.6.1.834", "Cardiac Ultrasound Area Measurement (12253)", TypeContextGroupName, false)

	// CardiacUltrasoundHemodynamicMeasurement12254 Cardiac Ultrasound Hemodynamic Measurement (12254)
	CardiacUltrasoundHemodynamicMeasurement12254 = New("1.2.840.10008.6.1.835", "Cardiac Ultrasound Hemodynamic Measurement (12254)", TypeContextGroupName, false)

	// CardiacUltrasoundMyocardiumMeasurement12255 Cardiac Ultrasound Myocardium Measurement (12255)
	CardiacUltrasoundMyocardiumMeasurement12255 = New("1.2.840.10008.6.1.836", "Cardiac Ultrasound Myocardium Measurement (12255)", TypeContextGroupName, false)

	// CardiacUltrasoundLeftVentricleMeasurement12257 Cardiac Ultrasound Left Ventricle Measurement (12257)
	CardiacUltrasoundLeftVentricleMeasurement12257 = New("1.2.840.10008.6.1.838", "Cardiac Ultrasound Left Ventricle Measurement (12257)", TypeContextGroupName, false)

	// CardiacUltrasoundRightVentricleMeasurement12258 Cardiac Ultrasound Right Ventricle Measurement (12258)
	CardiacUltrasoundRightVentricleMeasurement12258 = New("1.2.840.10008.6.1.839", "Cardiac Ultrasound Right Ventricle Measurement (12258)", TypeContextGroupName, false)

	// CardiacUltrasoundVentriclesMeasurement12259 Cardiac Ultrasound Ventricles Measurement (12259)
	CardiacUltrasoundVentriclesMeasurement12259 = New("1.2.840.10008.6.1.840", "Cardiac Ultrasound Ventricles Measurement (12259)", TypeContextGroupName, false)

	// CardiacUltrasoundPulmonaryArteryMeasurement12260 Cardiac Ultrasound Pulmonary Artery Measurement (12260)
	CardiacUltrasoundPulmonaryArteryMeasurement12260 = New("1.2.840.10008.6.1.841", "Cardiac Ultrasound Pulmonary Artery Measurement (12260)", TypeContextGroupName, false)

	// CardiacUltrasoundPulmonaryVein12261 Cardiac Ultrasound Pulmonary Vein (12261)
	CardiacUltrasoundPulmonaryVein12261 = New("1.2.840.10008.6.1.842", "Cardiac Ultrasound Pulmonary Vein (12261)", TypeContextGroupName, false)

	// CardiacUltrasoundPulmonaryValveMeasurement12262 Cardiac Ultrasound Pulmonary Valve Measurement (12262)
	CardiacUltrasoundPulmonaryValveMeasurement12262 = New("1.2.840.10008.6.1.843", "Cardiac Ultrasound Pulmonary Valve Measurement (12262)", TypeContextGroupName, false)

	// CardiacUltrasoundVenousReturnPulmonaryMeasurement12263 Cardiac Ultrasound Venous Return Pulmonary Measurement (12263)
	CardiacUltrasoundVenousReturnPulmonaryMeasurement12263 = New("1.2.840.10008.6.1.844", "Cardiac Ultrasound Venous Return Pulmonary Measurement (12263)", TypeContextGroupName, false)

	// CardiacUltrasoundVenousReturnSystemicMeasurement12264 Cardiac Ultrasound Venous Return Systemic Measurement (12264)
	CardiacUltrasoundVenousReturnSystemicMeasurement12264 = New("1.2.840.10008.6.1.845", "Cardiac Ultrasound Venous Return Systemic Measurement (12264)", TypeContextGroupName, false)

	// CardiacUltrasoundAtriaAndAtrialSeptumMeasurement12265 Cardiac Ultrasound Atria and Atrial Septum Measurement (12265)
	CardiacUltrasoundAtriaAndAtrialSeptumMeasurement12265 = New("1.2.840.10008.6.1.846", "Cardiac Ultrasound Atria and Atrial Septum Measurement (12265)", TypeContextGroupName, false)

	// CardiacUltrasoundMitralValveMeasurement12266 Cardiac Ultrasound Mitral Valve Measurement (12266)
	CardiacUltrasoundMitralValveMeasurement12266 = New("1.2.840.10008.6.1.847", "Cardiac Ultrasound Mitral Valve Measurement (12266)", TypeContextGroupName, false)

	// CardiacUltrasoundTricuspidValveMeasurement12267 Cardiac Ultrasound Tricuspid Valve Measurement (12267)
	CardiacUltrasoundTricuspidValveMeasurement12267 = New("1.2.840.10008.6.1.848", "Cardiac Ultrasound Tricuspid Valve Measurement (12267)", TypeContextGroupName, false)

	// CardiacUltrasoundAtrioventricularValveMeasurement12268 Cardiac Ultrasound Atrioventricular Valve Measurement (12268)
	CardiacUltrasoundAtrioventricularValveMeasurement12268 = New("1.2.840.10008.6.1.849", "Cardiac Ultrasound Atrioventricular Valve Measurement (12268)", TypeContextGroupName, false)

	// CardiacUltrasoundInterventricularSeptumMeasurement12269 Cardiac Ultrasound Interventricular Septum Measurement (12269)
	CardiacUltrasoundInterventricularSeptumMeasurement12269 = New("1.2.840.10008.6.1.850", "Cardiac Ultrasound Interventricular Septum Measurement (12269)", TypeContextGroupName, false)

	// CardiacUltrasoundAorticValveMeasurement12270 Cardiac Ultrasound Aortic Valve Measurement (12270)
	CardiacUltrasoundAorticValveMeasurement12270 = New("1.2.840.10008.6.1.851", "Cardiac Ultrasound Aortic Valve Measurement (12270)", TypeContextGroupName, false)

	// CardiacUltrasoundOutflowTractMeasurement12271 Cardiac Ultrasound Outflow Tract Measurement (12271)
	CardiacUltrasoundOutflowTractMeasurement12271 = New("1.2.840.10008.6.1.852", "Cardiac Ultrasound Outflow Tract Measurement (12271)", TypeContextGroupName, false)

	// CardiacUltrasoundSemilunarValveAnnulateAndSinusMeasurement12272 Cardiac Ultrasound Semilunar Valve, Annulate and Sinus Measurement (12272)
	CardiacUltrasoundSemilunarValveAnnulateAndSinusMeasurement12272 = New("1.2.840.10008.6.1.853", "Cardiac Ultrasound Semilunar Valve, Annulate and Sinus Measurement (12272)", TypeContextGroupName, false)

	// CardiacUltrasoundAorticSinotubularJunctionMeasurement12273 Cardiac Ultrasound Aortic Sinotubular Junction Measurement (12273)
	CardiacUltrasoundAorticSinotubularJunctionMeasurement12273 = New("1.2.840.10008.6.1.854", "Cardiac Ultrasound Aortic Sinotubular Junction Measurement (12273)", TypeContextGroupName, false)

	// CardiacUltrasoundAortaMeasurement12274 Cardiac Ultrasound Aorta Measurement (12274)
	CardiacUltrasoundAortaMeasurement12274 = New("1.2.840.10008.6.1.855", "Cardiac Ultrasound Aorta Measurement (12274)", TypeContextGroupName, false)

	// CardiacUltrasoundCoronaryArteryMeasurement12275 Cardiac Ultrasound Coronary Artery Measurement (12275)
	CardiacUltrasoundCoronaryArteryMeasurement12275 = New("1.2.840.10008.6.1.856", "Cardiac Ultrasound Coronary Artery Measurement (12275)", TypeContextGroupName, false)

	// CardiacUltrasoundAortoPulmonaryConnectionMeasurement12276 Cardiac Ultrasound Aorto Pulmonary Connection Measurement (12276)
	CardiacUltrasoundAortoPulmonaryConnectionMeasurement12276 = New("1.2.840.10008.6.1.857", "Cardiac Ultrasound Aorto Pulmonary Connection Measurement (12276)", TypeContextGroupName, false)

	// CardiacUltrasoundPericardiumAndPleuraMeasurement12277 Cardiac Ultrasound Pericardium and Pleura Measurement (12277)
	CardiacUltrasoundPericardiumAndPleuraMeasurement12277 = New("1.2.840.10008.6.1.858", "Cardiac Ultrasound Pericardium and Pleura Measurement (12277)", TypeContextGroupName, false)

	// CardiacUltrasoundFetalGeneralMeasurement12279 Cardiac Ultrasound Fetal General Measurement (12279)
	CardiacUltrasoundFetalGeneralMeasurement12279 = New("1.2.840.10008.6.1.859", "Cardiac Ultrasound Fetal General Measurement (12279)", TypeContextGroupName, false)

	// CardiacUltrasoundTargetSite12280 Cardiac Ultrasound Target Site (12280)
	CardiacUltrasoundTargetSite12280 = New("1.2.840.10008.6.1.860", "Cardiac Ultrasound Target Site (12280)", TypeContextGroupName, false)

	// CardiacUltrasoundTargetSiteModifier12281 Cardiac Ultrasound Target Site Modifier (12281)
	CardiacUltrasoundTargetSiteModifier12281 = New("1.2.840.10008.6.1.861", "Cardiac Ultrasound Target Site Modifier (12281)", TypeContextGroupName, false)

	// CardiacUltrasoundVenousReturnSystemicFindingSite12282 Cardiac Ultrasound Venous Return Systemic Finding Site (12282)
	CardiacUltrasoundVenousReturnSystemicFindingSite12282 = New("1.2.840.10008.6.1.862", "Cardiac Ultrasound Venous Return Systemic Finding Site (12282)", TypeContextGroupName, false)

	// CardiacUltrasoundVenousReturnPulmonaryFindingSite12283 Cardiac Ultrasound Venous Return Pulmonary Finding Site (12283)
	CardiacUltrasoundVenousReturnPulmonaryFindingSite12283 = New("1.2.840.10008.6.1.863", "Cardiac Ultrasound Venous Return Pulmonary Finding Site (12283)", TypeContextGroupName, false)

	// CardiacUltrasoundAtriaAndAtrialSeptumFindingSite12284 Cardiac Ultrasound Atria and Atrial Septum Finding Site (12284)
	CardiacUltrasoundAtriaAndAtrialSeptumFindingSite12284 = New("1.2.840.10008.6.1.864", "Cardiac Ultrasound Atria and Atrial Septum Finding Site (12284)", TypeContextGroupName, false)

	// CardiacUltrasoundAtrioventricularValveFindingSite12285 Cardiac Ultrasound Atrioventricular Valve Finding Site (12285)
	CardiacUltrasoundAtrioventricularValveFindingSite12285 = New("1.2.840.10008.6.1.865", "Cardiac Ultrasound Atrioventricular Valve Finding Site (12285)", TypeContextGroupName, false)

	// CardiacUltrasoundInterventricularSeptumFindingSite12286 Cardiac Ultrasound Interventricular Septum Finding Site (12286)
	CardiacUltrasoundInterventricularSeptumFindingSite12286 = New("1.2.840.10008.6.1.866", "Cardiac Ultrasound Interventricular Septum Finding Site (12286)", TypeContextGroupName, false)

	// CardiacUltrasoundVentricleFindingSite12287 Cardiac Ultrasound Ventricle Finding Site (12287)
	CardiacUltrasoundVentricleFindingSite12287 = New("1.2.840.10008.6.1.867", "Cardiac Ultrasound Ventricle Finding Site (12287)", TypeContextGroupName, false)

	// CardiacUltrasoundOutflowTractFindingSite12288 Cardiac Ultrasound Outflow Tract Finding Site (12288)
	CardiacUltrasoundOutflowTractFindingSite12288 = New("1.2.840.10008.6.1.868", "Cardiac Ultrasound Outflow Tract Finding Site (12288)", TypeContextGroupName, false)

	// CardiacUltrasoundSemilunarValveAnnulusAndSinusFindingSite12289 Cardiac Ultrasound Semilunar Valve, Annulus and Sinus Finding Site (12289)
	CardiacUltrasoundSemilunarValveAnnulusAndSinusFindingSite12289 = New("1.2.840.10008.6.1.869", "Cardiac Ultrasound Semilunar Valve, Annulus and Sinus Finding Site (12289)", TypeContextGroupName, false)

	// CardiacUltrasoundPulmonaryArteryFindingSite12290 Cardiac Ultrasound Pulmonary Artery Finding Site (12290)
	CardiacUltrasoundPulmonaryArteryFindingSite12290 = New("1.2.840.10008.6.1.870", "Cardiac Ultrasound Pulmonary Artery Finding Site (12290)", TypeContextGroupName, false)

	// CardiacUltrasoundAortaFindingSite12291 Cardiac Ultrasound Aorta Finding Site (12291)
	CardiacUltrasoundAortaFindingSite12291 = New("1.2.840.10008.6.1.871", "Cardiac Ultrasound Aorta Finding Site (12291)", TypeContextGroupName, false)

	// CardiacUltrasoundCoronaryArteryFindingSite12292 Cardiac Ultrasound Coronary Artery Finding Site (12292)
	CardiacUltrasoundCoronaryArteryFindingSite12292 = New("1.2.840.10008.6.1.872", "Cardiac Ultrasound Coronary Artery Finding Site (12292)", TypeContextGroupName, false)

	// CardiacUltrasoundAortopulmonaryConnectionFindingSite12293 Cardiac Ultrasound Aortopulmonary Connection Finding Site (12293)
	CardiacUltrasoundAortopulmonaryConnectionFindingSite12293 = New("1.2.840.10008.6.1.873", "Cardiac Ultrasound Aortopulmonary Connection Finding Site (12293)", TypeContextGroupName, false)

	// CardiacUltrasoundPericardiumAndPleuraFindingSite12294 Cardiac Ultrasound Pericardium and Pleura Finding Site (12294)
	CardiacUltrasoundPericardiumAndPleuraFindingSite12294 = New("1.2.840.10008.6.1.874", "Cardiac Ultrasound Pericardium and Pleura Finding Site (12294)", TypeContextGroupName, false)

	// OphthalmicUltrasoundAxialMeasurementsType4230 Ophthalmic Ultrasound Axial Measurements Type (4230)
	OphthalmicUltrasoundAxialMeasurementsType4230 = New("1.2.840.10008.6.1.876", "Ophthalmic Ultrasound Axial Measurements Type (4230)", TypeContextGroupName, false)

	// LensStatus4231 Lens Status (4231)
	LensStatus4231 = New("1.2.840.10008.6.1.877", "Lens Status (4231)", TypeContextGroupName, false)

	// VitreousStatus4232 Vitreous Status (4232)
	VitreousStatus4232 = New("1.2.840.10008.6.1.878", "Vitreous Status (4232)", TypeContextGroupName, false)

	// OphthalmicAxialLengthMeasurementsSegmentName4233 Ophthalmic Axial Length Measurements Segment Name (4233)
	OphthalmicAxialLengthMeasurementsSegmentName4233 = New("1.2.840.10008.6.1.879", "Ophthalmic Axial Length Measurements Segment Name (4233)", TypeContextGroupName, false)

	// RefractiveSurgeryType4234 Refractive Surgery Type (4234)
	RefractiveSurgeryType4234 = New("1.2.840.10008.6.1.880", "Refractive Surgery Type (4234)", TypeContextGroupName, false)

	// KeratometryDescriptor4235 Keratometry Descriptor (4235)
	KeratometryDescriptor4235 = New("1.2.840.10008.6.1.881", "Keratometry Descriptor (4235)", TypeContextGroupName, false)

	// IOLCalculationFormula4236 IOL Calculation Formula (4236)
	IOLCalculationFormula4236 = New("1.2.840.10008.6.1.882", "IOL Calculation Formula (4236)", TypeContextGroupName, false)

	// LensConstantType4237 Lens Constant Type (4237)
	LensConstantType4237 = New("1.2.840.10008.6.1.883", "Lens Constant Type (4237)", TypeContextGroupName, false)

	// RefractiveErrorType4238 Refractive Error Type (4238)
	RefractiveErrorType4238 = New("1.2.840.10008.6.1.884", "Refractive Error Type (4238)", TypeContextGroupName, false)

	// AnteriorChamberDepthDefinition4239 Anterior Chamber Depth Definition (4239)
	AnteriorChamberDepthDefinition4239 = New("1.2.840.10008.6.1.885", "Anterior Chamber Depth Definition (4239)", TypeContextGroupName, false)

	// OphthalmicMeasurementOrCalculationDataSource4240 Ophthalmic Measurement or Calculation Data Source (4240)
	OphthalmicMeasurementOrCalculationDataSource4240 = New("1.2.840.10008.6.1.886", "Ophthalmic Measurement or Calculation Data Source (4240)", TypeContextGroupName, false)

	// OphthalmicAxialLengthSelectionMethod4241 Ophthalmic Axial Length Selection Method (4241)
	OphthalmicAxialLengthSelectionMethod4241 = New("1.2.840.10008.6.1.887", "Ophthalmic Axial Length Selection Method (4241)", TypeContextGroupName, false)

	// OphthalmicQualityMetricType4243 Ophthalmic Quality Metric Type (4243)
	OphthalmicQualityMetricType4243 = New("1.2.840.10008.6.1.889", "Ophthalmic Quality Metric Type (4243)", TypeContextGroupName, false)

	// OphthalmicAgentConcentrationUnit4244 Ophthalmic Agent Concentration Unit (4244)
	OphthalmicAgentConcentrationUnit4244 = New("1.2.840.10008.6.1.890", "Ophthalmic Agent Concentration Unit (4244)", TypeContextGroupName, false)

	// FunctionalConditionPresentDuringAcquisition91 Functional Condition Present During Acquisition (91)
	FunctionalConditionPresentDuringAcquisition91 = New("1.2.840.10008.6.1.891", "Functional Condition Present During Acquisition (91)", TypeContextGroupName, false)

	// JointPositionDuringAcquisition92 Joint Position During Acquisition (92)
	JointPositionDuringAcquisition92 = New("1.2.840.10008.6.1.892", "Joint Position During Acquisition (92)", TypeContextGroupName, false)

	// JointPositioningMethod93 Joint Positioning Method (93)
	JointPositioningMethod93 = New("1.2.840.10008.6.1.893", "Joint Positioning Method (93)", TypeContextGroupName, false)

	// PhysicalForceAppliedDuringAcquisition94 Physical Force Applied During Acquisition (94)
	PhysicalForceAppliedDuringAcquisition94 = New("1.2.840.10008.6.1.894", "Physical Force Applied During Acquisition (94)", TypeContextGroupName, false)

	// ECGControlNumericVariable3690 ECG Control Numeric Variable (3690)
	ECGControlNumericVariable3690 = New("1.2.840.10008.6.1.895", "ECG Control Numeric Variable (3690)", TypeContextGroupName, false)

	// ECGControlTextVariable3691 ECG Control Text Variable (3691)
	ECGControlTextVariable3691 = New("1.2.840.10008.6.1.896", "ECG Control Text Variable (3691)", TypeContextGroupName, false)

	// WholeSlideMicroscopyImageReferencedImagePurposeOfReference8120 Whole Slide Microscopy Image Referenced Image Purpose of Reference (8120)
	WholeSlideMicroscopyImageReferencedImagePurposeOfReference8120 = New("1.2.840.10008.6.1.897", "Whole Slide Microscopy Image Referenced Image Purpose of Reference (8120)", TypeContextGroupName, false)

	// MicroscopyLensType8121 Microscopy Lens Type (8121)
	MicroscopyLensType8121 = New("1.2.840.10008.6.1.898", "Microscopy Lens Type (8121)", TypeContextGroupName, false)

	// MicroscopyIlluminatorAndSensorColor8122 Microscopy Illuminator and Sensor Color (8122)
	MicroscopyIlluminatorAndSensorColor8122 = New("1.2.840.10008.6.1.899", "Microscopy Illuminator and Sensor Color (8122)", TypeContextGroupName, false)

	// MicroscopyIlluminationMethod8123 Microscopy Illumination Method (8123)
	MicroscopyIlluminationMethod8123 = New("1.2.840.10008.6.1.900", "Microscopy Illumination Method (8123)", TypeContextGroupName, false)

	// MicroscopyFilter8124 Microscopy Filter (8124)
	MicroscopyFilter8124 = New("1.2.840.10008.6.1.901", "Microscopy Filter (8124)", TypeContextGroupName, false)

	// MicroscopyIlluminatorType8125 Microscopy Illuminator Type (8125)
	MicroscopyIlluminatorType8125 = New("1.2.840.10008.6.1.902", "Microscopy Illuminator Type (8125)", TypeContextGroupName, false)

	// AuditEventID400 Audit Event ID (400)
	AuditEventID400 = New("1.2.840.10008.6.1.903", "Audit Event ID (400)", TypeContextGroupName, false)

	// AuditEventTypeCode401 Audit Event Type Code (401)
	AuditEventTypeCode401 = New("1.2.840.10008.6.1.904", "Audit Event Type Code (401)", TypeContextGroupName, false)

	// AuditActiveParticipantRoleIDCode402 Audit Active Participant Role ID Code (402)
	AuditActiveParticipantRoleIDCode402 = New("1.2.840.10008.6.1.905", "Audit Active Participant Role ID Code (402)", TypeContextGroupName, false)

	// SecurityAlertTypeCode403 Security Alert Type Code (403)
	SecurityAlertTypeCode403 = New("1.2.840.10008.6.1.906", "Security Alert Type Code (403)", TypeContextGroupName, false)

	// AuditParticipantObjectIDTypeCode404 Audit Participant Object ID Type Code (404)
	AuditParticipantObjectIDTypeCode404 = New("1.2.840.10008.6.1.907", "Audit Participant Object ID Type Code (404)", TypeContextGroupName, false)

	// MediaTypeCode405 Media Type Code (405)
	MediaTypeCode405 = New("1.2.840.10008.6.1.908", "Media Type Code (405)", TypeContextGroupName, false)

	// VisualFieldStaticPerimetryTestPattern4250 Visual Field Static Perimetry Test Pattern (4250)
	VisualFieldStaticPerimetryTestPattern4250 = New("1.2.840.10008.6.1.909", "Visual Field Static Perimetry Test Pattern (4250)", TypeContextGroupName, false)

	// VisualFieldStaticPerimetryTestStrategy4251 Visual Field Static Perimetry Test Strategy (4251)
	VisualFieldStaticPerimetryTestStrategy4251 = New("1.2.840.10008.6.1.910", "Visual Field Static Perimetry Test Strategy (4251)", TypeContextGroupName, false)

	// VisualFieldStaticPerimetryScreeningTestMode4252 Visual Field Static Perimetry Screening Test Mode (4252)
	VisualFieldStaticPerimetryScreeningTestMode4252 = New("1.2.840.10008.6.1.911", "Visual Field Static Perimetry Screening Test Mode (4252)", TypeContextGroupName, false)

	// VisualFieldStaticPerimetryFixationStrategy4253 Visual Field Static Perimetry Fixation Strategy (4253)
	VisualFieldStaticPerimetryFixationStrategy4253 = New("1.2.840.10008.6.1.912", "Visual Field Static Perimetry Fixation Strategy (4253)", TypeContextGroupName, false)

	// VisualFieldStaticPerimetryTestAnalysisResult4254 Visual Field Static Perimetry Test Analysis Result (4254)
	VisualFieldStaticPerimetryTestAnalysisResult4254 = New("1.2.840.10008.6.1.913", "Visual Field Static Perimetry Test Analysis Result (4254)", TypeContextGroupName, false)

	// VisualFieldIlluminationColor4255 Visual Field Illumination Color (4255)
	VisualFieldIlluminationColor4255 = New("1.2.840.10008.6.1.914", "Visual Field Illumination Color (4255)", TypeContextGroupName, false)

	// VisualFieldProcedureModifier4256 Visual Field Procedure Modifier (4256)
	VisualFieldProcedureModifier4256 = New("1.2.840.10008.6.1.915", "Visual Field Procedure Modifier (4256)", TypeContextGroupName, false)

	// VisualFieldGlobalIndexName4257 Visual Field Global Index Name (4257)
	VisualFieldGlobalIndexName4257 = New("1.2.840.10008.6.1.916", "Visual Field Global Index Name (4257)", TypeContextGroupName, false)

	// AbstractMultiDimensionalImageModelComponentSemantic7180 Abstract Multi-dimensional Image Model Component Semantic (7180)
	AbstractMultiDimensionalImageModelComponentSemantic7180 = New("1.2.840.10008.6.1.917", "Abstract Multi-dimensional Image Model Component Semantic (7180)", TypeContextGroupName, false)

	// AbstractMultiDimensionalImageModelComponentUnit7181 Abstract Multi-dimensional Image Model Component Unit (7181)
	AbstractMultiDimensionalImageModelComponentUnit7181 = New("1.2.840.10008.6.1.918", "Abstract Multi-dimensional Image Model Component Unit (7181)", TypeContextGroupName, false)

	// AbstractMultiDimensionalImageModelDimensionSemantic7182 Abstract Multi-dimensional Image Model Dimension Semantic (7182)
	AbstractMultiDimensionalImageModelDimensionSemantic7182 = New("1.2.840.10008.6.1.919", "Abstract Multi-dimensional Image Model Dimension Semantic (7182)", TypeContextGroupName, false)

	// AbstractMultiDimensionalImageModelDimensionUnit7183 Abstract Multi-dimensional Image Model Dimension Unit (7183)
	AbstractMultiDimensionalImageModelDimensionUnit7183 = New("1.2.840.10008.6.1.920", "Abstract Multi-dimensional Image Model Dimension Unit (7183)", TypeContextGroupName, false)

	// AbstractMultiDimensionalImageModelAxisDirection7184 Abstract Multi-dimensional Image Model Axis Direction (7184)
	AbstractMultiDimensionalImageModelAxisDirection7184 = New("1.2.840.10008.6.1.921", "Abstract Multi-dimensional Image Model Axis Direction (7184)", TypeContextGroupName, false)

	// AbstractMultiDimensionalImageModelAxisOrientation7185 Abstract Multi-dimensional Image Model Axis Orientation (7185)
	AbstractMultiDimensionalImageModelAxisOrientation7185 = New("1.2.840.10008.6.1.922", "Abstract Multi-dimensional Image Model Axis Orientation (7185)", TypeContextGroupName, false)

	// AbstractMultiDimensionalImageModelQualitativeDimensionSampleSemantic7186 Abstract Multi-dimensional Image Model Qualitative Dimension Sample Semantic (7186)
	AbstractMultiDimensionalImageModelQualitativeDimensionSampleSemantic7186 = New("1.2.840.10008.6.1.923", "Abstract Multi-dimensional Image Model Qualitative Dimension Sample Semantic (7186)", TypeContextGroupName, false)

	// PlanningMethod7320 Planning Method (7320)
	PlanningMethod7320 = New("1.2.840.10008.6.1.924", "Planning Method (7320)", TypeContextGroupName, false)

	// DeIdentificationMethod7050 De-identification Method (7050)
	DeIdentificationMethod7050 = New("1.2.840.10008.6.1.925", "De-identification Method (7050)", TypeContextGroupName, false)

	// MeasurementOrientation12118 Measurement Orientation (12118)
	MeasurementOrientation12118 = New("1.2.840.10008.6.1.926", "Measurement Orientation (12118)", TypeContextGroupName, false)

	// ECGGlobalWaveformDuration3689 ECG Global Waveform Duration (3689)
	ECGGlobalWaveformDuration3689 = New("1.2.840.10008.6.1.927", "ECG Global Waveform Duration (3689)", TypeContextGroupName, false)

	// ICD3692 ICD (3692)
	ICD3692 = New("1.2.840.10008.6.1.930", "ICD (3692)", TypeContextGroupName, false)

	// RadiotherapyGeneralWorkitemDefinition9241 Radiotherapy General Workitem Definition (9241)
	RadiotherapyGeneralWorkitemDefinition9241 = New("1.2.840.10008.6.1.931", "Radiotherapy General Workitem Definition (9241)", TypeContextGroupName, false)

	// RadiotherapyAcquisitionWorkitemDefinition9242 Radiotherapy Acquisition Workitem Definition (9242)
	RadiotherapyAcquisitionWorkitemDefinition9242 = New("1.2.840.10008.6.1.932", "Radiotherapy Acquisition Workitem Definition (9242)", TypeContextGroupName, false)

	// RadiotherapyRegistrationWorkitemDefinition9243 Radiotherapy Registration Workitem Definition (9243)
	RadiotherapyRegistrationWorkitemDefinition9243 = New("1.2.840.10008.6.1.933", "Radiotherapy Registration Workitem Definition (9243)", TypeContextGroupName, false)

	// ContrastBolusSubstance3850 Contrast Bolus Substance (3850)
	ContrastBolusSubstance3850 = New("1.2.840.10008.6.1.934", "Contrast Bolus Substance (3850)", TypeContextGroupName, false)

	// LabelType10022 Label Type (10022)
	LabelType10022 = New("1.2.840.10008.6.1.935", "Label Type (10022)", TypeContextGroupName, false)

	// OphthalmicMappingUnitForRealWorldValueMapping4260 Ophthalmic Mapping Unit for Real World Value Mapping (4260)
	OphthalmicMappingUnitForRealWorldValueMapping4260 = New("1.2.840.10008.6.1.936", "Ophthalmic Mapping Unit for Real World Value Mapping (4260)", TypeContextGroupName, false)

	// OphthalmicMappingAcquisitionMethod4261 Ophthalmic Mapping Acquisition Method (4261)
	OphthalmicMappingAcquisitionMethod4261 = New("1.2.840.10008.6.1.937", "Ophthalmic Mapping Acquisition Method (4261)", TypeContextGroupName, false)

	// RetinalThicknessDefinition4262 Retinal Thickness Definition (4262)
	RetinalThicknessDefinition4262 = New("1.2.840.10008.6.1.938", "Retinal Thickness Definition (4262)", TypeContextGroupName, false)

	// OphthalmicThicknessMapValueType4263 Ophthalmic Thickness Map Value Type (4263)
	OphthalmicThicknessMapValueType4263 = New("1.2.840.10008.6.1.939", "Ophthalmic Thickness Map Value Type (4263)", TypeContextGroupName, false)

	// OphthalmicMapPurposeOfReference4264 Ophthalmic Map Purpose of Reference (4264)
	OphthalmicMapPurposeOfReference4264 = New("1.2.840.10008.6.1.940", "Ophthalmic Map Purpose of Reference (4264)", TypeContextGroupName, false)

	// OphthalmicThicknessDeviationCategory4265 Ophthalmic Thickness Deviation Category (4265)
	OphthalmicThicknessDeviationCategory4265 = New("1.2.840.10008.6.1.941", "Ophthalmic Thickness Deviation Category (4265)", TypeContextGroupName, false)

	// OphthalmicAnatomicStructureReferencePoint4266 Ophthalmic Anatomic Structure Reference Point (4266)
	OphthalmicAnatomicStructureReferencePoint4266 = New("1.2.840.10008.6.1.942", "Ophthalmic Anatomic Structure Reference Point (4266)", TypeContextGroupName, false)

	// CardiacSynchronizationTechnique3104 Cardiac Synchronization Technique (3104)
	CardiacSynchronizationTechnique3104 = New("1.2.840.10008.6.1.943", "Cardiac Synchronization Technique (3104)", TypeContextGroupName, false)

	// StainingProtocol8130 Staining Protocol (8130)
	StainingProtocol8130 = New("1.2.840.10008.6.1.944", "Staining Protocol (8130)", TypeContextGroupName, false)

	// SizeSpecificDoseEstimationMethodForCT10023 Size Specific Dose Estimation Method for CT (10023)
	SizeSpecificDoseEstimationMethodForCT10023 = New("1.2.840.10008.6.1.947", "Size Specific Dose Estimation Method for CT (10023)", TypeContextGroupName, false)

	// PathologyImagingProtocol8131 Pathology Imaging Protocol (8131)
	PathologyImagingProtocol8131 = New("1.2.840.10008.6.1.948", "Pathology Imaging Protocol (8131)", TypeContextGroupName, false)

	// MagnificationSelection8132 Magnification Selection (8132)
	MagnificationSelection8132 = New("1.2.840.10008.6.1.949", "Magnification Selection (8132)", TypeContextGroupName, false)

	// TissueSelection8133 Tissue Selection (8133)
	TissueSelection8133 = New("1.2.840.10008.6.1.950", "Tissue Selection (8133)", TypeContextGroupName, false)

	// GeneralRegionOfInterestMeasurementModifier7464 General Region of Interest Measurement Modifier (7464)
	GeneralRegionOfInterestMeasurementModifier7464 = New("1.2.840.10008.6.1.951", "General Region of Interest Measurement Modifier (7464)", TypeContextGroupName, false)

	// MeasurementDerivedFromMultipleROIMeasurements7465 Measurement Derived From Multiple ROI Measurements (7465)
	MeasurementDerivedFromMultipleROIMeasurements7465 = New("1.2.840.10008.6.1.952", "Measurement Derived From Multiple ROI Measurements (7465)", TypeContextGroupName, false)

	// SurfaceScanAcquisitionType8201 Surface Scan Acquisition Type (8201)
	SurfaceScanAcquisitionType8201 = New("1.2.840.10008.6.1.953", "Surface Scan Acquisition Type (8201)", TypeContextGroupName, false)

	// SurfaceScanModeType8202 Surface Scan Mode Type (8202)
	SurfaceScanModeType8202 = New("1.2.840.10008.6.1.954", "Surface Scan Mode Type (8202)", TypeContextGroupName, false)

	// SurfaceScanRegistrationMethodType8203 Surface Scan Registration Method Type (8203)
	SurfaceScanRegistrationMethodType8203 = New("1.2.840.10008.6.1.956", "Surface Scan Registration Method Type (8203)", TypeContextGroupName, false)

	// BasicCardiacView27 Basic Cardiac View (27)
	BasicCardiacView27 = New("1.2.840.10008.6.1.957", "Basic Cardiac View (27)", TypeContextGroupName, false)

	// CTReconstructionAlgorithm10033 CT Reconstruction Algorithm (10033)
	CTReconstructionAlgorithm10033 = New("1.2.840.10008.6.1.958", "CT Reconstruction Algorithm (10033)", TypeContextGroupName, false)

	// DetectorType10030 Detector Type (10030)
	DetectorType10030 = New("1.2.840.10008.6.1.959", "Detector Type (10030)", TypeContextGroupName, false)

	// CRDRMechanicalConfiguration10031 CR/DR Mechanical Configuration (10031)
	CRDRMechanicalConfiguration10031 = New("1.2.840.10008.6.1.960", "CR/DR Mechanical Configuration (10031)", TypeContextGroupName, false)

	// ProjectionXRayAcquisitionDeviceType10032 Projection X-Ray Acquisition Device Type (10032)
	ProjectionXRayAcquisitionDeviceType10032 = New("1.2.840.10008.6.1.961", "Projection X-Ray Acquisition Device Type (10032)", TypeContextGroupName, false)

	// AbstractSegmentationType7165 Abstract Segmentation Type (7165)
	AbstractSegmentationType7165 = New("1.2.840.10008.6.1.962", "Abstract Segmentation Type (7165)", TypeContextGroupName, false)

	// CommonTissueSegmentationType7166 Common Tissue Segmentation Type (7166)
	CommonTissueSegmentationType7166 = New("1.2.840.10008.6.1.963", "Common Tissue Segmentation Type (7166)", TypeContextGroupName, false)

	// PeripheralNervousSystemSegmentationType7167 Peripheral Nervous System Segmentation Type (7167)
	PeripheralNervousSystemSegmentationType7167 = New("1.2.840.10008.6.1.964", "Peripheral Nervous System Segmentation Type (7167)", TypeContextGroupName, false)

	// CornealTopographyMappingUnitForRealWorldValueMapping4267 Corneal Topography Mapping Unit for Real World Value Mapping (4267)
	CornealTopographyMappingUnitForRealWorldValueMapping4267 = New("1.2.840.10008.6.1.965", "Corneal Topography Mapping Unit for Real World Value Mapping (4267)", TypeContextGroupName, false)

	// CornealTopographyMapValueType4268 Corneal Topography Map Value Type (4268)
	CornealTopographyMapValueType4268 = New("1.2.840.10008.6.1.966", "Corneal Topography Map Value Type (4268)", TypeContextGroupName, false)

	// BrainStructureForVolumetricMeasurement7140 Brain Structure for Volumetric Measurement (7140)
	BrainStructureForVolumetricMeasurement7140 = New("1.2.840.10008.6.1.967", "Brain Structure for Volumetric Measurement (7140)", TypeContextGroupName, false)

	// RTDoseDerivation7220 RT Dose Derivation (7220)
	RTDoseDerivation7220 = New("1.2.840.10008.6.1.968", "RT Dose Derivation (7220)", TypeContextGroupName, false)

	// RTDosePurposeOfReference7221 RT Dose Purpose of Reference (7221)
	RTDosePurposeOfReference7221 = New("1.2.840.10008.6.1.969", "RT Dose Purpose of Reference (7221)", TypeContextGroupName, false)

	// SpectroscopyPurposeOfReference7215 Spectroscopy Purpose of Reference (7215)
	SpectroscopyPurposeOfReference7215 = New("1.2.840.10008.6.1.970", "Spectroscopy Purpose of Reference (7215)", TypeContextGroupName, false)

	// ScheduledProcessingParameterConceptCodesForRTTreatment9250 Scheduled Processing Parameter Concept Codes for RT Treatment (9250)
	ScheduledProcessingParameterConceptCodesForRTTreatment9250 = New("1.2.840.10008.6.1.971", "Scheduled Processing Parameter Concept Codes for RT Treatment (9250)", TypeContextGroupName, false)

	// RadiopharmaceuticalOrganDoseReferenceAuthority10040 Radiopharmaceutical Organ Dose Reference Authority (10040)
	RadiopharmaceuticalOrganDoseReferenceAuthority10040 = New("1.2.840.10008.6.1.972", "Radiopharmaceutical Organ Dose Reference Authority (10040)", TypeContextGroupName, false)

	// SourceOfRadioisotopeActivityInformation10041 Source of Radioisotope Activity Information (10041)
	SourceOfRadioisotopeActivityInformation10041 = New("1.2.840.10008.6.1.973", "Source of Radioisotope Activity Information (10041)", TypeContextGroupName, false)

	// IntravenousExtravasationSymptom10043 Intravenous Extravasation Symptom (10043)
	IntravenousExtravasationSymptom10043 = New("1.2.840.10008.6.1.975", "Intravenous Extravasation Symptom (10043)", TypeContextGroupName, false)

	// RadiosensitiveOrgan10044 Radiosensitive Organ (10044)
	RadiosensitiveOrgan10044 = New("1.2.840.10008.6.1.976", "Radiosensitive Organ (10044)", TypeContextGroupName, false)

	// RadiopharmaceuticalPatientState10045 Radiopharmaceutical Patient State (10045)
	RadiopharmaceuticalPatientState10045 = New("1.2.840.10008.6.1.977", "Radiopharmaceutical Patient State (10045)", TypeContextGroupName, false)

	// GFRMeasurement10046 GFR Measurement (10046)
	GFRMeasurement10046 = New("1.2.840.10008.6.1.978", "GFR Measurement (10046)", TypeContextGroupName, false)

	// GFRMeasurementMethod10047 GFR Measurement Method (10047)
	GFRMeasurementMethod10047 = New("1.2.840.10008.6.1.979", "GFR Measurement Method (10047)", TypeContextGroupName, false)

	// VisualEvaluationMethod8300 Visual Evaluation Method (8300)
	VisualEvaluationMethod8300 = New("1.2.840.10008.6.1.980", "Visual Evaluation Method (8300)", TypeContextGroupName, false)

	// TestPatternCode8301 Test Pattern Code (8301)
	TestPatternCode8301 = New("1.2.840.10008.6.1.981", "Test Pattern Code (8301)", TypeContextGroupName, false)

	// MeasurementPatternCode8302 Measurement Pattern Code (8302)
	MeasurementPatternCode8302 = New("1.2.840.10008.6.1.982", "Measurement Pattern Code (8302)", TypeContextGroupName, false)

	// DisplayDeviceType8303 Display Device Type (8303)
	DisplayDeviceType8303 = New("1.2.840.10008.6.1.983", "Display Device Type (8303)", TypeContextGroupName, false)

	// SUVUnit85 SUV Unit (85)
	SUVUnit85 = New("1.2.840.10008.6.1.984", "SUV Unit (85)", TypeContextGroupName, false)

	// T1MeasurementMethod4100 T1 Measurement Method (4100)
	T1MeasurementMethod4100 = New("1.2.840.10008.6.1.985", "T1 Measurement Method (4100)", TypeContextGroupName, false)

	// TracerKineticModel4101 Tracer Kinetic Model (4101)
	TracerKineticModel4101 = New("1.2.840.10008.6.1.986", "Tracer Kinetic Model (4101)", TypeContextGroupName, false)

	// PerfusionMeasurementMethod4102 Perfusion Measurement Method (4102)
	PerfusionMeasurementMethod4102 = New("1.2.840.10008.6.1.987", "Perfusion Measurement Method (4102)", TypeContextGroupName, false)

	// ArterialInputFunctionMeasurementMethod4103 Arterial Input Function Measurement Method (4103)
	ArterialInputFunctionMeasurementMethod4103 = New("1.2.840.10008.6.1.988", "Arterial Input Function Measurement Method (4103)", TypeContextGroupName, false)

	// BolusArrivalTimeDerivationMethod4104 Bolus Arrival Time Derivation Method (4104)
	BolusArrivalTimeDerivationMethod4104 = New("1.2.840.10008.6.1.989", "Bolus Arrival Time Derivation Method (4104)", TypeContextGroupName, false)

	// PerfusionAnalysisMethod4105 Perfusion Analysis Method (4105)
	PerfusionAnalysisMethod4105 = New("1.2.840.10008.6.1.990", "Perfusion Analysis Method (4105)", TypeContextGroupName, false)

	// QuantitativeMethodUsedForPerfusionAndTracerKineticModel4106 Quantitative Method Used for Perfusion and Tracer Kinetic Model (4106)
	QuantitativeMethodUsedForPerfusionAndTracerKineticModel4106 = New("1.2.840.10008.6.1.991", "Quantitative Method Used for Perfusion and Tracer Kinetic Model (4106)", TypeContextGroupName, false)

	// TracerKineticModelParameter4107 Tracer Kinetic Model Parameter (4107)
	TracerKineticModelParameter4107 = New("1.2.840.10008.6.1.992", "Tracer Kinetic Model Parameter (4107)", TypeContextGroupName, false)

	// PerfusionModelParameter4108 Perfusion Model Parameter (4108)
	PerfusionModelParameter4108 = New("1.2.840.10008.6.1.993", "Perfusion Model Parameter (4108)", TypeContextGroupName, false)

	// ModelIndependentDynamicContrastAnalysisParameter4109 Model-Independent Dynamic Contrast Analysis Parameter (4109)
	ModelIndependentDynamicContrastAnalysisParameter4109 = New("1.2.840.10008.6.1.994", "Model-Independent Dynamic Contrast Analysis Parameter (4109)", TypeContextGroupName, false)

	// TracerKineticModelingCovariate4110 Tracer Kinetic Modeling Covariate (4110)
	TracerKineticModelingCovariate4110 = New("1.2.840.10008.6.1.995", "Tracer Kinetic Modeling Covariate (4110)", TypeContextGroupName, false)

	// ContrastCharacteristic4111 Contrast Characteristic (4111)
	ContrastCharacteristic4111 = New("1.2.840.10008.6.1.996", "Contrast Characteristic (4111)", TypeContextGroupName, false)

	// MeasurementReportDocumentTitle7021 Measurement Report Document Title (7021)
	MeasurementReportDocumentTitle7021 = New("1.2.840.10008.6.1.997", "Measurement Report Document Title (7021)", TypeContextGroupName, false)

	// QuantitativeDiagnosticImagingProcedure100 Quantitative Diagnostic Imaging Procedure (100)
	QuantitativeDiagnosticImagingProcedure100 = New("1.2.840.10008.6.1.998", "Quantitative Diagnostic Imaging Procedure (100)", TypeContextGroupName, false)

	// PETRegionOfInterestMeasurement7466 PET Region of Interest Measurement (7466)
	PETRegionOfInterestMeasurement7466 = New("1.2.840.10008.6.1.999", "PET Region of Interest Measurement (7466)", TypeContextGroupName, false)

	// GrayLevelCoOccurrenceMatrixMeasurement7467 Gray Level Co-occurrence Matrix Measurement (7467)
	GrayLevelCoOccurrenceMatrixMeasurement7467 = New("1.2.840.10008.6.1.1000", "Gray Level Co-occurrence Matrix Measurement (7467)", TypeContextGroupName, false)

	// TextureMeasurement7468 Texture Measurement (7468)
	TextureMeasurement7468 = New("1.2.840.10008.6.1.1001", "Texture Measurement (7468)", TypeContextGroupName, false)

	// TimePointType6146 Time Point Type (6146)
	TimePointType6146 = New("1.2.840.10008.6.1.1002", "Time Point Type (6146)", TypeContextGroupName, false)

	// GenericIntensityAndSizeMeasurement7469 Generic Intensity and Size Measurement (7469)
	GenericIntensityAndSizeMeasurement7469 = New("1.2.840.10008.6.1.1003", "Generic Intensity and Size Measurement (7469)", TypeContextGroupName, false)

	// ResponseCriteria6147 Response Criteria (6147)
	ResponseCriteria6147 = New("1.2.840.10008.6.1.1004", "Response Criteria (6147)", TypeContextGroupName, false)

	// FetalBiometryAnatomicSite12020 Fetal Biometry Anatomic Site (12020)
	FetalBiometryAnatomicSite12020 = New("1.2.840.10008.6.1.1005", "Fetal Biometry Anatomic Site (12020)", TypeContextGroupName, false)

	// FetalLongBoneAnatomicSite12021 Fetal Long Bone Anatomic Site (12021)
	FetalLongBoneAnatomicSite12021 = New("1.2.840.10008.6.1.1006", "Fetal Long Bone Anatomic Site (12021)", TypeContextGroupName, false)

	// FetalCraniumAnatomicSite12022 Fetal Cranium Anatomic Site (12022)
	FetalCraniumAnatomicSite12022 = New("1.2.840.10008.6.1.1007", "Fetal Cranium Anatomic Site (12022)", TypeContextGroupName, false)

	// PelvisAndUterusAnatomicSite12023 Pelvis and Uterus Anatomic Site (12023)
	PelvisAndUterusAnatomicSite12023 = New("1.2.840.10008.6.1.1008", "Pelvis and Uterus Anatomic Site (12023)", TypeContextGroupName, false)

	// ParametricMapDerivationImagePurposeOfReference7222 Parametric Map Derivation Image Purpose of Reference (7222)
	ParametricMapDerivationImagePurposeOfReference7222 = New("1.2.840.10008.6.1.1009", "Parametric Map Derivation Image Purpose of Reference (7222)", TypeContextGroupName, false)

	// PhysicalQuantityDescriptor9000 Physical Quantity Descriptor (9000)
	PhysicalQuantityDescriptor9000 = New("1.2.840.10008.6.1.1010", "Physical Quantity Descriptor (9000)", TypeContextGroupName, false)

	// LymphNodeAnatomicSite7600 Lymph Node Anatomic Site (7600)
	LymphNodeAnatomicSite7600 = New("1.2.840.10008.6.1.1011", "Lymph Node Anatomic Site (7600)", TypeContextGroupName, false)

	// HeadAndNeckCancerAnatomicSite7601 Head and Neck Cancer Anatomic Site (7601)
	HeadAndNeckCancerAnatomicSite7601 = New("1.2.840.10008.6.1.1012", "Head and Neck Cancer Anatomic Site (7601)", TypeContextGroupName, false)

	// FiberTractInBrainstem7701 Fiber Tract In Brainstem (7701)
	FiberTractInBrainstem7701 = New("1.2.840.10008.6.1.1013", "Fiber Tract In Brainstem (7701)", TypeContextGroupName, false)

	// ProjectionAndThalamicFiber7702 Projection and Thalamic Fiber (7702)
	ProjectionAndThalamicFiber7702 = New("1.2.840.10008.6.1.1014", "Projection and Thalamic Fiber (7702)", TypeContextGroupName, false)

	// AssociationFiber7703 Association Fiber (7703)
	AssociationFiber7703 = New("1.2.840.10008.6.1.1015", "Association Fiber (7703)", TypeContextGroupName, false)

	// LimbicSystemTract7704 Limbic System Tract (7704)
	LimbicSystemTract7704 = New("1.2.840.10008.6.1.1016", "Limbic System Tract (7704)", TypeContextGroupName, false)

	// CommissuralFiber7705 Commissural Fiber (7705)
	CommissuralFiber7705 = New("1.2.840.10008.6.1.1017", "Commissural Fiber (7705)", TypeContextGroupName, false)

	// CranialNerve7706 Cranial Nerve (7706)
	CranialNerve7706 = New("1.2.840.10008.6.1.1018", "Cranial Nerve (7706)", TypeContextGroupName, false)

	// SpinalCordFiber7707 Spinal Cord Fiber (7707)
	SpinalCordFiber7707 = New("1.2.840.10008.6.1.1019", "Spinal Cord Fiber (7707)", TypeContextGroupName, false)

	// TractographyAnatomicSite7710 Tractography Anatomic Site (7710)
	TractographyAnatomicSite7710 = New("1.2.840.10008.6.1.1020", "Tractography Anatomic Site (7710)", TypeContextGroupName, false)

	// PrimaryAnatomicStructureForIntraOralRadiographySupernumeraryDentitionDesignationOfTeeth4025 Primary Anatomic Structure for Intra-oral Radiography (Supernumerary Dentition - Designation of Teeth) (4025)
	PrimaryAnatomicStructureForIntraOralRadiographySupernumeraryDentitionDesignationOfTeeth4025 = New("1.2.840.10008.6.1.1021", "Primary Anatomic Structure for Intra-oral Radiography (Supernumerary Dentition - Designation of Teeth) (4025)", TypeContextGroupName, false)

	// PrimaryAnatomicStructureForIntraOralAndCraniofacialRadiographyTeeth4026 Primary Anatomic Structure for Intra-oral and Craniofacial Radiography - Teeth (4026)
	PrimaryAnatomicStructureForIntraOralAndCraniofacialRadiographyTeeth4026 = New("1.2.840.10008.6.1.1022", "Primary Anatomic Structure for Intra-oral and Craniofacial Radiography - Teeth (4026)", TypeContextGroupName, false)

	// IEC61217DevicePositionParameter9401 IEC61217 Device Position Parameter (9401)
	IEC61217DevicePositionParameter9401 = New("1.2.840.10008.6.1.1023", "IEC61217 Device Position Parameter (9401)", TypeContextGroupName, false)

	// IEC61217GantryPositionParameter9402 IEC61217 Gantry Position Parameter (9402)
	IEC61217GantryPositionParameter9402 = New("1.2.840.10008.6.1.1024", "IEC61217 Gantry Position Parameter (9402)", TypeContextGroupName, false)

	// IEC61217PatientSupportPositionParameter9403 IEC61217 Patient Support Position Parameter (9403)
	IEC61217PatientSupportPositionParameter9403 = New("1.2.840.10008.6.1.1025", "IEC61217 Patient Support Position Parameter (9403)", TypeContextGroupName, false)

	// ActionableFindingClassification7035 Actionable Finding Classification (7035)
	ActionableFindingClassification7035 = New("1.2.840.10008.6.1.1026", "Actionable Finding Classification (7035)", TypeContextGroupName, false)

	// ImageQualityAssessment7036 Image Quality Assessment (7036)
	ImageQualityAssessment7036 = New("1.2.840.10008.6.1.1027", "Image Quality Assessment (7036)", TypeContextGroupName, false)

	// SummaryRadiationExposureQuantity10050 Summary Radiation Exposure Quantity (10050)
	SummaryRadiationExposureQuantity10050 = New("1.2.840.10008.6.1.1028", "Summary Radiation Exposure Quantity (10050)", TypeContextGroupName, false)

	// WideFieldOphthalmicPhotographyTransformationMethod4245 Wide Field Ophthalmic Photography Transformation Method (4245)
	WideFieldOphthalmicPhotographyTransformationMethod4245 = New("1.2.840.10008.6.1.1029", "Wide Field Ophthalmic Photography Transformation Method (4245)", TypeContextGroupName, false)

	// PETUnit84 PET Unit (84)
	PETUnit84 = New("1.2.840.10008.6.1.1030", "PET Unit (84)", TypeContextGroupName, false)

	// ImplantMaterial7300 Implant Material (7300)
	ImplantMaterial7300 = New("1.2.840.10008.6.1.1031", "Implant Material (7300)", TypeContextGroupName, false)

	// InterventionType7301 Intervention Type (7301)
	InterventionType7301 = New("1.2.840.10008.6.1.1032", "Intervention Type (7301)", TypeContextGroupName, false)

	// ImplantTemplateViewOrientation7302 Implant Template View Orientation (7302)
	ImplantTemplateViewOrientation7302 = New("1.2.840.10008.6.1.1033", "Implant Template View Orientation (7302)", TypeContextGroupName, false)

	// ImplantTemplateModifiedViewOrientation7303 Implant Template Modified View Orientation (7303)
	ImplantTemplateModifiedViewOrientation7303 = New("1.2.840.10008.6.1.1034", "Implant Template Modified View Orientation (7303)", TypeContextGroupName, false)

	// ImplantTargetAnatomy7304 Implant Target Anatomy (7304)
	ImplantTargetAnatomy7304 = New("1.2.840.10008.6.1.1035", "Implant Target Anatomy (7304)", TypeContextGroupName, false)

	// ImplantPlanningLandmark7305 Implant Planning Landmark (7305)
	ImplantPlanningLandmark7305 = New("1.2.840.10008.6.1.1036", "Implant Planning Landmark (7305)", TypeContextGroupName, false)

	// HumanHipImplantPlanningLandmark7306 Human Hip Implant Planning Landmark (7306)
	HumanHipImplantPlanningLandmark7306 = New("1.2.840.10008.6.1.1037", "Human Hip Implant Planning Landmark (7306)", TypeContextGroupName, false)

	// ImplantComponentType7307 Implant Component Type (7307)
	ImplantComponentType7307 = New("1.2.840.10008.6.1.1038", "Implant Component Type (7307)", TypeContextGroupName, false)

	// HumanHipImplantComponentType7308 Human Hip Implant Component Type (7308)
	HumanHipImplantComponentType7308 = New("1.2.840.10008.6.1.1039", "Human Hip Implant Component Type (7308)", TypeContextGroupName, false)

	// HumanTraumaImplantComponentType7309 Human Trauma Implant Component Type (7309)
	HumanTraumaImplantComponentType7309 = New("1.2.840.10008.6.1.1040", "Human Trauma Implant Component Type (7309)", TypeContextGroupName, false)

	// ImplantFixationMethod7310 Implant Fixation Method (7310)
	ImplantFixationMethod7310 = New("1.2.840.10008.6.1.1041", "Implant Fixation Method (7310)", TypeContextGroupName, false)

	// DeviceParticipatingRole7445 Device Participating Role (7445)
	DeviceParticipatingRole7445 = New("1.2.840.10008.6.1.1042", "Device Participating Role (7445)", TypeContextGroupName, false)

	// ContainerType8101 Container Type (8101)
	ContainerType8101 = New("1.2.840.10008.6.1.1043", "Container Type (8101)", TypeContextGroupName, false)

	// ContainerComponentType8102 Container Component Type (8102)
	ContainerComponentType8102 = New("1.2.840.10008.6.1.1044", "Container Component Type (8102)", TypeContextGroupName, false)

	// AnatomicPathologySpecimenType8103 Anatomic Pathology Specimen Type (8103)
	AnatomicPathologySpecimenType8103 = New("1.2.840.10008.6.1.1045", "Anatomic Pathology Specimen Type (8103)", TypeContextGroupName, false)

	// BreastTissueSpecimenType8104 Breast Tissue Specimen Type (8104)
	BreastTissueSpecimenType8104 = New("1.2.840.10008.6.1.1046", "Breast Tissue Specimen Type (8104)", TypeContextGroupName, false)

	// SpecimenCollectionProcedure8109 Specimen Collection Procedure (8109)
	SpecimenCollectionProcedure8109 = New("1.2.840.10008.6.1.1047", "Specimen Collection Procedure (8109)", TypeContextGroupName, false)

	// SpecimenSamplingProcedure8110 Specimen Sampling Procedure (8110)
	SpecimenSamplingProcedure8110 = New("1.2.840.10008.6.1.1048", "Specimen Sampling Procedure (8110)", TypeContextGroupName, false)

	// SpecimenPreparationProcedure8111 Specimen Preparation Procedure (8111)
	SpecimenPreparationProcedure8111 = New("1.2.840.10008.6.1.1049", "Specimen Preparation Procedure (8111)", TypeContextGroupName, false)

	// SpecimenStain8112 Specimen Stain (8112)
	SpecimenStain8112 = New("1.2.840.10008.6.1.1050", "Specimen Stain (8112)", TypeContextGroupName, false)

	// SpecimenPreparationStep8113 Specimen Preparation Step (8113)
	SpecimenPreparationStep8113 = New("1.2.840.10008.6.1.1051", "Specimen Preparation Step (8113)", TypeContextGroupName, false)

	// SpecimenFixative8114 Specimen Fixative (8114)
	SpecimenFixative8114 = New("1.2.840.10008.6.1.1052", "Specimen Fixative (8114)", TypeContextGroupName, false)

	// SpecimenEmbeddingMedia8115 Specimen Embedding Media (8115)
	SpecimenEmbeddingMedia8115 = New("1.2.840.10008.6.1.1053", "Specimen Embedding Media (8115)", TypeContextGroupName, false)

	// SourceOfProjectionXRayDoseInformation10020 Source of Projection X-Ray Dose Information (10020)
	SourceOfProjectionXRayDoseInformation10020 = New("1.2.840.10008.6.1.1054", "Source of Projection X-Ray Dose Information (10020)", TypeContextGroupName, false)

	// SourceOfCTDoseInformation10021 Source of CT Dose Information (10021)
	SourceOfCTDoseInformation10021 = New("1.2.840.10008.6.1.1055", "Source of CT Dose Information (10021)", TypeContextGroupName, false)

	// RadiationDoseReferencePoint10025 Radiation Dose Reference Point (10025)
	RadiationDoseReferencePoint10025 = New("1.2.840.10008.6.1.1056", "Radiation Dose Reference Point (10025)", TypeContextGroupName, false)

	// VolumetricViewDescription501 Volumetric View Description (501)
	VolumetricViewDescription501 = New("1.2.840.10008.6.1.1057", "Volumetric View Description (501)", TypeContextGroupName, false)

	// VolumetricViewModifier502 Volumetric View Modifier (502)
	VolumetricViewModifier502 = New("1.2.840.10008.6.1.1058", "Volumetric View Modifier (502)", TypeContextGroupName, false)

	// DiffusionAcquisitionValueType7260 Diffusion Acquisition Value Type (7260)
	DiffusionAcquisitionValueType7260 = New("1.2.840.10008.6.1.1059", "Diffusion Acquisition Value Type (7260)", TypeContextGroupName, false)

	// DiffusionModelValueType7261 Diffusion Model Value Type (7261)
	DiffusionModelValueType7261 = New("1.2.840.10008.6.1.1060", "Diffusion Model Value Type (7261)", TypeContextGroupName, false)

	// DiffusionTractographyAlgorithmFamily7262 Diffusion Tractography Algorithm Family (7262)
	DiffusionTractographyAlgorithmFamily7262 = New("1.2.840.10008.6.1.1061", "Diffusion Tractography Algorithm Family (7262)", TypeContextGroupName, false)

	// DiffusionTractographyMeasurementType7263 Diffusion Tractography Measurement Type (7263)
	DiffusionTractographyMeasurementType7263 = New("1.2.840.10008.6.1.1062", "Diffusion Tractography Measurement Type (7263)", TypeContextGroupName, false)

	// ResearchAnimalSourceRegistry7490 Research Animal Source Registry (7490)
	ResearchAnimalSourceRegistry7490 = New("1.2.840.10008.6.1.1063", "Research Animal Source Registry (7490)", TypeContextGroupName, false)

	// YesNoOnly231 Yes-No Only (231)
	YesNoOnly231 = New("1.2.840.10008.6.1.1064", "Yes-No Only (231)", TypeContextGroupName, false)

	// BiosafetyLevel601 Biosafety Level (601)
	BiosafetyLevel601 = New("1.2.840.10008.6.1.1065", "Biosafety Level (601)", TypeContextGroupName, false)

	// BiosafetyControlReason602 Biosafety Control Reason (602)
	BiosafetyControlReason602 = New("1.2.840.10008.6.1.1066", "Biosafety Control Reason (602)", TypeContextGroupName, false)

	// SexMaleFemaleOrBoth7457 Sex - Male Female or Both (7457)
	SexMaleFemaleOrBoth7457 = New("1.2.840.10008.6.1.1067", "Sex - Male Female or Both (7457)", TypeContextGroupName, false)

	// AnimalRoomType603 Animal Room Type (603)
	AnimalRoomType603 = New("1.2.840.10008.6.1.1068", "Animal Room Type (603)", TypeContextGroupName, false)

	// DeviceReuse604 Device Reuse (604)
	DeviceReuse604 = New("1.2.840.10008.6.1.1069", "Device Reuse (604)", TypeContextGroupName, false)

	// AnimalBeddingMaterial605 Animal Bedding Material (605)
	AnimalBeddingMaterial605 = New("1.2.840.10008.6.1.1070", "Animal Bedding Material (605)", TypeContextGroupName, false)

	// AnimalShelterType606 Animal Shelter Type (606)
	AnimalShelterType606 = New("1.2.840.10008.6.1.1071", "Animal Shelter Type (606)", TypeContextGroupName, false)

	// AnimalFeedType607 Animal Feed Type (607)
	AnimalFeedType607 = New("1.2.840.10008.6.1.1072", "Animal Feed Type (607)", TypeContextGroupName, false)

	// AnimalFeedSource608 Animal Feed Source (608)
	AnimalFeedSource608 = New("1.2.840.10008.6.1.1073", "Animal Feed Source (608)", TypeContextGroupName, false)

	// AnimalFeedingMethod609 Animal Feeding Method (609)
	AnimalFeedingMethod609 = New("1.2.840.10008.6.1.1074", "Animal Feeding Method (609)", TypeContextGroupName, false)

	// WaterType610 Water Type (610)
	WaterType610 = New("1.2.840.10008.6.1.1075", "Water Type (610)", TypeContextGroupName, false)

	// AnesthesiaCategoryCodeTypeForSmallAnimalAnesthesia611 Anesthesia Category Code Type for Small Animal Anesthesia (611)
	AnesthesiaCategoryCodeTypeForSmallAnimalAnesthesia611 = New("1.2.840.10008.6.1.1076", "Anesthesia Category Code Type for Small Animal Anesthesia (611)", TypeContextGroupName, false)

	// AnesthesiaCategoryCodeTypeFromAnesthesiaQualityInitiative612 Anesthesia Category Code Type from Anesthesia Quality Initiative (612)
	AnesthesiaCategoryCodeTypeFromAnesthesiaQualityInitiative612 = New("1.2.840.10008.6.1.1077", "Anesthesia Category Code Type from Anesthesia Quality Initiative (612)", TypeContextGroupName, false)

	// AnesthesiaInductionCodeTypeForSmallAnimalAnesthesia613 Anesthesia Induction Code Type for Small Animal Anesthesia (613)
	AnesthesiaInductionCodeTypeForSmallAnimalAnesthesia613 = New("1.2.840.10008.6.1.1078", "Anesthesia Induction Code Type for Small Animal Anesthesia (613)", TypeContextGroupName, false)

	// AnesthesiaInductionCodeTypeFromAnesthesiaQualityInitiative614 Anesthesia Induction Code Type from Anesthesia Quality Initiative (614)
	AnesthesiaInductionCodeTypeFromAnesthesiaQualityInitiative614 = New("1.2.840.10008.6.1.1079", "Anesthesia Induction Code Type from Anesthesia Quality Initiative (614)", TypeContextGroupName, false)

	// AnesthesiaMaintenanceCodeTypeForSmallAnimalAnesthesia615 Anesthesia Maintenance Code Type for Small Animal Anesthesia (615)
	AnesthesiaMaintenanceCodeTypeForSmallAnimalAnesthesia615 = New("1.2.840.10008.6.1.1080", "Anesthesia Maintenance Code Type for Small Animal Anesthesia (615)", TypeContextGroupName, false)

	// AnesthesiaMaintenanceCodeTypeFromAnesthesiaQualityInitiative616 Anesthesia Maintenance Code Type from Anesthesia Quality Initiative (616)
	AnesthesiaMaintenanceCodeTypeFromAnesthesiaQualityInitiative616 = New("1.2.840.10008.6.1.1081", "Anesthesia Maintenance Code Type from Anesthesia Quality Initiative (616)", TypeContextGroupName, false)

	// AirwayManagementMethodCodeTypeForSmallAnimalAnesthesia617 Airway Management Method Code Type for Small Animal Anesthesia (617)
	AirwayManagementMethodCodeTypeForSmallAnimalAnesthesia617 = New("1.2.840.10008.6.1.1082", "Airway Management Method Code Type for Small Animal Anesthesia (617)", TypeContextGroupName, false)

	// AirwayManagementMethodCodeTypeFromAnesthesiaQualityInitiative618 Airway Management Method Code Type from Anesthesia Quality Initiative (618)
	AirwayManagementMethodCodeTypeFromAnesthesiaQualityInitiative618 = New("1.2.840.10008.6.1.1083", "Airway Management Method Code Type from Anesthesia Quality Initiative (618)", TypeContextGroupName, false)

	// AirwayManagementSubMethodCodeTypeForSmallAnimalAnesthesia619 Airway Management Sub-Method Code Type for Small Animal Anesthesia (619)
	AirwayManagementSubMethodCodeTypeForSmallAnimalAnesthesia619 = New("1.2.840.10008.6.1.1084", "Airway Management Sub-Method Code Type for Small Animal Anesthesia (619)", TypeContextGroupName, false)

	// AirwayManagementSubMethodCodeTypeFromAnesthesiaQualityInitiative620 Airway Management Sub-Method Code Type from Anesthesia Quality Initiative (620)
	AirwayManagementSubMethodCodeTypeFromAnesthesiaQualityInitiative620 = New("1.2.840.10008.6.1.1085", "Airway Management Sub-Method Code Type from Anesthesia Quality Initiative (620)", TypeContextGroupName, false)

	// MedicationTypeForSmallAnimalAnesthesia621 Medication Type for Small Animal Anesthesia (621)
	MedicationTypeForSmallAnimalAnesthesia621 = New("1.2.840.10008.6.1.1086", "Medication Type for Small Animal Anesthesia (621)", TypeContextGroupName, false)

	// MedicationTypeCodeTypeFromAnesthesiaQualityInitiative622 Medication Type Code Type from Anesthesia Quality Initiative (622)
	MedicationTypeCodeTypeFromAnesthesiaQualityInitiative622 = New("1.2.840.10008.6.1.1087", "Medication Type Code Type from Anesthesia Quality Initiative (622)", TypeContextGroupName, false)

	// MedicationForSmallAnimalAnesthesia623 Medication for Small Animal Anesthesia (623)
	MedicationForSmallAnimalAnesthesia623 = New("1.2.840.10008.6.1.1088", "Medication for Small Animal Anesthesia (623)", TypeContextGroupName, false)

	// InhalationalAnesthesiaAgentForSmallAnimalAnesthesia624 Inhalational Anesthesia Agent for Small Animal Anesthesia (624)
	InhalationalAnesthesiaAgentForSmallAnimalAnesthesia624 = New("1.2.840.10008.6.1.1089", "Inhalational Anesthesia Agent for Small Animal Anesthesia (624)", TypeContextGroupName, false)

	// InjectableAnesthesiaAgentForSmallAnimalAnesthesia625 Injectable Anesthesia Agent for Small Animal Anesthesia (625)
	InjectableAnesthesiaAgentForSmallAnimalAnesthesia625 = New("1.2.840.10008.6.1.1090", "Injectable Anesthesia Agent for Small Animal Anesthesia (625)", TypeContextGroupName, false)

	// PremedicationAgentForSmallAnimalAnesthesia626 Premedication Agent for Small Animal Anesthesia (626)
	PremedicationAgentForSmallAnimalAnesthesia626 = New("1.2.840.10008.6.1.1091", "Premedication Agent for Small Animal Anesthesia (626)", TypeContextGroupName, false)

	// NeuromuscularBlockingAgentForSmallAnimalAnesthesia627 Neuromuscular Blocking Agent for Small Animal Anesthesia (627)
	NeuromuscularBlockingAgentForSmallAnimalAnesthesia627 = New("1.2.840.10008.6.1.1092", "Neuromuscular Blocking Agent for Small Animal Anesthesia (627)", TypeContextGroupName, false)

	// AncillaryMedicationsForSmallAnimalAnesthesia628 Ancillary Medications for Small Animal Anesthesia (628)
	AncillaryMedicationsForSmallAnimalAnesthesia628 = New("1.2.840.10008.6.1.1093", "Ancillary Medications for Small Animal Anesthesia (628)", TypeContextGroupName, false)

	// CarrierGasesForSmallAnimalAnesthesia629 Carrier Gases for Small Animal Anesthesia (629)
	CarrierGasesForSmallAnimalAnesthesia629 = New("1.2.840.10008.6.1.1094", "Carrier Gases for Small Animal Anesthesia (629)", TypeContextGroupName, false)

	// LocalAnestheticsForSmallAnimalAnesthesia630 Local Anesthetics for Small Animal Anesthesia (630)
	LocalAnestheticsForSmallAnimalAnesthesia630 = New("1.2.840.10008.6.1.1095", "Local Anesthetics for Small Animal Anesthesia (630)", TypeContextGroupName, false)

	// ProcedurePhaseRequiringAnesthesia631 Procedure Phase Requiring Anesthesia (631)
	ProcedurePhaseRequiringAnesthesia631 = New("1.2.840.10008.6.1.1096", "Procedure Phase Requiring Anesthesia (631)", TypeContextGroupName, false)

	// SurgicalProcedurePhaseRequiringAnesthesia632 Surgical Procedure Phase Requiring Anesthesia (632)
	SurgicalProcedurePhaseRequiringAnesthesia632 = New("1.2.840.10008.6.1.1097", "Surgical Procedure Phase Requiring Anesthesia (632)", TypeContextGroupName, false)

	// PhaseOfImagingProcedureRequiringAnesthesia633RETIRED Phase of Imaging Procedure Requiring Anesthesia (Retired) (633)
	PhaseOfImagingProcedureRequiringAnesthesia633RETIRED = New("1.2.840.10008.6.1.1098", "Phase of Imaging Procedure Requiring Anesthesia (Retired) (633)", TypeContextGroupName, true)

	// AnimalHandlingPhase634 Animal Handling Phase (634)
	AnimalHandlingPhase634 = New("1.2.840.10008.6.1.1099", "Animal Handling Phase (634)", TypeContextGroupName, false)

	// HeatingMethod635 Heating Method (635)
	HeatingMethod635 = New("1.2.840.10008.6.1.1100", "Heating Method (635)", TypeContextGroupName, false)

	// TemperatureSensorDeviceComponentTypeForSmallAnimalProcedure636 Temperature Sensor Device Component Type for Small Animal Procedure (636)
	TemperatureSensorDeviceComponentTypeForSmallAnimalProcedure636 = New("1.2.840.10008.6.1.1101", "Temperature Sensor Device Component Type for Small Animal Procedure (636)", TypeContextGroupName, false)

	// ExogenousSubstanceType637 Exogenous Substance Type (637)
	ExogenousSubstanceType637 = New("1.2.840.10008.6.1.1102", "Exogenous Substance Type (637)", TypeContextGroupName, false)

	// ExogenousSubstance638 Exogenous Substance (638)
	ExogenousSubstance638 = New("1.2.840.10008.6.1.1103", "Exogenous Substance (638)", TypeContextGroupName, false)

	// TumorGraftHistologicType639 Tumor Graft Histologic Type (639)
	TumorGraftHistologicType639 = New("1.2.840.10008.6.1.1104", "Tumor Graft Histologic Type (639)", TypeContextGroupName, false)

	// Fibril640 Fibril (640)
	Fibril640 = New("1.2.840.10008.6.1.1105", "Fibril (640)", TypeContextGroupName, false)

	// Virus641 Virus (641)
	Virus641 = New("1.2.840.10008.6.1.1106", "Virus (641)", TypeContextGroupName, false)

	// Cytokine642 Cytokine (642)
	Cytokine642 = New("1.2.840.10008.6.1.1107", "Cytokine (642)", TypeContextGroupName, false)

	// Toxin643 Toxin (643)
	Toxin643 = New("1.2.840.10008.6.1.1108", "Toxin (643)", TypeContextGroupName, false)

	// ExogenousSubstanceAdministrationSite644 Exogenous Substance Administration Site (644)
	ExogenousSubstanceAdministrationSite644 = New("1.2.840.10008.6.1.1109", "Exogenous Substance Administration Site (644)", TypeContextGroupName, false)

	// ExogenousSubstanceOriginTissue645 Exogenous Substance Origin Tissue (645)
	ExogenousSubstanceOriginTissue645 = New("1.2.840.10008.6.1.1110", "Exogenous Substance Origin Tissue (645)", TypeContextGroupName, false)

	// PreclinicalSmallAnimalImagingProcedure646 Preclinical Small Animal Imaging Procedure (646)
	PreclinicalSmallAnimalImagingProcedure646 = New("1.2.840.10008.6.1.1111", "Preclinical Small Animal Imaging Procedure (646)", TypeContextGroupName, false)

	// PositionReferenceIndicatorForFrameOfReference647 Position Reference Indicator for Frame of Reference (647)
	PositionReferenceIndicatorForFrameOfReference647 = New("1.2.840.10008.6.1.1112", "Position Reference Indicator for Frame of Reference (647)", TypeContextGroupName, false)

	// PresentAbsentOnly241 Present-Absent Only (241)
	PresentAbsentOnly241 = New("1.2.840.10008.6.1.1113", "Present-Absent Only (241)", TypeContextGroupName, false)

	// WaterEquivalentDiameterMethod10024 Water Equivalent Diameter Method (10024)
	WaterEquivalentDiameterMethod10024 = New("1.2.840.10008.6.1.1114", "Water Equivalent Diameter Method (10024)", TypeContextGroupName, false)

	// RadiotherapyPurposeOfReference7022 Radiotherapy Purpose of Reference (7022)
	RadiotherapyPurposeOfReference7022 = New("1.2.840.10008.6.1.1115", "Radiotherapy Purpose of Reference (7022)", TypeContextGroupName, false)

	// ContentAssessmentType701 Content Assessment Type (701)
	ContentAssessmentType701 = New("1.2.840.10008.6.1.1116", "Content Assessment Type (701)", TypeContextGroupName, false)

	// RTContentAssessmentType702 RT Content Assessment Type (702)
	RTContentAssessmentType702 = New("1.2.840.10008.6.1.1117", "RT Content Assessment Type (702)", TypeContextGroupName, false)

	// AssessmentBasis703 Assessment Basis (703)
	AssessmentBasis703 = New("1.2.840.10008.6.1.1118", "Assessment Basis (703)", TypeContextGroupName, false)

	// ReaderSpecialty7449 Reader Specialty (7449)
	ReaderSpecialty7449 = New("1.2.840.10008.6.1.1119", "Reader Specialty (7449)", TypeContextGroupName, false)

	// RequestedReportType9233 Requested Report Type (9233)
	RequestedReportType9233 = New("1.2.840.10008.6.1.1120", "Requested Report Type (9233)", TypeContextGroupName, false)

	// CTTransversePlaneReferenceBasis1000 CT Transverse Plane Reference Basis (1000)
	CTTransversePlaneReferenceBasis1000 = New("1.2.840.10008.6.1.1121", "CT Transverse Plane Reference Basis (1000)", TypeContextGroupName, false)

	// AnatomicalReferenceBasis1001 Anatomical Reference Basis (1001)
	AnatomicalReferenceBasis1001 = New("1.2.840.10008.6.1.1122", "Anatomical Reference Basis (1001)", TypeContextGroupName, false)

	// AnatomicalReferenceBasisHead1002 Anatomical Reference Basis - Head (1002)
	AnatomicalReferenceBasisHead1002 = New("1.2.840.10008.6.1.1123", "Anatomical Reference Basis - Head (1002)", TypeContextGroupName, false)

	// AnatomicalReferenceBasisSpine1003 Anatomical Reference Basis - Spine (1003)
	AnatomicalReferenceBasisSpine1003 = New("1.2.840.10008.6.1.1124", "Anatomical Reference Basis - Spine (1003)", TypeContextGroupName, false)

	// AnatomicalReferenceBasisChest1004 Anatomical Reference Basis - Chest (1004)
	AnatomicalReferenceBasisChest1004 = New("1.2.840.10008.6.1.1125", "Anatomical Reference Basis - Chest (1004)", TypeContextGroupName, false)

	// AnatomicalReferenceBasisAbdomenPelvis1005 Anatomical Reference Basis - Abdomen/Pelvis (1005)
	AnatomicalReferenceBasisAbdomenPelvis1005 = New("1.2.840.10008.6.1.1126", "Anatomical Reference Basis - Abdomen/Pelvis (1005)", TypeContextGroupName, false)

	// AnatomicalReferenceBasisExtremity1006 Anatomical Reference Basis - Extremity (1006)
	AnatomicalReferenceBasisExtremity1006 = New("1.2.840.10008.6.1.1127", "Anatomical Reference Basis - Extremity (1006)", TypeContextGroupName, false)

	// ReferenceGeometryPlane1010 Reference Geometry - Plane (1010)
	ReferenceGeometryPlane1010 = New("1.2.840.10008.6.1.1128", "Reference Geometry - Plane (1010)", TypeContextGroupName, false)

	// ReferenceGeometryPoint1011 Reference Geometry - Point (1011)
	ReferenceGeometryPoint1011 = New("1.2.840.10008.6.1.1129", "Reference Geometry - Point (1011)", TypeContextGroupName, false)

	// PatientAlignmentMethod1015 Patient Alignment Method (1015)
	PatientAlignmentMethod1015 = New("1.2.840.10008.6.1.1130", "Patient Alignment Method (1015)", TypeContextGroupName, false)

	// ContraindicationsForCTImaging1200 Contraindications For CT Imaging (1200)
	ContraindicationsForCTImaging1200 = New("1.2.840.10008.6.1.1131", "Contraindications For CT Imaging (1200)", TypeContextGroupName, false)

	// FiducialCategory7110 Fiducial Category (7110)
	FiducialCategory7110 = New("1.2.840.10008.6.1.1132", "Fiducial Category (7110)", TypeContextGroupName, false)

	// Fiducial7111 Fiducial (7111)
	Fiducial7111 = New("1.2.840.10008.6.1.1133", "Fiducial (7111)", TypeContextGroupName, false)

	// NonImageSourceInstancePurposeOfReference7013 Non-Image Source Instance Purpose of Reference (7013)
	NonImageSourceInstancePurposeOfReference7013 = New("1.2.840.10008.6.1.1134", "Non-Image Source Instance Purpose of Reference (7013)", TypeContextGroupName, false)

	// RTProcessOutput7023 RT Process Output (7023)
	RTProcessOutput7023 = New("1.2.840.10008.6.1.1135", "RT Process Output (7023)", TypeContextGroupName, false)

	// RTProcessInput7024 RT Process Input (7024)
	RTProcessInput7024 = New("1.2.840.10008.6.1.1136", "RT Process Input (7024)", TypeContextGroupName, false)

	// RTProcessInputUsed7025 RT Process Input Used (7025)
	RTProcessInputUsed7025 = New("1.2.840.10008.6.1.1137", "RT Process Input Used (7025)", TypeContextGroupName, false)

	// ProstateAnatomy6300 Prostate Anatomy (6300)
	ProstateAnatomy6300 = New("1.2.840.10008.6.1.1138", "Prostate Anatomy (6300)", TypeContextGroupName, false)

	// ProstateSectorAnatomyFromPIRADSV26301 Prostate Sector Anatomy from PI-RADS v2 (6301)
	ProstateSectorAnatomyFromPIRADSV26301 = New("1.2.840.10008.6.1.1139", "Prostate Sector Anatomy from PI-RADS v2 (6301)", TypeContextGroupName, false)

	// ProstateSectorAnatomyFromEuropeanConcensus16SectorMinimalModel6302 Prostate Sector Anatomy from European Concensus 16 Sector (Minimal) Model (6302)
	ProstateSectorAnatomyFromEuropeanConcensus16SectorMinimalModel6302 = New("1.2.840.10008.6.1.1140", "Prostate Sector Anatomy from European Concensus 16 Sector (Minimal) Model (6302)", TypeContextGroupName, false)

	// ProstateSectorAnatomyFromEuropeanConcensus27SectorOptimalModel6303 Prostate Sector Anatomy from European Concensus 27 Sector (Optimal) Model (6303)
	ProstateSectorAnatomyFromEuropeanConcensus27SectorOptimalModel6303 = New("1.2.840.10008.6.1.1141", "Prostate Sector Anatomy from European Concensus 27 Sector (Optimal) Model (6303)", TypeContextGroupName, false)

	// MeasurementSelectionReason12301 Measurement Selection Reason (12301)
	MeasurementSelectionReason12301 = New("1.2.840.10008.6.1.1142", "Measurement Selection Reason (12301)", TypeContextGroupName, false)

	// EchoFindingObservationType12302 Echo Finding Observation Type (12302)
	EchoFindingObservationType12302 = New("1.2.840.10008.6.1.1143", "Echo Finding Observation Type (12302)", TypeContextGroupName, false)

	// EchoMeasurementType12303 Echo Measurement Type (12303)
	EchoMeasurementType12303 = New("1.2.840.10008.6.1.1144", "Echo Measurement Type (12303)", TypeContextGroupName, false)

	// CardiovascularMeasuredProperty12304 Cardiovascular Measured Property (12304)
	CardiovascularMeasuredProperty12304 = New("1.2.840.10008.6.1.1145", "Cardiovascular Measured Property (12304)", TypeContextGroupName, false)

	// BasicEchoAnatomicSite12305 Basic Echo Anatomic Site (12305)
	BasicEchoAnatomicSite12305 = New("1.2.840.10008.6.1.1146", "Basic Echo Anatomic Site (12305)", TypeContextGroupName, false)

	// EchoFlowDirection12306 Echo Flow Direction (12306)
	EchoFlowDirection12306 = New("1.2.840.10008.6.1.1147", "Echo Flow Direction (12306)", TypeContextGroupName, false)

	// CardiacPhaseAndTimePoint12307 Cardiac Phase and Time Point (12307)
	CardiacPhaseAndTimePoint12307 = New("1.2.840.10008.6.1.1148", "Cardiac Phase and Time Point (12307)", TypeContextGroupName, false)

	// CoreEchoMeasurement12300 Core Echo Measurement (12300)
	CoreEchoMeasurement12300 = New("1.2.840.10008.6.1.1149", "Core Echo Measurement (12300)", TypeContextGroupName, false)

	// OCTAProcessingAlgorithmFamily4270 OCT-A Processing Algorithm Family (4270)
	OCTAProcessingAlgorithmFamily4270 = New("1.2.840.10008.6.1.1150", "OCT-A Processing Algorithm Family (4270)", TypeContextGroupName, false)

	// EnFaceImageType4271 En Face Image Type (4271)
	EnFaceImageType4271 = New("1.2.840.10008.6.1.1151", "En Face Image Type (4271)", TypeContextGroupName, false)

	// OPTScanPatternType4272 OPT Scan Pattern Type (4272)
	OPTScanPatternType4272 = New("1.2.840.10008.6.1.1152", "OPT Scan Pattern Type (4272)", TypeContextGroupName, false)

	// RetinalSegmentationSurface4273 Retinal Segmentation Surface (4273)
	RetinalSegmentationSurface4273 = New("1.2.840.10008.6.1.1153", "Retinal Segmentation Surface (4273)", TypeContextGroupName, false)

	// OrganForRadiationDoseEstimate10060 Organ for Radiation Dose Estimate (10060)
	OrganForRadiationDoseEstimate10060 = New("1.2.840.10008.6.1.1154", "Organ for Radiation Dose Estimate (10060)", TypeContextGroupName, false)

	// AbsorbedRadiationDoseType10061 Absorbed Radiation Dose Type (10061)
	AbsorbedRadiationDoseType10061 = New("1.2.840.10008.6.1.1155", "Absorbed Radiation Dose Type (10061)", TypeContextGroupName, false)

	// EquivalentRadiationDoseType10062 Equivalent Radiation Dose Type (10062)
	EquivalentRadiationDoseType10062 = New("1.2.840.10008.6.1.1156", "Equivalent Radiation Dose Type (10062)", TypeContextGroupName, false)

	// RadiationDoseEstimateDistributionRepresentation10063 Radiation Dose Estimate Distribution Representation (10063)
	RadiationDoseEstimateDistributionRepresentation10063 = New("1.2.840.10008.6.1.1157", "Radiation Dose Estimate Distribution Representation (10063)", TypeContextGroupName, false)

	// PatientModelType10064 Patient Model Type (10064)
	PatientModelType10064 = New("1.2.840.10008.6.1.1158", "Patient Model Type (10064)", TypeContextGroupName, false)

	// RadiationTransportModelType10065 Radiation Transport Model Type (10065)
	RadiationTransportModelType10065 = New("1.2.840.10008.6.1.1159", "Radiation Transport Model Type (10065)", TypeContextGroupName, false)

	// AttenuatorCategory10066 Attenuator Category (10066)
	AttenuatorCategory10066 = New("1.2.840.10008.6.1.1160", "Attenuator Category (10066)", TypeContextGroupName, false)

	// RadiationAttenuatorMaterial10067 Radiation Attenuator Material (10067)
	RadiationAttenuatorMaterial10067 = New("1.2.840.10008.6.1.1161", "Radiation Attenuator Material (10067)", TypeContextGroupName, false)

	// EstimateMethodType10068 Estimate Method Type (10068)
	EstimateMethodType10068 = New("1.2.840.10008.6.1.1162", "Estimate Method Type (10068)", TypeContextGroupName, false)

	// RadiationDoseEstimateParameter10069 Radiation Dose Estimate Parameter (10069)
	RadiationDoseEstimateParameter10069 = New("1.2.840.10008.6.1.1163", "Radiation Dose Estimate Parameter (10069)", TypeContextGroupName, false)

	// RadiationDoseType10070 Radiation Dose Type (10070)
	RadiationDoseType10070 = New("1.2.840.10008.6.1.1164", "Radiation Dose Type (10070)", TypeContextGroupName, false)

	// MRDiffusionComponentSemantic7270 MR Diffusion Component Semantic (7270)
	MRDiffusionComponentSemantic7270 = New("1.2.840.10008.6.1.1165", "MR Diffusion Component Semantic (7270)", TypeContextGroupName, false)

	// MRDiffusionAnisotropyIndex7271 MR Diffusion Anisotropy Index (7271)
	MRDiffusionAnisotropyIndex7271 = New("1.2.840.10008.6.1.1166", "MR Diffusion Anisotropy Index (7271)", TypeContextGroupName, false)

	// MRDiffusionModelParameter7272 MR Diffusion Model Parameter (7272)
	MRDiffusionModelParameter7272 = New("1.2.840.10008.6.1.1167", "MR Diffusion Model Parameter (7272)", TypeContextGroupName, false)

	// MRDiffusionModel7273 MR Diffusion Model (7273)
	MRDiffusionModel7273 = New("1.2.840.10008.6.1.1168", "MR Diffusion Model (7273)", TypeContextGroupName, false)

	// MRDiffusionModelFittingMethod7274 MR Diffusion Model Fitting Method (7274)
	MRDiffusionModelFittingMethod7274 = New("1.2.840.10008.6.1.1169", "MR Diffusion Model Fitting Method (7274)", TypeContextGroupName, false)

	// MRDiffusionModelSpecificMethod7275 MR Diffusion Model Specific Method (7275)
	MRDiffusionModelSpecificMethod7275 = New("1.2.840.10008.6.1.1170", "MR Diffusion Model Specific Method (7275)", TypeContextGroupName, false)

	// MRDiffusionModelInput7276 MR Diffusion Model Input (7276)
	MRDiffusionModelInput7276 = New("1.2.840.10008.6.1.1171", "MR Diffusion Model Input (7276)", TypeContextGroupName, false)

	// DiffusionRateAreaOverTimeUnit7277 Diffusion Rate Area Over Time Unit (7277)
	DiffusionRateAreaOverTimeUnit7277 = New("1.2.840.10008.6.1.1172", "Diffusion Rate Area Over Time Unit (7277)", TypeContextGroupName, false)

	// PediatricSizeCategory7039 Pediatric Size Category (7039)
	PediatricSizeCategory7039 = New("1.2.840.10008.6.1.1173", "Pediatric Size Category (7039)", TypeContextGroupName, false)

	// CalciumScoringPatientSizeCategory7041 Calcium Scoring Patient Size Category (7041)
	CalciumScoringPatientSizeCategory7041 = New("1.2.840.10008.6.1.1174", "Calcium Scoring Patient Size Category (7041)", TypeContextGroupName, false)

	// ReasonForRepeatingAcquisition10034 Reason for Repeating Acquisition (10034)
	ReasonForRepeatingAcquisition10034 = New("1.2.840.10008.6.1.1175", "Reason for Repeating Acquisition (10034)", TypeContextGroupName, false)

	// ProtocolAssertion800 Protocol Assertion (800)
	ProtocolAssertion800 = New("1.2.840.10008.6.1.1176", "Protocol Assertion (800)", TypeContextGroupName, false)

	// RadiotherapeuticDoseMeasurementDevice7026 Radiotherapeutic Dose Measurement Device (7026)
	RadiotherapeuticDoseMeasurementDevice7026 = New("1.2.840.10008.6.1.1177", "Radiotherapeutic Dose Measurement Device (7026)", TypeContextGroupName, false)

	// ExportAdditionalInformationDocumentTitle7014 Export Additional Information Document Title (7014)
	ExportAdditionalInformationDocumentTitle7014 = New("1.2.840.10008.6.1.1178", "Export Additional Information Document Title (7014)", TypeContextGroupName, false)

	// ExportDelayReason7015 Export Delay Reason (7015)
	ExportDelayReason7015 = New("1.2.840.10008.6.1.1179", "Export Delay Reason (7015)", TypeContextGroupName, false)

	// LevelOfDifficulty7016 Level of Difficulty (7016)
	LevelOfDifficulty7016 = New("1.2.840.10008.6.1.1180", "Level of Difficulty (7016)", TypeContextGroupName, false)

	// CategoryOfTeachingMaterialImaging7017 Category of Teaching Material - Imaging (7017)
	CategoryOfTeachingMaterialImaging7017 = New("1.2.840.10008.6.1.1181", "Category of Teaching Material - Imaging (7017)", TypeContextGroupName, false)

	// MiscellaneousDocumentTitle7018 Miscellaneous Document Title (7018)
	MiscellaneousDocumentTitle7018 = New("1.2.840.10008.6.1.1182", "Miscellaneous Document Title (7018)", TypeContextGroupName, false)

	// SegmentationNonImageSourcePurposeOfReference7019 Segmentation Non-Image Source Purpose of Reference (7019)
	SegmentationNonImageSourcePurposeOfReference7019 = New("1.2.840.10008.6.1.1183", "Segmentation Non-Image Source Purpose of Reference (7019)", TypeContextGroupName, false)

	// LongitudinalTemporalEventType280 Longitudinal Temporal Event Type (280)
	LongitudinalTemporalEventType280 = New("1.2.840.10008.6.1.1184", "Longitudinal Temporal Event Type (280)", TypeContextGroupName, false)

	// NonLesionObjectTypePhysicalObject6401 Non-lesion Object Type - Physical Object (6401)
	NonLesionObjectTypePhysicalObject6401 = New("1.2.840.10008.6.1.1185", "Non-lesion Object Type - Physical Object (6401)", TypeContextGroupName, false)

	// NonLesionObjectTypeSubstance6402 Non-lesion Object Type - Substance (6402)
	NonLesionObjectTypeSubstance6402 = New("1.2.840.10008.6.1.1186", "Non-lesion Object Type - Substance (6402)", TypeContextGroupName, false)

	// NonLesionObjectTypeTissue6403 Non-lesion Object Type - Tissue (6403)
	NonLesionObjectTypeTissue6403 = New("1.2.840.10008.6.1.1187", "Non-lesion Object Type - Tissue (6403)", TypeContextGroupName, false)

	// ChestNonLesionObjectTypePhysicalObject6404 Chest Non-lesion Object Type - Physical Object (6404)
	ChestNonLesionObjectTypePhysicalObject6404 = New("1.2.840.10008.6.1.1188", "Chest Non-lesion Object Type - Physical Object (6404)", TypeContextGroupName, false)

	// ChestNonLesionObjectTypeTissue6405 Chest Non-lesion Object Type - Tissue (6405)
	ChestNonLesionObjectTypeTissue6405 = New("1.2.840.10008.6.1.1189", "Chest Non-lesion Object Type - Tissue (6405)", TypeContextGroupName, false)

	// TissueSegmentationPropertyType7191 Tissue Segmentation Property Type (7191)
	TissueSegmentationPropertyType7191 = New("1.2.840.10008.6.1.1190", "Tissue Segmentation Property Type (7191)", TypeContextGroupName, false)

	// AnatomicalStructureSegmentationPropertyType7192 Anatomical Structure Segmentation Property Type (7192)
	AnatomicalStructureSegmentationPropertyType7192 = New("1.2.840.10008.6.1.1191", "Anatomical Structure Segmentation Property Type (7192)", TypeContextGroupName, false)

	// PhysicalObjectSegmentationPropertyType7193 Physical Object Segmentation Property Type (7193)
	PhysicalObjectSegmentationPropertyType7193 = New("1.2.840.10008.6.1.1192", "Physical Object Segmentation Property Type (7193)", TypeContextGroupName, false)

	// MorphologicallyAbnormalStructureSegmentationPropertyType7194 Morphologically Abnormal Structure Segmentation Property Type (7194)
	MorphologicallyAbnormalStructureSegmentationPropertyType7194 = New("1.2.840.10008.6.1.1193", "Morphologically Abnormal Structure Segmentation Property Type (7194)", TypeContextGroupName, false)

	// FunctionSegmentationPropertyType7195 Function Segmentation Property Type (7195)
	FunctionSegmentationPropertyType7195 = New("1.2.840.10008.6.1.1194", "Function Segmentation Property Type (7195)", TypeContextGroupName, false)

	// SpatialAndRelationalConceptSegmentationPropertyType7196 Spatial and Relational Concept Segmentation Property Type (7196)
	SpatialAndRelationalConceptSegmentationPropertyType7196 = New("1.2.840.10008.6.1.1195", "Spatial and Relational Concept Segmentation Property Type (7196)", TypeContextGroupName, false)

	// BodySubstanceSegmentationPropertyType7197 Body Substance Segmentation Property Type (7197)
	BodySubstanceSegmentationPropertyType7197 = New("1.2.840.10008.6.1.1196", "Body Substance Segmentation Property Type (7197)", TypeContextGroupName, false)

	// SubstanceSegmentationPropertyType7198 Substance Segmentation Property Type (7198)
	SubstanceSegmentationPropertyType7198 = New("1.2.840.10008.6.1.1197", "Substance Segmentation Property Type (7198)", TypeContextGroupName, false)

	// InterpretationRequestDiscontinuationReason9303 Interpretation Request Discontinuation Reason (9303)
	InterpretationRequestDiscontinuationReason9303 = New("1.2.840.10008.6.1.1198", "Interpretation Request Discontinuation Reason (9303)", TypeContextGroupName, false)

	// GrayLevelRunLengthBasedFeature7475 Gray Level Run Length Based Feature (7475)
	GrayLevelRunLengthBasedFeature7475 = New("1.2.840.10008.6.1.1199", "Gray Level Run Length Based Feature (7475)", TypeContextGroupName, false)

	// GrayLevelSizeZoneBasedFeature7476 Gray Level Size Zone Based Feature (7476)
	GrayLevelSizeZoneBasedFeature7476 = New("1.2.840.10008.6.1.1200", "Gray Level Size Zone Based Feature (7476)", TypeContextGroupName, false)

	// EncapsulatedDocumentSourcePurposeOfReference7060 Encapsulated Document Source Purpose of Reference (7060)
	EncapsulatedDocumentSourcePurposeOfReference7060 = New("1.2.840.10008.6.1.1201", "Encapsulated Document Source Purpose of Reference (7060)", TypeContextGroupName, false)

	// ModelDocumentTitle7061 Model Document Title (7061)
	ModelDocumentTitle7061 = New("1.2.840.10008.6.1.1202", "Model Document Title (7061)", TypeContextGroupName, false)

	// PurposeOfReferenceToPredecessor3DModel7062 Purpose of Reference to Predecessor 3D Model (7062)
	PurposeOfReferenceToPredecessor3DModel7062 = New("1.2.840.10008.6.1.1203", "Purpose of Reference to Predecessor 3D Model (7062)", TypeContextGroupName, false)

	// ModelScaleUnit7063 Model Scale Unit (7063)
	ModelScaleUnit7063 = New("1.2.840.10008.6.1.1204", "Model Scale Unit (7063)", TypeContextGroupName, false)

	// ModelUsage7064 Model Usage (7064)
	ModelUsage7064 = New("1.2.840.10008.6.1.1205", "Model Usage (7064)", TypeContextGroupName, false)

	// RadiationDoseUnit10071 Radiation Dose Unit (10071)
	RadiationDoseUnit10071 = New("1.2.840.10008.6.1.1206", "Radiation Dose Unit (10071)", TypeContextGroupName, false)

	// RadiotherapyFiducial7112 Radiotherapy Fiducial (7112)
	RadiotherapyFiducial7112 = New("1.2.840.10008.6.1.1207", "Radiotherapy Fiducial (7112)", TypeContextGroupName, false)

	// MultiEnergyRelevantMaterial300 Multi-energy Relevant Material (300)
	MultiEnergyRelevantMaterial300 = New("1.2.840.10008.6.1.1208", "Multi-energy Relevant Material (300)", TypeContextGroupName, false)

	// MultiEnergyMaterialUnit301 Multi-energy Material Unit (301)
	MultiEnergyMaterialUnit301 = New("1.2.840.10008.6.1.1209", "Multi-energy Material Unit (301)", TypeContextGroupName, false)

	// DosimetricObjectiveType9500 Dosimetric Objective Type (9500)
	DosimetricObjectiveType9500 = New("1.2.840.10008.6.1.1210", "Dosimetric Objective Type (9500)", TypeContextGroupName, false)

	// PrescriptionAnatomyCategory9501 Prescription Anatomy Category (9501)
	PrescriptionAnatomyCategory9501 = New("1.2.840.10008.6.1.1211", "Prescription Anatomy Category (9501)", TypeContextGroupName, false)

	// RTSegmentAnnotationCategory9502 RT Segment Annotation Category (9502)
	RTSegmentAnnotationCategory9502 = New("1.2.840.10008.6.1.1212", "RT Segment Annotation Category (9502)", TypeContextGroupName, false)

	// RadiotherapyTherapeuticRoleCategory9503 Radiotherapy Therapeutic Role Category (9503)
	RadiotherapyTherapeuticRoleCategory9503 = New("1.2.840.10008.6.1.1213", "Radiotherapy Therapeutic Role Category (9503)", TypeContextGroupName, false)

	// RTGeometricInformation9504 RT Geometric Information (9504)
	RTGeometricInformation9504 = New("1.2.840.10008.6.1.1214", "RT Geometric Information (9504)", TypeContextGroupName, false)

	// FixationOrPositioningDevice9505 Fixation or Positioning Device (9505)
	FixationOrPositioningDevice9505 = New("1.2.840.10008.6.1.1215", "Fixation or Positioning Device (9505)", TypeContextGroupName, false)

	// BrachytherapyDevice9506 Brachytherapy Device (9506)
	BrachytherapyDevice9506 = New("1.2.840.10008.6.1.1216", "Brachytherapy Device (9506)", TypeContextGroupName, false)

	// ExternalBodyModel9507 External Body Model (9507)
	ExternalBodyModel9507 = New("1.2.840.10008.6.1.1217", "External Body Model (9507)", TypeContextGroupName, false)

	// NonSpecificVolume9508 Non-specific Volume (9508)
	NonSpecificVolume9508 = New("1.2.840.10008.6.1.1218", "Non-specific Volume (9508)", TypeContextGroupName, false)

	// PurposeOfReferenceForRTPhysicianIntentInput9509 Purpose of Reference For RT Physician Intent Input (9509)
	PurposeOfReferenceForRTPhysicianIntentInput9509 = New("1.2.840.10008.6.1.1219", "Purpose of Reference For RT Physician Intent Input (9509)", TypeContextGroupName, false)

	// PurposeOfReferenceForRTTreatmentPlanningInput9510 Purpose of Reference For RT Treatment Planning Input (9510)
	PurposeOfReferenceForRTTreatmentPlanningInput9510 = New("1.2.840.10008.6.1.1220", "Purpose of Reference For RT Treatment Planning Input (9510)", TypeContextGroupName, false)

	// GeneralExternalRadiotherapyProcedureTechnique9511 General External Radiotherapy Procedure Technique (9511)
	GeneralExternalRadiotherapyProcedureTechnique9511 = New("1.2.840.10008.6.1.1221", "General External Radiotherapy Procedure Technique (9511)", TypeContextGroupName, false)

	// TomotherapeuticRadiotherapyProcedureTechnique9512 Tomotherapeutic Radiotherapy Procedure Technique (9512)
	TomotherapeuticRadiotherapyProcedureTechnique9512 = New("1.2.840.10008.6.1.1222", "Tomotherapeutic Radiotherapy Procedure Technique (9512)", TypeContextGroupName, false)

	// FixationDevice9513 Fixation Device (9513)
	FixationDevice9513 = New("1.2.840.10008.6.1.1223", "Fixation Device (9513)", TypeContextGroupName, false)

	// AnatomicalStructureForRadiotherapy9514 Anatomical Structure For Radiotherapy (9514)
	AnatomicalStructureForRadiotherapy9514 = New("1.2.840.10008.6.1.1224", "Anatomical Structure For Radiotherapy (9514)", TypeContextGroupName, false)

	// RTPatientSupportDevice9515 RT Patient Support Device (9515)
	RTPatientSupportDevice9515 = New("1.2.840.10008.6.1.1225", "RT Patient Support Device (9515)", TypeContextGroupName, false)

	// RadiotherapyBolusDeviceType9516 Radiotherapy Bolus Device Type (9516)
	RadiotherapyBolusDeviceType9516 = New("1.2.840.10008.6.1.1226", "Radiotherapy Bolus Device Type (9516)", TypeContextGroupName, false)

	// RadiotherapyBlockDeviceType9517 Radiotherapy Block Device Type (9517)
	RadiotherapyBlockDeviceType9517 = New("1.2.840.10008.6.1.1227", "Radiotherapy Block Device Type (9517)", TypeContextGroupName, false)

	// RadiotherapyAccessoryNoSlotHolderDeviceType9518 Radiotherapy Accessory No-slot Holder Device Type (9518)
	RadiotherapyAccessoryNoSlotHolderDeviceType9518 = New("1.2.840.10008.6.1.1228", "Radiotherapy Accessory No-slot Holder Device Type (9518)", TypeContextGroupName, false)

	// RadiotherapyAccessorySlotHolderDeviceType9519 Radiotherapy Accessory Slot Holder Device Type (9519)
	RadiotherapyAccessorySlotHolderDeviceType9519 = New("1.2.840.10008.6.1.1229", "Radiotherapy Accessory Slot Holder Device Type (9519)", TypeContextGroupName, false)

	// SegmentedRTAccessoryDevice9520 Segmented RT Accessory Device (9520)
	SegmentedRTAccessoryDevice9520 = New("1.2.840.10008.6.1.1230", "Segmented RT Accessory Device (9520)", TypeContextGroupName, false)

	// RadiotherapyTreatmentEnergyUnit9521 Radiotherapy Treatment Energy Unit (9521)
	RadiotherapyTreatmentEnergyUnit9521 = New("1.2.840.10008.6.1.1231", "Radiotherapy Treatment Energy Unit (9521)", TypeContextGroupName, false)

	// MultiSourceRadiotherapyProcedureTechnique9522 Multi-source Radiotherapy Procedure Technique (9522)
	MultiSourceRadiotherapyProcedureTechnique9522 = New("1.2.840.10008.6.1.1232", "Multi-source Radiotherapy Procedure Technique (9522)", TypeContextGroupName, false)

	// RoboticRadiotherapyProcedureTechnique9523 Robotic Radiotherapy Procedure Technique (9523)
	RoboticRadiotherapyProcedureTechnique9523 = New("1.2.840.10008.6.1.1233", "Robotic Radiotherapy Procedure Technique (9523)", TypeContextGroupName, false)

	// RadiotherapyProcedureTechnique9524 Radiotherapy Procedure Technique (9524)
	RadiotherapyProcedureTechnique9524 = New("1.2.840.10008.6.1.1234", "Radiotherapy Procedure Technique (9524)", TypeContextGroupName, false)

	// RadiationTherapyParticle9525 Radiation Therapy Particle (9525)
	RadiationTherapyParticle9525 = New("1.2.840.10008.6.1.1235", "Radiation Therapy Particle (9525)", TypeContextGroupName, false)

	// IonTherapyParticle9526 Ion Therapy Particle (9526)
	IonTherapyParticle9526 = New("1.2.840.10008.6.1.1236", "Ion Therapy Particle (9526)", TypeContextGroupName, false)

	// TeletherapyIsotope9527 Teletherapy Isotope (9527)
	TeletherapyIsotope9527 = New("1.2.840.10008.6.1.1237", "Teletherapy Isotope (9527)", TypeContextGroupName, false)

	// BrachytherapyIsotope9528 Brachytherapy Isotope (9528)
	BrachytherapyIsotope9528 = New("1.2.840.10008.6.1.1238", "Brachytherapy Isotope (9528)", TypeContextGroupName, false)

	// SingleDoseDosimetricObjective9529 Single Dose Dosimetric Objective (9529)
	SingleDoseDosimetricObjective9529 = New("1.2.840.10008.6.1.1239", "Single Dose Dosimetric Objective (9529)", TypeContextGroupName, false)

	// PercentageAndDoseDosimetricObjective9530 Percentage and Dose Dosimetric Objective (9530)
	PercentageAndDoseDosimetricObjective9530 = New("1.2.840.10008.6.1.1240", "Percentage and Dose Dosimetric Objective (9530)", TypeContextGroupName, false)

	// VolumeAndDoseDosimetricObjective9531 Volume and Dose Dosimetric Objective (9531)
	VolumeAndDoseDosimetricObjective9531 = New("1.2.840.10008.6.1.1241", "Volume and Dose Dosimetric Objective (9531)", TypeContextGroupName, false)

	// NoParameterDosimetricObjective9532 No-Parameter Dosimetric Objective (9532)
	NoParameterDosimetricObjective9532 = New("1.2.840.10008.6.1.1242", "No-Parameter Dosimetric Objective (9532)", TypeContextGroupName, false)

	// DeliveryTimeStructure9533 Delivery Time Structure (9533)
	DeliveryTimeStructure9533 = New("1.2.840.10008.6.1.1243", "Delivery Time Structure (9533)", TypeContextGroupName, false)

	// RadiotherapyTarget9534 Radiotherapy Target (9534)
	RadiotherapyTarget9534 = New("1.2.840.10008.6.1.1244", "Radiotherapy Target (9534)", TypeContextGroupName, false)

	// RadiotherapyDoseCalculationRole9535 Radiotherapy Dose Calculation Role (9535)
	RadiotherapyDoseCalculationRole9535 = New("1.2.840.10008.6.1.1245", "Radiotherapy Dose Calculation Role (9535)", TypeContextGroupName, false)

	// RadiotherapyPrescribingAndSegmentingPersonRole9536 Radiotherapy Prescribing and Segmenting Person Role (9536)
	RadiotherapyPrescribingAndSegmentingPersonRole9536 = New("1.2.840.10008.6.1.1246", "Radiotherapy Prescribing and Segmenting Person Role (9536)", TypeContextGroupName, false)

	// EffectiveDoseCalculationMethodCategory9537 Effective Dose Calculation Method Category (9537)
	EffectiveDoseCalculationMethodCategory9537 = New("1.2.840.10008.6.1.1247", "Effective Dose Calculation Method Category (9537)", TypeContextGroupName, false)

	// RadiationTransportBasedEffectiveDoseMethodModifier9538 Radiation Transport-based Effective Dose Method Modifier (9538)
	RadiationTransportBasedEffectiveDoseMethodModifier9538 = New("1.2.840.10008.6.1.1248", "Radiation Transport-based Effective Dose Method Modifier (9538)", TypeContextGroupName, false)

	// FractionationBasedEffectiveDoseMethodModifier9539 Fractionation-based Effective Dose Method Modifier (9539)
	FractionationBasedEffectiveDoseMethodModifier9539 = New("1.2.840.10008.6.1.1249", "Fractionation-based Effective Dose Method Modifier (9539)", TypeContextGroupName, false)

	// ImagingAgentAdministrationAdverseEvent60 Imaging Agent Administration Adverse Event (60)
	ImagingAgentAdministrationAdverseEvent60 = New("1.2.840.10008.6.1.1250", "Imaging Agent Administration Adverse Event (60)", TypeContextGroupName, false)

	// TimeRelativeToProcedure61RETIRED Time Relative to Procedure (Retired) (61)
	TimeRelativeToProcedure61RETIRED = New("1.2.840.10008.6.1.1251", "Time Relative to Procedure (Retired) (61)", TypeContextGroupName, true)

	// ImagingAgentAdministrationPhaseType62 Imaging Agent Administration Phase Type (62)
	ImagingAgentAdministrationPhaseType62 = New("1.2.840.10008.6.1.1252", "Imaging Agent Administration Phase Type (62)", TypeContextGroupName, false)

	// ImagingAgentAdministrationMode63 Imaging Agent Administration Mode (63)
	ImagingAgentAdministrationMode63 = New("1.2.840.10008.6.1.1253", "Imaging Agent Administration Mode (63)", TypeContextGroupName, false)

	// ImagingAgentAdministrationPatientState64 Imaging Agent Administration Patient State (64)
	ImagingAgentAdministrationPatientState64 = New("1.2.840.10008.6.1.1254", "Imaging Agent Administration Patient State (64)", TypeContextGroupName, false)

	// ImagingAgentAdministrationPremedication65 Imaging Agent Administration Premedication (65)
	ImagingAgentAdministrationPremedication65 = New("1.2.840.10008.6.1.1255", "Imaging Agent Administration Premedication (65)", TypeContextGroupName, false)

	// ImagingAgentAdministrationMedication66 Imaging Agent Administration Medication (66)
	ImagingAgentAdministrationMedication66 = New("1.2.840.10008.6.1.1256", "Imaging Agent Administration Medication (66)", TypeContextGroupName, false)

	// ImagingAgentAdministrationCompletionStatus67 Imaging Agent Administration Completion Status (67)
	ImagingAgentAdministrationCompletionStatus67 = New("1.2.840.10008.6.1.1257", "Imaging Agent Administration Completion Status (67)", TypeContextGroupName, false)

	// ImagingAgentAdministrationPharmaceuticalPresentationUnit68 Imaging Agent Administration Pharmaceutical Presentation Unit (68)
	ImagingAgentAdministrationPharmaceuticalPresentationUnit68 = New("1.2.840.10008.6.1.1258", "Imaging Agent Administration Pharmaceutical Presentation Unit (68)", TypeContextGroupName, false)

	// ImagingAgentAdministrationConsumable69 Imaging Agent Administration Consumable (69)
	ImagingAgentAdministrationConsumable69 = New("1.2.840.10008.6.1.1259", "Imaging Agent Administration Consumable (69)", TypeContextGroupName, false)

	// Flush70 Flush (70)
	Flush70 = New("1.2.840.10008.6.1.1260", "Flush (70)", TypeContextGroupName, false)

	// ImagingAgentAdministrationInjectorEventType71 Imaging Agent Administration Injector Event Type (71)
	ImagingAgentAdministrationInjectorEventType71 = New("1.2.840.10008.6.1.1261", "Imaging Agent Administration Injector Event Type (71)", TypeContextGroupName, false)

	// ImagingAgentAdministrationStepType72 Imaging Agent Administration Step Type (72)
	ImagingAgentAdministrationStepType72 = New("1.2.840.10008.6.1.1262", "Imaging Agent Administration Step Type (72)", TypeContextGroupName, false)

	// BolusShapingCurve73 Bolus Shaping Curve (73)
	BolusShapingCurve73 = New("1.2.840.10008.6.1.1263", "Bolus Shaping Curve (73)", TypeContextGroupName, false)

	// ImagingAgentAdministrationConsumableCatheterType74 Imaging Agent Administration Consumable Catheter Type (74)
	ImagingAgentAdministrationConsumableCatheterType74 = New("1.2.840.10008.6.1.1264", "Imaging Agent Administration Consumable Catheter Type (74)", TypeContextGroupName, false)

	// LowHighOrEqual75 Low High or Equal (75)
	LowHighOrEqual75 = New("1.2.840.10008.6.1.1265", "Low High or Equal (75)", TypeContextGroupName, false)

	// PremedicationType76 Premedication Type (76)
	PremedicationType76 = New("1.2.840.10008.6.1.1266", "Premedication Type (76)", TypeContextGroupName, false)

	// LateralityWithMedian245 Laterality with Median (245)
	LateralityWithMedian245 = New("1.2.840.10008.6.1.1267", "Laterality with Median (245)", TypeContextGroupName, false)

	// DermatologyAnatomicSite4029 Dermatology Anatomic Site (4029)
	DermatologyAnatomicSite4029 = New("1.2.840.10008.6.1.1268", "Dermatology Anatomic Site (4029)", TypeContextGroupName, false)

	// QuantitativeImageFeature218 Quantitative Image Feature (218)
	QuantitativeImageFeature218 = New("1.2.840.10008.6.1.1269", "Quantitative Image Feature (218)", TypeContextGroupName, false)

	// GlobalShapeDescriptor7477 Global Shape Descriptor (7477)
	GlobalShapeDescriptor7477 = New("1.2.840.10008.6.1.1270", "Global Shape Descriptor (7477)", TypeContextGroupName, false)

	// IntensityHistogramFeature7478 Intensity Histogram Feature (7478)
	IntensityHistogramFeature7478 = New("1.2.840.10008.6.1.1271", "Intensity Histogram Feature (7478)", TypeContextGroupName, false)

	// GreyLevelDistanceZoneBasedFeature7479 Grey Level Distance Zone Based Feature (7479)
	GreyLevelDistanceZoneBasedFeature7479 = New("1.2.840.10008.6.1.1272", "Grey Level Distance Zone Based Feature (7479)", TypeContextGroupName, false)

	// NeighbourhoodGreyToneDifferenceBasedFeature7500 Neighbourhood Grey Tone Difference Based Feature (7500)
	NeighbourhoodGreyToneDifferenceBasedFeature7500 = New("1.2.840.10008.6.1.1273", "Neighbourhood Grey Tone Difference Based Feature (7500)", TypeContextGroupName, false)

	// NeighbouringGreyLevelDependenceBasedFeature7501 Neighbouring Grey Level Dependence Based Feature (7501)
	NeighbouringGreyLevelDependenceBasedFeature7501 = New("1.2.840.10008.6.1.1274", "Neighbouring Grey Level Dependence Based Feature (7501)", TypeContextGroupName, false)

	// CorneaMeasurementMethodDescriptor4242 Cornea Measurement Method Descriptor (4242)
	CorneaMeasurementMethodDescriptor4242 = New("1.2.840.10008.6.1.1275", "Cornea Measurement Method Descriptor (4242)", TypeContextGroupName, false)

	// SegmentedRadiotherapeuticDoseMeasurementDevice7027 Segmented Radiotherapeutic Dose Measurement Device (7027)
	SegmentedRadiotherapeuticDoseMeasurementDevice7027 = New("1.2.840.10008.6.1.1276", "Segmented Radiotherapeutic Dose Measurement Device (7027)", TypeContextGroupName, false)

	// ClinicalCourseOfDisease6098 Clinical Course of Disease (6098)
	ClinicalCourseOfDisease6098 = New("1.2.840.10008.6.1.1277", "Clinical Course of Disease (6098)", TypeContextGroupName, false)

	// RacialGroup6099 Racial Group (6099)
	RacialGroup6099 = New("1.2.840.10008.6.1.1278", "Racial Group (6099)", TypeContextGroupName, false)

	// RelativeLaterality246 Relative Laterality (246)
	RelativeLaterality246 = New("1.2.840.10008.6.1.1279", "Relative Laterality (246)", TypeContextGroupName, false)

	// BrainLesionSegmentationTypeWithNecrosis7168 Brain Lesion Segmentation Type With Necrosis (7168)
	BrainLesionSegmentationTypeWithNecrosis7168 = New("1.2.840.10008.6.1.1280", "Brain Lesion Segmentation Type With Necrosis (7168)", TypeContextGroupName, false)

	// BrainLesionSegmentationTypeWithoutNecrosis7169 Brain Lesion Segmentation Type Without Necrosis (7169)
	BrainLesionSegmentationTypeWithoutNecrosis7169 = New("1.2.840.10008.6.1.1281", "Brain Lesion Segmentation Type Without Necrosis (7169)", TypeContextGroupName, false)

	// NonAcquisitionModality32 Non-Acquisition Modality (32)
	NonAcquisitionModality32 = New("1.2.840.10008.6.1.1282", "Non-Acquisition Modality (32)", TypeContextGroupName, false)

	// Modality33 Modality (33)
	Modality33 = New("1.2.840.10008.6.1.1283", "Modality (33)", TypeContextGroupName, false)

	// LateralityLeftRightOnly247 Laterality Left-Right Only (247)
	LateralityLeftRightOnly247 = New("1.2.840.10008.6.1.1284", "Laterality Left-Right Only (247)", TypeContextGroupName, false)

	// QualitativeEvaluationModifierType210 Qualitative Evaluation Modifier Type (210)
	QualitativeEvaluationModifierType210 = New("1.2.840.10008.6.1.1285", "Qualitative Evaluation Modifier Type (210)", TypeContextGroupName, false)

	// QualitativeEvaluationModifierValue211 Qualitative Evaluation Modifier Value (211)
	QualitativeEvaluationModifierValue211 = New("1.2.840.10008.6.1.1286", "Qualitative Evaluation Modifier Value (211)", TypeContextGroupName, false)

	// GenericAnatomicLocationModifier212 Generic Anatomic Location Modifier (212)
	GenericAnatomicLocationModifier212 = New("1.2.840.10008.6.1.1287", "Generic Anatomic Location Modifier (212)", TypeContextGroupName, false)

	// BeamLimitingDeviceType9541 Beam Limiting Device Type (9541)
	BeamLimitingDeviceType9541 = New("1.2.840.10008.6.1.1288", "Beam Limiting Device Type (9541)", TypeContextGroupName, false)

	// CompensatorDeviceType9542 Compensator Device Type (9542)
	CompensatorDeviceType9542 = New("1.2.840.10008.6.1.1289", "Compensator Device Type (9542)", TypeContextGroupName, false)

	// RadiotherapyTreatmentMachineMode9543 Radiotherapy Treatment Machine Mode (9543)
	RadiotherapyTreatmentMachineMode9543 = New("1.2.840.10008.6.1.1290", "Radiotherapy Treatment Machine Mode (9543)", TypeContextGroupName, false)

	// RadiotherapyDistanceReferenceLocation9544 Radiotherapy Distance Reference Location (9544)
	RadiotherapyDistanceReferenceLocation9544 = New("1.2.840.10008.6.1.1291", "Radiotherapy Distance Reference Location (9544)", TypeContextGroupName, false)

	// FixedBeamLimitingDeviceType9545 Fixed Beam Limiting Device Type (9545)
	FixedBeamLimitingDeviceType9545 = New("1.2.840.10008.6.1.1292", "Fixed Beam Limiting Device Type (9545)", TypeContextGroupName, false)

	// RadiotherapyWedgeType9546 Radiotherapy Wedge Type (9546)
	RadiotherapyWedgeType9546 = New("1.2.840.10008.6.1.1293", "Radiotherapy Wedge Type (9546)", TypeContextGroupName, false)

	// RTBeamLimitingDeviceOrientationLabel9547 RT Beam Limiting Device Orientation Label (9547)
	RTBeamLimitingDeviceOrientationLabel9547 = New("1.2.840.10008.6.1.1294", "RT Beam Limiting Device Orientation Label (9547)", TypeContextGroupName, false)

	// GeneralAccessoryDeviceType9548 General Accessory Device Type (9548)
	GeneralAccessoryDeviceType9548 = New("1.2.840.10008.6.1.1295", "General Accessory Device Type (9548)", TypeContextGroupName, false)

	// RadiationGenerationModeType9549 Radiation Generation Mode Type (9549)
	RadiationGenerationModeType9549 = New("1.2.840.10008.6.1.1296", "Radiation Generation Mode Type (9549)", TypeContextGroupName, false)

	// CArmPhotonElectronDeliveryRateUnit9550 C-Arm Photon-Electron Delivery Rate Unit (9550)
	CArmPhotonElectronDeliveryRateUnit9550 = New("1.2.840.10008.6.1.1297", "C-Arm Photon-Electron Delivery Rate Unit (9550)", TypeContextGroupName, false)

	// TreatmentDeliveryDeviceType9551 Treatment Delivery Device Type (9551)
	TreatmentDeliveryDeviceType9551 = New("1.2.840.10008.6.1.1298", "Treatment Delivery Device Type (9551)", TypeContextGroupName, false)

	// CArmPhotonElectronDosimeterUnit9552 C-Arm Photon-Electron Dosimeter Unit (9552)
	CArmPhotonElectronDosimeterUnit9552 = New("1.2.840.10008.6.1.1299", "C-Arm Photon-Electron Dosimeter Unit (9552)", TypeContextGroupName, false)

	// TreatmentPoint9553 Treatment Point (9553)
	TreatmentPoint9553 = New("1.2.840.10008.6.1.1300", "Treatment Point (9553)", TypeContextGroupName, false)

	// EquipmentReferencePoint9554 Equipment Reference Point (9554)
	EquipmentReferencePoint9554 = New("1.2.840.10008.6.1.1301", "Equipment Reference Point (9554)", TypeContextGroupName, false)

	// RadiotherapyTreatmentPlanningPersonRole9555 Radiotherapy Treatment Planning Person Role (9555)
	RadiotherapyTreatmentPlanningPersonRole9555 = New("1.2.840.10008.6.1.1302", "Radiotherapy Treatment Planning Person Role (9555)", TypeContextGroupName, false)

	// RealTimeVideoRenditionTitle7070 Real Time Video Rendition Title (7070)
	RealTimeVideoRenditionTitle7070 = New("1.2.840.10008.6.1.1303", "Real Time Video Rendition Title (7070)", TypeContextGroupName, false)

	// GeometryGraphicalRepresentation219 Geometry Graphical Representation (219)
	GeometryGraphicalRepresentation219 = New("1.2.840.10008.6.1.1304", "Geometry Graphical Representation (219)", TypeContextGroupName, false)

	// VisualExplanation217 Visual Explanation (217)
	VisualExplanation217 = New("1.2.840.10008.6.1.1305", "Visual Explanation (217)", TypeContextGroupName, false)

	// ProstateSectorAnatomyFromPIRADSV216304 Prostate Sector Anatomy from PI-RADS v2.1 (6304)
	ProstateSectorAnatomyFromPIRADSV216304 = New("1.2.840.10008.6.1.1306", "Prostate Sector Anatomy from PI-RADS v2.1 (6304)", TypeContextGroupName, false)

	// RadiotherapyRoboticNodeSet9556 Radiotherapy Robotic Node Set (9556)
	RadiotherapyRoboticNodeSet9556 = New("1.2.840.10008.6.1.1307", "Radiotherapy Robotic Node Set (9556)", TypeContextGroupName, false)

	// TomotherapeuticDosimeterUnit9557 Tomotherapeutic Dosimeter Unit (9557)
	TomotherapeuticDosimeterUnit9557 = New("1.2.840.10008.6.1.1308", "Tomotherapeutic Dosimeter Unit (9557)", TypeContextGroupName, false)

	// TomotherapeuticDoseRateUnit9558 Tomotherapeutic Dose Rate Unit (9558)
	TomotherapeuticDoseRateUnit9558 = New("1.2.840.10008.6.1.1309", "Tomotherapeutic Dose Rate Unit (9558)", TypeContextGroupName, false)

	// RoboticDeliveryDeviceDosimeterUnit9559 Robotic Delivery Device Dosimeter Unit (9559)
	RoboticDeliveryDeviceDosimeterUnit9559 = New("1.2.840.10008.6.1.1310", "Robotic Delivery Device Dosimeter Unit (9559)", TypeContextGroupName, false)

	// RoboticDeliveryDeviceDoseRateUnit9560 Robotic Delivery Device Dose Rate Unit (9560)
	RoboticDeliveryDeviceDoseRateUnit9560 = New("1.2.840.10008.6.1.1311", "Robotic Delivery Device Dose Rate Unit (9560)", TypeContextGroupName, false)

	// AnatomicStructure8134 Anatomic Structure (8134)
	AnatomicStructure8134 = New("1.2.840.10008.6.1.1312", "Anatomic Structure (8134)", TypeContextGroupName, false)

	// MediastinumFindingOrFeature6148 Mediastinum Finding or Feature (6148)
	MediastinumFindingOrFeature6148 = New("1.2.840.10008.6.1.1313", "Mediastinum Finding or Feature (6148)", TypeContextGroupName, false)

	// MediastinumAnatomy6149 Mediastinum Anatomy (6149)
	MediastinumAnatomy6149 = New("1.2.840.10008.6.1.1314", "Mediastinum Anatomy (6149)", TypeContextGroupName, false)

	// VascularUltrasoundReportDocumentTitle12100 Vascular Ultrasound Report Document Title (12100)
	VascularUltrasoundReportDocumentTitle12100 = New("1.2.840.10008.6.1.1315", "Vascular Ultrasound Report Document Title (12100)", TypeContextGroupName, false)

	// OrganPartNonLateralized12130 Organ Part (Non-Lateralized) (12130)
	OrganPartNonLateralized12130 = New("1.2.840.10008.6.1.1316", "Organ Part (Non-Lateralized) (12130)", TypeContextGroupName, false)

	// OrganPartLateralized12131 Organ Part (Lateralized) (12131)
	OrganPartLateralized12131 = New("1.2.840.10008.6.1.1317", "Organ Part (Lateralized) (12131)", TypeContextGroupName, false)

	// TreatmentTerminationReason9561 Treatment Termination Reason (9561)
	TreatmentTerminationReason9561 = New("1.2.840.10008.6.1.1318", "Treatment Termination Reason (9561)", TypeContextGroupName, false)

	// RadiotherapyTreatmentDeliveryPersonRole9562 Radiotherapy Treatment Delivery Person Role (9562)
	RadiotherapyTreatmentDeliveryPersonRole9562 = New("1.2.840.10008.6.1.1319", "Radiotherapy Treatment Delivery Person Role (9562)", TypeContextGroupName, false)

	// RadiotherapyInterlockResolution9563 Radiotherapy Interlock Resolution (9563)
	RadiotherapyInterlockResolution9563 = New("1.2.840.10008.6.1.1320", "Radiotherapy Interlock Resolution (9563)", TypeContextGroupName, false)

	// TreatmentSessionConfirmationAssertion9564 Treatment Session Confirmation Assertion (9564)
	TreatmentSessionConfirmationAssertion9564 = New("1.2.840.10008.6.1.1321", "Treatment Session Confirmation Assertion (9564)", TypeContextGroupName, false)

	// TreatmentToleranceViolationCause9565 Treatment Tolerance Violation Cause (9565)
	TreatmentToleranceViolationCause9565 = New("1.2.840.10008.6.1.1322", "Treatment Tolerance Violation Cause (9565)", TypeContextGroupName, false)

	// ClinicalToleranceViolationType9566 Clinical Tolerance Violation Type (9566)
	ClinicalToleranceViolationType9566 = New("1.2.840.10008.6.1.1323", "Clinical Tolerance Violation Type (9566)", TypeContextGroupName, false)

	// MachineToleranceViolationType9567 Machine Tolerance Violation Type (9567)
	MachineToleranceViolationType9567 = New("1.2.840.10008.6.1.1324", "Machine Tolerance Violation Type (9567)", TypeContextGroupName, false)

	// RadiotherapyTreatmentInterlock9568 Radiotherapy Treatment Interlock (9568)
	RadiotherapyTreatmentInterlock9568 = New("1.2.840.10008.6.1.1325", "Radiotherapy Treatment Interlock (9568)", TypeContextGroupName, false)

	// IsocentricPatientSupportPositionParameter9569 Isocentric Patient Support Position Parameter (9569)
	IsocentricPatientSupportPositionParameter9569 = New("1.2.840.10008.6.1.1326", "Isocentric Patient Support Position Parameter (9569)", TypeContextGroupName, false)

	// RTOverriddenTreatmentParameter9570 RT Overridden Treatment Parameter (9570)
	RTOverriddenTreatmentParameter9570 = New("1.2.840.10008.6.1.1327", "RT Overridden Treatment Parameter (9570)", TypeContextGroupName, false)

	// EEGLead3030 EEG Lead (3030)
	EEGLead3030 = New("1.2.840.10008.6.1.1328", "EEG Lead (3030)", TypeContextGroupName, false)

	// LeadLocationNearOrInMuscle3031 Lead Location Near or in Muscle (3031)
	LeadLocationNearOrInMuscle3031 = New("1.2.840.10008.6.1.1329", "Lead Location Near or in Muscle (3031)", TypeContextGroupName, false)

	// LeadLocationNearPeripheralNerve3032 Lead Location Near Peripheral Nerve (3032)
	LeadLocationNearPeripheralNerve3032 = New("1.2.840.10008.6.1.1330", "Lead Location Near Peripheral Nerve (3032)", TypeContextGroupName, false)

	// EOGLead3033 EOG Lead (3033)
	EOGLead3033 = New("1.2.840.10008.6.1.1331", "EOG Lead (3033)", TypeContextGroupName, false)

	// BodyPositionChannel3034 Body Position Channel (3034)
	BodyPositionChannel3034 = New("1.2.840.10008.6.1.1332", "Body Position Channel (3034)", TypeContextGroupName, false)

	// EEGAnnotationNeurophysiologicEnumeration3035 EEG Annotation – Neurophysiologic Enumeration (3035)
	EEGAnnotationNeurophysiologicEnumeration3035 = New("1.2.840.10008.6.1.1333", "EEG Annotation – Neurophysiologic Enumeration (3035)", TypeContextGroupName, false)

	// EMGAnnotationNeurophysiologicalEnumeration3036 EMG Annotation – Neurophysiological Enumeration (3036)
	EMGAnnotationNeurophysiologicalEnumeration3036 = New("1.2.840.10008.6.1.1334", "EMG Annotation – Neurophysiological Enumeration (3036)", TypeContextGroupName, false)

	// EOGAnnotationNeurophysiologicalEnumeration3037 EOG Annotation – Neurophysiological Enumeration (3037)
	EOGAnnotationNeurophysiologicalEnumeration3037 = New("1.2.840.10008.6.1.1335", "EOG Annotation – Neurophysiological Enumeration (3037)", TypeContextGroupName, false)

	// PatternEvent3038 Pattern Event (3038)
	PatternEvent3038 = New("1.2.840.10008.6.1.1336", "Pattern Event (3038)", TypeContextGroupName, false)

	// DeviceRelatedAndEnvironmentRelatedEvent3039 Device-related and Environment-related Event (3039)
	DeviceRelatedAndEnvironmentRelatedEvent3039 = New("1.2.840.10008.6.1.1337", "Device-related and Environment-related Event (3039)", TypeContextGroupName, false)

	// EEGAnnotationNeurologicalMonitoringMeasurement3040 EEG Annotation - Neurological Monitoring Measurement (3040)
	EEGAnnotationNeurologicalMonitoringMeasurement3040 = New("1.2.840.10008.6.1.1338", "EEG Annotation - Neurological Monitoring Measurement (3040)", TypeContextGroupName, false)

	// OBGYNUltrasoundReportDocumentTitle12024 OB-GYN Ultrasound Report Document Title (12024)
	OBGYNUltrasoundReportDocumentTitle12024 = New("1.2.840.10008.6.1.1339", "OB-GYN Ultrasound Report Document Title (12024)", TypeContextGroupName, false)

	// AutomationOfMeasurement7230 Automation of Measurement (7230)
	AutomationOfMeasurement7230 = New("1.2.840.10008.6.1.1340", "Automation of Measurement (7230)", TypeContextGroupName, false)

	// OBGYNUltrasoundBeamPath12025 OB-GYN Ultrasound Beam Path (12025)
	OBGYNUltrasoundBeamPath12025 = New("1.2.840.10008.6.1.1341", "OB-GYN Ultrasound Beam Path (12025)", TypeContextGroupName, false)

	// AngleMeasurement7550 Angle Measurement (7550)
	AngleMeasurement7550 = New("1.2.840.10008.6.1.1342", "Angle Measurement (7550)", TypeContextGroupName, false)

	// GenericPurposeOfReferenceToImagesAndCoordinatesInMeasurement7551 Generic Purpose of Reference to Images and Coordinates in Measurement (7551)
	GenericPurposeOfReferenceToImagesAndCoordinatesInMeasurement7551 = New("1.2.840.10008.6.1.1343", "Generic Purpose of Reference to Images and Coordinates in Measurement (7551)", TypeContextGroupName, false)

	// GenericPurposeOfReferenceToImagesInMeasurement7552 Generic Purpose of Reference to Images in Measurement (7552)
	GenericPurposeOfReferenceToImagesInMeasurement7552 = New("1.2.840.10008.6.1.1344", "Generic Purpose of Reference to Images in Measurement (7552)", TypeContextGroupName, false)

	// GenericPurposeOfReferenceToCoordinatesInMeasurement7553 Generic Purpose of Reference to Coordinates in Measurement (7553)
	GenericPurposeOfReferenceToCoordinatesInMeasurement7553 = New("1.2.840.10008.6.1.1345", "Generic Purpose of Reference to Coordinates in Measurement (7553)", TypeContextGroupName, false)

	// FitzpatrickSkinType4401 Fitzpatrick Skin Type (4401)
	FitzpatrickSkinType4401 = New("1.2.840.10008.6.1.1346", "Fitzpatrick Skin Type (4401)", TypeContextGroupName, false)

	// HistoryOfMalignantMelanoma4402 History of Malignant Melanoma (4402)
	HistoryOfMalignantMelanoma4402 = New("1.2.840.10008.6.1.1347", "History of Malignant Melanoma (4402)", TypeContextGroupName, false)

	// HistoryOfMelanomaInSitu4403 History of Melanoma in Situ (4403)
	HistoryOfMelanomaInSitu4403 = New("1.2.840.10008.6.1.1348", "History of Melanoma in Situ (4403)", TypeContextGroupName, false)

	// HistoryOfNonMelanomaSkinCancer4404 History of Non-Melanoma Skin Cancer (4404)
	HistoryOfNonMelanomaSkinCancer4404 = New("1.2.840.10008.6.1.1349", "History of Non-Melanoma Skin Cancer (4404)", TypeContextGroupName, false)

	// SkinDisorder4405 Skin Disorder (4405)
	SkinDisorder4405 = New("1.2.840.10008.6.1.1350", "Skin Disorder (4405)", TypeContextGroupName, false)

	// PatientReportedLesionCharacteristic4406 Patient Reported Lesion Characteristic (4406)
	PatientReportedLesionCharacteristic4406 = New("1.2.840.10008.6.1.1351", "Patient Reported Lesion Characteristic (4406)", TypeContextGroupName, false)

	// LesionPalpationFinding4407 Lesion Palpation Finding (4407)
	LesionPalpationFinding4407 = New("1.2.840.10008.6.1.1352", "Lesion Palpation Finding (4407)", TypeContextGroupName, false)

	// LesionVisualFinding4408 Lesion Visual Finding (4408)
	LesionVisualFinding4408 = New("1.2.840.10008.6.1.1353", "Lesion Visual Finding (4408)", TypeContextGroupName, false)

	// SkinProcedure4409 Skin Procedure (4409)
	SkinProcedure4409 = New("1.2.840.10008.6.1.1354", "Skin Procedure (4409)", TypeContextGroupName, false)

	// AbdominopelvicVessel12125 Abdominopelvic Vessel (12125)
	AbdominopelvicVessel12125 = New("1.2.840.10008.6.1.1355", "Abdominopelvic Vessel (12125)", TypeContextGroupName, false)

	// NumericValueFailureQualifier43 Numeric Value Failure Qualifier (43)
	NumericValueFailureQualifier43 = New("1.2.840.10008.6.1.1356", "Numeric Value Failure Qualifier (43)", TypeContextGroupName, false)

	// NumericValueUnknownQualifier44 Numeric Value Unknown Qualifier (44)
	NumericValueUnknownQualifier44 = New("1.2.840.10008.6.1.1357", "Numeric Value Unknown Qualifier (44)", TypeContextGroupName, false)

	// CouinaudLiverSegment7170 Couinaud Liver Segment (7170)
	CouinaudLiverSegment7170 = New("1.2.840.10008.6.1.1358", "Couinaud Liver Segment (7170)", TypeContextGroupName, false)

	// LiverSegmentationType7171 Liver Segmentation Type (7171)
	LiverSegmentationType7171 = New("1.2.840.10008.6.1.1359", "Liver Segmentation Type (7171)", TypeContextGroupName, false)

	// ContraindicationsForXAImaging1201 Contraindications For XA Imaging (1201)
	ContraindicationsForXAImaging1201 = New("1.2.840.10008.6.1.1360", "Contraindications For XA Imaging (1201)", TypeContextGroupName, false)

	// NeurophysiologicStimulationMode3041 Neurophysiologic Stimulation Mode (3041)
	NeurophysiologicStimulationMode3041 = New("1.2.840.10008.6.1.1361", "Neurophysiologic Stimulation Mode (3041)", TypeContextGroupName, false)

	// ReportedValueType10072 Reported Value Type (10072)
	ReportedValueType10072 = New("1.2.840.10008.6.1.1362", "Reported Value Type (10072)", TypeContextGroupName, false)

	// ValueTiming10073 Value Timing (10073)
	ValueTiming10073 = New("1.2.840.10008.6.1.1363", "Value Timing (10073)", TypeContextGroupName, false)

	// RDSRFrameOfReferenceOrigin10074 RDSR Frame of Reference Origin (10074)
	RDSRFrameOfReferenceOrigin10074 = New("1.2.840.10008.6.1.1364", "RDSR Frame of Reference Origin (10074)", TypeContextGroupName, false)

	// MicroscopyAnnotationPropertyType8135 Microscopy Annotation Property Type (8135)
	MicroscopyAnnotationPropertyType8135 = New("1.2.840.10008.6.1.1365", "Microscopy Annotation Property Type (8135)", TypeContextGroupName, false)

	// MicroscopyMeasurementType8136 Microscopy Measurement Type (8136)
	MicroscopyMeasurementType8136 = New("1.2.840.10008.6.1.1366", "Microscopy Measurement Type (8136)", TypeContextGroupName, false)

	// ProstateReportingSystem6310 Prostate Reporting System (6310)
	ProstateReportingSystem6310 = New("1.2.840.10008.6.1.1367", "Prostate Reporting System (6310)", TypeContextGroupName, false)

	// MRSignalIntensity6311 MR Signal Intensity (6311)
	MRSignalIntensity6311 = New("1.2.840.10008.6.1.1368", "MR Signal Intensity (6311)", TypeContextGroupName, false)

	// CrossSectionalScanPlaneOrientation6312 Cross-sectional Scan Plane Orientation (6312)
	CrossSectionalScanPlaneOrientation6312 = New("1.2.840.10008.6.1.1369", "Cross-sectional Scan Plane Orientation (6312)", TypeContextGroupName, false)

	// HistoryOfProstateDisease6313 History of Prostate Disease (6313)
	HistoryOfProstateDisease6313 = New("1.2.840.10008.6.1.1370", "History of Prostate Disease (6313)", TypeContextGroupName, false)

	// ProstateMRIStudyQualityFinding6314 Prostate MRI Study Quality Finding (6314)
	ProstateMRIStudyQualityFinding6314 = New("1.2.840.10008.6.1.1371", "Prostate MRI Study Quality Finding (6314)", TypeContextGroupName, false)

	// ProstateMRISeriesQualityFinding6315 Prostate MRI Series Quality Finding (6315)
	ProstateMRISeriesQualityFinding6315 = New("1.2.840.10008.6.1.1372", "Prostate MRI Series Quality Finding (6315)", TypeContextGroupName, false)

	// MRImagingArtifact6316 MR Imaging Artifact (6316)
	MRImagingArtifact6316 = New("1.2.840.10008.6.1.1373", "MR Imaging Artifact (6316)", TypeContextGroupName, false)

	// ProstateDCEMRIQualityFinding6317 Prostate DCE MRI Quality Finding (6317)
	ProstateDCEMRIQualityFinding6317 = New("1.2.840.10008.6.1.1374", "Prostate DCE MRI Quality Finding (6317)", TypeContextGroupName, false)

	// ProstateDWIMRIQualityFinding6318 Prostate DWI MRI Quality Finding (6318)
	ProstateDWIMRIQualityFinding6318 = New("1.2.840.10008.6.1.1375", "Prostate DWI MRI Quality Finding (6318)", TypeContextGroupName, false)

	// AbdominalInterventionType6319 Abdominal Intervention Type (6319)
	AbdominalInterventionType6319 = New("1.2.840.10008.6.1.1376", "Abdominal Intervention Type (6319)", TypeContextGroupName, false)

	// AbdominopelvicIntervention6320 Abdominopelvic Intervention (6320)
	AbdominopelvicIntervention6320 = New("1.2.840.10008.6.1.1377", "Abdominopelvic Intervention (6320)", TypeContextGroupName, false)

	// ProstateCancerDiagnosticProcedure6321 Prostate Cancer Diagnostic Procedure (6321)
	ProstateCancerDiagnosticProcedure6321 = New("1.2.840.10008.6.1.1378", "Prostate Cancer Diagnostic Procedure (6321)", TypeContextGroupName, false)

	// ProstateCancerFamilyHistory6322 Prostate Cancer Family History (6322)
	ProstateCancerFamilyHistory6322 = New("1.2.840.10008.6.1.1379", "Prostate Cancer Family History (6322)", TypeContextGroupName, false)

	// ProstateCancerTherapy6323 Prostate Cancer Therapy (6323)
	ProstateCancerTherapy6323 = New("1.2.840.10008.6.1.1380", "Prostate Cancer Therapy (6323)", TypeContextGroupName, false)

	// ProstateMRIAssessment6324 Prostate MRI Assessment (6324)
	ProstateMRIAssessment6324 = New("1.2.840.10008.6.1.1381", "Prostate MRI Assessment (6324)", TypeContextGroupName, false)

	// OverallAssessmentFromPIRADS6325 Overall Assessment from PI-RADS® (6325)
	OverallAssessmentFromPIRADS6325 = New("1.2.840.10008.6.1.1382", "Overall Assessment from PI-RADS® (6325)", TypeContextGroupName, false)

	// ImageQualityControlStandard6326 Image Quality Control Standard (6326)
	ImageQualityControlStandard6326 = New("1.2.840.10008.6.1.1383", "Image Quality Control Standard (6326)", TypeContextGroupName, false)

	// ProstateImagingIndication6327 Prostate Imaging Indication (6327)
	ProstateImagingIndication6327 = New("1.2.840.10008.6.1.1384", "Prostate Imaging Indication (6327)", TypeContextGroupName, false)

	// PIRADSV2LesionAssessmentCategory6328 PI-RADS® v2 Lesion Assessment Category (6328)
	PIRADSV2LesionAssessmentCategory6328 = New("1.2.840.10008.6.1.1385", "PI-RADS® v2 Lesion Assessment Category (6328)", TypeContextGroupName, false)

	// PIRADSV2T2WIPZLesionAssessmentCategory6329 PI-RADS® v2 T2WI PZ Lesion Assessment Category (6329)
	PIRADSV2T2WIPZLesionAssessmentCategory6329 = New("1.2.840.10008.6.1.1386", "PI-RADS® v2 T2WI PZ Lesion Assessment Category (6329)", TypeContextGroupName, false)

	// PIRADSV2T2WITZLesionAssessmentCategory6330 PI-RADS® v2 T2WI TZ Lesion Assessment Category (6330)
	PIRADSV2T2WITZLesionAssessmentCategory6330 = New("1.2.840.10008.6.1.1387", "PI-RADS® v2 T2WI TZ Lesion Assessment Category (6330)", TypeContextGroupName, false)

	// PIRADSV2DWILesionAssessmentCategory6331 PI-RADS® v2 DWI Lesion Assessment Category (6331)
	PIRADSV2DWILesionAssessmentCategory6331 = New("1.2.840.10008.6.1.1388", "PI-RADS® v2 DWI Lesion Assessment Category (6331)", TypeContextGroupName, false)

	// PIRADSV2DCELesionAssessmentCategory6332 PI-RADS® v2 DCE Lesion Assessment Category (6332)
	PIRADSV2DCELesionAssessmentCategory6332 = New("1.2.840.10008.6.1.1389", "PI-RADS® v2 DCE Lesion Assessment Category (6332)", TypeContextGroupName, false)

	// mpMRIAssessmentType6333 mpMRI Assessment Type (6333)
	mpMRIAssessmentType6333 = New("1.2.840.10008.6.1.1390", "mpMRI Assessment Type (6333)", TypeContextGroupName, false)

	// mpMRIAssessmentTypeFromPIRADS6334 mpMRI Assessment Type from PI-RADS® (6334)
	mpMRIAssessmentTypeFromPIRADS6334 = New("1.2.840.10008.6.1.1391", "mpMRI Assessment Type from PI-RADS® (6334)", TypeContextGroupName, false)

	// mpMRIAssessmentValue6335 mpMRI Assessment Value (6335)
	mpMRIAssessmentValue6335 = New("1.2.840.10008.6.1.1392", "mpMRI Assessment Value (6335)", TypeContextGroupName, false)

	// MRIAbnormality6336 MRI Abnormality (6336)
	MRIAbnormality6336 = New("1.2.840.10008.6.1.1393", "MRI Abnormality (6336)", TypeContextGroupName, false)

	// mpMRIProstateAbnormalityFromPIRADS6337 mpMRI Prostate Abnormality from PI-RADS® (6337)
	mpMRIProstateAbnormalityFromPIRADS6337 = New("1.2.840.10008.6.1.1394", "mpMRI Prostate Abnormality from PI-RADS® (6337)", TypeContextGroupName, false)

	// mpMRIBenignProstateAbnormalityFromPIRADS6338 mpMRI Benign Prostate Abnormality from PI-RADS® (6338)
	mpMRIBenignProstateAbnormalityFromPIRADS6338 = New("1.2.840.10008.6.1.1395", "mpMRI Benign Prostate Abnormality from PI-RADS® (6338)", TypeContextGroupName, false)

	// MRIShapeCharacteristic6339 MRI Shape Characteristic (6339)
	MRIShapeCharacteristic6339 = New("1.2.840.10008.6.1.1396", "MRI Shape Characteristic (6339)", TypeContextGroupName, false)

	// ProstateMRIShapeCharacteristicFromPIRADS6340 Prostate MRI Shape Characteristic from PI-RADS® (6340)
	ProstateMRIShapeCharacteristicFromPIRADS6340 = New("1.2.840.10008.6.1.1397", "Prostate MRI Shape Characteristic from PI-RADS® (6340)", TypeContextGroupName, false)

	// MRIMarginCharacteristic6341 MRI Margin Characteristic (6341)
	MRIMarginCharacteristic6341 = New("1.2.840.10008.6.1.1398", "MRI Margin Characteristic (6341)", TypeContextGroupName, false)

	// ProstateMRIMarginCharacteristicFromPIRADS6342 Prostate MRI Margin Characteristic from PI-RADS® (6342)
	ProstateMRIMarginCharacteristicFromPIRADS6342 = New("1.2.840.10008.6.1.1399", "Prostate MRI Margin Characteristic from PI-RADS® (6342)", TypeContextGroupName, false)

	// MRISignalCharacteristic6343 MRI Signal Characteristic (6343)
	MRISignalCharacteristic6343 = New("1.2.840.10008.6.1.1400", "MRI Signal Characteristic (6343)", TypeContextGroupName, false)

	// ProstateMRISignalCharacteristicFromPIRADS6344 Prostate MRI Signal Characteristic from PI-RADS® (6344)
	ProstateMRISignalCharacteristicFromPIRADS6344 = New("1.2.840.10008.6.1.1401", "Prostate MRI Signal Characteristic from PI-RADS® (6344)", TypeContextGroupName, false)

	// MRIEnhancementPattern6345 MRI Enhancement Pattern (6345)
	MRIEnhancementPattern6345 = New("1.2.840.10008.6.1.1402", "MRI Enhancement Pattern (6345)", TypeContextGroupName, false)

	// ProstateMRIEnhancementPatternFromPIRADS6346 Prostate MRI Enhancement Pattern from PI-RADS® (6346)
	ProstateMRIEnhancementPatternFromPIRADS6346 = New("1.2.840.10008.6.1.1403", "Prostate MRI Enhancement Pattern from PI-RADS® (6346)", TypeContextGroupName, false)

	// ProstateMRIExtraProstaticFinding6347 Prostate MRI Extra-prostatic Finding (6347)
	ProstateMRIExtraProstaticFinding6347 = New("1.2.840.10008.6.1.1404", "Prostate MRI Extra-prostatic Finding (6347)", TypeContextGroupName, false)

	// ProstateMRIAssessmentOfExtraProstaticAnatomicSite6348 Prostate MRI Assessment of Extra-prostatic Anatomic Site (6348)
	ProstateMRIAssessmentOfExtraProstaticAnatomicSite6348 = New("1.2.840.10008.6.1.1405", "Prostate MRI Assessment of Extra-prostatic Anatomic Site (6348)", TypeContextGroupName, false)

	// MRCoilType6349 MR Coil Type (6349)
	MRCoilType6349 = New("1.2.840.10008.6.1.1406", "MR Coil Type (6349)", TypeContextGroupName, false)

	// EndorectalCoilFillSubstance6350 Endorectal Coil Fill Substance (6350)
	EndorectalCoilFillSubstance6350 = New("1.2.840.10008.6.1.1407", "Endorectal Coil Fill Substance (6350)", TypeContextGroupName, false)

	// ProstateRelationalMeasurement6351 Prostate Relational Measurement (6351)
	ProstateRelationalMeasurement6351 = New("1.2.840.10008.6.1.1408", "Prostate Relational Measurement (6351)", TypeContextGroupName, false)

	// ProstateCancerDiagnosticBloodLabMeasurement6352 Prostate Cancer Diagnostic Blood Lab Measurement (6352)
	ProstateCancerDiagnosticBloodLabMeasurement6352 = New("1.2.840.10008.6.1.1409", "Prostate Cancer Diagnostic Blood Lab Measurement (6352)", TypeContextGroupName, false)

	// ProstateImagingTypesOfQualityControlStandard6353 Prostate Imaging Types of Quality Control Standard (6353)
	ProstateImagingTypesOfQualityControlStandard6353 = New("1.2.840.10008.6.1.1410", "Prostate Imaging Types of Quality Control Standard (6353)", TypeContextGroupName, false)

	// UltrasoundShearWaveMeasurement12308 Ultrasound Shear Wave Measurement (12308)
	UltrasoundShearWaveMeasurement12308 = New("1.2.840.10008.6.1.1411", "Ultrasound Shear Wave Measurement (12308)", TypeContextGroupName, false)

	// LeftVentricleMyocardialWall16SegmentModel3780RETIRED Left Ventricle Myocardial Wall 16 Segment Model (Retired) (3780)
	LeftVentricleMyocardialWall16SegmentModel3780RETIRED = New("1.2.840.10008.6.1.1412", "Left Ventricle Myocardial Wall 16 Segment Model (Retired) (3780)", TypeContextGroupName, true)

	// LeftVentricleMyocardialWall18SegmentModel3781 Left Ventricle Myocardial Wall 18 Segment Model (3781)
	LeftVentricleMyocardialWall18SegmentModel3781 = New("1.2.840.10008.6.1.1413", "Left Ventricle Myocardial Wall 18 Segment Model (3781)", TypeContextGroupName, false)

	// LeftVentricleBasalWall6Segments3782 Left Ventricle Basal Wall 6 Segments (3782)
	LeftVentricleBasalWall6Segments3782 = New("1.2.840.10008.6.1.1414", "Left Ventricle Basal Wall 6 Segments (3782)", TypeContextGroupName, false)

	// LeftVentricleMidlevelWall6Segments3783 Left Ventricle Midlevel Wall 6 Segments (3783)
	LeftVentricleMidlevelWall6Segments3783 = New("1.2.840.10008.6.1.1415", "Left Ventricle Midlevel Wall 6 Segments (3783)", TypeContextGroupName, false)

	// LeftVentricleApicalWall4Segments3784 Left Ventricle Apical Wall 4 Segments (3784)
	LeftVentricleApicalWall4Segments3784 = New("1.2.840.10008.6.1.1416", "Left Ventricle Apical Wall 4 Segments (3784)", TypeContextGroupName, false)

	// LeftVentricleApicalWall6Segments3785 Left Ventricle Apical Wall 6 Segments (3785)
	LeftVentricleApicalWall6Segments3785 = New("1.2.840.10008.6.1.1417", "Left Ventricle Apical Wall 6 Segments (3785)", TypeContextGroupName, false)

	// PatientTreatmentPreparationMethod9571 Patient Treatment Preparation Method (9571)
	PatientTreatmentPreparationMethod9571 = New("1.2.840.10008.6.1.1418", "Patient Treatment Preparation Method (9571)", TypeContextGroupName, false)

	// PatientShieldingDevice9572 Patient Shielding Device (9572)
	PatientShieldingDevice9572 = New("1.2.840.10008.6.1.1419", "Patient Shielding Device (9572)", TypeContextGroupName, false)

	// PatientTreatmentPreparationDevice9573 Patient Treatment Preparation Device (9573)
	PatientTreatmentPreparationDevice9573 = New("1.2.840.10008.6.1.1420", "Patient Treatment Preparation Device (9573)", TypeContextGroupName, false)

	// PatientPositionDisplacementReferencePoint9574 Patient Position Displacement Reference Point (9574)
	PatientPositionDisplacementReferencePoint9574 = New("1.2.840.10008.6.1.1421", "Patient Position Displacement Reference Point (9574)", TypeContextGroupName, false)

	// PatientAlignmentDevice9575 Patient Alignment Device (9575)
	PatientAlignmentDevice9575 = New("1.2.840.10008.6.1.1422", "Patient Alignment Device (9575)", TypeContextGroupName, false)

	// ReasonsForRTRadiationTreatmentOmission9576 Reasons for RT Radiation Treatment Omission (9576)
	ReasonsForRTRadiationTreatmentOmission9576 = New("1.2.840.10008.6.1.1423", "Reasons for RT Radiation Treatment Omission (9576)", TypeContextGroupName, false)

	// PatientTreatmentPreparationProcedure9577 Patient Treatment Preparation Procedure (9577)
	PatientTreatmentPreparationProcedure9577 = New("1.2.840.10008.6.1.1424", "Patient Treatment Preparation Procedure (9577)", TypeContextGroupName, false)

	// MotionManagementSetupDevice9578 Motion Management Setup Device (9578)
	MotionManagementSetupDevice9578 = New("1.2.840.10008.6.1.1425", "Motion Management Setup Device (9578)", TypeContextGroupName, false)

	// CoreEchoStrainMeasurement12309 Core Echo Strain Measurement (12309)
	CoreEchoStrainMeasurement12309 = New("1.2.840.10008.6.1.1426", "Core Echo Strain Measurement (12309)", TypeContextGroupName, false)

	// MyocardialStrainMethod12310 Myocardial Strain Method (12310)
	MyocardialStrainMethod12310 = New("1.2.840.10008.6.1.1427", "Myocardial Strain Method (12310)", TypeContextGroupName, false)

	// EchoMeasuredStrainProperty12311 Echo Measured Strain Property (12311)
	EchoMeasuredStrainProperty12311 = New("1.2.840.10008.6.1.1428", "Echo Measured Strain Property (12311)", TypeContextGroupName, false)

	// AssessmentFromCADRADS3020 Assessment from CAD-RADS™ (3020)
	AssessmentFromCADRADS3020 = New("1.2.840.10008.6.1.1429", "Assessment from CAD-RADS™ (3020)", TypeContextGroupName, false)

	// CADRADSStenosisAssessmentModifier3021 CAD-RADS™ Stenosis Assessment Modifier (3021)
	CADRADSStenosisAssessmentModifier3021 = New("1.2.840.10008.6.1.1430", "CAD-RADS™ Stenosis Assessment Modifier (3021)", TypeContextGroupName, false)

	// CADRADSAssessmentModifier3022 CAD-RADS™ Assessment Modifier (3022)
	CADRADSAssessmentModifier3022 = New("1.2.840.10008.6.1.1431", "CAD-RADS™ Assessment Modifier (3022)", TypeContextGroupName, false)

	// RTSegmentMaterial9579 RT Segment Material (9579)
	RTSegmentMaterial9579 = New("1.2.840.10008.6.1.1432", "RT Segment Material (9579)", TypeContextGroupName, false)

	// VertebralAnatomicStructure7602 Vertebral Anatomic Structure (7602)
	VertebralAnatomicStructure7602 = New("1.2.840.10008.6.1.1433", "Vertebral Anatomic Structure (7602)", TypeContextGroupName, false)

	// Vertebra7603 Vertebra (7603)
	Vertebra7603 = New("1.2.840.10008.6.1.1434", "Vertebra (7603)", TypeContextGroupName, false)

	// IntervertebralDisc7604 Intervertebral Disc (7604)
	IntervertebralDisc7604 = New("1.2.840.10008.6.1.1435", "Intervertebral Disc (7604)", TypeContextGroupName, false)

	// ImagingProcedure101 Imaging Procedure (101)
	ImagingProcedure101 = New("1.2.840.10008.6.1.1436", "Imaging Procedure (101)", TypeContextGroupName, false)

	// NICIPShortCodeImagingProcedure103 NICIP Short Code Imaging Procedure (103)
	NICIPShortCodeImagingProcedure103 = New("1.2.840.10008.6.1.1437", "NICIP Short Code Imaging Procedure (103)", TypeContextGroupName, false)

	// NICIPSNOMEDImagingProcedure104 NICIP SNOMED Imaging Procedure (104)
	NICIPSNOMEDImagingProcedure104 = New("1.2.840.10008.6.1.1438", "NICIP SNOMED Imaging Procedure (104)", TypeContextGroupName, false)

	// ICD10PCSImagingProcedure105 ICD-10-PCS Imaging Procedure (105)
	ICD10PCSImagingProcedure105 = New("1.2.840.10008.6.1.1439", "ICD-10-PCS Imaging Procedure (105)", TypeContextGroupName, false)

	// ICD10PCSNuclearMedicineProcedure106 ICD-10-PCS Nuclear Medicine Procedure (106)
	ICD10PCSNuclearMedicineProcedure106 = New("1.2.840.10008.6.1.1440", "ICD-10-PCS Nuclear Medicine Procedure (106)", TypeContextGroupName, false)

	// ICD10PCSRadiationTherapyProcedure107 ICD-10-PCS Radiation Therapy Procedure (107)
	ICD10PCSRadiationTherapyProcedure107 = New("1.2.840.10008.6.1.1441", "ICD-10-PCS Radiation Therapy Procedure (107)", TypeContextGroupName, false)

	// RTSegmentationPropertyCategory9580 RT Segmentation Property Category (9580)
	RTSegmentationPropertyCategory9580 = New("1.2.840.10008.6.1.1442", "RT Segmentation Property Category (9580)", TypeContextGroupName, false)

	// RadiotherapyRegistrationMark9581 Radiotherapy Registration Mark (9581)
	RadiotherapyRegistrationMark9581 = New("1.2.840.10008.6.1.1443", "Radiotherapy Registration Mark (9581)", TypeContextGroupName, false)

	// RadiotherapyDoseRegion9582 Radiotherapy Dose Region (9582)
	RadiotherapyDoseRegion9582 = New("1.2.840.10008.6.1.1444", "Radiotherapy Dose Region (9582)", TypeContextGroupName, false)

	// AnatomicallyLocalizedLesionSegmentationType7199 Anatomically Localized Lesion Segmentation Type (7199)
	AnatomicallyLocalizedLesionSegmentationType7199 = New("1.2.840.10008.6.1.1445", "Anatomically Localized Lesion Segmentation Type (7199)", TypeContextGroupName, false)

	// ReasonForRemovalFromOperationalUse7031 Reason for Removal from Operational Use (7031)
	ReasonForRemovalFromOperationalUse7031 = New("1.2.840.10008.6.1.1446", "Reason for Removal from Operational Use (7031)", TypeContextGroupName, false)

	// GeneralUltrasoundReportDocumentTitle12320 General Ultrasound Report Document Title (12320)
	GeneralUltrasoundReportDocumentTitle12320 = New("1.2.840.10008.6.1.1447", "General Ultrasound Report Document Title (12320)", TypeContextGroupName, false)

	// ElastographySite12321 Elastography Site (12321)
	ElastographySite12321 = New("1.2.840.10008.6.1.1448", "Elastography Site (12321)", TypeContextGroupName, false)

	// ElastographyMeasurementSite12322 Elastography Measurement Site (12322)
	ElastographyMeasurementSite12322 = New("1.2.840.10008.6.1.1449", "Elastography Measurement Site (12322)", TypeContextGroupName, false)

	// UltrasoundRelevantPatientCondition12323 Ultrasound Relevant Patient Condition (12323)
	UltrasoundRelevantPatientCondition12323 = New("1.2.840.10008.6.1.1450", "Ultrasound Relevant Patient Condition (12323)", TypeContextGroupName, false)

	// ShearWaveDetectionMethod12324 Shear Wave Detection Method (12324)
	ShearWaveDetectionMethod12324 = New("1.2.840.10008.6.1.1451", "Shear Wave Detection Method (12324)", TypeContextGroupName, false)

	// LiverUltrasoundStudyIndication12325 Liver Ultrasound Study Indication (12325)
	LiverUltrasoundStudyIndication12325 = New("1.2.840.10008.6.1.1452", "Liver Ultrasound Study Indication (12325)", TypeContextGroupName, false)

	// AnalogWaveformFilter3042 Analog Waveform Filter (3042)
	AnalogWaveformFilter3042 = New("1.2.840.10008.6.1.1453", "Analog Waveform Filter (3042)", TypeContextGroupName, false)

	// DigitalWaveformFilter3043 Digital Waveform Filter (3043)
	DigitalWaveformFilter3043 = New("1.2.840.10008.6.1.1454", "Digital Waveform Filter (3043)", TypeContextGroupName, false)

	// WaveformFilterLookupTableInputFrequencyUnit3044 Waveform Filter Lookup Table Input Frequency Unit (3044)
	WaveformFilterLookupTableInputFrequencyUnit3044 = New("1.2.840.10008.6.1.1455", "Waveform Filter Lookup Table Input Frequency Unit (3044)", TypeContextGroupName, false)

	// WaveformFilterLookupTableOutputMagnitudeUnit3045 Waveform Filter Lookup Table Output Magnitude Unit (3045)
	WaveformFilterLookupTableOutputMagnitudeUnit3045 = New("1.2.840.10008.6.1.1456", "Waveform Filter Lookup Table Output Magnitude Unit (3045)", TypeContextGroupName, false)

	// SpecificObservationSubjectClass272 Specific Observation Subject Class (272)
	SpecificObservationSubjectClass272 = New("1.2.840.10008.6.1.1457", "Specific Observation Subject Class (272)", TypeContextGroupName, false)

	// MovableBeamLimitingDeviceType9540 Movable Beam Limiting Device Type (9540)
	MovableBeamLimitingDeviceType9540 = New("1.2.840.10008.6.1.1458", "Movable Beam Limiting Device Type (9540)", TypeContextGroupName, false)

	// RadiotherapyAcquisitionWorkItemSubtasks9260 Radiotherapy Acquisition WorkItem Subtasks (9260)
	RadiotherapyAcquisitionWorkItemSubtasks9260 = New("1.2.840.10008.6.1.1459", "Radiotherapy Acquisition WorkItem Subtasks (9260)", TypeContextGroupName, false)

	// PatientPositionAcquisitionRadiationSourceLocations9261 Patient Position Acquisition Radiation Source Locations (9261)
	PatientPositionAcquisitionRadiationSourceLocations9261 = New("1.2.840.10008.6.1.1460", "Patient Position Acquisition Radiation Source Locations (9261)", TypeContextGroupName, false)

	// EnergyDerivationTypes9262 Energy Derivation Types (9262)
	EnergyDerivationTypes9262 = New("1.2.840.10008.6.1.1461", "Energy Derivation Types (9262)", TypeContextGroupName, false)

	// KVImagingAcquisitionTechniques9263 KV Imaging Acquisition Techniques (9263)
	KVImagingAcquisitionTechniques9263 = New("1.2.840.10008.6.1.1462", "KV Imaging Acquisition Techniques (9263)", TypeContextGroupName, false)

	// MVImagingAcquisitionTechniques9264 MV Imaging Acquisition Techniques (9264)
	MVImagingAcquisitionTechniques9264 = New("1.2.840.10008.6.1.1463", "MV Imaging Acquisition Techniques (9264)", TypeContextGroupName, false)

	// PatientPositionAcquisitionProjectionTechniques9265 Patient Position Acquisition - Projection Techniques (9265)
	PatientPositionAcquisitionProjectionTechniques9265 = New("1.2.840.10008.6.1.1464", "Patient Position Acquisition - Projection Techniques (9265)", TypeContextGroupName, false)

	// PatientPositionAcquisitionCTTechniques9266 Patient Position Acquisition – CT Techniques (9266)
	PatientPositionAcquisitionCTTechniques9266 = New("1.2.840.10008.6.1.1465", "Patient Position Acquisition – CT Techniques (9266)", TypeContextGroupName, false)

	// PatientPositioningRelatedObjectPurposes9267 Patient Positioning Related Object Purposes (9267)
	PatientPositioningRelatedObjectPurposes9267 = New("1.2.840.10008.6.1.1466", "Patient Positioning Related Object Purposes (9267)", TypeContextGroupName, false)

	// PatientPositionAcquisitionDevices9268 Patient Position Acquisition Devices (9268)
	PatientPositionAcquisitionDevices9268 = New("1.2.840.10008.6.1.1467", "Patient Position Acquisition Devices (9268)", TypeContextGroupName, false)

	// RTRadiationMetersetUnits9269 RT Radiation Meterset Units (9269)
	RTRadiationMetersetUnits9269 = New("1.2.840.10008.6.1.1468", "RT Radiation Meterset Units (9269)", TypeContextGroupName, false)

	// AcquisitionInitiationTypes9270 Acquisition Initiation Types (9270)
	AcquisitionInitiationTypes9270 = New("1.2.840.10008.6.1.1469", "Acquisition Initiation Types (9270)", TypeContextGroupName, false)

	// RTImagePatientPositionAcquisitionDevices9271 RT Image Patient Position Acquisition Devices (9271)
	RTImagePatientPositionAcquisitionDevices9271 = New("1.2.840.10008.6.1.1470", "RT Image Patient Position Acquisition Devices (9271)", TypeContextGroupName, false)

	// PhotoacousticIlluminationMethod11001 Photoacoustic Illumination Method (11001)
	PhotoacousticIlluminationMethod11001 = New("1.2.840.10008.6.1.1471", "Photoacoustic Illumination Method (11001)", TypeContextGroupName, false)

	// AcousticCouplingMedium11002 Acoustic Coupling Medium (11002)
	AcousticCouplingMedium11002 = New("1.2.840.10008.6.1.1472", "Acoustic Coupling Medium (11002)", TypeContextGroupName, false)

	// UltrasoundTransducerTechnology11003 Ultrasound Transducer Technology (11003)
	UltrasoundTransducerTechnology11003 = New("1.2.840.10008.6.1.1473", "Ultrasound Transducer Technology (11003)", TypeContextGroupName, false)

	// SpeedOfSoundCorrectionMechanisms11004 Speed of Sound Correction Mechanisms (11004)
	SpeedOfSoundCorrectionMechanisms11004 = New("1.2.840.10008.6.1.1474", "Speed of Sound Correction Mechanisms (11004)", TypeContextGroupName, false)

	// PhotoacousticReconstructionAlgorithmFamily11005 Photoacoustic Reconstruction Algorithm Family (11005)
	PhotoacousticReconstructionAlgorithmFamily11005 = New("1.2.840.10008.6.1.1475", "Photoacoustic Reconstruction Algorithm Family (11005)", TypeContextGroupName, false)

	// PhotoacousticImagedProperty11006 Photoacoustic Imaged Property (11006)
	PhotoacousticImagedProperty11006 = New("1.2.840.10008.6.1.1476", "Photoacoustic Imaged Property (11006)", TypeContextGroupName, false)

	// XRayRadiationDoseProcedureTypeReported10005 X-Ray Radiation Dose Procedure Type Reported (10005)
	XRayRadiationDoseProcedureTypeReported10005 = New("1.2.840.10008.6.1.1477", "X-Ray Radiation Dose Procedure Type Reported (10005)", TypeContextGroupName, false)

	// TopicalTreatment4410 Topical Treatment (4410)
	TopicalTreatment4410 = New("1.2.840.10008.6.1.1478", "Topical Treatment (4410)", TypeContextGroupName, false)

	// LesionColor4411 Lesion Color (4411)
	LesionColor4411 = New("1.2.840.10008.6.1.1479", "Lesion Color (4411)", TypeContextGroupName, false)

	// SpecimenStainForConfocalMicroscopy4412 Specimen Stain for Confocal Microscopy (4412)
	SpecimenStainForConfocalMicroscopy4412 = New("1.2.840.10008.6.1.1480", "Specimen Stain for Confocal Microscopy (4412)", TypeContextGroupName, false)

	// RTROIImageAcquisitionContext9272 RT ROI Image Acquisition Context (9272)
	RTROIImageAcquisitionContext9272 = New("1.2.840.10008.6.1.1481", "RT ROI Image Acquisition Context (9272)", TypeContextGroupName, false)

	// LobeOfLung6170 Lobe of Lung (6170)
	LobeOfLung6170 = New("1.2.840.10008.6.1.1482", "Lobe of Lung (6170)", TypeContextGroupName, false)

	// ZoneOfLung6171 Zone of Lung (6171)
	ZoneOfLung6171 = New("1.2.840.10008.6.1.1483", "Zone of Lung (6171)", TypeContextGroupName, false)

	// SleepStage3046 Sleep Stage (3046)
	SleepStage3046 = New("1.2.840.10008.6.1.1484", "Sleep Stage (3046)", TypeContextGroupName, false)

	// PatientPositionAcquisitionMRTechniques9273 Patient Position Acquisition – MR Techniques (9273)
	PatientPositionAcquisitionMRTechniques9273 = New("1.2.840.10008.6.1.1485", "Patient Position Acquisition – MR Techniques (9273)", TypeContextGroupName, false)

	// RTPlanRadiotherapyProcedureTechnique9583 RT Plan Radiotherapy Procedure Technique (9583)
	RTPlanRadiotherapyProcedureTechnique9583 = New("1.2.840.10008.6.1.1486", "RT Plan Radiotherapy Procedure Technique (9583)", TypeContextGroupName, false)

	// WaveformAnnotationClassification3047 Waveform Annotation Classification (3047)
	WaveformAnnotationClassification3047 = New("1.2.840.10008.6.1.1487", "Waveform Annotation Classification (3047)", TypeContextGroupName, false)

	// WaveformAnnotationsDocumentTitle3048 Waveform Annotations Document Title  (3048)
	WaveformAnnotationsDocumentTitle3048 = New("1.2.840.10008.6.1.1488", "Waveform Annotations Document Title  (3048)", TypeContextGroupName, false)

	// EEGProcedure3049 EEG Procedure  (3049)
	EEGProcedure3049 = New("1.2.840.10008.6.1.1489", "EEG Procedure  (3049)", TypeContextGroupName, false)

	// PatientConsciousness3050 Patient Consciousness  (3050)
	PatientConsciousness3050 = New("1.2.840.10008.6.1.1490", "Patient Consciousness  (3050)", TypeContextGroupName, false)

	// FollicleType12010 Follicle Type (12010)
	FollicleType12010 = New("1.2.840.10008.6.1.1491", "Follicle Type (12010)", TypeContextGroupName, false)

	// BreastSegmentationTypes7163 Breast Segmentation Types (7163)
	BreastSegmentationTypes7163 = New("1.2.840.10008.6.1.1492", "Breast Segmentation Types (7163)", TypeContextGroupName, false)

	// ImplantedDevice3779 Implanted Device (3779)
	ImplantedDevice3779 = New("1.2.840.10008.6.1.1493", "Implanted Device (3779)", TypeContextGroupName, false)

	// SimilarityMeasure281 Similarity Measure (281)
	SimilarityMeasure281 = New("1.2.840.10008.6.1.1494", "Similarity Measure (281)", TypeContextGroupName, false)

	// WaveformAcquisitionModality34 Waveform Acquisition Modality (34)
	WaveformAcquisitionModality34 = New("1.2.840.10008.6.1.1495", "Waveform Acquisition Modality (34)", TypeContextGroupName, false)

	// EnFaceProcessingAlgorithmFamily4274 En Face Processing Algorithm Family (4274)
	EnFaceProcessingAlgorithmFamily4274 = New("1.2.840.10008.6.1.1496", "En Face Processing Algorithm Family (4274)", TypeContextGroupName, false)

	// AnteriorEyeSegmentationSurface4275 Anterior Eye Segmentation Surface (4275)
	AnteriorEyeSegmentationSurface4275 = New("1.2.840.10008.6.1.1497", "Anterior Eye Segmentation Surface (4275)", TypeContextGroupName, false)

	// FetalEchocardiographyImageView12312 Fetal Echocardiography Image View (12312)
	FetalEchocardiographyImageView12312 = New("1.2.840.10008.6.1.1498", "Fetal Echocardiography Image View (12312)", TypeContextGroupName, false)

	// CardiacUltrasoundFetalArrhythmiaMeasurements12313 Cardiac Ultrasound Fetal Arrhythmia Measurements (12313)
	CardiacUltrasoundFetalArrhythmiaMeasurements12313 = New("1.2.840.10008.6.1.1499", "Cardiac Ultrasound Fetal Arrhythmia Measurements (12313)", TypeContextGroupName, false)

	// CommonFetalEchocardiographyMeasurements12314 Common Fetal Echocardiography Measurements (12314)
	CommonFetalEchocardiographyMeasurements12314 = New("1.2.840.10008.6.1.1500", "Common Fetal Echocardiography Measurements (12314)", TypeContextGroupName, false)

	// HeadAndNeckPrimaryAnatomicStructure4061 Head and Neck Primary Anatomic Structure (4061)
	HeadAndNeckPrimaryAnatomicStructure4061 = New("1.2.840.10008.6.1.1501", "Head and Neck Primary Anatomic Structure (4061)", TypeContextGroupName, false)

	// VLView4062 VL View (4062)
	VLView4062 = New("1.2.840.10008.6.1.1502", "VL View (4062)", TypeContextGroupName, false)

	// VLDentalView4063 VL Dental View (4063)
	VLDentalView4063 = New("1.2.840.10008.6.1.1503", "VL Dental View (4063)", TypeContextGroupName, false)

	// VLViewModifier4064 VL View Modifier (4064)
	VLViewModifier4064 = New("1.2.840.10008.6.1.1504", "VL View Modifier (4064)", TypeContextGroupName, false)

	// VLDentalViewModifier4065 VL Dental View Modifier (4065)
	VLDentalViewModifier4065 = New("1.2.840.10008.6.1.1505", "VL Dental View Modifier (4065)", TypeContextGroupName, false)

	// OrthognathicFunctionalCondition4066 Orthognathic Functional Condition (4066)
	OrthognathicFunctionalCondition4066 = New("1.2.840.10008.6.1.1506", "Orthognathic Functional Condition (4066)", TypeContextGroupName, false)

	// OrthodonticFindingByInspection4067 Orthodontic Finding by Inspection (4067)
	OrthodonticFindingByInspection4067 = New("1.2.840.10008.6.1.1507", "Orthodontic Finding by Inspection (4067)", TypeContextGroupName, false)

	// OrthodonticObservableEntity4068 Orthodontic Observable Entity (4068)
	OrthodonticObservableEntity4068 = New("1.2.840.10008.6.1.1508", "Orthodontic Observable Entity (4068)", TypeContextGroupName, false)

	// DentalOcclusion4069 Dental Occlusion (4069)
	DentalOcclusion4069 = New("1.2.840.10008.6.1.1509", "Dental Occlusion (4069)", TypeContextGroupName, false)

	// OrthodonticTreatmentProgress4070 Orthodontic Treatment Progress (4070)
	OrthodonticTreatmentProgress4070 = New("1.2.840.10008.6.1.1510", "Orthodontic Treatment Progress (4070)", TypeContextGroupName, false)

	// GeneralPhotographyDevice4071 General Photography Device (4071)
	GeneralPhotographyDevice4071 = New("1.2.840.10008.6.1.1511", "General Photography Device (4071)", TypeContextGroupName, false)

	// DevicesForThePurposeOfDentalPhotography4072 Devices for the Purpose of Dental Photography (4072)
	DevicesForThePurposeOfDentalPhotography4072 = New("1.2.840.10008.6.1.1512", "Devices for the Purpose of Dental Photography (4072)", TypeContextGroupName, false)

	// CTDIPhantomDevice4053 CTDI Phantom Device (4053)
	CTDIPhantomDevice4053 = New("1.2.840.10008.6.1.1513", "CTDI Phantom Device (4053)", TypeContextGroupName, false)

	// DiagnosticImagingProcedureWithoutIVContrast108 Diagnostic Imaging Procedure without IV Contrast (108)
	DiagnosticImagingProcedureWithoutIVContrast108 = New("1.2.840.10008.6.1.1514", "Diagnostic Imaging Procedure without IV Contrast (108)", TypeContextGroupName, false)

	// DiagnosticImagingProcedureWithIVContrast109 Diagnostic Imaging Procedure with IV Contrast (109)
	DiagnosticImagingProcedureWithIVContrast109 = New("1.2.840.10008.6.1.1515", "Diagnostic Imaging Procedure with IV Contrast (109)", TypeContextGroupName, false)

	// StructuralHeartProcedure12331 Structural Heart Procedure (12331)
	StructuralHeartProcedure12331 = New("1.2.840.10008.6.1.1516", "Structural Heart Procedure (12331)", TypeContextGroupName, false)

	// StructuralHeartDevice12332 Structural Heart Device (12332)
	StructuralHeartDevice12332 = New("1.2.840.10008.6.1.1517", "Structural Heart Device (12332)", TypeContextGroupName, false)

	// StructuralHeartMeasurement12333 Structural Heart Measurement (12333)
	StructuralHeartMeasurement12333 = New("1.2.840.10008.6.1.1518", "Structural Heart Measurement (12333)", TypeContextGroupName, false)

	// AorticValveStructuralMeasurement12334 Aortic Valve Structural Measurement (12334)
	AorticValveStructuralMeasurement12334 = New("1.2.840.10008.6.1.1519", "Aortic Valve Structural Measurement (12334)", TypeContextGroupName, false)

	// MitralValveStructuralMeasurement12335 Mitral Valve Structural Measurement (12335)
	MitralValveStructuralMeasurement12335 = New("1.2.840.10008.6.1.1520", "Mitral Valve Structural Measurement (12335)", TypeContextGroupName, false)

	// TricuspidValveStructuralMeasurement12336 Tricuspid Valve Structural Measurement (12336)
	TricuspidValveStructuralMeasurement12336 = New("1.2.840.10008.6.1.1521", "Tricuspid Valve Structural Measurement (12336)", TypeContextGroupName, false)

	// StructuralHeartEchoMeasurement12337 Structural Heart Echo Measurement (12337)
	StructuralHeartEchoMeasurement12337 = New("1.2.840.10008.6.1.1522", "Structural Heart Echo Measurement (12337)", TypeContextGroupName, false)

	// LeftAtrialAppendageClosureMeasurement12338 Left Atrial Appendage Closure Measurement (12338)
	LeftAtrialAppendageClosureMeasurement12338 = New("1.2.840.10008.6.1.1523", "Left Atrial Appendage Closure Measurement (12338)", TypeContextGroupName, false)

	// StructuralHeartProcedureAnatomicSite12339 Structural Heart Procedure Anatomic Site (12339)
	StructuralHeartProcedureAnatomicSite12339 = New("1.2.840.10008.6.1.1524", "Structural Heart Procedure Anatomic Site (12339)", TypeContextGroupName, false)

	// IndicationForStructuralHeartProcedure12341 Indication for Structural Heart Procedure (12341)
	IndicationForStructuralHeartProcedure12341 = New("1.2.840.10008.6.1.1525", "Indication for Structural Heart Procedure (12341)", TypeContextGroupName, false)

	// BradycardiacAgent12342 Bradycardiac Agent (12342)
	BradycardiacAgent12342 = New("1.2.840.10008.6.1.1526", "Bradycardiac Agent (12342)", TypeContextGroupName, false)

	// TransesophagealEchocardiographyScanPlane12343 Transesophageal Echocardiography Scan Plane (12343)
	TransesophagealEchocardiographyScanPlane12343 = New("1.2.840.10008.6.1.1527", "Transesophageal Echocardiography Scan Plane (12343)", TypeContextGroupName, false)

	// StructuralHeartMeasurementReportDocumentTitle12344 Structural Heart Measurement Report Document Title (12344)
	StructuralHeartMeasurementReportDocumentTitle12344 = New("1.2.840.10008.6.1.1528", "Structural Heart Measurement Report Document Title (12344)", TypeContextGroupName, false)

	// PersonGenderIdentity7458 Person Gender Identity (7458)
	PersonGenderIdentity7458 = New("1.2.840.10008.6.1.1529", "Person Gender Identity (7458)", TypeContextGroupName, false)

	// CategoryOfSexParametersForClinicalUse7459 Category of Sex Parameters for Clinical Use (7459)
	CategoryOfSexParametersForClinicalUse7459 = New("1.2.840.10008.6.1.1530", "Category of Sex Parameters for Clinical Use (7459)", TypeContextGroupName, false)

	// ThirdPersonPronounSet7448 Third Person Pronoun Set (7448)
	ThirdPersonPronounSet7448 = New("1.2.840.10008.6.1.1531", "Third Person Pronoun Set (7448)", TypeContextGroupName, false)

	// CardiacStructureCalcificationQualitativeEvaluation12345 Cardiac Structure Calcification Qualitative Evaluation (12345)
	CardiacStructureCalcificationQualitativeEvaluation12345 = New("1.2.840.10008.6.1.1532", "Cardiac Structure Calcification Qualitative Evaluation (12345)", TypeContextGroupName, false)

	// VisualFieldMeasurements4280 Visual Field Measurements (4280)
	VisualFieldMeasurements4280 = New("1.2.840.10008.6.1.1533", "Visual Field Measurements (4280)", TypeContextGroupName, false)

	// OpticDiscKeyMeasurements4281 Optic Disc Key Measurements (4281)
	OpticDiscKeyMeasurements4281 = New("1.2.840.10008.6.1.1534", "Optic Disc Key Measurements (4281)", TypeContextGroupName, false)

	// RetinalSectorMethods4282 Retinal Sector Methods (4282)
	RetinalSectorMethods4282 = New("1.2.840.10008.6.1.1535", "Retinal Sector Methods (4282)", TypeContextGroupName, false)

	// RNFLSectorMeasurements4283 RNFL Sector Measurements (4283)
	RNFLSectorMeasurements4283 = New("1.2.840.10008.6.1.1536", "RNFL Sector Measurements (4283)", TypeContextGroupName, false)

	// RNFLClockfaceMeasurements4284 RNFL Clockface Measurements (4284)
	RNFLClockfaceMeasurements4284 = New("1.2.840.10008.6.1.1537", "RNFL Clockface Measurements (4284)", TypeContextGroupName, false)

	// MacularThicknessKeyMeasurements4285 Macular Thickness Key Measurements (4285)
	MacularThicknessKeyMeasurements4285 = New("1.2.840.10008.6.1.1538", "Macular Thickness Key Measurements (4285)", TypeContextGroupName, false)

	// GanglionCellMeasurementExtent4286 Ganglion Cell Measurement Extent (4286)
	GanglionCellMeasurementExtent4286 = New("1.2.840.10008.6.1.1539", "Ganglion Cell Measurement Extent (4286)", TypeContextGroupName, false)

	// GanglionCellKeyMeasurements4287 Ganglion Cell Key Measurements (4287)
	GanglionCellKeyMeasurements4287 = New("1.2.840.10008.6.1.1540", "Ganglion Cell Key Measurements (4287)", TypeContextGroupName, false)

	// GanglionCellSectorMeasurements4288 Ganglion Cell Sector Measurements (4288)
	GanglionCellSectorMeasurements4288 = New("1.2.840.10008.6.1.1541", "Ganglion Cell Sector Measurements (4288)", TypeContextGroupName, false)

	// GanglionCellSectorMethods4289 Ganglion Cell Sector Methods (4289)
	GanglionCellSectorMethods4289 = New("1.2.840.10008.6.1.1542", "Ganglion Cell Sector Methods (4289)", TypeContextGroupName, false)

	// EndothelialCellCountMeasurements4290 Endothelial Cell Count Measurements (4290)
	EndothelialCellCountMeasurements4290 = New("1.2.840.10008.6.1.1543", "Endothelial Cell Count Measurements (4290)", TypeContextGroupName, false)

	// OphthalmicImageROIMeasurements4291 Ophthalmic Image ROI Measurements (4291)
	OphthalmicImageROIMeasurements4291 = New("1.2.840.10008.6.1.1544", "Ophthalmic Image ROI Measurements (4291)", TypeContextGroupName, false)

	// RTPlanApprovalAssertion9584 RT Plan Approval Assertion (9584)
	RTPlanApprovalAssertion9584 = New("1.2.840.10008.6.1.1545", "RT Plan Approval Assertion (9584)", TypeContextGroupName, false)

	// EstimatedDeliveryDateMethod12026 Estimated Delivery Date Method (12026)
	EstimatedDeliveryDateMethod12026 = New("1.2.840.10008.6.1.1546", "Estimated Delivery Date Method (12026)", TypeContextGroupName, false)

	// RTDoseCalculationAlgorithmFamily9585 RT Dose Calculation Algorithm Family (9585)
	RTDoseCalculationAlgorithmFamily9585 = New("1.2.840.10008.6.1.1547", "RT Dose Calculation Algorithm Family (9585)", TypeContextGroupName, false)

	// DoseIndexForDoseCalibration10012 Dose Index for Dose Calibration (10012)
	DoseIndexForDoseCalibration10012 = New("1.2.840.10008.6.1.1548", "Dose Index for Dose Calibration (10012)", TypeContextGroupName, false)

	// UltrasoundAttenuationImagingSite12036 Ultrasound Attenuation Imaging Site (12036)
	UltrasoundAttenuationImagingSite12036 = New("1.2.840.10008.6.1.1549", "Ultrasound Attenuation Imaging Site (12036)", TypeContextGroupName, false)
)

func init() {
	// Register all standard UIDs
	Register(Verification)
	Register(ImplicitVRLittleEndian)
	Register(ExplicitVRLittleEndian)
	Register(EncapsulatedUncompressedExplicitVRLittleEndian)
	Register(DeflatedExplicitVRLittleEndian)
	Register(ExplicitVRBigEndianRETIRED)
	Register(JPEGBaseline8Bit)
	Register(JPEGExtended12Bit)
	Register(JPEGExtended35RETIRED)
	Register(JPEGSpectralSelectionNonHierarchical68RETIRED)
	Register(JPEGSpectralSelectionNonHierarchical79RETIRED)
	Register(JPEGFullProgressionNonHierarchical1012RETIRED)
	Register(JPEGFullProgressionNonHierarchical1113RETIRED)
	Register(JPEGLossless)
	Register(JPEGLosslessNonHierarchical15RETIRED)
	Register(JPEGExtendedHierarchical1618RETIRED)
	Register(JPEGExtendedHierarchical1719RETIRED)
	Register(JPEGSpectralSelectionHierarchical2022RETIRED)
	Register(JPEGSpectralSelectionHierarchical2123RETIRED)
	Register(JPEGFullProgressionHierarchical2426RETIRED)
	Register(JPEGFullProgressionHierarchical2527RETIRED)
	Register(JPEGLosslessHierarchical28RETIRED)
	Register(JPEGLosslessHierarchical29RETIRED)
	Register(JPEGLosslessSV1)
	Register(JPEGLSLossless)
	Register(JPEGLSNearLossless)
	Register(JPEG2000Lossless)
	Register(JPEG2000)
	Register(JPEG2000MCLossless)
	Register(JPEG2000MC)
	Register(JPIPReferenced)
	Register(JPIPReferencedDeflate)
	Register(MPEG2MPML)
	Register(MPEG2MPMLF)
	Register(MPEG2MPHL)
	Register(MPEG2MPHLF)
	Register(MPEG4HP41)
	Register(MPEG4HP41F)
	Register(MPEG4HP41BD)
	Register(MPEG4HP41BDF)
	Register(MPEG4HP422D)
	Register(MPEG4HP422DF)
	Register(MPEG4HP423D)
	Register(MPEG4HP423DF)
	Register(MPEG4HP42STEREO)
	Register(MPEG4HP42STEREOF)
	Register(HEVCMP51)
	Register(HEVCM10P51)
	Register(JPEGXLLossless)
	Register(JPEGXLJPEGRecompression)
	Register(JPEGXL)
	Register(HTJ2KLossless)
	Register(HTJ2KLosslessRPCL)
	Register(HTJ2K)
	Register(JPIPHTJ2KReferenced)
	Register(JPIPHTJ2KReferencedDeflate)
	Register(RLELossless)
	Register(RFC2557MIMEEncapsulationRETIRED)
	Register(XMLEncodingRETIRED)
	Register(SMPTEST211020UncompressedProgressiveActiveVideo)
	Register(SMPTEST211020UncompressedInterlacedActiveVideo)
	Register(SMPTEST211030PCMDigitalAudio)
	Register(DeflatedImageFrameCompression)
	Register(MediaStorageDirectoryStorage)
	Register(HotIronPalette)
	Register(PETPalette)
	Register(HotMetalBluePalette)
	Register(PET20StepPalette)
	Register(SpringPalette)
	Register(SummerPalette)
	Register(FallPalette)
	Register(WinterPalette)
	Register(BasicStudyContentNotificationRETIRED)
	Register(Papyrus3ImplicitVRLittleEndianRETIRED)
	Register(StorageCommitmentPushModel)
	Register(StorageCommitmentPushModelInstance)
	Register(StorageCommitmentPullModelRETIRED)
	Register(StorageCommitmentPullModelInstanceRETIRED)
	Register(ProceduralEventLogging)
	Register(ProceduralEventLoggingInstance)
	Register(SubstanceAdministrationLogging)
	Register(SubstanceAdministrationLoggingInstance)
	Register(DCMUID)
	Register(DCM)
	Register(MA)
	Register(UBERON)
	Register(ITIS_TSN)
	Register(MGI)
	Register(PUBCHEM_CID)
	Register(DC)
	Register(NYUMCCG)
	Register(MAYONRISBSASRG)
	Register(IBSI)
	Register(RO)
	Register(RADELEMENT)
	Register(I11)
	Register(UNS)
	Register(RRID)
	Register(DICOMApplicationContext)
	Register(DetachedPatientManagementRETIRED)
	Register(DetachedPatientManagementMetaRETIRED)
	Register(DetachedVisitManagementRETIRED)
	Register(DetachedStudyManagementRETIRED)
	Register(StudyComponentManagementRETIRED)
	Register(ModalityPerformedProcedureStep)
	Register(ModalityPerformedProcedureStepRetrieve)
	Register(ModalityPerformedProcedureStepNotification)
	Register(DetachedResultsManagementRETIRED)
	Register(DetachedResultsManagementMetaRETIRED)
	Register(DetachedStudyManagementMetaRETIRED)
	Register(DetachedInterpretationManagementRETIRED)
	Register(Storage)
	Register(BasicFilmSession)
	Register(BasicFilmBox)
	Register(BasicGrayscaleImageBox)
	Register(BasicColorImageBox)
	Register(ReferencedImageBoxRETIRED)
	Register(BasicGrayscalePrintManagementMeta)
	Register(ReferencedGrayscalePrintManagementMetaRETIRED)
	Register(PrintJob)
	Register(BasicAnnotationBox)
	Register(Printer)
	Register(PrinterConfigurationRetrieval)
	Register(PrinterInstance)
	Register(PrinterConfigurationRetrievalInstance)
	Register(BasicColorPrintManagementMeta)
	Register(ReferencedColorPrintManagementMetaRETIRED)
	Register(VOILUTBox)
	Register(PresentationLUT)
	Register(ImageOverlayBoxRETIRED)
	Register(BasicPrintImageOverlayBoxRETIRED)
	Register(PrintQueueInstanceRETIRED)
	Register(PrintQueueManagementRETIRED)
	Register(StoredPrintStorageRETIRED)
	Register(HardcopyGrayscaleImageStorageRETIRED)
	Register(HardcopyColorImageStorageRETIRED)
	Register(PullPrintRequestRETIRED)
	Register(PullStoredPrintManagementMetaRETIRED)
	Register(MediaCreationManagement)
	Register(DisplaySystem)
	Register(DisplaySystemInstance)
	Register(ComputedRadiographyImageStorage)
	Register(DigitalXRayImageStorageForPresentation)
	Register(DigitalXRayImageStorageForProcessing)
	Register(DigitalMammographyXRayImageStorageForPresentation)
	Register(DigitalMammographyXRayImageStorageForProcessing)
	Register(DigitalIntraOralXRayImageStorageForPresentation)
	Register(DigitalIntraOralXRayImageStorageForProcessing)
	Register(CTImageStorage)
	Register(EnhancedCTImageStorage)
	Register(LegacyConvertedEnhancedCTImageStorage)
	Register(UltrasoundMultiFrameImageStorageRetiredRETIRED)
	Register(UltrasoundMultiFrameImageStorage)
	Register(MRImageStorage)
	Register(EnhancedMRImageStorage)
	Register(MRSpectroscopyStorage)
	Register(EnhancedMRColorImageStorage)
	Register(LegacyConvertedEnhancedMRImageStorage)
	Register(NuclearMedicineImageStorageRetiredRETIRED)
	Register(UltrasoundImageStorageRetiredRETIRED)
	Register(UltrasoundImageStorage)
	Register(EnhancedUSVolumeStorage)
	Register(PhotoacousticImageStorage)
	Register(SecondaryCaptureImageStorage)
	Register(MultiFrameSingleBitSecondaryCaptureImageStorage)
	Register(MultiFrameGrayscaleByteSecondaryCaptureImageStorage)
	Register(MultiFrameGrayscaleWordSecondaryCaptureImageStorage)
	Register(MultiFrameTrueColorSecondaryCaptureImageStorage)
	Register(StandaloneOverlayStorageRETIRED)
	Register(StandaloneCurveStorageRETIRED)
	Register(WaveformStorageTrialRETIRED)
	Register(TwelveLeadECGWaveformStorage)
	Register(GeneralECGWaveformStorage)
	Register(AmbulatoryECGWaveformStorage)
	Register(General32bitECGWaveformStorage)
	Register(HemodynamicWaveformStorage)
	Register(CardiacElectrophysiologyWaveformStorage)
	Register(BasicVoiceAudioWaveformStorage)
	Register(GeneralAudioWaveformStorage)
	Register(ArterialPulseWaveformStorage)
	Register(RespiratoryWaveformStorage)
	Register(MultichannelRespiratoryWaveformStorage)
	Register(RoutineScalpElectroencephalogramWaveformStorage)
	Register(ElectromyogramWaveformStorage)
	Register(ElectrooculogramWaveformStorage)
	Register(SleepElectroencephalogramWaveformStorage)
	Register(BodyPositionWaveformStorage)
	Register(WaveformPresentationStateStorage)
	Register(WaveformAcquisitionPresentationStateStorage)
	Register(StandaloneModalityLUTStorageRETIRED)
	Register(StandaloneVOILUTStorageRETIRED)
	Register(GrayscaleSoftcopyPresentationStateStorage)
	Register(ColorSoftcopyPresentationStateStorage)
	Register(PseudoColorSoftcopyPresentationStateStorage)
	Register(BlendingSoftcopyPresentationStateStorage)
	Register(XAXRFGrayscaleSoftcopyPresentationStateStorage)
	Register(GrayscalePlanarMPRVolumetricPresentationStateStorage)
	Register(CompositingPlanarMPRVolumetricPresentationStateStorage)
	Register(AdvancedBlendingPresentationStateStorage)
	Register(VolumeRenderingVolumetricPresentationStateStorage)
	Register(SegmentedVolumeRenderingVolumetricPresentationStateStorage)
	Register(MultipleVolumeRenderingVolumetricPresentationStateStorage)
	Register(VariableModalityLUTSoftcopyPresentationStateStorage)
	Register(XRayAngiographicImageStorage)
	Register(EnhancedXAImageStorage)
	Register(XRayRadiofluoroscopicImageStorage)
	Register(EnhancedXRFImageStorage)
	Register(XRayAngiographicBiPlaneImageStorageRETIRED)
	Register(XRay3DAngiographicImageStorage)
	Register(XRay3DCraniofacialImageStorage)
	Register(BreastTomosynthesisImageStorage)
	Register(BreastProjectionXRayImageStorageForPresentation)
	Register(BreastProjectionXRayImageStorageForProcessing)
	Register(IntravascularOpticalCoherenceTomographyImageStorageForPresentation)
	Register(IntravascularOpticalCoherenceTomographyImageStorageForProcessing)
	Register(NuclearMedicineImageStorage)
	Register(ParametricMapStorage)
	Register(RawDataStorage)
	Register(SpatialRegistrationStorage)
	Register(SpatialFiducialsStorage)
	Register(DeformableSpatialRegistrationStorage)
	Register(SegmentationStorage)
	Register(SurfaceSegmentationStorage)
	Register(TractographyResultsStorage)
	Register(LabelMapSegmentationStorage)
	Register(HeightMapSegmentationStorage)
	Register(RealWorldValueMappingStorage)
	Register(SurfaceScanMeshStorage)
	Register(SurfaceScanPointCloudStorage)
	Register(VLImageStorageTrialRETIRED)
	Register(VLMultiFrameImageStorageTrialRETIRED)
	Register(VLEndoscopicImageStorage)
	Register(VideoEndoscopicImageStorage)
	Register(VLMicroscopicImageStorage)
	Register(VideoMicroscopicImageStorage)
	Register(VLSlideCoordinatesMicroscopicImageStorage)
	Register(VLPhotographicImageStorage)
	Register(VideoPhotographicImageStorage)
	Register(OphthalmicPhotography8BitImageStorage)
	Register(OphthalmicPhotography16BitImageStorage)
	Register(StereometricRelationshipStorage)
	Register(OphthalmicTomographyImageStorage)
	Register(WideFieldOphthalmicPhotographyStereographicProjectionImageStorage)
	Register(WideFieldOphthalmicPhotography3DCoordinatesImageStorage)
	Register(OphthalmicOpticalCoherenceTomographyEnFaceImageStorage)
	Register(OphthalmicOpticalCoherenceTomographyBscanVolumeAnalysisStorage)
	Register(VLWholeSlideMicroscopyImageStorage)
	Register(DermoscopicPhotographyImageStorage)
	Register(ConfocalMicroscopyImageStorage)
	Register(ConfocalMicroscopyTiledPyramidalImageStorage)
	Register(LensometryMeasurementsStorage)
	Register(AutorefractionMeasurementsStorage)
	Register(KeratometryMeasurementsStorage)
	Register(SubjectiveRefractionMeasurementsStorage)
	Register(VisualAcuityMeasurementsStorage)
	Register(SpectaclePrescriptionReportStorage)
	Register(OphthalmicAxialMeasurementsStorage)
	Register(IntraocularLensCalculationsStorage)
	Register(MacularGridThicknessAndVolumeReportStorage)
	Register(OphthalmicVisualFieldStaticPerimetryMeasurementsStorage)
	Register(OphthalmicThicknessMapStorage)
	Register(CornealTopographyMapStorage)
	Register(TextSRStorageTrialRETIRED)
	Register(AudioSRStorageTrialRETIRED)
	Register(DetailSRStorageTrialRETIRED)
	Register(ComprehensiveSRStorageTrialRETIRED)
	Register(BasicTextSRStorage)
	Register(EnhancedSRStorage)
	Register(ComprehensiveSRStorage)
	Register(Comprehensive3DSRStorage)
	Register(ExtensibleSRStorage)
	Register(ProcedureLogStorage)
	Register(MammographyCADSRStorage)
	Register(KeyObjectSelectionDocumentStorage)
	Register(ChestCADSRStorage)
	Register(XRayRadiationDoseSRStorage)
	Register(RadiopharmaceuticalRadiationDoseSRStorage)
	Register(ColonCADSRStorage)
	Register(ImplantationPlanSRStorage)
	Register(AcquisitionContextSRStorage)
	Register(SimplifiedAdultEchoSRStorage)
	Register(PatientRadiationDoseSRStorage)
	Register(PlannedImagingAgentAdministrationSRStorage)
	Register(PerformedImagingAgentAdministrationSRStorage)
	Register(EnhancedXRayRadiationDoseSRStorage)
	Register(WaveformAnnotationSRStorage)
	Register(ContentAssessmentResultsStorage)
	Register(MicroscopyBulkSimpleAnnotationsStorage)
	Register(EncapsulatedPDFStorage)
	Register(EncapsulatedCDAStorage)
	Register(EncapsulatedSTLStorage)
	Register(EncapsulatedOBJStorage)
	Register(EncapsulatedMTLStorage)
	Register(PositronEmissionTomographyImageStorage)
	Register(LegacyConvertedEnhancedPETImageStorage)
	Register(StandalonePETCurveStorageRETIRED)
	Register(EnhancedPETImageStorage)
	Register(BasicStructuredDisplayStorage)
	Register(CTDefinedProcedureProtocolStorage)
	Register(CTPerformedProcedureProtocolStorage)
	Register(ProtocolApprovalStorage)
	Register(ProtocolApprovalInformationModelFind)
	Register(ProtocolApprovalInformationModelMove)
	Register(ProtocolApprovalInformationModelGet)
	Register(XADefinedProcedureProtocolStorage)
	Register(XAPerformedProcedureProtocolStorage)
	Register(InventoryStorage)
	Register(InventoryFind)
	Register(InventoryMove)
	Register(InventoryGet)
	Register(InventoryCreation)
	Register(RepositoryQuery)
	Register(StorageManagementInstance)
	Register(RTImageStorage)
	Register(RTDoseStorage)
	Register(RTStructureSetStorage)
	Register(RTBeamsTreatmentRecordStorage)
	Register(RTPlanStorage)
	Register(RTBrachyTreatmentRecordStorage)
	Register(RTTreatmentSummaryRecordStorage)
	Register(RTIonPlanStorage)
	Register(RTIonBeamsTreatmentRecordStorage)
	Register(RTPhysicianIntentStorage)
	Register(RTSegmentAnnotationStorage)
	Register(RTRadiationSetStorage)
	Register(CArmPhotonElectronRadiationStorage)
	Register(TomotherapeuticRadiationStorage)
	Register(RoboticArmRadiationStorage)
	Register(RTRadiationRecordSetStorage)
	Register(RTRadiationSalvageRecordStorage)
	Register(TomotherapeuticRadiationRecordStorage)
	Register(CArmPhotonElectronRadiationRecordStorage)
	Register(RoboticRadiationRecordStorage)
	Register(RTRadiationSetDeliveryInstructionStorage)
	Register(RTTreatmentPreparationStorage)
	Register(EnhancedRTImageStorage)
	Register(EnhancedContinuousRTImageStorage)
	Register(RTPatientPositionAcquisitionInstructionStorage)
	Register(DICOSCTImageStorage)
	Register(DICOSDigitalXRayImageStorageForPresentation)
	Register(DICOSDigitalXRayImageStorageForProcessing)
	Register(DICOSThreatDetectionReportStorage)
	Register(DICOS2DAITStorage)
	Register(DICOS3DAITStorage)
	Register(DICOSQuadrupoleResonanceStorage)
	Register(EddyCurrentImageStorage)
	Register(EddyCurrentMultiFrameImageStorage)
	Register(ThermographyImageStorage)
	Register(ThermographyMultiFrameImageStorage)
	Register(UltrasoundWaveformStorage)
	Register(PatientRootQueryRetrieveInformationModelFind)
	Register(PatientRootQueryRetrieveInformationModelMove)
	Register(PatientRootQueryRetrieveInformationModelGet)
	Register(StudyRootQueryRetrieveInformationModelFind)
	Register(StudyRootQueryRetrieveInformationModelMove)
	Register(StudyRootQueryRetrieveInformationModelGet)
	Register(PatientStudyOnlyQueryRetrieveInformationModelFindRETIRED)
	Register(PatientStudyOnlyQueryRetrieveInformationModelMoveRETIRED)
	Register(PatientStudyOnlyQueryRetrieveInformationModelGetRETIRED)
	Register(CompositeInstanceRootRetrieveMove)
	Register(CompositeInstanceRootRetrieveGet)
	Register(CompositeInstanceRetrieveWithoutBulkDataGet)
	Register(DefinedProcedureProtocolInformationModelFind)
	Register(DefinedProcedureProtocolInformationModelMove)
	Register(DefinedProcedureProtocolInformationModelGet)
	Register(ModalityWorklistInformationModelFind)
	Register(GeneralPurposeWorklistManagementMetaRETIRED)
	Register(GeneralPurposeWorklistInformationModelFindRETIRED)
	Register(GeneralPurposeScheduledProcedureStepRETIRED)
	Register(GeneralPurposePerformedProcedureStepRETIRED)
	Register(InstanceAvailabilityNotification)
	Register(RTBeamsDeliveryInstructionStorageTrialRETIRED)
	Register(RTConventionalMachineVerificationTrialRETIRED)
	Register(RTIonMachineVerificationTrialRETIRED)
	Register(UnifiedWorklistAndProcedureStepTrialRETIRED)
	Register(UnifiedProcedureStepPushTrialRETIRED)
	Register(UnifiedProcedureStepWatchTrialRETIRED)
	Register(UnifiedProcedureStepPullTrialRETIRED)
	Register(UnifiedProcedureStepEventTrialRETIRED)
	Register(UPSGlobalSubscriptionInstance)
	Register(UPSFilteredGlobalSubscriptionInstance)
	Register(UnifiedWorklistAndProcedureStep)
	Register(UnifiedProcedureStepPush)
	Register(UnifiedProcedureStepWatch)
	Register(UnifiedProcedureStepPull)
	Register(UnifiedProcedureStepEvent)
	Register(UnifiedProcedureStepQuery)
	Register(RTBeamsDeliveryInstructionStorage)
	Register(RTConventionalMachineVerification)
	Register(RTIonMachineVerification)
	Register(RTBrachyApplicationSetupDeliveryInstructionStorage)
	Register(GeneralRelevantPatientInformationQuery)
	Register(BreastImagingRelevantPatientInformationQuery)
	Register(CardiacRelevantPatientInformationQuery)
	Register(HangingProtocolStorage)
	Register(HangingProtocolInformationModelFind)
	Register(HangingProtocolInformationModelMove)
	Register(HangingProtocolInformationModelGet)
	Register(ColorPaletteStorage)
	Register(ColorPaletteQueryRetrieveInformationModelFind)
	Register(ColorPaletteQueryRetrieveInformationModelMove)
	Register(ColorPaletteQueryRetrieveInformationModelGet)
	Register(ProductCharacteristicsQuery)
	Register(SubstanceApprovalQuery)
	Register(GenericImplantTemplateStorage)
	Register(GenericImplantTemplateInformationModelFind)
	Register(GenericImplantTemplateInformationModelMove)
	Register(GenericImplantTemplateInformationModelGet)
	Register(ImplantAssemblyTemplateStorage)
	Register(ImplantAssemblyTemplateInformationModelFind)
	Register(ImplantAssemblyTemplateInformationModelMove)
	Register(ImplantAssemblyTemplateInformationModelGet)
	Register(ImplantTemplateGroupStorage)
	Register(ImplantTemplateGroupInformationModelFind)
	Register(ImplantTemplateGroupInformationModelMove)
	Register(ImplantTemplateGroupInformationModelGet)
	Register(NativeDICOMModel)
	Register(AbstractMultiDimensionalImageModel)
	Register(DICOMContentMappingResource)
	Register(VideoEndoscopicImageRealTimeCommunication)
	Register(VideoPhotographicImageRealTimeCommunication)
	Register(AudioWaveformRealTimeCommunication)
	Register(RenditionSelectionDocumentRealTimeCommunication)
	Register(dicomDeviceName)
	Register(dicomDescription)
	Register(dicomManufacturer)
	Register(dicomManufacturerModelName)
	Register(dicomSoftwareVersion)
	Register(dicomVendorData)
	Register(dicomAETitle)
	Register(dicomNetworkConnectionReference)
	Register(dicomApplicationCluster)
	Register(dicomAssociationInitiator)
	Register(dicomAssociationAcceptor)
	Register(dicomHostname)
	Register(dicomPort)
	Register(dicomSOPClass)
	Register(dicomTransferRole)
	Register(dicomTransferSyntax)
	Register(dicomPrimaryDeviceType)
	Register(dicomRelatedDeviceReference)
	Register(dicomPreferredCalledAETitle)
	Register(dicomTLSCyphersuite)
	Register(dicomAuthorizedNodeCertificateReference)
	Register(dicomThisNodeCertificateReference)
	Register(dicomInstalled)
	Register(dicomStationName)
	Register(dicomDeviceSerialNumber)
	Register(dicomInstitutionName)
	Register(dicomInstitutionAddress)
	Register(dicomInstitutionDepartmentName)
	Register(dicomIssuerOfPatientID)
	Register(dicomPreferredCallingAETitle)
	Register(dicomSupportedCharacterSet)
	Register(dicomConfigurationRoot)
	Register(dicomDevicesRoot)
	Register(dicomUniqueAETitlesRegistryRoot)
	Register(dicomDevice)
	Register(dicomNetworkAE)
	Register(dicomNetworkConnection)
	Register(dicomUniqueAETitle)
	Register(dicomTransferCapability)
	Register(UTC)
	Register(AnatomicModifier2)
	Register(AnatomicRegion4)
	Register(TransducerApproach5)
	Register(TransducerOrientation6)
	Register(UltrasoundBeamPath7)
	Register(AngiographicInterventionalDevice8)
	Register(ImageGuidedTherapeuticProcedure9)
	Register(InterventionalDrug10)
	Register(AdministrationRoute11)
	Register(ImagingContrastAgent12)
	Register(ImagingContrastAgentIngredient13)
	Register(RadiopharmaceuticalIsotope18)
	Register(PatientOrientation19)
	Register(PatientOrientationModifier20)
	Register(PatientEquipmentRelationship21)
	Register(CranioCaudadAngulation23)
	Register(Radiopharmaceutical25)
	Register(NuclearMedicineProjection26)
	Register(AcquisitionModality29)
	Register(DICOMDevice30)
	Register(AbstractPrior31)
	Register(NumericValueQualifier42)
	Register(MeasurementUnit82)
	Register(RealWorldValueMappingUnit83)
	Register(SignificanceLevel220)
	Register(MeasurementRangeConcept221)
	Register(Normality222)
	Register(NormalRangeValue223)
	Register(SelectionMethod224)
	Register(MeasurementUncertaintyConcept225)
	Register(PopulationStatisticalDescriptor226)
	Register(SampleStatisticalDescriptor227)
	Register(EquationOrTable228)
	Register(YesNo230)
	Register(PresentAbsent240)
	Register(NormalAbnormal242)
	Register(Laterality244)
	Register(PositiveNegative250)
	Register(ComplicationSeverity251)
	Register(ObserverType270)
	Register(ObservationSubjectClass271)
	Register(AudioChannelSource3000)
	Register(ECGLead3001)
	Register(HemodynamicWaveformSource3003)
	Register(CardiovascularAnatomicStructure3010)
	Register(ElectrophysiologyAnatomicLocation3011)
	Register(CoronaryArterySegment3014)
	Register(CoronaryArtery3015)
	Register(CardiovascularAnatomicStructureModifier3019)
	Register(CardiologyMeasurementUnit3082RETIRED)
	Register(TimeSynchronizationChannelType3090)
	Register(CardiacProceduralStateValue3101)
	Register(ElectrophysiologyMeasurementFunctionTechnique3240)
	Register(HemodynamicMeasurementTechnique3241)
	Register(CatheterizationProcedurePhase3250)
	Register(ElectrophysiologyProcedurePhase3254)
	Register(StressProtocol3261)
	Register(ECGPatientStateValue3262)
	Register(ElectrodePlacementValue3263)
	Register(XYZElectrodePlacementValues3264RETIRED)
	Register(HemodynamicPhysiologicalChallenge3271)
	Register(ECGAnnotation3335)
	Register(HemodynamicAnnotation3337)
	Register(ElectrophysiologyAnnotation3339)
	Register(ProcedureLogTitle3400)
	Register(LogNoteType3401)
	Register(PatientStatusAndEvent3402)
	Register(PercutaneousEntry3403)
	Register(StaffAction3404)
	Register(ProcedureActionValue3405)
	Register(NonCoronaryTranscatheterIntervention3406)
	Register(ObjectReferencePurpose3407)
	Register(ConsumableAction3408)
	Register(DrugContrastAdministration3409)
	Register(DrugContrastNumericParameter3410)
	Register(IntracoronaryDevice3411)
	Register(InterventionActionStatus3412)
	Register(AdverseOutcome3413)
	Register(ProcedureUrgency3414)
	Register(CardiacRhythm3415)
	Register(RespirationRhythm3416)
	Register(LesionRisk3418)
	Register(FindingTitle3419)
	Register(ProcedureAction3421)
	Register(DeviceUseAction3422)
	Register(NumericDeviceCharacteristic3423)
	Register(InterventionParameter3425)
	Register(ConsumablesParameter3426)
	Register(EquipmentEvent3427)
	Register(CardiovascularImagingProcedure3428)
	Register(CatheterizationDevice3429)
	Register(DateTimeQualifier3430)
	Register(PeripheralPulseLocation3440)
	Register(PatientAssessment3441)
	Register(PeripheralPulseMethod3442)
	Register(SkinCondition3446)
	Register(AirwayAssessment3448)
	Register(CalibrationObject3451)
	Register(CalibrationMethod3452)
	Register(CardiacVolumeMethod3453)
	Register(IndexMethod3455)
	Register(SubSegmentMethod3456)
	Register(ContourRealignment3458)
	Register(CircumferentialExtent3460)
	Register(RegionalExtent3461)
	Register(ChamberIdentification3462)
	Register(QAReferenceMethod3465)
	Register(PlaneIdentification3466)
	Register(EjectionFraction3467)
	Register(EDVolume3468)
	Register(ESVolume3469)
	Register(VesselLumenCrossSectionalAreaCalculationMethod3470)
	Register(EstimatedVolume3471)
	Register(CardiacContractionPhase3472)
	Register(IVUSProcedurePhase3480)
	Register(IVUSDistanceMeasurement3481)
	Register(IVUSAreaMeasurement3482)
	Register(IVUSLongitudinalMeasurement3483)
	Register(IVUSIndexRatio3484)
	Register(IVUSVolumeMeasurement3485)
	Register(VascularMeasurementSite3486)
	Register(IntravascularVolumetricRegion3487)
	Register(MinMaxMean3488)
	Register(CalciumDistribution3489)
	Register(IVUSLesionMorphology3491)
	Register(VascularDissectionClassification3492)
	Register(IVUSRelativeStenosisSeverity3493)
	Register(IVUSNonMorphologicalFinding3494)
	Register(IVUSPlaqueComposition3495)
	Register(IVUSFiducialPoint3496)
	Register(IVUSArterialMorphology3497)
	Register(PressureUnit3500)
	Register(HemodynamicResistanceUnit3502)
	Register(IndexedHemodynamicResistanceUnit3503)
	Register(CatheterSizeUnit3510)
	Register(SpecimenCollection3515)
	Register(BloodSourceType3520)
	Register(BloodGasPressure3524)
	Register(BloodGasContent3525)
	Register(BloodGasSaturation3526)
	Register(BloodBaseExcess3527)
	Register(BloodPH3528)
	Register(ArterialVenousContent3529)
	Register(OxygenAdministrationAction3530)
	Register(OxygenAdministration3531)
	Register(CirculatorySupportAction3550)
	Register(VentilationAction3551)
	Register(PacingAction3552)
	Register(CirculatorySupport3553)
	Register(Ventilation3554)
	Register(Pacing3555)
	Register(BloodPressureMethod3560)
	Register(RelativeTime3600)
	Register(HemodynamicPatientState3602)
	Register(ArterialLesionLocation3604)
	Register(ArterialSourceLocation3606)
	Register(VenousSourceLocation3607)
	Register(AtrialSourceLocation3608)
	Register(VentricularSourceLocation3609)
	Register(GradientSourceLocation3610)
	Register(PressureMeasurement3611)
	Register(BloodVelocityMeasurement3612)
	Register(HemodynamicTimeMeasurement3613)
	Register(NonMitralValveArea3614)
	Register(ValveArea3615)
	Register(HemodynamicPeriodMeasurement3616)
	Register(ValveFlow3617)
	Register(HemodynamicFlow3618)
	Register(HemodynamicResistanceMeasurement3619)
	Register(HemodynamicRatio3620)
	Register(FractionalFlowReserve3621)
	Register(MeasurementType3627)
	Register(CardiacOutputMethod3628)
	Register(ProcedureIntent3629)
	Register(CardiovascularAnatomicLocation3630)
	Register(Hypertension3640)
	Register(HemodynamicAssessment3641)
	Register(DegreeFinding3642)
	Register(HemodynamicMeasurementPhase3651)
	Register(BodySurfaceAreaEquation3663)
	Register(OxygenConsumptionEquationTable3664)
	Register(P50Equation3666)
	Register(FraminghamScore3667)
	Register(FraminghamTable3668)
	Register(ECGProcedureType3670)
	Register(ReasonForECGStudy3671)
	Register(Pacemaker3672)
	Register(Diagnosis3673RETIRED)
	Register(OtherFilters3675RETIRED)
	Register(LeadMeasurementTechnique3676)
	Register(SummaryCodesECG3677)
	Register(QTCorrectionAlgorithm3678)
	Register(ECGMorphologyDescription3679RETIRED)
	Register(ECGLeadNoiseDescription3680)
	Register(ECGLeadNoiseModifier3681RETIRED)
	Register(Probability3682RETIRED)
	Register(Modifier3683RETIRED)
	Register(Trend3684RETIRED)
	Register(ConjunctiveTerm3685RETIRED)
	Register(ECGInterpretiveStatement3686RETIRED)
	Register(ElectrophysiologyWaveformDuration3687)
	Register(ElectrophysiologyWaveformVoltage3688)
	Register(CathDiagnosis3700)
	Register(CardiacValveTract3701)
	Register(WallMotion3703)
	Register(MyocardiumWallMorphologyFinding3704)
	Register(ChamberSize3705)
	Register(OverallContractility3706)
	Register(VSDDescription3707)
	Register(AorticRootDescription3709)
	Register(CoronaryDominance3710)
	Register(ValvularAbnormality3711)
	Register(VesselDescriptor3712)
	Register(TIMIFlowCharacteristic3713)
	Register(Thrombus3714)
	Register(LesionMargin3715)
	Register(Severity3716)
	Register(LeftVentricleMyocardialWall17SegmentModel3717)
	Register(MyocardialWallSegmentsInProjection3718)
	Register(CanadianClinicalClassification3719)
	Register(CardiacHistoryDate3720RETIRED)
	Register(CardiovascularSurgery3721)
	Register(DiabeticTherapy3722)
	Register(MIType3723)
	Register(SmokingHistory3724)
	Register(CoronaryInterventionIndication3726)
	Register(CatheterizationIndication3727)
	Register(CathFinding3728)
	Register(AdmissionStatus3729)
	Register(InsurancePayor3730)
	Register(PrimaryCauseOfDeath3733)
	Register(AcuteCoronarySyndromeTimePeriod3735)
	Register(NYHAClassification3736)
	Register(IschemiaNonInvasiveTest3737)
	Register(PreCathAnginaType3738)
	Register(CathProcedureType3739)
	Register(ThrombolyticAdministration3740)
	Register(LabVisitMedicationAdministration3741)
	Register(PCIMedicationAdministration3742)
	Register(ClopidogrelTiclopidineAdministration3743)
	Register(EFTestingMethod3744)
	Register(CalculationMethod3745)
	Register(PercutaneousEntrySite3746)
	Register(PercutaneousClosure3747)
	Register(AngiographicEFTestingMethod3748)
	Register(PCIProcedureResult3749)
	Register(PreviouslyDilatedLesion3750)
	Register(GuidewireCrossing3752)
	Register(VascularComplication3754)
	Register(CathComplication3755)
	Register(CardiacPatientRiskFactor3756)
	Register(CardiacDiagnosticProcedure3757)
	Register(CardiovascularFamilyHistory3758)
	Register(HypertensionTherapy3760)
	Register(AntilipemicAgent3761)
	Register(AntiarrhythmicAgent3762)
	Register(MyocardialInfarctionTherapy3764)
	Register(ConcernType3769)
	Register(ProblemStatus3770)
	Register(HealthStatus3772)
	Register(UseStatus3773)
	Register(SocialHistory3774)
	Register(CardiovascularImplant3777)
	Register(PlaqueStructure3802)
	Register(StenosisMeasurementMethod3804)
	Register(StenosisType3805)
	Register(StenosisShape3806)
	Register(VolumeMeasurementMethod3807)
	Register(AneurysmType3808)
	Register(AssociatedCondition3809)
	Register(VascularMorphology3810)
	Register(StentFinding3813)
	Register(StentComposition3814)
	Register(SourceOfVascularFinding3815)
	Register(VascularSclerosisType3817)
	Register(NonInvasiveVascularProcedure3820)
	Register(PapillaryMuscleIncludedExcluded3821)
	Register(RespiratoryStatus3823)
	Register(HeartRhythm3826)
	Register(VesselSegment3827)
	Register(PulmonaryArtery3829)
	Register(StenosisLength3831)
	Register(StenosisGrade3832)
	Register(CardiacEjectionFraction3833)
	Register(CardiacVolumeMeasurement3835)
	Register(TimeBasedPerfusionMeasurement3836)
	Register(FiducialFeature3837)
	Register(DiameterDerivation3838)
	Register(CoronaryVein3839)
	Register(PulmonaryVein3840)
	Register(MyocardialSubsegment3843)
	Register(PartialViewSectionForMammography4005)
	Register(DXAnatomyImaged4009)
	Register(DXView4010)
	Register(DXViewModifier4011)
	Register(ProjectionEponymousName4012)
	Register(AnatomicRegionForMammography4013)
	Register(ViewForMammography4014)
	Register(ViewModifierForMammography4015)
	Register(AnatomicRegionForIntraOralRadiography4016)
	Register(AnatomicRegionModifierForIntraOralRadiography4017)
	Register(PrimaryAnatomicStructureForIntraOralRadiographyPermanentDentitionDesignationOfTeeth4018)
	Register(PrimaryAnatomicStructureForIntraOralRadiographyDeciduousDentitionDesignationOfTeeth4019)
	Register(PETRadionuclide4020)
	Register(PETRadiopharmaceutical4021)
	Register(CraniofacialAnatomicRegion4028)
	Register(CTMRAndPETAnatomyImaged4030)
	Register(CommonAnatomicRegion4031)
	Register(MRSpectroscopyMetabolite4032)
	Register(MRProtonSpectroscopyMetabolite4033)
	Register(EndoscopyAnatomicRegion4040)
	Register(XAXRFAnatomyImaged4042)
	Register(DrugOrContrastAgentCharacteristic4050)
	Register(GeneralDevice4051)
	Register(PhantomDevice4052)
	Register(OphthalmicImagingAgent4200)
	Register(PatientEyeMovementCommand4201)
	Register(OphthalmicPhotographyAcquisitionDevice4202)
	Register(OphthalmicPhotographyIllumination4203)
	Register(OphthalmicFilter4204)
	Register(OphthalmicLens4205)
	Register(OphthalmicChannelDescription4206)
	Register(OphthalmicImagePosition4207)
	Register(MydriaticAgent4208)
	Register(OphthalmicAnatomicStructureImaged4209)
	Register(OphthalmicTomographyAcquisitionDevice4210)
	Register(OphthalmicOCTAnatomicStructureImaged4211)
	Register(Language5000)
	Register(Country5001)
	Register(OverallBreastComposition6000)
	Register(OverallBreastCompositionFromBIRADS6001)
	Register(ChangeSinceLastMammogramOrPriorSurgery6002)
	Register(ChangeSinceLastMammogramOrPriorSurgeryFromBIRADS6003)
	Register(MammographyShapeCharacteristic6004)
	Register(ShapeCharacteristicFromBIRADS6005)
	Register(MammographyMarginCharacteristic6006)
	Register(MarginCharacteristicFromBIRADS6007)
	Register(DensityModifier6008)
	Register(DensityModifierFromBIRADS6009)
	Register(MammographyCalcificationType6010)
	Register(CalcificationTypeFromBIRADS6011)
	Register(CalcificationDistributionModifier6012)
	Register(CalcificationDistributionModifierFromBIRADS6013)
	Register(MammographySingleImageFinding6014)
	Register(SingleImageFindingFromBIRADS6015)
	Register(MammographyCompositeFeature6016)
	Register(CompositeFeatureFromBIRADS6017)
	Register(ClockfaceLocationOrRegion6018)
	Register(ClockfaceLocationOrRegionFromBIRADS6019)
	Register(QuadrantLocation6020)
	Register(QuadrantLocationFromBIRADS6021)
	Register(Side6022)
	Register(SideFromBIRADS6023)
	Register(Depth6024)
	Register(DepthFromBIRADS6025)
	Register(MammographyAssessment6026)
	Register(AssessmentFromBIRADS6027)
	Register(MammographyRecommendedFollowUp6028)
	Register(RecommendedFollowUpFromBIRADS6029)
	Register(MammographyPathologyCode6030)
	Register(BenignPathologyCodeFromBIRADS6031)
	Register(HighRiskLesionPathologyCodeFromBIRADS6032)
	Register(MalignantPathologyCodeFromBIRADS6033)
	Register(CADOutputIntendedUse6034)
	Register(CompositeFeatureRelation6035)
	Register(FeatureScope6036)
	Register(MammographyQuantitativeTemporalDifferenceType6037)
	Register(MammographyQualitativeTemporalDifferenceType6038)
	Register(NippleCharacteristic6039)
	Register(NonLesionObjectType6040)
	Register(MammographyImageQualityFinding6041)
	Register(ResultStatus6042)
	Register(MammographyCADAnalysisType6043)
	Register(ImageQualityAssessmentType6044)
	Register(MammographyQualityControlStandardType6045)
	Register(FollowUpIntervalUnit6046)
	Register(CADProcessingAndFindingSummary6047)
	Register(CADOperatingPointAxisLabel6048)
	Register(BreastProcedureReported6050)
	Register(BreastProcedureReason6051)
	Register(BreastImagingReportSectionTitle6052)
	Register(BreastImagingReportElement6053)
	Register(BreastImagingFinding6054)
	Register(BreastClinicalFindingOrIndicatedProblem6055)
	Register(AssociatedFindingForBreast6056)
	Register(DuctographyFindingForBreast6057)
	Register(ProcedureModifiersForBreast6058)
	Register(BreastImplantType6059)
	Register(BreastBiopsyTechnique6060)
	Register(BreastImagingProcedureModifier6061)
	Register(InterventionalProcedureComplication6062)
	Register(InterventionalProcedureResult6063)
	Register(UltrasoundFindingForBreast6064)
	Register(InstrumentApproach6065)
	Register(TargetConfirmation6066)
	Register(FluidColor6067)
	Register(TumorStagesFromAJCC6068)
	Register(NottinghamCombinedHistologicGrade6069)
	Register(BloomRichardsonHistologicGrade6070)
	Register(HistologicGradingMethod6071)
	Register(BreastImplantFinding6072)
	Register(GynecologicalHormone6080)
	Register(BreastCancerRiskFactor6081)
	Register(GynecologicalProcedure6082)
	Register(ProceduresForBreast6083)
	Register(MammoplastyProcedure6084)
	Register(TherapiesForBreast6085)
	Register(MenopausalPhase6086)
	Register(GeneralRiskFactor6087)
	Register(OBGYNMaternalRiskFactor6088)
	Register(Substance6089)
	Register(RelativeUsageExposureAmount6090)
	Register(RelativeFrequencyOfEventValue6091)
	Register(UsageExposureQualitativeConcept6092)
	Register(UsageExposureAmountQualitativeConcept6093)
	Register(UsageExposureFrequencyQualitativeConcept6094)
	Register(ProcedureNumericProperty6095)
	Register(PregnancyStatus6096)
	Register(SideOfFamily6097)
	Register(ChestComponentCategory6100)
	Register(ChestFindingOrFeature6101)
	Register(ChestFindingOrFeatureModifier6102)
	Register(AbnormalLinesFindingOrFeature6103)
	Register(AbnormalOpacityFindingOrFeature6104)
	Register(AbnormalLucencyFindingOrFeature6105)
	Register(AbnormalTextureFindingOrFeature6106)
	Register(WidthDescriptor6107)
	Register(ChestAnatomicStructureAbnormalDistribution6108)
	Register(RadiographicAnatomyFindingOrFeature6109)
	Register(LungAnatomyFindingOrFeature6110)
	Register(BronchovascularAnatomyFindingOrFeature6111)
	Register(PleuraAnatomyFindingOrFeature6112)
	Register(MediastinumAnatomyFindingOrFeature6113)
	Register(OsseousAnatomyFindingOrFeature6114)
	Register(OsseousAnatomyModifier6115)
	Register(MuscularAnatomy6116)
	Register(VascularAnatomy6117)
	Register(SizeDescriptor6118)
	Register(ChestBorderShape6119)
	Register(ChestBorderDefinition6120)
	Register(ChestOrientationDescriptor6121)
	Register(ChestContentDescriptor6122)
	Register(ChestOpacityDescriptor6123)
	Register(LocationInChest6124)
	Register(GeneralChestLocation6125)
	Register(LocationInLung6126)
	Register(SegmentLocationInLung6127)
	Register(ChestDistributionDescriptor6128)
	Register(ChestSiteInvolvement6129)
	Register(SeverityDescriptor6130)
	Register(ChestTextureDescriptor6131)
	Register(ChestCalcificationDescriptor6132)
	Register(ChestQuantitativeTemporalDifferenceType6133)
	Register(ChestQualitativeTemporalDifferenceType6134)
	Register(ImageQualityFinding6135)
	Register(ChestTypesOfQualityControlStandard6136)
	Register(CADAnalysisType6137)
	Register(ChestNonLesionObjectType6138)
	Register(NonLesionModifier6139)
	Register(CalculationMethod6140)
	Register(AttenuationCoefficientMeasurement6141)
	Register(CalculatedValue6142)
	Register(LesionResponse6143)
	Register(RECISTDefinedLesionResponse6144)
	Register(BaselineCategory6145)
	Register(BackgroundEchotexture6151)
	Register(Orientation6152)
	Register(LesionBoundary6153)
	Register(EchoPattern6154)
	Register(PosteriorAcousticFeature6155)
	Register(Vascularity6157)
	Register(CorrelationToOtherFinding6158)
	Register(MalignancyType6159)
	Register(BreastPrimaryTumorAssessmentFromAJCC6160)
	Register(PathologicalRegionalLymphNodeAssessmentForBreast6161)
	Register(AssessmentOfMetastasisForBreast6162)
	Register(MenstrualCyclePhase6163)
	Register(TimeInterval6164)
	Register(BreastLinearMeasurement6165)
	Register(CADGeometrySecondaryGraphicalRepresentation6166)
	Register(DiagnosticImagingReportDocumentTitle7000)
	Register(DiagnosticImagingReportHeading7001)
	Register(DiagnosticImagingReportElement7002)
	Register(DiagnosticImagingReportPurposeOfReference7003)
	Register(WaveformPurposeOfReference7004)
	Register(ContributingEquipmentPurposeOfReference7005)
	Register(SRDocumentPurposeOfReference7006)
	Register(SignaturePurpose7007)
	Register(MediaImport7008)
	Register(KeyObjectSelectionDocumentTitle7010)
	Register(RejectedForQualityReason7011)
	Register(BestInSet7012)
	Register(DocumentTitle7020)
	Register(RCSRegistrationMethodType7100)
	Register(BrainAtlasFiducial7101)
	Register(SegmentationPropertyCategory7150)
	Register(SegmentationPropertyType7151)
	Register(CardiacStructureSegmentationType7152)
	Register(CNSSegmentationType7153)
	Register(AbdominalSegmentationType7154)
	Register(ThoracicSegmentationType7155)
	Register(VascularSegmentationType7156)
	Register(DeviceSegmentationType7157)
	Register(ArtifactSegmentationType7158)
	Register(LesionSegmentationType7159)
	Register(PelvicOrganSegmentationType7160)
	Register(PhysiologySegmentationType7161)
	Register(ReferencedImagePurposeOfReference7201)
	Register(SourceImagePurposeOfReference7202)
	Register(ImageDerivation7203)
	Register(PurposeOfReferenceToAlternateRepresentation7205)
	Register(RelatedSeriesPurposeOfReference7210)
	Register(MultiFrameSubsetType7250)
	Register(PersonRole7450)
	Register(FamilyMember7451)
	Register(OrganizationalRole7452)
	Register(PerformingRole7453)
	Register(AnimalTaxonomicRankValue7454)
	Register(Sex7455)
	Register(AgeUnit7456)
	Register(LinearMeasurementUnit7460)
	Register(AreaMeasurementUnit7461)
	Register(VolumeMeasurementUnit7462)
	Register(LinearMeasurement7470)
	Register(AreaMeasurement7471)
	Register(VolumeMeasurement7472)
	Register(GeneralAreaCalculationMethod7473)
	Register(GeneralVolumeCalculationMethod7474)
	Register(Breed7480)
	Register(BreedRegistry7481)
	Register(WorkitemDefinition9231)
	Register(NonDICOMOutputTypes9232RETIRED)
	Register(ProcedureDiscontinuationReason9300)
	Register(ScopeOfAccumulation10000)
	Register(UIDType10001)
	Register(IrradiationEventType10002)
	Register(EquipmentPlaneIdentification10003)
	Register(FluoroMode10004)
	Register(XRayFilterMaterial10006)
	Register(XRayFilterType10007)
	Register(DoseRelatedDistanceMeasurement10008)
	Register(MeasuredCalculated10009)
	Register(DoseMeasurementDevice10010)
	Register(EffectiveDoseEvaluationMethod10011)
	Register(CTAcquisitionType10013)
	Register(CTIVContrastImagingTechnique10014)
	Register(CTDoseReferenceAuthority10015)
	Register(AnodeTargetMaterial10016)
	Register(XRayGrid10017)
	Register(UltrasoundProtocolType12001)
	Register(UltrasoundProtocolStageType12002)
	Register(OBGYNDate12003)
	Register(FetalBiometryRatio12004)
	Register(FetalBiometryMeasurement12005)
	Register(FetalLongBonesBiometryMeasurement12006)
	Register(FetalCraniumMeasurement12007)
	Register(OBGYNAmnioticSacMeasurement12008)
	Register(EarlyGestationBiometryMeasurement12009)
	Register(UltrasoundPelvisAndUterusMeasurement12011)
	Register(OBEquationTable12012)
	Register(GestationalAgeEquationTable12013)
	Register(OBFetalBodyWeightEquationTable12014)
	Register(FetalGrowthEquationTable12015)
	Register(EstimatedFetalWeightPercentileEquationTable12016)
	Register(GrowthDistributionRank12017)
	Register(OBGYNSummary12018)
	Register(OBGYNFetusSummary12019)
	Register(VascularSummary12101)
	Register(TemporalPeriodRelatingToProcedureOrTherapy12102)
	Register(VascularUltrasoundAnatomicLocation12103)
	Register(ExtracranialArtery12104)
	Register(IntracranialCerebralVessel12105)
	Register(IntracranialCerebralVesselUnilateral12106)
	Register(UpperExtremityArtery12107)
	Register(UpperExtremityVein12108)
	Register(LowerExtremityArtery12109)
	Register(LowerExtremityVein12110)
	Register(AbdominopelvicArteryPaired12111)
	Register(AbdominopelvicArteryUnpaired12112)
	Register(AbdominopelvicVeinPaired12113)
	Register(AbdominopelvicVeinUnpaired12114)
	Register(RenalVessel12115)
	Register(VesselSegmentModifier12116)
	Register(VesselBranchModifier12117)
	Register(VascularUltrasoundProperty12119)
	Register(UltrasoundBloodVelocityMeasurement12120)
	Register(VascularIndexRatio12121)
	Register(OtherVascularProperty12122)
	Register(CarotidRatio12123)
	Register(RenalRatio12124)
	Register(PelvicVasculatureAnatomicalLocation12140)
	Register(FetalVasculatureAnatomicalLocation12141)
	Register(EchocardiographyLeftVentricleMeasurement12200)
	Register(LeftVentricleLinearMeasurement12201)
	Register(LeftVentricleVolumeMeasurement12202)
	Register(LeftVentricleOtherMeasurement12203)
	Register(EchocardiographyRightVentricleMeasurement12204)
	Register(EchocardiographyLeftAtriumMeasurement12205)
	Register(EchocardiographyRightAtriumMeasurement12206)
	Register(EchocardiographyMitralValveMeasurement12207)
	Register(EchocardiographyTricuspidValveMeasurement12208)
	Register(EchocardiographyPulmonicValveMeasurement12209)
	Register(EchocardiographyPulmonaryArteryMeasurement12210)
	Register(EchocardiographyAorticValveMeasurement12211)
	Register(EchocardiographyAortaMeasurement12212)
	Register(EchocardiographyPulmonaryVeinMeasurement12214)
	Register(EchocardiographyVenaCavaMeasurement12215)
	Register(EchocardiographyHepaticVeinMeasurement12216)
	Register(EchocardiographyCardiacShuntMeasurement12217)
	Register(EchocardiographyCongenitalAnomalyMeasurement12218)
	Register(PulmonaryVeinModifier12219)
	Register(EchocardiographyCommonMeasurement12220)
	Register(FlowDirection12221)
	Register(OrificeFlowProperty12222)
	Register(EchocardiographyStrokeVolumeOrigin12223)
	Register(UltrasoundImageMode12224)
	Register(EchocardiographyImageView12226)
	Register(EchocardiographyMeasurementMethod12227)
	Register(EchocardiographyVolumeMethod12228)
	Register(EchocardiographyAreaMethod12229)
	Register(GradientMethod12230)
	Register(VolumeFlowMethod12231)
	Register(MyocardiumMassMethod12232)
	Register(CardiacPhase12233)
	Register(RespirationState12234)
	Register(MitralValveAnatomicSite12235)
	Register(EchocardiographyAnatomicSite12236)
	Register(EchocardiographyAnatomicSiteModifier12237)
	Register(WallMotionScoringScheme12238)
	Register(CardiacOutputProperty12239)
	Register(LeftVentricleAreaMeasurement12240)
	Register(TricuspidValveFindingSite12241)
	Register(AorticValveFindingSite12242)
	Register(LeftVentricleFindingSite12243)
	Register(CongenitalFindingSite12244)
	Register(SurfaceProcessingAlgorithmFamily7162)
	Register(StressTestProcedurePhase3207)
	Register(Stage3778)
	Register(SMLSizeDescriptor252)
	Register(MajorCoronaryArtery3016)
	Register(RadioactivityUnit3083)
	Register(RestStressState3102)
	Register(PETCardiologyProtocol3106)
	Register(PETCardiologyRadiopharmaceutical3107)
	Register(NMPETProcedure3108)
	Register(NuclearCardiologyProtocol3110)
	Register(NuclearCardiologyRadiopharmaceutical3111)
	Register(AttenuationCorrection3112)
	Register(PerfusionDefectType3113)
	Register(StudyQuality3114)
	Register(StressImagingQualityIssue3115)
	Register(NMExtracardiacFinding3116)
	Register(AttenuationCorrectionMethod3117)
	Register(LevelOfRisk3118)
	Register(LVFunction3119)
	Register(PerfusionFinding3120)
	Register(PerfusionMorphology3121)
	Register(VentricularEnlargement3122)
	Register(StressTestProcedure3200)
	Register(IndicationsForStressTest3201)
	Register(ChestPain3202)
	Register(ExerciserDevice3203)
	Register(StressAgent3204)
	Register(IndicationsForPharmacologicalStressTest3205)
	Register(NonInvasiveCardiacImagingProcedure3206)
	Register(ExerciseECGSummaryCode3208)
	Register(StressImagingSummaryCode3209)
	Register(SpeedOfResponse3210)
	Register(BPResponse3211)
	Register(TreadmillSpeed3212)
	Register(StressHemodynamicFinding3213)
	Register(PerfusionFindingMethod3215)
	Register(ComparisonFinding3217)
	Register(StressSymptom3220)
	Register(StressTestTerminationReason3221)
	Register(QTcMeasurement3227)
	Register(ECGTimingMeasurement3228)
	Register(ECGAxisMeasurement3229)
	Register(ECGFinding3230)
	Register(STSegmentFinding3231)
	Register(STSegmentLocation3232)
	Register(STSegmentMorphology3233)
	Register(EctopicBeatMorphology3234)
	Register(PerfusionComparisonFinding3235)
	Register(ToleranceComparisonFinding3236)
	Register(WallMotionComparisonFinding3237)
	Register(StressScoringScale3238)
	Register(PerceivedExertionScale3239)
	Register(VentricleIdentification3463)
	Register(ColonOverallAssessment6200)
	Register(ColonFindingOrFeature6201)
	Register(ColonFindingOrFeatureModifier6202)
	Register(ColonNonLesionObjectType6203)
	Register(AnatomicNonColonFinding6204)
	Register(ClockfaceLocationForColon6205)
	Register(RecumbentPatientOrientationForColon6206)
	Register(ColonQuantitativeTemporalDifferenceType6207)
	Register(ColonTypesOfQualityControlStandard6208)
	Register(ColonMorphologyDescriptor6209)
	Register(LocationInIntestinalTract6210)
	Register(ColonCADMaterialDescription6211)
	Register(CalculatedValueForColonFinding6212)
	Register(OphthalmicHorizontalDirection4214)
	Register(OphthalmicVerticalDirection4215)
	Register(OphthalmicVisualAcuityType4216)
	Register(ArterialPulseWaveform3004)
	Register(RespirationWaveform3005)
	Register(UltrasoundContrastBolusAgent12030)
	Register(ProtocolIntervalEvent12031)
	Register(TransducerScanPattern12032)
	Register(UltrasoundTransducerGeometry12033)
	Register(UltrasoundTransducerBeamSteering12034)
	Register(UltrasoundTransducerApplication12035)
	Register(InstanceAvailabilityStatus50)
	Register(ModalityPPSDiscontinuationReason9301)
	Register(MediaImportPPSDiscontinuationReason9302)
	Register(DXAnatomyImagedForAnimal7482)
	Register(CommonAnatomicRegionsForAnimal7483)
	Register(DXViewForAnimal7484)
	Register(InstitutionalDepartmentUnitService7030)
	Register(PurposeOfReferenceToPredecessorReport7009)
	Register(VisualFixationQualityDuringAcquisition4220)
	Register(VisualFixationQualityProblem4221)
	Register(OphthalmicMacularGridProblem4222)
	Register(Organization5002)
	Register(MixedBreed7486)
	Register(BroselowLutenPediatricSizeCategory7040)
	Register(CMDCTECCCalciumScoringPatientSizeCategory7042)
	Register(CardiacUltrasoundReportTitle12245)
	Register(CardiacUltrasoundIndicationForStudy12246)
	Register(PediatricFetalAndCongenitalCardiacSurgicalIntervention12247)
	Register(CardiacUltrasoundSummaryCode12248)
	Register(CardiacUltrasoundFetalSummaryCode12249)
	Register(CardiacUltrasoundCommonLinearMeasurement12250)
	Register(CardiacUltrasoundLinearValveMeasurement12251)
	Register(CardiacUltrasoundCardiacFunction12252)
	Register(CardiacUltrasoundAreaMeasurement12253)
	Register(CardiacUltrasoundHemodynamicMeasurement12254)
	Register(CardiacUltrasoundMyocardiumMeasurement12255)
	Register(CardiacUltrasoundLeftVentricleMeasurement12257)
	Register(CardiacUltrasoundRightVentricleMeasurement12258)
	Register(CardiacUltrasoundVentriclesMeasurement12259)
	Register(CardiacUltrasoundPulmonaryArteryMeasurement12260)
	Register(CardiacUltrasoundPulmonaryVein12261)
	Register(CardiacUltrasoundPulmonaryValveMeasurement12262)
	Register(CardiacUltrasoundVenousReturnPulmonaryMeasurement12263)
	Register(CardiacUltrasoundVenousReturnSystemicMeasurement12264)
	Register(CardiacUltrasoundAtriaAndAtrialSeptumMeasurement12265)
	Register(CardiacUltrasoundMitralValveMeasurement12266)
	Register(CardiacUltrasoundTricuspidValveMeasurement12267)
	Register(CardiacUltrasoundAtrioventricularValveMeasurement12268)
	Register(CardiacUltrasoundInterventricularSeptumMeasurement12269)
	Register(CardiacUltrasoundAorticValveMeasurement12270)
	Register(CardiacUltrasoundOutflowTractMeasurement12271)
	Register(CardiacUltrasoundSemilunarValveAnnulateAndSinusMeasurement12272)
	Register(CardiacUltrasoundAorticSinotubularJunctionMeasurement12273)
	Register(CardiacUltrasoundAortaMeasurement12274)
	Register(CardiacUltrasoundCoronaryArteryMeasurement12275)
	Register(CardiacUltrasoundAortoPulmonaryConnectionMeasurement12276)
	Register(CardiacUltrasoundPericardiumAndPleuraMeasurement12277)
	Register(CardiacUltrasoundFetalGeneralMeasurement12279)
	Register(CardiacUltrasoundTargetSite12280)
	Register(CardiacUltrasoundTargetSiteModifier12281)
	Register(CardiacUltrasoundVenousReturnSystemicFindingSite12282)
	Register(CardiacUltrasoundVenousReturnPulmonaryFindingSite12283)
	Register(CardiacUltrasoundAtriaAndAtrialSeptumFindingSite12284)
	Register(CardiacUltrasoundAtrioventricularValveFindingSite12285)
	Register(CardiacUltrasoundInterventricularSeptumFindingSite12286)
	Register(CardiacUltrasoundVentricleFindingSite12287)
	Register(CardiacUltrasoundOutflowTractFindingSite12288)
	Register(CardiacUltrasoundSemilunarValveAnnulusAndSinusFindingSite12289)
	Register(CardiacUltrasoundPulmonaryArteryFindingSite12290)
	Register(CardiacUltrasoundAortaFindingSite12291)
	Register(CardiacUltrasoundCoronaryArteryFindingSite12292)
	Register(CardiacUltrasoundAortopulmonaryConnectionFindingSite12293)
	Register(CardiacUltrasoundPericardiumAndPleuraFindingSite12294)
	Register(OphthalmicUltrasoundAxialMeasurementsType4230)
	Register(LensStatus4231)
	Register(VitreousStatus4232)
	Register(OphthalmicAxialLengthMeasurementsSegmentName4233)
	Register(RefractiveSurgeryType4234)
	Register(KeratometryDescriptor4235)
	Register(IOLCalculationFormula4236)
	Register(LensConstantType4237)
	Register(RefractiveErrorType4238)
	Register(AnteriorChamberDepthDefinition4239)
	Register(OphthalmicMeasurementOrCalculationDataSource4240)
	Register(OphthalmicAxialLengthSelectionMethod4241)
	Register(OphthalmicQualityMetricType4243)
	Register(OphthalmicAgentConcentrationUnit4244)
	Register(FunctionalConditionPresentDuringAcquisition91)
	Register(JointPositionDuringAcquisition92)
	Register(JointPositioningMethod93)
	Register(PhysicalForceAppliedDuringAcquisition94)
	Register(ECGControlNumericVariable3690)
	Register(ECGControlTextVariable3691)
	Register(WholeSlideMicroscopyImageReferencedImagePurposeOfReference8120)
	Register(MicroscopyLensType8121)
	Register(MicroscopyIlluminatorAndSensorColor8122)
	Register(MicroscopyIlluminationMethod8123)
	Register(MicroscopyFilter8124)
	Register(MicroscopyIlluminatorType8125)
	Register(AuditEventID400)
	Register(AuditEventTypeCode401)
	Register(AuditActiveParticipantRoleIDCode402)
	Register(SecurityAlertTypeCode403)
	Register(AuditParticipantObjectIDTypeCode404)
	Register(MediaTypeCode405)
	Register(VisualFieldStaticPerimetryTestPattern4250)
	Register(VisualFieldStaticPerimetryTestStrategy4251)
	Register(VisualFieldStaticPerimetryScreeningTestMode4252)
	Register(VisualFieldStaticPerimetryFixationStrategy4253)
	Register(VisualFieldStaticPerimetryTestAnalysisResult4254)
	Register(VisualFieldIlluminationColor4255)
	Register(VisualFieldProcedureModifier4256)
	Register(VisualFieldGlobalIndexName4257)
	Register(AbstractMultiDimensionalImageModelComponentSemantic7180)
	Register(AbstractMultiDimensionalImageModelComponentUnit7181)
	Register(AbstractMultiDimensionalImageModelDimensionSemantic7182)
	Register(AbstractMultiDimensionalImageModelDimensionUnit7183)
	Register(AbstractMultiDimensionalImageModelAxisDirection7184)
	Register(AbstractMultiDimensionalImageModelAxisOrientation7185)
	Register(AbstractMultiDimensionalImageModelQualitativeDimensionSampleSemantic7186)
	Register(PlanningMethod7320)
	Register(DeIdentificationMethod7050)
	Register(MeasurementOrientation12118)
	Register(ECGGlobalWaveformDuration3689)
	Register(ICD3692)
	Register(RadiotherapyGeneralWorkitemDefinition9241)
	Register(RadiotherapyAcquisitionWorkitemDefinition9242)
	Register(RadiotherapyRegistrationWorkitemDefinition9243)
	Register(ContrastBolusSubstance3850)
	Register(LabelType10022)
	Register(OphthalmicMappingUnitForRealWorldValueMapping4260)
	Register(OphthalmicMappingAcquisitionMethod4261)
	Register(RetinalThicknessDefinition4262)
	Register(OphthalmicThicknessMapValueType4263)
	Register(OphthalmicMapPurposeOfReference4264)
	Register(OphthalmicThicknessDeviationCategory4265)
	Register(OphthalmicAnatomicStructureReferencePoint4266)
	Register(CardiacSynchronizationTechnique3104)
	Register(StainingProtocol8130)
	Register(SizeSpecificDoseEstimationMethodForCT10023)
	Register(PathologyImagingProtocol8131)
	Register(MagnificationSelection8132)
	Register(TissueSelection8133)
	Register(GeneralRegionOfInterestMeasurementModifier7464)
	Register(MeasurementDerivedFromMultipleROIMeasurements7465)
	Register(SurfaceScanAcquisitionType8201)
	Register(SurfaceScanModeType8202)
	Register(SurfaceScanRegistrationMethodType8203)
	Register(BasicCardiacView27)
	Register(CTReconstructionAlgorithm10033)
	Register(DetectorType10030)
	Register(CRDRMechanicalConfiguration10031)
	Register(ProjectionXRayAcquisitionDeviceType10032)
	Register(AbstractSegmentationType7165)
	Register(CommonTissueSegmentationType7166)
	Register(PeripheralNervousSystemSegmentationType7167)
	Register(CornealTopographyMappingUnitForRealWorldValueMapping4267)
	Register(CornealTopographyMapValueType4268)
	Register(BrainStructureForVolumetricMeasurement7140)
	Register(RTDoseDerivation7220)
	Register(RTDosePurposeOfReference7221)
	Register(SpectroscopyPurposeOfReference7215)
	Register(ScheduledProcessingParameterConceptCodesForRTTreatment9250)
	Register(RadiopharmaceuticalOrganDoseReferenceAuthority10040)
	Register(SourceOfRadioisotopeActivityInformation10041)
	Register(IntravenousExtravasationSymptom10043)
	Register(RadiosensitiveOrgan10044)
	Register(RadiopharmaceuticalPatientState10045)
	Register(GFRMeasurement10046)
	Register(GFRMeasurementMethod10047)
	Register(VisualEvaluationMethod8300)
	Register(TestPatternCode8301)
	Register(MeasurementPatternCode8302)
	Register(DisplayDeviceType8303)
	Register(SUVUnit85)
	Register(T1MeasurementMethod4100)
	Register(TracerKineticModel4101)
	Register(PerfusionMeasurementMethod4102)
	Register(ArterialInputFunctionMeasurementMethod4103)
	Register(BolusArrivalTimeDerivationMethod4104)
	Register(PerfusionAnalysisMethod4105)
	Register(QuantitativeMethodUsedForPerfusionAndTracerKineticModel4106)
	Register(TracerKineticModelParameter4107)
	Register(PerfusionModelParameter4108)
	Register(ModelIndependentDynamicContrastAnalysisParameter4109)
	Register(TracerKineticModelingCovariate4110)
	Register(ContrastCharacteristic4111)
	Register(MeasurementReportDocumentTitle7021)
	Register(QuantitativeDiagnosticImagingProcedure100)
	Register(PETRegionOfInterestMeasurement7466)
	Register(GrayLevelCoOccurrenceMatrixMeasurement7467)
	Register(TextureMeasurement7468)
	Register(TimePointType6146)
	Register(GenericIntensityAndSizeMeasurement7469)
	Register(ResponseCriteria6147)
	Register(FetalBiometryAnatomicSite12020)
	Register(FetalLongBoneAnatomicSite12021)
	Register(FetalCraniumAnatomicSite12022)
	Register(PelvisAndUterusAnatomicSite12023)
	Register(ParametricMapDerivationImagePurposeOfReference7222)
	Register(PhysicalQuantityDescriptor9000)
	Register(LymphNodeAnatomicSite7600)
	Register(HeadAndNeckCancerAnatomicSite7601)
	Register(FiberTractInBrainstem7701)
	Register(ProjectionAndThalamicFiber7702)
	Register(AssociationFiber7703)
	Register(LimbicSystemTract7704)
	Register(CommissuralFiber7705)
	Register(CranialNerve7706)
	Register(SpinalCordFiber7707)
	Register(TractographyAnatomicSite7710)
	Register(PrimaryAnatomicStructureForIntraOralRadiographySupernumeraryDentitionDesignationOfTeeth4025)
	Register(PrimaryAnatomicStructureForIntraOralAndCraniofacialRadiographyTeeth4026)
	Register(IEC61217DevicePositionParameter9401)
	Register(IEC61217GantryPositionParameter9402)
	Register(IEC61217PatientSupportPositionParameter9403)
	Register(ActionableFindingClassification7035)
	Register(ImageQualityAssessment7036)
	Register(SummaryRadiationExposureQuantity10050)
	Register(WideFieldOphthalmicPhotographyTransformationMethod4245)
	Register(PETUnit84)
	Register(ImplantMaterial7300)
	Register(InterventionType7301)
	Register(ImplantTemplateViewOrientation7302)
	Register(ImplantTemplateModifiedViewOrientation7303)
	Register(ImplantTargetAnatomy7304)
	Register(ImplantPlanningLandmark7305)
	Register(HumanHipImplantPlanningLandmark7306)
	Register(ImplantComponentType7307)
	Register(HumanHipImplantComponentType7308)
	Register(HumanTraumaImplantComponentType7309)
	Register(ImplantFixationMethod7310)
	Register(DeviceParticipatingRole7445)
	Register(ContainerType8101)
	Register(ContainerComponentType8102)
	Register(AnatomicPathologySpecimenType8103)
	Register(BreastTissueSpecimenType8104)
	Register(SpecimenCollectionProcedure8109)
	Register(SpecimenSamplingProcedure8110)
	Register(SpecimenPreparationProcedure8111)
	Register(SpecimenStain8112)
	Register(SpecimenPreparationStep8113)
	Register(SpecimenFixative8114)
	Register(SpecimenEmbeddingMedia8115)
	Register(SourceOfProjectionXRayDoseInformation10020)
	Register(SourceOfCTDoseInformation10021)
	Register(RadiationDoseReferencePoint10025)
	Register(VolumetricViewDescription501)
	Register(VolumetricViewModifier502)
	Register(DiffusionAcquisitionValueType7260)
	Register(DiffusionModelValueType7261)
	Register(DiffusionTractographyAlgorithmFamily7262)
	Register(DiffusionTractographyMeasurementType7263)
	Register(ResearchAnimalSourceRegistry7490)
	Register(YesNoOnly231)
	Register(BiosafetyLevel601)
	Register(BiosafetyControlReason602)
	Register(SexMaleFemaleOrBoth7457)
	Register(AnimalRoomType603)
	Register(DeviceReuse604)
	Register(AnimalBeddingMaterial605)
	Register(AnimalShelterType606)
	Register(AnimalFeedType607)
	Register(AnimalFeedSource608)
	Register(AnimalFeedingMethod609)
	Register(WaterType610)
	Register(AnesthesiaCategoryCodeTypeForSmallAnimalAnesthesia611)
	Register(AnesthesiaCategoryCodeTypeFromAnesthesiaQualityInitiative612)
	Register(AnesthesiaInductionCodeTypeForSmallAnimalAnesthesia613)
	Register(AnesthesiaInductionCodeTypeFromAnesthesiaQualityInitiative614)
	Register(AnesthesiaMaintenanceCodeTypeForSmallAnimalAnesthesia615)
	Register(AnesthesiaMaintenanceCodeTypeFromAnesthesiaQualityInitiative616)
	Register(AirwayManagementMethodCodeTypeForSmallAnimalAnesthesia617)
	Register(AirwayManagementMethodCodeTypeFromAnesthesiaQualityInitiative618)
	Register(AirwayManagementSubMethodCodeTypeForSmallAnimalAnesthesia619)
	Register(AirwayManagementSubMethodCodeTypeFromAnesthesiaQualityInitiative620)
	Register(MedicationTypeForSmallAnimalAnesthesia621)
	Register(MedicationTypeCodeTypeFromAnesthesiaQualityInitiative622)
	Register(MedicationForSmallAnimalAnesthesia623)
	Register(InhalationalAnesthesiaAgentForSmallAnimalAnesthesia624)
	Register(InjectableAnesthesiaAgentForSmallAnimalAnesthesia625)
	Register(PremedicationAgentForSmallAnimalAnesthesia626)
	Register(NeuromuscularBlockingAgentForSmallAnimalAnesthesia627)
	Register(AncillaryMedicationsForSmallAnimalAnesthesia628)
	Register(CarrierGasesForSmallAnimalAnesthesia629)
	Register(LocalAnestheticsForSmallAnimalAnesthesia630)
	Register(ProcedurePhaseRequiringAnesthesia631)
	Register(SurgicalProcedurePhaseRequiringAnesthesia632)
	Register(PhaseOfImagingProcedureRequiringAnesthesia633RETIRED)
	Register(AnimalHandlingPhase634)
	Register(HeatingMethod635)
	Register(TemperatureSensorDeviceComponentTypeForSmallAnimalProcedure636)
	Register(ExogenousSubstanceType637)
	Register(ExogenousSubstance638)
	Register(TumorGraftHistologicType639)
	Register(Fibril640)
	Register(Virus641)
	Register(Cytokine642)
	Register(Toxin643)
	Register(ExogenousSubstanceAdministrationSite644)
	Register(ExogenousSubstanceOriginTissue645)
	Register(PreclinicalSmallAnimalImagingProcedure646)
	Register(PositionReferenceIndicatorForFrameOfReference647)
	Register(PresentAbsentOnly241)
	Register(WaterEquivalentDiameterMethod10024)
	Register(RadiotherapyPurposeOfReference7022)
	Register(ContentAssessmentType701)
	Register(RTContentAssessmentType702)
	Register(AssessmentBasis703)
	Register(ReaderSpecialty7449)
	Register(RequestedReportType9233)
	Register(CTTransversePlaneReferenceBasis1000)
	Register(AnatomicalReferenceBasis1001)
	Register(AnatomicalReferenceBasisHead1002)
	Register(AnatomicalReferenceBasisSpine1003)
	Register(AnatomicalReferenceBasisChest1004)
	Register(AnatomicalReferenceBasisAbdomenPelvis1005)
	Register(AnatomicalReferenceBasisExtremity1006)
	Register(ReferenceGeometryPlane1010)
	Register(ReferenceGeometryPoint1011)
	Register(PatientAlignmentMethod1015)
	Register(ContraindicationsForCTImaging1200)
	Register(FiducialCategory7110)
	Register(Fiducial7111)
	Register(NonImageSourceInstancePurposeOfReference7013)
	Register(RTProcessOutput7023)
	Register(RTProcessInput7024)
	Register(RTProcessInputUsed7025)
	Register(ProstateAnatomy6300)
	Register(ProstateSectorAnatomyFromPIRADSV26301)
	Register(ProstateSectorAnatomyFromEuropeanConcensus16SectorMinimalModel6302)
	Register(ProstateSectorAnatomyFromEuropeanConcensus27SectorOptimalModel6303)
	Register(MeasurementSelectionReason12301)
	Register(EchoFindingObservationType12302)
	Register(EchoMeasurementType12303)
	Register(CardiovascularMeasuredProperty12304)
	Register(BasicEchoAnatomicSite12305)
	Register(EchoFlowDirection12306)
	Register(CardiacPhaseAndTimePoint12307)
	Register(CoreEchoMeasurement12300)
	Register(OCTAProcessingAlgorithmFamily4270)
	Register(EnFaceImageType4271)
	Register(OPTScanPatternType4272)
	Register(RetinalSegmentationSurface4273)
	Register(OrganForRadiationDoseEstimate10060)
	Register(AbsorbedRadiationDoseType10061)
	Register(EquivalentRadiationDoseType10062)
	Register(RadiationDoseEstimateDistributionRepresentation10063)
	Register(PatientModelType10064)
	Register(RadiationTransportModelType10065)
	Register(AttenuatorCategory10066)
	Register(RadiationAttenuatorMaterial10067)
	Register(EstimateMethodType10068)
	Register(RadiationDoseEstimateParameter10069)
	Register(RadiationDoseType10070)
	Register(MRDiffusionComponentSemantic7270)
	Register(MRDiffusionAnisotropyIndex7271)
	Register(MRDiffusionModelParameter7272)
	Register(MRDiffusionModel7273)
	Register(MRDiffusionModelFittingMethod7274)
	Register(MRDiffusionModelSpecificMethod7275)
	Register(MRDiffusionModelInput7276)
	Register(DiffusionRateAreaOverTimeUnit7277)
	Register(PediatricSizeCategory7039)
	Register(CalciumScoringPatientSizeCategory7041)
	Register(ReasonForRepeatingAcquisition10034)
	Register(ProtocolAssertion800)
	Register(RadiotherapeuticDoseMeasurementDevice7026)
	Register(ExportAdditionalInformationDocumentTitle7014)
	Register(ExportDelayReason7015)
	Register(LevelOfDifficulty7016)
	Register(CategoryOfTeachingMaterialImaging7017)
	Register(MiscellaneousDocumentTitle7018)
	Register(SegmentationNonImageSourcePurposeOfReference7019)
	Register(LongitudinalTemporalEventType280)
	Register(NonLesionObjectTypePhysicalObject6401)
	Register(NonLesionObjectTypeSubstance6402)
	Register(NonLesionObjectTypeTissue6403)
	Register(ChestNonLesionObjectTypePhysicalObject6404)
	Register(ChestNonLesionObjectTypeTissue6405)
	Register(TissueSegmentationPropertyType7191)
	Register(AnatomicalStructureSegmentationPropertyType7192)
	Register(PhysicalObjectSegmentationPropertyType7193)
	Register(MorphologicallyAbnormalStructureSegmentationPropertyType7194)
	Register(FunctionSegmentationPropertyType7195)
	Register(SpatialAndRelationalConceptSegmentationPropertyType7196)
	Register(BodySubstanceSegmentationPropertyType7197)
	Register(SubstanceSegmentationPropertyType7198)
	Register(InterpretationRequestDiscontinuationReason9303)
	Register(GrayLevelRunLengthBasedFeature7475)
	Register(GrayLevelSizeZoneBasedFeature7476)
	Register(EncapsulatedDocumentSourcePurposeOfReference7060)
	Register(ModelDocumentTitle7061)
	Register(PurposeOfReferenceToPredecessor3DModel7062)
	Register(ModelScaleUnit7063)
	Register(ModelUsage7064)
	Register(RadiationDoseUnit10071)
	Register(RadiotherapyFiducial7112)
	Register(MultiEnergyRelevantMaterial300)
	Register(MultiEnergyMaterialUnit301)
	Register(DosimetricObjectiveType9500)
	Register(PrescriptionAnatomyCategory9501)
	Register(RTSegmentAnnotationCategory9502)
	Register(RadiotherapyTherapeuticRoleCategory9503)
	Register(RTGeometricInformation9504)
	Register(FixationOrPositioningDevice9505)
	Register(BrachytherapyDevice9506)
	Register(ExternalBodyModel9507)
	Register(NonSpecificVolume9508)
	Register(PurposeOfReferenceForRTPhysicianIntentInput9509)
	Register(PurposeOfReferenceForRTTreatmentPlanningInput9510)
	Register(GeneralExternalRadiotherapyProcedureTechnique9511)
	Register(TomotherapeuticRadiotherapyProcedureTechnique9512)
	Register(FixationDevice9513)
	Register(AnatomicalStructureForRadiotherapy9514)
	Register(RTPatientSupportDevice9515)
	Register(RadiotherapyBolusDeviceType9516)
	Register(RadiotherapyBlockDeviceType9517)
	Register(RadiotherapyAccessoryNoSlotHolderDeviceType9518)
	Register(RadiotherapyAccessorySlotHolderDeviceType9519)
	Register(SegmentedRTAccessoryDevice9520)
	Register(RadiotherapyTreatmentEnergyUnit9521)
	Register(MultiSourceRadiotherapyProcedureTechnique9522)
	Register(RoboticRadiotherapyProcedureTechnique9523)
	Register(RadiotherapyProcedureTechnique9524)
	Register(RadiationTherapyParticle9525)
	Register(IonTherapyParticle9526)
	Register(TeletherapyIsotope9527)
	Register(BrachytherapyIsotope9528)
	Register(SingleDoseDosimetricObjective9529)
	Register(PercentageAndDoseDosimetricObjective9530)
	Register(VolumeAndDoseDosimetricObjective9531)
	Register(NoParameterDosimetricObjective9532)
	Register(DeliveryTimeStructure9533)
	Register(RadiotherapyTarget9534)
	Register(RadiotherapyDoseCalculationRole9535)
	Register(RadiotherapyPrescribingAndSegmentingPersonRole9536)
	Register(EffectiveDoseCalculationMethodCategory9537)
	Register(RadiationTransportBasedEffectiveDoseMethodModifier9538)
	Register(FractionationBasedEffectiveDoseMethodModifier9539)
	Register(ImagingAgentAdministrationAdverseEvent60)
	Register(TimeRelativeToProcedure61RETIRED)
	Register(ImagingAgentAdministrationPhaseType62)
	Register(ImagingAgentAdministrationMode63)
	Register(ImagingAgentAdministrationPatientState64)
	Register(ImagingAgentAdministrationPremedication65)
	Register(ImagingAgentAdministrationMedication66)
	Register(ImagingAgentAdministrationCompletionStatus67)
	Register(ImagingAgentAdministrationPharmaceuticalPresentationUnit68)
	Register(ImagingAgentAdministrationConsumable69)
	Register(Flush70)
	Register(ImagingAgentAdministrationInjectorEventType71)
	Register(ImagingAgentAdministrationStepType72)
	Register(BolusShapingCurve73)
	Register(ImagingAgentAdministrationConsumableCatheterType74)
	Register(LowHighOrEqual75)
	Register(PremedicationType76)
	Register(LateralityWithMedian245)
	Register(DermatologyAnatomicSite4029)
	Register(QuantitativeImageFeature218)
	Register(GlobalShapeDescriptor7477)
	Register(IntensityHistogramFeature7478)
	Register(GreyLevelDistanceZoneBasedFeature7479)
	Register(NeighbourhoodGreyToneDifferenceBasedFeature7500)
	Register(NeighbouringGreyLevelDependenceBasedFeature7501)
	Register(CorneaMeasurementMethodDescriptor4242)
	Register(SegmentedRadiotherapeuticDoseMeasurementDevice7027)
	Register(ClinicalCourseOfDisease6098)
	Register(RacialGroup6099)
	Register(RelativeLaterality246)
	Register(BrainLesionSegmentationTypeWithNecrosis7168)
	Register(BrainLesionSegmentationTypeWithoutNecrosis7169)
	Register(NonAcquisitionModality32)
	Register(Modality33)
	Register(LateralityLeftRightOnly247)
	Register(QualitativeEvaluationModifierType210)
	Register(QualitativeEvaluationModifierValue211)
	Register(GenericAnatomicLocationModifier212)
	Register(BeamLimitingDeviceType9541)
	Register(CompensatorDeviceType9542)
	Register(RadiotherapyTreatmentMachineMode9543)
	Register(RadiotherapyDistanceReferenceLocation9544)
	Register(FixedBeamLimitingDeviceType9545)
	Register(RadiotherapyWedgeType9546)
	Register(RTBeamLimitingDeviceOrientationLabel9547)
	Register(GeneralAccessoryDeviceType9548)
	Register(RadiationGenerationModeType9549)
	Register(CArmPhotonElectronDeliveryRateUnit9550)
	Register(TreatmentDeliveryDeviceType9551)
	Register(CArmPhotonElectronDosimeterUnit9552)
	Register(TreatmentPoint9553)
	Register(EquipmentReferencePoint9554)
	Register(RadiotherapyTreatmentPlanningPersonRole9555)
	Register(RealTimeVideoRenditionTitle7070)
	Register(GeometryGraphicalRepresentation219)
	Register(VisualExplanation217)
	Register(ProstateSectorAnatomyFromPIRADSV216304)
	Register(RadiotherapyRoboticNodeSet9556)
	Register(TomotherapeuticDosimeterUnit9557)
	Register(TomotherapeuticDoseRateUnit9558)
	Register(RoboticDeliveryDeviceDosimeterUnit9559)
	Register(RoboticDeliveryDeviceDoseRateUnit9560)
	Register(AnatomicStructure8134)
	Register(MediastinumFindingOrFeature6148)
	Register(MediastinumAnatomy6149)
	Register(VascularUltrasoundReportDocumentTitle12100)
	Register(OrganPartNonLateralized12130)
	Register(OrganPartLateralized12131)
	Register(TreatmentTerminationReason9561)
	Register(RadiotherapyTreatmentDeliveryPersonRole9562)
	Register(RadiotherapyInterlockResolution9563)
	Register(TreatmentSessionConfirmationAssertion9564)
	Register(TreatmentToleranceViolationCause9565)
	Register(ClinicalToleranceViolationType9566)
	Register(MachineToleranceViolationType9567)
	Register(RadiotherapyTreatmentInterlock9568)
	Register(IsocentricPatientSupportPositionParameter9569)
	Register(RTOverriddenTreatmentParameter9570)
	Register(EEGLead3030)
	Register(LeadLocationNearOrInMuscle3031)
	Register(LeadLocationNearPeripheralNerve3032)
	Register(EOGLead3033)
	Register(BodyPositionChannel3034)
	Register(EEGAnnotationNeurophysiologicEnumeration3035)
	Register(EMGAnnotationNeurophysiologicalEnumeration3036)
	Register(EOGAnnotationNeurophysiologicalEnumeration3037)
	Register(PatternEvent3038)
	Register(DeviceRelatedAndEnvironmentRelatedEvent3039)
	Register(EEGAnnotationNeurologicalMonitoringMeasurement3040)
	Register(OBGYNUltrasoundReportDocumentTitle12024)
	Register(AutomationOfMeasurement7230)
	Register(OBGYNUltrasoundBeamPath12025)
	Register(AngleMeasurement7550)
	Register(GenericPurposeOfReferenceToImagesAndCoordinatesInMeasurement7551)
	Register(GenericPurposeOfReferenceToImagesInMeasurement7552)
	Register(GenericPurposeOfReferenceToCoordinatesInMeasurement7553)
	Register(FitzpatrickSkinType4401)
	Register(HistoryOfMalignantMelanoma4402)
	Register(HistoryOfMelanomaInSitu4403)
	Register(HistoryOfNonMelanomaSkinCancer4404)
	Register(SkinDisorder4405)
	Register(PatientReportedLesionCharacteristic4406)
	Register(LesionPalpationFinding4407)
	Register(LesionVisualFinding4408)
	Register(SkinProcedure4409)
	Register(AbdominopelvicVessel12125)
	Register(NumericValueFailureQualifier43)
	Register(NumericValueUnknownQualifier44)
	Register(CouinaudLiverSegment7170)
	Register(LiverSegmentationType7171)
	Register(ContraindicationsForXAImaging1201)
	Register(NeurophysiologicStimulationMode3041)
	Register(ReportedValueType10072)
	Register(ValueTiming10073)
	Register(RDSRFrameOfReferenceOrigin10074)
	Register(MicroscopyAnnotationPropertyType8135)
	Register(MicroscopyMeasurementType8136)
	Register(ProstateReportingSystem6310)
	Register(MRSignalIntensity6311)
	Register(CrossSectionalScanPlaneOrientation6312)
	Register(HistoryOfProstateDisease6313)
	Register(ProstateMRIStudyQualityFinding6314)
	Register(ProstateMRISeriesQualityFinding6315)
	Register(MRImagingArtifact6316)
	Register(ProstateDCEMRIQualityFinding6317)
	Register(ProstateDWIMRIQualityFinding6318)
	Register(AbdominalInterventionType6319)
	Register(AbdominopelvicIntervention6320)
	Register(ProstateCancerDiagnosticProcedure6321)
	Register(ProstateCancerFamilyHistory6322)
	Register(ProstateCancerTherapy6323)
	Register(ProstateMRIAssessment6324)
	Register(OverallAssessmentFromPIRADS6325)
	Register(ImageQualityControlStandard6326)
	Register(ProstateImagingIndication6327)
	Register(PIRADSV2LesionAssessmentCategory6328)
	Register(PIRADSV2T2WIPZLesionAssessmentCategory6329)
	Register(PIRADSV2T2WITZLesionAssessmentCategory6330)
	Register(PIRADSV2DWILesionAssessmentCategory6331)
	Register(PIRADSV2DCELesionAssessmentCategory6332)
	Register(mpMRIAssessmentType6333)
	Register(mpMRIAssessmentTypeFromPIRADS6334)
	Register(mpMRIAssessmentValue6335)
	Register(MRIAbnormality6336)
	Register(mpMRIProstateAbnormalityFromPIRADS6337)
	Register(mpMRIBenignProstateAbnormalityFromPIRADS6338)
	Register(MRIShapeCharacteristic6339)
	Register(ProstateMRIShapeCharacteristicFromPIRADS6340)
	Register(MRIMarginCharacteristic6341)
	Register(ProstateMRIMarginCharacteristicFromPIRADS6342)
	Register(MRISignalCharacteristic6343)
	Register(ProstateMRISignalCharacteristicFromPIRADS6344)
	Register(MRIEnhancementPattern6345)
	Register(ProstateMRIEnhancementPatternFromPIRADS6346)
	Register(ProstateMRIExtraProstaticFinding6347)
	Register(ProstateMRIAssessmentOfExtraProstaticAnatomicSite6348)
	Register(MRCoilType6349)
	Register(EndorectalCoilFillSubstance6350)
	Register(ProstateRelationalMeasurement6351)
	Register(ProstateCancerDiagnosticBloodLabMeasurement6352)
	Register(ProstateImagingTypesOfQualityControlStandard6353)
	Register(UltrasoundShearWaveMeasurement12308)
	Register(LeftVentricleMyocardialWall16SegmentModel3780RETIRED)
	Register(LeftVentricleMyocardialWall18SegmentModel3781)
	Register(LeftVentricleBasalWall6Segments3782)
	Register(LeftVentricleMidlevelWall6Segments3783)
	Register(LeftVentricleApicalWall4Segments3784)
	Register(LeftVentricleApicalWall6Segments3785)
	Register(PatientTreatmentPreparationMethod9571)
	Register(PatientShieldingDevice9572)
	Register(PatientTreatmentPreparationDevice9573)
	Register(PatientPositionDisplacementReferencePoint9574)
	Register(PatientAlignmentDevice9575)
	Register(ReasonsForRTRadiationTreatmentOmission9576)
	Register(PatientTreatmentPreparationProcedure9577)
	Register(MotionManagementSetupDevice9578)
	Register(CoreEchoStrainMeasurement12309)
	Register(MyocardialStrainMethod12310)
	Register(EchoMeasuredStrainProperty12311)
	Register(AssessmentFromCADRADS3020)
	Register(CADRADSStenosisAssessmentModifier3021)
	Register(CADRADSAssessmentModifier3022)
	Register(RTSegmentMaterial9579)
	Register(VertebralAnatomicStructure7602)
	Register(Vertebra7603)
	Register(IntervertebralDisc7604)
	Register(ImagingProcedure101)
	Register(NICIPShortCodeImagingProcedure103)
	Register(NICIPSNOMEDImagingProcedure104)
	Register(ICD10PCSImagingProcedure105)
	Register(ICD10PCSNuclearMedicineProcedure106)
	Register(ICD10PCSRadiationTherapyProcedure107)
	Register(RTSegmentationPropertyCategory9580)
	Register(RadiotherapyRegistrationMark9581)
	Register(RadiotherapyDoseRegion9582)
	Register(AnatomicallyLocalizedLesionSegmentationType7199)
	Register(ReasonForRemovalFromOperationalUse7031)
	Register(GeneralUltrasoundReportDocumentTitle12320)
	Register(ElastographySite12321)
	Register(ElastographyMeasurementSite12322)
	Register(UltrasoundRelevantPatientCondition12323)
	Register(ShearWaveDetectionMethod12324)
	Register(LiverUltrasoundStudyIndication12325)
	Register(AnalogWaveformFilter3042)
	Register(DigitalWaveformFilter3043)
	Register(WaveformFilterLookupTableInputFrequencyUnit3044)
	Register(WaveformFilterLookupTableOutputMagnitudeUnit3045)
	Register(SpecificObservationSubjectClass272)
	Register(MovableBeamLimitingDeviceType9540)
	Register(RadiotherapyAcquisitionWorkItemSubtasks9260)
	Register(PatientPositionAcquisitionRadiationSourceLocations9261)
	Register(EnergyDerivationTypes9262)
	Register(KVImagingAcquisitionTechniques9263)
	Register(MVImagingAcquisitionTechniques9264)
	Register(PatientPositionAcquisitionProjectionTechniques9265)
	Register(PatientPositionAcquisitionCTTechniques9266)
	Register(PatientPositioningRelatedObjectPurposes9267)
	Register(PatientPositionAcquisitionDevices9268)
	Register(RTRadiationMetersetUnits9269)
	Register(AcquisitionInitiationTypes9270)
	Register(RTImagePatientPositionAcquisitionDevices9271)
	Register(PhotoacousticIlluminationMethod11001)
	Register(AcousticCouplingMedium11002)
	Register(UltrasoundTransducerTechnology11003)
	Register(SpeedOfSoundCorrectionMechanisms11004)
	Register(PhotoacousticReconstructionAlgorithmFamily11005)
	Register(PhotoacousticImagedProperty11006)
	Register(XRayRadiationDoseProcedureTypeReported10005)
	Register(TopicalTreatment4410)
	Register(LesionColor4411)
	Register(SpecimenStainForConfocalMicroscopy4412)
	Register(RTROIImageAcquisitionContext9272)
	Register(LobeOfLung6170)
	Register(ZoneOfLung6171)
	Register(SleepStage3046)
	Register(PatientPositionAcquisitionMRTechniques9273)
	Register(RTPlanRadiotherapyProcedureTechnique9583)
	Register(WaveformAnnotationClassification3047)
	Register(WaveformAnnotationsDocumentTitle3048)
	Register(EEGProcedure3049)
	Register(PatientConsciousness3050)
	Register(FollicleType12010)
	Register(BreastSegmentationTypes7163)
	Register(ImplantedDevice3779)
	Register(SimilarityMeasure281)
	Register(WaveformAcquisitionModality34)
	Register(EnFaceProcessingAlgorithmFamily4274)
	Register(AnteriorEyeSegmentationSurface4275)
	Register(FetalEchocardiographyImageView12312)
	Register(CardiacUltrasoundFetalArrhythmiaMeasurements12313)
	Register(CommonFetalEchocardiographyMeasurements12314)
	Register(HeadAndNeckPrimaryAnatomicStructure4061)
	Register(VLView4062)
	Register(VLDentalView4063)
	Register(VLViewModifier4064)
	Register(VLDentalViewModifier4065)
	Register(OrthognathicFunctionalCondition4066)
	Register(OrthodonticFindingByInspection4067)
	Register(OrthodonticObservableEntity4068)
	Register(DentalOcclusion4069)
	Register(OrthodonticTreatmentProgress4070)
	Register(GeneralPhotographyDevice4071)
	Register(DevicesForThePurposeOfDentalPhotography4072)
	Register(CTDIPhantomDevice4053)
	Register(DiagnosticImagingProcedureWithoutIVContrast108)
	Register(DiagnosticImagingProcedureWithIVContrast109)
	Register(StructuralHeartProcedure12331)
	Register(StructuralHeartDevice12332)
	Register(StructuralHeartMeasurement12333)
	Register(AorticValveStructuralMeasurement12334)
	Register(MitralValveStructuralMeasurement12335)
	Register(TricuspidValveStructuralMeasurement12336)
	Register(StructuralHeartEchoMeasurement12337)
	Register(LeftAtrialAppendageClosureMeasurement12338)
	Register(StructuralHeartProcedureAnatomicSite12339)
	Register(IndicationForStructuralHeartProcedure12341)
	Register(BradycardiacAgent12342)
	Register(TransesophagealEchocardiographyScanPlane12343)
	Register(StructuralHeartMeasurementReportDocumentTitle12344)
	Register(PersonGenderIdentity7458)
	Register(CategoryOfSexParametersForClinicalUse7459)
	Register(ThirdPersonPronounSet7448)
	Register(CardiacStructureCalcificationQualitativeEvaluation12345)
	Register(VisualFieldMeasurements4280)
	Register(OpticDiscKeyMeasurements4281)
	Register(RetinalSectorMethods4282)
	Register(RNFLSectorMeasurements4283)
	Register(RNFLClockfaceMeasurements4284)
	Register(MacularThicknessKeyMeasurements4285)
	Register(GanglionCellMeasurementExtent4286)
	Register(GanglionCellKeyMeasurements4287)
	Register(GanglionCellSectorMeasurements4288)
	Register(GanglionCellSectorMethods4289)
	Register(EndothelialCellCountMeasurements4290)
	Register(OphthalmicImageROIMeasurements4291)
	Register(RTPlanApprovalAssertion9584)
	Register(EstimatedDeliveryDateMethod12026)
	Register(RTDoseCalculationAlgorithmFamily9585)
	Register(DoseIndexForDoseCalibration10012)
	Register(UltrasoundAttenuationImagingSite12036)
}
