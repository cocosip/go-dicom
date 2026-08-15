// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package reconstruction provides validated CT/MR volume reconstruction and
// multi-planar reformation in DICOM patient coordinates.
//
// Source datasets are represented as immutable ImageData frames. VolumeData
// sorts compatible frames without expanding them into a dense float64 volume.
// Sample and Cut perform bilinear in-frame and actual-distance between-frame
// interpolation on demand. Slice.Valid distinguishes out-of-volume and padding
// samples from genuine zero values.
//
// Stack lazily describes standard axial, coronal, or sagittal series.
// DicomGenerator streams classic single-frame CT/MR derived datasets with new
// UIDs, source references, image-plane geometry, native 16-bit pixels, and
// Explicit VR Little Endian transfer syntax.
//
// Enhanced CT/MR is supported as input when frames belong to one Stack ID.
// Non-spatial Dimension Index values must remain constant. Enhanced multi-frame
// output and varying temporal/diffusion/cardiac dimensions are outside this
// package's current scope.
package reconstruction
