# Imaging Package

This package provides DICOM image processing functionality for the go-dicom library.

## Overview

The imaging package handles pixel data representation, image codecs, and basic image processing operations for DICOM files.

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

### Codec Framework

The `codec` subpackage (`pkg/imaging/codec/`) provides image compression and decompression:

- **Codec Interface**: Generic interface for image codecs
- **PixelData**: Lightweight pixel data structure for codec operations
- **Parameters**: Codec-specific parameters interface

### Implemented Codecs

#### RLE Lossless (Run-Length Encoding)

**Implementation**: `codec/rle.go`

The RLE codec implements DICOM RLE Lossless compression (Transfer Syntax UID: 1.2.840.10008.1.2.5) according to DICOM Part 5, Annex G.

**Features**:
- ✅ Encode pixel data to RLE-compressed format
- ✅ Decode RLE-compressed pixel data
- ✅ Support for 8-bit and 16-bit images
- ✅ Support for grayscale and RGB images
- ✅ Support for interleaved and planar configurations
- ✅ Multi-frame support

**Usage Example**:
```go
import "github.com/cocosip/go-dicom/pkg/imaging/codec"

// Create codec
rleCodec := codec.NewRLECodec()

// Prepare source pixel data
src := &codec.PixelData{
    Data:                      pixelBytes,
    Width:                     512,
    Height:                    512,
    NumberOfFrames:            1,
    BitsAllocated:             8,
    BitsStored:                8,
    HighBit:                   7,
    SamplesPerPixel:           1,
    PixelRepresentation:       0, // unsigned
    PlanarConfiguration:       0, // interleaved
    PhotometricInterpretation: "MONOCHROME2",
}

// Encode
encoded := &codec.PixelData{ /* same metadata as src */ }
err := rleCodec.Encode(src, encoded, nil)

// Decode
decoded := &codec.PixelData{ /* same metadata as src */ }
err = rleCodec.Decode(encoded, decoded, nil)
```

### Deferred Codecs

The following codecs are **not yet implemented** due to lack of suitable Go libraries:

- ❌ **JPEG Baseline** (Transfer Syntax UID: 1.2.840.10008.1.2.4.50)
  - Reason: DICOM JPEG uses non-standard markers not compatible with Go's standard `image/jpeg`
  - Future: May implement via CGO bindings to libjpeg or wait for Go library support

- ❌ **JPEG Lossless** (Transfer Syntax UID: 1.2.840.10008.1.2.4.57, .70)
  - Reason: No mature Go library for JPEG lossless
  - Future: Requires specialized JPEG lossless decoder

- ❌ **JPEG-LS** (Transfer Syntax UID: 1.2.840.10008.1.2.4.80, .81)
  - Reason: No mature Go library for JPEG-LS
  - Future: Wait for library or implement from spec

- ❌ **JPEG 2000** (Transfer Syntax UID: 1.2.840.10008.1.2.4.90, .91)
  - Reason: No mature Go library for JPEG 2000
  - Future: Consider CGO bindings to OpenJPEG

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
- ✅ RLE Codec: 100%

## Future Work

### Short Term
1. Integrate codec with DicomDataset for full read/write operations
2. Add DicomPixelData wrapper class for easier dataset integration
3. Implement palette color LUT support

### Medium Term
1. Research and evaluate JPEG codec options:
   - Pure Go implementation
   - CGO bindings to libjpeg/libjpeg-turbo
2. Implement basic image rendering (VOI LUT, modality LUT)
3. Add image export to standard formats (PNG, JPEG)

### Long Term
1. JPEG-LS codec (pending library availability)
2. JPEG 2000 codec (likely via CGO)
3. Advanced rendering features
4. GPU-accelerated processing

## References

- DICOM Standard Part 5: Data Structures and Encoding
  - Annex G: RLE Lossless Compression
- DICOM Standard Part 3: Information Object Definitions
  - Annex C: Image Pixel Module
- fo-dicom C# library: Reference implementation

## License

Copyright (c) 2025 go-dicom contributors.
Licensed under the Microsoft Public License (MS-PL).
