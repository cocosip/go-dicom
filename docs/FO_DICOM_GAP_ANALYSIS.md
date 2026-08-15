# fo-dicom Capability Gap Analysis

[简体中文](FO_DICOM_GAP_ANALYSIS.zh-CN.md)

This document tracks capability gaps between `go-dicom` and the reference
`fo-dicom` implementation. It is a versioned engineering backlog, not a claim
that every .NET API must be reproduced in Go.

## Audit Baseline

The initial audit was performed on 2026-08-14 against these revisions:

- `go-dicom`: `d5970f342973c0d659c3eab1b7cee8563a7f5dda`
- `fo-dicom`: `7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2`
  (`5.2.6-101-g7ea6d424`)

Source code, tests, and examples are the evidence for this comparison. README
feature lists are not considered implementation evidence.

The following command passed at the audited `go-dicom` revision:

```powershell
go test ./cmd/... ./examples/... ./pkg/... ./tools/...
```

Recheck both repositories before implementing an item because either side may
have changed since this baseline.

## Classification

| Status | Meaning |
| --- | --- |
| `Open` | No equivalent domain capability exists in `go-dicom`. |
| `Partial` | Some layers or value types exist, but the public workflow is incomplete. |
| `External` | The capability is intentionally supplied by a companion module. |
| `Not a gap` | The difference is platform-specific or absent from both libraries. |
| `Complete` | Acceptance criteria were met and current evidence is recorded. |

Priority indicates implementation order, not estimated effort:

- `P0`: correctness, interoperability, or misleading public contract
- `P1`: major fo-dicom domain workflow required for practical parity
- `P2`: important supporting API or operational capability
- `P3`: large or specialized capability that should follow its prerequisites

## Executive Summary

| ID | Priority | Status | Capability |
| --- | --- | --- | --- |
| DOC-001 | P0 | Partial | Public capability statements match implemented behavior |
| NET-001 | P0 | Complete | TLS-enabled high-level DICOM client |
| NET-002 | P0 | Complete | C-STORE transfer syntax selection and automatic transcoding |
| STD-001 | P0 | Complete | Reproducible and current generated standard tables |
| MED-001 | P1 | Complete | DICOMDIR media directory model and read/write workflow |
| NET-003 | P1 | Complete | Advanced association negotiation through the high-level client |
| NET-004 | P1 | Complete | SOP Class Common Extended Negotiation |
| SR-001 | P1 | Partial | Complete Structured Report value types and file workflow |
| IMG-001 | P1 | Partial | Dataset-driven image rendering pipeline |
| CORE-001 | P1 | Complete | Recursive Dataset and Sequence validation |
| IMG-002 | P2 | Open | Frame geometry, spatial transforms, and interpolation tools |
| CORE-002 | P2 | Open | Dataset walker, match rules, and transform rules |
| DICT-001 | P2 | Complete | Runtime XML dictionary loading |
| ANON-001 | P2 | Complete | Complete custom anonymization profile loading |
| PRINT-001 | P2 | Partial | Dataset-backed DICOM Print Management models |
| OBS-001 | P2 | Open | Structured network logging, request events, and metrics hooks |
| IMG-003 | P3 | Partial | Volume reconstruction and MPR |
| MED-002 | P3 | Open | DICOM file scanner workflow |

## Implementation Progress

Progress as of 2026-08-15:

- **Complete:** NET-001, NET-002, STD-001, MED-001, NET-003, NET-004,
  CORE-001, ANON-001, and DICT-001.
- **Not complete:** DOC-001, SR-001, IMG-001, IMG-002, CORE-002, PRINT-001,
  OBS-001, IMG-003, and MED-002 retain their `Partial` or `Open` status.
- Phase 0 is not complete. NET-001, NET-002, and STD-001 are complete; the
  remaining Phase 0 work is tracked by DOC-001.
- **Next item:** SR-001, the first incomplete item in the planned development
  order below.

## Planned Development Order

Priority describes capability importance; this order describes the intended
implementation sequence. Sequence numbers remain stable as work is completed.
Mark completed rows `Complete`; the first row that is not complete is the next
item to implement. Reorder rows only when new dependency or scope evidence is
recorded in this document.

