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
