// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package scanner

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sync/atomic"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestNewRejectsInvalidOptions(t *testing.T) {
	tests := []struct {
		name   string
		option Option
	}{
		{name: "zero workers", option: WithWorkers(0)},
		{name: "negative workers", option: WithWorkers(-1)},
		{name: "unknown symlink policy", option: WithSymlinkPolicy(SymlinkPolicy(99))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(tt.option); err == nil {
				t.Fatal("New() error = nil, want validation error")
			}
		})
	}
}

func TestNewAcceptsSupportedOptions(t *testing.T) {
	value, err := New(
		WithRecursive(false),
		WithWorkers(2),
		WithStopOnError(true),
		WithSymlinkPolicy(FollowFileSymlinksWithinRoot),
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if value == nil {
		t.Fatal("New() returned nil Scanner")
	}
}

func TestScanRejectsEmptyRoots(t *testing.T) {
	value, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if _, err := value.Scan(context.Background(), nil, nil); err == nil {
		t.Fatal("Scan() error = nil, want empty roots error")
	}
}

func TestScanReturnsPreCancelledContext(t *testing.T) {
	value, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := value.Scan(ctx, []string{"unused"}, nil); !errors.Is(err, context.Canceled) {
		t.Fatalf("Scan() error = %v, want context.Canceled", err)
	}
}

func TestScanErrorUnwrapsResultError(t *testing.T) {
	cause := errors.New("unreadable")
	err := &ScanError{Result: Result{Path: "image.dcm", Kind: ResultReadError, Err: cause}}

	if !errors.Is(err, cause) {
		t.Fatalf("errors.Is(%v, cause) = false", err)
	}
}

func TestScanFileRootReturnsDICOMMetadata(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "image.dcm")
	writeScannerDICOM(t, path, "File^Root")
	value, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	var results []Result
	summary, err := value.Scan(context.Background(), []string{path}, func(result Result) error {
		results = append(results, result)
		return nil
	})
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		t.Fatalf("filepath.Abs() error = %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("results = %d, want 1", len(results))
	}
	result := results[0]
	if result.Root != absPath || result.Path != absPath || result.RelativePath != "image.dcm" {
		t.Fatalf("result paths = root %q path %q relative %q", result.Root, result.Path, result.RelativePath)
	}
	if result.Kind != ResultDICOM || result.File == nil || result.Err != nil {
		t.Fatalf("result = %#v, want DICOM metadata", result)
	}
	if !result.File.IsPartial {
		t.Fatal("Result File IsPartial = false, want true")
	}
	if got := result.File.Dataset.TryGetString(tag.PatientName); got != "File^Root" {
		t.Fatalf("PatientName = %q, want File^Root", got)
	}
	wantSummary := Summary{Results: 1, DICOMFiles: 1}
	if summary != wantSummary {
		t.Fatalf("Summary = %#v, want %#v", summary, wantSummary)
	}
}

func TestScanMixedTreeClassifiesAndOrdersResults(t *testing.T) {
	root := t.TempDir()
	validPath := filepath.Join(root, "a-valid.dcm")
	invalidPath := filepath.Join(root, "b-invalid.txt")
	readErrorPath := filepath.Join(root, "c-read-error.dcm")
	nestedDir := filepath.Join(root, "nested")
	nestedPath := filepath.Join(nestedDir, "d-valid.dcm")
	if err := os.Mkdir(nestedDir, 0o755); err != nil {
		t.Fatalf("Mkdir() error = %v", err)
	}
	writeScannerDICOM(t, validPath, "First^Patient")
	if err := os.WriteFile(invalidPath, []byte("not dicom"), 0o600); err != nil {
		t.Fatalf("WriteFile(invalid) error = %v", err)
	}
	writeScannerDICOM(t, readErrorPath, "Unreadable^Patient")
	writeScannerDICOM(t, nestedPath, "Nested^Patient")

	value, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	realOpen := value.openFile
	value.openFile = func(path string) (*os.File, error) {
		if path == readErrorPath {
			return nil, &fs.PathError{Op: "open", Path: path, Err: fs.ErrPermission}
		}
		return realOpen(path)
	}

	var gotPaths []string
	var gotKinds []ResultKind
	summary, err := value.Scan(context.Background(), []string{root}, func(result Result) error {
		gotPaths = append(gotPaths, result.RelativePath)
		gotKinds = append(gotKinds, result.Kind)
		return nil
	})
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	wantPaths := []string{"a-valid.dcm", "b-invalid.txt", "c-read-error.dcm", filepath.Join("nested", "d-valid.dcm")}
	wantKinds := []ResultKind{ResultDICOM, ResultInvalid, ResultReadError, ResultDICOM}
	if !reflect.DeepEqual(gotPaths, wantPaths) {
		t.Fatalf("paths = %#v, want %#v", gotPaths, wantPaths)
	}
	if !reflect.DeepEqual(gotKinds, wantKinds) {
		t.Fatalf("kinds = %#v, want %#v", gotKinds, wantKinds)
	}
	wantSummary := Summary{Results: 4, DICOMFiles: 2, InvalidFiles: 1, ReadErrors: 1}
	if summary != wantSummary {
		t.Fatalf("Summary = %#v, want %#v", summary, wantSummary)
	}
}

