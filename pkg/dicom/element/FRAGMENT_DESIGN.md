# Fragment Sequence Design

## Overview

Fragment Sequence 是 DICOM 中用于存储封装（encapsulated）像素数据的特殊结构，主要用于压缩图像格式（JPEG, JPEG 2000, JPEG-LS, RLE等）。

## DICOM Fragment Sequence 结构

```
PixelData Tag (7FE0,0010) + VR (OB/OW) + Length (FFFFFFFF)
├─ Item Tag (FFFE,E000) + Length          # Offset Table (第一个 Item)
│  └─ Data: [offset1, offset2, ...]       # 每帧的字节偏移量（uint32数组）
├─ Item Tag (FFFE,E000) + Length          # Fragment 1
│  └─ Data: compressed frame data
├─ Item Tag (FFFE,E000) + Length          # Fragment 2
│  └─ Data: compressed frame data
├─ ...
└─ Sequence Delimitation Item (FFFE,E0DD) + Length (00000000)
```

## Key Points

1. **Undefined Length**: Fragment Sequence 总是使用 undefined length (0xFFFFFFFF)
2. **Offset Table**:
   - 第一个 Item 是 Offset Table
   - 可以为空（length=0）表示单帧图像或未使用 offset table
   - 非空时包含 uint32 数组，每个值是对应帧相对于第一个 fragment 的字节偏移
3. **Fragments**:
   - 每个 fragment 包含一帧或一帧的一部分
   - 多个 fragments 可以组成一帧
4. **Delimitation**: 必须以 Sequence Delimitation Item 结束

## Implementation

### FragmentSequence struct
```go
type FragmentSequence struct {
    tag         *tag.Tag
    vr          *vr.VR
    offsetTable []uint32            // Offset table entries
    fragments   []buffer.ByteBuffer // Fragment buffers
}
```

### Parser Changes

需要在 `parser.go` 中添加：

1. **检测 Fragment Sequence**:
   ```go
   // 当遇到 OB/OW 元素且 length == 0xFFFFFFFF 时，表示这是 Fragment Sequence
   if (vrValue == vr.OB || vrValue == vr.OW) && length == 0xFFFFFFFF {
       return p.readFragmentSequence(t, vrValue)
   }
   ```

2. **读取 Fragment Sequence**:
   ```go
   func (p *parseContext) readFragmentSequence(t *tag.Tag, vrValue *vr.VR) (element.Element, error) {
       fs := element.NewFragmentSequence(t, vrValue)

       // Read first item (offset table)
       // Read subsequent items (fragments)
       // Until we hit Sequence Delimitation Item (FFFE,E0DD)

       return fs, nil
   }
   ```

### Writer Changes

需要在 `writer.go` 中添加：

1. **类型检测**:
   ```go
   if fs, ok := elem.(*element.FragmentSequence); ok {
       return w.writeFragmentSequence(fs)
   }
   ```

2. **写入 Fragment Sequence**:
   ```go
   func (w *Writer) writeFragmentSequence(fs *element.FragmentSequence) error {
       // Write tag + VR + undefined length (FFFFFFFF)
       // Write offset table item
       // Write fragment items
       // Write sequence delimitation item
   }
   ```

## Testing Strategy

1. **Unit Tests**: Test FragmentSequence creation and manipulation
2. **Parser Tests**: Test reading encapsulated pixel data
3. **Writer Tests**: Test writing fragment sequences
4. **Round-trip Tests**: Read → Write → Read should preserve data
5. **Real DICOM Files**: Test with actual compressed DICOM images

## Transfer Syntaxes Requiring Fragment Sequence

- JPEG Baseline (1.2.840.10008.1.2.4.50)
- JPEG Extended (1.2.840.10008.1.2.4.51)
- JPEG Lossless (1.2.840.10008.1.2.4.57, 1.2.840.10008.1.2.4.70)
- JPEG 2000 (1.2.840.10008.1.2.4.90, 1.2.840.10008.1.2.4.91)
- JPEG-LS (1.2.840.10008.1.2.4.80, 1.2.840.10008.1.2.4.81)
- RLE (1.2.840.10008.1.2.5)

## References

- DICOM Standard Part 5: Data Structures and Encoding
- DICOM Standard Part 3: Information Object Definitions
- fo-dicom: DicomFragmentSequence.cs, DicomReader.cs (ParseFragmentSequence)
