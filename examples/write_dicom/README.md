# DICOM File Writing Example

This example demonstrates how to create and write DICOM files using the go-dicom library.

## Features

- Create DICOM datasets from scratch
- Add patient, study, series, and image information
- Generate pixel data
- Write files with various options
- Customize File Meta Information

## Usage

### Basic Usage

```bash
go run main.go
```

This will create `output.dcm` with sample patient and image data.

## Advanced Options

### Custom Transfer Syntax

```go
// Implicit VR Little Endian
err := writer.WriteFile("output.dcm", ds,
    writer.WithTransferSyntax(transfer.ImplicitVRLittleEndian))

// Explicit VR Little Endian (default)
err := writer.WriteFile("output.dcm", ds,
    writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian))
```

### Custom Implementation Information

You can customize the Implementation Class UID and Version Name that appear in the File Meta Information:

```go
err := writer.WriteFile("output.dcm", ds,
    writer.WithImplementationClassUID("1.2.840.12345.6.7.8"),
    writer.WithImplementationVersionName("MYAPP_2.0"))
```

This is useful for:
- Identifying files created by your application
- Version tracking
- Compliance with your organization's UID scheme

### Write Without Preamble

For datasets that are not standalone files (e.g., network transmission):

```go
err := writer.Write(outputStream, ds,
    writer.WithoutPreamble())
```

### Custom File Meta Information

```go
fileMetaInfo := dataset.New()
fileMetaInfo.Add(element.NewString(tag.SourceApplicationEntityTitle, vr.AE,
    []string{"MY_WORKSTATION"}))

err := writer.WriteFile("output.dcm", ds,
    writer.WithFileMetaInfo(fileMetaInfo))
```

### Explicit Length Sequences

By default, sequences use undefined length (0xFFFFFFFF). You can use explicit lengths:

```go
err := writer.WriteFile("output.dcm", ds,
    writer.WithExplicitLengthSequences(true),
    writer.WithExplicitLengthSequenceItems())
```

### Complete Example with All Options

```go
package main

import (
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/transfer"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
    ds := dataset.New()

    // Add required elements
    ds.Add(element.NewString(tag.SOPClassUID, vr.UI,
        []string{"1.2.840.10008.5.1.4.1.1.2"})) // CT Image Storage
    ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI,
        []string{"1.2.840.113619.2.55.3.12345"}))

    // Add patient info
    ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
    ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

    // Write with custom options
    err := writer.WriteFile("custom_output.dcm", ds,
        writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian),
        writer.WithImplementationClassUID("1.2.840.12345.1"),
        writer.WithImplementationVersionName("MyApp_3.0.1"),
        writer.WithExplicitLengthSequences(true))

    if err != nil {
        log.Fatalf("Failed to write file: %v", err)
    }

    log.Println("DICOM file written successfully!")
}
```

## File Structure

The written DICOM file contains:

1. **Preamble (128 bytes)** - All zeros
2. **DICM Prefix (4 bytes)** - Magic string "DICM"
3. **File Meta Information (Group 0002)** - Always Explicit VR Little Endian
   - FileMetaInformationGroupLength (0002,0000)
   - FileMetaInformationVersion (0002,0001)
   - MediaStorageSOPClassUID (0002,0002)
   - MediaStorageSOPInstanceUID (0002,0003)
   - TransferSyntaxUID (0002,0010)
   - ImplementationClassUID (0002,0012)
   - ImplementationVersionName (0002,0013)
4. **Dataset** - Encoded according to Transfer Syntax

## Verifying Output

You can verify the written file using:

### Using go-dicom

```bash
cd ../read_dicom
# Update main.go to point to output.dcm
go run main.go
```

### Using dcmdump (DCMTK)

```bash
dcmdump output.dcm
```

### Using Python pydicom

```python
import pydicom
ds = pydicom.dcmread('output.dcm')
print(ds)
```

## Common Use Cases

### 1. Creating Test Files

```go
// Minimal test file
ds := dataset.New()
ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
writer.WriteFile("test.dcm", ds)
```

### 2. Converting Formats

```go
// Read from one format, write to another
result, _ := parser.ParseFile("input.dcm")
writer.WriteFile("output.dcm", result.Dataset,
    writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian))
```

### 3. Modifying Existing Files

```go
// Read, modify, and save
result, _ := parser.ParseFile("original.dcm")
result.Dataset.AddOrUpdate(element.NewString(tag.PatientName, vr.PN, []string{"Anonymous"}))
writer.WriteFile("modified.dcm", result.Dataset)
```

## Requirements

- Go 1.21 or later
- Valid DICOM dataset with required elements:
  - SOPClassUID (0008,0016)
  - SOPInstanceUID (0008,0018)

## Related Examples

- [read_dicom](../read_dicom) - Read and display DICOM files
- [anonymize](../anonymize) - Anonymize DICOM files
- [cstore_scu](../cstore_scu) - Send DICOM files over network

## Notes

- The library automatically generates File Meta Information if not provided
- SOPClassUID and SOPInstanceUID from the dataset are automatically copied to File Meta Information
- By default, sequences use undefined length (0xFFFFFFFF) which is more common
- Implementation Class UID should be unique to your organization/application
- Implementation Version Name should identify your application version

## Troubleshooting

### "Missing SOPClassUID"

Ensure your dataset contains:
```go
ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
```

### "Missing SOPInstanceUID"

Ensure your dataset contains:
```go
ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6"}))
```

### Invalid UID Format

UIDs must:
- Contain only digits and dots
- Not start or end with a dot
- Have components separated by dots
- Be unique (use a proper UID root for your organization)
