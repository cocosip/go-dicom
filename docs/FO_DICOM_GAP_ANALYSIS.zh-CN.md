# fo-dicom 能力差距分析

[English](FO_DICOM_GAP_ANALYSIS.md)

本文档跟踪 `go-dicom` 与参考实现 `fo-dicom` 之间的能力差距。它是一份纳入版本管理的工程待办清单，并不意味着需要在 Go 中复刻每一个 .NET API。

## 审计基线

初次审计于 2026-08-14 针对以下版本完成：

- `go-dicom`: `d5970f342973c0d659c3eab1b7cee8563a7f5dda`
- `fo-dicom`: `7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2`
  (`5.2.6-101-g7ea6d424`)

本次对比以源代码、测试和示例作为证据。README 中的功能列表不作为实现证据。

以下命令在被审计的 `go-dicom` 版本上执行通过：

```powershell
go test ./cmd/... ./examples/... ./pkg/... ./tools/...
```

由于任一仓库都可能在基线之后发生变化，实施具体条目前应重新检查两个仓库。

## 分类

| 状态 | 含义 |
| --- | --- |
| `Open` | `go-dicom` 中不存在等价的领域能力。 |
| `Partial` | 已存在部分层次或值类型，但公共工作流不完整。 |
| `External` | 该能力有意由配套模块提供。 |
| `Not a gap` | 差异属于平台特性，或两个库都不具备该能力。 |
| `Complete` | 已满足验收标准，并记录了当前验证证据。 |

优先级表示实施顺序，而不是工作量估算：

- `P0`: 正确性、互操作性或具有误导性的公共契约
- `P1`: 实现实用对标所需的主要 fo-dicom 领域工作流
- `P2`: 重要的支撑 API 或运维能力
- `P3`: 应在其前置能力完成后实施的大型或专用能力

## 执行摘要

| ID | 优先级 | 状态 | 能力 |
| --- | --- | --- | --- |
| DOC-001 | P0 | Partial | 公共能力声明与实际实现行为一致 |
| NET-001 | P0 | Complete | 支持 TLS 的高层 DICOM 客户端 |
| NET-002 | P0 | Partial | C-STORE 传输语法选择与自动转码 |
| STD-001 | P0 | Partial | 可复现且保持最新的标准生成表 |
| MED-001 | P1 | Open | DICOMDIR 介质目录模型及读写工作流 |
| NET-003 | P1 | Partial | 通过高层客户端执行高级关联协商 |
| NET-004 | P1 | Open | SOP Class Common Extended Negotiation |
| SR-001 | P1 | Partial | 完整的结构化报告值类型和文件工作流 |
| IMG-001 | P1 | Partial | Dataset 驱动的图像渲染管线 |
| CORE-001 | P1 | Partial | Dataset 和 Sequence 递归验证 |
| IMG-002 | P2 | Open | 帧几何、空间变换和插值工具 |
| CORE-002 | P2 | Open | Dataset 遍历器、匹配规则和转换规则 |
| DICT-001 | P2 | Partial | 运行时 XML 字典加载 |
| ANON-001 | P2 | Complete | 完整的自定义匿名化配置加载 |
| PRINT-001 | P2 | Partial | 基于 Dataset 的 DICOM 打印管理模型 |
| OBS-001 | P2 | Open | 结构化网络日志、请求事件和指标钩子 |
| IMG-003 | P3 | Partial | 体数据重建和 MPR |
| MED-002 | P3 | Open | DICOM 文件扫描工作流 |

## 实施进度

截至 2026-08-14：

- **已完成：** NET-001 和 ANON-001。
- **未完成：** DOC-001、NET-002、STD-001、MED-001、NET-003、NET-004、
  SR-001、IMG-001、CORE-001、IMG-002、CORE-002、DICT-001、PRINT-001、
  OBS-001、IMG-003 和 MED-002 继续保持 `Partial` 或 `Open` 状态。
- Phase 0 尚未完成。NET-001 已完成，DOC-001 中与 TLS 相关的部分已经修复；
  其余 Phase 0 工作仍由 DOC-001、NET-002 和 STD-001 跟踪。

