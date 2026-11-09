# DIMSE Package

DIMSE (DICOM Message Service Element) is the core message layer of DICOM network communication.

## Features

- ✅ C-ECHO: verify connectivity
- ✅ C-STORE: store DICOM instances
- ✅ C-FIND: query the DICOM archive
- ✅ C-MOVE: retrieve instances to a third-party destination
- ✅ C-GET: retrieve instances over the same association
- ✅ C-CANCEL: cancel a pending C-FIND/C-MOVE/C-GET operation
- ✅ N-EVENT-REPORT, N-GET, N-SET, N-ACTION, N-CREATE, N-DELETE
- ✅ Automatic MessageID management
- ✅ Priority support
- ✅ Status code handling
- ✅ Outbound required-field validation (`ValidateRequestForSend`)
- ✅ Optional response parameters (Type U(=)) are omitted when empty and tolerated when absent
- ✅ Missing Status element is classified as Failure instead of an all-false status

## Automatic MessageID Management

### Problem
In DICOM network communication, every request needs a unique MessageID. Managing MessageID manually is error-prone:

```go
// ❌ Not recommended: manual MessageID management
msgID := uint16(1)
req1 := dimse.NewCEchoRequest()
_ = req1.SetMessageID(msgID)
msgID++
req2, _ := dimse.NewCStoreRequest(dataset)
_ = req2.SetMessageID(msgID)
msgID++
// ... easy to forget to increment or to duplicate
```

### Solution: MessageIDGenerator

Use `MessageIDGenerator` to manage MessageIDs automatically:

```go
// ✅ Recommended: use MessageIDGenerator
idGen := dimse.NewMessageIDGenerator()

// Create requests without a MessageID; the generator assigns one before sending.
req1 := dimse.NewCEchoRequest()
req2, _ := dimse.NewCStoreRequest(dataset)
req3 := dimse.NewCFindRequestPatientRoot(level, query)

// Assign a MessageID automatically before sending
idGen.AssignMessageID(req1)  // MessageID = 1
idGen.AssignMessageID(req2)  // MessageID = 2
idGen.AssignMessageID(req3)  // MessageID = 3
```

### Features

1. **Thread-safe**: can be called concurrently
2. **Auto-increment**: starts at 1, range 1-65535
3. **Skips 0 automatically**: MessageID=0 means unassigned
4. **Keeps existing values**: does not overwrite an existing MessageID
5. **Supports reset**: the counter can be reset on a new connection

## Basic Usage

### C-ECHO (verify connectivity)

```go
package main

import (
    "github.com/cocosip/go-dicom/pkg/network/dimse"
)

func main() {
    // Create a MessageID generator
    idGen := dimse.NewMessageIDGenerator()

    // Create a C-ECHO request (MessageID=0 means unassigned)
    req := dimse.NewCEchoRequest()

    // Assign a MessageID
    msgID, err := idGen.AssignMessageID(req)
    if err != nil {
        panic(err)
    }

    // Now req can be sent; its MessageID has been set to 1
    println("MessageID:", msgID)  // Output: MessageID: 1
}
```

### C-STORE (store DICOM instances)

```go
package main

import (
    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
)

func main() {
    idGen := dimse.NewMessageIDGenerator()

    // Create a DICOM dataset
    ds := dataset.New()
    ds.Add(element.NewString(tag.SOPClassUID, vr.UI,
        []string{"1.2.840.10008.5.1.4.1.1.2"}))  // CT Image Storage
    ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI,
        []string{"1.2.3.4.5.6.7.8.9"}))
    ds.Add(element.NewString(tag.PatientName, vr.PN,
        []string{"Doe^John"}))

    // Create a C-STORE request
    req, err := dimse.NewCStoreRequest(ds)
    if err != nil {
        panic(err)
    }

    // Set priority (optional)
    req.SetPriority(uint16(dimse.PriorityHigh))

    // Assign a MessageID
    msgID, _ := idGen.AssignMessageID(req)
    println("MessageID:", msgID)
}
```

### C-FIND (query the DICOM archive)

```go
package main

import (
    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
)

func main() {
    idGen := dimse.NewMessageIDGenerator()

    // Create a query dataset
    query := dataset.New()
    query.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^*"}))
    query.Add(element.NewString(tag.PatientID, vr.LO, []string{""}))

    // Create a C-FIND request (Patient Root Query)
    req := dimse.NewCFindRequestPatientRoot(
        dimse.QueryRetrieveLevelPatient,
        query,
    )

    // Assign a MessageID
    msgID, _ := idGen.AssignMessageID(req)
    println("MessageID:", msgID)
}
```

## Integration in Client/Association

In real applications, MessageIDGenerator should be managed by the Client or Association:

```go
type Client struct {
    messageIDGen *dimse.MessageIDGenerator
    // ... other fields
}

func NewClient() *Client {
    return &Client{
        messageIDGen: dimse.NewMessageIDGenerator(),
    }
}

func (c *Client) Send(msg dimse.Message) error {
    // Assign the MessageID automatically
    _, err := c.messageIDGen.AssignMessageID(msg)
    if err != nil {
        return err
    }

    // Send the message
    // ...
    return nil
}

// User code
func main() {
    client := NewClient()

    // No need to worry about MessageID!
    req := dimse.NewCEchoRequest()
    client.Send(req)
}
```

