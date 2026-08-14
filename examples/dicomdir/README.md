# DICOMDIR Example

This example creates a DICOMDIR from explicit DICOM file arguments. It does not
scan directories, copy files, rename files, or change the supplied DICOM files.

Each input has the form `FILE_ID=path`. The File ID is the relative DICOM media
identifier stored in the directory record, not a filesystem operation.

```powershell
go run ./examples/dicomdir `
  -output DICOMDIR `
  'IMAGES/IMG0001=D:\media\IMAGES\IMG0001' `
  'IMAGES/IMG0002=D:\media\IMAGES\IMG0002'
```

Add `-icons` to generate optional, aspect-preserving 8-bit grayscale icons with
a maximum size of 128x128 pixels:

```powershell
go run ./examples/dicomdir -icons 'IMAGES/IMG0001=D:\media\IMAGES\IMG0001'
```

Compressed source images require the matching pure-Go codec to be registered in
the global imaging codec registry before icon generation. Directory creation
continues without an icon when rendering fails and records a diagnostic.

After saving, the example reopens the result with strict offset validation and
prints only record type counts.
