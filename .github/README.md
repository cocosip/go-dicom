# GitHub Workflows

This directory contains GitHub Actions workflows for CI/CD.

## Workflows

### CI (`ci.yml`)
Runs on every push and pull request to main branches.

**Jobs:**
- **Test**: Runs tests on Ubuntu, Windows, and macOS with Go 1.21, 1.22, and 1.23
  - Runs tests with race detector
  - Generates coverage reports
  - Uploads coverage to Codecov
- **Lint**: Runs golangci-lint with configured linters
- **Build**: Verifies that the code builds successfully
- **Benchmark**: Runs benchmarks on pull requests

### Release (`release.yml`)
Runs when a new tag is pushed (e.g., `v1.0.0`).

**Jobs:**
- Runs tests
- Builds the project
- Creates a GitHub release

### CodeQL (`codeql.yml`)
Security scanning with GitHub's CodeQL.

**Schedule:**
- Runs on every push/PR to main branch
- Weekly scheduled scan every Monday

## Dependabot (`dependabot.yml`)
Automatically updates:
- GitHub Actions versions
- Go module dependencies

Updates are checked weekly and create PRs with the `dependencies` label.

## Badge Examples

Add these to your main README.md:

```markdown
[![CI](https://github.com/cocosip/go-dicom/workflows/CI/badge.svg)](https://github.com/cocosip/go-dicom/actions?query=workflow%3ACI)
[![codecov](https://codecov.io/gh/cocosip/go-dicom/branch/master/graph/badge.svg)](https://codecov.io/gh/cocosip/go-dicom)
[![Go Report Card](https://goreportcard.com/badge/github.com/cocosip/go-dicom)](https://goreportcard.com/report/github.com/cocosip/go-dicom)
[![GoDoc](https://pkg.go.dev/badge/github.com/cocosip/go-dicom)](https://pkg.go.dev/github.com/cocosip/go-dicom)
```

## Local Development

### Running linters locally

Install golangci-lint:
```bash
# macOS/Linux
brew install golangci-lint

# Windows (with Chocolatey)
choco install golangci-lint

# Or download from https://github.com/golangci/golangci-lint/releases
```

Run linters:
```bash
golangci-lint run
```

### Running tests with coverage

```bash
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
```

### Running benchmarks

```bash
go test -bench=. -benchmem ./pkg/dicom/...
```
