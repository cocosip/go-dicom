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
| NET-002 | P0 | Complete | C-STORE 传输语法选择与自动转码 |
| STD-001 | P0 | Complete | 可复现且保持最新的标准生成表 |
| MED-001 | P1 | Complete | DICOMDIR 介质目录模型及读写工作流 |
| NET-003 | P1 | Complete | 通过高层客户端执行高级关联协商 |
| NET-004 | P1 | Complete | SOP Class Common Extended Negotiation |
| SR-001 | P1 | Complete | 完整的结构化报告值类型和文件工作流 |
| IMG-001 | P1 | Complete | Dataset 驱动的图像渲染管线 |
| CORE-001 | P1 | Complete | Dataset 和 Sequence 递归验证 |
| IMG-002 | P2 | Complete | 帧几何、空间变换和插值工具 |
| CORE-002 | P2 | Complete | Dataset 遍历器、匹配规则和转换规则 |
| DICT-001 | P2 | Complete | 运行时 XML 字典加载 |
| ANON-001 | P2 | Complete | 完整的自定义匿名化配置加载 |
| PRINT-001 | P2 | Complete | 基于 Dataset 的 DICOM 打印管理模型 |
| OBS-001 | P2 | Open | 结构化网络日志、请求事件和指标钩子 |
| IMG-003 | P3 | Complete | 体数据重建和 MPR |
| MED-002 | P3 | Open | DICOM 文件扫描工作流 |

## 实施进度

截至 2026-08-16：

- **已完成：** NET-001、NET-002、STD-001、MED-001、NET-003、NET-004、
  SR-001、IMG-001、IMG-002、IMG-003、CORE-001、CORE-002、ANON-001、
  DICT-001 和 PRINT-001。
- **未完成：** DOC-001、OBS-001 和 MED-002 继续保持 `Partial` 或 `Open` 状态。
- Phase 0 尚未完成。NET-001、NET-002 和 STD-001 已完成；其余 Phase 0
  工作由 DOC-001 跟踪。
- **下一项：** OBS-001，即下方计划开发顺序中的第一个未完成条目。

## 计划开发顺序

优先级表示能力的重要程度；此顺序表示计划实施的先后次序。工作完成后保持顺序号稳定，
只将对应行更新为 `Complete`；表中第一个未完成条目就是下一项开发内容。只有在本文档中
记录了新的依赖或范围证据后，才能调整顺序。

| 顺序 | ID | 优先级 | 当前状态 | 排序理由 |
| ---: | --- | --- | --- | --- |
| 1 | NET-004 | P1 | Complete | 趁 NET-003 的 Association API 和测试上下文仍然清晰，补完整个协商能力族。 |
| 2 | CORE-001 | P1 | Complete | 在完成嵌套领域工作流之前，先建立 Dataset 和 Sequence 的递归正确性保障。 |
| 3 | SR-001 | P1 | Complete | 基于 CORE-001 的递归验证行为完成强类型 SR 树和文件工作流。 |
| 4 | IMG-001 | P1 | Complete | 在加入共享空间工具之前，先建立 Dataset 驱动的渲染管线。 |
| 5 | IMG-002 | P2 | Complete | 渲染需求稳定后补齐几何、变换和插值；该项也是 IMG-003 的前置条件。 |
| 6 | CORE-002 | P2 | Complete | 待 CORE-001 和 SR-001 明确遍历语义后，再将 walker、路径、匹配和转换 API 通用化。 |
| 7 | PRINT-001 | P2 | Complete | 核心 Dataset 能力稳定后，完成 Dataset-backed 打印模型和 N-service 工作流。 |
| 8 | OBS-001 | P2 | Open | 协商和打印网络工作流稳定后，再加入横切的网络诊断能力。 |
| 9 | MED-002 | P3 | Open | 在更大型的重建工作之前，先交付独立且边界明确的扫描器工作流。 |
| 10 | IMG-003 | P3 | Complete | 仅在 IMG-001 和 IMG-002 完成后实施体重建和 MPR。 |
| 11 | DOC-001 | P0 | Partial | 主要能力完成后再做最终公共 API 和 README 审计，避免文档反复调整。 |

## 详细差距

### DOC-001: 公共能力声明

**状态：** `Partial`  
**优先级：** `P0`

在审计基线中，README 的若干声明超出了当时已实现的公共工作流，包括完整的客户端 TLS、完整的 SR 值类型、图像重建、通过客户端进行高级协商，以及打印任务创建。TLS 示例调用了不存在的 `client.WithTLS` 选项；另一个渲染示例向 `NewDicomImage` 传入 Dataset，但该函数实际接收 `*DicomPixelData`。

