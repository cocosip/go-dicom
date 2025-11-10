# go-dicom

A pure Go implementation of the DICOM (Digital Imaging and Communications in Medicine) standard, ported from the [fo-dicom](https://github.com/fo-dicom/fo-dicom) C# library.

[![Go Reference](https://pkg.go.dev/badge/github.com/cocosip/go-dicom.svg)](https://pkg.go.dev/github.com/cocosip/go-dicom)
[![Go Report Card](https://goreportcard.com/badge/github.com/cocosip/go-dicom)](https://goreportcard.com/report/github.com/cocosip/go-dicom)
[![License](https://img.shields.io/badge/license-MS--PL-blue.svg)](LICENSE)

## ⚠️ Work in Progress

This library is currently under active development. The API is not stable and may change significantly.

## Features

### Core Capabilities ✅

- ✅ **DICOM File I/O** - Read and write DICOM files with full standard compliance
- ✅ **Transfer Syntax Support** - Explicit/Implicit VR, Big/Little Endian, RLE, JPEG compression
- ✅ **Multi-Frame Images** - Full support for multi-frame and video DICOM files
- ✅ **Character Encoding** - 30+ character sets (UTF-8, Latin, Chinese, Japanese, Korean, Arabic, etc.)
- ✅ **Fragment Sequences** - Compressed pixel data handling (JPEG, RLE, etc.)
- ✅ **Structured Reports (SR)** - Parse and create SR documents with hierarchical content
- ✅ **Dataset Operations** - Rich API for accessing and manipulating DICOM elements
- ✅ **JSON/XML Serialization** - Export/Import DICOM data to JSON (Part 18) and XML
- ✅ **Anonymization** - Remove/Replace patient identifiable information
- ✅ **Pixel Data Processing** - Access raw pixel data, color space conversion, LUT operations

### Detailed Status (~75% Complete)

- [x] **Core DICOM data types**
  - [x] Tag (5338 standard tags + private tag support)
  - [x] VR (35 value representations with validation)
  - [x] VM (15 value multiplicities)
  - [x] Element (string, numeric, binary types)
  - [x] Dataset & Sequence (full support)
  - [x] Dictionary (tag/keyword lookup, Default global dictionary)
  - [x] UID (1965 standard UIDs)
  - [x] Transfer Syntax (15+ syntaxes including JPEG, RLE, MPEG)
  - [x] Character Set (30+ encodings)

- [x] **DICOM file reading**
  - [x] Explicit/Implicit VR parsing
  - [x] Sequence parsing (defined/undefined length)
  - [x] Fragment sequence support (compressed images)
  - [x] Multi-frame image support
  - [x] **ReadOption**: SkipLargeTags, ReadLargeOnDemand, ReadAll
  - [x] **FileFormat** detection: DICOM3, DICOM3NoPreamble, etc.
  - [x] Large object handling with configurable thresholds
  - [x] Tag.DictionaryEntry() for metadata lookup

- [x] **DICOM file writing**
  - [x] Explicit/Implicit VR writing
  - [x] Auto-generated File Meta Information
  - [x] Single and multi-frame image creation
  - [x] **DicomWriteOptions**:
    - ExplicitLengthSequences/Items
    - KeepGroupLengths
    - LargeObjectSize
    - Transfer syntax selection
  - [x] Group length auto-filtering
  - [x] Explicit/undefined length sequences

- [x] **Special Format Support**
  - [x] Multi-frame images (verified up to 100 frames)
  - [x] Fragment sequences (JPEG, RLE compressed data)
  - [x] Video DICOM (MPEG2)
  - [x] RGB color images (planar and interleaved)
  - [x] Structured Reports (SR) with hierarchical content
  - [x] Modality LUT Sequences
  - [x] Character set variants (17+ tested encodings)

- [x] **JSON/XML Serialization**
  - [x] DICOM JSON Model (Part 18)
  - [x] Native XML format
  - [x] Bulkdata handling
  - [x] Pretty-print options

- [x] **Anonymization**
  - [x] Basic anonymization profile
  - [x] Custom anonymization rules
  - [x] Patient/Study/Series level anonymization

- [x] **Imaging Support**
  - [x] Pixel data handling
  - [x] Color space conversion (YBR↔RGB)
  - [x] Planar/Interleaved conversion
  - [x] LUT (Lookup Table) operations
  - [x] VOI windowing
  - [x] Overlay data support
  - [x] Palette color support

- [x] **Structured Reports**
  - [x] SR content items (TEXT, NUM, CODE, CONTAINER, IMAGE, COMPOSITE)
  - [x] Hierarchical structure
  - [x] Code items and measured values
  - [x] Relationship types (CONTAINS, HAS OBS CONTEXT, etc.)

- [ ] **DICOM Networking** (Future)
  - [ ] DIMSE services (C-ECHO, C-FIND, C-STORE, etc.)
  - [ ] Association management
  - [ ] SCP/SCU implementation

- [ ] **Image Codecs** (Future - Decompression)
  - [ ] JPEG codec
  - [ ] JPEG-LS codec
  - [ ] JPEG 2000 codec
  - [ ] RLE codec

See [TODO.md](TODO.md) for detailed development roadmap.

## Performance

go-dicom is designed for high performance with minimal memory allocations. See [BENCHMARKS.md](BENCHMARKS.md) for detailed performance metrics.

**Key Performance Highlights:**
- Dataset Get/Add operations: ~2-13 ns/op with zero allocations
- Parser: ~1.4 μs/op for typical datasets
- Writer: Scales linearly with dataset size (~1.2 μs for small datasets)
- Zero-allocation endian swapping operations

Run benchmarks yourself:
```bash
go test -bench=. -benchmem ./pkg/dicom/...
```

## Installation

```bash
go get github.com/cocosip/go-dicom
```

## Quick Start

### Reading a DICOM File

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
    // Open DICOM file
    file, err := os.Open("example.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Parse the file
    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Get patient information
    patientName, exists := result.Dataset.GetString(tag.PatientName)
    if exists {
        fmt.Printf("Patient Name: %s\n", patientName)
    }

    patientID, _ := result.Dataset.GetString(tag.PatientID)
    fmt.Printf("Patient ID: %s\n", patientID)

    // Get study information
    studyDate, _ := result.Dataset.GetString(tag.StudyDate)
    modality, _ := result.Dataset.GetString(tag.Modality)
    fmt.Printf("Study Date: %s\n", studyDate)
    fmt.Printf("Modality: %s\n", modality)

    // Get file format
    fmt.Printf("Format: %s\n", result.Format)
    fmt.Printf("Total elements: %d\n", result.Dataset.Count())
}
```

### Reading Image Properties

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
    file, err := os.Open("image.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Get image dimensions
    rows, err := result.Dataset.GetUInt16(tag.Rows, 0)
    if err == nil {
        fmt.Printf("Rows: %d\n", rows)
    }

    cols, err := result.Dataset.GetUInt16(tag.Columns, 0)
    if err == nil {
        fmt.Printf("Columns: %d\n", cols)
    }

    // Get bit depth
    bitsAllocated, _ := result.Dataset.GetUInt16(tag.BitsAllocated, 0)
    bitsStored, _ := result.Dataset.GetUInt16(tag.BitsStored, 0)
    fmt.Printf("Bits Allocated: %d\n", bitsAllocated)
    fmt.Printf("Bits Stored: %d\n", bitsStored)

    // Get photometric interpretation
    photoInterp, _ := result.Dataset.GetString(tag.PhotometricInterpretation)
    fmt.Printf("Photometric Interpretation: %s\n", photoInterp)

    // Check if multi-frame
    numFrames, exists := result.Dataset.GetString(tag.NumberOfFrames)
    if exists {
        fmt.Printf("Number of Frames: %s\n", numFrames)
    } else {
        fmt.Println("Single frame image")
    }
}
```

### Accessing Pixel Data

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
    file, err := os.Open("image.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Get pixel data element
    pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
    if !exists {
        log.Fatal("No pixel data found")
    }

    fmt.Printf("Pixel data type: %T\n", pixelDataElem)

    // For uncompressed data (OtherWord/OtherByte)
    if pd, ok := pixelDataElem.(interface{ GetData() []byte }); ok {
        pixelData := pd.GetData()
        fmt.Printf("Pixel data size: %d bytes\n", len(pixelData))
        // Process raw pixel data...
    }

    // For compressed data (fragment sequences)
    if pd, ok := pixelDataElem.(interface {
        FragmentCount() int
        Fragments() interface{}
    }); ok {
        fragmentCount := pd.FragmentCount()
        fmt.Printf("Compressed data with %d fragments\n", fragmentCount)
        // Access individual fragments...
    }
}
```

### Reading Multi-Frame Images

```go
package main

import (
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
    file, err := os.Open("multiframe.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Get number of frames
    numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames)
    if !exists {
        fmt.Println("Single frame image")
        return
    }

    numFrames := strings.TrimSpace(numFramesStr)
    fmt.Printf("Multi-frame image with %s frames\n", numFrames)

    // Get frame dimensions
    rows, _ := result.Dataset.GetUInt16(tag.Rows, 0)
    cols, _ := result.Dataset.GetUInt16(tag.Columns, 0)
    bitsAllocated, _ := result.Dataset.GetUInt16(tag.BitsAllocated, 0)

    // Calculate frame size
    bytesPerPixel := int(bitsAllocated) / 8
    frameSize := int(rows) * int(cols) * bytesPerPixel
    fmt.Printf("Each frame: %dx%d, %d bytes\n", cols, rows, frameSize)

    // Get pixel data
    pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
    if !exists {
        log.Fatal("No pixel data found")
    }

    if pd, ok := pixelDataElem.(interface{ GetData() []byte }); ok {
        allFramesData := pd.GetData()
        totalFrames := len(allFramesData) / frameSize
        fmt.Printf("Pixel data contains %d frames\n", totalFrames)

        // Extract individual frames
        for frame := 0; frame < totalFrames; frame++ {
            frameOffset := frame * frameSize
            frameData := allFramesData[frameOffset : frameOffset+frameSize]
            fmt.Printf("Frame %d: %d bytes\n", frame, len(frameData))
            // Process individual frame...
        }
    }
}
```

### Reading with Options (Skip Large Tags)

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
    file, err := os.Open("large_image.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Parse with options: skip large tags like pixel data
    result, err := parser.Parse(file,
        parser.WithReadOption(parser.SkipLargeTags),
        parser.WithLargeObjectSize(128*1024), // 128KB threshold
    )
    if err != nil {
        log.Fatal(err)
    }

    // Access metadata (small tags still available)
    patientName, _ := result.Dataset.GetString(tag.PatientName)
    modality, _ := result.Dataset.GetString(tag.Modality)

    fmt.Printf("Patient: %s\n", patientName)
    fmt.Printf("Modality: %s\n", modality)
    fmt.Printf("Loaded elements: %d\n", result.Dataset.Count())

    // Pixel data was skipped, so it won't be in the dataset
    _, exists := result.Dataset.Get(tag.PixelData)
    fmt.Printf("Pixel data loaded: %v\n", exists)
}
```

## Writing DICOM Files

### Creating a Simple DICOM File

```go
package main

import (
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
    // Create new dataset
    ds := dataset.New()

    // Add required elements
    ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
    ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))

    // Add patient information
    ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
    ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
    ds.Add(element.NewString(tag.PatientBirthDate, vr.DA, []string{"19800101"}))
    ds.Add(element.NewString(tag.PatientSex, vr.CS, []string{"M"}))

    // Add study information
    ds.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8"}))
    ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250110"}))
    ds.Add(element.NewString(tag.StudyTime, vr.TM, []string{"120000"}))
    ds.Add(element.NewString(tag.StudyDescription, vr.LO, []string{"Test Study"}))

    // Add series information
    ds.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7"}))
    ds.Add(element.NewString(tag.Modality, vr.CS, []string{"CT"}))
    ds.Add(element.NewString(tag.SeriesNumber, vr.IS, []string{"1"}))

    // Write to file
    file, err := os.Create("output.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write with default options (Explicit VR Little Endian)
    if err := writer.Write(file, ds); err != nil {
        log.Fatal(err)
    }

    log.Println("DICOM file created successfully")
}
```

### Creating a Single Frame Image

```go
package main

import (
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
    ds := dataset.New()

    // Add required metadata
    ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
    ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
    ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))
    ds.Add(element.NewString(tag.Modality, vr.CS, []string{"CT"}))

    // Add image properties
    rows := uint16(512)
    cols := uint16(512)
    ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{rows}))
    ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{cols}))
    ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}))
    ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{16}))
    ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{15}))
    ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
    ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
    ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

    // Create pixel data (512x512x2 bytes = 524,288 bytes)
    pixelDataSize := int(rows) * int(cols) * 2
    pixelData := make([]byte, pixelDataSize)

    // Fill with test pattern (gradient)
    for i := 0; i < pixelDataSize/2; i++ {
        value := uint16(i % 65536)
        pixelData[i*2] = byte(value & 0xFF)
        pixelData[i*2+1] = byte((value >> 8) & 0xFF)
    }

    // Add pixel data
    ds.Add(element.NewOtherWord(tag.PixelData, pixelData))

    // Write to file
    file, err := os.Create("single_frame.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    if err := writer.Write(file, ds); err != nil {
        log.Fatal(err)
    }

    log.Printf("Created single frame image: %dx%d\n", cols, rows)
}
```

### Creating a Multi-Frame Image

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
    ds := dataset.New()

    // Add required metadata
    ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
    ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6"}))
    ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"MultiFrame^Test"}))
    ds.Add(element.NewString(tag.Modality, vr.CS, []string{"MR"}))

    // Add image properties
    rows := uint16(256)
    cols := uint16(256)
    numFrames := 10

    ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{rows}))
    ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{cols}))
    ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}))
    ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{16}))
    ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{15}))
    ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
    ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
    ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

    // IMPORTANT: Add NumberOfFrames for multi-frame images
    ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{fmt.Sprintf("%d", numFrames)}))

    // Create multi-frame pixel data
    frameSize := int(rows) * int(cols) * 2 // 2 bytes per pixel
    totalPixelDataSize := frameSize * numFrames
    pixelData := make([]byte, totalPixelDataSize)

    // Fill each frame with different pattern
    for frame := 0; frame < numFrames; frame++ {
        frameOffset := frame * frameSize
        baseValue := uint16(frame * 6000) // Different intensity per frame

        for i := 0; i < frameSize/2; i++ {
            value := baseValue + uint16(i%1000)
            pixelData[frameOffset+i*2] = byte(value & 0xFF)
            pixelData[frameOffset+i*2+1] = byte((value >> 8) & 0xFF)
        }
    }

    // Add pixel data
    ds.Add(element.NewOtherWord(tag.PixelData, pixelData))

    // Write to file
    file, err := os.Create("multi_frame.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    if err := writer.Write(file, ds); err != nil {
        log.Fatal(err)
    }

    log.Printf("Created multi-frame image: %d frames of %dx%d\n", numFrames, cols, rows)
}
```

### Writing with Custom Transfer Syntax

```go
package main

