# DICOM Transcode Example

This example demonstrates how to transcode DICOM files between different transfer syntaxes using the go-dicom library's codec package.

## What is DICOM Transcoding?

DICOM transcoding is the process of converting pixel data from one transfer syntax (encoding format) to another. This can include:

- **Compression**: Converting uncompressed pixel data to a compressed format (JPEG, JPEG 2000, RLE, etc.)
- **Decompression**: Converting compressed pixel data back to uncompressed format
- **Format Conversion**: Converting between different compression formats (e.g., JPEG to JPEG 2000)

## Features

This example supports the following transfer syntaxes:

### Uncompressed Formats
- **uncompressed** / **explicit-vr**: Explicit VR Little Endian (1.2.840.10008.1.2.1)
- **implicit-vr**: Implicit VR Little Endian (1.2.840.10008.1.2)

### Lossless Compression
- **jpeg-lossless**: JPEG Lossless, Non-Hierarchical (Process 14)
- **jpeg-ls**: JPEG-LS Lossless
- **jpeg2000-lossless**: JPEG 2000 Lossless
- **rle**: RLE (Run-Length Encoding) Lossless

### Lossy Compression
- **jpeg** / **jpeg-baseline**: JPEG Baseline (Process 1) - 8-bit lossy
- **jpeg-extended** / **jpeg-12bit**: JPEG Extended (Process 2 & 4) - up to 12-bit lossy
- **jpeg-ls-lossy**: JPEG-LS Near-Lossless
- **jpeg2000** / **j2k**: JPEG 2000 with lossy options

## Building

This example uses a separate Go module with the `go-dicom-codec` dependency, which provides the actual codec implementations for compression and decompression.

```bash
# From the examples/transcode directory
cd examples/transcode
go mod download  # Download dependencies
go build         # Build the executable

# Or build from the project root
go build -o bin/transcode.exe ./examples/transcode
```

