# CORE-001 Recursive Validation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [x]`) syntax for tracking.

**Goal:** Implement fo-dicom-aligned explicit recursive validation and default Dataset insertion validation without global validation-state races.

**Architecture:** A single Dataset validation engine owns traversal, paths, VM checks, and error wrapping. Built-in elements expose an unconditional value-validation path; explicit validation and automatic insertion both call the same engine, while parsers hydrate validation-disabled Datasets.

**Tech Stack:** Go 1.24, standard `errors`/`fmt`, existing `element`, `vr`, `vm`, `dictif`, parser, and serialization packages.

**Spec:** `docs/CORE_001_RECURSIVE_VALIDATION_DESIGN.md`

**Status:** Implemented; the local Windows race runtime remains unavailable
with exit status `0xc0000139`.

## Global Constraints

- Align behavior with fo-dicom revision `7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2`.
- Explicit `Dataset.Validate()` always validates and fails on the first error.
- Automatic validation is enabled by default, including for a zero-value Dataset.
- Dataset validation must not read or mutate `vr.PerformValidation`.
- Invalid binary, JSON, and XML input remains readable and can fail later explicit validation.
- Preserve the original validation error through `Unwrap()`.
- Do not add IOD Type 1/2 or pixel semantic validation.

---

### Task 1: Unconditional Built-in Element Validation

**Files:**
- Modify: `pkg/dicom/vr/vr.go`
- Modify: `pkg/dicom/element/element.go`
- Modify: `pkg/dicom/element/string.go`
- Modify: `pkg/dicom/element/date.go`
- Modify: `pkg/dicom/element/numeric_string.go`
- Modify: `pkg/dicom/element/person_name.go`
- Test: `pkg/dicom/element/string_test.go`

**Interfaces:**
- Produces: `vr.(*VR).ValidateStringValue(string) error`, which ignores the legacy global switch.
- Produces: `element.ValidateValue(Element) error`, which forces built-in value validation and falls back to `Element.Validate()` for external implementations.

- [x] Write a test that sets `vr.PerformValidation = false`, verifies direct legacy `elem.Validate()` skips the error, and verifies `element.ValidateValue(elem)` still rejects an invalid UID.
- [x] Run `go test ./pkg/dicom/element -run TestValidateValueIgnoresGlobalSwitch -count=1` and verify it fails because `ValidateValue` is undefined.
- [x] Add `ValidateStringValue`, make legacy `ValidateString` gate and delegate to it, and add an internal `validateValue() error` interface implemented by built-in string wrappers.
- [x] Run `go test ./pkg/dicom/element -run TestValidateValueIgnoresGlobalSwitch -count=1` and verify it passes.

### Task 2: Dataset Recursive Engine, VM, and Paths

**Files:**
- Create: `pkg/dicom/dataset/validation.go`
- Create: `pkg/dicom/dataset/validation_test.go`
- Modify: `pkg/dicom/dataset/sequence.go`
- Modify: `pkg/dicom/dictif/interface.go` only if a typed VM lookup helper is required.

**Interfaces:**
- Produces: `func (ds *Dataset) Validate() error`.
- Produces: `ValidationKind`, `ValidationPathSegment`, and `ValidationError` with `Error()` and `Unwrap()`.
- Consumes: `element.ValidateValue(Element) error`.

- [x] Write failing tests for sorted fail-fast behavior, public VM, `Count()==0`, private/unknown VM bypass, VM-exempt VRs, and nested path `(0008,1115)[0]/(0008,1140)[0]/(0008,1155)`.
- [x] Run `go test ./pkg/dicom/dataset -run 'TestDatasetValidate|TestSequenceValidate' -count=1` and verify missing API/behavior failures.
- [x] Implement recursive traversal, structural checks, forced value validation, `vm.Parse(...).IsValid`, VM exemptions, path formatting, and error unwrapping.
- [x] Change `Sequence.Validate()` to call the shared recursive engine and skip nil items.
- [x] Run the focused Dataset tests and verify they pass.

### Task 3: Automatic Validation and Mutation APIs

**Files:**
- Modify: `pkg/dicom/dataset/dataset.go`
- Modify: `pkg/dicom/dataset/dataset_test.go`
- Modify: internal callers affected by `NewWithElements` and `Merge` signatures.

**Interfaces:**
- Produces: `func (ds *Dataset) AutoValidate() bool`.
- Produces: `func (ds *Dataset) SetAutoValidate(enabled bool)`.
- Changes: `NewWithElements([]element.Element) (*Dataset, error)`.
- Changes: `Merge(*Dataset, bool) error`.

- [x] Write failing tests for default/zero-value automatic validation, disabled insertion, failed Add/AddOrUpdate state preservation, constructor failure returning nil Dataset, atomic Merge, and Clone/Filter mode preservation.
- [x] Run the focused tests and verify the old insertion behavior fails them.
- [x] Add inverse `skipValidation` state, validate before mutation, prevalidate Merge candidates, and direct-copy Clone/Filter without retroactive validation.
- [x] Update all compile-time callers for the constructor and Merge signatures.
- [x] Run `go test ./pkg/dicom/dataset -count=1` and verify it passes.

### Task 4: Validation-disabled Hydration

**Files:**
- Modify: `pkg/dicom/parser/parser.go`
- Modify: `pkg/dicom/serialization/json_reader.go`
- Modify: `pkg/dicom/serialization/xml.go`
- Test: `pkg/dicom/parser/parser_test.go` or nearest parser round-trip test.
- Test: `pkg/dicom/serialization/json_test.go`
- Test: `pkg/dicom/serialization/xml_test.go`

**Interfaces:**
- Consumes: `Dataset.SetAutoValidate(false)` during hydration.
- Produces: returned root and nested Datasets restored to `AutoValidate()==true` without retroactive validation.

- [x] Write JSON and XML tests containing an invalid UID that deserialize successfully and then fail `Dataset.Validate()`.
- [x] Add an equivalent binary parser regression using a written invalid Dataset or a minimal encoded fixture.
- [x] Run the three focused tests and verify automatic insertion initially rejects hydration.
- [x] Disable automatic validation immediately after constructing every hydrated root/item Dataset and restore it before return.
- [x] Stop ignoring XML item `Add` structural errors.
- [x] Run `go test ./pkg/dicom/parser ./pkg/dicom/serialization -count=1` and verify it passes.

### Task 5: Full Verification and Backlog Update

**Files:**
- Modify: `docs/FO_DICOM_GAP_ANALYSIS.md`
- Modify: `docs/FO_DICOM_GAP_ANALYSIS.zh-CN.md`

**Interfaces:**
- Produces: CORE-001 marked `Complete`; the next item remains the first incomplete row in planned order.

- [x] Run `gofmt` on every changed Go file.
- [x] Run `go test ./pkg/dicom/dataset ./pkg/dicom/element ./pkg/dicom/parser ./pkg/dicom/serialization -count=1`.
- [x] Run `go test ./cmd/... ./examples/... ./pkg/... ./tools/...`.
- [x] Run `go build ./...`.
- [x] Run affected-package race tests; if Windows returns `0xc0000139`, record it as an unverified race gate.
- [x] Run `golangci-lint run` and report any environment blocker separately.
- [x] Compare implemented behavior against every acceptance item in the design spec.
- [x] Mark CORE-001 complete in both gap documents and identify `SR-001` as next only after all required gates pass.
- [x] Run `git diff --check` and review `git diff --stat` before committing the implementation.
