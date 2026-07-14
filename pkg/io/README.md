# IO Package

This package contains reusable byte-buffer and random-access I/O helpers used by the DICOM parser and related packages.

## HTTP Range-backed DICOM parsing

Large DICOM objects can be parsed from an upstream service that supports HTTP `Range` without downloading the full object eagerly. The parser does not need a special HTTP code path. Instead, wrap the remote object with a seekable reader:

```go
package main

import (
	"context"
	"log"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/io/rangehttp"
	"github.com/cocosip/go-dicom/pkg/io/rangeio"
)

func main() {
	ctx := context.Background()

	fetcher, err := rangehttp.NewFetcher(ctx, "https://example.com/image.dcm",
		rangehttp.WithHeader("Authorization", "Bearer <token>"),
	)
	if err != nil {
		log.Fatal(err)
	}

	reader := rangeio.NewReadSeeker(fetcher,
		rangeio.WithContext(ctx),
		rangeio.WithBlockSize(512*1024),
		rangeio.WithMaxCachedBlocks(16),
	)

	result, err := parser.Parse(reader,
		parser.WithReadOption(parser.ReadLargeOnDemand),
	)
	if err != nil {
		log.Fatal(err)
	}

	_ = result.Dataset
}
```

`parser.ReadLargeOnDemand` can keep large elements and encapsulated pixel fragments lazy when the input implements `io.ReadSeeker`. The `rangeio.ReadSeeker` provides that interface while fetching the required byte blocks on demand.

## Packages

### `rangeio`

`rangeio` is transport-neutral. It defines a small random-access source interface and adapts it to `io.ReadSeeker`:

```go
type FetchRequest struct {
	Offset int64
	Length int64
}

type Fetcher interface {
	Size() int64
	Fetch(ctx context.Context, req FetchRequest) ([]byte, error)
}
```

Use this package when the backing store is not necessarily HTTP. Any source that can return exact byte ranges can implement `Fetcher`, including S3/OSS SDKs, database blobs, RPC services, or local test fixtures.

`NewReadSeeker` coalesces reads into fixed-size blocks and keeps a small in-memory cache:

- `WithBlockSize(size)` controls the range request size used for cache blocks.
- `WithMaxCachedBlocks(count)` controls how many fetched blocks remain cached.
- `WithContext(ctx)` sets the context used for fetch calls.

### `rangehttp`

`rangehttp` is the default HTTP implementation of `rangeio.Fetcher`.

`NewFetcher(ctx, url, opts...)` probes the upstream object before returning:

- It tries `HEAD` first.
- If `HEAD` is not usable, it falls back to `GET` with `Range: bytes=0-0`.
- It records the object size.
- It records `ETag` or `Last-Modified` when present.

Each `Fetch` call sends a single HTTP byte range request and validates the response:

- Sends `Range: bytes=start-end`.
- Sends `If-Range` when an `ETag` or `Last-Modified` validator is available.
- Requires `206 Partial Content`.
- Validates `Content-Range`.
- Returns `io.ErrUnexpectedEOF` if the response body is shorter than requested.

Available options:

```go
rangehttp.WithClient(client)
rangehttp.WithHeader("Authorization", "Bearer <token>")
rangehttp.WithHeaders(headers)
rangehttp.WithRequestOption(func(req *http.Request) {
	req.Header.Set("X-Tenant-ID", "tenant-1")
})
```

Use `WithRequestOption` for request customization that is easier to express with direct access to `*http.Request`, such as tenant headers, tracing headers, or per-request signing.

## Notes and limits

- The upstream must support single byte-range responses.
- Multipart range responses are not used.
- The current HTTP fetcher expects a stable object during parsing. `If-Range` helps detect changes when the server provides a validator.
- The DICOM parser still controls what data is accessed. If a high-level API reads all frames, the range reader will fetch the data needed for all frames. If the API only touches one lazy fragment/frame, only those blocks are fetched.
