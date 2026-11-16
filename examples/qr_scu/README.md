# DICOM Query/Retrieve SCU Example

This example demonstrates a complete DICOM Query/Retrieve Service Class User (SCU) client workflow.

## Features

Complete QR workflow:
1. **C-ECHO**: Verify connectivity to QR SCP
2. **C-FIND**: Query for studies, series, or images
3. **C-GET** or **C-MOVE**: Retrieve DICOM instances

For C-GET, the SCU registers a C-STORE handler and serves the Storage SCP role on
the same association, so no separate Storage SCP is required.

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
1. Send a C-FIND query
2. Send a C-GET request for the first matching study
3. Receive C-STORE sub-operations over the same association and save files to `./my_dicom_files/`

C-GET needs a C-STORE handler and Storage SOP Classes negotiated with the SCP
role; the example configures both automatically.

### C-MOVE Retrieval

Query and request C-MOVE to another AE:

```bash
go run main.go -patient-name "DOE^*" -retrieve move -move-destination QRDEST
```

The example starts an embedded Storage SCP on `-store-port` and advertises it as
the `-move-destination` AE. The QR SCP must map that AE to this host and port
(for `qr_scp`, pass `-ae-routes QRDEST=host:11114`).

## Command-Line Flags

### Connection
- `-host string`: QR SCP hostname (default: "localhost")
- `-port int`: QR SCP port (default: 11113)
- `-calling-ae string`: Calling AE Title (default: "QRSCU")
- `-called-ae string`: Called AE Title (default: "QRSCP")
- `-timeout duration`: Overall timeout (default: 180s)

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
- `-output-dir string`: Directory to save received DICOM files (default: "./received_qr")
- `-store-port int`: Local Storage SCP port advertised as the C-MOVE destination (default: 11114)

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

[Step 1/4] Performing C-ECHO verification...
✓ C-ECHO succeeded

[Step 2/4] Performing C-FIND query (level=STUDY)...
✓ C-FIND completed: found 1 result(s)

C-FIND Results (STUDY level):
[1] Patient=DOE^JOHN (P001) StudyDate=20260220 Modality=CT
    StudyInstanceUID:  1.2.840.113619.2.55.3.2831164357.781.170000001.1

[Step 3/4] Performing C-GET...
  Target Study: 1.2.840.113619.2.55.3.2831164357.781.170000001.1
  Output Dir:   ./received_qr
  (C-STORE sub-operations received on the same association)
  Progress: remaining=2 completed=0 failed=0 warning=0
  Progress: remaining=0 completed=2 failed=0 warning=0
✓ C-GET completed

✓ Received 2 DICOM instance(s)

=== Query/Retrieve workflow completed successfully ===
```

## Architecture Notes

### C-GET Workflow

When using `-retrieve get`, the SCU:

1. Negotiates the Query/Retrieve SOP Class and the Storage SOP Classes with the SCP role
2. Registers a C-STORE handler on the client
3. Sends the C-GET request to the QR SCP
4. Receives C-STORE sub-operations over the same association and saves each instance to `-output-dir`

### C-MOVE Workflow

When using `-retrieve move`, the SCU:

1. Starts an embedded Storage SCP on `-store-port`
2. Sends a C-MOVE request whose destination is the `-move-destination` AE
3. The QR SCP resolves that AE through its own routing table and opens a separate Association for the C-STORE sub-operations
4. The embedded Storage SCP receives the instances and saves them to `-output-dir`

**Note**: C-MOVE requires the QR SCP to know the destination AE. For `qr_scp`,
start it with `-ae-routes QRDEST=<this host>:11114`.

## DICOM Standard References

- **PS3.4 C.4**: Query/Retrieve Service Class
- **PS3.7 9.1.2**: C-FIND DIMSE Service
- **PS3.7 9.1.4**: C-MOVE DIMSE Service
- **PS3.7 9.1.3**: C-GET DIMSE Service
