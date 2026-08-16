// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package scanner

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
)

// Scanner is an immutable filesystem scanner configuration.
type Scanner struct {
	config       config
	openFile     func(string) (*os.File, error)
	evalSymlinks func(string) (string, error)
	classifyFile func(context.Context, string, string, string) (Result, error)
}

type scanJob struct {
	sequence          int
	root              string
	path              string
	relative          string
	result            *Result
	skippedSymlink    bool
	skippedNonRegular bool
}

type scanCompletion struct {
	job    scanJob
	result Result
	err    error
}

// New constructs a Scanner from options.
func New(options ...Option) (*Scanner, error) {
	value, err := newConfig(options...)
	if err != nil {
		return nil, err
	}
	return &Scanner{
		config:       value,
		openFile:     os.Open,
		evalSymlinks: filepath.EvalSymlinks,
	}, nil
}

// Scan scans file and directory roots and delivers metadata results in discovery order.
func (s *Scanner) Scan(ctx context.Context, roots []string, handler Handler) (Summary, error) {
	if s == nil {
		return Summary{}, fmt.Errorf("scanner cannot be nil")
	}
	if len(roots) == 0 {
		return Summary{}, fmt.Errorf("scanner roots cannot be empty")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return Summary{}, err
	}

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	jobs := make(chan scanJob)
	enumerationDone := make(chan error, 1)
	outstanding := make(chan struct{}, s.config.workers*2)

	go func() {
		defer close(jobs)
		sequence := 0
		submit := func(job scanJob) error {
			select {
			case outstanding <- struct{}{}:
			case <-runCtx.Done():
				return runCtx.Err()
			}
			job.sequence = sequence
			sequence++
			select {
			case jobs <- job:
				return nil
			case <-runCtx.Done():
				<-outstanding
				return runCtx.Err()
			}
		}
		enumerationDone <- s.enumerate(runCtx, roots, submit)
	}()
	completions, workersDone := s.startWorkers(runCtx, jobs)

	var summary Summary
	pending := make(map[int]scanCompletion)
	next := 0
	for {
		select {
		case <-ctx.Done():
			cancel()
			<-enumerationDone
			<-workersDone
			return summary, ctx.Err()
		case completion, ok := <-completions:
			if !ok {
				enumerationErr := <-enumerationDone
				if enumerationErr != nil {
					if err := ctx.Err(); err != nil {
						return summary, err
					}
					return summary, enumerationErr
				}
				return summary, nil
			}
			pending[completion.job.sequence] = completion
			for {
				current, exists := pending[next]
				if !exists {
					break
				}
				delete(pending, next)
				next++
				<-outstanding
				if current.err != nil {
					cancel()
					<-enumerationDone
					<-workersDone
					if err := ctx.Err(); err != nil {
						return summary, err
					}
					return summary, current.err
				}
				if current.job.skippedSymlink {
					summary.SkippedSymlinks++
					continue
				}
				if current.job.skippedNonRegular {
					summary.SkippedNonRegular++
					continue
				}
				if err := deliver(&summary, current.result, handler); err != nil {
					cancel()
					<-enumerationDone
					<-workersDone
					return summary, err
				}
				if s.config.stopOnError && current.result.Kind != ResultDICOM {
					cancel()
					<-enumerationDone
					<-workersDone
					return summary, &ScanError{Result: current.result}
				}
			}
		}
	}
}

