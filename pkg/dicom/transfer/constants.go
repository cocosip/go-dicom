// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transfer

import (
	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

// Standard DICOM Transfer Syntaxes
var (
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

	// JPEGBaseline8Bit is JPEG Baseline (Process 1).
	// Lossy JPEG compression, 8-bit.
	JPEGBaseline8Bit = NewBuilder(uid.JPEGBaseline8Bit).
				SetExplicitVR(true).
				SetEncapsulated(true).
				SetLossy(true, "ISO_10918_1").
				SetEndian(endian.Little).
				Build()

	// JPEGExtended12Bit is JPEG Extended (Process 2 & 4).
	// Lossy JPEG compression, 12-bit.
	JPEGExtended12Bit = NewBuilder(uid.JPEGExtended12Bit).
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

	// JPEGLosslessSV1 is JPEG Lossless, Non-Hierarchical, First-Order Prediction.
	// JPEG Lossless, Process 14 [Selection Value 1].
	JPEGLosslessSV1 = NewBuilder(uid.JPEGLosslessSV1).
			SetExplicitVR(true).
			SetEncapsulated(true).
			SetEndian(endian.Little).
			Build()

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

	// RLELossless is RLE Lossless compression.
	RLELossless = NewBuilder(uid.RLELossless).
			SetExplicitVR(true).
			SetEncapsulated(true).
			SetEndian(endian.Little).
			Build()

	// GEPrivateImplicitVRBigEndian is a private GE transfer syntax.
	// Same as Implicit VR Little Endian except for big endian pixel data.
	GEPrivateImplicitVRBigEndian = NewBuilder(uid.GEPrivateImplicitVRBigEndian).
					SetExplicitVR(false).
					SetEndian(endian.Little).
					SetSwapPixelData(true).
					Build()
)

// Common transfer syntax groups for convenience
var (
	// UncompressedTransferSyntaxes are transfer syntaxes without compression.
	UncompressedTransferSyntaxes = []*TransferSyntax{
		ImplicitVRLittleEndian,
		ExplicitVRLittleEndian,
		ExplicitVRBigEndian,
	}

	// LosslessTransferSyntaxes are transfer syntaxes with lossless compression.
	LosslessTransferSyntaxes = []*TransferSyntax{
		JPEGLossless,
		JPEGLosslessSV1,
		JPEGLSLossless,
		JPEG2000Lossless,
		RLELossless,
	}

	// LossyTransferSyntaxes are transfer syntaxes with lossy compression.
	LossyTransferSyntaxes = []*TransferSyntax{
		JPEGBaseline8Bit,
		JPEGExtended12Bit,
		JPEGLSNearLossless,
		JPEG2000,
	}
)

// IsUncompressed returns true if the transfer syntax does not use compression.
func IsUncompressed(ts *TransferSyntax) bool {
	return !ts.IsEncapsulated() && !ts.IsDeflate()
}

// IsLosslessCompressed returns true if the transfer syntax uses lossless compression.
func IsLosslessCompressed(ts *TransferSyntax) bool {
	return ts.IsEncapsulated() && !ts.IsLossy()
}

// IsLossyCompressed returns true if the transfer syntax uses lossy compression.
func IsLossyCompressed(ts *TransferSyntax) bool {
	return ts.IsEncapsulated() && ts.IsLossy()
}
