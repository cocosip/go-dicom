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

## Dependencies Required

Before this package can be fully implemented, the following components need to be developed:

### 1. 3D Geometry Library (`pkg/imaging/math3d`)
```go
type Point3D struct {
    X, Y, Z float64
}

type Vector3D struct {
    X, Y, Z float64
}

type Plane3D struct {
    Normal Vector3D
    Point  Point3D
}
```

### 2. Frame Geometry (`pkg/imaging/geometry`)
```go
type FrameGeometry struct {
    ImagePositionPatient    Point3D
    ImageOrientationPatient [6]float64
    PixelSpacing           [2]float64
    SliceThickness         float64
}
```

### 3. Pixel Data Interface (`pkg/imaging`)
```go
type IPixelData interface {
    Width() int
    Height() int
    GetPixel(x, y int) float64
}
```

## Implementation Roadmap

### Phase 1: Foundation (0%)
- [ ] Implement 3D geometry library (Point3D, Vector3D, Plane3D, Matrix3D)
- [ ] Implement geometric helper functions (dot product, cross product, interpolation)
- [ ] Add FrameGeometry for image plane definitions
- [ ] Create IPixelData interface and implementations

### Phase 2: Volume Reconstruction (0%)
- [ ] Implement ImageData wrapper
- [ ] Implement VolumeData construction
- [ ] Add slice sorting and validation
- [ ] Calculate bounding box
- [ ] Handle slice spacing validation

### Phase 3: MPR (0%)
- [ ] Implement Slice extraction
- [ ] Add trilinear interpolation
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

1. 3D geometry math library
2. Interpolation algorithms
3. Performance optimization
4. Testing with various scanner types

## License

Microsoft Public License (MS-PL) - Same as parent project

## References

- DICOM Standard Part 3: Information Object Definitions
  - Section C.7.6.2: Image Plane Module
  - Section C.7.6.1: Image Pixel Module
- Multi-Planar Reformation techniques in medical imaging
