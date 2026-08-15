# CORE-001 Recursive Validation Design

Date: 2026-08-15
Status: Implemented

## Context

go-dicom has value validation on several concrete element types, but it does
not have a public Dataset validation workflow. `Sequence.Validate` is a
placeholder, Dataset insertion does not validate values, and the existing
`vr.PerformValidation` switch is mutable global state.

This design aligns CORE-001 with fo-dicom revision
`7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2`. Go APIs and implementation
details may be idiomatic, but validation behavior, ordering, and exceptions
must remain materially equivalent.

## Decisions

CORE-001 combines two complementary capabilities:

1. An explicit recursive validation engine used by `Dataset.Validate()`.
2. Dataset-level automatic validation used by insertion APIs when enabled.

Both capabilities use the same internal validation engine. There will not be
separate explicit-validation and insertion-validation implementations.

Automatic validation is enabled by default. Explicit `Validate()` always
validates, regardless of the automatic-validation setting. This matches the
fo-dicom distinction between `DicomDataset.Validate()` and `ValidateItems`.

## Goals

- Validate Dataset elements deterministically and fail on the first error.
- Recursively validate Sequence items in item order.
- Validate an element's actual VR value before validating dictionary VM.
- Match fo-dicom's private-tag, empty-value, and VR-specific VM exceptions.
- Return a structured error containing the complete nested tag/item path.
- Validate `Add`, `AddOrUpdate`, and bulk insertion when automatic validation
  is enabled, before committing the affected element.
- Allow parsing and deserialization of invalid DICOM data while retaining the
  ability to validate the resulting Dataset explicitly.
- Replace validation control in Dataset workflows with per-Dataset state so
  independent Dataset operations do not race through a global switch.

## Non-goals

- IOD or module conformance, including Type 1, Type 1C, Type 2, and Type 2C
  presence rules.
- Pixel-data semantic validation.
- File Meta Information conformance beyond the existing
  `FileMetaInformation.Validate()` workflow.
- Rejecting an element because its actual VR is not one of the dictionary's
  listed VRs. fo-dicom's element validation validates the actual VR value and
  VM but does not perform this dictionary-VR consistency check here.
- Making Dataset mutation goroutine-safe. The existing external
  synchronization requirement remains.

## Public API

The Dataset API will add:

```go
func (ds *Dataset) Validate() error
func (ds *Dataset) AutoValidate() bool
func (ds *Dataset) SetAutoValidate(enabled bool)
```

`Validate()` performs full explicit validation even when `AutoValidate()` is
false. `SetAutoValidate` only controls automatic validation performed by
subsequent mutations.

`New()` and `NewWithTransferSyntax()` return Datasets with automatic
validation enabled. A zero-value `Dataset` also behaves as enabled. The
implementation should therefore store the inverse state, for example
`skipValidation bool`, so the Go zero value has the required behavior.

The existing constructor changes from:

```go
func NewWithElements(elements []element.Element) *Dataset
```

to:

```go
func NewWithElements(elements []element.Element) (*Dataset, error)
```

It retains the current replacement rule: when a tag occurs more than once,
the later element replaces the earlier element. Every candidate is validated
before insertion. On the first error it returns `nil` and the error, never a
partially constructed Dataset.

`Merge` will return an error so automatic-validation failures can be reported:

```go
func (ds *Dataset) Merge(other *Dataset, overwrite bool) error
```

Existing Go call sites that use `ds.Merge(...)` as a statement remain source
compatible because a function result may be discarded. New and updated
internal call sites must handle the result where failure matters.

## Validation Architecture

### Single engine

Dataset owns traversal, path construction, dictionary VM validation, and
error wrapping. Element owns validation of its concrete value representation.
Sequence owns its item collection but delegates recursive Dataset validation
to the shared Dataset engine.

Built-in elements will have an internal forced-validation path that does not
read `vr.PerformValidation`. Existing direct `Element.Validate()` calls remain
source compatible; they may retain legacy global-switch behavior during this
change. Dataset explicit and automatic validation must never temporarily
modify the global variable.

