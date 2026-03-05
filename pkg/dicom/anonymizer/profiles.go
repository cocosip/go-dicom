// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package anonymizer

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

// AdditionalSecurityProfiles provides additional pre-defined security profiles
// beyond the default BasicProfile.
type AdditionalSecurityProfiles struct{}

// NewRetainUIDsProfile creates a profile that retains all UIDs.
// This is useful for research where UID consistency is required.
func NewRetainUIDsProfile() *SecurityProfile {
	profile := NewSecurityProfile(RetainUIDs)

	// Override UID-related tags to keep them
	profile.rules = append(profile.rules, []profileRule{
		{mustCompile("(?i)" + `0008,0018`), ActionK}, // SOP Instance UID
		{mustCompile("(?i)" + `0008,0020`), ActionK}, // Study Instance UID
		{mustCompile("(?i)" + `0008,0021`), ActionK}, // Series Instance UID
		{mustCompile("(?i)" + `0008,0030`), ActionK}, // Study Time (keep for research)
		{mustCompile("(?i)" + `0008,0050`), ActionK}, // Accession Number
		{mustCompile("(?i)" + `0010,0020`), ActionK}, // Patient ID
		{mustCompile("(?i)" + `0020,000D`), ActionK}, // Study Instance UID
		{mustCompile("(?i)" + `0020,000E`), ActionK}, // Series Instance UID
		{mustCompile("(?i)" + `0020,0010`), ActionK}, // Study ID
		{mustCompile("(?i)" + `0020,0052`), ActionK}, // Frame of Reference UID
		{mustCompile("(?i)" + `0020,0200`), ActionK}, // Synchronization Frame of Reference UID
		{mustCompile("(?i)" + `0020,3401`), ActionK}, // RT Frame of Reference UID
	}...)

	return profile
}

// NewRetainInstitutionProfile creates a profile that retains institution identifiers.
// This is useful for multi-center studies where institution tracking is needed.
func NewRetainInstitutionProfile() *SecurityProfile {
	profile := NewSecurityProfile(RetainInstitutionIdent)

	// Override institution-related tags to keep them
	profile.rules = append(profile.rules, []profileRule{
		{mustCompile("(?i)" + `0008,0080`), ActionK}, // Institution Name
		{mustCompile("(?i)" + `0008,0081`), ActionK}, // Institution Address
		{mustCompile("(?i)" + `0008,0082`), ActionK}, // Institution Code Sequence
		{mustCompile("(?i)" + `0008,0090`), ActionK}, // Referring Physician's Name
		{mustCompile("(?i)" + `0008,0092`), ActionK}, // Referring Physician's Address
		{mustCompile("(?i)" + `0008,0094`), ActionK}, // Referring Physician's Telephone Numbers
		{mustCompile("(?i)" + `0008,0096`), ActionK}, // Referring Physician Identification Sequence
		{mustCompile("(?i)" + `0008,1040`), ActionK}, // Institutional Department Name
		{mustCompile("(?i)" + `0008,1048`), ActionK}, // Physician(s) of Record
		{mustCompile("(?i)" + `0008,1050`), ActionK}, // Performing Physician's Name
		{mustCompile("(?i)" + `0008,1060`), ActionK}, // Name of Physician(s) Reading Study
		{mustCompile("(?i)" + `0008,1070`), ActionK}, // Operators' Name
	}...)

	return profile
}

