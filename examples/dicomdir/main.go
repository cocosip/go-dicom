// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Command dicomdir creates a DICOMDIR from explicitly listed DICOM files.
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/imaging"
	"github.com/cocosip/go-dicom/pkg/media"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	output := flag.String("output", "DICOMDIR", "output DICOMDIR path")
	icons := flag.Bool("icons", false, "generate 128x128 DICOMDIR icons")
	flag.Parse()
	if flag.NArg() == 0 {
		return fmt.Errorf("provide at least one FILE_ID=path input")
	}

	options := make([]media.Option, 0, 2)
	if *icons {
		options = append(options,
			media.WithImageIcons(true),
			media.WithIconGenerator(imaging.NewDirectoryIconGenerator()),
		)
	}
	directory, err := media.NewDirectory(options...)
	if err != nil {
		return err
	}

	for _, input := range flag.Args() {
		fileIDText, path, ok := strings.Cut(input, "=")
		if !ok || fileIDText == "" || path == "" {
			return fmt.Errorf("invalid input %q: want FILE_ID=path", input)
		}
		fileID, err := media.ParseFileID(fileIDText)
		if err != nil {
			return fmt.Errorf("parse File ID %q: %w", fileIDText, err)
		}
		file, err := parser.ParseFile(path)
		if err != nil {
			return fmt.Errorf("parse %s: %w", path, err)
		}
		if _, err := directory.AddFile(file, fileID); err != nil {
			return fmt.Errorf("add %s: %w", path, err)
		}
	}
	if err := directory.Save(*output); err != nil {
		return err
	}

	opened, err := media.Open(*output, media.WithOffsetPolicy(media.StrictOffsets))
	if err != nil {
		return fmt.Errorf("strictly reopen output: %w", err)
	}
	counts := make(map[media.RecordType]int)
	if err := opened.Walk(func(record *media.Record) error {
		counts[record.Type()]++
		return nil
	}); err != nil {
		return err
	}
	types := make([]string, 0, len(counts))
	for recordType := range counts {
		types = append(types, string(recordType))
	}
	sort.Strings(types)
	for _, recordType := range types {
		fmt.Printf("%s=%d\n", recordType, counts[media.RecordType(recordType)])
	}
	return nil
}
