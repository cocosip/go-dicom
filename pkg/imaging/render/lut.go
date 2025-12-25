// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import (
	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
)

// LUT is an alias for the common LUT interface defined in types package.
// This alias is provided for backward compatibility and convenience.
type LUT = imagetypes.LUT

// ModalityLUT defines the interface for Modality lookup tables.
// Either Modality Rescale LUT or Modality Sequence LUT
type ModalityLUT interface {
	LUT
}
