// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import "testing"

func TestGlobalRegistryDoesNotRegisterRLE(t *testing.T) {
	const rleLosslessUID = "1.2.840.10008.1.2.5"

	for _, uid := range GetGlobalRegistry().ListCodecs() {
		if uid == rleLosslessUID {
			t.Fatal("global registry must not register an RLE codec")
		}
	}
}
