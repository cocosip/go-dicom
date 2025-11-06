# DICOM Element Package

`pkg/dicom/element` 包实现了DICOM数据元素的各种类型。

## 概述

Element是DICOM数据集的基本构建块，每个Element包含：
- **Tag**: 唯一标识符 (group, element)
- **VR (Value Representation)**: 值表示类型
- **Data**: 二进制数据，存储在ByteBuffer中

## 已实现的Element类型

### 1. 字符串类型 (String)

**文件**: `string.go`

支持所有字符串VR：
- AE (Application Entity)
- AS (Age String)
- CS (Code String)
- DA (Date)
- DS (Decimal String)
- DT (Date Time)
- IS (Integer String)
- LO (Long String)
- LT (Long Text)
- PN (Person Name)
- SH (Short String)
- ST (Short Text)
- TM (Time)
- UC (Unlimited Characters)
- UI (Unique Identifier)
- UR (URI/URL)
- UT (Unlimited Text)

**特点**:
- 支持多值（用反斜杠分隔）
- 支持字符集编码
- 自动去除尾随空格

**示例**:
```go
// 单值字符串
elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})
name := elem.GetString() // "Doe^John"

// 多值字符串
elem := element.NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY", "AXIAL"})
values := elem.GetValues() // ["ORIGINAL", "PRIMARY", "AXIAL"]
```

### 2. 特殊字符串类型

#### 2.1 人名类型 (PersonName)

**文件**: `person_name.go`

PN (Person Name) 的特殊实现，支持人名组件解析：
- FamilyName (姓)
- GivenName (名)
- MiddleName (中间名)
- NamePrefix (前缀，如 Dr., Prof.)
- NameSuffix (后缀，如 Jr., III)

**示例**:
```go
elem := element.NewPersonName(tag.PatientName, []string{"Doe^John^A^Dr.^Jr."})
components, _ := elem.GetComponents(0)
// components.FamilyName = "Doe"
// components.GivenName = "John"
// components.MiddleName = "A"
// components.NamePrefix = "Dr."
// components.NameSuffix = "Jr."

formatted, _ := elem.GetFormattedName(0) // "Dr. John A Doe, Jr."
```

#### 2.2 日期时间类型

**文件**: `date.go`

包含三种日期时间类型：

**Date (DA)** - 日期格式 YYYYMMDD
```go
elem := element.NewDate(tag.StudyDate, []string{"20250106"})
date, _ := elem.GetDate(0) // time.Time

// 从 time.Time 创建
testTime := time.Date(2025, 1, 6, 0, 0, 0, 0, time.UTC)
elem := element.NewDateFromTime(tag.StudyDate, []time.Time{testTime})
```

**Time (TM)** - 时间格式 HHMMSS.FFFFFF
```go
elem := element.NewTime(tag.StudyTime, []string{"123456.789"})
tm, _ := elem.GetTime(0) // time.Time

// 支持多种格式: HH, HHMM, HHMMSS, HHMMSS.FFF, HHMMSS.FFFFFF
elem := element.NewTime(tag.StudyTime, []string{"1234"}) // HHMM格式
```

**DateTime (DT)** - 日期时间格式 YYYYMMDDHHMMSS.FFFFFF&ZZXX
```go
elem := element.NewDateTime(tag.AcquisitionDateTime, []string{"20250106123456"})
dt, _ := elem.GetDateTime(0) // time.Time

// 支持时区
elem := element.NewDateTimeFromTime(tag.AcquisitionDateTime, times)
```

#### 2.3 数字字符串类型

**文件**: `numeric_string.go`

**DecimalString (DS)** - 十进制字符串转浮点数
```go
elem := element.NewDecimalString(tag.SliceThickness, []string{"1.5", "2.0"})
f, _ := elem.GetFloat(0) // 1.5

// 从 float 创建
elem := element.NewDecimalStringFromFloat(tag.SliceThickness, []float64{1.5, 2.0})
floats, _ := elem.GetFloats() // []float64{1.5, 2.0}

// 支持 float32
f32, _ := elem.GetFloat32(0)
```

**IntegerString (IS)** - 整数字符串转整数
```go
elem := element.NewIntegerString(tag.NumberOfFrames, []string{"100"})
i, _ := elem.GetInt(0) // 100

// 从 int 创建
elem := element.NewIntegerStringFromInt(tag.NumberOfFrames, []int{100})
ints, _ := elem.GetInts() // []int{100}

// 支持 int32, int64
i32, _ := elem.GetInt32(0)
i64, _ := elem.GetInt64(0)
```

### 3. 数值类型 (Numeric)

**文件**: `numeric.go`

| 类型 | VR | Go类型 | 大小 | 范围 |
|------|----|----|------|------|
| UnsignedShort | US | uint16 | 2 bytes | 0 ~ 65,535 |
| UnsignedLong | UL | uint32 | 4 bytes | 0 ~ 4,294,967,295 |
| SignedShort | SS | int16 | 2 bytes | -32,768 ~ 32,767 |
| SignedLong | SL | int32 | 4 bytes | -2,147,483,648 ~ 2,147,483,647 |
| Float | FL | float32 | 4 bytes | IEEE 754 单精度 |
| Double | FD | float64 | 8 bytes | IEEE 754 双精度 |
| SignedVeryLong | SV | int64 | 8 bytes | -2^63 ~ 2^63-1 |
| UnsignedVeryLong | UV | uint64 | 8 bytes | 0 ~ 2^64-1 |

