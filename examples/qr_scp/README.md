# DICOM Query/Retrieve SCP Example

This example demonstrates a complete DICOM Query/Retrieve Service Class Provider (SCP) server.

## Features

- **C-ECHO**: Verification service
- **C-FIND**: Query for patients, studies, series, and images
- **C-MOVE**: Request transfer of instances to another AE (simulated)
- **C-GET**: Request direct retrieval of instances (simulated)

## Usage

### Basic Usage (Sample Data)

Start the QR SCP with built-in sample data:

```bash
cd examples/qr_scp
go run main.go
```

This starts a server on port 11113 with 3 sample studies.

### Load Real DICOM Files

To index real DICOM files from a directory:

```bash
go run main.go -data-dir /path/to/dicom/files -verbose
```

The server will:
- Recursively scan the directory for DICOM files
- Extract metadata (PatientID, StudyUID, SeriesUID, SOPInstanceUID, etc.)
- Build an in-memory query index
- Track source file paths for future retrieval operations

### Command-Line Flags

- `-port int`: Port to listen on (default: 11113)
- `-data-dir string`: Directory containing DICOM files to index
- `-verbose`: Enable verbose logging

## Query Levels

The SCP supports all DICOM query levels:

- **PATIENT**: Query by PatientID, PatientName
- **STUDY**: Query by StudyInstanceUID, StudyDate, AccessionNumber
- **SERIES**: Query by SeriesInstanceUID, Modality
- **IMAGE**: Query by SOPInstanceUID, SOPClassUID

## Matching

The SCP supports:
- **Exact matching**: `PatientID=P001`
- **Wildcard matching**: `PatientName=DOE^*` or `PatientName=*JOHN*`
- **Date ranges**: `StudyDate=20260220-20260228`
- **Case-insensitive**: All string matching is case-insensitive

## Sample Data

When no `-data-dir` is specified, the server uses these sample records:

1. **Patient**: DOE^JOHN (P001)
   - **Study**: CT scan from 2026-02-20
   - **Modality**: CT
   - **Instances**: 2 images

2. **Patient**: SMITH^JANE (P002)
   - **Study**: MR scan from 2026-02-24
   - **Modality**: MR
   - **Instances**: 1 image

## Testing with QR SCU

See the companion `qr_scu` example for a complete client that performs:
1. C-ECHO verification
2. C-FIND queries
3. C-GET or C-MOVE retrieval

Example test workflow:

```bash
# Terminal 1: Start SCP
cd examples/qr_scp
go run main.go -verbose

# Terminal 2: Run SCU query
cd examples/qr_scu
go run main.go -patient-name "DOE^*" -retrieve get
```

## Implementation Notes

### C-MOVE
C-MOVE implementation requires:
- Configuration mapping AE titles to host:port addresses
- Client connection to the move destination AE
- C-STORE sub-operations to send instances

The current implementation simulates these operations and returns progress responses.

### C-GET
C-GET implementation requires:
- Access to the association/service to send C-STORE sub-operations
- Reading DICOM files from the indexed source paths
- Sending instances back to the requestor

The current implementation demonstrates the query matching and progress tracking,
but C-STORE sending requires architectural enhancements to the handler API.

## DICOM Standard References

- **PS3.4 C.4.1**: Patient Root Query/Retrieve Information Model
- **PS3.4 C.6.1**: Study Root Query/Retrieve Information Model
- **PS3.7**: DIMSE Services (C-FIND, C-MOVE, C-GET)
