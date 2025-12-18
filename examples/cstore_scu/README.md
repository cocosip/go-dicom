# DICOM C-STORE SCU Example

This example demonstrates how to send DICOM files to a DICOM server using the C-STORE operation (Service Class User - SCU).

## Features

- Connect to DICOM SCP (server)
- Verify connection with C-ECHO
- Send single DICOM file
- Send all DICOM files from a directory
- Support for multiple SOP Classes and Transfer Syntaxes
- Display file metadata before sending
- Track transfer statistics

## Usage

### Basic Usage - Send Single File

```bash
go run main.go -file path/to/file.dcm
```

### Send All Files from Directory

```bash
go run main.go -dir path/to/dicom/directory
```

### Specify Server Details

```bash
go run main.go \
  -host 192.168.1.100 \
  -port 104 \
  -calling-ae MY_SCU \
  -called-ae PACS_SERVER \
  -file path/to/file.dcm
```

### Verify Connection Only (C-ECHO)

```bash
go run main.go -verify
```

### Display Metadata

```bash
go run main.go -file path/to/file.dcm -metadata
```

## Command Line Options

| Option | Default | Description |
|--------|---------|-------------|
| `-host` | localhost | DICOM server hostname or IP address |
| `-port` | 11112 | DICOM server port |
| `-calling-ae` | STORE_SCU | Calling AE Title (your client) |
| `-called-ae` | STORE_SCP | Called AE Title (remote server) |
| `-file` | | Single DICOM file to send |
| `-dir` | | Directory containing DICOM files |
| `-timeout` | 30s | Operation timeout |
| `-verify` | false | Only verify connection (C-ECHO) |
| `-metadata` | false | Print file metadata before sending |

## Examples

### 1. Test Connection to Server

```bash
go run main.go -host pacs.hospital.local -port 104 -verify
```

### 2. Send Study to PACS

```bash
go run main.go \
  -host pacs.hospital.local \
  -port 104 \
  -calling-ae MODALITY_1 \
  -called-ae PACS \
  -dir /studies/CT_20240101
```

### 3. Send with Metadata Display

```bash
go run main.go \
  -file CT.1.2.3.4.dcm \
  -metadata
```

## Supported SOP Classes

The example is pre-configured to support:

- Verification SOP Class (C-ECHO)
- CT Image Storage
- MR Image Storage
- Secondary Capture Image Storage

Additional SOP Classes can be added by modifying the `main.go` file and adding more presentation contexts.

## Output Example

```
=== DICOM C-STORE SCU Example ===
Calling AE: STORE_SCU
Called AE:  STORE_SCP
Server:     localhost:11112

Connecting to localhost:11112...
Connected successfully!
Association established with 4 presentation contexts

Performing C-ECHO verification...
C-ECHO successful!

Sending 1 file(s)...

[1/1] Processing: image.dcm
  SOP Class:    1.2.840.10008.5.1.4.1.1.2
  SOP Instance: 1.2.840.113619.2.55.3.12345
  Sent in:      125ms
  SUCCESS

=== Summary ===
Total files: 1
Successful:  1
Failed:      0
```

## Requirements

- Go 1.21 or later
- Access to a DICOM SCP server (or use the companion `cstore_scp` example)

## Related Examples

- [cstore_scp](../cstore_scp) - DICOM C-STORE receiver (server)
- [read_dicom](../read_dicom) - Read and display DICOM file contents

## Error Handling

The example handles common errors:

- **Connection refused**: Check if the server is running and the port is correct
- **Association rejected**: Verify AE titles are configured correctly on the server
- **No presentation context**: The server doesn't support the SOP Class in the file
- **Parse error**: The file is not a valid DICOM file
- **Missing UIDs**: The file is missing required SOPClassUID or SOPInstanceUID

## Notes

- Default port 11112 is used instead of standard DICOM port 104 (which requires root/admin privileges)
- The example accepts the first transfer syntax proposed by the server for each SOP Class
- Files are validated during parsing, and only valid DICOM files will be sent