## 详细差距

### DOC-001: 公共能力声明

**状态：** `Partial`  
**优先级：** `P0`

在审计基线中，README 的若干声明超出了当时已实现的公共工作流，包括完整的客户端 TLS、完整的 SR 值类型、图像重建、通过客户端进行高级协商，以及打印任务创建。TLS 示例调用了不存在的 `client.WithTLS` 选项；另一个渲染示例向 `NewDicomImage` 传入 Dataset，但该函数实际接收 `*DicomPixelData`。

README 中的标签和 UID 数量也与被审计源码中去重后的标准生成条目数不一致。

2026-08-14 进度：NET-001 已补齐高层客户端 TLS API，并修复 README 中的
TLS 示例。SR、渲染、高级协商、打印、生成条目数量及其他公共能力声明尚未
重新审计或修复，因此 DOC-001 仍为 `Partial`。

**验收标准**

- 每项受检查的能力都有可编译的公共 API 和聚焦测试作为支撑。
- 明确标注部分实现和外部提供的能力。
- 所有 README 代码示例都能在 CI 中编译。
- 自动生成标签和 UID 数量，或者不再声明这些数量。

**建议验证**

- 为 README 代码片段增加编译测试，或将可运行片段迁移到 examples。
- 运行完整 Go 包测试树和文档链接检查。

### NET-001: 高层客户端 TLS

**状态：** `Complete`

**优先级：** `P0`

已于 2026-08-14 完成。高层客户端现在提供
`client.WithTLSConfig(*tls.Config)`。配置为 nil 时保留原有普通 TCP 路径；
配置非 nil 时，在 DICOM Association 协商前使用 `transport.DialTLS`。
`ConnectTimeout` 同时覆盖 TCP 建连和 TLS 握手，调用方取消可以正常传递，
transport 层会在应用默认值前克隆由调用方持有的 TLS 配置。

参考：[fo-dicom DicomClient](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Network/Client/DicomClient.cs)

**验收标准**

- 客户端接受调用方持有的 `*tls.Config`，且不修改该对象。
- 默认仍使用普通 TCP。
- 覆盖 TLS 握手、证书校验、超时、取消和关闭路径。
- README 的 TLS 示例可编译，并使用真实存在的公共选项名称。

**验证证据**

- 聚焦测试覆盖普通 TCP、经过验证的 TLS 握手、完整 TLS Association、
  主机名不匹配、不受信任证书、握手超时、调用方取消，以及多个客户端共享
  同一个 `tls.Config` 的并发场景。
- `go test ./pkg/network/... -count=1` 通过。
- `go test -race ./pkg/network/... -count=1` 通过。
- `go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1` 通过。
- `golangci-lint run` 报告 0 个问题；分析完成后，沙箱环境输出了缓存持久化警告。
- 实现和上述验证证据已提交于 `e8e44ef`。

### NET-002: C-STORE 协商传输语法与转码

**状态：** `Partial`  
**优先级：** `P0`

`go-dicom` 已有编解码器注册表和转码器，但网络发送路径不会调用它们。该路径按 SOP Class 选择表示上下文，并使用被接受的传输语法编码请求。当 Dataset 的像素数据传输语法与被接受的上下文不同时，这种处理并不充分。`CStoreMultiple` 会顺序发送每个 Dataset，并在第一个错误处停止。

参考：[fo-dicom DicomCStoreRequest](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Network/DicomCStoreRequest.cs)

**验收标准**

- 表示上下文选择同时考虑 SOP Class 和源传输语法。
- 被接受时优先使用原始语法。
- 必要时使用已注册的编解码器进行回退转码。
- 没有可直接使用或可转码的已接受语法时，返回明确错误。
- 不修改调用方的 Dataset。
- 批量发送明确定义顺序、部分成功、取消和并发语义。

**建议验证**

- 通过已注册的配套编解码器，执行覆盖原生、RLE、JPEG、JPEG-LS 和 JPEG 2000 的互操作测试。
- 验证线上 Dataset 的传输语法和解码后像素，而不只是 DIMSE 响应状态。