在审计基线中，README 的标签和 UID 数量与被审计源码中去重后的标准生成条目数不一致。

2026-08-14 进度：NET-001 已补齐高层客户端 TLS API，并修复 README 中的
TLS 示例；NET-002 已使 C-STORE 传输语法选择与转码声明和实际发送路径一致；
STD-001 已修正生成 Tag、UID 的数量和工具说明。SR、渲染、高级协商、打印及
其他公共能力声明尚未重新审计或修复，因此 DOC-001 仍为 `Partial`。

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

**状态：** `Complete`
**优先级：** `P0`

已于 2026-08-14 完成。共享 DIMSE 发送路径现在同时依据 SOP Class 和 Dataset
源传输语法选择 C-STORE 表示上下文，并优先使用与源语法完全匹配的上下文。
没有精确匹配时，会按确定性顺序选择 codec registry 可转码的第一个已接受语法，
并发送转码副本。显式指定的 Presentation Context ID 保持优先，并会在使用前验证。

包含 Pixel Data 的 Dataset 必须声明源传输语法；解析及 DIMSE 接收到的 Dataset
都会自动保留该语法，因此也支持接收后直接转发。源语法缺失或没有可用 codec 时，
会在写入网络前返回包含 SOP Class 和传输语法上下文的明确错误。没有 Pixel Data
的对象可以直接重新编码，不需要图像 codec，但显式 context ID 仍会被验证。
调用方 Dataset 和 codec 输入帧均与转码过程中的修改隔离。

`CStore`、`CStoreWithPriority` 和 C-GET 的 C-STORE 子操作均复用同一 service
路径。`CStoreMultiple` 保持顺序发送，在首次失败或取消时停止，并返回已经完成且
成功的对象数量。

参考：[fo-dicom DicomCStoreRequest](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Network/DicomCStoreRequest.cs)

**验收标准**

- 表示上下文选择同时考虑 SOP Class 和源传输语法。
- 被接受时优先使用原始语法。
- 必要时使用已注册的编解码器进行回退转码。
- 没有可直接使用或可转码的已接受语法时，返回明确错误。
- 不修改调用方的 Dataset。
- 批量发送明确定义顺序、部分成功、取消和并发语义。

**验证证据**

- Service 测试检查实际 PDV 的 Presentation Context，并解码线上 Dataset，验证
  原语法选择和 Little/Big Endian 转码。
- 测试覆盖已注册压缩 codec 回退、codec 缺失、源语法未知、显式上下文 ID
  （包括已拒绝和 SOP Class 不匹配的 context）、无 Pixel Data 对象、接收后转发，
  以及面对会修改输入帧的 codec 时调用方 Dataset 仍保持不变。
- Client 测试明确覆盖顺序批量发送、部分成功计数、首次失败和取消语义。
- `CGO_ENABLED=0 go test ./pkg/imaging/... ./pkg/network/... -count=1` 通过。
- `CGO_ENABLED=0 go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`
  通过。
- `CGO_ENABLED=0 golangci-lint run` 报告 0 个问题。实现未增加 CGo 指令、C 导入、
  原生依赖或模块变更。
- 与真实配套 RLE、JPEG、JPEG-LS 和 JPEG 2000 codec 的跨进程互操作仍属于独立
  验证层级。

### STD-001: 标准表生成与工具链

**状态：** `Complete`
**优先级：** `P0`

已于 2026-08-14 完成。旧生成数据包含 5,338 个 Tag、5,338 个标准字典条目和
1,906 个标准 UID；三个一次性工具还通过硬编码读取仓库中并不存在的
`fo-dicom-code` 目录。

现在将 fo-dicom 2026b 的 `DICOM Dictionary.xml` 和人工维护的
`Private Dictionary.xml` 固定保存在 `tools/data/2026b`。统一的
`tools/generate_dicom` 命令接收两个输入路径和仓库根目录，一次性重新生成四类数据：

- 5,347 个标准 Tag 常量
- 1,928 个标准 UID 常量
- 5,347 个标准字典条目
- 235 个 private creator，共 4,678 个私有条目

生成器按 fo-dicom 规则为 retired Tag 和 UID 的生成标识符追加 `RETIRED`，字典
keyword 则保持 XML 原值。`pkg/dicom/uid/uids_private.go` 中的 59 个私有 UID
在这两份权威 XML 中没有来源，因此本工具不会重新生成它们。更新基线采用显式复制
XML 并重新生成的方式，不提供 fo-dicom 源码下载工具。

**验收标准**

- 生成器输入采用明确的 CLI 参数、已记录的外部源码检出，或许可证允许提交的标准输入。
- 干净检出可以复现所有生成文件。
- 生成过程确定且由 CI 执行检查。
- 标签、字典和 UID 源从同一基线原子更新。

