// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"fmt"
	"strconv"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

const (
	basicFilmSessionSOPClassUID = "1.2.840.10008.5.1.1.1"
	basicFilmBoxSOPClassUID     = "1.2.840.10008.5.1.1.2"
	presentationLUTSOPClassUID  = "1.2.840.10008.5.1.1.23"
)

func newSOPInstanceUID() string {
	return uid.GenerateDerivedFromUUID().UID()
}

func addElement(ds *dataset.Dataset, elem element.Element) error {
	if err := ds.Add(elem); err != nil {
		return fmt.Errorf("add %s: %w", elem.Tag(), err)
	}
	return nil
}

func addString(ds *dataset.Dataset, t *tag.Tag, value string, valueRepresentation *vr.VR) error {
	if value == "" {
		return nil
	}
	return addElement(ds, element.NewString(t, valueRepresentation, []string{value}))
}

func requiredString(ds *dataset.Dataset, t *tag.Tag, name string) (string, error) {
	if ds == nil {
		return "", fmt.Errorf("printing: nil Dataset")
	}
	value, ok := ds.GetString(t)
	if !ok || value == "" {
		return "", fmt.Errorf("printing: Dataset has no %s", name)
	}
	return value, nil
}

func overrideOrDatasetUID(override string, ds *dataset.Dataset, t *tag.Tag, name string) (string, error) {
	if override != "" {
		return override, nil
	}
	return requiredString(ds, t, name)
}

func optionalInteger(ds *dataset.Dataset, t *tag.Tag) (int, bool, error) {
	elem, ok := ds.Get(t)
	if !ok {
		return 0, false, nil
	}
	var raw string
	switch value := elem.(type) {
	case *element.IntegerString:
		raw = value.GetValue(0)
	case *element.String:
		if value.ValueRepresentation() != vr.IS {
			return 0, true, fmt.Errorf("printing: %s has VR %s, want IS", t, value.ValueRepresentation())
		}
		raw = value.GetValue(0)
	default:
		return 0, true, fmt.Errorf("printing: %s is not an Integer String", t)
	}
	parsed, err := strconv.Atoi(raw)
	if err != nil {
		return 0, true, fmt.Errorf("printing: invalid %s: %w", t, err)
	}
	return parsed, true, nil
}

func optionalDecimalString(ds *dataset.Dataset, t *tag.Tag) (string, bool, error) {
	elem, ok := ds.Get(t)
	if !ok {
		return "", false, nil
	}
	switch value := elem.(type) {
	case *element.DecimalString:
		return value.GetValue(0), true, nil
	case *element.String:
		if value.ValueRepresentation() != vr.DS {
			return "", true, fmt.Errorf("printing: %s has VR %s, want DS", t, value.ValueRepresentation())
		}
		return value.GetValue(0), true, nil
	default:
		return "", true, fmt.Errorf("printing: %s is not a Decimal String", t)
	}
}

func optionalUInt16(ds *dataset.Dataset, t *tag.Tag) (uint16, bool, error) {
	elem, ok := ds.Get(t)
	if !ok {
		return 0, false, nil
	}
	if elem.Count() != 1 {
		return 0, true, fmt.Errorf("printing: %s has VM %d, want 1", t, elem.Count())
	}
	value, err := ds.GetUInt16(t, 0)
	if err != nil {
		return 0, true, fmt.Errorf("printing: invalid %s: %w", t, err)
	}
	return value, true, nil
}

func addReferenceSequence(ds *dataset.Dataset, sequenceTag *tag.Tag, references []SOPReference) error {
	if len(references) == 0 {
		return nil
	}
	items := make([]*dataset.Dataset, 0, len(references))
	for index, reference := range references {
		if reference.SOPClassUID == "" || reference.SOPInstanceUID == "" {
			return fmt.Errorf("printing: %s reference %d requires SOP Class and Instance UIDs", sequenceTag, index)
		}
		item := dataset.New()
		if err := addString(item, tag.ReferencedSOPClassUID, reference.SOPClassUID, vr.UI); err != nil {
			return err
		}
		if err := addString(item, tag.ReferencedSOPInstanceUID, reference.SOPInstanceUID, vr.UI); err != nil {
			return err
		}
		items = append(items, item)
	}
	return addElement(ds, dataset.NewSequenceWithItems(sequenceTag, items))
}

