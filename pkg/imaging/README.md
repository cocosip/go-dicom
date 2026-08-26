# Imaging Package

This package provides DICOM image processing functionality for the go-dicom library.

## Overview

The imaging package handles Dataset-driven rendering, pixel data, codecs,
patient-space frame geometry, spatial transforms, and interpolation for DICOM
files.

## Components

### Core Types

- **ColorSpace** (`colorspace.go`): Defines color spaces and their components
  - Grayscale, RGB, BGR, RGBA, YCbCr, etc.
  - Component-based representation with subsampling information

- **PhotometricInterpretation** (`photometric.go`): Represents DICOM photometric interpretation
  - MONOCHROME1, MONOCHROME2 (grayscale)
  - PALETTE COLOR (indexed color)
  - RGB (true color)
  - YBR_FULL, YBR_FULL_422, YBR_PARTIAL_422, YBR_PARTIAL_420 (YCbCr variants)
  - YBR_ICT, YBR_RCT (JPEG 2000 color transforms)

- **PixelConfiguration** (`pixel_config.go`): Pixel representation and planar configuration
  - PixelRepresentation: Unsigned vs Signed integers
  - PlanarConfiguration: Interleaved vs Planar color data

### Geometry and Spatial Tools

- **geometry**: Classic and enhanced multi-frame geometry, orientation,
  patient/image coordinate conversion, bounds, and localizers
- **math3d**: Point, vector, plane, segment, bounding-box, rectangle, and 4x4
  matrix primitives
- **transform**: Composable affine transforms and scale/rotate/flip/pan viewer
  state with best-fit placement
- **interpolation**: Stride-aware scalar grids with nearest-neighbor and
  bilinear sampling and resizing
- **Histogram**: Integer-bin counting and explicit or percentage windows

## Geometry and Spatial Usage

### Extract Frame Geometry

Use `geometry.NewFrameGeometry` with a parsed `*dataset.Dataset`. The frame
index is zero-based for both classic and enhanced multi-frame datasets:

```go
frameGeometry, err := geometry.NewFrameGeometry(ds, frameIndex)
if err != nil {
    return err
}

switch frameGeometry.Type {
case geometry.GeometryVolume:
    // Complete patient position, orientation, and spacing are available.
case geometry.GeometryPlane:
    // Pixel spacing is available, but patient position/orientation are not.
case geometry.GeometryNone:
    // Dimensions are available, but spatial transforms cannot be performed.
}
```

For enhanced multi-frame datasets, shared functional-group values are loaded
first and matching per-frame values override them. Top-level spacing follows
fo-dicom precedence: Imager Pixel Spacing, Pixel Spacing, Nominal Scanned Pixel
Spacing, then functional-group spacing.

The complete Dataset construction and coordinate round-trip are compiled as
part of [`ExampleNewFrameGeometry`](geometry/example_test.go).

### Coordinate Convention

Image coordinates are zero-based pixel centers. `(0,0)` is the center of the
first pixel and `(columns-1, rows-1)` is the center of the last pixel. `X`
advances across columns and uses `PixelSpacingColumns`; `Y` advances across
rows and uses `PixelSpacingRows`. DICOM Pixel Spacing values are supplied in
`[row, column]` order.

```go
patientPoint, err := frameGeometry.ImageToPatient(math3d.Point2{X: 12, Y: 8})
if err != nil {
    return err
}
imagePoint, err := frameGeometry.PatientToImage(patientPoint)
if err != nil {
    return err
}
```

`TopLeft`, `TopRight`, `BottomLeft`, `BottomRight`, and `BoundingBox()` describe
pixel centers, not the outer physical edges of the pixels. Require
`GeometryVolume` when the result must represent a real DICOM patient-space
location.

### Localizers

Localizers require source and destination frames from the same Frame of
Reference and with different orientations. Check compatibility before drawing:

```go
if geometry.CanDrawLocalizer(source, destination) {
    start, end, ok := geometry.IntersectionLocalizer(source, destination)
    if ok {
        drawLine(start, end)
    }
}
```

`IntersectionLocalizer` returns the clipped intersection line in destination
image coordinates. `ProjectionLocalizer` projects all four source corners into
destination image coordinates. See the executable
[`ExampleIntersectionLocalizer`](geometry/example_test.go).

