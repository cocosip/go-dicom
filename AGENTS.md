# Repository Guidelines

## Project Structure & Module Organization
`go-dicom` is a Go module for DICOM parsing, writing, imaging, networking, and anonymization.

- `pkg/` contains the library code by domain (`pkg/dicom`, `pkg/imaging`, `pkg/network`, `pkg/io`, `pkg/printing`, `pkg/sr`).
- `cmd/` contains CLI tools such as `dicominfo`, `dicomdump`, and `dicom2json`.
- `examples/` holds runnable samples.
- `test-data/` stores sample DICOM files used by tests and manual checks.
- `tools/` contains code generators for tags, dictionaries, and UID data.

Generated tables and lookup data should be updated with the relevant tool, not by hand.

## Build, Test, and Development Commands

- `go build ./...` builds all packages and commands.
- `go test ./cmd/... ./examples/... ./pkg/... ./tools/...` runs the full unit test suite.
- `go test -v -race -coverprofile=coverage.txt -covermode=atomic ./cmd/... ./examples/... ./pkg/... ./tools/...` matches CI coverage and race checks.
- `golangci-lint run` runs the configured linter set from `.golangci.yml`.
- `go test -bench='.' -benchmem ./pkg/dicom/...` runs the main benchmark suite.

For focused debugging, run a package or test name directly, for example `go test -v ./pkg/dicom/parser -run TestMultiFrame`.

## Coding Style & Naming Conventions

- Use `gofmt` on all Go files; keep imports formatted by the toolchain.
- Follow standard Go naming: short package names, exported identifiers in `CamelCase`, unexported identifiers in `camelCase`.
- Keep package layout aligned with existing domain boundaries; avoid cross-package helpers unless broadly useful.
- Prefer small, focused changes that preserve parser and transfer-syntax behavior.

## Testing Guidelines

- Tests live beside the code they cover and usually use `*_test.go`.
- Add regression tests for parser, transfer, or buffer changes; this repo relies heavily on file-based fixtures in `test-data/`.
- Use benchmarks when changing hot paths or allocation behavior, and note any meaningful result shifts.

## Commit & Pull Request Guidelines

Keep changes focused on one topic per commit. Follow any commit-message prompt or template requested by the host app or release workflow.

Pull requests should include a short summary, the packages affected, and the verification performed (`go test`, `go test -race`, `golangci-lint`, benchmarks if relevant). Include sample output or screenshots only for CLI or example changes.

## Security & Configuration Tips

- Avoid committing local build artifacts or temporary files.
- Treat DICOM test files as fixtures; keep new samples small and representative.
- If you add a new command or generator, document its purpose in the nearest README or package doc.
