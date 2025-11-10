# DICOM Structured Report (SR) Package

## Overview

This package provides support for creating and parsing DICOM Structured Reports (SR).

## Current Status

🚧 **PARTIALLY IMPLEMENTED - REQUIRES DICOM DATASET API** 🚧

The core SR types and structures are implemented, but full functionality requires a complete `dicom.Dataset` API which is still under development.

## What are DICOM Structured Reports?

Structured Reports are used to encode clinical findings, measurements, and observations in a structured, machine-readable format. They represent information as a tree of content items with defined relationships.

### Key Components

**Content Items**: The building blocks of an SR
- Concept Name: Coded entry describing what the item represents
- Value Type: TEXT, NUM, CODE, IMAGE, CONTAINER, etc.
- Relationship Type: CONTAINS, HAS PROPERTIES, INFERRED FROM, etc.
- Value: The actual content

**Value Types**:
- `TEXT`: Text values
- `NUM`: Numeric measurements with units
- `CODE`: Coded concepts
- `CONTAINER`: Grouping of other content items
- `IMAGE`: Reference to an image SOP instance
- `COMPOSITE`: Reference to a composite SOP instance
- `PNAME`: Person name
- `DATE`, `TIME`, `DATETIME`: Temporal values
- `UIDREF`: UID reference

**Relationship Types**:
- `CONTAINS`: Parent-child containment
- `HAS PROPERTIES`: Describes properties
- `INFERRED FROM`: Indicates source
- `SELECTED FROM`: Selection from options
- `HAS OBS CONTEXT`: Observation context
- `HAS ACQ CONTEXT`: Acquisition context
- `HAS CONCEPT MOD`: Concept modifier

## Implemented Features

### ✅ Core Types
- [x] `CodeItem`: Coded entries (code value + scheme + meaning)
- [x] `MeasuredValue`: Numeric measurements with units
- [x] `ReferencedSOP`: References to SOP instances
- [x] Value types (TEXT, NUM, CODE, CONTAINER, etc.)
- [x] Relationship types (CONTAINS, HAS PROPERTIES, etc.)
- [x] Continuity types (SEPARATE, CONTINUOUS)

### ✅ Content Items
- [x] Text content items
- [x] Code content items
- [x] Numeric content items
- [x] Container content items
- [x] Content item hierarchy

### ✅ Structured Reports
- [x] SR creation with root code
- [x] Adding content items
- [x] Convenience methods (AddText, AddCode, AddNumeric, AddContainer)

### ⏳ Pending Implementation
- [ ] Full integration with `dicom.Dataset` API
- [ ] File I/O (Open, Save methods)
- [ ] Image/Composite/Waveform content items
- [ ] Spatial/Temporal coordinate items
- [ ] SR template support
- [ ] SR validation
- [ ] SR serialization/deserialization

## Dependencies Required

Before full implementation, the following need to be completed in the `dicom` package:

1. **Complete Dataset API**:
   ```go
   type Dataset struct {
       // Full implementation with element access
   }

   func (d *Dataset) GetString(tag Tag) (string, error)
   func (d *Dataset) PutString(tag Tag, value string) error
   func (d *Dataset) GetFloat64s(tag Tag) ([]float64, error)
   func (d *Dataset) FindElementByTag(tag Tag) (*Element, error)
   // ... etc
   ```

2. **Element and Sequence Support**:
   ```go
   type Element struct {
       Tag       Tag
       ValueType VR
       Value     Value
   }

   type SequencesValue struct {
       Values []*Dataset
   }
   ```

3. **File I/O**:
   ```go
   func ReadFile(filename string) (*File, error)
   func (f *File) Save(filename string) error
   ```

## Planned Usage (When Complete)

```go
// Create a new structured report
rootCode := sr.NewCodeItem("113704", "DCM", "SR Document")
report, _ := sr.NewStructuredReport(rootCode)

// Add a text finding
findingCode := sr.NewCodeItem("121071", "DCM", "Finding")
report.AddText(findingCode, sr.RelationshipContains, "No abnormalities detected")

// Add a numeric measurement
measureCode := sr.NewCodeItem("33728-7", "LN", "Size")
units := sr.NewCodeItem("mm", "UCUM", "millimeter")
value := sr.NewMeasuredValue(25.5, units)
report.AddNumeric(measureCode, sr.RelationshipContains, value)

// Add a container with nested items
containerCode := sr.NewCodeItem("121111", "DCM", "Summary")
item1, _ := sr.NewContentItemText(code1, sr.RelationshipContains, "Summary text")
item2, _ := sr.NewContentItemNumeric(code2, sr.RelationshipContains, measurement)
report.AddContainer(containerCode, sr.RelationshipContains, sr.ContinuitySeparate, item1, item2)

// Save to file (when implemented)
// report.Save("report.dcm")
```

## Implementation Roadmap

### Phase 1: Core Types ✅
- [x] Value types and relationships
- [x] CodeItem implementation
- [x] MeasuredValue implementation
- [x] ReferencedSOP implementation
- [x] Error handling

### Phase 2: Content Items ✅
- [x] ContentItem base implementation
- [x] Text, Code, Numeric types
- [x] Container type
- [x] Hierarchy support

### Phase 3: Structured Report ✅ (Partial)
- [x] SR root creation
- [x] Adding content items
- [x] Convenience methods
- [ ] Full Dataset integration

### Phase 4: File I/O (Pending)
- [ ] Open from file
- [ ] Save to file
- [ ] Stream support

### Phase 5: Advanced Features (Pending)
- [ ] SR templates (TID 1500, TID 1400, etc.)
- [ ] Image/Waveform references
- [ ] Spatial/Temporal coordinates
- [ ] SR validation against templates
- [ ] Comprehensive examples

## Testing

Basic unit tests are implemented for core types. Run tests with:

```bash
go test ./pkg/sr
```

Note: Some tests may fail until the full `dicom.Dataset` API is implemented.

## References

- DICOM Part 3: Information Object Definitions, Annex A.35 (SR IOD)
- DICOM Part 3, Section C.17.3 (SR Document Content Module)
- DICOM Part 16: Content Mapping Resource (SR Templates)
- https://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_A.35.html

## Contributing

This package is under active development. Contributions are welcome, especially for:
- Completing `dicom.Dataset` integration
- Implementing file I/O
- SR template support
- Validation logic
- Examples and documentation

## License

Microsoft Public License (MS-PL) - Same as parent project
