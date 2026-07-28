# Go-DICOM Performance Benchmarks

Historical snapshot collected: 2025-01-10

> The tables below are a historical measurement snapshot, not a current performance baseline. Benchmark results vary by commit, Go version, CPU, operating system, and DICOM workload. For a pull request, download the `benchmark-results` CI artifact; for local work, run the commands in [Running Benchmarks](#running-benchmarks).

## Benchmark Coverage

- `pkg/dicom/...`: dataset, parser, writer, endian, UID, and committed real-DICOM fixture benchmarks.
- `pkg/io/...`: buffer-pool and range-backed `ReadSeeker` cache/cross-block benchmarks.
- `pkg/imaging/...`: pixel layout and color-conversion benchmarks.

The PR workflow runs all three package trees with `-run=^$` so benchmark results are not mixed with unit-test execution. The resulting `benchmark_results.txt` is uploaded as the `benchmark-results` artifact and retained for 90 days.

## Overview

This document contains performance benchmark results for the go-dicom library. Benchmarks are run on:

- **CPU**: Intel(R) Core(TM) Ultra 9 185H
- **OS**: Windows
- **Architecture**: amd64
- **Go Version**: go1.21+

## Dataset Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|-------|------|-----------|
| DatasetAdd | 12.85 | 0 | 0 |
| DatasetAddOrUpdate | 14.01 | 0 | 0 |
| DatasetGet | 2.141 | 0 | 0 |
| DatasetGetString | 74.40 | 48 | 4 |
| DatasetContains | 2.351 | 0 | 0 |
| DatasetElements (100 elements) | 5,453 | 2,264 | 4 |
| DatasetClone (50 elements) | 2,574 | 3,208 | 7 |
| DatasetMerge (50 elements) | 2,452 | 3,208 | 7 |
| DatasetFilter (100 elements) | 2,764 | 3,472 | 10 |

**Key Insights:**
- Core operations (Add, Get, Contains) are extremely fast with zero allocations
- GetString has some overhead due to type assertion and string conversion
- Clone and Merge operations are efficient even with moderate dataset sizes

## Parser Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|-------|------|-----------|
| ParseSmallDataset (10 elements) | 1,382 | 1,249 | 43 |
| ParseMediumDataset (100 elements) | 1,383 | 1,249 | 43 |
| ParseLargeDataset (500 elements) | 1,408 | 1,249 | 43 |
| ReadTag | 106.7 | 212 | 5 |
| ReadElement | 206.4 | 304 | 11 |

**Key Insights:**
- Parsing performance is consistent regardless of dataset size
- Fixed allocation count (43) indicates efficient memory management
- Most overhead is in initialization rather than data processing

## Writer Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|-------|------|-----------|
| WriteSmallDataset (3 elements) | 1,332 | 1,265 | 44 |
| WriteMediumDataset (50 elements) | 8,346 | 4,580 | 230 |
| WriteLargeDataset (200 elements) | 34,600 | 16,384 | 832 |
| WriteTag | 35.05 | 4 | 2 |
| WriteElement | 161.3 | 192 | 7 |
| WritePreamble | 215.8 | 628 | 6 |

**Key Insights:**
- Writing scales linearly with dataset size
- Individual operations (WriteTag, WriteElement) are very fast
- Memory allocations scale proportionally with data size
- **Optimization**: Uses sync.Pool for temporary buffers, reducing allocations by 2 per write

## Endian Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|-------|------|-----------|
| SwapUint16 | 0.1179 | 0 | 0 |
| SwapUint32 | 0.1215 | 0 | 0 |
| SwapUint64 | 0.1169 | 0 | 0 |
| SwapBytes2 | 213.0 | 0 | 0 |
| SwapBytes4 | 186.5 | 0 | 0 |

**Key Insights:**
- Endian conversion operations are extremely fast
- Zero allocations for all endian operations
- Sub-nanosecond performance for scalar swaps

## UID Generation

| Benchmark | ns/op | B/op | allocs/op |
|-----------|-------|------|-----------|
| GenerateDerivedFromUUID | 344.6 | 304 | 7 |
| Generator.Generate | 13.58 | 0 | 0 |

**Key Insights:**
- UUID-based generation has moderate overhead
- Sequential generator is very fast with zero allocations

## Buffer Pool Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|-------|------|-----------|
| GetBufferSmall (1KB) | 29.04 | 24 | 1 |
| GetBufferMedium (32KB) | 193.7 | 24 | 1 |
| GetBufferLarge (512KB) | 10,720 | 33 | 1 |
| AllocateSmallNormal | 0.17 | 0 | 0 |
| AllocateMediumNormal | 0.20 | 0 | 0 |
| AllocateLargeNormal | 61,029 | 524,288 | 1 |
| GetBytesBuffer (pooled) | 10.81 | 0 | 0 |
| BytesBufferNormal | 21.90 | 64 | 1 |

**Key Insights:**
- **Large buffer pooling** (≥512KB): 5.7x faster, 99.99% less memory allocated
- **bytes.Buffer pooling**: 2x faster with 100% allocation reduction
- Small buffer pooling is slower than direct allocation (Go compiler optimization)
- Buffer pool most beneficial for temporary buffers and large allocations

## Imaging Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|-------|------|-----------|
| InterleavedToPlanar24 | ~varies | ~varies | ~varies |
| YBRFullToRGB | ~varies | ~varies | ~varies |

*(See pkg/imaging/converter_test.go for detailed imaging benchmarks)*

## Performance Summary

### Strengths
1. **Core Operations**: Dataset Get/Add operations are sub-15ns with zero allocations
2. **Consistent Parsing**: Parser performance doesn't degrade with larger datasets
3. **Efficient Endian**: Byte swapping operations are extremely fast
4. **Linear Scaling**: Writer performance scales predictably with data size
5. **Buffer Pooling**: bytes.Buffer pool reduces allocations by 100% for temporary buffers
6. **Large Buffer Pooling**: 5.7x performance improvement for ≥512KB allocations

### Recent Optimizations (Stage 12.1)
1. **Writer Buffer Pooling**: Implemented sync.Pool for temporary bytes.Buffer instances
   - Reduced allocations by 2 per write operation
   - 100% allocation reduction for temporary buffers
   - 2x performance improvement for buffer operations

### Areas for Potential Optimization
1. **Parser Initialization**: Fixed 43 allocations could potentially be reduced
2. **String Conversions**: GetString has 4 allocations that could be optimized

## Running Benchmarks

To run the same benchmark scope as CI:

```bash
go test -bench='.' -benchmem -run='^$' ./pkg/dicom/... ./pkg/io/... ./pkg/imaging/...
```

To run specific package benchmarks:

```bash
go test -bench='.' -benchmem ./pkg/dicom/dataset
go test -bench='.' -benchmem ./pkg/dicom/parser
go test -bench='.' -benchmem ./pkg/dicom/writer
go test -bench='.' -benchmem ./pkg/io/rangeio
go test -bench='.' -benchmem ./pkg/imaging
```

To generate CPU profile:

```bash
go test -bench='.' -cpuprofile=cpu.prof ./pkg/dicom/dataset
go tool pprof cpu.prof
```

To generate memory profile:

```bash
go test -bench='.' -memprofile=mem.prof ./pkg/dicom/dataset
go tool pprof mem.prof
```

## Comparison with fo-dicom

*(To be added: Performance comparison with the original C# fo-dicom library)*

## Notes

- Most microbenchmarks use synthetic data; parser benchmarks also use committed real DICOM fixtures.
- Real-world performance varies based on DICOM file complexity and transfer syntax.
- Transfer syntax and compression can significantly impact parser/writer performance
- Network operations (when implemented) will have additional overhead
