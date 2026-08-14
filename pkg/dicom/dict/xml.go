// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dict

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vm"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

type xmlDictionaryCollection struct {
	XMLName      xml.Name
	Dictionaries []xmlDictionary     `xml:"dictionary"`
	Unknown      []xmlUnknownElement `xml:",any"`
	Attributes   []xml.Attr          `xml:",any,attr"`
	Text         string              `xml:",chardata"`
}

type xmlDictionary struct {
	XMLName    xml.Name
	Version    string              `xml:"version,attr"`
	Creator    string              `xml:"creator,attr"`
	Tags       []xmlDictionaryTag  `xml:"tag"`
	UIDs       []xmlIgnoredUID     `xml:"uid"`
	Unknown    []xmlUnknownElement `xml:",any"`
	Attributes []xml.Attr          `xml:",any,attr"`
	Text       string              `xml:",chardata"`
}

type xmlDictionaryTag struct {
	XMLName    xml.Name
	Group      string              `xml:"group,attr"`
	Element    string              `xml:"element,attr"`
	Keyword    string              `xml:"keyword,attr"`
	VR         string              `xml:"vr,attr"`
	VM         string              `xml:"vm,attr"`
	Retired    string              `xml:"retired,attr"`
	Name       string              `xml:",chardata"`
	Unknown    []xmlUnknownElement `xml:",any"`
	Attributes []xml.Attr          `xml:",any,attr"`
}

type xmlUnknownElement struct {
	XMLName xml.Name
}

type xmlIgnoredUID struct {
	XMLName xml.Name
}

func (dictionary *xmlDictionary) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	dictionary.XMLName = start.Name
	for _, attribute := range start.Attr {
		switch attribute.Name.Local {
		case "version":
			dictionary.Version = attribute.Value
		case "creator":
			dictionary.Creator = attribute.Value
		default:
			dictionary.Attributes = append(dictionary.Attributes, attribute)
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		switch value := token.(type) {
		case xml.CharData:
			dictionary.Text += string(value)
		case xml.StartElement:
			switch value.Name.Local {
			case "tag":
				contextTag := xmlDictionaryTagFromStart(value)
				var dictionaryTag xmlDictionaryTag
				if err := decoder.DecodeElement(&dictionaryTag, &value); err != nil {
					return fmt.Errorf(
						"%s: %w",
						xmlTagContext(strings.TrimSpace(dictionary.Creator), len(dictionary.Tags)+1, contextTag),
						err,
					)
				}
				dictionary.Tags = append(dictionary.Tags, dictionaryTag)
			case "uid":
				if err := decoder.Skip(); err != nil {
					return err
				}
				dictionary.UIDs = append(dictionary.UIDs, xmlIgnoredUID{XMLName: value.Name})
			default:
				if err := decoder.Skip(); err != nil {
					return err
				}
				dictionary.Unknown = append(dictionary.Unknown, xmlUnknownElement{XMLName: value.Name})
			}
		case xml.EndElement:
			if value.Name == start.Name {
				return nil
			}
		}
	}
}

func xmlDictionaryTagFromStart(start xml.StartElement) xmlDictionaryTag {
	dictionaryTag := xmlDictionaryTag{XMLName: start.Name}
	for _, attribute := range start.Attr {
		switch attribute.Name.Local {
		case "group":
			dictionaryTag.Group = attribute.Value
		case "element":
			dictionaryTag.Element = attribute.Value
		}
	}
	return dictionaryTag
}

type parsedXMLDictionary struct {
	creator string
	entries []*Entry
}

// NewFromXML creates a dictionary from a fo-dicom-compatible XML stream.
func NewFromXML(reader io.Reader) (*Dictionary, error) {
	dictionary := New()
	if err := dictionary.LoadXML(reader); err != nil {
		return nil, err
	}
	return dictionary, nil
}

// LoadXML validates and merges a fo-dicom-compatible XML dictionary stream.
// Exact and masked duplicates use last-loaded-wins semantics. The receiver is
// unchanged when parsing or validation fails.
func (d *Dictionary) LoadXML(reader io.Reader) error {
	if d == nil {
		return fmt.Errorf("dictionary is nil")
	}
	if reader == nil {
		return fmt.Errorf("dictionary XML reader is nil")
	}

	parsed, err := parseDictionaryXML(reader)
	if err != nil {
		return err
	}

	for _, parsedDictionary := range parsed {
		target := d
		if parsedDictionary.creator != "" {
			target = d.GetPrivateDictionary(parsedDictionary.creator)
		}
		for _, entry := range parsedDictionary.entries {
			target.Add(entry)
		}
	}
	return nil
}

