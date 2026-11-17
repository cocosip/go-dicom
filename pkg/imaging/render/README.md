# DICOM Image Rendering Package

This package provides comprehensive DICOM image rendering capabilities, including LUT (Lookup Table) transformations, color space conversions, and image export functionality.

## Features

### LUT (Lookup Table) System

#### Base LUT Interface
All LUTs implement the `LUT` interface:
```go
type LUT interface {
    IsValid() bool
    MinimumOutputValue() float64
    MaximumOutputValue() float64
    Transform(input float64) float64
    Recalculate()
}
```

#### Modality LUT
Converts stored pixel values to modality-specific units (e.g., Hounsfield Units for CT).

**Modality Rescale LUT**:
```go
table := lut.NewModalityRescaleLUT(slope, intercept, minInput, maxInput)
output := table.Transform(pixelValue)  // output = pixelValue * slope + intercept
```

**Modality Sequence LUT**:
```go
lutData := []float64{10, 20, 30, 40, 50}
table := lut.NewModalitySequenceLUT(lutData, firstValueMapped, isSigned)
```

#### VOI LUT (Value of Interest)
Applies window/level transformations for visualization.

**Supported Functions**:
- `LINEAR`: Standard linear window/level
- `LINEAR_EXACT`: DICOM C.11.2.1.3.2 exact formula
- `SIGMOID`: Sigmoid curve transformation

```go
// Create LINEAR VOI LUT
table := lut.NewVOILinearLUT(windowCenter, windowWidth)

// Adjust window
table.SetWindow(newCenter, newWidth)

// Create using factory (auto-selects LINEAR_EXACT for width < 1.0)
voiLUT := lut.CreateVOILUT(lut.VOILUTFunctionLinear, center, width)
```

#### Composite LUT
Chains multiple LUTs together for complex transformations.

```go
modalityLUT := lut.NewModalityRescaleLUT(2.0, 0, 0, 4095)
voiLUT := lut.NewVOILinearLUT(2048, 4096)
invertLUT := lut.NewInvertLUT(0, 255)

compositeLUT := lut.NewCompositeLUT()
compositeLUT.Add(modalityLUT)
compositeLUT.Add(voiLUT)
compositeLUT.Add(invertLUT)
output := compositeLUT.Transform(input)
```

### Color Space Conversion

Supports conversion between multiple DICOM color spaces:

```go
// RGB <-> YBR_FULL
y, cb, cr := colorconv.RGBToYBRFull(r, g, b)
r, g, b = colorconv.YBRFullToRGB(y, cb, cr)

// YBR_FULL_422 (4:2:2 subsampling)
y1, y2, cb, cr := colorconv.RGBToYBRFull422(r1, g1, b1, r2, g2, b2)
r1, g1, b1, r2, g2, b2 = colorconv.YBRFull422ToRGB(y1, y2, cb, cr)

// YBR_PARTIAL_422 (BT.601 limited range)
y, cb, cr := colorconv.RGBToYBRPartial422(r, g, b)
r, g, b = colorconv.YBRPartial422ToRGB(y, cb, cr)

// YBR_ICT/RCT (JPEG 2000)
y, cb, cr := colorconv.RGBToYBRICT(r, g, b)
r, g, b = colorconv.YBRICTToRGB(y, cb, cr)

// Batch conversion
rgbData, err := colorconv.ConvertToRGB(ybrData, width, height, *pixel.YbrFull, planarConfig)
```

### Rendering Pipeline

The `GrayscalePipeline` provides a complete rendering pipeline:

```go
// Create pipeline
pipeline := render.NewGrayscalePipeline(
    rescaleSlope,      // Modality rescale slope
    rescaleIntercept,  // Modality rescale intercept
    windowCenter,      // Window center
    windowWidth,       // Window width
    minInput,          // Minimum input value
    maxInput,          // Maximum input value
    invert,            // Whether to invert output
)

// Adjust window
pipeline.SetWindow(center, width)
pipeline.SetWindowWidth(newWidth)
pipeline.SetWindowCenter(newCenter)

// Toggle inversion
pipeline.SetInvert(true)

// Get LUT for pixel transformation
lut := pipeline.LUT()
outputValue := lut.Transform(pixelValue)

// Clear cached LUT (forces rebuild)
pipeline.ClearCache()
```

**Pipeline Stages**:
1. **Modality LUT**: Converts stored values to modality units
2. **VOI LUT**: Applies window/level for visualization
3. **Invert LUT** (optional): Inverts output for MONOCHROME1

The pipeline is thread-safe and caches the composite LUT for performance.

### Image Export

