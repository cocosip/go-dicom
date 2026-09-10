// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DICOM Dictionary.xml (version 2026b). DO NOT EDIT.

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
	PhotoacousticImageStorage = New("1.2.840.10008.5.1.4.1.1.6.3", "Photoacoustic Image Storage", TypeSOPClass, false)

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

	// WaveformAnnotationsDocumentTitle3048 Waveform Annotations Document Title (3048)
	WaveformAnnotationsDocumentTitle3048 = New("1.2.840.10008.6.1.1488", "Waveform Annotations Document Title (3048)", TypeContextGroupName, false)

	// EEGProcedure3049 EEG Procedure (3049)
	EEGProcedure3049 = New("1.2.840.10008.6.1.1489", "EEG Procedure (3049)", TypeContextGroupName, false)

	// PatientConsciousness3050 Patient Consciousness (3050)
	PatientConsciousness3050 = New("1.2.840.10008.6.1.1490", "Patient Consciousness (3050)", TypeContextGroupName, false)

	// FollicleType12010 Follicle Type (12010)
	FollicleType12010 = New("1.2.840.10008.6.1.1491", "Follicle Type (12010)", TypeContextGroupName, false)

	// BreastTissueSegmentationType7163 Breast Tissue Segmentation Type (7163)
	BreastTissueSegmentationType7163 = New("1.2.840.10008.6.1.1492", "Breast Tissue Segmentation Type (7163)", TypeContextGroupName, false)

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

	// FetalAnatomySurveyAssessment12040 Fetal Anatomy Survey Assessment (12040)
	FetalAnatomySurveyAssessment12040 = New("1.2.840.10008.6.1.1550", "Fetal Anatomy Survey Assessment (12040)", TypeContextGroupName, false)

	// FetalAnatomySurveyAssessmentHead12041 Fetal Anatomy Survey Assessment - Head (12041)
	FetalAnatomySurveyAssessmentHead12041 = New("1.2.840.10008.6.1.1551", "Fetal Anatomy Survey Assessment - Head (12041)", TypeContextGroupName, false)

	// FetalAnatomySurveyAssessmentFaceAndNeck12042 Fetal Anatomy Survey Assessment - Face and Neck (12042)
	FetalAnatomySurveyAssessmentFaceAndNeck12042 = New("1.2.840.10008.6.1.1552", "Fetal Anatomy Survey Assessment - Face and Neck (12042)", TypeContextGroupName, false)

	// FetalAnatomySurveyAssessmentChest12043 Fetal Anatomy Survey Assessment - Chest (12043)
	FetalAnatomySurveyAssessmentChest12043 = New("1.2.840.10008.6.1.1553", "Fetal Anatomy Survey Assessment - Chest (12043)", TypeContextGroupName, false)

	// FetalAnatomySurveyAssessmentHeart12044 Fetal Anatomy Survey Assessment - Heart (12044)
	FetalAnatomySurveyAssessmentHeart12044 = New("1.2.840.10008.6.1.1554", "Fetal Anatomy Survey Assessment - Heart (12044)", TypeContextGroupName, false)

	// FetalAnatomySurveyAssessmentAbdomenAndPelvis12045 Fetal Anatomy Survey Assessment - Abdomen and Pelvis (12045)
	FetalAnatomySurveyAssessmentAbdomenAndPelvis12045 = New("1.2.840.10008.6.1.1555", "Fetal Anatomy Survey Assessment - Abdomen and Pelvis (12045)", TypeContextGroupName, false)

	// FetalAnatomySurveyAssessmentSpine12046 Fetal Anatomy Survey Assessment - Spine (12046)
	FetalAnatomySurveyAssessmentSpine12046 = New("1.2.840.10008.6.1.1556", "Fetal Anatomy Survey Assessment - Spine (12046)", TypeContextGroupName, false)

	// FetalAnatomySurveyAssessmentExtremities12047 Fetal Anatomy Survey Assessment - Extremities (12047)
	FetalAnatomySurveyAssessmentExtremities12047 = New("1.2.840.10008.6.1.1557", "Fetal Anatomy Survey Assessment - Extremities (12047)", TypeContextGroupName, false)

	// FetalAnatomySurveyAssessmentMaternal12048 Fetal Anatomy Survey Assessment - Maternal (12048)
	FetalAnatomySurveyAssessmentMaternal12048 = New("1.2.840.10008.6.1.1558", "Fetal Anatomy Survey Assessment - Maternal (12048)", TypeContextGroupName, false)

	// FetalAnatomySurveyPracticeGuideline12049 Fetal Anatomy Survey Practice Guideline (12049)
	FetalAnatomySurveyPracticeGuideline12049 = New("1.2.840.10008.6.1.1559", "Fetal Anatomy Survey Practice Guideline (12049)", TypeContextGroupName, false)

	// SensitiveContentCategory900 Sensitive Content Category (900)
	SensitiveContentCategory900 = New("1.2.840.10008.6.1.1560", "Sensitive Content Category (900)", TypeContextGroupName, false)

	// SensitiveContentDetail901 Sensitive Content Detail (901)
	SensitiveContentDetail901 = New("1.2.840.10008.6.1.1561", "Sensitive Content Detail (901)", TypeContextGroupName, false)

	// ApplicationTypeCode406 Application Type Code (406)
	ApplicationTypeCode406 = New("1.2.840.10008.6.1.1562", "Application Type Code (406)", TypeContextGroupName, false)

	// XRayModulationType10035 X-Ray Modulation Type (10035)
	XRayModulationType10035 = New("1.2.840.10008.6.1.1563", "X-Ray Modulation Type (10035)", TypeContextGroupName, false)

	// RadiotherapyDoseRealWorldUnits9586 Radiotherapy Dose Real World Units (9586)
	RadiotherapyDoseRealWorldUnits9586 = New("1.2.840.10008.6.1.1564", "Radiotherapy Dose Real World Units (9586)", TypeContextGroupName, false)

	// RadiotherapyDoseInterpretedTypeCodes9587 Radiotherapy Dose Interpreted Type Codes (9587)
	RadiotherapyDoseInterpretedTypeCodes9587 = New("1.2.840.10008.6.1.1565", "Radiotherapy Dose Interpreted Type Codes (9587)", TypeContextGroupName, false)

	// RadiotherapyDoseInterpretedTypeModifierCodes9588 Radiotherapy Dose Interpreted Type Modifier Codes (9588)
	RadiotherapyDoseInterpretedTypeModifierCodes9588 = New("1.2.840.10008.6.1.1566", "Radiotherapy Dose Interpreted Type Modifier Codes (9588)", TypeContextGroupName, false)

	// RadiotherapyDoseIntentCodes9589 Radiotherapy Dose Intent Codes (9589)
	RadiotherapyDoseIntentCodes9589 = New("1.2.840.10008.6.1.1567", "Radiotherapy Dose Intent Codes (9589)", TypeContextGroupName, false)

	// QualitySegmentationPropertyType7164 Quality Segmentation Property Type (7164)
	QualitySegmentationPropertyType7164 = New("1.2.840.10008.6.1.1568", "Quality Segmentation Property Type (7164)", TypeContextGroupName, false)

	// UltrasoundZScorePopulationIndex12027 Ultrasound Z-Score Population Index (12027)
	UltrasoundZScorePopulationIndex12027 = New("1.2.840.10008.6.1.1569", "Ultrasound Z-Score Population Index (12027)", TypeContextGroupName, false)

	// FetalUltrasoundZScoreReferenceAuthority12028 Fetal Ultrasound Z-Score Reference Authority (12028)
	FetalUltrasoundZScoreReferenceAuthority12028 = New("1.2.840.10008.6.1.1570", "Fetal Ultrasound Z-Score Reference Authority (12028)", TypeContextGroupName, false)

	// MetalArtifactReductionAlgorithmFamily10036 Metal Artifact Reduction Algorithm Family (10036)
	MetalArtifactReductionAlgorithmFamily10036 = New("1.2.840.10008.6.1.1571", "Metal Artifact Reduction Algorithm Family (10036)", TypeContextGroupName, false)
)

