// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

//revive:disable:var-naming // package name must match public import path (pkg/dicom/parser)
package parser

import (
	"bytes"
	"encoding/binary"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestReadLargeOnDemand_ConcurrentLazyLoads(t *testing.T) {
	var dicom bytes.Buffer
	dicom.Write(make([]byte, 128))
	dicom.WriteString("DICM")

	// File meta: Transfer Syntax UID (Explicit VR Little Endian)
	_ = binary.Write(&dicom, binary.LittleEndian, uint16(0x0002))
	_ = binary.Write(&dicom, binary.LittleEndian, uint16(0x0010))
	_, _ = dicom.WriteString("UI")
	tsUID := []byte("1.2.840.10008.1.2.1\x00")
	_ = binary.Write(&dicom, binary.LittleEndian, uint16(len(tsUID)))
	_, _ = dicom.Write(tsUID)

	writeOB := func(group, elem uint16, payload []byte) {
		_ = binary.Write(&dicom, binary.LittleEndian, group)
		_ = binary.Write(&dicom, binary.LittleEndian, elem)
		_, _ = dicom.WriteString("OB")
		_ = binary.Write(&dicom, binary.LittleEndian, uint16(0))
		_ = binary.Write(&dicom, binary.LittleEndian, uint32(len(payload)))
		_, _ = dicom.Write(payload)
	}

	data1 := bytes.Repeat([]byte{0x11}, 128*1024)
	data2 := bytes.Repeat([]byte{0x22}, 128*1024)
	writeOB(0x7FE1, 0x0010, data1)
	writeOB(0x7FE1, 0x0011, data2)

	result, err := Parse(bytes.NewReader(dicom.Bytes()),
		WithReadOption(ReadLargeOnDemand),
		WithLargeObjectSize(1024),
	)
	if err != nil {
		t.Fatalf("Parse() error: %v", err)
	}

	elem1, ok := result.Dataset.Get(tag.New(0x7FE1, 0x0010))
	if !ok {
		t.Fatal("missing first large element")
	}
	elem2, ok := result.Dataset.Get(tag.New(0x7FE1, 0x0011))
	if !ok {
		t.Fatal("missing second large element")
	}

	const workers = 32
	var wg sync.WaitGroup
	wg.Add(workers * 2)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			got := elem1.Buffer().Data()
			if len(got) != len(data1) || got[0] != data1[0] {
				t.Errorf("first lazy buffer data mismatch: len=%d first=%d", len(got), got[0])
			}
		}()
		go func() {
			defer wg.Done()
			got := elem2.Buffer().Data()
			if len(got) != len(data2) || got[0] != data2[0] {
				t.Errorf("second lazy buffer data mismatch: len=%d first=%d", len(got), got[0])
			}
		}()
	}
	wg.Wait()
}