func TestScanNonRecursiveSkipsNestedFiles(t *testing.T) {
	root := t.TempDir()
	writeScannerDICOM(t, filepath.Join(root, "top.dcm"), "Top^Patient")
	nestedDir := filepath.Join(root, "nested")
	if err := os.Mkdir(nestedDir, 0o755); err != nil {
		t.Fatalf("Mkdir() error = %v", err)
	}
	writeScannerDICOM(t, filepath.Join(nestedDir, "nested.dcm"), "Nested^Patient")
	value, err := New(WithRecursive(false))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	var paths []string
	summary, err := value.Scan(context.Background(), []string{root}, func(result Result) error {
		paths = append(paths, result.RelativePath)
		return nil
	})
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	if !reflect.DeepEqual(paths, []string{"top.dcm"}) {
		t.Fatalf("paths = %#v, want top-level file only", paths)
	}
	if summary != (Summary{Results: 1, DICOMFiles: 1}) {
		t.Fatalf("Summary = %#v", summary)
	}
}

func TestScanNilHandlerStillCountsResults(t *testing.T) {
	path := filepath.Join(t.TempDir(), "image.dcm")
	writeScannerDICOM(t, path, "Count^Only")
	value, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	summary, err := value.Scan(nil, []string{path}, nil) //nolint:staticcheck // Scan explicitly accepts a nil context.
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	if summary != (Summary{Results: 1, DICOMFiles: 1}) {
		t.Fatalf("Summary = %#v", summary)
	}
}

func TestScanBoundedByConfiguredWorkers(t *testing.T) {
	root := t.TempDir()
	writePlaceholderFiles(t, root, "a.dcm", "b.dcm", "c.dcm", "d.dcm")
	value, err := New(WithWorkers(2))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	started := make(chan string, 4)
	release := make(chan struct{})
	value.classifyFile = func(ctx context.Context, root, path, relative string) (Result, error) {
		started <- relative
		select {
		case <-release:
			return Result{Root: root, Path: path, RelativePath: relative, Kind: ResultDICOM}, nil
		case <-ctx.Done():
			return Result{}, ctx.Err()
		}
	}

	done := make(chan error, 1)
	go func() {
		_, scanErr := value.Scan(context.Background(), []string{root}, nil)
		done <- scanErr
	}()
	for range 2 {
		select {
		case <-started:
		case <-time.After(2 * time.Second):
			t.Fatal("configured workers did not start concurrently")
		}
	}
	select {
	case extra := <-started:
		t.Fatalf("third classifier %q started while two workers were blocked", extra)
	case <-time.After(50 * time.Millisecond):
	}
	close(release)
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Scan() did not return after workers were released")
	}
}

func TestScanDeterministicOrderAndSerialHandler(t *testing.T) {
	root := t.TempDir()
	writePlaceholderFiles(t, root, "a.dcm", "b.dcm", "c.dcm", "d.dcm")
	value, err := New(WithWorkers(4))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	delays := map[string]time.Duration{
		"a.dcm": 40 * time.Millisecond,
		"b.dcm": 30 * time.Millisecond,
		"c.dcm": 20 * time.Millisecond,
		"d.dcm": 10 * time.Millisecond,
	}
	value.classifyFile = func(_ context.Context, root, path, relative string) (Result, error) {
		time.Sleep(delays[relative])
		return Result{Root: root, Path: path, RelativePath: relative, Kind: ResultDICOM}, nil
	}

	var handlerActive atomic.Int32
	var handlerMaximum atomic.Int32
	var got []string
	summary, err := value.Scan(context.Background(), []string{root}, func(result Result) error {
		active := handlerActive.Add(1)
		for {
			maximum := handlerMaximum.Load()
			if active <= maximum || handlerMaximum.CompareAndSwap(maximum, active) {
				break
			}
		}
		time.Sleep(5 * time.Millisecond)
		got = append(got, result.RelativePath)
		handlerActive.Add(-1)
		return nil
	})
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	want := []string{"a.dcm", "b.dcm", "c.dcm", "d.dcm"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("paths = %#v, want %#v", got, want)
	}
	if handlerMaximum.Load() != 1 {
		t.Fatalf("maximum concurrent handlers = %d, want 1", handlerMaximum.Load())
	}
	if summary != (Summary{Results: 4, DICOMFiles: 4}) {
		t.Fatalf("Summary = %#v", summary)
	}
}

