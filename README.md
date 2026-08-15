# go-dicom

A pure Go implementation of the DICOM (Digital Imaging and Communications in Medicine) standard, ported from the [fo-dicom](https://github.com/fo-dicom/fo-dicom) C# library.

[![Go Reference](https://pkg.go.dev/badge/github.com/cocosip/go-dicom.svg)](https://pkg.go.dev/github.com/cocosip/go-dicom)
[![Go Report Card](https://goreportcard.com/badge/github.com/cocosip/go-dicom)](https://goreportcard.com/report/github.com/cocosip/go-dicom)
[![License](https://img.shields.io/badge/license-MS--PL-blue.svg)](LICENSE)

## Release Status

`go-dicom` is a released DICOM library.
This README describes the supported library surface rather than a development roadmap.
Known parity gaps and the staged plan for addressing them are tracked in
[fo-dicom Capability Gap Analysis](docs/FO_DICOM_GAP_ANALYSIS.md)
([简体中文](docs/FO_DICOM_GAP_ANALYSIS.zh-CN.md)).

## Features

### Core Capabilities ✅

- ✅ **DICOM File I/O** - Read and write DICOM files with full standard compliance
- ✅ **DICOMDIR Media** - Create, read, traverse, repair, and save media directories with optional image icons
- ✅ **Transfer Syntax Support** - Explicit/Implicit VR, Big/Little Endian
- ✅ **Multi-Frame Images** - Full support for multi-frame and video DICOM files
- ✅ **Character Encoding** - 30+ character sets with auto-detection (UTF-8, Latin, Chinese, Japanese, Korean, Arabic, etc.)
- ✅ **Fragment Sequences** - Read and write compressed pixel data; compressed transfer syntaxes are handled by [go-dicom-codecs](https://github.com/cocosip/go-dicom-codecs)
- ✅ **Structured Reports (SR)** - Parse and create SR documents with hierarchical content
- ✅ **Dataset Operations** - Rich API for accessing and manipulating DICOM elements
- ✅ **JSON/XML Serialization** - Export/Import DICOM data to JSON (Part 18) and XML formats
- ✅ **Anonymization** - Remove/Replace patient identifiable information with configurable profiles
- ✅ **Pixel Data Processing** - Access raw pixel data, color space conversion, LUT operations, image rendering
- ✅ **DICOM Networking** - Complete DIMSE services (C-ECHO/STORE/FIND/MOVE/GET + N-CREATE/GET/SET/DELETE/ACTION/EVENT-REPORT), TLS support
- ✅ **Image Processing** - Rendering pipeline, windowing, LUT operations, color space conversion
- ✅ **DICOM Printing** - Film Session, Film Box, Image Box, Printer status management

### Included Capabilities

- [x] **Core DICOM data types**
  - [x] Tag (5347 standard tags + private tag support)
  - [x] VR (35 value representations with validation)
  - [x] VM (15 value multiplicities)
  - [x] Element (string, numeric, binary, date, person name types)
  - [x] Dataset & Sequence (full support with lazy loading)
  - [x] Dictionary (tag/keyword lookup with global default instance)
  - [x] UID (1928 standard UIDs + 59 private UIDs)
  - [x] Transfer Syntax (15+ syntaxes including JPEG and MPEG)
  - [x] Character Set (30+ encodings with auto-detection)

- [x] **DICOM file reading**
  - [x] Explicit/Implicit VR parsing
  - [x] Sequence parsing (defined/undefined length)
  - [x] Fragment sequence support (compressed images)
  - [x] Multi-frame image support
  - [x] ReadOptions: SkipLargeTags, ReadLargeOnDemand, ReadAll
  - [x] FileFormat detection: DICOM3, DICOM3NoPreamble, ACR-NEMA
  - [x] Large object handling with configurable thresholds
  - [x] Automatic character set detection and conversion
  - [x] Byte order handling (Little/Big Endian)

- [x] **DICOM file writing**
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
  - [x] Fragment sequences (compressed pixel data)
  - [x] Video DICOM (MPEG2)
  - [x] RGB color images (planar and interleaved)
  - [x] Structured Reports (SR) with hierarchical content
  - [x] Modality LUT Sequences
  - [x] Character set variants (17+ tested encodings)

- [x] **DICOMDIR Media**
  - [x] Patient/study/series/instance grouping and traversal
  - [x] Strict validation and bounded compatible offset recovery
  - [x] Two-pass Explicit/Implicit VR Little Endian writing
  - [x] Optional pure-Go 8-bit grayscale icon generation
  - [x] fo-dicom fixture and cross-library interoperability verification

- [x] **JSON/XML Serialization**
  - [x] DICOM JSON Model (Part 18 compliant)
  - [x] Native XML format
  - [x] Bulkdata handling with base64 encoding
  - [x] Pretty-print options
  - [x] PersonName component groups support
  - [x] Sequence nesting support

- [x] **Anonymization**
  - [x] Basic anonymization profile (patient identifiable information)
  - [x] Custom anonymization rules (Remove, Replace, Keep)
  - [x] Patient/Study/Series level anonymization
  - [x] Date shifting and UID remapping
  - [x] Recursive sequence anonymization

- [x] **Imaging Support**
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

- [x] **Structured Reports**
  - [x] Typed content items (TEXT, NUM, CODE, CONTAINER, PNAME, DATE, TIME, DATETIME, UIDREF)
  - [x] IMAGE, COMPOSITE, WAVEFORM, SCOORD, and TCOORD values
  - [x] Hierarchical structure with parent-child relationships
  - [x] Code Value, Long Code Value, and URN Code Value support
  - [x] Measured values with units
  - [x] Referenced SOP instances
  - [x] Relationship types (CONTAINS, HAS OBS CONTEXT, INFERRED FROM, etc.)
  - [x] Recursive semantic validation with nested content paths
  - [x] Validated file and stream I/O with preserved File Meta and transfer syntax

- [x] **DICOM Networking**
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
  - [x] Asynchronous Operations Window negotiation with request throttling
  - [x] Advanced role negotiation (SCP/SCU Role Selection)
  - [x] Extended negotiation items (SOP Class Extended Negotiation)
  - [x] User Identity negotiation (Username, Username/Password, Kerberos, SAML, JWT)
  - [x] ServiceApplicationInfo helper type

- [x] **Image Codecs**
  - [x] Native codecs (uncompressed data - Explicit/Implicit VR, Little/Big Endian)
  - [x] Transcoder framework for format conversion between transfer syntaxes
  - [x] Codec registry and plugin architecture
  - [x] Compressed transfer syntax codecs supplied by [go-dicom-codecs](https://github.com/cocosip/go-dicom-codecs)

  **Note**: `go-dicom` supplies the codec registry, transcoder, and DICOM encapsulation support.
  Add the selected [go-dicom-codecs](https://github.com/cocosip/go-dicom-codecs) codec package as a blank import to register its compressed transfer syntax.

- [x] **DICOM Printing**
  - [x] Film Session management
  - [x] Film Box configuration
  - [x] Image Box handling
  - [x] Presentation LUT
  - [x] Print job creation
  - [x] Printer status management

## Performance

go-dicom includes benchmarks for dataset operations, parsing, writing, I/O, and image conversion. See [BENCHMARKS.md](BENCHMARKS.md) for the benchmark scope, historical snapshots, and result provenance. Pull requests publish their full benchmark output as a `benchmark-results` workflow artifact.

Benchmark numbers vary by commit, Go version, CPU, operating system, and DICOM workload. Use the command below to collect a result for the environment being evaluated.

Run benchmarks yourself:
```bash
go test -bench='.' -benchmem -run='^$' ./pkg/dicom/... ./pkg/io/... ./pkg/imaging/...
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

### Creating and Updating Sequences

```go
package main

import (
    "fmt"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func mustAdd(ds *dataset.Dataset, elem element.Element) {
    if err := ds.Add(elem); err != nil {
        log.Fatal(err)
    }
}

func main() {
    ds := dataset.New()

    // Build ReferencedStudySequence with one study item.
    studies := dataset.NewSequence(tag.ReferencedStudySequence)
    study := dataset.New()
    mustAdd(study, element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

    // A sequence item is itself a Dataset, so sequences can be nested.
    series := dataset.NewSequence(tag.ReferencedSeriesSequence)
    seriesItem := dataset.New()
    mustAdd(seriesItem, element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.4.5.1"}))
    mustAdd(seriesItem, element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.1.1"}))
    series.AddItem(seriesItem)
    mustAdd(study, series)

    studies.AddItem(study)
    mustAdd(ds, studies)

    // Retrieve the sequence with a typed accessor and update its first item.
    referencedStudies, err := ds.GetSequence(tag.ReferencedStudySequence)
    if err != nil {
        log.Fatal(err)
    }

    firstStudy := referencedStudies.GetItem(0)
    if firstStudy == nil {
        log.Fatal("ReferencedStudySequence has no items")
    }
    if err := firstStudy.AddOrUpdate(element.NewString(tag.StudyDescription, vr.LO, []string{"Follow-up study"})); err != nil {
        log.Fatal(err)
    }

    referencedSeries, err := firstStudy.GetSequence(tag.ReferencedSeriesSequence)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Referenced studies: %d, referenced series: %d\n", referencedStudies.Count(), referencedSeries.Count())

    // Remove an item when it is no longer needed.
    referencedSeries.RemoveItem(0)
    fmt.Printf("Referenced series after removal: %d\n", referencedSeries.Count())
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

    "github.com/cocosip/go-dicom/pkg/sr"
)

func main() {
    report, err := sr.Open("sr_report.dcm")
    if err != nil {
        log.Fatal(err)
    }

    children, err := report.Children()
    if err != nil {
        log.Fatal(err)
    }
    for i, item := range children {
        valueType, err := item.ValueType()
        if err != nil {
            // By-reference items intentionally have no Value Type.
            fmt.Printf("Item %d: by-reference\n", i)
            continue
        }
        fmt.Printf("Item %d: %s\n", i, valueType)
    }
}
```

### Handling Compressed Images

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
    file, err := os.Open("compressed.dcm")
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

        // To decode or transcode this data, add the selected go-dicom-codecs
        // package as a blank import so it registers its codec, then use
        // imaging/codec.NewTranscoder with the required transfer syntaxes.
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

### DICOM Networking - Advanced Association Negotiation

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/network/association"
    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    const ctImageStorage = "1.2.840.10008.5.1.4.1.1.2"

    c := client.New(
        client.WithCallingAE("GO-SCU"),
        client.WithCalledAE("ANY-SCP"),
        client.WithAsynchronousOperations(4, 2),
        client.WithExtendedNegotiation(
            association.NewExtendedNegotiation(ctImageStorage, []byte{1, 1, 0}),
        ),
        client.WithUserIdentity(
            association.NewUserIdentityJWT([]byte(os.Getenv("DICOM_JWT")), true),
        ),
    )
    c.AddPresentationContextWithRoles(
        ctImageStorage,
        true, // request the SCU role
        true, // request the SCP role for C-GET C-STORE sub-operations
        "1.2.840.10008.1.2.1",
    )

    ctx := context.Background()
    if err := c.Connect(ctx, "localhost", 11112); err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    accepted := c.GetAssociation()
    log.Printf("accepted async window: %d/%d",
        accepted.AsynchronousOperations.MaxInvokedOperations,
        accepted.AsynchronousOperations.MaxPerformedOperations,
    )
}
```

When a positive User Identity response is requested, the client rejects an
association that omits it by default. Use
`client.WithRequireSuccessfulUserIdentityNegotiation(false)` only when an
empty response is an intentional compatibility requirement.

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
        MinVersion:         tls.VersionTLS12,
    }

    // Create DICOM client with TLS
    c := client.New(
        client.WithCallingAE("SECURE-SCU"),
        client.WithCalledAE("SECURE-SCP"),
        client.WithTLSConfig(tlsConfig),
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

### DICOM Networking - C-STORE SCU (Send Files)

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/dicom/transfer"
    "github.com/cocosip/go-dicom/pkg/dicom/uid"
    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    c := client.New(
        client.WithCallingAE("GO-SCU"),
        client.WithCalledAE("STORE-SCP"),
    )

    // Propose CT Image Storage with common transfer syntaxes
    c.AddPresentationContext(
        uid.CTImageStorage.UID(),
        transfer.ExplicitVRLittleEndian.UID().UID(),
        transfer.ImplicitVRLittleEndian.UID().UID(),
    )

    ctx := context.Background()
    if err := c.Connect(ctx, "localhost", 11112); err != nil {
        log.Fatalf("connect: %v", err)
    }
    defer c.Close()

    // Read DICOM file
    result, err := parser.ParseFile("scan.dcm")
    if err != nil {
        log.Fatalf("open: %v", err)
    }

    // Send via C-STORE
    if err := c.CStore(ctx, result.Dataset); err != nil {
        log.Fatalf("C-STORE: %v", err)
    }
    log.Println("C-STORE successful")
}
```

The client prefers the Dataset's original transfer syntax when the peer accepts
it. Otherwise it transcodes a copy to the first accepted syntax supported by the
registered codecs. Parsed Datasets retain their source syntax automatically;
programmatically constructed Datasets with Pixel Data must set it with
`dataset.NewWithTransferSyntax`.

### DICOM Networking - C-FIND SCU (Query)

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/client"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
    "github.com/cocosip/go-dicom/pkg/uid"
)

func main() {
    c := client.New(
        client.WithCallingAE("GO-SCU"),
        client.WithCalledAE("QR-SCP"),
    )
    c.AddPresentationContext(
        uid.StudyRootQueryRetrieveInformationModelFind,
        uid.ExplicitVRLittleEndian,
    )

    ctx := context.Background()
    if err := c.Connect(ctx, "localhost", 11112); err != nil {
        log.Fatalf("connect: %v", err)
    }
    defer c.Close()

    // Build query identifier
    query := dataset.New()
    _ = query.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{"STUDY"}))
    _ = query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))
    _ = query.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{""})) // empty = return all

    responses, err := c.CFind(ctx, dimse.QueryRetrieveLevelStudy, query)
    if err != nil {
        log.Fatalf("C-FIND: %v", err)
    }
    for _, resp := range responses {
        log.Printf("Found study: %v", resp.DataDataset())
    }
}
```

### DICOM Networking - C-MOVE SCU (Retrieve to third-party AE)

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/client"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
    "github.com/cocosip/go-dicom/pkg/uid"
)

func main() {
    c := client.New(
        client.WithCallingAE("GO-SCU"),
        client.WithCalledAE("QR-SCP"),
    )
    c.AddPresentationContext(
        uid.StudyRootQueryRetrieveInformationModelMove,
        uid.ExplicitVRLittleEndian,
    )

    ctx := context.Background()
    if err := c.Connect(ctx, "localhost", 11112); err != nil {
        log.Fatalf("connect: %v", err)
    }
    defer c.Close()

    // Identifier dataset: which study to move
    identifier := dataset.New()
    _ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.840.999.1"}))

    // "DEST-AE" must be known to the SCP (registered in its AE routing table)
    responses, err := c.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST-AE", identifier)
    if err != nil {
        log.Fatalf("C-MOVE: %v", err)
    }
    for _, resp := range responses {
        log.Printf("C-MOVE progress: completed=%d failed=%d",
            resp.NumberOfCompletedSubOperations(),
            resp.NumberOfFailedSubOperations())
    }
}
```

### DICOM Networking - C-GET SCU (Retrieve to self)

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/client"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
    "github.com/cocosip/go-dicom/pkg/uid"
)

func main() {
    c := client.New(
        client.WithCallingAE("GO-SCU"),
        client.WithCalledAE("QR-SCP"),
    )
    // Propose both C-GET and the storage SOP classes you expect to receive
    c.AddPresentationContext(
        uid.StudyRootQueryRetrieveInformationModelGet,
        uid.ExplicitVRLittleEndian,
    )
    c.AddPresentationContext(
        uid.CTImageStorage,
        uid.ExplicitVRLittleEndian,
    )

    ctx := context.Background()
    if err := c.Connect(ctx, "localhost", 11112); err != nil {
        log.Fatalf("connect: %v", err)
    }
    defer c.Close()

    identifier := dataset.New()
    _ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.840.999.1"}))

    // C-GET: SCP sends C-STORE sub-operations back over the same association.
    // The client's CStoreHandler is called for each received instance.
    c.SetCStoreHandler(func(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
        ds := req.DataDataset()
        log.Printf("Received instance: %v", ds)
        return dimse.NewCStoreResponseFromRequest(req, 0x0000), nil
    })

    responses, err := c.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier)
    if err != nil {
        log.Fatalf("C-GET: %v", err)
    }
    for _, resp := range responses {
        log.Printf("C-GET progress: completed=%d failed=%d",
            resp.NumberOfCompletedSubOperations(),
            resp.NumberOfFailedSubOperations())
    }
}
```

### DICOM Networking - SCP Server (C-MOVE and C-GET handlers)

The SCP handlers for C-MOVE and C-GET use an **operation object** that implements
`SubOperationResponder`. This enables streaming progress — each call to
`op.SendPending` immediately writes a response to the wire, equivalent to
fo-dicom's `IAsyncEnumerable<DicomCMoveResponse>` pattern.

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/network/server"
    "github.com/cocosip/go-dicom/pkg/network/service"
)

func main() {
    srv := server.New(server.WithPort(11112))

    // ── C-MOVE SCP ────────────────────────────────────────────────────────────
    // CMoveOperation embeds SubOperationResponder and exposes:
    //   op.QueryLevel()       dimse.QueryRetrieveLevel
    //   op.MoveDestination()  string  (destination AE title)
    //   op.Identifier()       *dataset.Dataset
    //   op.SendPending(remaining, completed, failed, warning uint16) error
    //   op.SendSuccess() / op.SendWarning() / op.SendFailure(code) error
    srv.SetCMoveHandler(func(ctx context.Context, op service.CMoveOperation) error {
        destination := op.MoveDestination()
        log.Printf("C-MOVE to %s level=%s", destination, op.QueryLevel())

        // Look up destination address from your AE routing table
        // destHost, destPort := aeRoutes[destination]
        // ... connect and send C-STORE to destination ...

        // Report progress after each sub-operation
        _ = op.SendPending(9, 1, 0, 0)  // 1 sent, 9 remaining
        _ = op.SendPending(0, 10, 0, 0) // all done

        return op.SendSuccess()
    })

    // ── C-GET SCP ─────────────────────────────────────────────────────────────
    // CGetOperation embeds SubOperationResponder and exposes:
    //   op.QueryLevel()   dimse.QueryRetrieveLevel
    //   op.Identifier()   *dataset.Dataset
    //   op.SendCStore(ctx, ds) (*dimse.CStoreResponse, error)  ← same association
    //   op.SendPending / op.SendSuccess / op.SendWarning / op.SendFailure
    srv.SetCGetHandler(func(ctx context.Context, op service.CGetOperation) error {
        log.Printf("C-GET level=%s", op.QueryLevel())

        // For each matching instance, push it back to the SCU
        // ds, _ := loadInstance(...)
        // resp, err := op.SendCStore(ctx, ds)

        // Report progress after each push
        _ = op.SendPending(4, 1, 0, 0)
        _ = op.SendPending(0, 5, 0, 0)

        return op.SendSuccess()
    })

    ctx := context.Background()
    if err := srv.ListenAndServe(ctx); err != nil {
        log.Fatal(err)
    }
}
```

**`SubOperationResponder` interface** (embedded by both `CMoveOperation` and `CGetOperation`):

| Method | Status code | Description |
|---|---|---|
| `SendPending(remaining, completed, failed, warning uint16) error` | 0xFF00 | Progress update after each sub-operation |
| `SendSuccess() error` | 0x0000 | All sub-operations completed successfully |
| `SendWarning() error` | 0xB000 | Completed but some sub-operations failed |
| `SendFailure(code uint16) error` | custom | Fatal failure (e.g. 0xA801 = Move Destination Unknown) |

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
go test ./cmd/... ./examples/... ./pkg/... ./tools/...

# Run parser tests
go test -v ./pkg/dicom/parser

# Run specific test
go test -v ./pkg/dicom/parser -run TestMultiFrame

# Run with coverage
go test -cover ./cmd/... ./examples/... ./pkg/... ./tools/...
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
│   │   ├── tag/            # DICOM tags (5347 standard tags)
│   │   ├── vr/             # Value representations (35 VRs)
│   │   ├── vm/             # Value multiplicities
│   │   ├── element/        # DICOM elements (all VR types)
│   │   ├── dataset/        # Dataset, sequences, file meta info
│   │   ├── parser/         # DICOM file parsing
│   │   ├── writer/         # DICOM file writing
│   │   ├── dict/           # Tag dictionary
│   │   ├── transfer/       # Transfer syntaxes
│   │   ├── uid/            # Standard UIDs (1928 UIDs)
│   │   ├── charset/        # Character encoding (30+ charsets)
│   │   ├── serialization/  # JSON/XML conversion
│   │   ├── anonymizer/     # Anonymization profiles
│   │   ├── endian/         # Byte order handling
│   │   └── testutil/       # Test utilities
│   ├── imaging/            # Image processing
│   │   ├── codec/          # Image codecs (native, transcoder)
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
│   ├── data/2026b/         # Pinned fo-dicom dictionary XML inputs
│   └── generate_dicom/     # Generate tags, UIDs, and dictionaries
├── BENCHMARKS.md           # Performance benchmarks
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
go test ./cmd/... ./examples/... ./pkg/... ./tools/...

# Run benchmarks
go test -bench='.' -benchmem ./pkg/dicom/...
```

### Regenerating DICOM Data

The repository pins fo-dicom's standard and manually maintained private
dictionary XML files under `tools/data/2026b`. Regenerate all XML-derived data
with the unified tool:

```powershell
go run ./tools/generate_dicom `
  -standard "tools/data/2026b/DICOM Dictionary.xml" `
  -private "tools/data/2026b/Private Dictionary.xml" `
  -root "."
```

The command derives in one run the standard Tag constants, standard UID
constants, standard dictionary, and private dictionary from one pinned
baseline. To update the baseline, copy both XML files from the selected
fo-dicom release into a new versioned directory, run the command with those
paths, and commit the XML and all four generated outputs together. Private UIDs
in `pkg/dicom/uid/uids_private.go` are not present in either XML input and are
therefore outside this generator.

### Code Quality

```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Run go vet
go vet ./...
```

## Relationship to fo-dicom (C#)

This project is a Go port inspired by the fo-dicom library:

1. ✅ Core data types (Tag, VR, Dictionary)
2. ✅ Data structures (Element, Dataset)
3. ✅ I/O capabilities (Reader, Writer)
4. ✅ Advanced features (Networking, Codecs)

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
