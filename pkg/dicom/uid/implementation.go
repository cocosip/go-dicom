// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package uid

// Implementation identification shared by the file writer, the dataset defaults
// and the network layer.
//
// Both values are placeholders pending ISO/UID registration: applications
// embedding go-dicom should set their own registered Implementation Class UID
// and Implementation Version Name instead of relying on these defaults.
const (
	// GoDicomImplementationClassUID is the Implementation Class UID (0002,0012)
	// used for files and for A-ASSOCIATE user information.
	GoDicomImplementationClassUID = "1.2.826.0.1.3680043.10.854"
	// GoDicomImplementationVersionName is the Implementation Version Name
	// (0002,0013). The SH VR limits it to 16 characters.
	GoDicomImplementationVersionName = "GO-DICOM-0.1"
)
