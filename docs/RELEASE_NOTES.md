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
- `Service.SendCFindWithError`, `SendCMoveWithError`, and
  `SendCGetWithError` expose asynchronous progressive-request terminal errors
  while the existing response-channel APIs remain available.
- Progressive C-FIND, C-MOVE, and C-GET callbacks now send C-CANCEL and drain
  the final response when the callback stops. Caller context cancellation makes
  one bounded best-effort C-CANCEL attempt and preserves the context error.
- Incoming associations without DICOM UL Protocol Version 1 bit 0 are rejected
  with the standard ACSE protocol-version-not-supported diagnostic.
- Server lifecycle publication is atomic, and graceful shutdown now waits for
  the accept loop to finish before waiting for active connections.
- Profile-based anonymization uses structure-preserving composite actions,
  VR-valid dummy values, recursive Sequence processing, and top-level
  de-identification declarations. `RetainUIDs` retains only UIDs; other
  exceptions such as Patient ID require an explicit `ActionK` override.
- `Anonymizer.AnonymizeFileInPlace` returns fresh File Meta Information so
  source AE titles, private metadata, and implementation identifiers are not
  reused in anonymized files.