**示例**:
```go
// 创建 US 元素
elem := element.NewUnsignedShort(tag.Rows, []uint16{512, 512})
rows, _ := elem.GetValue(0) // 512

// 创建 FD 元素
elem := element.NewDouble(tag.RealWorldValueSlope, []float64{1.234567890123})
slope, _ := elem.GetValue(0) // 1.234567890123
```

### 4. 属性标签类型 (AttributeTag)

**文件**: `attribute_tag.go`

**VR**: AT

存储DICOM标签引用，每个值是一个32位标签(group + element)。

**示例**:
```go
tags := []*tag.Tag{tag.PatientName, tag.StudyDate}
elem := element.NewAttributeTag(tag.OffendingElement, tags)
tag, _ := elem.GetValue(0) // tag.PatientName
```

### 5. 二进制类型 (Binary)

**文件**: `binary.go`

| 类型 | VR | 描述 |
|------|----|----|
| OtherByte | OB | 字节序列 |
| OtherWord | OW | 16位字序列 |
| Unknown | UN | 未知数据 |
| OtherDouble | OD | 64位浮点序列 |
| OtherFloat | OF | 32位浮点序列 |
| OtherLong | OL | 32位整数序列 |
| OtherVeryLong | OV | 64位整数序列 |

**特点**:
- 用于存储像素数据、波形数据等
- Count() 返回 1（整体作为单个值）
- GetData() 返回原始字节数组

**示例**:
```go
// 存储像素数据
data := []byte{0x00, 0xFF, 0x80, 0x7F}
elem := element.NewOtherWord(tag.PixelData, data)
pixelData := elem.GetData() // [0x00, 0xFF, 0x80, 0x7F]
```

## 接口设计

### Element 接口

所有Element类型都实现此接口：

```go
type Element interface {
    Tag() *tag.Tag
    ValueRepresentation() *vr.VR
    Buffer() buffer.ByteBuffer
    Length() uint32
    Count() int
    String() string
    Validate() error
}
```

### 继承结构

```
Element (interface)
  └─ base (struct) - 所有类型的公共基础
      ├─ String
      │   └─ PersonName
      ├─ UnsignedShort
      ├─ UnsignedLong
      ├─ SignedShort
      ├─ SignedLong
      ├─ Float
      ├─ Double
      ├─ SignedVeryLong
      ├─ UnsignedVeryLong
      ├─ AttributeTag
      ├─ OtherByte
      ├─ OtherWord
      ├─ OtherDouble
      ├─ OtherFloat
      ├─ OtherLong
      ├─ OtherVeryLong
      └─ Unknown
```

## 文件清单

| 文件 | 说明 | 行数 |
|------|------|------|
| `doc.go` | 包文档 | ~50 |
| `element.go` | Element接口和base实现 | ~95 |
| `string.go` | 通用字符串类型 | ~165 |
| `person_name.go` | 人名类型 (PN) | ~135 |
| `date.go` | 日期时间类型 (DA/TM/DT) | ~230 |
| `numeric_string.go` | 数字字符串类型 (DS/IS) | ~235 |
| `numeric.go` | 数值类型 (US/UL/SS/SL/FL/FD/SV/UV) | ~605 |
| `attribute_tag.go` | 属性标签类型 (AT) | ~95 |
| `binary.go` | 二进制类型 (OB/OW/OD/OF/OL/OV/UN) | ~275 |
| `element_test.go` | 接口测试 | ~55 |
| `string_test.go` | 字符串测试 | ~110 |
| `date_test.go` | 日期时间测试 | ~125 |
| `numeric_string_test.go` | 数字字符串测试 | ~150 |
| `numeric_test.go` | 数值测试 | ~370 |
| `PLAN.md` | 实现计划 | ~140 |
| `README.md` | 文档 | ~280 |
| **总计** |  | **~2,740行** |

## 测试覆盖

- **测试文件**: 5个
- **测试函数**: 21个
- **测试用例**: 70+个子测试
- **测试覆盖率**: 主要功能全覆盖
- **测试通过率**: 100% ✅

```bash
# 运行所有测试
cd pkg/dicom/element
go test -v

# 运行特定测试
go test -v -run TestString
go test -v -run TestNumeric
```

## 使用示例

### 创建Element

```go
// 字符串
elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})

// 数值
elem := element.NewUnsignedShort(tag.Rows, []uint16{512})

// 二进制
data := []byte{0x01, 0x02, 0x03, 0x04}
elem := element.NewOtherByte(tag.PixelData, data)
```

### 从Buffer创建

```go
buf := buffer.NewMemory([]byte("Patient Name"))
elem := element.NewStringFromBuffer(tag.PatientName, vr.PN, buf, nil)
```

### 访问值

```go
// 字符串
name := elem.GetString()
values := elem.GetValues()

// 数值
value, err := elem.GetValue(0)
values, err := elem.GetValues()

// 二进制
data := elem.GetData()
```

## 未来工作

### 待实现类型

- [ ] Sequence (SQ) - 最复杂，需要Dataset支持才能实现

### 待增强功能

- [ ] VR特定的验证规则
- [ ] 更完善的错误处理
- [ ] 性能优化
- [ ] 更多测试用例

## 相关包

- `pkg/dicom/tag` - DICOM标签
- `pkg/dicom/vr` - 值表示
- `pkg/io/buffer` - 缓冲区接口
- `pkg/dicom/charset` - 字符集编码

## 参考

- DICOM标准 PS3.5: Data Structures and Encoding
- fo-dicom库: https://github.com/fo-dicom/fo-dicom
