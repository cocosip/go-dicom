# DICOM Reconstruction

`reconstruction` builds validated CT/MR volumes, samples arbitrary
patient-space planes, creates lazy axial/coronal/sagittal stacks, and writes
classic single-frame derived DICOM series.

## Supported workflow

Input:

- Classic single-frame CT Image Storage and MR Image Storage datasets.
- Enhanced CT Image Storage and Enhanced MR Image Storage multi-frame datasets.
- Native 8-, 16-, or 32-bit integer pixels. CT requires MONOCHROME2; MR
  accepts MONOCHROME1 or MONOCHROME2 and preserves that polarity in output.
- Encapsulated pixels when a matching codec has been registered with
  `pkg/imaging/codec`.
- Shared and per-frame Enhanced geometry and rescale transformations.
- One Enhanced Stack ID. Non-spatial Dimension Index values may be present
  only when they are constant across every reconstructed frame.

Output:

- Classic single-frame CT or MR Image Storage datasets.
- Explicit VR Little Endian, native 16-bit integer pixel data.
- New Series/SOP Instance UIDs, `DERIVED\SECONDARY\MPR` Image Type,
  derivation description, source references, and complete image-plane geometry.
- `Modality` forced from the generated CT/MR SOP Class rather than inherited
  from source metadata, with source monochrome polarity preserved where legal.
- Enhanced functional-group, dimension, concatenation, and multi-frame tags
  removed from every output instance.

Enhanced multi-frame output is intentionally outside IMG-003.

## Build a volume

```go
result, err := parser.ParseFile("source.dcm")
if err != nil {
    return err
}

images, err := reconstruction.NewImageDataFromDataset(result.Dataset)
if err != nil {
    return err
}

volume, err := reconstruction.NewVolumeData(images)
if err != nil {
    return err
}
```

For a classic series, parse each file, call `NewImageData(dataset, 0)`, and
append the returned frames before calling `NewVolumeData`.

Volume construction requires:

- a non-empty, matching Frame of Reference UID;
- at least two unique slice positions;
- matching dimensions, pixel spacing, orientation, modality, and pixel
  semantics;
- parallel source frames;
- regular slice spacing by default.

Irregular spacing must be acknowledged explicitly:

```go
volume, err := reconstruction.NewVolumeData(
    images,
    reconstruction.WithIrregularSpacingAllowed(),
)
```

Sampling still uses each adjacent pair's actual patient-space distance.

`ImageData` and `VolumeData` do not expose mutable internal fields. Metadata
accessors return values or independent Dataset/slice copies:

```go
frameGeometry := images[0].Geometry()
sourceMetadata := images[0].SourceDataset()

positions := volume.SlicePositions()
bounds := volume.Bounds()
commonMetadata := volume.CommonDataset()
```

Other inspection methods include `ImageData.FrameIndex`,
`SourceSOPClassUID`, `SourceSOPInstanceUID`, `SortingPosition`,
`RescaleSlope`, and `RescaleIntercept`, plus `VolumeData.Len`, `Normal`,
`MinSliceSpacing`, `MaxSliceSpacing`, and `IrregularSpacing`.

## Sample and cut

`VolumeData.Sample` returns a modality-space value and a validity flag. It uses
bilinear interpolation inside each source frame and linear interpolation
between adjacent slice positions. The first and last pixel/slice centers are
valid. Points outside the volume, and values depending on Pixel Padding, return
`valid == false`; they are not converted to numeric zero.

```go
value, valid, err := volume.Sample(math3d.Point3{X: 20, Y: 30, Z: 40})
```

An arbitrary plane is described with `CutSpec`:

```go
slice, err := volume.Cut(ctx, reconstruction.CutSpec{
    TopLeft:             math3d.Point3{X: 0, Y: 0, Z: 10},
    RowDirection:        math3d.Vector3{X: 1},
    ColumnDirection:     math3d.Vector3{Y: 1},
    Rows:                512,
    Columns:             512,
    PixelSpacingRows:    0.8,
    PixelSpacingColumns: 0.8,
}, reconstruction.CutOptions{Workers: 4})
```

Worker counts are bounded by the output pixel count. Workers write disjoint
indices, and output order and values do not depend on the worker count. A
non-positive worker count uses one worker. Cancellation returns the context
error.

`Slice.MinMax` ignores invalid pixels. `Slice.Render8Bit` applies the DICOM
linear window function and lets the caller select the byte used for invalid
pixels.

## Lazy standard stacks

```go
stack, err := reconstruction.NewStack(
    volume,
    reconstruction.StackTypeCoronal,
    1.0, // in-plane pixel spacing, mm
    2.0, // distance between output slices, mm
)
if err != nil {
    return err
}

slice, err := stack.Materialize(ctx, 0, reconstruction.CutOptions{Workers: 4})
```

`Stack` stores only slice specifications. Pixel arrays are allocated only by
`Materialize`, `Stream`, or `Generate`. Standard-plane counts include both
patient-space endpoints; a non-divisible final interval is shortened so the
last specification lands on the volume bound.

## Generate derived DICOM

Use `Stream` for bounded memory:

```go
generator, err := reconstruction.NewDicomGenerator(volume)
if err != nil {
    return err
}

err = generator.Stream(ctx, stack, "Coronal MPR",
    reconstruction.CutOptions{Workers: 4},
    func(index int, output *dataset.Dataset) error {
        path := fmt.Sprintf("mpr-%04d.dcm", index+1)
        return writer.WriteFile(path, output,
            writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian))
    },
)
```

`Generate` is a convenience wrapper that returns all datasets in memory.
`WithGeneratorClock` and `WithGeneratorUIDFactory` make dates and UIDs
deterministic in tests or controlled workflows.

Generated samples are modality-space values rounded to integers. Unsigned
output is selected when every valid value is non-negative; otherwise signed
output is used. An invalid-pixel mask reserves the corresponding 16-bit extreme
as Pixel Padding Value. Generation fails instead of clipping when valid values
do not fit the remaining 16-bit range.

Floating-point geometry, spacing, window, and rescale values are encoded as
valid DICOM DS values of at most 16 characters. Non-finite values are rejected.

## Executable example

The repository example reads one Enhanced file or multiple classic files and
writes a derived series:

```text
go run ./examples/reconstruction \
  -plane coronal \
  -spacing 1 \
  -slice-distance 2 \
  -output ./mpr-output \
  source-001.dcm source-002.dcm source-003.dcm
```

Use `-allow-irregular` only after deciding that interpolation across irregular
source spacing is appropriate for the application.

## Deliberate limits

- CT and MR only; no PET, parametric map, segmentation, color, or floating
  pixel sources.
- One compatible spatial stack only; temporal, cardiac, diffusion, and other
  Enhanced dimensions may be declared only when their Dimension Index Values
  stay constant; varying non-spatial dimensions and multiple Stack IDs are
  rejected.
- Source limits are 65,535 frames, 16,384 rows/columns, and 1 GiB of expanded
  integer pixel data. Output limits are 65,535 stack slices, DICOM US-range
  rows/columns, and 4,096 x 4,096 pixels per materialized cut.
- No gantry-tilt correction beyond sampling parallel oblique source planes.
- No thick-slab projection or slice-thickness integration.
- No Enhanced CT/MR output.
- No implicit clipping or automatic resampling of incompatible input series.

Applications should validate generated instances against their target archive
and modality-specific profile before clinical use.
