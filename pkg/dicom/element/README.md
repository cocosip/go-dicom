# DICOM Element Package

The `pkg/dicom/element` package implements the various element types for DICOM data elements.

> Note: when application code reads and updates common values in a Dataset, prefer
> `pkg/dicom/dataset`'s `GetString`, `GetBytes`, `GetUInt16s`,
> `AddValue`, and `AddOrUpdateValue`. Use this package's concrete types and
> constructors directly only when you need to build a specific Element, handle
> private tags, or explicitly specify the VR.

## Overview

An Element is the basic building block of a DICOM dataset. Each Element contains:
- **Tag**: unique identifier (group, element)
- **VR (Value Representation)**: value representation type
- **Data**: binary data, stored in a ByteBuffer

## Implemented Element types

### 1. String types (String)

**File**: `string.go`

Supports all string VRs:
- AE (Application Entity)
- AS (Age String)
- CS (Code String)
- DA (Date)
- DS (Decimal String)
- DT (Date Time)
- IS (Integer String)
- LO (Long String)
- LT (Long Text)
- PN (Person Name)
- SH (Short String)
- ST (Short Text)
- TM (Time)
- UC (Unlimited Characters)
- UI (Unique Identifier)
- UR (URI/URL)
- UT (Unlimited Text)

**Features**:
- Supports multiple values (separated by backslash)
- Supports character set encoding
- Automatically strips trailing spaces

**Example**:
```go
// single-value string
elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})
name := elem.GetString() // "Doe^John"

// multi-value string
elem := element.NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY", "AXIAL"})
values := elem.GetValues() // ["ORIGINAL", "PRIMARY", "AXIAL"]
```

### 2. Special string types

#### 2.1 Person Name type (PersonName)

**File**: `person_name.go`

A special implementation of PN (Person Name) that supports parsing person name components:
- FamilyName
- GivenName
- MiddleName
- NamePrefix (prefix, e.g. Dr., Prof.)
- NameSuffix (suffix, e.g. Jr., III)

**Example**:
```go
elem := element.NewPersonName(tag.PatientName, []string{"Doe^John^A^Dr.^Jr."})
components, _ := elem.GetComponents(0)
// components.FamilyName = "Doe"
// components.GivenName = "John"
// components.MiddleName = "A"
// components.NamePrefix = "Dr."
// components.NameSuffix = "Jr."

formatted, _ := elem.GetFormattedName(0) // "Dr. John A Doe, Jr."
```

#### 2.2 Date and time types

**File**: `date.go`

Contains three date and time types:

**Date (DA)** - date format YYYYMMDD
```go
elem := element.NewDate(tag.StudyDate, []string{"20250106"})
date, _ := elem.GetDate(0) // time.Time

// create from time.Time
testTime := time.Date(2025, 1, 6, 0, 0, 0, 0, time.UTC)
elem := element.NewDateFromTime(tag.StudyDate, []time.Time{testTime})
```

**Time (TM)** - time format HHMMSS.FFFFFF
```go
elem := element.NewTime(tag.StudyTime, []string{"123456.789"})
tm, _ := elem.GetTime(0) // time.Time

// supports multiple formats: HH, HHMM, HHMMSS, HHMMSS.FFF, HHMMSS.FFFFFF
elem := element.NewTime(tag.StudyTime, []string{"1234"}) // HHMM format
```

**DateTime (DT)** - date-time format YYYYMMDDHHMMSS.FFFFFF&ZZXX
```go
elem := element.NewDateTime(tag.AcquisitionDateTime, []string{"20250106123456"})
dt, _ := elem.GetDateTime(0) // time.Time

// supports time zones
elem := element.NewDateTimeFromTime(tag.AcquisitionDateTime, times)
```

#### 2.3 Numeric string types

**File**: `numeric_string.go`

**DecimalString (DS)** - decimal string to floating point
```go
elem := element.NewDecimalString(tag.SliceThickness, []string{"1.5", "2.0"})
f, _ := elem.GetFloat(0) // 1.5

// create from float
elem := element.NewDecimalStringFromFloat(tag.SliceThickness, []float64{1.5, 2.0})
floats, _ := elem.GetFloats() // []float64{1.5, 2.0}

// supports float32
f32, _ := elem.GetFloat32(0)
```

**IntegerString (IS)** - integer string to integer
```go
elem := element.NewIntegerString(tag.NumberOfFrames, []string{"100"})
i, _ := elem.GetInt(0) // 100

// create from int
elem := element.NewIntegerStringFromInt(tag.NumberOfFrames, []int{100})
ints, _ := elem.GetInts() // []int{100}

// supports int32, int64
i32, _ := elem.GetInt32(0)
i64, _ := elem.GetInt64(0)
```

### 3. Numeric types (Numeric)

**File**: `numeric.go`

| Type | VR | Go type | Size | Range |
|------|----|----|------|------|
| UnsignedShort | US | uint16 | 2 bytes | 0 ~ 65,535 |
| UnsignedLong | UL | uint32 | 4 bytes | 0 ~ 4,294,967,295 |
| SignedShort | SS | int16 | 2 bytes | -32,768 ~ 32,767 |
| SignedLong | SL | int32 | 4 bytes | -2,147,483,648 ~ 2,147,483,647 |
| Float | FL | float32 | 4 bytes | IEEE 754 single precision |
| Double | FD | float64 | 8 bytes | IEEE 754 double precision |
| SignedVeryLong | SV | int64 | 8 bytes | -2^63 ~ 2^63-1 |
| UnsignedVeryLong | UV | uint64 | 8 bytes | 0 ~ 2^64-1 |