func parseDictionaryXML(reader io.Reader) ([]parsedXMLDictionary, error) {
	decoder := xml.NewDecoder(reader)
	start, err := firstStartElement(decoder)
	if err != nil {
		return nil, err
	}

	var dictionaries []xmlDictionary
	switch start.Name.Local {
	case "dictionary":
		var dictionary xmlDictionary
		if err := decoder.DecodeElement(&dictionary, &start); err != nil {
			return nil, fmt.Errorf("parse dictionary XML: %w", err)
		}
		dictionaries = []xmlDictionary{dictionary}
	case "dictionaries":
		var collection xmlDictionaryCollection
		if err := decoder.DecodeElement(&collection, &start); err != nil {
			return nil, fmt.Errorf("parse dictionary XML: %w", err)
		}
		if len(collection.Attributes) > 0 {
			return nil, fmt.Errorf("<dictionaries>: unexpected attribute %q", collection.Attributes[0].Name.Local)
		}
		if len(collection.Unknown) > 0 {
			return nil, fmt.Errorf("<dictionaries>: unexpected <%s> element", collection.Unknown[0].XMLName.Local)
		}
		if strings.TrimSpace(collection.Text) != "" {
			return nil, fmt.Errorf("<dictionaries>: unexpected text content")
		}
		dictionaries = collection.Dictionaries
	default:
		return nil, fmt.Errorf("expected <dictionary> or <dictionaries> root, got <%s>", start.Name.Local)
	}

	if err := ensureXMLComplete(decoder); err != nil {
		return nil, err
	}

	parsed := make([]parsedXMLDictionary, 0, len(dictionaries))
	for _, dictionary := range dictionaries {
		entries, err := parseXMLDictionary(dictionary)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, parsedXMLDictionary{
			creator: strings.TrimSpace(dictionary.Creator),
			entries: entries,
		})
	}
	return parsed, nil
}

func firstStartElement(decoder *xml.Decoder) (xml.StartElement, error) {
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return xml.StartElement{}, fmt.Errorf("dictionary XML is empty")
		}
		if err != nil {
			return xml.StartElement{}, fmt.Errorf("parse dictionary XML: %w", err)
		}
		switch value := token.(type) {
		case xml.StartElement:
			return value, nil
		case xml.CharData:
			text := strings.TrimPrefix(string(value), "\ufeff")
			if strings.TrimSpace(text) != "" {
				return xml.StartElement{}, fmt.Errorf("unexpected text before dictionary root")
			}
		}
	}
}

func ensureXMLComplete(decoder *xml.Decoder) error {
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("parse dictionary XML: %w", err)
		}
		switch value := token.(type) {
		case xml.StartElement:
			return fmt.Errorf("unexpected <%s> element after dictionary root", value.Name.Local)
		case xml.CharData:
			if strings.TrimSpace(string(value)) != "" {
				return fmt.Errorf("unexpected text after dictionary root")
			}
		}
	}
}