**验证证据**

- 仓库内两份 XML 的 SHA-256 与本地 fo-dicom 2026b 源文件逐字节一致。
- 独立解析 XML 确认 5,347 个 Tag、1,928 个 UID、235 个 private creator 和
  4,678 个私有条目。
- 生成器集成测试会在临时目录中用仓库 XML 重新生成四个输出，断言精确数量，并在
  归一化平台换行符后与已提交文件比较；该测试位于现有 CI 的 `./tools/...` 测试范围内。
- 聚焦测试覆盖 fo-dicom retired 标识符规则和当前 `vm` 符号映射。

### MED-001: DICOMDIR

**状态：** `Complete`
**优先级：** `P1`

已于 2026-08-15 完成。`pkg/media` 现在提供 DICOMDIR 创建、严格/兼容读取、
层级遍历、确定性文件分组、受限偏移修复和两遍写入。支持 PATIENT、STUDY、
SERIES、IMAGE、SR DOCUMENT 和 PRESENTATION 记录，并在读取时保留未知记录类型。
不会扫描、复制、移动、重命名或重写引用文件。

可选图标生成由 `pkg/imaging` 通过结构化接口实现，`pkg/media` 与 `pkg/imaging`
互不导入。图标使用现有纯 Go codec registry，渲染代表帧、保持宽高比，并生成
最大 128x128 的 8 位 MONOCHROME2 图像。

参考：[fo-dicom Media](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Media)

**验收标准**

- 打开和保存 DICOMDIR 文件时保持有效的记录偏移。
- 提供患者、检查、序列和实例记录的遍历能力。
- 添加文件时生成有效的 Referenced File ID，并进行确定性分组。
- 对错误或过期偏移提供有文档说明的严格/兼容行为。
- 图标生成保持可选，并与目录核心解耦。

**验证证据**

- 聚焦测试覆盖 File ID 校验、重复和匿名化分组、缺失属性、精确偏移、固定差值与
  类型/物理顺序恢复、歧义、循环、重复引用、不可达记录、writer 错误、源对象不变、
  代表帧和图标失败。
- 13,796 字节的 fo-dicom DICOMDIR 固件可在严格模式下读取 80 条记录，并通过
  两种受支持的目录传输语法往返。
- fo-dicom 6.0.0-alpha1 成功打开 Go 生成的 Explicit 和 Implicit VR Little
  Endian DICOMDIR；9 条记录的层级、类型计数和 Referenced File ID 完全一致。
- `CGO_ENABLED=0 go test ./pkg/media ./pkg/imaging/... -count=1` 通过；未增加 CGo
  指令、C 导入或原生依赖。
- `CGO_ENABLED=0 go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`
  通过。
- `CGO_ENABLED=0 golangci-lint run` 报告 0 个问题。

### NET-003: 高级关联协商

**状态：** `Complete`
**优先级：** `P1`

高层客户端现已提供 Asynchronous Operations Window、SCP/SCU Role Selection、
SOP Class Extended Negotiation 和 User Identity 配置入口。User Identity 支持
Username、Username/Password、Kerberos、SAML 和 JWT；关联建立后同时保留请求值和
接受值。

请求 Positive User Identity 响应时，客户端默认要求服务端返回响应；兼容选项可显式
允许响应缺失。协商后的最大 invoked operations 会实际限制尚未完成的 DIMSE 请求数，
`0` 表示无限；最大 performed operations 与 fo-dicom 一致，仅协商并暴露，不限制
接收线程。Role Selection 与 Presentation Context 绑定，并约束普通请求和反向
C-STORE 子操作。Common Extended Negotiation `0x57` 由 NET-004 单独完成。

**验收标准**

- 高层客户端选项公开所有已支持的协商项。
- 关联建立后可以访问请求值和接受值。
- 明确定义 Positive User Identity 响应和必需响应缺失时的失败行为。
- 异步操作限制由请求调度器实际执行，而不只是在线路上编码。
- Role Selection 在适用时影响请求和子操作行为。

**验证证据**

- 真实客户端/服务端关联集成测试往返全部 NET-003 协商值，并验证接受后的异步请求限制。
- 聚焦测试覆盖有限和无限请求窗口、等待时取消、默认和协商角色、反向 C-STORE 拒绝、
  Positive Identity 响应缺失以及异常 Role 响应。
- PDU 往返测试覆盖 AC 异步窗口、Role Selection、Extended Negotiation 和存在但为空的
  User Identity 响应。
- `CGO_ENABLED=0 go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`
  在 Go 1.26.6 Windows/amd64 上通过。
- `CGO_ENABLED=0 go build ./...` 通过，`CGO_ENABLED=0 golangci-lint run`
  报告 0 个问题。

