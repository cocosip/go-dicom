# Buffer 实现规划

基于 fo-dicom 的 Buffer 实现，go-dicom 需要实现以下 Buffer 类型：

## 已完成的 Buffer (12/15) ✅

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

### 8. ✅ RangeByteBuffer
**文件**: `range.go`
**用途**: 表示另一个 buffer 的子范围（offset + length）
**参考**: `fo-dicom-code/IO/Buffer/RangeByteBuffer.cs`
**实现日期**: 2025-11-07
**测试**: `range_test.go` (12 个测试用例，全部通过)

**说明**:
- 用于高效地表示大 buffer 的一部分，无需复制数据
- 支持嵌套范围（一个范围的范围）
- 在 DICOM 序列和像素数据帧处理中非常有用
- 实现了分块写入以优化大数据处理

### 9. ✅ FileByteBuffer
**文件**: `file.go`
**用途**: 基于文件的 buffer
**参考**: `fo-dicom-code/IO/Buffer/FileByteBuffer.cs`
**实现日期**: 2025-11-07
**测试**: `file_test.go` (13 个测试用例，全部通过)

**说明**:
- 提供文件特定范围的高效访问，无需加载整个文件到内存
- 支持分块读写（1MB chunks），适合大文件处理
- 实现了 FilePath() 和 Position() 方法用于文件信息查询
- 在读取大型 DICOM 文件时非常有用

---

### 10. ✅ TempFileBuffer
**文件**: `tempfile.go`
**用途**: 基于临时文件的 buffer，用于大数据
**参考**: `fo-dicom-code/IO/Buffer/TempFileBuffer.cs`
**实现日期**: 2025-11-07
**测试**: `tempfile_test.go` (14 个测试用例，全部通过)

**说明**:
- 当数据太大无法放入内存时，使用临时文件存储
- 自动创建和管理临时文件
- 提供 Close() 方法清理临时文件
- 支持多次读取操作

### 11. ✅ SwapByteBuffer
**文件**: `swap.go`
**用途**: 字节交换包装器（不同于 endian 转换）
**参考**: `fo-dicom-code/IO/Buffer/SwapByteBuffer.cs`
**实现日期**: 2025-11-07
**测试**: `swap_test.go` (15 个测试用例，全部通过)

**说明**:
- 用于特定的 GE 私有传输语法
- 与 EndianByteBuffer 不同：SwapByteBuffer 总是交换字节，不考虑机器字节序
- 支持不同的 unitSize (2, 4, 8 等)
- 使用 endian.SwapBytes 进行高效的字节交换

### 12. ✅ LazyByteBuffer
**文件**: `lazy.go`
**用途**: 延迟加载的 buffer
**参考**: `fo-dicom-code/IO/Buffer/LazyByteBuffer.cs`
**实现日期**: 2025-11-07
**测试**: `lazy_test.go` (14 个测试用例，全部通过)

**说明**:
- 使用函数延迟创建实际的 buffer
- 数据只在首次访问时加载，然后缓存
- 使用 sync.Once 确保线程安全且只加载一次
- 提供 IsLoaded() 方法检查是否已加载
- 适用于延迟昂贵的数据加载操作

---

## 待实现的 Buffer (3/15)

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
- ✅ RangeByteBuffer

### 第二优先级（短期实现）
这些在文件 I/O 阶段需要：
- ✅ FileByteBuffer
- ⏳ ByteBufferByteSource (需要先实现 ByteSource 接口层)
- ✅ TempFileBuffer

### 第三优先级（中期实现）
这些用于高级功能：
- ✅ LazyByteBuffer
- ✅ SwapByteBuffer

---

## 当前状态

**已完成**: 12/15 (80%)
**核心功能**: ✅ 完成

**最新完成**: LazyByteBuffer (2025-11-07)
- 实现了完整的 LazyByteBuffer 及测试
- 使用函数延迟创建实际的 buffer，数据只在首次访问时加载
- 使用 sync.Once 确保线程安全且只加载一次
- 14 个单元测试，全部通过

**本次会话完成的 Buffer**:
1. RangeByteBuffer - 12 个测试，支持嵌套范围和分块写入
2. FileByteBuffer - 13 个测试，文件范围高效访问
3. TempFileBuffer - 14 个测试，临时文件存储大数据
4. SwapByteBuffer - 15 个测试，字节交换用于 GE 私有语法
5. LazyByteBuffer - 14 个测试，延迟加载数据

**下一步**: 实现 ByteSource 接口层，然后完成 ByteBufferByteSource

**里程碑**:
- Buffer 抽象层核心完成 ✅
- JSON 序列化基础完成 ✅
- BulkDataURI 支持完成 ✅
- DicomElement/DicomDataset 已实现 ✅
