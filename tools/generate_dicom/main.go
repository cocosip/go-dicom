// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

// Command generate_dicom regenerates all XML-derived DICOM data files.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cocosip/go-dicom/tools/internal/dicomxml"
)

const (
	tagsOutput              = "pkg/dicom/tag/tags_generated.go"
	uidsOutput              = "pkg/dicom/uid/uids_generated.go"
	dictionaryOutput        = "pkg/dicom/dict/dictionary_data.go"
	privateDictionaryOutput = "pkg/dicom/dict/private_dictionary_data.go"
)

type generatedFile struct {
	path string
	data []byte
}

func main() {
	if err := runCommand(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runCommand(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("generate_dicom", flag.ContinueOnError)
	flags.SetOutput(stdout)
	standardPath := flags.String("standard", "", "path to DICOM Dictionary.xml")
	privatePath := flags.String("private", "", "path to Private Dictionary.xml")
	supplementPath := flags.String("supplement", "", "optional path to supplemental standard dictionary XML rows")
	root := flags.String("root", "", "repository root receiving generated files")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	if *standardPath == "" || *privatePath == "" || *root == "" {
		return fmt.Errorf("-standard, -private, and -root are required")
	}

	standard, err := dicomxml.Load(*standardPath)
	if err != nil {
		return err
	}
	if *supplementPath != "" {
		// The pinned baseline is older than the current PS3.6 edition; the
		// supplemental document carries the registry rows it does not have.
		supplement, err := dicomxml.Load(*supplementPath)
		if err != nil {
			return err
		}
		standard.Tags = append(standard.Tags, supplement.Tags...)
		standard.UIDs = append(standard.UIDs, supplement.UIDs...)
	}
	privateDictionaries, err := dicomxml.LoadPrivate(*privatePath)
	if err != nil {
		return err
	}

	tags, err := generateTags(standard)
	if err != nil {
		return err
	}
	uids, err := generateUIDs(standard)
	if err != nil {
		return err
	}
	dictionary, err := generateDictionary(standard)
	if err != nil {
		return err
	}
	privateDictionary, privateEntryCount, err := generatePrivateDictionary(privateDictionaries)
	if err != nil {
		return err
	}

	files := []generatedFile{
		{path: tagsOutput, data: tags},
		{path: uidsOutput, data: uids},
		{path: dictionaryOutput, data: dictionary},
		{path: privateDictionaryOutput, data: privateDictionary},
	}
	for _, file := range files {
		if err := writeGeneratedFile(*root, file); err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(
		stdout,
		"Generated %d tags, %d UIDs, %d private creators, %d private entries from DICOM %s\n",
		len(standard.Tags),
		len(standard.UIDs),
		len(privateDictionaries.Dictionaries),
		privateEntryCount,
		standard.Version,
	)
	return err
}

func writeGeneratedFile(root string, file generatedFile) error {
	path := filepath.Join(root, filepath.FromSlash(file.path))
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return fmt.Errorf("create generated output directory for %q: %w", path, err)
	}
	// Write through a temporary file in the same directory and rename on
	// success, so an interrupted run cannot leave a truncated generated source.
	temp, err := os.CreateTemp(dir, filepath.Base(path)+".tmp*")
	if err != nil {
		return fmt.Errorf("create temporary file for %q: %w", path, err)
	}
	tempPath := temp.Name()
	removeTemp := func() { _ = os.Remove(tempPath) }

	if _, err := temp.Write(file.data); err != nil {
		_ = temp.Close()
		removeTemp()
		return fmt.Errorf("write generated file %q: %w", path, err)
	}
	if err := temp.Chmod(0o644); err != nil { // #nosec G302 -- generated source is repository-readable.
		_ = temp.Close()
		removeTemp()
		return fmt.Errorf("set generated file permissions %q: %w", path, err)
	}
	if err := temp.Close(); err != nil {
		removeTemp()
		return fmt.Errorf("close generated file %q: %w", path, err)
	}
	if err := os.Rename(tempPath, path); err != nil {
		removeTemp()
		return fmt.Errorf("replace generated file %q: %w", path, err)
	}
	return nil
}
