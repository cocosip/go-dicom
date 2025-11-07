// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package serialization

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// ToXML serializes a DICOM dataset to XML format (NativeDicomModel)
// according to DICOM Part 19: Native DICOM Model
func ToXML(ds *dataset.Dataset, opts ...XMLOption) ([]byte, error) {
	cfg := defaultXMLConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	w := &xmlWriter{
		config: cfg,
		buf:    &bytes.Buffer{},
	}

	if err := w.writeDataset(ds); err != nil {
		return nil, err
	}

	return w.buf.Bytes(), nil
}

// xmlWriter handles the XML serialization process
type xmlWriter struct {
	config xmlConfig
	buf    *bytes.Buffer
	indent int // current indentation level
}

// writeDataset writes a DICOM dataset as XML
func (w *xmlWriter) writeDataset(ds *dataset.Dataset) error {
	// XML declaration
	w.writeLine(`<?xml version="1.0" encoding="utf-8"?>`)
	w.writeLine(`<NativeDicomModel>`)
	w.indent++

	// Write all elements
	if err := w.writeElements(ds); err != nil {
		return err
	}

	w.indent--
	w.writeLine(`</NativeDicomModel>`)
	return nil
}

// writeElements writes all elements in a dataset
func (w *xmlWriter) writeElements(ds *dataset.Dataset) error {
	for _, elem := range ds.Elements() {
		if seq, ok := elem.(*dataset.Sequence); ok {
			if err := w.writeSequence(seq); err != nil {
				return err
			}
		} else {
			if err := w.writeElement(elem); err != nil {
				return err
			}
		}
	}
	return nil
}

// writeElement writes a single DICOM element as XML
func (w *xmlWriter) writeElement(elem element.Element) error {
	// Write opening DicomAttribute tag
	w.writeIndent()
	w.buf.WriteString(`<DicomAttribute`)
	w.writeAttribute("tag", fmt.Sprintf("%04X%04X", elem.Tag().Group(), elem.Tag().Element()))
	w.writeAttribute("vr", elem.ValueRepresentation().String())

	// Add keyword if available
	entryIface := elem.Tag().DictionaryEntry()
	if entryIface != nil {
		if entry, ok := entryIface.(*dict.Entry); ok && entry != nil {
			if kw := entry.Keyword(); kw != "" {
				w.writeAttribute("keyword", kw)
			}
		}
	}

	// Add private creator if present
	if elem.Tag().IsPrivate() {
		pc := elem.Tag().PrivateCreator()
		if pc != nil && pc.Creator() != "" {
			w.writeAttribute("privateCreator", pc.Creator())
		}
	}

	w.buf.WriteString(">")
	if w.config.indent != "" {
		w.buf.WriteString("\n")
	}

	// Write element value based on VR
	vrCode := elem.ValueRepresentation()
	if err := w.writeElementValue(elem, vrCode); err != nil {
		return err
	}

	// Write closing tag
	w.writeIndentedLine(`</DicomAttribute>`)
	return nil
}

// writeElementValue writes the value of an element based on its VR
func (w *xmlWriter) writeElementValue(elem element.Element, vrValue *vr.VR) error {
	vrCode := vrValue.Code()

	// Binary VRs - write as InlineBinary base64
	if vrCode == vr.CodeOB || vrCode == vr.CodeOD || vrCode == vr.CodeOF ||
		vrCode == vr.CodeOW || vrCode == vr.CodeOL || vrCode == vr.CodeOV || vrCode == vr.CodeUN {
		return w.writeBinaryValue(elem)
	}

	// PersonName - special structured format
	if vrCode == vr.CodePN {
		return w.writePersonNameValue(elem)
	}

	// All other VRs - write as Value elements
	return w.writeStringValues(elem)
}

// writeBinaryValue writes binary data as base64-encoded InlineBinary
func (w *xmlWriter) writeBinaryValue(elem element.Element) error {
	buf := elem.Buffer()
	if buf == nil || buf.Size() == 0 {
		return nil
	}

	data := buf.Data()
	encoded := base64.StdEncoding.EncodeToString(data)

	w.indent++
	w.writeIndent()
	w.buf.WriteString(`<InlineBinary>`)
	w.buf.WriteString(encoded)
	w.buf.WriteString(`</InlineBinary>`)
	if w.config.indent != "" {
		w.buf.WriteString("\n")
	}
	w.indent--

	return nil
}

