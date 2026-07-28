package rangeio

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"
)

type recordingFetcher struct {
	data  []byte
	calls []FetchRequest
}

type benchmarkFetcher struct {
	data []byte
}

func (f benchmarkFetcher) Size() int64 {
	return int64(len(f.data))
}

func (f benchmarkFetcher) Fetch(_ context.Context, req FetchRequest) ([]byte, error) {
	return f.data[req.Offset : req.Offset+req.Length], nil
}

func (f *recordingFetcher) Size() int64 {
	return int64(len(f.data))
}

func (f *recordingFetcher) Fetch(_ context.Context, req FetchRequest) ([]byte, error) {
	f.calls = append(f.calls, req)
	if req.Offset < 0 || req.Length < 0 {
		return nil, errors.New("invalid request")
	}
	if req.Offset >= int64(len(f.data)) {
		return []byte{}, nil
	}
	end := req.Offset + req.Length
	if end > int64(len(f.data)) {
		end = int64(len(f.data))
	}
	return append([]byte(nil), f.data[req.Offset:end]...), nil
}

func TestReadSeekerFetchesBlocksAndCachesThem(t *testing.T) {
	fetcher := &recordingFetcher{data: []byte("abcdefghijklmnopqrstuvwxyz")}
	rs := NewReadSeeker(fetcher, WithBlockSize(8), WithMaxCachedBlocks(2))

	buf := make([]byte, 5)
	n, err := rs.Read(buf)
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	if n != 5 || string(buf) != "abcde" {
		t.Fatalf("Read() = %d %q, want 5 %q", n, buf, "abcde")
	}
	if len(fetcher.calls) != 1 || fetcher.calls[0] != (FetchRequest{Offset: 0, Length: 8}) {
		t.Fatalf("Fetch calls = %+v, want one 0..8 block", fetcher.calls)
	}

	pos, err := rs.Seek(1, io.SeekStart)
	if err != nil {
		t.Fatalf("Seek() error = %v", err)
	}
	if pos != 1 {
		t.Fatalf("Seek() = %d, want 1", pos)
	}

	buf = make([]byte, 4)
	n, err = rs.Read(buf)
	if err != nil {
		t.Fatalf("Read() after seek error = %v", err)
	}
	if n != 4 || string(buf) != "bcde" {
		t.Fatalf("Read() after seek = %d %q, want 4 %q", n, buf, "bcde")
	}
	if len(fetcher.calls) != 1 {
		t.Fatalf("Read() should reuse cached block, calls = %+v", fetcher.calls)
	}
}

func TestReadSeekerReadsAcrossBlocks(t *testing.T) {
	fetcher := &recordingFetcher{data: []byte("abcdefghijklmnopqrstuvwxyz")}
	rs := NewReadSeeker(fetcher, WithBlockSize(8))

	if _, err := rs.Seek(6, io.SeekStart); err != nil {
		t.Fatalf("Seek() error = %v", err)
	}

	buf := make([]byte, 6)
	n, err := rs.Read(buf)
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	if n != 6 || string(buf) != "ghijkl" {
		t.Fatalf("Read() = %d %q, want 6 %q", n, buf, "ghijkl")
	}

	want := []FetchRequest{{Offset: 0, Length: 8}, {Offset: 8, Length: 8}}
	if !equalFetchRequests(fetcher.calls, want) {
		t.Fatalf("Fetch calls = %+v, want %+v", fetcher.calls, want)
	}
}

func TestReadSeekerEOFAndSeekBounds(t *testing.T) {
	fetcher := &recordingFetcher{data: []byte("abc")}
	rs := NewReadSeeker(fetcher, WithBlockSize(8))

	if _, err := rs.Seek(-1, io.SeekStart); err == nil {
		t.Fatal("Seek(-1) error = nil, want error")
	}

	pos, err := rs.Seek(2, io.SeekStart)
	if err != nil {
		t.Fatalf("Seek(2) error = %v", err)
	}
	if pos != 2 {
		t.Fatalf("Seek(2) = %d, want 2", pos)
	}

	buf := make([]byte, 4)
	n, err := rs.Read(buf)
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	if n != 1 || !bytes.Equal(buf[:n], []byte("c")) {
		t.Fatalf("Read() = %d %q, want 1 %q", n, buf[:n], "c")
	}

	n, err = rs.Read(buf)
	if n != 0 || err != io.EOF {
		t.Fatalf("Read() at EOF = %d, %v; want 0, io.EOF", n, err)
	}
}

func BenchmarkReadSeekerCacheHit(b *testing.B) {
	data := make([]byte, 64*1024)
	reader := NewReadSeeker(benchmarkFetcher{data: data}, WithBlockSize(64*1024))
	buf := make([]byte, 4*1024)
	if _, err := reader.Read(buf); err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.SetBytes(int64(len(buf)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := reader.Seek(0, io.SeekStart); err != nil {
			b.Fatal(err)
		}
		if n, err := reader.Read(buf); err != nil || n != len(buf) {
			b.Fatalf("Read() = %d, %v", n, err)
		}
	}
}

func BenchmarkReadSeekerCrossBlockRead(b *testing.B) {
	data := make([]byte, 128*1024)
	fetcher := benchmarkFetcher{data: data}
	buf := make([]byte, 96*1024)

	b.ReportAllocs()
	b.SetBytes(int64(len(buf)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := NewReadSeeker(fetcher, WithBlockSize(64*1024), WithMaxCachedBlocks(2))
		if n, err := reader.Read(buf); err != nil || n != len(buf) {
			b.Fatalf("Read() = %d, %v", n, err)
		}
	}
}

func equalFetchRequests(got, want []FetchRequest) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