// NewRetainDatesProfile creates a profile that retains all dates.
// This is useful for longitudinal studies where date information is critical.
// Note: This may compromise privacy in some jurisdictions.
func NewRetainDatesProfile() *SecurityProfile {
	profile := NewSecurityProfile(RetainLongFullDates | RetainLongModifDates)

	// Override date-related tags to keep them
	profile.rules = append(profile.rules, []profileRule{
		{mustCompile("(?i)" + `0008,0012`), ActionK}, // Instance Creation Date
		{mustCompile("(?i)" + `0008,0013`), ActionK}, // Instance Creation Time
		{mustCompile("(?i)" + `0008,0020`), ActionK}, // Study Date
		{mustCompile("(?i)" + `0008,0021`), ActionK}, // Series Date
		{mustCompile("(?i)" + `0008,0022`), ActionK}, // Acquisition Date
		{mustCompile("(?i)" + `0008,0023`), ActionK}, // Content Date
		{mustCompile("(?i)" + `0008,0030`), ActionK}, // Study Time
		{mustCompile("(?i)" + `0008,0031`), ActionK}, // Series Time
		{mustCompile("(?i)" + `0008,0032`), ActionK}, // Acquisition Time
		{mustCompile("(?i)" + `0008,0033`), ActionK}, // Content Time
		{mustCompile("(?i)" + `0008,0050`), ActionK}, // Accession Number (often contains date info)
		{mustCompile("(?i)" + `0008,0060`), ActionK}, // Modality
		{mustCompile("(?i)" + `0010,0030`), ActionK}, // Patient's Birth Date
		{mustCompile("(?i)" + `0010,0032`), ActionK}, // Patient's Birth Time
		{mustCompile("(?i)" + `0010,0033`), ActionK}, // Patient's Birth Date in Alternative Calendar
		{mustCompile("(?i)" + `0010,0034`), ActionK}, // Patient's Alternative Calendar
		{mustCompile("(?i)" + `0012,0050`), ActionK}, // Clinical Trial Time Point ID
		{mustCompile("(?i)" + `0012,0051`), ActionK}, // Clinical Trial Time Point Description
		{mustCompile("(?i)" + `0012,0071`), ActionK}, // Clinical Trial Coordinating Center Name
		{mustCompile("(?i)" + `0012,0072`), ActionK}, // Clinical Trial Protocol ID
		{mustCompile("(?i)" + `0012,0073`), ActionK}, // Clinical Trial Protocol Name
		{mustCompile("(?i)" + `0012,0081`), ActionK}, // Clinical Trial Site ID
		{mustCompile("(?i)" + `0012,0082`), ActionK}, // Clinical Trial Site Name
		{mustCompile("(?i)" + `0040,0002`), ActionK}, // Study Entry Date
		{mustCompile("(?i)" + `0040,0003`), ActionK}, // Study Entry Time
		{mustCompile("(?i)" + `0040,0004`), ActionK}, // Study End Date
		{mustCompile("(?i)" + `0040,0005`), ActionK}, // Study End Time
		{mustCompile("(?i)" + `0040,0244`), ActionK}, // Scheduled Procedure Step Start Date
		{mustCompile("(?i)" + `0040,0245`), ActionK}, // Scheduled Procedure Step Start Time
		{mustCompile("(?i)" + `0040,0250`), ActionK}, // Scheduled Procedure Step End Date
		{mustCompile("(?i)" + `0040,0251`), ActionK}, // Scheduled Procedure Step End Time
	}...)

	return profile
}

// NewMinimalProfile creates a minimal anonymization profile that only removes
// direct patient identifiers. This is the least aggressive profile.
// WARNING: This may not be sufficient for HIPAA or GDPR compliance.
func NewMinimalProfile() *SecurityProfile {
	profile := &SecurityProfile{
		rules: make([]profileRule, 0),
	}

	// Only remove the most critical patient identifiers
	minimalRules := []string{
		// Patient identifiers - remove
		`0010,0010`, // Patient's Name - Z
		`0010,0020`, // Patient ID - Z
		`0010,0030`, // Patient's Birth Date - Z
		`0010,0032`, // Patient's Birth Time - Z
		`0010,0033`, // Patient's Birth Date in Alternative Calendar - Z
		`0010,0040`, // Patient's Sex - Z
		`0010,0050`, // Insurance Plan Identification - X
		`0010,0101`, // Patient's Primary Language - X
		`0010,1000`, // Other Patient IDs - X
		`0010,1001`, // Other Patient Names - X
		`0010,1002`, // Other Patient IDs Sequence - X
		`0010,1005`, // Patient's Birth Name - X
		`0010,1010`, // Patient's Age - Z
		`0010,1020`, // Patient's Size - X
		`0010,1030`, // Patient's Weight - X
		`0010,1040`, // Patient's Address - X
		`0010,1050`, // Insurance Plan Identification Sequence - X
		`0010,1060`, // Patient's Mother's Birth Name - X
		`0010,1080`, // Military Rank - X
		`0010,1081`, // Branch of Service - X
		`0010,1090`, // Medical Record Locator - X
		`0010,2000`, // Medical Alerts - X
		`0010,2110`, // Allergies - X
		`0010,2150`, // Country of Residence - X
		`0010,2152`, // Region of Residence - X
		`0010,2154`, // Home Telephone Numbers - X
		`0010,2160`, // Ethnic Group - X
		`0010,2180`, // Occupation - X
		`0010,21A0`, // Smoking Status - X
		`0010,21B0`, // Additional Patient History - X
		`0010,21C0`, // Pregnancy Status - X
		`0010,21D0`, // Last Menstrual Date - X
		`0010,21F0`, // Patient's Religious Preferences - X
		`0010,2203`, // Last Menstrual Date Approximate - X
		`0010,4000`, // Patient Comments - X
		`0038,0010`, // Admission ID - X
		`0038,0011`, // Issuer of Admission ID - X
		`0038,001E`, // Scheduled Patient Admission Date - X
		`0038,0020`, // Admitting Date - Z
		`0038,0021`, // Discharge Date - Z
		`0038,0040`, // Discharge Time - Z
		`0038,0050`, // Hospital Admitting Diagnoses - X
		`0038,0060`, // Payer Name - X
		`0038,0300`, // Current Patient Location - X
		`0038,0400`, // Patient's Institution Residence - X
		`0038,0500`, // Special Needs - X
		`0040,0280`, // Comments on the Scheduled Procedure Step - X
		`0040,1001`, // Requested Procedure ID - X
		`0040,1004`, // Patient Transport Arrangements - X
		`0040,1005`, // Requested Procedure Location - X
	}

	for _, tagPattern := range minimalRules {
		// Determine action (Z for most, X for comments/history)
		action := ActionZ
		if strings.Contains(tagPattern, "0010,4000") || // Patient Comments
			strings.Contains(tagPattern, "0038,0010") || // Admission ID
			strings.Contains(tagPattern, "0038,0050") || // Hospital Admitting Diagnoses
			strings.Contains(tagPattern, "0038,0060") || // Payer Name
			strings.Contains(tagPattern, "0038,0300") || // Current Patient Location
			strings.Contains(tagPattern, "0038,0400") || // Patient's Institution Residence
			strings.Contains(tagPattern, "0038,0500") { // Special Needs
			action = ActionX
		}

		_ = profile.AddRule(tagPattern, action)
	}

	return profile
}