func readReferenceSequence(ds *dataset.Dataset, sequenceTag *tag.Tag) ([]SOPReference, error) {
	sequence, err := ds.GetSequence(sequenceTag)
	if err != nil {
		if _, ok := ds.Get(sequenceTag); !ok {
			return nil, nil
		}
		return nil, fmt.Errorf("printing: read %s: %w", sequenceTag, err)
	}
	references := make([]SOPReference, 0, sequence.Count())
	for index := 0; index < sequence.Count(); index++ {
		item := sequence.GetItem(index)
		if item == nil {
			return nil, fmt.Errorf("printing: %s item %d is nil", sequenceTag, index)
		}
		classUID, err := requiredString(item, tag.ReferencedSOPClassUID, "Referenced SOP Class UID")
		if err != nil {
			return nil, fmt.Errorf("printing: %s item %d: %w", sequenceTag, index, err)
		}
		instanceUID, err := requiredString(item, tag.ReferencedSOPInstanceUID, "Referenced SOP Instance UID")
		if err != nil {
			return nil, fmt.Errorf("printing: %s item %d: %w", sequenceTag, index, err)
		}
		references = append(references, SOPReference{SOPClassUID: classUID, SOPInstanceUID: instanceUID})
	}
	return references, nil
}

// ToDataset returns an independent Dataset containing the Film Session attributes.
func (fs *FilmSession) ToDataset() (*dataset.Dataset, error) {
	if fs == nil {
		return nil, fmt.Errorf("printing: nil FilmSession")
	}
	if fs.SOPClassUID == "" || fs.SOPInstanceUID == "" {
		return nil, fmt.Errorf("printing: FilmSession requires SOP Class and Instance UIDs")
	}

	ds := dataset.New()
	values := []struct {
		tag   *tag.Tag
		value string
		vr    *vr.VR
	}{
		{tag.SOPClassUID, fs.SOPClassUID, vr.UI},
		{tag.SOPInstanceUID, fs.SOPInstanceUID, vr.UI},
		{tag.FilmDestination, string(fs.FilmDestination), vr.CS},
		{tag.FilmSessionLabel, fs.FilmSessionLabel, vr.LO},
		{tag.MediumType, string(fs.MediumType), vr.CS},
		{tag.PrintPriority, string(fs.PrintPriority), vr.CS},
	}
	for _, value := range values {
		if err := addString(ds, value.tag, value.value, value.vr); err != nil {
			return nil, err
		}
	}
	if err := addElement(ds, element.NewIntegerStringFromInt(tag.MemoryAllocation, []int{fs.MemoryAllocation})); err != nil {
		return nil, err
	}
	if err := addElement(ds, element.NewIntegerStringFromInt(tag.NumberOfCopies, []int{fs.NumberOfCopies})); err != nil {
		return nil, err
	}
	return ds, nil
}

// NewFilmSessionFromDataset constructs a Film Session from an independent copy of Dataset values.
func NewFilmSessionFromDataset(sopClassUID, sopInstanceUID string, ds *dataset.Dataset, isColor bool) (*FilmSession, error) {
	classUID, err := overrideOrDatasetUID(sopClassUID, ds, tag.SOPClassUID, "SOP Class UID")
	if err != nil {
		return nil, err
	}
	instanceUID, err := overrideOrDatasetUID(sopInstanceUID, ds, tag.SOPInstanceUID, "SOP Instance UID")
	if err != nil {
		return nil, err
	}

	fs := NewFilmSession(classUID, instanceUID, isColor)
	if value, ok := ds.GetString(tag.FilmDestination); ok {
		fs.FilmDestination = FilmDestination(value)
	}
	if value, ok := ds.GetString(tag.FilmSessionLabel); ok {
		fs.FilmSessionLabel = value
	}
	if value, ok, err := optionalInteger(ds, tag.MemoryAllocation); err != nil {
		return nil, err
	} else if ok {
		fs.MemoryAllocation = value
	}
	if value, ok := ds.GetString(tag.MediumType); ok {
		fs.MediumType = MediumType(value)
	}
	if value, ok := ds.GetString(tag.PrintPriority); ok {
		fs.PrintPriority = PrintPriority(value)
	}
	if value, ok, err := optionalInteger(ds, tag.NumberOfCopies); err != nil {
		return nil, err
	} else if ok {
		fs.NumberOfCopies = value
	}
	return fs, nil
}

