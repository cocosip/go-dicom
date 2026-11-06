# go-dicom Documentation

This directory is the single entry point for go-dicom documentation. Guides are
written in English; package-specific details live in each package's `README.md`.

## Guides

- [Architecture](ARCHITECTURE.md) — layering, package responsibilities, key data
  flows, extension points, and design conventions.
- [Networking guide](network/README.md) — SCU client (session style and managed
  queue style), SCP server, presentation-context and SCU/SCP role negotiation,
  C-GET storage contexts and factories, DIMSE messages, configuration, and status
  codes.
- [Networking troubleshooting](network/TROUBLESHOOTING.md) — common connection,
  association, and DIMSE problems, plus debugging tips.

## Changelog and benchmarks

- [Changelog](CHANGELOG.md)
- [Benchmarks](BENCHMARKS.md)

## Package references

Each package keeps a README with package-specific details:

- [pkg/network](../pkg/network/README.md)
- [pkg/network/dimse](../pkg/network/dimse/README.md)
- [pkg/dicom/element](../pkg/dicom/element/README.md)
- [pkg/dicom/transcode](../pkg/dicom/transcode/README.md)
- [pkg/imaging](../pkg/imaging/README.md)
- [pkg/imaging/render](../pkg/imaging/render/README.md)
- [pkg/imaging/reconstruction](../pkg/imaging/reconstruction/README.md)
- [pkg/io](../pkg/io/README.md)
- [pkg/media](../pkg/media/README.md)
- [pkg/sr](../pkg/sr/README.md)

## Examples

Runnable examples live under [`examples/`](../examples/). Most example
directories include their own README.

## Other

- [Repository guidelines](../AGENTS.md)