### STD-001: 标准表生成与工具链

**状态：** `Partial`  
**优先级：** `P0`

被审计的生成源码包含 5,334 个唯一标准标签，fo-dicom 中为 5,343 个；包含 1,906 个唯一标准 UID，fo-dicom 中为 1,928 个。

缺失标签：

- `(0008,001D)` Sensitive Content Code Sequence
- `(0018,9390)` 至 `(0018,9392)` Metal Artifact Reduction 属性
- `(3004,0020)` 至 `(3004,0024)` RT Dose 属性

缺失的 UID 是上下文组 UID `1.2.840.10008.6.1.1550` 至 `1.2.840.10008.6.1.1571`。

三个生成器都依赖仓库本地的 `fo-dicom-code` 目录，但干净检出中不存在该目录。`go run ./tools/generate_tags` 会因为无法打开源文件而在生成前失败。

**验收标准**

- 生成器输入采用明确的 CLI 参数、已记录的外部源码检出，或许可证允许提交的标准输入。
- 干净检出可以复现所有生成文件。
- 生成过程确定且由 CI 执行检查。
- 标签、字典和 UID 源从同一基线原子更新。

**建议验证**

- 在干净的临时检出中运行所有生成器。
- 运行 `gofmt`、完整测试套件以及再生成无差异检查。
- 与固定参考版本比较唯一标签集和 UID 集。

### MED-001: DICOMDIR

**状态：** `Open`  
**优先级：** `P1`

通用解析可以读取 DICOMDIR 元素，但不存在介质目录领域模型。fo-dicom 提供目录记录、层级遍历、偏移修复、文件添加、读写工作流和可选的图标图像生成。

参考：[fo-dicom Media](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Media)

**验收标准**

- 打开和保存 DICOMDIR 文件时保持有效的记录偏移。
- 提供患者、检查、序列和实例记录的遍历能力。
- 添加文件时生成有效的 Referenced File ID，并进行确定性分组。
- 对错误或过期偏移提供有文档说明的严格/兼容行为。
- 图标生成保持可选，并与目录核心解耦。

**建议验证**

- 对真实 DICOMDIR 固件执行往返测试。
- 测试重复或匿名化标识符、长文件 ID、无效偏移和缺失可选属性。
- 使用 fo-dicom 交叉打开生成的 DICOMDIR 文件。

### NET-003: 高级关联协商

**状态：** `Partial`  
**优先级：** `P1`

PDU 和 association 包已经建模 User Identity、Asynchronous Operations Window、SCP/SCU Role Selection 和 SOP Class Extended Negotiation。高层客户端没有这些值的配置入口，`buildUserInformation` 只发送最大 PDU 长度和实现标识。

**验收标准**

- 高层客户端选项公开所有已支持的协商项。
- 关联建立后可以访问请求值和接受值。
- 明确定义 Positive User Identity 响应和必需响应缺失时的失败行为。
- 异步操作限制由请求调度器实际执行，而不只是在线路上编码。
- Role Selection 在适用时影响请求和子操作行为。

**建议验证**

- 执行编码/解码测试以及客户端/服务端关联集成测试。
- 在协商异步窗口限制之下、等于限制和超过限制时执行并发请求测试。
- 覆盖拒绝和异常响应。

### NET-004: SOP Class Common Extended Negotiation

**状态：** `Open`  
**优先级：** `P1`

PDU 项类型常量 `0x57` 已存在，但缺少针对 Service Class UID 和 Related General SOP Class UID 的完整结构、编码器、解码器、关联表示以及客户端/服务端集成。

参考：[fo-dicom DicomExtendedNegotiation](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Network/DicomExtendedNegotiation.cs)

**验收标准**

- 对完整 `0x57` 项进行编码和解码，并校验长度。
- 保留 Service Class UID 以及所有 Related General SOP Class UID。
- 通过关联 API 和高层客户端 API 公开该请求。
- 拒绝异常长度，且不发生 panic 或留下部分状态。