var generatedStandardUIDEntries = []*UID{
	Verification,
	ImplicitVRLittleEndian,
	ExplicitVRLittleEndian,
	EncapsulatedUncompressedExplicitVRLittleEndian,
	DeflatedExplicitVRLittleEndian,
	ExplicitVRBigEndianRETIRED,
	JPEGBaseline8Bit,
	JPEGExtended12Bit,
	JPEGExtended35RETIRED,
	JPEGSpectralSelectionNonHierarchical68RETIRED,
	JPEGSpectralSelectionNonHierarchical79RETIRED,
	JPEGFullProgressionNonHierarchical1012RETIRED,
	JPEGFullProgressionNonHierarchical1113RETIRED,
	JPEGLossless,
	JPEGLosslessNonHierarchical15RETIRED,
	JPEGExtendedHierarchical1618RETIRED,
	JPEGExtendedHierarchical1719RETIRED,
	JPEGSpectralSelectionHierarchical2022RETIRED,
	JPEGSpectralSelectionHierarchical2123RETIRED,
	JPEGFullProgressionHierarchical2426RETIRED,
	JPEGFullProgressionHierarchical2527RETIRED,
	JPEGLosslessHierarchical28RETIRED,
	JPEGLosslessHierarchical29RETIRED,
	JPEGLosslessSV1,
	JPEGLSLossless,
	JPEGLSNearLossless,
	JPEG2000Lossless,
	JPEG2000,
	JPEG2000MCLossless,
	JPEG2000MC,
	JPIPReferenced,
	JPIPReferencedDeflate,
	MPEG2MPML,
	MPEG2MPMLF,
	MPEG2MPHL,
	MPEG2MPHLF,
	MPEG4HP41,
	MPEG4HP41F,
	MPEG4HP41BD,
	MPEG4HP41BDF,
	MPEG4HP422D,
	MPEG4HP422DF,
	MPEG4HP423D,
	MPEG4HP423DF,
	MPEG4HP42STEREO,
	MPEG4HP42STEREOF,
	HEVCMP51,
	HEVCM10P51,
	JPEGXLLossless,
	JPEGXLJPEGRecompression,
	JPEGXL,
	HTJ2KLossless,
	HTJ2KLosslessRPCL,
	HTJ2K,
	JPIPHTJ2KReferenced,
	JPIPHTJ2KReferencedDeflate,
	RLELossless,
	RFC2557MIMEEncapsulationRETIRED,
	XMLEncodingRETIRED,
	SMPTEST211020UncompressedProgressiveActiveVideo,
	SMPTEST211020UncompressedInterlacedActiveVideo,
	SMPTEST211030PCMDigitalAudio,
	DeflatedImageFrameCompression,
	MediaStorageDirectoryStorage,
	HotIronPalette,
	PETPalette,
	HotMetalBluePalette,
	PET20StepPalette,
	SpringPalette,
	SummerPalette,
	FallPalette,
	WinterPalette,
	BasicStudyContentNotificationRETIRED,
	Papyrus3ImplicitVRLittleEndianRETIRED,
	StorageCommitmentPushModel,
	StorageCommitmentPushModelInstance,
	StorageCommitmentPullModelRETIRED,
	StorageCommitmentPullModelInstanceRETIRED,
	ProceduralEventLogging,
	ProceduralEventLoggingInstance,
	SubstanceAdministrationLogging,
	SubstanceAdministrationLoggingInstance,
	DCMUID,
	DCM,
	MA,
	UBERON,
	ITIS_TSN,
	MGI,
	PUBCHEM_CID,
	DC,
	NYUMCCG,
	MAYONRISBSASRG,
	IBSI,
	RO,
	RADELEMENT,
	I11,
	UNS,
	RRID,
	DICOMApplicationContext,
	DetachedPatientManagementRETIRED,
	DetachedPatientManagementMetaRETIRED,
	DetachedVisitManagementRETIRED,
	DetachedStudyManagementRETIRED,
	StudyComponentManagementRETIRED,
	ModalityPerformedProcedureStep,
	ModalityPerformedProcedureStepRetrieve,
	ModalityPerformedProcedureStepNotification,
	DetachedResultsManagementRETIRED,
	DetachedResultsManagementMetaRETIRED,
	DetachedStudyManagementMetaRETIRED,
	DetachedInterpretationManagementRETIRED,
	Storage,
	BasicFilmSession,
	BasicFilmBox,
	BasicGrayscaleImageBox,
	BasicColorImageBox,
	ReferencedImageBoxRETIRED,
	BasicGrayscalePrintManagementMeta,
	ReferencedGrayscalePrintManagementMetaRETIRED,
	PrintJob,
	BasicAnnotationBox,
	Printer,
	PrinterConfigurationRetrieval,
	PrinterInstance,
	PrinterConfigurationRetrievalInstance,
	BasicColorPrintManagementMeta,
	ReferencedColorPrintManagementMetaRETIRED,
	VOILUTBox,
	PresentationLUT,
	ImageOverlayBoxRETIRED,
	BasicPrintImageOverlayBoxRETIRED,
	PrintQueueInstanceRETIRED,
	PrintQueueManagementRETIRED,
	StoredPrintStorageRETIRED,
	HardcopyGrayscaleImageStorageRETIRED,
	HardcopyColorImageStorageRETIRED,
	PullPrintRequestRETIRED,
	PullStoredPrintManagementMetaRETIRED,
	MediaCreationManagement,
	DisplaySystem,
	DisplaySystemInstance,
	ComputedRadiographyImageStorage,
	DigitalXRayImageStorageForPresentation,
	DigitalXRayImageStorageForProcessing,
	DigitalMammographyXRayImageStorageForPresentation,
	DigitalMammographyXRayImageStorageForProcessing,
	DigitalIntraOralXRayImageStorageForPresentation,
	DigitalIntraOralXRayImageStorageForProcessing,
	CTImageStorage,
	EnhancedCTImageStorage,
	LegacyConvertedEnhancedCTImageStorage,
	UltrasoundMultiFrameImageStorageRetiredRETIRED,
	UltrasoundMultiFrameImageStorage,
	MRImageStorage,
	EnhancedMRImageStorage,
	MRSpectroscopyStorage,
	EnhancedMRColorImageStorage,
	LegacyConvertedEnhancedMRImageStorage,
	NuclearMedicineImageStorageRetiredRETIRED,
	UltrasoundImageStorageRetiredRETIRED,
	UltrasoundImageStorage,
	EnhancedUSVolumeStorage,
	PhotoacousticImageStorage,
	SecondaryCaptureImageStorage,
	MultiFrameSingleBitSecondaryCaptureImageStorage,
	MultiFrameGrayscaleByteSecondaryCaptureImageStorage,
	MultiFrameGrayscaleWordSecondaryCaptureImageStorage,
	MultiFrameTrueColorSecondaryCaptureImageStorage,
	StandaloneOverlayStorageRETIRED,
	StandaloneCurveStorageRETIRED,
	WaveformStorageTrialRETIRED,
	TwelveLeadECGWaveformStorage,
	GeneralECGWaveformStorage,
	AmbulatoryECGWaveformStorage,
	General32bitECGWaveformStorage,
	HemodynamicWaveformStorage,
	CardiacElectrophysiologyWaveformStorage,
	BasicVoiceAudioWaveformStorage,
	GeneralAudioWaveformStorage,
	ArterialPulseWaveformStorage,
	RespiratoryWaveformStorage,
	MultichannelRespiratoryWaveformStorage,
	RoutineScalpElectroencephalogramWaveformStorage,
	ElectromyogramWaveformStorage,
	ElectrooculogramWaveformStorage,
	SleepElectroencephalogramWaveformStorage,
	BodyPositionWaveformStorage,
	WaveformPresentationStateStorage,
	WaveformAcquisitionPresentationStateStorage,
	StandaloneModalityLUTStorageRETIRED,
	StandaloneVOILUTStorageRETIRED,
	GrayscaleSoftcopyPresentationStateStorage,
	ColorSoftcopyPresentationStateStorage,
	PseudoColorSoftcopyPresentationStateStorage,
	BlendingSoftcopyPresentationStateStorage,
	XAXRFGrayscaleSoftcopyPresentationStateStorage,
	GrayscalePlanarMPRVolumetricPresentationStateStorage,
	CompositingPlanarMPRVolumetricPresentationStateStorage,
	AdvancedBlendingPresentationStateStorage,
	VolumeRenderingVolumetricPresentationStateStorage,
	SegmentedVolumeRenderingVolumetricPresentationStateStorage,
	MultipleVolumeRenderingVolumetricPresentationStateStorage,
	VariableModalityLUTSoftcopyPresentationStateStorage,
	XRayAngiographicImageStorage,
	EnhancedXAImageStorage,
	XRayRadiofluoroscopicImageStorage,
	EnhancedXRFImageStorage,
	XRayAngiographicBiPlaneImageStorageRETIRED,
	XRay3DAngiographicImageStorage,
	XRay3DCraniofacialImageStorage,
	BreastTomosynthesisImageStorage,
	BreastProjectionXRayImageStorageForPresentation,
	BreastProjectionXRayImageStorageForProcessing,
	IntravascularOpticalCoherenceTomographyImageStorageForPresentation,
	IntravascularOpticalCoherenceTomographyImageStorageForProcessing,
	NuclearMedicineImageStorage,
	ParametricMapStorage,
	RawDataStorage,
	SpatialRegistrationStorage,
	SpatialFiducialsStorage,
	DeformableSpatialRegistrationStorage,
	SegmentationStorage,
	SurfaceSegmentationStorage,
	TractographyResultsStorage,
	LabelMapSegmentationStorage,
	HeightMapSegmentationStorage,
	RealWorldValueMappingStorage,
	SurfaceScanMeshStorage,
	SurfaceScanPointCloudStorage,
	VLImageStorageTrialRETIRED,
	VLMultiFrameImageStorageTrialRETIRED,
	VLEndoscopicImageStorage,
	VideoEndoscopicImageStorage,
	VLMicroscopicImageStorage,
	VideoMicroscopicImageStorage,
	VLSlideCoordinatesMicroscopicImageStorage,
	VLPhotographicImageStorage,
	VideoPhotographicImageStorage,
	OphthalmicPhotography8BitImageStorage,
	OphthalmicPhotography16BitImageStorage,
	StereometricRelationshipStorage,
	OphthalmicTomographyImageStorage,
	WideFieldOphthalmicPhotographyStereographicProjectionImageStorage,
	WideFieldOphthalmicPhotography3DCoordinatesImageStorage,
	OphthalmicOpticalCoherenceTomographyEnFaceImageStorage,
	OphthalmicOpticalCoherenceTomographyBscanVolumeAnalysisStorage,
	VLWholeSlideMicroscopyImageStorage,
	DermoscopicPhotographyImageStorage,
	ConfocalMicroscopyImageStorage,
	ConfocalMicroscopyTiledPyramidalImageStorage,
	LensometryMeasurementsStorage,
	AutorefractionMeasurementsStorage,
	KeratometryMeasurementsStorage,
	SubjectiveRefractionMeasurementsStorage,
	VisualAcuityMeasurementsStorage,
	SpectaclePrescriptionReportStorage,
	OphthalmicAxialMeasurementsStorage,
	IntraocularLensCalculationsStorage,
	MacularGridThicknessAndVolumeReportStorage,
	OphthalmicVisualFieldStaticPerimetryMeasurementsStorage,
	OphthalmicThicknessMapStorage,
	CornealTopographyMapStorage,
	TextSRStorageTrialRETIRED,
	AudioSRStorageTrialRETIRED,
	DetailSRStorageTrialRETIRED,
	ComprehensiveSRStorageTrialRETIRED,
	BasicTextSRStorage,
	EnhancedSRStorage,
	ComprehensiveSRStorage,
	Comprehensive3DSRStorage,
	ExtensibleSRStorage,
	ProcedureLogStorage,
	MammographyCADSRStorage,
	KeyObjectSelectionDocumentStorage,
	ChestCADSRStorage,
	XRayRadiationDoseSRStorage,
	RadiopharmaceuticalRadiationDoseSRStorage,
	ColonCADSRStorage,
	ImplantationPlanSRStorage,
	AcquisitionContextSRStorage,
	SimplifiedAdultEchoSRStorage,
	PatientRadiationDoseSRStorage,
	PlannedImagingAgentAdministrationSRStorage,
	PerformedImagingAgentAdministrationSRStorage,
	EnhancedXRayRadiationDoseSRStorage,
	WaveformAnnotationSRStorage,
	ContentAssessmentResultsStorage,
	MicroscopyBulkSimpleAnnotationsStorage,
	EncapsulatedPDFStorage,
	EncapsulatedCDAStorage,
	EncapsulatedSTLStorage,
	EncapsulatedOBJStorage,
	EncapsulatedMTLStorage,
	PositronEmissionTomographyImageStorage,
	LegacyConvertedEnhancedPETImageStorage,
	StandalonePETCurveStorageRETIRED,
	EnhancedPETImageStorage,
	BasicStructuredDisplayStorage,
	CTDefinedProcedureProtocolStorage,
	CTPerformedProcedureProtocolStorage,
	ProtocolApprovalStorage,
	ProtocolApprovalInformationModelFind,
	ProtocolApprovalInformationModelMove,
	ProtocolApprovalInformationModelGet,
	XADefinedProcedureProtocolStorage,
	XAPerformedProcedureProtocolStorage,
	InventoryStorage,
	InventoryFind,
	InventoryMove,
	InventoryGet,
	InventoryCreation,
	RepositoryQuery,
	StorageManagementInstance,
	RTImageStorage,
	RTDoseStorage,
	RTStructureSetStorage,
	RTBeamsTreatmentRecordStorage,
	RTPlanStorage,
	RTBrachyTreatmentRecordStorage,
	RTTreatmentSummaryRecordStorage,
	RTIonPlanStorage,
	RTIonBeamsTreatmentRecordStorage,
	RTPhysicianIntentStorage,
	RTSegmentAnnotationStorage,
	RTRadiationSetStorage,
	CArmPhotonElectronRadiationStorage,
	TomotherapeuticRadiationStorage,
	RoboticArmRadiationStorage,
	RTRadiationRecordSetStorage,
	RTRadiationSalvageRecordStorage,
	TomotherapeuticRadiationRecordStorage,
	CArmPhotonElectronRadiationRecordStorage,
	RoboticRadiationRecordStorage,
	RTRadiationSetDeliveryInstructionStorage,
	RTTreatmentPreparationStorage,
	EnhancedRTImageStorage,
	EnhancedContinuousRTImageStorage,
	RTPatientPositionAcquisitionInstructionStorage,
	DICOSCTImageStorage,
	DICOSDigitalXRayImageStorageForPresentation,
	DICOSDigitalXRayImageStorageForProcessing,
	DICOSThreatDetectionReportStorage,
	DICOS2DAITStorage,
	DICOS3DAITStorage,
	DICOSQuadrupoleResonanceStorage,
	EddyCurrentImageStorage,
	EddyCurrentMultiFrameImageStorage,
	ThermographyImageStorage,
	ThermographyMultiFrameImageStorage,
	UltrasoundWaveformStorage,
	PatientRootQueryRetrieveInformationModelFind,
	PatientRootQueryRetrieveInformationModelMove,
	PatientRootQueryRetrieveInformationModelGet,
	StudyRootQueryRetrieveInformationModelFind,
	StudyRootQueryRetrieveInformationModelMove,
	StudyRootQueryRetrieveInformationModelGet,
	PatientStudyOnlyQueryRetrieveInformationModelFindRETIRED,
	PatientStudyOnlyQueryRetrieveInformationModelMoveRETIRED,
	PatientStudyOnlyQueryRetrieveInformationModelGetRETIRED,
	CompositeInstanceRootRetrieveMove,
	CompositeInstanceRootRetrieveGet,
	CompositeInstanceRetrieveWithoutBulkDataGet,
	DefinedProcedureProtocolInformationModelFind,
	DefinedProcedureProtocolInformationModelMove,
	DefinedProcedureProtocolInformationModelGet,
	ModalityWorklistInformationModelFind,
	GeneralPurposeWorklistManagementMetaRETIRED,
	GeneralPurposeWorklistInformationModelFindRETIRED,
	GeneralPurposeScheduledProcedureStepRETIRED,
	GeneralPurposePerformedProcedureStepRETIRED,
	InstanceAvailabilityNotification,
	RTBeamsDeliveryInstructionStorageTrialRETIRED,
	RTConventionalMachineVerificationTrialRETIRED,
	RTIonMachineVerificationTrialRETIRED,
	UnifiedWorklistAndProcedureStepTrialRETIRED,
	UnifiedProcedureStepPushTrialRETIRED,
	UnifiedProcedureStepWatchTrialRETIRED,
	UnifiedProcedureStepPullTrialRETIRED,
	UnifiedProcedureStepEventTrialRETIRED,
	UPSGlobalSubscriptionInstance,
	UPSFilteredGlobalSubscriptionInstance,
	UnifiedWorklistAndProcedureStep,
	UnifiedProcedureStepPush,
	UnifiedProcedureStepWatch,
	UnifiedProcedureStepPull,
	UnifiedProcedureStepEvent,
	UnifiedProcedureStepQuery,
	RTBeamsDeliveryInstructionStorage,
	RTConventionalMachineVerification,
	RTIonMachineVerification,
	RTBrachyApplicationSetupDeliveryInstructionStorage,
	GeneralRelevantPatientInformationQuery,
	BreastImagingRelevantPatientInformationQuery,
	CardiacRelevantPatientInformationQuery,
	HangingProtocolStorage,
	HangingProtocolInformationModelFind,
	HangingProtocolInformationModelMove,
	HangingProtocolInformationModelGet,
	ColorPaletteStorage,
	ColorPaletteQueryRetrieveInformationModelFind,
	ColorPaletteQueryRetrieveInformationModelMove,
	ColorPaletteQueryRetrieveInformationModelGet,
	ProductCharacteristicsQuery,
	SubstanceApprovalQuery,
	GenericImplantTemplateStorage,
	GenericImplantTemplateInformationModelFind,
	GenericImplantTemplateInformationModelMove,
	GenericImplantTemplateInformationModelGet,
	ImplantAssemblyTemplateStorage,
	ImplantAssemblyTemplateInformationModelFind,
	ImplantAssemblyTemplateInformationModelMove,
	ImplantAssemblyTemplateInformationModelGet,
	ImplantTemplateGroupStorage,
	ImplantTemplateGroupInformationModelFind,
	ImplantTemplateGroupInformationModelMove,
	ImplantTemplateGroupInformationModelGet,
	NativeDICOMModel,
	AbstractMultiDimensionalImageModel,
	DICOMContentMappingResource,
	VideoEndoscopicImageRealTimeCommunication,
	VideoPhotographicImageRealTimeCommunication,
	AudioWaveformRealTimeCommunication,
	RenditionSelectionDocumentRealTimeCommunication,
	dicomDeviceName,
	dicomDescription,
	dicomManufacturer,
	dicomManufacturerModelName,
	dicomSoftwareVersion,
	dicomVendorData,
	dicomAETitle,
	dicomNetworkConnectionReference,
	dicomApplicationCluster,
	dicomAssociationInitiator,
	dicomAssociationAcceptor,
	dicomHostname,
	dicomPort,
	dicomSOPClass,
	dicomTransferRole,
	dicomTransferSyntax,
	dicomPrimaryDeviceType,
	dicomRelatedDeviceReference,
	dicomPreferredCalledAETitle,
	dicomTLSCyphersuite,
	dicomAuthorizedNodeCertificateReference,
	dicomThisNodeCertificateReference,
	dicomInstalled,
	dicomStationName,
	dicomDeviceSerialNumber,
	dicomInstitutionName,
	dicomInstitutionAddress,
	dicomInstitutionDepartmentName,
	dicomIssuerOfPatientID,
	dicomPreferredCallingAETitle,
	dicomSupportedCharacterSet,
	dicomConfigurationRoot,
	dicomDevicesRoot,
	dicomUniqueAETitlesRegistryRoot,
	dicomDevice,
	dicomNetworkAE,
	dicomNetworkConnection,
	dicomUniqueAETitle,
	dicomTransferCapability,
	UTC,
	AnatomicModifier2,
	AnatomicRegion4,
	TransducerApproach5,
	TransducerOrientation6,
	UltrasoundBeamPath7,
	AngiographicInterventionalDevice8,
	ImageGuidedTherapeuticProcedure9,
	InterventionalDrug10,
	AdministrationRoute11,
	ImagingContrastAgent12,
	ImagingContrastAgentIngredient13,
	RadiopharmaceuticalIsotope18,
	PatientOrientation19,
	PatientOrientationModifier20,
	PatientEquipmentRelationship21,
	CranioCaudadAngulation23,
	Radiopharmaceutical25,
	NuclearMedicineProjection26,
	AcquisitionModality29,
	DICOMDevice30,
	AbstractPrior31,
	NumericValueQualifier42,
	MeasurementUnit82,
	RealWorldValueMappingUnit83,
	SignificanceLevel220,
	MeasurementRangeConcept221,
	Normality222,
	NormalRangeValue223,
	SelectionMethod224,
	MeasurementUncertaintyConcept225,
	PopulationStatisticalDescriptor226,
	SampleStatisticalDescriptor227,
	EquationOrTable228,
	YesNo230,
	PresentAbsent240,
	NormalAbnormal242,
	Laterality244,
	PositiveNegative250,
	ComplicationSeverity251,
	ObserverType270,
	ObservationSubjectClass271,
	AudioChannelSource3000,
	ECGLead3001,
	HemodynamicWaveformSource3003,
	CardiovascularAnatomicStructure3010,
	ElectrophysiologyAnatomicLocation3011,
	CoronaryArterySegment3014,
	CoronaryArtery3015,
	CardiovascularAnatomicStructureModifier3019,
	CardiologyMeasurementUnit3082RETIRED,
	TimeSynchronizationChannelType3090,
	CardiacProceduralStateValue3101,
	ElectrophysiologyMeasurementFunctionTechnique3240,
	HemodynamicMeasurementTechnique3241,
	CatheterizationProcedurePhase3250,
	ElectrophysiologyProcedurePhase3254,
	StressProtocol3261,
	ECGPatientStateValue3262,
	ElectrodePlacementValue3263,
	XYZElectrodePlacementValues3264RETIRED,
	HemodynamicPhysiologicalChallenge3271,
	ECGAnnotation3335,
	HemodynamicAnnotation3337,
	ElectrophysiologyAnnotation3339,
	ProcedureLogTitle3400,
	LogNoteType3401,
	PatientStatusAndEvent3402,
	PercutaneousEntry3403,
	StaffAction3404,
	ProcedureActionValue3405,
	NonCoronaryTranscatheterIntervention3406,
	ObjectReferencePurpose3407,
	ConsumableAction3408,
	DrugContrastAdministration3409,
	DrugContrastNumericParameter3410,
	IntracoronaryDevice3411,
	InterventionActionStatus3412,
	AdverseOutcome3413,
	ProcedureUrgency3414,
	CardiacRhythm3415,
	RespirationRhythm3416,
	LesionRisk3418,
	FindingTitle3419,
	ProcedureAction3421,
	DeviceUseAction3422,
	NumericDeviceCharacteristic3423,
	InterventionParameter3425,
	ConsumablesParameter3426,
	EquipmentEvent3427,
	CardiovascularImagingProcedure3428,
	CatheterizationDevice3429,
	DateTimeQualifier3430,
	PeripheralPulseLocation3440,
	PatientAssessment3441,
	PeripheralPulseMethod3442,
	SkinCondition3446,
	AirwayAssessment3448,
	CalibrationObject3451,
	CalibrationMethod3452,
	CardiacVolumeMethod3453,
	IndexMethod3455,
	SubSegmentMethod3456,
	ContourRealignment3458,
	CircumferentialExtent3460,
	RegionalExtent3461,
	ChamberIdentification3462,
	QAReferenceMethod3465,
	PlaneIdentification3466,
	EjectionFraction3467,
	EDVolume3468,
	ESVolume3469,
	VesselLumenCrossSectionalAreaCalculationMethod3470,
	EstimatedVolume3471,
	CardiacContractionPhase3472,
	IVUSProcedurePhase3480,
	IVUSDistanceMeasurement3481,
	IVUSAreaMeasurement3482,
	IVUSLongitudinalMeasurement3483,
	IVUSIndexRatio3484,
	IVUSVolumeMeasurement3485,
	VascularMeasurementSite3486,
	IntravascularVolumetricRegion3487,
	MinMaxMean3488,
	CalciumDistribution3489,
	IVUSLesionMorphology3491,
	VascularDissectionClassification3492,
	IVUSRelativeStenosisSeverity3493,
	IVUSNonMorphologicalFinding3494,
	IVUSPlaqueComposition3495,
	IVUSFiducialPoint3496,
	IVUSArterialMorphology3497,
	PressureUnit3500,
	HemodynamicResistanceUnit3502,
	IndexedHemodynamicResistanceUnit3503,
	CatheterSizeUnit3510,
	SpecimenCollection3515,
	BloodSourceType3520,
	BloodGasPressure3524,
	BloodGasContent3525,
	BloodGasSaturation3526,
	BloodBaseExcess3527,
	BloodPH3528,
	ArterialVenousContent3529,
	OxygenAdministrationAction3530,
	OxygenAdministration3531,
	CirculatorySupportAction3550,
	VentilationAction3551,
	PacingAction3552,
	CirculatorySupport3553,
	Ventilation3554,
	Pacing3555,
	BloodPressureMethod3560,
	RelativeTime3600,
	HemodynamicPatientState3602,
	ArterialLesionLocation3604,
	ArterialSourceLocation3606,
	VenousSourceLocation3607,
	AtrialSourceLocation3608,
	VentricularSourceLocation3609,
	GradientSourceLocation3610,
	PressureMeasurement3611,
	BloodVelocityMeasurement3612,
	HemodynamicTimeMeasurement3613,
	NonMitralValveArea3614,
	ValveArea3615,
	HemodynamicPeriodMeasurement3616,
	ValveFlow3617,
	HemodynamicFlow3618,
	HemodynamicResistanceMeasurement3619,
	HemodynamicRatio3620,
	FractionalFlowReserve3621,
	MeasurementType3627,
	CardiacOutputMethod3628,
	ProcedureIntent3629,
	CardiovascularAnatomicLocation3630,
	Hypertension3640,
	HemodynamicAssessment3641,
	DegreeFinding3642,
	HemodynamicMeasurementPhase3651,
	BodySurfaceAreaEquation3663,
	OxygenConsumptionEquationTable3664,
	P50Equation3666,
	FraminghamScore3667,
	FraminghamTable3668,
	ECGProcedureType3670,
	ReasonForECGStudy3671,
	Pacemaker3672,
	Diagnosis3673RETIRED,
	OtherFilters3675RETIRED,
	LeadMeasurementTechnique3676,
	SummaryCodesECG3677,
	QTCorrectionAlgorithm3678,
	ECGMorphologyDescription3679RETIRED,
	ECGLeadNoiseDescription3680,
	ECGLeadNoiseModifier3681RETIRED,
	Probability3682RETIRED,
	Modifier3683RETIRED,
	Trend3684RETIRED,
	ConjunctiveTerm3685RETIRED,
	ECGInterpretiveStatement3686RETIRED,
	ElectrophysiologyWaveformDuration3687,
	ElectrophysiologyWaveformVoltage3688,
	CathDiagnosis3700,
	CardiacValveTract3701,
	WallMotion3703,
	MyocardiumWallMorphologyFinding3704,
	ChamberSize3705,
	OverallContractility3706,
	VSDDescription3707,
	AorticRootDescription3709,
	CoronaryDominance3710,
	ValvularAbnormality3711,
	VesselDescriptor3712,
	TIMIFlowCharacteristic3713,
	Thrombus3714,
	LesionMargin3715,
	Severity3716,
	LeftVentricleMyocardialWall17SegmentModel3717,
	MyocardialWallSegmentsInProjection3718,
	CanadianClinicalClassification3719,
	CardiacHistoryDate3720RETIRED,
	CardiovascularSurgery3721,
	DiabeticTherapy3722,
	MIType3723,
	SmokingHistory3724,
	CoronaryInterventionIndication3726,
	CatheterizationIndication3727,
	CathFinding3728,
	AdmissionStatus3729,
	InsurancePayor3730,
	PrimaryCauseOfDeath3733,
	AcuteCoronarySyndromeTimePeriod3735,
	NYHAClassification3736,
	IschemiaNonInvasiveTest3737,
	PreCathAnginaType3738,
	CathProcedureType3739,
	ThrombolyticAdministration3740,
	LabVisitMedicationAdministration3741,
	PCIMedicationAdministration3742,
	ClopidogrelTiclopidineAdministration3743,
	EFTestingMethod3744,
	CalculationMethod3745,
	PercutaneousEntrySite3746,
	PercutaneousClosure3747,
	AngiographicEFTestingMethod3748,
	PCIProcedureResult3749,
	PreviouslyDilatedLesion3750,
	GuidewireCrossing3752,
	VascularComplication3754,
	CathComplication3755,
	CardiacPatientRiskFactor3756,
	CardiacDiagnosticProcedure3757,
	CardiovascularFamilyHistory3758,
	HypertensionTherapy3760,
	AntilipemicAgent3761,
	AntiarrhythmicAgent3762,
	MyocardialInfarctionTherapy3764,
	ConcernType3769,
	ProblemStatus3770,
	HealthStatus3772,
	UseStatus3773,
	SocialHistory3774,
	CardiovascularImplant3777,
	PlaqueStructure3802,
	StenosisMeasurementMethod3804,
	StenosisType3805,
	StenosisShape3806,
	VolumeMeasurementMethod3807,
	AneurysmType3808,
	AssociatedCondition3809,
	VascularMorphology3810,
	StentFinding3813,
	StentComposition3814,
	SourceOfVascularFinding3815,
	VascularSclerosisType3817,
	NonInvasiveVascularProcedure3820,
	PapillaryMuscleIncludedExcluded3821,
	RespiratoryStatus3823,
	HeartRhythm3826,
	VesselSegment3827,
	PulmonaryArtery3829,
	StenosisLength3831,
	StenosisGrade3832,
	CardiacEjectionFraction3833,
	CardiacVolumeMeasurement3835,
	TimeBasedPerfusionMeasurement3836,
	FiducialFeature3837,
	DiameterDerivation3838,
	CoronaryVein3839,
	PulmonaryVein3840,
	MyocardialSubsegment3843,
	PartialViewSectionForMammography4005,
	DXAnatomyImaged4009,
	DXView4010,
	DXViewModifier4011,
	ProjectionEponymousName4012,
	AnatomicRegionForMammography4013,
	ViewForMammography4014,
	ViewModifierForMammography4015,
	AnatomicRegionForIntraOralRadiography4016,
	AnatomicRegionModifierForIntraOralRadiography4017,
	PrimaryAnatomicStructureForIntraOralRadiographyPermanentDentitionDesignationOfTeeth4018,
	PrimaryAnatomicStructureForIntraOralRadiographyDeciduousDentitionDesignationOfTeeth4019,
	PETRadionuclide4020,
	PETRadiopharmaceutical4021,
	CraniofacialAnatomicRegion4028,
	CTMRAndPETAnatomyImaged4030,
	CommonAnatomicRegion4031,
	MRSpectroscopyMetabolite4032,
	MRProtonSpectroscopyMetabolite4033,
	EndoscopyAnatomicRegion4040,
	XAXRFAnatomyImaged4042,
	DrugOrContrastAgentCharacteristic4050,
	GeneralDevice4051,
	PhantomDevice4052,
	OphthalmicImagingAgent4200,
	PatientEyeMovementCommand4201,
	OphthalmicPhotographyAcquisitionDevice4202,
	OphthalmicPhotographyIllumination4203,
	OphthalmicFilter4204,
	OphthalmicLens4205,
	OphthalmicChannelDescription4206,
	OphthalmicImagePosition4207,
	MydriaticAgent4208,
	OphthalmicAnatomicStructureImaged4209,
	OphthalmicTomographyAcquisitionDevice4210,
	OphthalmicOCTAnatomicStructureImaged4211,
	Language5000,
	Country5001,
	OverallBreastComposition6000,
	OverallBreastCompositionFromBIRADS6001,
	ChangeSinceLastMammogramOrPriorSurgery6002,
	ChangeSinceLastMammogramOrPriorSurgeryFromBIRADS6003,
	MammographyShapeCharacteristic6004,
	ShapeCharacteristicFromBIRADS6005,
	MammographyMarginCharacteristic6006,
	MarginCharacteristicFromBIRADS6007,
	DensityModifier6008,
	DensityModifierFromBIRADS6009,
	MammographyCalcificationType6010,
	CalcificationTypeFromBIRADS6011,
	CalcificationDistributionModifier6012,
	CalcificationDistributionModifierFromBIRADS6013,
	MammographySingleImageFinding6014,
	SingleImageFindingFromBIRADS6015,
	MammographyCompositeFeature6016,
	CompositeFeatureFromBIRADS6017,
	ClockfaceLocationOrRegion6018,
	ClockfaceLocationOrRegionFromBIRADS6019,
	QuadrantLocation6020,
	QuadrantLocationFromBIRADS6021,
	Side6022,
	SideFromBIRADS6023,
	Depth6024,
	DepthFromBIRADS6025,
	MammographyAssessment6026,
	AssessmentFromBIRADS6027,
	MammographyRecommendedFollowUp6028,
	RecommendedFollowUpFromBIRADS6029,
	MammographyPathologyCode6030,
	BenignPathologyCodeFromBIRADS6031,
	HighRiskLesionPathologyCodeFromBIRADS6032,
	MalignantPathologyCodeFromBIRADS6033,
	CADOutputIntendedUse6034,
	CompositeFeatureRelation6035,
	FeatureScope6036,
	MammographyQuantitativeTemporalDifferenceType6037,
	MammographyQualitativeTemporalDifferenceType6038,
	NippleCharacteristic6039,
	NonLesionObjectType6040,
	MammographyImageQualityFinding6041,
	ResultStatus6042,
	MammographyCADAnalysisType6043,
	ImageQualityAssessmentType6044,
	MammographyQualityControlStandardType6045,
	FollowUpIntervalUnit6046,
	CADProcessingAndFindingSummary6047,
	CADOperatingPointAxisLabel6048,
	BreastProcedureReported6050,
	BreastProcedureReason6051,
	BreastImagingReportSectionTitle6052,
	BreastImagingReportElement6053,
	BreastImagingFinding6054,
	BreastClinicalFindingOrIndicatedProblem6055,
	AssociatedFindingForBreast6056,
	DuctographyFindingForBreast6057,
	ProcedureModifiersForBreast6058,
	BreastImplantType6059,
	BreastBiopsyTechnique6060,
	BreastImagingProcedureModifier6061,
	InterventionalProcedureComplication6062,
	InterventionalProcedureResult6063,
	UltrasoundFindingForBreast6064,
	InstrumentApproach6065,
	TargetConfirmation6066,
	FluidColor6067,
	TumorStagesFromAJCC6068,
	NottinghamCombinedHistologicGrade6069,
	BloomRichardsonHistologicGrade6070,
	HistologicGradingMethod6071,
	BreastImplantFinding6072,
	GynecologicalHormone6080,
	BreastCancerRiskFactor6081,
	GynecologicalProcedure6082,
	ProceduresForBreast6083,
	MammoplastyProcedure6084,
	TherapiesForBreast6085,
	MenopausalPhase6086,
	GeneralRiskFactor6087,
	OBGYNMaternalRiskFactor6088,
	Substance6089,
	RelativeUsageExposureAmount6090,
	RelativeFrequencyOfEventValue6091,
	UsageExposureQualitativeConcept6092,
	UsageExposureAmountQualitativeConcept6093,
	UsageExposureFrequencyQualitativeConcept6094,
	ProcedureNumericProperty6095,
	PregnancyStatus6096,
	SideOfFamily6097,
	ChestComponentCategory6100,
	ChestFindingOrFeature6101,
	ChestFindingOrFeatureModifier6102,
	AbnormalLinesFindingOrFeature6103,
	AbnormalOpacityFindingOrFeature6104,
	AbnormalLucencyFindingOrFeature6105,
	AbnormalTextureFindingOrFeature6106,
	WidthDescriptor6107,
	ChestAnatomicStructureAbnormalDistribution6108,
	RadiographicAnatomyFindingOrFeature6109,
	LungAnatomyFindingOrFeature6110,
	BronchovascularAnatomyFindingOrFeature6111,
	PleuraAnatomyFindingOrFeature6112,
	MediastinumAnatomyFindingOrFeature6113,
	OsseousAnatomyFindingOrFeature6114,
	OsseousAnatomyModifier6115,
	MuscularAnatomy6116,
	VascularAnatomy6117,
	SizeDescriptor6118,
	ChestBorderShape6119,
	ChestBorderDefinition6120,
	ChestOrientationDescriptor6121,
	ChestContentDescriptor6122,
	ChestOpacityDescriptor6123,
	LocationInChest6124,
	GeneralChestLocation6125,
	LocationInLung6126,
	SegmentLocationInLung6127,
	ChestDistributionDescriptor6128,
	ChestSiteInvolvement6129,
	SeverityDescriptor6130,
	ChestTextureDescriptor6131,
	ChestCalcificationDescriptor6132,
	ChestQuantitativeTemporalDifferenceType6133,
	ChestQualitativeTemporalDifferenceType6134,
	ImageQualityFinding6135,
	ChestTypesOfQualityControlStandard6136,
	CADAnalysisType6137,
	ChestNonLesionObjectType6138,
	NonLesionModifier6139,
	CalculationMethod6140,
	AttenuationCoefficientMeasurement6141,
	CalculatedValue6142,
	LesionResponse6143,
	RECISTDefinedLesionResponse6144,
	BaselineCategory6145,
	BackgroundEchotexture6151,
	Orientation6152,
	LesionBoundary6153,
	EchoPattern6154,
	PosteriorAcousticFeature6155,
	Vascularity6157,
	CorrelationToOtherFinding6158,
	MalignancyType6159,
	BreastPrimaryTumorAssessmentFromAJCC6160,
	PathologicalRegionalLymphNodeAssessmentForBreast6161,
	AssessmentOfMetastasisForBreast6162,
	MenstrualCyclePhase6163,
	TimeInterval6164,
	BreastLinearMeasurement6165,
	CADGeometrySecondaryGraphicalRepresentation6166,
	DiagnosticImagingReportDocumentTitle7000,
	DiagnosticImagingReportHeading7001,
	DiagnosticImagingReportElement7002,
	DiagnosticImagingReportPurposeOfReference7003,
	WaveformPurposeOfReference7004,
	ContributingEquipmentPurposeOfReference7005,
	SRDocumentPurposeOfReference7006,
	SignaturePurpose7007,
	MediaImport7008,
	KeyObjectSelectionDocumentTitle7010,
	RejectedForQualityReason7011,
	BestInSet7012,
	DocumentTitle7020,
	RCSRegistrationMethodType7100,
	BrainAtlasFiducial7101,
	SegmentationPropertyCategory7150,
	SegmentationPropertyType7151,
	CardiacStructureSegmentationType7152,
	CNSSegmentationType7153,
	AbdominalSegmentationType7154,
	ThoracicSegmentationType7155,
	VascularSegmentationType7156,
	DeviceSegmentationType7157,
	ArtifactSegmentationType7158,
	LesionSegmentationType7159,
	PelvicOrganSegmentationType7160,
	PhysiologySegmentationType7161,
	ReferencedImagePurposeOfReference7201,
	SourceImagePurposeOfReference7202,
	ImageDerivation7203,
	PurposeOfReferenceToAlternateRepresentation7205,
	RelatedSeriesPurposeOfReference7210,
	MultiFrameSubsetType7250,
	PersonRole7450,
	FamilyMember7451,
	OrganizationalRole7452,
	PerformingRole7453,
	AnimalTaxonomicRankValue7454,
	Sex7455,
	AgeUnit7456,
	LinearMeasurementUnit7460,
	AreaMeasurementUnit7461,
	VolumeMeasurementUnit7462,
	LinearMeasurement7470,
	AreaMeasurement7471,
	VolumeMeasurement7472,
	GeneralAreaCalculationMethod7473,
	GeneralVolumeCalculationMethod7474,
	Breed7480,
	BreedRegistry7481,
	WorkitemDefinition9231,
	NonDICOMOutputTypes9232RETIRED,
	ProcedureDiscontinuationReason9300,
	ScopeOfAccumulation10000,
	UIDType10001,
	IrradiationEventType10002,
	EquipmentPlaneIdentification10003,
	FluoroMode10004,
	XRayFilterMaterial10006,
	XRayFilterType10007,
	DoseRelatedDistanceMeasurement10008,
	MeasuredCalculated10009,
	DoseMeasurementDevice10010,
	EffectiveDoseEvaluationMethod10011,
	CTAcquisitionType10013,
	CTIVContrastImagingTechnique10014,
	CTDoseReferenceAuthority10015,
	AnodeTargetMaterial10016,
	XRayGrid10017,
	UltrasoundProtocolType12001,
	UltrasoundProtocolStageType12002,
	OBGYNDate12003,
	FetalBiometryRatio12004,
	FetalBiometryMeasurement12005,
	FetalLongBonesBiometryMeasurement12006,
	FetalCraniumMeasurement12007,
	OBGYNAmnioticSacMeasurement12008,
	EarlyGestationBiometryMeasurement12009,
	UltrasoundPelvisAndUterusMeasurement12011,
	OBEquationTable12012,
	GestationalAgeEquationTable12013,
	OBFetalBodyWeightEquationTable12014,
	FetalGrowthEquationTable12015,
	EstimatedFetalWeightPercentileEquationTable12016,
	GrowthDistributionRank12017,
	OBGYNSummary12018,
	OBGYNFetusSummary12019,
	VascularSummary12101,
	TemporalPeriodRelatingToProcedureOrTherapy12102,
	VascularUltrasoundAnatomicLocation12103,
	ExtracranialArtery12104,
	IntracranialCerebralVessel12105,
	IntracranialCerebralVesselUnilateral12106,
	UpperExtremityArtery12107,
	UpperExtremityVein12108,
	LowerExtremityArtery12109,
	LowerExtremityVein12110,
	AbdominopelvicArteryPaired12111,
	AbdominopelvicArteryUnpaired12112,
	AbdominopelvicVeinPaired12113,
	AbdominopelvicVeinUnpaired12114,
	RenalVessel12115,
	VesselSegmentModifier12116,
	VesselBranchModifier12117,
	VascularUltrasoundProperty12119,
	UltrasoundBloodVelocityMeasurement12120,
	VascularIndexRatio12121,
	OtherVascularProperty12122,
	CarotidRatio12123,
	RenalRatio12124,
	PelvicVasculatureAnatomicalLocation12140,
	FetalVasculatureAnatomicalLocation12141,
	EchocardiographyLeftVentricleMeasurement12200,
	LeftVentricleLinearMeasurement12201,
	LeftVentricleVolumeMeasurement12202,
	LeftVentricleOtherMeasurement12203,
	EchocardiographyRightVentricleMeasurement12204,
	EchocardiographyLeftAtriumMeasurement12205,
	EchocardiographyRightAtriumMeasurement12206,
	EchocardiographyMitralValveMeasurement12207,
	EchocardiographyTricuspidValveMeasurement12208,
	EchocardiographyPulmonicValveMeasurement12209,
	EchocardiographyPulmonaryArteryMeasurement12210,
	EchocardiographyAorticValveMeasurement12211,
	EchocardiographyAortaMeasurement12212,
	EchocardiographyPulmonaryVeinMeasurement12214,
	EchocardiographyVenaCavaMeasurement12215,
	EchocardiographyHepaticVeinMeasurement12216,
	EchocardiographyCardiacShuntMeasurement12217,
	EchocardiographyCongenitalAnomalyMeasurement12218,
	PulmonaryVeinModifier12219,
	EchocardiographyCommonMeasurement12220,
	FlowDirection12221,
	OrificeFlowProperty12222,
	EchocardiographyStrokeVolumeOrigin12223,
	UltrasoundImageMode12224,
	EchocardiographyImageView12226,
	EchocardiographyMeasurementMethod12227,
	EchocardiographyVolumeMethod12228,
	EchocardiographyAreaMethod12229,
	GradientMethod12230,
	VolumeFlowMethod12231,
	MyocardiumMassMethod12232,
	CardiacPhase12233,
	RespirationState12234,
	MitralValveAnatomicSite12235,
	EchocardiographyAnatomicSite12236,
	EchocardiographyAnatomicSiteModifier12237,
	WallMotionScoringScheme12238,
	CardiacOutputProperty12239,
	LeftVentricleAreaMeasurement12240,
	TricuspidValveFindingSite12241,
	AorticValveFindingSite12242,
	LeftVentricleFindingSite12243,
	CongenitalFindingSite12244,
	SurfaceProcessingAlgorithmFamily7162,
	StressTestProcedurePhase3207,
	Stage3778,
	SMLSizeDescriptor252,
	MajorCoronaryArtery3016,
	RadioactivityUnit3083,
	RestStressState3102,
	PETCardiologyProtocol3106,
	PETCardiologyRadiopharmaceutical3107,
	NMPETProcedure3108,
	NuclearCardiologyProtocol3110,
	NuclearCardiologyRadiopharmaceutical3111,
	AttenuationCorrection3112,
	PerfusionDefectType3113,
	StudyQuality3114,
	StressImagingQualityIssue3115,
	NMExtracardiacFinding3116,
	AttenuationCorrectionMethod3117,
	LevelOfRisk3118,
	LVFunction3119,
	PerfusionFinding3120,
	PerfusionMorphology3121,
	VentricularEnlargement3122,
	StressTestProcedure3200,
	IndicationsForStressTest3201,
	ChestPain3202,
	ExerciserDevice3203,
	StressAgent3204,
	IndicationsForPharmacologicalStressTest3205,
	NonInvasiveCardiacImagingProcedure3206,
	ExerciseECGSummaryCode3208,
	StressImagingSummaryCode3209,
	SpeedOfResponse3210,
	BPResponse3211,
	TreadmillSpeed3212,
	StressHemodynamicFinding3213,
	PerfusionFindingMethod3215,
	ComparisonFinding3217,
	StressSymptom3220,
	StressTestTerminationReason3221,
	QTcMeasurement3227,
	ECGTimingMeasurement3228,
	ECGAxisMeasurement3229,
	ECGFinding3230,
	STSegmentFinding3231,
	STSegmentLocation3232,
	STSegmentMorphology3233,
	EctopicBeatMorphology3234,
	PerfusionComparisonFinding3235,
	ToleranceComparisonFinding3236,
	WallMotionComparisonFinding3237,
	StressScoringScale3238,
	PerceivedExertionScale3239,
	VentricleIdentification3463,
	ColonOverallAssessment6200,
	ColonFindingOrFeature6201,
	ColonFindingOrFeatureModifier6202,
	ColonNonLesionObjectType6203,
	AnatomicNonColonFinding6204,
	ClockfaceLocationForColon6205,
	RecumbentPatientOrientationForColon6206,
	ColonQuantitativeTemporalDifferenceType6207,
	ColonTypesOfQualityControlStandard6208,
	ColonMorphologyDescriptor6209,
	LocationInIntestinalTract6210,
	ColonCADMaterialDescription6211,
	CalculatedValueForColonFinding6212,
	OphthalmicHorizontalDirection4214,
	OphthalmicVerticalDirection4215,
	OphthalmicVisualAcuityType4216,
	ArterialPulseWaveform3004,
	RespirationWaveform3005,
	UltrasoundContrastBolusAgent12030,
	ProtocolIntervalEvent12031,
	TransducerScanPattern12032,
	UltrasoundTransducerGeometry12033,
	UltrasoundTransducerBeamSteering12034,
	UltrasoundTransducerApplication12035,
	InstanceAvailabilityStatus50,
	ModalityPPSDiscontinuationReason9301,
	MediaImportPPSDiscontinuationReason9302,
	DXAnatomyImagedForAnimal7482,
	CommonAnatomicRegionsForAnimal7483,
	DXViewForAnimal7484,
	InstitutionalDepartmentUnitService7030,
	PurposeOfReferenceToPredecessorReport7009,
	VisualFixationQualityDuringAcquisition4220,
	VisualFixationQualityProblem4221,
	OphthalmicMacularGridProblem4222,
	Organization5002,
	MixedBreed7486,
	BroselowLutenPediatricSizeCategory7040,
	CMDCTECCCalciumScoringPatientSizeCategory7042,
	CardiacUltrasoundReportTitle12245,
	CardiacUltrasoundIndicationForStudy12246,
	PediatricFetalAndCongenitalCardiacSurgicalIntervention12247,
	CardiacUltrasoundSummaryCode12248,
	CardiacUltrasoundFetalSummaryCode12249,
	CardiacUltrasoundCommonLinearMeasurement12250,
	CardiacUltrasoundLinearValveMeasurement12251,
	CardiacUltrasoundCardiacFunction12252,
	CardiacUltrasoundAreaMeasurement12253,
	CardiacUltrasoundHemodynamicMeasurement12254,
	CardiacUltrasoundMyocardiumMeasurement12255,
	CardiacUltrasoundLeftVentricleMeasurement12257,
	CardiacUltrasoundRightVentricleMeasurement12258,
	CardiacUltrasoundVentriclesMeasurement12259,
	CardiacUltrasoundPulmonaryArteryMeasurement12260,
	CardiacUltrasoundPulmonaryVein12261,
	CardiacUltrasoundPulmonaryValveMeasurement12262,
	CardiacUltrasoundVenousReturnPulmonaryMeasurement12263,
	CardiacUltrasoundVenousReturnSystemicMeasurement12264,
	CardiacUltrasoundAtriaAndAtrialSeptumMeasurement12265,
	CardiacUltrasoundMitralValveMeasurement12266,
	CardiacUltrasoundTricuspidValveMeasurement12267,
	CardiacUltrasoundAtrioventricularValveMeasurement12268,
	CardiacUltrasoundInterventricularSeptumMeasurement12269,
	CardiacUltrasoundAorticValveMeasurement12270,
	CardiacUltrasoundOutflowTractMeasurement12271,
	CardiacUltrasoundSemilunarValveAnnulateAndSinusMeasurement12272,
	CardiacUltrasoundAorticSinotubularJunctionMeasurement12273,
	CardiacUltrasoundAortaMeasurement12274,
	CardiacUltrasoundCoronaryArteryMeasurement12275,
	CardiacUltrasoundAortoPulmonaryConnectionMeasurement12276,
	CardiacUltrasoundPericardiumAndPleuraMeasurement12277,
	CardiacUltrasoundFetalGeneralMeasurement12279,
	CardiacUltrasoundTargetSite12280,
	CardiacUltrasoundTargetSiteModifier12281,
	CardiacUltrasoundVenousReturnSystemicFindingSite12282,
	CardiacUltrasoundVenousReturnPulmonaryFindingSite12283,
	CardiacUltrasoundAtriaAndAtrialSeptumFindingSite12284,
	CardiacUltrasoundAtrioventricularValveFindingSite12285,
	CardiacUltrasoundInterventricularSeptumFindingSite12286,
	CardiacUltrasoundVentricleFindingSite12287,
	CardiacUltrasoundOutflowTractFindingSite12288,
	CardiacUltrasoundSemilunarValveAnnulusAndSinusFindingSite12289,
	CardiacUltrasoundPulmonaryArteryFindingSite12290,
	CardiacUltrasoundAortaFindingSite12291,
	CardiacUltrasoundCoronaryArteryFindingSite12292,
	CardiacUltrasoundAortopulmonaryConnectionFindingSite12293,
	CardiacUltrasoundPericardiumAndPleuraFindingSite12294,
	OphthalmicUltrasoundAxialMeasurementsType4230,
	LensStatus4231,
	VitreousStatus4232,
	OphthalmicAxialLengthMeasurementsSegmentName4233,
	RefractiveSurgeryType4234,
	KeratometryDescriptor4235,
	IOLCalculationFormula4236,
	LensConstantType4237,
	RefractiveErrorType4238,
	AnteriorChamberDepthDefinition4239,
	OphthalmicMeasurementOrCalculationDataSource4240,
	OphthalmicAxialLengthSelectionMethod4241,
	OphthalmicQualityMetricType4243,
	OphthalmicAgentConcentrationUnit4244,
	FunctionalConditionPresentDuringAcquisition91,
	JointPositionDuringAcquisition92,
	JointPositioningMethod93,
	PhysicalForceAppliedDuringAcquisition94,
	ECGControlNumericVariable3690,
	ECGControlTextVariable3691,
	WholeSlideMicroscopyImageReferencedImagePurposeOfReference8120,
	MicroscopyLensType8121,
	MicroscopyIlluminatorAndSensorColor8122,
	MicroscopyIlluminationMethod8123,
	MicroscopyFilter8124,
	MicroscopyIlluminatorType8125,
	AuditEventID400,
	AuditEventTypeCode401,
	AuditActiveParticipantRoleIDCode402,
	SecurityAlertTypeCode403,
	AuditParticipantObjectIDTypeCode404,
	MediaTypeCode405,
	VisualFieldStaticPerimetryTestPattern4250,
	VisualFieldStaticPerimetryTestStrategy4251,
	VisualFieldStaticPerimetryScreeningTestMode4252,
	VisualFieldStaticPerimetryFixationStrategy4253,
	VisualFieldStaticPerimetryTestAnalysisResult4254,
	VisualFieldIlluminationColor4255,
	VisualFieldProcedureModifier4256,
	VisualFieldGlobalIndexName4257,
	AbstractMultiDimensionalImageModelComponentSemantic7180,
	AbstractMultiDimensionalImageModelComponentUnit7181,
	AbstractMultiDimensionalImageModelDimensionSemantic7182,
	AbstractMultiDimensionalImageModelDimensionUnit7183,
	AbstractMultiDimensionalImageModelAxisDirection7184,
	AbstractMultiDimensionalImageModelAxisOrientation7185,
	AbstractMultiDimensionalImageModelQualitativeDimensionSampleSemantic7186,
	PlanningMethod7320,
	DeIdentificationMethod7050,
	MeasurementOrientation12118,
	ECGGlobalWaveformDuration3689,
	ICD3692,
	RadiotherapyGeneralWorkitemDefinition9241,
	RadiotherapyAcquisitionWorkitemDefinition9242,
	RadiotherapyRegistrationWorkitemDefinition9243,
	ContrastBolusSubstance3850,
	LabelType10022,
	OphthalmicMappingUnitForRealWorldValueMapping4260,
	OphthalmicMappingAcquisitionMethod4261,
	RetinalThicknessDefinition4262,
	OphthalmicThicknessMapValueType4263,
	OphthalmicMapPurposeOfReference4264,
	OphthalmicThicknessDeviationCategory4265,
	OphthalmicAnatomicStructureReferencePoint4266,
	CardiacSynchronizationTechnique3104,
	StainingProtocol8130,
	SizeSpecificDoseEstimationMethodForCT10023,
	PathologyImagingProtocol8131,
	MagnificationSelection8132,
	TissueSelection8133,
	GeneralRegionOfInterestMeasurementModifier7464,
	MeasurementDerivedFromMultipleROIMeasurements7465,
	SurfaceScanAcquisitionType8201,
	SurfaceScanModeType8202,
	SurfaceScanRegistrationMethodType8203,
	BasicCardiacView27,
	CTReconstructionAlgorithm10033,
	DetectorType10030,
	CRDRMechanicalConfiguration10031,
	ProjectionXRayAcquisitionDeviceType10032,
	AbstractSegmentationType7165,
	CommonTissueSegmentationType7166,
	PeripheralNervousSystemSegmentationType7167,
	CornealTopographyMappingUnitForRealWorldValueMapping4267,
	CornealTopographyMapValueType4268,
	BrainStructureForVolumetricMeasurement7140,
	RTDoseDerivation7220,
	RTDosePurposeOfReference7221,
	SpectroscopyPurposeOfReference7215,
	ScheduledProcessingParameterConceptCodesForRTTreatment9250,
	RadiopharmaceuticalOrganDoseReferenceAuthority10040,
	SourceOfRadioisotopeActivityInformation10041,
	IntravenousExtravasationSymptom10043,
	RadiosensitiveOrgan10044,
	RadiopharmaceuticalPatientState10045,
	GFRMeasurement10046,
	GFRMeasurementMethod10047,
	VisualEvaluationMethod8300,
	TestPatternCode8301,
	MeasurementPatternCode8302,
	DisplayDeviceType8303,
	SUVUnit85,
	T1MeasurementMethod4100,
	TracerKineticModel4101,
	PerfusionMeasurementMethod4102,
	ArterialInputFunctionMeasurementMethod4103,
	BolusArrivalTimeDerivationMethod4104,
	PerfusionAnalysisMethod4105,
	QuantitativeMethodUsedForPerfusionAndTracerKineticModel4106,
	TracerKineticModelParameter4107,
	PerfusionModelParameter4108,
	ModelIndependentDynamicContrastAnalysisParameter4109,
	TracerKineticModelingCovariate4110,
	ContrastCharacteristic4111,
	MeasurementReportDocumentTitle7021,
	QuantitativeDiagnosticImagingProcedure100,
	PETRegionOfInterestMeasurement7466,
	GrayLevelCoOccurrenceMatrixMeasurement7467,
	TextureMeasurement7468,
	TimePointType6146,
	GenericIntensityAndSizeMeasurement7469,
	ResponseCriteria6147,
	FetalBiometryAnatomicSite12020,
	FetalLongBoneAnatomicSite12021,
	FetalCraniumAnatomicSite12022,
	PelvisAndUterusAnatomicSite12023,
	ParametricMapDerivationImagePurposeOfReference7222,
	PhysicalQuantityDescriptor9000,
	LymphNodeAnatomicSite7600,
	HeadAndNeckCancerAnatomicSite7601,
	FiberTractInBrainstem7701,
	ProjectionAndThalamicFiber7702,
	AssociationFiber7703,
	LimbicSystemTract7704,
	CommissuralFiber7705,
	CranialNerve7706,
	SpinalCordFiber7707,
	TractographyAnatomicSite7710,
	PrimaryAnatomicStructureForIntraOralRadiographySupernumeraryDentitionDesignationOfTeeth4025,
	PrimaryAnatomicStructureForIntraOralAndCraniofacialRadiographyTeeth4026,
	IEC61217DevicePositionParameter9401,
	IEC61217GantryPositionParameter9402,
	IEC61217PatientSupportPositionParameter9403,
	ActionableFindingClassification7035,
	ImageQualityAssessment7036,
	SummaryRadiationExposureQuantity10050,
	WideFieldOphthalmicPhotographyTransformationMethod4245,
	PETUnit84,
	ImplantMaterial7300,
	InterventionType7301,
	ImplantTemplateViewOrientation7302,
	ImplantTemplateModifiedViewOrientation7303,
	ImplantTargetAnatomy7304,
	ImplantPlanningLandmark7305,
	HumanHipImplantPlanningLandmark7306,
	ImplantComponentType7307,
	HumanHipImplantComponentType7308,
	HumanTraumaImplantComponentType7309,
	ImplantFixationMethod7310,
	DeviceParticipatingRole7445,
	ContainerType8101,
	ContainerComponentType8102,
	AnatomicPathologySpecimenType8103,
	BreastTissueSpecimenType8104,
	SpecimenCollectionProcedure8109,
	SpecimenSamplingProcedure8110,
	SpecimenPreparationProcedure8111,
	SpecimenStain8112,
	SpecimenPreparationStep8113,
	SpecimenFixative8114,
	SpecimenEmbeddingMedia8115,
	SourceOfProjectionXRayDoseInformation10020,
	SourceOfCTDoseInformation10021,
	RadiationDoseReferencePoint10025,
	VolumetricViewDescription501,
	VolumetricViewModifier502,
	DiffusionAcquisitionValueType7260,
	DiffusionModelValueType7261,
	DiffusionTractographyAlgorithmFamily7262,
	DiffusionTractographyMeasurementType7263,
	ResearchAnimalSourceRegistry7490,
	YesNoOnly231,
	BiosafetyLevel601,
	BiosafetyControlReason602,
	SexMaleFemaleOrBoth7457,
	AnimalRoomType603,
	DeviceReuse604,
	AnimalBeddingMaterial605,
	AnimalShelterType606,
	AnimalFeedType607,
	AnimalFeedSource608,
	AnimalFeedingMethod609,
	WaterType610,
	AnesthesiaCategoryCodeTypeForSmallAnimalAnesthesia611,
	AnesthesiaCategoryCodeTypeFromAnesthesiaQualityInitiative612,
	AnesthesiaInductionCodeTypeForSmallAnimalAnesthesia613,
	AnesthesiaInductionCodeTypeFromAnesthesiaQualityInitiative614,
	AnesthesiaMaintenanceCodeTypeForSmallAnimalAnesthesia615,
	AnesthesiaMaintenanceCodeTypeFromAnesthesiaQualityInitiative616,
	AirwayManagementMethodCodeTypeForSmallAnimalAnesthesia617,
	AirwayManagementMethodCodeTypeFromAnesthesiaQualityInitiative618,
	AirwayManagementSubMethodCodeTypeForSmallAnimalAnesthesia619,
	AirwayManagementSubMethodCodeTypeFromAnesthesiaQualityInitiative620,
	MedicationTypeForSmallAnimalAnesthesia621,
	MedicationTypeCodeTypeFromAnesthesiaQualityInitiative622,
	MedicationForSmallAnimalAnesthesia623,
	InhalationalAnesthesiaAgentForSmallAnimalAnesthesia624,
	InjectableAnesthesiaAgentForSmallAnimalAnesthesia625,
	PremedicationAgentForSmallAnimalAnesthesia626,
	NeuromuscularBlockingAgentForSmallAnimalAnesthesia627,
	AncillaryMedicationsForSmallAnimalAnesthesia628,
	CarrierGasesForSmallAnimalAnesthesia629,
	LocalAnestheticsForSmallAnimalAnesthesia630,
	ProcedurePhaseRequiringAnesthesia631,
	SurgicalProcedurePhaseRequiringAnesthesia632,
	PhaseOfImagingProcedureRequiringAnesthesia633RETIRED,
	AnimalHandlingPhase634,
	HeatingMethod635,
	TemperatureSensorDeviceComponentTypeForSmallAnimalProcedure636,
	ExogenousSubstanceType637,
	ExogenousSubstance638,
	TumorGraftHistologicType639,
	Fibril640,
	Virus641,
	Cytokine642,
	Toxin643,
	ExogenousSubstanceAdministrationSite644,
	ExogenousSubstanceOriginTissue645,
	PreclinicalSmallAnimalImagingProcedure646,
	PositionReferenceIndicatorForFrameOfReference647,
	PresentAbsentOnly241,
	WaterEquivalentDiameterMethod10024,
	RadiotherapyPurposeOfReference7022,
	ContentAssessmentType701,
	RTContentAssessmentType702,
	AssessmentBasis703,
	ReaderSpecialty7449,
	RequestedReportType9233,
	CTTransversePlaneReferenceBasis1000,
	AnatomicalReferenceBasis1001,
	AnatomicalReferenceBasisHead1002,
	AnatomicalReferenceBasisSpine1003,
	AnatomicalReferenceBasisChest1004,
	AnatomicalReferenceBasisAbdomenPelvis1005,
	AnatomicalReferenceBasisExtremity1006,
	ReferenceGeometryPlane1010,
	ReferenceGeometryPoint1011,
	PatientAlignmentMethod1015,
	ContraindicationsForCTImaging1200,
	FiducialCategory7110,
	Fiducial7111,
	NonImageSourceInstancePurposeOfReference7013,
	RTProcessOutput7023,
	RTProcessInput7024,
	RTProcessInputUsed7025,
	ProstateAnatomy6300,
	ProstateSectorAnatomyFromPIRADSV26301,
	ProstateSectorAnatomyFromEuropeanConcensus16SectorMinimalModel6302,
	ProstateSectorAnatomyFromEuropeanConcensus27SectorOptimalModel6303,
	MeasurementSelectionReason12301,
	EchoFindingObservationType12302,
	EchoMeasurementType12303,
	CardiovascularMeasuredProperty12304,
	BasicEchoAnatomicSite12305,
	EchoFlowDirection12306,
	CardiacPhaseAndTimePoint12307,
	CoreEchoMeasurement12300,
	OCTAProcessingAlgorithmFamily4270,
	EnFaceImageType4271,
	OPTScanPatternType4272,
	RetinalSegmentationSurface4273,
	OrganForRadiationDoseEstimate10060,
	AbsorbedRadiationDoseType10061,
	EquivalentRadiationDoseType10062,
	RadiationDoseEstimateDistributionRepresentation10063,
	PatientModelType10064,
	RadiationTransportModelType10065,
	AttenuatorCategory10066,
	RadiationAttenuatorMaterial10067,
	EstimateMethodType10068,
	RadiationDoseEstimateParameter10069,
	RadiationDoseType10070,
	MRDiffusionComponentSemantic7270,
	MRDiffusionAnisotropyIndex7271,
	MRDiffusionModelParameter7272,
	MRDiffusionModel7273,
	MRDiffusionModelFittingMethod7274,
	MRDiffusionModelSpecificMethod7275,
	MRDiffusionModelInput7276,
	DiffusionRateAreaOverTimeUnit7277,
	PediatricSizeCategory7039,
	CalciumScoringPatientSizeCategory7041,
	ReasonForRepeatingAcquisition10034,
	ProtocolAssertion800,
	RadiotherapeuticDoseMeasurementDevice7026,
	ExportAdditionalInformationDocumentTitle7014,
	ExportDelayReason7015,
	LevelOfDifficulty7016,
	CategoryOfTeachingMaterialImaging7017,
	MiscellaneousDocumentTitle7018,
	SegmentationNonImageSourcePurposeOfReference7019,
	LongitudinalTemporalEventType280,
	NonLesionObjectTypePhysicalObject6401,
	NonLesionObjectTypeSubstance6402,
	NonLesionObjectTypeTissue6403,
	ChestNonLesionObjectTypePhysicalObject6404,
	ChestNonLesionObjectTypeTissue6405,
	TissueSegmentationPropertyType7191,
	AnatomicalStructureSegmentationPropertyType7192,
	PhysicalObjectSegmentationPropertyType7193,
	MorphologicallyAbnormalStructureSegmentationPropertyType7194,
	FunctionSegmentationPropertyType7195,
	SpatialAndRelationalConceptSegmentationPropertyType7196,
	BodySubstanceSegmentationPropertyType7197,
	SubstanceSegmentationPropertyType7198,
	InterpretationRequestDiscontinuationReason9303,
	GrayLevelRunLengthBasedFeature7475,
	GrayLevelSizeZoneBasedFeature7476,
	EncapsulatedDocumentSourcePurposeOfReference7060,
	ModelDocumentTitle7061,
	PurposeOfReferenceToPredecessor3DModel7062,
	ModelScaleUnit7063,
	ModelUsage7064,
	RadiationDoseUnit10071,
	RadiotherapyFiducial7112,
	MultiEnergyRelevantMaterial300,
	MultiEnergyMaterialUnit301,
	DosimetricObjectiveType9500,
	PrescriptionAnatomyCategory9501,
	RTSegmentAnnotationCategory9502,
	RadiotherapyTherapeuticRoleCategory9503,
	RTGeometricInformation9504,
	FixationOrPositioningDevice9505,
	BrachytherapyDevice9506,
	ExternalBodyModel9507,
	NonSpecificVolume9508,
	PurposeOfReferenceForRTPhysicianIntentInput9509,
	PurposeOfReferenceForRTTreatmentPlanningInput9510,
	GeneralExternalRadiotherapyProcedureTechnique9511,
	TomotherapeuticRadiotherapyProcedureTechnique9512,
	FixationDevice9513,
	AnatomicalStructureForRadiotherapy9514,
	RTPatientSupportDevice9515,
	RadiotherapyBolusDeviceType9516,
	RadiotherapyBlockDeviceType9517,
	RadiotherapyAccessoryNoSlotHolderDeviceType9518,
	RadiotherapyAccessorySlotHolderDeviceType9519,
	SegmentedRTAccessoryDevice9520,
	RadiotherapyTreatmentEnergyUnit9521,
	MultiSourceRadiotherapyProcedureTechnique9522,
	RoboticRadiotherapyProcedureTechnique9523,
	RadiotherapyProcedureTechnique9524,
	RadiationTherapyParticle9525,
	IonTherapyParticle9526,
	TeletherapyIsotope9527,
	BrachytherapyIsotope9528,
	SingleDoseDosimetricObjective9529,
	PercentageAndDoseDosimetricObjective9530,
	VolumeAndDoseDosimetricObjective9531,
	NoParameterDosimetricObjective9532,
	DeliveryTimeStructure9533,
	RadiotherapyTarget9534,
	RadiotherapyDoseCalculationRole9535,
	RadiotherapyPrescribingAndSegmentingPersonRole9536,
	EffectiveDoseCalculationMethodCategory9537,
	RadiationTransportBasedEffectiveDoseMethodModifier9538,
	FractionationBasedEffectiveDoseMethodModifier9539,
	ImagingAgentAdministrationAdverseEvent60,
	TimeRelativeToProcedure61RETIRED,
	ImagingAgentAdministrationPhaseType62,
	ImagingAgentAdministrationMode63,
	ImagingAgentAdministrationPatientState64,
	ImagingAgentAdministrationPremedication65,
	ImagingAgentAdministrationMedication66,
	ImagingAgentAdministrationCompletionStatus67,
	ImagingAgentAdministrationPharmaceuticalPresentationUnit68,
	ImagingAgentAdministrationConsumable69,
	Flush70,
	ImagingAgentAdministrationInjectorEventType71,
	ImagingAgentAdministrationStepType72,
	BolusShapingCurve73,
	ImagingAgentAdministrationConsumableCatheterType74,
	LowHighOrEqual75,
	PremedicationType76,
	LateralityWithMedian245,
	DermatologyAnatomicSite4029,
	QuantitativeImageFeature218,
	GlobalShapeDescriptor7477,
	IntensityHistogramFeature7478,
	GreyLevelDistanceZoneBasedFeature7479,
	NeighbourhoodGreyToneDifferenceBasedFeature7500,
	NeighbouringGreyLevelDependenceBasedFeature7501,
	CorneaMeasurementMethodDescriptor4242,
	SegmentedRadiotherapeuticDoseMeasurementDevice7027,
	ClinicalCourseOfDisease6098,
	RacialGroup6099,
	RelativeLaterality246,
	BrainLesionSegmentationTypeWithNecrosis7168,
	BrainLesionSegmentationTypeWithoutNecrosis7169,
	NonAcquisitionModality32,
	Modality33,
	LateralityLeftRightOnly247,
	QualitativeEvaluationModifierType210,
	QualitativeEvaluationModifierValue211,
	GenericAnatomicLocationModifier212,
	BeamLimitingDeviceType9541,
	CompensatorDeviceType9542,
	RadiotherapyTreatmentMachineMode9543,
	RadiotherapyDistanceReferenceLocation9544,
	FixedBeamLimitingDeviceType9545,
	RadiotherapyWedgeType9546,
	RTBeamLimitingDeviceOrientationLabel9547,
	GeneralAccessoryDeviceType9548,
	RadiationGenerationModeType9549,
	CArmPhotonElectronDeliveryRateUnit9550,
	TreatmentDeliveryDeviceType9551,
	CArmPhotonElectronDosimeterUnit9552,
	TreatmentPoint9553,
	EquipmentReferencePoint9554,
	RadiotherapyTreatmentPlanningPersonRole9555,
	RealTimeVideoRenditionTitle7070,
	GeometryGraphicalRepresentation219,
	VisualExplanation217,
	ProstateSectorAnatomyFromPIRADSV216304,
	RadiotherapyRoboticNodeSet9556,
	TomotherapeuticDosimeterUnit9557,
	TomotherapeuticDoseRateUnit9558,
	RoboticDeliveryDeviceDosimeterUnit9559,
	RoboticDeliveryDeviceDoseRateUnit9560,
	AnatomicStructure8134,
	MediastinumFindingOrFeature6148,
	MediastinumAnatomy6149,
	VascularUltrasoundReportDocumentTitle12100,
	OrganPartNonLateralized12130,
	OrganPartLateralized12131,
	TreatmentTerminationReason9561,
	RadiotherapyTreatmentDeliveryPersonRole9562,
	RadiotherapyInterlockResolution9563,
	TreatmentSessionConfirmationAssertion9564,
	TreatmentToleranceViolationCause9565,
	ClinicalToleranceViolationType9566,
	MachineToleranceViolationType9567,
	RadiotherapyTreatmentInterlock9568,
	IsocentricPatientSupportPositionParameter9569,
	RTOverriddenTreatmentParameter9570,
	EEGLead3030,
	LeadLocationNearOrInMuscle3031,
	LeadLocationNearPeripheralNerve3032,
	EOGLead3033,
	BodyPositionChannel3034,
	EEGAnnotationNeurophysiologicEnumeration3035,
	EMGAnnotationNeurophysiologicalEnumeration3036,
	EOGAnnotationNeurophysiologicalEnumeration3037,
	PatternEvent3038,
	DeviceRelatedAndEnvironmentRelatedEvent3039,
	EEGAnnotationNeurologicalMonitoringMeasurement3040,
	OBGYNUltrasoundReportDocumentTitle12024,
	AutomationOfMeasurement7230,
	OBGYNUltrasoundBeamPath12025,
	AngleMeasurement7550,
	GenericPurposeOfReferenceToImagesAndCoordinatesInMeasurement7551,
	GenericPurposeOfReferenceToImagesInMeasurement7552,
	GenericPurposeOfReferenceToCoordinatesInMeasurement7553,
	FitzpatrickSkinType4401,
	HistoryOfMalignantMelanoma4402,
	HistoryOfMelanomaInSitu4403,
	HistoryOfNonMelanomaSkinCancer4404,
	SkinDisorder4405,
	PatientReportedLesionCharacteristic4406,
	LesionPalpationFinding4407,
	LesionVisualFinding4408,
	SkinProcedure4409,
	AbdominopelvicVessel12125,
	NumericValueFailureQualifier43,
	NumericValueUnknownQualifier44,
	CouinaudLiverSegment7170,
	LiverSegmentationType7171,
	ContraindicationsForXAImaging1201,
	NeurophysiologicStimulationMode3041,
	ReportedValueType10072,
	ValueTiming10073,
	RDSRFrameOfReferenceOrigin10074,
	MicroscopyAnnotationPropertyType8135,
	MicroscopyMeasurementType8136,
	ProstateReportingSystem6310,
	MRSignalIntensity6311,
	CrossSectionalScanPlaneOrientation6312,
	HistoryOfProstateDisease6313,
	ProstateMRIStudyQualityFinding6314,
	ProstateMRISeriesQualityFinding6315,
	MRImagingArtifact6316,
	ProstateDCEMRIQualityFinding6317,
	ProstateDWIMRIQualityFinding6318,
	AbdominalInterventionType6319,
	AbdominopelvicIntervention6320,
	ProstateCancerDiagnosticProcedure6321,
	ProstateCancerFamilyHistory6322,
	ProstateCancerTherapy6323,
	ProstateMRIAssessment6324,
	OverallAssessmentFromPIRADS6325,
	ImageQualityControlStandard6326,
	ProstateImagingIndication6327,
	PIRADSV2LesionAssessmentCategory6328,
	PIRADSV2T2WIPZLesionAssessmentCategory6329,
	PIRADSV2T2WITZLesionAssessmentCategory6330,
	PIRADSV2DWILesionAssessmentCategory6331,
	PIRADSV2DCELesionAssessmentCategory6332,
	mpMRIAssessmentType6333,
	mpMRIAssessmentTypeFromPIRADS6334,
	mpMRIAssessmentValue6335,
	MRIAbnormality6336,
	mpMRIProstateAbnormalityFromPIRADS6337,
	mpMRIBenignProstateAbnormalityFromPIRADS6338,
	MRIShapeCharacteristic6339,
	ProstateMRIShapeCharacteristicFromPIRADS6340,
	MRIMarginCharacteristic6341,
	ProstateMRIMarginCharacteristicFromPIRADS6342,
	MRISignalCharacteristic6343,
	ProstateMRISignalCharacteristicFromPIRADS6344,
	MRIEnhancementPattern6345,
	ProstateMRIEnhancementPatternFromPIRADS6346,
	ProstateMRIExtraProstaticFinding6347,
	ProstateMRIAssessmentOfExtraProstaticAnatomicSite6348,
	MRCoilType6349,
	EndorectalCoilFillSubstance6350,
	ProstateRelationalMeasurement6351,
	ProstateCancerDiagnosticBloodLabMeasurement6352,
	ProstateImagingTypesOfQualityControlStandard6353,
	UltrasoundShearWaveMeasurement12308,
	LeftVentricleMyocardialWall16SegmentModel3780RETIRED,
	LeftVentricleMyocardialWall18SegmentModel3781,
	LeftVentricleBasalWall6Segments3782,
	LeftVentricleMidlevelWall6Segments3783,
	LeftVentricleApicalWall4Segments3784,
	LeftVentricleApicalWall6Segments3785,
	PatientTreatmentPreparationMethod9571,
	PatientShieldingDevice9572,
	PatientTreatmentPreparationDevice9573,
	PatientPositionDisplacementReferencePoint9574,
	PatientAlignmentDevice9575,
	ReasonsForRTRadiationTreatmentOmission9576,
	PatientTreatmentPreparationProcedure9577,
	MotionManagementSetupDevice9578,
	CoreEchoStrainMeasurement12309,
	MyocardialStrainMethod12310,
	EchoMeasuredStrainProperty12311,
	AssessmentFromCADRADS3020,
	CADRADSStenosisAssessmentModifier3021,
	CADRADSAssessmentModifier3022,
	RTSegmentMaterial9579,
	VertebralAnatomicStructure7602,
	Vertebra7603,
	IntervertebralDisc7604,
	ImagingProcedure101,
	NICIPShortCodeImagingProcedure103,
	NICIPSNOMEDImagingProcedure104,
	ICD10PCSImagingProcedure105,
	ICD10PCSNuclearMedicineProcedure106,
	ICD10PCSRadiationTherapyProcedure107,
	RTSegmentationPropertyCategory9580,
	RadiotherapyRegistrationMark9581,
	RadiotherapyDoseRegion9582,
	AnatomicallyLocalizedLesionSegmentationType7199,
	ReasonForRemovalFromOperationalUse7031,
	GeneralUltrasoundReportDocumentTitle12320,
	ElastographySite12321,
	ElastographyMeasurementSite12322,
	UltrasoundRelevantPatientCondition12323,
	ShearWaveDetectionMethod12324,
	LiverUltrasoundStudyIndication12325,
	AnalogWaveformFilter3042,
	DigitalWaveformFilter3043,
	WaveformFilterLookupTableInputFrequencyUnit3044,
	WaveformFilterLookupTableOutputMagnitudeUnit3045,
	SpecificObservationSubjectClass272,
	MovableBeamLimitingDeviceType9540,
	RadiotherapyAcquisitionWorkItemSubtasks9260,
	PatientPositionAcquisitionRadiationSourceLocations9261,
	EnergyDerivationTypes9262,
	KVImagingAcquisitionTechniques9263,
	MVImagingAcquisitionTechniques9264,
	PatientPositionAcquisitionProjectionTechniques9265,
	PatientPositionAcquisitionCTTechniques9266,
	PatientPositioningRelatedObjectPurposes9267,
	PatientPositionAcquisitionDevices9268,
	RTRadiationMetersetUnits9269,
	AcquisitionInitiationTypes9270,
	RTImagePatientPositionAcquisitionDevices9271,
	PhotoacousticIlluminationMethod11001,
	AcousticCouplingMedium11002,
	UltrasoundTransducerTechnology11003,
	SpeedOfSoundCorrectionMechanisms11004,
	PhotoacousticReconstructionAlgorithmFamily11005,
	PhotoacousticImagedProperty11006,
	XRayRadiationDoseProcedureTypeReported10005,
	TopicalTreatment4410,
	LesionColor4411,
	SpecimenStainForConfocalMicroscopy4412,
	RTROIImageAcquisitionContext9272,
	LobeOfLung6170,
	ZoneOfLung6171,
	SleepStage3046,
	PatientPositionAcquisitionMRTechniques9273,
	RTPlanRadiotherapyProcedureTechnique9583,
	WaveformAnnotationClassification3047,
	WaveformAnnotationsDocumentTitle3048,
	EEGProcedure3049,
	PatientConsciousness3050,
	FollicleType12010,
	BreastTissueSegmentationType7163,
	ImplantedDevice3779,
	SimilarityMeasure281,
	WaveformAcquisitionModality34,
	EnFaceProcessingAlgorithmFamily4274,
	AnteriorEyeSegmentationSurface4275,
	FetalEchocardiographyImageView12312,
	CardiacUltrasoundFetalArrhythmiaMeasurements12313,
	CommonFetalEchocardiographyMeasurements12314,
	HeadAndNeckPrimaryAnatomicStructure4061,
	VLView4062,
	VLDentalView4063,
	VLViewModifier4064,
	VLDentalViewModifier4065,
	OrthognathicFunctionalCondition4066,
	OrthodonticFindingByInspection4067,
	OrthodonticObservableEntity4068,
	DentalOcclusion4069,
	OrthodonticTreatmentProgress4070,
	GeneralPhotographyDevice4071,
	DevicesForThePurposeOfDentalPhotography4072,
	CTDIPhantomDevice4053,
	DiagnosticImagingProcedureWithoutIVContrast108,
	DiagnosticImagingProcedureWithIVContrast109,
	StructuralHeartProcedure12331,
	StructuralHeartDevice12332,
	StructuralHeartMeasurement12333,
	AorticValveStructuralMeasurement12334,
	MitralValveStructuralMeasurement12335,
	TricuspidValveStructuralMeasurement12336,
	StructuralHeartEchoMeasurement12337,
	LeftAtrialAppendageClosureMeasurement12338,
	StructuralHeartProcedureAnatomicSite12339,
	IndicationForStructuralHeartProcedure12341,
	BradycardiacAgent12342,
	TransesophagealEchocardiographyScanPlane12343,
	StructuralHeartMeasurementReportDocumentTitle12344,
	PersonGenderIdentity7458,
	CategoryOfSexParametersForClinicalUse7459,
	ThirdPersonPronounSet7448,
	CardiacStructureCalcificationQualitativeEvaluation12345,
	VisualFieldMeasurements4280,
	OpticDiscKeyMeasurements4281,
	RetinalSectorMethods4282,
	RNFLSectorMeasurements4283,
	RNFLClockfaceMeasurements4284,
	MacularThicknessKeyMeasurements4285,
	GanglionCellMeasurementExtent4286,
	GanglionCellKeyMeasurements4287,
	GanglionCellSectorMeasurements4288,
	GanglionCellSectorMethods4289,
	EndothelialCellCountMeasurements4290,
	OphthalmicImageROIMeasurements4291,
	RTPlanApprovalAssertion9584,
	EstimatedDeliveryDateMethod12026,
	RTDoseCalculationAlgorithmFamily9585,
	DoseIndexForDoseCalibration10012,
	UltrasoundAttenuationImagingSite12036,
	FetalAnatomySurveyAssessment12040,
	FetalAnatomySurveyAssessmentHead12041,
	FetalAnatomySurveyAssessmentFaceAndNeck12042,
	FetalAnatomySurveyAssessmentChest12043,
	FetalAnatomySurveyAssessmentHeart12044,
	FetalAnatomySurveyAssessmentAbdomenAndPelvis12045,
	FetalAnatomySurveyAssessmentSpine12046,
	FetalAnatomySurveyAssessmentExtremities12047,
	FetalAnatomySurveyAssessmentMaternal12048,
	FetalAnatomySurveyPracticeGuideline12049,
	SensitiveContentCategory900,
	SensitiveContentDetail901,
	ApplicationTypeCode406,
	XRayModulationType10035,
	RadiotherapyDoseRealWorldUnits9586,
	RadiotherapyDoseInterpretedTypeCodes9587,
	RadiotherapyDoseInterpretedTypeModifierCodes9588,
	RadiotherapyDoseIntentCodes9589,
	QualitySegmentationPropertyType7164,
	UltrasoundZScorePopulationIndex12027,
	FetalUltrasoundZScoreReferenceAuthority12028,
	MetalArtifactReductionAlgorithmFamily10036,
}