### NET-004: SOP Class Common Extended Negotiation

**状态：** `Complete`
**优先级：** `P1`

已于 2026-08-15 完成。PDU 层现在可以完整编码和解码 `0x57` 项；关联层和高层
客户端会按 SOP Class 合并其 Service Class UID、保持顺序的 Related General SOP
Class UID，以及同一 SOP Class 的 `0x56` Application Information。客户端选项会复制
调用方数据；显式但无效的 Common 请求会保留到 PDU 编码阶段并返回错误。

Common Extended Negotiation 仍然只用于请求：A-ASSOCIATE-AC 既不发送也不接纳
`0x57`。空 Related General SOP Class UID 列表合法；SOP Class UID、Service Class UID
以及每个 Related UID 必须非空。写入或分配前会校验嵌套和外层 16 位长度。

参考：[fo-dicom DicomExtendedNegotiation](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Network/DicomExtendedNegotiation.cs)

**验收标准**

- 对完整 `0x57` 项进行编码和解码，并校验长度。
- 保留 Service Class UID 以及所有 Related General SOP Class UID。
- 通过关联 API 和高层客户端 API 公开该请求。
- 拒绝异常长度，且不发生 panic 或留下部分状态。

**验证证据**

- 精确字节和往返测试覆盖组合的 `0x56`/`0x57` 值、空 Related 列表、多个 Related
  UID、必填 UID 为空、嵌套数据过大、异常长度、不完整项头以及 AC 方向约束。
- 真实 Go 客户端/服务端关联通过高层客户端和服务端 Association API 往返 Common 值。
- 与 fo-dicom 修订 `7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2` 的双向完整 PDU
  检查通过：fo-dicom 可解析 Go 请求，Go 也可解析 fo-dicom 请求。
- `go test ./pkg/network/... -count=1`、仓库完整测试、`go build ./...` 和
  `golangci-lint run` 在 Go 1.26.6 Windows/amd64 上通过。
- Windows race 门禁受环境阻塞：即使执行
  `go test -race fmt -run '^$' -count=1`，也会在进入仓库测试代码前以
  `0xc0000139` 状态退出。

### SR-001: 完整结构化报告工作流

**状态：** `Complete`
**优先级：** `P1`

SR 包现在为全部已声明值类型提供类型化构造和读取，包括 SCOORD/TCOORD 基数约束、
引用 SOP 校验、长代码值与 URN 代码值、解析 by-reference 内容，以及带完整路径的递归
语义校验。Open、Read、Write 和 Save 会保留解析得到的 File Meta Information 与传输
语法，拒绝 partial 或非法报告，并强制使用显式 Sequence 和 Item 长度，与 fo-dicom 的
SR 输出行为保持一致。

参考：[fo-dicom StructuredReport](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/StructuredReport)

**验收标准**

- 每个已声明的 Value Type 都有对称的构造器和类型化读取支持。
- 校验引用 SOP、空间坐标和时间坐标约束。
- 文件和流的打开/保存工作流保持 SR 内容树。
- 校验根关系规则和子关系规则。

**已执行验证**

- 构造器/getter 往返覆盖全部已声明 Value Type、代码值标签选择、坐标选择/基数和切片隔离。
- 递归测试覆盖根与子关系、冲突值标签、损坏 sequence、by-reference 条目、引用 UID 和嵌套路径。
- `test-data/test_SR.dcm` 与 fo-dicom fixture 字节完全一致，并完成 UIDREF、SCOORD、
  TCOORD、by-reference 和深层嵌套内容的往返，同时保留 File Meta Information 与传输语法。
- SR/parser/writer 聚焦测试、仓库完整测试、`go build ./...` 和
  `golangci-lint run` 均通过。
- Windows race runtime 对受影响包以及 `go test -race fmt -run '^$'` 都以
  `0xc0000139` 退出；race 仍需由 CI 验证，不能记为本地通过。

### IMG-001: Dataset 驱动的渲染

**状态：** `Complete`
**优先级：** `P1`

已于 2026-08-15 完成。`DicomImage` 现在可以从 Dataset、解析结果或文件构造；
Dataset 构造会保留私有克隆，通过可注入 codec registry 解码封装帧，并根据顶层及
功能组元数据创建逐帧灰度管线。

窗选择与 fo-dicom 保持同一优先级：有效的顶层窗、功能组窗、有效的
Smallest/Largest Image Pixel Value，最后是排除 padding 后的像素最小/最大值。
回退范围会先经过显式 Modality LUT 或 rescale，再计算窗宽窗位。显式 Modality/VOI
LUT Sequence、VOI LUT Function、MONOCHROME1 显示、palette color、调用方反色及
灰度色表都会影响输出。