| Order | ID | Priority | Current status | Ordering rationale |
| ---: | --- | --- | --- | --- |
| 1 | NET-004 | P1 | Complete | Complete the NET-003 negotiation family while its association APIs and tests are current. |
| 2 | CORE-001 | P1 | Complete | Establish recursive Dataset and Sequence correctness before completing nested domain workflows. |
| 3 | SR-001 | P1 | Partial | Build typed SR trees and file workflows on the recursive validation behavior from CORE-001. |
| 4 | IMG-001 | P1 | Partial | Establish the Dataset-driven rendering pipeline before adding shared spatial tooling. |
| 5 | IMG-002 | P2 | Open | Add geometry, transforms, and interpolation after rendering requirements are stable; this is also an IMG-003 prerequisite. |
| 6 | CORE-002 | P2 | Open | Generalize walking, path, match, and transform APIs after CORE-001 and SR-001 establish their traversal semantics. |
| 7 | PRINT-001 | P2 | Partial | Complete Dataset-backed print models and the N-service workflow after the core Dataset work is stable. |
| 8 | OBS-001 | P2 | Open | Add cross-cutting network diagnostics after negotiation and print network workflows have settled. |
| 9 | MED-002 | P3 | Open | Deliver the independent, bounded scanner workflow before the larger reconstruction item. |
| 10 | IMG-003 | P3 | Partial | Implement volume reconstruction and MPR only after IMG-001 and IMG-002 are complete. |
| 11 | DOC-001 | P0 | Partial | Perform the final public API and README audit after the major capabilities are complete, avoiding repeated documentation churn. |

## Detailed Gaps

### DOC-001: Public Capability Statements

**Status:** `Partial`  
**Priority:** `P0`

At the audit baseline, several README statements were broader than the
implemented public workflow. Examples included complete client TLS, complete
SR value types, image reconstruction, advanced negotiation through the client,
and print job creation. The TLS example called a nonexistent `client.WithTLS`
option, and one rendering example passed a Dataset to `NewDicomImage` even
though the function accepts `*DicomPixelData`.

At the audit baseline, the README's tag and UID totals did not match the unique
generated standard entries counted in the audited source.

Progress on 2026-08-14: NET-001 repaired the high-level client TLS API and its
README examples. NET-002 aligned C-STORE transfer syntax selection and
transcoding claims with the implemented send path. STD-001 corrected the
generated Tag and UID totals and the tooling description. The remaining SR,
rendering, advanced negotiation, print, and other public capability statements
have not been re-audited or repaired, so DOC-001 remains `Partial`.

**Acceptance criteria**

- Every checked capability is backed by a compiling public API and focused test.
- Partial and external capabilities are labeled explicitly.
- All README code examples compile in CI.
- Generated tag and UID totals are derived automatically or omitted.

**Suggested verification**

- Add compile tests for README snippets or move runnable snippets to examples.
- Run the full Go package tree and documentation link checks.

### NET-001: High-Level Client TLS

**Status:** `Complete`

**Priority:** `P0`

Completed on 2026-08-14. The high-level client now exposes
`client.WithTLSConfig(*tls.Config)`. A nil configuration preserves the existing
plain-TCP path, while a non-nil configuration uses `transport.DialTLS` before
DICOM association negotiation. `ConnectTimeout` covers both TCP establishment
and the TLS handshake, caller cancellation is preserved, and the transport
layer clones the caller-owned TLS configuration before applying defaults.

Reference: [fo-dicom DicomClient](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Network/Client/DicomClient.cs)

**Acceptance criteria**

- The client accepts a caller-owned `*tls.Config` without mutating it.
- Plain TCP remains the default.
- TLS handshake, certificate validation, timeout, cancellation, and close paths
  are covered.
- The README TLS example compiles and uses the actual public option name.

**Verification evidence**

- Focused tests cover plain TCP, a verified TLS handshake, a complete TLS
  association, hostname mismatch, untrusted certificates, handshake timeout,
  caller cancellation, and concurrent clients sharing one `tls.Config`.
- `go test ./pkg/network/... -count=1` passed.
- `go test -race ./pkg/network/... -count=1` passed.
- `go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1` passed.
- `golangci-lint run` reported 0 issues; the sandbox emitted cache-persistence
  warnings after analysis completed.
- The implementation and this verification evidence were committed in
  `e8e44ef`.

### NET-002: C-STORE Negotiated Transfer Syntax and Transcoding

**Status:** `Complete`
**Priority:** `P0`

Completed on 2026-08-14. The shared DIMSE send path now selects C-STORE
presentation contexts by SOP Class and source transfer syntax. An exact source
syntax match is preferred. Otherwise, the first accepted syntax that the codec
registry can transcode to is selected deterministically and a transcoded copy
is sent. Explicit presentation context IDs remain authoritative and are
validated before use.

Datasets containing Pixel Data must identify their source transfer syntax;
parsed and DIMSE-received Datasets do so automatically, which also supports
receive-and-forward workflows. Missing source syntax and unavailable codec
paths fail before any network write with SOP Class and transfer syntax context.
Datasets without Pixel Data can be re-encoded directly and do not require image
codecs, while explicit context IDs are still validated. Caller-owned Dataset
and codec input frames are isolated from transcoding mutations.

