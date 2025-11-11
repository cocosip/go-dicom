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
- ✅ **Transfer Syntax Support** - Explicit/Implicit VR, Big/Little Endian (RLE codec implemented, JPEG/JPEG2K planned)
- ✅ **Multi-Frame Images** - Full support for multi-frame and video DICOM files
- ✅ **Character Encoding** - 30+ character sets with auto-detection (UTF-8, Latin, Chinese, Japanese, Korean, Arabic, etc.)
- ✅ **Fragment Sequences** - Compressed pixel data handling (read/write, codec framework ready)
- ✅ **Structured Reports (SR)** - Parse and create SR documents with hierarchical content
- ✅ **Dataset Operations** - Rich API for accessing and manipulating DICOM elements
- ✅ **JSON/XML Serialization** - Export/Import DICOM data to JSON (Part 18) and XML formats
- ✅ **Anonymization** - Remove/Replace patient identifiable information with configurable profiles
- ✅ **Pixel Data Processing** - Access raw pixel data, color space conversion, LUT operations, image rendering
- ✅ **DICOM Networking** - Complete DIMSE services (C-ECHO/STORE/FIND/MOVE/GET + N-CREATE/GET/SET/DELETE/ACTION/EVENT-REPORT), TLS support
- ✅ **Image Processing** - Rendering pipeline, windowing, LUT operations, color space conversion
- ✅ **DICOM Printing** - Film Session, Film Box, Image Box support

### Detailed Status

- [x] **Core DICOM data types** (~100% Complete)
  - [x] Tag (5338 standard tags + private tag support)
  - [x] VR (35 value representations with validation)
  - [x] VM (15 value multiplicities)
  - [x] Element (string, numeric, binary, date, person name types)
  - [x] Dataset & Sequence (full support with lazy loading)
  - [x] Dictionary (tag/keyword lookup with global default instance)
  - [x] UID (1965 standard UIDs)
  - [x] Transfer Syntax (15+ syntaxes including JPEG, RLE, MPEG)
  - [x] Character Set (30+ encodings with auto-detection)

- [x] **DICOM file reading** (~100% Complete)
  - [x] Explicit/Implicit VR parsing
  - [x] Sequence parsing (defined/undefined length)
  - [x] Fragment sequence support (compressed images)
  - [x] Multi-frame image support
  - [x] ReadOptions: SkipLargeTags, ReadLargeOnDemand, ReadAll
  - [x] FileFormat detection: DICOM3, DICOM3NoPreamble, ACR-NEMA
  - [x] Large object handling with configurable thresholds
  - [x] Automatic character set detection and conversion
  - [x] Byte order handling (Little/Big Endian)

- [x] **DICOM file writing** (~100% Complete)
  - [x] Explicit/Implicit VR writing
  - [x] Auto-generated File Meta Information (FMI)
  - [x] Single and multi-frame image creation
  - [x] WriteOptions:
    - ExplicitLengthSequences/Items (vs undefined length)
    - KeepGroupLengths (default: auto-remove deprecated group lengths)
    - LargeObjectSize threshold configuration
    - Transfer syntax selection (Explicit/Implicit VR, Big/Little Endian)
  - [x] Group length auto-filtering (removes deprecated (GGGG,0000) tags)
  - [x] Proper byte order handling

- [x] **Special Format Support**
  - [x] Multi-frame images (verified up to 100 frames)
  - [x] Fragment sequences (JPEG, RLE compressed data)
  - [x] Video DICOM (MPEG2)
  - [x] RGB color images (planar and interleaved)
  - [x] Structured Reports (SR) with hierarchical content
  - [x] Modality LUT Sequences
  - [x] Character set variants (17+ tested encodings)

- [x] **JSON/XML Serialization** (~100% Complete)
  - [x] DICOM JSON Model (Part 18 compliant)
  - [x] Native XML format
  - [x] Bulkdata handling with base64 encoding
  - [x] Pretty-print options
  - [x] PersonName component groups support
  - [x] Sequence nesting support

- [x] **Anonymization** (~95% Complete)
  - [x] Basic anonymization profile (patient identifiable information)
  - [x] Custom anonymization rules (Remove, Replace, Keep)
  - [x] Patient/Study/Series level anonymization
  - [x] Date shifting and UID remapping
  - [x] Recursive sequence anonymization

- [x] **Imaging Support** (~90% Complete)
  - [x] Pixel data extraction and handling
  - [x] Color space conversion (YBR↔RGB)
  - [x] Planar/Interleaved conversion
  - [x] LUT (Lookup Table) operations (Modality LUT, VOI LUT)
  - [x] VOI windowing (window center/width)
  - [x] Overlay data extraction
  - [x] Palette color LUT support
  - [x] Bit depth conversion and scaling
  - [x] Image reconstruction from pixel data
  - [x] Rendering pipeline with configurable options