Export DICOM images to standard formats (PNG, JPEG):

```go
// Create exporter with pipeline
pipeline := render.NewGrayscalePipeline(...)
exporter := render.NewImageExporter(pipeline)

// Configure export options
options := &render.ExportOptions{
    Format:      render.FormatPNG,  // or FormatJPEG
    JPEGQuality: 90,                 // Only for JPEG
}

// Export grayscale image
err := exporter.ExportGrayscale(
    writer,
    pixelData,
    width, height,
    bitsAllocated, bitsStored,
    isSigned,
    photometricInterpretation,
    options,
)

// Export RGB image
err = exporter.ExportRGB(
    writer,
    pixelData,
    width, height,
    photometricInterpretation,
    planarConfiguration,
    options,
)

// Convenience method
err = exporter.RenderFrame(
    writer,
    frameData,
    width, height,
    bitsAllocated, bitsStored,
    samplesPerPixel,
    isSigned,
    photometricInterpretation,
    planarConfiguration,
    options,
)
```

## Complete Example

```go
package main

import (
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/imaging/render"
)

func main() {
    // 1. Create rendering pipeline
    pipeline := render.NewGrayscalePipeline(
        1.0,    // slope
        -1024,  // intercept (for CT)
        0,      // window center (lung window)
        1500,   // window width
        -1024,  // min input
        3071,   // max input
        false,  // don't invert
    )

    // 2. Adjust window for soft tissue
    pipeline.SetWindow(40, 400)

    // 3. Create exporter
    exporter := render.NewImageExporter(pipeline)

    // 4. Export to PNG
    file, err := os.Create("output.png")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    options := &render.ExportOptions{
        Format: render.FormatPNG,
    }

    // 512x512 16-bit signed MONOCHROME2 frame data
    pixelData := make([]byte, 512*512*2)

    if err := exporter.ExportGrayscale(
        file,
        pixelData,
        512, 512,      // dimensions
        16, 12,        // bits allocated, bits stored
        true,          // signed
        "MONOCHROME2",
        options,
    ); err != nil {
        log.Fatal(err)
    }
}
```

## Architecture Notes

### LUT Design
- **Modality LUT**: Always returns `IsValid() = true`
- **VOI LUT**: `IsValid()` reports whether no recalculation is pending; a window change marks it stale again
- **Composite LUT**: Chains transformations efficiently
- **Precalculated table**: `GrayscalePipeline.PrecalculatedLUT` caches the stored-value table across frames and drops it whenever the pipeline LUT changes

### Color Space Formulas

**YBR_FULL (BT.601 Full Range)**:
- Y = 0.299*R + 0.587*G + 0.114*B
- Cb = -0.168736*R - 0.331264*G + 0.5*B + 128
- Cr = 0.5*R - 0.418688*G - 0.081312*B + 128

**YBR_PARTIAL_422 (BT.601 Limited Range)**:
- Y' = 0.2568*R + 0.5041*G + 0.0979*B + 16
- Cb = -0.1482*R - 0.2910*G + 0.4392*B + 128
- Cr = 0.4392*R - 0.3678*G - 0.0714*B + 128

**YBR_ICT (Irreversible Color Transform - JPEG 2000)**:
- Y = 0.299*R + 0.587*G + 0.114*B
- Cb = -0.1687*R - 0.3313*G + 0.5*B
- Cr = 0.5*R - 0.4187*G - 0.0813*B

**YBR_RCT (Reversible Color Transform - JPEG 2000 Lossless)**:
- Y = ⌊(R + 2*G + B) / 4⌋
- Cb = B - G
- Cr = R - G

### Performance

- **Pipeline caching**: Composite LUT is cached and only rebuilt when parameters change
- **Thread-safe**: All components use appropriate synchronization
- **Memory efficient**: Pixel data transformations are done in-place where possible

## Testing

The package includes comprehensive unit tests:

- **Pipeline Tests** (`pipeline_test.go`): grayscale pipeline, VOI/window updates, and concurrency
- **Export Tests** (`export_test.go`): grayscale export and bit-depth conversion
- **Affine Graphics Tests** (`affine_graphics_test.go`): overlay/supersampling geometry
- **LUT Tests**: `pkg/imaging/lut`
- **Color Space Tests**: `pkg/imaging/colorconv`

Run tests:
```bash
go test -v ./pkg/imaging/render/...
```

## References

- DICOM Standard Part 3: Information Object Definitions
- DICOM Standard Part 5: Data Structures and Encoding
- BT.601: Studio encoding parameters of digital television
- fo-dicom: Reference C# implementation