`CStore`, `CStoreWithPriority`, and C-GET C-STORE sub-operations all use the
shared service path. `CStoreMultiple` remains sequential, stops at the first
failure or cancellation, and returns the number of completed successful stores.

Reference: [fo-dicom DicomCStoreRequest](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Network/DicomCStoreRequest.cs)

**Acceptance criteria**

- Presentation context selection considers both SOP Class and source transfer
  syntax.
- The original syntax is preferred when accepted.
- A registered codec is used for fallback transcoding when necessary.
- Failure is explicit when no accepted syntax is directly usable or
  transcodable.
- The caller's Dataset is not mutated.
- Batch sending defines ordering, partial-success, cancellation, and concurrency
  semantics.

**Verification evidence**

- Service tests inspect the actual PDV presentation context and decode the wire
  Dataset to verify exact-syntax selection and Little/Big Endian transcoding.
- Tests cover registered compressed-codec fallback, unavailable codecs, unknown
  source syntax, explicit context IDs (including rejected and wrong-SOP
  contexts), non-pixel objects, receive-and-forward behavior, and caller
  Dataset immutability against mutating codecs.
- Client tests define sequential batch ordering, partial-success count, first
  failure, and cancellation behavior.
- `CGO_ENABLED=0 go test ./pkg/imaging/... ./pkg/network/... -count=1` passed.
- `CGO_ENABLED=0 go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`
  passed.
- `CGO_ENABLED=0 golangci-lint run` reported 0 issues. The implementation adds
  no CGo directives, C imports, native dependencies, or module changes.
- Cross-process interoperability with the real companion RLE, JPEG, JPEG-LS,
  and JPEG 2000 codecs remains a separate verification level.

### STD-001: Generated Standard Tables and Tooling

**Status:** `Complete`
**Priority:** `P0`

Completed on 2026-08-14. The previous generated sources contained 5,338 tags,
5,338 standard dictionary entries, and 1,906 standard UIDs. Three separate
one-off tools also read hard-coded files from an absent repository-local
`fo-dicom-code` directory.

The fo-dicom 2026b `DICOM Dictionary.xml` and manually maintained
`Private Dictionary.xml` are now pinned under `tools/data/2026b`. One
`tools/generate_dicom` command accepts both input paths and the repository root,
then regenerates all four XML-derived outputs together:

- 5,347 standard Tag constants
- 1,928 standard UID constants
- 5,347 standard dictionary entries
- 235 private creators containing 4,678 private entries

The generator preserves fo-dicom's `RETIRED` suffix convention for generated
Tag and UID identifiers while retaining the original XML keyword in dictionary
entries. The 59 private UIDs in `pkg/dicom/uid/uids_private.go` have no source
in either authoritative XML file and are intentionally not regenerated by this
tool. Updating the baseline is an explicit copy-and-regenerate operation; no
fo-dicom source downloader is included.

**Acceptance criteria**

- Generator inputs are explicit CLI arguments, a documented checked-out source,
  or committed standards inputs with acceptable licensing.
- A clean checkout can reproduce all generated files.
- Generation is deterministic and checked by CI.
- Tag, dictionary, and UID sources are updated atomically from one baseline.

**Verification evidence**

- SHA-256 hashes of both committed XML files match the local fo-dicom 2026b
  source files byte for byte.
- Independent XML parsing confirms 5,347 tags, 1,928 UIDs, 235 private
  creators, and 4,678 private entries.
- A generator integration test regenerates all four outputs from the bundled
  XML in a temporary directory, asserts the exact counts, and compares each
  file with the committed output after normalizing platform line endings. It
  runs through the existing CI `./tools/...` test scope.
- Focused tests cover the fo-dicom retired-identifier convention and current
  `vm` symbol mappings.

### MED-001: DICOMDIR

**Status:** `Complete`
**Priority:** `P1`

Completed on 2026-08-15. Package `pkg/media` now provides DICOMDIR creation,
strict and compatible reading, hierarchical traversal, deterministic file
grouping, bounded offset recovery, and two-pass writing. It supports PATIENT,
STUDY, SERIES, IMAGE, SR DOCUMENT, and PRESENTATION records and preserves
unknown record types. Referenced files are not scanned, copied, moved, renamed,
or rewritten.

Optional icon generation is implemented by `pkg/imaging` through a structural
interface, so `pkg/media` and `pkg/imaging` do not import each other. Icons use
the existing pure-Go codec registry, render a representative frame, preserve
aspect ratio, and produce 8-bit MONOCHROME2 images no larger than 128x128.

