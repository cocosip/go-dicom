# DICOM Dataset Package

`pkg/dicom/dataset` 包实现了DICOM数据集（Dataset）和序列（Sequence）。

## 概述

Dataset是DICOM数据的核心容器，用于：
- 存储和管理Element集合
- 按Tag进行快速查找
- 支持嵌套结构（通过Sequence）
- 用于DICOM文件、网络消息、序列项

## 核心类型

### 1. Dataset

Dataset是Element的集合，内部使用map进行存储，保证按Tag排序输出。

**基本操作：**

```go
// 创建Dataset
ds := dataset.New()

// 添加元素
ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))

// 获取元素
elem, exists := ds.Get(tag.PatientName)
if exists {
    str := elem.(*element.String)
    name := str.GetString() // "Doe^John"
}

// 检查是否存在
if ds.Contains(tag.PatientName) {
    // ...
}

// 删除元素
ds.Remove(tag.PatientName)

// 获取所有元素（按Tag排序）
elements := ds.Elements()

// 获取所有Tag
tags := ds.Tags()

// 清空Dataset
ds.Clear()
```

**高级操作：**

```go
// 克隆Dataset
clone := ds.Clone()

// 合并Dataset
ds1.Merge(ds2, true) // true = 覆盖已存在的元素

// 过滤元素
stringElements := ds.Filter(func(elem element.Element) bool {
    _, ok := elem.(*element.String)
    return ok
})

// 统计
count := ds.Count()
empty := ds.IsEmpty()
```

### 2. Accessor方法

Dataset提供了便捷的类型安全访问器：

```go
// 字符串访问
name, exists := ds.GetString(tag.PatientName)
strs, exists := ds.GetStrings(tag.ImageType)

// 数值访问
rows, err := ds.GetUInt16(tag.Rows, 0)
rowsAll, err := ds.GetUInt16s(tag.Rows)

length, err := ds.GetUInt32(tag.FileMetaInformationGroupLength, 0)
value, err := ds.GetInt16(tag.SmallestImagePixelValue, 0)
value, err := ds.GetInt32(tag.ReferencePixelX0, 0)

// 浮点数访问
pos, err := ds.GetFloat32(tag.ImagePositionPatient, 0)
slope, err := ds.GetFloat64(tag.RealWorldValueSlope, 0)

// 序列访问
seq, err := ds.GetSequence(tag.ReferencedImageSequence)
```

**Try* 方法** - 不抛错误，失败返回零值：

```go
name := ds.TryGetString(tag.PatientName) // 不存在返回 ""
rows := ds.TryGetUInt16(tag.Rows, 0)      // 不存在返回 0
length := ds.TryGetUInt32(tag.FileMetaInformationGroupLength, 0)
```

### 3. Sequence

Sequence（VR=SQ）表示嵌套的Dataset集合。

**基本操作：**

```go
// 创建Sequence
seq := dataset.NewSequence(tag.ReferencedImageSequence)

// 添加Item (每个Item是一个Dataset)
item1 := dataset.New()
item1.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3"}))
item1.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4"}))
seq.AddItem(item1)

item2 := dataset.New()
item2.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.5"}))
seq.AddItem(item2)

// 获取Item
item := seq.GetItem(0)
if item != nil {
    uid, _ := item.GetString(tag.StudyInstanceUID)
}

// 获取所有Items
items := seq.GetItems()

// 删除Item
seq.RemoveItem(1)

// 清空Sequence
seq.Clear()

// 统计
count := seq.Count()
empty := seq.IsEmpty()
```

**将Sequence添加到Dataset：**

```go
ds := dataset.New()
ds.Add(seq) // Sequence实现了Element接口

// 从Dataset获取Sequence
seq, err := ds.GetSequence(tag.ReferencedImageSequence)
if err == nil {
    for i := 0; i < seq.Count(); i++ {
        item := seq.GetItem(i)
        // 处理item...
    }
}
```

**Sequence高级操作：**

```go
// 克隆Sequence
clone := seq.Clone()

// 过滤Items
ctItems := seq.Filter(func(ds *dataset.Dataset) bool {
    modality, _ := ds.GetString(tag.Modality)
    return modality == "CT"
})
```

## 文件清单