- [x] **Structured Reports** (~95% Complete)
  - [x] SR content items (TEXT, NUM, CODE, CONTAINER, IMAGE, COMPOSITE, etc.)
  - [x] Hierarchical structure with parent-child relationships
  - [x] Code items with coding scheme designators
  - [x] Measured values with units
  - [x] Referenced SOP instances
  - [x] Relationship types (CONTAINS, HAS OBS CONTEXT, INFERRED FROM, etc.)
  - [x] Content sequence parsing and creation

- [x] **DICOM Networking** (~95% Complete)
  - [x] PDU (Protocol Data Unit) encoding/decoding (7 PDU types)
  - [x] Association negotiation (A-ASSOCIATE-RQ/AC/RJ, A-RELEASE, A-ABORT)
  - [x] Presentation context negotiation with transfer syntax support
  - [x] DIMSE C-services:
    - [x] C-ECHO (verification)
    - [x] C-STORE (image storage)
    - [x] C-FIND (query/retrieve)
    - [x] C-MOVE (retrieval - pull mode)
    - [x] C-GET (retrieval - push mode)
  - [x] DIMSE N-services:
    - [x] N-EVENT-REPORT (event notification)
    - [x] N-GET (retrieve attributes)
    - [x] N-SET (modify attributes)
    - [x] N-ACTION (perform action)
    - [x] N-CREATE (create instance)
    - [x] N-DELETE (delete instance)
  - [x] DICOM Client (SCU) implementation with sync/async API
  - [x] DICOM Server (SCP) framework with handler pattern
  - [x] Network transport abstraction (TCP/TLS)
  - [x] TLS support (secure DICOM connections)
  - [x] Concurrent-safe operations
  - [x] 401+ unit tests with >85% code coverage
  - [ ] Advanced role negotiation - Planned
  - [ ] Extended negotiation items - Planned

- [x] **Image Codecs** (~60% Complete)
  - [x] Native codecs (uncompressed data - Explicit/Implicit VR, Little/Big Endian)
  - [x] RLE codec (RLE Lossless encode/decode - fully functional)
  - [x] Transcoder framework for format conversion between transfer syntaxes
  - [x] Codec registry and plugin architecture
  - [ ] JPEG Baseline codec (1.2.840.10008.1.2.4.50) - Planned
  - [ ] JPEG Lossless codec (1.2.840.10008.1.2.4.57/70) - Planned
  - [ ] JPEG-LS Lossless/Near-Lossless (1.2.840.10008.1.2.4.80/81) - Planned
  - [ ] JPEG 2000 Lossless/Lossy (1.2.840.10008.1.2.4.90/91) - Planned
  - [ ] MPEG-2/MPEG-4 Video codecs - Planned

  **Note**: The codec framework is complete and extensible. Fragment sequence reading/writing is fully implemented.
  Images compressed with JPEG/JPEG2K can be read (fragment sequences extracted), but decompression requires
  codec implementation or external libraries.

- [x] **DICOM Printing** (~90% Complete)
  - [x] Film Session management
  - [x] Film Box configuration
  - [x] Image Box handling
  - [x] Presentation LUT
  - [x] Print job creation
  - [ ] Printer status management - Planned

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

### DICOM Networking - C-ECHO

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    // Create DICOM client
    c := client.New(
        client.WithCallingAE("GO-SCU"),
        client.WithCalledAE("ANY-SCP"),
    )

    // Add Verification SOP Class
    c.AddPresentationContext(
        "1.2.840.10008.1.1",   // Verification SOP Class
        "1.2.840.10008.1.2.1", // Explicit VR Little Endian
    )

    // Connect to DICOM server
    ctx := context.Background()
    if err := c.Connect(ctx, "localhost", 11112); err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer c.Close()

    // Send C-ECHO request
    if err := c.CEcho(ctx); err != nil {
        log.Fatalf("C-ECHO failed: %v", err)
    }

    fmt.Println("C-ECHO successful - DICOM server is alive")
}
```

### DICOM Networking - TLS Secure Connection

```go
package main

import (
    "context"
    "crypto/tls"
    "log"

    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    // Create TLS configuration
    tlsConfig := &tls.Config{
        ServerName:         "pacs.example.com",
        InsecureSkipVerify: false, // Set to true only for testing
    }

    // Create DICOM client with TLS
    c := client.New(
        client.WithCallingAE("SECURE-SCU"),
        client.WithCalledAE("SECURE-SCP"),
        client.WithTLS(tlsConfig),
    )

    // Add presentation context
    c.AddPresentationContext(
        "1.2.840.10008.1.1",   // Verification SOP Class
        "1.2.840.10008.1.2.1", // Explicit VR Little Endian
    )

    // Connect with TLS
    ctx := context.Background()
    if err := c.Connect(ctx, "pacs.example.com", 11112); err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer c.Close()

    // Send C-ECHO over secure connection
    if err := c.CEcho(ctx); err != nil {
        log.Fatalf("C-ECHO failed: %v", err)
    }

    log.Println("Secure DICOM connection established successfully!")
}
```

### Image Processing and Rendering

```go
package main