**建议验证**

- 针对代表性的 `0x57` 载荷执行字节级编码/解码往返测试。
- 使用多个 Related SOP Class UID 执行客户端/服务端关联测试。
- 与 fo-dicom 执行互操作测试，并测试异常和截断的项。

### SR-001: 完整结构化报告工作流

**状态：** `Partial`  
**优先级：** `P1`

标准 SR 类型均已有值类型常量，但构造和类型化读取主要集中在 TEXT、CODE、NUM 和 CONTAINER。PNAME、DATE、TIME、DATETIME、UIDREF、COMPOSITE、IMAGE、WAVEFORM、SCOORD 和 TCOORD 缺少完整的类型化 API。SR 专用的 Open 和 Save 方法仍是被注释的占位实现。

参考：[fo-dicom StructuredReport](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/StructuredReport)

**验收标准**

- 每个已声明的 Value Type 都有对称的构造器和类型化读取支持。
- 校验引用 SOP、空间坐标和时间坐标约束。
- 文件和流的打开/保存工作流保持 SR 内容树。
- 校验根关系规则和子关系规则。

**建议验证**

- 对每个 Value Type 执行表驱动往返测试。
- 测试嵌套内容树和非法关系。
- 使用 fo-dicom 交叉读取具有代表性的 SR 文档。

### IMG-001: Dataset 驱动的渲染

**状态：** `Partial`  
**优先级：** `P1`

LUT 和 overlay 基础能力已存在，但 `DicomImage` 由 `DicomPixelData` 构造，而不能直接由 Dataset 或文件构造。其默认灰度管线将 rescale slope/intercept 固定为 `1/0`，从像素计算最优窗，并未完整使用 Dataset 中的 Modality LUT、VOI LUT、窗、显示以及逐帧元数据。已保存的 `scale` 和 `showOverlays` 状态不会影响 `RenderFrame`。

参考：[fo-dicom DicomImage](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Imaging/DicomImage.cs)

**验收标准**

- 可以从 Dataset 和解析后的文件结果构造可渲染图像。
- 根据 Dataset 和逐帧元数据构建 modality、VOI、presentation、palette 和 inversion 阶段，并记录明确的优先级。
- 缩放会改变输出尺寸，并使用定义明确的插值模式。
- overlay 可见性、原点、帧范围和颜色会影响渲染输出。
- 保留调用方对窗、LUT、缩放和反色的显式覆盖。

**建议验证**

- 为 CT rescale/windowing、MONOCHROME1、调色板彩色、overlay 和多帧功能组增加黄金图像测试。
- 与 fo-dicom 比较代表性渲染像素，误差保持在声明的容差内。

### CORE-001: 递归验证

**状态：** `Partial`  
**优先级：** `P1`

很多具体 Element 类型实现了 VR 验证，但 Dataset 没有公共递归验证工作流。`Sequence.Validate` 当前直接返回 `nil`，不会验证子 Dataset。

**验收标准**

- Dataset 验证访问每个元素，并递归验证 sequence。
- 错误包含嵌套标签/条目路径和原始验证原因。
- 可以显式启用或禁用验证，且不会产生全局数据竞态。
- 明确定义 VM、VR、必需结构字段和 sequence 子项的验证范围。

**建议验证**

- 构造嵌套 sequence 失败，并精确断言错误路径。
- 测试多值、私有显式 VR、异常日期/时间、UID 和数值。
- 如果保留全局配置，对并发验证配置执行竞态测试。

### IMG-002: 几何与空间图像工具

**状态：** `Open`  
**优先级：** `P2`

fo-dicom 包含 FrameGeometry、患者/图像坐标转换、方向和定位支持、空间变换、直方图辅助工具、数学几何类型以及最近邻/双线性插值。go-dicom 没有等价的内聚几何层。

参考：[fo-dicom FrameGeometry](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Imaging/FrameGeometry.cs)

**验收标准**