// ToDataset returns an independent Dataset containing the Film Box attributes.
func (fb *FilmBox) ToDataset() (*dataset.Dataset, error) {
	if fb == nil {
		return nil, fmt.Errorf("printing: nil FilmBox")
	}
	if fb.SOPInstanceUID == "" {
		return nil, fmt.Errorf("printing: FilmBox requires an SOP Instance UID")
	}

	ds := dataset.New()
	values := []struct {
		tag   *tag.Tag
		value string
		vr    *vr.VR
	}{
		{tag.SOPClassUID, basicFilmBoxSOPClassUID, vr.UI},
		{tag.SOPInstanceUID, fb.SOPInstanceUID, vr.UI},
		{tag.ImageDisplayFormat, fb.ImageDisplayFormat, vr.ST},
		{tag.FilmOrientation, string(fb.FilmOrientation), vr.CS},
		{tag.FilmSizeID, string(fb.FilmSizeID), vr.CS},
		{tag.MagnificationType, string(fb.MagnificationType), vr.CS},
		{tag.Trim, string(fb.Trim), vr.CS},
		{tag.BorderDensity, string(fb.BorderDensity), vr.CS},
		{tag.EmptyImageDensity, string(fb.EmptyImageDensity), vr.CS},
		{tag.ConfigurationInformation, fb.ConfigurationInformation, vr.ST},
		{tag.AnnotationDisplayFormatID, fb.AnnotationDisplayFormatID, vr.CS},
		{tag.SmoothingType, fb.SmoothingType, vr.CS},
		{tag.RequestedResolutionID, fb.RequestedResolutionID, vr.CS},
	}
	for _, value := range values {
		if err := addString(ds, value.tag, value.value, value.vr); err != nil {
			return nil, err
		}
	}
	if err := addElement(ds, element.NewUnsignedShort(tag.MaxDensity, []uint16{fb.MaxDensity})); err != nil {
		return nil, err
	}
	if err := addElement(ds, element.NewUnsignedShort(tag.MinDensity, []uint16{fb.MinDensity})); err != nil {
		return nil, err
	}
	if err := addElement(ds, element.NewUnsignedShort(tag.Illumination, []uint16{fb.Illumination})); err != nil {
		return nil, err
	}
	if err := addElement(ds, element.NewUnsignedShort(tag.ReflectedAmbientLight, []uint16{fb.ReflectedAmbientLight})); err != nil {
		return nil, err
	}
	sessionReference := fb.ReferencedFilmSession
	if fb.filmSession != nil {
		sessionReference = SOPReference{SOPClassUID: fb.filmSession.SOPClassUID, SOPInstanceUID: fb.filmSession.SOPInstanceUID}
	}
	if sessionReference != (SOPReference{}) {
		if err := addReferenceSequence(ds, tag.ReferencedFilmSessionSequence, []SOPReference{sessionReference}); err != nil {
			return nil, err
		}
	}
	imageReferences := append([]SOPReference(nil), fb.ReferencedImageBoxes...)
	if len(fb.BasicImageBoxes) > 0 {
		imageReferences = make([]SOPReference, 0, len(fb.BasicImageBoxes))
		for index, imageBox := range fb.BasicImageBoxes {
			if imageBox == nil {
				return nil, fmt.Errorf("printing: nil ImageBox at index %d", index)
			}
			imageReferences = append(imageReferences, SOPReference{
				SOPClassUID: imageBox.SOPClassUID, SOPInstanceUID: imageBox.SOPInstanceUID,
			})
		}
	}
	if err := addReferenceSequence(ds, tag.ReferencedImageBoxSequence, imageReferences); err != nil {
		return nil, err
	}
	if fb.ReferencedPresentationLUT != (SOPReference{}) {
		if err := addReferenceSequence(ds, tag.ReferencedPresentationLUTSequence, []SOPReference{fb.ReferencedPresentationLUT}); err != nil {
			return nil, err
		}
	}
	return ds, nil
}

