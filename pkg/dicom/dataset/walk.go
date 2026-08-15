// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"
	"reflect"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// WalkEventKind identifies a Dataset traversal event.
type WalkEventKind uint8

const (
	// WalkInvalid identifies errors that occur before a traversal event.
	WalkInvalid WalkEventKind = iota
	// WalkElement visits an ordinary non-container element.
	WalkElement
	// WalkSequenceBegin starts a Sequence container.
	WalkSequenceBegin
	// WalkSequenceItemBegin starts a Sequence item Dataset.
	WalkSequenceItemBegin
	// WalkSequenceItemEnd ends a Sequence item Dataset.
	WalkSequenceItemEnd
	// WalkSequenceEnd ends a Sequence container.
	WalkSequenceEnd
	// WalkFragmentBegin starts an encapsulated fragment sequence.
	WalkFragmentBegin
	// WalkFragmentItem visits one encoded fragment buffer.
	WalkFragmentItem
	// WalkFragmentEnd ends an encapsulated fragment sequence.
	WalkFragmentEnd
)

// String returns a stable event name.
func (kind WalkEventKind) String() string {
	switch kind {
	case WalkElement:
		return "element"
	case WalkSequenceBegin:
		return "sequence-begin"
	case WalkSequenceItemBegin:
		return "sequence-item-begin"
	case WalkSequenceItemEnd:
		return "sequence-item-end"
	case WalkSequenceEnd:
		return "sequence-end"
	case WalkFragmentBegin:
		return "fragment-begin"
	case WalkFragmentItem:
		return "fragment-item"
	case WalkFragmentEnd:
		return "fragment-end"
	default:
		return "invalid"
	}
}

// WalkEvent describes one location visited by Walk.
type WalkEvent struct {
	Kind     WalkEventKind
	Path     Path
	Element  element.Element
	Dataset  *Dataset
	Fragment buffer.ByteBuffer
}

// WalkAction controls traversal after a visitor callback.
type WalkAction uint8

const (
	// WalkContinue continues traversal normally.
	WalkContinue WalkAction = iota
	// WalkSkipChildren skips a begin event's children but still emits its end event.
	WalkSkipChildren
	// WalkStop ends traversal successfully without emitting later events.
	WalkStop
)

// VisitFunc handles one synchronous Dataset traversal event.
type VisitFunc func(WalkEvent) (WalkAction, error)

// WalkError adds event and path context to a traversal failure.
type WalkError struct {
	Kind  WalkEventKind
	Path  Path
	Cause error
}

func (e *WalkError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Cause == nil {
		return fmt.Sprintf("DICOM walk failed during %s at %s", e.Kind, FormatPath(e.Path))
	}
	return fmt.Sprintf("DICOM walk failed during %s at %s: %v", e.Kind, FormatPath(e.Path), e.Cause)
}