import (
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/transfer"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
    ds := dataset.New()

    // Add elements...
    ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))
    // ... more elements ...

    file, err := os.Create("output_explicit.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write with explicit VR Little Endian
    err = writer.Write(file, ds,
        writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian),
        writer.WithExplicitLengthSequences(true),  // Use explicit sequence lengths
        writer.WithLargeObjectSize(1024*1024),     // 1MB threshold for large objects
    )
    if err != nil {
        log.Fatal(err)
    }

    log.Println("DICOM file created with explicit VR Little Endian")
}
```

## Dataset Operations

### Iterating Through Elements

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
)

func main() {
    file, err := os.Open("example.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Get all elements
    elements := result.Dataset.Elements()

    fmt.Printf("Total elements: %d\n", len(elements))
    fmt.Println("\nAll DICOM tags:")

    for _, elem := range elements {
        tag := elem.Tag()
        vr := elem.VR()

        // Get tag dictionary entry for tag name
        entry := tag.DictionaryEntry()
        tagName := "Unknown"
        if entry != nil {
            tagName = entry.Name
        }

        fmt.Printf("(%04X,%04X) %s [%s]: ",
            tag.Group(), tag.Element(), tagName, vr.String())

        // Try to get string value
        if strElem, ok := elem.(interface{ GetString() string }); ok {
            value := strElem.GetString()
            if len(value) > 50 {
                value = value[:50] + "..."
            }
            fmt.Printf("%s\n", value)
        } else {
            fmt.Printf("[%T]\n", elem)
        }
    }
}
```