- 解析经典和增强多帧几何信息。
- 按有文档说明的约定在像素坐标和患者坐标之间转换。
- 计算帧方向、法向量、包围盒和定位线。
- 提供旋转、翻转、平移、缩放和最佳适配变换。
- 提供经过测试、可被渲染和 MPR 复用的插值基础能力。

**建议验证**

- 对合成的轴位、矢状位、冠状位和斜位帧执行坐标往返测试。
- 使用具有已知患者空间几何的经典和增强多帧固件。
- 为变换和最近邻/双线性插值增加黄金测试。

### CORE-002: Dataset 遍历器与规则

**状态：** `Open`  
**优先级：** `P2`

fo-dicom 提供递归 Dataset 遍历器，以及可组合的匹配规则和转换规则。go-dicom 使用方目前必须自行构建这些行为。

参考：[fo-dicom DicomDatasetWalker](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/DicomDatasetWalker.cs)

**验收标准**

- 使用稳定的路径信息访问 element、sequence、item 和 fragment。
- 支持提前停止和错误传播。
- 提供可组合的存在、空值、相等、通配符和正则匹配规则。
- 提供常用的删除、设置、映射、复制、正则、大小写和 UID 转换规则。
- 明确定义转换是修改源 Dataset 还是克隆 Dataset。

**建议验证**

- 对嵌套 sequence 和像素 fragment 断言遍历顺序及路径。
- 覆盖提前停止、访问器错误和规则组合优先级。
- 验证修改型和克隆型转换都保留所有无关元素。

### DICT-001: 运行时 XML 字典加载

**状态：** `Partial`  
**优先级：** `P2`

`Dictionary.Add` 支持编程式扩展，但没有运行时读取 fo-dicom 兼容 XML 字典的能力，包括私有创建者字典和掩码标签。

参考：[fo-dicom DicomDictionaryReader](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/DicomDictionaryReader.cs)

**验收标准**

- 从 `io.Reader` 加载标准和私有字典。
- 支持精确标签和掩码标签、多个 VR、VM、keyword、retired 标志和 private creator。
- 拒绝异常输入，并提供元素级上下文。
- 明确定义重复条目和字典合并行为。

**建议验证**

- 加载具有代表性的 fo-dicom 标准和私有 XML 字典。
- 测试掩码标签查找优先级、重复条目和合并行为。
- 覆盖异常 XML、无效 VR/VM 值和不完整私有条目。

### ANON-001: 自定义配置加载

**状态：** `Complete`
**优先级：** `P2`

已于 2026-08-14 完成。`NewProfileFromReader` 现在接受任意 `io.Reader`；
`LoadProfileFromFile` 会打开指定路径并委托给同一个解析器。现有两列
`pattern;action` 配置保持兼容，同时支持 fo-dicom 兼容的 12 列配置，且全部
11 个 `SecurityProfileOptions` 列与内置 profile 加载器采用相同的优先级语义。

解析现在默认为严格模式：非注释输入必须使用两种受支持的列布局之一，pattern
和 action 都会被验证，错误包含源文件行号。空行和以 `#` 开头的注释仍受支持。
当前未提供宽松模式。

**验收标准**

- 文件加载委托给同一个基于 Reader 的解析器。
- Reader API 接受通用 `io.Reader`，而不只是 `*strings.Reader`。
- 内置输入和自定义输入采用一致的 profile option 语义。
- 对无效行和 action 返回可操作的错误；除非显式选择宽松模式，否则不能静默跳过。

**验证证据**

- 文件 API 和 Reader API 已使用相同输入比较规则输出。
- 测试覆盖通用 `io.Reader`、每个 profile option 列、组合 option 优先级、
  注释、异常列数、未知 action、空 pattern 和无效正则表达式。
- `ExampleNewProfileFromReader` 由 Go 测试套件编译并执行。
- `go test ./pkg/dicom/anonymizer -count=1` 通过。
- `go test -race ./pkg/dicom/anonymizer -count=1` 通过。
- `go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1` 通过。
- `golangci-lint run` 报告 0 个问题。
- 当前验证对应工作区代码，尚未创建 ANON-001 实现提交。

### PRINT-001: 基于 Dataset 的打印管理模型

