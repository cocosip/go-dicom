// Package rangeio adapts random-access byte fetchers into seekable readers.
package rangeio

import (
	"context"
	"errors"
	"io"
	"sync"
)

const defaultBlockSize int64 = 512 * 1024

// FetchRequest describes a byte range to retrieve from a random-access source.
type FetchRequest struct {
	Offset int64
	Length int64
}

// Fetcher retrieves byte ranges from a finite random-access source.
type Fetcher interface {
	Size() int64
	Fetch(ctx context.Context, req FetchRequest) ([]byte, error)
}

// Option configures a ReadSeeker.
type Option func(*ReadSeeker)

// WithBlockSize sets the cache block size used to coalesce small reads.
func WithBlockSize(size int64) Option {
	return func(r *ReadSeeker) {
		if size > 0 {
			r.blockSize = size
		}
	}
}

// WithMaxCachedBlocks sets the maximum number of blocks retained in memory.
func WithMaxCachedBlocks(maxBlocks int) Option {
	return func(r *ReadSeeker) {
		if maxBlocks >= 0 {
			r.maxCachedBlocks = maxBlocks
		}
	}
}

// WithContext sets the context used for range fetches.
func WithContext(ctx context.Context) Option {
	return func(r *ReadSeeker) {
		if ctx != nil {
			r.ctx = ctx
		}
	}
}

// ReadSeeker adapts a Fetcher into an io.ReadSeeker.
type ReadSeeker struct {
	fetcher Fetcher
	size    int64
	pos     int64

	ctx             context.Context
	blockSize       int64
	maxCachedBlocks int
	cache           map[int64][]byte
	cacheOrder      []int64

	mu sync.Mutex
}

var _ io.ReadSeeker = (*ReadSeeker)(nil)

// NewReadSeeker creates a seekable reader over a range fetcher.
func NewReadSeeker(fetcher Fetcher, opts ...Option) *ReadSeeker {
	r := &ReadSeeker{
		fetcher:         fetcher,
		ctx:             context.Background(),
		blockSize:       defaultBlockSize,
		maxCachedBlocks: 8,
		cache:           make(map[int64][]byte),
	}
	if fetcher != nil {
		r.size = fetcher.Size()
	}
	for _, opt := range opts {
		if opt != nil {
			opt(r)
		}
	}
	if r.blockSize <= 0 {
		r.blockSize = defaultBlockSize
	}
	return r
}

func (r *ReadSeeker) Read(p []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(p) == 0 {
		return 0, nil
	}
	if r.fetcher == nil {
		return 0, errors.New("rangeio: nil fetcher")
	}
	if r.pos >= r.size {
		return 0, io.EOF
	}

	total := 0
	for total < len(p) && r.pos < r.size {
		blockIndex := r.pos / r.blockSize
		block, err := r.block(blockIndex)
		if err != nil {
			if total > 0 {
				return total, err
			}
			return 0, err
		}
		offsetInBlock := r.pos - blockIndex*r.blockSize
		if offsetInBlock >= int64(len(block)) {
			break
		}

		n := copy(p[total:], block[offsetInBlock:])
		total += n
		r.pos += int64(n)
	}

	if r.pos >= r.size {
		if total > 0 {
			return total, nil
		}
		return 0, io.EOF
	}
	if total == 0 {
		return 0, io.EOF
	}
	return total, nil
}

// Seek sets the offset for the next Read.
func (r *ReadSeeker) Seek(offset int64, whence int) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var next int64
	switch whence {
	case io.SeekStart:
		next = offset
	case io.SeekCurrent:
		next = r.pos + offset
	case io.SeekEnd:
		next = r.size + offset
	default:
		return 0, errors.New("rangeio: invalid seek whence")
	}
	if next < 0 {
		return 0, errors.New("rangeio: negative position")
	}
	r.pos = next
	return r.pos, nil
}

func (r *ReadSeeker) block(index int64) ([]byte, error) {
	if data, ok := r.cache[index]; ok {
		return data, nil
	}

	offset := index * r.blockSize
	length := r.blockSize
	if remaining := r.size - offset; remaining < length {
		length = remaining
	}
	if length <= 0 {
		return nil, io.EOF
	}

	data, err := r.fetcher.Fetch(r.ctx, FetchRequest{Offset: offset, Length: length})
	if err != nil {
		return nil, err
	}
	if int64(len(data)) != length {
		return nil, io.ErrUnexpectedEOF
	}

	r.cache[index] = data
	r.cacheOrder = append(r.cacheOrder, index)
	r.evictCache()
	return data, nil
}

func (r *ReadSeeker) evictCache() {
	if r.maxCachedBlocks < 0 {
		return
	}
	for len(r.cacheOrder) > r.maxCachedBlocks {
		oldest := r.cacheOrder[0]
		copy(r.cacheOrder, r.cacheOrder[1:])
		r.cacheOrder = r.cacheOrder[:len(r.cacheOrder)-1]
		delete(r.cache, oldest)
	}
}
