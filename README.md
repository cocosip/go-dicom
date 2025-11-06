# go-dicom

A pure Go implementation of the DICOM (Digital Imaging and Communications in Medicine) standard, ported from the [fo-dicom](https://github.com/fo-dicom/fo-dicom) C# library.

[![Go Reference](https://pkg.go.dev/badge/github.com/cocosip/go-dicom.svg)](https://pkg.go.dev/github.com/cocosip/go-dicom)
[![Go Report Card](https://goreportcard.com/badge/github.com/cocosip/go-dicom)](https://goreportcard.com/report/github.com/cocosip/go-dicom)
[![License](https://img.shields.io/badge/license-MS--PL-blue.svg)](LICENSE)

## ⚠️ Work in Progress

This library is currently under active development. The API is not stable and may change significantly.

## Features

### Current Status (~60% Complete)

- [x] **Core DICOM data types**
  - [x] Tag (5338 standard tags + private tag support)
  - [x] VR (35 value representations with validation)
  - [x] VM (15 value multiplicities)
  - [x] Element (string, numeric, binary types)
  - [x] Dataset & Sequence (full support)
  - [x] Dictionary (tag/keyword lookup, Default global dictionary)
  - [x] UID (1965 standard UIDs)
  - [x] Transfer Syntax (15+ syntaxes)
  - [x] Character Set (30+ encodings)

- [x] **DICOM file reading**
  - [x] Explicit/Implicit VR parsing
  - [x] Sequence parsing (defined/undefined length)
  - [x] **ReadOption**: SkipLargeTags, ReadLargeOnDemand, ReadAll
  - [x] **FileFormat** detection: DICOM3, DICOM3NoPreamble, etc.
  - [x] Large object handling with configurable thresholds
  - [x] Tag.DictionaryEntry() for metadata lookup

- [x] **DICOM file writing**
  - [x] Explicit/Implicit VR writing
  - [x] Auto-generated File Meta Information
  - [x] **DicomWriteOptions**:
    - ExplicitLengthSequences/Items
    - KeepGroupLengths
    - LargeObjectSize
  - [x] Group length auto-filtering
  - [x] Explicit/undefined length sequences

- [ ] JSON serialization (DICOM JSON Model)
- [ ] DICOM networking (DIMSE services)
- [ ] Image codec support

See [TODO.md](TODO.md) for detailed development roadmap.

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

    // Get patient name
    patientName, exists := result.Dataset.GetString(tag.PatientName)
    if exists {
        fmt.Printf("Patient Name: %s\n", patientName)
    }

    // Get file format
    fmt.Printf("Format: %s\n", result.Format)
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
    fmt.Printf("Patient: %s\n", patientName)
}
```

### Writing a DICOM File

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

    // Add elements
    ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
    ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
    ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250106"}))

    // Write to file
    file, err := os.Create("output.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write with default options
    if err := writer.Write(file, ds); err != nil {
        log.Fatal(err)
    }
}
```

### Writing with Options

```go
package main

import (
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/transfer"
    "github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
    ds := dataset.New()
    // ... add elements ...

    file, err := os.Create("output.dcm")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write with custom options
    err = writer.Write(file, ds,
        writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian),
        writer.WithExplicitLengthSequences(),  // Use explicit lengths
        writer.WithLargeObjectSize(1024*1024), // 1MB threshold
    )
    if err != nil {
        log.Fatal(err)
    }
}
```

## Development

### Prerequisites

- Go 1.21 or later
- golangci-lint (for linting)

### Building

```bash
# Build all packages
make build

# Or using go directly
go build ./...
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run tests for a specific package
make test-pkg PKG=./pkg/dicom
```

### Code Quality

```bash
# Format code
make fmt

# Run linter
make lint

# Run go vet
make vet
```

## Project Structure

```
go-dicom/
├── pkg/
│   ├── dicom/          # Core DICOM types (Tag, VR, Element, Dataset)
│   ├── io/             # I/O operations (Reader, Writer, Buffer)
│   │   ├── buffer/     # Byte buffer abstractions
│   │   ├── reader/     # DICOM file reading
│   │   └── writer/     # DICOM file writing
│   └── transfer/       # Transfer syntax definitions
├── internal/           # Internal utilities
├── cmd/                # Command-line tools
│   └── dicominfo/      # Display DICOM file information
├── examples/           # Usage examples
├── tests/              # Integration tests
├── CLAUDE.md           # Claude Code guidance
├── TODO.md             # Development roadmap
└── Makefile            # Build automation
```

## Migration from fo-dicom (C#)

This project is a port of the fo-dicom library from C# to Go. The migration follows an incremental approach:

1. Core data types (Tag, VR, Dictionary)
2. Data structures (Element, Dataset)
3. I/O capabilities (Reader, Writer)
4. Advanced features (Networking, Imaging)

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
4. Run `make all` before submitting PRs

## Status

**Current Phase**: Project Infrastructure Setup

**Next Steps**: Implementing core DICOM data types (VR, Tag, Dictionary)

See [TODO.md](TODO.md) for complete task list and progress tracking.
