// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package buffer internal tests for FileByteBuffer position overflow behaviour.
package buffer

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestFileByteBufferPositionOverflowPanics(t *testing.T) {
	buf := &FileByteBuffer{position: math.MaxUint32 + 1}
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Position() should panic on overflow")
		}
		msg := fmt.Sprintf("%v", r)
		if !strings.Contains(msg, "overflows uint32") {
			t.Errorf("panic message %q does not contain \"overflows uint32\"", msg)
		}
	}()
	buf.Position()
}
