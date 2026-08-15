// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

type walkLengths struct {
	sequences map[string]uint32
	items     map[string]uint32
}

type lengthFrame struct {
	kind      dataset.WalkEventKind
	path      string
	content   uint64
	fragments *element.FragmentSequence
}

type walkLengthCalculator struct {
	writer  *Writer
	root    *dataset.Dataset
	lengths *walkLengths
	frames  []lengthFrame
	skipped map[string]dataset.WalkEventKind
}

func (w *Writer) calculateWalkLengths(ds *dataset.Dataset) (*walkLengths, error) {
	lengths := &walkLengths{
		sequences: make(map[string]uint32),
		items:     make(map[string]uint32),
	}
	calculator := &walkLengthCalculator{
		writer:  w,
		root:    ds,
		lengths: lengths,
		frames:  []lengthFrame{{kind: dataset.WalkInvalid}},
		skipped: make(map[string]dataset.WalkEventKind),
	}
	if err := dataset.Walk(ds, calculator.visit); err != nil {
		return nil, fmt.Errorf("failed to calculate encoded lengths: %w", err)
	}
	if len(calculator.frames) != 1 {
		return nil, fmt.Errorf("unbalanced Dataset walk while calculating encoded lengths")
	}
	return lengths, nil
}

func (c *walkLengthCalculator) visit(event dataset.WalkEvent) (dataset.WalkAction, error) {
	path := dataset.FormatPath(event.Path)
	switch event.Kind {
	case dataset.WalkElement:
		if c.writer.skipRootElement(c.root, event) {
			return dataset.WalkContinue, nil
		}
		length, err := c.writer.encodedElementLength(event.Element)
		if err != nil {
			return dataset.WalkContinue, err
		}
		return dataset.WalkContinue, c.add(length)

	case dataset.WalkSequenceBegin:
		if c.writer.skipRootElement(c.root, event) {
			c.skipped[path] = event.Kind
			return dataset.WalkSkipChildren, nil
		}
		c.frames = append(c.frames, lengthFrame{kind: event.Kind, path: path})

	case dataset.WalkSequenceItemBegin:
		c.frames = append(c.frames, lengthFrame{kind: event.Kind, path: path})

	case dataset.WalkSequenceItemEnd:
		frame, err := c.pop(dataset.WalkSequenceItemBegin, path)
		if err != nil {
			return dataset.WalkContinue, err
		}
		content, err := uint32Length(frame.content, "sequence item", path)
		if err != nil {
			return dataset.WalkContinue, err
		}
		c.lengths.items[path] = content
		total := uint64(8) + frame.content
		if !c.writer.explicitLengthSequenceItems {
			total += 8
		}
		return dataset.WalkContinue, c.add(total)

	case dataset.WalkSequenceEnd:
		if c.consumeSkipped(path, dataset.WalkSequenceBegin) {
			return dataset.WalkContinue, nil
		}
		frame, err := c.pop(dataset.WalkSequenceBegin, path)
		if err != nil {
			return dataset.WalkContinue, err
		}
		content, err := uint32Length(frame.content, "sequence", path)
		if err != nil {
			return dataset.WalkContinue, err
		}
		c.lengths.sequences[path] = content
		total := uint64(c.writer.elementHeaderLength(vr.SQ)) + frame.content
		if !c.writer.explicitLengthSequences {
			total += 8
		}
		return dataset.WalkContinue, c.add(total)

	case dataset.WalkFragmentBegin:
		if c.writer.skipRootElement(c.root, event) {
			c.skipped[path] = event.Kind
			return dataset.WalkSkipChildren, nil
		}
		fragments, ok := event.Element.(*element.FragmentSequence)
		if !ok || fragments == nil {
			return dataset.WalkContinue, fmt.Errorf("fragment event contains %T", event.Element)
		}
		c.frames = append(c.frames, lengthFrame{kind: event.Kind, path: path, fragments: fragments})

	case dataset.WalkFragmentItem:
		if event.Fragment == nil {
			return dataset.WalkContinue, fmt.Errorf("fragment item buffer is nil")
		}
		length := uint64(event.Fragment.Size())
		if length%2 != 0 {
			length++
		}
		return dataset.WalkContinue, c.add(8 + length)

	case dataset.WalkFragmentEnd:
		if c.consumeSkipped(path, dataset.WalkFragmentBegin) {
			return dataset.WalkContinue, nil
		}
		frame, err := c.pop(dataset.WalkFragmentBegin, path)
		if err != nil {
			return dataset.WalkContinue, err
		}
		offsetCount := effectiveOffsetCount(frame.fragments)
		total := uint64(c.writer.elementHeaderLength(frame.fragments.ValueRepresentation())) +
			8 + uint64(offsetCount)*4 + frame.content + 8
		return dataset.WalkContinue, c.add(total)
	}
	return dataset.WalkContinue, nil
}

