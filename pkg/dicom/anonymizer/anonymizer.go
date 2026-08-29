// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package anonymizer provides DICOM anonymizer.
package anonymizer

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// SecurityProfileOptions represents profile options as described in DICOM PS 3.15
// See: http://dicom.nema.org/medical/dicom/current/output/chtml/part15/PS3.15.html
type SecurityProfileOptions int

const (
	// BasicProfile - Basic profile options
	BasicProfile SecurityProfileOptions = 1 << iota
	// RetainSafePrivate - Retain private tags with safe VRs
	RetainSafePrivate
	// RetainUIDs - Retain UIDs
	RetainUIDs
	// RetainDeviceIdent - Retain device identifiers
	RetainDeviceIdent
	// RetainInstitutionIdent - Retain institution identifiers
	RetainInstitutionIdent
	// RetainPatientChars - Retain patient characters
	RetainPatientChars
	// RetainLongFullDates - Retain long full dates
	RetainLongFullDates
	// RetainLongModifDates - Retain long modification dates
	RetainLongModifDates
	// CleanDesc - Clean descriptions
	CleanDesc
	// CleanStructdCont - Clean structured content
	CleanStructdCont
	// CleanGraph - Clean graphs
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
	rules          []profileRule
	exactOverrides map[uint32]SecurityProfileAction
	PatientName    string // Optional replacement patient name
	PatientID      string // Optional replacement patient ID
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

// OverrideAction applies action to an exact tag before all profile rules.
//
// ActionK retains the original value and can preserve identifiers that make
// datasets linkable, so callers must explicitly authorize its use.
func (sp *SecurityProfile) OverrideAction(t *tag.Tag, action SecurityProfileAction) {
	if sp.exactOverrides == nil {
		sp.exactOverrides = make(map[uint32]SecurityProfileAction)
	}
	sp.exactOverrides[t.ToUint32()] = action
	tagStr := t.String()
	tagStr = tagStr[1 : len(tagStr)-1]
	pattern := regexp.MustCompile("(?i)^" + regexp.QuoteMeta(tagStr) + "$")
	sp.rules = append([]profileRule{{pattern: pattern, action: action}}, sp.rules...)
}

func (sp *SecurityProfile) exactOverrideAction(t *tag.Tag) (SecurityProfileAction, bool) {
	if sp == nil || sp.exactOverrides == nil {
		return 0, false
	}
	action, ok := sp.exactOverrides[t.ToUint32()]
	return action, ok
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
	state := &anonymizationState{retainedOverrides: make(map[uint32]*tag.Tag)}
	return a.anonymizeInPlace(ds, true, state)
}

type anonymizationState struct {
	retainedOverrides map[uint32]*tag.Tag
}

func (a *Anonymizer) anonymizeInPlace(ds *dataset.Dataset, topLevel bool, state *anonymizationState) error {
	var toRemove []element.Element
	var patientNameElement element.Element
	var patientIDElement element.Element

	for _, elem := range ds.Elements() {
		if elem.Tag().Equals(tag.PatientName) && a.Profile.PatientName != "" {
			patientNameElement = elem
		}
		if elem.Tag().Equals(tag.PatientID) && a.Profile.PatientID != "" {
			patientIDElement = elem
		}

		action, hasAction := a.Profile.FindAction(elem.Tag())
		if override, exact := a.Profile.exactOverrideAction(elem.Tag()); exact && override == ActionK &&
			hasAction && action == ActionK && !a.hasConfiguredReplacement(elem.Tag()) {
			state.retainedOverrides[elem.Tag().ToUint32()] = elem.Tag()
		}

		skipAction, err := a.anonymizeSequence(elem, hasAction, action, state)
		if err != nil {
			return err
		}
		if skipAction {
			continue
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
	}

	// Remove marked elements first, so patient name/id replacements
	// below are not immediately deleted when the profile uses ActionX.
	for _, elem := range toRemove {
		ds.Remove(elem.Tag())
	}

	// Re-add configured replacements only when the source attribute existed.
	if patientNameElement != nil {
		if err := a.replaceString(ds, patientNameElement, a.Profile.PatientName); err != nil {
			return fmt.Errorf("failed to replace patient name: %w", err)
		}
	}
	if patientIDElement != nil {
		if err := a.replaceString(ds, patientIDElement, a.Profile.PatientID); err != nil {
			return fmt.Errorf("failed to replace patient ID: %w", err)
		}
	}

	if topLevel {
		return addDeidentificationDeclaration(ds, state)
	}

	return nil
}

func (a *Anonymizer) anonymizeSequence(elem element.Element, hasAction bool, action SecurityProfileAction, state *anonymizationState) (bool, error) {
	seq, ok := elem.(*dataset.Sequence)
	if !ok {
		return false, nil
	}
	if hasAction && action != ActionK && action != ActionC && action != ActionD && action != ActionU {
		return false, nil
	}
	for i := 0; i < seq.Count(); i++ {
		if err := a.anonymizeInPlace(seq.GetItem(i), false, state); err != nil {
			return false, err
		}
	}
	return hasAction && action != ActionK, nil
}

func addDeidentificationDeclaration(ds *dataset.Dataset, state *anonymizationState) error {
	identityRemoved := "YES"
	methods := []string{"go-dicom profile-based de-identification"}
	if len(state.retainedOverrides) > 0 {
		identityRemoved = "NO"
		retained := make([]string, 0, len(state.retainedOverrides))
		for _, retainedTag := range state.retainedOverrides {
			retained = append(retained, retainedTag.String())
		}
		sort.Strings(retained)
		methods = make([]string, 0, len(retained)+1)
		for _, retainedTag := range retained {
			methods = append(methods, "go-dicom retained explicit override: "+retainedTag)
		}
		methods = append(methods, "go-dicom profile-based de-identification")
	}
	if err := ds.AddOrUpdate(element.NewString(tag.PatientIdentityRemoved, vr.CS, []string{identityRemoved})); err != nil {
		return fmt.Errorf("failed to set Patient Identity Removed: %w", err)
	}
	if err := ds.AddOrUpdate(element.NewString(tag.DeidentificationMethod, vr.LO, methods)); err != nil {
		return fmt.Errorf("failed to set De-identification Method: %w", err)
	}
	return nil
}

func (a *Anonymizer) hasConfiguredReplacement(t *tag.Tag) bool {
	return t.Equals(tag.PatientName) && a.Profile.PatientName != "" ||
		t.Equals(tag.PatientID) && a.Profile.PatientID != ""
}

// Anonymize clones and anonymizes a dataset
func (a *Anonymizer) Anonymize(ds *dataset.Dataset) (*dataset.Dataset, error) {
	clone, err := ds.DeepCloneChecked()
	if err != nil {
		return nil, fmt.Errorf("clone dataset for anonymization: %w", err)
	}
	if err := a.AnonymizeInPlace(clone); err != nil {
		return nil, err
	}
	return clone, nil
}

// AnonymizeFileInPlace anonymizes a main Data Set and returns newly generated
// File Meta Information. The source metadata is used only as a transfer syntax
// fallback; optional AE titles, private information, and implementation
// identifiers are deliberately not copied.
func (a *Anonymizer) AnonymizeFileInPlace(ds *dataset.Dataset, source *dataset.FileMetaInformation) (*dataset.FileMetaInformation, error) {
	if ds == nil {
		return nil, fmt.Errorf("dataset is nil")
	}

	ts := ds.InternalTransferSyntax()
	if ts == nil && source != nil {
		if sourceTS, ok := source.TransferSyntax(); ok {
			ts = sourceTS
		}
	}
	if ts == nil {
		ts = transfer.ExplicitVRLittleEndian
	}
	if ds.InternalTransferSyntax() == nil {
		ds.SetInternalTransferSyntax(ts)
	}

	if err := a.AnonymizeInPlace(ds); err != nil {
		return nil, err
	}

	sopClassUID, ok := ds.GetString(tag.SOPClassUID)
	if !ok || sopClassUID == "" {
		return nil, fmt.Errorf("dataset is missing SOPClassUID")
	}
	sopInstanceUID, ok := ds.GetString(tag.SOPInstanceUID)
	if !ok || sopInstanceUID == "" {
		return nil, fmt.Errorf("dataset is missing SOPInstanceUID")
	}

	fresh := dataset.NewDefaultFileMetaInformation()
	if err := fresh.SetMediaStorageSOPClassUID(sopClassUID); err != nil {
		return nil, fmt.Errorf("set Media Storage SOP Class UID: %w", err)
	}
	if err := fresh.SetMediaStorageSOPInstanceUID(sopInstanceUID); err != nil {
		return nil, fmt.Errorf("set Media Storage SOP Instance UID: %w", err)
	}
	if err := fresh.SetTransferSyntax(ts); err != nil {
		return nil, fmt.Errorf("set Transfer Syntax UID: %w", err)
	}
	return fresh, nil
}

// applyAction applies the specified action to an element
func (a *Anonymizer) applyAction(ds *dataset.Dataset, elem element.Element, action SecurityProfileAction) error {
	vrValue := elem.ValueRepresentation()

	switch action {
	case ActionU, ActionC, ActionD:
		if vrValue.Code() == vr.CodeUI {
			return a.replaceUID(ds, elem)
		} else if isStringVR(vrValue) {
			return a.replaceString(ds, elem, dummyStringValue(vrValue))
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
			return ds.AddOrUpdate(element.NewString(t, vrValue, []string{dummyStringValue(vrValue)}))
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
	case vr.CodeSV:
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewSignedVeryLong(t, []int64{0}))
		}
		return ds.AddOrUpdate(element.NewSignedVeryLong(t, []int64{}))
	case vr.CodeUV:
		if nonZeroLength {
			return ds.AddOrUpdate(element.NewUnsignedVeryLong(t, []uint64{0}))
		}
		return ds.AddOrUpdate(element.NewUnsignedVeryLong(t, []uint64{}))
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

func dummyStringValue(vrValue *vr.VR) string {
	switch vrValue.Code() {
	case vr.CodeAS:
		return "000Y"
	case vr.CodeDA:
		return "19000101"
	case vr.CodeDS, vr.CodeIS:
		return "0"
	case vr.CodeDT:
		return "19000101000000"
	case vr.CodeTM:
		return "000000"
	case vr.CodeAE:
		return "ANON"
	case vr.CodeCS:
		return "ANONYMIZED"
	case vr.CodeUR:
		return "about:blank"
	default:
		return "ANONYMOUS"
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
