# Fragment Sequence Implementation Summary

## 概述

成功实现了 DICOM Fragment Sequence 功能，用于支持封装（encapsulated）像素数据，特别是多帧压缩图像（JPEG, JPEG 2000, JPEG-LS, RLE 等）。

## 实现的功能

### 1. FragmentSequence 类型 (`fragment.go`)

- ✅ `FragmentSequence` 结构体：存储 offset table 和 fragments
- ✅ `OtherByteFragment` (OB VR)
- ✅ `OtherWordFragment` (OW VR)
- ✅ Offset Table 管理（`SetOffsetTable`, `OffsetTable`）
- ✅ Fragment 管理（`AddFragment`, `GetFragment`, `FragmentCount`）
- ✅ 实现了 `Element` 接口

**关键方法：**
```go
type FragmentSequence struct {
    tag         *tag.Tag
    vr          *vr.VR
    offsetTable []uint32            // 每帧的字节偏移量
    fragments   []buffer.ByteBuffer // 压缩帧数据
}
```

### 2. Parser 支持 (`parser.go`)

- ✅ 检测 Fragment Sequence（OB/OW + undefined length 0xFFFFFFFF）
- ✅ `readFragmentSequence()` 方法
- ✅ 正确处理 Offset Table（第一个 Item）
- ✅ 读取所有 Fragments
- ✅ 处理 Sequence Delimitation Item (FFFE,E0DD)
- ✅ 处理 EOF 边界情况（某些文件可能缺少 delimitation）

**关键代码：**
```go
// 检测
if (vrValue.Code() == vr.CodeOB || vrValue.Code() == vr.CodeOW) && length == 0xFFFFFFFF {
    return p.readFragmentSequence(t, vrValue)
}
```

### 3. Writer 支持 (`writer.go`)

- ✅ 检测 Fragment Sequence 类型
- ✅ `writeFragmentSequence()` 方法
- ✅ 写入正确的 Tag + VR + Undefined Length
- ✅ 写入 Offset Table Item
- ✅ 写入所有 Fragment Items
- ✅ 写入 Sequence Delimitation Item

**DICOM Fragment Sequence 写入格式：**
```
PixelData (7FE0,0010) + OB/OW + 00 00 + FFFFFFFF
├─ Item (FFFE,E000) + Length                # Offset Table
│  └─ [offset1, offset2, ...] (uint32 数组)
├─ Item (FFFE,E000) + Length                # Fragment 1
│  └─ [compressed data]
├─ Item (FFFE,E000) + Length                # Fragment 2
│  └─ [compressed data]
├─ ...
└─ Delimitation (FFFE,E0DD) + 00000000
```

### 4. 测试覆盖

#### Element 测试 (`fragment_test.go`)
- ✅ FragmentSequence 创建和基本操作
- ✅ Offset Table 操作
- ✅ Fragment 添加和获取
- ✅ OtherByteFragment 和 OtherWordFragment

#### Parser 测试 (`parser_fragment_test.go`)
- ✅ 解析空 Offset Table 的 Fragment Sequence
- ✅ 解析带 Offset Table 的 Fragment Sequence
- ✅ 验证 Fragment 数量和大小

#### Writer 测试 (`writer_fragment_test.go`)
- ✅ 写入空 Offset Table 的 Fragment Sequence
- ✅ 写入带 Offset Table 的 Fragment Sequence
- ✅ **Round-trip 测试：Write → Read → 验证数据完整性**

**所有测试均通过！**

## 技术细节

### Offset Table 说明

Offset Table 是 Fragment Sequence 的第一个 Item，包含每帧相对于第一个 fragment 的字节偏移量：

- **单帧图像**：Offset Table 可以为空（Length = 0）
- **多帧图像**：包含 N 个 uint32 offsets（N = 帧数）
- **格式**：Little Endian uint32 数组

示例（2 帧图像）：
```
Item (FFFE,E000) + 00000008  # Length = 8 bytes (2 offsets)
├─ 00000000                  # Frame 1 at offset 0
└─ 00000010                  # Frame 2 at offset 16
```

### 支持的 Transfer Syntaxes

Fragment Sequence 用于以下压缩格式：

- **JPEG Baseline** (1.2.840.10008.1.2.4.50)
- **JPEG Extended** (1.2.840.10008.1.2.4.51)
- **JPEG Lossless** (1.2.840.10008.1.2.4.57, 1.2.840.10008.1.2.4.70)
- **JPEG 2000** (1.2.840.10008.1.2.4.90, 1.2.840.10008.1.2.4.91)
- **JPEG-LS** (1.2.840.10008.1.2.4.80, 1.2.840.10008.1.2.4.81)
- **RLE** (1.2.840.10008.1.2.5)