// NewFilmBoxFromDataset constructs a Film Box from Dataset values.
func NewFilmBoxFromDataset(sopInstanceUID string, ds *dataset.Dataset) (*FilmBox, error) {
	if _, err := requiredString(ds, tag.SOPClassUID, "SOP Class UID"); err != nil {
		return nil, err
	}
	instanceUID, err := overrideOrDatasetUID(sopInstanceUID, ds, tag.SOPInstanceUID, "SOP Instance UID")
	if err != nil {
		return nil, err
	}
	format, _ := ds.GetString(tag.ImageDisplayFormat)
	fb := NewFilmBox(instanceUID, format)
	if value, ok := ds.GetString(tag.FilmOrientation); ok {
		fb.FilmOrientation = FilmOrientation(value)
	}
	if value, ok := ds.GetString(tag.FilmSizeID); ok {
		fb.FilmSizeID = FilmSize(value)
	}
	if value, ok := ds.GetString(tag.MagnificationType); ok {
		fb.MagnificationType = MagnificationType(value)
	}
	if value, err := ds.GetUInt16(tag.MaxDensity, 0); err == nil {
		fb.MaxDensity = value
	}
	if value, err := ds.GetUInt16(tag.MinDensity, 0); err == nil {
		fb.MinDensity = value
	}
	if value, ok := ds.GetString(tag.Trim); ok {
		fb.Trim = TrimMode(value)
	}
	if value, ok := ds.GetString(tag.BorderDensity); ok {
		fb.BorderDensity = BorderDensity(value)
	}
	if value, ok := ds.GetString(tag.EmptyImageDensity); ok {
		fb.EmptyImageDensity = EmptyImageDensity(value)
	}
	if value, ok := ds.GetString(tag.ConfigurationInformation); ok {
		fb.ConfigurationInformation = value
	}
	if value, ok := ds.GetString(tag.AnnotationDisplayFormatID); ok {
		fb.AnnotationDisplayFormatID = value
	}
	if value, ok := ds.GetString(tag.SmoothingType); ok {
		fb.SmoothingType = value
	}
	if value, ok, err := optionalUInt16(ds, tag.Illumination); err != nil {
		return nil, err
	} else if ok {
		fb.Illumination = value
	}
	if value, ok, err := optionalUInt16(ds, tag.ReflectedAmbientLight); err != nil {
		return nil, err
	} else if ok {
		fb.ReflectedAmbientLight = value
	}
	if value, ok := ds.GetString(tag.RequestedResolutionID); ok {
		fb.RequestedResolutionID = value
	}
	if references, err := readReferenceSequence(ds, tag.ReferencedFilmSessionSequence); err != nil {
		return nil, err
	} else if len(references) > 1 {
		return nil, fmt.Errorf("printing: Referenced Film Session Sequence has %d items, want at most 1", len(references))
	} else if len(references) == 1 {
		fb.ReferencedFilmSession = references[0]
	}
	references, err := readReferenceSequence(ds, tag.ReferencedImageBoxSequence)
	if err != nil {
		return nil, err
	}
	fb.ReferencedImageBoxes = references
	if references, err := readReferenceSequence(ds, tag.ReferencedPresentationLUTSequence); err != nil {
		return nil, err
	} else if len(references) > 1 {
		return nil, fmt.Errorf("printing: Referenced Presentation LUT Sequence has %d items, want at most 1", len(references))
	} else if len(references) == 1 {
		fb.ReferencedPresentationLUT = references[0]
	}
	return fb, nil
}