### Affine and Viewer Transforms

`Affine2D.Then` composes in execution order: `a.Then(b)` applies `a` first and
`b` second.

```go
matrix := transform.Identity().
    Then(transform.Scale(2, 2)).
    Then(transform.Rotate(90)).
    Then(transform.Translate(100, 50))
screenPoint := matrix.Apply(imagePoint)
```

Use `SpatialTransform` for viewer state and `BestFit` to center an image while
preserving its aspect ratio. Invalid or empty rectangles and non-positive
scales return errors. The executable examples cover
[`Affine2D.Then`](transform/example_test.go),
[`SpatialTransform.Affine`](transform/example_test.go), and
[`BestFit`](transform/example_test.go).

### Interpolation

`interpolation.Grid` is a read-only view over row-major `float64` samples.
Stride is measured in values, not bytes, and can be larger than the image
width. Sampling uses the same zero-based pixel-center convention as geometry:

```go
grid, err := interpolation.NewGrid(samples, width, height, stride)
if err != nil {
    return err
}
value, inside := grid.Bilinear(x, y)
resized, err := grid.Resize(outputWidth, outputHeight, interpolation.ModeBilinear)
```

`Nearest` and `Bilinear` return `inside=false` for non-finite or out-of-bounds
coordinates. `Resize` returns tightly packed output and aligns source and
destination endpoints. See the stride-aware executable
[`ExampleGrid.Resize`](interpolation/example_test.go).

### Histogram Windows

`Histogram` counts integer values in an inclusive range. Values outside that
range are ignored. A percentage window removes less-populated edge bins while
retaining the requested share of the total count; an explicit window selects
an inclusive absolute bin range:

```go
histogram, err := imaging.NewHistogram(minimum, maximum)
if err != nil {
    return err
}
for _, value := range samples {
    histogram.Add(value)
}
if err := histogram.ApplyPercentWindow(95); err != nil {
    return err
}
start, end := histogram.WindowStart(), histogram.WindowEnd()
```

See executable [`ExampleHistogram.ApplyPercentWindow`](histogram_example_test.go).

### Errors and Reconstruction Boundary

`NewFrameGeometry` returns errors for invalid frame indices, dimensions,
malformed values, incomplete position/orientation pairs, and invalid direction
cosines. Missing spacing is not an error; it produces `GeometryNone`, and
coordinate conversion then returns an error. Constructors and resizing methods
also reject non-finite or invalid dimensions instead of silently repairing
them.

These APIs provide the spatial and sampling foundation for IMG-003. Volume
assembly, slice sorting, trilinear MPR sampling, and derived DICOM generation
remain in the `reconstruction` package and are not implemented by IMG-002.

### Codec Framework

The `codec` subpackage (`pkg/imaging/codec/`) provides codec interfaces and native pixel-data handling. Compressed codecs are supplied by `go-dicom-codecs`.

- **Codec Interface**: Generic interface for image codecs
- **PixelData**: Lightweight pixel data structure for codec operations
- **Parameters**: Codec-specific parameters interface

### DicomPixelData

**Implementation**: `pixeldata.go`

The `DicomPixelData` type provides a high-level interface for managing DICOM pixel data:

- **PixelDataInfo**: Metadata container for pixel data attributes
  - Image dimensions (width, height)
  - Bit depth (bits allocated/stored, high bit)
  - Multi-frame support
  - Photometric interpretation
  - Validation and size calculation helpers

- **DicomPixelData**: Main pixel data management class
  - Frame-based storage
  - Add/Get individual frames
  - Codec integration (Encode/Decode methods)
  - Conversion to/from codec.PixelData

**Usage Example**:
```go
import "github.com/cocosip/go-dicom/pkg/imaging"

// Create pixel data info
info := &imaging.PixelDataInfo{
    Width:                     512,
    Height:                    512,
    NumberOfFrames:            1,
    BitsAllocated:             16,
    BitsStored:                12,
    HighBit:                   11,
    SamplesPerPixel:           1,
    PixelRepresentation:       imaging.UnsignedPixels,
    PlanarConfiguration:       imaging.InterleavedPlanar,
    PhotometricInterpretation: imaging.Monochrome2,
}

// Create pixel data
pixelData, err := imaging.NewDicomPixelData(info)

// Add frames
frameBytes := make([]byte, info.UncompressedFrameSize())
// ... fill frameBytes ...
err = pixelData.AddFrame(frameBytes)

```

