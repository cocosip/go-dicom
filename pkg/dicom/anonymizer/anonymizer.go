// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package anonymizer

import (
	"fmt"
	"regexp"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// SecurityProfileOptions represents profile options as described in DICOM PS 3.15
// See: http://dicom.nema.org/medical/dicom/current/output/chtml/part15/PS3.15.html
type SecurityProfileOptions int

const (
	BasicProfile SecurityProfileOptions = 1 << iota
	RetainSafePrivate
	RetainUIDs
	RetainDeviceIdent
	RetainInstitutionIdent
	RetainPatientChars
	RetainLongFullDates
	RetainLongModifDates
	CleanDesc
	CleanStructdCont
	CleanGraph
)

// SecurityProfileAction represents profile actions per tag as described in DICOM PS 3.15
type SecurityProfileAction byte

const (
	// ActionD - Replace with a non-zero length value that may be a dummy value and consistent with the VR
	ActionD SecurityProfileAction = 1 << iota
	// ActionZ - Replace with a zero length value, or a non-zero length value that may be a dummy value and consistent with the VR
	ActionZ
	// ActionX - Remove
	ActionX
	// ActionK - Keep (unchanged for non-sequence attributes, cleaned for sequences)
	ActionK
	// ActionC - Clean, that is replace with values of similar meaning known not to contain identifying information and consistent with the VR
	ActionC
	// ActionU - Replace with a non-zero length UID that is internally consistent within a set of Instances
	ActionU
)

// String returns the string representation of the action
func (a SecurityProfileAction) String() string {
	switch a {
	case ActionD:
		return "D"
	case ActionZ:
		return "Z"
	case ActionX:
		return "X"
	case ActionK:
		return "K"
	case ActionC:
		return "C"
	case ActionU:
		return "U"
	default:
		return "Unknown"
	}
}

// SecurityProfile represents a mapping from tag patterns to anonymization actions
type SecurityProfile struct {
	rules       []profileRule
	PatientName string // Optional replacement patient name
	PatientID   string // Optional replacement patient ID
}

type profileRule struct {
	pattern *regexp.Regexp
	action  SecurityProfileAction
}

// NewSecurityProfile creates a new security profile with the specified options
func NewSecurityProfile(options SecurityProfileOptions) *SecurityProfile {
	profile := &SecurityProfile{
		rules: make([]profileRule, 0),
	}
	profile.loadDefaultProfile(options)
	return profile
}

// AddRule adds a custom rule to the profile
// The pattern is case-insensitive
func (sp *SecurityProfile) AddRule(pattern string, action SecurityProfileAction) error {
	re, err := regexp.Compile("(?i)" + pattern)
	if err != nil {
		return fmt.Errorf("invalid pattern %q: %w", pattern, err)
	}
	sp.rules = append(sp.rules, profileRule{
		pattern: re,
		action:  action,
	})
	return nil
}

// FindAction finds the action for a given tag
func (sp *SecurityProfile) FindAction(t *tag.Tag) (SecurityProfileAction, bool) {
	tagStr := t.String() // Format: (GGGG,EEEE)
	// Remove parentheses for matching
	tagStr = tagStr[1 : len(tagStr)-1] // "GGGG,EEEE"

	for _, rule := range sp.rules {
		if rule.pattern.MatchString(tagStr) {
			return rule.action, true
		}
	}
	return 0, false
}

// Anonymizer performs anonymization actions on DICOM datasets
type Anonymizer struct {
	Profile      *SecurityProfile
	ReplacedUIDs map[string]string // Context for UID replacement consistency
}

// NewAnonymizer creates a new Anonymizer with the specified profile
// If profile is nil, uses the default BasicProfile
func NewAnonymizer(profile *SecurityProfile) *Anonymizer {
	if profile == nil {
		profile = NewSecurityProfile(BasicProfile)
	}
	return &Anonymizer{
		Profile:      profile,
		ReplacedUIDs: make(map[string]string),
	}
}

// AnonymizeInPlace anonymizes a dataset without cloning
func (a *Anonymizer) AnonymizeInPlace(ds *dataset.Dataset) error {
	var toRemove []element.Element

	for _, elem := range ds.Elements() {
		action, hasAction := a.Profile.FindAction(elem.Tag())

		// Handle sequences - recursively anonymize
		if seq, ok := elem.(*dataset.Sequence); ok {
			if !hasAction || action == ActionK {
				for i := 0; i < seq.Count(); i++ {
					item := seq.GetItem(i)
					if err := a.AnonymizeInPlace(item); err != nil {
						return err
					}
				}
			}
		}

		// Apply action if found
		if hasAction {
			if err := a.applyAction(ds, elem, action); err != nil {
				return err
			}
			if action == ActionX {
				toRemove = append(toRemove, elem)
			}
		}

		// Handle special replacement fields
		if elem.Tag().Equals(tag.PatientName) && a.Profile.PatientName != "" {
			if err := a.replaceString(ds, elem, a.Profile.PatientName); err != nil {
				return fmt.Errorf("failed to replace patient name: %w", err)
			}
		} else if elem.Tag().Equals(tag.PatientID) && a.Profile.PatientID != "" {
			if err := a.replaceString(ds, elem, a.Profile.PatientID); err != nil {
				return fmt.Errorf("failed to replace patient ID: %w", err)
			}
		}
	}

	// Remove marked elements
	for _, elem := range toRemove {
		ds.Remove(elem.Tag())
	}

	return nil
}

// Anonymize clones and anonymizes a dataset
func (a *Anonymizer) Anonymize(ds *dataset.Dataset) (*dataset.Dataset, error) {
	clone := ds.Clone()
	if err := a.AnonymizeInPlace(clone); err != nil {
		return nil, err
	}
	return clone, nil
}

// applyAction applies the specified action to an element
func (a *Anonymizer) applyAction(ds *dataset.Dataset, elem element.Element, action SecurityProfileAction) error {
	vrValue := elem.ValueRepresentation()

	switch action {
	case ActionU, ActionC, ActionD:
		if vrValue.Code() == vr.CodeUI {
			return a.replaceUID(ds, elem)
		} else if isStringVR(vrValue) {
			return a.replaceString(ds, elem, "ANONYMOUS")
		}
		return a.blankItem(ds, elem, true)
	case ActionK:
		// Keep - do nothing
	case ActionX:
		// Remove - will be handled by caller
	case ActionZ:
		return a.blankItem(ds, elem, false)
	default:
		return fmt.Errorf("unknown action %v", action)
	}

	return nil
}

// replaceUID replaces a UID element with a new anonymized UID
func (a *Anonymizer) replaceUID(ds *dataset.Dataset, elem element.Element) error {
	// Get the old UID value
	oldUID, ok := ds.GetString(elem.Tag())
	if !ok || oldUID == "" {
		return nil // No need to replace empty values
	}

	var newUID string
	if replaced, exists := a.ReplacedUIDs[oldUID]; exists {
		// Use previously replaced UID for consistency
		newUID = replaced
	} else {
		// Generate new UID
		generatedUID := uid.GenerateDerivedFromUUID()
		newUID = generatedUID.UID()
		a.ReplacedUIDs[oldUID] = newUID
	}

	// Replace in dataset
	if err := ds.AddOrUpdate(element.NewString(elem.Tag(), vr.UI, []string{newUID})); err != nil {
		return fmt.Errorf("failed to update UID element: %w", err)
	}
	return nil
}

// replaceString replaces a string element with a new value
func (a *Anonymizer) replaceString(ds *dataset.Dataset, elem element.Element, newValue string) error {
	return ds.AddOrUpdate(element.NewString(elem.Tag(), elem.ValueRepresentation(), []string{newValue}))
}

// blankItem blanks an item to a value suitable for the element type
func (a *Anonymizer) blankItem(ds *dataset.Dataset, elem element.Element, nonZeroLength bool) error {
	t := elem.Tag()
	vrValue := elem.ValueRepresentation()
	vrCode := vrValue.Code()

	// Sequence - replace with empty sequence
	if vrCode == vr.CodeSQ {
		return ds.AddOrUpdate(dataset.NewSequence(t))
	}

	// String types - replace with empty string or dummy
	if isStringVR(vrValue) {
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewString(t, vrValue, []string{""}))
		}
		return ds.AddOrUpdate(element.NewString(t, vrValue, []string{}))
	}

	// Binary types (OB, OW, OD, OF, OL, OV, UN) - replace with empty buffer
	if vrCode == vr.CodeOB || vrCode == vr.CodeOW || vrCode == vr.CodeOD ||
		vrCode == vr.CodeOF || vrCode == vr.CodeOL || vrCode == vr.CodeOV ||
		vrCode == vr.CodeUN {
		return ds.AddOrUpdate(element.NewOtherByte(t, []byte{}))
	}

	// Numeric types - replace with zero value or empty array
	switch vrCode {
	case vr.CodeUS:
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewUnsignedShort(t, []uint16{0}))
		}
		return ds.AddOrUpdate(element.NewUnsignedShort(t, []uint16{}))
	case vr.CodeUL:
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewUnsignedLong(t, []uint32{0}))
		}
		return ds.AddOrUpdate(element.NewUnsignedLong(t, []uint32{}))
	case vr.CodeSS:
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewSignedShort(t, []int16{0}))
		}
		return ds.AddOrUpdate(element.NewSignedShort(t, []int16{}))
	case vr.CodeSL:
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewSignedLong(t, []int32{0}))
		}
		return ds.AddOrUpdate(element.NewSignedLong(t, []int32{}))
	case vr.CodeFL:
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewFloat(t, []float32{0}))
		}
		return ds.AddOrUpdate(element.NewFloat(t, []float32{}))
	case vr.CodeFD:
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewDouble(t, []float64{0}))
		}
		return ds.AddOrUpdate(element.NewDouble(t, []float64{}))
	case vr.CodeAT:
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewAttributeTag(t, []*tag.Tag{tag.Item}))
		}
		return ds.AddOrUpdate(element.NewAttributeTag(t, []*tag.Tag{}))
	default:
		// For other types, create empty buffer
		return ds.AddOrUpdate(element.NewOtherByte(t, []byte{}))
	}
}

// isStringVR checks if a VR is a string type
func isStringVR(vrValue *vr.VR) bool {
	code := vrValue.Code()
	return code == vr.CodeAE || code == vr.CodeAS || code == vr.CodeCS ||
		code == vr.CodeDA || code == vr.CodeDS || code == vr.CodeDT ||
		code == vr.CodeIS || code == vr.CodeLO || code == vr.CodeLT ||
		code == vr.CodePN || code == vr.CodeSH || code == vr.CodeST ||
		code == vr.CodeTM || code == vr.CodeUC || code == vr.CodeUI ||
		code == vr.CodeUR || code == vr.CodeUT
}