// writePersonNameValue writes PersonName values in structured format
func (w *xmlWriter) writePersonNameValue(elem element.Element) error {
	w.indent++

	// Try to cast to String element to access GetValue
	strElem, ok := elem.(interface{ GetValue(int) string })
	if !ok {
		w.indent--
		return nil
	}

	for i := 0; i < elem.Count(); i++ {
		value := strElem.GetValue(i)
		if value == "" {
			continue
		}

		// Parse PersonName components (Alphabetic=Ideographic=Phonetic)
		components := strings.Split(value, "=")

		w.writeIndent()
		w.buf.WriteString(`<PersonName`)
		w.writeAttribute("number", fmt.Sprintf("%d", i+1))
		w.buf.WriteString(">")
		if w.config.indent != "" {
			w.buf.WriteString("\n")
		}

		w.indent++

		// Write each component group (usually just Alphabetic)
		for groupIdx, component := range components {
			if component == "" {
				continue
			}

			groupName := "Alphabetic"
			switch groupIdx {
			case 1:
				groupName = "Ideographic"
			case 2:
				groupName = "Phonetic"
			}

			// Parse component parts (FamilyName^GivenName^MiddleName^Prefix^Suffix)
			parts := strings.Split(component, "^")

			w.writeIndentedLine(fmt.Sprintf("<%s>", groupName))
			w.indent++

			// Write each part if not empty
			partNames := []string{"FamilyName", "GivenName", "MiddleName", "NamePrefix", "NameSuffix"}
			for partIdx, part := range parts {
				if part != "" && partIdx < len(partNames) {
					w.writeIndent()
					w.buf.WriteString(fmt.Sprintf("<%s>%s</%s>", partNames[partIdx], escapeXML(part), partNames[partIdx]))
					if w.config.indent != "" {
						w.buf.WriteString("\n")
					}
				}
			}

			w.indent--
			w.writeIndentedLine(fmt.Sprintf("</%s>", groupName))
		}

		w.indent--
		w.writeIndentedLine(`</PersonName>`)
	}

	w.indent--
	return nil
}

// writeStringValues writes string values as Value elements
func (w *xmlWriter) writeStringValues(elem element.Element) error {
	w.indent++

	// Get string values from element based on type
	var values []string

	// Try string element first
	if strElem, ok := elem.(interface{ GetValue(int) string }); ok {
		for i := 0; i < elem.Count(); i++ {
			values = append(values, strElem.GetValue(i))
		}
	} else {
		// For numeric elements, try type-specific GetValues methods
		values = w.getNumericValuesAsStrings(elem)
	}

	// Write values
	for i, value := range values {
		if value == "" {
			continue
		}

		w.writeIndent()
		w.buf.WriteString(`<Value`)
		w.writeAttribute("number", fmt.Sprintf("%d", i+1))
		w.buf.WriteString(">")
		w.buf.WriteString(escapeXML(value))
		w.buf.WriteString(`</Value>`)
		if w.config.indent != "" {
			w.buf.WriteString("\n")
		}
	}

	w.indent--
	return nil
}