### Implemented Codecs

#### Compressed Codecs

**Implementation**: `go-dicom-codecs`

Compressed DICOM image codecs are implemented by `go-dicom-codecs`.

**Features**:
- ✅ Encode pixel data to compressed format
- ✅ Decode compressed pixel data
- ✅ Support for 8-bit and 16-bit images
- ✅ Support for grayscale and RGB images
- ✅ Support for interleaved and planar configurations
- ✅ Multi-frame support

#### Native/Uncompressed

**Implementation**: `codec/native.go`

The Native codec handles uncompressed pixel data with various byte orders:

**Features**:
- ✅ Little Endian / Big Endian support
- ✅ Byte swapping for multi-byte samples (16-bit, 32-bit, 64-bit)
- ✅ Direct copy for single-byte samples
- ✅ Read/Write helpers for uint16/uint32 with endianness
- ✅ Utility function for endianness conversion

**Usage Example**:
```go
import "github.com/cocosip/go-dicom/pkg/imaging/codec"

// Little Endian codec
leCodec := codec.NewExplicitVRLittleEndianCodec()

// Big Endian codec
beCodec := codec.NewExplicitVRBigEndianCodec()

// Convert between endianness
swapped, err := codec.ConvertEndianness(pixelBytes, 2) // 2 bytes per sample

// Byte swapping during encode/decode
params := codec.NewBaseParameters()
params.SetParameter("swap_bytes", true)
err = leCodec.Encode(src, dst, params)
```

### Compressed Codec Registration

Compressed codecs live in the companion `go-dicom-codecs` module so core users
only link the codecs their application needs. Register a codec with a blank
import before decoding or transcoding its transfer syntax:

```go
import _ "github.com/cocosip/go-dicom-codecs/jpeg/baseline"
```

If no matching codec was registered, decode/transcode returns the registry
lookup error; registering a codec does not imply that a peer accepted its
transfer syntax during association negotiation.

## Architecture Notes

### Package Design

The imaging package is designed to avoid circular dependencies:

1. **Core types** (`pkg/imaging/*`): Basic types with no heavy dependencies
2. **Codec package** (`pkg/imaging/codec/`): Codec interfaces and implementations
3. **Minimal dependencies**: Only depends on `pkg/dicom/transfer` for transfer syntax references

### PixelData Structure

The codec `PixelData` type is intentionally simplified to avoid circular dependencies with the full `DicomDataset`. It contains only the essential information needed for encoding/decoding operations.

In a full implementation, you would:
1. Extract pixel information from a `DicomDataset` into a `codec.PixelData`
2. Perform codec operations
3. Update the dataset with the processed pixel data

## Testing

All components include comprehensive unit tests:

```bash
# Test all imaging components
go test ./pkg/imaging/...

# Test specific codec
go test ./pkg/imaging/codec -v

# Test with coverage
go test ./pkg/imaging/... -cover
```

Current test coverage:
- ✅ ColorSpace: 100%
- ✅ PhotometricInterpretation: 100%
- ✅ PixelConfiguration: 100%
- ✅ PixelDataInfo/DicomPixelData: 100%
- ✅ Native Codec: 100%

## Future Work

### Short Term
1. Build `reconstruction.ImageData` and volume validation on the completed
   geometry and interpolation APIs
2. Implement volume slice sorting and trilinear sampling for MPR
3. Generate DICOM datasets from reformatted stacks

### Medium Term
1. Add further rendering performance benchmarks

### Long Term
1. GPU-accelerated processing

## References

- DICOM Standard Part 5: Data Structures and Encoding
- DICOM Standard Part 3: Information Object Definitions
  - Annex C: Image Pixel Module
- fo-dicom C# library: Reference implementation

## License

Copyright (c) 2025 go-dicom contributors.
Licensed under the Microsoft Public License (MS-PL).
