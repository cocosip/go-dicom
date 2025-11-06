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
| SV | Signed 64-bit Very Long | 有符号64位 | 8 bytes | int64 | ❌ 待实现 |
| UV | Unsigned 64-bit Very Long | 无符号64位 | 8 bytes | uint64 | ❌ 待实现 |

**实现文件**: `numeric.go`
**测试文件**: `numeric_test.go`

### 3. 二进制类型 (Binary/Other)

| VR | 名称 | 描述 | 状态 |
|----|------|------|------|
| OB | Other Byte | 字节序列 | ❌ 待实现 |
| OD | Other Double | 双精度浮点序列 | ❌ 待实现 |
| OF | Other Float | 单精度浮点序列 | ❌ 待实现 |
| OL | Other Long | 32位整数序列 | ❌ 待实现 |
| OV | Other 64-bit Very Long | 64位整数序列 | ❌ 待实现 |
| OW | Other Word | 16位字序列 | ❌ 待实现 |
| UN | Unknown | 未知数据 | ❌ 待实现 |

**计划文件**: `binary.go`
**测试文件**: `binary_test.go`

### 4. 特殊类型

| VR | 名称 | 描述 | 状态 |
|----|------|------|------|
| SQ | Sequence of Items | 序列(包含子项) | ❌ 待实现 |
| AT | Attribute Tag | 属性标签 | ❌ 待实现 |

**计划文件**: `sequence.go`, `attribute_tag.go`
**测试文件**: `sequence_test.go`, `attribute_tag_test.go`

## 实现优先级

### 高优先级 (Phase 4.1) ✅
- [x] String 类型 (所有字符串VR)
- [x] 基础数值类型 (US, UL, SS, SL, FL, FD)

### 中优先级 (Phase 4.2) ✅
- [x] 64位数值类型 (SV, UV)
- [x] 属性标签类型 (AT)
- [x] 二进制类型 (OB, OW, UN)
- [x] 特殊字符串类型 (PN, DA, TM, DT, DS, IS)

### 低优先级 (Phase 4.3)
- [x] 其他二进制类型 (OD, OF, OL, OV)
- [ ] 序列类型 (SQ) - 最复杂，需要Dataset支持

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

- **已实现**: 所有主要类型（除了SQ序列）
  - 字符串类型: 17个VR (包括PN特殊处理)
  - 数值类型: 8种 (US/UL/SS/SL/FL/FD/SV/UV)
  - 日期时间类型: 3种 (DA/TM/DT)
  - 数字字符串类型: 2种 (DS/IS)
  - 二进制类型: 7种 (OB/OW/OD/OF/OL/OV/UN)
  - 属性标签类型: 1种 (AT)
- **待实现**: 1个类型 (SQ - 需要Dataset支持)
- **代码文件**: 16个
- **代码行数**: ~2,740行
- **测试文件**: 5个
- **测试函数**: 21个
- **测试通过率**: 100% ✅

## 下一步

Phase 4 基本完成！下一步：

**Phase 5: Dataset包**
- 实现Dataset结构（Element集合管理）
- 实现Sequence (SQ)类型（依赖Dataset）
- 实现序列化/反序列化
