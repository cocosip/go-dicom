// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package uid

import (
	"runtime/debug"
	"strings"
)

// Implementation identification shared by the file writer, the dataset defaults
// and the network layer.
//
// The Implementation Class UID is a placeholder pending ISO/UID registration:
// applications embedding go-dicom should set their own registered
// Implementation Class UID instead of relying on this default.
const (
	// GoDicomImplementationClassUID is the Implementation Class UID (0002,0012)
	// used for files and for A-ASSOCIATE user information.
	GoDicomImplementationClassUID = "1.2.826.0.1.3680043.10.854"
	// GoDicomModulePath is the module whose build version is reported in
	// ImplementationVersionName.
	GoDicomModulePath = "github.com/cocosip/go-dicom"
)

const (
	// implementationVersionNamePrefix precedes the module version.
	implementationVersionNamePrefix = "GO_DICOM_"
	// implementationVersionNameMaxLength is the SH VR limit (PS3.5 6.2).
	implementationVersionNameMaxLength = 16
	// developmentVersionName is reported when the binary records no version:
	// `go run`, a source archive, or a module built without build information.
	developmentVersionName = "GO_DICOM_DEV"
)

// GoDicomImplementationVersionName is the Implementation Version Name
// (0002,0013). It is derived from the go-dicom module version recorded in the
// binary's build information, so a program that depends on go-dicom v0.9.0
// reports "GO_DICOM_0_9_0". PS3.5 A.1 limits the value to 16 characters of
// uppercase letters, digits and underscores.
var GoDicomImplementationVersionName = implementationVersionName(developmentVersionName)

func implementationVersionName(fallback string) string {
	version := buildModuleVersion()
	if version == "" {
		return fallback
	}
	return formatImplementationVersionName(version)
}

// buildModuleVersion returns the go-dicom release version recorded in the
// binary's build information: the main module version when go-dicom is the
// program, otherwise the version of go-dicom as a dependency. It returns ""
// when no release version is recorded.
func buildModuleVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	if info.Main.Path == GoDicomModulePath && isReleaseVersion(info.Main.Version) {
		return info.Main.Version
	}
	for _, dep := range info.Deps {
		if dep == nil || dep.Path != GoDicomModulePath {
			continue
		}
		if isReleaseVersion(dep.Version) {
			return dep.Version
		}
	}
	return ""
}

// isReleaseVersion reports whether a module version is a tag or pseudo-version
// rather than the "(devel)" placeholder of local builds.
func isReleaseVersion(version string) bool {
	return len(version) > 1 && version[0] == 'v'
}

// formatImplementationVersionName converts a module version (v0.9.0) into an
// Implementation Version Name (GO_DICOM_0_9_0), truncated to the SH limit.
func formatImplementationVersionName(version string) string {
	version = strings.TrimPrefix(version, "v")
	var b strings.Builder
	b.Grow(len(implementationVersionNamePrefix) + len(version))
	b.WriteString(implementationVersionNamePrefix)
	for i := 0; i < len(version); i++ {
		c := version[i]
		switch {
		case c >= 'a' && c <= 'z':
			b.WriteByte(c - 'a' + 'A')
		case c >= 'A' && c <= 'Z', c >= '0' && c <= '9':
			b.WriteByte(c)
		case c == '.' || c == '-' || c == '+' || c == '_':
			b.WriteByte('_')
		}
	}
	name := b.String()
	if len(name) > implementationVersionNameMaxLength {
		name = name[:implementationVersionNameMaxLength]
		// Cut at a version-separator boundary so the name never ends mid-token.
		if i := strings.LastIndexByte(name, '_'); i > 0 {
			name = name[:i]
		}
	}
	return name
}