### Working with Sequences

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
    file, err := os.Open("sr_document.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Access a sequence (e.g., Content Sequence in SR)
    contentSeqElem, exists := result.Dataset.Get(tag.ContentSequence)
    if !exists {
        log.Fatal("Content Sequence not found")
    }

    // Cast to Sequence type
    contentSeq, ok := contentSeqElem.(*dataset.Sequence)
    if !ok {
        log.Fatal("Element is not a sequence")
    }

    fmt.Printf("Content Sequence has %d items\n", contentSeq.Count())

    // Iterate through sequence items
    for i := 0; i < contentSeq.Count(); i++ {
        item := contentSeq.GetItem(i)

        fmt.Printf("\n--- Item %d ---\n", i)

        // Get specific tags from item
        if relType, exists := item.GetString(tag.RelationshipType); exists {
            fmt.Printf("Relationship Type: %s\n", relType)
        }

        if valueType, exists := item.GetString(tag.ValueType); exists {
            fmt.Printf("Value Type: %s\n", valueType)

            // Get value based on type
            if valueType == "TEXT" {
                if textValue, exists := item.GetString(tag.TextValue); exists {
                    fmt.Printf("Text Value: %s\n", textValue)
                }
            }
        }

        // Check for nested sequences
        if nestedSeq, exists := item.Get(tag.ContentSequence); exists {
            fmt.Println("  [Contains nested Content Sequence]")
        }
    }
}
```

### Modifying and Anonymizing Data

```go
package main

