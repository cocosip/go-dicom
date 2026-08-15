# DICOM Structured Reports

Package `sr` provides typed construction, reading, validation, and file I/O for
DICOM Structured Report content trees.

## Supported content

- TEXT, NUM, CODE, and CONTAINER
- PNAME, DATE, TIME, DATETIME, and UIDREF
- COMPOSITE, IMAGE, and WAVEFORM references
- SCOORD and TCOORD coordinates
- Nested by-value content and parsed by-reference content items
- Code Value, Long Code Value, and URN Code Value

The package validates root and child relationships, value-type/tag consistency,
single-item value sequences, code items, measurement units, referenced SOP UIDs,
coordinate cardinality, and nested content. Validation errors include the
`ContentSequence[index]` path to the failing item.

Template-specific relationship matrices, TID implementations, and SCOORD3D are
outside the current package scope.

## Create a report

```go
root := sr.NewCodeItem("113704", "DCM", "SR Document")
report, err := sr.NewStructuredReport(root)
if err != nil {
    return err
}

finding := sr.NewCodeItem("121071", "DCM", "Finding")
if err := report.AddText(
    finding,
    sr.RelationshipContains,
    "No abnormalities detected",
); err != nil {
    return err
}

units := sr.NewCodeItem("mm", "UCUM", "millimeter")
measurement := sr.NewMeasuredValue(25.5, units)
if err := report.AddNumeric(
    sr.NewCodeItem("33728-7", "LN", "Size"),
    sr.RelationshipContains,
    measurement,
); err != nil {
    return err
}

if err := report.Save("report.dcm"); err != nil {
    return err
}
```

## Read and write

```go
report, err := sr.Open("report.dcm")
if err != nil {
    return err
}

children, err := report.Children()
if err != nil {
    return err
}

if err := report.Write(output); err != nil {
    return err
}
```

`Open` and `Read` reject partial parses and invalid SR trees. Parsed File Meta
Information and transfer syntax are preserved by default when using `Write` or
`Save`. SR output always uses explicit lengths for Sequences and Sequence Items,
matching fo-dicom's Structured Report save behavior.

`NewStructuredReportFromDataset` remains permissive for callers that need to
inspect or repair a Dataset before calling `Validate`.

## Testing

```bash
go test ./pkg/sr
```

The test suite round-trips `test-data/test_SR.dcm`, which is byte-identical to
fo-dicom's Structured Report fixture, and exercises UIDREF, SCOORD, TCOORD,
by-reference content, nested validation paths, File Meta preservation, and
explicit-length writing.

## References

- DICOM PS3.3, Section C.17.3 and Annex A.35
- DICOM PS3.16, Structured Reporting templates and context groups
- [fo-dicom StructuredReport](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/StructuredReport)
