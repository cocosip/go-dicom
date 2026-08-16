// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package scanner_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cocosip/go-dicom/pkg/media/scanner"
)

func ExampleScanner_Scan() {
	root, err := os.MkdirTemp("", "go-dicom-scanner-example")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := os.RemoveAll(root); err != nil {
			panic(err)
		}
	}()
	path := filepath.Join(root, "IMAGE001")
	if err := writeExternalScannerDICOM(path); err != nil {
		panic(err)
	}

	value, err := scanner.New()
	if err != nil {
		panic(err)
	}
	summary, err := value.Scan(context.Background(), []string{path}, func(result scanner.Result) error {
		fmt.Println(result.RelativePath, result.File.IsPartial)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(summary.DICOMFiles)

	// Output:
	// IMAGE001 true
	// 1
}