func (c *walkLengthCalculator) add(length uint64) error {
	if len(c.frames) == 0 || math.MaxUint64-c.frames[len(c.frames)-1].content < length {
		return fmt.Errorf("encoded length overflow")
	}
	c.frames[len(c.frames)-1].content += length
	return nil
}

func (c *walkLengthCalculator) pop(kind dataset.WalkEventKind, path string) (lengthFrame, error) {
	if len(c.frames) <= 1 {
		return lengthFrame{}, fmt.Errorf("unexpected %s end at %s", kind, path)
	}
	last := len(c.frames) - 1
	frame := c.frames[last]
	c.frames = c.frames[:last]
	if frame.kind != kind || frame.path != path {
		return lengthFrame{}, fmt.Errorf("unbalanced %s at %s", kind, path)
	}
	return frame, nil
}

func (c *walkLengthCalculator) consumeSkipped(path string, kind dataset.WalkEventKind) bool {
	if c.skipped[path] != kind {
		return false
	}
	delete(c.skipped, path)
	return true
}

func uint32Length(length uint64, name, path string) (uint32, error) {
	if length > math.MaxUint32 {
		return 0, fmt.Errorf("%s at %s exceeds maximum encoded length", name, path)
	}
	return uint32(length), nil
}

func (w *Writer) encodedElementLength(elem element.Element) (uint64, error) {
	if elem == nil || elem.ValueRepresentation() == nil {
		return 0, fmt.Errorf("element or value representation is nil")
	}
	valueLength := uint64(0)
	if elem.Buffer() != nil {
		valueLength = uint64(elem.Buffer().Size())
	}
	if valueLength%2 != 0 {
		valueLength++
	}
	return uint64(w.elementHeaderLength(elem.ValueRepresentation())) + valueLength, nil
}

func (w *Writer) elementHeaderLength(valueRepresentation *vr.VR) uint32 {
	if !w.isExplicitVR || valueRepresentation.Is16bitLength() {
		return 8
	}
	return 12
}

func (w *Writer) skipRootElement(root *dataset.Dataset, event dataset.WalkEvent) bool {
	if event.Dataset != root || event.Element == nil || event.Element.Tag() == nil {
		return false
	}
	t := event.Element.Tag()
	return t.Group() == 0x0002 || (!w.keepGroupLengths && t.Element() == 0x0000)
}

type walkDatasetEncoder struct {
	writer  *Writer
	root    *dataset.Dataset
	lengths *walkLengths
	skipped map[string]dataset.WalkEventKind
}

func (w *Writer) writeDatasetWithWalk(ds *dataset.Dataset, lengths *walkLengths) error {
	encoder := &walkDatasetEncoder{
		writer:  w,
		root:    ds,
		lengths: lengths,
		skipped: make(map[string]dataset.WalkEventKind),
	}
	if err := dataset.Walk(ds, encoder.visit); err != nil {
		return fmt.Errorf("failed to encode Dataset walk: %w", err)
	}
	return nil
}

func (e *walkDatasetEncoder) visit(event dataset.WalkEvent) (dataset.WalkAction, error) {
	path := dataset.FormatPath(event.Path)
	switch event.Kind {
	case dataset.WalkElement:
		if e.writer.skipRootElement(e.root, event) {
			return dataset.WalkContinue, nil
		}
		return dataset.WalkContinue, e.writer.writeElement(event.Element)

	case dataset.WalkSequenceBegin:
		if e.writer.skipRootElement(e.root, event) {
			e.skipped[path] = event.Kind
			return dataset.WalkSkipChildren, nil
		}
		sequence := event.Element.(*dataset.Sequence)
		if err := e.writer.writeTag(sequence.Tag()); err != nil {
			return dataset.WalkContinue, fmt.Errorf("failed to write sequence tag %s: %w", sequence.Tag(), err)
		}
		if e.writer.isExplicitVR {
			if err := e.writer.writeVR(vr.SQ); err != nil {
				return dataset.WalkContinue, fmt.Errorf("failed to write SQ VR: %w", err)
			}
		}
		length := uint32(math.MaxUint32)
		if e.writer.explicitLengthSequences {
			length = e.lengths.sequences[path]
		}
		return dataset.WalkContinue, e.writer.writeLength(vr.SQ, length)

	case dataset.WalkSequenceItemBegin:
		if err := e.writer.observeSequenceItem(event.Element.Tag(), event.Dataset); err != nil {
			return dataset.WalkContinue, err
		}
		if err := e.writer.writeTag(tag.New(0xFFFE, 0xE000)); err != nil {
			return dataset.WalkContinue, err
		}
		length := uint32(math.MaxUint32)
		if e.writer.explicitLengthSequenceItems {
			length = e.lengths.items[path]
		}
		return dataset.WalkContinue, binary.Write(e.writer.writer, e.writer.byteOrder, length)

	case dataset.WalkSequenceItemEnd:
		if e.writer.explicitLengthSequenceItems {
			return dataset.WalkContinue, nil
		}
		return dataset.WalkContinue, e.writer.writeDelimiter(0xE00D)

	case dataset.WalkSequenceEnd:
		if e.consumeSkipped(path, dataset.WalkSequenceBegin) || e.writer.explicitLengthSequences {
			return dataset.WalkContinue, nil
		}
		return dataset.WalkContinue, e.writer.writeDelimiter(0xE0DD)

	case dataset.WalkFragmentBegin:
		if e.writer.skipRootElement(e.root, event) {
			e.skipped[path] = event.Kind
			return dataset.WalkSkipChildren, nil
		}
		fragments := event.Element.(*element.FragmentSequence)
		return dataset.WalkContinue, e.writer.writeFragmentBegin(fragments)

	case dataset.WalkFragmentItem:
		return dataset.WalkContinue, e.writer.writeFragmentItem(event.Fragment)

	case dataset.WalkFragmentEnd:
		if e.consumeSkipped(path, dataset.WalkFragmentBegin) {
			return dataset.WalkContinue, nil
		}
		return dataset.WalkContinue, e.writer.writeDelimiter(0xE0DD)
	}
	return dataset.WalkContinue, nil
}

