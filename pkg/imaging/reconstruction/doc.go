// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

// Package reconstruction provides 3D volume reconstruction and Multi-Planar Reformation (MPR) capabilities.
//
// This package enables:
//   - Building 3D volumes from multiple 2D DICOM slices
//   - Generating arbitrary plane cuts (MPR) from reconstructed volumes
//   - Creating Axial, Coronal, and Sagittal reformatted series
//   - Interpolation between slices for smooth reconstructions
//
// # Available Foundations
//
// The shared prerequisites are available in:
//   - pkg/imaging/math3d for point, vector, plane, bounds, and matrix operations
//   - pkg/imaging/geometry for classic and enhanced multi-frame geometry
//   - pkg/imaging/interpolation for scalar nearest-neighbor and bilinear sampling
//
// Reconstruction still needs pixel-access integration, volume construction,
// slice generation, DICOM output, and their end-to-end tests.
//
// # Basic Usage (Planned)
//
//	// Load DICOM files
//	images := []reconstruction.ImageData{...}
//
//	// Build 3D volume
//	volume := reconstruction.NewVolumeData(images)
//
//	// Create reformatted stack (e.g., sagittal view)
//	stack := reconstruction.NewStack(volume, reconstruction.StackTypeSagittal, spacing, sliceDistance)
//
//	// Generate DICOM files from stack
//	generator := reconstruction.NewDicomGenerator(volume.CommonData)
//	datasets := generator.StoreAsDicom(stack, "Sagittal Reformation")
//
// # Reference
//
// Based on fo-dicom Imaging.Reconstruction package.
// Reference: https://github.com/fo-dicom/fo-dicom/tree/development/FO-DICOM.Core/Imaging/Reconstruction
//
// # TODO
//
// This package is currently a placeholder. Full implementation requires:
//  1. Implement pixel access for reconstruction source frames
//  2. Implement volume reconstruction algorithms
//  3. Implement MPR (Multi-Planar Reformation) with trilinear interpolation
//  4. Implement DICOM output for generated stacks
//  5. Add performance benchmarks and end-to-end orientation tests
package reconstruction