For an external implementation of `element.Element`, the Dataset engine calls
its public `Validate()` method because go-dicom cannot inspect custom value
semantics. Path and VM handling still come from the Dataset engine.

### Dataset traversal

Validation iterates `Dataset.Elements()`, which is sorted by tag. It stops at
the first failure. An empty Dataset is valid. Calling `Validate()` on a nil
Dataset receiver returns a structural validation error rather than panicking.

For each item:

1. Check required object structure, including a non-nil element, tag, and VR.
2. If it is a Sequence, recursively validate its non-nil child Datasets in
   item order and do not run normal element VM validation on the Sequence.
3. Otherwise validate the value according to the element's actual VR.
4. If applicable, look up and validate the public dictionary VM.

This is fail-fast behavior. CORE-001 does not add an aggregate-all-errors API.

### Value and VM rules

VR value validation precedes VM validation, matching fo-dicom.

Dictionary VM validation applies only when all of the following are true:

- the tag is public;
- `Count() > 0`;
- dictionary metadata is available; and
- the element type is not one of fo-dicom's VM-exempt types.

VM validity includes minimum, maximum, and stepped multiplicity such as
`2-2n`; it must use `vm.Parse(...).IsValid(count)` rather than a simple range
check.

The fo-dicom VM exemptions retained by this design are:

- Sequence (`SQ`);
- other/bulk value representations `OB`, `OW`, `OL`, `OD`, `OF`, and `OV`;
- Unlimited Characters (`UC`); and
- Unknown (`UN`).

Private elements validate their actual VR values but skip standard dictionary
VM validation. Unknown public tags also validate their actual VR values and
skip VM when no dictionary entry exists.

Nil Sequence items are skipped, matching fo-dicom. Empty Sequence items are
valid empty Datasets. `NewSequenceWithItems` may therefore preserve nil items,
even though `AddItem` currently ignores nil arguments.

## Automatic Validation Lifecycle

When automatic validation is enabled:

- `Add` validates the candidate before inserting it. A validation failure or
  duplicate-tag failure leaves the Dataset unchanged.
- `AddOrUpdate` validates the candidate before replacing an existing element.
  A failure preserves the previous element.
- `Merge` first validates every candidate that would be inserted or replaced,
  then applies the mutation. A failure leaves the destination Dataset
  unchanged. This gives Go callers an atomic bulk-operation contract while
  retaining the same validation rules as fo-dicom.
- Adding a Sequence validates all nested items even if a child Dataset has
  automatic validation disabled, because Sequence validation is explicit.

When automatic validation is disabled, mutation APIs retain their existing
structural checks, such as nil elements and duplicate tags, but skip VR, VM,
and recursive value validation.

`Clone` and `Filter` are structural copies, not new user insertions. They copy
elements without retroactive validation and preserve the source Dataset's
automatic-validation setting. This prevents invalid-but-readable input from
being silently dropped during a copy.

## Parser and Deserializer Behavior

Binary parser, JSON reader, and XML reader hydration paths must disable
automatic validation on every Dataset they construct, including Sequence item
Datasets. After hydration they restore automatic validation to enabled without
retroactively validating existing values.

This matches fo-dicom's ability to read an invalid file, inspect its contents,
and later receive an error from explicit `Dataset.Validate()`.

Normal application-created Datasets continue to default to automatic
validation enabled.

## Error Model

The dataset package will expose a structured validation error:

```go
type ValidationKind string

const (
    ValidationStructural ValidationKind = "structural"
    ValidationValue      ValidationKind = "value"
    ValidationVM         ValidationKind = "vm"
)

type ValidationPathSegment struct {
    Tag       *tag.Tag
    ItemIndex *int
}

type ValidationError struct {
    Kind  ValidationKind
    Path  []ValidationPathSegment
    Cause error
}

func (e *ValidationError) Error() string
func (e *ValidationError) Unwrap() error
```

