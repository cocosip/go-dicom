# DICOM Query/Retrieve SCU Example

This example demonstrates a complete DICOM Query/Retrieve Service Class User (SCU) client workflow.

## Features

Complete QR workflow:
1. **C-ECHO**: Verify connectivity to QR SCP
2. **C-FIND**: Query for studies, series, or images
3. **C-GET** or **C-MOVE**: Retrieve DICOM instances

For C-GET, the SCU acts as a temporary Storage SCP to receive C-STORE sub-operations from the QR SCP.

## Usage

### Basic Query (No Retrieval)

Query for all studies:

```bash
cd examples/qr_scu
go run main.go -level STUDY
```

### Query with Wildcard

Find all patients with last name "DOE":

```bash
go run main.go -patient-name "DOE^*"
```

### C-GET Retrieval

Query and retrieve using C-GET:

```bash
go run main.go -patient-name "DOE^*" -retrieve get -output-dir ./my_dicom_files
```

This will:
1. Start a temporary Storage SCP on port 11114
2. Send C-FIND query
3. Send C-GET request for the first matching study
4. Receive C-STORE sub-operations and save files to `./my_dicom_files/`

### C-MOVE Retrieval

Query and request C-MOVE to another AE:

```bash
go run main.go -patient-name "DOE^*" -retrieve move -move-destination STORESCP
```

The QR SCP will send instances to the destination AE (requires that AE to be running and accepting connections).

## Command-Line Flags

### Connection
- `-host string`: QR SCP hostname (default: "localhost")
- `-port int`: QR SCP port (default: 11113)
- `-calling-ae string`: Calling AE Title (default: "QRSCU")
- `-called-ae string`: Called AE Title (default: "QRSCP")
- `-timeout duration`: Overall timeout (default: 60s)

### Query Parameters
- `-level string`: Query level - PATIENT, STUDY, SERIES, or IMAGE (default: "STUDY")
- `-patient-name string`: PatientName match key (supports wildcards)
- `-patient-id string`: PatientID match key
- `-study-date string`: StudyDate (YYYYMMDD or YYYYMMDD-YYYYMMDD for range)
- `-study-uid string`: StudyInstanceUID match key

### Retrieve Options
- `-retrieve string`: Retrieve mode - `none`, `move`, or `get` (default: "none")
- `-retrieve-study-uid string`: Specific StudyUID to retrieve (defaults to first C-FIND result)
- `-move-destination string`: Destination AE for C-MOVE (default: "QRDEST")
- `-output-dir string`: Directory to save received files for C-GET (default: "./received_qr")
- `-scp-port int`: Local SCP port for receiving C-GET results (default: 11114)

### Other
- `-verbose`: Enable verbose logging

## Query Levels

### PATIENT Level
Returns unique patients matching the criteria.

```bash
go run main.go -level PATIENT -patient-name "SMITH^*"
```

### STUDY Level
Returns unique studies matching the criteria.

```bash
go run main.go -level STUDY -study-date 20260220-20260228
```

### SERIES Level
Returns unique series within matching studies.

```bash
go run main.go -level SERIES -study-uid 1.2.840.113619.2.55.3.2831164357.781.170000001.1
```

### IMAGE Level
Returns individual SOP instances.

```bash
go run main.go -level IMAGE -study-uid 1.2.840.113619.2.55.3.2831164357.781.170000001.1
```

## Complete Workflow Example

### Setup

Terminal 1 - Start QR SCP:
```bash
cd examples/qr_scp
go run main.go -verbose
```

Terminal 2 - Run full QR workflow:
```bash
cd examples/qr_scu

# Step 1: Query for all studies
go run main.go -level STUDY

# Step 2: Query specific patient
go run main.go -patient-name "DOE^JOHN"

# Step 3: Retrieve via C-GET
go run main.go -patient-name "DOE^JOHN" -retrieve get -verbose

# Check received files
ls -lh received_qr/
```

## Output Example

```
=== DICOM Query/Retrieve SCU - Complete Workflow ===
Target:          localhost:11113
Calling AE:      QRSCU
Called AE:       QRSCP
Query Level:     STUDY
Retrieve Mode:   get
Output Dir:      ./received_qr
Local SCP Port:  11114

[Step 1/4] Performing C-ECHO verification...
✓ C-ECHO succeeded

[Step 2/4] Performing C-FIND query (level=STUDY)...
✓ C-FIND completed: found 1 result(s)

C-FIND Results (STUDY level):
[1] Patient=DOE^JOHN (P001) StudyDate=20260220 Modality=CT
    StudyInstanceUID:  1.2.840.113619.2.55.3.2831164357.781.170000001.1

[Step 3/4] Starting local SCP on port 11114 to receive C-STORE...

[Step 4/4] Performing C-GET...
  Target Study: 1.2.840.113619.2.55.3.2831164357.781.170000001.1
  Output Dir:   ./received_qr
  Progress: remaining=2 completed=0 failed=0 warning=0
  Progress: remaining=0 completed=2 failed=0 warning=0
✓ C-GET completed

✓ Received 2 DICOM instance(s)

=== Query/Retrieve workflow completed successfully ===
```

## Architecture Notes

### C-GET Workflow

When using `-retrieve get`, the SCU:

1. Starts a temporary DICOM Storage SCP on `-scp-port`
2. Sends C-GET request to the QR SCP
3. QR SCP sends C-STORE sub-operations back to the SCU
4. SCU receives and saves each instance to `-output-dir`
5. SCU gives the temporary SCP five seconds to release active Associations, then closes any that remain

### C-MOVE Workflow

When using `-retrieve move`, the SCU:

1. Sends C-MOVE request to the QR SCP
2. QR SCP initiates C-STORE operations to the `-move-destination` AE
3. The move destination (a separate Storage SCP) receives the instances

**Note**: C-MOVE requires that the destination AE is running and configured at the QR SCP.

## DICOM Standard References

- **PS3.4 C.4**: Query/Retrieve Service Class
- **PS3.7 9.1.2**: C-FIND DIMSE Service
- **PS3.7 9.1.4**: C-MOVE DIMSE Service
- **PS3.7 9.1.3**: C-GET DIMSE Service