func TestScanCancellationStopsBlockedWorkers(t *testing.T) {
	root := t.TempDir()
	writePlaceholderFiles(t, root, "a.dcm", "b.dcm", "c.dcm")
	value, err := New(WithWorkers(2))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	started := make(chan struct{}, 3)
	value.classifyFile = func(ctx context.Context, _, _, _ string) (Result, error) {
		started <- struct{}{}
		<-ctx.Done()
		return Result{}, ctx.Err()
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		_, scanErr := value.Scan(ctx, []string{root}, nil)
		done <- scanErr
	}()
	for range 2 {
		select {
		case <-started:
		case <-time.After(2 * time.Second):
			t.Fatal("workers did not start")
		}
	}
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("Scan() error = %v, want context.Canceled", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Scan() did not join cancelled workers")
	}
}

func TestScanStopOnErrorReturnsDeterministicPrefix(t *testing.T) {
	root := t.TempDir()
	writePlaceholderFiles(t, root, "a.dcm", "b.dcm", "c.dcm")
	cause := errors.New("invalid fixture")
	classifier := func(_ context.Context, root, path, relative string) (Result, error) {
		result := Result{Root: root, Path: path, RelativePath: relative, Kind: ResultDICOM}
		if relative == "b.dcm" {
			result.Kind = ResultInvalid
			result.Err = cause
		}
		return result, nil
	}

	continueScanner, err := New(WithWorkers(2))
	if err != nil {
		t.Fatalf("New(continue) error = %v", err)
	}
	continueScanner.classifyFile = classifier
	var continued []string
	continueSummary, err := continueScanner.Scan(context.Background(), []string{root}, func(result Result) error {
		continued = append(continued, result.RelativePath)
		return nil
	})
	if err != nil {
		t.Fatalf("continue Scan() error = %v", err)
	}
	if !reflect.DeepEqual(continued, []string{"a.dcm", "b.dcm", "c.dcm"}) {
		t.Fatalf("continued paths = %#v", continued)
	}
	if continueSummary != (Summary{Results: 3, DICOMFiles: 2, InvalidFiles: 1}) {
		t.Fatalf("continue Summary = %#v", continueSummary)
	}

	stopScanner, err := New(WithWorkers(2), WithStopOnError(true))
	if err != nil {
		t.Fatalf("New(stop) error = %v", err)
	}
	stopScanner.classifyFile = classifier
	var stopped []string
	stopSummary, err := stopScanner.Scan(context.Background(), []string{root}, func(result Result) error {
		stopped = append(stopped, result.RelativePath)
		return nil
	})
	if !errors.Is(err, cause) {
		t.Fatalf("stop Scan() error = %v, want cause", err)
	}
	var scanErr *ScanError
	if !errors.As(err, &scanErr) {
		t.Fatalf("stop Scan() error type = %T, want *ScanError", err)
	}
	if !reflect.DeepEqual(stopped, []string{"a.dcm", "b.dcm"}) {
		t.Fatalf("stopped paths = %#v", stopped)
	}
	if stopSummary != (Summary{Results: 2, DICOMFiles: 1, InvalidFiles: 1}) {
		t.Fatalf("stop Summary = %#v", stopSummary)
	}
}

func TestScanHandlerErrorStopsDelivery(t *testing.T) {
	root := t.TempDir()
	writePlaceholderFiles(t, root, "a.dcm", "b.dcm")
	value, err := New(WithWorkers(2))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	value.classifyFile = func(_ context.Context, root, path, relative string) (Result, error) {
		return Result{Root: root, Path: path, RelativePath: relative, Kind: ResultDICOM}, nil
	}
	handlerErr := errors.New("handler stopped")
	var delivered int
	summary, err := value.Scan(context.Background(), []string{root}, func(Result) error {
		delivered++
		return handlerErr
	})
	if !errors.Is(err, handlerErr) {
		t.Fatalf("Scan() error = %v, want handler error", err)
	}
	if delivered != 1 {
		t.Fatalf("delivered = %d, want 1", delivered)
	}
	if summary != (Summary{Results: 1, DICOMFiles: 1}) {
		t.Fatalf("Summary = %#v", summary)
	}
}