// Unwrap returns the cancellation, visitor, cycle, input, or action error.
func (e *WalkError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

type walkTaskKind uint8

const (
	walkTaskEvent walkTaskKind = iota
	walkTaskDataset
	walkTaskSequenceItems
	walkTaskFragments
	walkTaskLeaveDataset
)

type walkTask struct {
	kind walkTaskKind

	event   WalkEvent
	canSkip bool

	dataset     *Dataset
	path        Path
	elements    []element.Element
	tags        []*tag.Tag
	index       int
	initialized bool

	sequence *Sequence
	items    []*Dataset

	fragmentSequence *element.FragmentSequence
	fragments        []buffer.ByteBuffer
}

// Walk traverses ds depth-first in ascending Dataset tag order.
func Walk(ds *Dataset, visit VisitFunc) error {
	if ds == nil {
		return newWalkError(WalkInvalid, nil, fmt.Errorf("walk Dataset is nil"))
	}
	if visit == nil {
		return newWalkError(WalkInvalid, nil, fmt.Errorf("walk visitor is nil"))
	}

	stack := []walkTask{newDatasetTask(ds, nil)}
	activeDatasets := map[*Dataset]struct{}{ds: {}}
	for len(stack) > 0 {
		last := len(stack) - 1
		task := stack[last]
		stack = stack[:last]

		switch task.kind {
		case walkTaskEvent:
			stop, err := dispatchWalkEvent(task, &stack, visit)
			if err != nil {
				return err
			}
			if stop {
				return nil
			}

		case walkTaskDataset:
			if !task.initialized {
				tagValues := append([]uint32(nil), task.dataset.sortedTagValues()...)
				task.elements = make([]element.Element, len(tagValues))
				task.tags = make([]*tag.Tag, len(tagValues))
				for index, tagValue := range tagValues {
					task.elements[index] = task.dataset.items[tagValue]
					task.tags[index] = tag.New(uint16(tagValue>>16), uint16(tagValue))
					if elem := task.elements[index]; !isNilElement(elem) && elem.Tag() != nil {
						task.tags[index] = elem.Tag().Clone()
					}
				}
				task.initialized = true
			}
			if task.index >= len(task.elements) {
				continue
			}
			elem := task.elements[task.index]
			storedTag := task.tags[task.index]
			task.index++
			stack = append(stack, task)
			if sequence, ok := elem.(*Sequence); ok {
				containerPath := appendPath(task.path, PathSegment{Tag: storedTag})
				if sequence == nil {
					return newWalkError(WalkSequenceBegin, containerPath, fmt.Errorf("sequence is nil"))
				}
				if sequence.Tag() == nil {
					return newWalkError(WalkSequenceBegin, containerPath, fmt.Errorf("sequence tag is nil"))
				}
				items := append([]*Dataset(nil), sequence.GetItems()...)
				stack = append(stack,
					walkTask{kind: walkTaskEvent, event: WalkEvent{Kind: WalkSequenceEnd, Path: containerPath, Element: sequence, Dataset: task.dataset}},
					walkTask{kind: walkTaskSequenceItems, path: containerPath, sequence: sequence, items: items, initialized: true},
					walkTask{kind: walkTaskEvent, event: WalkEvent{Kind: WalkSequenceBegin, Path: containerPath, Element: sequence, Dataset: task.dataset}, canSkip: true},
				)
				continue
			}
			if fragments, ok, err := fragmentSequenceOf(elem); ok {
				containerPath := appendPath(task.path, PathSegment{Tag: storedTag})
				if err != nil {
					return newWalkError(WalkFragmentBegin, containerPath, err)
				}
				if fragments.Tag() == nil {
					return newWalkError(WalkFragmentBegin, containerPath, fmt.Errorf("fragment sequence tag is nil"))
				}
				fragmentItems := append([]buffer.ByteBuffer(nil), fragments.Fragments()...)
				stack = append(stack,
					walkTask{kind: walkTaskEvent, event: WalkEvent{Kind: WalkFragmentEnd, Path: containerPath, Element: fragments, Dataset: task.dataset}},
					walkTask{kind: walkTaskFragments, path: containerPath, dataset: task.dataset, fragmentSequence: fragments, fragments: fragmentItems, initialized: true},
					walkTask{kind: walkTaskEvent, event: WalkEvent{Kind: WalkFragmentBegin, Path: containerPath, Element: fragments, Dataset: task.dataset}, canSkip: true},
				)
				continue
			}
			path := appendPath(task.path, PathSegment{Tag: storedTag})
			if isNilElement(elem) {
				return newWalkError(WalkElement, path, fmt.Errorf("element is nil"))
			}
			if elem.Tag() == nil {
				return newWalkError(WalkElement, path, fmt.Errorf("element tag is nil"))
			}
			stack = append(stack, walkTask{kind: walkTaskEvent, event: WalkEvent{
				Kind: WalkElement, Path: path, Element: elem, Dataset: task.dataset,
			}})

		case walkTaskSequenceItems:
			if task.index >= len(task.items) {
				continue
			}
			itemIndex := task.index
			item := task.items[itemIndex]
			if item == nil {
				item = New()
			}
			task.index++
			itemPath := ClonePath(task.path)
			itemPath[len(itemPath)-1].ItemIndex = &itemIndex
			if _, exists := activeDatasets[item]; exists {
				return newWalkError(WalkSequenceItemBegin, itemPath, fmt.Errorf("Dataset cycle detected"))
			}
			activeDatasets[item] = struct{}{}
			stack = append(stack,
				task,
				walkTask{kind: walkTaskLeaveDataset, dataset: item},
				walkTask{kind: walkTaskEvent, event: WalkEvent{Kind: WalkSequenceItemEnd, Path: itemPath, Element: task.sequence, Dataset: item}},
				newDatasetTask(item, itemPath),
				walkTask{kind: walkTaskEvent, event: WalkEvent{Kind: WalkSequenceItemBegin, Path: itemPath, Element: task.sequence, Dataset: item}, canSkip: true},
			)

		case walkTaskFragments:
			if task.index >= len(task.fragments) {
				continue
			}
			fragmentIndex := task.index
			fragment := task.fragments[fragmentIndex]
			task.index++
			fragmentPath := ClonePath(task.path)
			fragmentPath[len(fragmentPath)-1].FragmentIndex = &fragmentIndex
			stack = append(stack,
				task,
				walkTask{kind: walkTaskEvent, event: WalkEvent{
					Kind: WalkFragmentItem, Path: fragmentPath, Element: task.fragmentSequence, Dataset: task.dataset, Fragment: fragment,
				}},
			)

		case walkTaskLeaveDataset:
			delete(activeDatasets, task.dataset)
		}
	}
	return nil
}

func dispatchWalkEvent(task walkTask, stack *[]walkTask, visit VisitFunc) (bool, error) {
	event := task.event
	event.Path = ClonePath(event.Path)
	action, err := visit(event)
	if err != nil {
		return false, newWalkError(event.Kind, event.Path, err)
	}
	switch action {
	case WalkContinue:
		return false, nil
	case WalkSkipChildren:
		if !task.canSkip {
			return false, newWalkError(event.Kind, event.Path, fmt.Errorf("cannot skip children for %s event", event.Kind))
		}
		if len(*stack) == 0 {
			return false, newWalkError(event.Kind, event.Path, fmt.Errorf("missing child traversal state"))
		}
		*stack = (*stack)[:len(*stack)-1]
		return false, nil
	case WalkStop:
		return true, nil
	default:
		return false, newWalkError(event.Kind, event.Path, fmt.Errorf("invalid walk action %d", action))
	}
}

func newWalkError(kind WalkEventKind, path Path, cause error) error {
	return &WalkError{Kind: kind, Path: ClonePath(path), Cause: cause}
}

func newDatasetTask(ds *Dataset, path Path) walkTask {
	return walkTask{
		kind:    walkTaskDataset,
		dataset: ds,
		path:    ClonePath(path),
	}
}

func fragmentSequenceOf(elem element.Element) (*element.FragmentSequence, bool, error) {
	switch value := elem.(type) {
	case *element.FragmentSequence:
		if value == nil {
			return nil, true, fmt.Errorf("fragment sequence is nil")
		}
		return value, true, nil
	case *element.OtherByteFragment:
		if value == nil || value.FragmentSequence == nil {
			return nil, true, fmt.Errorf("fragment sequence is nil")
		}
		return value.FragmentSequence, true, nil
	case *element.OtherWordFragment:
		if value == nil || value.FragmentSequence == nil {
			return nil, true, fmt.Errorf("fragment sequence is nil")
		}
		return value.FragmentSequence, true, nil
	default:
		return nil, false, nil
	}
}

func isNilElement(elem element.Element) bool {
	if elem == nil {
		return true
	}
	value := reflect.ValueOf(elem)
	return value.Kind() == reflect.Pointer && value.IsNil()
}
