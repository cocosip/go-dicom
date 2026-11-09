# DICOM Networking

Short package overview for go-dicom's networking layer.

- Full usage guide: [docs/network/README.md](../../docs/network/README.md)
- Troubleshooting: [docs/network/TROUBLESHOOTING.md](../../docs/network/TROUBLESHOOTING.md)
- DIMSE message model: [pkg/network/dimse/README.md](dimse/README.md)

## Overview

go-dicom implements DICOM DIMSE networking: an SCU client, an SCP server, DIMSE
message models, association negotiation (including SCU/SCP role selection), and
observability. Supported operations: C-ECHO, C-STORE, C-FIND, C-MOVE, C-GET, and
all six N-Service operations.

The client is available in two styles:

- `client.Client` — session style: `Connect` first, then call `CEcho`, `CStore`,
  `CFind`, `CMove`, `CGet`, or the N-Service methods.
- `client.ManagedClient` — queue style (analogous to fo-dicom `DicomClient`):
  enqueue Jobs with `Add`, then `Send` batches them over automatically configured
  associations.

For C-GET (and C-MOVE delivery back to this AE), Storage SOP Classes must be
negotiated with the SCP role and a `CStoreHandler` must be configured; see the
full guide for `AddStoragePresentationContexts`, `AddStoragePresentationContextsByCategory`,
and the `PresentationContextSpec` factories.

## Behavior guarantees

Recent conformance hardening is documented in the full guide; the short version:

- Outgoing requests are validated before encoding (`dimse.ValidateRequestForSend`);
  present-but-empty mandatory UIDs are rejected, never written as zero-length
  Type-1 elements.
- C-STORE-RSP/C-FIND-RSP optional UIDs (Type U(=)) may be omitted by peers and are
  omitted by the response constructors when empty.
- Unrequested role selections in an AC are ignored and over-granted roles are
  clamped (PS3.7 D.3.3.4) instead of closing the association.
- Handler panics are recovered per request and abort only that association; the
  process keeps serving other associations.
- `Close`/`GracefulRelease` are bounded by the caller's context; a peer A-ABORT
  leaves the association in `StateAborted`.
- Cancelling a progressive C-FIND/C-MOVE/C-GET (or its request timeout) sends one
  best-effort C-CANCEL; drain the event stream until it closes.
- `ManagedClient` re-establishes a dead association for Jobs not yet attempted;
  Jobs that already returned an error are not retried.
- Over-limit connections get A-ASSOCIATE-RJ(local-limit-exceeded); reject
  negotiations are capped at MaxConnections and shut down with the server. Idle
  accept timeouts are silent, and persistent accept errors back off.
- `DialTLS` derives `ServerName` from the dial address when it is unset.

## Package layout

| Path | Purpose |
|------|---------|
| `association` | Association state and presentation-context / role negotiation |
| `client` | SCU client (`client.Client` session style, `client.ManagedClient` queue style) |
| `dimse` | DIMSE request/response messages |
| `server` | SCP server and handler registration |
| `service` | Association-scoped DIMSE engine shared by client and server |
| `pdu` | DUL PDU encoding/decoding |
| `transport` | TCP/TLS transport |
| `observability` | Lifecycle events, structured logs, and metrics |
| `status` | DIMSE status codes |

## Quick start

```go
c := client.New(
    client.WithCallingAE("MY_SCU"),
    client.WithCalledAE("PACS"),
)
_ = c.AddPresentationContext("1.2.840.10008.1.1", "1.2.840.10008.1.2.1")
if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
    return err
}
defer c.Close()
if err := c.CEcho(ctx); err != nil {
    return err
}
```

See the [networking guide](../../docs/network/README.md) for client/server
examples, configuration, role negotiation, DIMSE messages, status codes, and
performance tips.
