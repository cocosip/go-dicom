// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package pixeldata owns DICOM pixel metadata, frames, and Dataset conversion.
//
// New Pixel Data can derive its OB, OW, or encapsulated representation from a
// Dataset's image metadata and transfer syntax:
//
//	pixels, err := pixeldata.NewForDataset(ds)
//	if err != nil {
//		return err
//	}
//	for _, frame := range frames {
//		if err := pixels.AddFrame(ctx, frame); err != nil {
//			return err
//		}
//	}
//	if err := pixels.WriteToDataset(ds); err != nil {
//		return err
//	}
//
// FromDataset opens existing native or encapsulated Pixel Data. The lower-level
// dicom/encapsulated package remains responsible for fragment and offset-table
// frame assembly.
package pixeldata