**Example**:
```go
// create a US element
elem := element.NewUnsignedShort(tag.Rows, []uint16{512, 512})
rows, _ := elem.GetValue(0) // 512

// create an FD element
elem := element.NewDouble(tag.RealWorldValueSlope, []float64{1.234567890123})
slope, _ := elem.GetValue(0) // 1.234567890123
```

### 4. Attribute Tag type (AttributeTag)

**File**: `attribute_tag.go`

**VR**: AT

Stores DICOM tag references; each value is a 32-bit tag (group + element).

**Example**:
```go
tags := []*tag.Tag{tag.PatientName, tag.StudyDate}
elem := element.NewAttributeTag(tag.OffendingElement, tags)
tag, _ := elem.GetValue(0) // tag.PatientName
```

### 5. Binary types (Binary)

**File**: `binary.go`

| Type | VR | Description |
|------|----|----|
| OtherByte | OB | byte sequence |
| OtherWord | OW | 16-bit word sequence |
| Unknown | UN | unknown data |
| OtherDouble | OD | 64-bit floating point sequence |
| OtherFloat | OF | 32-bit floating point sequence |
| OtherLong | OL | 32-bit integer sequence |
| OtherVeryLong | OV | 64-bit integer sequence |

**Features**:
- Used to store pixel data, waveform data, etc.
- Count() returns 0 for empty data; non-empty data is returned as a single value, i.e. 1
- GetData() returns the raw byte array

**Example**:
```go
// construct a binary Element directly (for Pixel Data in ordinary applications,
// prefer imaging/pixeldata.FromDataset or pixeldata.NewForDataset)
data := []byte{0x00, 0xFF, 0x80, 0x7F}
elem := element.NewOtherWord(tag.PixelData, data)
pixelData := elem.GetData() // [0x00, 0xFF, 0x80, 0x7F]
```

## Interface design

### Element interface

All Element types implement this interface:

```go
type Element interface {
    Tag() *tag.Tag
    ValueRepresentation() *vr.VR
    Buffer() buffer.ByteBuffer
    Length() uint32
    Count() int
    String() string
    Validate() error
}
```

### Inheritance structure

```
Element (interface)
  └─ base (struct) - common base for all types
      ├─ String
      │   └─ PersonName
      ├─ UnsignedShort
      ├─ UnsignedLong
      ├─ SignedShort
      ├─ SignedLong
      ├─ Float
      ├─ Double
      ├─ SignedVeryLong
      ├─ UnsignedVeryLong
      ├─ AttributeTag
      ├─ OtherByte
      ├─ OtherWord
      ├─ OtherDouble
      ├─ OtherFloat
      ├─ OtherLong
      ├─ OtherVeryLong
      └─ Unknown
```

## File list

| File | Description | Lines |
|------|------|------|
| `doc.go` | package documentation | ~50 |
| `element.go` | Element interface and base implementation | ~95 |
| `string.go` | generic string type | ~165 |
| `person_name.go` | person name type (PN) | ~135 |
| `date.go` | date and time types (DA/TM/DT) | ~230 |
| `numeric_string.go` | numeric string types (DS/IS) | ~235 |
| `numeric.go` | numeric types (US/UL/SS/SL/FL/FD/SV/UV) | ~605 |
| `attribute_tag.go` | attribute tag type (AT) | ~95 |
| `binary.go` | binary types (OB/OW/OD/OF/OL/OV/UN) | ~275 |
| `element_test.go` | interface tests | ~55 |
| `string_test.go` | string tests | ~110 |
| `date_test.go` | date and time tests | ~125 |
| `numeric_string_test.go` | numeric string tests | ~150 |
| `numeric_test.go` | numeric tests | ~370 |
| `README.md` | documentation | ~280 |
| **Total** |  | **~2,740 lines** |

## Test coverage

- **Test files**: 5
- **Test functions**: 21
- **Test cases**: 70+ subtests
- **Test coverage**: full coverage of main functionality
- **Test pass rate**: 100% ✅

```bash
# run all tests
cd pkg/dicom/element
go test -v

# run a specific test
go test -v -run TestString
go test -v -run TestNumeric
```

## Usage examples

### Creating an Element

```go
// string
elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})

// numeric
elem := element.NewUnsignedShort(tag.Rows, []uint16{512})

// binary
data := []byte{0x01, 0x02, 0x03, 0x04}
elem := element.NewOtherByte(tag.PixelData, data)
```

### Creating from a Buffer

```go
buf := buffer.NewMemory([]byte("Patient Name"))
elem := element.NewStringFromBuffer(tag.PatientName, vr.PN, buf, nil)
```

### Accessing values

```go
// string
name := elem.GetString()
values := elem.GetValues()

// numeric
value, err := elem.GetValue(0)
values, err := elem.GetValues()

// binary
data := elem.GetData()
```

## Cross-package implementations

- [x] Sequence (SQ) - located in `pkg/dicom/dataset` and implements `element.Element`

## Future work

### Features to enhance

- [ ] VR-specific validation rules
- [ ] more complete error handling
- [ ] performance optimization
- [ ] more test cases

## Related packages

- `pkg/dicom/tag` - DICOM tags
- `pkg/dicom/vr` - value representations
- `pkg/io/buffer` - buffer interface
- `pkg/dicom/charset` - character set encoding

## References

- DICOM standard PS3.5: Data Structures and Encoding
- fo-dicom library: https://github.com/fo-dicom/fo-dicom