渲染现已支持打包 1-bit 和有符号/无符号 32-bit 灰度、双线性缩放（1-bit 使用最近邻）、
显式及嵌入 overlay、overlay 原点/帧范围/裁剪/颜色/可见性，以及调用方切换 VOI LUT
和控制 LUT 是否传播到所有帧。源 Dataset 不会被修改。

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

**已执行验证**

- 合成像素测试覆盖 rescale/windowing、Modality/VOI LUT 优先级、16-bit VOI 输出
  归一化、SIGMOID VOI、MONOCHROME 显示、原生及封装 palette 转换、1/32-bit 样本、
  32-bit Explicit VR Big Endian 像素、Big Endian Modality/VOI/Palette LUT Data、
  功能组、缩放、色表和显式/嵌入 overlay。
- Dataset、解析结果、文件、注入 codec、克隆和源数据保持工作流均有聚焦测试，
  包括调用方 pipeline 覆盖，以及无 Basic Offset Table 多帧 fragment 的就地修改 codec。
- `go test ./pkg/imaging/... -count=1`、
  `go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`、
  `go build ./...`、`golangci-lint run` 和 `git diff --check` 均通过。构建明确退出成功，
  但会提示只读的默认 Go module stat cache 无法更新。
- `go test -race fmt -run '^$'` 在当前 Windows 主机上以 `0xc0000139` 无法启动；
  race 仍需由 CI 验证，不能记为本地通过。

### CORE-001: 递归验证

**状态：** `Complete`
**优先级：** `P1`

Dataset 现在提供显式递归验证和默认开启的写入自动验证。Dataset 与 Sequence
共用一套遇错即停的验证引擎，按确定的标签和条目顺序遍历，先验证实际 VR 值再验证字典
VM，并保留 fo-dicom 的 VM 例外、完整嵌套路径和原始错误原因。二进制、JSON 和 XML
装载仍可读取无效值，并在返回前恢复自动验证。

**验收标准**

- Dataset 验证访问每个元素，并递归验证 sequence。
- 错误包含嵌套标签/条目路径和原始验证原因。
- 可以显式启用或禁用验证，且不会产生全局数据竞态。
- 明确定义 VM、VR、必需结构字段和 sequence 子项的验证范围。

**已执行验证**

- 嵌套 sequence、精确路径、VM、私有/未知标签、异常值、自动写入、失败回滚和装载
  回归测试均通过。
- `go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`、
  `go build ./...` 和 `golangci-lint run` 通过。
- 受影响包的 race 命令在当前 Windows 主机上以 `0xc0000139` 无法启动；
  `go test -race fmt` 也以相同状态失败，因此 race runtime 仍需由 CI 验证，不能记为
  本地通过。

### IMG-002: 几何与空间图像工具

**状态：** `Complete`
**优先级：** `P2`

已于 2026-08-15 完成。`pkg/imaging/geometry` 现在可按与 fo-dicom 一致的
spacing 和功能组优先级提取经典及增强多帧几何，提供方向、法向量、像素中心角点、
包围盒、投影/相交定位线，以及患者坐标与图像坐标互转。图像坐标采用零基像素中心约定：
`(0,0)` 是首像素中心，末像素中心是 `(columns-1, rows-1)`。

`pkg/imaging/math3d`、`pkg/imaging/transform` 和
`pkg/imaging/interpolation` 分别提供有限值校验的几何/4x4 矩阵、可组合 affine 与
viewer 空间变换及最佳适配，以及支持 stride 的标量最近邻/双线性采样与缩放。
imaging 包同时提供已修正窗口累计行为的整数直方图。

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

**已执行验证**

- 合成轴位、矢状位、冠状位和斜位帧覆盖坐标往返、方向、法向量、零基像素中心角点、
  包围盒、投影定位框和相交定位线。
- 经典及增强多帧 Dataset 覆盖顶层/共享/逐帧优先级、帧边界、缺失几何、非法值和
  非有限值。
- 确定性测试覆盖向量、平面、矩阵、直方图、affine、最佳适配、最近邻、双线性、
  stride、端点及单像素场景。
- 可执行 Go 示例覆盖 Dataset 几何提取和坐标转换、定位线、affine/viewer 变换、
  插值及直方图窗口；详细约定与错误行为记录在 `pkg/imaging/README.md`。
- `CGO_ENABLED=0 go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`、
  `CGO_ENABLED=0 go build ./...` 和 `CGO_ENABLED=0 golangci-lint run` 均通过。
  本地验收固定禁用 CGO，因此未运行 race 命令。

### CORE-002: Dataset 遍历器与规则

**状态：** `Complete`
**优先级：** `P2`