func parseXMLDictionary(dictionary xmlDictionary) ([]*Entry, error) {
	if len(dictionary.Attributes) > 0 {
		return nil, fmt.Errorf("<dictionary>: unexpected attribute %q", dictionary.Attributes[0].Name.Local)
	}
	if len(dictionary.Unknown) > 0 {
		return nil, fmt.Errorf("<dictionary>: unexpected <%s> element", dictionary.Unknown[0].XMLName.Local)
	}
	if strings.TrimSpace(dictionary.Text) != "" {
		return nil, fmt.Errorf("<dictionary>: unexpected text content")
	}

	creatorName := strings.TrimSpace(dictionary.Creator)
	entries := make([]*Entry, 0, len(dictionary.Tags))
	for index, xmlTag := range dictionary.Tags {
		entry, err := parseXMLDictionaryTag(xmlTag, creatorName)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", xmlTagContext(creatorName, index+1, xmlTag), err)
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func xmlTagContext(creator string, index int, xmlTag xmlDictionaryTag) string {
	prefix := fmt.Sprintf("tag %d", index)
	if creator != "" {
		prefix = fmt.Sprintf("dictionary %q %s", creator, prefix)
	}
	if xmlTag.Group != "" && xmlTag.Element != "" {
		prefix += fmt.Sprintf(" (%s,%s)", xmlTag.Group, xmlTag.Element)
	}
	return prefix
}

func parseXMLDictionaryTag(xmlTag xmlDictionaryTag, creatorName string) (*Entry, error) {
	if len(xmlTag.Attributes) > 0 {
		return nil, fmt.Errorf("unexpected attribute %q", xmlTag.Attributes[0].Name.Local)
	}
	if len(xmlTag.Unknown) > 0 {
		return nil, fmt.Errorf("unexpected <%s> element", xmlTag.Unknown[0].XMLName.Local)
	}
	if xmlTag.Group == "" {
		return nil, fmt.Errorf("missing group")
	}
	if xmlTag.Element == "" {
		return nil, fmt.Errorf("missing element")
	}
	if err := validateXMLTagPart("group", xmlTag.Group); err != nil {
		return nil, err
	}
	if err := validateXMLTagPart("element", xmlTag.Element); err != nil {
		return nil, err
	}
	if xmlTag.VM == "" {
		return nil, fmt.Errorf("missing VM")
	}

	valueMultiplicity, err := parseXMLVM(xmlTag.VM)
	if err != nil {
		return nil, err
	}
	valueRepresentations, err := parseXMLVRs(xmlTag.VR)
	if err != nil {
		return nil, err
	}
	retired, err := parseXMLRetired(xmlTag.Retired)
	if err != nil {
		return nil, err
	}
	name := strings.TrimSpace(xmlTag.Name)
	if name == "" {
		return nil, fmt.Errorf("missing name")
	}

	privateCreator := (*tag.PrivateCreator)(nil)
	if creatorName != "" {
		privateCreator = tag.NewPrivateCreator(creatorName)
	}
	pattern := fmt.Sprintf("(%s,%s)", xmlTag.Group, xmlTag.Element)
	if strings.ContainsAny(strings.ToLower(pattern), "x") {
		maskedTag, err := tag.ParseMaskedTag(pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid masked tag: %w", err)
		}
		maskedTag.Tag().SetPrivateCreator(privateCreator)
		return &Entry{
			tag:                  maskedTag.Tag(),
			maskTag:              maskedTag,
			name:                 name,
			keyword:              strings.TrimSpace(xmlTag.Keyword),
			valueRepresentations: valueRepresentations,
			valueMultiplicity:    valueMultiplicity,
			isRetired:            retired,
		}, nil
	}

	exactTag, err := tag.Parse(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid tag: %w", err)
	}
	exactTag.SetPrivateCreator(privateCreator)
	return &Entry{
		tag:                  exactTag,
		name:                 name,
		keyword:              strings.TrimSpace(xmlTag.Keyword),
		valueRepresentations: valueRepresentations,
		valueMultiplicity:    valueMultiplicity,
		isRetired:            retired,
	}, nil
}

func validateXMLTagPart(name, value string) error {
	if len(value) != 4 {
		return fmt.Errorf("invalid %s %q: expected four hexadecimal or wildcard digits", name, value)
	}
	for _, character := range value {
		if (character >= '0' && character <= '9') ||
			(character >= 'a' && character <= 'f') ||
			(character >= 'A' && character <= 'F') ||
			character == 'x' || character == 'X' {
			continue
		}
		return fmt.Errorf("invalid %s %q: unexpected character %q", name, value, character)
	}
	return nil
}

func parseXMLVRs(value string) ([]*vr.VR, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return []*vr.VR{vr.None}, nil
	}

	codes, err := splitXMLVRs(value)
	if err != nil {
		return nil, err
	}
	valueRepresentations := make([]*vr.VR, 0, len(codes))
	for _, code := range codes {
		valueRepresentation, err := vr.Parse(code)
		if err != nil {
			return nil, fmt.Errorf("invalid VR %q: %w", code, err)
		}
		valueRepresentations = append(valueRepresentations, valueRepresentation)
	}
	return valueRepresentations, nil
}

func splitXMLVRs(value string) ([]string, error) {
	codes := make([]string, 0, 1)
	start := 0
	for index, character := range value {
		if !isXMLVRSeparator(character) {
			continue
		}
		code := strings.TrimSpace(value[start:index])
		if code == "" {
			return nil, fmt.Errorf("invalid VR list %q: empty component", value)
		}
		codes = append(codes, code)
		start = index + 1
	}
	code := strings.TrimSpace(value[start:])
	if code == "" {
		return nil, fmt.Errorf("invalid VR list %q: empty component", value)
	}
	return append(codes, code), nil
}

func isXMLVRSeparator(character rune) bool {
	switch character {
	case '_', '/', '\\', ',', '|':
		return true
	default:
		return false
	}
}

func parseXMLVM(value string) (*vm.VM, error) {
	value = strings.TrimSpace(value)
	parts := strings.Split(value, " or ")
	var selected *vm.VM
	for _, part := range parts {
		part = strings.TrimSpace(part)
		parsed, err := vm.Parse(part)
		if err != nil {
			return nil, fmt.Errorf("invalid VM %q: %w", value, err)
		}
		if selected == nil {
			selected = parsed
		}
	}
	return selected, nil
}

func parseXMLRetired(value string) (bool, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return false, nil
	}
	if !strings.EqualFold(value, "true") && !strings.EqualFold(value, "false") {
		return false, fmt.Errorf("invalid retired value %q", value)
	}
	retired, err := strconv.ParseBool(strings.ToLower(value))
	if err != nil {
		return false, fmt.Errorf("invalid retired value %q: %w", value, err)
	}
	return retired, nil
}