// ToDataset returns an independent Dataset containing the Image Box attributes.
func (ib *ImageBox) ToDataset() (*dataset.Dataset, error) {
	if ib == nil {
		return nil, fmt.Errorf("printing: nil ImageBox")
	}
	if ib.SOPClassUID == "" || ib.SOPInstanceUID == "" {
		return nil, fmt.Errorf("printing: ImageBox requires SOP Class and Instance UIDs")
	}

	ds := dataset.New()
	values := []struct {
		tag   *tag.Tag
		value string
		vr    *vr.VR
	}{
		{tag.SOPClassUID, ib.SOPClassUID, vr.UI},
		{tag.SOPInstanceUID, ib.SOPInstanceUID, vr.UI},
		{tag.Polarity, string(ib.Polarity), vr.CS},
		{tag.MagnificationType, string(ib.MagnificationType), vr.CS},
		{tag.SmoothingType, ib.SmoothingType, vr.CS},
		{tag.RequestedDecimateCropBehavior, ib.RequestedDecimateCropBehavior, vr.CS},
	}
	for _, value := range values {
		if err := addString(ds, value.tag, value.value, value.vr); err != nil {
			return nil, err
		}
	}
	if err := addElement(ds, element.NewUnsignedShort(tag.ImageBoxPosition, []uint16{ib.ImageBoxPosition})); err != nil {
		return nil, err
	}
	if ib.RequestedImageSize != "" {
		if err := addElement(ds, element.NewDecimalString(tag.RequestedImageSize, []string{ib.RequestedImageSize})); err != nil {
			return nil, err
		}
	}
	if ib.MaxDensity != nil {
		if err := addElement(ds, element.NewUnsignedShort(tag.MaxDensity, []uint16{*ib.MaxDensity})); err != nil {
			return nil, err
		}
	}
	if ib.MinDensity != nil {
		if err := addElement(ds, element.NewUnsignedShort(tag.MinDensity, []uint16{*ib.MinDensity})); err != nil {
			return nil, err
		}
	}
	if ib.ConfigurationInformation != nil {
		if err := addElement(ds, element.NewString(tag.ConfigurationInformation, vr.ST, []string{*ib.ConfigurationInformation})); err != nil {
			return nil, err
		}
	}
	imageSequence, err := ib.ImageSequence()
	if err != nil {
		return nil, err
	}
	if imageSequence != nil {
		sequenceTag := tag.BasicGrayscaleImageSequence
		if ib.IsColor {
			sequenceTag = tag.BasicColorImageSequence
		}
		if err := addElement(ds, dataset.NewSequenceWithItems(sequenceTag, []*dataset.Dataset{imageSequence})); err != nil {
			return nil, err
		}
	}
	return ds, nil
}

// NewImageBoxFromDataset constructs an Image Box from Dataset values.
func NewImageBoxFromDataset(sopInstanceUID string, ds *dataset.Dataset, isColor bool) (*ImageBox, error) {
	classUID, err := requiredString(ds, tag.SOPClassUID, "SOP Class UID")
	if err != nil {
		return nil, err
	}
	instanceUID, err := overrideOrDatasetUID(sopInstanceUID, ds, tag.SOPInstanceUID, "SOP Instance UID")
	if err != nil {
		return nil, err
	}
	isColor = isColor || classUID == SOPClassColorImageBox
	ib := NewImageBox(instanceUID, isColor)
	ib.SOPClassUID = classUID
	if value, err := ds.GetUInt16(tag.ImageBoxPosition, 0); err == nil {
		ib.ImageBoxPosition = value
	}
	if value, ok := ds.GetString(tag.Polarity); ok {
		ib.Polarity = Polarity(value)
	}
	if value, ok := ds.GetString(tag.MagnificationType); ok {
		ib.MagnificationType = MagnificationType(value)
	}
	if value, ok := ds.GetString(tag.SmoothingType); ok {
		ib.SmoothingType = value
	}
	if value, ok, err := optionalDecimalString(ds, tag.RequestedImageSize); err != nil {
		return nil, err
	} else if ok {
		ib.RequestedImageSize = value
	}
	if value, ok, err := optionalUInt16(ds, tag.MaxDensity); err != nil {
		return nil, err
	} else if ok {
		ib.MaxDensity = &value
	}
	if value, ok, err := optionalUInt16(ds, tag.MinDensity); err != nil {
		return nil, err
	} else if ok {
		ib.MinDensity = &value
	}
	if configElement, ok := ds.Get(tag.ConfigurationInformation); ok {
		if configElement.ValueRepresentation() != vr.ST {
			return nil, fmt.Errorf("printing: %s has VR %s, want ST", tag.ConfigurationInformation, configElement.ValueRepresentation())
		}
		if configElement.Count() != 1 {
			return nil, fmt.Errorf("printing: %s has VM %d, want 1", tag.ConfigurationInformation, configElement.Count())
		}
		value, valid := ds.GetString(tag.ConfigurationInformation)
		if !valid {
			return nil, fmt.Errorf("printing: invalid %s: expected a string value", tag.ConfigurationInformation)
		}
		ib.ConfigurationInformation = &value
	}
	if value, ok := ds.GetString(tag.RequestedDecimateCropBehavior); ok {
		ib.RequestedDecimateCropBehavior = value
	}
	sequenceTag := tag.BasicGrayscaleImageSequence
	if isColor {
		sequenceTag = tag.BasicColorImageSequence
	}
	sequence, sequenceErr := ds.GetSequence(sequenceTag)
	if sequenceErr != nil {
		if _, exists := ds.Get(sequenceTag); exists {
			return nil, fmt.Errorf("printing: read %s: %w", sequenceTag, sequenceErr)
		}
	} else {
		if sequence.Count() != 1 || sequence.GetItem(0) == nil {
			return nil, fmt.Errorf("printing: %s must contain exactly one item", sequenceTag)
		}
		if err := ib.SetImageSequence(sequence.GetItem(0)); err != nil {
			return nil, err
		}
	}
	return ib, nil
}

