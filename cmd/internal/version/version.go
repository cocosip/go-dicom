// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

// Package version reports the go-dicom build version for the command-line
// tools, so -version never drifts from the released module version.
package version

import "runtime/debug"

// String returns the version recorded in the binary's build information: the
// module version for `go install ...@vX.Y.Z` or a release build, and "devel"
// for a working-tree build that records no release version.
func String() string {
	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" || info.Main.Version == "(devel)" {
		return "devel"
	}
	return info.Main.Version
}
