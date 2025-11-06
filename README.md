# go-dicom

A pure Go implementation of the DICOM (Digital Imaging and Communications in Medicine) standard, ported from the [fo-dicom](https://github.com/fo-dicom/fo-dicom) C# library.

[![Go Reference](https://pkg.go.dev/badge/github.com/yourusername/go-dicom.svg)](https://pkg.go.dev/github.com/yourusername/go-dicom)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/go-dicom)](https://goreportcard.com/report/github.com/yourusername/go-dicom)
[![License](https://img.shields.io/badge/license-MS--PL-blue.svg)](LICENSE)

## ⚠️ Work in Progress

This library is currently under active development. The API is not stable and may change significantly.

## Features

### Current Status

- [ ] Core DICOM data types (Tag, VR, Element, Dataset)
- [ ] DICOM file reading
- [ ] DICOM file writing
- [ ] Transfer syntax support
- [ ] JSON serialization (DICOM JSON Model)
- [ ] DICOM networking (DIMSE services)
- [ ] Image codec support

See [TODO.md](TODO.md) for detailed development roadmap.

## Installation

```bash
go get github.com/yourusername/go-dicom
```

## Quick Start

### Reading a DICOM File

```go
package main

import (
    "fmt"
    "log"

    "github.com/yourusername/go-dicom/pkg/dicom"
)

func main() {
    // Open DICOM file
    file, err := dicom.Open("example.dcm")
    if err != nil {
        log.Fatal(err)
    }

    // Access dataset
    dataset := file.Dataset

    // Get patient name
    patientName, err := dataset.GetString(dicom.TagPatientName)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Patient Name: %s\n", patientName)
}
```

### Creating a DICOM File

```go
package main

import (
    "log"

    "github.com/yourusername/go-dicom/pkg/dicom"
)

func main() {
    // Create new dataset
    dataset := dicom.NewDataset()

    // Add elements
    dataset.AddString(dicom.TagPatientName, "Doe^John")
    dataset.AddString(dicom.TagPatientID, "12345")
    dataset.AddString(dicom.TagStudyDate, "20250106")

    // Create DICOM file
    file := dicom.NewFile(dataset)

    // Save to disk
    if err := file.Save("output.dcm"); err != nil {
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
