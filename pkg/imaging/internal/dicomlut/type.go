// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dicomlut

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// RequiredLongString reads a required non-empty LO value with VM 1.
func RequiredLongString(ds *dataset.Dataset, t *tag.Tag, name string) (string, error) {
	if ds == nil {
		return "", fmt.Errorf("%s is missing", name)
	}
	valueElement, ok := ds.Get(t)
	if !ok {
		return "", fmt.Errorf("%s is missing", name)
	}
	value, ok := valueElement.(*element.String)
	if !ok || value.ValueRepresentation() != vr.LO {
		return "", fmt.Errorf("%s must use LO VR", name)
	}
	if value.Count() != 1 {
		return "", fmt.Errorf("%s must contain exactly one value", name)
	}
	result := strings.TrimSpace(value.GetValue(0))
	if result == "" {
		return "", fmt.Errorf("%s must not be empty", name)
	}
	return result, nil
}
