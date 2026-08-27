# Release Notes

## Next

- `Client.AddPresentationContext` and `Client.AddPresentationContextWithRoles`
  now return an error. They reject invalid UIDs, invalid context IDs, more than
  128 contexts, and calls made after connection lifecycle start.
- `Client.RequestTimeout` now controls an individual DIMSE response idle
  timeout. It no longer closes an otherwise idle association; use
  `WithTransportReadTimeout` for a per-PDU transport deadline.
- Server `AcceptTimeout` and `RequestTimeout` now reach the listener and
  per-connection service runtime. `AssociationTimeout` is negotiation-only;
  use the new server `WithTransportReadTimeout` and
  `WithTransportWriteTimeout` options for per-PDU deadlines.
- `Service.SendCFindWithError` exposes asynchronous C-FIND terminal errors
  while the existing `SendCFind` response-channel API remains available.
