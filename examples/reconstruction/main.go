// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Command reconstruction generates a classic CT/MR MPR series from one
// Enhanced dataset or a list of classic source instances.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/imaging/reconstruction"
)

func main() {
	plane := flag.String("plane", "coronal", "output plane: axial, coronal, or sagittal")
	spacing := flag.Float64("spacing", 1, "output in-plane pixel spacing in mm")
	sliceDistance := flag.Float64("slice-distance", 1, "output slice distance in mm")
	outputDirectory := flag.String("output", "mpr-output", "output directory")
	workers := flag.Int("workers", 1, "parallel workers per output slice")
	allowIrregular := flag.Bool("allow-irregular", false, "allow irregular source slice spacing")
	flag.Parse()
	if err := run(context.Background(), flag.Args(), *outputDirectory, *plane, *spacing, *sliceDistance, *workers, *allowIrregular); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, inputs []string, outputDirectory, plane string, spacing, sliceDistance float64, workers int, allowIrregular bool) error {
	if len(inputs) == 0 {
		return fmt.Errorf("provide one Enhanced DICOM file or at least two classic DICOM files")
	}
	stackType, err := parseStackType(plane)
	if err != nil {
		return err
	}
	images := make([]*reconstruction.ImageData, 0, len(inputs))
	for _, input := range inputs {
		result, err := parser.ParseFile(input)
		if err != nil {
			return fmt.Errorf("parse %s: %w", input, err)
		}
		frames, err := reconstruction.NewImageDataFromDataset(result.Dataset)
		if err != nil {
			return fmt.Errorf("read reconstruction frames from %s: %w", input, err)
		}
		images = append(images, frames...)
	}
	volumeOptions := []reconstruction.VolumeOption{}
	if allowIrregular {
		volumeOptions = append(volumeOptions, reconstruction.WithIrregularSpacingAllowed())
	}
	volume, err := reconstruction.NewVolumeData(images, volumeOptions...)
	if err != nil {
		return err
	}
	stack, err := reconstruction.NewStack(volume, stackType, spacing, sliceDistance)
	if err != nil {
		return err
	}
	generator, err := reconstruction.NewDicomGenerator(volume)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(outputDirectory, 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	description := stackType.String() + " MPR"
	return generator.Stream(ctx, stack, description, reconstruction.CutOptions{Workers: workers}, func(index int, output *dataset.Dataset) error {
		path := filepath.Join(outputDirectory, fmt.Sprintf("mpr-%04d.dcm", index+1))
		if err := writer.WriteFile(path, output, writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
}

func parseStackType(value string) (reconstruction.StackType, error) {
	switch strings.ToLower(value) {
	case "axial":
		return reconstruction.StackTypeAxial, nil
	case "coronal":
		return reconstruction.StackTypeCoronal, nil
	case "sagittal":
		return reconstruction.StackTypeSagittal, nil
	default:
		return 0, fmt.Errorf("unknown plane %q", value)
	}
}
