# Changelog

All notable changes to go-dicom are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/).


## [v0.8.4] - 2026-11-07

### Added

- **network/client**: `AddStoragePresentationContexts`, `AddStoragePresentationContextsByCategory`, and `AddPresentationContexts` to declare Storage SOP Classes with the SCP role before connecting, plus the shared `PresentationContextSpec` type with `NewPresentationContextSpec`, `StoragePresentationContextSpecs`, and `StoragePresentationContextSpecsFor` factories.
- **network/client**: `NewCGetJob` accepts optional storage SOP Class UIDs and negotiates the SCP role for them automatically.
- **dicom/uid**: `StorageSOPClassUIDs(category)` enumerates the active standard Storage SOP Classes for a storage category.
- **docs**: consolidated networking guide, architecture overview, documentation index, and this changelog under `docs/`.

### Changed

- **network/client**: C-GET now preflight-validates that a C-STORE handler is configured and that at least one Storage SOP Class is negotiated with the SCP role, returning an actionable error instead of a silent `0xB000`.
- **license**: relicensed the project from MS-PL to the MIT License; Go source headers, license badges, and documentation were updated.
- **network/dimse**: request constructors no longer accept caller-supplied Message IDs; the service assigns IDs before sending.
- **network**: association defaults now use one implementation UID and version name across PDU, client, server, and service layers.

### Fixed

- **network/client**: C-GET no longer reports success while delivering zero instances when storage sub-operations cannot be sent back to the requestor.
- **network**: reject malformed DIMSE command sets with zero or duplicate request Message IDs, invalid priorities, or inconsistent dataset flags.
- **network**: validate association implementation identity and negotiated PDU length accounting before transmission.

## [v0.8.3] - 2026-11-03

### Added
- unify diagnostic rendering for datasets, elements, and associations

### Fixed
- correct diagnostic rendering and presentation context output

### CI
- **github**: simplify CI pipelines and add concurrency control

## [v0.8.2] - 2026-10-12

### CI
- **github**: upgrade golangci-lint to v2.14.0

### Chores
- **deps**: update golang.org/x/image and golang.org/x/text

## [v0.8.1] - 2026-09-15

### Fixed
- **transcode**: preserve dataset validation policy during transcoding

### Tests
- **transcode**: centralize invalid patient age test value

## [v0.8.0] - 2026-09-15

### Added
- strengthen DICOM pixel transforms and LUT conformance
- **dataset**: introduce dictionary-driven value APIs and typed element factories
- **dicom**: add charset-aware dataset values and pixel data writing
- **dicom**: add typed VR, pixel, LUT, and charset support

### Fixed
- **dicom**: prevent mutation of registered UIDs and transfer syntaxes
- **dicom**: align pixel VR, frame offsets, and LUT precedence
- **imaging**: handle native byte order and configurable LUT VR validation
- replace element usage with value-based methods for dataset updates
- **dicom**: enforce character set validation and propagate decoding errors
- **network**: discard stale wake notifications before waiting
- **dicom**: validate character sets and propagate transcoder errors
- harden DICOM data and client contracts

### Performance
- **dicom**: reduce dataset value conversion allocations

### Changed
- **packages**: consolidate ownership and registry boundaries

### Changed
- **packages**: consolidate ownership and registry boundaries

### Documentation
- **api**: refresh dataset, pixel data, and DICOM examples

## [v0.7.5] - 2026-09-10

### Added
- **network**: enrich DIMSE and association observability
- **imaging**: validate LUT rendering and support alpha palette output

### Fixed
- **lut**: align DICOM LUT parsing and presentation modes
- enforce DICOM Pixel Data VR and LUT encoding rules
- **imaging**: enforce LUT descriptor and data validation

### Changed
- **logging**: simplify structured log attribute initialization

### Changed
- **logging**: simplify structured log attribute initialization

## [v0.7.4] - 2026-08-31

### Added
- **logging**: add structured slog observability across DICOM, imaging, and network flows
- **logging**: centralize optional process-wide slog emission

## [v0.7.3] - 2026-08-29

### Fixed
- **anonymizer**: apply retain profiles on top of the basic anonymization rules
- **network**: add forced server close after graceful shutdown timeout
- tighten anonymizer file metadata handling and progressive DIMSE error propagation
- **anonymizer**: preserve explicit keep overrides and startup/final-response handling

## [v0.7.2] - 2026-08-27

### Fixed
- **network**: scope server timeouts to accept, negotiation, and per-PDU I/O
- **imaging**: preserve transcoded pixel metadata and lossy compression attributes

## [v0.7.1] - 2026-08-26

### Added
- **network**: add managed SCU workflows, streaming C-FIND, and rendering text transforms