// NewResearchProfile creates a profile optimized for research use.
// It retains more information while still protecting patient identity.
func NewResearchProfile() *SecurityProfile {
	profile := NewSecurityProfile(RetainUIDs | RetainLongFullDates)

	// Research-friendly rules
	researchRules := []struct {
		pattern string
		action  SecurityProfileAction
	}{
		// Keep UIDs for study consistency
		{`0008,0018`, ActionK}, // SOP Instance UID
		{`0008,0020`, ActionK}, // Study Date (keep for temporal analysis)
		{`0008,0021`, ActionK}, // Series Date
		{`0008,0030`, ActionK}, // Study Time
		{`0008,0031`, ActionK}, // Series Time
		{`0008,0050`, ActionZ}, // Accession Number (zero out)
		{`0008,0080`, ActionC}, // Institution Name (clean)
		{`0008,0090`, ActionZ}, // Referring Physician's Name (zero out)
		{`0008,1030`, ActionC}, // Requested Procedure Description (clean)
		{`0008,103E`, ActionC}, // Series Description (clean)

		// Patient demographics - keep some for research
		{`0010,0010`, ActionZ}, // Patient's Name (zero out)
		{`0010,0020`, ActionZ}, // Patient ID (zero out)
		{`0010,0030`, ActionZ}, // Patient's Birth Date (zero out)
		{`0010,0040`, ActionC}, // Patient's Sex (clean - keep value)
		{`0010,0050`, ActionX}, // Insurance Plan Identification (remove)
		{`0010,1010`, ActionK}, // Patient's Age (keep for analysis)
		{`0010,1020`, ActionK}, // Patient's Size (keep for analysis)
		{`0010,1030`, ActionK}, // Patient's Weight (keep for analysis)
		{`0010,1040`, ActionX}, // Patient's Address (remove)
		{`0010,1060`, ActionX}, // Patient's Mother's Birth Name (remove)
		{`0010,1080`, ActionX}, // Military Rank (remove)
		{`0010,2000`, ActionX}, // Medical Alerts (remove)
		{`0010,2110`, ActionX}, // Allergies (remove for privacy)
		{`0010,2150`, ActionX}, // Country of Residence (remove)
		{`0010,2160`, ActionK}, // Ethnic Group (keep for research)
		{`0010,2180`, ActionX}, // Occupation (remove)
		{`0010,21A0`, ActionK}, // Smoking Status (keep for research)
		{`0010,21B0`, ActionX}, // Additional Patient History (remove)
		{`0010,21F0`, ActionX}, // Patient's Religious Preferences (remove)
		{`0010,4000`, ActionX}, // Patient Comments (remove)

		// Study/Series info - keep for research
		{`0008,0060`, ActionK}, // Modality
		{`0008,0061`, ActionK}, // Modalities in Study
		{`0008,0070`, ActionC}, // Manufacturer (clean)
		{`0008,0080`, ActionC}, // Institution Name (clean)
		{`0008,1010`, ActionC}, // Station Name (clean)
		{`0008,1030`, ActionC}, // Requested Procedure Description (clean)
		{`0008,103E`, ActionC}, // Series Description (clean)
		{`0008,1040`, ActionK}, // Institutional Department Name
		{`0008,1050`, ActionZ}, // Performing Physician's Name (zero out)
		{`0008,1060`, ActionZ}, // Name of Physician(s) Reading Study (zero out)
		{`0008,1070`, ActionZ}, // Operators' Name (zero out)
		{`0008,1080`, ActionC}, // Admitting Diagnoses Description (clean)
		{`0008,1090`, ActionC}, // Manufacturer's Model Name (clean)

		// Device info - keep for technical analysis
		{`0018,0010`, ActionC}, // Contrast/Bolus Agent (clean)
		{`0018,0015`, ActionK}, // Chest Tube Position (keep)
		{`0018,0050`, ActionK}, // Slice Thickness (keep)
		{`0018,0060`, ActionK}, // KVP (keep)
		{`0018,0070`, ActionK}, // Counts Accumulated (keep)
		{`0018,0080`, ActionK}, // Acquisition Time (keep)
		{`0018,0081`, ActionK}, // Acquisition Duration (keep)
		{`0018,0083`, ActionK}, // Number of Slices (keep)
		{`0018,0084`, ActionK}, // Data Collection Center Name (clean)
		{`0018,0085`, ActionK}, // Data Collection Description (clean)
		{`0018,0086`, ActionK}, // Acquisition Protocol Description (clean)

		// Image properties - keep for analysis
		{`0018,1000`, ActionC}, // Device Serial Number (clean)
		{`0018,1002`, ActionC}, // Device UID (clean)
		{`0018,1007`, ActionC}, // Station Name (clean)
		{`0018,1008`, ActionC}, // Device ID (clean)
		{`0018,1009`, ActionC}, // Device Manufacturer (clean)
		{`0018,1010`, ActionC}, // Device Manufacturer's Model Name (clean)
		{`0018,1020`, ActionC}, // Software Versions (clean)
		{`0018,1030`, ActionC}, // Protocol Name (clean)
		{`0018,1100`, ActionK}, // Reconstruction Diameter (keep)
		{`0018,1110`, ActionK}, // Distance Source to Detector (keep)
		{`0018,1111`, ActionK}, // Distance Source to Patient (keep)
		{`0018,1120`, ActionK}, // Gantry/Detector Tilt (keep)
		{`0018,1130`, ActionK}, // Table Height (keep)
		{`0018,1131`, ActionK}, // Table Traverse (keep)
		{`0018,1134`, ActionK}, // Rotation Direction (keep)
		{`0018,1135`, ActionK}, // Angular Position (keep)
		{`0018,1140`, ActionK}, // Field of View Shape (keep)
		{`0018,1141`, ActionK}, // Field of View Dimensions (keep)
		{`0018,1142`, ActionK}, // Flow Identifier (keep)
		{`0018,1143`, ActionK}, // Frame Identifier (keep)
		{`0018,1144`, ActionK}, // Time Identifier (keep)
		{`0018,1145`, ActionK}, // Time (keep)
		{`0018,1150`, ActionK}, // Collimator/Filter Type (keep)
		{`0018,1151`, ActionK}, // X-Ray Tube Current (keep)
		{`0018,1152`, ActionK}, // Exposure (keep)
		{`0018,1153`, ActionK}, // Exposure in uAs (keep)
		{`0018,1154`, ActionK}, // Average Pulse Width (keep)
		{`0018,1155`, ActionK}, // Radiation Setting (keep)
		{`0018,1156`, ActionK}, // Rectification Mode (keep)
		{`0018,1157`, ActionK}, // Radiation Asymmetry Factor (keep)
		{`0018,1158`, ActionK}, // Exposure Time (keep)
		{`0018,1159`, ActionK}, // X-Ray Tube Current in uA (keep)
		{`0018,115A`, ActionK}, // X-Ray Tube Current Regulation (keep)
		{`0018,115B`, ActionK}, // Pulse Width (keep)
		{`0018,115C`, ActionK}, // Distal Focal Spot Diameter (keep)
		{`0018,115D`, ActionK}, // Proximal Focal Spot Diameter (keep)
		{`0018,115E`, ActionK}, // Distance Source to Entrance (keep)
		{`0018,1160`, ActionK}, // Filter Type (keep)
		{`0018,1161`, ActionK}, // Filter Material (keep)
		{`0018,1162`, ActionK}, // Filter Thickness Minimum (keep)
		{`0018,1163`, ActionK}, // Filter Thickness Maximum (keep)
		{`0018,1164`, ActionK}, // Total Collimation Width (keep)
		{`0018,1170`, ActionK}, // Image and Flow Acquisition Description (keep)

		// Keep study/Series UIDs
		{`0020,000D`, ActionK}, // Study Instance UID
		{`0020,000E`, ActionK}, // Series Instance UID
		{`0020,0010`, ActionZ}, // Study ID (zero out)
		{`0020,0011`, ActionK}, // Series Number (keep)
		{`0020,0012`, ActionK}, // Acquisition Number (keep)
		{`0020,0013`, ActionK}, // Instance Number (keep)
		{`0020,0020`, ActionK}, // Patient Orientation (keep)
		{`0020,0030`, ActionK}, // Image Position (Patient) (keep)
		{`0020,0032`, ActionK}, // Image Position (Patient) (keep)
		{`0020,0035`, ActionK}, // Image Orientation (Patient) (keep)
		{`0020,0052`, ActionK}, // Frame of Reference UID
		{`0020,0060`, ActionK}, // Laterality (keep)
		{`0020,0062`, ActionK}, // Image Laterality (keep)
		{`0020,0070`, ActionK}, // Geometry of the K-Space Traversal (keep)
		{`0020,0071`, ActionK}, // Segmentation Start Frame (keep)
		{`0020,0072`, ActionK}, // Segmentation End Frame (keep)
		{`0020,0100`, ActionK}, // Temporal Position Identifier (keep)
		{`0020,0105`, ActionK}, // Number of Temporal Positions (keep)
		{`0020,0110`, ActionK}, // Temporal Resolution (keep)
		{`0020,1001`, ActionK}, // Source Image IDs Sequence (keep references)
		{`0020,1002`, ActionK}, // Source Image Sequence (keep references)
		{`0020,103F`, ActionK}, // Target Adapter (keep)
		{`0020,1040`, ActionK}, // Position Reference Indicator (keep)
		{`0020,1041`, ActionK}, // Slice Location (keep)
		{`0020,1070`, ActionK}, // Other Study Numbers (keep)
		{`0020,1071`, ActionK}, // Other Series Numbers (keep)
		{`0020,1072`, ActionK}, // Other Instance Numbers (keep)
	}

	for _, rule := range researchRules {
		_ = profile.AddRule(rule.pattern, rule.action)
	}

	return profile
}