Reference: [fo-dicom Media](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Media)

**Acceptance criteria**

- Open and save DICOMDIR files while preserving valid record offsets.
- Expose patient, study, series, and instance record traversal.
- Add files with valid Referenced File IDs and deterministic grouping.
- Support malformed or stale offsets with documented strict/compatible behavior.
- Make icon generation optional and independent from the directory core.

**Verification evidence**

- Focused tests cover File ID validation, duplicate and anonymized grouping,
  missing attributes, exact offsets, fixed-delta and type/order recovery,
  ambiguity, cycles, duplicate references, unreachable records, writer errors,
  source immutability, representative frames, and icon failures.
- The 13,796-byte fo-dicom DICOMDIR fixture opens in strict mode with 80 records
  and round-trips through both supported directory transfer syntaxes.
- fo-dicom 6.0.0-alpha1 opened Go-generated Explicit and Implicit VR Little
  Endian DICOMDIR files with identical 9-record hierarchy, type counts, and
  Referenced File IDs.
- `CGO_ENABLED=0 go test ./pkg/media ./pkg/imaging/... -count=1` passed. No CGo
  directives, C imports, or native dependencies were added.
- `CGO_ENABLED=0 go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`
  passed.
- `CGO_ENABLED=0 golangci-lint run` reported 0 issues.

### NET-003: Advanced Association Negotiation

**Status:** `Complete`
**Priority:** `P1`

The high-level client now exposes Asynchronous Operations Window, SCP/SCU Role
Selection, SOP Class Extended Negotiation, and User Identity options. User
Identity supports Username, Username/Password, Kerberos, SAML, and JWT. The
established association retains both requested and accepted values.

Positive User Identity responses are required by default when requested; a
compatibility option can explicitly allow an omitted response. The negotiated
maximum invoked operations limits unfinished DIMSE requests (`0` means
unlimited), while maximum performed operations remains negotiated metadata,
matching fo-dicom behavior. Role Selection is bound to Presentation Contexts
and constrains ordinary requests and reverse C-STORE sub-operations. Common
Extended Negotiation item `0x57` is completed separately by NET-004.

**Acceptance criteria**

- High-level client options expose all supported negotiation items.
- Requested and accepted values are available after association establishment.
- Positive User Identity responses and required-response failure behavior are
  defined.
- Async operation limits are enforced by the request dispatcher rather than
  merely encoded on the wire.
- Role Selection affects request and sub-operation behavior where applicable.

**Verification evidence**

- A real client/server association integration test round-trips all NET-003
  negotiation values and verifies the accepted asynchronous request limit.
- Focused tests cover finite and unlimited request windows, cancellation while
  waiting, default and negotiated roles, reverse C-STORE rejection, missing
  positive identity responses, and malformed role responses.
- PDU round-trip tests cover AC asynchronous window, Role Selection, Extended
  Negotiation, and present-empty User Identity responses.
- `CGO_ENABLED=0 go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`
  passed on Go 1.26.6 for Windows/amd64.
- `CGO_ENABLED=0 go build ./...` passed and `CGO_ENABLED=0 golangci-lint run`
  reported 0 issues.

### NET-004: SOP Class Common Extended Negotiation

**Status:** `Complete`
**Priority:** `P1`

Completed on 2026-08-15. The PDU layer now encodes and decodes the complete
`0x57` item, while the association and high-level client layers merge its
Service Class UID and ordered Related General SOP Class UIDs with `0x56`
application information for the same SOP Class. Caller-owned client option
data is copied, and an explicitly invalid common request is preserved until
the PDU encoder rejects it.

Common Extended Negotiation remains request-only: A-ASSOCIATE-AC neither emits
nor accepts `0x57`. Empty Related General SOP Class UID lists are valid;
required SOP Class UID, Service Class UID, and individual related UIDs must be
non-empty. Nested and outer 16-bit lengths are checked before allocation or
write.

Reference: [fo-dicom DicomExtendedNegotiation](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Network/DicomExtendedNegotiation.cs)

**Acceptance criteria**

- Encode and decode the complete `0x57` item with length validation.
- Preserve Service Class UID and all Related General SOP Class UIDs.
- Expose the request through association and high-level client APIs.
- Reject malformed lengths without panics or partial state.

**Verification evidence**

- Exact-byte and round-trip tests cover combined `0x56`/`0x57` values, an empty
  related list, multiple related UIDs, invalid required UIDs, oversized nested
  data, malformed lengths, partial headers, and AC directionality.
- A real Go client/server association round-trips the common values through the
  high-level client and server association APIs.