### CI
- **github**: update GitHub Actions versions in workflow configs
- **github**: update Go version to 1.27 in workflow configs

## [v0.7.0] - 2026-08-18

### Added
- **network**: add TLS support to high-level DICOM client
- **dictionary**: align data with fo-dicom 2026b
- **network**: add negotiated C-STORE transcoding
- **media**: add DICOMDIR media directory support
- **network**: add SOP Class Common Extended Negotiation support
- **core**: add recursive Dataset validation
- **sr**: complete typed structured report I/O and validation
- **imaging**: render images from dataset metadata
- **reconstruction**: implement CT/MR volume reconstruction and derived MPR output
- **dicom**: preserve typed binary and long numeric VRs in XML and parsing
- **printing**: implement dataset-backed print management and workflow support
- **network**: add structured observability hooks for client, server, and service
- **printing**: expand dataset-backed print workflow and folder persistence

### Fixed
- **anonymizer**: complete custom profile loading for security profiles
- **network**: enforce negotiated association roles and async request limits
- **network**: publish release-requested state before sending A-RELEASE-RQ
- **dicom/parser**: stop parsing before pixel data
- **printing**: validate print datasets and clean stale image files
- **printing**: allow empty optional US values and configuration text
- tighten DICOM encoding, negotiation, and default failure handling

### Changed
- **imaging**: add shared spatial foundations
- **dicom**: centralize dataset walking across validation and writing
- **dicom**: make dataset snapshots and transforms fully independent
- centralize shared test fixtures and refresh CI lint action

### Changed
- **imaging**: add shared spatial foundations
- **dicom**: centralize dataset walking across validation and writing
- **dicom**: make dataset snapshots and transforms fully independent
- centralize shared test fixtures and refresh CI lint action

### Documentation
- **docs**: add fo-dicom capability gap analysis reference
- **core**: add CORE-001 recursive validation design
- **core**: remove obsolete CORE-001 implementation docs
- remove retired capability gap analysis documents and README link

### Tests
- **generator**: normalize platform line endings
- **generate_dicom**: ensure bundled 2026b generation is deterministic

### Chores
- bump golang.org/x/text to v0.41.0

## [v0.6.2] - 2026-08-03

### Added
- **examples**: parameterize example input and output paths

### Fixed
- **network**: support early C-MOVE peer release
- **element**: validate string VR maximum length per encoded value

### Documentation
- **readme**: update release status and codec support wording

### Tests
- **parser**: add benchmark coverage and stabilize fixture-based parser tests

## [v0.6.1] - 2026-07-27

### Fixed
- **parser**: ignore zero trailing padding after datasets

## [v0.6.0] - 2026-07-22

### Changed
- **imaging**: remove built-in RLE codec support

### Changed
- **imaging**: remove built-in RLE codec support

### Build
- **deps**: bump actions/setup-go from 5 to 7

## [v0.5.3] - 2026-07-14

### Added
- **io**: add range HTTP and range IO helpers

## [v0.5.2] - 2026-07-10

### Chores
- bump golang.org/x/text to v0.40.0

## [v0.5.1] - 2026-06-30

### Added
- **pkg/network**: implement C-CANCEL-RQ support for DICOM operations

### CI
- **github**: disable fail-fast in matrix workflows

### Build
- **deps**: bump actions/checkout from 4 to 7

## [v0.5.0] - 2026-06-16

### Added
- **service**: add DIMSE-N request handling and runtime safety guards

### Fixed
- **parser**: preserve lazy loading in nested defined-length sequences
- **parser**: remove redundant lazy test conversion
- **network**: preserve presentation context on DIMSE responses and avoid abort shutdown deadlock
- **network**: preserve query model UIDs across DIMSE requests and service startup
- **dicom**: harden dataset, parser, and serialization edge cases
- **network**: add C-CANCEL request support and harden DICOM runtime edge cases

### Changed
- **network**: centralize PDU writes and tighten PDV fragmentation edge handling
- extract shared string constants across tests and protocol helpers

### Changed
- **network**: centralize PDU writes and tighten PDV fragmentation edge handling
- extract shared string constants across tests and protocol helpers

### Documentation
- simplify agent guidance documentation

### Build
- **deps**: bump codecov/codecov-action from 4 to 7
- **deps**: bump golang.org/x/text in the go-dependencies group

## [v0.4.0] - 2026-05-12

### Added
- **writer**: add deflated explicit VR little endian support
- **dicom**: add parser context cancellation and isolate transfer registry

