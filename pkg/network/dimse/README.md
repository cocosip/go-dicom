# DIMSE Package

DIMSE (DICOM Message Service Element) is the core message layer of DICOM network communication.

## Features

- ✅ C-ECHO: verify connectivity
- ✅ C-STORE: store DICOM instances
- ✅ C-FIND: query the DICOM archive
- ✅ Automatic MessageID management
- ✅ Priority support
- ✅ Status code handling

## Automatic MessageID Management

### Problem
In DICOM network communication, every request needs a unique MessageID. Managing MessageID manually is error-prone:

```go
// ❌ Not recommended: manual MessageID management
msgID := 1
req1 := dimse.NewCEchoRequest(msgID)
msgID++
req2 := dimse.NewCStoreRequest(msgID, dataset)
msgID++
// ... easy to forget to increment or to duplicate
```

### Solution: MessageIDGenerator

Use `MessageIDGenerator` to manage MessageIDs automatically:

```go
// ✅ Recommended: use MessageIDGenerator
idGen := dimse.NewMessageIDGenerator()

// Pass 0 when creating a request (means unassigned)
req1 := dimse.NewCEchoRequest(0)
req2 := dimse.NewCStoreRequest(0, dataset)
req3 := dimse.NewCFindRequestPatientRoot(0, level, query)

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
    req := dimse.NewCEchoRequest(0)

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
    req, err := dimse.NewCStoreRequest(0, ds)
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
        0,
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
    req := dimse.NewCEchoRequest(0)
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

## Advanced Features

### Priority

```go
req := dimse.NewCStoreRequest(0, dataset)
req.SetPriority(uint16(dimse.PriorityHigh))    // High priority
req.SetPriority(uint16(dimse.PriorityMedium))  // Medium priority (default)
req.SetPriority(uint16(dimse.PriorityLow))     // Low priority
```

### Move Originator (C-MOVE sub-operations)

```go
req, _ := dimse.NewCStoreRequest(0, dataset)
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

## Future Development

In Phase 9.7 (Client API), we will provide a higher-level API:

```go
// Future API (coming soon)
client := dicom.NewClient("MY_SCU")
client.Connect("localhost", 11112, "REMOTE_SCP")

// No need to create Request objects at all!
err := client.Echo()
err := client.Store(dataset)
results := client.Find(query)
```

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
