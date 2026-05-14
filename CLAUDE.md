# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build ./...

# Test all packages
go test ./...

# Test with race detector and coverage (matches CI)
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run a single test
go test -v ./pkg/dicom/parser -run TestMultiFrame

# Lint (matches CI: golangci-lint v2.10.1)
golangci-lint run

# Format and vet
go fmt ./...
go vet ./...

# Benchmarks
go test -bench=. -benchmem ./pkg/dicom/...
```

## Commit Convention

Commits use emoji-prefixed conventional commits: `🐛 fix(scope): desc`, `✨ feat(scope): desc`, `🏗️ build(deps): desc`. Match the style of recent git log entries.

## Architecture

This is a pure-Go port of the C# [fo-dicom](https://github.com/fo-dicom/fo-dicom) library. Go 1.25+, module `github.com/cocosip/go-dicom`, only two dependencies: `google/uuid` and `golang.org/x/text`.

### Three-layer structure

**Core DICOM (`pkg/dicom/`)** — foundational types with no external dependencies beyond the two above:

| Package | Role |
|---|---|
| `tag/` | 5338 standard tag constants; `tags_generated.go` is code-generated — do not hand-edit |
| `vr/` | 35 Value Representations with validation |
| `vm/` | Value multiplicities |
| `element/` | All element types (see below) |
| `dataset/` | `Dataset`, `Sequence`, `FileMetaInformation` containers |
| `dict/` | Tag dictionary; `dictionary_data.go` is code-generated |
| `uid/` | 1965 standard UIDs; `uids_generated.go` is code-generated |
| `transfer/` | Transfer syntax constants + `Syntax` struct (builder pattern) |
| `charset/` | 30+ character encodings via `golang.org/x/text` |
| `serialization/` | DICOM JSON (Part 18) and XML |
| `anonymizer/` | Configurable anonymization profiles |
| `endian/` | Zero-allocation byte-order helpers |
| `parser/` | DICOM file parsing |
| `writer/` | DICOM file writing |

**Imaging (`pkg/imaging/`)** — image processing on top of core types:
- `codec/` — codec registry; native (uncompressed) and RLE codecs; transcoder framework
- `lut/` — Modality, VOI, and Palette lookup tables
- `render/` — rendering pipeline
- `reconstruction/` — pixel data reconstruction

**Networking (`pkg/network/`)** — full DIMSE implementation:
- `pdu/` — 7 PDU types (A-ASSOCIATE-RQ/AC/RJ, A-RELEASE, A-ABORT, P-DATA)
- `dimse/` — all C-services (ECHO/STORE/FIND/MOVE/GET) and N-services
- `association/` — association negotiation
- `service/` — DIMSE service layer; `CMoveOperation`/`CGetOperation` embed `SubOperationResponder` for streaming progress
- `client/` — SCU (caller side); `server/` — SCP (handler side)
- `transport/` — TCP/TLS transport abstraction

**Other packages**: `pkg/sr/` (Structured Reports), `pkg/printing/` (Film/Image Box), `pkg/io/buffer/` (`ByteBuffer` abstraction used throughout).

**CLI tools** (`cmd/`): `dicominfo`, `dicomdump`, `dicom2json`.

**Code-gen tools** (`tools/`): `generate_tags`, `generate_uids`, `generate_dict` — re-run to regenerate the `*_generated.go` files.

### Key design decisions

**Element type hierarchy** — `Element` interface with concrete typed structs (`String`, `UnsignedShort`, `OtherWord`, etc.). The parser constructs the right concrete type from the VR field. Type-assert to the concrete type to access typed accessors (`GetString()`, `GetValue(index)`, `GetData()`).

**Lazy loading** — `LazyByteBuffer` defers reading large elements until accessed, provided the underlying `io.Reader` is seekable. `ReadLargeOnDemand` and `SkipLargeTags` parser options control this. `ReadAll` forces eager loading.

**Parser result** — `ParseResult` carries `FileMetaInformation` (always Explicit VR Little Endian, Group 0002), `Dataset` (main payload), `TransferSyntax`, and `Format`. Raw datasets without File Meta Information require `parser.WithAssumedTransferSyntax(...)`.

**Writer global defaults** — call `writer.SetDefaultImplementationClassUID` and `writer.SetDefaultImplementationVersionName` once at application startup rather than per-file.

**DIMSE streaming** — C-MOVE and C-GET SCP handlers receive an operation object (`service.CMoveOperation` / `service.CGetOperation`) that embeds `SubOperationResponder`. Call `op.SendPending(remaining, completed, failed, warning)` after each sub-operation, then `op.SendSuccess()` or `op.SendFailure(code)`.

**Transfer syntax builder** — new syntaxes are constructed via `transfer.NewBuilder(uid).SetExplicitVR(bool).SetEndian(...).Build()`.