### 性能考虑

1. **内存管理**：Fragments 使用 `buffer.ByteBuffer` 接口，支持内存池优化
2. **懒加载**：可以与 Parser 的 `ReadLargeOnDemand` 选项结合
3. **流式处理**：逐 Fragment 读写，避免一次性加载所有数据

## 使用示例

### 创建 Fragment Sequence

```go
// 创建 Fragment Sequence
obf := element.NewOtherByteFragment(tag.PixelData)

// 设置 Offset Table（可选）
obf.SetOffsetTable([]uint32{0, 1024, 2048})

// 添加 Fragments
frame1 := compressFrame(imageData1) // JPEG 压缩
frame2 := compressFrame(imageData2)
obf.AddFragment(buffer.NewMemory(frame1))
obf.AddFragment(buffer.NewMemory(frame2))

// 添加到 Dataset
ds.Add(obf)
```

### 读取 Fragment Sequence

```go
// 解析 DICOM 文件
result, _ := parser.Parse(file)

// 获取 PixelData
pixelData, exists := result.Dataset.Get(tag.PixelData)
if !exists {
    return
}

// 类型断言
if obf, ok := pixelData.(*element.OtherByteFragment); ok {
    // 获取 Offset Table
    offsets := obf.OffsetTable()

    // 遍历 Fragments
    for i := 0; i < obf.FragmentCount(); i++ {
        frag, _ := obf.GetFragment(i)
        compressedData := frag.Data()

        // 解压缩
        imageData := decompressFrame(compressedData)
        // ... 处理图像数据
    }
}
```

### Round-trip（读取→修改→写入）

```go
// 读取
result, _ := parser.Parse(inputFile)

// 修改或添加 Fragments
pixelData, _ := result.Dataset.Get(tag.PixelData)
obf := pixelData.(*element.OtherByteFragment)
obf.AddFragment(buffer.NewMemory(newFrameData))

// 写入
outputFile, _ := os.Create("output.dcm")
writer.Write(outputFile, result.Dataset)
```

## 与 fo-dicom 的兼容性

实现参考了 fo-dicom 的设计：

- ✅ 相同的 Fragment Sequence 结构
- ✅ 相同的 Offset Table 格式
- ✅ 兼容的文件格式
- ✅ 支持所有标准压缩格式

## 已知限制

1. **压缩/解压缩**：当前实现只处理 Fragment Sequence 的读写，不包含实际的图像压缩/解压缩算法
   - JPEG/JPEG2000/RLE 编解码器需要单独实现或使用外部库

2. **大文件优化**：虽然支持懒加载，但 Fragment Sequence 本身目前会一次性读入内存
   - 未来可以实现流式 Fragment 访问

## 后续工作

1. **图像编解码器集成**：
   - 实现或集成 JPEG, JPEG 2000, RLE 编解码器
   - 提供 `DecompressFrame()` 和 `CompressFrame()` 辅助函数

2. **高级 API**：
   - 提供 `DicomImage` 类型，自动处理 Fragment 解压缩
   - 支持多帧图像的逐帧访问

3. **性能优化**：
   - Fragment 的流式读写
   - 并行解压缩多个 Fragments

## 测试结果

```bash
# Element 测试
PASS: TestFragmentSequence

# Parser 测试
PASS: TestParseFragmentSequence/EmptyOffsetTable
PASS: TestParseFragmentSequence/WithOffsetTable

# Writer 测试
PASS: TestWriteFragmentSequence/EmptyOffsetTable
PASS: TestWriteFragmentSequence/WithOffsetTable
PASS: TestFragmentSequenceRoundTrip/RoundTrip

✅ 所有测试通过！
```

## 文件清单

- `pkg/dicom/element/fragment.go` - FragmentSequence 类型定义
- `pkg/dicom/element/fragment_test.go` - 单元测试
- `pkg/dicom/element/FRAGMENT_DESIGN.md` - 设计文档
- `pkg/dicom/element/FRAGMENT_SUMMARY.md` - 本文档
- `pkg/dicom/parser/parser.go` - Parser 实现（添加了 readFragmentSequence）
- `pkg/dicom/parser/parser_fragment_test.go` - Parser 测试
- `pkg/dicom/writer/writer.go` - Writer 实现（添加了 writeFragmentSequence）
- `pkg/dicom/writer/writer_fragment_test.go` - Writer 和 round-trip 测试

## 结论

✅ **Fragment Sequence 功能已完整实现并通过测试**

该实现为 go-dicom 提供了完整的多帧压缩图像支持基础，符合 DICOM 标准，并与 fo-dicom 兼容。