### Fixed
- **buffer**: detect and handle composite buffer size overflow
- tighten invalid input handling across dataset, SR, and DIMSE parsing
- restore review-regressed response and constructor behavior
- **dicom**: correct fragment offset table to include item headers
- harden release handling and lazy parsing semantics
- correct charset handling, RLE decoding, and add safety checks
- support files above 4GiB and fix fragment sequence handling
- **imaging**: correct encapsulated pixel data handling and add lazy fragment support
- **parser**: add deflated transfer syntax and multi-encoding support
- **network,charset,io**: wire ctx deadline to ReadPDU, guard ISO 2022 escape fallback, default PDU length, panic on Position overflow
- **network,io**: correct readTimeoutFromContext edge cases and add coverage
- **dicom**: preserve parser context and registry semantics
- **parser,transfer**: fix context cancellation and isolate the registry

### Performance
- **imaging**: add lazy fragment loading for single-frame decode

### Changed
- rename swapPixelDataBytes and document lazy-load fallback
- **imaging/codec**: consolidate StripTrailingPadding and fix framesFromFragments

### Changed
- rename swapPixelDataBytes and document lazy-load fallback
- **imaging/codec**: consolidate StripTrailingPadding and fix framesFromFragments

### Documentation
- add AGENTS.md with Codex guidance for DICOM library port

### Build
- **deps**: update golang.org/x/text to v0.37.0

## [v0.3.1] - 2026-03-17

### CI
- **github**: update golangci-lint to v2.10.1
- **github**: replace deprecated create-release with action-gh-release@v2
- **github**: generate custom changelog for releases

### Chores
- **examples**: remove transcode example directory

## [v0.3.0] - 2026-03-17

### Added
- **printing**: add printer status management to DICOM printing module
- add QR SCP and SCU example implementations
- **parser**: implement streaming parsing for defined-length sequences and items
- **dicom**: add advanced DICOM utilities for parsing, ranges, XML deserialization, and anonymization profiles
- **examples**: enhance QR examples with complete workflow support
- **network**: support C-STORE on client association and expand DIMSE factory
- **examples/qr_scu**: add embedded Storage SCP for C-MOVE receive
- **network**: add streaming operation interfaces for C-MOVE/C-GET and standardize status codes

### Fixed
- **network**: look up VR from DICOM dictionary for implicit VR syntaxes
- **network**: improve DICOM compliance and service error handling
- **network**: dispatch requests in goroutines to prevent C-GET deadlock

### Performance
- **dataset**: add sorted tag cache for deterministic iteration performance

### Changed
- **parser**: improve DICOM format detection and parsing logic
- **examples**: improve code quality in QR examples
- **network**: inline createCCancelRequest and remove unused function

### Changed
- **parser**: improve DICOM format detection and parsing logic
- **examples**: improve code quality in QR examples
- **network**: inline createCCancelRequest and remove unused function

### Tests
- **parser**: add benchmark tests for sequence parsing

### Build
- **deps**: bump golang.org/x/text in the go-dependencies group
- **deps**: bump actions/upload-artifact from 4 to 7

### Chores
- **deps**: update golang.org/x/text to v0.35.0

## [v0.2.0] - 2026-02-12

### Added
- **cstore_scp**: add AE title validation support

### Changed
- **examples**: add package documentation and improve code formatting

### Changed
- **examples**: add package documentation and improve code formatting

### Build
- **deps**: bump golang.org/x/text in the go-dependencies group

## [v0.1.9] - 2026-01-04

### Added
- **dicom/writer**: add SetOffsetTableForFrames function for fragment sequence management

## [v0.1.8] - 2025-12-30

### Build
- **deps**: bump golang.org/x/text in the go-dependencies group

## [v0.1.7] - 2025-12-26

### Added
- **imaging/codec**: add VR selection control for DICOM pixel data transcoding

### Changed
- **dicom**: add compile-time interface checks for Element implementations

### Changed
- **dicom**: add compile-time interface checks for Element implementations

## [v0.1.6] - 2025-12-25

### Added
- **imaging**: add TranscodeWithMetadata method to preserve File Meta Information
- **imaging**: add TranscodeWithMetadata method to preserve File Meta Information

### Fixed
- **imaging**: correct VR selection for encapsulated DICOM formats

### Changed
- **examples**: remove trailing newline from transcode example
- **imaging**: rename types package to imagetypes and update VM documentation

### Changed
- **examples**: remove trailing newline from transcode example
- **imaging**: rename types package to imagetypes and update VM documentation

## [v0.1.5] - 2025-12-24

### Added
- **imaging**: enhance CreatePixelData to prioritize transfer syntax sources

## [v0.1.4] - 2025-12-23