## Response Handling

### C-ECHO response

```go
// Create a success response
resp := dimse.NewCEchoResponseSuccess(requestMessageID)

// Check status
if resp.IsSuccess() {
    println("Echo successful")
}
```

### C-STORE response

```go
resp := dimse.NewCStoreResponseSuccess(
    requestMessageID,
    sopClassUID,
    sopInstanceUID,
)

if resp.IsSuccess() {
    println("Store successful")
}
```

### C-FIND response

```go
// Pending response (contains query results)
resp := dimse.NewCFindResponsePending(
    requestMessageID,
    sopClassUID,
    identifier,
)

if resp.IsPending() {
    // Handle the query results
    if resp.HasIdentifier() {
        data := resp.DataDataset()
        // ... process the data
    }
}

// Final response (query complete)
finalResp := dimse.NewCFindResponseSuccess(requestMessageID, sopClassUID)
if finalResp.IsSuccess() {
    println("Query completed")
}
```

### Optional response UIDs

`Affected SOP Class UID` and `Affected SOP Instance UID` are Type U(=) in
C-STORE-RSP, and `Affected SOP Class UID` is Type U(=) in C-FIND-RSP (PS3.7
9.1.1.1.3/9.1.1.1.4, 9.1.2.1.3). The constructors write the element only when
the argument is non-empty:

```go
// Minimal C-STORE-RSP (Status + MessageIDBeingRespondedTo only)
resp := dimse.NewCStoreResponse(requestMessageID, status.Success, "", "")
resp := dimse.NewCFindResponse(requestMessageID, status.Success, "", nil)

// With the optional UIDs, as before
resp := dimse.NewCStoreResponse(requestMessageID, status.Success, sopClassUID, sopInstanceUID)
```

`NewCStoreResponseFromRequest` / `NewCFindResponseFromRequest` propagate the
request's UIDs, so absence round-trips. The service accepts responses that omit
these optional elements instead of aborting the association.

### Required-field validation

`ValidateRequestForSend(req)` is called by the service before encoding a request.
It rejects an empty value for every mandatory identifier (C-STORE SOP
Class/Instance, C-FIND/C-GET SOP Class, C-MOVE SOP Class and destination, N-GET/
N-SET/N-DELETE/N-ACTION requested Class/Instance, N-CREATE Class, N-EVENT-REPORT
Class/Instance plus Event Type ID presence). C-ECHO and C-CANCEL always pass, and
the N-CREATE instance UID may be empty so the SCP can assign it.

`NewCStoreRequest` additionally rejects a dataset whose `SOPClassUID` or
`SOPInstanceUID` element exists but is zero-length. The validator switches on
the built-in request types; a custom `Request` implementation is not
introspected and is sent as-is.

### Unknown status

`BaseResponse.Status()` reports a response that lacks the Status element
(0000,0900) as a generic Failure (`status.LookupStatus(0xFFFF)`), so
`IsFailure()` is true. Predicates never all return false.

### SetAffectedSOPClassUID

For C-FIND, C-GET, and C-MOVE requests, `SetAffectedSOPClassUID` updates both the
command element and the cached field, so `AffectedSOPClassUID()` and `String()`
return the new UID after a request is rebuilt from an inbound command set.

## Advanced Features

### Priority

```go
req, _ := dimse.NewCStoreRequest(dataset)
req.SetPriority(uint16(dimse.PriorityHigh))    // High priority
req.SetPriority(uint16(dimse.PriorityMedium))  // Medium priority (default)
req.SetPriority(uint16(dimse.PriorityLow))     // Low priority
```

### Move Originator (C-MOVE sub-operations)

```go
req, _ := dimse.NewCStoreRequest(dataset)
req.SetMoveOriginator("MOVE_SCU", 999)
```

### Reset the MessageID counter

```go
idGen := dimse.NewMessageIDGenerator()

// After using it for a while...
idGen.Next()  // 1
idGen.Next()  // 2

// A new connection starts; reset the counter
idGen.Reset()
idGen.Next()  // 1 (starts over)
```

## Higher-level Client API

For most applications the package-level messages above are created for you by
`github.com/cocosip/go-dicom/pkg/network/client`:

```go
c := client.New(client.WithCallingAE("MY_SCU"), client.WithCalledAE("PACS"))
if err := c.AddPresentationContextWithRoles(
    "1.2.840.10008.5.1.4.1.1.2", true, false, "1.2.840.10008.1.2.1",
); err != nil {
    return err
}
if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
    return err
}
defer c.Close()

if err := c.CEcho(ctx); err != nil {
    return err
}
```

The client exposes `CEcho`, `CStore`, `CFind`, `CMove`, `CGet`, and the six
N-Service methods; `client.ManagedClient` provides a queued alternative. See
[docs/network/README.md](../../../docs/network/README.md) for the full guide.

## Testing

Run the tests:

```bash
go test ./pkg/network/dimse
go test -v ./pkg/network/dimse -run TestMessageIDGenerator
```

## References

- DICOM Standard: https://www.dicomstandard.org/
- DICOM Part 7: Message Exchange
- DICOM Part 8: Network Communication Support
