// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"context"
	"fmt"
)

// testPixelData is a simple in-memory implementation of the codec frame ports.
type testPixelData struct {
	frames       [][]byte
	info         FrameInfo
	encapsulated bool
}

// newTestPixelData creates a new test pixel data instance.
func newTestPixelData(info FrameInfo) *testPixelData {
	return &testPixelData{
		frames: make([][]byte, 0),
		info:   info,
	}
}

// Frame returns the pixel data for the specified frame.
func (pd *testPixelData) Frame(ctx context.Context, frameIndex int) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if frameIndex < 0 || frameIndex >= len(pd.frames) {
		return nil, fmt.Errorf("frame index %d out of range [0, %d)", frameIndex, len(pd.frames))
	}
	return append([]byte(nil), pd.frames[frameIndex]...), nil
}

// AddFrame appends a new frame to the pixel data.
func (pd *testPixelData) AddFrame(ctx context.Context, frameData []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	pd.frames = append(pd.frames, append([]byte(nil), frameData...))
	return nil
}

// FrameCount returns the number of frames.
func (pd *testPixelData) FrameCount() int {
	return len(pd.frames)
}

// FrameInfo returns frame metadata.
func (pd *testPixelData) FrameInfo() FrameInfo {
	return pd.info
}

// SetFrameInfo updates frame metadata.
func (pd *testPixelData) SetFrameInfo(info FrameInfo) error {
	pd.info = info
	return nil
}

// Encapsulated returns whether the pixel data is encapsulated.
func (pd *testPixelData) Encapsulated() bool {
	return pd.encapsulated
}