// getNumericValuesAsStrings extracts numeric values and converts them to strings
func (w *xmlWriter) getNumericValuesAsStrings(elem element.Element) []string {
	var values []string

	// Try US (unsigned short)
	if usElem, ok := elem.(interface{ GetValues() ([]uint16, error) }); ok {
		if vals, err := usElem.GetValues(); err == nil {
			for _, v := range vals {
				values = append(values, fmt.Sprintf("%d", v))
			}
			return values
		}
	}

	// Try UL (unsigned long)
	if ulElem, ok := elem.(interface{ GetValues() ([]uint32, error) }); ok {
		if vals, err := ulElem.GetValues(); err == nil {
			for _, v := range vals {
				values = append(values, fmt.Sprintf("%d", v))
			}
			return values
		}
	}

	// Try SS (signed short)
	if ssElem, ok := elem.(interface{ GetValues() ([]int16, error) }); ok {
		if vals, err := ssElem.GetValues(); err == nil {
			for _, v := range vals {
				values = append(values, fmt.Sprintf("%d", v))
			}
			return values
		}
	}

	// Try SL (signed long)
	if slElem, ok := elem.(interface{ GetValues() ([]int32, error) }); ok {
		if vals, err := slElem.GetValues(); err == nil {
			for _, v := range vals {
				values = append(values, fmt.Sprintf("%d", v))
			}
			return values
		}
	}

	// Try FL (float)
	if flElem, ok := elem.(interface{ GetValues() ([]float32, error) }); ok {
		if vals, err := flElem.GetValues(); err == nil {
			for _, v := range vals {
				values = append(values, fmt.Sprintf("%g", v))
			}
			return values
		}
	}

	// Try FD (double)
	if fdElem, ok := elem.(interface{ GetValues() ([]float64, error) }); ok {
		if vals, err := fdElem.GetValues(); err == nil {
			for _, v := range vals {
				values = append(values, fmt.Sprintf("%g", v))
			}
			return values
		}
	}

	// Try SV (signed very long)
	if svElem, ok := elem.(interface{ GetValues() ([]int64, error) }); ok {
		if vals, err := svElem.GetValues(); err == nil {
			for _, v := range vals {
				values = append(values, fmt.Sprintf("%d", v))
			}
			return values
		}
	}

	// Try UV (unsigned very long)
	if uvElem, ok := elem.(interface{ GetValues() ([]uint64, error) }); ok {
		if vals, err := uvElem.GetValues(); err == nil {
			for _, v := range vals {
				values = append(values, fmt.Sprintf("%d", v))
			}
			return values
		}
	}

	return values
}

// writeSequence writes a DICOM sequence as XML
func (w *xmlWriter) writeSequence(seq *dataset.Sequence) error {
	// Write opening DicomAttribute tag
	w.writeIndent()
	w.buf.WriteString(`<DicomAttribute`)
	w.writeAttribute("tag", fmt.Sprintf("%04X%04X", seq.Tag().Group(), seq.Tag().Element()))
	w.writeAttribute("vr", seq.ValueRepresentation().String())

	// Add keyword if available
	entryIface := seq.Tag().DictionaryEntry()
	if entryIface != nil {
		if entry, ok := entryIface.(*dict.Entry); ok && entry != nil {
			if kw := entry.Keyword(); kw != "" {
				w.writeAttribute("keyword", kw)
			}
		}
	}

	w.buf.WriteString(">")
	if w.config.indent != "" {
		w.buf.WriteString("\n")
	}

	w.indent++

	// Write each item
	for i := 0; i < seq.Count(); i++ {
		item := seq.GetItem(i)
		if item == nil {
			continue
		}

		w.writeIndent()
		w.buf.WriteString(`<Item`)
		w.writeAttribute("number", fmt.Sprintf("%d", i+1))
		w.buf.WriteString(">")
		if w.config.indent != "" {
			w.buf.WriteString("\n")
		}

		w.indent++
		if err := w.writeElements(item); err != nil {
			return err
		}
		w.indent--

		w.writeIndentedLine(`</Item>`)
	}

	w.indent--
	w.writeIndentedLine(`</DicomAttribute>`)

	return nil
}

// writeAttribute writes an XML attribute
func (w *xmlWriter) writeAttribute(name, value string) {
	w.buf.WriteString(` `)
	w.buf.WriteString(name)
	w.buf.WriteString(`="`)
	w.buf.WriteString(escapeXML(value))
	w.buf.WriteString(`"`)
}

// writeIndent writes the current indentation
func (w *xmlWriter) writeIndent() {
	if w.config.indent == "" {
		return
	}
	for i := 0; i < w.indent; i++ {
		w.buf.WriteString(w.config.indent)
	}
}

// writeLine writes a line with newline
func (w *xmlWriter) writeLine(s string) {
	w.buf.WriteString(s)
	if w.config.indent != "" {
		w.buf.WriteString("\n")
	}
}

// writeIndentedLine writes an indented line
func (w *xmlWriter) writeIndentedLine(s string) {
	w.writeIndent()
	w.writeLine(s)
}

// escapeXML escapes special XML characters
func escapeXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&apos;")
	return s
}
