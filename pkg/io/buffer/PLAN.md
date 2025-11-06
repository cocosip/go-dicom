# Buffer 实现规划

基于 fo-dicom 的 Buffer 实现，go-dicom 需要实现以下 Buffer 类型：

## 已完成的 Buffer (7/15) ✅

### 1. ✅ MemoryByteBuffer
**文件**: `memory.go`
**用途**: 内存中的字节缓冲区，最常用的类型
**参考**: `fo-dicom-code/IO/Buffer/MemoryByteBuffer.cs`

### 2. ✅ EmptyBuffer
**文件**: `empty.go`
**用途**: 空缓冲区（单例）
**参考**: `fo-dicom-code/IO/Buffer/EmptyBuffer.cs`

### 3. ✅ CompositeByteBuffer
**文件**: `composite.go`
**用途**: 组合多个 buffer，提供统一视图
**参考**: `fo-dicom-code/IO/Buffer/CompositeByteBuffer.cs`

### 4. ✅ StreamByteBuffer
**文件**: `stream.go`
**用途**: 基于流的 buffer，按需从 io.ReadSeeker 读取
**参考**: `fo-dicom-code/IO/Buffer/StreamByteBuffer.cs`

### 5. ✅ EndianByteBuffer
**文件**: `endian.go`
**用途**: 字节序转换包装器
**参考**: `fo-dicom-code/IO/Buffer/EndianByteBuffer.cs`

### 6. ✅ EvenLengthBuffer
**文件**: `evenlength.go`
**用途**: 确保偶数长度（DICOM 要求）
**参考**: `fo-dicom-code/IO/Buffer/EvenLengthBuffer.cs`

### 7. ✅ BulkDataUriByteBuffer
**文件**: `bulkdata.go`
**用途**: 表示 DICOM JSON 中的 BulkDataURI
**参考**: `fo-dicom-code/IO/Buffer/BulkDataUriByteBuffer.cs` 和 `IBulkDataUriByteBuffer.cs`
**实现日期**: 2025-11-07
**测试**: `bulkdata_test.go` (11 个测试用例，全部通过)

**说明**:
- 用于 DICOM JSON 序列化，引用外部数据 URI
- 支持延迟加载（lazy loading）模式
- 实现了完整的 ByteBuffer 接口
- 已集成到 JSON 序列化/反序列化中

---

## 待实现的 Buffer (8/15)

### 8. ⏳ RangeByteBuffer
**文件**: `range.go` (待创建)
**用途**: 表示另一个 buffer 的子范围（offset + length）
**参考**: `fo-dicom-code/IO/Buffer/RangeByteBuffer.cs`
**优先级**: 高
**预计工作量**: 0.5 天

**说明**: 用于高效地表示大 buffer 的一部分，无需复制数据

---

### 9. ⏳ FileByteBuffer
**文件**: `file.go` (待创建)
**用途**: 基于文件的 buffer
**参考**: `fo-dicom-code/IO/Buffer/FileByteBuffer.cs`
**优先级**: 中
**预计工作量**: 0.5 天

**说明**: 类似 StreamByteBuffer，但专门针对文件进行优化

---

### 10. ⏳ TempFileBuffer
**文件**: `tempfile.go` (待创建)
**用途**: 基于临时文件的 buffer，用于大数据
**参考**: `fo-dicom-code/IO/Buffer/TempFileBuffer.cs`
**优先级**: 中
**预计工作量**: 0.5-1 天

**说明**: 当数据太大无法放入内存时，使用临时文件存储

---

### 11. ⏳ SwapByteBuffer
**文件**: `swap.go` (待创建)
**用途**: 字节交换包装器（不同于 endian 转换）
**参考**: `fo-dicom-code/IO/Buffer/SwapByteBuffer.cs`
**优先级**: 低
**预计工作量**: 0.5 天

**说明**: 用于特定的 GE 私有传输语法

---

### 12. ⏳ LazyByteBuffer
**文件**: `lazy.go` (待创建)
**用途**: 延迟加载的 buffer
**参考**: `fo-dicom-code/IO/Buffer/LazyByteBuffer.cs`
**优先级**: 低
**预计工作量**: 0.5 天

**说明**: 使用函数延迟创建实际的 buffer

---

### 13. ⏳ ByteBufferByteSource
**文件**: `bytesource.go` (待创建)
**用途**: 将 ByteBuffer 适配为 ByteSource 接口
**参考**: `fo-dicom-code/IO/Buffer/ByteBufferByteSource.cs`
**优先级**: 中（文件读取时需要）
**预计工作量**: 0.5 天

**说明**: 桥接 Buffer 和 ByteSource 接口（第六阶段需要）

---

### 14. ByteBufferEnumerator
**文件**: N/A (Go 不需要)
**参考**: `fo-dicom-code/IO/Buffer/ByteBufferEnumerator.cs`
**优先级**: N/A
**说明**: C# 的迭代器模式，Go 中使用不同方式实现

---

### 15. ByteBufferExtensions
**文件**: N/A (Go 中使用普通函数)
**参考**: `fo-dicom-code/IO/Buffer/ByteBufferExtensions.cs`
**优先级**: N/A
**说明**: C# 扩展方法，Go 中作为包级函数实现

---

## 实现优先级建议

### 第一优先级（立即实现）
这些是文件读写的基础：
- ✅ MemoryByteBuffer
- ✅ EmptyBuffer
- ✅ CompositeByteBuffer
- ✅ StreamByteBuffer
- ✅ EndianByteBuffer
- ✅ EvenLengthBuffer
- ✅ BulkDataUriByteBuffer
- ⏳ RangeByteBuffer

### 第二优先级（短期实现）
这些在文件 I/O 阶段需要：
- ⏳ FileByteBuffer
- ⏳ ByteBufferByteSource
- ⏳ TempFileBuffer

### 第三优先级（中期实现）
这些用于高级功能：
- ⏳ LazyByteBuffer
- ⏳ SwapByteBuffer

---

## 当前状态

**已完成**: 7/15 (47%)
**核心功能**: ✅ 完成

**最新完成**: BulkDataUriByteBuffer (2025-11-07)
- 实现了完整的 BulkDataUriByteBuffer 及测试
- 已集成到 JSON 序列化/反序列化中
- 支持所有相关 VR 类型（OB, OW, OD, OF, OL, OV, UN）
- 11 个单元测试 + 7 个集成测试，全部通过

**下一步**: 继续 JSON 序列化的其他功能

**里程碑**:
- Buffer 抽象层核心完成 ✅
- JSON 序列化基础完成 ✅
- BulkDataURI 支持完成 ✅
- DicomElement/DicomDataset 已实现 ✅
