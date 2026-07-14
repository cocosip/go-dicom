package parser

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"github.com/cocosip/go-dicom/pkg/io/rangehttp"
	"github.com/cocosip/go-dicom/pkg/io/rangeio"
)

func TestParseFromHTTPRangeReadSeekerKeepsFragmentsLazy(t *testing.T) {
	dicomData := createFragmentSequenceDICOM().Bytes()
	var mu sync.Mutex
	var ranges []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		ranges = append(ranges, r.Header.Get("Range"))
		mu.Unlock()

		w.Header().Set("ETag", `"range-dicom"`)
		http.ServeContent(w, r, "fragment.dcm", time.Unix(1_786_147_200, 0), bytes.NewReader(dicomData))
	}))
	defer server.Close()

	fetcher, err := rangehttp.NewFetcher(t.Context(), server.URL)
	if err != nil {
		t.Fatalf("NewFetcher() error = %v", err)
	}

	reader := rangeio.NewReadSeeker(fetcher,
		rangeio.WithContext(t.Context()),
		rangeio.WithBlockSize(8),
		rangeio.WithMaxCachedBlocks(64),
	)
	result, err := Parse(reader, WithReadOption(ReadLargeOnDemand), WithLargeObjectSize(4))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("PixelData element not found")
	}
	obf, ok := pixelDataElem.(*element.OtherByteFragment)
	if !ok {
		t.Fatalf("PixelData is %T, want *element.OtherByteFragment", pixelDataElem)
	}

	frag, err := obf.GetFragment(0)
	if err != nil {
		t.Fatalf("GetFragment(0) error = %v", err)
	}
	lazy, ok := frag.(*buffer.LazyByteBuffer)
	if !ok {
		t.Fatalf("fragment buffer = %T, want *buffer.LazyByteBuffer", frag)
	}
	if lazy.IsLoaded() {
		t.Fatal("fragment should remain unloaded after Parse")
	}
	if got := frag.Data(); !bytes.Equal(got, []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}) {
		t.Fatalf("fragment data = %v", got)
	}
	if !lazy.IsLoaded() {
		t.Fatal("fragment should be loaded after Data()")
	}

	mu.Lock()
	defer mu.Unlock()
	if !containsRangeRequest(ranges) {
		t.Fatalf("server saw ranges %q, want at least one byte range request", ranges)
	}
}

func containsRangeRequest(ranges []string) bool {
	for _, value := range ranges {
		if strings.HasPrefix(value, "bytes=") {
			return true
		}
	}
	return false
}
