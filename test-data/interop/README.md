# Interoperability fixtures

These binaries are the reference data for the interoperability tests. They are
the only artifacts kept from the fo-dicom interop runs; no C# code is stored in
this repository and no interop test requires .NET at build or test time.

## Files

| File | Produced by | Purpose |
|---|---|---|
| `ct_small_explicit_le.dcm` | go-dicom `writer.Write` with Explicit VR Little Endian, defined-length sequences and items, Implementation Class UID `1.2.826.0.1.3680043.10.999.1`, Implementation Version Name `GODICOM-INTEROP` | Byte-level golden for `pkg/dicom/writer` (`TestWriteMatchesInteropGolden`) |
| `ct_small_explicit_le_fo-dicom.dcm` | fo-dicom v6.0.0-alpha1 `DicomFile.Save` of the fixture above (`ExplicitLengthSequences=true`), which re-stamps the Implementation Class UID/Version Name and recalculates the file-meta group length | Input for `TestParseFoDicomEncoding`: go-dicom must read what fo-dicom writes |

Both carry the same dataset: one CT instance with padded string values
(`PN`, `LO`, `UI`), nine numeric values, a defined-length referenced-image
sequence item, and 8 bytes of 16-bit pixel data.

## Verified interoperability (2026-11-17)

fo-dicom v6.0.0-alpha1 (checkout of `dotnet-source/fo-dicom`, .NET 10):

- **go-dicom → fo-dicom**: `DicomFile.Open` reads `ct_small_explicit_le.dcm`,
  reports Explicit VR Little Endian, 22 top-level items, and the expected
  values, including the 8-byte pixel data (SHA-256
  `66840dda154e8a113c31dd0ad32f7f3a366a80e8136979d8f5a101d3d29d6f72`).
- **go-dicom vs fo-dicom encoding**: re-encoding the fixture with
  `ExplicitLengthSequences=true` produces a byte-identical *dataset body*
  (494 bytes, SHA-256
  `37a46c87045f8bd2e5ee6b307e7f4ecae38303f774601ceb68d9eb045768e21f`) for
  both implementations. Only the file-meta group differs, because each
  implementation stamps its own Implementation Class UID and Version Name.
- **fo-dicom → go-dicom**: the re-encoded file is kept as
  `ct_small_explicit_le_fo-dicom.dcm` and is parsed by go-dicom with the same
  element values and pixel bytes.

## Regenerating `ct_small_explicit_le.dcm`

The dataset is defined once, in `pkg/dicom/writer/interop_test.go`:

```bash
GO_DICOM_UPDATE_INTEROP_GOLDEN=1 go test ./pkg/dicom/writer -run TestWriteMatchesInteropGolden
```

Only regenerate after re-running the interop check above with fo-dicom: the
golden exists to catch accidental encoding drift (element order, padding,
sequence lengths, file meta), so an unexplained byte change is a regression, not
a fixture update.
