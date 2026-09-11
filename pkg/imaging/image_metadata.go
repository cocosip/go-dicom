// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/internal/dicomlut"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
)

func imageSoftcopyVOI(presentationState, imageDataset *dataset.Dataset, frame int) (*dataset.Dataset, error) {
	if presentationState == nil || !presentationState.Contains(tag.SoftcopyVOILUTSequence) {
		return nil, nil
	}
	sequence, err := presentationState.GetSequence(tag.SoftcopyVOILUTSequence)
	if err != nil {
		return nil, fmt.Errorf("read Softcopy VOI LUT Sequence: %w", err)
	}
	if sequence.Count() == 0 {
		return nil, fmt.Errorf("softcopy VOI LUT Sequence is empty")
	}
	imageUID, _ := imageDataset.GetString(tag.SOPInstanceUID)
	var defaultItem, matchedItem *dataset.Dataset
	for itemIndex := 0; itemIndex < sequence.Count(); itemIndex++ {
		item := sequence.GetItem(itemIndex)
		if item == nil {
			return nil, fmt.Errorf("softcopy VOI LUT Sequence item %d is nil", itemIndex)
		}
		if !item.Contains(tag.ReferencedImageSequence) {
			if defaultItem != nil {
				return nil, fmt.Errorf("softcopy VOI LUT Sequence contains multiple default items")
			}
			defaultItem = item
			continue
		}
		matches, err := softcopyItemReferencesFrame(item, imageUID, frame)
		if err != nil {
			return nil, fmt.Errorf("softcopy VOI LUT Sequence item %d: %w", itemIndex, err)
		}
		if matches {
			if matchedItem != nil {
				return nil, fmt.Errorf("multiple Softcopy VOI LUT Sequence items reference frame %d", frame+1)
			}
			matchedItem = item
		}
	}
	if matchedItem != nil {
		return matchedItem, nil
	}
	return defaultItem, nil
}

func softcopyItemReferencesFrame(item *dataset.Dataset, imageUID string, frame int) (bool, error) {
	references, err := item.GetSequence(tag.ReferencedImageSequence)
	if err != nil || references.Count() == 0 {
		return false, fmt.Errorf("referenced Image Sequence is missing or empty")
	}
	for index := 0; index < references.Count(); index++ {
		reference := references.GetItem(index)
		if reference == nil {
			return false, fmt.Errorf("referenced Image Sequence item %d is nil", index)
		}
		referencedUID, ok := reference.GetString(tag.ReferencedSOPInstanceUID)
		if !ok || strings.TrimSpace(referencedUID) == "" {
			return false, fmt.Errorf("referenced SOP Instance UID is missing")
		}
		if imageUID == "" || strings.TrimSpace(referencedUID) != strings.TrimSpace(imageUID) {
			continue
		}
		frameNumbers, hasFrames := reference.GetStrings(tag.ReferencedFrameNumber)
		if !hasFrames || len(frameNumbers) == 0 {
			return true, nil
		}
		for _, value := range frameNumbers {
			referencedFrame, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil || referencedFrame < 1 {
				return false, fmt.Errorf("invalid Referenced Frame Number %q", value)
			}
			if referencedFrame == frame+1 {
				return true, nil
			}
		}
	}
	return false, nil
}

func imageWindowPairAt(primary, fallback *dataset.Dataset, index int) (float64, float64, error) {
	source := datasetWithCompletePair(primary, fallback, tag.WindowCenter, tag.WindowWidth)
	if source == nil {
		source = datasetWithAnyTag(primary, fallback, tag.WindowCenter, tag.WindowWidth)
	}
	if source == nil {
		return 0, 0, fmt.Errorf("window center and width are missing")
	}
	centerElement, centerOK := source.Get(tag.WindowCenter)
	widthElement, widthOK := source.Get(tag.WindowWidth)
	if !centerOK || !widthOK {
		return 0, 0, fmt.Errorf("window center and width must both be present")
	}
	centers, centerOK := centerElement.(*element.DecimalString)
	widths, widthOK := widthElement.(*element.DecimalString)
	if !centerOK || !widthOK {
		return 0, 0, fmt.Errorf("window center and width must use Decimal String VR")
	}
	if centers.Count() != widths.Count() {
		return 0, 0, fmt.Errorf("window center and width value counts differ: %d and %d", centers.Count(), widths.Count())
	}
	if index < 0 || index >= centers.Count() {
		return 0, 0, fmt.Errorf("window index %d is out of range [0, %d)", index, centers.Count())
	}
	center, err := centers.GetFloat(index)
	if err != nil {
		return 0, 0, err
	}
	width, err := widths.GetFloat(index)
	if err != nil {
		return 0, 0, err
	}
	return center, width, nil
}