**Note**: This example requires the [go-dicom-codec](https://github.com/cocosip/go-dicom-codec) package, which is automatically downloaded when you build.

## Usage

### Basic Usage

```bash
transcode -input <input-file> -ts <transfer-syntax> [-output <output-file>]
```

### Command-Line Flags

- `-input` (required): Input DICOM file path
- `-output`: Output DICOM file path (default: input filename with "_transcoded" suffix)
- `-ts`: Target transfer syntax (default: "jpeg")
  - Options: uncompressed, jpeg, jpeg-extended, jpeg-lossless, jpeg-ls, jpeg2000, jpeg2000-lossless, rle
- `-jpeg-quality`: JPEG quality for lossy compression, 1-100 (default: 90)
- `-verbose`: Show verbose output (default: true)
- `-all`: Generate output files for all compatible formats (ignores `-ts` and `-output`)

### Examples

#### 1. Compress to JPEG Baseline (Lossy)

Convert an uncompressed DICOM file to JPEG with 85% quality:

```bash
transcode -input uncompressed.dcm -ts jpeg -jpeg-quality 85
```

This is useful for:
- Reducing file size for storage or transmission
- Creating web-friendly DICOM images
- Archive storage where some quality loss is acceptable

#### 2. Compress to JPEG Lossless

Convert to lossless JPEG compression:

```bash
transcode -input uncompressed.dcm -ts jpeg-lossless
```

Benefits:
- No quality loss
- Typically 50-60% size reduction
- Maintains diagnostic quality
- Widely supported

#### 3. Compress to RLE Lossless

Convert to RLE (Run-Length Encoding) lossless compression:

```bash
transcode -input uncompressed.dcm -ts rle
```

Best for:
- Images with large uniform areas
- Screen captures or synthetic images
- Simple and fast compression/decompression

#### 4. Decompress to Uncompressed

Convert any compressed DICOM file back to uncompressed:

```bash
transcode -input compressed.dcm -ts uncompressed
```

Useful for:
- Processing with tools that don't support compressed formats
- Editing pixel data
- Maximum compatibility

#### 5. Convert Between Compression Formats

Convert from one compressed format to another (e.g., JPEG to JPEG 2000):

```bash
transcode -input jpeg_compressed.dcm -ts jpeg2000-lossless
```

Note: The transcoder will decompress the source format and then compress to the target format.

#### 6. Specify Custom Output Path

```bash
transcode -input input.dcm -ts jpeg2000-lossless -output D:\output\compressed.dcm
```

#### 7. Generate All Compatible Formats (Recommended for Comparison)

The most powerful feature - automatically generate output files for all compatible compression formats and compare results:

```bash
transcode -input myimage.dcm -all
```

**Note**: Output files are automatically named based on the input filename with a format suffix:
- Input: `myimage.dcm`
- Outputs: `myimage_uncompressed.dcm`, `myimage_rle.dcm`, `myimage_jpeg_lossless.dcm`, etc.

**Output Example**:
```
=== Generating All Compatible Formats ===

Found 6 compatible formats for 16-bit image:

[1/6] Processing Uncompressed (Explicit VR)...
  ✅ Success: myimage_uncompressed.dcm (819572 bytes, -0.0% reduction)
[2/6] Processing RLE Lossless...
  ✅ Success: myimage_rle.dcm (417130 bytes, 49.1% reduction)
[3/6] Processing JPEG Lossless...
  ✅ Success: myimage_jpeg_lossless.dcm (239646 bytes, 70.7% reduction)
[4/6] Processing JPEG-LS Lossless...
  ✅ Success: myimage_jpegls_lossless.dcm (198471 bytes, 75.8% reduction)
[5/6] Processing JPEG 2000 Lossless...
  ✅ Success: myimage_jpeg2000_lossless.dcm (218118 bytes, 73.4% reduction)
[6/6] Processing JPEG 2000 Lossy...
  ✅ Success: myimage_jpeg2000_lossy.dcm (113272 bytes, 86.2% reduction)

=== Compression Comparison ===

Input File: myimage.dcm

Format                         Output File                     Size       Compression
------------------------------------------------------------------------------------
Uncompressed (Explicit VR)     myimage_uncompressed.dcm        800.4 KB   0.0%↑
RLE Lossless                   myimage_rle.dcm                 407.4 KB   1.96:1 (49.1%↓)
JPEG Lossless                  myimage_jpeg_lossless.dcm       234.0 KB   3.42:1 (70.7%↓)
JPEG-LS Lossless               myimage_jpegls_lossless.dcm     193.8 KB   4.13:1 (75.8%↓)
JPEG 2000 Lossless             myimage_jpeg2000_lossless.dcm   213.0 KB   3.76:1 (73.4%↓)
JPEG 2000 Lossy                myimage_jpeg2000_lossy.dcm      110.6 KB   7.23:1 (86.2%↓)
------------------------------------------------------------------------------------

Best Compression: JPEG 2000 Lossy (7.23:1, 86.2% reduction)
Input Size:       800.1 KB
Smallest Output:  110.6 KB
```

**Benefits**:
- ✅ Automatically detects compatible formats based on image bit depth
- ✅ Generates all formats in one command
- ✅ Shows side-by-side compression comparison
- ✅ Identifies the best compression format
- ✅ Perfect for benchmarking and choosing the optimal format

**Smart Format Selection**:
- For 8-bit images: Generates 8 formats (includes JPEG Baseline & Extended)
- For 10/12-bit images: Generates 7 formats (includes JPEG Extended)
- For 16-bit images: Generates 6 formats (excludes JPEG Baseline & Extended)

## Understanding Transfer Syntaxes

### When to Use Each Format

| Format | Compression Ratio | Quality | Speed | Use Case |
|--------|------------------|---------|-------|----------|
| **Uncompressed** | 1:1 (none) | Perfect | Fastest | Maximum compatibility, editing |
| **RLE Lossless** | 1.2-2:1 | Perfect | Very Fast | Simple images, screen captures |
| **JPEG Lossless** | 2-3:1 | Perfect | Fast | General purpose, diagnostic images |
| **JPEG-LS Lossless** | 2-3:1 | Perfect | Fast | Efficient lossless compression |
| **JPEG 2000 Lossless** | 2-4:1 | Perfect | Moderate | High compression, modern format |
| **JPEG Baseline** | 5-20:1 | Good* | Very Fast | Non-diagnostic, preview images |
| **JPEG-LS Near-Lossless** | 3-6:1 | Very Good* | Fast | Controlled quality loss |
| **JPEG 2000** | 5-50:1 | Good-Excellent* | Moderate | Flexible quality/size trade-off |

\* Quality depends on compression settings

### Lossless vs Lossy

**Lossless Compression:**
- No information is lost
- Original pixel values can be perfectly reconstructed
- Required for diagnostic images in many jurisdictions
- Lower compression ratios

**Lossy Compression:**
- Some information is discarded to achieve higher compression
- Cannot perfectly reconstruct original pixel values
- Suitable for preview images, educational content, or non-diagnostic use
- Much higher compression ratios possible

## Implementation Details

This example uses the `codec.Transcoder` type from the `pkg/imaging/codec` package, which:

1. Detects the input transfer syntax from the DICOM file
2. Selects appropriate codecs for decompression and compression
3. Handles pixel data conversion, including:
   - Decompression of encapsulated pixel data
   - Frame-by-frame processing
   - Re-compression to target format
   - Proper metadata updates

The transcoder automatically handles:
- Single-frame and multi-frame images
- Different bit depths (8-bit, 12-bit, 16-bit)
- Color and grayscale images
- Various photometric interpretations

### Codec Registration

To use compression/decompression, you need to register codecs from the [go-dicom-codec](https://github.com/cocosip/go-dicom-codec) package:

```go
import (
    _ "github.com/cocosip/go-dicom-codec/jpeg/baseline"        // JPEG Baseline
    _ "github.com/cocosip/go-dicom-codec/jpeg/lossless"        // JPEG Lossless
    _ "github.com/cocosip/go-dicom-codec/jpegls/lossless"      // JPEG-LS Lossless
    _ "github.com/cocosip/go-dicom-codec/jpegls/nearlossless"  // JPEG-LS Near-Lossless
    _ "github.com/cocosip/go-dicom-codec/jpeg2000/lossless"    // JPEG 2000 Lossless
    _ "github.com/cocosip/go-dicom-codec/jpeg2000/lossy"       // JPEG 2000 Lossy
)
```

The blank imports automatically register the codecs with the global registry when the packages are imported.

**Current Status**: The go-dicom-codec package is production-ready for most codecs:
- ✅ JPEG Baseline, Extended, and Lossless - Fully functional
- ✅ JPEG-LS Lossless and Near-Lossless - Fully functional
- ✅ JPEG 2000 Lossless and Lossy - Fully functional
- 🧪 HTJ2K (High-Throughput JPEG 2000) - Experimental

## Code Overview

```go
// Create transcoder
transcoder := codec.NewTranscoder(
    inputTransferSyntax,
    outputTransferSyntax,
    codec.WithOutputParameters(params), // Optional codec parameters
)

// Transcode the dataset
newDataset, err := transcoder.Transcode(dataset)
if err != nil {
    log.Fatal(err)
}

// Write output file with new transfer syntax
writer := writer.NewWriter(outputFile, outputTransferSyntax)
err = writer.WriteDataset(newDataset)
```

## Performance Considerations

### Compression Speed (Fastest to Slowest)
1. RLE Lossless
2. JPEG Baseline (lossy)
3. JPEG Lossless
4. JPEG-LS
5. JPEG 2000

### Compression Ratio (Best to Worst)
1. JPEG 2000 Lossless
2. JPEG-LS Lossless
3. JPEG Lossless
4. RLE Lossless
5. Uncompressed

### File Size Examples

For a typical uncompressed CT image (512×512×16-bit, ~524 KB):

- **RLE Lossless**: ~400 KB (23% reduction)
- **JPEG Lossless**: ~350 KB (33% reduction)
- **JPEG-LS Lossless**: ~320 KB (39% reduction)
- **JPEG 2000 Lossless**: ~280 KB (47% reduction)
- **JPEG Baseline (Q=90)**: ~80 KB (85% reduction, lossy)*

\* Note: JPEG Baseline only supports 8-bit images

*Actual results vary based on image content and characteristics.*

## Limitations

### Bit Depth Constraints

Different compression formats have specific bit depth requirements:

| Format | Supported Bit Depths | Notes |
|--------|---------------------|-------|
| **JPEG Baseline** | 8-bit only | Most restrictive, lossy compression |
| **JPEG Extended** | 8-bit, 12-bit | Lossy compression for higher bit depths |
| **JPEG Lossless** | 8-bit, 10-bit, 12-bit, 16-bit | Recommended for medical images |
| **JPEG-LS** | 8-bit, 10-bit, 12-bit, 16-bit | Efficient lossless compression |
| **JPEG 2000** | 8-bit, 10-bit, 12-bit, 16-bit | Modern standard, flexible |
| **RLE** | All bit depths | Simple, reliable |
| **Uncompressed** | All bit depths | No restrictions |

**Smart Format Selection**: The transcoder automatically detects bit depth incompatibilities and switches to a compatible format while preserving the user's intent (lossy vs lossless).

**Example - Auto-switching**:
```bash
$ transcode -input 16bit_image.dcm -ts jpeg

Compatibility Error: JPEG Baseline only supports 8-bit images.
  Your image has 16 bits allocated.

  Auto-switching to compatible format...
  → Switching to JPEG 2000 Lossy (supports all bit depths)

Auto-switching to compatible format: JPEG 2000 Image Compression
Creating transcoder...
Transcoding dataset...
Transcoding complete!
```

**Auto-switching Logic**:

| Requested Format | Image Bit Depth | Auto-switched To | Reason |
|-----------------|----------------|------------------|---------|
| JPEG Baseline | 10/12/16-bit | JPEG Extended (≤12-bit) or JPEG 2000 Lossy (>12-bit) | Preserves lossy intent |
| JPEG Extended | 16-bit | JPEG 2000 Lossy | Supports all bit depths |
| Any incompatible | Any | JPEG Lossless | Safe lossless default |

### Other Limitations

1. **Codec Availability**: Some transfer syntaxes require codec libraries
2. **Multi-frame Support**: Complex multi-frame sequences may have limitations
3. **Color Space Conversion**: Some conversions may not preserve exact color representation
4. **Overlay Data**: Embedded overlays are preserved but not converted
5. **Metadata**: Some compression-specific metadata may need manual adjustment

## Error Handling

The example handles common errors:

- **Missing Input File**: Validates file exists before processing
- **Unsupported Transfer Syntax**: Checks if codec is available
- **Insufficient Memory**: Large images may fail on systems with limited RAM
- **Invalid Pixel Data**: Corrupted or malformed pixel data will cause transcode to fail
- **Write Errors**: Disk full or permission issues when saving output

## Related Examples

- `examples/read_dicom`: Reading DICOM files
- `examples/write_dicom`: Writing DICOM files
- `examples/dicom_to_image`: Converting DICOM to standard image formats

## References

- DICOM Standard Part 5 (Transfer Syntaxes): https://dicom.nema.org/medical/dicom/current/output/chtml/part05/chapter_10.html
- JPEG Compression: ISO/IEC 10918-1
- JPEG 2000: ISO/IEC 15444-1
- JPEG-LS: ISO/IEC 14495-1