已于 2026-08-16 完成。`pkg/dicom/dataset` 现提供确定性、迭代式 Dataset
walker，包含共享不可变路径、成对的 Sequence/item 与 fragment 事件、
continue/skip/stop 控制、visitor 错误上下文、循环检测、容器快照和明确的空 item
处理。Dataset 验证、writer 长度计算与编码、`dicomdump` 均复用该 walker。
writer 在不保留第二套递归遍历的前提下，继续支持显式/未定义长度、流式值、fragment、
group 处理、deflate 和 Sequence item offset observer。

`pkg/dicom/dataset/rules` 提供可组合的存在/内容/布尔匹配规则，以及按顺序执行的
条件转换，包括 remove、set、map、copy、regex、大小写、trim、pad、truncate 和
UID 操作。默认应用返回独立 clone；原地应用仅提交成功完成的转换 clone。按顺序记录的
`ChangeSet` 快照与分阶段错误会保留路径和原始原因。

fo-dicom 的读取期 observer 与 Dataset walker 是两条独立管线。因此 go-dicom 未增加
公开 Reader Walker；DIMSE Dataset 解码改为使用 assumed transfer syntax 和 eager read
复用 `pkg/dicom/parser`，由同一 parser 覆盖 Sequence、fragment、字符集、private creator
和 big endian。

参考：[fo-dicom DicomDatasetWalker](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/DicomDatasetWalker.cs)

**验收标准**

- 使用稳定的路径信息访问 element、sequence、item 和 fragment。
- 支持提前停止和错误传播。
- 提供可组合的存在、空值、相等、通配符和正则匹配规则。
- 提供常用的删除、设置、映射、复制、正则、大小写和 UID 转换规则。
- 明确定义转换是修改源 Dataset 还是克隆 Dataset。

**已执行验证**

- Dataset、rules、element、writer、parser、network service、media 和
  `dicomdump` 聚焦测试及示例均通过。
- 完整 `CGO_ENABLED=0 go test ./cmd/... ./examples/... ./pkg/... ./tools/...
  -count=1` 与 `CGO_ENABLED=0 go build ./...` 均通过。
- 使用仓库本地缓存执行 `golangci-lint run --allow-parallel-runners`，报告 0 个问题。
- Walker、match 和 transform benchmark 均完成并输出 allocation 数据。
- Windows 本地验收使用 `CGO_ENABLED=0`，未运行 race 测试。

**有意边界：** `SplitFormat`、隐式递归规则应用、公开 Reader Walker/Observer、
parser 状态机替换和 DIMSE streaming 重构不属于 CORE-002。

### DICT-001: 运行时 XML 字典加载

**状态：** `Complete`
**优先级：** `P2`

已于 2026-08-14 完成。`NewFromXML` 可以从任意 `io.Reader` 创建字典；
`Dictionary.LoadXML` 会先验证完整文档，再将其合并到已有字典。fo-dicom 使用的
单个 `<dictionary>` 标准布局和包含多个 creator 的 `<dictionaries>` 私有布局均受支持。

实现支持精确标签、掩码标签、fo-dicom 的全部 VR 分隔符、备选 VM、keyword、
retired 标志、UTF-8 BOM，以及组合源文件中已知但不属于字典条目的 `<uid>` 节点。
私有条目按 Private Creator 隔离到不同子字典；查询标签带 creator 时会自动路由。
精确标签优先于掩码标签；精确和掩码重复项均采用后加载覆盖语义。异常输入会尽可能
报告 creator、条目序号和标签值，加载失败不会向目标字典留下部分修改。

运行时加载的私有字典也已接入 Implicit VR 解析。Private Creator 预留元素固定按
`LO` 解码；后续私有元素按 group、已分配 block 和 creator 解析。Creator 预留关系在
根数据集以及每个序列 Item 中分别维护，因此不同 Item 可以安全地复用同一 block 来
表示不同 creator。

精确私有条目使用 group 和 element 低字节作为键，不依赖数据集中实际分配的 block。
程序化 `Add` 会把带 creator 的条目路由到对应私有子字典，并在应用该字典的 creator
前克隆条目，避免修改调用方持有的对象。即使 XML 语法错误发生在 `<tag>` 内部，错误
信息仍会保留 creator、条目序号和标签上下文。

