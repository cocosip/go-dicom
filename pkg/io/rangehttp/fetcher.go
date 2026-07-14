// Package rangehttp provides an HTTP Range implementation of rangeio.Fetcher.
package rangehttp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/io/rangeio"
)

type requestOption func(*http.Request)

// Option configures an HTTP range fetcher.
type Option func(*Fetcher)

// Fetcher implements rangeio.Fetcher using HTTP Range requests.
type Fetcher struct {
	client         *http.Client
	url            string
	headers        http.Header
	requestOptions []requestOption

	size         int64
	etag         string
	lastModified string
}

var _ rangeio.Fetcher = (*Fetcher)(nil)

// NewFetcher probes the remote object and creates a Fetcher for HTTP Range reads.
func NewFetcher(ctx context.Context, url string, opts ...Option) (*Fetcher, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	f := &Fetcher{
		client:  http.DefaultClient,
		url:     url,
		headers: make(http.Header),
	}
	for _, opt := range opts {
		if opt != nil {
			opt(f)
		}
	}
	if f.client == nil {
		f.client = http.DefaultClient
	}
	if err := f.probe(ctx); err != nil {
		return nil, err
	}
	return f, nil
}

// WithClient sets the HTTP client used for requests.
func WithClient(client *http.Client) Option {
	return func(f *Fetcher) {
		if client != nil {
			f.client = client
		}
	}
}

// WithHeader adds a default header to all requests.
func WithHeader(key, value string) Option {
	return func(f *Fetcher) {
		f.headers.Add(key, value)
	}
}

// WithHeaders adds default headers to all requests.
func WithHeaders(headers http.Header) Option {
	return func(f *Fetcher) {
		for key, values := range headers {
			for _, value := range values {
				f.headers.Add(key, value)
			}
		}
	}
}

// WithRequestOption registers a request mutator applied after default headers.
func WithRequestOption(option func(*http.Request)) Option {
	return func(f *Fetcher) {
		if option != nil {
			f.requestOptions = append(f.requestOptions, option)
		}
	}
}

// Size returns the probed object size in bytes.
func (f *Fetcher) Size() int64 {
	return f.size
}

// URL returns the configured URL.
func (f *Fetcher) URL() string {
	return f.url
}

// ETag returns the validator from the probe response, if present.
func (f *Fetcher) ETag() string {
	return f.etag
}

// LastModified returns the Last-Modified validator from the probe response, if present.
func (f *Fetcher) LastModified() string {
	return f.lastModified
}

// Fetch retrieves a byte range from the remote object.
func (f *Fetcher) Fetch(ctx context.Context, req rangeio.FetchRequest) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if req.Offset < 0 || req.Length < 0 {
		return nil, fmt.Errorf("rangehttp: invalid range offset=%d length=%d", req.Offset, req.Length)
	}
	if req.Length == 0 {
		return []byte{}, nil
	}
	end := req.Offset + req.Length - 1
	if end < req.Offset {
		return nil, fmt.Errorf("rangehttp: range overflow offset=%d length=%d", req.Offset, req.Length)
	}
	if req.Offset >= f.size || end >= f.size {
		return nil, fmt.Errorf("rangehttp: requested bytes %d-%d outside object size %d", req.Offset, end, f.size)
	}

	httpReq, err := f.newRequest(ctx, http.MethodGet)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", req.Offset, end))
	if validator := f.ifRangeValidator(); validator != "" {
		httpReq.Header.Set("If-Range", validator)
	}

	resp, err := f.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusPartialContent {
		return nil, fmt.Errorf("rangehttp: expected 206 Partial Content, got %s", resp.Status)
	}
	if err := validateContentRange(resp.Header.Get("Content-Range"), req.Offset, end, f.size); err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if int64(len(data)) != req.Length {
		return nil, io.ErrUnexpectedEOF
	}
	return data, nil
}

func (f *Fetcher) probe(ctx context.Context) error {
	if err := f.probeHEAD(ctx); err == nil {
		return nil
	}
	return f.probeRangeGET(ctx)
}

func (f *Fetcher) probeHEAD(ctx context.Context) error {
	req, err := f.newRequest(ctx, http.MethodHead)
	if err != nil {
		return err
	}
	resp, err := f.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("rangehttp: HEAD probe failed with %s", resp.Status)
	}
	size := resp.ContentLength
	if size < 0 {
		size, err = strconv.ParseInt(strings.TrimSpace(resp.Header.Get("Content-Length")), 10, 64)
		if err != nil {
			return fmt.Errorf("rangehttp: HEAD response missing valid Content-Length")
		}
	}
	f.setProbeMetadata(size, resp.Header)
	return nil
}

func (f *Fetcher) probeRangeGET(ctx context.Context) error {
	req, err := f.newRequest(ctx, http.MethodGet)
	if err != nil {
		return err
	}
	req.Header.Set("Range", "bytes=0-0")
	resp, err := f.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("rangehttp: range probe expected 206 Partial Content, got %s", resp.Status)
	}
	_, _, size, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return err
	}
	f.setProbeMetadata(size, resp.Header)
	_, _ = io.Copy(io.Discard, resp.Body)
	return nil
}

func (f *Fetcher) setProbeMetadata(size int64, header http.Header) {
	f.size = size
	f.etag = header.Get("ETag")
	f.lastModified = header.Get("Last-Modified")
}

func (f *Fetcher) newRequest(ctx context.Context, method string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, f.url, nil)
	if err != nil {
		return nil, err
	}
	for key, values := range f.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	for _, option := range f.requestOptions {
		option(req)
	}
	return req, nil
}

func (f *Fetcher) ifRangeValidator() string {
	if f.etag != "" {
		return f.etag
	}
	return f.lastModified
}

func validateContentRange(value string, wantStart, wantEnd, wantSize int64) error {
	start, end, size, err := parseContentRange(value)
	if err != nil {
		return err
	}
	if start != wantStart || end != wantEnd {
		return fmt.Errorf("rangehttp: Content-Range %q does not match requested bytes %d-%d", value, wantStart, wantEnd)
	}
	if wantSize >= 0 && size != wantSize {
		return fmt.Errorf("rangehttp: Content-Range size %d does not match probed size %d", size, wantSize)
	}
	return nil
}

func parseContentRange(value string) (start, end, size int64, err error) {
	value = strings.TrimSpace(value)
	if !strings.HasPrefix(value, "bytes ") {
		return 0, 0, 0, fmt.Errorf("rangehttp: invalid Content-Range %q", value)
	}
	parts := strings.SplitN(strings.TrimPrefix(value, "bytes "), "/", 2)
	if len(parts) != 2 {
		return 0, 0, 0, fmt.Errorf("rangehttp: invalid Content-Range %q", value)
	}
	rangeParts := strings.SplitN(parts[0], "-", 2)
	if len(rangeParts) != 2 {
		return 0, 0, 0, fmt.Errorf("rangehttp: invalid Content-Range %q", value)
	}
	start, err = strconv.ParseInt(rangeParts[0], 10, 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("rangehttp: invalid Content-Range start %q", value)
	}
	end, err = strconv.ParseInt(rangeParts[1], 10, 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("rangehttp: invalid Content-Range end %q", value)
	}
	size, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("rangehttp: invalid Content-Range size %q", value)
	}
	if start < 0 || end < start || size <= end {
		return 0, 0, 0, fmt.Errorf("rangehttp: invalid Content-Range bounds %q", value)
	}
	return start, end, size, nil
}