### Added
- **imaging**: add BitsAllocated-based VR selection for PixelData
- **imaging/codec**: support OW VR for multi-byte pixel data in transcoder
- **dicom**: add transfer syntax tracking to Dataset

### Fixed
- **dicom**: always include Basic Offset Table for single-frame images

## [v0.1.1] - 2025-12-22

### Added
- **dicom**: implement fragment padding and update offset table logic
- **imaging**: strip trailing padding from JPEG/JPEG-LS fragments

## [v0.1.0] - 2025-12-22

### Added
- **parser**: implement advanced DICOM file parsing with ReadOption and FileFormat support
- **writer**: update WithExplicitLengthSequences to accept boolean parameter
- **serialization**: implement DICOM JSON serialization with BulkDataURI support
- **serialization**: implement DICOM XML serialization with full NativeDicomModel support
- **dicomanonymizer**: implement DICOM anonymization based on PS 3.15
- **network**: update DIMSE layer documentation and implement core service components
- **network**: implement C-GET and C-MOVE DIMSE operations with full support
- **network**: implement association negotiation and responder interfaces
- **network**: complete transport, client, and server implementation
- **imaging**: implement image processing foundation with RLE codec
- **imaging**: complete Stage 11 with Native codec and DicomPixelData
- **imaging**: integrate transfer syntax constants with codecs
- **imaging**: complete Stage 11.3 image rendering implementation
- **imaging**: add BitDepth, Palette LUT, and Overlay support
- **printing**: implement DICOM Print Management package
- **imaging**: implement comprehensive LUT (Lookup Table) system
- **imaging**: add Reconstruction package framework for 3D MPR
- **sr**: implement DICOM Structured Report package
- **parser**: add benchmark tests for DICOM parsing performance
- **dicom**: add support for reading and writing DICOM fragment sequences
- **imaging**: implement DICOM transcoder core functionality
- **anonymizer**: improve error handling in anonymization process
- **docs**: add package documentation to command-line tools
- **dicom**: enforce golangci-lint compliance and improve error handling
- **dicom**: improve error handling and code clarity in DICOM packages
- **dicom**: add support for additional DICOM transfer syntaxes
- **writer**: enhance DICOM file writing with complete metadata and pixel data
- **examples**: add DICOM C-STORE SCP and SCU examples
- **network**: implement proper DIMSE message encoding and decoding
- **network**: improve message processing logic for fragmented PDVs
- **network**: improve DIMSE command dataset type handling
- **cstore_scu**: replace hardcoded config with command-line flags
- **imaging**: implement optimal window calculation for DICOM images
- **examples**: add DICOM transcode example with multiple transfer syntax support
- **transcode**: add support for JPEG Extended and auto-format selection
- **imaging**: implement per-frame decoding and encoding for transcoder
- **imaging**: remove BitDepth struct and its associated methods
- **imaging**: add detailed documentation and structure for reconstruction package
- **imaging**: refactor codec interface to align with fo-dicom architecture
- **parser**: handle undefined length UN VR as sequence
- **imaging**: improve photometric interpretation constants and pixel data handling
- **dicom**: implement dictif interface compliance for dictionary and tag packages
- **codec**: update go-dicom-codec dependency and improve frame count parsing

### Changed
- **tag**: replace manual suffix removal with strings.TrimSuffix
- **tests**: improve test structure and reduce duplication
- **imaging**: move Color32 type to imaging/types package
- **imaging**: move LUT interfaces and types to shared package

### Changed
- **tag**: replace manual suffix removal with strings.TrimSuffix
- **tests**: improve test structure and reduce duplication
- **imaging**: move Color32 type to imaging/types package
- **imaging**: move LUT interfaces and types to shared package

### Documentation
- remove outdated progress and status documentation
- **.golangci.yml**: update govet configuration to include shadow checking
- remove outdated README content about GitHub workflows
- **readme**: update roadmap checklist items
- **readme**: update feature status and documentation

### CI
- **github**: update go version matrix to only include 1.25
- **github**: replace golangci-lint action with manual installation
- **github**: split test runs for windows and unix systems
- **github**: replace manual golangci-lint installation with official action
- **github**: increase golangci-lint timeout to 10 minutes

### Build
- initialize project with core DICOM infrastructure
- **deps**: update Go versions in CI workflow
- **github**: update go version to 1.25 in ci/cd workflows

### Chores
- **ci**: update golangci-lint version to latest
- **golangci**: update linter output configuration
- **golangci**: remove unused configuration settings

### Other
- I notice that the "code differences" section is empty in your request. However, I'll provide you with a template and example that follows Angular commit conventions:
- I notice that the "code differences" section is empty in your request. However, I'll provide you with a template and example that follows Angular commit conventions:

