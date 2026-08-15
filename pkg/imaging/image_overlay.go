// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"encoding/binary"
	"image"
	"image/color"
	"image/draw"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
)

func imageOverlays(ds *dataset.Dataset, pixelData *DicomPixelData) []*DicomOverlayData {
	var overlays []*DicomOverlayData
	for _, group := range OverlayGroupNumbers() {
		rows, err := ds.GetUInt16(tag.New(group, 0x0010), 0)
		if err != nil {
			continue
		}
		columns, err := ds.GetUInt16(tag.New(group, 0x0011), 0)
		if err != nil {
			continue
		}
		overlay := NewDicomOverlayData(group)
		overlay.Rows = int(rows)
		overlay.Columns = int(columns)
		overlay.OriginRow = 1
		overlay.OriginColumn = 1
		overlay.ImageFrameOrigin = 1
		if origin, ok := ds.Get(tag.New(group, 0x0050)); ok {
			if values, ok := origin.(*element.SignedShort); ok {
				if row, err := values.GetValue(0); err == nil {
					overlay.OriginRow = int(row)
				}
				if column, err := values.GetValue(1); err == nil {
					overlay.OriginColumn = int(column)
				}
			}
		}
		if frames, ok := ds.Get(tag.New(group, 0x0015)); ok {
			if values, ok := frames.(*element.IntegerString); ok {
				if count, err := values.GetInt(0); err == nil && count > 0 {
					overlay.NumberOfFrames = count
				}
			}
		}
		if origin, err := ds.GetUInt16(tag.New(group, 0x0051), 0); err == nil && origin > 0 {
			overlay.ImageFrameOrigin = int(origin)
		}

		if dataElement, ok := ds.Get(tag.New(group, 0x3000)); ok {
			switch value := dataElement.(type) {
			case *element.OtherByte:
				overlay.Data = append([]byte(nil), value.GetData()...)
			case *element.OtherWord:
				overlay.Data = append([]byte(nil), value.GetData()...)
			}
		} else if bitsAllocated, err := ds.GetUInt16(tag.New(group, 0x0100), 0); err == nil && pixelData != nil && bitsAllocated == pixelData.Info.BitsAllocated {
			if bitPosition, err := ds.GetUInt16(tag.New(group, 0x0102), 0); err == nil && bitPosition < bitsAllocated {
				overlay.BitPosition = int(bitPosition)
				overlay.Data = extractEmbeddedOverlay(pixelData, overlay)
			}
		}
		if len(overlay.Data)*8 >= overlay.Rows*overlay.Columns*overlay.NumberOfFrames {
			overlays = append(overlays, overlay)
		}
	}
	return overlays
}

func extractEmbeddedOverlay(pixelData *DicomPixelData, overlay *DicomOverlayData) []byte {
	bytesPerSample := pixelData.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 && bytesPerSample != 4 {
		return nil
	}
	unpacked := make([]byte, overlay.UnpackedDataSize())
	framePixels := overlay.Rows * overlay.Columns
	for overlayFrame := 0; overlayFrame < overlay.NumberOfFrames; overlayFrame++ {
		imageFrame := overlay.ImageFrameOrigin - 1 + overlayFrame
		frame, err := pixelData.GetFrame(imageFrame)
		if err != nil {
			return nil
		}
		for pixel := 0; pixel < framePixels; pixel++ {
			offset := pixel * bytesPerSample
			if offset+bytesPerSample > len(frame) {
				break
			}
			var raw uint32
			switch bytesPerSample {
			case 1:
				raw = uint32(frame[offset])
			case 2:
				raw = uint32(binary.LittleEndian.Uint16(frame[offset:]))
			case 4:
				raw = binary.LittleEndian.Uint32(frame[offset:])
			}
			if raw&(uint32(1)<<uint(overlay.BitPosition)) != 0 {
				unpacked[overlayFrame*framePixels+pixel] = 255
			}
		}
	}
	packed := NewDicomOverlayData(overlay.Group)
	packed.Rows = overlay.Rows
	packed.Columns = overlay.Columns
	packed.NumberOfFrames = overlay.NumberOfFrames
	if err := packed.PackData(unpacked); err != nil {
		return nil
	}
	return packed.Data
}

func (img *DicomImage) applyOverlays(source image.Image, frame int) image.Image {
	img.mu.RLock()
	show := img.showOverlays
	overlays := append([]*DicomOverlayData(nil), img.overlays...)
	overlayColor := img.overlayColor
	img.mu.RUnlock()
	if !show || len(overlays) == 0 {
		return source
	}

	destination := image.NewRGBA(source.Bounds())
	draw.Draw(destination, destination.Bounds(), source, source.Bounds().Min, draw.Src)
	for _, overlay := range overlays {
		overlayFrame := frame + 1 - overlay.ImageFrameOrigin
		if overlayFrame < 0 || overlayFrame >= overlay.NumberOfFrames {
			continue
		}
		data, err := overlay.GetFrame(overlayFrame)
		if err != nil {
			continue
		}
		for row := 0; row < overlay.Rows; row++ {
			for column := 0; column < overlay.Columns; column++ {
				if data[row*overlay.Columns+column] == 0 {
					continue
				}
				x := overlay.OriginColumn - 1 + column
				y := overlay.OriginRow - 1 + row
				if image.Pt(x, y).In(destination.Bounds()) {
					destination.Set(x, y, color.NRGBA{R: overlayColor.R, G: overlayColor.G, B: overlayColor.B, A: overlayColor.A})
				}
			}
		}
	}
	return destination
}

// ShowOverlays reports whether Dataset overlays are rendered.
func (img *DicomImage) ShowOverlays() bool {
	img.mu.RLock()
	defer img.mu.RUnlock()
	return img.showOverlays
}

// SetShowOverlays enables or disables Dataset overlay rendering.
func (img *DicomImage) SetShowOverlays(show bool) {
	img.mu.Lock()
	defer img.mu.Unlock()
	img.showOverlays = show
}

// OverlayColor returns the color used to render Dataset overlays.
func (img *DicomImage) OverlayColor() imagetypes.Color32 {
	img.mu.RLock()
	defer img.mu.RUnlock()
	return img.overlayColor
}

// SetOverlayColor changes the color used to render Dataset overlays.
func (img *DicomImage) SetOverlayColor(value imagetypes.Color32) {
	img.mu.Lock()
	defer img.mu.Unlock()
	img.overlayColor = value
}
