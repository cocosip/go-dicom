# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based DICOM (Digital Imaging and Communications in Medicine) library being ported from the C# fo-dicom library. The source C# code is located in `fo-dicom-code/` directory. This is an incremental, feature-by-feature migration project that requires translating C# concepts to idiomatic Go.

**Key Context:**
- Source: fo-dicom v5.2.4 (C# .NET Standard 2.0 library)
- Target: Go DICOM library
- Migration Strategy: Incremental, small-feature, layer-by-layer development
- License: Microsoft Public License (MS-PL)

## Source Code Architecture (fo-dicom)

### Core Components

1. **Core DICOM Data Structures** (`fo-dicom-code/` root)
   - `DicomTag.cs` - DICOM tag representation (group, element pairs)
   - `DicomVR.cs` - Value Representation (data type system)
   - `DicomDataset.cs` (~1700 lines) - Main DICOM dataset container
   - `DicomElement.cs` (~2000 lines) - Individual DICOM elements
   - `DicomFile.cs` - DICOM file handling with metadata
   - `DicomDictionary.cs` - Tag/VR definitions
   - `DicomTransferSyntax.cs` - Encoding methods (explicit/implicit VR, endianness, compression)

2. **Generated Code** (from T4 templates)
   - `DicomTagGenerated.cs` (~16k lines) - All standard DICOM tags
   - `DicomUIDGenerated.cs` (~7k lines) - Standard UIDs
   - `DicomAnonymizerGenerated.cs` - Anonymization rules

3. **IO Layer** (`fo-dicom-code/IO/`)
   - `Buffer/` - Memory buffer abstractions (IByteBuffer implementations)
     - CompositeByteBuffer, EndianByteBuffer, MemoryByteBuffer, StreamByteBuffer, etc.
   - `Reader/` - DICOM parsing
     - `DicomReader.cs` - Main reader with observer pattern
     - `DicomFileReader.cs` - File-specific reading
     - `DicomReaderObserver` interfaces
   - `Writer/` - DICOM serialization
     - `DicomWriter.cs` - Main writer
     - `DicomFileWriter.cs` - File-specific writing
   - `IByteSource` / `IByteTarget` - Stream abstractions
   - `Endian.cs` - Byte order conversions

4. **Network Layer** (`fo-dicom-code/Network/`)
   - `DicomService.cs` (~2400 lines) - Base DICOM networking service
   - `DicomServer.cs` - DICOM SCP (Service Class Provider)
   - `Client/` - DICOM SCU (Service Class User) implementations
   - PDU (Protocol Data Unit) handling
   - DIMSE (DICOM Message Service Element) operations:
     - C-ECHO, C-FIND, C-GET, C-MOVE, C-STORE
     - N-CREATE, N-GET, N-SET, N-DELETE, N-ACTION, N-EVENT-REPORT
   - `DicomAssociation.cs` - DICOM association management
   - `DicomPresentationContext.cs` - Negotiation contexts

5. **Imaging** (`fo-dicom-code/Imaging/`)
   - `DicomImage.cs` - Image representation
   - `DicomPixelData.cs` - Pixel data handling
   - `Codec/` - Image compression/decompression
   - `Render/` - Image rendering pipeline
   - `LUT/` - Lookup tables for windowing/leveling
   - `PhotometricInterpretation.cs` - Color space handling

6. **Serialization** (`fo-dicom-code/Serialization/`)
   - `JsonDicomConverter.cs` - JSON serialization (DICOM JSON model)
   - `DicomXML.cs` - XML serialization

7. **Memory Management** (`fo-dicom-code/Memory/`)
   - `IMemoryProvider` - Abstraction for memory allocation
   - `ArrayPoolMemory` - Pooled memory implementation

## Migration Strategy

### Development Approach

Since fo-dicom is a large library (~50+ files with complex dependencies), follow this incremental approach:

1. **Start with Foundation Layers**
   - Begin with core types: DicomTag, DicomVR, DicomVRCode
   - Implement basic buffer abstractions
   - Add DicomDictionary (can use simplified version initially)

2. **Build Data Layer**
   - DicomElement and its variants
   - DicomDataset
   - Implement basic validation

3. **Add IO Capabilities**
   - Byte sources/targets
   - DicomReader with observer pattern
   - DicomWriter
   - Transfer syntax support

4. **Network Layer** (Complex - save for later)
   - Implement PDU structures
   - Association management
   - DIMSE operations one at a time

5. **Imaging Features** (Optional - depends on requirements)
   - Pixel data handling
   - Codec support

### C# to Go Translation Guidelines

**Type Mappings:**
- `ushort` → `uint16`
- `uint` → `uint32`
- `string` → `string`
- `byte[]` → `[]byte`
- `List<T>` → `[]T` or consider `container/list`
- `Dictionary<K,V>` → `map[K]V`
- `Stream` → `io.Reader`/`io.Writer`/`io.ReadWriteSeeker`

**Pattern Translations:**
- C# properties → Go getter methods or exported fields
- C# events → Go channels or callback functions
- C# async/await → Go goroutines and channels
- C# LINQ → Go iteration with explicit loops or functional helpers
- C# interfaces → Go interfaces (usually smaller, more focused)
- C# nullable reference types → Go pointers or zero values
- Observer pattern → Interface with callbacks or channels

**Dependency Injection:**
- fo-dicom uses Microsoft.Extensions.DependencyInjection
- In Go, use constructor injection with interfaces
- Consider using a simple DI container or manual wiring

**Memory Management:**
- C# uses IMemoryProvider/ArrayPool for performance
- In Go, leverage built-in GC, use sync.Pool for hot paths
- Be mindful of buffer reuse in IO operations

**Error Handling:**
- C# exceptions → Go error returns
- Validation exceptions → return errors with descriptive messages
- Use error wrapping (`fmt.Errorf` with `%w`) for context

## Development Commands

Since this is a fresh Go project being created from C# source, you'll need to:

### Initial Setup
```bash
# Initialize Go module (if not already done)
go mod init github.com/yourusername/go-dicom

# Add dependencies as needed
go get <dependency>
```

### Building
```bash
# Build the project
go build ./...

# Build a specific package
go build ./pkg/dicom
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests in a specific package
go test ./pkg/dicom

# Run a specific test
go test -run TestDicomTag ./pkg/dicom

# Run tests with verbose output
go test -v ./...
```

### Code Quality
```bash
# Format code
go fmt ./...

# Run linter (requires golangci-lint)
golangci-lint run

# Run static analysis
go vet ./...
```

**IMPORTANT: All code must comply with golangci-lint standards.**
- Run `golangci-lint run` before committing any code
- Fix all linter warnings and errors
- The project uses a configured `.golangci.yml` file with specific rules
- Code that doesn't pass golangci-lint checks should not be committed

## Key Implementation Notes

### DICOM Tag Structure
- Tags are 32-bit values: (group << 16) | element
- Private tags have odd group numbers
- Tag format: (GGGG,EEEE) where G=group, E=element

### Transfer Syntax
- Controls encoding: explicit/implicit VR, byte order, compression
- Must be handled in both reading and writing
- Common syntaxes:
  - Explicit VR Little Endian (default)
  - Implicit VR Little Endian
  - Explicit VR Big Endian
  - Compressed (JPEG, JPEG-LS, JPEG2000, RLE)

### Value Representation (VR)
- Defines data type of DICOM elements
- String types: AE, AS, CS, DA, DS, DT, IS, LO, LT, PN, SH, ST, TM, UC, UI, UR, UT
- Binary types: OB, OD, OF, OL, OV, OW, UN
- Numeric types: FL, FD, SL, SS, UL, US
- Special types: AT (attribute tag), SQ (sequence)

### Reading DICOM Files
- Parse meta information (group 0002) first to determine transfer syntax
- Then parse dataset with appropriate transfer syntax
- Use observer pattern to build dataset during parsing
- Handle sequences recursively

### Private Tags
- Require private creator at element 0x0010-0x00FF
- Private creator reserves block of 256 elements
- Implementation should track private creator mappings

## Code Organization Suggestions

Consider this Go project structure:
```
go-dicom/
├── pkg/
│   ├── dicom/          # Core DICOM types
│   │   ├── tag.go
│   │   ├── vr.go
│   │   ├── dataset.go
│   │   ├── element.go
│   │   └── dictionary.go
│   ├── io/             # Reading/Writing
│   │   ├── buffer/
│   │   ├── reader/
│   │   └── writer/
│   ├── transfer/       # Transfer syntax
│   ├── network/        # DICOM networking (later)
│   └── imaging/        # Image handling (later)
├── internal/           # Internal helpers
├── cmd/                # Command-line tools
└── examples/           # Usage examples
```

## Testing Strategy

- Write tests alongside each component as you port from C#
- Use table-driven tests for VR validation, tag parsing, etc.
- Include test DICOM files (small samples) for integration tests
- Reference fo-dicom test suite for expected behavior
- Test edge cases: private tags, sequences, large files, various transfer syntaxes

## References

- DICOM Standard: https://www.dicomstandard.org/
- fo-dicom GitHub: https://github.com/fo-dicom/fo-dicom
- fo-dicom Documentation: https://github.com/fo-dicom/fo-dicom/wiki