// mustCompile compiles a regex pattern or panics if invalid.
// Used for static profile rules that are known to be valid.
func mustCompile(pattern string) *regexp.Regexp {
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	return re
}

// NewProfileFromReader creates a custom security profile from a reader.
// Format: tag_pattern;action (one per line)
// Example:
//
//	0010,0010;Z
//	0010,0020;Z
//	0008,0018;K
func NewProfileFromReader(reader *strings.Reader, _ SecurityProfileOptions) (*SecurityProfile, error) {
	profile := &SecurityProfile{
		rules: make([]profileRule, 0),
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) != 2 {
			continue
		}

		pattern := strings.TrimSpace(parts[0])
		actionStr := strings.TrimSpace(parts[1])

		var action SecurityProfileAction
		switch actionStr {
		case "D":
			action = ActionD
		case "Z":
			action = ActionZ
		case "X":
			action = ActionX
		case "K":
			action = ActionK
		case "C":
			action = ActionC
		case "U":
			action = ActionU
		default:
			continue
		}

		if err := profile.AddRule(pattern, action); err != nil {
			return nil, err
		}
	}

	return profile, scanner.Err()
}

// LoadProfileFromFile loads a custom profile from a file.
// Format: tag_pattern;action (one per line)
// This function is provided for reference - actual implementation
// would need to use os.Open which requires the "os" import.
// Users should create their own file loading function as needed.
func LoadProfileFromFile(_ string, _ SecurityProfileOptions) (*SecurityProfile, error) {
	// Note: This is a placeholder. To load from file, use:
	// 1. Read file contents into a string
	// 2. Create strings.NewReader with the contents
	// 3. Call NewProfileFromReader
	return nil, fmt.Errorf("use NewProfileFromReader with file contents instead")
}
