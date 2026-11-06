// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package dataset

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// Layout selects how Dataset.Format renders a dataset.
type Layout int

const (
	// LayoutSingleLine renders the dataset on one line, which is suitable for
	// structured log attributes. Sequence children are summarized.
	LayoutSingleLine Layout = iota
	// LayoutRows renders one element per line without column alignment.
	LayoutRows
	// LayoutDump renders aligned rows with byte length and dictionary keyword.
	LayoutDump
)

// FormatOptions controls Dataset.Format rendering.
type FormatOptions struct {
	// Layout selects the rendering style. The zero value is LayoutSingleLine.
	Layout Layout
	// MaxDepth limits sequence nesting for LayoutRows and LayoutDump. A
	// negative value means unlimited; zero does not descend into sequences.
	MaxDepth int
	// MaxElements stops rendering after this many element rows. Zero means
	// unlimited.
	MaxElements int
	// MaxValueLength limits the runes rendered for one element value. Text and
	// numeric values are truncated; binary and container values are summarized
	// by size or item count. Zero uses element.DefaultValueLimit.
	MaxValueLength int
	// HideValues renders tags and VRs without element values.
	HideValues bool
}

// Format renders the dataset for diagnostics, logs, and CLI dumps. It is a
// best-effort view: rendering stops at MaxElements or at the first traversal
// error, and sequence children are only expanded by LayoutRows and LayoutDump.
func (ds *Dataset) Format(opts FormatOptions) string {
	if ds == nil {
		return ""
	}
	if opts.MaxValueLength <= 0 {
		opts.MaxValueLength = element.DefaultValueLimit
	}
	switch opts.Layout {
	case LayoutRows:
		return strings.Join(formatLines(ds, opts, false), "\n")
	case LayoutDump:
		return strings.Join(formatLines(ds, opts, true), "\n")
	default:
		return formatSingleLine(ds, opts)
	}
}

func formatSingleLine(ds *Dataset, opts FormatOptions) string {
	var parts []string
	truncated := false
	err := Walk(ds, func(event WalkEvent) (WalkAction, error) {
		switch event.Kind {
		case WalkElement, WalkFragmentBegin:
			if formatLimitReached(len(parts), opts.MaxElements) {
				truncated = true
				return WalkStop, nil
			}
			parts = append(parts, formatElementLine(event.Element, 0, opts, false))
		case WalkSequenceBegin:
			sequence, ok := event.Element.(*Sequence)
			if !ok || sequence == nil {
				return WalkContinue, nil
			}
			if formatLimitReached(len(parts), opts.MaxElements) {
				truncated = true
				return WalkStop, nil
			}
			parts = append(parts, formatSequenceLine(sequence, 0, false))
			return WalkSkipChildren, nil
		}
		return WalkContinue, nil
	})
	if len(parts) == 0 {
		return ""
	}
	result := strings.Join(parts, " | ")
	if truncated || err != nil {
		result += " | ..."
	}
	return result
}

func formatLines(ds *Dataset, opts FormatOptions, align bool) []string {
	var lines []string
	elements := 0
	truncated := false
	err := Walk(ds, func(event WalkEvent) (WalkAction, error) {
		switch event.Kind {
		case WalkElement, WalkFragmentBegin:
			if formatLimitReached(elements, opts.MaxElements) {
				truncated = true
				return WalkStop, nil
			}
			elements++
			lines = append(lines, formatElementLine(event.Element, formatDepth(event.Path), opts, align))
		case WalkSequenceBegin:
			sequence, ok := event.Element.(*Sequence)
			if !ok || sequence == nil {
				return WalkContinue, nil
			}
			if formatLimitReached(elements, opts.MaxElements) {
				truncated = true
				return WalkStop, nil
			}
			elements++
			depth := formatDepth(event.Path)
			lines = append(lines, formatSequenceLine(sequence, depth, align))
			if opts.MaxDepth >= 0 && depth >= opts.MaxDepth {
				return WalkSkipChildren, nil
			}
		case WalkSequenceItemBegin:
			index := formatItemIndex(event.Path)
			lines = append(lines, strings.Repeat("  ", formatDepth(event.Path))+"Item #"+strconv.Itoa(index+1)+":")
		}
		return WalkContinue, nil
	})
	if truncated || err != nil {
		lines = append(lines, "...")
	}
	return lines
}

func formatElementLine(elem element.Element, depth int, opts FormatOptions, align bool) string {
	tagText := formatTagText(elem)
	vrText := formatVRText(elem)
	indent := strings.Repeat("  ", depth)
	if align && depth > 0 {
		indent += "> "
	}

	line := indent + tagText + " " + vrText
	value := ""
	if !opts.HideValues {
		if value = element.FormatValue(elem, opts.MaxValueLength); value != "" {
			line += " " + value
		}
	}
	if !align {
		return line
	}

	width := utf8.RuneCountInString(indent+tagText+" "+vrText) + 1
	if value != "" {
		// Reserve the value column plus the brackets and ellipsis added by
		// element.FormatValue for truncated text values.
		width += 1 + opts.MaxValueLength + 5
	}
	line = padRight(line, width)
	line += fmt.Sprintf("#%6d", elementLength(elem))
	if keyword := formatKeyword(elem); keyword != "" {
		line += ", " + keyword
	}
	return line
}

func formatSequenceLine(sequence *Sequence, depth int, align bool) string {
	indent := strings.Repeat("  ", depth)
	if align && depth > 0 {
		indent += "> "
	}
	count := sequence.Count()
	label := "items"
	if count == 1 {
		label = "item"
	}
	line := fmt.Sprintf("%s%s %s (%d %s)", indent, formatTagText(sequence), vr.CodeSQ, count, label)
	if align {
		if keyword := formatKeyword(sequence); keyword != "" {
			line += " " + keyword
		}
	}
	return line
}

func formatTagText(elem element.Element) string {
	if elem == nil || elem.Tag() == nil {
		return "<nil-tag>"
	}
	return elem.Tag().String()
}

func formatVRText(elem element.Element) string {
	if elem == nil || elem.ValueRepresentation() == nil {
		return "<nil-vr>"
	}
	return elem.ValueRepresentation().Code()
}

func formatKeyword(elem element.Element) string {
	if elem == nil || elem.Tag() == nil {
		return ""
	}
	entry := dict.Default().Lookup(elem.Tag())
	if entry == nil {
		return ""
	}
	return entry.Keyword()
}

func elementLength(elem element.Element) int64 {
	if elem == nil {
		return 0
	}
	return int64(elem.Length())
}

func formatDepth(path Path) int {
	depth := 0
	for _, segment := range path {
		if segment.ItemIndex != nil {
			depth++
		}
	}
	return depth
}

func formatItemIndex(path Path) int {
	for index := len(path) - 1; index >= 0; index-- {
		if path[index].ItemIndex != nil {
			return *path[index].ItemIndex
		}
	}
	return 0
}

func formatLimitReached(count, limit int) bool {
	return limit > 0 && count >= limit
}

func padRight(value string, width int) string {
	if runes := utf8.RuneCountInString(value); runes < width {
		return value + strings.Repeat(" ", width-runes)
	}
	return value + " "
}