import (
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
    // Read existing DICOM file
    file, err := os.Open("original.dcm")
    if err != nil {
        log.Fatal(err)
    }

    result, err := parser.Parse(file)
    file.Close()
    if err != nil {
        log.Fatal(err)
    }

    ds := result.Dataset

    // Remove patient identifiable information
    ds.Remove(tag.PatientName)
    ds.Remove(tag.PatientID)
    ds.Remove(tag.PatientBirthDate)
    ds.Remove(tag.PatientAddress)
    ds.Remove(tag.PatientTelephoneNumbers)

    // Add anonymized values
    ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"ANONYMOUS"}))
    ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"ANON0001"}))

    // Modify study description
    ds.Remove(tag.StudyDescription)
    ds.Add(element.NewString(tag.StudyDescription, vr.LO, []string{"Anonymized Study"}))

    // Write anonymized file
    outFile, err := os.Create("anonymized.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer outFile.Close()

    if err := writer.Write(outFile, ds); err != nil {
        log.Fatal(err)
    }

    log.Println("Anonymized DICOM file created")
}
```

## Special Format Support

### Reading Structured Reports (SR)

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
    file, err := os.Open("sr_report.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Verify it's a Structured Report
    sopClassUID, _ := result.Dataset.GetString(tag.SOPClassUID)
    modality, _ := result.Dataset.GetString(tag.Modality)

    fmt.Printf("SOP Class: %s\n", sopClassUID)
    fmt.Printf("Modality: %s\n", modality)

    // Get SR document info
    verificationFlag, _ := result.Dataset.GetString(tag.VerificationFlag)
    completionFlag, _ := result.Dataset.GetString(tag.CompletionFlag)

    fmt.Printf("Verification: %s\n", verificationFlag)
    fmt.Printf("Completion: %s\n", completionFlag)

    // Access Content Sequence (main SR data structure)
    contentSeqElem, exists := result.Dataset.Get(tag.ContentSequence)
    if !exists {
        log.Fatal("Content Sequence not found")
    }

    contentSeq := contentSeqElem.(*dataset.Sequence)
    fmt.Printf("\nContent Sequence has %d items\n", contentSeq.Count())

    // Process each content item
    for i := 0; i < contentSeq.Count(); i++ {
        item := contentSeq.GetItem(i)

        relType, _ := item.GetString(tag.RelationshipType)
        valueType, _ := item.GetString(tag.ValueType)

        fmt.Printf("\nItem %d:\n", i)
        fmt.Printf("  Relationship: %s\n", relType)
        fmt.Printf("  Value Type: %s\n", valueType)

        switch valueType {
        case "TEXT":
            if text, exists := item.GetString(tag.TextValue); exists {
                fmt.Printf("  Text: %s\n", text)
            }
        case "CODE":
            if codeSeq, exists := item.Get(tag.ConceptCodeSequence); exists {
                fmt.Printf("  [Code Sequence]\n")
            }
        case "NUM":
            fmt.Printf("  [Numeric Measurement]\n")
        case "CONTAINER":
            fmt.Printf("  [Container]\n")
        }
    }
}
```

