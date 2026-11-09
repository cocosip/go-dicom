# DICOM Networking Guide

This guide is the single, consolidated reference for go-dicom's networking layer: the SCU (Service Class User) client, the SCP (Service Class Provider) server, the DIMSE message model, association management, and the operational details (TLS, timeouts, concurrency, status codes, performance) that apply across all of them.

## Table of Contents

- [Overview](#overview)
- [Architecture / Package Structure](#architecture--package-structure)
  - [Layered Architecture](#layered-architecture)
  - [Package Structure](#package-structure)
- [Quick Start](#quick-start)
- [Client (SCU)](#client-scu)
  - [Two Client Usage Styles: Session vs ManagedClient Queue](#two-client-usage-styles-session-vs-managedclient-queue)
  - [Creating a Client](#creating-a-client)
  - [Client Configuration Options](#client-configuration-options)
  - [Adding Presentation Contexts](#adding-presentation-contexts)
  - [Requesting SCU/SCP Roles and Factory Methods](#requesting-scuscpp-roles-and-factory-methods)
  - [C-GET Must Negotiate the Storage SCP Role First](#c-get-must-negotiate-the-storage-scp-role-first)
  - [Connecting and Closing](#connecting-and-closing)
  - [DIMSE Operations](#dimse-operations)
  - [Convenience Methods](#convenience-methods)
  - [Getting Association Information](#getting-association-information)
  - [Client Examples](#client-examples)
- [Server (SCP)](#server-scp)
  - [Creating a Server](#creating-a-server)
  - [Server Configuration Options](#server-configuration-options)
  - [DIMSE Handlers](#dimse-handlers)
  - [Lifecycle Handlers](#lifecycle-handlers)
  - [Starting and Stopping the Server](#starting-and-stopping-the-server)
  - [Server Examples](#server-examples)
- [DIMSE Messages](#dimse-messages)
  - [Message Structure](#message-structure)
  - [Automatic MessageID Management](#automatic-messageid-management)
  - [Creating Request Messages](#creating-request-messages)
  - [Outbound Validation](#outbound-validation)
  - [Creating Response Messages](#creating-response-messages)
  - [Checking Response Status](#checking-response-status)
  - [Response Handling](#response-handling)
  - [Priority, Move Originator, and Resetting the Counter](#priority-move-originator-and-resetting-the-counter)
- [Association Management](#association-management)
- [Advanced Features](#advanced-features)
  - [TLS Encrypted Connections](#tls-encrypted-connections)
  - [Association Negotiation Control](#association-negotiation-control)
  - [Observability](#observability)
  - [Timeout Control](#timeout-control)
  - [Concurrency and Timeout Contract](#concurrency-and-timeout-contract)
  - [Concurrent Processing](#concurrent-processing)
- [Status Codes / Error Handling](#status-codes--error-handling)
  - [Status Code Reference](#status-code-reference)
  - [Common Error Types](#common-error-types)
  - [Status Code Errors](#status-code-errors)
  - [Error Wrapping](#error-wrapping)
- [Performance Tips](#performance-tips)
- [Next Steps](#next-steps)
- [References](#references)

## Overview

This package provides the SCU (Service Class User) and SCP (Service Class Provider) foundation for the DICOM network protocol stack.

The DICOM network protocol is based on TCP/IP, using the Association mechanism for connection management and DIMSE (DICOM Message Service Element) for message exchange. This implementation supports the network services implemented in this repository; external PACS interoperability still depends on the negotiated SOP classes and transfer syntaxes.

### Key Features

- ✅ **Full PDU support**: 7 PDU types (A-ASSOCIATE-RQ/AC/RJ, P-DATA-TF, A-RELEASE-RQ/RP, A-ABORT)
- ✅ **Association management**: automatic negotiation of Presentation Contexts and user-information parameters, with a per-direction Maximum PDU Length advertisement (PS3.7 D.3.3.1)
- ✅ **DIMSE message model and covered service workflows**:
  - C-ECHO (connectivity verification)
  - C-STORE (store images)
  - C-FIND (query)
  - C-MOVE (retrieve — Pull mode)
  - C-GET (retrieve — Push mode)
  - N-* operations (N-EVENT-REPORT, N-GET, N-SET, N-ACTION, N-CREATE, N-DELETE)
- ✅ **SCU client**: high-level API supporting synchronous and asynchronous operations
- ✅ **SCP server**: handler mode with support for custom business logic
- ✅ **TLS support**: secure connections (the SNI/verification name defaults to the dial address when not set)
- ✅ **Client lifecycle concurrency control**: `Connect`, `Close`, `IsConnected`, and
  `GetAssociation` can be called concurrently; configuration and Presentation Contexts can only be modified before the first connection
- ✅ **Robust against conformant-but-minimal peers**: optional Type U(=) response
  parameters may be absent, unrequested role selections are ignored (PS3.7 D.3.3.4),
  and PDU decoding validates User Identity sub-items and trims UID padding
- ✅ **Failure containment**: handler panics are recovered per request, graceful
  close is bounded by the caller's context, timed-out/cancelled progressive
  operations send C-CANCEL, and the server rejects over-limit connections with
  A-ASSOCIATE-RJ instead of leaving them hanging

### Test Coverage

- **647 unit test functions** across the network packages, all passing
- Includes normal, error, concurrency, and race-detector scenarios
- Statement coverage per package is roughly 73–89%
  (`go test ./pkg/network/... -cover`); `client` and `service` are the lowest

## Architecture / Package Structure

### Layered Architecture

```
┌─────────────────────────────────────────────────────┐
│              Application Layer                       │
│         (Your DICOM Application)                     │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│            Client / Server API                       │
│  pkg/network/client  │  pkg/network/server          │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│              Service Layer                           │
│  pkg/network/service (State Machine + Handlers)     │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│              DIMSE Layer                             │
│  pkg/network/dimse (Message Encoding/Decoding)      │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│           Association Layer                          │
│  pkg/network/association (Negotiation)              │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│              PDU Layer                               │
│  pkg/network/pdu (Protocol Data Units)              │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│           Transport Layer                            │
│  pkg/network/transport (TCP/TLS)                    │
└─────────────────────────────────────────────────────┘
```

### Package Structure

```
pkg/network/
├── status/         # DICOM status code definitions
├── pdu/            # PDU encoding/decoding
├── association/    # Association management
├── dimse/          # DIMSE message definitions
├── transport/      # Network transport (TCP/TLS)
├── service/        # Core service layer (state machine, message handling)
├── client/         # SCU client API
└── server/         # SCP server API
```

## Quick Start

Create a client with `client.New()` and a server with `server.New()`, then add Presentation Contexts, connect, and send a request. The complete runnable programs are in [Client Examples](#client-examples) and [Server Examples](#server-examples).

```go
import "github.com/cocosip/go-dicom/pkg/network/client"

// Basic usage
c := client.New()
```

```go
import "github.com/cocosip/go-dicom/pkg/network/server"

// Basic usage
srv := server.New()
```

## Client (SCU)

### Two Client Usage Styles: Session vs ManagedClient Queue

go-dicom provides two client APIs; choose by scenario:

| | `client.Client` (session-style) | `client.ManagedClient` (queue-style, analogous to fo-dicom `DicomClient`) |
|---|---|---|
| Connection timing | `Connect` first, then call methods | `Add(job)` first; connects only on `Send` |
| Operation style | `c.CGet(...)` / `c.CFind(...)` synchronous methods | Job objects such as `NewCGetJob(...)` / `NewCFindJob(...)` |
| Context source | Manually call `AddPresentationContext*` before `Connect` | Merged from each Job's `PresentationContexts()` on `Send` |
| Suitable for | One-off, sequential operations, simple scripts | Batch, connection reuse, a style closer to fo-dicom |

`client.Client` suits scenarios where the caller needs to control Presentation
Contexts, connection timing, or advanced association negotiation itself.
`client.ManagedClient` retains that low-level API while, for common SCU workflows,
deriving contexts per Job, batching in FIFO order, splitting associations at the
128-context limit, and reusing compatible associations within a short linger
period.

> When doing C-GET either way, **you must negotiate the SCP role for the Storage SOP Class you want to receive** (see [C-GET Must Negotiate the Storage SCP Role First](#c-get-must-negotiate-the-storage-scp-role-first)), otherwise the SCP cannot send anything back and you only get 0xB000.

```go
managed := client.NewManaged(
    client.WithMaximumRequestsPerAssociation(32),
    client.WithAssociationLingerTimeout(50*time.Millisecond),
)
if err := managed.Add(client.NewCEchoJob(func(err error) {
    log.Printf("C-ECHO complete: %v", err)
})); err != nil {
    return err
}
if err := managed.Send(ctx, "pacs.example", 104); err != nil {
    return err
}
defer managed.Close()
```

Managed Job constructors are available for C-ECHO, C-STORE, C-FIND, C-MOVE,
C-GET and all six N-Service requests. Dataset-bearing Jobs deep-clone their
input while queued, so callers may safely reuse or change their source Dataset
after constructing the Job. `Close` rejects future Jobs, completes queued Jobs,
and aborts an active association.

`Send` drains queued Jobs in FIFO association batches. If the association dies
between Jobs, the remaining **not yet attempted** Jobs continue on a freshly
negotiated association; a Job whose execution already returned an error keeps
that error because it may have been processed by the peer (no implicit retry).
Death is detected from the client/service state and from connection-class job
errors (`ErrClientNotConnected`, `ErrClientClosing`, `service.ErrServiceClosed`,
`net.ErrClosed`, `io.EOF`, `io.ErrClosedPipe`, `*net.OpError`); a cancelled or
timed-out context is a caller decision, never a reconnect trigger.
Association-timeout connect retries remain bounded by
`WithMaximumConsecutiveAssociationTimeouts`.

Progressive C-FIND, C-MOVE, and C-GET operations propagate request timeout,
transport, and association-close errors. Their callbacks send C-CANCEL when
they return `false` and wait for the final response so the Association remains
usable. Cancelling the caller context, or letting `WithRequestTimeout` expire,
makes the **service layer** send one bounded best-effort C-CANCEL for those
three operations; the client wrapper does not send a second one. Resources
(pending request and asynchronous-operation slot) are released before the
terminal event is delivered. Low-level `Service.SendCFind`, `SendCMove`, and
`SendCGet` calls return a single ordered `ResponseEvent` stream: each event
contains either a DIMSE response or the one terminal error. The stream never
closes silently before a final response or terminal error, and callers must keep
draining it until it closes even after cancelling the context. Managed Jobs
inherit the same behavior.

### Creating a Client

```go
import "github.com/cocosip/go-dicom/pkg/network/client"

// Basic usage
c := client.New()

// Using configuration options
c := client.New(
    client.WithCallingAE("MY_SCU"),
    client.WithCalledAE("PACS_SERVER"),
    client.WithMaxPDULength(16384),
    client.WithConnectTimeout(10*time.Second),
    client.WithRequestTimeout(30*time.Second),
    client.WithAssociationTimeout(10*time.Second),
)
```

### Client Configuration Options

| Option | Type | Default | Description |
|------|------|--------|------|
| `WithCallingAE(ae)` | string | "GO_DICOM_SCU" | Local AE Title |
| `WithCalledAE(ae)` | string | "ANY_SCP" | Remote AE Title |
| `WithMaxPDULength(length)` | uint32 | 16384 | Maximum PDU length this AE can receive (per direction) |
| `WithMaxMessageSize(size)` | uint32 | 512 MiB | Bound for a reassembled inbound DIMSE message (0=unlimited) |
| `WithConnectTimeout(timeout)` | time.Duration | 10s | TCP connection timeout |
| `WithRequestTimeout(timeout)` | time.Duration | 30s | DIMSE request timeout |
| `WithTransportReadTimeout(timeout)` | time.Duration | 0 | Single PDU read timeout; 0 allows an established association to stay idle |
| `WithTransportWriteTimeout(timeout)` | time.Duration | 30s | Single PDU write timeout |
| `WithAssociationTimeout(timeout)` | time.Duration | 10s | Association negotiation timeout |
| `WithImplementationClassUID(uid)` | string | "1.2.826.0.1.3680043.8.498.1" | Implementation class UID |
| `WithImplementationVersionName(name)` | string | "go-dicom-1.0" | Implementation version name |
| `WithTransferSyntaxRegistry(registry)` | `*transfer.Registry` | Independent standard catalog view | Explicitly inject a custom Transfer Syntax; standard and unknown syntaxes can still be parsed when not registered |

### Adding Presentation Contexts

A Presentation Context defines the supported SOP Classes and Transfer Syntaxes:

```go
// Verification SOP Class
if err := c.AddPresentationContext(
    "1.2.840.10008.1.1",   // Abstract Syntax (SOP Class UID)
    "1.2.840.10008.1.2",   // Transfer Syntax: Implicit VR Little Endian
    "1.2.840.10008.1.2.1", // Transfer Syntax: Explicit VR Little Endian
); err != nil { return err }

// CT Image Storage
if err := c.AddPresentationContext(
    "1.2.840.10008.5.1.4.1.1.2",
    "1.2.840.10008.1.2.1",
); err != nil { return err }

// Patient Root Query/Retrieve
if err := c.AddPresentationContext(
    "1.2.840.10008.5.1.4.1.2.1.1",
    "1.2.840.10008.1.2.1",
); err != nil { return err }
```

Standard Transfer Syntax entries come from an immutable catalog. To use an
application-defined custom syntax, first call `Register` or `Replace` on the
instance returned by `transfer.NewRegistry()`, then inject it via
`WithTransferSyntaxRegistry`; `transfer.NewBuilder(...).Build()` only constructs
objects and does not modify any process-level state.

**Common SOP Class UIDs**:

| SOP Class | UID |
|-----------|-----|
| Verification | 1.2.840.10008.1.1 |
| CT Image Storage | 1.2.840.10008.5.1.4.1.1.2 |
| MR Image Storage | 1.2.840.10008.5.1.4.1.1.4 |
| Patient Root Q/R - FIND | 1.2.840.10008.5.1.4.1.2.1.1 |
| Study Root Q/R - FIND | 1.2.840.10008.5.1.4.1.2.2.1 |
| Patient Root Q/R - MOVE | 1.2.840.10008.5.1.4.1.2.1.2 |
| Study Root Q/R - MOVE | 1.2.840.10008.5.1.4.1.2.2.2 |

### Requesting SCU/SCP Roles and Factory Methods

A plain `AddPresentationContext` does not declare roles (default: requesting side SCU, accepting side SCP). For C-GET to receive C-STORE sent back by the SCP, you must declare the **SCP role** for the Storage SOP Class.

Role responses are applied defensively: a role selection item for a SOP Class
this AE did not propose is ignored (PS3.7 D.3.3.4: "The Association-requester
shall ignore the response if it has not proposed the role"), and a role the
acceptor granted beyond the proposal is clamped to the proposed value. Neither
case closes the association; only an invalid role value (> 1) is treated as a
protocol error.

```go
// Single: explicit role
c.AddPresentationContextWithRoles(uid.CTImageStorage.UID(), false, true, ts...)

// Batch: explicit UID list
c.AddStoragePresentationContexts([]string{uid.CTImageStorage.UID(), uid.MRImageStorage.UID()}, ts...)

// Negotiate all at once by storage category (equivalent to fo-dicom's
// DicomPresentationContext.GetScpRolePresentationContextsFromStorageUids)
c.AddStoragePresentationContextsByCategory(uid.StorageCategoryImage, ts...)
```

When you need to assemble contexts yourself (including private SOP Classes and explicit roles), use **value objects + factories**. `PresentationContextSpec` is reused by both `Client` and `ManagedClient`:

```go
specs := []client.PresentationContextSpec{
    client.NewPresentationContextSpec(uid.StudyRootQueryRetrieveInformationModelGet.UID(), false, false, ts...),
    client.NewPresentationContextSpec(uid.CTImageStorage.UID(), false, true, ts...),
}
if err := c.AddPresentationContexts(specs); err != nil { return err } // the session-style client accepts specs
// queue-style Jobs internally use the same PresentationContextSpec type
```

Factory functions: `client.NewPresentationContextSpec`, `client.StoragePresentationContextSpecs(category, ts...)`, `client.StoragePresentationContextSpecsFor(uids, ts...)`. This way, adding a new SOP Class/role policy only requires producing specs, without changing the client implementation.

### C-GET Must Negotiate the Storage SCP Role First

C-GET has the SCP send C-STORE sub-operations back over the same association. Per
the DICOM standard, the requester must declare the SCP role for each Storage SOP
Class that may be returned and configure a `CStoreHandler`. Complete this
negotiation before connecting, using `Client.AddStoragePresentationContexts` (or
the one-shot-by-category `Client.AddStoragePresentationContextsByCategory`, or
the lower-level `AddPresentationContextWithRoles`); otherwise the SCP cannot send
anything back and C-GET will only get 0xB000. When these prerequisites are
missing, `CGet` returns a readable error directly instead of silently receiving 0
files. Managed's `NewCGetJob` can complete the same negotiation through optional
storage SOP Class parameters.

Two styles (session-style / queue-style):

```go
// Session style
c := client.New(client.WithCStoreHandler(handleCStore))
c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelGet.UID(), ts...)
c.AddStoragePresentationContextsByCategory(uid.StorageCategoryImage, ts...)
c.Connect(ctx, host, port)
c.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier, cb)

// Queue style (≈ fo-dicom's AddRequest + Send)
managed := client.NewManaged(client.WithBaseClientOptions(
    client.WithCStoreHandler(handleCStore),
    client.WithCallingAE("MY_SCU"), client.WithCalledAE("QR_SCP"),
))
job, _ := client.NewCGetJob(dimse.QueryRetrieveLevelStudy, identifier, cb,
    func(err error) {}, uid.StorageSOPClassUIDs(uid.StorageCategoryImage)...)
managed.Add(job)
managed.Send(ctx, host, port)
```

### Connecting and Closing

```go
// Connect to the server
ctx := context.Background()
if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
    return err
}

// Check connection status
if c.IsConnected() {
    // connected
}

// Graceful close (sends A-RELEASE-RQ)
if err := c.Close(); err != nil {
    return err
}

// Forceful abort (sends A-ABORT)
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
if err := c.Abort(ctx); err != nil {
    return err
}
```

After an A-RELEASE-RP is sent, the association cannot carry any further P-DATA
(PS3.8 7.2.2), so the transport is closed. When both AEs request release at the
same time, the release collision is resolved per PS3.8: the association-requestor
answers first, and the association-acceptor waits for the peer's confirmation
before sending its own A-RELEASE-RP. An A-RELEASE-RP that was never requested is
treated as an abnormal condition and aborts the association (A-ABORT, AA-8).

`Close` is bounded: it gives the release handshake a 5-second context and does
not wait for the service's handler-drain timeout (60 seconds by default) once
that context expires. The connection is closed in every case, so an unresponsive
peer or a long-running inbound handler cannot block the caller indefinitely.
When the peer sends A-ABORT, the association is reported as `StateAborted` (not
`StateClosed`), and `CloseError()` records the received A-ABORT.

### DIMSE Operations

All operations are available as synchronous methods on `client.Client`. Complete runnable programs for the common operations are in [Client Examples](#client-examples).

**C-ECHO (verifying the connection):** call `c.CEcho(ctx)`; the basic usage is `if err := c.CEcho(ctx); err != nil { return fmt.Errorf("C-ECHO failed: %w", err) }`.

**C-STORE (storing images):** `c.CStore(ctx, result.Dataset)` stores a parsed dataset. Two additional forms:

```go
import "github.com/cocosip/go-dicom/pkg/network/dimse"

// Store with priority
if err := c.CStoreWithPriority(ctx, result.Dataset, dimse.PriorityHigh); err != nil {
    return err
}

// Batch store
datasets := []*dataset.Dataset{ds1, ds2, ds3}
count, err := c.CStoreMultiple(ctx, datasets)
if err != nil {
    log.Printf("Stored %d out of %d datasets: %v\n", count, len(datasets), err)
}
```

`CStore` prefers an accepted presentation context matching the Dataset's source
transfer syntax. If the peer accepts only another syntax, a registered codec is
used to transcode a copy; the caller's Dataset is not modified. Parsed Datasets
already retain their source syntax. For a programmatically constructed Dataset
containing Pixel Data, use `dataset.NewWithTransferSyntax` so the byte order or
encapsulation is known.

`CStoreMultiple` sends in slice order without concurrency. It stops at the first
failure or cancellation, and `count` reports only completed successful stores.

**C-FIND (query):** `c.CFind` collects all results; `c.CFindWithCallback` streams them (see [Client Examples](#client-examples)).

**Query/Retrieve Level**:

```go
dimse.QueryRetrieveLevelPatient  // PATIENT level
dimse.QueryRetrieveLevelStudy    // STUDY level
dimse.QueryRetrieveLevelSeries   // SERIES level
dimse.QueryRetrieveLevelImage    // IMAGE level
```

**DICOM query wildcards**:

| Wildcard | Meaning | Example |
|--------|------|------|
| `*` | Matches any characters | "DOE*" matches "DOE", "DOES", "DOE123" |
| `?` | Matches a single character | "DOE?" matches "DOES", not "DOE" |
| `-` | Date range | "20240101-20241231" |
| `` | Empty string | Returns the value of that field |

**C-MOVE (retrieve - Pull mode)**:

```go
// Build a retrieval identifier
identifier := dataset.New()
identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

// Execute C-MOVE (send images to a third party)
err := c.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST_AE", identifier,
    func(remaining, completed, failed, warning uint16) bool {
        fmt.Printf("Progress: %d remaining, %d completed, %d failed, %d warning\n",
            remaining, completed, failed, warning)
        return true  // continue receiving progress updates
    })
```

**C-GET (retrieve - Push mode):** C-GET sends images directly to this client over the same association. Both of the following must be satisfied, otherwise the SCP cannot send anything back and you only get 0xB000: (1) set a C-STORE handler to receive the images sent back; (2) before connecting, declare the SCP role for every Storage SOP Class that may be sent back (`AddStoragePresentationContexts` automatically requests the SCP role). See [C-GET Must Negotiate the Storage SCP Role First](#c-get-must-negotiate-the-storage-scp-role-first) for both session-style and queue-style forms.

### Convenience Methods

```go
// Patient Root Query
results, err := c.CFindPatientRoot(ctx, dimse.QueryRetrieveLevelStudy, query)

// Study Root Query
results, err := c.CFindStudyRoot(ctx, dimse.QueryRetrieveLevelStudy, query)
```

### Getting Association Information

```go
// Get the current Association
assoc := c.GetAssociation()
if assoc != nil {
    fmt.Printf("CallingAE: %s\n", assoc.CallingAE)
    fmt.Printf("CalledAE: %s\n", assoc.CalledAE)
    // Maximum PDU length the peer advertised for what it can receive.
    // Each direction has its own independent limit (PS3.7 D.3.3.1).
    fmt.Printf("Peer MaxPDULength: %d\n", assoc.MaxPDULength)

    // Get the accepted Presentation Contexts
    accepted := assoc.GetAcceptedPresentationContexts()
    for _, pc := range accepted {
        transferSyntax := ""
        if pc.AcceptedTransferSyntax != nil {
            transferSyntax = pc.AcceptedTransferSyntax.UID().UID()
        }
        fmt.Printf("Context ID %d: %s (%s)\n",
            pc.ID, pc.AbstractSyntax, transferSyntax)
    }
}

// Get the configuration
config := c.GetConfig()
fmt.Printf("CallingAE: %s\n", config.CallingAE)
```

### Client Examples

#### 1. C-ECHO (verify connection)

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    // Create the client
    c := client.New(
        client.WithCallingAE("MY_SCU"),
        client.WithCalledAE("PACS_SERVER"),
    )

    // Add Presentation Context (Verification SOP Class)
    if err := c.AddPresentationContext(
        "1.2.840.10008.1.1",   // Verification SOP Class
        "1.2.840.10008.1.2",   // Implicit VR Little Endian
        "1.2.840.10008.1.2.1", // Explicit VR Little Endian
    ); err != nil {
        log.Fatal(err)
    }

    // Connect to the server
    ctx := context.Background()
    if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // Send C-ECHO
    if err := c.CEcho(ctx); err != nil {
        log.Fatalf("C-ECHO failed: %v", err)
    }

    log.Println("C-ECHO successful!")
}
```

#### 2. C-STORE (store image)

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    // Read the DICOM file
    result, err := parser.ParseFile("image.dcm")
    if err != nil {
        log.Fatal(err)
    }

    // Create the client
    c := client.New(
        client.WithCallingAE("MY_SCU"),
        client.WithCalledAE("PACS_SERVER"),
    )

    // Add the Storage Presentation Context
    if err := c.AddPresentationContext(
        "1.2.840.10008.5.1.4.1.1.2", // CT Image Storage
        "1.2.840.10008.1.2.1",       // Explicit VR Little Endian
    ); err != nil {
        log.Fatal(err)
    }

    // Connect and send
    ctx := context.Background()
    if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // Send C-STORE
    if err := c.CStore(ctx, result.Dataset); err != nil {
        log.Fatalf("C-STORE failed: %v", err)
    }

    log.Println("Image stored successfully!")
}
```

#### 3. C-FIND (query)

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/client"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
)

func main() {
    // Create the client
    c := client.New(
        client.WithCallingAE("MY_SCU"),
        client.WithCalledAE("PACS_SERVER"),
    )

    // Add the Query/Retrieve Presentation Context
    if err := c.AddPresentationContext(
        "1.2.840.10008.5.1.4.1.2.2.1", // Study Root Query/Retrieve
        "1.2.840.10008.1.2.1",
    ); err != nil {
        log.Fatal(err)
    }

    // Connect
    ctx := context.Background()
    if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // Build the query criteria
    query := dataset.New()
    query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN*"}))
    query.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20240101-20241231"}))
    query.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{""})) // Return field

    // Execute the query
    results, err := c.CFind(ctx, dimse.QueryRetrieveLevelStudy, query)
    if err != nil {
        log.Fatalf("C-FIND failed: %v", err)
    }

    // Process the results
    log.Printf("Found %d studies\n", len(results))
    for i, result := range results {
        studyUID := result.TryGetString(tag.StudyInstanceUID)
        patientName := result.TryGetString(tag.PatientName)
        log.Printf("Study %d: %s - %s\n", i+1, studyUID, patientName)
    }
}
```

#### 4. Stream query results (avoid memory usage)

```go
// Use a callback to handle a large number of query results
err := c.CFindWithCallback(ctx, dimse.QueryRetrieveLevelStudy, query,
    func(result *dataset.Dataset) bool {
        // Process each result
        studyUID := result.TryGetString(tag.StudyInstanceUID)
        log.Printf("Found study: %s\n", studyUID)

        // Return true to keep receiving, false to stop
        return true
    })
```

## Server (SCP)

### Creating a Server

```go
import "github.com/cocosip/go-dicom/pkg/network/server"

// Basic usage
srv := server.New()

// Using configuration options
srv := server.New(
    server.WithAETitle("MY_SCP"),
    server.WithPort(11112),
    server.WithMaxPDULength(16384),
    server.WithMaxConnections(50),
    server.WithAssociationTimeout(30*time.Second),
    server.WithTransportReadTimeout(2*time.Minute),
)
```

### Server Configuration Options

| Option | Type | Default | Description |
|------|------|--------|------|
| `WithAETitle(ae)` | string | "GO_DICOM_SCP" | Local AE Title |
| `WithPort(port)` | int | 104 | Listening port |
| `WithMaxPDULength(length)` | uint32 | 16384 | Maximum PDU length this AE can receive (per direction) |
| `WithMaxMessageSize(size)` | uint32 | 512 MiB | Bound for a reassembled inbound DIMSE message (0=unlimited) |
| `WithMaxConnections(max)` | int | 0 | Maximum concurrent connections (0=unlimited); over-limit connections are rejected with A-ASSOCIATE-RJ(rejected-transient, local-limit-exceeded) |
| `WithAcceptTimeout(timeout)` | time.Duration | 0 | Single Accept polling deadline (0=obey only the Server context); expiry is silent and persistent accept errors back off up to 1s |
| `WithAssociationTimeout(timeout)` | time.Duration | 10s | Association negotiation timeout |
| `WithRequestTimeout(timeout)` | time.Duration | 30s | Response idle timeout for DIMSE requests actively initiated by the Server |
| `WithTransportReadTimeout(timeout)` | time.Duration | 0 | Single PDU read timeout (0=disabled) |
| `WithTransportWriteTimeout(timeout)` | time.Duration | 30s | Single PDU write timeout (0=disabled) |
| `WithAcceptedCallingAETitles(aes)` | []string | nil | Allowed Calling AEs (nil=allow all) |
| `WithStrictAECheck(strict)` | bool | false | Strict AE Title checking |
| `WithTLSConfig(config)` | *tls.Config | nil | TLS configuration |

### DIMSE Handlers

Handlers are registered per operation with `srv.SetCEchoHandler`, `srv.SetCStoreHandler`, `srv.SetCFindHandler`, `srv.SetCMoveHandler`, and `srv.SetCGetHandler`. The C-ECHO, C-STORE, and C-FIND handlers are shown in the complete programs in [Server Examples](#server-examples).

Handlers run in their own goroutine per inbound request. A panic inside a
handler is recovered and isolated: it is logged as a request failure with its
stack trace, the request lifecycle is completed with an error, and the affected
association is aborted with A-ABORT (service-provider, not-specified). A remote
peer therefore cannot crash the whole process by driving a handler into a
panic; other associations keep running. Handlers must still return promptly or
honor cancellation, because shutdown waits for them.

For C-FIND, C-MOVE, and C-GET, a peer C-CANCEL cancels the context the handler
received, but the inbound request registration is kept until the final
response has been sent. A handler that observes the cancellation may still
return a final response (or a failure reason); it is encoded against the
original MessageID. The registration is released only on the final response or
service shutdown, never on C-CANCEL alone.

The C-MOVE and C-GET handlers are:

#### C-MOVE Handler

```go
srv.SetCMoveHandler(func(ctx context.Context, req *dimse.CMoveRequest) ([]*dimse.CMoveResponse, error) {
    identifier := req.DataDataset()
    destination := req.MoveDestination()

    log.Printf("C-MOVE to %s\n", destination)

    // 1. Find matching images
    images := findImages(identifier)

    // 2. Create a connection to the destination
    destClient := client.New(
        client.WithCallingAE("MY_SCP"),
        client.WithCalledAE(destination),
    )
    // ... add presentation contexts
    destClient.Connect(ctx, getDestinationHost(destination), 104)
    defer destClient.Close()

    // 3. Send images
    responses := make([]*dimse.CMoveResponse, 0)
    remaining := uint16(len(images))
    completed := uint16(0)

    for _, image := range images {
        // Send a Pending response
        remaining--
        resp := dimse.NewCMoveResponsePending(req, remaining, completed, 0, 0)
        responses = append(responses, resp)

        // Send the image to the destination
        if err := destClient.CStore(ctx, image); err != nil {
            log.Printf("Failed to move image: %v\n", err)
            continue
        }
        completed++
    }

    // Final success response
    finalResp := dimse.NewCMoveResponseSuccess(req)
    responses = append(responses, finalResp)

    return responses, nil
})
```

#### C-GET Handler

```go
srv.SetCGetHandler(func(ctx context.Context, req *dimse.CGetRequest) ([]*dimse.CGetResponse, error) {
    // C-GET is like C-MOVE, but sends directly to the requester
    // The implementation is similar to the C-MOVE Handler
    // ...
    return responses, nil
})
```

### Lifecycle Handlers

#### Association release

```go
srv.SetAssociationReleaseHandlerFunc(func(ctx context.Context) error {
    log.Println("Client requesting release")

    // You can do cleanup work here
    // Return nil to accept the release, return an error to reject it
    return nil
})
```

#### Connection lifecycle

```go
srv.SetConnectionLifecycleHandlerFuncs(
    // OnAbort
    func(ctx context.Context, source byte, reason byte) {
        log.Printf("Connection aborted: source=%d, reason=%d\n", source, reason)
        // do cleanup work
    },
    // OnConnectionClosed
    func(ctx context.Context, err error) {
        if err != nil {
            log.Printf("Connection closed with error: %v\n", err)
        } else {
            log.Println("Connection closed normally")
        }
        // do cleanup work
    },
)
```

Association negotiation is also a lifecycle concern; see [Association Negotiation Control](#association-negotiation-control) for the negotiator function.

### Starting and Stopping the Server

```go
// Start the server (blocking)
ctx := context.Background()
if err := srv.ListenAndServe(ctx); err != nil {
    log.Fatal(err)
}

// Gracefully shut down in another goroutine
go func() {
    <-shutdownChan
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Printf("Graceful shutdown did not finish: %v\n", err)
        if closeErr := srv.Close(); closeErr != nil {
            log.Printf("Close error: %v\n", closeErr)
        }
    }
}()
```

`Shutdown` first stops accepting and waits for the accept loop to finish, then
waits for existing Associations to complete on their own, so no new processing
tasks are accepted after it starts waiting for connections. It runs under the
supplied deadline-bound context; when that context expires it does not force
transfers to be interrupted, so call `Close` to immediately close still-active
inbound connections. `Close` immediately closes the listener and all active
inbound connections, without sending A-RELEASE-RQ and without waiting for
running handlers to return.

An inbound A-ASSOCIATE-RQ must declare support for DICOM UL Protocol Version 1
(bit 0). If bit 0 is not set, the server returns an ACSE
`protocol-version-not-supported` rejection before invoking the application
negotiator.

```go
shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

if err := srv.Shutdown(shutdownCtx); err != nil {
    if closeErr := srv.Close(); closeErr != nil {
        log.Printf("force close failed: %v", closeErr)
    }
}
```

### Server Examples

#### Basic Echo SCP

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/network/dimse"
    "github.com/cocosip/go-dicom/pkg/network/server"
    "github.com/cocosip/go-dicom/pkg/network/service"
    "github.com/cocosip/go-dicom/pkg/network/status"
)

func main() {
    // Create the server
    srv := server.New(
        server.WithAETitle("MY_SCP"),
        server.WithPort(11112),
        server.WithMaxPDULength(16384),
    )

    // Set the C-ECHO handler
    srv.SetCEchoHandler(func(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
        log.Println("Received C-ECHO request")
        return dimse.NewCEchoResponseFromRequest(req, status.Success), nil
    })

    // Start the server
    log.Println("Starting Echo SCP on port 11112...")
    ctx := context.Background()
    if err := srv.ListenAndServe(ctx); err != nil {
        log.Fatal(err)
    }
}
```

#### Storage SCP (receive images)

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/writer"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
    "github.com/cocosip/go-dicom/pkg/network/server"
)

func main() {
    srv := server.New(
        server.WithAETitle("STORAGE_SCP"),
        server.WithPort(11112),
    )

    // Set the C-STORE handler
    srv.SetCStoreHandler(func(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
        ds := req.DataDataset()

        // Get the SOP Instance UID
        sopInstanceUID := ds.TryGetString(tag.SOPInstanceUID)
        log.Printf("Receiving image: %s\n", sopInstanceUID)

        // Save to file
        filename := fmt.Sprintf("./received/%s.dcm", sopInstanceUID)
        if err := writer.WriteFile(filename, ds); err != nil {
            log.Printf("Failed to save image: %v\n", err)
            return dimse.NewCStoreResponseFromRequest(req, 0xC000), nil // Failure
        }

        log.Printf("Image saved to %s\n", filename)
        return dimse.NewCStoreResponseFromRequest(req, 0x0000), nil // Success
    })

    // Ensure the receive directory exists
    os.MkdirAll("./received", 0755)

    // Start the server
    log.Println("Starting Storage SCP on port 11112...")
    ctx := context.Background()
    if err := srv.ListenAndServe(ctx); err != nil {
        log.Fatal(err)
    }
}
```

#### Query SCP (query service)

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
    "github.com/cocosip/go-dicom/pkg/network/server"
    "github.com/cocosip/go-dicom/pkg/network/service"
    "github.com/cocosip/go-dicom/pkg/network/status"
)

func main() {
    srv := server.New(
        server.WithAETitle("QUERY_SCP"),
        server.WithPort(11112),
    )

    // Stream C-FIND results; each SendPending call accepts send-queue backpressure.
    srv.SetCFindHandler(func(ctx context.Context, op service.CFindOperation) error {
        query := op.Identifier()
        level := op.QueryLevel()

        log.Printf("Received C-FIND query at level %s\n", level)

        // Query the database (mock data used here)
        results := searchDatabase(query, level)

        for _, result := range results {
            if err := op.SendPending(result); err != nil {
                return err
            }
        }
        return op.SendFinal(status.Success)
    })

    // Start the server
    log.Println("Starting Query SCP on port 11112...")
    ctx := context.Background()
    if err := srv.ListenAndServe(ctx); err != nil {
        log.Fatal(err)
    }
}

func searchDatabase(query *dataset.Dataset, level dimse.QueryRetrieveLevel) []*dataset.Dataset {
    // Mock query results
    result := dataset.New()
    result.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))
    result.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
    result.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
    result.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20240115"}))

    return []*dataset.Dataset{result}
}
```

## DIMSE Messages

DIMSE (DICOM Message Service Element) is the core message layer of DICOM network communication. It provides:

- ✅ C-ECHO: verify connectivity
- ✅ C-STORE: store DICOM instances
- ✅ C-FIND: query the DICOM archive
- ✅ Automatic MessageID management
- ✅ Priority support
- ✅ Status code handling

### Message Structure

All DIMSE messages implement the `dimse.Message` interface:

```go
type Message interface {
    CommandField() uint16
    MessageID() uint16
    SetMessageID(id uint16)
    CommandDataset() *dataset.Dataset
    DataDataset() *dataset.Dataset
    HasData() bool
}
```

Outbound encoding always writes a correct Command Group Length (0000,0000), even
when the supplied command set has no such element (for example a message rebuilt
from scratch): the encoder reserves the element, encodes once, and patches the
length from the actual byte count. The value excludes the group-length element
itself (PS3.7 7.1).

### Automatic MessageID Management

In DICOM network communication, every request needs a unique MessageID. Managing MessageID manually is error-prone:

```go
// ❌ Not recommended: manual MessageID management
msgID := 1
req1 := dimse.NewCEchoRequest()
_ = req1.SetMessageID(uint16(msgID))
msgID++
req2, _ := dimse.NewCStoreRequest(dataset)
_ = req2.SetMessageID(uint16(msgID))
msgID++
// ... easy to forget to increment or to duplicate
```

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

Features:

1. **Thread-safe**: can be called concurrently
2. **Auto-increment**: starts at 1, range 1-65535
3. **Skips 0 automatically**: MessageID=0 means unassigned
4. **Keeps existing values**: does not overwrite an existing MessageID
5. **Supports reset**: the counter can be reset on a new connection

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

### Creating Request Messages

```go
import "github.com/cocosip/go-dicom/pkg/network/dimse"

// C-ECHO
req := dimse.NewCEchoRequest()

// C-STORE
req, err := dimse.NewCStoreRequest(dataset)

// C-FIND
req := dimse.NewCFindRequest(level, query)
req := dimse.NewCFindRequestPatientRoot(level, query)
req := dimse.NewCFindRequestStudyRoot(level, query)

// C-MOVE
req := dimse.NewCMoveRequest(level, destination, identifier)
req := dimse.NewCMoveRequestPatientRoot(level, destination, identifier)
req := dimse.NewCMoveRequestStudyRoot(level, destination, identifier)

// C-GET
req := dimse.NewCGetRequest(level, identifier)
req := dimse.NewCGetRequestPatientRoot(level, identifier)
req := dimse.NewCGetRequestStudyRoot(level, identifier)
```

### Outbound Validation

Every outgoing request passes `dimse.ValidateRequestForSend` before it is
encoded. The validator rejects an empty value for a mandatory identifier:
C-STORE Affected SOP Class/Instance UID; C-FIND, C-GET, and C-MOVE Affected SOP
Class UID; C-MOVE Move Destination; N-GET/N-SET/N-DELETE/N-ACTION requested SOP
Class and Instance UID; N-CREATE Affected SOP Class UID (the instance UID may be
empty so the SCP can assign it); N-EVENT-REPORT Affected SOP Class/Instance UID
and Event Type ID presence. C-ECHO and C-CANCEL always pass. A present but
zero-length element is treated as missing, so no zero-length Type-1 element is
ever written to the wire. The validator covers the built-in request types; a
custom `dimse.Request` implementation is not introspected. Inbound C-STORE-RQ
also rejects a command `AffectedSOPClassUID`/`AffectedSOPInstanceUID` that is
present but empty, while a missing command element still falls back to the Data
Set value for lenient peers.

### Creating Response Messages

```go
import "github.com/cocosip/go-dicom/pkg/network/status"

// Create a response from a request. The status argument is *status.Status.
echoResp := dimse.NewCEchoResponseFromRequest(req, status.Success)
storeResp := dimse.NewCStoreResponseFromRequest(req, status.Success)
findResp := dimse.NewCFindResponseFromRequest(req, status.CFindPending, identifier)
moveResp := dimse.NewCMoveResponseFromRequest(req, status.CMovePending)
getResp := dimse.NewCGetResponseFromRequest(req, status.CGetPending)
```

```go
// Pending response (the constructors take the request's MessageID and SOP Class UID)
resp := dimse.NewCFindResponsePending(req.MessageID(), req.AffectedSOPClassUID(), identifier)
resp := dimse.NewCMoveResponsePending(
    req.MessageID(), req.AffectedSOPClassUID(), remaining, completed, failed, warning)
resp := dimse.NewCGetResponsePending(
    req.MessageID(), req.AffectedSOPClassUID(), remaining, completed, failed, warning)

// Success response
resp := dimse.NewCFindResponseSuccess(req.MessageID(), req.AffectedSOPClassUID())
resp := dimse.NewCMoveResponseSuccess(req.MessageID(), req.AffectedSOPClassUID())
resp := dimse.NewCGetResponseSuccess(req.MessageID(), req.AffectedSOPClassUID())
```

`Affected SOP Class UID` / `Affected SOP Instance UID` are Type U(=) in
C-STORE-RSP and `Affected SOP Class UID` is Type U(=) in C-FIND-RSP. The
constructors therefore omit those command elements when the UID argument is
empty, which allows a conformant minimal response (`Status` +
`MessageIDBeingRespondedTo` only). On the receive side, the service tolerates
their absence instead of aborting the association.

### Checking Response Status

```go
status := resp.Status()

if status.IsSuccess() {
    // Success (0x0000)
} else if status.IsPending() {
    // Pending (0xFF00)
} else if status.IsWarning() {
    // Warning (0x0001, 0x0107, 0x0116, 0xB000-0xBFFF)
} else if status.IsFailure() {
    // Failure (0xA700, 0xA900, 0xC000-0xCFFF)
}

// Get the status code
statusCode := resp.StatusCode()  // uint16
```

A response without a Status (0000,0900) element is malformed. `Status()` reports
it as a generic Failure (0xFFFF) so `IsFailure()` is true; it never returns a
value whose predicates are all false.

### Response Handling

```go
// C-ECHO response
resp := dimse.NewCEchoResponseSuccess(requestMessageID)

if resp.IsSuccess() {
    println("Echo successful")
}
```

```go
// C-STORE response
resp := dimse.NewCStoreResponseSuccess(
    requestMessageID,
    sopClassUID,
    sopInstanceUID,
)

if resp.IsSuccess() {
    println("Store successful")
}
```

```go
// C-FIND response
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

### Priority, Move Originator, and Resetting the Counter

```go
import "github.com/cocosip/go-dicom/pkg/network/dimse"

dimse.PriorityLow     // 0x0002
dimse.PriorityMedium  // 0x0000 (default)
dimse.PriorityHigh    // 0x0001
```

```go
// Priority
req, _ := dimse.NewCStoreRequest(dataset)
req.SetPriority(uint16(dimse.PriorityHigh))    // High priority
req.SetPriority(uint16(dimse.PriorityMedium))  // Medium priority (default)
req.SetPriority(uint16(dimse.PriorityLow))     // Low priority
```

```go
// Move Originator (C-MOVE sub-operations)
req, _ := dimse.NewCStoreRequest(dataset)
_ = req.SetMoveOriginator("MOVE_SCU", 999)
```

```go
// Reset the MessageID counter
idGen := dimse.NewMessageIDGenerator()

// After using it for a while...
idGen.Next()  // 1
idGen.Next()  // 2

// A new connection starts; reset the counter
idGen.Reset()
idGen.Next()  // 1 (starts over)
```

Run the DIMSE tests with:

```bash
go test ./pkg/network/dimse
go test -v ./pkg/network/dimse -run TestMessageIDGenerator
```

## Association Management

### Association structure

```go
type Association struct {
    CalledAE             string
    CallingAE            string
    ApplicationContext   string
    PresentationContexts []*PresentationContext
    MaxPDULength         uint32 // peer's receive limit
    LocalMaxPDULength    uint32 // local receive limit, used when building an RQ
    ImplementationClassUID    string
    ImplementationVersionName string
    LocalImplementationClassUID    string
    LocalImplementationVersionName string
    // ...
}
```

`ToAAssociateRQ` always advertises the local AE's implementation identity and
`LocalMaxPDULength`; it never echoes the peer's values back into an outgoing RQ.
`NewAssociation` initializes the `Local*` fields, so only callers that construct
an `Association` without it need to set them explicitly before converting.

PDU decoding is defensive: decoding into a reused `AAssociateRQ`/`AAssociateAC`
resets the receiver first, Application Context UIDs are trimmed of trailing NUL
or space padding, and User Identity sub-items are validated against PS3.7
Table D.3-14 (identity type 1-5, a single sub-item, a secondary field only for
username/password, no trailing bytes). A username/password sub-item with an
empty passcode is accepted: the standard only forbids a non-empty secondary
field for the other types. The encoder applies the same rules before writing an
A-ASSOCIATE-RQ, and both `FromAAssociateRQ` and `ToAAssociateRQ` copy the
identity slices and drop an invalid identity, so a hand-built PDU or a locally
constructed identity cannot bypass the check.

### Presentation Context

```go
type PresentationContext struct {
    ID             byte
    AbstractSyntax string
    TransferSyntax string
    Result         byte
}

// Find a Presentation Context
pc := assoc.FindPresentationContextByID(1)
pc := assoc.FindPresentationContextByAbstractSyntax("1.2.840.10008.1.1")

// Get the accepted Contexts
accepted := assoc.GetAcceptedPresentationContexts()

// Get the Transfer Syntax
ts := assoc.GetTransferSyntaxForAbstractSyntax("1.2.840.10008.1.1")
```

## Advanced Features

### TLS Encrypted Connections

#### Client using TLS

```go
import "crypto/tls"

// The high-level client completes the TLS handshake and DICOM Association negotiation
c := client.New(
    client.WithCallingAE("SECURE_SCU"),
    client.WithCalledAE("SECURE_SCP"),
    client.WithTLSConfig(&tls.Config{
        ServerName: "pacs.hospital.com",
        MinVersion: tls.VersionTLS12,
    }),
)
if err := c.AddPresentationContext(
    "1.2.840.10008.1.1",
    "1.2.840.10008.1.2.1",
); err != nil {
    log.Fatal(err)
}
err := c.Connect(ctx, "pacs.hospital.com", 11112)
```

`ServerName` may be omitted: when the TLS configuration has no explicit
`ServerName` and `InsecureSkipVerify` is false, the transport derives it from the
dial address host (an IP literal is used verbatim and checked against the
certificate's IP SANs). An explicitly set `ServerName` always wins, and it is
never derived when verification is disabled.

#### Server using TLS

```go
import "crypto/tls"

// Load certificates
cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
if err != nil {
    log.Fatal(err)
}

tlsConfig := &tls.Config{
    Certificates: []tls.Certificate{cert},
}

// Create the TLS server
srv := server.New(
    server.WithAETitle("SECURE_SCP"),
    server.WithPort(11112),
    server.WithTLSConfig(tlsConfig),
)
```

### Association Negotiation Control

```go
// Server side: custom Association negotiation
srv.SetAssociationNegotiatorFunc(func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
    // Check the Calling AE Title
    if assoc.CallingAE != "AUTHORIZED_SCU" {
        return responder.SendReject(ctx,
            pdu.ResultRejectedPermanent,
            pdu.SourceServiceUser,
            pdu.ReasonServiceUserCalledAETitleNotRecognized)
    }

    // Accept only specific SOP Classes
    acceptedSOPs := map[string]bool{
        "1.2.840.10008.1.1": true,          // Verification
        "1.2.840.10008.5.1.4.1.1.2": true,  // CT Image Storage
    }

    for _, pc := range assoc.PresentationContexts {
        if acceptedSOPs[pc.AbstractSyntax] && len(pc.ProposedTransferSyntaxes) > 0 {
            pc.Accept(pc.ProposedTransferSyntaxes[0])
        } else {
            pc.Reject(pdu.ResultAbstractSyntaxNotSupported)
        }
    }

    return responder.SendAccept(ctx, assoc)
})
```

### Observability

```go
// Check whether the server is running
if srv.IsRunning() {
    // The server is running
}

// Get the number of active connections
activeConns := srv.ActiveConnections()
log.Printf("Active connections: %d\n", activeConns)
```

### Timeout Control

```go
// Configure client timeouts
c := client.New(
    client.WithConnectTimeout(10*time.Second),      // TCP connection timeout
    client.WithAssociationTimeout(10*time.Second),  // Association negotiation timeout
    client.WithRequestTimeout(30*time.Second),      // DIMSE request timeout
)

// Use context to control the timeout of a single operation
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := c.CEcho(ctx); err != nil {
    // Handle the timeout error
}
```

```go
// Set a timeout for a single operation
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

err := client.CEcho(ctx)

// Context cancellation
ctx, cancel := context.WithCancel(context.Background())
go func() {
    time.Sleep(5 * time.Second)
    cancel()  // cancel the operation
}()
```

### Concurrency and Timeout Contract

The same `Client` can concurrently call `Connect`, `Close`, `IsConnected`, and
`GetAssociation`; only one connection attempt is allowed at a time. DIMSE requests
can run concurrently within the negotiated Asynchronous Operations Window.
Configuration and Presentation Contexts can only be added before the first
`Connect`; once the connection lifecycle has started, a status error is returned.

`WithRequestTimeout` is a single DIMSE response idle timeout and does not
disconnect because the association is idle. `WithTransportReadTimeout` and
`WithTransportWriteTimeout` each constrain only a single underlying PDU I/O.
When a C-FIND, C-MOVE, or C-GET request times out (`RequestTimeoutError`) or its
context is cancelled, the service sends exactly one best-effort C-CANCEL (bounded
by one second, and deduplicated when the timer and the cancellation race) so the
peer stops producing responses this side would discard; a C-ECHO
or N-Service request times out without C-CANCEL because the operation does not
support it. The terminal error is delivered after the pending request and
asynchronous-operation slot are released, and the returned event stream must be
drained until it closes.

Graceful shutdown (`Client.Close`, `Service.GracefulRelease`) is bounded by the
caller's context: the context covers both the A-RELEASE-RP wait and the drain of
in-flight request handlers. If it expires, the service closes the connection and
returns a timeout error instead of waiting for the default 60-second handler
shutdown budget. `Server.Shutdown(ctx)` likewise honors its context, while
`Server.Close` terminates connections immediately.

Client and Server use the same timeout boundaries: `WithAssociationTimeout`
constrains only association negotiation; the Server's `WithAcceptTimeout`
constrains a single accept attempt; the Server's `WithRequestTimeout` constrains
only DIMSE requests actively sent on that connection. The execution time of
inbound handlers is controlled by the context the handler receives and by
application logic.

### Concurrent Processing

```go
// The server handles concurrent connections automatically
srv := server.New(
    server.WithMaxConnections(100), // Maximum number of concurrent connections
)

// The client can send multiple requests concurrently (using different connections)
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()

        c := client.New(/* ... */)
        if err := c.Connect(ctx, host, port); err != nil {
            log.Printf("Client %d: connect failed: %v\n", id, err)
            return
        }
        defer c.Close()

        if err := c.CEcho(ctx); err != nil {
            log.Printf("Client %d: echo failed: %v\n", id, err)
        }
    }(i)
}
wg.Wait()
```

## Status Codes / Error Handling

### Status Code Reference

Common DICOM status codes:

```go
import "github.com/cocosip/go-dicom/pkg/network/status"

// Success
status.Success                    // 0x0000

// Pending (query/retrieve operations)
status.Pending                    // 0xFF00

// Warning
status.AttributeListError         // 0x0107
status.AttributeValueOutOfRange   // 0x0116

// Failure
status.ErrorProcessingFailure           // 0x0110
status.ErrorUnrecognizedOperation     // 0x0211
status.CStoreRefusedOutOfResources     // 0xA700
status.ErrorDataSetDoesNotMatchSOPClass // 0xA900
status.CStoreErrorCannotUnderstand     // 0xC000
```

Check status:

```go
resp, err := c.SendCEcho(ctx, req)
if err != nil {
    return err
}

if resp.Status().IsSuccess() {
    // Success
} else if resp.Status().IsPending() {
    // Waiting for more results
} else if resp.Status().IsFailure() {
    // Failure
}
```

### Common Error Types

```go
// Connection error
if err := c.Connect(ctx, host, port); err != nil {
    if strings.Contains(err.Error(), "connection refused") {
        // server not running or wrong port
    } else if strings.Contains(err.Error(), "timeout") {
        // connection timeout
    } else if strings.Contains(err.Error(), "no presentation contexts") {
        // no Presentation Context added
    }
}

// Association rejected
if err := c.Connect(ctx, host, port); err != nil {
    if strings.Contains(err.Error(), "association rejected") {
        // Association was rejected
        // Possible reasons: AE Title mismatch, unsupported SOP Class
    }
}

// DIMSE error
if err := c.CEcho(ctx); err != nil {
    if strings.Contains(err.Error(), "not connected") {
        // not connected
    } else if strings.Contains(err.Error(), "context deadline exceeded") {
        // timeout
    }
}
```

### Status Code Errors

```go
resp, err := c.SendCEcho(ctx, req)
if err != nil {
    return err
}

if !resp.Status().IsSuccess() {
    switch resp.StatusCode() {
    case 0xA700:
        // Out of Resources
    case 0xA900:
        // Data Set does not match SOP Class
    case 0xC000:
        // Cannot understand
    default:
        return fmt.Errorf("operation failed: status 0x%04X", resp.StatusCode())
    }
}
```

### Error Wrapping

```go
import "errors"

if err != nil {
    // Check for specific errors
    if errors.Is(err, context.DeadlineExceeded) {
        // timeout error
    }

    if errors.Is(err, context.Canceled) {
        // Context was canceled
    }

    // Wrap the error
    return fmt.Errorf("failed to store image: %w", err)
}
```

## Performance Tips

### 1. Reuse client connections

```go
// ✅ Good practice: reuse the connection
c := client.New(/* ... */)
c.Connect(ctx, host, port)
defer c.Close()

for i := 0; i < 100; i++ {
    c.CEcho(ctx)  // Reuse the same Association
}

// ❌ Bad practice: reconnect every time
for i := 0; i < 100; i++ {
    c := client.New(/* ... */)
    c.Connect(ctx, host, port)
    c.CEcho(ctx)
    c.Close()  // Renegotiates the Association every time — expensive
}
```

### 2. Tune MaxPDULength

```go
// Advertise a larger receive limit for large file transfers. MaxPDULength is a
// per-direction receive limit (PS3.7 D.3.3.1); the peer may advertise a
// different value. Outgoing PDUs use the smaller of the two limits.
c := client.New(
    client.WithMaxPDULength(65536), // Default 16384
)
```

### 3. Server resource limits

```go
srv := server.New(
    server.WithMaxConnections(50),                  // Limit concurrent connections
    server.WithAssociationTimeout(30*time.Second),  // Limit Association negotiation
    server.WithTransportReadTimeout(2*time.Minute), // Limit a single PDU read / connection idle
)
```

Connections that arrive while `WithMaxConnections` is exhausted are still
accepted: the server reads their A-ASSOCIATE-RQ and answers with
A-ASSOCIATE-RJ (result 2 rejected-transient, source 3
service-provider-presentation, reason 2 local-limit-exceeded) instead of leaving
the peer hanging, and they are never counted as active. Rejection is still
resource-bounded: at most `WithMaxConnections` reject negotiations run at once,
further connections are closed without reading their request, and each
rejection wait is capped at 30 seconds (or `WithAssociationTimeout` when that is
smaller), so a connection flood cannot spawn unbounded goroutines. A rejected
connection also emits its matching connection-closed event. Graceful
`Shutdown` closes connections that are still negotiating (including rejects)
before waiting, so it never waits for that negotiation bound. An expired
`WithAcceptTimeout` is treated as an idle poll and produces no `accept_failed`
error record; other persistent accept errors back off from 5 ms up to 1 second.

## Next Steps

- See the [Troubleshooting Guide](./TROUBLESHOOTING.md) for common issues
- Run the example programs in the `examples/` directory

## References

- [DICOM Standard - Part 7: Message Exchange](https://dicom.nema.org/medical/dicom/current/output/html/part07.html)
- [DICOM Standard - Part 8: Network Communication Support](https://dicom.nema.org/medical/dicom/current/output/html/part08.html)
- [fo-dicom project](https://github.com/fo-dicom/fo-dicom) - C# reference implementation