- Bidirectional full-PDU checks against fo-dicom revision
  `7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2` passed: fo-dicom parsed the Go
  request, and Go parsed the fo-dicom request.
- `go test ./pkg/network/... -count=1`, the full repository test suite,
  `go build ./...`, and `golangci-lint run` passed on Go 1.26.6 Windows/amd64.
- The Windows race gate is environment-blocked: even
  `go test -race fmt -run '^$' -count=1` exits with status `0xc0000139`, before
  running repository test code.

### SR-001: Complete Structured Report Workflow

**Status:** `Partial`  
**Priority:** `P1`

Value type constants exist for the standard SR types, but construction and
typed reading focus on TEXT, CODE, NUM, and CONTAINER. PNAME, DATE, TIME,
DATETIME, UIDREF, COMPOSITE, IMAGE, WAVEFORM, SCOORD, and TCOORD lack complete
typed APIs. SR-specific Open and Save methods remain commented placeholders.

Reference: [fo-dicom StructuredReport](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/StructuredReport)

**Acceptance criteria**

- Every declared Value Type has symmetric constructor and typed reader support.
- Referenced SOP, spatial coordinate, and temporal coordinate constraints are
  validated.
- File and stream open/save workflows preserve SR content trees.
- Root relationship rules and child relationship rules are validated.

**Suggested verification**

- Table-driven round trips for every Value Type.
- Nested content tree and invalid relationship tests.
- Cross-read representative SR documents with fo-dicom.

### IMG-001: Dataset-Driven Rendering

**Status:** `Partial`  
**Priority:** `P1`

The LUT and overlay primitives exist, but `DicomImage` is constructed from
`DicomPixelData` rather than directly from a Dataset or file. Its default
grayscale pipeline fixes rescale slope/intercept to `1/0`, computes an optimal
window from pixels, and does not fully consume Dataset Modality LUT, VOI LUT,
window, presentation, and per-frame metadata. Stored `scale` and
`showOverlays` state do not affect `RenderFrame`.

Reference: [fo-dicom DicomImage](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Imaging/DicomImage.cs)

**Acceptance criteria**

- Construct a renderable image from a Dataset and from a parsed file result.
- Build modality, VOI, presentation, palette, and inversion stages from Dataset
  and per-frame metadata with documented precedence.
- Scaling changes output dimensions and uses a defined interpolation mode.
- Overlay visibility, origin, frame range, and color affect rendered output.
- Preserve explicit caller overrides for window, LUT, scale, and inversion.

**Suggested verification**

- Golden-image tests for CT rescale/windowing, MONOCHROME1, palette color,
  overlays, and multi-frame functional groups.
- Compare representative rendered pixels with fo-dicom within declared
  tolerance.

### CORE-001: Recursive Validation

**Status:** `Complete`
**Priority:** `P1`

Dataset now provides explicit recursive validation and default automatic
validation for insertion APIs. Dataset and Sequence validation use one
fail-fast engine with deterministic tag/item traversal, actual-VR value
validation before dictionary VM validation, fo-dicom VM exceptions, nested
path errors, and retained original causes. Binary, JSON, and XML hydration can
still read invalid values and restore automatic validation before returning.

**Acceptance criteria**

- Dataset validation visits every element and recursively validates sequences.
- Errors identify the nested tag/item path and original validation cause.
- Validation can be explicitly enabled or disabled without global data races.
- VM, VR, required structural fields, and sequence children have defined scope.

**Verification performed**

- Nested sequence, exact path, VM, private/unknown tag, malformed value,
  automatic insertion, mutation rollback, and hydration regression tests pass.
- `go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`,
  `go build ./...`, and `golangci-lint run` pass.
- The affected-package race command cannot start on this Windows host and
  exits with `0xc0000139`; the same failure occurs for `go test -race fmt`, so
  the race runtime remains an environment gate for CI rather than a passing
  local result.

### IMG-002: Geometry and Spatial Image Tools

**Status:** `Open`  
**Priority:** `P2`

fo-dicom includes FrameGeometry, patient/image coordinate conversion,
orientation and localization support, spatial transforms, histogram helpers,
mathematical geometry types, and nearest-neighbor/bilinear interpolation.
go-dicom has no equivalent cohesive geometry layer.

Reference: [fo-dicom FrameGeometry](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Imaging/FrameGeometry.cs)

**Acceptance criteria**

- Parse classic and enhanced multi-frame geometry.
- Convert between pixel and patient coordinates with documented conventions.
- Calculate frame orientation, normals, bounding boxes, and localization lines.
- Provide rotate, flip, translate, scale, and best-fit transforms.
- Provide tested interpolation primitives reusable by rendering and MPR.

