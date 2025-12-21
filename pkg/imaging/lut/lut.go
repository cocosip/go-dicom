// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

// Package lut provides Lookup Table implementations for DICOM imaging.
// Lookup Tables (LUTs) are used to map input pixel values to output values
// in DICOM imaging. They are commonly used for windowing, normalization,
// and color mapping.
package lut

import (
	"github.com/cocosip/go-dicom/pkg/imaging/types"
)

// LUT is an alias for the common LUT interface defined in types package.
// This alias is provided for backward compatibility and convenience.
type LUT = types.LUT
