// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func BenchmarkWalkFlat(b *testing.B) {
	ds := New()
	ds.SetAutoValidate(false)
	for index := 0; index < 1024; index++ {
		t := tag.New(0x0011, uint16(index)) // #nosec G115 -- benchmark bound is 1024
		_ = ds.Add(element.NewString(t, vr.LO, []string{"value"}))
	}
	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		_ = Walk(ds, continueWalk)
	}
}

func BenchmarkWalkDeepSequence(b *testing.B) {
	ds := New()
	for depth := 0; depth < 512; depth++ {
		parent := New()
		parent.SetAutoValidate(false)
		_ = parent.Add(NewSequenceWithItems(tag.ReferencedImageSequence, []*Dataset{ds}))
		ds = parent
	}
	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		_ = Walk(ds, continueWalk)
	}
}

func BenchmarkWalkFragments(b *testing.B) {
	ds := New()
	fragments := element.NewOtherByteFragment(tag.PixelData)
	payload := buffer.NewMemory(make([]byte, 4096))
	for index := 0; index < 1024; index++ {
		fragments.AddFragment(payload)
	}
	_ = ds.Add(fragments)
	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		_ = Walk(ds, continueWalk)
	}
}

func continueWalk(WalkEvent) (WalkAction, error) { return WalkContinue, nil }
