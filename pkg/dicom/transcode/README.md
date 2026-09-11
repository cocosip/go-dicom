# Codec and Dataset Transcoding

The `imaging/codec` package owns only the frame-level codec SPI, native
transfer-syntax codecs, and codec registration. This `dicom/transcode` package
owns Dataset-level transcoding, File Meta Information updates, Pixel Data VR
selection, and lossy metadata history. Compressed codec implementations live in
`go-dicom-codecs`; the dependency direction is always
`go-dicom-codecs -> go-dicom`.

## Frame Contract

Codecs read frames through `FrameSource` and write through `FrameSink`:

```go
type FrameSource interface {
    FrameCount() int
    Frame(context.Context, int) ([]byte, error)
    FrameInfo() FrameInfo
    Encapsulated() bool
}

type FrameSink interface {
    AddFrame(context.Context, []byte) error
    SetFrameInfo(FrameInfo) error
}
```

`Frame` returns a caller-owned copy. `AddFrame` copies its input before the
call returns. Both methods honor cancellation while copying. `FrameInfo` is
passed by value and uses the shared value objects from `imaging/pixel`.

An output codec must return an error from `SetFrameInfo` instead of ignoring a
metadata update failure.

## Codec SPI

```go
type Codec interface {
    Name() string
    TransferSyntax() *transfer.Syntax
    DefaultParameters() Parameters
    Encode(context.Context, FrameSource, FrameSink, Parameters) error
    Decode(context.Context, FrameSource, FrameSink, Parameters) error
}
```

There is one context-aware encode/decode entry. Implementations must preserve
`context.Canceled` and `context.DeadlineExceeded` in the returned error chain.

## Parameters

Each codec owns a strongly typed parameter struct:

```go
type Parameters interface {
    Clone() Parameters
    Validate() error
}
```

`PrepareParameters` chooses user-supplied or default parameters, deep-clones
them for one call, and validates the clone. A validation failure wraps
`ErrInvalidParameters`. Codecs without settings use `NoParameters`.

Native byte swapping uses a three-state value so the Big Endian decode default
can be explicitly overridden:

```go
params := codec.NativeParameters{
    ByteSwap: codec.ByteSwapDisabled,
}
```

- `ByteSwapDefault` follows the native codec's transfer-syntax behavior.
- `ByteSwapDisabled` preserves source byte order.
- `ByteSwapEnabled` swaps bytes within each 16-bit native OW word.

## Pixel Data

`pixeldata.Data` is the public Pixel Data owner and implements both frame ports.
It retains Dataset extraction, Pixel Data element write-back, frame/BOT access,
sample reads, photometric conversion, windowing, padding, and codec operations.

```go
pixels, err := pixeldata.FromDataset(ds)
if err != nil {
    return err
}

decoded, err := pixels.Decode(ctx, registeredCodec, nil)
if err != nil {
    return err
}

pixelElement, err := decoded.ToElement()
```

A nil parameter value requests the codec defaults. `Data.Encode` and
`Data.Decode` pass only an owned, validated parameter clone to the codec.

## Dataset Transcoding

Bind a `Manager` to one explicit registry, then create a transcoder for the
source and destination transfer syntaxes:

```go
registry := codec.NewRegistry()
// Register application-owned compressed codecs on registry when needed.

manager, err := transcode.NewManager(registry)
if err != nil {
    return err
}

transcoder, err := manager.NewTranscoder(
    sourceSyntax,
    destinationSyntax,
    transcode.WithInputParameters(inputParameters),
    transcode.WithOutputParameters(outputParameters),
)
if err != nil {
    return err
}

result, err := transcoder.Transcode(ctx, ds)
```

The same transcoder also provides:

- `TranscodeWithMetadata` for Dataset plus File Meta Information.
- `DecodeFrame` for one encapsulated frame.
- `WithPixelDataReadMode` for standard validation or compatibility recovery.
- `WithPixelDataWriteMode` for standard or compatibility output VR selection.
- `WithStrictDICOMVR` as the compatibility entry point for the write mode.

All three operations accept `context.Context` as their first argument. There
are no non-context wrappers and no constructor that bypasses `Manager`.

Native-to-native conversion does not require a registered codec. Encapsulated
input or output requires the matching compressed codec.

## Registration

Applications that use compressed transfer syntaxes must register their codec
implementations before creating a Manager. Registration only supplies the
implementation; association negotiation still determines whether a peer
accepted the transfer syntax.

Use an isolated `Registry` when lifecycle isolation matters. The process-wide
registry is reserved for runtime codec plugins and remains mutable and
thread-safe.

```go
isolated := codec.NewRegistry()
if err := isolated.Register(compressedCodec); err != nil {
    return err
}

pluginRegistry := codec.GlobalRegistry()
previous, err := pluginRegistry.Replace(updatedCodec)
_ = previous
```

`Register` rejects duplicate transfer syntax UIDs; replacement must be
explicit through `Replace`. `Unregister`, `Lookup`, and `List` provide removal,
query, and a UID-sorted independent snapshot. Both `NewRegistry` and
`GlobalRegistry` initially contain the native codecs. Only composition roots
should select the global registry and pass the resulting Registry or Manager
to internal services.

General byte-order conversion is owned by `io/endian`:

```go
swapped, err := endian.ConvertEndianness(pixelBytes, bytesPerSample)
```

Native DICOM Pixel Data with VR OW is different: its endian unit is always one
16-bit word, even when `BitsAllocated` is 32 or greater. The transcoder handles
that conversion and callers should not reverse an entire multi-word sample.

## Pixel Data VR

- Reads default to compatibility mode; non-standard native OB and encapsulated
  OW are accepted with a structured warning.
- Encapsulated OW is treated as an opaque compressed byte stream in compatibility
  mode; it is not endian-swapped. Standard output normalizes it to OB.
- Writes default to standard mode; encapsulated Pixel Data uses OB.
- Standard read mode rejects native OB where OW is required and rejects
  encapsulated OW.
- Compatibility write mode may emit OW for multi-byte encapsulated data.
- Implicit VR Little Endian native Pixel Data uses OW.
- Explicit VR native Pixel Data uses OW above 8 Bits Allocated and may use OB
  or OW at 8 Bits Allocated or below.

## Verification

```powershell
go test ./pkg/imaging/codec ./pkg/imaging/pixeldata ./pkg/dicom/transcode -count=1
```

The tests cover frame ownership, cancellation, parameter ownership and
validation, NativeCodec byte swapping, fragment/BOT handling, VR selection,
metadata write-back, and Dataset transcoding.
