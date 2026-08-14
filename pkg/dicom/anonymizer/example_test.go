// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package anonymizer_test

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/anonymizer"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func ExampleNewProfileFromReader() {
	const profileData = "0010,0010;Z\n0008,0018;U;;K;;;;;;;;\n"
	profile, err := anonymizer.NewProfileFromReader(
		strings.NewReader(profileData),
		anonymizer.BasicProfile|anonymizer.RetainUIDs,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	patientNameAction, patientNameFound := profile.FindAction(tag.PatientName)
	sopInstanceUIDAction, sopInstanceUIDFound := profile.FindAction(tag.SOPInstanceUID)
	fmt.Println(patientNameAction, patientNameFound)
	fmt.Println(sopInstanceUIDAction, sopInstanceUIDFound)

	// Output:
	// Z true
	// K true
}