func TestScanSymlinkDefaultSkipsLinks(t *testing.T) {
	parent := t.TempDir()
	root := filepath.Join(parent, "root")
	if err := os.Mkdir(root, 0o755); err != nil {
		t.Fatalf("Mkdir() error = %v", err)
	}
	target := filepath.Join(parent, "outside.dcm")
	writeScannerDICOM(t, target, "Outside^Patient")
	requireSymlink(t, target, filepath.Join(root, "linked.dcm"))
	value, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	var results []Result
	summary, err := value.Scan(context.Background(), []string{root}, func(result Result) error {
		results = append(results, result)
		return nil
	})
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("results = %#v, want no delivered links", results)
	}
	if summary != (Summary{SkippedSymlinks: 1}) {
		t.Fatalf("Summary = %#v, want one skipped symlink", summary)
	}
}

func TestScanSymlinkFollowsRegularFileWithinRoot(t *testing.T) {
	root := t.TempDir()
	target := filepath.Join(root, "real.dcm")
	writeScannerDICOM(t, target, "Linked^Patient")
	requireSymlink(t, target, filepath.Join(root, "alias.dcm"))
	value, err := New(WithSymlinkPolicy(FollowFileSymlinksWithinRoot))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	var paths []string
	summary, err := value.Scan(context.Background(), []string{root}, func(result Result) error {
		paths = append(paths, result.RelativePath)
		if result.Kind != ResultDICOM {
			t.Fatalf("result %q kind = %v, want ResultDICOM", result.RelativePath, result.Kind)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	if !reflect.DeepEqual(paths, []string{"alias.dcm", "real.dcm"}) {
		t.Fatalf("paths = %#v, want alias and target in lexical order", paths)
	}
	if summary != (Summary{Results: 2, DICOMFiles: 2}) {
		t.Fatalf("Summary = %#v", summary)
	}
}

func TestScanSymlinkFollowSkipsOutsideAndDirectoryLinks(t *testing.T) {
	parent := t.TempDir()
	root := filepath.Join(parent, "root")
	outsideDir := filepath.Join(parent, "outside-dir")
	if err := os.MkdirAll(root, 0o755); err != nil {
		t.Fatalf("MkdirAll(root) error = %v", err)
	}
	if err := os.MkdirAll(outsideDir, 0o755); err != nil {
		t.Fatalf("MkdirAll(outside) error = %v", err)
	}
	outFile := filepath.Join(parent, "outside.dcm")
	writeScannerDICOM(t, outFile, "Outside^Patient")
	writeScannerDICOM(t, filepath.Join(outsideDir, "nested.dcm"), "Nested^Patient")
	requireSymlink(t, outFile, filepath.Join(root, "external.dcm"))
	requireSymlink(t, outsideDir, filepath.Join(root, "linked-dir"))
	value, err := New(WithSymlinkPolicy(FollowFileSymlinksWithinRoot))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	summary, err := value.Scan(context.Background(), []string{root}, func(result Result) error {
		t.Fatalf("unexpected delivered result: %#v", result)
		return nil
	})
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	if summary != (Summary{SkippedSymlinks: 2}) {
		t.Fatalf("Summary = %#v, want two skipped symlinks", summary)
	}
}

func TestScanSymlinkFollowReportsBrokenLink(t *testing.T) {
	root := t.TempDir()
	link := filepath.Join(root, "broken.dcm")
	requireSymlink(t, filepath.Join(root, "missing.dcm"), link)
	value, err := New(WithSymlinkPolicy(FollowFileSymlinksWithinRoot))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	var results []Result
	summary, err := value.Scan(context.Background(), []string{root}, func(result Result) error {
		results = append(results, result)
		return nil
	})
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("results = %d, want one broken-link result", len(results))
	}
	result := results[0]
	if result.Path != link || result.RelativePath != "broken.dcm" || result.Kind != ResultReadError || result.Err == nil || result.File != nil {
		t.Fatalf("result = %#v, want broken-link read error", result)
	}
	if summary != (Summary{Results: 1, ReadErrors: 1}) {
		t.Fatalf("Summary = %#v", summary)
	}
}

func TestSymlinkJobEnforcesPolicyWithoutHostSymlinkSupport(t *testing.T) {
	parent := t.TempDir()
	root := filepath.Join(parent, "root")
	inside := filepath.Join(root, "inside.dcm")
	outside := filepath.Join(parent, "outside.dcm")
	directory := filepath.Join(root, "directory")
	if err := os.MkdirAll(directory, 0o755); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}
	writePlaceholderFiles(t, root, "inside.dcm")
	writePlaceholderFiles(t, parent, "outside.dcm")
	link := filepath.Join(root, "link.dcm")

	t.Run("default skips without resolving", func(t *testing.T) {
		value, err := New()
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		value.evalSymlinks = func(string) (string, error) {
			t.Fatal("default skip policy resolved a symlink")
			return "", nil
		}
		job := value.symlinkJob(root, root, link, "link.dcm")
		if !job.skippedSymlink || job.result != nil {
			t.Fatalf("job = %#v, want skipped symlink", job)
		}
	})

	tests := []struct {
		name       string
		target     string
		resolveErr error
		wantFollow bool
		wantSkip   bool
		wantRead   bool
	}{
		{name: "regular file within root", target: inside, wantFollow: true},
		{name: "regular file outside root", target: outside, wantSkip: true},
		{name: "directory within root", target: directory, wantSkip: true},
		{name: "broken link", resolveErr: &fs.PathError{Op: "EvalSymlinks", Path: link, Err: fs.ErrNotExist}, wantRead: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := New(WithSymlinkPolicy(FollowFileSymlinksWithinRoot))
			if err != nil {
				t.Fatalf("New() error = %v", err)
			}
			value.evalSymlinks = func(path string) (string, error) {
				if path != link {
					t.Fatalf("EvalSymlinks path = %q, want %q", path, link)
				}
				return tt.target, tt.resolveErr
			}

			job := value.symlinkJob(root, root, link, "link.dcm")
			if tt.wantFollow && (job.path != link || job.root != root || job.relative != "link.dcm" || job.result != nil || job.skippedSymlink) {
				t.Fatalf("job = %#v, want followed regular-file job", job)
			}
			if tt.wantSkip && (!job.skippedSymlink || job.result != nil) {
				t.Fatalf("job = %#v, want skipped symlink", job)
			}
			if tt.wantRead && (job.result == nil || job.result.Kind != ResultReadError || job.result.Err == nil || job.result.Path != link) {
				t.Fatalf("job = %#v, want link read-error result", job)
			}
		})
	}
}

