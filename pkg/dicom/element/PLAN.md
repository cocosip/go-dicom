# DICOM Element实现计划

根据DICOM标准和fo-dicom库，Element类型按照Value Representation (VR)分类。

## Element类型分类

### 1. 字符串类型 (String-based)

| VR | 名称 | 描述 | 状态 |
|----|------|------|------|
| AE | Application Entity | 应用实体标题 | ✅ 已实现(String) |
| AS | Age String | 年龄字符串 | ✅ 已实现(String) |
| CS | Code String | 代码字符串 | ✅ 已实现(String) |
| DA | Date | 日期 | ✅ 已实现(String) |
| DS | Decimal String | 十进制字符串 | ✅ 已实现(String) |
| DT | Date Time | 日期时间 | ✅ 已实现(String) |
| IS | Integer String | 整数字符串 | ✅ 已实现(String) |
| LO | Long String | 长字符串 | ✅ 已实现(String) |
| LT | Long Text | 长文本 | ✅ 已实现(String) |
| PN | Person Name | 人名 | ✅ 已实现(String) |
| SH | Short String | 短字符串 | ✅ 已实现(String) |
| ST | Short Text | 短文本 | ✅ 已实现(String) |
| TM | Time | 时间 | ✅ 已实现(String) |
| UC | Unlimited Characters | 无限字符 | ✅ 已实现(String) |
| UI | Unique Identifier | 唯一标识符(UID) | ✅ 已实现(String) |
| UR | URI/URL | 统一资源标识符 | ✅ 已实现(String) |
| UT | Unlimited Text | 无限文本 | ✅ 已实现(String) |

**实现文件**: `string.go`
**测试文件**: `string_test.go`

### 2. 数值类型 (Numeric)

| VR | 名称 | 描述 | 大小 | 类型 | 状态 |
|----|------|------|------|------|------|
| FL | Floating Point Single | 单精度浮点 | 4 bytes | float32 | ✅ 已实现 |
| FD | Floating Point Double | 双精度浮点 | 8 bytes | float64 | ✅ 已实现 |
| SL | Signed Long | 有符号长整型 | 4 bytes | int32 | ✅ 已实现 |
| SS | Signed Short | 有符号短整型 | 2 bytes | int16 | ✅ 已实现 |
| UL | Unsigned Long | 无符号长整型 | 4 bytes | uint32 | ✅ 已实现 |
| US | Unsigned Short | 无符号短整型 | 2 bytes | uint16 | ✅ 已实现 |
| SV | Signed 64-bit Very Long | 有符号64位 | 8 bytes | int64 | ✅ 已实现 |
| UV | Unsigned 64-bit Very Long | 无符号64位 | 8 bytes | uint64 | ✅ 已实现 |

**实现文件**: `numeric.go`
**测试文件**: `numeric_test.go`

### 3. 二进制类型 (Binary/Other)

| VR | 名称 | 描述 | 状态 |
|----|------|------|------|
| OB | Other Byte | 字节序列 | ✅ 已实现 |
| OD | Other Double | 双精度浮点序列 | ✅ 已实现 |
| OF | Other Float | 单精度浮点序列 | ✅ 已实现 |
| OL | Other Long | 32位整数序列 | ✅ 已实现 |
| OV | Other 64-bit Very Long | 64位整数序列 | ✅ 已实现 |
| OW | Other Word | 16位字序列 | ✅ 已实现 |
| UN | Unknown | 未知数据 | ✅ 已实现 |

**实现文件**: `binary.go`
**测试覆盖**: `value_text_test.go` 及 parser/serialization 集成测试

### 4. 特殊类型

| VR | 名称 | 描述 | 状态 |
|----|------|------|------|
| SQ | Sequence of Items | 序列(包含子项) | ✅ 已实现 (`pkg/dicom/dataset`) |
| AT | Attribute Tag | 属性标签 | ✅ 已实现 |

**实现文件**: `../dataset/sequence.go`, `attribute_tag.go`
**测试文件**: `../dataset/sequence_test.go`, `../dataset/sequence_integration_test.go`, `value_text_test.go`

`Sequence` 包含 `*Dataset`，因此实现位于 `dataset` 包以避免 `element` 与
`dataset` 之间形成循环依赖；它仍完整实现 `element.Element` 接口。

## 实现优先级

### 高优先级 (Phase 4.1) ✅
- [x] String 类型 (所有字符串VR)
- [x] 基础数值类型 (US, UL, SS, SL, FL, FD)

### 中优先级 (Phase 4.2) ✅
- [x] 64位数值类型 (SV, UV)
- [x] 属性标签类型 (AT)
- [x] 二进制类型 (OB, OW, UN)
- [x] 特殊字符串类型 (PN, DA, TM, DT, DS, IS)

### 低优先级 (Phase 4.3) ✅
- [x] 其他二进制类型 (OD, OF, OL, OV)
- [x] 序列类型 (SQ) - 位于 `pkg/dicom/dataset`

## fo-dicom参考

在fo-dicom中，Element类型的实现：
- `DicomStringElement` - 所有字符串VR
- `DicomIntegerString` - IS (但也可以用StringElement)
- `DicomDecimalString` - DS (但也可以用StringElement)
- `DicomPersonName` - PN的特殊处理
- `DicomDate`, `DicomTime`, `DicomDateTime` - 日期时间特殊处理
- `DicomUnsignedShort` - US
- `DicomUnsignedLong` - UL
- `DicomSignedShort` - SS
- `DicomSignedLong` - SL
- `DicomFloatingPointSingle` - FL
- `DicomFloatingPointDouble` - FD
- `DicomOtherByte` - OB
- `DicomOtherWord` - OW
- `DicomOtherDouble` - OD
- `DicomOtherFloat` - OF
- `DicomOtherLong` - OL
- `DicomOtherVeryLong` - OV
- `DicomUnknown` - UN
- `DicomAttributeTag` - AT
- `DicomSequence` - SQ

## 当前统计

- **已实现**: 本计划列出的所有 Element 类型
  - 字符串类型: 17个VR (包括PN特殊处理)
  - 数值类型: 8种 (US/UL/SS/SL/FL/FD/SV/UV)
  - 日期时间类型: 3种 (DA/TM/DT)
  - 数字字符串类型: 2种 (DS/IS)
  - 二进制类型: 7种 (OB/OW/OD/OF/OL/OV/UN)
  - 属性标签类型: 1种 (AT)
  - 序列类型: 1种 (SQ，位于 `pkg/dicom/dataset`)
- **读写集成**: SQ 已接入 parser、writer、JSON 和 XML；parser 与 XML 对二进制数值及二进制 VR 保留具体 Element 类型
- **待实现**: 无

## 下一步

Phase 4 和本计划依赖的 Phase 5 能力均已完成。后续变更应继续保证：

- 新增或修改 VR 时同步更新 parser、writer、JSON/XML 序列化及回归测试。
- Sequence 的嵌套、递归验证、克隆和显式/未定义长度编码保持一致。