| 文件 | 说明 | 行数 |
|------|------|------|
| `doc.go` | 包文档 | ~50 |
| `dataset.go` | Dataset核心实现 | ~220 |
| `accessors.go` | 类型安全访问器 | ~200 |
| `sequence.go` | Sequence实现 | ~160 |
| `dataset_test.go` | Dataset测试 | ~260 |
| `accessors_test.go` | Accessor测试 | ~195 |
| `sequence_test.go` | Sequence测试 | ~200 |
| **总计** |  | **~1,285行** |

## 测试覆盖

- **测试文件**: 3个
- **测试函数**: 25个
- **测试用例**: 60+个子测试
- **测试通过率**: 100% ✅

```bash
# 运行所有测试
cd pkg/dicom/dataset
go test -v

# 运行特定测试
go test -v -run TestDataset
go test -v -run TestSequence
go test -v -run TestAccessors
```

## 使用示例

### 示例1: 创建Patient信息Dataset

```go
package main

import (
    "fmt"
    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func main() {
    ds := dataset.New()

    // 添加Patient信息
    ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
    ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
    ds.Add(element.NewString(tag.PatientBirthDate, vr.DA, []string{"19800101"}))
    ds.Add(element.NewString(tag.PatientSex, vr.CS, []string{"M"}))

    // 读取信息
    name := ds.TryGetString(tag.PatientName)
    id := ds.TryGetString(tag.PatientID)

    fmt.Printf("Patient: %s (ID: %s)\n", name, id)
    fmt.Printf("Total elements: %d\n", ds.Count())
}
```

### 示例2: 使用Sequence

```go
// 创建Referenced Image Sequence
seq := dataset.NewSequence(tag.ReferencedImageSequence)

// 添加多个引用图像
for i := 0; i < 3; i++ {
    item := dataset.New()
    item.Add(element.NewString(tag.ReferencedSOPClassUID, vr.UI,
        []string{"1.2.840.10008.5.1.4.1.1.2"}))
    item.Add(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI,
        []string{fmt.Sprintf("1.2.3.%d", i)}))
    item.Add(element.NewUnsignedShort(tag.ReferencedFrameNumber, []uint16{uint16(i + 1)}))
    seq.AddItem(item)
}

// 添加到主Dataset
mainDS := dataset.New()
mainDS.Add(seq)

// 遍历Sequence
if refSeq, err := mainDS.GetSequence(tag.ReferencedImageSequence); err == nil {
    fmt.Printf("Referenced images: %d\n", refSeq.Count())
    for i := 0; i < refSeq.Count(); i++ {
        item := refSeq.GetItem(i)
        uid, _ := item.GetString(tag.ReferencedSOPInstanceUID)
        fmt.Printf("  Image %d: %s\n", i+1, uid)
    }
}
```

### 示例3: Dataset合并和过滤

```go
// 创建两个Dataset
ds1 := dataset.New()
ds1.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
ds1.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250106"}))

ds2 := dataset.New()
ds2.Add(element.NewString(tag.PatientName, vr.PN, []string{"Smith^Jane"}))
ds2.Add(element.NewString(tag.StudyTime, vr.TM, []string{"120000"}))

// 合并（不覆盖）
merged := ds1.Clone()
merged.Merge(ds2, false) // PatientName保持 "Doe^John"

// 过滤只保留字符串元素
filtered := merged.Filter(func(elem element.Element) bool {
    _, ok := elem.(*element.String)
    return ok
})
```

## 设计原则

1. **性能优化**: 使用map存储，O(1)查找复杂度
2. **类型安全**: 提供类型安全的访问器方法
3. **一致性**: Elements()和Tags()保证排序输出
4. **灵活性**: 支持过滤、合并、克隆等高级操作
5. **嵌套支持**: Sequence实现完整的嵌套Dataset支持

## 与Element包的关系

- Dataset存储Element接口类型
- Sequence也是Element的一种（VR=SQ）
- 可以在Dataset中存储任何Element类型（包括Sequence）

## 相关包

- `pkg/dicom/element` - DICOM元素实现
- `pkg/dicom/tag` - DICOM标签
- `pkg/dicom/vr` - 值表示

## 下一步

- [ ] 实现DICOM文件读写（需要Parser和Writer）
- [ ] 实现网络传输（需要DIMSE协议）
- [ ] 添加数据验证功能
- [ ] 性能优化和基准测试
