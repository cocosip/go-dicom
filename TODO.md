# go-dicom 开发任务列表

本文档记录从 fo-dicom (C#) 迁移到 go-dicom 的开发任务清单。

## 图例
- [ ] 待完成
- [x] 已完成
- [~] 进行中
- [-] 已取消/暂缓

---

## 第一阶段：项目基础设施 ✅

### 1.1 项目初始化
- [x] 创建 Go module (go.mod)
- [x] 设置项目目录结构
  - [x] pkg/dicom (核心 DICOM 类型)
  - [x] pkg/io (IO 操作)
  - [x] pkg/transfer (传输语法)
  - [x] internal (内部工具)
  - [x] cmd (命令行工具)
  - [x] examples (示例代码)
- [x] 配置 golangci-lint
- [ ] 设置 GitHub Actions CI (可选)
- [x] 添加 README.md
- [x] 添加 LICENSE 文件 (MS-PL)

### 1.2 开发工具配置
- [x] 创建 Makefile 用于常用命令
- [x] 设置 .gitignore
- [ ] 配置 VSCode/GoLand 调试环境 (可选)

---

## 第二阶段：核心数据类型

### 2.1 DicomVR (Value Representation) ✅
**参考**: `fo-dicom-code/DicomVR.cs` (~878 lines)
**包**: `pkg/dicom/vr`

- [x] 定义 VR 结构体
- [x] 实现所有标准 VR 常量 (35个)
  - [x] 字符串类型: AE, AS, CS, DA, DS, DT, IS, LO, LT, PN, SH, ST, TM, UC, UI, UR, UT
  - [x] 二进制类型: OB, OD, OF, OL, OV, OW, UN
  - [x] 数值类型: FL, FD, SL, SS, UL, US, SV, UV
  - [x] 特殊类型: AT, SQ
- [x] 实现 Parse() 和 TryParse()
- [x] 实现 ParseBytes() (优化的字节解析)
- [x] 实现 VR 验证逻辑 (15+ 验证函数)
- [x] 编写 VR 单元测试 (100+ 测试用例)
- [x] 全局验证开关 PerformValidation

**依赖**: 无
**实际工作量**: 3 天
**状态**: ✅ 完成，所有测试通过

### 2.1.5 DicomVM (Value Multiplicity) ✅
**参考**: `fo-dicom-code/DicomVM.cs` (~127 lines)
**包**: `pkg/dicom/vm`

- [x] 定义 VM 结构体
- [x] 实现 VM 解析 (支持 "1", "1-2", "1-n", "2-2n" 等格式)
- [x] 实现 IsValid() 值数量验证
- [x] 实现所有标准 VM 常量 (15个)
- [x] 实现解析缓存 (线程安全)
- [x] 编写单元测试 (50+ 测试用例)

**依赖**: 无
**实际工作量**: 1 天
**状态**: ✅ 完成，所有测试通过

### 2.2 DicomTag ✅
**参考**: `fo-dicom-code/DicomTag.cs` (~272 lines), `DicomMaskedTag.cs`, `DicomTagGenerated.cs`
**包**: `pkg/dicom/tag`

- [x] 定义 Tag 结构体
- [x] 实现 Tag 构造函数 (New, NewWithPrivateCreator, FromUint32)
- [x] 实现 Tag 格式化输出 (String方法)
  - [x] 格式 "G": (0028,0010)
  - [x] 格式 "X": (0028,xx10) 用于私有标签
  - [x] 格式 "g": 0028,0010 (无括号)
  - [x] 格式 "J": 00280010 (紧凑格式)
- [x] 实现 Tag 比较 (Equals, Compare)
- [x] 实现 Tag 转换 (uint32 <-> Tag)
- [x] 实现 IsPrivate() 方法
- [x] 实现 Parse() 和 MustParse()
- [x] 实现 Hash() 用于 map
- [x] 实现 PrivateCreator 类型
- [x] 实现 MaskedTag 类型 (通配符匹配)
- [x] 生成 5338 个标准 Tag 常量
- [x] 编写 Tag 单元测试 (30+ 测试用例)

**依赖**: 无
**实际工作量**: 2 天
**状态**: ✅ 完成，所有测试通过

### 2.3 DicomDictionary ✅
**参考**: `fo-dicom-code/DicomDictionary.cs` (~363 lines), `DicomDictionaryEntry.cs`
**包**: `pkg/dicom/dict`

- [x] 定义 Entry 结构体 (Tag, VR, VM, Name, Keyword, IsRetired)
- [x] 实现 NewEntry() 和 NewEntryWithMask() 构造函数
- [x] 实现 Dictionary 基本结构 (使用 map)
- [x] 实现 Lookup(tag) 按标签查询
- [x] 实现 LookupKeyword(keyword) 按关键字查询
- [x] 实现 MaskedTag 匹配支持 (如 group length)
- [x] 实现私有字典支持 (NewPrivate)
- [x] 线程安全实现 (RWMutex)
- [x] 编写单元测试 (7 个测试用例)

**依赖**: DicomTag, DicomVR, DicomVM
**实际工作量**: 1 天
**状态**: ✅ 完成，所有测试通过

### 2.4 DicomUID ✅
**参考**: `fo-dicom-code/DicomUID.cs`, `DicomUIDGenerated.cs`, `DicomUIDPrivate.cs`
**包**: `pkg/dicom/uid`

- [x] 定义 UID 结构体
- [x] 定义 UID 类型枚举 (TransferSyntax, SOPClass, SOPInstance 等 - 13种类型)
- [x] 定义存储类别枚举 (Image, PresentationState, StructuredReport 等 - 10种类别)
- [x] 实现 UID 验证 (IsValid)
- [x] 实现 UID 生成 (GenerateFromRoot, Append, GenerateDerivedFromUUID)
- [x] 实现 UID 生成器 (Generator with source-to-destination mapping)
- [x] 实现 Parse() 和 Register()
- [x] 生成 1906 个标准 UID 常量 (从 DicomUIDGenerated.cs)
- [x] 生成 59 个私有 UID 常量 (从 DicomUIDPrivate.cs)
- [x] 实现 StorageCategory 检测 (IsImageStorage, IsVolumeStorage)
- [x] 编写单元测试 (22个测试函数)

**依赖**: 无
**实际工作量**: 2 天
**状态**: ✅ 完成，所有测试通过

### 2.5 DicomTransferSyntax ✅
**参考**: `fo-dicom-code/DicomTransferSyntax.cs`
**包**: `pkg/dicom/transfer`

- [x] 定义 TransferSyntax 结构体 (使用 Builder 模式)
- [x] 实现传输语法属性
  - [x] IsExplicitVR (显式/隐式 VR)
  - [x] IsEncapsulated (封装格式)
  - [x] IsLossy (有损压缩)
  - [x] IsDeflate (Deflate 压缩)
  - [x] LossyCompressionMethod (有损压缩方法)
  - [x] Endian (字节序)
  - [x] SwapPixelData (像素数据交换)
- [x] 定义标准传输语法常量 (15+ 个)
  - [x] ImplicitVRLittleEndian
  - [x] ExplicitVRLittleEndian
  - [x] ExplicitVRBigEndian (退役)
  - [x] DeflatedExplicitVRLittleEndian
  - [x] JPEG 系列 (Baseline, Extended, Lossless, JPEG-LS)
  - [x] JPEG 2000 (Lossless, Lossy)
  - [x] RLE Lossless
  - [x] GE Private Transfer Syntaxes
- [x] 实现 Parse() 和 Lookup() 方法
- [x] 实现 Register/Unregister/Query 方法
- [x] 实现辅助函数 (IsUncompressed, IsLosslessCompressed, IsLossyCompressed)
- [x] 编写单元测试 (17个测试函数)

**依赖**: DicomUID, Endian
**实际工作量**: 1 天
**状态**: ✅ 完成，所有测试通过

### 2.6 DicomEncoding (Charset) ✅
**参考**: `fo-dicom-code/DicomEncoding.cs`
**包**: `pkg/dicom/charset`

- [x] 定义字符集类型 (CharsetInfo)
- [x] 实现 DICOM 特定字符集映射 (30+ 字符集)
- [x] 实现字符编码转换
  - [x] ISO 8859 系列 (Latin-1 to Latin-5, Greek, Arabic, Hebrew, Cyrillic, Thai)
  - [x] UTF-8
  - [x] 中文 (GBK, GB18030, GB2312, HZ-GB2312)
  - [x] 日文 (Shift-JIS, ISO-2022-JP)
  - [x] 韩文 (EUC-KR)
  - [x] ISO 2022 扩展字符集系列
- [x] 实现 GetEncoding() 方法 (支持拼写容错)
- [x] 实现 GetEncodings() 方法 (多字符集支持)
- [x] 实现 EncodeString/DecodeString 方法
- [x] 实现 GetCharsetName() 反向查询
- [x] 实现 KnownCharsets() 和 GetCharsetInfo()
- [x] 编写单元测试 (11个测试函数，包含往返测试)

**依赖**: golang.org/x/text/encoding
**实际工作量**: 1 天
**状态**: ✅ 完成，所有测试通过

### 2.7 辅助类型
**参考**: `fo-dicom-code/DicomDateRange.cs`, `DicomParseable.cs`
**包**: `pkg/dicom`

- [ ] DicomDateRange
  - [ ] 日期范围解析
  - [ ] DICOM 查询格式支持 ("20240101-20241231")
  - [ ] 单元测试
- [ ] DicomParseable 接口
  - [ ] 定义 Parse() 接口规范
  - [ ] 文档说明

**依赖**: 无
**预计工作量**: 0.5-1 天

---

## 第三阶段：Buffer 抽象层 ✅

### 3.1 IByteBuffer 接口 ✅
**参考**: `fo-dicom-code/IO/Buffer/IByteBuffer.cs`
**包**: `pkg/io/buffer`

- [x] 定义 ByteBuffer 接口
  ```go
  type ByteBuffer interface {
      IsMemory() bool
      Size() uint32
      Data() []byte
      GetByteRange(offset, count uint32, output []byte) error
      WriteTo(w io.Writer) (int64, error)
  }
  ```
- [x] 实现 MemoryByteBuffer (基于 []byte)
- [x] 实现 EmptyBuffer (单例模式)
- [x] 编写单元测试 (8个测试函数)

**依赖**: 无
**实际工作量**: 0.5 天
**状态**: ✅ 完成，所有测试通过

### 3.2 核心 Buffer 实现 ✅
**参考**: `fo-dicom-code/IO/Buffer/*.cs`
**包**: `pkg/io/buffer`

- [x] 实现 CompositeByteBuffer (组合多个 buffer)
- [x] 实现 StreamByteBuffer (基于 io.ReadSeeker)
- [x] 实现 EndianByteBuffer (字节序转换)
- [x] 实现 EvenLengthBuffer (确保偶数长度)
- [x] 实现 RangeByteBuffer (子范围视图)
- [x] 编写单元测试 (14个测试函数，包含并发测试)

**依赖**: IByteBuffer, pkg/dicom/endian
**实际工作量**: 1 天
**状态**: ✅ 完成，所有测试通过

**已实现的 Buffer 类型** (7/15):
1. ✅ MemoryByteBuffer - 内存缓冲区
2. ✅ EmptyBuffer - 空缓冲区
3. ✅ CompositeByteBuffer - 组合缓冲区
4. ✅ StreamByteBuffer - 流缓冲区
5. ✅ EndianByteBuffer - 字节序转换
6. ✅ EvenLengthBuffer - 偶数长度包装
7. ✅ RangeByteBuffer - 子范围视图

**待实现的 Buffer 类型** (按需实现):
- FileByteBuffer - 文件缓冲区
- TempFileBuffer - 临时文件缓冲区
- LazyByteBuffer - 延迟加载
- SwapByteBuffer - 字节交换
- BulkDataUriByteBuffer - DICOM JSON 支持
- 其他辅助类型

### 3.3 Endian 字节序处理 ✅
**参考**: `fo-dicom-code/IO/Endian.cs`
**包**: `pkg/dicom/endian`

- [x] 定义 Endian 类型 (Little, Big)
- [x] 实现 SwapBytes 系列函数
  - [x] SwapUint16/SwapInt16
  - [x] SwapUint32/SwapInt32
  - [x] SwapUint64/SwapInt64
- [x] 实现批量字节序转换 (SwapBytes, SwapBytesN)
- [x] 实现字节序检测 (LocalMachine, Network)
- [x] 集成 encoding/binary
- [x] 编写单元测试 (18个测试函数)

**依赖**: 无
**实际工作量**: 0.5 天
**状态**: ✅ 完成，所有测试通过

---

## 第四阶段：DICOM 元素和数据集 ✅

### 4.1 DicomElement (基础) ✅
**参考**: `fo-dicom-code/DicomElement.cs` (~2049 lines)
**包**: `pkg/dicom/element`

- [x] 定义 Element 接口
- [x] 实现基础 Element 结构
- [x] 实现字符串类型元素工厂函数
  - [x] NewString() - 通用字符串元素
  - [x] 支持多值字符串
- [x] 实现数值类型元素
  - [x] NewUnsignedShort (US)
  - [x] NewUnsignedLong (UL)
  - [x] NewSignedShort (SS)
  - [x] NewSignedLong (SL)
  - [x] NewFloat (FL)
  - [x] NewDouble (FD)
- [x] 实现二进制元素
  - [x] NewOtherByte (OB)
  - [x] NewOtherWord (OW)
- [x] 编写单元测试 (13个测试函数)

**依赖**: DicomTag, DicomVR, ByteBuffer
**实际工作量**: 2 天
**状态**: ✅ 完成，所有测试通过

### 4.2 DicomSequence ✅
**包**: `pkg/dicom/dataset`

- [x] 实现 Sequence 类型
- [x] 实现 NewSequence() 构造函数
- [x] 实现 AddItem/GetItem/Count 方法
- [x] 实现 Element 接口
- [x] 编写单元测试

**依赖**: DicomElement
**实际工作量**: 0.5 天
**状态**: ✅ 完成，所有测试通过

### 4.3 DicomDataset ✅
**参考**: `fo-dicom-code/DicomDataset.cs` (~1691 lines)
**包**: `pkg/dicom/dataset`

- [x] 定义 Dataset 结构体 (使用切片保持插入顺序)
- [x] 实现 Add/Remove/Get 方法
- [x] 实现类型安全的 GetString/GetUInt16/GetUInt32/GetFloat32/GetFloat64 方法
- [x] 实现 Elements() 遍历方法
- [x] 实现 Clear() 方法
- [x] 编写单元测试 (12个测试函数)

**依赖**: DicomElement
**实际工作量**: 1 天
**状态**: ✅ 完成，所有测试通过

---

## 第五阶段：传输语法 (Transfer Syntax)

### 5.1 DicomTransferSyntax
**参考**: `fo-dicom-code/DicomTransferSyntax.cs` (~920 lines)

- [ ] 定义 TransferSyntax 结构体
  ```go
  type TransferSyntax struct {
      UID              UID
      IsExplicitVR     bool
      IsLittleEndian   bool
      IsDeflated       bool
      IsEncapsulated   bool
      IsLossy          bool
      // ...
  }
  ```
- [ ] 定义常用传输语法常量
  - [ ] ExplicitVRLittleEndian (默认)
  - [ ] ImplicitVRLittleEndian
  - [ ] ExplicitVRBigEndian
  - [ ] JPEGBaseline
  - [ ] JPEG2000Lossless
  - [ ] RLELossless
  - [ ] 等等...
- [ ] 实现 TransferSyntax.Lookup(uid) 查找
- [ ] 编写单元测试

**依赖**: DicomUID
**预计工作量**: 2-3 天

### 5.2 DicomUID
**参考**: `fo-dicom-code/DicomUID.cs`

- [ ] 定义 UID 结构体
- [ ] 实现常用 UID 常量 (可先实现部分)
- [ ] 实现 UID 验证
- [ ] 实现 UID 生成器
- [ ] 编写单元测试

**依赖**: 无
**预计工作量**: 1-2 天

---

## 第六阶段：DICOM 文件读写 ✅

### 6.1 DicomParser (DICOM 文件读取) ✅
**参考**: `fo-dicom-code/IO/Reader/DicomReader.cs` (~1300 lines)
**包**: `pkg/dicom/parser`

- [x] 实现 Parser 结构
- [x] 实现 Preamble 和 DICM 前缀读取
- [x] 实现 File Meta Information (Group 0x0002) 解析
- [x] 实现显式 VR 解析
- [x] 实现隐式 VR 解析 (使用字典查找)
- [x] 实现长度解析 (定长 vs 未定义长度)
- [x] 实现序列解析 (支持未定义长度和定义长度)
- [x] 实现 Item 解析 (支持未定义长度和定义长度)
- [x] 实现 ParseResult 结构 (包含 FileMetaInformation 和 Dataset)
- [x] 实现 Parse() 和 ParseFile() 方法
- [x] 实现解析选项 (WithMaxElementSize, WithStopAtTag, WithDictionary)
- [x] 修复序列读取时的 delimitation tag 处理问题
- [x] 编写单元测试 (8个测试函数)
- [x] **新增**: 实现 ReadOption 枚举 (Default, ReadLargeOnDemand, SkipLargeTags, ReadAll)
- [x] **新增**: 实现 FileFormat 枚举 (DICOM3, DICOM3NoPreamble, 等)
- [x] **新增**: ParseResult 添加 Format 和 IsPartial 字段
- [x] **新增**: 添加 LargeObjectSize 配置选项 (默认 64KB)
- [x] **新增**: 实现大数据标签处理逻辑
- [x] **新增**: 重构为 parseContext (内部实现)，只暴露包级函数
- [x] **新增**: Tag.DictionaryEntry() 方法 (查找字典条目)

**依赖**: DicomTag, DicomVR, TransferSyntax, DicomDataset, DicomDictionary
**实际工作量**: 5 天 (含新增功能)
**状态**: ✅ 完成，所有测试通过，功能完整对标 fo-dicom DicomFile

### 6.2 DicomWriter (DICOM 文件写入) ✅
**参考**: `fo-dicom-code/IO/Writer/DicomWriter.cs`, `DicomWriteOptions.cs`
**包**: `pkg/dicom/writer`

- [x] 实现 Writer 结构
- [x] 实现 Preamble 和 DICM 前缀写入
- [x] 实现 File Meta Information Group Length 自动计算
- [x] 实现 Tag 写入 (支持 Little/Big Endian)
- [x] 实现 VR 写入 (Explicit VR)
- [x] 实现 Length 写入 (16-bit/32-bit 自动选择)
- [x] 实现 Element 写入
- [x] 实现 Sequence 写入 (使用未定义长度)
- [x] 实现 Item 写入 (使用未定义长度)
- [x] 实现 Dataset 写入
- [x] 采用 Options 模式重构 API
  - [x] WithTransferSyntax() - 指定传输语法
  - [x] WithFileMetaInfo() - 指定 File Meta Information
  - [x] WithoutPreamble() - 跳过 preamble
- [x] 实现 Write() 函数 (写入到 io.Writer)
- [x] 实现 WriteFile() 函数 (写入到文件)
- [x] 自动生成 File Meta Information
- [x] 编写单元测试 (17个测试函数)
- [x] 编写 Roundtrip 测试 (写入后读取验证)
- [x] **新增**: 实现 DicomWriteOptions 全部选项
  - [x] WithExplicitLengthSequences() - 序列使用显式长度
  - [x] WithExplicitLengthSequenceItems() - 序列项使用显式长度
  - [x] WithKeepGroupLengths() - 保留 group length 标签
  - [x] WithLargeObjectSize(size) - 大对象阈值 (默认 1MB)
- [x] **新增**: 实现 group length 标签过滤逻辑
- [x] **新增**: 实现显式长度序列写入 (自动计算长度)
- [x] **新增**: 实现显式长度序列项写入

**依赖**: DicomTag, DicomVR, TransferSyntax, DicomDataset
**实际工作量**: 4 天 (含新增功能)
**状态**: ✅ 完成，所有测试通过，功能完整对标 fo-dicom DicomWriteOptions

---

## 第七阶段：序列化和工具 ~

### 7.1 JSON 序列化 ✅
**参考**: `fo-dicom-code/Serialization/JsonDicomConverter.cs` (~1400 lines)
**包**: `pkg/dicom/serialization`

- [x] 实现 DICOM JSON 模型序列化 (ToJSON)
- [x] 实现 DICOM JSON 反序列化 (FromJSON)
- [x] 实现序列化选项 (Options 模式)
  - [x] WithWriteTagsAsKeywords (关键字作为 key)
  - [x] WithWriteKeyword / WithWriteName (额外字段)
  - [x] WithNumberSerializationMode (数字序列化模式)
  - [x] WithIndent (格式化输出)
  - [x] WithAutoValidate (反序列化验证)
- [x] 支持所有 DICOM VR 类型
  - [x] 字符串类型 (AE, CS, DA, LO, PN, UI 等)
  - [x] 数值类型 (FL, FD, SL, SS, SV, UL, US, UV)
  - [x] 数值字符串 (IS, DS)
  - [x] 二进制类型 (OB, OW, InlineBinary)
  - [x] 序列 (SQ)
  - [x] PersonName (分组格式)
  - [x] AttributeTag (AT)
- [x] 特殊值处理
  - [x] NaN, Infinity, -Infinity
  - [x] null 值
  - [x] 空元素
- [x] 支持 BulkDataURI
  - [x] 实现 BulkDataUriByteBuffer (pkg/io/buffer/bulkdata.go)
  - [x] 在 JSON 序列化中集成 BulkDataURI 检测和输出
  - [x] 在 JSON 反序列化中支持 BulkDataURI 解析
  - [x] 支持所有相关 VR 类型 (OB, OW, OD, OF, OL, OV, UN)
  - [x] 实现延迟加载（lazy loading）模式
- [x] 编写单元测试
  - [x] BulkDataUriByteBuffer 单元测试 (11 个测试用例)
  - [x] JSON 序列化集成测试 (7 个测试用例)
  - [x] 所有测试通过

**依赖**: DicomDataset, DicomElement, DicomTag, DicomVR
**实际工作量**: 2 天
**状态**: ✅ 完整功能完成，所有测试通过
**完成日期**: 2025-11-07

### 7.2 XML 序列化 ✅
**参考**: `fo-dicom-code/Serialization/DicomXML.cs`
**包**: `pkg/dicom/serialization`

- [x] 实现 DICOM XML 序列化 (NativeDicomModel格式)
- [x] 实现 ToXML() 函数
- [x] 实现 XML 选项配置 (WithXMLIndent, WithXMLCompact)
- [x] 支持所有 DICOM 元素类型
  - [x] 字符串元素 (使用 `<Value>` 标签)
  - [x] 数值元素 (US, UL, SS, SL, FL, FD, SV, UV - 自动转换为字符串)
  - [x] 二进制元素 (使用 `<InlineBinary>` base64编码)
  - [x] PersonName (结构化格式: Alphabetic/Ideographic/Phonetic)
  - [x] 序列 (使用 `<Item>` 标签)
- [x] 支持嵌套序列
- [x] XML 特殊字符转义 (&, <, >, ", ')
- [x] 支持私有标签 (privateCreator 属性)
- [x] 支持 keyword 属性 (从字典查找)
- [x] 编写单元测试 (16个测试用例，全部通过)
- [ ] 实现 FromXML() 反序列化 (未实现 - 可选功能)

**依赖**: DicomDataset, DicomElement, DicomDictionary
**实际工作量**: 1 天
**状态**: ✅ 完成 XML 序列化，所有测试通过
**完成日期**: 2025-11-07
**优先级**: 低 (FromXML 反序列化可根据需求决定是否实现)

### 7.3 DicomAnonymizer ✅
**参考**: `fo-dicom-code/DicomAnonymizer.cs`, `DicomAnonymizerGenerated.cs`
**包**: `pkg/dicom/anonymizer`

- [x] 实现匿名化规则 (SecurityProfile, SecurityProfileActions)
- [x] 实现常用匿名化操作 (Remove, Replace, Clear, Blank, UID replacement)
- [x] 支持自定义匿名化规则 (AddRule)
- [x] 实现安全配置选项 (BasicProfile, RetainUIDs, RetainDeviceIdent 等)
- [x] 实现默认 DICOM PS 3.15 配置文件
- [x] 支持序列递归匿名化
- [x] 支持自定义 PatientName 和 PatientID 替换
- [x] 实现 UID 一致性替换 (ReplacedUIDs 映射)
- [x] 编写单元测试 (13个测试函数，全部通过)

**依赖**: DicomDataset, DicomElement
**实际工作量**: 1 天
**状态**: ✅ 完成，所有测试通过
**完成日期**: 2025-11-07

---

## 第八阶段：命令行工具和示例 ✅

### 8.1 命令行工具 ✅
**包**: `cmd/*`

- [x] 创建 `dicominfo` 工具 (显示 DICOM 文件信息)
  - 显示文件格式和传输语法
  - 显示 File Meta Information
  - 显示患者、检查、系列、图像信息
  - 支持详细输出模式 (-v)
- [x] 创建 `dicomdump` 工具 (转储所有标签)
  - 递归显示所有元素
  - 支持序列嵌套显示
  - 可配置最大深度
  - 支持紧凑模式输出
- [x] 创建 `dicom2json` 工具 (转换为 JSON)
  - 输出 DICOM JSON Model 格式
  - 支持关键字作为 JSON key
  - 支持紧凑和格式化输出
  - 可选包含 File Meta Information

**依赖**: 前述所有功能
**实际工作量**: 0.5 天
**状态**: ✅ 完成，3个命令行工具全部实现

### 8.2 示例代码 ✅
**目录**: `examples/*`

- [x] 读取 DICOM 文件示例 (`examples/read_dicom/`)
- [x] 创建和保存 DICOM 文件示例 (`examples/write_dicom/`)
- [x] 匿名化示例 (`examples/anonymize/`)
- [x] JSON 序列化示例 (`examples/json_conversion/`)

**依赖**: 核心功能
**实际工作量**: 0.5 天
**状态**: ✅ 完成，4个示例程序全部实现
**完成日期**: 2025-11-07

---

## 第九阶段：网络功能 (DIMSE/DICOM 网络) [可选]

**注**: 这是最复杂的部分，建议在核心功能稳定后再实现。本阶段采用自底向上的增量开发策略。

**Go 设计原则**:
- ✅ 使用 `net.Conn` 替代 C# 的 `INetworkStream`
- ✅ 使用 goroutines + channels 替代 C# 的 async/await
- ✅ 使用 `context.Context` 进行超时和取消控制
- ✅ 使用组合（composition）替代继承（inheritance）
- ✅ 使用回调函数或 channels 替代 C# 事件
- ✅ 使用小而专注的接口，避免大接口

**包结构设计** (Go 风格):
```
pkg/network/
├── pdu/              # PDU 层 (最底层，无网络依赖)
├── association/      # Association 管理 (无网络依赖)
├── dimse/            # DIMSE 消息定义 (无网络依赖)
├── transport/        # 网络传输抽象 (net.Conn 封装)
├── service/          # 核心服务层 (PDU I/O + 状态机)
├── client/           # SCU 客户端实现
├── server/           # SCP 服务端实现
└── status/           # DICOM 状态码
```

---

### 9.1 基础类型和常量 ✅
**参考**: `fo-dicom-code/Network/DicomStatus.cs`, `DicomCommandField.cs`, `DicomPriority.cs`
**包**: `pkg/network/status`, `pkg/network/dimse`

**任务**:
- [x] 实现 DicomStatus 类型和常量 (800+ 标准状态码)
  - [x] 定义 Status 结构体 (Code uint16, State string, Description string)
  - [x] 实现分类常量: Success, Pending, Cancel, Warning, Failure
  - [x] 实现具体状态码常量 (Success = 0x0000, Pending = 0xFF00, 等)
  - [x] 实现状态查询方法: IsSuccess(), IsPending(), IsWarning(), IsFailure()
  - [x] 实现 LookupStatus() 方法
- [x] 实现 CommandField 枚举
  - [x] C-STORE-RQ/RSP (0x0001/0x8001)
  - [x] C-FIND-RQ/RSP (0x0020/0x8020)
  - [x] C-GET-RQ/RSP (0x0010/0x8010)
  - [x] C-MOVE-RQ/RSP (0x0021/0x8021)
  - [x] C-ECHO-RQ/RSP (0x0030/0x8030)
  - [x] N-EVENT-REPORT-RQ/RSP, N-GET-RQ/RSP, 等 (7 对)
  - [x] 实现 String(), IsRequest(), IsResponse() 方法
- [x] 实现 Priority 枚举 (Low=0x0002, Medium=0x0000, High=0x0001)
  - [x] 实现 String() 方法
- [x] 编写单元测试 (13个测试函数，全部通过)

**依赖**: 无
**实际工作量**: 1 天
**优先级**: P0 (必须先实现，其他模块依赖)
**状态**: ✅ 完成，所有测试通过

---

### 9.2 PDU 层 (Protocol Data Unit) - 二进制协议 ✅
**参考**: `fo-dicom-code/Network/PDU.cs` (~2100 lines)
**包**: `pkg/network/pdu`
**实际工作量**: 4 天
**状态**: ✅ 完成，所有测试通过 (72个测试用例)

**设计要点**:
- PDU 是纯数据结构 + 编解码器，不涉及网络 I/O
- 使用 `encoding/binary` 进行字节序处理
- 使用 `io.Reader`/`io.Writer` 接口（而非具体的 net.Conn）

**任务**:

#### 9.2.1 RawPDU 基础结构 ✅
- [x] 定义 PDU 类型常量
  - [x] A-ASSOCIATE-RQ (0x01)
  - [x] A-ASSOCIATE-AC (0x02)
  - [x] A-ASSOCIATE-RJ (0x03)
  - [x] P-DATA-TF (0x04)
  - [x] A-RELEASE-RQ (0x05)
  - [x] A-RELEASE-RP (0x06)
  - [x] A-ABORT (0x07)
- [x] 实现 RawPDU 结构体
  ```go
  type RawPDU struct {
      Type     byte   // PDU type
      Reserved byte   // Reserved (0x00)
      Length   uint32 // PDU length (excludes 6-byte header)
      Data     []byte // PDU data
  }
  ```
- [x] 实现 Read(r io.Reader) error - 从流读取 PDU
- [x] 实现 Write(w io.Writer) error - 写入 PDU 到流
- [x] 实现辅助函数: NewRawPDU(), String(), TotalLength(), PDUTypeString()
- [x] 实现便利函数: ReadPDU(), WritePDU()
- [x] 编写单元测试 (11个测试函数，全部通过)
  - [x] PDU 类型字符串测试
  - [x] 构造函数测试
  - [x] 读写往返测试
  - [x] 错误处理测试（无效类型、不完整数据等）

**状态**: ✅ 完成，所有测试通过

#### 9.2.2 A-ASSOCIATE-RQ PDU ✅
- [x] 定义 Item 类型常量 (ApplicationContext=0x10, PresentationContext=0x20, 等 - 15个)
- [x] 实现 AAssociateRQ 结构体
  ```go
  type AAssociateRQ struct {
      ProtocolVersion       uint16
      CalledAETitle         string  // 被调用 AE (max 16 bytes)
      CallingAETitle        string  // 调用方 AE (max 16 bytes)
      ApplicationContext    string  // 通常是 "1.2.840.10008.3.1.1.1"
      PresentationContexts  []PresentationContextRQ
      UserInformation       *UserInformation
  }
  ```
- [x] 实现 PresentationContextRQ 结构体
  ```go
  type PresentationContextRQ struct {
      ID                byte   // Context ID (奇数: 1, 3, 5, ...)
      AbstractSyntax    string // SOP Class UID
      TransferSyntaxes  []string // 提议的传输语法列表
  }
  ```
- [x] 实现 UserInformation 结构体
  - [x] MaximumLength (0x51) - 最大 PDU 长度
  - [x] ImplementationClassUID (0x52)
  - [x] ImplementationVersionName (0x55)
  - [x] AsynchronousOperationsWindow (0x53) - 可选
  - [x] SCP_SCU_RoleSelection (0x54) - 可选
  - [x] ExtendedNegotiation (0x56) - 可选
  - [x] UserIdentityNegotiation (0x58) - 可选 (支持5种类型)
- [x] 实现 Encode() (*RawPDU, error) - 编码为 PDU
- [x] 实现 Decode(*RawPDU) error - 从 PDU 解码
- [x] 实现 Item 辅助函数 (readItem, writeItem, readString, writeString)
- [x] 编写单元测试 (10个测试函数，全部通过)
  - 基础编解码往返测试
  - 多 Presentation Context 测试
  - UserIdentity 测试 (Username, Username+Password)
  - ExtendedNegotiation 测试
  - RoleSelection 测试
  - ImplementationVersion 测试
  - AE Title 空格填充测试

**状态**: ✅ 完成，10个测试全部通过
**实现文件**:
- `pkg/network/pdu/items.go` (180 lines)
- `pkg/network/pdu/associate_rq.go` (520 lines)
- `pkg/network/pdu/associate_rq_test.go` (398 lines)

#### 9.2.3 A-ASSOCIATE-AC PDU ✅
- [x] 实现 AAssociateAC 结构体 (类似 RQ，但包含协商结果)
- [x] 实现 PresentationContextAC 结构体
- [x] 实现编解码方法
- [x] 编写单元测试 (10个测试用例)
- [x] 实现 Result 枚举和 ResultString()

**状态**: ✅ 完成

#### 9.2.4 A-ASSOCIATE-RJ PDU ✅
- [x] 实现 AAssociateRJ 结构体
- [x] 实现编解码方法
- [x] 实现原因码解释方法 (ResultString, SourceString, ReasonString)
- [x] 编写单元测试 (10个测试用例)
- [x] 支持所有标准拒绝原因

**状态**: ✅ 完成

#### 9.2.5 P-DATA-TF PDU (数据传输) ✅
- [x] 实现 PDataTF 结构体
- [x] 实现 PDV (Presentation Data Value) 结构体
- [x] 实现编解码方法
- [x] 实现分片逻辑 (FragmentData)
- [x] 实现重组逻辑 (ReassembleData)
- [x] 编写单元测试 (20个测试用例)
  - 单/多 PDV 测试
  - 大消息分片测试
  - 混合 Context 和类型测试

**状态**: ✅ 完成

#### 9.2.6 A-RELEASE-RQ/RP PDU ✅
- [x] 实现 AReleaseRQ 结构体 (无数据字段，仅有 header)
- [x] 实现 AReleaseRP 结构体 (无数据字段，仅有 header)
- [x] 实现编解码方法
- [x] 编写单元测试 (6个测试用例)
- [x] Release 往返测试

**状态**: ✅ 完成

#### 9.2.7 A-ABORT PDU ✅
- [x] 实现 AAbort 结构体
- [x] 实现编解码方法
- [x] 实现原因码解释方法 (SourceString, ReasonString)
- [x] 编写单元测试 (14个测试用例)
- [x] 支持所有服务用户和服务提供者原因

**状态**: ✅ 完成

**依赖**: encoding/binary, io, DicomUID, DicomTransferSyntax
**实际工作量**: 4 天
**优先级**: P0 (基础协议层)
**总测试数**: 72个测试用例，全部通过

---

### 9.3 Association 管理 (无网络依赖) ✅
**参考**: `fo-dicom-code/Network/DicomAssociation.cs`, `DicomPresentationContext.cs`
**包**: `pkg/network/association`
**实际工作量**: 2 天
**状态**: ✅ 完成，所有测试通过 (14个测试用例)

**设计要点**:
- Association 是纯数据结构，描述连接元数据
- 不包含网络逻辑，只负责协商状态管理

**任务**:

#### 9.3.1 PresentationContext ✅
- [x] 实现 PresentationContext 结构体
- [x] 实现 AcceptanceResult 枚举 (Acceptance, UserRejection, NoReason, etc.)
- [x] 实现 Accept/Reject 方法
- [x] 实现 IsAccepted() bool 方法
- [x] 编写单元测试

**状态**: ✅ 完成

#### 9.3.2 ExtendedNegotiation ✅
- [x] 实现 ExtendedNegotiation 结构体
- [x] 支持任意 SOP Class 扩展协商数据
- [x] 编写单元测试

**状态**: ✅ 完成

#### 9.3.3 UserIdentity ✅
- [x] 实现 UserIdentity 结构体
- [x] 支持 5 种认证类型
  - Username (Type=1)
  - Username+Password (Type=2)
  - Kerberos (Type=3)
  - SAML (Type=4)
  - JWT (Type=5)
- [x] 实现 NewUserIdentityUsername 和 NewUserIdentityUsernamePassword
- [x] 编写单元测试

**状态**: ✅ 完成

#### 9.3.4 Association ✅
- [x] 实现 Association 结构体 (包含 MessageIDGenerator)
- [x] 实现 AddPresentationContext() 方法
- [x] 实现 FindPresentationContextByID(id) 方法
- [x] 实现 FindPresentationContextByAbstractSyntax(uid) 方法
- [x] 实现 GetAcceptedPresentationContexts() 方法
- [x] 实现 ExtendedNegotiation 管理
- [x] 实现 RoleSelection 支持 (AsynchronousOperationsWindow)
- [x] 实现 SetEstablished() 状态管理
- [x] 实现 GetTransferSyntaxForAbstractSyntax() 辅助方法
- [x] 编写单元测试 (14个测试用例)

**状态**: ✅ 完成

**依赖**: DicomUID, DicomTransferSyntax, dimse.MessageIDGenerator
**实际工作量**: 2 天
**优先级**: P0
**总测试数**: 14个测试用例，全部通过

---

### 9.4 DIMSE 消息层 (无网络依赖) ✅
**参考**: `fo-dicom-code/Network/DicomMessage.cs`, `DicomRequest.cs`, `DicomResponse.cs`
**包**: `pkg/network/dimse`
**实际工作量**: 3 天
**状态**: ✅ 完成，所有测试通过 (31个测试用例 + 3个示例)

**设计要点**:
- DIMSE 消息本质是 DICOM Dataset + 命令字段
- 使用已有的 dataset.Dataset，添加网络特定字段
- MessageID 自动生成机制 (每个 Association 拥有独立的 MessageIDGenerator)

**任务**:

#### 9.4.1 Message 基础 ✅
- [x] 实现 Message, Request, Response 接口
- [x] 实现 BaseMessage 结构体 (所有消息的基础)
- [x] 实现 BaseRequest 和 BaseResponse
- [x] 实现辅助方法: CommandField(), MessageID(), SetMessageID(), 等
- [x] 实现 MessageID 缓存机制
- [x] 编写单元测试

**状态**: ✅ 完成

#### 9.4.2 MessageID 自动生成器 ✅
- [x] 实现 MessageIDGenerator (线程安全的原子计数器)
- [x] 实现 Next() 方法 (1-65535 循环)
- [x] 实现 AssignMessageID() 方法 (自动分配或保留已设置的 ID)
- [x] 实现 Reset() 方法
- [x] 实现 Wraparound 处理 (65535 → 1, 跳过 0)
- [x] 集成到 Association 结构 (每个 Association 拥有独立的生成器)
- [x] 编写并发测试 (100 goroutines * 100 IDs)

**状态**: ✅ 完成

#### 9.4.3 C-ECHO (验证连通性) ✅
- [x] 实现 CEchoRequest 结构体
- [x] 实现 NewCEchoRequest() 构造函数 (无需 messageID 参数)
- [x] 实现 CEchoResponse 结构体
- [x] 实现 NewCEchoResponseSuccess() 构造函数
- [x] 实现 AffectedSOPClassUID(), Priority() 方法
- [x] 实现 Status 查询方法 (IsSuccess, IsPending, IsFailure)
- [x] 编写单元测试 (序列化/反序列化往返)

**状态**: ✅ 完成

#### 9.4.4 C-STORE (存储图像) ✅
- [x] 实现 CStoreRequest 结构体
- [x] 实现 NewCStoreRequest(ds) 构造函数 (自动提取 SOP Class/Instance UID)
- [x] 实现 SetPriority() 方法
- [x] 实现 SetMoveOriginator() 方法 (C-MOVE 子操作支持)
- [x] 实现 CStoreResponse 结构体
- [x] 实现 NewCStoreResponseSuccess() 构造函数
- [x] 实现 SOP Class/Instance UID 验证
- [x] 编写单元测试 (包含缺失 UID 的错误测试)

**状态**: ✅ 完成

#### 9.4.5 C-FIND (查询) ✅
- [x] 实现 CFindRequest 结构体
- [x] 实现 QueryRetrieveLevel 枚举 (PATIENT, STUDY, SERIES, IMAGE)
- [x] 实现 NewCFindRequest() 通用构造函数
- [x] 实现 NewCFindRequestPatientRoot() 便捷构造函数
- [x] 实现 NewCFindRequestStudyRoot() 便捷构造函数
- [x] 实现 SetPriority() 方法
- [x] 自动添加 QueryRetrieveLevel 到 query dataset
- [x] 实现 CFindResponse 结构体
- [x] 实现 NewCFindResponsePending() (带 identifier)
- [x] 实现 NewCFindResponseSuccess() (最终响应)
- [x] 实现 HasIdentifier() 方法
- [x] 实现 IsPending() 方法
- [x] 编写单元测试 (包含 Patient Root 和 Study Root)

**状态**: ✅ 完成

#### 9.4.6 C-GET (检索 - Push 模式)
- [ ] 实现 CGetRequest (类似 C-FIND)
- [ ] 实现 CGetResponse
  ```go
  type CGetResponse struct {
      *BaseMessage
      messageIDBeingRespondedTo     uint16
      status                        *status.Status
      affectedSOPClassUID           string
      numberOfRemainingSuboperations   uint16
      numberOfCompletedSuboperations   uint16
      numberOfFailedSuboperations      uint16
      numberOfWarningSuboperations     uint16
  }
  ```
- [ ] 编写单元测试

#### 9.4.7 C-MOVE (检索 - Pull 模式)
- [ ] 实现 CMoveRequest
  ```go
  type CMoveRequest struct {
      *BaseMessage
      messageID             uint16
      affectedSOPClassUID   string
      priority              Priority
      moveDestination       string // 目标 AE Title
      // data 字段包含查询条件
  }
  ```
- [ ] 实现 CMoveResponse (类似 C-GET)
- [ ] 编写单元测试

#### 9.4.8 N-* 操作 ✅
- [x] N-EVENT-REPORT-RQ/RSP
- [x] N-GET-RQ/RSP
- [x] N-SET-RQ/RSP
- [x] N-ACTION-RQ/RSP
- [x] N-CREATE-RQ/RSP
- [x] N-DELETE-RQ/RSP

**依赖**: dataset.Dataset, DicomTag, Status, association.Association
**实际工作量**: 3 天
**优先级**: P1
**总测试数**: 102个测试用例 + 3个示例，全部通过

**API 设计**:
```go
// C-* 操作 (用户代码)
req := dimse.NewCEchoRequest()           // MessageID 默认为 0
storeReq, _ := dimse.NewCStoreRequest(ds)
findReq := dimse.NewCFindRequestPatientRoot(level, query)

// N-* 操作
eventReq := dimse.NewNEventReportRequest(sopClassUID, sopInstanceUID, eventTypeID, eventInfo)
getReq := dimse.NewNGetRequest(sopClassUID, sopInstanceUID, attrList)
setReq := dimse.NewNSetRequest(sopClassUID, sopInstanceUID, modList)
actionReq := dimse.NewNActionRequest(sopClassUID, sopInstanceUID, actionTypeID, actionInfo)
createReq := dimse.NewNCreateRequest(sopClassUID, "", attrList) // "" = server assigns UID
deleteReq := dimse.NewNDeleteRequest(sopClassUID, sopInstanceUID)

// Association 自动分配 MessageID
msgID, _ := association.AssignMessageID(req)

// 单元测试手动设置
req.SetMessageID(123)
```

**实现文件**:
- `pkg/network/dimse/message.go` (265 lines) - 基础消息接口和结构
- `pkg/network/dimse/cecho.go` (104 lines) - C-ECHO 实现
- `pkg/network/dimse/cstore.go` (144 lines) - C-STORE 实现
- `pkg/network/dimse/cfind.go` (168 lines) - C-FIND 实现
- `pkg/network/dimse/neventreport.go` (207 lines) - N-EVENT-REPORT 实现
- `pkg/network/dimse/nget.go` (179 lines) - N-GET 实现
- `pkg/network/dimse/nset.go` (158 lines) - N-SET 实现
- `pkg/network/dimse/naction.go` (207 lines) - N-ACTION 实现
- `pkg/network/dimse/ncreate.go` (193 lines) - N-CREATE 实现
- `pkg/network/dimse/ndelete.go` (142 lines) - N-DELETE 实现
- `pkg/network/dimse/messageid.go` (110 lines) - MessageID 生成器
- `pkg/network/dimse/dimse_test.go` (374 lines) - C-* 操作测试
- `pkg/network/dimse/noperation_test.go` (504 lines) - N-* 操作测试
- `pkg/network/dimse/messageid_test.go` (203 lines) - MessageID 测试
- `pkg/network/dimse/example_messageid_test.go` (103 lines) - 示例程序
- `pkg/network/association/association.go` - MessageIDGenerator 集成

---

### 9.5 网络传输层 (Go 风格抽象)
**参考**: `fo-dicom-code/Network/INetworkStream.cs`, `NetworkManager.cs`
**包**: `pkg/network/transport`

**设计要点**:
- 使用 Go 的 `net.Conn` 接口，无需自定义 NetworkStream
- 提供 TLS 封装和配置
- 提供超时控制

**任务**:

#### 9.5.1 TCP 连接封装
- [ ] 实现 Dialer 结构体
  ```go
  type Dialer struct {
      Timeout       time.Duration
      KeepAlive     time.Duration
      TLSConfig     *tls.Config // 可选
  }
  ```
- [ ] 实现 Dial(ctx, host, port) (net.Conn, error) 方法
- [ ] 实现 DialTLS(ctx, host, port) (net.Conn, error) 方法
- [ ] 编写单元测试 (需要 mock 服务器)

#### 9.5.2 TCP 监听器封装
- [ ] 实现 Listener 结构体
  ```go
  type Listener struct {
      listener  net.Listener
      TLSConfig *tls.Config // 可选
  }
  ```
- [ ] 实现 Listen(host, port) (*Listener, error)
- [ ] 实现 Accept(ctx) (net.Conn, error) 方法
- [ ] 实现 Close() error 方法
- [ ] 编写单元测试

#### 9.5.3 连接读写辅助
- [ ] 实现 ReadPDU(conn, timeout) (*pdu.RawPDU, error)
  - 使用 conn.SetReadDeadline() 设置超时
- [ ] 实现 WritePDU(conn, timeout, pdu) error
  - 使用 conn.SetWriteDeadline() 设置超时
- [ ] 编写单元测试

**依赖**: net, crypto/tls, pdu
**预计工作量**: 1-2 天
**优先级**: P1

---

### 9.6 核心服务层 (状态机 + PDU I/O)
**参考**: `fo-dicom-code/Network/DicomService.cs` (~2400 lines)
**包**: `pkg/network/service`

**设计要点**:
- 不使用继承，使用组合和接口
- 使用 goroutines 分离发送/接收循环
- 使用 channels 进行消息队列管理
- 使用 context.Context 进行生命周期管理

**任务**:

#### 9.6.1 服务状态机
- [ ] 定义 ServiceState 枚举
  ```go
  type State int
  const (
      StateIdle State = iota
      StateAssociationRequested  // SCU 已发送 A-ASSOCIATE-RQ
      StateAssociationAccepted   // 关联已建立
      StateTransferring          // 正在传输 DIMSE 消息
      StateReleaseRequested      // 已发送 A-RELEASE-RQ
      StateClosed                // 连接已关闭
  )
  ```
- [ ] 实现状态转换验证逻辑
- [ ] 编写单元测试

#### 9.6.2 Service 核心结构
- [ ] 实现 Service 结构体
  ```go
  type Service struct {
      conn           net.Conn
      assoc          *association.Association
      state          State
      stateMu        sync.RWMutex

      // Goroutine 通信
      sendQueue      chan SendRequest
      recvQueue      chan dimse.Message
      closeOnce      sync.Once
      closeCh        chan struct{}
      errCh          chan error

      // 配置
      Options        *Options

      // 消息处理
      pendingRequests map[uint16]*PendingRequest // MessageID -> Request
      handlers        *Handlers

      // 接收缓冲
      pduBuffer      *PDUBuffer // 重组 PDV
  }
  ```
- [ ] 实现 NewService(conn, assoc, options) *Service
- [ ] 编写单元测试

#### 9.6.3 发送循环 (sendLoop)
- [ ] 实现 sendLoop(ctx) error
  ```go
  func (s *Service) sendLoop(ctx context.Context) error {
      for {
          select {
          case <-ctx.Done():
              return ctx.Err()
          case <-s.closeCh:
              return nil
          case req := <-s.sendQueue:
              // 1. 将 DIMSE 消息编码为 Command Dataset + Data Dataset
              // 2. 分片为 PDV (如果 > MaxPDULength)
              // 3. 封装为 P-DATA-TF PDU
              // 4. 写入 conn
              // 5. 如果是 Request，添加到 pendingRequests
          }
      }
  }
  ```
- [ ] 实现消息分片逻辑 fragmentMessage()
- [ ] 编写单元测试

#### 9.6.4 接收循环 (recvLoop)
- [ ] 实现 recvLoop(ctx) error
  ```go
  func (s *Service) recvLoop(ctx context.Context) error {
      for {
          select {
          case <-ctx.Done():
              return ctx.Err()
          case <-s.closeCh:
              return nil
          default:
              // 1. 读取 RawPDU
              // 2. 根据 PDU 类型分发
              // 3. P-DATA-TF: 重组 PDV -> DIMSE 消息 -> recvQueue
              // 4. A-RELEASE-RQ: 发送 A-RELEASE-RP, 关闭连接
              // 5. A-ABORT: 立即关闭
          }
      }
  }
  ```
- [ ] 实现 PDV 重组逻辑 (PDUBuffer)
- [ ] 实现 DIMSE 消息解码 decodeDIMSE()
- [ ] 编写单元测试

#### 9.6.5 消息发送 API
- [ ] 实现 Send(ctx, msg dimse.Message) error
  ```go
  func (s *Service) Send(ctx context.Context, msg dimse.Message) error {
      select {
      case s.sendQueue <- SendRequest{Message: msg}:
          return nil
      case <-ctx.Done():
          return ctx.Err()
      case <-s.closeCh:
          return ErrServiceClosed
      }
  }
  ```
- [ ] 实现 SendRequest(ctx, req dimse.Request) (dimse.Response, error)
  - 发送 Request
  - 等待 Response (通过 pendingRequests 映射)
  - 支持超时
- [ ] 编写单元测试

#### 9.6.6 消息处理器 (Handlers)
- [ ] 定义 Handler 接口
  ```go
  type Handler interface {
      Handle(ctx context.Context, msg dimse.Message) (dimse.Message, error)
  }
  ```
- [ ] 实现 Handlers 结构体 (消息路由)
  ```go
  type Handlers struct {
      CEchoHandler  func(context.Context, *dimse.CEchoRequest) (*dimse.CEchoResponse, error)
      CStoreHandler func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error)
      CFindHandler  func(context.Context, *dimse.CFindRequest) ([]*dimse.CFindResponse, error) // 返回多个 Pending + 1 个 Success
      CGetHandler   func(context.Context, *dimse.CGetRequest) error // 触发多个 C-STORE
      CMoveHandler  func(context.Context, *dimse.CMoveRequest) error
  }
  ```
- [ ] 实现消息分发逻辑 dispatchMessage()
- [ ] 编写单元测试

#### 9.6.7 关联管理
- [ ] 实现 SendAssociationRequest(ctx, assoc) error
- [ ] 实现 SendAssociationAccept(ctx, assoc) error
- [ ] 实现 SendAssociationReject(ctx, result, source, reason) error
- [ ] 实现 SendReleaseRequest(ctx) error
- [ ] 实现 SendAbort(ctx, source, reason) error
- [ ] 编写单元测试

#### 9.6.8 生命周期管理
- [ ] 实现 Run(ctx) error - 启动 sendLoop 和 recvLoop
- [ ] 实现 Close() error - 优雅关闭
- [ ] 实现 Abort(source, reason) - 强制关闭
- [ ] 编写单元测试

**依赖**: pdu, association, dimse, transport, context
**预计工作量**: 7-8 天
**优先级**: P1 (核心，但依赖前面所有模块)

---

### 9.7 Client 实现 (SCU - Service Class User)
**参考**: `fo-dicom-code/Network/Client/DicomClient.cs`
**包**: `pkg/network/client`

**设计要点**:
- 封装 Service，提供更高级的 API
- 支持连接池和重连
- 提供同步和异步 API

**任务**:

#### 9.7.1 Client 结构
- [ ] 实现 Client 结构体
  ```go
  type Client struct {
      service  *service.Service
      options  *Options
      assoc    *association.Association
  }
  ```
- [ ] 实现 ClientOptions
  ```go
  type Options struct {
      CallingAE          string
      CalledAE           string
      MaxPDULength       uint32
      ConnectTimeout     time.Duration
      RequestTimeout     time.Duration
      AssociationTimeout time.Duration
  }
  ```
- [ ] 编写单元测试

#### 9.7.2 连接和关联
- [ ] 实现 Connect(ctx, host, port) error
  ```go
  func (c *Client) Connect(ctx context.Context, host string, port int) error {
      // 1. TCP 连接
      // 2. 发送 A-ASSOCIATE-RQ
      // 3. 等待 A-ASSOCIATE-AC/RJ
      // 4. 启动 Service.Run()
  }
  ```
- [ ] 实现 AddPresentationContext(abstractSyntax, transferSyntaxes) 方法
- [ ] 实现 Disconnect(ctx) error - 发送 A-RELEASE-RQ
- [ ] 编写单元测试 (需要 mock SCP 服务器)

#### 9.7.3 DIMSE 操作 API
- [ ] 实现 CEcho(ctx) error
  ```go
  func (c *Client) CEcho(ctx context.Context) error {
      req := dimse.NewCEchoRequest()
      resp, err := c.service.SendRequest(ctx, req)
      // 检查 resp.Status
  }
  ```
- [ ] 实现 CStore(ctx, dataset) error
  ```go
  func (c *Client) CStore(ctx context.Context, ds *dataset.Dataset) error {
      req := dimse.NewCStoreRequest(ds)
      resp, err := c.service.SendRequest(ctx, req)
      // 检查 resp.Status
  }
  ```
- [ ] 实现 CFind(ctx, level, query) ([]dataset.Dataset, error)
  ```go
  func (c *Client) CFind(ctx context.Context, level QueryRetrieveLevel, query *dataset.Dataset) ([]*dataset.Dataset, error) {
      req := dimse.NewCFindRequest(level, query)
      // 发送请求，收集所有 Pending 响应，直到 Success
  }
  ```
- [ ] 实现 CGet(ctx, level, query, storeHandler) error
- [ ] 实现 CMove(ctx, level, query, destination) error
- [ ] 编写单元测试

#### 9.7.4 高级特性
- [ ] 实现连接池 ClientPool
- [ ] 实现自动重连逻辑
- [ ] 实现请求重试机制
- [ ] 编写单元测试

**依赖**: service, association, dimse, transport
**预计工作量**: 3-4 天
**优先级**: P2

---

### 9.8 Server 实现 (SCP - Service Class Provider)
**参考**: `fo-dicom-code/Network/DicomServer.cs`
**包**: `pkg/network/server`

**设计要点**:
- 监听 TCP 端口
- 为每个连接创建独立的 Service
- 使用用户提供的 Handlers 处理请求

**任务**:

#### 9.8.1 Server 结构
- [ ] 实现 Server 结构体
  ```go
  type Server struct {
      listener *transport.Listener
      options  *Options
      handlers *service.Handlers

      // 连接管理
      conns    map[string]*service.Service // 跟踪活跃连接
      connsMu  sync.RWMutex

      closeOnce sync.Once
      closeCh   chan struct{}
  }
  ```
- [ ] 实现 ServerOptions
  ```go
  type Options struct {
      AETitle           string
      Port              int
      MaxConnections    int
      MaxPDULength      uint32
      AssociationTimeout time.Duration
      IdleTimeout       time.Duration
      TLSConfig         *tls.Config // 可选
  }
  ```
- [ ] 编写单元测试

#### 9.8.2 监听和接受连接
- [ ] 实现 ListenAndServe(ctx) error
  ```go
  func (s *Server) ListenAndServe(ctx context.Context) error {
      // 1. 启动 TCP 监听
      // 2. 循环 Accept 连接
      // 3. 为每个连接启动 goroutine 处理
  }
  ```
- [ ] 实现 handleConnection(conn) 方法
  ```go
  func (s *Server) handleConnection(conn net.Conn) {
      // 1. 等待 A-ASSOCIATE-RQ
      // 2. 调用 OnAssociationRequest 回调
      // 3. 发送 A-ASSOCIATE-AC/RJ
      // 4. 创建 Service 并启动
      // 5. 连接关闭后清理
  }
  ```
- [ ] 实现连接数限制逻辑
- [ ] 编写单元测试

#### 9.8.3 关联协商回调
- [ ] 实现 OnAssociationRequest 回调接口
  ```go
  type AssociationNegotiator interface {
      OnAssociationRequest(*association.Association) (*association.Association, error)
  }
  ```
- [ ] 提供默认实现 (接受所有 Presentation Contexts)
- [ ] 编写单元测试

#### 9.8.4 生命周期管理
- [ ] 实现 Shutdown(ctx) error - 优雅关闭所有连接
- [ ] 实现 Close() error - 立即关闭
- [ ] 实现连接数统计和监控
- [ ] 编写单元测试

**依赖**: service, association, transport
**预计工作量**: 3-4 天
**优先级**: P2

---

### 9.9 集成测试和示例
**包**: `pkg/network` (测试), `examples/dicom_server`, `examples/dicom_client`

**任务**:

#### 9.9.1 本地集成测试
- [ ] 实现 TestClientServerCEcho - 本地 SCU-SCP C-ECHO 测试
- [ ] 实现 TestClientServerCStore - C-STORE 往返测试
- [ ] 实现 TestClientServerCFind - C-FIND 查询测试
- [ ] 实现 TestAssociationReject - 拒绝场景测试
- [ ] 实现 TestAbortAndRelease - 异常处理测试
- [ ] 实现并发测试 (多个并发连接)
- [ ] 实现性能基准测试 (吞吐量、延迟)

#### 9.9.2 真实 PACS 互操作性测试 (手动)
- [ ] 测试与 Orthanc 的互操作性
- [ ] 测试与 dcm4che 的互操作性
- [ ] 测试与商业 PACS 的互操作性
- [ ] 记录测试结果和兼容性矩阵

#### 9.9.3 示例程序
- [ ] 创建 examples/dicom_echo_scu - C-ECHO 客户端
  ```bash
  dicom_echo_scu -aec PACS_SERVER -aet MY_SCU 192.168.1.100 104
  ```
- [ ] 创建 examples/dicom_store_scu - C-STORE 客户端
  ```bash
  dicom_store_scu -aec PACS_SERVER file1.dcm file2.dcm
  ```
- [ ] 创建 examples/dicom_find_scu - C-FIND 查询客户端
- [ ] 创建 examples/dicom_storage_scp - C-STORE SCP 服务器
  ```bash
  dicom_storage_scp -aet MY_SCP -port 11112 -dir ./received
  ```

#### 9.9.4 文档
- [ ] 编写网络功能 README (pkg/network/README.md)
- [ ] 编写 API 使用指南
- [ ] 编写故障排查指南 (常见错误和解决方案)
- [ ] 添加网络抓包示例 (Wireshark DICOM 解析)

**预计工作量**: 4-5 天
**优先级**: P2

---

## Phase 9 总结

**包结构** (8 个新包):
```
pkg/network/
├── status/         # 状态码 (P0) ✅
├── pdu/            # PDU 层 (P0) ✅
├── association/    # Association 管理 (P0) ✅
├── dimse/          # DIMSE 消息 (P1) ✅
├── transport/      # 网络抽象 (P1) ⏳
├── service/        # 核心服务 (P1) ⏳
├── client/         # SCU 客户端 (P2) ⏳
└── server/         # SCP 服务器 (P2) ⏳
```

**开发顺序** (自底向上):
1. **P0 阶段 (基础) ✅ 完成**:
   - ✅ 9.1 基础类型和常量 (1 天) - 13个测试
   - ✅ 9.2 PDU 层 (4 天) - 72个测试
   - ✅ 9.3 Association 管理 (2 天) - 14个测试
   - **总计**: 7 天，99个测试用例全部通过

2. **P1 阶段 (核心) - 部分完成**:
   - ✅ 9.4 DIMSE 消息 (3 天) - 31个测试 + 3个示例
   - ⏳ 9.5 网络传输 (1-2 天)
   - ⏳ 9.6 核心服务 (7-8 天)
   - **当前进度**: 1/3 完成

3. **P2 阶段 (应用) - 待开始**:
   - ⏳ 9.7 Client (3-4 天)
   - ⏳ 9.8 Server (3-4 天)
   - ⏳ 9.9 集成测试和示例 (4-5 天)

**已完成工作量**: 10 天
**剩余预计工作量**: 20-28 天
**总预计工作量**: 30-38 天 (6-8 周)

**里程碑**:
- ✅ Milestone 1: P0 完成 (2025-11-07) - 可以编解码所有 PDU 类型
- ⏳ Milestone 2: P1 进行中 - 已完成 DIMSE 消息层
- ⏳ Milestone 3: P2 待开始 - 功能完整的 SCU/SCP 实现

**当前测试统计** (Phase 9):
- Status: 20个测试 ✅
- PDU: 72个测试 ✅
- Association: 14个测试 ✅
- DIMSE: 31个测试 + 3个示例 ✅
- **总计**: 140个测试用例，全部通过

**风险和挑战**:
- PDU 编解码复杂，需要仔细处理字节对齐
- 状态机管理需要考虑并发安全
- 协议兼容性测试需要多个 PACS 系统
- 性能优化可能需要额外时间

---

## 第十一阶段：图像处理 (可选)

**注**: 根据项目需求决定是否实现

### 11.1 DicomPixelData
**参考**: `fo-dicom-code/Imaging/DicomPixelData.cs`

- [ ] 实现像素数据提取
- [ ] 支持多帧图像
- [ ] 支持不同的 Photometric Interpretation
- [ ] 编写单元测试

**预计工作量**: 4-5 天

### 11.2 图像解码器
**参考**: `fo-dicom-code/Imaging/Codec/`

- [ ] 实现 RLE 解码器
- [ ] 集成 JPEG 解码器 (可使用第三方库)
- [ ] 集成 JPEG-LS 解码器
- [ ] 集成 JPEG 2000 解码器
- [ ] 编写单元测试

**预计工作量**: 7-10 天

### 11.3 图像渲染
**参考**: `fo-dicom-code/Imaging/Render/`

- [ ] 实现 VOI LUT (Window/Level)
- [ ] 实现 Modality LUT
- [ ] 实现颜色空间转换
- [ ] 实现图像导出 (PNG, JPEG)
- [ ] 编写单元测试

**预计工作量**: 5-7 天

---

## 第十二阶段：性能优化和文档

### 12.1 性能优化
- [ ] 添加性能基准测试 (benchmarks)
- [ ] 优化内存分配 (使用 sync.Pool)
- [ ] 优化大文件读取
- [ ] Profile 和优化热点代码
- [ ] 添加并发支持 (适当的地方)

**预计工作量**: 5-7 天

### 12.2 文档完善
- [ ] 完善 README.md
  - [ ] 功能列表
  - [ ] 快速开始指南
  - [ ] API 文档链接
- [ ] 编写 GoDoc 注释 (所有公开 API)
- [ ] 创建使用指南文档
- [ ] 创建迁移指南 (从 fo-dicom 到 go-dicom)
- [ ] 添加架构设计文档

**预计工作量**: 4-5 天

### 12.3 完整性测试
- [ ] 使用真实 DICOM 文件进行测试
- [ ] 测试各种 Modality (CT, MR, US, CR, DX 等)
- [ ] 测试各种传输语法
- [ ] 测试边界情况和错误处理
- [ ] 代码覆盖率 > 80%

**预计工作量**: 5-7 天

---

## 里程碑

### 里程碑 1: 基础功能 (预计 3-4 周)
- 项目初始化
- 核心数据类型 (Tag, VR, Element, Dataset)
- Buffer 抽象层
- 传输语法支持

### 里程碑 2: 文件 I/O (预计 3-4 周)
- DICOM 文件读取
- DICOM 文件写入
- 基本命令行工具

### 里程碑 3: 高级功能 (预计 2-3 周)
- JSON/XML 序列化
- 匿名化工具
- 示例和文档

### 里程碑 4: 网络功能 (预计 6-8 周)
- PDU 和 Association
- DIMSE 消息
- DicomClient 和 DicomServer
- 集成测试

### 里程碑 5: 图像处理 (可选，预计 4-5 周)
- 像素数据处理
- 图像解码
- 图像渲染

### 里程碑 6: 优化和发布 (预计 2-3 周)
- 性能优化
- 文档完善
- 发布 v1.0

---

## 注意事项

1. **增量开发**: 每个功能模块完成后都要编写测试，确保质量
2. **参考测试**: fo-dicom 有完善的测试套件，可以参考其测试用例
3. **DICOM 标准**: 遇到不确定的地方，查阅 DICOM 标准文档
4. **Go 惯例**: 遵循 Go 语言的最佳实践和代码风格
5. **依赖管理**: 尽量减少外部依赖，必要时选择维护良好的库
6. **错误处理**: Go 风格的错误处理，提供清晰的错误信息
7. **并发安全**: 考虑并发场景，必要时添加同步机制
8. **向后兼容**: API 设计时考虑未来扩展性

---

## 项目配置更新

### 模块名称
- [x] 模块名称从 `github.com/yourusername/go-dicom` 更新为 `github.com/cocosip/go-dicom`
- [x] 所有内部导入路径已更新 (23个文件)
- [x] go.mod 已更新

### 代码生成工具
- [x] 修复 tools 目录结构
  - [x] `tools/generate_tags/` - Tag 常量生成工具
  - [x] `tools/generate_uids/` - UID 常量生成工具
- [x] 解决多个 main() 函数冲突问题

---

## 当前进度

**总体进度**: ~65% (第一至第八阶段完成，第九阶段 P0+P1 部分完成)

**当前阶段**:
- ✅ 第一至第八阶段完成 - 核心 DICOM 功能 + 命令行工具
- ✅ 第九阶段 P0 完成 - PDU 层、Association 管理、基础类型
- ✅ 第九阶段 P1 DIMSE 层完成 - 所有 C-* 和 N-* 操作
- ⏳ 第九阶段 P1 待完成 - 网络传输层、核心服务层
- ⏳ 第九阶段 P2 待开始 - Client/Server 实现

**Phase 1-3 Summary** (基础设施):
- ✅ 项目初始化和工具配置
- ✅ 核心数据类型 (8个包)
- ✅ Buffer 抽象层 (7种 Buffer 类型)

**Phase 4 Summary** (数据结构):
- ✅ DicomElement - 多种类型元素支持
- ✅ DicomSequence - 序列和嵌套数据集
- ✅ DicomDataset - 完整数据集操作

**Phase 5 Summary** (传输和编码):
- ✅ TransferSyntax - 15+ 标准传输语法
- ✅ UID - 1965个标准UID
- ✅ Charset - 30+ 字符集支持

**Phase 6 Summary** (文件 I/O):
- ✅ **Parser** - 完整的 DICOM 文件解析
  - 支持显式/隐式 VR
  - 支持序列和嵌套结构
  - **ReadOption**: Default/ReadLargeOnDemand/SkipLargeTags/ReadAll
  - **FileFormat**: DICOM3/DICOM3NoPreamble 等格式检测
  - **大数据处理**: 可配置阈值和处理策略
  - Tag.DictionaryEntry() 字典查询
- ✅ **Writer** - 完整的 DICOM 文件写入
  - 自动生成 File Meta Information
  - 支持显式/隐式 VR
  - **DicomWriteOptions**:
    - ExplicitLengthSequences/Items
    - KeepGroupLengths
    - LargeObjectSize
  - Group length 自动过滤
  - 显式/未定义长度序列

**已完成的包** (15个核心包):
1. `pkg/dicom/vr` - Value Representation (35种标准VR)
2. `pkg/dicom/vm` - Value Multiplicity (15种标准VM)
3. `pkg/dicom/tag` - DICOM Tags (5338个标准Tag常量 + DictionaryEntry)
4. `pkg/dicom/dict` - DICOM Dictionary (支持全局Default字典)
5. `pkg/dicom/uid` - DICOM UIDs (1965个UID常量 + go-dicom实现标识符)
6. `pkg/dicom/endian` - Endian 字节序处理
7. `pkg/dicom/transfer` - Transfer Syntax (15+ 标准传输语法)
8. `pkg/dicom/charset` - Character Set Encoding (30+ 字符集)
9. `pkg/io/buffer` - ByteBuffer 抽象层 (7种 Buffer 类型)
10. `pkg/dicom/element` - DICOM Elements (多种类型)
11. `pkg/dicom/dataset` - Dataset, Sequence, 和 FileMetaInformation (Group 0x0002)
12. `pkg/dicom/parser` - 文件解析 (完整 ReadOption 支持)
13. `pkg/dicom/writer` - 文件写入 (完整 WriteOption 支持)
14. `pkg/dicom/serialization` - JSON 和 XML 序列化
15. `pkg/dicom/anonymizer` - DICOM 匿名化 (基于 DICOM PS 3.15)

**功能对比 fo-dicom**:
- ✅ DicomTag → Tag (完整实现 + DictionaryEntry)
- ✅ DicomVR → VR (完整实现)
- ✅ DicomVM → VM (完整实现)
- ✅ DicomDictionary → Dictionary (完整实现 + Default字典)
- ✅ DicomUID → UID (完整实现)
- ✅ DicomTransferSyntax → TransferSyntax (完整实现)
- ✅ DicomEncoding → Charset (完整实现)
- ✅ DicomElement → Element (完整实现)
- ✅ DicomDataset → Dataset (完整实现)
- ✅ DicomFile (ReadOption) → Parser (ReadOption) ✓
- ✅ DicomFile (WriteOption) → Writer (WriteOption) ✓
- ✅ FileReadOption → ReadOption (完整实现)
- ✅ DicomWriteOptions → WriteOption (完整实现)
- ✅ DicomFileFormat → FileFormat (完整实现)
- ✅ DicomFileMetaInformation → FileMetaInformation (完整实现) ✓

**命令行工具** (3个):
1. `dicominfo` - 显示 DICOM 文件信息
2. `dicomdump` - 转储所有 DICOM 标签
3. `dicom2json` - 转换为 JSON 格式

**示例程序** (4个):
1. `examples/read_dicom` - 读取 DICOM 文件
2. `examples/write_dicom` - 创建和写入 DICOM 文件
3. `examples/anonymize` - 匿名化 DICOM 文件
4. `examples/json_conversion` - JSON 序列化和反序列化

**总计**:
- ✅ 所有测试通过 (248+ 测试函数)
- ✅ 代码总量约 40,000 行 (含生成代码)
- ✅ 核心功能完整，可用于生产环境
- ✅ 提供 3 个命令行工具
- ✅ 提供 4 个示例程序

**下一步**: Phase 9 - 网络功能 (DIMSE/DICOM 网络) [可选]

**Phase 9 Network 进度** (网络功能):
- ✅ Phase 9.1: 基础类型和常量 (Status, CommandField, Priority) - 20个测试
- ✅ Phase 9.2: PDU 层 (7种 PDU 类型) - 72个测试
- ✅ Phase 9.3: Association 管理 - 14个测试
- ✅ Phase 9.4: DIMSE 消息层 (所有 C-* 和 N-* 操作) - 102个测试 + 3个示例
  - MessageIDGenerator 自动生成机制
  - 简化 API (无需手动传入 messageID)
  - C-ECHO, C-STORE, C-FIND 完成
  - N-EVENT-REPORT, N-GET, N-SET, N-ACTION, N-CREATE, N-DELETE 完成
  - C-GET, C-MOVE 待后续实现
- ⏳ Phase 9.5-9.9: 待实现

**最近更新**: 2025-11-07
- ✅ 完成 Phase 9.4 DIMSE 消息层 (包括所有 N-* 操作)
  - 实现 C-ECHO (验证连通性)
  - 实现 C-STORE (存储图像)
  - 实现 C-FIND (查询) - Patient Root / Study Root
  - 实现 N-EVENT-REPORT (事件报告)
  - 实现 N-GET (获取属性)
  - 实现 N-SET (设置属性)
  - 实现 N-ACTION (执行操作)
  - 实现 N-CREATE (创建对象)
  - 实现 N-DELETE (删除对象)
  - 实现 MessageIDGenerator (线程安全，自动分配，支持所有操作类型)
  - 102个单元测试 + 3个示例程序全部通过
- ✅ 完成 Phase 9.3 Association 管理
  - PresentationContext 协商
  - ExtendedNegotiation 支持
  - UserIdentity 5种认证类型
  - 14个单元测试全部通过
- ✅ 完成 Phase 9.2 PDU 层 (7种 PDU)
  - A-ASSOCIATE-RQ/AC/RJ
  - P-DATA-TF (分片/重组)
  - A-RELEASE-RQ/RP
  - A-ABORT
  - 72个单元测试全部通过
- ✅ 完成 Phase 9.1 基础类型
  - Status (800+ 状态码)
  - CommandField (23 个命令)
  - Priority (Low/Medium/High)
  - 20个单元测试全部通过
