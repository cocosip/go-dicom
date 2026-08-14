// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dict_test

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func ExampleNewFromXML() {
	const source = `<dictionary creator="VENDOR">
  <tag group="0011" element="1010" keyword="VendorValue" vr="LO" vm="1">Vendor Value</tag>
</dictionary>`

	dictionary, err := dict.NewFromXML(strings.NewReader(source))
	if err != nil {
		panic(err)
	}

	creator := dictionary.GetPrivateCreator("VENDOR")
	entry := dictionary.Lookup(tag.NewWithPrivateCreator(0x0011, 0x1210, creator))
	fmt.Println(entry.Name(), entry.VRs(), entry.VM())
	// Output: Vendor Value [LO] 1
}
