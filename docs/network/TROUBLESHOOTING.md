# DICOM Network Troubleshooting Guide

This document lists common problems and solutions when using go-dicom network features.

## Table of Contents

- [Connection Issues](#connection-issues)
- [Association Issues](#association-issues)
- [DIMSE Operation Issues](#dimse-operation-issues)
- [Performance Issues](#performance-issues)
- [Debugging Tips](#debugging-tips)

---

## Connection Issues

### Problem: Connection refused

**Error message**:
```
failed to connect to 192.168.1.100:104: dial tcp 192.168.1.100:104: connect: connection refused
```

**Possible causes**:
1. The server is not running
2. Wrong port number
3. A firewall is blocking the connection
4. The server is listening on the wrong address

**Solution**:

```bash
# 1. Check whether the server is running
netstat -an | grep 104
# or
ss -tlnp | grep 104

# 2. Check whether the port is reachable
telnet 192.168.1.100 104
# or
nc -vz 192.168.1.100 104

# 3. Check the firewall
# Linux
sudo iptables -L -n | grep 104
# Windows
netsh advfirewall firewall show rule name=all | findstr 104

# 4. Verify the server configuration
# Make sure the server listens on the correct address
srv := server.New(
    server.WithPort(104),  // Make sure the port is correct
)
```

### Problem: Connection timeout

**Error message**:
```
failed to connect: context deadline exceeded
```

**Possible causes**:
1. Network latency is too high
2. The server responds slowly
3. The network is unreachable

**Solution**:

```go
// Increase the connection timeout
c := client.New(
    client.WithConnectTimeout(30*time.Second),  // Increase to 30 seconds
)

// Or control it with a context
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
err := c.Connect(ctx, host, port)
```

### Problem: TLS handshake failure

**Error message**:
```
TLS handshake failed: x509: certificate signed by unknown authority
```

**Solution**:

```go
import "crypto/tls"

// Development: skip certificate verification (not recommended for production)
tlsConfig := &tls.Config{
    InsecureSkipVerify: true,
}

// Production: add a custom CA certificate
caCert, err := os.ReadFile("ca.crt")
if err != nil {
    log.Fatal(err)
}

caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)

tlsConfig := &tls.Config{
    RootCAs: caCertPool,
    ServerName: "pacs.hospital.com",
}

conn, err := transport.DialTLS(ctx, "tcp", "pacs.hospital.com:11112",
    transport.WithTLSConfig(tlsConfig))
```

---

## Association Issues

### Problem: Association rejected - AE Title mismatch

**Error message**:
```
association rejected: called AE title not recognized
```

**Possible causes**:
1. The Called AE Title is incorrect
2. The server has strict AE Title checking enabled
3. AE Titles are case-sensitive

**Solution**:

```go
// 1. Confirm the server's AE Title
c := client.New(
    client.WithCalledAE("EXACT_SERVER_AE"),  // must match exactly
)

// 2. Check the server configuration
srv := server.New(
    server.WithAETitle("MY_SCP"),
    server.WithStrictAECheck(false),  // Disable strict checking (development environment)
)

// 3. Or configure the allowed Calling AE Titles
srv := server.New(
    server.WithAcceptedCallingAETitles([]string{
        "SCU_1",
        "SCU_2",
        "WORKSTATION_*",  // Wildcards are supported
    }),
)
```

### Problem: No Presentation Context accepted

**Error message**:
```
invalid A-ASSOCIATE-AC: all presentation contexts were rejected
```

**Possible causes**:
1. No Presentation Context was added
2. The SOP Class UID does not match
3. The Transfer Syntax is not supported

**Solution**:

```go
// 1. Add Presentation Contexts
c := client.New(/* ... */)

// Add Verification
c.AddPresentationContext(
    "1.2.840.10008.1.1",      // Verification SOP Class
    "1.2.840.10008.1.2",      // Implicit VR Little Endian
    "1.2.840.10008.1.2.1",    // Explicit VR Little Endian
)

// Add Storage
c.AddPresentationContext(
    "1.2.840.10008.5.1.4.1.1.2",  // CT Image Storage
    "1.2.840.10008.1.2.1",        // Explicit VR Little Endian
)

// 2. Server side: make sure these SOP Classes are accepted
srv.SetAssociationNegotiatorFunc(
    func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
        for _, pc := range assoc.PresentationContexts {
            // Accept all SOP Classes (development environment)
            if len(pc.ProposedTransferSyntaxes) > 0 {
                pc.Accept(pc.ProposedTransferSyntaxes[0])
            }
        }
        return responder.SendAccept(ctx, assoc)
    })
```

### Problem: MaxPDULength mismatch

**Error message**:
```
PDU length 65536 exceeds maximum allowed 16384
```

**Cause**: Each DICOM AE advertises the maximum PDU length **it can receive**
in the Maximum Length Sub-item, independently for each direction of data flow
(PS3.7 D.3.3.1, PS3.8 D.1.1). The limits do not have to match, and a mismatch is
legal: a sender must simply never exceed the receiver's advertised value. The
error above means the peer sent a P-DATA-TF PDU larger than this AE advertised.

**Solution**:

```go
// The local value is the receive limit; it is advertised to the peer.
// A peer with a different limit is still compliant.

// Client: advertise that this SCU can receive PDUs up to 32 KB (0 = unlimited)
c := client.New(
    client.WithMaxPDULength(32768),
)

// Server: advertise that this SCP can receive PDUs up to 64 KB
srv := server.New(
    server.WithMaxPDULength(65536),
)
```

Outgoing PDUs are fragmented to the smaller of the local configured limit and
the peer's advertised limit, so raising only one side is always safe.

---

## DIMSE Operation Issues

### Problem: C-STORE failure - missing SOP Class UID

**Error message**:
```
failed to create C-STORE request: SOPClassUID not found in dataset
```

**Solution**:

```go
import (
    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// Make sure the Dataset contains the required elements
ds := dataset.New()
ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6"}))
// ... other required elements

// Then send
err := c.CStore(ctx, ds)
```

### Problem: C-FIND returns empty results

**Possible causes**:
1. The query conditions do not match
2. The Query/Retrieve Level is wrong
3. Required fields are missing

**Solution**:

```go
// 1. Check the query conditions
query := dataset.New()
query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))  // exact match
query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE*"}))      // wildcard match
query.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20240101-"}))   // date range

// 2. Add return fields (an empty string means return that field)
query.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{""}))
query.Add(element.NewString(tag.StudyDescription, vr.LO, []string{""}))
query.Add(element.NewString(tag.StudyDate, vr.DA, []string{""}))

// 3. Make sure the Level is correct
results, err := c.CFind(ctx, dimse.QueryRetrieveLevelStudy, query)

// 4. Debug: print the query conditions
for _, elem := range query.Elements() {
    log.Printf("Query: %s = %v\n", elem.Tag(), elem.Value())
}
```

### Problem: C-MOVE/C-GET failure - destination unreachable

**Error message**:
```
C-MOVE failed: destination AE not found
```

**Solution**:

```go
// 1. Make sure the destination AE Title is correct
err := c.CMove(ctx, level, "EXACT_DEST_AE", identifier, callback)

// 2. Server side: verify the destination exists
// You can maintain an AE Title -> Host:Port mapping
var knownAEs = map[string]string{
    "DEST_SCP_1": "192.168.1.101:104",
    "DEST_SCP_2": "192.168.1.102:104",
}

func getDestinationHost(ae string) (string, error) {
    host, ok := knownAEs[ae]
    if !ok {
        return "", fmt.Errorf("destination AE not found: %s", ae)
    }
    return host, nil
}
```

### Problem: Request timeout

**Error message**:
```
C-FIND failed: context deadline exceeded
```

**Solution**:

```go
// 1. Increase the request timeout
c := client.New(
    client.WithRequestTimeout(60*time.Second),  // Increase to 60 seconds
)

// 2. Use a custom timeout for a specific operation
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

results, err := c.CFind(ctx, level, query)

// 3. For long-running operations, use streaming
err := c.CFindWithCallback(ctx, level, query,
    func(result *dataset.Dataset) bool {
        // Process each result immediately instead of waiting for all to finish
        processResult(result)
        return true
    })
```

---

## Performance Issues

### Problem: Large file transfers are slow

**Symptom**:
Transferring large DICOM files (such as CT series) is very slow

**Solution**:

```go
// 1. Increase MaxPDULength
c := client.New(
    client.WithMaxPDULength(65536),  // 64 KB (default 16 KB)
)

srv := server.New(
    server.WithMaxPDULength(65536),
)

// 2. Make sure network bandwidth is sufficient
// Use iperf to test network speed
// iperf -s  # server side
// iperf -c <server-ip>  # client

// 3. Reuse the connection when transferring in batches
c := client.New(/* ... */)
c.Connect(ctx, host, port)
defer c.Close()

for _, file := range files {
    ds, _ := parser.ParseFile(file)
    c.CStore(ctx, ds)  // reuse the same connection
}
```

### Problem: Too many concurrent connections cause resource exhaustion

**Symptom**:
High server memory usage, or "too many open files" errors

**Solution**:

```go
// 1. Limit the server's concurrent connections
srv := server.New(
    server.WithMaxConnections(50),  // at most 50 concurrent connections
)

// 2. Set PDU read/idle timeouts for established connections
srv := server.New(
    server.WithTransportReadTimeout(30*time.Second), // Disconnect when a complete PDU is not received within 30 seconds
)

// 3. Increase the file descriptor limit (Linux)
// ulimit -n 65536

// 4. Use a connection pool on the client (if frequent connections are needed)
// Note: DICOM connections are usually long-lived and do not need a pool
```

### Problem: Excessive memory usage

**Symptom**:
High memory usage when processing a large number of query results

**Solution**:

```go
// ❌ Bad practice: collect all results into memory
results, err := c.CFind(ctx, level, query)  // If there are 10000 results this uses a lot of memory

// ✅ Good practice: use streaming
err := c.CFindWithCallback(ctx, level, query,
    func(result *dataset.Dataset) bool {
        // Process each result immediately
        processAndSaveResult(result)

        // Once a result is processed it can be garbage collected
        return true
    })
```

---

## Debugging Tips

### Structured network logs

Network logs are output uniformly via `pkg/logging.Configure`. When querying QR
operations, prefer the following fields over looking only at the raw `command`
value:

- `operation`: `C-FIND`, `C-MOVE`, `C-GET` or `C-STORE`
- `command_name`: e.g. `C-GET-RQ`, `C-GET-RSP`
- `message_id`: DICOM request/response correlation identifier
- `operation_id`, `parent_operation_id`: in-process auxiliary correlation identifiers, applicable only to confirmable parent-child relationships
- `query_retrieve_level`: `PATIENT`, `STUDY`, `SERIES`, `IMAGE`
- `sop_class_uid`, `transfer_syntax`, `presentation_context_id`: negotiation and encoding context
- `status`, `status_hex`, `status_state`: DICOM status code and human-readable state
- `remaining_suboperations`, `completed_suboperations`,
  `failed_suboperations`, `warning_suboperations`: C-MOVE/C-GET progress

Association logs also include `local_addr`, `remote_addr`, Calling/Called AE,
implementation UID/version, maximum PDU, async operation window, and the number
of accepted Presentation Contexts. `presentation_contexts` is a structured value:
the text handler outputs
`1:1.2.840.10008.1.1 -> 1.2.840.10008.1.2 [acceptance]; 3:...`, while the JSON
handler outputs a native array (without double-escaped `\"`), containing each
context's ID, Abstract Syntax, Transfer Syntax, and negotiation result, which is
used to judge whether negotiation is as expected. At INFO level it outputs
Association and all DIMSE message summaries (including pending progress);
timeouts, cancellations, and transport details are output at WARN/DEBUG. A C-MOVE
destination AE usually uses another Association; the source-side C-MOVE and the
destination-side C-STORE cannot be automatically correlated merely from a local
`operation_id` — the business layer needs to provide an external correlation ID.

Logs do not output full datasets, PDU contents, or patient identifiers. To
troubleshoot data issues you must format explicitly:

```go
// Single line, suitable as a log field; note that it includes patient information, so use it only in controlled scenarios.
slog.String("dataset", ds.Format(dataset.FormatOptions{MaxElements: 100}))

// Multi-line aligned output, suitable for local debugging or a CLI.
fmt.Println(ds.Format(dataset.FormatOptions{
    Layout:   dataset.LayoutDump,
    MaxDepth: -1,
}))
```

### Enabling verbose logging

```go
import "log"

// Set the log level
log.SetFlags(log.LstdFlags | log.Lshortfile)

// Add logs at key points
log.Printf("Connecting to %s:%d\n", host, port)
log.Printf("Sending C-ECHO request\n")
log.Printf("Received response: status=0x%04X\n", resp.StatusCode())
```

### Capturing packets with Wireshark

Use Wireshark with the DICOM dissector to capture traffic on the DICOM port (typically 104)
and inspect the association and DIMSE exchanges.

### Checking Association details

```go
// Check the Association after connecting
assoc := c.GetAssociation()
if assoc != nil {
    log.Printf("Association established:\n")
    log.Printf("  Called AE:  %s\n", assoc.CalledAE)
    log.Printf("  Calling AE: %s\n", assoc.CallingAE)
    log.Printf("  Max PDU:    %d\n", assoc.MaxPDULength)

    log.Printf("Accepted Presentation Contexts:\n")
    for _, pc := range assoc.GetAcceptedPresentationContexts() {
        transferSyntax := ""
        if pc.AcceptedTransferSyntax != nil {
            transferSyntax = pc.AcceptedTransferSyntax.UID().UID()
        }
        log.Printf("  ID=%d: %s (%s)\n",
            pc.ID, pc.AbstractSyntax, transferSyntax)
    }
}
```

### Printing DIMSE messages

```go
// Print request information
log.Printf("C-STORE Request:\n")
log.Printf("  SOP Class UID:    %s\n", req.AffectedSOPClassUID())
log.Printf("  SOP Instance UID: %s\n", req.AffectedSOPInstanceUID())
log.Printf("  Message ID:       %d\n", req.MessageID())

// Print Dataset contents
ds := req.DataDataset()
for _, elem := range ds.Elements() {
    log.Printf("  %s: %v\n", elem.Tag(), elem.Value())
}
```

### Testing network connectivity

```go
// Simple Echo test program
package main

import (
    "context"
    "log"
    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    c := client.New(
        client.WithCallingAE("TEST_SCU"),
        client.WithCalledAE("PACS_SCP"),
    )

    c.AddPresentationContext(
        "1.2.840.10008.1.1",   // Verification
        "1.2.840.10008.1.2.1",
    )

    ctx := context.Background()
    log.Println("Connecting...")
    if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
        log.Fatalf("Connect failed: %v", err)
    }
    defer c.Close()

    log.Println("Sending C-ECHO...")
    if err := c.CEcho(ctx); err != nil {
        log.Fatalf("C-ECHO failed: %v", err)
    }

    log.Println("Success!")
}
```

### Common DICOM tools

For comparison testing:

- **dcmtk**: DICOM Toolkit (command-line tool)
  ```bash
  # C-ECHO test
  echoscu -aet MY_SCU -aec PACS_SCP 192.168.1.100 104

  # C-STORE
  storescu -aet MY_SCU -aec PACS_SCP 192.168.1.100 104 image.dcm

  # C-FIND
  findscu -aet MY_SCU -aec PACS_SCP 192.168.1.100 104 query.dcm
  ```

- **Orthanc**: lightweight PACS server
  ```bash
  # Run with Docker
  docker run -p 4242:4242 -p 8042:8042 jodogne/orthanc
  ```

- **DCMTK's storescp**: a simple Storage SCP
  ```bash
  # Start the Storage SCP
  storescp -aet STORAGE_SCP 11112
  ```

---

## Common Error Codes

| Status Code | Meaning | Possible Cause | Solution |
|--------|------|----------|----------|
| 0x0000 | Success | Success | - |
| 0xFF00 | Pending | Waiting for more results | Continue receiving |
| 0x0107 | Attribute list error | Query field error | Check the query conditions |
| 0x0116 | Attribute value out of range | Field value out of range | Check the value range |
| 0xA700 | Out of resources | Insufficient server resources | Reduce concurrency or add server resources |
| 0xA900 | Dataset does not match SOP Class | Dataset does not match the SOP Class | Check required fields |
| 0xC000 | Cannot understand | Cannot understand the request | Check the request format |
| 0xC001 | Unrecognized operation | Unsupported operation | Check the SOP Class |

---

## Getting Help

If the problem is still unresolved:

1. **Check the logs**: enable verbose logging and look at the error messages
2. **Packet capture analysis**: use Wireshark to analyze DICOM traffic
3. **Comparison testing**: use tools such as dcmtk to test the server
4. **Read the source**: read the go-dicom source to understand implementation details
5. **Submit an Issue**: submit an Issue on GitHub with the logs and packet captures

---

## Next Steps

- See the [networking guide](./README.md) for detailed API documentation
