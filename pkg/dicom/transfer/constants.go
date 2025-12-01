// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transfer

import (
	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

// Standard DICOM Transfer Syntaxes
var (
	// ImplicitVRBigEndian is a virtual transfer syntax for implicit VR with big endian byte order.
	ImplicitVRBigEndian = NewBuilder(uid.New(
		uid.ExplicitVRBigEndianRETIRED.UID()+".123456",
		"Implicit VR Big Endian",
		uid.TypeTransferSyntax,
		false)).
		SetExplicitVR(false).
		SetEndian(endian.Big).
		Build()

	// GEPrivateImplicitVRBigEndian is a private GE transfer syntax.
	// Same as Implicit VR Little Endian except for big endian pixel data.
	GEPrivateImplicitVRBigEndian = NewBuilder(uid.GEPrivateImplicitVRBigEndian).
					SetExplicitVR(false).
					SetEndian(endian.Little).
					SetSwapPixelData(true).
					Build()

	// ImplicitVRLittleEndian is the default DICOM transfer syntax.
	// Implicit VR, Little Endian.
	ImplicitVRLittleEndian = NewBuilder(uid.ImplicitVRLittleEndian).
				SetExplicitVR(false).
				SetEndian(endian.Little).
				Build()

	// ExplicitVRLittleEndian is the most common explicit VR transfer syntax.
	// Explicit VR, Little Endian.
	ExplicitVRLittleEndian = NewBuilder(uid.ExplicitVRLittleEndian).
				SetExplicitVR(true).
				SetEndian(endian.Little).
				Build()

	// ExplicitVRBigEndian is a retired transfer syntax.
	// Explicit VR, Big Endian (Retired).
	ExplicitVRBigEndian = NewBuilder(uid.ExplicitVRBigEndianRETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEndian(endian.Big).
				Build()

	// DeflatedExplicitVRLittleEndian uses deflate compression.
	// Deflated Explicit VR Little Endian.
	DeflatedExplicitVRLittleEndian = NewBuilder(uid.DeflatedExplicitVRLittleEndian).
					SetExplicitVR(true).
					SetDeflate(true).
					SetEndian(endian.Little).
					Build()

	// Papyrus3ImplicitVRLittleEndianRetired is a retired Papyrus transfer syntax.
	Papyrus3ImplicitVRLittleEndianRetired = NewBuilder(uid.Papyrus3ImplicitVRLittleEndianRETIRED).
						SetRetired(true).
						SetExplicitVR(false).
						SetEndian(endian.Little).
						Build()

	// JPEGBaseline8Bit is JPEG Baseline (Process 1).
	// Lossy JPEG compression, 8-bit.
	JPEGBaseline8Bit = NewBuilder(uid.JPEGBaseline8Bit).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess1 is an alias for JPEGBaseline8Bit.
	JPEGProcess1 = JPEGBaseline8Bit

	// JPEGExtended12Bit is JPEG Extended (Process 2 & 4).
	// Lossy JPEG compression, 12-bit.
	JPEGExtended12Bit = NewBuilder(uid.JPEGExtended12Bit).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess2_4 is an alias for JPEGExtended12Bit.
	JPEGProcess2_4 = JPEGExtended12Bit

	// JPEGProcess3_5Retired is JPEG Extended (Process 3 & 5) (Retired).
	JPEGProcess3_5Retired = NewBuilder(uid.JPEGExtended35RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess6_8Retired is JPEG Spectral Selection, Non-Hierarchical (Process 6 & 8) (Retired).
	JPEGProcess6_8Retired = NewBuilder(uid.JPEGSpectralSelectionNonHierarchical68RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess7_9Retired is JPEG Spectral Selection, Non-Hierarchical (Process 7 & 9) (Retired).
	JPEGProcess7_9Retired = NewBuilder(uid.JPEGSpectralSelectionNonHierarchical79RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess10_12Retired is JPEG Full Progression, Non-Hierarchical (Process 10 & 12) (Retired).
	JPEGProcess10_12Retired = NewBuilder(uid.JPEGFullProgressionNonHierarchical1012RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess11_13Retired is JPEG Full Progression, Non-Hierarchical (Process 11 & 13) (Retired).
	JPEGProcess11_13Retired = NewBuilder(uid.JPEGFullProgressionNonHierarchical1113RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGLossless is JPEG Lossless, Non-Hierarchical (Process 14).
	// Lossless JPEG compression.
	JPEGLossless = NewBuilder(uid.JPEGLossless).
			SetExplicitVR(true).
			SetEncapsulated(true).
			SetEndian(endian.Little).
			Build()

	// JPEGProcess14 is an alias for JPEGLossless.
	JPEGProcess14 = JPEGLossless

	// JPEGProcess15Retired is JPEG Lossless, Non-Hierarchical (Process 15) (Retired).
	JPEGProcess15Retired = NewBuilder(uid.JPEGLosslessNonHierarchical15RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetEndian(endian.Little).
				Build()

	// JPEGProcess16_18Retired is JPEG Extended, Hierarchical (Process 16 & 18) (Retired).
	JPEGProcess16_18Retired = NewBuilder(uid.JPEGExtendedHierarchical1618RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess17_19Retired is JPEG Extended, Hierarchical (Process 17 & 19) (Retired).
	JPEGProcess17_19Retired = NewBuilder(uid.JPEGExtendedHierarchical1719RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess20_22Retired is JPEG Spectral Selection, Hierarchical (Process 20 & 22) (Retired).
	JPEGProcess20_22Retired = NewBuilder(uid.JPEGSpectralSelectionHierarchical2022RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess21_23Retired is JPEG Spectral Selection, Hierarchical (Process 21 & 23) (Retired).
	JPEGProcess21_23Retired = NewBuilder(uid.JPEGSpectralSelectionHierarchical2123RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess24_26Retired is JPEG Full Progression, Hierarchical (Process 24 & 26) (Retired).
	JPEGProcess24_26Retired = NewBuilder(uid.JPEGFullProgressionHierarchical2426RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess25_27Retired is JPEG Full Progression, Hierarchical (Process 25 & 27) (Retired).
	JPEGProcess25_27Retired = NewBuilder(uid.JPEGFullProgressionHierarchical2527RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGProcess28Retired is JPEG Lossless, Hierarchical (Process 28) (Retired).
	JPEGProcess28Retired = NewBuilder(uid.JPEGLosslessHierarchical28RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetEndian(endian.Little).
				Build()

	// JPEGProcess29Retired is JPEG Lossless, Hierarchical (Process 29) (Retired).
	JPEGProcess29Retired = NewBuilder(uid.JPEGLosslessHierarchical29RETIRED).
				SetRetired(true).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetEndian(endian.Little).
				Build()

	// JPEGLosslessSV1 is JPEG Lossless, Non-Hierarchical, First-Order Prediction.
	// JPEG Lossless, Process 14 [Selection Value 1].
	JPEGLosslessSV1 = NewBuilder(uid.JPEGLosslessSV1).
			SetExplicitVR(true).
			SetEncapsulated(true).
			SetEndian(endian.Little).
			Build()

	// JPEGProcess14SV1 is an alias for JPEGLosslessSV1.
	JPEGProcess14SV1 = JPEGLosslessSV1

	// JPEGLSLossless is JPEG-LS Lossless Image Compression.
	JPEGLSLossless = NewBuilder(uid.JPEGLSLossless).
			SetExplicitVR(true).
			SetEncapsulated(true).
			SetEndian(endian.Little).
			Build()

	// JPEGLSNearLossless is JPEG-LS Lossy (Near-Lossless) Image Compression.
	JPEGLSNearLossless = NewBuilder(uid.JPEGLSNearLossless).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_14495_1").
				SetEndian(endian.Little).
				Build()

	// JPEG2000Lossless is JPEG 2000 Image Compression (Lossless Only).
	JPEG2000Lossless = NewBuilder(uid.JPEG2000Lossless).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetEndian(endian.Little).
				Build()

	// JPEG2000 is JPEG 2000 Image Compression.
	// Can be lossy or lossless.
	JPEG2000 = NewBuilder(uid.JPEG2000).
			SetExplicitVR(true).
			SetEncapsulated(true).
			SetLossy(true, "ISO_15444_1").
			SetEndian(endian.Little).
			Build()

	// JPEG2000Lossy is an alias for JPEG2000.
	JPEG2000Lossy = JPEG2000

	// JPEG2000Part2MultiComponentLosslessOnly is JPEG 2000 Part 2 Multi-component Image Compression (Lossless Only).
	JPEG2000Part2MultiComponentLosslessOnly = NewBuilder(uid.JPEG2000MCLossless).
						SetExplicitVR(true).
						SetEncapsulated(true).
						SetEndian(endian.Little).
						Build()

	// JPEG2000Part2MultiComponent is JPEG 2000 Part 2 Multi-component Image Compression.
	JPEG2000Part2MultiComponent = NewBuilder(uid.JPEG2000MC).
					SetExplicitVR(true).
					SetEncapsulated(true).
					SetLossy(true, "ISO_15444_2").
					SetEndian(endian.Little).
					Build()

	// JPIPReferenced is JPIP Referenced.
	JPIPReferenced = NewBuilder(uid.JPIPReferenced).
			SetExplicitVR(true).
			SetEndian(endian.Little).
			Build()

	// JPIPReferencedDeflate is JPIP Referenced with deflate compression.
	JPIPReferencedDeflate = NewBuilder(uid.JPIPReferencedDeflate).
				SetExplicitVR(true).
				SetDeflate(true).
				SetEndian(endian.Little).
				Build()

	// MPEG2 is MPEG2 Main Profile @ Main Level.
	MPEG2 = NewBuilder(uid.MPEG2MPML).
		SetExplicitVR(true).
		SetEncapsulated(true).
		SetLossy(true, "ISO_13818_2").
		SetEndian(endian.Little).
		Build()

	// FragmentableMPEG2 is fragmentable MPEG2 Main Profile @ Main Level.
	FragmentableMPEG2 = NewBuilder(uid.MPEG2MPMLF).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_13818_2").
				SetEndian(endian.Little).
				Build()

	// MPEG2MainProfileHighLevel is MPEG2 Main Profile / High Level.
	MPEG2MainProfileHighLevel = NewBuilder(uid.MPEG2MPHL).
					SetExplicitVR(true).
					SetEncapsulated(true).
					SetLossy(true, "ISO_13818_2").
					SetEndian(endian.Little).
					Build()

	// FragmentableMPEG2MainProfileHighLevel is fragmentable MPEG2 Main Profile / High Level.
	FragmentableMPEG2MainProfileHighLevel = NewBuilder(uid.MPEG2MPHLF).
						SetExplicitVR(true).
						SetEncapsulated(true).
						SetLossy(true, "ISO_13818_2").
						SetEndian(endian.Little).
						Build()

	// MPEG4AVCH264HighProfileLevel41 is MPEG-4 AVC/H.264 High Profile / Level 4.1.
	MPEG4AVCH264HighProfileLevel41 = NewBuilder(uid.MPEG4HP41).
					SetExplicitVR(true).
					SetEncapsulated(true).
					SetLossy(true, "ISO_14496_10").
					SetEndian(endian.Little).
					Build()

	// FragmentableMPEG4AVCH264HighProfileLevel41 is fragmentable MPEG-4 AVC/H.264 High Profile / Level 4.1.
	FragmentableMPEG4AVCH264HighProfileLevel41 = NewBuilder(uid.MPEG4HP41F).
							SetExplicitVR(true).
							SetEncapsulated(true).
							SetLossy(true, "ISO_14496_10").
							SetEndian(endian.Little).
							Build()

	// MPEG4AVCH264BDCompatibleHighProfileLevel41 is MPEG-4 AVC/H.264 BD-compatible High Profile / Level 4.1.
	MPEG4AVCH264BDCompatibleHighProfileLevel41 = NewBuilder(uid.MPEG4HP41BD).
							SetExplicitVR(true).
							SetEncapsulated(true).
							SetLossy(true, "ISO_14496_10").
							SetEndian(endian.Little).
							Build()

	// FragmentableMPEG4AVCH264BDCompatibleHighProfileLevel41 is fragmentable MPEG-4 AVC/H.264 BD-compatible High Profile / Level 4.1.
	FragmentableMPEG4AVCH264BDCompatibleHighProfileLevel41 = NewBuilder(uid.MPEG4HP41BDF).
								SetExplicitVR(true).
								SetEncapsulated(true).
								SetLossy(true, "ISO_14496_10").
								SetEndian(endian.Little).
								Build()

	// MPEG4AVCH264HighProfileLevel42For2DVideo is MPEG-4 AVC/H.264 High Profile / Level 4.2 For 2D Video.
	MPEG4AVCH264HighProfileLevel42For2DVideo = NewBuilder(uid.MPEG4HP422D).
							SetExplicitVR(true).
							SetEncapsulated(true).
							SetLossy(true, "ISO_14496_10").
							SetEndian(endian.Little).
							Build()

	// FragmentableMPEG4AVCH264HighProfileLevel42For2DVideo is fragmentable MPEG-4 AVC/H.264 High Profile / Level 4.2 For 2D Video.
	FragmentableMPEG4AVCH264HighProfileLevel42For2DVideo = NewBuilder(uid.MPEG4HP422DF).
								SetExplicitVR(true).
								SetEncapsulated(true).
								SetLossy(true, "ISO_14496_10").
								SetEndian(endian.Little).
								Build()

	// MPEG4AVCH264HighProfileLevel42For3DVideo is MPEG-4 AVC/H.264 High Profile / Level 4.2 For 3D Video.
	MPEG4AVCH264HighProfileLevel42For3DVideo = NewBuilder(uid.MPEG4HP423D).
							SetExplicitVR(true).
							SetEncapsulated(true).
							SetLossy(true, "ISO_14496_10").
							SetEndian(endian.Little).
							Build()

	// FragmentableMPEG4AVCH264HighProfileLevel42For3DVideo is fragmentable MPEG-4 AVC/H.264 High Profile / Level 4.2 For 3D Video.
	FragmentableMPEG4AVCH264HighProfileLevel42For3DVideo = NewBuilder(uid.MPEG4HP423DF).
								SetExplicitVR(true).
								SetEncapsulated(true).
								SetLossy(true, "ISO_14496_10").
								SetEndian(endian.Little).
								Build()

	// MPEG4AVCH264StereoHighProfileLevel42 is MPEG-4 AVC/H.264 Stereo High Profile / Level 4.2.
	MPEG4AVCH264StereoHighProfileLevel42 = NewBuilder(uid.MPEG4HP42STEREO).
						SetExplicitVR(true).
						SetEncapsulated(true).
						SetLossy(true, "ISO_14496_10").
						SetEndian(endian.Little).
						Build()

	// FragmentableMPEG4AVCH264StereoHighProfileLevel42 is fragmentable MPEG-4 AVC/H.264 Stereo High Profile / Level 4.2.
	FragmentableMPEG4AVCH264StereoHighProfileLevel42 = NewBuilder(uid.MPEG4HP42STEREOF).
								SetExplicitVR(true).
								SetEncapsulated(true).
								SetLossy(true, "ISO_14496_10").
								SetEndian(endian.Little).
								Build()

	// HEVCH265MainProfileLevel51 is HEVC/H.265 Main Profile / Level 5.1.
	HEVCH265MainProfileLevel51 = NewBuilder(uid.HEVCMP51).
					SetExplicitVR(true).
					SetEncapsulated(true).
					SetLossy(true, "ISO_23008_2").
					SetEndian(endian.Little).
					Build()

	// HEVCH265Main10ProfileLevel51 is HEVC/H.265 Main 10 Profile / Level 5.1.
	HEVCH265Main10ProfileLevel51 = NewBuilder(uid.HEVCM10P51).
					SetExplicitVR(true).
					SetEncapsulated(true).
					SetLossy(true, "ISO_23008_2").
					SetEndian(endian.Little).
					Build()

	// HTJ2KLossless is High-Throughput JPEG 2000 Image Compression (Lossless Only).
	HTJ2KLossless = NewBuilder(uid.HTJ2KLossless).
			SetExplicitVR(true).
			SetEncapsulated(true).
			SetEndian(endian.Little).
			Build()

	// HTJ2KLosslessRPCL is High-Throughput JPEG 2000 with RPCL Options Image Compression (Lossless Only).
	HTJ2KLosslessRPCL = NewBuilder(uid.HTJ2KLosslessRPCL).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetEndian(endian.Little).
				Build()

	// HTJ2K is High-Throughput JPEG 2000 Image Compression.
	HTJ2K = NewBuilder(uid.HTJ2K).
		SetExplicitVR(true).
		SetEncapsulated(true).
		SetLossy(true, "ISO_15444_15").
		SetEndian(endian.Little).
		Build()

	// JPIPHTJ2KReferenced is JPIP HTJ2K Referenced.
	JPIPHTJ2KReferenced = NewBuilder(uid.JPIPHTJ2KReferenced).
				SetExplicitVR(true).
				SetEndian(endian.Little).
				Build()

	// JPIPHTJ2KReferencedDeflate is JPIP HTJ2K Referenced Deflate.
	JPIPHTJ2KReferencedDeflate = NewBuilder(uid.JPIPHTJ2KReferencedDeflate).
					SetExplicitVR(true).
					SetDeflate(true).
					SetEndian(endian.Little).
					Build()

	// RLELossless is RLE Lossless compression.
	RLELossless = NewBuilder(uid.RLELossless).
			SetExplicitVR(true).
			SetEncapsulated(true).
			SetEndian(endian.Little).
			Build()

	// RFC2557MIMEEncapsulation is RFC 2557 MIME encapsulation.
	RFC2557MIMEEncapsulation = NewBuilder(uid.RFC2557MIMEEncapsulationRETIRED).
					SetExplicitVR(true).
					SetRetired(true).
					SetEndian(endian.Little).
					Build()

	// XMLEncoding is XML Encoding (Retired).
	XMLEncoding = NewBuilder(uid.XMLEncodingRETIRED).
			SetExplicitVR(true).
			SetRetired(true).
			SetEndian(endian.Little).
			Build()
)

// Common transfer syntax groups for convenience
var (
	// UncompressedTransferSyntaxes are transfer syntaxes without compression.
	UncompressedTransferSyntaxes = []*Syntax{
		ImplicitVRLittleEndian,
		ExplicitVRLittleEndian,
		ExplicitVRBigEndian,
		ImplicitVRBigEndian,
		GEPrivateImplicitVRBigEndian,
		Papyrus3ImplicitVRLittleEndianRetired,
	}

	// LosslessTransferSyntaxes are transfer syntaxes with lossless compression.
	LosslessTransferSyntaxes = []*Syntax{
		DeflatedExplicitVRLittleEndian,
		JPEGLossless,
		JPEGProcess15Retired,
		JPEGProcess28Retired,
		JPEGProcess29Retired,
		JPEGLosslessSV1,
		JPEGLSLossless,
		JPEG2000Lossless,
		JPEG2000Part2MultiComponentLosslessOnly,
		HTJ2KLossless,
		HTJ2KLosslessRPCL,
		RLELossless,
	}

	// LossyTransferSyntaxes are transfer syntaxes with lossy compression.
	LossyTransferSyntaxes = []*Syntax{
		JPEGBaseline8Bit,
		JPEGExtended12Bit,
		JPEGProcess3_5Retired,
		JPEGProcess6_8Retired,
		JPEGProcess7_9Retired,
		JPEGProcess10_12Retired,
		JPEGProcess11_13Retired,
		JPEGProcess16_18Retired,
		JPEGProcess17_19Retired,
		JPEGProcess20_22Retired,
		JPEGProcess21_23Retired,
		JPEGProcess24_26Retired,
		JPEGProcess25_27Retired,
		JPEGLSNearLossless,
		JPEG2000,
		JPEG2000Part2MultiComponent,
		MPEG2,
		FragmentableMPEG2,
		MPEG2MainProfileHighLevel,
		FragmentableMPEG2MainProfileHighLevel,
		MPEG4AVCH264HighProfileLevel41,
		FragmentableMPEG4AVCH264HighProfileLevel41,
		MPEG4AVCH264BDCompatibleHighProfileLevel41,
		FragmentableMPEG4AVCH264BDCompatibleHighProfileLevel41,
		MPEG4AVCH264HighProfileLevel42For2DVideo,
		FragmentableMPEG4AVCH264HighProfileLevel42For2DVideo,
		MPEG4AVCH264HighProfileLevel42For3DVideo,
		FragmentableMPEG4AVCH264HighProfileLevel42For3DVideo,
		MPEG4AVCH264StereoHighProfileLevel42,
		FragmentableMPEG4AVCH264StereoHighProfileLevel42,
		HEVCH265MainProfileLevel51,
		HEVCH265Main10ProfileLevel51,
		HTJ2K,
	}
)

// IsUncompressed returns true if the transfer syntax does not use compression.
func IsUncompressed(ts *Syntax) bool {
	return !ts.IsEncapsulated() && !ts.IsDeflate()
}

// IsLosslessCompressed returns true if the transfer syntax uses lossless compression.
func IsLosslessCompressed(ts *Syntax) bool {
	return ts.IsEncapsulated() && !ts.IsLossy()
}

// IsLossyCompressed returns true if the transfer syntax uses lossy compression.
func IsLossyCompressed(ts *Syntax) bool {
	return ts.IsEncapsulated() && ts.IsLossy()
}
