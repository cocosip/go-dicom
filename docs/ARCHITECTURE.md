# Architecture

This document gives a high-level view of how go-dicom is organized: the layering,
the responsibility of each package, the main data flows, and the intended
extension points. It is a map, not an API reference; see the package `README.md`
files and the [networking guide](network/README.md) for details.

## Overview

go-dicom is a pure Go implementation of the DICOM standard, inspired by the
fo-dicom C# library. The code is split by domain so that the core data model
stays independent of higher-level concerns such as imaging, networking, media
handling, structured reporting, and printing.

## Layering

Dependencies point downward. A package may import packages below it, never the
other way around, and import cycles are avoided by defining interfaces in the
consumer package.

```
cmd/ and examples/            command-line tools and runnable samples
        |
pkg/network  pkg/printing  pkg/media  pkg/sr
        \        |             |        /
         \       |             |       /
            pkg/imaging  pkg/dicom/transcode
                    \        /
                    pkg/dicom          core data model (tags, VR, elements, dataset,
                        |              parser, writer, dictionary, UID, charset,
                        |              anonymizer, transfer syntax, serialization)
                    pkg/io             buffer, endian, range-backed readers
                    pkg/logging        optional process-wide logger (used by all)
```

- `pkg/io` is the lowest layer: byte buffers, endian conversion, and range-backed
  readers used by the parser and imaging code.
- `pkg/dicom` is the core DICOM data model and file format. It builds on
  `pkg/io` and has no dependency on imaging or networking.
- Domain packages (`pkg/imaging`, `pkg/network`, `pkg/media`, `pkg/sr`,
  `pkg/printing`) build on `pkg/dicom` and, where needed, `pkg/io`.
- `pkg/logging` is standalone and is the single optional logging channel for all
  packages.

## Package responsibilities

| Package | Responsibility |
|---------|----------------|
| `pkg/io` | Byte buffers (`buffer`), endian helpers (`endian`), range-backed and HTTP readers (`rangeio`, `rangehttp`) |
| `pkg/dicom/tag`, `vr`, `vm` | DICOM tags, value representations, value multiplicities |
| `pkg/dicom/element`, `dataset` | Element types, datasets, sequences, and file meta information |
| `pkg/dicom/parser`, `writer` | DICOM file reading and writing |
| `pkg/dicom/dict`, `uid`, `transfer`, `charset` | Tag dictionary, standard UIDs, transfer syntaxes, character sets |
| `pkg/dicom/serialization`, `anonymizer` | JSON/XML serialization and de-identification profiles |
| `pkg/dicom/encapsulated`, `transcode`, `parseable` | Encapsulated pixel data, dataset-level transcoding, parseable values |
| `pkg/imaging/codec` | Image codec SPI and registry (native and external codecs) |
| `pkg/imaging/pixel`, `pixeldata` | Pixel format value objects and frame-aware pixel data access |
| `pkg/imaging/geometry`, `math3d`, `transform`, `interpolation` | Frame geometry, spatial primitives, transforms, sampling |
| `pkg/imaging/lut`, `colorconv`, `render`, `reconstruction` | LUTs, color conversion, dataset-driven rendering, volume reconstruction and MPR |
| `pkg/network/pdu`, `transport` | DUL PDU encode/decode and TCP/TLS transport |
| `pkg/network/association`, `dimse`, `status` | Association negotiation, DIMSE messages, status codes |
| `pkg/network/service`, `client`, `server` | Association-scoped DIMSE engine, SCU client, SCP server |
| `pkg/network/observability` | Structured logs, lifecycle events, and metrics hooks |
| `pkg/media`, `pkg/media/scanner` | DICOMDIR and filesystem media workflows, bounded metadata scanning |
| `pkg/sr` | Structured Report parsing and construction |
| `pkg/printing` | Film Session / Film Box / Image Box print management |
| `pkg/logging` | Optional process-wide logger |

## Key data flows

- **Reading a file**: `pkg/io` readers feed `pkg/dicom/parser`, which detects the
  file format, decodes elements and sequences into a `dataset.Dataset`, and can
  defer large values (for example pixel data) until they are accessed.
- **Writing a file**: a `dataset.Dataset` is passed to `pkg/dicom/writer`, which
  serializes elements, generates file meta information, and writes through
  `pkg/io`.
- **Imaging**: `pkg/imaging/pixeldata` extracts frame-aware pixel data from a
  dataset; `pkg/imaging/codec` and `pkg/dicom/transcode` handle (de)compression
  and transfer-syntax conversion; `pkg/imaging/render`, `lut`, and `colorconv`
  produce displayable images.
- **Networking**: TCP/TLS transport carries DUL PDUs; association negotiation
  selects presentation contexts and SCU/SCP roles; the DIMSE `service` layer
  dispatches requests to `client` (SCU) or `server` (SCP) handlers. C-GET uses
  C-STORE sub-operations over the same association, which is why storage SOP
  Classes must be negotiated with the SCP role (see the
  [networking guide](network/README.md)).
- **Anonymization**: `pkg/dicom/anonymizer` applies a profile of actions to a
  dataset, walking sequences recursively and rewriting values while preserving
  structure.

## Extension points

- **Codecs**: implement the `pkg/imaging/codec` SPI and register it; the registry
  is what the transcode layer uses.
- **Transfer syntaxes and UIDs**: the registries in `pkg/dicom/transfer` and
  `pkg/dicom/uid` describe supported syntaxes and SOP Classes.
- **Presentation contexts**: the client exposes `PresentationContextSpec`
  factories and `AddPresentationContexts`, plus `AddStoragePresentationContexts`
  and `AddStoragePresentationContextsByCategory`; ManagedClient jobs derive
  contexts from their own specification.
- **Server handlers**: register DIMSE and lifecycle handlers on the SCP server;
  C-GET requires a C-STORE handler so sub-operations can be delivered.
- **Logging**: configure the optional logger once with `pkg/logging.Configure`;
  leaving it unconfigured is fully supported and silent.

## Design conventions

- Generated tables (tags, dictionaries, UIDs) are produced by `tools/`; do not
  edit them by hand.
- Dataset value access is centered on `pkg/dicom/dataset` accessors; pixel data
  ownership is centered on `pkg/imaging/pixeldata.Data`.
- Dataset-level transcoding lives in `pkg/dicom/transcode`; the codec SPI lives in
  `pkg/imaging/codec`; pixel format value objects live in `pkg/imaging/pixel`.
- Networking operations take a `context.Context` and honor configurable timeouts;
  errors are wrapped with `%w` so callers can use `errors.Is` / `errors.As`.
- All packages log through the single optional `pkg/logging` channel and never
  read or modify `slog.Default()`.

## See also

- [Documentation index](README.md)
- [Networking guide](network/README.md)
- [Networking troubleshooting](network/TROUBLESHOOTING.md)
- [Changelog](CHANGELOG.md)
- [Benchmarks](BENCHMARKS.md)