func (e *walkDatasetEncoder) consumeSkipped(path string, kind dataset.WalkEventKind) bool {
	if e.skipped[path] != kind {
		return false
	}
	delete(e.skipped, path)
	return true
}

func (w *Writer) writeDelimiter(elementNumber uint16) error {
	if err := w.writeTag(tag.New(0xFFFE, elementNumber)); err != nil {
		return err
	}
	return binary.Write(w.writer, w.byteOrder, uint32(0))
}

func (w *Writer) writeFragmentBegin(fs *element.FragmentSequence) error {
	if err := w.writeTag(fs.Tag()); err != nil {
		return fmt.Errorf("failed to write fragment sequence tag: %w", err)
	}
	if w.isExplicitVR {
		if err := w.writeVR(fs.ValueRepresentation()); err != nil {
			return fmt.Errorf("failed to write fragment sequence VR: %w", err)
		}
	}
	if err := w.writeLength(fs.ValueRepresentation(), math.MaxUint32); err != nil {
		return fmt.Errorf("failed to write undefined fragment length: %w", err)
	}
	if err := w.writeTag(tag.New(0xFFFE, 0xE000)); err != nil {
		return fmt.Errorf("failed to write offset table item tag: %w", err)
	}

	offsets := fs.OffsetTable()
	if len(offsets) == 0 && fs.FragmentCount() == 1 {
		offsets = []uint32{0}
	}
	if len(offsets) > int(math.MaxUint32/4) {
		return fmt.Errorf("offset table too large: %d entries", len(offsets))
	}
	if err := binary.Write(w.writer, w.byteOrder, uint32(len(offsets))*4); err != nil { // #nosec G115 -- checked above
		return fmt.Errorf("failed to write offset table length: %w", err)
	}
	for _, offset := range offsets {
		if err := binary.Write(w.writer, w.byteOrder, offset); err != nil {
			return fmt.Errorf("failed to write offset value: %w", err)
		}
	}
	return nil
}

func (w *Writer) writeFragmentItem(fragment buffer.ByteBuffer) error {
	if fragment == nil {
		return fmt.Errorf("fragment buffer is nil")
	}
	length := fragment.Size()
	paddedLength := length
	if length%2 != 0 {
		if length == math.MaxUint32 {
			return fmt.Errorf("fragment is too large to pad")
		}
		paddedLength++
	}
	if err := w.writeTag(tag.New(0xFFFE, 0xE000)); err != nil {
		return fmt.Errorf("failed to write fragment item tag: %w", err)
	}
	if err := binary.Write(w.writer, w.byteOrder, paddedLength); err != nil {
		return fmt.Errorf("failed to write fragment length: %w", err)
	}
	written, err := fragment.WriteTo(w.writer)
	if err != nil {
		return fmt.Errorf("failed to write fragment data: %w", err)
	}
	if written != int64(length) {
		return fmt.Errorf("short fragment write: wrote %d bytes, expected %d", written, length)
	}
	if length%2 != 0 {
		return writeAll(w.writer, []byte{0})
	}
	return nil
}

func effectiveOffsetCount(fs *element.FragmentSequence) int {
	if len(fs.OffsetTable()) == 0 && fs.FragmentCount() == 1 {
		return 1
	}
	return len(fs.OffsetTable())
}
