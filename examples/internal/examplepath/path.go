// Package examplepath provides path validation shared by runnable examples.
package examplepath

import (
	"fmt"
	"os"
	"path/filepath"
)

// RequireInputFile verifies that path identifies an existing non-directory file.
func RequireInputFile(path string) error {
	if path == "" {
		return fmt.Errorf("input file is required")
	}

	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("inspect input file %q: %w", path, err)
	}
	if info.IsDir() {
		return fmt.Errorf("input path %q is a directory; a DICOM file is required", path)
	}
	return nil
}

// PrepareOutputFile creates the parent directory for a file output path.
func PrepareOutputFile(path string) error {
	if path == "" {
		return fmt.Errorf("output file is required")
	}

	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return fmt.Errorf("output path %q is a directory; a file path is required", path)
	} else if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("inspect output file %q: %w", path, err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("create output directory for %q: %w", path, err)
	}
	return nil
}

// PrepareOutputDir creates path unless it already names a non-directory file.
func PrepareOutputDir(path string) error {
	if path == "" {
		return fmt.Errorf("output directory is required")
	}

	if info, err := os.Stat(path); err == nil && !info.IsDir() {
		return fmt.Errorf("output path %q is a file; a directory path is required", path)
	} else if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("inspect output directory %q: %w", path, err)
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("create output directory %q: %w", path, err)
	}
	return nil
}
