# DICOM Transcoder 设计文档

## 概述

DicomTranscoder 是 DICOM 图像编解码的核心组件，负责在不同 Transfer Syntax 之间转换像素数据。

## 功能需求

### 核心功能

1. **转码（Transcode）**
   - 压缩 → 压缩：解压缩 → 重新压缩
   - 压缩 → 未压缩：解压缩
   - 未压缩 → 压缩：压缩
   - 未压缩 → 未压缩：像素格式转换（字节序、平面配置等）

2. **单帧解码（DecodeFrame）**
   - 从压缩数据中提取并解压单个帧
   - 返回未压缩的帧数据

3. **像素数据解码（DecodePixelData）**
   - 解码单帧并返回可渲染的像素数据
   - 处理颜色空间转换

### 辅助功能

1. **Overlay 处理**
   - 提取嵌入在像素数据中的 overlay
   - 转换为独立的 overlay 数据元素

2. **有损压缩元数据**
   - 记录有损压缩方法
   - 计算和记录压缩比

## 架构设计

### 组件关系

```
DicomTranscoder
├─ InputCodec (IDicomCodec)
│  └─ Decode(PixelData) -> UncompressedPixelData
├─ OutputCodec (IDicomCodec)
│  └─ Encode(PixelData) -> CompressedPixelData
└─ TranscoderManager
   └─ GetCodec(TransferSyntax) -> IDicomCodec
```

### 接口定义

```go
// Transcoder interface
type Transcoder interface {
    // Transcode dataset from InputSyntax to OutputSyntax
    Transcode(ds *dataset.Dataset) (*dataset.Dataset, error)

    // DecodeFrame extracts and decodes a single frame
    DecodeFrame(ds *dataset.Dataset, frameIndex int) ([]byte, error)

    // DecodePixelData decodes pixel data for rendering
    DecodePixelData(ds *dataset.Dataset, frameIndex int) (*PixelData, error)

    // Properties
    InputSyntax() *transfer.TransferSyntax
    OutputSyntax() *transfer.TransferSyntax
}
```

## 转码场景

### 1. 压缩 → 未压缩

```
JPEG Baseline → Explicit VR Little Endian

Steps:
1. Extract compressed pixel data (FragmentSequence)
2. For each fragment:
   - Decode using JPEG codec
   - Store in uncompressed pixel data
3. Update Transfer Syntax UID
4. Update Photometric Interpretation (if needed)
```

### 2. 未压缩 → 压缩

```
Explicit VR Little Endian → JPEG 2000

Steps:
1. Extract uncompressed pixel data
2. For each frame:
   - Encode using JPEG 2000 codec
   - Add to FragmentSequence
3. Build offset table
4. Update Transfer Syntax UID
5. Add lossy compression metadata
```

### 3. 压缩 → 压缩

```
JPEG Baseline → JPEG 2000

Steps:
1. Decode to uncompressed (temp)
2. Encode from uncompressed to target format
```

### 4. 未压缩 → 未压缩

```
Implicit VR Little Endian → Explicit VR Big Endian

Steps:
1. Extract pixel data
2. Apply byte order conversion
3. Update Transfer Syntax UID
```

## Codec Registry

### 内置 Codecs

- ✅ **Native**: 未压缩格式
- ✅ **RLE**: Run-Length Encoding
- ⏳ **JPEG Baseline**: 需要实现
- ⏳ **JPEG Lossless**: 需要实现
- ⏳ **JPEG 2000**: 需要实现
- ⏳ **JPEG-LS**: 需要实现

### TranscoderManager

```go
type TranscoderManager interface {
    // Register a codec for a transfer syntax
    RegisterCodec(ts *transfer.TransferSyntax, codec Codec)

    // Get codec for a transfer syntax
    GetCodec(ts *transfer.TransferSyntax) (Codec, bool)

    // Check if codec exists
    HasCodec(ts *transfer.TransferSyntax) bool

    // Create transcoder
    CreateTranscoder(inputTS, outputTS *transfer.TransferSyntax) (*Transcoder, error)
}
```

## 实现策略

### Phase 1: 基础框架（当前）
- ✅ Codec 接口
- ✅ PixelData 类型
- ✅ Native codec
- ✅ RLE codec
- ⏳ Transcoder 接口
- ⏳ TranscoderManager

### Phase 2: JPEG 支持
- JPEG Baseline codec
- JPEG Lossless codec
- 使用 Go image/jpeg 或外部库

### Phase 3: JPEG 2000 支持
- JPEG 2000 codec
- 需要外部库（如 openjpeg binding）

### Phase 4: 高级功能
- Overlay 提取和处理
- 多线程并行编解码
- 流式处理大文件

## 外部依赖

### JPEG Support
- **Option 1**: Go 标准库 `image/jpeg`
  - ✅ 无外部依赖
  - ❌ 仅支持 Baseline (8-bit)
  - ❌ 不支持 12-bit 和 Lossless

- **Option 2**: libjpeg-turbo (CGO)
  - ✅ 完整 JPEG 支持
  - ✅ 高性能
  - ❌ 需要 CGO
  - ❌ 需要安装 libjpeg

### JPEG 2000 Support
- **Option 1**: openjpeg (CGO)
  - ✅ 完整 JPEG 2000 支持
  - ❌ 需要 CGO
  - ❌ 需要安装 openjpeg

- **Option 2**: Pure Go 实现
  - ✅ 无外部依赖
  - ❌ 性能较低
  - ❌ 需要大量开发工作

## 建议的实现顺序

1. **立即实现**：
   - Transcoder 接口和基础实现
   - TranscoderManager
   - 未压缩 ↔ 未压缩转码
   - RLE ↔ 未压缩转码

2. **短期实现**：
   - JPEG Baseline codec（使用标准库）
   - 基本的 JPEG ↔ 未压缩转码

3. **中期实现**：
   - 完整 JPEG 支持（考虑 CGO）
   - JPEG-LS codec

4. **长期实现**：
   - JPEG 2000 support
   - 性能优化
   - Overlay 处理

## 使用示例

### 基本转码

```go
// Create transcoder
transcoder := NewTranscoder(
    transfer.JPEGBaseline,
    transfer.ExplicitVRLittleEndian,
    nil, // input params
    nil, // output params
)

// Transcode dataset
outputDS, err := transcoder.Transcode(inputDS)
```

### 单帧解码

```go
transcoder := NewTranscoder(
    transfer.JPEG2000Lossless,
    transfer.ExplicitVRLittleEndian,
    nil, nil,
)

// Decode specific frame
frameData, err := transcoder.DecodeFrame(ds, 5) // Frame #5
```

### 使用 TranscoderManager

```go
// Register codecs
manager := NewTranscoderManager()
manager.RegisterCodec(transfer.JPEGBaseline, NewJPEGCodec())
manager.RegisterCodec(transfer.RLE, NewRLECodec())

// Create transcoder
transcoder, err := manager.CreateTranscoder(
    transfer.JPEGBaseline,
    transfer.ExplicitVRLittleEndian,
)

// Transcode
outputDS, err := transcoder.Transcode(inputDS)
```

## 参考

- fo-dicom: `DicomTranscoder.cs`
- DICOM Standard Part 5: Data Structures and Encoding
- DICOM Standard Part 6: Data Dictionary (Transfer Syntaxes)