参考：[fo-dicom DicomDictionaryReader](https://github.com/fo-dicom/fo-dicom/blob/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/DicomDictionaryReader.cs)

**验收标准**

- 从 `io.Reader` 加载标准和私有字典。
- 支持精确标签和掩码标签、多个 VR、VM、keyword、retired 标志和 private creator。
- 拒绝异常输入，并提供元素级上下文。
- 明确定义重复条目和字典合并行为。

**验证证据**

- 本地 fo-dicom 2026b 的完整 `DICOM Dictionary.xml` 和
  `Private Dictionary.xml` 均成功加载；已验证 `MED NM` 的已知私有掩码标签可通过
  creator 子字典查询。
- 测试覆盖标准/私有布局、精确/掩码优先级、重复项覆盖、失败时原子合并、
  fo-dicom 的全部 VR 分隔符、备选 VM、BOM、组合源 UID 节点、异常 XML、
  XML 标签上下文、无效 VR/VM、缺失字段、私有精确标签 block 归一化、creator 不匹配
  拒绝、程序化私有条目路由、调用方条目不可变性和并发创建私有字典。
- Implicit VR 解析器测试覆盖 Private Creator 预留元素解码、私有 VR 查询，以及多个
  序列 Item 复用同一已分配 block 时的 creator 隔离。
- `ExampleNewFromXML` 由 Go 测试套件编译并执行。
- `go test ./pkg/dicom/dict -count=1` 通过。
- `go test -race ./pkg/dicom/dict -count=1` 通过。
- `go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1` 通过。
- `golangci-lint run` 报告 0 个问题。
- 已实现于提交 `6dc89f7`。

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
- 已在提交 `a9e4301` 中实现。

### PRINT-001: 基于 Dataset 的打印管理模型

**状态：** `Complete`
**优先级：** `P2`

已于 2026-08-16 完成。FilmSession、FilmBox、ImageBox 和 PresentationLUT
现在可通过相互独立的 Dataset 对象往返全部已支持属性。空 SOP Instance UID
使用 UUID 派生的 `2.25.*` 值。FilmSession 提供基于 UID 的创建/查找/删除、
修复父引用的递归独立克隆，以及使用 Explicit VR Little Endian 的 DICOM Part 10
加载/保存；原有索引读取和直接字段 API 保持兼容。

打印客户端会验证完整层级，并按确定顺序执行 Basic Print：Film Session N-CREATE、
Presentation LUT N-CREATE、Film Box N-CREATE、Image Box N-SET，最后执行
Film Session N-ACTION。新增的同步强类型 N-CREATE/N-SET/N-ACTION/N-DELETE
service API 复用现有消息 ID、待处理响应、取消和异步操作控制。工作流在第一个传输
错误或 DIMSE 失败状态处停止，不承诺远端回滚。

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

**验证证据**

- Dataset 往返覆盖 FilmSession、FilmBox、ImageBox 和 PresentationLUT 的所有
  已支持属性，并验证源 Dataset、图像字节和 LUT 切片相互独立。
- 测试覆盖唯一 UID 生成、克隆独立性、父引用修复、UID 查找/删除、重复拒绝、
  显示格式边界和 DICOM Part 10 FilmSession 文件往返。
- 真实 `net.Pipe` 客户端/服务端 service 对完整 N-CREATE/N-SET/N-ACTION 工作流
  进行编解码，并验证代表性 Dataset 值、SOP Class/Instance UID、操作顺序、
  首错停止和取消行为。
- `CGO_ENABLED=0 go test ./pkg/printing ./pkg/network/... -count=1` 通过。
- `CGO_ENABLED=0 go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`
  和 `CGO_ENABLED=0 go build ./...` 通过。build 输出一条非致命的 Windows 用户模块
  stat cache 访问警告，最终退出码为 0。
- `golangci-lint run --allow-parallel-runners` 报告 0 个问题，`git diff --check` 通过。

**互操作边界：** service-pair 测试验证了真实 go-dicom DIMSE 编解码。本次完成验证中，
未由独立 fo-dicom 进程交叉读取生成的 FilmSession 文件，也未让其参与打印关联。

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

**状态：** `Complete`
**优先级：** `P3`

已于 2026-08-15 完成。reconstruction 包现可将经典或 Enhanced CT/MR Dataset 展开为逐帧 `ImageData`，验证并排序不可变的 `VolumeData`，采样任意患者空间切面，创建惰性的标准平面 Stack，并流式生成经典 CT/MR 派生 DICOM 实例。

实现没有照搬被审计的 fo-dicom 框架：它会校验帧法向量，并发路径不共享可变搜索状态，允许最后一行/列和首末层的像素中心，包含 Stack 端点；显式允许不规则层距后，按真实相邻层距离插值。生成实例使用 Explicit VR Little Endian、新的 Series/SOP UID、`DERIVED\SECONDARY\MPR`、源帧引用、有效图像平面标签、16 位有符号或无符号像素，并使用 Pixel Padding Value 表示无效样本。经典输出会移除 Enhanced 功能组和维度标签。

完成后的走查还将源帧和 Volume 公共状态收紧为只读访问器，原生 Enhanced 多帧只创建并共享一个 Pixel Data 容器，并按帧校验 Dimension Index Values。重建只接受一个 Stack ID；非空间维度可以声明，但必须在全部帧中保持不变，多 Stack 以及变化的时间、扩散、心动周期维度会被拒绝。生成实例的 Modality 由输出 SOP Class 强制确定；CT MONOCHROME1 会被拒绝，MR 合法的 MONOCHROME1 极性保持不变。源数据、Cut 和 Stack 的 checked 上限可阻止整数溢出及无界预分配，生成的浮点标签使用不超过 16 字符的合法 DICOM DS 表示。

参考：[fo-dicom Reconstruction](https://github.com/fo-dicom/fo-dicom/tree/7ea6d424d0b0e11ecf6a55e81a8ac58b05d5e3e2/FO-DICOM.Core/Imaging/Reconstruction)

本条目依赖 IMG-001 和 IMG-002；两个前置项现均已完成。

**验收标准**

- 只从几何兼容的切片构建体数据。
- 对切片排序，并检测不规则间距和 Frame of Reference 不匹配。
- 以定义明确的插值和体外行为生成轴位、冠状位、矢状位和任意切面。
- 生成具有有效几何、派生元数据、UID 和像素表示的派生 DICOM 实例。
- 明确定义大型检查的内存和并发行为。

**验证**

- 合成测试覆盖排序、兼容性错误、包含端点的边界、双线性与真实层距插值、Padding 掩码、取消和确定性 worker 数。
- 轴位、冠状位、矢状位和任意切面均有可预测像素值测试，包括不规则层距和 Enhanced 多帧输入。
- 派生 CT/MR Dataset 会再次经过 writer/parser 往返，验证 Transfer Syntax、元数据、源引用和代表性像素值。
- 仓库的 `TestMultiFrame.dcm` 与 fo-dicom 的 `GH1876.dcm` 字节完全一致；集成回归测试覆盖其 7 帧几何、stored 代表像素、层距、参考 Cut、modality-space rescale 以及派生序列往返。
- 使用流式生成器时，内存上界为已解码源帧加一个已物化输出切片。

**有意边界：** IMG-003 不生成 Enhanced 多帧输出；对于彩色、浮点像素、非 CT/MR、方向不兼容和不受支持的多维输入，会显式拒绝，而不是生成语义不明确的结果。

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

当前进度：NET-001、NET-002 和 STD-001 已完成；DOC-001 仍为部分完成，
因此 Phase 0 尚未完成。

阶段验收：

- README 示例可编译，能力描述与实际公共 API 一致。
- 高层 SCU 可以通过普通 TCP 和经验证的 TLS 建立连接。
- C-STORE 使用协商后的传输语法发送，或执行经过验证的转码。
- 可以从干净检出复现标准生成数据，且不存在差异。
- 聚焦测试、完整包测试、变更后的共享/网络代码竞态测试和相关互操作检查全部通过。

### Phase 1: 主要领域能力对标

范围：MED-001、NET-003、NET-004、SR-001、IMG-001、CORE-001。

当前进度：MED-001、NET-003、NET-004、SR-001、IMG-001 和 CORE-001 均已完成，
因此 Phase 1 已完成。

阶段验收：

- DICOMDIR 创建和往返可在两个库之间工作。
- 高级关联值得到协商，并由客户端行为实际执行。
- 所有已声明 SR Value Type 均可通过 Dataset 和文件工作流往返。
- 渲染使用 Dataset 元数据并生成经过验证的输出。
- Dataset 验证报告嵌套路径并验证 sequence 子项。

每个条目都应独立实现和发布；Phase 1 不是一个单独的拉取请求。

### Phase 2: 支撑 API 与运维能力

范围：IMG-002、CORE-002、DICT-001、ANON-001、PRINT-001、OBS-001。

当前进度：IMG-002、CORE-002、DICT-001、ANON-001 和 PRINT-001 已完成；
OBS-001 尚未完成，因此 Phase 2 尚未完成。

阶段验收：

- 在开始 MPR 工作之前稳定共享几何和遍历器 API。
- 运行时字典和匿名化配置具备严格的解析测试。
- 打印对象通过 Dataset 和 DIMSE N-service 工作流往返。
- 网络诊断可注入、结构化，且默认静默。

### Phase 3: 专用工作流

范围：IMG-003 和 MED-002。

当前进度：IMG-003 已完成；MED-002 仍为 Open，因此 Phase 3 尚未完成。

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
- 保持计划开发顺序中的状态最新，已完成行不得重新编号。只有依赖或范围证据发生变化时
  才能调整顺序，并在排序表中记录原因。