func datasetWithCompletePair(primary, fallback *dataset.Dataset, first, second *tag.Tag) *dataset.Dataset {
	for _, ds := range []*dataset.Dataset{primary, fallback} {
		if ds != nil && ds.Contains(first) && ds.Contains(second) {
			return ds
		}
	}
	return nil
}

func datasetWithAnyTag(primary, fallback *dataset.Dataset, tags ...*tag.Tag) *dataset.Dataset {
	for _, ds := range []*dataset.Dataset{primary, fallback} {
		if ds == nil {
			continue
		}
		for _, t := range tags {
			if ds.Contains(t) {
				return ds
			}
		}
	}
	return nil
}

func imageModalityTransform(primary, fallback *dataset.Dataset, pixelSigned bool, minInput, maxInput float64) (
	modalityLUT lut.LUT,
	slope float64,
	intercept float64,
	voiDescriptorSigned bool,
	err error,
) {
	slope, intercept = 1, 0
	// Modality attributes are selected as one precedence group. A Functional
	// Group value must override top-level values, including the other legal
	// representation of the modality transform.
	source := datasetWithAnyTag(
		primary,
		nil,
		tag.ModalityLUTSequence,
		tag.RescaleSlope,
		tag.RescaleIntercept,
	)
	if source == nil {
		source = datasetWithAnyTag(
			fallback,
			nil,
			tag.ModalityLUTSequence,
			tag.RescaleSlope,
			tag.RescaleIntercept,
		)
	}
	if source == nil {
		return nil, slope, intercept, pixelSigned, nil
	}
	modalitySource := source
	rescaleSource := source
	if !source.Contains(tag.ModalityLUTSequence) {
		modalitySource = nil
	}
	if !source.Contains(tag.RescaleSlope) && !source.Contains(tag.RescaleIntercept) {
		rescaleSource = nil
	}
	if modalitySource != nil && rescaleSource != nil {
		return nil, 0, 0, false, fmt.Errorf("modality LUT sequence cannot coexist with rescale slope/intercept")
	}
	if modalitySource != nil {
		modalityLUT, err = pixeldata.ModalityLUT(modalitySource, pixelSigned)
		if err != nil {
			return nil, 0, 0, false, err
		}
		return modalityLUT, slope, intercept, false, nil
	}
	if rescaleSource == nil {
		return nil, slope, intercept, pixelSigned, nil
	}
	if !rescaleSource.Contains(tag.RescaleSlope) || !rescaleSource.Contains(tag.RescaleIntercept) {
		return nil, 0, 0, false, fmt.Errorf("rescale slope and intercept must both be present")
	}
	if _, err := dicomlut.RequiredLongString(rescaleSource, tag.RescaleType, "Rescale Type"); err != nil {
		return nil, 0, 0, false, err
	}
	slope, err = imageSingleDecimal(rescaleSource, tag.RescaleSlope)
	if err != nil {
		return nil, 0, 0, false, err
	}
	if slope == 0 || math.IsNaN(slope) || math.IsInf(slope, 0) {
		return nil, 0, 0, false, fmt.Errorf("rescale slope must be finite and non-zero")
	}
	intercept, err = imageSingleDecimal(rescaleSource, tag.RescaleIntercept)
	if err != nil {
		return nil, 0, 0, false, err
	}
	if math.IsNaN(intercept) || math.IsInf(intercept, 0) {
		return nil, 0, 0, false, fmt.Errorf("rescale intercept must be finite")
	}
	minimum := math.Min(minInput*slope+intercept, maxInput*slope+intercept)
	return nil, slope, intercept, minimum < 0, nil
}

func imageSingleDecimal(ds *dataset.Dataset, t *tag.Tag) (float64, error) {
	elem, ok := ds.Get(t)
	if !ok {
		return 0, fmt.Errorf("element %s not found", t)
	}
	value, ok := elem.(*element.DecimalString)
	if !ok || value.Count() != 1 {
		return 0, fmt.Errorf("element %s must contain exactly one Decimal String value", t)
	}
	return value.GetFloat(0)
}

func imageStringFrom(primary, fallback *dataset.Dataset, t *tag.Tag) (string, bool) {
	for _, source := range []*dataset.Dataset{primary, fallback} {
		if value, ok := source.GetString(t); ok {
			value = strings.TrimSpace(value)
			if value != "" {
				return value, true
			}
		}
	}
	return "", false
}

// imageFunctionalGroupValues flattens the first item of each shared and
// per-frame functional-group macro. Per-frame values replace shared values.
func imageFunctionalGroupValues(ds *dataset.Dataset, frame int) *dataset.Dataset {
	return dicomlut.FunctionalGroupValues(ds, frame)
}