**Suggested verification**

- Synthetic axial, sagittal, coronal, and oblique frame coordinate round trips.
- Classic and enhanced multi-frame fixtures with known patient-space geometry.
- Golden tests for transforms and nearest-neighbor/bilinear interpolation.

### CORE-002: Dataset Walker and Rules

**Status:** `Open`  
**Priority:** `P2`

fo-dicom provides a recursive Dataset walker plus composable match and transform
rules. go-dicom consumers must currently build this behavior ad hoc.

Reference: [fo-dicom DicomDatasetWalker](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/DicomDatasetWalker.cs)

**Acceptance criteria**

- Visit elements, sequences, items, and fragments with stable path information.
- Allow early stop and error propagation.
- Supply composable exists/empty/equality/wildcard/regex match rules.
- Supply common remove/set/map/copy/regex/case/UID transform rules.
- Define whether transforms mutate or clone the source Dataset.

**Suggested verification**

- Assert traversal order and paths for nested sequences and pixel fragments.
- Cover early stop, visitor errors, and rule composition precedence.
- Verify mutating and cloning transforms preserve all unrelated elements.

### DICT-001: Runtime XML Dictionary Loading

**Status:** `Complete`
**Priority:** `P2`

Completed on 2026-08-14. `NewFromXML` creates a dictionary from any `io.Reader`,
and `Dictionary.LoadXML` validates an entire document before merging it into an
existing dictionary. Both the single `<dictionary>` standard layout and the
multi-creator `<dictionaries>` private layout used by fo-dicom are supported.

Exact and masked tags, all fo-dicom VR separators, alternative VM values,
keywords, retired flags, UTF-8 BOM input, and known combined-source `<uid>`
elements are handled. Private creator entries are isolated in creator-specific
sub-dictionaries and are selected automatically when a lookup tag carries a
private creator. Exact entries take precedence over masked entries; exact and
masked duplicates use last-loaded-wins semantics. Malformed input reports the
dictionary creator, tag index, and tag value when available, and a failed load
does not partially modify the target dictionary.

Runtime-loaded private dictionaries are also used during implicit VR parsing.
Private Creator reservation elements are decoded as `LO`; subsequent private
elements are resolved by group, allocated block, and creator. Creator
reservations are scoped independently for the root dataset and for every
sequence item, so the same block may safely identify different creators in
different items.

Exact private entries are keyed by their group and low element byte, making
them independent of the block allocated in a dataset. Programmatic `Add` calls
route creator-tagged entries into the matching private sub-dictionary and clone
entries before applying that dictionary's creator, leaving caller-owned entries
unchanged. XML syntax errors inside a `<tag>` retain creator, entry index, and
tag context.

Reference: [fo-dicom DicomDictionaryReader](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/DicomDictionaryReader.cs)

**Acceptance criteria**

- Load standard and private dictionaries from `io.Reader`.
- Support exact and masked tags, multiple VRs, VM, keyword, retired flag, and
  private creator.
- Reject malformed input with element-level context.
- Define duplicate-entry and dictionary merge behavior.

**Verification evidence**

- The complete local fo-dicom 2026b `DICOM Dictionary.xml` and
  `Private Dictionary.xml` files load successfully; a known `MED NM` private
  masked tag resolves through its creator dictionary.
- Tests cover standard and private layouts, exact and masked lookup precedence,
  duplicate replacement, atomic merge failure, every fo-dicom VR separator,
  alternative VM values, BOM input, combined-source UID elements, malformed
  XML with tag context, invalid VR/VM values, missing fields, private exact-tag
  block normalization, creator-mismatch rejection, programmatic private-entry
  routing, caller entry immutability, and concurrent private dictionary
  creation.
- Implicit VR parser tests cover Private Creator reservation decoding, private
  VR lookup, and creator isolation between sequence items that reuse the same
  allocated block.
- `ExampleNewFromXML` is compiled and executed by the Go test suite.
- `go test ./pkg/dicom/dict -count=1` passes.
- `go test -race ./pkg/dicom/dict -count=1` passes.
- `go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1` passes.
- `golangci-lint run` reports 0 issues.
- Implemented in commit `6dc89f7`.

### ANON-001: Custom Profile Loading

**Status:** `Complete`
**Priority:** `P2`

Completed on 2026-08-14. `NewProfileFromReader` now accepts any `io.Reader`,
and `LoadProfileFromFile` opens the requested path and delegates to that same
parser. Existing two-column `pattern;action` profiles remain supported. The
parser also accepts fo-dicom-compatible 12-column profiles and applies all 11
`SecurityProfileOptions` columns with the same precedence as the built-in
profile loader.