func (s *Scanner) startWorkers(ctx context.Context, jobs <-chan scanJob) (<-chan scanCompletion, <-chan struct{}) {
	completions := make(chan scanCompletion)
	done := make(chan struct{})
	classifier := s.classifyFile
	if classifier == nil {
		classifier = s.scanFile
	}
	var workers sync.WaitGroup
	workers.Add(s.config.workers)
	for range s.config.workers {
		go func() {
			defer workers.Done()
			for job := range jobs {
				completion := scanCompletion{job: job}
				switch {
				case job.result != nil:
					completion.result = *job.result
				case job.skippedSymlink, job.skippedNonRegular:
				default:
					completion.result, completion.err = classifier(ctx, job.root, job.path, job.relative)
				}
				select {
				case completions <- completion:
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	go func() {
		workers.Wait()
		close(completions)
		close(done)
	}()
	return completions, done
}

func (s *Scanner) enumerate(ctx context.Context, roots []string, submit func(scanJob) error) error {
	for _, root := range roots {
		if err := ctx.Err(); err != nil {
			return err
		}
		absoluteRoot, err := filepath.Abs(root)
		if err != nil {
			return fmt.Errorf("resolve scanner root %q: %w", root, err)
		}
		absoluteRoot = filepath.Clean(absoluteRoot)
		info, err := os.Lstat(absoluteRoot)
		if err != nil {
			result := Result{Root: absoluteRoot, Path: absoluteRoot, RelativePath: ".", Kind: ResultReadError, Err: err}
			if err := submit(scanJob{result: &result}); err != nil {
				return err
			}
			continue
		}
		if info.Mode()&os.ModeSymlink != 0 {
			canonicalRoot := ""
			if s.config.symlinkPolicy == FollowFileSymlinksWithinRoot {
				canonicalRoot, err = s.evalSymlinks(filepath.Dir(absoluteRoot))
				if err != nil {
					result := Result{Root: absoluteRoot, Path: absoluteRoot, RelativePath: filepath.Base(absoluteRoot), Kind: ResultReadError, Err: err}
					if submitErr := submit(scanJob{result: &result}); submitErr != nil {
						return submitErr
					}
					continue
				}
			}
			job := s.symlinkJob(absoluteRoot, canonicalRoot, absoluteRoot, filepath.Base(absoluteRoot))
			if err := submit(job); err != nil {
				return err
			}
			continue
		}
		if info.IsDir() {
			if err := s.enumerateDirectory(ctx, absoluteRoot, submit); err != nil {
				return err
			}
			continue
		}
		if !info.Mode().IsRegular() {
			if err := submit(scanJob{skippedNonRegular: true}); err != nil {
				return err
			}
			continue
		}
		if err := submit(scanJob{root: absoluteRoot, path: absoluteRoot, relative: filepath.Base(absoluteRoot)}); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scanner) enumerateDirectory(ctx context.Context, root string, submit func(scanJob) error) error {
	canonicalRoot := root
	if s.config.symlinkPolicy == FollowFileSymlinksWithinRoot {
		resolved, err := s.evalSymlinks(root)
		if err != nil {
			result := Result{Root: root, Path: root, RelativePath: ".", Kind: ResultReadError, Err: err}
			return submit(scanJob{result: &result})
		}
		canonicalRoot = resolved
	}
	return filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if err := ctx.Err(); err != nil {
			return err
		}
		relative, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return fmt.Errorf("make scanner path relative to %s: %w", root, relErr)
		}
		if walkErr != nil {
			result := Result{Root: root, Path: path, RelativePath: relative, Kind: ResultReadError, Err: walkErr}
			return submit(scanJob{result: &result})
		}
		if path == root {
			return nil
		}
		if entry.IsDir() {
			if !s.config.recursive {
				return filepath.SkipDir
			}
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return submit(s.symlinkJob(root, canonicalRoot, path, relative))
		}
		info, err := entry.Info()
		if err != nil {
			result := Result{Root: root, Path: path, RelativePath: relative, Kind: ResultReadError, Err: err}
			return submit(scanJob{result: &result})
		}
		if !info.Mode().IsRegular() {
			return submit(scanJob{skippedNonRegular: true})
		}
		return submit(scanJob{root: root, path: path, relative: relative})
	})
}

func (s *Scanner) symlinkJob(root, canonicalRoot, path, relative string) scanJob {
	if s.config.symlinkPolicy == SkipSymlinks {
		return scanJob{skippedSymlink: true}
	}
	target, err := s.evalSymlinks(path)
	if err != nil {
		result := Result{Root: root, Path: path, RelativePath: relative, Kind: ResultReadError, Err: err}
		return scanJob{result: &result}
	}
	withinRoot, err := pathWithinRoot(canonicalRoot, target)
	if err != nil || !withinRoot {
		return scanJob{skippedSymlink: true}
	}
	info, err := os.Stat(target)
	if err != nil {
		result := Result{Root: root, Path: path, RelativePath: relative, Kind: ResultReadError, Err: err}
		return scanJob{result: &result}
	}
	if !info.Mode().IsRegular() {
		return scanJob{skippedSymlink: true}
	}
	return scanJob{root: root, path: path, relative: relative}
}

func pathWithinRoot(root, path string) (bool, error) {
	relative, err := filepath.Rel(filepath.Clean(root), filepath.Clean(path))
	if err != nil {
		return false, err
	}
	return relative != ".." && !filepath.IsAbs(relative) && !strings.HasPrefix(relative, ".."+string(filepath.Separator)), nil
}

func (s *Scanner) scanFile(ctx context.Context, root, path, relative string) (Result, error) {
	result := Result{Root: root, Path: path, RelativePath: relative}
	file, err := s.openFile(path)
	if err != nil {
		result.Kind = ResultReadError
		result.Err = err
		return result, nil
	}
	options := []parser.Option{
		parser.WithContext(ctx),
		parser.WithStopBeforePixelData(),
	}
	if s.config.assumedTransferSyntax != nil {
		options = append(options, parser.WithAssumedTransferSyntax(s.config.assumedTransferSyntax))
	}
	parsed, parseErr := parser.Parse(file, options...)
	closeErr := file.Close()
	if err := ctx.Err(); err != nil {
		return Result{}, err
	}
	if parseErr != nil {
		result.Kind = ResultInvalid
		result.Err = parseErr
		var pathErr *fs.PathError
		if errors.As(parseErr, &pathErr) {
			result.Kind = ResultReadError
		}
		return result, nil
	}
	if closeErr != nil {
		result.Kind = ResultReadError
		result.Err = closeErr
		return result, nil
	}
	result.Kind = ResultDICOM
	result.File = parsed
	return result, nil
}

func deliver(summary *Summary, result Result, handler Handler) error {
	summary.Results++
	switch result.Kind {
	case ResultDICOM:
		summary.DICOMFiles++
	case ResultInvalid:
		summary.InvalidFiles++
	case ResultReadError:
		summary.ReadErrors++
	}
	if handler == nil {
		return nil
	}
	return handler(result)
}
