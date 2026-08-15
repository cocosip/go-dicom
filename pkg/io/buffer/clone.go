// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"bytes"
	"fmt"
	"reflect"
)

// Clone returns an independent buffer while preserving wrapper and bulk-data
// semantics. Any I/O required to detach external storage is reported.
func Clone(source ByteBuffer) (ByteBuffer, error) {
	if source == nil || isNilByteBuffer(source) {
		return nil, nil
	}
	if source == Empty {
		return Empty, nil
	}

	switch value := source.(type) {
	case *MemoryByteBuffer:
		return NewMemory(append([]byte(nil), value.Data()...)), nil
	case *PooledBuffer:
		return NewMemory(append([]byte(nil), value.Data()...)), nil
	case *BulkDataURIByteBuffer:
		if value.data == nil {
			return NewBulkDataURI(value.bulkDataURI), nil
		}
		return NewBulkDataURIWithData(value.bulkDataURI, append([]byte(nil), value.data...)), nil
	case *RangeByteBuffer:
		internal, err := Clone(value.internal)
		if err != nil {
			return nil, fmt.Errorf("clone range buffer: %w", err)
		}
		cloned, err := NewRange(internal, value.offset, value.length)
		if err != nil {
			return nil, fmt.Errorf("clone range buffer: %w", err)
		}
		return cloned, nil
	case *CompositeByteBuffer:
		buffers := make([]ByteBuffer, len(value.buffers))
		for index, child := range value.buffers {
			cloned, err := Clone(child)
			if err != nil {
				return nil, fmt.Errorf("clone composite buffer %d: %w", index, err)
			}
			buffers[index] = cloned
		}
		return NewComposite(buffers...), nil
	case *EvenLengthBuffer:
		internal, err := Clone(value.internal)
		if err != nil {
			return nil, fmt.Errorf("clone even-length buffer: %w", err)
		}
		return NewEvenLength(internal), nil
	case *EndianByteBuffer:
		internal, err := Clone(value.internal)
		if err != nil {
			return nil, fmt.Errorf("clone endian buffer: %w", err)
		}
		return NewEndian(internal, value.endian, value.unitSize), nil
	case *SwapByteBuffer:
		internal, err := Clone(value.internal)
		if err != nil {
			return nil, fmt.Errorf("clone swap buffer: %w", err)
		}
		cloned, err := NewSwap(internal, value.unitSize)
		if err != nil {
			return nil, fmt.Errorf("clone swap buffer: %w", err)
		}
		return cloned, nil
	default:
		return materializeClone(source)
	}
}

func materializeClone(source ByteBuffer) (ByteBuffer, error) {
	var destination bytes.Buffer
	written, err := source.WriteTo(&destination)
	if err != nil {
		return nil, fmt.Errorf("read source buffer: %w", err)
	}
	if written != int64(source.Size()) || written != int64(destination.Len()) {
		return nil, fmt.Errorf("read source buffer: wrote %d bytes, expected %d", written, source.Size())
	}
	return NewMemory(append([]byte(nil), destination.Bytes()...)), nil
}

func isNilByteBuffer(source ByteBuffer) bool {
	value := reflect.ValueOf(source)
	return value.Kind() == reflect.Ptr && value.IsNil()
}
