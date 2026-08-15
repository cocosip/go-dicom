# DICOM Reconstruction Package

## Overview

This package provides 3D volume reconstruction and Multi-Planar Reformation (MPR) capabilities for DICOM medical images.

## Current Status

🚧 **UNDER DEVELOPMENT - NOT YET FUNCTIONAL** 🚧

This package is currently a placeholder/framework. Full implementation is pending.

## What is MPR (Multi-Planar Reformation)?

MPR allows viewing a 3D volume from different perspectives by extracting slices along arbitrary planes:

- **Axial**: Standard CT/MR slices (looking from feet to head)
- **Coronal**: Front-to-back slices
- **Sagittal**: Side-to-side slices
- **Oblique**: Any custom plane orientation

## Planned Features

### 1. Volume Reconstruction
- Load multiple 2D DICOM slices
- Validate geometric consistency
- Build 3D volume with proper spacing
- Handle varying slice spacing

### 2. Multi-Planar Reformation (MPR)
- Extract slices along standard planes (Axial, Coronal, Sagittal)
- Support arbitrary oblique planes
- Trilinear interpolation for smooth reconstructions
- Configurable slice spacing and thickness

### 3. DICOM Generation
- Generate new DICOM series from reformatted slices
- Preserve patient/study metadata
- Update geometric tags (Image Position, Orientation)
- Maintain DICOM conformance

## Foundation Status

The shared spatial foundations are now available:

- `pkg/imaging/math3d`: points, vectors, planes, bounds, rectangles, and matrices
- `pkg/imaging/geometry`: classic and enhanced multi-frame geometry, coordinate conversion, and localizers
- `pkg/imaging/interpolation`: reusable scalar nearest-neighbor and bilinear sampling

This package still needs reconstruction-specific pixel access, volume
construction, slice generation, and DICOM output.

## Implementation Roadmap

### Phase 1: Foundation
- [x] Implement 3D geometry library and matrices
- [x] Implement geometric helper functions and interpolation primitives
- [x] Add FrameGeometry for classic and enhanced image planes
- [ ] Create IPixelData interface and implementations

### Phase 2: Volume Reconstruction (0%)
- [ ] Implement ImageData wrapper
- [ ] Implement VolumeData construction
- [ ] Add slice sorting and validation
- [ ] Calculate bounding box
- [ ] Handle slice spacing validation

### Phase 3: MPR (0%)
- [ ] Implement Slice extraction
- [ ] Compose the available bilinear primitives with between-slice interpolation
- [ ] Implement Stack generation (Axial, Coronal, Sagittal)
- [ ] Support arbitrary plane orientations
- [ ] Optimize with parallel processing

### Phase 4: DICOM Generation (0%)
- [ ] Implement DicomGenerator
- [ ] Metadata preservation
- [ ] Geometric tag updates
- [ ] Pixel data encoding
- [ ] UID generation

### Phase 5: Testing & Optimization (0%)
- [ ] Unit tests for each component
- [ ] Integration tests with real DICOM data
- [ ] Performance benchmarks
- [ ] Memory usage optimization
- [ ] Documentation and examples

## Similar Libraries

For reference, see these implementations in other languages:

- **fo-dicom (C#)**: https://github.com/fo-dicom/fo-dicom/tree/development/FO-DICOM.Core/Imaging/Reconstruction
- **dcmtk (C++)**: DCMTK's dcmmkdir and related tools
- **pydicom + numpy**: Python-based MPR implementations

## Contributing

This is a complex feature that requires significant implementation effort. Contributions are welcome, especially for:

1. Volume construction and validation
2. Slice and stack generation
3. DICOM output for reformatted stacks
4. Performance optimization and scanner-specific testing

## License

Microsoft Public License (MS-PL) - Same as parent project

## References

- DICOM Standard Part 3: Information Object Definitions
  - Section C.7.6.2: Image Plane Module
  - Section C.7.6.1: Image Pixel Module
- Multi-Planar Reformation techniques in medical imaging