**状态：** `Partial`  
**优先级：** `P2`

FilmSession、FilmBox、ImageBox、PresentationLUT 和打印机状态辅助能力已经存在，但主要模型是简化的 Go struct，而不是基于 Dataset 的 DICOM 对象。它们只覆盖部分属性，缺少 fo-dicom 的克隆、加载/保存、UID 查找/删除和完整布局行为。空 UID 使用固定占位值，而不是生成的新 UID。

参考：[fo-dicom Printing](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Printing)

**验收标准**

- 模型通过 Dataset 对象往返所有受支持的 DICOM 属性。
- 空 SOP Instance UID 会生成唯一值。
- 克隆、加载/保存、查找和删除操作保持父子引用。
- 胶片布局和 image box 创建支持已声明范围内的标准显示格式。
- 端到端演示 N-CREATE/N-SET/N-ACTION 集成。

**建议验证**

- 对每个受支持的打印对象和显示格式执行 Dataset 往返测试。
- 验证克隆之间的 UID 唯一性和父子引用完整性。
- 执行端到端打印工作流，并使用 fo-dicom 交叉读取生成的对象。

### OBS-001: 网络可观测性

**状态：** `Open`  
**优先级：** `P2`

fo-dicom 提供结构化日志、请求已发送/待处理/已完成/已超时事件，以及网络指标收集器钩子。go-dicom 有生命周期回调，但没有内聚的日志或指标接口；部分 PDU 解码警告还会直接通过 `fmt.Printf` 输出。

参考：[fo-dicom Network Metrics](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Log/Metrics)

**验收标准**

- 库代码绝不直接写入 stdout/stderr。
- 客户端和服务端接受默认无操作的结构化日志钩子。
- 请求生命周期事件包含关联、消息 ID、命令、状态、耗时以及超时/取消结果。
- 指标钩子提供连接、关联、DIMSE、字节、错误和延迟观测，且不依赖特定遥测厂商。

**建议验证**

- 断言默认配置不会产生任何进程输出。
- 集成测试成功、待处理、超时、取消、拒绝和传输失败路径中的钩子顺序。
- 对并发关联以及缓慢或失败的 observer 执行竞态测试。

### IMG-003: 体数据重建与 MPR

**状态：** `Partial`  
**优先级：** `P3`

reconstruction 包记录了 ImageData、VolumeData、Slice、Stack 和 DicomGenerator，但这些类型仍是占位实现。构造器返回 `ErrNotImplemented`，`NewDicomGenerator` 返回 `nil`。

参考：[fo-dicom Reconstruction](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Imaging/Reconstruction)

本条目依赖 IMG-001 和 IMG-002。

**验收标准**

- 只从几何兼容的切片构建体数据。
- 对切片排序，并检测不规则间距和 Frame of Reference 不匹配。
- 以定义明确的插值和体外行为生成轴位、冠状位、矢状位和任意切面。
- 生成具有有效几何、派生元数据、UID 和像素表示的派生 DICOM 实例。
- 明确定义大型检查的内存和并发行为。

**建议验证**

- 使用切面结果可解析预测的合成体数据。
- 测试不规则间距、逆序、斜位方向和多帧 Dataset。
- 与 fo-dicom 交叉检查几何和代表性像素值。
- 在优化体数据和切面生成之前增加基准测试。

### MED-002: DICOM 文件扫描器

**状态：** `Open`  
**优先级：** `P3`

fo-dicom 提供扫描文件、报告 DICOM/非 DICOM 结果并与介质工作流集成的扫描器。go-dicom 有解析器和 CLI 示例，但没有可复用的扫描器抽象。

参考：[fo-dicom DicomFileScanner](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Media/DicomFileScanner.cs)

**验收标准**

- 使用 context cancellation 扫描文件和目录树。
- 分别报告有效 DICOM 文件、无效文件和读取错误。
- 明确定义递归、符号链接、并发和遇错停止行为。
- 只需分类或元数据时，避免加载大型像素数据。

**建议验证**

