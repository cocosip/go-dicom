// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

// StripTrailingPadding removes a single trailing 0x00 that follows the JPEG/JPEG-LS EOI
// marker (0xFF 0xD9). Some encoders pad encapsulated fragments to even length; the padding
// byte must not be passed to a codec decoder.
func StripTrailingPadding(data []byte) []byte {
	if len(data) >= 3 && data[len(data)-1] == 0x00 &&
		data[len(data)-2] == 0xD9 && data[len(data)-3] == 0xFF {
		return data[:len(data)-1]
	}
	return data
}
