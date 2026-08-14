// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Command generate_dicom regenerates all XML-derived DICOM data files.
package main

import (
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
	root := flags.String("root", "", "repository root receiving generated files")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *standardPath == "" || *privatePath == "" || *root == "" {
		return fmt.Errorf("-standard, -private, and -root are required")
	}

	standard, err := dicomxml.Load(*standardPath)
	if err != nil {
		return err
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
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return fmt.Errorf("create generated output directory for %q: %w", path, err)
	}
	if err := os.WriteFile(path, file.data, 0o644); err != nil { // #nosec G306 -- generated source is repository-readable.
		return fmt.Errorf("write generated file %q: %w", path, err)
	}
	return nil
}
