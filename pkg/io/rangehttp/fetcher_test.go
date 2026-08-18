package rangehttp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/rangeio"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

type countingReadCloser struct {
	reader *bytes.Reader
	read   int
}

func (r *countingReadCloser) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.read += n
	return n, err
}

func (r *countingReadCloser) Close() error { return nil }

func TestFetcherProbesAndFetchesRanges(t *testing.T) {
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	var sawHead bool
	var sawGet bool

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer token" {
			t.Fatalf("Authorization header = %q, want %q", got, "Bearer token")
		}
		if got := r.Header.Get("X-Tenant-ID"); got != "tenant-1" {
			t.Fatalf("X-Tenant-ID header = %q, want %q", got, "tenant-1")
		}

		switch r.Method {
		case http.MethodHead:
			sawHead = true
			w.Header().Set("Accept-Ranges", "bytes")
			w.Header().Set("Content-Length", fmt.Sprint(len(data)))
			w.Header().Set("ETag", `"abc123"`)
			w.Header().Set("Last-Modified", "Tue, 14 Jul 2026 00:00:00 GMT")
		case http.MethodGet:
			sawGet = true
			if got := r.Header.Get("Range"); got != "bytes=2-5" {
				t.Fatalf("Range header = %q, want %q", got, "bytes=2-5")
			}
			if got := r.Header.Get("If-Range"); got != `"abc123"` {
				t.Fatalf("If-Range header = %q, want %q", got, `"abc123"`)
			}
			w.Header().Set("Content-Range", fmt.Sprintf("bytes 2-5/%d", len(data)))
			w.WriteHeader(http.StatusPartialContent)
			_, _ = w.Write(data[2:6])
		default:
			t.Fatalf("unexpected method %s", r.Method)
		}
	}))
	defer server.Close()

	fetcher, err := NewFetcher(context.Background(), server.URL,
		WithHeader("Authorization", "Bearer token"),
		WithRequestOption(func(req *http.Request) {
			req.Header.Set("X-Tenant-ID", "tenant-1")
		}),
	)
	if err != nil {
		t.Fatalf("NewFetcher() error = %v", err)
	}
	if fetcher.Size() != int64(len(data)) {
		t.Fatalf("Size() = %d, want %d", fetcher.Size(), len(data))
	}
	if fetcher.ETag() != `"abc123"` {
		t.Fatalf("ETag() = %q, want %q", fetcher.ETag(), `"abc123"`)
	}
	if fetcher.LastModified() != "Tue, 14 Jul 2026 00:00:00 GMT" {
		t.Fatalf("LastModified() = %q", fetcher.LastModified())
	}

	got, err := fetcher.Fetch(context.Background(), rangeio.FetchRequest{Offset: 2, Length: 4})
	if err != nil {
		t.Fatalf("Fetch() error = %v", err)
	}
	if string(got) != "cdef" {
		t.Fatalf("Fetch() = %q, want %q", got, "cdef")
	}
	if !sawHead || !sawGet {
		t.Fatalf("sawHead=%v sawGet=%v, want both true", sawHead, sawGet)
	}
}

func TestFetcherRejectsNonPartialRangeResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Content-Length", "6")
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("abcdef"))
		}
	}))
	defer server.Close()

	fetcher, err := NewFetcher(context.Background(), server.URL)
	if err != nil {
		t.Fatalf("NewFetcher() error = %v", err)
	}

	if _, err := fetcher.Fetch(context.Background(), rangeio.FetchRequest{Offset: 0, Length: 3}); err == nil {
		t.Fatal("Fetch() error = nil, want error for non-206 response")
	}
}

func TestFetcherRejectsOutOfBoundsRequestBeforeHTTPGet(t *testing.T) {
	var getCount int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Content-Length", "6")
		if r.Method == http.MethodGet {
			getCount++
		}
	}))
	defer server.Close()

	fetcher, err := NewFetcher(context.Background(), server.URL)
	if err != nil {
		t.Fatalf("NewFetcher() error = %v", err)
	}

	if _, err := fetcher.Fetch(context.Background(), rangeio.FetchRequest{Offset: 4, Length: 3}); err == nil {
		t.Fatal("Fetch() error = nil, want error for out-of-bounds request")
	}
	if getCount != 0 {
		t.Fatalf("GET count = %d, want 0", getCount)
	}
}

func TestFetcherFallsBackToRangeGetProbe(t *testing.T) {
	data := []byte("abcdef")
	var probeRange string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method %s", r.Method)
		}
		if r.Header.Get("Range") == "bytes=0-0" {
			probeRange = r.Header.Get("Range")
			w.Header().Set("Accept-Ranges", "bytes")
			w.Header().Set("Content-Range", fmt.Sprintf("bytes 0-0/%d", len(data)))
			w.WriteHeader(http.StatusPartialContent)
			_, _ = w.Write(data[:1])
			return
		}
		w.Header().Set("Content-Range", fmt.Sprintf("bytes 1-2/%d", len(data)))
		w.WriteHeader(http.StatusPartialContent)
		_, _ = w.Write(data[1:3])
	}))
	defer server.Close()

	fetcher, err := NewFetcher(context.Background(), server.URL)
	if err != nil {
		t.Fatalf("NewFetcher() error = %v", err)
	}
	if fetcher.Size() != int64(len(data)) {
		t.Fatalf("Size() = %d, want %d", fetcher.Size(), len(data))
	}
	if probeRange != "bytes=0-0" {
		t.Fatalf("probe Range = %q, want bytes=0-0", probeRange)
	}

	got, err := fetcher.Fetch(context.Background(), rangeio.FetchRequest{Offset: 1, Length: 2})
	if err != nil {
		t.Fatalf("Fetch() error = %v", err)
	}
	if string(got) != "bc" {
		t.Fatalf("Fetch() = %q, want %q", got, "bc")
	}
}

func TestFetcherBoundsOversizedRangeResponseRead(t *testing.T) {
	body := &countingReadCloser{reader: bytes.NewReader(bytes.Repeat([]byte{'x'}, 1024))}
	client := &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusPartialContent,
			Status:     "206 Partial Content",
			Header:     http.Header{"Content-Range": []string{"bytes 0-3/10"}},
			Body:       body,
		}, nil
	})}
	fetcher := &Fetcher{client: client, url: "http://example.test/object", headers: make(http.Header), size: 10}

	if _, err := fetcher.Fetch(context.Background(), rangeio.FetchRequest{Offset: 0, Length: 4}); err == nil {
		t.Fatal("Fetch() error = nil, want error for oversized response body")
	}
	if body.read > 5 {
		t.Fatalf("Fetch() read %d bytes, want at most expected length plus one", body.read)
	}
}

func TestFetcherBoundsOversizedRangeProbeRead(t *testing.T) {
	body := &countingReadCloser{reader: bytes.NewReader(bytes.Repeat([]byte{'x'}, 1024))}
	client := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodHead {
			return &http.Response{
				StatusCode: http.StatusMethodNotAllowed,
				Status:     "405 Method Not Allowed",
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader(nil)),
			}, nil
		}
		return &http.Response{
			StatusCode: http.StatusPartialContent,
			Status:     "206 Partial Content",
			Header:     http.Header{"Content-Range": []string{"bytes 0-0/10"}},
			Body:       body,
		}, nil
	})}

	if _, err := NewFetcher(context.Background(), "http://example.test/object", WithClient(client)); err == nil {
		t.Fatal("NewFetcher() error = nil, want error for oversized probe body")
	}
	if body.read > 2 {
		t.Fatalf("range probe read %d bytes, want at most expected length plus one", body.read)
	}
}