import (
    "image/png"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/imaging"
    "github.com/cocosip/go-dicom/pkg/imaging/render"
)

func main() {
    // Read DICOM file
    file, err := os.Open("ct_image.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    result, err := parser.Parse(file)
    if err != nil {
        log.Fatal(err)
    }

    // Create DICOM image
    dicomImage, err := imaging.NewDicomImage(result.Dataset)
    if err != nil {
        log.Fatal(err)
    }

    // Render with default options
    img, err := render.Render(dicomImage, nil)
    if err != nil {
        log.Fatal(err)
    }

    // Save as PNG
    outFile, err := os.Create("output.png")
    if err != nil {
        log.Fatal(err)
    }
    defer outFile.Close()

    if err := png.Encode(outFile, img); err != nil {
        log.Fatal(err)
    }

    log.Println("Image rendered and saved as PNG")
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
│   │   ├── tag/            # DICOM tags (5338 standard tags)
│   │   ├── vr/             # Value representations (35 VRs)
│   │   ├── vm/             # Value multiplicities
│   │   ├── element/        # DICOM elements (all VR types)
│   │   ├── dataset/        # Dataset, sequences, file meta info
│   │   ├── parser/         # DICOM file parsing
│   │   ├── writer/         # DICOM file writing
│   │   ├── dict/           # Tag dictionary
│   │   ├── transfer/       # Transfer syntaxes
│   │   ├── uid/            # Standard UIDs (1965 UIDs)
│   │   ├── charset/        # Character encoding (30+ charsets)
│   │   ├── serialization/  # JSON/XML conversion
│   │   ├── anonymizer/     # Anonymization profiles
│   │   ├── endian/         # Byte order handling
│   │   └── testutil/       # Test utilities
│   ├── imaging/            # Image processing
│   │   ├── codec/          # Image codecs (RLE, native, transcoder)
│   │   ├── lut/            # Lookup tables (Modality, VOI, Palette)
│   │   ├── render/         # Image rendering pipeline
│   │   └── reconstruction/ # Image reconstruction
│   ├── network/            # DICOM networking
│   │   ├── pdu/            # Protocol Data Units
│   │   ├── dimse/          # DIMSE messages
│   │   ├── association/    # Association management
│   │   ├── client/         # SCU (Service Class User)
│   │   ├── server/         # SCP (Service Class Provider)
│   │   ├── service/        # DIMSE services
│   │   ├── transport/      # Network transport
│   │   └── status/         # DIMSE status codes
│   ├── sr/                 # Structured Reports
│   ├── printing/           # DICOM printing
│   └── io/                 # I/O operations
│       └── buffer/         # Byte buffer abstractions
├── cmd/                    # Command-line tools
│   ├── dicominfo/          # Display DICOM file information
│   ├── dicomdump/          # Dump DICOM file contents
│   └── dicom2json/         # Convert DICOM to JSON
├── examples/               # Usage examples
│   ├── read_dicom/         # Reading DICOM files
│   ├── write_dicom/        # Writing DICOM files
│   ├── json_conversion/    # JSON/XML serialization
│   └── anonymize/          # Anonymization examples
├── test-data/              # Test DICOM files
├── tools/                  # Code generation tools
│   ├── generate_tags/      # Generate tag constants
│   ├── generate_uids/      # Generate UID constants
│   └── generate_dict/      # Generate dictionary data
├── BENCHMARKS.md           # Performance benchmarks
├── TODO.md                 # Development roadmap (Chinese)
└── CLAUDE.md               # Development guide for AI assistants
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

**Current Phase**: Advanced Features (~90% Complete)

**Working Features**:
- ✅ Complete DICOM file I/O (read/write with all transfer syntaxes)
- ✅ Multi-frame and single-frame image support
- ✅ Character encoding (30+ encodings with auto-detection)
- ✅ Fragment sequences (compressed pixel data)
- ✅ Structured Reports (SR) with hierarchical content
- ✅ Dataset manipulation with rich API
- ✅ JSON/XML serialization (DICOM Part 18 compliant)
- ✅ Anonymization with configurable profiles
- ✅ Image processing (LUT, windowing, color conversion, rendering)
- ✅ DICOM networking (All C-services + N-services, TLS secure connections)
- ✅ RLE codec and transcoding framework
- ✅ DICOM printing (Film Session, Film Box, Image Box)

**In Progress**:
- [ ] Advanced image codecs (JPEG, JPEG-LS, JPEG2000 decompression)
- [ ] Extended DICOM networking features (role negotiation, SOP class extended negotiation)
- [ ] Additional anonymization profiles (enhanced privacy options)

**Command-Line Tools**:
- `dicominfo` - Display DICOM file metadata
- `dicomdump` - Dump complete DICOM file structure
- `dicom2json` - Convert DICOM to JSON format

See [TODO.md](TODO.md) (Chinese) for complete task list and detailed progress tracking.