### Handling Compressed Images (RLE, JPEG)

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/io/buffer"
)

func main() {
    file, err := os.Open("compressed_rle.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Check transfer syntax
    if result.FileMetaInformation != nil {
        tsUID, _ := result.FileMetaInformation.TransferSyntaxUID()
        fmt.Printf("Transfer Syntax: %s\n", tsUID)

        // 1.2.840.10008.1.2.5 = RLE Lossless
        if tsUID == "1.2.840.10008.1.2.5" {
            fmt.Println("This is RLE compressed")
        }
    }

    // Get pixel data
    pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
    if !exists {
        log.Fatal("No pixel data")
    }

    // For compressed data, pixel data is stored as fragment sequence
    if fragSeq, ok := pixelDataElem.(interface {
        FragmentCount() int
        Fragments() []buffer.ByteBuffer
    }); ok {
        fragmentCount := fragSeq.FragmentCount()
        fmt.Printf("Compressed data with %d fragments\n", fragmentCount)

        fragments := fragSeq.Fragments()
        totalSize := 0

        for i, buf := range fragments {
            data := buf.Data()
            fragmentSize := len(data)
            totalSize += fragmentSize

            fmt.Printf("Fragment %d: %d bytes\n", i, fragmentSize)
        }

        fmt.Printf("Total compressed size: %d bytes\n", totalSize)

        // Note: To decompress, you would need a codec
        // The raw compressed data is accessible via fragments
    }
}
```

## Advanced Features

### Character Encoding Support

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
    // The parser automatically handles character encoding
    file, err := os.Open("chinese_patient.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Get specific character set
    charset, exists := result.Dataset.GetString(tag.SpecificCharacterSet)
    if exists {
        fmt.Printf("Character Set: %s\n", charset)
    }

    // Patient name is automatically decoded using the correct character set
    patientName, _ := result.Dataset.GetString(tag.PatientName)
    fmt.Printf("Patient Name: %s\n", patientName)

    // Supported character sets include:
    // - ISO_IR 100 (Latin-1)
    // - ISO_IR 192 (UTF-8)
    // - ISO_IR 126 (Greek)
    // - ISO_IR 127 (Arabic)
    // - ISO_IR 138 (Hebrew)
    // - ISO_IR 144 (Cyrillic)
    // - GB18030 (Chinese Simplified)
    // - GBK (Chinese Simplified)
    // - ISO 2022 IR 87 (Japanese)
    // - ISO 2022 IR 149 (Korean)
    // And many more...
}
```

### Large File Handling

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
    file, err := os.Open("very_large_image.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Option 1: Skip large tags completely
    result, err := parser.Parse(file,
        parser.WithReadOption(parser.SkipLargeTags),
        parser.WithLargeObjectSize(10*1024*1024), // Skip tags > 10MB
    )
    if err != nil {
        log.Fatal(err)
    }

    // Metadata is available, but pixel data was skipped
    patientName, _ := result.Dataset.GetString(tag.PatientName)
    fmt.Printf("Patient: %s\n", patientName)

    _, hasPixelData := result.Dataset.Get(tag.PixelData)
    fmt.Printf("Pixel data loaded: %v\n", hasPixelData) // false

    // Option 2: Read all data (default)
    file2, _ := os.Open("large_image.dcm")
    defer file2.Close()

    resultFull, err := parser.Parse(file2,
        parser.WithReadOption(parser.ReadAll),
    )
    if err != nil {
        log.Fatal(err)
    }

    _, hasPixelData2 := resultFull.Dataset.Get(tag.PixelData)
    fmt.Printf("Full load - Pixel data loaded: %v\n", hasPixelData2) // true
}
```

## Testing

The library includes comprehensive test coverage for various DICOM formats:

```bash
# Run all tests
go test ./...

# Run parser tests
go test -v ./pkg/dicom/parser

# Run specific test
go test -v ./pkg/dicom/parser -run TestMultiFrame

# Run with coverage
go test -cover ./...
```

### Tested DICOM Formats

✅ **Character Sets** (17 encodings tested)
- Latin-1, UTF-8, Greek, Arabic, Hebrew, Cyrillic
- Chinese (GB2312, GBK, GB18030)
- Japanese (Shift-JIS, ISO-2022-JP)
- Korean (EUC-KR)

✅ **Image Formats**
- Single frame images
- Multi-frame images (tested up to 100 frames)
- RGB color images (planar and interleaved)
- Video DICOM (MPEG2)

✅ **Compression**
- RLE Lossless (fragment sequences)
- JPEG compressed (fragment sequences)
- Uncompressed

✅ **Special Formats**
- Structured Reports (SR)
- Modality LUT Sequences
- Fragment sequences with offset tables

✅ **Write/Read Cycle**
- Single and multi-frame image creation
- Parser can read writer-generated files
- Frame count and pixel data integrity verified

## Project Structure

```
go-dicom/
├── pkg/
│   ├── dicom/              # Core DICOM types
│   │   ├── tag/            # DICOM tags
│   │   ├── vr/             # Value representations
│   │   ├── element/        # DICOM elements
│   │   ├── dataset/        # Dataset and sequences
│   │   ├── parser/         # File parsing
│   │   ├── writer/         # File writing
│   │   ├── dictionary/     # Tag dictionary
│   │   ├── transfer/       # Transfer syntaxes
│   │   └── uid/            # Standard UIDs
│   ├── io/                 # I/O operations
│   │   ├── buffer/         # Byte buffer abstractions
│   │   ├── reader/         # Low-level DICOM reading
│   │   └── writer/         # Low-level DICOM writing
│   └── charset/            # Character encoding support
├── test-data/              # Test DICOM files
├── examples/               # Usage examples
├── BENCHMARKS.md           # Performance benchmarks
├── TODO.md                 # Development roadmap
└── CLAUDE.md               # Architecture documentation
```

## Development

### Prerequisites

- Go 1.21 or later
- golangci-lint (for linting)

### Building

```bash
# Build all packages
go build ./...

# Run tests
go test ./...

# Run benchmarks
go test -bench=. -benchmem ./pkg/dicom/...
```

### Code Quality

```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Run go vet
go vet ./...
```

## Migration from fo-dicom (C#)

This project is a port of the fo-dicom library from C# to Go. The migration follows an incremental approach:

1. ✅ Core data types (Tag, VR, Dictionary)
2. ✅ Data structures (Element, Dataset)
3. ✅ I/O capabilities (Reader, Writer)
4. ⏳ Advanced features (Networking, Codecs)

See [CLAUDE.md](CLAUDE.md) for detailed architecture documentation.

## DICOM Resources

- [DICOM Standard](https://www.dicomstandard.org/)
- [fo-dicom GitHub Repository](https://github.com/fo-dicom/fo-dicom)
- [DICOM Library](https://dicomlibrary.com/) - Free DICOM images for testing

## License

This project is licensed under the Microsoft Public License (MS-PL), the same license as fo-dicom.

See [LICENSE](LICENSE) for details.

## Acknowledgments

This project is heavily inspired by and based on [fo-dicom](https://github.com/fo-dicom/fo-dicom), an excellent DICOM library for .NET.

Special thanks to the fo-dicom contributors for their comprehensive implementation of the DICOM standard.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

### Development Guidelines

1. Follow Go best practices and idioms
2. Write tests for all new functionality
3. Update documentation as needed
4. Ensure all tests pass before submitting PRs

## Status

**Current Phase**: Core Implementation Complete (~75%)

**Working Features**:
- ✅ Complete DICOM file I/O (read/write)
- ✅ Multi-frame image support
- ✅ Character encoding (30+ encodings)
- ✅ Fragment sequences (compressed data)
- ✅ Structured Reports
- ✅ Dataset manipulation
- ✅ JSON/XML serialization

**Next Steps**:
- DICOM networking (C-ECHO, C-STORE, C-FIND)
- Image codecs (JPEG, RLE decompression)

See [TODO.md](TODO.md) for complete task list and progress tracking.