func requireSymlink(t *testing.T, target, link string) {
	t.Helper()
	if err := os.Symlink(target, link); err != nil {
		t.Skipf("symbolic links are unavailable on this host: %v", err)
	}
}

func writePlaceholderFiles(t *testing.T, root string, names ...string) {
	t.Helper()
	for _, name := range names {
		if err := os.WriteFile(filepath.Join(root, name), []byte(name), 0o600); err != nil {
			t.Fatalf("WriteFile(%s) error = %v", name, err)
		}
	}
}

type scannerTestTB interface {
	Helper()
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

func writeScannerDICOM(t scannerTestTB, path, patientName string) {
	t.Helper()
	var file bytes.Buffer
	file.Write(make([]byte, 128))
	file.WriteString("DICM")
	writeScannerShortElement(t, &file, 0x0002, 0x0010, "UI", []byte("1.2.840.10008.1.2.1"))
	writeScannerShortElement(t, &file, 0x0010, 0x0010, "PN", []byte(patientName))
	writeScannerLongElement(t, &file, 0x7FE0, 0x0010, "OB", bytes.Repeat([]byte{0x5A}, 1024))
	if err := os.WriteFile(path, file.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile(%s) error = %v", path, err)
	}
}

func writeScannerShortElement(t scannerTestTB, dst *bytes.Buffer, group, element uint16, vr string, value []byte) {
	t.Helper()
	if err := binary.Write(dst, binary.LittleEndian, group); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(dst, binary.LittleEndian, element); err != nil {
		t.Fatal(err)
	}
	dst.WriteString(vr)
	if err := binary.Write(dst, binary.LittleEndian, uint16(len(value))); err != nil {
		t.Fatal(err)
	}
	dst.Write(value)
}

func writeScannerLongElement(t scannerTestTB, dst *bytes.Buffer, group, element uint16, vr string, value []byte) {
	t.Helper()
	if err := binary.Write(dst, binary.LittleEndian, group); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(dst, binary.LittleEndian, element); err != nil {
		t.Fatal(err)
	}
	dst.WriteString(vr)
	if err := binary.Write(dst, binary.LittleEndian, uint16(0)); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(dst, binary.LittleEndian, uint32(len(value))); err != nil {
		t.Fatal(err)
	}
	dst.Write(value)
}