var generatedStandardUIDIndex = map[string]*UID{
	"1.2.840.10008.1.1":                Verification,
	"1.2.840.10008.1.2":                ImplicitVRLittleEndian,
	"1.2.840.10008.1.2.1":              ExplicitVRLittleEndian,
	"1.2.840.10008.1.2.1.98":           EncapsulatedUncompressedExplicitVRLittleEndian,
	"1.2.840.10008.1.2.1.99":           DeflatedExplicitVRLittleEndian,
	"1.2.840.10008.1.2.2":              ExplicitVRBigEndianRETIRED,
	"1.2.840.10008.1.2.4.50":           JPEGBaseline8Bit,
	"1.2.840.10008.1.2.4.51":           JPEGExtended12Bit,
	"1.2.840.10008.1.2.4.52":           JPEGExtended35RETIRED,
	"1.2.840.10008.1.2.4.53":           JPEGSpectralSelectionNonHierarchical68RETIRED,
	"1.2.840.10008.1.2.4.54":           JPEGSpectralSelectionNonHierarchical79RETIRED,
	"1.2.840.10008.1.2.4.55":           JPEGFullProgressionNonHierarchical1012RETIRED,
	"1.2.840.10008.1.2.4.56":           JPEGFullProgressionNonHierarchical1113RETIRED,
	"1.2.840.10008.1.2.4.57":           JPEGLossless,
	"1.2.840.10008.1.2.4.58":           JPEGLosslessNonHierarchical15RETIRED,
	"1.2.840.10008.1.2.4.59":           JPEGExtendedHierarchical1618RETIRED,
	"1.2.840.10008.1.2.4.60":           JPEGExtendedHierarchical1719RETIRED,
	"1.2.840.10008.1.2.4.61":           JPEGSpectralSelectionHierarchical2022RETIRED,
	"1.2.840.10008.1.2.4.62":           JPEGSpectralSelectionHierarchical2123RETIRED,
	"1.2.840.10008.1.2.4.63":           JPEGFullProgressionHierarchical2426RETIRED,
	"1.2.840.10008.1.2.4.64":           JPEGFullProgressionHierarchical2527RETIRED,
	"1.2.840.10008.1.2.4.65":           JPEGLosslessHierarchical28RETIRED,
	"1.2.840.10008.1.2.4.66":           JPEGLosslessHierarchical29RETIRED,
	"1.2.840.10008.1.2.4.70":           JPEGLosslessSV1,
	"1.2.840.10008.1.2.4.80":           JPEGLSLossless,
	"1.2.840.10008.1.2.4.81":           JPEGLSNearLossless,
	"1.2.840.10008.1.2.4.90":           JPEG2000Lossless,
	"1.2.840.10008.1.2.4.91":           JPEG2000,
	"1.2.840.10008.1.2.4.92":           JPEG2000MCLossless,
	"1.2.840.10008.1.2.4.93":           JPEG2000MC,
	"1.2.840.10008.1.2.4.94":           JPIPReferenced,
	"1.2.840.10008.1.2.4.95":           JPIPReferencedDeflate,
	"1.2.840.10008.1.2.4.100":          MPEG2MPML,
	"1.2.840.10008.1.2.4.100.1":        MPEG2MPMLF,
	"1.2.840.10008.1.2.4.101":          MPEG2MPHL,
	"1.2.840.10008.1.2.4.101.1":        MPEG2MPHLF,
	"1.2.840.10008.1.2.4.102":          MPEG4HP41,
	"1.2.840.10008.1.2.4.102.1":        MPEG4HP41F,
	"1.2.840.10008.1.2.4.103":          MPEG4HP41BD,
	"1.2.840.10008.1.2.4.103.1":        MPEG4HP41BDF,
	"1.2.840.10008.1.2.4.104":          MPEG4HP422D,
	"1.2.840.10008.1.2.4.104.1":        MPEG4HP422DF,
	"1.2.840.10008.1.2.4.105":          MPEG4HP423D,
	"1.2.840.10008.1.2.4.105.1":        MPEG4HP423DF,
	"1.2.840.10008.1.2.4.106":          MPEG4HP42STEREO,
	"1.2.840.10008.1.2.4.106.1":        MPEG4HP42STEREOF,
	"1.2.840.10008.1.2.4.107":          HEVCMP51,
	"1.2.840.10008.1.2.4.108":          HEVCM10P51,
	"1.2.840.10008.1.2.4.110":          JPEGXLLossless,
	"1.2.840.10008.1.2.4.111":          JPEGXLJPEGRecompression,
	"1.2.840.10008.1.2.4.112":          JPEGXL,
	"1.2.840.10008.1.2.4.201":          HTJ2KLossless,
	"1.2.840.10008.1.2.4.202":          HTJ2KLosslessRPCL,
	"1.2.840.10008.1.2.4.203":          HTJ2K,
	"1.2.840.10008.1.2.4.204":          JPIPHTJ2KReferenced,
	"1.2.840.10008.1.2.4.205":          JPIPHTJ2KReferencedDeflate,
	"1.2.840.10008.1.2.5":              RLELossless,
	"1.2.840.10008.1.2.6.1":            RFC2557MIMEEncapsulationRETIRED,
	"1.2.840.10008.1.2.6.2":            XMLEncodingRETIRED,
	"1.2.840.10008.1.2.7.1":            SMPTEST211020UncompressedProgressiveActiveVideo,
	"1.2.840.10008.1.2.7.2":            SMPTEST211020UncompressedInterlacedActiveVideo,
	"1.2.840.10008.1.2.7.3":            SMPTEST211030PCMDigitalAudio,
	"1.2.840.10008.1.2.8.1":            DeflatedImageFrameCompression,
	"1.2.840.10008.1.3.10":             MediaStorageDirectoryStorage,
	"1.2.840.10008.1.5.1":              HotIronPalette,
	"1.2.840.10008.1.5.2":              PETPalette,
	"1.2.840.10008.1.5.3":              HotMetalBluePalette,
	"1.2.840.10008.1.5.4":              PET20StepPalette,
	"1.2.840.10008.1.5.5":              SpringPalette,
	"1.2.840.10008.1.5.6":              SummerPalette,
	"1.2.840.10008.1.5.7":              FallPalette,
	"1.2.840.10008.1.5.8":              WinterPalette,
	"1.2.840.10008.1.9":                BasicStudyContentNotificationRETIRED,
	"1.2.840.10008.1.20":               Papyrus3ImplicitVRLittleEndianRETIRED,
	"1.2.840.10008.1.20.1":             StorageCommitmentPushModel,
	"1.2.840.10008.1.20.1.1":           StorageCommitmentPushModelInstance,
	"1.2.840.10008.1.20.2":             StorageCommitmentPullModelRETIRED,
	"1.2.840.10008.1.20.2.1":           StorageCommitmentPullModelInstanceRETIRED,
	"1.2.840.10008.1.40":               ProceduralEventLogging,
	"1.2.840.10008.1.40.1":             ProceduralEventLoggingInstance,
	"1.2.840.10008.1.42":               SubstanceAdministrationLogging,
	"1.2.840.10008.1.42.1":             SubstanceAdministrationLoggingInstance,
	"1.2.840.10008.2.6.1":              DCMUID,
	"1.2.840.10008.2.16.4":             DCM,
	"1.2.840.10008.2.16.5":             MA,
	"1.2.840.10008.2.16.6":             UBERON,
	"1.2.840.10008.2.16.7":             ITIS_TSN,
	"1.2.840.10008.2.16.8":             MGI,
	"1.2.840.10008.2.16.9":             PUBCHEM_CID,
	"1.2.840.10008.2.16.10":            DC,
	"1.2.840.10008.2.16.11":            NYUMCCG,
	"1.2.840.10008.2.16.12":            MAYONRISBSASRG,
	"1.2.840.10008.2.16.13":            IBSI,
	"1.2.840.10008.2.16.14":            RO,
	"1.2.840.10008.2.16.15":            RADELEMENT,
	"1.2.840.10008.2.16.16":            I11,
	"1.2.840.10008.2.16.17":            UNS,
	"1.2.840.10008.2.16.18":            RRID,
	"1.2.840.10008.3.1.1.1":            DICOMApplicationContext,
	"1.2.840.10008.3.1.2.1.1":          DetachedPatientManagementRETIRED,
	"1.2.840.10008.3.1.2.1.4":          DetachedPatientManagementMetaRETIRED,
	"1.2.840.10008.3.1.2.2.1":          DetachedVisitManagementRETIRED,
	"1.2.840.10008.3.1.2.3.1":          DetachedStudyManagementRETIRED,
	"1.2.840.10008.3.1.2.3.2":          StudyComponentManagementRETIRED,
	"1.2.840.10008.3.1.2.3.3":          ModalityPerformedProcedureStep,
	"1.2.840.10008.3.1.2.3.4":          ModalityPerformedProcedureStepRetrieve,
	"1.2.840.10008.3.1.2.3.5":          ModalityPerformedProcedureStepNotification,
	"1.2.840.10008.3.1.2.5.1":          DetachedResultsManagementRETIRED,
	"1.2.840.10008.3.1.2.5.4":          DetachedResultsManagementMetaRETIRED,
	"1.2.840.10008.3.1.2.5.5":          DetachedStudyManagementMetaRETIRED,
	"1.2.840.10008.3.1.2.6.1":          DetachedInterpretationManagementRETIRED,
	"1.2.840.10008.4.2":                Storage,
	"1.2.840.10008.5.1.1.1":            BasicFilmSession,
	"1.2.840.10008.5.1.1.2":            BasicFilmBox,
	"1.2.840.10008.5.1.1.4":            BasicGrayscaleImageBox,
	"1.2.840.10008.5.1.1.4.1":          BasicColorImageBox,
	"1.2.840.10008.5.1.1.4.2":          ReferencedImageBoxRETIRED,
	"1.2.840.10008.5.1.1.9":            BasicGrayscalePrintManagementMeta,
	"1.2.840.10008.5.1.1.9.1":          ReferencedGrayscalePrintManagementMetaRETIRED,
	"1.2.840.10008.5.1.1.14":           PrintJob,
	"1.2.840.10008.5.1.1.15":           BasicAnnotationBox,
	"1.2.840.10008.5.1.1.16":           Printer,
	"1.2.840.10008.5.1.1.16.376":       PrinterConfigurationRetrieval,
	"1.2.840.10008.5.1.1.17":           PrinterInstance,
	"1.2.840.10008.5.1.1.17.376":       PrinterConfigurationRetrievalInstance,
	"1.2.840.10008.5.1.1.18":           BasicColorPrintManagementMeta,
	"1.2.840.10008.5.1.1.18.1":         ReferencedColorPrintManagementMetaRETIRED,
	"1.2.840.10008.5.1.1.22":           VOILUTBox,
	"1.2.840.10008.5.1.1.23":           PresentationLUT,
	"1.2.840.10008.5.1.1.24":           ImageOverlayBoxRETIRED,
	"1.2.840.10008.5.1.1.24.1":         BasicPrintImageOverlayBoxRETIRED,
	"1.2.840.10008.5.1.1.25":           PrintQueueInstanceRETIRED,
	"1.2.840.10008.5.1.1.26":           PrintQueueManagementRETIRED,
	"1.2.840.10008.5.1.1.27":           StoredPrintStorageRETIRED,
	"1.2.840.10008.5.1.1.29":           HardcopyGrayscaleImageStorageRETIRED,
	"1.2.840.10008.5.1.1.30":           HardcopyColorImageStorageRETIRED,
	"1.2.840.10008.5.1.1.31":           PullPrintRequestRETIRED,
	"1.2.840.10008.5.1.1.32":           PullStoredPrintManagementMetaRETIRED,
	"1.2.840.10008.5.1.1.33":           MediaCreationManagement,
	"1.2.840.10008.5.1.1.40":           DisplaySystem,
	"1.2.840.10008.5.1.1.40.1":         DisplaySystemInstance,
	"1.2.840.10008.5.1.4.1.1.1":        ComputedRadiographyImageStorage,
	"1.2.840.10008.5.1.4.1.1.1.1":      DigitalXRayImageStorageForPresentation,
	"1.2.840.10008.5.1.4.1.1.1.1.1":    DigitalXRayImageStorageForProcessing,
	"1.2.840.10008.5.1.4.1.1.1.2":      DigitalMammographyXRayImageStorageForPresentation,
	"1.2.840.10008.5.1.4.1.1.1.2.1":    DigitalMammographyXRayImageStorageForProcessing,
	"1.2.840.10008.5.1.4.1.1.1.3":      DigitalIntraOralXRayImageStorageForPresentation,
	"1.2.840.10008.5.1.4.1.1.1.3.1":    DigitalIntraOralXRayImageStorageForProcessing,
	"1.2.840.10008.5.1.4.1.1.2":        CTImageStorage,
	"1.2.840.10008.5.1.4.1.1.2.1":      EnhancedCTImageStorage,
	"1.2.840.10008.5.1.4.1.1.2.2":      LegacyConvertedEnhancedCTImageStorage,
	"1.2.840.10008.5.1.4.1.1.3":        UltrasoundMultiFrameImageStorageRetiredRETIRED,
	"1.2.840.10008.5.1.4.1.1.3.1":      UltrasoundMultiFrameImageStorage,
	"1.2.840.10008.5.1.4.1.1.4":        MRImageStorage,
	"1.2.840.10008.5.1.4.1.1.4.1":      EnhancedMRImageStorage,
	"1.2.840.10008.5.1.4.1.1.4.2":      MRSpectroscopyStorage,
	"1.2.840.10008.5.1.4.1.1.4.3":      EnhancedMRColorImageStorage,
	"1.2.840.10008.5.1.4.1.1.4.4":      LegacyConvertedEnhancedMRImageStorage,
	"1.2.840.10008.5.1.4.1.1.5":        NuclearMedicineImageStorageRetiredRETIRED,
	"1.2.840.10008.5.1.4.1.1.6":        UltrasoundImageStorageRetiredRETIRED,
	"1.2.840.10008.5.1.4.1.1.6.1":      UltrasoundImageStorage,
	"1.2.840.10008.5.1.4.1.1.6.2":      EnhancedUSVolumeStorage,
	"1.2.840.10008.5.1.4.1.1.6.3":      PhotoacousticImageStorage,
	"1.2.840.10008.5.1.4.1.1.7":        SecondaryCaptureImageStorage,
	"1.2.840.10008.5.1.4.1.1.7.1":      MultiFrameSingleBitSecondaryCaptureImageStorage,
	"1.2.840.10008.5.1.4.1.1.7.2":      MultiFrameGrayscaleByteSecondaryCaptureImageStorage,
	"1.2.840.10008.5.1.4.1.1.7.3":      MultiFrameGrayscaleWordSecondaryCaptureImageStorage,
	"1.2.840.10008.5.1.4.1.1.7.4":      MultiFrameTrueColorSecondaryCaptureImageStorage,
	"1.2.840.10008.5.1.4.1.1.8":        StandaloneOverlayStorageRETIRED,
	"1.2.840.10008.5.1.4.1.1.9":        StandaloneCurveStorageRETIRED,
	"1.2.840.10008.5.1.4.1.1.9.1":      WaveformStorageTrialRETIRED,
	"1.2.840.10008.5.1.4.1.1.9.1.1":    TwelveLeadECGWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.1.2":    GeneralECGWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.1.3":    AmbulatoryECGWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.1.4":    General32bitECGWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.2.1":    HemodynamicWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.3.1":    CardiacElectrophysiologyWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.4.1":    BasicVoiceAudioWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.4.2":    GeneralAudioWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.5.1":    ArterialPulseWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.6.1":    RespiratoryWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.6.2":    MultichannelRespiratoryWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.7.1":    RoutineScalpElectroencephalogramWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.7.2":    ElectromyogramWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.7.3":    ElectrooculogramWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.7.4":    SleepElectroencephalogramWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.8.1":    BodyPositionWaveformStorage,
	"1.2.840.10008.5.1.4.1.1.9.100.1":  WaveformPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.9.100.2":  WaveformAcquisitionPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.10":       StandaloneModalityLUTStorageRETIRED,
	"1.2.840.10008.5.1.4.1.1.11":       StandaloneVOILUTStorageRETIRED,
	"1.2.840.10008.5.1.4.1.1.11.1":     GrayscaleSoftcopyPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.2":     ColorSoftcopyPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.3":     PseudoColorSoftcopyPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.4":     BlendingSoftcopyPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.5":     XAXRFGrayscaleSoftcopyPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.6":     GrayscalePlanarMPRVolumetricPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.7":     CompositingPlanarMPRVolumetricPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.8":     AdvancedBlendingPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.9":     VolumeRenderingVolumetricPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.10":    SegmentedVolumeRenderingVolumetricPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.11":    MultipleVolumeRenderingVolumetricPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.11.12":    VariableModalityLUTSoftcopyPresentationStateStorage,
	"1.2.840.10008.5.1.4.1.1.12.1":     XRayAngiographicImageStorage,
	"1.2.840.10008.5.1.4.1.1.12.1.1":   EnhancedXAImageStorage,
	"1.2.840.10008.5.1.4.1.1.12.2":     XRayRadiofluoroscopicImageStorage,
	"1.2.840.10008.5.1.4.1.1.12.2.1":   EnhancedXRFImageStorage,
	"1.2.840.10008.5.1.4.1.1.12.3":     XRayAngiographicBiPlaneImageStorageRETIRED,
	"1.2.840.10008.5.1.4.1.1.13.1.1":   XRay3DAngiographicImageStorage,
	"1.2.840.10008.5.1.4.1.1.13.1.2":   XRay3DCraniofacialImageStorage,
	"1.2.840.10008.5.1.4.1.1.13.1.3":   BreastTomosynthesisImageStorage,
	"1.2.840.10008.5.1.4.1.1.13.1.4":   BreastProjectionXRayImageStorageForPresentation,
	"1.2.840.10008.5.1.4.1.1.13.1.5":   BreastProjectionXRayImageStorageForProcessing,
	"1.2.840.10008.5.1.4.1.1.14.1":     IntravascularOpticalCoherenceTomographyImageStorageForPresentation,
	"1.2.840.10008.5.1.4.1.1.14.2":     IntravascularOpticalCoherenceTomographyImageStorageForProcessing,
	"1.2.840.10008.5.1.4.1.1.20":       NuclearMedicineImageStorage,
	"1.2.840.10008.5.1.4.1.1.30":       ParametricMapStorage,
	"1.2.840.10008.5.1.4.1.1.66":       RawDataStorage,
	"1.2.840.10008.5.1.4.1.1.66.1":     SpatialRegistrationStorage,
	"1.2.840.10008.5.1.4.1.1.66.2":     SpatialFiducialsStorage,
	"1.2.840.10008.5.1.4.1.1.66.3":     DeformableSpatialRegistrationStorage,
	"1.2.840.10008.5.1.4.1.1.66.4":     SegmentationStorage,
	"1.2.840.10008.5.1.4.1.1.66.5":     SurfaceSegmentationStorage,
	"1.2.840.10008.5.1.4.1.1.66.6":     TractographyResultsStorage,
	"1.2.840.10008.5.1.4.1.1.66.7":     LabelMapSegmentationStorage,
	"1.2.840.10008.5.1.4.1.1.66.8":     HeightMapSegmentationStorage,
	"1.2.840.10008.5.1.4.1.1.67":       RealWorldValueMappingStorage,
	"1.2.840.10008.5.1.4.1.1.68.1":     SurfaceScanMeshStorage,
	"1.2.840.10008.5.1.4.1.1.68.2":     SurfaceScanPointCloudStorage,
	"1.2.840.10008.5.1.4.1.1.77.1":     VLImageStorageTrialRETIRED,
	"1.2.840.10008.5.1.4.1.1.77.2":     VLMultiFrameImageStorageTrialRETIRED,
	"1.2.840.10008.5.1.4.1.1.77.1.1":   VLEndoscopicImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.1.1": VideoEndoscopicImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.2":   VLMicroscopicImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.2.1": VideoMicroscopicImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.3":   VLSlideCoordinatesMicroscopicImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.4":   VLPhotographicImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.4.1": VideoPhotographicImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.5.1": OphthalmicPhotography8BitImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.5.2": OphthalmicPhotography16BitImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.5.3": StereometricRelationshipStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.5.4": OphthalmicTomographyImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.5.5": WideFieldOphthalmicPhotographyStereographicProjectionImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.5.6": WideFieldOphthalmicPhotography3DCoordinatesImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.5.7": OphthalmicOpticalCoherenceTomographyEnFaceImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.5.8": OphthalmicOpticalCoherenceTomographyBscanVolumeAnalysisStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.6":   VLWholeSlideMicroscopyImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.7":   DermoscopicPhotographyImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.8":   ConfocalMicroscopyImageStorage,
	"1.2.840.10008.5.1.4.1.1.77.1.9":   ConfocalMicroscopyTiledPyramidalImageStorage,
	"1.2.840.10008.5.1.4.1.1.78.1":     LensometryMeasurementsStorage,
	"1.2.840.10008.5.1.4.1.1.78.2":     AutorefractionMeasurementsStorage,
	"1.2.840.10008.5.1.4.1.1.78.3":     KeratometryMeasurementsStorage,
	"1.2.840.10008.5.1.4.1.1.78.4":     SubjectiveRefractionMeasurementsStorage,
	"1.2.840.10008.5.1.4.1.1.78.5":     VisualAcuityMeasurementsStorage,
	"1.2.840.10008.5.1.4.1.1.78.6":     SpectaclePrescriptionReportStorage,
	"1.2.840.10008.5.1.4.1.1.78.7":     OphthalmicAxialMeasurementsStorage,
	"1.2.840.10008.5.1.4.1.1.78.8":     IntraocularLensCalculationsStorage,
	"1.2.840.10008.5.1.4.1.1.79.1":     MacularGridThicknessAndVolumeReportStorage,
	"1.2.840.10008.5.1.4.1.1.80.1":     OphthalmicVisualFieldStaticPerimetryMeasurementsStorage,
	"1.2.840.10008.5.1.4.1.1.81.1":     OphthalmicThicknessMapStorage,
	"1.2.840.10008.5.1.4.1.1.82.1":     CornealTopographyMapStorage,
	"1.2.840.10008.5.1.4.1.1.88.1":     TextSRStorageTrialRETIRED,
	"1.2.840.10008.5.1.4.1.1.88.2":     AudioSRStorageTrialRETIRED,
	"1.2.840.10008.5.1.4.1.1.88.3":     DetailSRStorageTrialRETIRED,
	"1.2.840.10008.5.1.4.1.1.88.4":     ComprehensiveSRStorageTrialRETIRED,
	"1.2.840.10008.5.1.4.1.1.88.11":    BasicTextSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.22":    EnhancedSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.33":    ComprehensiveSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.34":    Comprehensive3DSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.35":    ExtensibleSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.40":    ProcedureLogStorage,
	"1.2.840.10008.5.1.4.1.1.88.50":    MammographyCADSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.59":    KeyObjectSelectionDocumentStorage,
	"1.2.840.10008.5.1.4.1.1.88.65":    ChestCADSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.67":    XRayRadiationDoseSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.68":    RadiopharmaceuticalRadiationDoseSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.69":    ColonCADSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.70":    ImplantationPlanSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.71":    AcquisitionContextSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.72":    SimplifiedAdultEchoSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.73":    PatientRadiationDoseSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.74":    PlannedImagingAgentAdministrationSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.75":    PerformedImagingAgentAdministrationSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.76":    EnhancedXRayRadiationDoseSRStorage,
	"1.2.840.10008.5.1.4.1.1.88.77":    WaveformAnnotationSRStorage,
	"1.2.840.10008.5.1.4.1.1.90.1":     ContentAssessmentResultsStorage,
	"1.2.840.10008.5.1.4.1.1.91.1":     MicroscopyBulkSimpleAnnotationsStorage,
	"1.2.840.10008.5.1.4.1.1.104.1":    EncapsulatedPDFStorage,
	"1.2.840.10008.5.1.4.1.1.104.2":    EncapsulatedCDAStorage,
	"1.2.840.10008.5.1.4.1.1.104.3":    EncapsulatedSTLStorage,
	"1.2.840.10008.5.1.4.1.1.104.4":    EncapsulatedOBJStorage,
	"1.2.840.10008.5.1.4.1.1.104.5":    EncapsulatedMTLStorage,
	"1.2.840.10008.5.1.4.1.1.128":      PositronEmissionTomographyImageStorage,
	"1.2.840.10008.5.1.4.1.1.128.1":    LegacyConvertedEnhancedPETImageStorage,
	"1.2.840.10008.5.1.4.1.1.129":      StandalonePETCurveStorageRETIRED,
	"1.2.840.10008.5.1.4.1.1.130":      EnhancedPETImageStorage,
	"1.2.840.10008.5.1.4.1.1.131":      BasicStructuredDisplayStorage,
	"1.2.840.10008.5.1.4.1.1.200.1":    CTDefinedProcedureProtocolStorage,
	"1.2.840.10008.5.1.4.1.1.200.2":    CTPerformedProcedureProtocolStorage,
	"1.2.840.10008.5.1.4.1.1.200.3":    ProtocolApprovalStorage,
	"1.2.840.10008.5.1.4.1.1.200.4":    ProtocolApprovalInformationModelFind,
	"1.2.840.10008.5.1.4.1.1.200.5":    ProtocolApprovalInformationModelMove,
	"1.2.840.10008.5.1.4.1.1.200.6":    ProtocolApprovalInformationModelGet,
	"1.2.840.10008.5.1.4.1.1.200.7":    XADefinedProcedureProtocolStorage,
	"1.2.840.10008.5.1.4.1.1.200.8":    XAPerformedProcedureProtocolStorage,
	"1.2.840.10008.5.1.4.1.1.201.1":    InventoryStorage,
	"1.2.840.10008.5.1.4.1.1.201.2":    InventoryFind,
	"1.2.840.10008.5.1.4.1.1.201.3":    InventoryMove,
	"1.2.840.10008.5.1.4.1.1.201.4":    InventoryGet,
	"1.2.840.10008.5.1.4.1.1.201.5":    InventoryCreation,
	"1.2.840.10008.5.1.4.1.1.201.6":    RepositoryQuery,
	"1.2.840.10008.5.1.4.1.1.201.1.1":  StorageManagementInstance,
	"1.2.840.10008.5.1.4.1.1.481.1":    RTImageStorage,
	"1.2.840.10008.5.1.4.1.1.481.2":    RTDoseStorage,
	"1.2.840.10008.5.1.4.1.1.481.3":    RTStructureSetStorage,
	"1.2.840.10008.5.1.4.1.1.481.4":    RTBeamsTreatmentRecordStorage,
	"1.2.840.10008.5.1.4.1.1.481.5":    RTPlanStorage,
	"1.2.840.10008.5.1.4.1.1.481.6":    RTBrachyTreatmentRecordStorage,
	"1.2.840.10008.5.1.4.1.1.481.7":    RTTreatmentSummaryRecordStorage,
	"1.2.840.10008.5.1.4.1.1.481.8":    RTIonPlanStorage,
	"1.2.840.10008.5.1.4.1.1.481.9":    RTIonBeamsTreatmentRecordStorage,
	"1.2.840.10008.5.1.4.1.1.481.10":   RTPhysicianIntentStorage,
	"1.2.840.10008.5.1.4.1.1.481.11":   RTSegmentAnnotationStorage,
	"1.2.840.10008.5.1.4.1.1.481.12":   RTRadiationSetStorage,
	"1.2.840.10008.5.1.4.1.1.481.13":   CArmPhotonElectronRadiationStorage,
	"1.2.840.10008.5.1.4.1.1.481.14":   TomotherapeuticRadiationStorage,
	"1.2.840.10008.5.1.4.1.1.481.15":   RoboticArmRadiationStorage,
	"1.2.840.10008.5.1.4.1.1.481.16":   RTRadiationRecordSetStorage,
	"1.2.840.10008.5.1.4.1.1.481.17":   RTRadiationSalvageRecordStorage,
	"1.2.840.10008.5.1.4.1.1.481.18":   TomotherapeuticRadiationRecordStorage,
	"1.2.840.10008.5.1.4.1.1.481.19":   CArmPhotonElectronRadiationRecordStorage,
	"1.2.840.10008.5.1.4.1.1.481.20":   RoboticRadiationRecordStorage,
	"1.2.840.10008.5.1.4.1.1.481.21":   RTRadiationSetDeliveryInstructionStorage,
	"1.2.840.10008.5.1.4.1.1.481.22":   RTTreatmentPreparationStorage,
	"1.2.840.10008.5.1.4.1.1.481.23":   EnhancedRTImageStorage,
	"1.2.840.10008.5.1.4.1.1.481.24":   EnhancedContinuousRTImageStorage,
	"1.2.840.10008.5.1.4.1.1.481.25":   RTPatientPositionAcquisitionInstructionStorage,
	"1.2.840.10008.5.1.4.1.1.501.1":    DICOSCTImageStorage,
	"1.2.840.10008.5.1.4.1.1.501.2.1":  DICOSDigitalXRayImageStorageForPresentation,
	"1.2.840.10008.5.1.4.1.1.501.2.2":  DICOSDigitalXRayImageStorageForProcessing,
	"1.2.840.10008.5.1.4.1.1.501.3":    DICOSThreatDetectionReportStorage,
	"1.2.840.10008.5.1.4.1.1.501.4":    DICOS2DAITStorage,
	"1.2.840.10008.5.1.4.1.1.501.5":    DICOS3DAITStorage,
	"1.2.840.10008.5.1.4.1.1.501.6":    DICOSQuadrupoleResonanceStorage,
	"1.2.840.10008.5.1.4.1.1.601.1":    EddyCurrentImageStorage,
	"1.2.840.10008.5.1.4.1.1.601.2":    EddyCurrentMultiFrameImageStorage,
	"1.2.840.10008.5.1.4.1.1.601.3":    ThermographyImageStorage,
	"1.2.840.10008.5.1.4.1.1.601.4":    ThermographyMultiFrameImageStorage,
	"1.2.840.10008.5.1.4.1.1.601.5":    UltrasoundWaveformStorage,
	"1.2.840.10008.5.1.4.1.2.1.1":      PatientRootQueryRetrieveInformationModelFind,
	"1.2.840.10008.5.1.4.1.2.1.2":      PatientRootQueryRetrieveInformationModelMove,
	"1.2.840.10008.5.1.4.1.2.1.3":      PatientRootQueryRetrieveInformationModelGet,
	"1.2.840.10008.5.1.4.1.2.2.1":      StudyRootQueryRetrieveInformationModelFind,
	"1.2.840.10008.5.1.4.1.2.2.2":      StudyRootQueryRetrieveInformationModelMove,
	"1.2.840.10008.5.1.4.1.2.2.3":      StudyRootQueryRetrieveInformationModelGet,
	"1.2.840.10008.5.1.4.1.2.3.1":      PatientStudyOnlyQueryRetrieveInformationModelFindRETIRED,
	"1.2.840.10008.5.1.4.1.2.3.2":      PatientStudyOnlyQueryRetrieveInformationModelMoveRETIRED,
	"1.2.840.10008.5.1.4.1.2.3.3":      PatientStudyOnlyQueryRetrieveInformationModelGetRETIRED,
	"1.2.840.10008.5.1.4.1.2.4.2":      CompositeInstanceRootRetrieveMove,
	"1.2.840.10008.5.1.4.1.2.4.3":      CompositeInstanceRootRetrieveGet,
	"1.2.840.10008.5.1.4.1.2.5.3":      CompositeInstanceRetrieveWithoutBulkDataGet,
	"1.2.840.10008.5.1.4.20.1":         DefinedProcedureProtocolInformationModelFind,
	"1.2.840.10008.5.1.4.20.2":         DefinedProcedureProtocolInformationModelMove,
	"1.2.840.10008.5.1.4.20.3":         DefinedProcedureProtocolInformationModelGet,
	"1.2.840.10008.5.1.4.31":           ModalityWorklistInformationModelFind,
	"1.2.840.10008.5.1.4.32":           GeneralPurposeWorklistManagementMetaRETIRED,
	"1.2.840.10008.5.1.4.32.1":         GeneralPurposeWorklistInformationModelFindRETIRED,
	"1.2.840.10008.5.1.4.32.2":         GeneralPurposeScheduledProcedureStepRETIRED,
	"1.2.840.10008.5.1.4.32.3":         GeneralPurposePerformedProcedureStepRETIRED,
	"1.2.840.10008.5.1.4.33":           InstanceAvailabilityNotification,
	"1.2.840.10008.5.1.4.34.1":         RTBeamsDeliveryInstructionStorageTrialRETIRED,
	"1.2.840.10008.5.1.4.34.2":         RTConventionalMachineVerificationTrialRETIRED,
	"1.2.840.10008.5.1.4.34.3":         RTIonMachineVerificationTrialRETIRED,
	"1.2.840.10008.5.1.4.34.4":         UnifiedWorklistAndProcedureStepTrialRETIRED,
	"1.2.840.10008.5.1.4.34.4.1":       UnifiedProcedureStepPushTrialRETIRED,
	"1.2.840.10008.5.1.4.34.4.2":       UnifiedProcedureStepWatchTrialRETIRED,
	"1.2.840.10008.5.1.4.34.4.3":       UnifiedProcedureStepPullTrialRETIRED,
	"1.2.840.10008.5.1.4.34.4.4":       UnifiedProcedureStepEventTrialRETIRED,
	"1.2.840.10008.5.1.4.34.5":         UPSGlobalSubscriptionInstance,
	"1.2.840.10008.5.1.4.34.5.1":       UPSFilteredGlobalSubscriptionInstance,
	"1.2.840.10008.5.1.4.34.6":         UnifiedWorklistAndProcedureStep,
	"1.2.840.10008.5.1.4.34.6.1":       UnifiedProcedureStepPush,
	"1.2.840.10008.5.1.4.34.6.2":       UnifiedProcedureStepWatch,
	"1.2.840.10008.5.1.4.34.6.3":       UnifiedProcedureStepPull,
	"1.2.840.10008.5.1.4.34.6.4":       UnifiedProcedureStepEvent,
	"1.2.840.10008.5.1.4.34.6.5":       UnifiedProcedureStepQuery,
	"1.2.840.10008.5.1.4.34.7":         RTBeamsDeliveryInstructionStorage,
	"1.2.840.10008.5.1.4.34.8":         RTConventionalMachineVerification,
	"1.2.840.10008.5.1.4.34.9":         RTIonMachineVerification,
	"1.2.840.10008.5.1.4.34.10":        RTBrachyApplicationSetupDeliveryInstructionStorage,
	"1.2.840.10008.5.1.4.37.1":         GeneralRelevantPatientInformationQuery,
	"1.2.840.10008.5.1.4.37.2":         BreastImagingRelevantPatientInformationQuery,
	"1.2.840.10008.5.1.4.37.3":         CardiacRelevantPatientInformationQuery,
	"1.2.840.10008.5.1.4.38.1":         HangingProtocolStorage,
	"1.2.840.10008.5.1.4.38.2":         HangingProtocolInformationModelFind,
	"1.2.840.10008.5.1.4.38.3":         HangingProtocolInformationModelMove,
	"1.2.840.10008.5.1.4.38.4":         HangingProtocolInformationModelGet,
	"1.2.840.10008.5.1.4.39.1":         ColorPaletteStorage,
	"1.2.840.10008.5.1.4.39.2":         ColorPaletteQueryRetrieveInformationModelFind,
	"1.2.840.10008.5.1.4.39.3":         ColorPaletteQueryRetrieveInformationModelMove,
	"1.2.840.10008.5.1.4.39.4":         ColorPaletteQueryRetrieveInformationModelGet,
	"1.2.840.10008.5.1.4.41":           ProductCharacteristicsQuery,
	"1.2.840.10008.5.1.4.42":           SubstanceApprovalQuery,
	"1.2.840.10008.5.1.4.43.1":         GenericImplantTemplateStorage,
	"1.2.840.10008.5.1.4.43.2":         GenericImplantTemplateInformationModelFind,
	"1.2.840.10008.5.1.4.43.3":         GenericImplantTemplateInformationModelMove,
	"1.2.840.10008.5.1.4.43.4":         GenericImplantTemplateInformationModelGet,
	"1.2.840.10008.5.1.4.44.1":         ImplantAssemblyTemplateStorage,
	"1.2.840.10008.5.1.4.44.2":         ImplantAssemblyTemplateInformationModelFind,
	"1.2.840.10008.5.1.4.44.3":         ImplantAssemblyTemplateInformationModelMove,
	"1.2.840.10008.5.1.4.44.4":         ImplantAssemblyTemplateInformationModelGet,
	"1.2.840.10008.5.1.4.45.1":         ImplantTemplateGroupStorage,
	"1.2.840.10008.5.1.4.45.2":         ImplantTemplateGroupInformationModelFind,
	"1.2.840.10008.5.1.4.45.3":         ImplantTemplateGroupInformationModelMove,
	"1.2.840.10008.5.1.4.45.4":         ImplantTemplateGroupInformationModelGet,
	"1.2.840.10008.7.1.1":              NativeDICOMModel,
	"1.2.840.10008.7.1.2":              AbstractMultiDimensionalImageModel,
	"1.2.840.10008.8.1.1":              DICOMContentMappingResource,
	"1.2.840.10008.10.1":               VideoEndoscopicImageRealTimeCommunication,
	"1.2.840.10008.10.2":               VideoPhotographicImageRealTimeCommunication,
	"1.2.840.10008.10.3":               AudioWaveformRealTimeCommunication,
	"1.2.840.10008.10.4":               RenditionSelectionDocumentRealTimeCommunication,
	"1.2.840.10008.15.0.3.1":           dicomDeviceName,
	"1.2.840.10008.15.0.3.2":           dicomDescription,
	"1.2.840.10008.15.0.3.3":           dicomManufacturer,
	"1.2.840.10008.15.0.3.4":           dicomManufacturerModelName,
	"1.2.840.10008.15.0.3.5":           dicomSoftwareVersion,
	"1.2.840.10008.15.0.3.6":           dicomVendorData,
	"1.2.840.10008.15.0.3.7":           dicomAETitle,
	"1.2.840.10008.15.0.3.8":           dicomNetworkConnectionReference,
	"1.2.840.10008.15.0.3.9":           dicomApplicationCluster,
	"1.2.840.10008.15.0.3.10":          dicomAssociationInitiator,
	"1.2.840.10008.15.0.3.11":          dicomAssociationAcceptor,
	"1.2.840.10008.15.0.3.12":          dicomHostname,
	"1.2.840.10008.15.0.3.13":          dicomPort,
	"1.2.840.10008.15.0.3.14":          dicomSOPClass,
	"1.2.840.10008.15.0.3.15":          dicomTransferRole,
	"1.2.840.10008.15.0.3.16":          dicomTransferSyntax,
	"1.2.840.10008.15.0.3.17":          dicomPrimaryDeviceType,
	"1.2.840.10008.15.0.3.18":          dicomRelatedDeviceReference,
	"1.2.840.10008.15.0.3.19":          dicomPreferredCalledAETitle,
	"1.2.840.10008.15.0.3.20":          dicomTLSCyphersuite,
	"1.2.840.10008.15.0.3.21":          dicomAuthorizedNodeCertificateReference,
	"1.2.840.10008.15.0.3.22":          dicomThisNodeCertificateReference,
	"1.2.840.10008.15.0.3.23":          dicomInstalled,
	"1.2.840.10008.15.0.3.24":          dicomStationName,
	"1.2.840.10008.15.0.3.25":          dicomDeviceSerialNumber,
	"1.2.840.10008.15.0.3.26":          dicomInstitutionName,
	"1.2.840.10008.15.0.3.27":          dicomInstitutionAddress,
	"1.2.840.10008.15.0.3.28":          dicomInstitutionDepartmentName,
	"1.2.840.10008.15.0.3.29":          dicomIssuerOfPatientID,
	"1.2.840.10008.15.0.3.30":          dicomPreferredCallingAETitle,
	"1.2.840.10008.15.0.3.31":          dicomSupportedCharacterSet,
	"1.2.840.10008.15.0.4.1":           dicomConfigurationRoot,
	"1.2.840.10008.15.0.4.2":           dicomDevicesRoot,
	"1.2.840.10008.15.0.4.3":           dicomUniqueAETitlesRegistryRoot,
	"1.2.840.10008.15.0.4.4":           dicomDevice,
	"1.2.840.10008.15.0.4.5":           dicomNetworkAE,
	"1.2.840.10008.15.0.4.6":           dicomNetworkConnection,
	"1.2.840.10008.15.0.4.7":           dicomUniqueAETitle,
	"1.2.840.10008.15.0.4.8":           dicomTransferCapability,
	"1.2.840.10008.15.1.1":             UTC,
	"1.2.840.10008.6.1.1":              AnatomicModifier2,
	"1.2.840.10008.6.1.2":              AnatomicRegion4,
	"1.2.840.10008.6.1.3":              TransducerApproach5,
	"1.2.840.10008.6.1.4":              TransducerOrientation6,
	"1.2.840.10008.6.1.5":              UltrasoundBeamPath7,
	"1.2.840.10008.6.1.6":              AngiographicInterventionalDevice8,
	"1.2.840.10008.6.1.7":              ImageGuidedTherapeuticProcedure9,
	"1.2.840.10008.6.1.8":              InterventionalDrug10,
	"1.2.840.10008.6.1.9":              AdministrationRoute11,
	"1.2.840.10008.6.1.10":             ImagingContrastAgent12,
	"1.2.840.10008.6.1.11":             ImagingContrastAgentIngredient13,
	"1.2.840.10008.6.1.12":             RadiopharmaceuticalIsotope18,
	"1.2.840.10008.6.1.13":             PatientOrientation19,
	"1.2.840.10008.6.1.14":             PatientOrientationModifier20,
	"1.2.840.10008.6.1.15":             PatientEquipmentRelationship21,
	"1.2.840.10008.6.1.16":             CranioCaudadAngulation23,
	"1.2.840.10008.6.1.17":             Radiopharmaceutical25,
	"1.2.840.10008.6.1.18":             NuclearMedicineProjection26,
	"1.2.840.10008.6.1.19":             AcquisitionModality29,
	"1.2.840.10008.6.1.20":             DICOMDevice30,
	"1.2.840.10008.6.1.21":             AbstractPrior31,
	"1.2.840.10008.6.1.22":             NumericValueQualifier42,
	"1.2.840.10008.6.1.23":             MeasurementUnit82,
	"1.2.840.10008.6.1.24":             RealWorldValueMappingUnit83,
	"1.2.840.10008.6.1.25":             SignificanceLevel220,
	"1.2.840.10008.6.1.26":             MeasurementRangeConcept221,
	"1.2.840.10008.6.1.27":             Normality222,
	"1.2.840.10008.6.1.28":             NormalRangeValue223,
	"1.2.840.10008.6.1.29":             SelectionMethod224,
	"1.2.840.10008.6.1.30":             MeasurementUncertaintyConcept225,
	"1.2.840.10008.6.1.31":             PopulationStatisticalDescriptor226,
	"1.2.840.10008.6.1.32":             SampleStatisticalDescriptor227,
	"1.2.840.10008.6.1.33":             EquationOrTable228,
	"1.2.840.10008.6.1.34":             YesNo230,
	"1.2.840.10008.6.1.35":             PresentAbsent240,
	"1.2.840.10008.6.1.36":             NormalAbnormal242,
	"1.2.840.10008.6.1.37":             Laterality244,
	"1.2.840.10008.6.1.38":             PositiveNegative250,
	"1.2.840.10008.6.1.39":             ComplicationSeverity251,
	"1.2.840.10008.6.1.40":             ObserverType270,
	"1.2.840.10008.6.1.41":             ObservationSubjectClass271,
	"1.2.840.10008.6.1.42":             AudioChannelSource3000,
	"1.2.840.10008.6.1.43":             ECGLead3001,
	"1.2.840.10008.6.1.44":             HemodynamicWaveformSource3003,
	"1.2.840.10008.6.1.45":             CardiovascularAnatomicStructure3010,
	"1.2.840.10008.6.1.46":             ElectrophysiologyAnatomicLocation3011,
	"1.2.840.10008.6.1.47":             CoronaryArterySegment3014,
	"1.2.840.10008.6.1.48":             CoronaryArtery3015,
	"1.2.840.10008.6.1.49":             CardiovascularAnatomicStructureModifier3019,
	"1.2.840.10008.6.1.50":             CardiologyMeasurementUnit3082RETIRED,
	"1.2.840.10008.6.1.51":             TimeSynchronizationChannelType3090,
	"1.2.840.10008.6.1.52":             CardiacProceduralStateValue3101,
	"1.2.840.10008.6.1.53":             ElectrophysiologyMeasurementFunctionTechnique3240,
	"1.2.840.10008.6.1.54":             HemodynamicMeasurementTechnique3241,
	"1.2.840.10008.6.1.55":             CatheterizationProcedurePhase3250,
	"1.2.840.10008.6.1.56":             ElectrophysiologyProcedurePhase3254,
	"1.2.840.10008.6.1.57":             StressProtocol3261,
	"1.2.840.10008.6.1.58":             ECGPatientStateValue3262,
	"1.2.840.10008.6.1.59":             ElectrodePlacementValue3263,
	"1.2.840.10008.6.1.60":             XYZElectrodePlacementValues3264RETIRED,
	"1.2.840.10008.6.1.61":             HemodynamicPhysiologicalChallenge3271,
	"1.2.840.10008.6.1.62":             ECGAnnotation3335,
	"1.2.840.10008.6.1.63":             HemodynamicAnnotation3337,
	"1.2.840.10008.6.1.64":             ElectrophysiologyAnnotation3339,
	"1.2.840.10008.6.1.65":             ProcedureLogTitle3400,
	"1.2.840.10008.6.1.66":             LogNoteType3401,
	"1.2.840.10008.6.1.67":             PatientStatusAndEvent3402,
	"1.2.840.10008.6.1.68":             PercutaneousEntry3403,
	"1.2.840.10008.6.1.69":             StaffAction3404,
	"1.2.840.10008.6.1.70":             ProcedureActionValue3405,
	"1.2.840.10008.6.1.71":             NonCoronaryTranscatheterIntervention3406,
	"1.2.840.10008.6.1.72":             ObjectReferencePurpose3407,
	"1.2.840.10008.6.1.73":             ConsumableAction3408,
	"1.2.840.10008.6.1.74":             DrugContrastAdministration3409,
	"1.2.840.10008.6.1.75":             DrugContrastNumericParameter3410,
	"1.2.840.10008.6.1.76":             IntracoronaryDevice3411,
	"1.2.840.10008.6.1.77":             InterventionActionStatus3412,
	"1.2.840.10008.6.1.78":             AdverseOutcome3413,
	"1.2.840.10008.6.1.79":             ProcedureUrgency3414,
	"1.2.840.10008.6.1.80":             CardiacRhythm3415,
	"1.2.840.10008.6.1.81":             RespirationRhythm3416,
	"1.2.840.10008.6.1.82":             LesionRisk3418,
	"1.2.840.10008.6.1.83":             FindingTitle3419,
	"1.2.840.10008.6.1.84":             ProcedureAction3421,
	"1.2.840.10008.6.1.85":             DeviceUseAction3422,
	"1.2.840.10008.6.1.86":             NumericDeviceCharacteristic3423,
	"1.2.840.10008.6.1.87":             InterventionParameter3425,
	"1.2.840.10008.6.1.88":             ConsumablesParameter3426,
	"1.2.840.10008.6.1.89":             EquipmentEvent3427,
	"1.2.840.10008.6.1.90":             CardiovascularImagingProcedure3428,
	"1.2.840.10008.6.1.91":             CatheterizationDevice3429,
	"1.2.840.10008.6.1.92":             DateTimeQualifier3430,
	"1.2.840.10008.6.1.93":             PeripheralPulseLocation3440,
	"1.2.840.10008.6.1.94":             PatientAssessment3441,
	"1.2.840.10008.6.1.95":             PeripheralPulseMethod3442,
	"1.2.840.10008.6.1.96":             SkinCondition3446,
	"1.2.840.10008.6.1.97":             AirwayAssessment3448,
	"1.2.840.10008.6.1.98":             CalibrationObject3451,
	"1.2.840.10008.6.1.99":             CalibrationMethod3452,
	"1.2.840.10008.6.1.100":            CardiacVolumeMethod3453,
	"1.2.840.10008.6.1.101":            IndexMethod3455,
	"1.2.840.10008.6.1.102":            SubSegmentMethod3456,
	"1.2.840.10008.6.1.103":            ContourRealignment3458,
	"1.2.840.10008.6.1.104":            CircumferentialExtent3460,
	"1.2.840.10008.6.1.105":            RegionalExtent3461,
	"1.2.840.10008.6.1.106":            ChamberIdentification3462,
	"1.2.840.10008.6.1.107":            QAReferenceMethod3465,
	"1.2.840.10008.6.1.108":            PlaneIdentification3466,
	"1.2.840.10008.6.1.109":            EjectionFraction3467,
	"1.2.840.10008.6.1.110":            EDVolume3468,
	"1.2.840.10008.6.1.111":            ESVolume3469,
	"1.2.840.10008.6.1.112":            VesselLumenCrossSectionalAreaCalculationMethod3470,
	"1.2.840.10008.6.1.113":            EstimatedVolume3471,
	"1.2.840.10008.6.1.114":            CardiacContractionPhase3472,
	"1.2.840.10008.6.1.115":            IVUSProcedurePhase3480,
	"1.2.840.10008.6.1.116":            IVUSDistanceMeasurement3481,
	"1.2.840.10008.6.1.117":            IVUSAreaMeasurement3482,
	"1.2.840.10008.6.1.118":            IVUSLongitudinalMeasurement3483,
	"1.2.840.10008.6.1.119":            IVUSIndexRatio3484,
	"1.2.840.10008.6.1.120":            IVUSVolumeMeasurement3485,
	"1.2.840.10008.6.1.121":            VascularMeasurementSite3486,
	"1.2.840.10008.6.1.122":            IntravascularVolumetricRegion3487,
	"1.2.840.10008.6.1.123":            MinMaxMean3488,
	"1.2.840.10008.6.1.124":            CalciumDistribution3489,
	"1.2.840.10008.6.1.125":            IVUSLesionMorphology3491,
	"1.2.840.10008.6.1.126":            VascularDissectionClassification3492,
	"1.2.840.10008.6.1.127":            IVUSRelativeStenosisSeverity3493,
	"1.2.840.10008.6.1.128":            IVUSNonMorphologicalFinding3494,
	"1.2.840.10008.6.1.129":            IVUSPlaqueComposition3495,
	"1.2.840.10008.6.1.130":            IVUSFiducialPoint3496,
	"1.2.840.10008.6.1.131":            IVUSArterialMorphology3497,
	"1.2.840.10008.6.1.132":            PressureUnit3500,
	"1.2.840.10008.6.1.133":            HemodynamicResistanceUnit3502,
	"1.2.840.10008.6.1.134":            IndexedHemodynamicResistanceUnit3503,
	"1.2.840.10008.6.1.135":            CatheterSizeUnit3510,
	"1.2.840.10008.6.1.136":            SpecimenCollection3515,
	"1.2.840.10008.6.1.137":            BloodSourceType3520,
	"1.2.840.10008.6.1.138":            BloodGasPressure3524,
	"1.2.840.10008.6.1.139":            BloodGasContent3525,
	"1.2.840.10008.6.1.140":            BloodGasSaturation3526,
	"1.2.840.10008.6.1.141":            BloodBaseExcess3527,
	"1.2.840.10008.6.1.142":            BloodPH3528,
	"1.2.840.10008.6.1.143":            ArterialVenousContent3529,
	"1.2.840.10008.6.1.144":            OxygenAdministrationAction3530,
	"1.2.840.10008.6.1.145":            OxygenAdministration3531,
	"1.2.840.10008.6.1.146":            CirculatorySupportAction3550,
	"1.2.840.10008.6.1.147":            VentilationAction3551,
	"1.2.840.10008.6.1.148":            PacingAction3552,
	"1.2.840.10008.6.1.149":            CirculatorySupport3553,
	"1.2.840.10008.6.1.150":            Ventilation3554,
	"1.2.840.10008.6.1.151":            Pacing3555,
	"1.2.840.10008.6.1.152":            BloodPressureMethod3560,
	"1.2.840.10008.6.1.153":            RelativeTime3600,
	"1.2.840.10008.6.1.154":            HemodynamicPatientState3602,
	"1.2.840.10008.6.1.155":            ArterialLesionLocation3604,
	"1.2.840.10008.6.1.156":            ArterialSourceLocation3606,
	"1.2.840.10008.6.1.157":            VenousSourceLocation3607,
	"1.2.840.10008.6.1.158":            AtrialSourceLocation3608,
	"1.2.840.10008.6.1.159":            VentricularSourceLocation3609,
	"1.2.840.10008.6.1.160":            GradientSourceLocation3610,
	"1.2.840.10008.6.1.161":            PressureMeasurement3611,
	"1.2.840.10008.6.1.162":            BloodVelocityMeasurement3612,
	"1.2.840.10008.6.1.163":            HemodynamicTimeMeasurement3613,
	"1.2.840.10008.6.1.164":            NonMitralValveArea3614,
	"1.2.840.10008.6.1.165":            ValveArea3615,
	"1.2.840.10008.6.1.166":            HemodynamicPeriodMeasurement3616,
	"1.2.840.10008.6.1.167":            ValveFlow3617,
	"1.2.840.10008.6.1.168":            HemodynamicFlow3618,
	"1.2.840.10008.6.1.169":            HemodynamicResistanceMeasurement3619,
	"1.2.840.10008.6.1.170":            HemodynamicRatio3620,
	"1.2.840.10008.6.1.171":            FractionalFlowReserve3621,
	"1.2.840.10008.6.1.172":            MeasurementType3627,
	"1.2.840.10008.6.1.173":            CardiacOutputMethod3628,
	"1.2.840.10008.6.1.174":            ProcedureIntent3629,
	"1.2.840.10008.6.1.175":            CardiovascularAnatomicLocation3630,
	"1.2.840.10008.6.1.176":            Hypertension3640,
	"1.2.840.10008.6.1.177":            HemodynamicAssessment3641,
	"1.2.840.10008.6.1.178":            DegreeFinding3642,
	"1.2.840.10008.6.1.179":            HemodynamicMeasurementPhase3651,
	"1.2.840.10008.6.1.180":            BodySurfaceAreaEquation3663,
	"1.2.840.10008.6.1.181":            OxygenConsumptionEquationTable3664,
	"1.2.840.10008.6.1.182":            P50Equation3666,
	"1.2.840.10008.6.1.183":            FraminghamScore3667,
	"1.2.840.10008.6.1.184":            FraminghamTable3668,
	"1.2.840.10008.6.1.185":            ECGProcedureType3670,
	"1.2.840.10008.6.1.186":            ReasonForECGStudy3671,
	"1.2.840.10008.6.1.187":            Pacemaker3672,
	"1.2.840.10008.6.1.188":            Diagnosis3673RETIRED,
	"1.2.840.10008.6.1.189":            OtherFilters3675RETIRED,
	"1.2.840.10008.6.1.190":            LeadMeasurementTechnique3676,
	"1.2.840.10008.6.1.191":            SummaryCodesECG3677,
	"1.2.840.10008.6.1.192":            QTCorrectionAlgorithm3678,
	"1.2.840.10008.6.1.193":            ECGMorphologyDescription3679RETIRED,
	"1.2.840.10008.6.1.194":            ECGLeadNoiseDescription3680,
	"1.2.840.10008.6.1.195":            ECGLeadNoiseModifier3681RETIRED,
	"1.2.840.10008.6.1.196":            Probability3682RETIRED,
	"1.2.840.10008.6.1.197":            Modifier3683RETIRED,
	"1.2.840.10008.6.1.198":            Trend3684RETIRED,
	"1.2.840.10008.6.1.199":            ConjunctiveTerm3685RETIRED,
	"1.2.840.10008.6.1.200":            ECGInterpretiveStatement3686RETIRED,
	"1.2.840.10008.6.1.201":            ElectrophysiologyWaveformDuration3687,
	"1.2.840.10008.6.1.202":            ElectrophysiologyWaveformVoltage3688,
	"1.2.840.10008.6.1.203":            CathDiagnosis3700,
	"1.2.840.10008.6.1.204":            CardiacValveTract3701,
	"1.2.840.10008.6.1.205":            WallMotion3703,
	"1.2.840.10008.6.1.206":            MyocardiumWallMorphologyFinding3704,
	"1.2.840.10008.6.1.207":            ChamberSize3705,
	"1.2.840.10008.6.1.208":            OverallContractility3706,
	"1.2.840.10008.6.1.209":            VSDDescription3707,
	"1.2.840.10008.6.1.210":            AorticRootDescription3709,
	"1.2.840.10008.6.1.211":            CoronaryDominance3710,
	"1.2.840.10008.6.1.212":            ValvularAbnormality3711,
	"1.2.840.10008.6.1.213":            VesselDescriptor3712,
	"1.2.840.10008.6.1.214":            TIMIFlowCharacteristic3713,
	"1.2.840.10008.6.1.215":            Thrombus3714,
	"1.2.840.10008.6.1.216":            LesionMargin3715,
	"1.2.840.10008.6.1.217":            Severity3716,
	"1.2.840.10008.6.1.218":            LeftVentricleMyocardialWall17SegmentModel3717,
	"1.2.840.10008.6.1.219":            MyocardialWallSegmentsInProjection3718,
	"1.2.840.10008.6.1.220":            CanadianClinicalClassification3719,
	"1.2.840.10008.6.1.221":            CardiacHistoryDate3720RETIRED,
	"1.2.840.10008.6.1.222":            CardiovascularSurgery3721,
	"1.2.840.10008.6.1.223":            DiabeticTherapy3722,
	"1.2.840.10008.6.1.224":            MIType3723,
	"1.2.840.10008.6.1.225":            SmokingHistory3724,
	"1.2.840.10008.6.1.226":            CoronaryInterventionIndication3726,
	"1.2.840.10008.6.1.227":            CatheterizationIndication3727,
	"1.2.840.10008.6.1.228":            CathFinding3728,
	"1.2.840.10008.6.1.229":            AdmissionStatus3729,
	"1.2.840.10008.6.1.230":            InsurancePayor3730,
	"1.2.840.10008.6.1.231":            PrimaryCauseOfDeath3733,
	"1.2.840.10008.6.1.232":            AcuteCoronarySyndromeTimePeriod3735,
	"1.2.840.10008.6.1.233":            NYHAClassification3736,
	"1.2.840.10008.6.1.234":            IschemiaNonInvasiveTest3737,
	"1.2.840.10008.6.1.235":            PreCathAnginaType3738,
	"1.2.840.10008.6.1.236":            CathProcedureType3739,
	"1.2.840.10008.6.1.237":            ThrombolyticAdministration3740,
	"1.2.840.10008.6.1.238":            LabVisitMedicationAdministration3741,
	"1.2.840.10008.6.1.239":            PCIMedicationAdministration3742,
	"1.2.840.10008.6.1.240":            ClopidogrelTiclopidineAdministration3743,
	"1.2.840.10008.6.1.241":            EFTestingMethod3744,
	"1.2.840.10008.6.1.242":            CalculationMethod3745,
	"1.2.840.10008.6.1.243":            PercutaneousEntrySite3746,
	"1.2.840.10008.6.1.244":            PercutaneousClosure3747,
	"1.2.840.10008.6.1.245":            AngiographicEFTestingMethod3748,
	"1.2.840.10008.6.1.246":            PCIProcedureResult3749,
	"1.2.840.10008.6.1.247":            PreviouslyDilatedLesion3750,
	"1.2.840.10008.6.1.248":            GuidewireCrossing3752,
	"1.2.840.10008.6.1.249":            VascularComplication3754,
	"1.2.840.10008.6.1.250":            CathComplication3755,
	"1.2.840.10008.6.1.251":            CardiacPatientRiskFactor3756,
	"1.2.840.10008.6.1.252":            CardiacDiagnosticProcedure3757,
	"1.2.840.10008.6.1.253":            CardiovascularFamilyHistory3758,
	"1.2.840.10008.6.1.254":            HypertensionTherapy3760,
	"1.2.840.10008.6.1.255":            AntilipemicAgent3761,
	"1.2.840.10008.6.1.256":            AntiarrhythmicAgent3762,
	"1.2.840.10008.6.1.257":            MyocardialInfarctionTherapy3764,
	"1.2.840.10008.6.1.258":            ConcernType3769,
	"1.2.840.10008.6.1.259":            ProblemStatus3770,
	"1.2.840.10008.6.1.260":            HealthStatus3772,
	"1.2.840.10008.6.1.261":            UseStatus3773,
	"1.2.840.10008.6.1.262":            SocialHistory3774,
	"1.2.840.10008.6.1.263":            CardiovascularImplant3777,
	"1.2.840.10008.6.1.264":            PlaqueStructure3802,
	"1.2.840.10008.6.1.265":            StenosisMeasurementMethod3804,
	"1.2.840.10008.6.1.266":            StenosisType3805,
	"1.2.840.10008.6.1.267":            StenosisShape3806,
	"1.2.840.10008.6.1.268":            VolumeMeasurementMethod3807,
	"1.2.840.10008.6.1.269":            AneurysmType3808,
	"1.2.840.10008.6.1.270":            AssociatedCondition3809,
	"1.2.840.10008.6.1.271":            VascularMorphology3810,
	"1.2.840.10008.6.1.272":            StentFinding3813,
	"1.2.840.10008.6.1.273":            StentComposition3814,
	"1.2.840.10008.6.1.274":            SourceOfVascularFinding3815,
	"1.2.840.10008.6.1.275":            VascularSclerosisType3817,
	"1.2.840.10008.6.1.276":            NonInvasiveVascularProcedure3820,
	"1.2.840.10008.6.1.277":            PapillaryMuscleIncludedExcluded3821,
	"1.2.840.10008.6.1.278":            RespiratoryStatus3823,
	"1.2.840.10008.6.1.279":            HeartRhythm3826,
	"1.2.840.10008.6.1.280":            VesselSegment3827,
	"1.2.840.10008.6.1.281":            PulmonaryArtery3829,
	"1.2.840.10008.6.1.282":            StenosisLength3831,
	"1.2.840.10008.6.1.283":            StenosisGrade3832,
	"1.2.840.10008.6.1.284":            CardiacEjectionFraction3833,
	"1.2.840.10008.6.1.285":            CardiacVolumeMeasurement3835,
	"1.2.840.10008.6.1.286":            TimeBasedPerfusionMeasurement3836,
	"1.2.840.10008.6.1.287":            FiducialFeature3837,
	"1.2.840.10008.6.1.288":            DiameterDerivation3838,
	"1.2.840.10008.6.1.289":            CoronaryVein3839,
	"1.2.840.10008.6.1.290":            PulmonaryVein3840,
	"1.2.840.10008.6.1.291":            MyocardialSubsegment3843,
	"1.2.840.10008.6.1.292":            PartialViewSectionForMammography4005,
	"1.2.840.10008.6.1.293":            DXAnatomyImaged4009,
	"1.2.840.10008.6.1.294":            DXView4010,
	"1.2.840.10008.6.1.295":            DXViewModifier4011,
	"1.2.840.10008.6.1.296":            ProjectionEponymousName4012,
	"1.2.840.10008.6.1.297":            AnatomicRegionForMammography4013,
	"1.2.840.10008.6.1.298":            ViewForMammography4014,
	"1.2.840.10008.6.1.299":            ViewModifierForMammography4015,
	"1.2.840.10008.6.1.300":            AnatomicRegionForIntraOralRadiography4016,
	"1.2.840.10008.6.1.301":            AnatomicRegionModifierForIntraOralRadiography4017,
	"1.2.840.10008.6.1.302":            PrimaryAnatomicStructureForIntraOralRadiographyPermanentDentitionDesignationOfTeeth4018,
	"1.2.840.10008.6.1.303":            PrimaryAnatomicStructureForIntraOralRadiographyDeciduousDentitionDesignationOfTeeth4019,
	"1.2.840.10008.6.1.304":            PETRadionuclide4020,
	"1.2.840.10008.6.1.305":            PETRadiopharmaceutical4021,
	"1.2.840.10008.6.1.306":            CraniofacialAnatomicRegion4028,
	"1.2.840.10008.6.1.307":            CTMRAndPETAnatomyImaged4030,
	"1.2.840.10008.6.1.308":            CommonAnatomicRegion4031,
	"1.2.840.10008.6.1.309":            MRSpectroscopyMetabolite4032,
	"1.2.840.10008.6.1.310":            MRProtonSpectroscopyMetabolite4033,
	"1.2.840.10008.6.1.311":            EndoscopyAnatomicRegion4040,
	"1.2.840.10008.6.1.312":            XAXRFAnatomyImaged4042,
	"1.2.840.10008.6.1.313":            DrugOrContrastAgentCharacteristic4050,
	"1.2.840.10008.6.1.314":            GeneralDevice4051,
	"1.2.840.10008.6.1.315":            PhantomDevice4052,
	"1.2.840.10008.6.1.316":            OphthalmicImagingAgent4200,
	"1.2.840.10008.6.1.317":            PatientEyeMovementCommand4201,
	"1.2.840.10008.6.1.318":            OphthalmicPhotographyAcquisitionDevice4202,
	"1.2.840.10008.6.1.319":            OphthalmicPhotographyIllumination4203,
	"1.2.840.10008.6.1.320":            OphthalmicFilter4204,
	"1.2.840.10008.6.1.321":            OphthalmicLens4205,
	"1.2.840.10008.6.1.322":            OphthalmicChannelDescription4206,
	"1.2.840.10008.6.1.323":            OphthalmicImagePosition4207,
	"1.2.840.10008.6.1.324":            MydriaticAgent4208,
	"1.2.840.10008.6.1.325":            OphthalmicAnatomicStructureImaged4209,
	"1.2.840.10008.6.1.326":            OphthalmicTomographyAcquisitionDevice4210,
	"1.2.840.10008.6.1.327":            OphthalmicOCTAnatomicStructureImaged4211,
	"1.2.840.10008.6.1.328":            Language5000,
	"1.2.840.10008.6.1.329":            Country5001,
	"1.2.840.10008.6.1.330":            OverallBreastComposition6000,
	"1.2.840.10008.6.1.331":            OverallBreastCompositionFromBIRADS6001,
	"1.2.840.10008.6.1.332":            ChangeSinceLastMammogramOrPriorSurgery6002,
	"1.2.840.10008.6.1.333":            ChangeSinceLastMammogramOrPriorSurgeryFromBIRADS6003,
	"1.2.840.10008.6.1.334":            MammographyShapeCharacteristic6004,
	"1.2.840.10008.6.1.335":            ShapeCharacteristicFromBIRADS6005,
	"1.2.840.10008.6.1.336":            MammographyMarginCharacteristic6006,
	"1.2.840.10008.6.1.337":            MarginCharacteristicFromBIRADS6007,
	"1.2.840.10008.6.1.338":            DensityModifier6008,
	"1.2.840.10008.6.1.339":            DensityModifierFromBIRADS6009,
	"1.2.840.10008.6.1.340":            MammographyCalcificationType6010,
	"1.2.840.10008.6.1.341":            CalcificationTypeFromBIRADS6011,
	"1.2.840.10008.6.1.342":            CalcificationDistributionModifier6012,
	"1.2.840.10008.6.1.343":            CalcificationDistributionModifierFromBIRADS6013,
	"1.2.840.10008.6.1.344":            MammographySingleImageFinding6014,
	"1.2.840.10008.6.1.345":            SingleImageFindingFromBIRADS6015,
	"1.2.840.10008.6.1.346":            MammographyCompositeFeature6016,
	"1.2.840.10008.6.1.347":            CompositeFeatureFromBIRADS6017,
	"1.2.840.10008.6.1.348":            ClockfaceLocationOrRegion6018,
	"1.2.840.10008.6.1.349":            ClockfaceLocationOrRegionFromBIRADS6019,
	"1.2.840.10008.6.1.350":            QuadrantLocation6020,
	"1.2.840.10008.6.1.351":            QuadrantLocationFromBIRADS6021,
	"1.2.840.10008.6.1.352":            Side6022,
	"1.2.840.10008.6.1.353":            SideFromBIRADS6023,
	"1.2.840.10008.6.1.354":            Depth6024,
	"1.2.840.10008.6.1.355":            DepthFromBIRADS6025,
	"1.2.840.10008.6.1.356":            MammographyAssessment6026,
	"1.2.840.10008.6.1.357":            AssessmentFromBIRADS6027,
	"1.2.840.10008.6.1.358":            MammographyRecommendedFollowUp6028,
	"1.2.840.10008.6.1.359":            RecommendedFollowUpFromBIRADS6029,
	"1.2.840.10008.6.1.360":            MammographyPathologyCode6030,
	"1.2.840.10008.6.1.361":            BenignPathologyCodeFromBIRADS6031,
	"1.2.840.10008.6.1.362":            HighRiskLesionPathologyCodeFromBIRADS6032,
	"1.2.840.10008.6.1.363":            MalignantPathologyCodeFromBIRADS6033,
	"1.2.840.10008.6.1.364":            CADOutputIntendedUse6034,
	"1.2.840.10008.6.1.365":            CompositeFeatureRelation6035,
	"1.2.840.10008.6.1.366":            FeatureScope6036,
	"1.2.840.10008.6.1.367":            MammographyQuantitativeTemporalDifferenceType6037,
	"1.2.840.10008.6.1.368":            MammographyQualitativeTemporalDifferenceType6038,
	"1.2.840.10008.6.1.369":            NippleCharacteristic6039,
	"1.2.840.10008.6.1.370":            NonLesionObjectType6040,
	"1.2.840.10008.6.1.371":            MammographyImageQualityFinding6041,
	"1.2.840.10008.6.1.372":            ResultStatus6042,
	"1.2.840.10008.6.1.373":            MammographyCADAnalysisType6043,
	"1.2.840.10008.6.1.374":            ImageQualityAssessmentType6044,
	"1.2.840.10008.6.1.375":            MammographyQualityControlStandardType6045,
	"1.2.840.10008.6.1.376":            FollowUpIntervalUnit6046,
	"1.2.840.10008.6.1.377":            CADProcessingAndFindingSummary6047,
	"1.2.840.10008.6.1.378":            CADOperatingPointAxisLabel6048,
	"1.2.840.10008.6.1.379":            BreastProcedureReported6050,
	"1.2.840.10008.6.1.380":            BreastProcedureReason6051,
	"1.2.840.10008.6.1.381":            BreastImagingReportSectionTitle6052,
	"1.2.840.10008.6.1.382":            BreastImagingReportElement6053,
	"1.2.840.10008.6.1.383":            BreastImagingFinding6054,
	"1.2.840.10008.6.1.384":            BreastClinicalFindingOrIndicatedProblem6055,
	"1.2.840.10008.6.1.385":            AssociatedFindingForBreast6056,
	"1.2.840.10008.6.1.386":            DuctographyFindingForBreast6057,
	"1.2.840.10008.6.1.387":            ProcedureModifiersForBreast6058,
	"1.2.840.10008.6.1.388":            BreastImplantType6059,
	"1.2.840.10008.6.1.389":            BreastBiopsyTechnique6060,
	"1.2.840.10008.6.1.390":            BreastImagingProcedureModifier6061,
	"1.2.840.10008.6.1.391":            InterventionalProcedureComplication6062,
	"1.2.840.10008.6.1.392":            InterventionalProcedureResult6063,
	"1.2.840.10008.6.1.393":            UltrasoundFindingForBreast6064,
	"1.2.840.10008.6.1.394":            InstrumentApproach6065,
	"1.2.840.10008.6.1.395":            TargetConfirmation6066,
	"1.2.840.10008.6.1.396":            FluidColor6067,
	"1.2.840.10008.6.1.397":            TumorStagesFromAJCC6068,
	"1.2.840.10008.6.1.398":            NottinghamCombinedHistologicGrade6069,
	"1.2.840.10008.6.1.399":            BloomRichardsonHistologicGrade6070,
	"1.2.840.10008.6.1.400":            HistologicGradingMethod6071,
	"1.2.840.10008.6.1.401":            BreastImplantFinding6072,
	"1.2.840.10008.6.1.402":            GynecologicalHormone6080,
	"1.2.840.10008.6.1.403":            BreastCancerRiskFactor6081,
	"1.2.840.10008.6.1.404":            GynecologicalProcedure6082,
	"1.2.840.10008.6.1.405":            ProceduresForBreast6083,
	"1.2.840.10008.6.1.406":            MammoplastyProcedure6084,
	"1.2.840.10008.6.1.407":            TherapiesForBreast6085,
	"1.2.840.10008.6.1.408":            MenopausalPhase6086,
	"1.2.840.10008.6.1.409":            GeneralRiskFactor6087,
	"1.2.840.10008.6.1.410":            OBGYNMaternalRiskFactor6088,
	"1.2.840.10008.6.1.411":            Substance6089,
	"1.2.840.10008.6.1.412":            RelativeUsageExposureAmount6090,
	"1.2.840.10008.6.1.413":            RelativeFrequencyOfEventValue6091,
	"1.2.840.10008.6.1.414":            UsageExposureQualitativeConcept6092,
	"1.2.840.10008.6.1.415":            UsageExposureAmountQualitativeConcept6093,
	"1.2.840.10008.6.1.416":            UsageExposureFrequencyQualitativeConcept6094,
	"1.2.840.10008.6.1.417":            ProcedureNumericProperty6095,
	"1.2.840.10008.6.1.418":            PregnancyStatus6096,
	"1.2.840.10008.6.1.419":            SideOfFamily6097,
	"1.2.840.10008.6.1.420":            ChestComponentCategory6100,
	"1.2.840.10008.6.1.421":            ChestFindingOrFeature6101,
	"1.2.840.10008.6.1.422":            ChestFindingOrFeatureModifier6102,
	"1.2.840.10008.6.1.423":            AbnormalLinesFindingOrFeature6103,
	"1.2.840.10008.6.1.424":            AbnormalOpacityFindingOrFeature6104,
	"1.2.840.10008.6.1.425":            AbnormalLucencyFindingOrFeature6105,
	"1.2.840.10008.6.1.426":            AbnormalTextureFindingOrFeature6106,
	"1.2.840.10008.6.1.427":            WidthDescriptor6107,
	"1.2.840.10008.6.1.428":            ChestAnatomicStructureAbnormalDistribution6108,
	"1.2.840.10008.6.1.429":            RadiographicAnatomyFindingOrFeature6109,
	"1.2.840.10008.6.1.430":            LungAnatomyFindingOrFeature6110,
	"1.2.840.10008.6.1.431":            BronchovascularAnatomyFindingOrFeature6111,
	"1.2.840.10008.6.1.432":            PleuraAnatomyFindingOrFeature6112,
	"1.2.840.10008.6.1.433":            MediastinumAnatomyFindingOrFeature6113,
	"1.2.840.10008.6.1.434":            OsseousAnatomyFindingOrFeature6114,
	"1.2.840.10008.6.1.435":            OsseousAnatomyModifier6115,
	"1.2.840.10008.6.1.436":            MuscularAnatomy6116,
	"1.2.840.10008.6.1.437":            VascularAnatomy6117,
	"1.2.840.10008.6.1.438":            SizeDescriptor6118,
	"1.2.840.10008.6.1.439":            ChestBorderShape6119,
	"1.2.840.10008.6.1.440":            ChestBorderDefinition6120,
	"1.2.840.10008.6.1.441":            ChestOrientationDescriptor6121,
	"1.2.840.10008.6.1.442":            ChestContentDescriptor6122,
	"1.2.840.10008.6.1.443":            ChestOpacityDescriptor6123,
	"1.2.840.10008.6.1.444":            LocationInChest6124,
	"1.2.840.10008.6.1.445":            GeneralChestLocation6125,
	"1.2.840.10008.6.1.446":            LocationInLung6126,
	"1.2.840.10008.6.1.447":            SegmentLocationInLung6127,
	"1.2.840.10008.6.1.448":            ChestDistributionDescriptor6128,
	"1.2.840.10008.6.1.449":            ChestSiteInvolvement6129,
	"1.2.840.10008.6.1.450":            SeverityDescriptor6130,
	"1.2.840.10008.6.1.451":            ChestTextureDescriptor6131,
	"1.2.840.10008.6.1.452":            ChestCalcificationDescriptor6132,
	"1.2.840.10008.6.1.453":            ChestQuantitativeTemporalDifferenceType6133,
	"1.2.840.10008.6.1.454":            ChestQualitativeTemporalDifferenceType6134,
	"1.2.840.10008.6.1.455":            ImageQualityFinding6135,
	"1.2.840.10008.6.1.456":            ChestTypesOfQualityControlStandard6136,
	"1.2.840.10008.6.1.457":            CADAnalysisType6137,
	"1.2.840.10008.6.1.458":            ChestNonLesionObjectType6138,
	"1.2.840.10008.6.1.459":            NonLesionModifier6139,
	"1.2.840.10008.6.1.460":            CalculationMethod6140,
	"1.2.840.10008.6.1.461":            AttenuationCoefficientMeasurement6141,
	"1.2.840.10008.6.1.462":            CalculatedValue6142,
	"1.2.840.10008.6.1.463":            LesionResponse6143,
	"1.2.840.10008.6.1.464":            RECISTDefinedLesionResponse6144,
	"1.2.840.10008.6.1.465":            BaselineCategory6145,
	"1.2.840.10008.6.1.466":            BackgroundEchotexture6151,
	"1.2.840.10008.6.1.467":            Orientation6152,
	"1.2.840.10008.6.1.468":            LesionBoundary6153,
	"1.2.840.10008.6.1.469":            EchoPattern6154,
	"1.2.840.10008.6.1.470":            PosteriorAcousticFeature6155,
	"1.2.840.10008.6.1.471":            Vascularity6157,
	"1.2.840.10008.6.1.472":            CorrelationToOtherFinding6158,
	"1.2.840.10008.6.1.473":            MalignancyType6159,
	"1.2.840.10008.6.1.474":            BreastPrimaryTumorAssessmentFromAJCC6160,
	"1.2.840.10008.6.1.475":            PathologicalRegionalLymphNodeAssessmentForBreast6161,
	"1.2.840.10008.6.1.476":            AssessmentOfMetastasisForBreast6162,
	"1.2.840.10008.6.1.477":            MenstrualCyclePhase6163,
	"1.2.840.10008.6.1.478":            TimeInterval6164,
	"1.2.840.10008.6.1.479":            BreastLinearMeasurement6165,
	"1.2.840.10008.6.1.480":            CADGeometrySecondaryGraphicalRepresentation6166,
	"1.2.840.10008.6.1.481":            DiagnosticImagingReportDocumentTitle7000,
	"1.2.840.10008.6.1.482":            DiagnosticImagingReportHeading7001,
	"1.2.840.10008.6.1.483":            DiagnosticImagingReportElement7002,
	"1.2.840.10008.6.1.484":            DiagnosticImagingReportPurposeOfReference7003,
	"1.2.840.10008.6.1.485":            WaveformPurposeOfReference7004,
	"1.2.840.10008.6.1.486":            ContributingEquipmentPurposeOfReference7005,
	"1.2.840.10008.6.1.487":            SRDocumentPurposeOfReference7006,
	"1.2.840.10008.6.1.488":            SignaturePurpose7007,
	"1.2.840.10008.6.1.489":            MediaImport7008,
	"1.2.840.10008.6.1.490":            KeyObjectSelectionDocumentTitle7010,
	"1.2.840.10008.6.1.491":            RejectedForQualityReason7011,
	"1.2.840.10008.6.1.492":            BestInSet7012,
	"1.2.840.10008.6.1.493":            DocumentTitle7020,
	"1.2.840.10008.6.1.494":            RCSRegistrationMethodType7100,
	"1.2.840.10008.6.1.495":            BrainAtlasFiducial7101,
	"1.2.840.10008.6.1.496":            SegmentationPropertyCategory7150,
	"1.2.840.10008.6.1.497":            SegmentationPropertyType7151,
	"1.2.840.10008.6.1.498":            CardiacStructureSegmentationType7152,
	"1.2.840.10008.6.1.499":            CNSSegmentationType7153,
	"1.2.840.10008.6.1.500":            AbdominalSegmentationType7154,
	"1.2.840.10008.6.1.501":            ThoracicSegmentationType7155,
	"1.2.840.10008.6.1.502":            VascularSegmentationType7156,
	"1.2.840.10008.6.1.503":            DeviceSegmentationType7157,
	"1.2.840.10008.6.1.504":            ArtifactSegmentationType7158,
	"1.2.840.10008.6.1.505":            LesionSegmentationType7159,
	"1.2.840.10008.6.1.506":            PelvicOrganSegmentationType7160,
	"1.2.840.10008.6.1.507":            PhysiologySegmentationType7161,
	"1.2.840.10008.6.1.508":            ReferencedImagePurposeOfReference7201,
	"1.2.840.10008.6.1.509":            SourceImagePurposeOfReference7202,
	"1.2.840.10008.6.1.510":            ImageDerivation7203,
	"1.2.840.10008.6.1.511":            PurposeOfReferenceToAlternateRepresentation7205,
	"1.2.840.10008.6.1.512":            RelatedSeriesPurposeOfReference7210,
	"1.2.840.10008.6.1.513":            MultiFrameSubsetType7250,
	"1.2.840.10008.6.1.514":            PersonRole7450,
	"1.2.840.10008.6.1.515":            FamilyMember7451,
	"1.2.840.10008.6.1.516":            OrganizationalRole7452,
	"1.2.840.10008.6.1.517":            PerformingRole7453,
	"1.2.840.10008.6.1.518":            AnimalTaxonomicRankValue7454,
	"1.2.840.10008.6.1.519":            Sex7455,
	"1.2.840.10008.6.1.520":            AgeUnit7456,
	"1.2.840.10008.6.1.521":            LinearMeasurementUnit7460,
	"1.2.840.10008.6.1.522":            AreaMeasurementUnit7461,
	"1.2.840.10008.6.1.523":            VolumeMeasurementUnit7462,
	"1.2.840.10008.6.1.524":            LinearMeasurement7470,
	"1.2.840.10008.6.1.525":            AreaMeasurement7471,
	"1.2.840.10008.6.1.526":            VolumeMeasurement7472,
	"1.2.840.10008.6.1.527":            GeneralAreaCalculationMethod7473,
	"1.2.840.10008.6.1.528":            GeneralVolumeCalculationMethod7474,
	"1.2.840.10008.6.1.529":            Breed7480,
	"1.2.840.10008.6.1.530":            BreedRegistry7481,
	"1.2.840.10008.6.1.531":            WorkitemDefinition9231,
	"1.2.840.10008.6.1.532":            NonDICOMOutputTypes9232RETIRED,
	"1.2.840.10008.6.1.533":            ProcedureDiscontinuationReason9300,
	"1.2.840.10008.6.1.534":            ScopeOfAccumulation10000,
	"1.2.840.10008.6.1.535":            UIDType10001,
	"1.2.840.10008.6.1.536":            IrradiationEventType10002,
	"1.2.840.10008.6.1.537":            EquipmentPlaneIdentification10003,
	"1.2.840.10008.6.1.538":            FluoroMode10004,
	"1.2.840.10008.6.1.539":            XRayFilterMaterial10006,
	"1.2.840.10008.6.1.540":            XRayFilterType10007,
	"1.2.840.10008.6.1.541":            DoseRelatedDistanceMeasurement10008,
	"1.2.840.10008.6.1.542":            MeasuredCalculated10009,
	"1.2.840.10008.6.1.543":            DoseMeasurementDevice10010,
	"1.2.840.10008.6.1.544":            EffectiveDoseEvaluationMethod10011,
	"1.2.840.10008.6.1.545":            CTAcquisitionType10013,
	"1.2.840.10008.6.1.546":            CTIVContrastImagingTechnique10014,
	"1.2.840.10008.6.1.547":            CTDoseReferenceAuthority10015,
	"1.2.840.10008.6.1.548":            AnodeTargetMaterial10016,
	"1.2.840.10008.6.1.549":            XRayGrid10017,
	"1.2.840.10008.6.1.550":            UltrasoundProtocolType12001,
	"1.2.840.10008.6.1.551":            UltrasoundProtocolStageType12002,
	"1.2.840.10008.6.1.552":            OBGYNDate12003,
	"1.2.840.10008.6.1.553":            FetalBiometryRatio12004,
	"1.2.840.10008.6.1.554":            FetalBiometryMeasurement12005,
	"1.2.840.10008.6.1.555":            FetalLongBonesBiometryMeasurement12006,
	"1.2.840.10008.6.1.556":            FetalCraniumMeasurement12007,
	"1.2.840.10008.6.1.557":            OBGYNAmnioticSacMeasurement12008,
	"1.2.840.10008.6.1.558":            EarlyGestationBiometryMeasurement12009,
	"1.2.840.10008.6.1.559":            UltrasoundPelvisAndUterusMeasurement12011,
	"1.2.840.10008.6.1.560":            OBEquationTable12012,
	"1.2.840.10008.6.1.561":            GestationalAgeEquationTable12013,
	"1.2.840.10008.6.1.562":            OBFetalBodyWeightEquationTable12014,
	"1.2.840.10008.6.1.563":            FetalGrowthEquationTable12015,
	"1.2.840.10008.6.1.564":            EstimatedFetalWeightPercentileEquationTable12016,
	"1.2.840.10008.6.1.565":            GrowthDistributionRank12017,
	"1.2.840.10008.6.1.566":            OBGYNSummary12018,
	"1.2.840.10008.6.1.567":            OBGYNFetusSummary12019,
	"1.2.840.10008.6.1.568":            VascularSummary12101,
	"1.2.840.10008.6.1.569":            TemporalPeriodRelatingToProcedureOrTherapy12102,
	"1.2.840.10008.6.1.570":            VascularUltrasoundAnatomicLocation12103,
	"1.2.840.10008.6.1.571":            ExtracranialArtery12104,
	"1.2.840.10008.6.1.572":            IntracranialCerebralVessel12105,
	"1.2.840.10008.6.1.573":            IntracranialCerebralVesselUnilateral12106,
	"1.2.840.10008.6.1.574":            UpperExtremityArtery12107,
	"1.2.840.10008.6.1.575":            UpperExtremityVein12108,
	"1.2.840.10008.6.1.576":            LowerExtremityArtery12109,
	"1.2.840.10008.6.1.577":            LowerExtremityVein12110,
	"1.2.840.10008.6.1.578":            AbdominopelvicArteryPaired12111,
	"1.2.840.10008.6.1.579":            AbdominopelvicArteryUnpaired12112,
	"1.2.840.10008.6.1.580":            AbdominopelvicVeinPaired12113,
	"1.2.840.10008.6.1.581":            AbdominopelvicVeinUnpaired12114,
	"1.2.840.10008.6.1.582":            RenalVessel12115,
	"1.2.840.10008.6.1.583":            VesselSegmentModifier12116,
	"1.2.840.10008.6.1.584":            VesselBranchModifier12117,
	"1.2.840.10008.6.1.585":            VascularUltrasoundProperty12119,
	"1.2.840.10008.6.1.586":            UltrasoundBloodVelocityMeasurement12120,
	"1.2.840.10008.6.1.587":            VascularIndexRatio12121,
	"1.2.840.10008.6.1.588":            OtherVascularProperty12122,
	"1.2.840.10008.6.1.589":            CarotidRatio12123,
	"1.2.840.10008.6.1.590":            RenalRatio12124,
	"1.2.840.10008.6.1.591":            PelvicVasculatureAnatomicalLocation12140,
	"1.2.840.10008.6.1.592":            FetalVasculatureAnatomicalLocation12141,
	"1.2.840.10008.6.1.593":            EchocardiographyLeftVentricleMeasurement12200,
	"1.2.840.10008.6.1.594":            LeftVentricleLinearMeasurement12201,
	"1.2.840.10008.6.1.595":            LeftVentricleVolumeMeasurement12202,
	"1.2.840.10008.6.1.596":            LeftVentricleOtherMeasurement12203,
	"1.2.840.10008.6.1.597":            EchocardiographyRightVentricleMeasurement12204,
	"1.2.840.10008.6.1.598":            EchocardiographyLeftAtriumMeasurement12205,
	"1.2.840.10008.6.1.599":            EchocardiographyRightAtriumMeasurement12206,
	"1.2.840.10008.6.1.600":            EchocardiographyMitralValveMeasurement12207,
	"1.2.840.10008.6.1.601":            EchocardiographyTricuspidValveMeasurement12208,
	"1.2.840.10008.6.1.602":            EchocardiographyPulmonicValveMeasurement12209,
	"1.2.840.10008.6.1.603":            EchocardiographyPulmonaryArteryMeasurement12210,
	"1.2.840.10008.6.1.604":            EchocardiographyAorticValveMeasurement12211,
	"1.2.840.10008.6.1.605":            EchocardiographyAortaMeasurement12212,
	"1.2.840.10008.6.1.606":            EchocardiographyPulmonaryVeinMeasurement12214,
	"1.2.840.10008.6.1.607":            EchocardiographyVenaCavaMeasurement12215,
	"1.2.840.10008.6.1.608":            EchocardiographyHepaticVeinMeasurement12216,
	"1.2.840.10008.6.1.609":            EchocardiographyCardiacShuntMeasurement12217,
	"1.2.840.10008.6.1.610":            EchocardiographyCongenitalAnomalyMeasurement12218,
	"1.2.840.10008.6.1.611":            PulmonaryVeinModifier12219,
	"1.2.840.10008.6.1.612":            EchocardiographyCommonMeasurement12220,
	"1.2.840.10008.6.1.613":            FlowDirection12221,
	"1.2.840.10008.6.1.614":            OrificeFlowProperty12222,
	"1.2.840.10008.6.1.615":            EchocardiographyStrokeVolumeOrigin12223,
	"1.2.840.10008.6.1.616":            UltrasoundImageMode12224,
	"1.2.840.10008.6.1.617":            EchocardiographyImageView12226,
	"1.2.840.10008.6.1.618":            EchocardiographyMeasurementMethod12227,
	"1.2.840.10008.6.1.619":            EchocardiographyVolumeMethod12228,
	"1.2.840.10008.6.1.620":            EchocardiographyAreaMethod12229,
	"1.2.840.10008.6.1.621":            GradientMethod12230,
	"1.2.840.10008.6.1.622":            VolumeFlowMethod12231,
	"1.2.840.10008.6.1.623":            MyocardiumMassMethod12232,
	"1.2.840.10008.6.1.624":            CardiacPhase12233,
	"1.2.840.10008.6.1.625":            RespirationState12234,
	"1.2.840.10008.6.1.626":            MitralValveAnatomicSite12235,
	"1.2.840.10008.6.1.627":            EchocardiographyAnatomicSite12236,
	"1.2.840.10008.6.1.628":            EchocardiographyAnatomicSiteModifier12237,
	"1.2.840.10008.6.1.629":            WallMotionScoringScheme12238,
	"1.2.840.10008.6.1.630":            CardiacOutputProperty12239,
	"1.2.840.10008.6.1.631":            LeftVentricleAreaMeasurement12240,
	"1.2.840.10008.6.1.632":            TricuspidValveFindingSite12241,
	"1.2.840.10008.6.1.633":            AorticValveFindingSite12242,
	"1.2.840.10008.6.1.634":            LeftVentricleFindingSite12243,
	"1.2.840.10008.6.1.635":            CongenitalFindingSite12244,
	"1.2.840.10008.6.1.636":            SurfaceProcessingAlgorithmFamily7162,
	"1.2.840.10008.6.1.637":            StressTestProcedurePhase3207,
	"1.2.840.10008.6.1.638":            Stage3778,
	"1.2.840.10008.6.1.735":            SMLSizeDescriptor252,
	"1.2.840.10008.6.1.736":            MajorCoronaryArtery3016,
	"1.2.840.10008.6.1.737":            RadioactivityUnit3083,
	"1.2.840.10008.6.1.738":            RestStressState3102,
	"1.2.840.10008.6.1.739":            PETCardiologyProtocol3106,
	"1.2.840.10008.6.1.740":            PETCardiologyRadiopharmaceutical3107,
	"1.2.840.10008.6.1.741":            NMPETProcedure3108,
	"1.2.840.10008.6.1.742":            NuclearCardiologyProtocol3110,
	"1.2.840.10008.6.1.743":            NuclearCardiologyRadiopharmaceutical3111,
	"1.2.840.10008.6.1.744":            AttenuationCorrection3112,
	"1.2.840.10008.6.1.745":            PerfusionDefectType3113,
	"1.2.840.10008.6.1.746":            StudyQuality3114,
	"1.2.840.10008.6.1.747":            StressImagingQualityIssue3115,
	"1.2.840.10008.6.1.748":            NMExtracardiacFinding3116,
	"1.2.840.10008.6.1.749":            AttenuationCorrectionMethod3117,
	"1.2.840.10008.6.1.750":            LevelOfRisk3118,
	"1.2.840.10008.6.1.751":            LVFunction3119,
	"1.2.840.10008.6.1.752":            PerfusionFinding3120,
	"1.2.840.10008.6.1.753":            PerfusionMorphology3121,
	"1.2.840.10008.6.1.754":            VentricularEnlargement3122,
	"1.2.840.10008.6.1.755":            StressTestProcedure3200,
	"1.2.840.10008.6.1.756":            IndicationsForStressTest3201,
	"1.2.840.10008.6.1.757":            ChestPain3202,
	"1.2.840.10008.6.1.758":            ExerciserDevice3203,
	"1.2.840.10008.6.1.759":            StressAgent3204,
	"1.2.840.10008.6.1.760":            IndicationsForPharmacologicalStressTest3205,
	"1.2.840.10008.6.1.761":            NonInvasiveCardiacImagingProcedure3206,
	"1.2.840.10008.6.1.763":            ExerciseECGSummaryCode3208,
	"1.2.840.10008.6.1.764":            StressImagingSummaryCode3209,
	"1.2.840.10008.6.1.765":            SpeedOfResponse3210,
	"1.2.840.10008.6.1.766":            BPResponse3211,
	"1.2.840.10008.6.1.767":            TreadmillSpeed3212,
	"1.2.840.10008.6.1.768":            StressHemodynamicFinding3213,
	"1.2.840.10008.6.1.769":            PerfusionFindingMethod3215,
	"1.2.840.10008.6.1.770":            ComparisonFinding3217,
	"1.2.840.10008.6.1.771":            StressSymptom3220,
	"1.2.840.10008.6.1.772":            StressTestTerminationReason3221,
	"1.2.840.10008.6.1.773":            QTcMeasurement3227,
	"1.2.840.10008.6.1.774":            ECGTimingMeasurement3228,
	"1.2.840.10008.6.1.775":            ECGAxisMeasurement3229,
	"1.2.840.10008.6.1.776":            ECGFinding3230,
	"1.2.840.10008.6.1.777":            STSegmentFinding3231,
	"1.2.840.10008.6.1.778":            STSegmentLocation3232,
	"1.2.840.10008.6.1.779":            STSegmentMorphology3233,
	"1.2.840.10008.6.1.780":            EctopicBeatMorphology3234,
	"1.2.840.10008.6.1.781":            PerfusionComparisonFinding3235,
	"1.2.840.10008.6.1.782":            ToleranceComparisonFinding3236,
	"1.2.840.10008.6.1.783":            WallMotionComparisonFinding3237,
	"1.2.840.10008.6.1.784":            StressScoringScale3238,
	"1.2.840.10008.6.1.785":            PerceivedExertionScale3239,
	"1.2.840.10008.6.1.786":            VentricleIdentification3463,
	"1.2.840.10008.6.1.787":            ColonOverallAssessment6200,
	"1.2.840.10008.6.1.788":            ColonFindingOrFeature6201,
	"1.2.840.10008.6.1.789":            ColonFindingOrFeatureModifier6202,
	"1.2.840.10008.6.1.790":            ColonNonLesionObjectType6203,
	"1.2.840.10008.6.1.791":            AnatomicNonColonFinding6204,
	"1.2.840.10008.6.1.792":            ClockfaceLocationForColon6205,
	"1.2.840.10008.6.1.793":            RecumbentPatientOrientationForColon6206,
	"1.2.840.10008.6.1.794":            ColonQuantitativeTemporalDifferenceType6207,
	"1.2.840.10008.6.1.795":            ColonTypesOfQualityControlStandard6208,
	"1.2.840.10008.6.1.796":            ColonMorphologyDescriptor6209,
	"1.2.840.10008.6.1.797":            LocationInIntestinalTract6210,
	"1.2.840.10008.6.1.798":            ColonCADMaterialDescription6211,
	"1.2.840.10008.6.1.799":            CalculatedValueForColonFinding6212,
	"1.2.840.10008.6.1.800":            OphthalmicHorizontalDirection4214,
	"1.2.840.10008.6.1.801":            OphthalmicVerticalDirection4215,
	"1.2.840.10008.6.1.802":            OphthalmicVisualAcuityType4216,
	"1.2.840.10008.6.1.803":            ArterialPulseWaveform3004,
	"1.2.840.10008.6.1.804":            RespirationWaveform3005,
	"1.2.840.10008.6.1.805":            UltrasoundContrastBolusAgent12030,
	"1.2.840.10008.6.1.806":            ProtocolIntervalEvent12031,
	"1.2.840.10008.6.1.807":            TransducerScanPattern12032,
	"1.2.840.10008.6.1.808":            UltrasoundTransducerGeometry12033,
	"1.2.840.10008.6.1.809":            UltrasoundTransducerBeamSteering12034,
	"1.2.840.10008.6.1.810":            UltrasoundTransducerApplication12035,
	"1.2.840.10008.6.1.811":            InstanceAvailabilityStatus50,
	"1.2.840.10008.6.1.812":            ModalityPPSDiscontinuationReason9301,
	"1.2.840.10008.6.1.813":            MediaImportPPSDiscontinuationReason9302,
	"1.2.840.10008.6.1.814":            DXAnatomyImagedForAnimal7482,
	"1.2.840.10008.6.1.815":            CommonAnatomicRegionsForAnimal7483,
	"1.2.840.10008.6.1.816":            DXViewForAnimal7484,
	"1.2.840.10008.6.1.817":            InstitutionalDepartmentUnitService7030,
	"1.2.840.10008.6.1.818":            PurposeOfReferenceToPredecessorReport7009,
	"1.2.840.10008.6.1.819":            VisualFixationQualityDuringAcquisition4220,
	"1.2.840.10008.6.1.820":            VisualFixationQualityProblem4221,
	"1.2.840.10008.6.1.821":            OphthalmicMacularGridProblem4222,
	"1.2.840.10008.6.1.822":            Organization5002,
	"1.2.840.10008.6.1.823":            MixedBreed7486,
	"1.2.840.10008.6.1.824":            BroselowLutenPediatricSizeCategory7040,
	"1.2.840.10008.6.1.825":            CMDCTECCCalciumScoringPatientSizeCategory7042,
	"1.2.840.10008.6.1.826":            CardiacUltrasoundReportTitle12245,
	"1.2.840.10008.6.1.827":            CardiacUltrasoundIndicationForStudy12246,
	"1.2.840.10008.6.1.828":            PediatricFetalAndCongenitalCardiacSurgicalIntervention12247,
	"1.2.840.10008.6.1.829":            CardiacUltrasoundSummaryCode12248,
	"1.2.840.10008.6.1.830":            CardiacUltrasoundFetalSummaryCode12249,
	"1.2.840.10008.6.1.831":            CardiacUltrasoundCommonLinearMeasurement12250,
	"1.2.840.10008.6.1.832":            CardiacUltrasoundLinearValveMeasurement12251,
	"1.2.840.10008.6.1.833":            CardiacUltrasoundCardiacFunction12252,
	"1.2.840.10008.6.1.834":            CardiacUltrasoundAreaMeasurement12253,
	"1.2.840.10008.6.1.835":            CardiacUltrasoundHemodynamicMeasurement12254,
	"1.2.840.10008.6.1.836":            CardiacUltrasoundMyocardiumMeasurement12255,
	"1.2.840.10008.6.1.838":            CardiacUltrasoundLeftVentricleMeasurement12257,
	"1.2.840.10008.6.1.839":            CardiacUltrasoundRightVentricleMeasurement12258,
	"1.2.840.10008.6.1.840":            CardiacUltrasoundVentriclesMeasurement12259,
	"1.2.840.10008.6.1.841":            CardiacUltrasoundPulmonaryArteryMeasurement12260,
	"1.2.840.10008.6.1.842":            CardiacUltrasoundPulmonaryVein12261,
	"1.2.840.10008.6.1.843":            CardiacUltrasoundPulmonaryValveMeasurement12262,
	"1.2.840.10008.6.1.844":            CardiacUltrasoundVenousReturnPulmonaryMeasurement12263,
	"1.2.840.10008.6.1.845":            CardiacUltrasoundVenousReturnSystemicMeasurement12264,
	"1.2.840.10008.6.1.846":            CardiacUltrasoundAtriaAndAtrialSeptumMeasurement12265,
	"1.2.840.10008.6.1.847":            CardiacUltrasoundMitralValveMeasurement12266,
	"1.2.840.10008.6.1.848":            CardiacUltrasoundTricuspidValveMeasurement12267,
	"1.2.840.10008.6.1.849":            CardiacUltrasoundAtrioventricularValveMeasurement12268,
	"1.2.840.10008.6.1.850":            CardiacUltrasoundInterventricularSeptumMeasurement12269,
	"1.2.840.10008.6.1.851":            CardiacUltrasoundAorticValveMeasurement12270,
	"1.2.840.10008.6.1.852":            CardiacUltrasoundOutflowTractMeasurement12271,
	"1.2.840.10008.6.1.853":            CardiacUltrasoundSemilunarValveAnnulateAndSinusMeasurement12272,
	"1.2.840.10008.6.1.854":            CardiacUltrasoundAorticSinotubularJunctionMeasurement12273,
	"1.2.840.10008.6.1.855":            CardiacUltrasoundAortaMeasurement12274,
	"1.2.840.10008.6.1.856":            CardiacUltrasoundCoronaryArteryMeasurement12275,
	"1.2.840.10008.6.1.857":            CardiacUltrasoundAortoPulmonaryConnectionMeasurement12276,
	"1.2.840.10008.6.1.858":            CardiacUltrasoundPericardiumAndPleuraMeasurement12277,
	"1.2.840.10008.6.1.859":            CardiacUltrasoundFetalGeneralMeasurement12279,
	"1.2.840.10008.6.1.860":            CardiacUltrasoundTargetSite12280,
	"1.2.840.10008.6.1.861":            CardiacUltrasoundTargetSiteModifier12281,
	"1.2.840.10008.6.1.862":            CardiacUltrasoundVenousReturnSystemicFindingSite12282,
	"1.2.840.10008.6.1.863":            CardiacUltrasoundVenousReturnPulmonaryFindingSite12283,
	"1.2.840.10008.6.1.864":            CardiacUltrasoundAtriaAndAtrialSeptumFindingSite12284,
	"1.2.840.10008.6.1.865":            CardiacUltrasoundAtrioventricularValveFindingSite12285,
	"1.2.840.10008.6.1.866":            CardiacUltrasoundInterventricularSeptumFindingSite12286,
	"1.2.840.10008.6.1.867":            CardiacUltrasoundVentricleFindingSite12287,
	"1.2.840.10008.6.1.868":            CardiacUltrasoundOutflowTractFindingSite12288,
	"1.2.840.10008.6.1.869":            CardiacUltrasoundSemilunarValveAnnulusAndSinusFindingSite12289,
	"1.2.840.10008.6.1.870":            CardiacUltrasoundPulmonaryArteryFindingSite12290,
	"1.2.840.10008.6.1.871":            CardiacUltrasoundAortaFindingSite12291,
	"1.2.840.10008.6.1.872":            CardiacUltrasoundCoronaryArteryFindingSite12292,
	"1.2.840.10008.6.1.873":            CardiacUltrasoundAortopulmonaryConnectionFindingSite12293,
	"1.2.840.10008.6.1.874":            CardiacUltrasoundPericardiumAndPleuraFindingSite12294,
	"1.2.840.10008.6.1.876":            OphthalmicUltrasoundAxialMeasurementsType4230,
	"1.2.840.10008.6.1.877":            LensStatus4231,
	"1.2.840.10008.6.1.878":            VitreousStatus4232,
	"1.2.840.10008.6.1.879":            OphthalmicAxialLengthMeasurementsSegmentName4233,
	"1.2.840.10008.6.1.880":            RefractiveSurgeryType4234,
	"1.2.840.10008.6.1.881":            KeratometryDescriptor4235,
	"1.2.840.10008.6.1.882":            IOLCalculationFormula4236,
	"1.2.840.10008.6.1.883":            LensConstantType4237,
	"1.2.840.10008.6.1.884":            RefractiveErrorType4238,
	"1.2.840.10008.6.1.885":            AnteriorChamberDepthDefinition4239,
	"1.2.840.10008.6.1.886":            OphthalmicMeasurementOrCalculationDataSource4240,
	"1.2.840.10008.6.1.887":            OphthalmicAxialLengthSelectionMethod4241,
	"1.2.840.10008.6.1.889":            OphthalmicQualityMetricType4243,
	"1.2.840.10008.6.1.890":            OphthalmicAgentConcentrationUnit4244,
	"1.2.840.10008.6.1.891":            FunctionalConditionPresentDuringAcquisition91,
	"1.2.840.10008.6.1.892":            JointPositionDuringAcquisition92,
	"1.2.840.10008.6.1.893":            JointPositioningMethod93,
	"1.2.840.10008.6.1.894":            PhysicalForceAppliedDuringAcquisition94,
	"1.2.840.10008.6.1.895":            ECGControlNumericVariable3690,
	"1.2.840.10008.6.1.896":            ECGControlTextVariable3691,
	"1.2.840.10008.6.1.897":            WholeSlideMicroscopyImageReferencedImagePurposeOfReference8120,
	"1.2.840.10008.6.1.898":            MicroscopyLensType8121,
	"1.2.840.10008.6.1.899":            MicroscopyIlluminatorAndSensorColor8122,
	"1.2.840.10008.6.1.900":            MicroscopyIlluminationMethod8123,
	"1.2.840.10008.6.1.901":            MicroscopyFilter8124,
	"1.2.840.10008.6.1.902":            MicroscopyIlluminatorType8125,
	"1.2.840.10008.6.1.903":            AuditEventID400,
	"1.2.840.10008.6.1.904":            AuditEventTypeCode401,
	"1.2.840.10008.6.1.905":            AuditActiveParticipantRoleIDCode402,
	"1.2.840.10008.6.1.906":            SecurityAlertTypeCode403,
	"1.2.840.10008.6.1.907":            AuditParticipantObjectIDTypeCode404,
	"1.2.840.10008.6.1.908":            MediaTypeCode405,
	"1.2.840.10008.6.1.909":            VisualFieldStaticPerimetryTestPattern4250,
	"1.2.840.10008.6.1.910":            VisualFieldStaticPerimetryTestStrategy4251,
	"1.2.840.10008.6.1.911":            VisualFieldStaticPerimetryScreeningTestMode4252,
	"1.2.840.10008.6.1.912":            VisualFieldStaticPerimetryFixationStrategy4253,
	"1.2.840.10008.6.1.913":            VisualFieldStaticPerimetryTestAnalysisResult4254,
	"1.2.840.10008.6.1.914":            VisualFieldIlluminationColor4255,
	"1.2.840.10008.6.1.915":            VisualFieldProcedureModifier4256,
	"1.2.840.10008.6.1.916":            VisualFieldGlobalIndexName4257,
	"1.2.840.10008.6.1.917":            AbstractMultiDimensionalImageModelComponentSemantic7180,
	"1.2.840.10008.6.1.918":            AbstractMultiDimensionalImageModelComponentUnit7181,
	"1.2.840.10008.6.1.919":            AbstractMultiDimensionalImageModelDimensionSemantic7182,
	"1.2.840.10008.6.1.920":            AbstractMultiDimensionalImageModelDimensionUnit7183,
	"1.2.840.10008.6.1.921":            AbstractMultiDimensionalImageModelAxisDirection7184,
	"1.2.840.10008.6.1.922":            AbstractMultiDimensionalImageModelAxisOrientation7185,
	"1.2.840.10008.6.1.923":            AbstractMultiDimensionalImageModelQualitativeDimensionSampleSemantic7186,
	"1.2.840.10008.6.1.924":            PlanningMethod7320,
	"1.2.840.10008.6.1.925":            DeIdentificationMethod7050,
	"1.2.840.10008.6.1.926":            MeasurementOrientation12118,
	"1.2.840.10008.6.1.927":            ECGGlobalWaveformDuration3689,
	"1.2.840.10008.6.1.930":            ICD3692,
	"1.2.840.10008.6.1.931":            RadiotherapyGeneralWorkitemDefinition9241,
	"1.2.840.10008.6.1.932":            RadiotherapyAcquisitionWorkitemDefinition9242,
	"1.2.840.10008.6.1.933":            RadiotherapyRegistrationWorkitemDefinition9243,
	"1.2.840.10008.6.1.934":            ContrastBolusSubstance3850,
	"1.2.840.10008.6.1.935":            LabelType10022,
	"1.2.840.10008.6.1.936":            OphthalmicMappingUnitForRealWorldValueMapping4260,
	"1.2.840.10008.6.1.937":            OphthalmicMappingAcquisitionMethod4261,
	"1.2.840.10008.6.1.938":            RetinalThicknessDefinition4262,
	"1.2.840.10008.6.1.939":            OphthalmicThicknessMapValueType4263,
	"1.2.840.10008.6.1.940":            OphthalmicMapPurposeOfReference4264,
	"1.2.840.10008.6.1.941":            OphthalmicThicknessDeviationCategory4265,
	"1.2.840.10008.6.1.942":            OphthalmicAnatomicStructureReferencePoint4266,
	"1.2.840.10008.6.1.943":            CardiacSynchronizationTechnique3104,
	"1.2.840.10008.6.1.944":            StainingProtocol8130,
	"1.2.840.10008.6.1.947":            SizeSpecificDoseEstimationMethodForCT10023,
	"1.2.840.10008.6.1.948":            PathologyImagingProtocol8131,
	"1.2.840.10008.6.1.949":            MagnificationSelection8132,
	"1.2.840.10008.6.1.950":            TissueSelection8133,
	"1.2.840.10008.6.1.951":            GeneralRegionOfInterestMeasurementModifier7464,
	"1.2.840.10008.6.1.952":            MeasurementDerivedFromMultipleROIMeasurements7465,
	"1.2.840.10008.6.1.953":            SurfaceScanAcquisitionType8201,
	"1.2.840.10008.6.1.954":            SurfaceScanModeType8202,
	"1.2.840.10008.6.1.956":            SurfaceScanRegistrationMethodType8203,
	"1.2.840.10008.6.1.957":            BasicCardiacView27,
	"1.2.840.10008.6.1.958":            CTReconstructionAlgorithm10033,
	"1.2.840.10008.6.1.959":            DetectorType10030,
	"1.2.840.10008.6.1.960":            CRDRMechanicalConfiguration10031,
	"1.2.840.10008.6.1.961":            ProjectionXRayAcquisitionDeviceType10032,
	"1.2.840.10008.6.1.962":            AbstractSegmentationType7165,
	"1.2.840.10008.6.1.963":            CommonTissueSegmentationType7166,
	"1.2.840.10008.6.1.964":            PeripheralNervousSystemSegmentationType7167,
	"1.2.840.10008.6.1.965":            CornealTopographyMappingUnitForRealWorldValueMapping4267,
	"1.2.840.10008.6.1.966":            CornealTopographyMapValueType4268,
	"1.2.840.10008.6.1.967":            BrainStructureForVolumetricMeasurement7140,
	"1.2.840.10008.6.1.968":            RTDoseDerivation7220,
	"1.2.840.10008.6.1.969":            RTDosePurposeOfReference7221,
	"1.2.840.10008.6.1.970":            SpectroscopyPurposeOfReference7215,
	"1.2.840.10008.6.1.971":            ScheduledProcessingParameterConceptCodesForRTTreatment9250,
	"1.2.840.10008.6.1.972":            RadiopharmaceuticalOrganDoseReferenceAuthority10040,
	"1.2.840.10008.6.1.973":            SourceOfRadioisotopeActivityInformation10041,
	"1.2.840.10008.6.1.975":            IntravenousExtravasationSymptom10043,
	"1.2.840.10008.6.1.976":            RadiosensitiveOrgan10044,
	"1.2.840.10008.6.1.977":            RadiopharmaceuticalPatientState10045,
	"1.2.840.10008.6.1.978":            GFRMeasurement10046,
	"1.2.840.10008.6.1.979":            GFRMeasurementMethod10047,
	"1.2.840.10008.6.1.980":            VisualEvaluationMethod8300,
	"1.2.840.10008.6.1.981":            TestPatternCode8301,
	"1.2.840.10008.6.1.982":            MeasurementPatternCode8302,
	"1.2.840.10008.6.1.983":            DisplayDeviceType8303,
	"1.2.840.10008.6.1.984":            SUVUnit85,
	"1.2.840.10008.6.1.985":            T1MeasurementMethod4100,
	"1.2.840.10008.6.1.986":            TracerKineticModel4101,
	"1.2.840.10008.6.1.987":            PerfusionMeasurementMethod4102,
	"1.2.840.10008.6.1.988":            ArterialInputFunctionMeasurementMethod4103,
	"1.2.840.10008.6.1.989":            BolusArrivalTimeDerivationMethod4104,
	"1.2.840.10008.6.1.990":            PerfusionAnalysisMethod4105,
	"1.2.840.10008.6.1.991":            QuantitativeMethodUsedForPerfusionAndTracerKineticModel4106,
	"1.2.840.10008.6.1.992":            TracerKineticModelParameter4107,
	"1.2.840.10008.6.1.993":            PerfusionModelParameter4108,
	"1.2.840.10008.6.1.994":            ModelIndependentDynamicContrastAnalysisParameter4109,
	"1.2.840.10008.6.1.995":            TracerKineticModelingCovariate4110,
	"1.2.840.10008.6.1.996":            ContrastCharacteristic4111,
	"1.2.840.10008.6.1.997":            MeasurementReportDocumentTitle7021,
	"1.2.840.10008.6.1.998":            QuantitativeDiagnosticImagingProcedure100,
	"1.2.840.10008.6.1.999":            PETRegionOfInterestMeasurement7466,
	"1.2.840.10008.6.1.1000":           GrayLevelCoOccurrenceMatrixMeasurement7467,
	"1.2.840.10008.6.1.1001":           TextureMeasurement7468,
	"1.2.840.10008.6.1.1002":           TimePointType6146,
	"1.2.840.10008.6.1.1003":           GenericIntensityAndSizeMeasurement7469,
	"1.2.840.10008.6.1.1004":           ResponseCriteria6147,
	"1.2.840.10008.6.1.1005":           FetalBiometryAnatomicSite12020,
	"1.2.840.10008.6.1.1006":           FetalLongBoneAnatomicSite12021,
	"1.2.840.10008.6.1.1007":           FetalCraniumAnatomicSite12022,
	"1.2.840.10008.6.1.1008":           PelvisAndUterusAnatomicSite12023,
	"1.2.840.10008.6.1.1009":           ParametricMapDerivationImagePurposeOfReference7222,
	"1.2.840.10008.6.1.1010":           PhysicalQuantityDescriptor9000,
	"1.2.840.10008.6.1.1011":           LymphNodeAnatomicSite7600,
	"1.2.840.10008.6.1.1012":           HeadAndNeckCancerAnatomicSite7601,
	"1.2.840.10008.6.1.1013":           FiberTractInBrainstem7701,
	"1.2.840.10008.6.1.1014":           ProjectionAndThalamicFiber7702,
	"1.2.840.10008.6.1.1015":           AssociationFiber7703,
	"1.2.840.10008.6.1.1016":           LimbicSystemTract7704,
	"1.2.840.10008.6.1.1017":           CommissuralFiber7705,
	"1.2.840.10008.6.1.1018":           CranialNerve7706,
	"1.2.840.10008.6.1.1019":           SpinalCordFiber7707,
	"1.2.840.10008.6.1.1020":           TractographyAnatomicSite7710,
	"1.2.840.10008.6.1.1021":           PrimaryAnatomicStructureForIntraOralRadiographySupernumeraryDentitionDesignationOfTeeth4025,
	"1.2.840.10008.6.1.1022":           PrimaryAnatomicStructureForIntraOralAndCraniofacialRadiographyTeeth4026,
	"1.2.840.10008.6.1.1023":           IEC61217DevicePositionParameter9401,
	"1.2.840.10008.6.1.1024":           IEC61217GantryPositionParameter9402,
	"1.2.840.10008.6.1.1025":           IEC61217PatientSupportPositionParameter9403,
	"1.2.840.10008.6.1.1026":           ActionableFindingClassification7035,
	"1.2.840.10008.6.1.1027":           ImageQualityAssessment7036,
	"1.2.840.10008.6.1.1028":           SummaryRadiationExposureQuantity10050,
	"1.2.840.10008.6.1.1029":           WideFieldOphthalmicPhotographyTransformationMethod4245,
	"1.2.840.10008.6.1.1030":           PETUnit84,
	"1.2.840.10008.6.1.1031":           ImplantMaterial7300,
	"1.2.840.10008.6.1.1032":           InterventionType7301,
	"1.2.840.10008.6.1.1033":           ImplantTemplateViewOrientation7302,
	"1.2.840.10008.6.1.1034":           ImplantTemplateModifiedViewOrientation7303,
	"1.2.840.10008.6.1.1035":           ImplantTargetAnatomy7304,
	"1.2.840.10008.6.1.1036":           ImplantPlanningLandmark7305,
	"1.2.840.10008.6.1.1037":           HumanHipImplantPlanningLandmark7306,
	"1.2.840.10008.6.1.1038":           ImplantComponentType7307,
	"1.2.840.10008.6.1.1039":           HumanHipImplantComponentType7308,
	"1.2.840.10008.6.1.1040":           HumanTraumaImplantComponentType7309,
	"1.2.840.10008.6.1.1041":           ImplantFixationMethod7310,
	"1.2.840.10008.6.1.1042":           DeviceParticipatingRole7445,
	"1.2.840.10008.6.1.1043":           ContainerType8101,
	"1.2.840.10008.6.1.1044":           ContainerComponentType8102,
	"1.2.840.10008.6.1.1045":           AnatomicPathologySpecimenType8103,
	"1.2.840.10008.6.1.1046":           BreastTissueSpecimenType8104,
	"1.2.840.10008.6.1.1047":           SpecimenCollectionProcedure8109,
	"1.2.840.10008.6.1.1048":           SpecimenSamplingProcedure8110,
	"1.2.840.10008.6.1.1049":           SpecimenPreparationProcedure8111,
	"1.2.840.10008.6.1.1050":           SpecimenStain8112,
	"1.2.840.10008.6.1.1051":           SpecimenPreparationStep8113,
	"1.2.840.10008.6.1.1052":           SpecimenFixative8114,
	"1.2.840.10008.6.1.1053":           SpecimenEmbeddingMedia8115,
	"1.2.840.10008.6.1.1054":           SourceOfProjectionXRayDoseInformation10020,
	"1.2.840.10008.6.1.1055":           SourceOfCTDoseInformation10021,
	"1.2.840.10008.6.1.1056":           RadiationDoseReferencePoint10025,
	"1.2.840.10008.6.1.1057":           VolumetricViewDescription501,
	"1.2.840.10008.6.1.1058":           VolumetricViewModifier502,
	"1.2.840.10008.6.1.1059":           DiffusionAcquisitionValueType7260,
	"1.2.840.10008.6.1.1060":           DiffusionModelValueType7261,
	"1.2.840.10008.6.1.1061":           DiffusionTractographyAlgorithmFamily7262,
	"1.2.840.10008.6.1.1062":           DiffusionTractographyMeasurementType7263,
	"1.2.840.10008.6.1.1063":           ResearchAnimalSourceRegistry7490,
	"1.2.840.10008.6.1.1064":           YesNoOnly231,
	"1.2.840.10008.6.1.1065":           BiosafetyLevel601,
	"1.2.840.10008.6.1.1066":           BiosafetyControlReason602,
	"1.2.840.10008.6.1.1067":           SexMaleFemaleOrBoth7457,
	"1.2.840.10008.6.1.1068":           AnimalRoomType603,
	"1.2.840.10008.6.1.1069":           DeviceReuse604,
	"1.2.840.10008.6.1.1070":           AnimalBeddingMaterial605,
	"1.2.840.10008.6.1.1071":           AnimalShelterType606,
	"1.2.840.10008.6.1.1072":           AnimalFeedType607,
	"1.2.840.10008.6.1.1073":           AnimalFeedSource608,
	"1.2.840.10008.6.1.1074":           AnimalFeedingMethod609,
	"1.2.840.10008.6.1.1075":           WaterType610,
	"1.2.840.10008.6.1.1076":           AnesthesiaCategoryCodeTypeForSmallAnimalAnesthesia611,
	"1.2.840.10008.6.1.1077":           AnesthesiaCategoryCodeTypeFromAnesthesiaQualityInitiative612,
	"1.2.840.10008.6.1.1078":           AnesthesiaInductionCodeTypeForSmallAnimalAnesthesia613,
	"1.2.840.10008.6.1.1079":           AnesthesiaInductionCodeTypeFromAnesthesiaQualityInitiative614,
	"1.2.840.10008.6.1.1080":           AnesthesiaMaintenanceCodeTypeForSmallAnimalAnesthesia615,
	"1.2.840.10008.6.1.1081":           AnesthesiaMaintenanceCodeTypeFromAnesthesiaQualityInitiative616,
	"1.2.840.10008.6.1.1082":           AirwayManagementMethodCodeTypeForSmallAnimalAnesthesia617,
	"1.2.840.10008.6.1.1083":           AirwayManagementMethodCodeTypeFromAnesthesiaQualityInitiative618,
	"1.2.840.10008.6.1.1084":           AirwayManagementSubMethodCodeTypeForSmallAnimalAnesthesia619,
	"1.2.840.10008.6.1.1085":           AirwayManagementSubMethodCodeTypeFromAnesthesiaQualityInitiative620,
	"1.2.840.10008.6.1.1086":           MedicationTypeForSmallAnimalAnesthesia621,
	"1.2.840.10008.6.1.1087":           MedicationTypeCodeTypeFromAnesthesiaQualityInitiative622,
	"1.2.840.10008.6.1.1088":           MedicationForSmallAnimalAnesthesia623,
	"1.2.840.10008.6.1.1089":           InhalationalAnesthesiaAgentForSmallAnimalAnesthesia624,
	"1.2.840.10008.6.1.1090":           InjectableAnesthesiaAgentForSmallAnimalAnesthesia625,
	"1.2.840.10008.6.1.1091":           PremedicationAgentForSmallAnimalAnesthesia626,
	"1.2.840.10008.6.1.1092":           NeuromuscularBlockingAgentForSmallAnimalAnesthesia627,
	"1.2.840.10008.6.1.1093":           AncillaryMedicationsForSmallAnimalAnesthesia628,
	"1.2.840.10008.6.1.1094":           CarrierGasesForSmallAnimalAnesthesia629,
	"1.2.840.10008.6.1.1095":           LocalAnestheticsForSmallAnimalAnesthesia630,
	"1.2.840.10008.6.1.1096":           ProcedurePhaseRequiringAnesthesia631,
	"1.2.840.10008.6.1.1097":           SurgicalProcedurePhaseRequiringAnesthesia632,
	"1.2.840.10008.6.1.1098":           PhaseOfImagingProcedureRequiringAnesthesia633RETIRED,
	"1.2.840.10008.6.1.1099":           AnimalHandlingPhase634,
	"1.2.840.10008.6.1.1100":           HeatingMethod635,
	"1.2.840.10008.6.1.1101":           TemperatureSensorDeviceComponentTypeForSmallAnimalProcedure636,
	"1.2.840.10008.6.1.1102":           ExogenousSubstanceType637,
	"1.2.840.10008.6.1.1103":           ExogenousSubstance638,
	"1.2.840.10008.6.1.1104":           TumorGraftHistologicType639,
	"1.2.840.10008.6.1.1105":           Fibril640,
	"1.2.840.10008.6.1.1106":           Virus641,
	"1.2.840.10008.6.1.1107":           Cytokine642,
	"1.2.840.10008.6.1.1108":           Toxin643,
	"1.2.840.10008.6.1.1109":           ExogenousSubstanceAdministrationSite644,
	"1.2.840.10008.6.1.1110":           ExogenousSubstanceOriginTissue645,
	"1.2.840.10008.6.1.1111":           PreclinicalSmallAnimalImagingProcedure646,
	"1.2.840.10008.6.1.1112":           PositionReferenceIndicatorForFrameOfReference647,
	"1.2.840.10008.6.1.1113":           PresentAbsentOnly241,
	"1.2.840.10008.6.1.1114":           WaterEquivalentDiameterMethod10024,
	"1.2.840.10008.6.1.1115":           RadiotherapyPurposeOfReference7022,
	"1.2.840.10008.6.1.1116":           ContentAssessmentType701,
	"1.2.840.10008.6.1.1117":           RTContentAssessmentType702,
	"1.2.840.10008.6.1.1118":           AssessmentBasis703,
	"1.2.840.10008.6.1.1119":           ReaderSpecialty7449,
	"1.2.840.10008.6.1.1120":           RequestedReportType9233,
	"1.2.840.10008.6.1.1121":           CTTransversePlaneReferenceBasis1000,
	"1.2.840.10008.6.1.1122":           AnatomicalReferenceBasis1001,
	"1.2.840.10008.6.1.1123":           AnatomicalReferenceBasisHead1002,
	"1.2.840.10008.6.1.1124":           AnatomicalReferenceBasisSpine1003,
	"1.2.840.10008.6.1.1125":           AnatomicalReferenceBasisChest1004,
	"1.2.840.10008.6.1.1126":           AnatomicalReferenceBasisAbdomenPelvis1005,
	"1.2.840.10008.6.1.1127":           AnatomicalReferenceBasisExtremity1006,
	"1.2.840.10008.6.1.1128":           ReferenceGeometryPlane1010,
	"1.2.840.10008.6.1.1129":           ReferenceGeometryPoint1011,
	"1.2.840.10008.6.1.1130":           PatientAlignmentMethod1015,
	"1.2.840.10008.6.1.1131":           ContraindicationsForCTImaging1200,
	"1.2.840.10008.6.1.1132":           FiducialCategory7110,
	"1.2.840.10008.6.1.1133":           Fiducial7111,
	"1.2.840.10008.6.1.1134":           NonImageSourceInstancePurposeOfReference7013,
	"1.2.840.10008.6.1.1135":           RTProcessOutput7023,
	"1.2.840.10008.6.1.1136":           RTProcessInput7024,
	"1.2.840.10008.6.1.1137":           RTProcessInputUsed7025,
	"1.2.840.10008.6.1.1138":           ProstateAnatomy6300,
	"1.2.840.10008.6.1.1139":           ProstateSectorAnatomyFromPIRADSV26301,
	"1.2.840.10008.6.1.1140":           ProstateSectorAnatomyFromEuropeanConcensus16SectorMinimalModel6302,
	"1.2.840.10008.6.1.1141":           ProstateSectorAnatomyFromEuropeanConcensus27SectorOptimalModel6303,
	"1.2.840.10008.6.1.1142":           MeasurementSelectionReason12301,
	"1.2.840.10008.6.1.1143":           EchoFindingObservationType12302,
	"1.2.840.10008.6.1.1144":           EchoMeasurementType12303,
	"1.2.840.10008.6.1.1145":           CardiovascularMeasuredProperty12304,
	"1.2.840.10008.6.1.1146":           BasicEchoAnatomicSite12305,
	"1.2.840.10008.6.1.1147":           EchoFlowDirection12306,
	"1.2.840.10008.6.1.1148":           CardiacPhaseAndTimePoint12307,
	"1.2.840.10008.6.1.1149":           CoreEchoMeasurement12300,
	"1.2.840.10008.6.1.1150":           OCTAProcessingAlgorithmFamily4270,
	"1.2.840.10008.6.1.1151":           EnFaceImageType4271,
	"1.2.840.10008.6.1.1152":           OPTScanPatternType4272,
	"1.2.840.10008.6.1.1153":           RetinalSegmentationSurface4273,
	"1.2.840.10008.6.1.1154":           OrganForRadiationDoseEstimate10060,
	"1.2.840.10008.6.1.1155":           AbsorbedRadiationDoseType10061,
	"1.2.840.10008.6.1.1156":           EquivalentRadiationDoseType10062,
	"1.2.840.10008.6.1.1157":           RadiationDoseEstimateDistributionRepresentation10063,
	"1.2.840.10008.6.1.1158":           PatientModelType10064,
	"1.2.840.10008.6.1.1159":           RadiationTransportModelType10065,
	"1.2.840.10008.6.1.1160":           AttenuatorCategory10066,
	"1.2.840.10008.6.1.1161":           RadiationAttenuatorMaterial10067,
	"1.2.840.10008.6.1.1162":           EstimateMethodType10068,
	"1.2.840.10008.6.1.1163":           RadiationDoseEstimateParameter10069,
	"1.2.840.10008.6.1.1164":           RadiationDoseType10070,
	"1.2.840.10008.6.1.1165":           MRDiffusionComponentSemantic7270,
	"1.2.840.10008.6.1.1166":           MRDiffusionAnisotropyIndex7271,
	"1.2.840.10008.6.1.1167":           MRDiffusionModelParameter7272,
	"1.2.840.10008.6.1.1168":           MRDiffusionModel7273,
	"1.2.840.10008.6.1.1169":           MRDiffusionModelFittingMethod7274,
	"1.2.840.10008.6.1.1170":           MRDiffusionModelSpecificMethod7275,
	"1.2.840.10008.6.1.1171":           MRDiffusionModelInput7276,
	"1.2.840.10008.6.1.1172":           DiffusionRateAreaOverTimeUnit7277,
	"1.2.840.10008.6.1.1173":           PediatricSizeCategory7039,
	"1.2.840.10008.6.1.1174":           CalciumScoringPatientSizeCategory7041,
	"1.2.840.10008.6.1.1175":           ReasonForRepeatingAcquisition10034,
	"1.2.840.10008.6.1.1176":           ProtocolAssertion800,
	"1.2.840.10008.6.1.1177":           RadiotherapeuticDoseMeasurementDevice7026,
	"1.2.840.10008.6.1.1178":           ExportAdditionalInformationDocumentTitle7014,
	"1.2.840.10008.6.1.1179":           ExportDelayReason7015,
	"1.2.840.10008.6.1.1180":           LevelOfDifficulty7016,
	"1.2.840.10008.6.1.1181":           CategoryOfTeachingMaterialImaging7017,
	"1.2.840.10008.6.1.1182":           MiscellaneousDocumentTitle7018,
	"1.2.840.10008.6.1.1183":           SegmentationNonImageSourcePurposeOfReference7019,
	"1.2.840.10008.6.1.1184":           LongitudinalTemporalEventType280,
	"1.2.840.10008.6.1.1185":           NonLesionObjectTypePhysicalObject6401,
	"1.2.840.10008.6.1.1186":           NonLesionObjectTypeSubstance6402,
	"1.2.840.10008.6.1.1187":           NonLesionObjectTypeTissue6403,
	"1.2.840.10008.6.1.1188":           ChestNonLesionObjectTypePhysicalObject6404,
	"1.2.840.10008.6.1.1189":           ChestNonLesionObjectTypeTissue6405,
	"1.2.840.10008.6.1.1190":           TissueSegmentationPropertyType7191,
	"1.2.840.10008.6.1.1191":           AnatomicalStructureSegmentationPropertyType7192,
	"1.2.840.10008.6.1.1192":           PhysicalObjectSegmentationPropertyType7193,
	"1.2.840.10008.6.1.1193":           MorphologicallyAbnormalStructureSegmentationPropertyType7194,
	"1.2.840.10008.6.1.1194":           FunctionSegmentationPropertyType7195,
	"1.2.840.10008.6.1.1195":           SpatialAndRelationalConceptSegmentationPropertyType7196,
	"1.2.840.10008.6.1.1196":           BodySubstanceSegmentationPropertyType7197,
	"1.2.840.10008.6.1.1197":           SubstanceSegmentationPropertyType7198,
	"1.2.840.10008.6.1.1198":           InterpretationRequestDiscontinuationReason9303,
	"1.2.840.10008.6.1.1199":           GrayLevelRunLengthBasedFeature7475,
	"1.2.840.10008.6.1.1200":           GrayLevelSizeZoneBasedFeature7476,
	"1.2.840.10008.6.1.1201":           EncapsulatedDocumentSourcePurposeOfReference7060,
	"1.2.840.10008.6.1.1202":           ModelDocumentTitle7061,
	"1.2.840.10008.6.1.1203":           PurposeOfReferenceToPredecessor3DModel7062,
	"1.2.840.10008.6.1.1204":           ModelScaleUnit7063,
	"1.2.840.10008.6.1.1205":           ModelUsage7064,
	"1.2.840.10008.6.1.1206":           RadiationDoseUnit10071,
	"1.2.840.10008.6.1.1207":           RadiotherapyFiducial7112,
	"1.2.840.10008.6.1.1208":           MultiEnergyRelevantMaterial300,
	"1.2.840.10008.6.1.1209":           MultiEnergyMaterialUnit301,
	"1.2.840.10008.6.1.1210":           DosimetricObjectiveType9500,
	"1.2.840.10008.6.1.1211":           PrescriptionAnatomyCategory9501,
	"1.2.840.10008.6.1.1212":           RTSegmentAnnotationCategory9502,
	"1.2.840.10008.6.1.1213":           RadiotherapyTherapeuticRoleCategory9503,
	"1.2.840.10008.6.1.1214":           RTGeometricInformation9504,
	"1.2.840.10008.6.1.1215":           FixationOrPositioningDevice9505,
	"1.2.840.10008.6.1.1216":           BrachytherapyDevice9506,
	"1.2.840.10008.6.1.1217":           ExternalBodyModel9507,
	"1.2.840.10008.6.1.1218":           NonSpecificVolume9508,
	"1.2.840.10008.6.1.1219":           PurposeOfReferenceForRTPhysicianIntentInput9509,
	"1.2.840.10008.6.1.1220":           PurposeOfReferenceForRTTreatmentPlanningInput9510,
	"1.2.840.10008.6.1.1221":           GeneralExternalRadiotherapyProcedureTechnique9511,
	"1.2.840.10008.6.1.1222":           TomotherapeuticRadiotherapyProcedureTechnique9512,
	"1.2.840.10008.6.1.1223":           FixationDevice9513,
	"1.2.840.10008.6.1.1224":           AnatomicalStructureForRadiotherapy9514,
	"1.2.840.10008.6.1.1225":           RTPatientSupportDevice9515,
	"1.2.840.10008.6.1.1226":           RadiotherapyBolusDeviceType9516,
	"1.2.840.10008.6.1.1227":           RadiotherapyBlockDeviceType9517,
	"1.2.840.10008.6.1.1228":           RadiotherapyAccessoryNoSlotHolderDeviceType9518,
	"1.2.840.10008.6.1.1229":           RadiotherapyAccessorySlotHolderDeviceType9519,
	"1.2.840.10008.6.1.1230":           SegmentedRTAccessoryDevice9520,
	"1.2.840.10008.6.1.1231":           RadiotherapyTreatmentEnergyUnit9521,
	"1.2.840.10008.6.1.1232":           MultiSourceRadiotherapyProcedureTechnique9522,
	"1.2.840.10008.6.1.1233":           RoboticRadiotherapyProcedureTechnique9523,
	"1.2.840.10008.6.1.1234":           RadiotherapyProcedureTechnique9524,
	"1.2.840.10008.6.1.1235":           RadiationTherapyParticle9525,
	"1.2.840.10008.6.1.1236":           IonTherapyParticle9526,
	"1.2.840.10008.6.1.1237":           TeletherapyIsotope9527,
	"1.2.840.10008.6.1.1238":           BrachytherapyIsotope9528,
	"1.2.840.10008.6.1.1239":           SingleDoseDosimetricObjective9529,
	"1.2.840.10008.6.1.1240":           PercentageAndDoseDosimetricObjective9530,
	"1.2.840.10008.6.1.1241":           VolumeAndDoseDosimetricObjective9531,
	"1.2.840.10008.6.1.1242":           NoParameterDosimetricObjective9532,
	"1.2.840.10008.6.1.1243":           DeliveryTimeStructure9533,
	"1.2.840.10008.6.1.1244":           RadiotherapyTarget9534,
	"1.2.840.10008.6.1.1245":           RadiotherapyDoseCalculationRole9535,
	"1.2.840.10008.6.1.1246":           RadiotherapyPrescribingAndSegmentingPersonRole9536,
	"1.2.840.10008.6.1.1247":           EffectiveDoseCalculationMethodCategory9537,
	"1.2.840.10008.6.1.1248":           RadiationTransportBasedEffectiveDoseMethodModifier9538,
	"1.2.840.10008.6.1.1249":           FractionationBasedEffectiveDoseMethodModifier9539,
	"1.2.840.10008.6.1.1250":           ImagingAgentAdministrationAdverseEvent60,
	"1.2.840.10008.6.1.1251":           TimeRelativeToProcedure61RETIRED,
	"1.2.840.10008.6.1.1252":           ImagingAgentAdministrationPhaseType62,
	"1.2.840.10008.6.1.1253":           ImagingAgentAdministrationMode63,
	"1.2.840.10008.6.1.1254":           ImagingAgentAdministrationPatientState64,
	"1.2.840.10008.6.1.1255":           ImagingAgentAdministrationPremedication65,
	"1.2.840.10008.6.1.1256":           ImagingAgentAdministrationMedication66,
	"1.2.840.10008.6.1.1257":           ImagingAgentAdministrationCompletionStatus67,
	"1.2.840.10008.6.1.1258":           ImagingAgentAdministrationPharmaceuticalPresentationUnit68,
	"1.2.840.10008.6.1.1259":           ImagingAgentAdministrationConsumable69,
	"1.2.840.10008.6.1.1260":           Flush70,
	"1.2.840.10008.6.1.1261":           ImagingAgentAdministrationInjectorEventType71,
	"1.2.840.10008.6.1.1262":           ImagingAgentAdministrationStepType72,
	"1.2.840.10008.6.1.1263":           BolusShapingCurve73,
	"1.2.840.10008.6.1.1264":           ImagingAgentAdministrationConsumableCatheterType74,
	"1.2.840.10008.6.1.1265":           LowHighOrEqual75,
	"1.2.840.10008.6.1.1266":           PremedicationType76,
	"1.2.840.10008.6.1.1267":           LateralityWithMedian245,
	"1.2.840.10008.6.1.1268":           DermatologyAnatomicSite4029,
	"1.2.840.10008.6.1.1269":           QuantitativeImageFeature218,
	"1.2.840.10008.6.1.1270":           GlobalShapeDescriptor7477,
	"1.2.840.10008.6.1.1271":           IntensityHistogramFeature7478,
	"1.2.840.10008.6.1.1272":           GreyLevelDistanceZoneBasedFeature7479,
	"1.2.840.10008.6.1.1273":           NeighbourhoodGreyToneDifferenceBasedFeature7500,
	"1.2.840.10008.6.1.1274":           NeighbouringGreyLevelDependenceBasedFeature7501,
	"1.2.840.10008.6.1.1275":           CorneaMeasurementMethodDescriptor4242,
	"1.2.840.10008.6.1.1276":           SegmentedRadiotherapeuticDoseMeasurementDevice7027,
	"1.2.840.10008.6.1.1277":           ClinicalCourseOfDisease6098,
	"1.2.840.10008.6.1.1278":           RacialGroup6099,
	"1.2.840.10008.6.1.1279":           RelativeLaterality246,
	"1.2.840.10008.6.1.1280":           BrainLesionSegmentationTypeWithNecrosis7168,
	"1.2.840.10008.6.1.1281":           BrainLesionSegmentationTypeWithoutNecrosis7169,
	"1.2.840.10008.6.1.1282":           NonAcquisitionModality32,
	"1.2.840.10008.6.1.1283":           Modality33,
	"1.2.840.10008.6.1.1284":           LateralityLeftRightOnly247,
	"1.2.840.10008.6.1.1285":           QualitativeEvaluationModifierType210,
	"1.2.840.10008.6.1.1286":           QualitativeEvaluationModifierValue211,
	"1.2.840.10008.6.1.1287":           GenericAnatomicLocationModifier212,
	"1.2.840.10008.6.1.1288":           BeamLimitingDeviceType9541,
	"1.2.840.10008.6.1.1289":           CompensatorDeviceType9542,
	"1.2.840.10008.6.1.1290":           RadiotherapyTreatmentMachineMode9543,
	"1.2.840.10008.6.1.1291":           RadiotherapyDistanceReferenceLocation9544,
	"1.2.840.10008.6.1.1292":           FixedBeamLimitingDeviceType9545,
	"1.2.840.10008.6.1.1293":           RadiotherapyWedgeType9546,
	"1.2.840.10008.6.1.1294":           RTBeamLimitingDeviceOrientationLabel9547,
	"1.2.840.10008.6.1.1295":           GeneralAccessoryDeviceType9548,
	"1.2.840.10008.6.1.1296":           RadiationGenerationModeType9549,
	"1.2.840.10008.6.1.1297":           CArmPhotonElectronDeliveryRateUnit9550,
	"1.2.840.10008.6.1.1298":           TreatmentDeliveryDeviceType9551,
	"1.2.840.10008.6.1.1299":           CArmPhotonElectronDosimeterUnit9552,
	"1.2.840.10008.6.1.1300":           TreatmentPoint9553,
	"1.2.840.10008.6.1.1301":           EquipmentReferencePoint9554,
	"1.2.840.10008.6.1.1302":           RadiotherapyTreatmentPlanningPersonRole9555,
	"1.2.840.10008.6.1.1303":           RealTimeVideoRenditionTitle7070,
	"1.2.840.10008.6.1.1304":           GeometryGraphicalRepresentation219,
	"1.2.840.10008.6.1.1305":           VisualExplanation217,
	"1.2.840.10008.6.1.1306":           ProstateSectorAnatomyFromPIRADSV216304,
	"1.2.840.10008.6.1.1307":           RadiotherapyRoboticNodeSet9556,
	"1.2.840.10008.6.1.1308":           TomotherapeuticDosimeterUnit9557,
	"1.2.840.10008.6.1.1309":           TomotherapeuticDoseRateUnit9558,
	"1.2.840.10008.6.1.1310":           RoboticDeliveryDeviceDosimeterUnit9559,
	"1.2.840.10008.6.1.1311":           RoboticDeliveryDeviceDoseRateUnit9560,
	"1.2.840.10008.6.1.1312":           AnatomicStructure8134,
	"1.2.840.10008.6.1.1313":           MediastinumFindingOrFeature6148,
	"1.2.840.10008.6.1.1314":           MediastinumAnatomy6149,
	"1.2.840.10008.6.1.1315":           VascularUltrasoundReportDocumentTitle12100,
	"1.2.840.10008.6.1.1316":           OrganPartNonLateralized12130,
	"1.2.840.10008.6.1.1317":           OrganPartLateralized12131,
	"1.2.840.10008.6.1.1318":           TreatmentTerminationReason9561,
	"1.2.840.10008.6.1.1319":           RadiotherapyTreatmentDeliveryPersonRole9562,
	"1.2.840.10008.6.1.1320":           RadiotherapyInterlockResolution9563,
	"1.2.840.10008.6.1.1321":           TreatmentSessionConfirmationAssertion9564,
	"1.2.840.10008.6.1.1322":           TreatmentToleranceViolationCause9565,
	"1.2.840.10008.6.1.1323":           ClinicalToleranceViolationType9566,
	"1.2.840.10008.6.1.1324":           MachineToleranceViolationType9567,
	"1.2.840.10008.6.1.1325":           RadiotherapyTreatmentInterlock9568,
	"1.2.840.10008.6.1.1326":           IsocentricPatientSupportPositionParameter9569,
	"1.2.840.10008.6.1.1327":           RTOverriddenTreatmentParameter9570,
	"1.2.840.10008.6.1.1328":           EEGLead3030,
	"1.2.840.10008.6.1.1329":           LeadLocationNearOrInMuscle3031,
	"1.2.840.10008.6.1.1330":           LeadLocationNearPeripheralNerve3032,
	"1.2.840.10008.6.1.1331":           EOGLead3033,
	"1.2.840.10008.6.1.1332":           BodyPositionChannel3034,
	"1.2.840.10008.6.1.1333":           EEGAnnotationNeurophysiologicEnumeration3035,
	"1.2.840.10008.6.1.1334":           EMGAnnotationNeurophysiologicalEnumeration3036,
	"1.2.840.10008.6.1.1335":           EOGAnnotationNeurophysiologicalEnumeration3037,
	"1.2.840.10008.6.1.1336":           PatternEvent3038,
	"1.2.840.10008.6.1.1337":           DeviceRelatedAndEnvironmentRelatedEvent3039,
	"1.2.840.10008.6.1.1338":           EEGAnnotationNeurologicalMonitoringMeasurement3040,
	"1.2.840.10008.6.1.1339":           OBGYNUltrasoundReportDocumentTitle12024,
	"1.2.840.10008.6.1.1340":           AutomationOfMeasurement7230,
	"1.2.840.10008.6.1.1341":           OBGYNUltrasoundBeamPath12025,
	"1.2.840.10008.6.1.1342":           AngleMeasurement7550,
	"1.2.840.10008.6.1.1343":           GenericPurposeOfReferenceToImagesAndCoordinatesInMeasurement7551,
	"1.2.840.10008.6.1.1344":           GenericPurposeOfReferenceToImagesInMeasurement7552,
	"1.2.840.10008.6.1.1345":           GenericPurposeOfReferenceToCoordinatesInMeasurement7553,
	"1.2.840.10008.6.1.1346":           FitzpatrickSkinType4401,
	"1.2.840.10008.6.1.1347":           HistoryOfMalignantMelanoma4402,
	"1.2.840.10008.6.1.1348":           HistoryOfMelanomaInSitu4403,
	"1.2.840.10008.6.1.1349":           HistoryOfNonMelanomaSkinCancer4404,
	"1.2.840.10008.6.1.1350":           SkinDisorder4405,
	"1.2.840.10008.6.1.1351":           PatientReportedLesionCharacteristic4406,
	"1.2.840.10008.6.1.1352":           LesionPalpationFinding4407,
	"1.2.840.10008.6.1.1353":           LesionVisualFinding4408,
	"1.2.840.10008.6.1.1354":           SkinProcedure4409,
	"1.2.840.10008.6.1.1355":           AbdominopelvicVessel12125,
	"1.2.840.10008.6.1.1356":           NumericValueFailureQualifier43,
	"1.2.840.10008.6.1.1357":           NumericValueUnknownQualifier44,
	"1.2.840.10008.6.1.1358":           CouinaudLiverSegment7170,
	"1.2.840.10008.6.1.1359":           LiverSegmentationType7171,
	"1.2.840.10008.6.1.1360":           ContraindicationsForXAImaging1201,
	"1.2.840.10008.6.1.1361":           NeurophysiologicStimulationMode3041,
	"1.2.840.10008.6.1.1362":           ReportedValueType10072,
	"1.2.840.10008.6.1.1363":           ValueTiming10073,
	"1.2.840.10008.6.1.1364":           RDSRFrameOfReferenceOrigin10074,
	"1.2.840.10008.6.1.1365":           MicroscopyAnnotationPropertyType8135,
	"1.2.840.10008.6.1.1366":           MicroscopyMeasurementType8136,
	"1.2.840.10008.6.1.1367":           ProstateReportingSystem6310,
	"1.2.840.10008.6.1.1368":           MRSignalIntensity6311,
	"1.2.840.10008.6.1.1369":           CrossSectionalScanPlaneOrientation6312,
	"1.2.840.10008.6.1.1370":           HistoryOfProstateDisease6313,
	"1.2.840.10008.6.1.1371":           ProstateMRIStudyQualityFinding6314,
	"1.2.840.10008.6.1.1372":           ProstateMRISeriesQualityFinding6315,
	"1.2.840.10008.6.1.1373":           MRImagingArtifact6316,
	"1.2.840.10008.6.1.1374":           ProstateDCEMRIQualityFinding6317,
	"1.2.840.10008.6.1.1375":           ProstateDWIMRIQualityFinding6318,
	"1.2.840.10008.6.1.1376":           AbdominalInterventionType6319,
	"1.2.840.10008.6.1.1377":           AbdominopelvicIntervention6320,
	"1.2.840.10008.6.1.1378":           ProstateCancerDiagnosticProcedure6321,
	"1.2.840.10008.6.1.1379":           ProstateCancerFamilyHistory6322,
	"1.2.840.10008.6.1.1380":           ProstateCancerTherapy6323,
	"1.2.840.10008.6.1.1381":           ProstateMRIAssessment6324,
	"1.2.840.10008.6.1.1382":           OverallAssessmentFromPIRADS6325,
	"1.2.840.10008.6.1.1383":           ImageQualityControlStandard6326,
	"1.2.840.10008.6.1.1384":           ProstateImagingIndication6327,
	"1.2.840.10008.6.1.1385":           PIRADSV2LesionAssessmentCategory6328,
	"1.2.840.10008.6.1.1386":           PIRADSV2T2WIPZLesionAssessmentCategory6329,
	"1.2.840.10008.6.1.1387":           PIRADSV2T2WITZLesionAssessmentCategory6330,
	"1.2.840.10008.6.1.1388":           PIRADSV2DWILesionAssessmentCategory6331,
	"1.2.840.10008.6.1.1389":           PIRADSV2DCELesionAssessmentCategory6332,
	"1.2.840.10008.6.1.1390":           mpMRIAssessmentType6333,
	"1.2.840.10008.6.1.1391":           mpMRIAssessmentTypeFromPIRADS6334,
	"1.2.840.10008.6.1.1392":           mpMRIAssessmentValue6335,
	"1.2.840.10008.6.1.1393":           MRIAbnormality6336,
	"1.2.840.10008.6.1.1394":           mpMRIProstateAbnormalityFromPIRADS6337,
	"1.2.840.10008.6.1.1395":           mpMRIBenignProstateAbnormalityFromPIRADS6338,
	"1.2.840.10008.6.1.1396":           MRIShapeCharacteristic6339,
	"1.2.840.10008.6.1.1397":           ProstateMRIShapeCharacteristicFromPIRADS6340,
	"1.2.840.10008.6.1.1398":           MRIMarginCharacteristic6341,
	"1.2.840.10008.6.1.1399":           ProstateMRIMarginCharacteristicFromPIRADS6342,
	"1.2.840.10008.6.1.1400":           MRISignalCharacteristic6343,
	"1.2.840.10008.6.1.1401":           ProstateMRISignalCharacteristicFromPIRADS6344,
	"1.2.840.10008.6.1.1402":           MRIEnhancementPattern6345,
	"1.2.840.10008.6.1.1403":           ProstateMRIEnhancementPatternFromPIRADS6346,
	"1.2.840.10008.6.1.1404":           ProstateMRIExtraProstaticFinding6347,
	"1.2.840.10008.6.1.1405":           ProstateMRIAssessmentOfExtraProstaticAnatomicSite6348,
	"1.2.840.10008.6.1.1406":           MRCoilType6349,
	"1.2.840.10008.6.1.1407":           EndorectalCoilFillSubstance6350,
	"1.2.840.10008.6.1.1408":           ProstateRelationalMeasurement6351,
	"1.2.840.10008.6.1.1409":           ProstateCancerDiagnosticBloodLabMeasurement6352,
	"1.2.840.10008.6.1.1410":           ProstateImagingTypesOfQualityControlStandard6353,
	"1.2.840.10008.6.1.1411":           UltrasoundShearWaveMeasurement12308,
	"1.2.840.10008.6.1.1412":           LeftVentricleMyocardialWall16SegmentModel3780RETIRED,
	"1.2.840.10008.6.1.1413":           LeftVentricleMyocardialWall18SegmentModel3781,
	"1.2.840.10008.6.1.1414":           LeftVentricleBasalWall6Segments3782,
	"1.2.840.10008.6.1.1415":           LeftVentricleMidlevelWall6Segments3783,
	"1.2.840.10008.6.1.1416":           LeftVentricleApicalWall4Segments3784,
	"1.2.840.10008.6.1.1417":           LeftVentricleApicalWall6Segments3785,
	"1.2.840.10008.6.1.1418":           PatientTreatmentPreparationMethod9571,
	"1.2.840.10008.6.1.1419":           PatientShieldingDevice9572,
	"1.2.840.10008.6.1.1420":           PatientTreatmentPreparationDevice9573,
	"1.2.840.10008.6.1.1421":           PatientPositionDisplacementReferencePoint9574,
	"1.2.840.10008.6.1.1422":           PatientAlignmentDevice9575,
	"1.2.840.10008.6.1.1423":           ReasonsForRTRadiationTreatmentOmission9576,
	"1.2.840.10008.6.1.1424":           PatientTreatmentPreparationProcedure9577,
	"1.2.840.10008.6.1.1425":           MotionManagementSetupDevice9578,
	"1.2.840.10008.6.1.1426":           CoreEchoStrainMeasurement12309,
	"1.2.840.10008.6.1.1427":           MyocardialStrainMethod12310,
	"1.2.840.10008.6.1.1428":           EchoMeasuredStrainProperty12311,
	"1.2.840.10008.6.1.1429":           AssessmentFromCADRADS3020,
	"1.2.840.10008.6.1.1430":           CADRADSStenosisAssessmentModifier3021,
	"1.2.840.10008.6.1.1431":           CADRADSAssessmentModifier3022,
	"1.2.840.10008.6.1.1432":           RTSegmentMaterial9579,
	"1.2.840.10008.6.1.1433":           VertebralAnatomicStructure7602,
	"1.2.840.10008.6.1.1434":           Vertebra7603,
	"1.2.840.10008.6.1.1435":           IntervertebralDisc7604,
	"1.2.840.10008.6.1.1436":           ImagingProcedure101,
	"1.2.840.10008.6.1.1437":           NICIPShortCodeImagingProcedure103,
	"1.2.840.10008.6.1.1438":           NICIPSNOMEDImagingProcedure104,
	"1.2.840.10008.6.1.1439":           ICD10PCSImagingProcedure105,
	"1.2.840.10008.6.1.1440":           ICD10PCSNuclearMedicineProcedure106,
	"1.2.840.10008.6.1.1441":           ICD10PCSRadiationTherapyProcedure107,
	"1.2.840.10008.6.1.1442":           RTSegmentationPropertyCategory9580,
	"1.2.840.10008.6.1.1443":           RadiotherapyRegistrationMark9581,
	"1.2.840.10008.6.1.1444":           RadiotherapyDoseRegion9582,
	"1.2.840.10008.6.1.1445":           AnatomicallyLocalizedLesionSegmentationType7199,
	"1.2.840.10008.6.1.1446":           ReasonForRemovalFromOperationalUse7031,
	"1.2.840.10008.6.1.1447":           GeneralUltrasoundReportDocumentTitle12320,
	"1.2.840.10008.6.1.1448":           ElastographySite12321,
	"1.2.840.10008.6.1.1449":           ElastographyMeasurementSite12322,
	"1.2.840.10008.6.1.1450":           UltrasoundRelevantPatientCondition12323,
	"1.2.840.10008.6.1.1451":           ShearWaveDetectionMethod12324,
	"1.2.840.10008.6.1.1452":           LiverUltrasoundStudyIndication12325,
	"1.2.840.10008.6.1.1453":           AnalogWaveformFilter3042,
	"1.2.840.10008.6.1.1454":           DigitalWaveformFilter3043,
	"1.2.840.10008.6.1.1455":           WaveformFilterLookupTableInputFrequencyUnit3044,
	"1.2.840.10008.6.1.1456":           WaveformFilterLookupTableOutputMagnitudeUnit3045,
	"1.2.840.10008.6.1.1457":           SpecificObservationSubjectClass272,
	"1.2.840.10008.6.1.1458":           MovableBeamLimitingDeviceType9540,
	"1.2.840.10008.6.1.1459":           RadiotherapyAcquisitionWorkItemSubtasks9260,
	"1.2.840.10008.6.1.1460":           PatientPositionAcquisitionRadiationSourceLocations9261,
	"1.2.840.10008.6.1.1461":           EnergyDerivationTypes9262,
	"1.2.840.10008.6.1.1462":           KVImagingAcquisitionTechniques9263,
	"1.2.840.10008.6.1.1463":           MVImagingAcquisitionTechniques9264,
	"1.2.840.10008.6.1.1464":           PatientPositionAcquisitionProjectionTechniques9265,
	"1.2.840.10008.6.1.1465":           PatientPositionAcquisitionCTTechniques9266,
	"1.2.840.10008.6.1.1466":           PatientPositioningRelatedObjectPurposes9267,
	"1.2.840.10008.6.1.1467":           PatientPositionAcquisitionDevices9268,
	"1.2.840.10008.6.1.1468":           RTRadiationMetersetUnits9269,
	"1.2.840.10008.6.1.1469":           AcquisitionInitiationTypes9270,
	"1.2.840.10008.6.1.1470":           RTImagePatientPositionAcquisitionDevices9271,
	"1.2.840.10008.6.1.1471":           PhotoacousticIlluminationMethod11001,
	"1.2.840.10008.6.1.1472":           AcousticCouplingMedium11002,
	"1.2.840.10008.6.1.1473":           UltrasoundTransducerTechnology11003,
	"1.2.840.10008.6.1.1474":           SpeedOfSoundCorrectionMechanisms11004,
	"1.2.840.10008.6.1.1475":           PhotoacousticReconstructionAlgorithmFamily11005,
	"1.2.840.10008.6.1.1476":           PhotoacousticImagedProperty11006,
	"1.2.840.10008.6.1.1477":           XRayRadiationDoseProcedureTypeReported10005,
	"1.2.840.10008.6.1.1478":           TopicalTreatment4410,
	"1.2.840.10008.6.1.1479":           LesionColor4411,
	"1.2.840.10008.6.1.1480":           SpecimenStainForConfocalMicroscopy4412,
	"1.2.840.10008.6.1.1481":           RTROIImageAcquisitionContext9272,
	"1.2.840.10008.6.1.1482":           LobeOfLung6170,
	"1.2.840.10008.6.1.1483":           ZoneOfLung6171,
	"1.2.840.10008.6.1.1484":           SleepStage3046,
	"1.2.840.10008.6.1.1485":           PatientPositionAcquisitionMRTechniques9273,
	"1.2.840.10008.6.1.1486":           RTPlanRadiotherapyProcedureTechnique9583,
	"1.2.840.10008.6.1.1487":           WaveformAnnotationClassification3047,
	"1.2.840.10008.6.1.1488":           WaveformAnnotationsDocumentTitle3048,
	"1.2.840.10008.6.1.1489":           EEGProcedure3049,
	"1.2.840.10008.6.1.1490":           PatientConsciousness3050,
	"1.2.840.10008.6.1.1491":           FollicleType12010,
	"1.2.840.10008.6.1.1492":           BreastTissueSegmentationType7163,
	"1.2.840.10008.6.1.1493":           ImplantedDevice3779,
	"1.2.840.10008.6.1.1494":           SimilarityMeasure281,
	"1.2.840.10008.6.1.1495":           WaveformAcquisitionModality34,
	"1.2.840.10008.6.1.1496":           EnFaceProcessingAlgorithmFamily4274,
	"1.2.840.10008.6.1.1497":           AnteriorEyeSegmentationSurface4275,
	"1.2.840.10008.6.1.1498":           FetalEchocardiographyImageView12312,
	"1.2.840.10008.6.1.1499":           CardiacUltrasoundFetalArrhythmiaMeasurements12313,
	"1.2.840.10008.6.1.1500":           CommonFetalEchocardiographyMeasurements12314,
	"1.2.840.10008.6.1.1501":           HeadAndNeckPrimaryAnatomicStructure4061,
	"1.2.840.10008.6.1.1502":           VLView4062,
	"1.2.840.10008.6.1.1503":           VLDentalView4063,
	"1.2.840.10008.6.1.1504":           VLViewModifier4064,
	"1.2.840.10008.6.1.1505":           VLDentalViewModifier4065,
	"1.2.840.10008.6.1.1506":           OrthognathicFunctionalCondition4066,
	"1.2.840.10008.6.1.1507":           OrthodonticFindingByInspection4067,
	"1.2.840.10008.6.1.1508":           OrthodonticObservableEntity4068,
	"1.2.840.10008.6.1.1509":           DentalOcclusion4069,
	"1.2.840.10008.6.1.1510":           OrthodonticTreatmentProgress4070,
	"1.2.840.10008.6.1.1511":           GeneralPhotographyDevice4071,
	"1.2.840.10008.6.1.1512":           DevicesForThePurposeOfDentalPhotography4072,
	"1.2.840.10008.6.1.1513":           CTDIPhantomDevice4053,
	"1.2.840.10008.6.1.1514":           DiagnosticImagingProcedureWithoutIVContrast108,
	"1.2.840.10008.6.1.1515":           DiagnosticImagingProcedureWithIVContrast109,
	"1.2.840.10008.6.1.1516":           StructuralHeartProcedure12331,
	"1.2.840.10008.6.1.1517":           StructuralHeartDevice12332,
	"1.2.840.10008.6.1.1518":           StructuralHeartMeasurement12333,
	"1.2.840.10008.6.1.1519":           AorticValveStructuralMeasurement12334,
	"1.2.840.10008.6.1.1520":           MitralValveStructuralMeasurement12335,
	"1.2.840.10008.6.1.1521":           TricuspidValveStructuralMeasurement12336,
	"1.2.840.10008.6.1.1522":           StructuralHeartEchoMeasurement12337,
	"1.2.840.10008.6.1.1523":           LeftAtrialAppendageClosureMeasurement12338,
	"1.2.840.10008.6.1.1524":           StructuralHeartProcedureAnatomicSite12339,
	"1.2.840.10008.6.1.1525":           IndicationForStructuralHeartProcedure12341,
	"1.2.840.10008.6.1.1526":           BradycardiacAgent12342,
	"1.2.840.10008.6.1.1527":           TransesophagealEchocardiographyScanPlane12343,
	"1.2.840.10008.6.1.1528":           StructuralHeartMeasurementReportDocumentTitle12344,
	"1.2.840.10008.6.1.1529":           PersonGenderIdentity7458,
	"1.2.840.10008.6.1.1530":           CategoryOfSexParametersForClinicalUse7459,
	"1.2.840.10008.6.1.1531":           ThirdPersonPronounSet7448,
	"1.2.840.10008.6.1.1532":           CardiacStructureCalcificationQualitativeEvaluation12345,
	"1.2.840.10008.6.1.1533":           VisualFieldMeasurements4280,
	"1.2.840.10008.6.1.1534":           OpticDiscKeyMeasurements4281,
	"1.2.840.10008.6.1.1535":           RetinalSectorMethods4282,
	"1.2.840.10008.6.1.1536":           RNFLSectorMeasurements4283,
	"1.2.840.10008.6.1.1537":           RNFLClockfaceMeasurements4284,
	"1.2.840.10008.6.1.1538":           MacularThicknessKeyMeasurements4285,
	"1.2.840.10008.6.1.1539":           GanglionCellMeasurementExtent4286,
	"1.2.840.10008.6.1.1540":           GanglionCellKeyMeasurements4287,
	"1.2.840.10008.6.1.1541":           GanglionCellSectorMeasurements4288,
	"1.2.840.10008.6.1.1542":           GanglionCellSectorMethods4289,
	"1.2.840.10008.6.1.1543":           EndothelialCellCountMeasurements4290,
	"1.2.840.10008.6.1.1544":           OphthalmicImageROIMeasurements4291,
	"1.2.840.10008.6.1.1545":           RTPlanApprovalAssertion9584,
	"1.2.840.10008.6.1.1546":           EstimatedDeliveryDateMethod12026,
	"1.2.840.10008.6.1.1547":           RTDoseCalculationAlgorithmFamily9585,
	"1.2.840.10008.6.1.1548":           DoseIndexForDoseCalibration10012,
	"1.2.840.10008.6.1.1549":           UltrasoundAttenuationImagingSite12036,
	"1.2.840.10008.6.1.1550":           FetalAnatomySurveyAssessment12040,
	"1.2.840.10008.6.1.1551":           FetalAnatomySurveyAssessmentHead12041,
	"1.2.840.10008.6.1.1552":           FetalAnatomySurveyAssessmentFaceAndNeck12042,
	"1.2.840.10008.6.1.1553":           FetalAnatomySurveyAssessmentChest12043,
	"1.2.840.10008.6.1.1554":           FetalAnatomySurveyAssessmentHeart12044,
	"1.2.840.10008.6.1.1555":           FetalAnatomySurveyAssessmentAbdomenAndPelvis12045,
	"1.2.840.10008.6.1.1556":           FetalAnatomySurveyAssessmentSpine12046,
	"1.2.840.10008.6.1.1557":           FetalAnatomySurveyAssessmentExtremities12047,
	"1.2.840.10008.6.1.1558":           FetalAnatomySurveyAssessmentMaternal12048,
	"1.2.840.10008.6.1.1559":           FetalAnatomySurveyPracticeGuideline12049,
	"1.2.840.10008.6.1.1560":           SensitiveContentCategory900,
	"1.2.840.10008.6.1.1561":           SensitiveContentDetail901,
	"1.2.840.10008.6.1.1562":           ApplicationTypeCode406,
	"1.2.840.10008.6.1.1563":           XRayModulationType10035,
	"1.2.840.10008.6.1.1564":           RadiotherapyDoseRealWorldUnits9586,
	"1.2.840.10008.6.1.1565":           RadiotherapyDoseInterpretedTypeCodes9587,
	"1.2.840.10008.6.1.1566":           RadiotherapyDoseInterpretedTypeModifierCodes9588,
	"1.2.840.10008.6.1.1567":           RadiotherapyDoseIntentCodes9589,
	"1.2.840.10008.6.1.1568":           QualitySegmentationPropertyType7164,
	"1.2.840.10008.6.1.1569":           UltrasoundZScorePopulationIndex12027,
	"1.2.840.10008.6.1.1570":           FetalUltrasoundZScoreReferenceAuthority12028,
	"1.2.840.10008.6.1.1571":           MetalArtifactReductionAlgorithmFamily10036,
}