// ToDataset returns an independent Dataset containing the Presentation LUT attributes.
func (p *PresentationLUT) ToDataset() (*dataset.Dataset, error) {
	if p == nil {
		return nil, fmt.Errorf("printing: nil PresentationLUT")
	}
	if p.SOPInstanceUID == "" {
		return nil, fmt.Errorf("printing: PresentationLUT requires an SOP Instance UID")
	}
	ds := dataset.New()
	if err := addString(ds, tag.SOPClassUID, presentationLUTSOPClassUID, vr.UI); err != nil {
		return nil, err
	}
	if err := addString(ds, tag.SOPInstanceUID, p.SOPInstanceUID, vr.UI); err != nil {
		return nil, err
	}
	lutItem := dataset.New()
	if len(p.LUTDescriptor) > 0 {
		if err := addElement(lutItem, element.NewUnsignedShort(tag.LUTDescriptor, append([]uint16(nil), p.LUTDescriptor...))); err != nil {
			return nil, err
		}
	}
	if err := addString(lutItem, tag.LUTExplanation, p.LUTExplanation, vr.LO); err != nil {
		return nil, err
	}
	if len(p.LUTData) > 0 {
		if err := addElement(lutItem, element.NewUnsignedShort(tag.LUTData, append([]uint16(nil), p.LUTData...))); err != nil {
			return nil, err
		}
	}
	if err := addElement(ds, dataset.NewSequenceWithItems(tag.PresentationLUTSequence, []*dataset.Dataset{lutItem})); err != nil {
		return nil, err
	}
	if err := addString(ds, tag.PresentationLUTShape, string(p.PresentationLUTShape), vr.CS); err != nil {
		return nil, err
	}
	return ds, nil
}

// NewPresentationLUTFromDataset constructs a Presentation LUT from Dataset values.
func NewPresentationLUTFromDataset(sopInstanceUID string, ds *dataset.Dataset) (*PresentationLUT, error) {
	if _, err := requiredString(ds, tag.SOPClassUID, "SOP Class UID"); err != nil {
		return nil, err
	}
	instanceUID, err := overrideOrDatasetUID(sopInstanceUID, ds, tag.SOPInstanceUID, "SOP Instance UID")
	if err != nil {
		return nil, err
	}
	p := NewPresentationLUT(instanceUID)
	lutValues := ds
	sequence, sequenceErr := ds.GetSequence(tag.PresentationLUTSequence)
	if sequenceErr != nil {
		if _, exists := ds.Get(tag.PresentationLUTSequence); exists {
			return nil, fmt.Errorf("printing: read Presentation LUT Sequence: %w", sequenceErr)
		}
	} else {
		if sequence.Count() != 1 || sequence.GetItem(0) == nil {
			return nil, fmt.Errorf("printing: Presentation LUT Sequence must contain exactly one item")
		}
		lutValues = sequence.GetItem(0)
	}
	if values, err := lutValues.GetUInt16s(tag.LUTDescriptor); err == nil {
		p.LUTDescriptor = append([]uint16(nil), values...)
	}
	if value, ok := lutValues.GetString(tag.LUTExplanation); ok {
		p.LUTExplanation = value
	}
	if values, err := lutValues.GetUInt16s(tag.LUTData); err == nil {
		p.LUTData = append([]uint16(nil), values...)
	}
	if value, ok := ds.GetString(tag.PresentationLUTShape); ok {
		p.PresentationLUTShape = PresentationLUTShape(value)
	}
	return p, nil
}