Parsing is strict: non-comment input must use one of the two supported column
layouts, patterns and actions are validated, and errors include the source line
number. Blank lines and lines beginning with `#` remain supported. No lenient
mode is currently exposed.

**Acceptance criteria**

- File loading delegates to the same Reader-based parser.
- Reader APIs accept general `io.Reader`, not only `*strings.Reader`.
- Profile option semantics are consistent for built-in and custom inputs.
- Invalid lines and actions return actionable errors instead of being silently
  skipped unless a lenient mode is explicitly selected.

**Verification evidence**

- File and Reader APIs are tested against identical input and rule output.
- Tests cover a general `io.Reader`, every profile option column, combined
  option precedence, comments, malformed column counts, unknown actions, empty
  patterns, and invalid regular expressions.
- `ExampleNewProfileFromReader` is compiled and executed by the Go test suite.
- `go test ./pkg/dicom/anonymizer -count=1` passes.
- `go test -race ./pkg/dicom/anonymizer -count=1` passes.
- `go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1` passes.
- `golangci-lint run` reports 0 issues.
- Implemented in commit `a9e4301`.

### PRINT-001: Dataset-Backed Print Management Models

**Status:** `Partial`  
**Priority:** `P2`

FilmSession, FilmBox, ImageBox, PresentationLUT, and printer status helpers
exist, but the principal models are simplified Go structs rather than
Dataset-backed DICOM objects. They cover a subset of attributes and lack the
fo-dicom clone, load/save, UID lookup/removal, and complete layout behavior.
Empty UIDs use fixed placeholder values rather than generated UIDs.

Reference: [fo-dicom Printing](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Printing)

**Acceptance criteria**

- Models round-trip all supported DICOM attributes through Dataset objects.
- Empty SOP Instance UIDs are generated uniquely.
- Clone, load/save, find, and delete preserve parent/child references.
- Film layout and image box creation support the standard display formats in
  the declared scope.
- N-CREATE/N-SET/N-ACTION integration is demonstrated end to end.

**Suggested verification**

- Dataset round trips for every supported print object and display format.
- Verify UID uniqueness and parent/child reference integrity across clones.
- Run an end-to-end print workflow and cross-read generated objects with
  fo-dicom.

### OBS-001: Network Observability

**Status:** `Open`  
**Priority:** `P2`

fo-dicom exposes structured logging, request sent/pending/completed/timed-out
events, and a network metrics collector hook. go-dicom has lifecycle callbacks
but no cohesive logger or metrics interface; some PDU decoding warnings are
written directly with `fmt.Printf`.

Reference: [fo-dicom Network Metrics](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Log/Metrics)

**Acceptance criteria**

- Library code never writes directly to stdout/stderr.
- Client and server accept no-op-by-default structured logging hooks.
- Request lifecycle events include association, message ID, command, status,
  duration, and timeout/cancellation outcome.
- Metrics hooks expose connection, association, DIMSE, byte, error, and latency
  observations without requiring a specific telemetry vendor.

**Suggested verification**

- Assert the default configuration writes no process output.
- Integration-test hook ordering for success, pending, timeout, cancellation,
  rejection, and transport failure paths.
- Run race tests with concurrent associations and slow or failing observers.

### IMG-003: Volume Reconstruction and MPR

**Status:** `Partial`  
**Priority:** `P3`

The reconstruction package documents ImageData, VolumeData, Slice, Stack, and
DicomGenerator but implements them as placeholders. Constructors return
`ErrNotImplemented`, and `NewDicomGenerator` returns `nil`.

Reference: [fo-dicom Reconstruction](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Imaging/Reconstruction)

This item depends on IMG-001 and IMG-002.

**Acceptance criteria**

- Build a volume only from geometrically compatible slices.
- Sort slices and detect irregular spacing and frame-of-reference mismatches.
- Generate axial, coronal, sagittal, and arbitrary cuts with defined
  interpolation and out-of-volume behavior.
- Generate derived DICOM instances with valid geometry, derivation metadata,
  UIDs, and pixel representation.
- Define memory and concurrency behavior for large studies.

**Suggested verification**

- Synthetic volumes with analytically predictable cuts.
- Irregular spacing, reversed ordering, oblique orientation, and multi-frame
  datasets.
- Cross-check geometry and representative pixel values with fo-dicom.
- Add benchmarks before optimizing volume and cut generation.

### MED-002: DICOM File Scanner

**Status:** `Open`  
**Priority:** `P3`

fo-dicom provides a scanner that walks files, reports DICOM/non-DICOM results,
and integrates with media workflows. go-dicom has parsers and CLI examples but
no reusable scanner abstraction.

