# DICOM to Image Converter

This example demonstrates how to extract pixel data from DICOM files and convert them to standard image formats (PNG, JPEG) using the `imaging` package.

## Features

- Extract pixel data from DICOM files
- Support for grayscale images (MONOCHROME1, MONOCHROME2)
- Support for RGB color images
- Apply window center/width for optimal display
- Auto-detect windowing from DICOM tags
- Support for multi-frame DICOM files
- Export as PNG or JPEG format
- Grayscale inversion option

## Usage

### Basic Usage

Convert a DICOM file to PNG (uses auto-detected settings):

```bash
go run main.go -input sample.dcm
```

This will create `sample.png` in the same directory.

### Specify Output File

```bash
go run main.go -input input.dcm -output output.png
```

### Custom Windowing

For CT scans, you may want to adjust the window center and width:

```bash
# Lung window
go run main.go -input ct_chest.dcm -window-center -600 -window-width 1500

# Bone window
go run main.go -input ct_chest.dcm -window-center 400 -window-width 1800

# Soft tissue window
go run main.go -input ct_chest.dcm -window-center 40 -window-width 400
```

### Multi-Frame DICOM

Extract a specific frame from multi-frame DICOM:

```bash
# Extract first frame (default)
go run main.go -input multiframe.dcm -frame 0

# Extract third frame
go run main.go -input multiframe.dcm -frame 2
```

### Export as JPEG

```bash
go run main.go -input sample.dcm -format jpeg -jpeg-quality 95
```

### Invert Grayscale

```bash
go run main.go -input sample.dcm -invert
```

## Command-Line Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-input` | string | (required) | Input DICOM file path |
| `-output` | string | auto | Output image file path (default: input filename with .png extension) |
| `-window-center` | float64 | 0 | Window center for display (0 = use DICOM value or auto) |
| `-window-width` | float64 | 0 | Window width for display (0 = use DICOM value or auto) |
| `-invert` | bool | false | Invert grayscale (swap black/white) |
| `-frame` | int | 0 | Frame number to extract (0-indexed, for multi-frame DICOM) |
| `-format` | string | png | Output format: png or jpeg |
| `-jpeg-quality` | int | 90 | JPEG quality (1-100, only for JPEG format) |

## Supported Image Types

### Grayscale Images
- MONOCHROME1 (inverted grayscale - automatically handled)
- MONOCHROME2 (normal grayscale)
- 8-bit and 16-bit pixel data
- Signed and unsigned pixel representation

### Color Images
- RGB
- YBR_FULL (automatically converted to RGB)
- YBR_FULL_422 (automatically converted to RGB)
- 8-bit and 16-bit pixel data per sample

### Multi-Frame Support
- Extracts individual frames from multi-frame DICOM files
- Useful for DICOM videos, time series, or multi-slice acquisitions

## Building

```bash
go build -o dicom_to_image.exe
```

## Examples

### Convert CT Image

```bash
dicom_to_image.exe -input ct_scan.dcm -window-center 40 -window-width 400
```

### Convert MRI Image

```bash
dicom_to_image.exe -input mri_brain.dcm
```

### Convert X-Ray Image

```bash
dicom_to_image.exe -input xray.dcm -window-center 2048 -window-width 4096
```

### Convert Multi-Frame DICOM to Series of Images

```bash
# Extract all frames from a multi-frame DICOM
for /L %i in (0,1,10) do dicom_to_image.exe -input multiframe.dcm -frame %i -output frame_%i.png
```

### Convert to JPEG with Custom Quality

```bash
dicom_to_image.exe -input sample.dcm -format jpeg -jpeg-quality 100 -output high_quality.jpg
```

## How It Works

1. **Parse DICOM File**: Uses the `parser` package to read the DICOM file
2. **Extract Image Information**: Retrieves image dimensions, bit depth, samples per pixel, etc.
3. **Create Pixel Data**: Uses the `imaging` package to create a structured representation
4. **Apply Rendering Pipeline**: Applies window/level adjustments and color space conversions
5. **Export Image**: Renders to PNG or JPEG using the `render` package

## Notes

- The program automatically detects image properties from DICOM tags
- Window center/width values are applied for optimal grayscale display
- If window parameters are not specified, the program will use values from DICOM tags or calculate them automatically
- RGB images ignore window parameters
- MONOCHROME1 images are automatically inverted (unless `-invert` is used to un-invert them)
- Multi-frame DICOM files are supported - specify the frame index with `-frame`
- Compressed DICOM files (JPEG, JPEG-LS, etc.) are not currently supported in this example
- Press Enter to exit after completion or error

## Troubleshooting

### "Missing Pixel Data" Error
- Ensure the DICOM file contains image data
- Some DICOM files (like SR, PR) don't contain pixel data

### "Frame index out of range" Error
- Check the number of frames in the DICOM file
- Use `-frame` with a valid frame index (0 to NumberOfFrames-1)

### Image Appears Too Dark or Too Bright
- Adjust window center and width using `-window-center` and `-window-width`
- Try different preset windows (lung, bone, soft tissue)

### Wrong Colors in RGB Images
- The program automatically handles YBR color space conversion
- If colors still look wrong, the DICOM file may have non-standard encoding