- 扫描同时包含 DICOM、非 DICOM、不可读文件和符号链接的混合目录树。
- 测试取消、有界并发、确定性结果统计以及两种遇错停止模式。
- 检测大型固件，确认仅元数据扫描不会读取像素载荷。

## 外部能力边界

### 压缩编解码器

**状态：** `External`

核心仓库有意只提供原生传输语法编解码器、注册表、转码器和封装支持。压缩编解码器由 `github.com/cocosip/go-dicom-codecs` 通过空白导入注册提供。

fo-dicom Core 包含 RLE 和 JPEG Lossless 解码器实现，但这并不要求把所有压缩编解码器移回 `go-dicom`。能力声明必须指出配套模块，验证必须覆盖组合后的运行时。

### 平台集成

**状态：** `Not a gap`

WPF、ImageSharp、SkiaSharp、ASP.NET 依赖注入以及 .NET 特有的异步 API 形式均属于平台集成。除非单独批准了相应的 Go 使用场景和包边界，否则不应移植这些能力。

### 全视野切片成像

**状态：** `Not a one-sided gap`

两个库都提供与 WSI 兼容的低层标签和多帧像素数据，但被审计版本都没有完整的金字塔/层级/坐标 WSI 领域 API 或 WSI IOD 验证器。未来的 WSI 提案应作为新的跨库能力单独确定范围，而不是作为 fo-dicom 对标工作。

## 交付阶段

### Phase 0: 公共契约与互操作性

范围：DOC-001、NET-001、NET-002、STD-001。

当前进度：NET-001 已完成；DOC-001 仍为部分完成，NET-002 和 STD-001 尚未
完成，因此 Phase 0 尚未完成。

阶段验收：

- README 示例可编译，能力描述与实际公共 API 一致。
- 高层 SCU 可以通过普通 TCP 和经验证的 TLS 建立连接。
- C-STORE 使用协商后的传输语法发送，或执行经过验证的转码。
- 可以从干净检出复现标准生成数据，且不存在差异。
- 聚焦测试、完整包测试、变更后的共享/网络代码竞态测试和相关互操作检查全部通过。

### Phase 1: 主要领域能力对标

范围：MED-001、NET-003、NET-004、SR-001、IMG-001、CORE-001。

阶段验收：

- DICOMDIR 创建和往返可在两个库之间工作。
- 高级关联值得到协商，并由客户端行为实际执行。
- 所有已声明 SR Value Type 均可通过 Dataset 和文件工作流往返。
- 渲染使用 Dataset 元数据并生成经过验证的输出。
- Dataset 验证报告嵌套路径并验证 sequence 子项。

每个条目都应独立实现和发布；Phase 1 不是一个单独的拉取请求。

### Phase 2: 支撑 API 与运维能力

范围：IMG-002、CORE-002、DICT-001、ANON-001、PRINT-001、OBS-001。

当前进度：ANON-001 已完成；IMG-002、CORE-002、DICT-001、PRINT-001 和
OBS-001 尚未完成，因此 Phase 2 尚未完成。

阶段验收：

- 在开始 MPR 工作之前稳定共享几何和遍历器 API。
- 运行时字典和匿名化配置具备严格的解析测试。
- 打印对象通过 Dataset 和 DIMSE N-service 工作流往返。
- 网络诊断可注入、结构化，且默认静默。

### Phase 3: 专用工作流

范围：IMG-003 和 MED-002。

阶段验收：

- 针对合成用例和参考用例验证 MPR 几何和像素。
- 扫描大型目录树时，扫描器行为保持有界且可取消。
- 通过性能基准建立大型检查和扫描的基线。

## 维护规则

- 发布后保持 ID 稳定。
- 不删除已完成条目；将其标记为 `Complete`，并记录验证提交、测试和任何有意延后的范围。
- 只有在重新执行源码对比之后才能更新审计基线。
- 不能仅根据 README 变更或类型声明将条目标记为完成。
- 将源码/单元测试证据、集成行为和跨库互操作性视为不同层级的验证。
- 开始实施前，应将新发现的差距同时加入详细列表和阶段表。
