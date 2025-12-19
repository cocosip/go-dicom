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
- **rle**: RLE Lossless

### Lossy Compression
- **jpeg** / **jpeg-baseline**: JPEG Baseline (Process 1) - 8-bit lossy
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
  - Options: uncompressed, jpeg, jpeg-lossless, jpeg-ls, jpeg2000, jpeg2000-lossless, rle
- `-jpeg-quality`: JPEG quality for lossy compression, 1-100 (default: 90)
- `-verbose`: Show verbose output (default: true)

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
transcode -input input.dcm -ts rle -output D:\output\compressed.dcm
```

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
    _ "github.com/cocosip/go-dicom-codec/jpeg"     // JPEG codecs
    _ "github.com/cocosip/go-dicom-codec/jpegls"   // JPEG-LS codecs
    _ "github.com/cocosip/go-dicom-codec/jpeg2000" // JPEG 2000 codecs
    _ "github.com/cocosip/go-dicom-codec/rle"      // RLE codec
)
```

The blank imports automatically register the codecs with the global registry when the packages are imported.

**Current Status**: The current implementation provides the transcoding infrastructure. Full codec support will be available once the go-dicom-codec package implementations are complete.

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
- **JPEG Baseline (Q=90)**: ~80 KB (85% reduction, lossy)

*Actual results vary based on image content and characteristics.*

## Limitations

Current implementation limitations:

1. **Codec Availability**: Some transfer syntaxes require external codec libraries
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