An item index belongs to the Sequence tag that selected that item. A nested
path is formatted consistently as:

```text
(0008,1115)[0]/(0008,1140)[2]/(0008,1155)
```

The original element or VR error is retained as `Cause`, and `Unwrap()` allows
`errors.Is` and `errors.As` to inspect it. Recursive layers prepend path
segments without replacing an existing `ValidationError` cause.

Automatic insertion returns the same `ValidationError` shape as explicit
validation. Constructor and Merge errors retain the path of the failing
candidate.

## Global Validation Compatibility

`vr.PerformValidation` is an existing exported bool and cannot be made
race-safe while preserving direct external assignment compatibility. CORE-001
will mark it deprecated and stop using it in new Dataset validation paths.

Per-Dataset automatic validation and explicit Dataset validation therefore do
not produce a global configuration race. Code that continues to mutate
`vr.PerformValidation` concurrently remains outside the new guarantee and
should migrate to Dataset-level control.

## Expected Implementation Areas

- `pkg/dicom/dataset/dataset.go`: automatic-validation state and mutation API
  changes.
- `pkg/dicom/dataset/validation.go`: traversal, VM rules, paths, and errors.
- `pkg/dicom/dataset/sequence.go`: shared recursive validation integration.
- `pkg/dicom/element`: forced built-in value validation without global state.
- `pkg/dicom/vr`: unconditional internal VR validation primitive and legacy
  wrapper compatibility.
- `pkg/dicom/parser`: validation-disabled binary hydration.
- `pkg/dicom/serialization`: validation-disabled JSON and XML hydration.
- `docs/FO_DICOM_GAP_ANALYSIS.md` and its Chinese counterpart: completion
  status after all acceptance gates pass.

## Testing Strategy

Focused tests will cover:

- direct Dataset value-validation success and failure;
- deterministic first failure by sorted tag;
- two-level nested Sequence failure with exact path and retained cause;
- nil and empty Sequence items;
- DA, TM, DT, UI, DS, IS, PN, and multi-value strings;
- VM minimum, maximum, fixed, unlimited, and stepped multiplicity;
- `Count() == 0` VM bypass;
- private explicit-VR value validation with VM bypass;
- unknown public explicit-VR value validation without dictionary VM;
- VM exemptions for SQ, OB, OW, OL, OD, OF, OV, UC, and UN;
- default automatic validation for `New`, `NewWithTransferSyntax`, and a
  zero-value Dataset;
- disabled automatic validation followed by failing explicit validation;
- `Add` and `AddOrUpdate` preserving Dataset state on failure;
- `NewWithElements` returning no partial Dataset on failure;
- atomic `Merge` failure;
- Clone and Filter preserving invalid readable content and validation mode;
- parser, JSON, and XML reading invalid values successfully, followed by an
  explicit validation failure; and
- independent concurrent validation of separate Datasets with different
  automatic-validation settings, without changing global state.

Verification gates:

```text
go test ./pkg/dicom/dataset ./pkg/dicom/element ./pkg/dicom/parser ./pkg/dicom/serialization
go test ./cmd/... ./examples/... ./pkg/... ./tools/...
go build ./...
golangci-lint run
```

Race tests should run for the affected packages and full suite where the local
Go runtime supports `-race`. If the known Windows runtime startup failure
`0xc0000139` persists, it must be reported separately rather than treated as a
passing race gate.

## Acceptance

CORE-001 is complete when:

- explicit and automatic Dataset validation use the same recursive engine;
- the validation order and exceptions above match fo-dicom;
- invalid input can still be parsed with later explicit validation;
- errors contain an exact nested path and unwrap to the original cause;
- Dataset-level validation control does not depend on mutable global state;
- focused and full non-race verification gates pass; and
- the gap-analysis documents mark only CORE-001 complete and advance the next
  item according to the existing planned development order.
