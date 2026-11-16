# DICOM Query/Retrieve SCP Example

This example demonstrates a complete DICOM Query/Retrieve Service Class Provider (SCP) server.

## Features

- **C-ECHO**: Verification service
- **C-FIND**: Query for patients, studies, series, and images
- **C-MOVE**: Request transfer of instances to another AE (requires `-ae-routes`)
- **C-GET**: Request direct retrieval of instances over the same association

## Usage

### Basic Usage (Sample Data)

Start the QR SCP with built-in sample data:

```bash
cd examples/qr_scp
go run main.go
```

This starts a server on port 11113 with 2 sample studies (3 instances).

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
- `-ae-routes string`: C-MOVE routing table, `AE1=host:port,AE2=host:port`
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

Sample records have no source files, so they support C-FIND but not C-MOVE or
C-GET. Start the server with `-data-dir` to retrieve real instances.

## Testing with QR SCU

See the companion `qr_scu` example for a complete client that performs:
1. C-ECHO verification
2. C-FIND queries
3. C-GET or C-MOVE retrieval

Example test workflow:

```bash
# Terminal 1: Start SCP with files that retrieval can actually send
cd examples/qr_scp
go run main.go -data-dir /path/to/dicom/files -verbose

# Terminal 2: Run SCU query and C-GET
cd examples/qr_scu
go run main.go -patient-name "DOE^*" -retrieve get

# Or C-MOVE: the SCP must know the destination AE/host:port
go run main.go -patient-name "DOE^*" -retrieve move -move-destination QRDEST
# (start qr_scp with -ae-routes QRDEST=<this host>:11114)
```

## Implementation Notes

### C-MOVE
C-MOVE requires a routing table that maps AE titles to `host:port` addresses,
supplied with `-ae-routes` (for example `-ae-routes STORESCP=127.0.0.1:11112`).
The handler connects to the destination, sends each matching instance with
C-STORE, and reports progress through `op.SendPending`. An unknown destination
returns `0xA801`; a connection or sub-operation failure is reported through the
final warning/failure status.

### C-GET
C-GET sends each matching instance back over the requesting Association with
`op.SendCStore`, so no separate destination address is needed. Instances are
read from the paths captured while indexing `-data-dir`; records without a
source file (the built-in sample data) cannot be retrieved.

## DICOM Standard References

- **PS3.4 C.4.1**: Patient Root Query/Retrieve Information Model
- **PS3.4 C.6.1**: Study Root Query/Retrieve Information Model
- **PS3.7**: DIMSE Services (C-FIND, C-MOVE, C-GET)