Reference: [fo-dicom DicomFileScanner](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Media/DicomFileScanner.cs)

**Acceptance criteria**

- Scan files and directory trees with context cancellation.
- Report valid DICOM files, invalid files, and read errors independently.
- Make recursion, symlink, concurrency, and stop-on-error behavior explicit.
- Avoid loading large pixel data when only classification or metadata is needed.

**Suggested verification**

- Scan a mixed tree containing DICOM, non-DICOM, unreadable, and symlinked files.
- Test cancellation, bounded concurrency, deterministic result accounting, and
  both stop-on-error modes.
- Instrument large fixtures to verify metadata-only scans do not read pixel
  payloads.

## External Capability Boundary

### Compressed Codecs

**Status:** `External`

The core repository intentionally provides native transfer syntax codecs, the
registry, transcoder, and encapsulation support. Compressed codecs are supplied
by `github.com/cocosip/go-dicom-codecs` through blank-import registration.

fo-dicom Core includes RLE and JPEG Lossless decoder implementations, but this
does not require moving all compressed codecs back into `go-dicom`. Capability
claims must name the companion module and verification must cover the combined
runtime.

### Platform Integrations

**Status:** `Not a gap`

WPF, ImageSharp, SkiaSharp, ASP.NET dependency injection, and .NET-specific
async API shapes are platform integrations. They should not be ported unless a
separate Go use case and package boundary are approved.

### Whole Slide Imaging

**Status:** `Not a one-sided gap`

Both libraries expose low-level WSI-compatible tags and multi-frame pixel data,
but neither audited revision provides a complete pyramid/level/coordinate WSI
domain API or WSI IOD validator. A future WSI proposal should be scoped as a new
cross-library capability rather than fo-dicom parity work.

## Delivery Phases

### Phase 0: Public Contract and Interoperability

Scope: DOC-001, NET-001, NET-002, STD-001.

Current progress: NET-001, NET-002, and STD-001 are complete. DOC-001 remains
partial; therefore Phase 0 is not complete.

Phase acceptance:

- README examples compile and capability wording matches actual public APIs.
- High-level SCU connections work over plain TCP and validated TLS.
- C-STORE sends the negotiated transfer syntax or performs verified transcoding.
- Generated standards data is reproducible from a clean checkout with no diff.
- Focused tests, full package tests, race tests for changed shared/network code,
  and relevant interoperability checks pass.

### Phase 1: Major Domain Parity

Scope: MED-001, NET-003, NET-004, SR-001, IMG-001, CORE-001.

Current progress: MED-001, NET-003, NET-004, and CORE-001 are complete. SR-001
and IMG-001 remain incomplete; therefore Phase 1 is not complete.

Phase acceptance:

- DICOMDIR creation and round trip work with both libraries.
- Advanced association values are negotiated and enforced by client behavior.
- All declared SR Value Types round-trip through Dataset and file workflows.
- Rendering consumes Dataset metadata and produces verified output.
- Dataset validation reports nested paths and validates sequence children.

Each item should be implemented and released independently; Phase 1 is not a
single pull request.

### Phase 2: Supporting APIs and Operations

Scope: IMG-002, CORE-002, DICT-001, ANON-001, PRINT-001, OBS-001.

Current progress: DICT-001 and ANON-001 are complete. IMG-002, CORE-002,
PRINT-001, and OBS-001 remain incomplete; therefore Phase 2 is not complete.

Phase acceptance:

- Shared geometry and walker APIs are stable before MPR work begins.
- Runtime dictionaries and anonymization profiles have strict parsing tests.
- Print objects round-trip through Dataset and the DIMSE N-service workflow.
- Network diagnostics are injectable, structured, and silent by default.

### Phase 3: Specialized Workflows

Scope: IMG-003 and MED-002.

Phase acceptance:

- MPR geometry and pixels are validated against synthetic and reference cases.
- Scanner behavior remains bounded and cancellable for large directory trees.
- Performance benchmarks establish baselines for large studies and scans.

## Maintenance Rules

- Keep IDs stable after publication.
- Do not delete completed items; mark them `Complete` and record the verifying
  commit, tests, and any intentionally deferred scope.
- Update the audit baseline only after rerunning the source comparison.
- Do not mark an item complete from a README change or type declaration alone.
- Treat source/unit-test evidence, integration behavior, and cross-library
  interoperability as separate levels of verification.
- Add newly discovered gaps to the detailed list and the phase table before
  beginning implementation.
- Keep the planned development order statuses current without renumbering
  completed rows. Change the order only when dependency or scope evidence
  changes, and record the reason in the ordering table.
