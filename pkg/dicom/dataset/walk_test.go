// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"errors"
	"reflect"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

type observedWalkEvent struct {
	kind WalkEventKind
	path string
}

func TestWalkVisitsDepthFirstInTagOrder(t *testing.T) {
	root := New()
	item := New()
	requireWalkAdd(t, item, element.NewString(tag.SOPInstanceUID, vr.UI, []string{testStudyInstanceUID}))
	requireWalkAdd(t, root, element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	requireWalkAdd(t, root, NewSequenceWithItems(tag.ReferencedImageSequence, []*Dataset{item}))

	var got []observedWalkEvent
	err := Walk(root, func(event WalkEvent) (WalkAction, error) {
		got = append(got, observedWalkEvent{kind: event.Kind, path: FormatPath(event.Path)})
		return WalkContinue, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	want := []observedWalkEvent{
		{kind: WalkSequenceBegin, path: referencedImageSequencePath},
		{kind: WalkSequenceItemBegin, path: testSourceImagePath},
		{kind: WalkElement, path: "(0008,1140)[0]/(0008,0018)"},
		{kind: WalkSequenceItemEnd, path: testSourceImagePath},
		{kind: WalkSequenceEnd, path: referencedImageSequencePath},
		{kind: WalkElement, path: "(0010,0010)"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Walk events = %#v, want %#v", got, want)
	}
}

func TestWalkPathPreservesPrivateCreatorWithoutAliasingSourceTag(t *testing.T) {
	privateTag := tag.NewWithPrivateCreator(0x0011, 0x1010, tag.NewPrivateCreator(privateCreatorOriginal))
	root := New()
	requireWalkAdd(t, root, element.NewString(privateTag, vr.LO, []string{"value"}))

	err := Walk(root, func(event WalkEvent) (WalkAction, error) {
		if event.Kind != WalkElement {
			return WalkContinue, nil
		}
		if creator := event.Path[0].Tag.PrivateCreator(); creator == nil || creator.Creator() != privateCreatorOriginal {
			t.Fatalf("walk path Tag = %v, want private creator ORIGINAL", event.Path[0].Tag)
		}
		event.Path[0].Tag.SetPrivateCreator(tag.NewPrivateCreator("CHANGED"))
		return WalkContinue, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if creator := root.GetOrNil(privateTag).Tag().PrivateCreator(); creator == nil || creator.Creator() != privateCreatorOriginal {
		t.Fatalf("walk path mutation changed source Tag: %v", root.GetOrNil(privateTag).Tag())
	}
}

func TestWalkEmptyContainersAreBalancedAndNilItemsAreEmptyDatasets(t *testing.T) {
	root := New()
	emptyItem := New()
	sequence := NewSequenceWithItems(tag.ReferencedImageSequence, []*Dataset{nil, emptyItem})
	requireWalkAdd(t, root, sequence)
	fragments := element.NewOtherByteFragment(tag.PixelData)
	requireWalkAdd(t, root, fragments)

	var got []observedWalkEvent
	err := Walk(root, func(event WalkEvent) (WalkAction, error) {
		got = append(got, observedWalkEvent{kind: event.Kind, path: FormatPath(event.Path)})
		if (event.Kind == WalkSequenceItemBegin || event.Kind == WalkSequenceItemEnd) && event.Dataset == nil {
			t.Fatalf("nil item event Dataset at %s", FormatPath(event.Path))
		}
		return WalkContinue, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	want := []observedWalkEvent{
		{kind: WalkSequenceBegin, path: referencedImageSequencePath},
		{kind: WalkSequenceItemBegin, path: testSourceImagePath},
		{kind: WalkSequenceItemEnd, path: testSourceImagePath},
		{kind: WalkSequenceItemBegin, path: "(0008,1140)[1]"},
		{kind: WalkSequenceItemEnd, path: "(0008,1140)[1]"},
		{kind: WalkSequenceEnd, path: referencedImageSequencePath},
		{kind: WalkFragmentBegin, path: "(7fe0,0010)"},
		{kind: WalkFragmentEnd, path: "(7fe0,0010)"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Walk events = %#v, want %#v", got, want)
	}
}

func TestWalkFragmentBeginExposesOffsetTableAndItemsExposeBuffers(t *testing.T) {
	root := New()
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.SetOffsetTable([]uint32{0, 4})
	first := buffer.NewMemory([]byte{1, 2, 3, 4})
	second := buffer.NewMemory([]byte{5, 6})
	fragments.AddFragment(first)
	fragments.AddFragment(second)
	requireWalkAdd(t, root, fragments)

	var begin *element.FragmentSequence
	var gotPaths []string
	var gotFragments []buffer.ByteBuffer
	err := Walk(root, func(event WalkEvent) (WalkAction, error) {
		switch event.Kind {
		case WalkFragmentBegin:
			var ok bool
			begin, ok = event.Element.(*element.FragmentSequence)
			if !ok {
				t.Fatalf("fragment begin Element type = %T, want *element.FragmentSequence", event.Element)
			}
		case WalkFragmentItem:
			gotPaths = append(gotPaths, FormatPath(event.Path))
			gotFragments = append(gotFragments, event.Fragment)
		}
		return WalkContinue, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(begin.OffsetTable(), []uint32{0, 4}) {
		t.Fatalf("offset table = %v, want [0 4]", begin.OffsetTable())
	}
	if !reflect.DeepEqual(gotPaths, []string{"(7fe0,0010)#0", "(7fe0,0010)#1"}) {
		t.Fatalf("fragment paths = %v", gotPaths)
	}
	if len(gotFragments) != 2 || gotFragments[0] != first || gotFragments[1] != second {
		t.Fatal("fragment item buffers were copied or reordered")
	}
}

func TestWalkSnapshotsSequenceItemsBeforeBeginCallback(t *testing.T) {
	first := New()
	requireWalkAdd(t, first, element.NewString(tag.PatientID, vr.LO, []string{"first"}))
	second := New()
	requireWalkAdd(t, second, element.NewString(tag.PatientID, vr.LO, []string{"second"}))
	sequence := NewSequenceWithItems(tag.ReferencedImageSequence, []*Dataset{first})
	root := New()
	requireWalkAdd(t, root, sequence)

	var values []string
	err := Walk(root, func(event WalkEvent) (WalkAction, error) {
		if event.Kind == WalkSequenceBegin {
			sequence.AddItem(second)
		}
		if event.Kind == WalkElement && event.Element.Tag().Equals(tag.PatientID) {
			value, _ := event.Dataset.GetString(tag.PatientID)
			values = append(values, value)
		}
		return WalkContinue, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(values, []string{"first"}) {
		t.Fatalf("visited Sequence item values = %v, want pre-callback snapshot", values)
	}
}

func TestWalkSnapshotsFragmentsBeforeBeginCallback(t *testing.T) {
	first := buffer.NewMemory([]byte{1, 2})
	second := buffer.NewMemory([]byte{3, 4})
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(first)
	root := New()
	requireWalkAdd(t, root, fragments)

	var visited []buffer.ByteBuffer
	err := Walk(root, func(event WalkEvent) (WalkAction, error) {
		if event.Kind == WalkFragmentBegin {
			fragments.AddFragment(second)
		}
		if event.Kind == WalkFragmentItem {
			visited = append(visited, event.Fragment)
		}
		return WalkContinue, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(visited, []buffer.ByteBuffer{first}) {
		t.Fatalf("visited fragments = %v, want pre-callback snapshot", visited)
	}
}

func TestWalkEventFieldsFollowEventKind(t *testing.T) {
	root := New()
	item := New()
	sequence := NewSequenceWithItems(tag.ReferencedImageSequence, []*Dataset{item})
	requireWalkAdd(t, root, sequence)

	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{1, 2}))
	requireWalkAdd(t, root, fragments)

	err := Walk(root, func(event WalkEvent) (WalkAction, error) {
		switch event.Kind {
		case WalkSequenceBegin, WalkSequenceEnd:
			if event.Element != sequence || event.Dataset != root || event.Fragment != nil {
				t.Fatalf("sequence event fields = %#v", event)
			}
		case WalkSequenceItemBegin, WalkSequenceItemEnd:
			if event.Element != sequence || event.Dataset != item || event.Fragment != nil {
				t.Fatalf("sequence item event fields = %#v", event)
			}
		case WalkFragmentBegin, WalkFragmentItem, WalkFragmentEnd:
			if event.Element != fragments.FragmentSequence || event.Dataset != root {
				t.Fatalf("fragment event fields = %#v", event)
			}
		case WalkElement:
			if event.Dataset == nil {
				t.Fatalf("element event Dataset is nil: %#v", event)
			}
		}
		return WalkContinue, nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestWalkStopReturnsWithoutLaterEvents(t *testing.T) {
	root := newWalkControlFixture(t)
	allKinds := collectWalkKinds(t, root, WalkInvalid, WalkContinue)

	for _, stopKind := range allKinds {
		stopKind := stopKind
		t.Run(stopKind.String(), func(t *testing.T) {
			var got []WalkEventKind
			err := Walk(root, func(event WalkEvent) (WalkAction, error) {
				got = append(got, event.Kind)
				if event.Kind == stopKind {
					return WalkStop, nil
				}
				return WalkContinue, nil
			})
			if err != nil {
				t.Fatal(err)
			}
			if len(got) == 0 || got[len(got)-1] != stopKind {
				t.Fatalf("last event = %v, want %v", got, stopKind)
			}
		})
	}
}

func TestWalkSkipChildrenEmitsMatchingEnd(t *testing.T) {
	tests := []struct {
		name string
		skip WalkEventKind
		want []WalkEventKind
	}{
		{
			name: "sequence",
			skip: WalkSequenceBegin,
			want: []WalkEventKind{WalkSequenceBegin, WalkSequenceEnd, WalkFragmentBegin, WalkFragmentItem, WalkFragmentEnd},
		},
		{
			name: "sequence item",
			skip: WalkSequenceItemBegin,
			want: []WalkEventKind{WalkSequenceBegin, WalkSequenceItemBegin, WalkSequenceItemEnd, WalkSequenceEnd, WalkFragmentBegin, WalkFragmentItem, WalkFragmentEnd},
		},
		{
			name: "fragment",
			skip: WalkFragmentBegin,
			want: []WalkEventKind{WalkSequenceBegin, WalkSequenceItemBegin, WalkElement, WalkSequenceItemEnd, WalkSequenceEnd, WalkFragmentBegin, WalkFragmentEnd},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := collectWalkKinds(t, newWalkControlFixture(t), tc.skip, WalkSkipChildren)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("events = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestWalkRejectsSkipChildrenOnLeafAndEndEvents(t *testing.T) {
	invalidKinds := []WalkEventKind{
		WalkElement,
		WalkSequenceItemEnd,
		WalkSequenceEnd,
		WalkFragmentItem,
		WalkFragmentEnd,
	}
	for _, invalidKind := range invalidKinds {
		invalidKind := invalidKind
		t.Run(invalidKind.String(), func(t *testing.T) {
			err := Walk(newWalkControlFixture(t), func(event WalkEvent) (WalkAction, error) {
				if event.Kind == invalidKind {
					return WalkSkipChildren, nil
				}
				return WalkContinue, nil
			})
			var walkErr *WalkError
			if !errors.As(err, &walkErr) {
				t.Fatalf("error = %v, want *WalkError", err)
			}
			if walkErr.Kind != invalidKind || len(walkErr.Path) == 0 || walkErr.Cause == nil {
				t.Fatalf("WalkError = %#v", walkErr)
			}
		})
	}
}

func TestWalkWrapsVisitorErrorWithEventContext(t *testing.T) {
	sentinel := errors.New("visitor failed")
	err := Walk(newWalkControlFixture(t), func(event WalkEvent) (WalkAction, error) {
		if event.Kind == WalkSequenceItemBegin {
			return WalkContinue, sentinel
		}
		return WalkContinue, nil
	})

	var walkErr *WalkError
	if !errors.As(err, &walkErr) || !errors.Is(err, sentinel) {
		t.Fatalf("error = %v, want wrapping *WalkError", err)
	}
	if walkErr.Kind != WalkSequenceItemBegin || FormatPath(walkErr.Path) != testSourceImagePath {
		t.Fatalf("WalkError = %#v", walkErr)
	}
}

func TestWalkRejectsNilInputsWithWalkError(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{name: "dataset", err: Walk(nil, func(WalkEvent) (WalkAction, error) { return WalkContinue, nil })},
		{name: "visitor", err: Walk(New(), nil)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var walkErr *WalkError
			if !errors.As(tc.err, &walkErr) || walkErr.Kind != WalkInvalid {
				t.Fatalf("error = %v, want invalid *WalkError", tc.err)
			}
		})
	}
}

func TestWalkTypedNilSequenceReturnsContextualError(t *testing.T) {
	ds := New()
	var sequence *Sequence
	ds.items[tag.ReferencedImageSequence.ToUint32()] = sequence

	err := Walk(ds, func(WalkEvent) (WalkAction, error) {
		t.Fatal("visitor called for typed-nil Sequence")
		return WalkContinue, nil
	})
	var walkErr *WalkError
	if !errors.As(err, &walkErr) {
		t.Fatalf("error = %v, want *WalkError", err)
	}
	if walkErr.Kind != WalkSequenceBegin || FormatPath(walkErr.Path) != referencedImageSequencePath {
		t.Fatalf("WalkError = %#v", walkErr)
	}
}

func TestWalkDetectsDirectAndIndirectDatasetCycles(t *testing.T) {
	direct := New()
	directSequence := NewSequence(tag.ReferencedImageSequence)
	requireWalkAdd(t, direct, directSequence)
	directSequence.AddItem(direct)

	child := New()
	root := New()
	rootSequence := NewSequenceWithItems(tag.ReferencedImageSequence, []*Dataset{child})
	requireWalkAdd(t, root, rootSequence)
	child.SetAutoValidate(false)
	reverseSequence := NewSequenceWithItems(tag.SourceImageSequence, []*Dataset{root})
	requireWalkAdd(t, child, reverseSequence)

	for _, tc := range []struct {
		name string
		ds   *Dataset
		path string
	}{
		{name: "direct", ds: direct, path: testSourceImagePath},
		{name: "indirect", ds: root, path: "(0008,1140)[0]/(0008,2112)[0]"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := Walk(tc.ds, func(WalkEvent) (WalkAction, error) { return WalkContinue, nil })
			var walkErr *WalkError
			if !errors.As(err, &walkErr) {
				t.Fatalf("error = %v, want *WalkError", err)
			}
			if walkErr.Kind != WalkSequenceItemBegin || FormatPath(walkErr.Path) != tc.path {
				t.Fatalf("WalkError = %#v, want path %s", walkErr, tc.path)
			}
		})
	}
}

func TestWalkAllowsSharedSiblingDatasetAndRetainedPaths(t *testing.T) {
	child := New()
	requireWalkAdd(t, child, element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	root := New()
	requireWalkAdd(t, root, NewSequenceWithItems(tag.ReferencedImageSequence, []*Dataset{child, child}))

	var retained []Path
	err := Walk(root, func(event WalkEvent) (WalkAction, error) {
		retained = append(retained, event.Path)
		return WalkContinue, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := FormatPath(retained[0]); got != referencedImageSequencePath {
		t.Fatalf("retained first path changed to %s", got)
	}
	wantLast := "(0008,1140)[1]/(0010,0010)"
	if got := FormatPath(retained[len(retained)-3]); got != wantLast {
		t.Fatalf("second sibling element path = %s, want %s", got, wantLast)
	}
}

func newWalkControlFixture(t *testing.T) *Dataset {
	t.Helper()
	item := New()
	requireWalkAdd(t, item, element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	root := New()
	requireWalkAdd(t, root, NewSequenceWithItems(tag.ReferencedImageSequence, []*Dataset{item}))
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{1, 2}))
	requireWalkAdd(t, root, fragments)
	return root
}

func collectWalkKinds(t *testing.T, ds *Dataset, target WalkEventKind, action WalkAction) []WalkEventKind {
	t.Helper()
	var got []WalkEventKind
	err := Walk(ds, func(event WalkEvent) (WalkAction, error) {
		got = append(got, event.Kind)
		if event.Kind == target {
			return action, nil
		}
		return WalkContinue, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func requireWalkAdd(t *testing.T, ds *Dataset, elem element.Element) {
	t.Helper()
	if err := ds.Add(elem); err != nil {
		t.Fatal(err)
	}
}
