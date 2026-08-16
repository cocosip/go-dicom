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
	if imageData := ib.GetImageData(); len(imageData) > 0 {
		sequenceTag := tag.BasicGrayscaleImageSequence
		if ib.IsColor {
			sequenceTag = tag.BasicColorImageSequence
		}
		item := dataset.New()
		if err := addElement(item, element.NewOtherByte(tag.PixelData, append([]byte(nil), imageData...))); err != nil {
			return nil, err
		}
		if err := addElement(ds, dataset.NewSequenceWithItems(sequenceTag, []*dataset.Dataset{item})); err != nil {
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
	sequenceTag := tag.BasicGrayscaleImageSequence
	if isColor {
		sequenceTag = tag.BasicColorImageSequence
	}
	if sequence, err := ds.GetSequence(sequenceTag); err == nil && sequence.Count() > 0 {
		if item := sequence.GetItem(0); item != nil {
			if pixelData := item.GetOrNil(tag.PixelData); pixelData != nil && pixelData.Buffer() != nil {
				ib.SetImageData(append([]byte(nil), pixelData.Buffer().Data()...))
			}
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
	if len(p.LUTDescriptor) > 0 {
		if err := addElement(ds, element.NewUnsignedShort(tag.LUTDescriptor, append([]uint16(nil), p.LUTDescriptor...))); err != nil {
			return nil, err
		}
	}
	if err := addString(ds, tag.LUTExplanation, p.LUTExplanation, vr.LO); err != nil {
		return nil, err
	}
	if len(p.LUTData) > 0 {
		if err := addElement(ds, element.NewUnsignedShort(tag.LUTData, append([]uint16(nil), p.LUTData...))); err != nil {
			return nil, err
		}
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
	if values, err := ds.GetUInt16s(tag.LUTDescriptor); err == nil {
		p.LUTDescriptor = append([]uint16(nil), values...)
	}
	if value, ok := ds.GetString(tag.LUTExplanation); ok {
		p.LUTExplanation = value
	}
	if values, err := ds.GetUInt16s(tag.LUTData); err == nil {
		p.LUTData = append([]uint16(nil), values...)
	}
	if value, ok := ds.GetString(tag.PresentationLUTShape); ok {
		p.PresentationLUTShape = PresentationLUTShape(value)
	}
	return p, nil
}
