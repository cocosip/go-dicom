# DICOM Media Directory

Package `media` provides DICOMDIR creation, reading, traversal, offset recovery,
and two-pass writing. It supports patient, study, series, image, SR document,
and presentation records while preserving unknown record types when reading.

## Create

```go
directory, err := media.NewDirectory()
if err != nil {
    return err
}

file, err := parser.ParseFile("IMAGES/IMG0001")
if err != nil {
    return err
}
fileID, err := media.ParseFileID("IMAGES/IMG0001")
if err != nil {
    return err
}
if _, err := directory.AddFile(file, fileID); err != nil {
    return err
}
return directory.Save("DICOMDIR")
```

`AddFile` records the supplied File ID but does not scan, copy, move, rename, or
rewrite the referenced file. File IDs contain one to eight components; each
component contains one to eight uppercase ASCII letters, digits, or underscores.

## Read

Strict mode rejects every stale, missing, cyclic, duplicate, or unreachable
record reference:

```go
directory, err := media.Open("DICOMDIR", media.WithOffsetPolicy(media.StrictOffsets))
```

Compatible mode is the default. It repairs offsets only when a fixed offset
delta or the supported physical record hierarchy produces one complete,
acyclic result. Ambiguous recovery fails instead of selecting the nearest item.
Successful repairs are available through `Directory.Diagnostics()`.

## Icons

Icon generation is opt-in and structurally decoupled from this package:

```go
directory, err := media.NewDirectory(
    media.WithImageIcons(true),
    media.WithIconGenerator(imaging.NewDirectoryIconGenerator()),
)
```

The imaging generator renders a representative frame, preserves aspect ratio,
does not upscale small images, and produces an 8-bit `MONOCHROME2` icon no
larger than 128x128. Compressed images require their codec to be registered in
the global imaging codec registry. Icon failures are non-fatal and add a
non-identifying diagnostic.

## Ownership

Directory mutation and writing are not safe for concurrent use. Returned root,
child, and diagnostic slices are defensive copies, while record Dataset
pointers expose the stored DICOM attributes. The directory does not mutate the
`parser.ParseResult`, Dataset, File Meta Information, or pixel buffers passed to
`AddFile`.
