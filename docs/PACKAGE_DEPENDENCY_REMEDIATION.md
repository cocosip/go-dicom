# go-dicom 包依赖架构走查与修复总纲

**文档状态：** 阶段 1–8 已按“唯一所有者、唯一入口、功能守恒”原则完成；`go-dicom-codecs` 外部集成待本仓发布后验证

**维护方式：** 仓库长期架构文档；不属于 `docs/superpowers` 临时规格或实施计划

**审计范围：** `pkg/...` 下全部生产包，以及会进入生产编译集合的辅助代码

## 1. 文档目的

本文档合并以下两次走查结论：

1. `pkg/imaging/codec` 专项走查，包括 codec SPI、Pixel Data 抽象、转码编排、
   registry 和外部 codec 扩展边界。
2. `go-dicom` 全仓包依赖走查，包括父包与子包、跨领域依赖、潜在循环依赖、
   为规避循环而引入的接口包，以及测试代码对生产包的污染。

本文档不是逐文件实施步骤。它定义稳定的问题 ID、问题证据、目标依赖方向、修复方案、
迁移约束和验收标准。进入任一实施阶段前，应以当时源码重新确认文件位置并生成该阶段的
具体实施清单。

## 2. 审计基线

审计日期：2026-09-10。

| 项目 | 值 |
| --- | --- |
| 仓库 | `github.com/cocosip/go-dicom` |
| 提交 | `3ced84a` |
| 生产包数量 | 47 |
| 仓库内直接依赖边 | 178 |
| 编译级 import cycle | 0 |

执行过的验证：

```powershell
go list -mod=readonly ./pkg/...
go test ./pkg/... -count=1
git diff --check
git status --short
```

上述命令在原始审计时均通过。当前重新审查已再次确认提交、包数和依赖边数量；工作区没有
生产代码改动，仅有本整改文档处于未跟踪状态。测试通过只证明当前实现可编译且现有测试覆盖内
行为正常，不能否定本文记录的初始化顺序、职责漂移、隐式依赖获取和潜在循环依赖风险。

## 3. 范围与非目标

### 3.1 本文范围

- `pkg/dicom`、`pkg/imaging`、`pkg/io`、`pkg/network`、`pkg/media`、
  `pkg/printing`、`pkg/sr` 和 `pkg/logging` 的全部包依赖。
- codec 接口作为 `go-dicom` 对外提供的扩展契约。
- Dataset 级转码、C-STORE 自动转码和 reconstruction 解码的所有权。
- Registry 的生命周期、注册契约、隐式依赖获取和只为规避循环而存在的类型包。
- 公共 API 和包路径的所有权调整，以及仓库内外调用方的同步迁移。

### 3.2 明确非目标

- 不重写 `go-dicom-codecs` 中 JPEG、JPEG-LS、JPEG 2000、HTJ2K、RLE 等压缩算法；但
  Codec SPI 变化必须联动迁移该仓库的适配代码、注册入口、参数类型、示例和测试。
- 不因为某个类型是 interface 就一律删除接口。
- 不在一次提交中完成全部迁移；每个阶段必须可独立测试和回退。
- 不把目录层级机械等同于 Go 的语言层级。父目录与子目录只有在职责和依赖方向冲突时
  才构成架构问题。

### 3.3 迁移硬约束

1. 一个概念只能有一个公共所有者和一套生产实现；不得用 type alias、变量别名、转发函数、
   deprecated wrapper 或 legacy adapter 在旧包继续暴露第二入口。
2. 允许调整公开 API 和 import path。每个阶段必须同步迁移本仓库全部调用方；涉及 Codec SPI
   时按 3.5 的正式版本发布链后续迁移 `go-dicom-codecs`，不能用旧接口兜底。
3. 本轮包整改不授权删除任何现有功能。所有已导出的类型、函数、方法、字段、接口及其可观察
   行为都按有效公共能力处理；不得以“仓库内没有调用”“只有测试调用”或“可以简化 API”为由
   移除。若以后确需废止某项功能，必须脱离本整改另行走查并取得明确确认。
4. 旧包或旧 API 只允许在完整能力迁移后退役，退役的是错误路径，不是功能。每个受影响的导出
   标识符都必须记录旧入口、行为语义、新所有者、新入口、调用方迁移和回归测试；任一项无法
   映射时不得退役旧定义。
5. 每个阶段结束时，旧入口、兼容层和重复实现必须在同一阶段删除；不得把清理推迟到未定义的
   “下一次大版本”。
6. 若某项能力无法在当前阶段完成调用方联动迁移，该阶段不得合入，而不是临时保留双入口。

### 3.4 功能守恒门禁

每个实施阶段必须先从源码生成“公共能力迁移台账”，至少覆盖该阶段涉及包的全部导出标识符，
不能只统计当前仓库内引用。台账的每一行都必须包含：

1. 旧包路径和完整 API 名称。
2. 当前功能、错误语义、可变性、并发语义和必要副作用。
3. 唯一的新包路径和新 API 名称。
4. 仓库内调用方、示例以及 `go-dicom-codecs` 等受控外部调用方的迁移位置。
5. 证明行为仍然存在的迁移前快照或回归测试。

只有台账全部映射、调用方全部切换且行为回归通过后，才能移除旧路径、重复入口或重复实现。
删除 alias、wrapper 和旧包不能作为功能完成的证据；新的唯一入口必须能够实际完成原功能。

### 3.5 跨仓发布与验证顺序

`go-dicom-codecs` 依赖 `go-dicom`，因此 Codec SPI 不能要求下游在未发布源码上提前完成正式集成。
涉及两仓库的阶段按以下顺序执行：

1. 在 `go-dicom` 完成公共能力台账、新 Codec SPI、native codec、Registry、transcode 和本仓全部
   调用方迁移，以本仓单元测试、契约测试、build、vet 和 lint 验证后发布正式版本。
2. 发布前不执行 `go-dicom-codecs` 集成测试，也不使用本地 `replace`、`go.work` 或复制源码把
   未发布版本伪装成正式依赖；此时状态只能记录为“go-dicom 已完成，外部 codec 集成待发布后
   验证”，不能宣称跨仓整改全部完成。
3. `go-dicom-codecs` 引用已发布的 `go-dicom` 版本后，迁移各 codec 实现、强类型参数、注册入口、
   示例和测试，并执行 JPEG、JPEG-LS、JPEG 2000、HTJ2K、RLE、blank import 与 Dataset
   transcode 集成回归。
4. `go-dicom-codecs` 验证通过并发布后，再把跨仓项目标记为完成。依赖方向始终保持
   `go-dicom-codecs -> go-dicom`，`go-dicom` 不得为了验证而反向引用下游模块。

## 4. 当前依赖结构

### 4.1 领域级依赖

```text
media ---------------------> dicom
sr ------------------------> dicom
printing ------------------> dicom
printing ------------------> network/dimse

network -------------------> dicom
network/service -----------> imaging/codec       [跨领域编排泄漏]

imaging -------------------> dicom
imaging -------------------> io/buffer
imaging/reconstruction ----> imaging              [子包反向依赖根包]

dicom ---------------------> io/buffer
io/buffer -----------------> dicom/endian         [领域级双向依赖]
```

当前没有编译器可见的循环，但存在两个需要优先消除的结构风险：

1. `dicom -> io/buffer -> dicom/endian` 已形成领域级双向依赖。
2. 根 `imaging` 导入多个 imaging 子包，而 `imaging/reconstruction` 又导入根
   `imaging`。这会阻止根包将来聚合 reconstruction，并使公共类型无法向下沉降。

### 4.2 全部包走查结果

`out` 是该包直接导入的仓库内部包数量，`in` 是直接导入该包的仓库内部包数量。

| 包 | out | in | 结论 |
| --- | ---: | ---: | --- |
| `dicom` | 0 | 0 | 仅作为根入口/文档包，正常 |
| `dicom/anonymizer` | 6 | 0 | 上层业务包依赖 Dataset 基础能力，正常 |
| `dicom/charset` | 0 | 3 | 叶子能力，正常 |
| `dicom/dataset` | 9 | 15 | 核心聚合包，依赖较多但职责基本合理 |
| `dicom/dataset/rules` | 7 | 0 | 子包扩展父包且父包不回调 rules，方向合理 |
| `dicom/daterange` | 0 | 2 | 叶子能力，正常 |
| `dicom/dict` | 4 | 4 | 与 `tag/dictif` 形成运行时概念环，需修复 |
| `dicom/dictif` | 0 | 2 | 专门规避循环的服务定位器，应退役 |
| `dicom/element` | 6 | 15 | 核心抽象合理，但受 `io/buffer` 分层影响 |
| `dicom/endian` | 0 | 7 | 内容通用却位于 DICOM 域，导致 io 反向依赖 |
| `dicom/parseable` | 0 | 3 | 小型叶子抽象，合理 |
| `dicom/parser` | 10 | 6 | 上层解析编排，依赖规模与职责相符 |
| `dicom/serialization` | 6 | 0 | 上层序列化包，字典查询路径需调整 |
| `dicom/tag` | 1 | 18 | 应是基础叶子，当前被 `dictif` 污染 |
| `dicom/testutil` | 0 | 0 | 误放在公开路径的仓库测试辅助包，应删除且不新增共享导出 helper |
| `dicom/transfer` | 3 | 13 | 基础能力合理；默认 registry 契约可改进 |
| `dicom/uid` | 1 | 10 | 基础能力合理；标准目录由可变全局注册构建，需去副作用 |
| `dicom/vm` | 0 | 1 | 叶子能力，正常 |
| `dicom/vr` | 0 | 16 | 叶子能力，正常 |
| `dicom/writer` | 9 | 4 | 上层写入编排，依赖规模与职责相符 |
| `imaging` | 15 | 1 | 根包过重，并被 reconstruction 反向依赖 |
| `imaging/codec` | 8 | 3 | SPI 与 Dataset 转码混合，未成为叶子扩展点 |
| `imaging/geometry` | 5 | 1 | Dataset 几何适配层，方向合理 |
| `imaging/imagetypes` | 0 | 4 | 循环规避型杂物包，应按所有权拆分 |
| `imaging/interpolation` | 0 | 2 | 数学叶子能力，正常 |
| `imaging/lut` | 1 | 2 | 所有权方向合理，但与 render 有重复实现 |
| `imaging/math3d` | 0 | 5 | 数学叶子能力，正常 |
| `imaging/reconstruction` | 10 | 0 | 依赖根 imaging，形成潜在循环屏障 |
| `imaging/render` | 5 | 1 | 渲染编排合理，但重复持有颜色和 LUT 算法 |
| `imaging/transform` | 1 | 2 | 依赖 math3d 的叶子变换能力，正常 |
| `io/buffer` | 1 | 7 | 反向依赖 `dicom/endian`，分层错误 |
| `io/rangehttp` | 1 | 0 | `rangeio` 的 HTTP adapter，方向合理 |
| `io/rangeio` | 0 | 1 | 定义消费方 `Fetcher` 端口，方向合理 |
| `logging` | 0 | 6 | 独立叶子包，正常 |
| `media` | 8 | 0 | DICOMDIR 上层领域包，依赖方向合理 |
| `media/scanner` | 2 | 0 | 独立扫描入口，不反向依赖 media，合理 |
| `network/association` | 4 | 3 | 对 `dimse` 的 MessageID 依赖层级不合理 |
| `network/client` | 9 | 0 | 上层入口合理，但内部 service 接口过宽 |
| `network/dimse` | 5 | 5 | DIMSE 消息模型层，方向基本合理 |
| `network/observability` | 1 | 3 | 只依赖 logging，合理 |
| `network/pdu` | 1 | 5 | 协议叶子基本合理，但生产代码导入 testing |
| `network/server` | 6 | 0 | 上层入口，依赖方向合理 |
| `network/service` | 14 | 2 | 网络核心编排过重，并直接读取全局 codec manager |
| `network/status` | 0 | 3 | 叶子能力，正常 |
| `network/transport` | 1 | 3 | 依赖 PDU 的传输适配合理；测试证书代码位置错误 |
| `printing` | 10 | 0 | 上层领域包，消费窄 DIMSE 接口，合理 |
| `sr` | 8 | 0 | 上层领域包，只依赖 DICOM 基础能力，合理 |

## 5. 接口分类原则

本轮不采用“接口越少越好”或“为了避免 import 就定义接口”的单一判断方式。接口合理与否
按以下标准判断：

1. 接口由消费方还是实现方定义。
2. 接口是否通过构造函数或 option 显式注入。
3. 接口是否足够窄，并表达稳定的业务能力。
4. 生产代码是否真的存在多个实现或隔离边界。
5. 接口是否隐藏了全局初始化顺序或类型断言。

以下接口应保留：

- `media.IconGenerator`：消费方定义、单方法、显式注入，用于保持 media 不依赖 imaging。
- `printing.DIMSEService`：消费方定义、只包含打印所需的四个 N-Service 方法。
- `rangeio.Fetcher`：消费方端口，`rangehttp.Fetcher` 是清晰的 adapter。
- `dataset/rules` 的 Match、Transform 和 UID generator：正常扩展点。
- network observer、handler、operation、PDU encoder 和 listener 接口：协议扩展或边界测试端口。
- `element.Element`、`buffer.ByteBuffer`、`parseable.Parseable`：稳定的领域基础抽象。

以下接口或接口包需要调整：

- `dictif.Lookup`：通过全局 setter/getter 工作，是服务定位器而非依赖注入。
- `imagetypes.PixelData`：同时表达输入、输出、元数据和可变状态，职责过宽。
- `network/client.serviceInterface`：约 20 个方法，只为内部 mock 聚合全部 Service 能力。
- 三个 `cFind/cMove/cGetErrorService`：让同一调用在不同实现下产生不同终态错误语义。

## 6. 总体修复策略

### 6.1 不采用：只修补当前 import 边

只移动 `endian`、测试辅助文件和少量接口，不调整 codec、transcoder 和全局 registry。

优点是改动小；缺点是 `dictif`、`imagetypes`、network 自动转码和 imaging 根包反向依赖仍然
存在，后续功能仍会再次遇到同类循环问题。该方案不推荐。

### 6.2 不采用：分阶段迁移并保留兼容层

建立明确叶子包，将 Dataset 转码提升为独立编排层，但旧公开类型通过 alias、deprecated
wrapper 或 legacy adapter 暂时保留。

该方案表面降低迁移成本，实际会长期保留新旧两套入口，模糊所有权，并使仓库内外调用方没有
完成迁移的强制边界。本文不采用。

### 6.3 采用：按所有权分阶段、阶段内一次性切换

先按依赖顺序拆成可独立验证的阶段；每个阶段内先建立行为回归，再迁移实现和全部调用方，最后
删除旧包、旧 API、兼容入口和重复实现。阶段之间允许小步实施，阶段内部不保留双入口。

该方案允许公开 API 发生必要变化，但要求功能守恒。Codec SPI 先由 `go-dicom` 完成并发布，
再由 `go-dicom-codecs` 引用正式版本适配；压缩算法本身不重写，只适配唯一的新契约。本文后续
方案均按此策略设计。

### 6.4 现有 alias 与 compatibility 入口的处理边界

源码中已经存在的 alias、deprecated wrapper 和 legacy 双入口也必须随其所有者阶段收敛，不能只
禁止本轮新增。当前逐项审计得到以下明确映射：

| 当前重复入口 | 唯一入口 | 处理阶段与功能守恒要求 |
| --- | --- | --- |
| `dataset.ValidationPathSegment` | `dataset.PathSegment` | 阶段 6 更新 `ValidationError.Path` 和内部验证签名后移除 type alias；路径结构、clone 和格式化行为不变 |
| `Tag.Uint32()` | `Tag.ToUint32()` | 阶段 6 迁移全部调用后移除纯转发方法；group/element 的 32 位编码不变 |
| `Client.Ping(ctx)` | `Client.CEcho(ctx)` | 阶段 7 迁移调用后移除纯转发方法；连接校验仍执行同一 C-ECHO 请求和错误路径 |
| `service.WithDIMSETimeout` | `WithHandlerShutdownTimeout` | 阶段 7 按旧函数的真实行为迁移；出站响应空闲超时继续由独立的 `WithRequestTimeout` 配置，不保留含糊名称 |
| legacy `CFindHandler` 与 `CFindStreamHandler` | `CFindHandler func(context.Context, CFindOperation) error` | 阶段 7 直接把无后缀名称切换为唯一 operation 签名；旧 slice handler 的 pending/final/status/identifier 能力全部由 operation 发送，不保留 Stream/legacy 双入口 |
| `ExtendedNegotiation.ServiceClassAppInfo` | `RequestedApplicationInfo`、`AcceptedApplicationInfo` | 阶段 7 按协商方向迁移后移除含糊字段；request/accept 两个值及其独立 copy 语义均保留 |
| `JPEGProcess1`、`JPEGProcess2_4`、`JPEGProcess14`、`JPEGProcess14SV1`、`JPEG2000Lossy` | `JPEGBaseline8Bit`、`JPEGExtended12Bit`、`JPEGLossless`、`JPEGLosslessSV1`、`JPEG2000` | 阶段 8 迁移调用和生成的标准列表后移除变量 alias；每个 UID 及全部 transfer 属性仍只有一个标准值 |
| `imagetypes`、`render` 中的 LUT/VOI alias | `imaging/lut` 的实际定义 | 按阶段 3 和 DEP-013 迁移，不在旧包保留 alias |

“底层值相同”不等于“重复入口”。以下名称表达不同协议上下文，不能按 compatibility alias 机械
删除：C-ECHO、N-GET、N-SET、N-ACTION 等操作各自的 Success 状态名；`endian.Network` 表达
网络字节序；PDU request/accept 结构各自的 `ServiceClassAppInfo` 表达线上的单向字段。它们不是
旧 API 到新 API 的迁移桥。高层 convenience API 若组合了独立工作流，也必须先完成行为审计，
不能仅因内部调用另一个方法而认定为 alias。

## 7. 详细问题与修复方案

### DEP-001：`dictif` 形成运行时概念循环

**优先级：** P1
**性质：** 已确认的架构缺陷

#### 问题描述

`dict` 必须导入 `tag`，因为 Dictionary Entry 和查询键使用 `*tag.Tag`。`tag` 又希望提供
`DictionaryEntry()` 和 `ParseKeyword()`，因此通过 `dictif` 间接调用字典。`dict.init()`
把 Default Dictionary 写入 `dictif.globalLookup`，完成一次运行时回调。

这没有产生 Go import cycle，却带来以下问题：

- `tag` 的行为取决于应用是否间接导入过 `dict`。
- `Tag.DictionaryEntry()` 只能返回 `interface{}`，序列化包必须断言为 `*dict.Entry`。
- `dictif.SetGlobalLookup` 暴露可变进程全局状态。
- tag 测试把“字典未初始化时返回 nil”作为允许行为，掩盖了公共 API 的不确定性。
- private creator 路径从字典获取 creator 后又按字符串重新创建，没有获得对象缓存收益。

证据：

- `pkg/dicom/dictif/interface.go`
- `pkg/dicom/dict/dictionary.go`
- `pkg/dicom/tag/tag.go`
- `pkg/dicom/serialization/json.go`
- `pkg/dicom/serialization/xml.go`
- `pkg/dicom/serialization/json_reader.go`

#### 修复方案

1. 规定 `tag` 只负责 Tag 值、格式解析、比较、掩码和 private creator 值对象，不再查询字典。
2. 在 `dicom/dict` 定义面向字典消费者的 `Lookup` 契约，由 `Dictionary` 实现；标准路径使用
   `dict.Default()`，JSON/XML 的读写 option 增加显式 Lookup 注入，使原 `SetGlobalLookup` 承载的
   自定义字典能力不依赖进程全局状态也能继续使用。
3. JSON/XML writer 通过配置的 Lookup 查询 entry，未配置时使用 `dict.Default()`，不再调用
   `Tag.DictionaryEntry()`；JSON/XML reader 的 keyword 解析使用同一个 Lookup，保证读写选择
   同一字典。
4. `tag.Parse` 创建 private creator 时直接使用 `tag.NewPrivateCreator`；字典内部缓存只由
   Dictionary 自己管理。
5. 将 keyword 到 Tag 的查询能力唯一归属 `dict.Lookup.LookupKeyword`；
   `Tag.DictionaryEntry()` 映射为 `dict.Lookup.Lookup(tag)`，`tag.ParseKeyword()` 映射为
   `dict.Lookup.LookupKeyword(keyword)`，标准调用方使用 `dict.Default()`，自定义调用方使用自己
   持有的 Lookup。
6. `dictif.Tag`、`dictif.Entry` 和 `dictif.PrivateCreator` 的公开查询能力分别由规范的
   `tag.Tag`、`dict.Entry` 和 `tag.PrivateCreator` 类型承接；`dictif.Lookup` 的自定义实现能力由
   新 `dict.Lookup` 承接。逐项迁移后才退役 `dictif`、`dictionaryAdapter` 和全局 setter/getter。

#### 验收标准

- 仓库内外受控调用方均改用 `dict.Dictionary`，不再存在 `Tag.DictionaryEntry()` 或
  `tag.ParseKeyword()`。
- `tag` 不导入 `dictif`，导入 `tag` 的行为不再依赖其他包的 init 顺序。
- JSON/XML keyword round-trip 覆盖标准 tag、未知 tag、masked tag 和 private tag。
- `go list` 不再包含 `dicom/dictif`，且字典 lookup、keyword parse、private creator 功能均有
  新入口回归。
- 注入自定义 Lookup 的 JSON/XML 读写回归通过，证明移除全局 setter 没有删除自定义字典能力。

### DEP-002：`io/buffer` 反向依赖 `dicom/endian`

**优先级：** P1
**性质：** 已确认的分层缺陷

#### 问题描述

`dicom/dataset`、`dicom/element`、parser、writer 和 serialization 都依赖 `io/buffer`；
`io/buffer/endian.go` 又导入 `dicom/endian`。因此领域折叠后形成 `dicom <-> io`。

Endian、ByteOrder 和 SwapBytesN 都是通用二进制能力，不应由 DICOM 领域拥有。

#### 修复方案

1. 新建 `pkg/io/endian`，唯一拥有 `Endian`、`Little`/`Big`/`Network`、`ByteOrder` 和原始
   swap utilities；完整迁移 `String`、`IsBig`、`IsLittle`、`SwapUint16/32/64`、
   `SwapInt16/32/64`、`SwapBytes` 和 `SwapBytesN`。`Network` 使用常量，`LocalMachine` 的读取
   能力迁移为只读 `Native() Endian` 查询；原先通过改写该变量选择字节序的场景改用各类型已有
   的显式 `WithEndian` 构造入口，保留调用方选择字节序的能力而不再修改进程全局状态。
2. `EndianByteBuffer`、`SwapByteBuffer`、`NewEndian` 和 `NewSwap` 继续唯一归属 `io/buffer`，
   只调用 `io/endian` 的原始运算；不得把 buffer wrapper 复制到 `io/endian`。
3. `io/buffer` 改为只依赖 `io/endian`。
4. 一次性迁移本仓库全部导入到 `pkg/io/endian`，包括 parser、writer、element、transfer、
   printing、network 测试和示例。
5. 删除 `pkg/dicom/endian` 及其旧测试；测试迁移到新所有者，不保留 alias 或函数转发。

#### 验收标准

- `io/...` 不再导入任何 `dicom/...` 包。
- 仓库中不存在 `pkg/dicom/endian` 导入或兼容包，Endian 只有 `pkg/io/endian` 一个入口。
- 不存在可由调用方改写的 native/network byte order 全局变量；buffer wrapper 仍只有
  `io/buffer` 一套实现和入口。
- Big/Little Endian parser、writer 和 buffer swap 回归全部通过。
- `dicom/endian` 的每个导出常量、变量读取、方法和 swap 函数都有 `io/endian` 或显式 endian
  构造入口映射，不存在因包移动而缺失的能力。

### DEP-003：codec SPI 与 Dataset 转码编排混在同一包

**优先级：** P1
**性质：** 已确认的职责混合

#### 问题描述

`codec.Codec` 是供外部像素编解码器实现的低层 SPI；同一包中的 `Transcoder` 却直接操作
Dataset、Element、Tag、VR、ByteBuffer、File Meta 和 lossy metadata。这使 codec 包拥有
8 个内部依赖，无法成为稳定叶子扩展点。

#### 修复方案

1. `pkg/imaging/codec` 只保留 codec contract、frame contract、parameters contract、
   native codec 和 codec Registry。
2. 新建 `pkg/dicom/transcode`，承接当前 Dataset 级 `Transcoder`、Manager、File Meta 更新、
   VR 规范化、fragment/frame 组装和 lossy history。只保留以下构造路径和 context API：

   ```go
   func NewManager(registry *codec.Registry) (*Manager, error)
   func (m *Manager) CanTranscode(input, output *transfer.Syntax) bool
   func (m *Manager) NewTranscoder(input, output *transfer.Syntax, opts ...Option) (*Transcoder, error)

   func (t *Transcoder) InputSyntax() *transfer.Syntax
   func (t *Transcoder) OutputSyntax() *transfer.Syntax
   func (t *Transcoder) Transcode(ctx context.Context, ds *dataset.Dataset) (*dataset.Dataset, error)
   func (t *Transcoder) TranscodeWithMetadata(ctx context.Context, ds *dataset.Dataset, meta *dataset.FileMetaInformation) (*dataset.Dataset, *dataset.FileMetaInformation, error)
   func (t *Transcoder) DecodeFrame(ctx context.Context, ds *dataset.Dataset, frame int) ([]byte, error)

   func WithInputParameters(params codec.Parameters) Option
   func WithOutputParameters(params codec.Parameters) Option
   func WithStrictDICOMVR(strict bool) Option
   ```

   Registry 不得为 nil，`Transcoder` 不提供绕过 Manager 的公开构造函数。
3. 明确 `dicom/transcode` 是高层集成包，可以依赖 `dicom/dataset`、`imaging/codec` 和
   `imaging/pixeldata`；基础包不得反向导入它。
4. 新包不保留无 context wrapper，也不同时提供 `Transcode`/`TranscodeContext` 等重名路径；
   Dataset、File Meta 和单帧解码三种现有功能分别由上述唯一方法承接。
5. 现有 `CanTranscode`、输入/输出 syntax 查询、输入/输出参数和 strict DICOM VR 选择能力按上述
   API 迁移。`WithCodecRegistry` 由 Manager 构造参数取代；`WithInputCodec`/`WithOutputCodec`
   的自定义实现能力由调用方在目标 Registry 上 Register/Replace 后创建 Manager 承接，不再让
   单个 Transcoder 绕过 Registry 形成第二条 codec 选择路径。
6. 先迁移本仓库调用方到 `dicom/transcode`，再从 codec 退役 `Transcoder`、
   `TranscoderManager`、`NewTranscoder`、`GetDefaultManager` 及 Dataset 转码文件；codec 不保留
   wrapper 或旧实现。`go-dicom-codecs` 中涉及这些入口的示例或测试按 3.5 在正式版本发布后迁移。
7. NativeCodec 仍由 codec 唯一拥有，完整保留通用构造函数、三种内置 transfer syntax 构造函数、
   Name、TransferSyntax、DefaultParameters、Encode/Decode，以及按 codec endian 读写 uint16/uint32
   的公开能力。`StripTrailingPadding` 继续作为 encapsulated frame 公共辅助函数保留在 codec。
8. `ConvertEndianness` 的复制返回、2/4/8-byte sample 交换和错误语义完整迁入 `io/endian` 的唯一
   公开入口；它不应因移出 codec 而消失，也不在两个包同时保留 wrapper。

#### 验收标准

- codec SPI 文件不再导入 Dataset、Element、Tag、VR、writer 或 logging。
- Dataset metadata 变更只出现在 `dicom/transcode`。
- `codec` 中不存在 Dataset 级 Transcoder 类型、构造函数或转发入口。
- `dicom/transcode` 不存在绕过 Manager 的构造入口或无 context 的转码 wrapper。
- `CanTranscode` 仍供 C-STORE presentation context 选择使用；syntax、参数和 strict VR 行为均有
  迁移前后回归。
- codec 可单独使用内存 frame source/sink 完成 encode/decode 测试。
- native-to-native、native-to-encapsulated、encapsulated-to-native 和压缩到压缩路径保持一致。
- NativeCodec 的全部构造、数值读写、frame copy/swap 和 trailing padding 辅助能力均有新契约下
  的回归；`ConvertEndianness` 在 `io/endian` 保持迁移前行为。

### DEP-004：`imagetypes` 是循环规避型杂物包

**优先级：** P1
**性质：** 已确认的所有权不清

#### 问题描述

`imagetypes` 同时包含：

- codec 使用的 PixelData、FrameInfo 和 FrameInfoSetter。
- render/lut 使用的 LUT、VOILUTFunction 和 Color32。
- 当前只在自身测试中直接使用、但提供完整公开运算能力的 BitDepth。

根 `imaging` 还直接定义 PixelRepresentation、PlanarConfiguration、ColorSpace 和
PhotometricInterpretation。`pixeldata` 下沉后如果继续引用这些根包类型，就会再次形成
`imaging/pixeldata -> imaging` 的反向依赖，因此只拆 `imagetypes` 而不处理这些像素格式值对象
并不能完成分层。

这些类型没有共同生命周期，只因为“多个 imaging 包都能导入”而聚集。包注释也明确说明
其目的是避免循环依赖，这说明它是结构问题的结果，不是稳定领域抽象。

#### 修复方案

1. 新建独立叶子包 `pkg/imaging/pixel`，只拥有像素格式值对象：`BitDepth`、
   `Representation`、`PlanarConfiguration`、`Component`、`ColorSpace`、
   `PhotometricInterpretation`，以及这些类型现有的构造、解析、比较、枚举值和查询方法。
   该包不得包含 Dataset 提取、帧存储、codec、LUT、渲染或颜色转换算法，防止它变成新的
   `imagetypes` 杂物包。
2. `BitDepth` 是 `imaging/pixel` 对外提供的通用像素能力，不属于 codec。将 `BitDepth` 和
   `NewBitDepth` 的唯一实现迁入该包，完整保留 `BytesAllocated`、`IsValid`、`Mask`、
   `SignMask`、`ExtendSign`、`MinimumValue`、`MaximumValue` 和 `Range`，不因当前只有自身测试
   直接调用而删减。
3. `FrameInfo` 和方向明确的 frame 端口迁入 `imaging/codec`；FrameInfo 只消费
  `imaging/pixel` 的格式值对象，不在 codec 内重新定义位深、signed/unsigned、planar 或
   photometric 语义。
4. `Color32` 与 DICOM YBR/RGB 数值转换统一归属独立叶子包 `imaging/colorconv`；完整保留
   `NewColor32`、ARGB 打包和全部现有转换行为，`lut`、`render` 与 `pixeldata` 都只消费该唯一
   实现。
5. `LUT`、`VOILUTFunction` 及其全部枚举值迁入 `imaging/lut`，render 只消费 lut 契约。
6. 按 3.4 为 `imagetypes` 和被下沉的根 imaging 类型建立逐 API 迁移台账；本仓库引用和测试
   全部切换后退役旧定义及 `imagetypes` 包，不保留类型别名、转发或第二套实现。
   `go-dicom-codecs` 的引用和测试按 3.5 在正式版本发布后迁移并完成第二段验收。

#### 验收标准

- 每个共享类型都有唯一所有者。
- `pixel.BitDepth` 的构造、字段和全部公开计算能力与迁移前逐项对应；codec 只消费该类型，
  `FrameInfo`/`PixelDataInfo` 不存在第二份位深算法。
- `pixel` 的依赖图保持叶子化，且不吸收 frame I/O、Dataset、codec、LUT、render 或转换职责。
- Representation、PlanarConfiguration、ColorSpace 和 PhotometricInterpretation 的现有
  字符串、解析、比较、枚举值及颜色空间关联能力均由 `imaging/pixel` 的唯一入口承接。
- `go list` 不再包含 `imaging/imagetypes`，仓库内外受控调用方均使用唯一所有者。
- `lut` 与 `render` 不再相互复制接口定义。

### DEP-005：`reconstruction` 反向依赖 imaging 根包

**优先级：** P1
**性质：** 潜在循环屏障

#### 问题描述

根 `imaging` 导入 codec、lut、render、transform 等多个子包；`imaging/reconstruction`
为了使用 `DicomPixelData` 和 `CreatePixelData` 又导入根 `imaging`。当前根包没有导入
reconstruction，所以编译通过，但根包已经无法自然聚合 reconstruction API。

#### 修复方案

1. 新建 `pkg/imaging/pixeldata`，迁移 PixelDataInfo、DicomPixelData、Dataset 像素提取和
   Element 转换。新构造入口为 `pixeldata.New`、`pixeldata.NewFromBytes` 和
   `pixeldata.FromDataset`；删除根 imaging 下的旧构造函数，不保留同名转发。
2. reconstruction 直接依赖 `imaging/pixeldata`；需要 Dataset 解码时依赖
   `dicom/transcode`，不再依赖根 imaging。
3. `DicomPixelData`、`PixelDataInfo`、`CreatePixelData` 和 Pixel Element 转换入口只在
   `imaging/pixeldata` 提供；从根 imaging 删除同名类型和函数，不保留 alias 或转发。
4. 根 imaging 只保留自身拥有的高层 `DicomImage` 创建和渲染编排；它可依赖叶子功能包，但
   不再作为子包 API 的二次导出 façade，任何 imaging 子包也不得反向导入根 imaging。
5. DicomPixelData 的所有现有行为随实现整体迁移：帧读取/追加/计数/合并、sample 读取、BOT、
   encapsulated 判断、Pixel Element 写回、interleaved/MONO/YBR 转换、padding 判断与 mask、
   min/max、自动窗口、VOI LUT/window 到 8-bit，以及 codec encode/decode。Encode/Decode 将
   context 作为首参数并只调用 DEP-007 的唯一 Codec API，不保留无 context 方法。
6. PixelDataInfo 保留 frame size、total size 和完整校验能力；其中位深数据使用 DEP-004 唯一的
   `pixel.BitDepth`，signed/unsigned、planar、photometric 和 color space 同样使用
   `imaging/pixel` 的值对象。PixelDataInfo 的 Validate 只增加帧数、采样、photometric、VR、
   压缩和 padding 等高层约束，不复制底层格式校验算法。
7. 类型移动不能漏掉公开数据面：`DicomPixelData.Info` 的读取和更新能力，以及 PixelDataInfo 的
   Width、Height、NumberOfFrames、BitsAllocated/BitsStored/HighBit、SamplesPerPixel、
   PixelRepresentation、PlanarConfiguration、PhotometricInterpretation、VRCode、Encapsulated、
   TransferSyntaxUID、lossy metadata 和 padding 字段必须逐项映射到新类型；若改为受检 getter/
   setter，必须证明原先可表达的每一种更新仍可完成。
8. 方法台账必须逐项覆盖 `GetFrame`、`GetSample`、`IsPaddingSample`、`CalculateOptimalWindow`、
   `AddFrame`、`GetAllFrames`、`FrameCount`、`EnsureInterleaved`、
   `ConvertMonochrome1ToMonochrome2`、`ConvertYBRToRGB`、`WindowTo8bit`、`MinMax`、
   `MaskPadding`、`WindowOrLUTTo8bit`、`ToElement`、`IsEncapsulated`、`BasicOffsetTable`、
   `GetFrameInfo`、`Encode` 和 `Decode`。允许按 Go 命名收敛 Get 前缀并增加 context/error，但不得
   把未列入 FrameSource/FrameSink 的高层方法误判为可删除。

#### 验收标准

- `rg` 检查 `pkg/imaging/*` 子包不再导入 `pkg/imaging` 根路径。
- reconstruction 的现有几何、volume、MPR 和 DICOM generator 测试全部通过。
- 根 imaging 和 pixeldata 不存在同名公共类型或构造入口；像素提取、修改、写回和多帧功能
  均在 `imaging/pixeldata` 通过迁移测试。
- 原 DicomPixelData/PixelDataInfo 的每个公开行为均有新入口映射和回归，不以“当前无生产调用”
  为理由删除公开功能。

### DEP-006：PixelData 将 source、sink 和可选 metadata 混为一体

**优先级：** P1
**性质：** codec 公共契约缺陷

#### 问题描述

`imagetypes.PixelData` 同时要求 GetFrame、AddFrame、FrameCount、GetFrameInfo 和
IsEncapsulated。输入对象不需要 AddFrame，输出对象不应要求 GetFrame；输出 metadata 又通过
可选 `FrameInfoSetter` 和运行时 type assertion 写回。

此外，`GetFrame() []byte` 没有说明返回值是只读借用还是独立副本。现有 DicomPixelData 和
simplePixelData 都直接返回内部切片，因此 codec 可以意外修改调用方数据。

#### 修复方案

在 `imaging/codec` 中定义方向明确的新端口：

```go
type FrameSource interface {
    FrameCount() int
    Frame(ctx context.Context, index int) ([]byte, error)
    FrameInfo() FrameInfo
    Encapsulated() bool
}

type FrameSink interface {
    AddFrame(ctx context.Context, frame []byte) error
    SetFrameInfo(info FrameInfo) error
}
```

`FrameInfo` 使用 DEP-004 的 `pixel.BitDepth`、`pixel.Representation`、
`pixel.PlanarConfiguration` 和 `pixel.PhotometricInterpretation` 表达格式元数据；codec 不拥有
这些通用像素值对象，也不再用无类型的 `uint16`/`string` 复制其语义。

契约要求：

1. `Frame` 返回由调用方拥有的独立副本。Go 的 `[]byte` 无法表达只读借用，因此不能把
   “codec 不得修改”当作内存隔离保证；实现必须通过复制保证 codec 修改返回切片时不影响源数据。
2. `AddFrame` 的实现必须复制输入，调用返回后不得继续持有调用方可变内存。
3. `Frame`/`AddFrame` 必须在底层读取、复制和写入期间检查 context；不能只在 codec 外层调用
   前后检查。
4. FrameInfo 按值返回和传递，避免共享指针元数据。
5. 输出 metadata 成为 FrameSink 必选能力，不再通过可选 type assertion 静默忽略。输出 sink
   以输入 FrameInfo 初始化，codec 只在编码/解码改变 metadata 时显式调用 `SetFrameInfo`。
6. Codec 直接使用 `FrameSource`/`FrameSink`；迁移全部实现和调用方后删除旧 `PixelData`、
   `FrameInfoSetter`、`SetFrameInfo` 和所有 adapter。

#### 验收标准

- 输入 codec 无法通过接口调用输出方法，输出端不会被要求提供无意义读取能力。
- 增加恶意 codec 回归：修改输入 frame 后，原始 PixelData 不变。
- metadata 写回失败必须返回错误，不能返回 false 后继续成功。
- `imagetypes.PixelData` 和兼容 adapter 不再存在；同一帧读写能力没有第二套接口。

### DEP-007：codec 操作不可由 context 取消

**优先级：** P1
**性质：** 已确认的行为缺口

#### 问题描述

`TranscodeContext` 的 context 只传给 logging；`Codec.Encode/Decode` 不接收 context，
transcoder 内也没有检查 `ctx.Err()`。大图、多帧或 native codec 调用开始后无法响应取消。

#### 修复方案

将 context 直接纳入唯一的 Codec 契约：

```go
type Codec interface {
    Name() string
    TransferSyntax() *transfer.Syntax
    DefaultParameters() Parameters
    Encode(context.Context, FrameSource, FrameSink, Parameters) error
    Decode(context.Context, FrameSource, FrameSink, Parameters) error
}
```

1. 将 `Encode`/`Decode` 直接改为 context 首参的唯一签名；不保留旧无 context 签名，不增加
   `EncodeContext`/`DecodeContext` 方法，也不新增 `ContextCodec` 次级接口。
2. NativeCodec 在 `go-dicom` 发布前实现唯一契约并在帧循环和大块 copy/swap 边界检查取消；
   `go-dicom-codecs` 的所有 codec 按 3.5 在正式版本发布后实现相同契约。
3. `dicom/transcode` 只调用 context 方法，不做动态 type assertion 或旧 Codec fallback。
4. 错误必须保留 `context.Canceled` 或 `context.DeadlineExceeded`，支持 `errors.Is`。

#### 验收标准

- 已取消 context 不得开始 codec 工作。
- 多帧 encode/decode 可在帧边界稳定退出。
- network 请求取消不会继续在后台执行无界转码。
- 仓库中不存在无 context Codec 接口、`ContextCodec` 分支或 legacy adapter。

### DEP-008：Parameters 是可变的 string-any 参数袋

**优先级：** P2
**性质：** 公共契约设计债务

#### 问题描述

`codec.Parameters` 通过字符串读取和写入 `interface{}`；BaseParameters 内部是无复制、无并发
保护的 map。参数名拼写、值类型和默认值只能在运行时发现，也不清楚同一参数对象能否被多个
并发任务复用。

#### 修复方案

1. `Parameters` 改为以下调用所有权契约；各 codec 使用自己的强类型参数 struct 实现它：

   ```go
   type Parameters interface {
       Clone() Parameters
       Validate() error
   }
   ```
2. Codec 对收到的具体参数类型进行受检断言，类型不匹配时包装统一的
   `codec.ErrInvalidParameters`，不读取字符串键。
3. 删除 `BaseParameters`、`NewBaseParameters`、`GetParameter` 和 `SetParameter`；无参数 codec
   使用唯一的无状态 `codec.NoParameters` 值。
4. 每个参数构造函数负责填充默认值；`Validate` 只校验并返回错误，不得修正或补写字段。
   `Clone` 必须深复制 slice、map 和指针字段。
5. Manager 在每次调用前选择用户参数或 codec 默认参数，再执行 `Clone` 和 `Validate`；不得让
   Codec 修改调用方对象，也不得让一次调用修改后影响后续调用。
6. nil 表示使用该 codec 的默认参数；`DefaultParameters()` 每次返回独立值，传给 Codec 的参数
   始终非 nil、已校验且归当前调用独占。
7. string-any 入口退役前，必须把所有已实现参数转成对应 codec 参数 struct 的导出强类型字段，
   不能只迁移常用字段。至少包括 NativeCodec 的 swap bytes，JPEG quality/bit depth、JPEG
   Lossless predictor、JPEG-LS NEAR、HTJ2K quality/block size/levels/progression，以及 JPEG 2000
   的 irreversible、rate、rate levels、levels、layers、progression、target ratio、PCRD、
   lossless layer、verbosity、quantization、subband steps、signed handling、photometric update
   和 MCT 开关。
8. JPEG 2000 当前通过参数 map 承载的 `mctMatrix`、`inverseMctMatrix`、`mctOffsets`、
   `mctNormScale`、`mctAssocType`、`mctMatrixElementType`、`mcoPrecision`、`mcoRecordOrder` 和
   `mctBindings` 必须全部变成强类型字段并保持编码行为；矩阵、slice 和 bindings 必须深复制，
   不允许在移除参数 map 时一并丢失。
9. 任意 custom key 的扩展能力也不能直接丢失：外部 codec 改为通过自己拥有的强类型
   Parameters struct 定义和读取扩展字段，Manager 仍接受任何正确实现 `codec.Parameters` 的
   第三方参数类型。仓库内示例和测试中实际使用的 predictor、quality、bitDepth、near 及上述
   JPEG 2000 参数必须逐项迁移，不得以“custom key”名义忽略。

#### 验收标准

- 两个并发转码不会共享可变默认参数。
- 参数类型不匹配返回明确错误，不 panic。
- 新增 codec 不需要依赖字符串键才能暴露常用参数。
- `BaseParameters` 和 string-any 参数访问入口不再存在。
- 参数 Validate 不改变接收者，Clone 后修改嵌套 slice/map 不影响原对象。
- JPEG、JPEG-LS、JPEG 2000 和 HTJ2K 的参数能力迁移台账完整；MCT/MCO 的矩阵、offset、
  precision、record order 和 binding 回归证明类型化前后编码配置一致。

### DEP-009：codec Registry 存在重复入口和不一致注册契约

**优先级：** P1
**性质：** 已确认的扩展契约缺陷

#### 问题描述

`RegisterCodec(ts, codec)` 同时接收 Transfer Syntax 和 Codec，但不验证
`ts == codec.TransferSyntax()`；nil 输入可能 panic；相同 UID 注册会静默覆盖。与此同时，
`Register*` 包装函数、codec package 的 `init` 和 Registry 方法形成了多套注册入口，
`GetDefaultManager` 又把 Registry 获取和 Manager 创建叠成第二层默认入口。

Codec Registry 与 UID/Transfer Syntax 标准目录的性质不同：它就是运行时扩展点，必须允许
应用或新开发的 codec 在进程运行期间注册实现。需要修复的是注册契约和依赖获取路径，不是把
Registry 改成不可变对象，也不是删除全局插件注册能力。

#### 修复方案

Registry 保留全局插件实例和隔离实例两种明确生命周期；二者使用同一套受检方法，不各自定义
注册 API：

```go
func NewRegistry() *Registry
func GlobalRegistry() *Registry

func (r *Registry) Register(codec Codec) error
func (r *Registry) Replace(codec Codec) (previous Codec, error)
func (r *Registry) Unregister(syntax *transfer.Syntax) (previous Codec, found bool)
func (r *Registry) Lookup(syntax *transfer.Syntax) (Codec, bool)
func (r *Registry) List() []Codec
```

1. Register 从 codec 自身派生 syntax UID，校验 codec、syntax 和 UID 非 nil；分别返回可由
   `errors.Is` 判断的 `ErrNilCodec`、`ErrNilTransferSyntax`，不得 panic 或静默忽略。
2. 普通 Register 遇到重复 UID 返回 `ErrCodecAlreadyRegistered`，原先依靠静默覆盖来更换默认
   参数或实现的调用方必须显式调用 Replace。
3. 只有显式 Replace 允许覆盖，并返回旧实现。
4. Unregister、Lookup 和 List 分别保留现有注销、查询和枚举能力；nil syntax 的 Unregister/
   Lookup 返回未命中。List 返回独立 slice 快照并按 transfer syntax UID 排序，保证并发调用、
   诊断和测试确定性；codec 实例自身的线程安全仍由其实现负责。不再额外暴露同义的
   Get/Has/ListCodecs 转发入口。
5. `codec.GlobalRegistry()` 是唯一的进程级 codec 插件注册表，线程安全、运行时可变，初始包含
   `go-dicom` 自带的 native codec。`codec.NewRegistry()` 只用于需要隔离配置的调用方和测试，
   返回同样包含 native codec 的独立实例；两者是不同生命周期，不是 alias 或兼容入口。
6. 删除旧的 `GetGlobalRegistry`、`GetDefaultManager` 和 package-level `RegisterCodec` 转发。
   Manager 只通过构造函数接收一个 Registry，不再自行选择默认值。
7. `go-dicom-codecs` 单向依赖 `go-dicom` 的 Codec SPI。它保留各 transfer syntax 对应的 codec
   构造函数；各 codec package 的 `init` 使用唯一的
   `codec.GlobalRegistry().Register(codec)` 路径注册默认实现，继续支持 blank import 插件加载；
   init 不得忽略 Register error，重复或无效 codec 必须立即 panic 暴露程序配置错误。删除导出的
   `Register*` 包装函数，不增加 `RegisterAll`。当前这些包装函数携带的 quality、predictor、
   NEAR、rate 或多 transfer syntax 选择能力，分别由已存在的 codec 构造函数和显式
   Register/Replace 组合承接；不得只删除包装函数而丢掉可配置注册能力。需要隔离 Registry 时，
   调用方构造 codec 后仍只调用目标 Registry 的同一套 Register/Replace 方法。
8. `go-dicom` 不得 import `go-dicom-codecs`，也不得枚举或构造外部压缩 codec；它只提供 SPI、
   Registry、native codec 和消费已注册实现的能力。
9. Service、reconstruction 和 image/transcode 的内部对象必须接收明确的 Registry 或 Manager；
   只有 Client/Server 等 composition root 可以选择 `codec.GlobalRegistry()` 并向下传递。
10. 现有 Registry 公开能力逐项映射：`NewCodecRegistry` 到 `NewRegistry`，`GetGlobalRegistry` 到
    `GlobalRegistry`，`RegisterCodec` 到 `Register`/`Replace`，`UnregisterCodec` 到 `Unregister`，
    `GetCodec`/`HasCodec` 到 `Lookup` 的返回值，`ListCodecs` 到 `List` 中 codec 的 transfer syntax
    UID。方法名和返回类型允许收敛，但注册、覆盖、注销、查询、存在性判断和确定性枚举能力均
    不得消失。

#### 验收标准

- syntax 不一致、nil codec、nil syntax 和重复注册均有回归测试。
- Global Registry 在并发注册、查询、替换和注销下保持安全，新开发 codec 可通过唯一 Register
  方法在运行时加入。
- List 在并发修改期间返回按 UID 排序的独立快照；nil 和重复注册错误可分类，blank import 的
  init 注册不吞掉错误。
- 两个独立 Registry 的注册互不影响。
- Service、reconstruction 和 image/transcode 运行过程中不临时查找 Global Registry；默认
  Registry 的选择只发生在 composition root。
- `go-dicom` 的 module/import 图不包含 `go-dicom-codecs`；在 `go-dicom-codecs` 的消费方集成
  测试中，blank import 和显式实例注册两种场景均能完成原有 transfer syntax 转码。该项只在
  `go-dicom` 正式版本发布并被下游引用后执行，不属于发布前验证。

### DEP-010：network/service 直接承担 codec 发现和 Dataset 转码

**优先级：** P1
**性质：** 跨领域职责泄漏

#### 问题描述

C-STORE 发送路径在 presentation context 不匹配时直接获取全局 codec manager、判断转码能力、
创建 Transcoder 并转换 Dataset。网络层因而直接依赖 imaging/codec，且转码策略无法按 Service
配置。

#### 修复方案

1. 在 `dicom/transcode` 提供面向调用方的 Manager，输入 source/output syntax 和 Dataset。
2. Service 配置增加具体的 `*transcode.Manager`，不为测试再定义平行的 Transcoder 接口；
   测试使用带独立 registry 的真实 Manager。
3. Client/Server 构造时显式传递同一依赖；未配置自定义 Manager 时，面向用户的顶层构造函数
   使用 `codec.GlobalRegistry()` 创建并传入默认 manager。因此 blank import
   `go-dicom-codecs` 后已注册的压缩 codec 继续自动参与转码；Service 本身不得临时读取全局值。
   需要隔离配置时，调用方使用 `codec.NewRegistry()`，构造所需的 `go-dicom-codecs` Codec，
   通过该 Registry 唯一的 Register 方法注册后传入。
4. C-STORE 只负责协商候选 syntax、调用转码能力以及发送返回的 Dataset。
5. 无可用 codec、未配置转码和转码失败分别返回可分类错误。

#### 验收标准

- `network/service` 不导入 `imaging/codec`。
- Service 测试可以使用独立 registry，不修改进程全局状态。
- source syntax 直发、无 Pixel Data、可转码和不可转码四条路径保持现有协议语义。
- 未注册外部 codec 时返回明确能力错误；blank import 注册或向隔离 Registry 显式注册后，
  所有现有压缩传输语法仍可发送，不因包移动删除自动转码功能。

### DEP-011：association 为 MessageID 依赖 DIMSE

**优先级：** P2
**性质：** network 内部层级倒置

#### 问题描述

Association 注释称其为“不包含网络 I/O 的纯协商数据结构”，但它持有
`dimse.MessageIDGenerator`，并暴露 `AssignMessageID(dimse.Message)`。真正调用该方法的是
service。UL Association 因一个序号生成器依赖上层 DIMSE 消息模型。

#### 修复方案

1. Service 完整持有一个 `dimse.MessageIDGenerator`，其生命周期与该 Service/Association
   会话一致；Association 不再持有任何 MessageID 状态。
2. Service 在发送请求前调用 generator 分配 ID；`Association.NextMessageID` 的手工生成能力
   迁移到公开的 `dimse.MessageIDGenerator.Next`，`Association.AssignMessageID` 的手工分配能力
   迁移到 `dimse.MessageIDGenerator.AssignMessageID`，全部自动发送入口仍由 Service 代调用。
   完成调用方迁移后退役 Association 上的两个旧入口，不保留转发方法。
3. 将现有所有 request 已通过 `BaseRequest` 提供的 `SetMessageID(uint16) error` 纳入
   `dimse.Request` 接口；`MessageIDGenerator.AssignMessageID` 改为只接收 `dimse.Request`，
   直接调用 setter，删除针对每种具体消息的大型 type switch，不新增平行 setter 接口。

#### 验收标准

- `network/association` 不导入 `network/dimse`。
- 每个 Association/Service 的 ID 序列仍独立，范围保持 1..65535。
- 并发、回绕和保留调用方已有 MessageID 的测试通过。
- MessageID 自动分配功能仍由所有 C-Service/N-Service 发送入口提供，Association 中不存在
  同名兼容入口。
- 原 Association 两个公开方法的手工调用场景均有 MessageIDGenerator 新入口和回归，不删除
  独立生成、重置或向请求分配 MessageID 的能力。

### DEP-012：client 的 Service 测试接口过宽且存在双重语义

**优先级：** P2
**性质：** 可维护性和错误契约风险

#### 问题描述

`client.serviceInterface` 聚合 Association 生命周期和全部 C/N DIMSE 方法，约 20 个方法。
随后 C-FIND、C-MOVE、C-GET 又通过三个可选接口探测是否支持 terminal error channel。

生产 Client 总是创建 `*service.Service`，没有公开第三方 Service 注入入口，因此注释所称的
“第三方 adapter 兼容”没有公共路径。测试替身是否实现可选接口，会改变同一 Client API 的
错误传播行为。

#### 修复方案

1. 统一 Service 的多响应契约为单一 typed event stream：

   ```go
   type ResponseEvent[T dimse.Response] struct {
       Response T
       Err      error
   }

   func (s *Service) SendCFind(ctx context.Context, req *dimse.CFindRequest) (<-chan ResponseEvent[*dimse.CFindResponse], error)
   func (s *Service) SendCMove(ctx context.Context, req *dimse.CMoveRequest) (<-chan ResponseEvent[*dimse.CMoveResponse], error)
   func (s *Service) SendCGet(ctx context.Context, req *dimse.CGetRequest) (<-chan ResponseEvent[*dimse.CGetResponse], error)
   ```

   启动前的校验/写入失败直接返回 error；启动后的协议响应和本地终态错误按发生顺序进入同一
   channel。每个 event 只能设置 Response 或 Err 之一；协议 final response 发出后关闭，终态
   error 发出一次后关闭，不能静默关闭而没有 final/error。
2. 原 response-only 入口的全部响应序列与 `Send*WithError` 的 terminal error 能力统一映射到
   typed event stream；迁移调用方后退役两套旧签名和独立 terminal error channel，不丢弃任一
   正常响应、final response 或本地终态错误。
3. Client 生产字段改回具体 `*service.Service`，所有 Client C-FIND/MOVE/GET 路径消费同一 event
   契约。
4. 单元测试把纯消费循环提取为接收单个 typed event channel 的 helper；生命周期和发送测试
   使用真实 Service + `net.Pipe`，不要为 mock 维护与 Service 等宽的影子接口。
5. 入站 C-FIND 只保留 `CFindHandler func(context.Context, CFindOperation) error`：直接把
   `Handlers.CFindHandler`、`WithCFindHandler` 和 `Server.SetCFindHandler` 切换为 operation 签名，
   移除 `CFindStreamHandler` 后缀入口、旧 slice 签名和 `ErrCFindHandlerConflict`。迁移旧 handler
   时逐个转发原 `CFindResponse` 的 pending identifier 与 final status；若 operation 不能表达旧
   response 的某项行为，先扩充唯一的 CFindOperation 契约，不得丢弃该行为。
6. 移除 `WithDIMSETimeout`。现有调用若依赖其真实行为，迁移到
   `WithHandlerShutdownTimeout`；只有本来想配置出站响应空闲超时的调用才使用
   `WithRequestTimeout`，不得把两种超时重新合成一个含糊 option。
7. 移除只调用 `CEcho` 的 `Client.Ping`，调用方统一使用 DICOM 语义明确的 `Client.CEcho`；
   C-ECHO 的请求、状态检查、context 和错误链保持不变。
8. `ExtendedNegotiation` 只保留方向明确的 `RequestedApplicationInfo` 和
   `AcceptedApplicationInfo`。移除兼容字段 `ServiceClassAppInfo` 及 clone/merge 的双写逻辑；PDU
   RQ/AC 结构中的同名线上字段继续保留，并分别映射到 requested/accepted 值。

#### 验收标准

- 同一 Client 方法的错误语义不取决于底层动态类型。
- 三个 optional error service 接口承载的终态错误能力已进入 typed event stream，旧接口不再作为
  第二套动态分支存在。
- 每类多响应操作只有一个 Service 发送入口和一个有序 event channel，正常 final、异步错误、
  context 取消、association 关闭和 channel 关闭顺序均有回归。
- 测试替身只实现被测路径真正需要的方法。
- 入站 C-FIND 只有一套 handler 字段、option 和 Server setter；原 slice handler 能产生的响应
  序列均可由唯一 CFindOperation 表达。
- Client 不再存在 `Ping` 转发方法，Service 不再存在 `WithDIMSETimeout`，Association 模型不再
  双写 `ServiceClassAppInfo`；对应 C-ECHO、两类 timeout 和扩展协商能力均有唯一入口回归。

### DEP-013：颜色转换和 Modality LUT 重复实现

**优先级：** P2
**性质：** 已确认的重复所有权

#### 问题描述

根 imaging 的 PixelDataConverter 与 render.ColorSpaceConverter 分别实现 YBR/RGB 转换，
其系数精度、舍入方式和 YBR_FULL_422 布局处理并不完全一致。`imaging/lut` 与
`imaging/render` 也各自定义 `ModalityRescaleLUT`。

同一 DICOM 像素经不同入口可能得到不同结果，修复算法时也容易只更新一份。

#### 修复方案

1. `imaging/colorconv` 只提供无状态函数，不再定义空壳 converter struct。整块 buffer 转换统一
   使用 `Convert` 前缀，单像素数值转换不使用该前缀，避免 Go 不支持重载时再次制造 wrapper：

   ```go
   func PlanarToInterleaved(data []byte, samplesPerPixel, bytesPerSample int) ([]byte, error)
   func InterleavedToPlanar(data []byte, samplesPerPixel, bytesPerSample int) ([]byte, error)
   func ConvertMono1ToMono2(data []byte, bitsStored uint16, bytesPerSample int, signed bool) ([]byte, error)
   func ConvertYBRFullToRGB(data []byte) ([]byte, error)
   func ConvertYBRFull422ToRGB(data []byte, width int) ([]byte, error)
   func ConvertYBRPartialToRGB(data []byte) ([]byte, error)
   func ConvertYBRPartial422ToRGB(data []byte, width int) ([]byte, error)
   func ConvertRGBToYBRFull(data []byte) ([]byte, error)
   func ConvertYBRICTToRGB(data []byte, bitsAllocated int) ([]byte, error)
   func ConvertYBRRCTToRGB(data []byte, bitsAllocated int) ([]byte, error)
   func ConvertToRGB(data []byte, width, height int, photometric pixel.PhotometricInterpretation, planar pixel.PlanarConfiguration) ([]byte, error)

   func RGBToYBRFull(r, g, b uint8) (y, cb, cr uint8)
   func YBRFullToRGB(y, cb, cr uint8) (r, g, b uint8)
   func RGBToYBRFull422(r1, g1, b1, r2, g2, b2 uint8) (y1, y2, cb, cr uint8)
   func YBRFull422ToRGB(y1, y2, cb, cr uint8) (r1, g1, b1, r2, g2, b2 uint8)
   func RGBToYBRPartial422(r, g, b uint8) (y, cb, cr uint8)
   func YBRPartial422ToRGB(y, cb, cr uint8) (r, g, b uint8)
   func RGBToYBRICT(r, g, b uint8) (y, cb, cr int16)
   func YBRICTToRGB(y, cb, cr int16) (r, g, b uint8)
   func RGBToYBRRCT(r, g, b uint8) (y, cb, cr int16)
   func YBRRCTToRGB(y, cb, cr int16) (r, g, b uint8)
   ```

   `Color32`、planar/interleaved、MONOCHROME1 到 MONOCHROME2、YBR_FULL、YBR_FULL_422、
   YBR_PARTIAL_422、YBR_ICT、YBR_RCT 和 RGB 反向转换均由这些唯一入口承接。
2. `PixelDataConverter.SwapBytes16/32` 不属于颜色转换，迁移到 `io/endian` 的唯一 copy-and-swap
   slice 能力并保持“不修改输入、返回新 slice”的语义；InterleavedToPlanar24、
   PlanarToInterleaved24、YBRFullToRGB、YBRFull422ToRGB、YBRPartialToRGB 和 RGBToYBRFull 分别
   映射到上述通用或 Convert buffer 函数。逐方法回归通过后退役根 imaging 的
   `PixelDataConverter`、已有的同义 package wrapper 和构造函数。
3. render `ColorSpaceConverter` 的 ConvertToRGB 迁移为 buffer `ConvertToRGB`；其余十个公开
   RGB/YBR 方法逐项迁移为上表对应的标量函数。render/export 直接调用这些函数，只负责校验
   图像尺寸、选择 photometric 路径并映射为 Go image。逐方法回归通过后退役 converter struct
   和构造函数。
4. 以 `imaging/lut.ModalityRescaleLUT` 为唯一 rescale 实现，补齐 render 版本已有的输入范围
   行为后，删除 render 下的同名实现、构造函数、`LUT` type alias 和空的 `ModalityLUT` 包装接口。
5. 将当前只存在于 render 的 `ModalitySequenceLUT` 原样迁入 `imaging/lut`；构造时复制输入
   slice，避免调用方后续修改 LUT 内容。render/pipeline 直接消费 `lut.LUT`。
6. `go-dicom-codecs/jpeg2000/htj2k` adapter 层重复的 `convertFoDicomYBR*` 行为按 3.5 在正式
   版本发布后逐项迁移到 `imaging/colorconv`，回归通过后退役重复实现。JPEG 2000/OpenJPH 内部
   对有符号 component 执行的 ICT/RCT 算法属于 codec 内部数值域，必须保留，不与 DICOM
   8/16-bit photometric adapter 合并。
7. 合并前将两套现有测试迁到唯一所有者并建立共享 golden vectors，覆盖边界值、错误输入、
   舍入、奇数宽度、planar/interleaved、subsampling 和 8/16-bit 情况；同步修正 render README。

#### 验收标准

- 每种颜色转换和 Modality LUT 只有一个生产实现。
- 根 imaging/render 不再暴露同名 converter、LUT 类型或构造函数；旧功能全部迁入唯一新入口。
- 共享 golden vectors 同时覆盖原两套实现支持的全部颜色格式、planar/MONO/byte-swap 能力和
  LUT 行为，避免合并时丢功能。

### DEP-014：测试辅助代码进入生产编译集合

**优先级：** P2
**性质：** 已确认的构建边界缺陷

#### 问题描述

以下文件没有 `_test.go` 后缀，因此会进入生产 package 编译：

- `pkg/network/pdu/test_helpers.go`：直接导入标准库 `testing`。
- `pkg/network/transport/testutil.go`：只为测试生成自签名证书。
- `pkg/imaging/codec/testhelpers.go`：`newTestPixelData` 只被测试调用。

`pkg/dicom/testutil` 的包注释明确说明它为测试和 linter 服务。当前调用审计表明，只有
`SafeUint16FromInt` 被 dataset、parser、writer 的 `_test.go` 使用，用于把有界循环计数转换为
uint16 并满足 gosec G115；它在负数和溢出时静默钳制，既不是 DICOM 领域能力，也不是适合作为
通用公共 API 的 checked conversion。其全部调用值都已由固定循环边界或测试尺寸限定在 uint16
范围内，因此不需要再保留一个共享转换 helper。`SafeUint32FromInt`、`SafeJoin`、`OpenUnder`、
`ReadFileUnder` 均为仓库内零调用。整个包只服务内部测试，不应继续位于公开路径，也不需要迁移成
另一个导出入口。

`codec/transcoder.go` 中还残留一条声称使用 `testPixelData` 的过期注释，实际生产代码使用
的是 `newSimplePixelData`。

#### 修复方案

1. 将 `pkg/network/pdu/test_helpers.go` 重命名为 `reason_strings_test.go`，保留
   `testReasonStrings` 及其现有测试功能。
2. 将 `pkg/network/transport/testutil.go` 重命名为 `test_certificate_test.go`，保留测试证书生成
   功能及全部 TLS 测试调用。
3. 将 `pkg/imaging/codec/testhelpers.go` 重命名为 `test_pixeldata_test.go`，保留
   `newTestPixelData` 及 native codec 测试。
4. 修正 `codec/transcoder.go` 第 490 行附近把 `newSimplePixelData` 错写为 `testPixelData` 的注释。
5. 在 dataset 和 writer 的固定次数建数循环中直接使用 uint16 计数，避免 int 到 uint16 转换；
   parser 测试中保留 int 索引，并在已由图像尺寸、帧数和取模上界证明安全的转换点使用直接转换
   及精确的 `#nosec G115` 边界说明。不得为这些调用新增共享导出 helper，也不得保留静默钳制
   语义来掩盖越界。
6. 移除仅供测试使用的 `SafeUint16FromInt`，同时移除四个经全仓调用审计确认未使用的
   `SafeUint32FromInt`、`SafeJoin`、`OpenUnder`、`ReadFileUnder`，并退役公开的
   `pkg/dicom/testutil`。这是经用途审查和明确确认的测试辅助例外，不能作为删除其他零引用公开
   API 的依据。
7. 使用 `go list -f '{{.GoFiles}}|{{.Imports}}'` 验证测试 helper 不再混入所属生产包，且生产包、
   examples 和外部测试包均不再导入 `pkg/dicom/testutil`。

#### 验收标准

- production `network/pdu` 不再导入 `testing`。
- production `network/transport` 不再包含测试证书生成逻辑。
- production `imaging/codec` 不再编译只供本包测试使用的 testPixelData helper。
- dataset、parser、writer 的原有测试场景继续覆盖相同的数据范围；仓库中不再存在
  `SafeUint16FromInt` 或替代的共享导出 conversion helper，`go list` 不再包含
  `pkg/dicom/testutil`，且没有任何 DICOM 运行时功能因此改变。

### DEP-015：UID、Transfer Syntax 和 Codec 的全局状态语义混乱

**优先级：** P3
**性质：** 跨包一致性债务

#### 问题描述

- UID 标准条目本质上是生成的数据目录，却通过 generated `init` 写入可变 global map；公共
  `Register` 又允许静默覆盖标准条目。
- Transfer Syntax 既有标准常量和 `NewRegistry`，又由 `Builder.Build` 隐式写入可变
  `DefaultRegistry`，同时暴露 package-level 和实例级两套 Register/Lookup/Query 入口。
- Codec 是真正的运行时扩展点，允许 `go-dicom-codecs` 或应用自定义 codec 通过 blank import
  的 `init` 写入 Global Registry；它与只读标准目录不能采用相同的去全局化规则。

三者生命周期并不相同。问题不是 API 名字没有机械统一，而是 UID/Transfer Syntax 的标准数据
和自定义扩展混在可变默认对象中，同时 Codec 合法的运行时注册又存在重复入口和隐式获取。

#### 修复方案

1. UID 生成器直接生成只读的标准条目表和按 UID 字符串索引，不再生成注册调用或 generated
   `init`。当前 `RootUID` 没有被任何生成函数读取，把它改成常量既不能修复 Registry，又会删除
   公开变量的可写行为，因此本轮不改它；自定义根生成仍由 `GenerateFromRoot(root, suffix)`
   显式提供；`RootUID` 保持现状，不属于本轮整改项。
2. `uid.Parse(s, name, type)` 保留当前标准目录命中和未知值构造语义；`MustParse`、`UID.Parse`、
   `IsValid`、`Append`、`Enumerate` 当前承载的标准条目读取以及 UID 的全部查询方法均必须映射。
   新入口明确为：

   ```go
   func StandardEntries() []*UID
   func Parse(s, name string, uidType Type) *UID

   func NewRegistry() *Registry
   func (r *Registry) Register(value *UID) error
   func (r *Registry) Replace(value *UID) (previous *UID, err error)
   func (r *Registry) Unregister(value string) (previous *UID, found bool)
   func (r *Registry) Resolve(value string) (*UID, bool)
   func (r *Registry) Parse(s, name string, uidType Type) *UID
   func (r *Registry) Enumerate() []*UID
   ```

   Registry 是标准目录之上的隔离 overlay：默认可解析全部标准条目；Register 拒绝 nil、非法 UID
   和任何有效重复项；Replace 才能显式 shadow 标准或自定义条目；Unregister 删除当前有效项并在
   该 Registry 内 mask 同 UID 的标准条目，因此保持旧全局 map 覆盖后再注销不会自动恢复标准项
   的可观察语义。Resolve 只返回当前有效条目，Parse 则按“overlay、未被 mask 的标准条目、按传入
   name/type 构造未知值”的顺序处理。Enumerate 返回独立、按 UID 排序的当前有效视图。
   `uid.Register` 的普通新增迁移到 Register，原静默覆盖调用迁移到 Replace，package-level
   Enumerate 的标准读取迁移到 StandardEntries，自定义枚举迁移到 Registry.Enumerate；全部调用
   迁移后才退役 package-level Register 和可变全局 map。
3. Transfer Syntax 的 `Builder.Build` 只构造值，不注册；生成的唯一标准条目表按 UID 建立只读
   索引，不包含 `JPEGProcess1`、`JPEGProcess2_4`、`JPEGProcess14`、`JPEGProcess14SV1` 和
   `JPEG2000Lossy` 五个变量 alias。标准 `*Syntax` 的对象身份、属性及现有直接比较行为必须进入
   快照门禁，不能通过每次查询返回任意 clone 来破坏。
4. `transfer.Parse` 保留标准目录命中与未知 Transfer Syntax 的 Explicit VR、Encapsulated、
   Little Endian 默认属性构造语义，以满足 `Syntax.Parse` 的 parseable 契约。显式 Registry API
   固定为 `NewRegistry`、`Parse`、`Lookup`、`Query`、`Register`、`Replace`、`Unregister` 和
   `List`：Query 只返回当前有效的已知项；Lookup 校验 UID 类型，未命中时按现有默认属性构造；
   Parse 负责字符串规范化后执行同一解析；List 返回按 UID 排序的独立快照。Registry 与 UID
   Registry 一样使用 overlay + mask，使 Replace/Unregister 能保留现有覆盖和移除标准项的有效
   视图语义，而底层标准目录本身不被改写。
5. Association、parser 和需要自定义语法的调用方持有并传递明确的 `*transfer.Registry`；现有
   package-level Register 的普通新增/静默覆盖分别迁移到实例 Register/Replace，Unregister、
   Lookup、Query、KnownEntries 分别迁移到实例的同等方法。全部调用迁移后才退役
   `DefaultRegistry` 及这些 package-level 转发；不得把 Query 的“只查已知项”和 Lookup 的
   “未知项构造默认 Syntax”合并成一个含糊入口。
6. Codec 按 DEP-009 执行，保留线程安全且可变的 `codec.GlobalRegistry()` 作为唯一进程级插件
   注册表，同时保留 `codec.NewRegistry()` 的隔离生命周期；二者只使用 Registry 的
   `Register(codec)`/`Replace(codec)` 方法，不提供 package-level 注册转发或默认 manager。
7. 所有可扩展 Registry 的 Register 返回 error 并拒绝 nil/重复，Replace 才允许覆盖，列表按
   UID 排序。UID/Transfer Syntax 的隔离实例不共享 overlay 或 mask；Codec 的 Global Registry
   则明确用于进程级共享注册，不能与隔离实例混为同一生命周期。
8. UID/Transfer Syntax 的迁移晚于 codec，先建立标准条目数量、值、名称、类型、retired 标志、
   transfer 属性、标准对象身份和五个 transfer alias 的迁移快照，保证去掉 init 和重复名后标准
   数据没有缺项、属性漂移或直接比较回归。

#### 验收标准

- 标准 UID/Transfer Syntax 查询不受 import/init 顺序影响；显式 Registry 的 shadow/mask 只改变
  该实例的有效视图，不修改底层标准目录或其他 Registry。
- 自定义 UID、Transfer Syntax 和 Codec 的注册、替换、查询与枚举能力均由各自 Registry
  完成；Codec Global Registry 的并发访问和隔离 Registry 的并行测试均有覆盖。
- generated UID/Transfer Syntax 条目及全部属性与迁移前快照一致，初始化确定且无副作用。
- Transfer Syntax 的 Query/Lookup 未命中差异、未知属性构造、标准项覆盖后注销、标准对象直接
  比较和 UID 的自定义 name/type 解析均有回归。
- UID/Transfer Syntax 不存在可变默认 registry；Codec 只保留一个明确的可变 Global Registry
  和同一套实例方法，不存在 package-level 注册转发、默认 manager 或第二套注册动作。

## 8. 目标包结构与依赖方向

完成迁移后的核心方向应为：

```text
pkg/io/endian                 无仓库内依赖
pkg/io/buffer              -> io/endian

pkg/dicom/tag                不依赖 dict 或 dictif
pkg/dicom/dict             -> dicom/tag
pkg/dicom/element          -> dicom/tag + io/buffer
pkg/dicom/dataset          -> dicom/element

pkg/imaging/pixel            独立拥有像素格式值对象
pkg/imaging/codec          -> dicom/transfer + imaging/pixel
pkg/imaging/pixeldata      -> imaging/pixel + imaging/codec + dicom/dataset + dicom/element
pkg/imaging/colorconv        独立拥有颜色值与颜色转换
pkg/imaging/lut            -> imaging/colorconv
pkg/imaging/render         -> imaging/pixel + imaging/colorconv + imaging/lut + imaging/pixeldata
pkg/imaging/reconstruction -> imaging/pixeldata + dicom/transcode
pkg/imaging                -> 上述叶子/功能包，只编排 DicomImage 高层能力

pkg/dicom/transcode        -> dicom/dataset + imaging/codec + imaging/pixeldata
pkg/network/service        -> dicom/transcode + network 基础包
pkg/network/client/server  -> network/service
```

这里的 `imaging/pixel` 是低层值对象包，codec 和 pixeldata 都只能消费它；`dicom/transcode`
是高层集成子包，不允许 `dicom/dataset`、`dicom/element`、`imaging/pixel`、`imaging/codec`
或 `imaging/pixeldata` 反向依赖它。因此即使 transcode 同时使用 DICOM 与 imaging 能力，也不会
形成包级循环。

## 9. 推荐实施顺序

### 阶段 1：构建边界清理

处理 DEP-014。该阶段把所属 package 的测试辅助文件改为 `_test.go`，移除公开路径中的
`dicom/testutil` 及其仅供测试的 helper，并修正文档注释；不改变任何 DICOM 运行时行为。

#### 阶段 1 公共能力迁移台账

| 旧 API | 原语义 | 唯一归属/处理 | 调用方与回归 |
| --- | --- | --- | --- |
| `dicom/testutil.SafeUint16FromInt`、`SafeUint32FromInt` | 测试数据的有界整数转换，越界 panic | 无生产替代；各测试在自身已知边界内直接转换 | dataset/parser/writer/network 测试继续覆盖原测试场景 |
| `dicom/testutil.SafeJoin`、`OpenUnder`、`ReadFileUnder` | 仅测试夹具使用的目录越界防护和文件读取 | 无生产替代；全仓与受控外部调用审计确认没有生产调用方 | 原包删除；未引入新的公开 test helper |

`network/pdu` 的 reason string helper、`network/transport` 的证书生成器和 `imaging/codec` 的
内存 PixelData 均不是公开生产能力，已原样迁入所属 package 的 `_test.go`，并由原测试继续使用。

### 阶段 2：基础分层修复

处理 DEP-002。建立 `io/endian`、迁移全部调用和测试并删除 `dicom/endian`，消除当前唯一的
领域级双向依赖。

#### 阶段 2 公共能力迁移台账

| 旧 API | 原语义 | 唯一新 API | 调用方与回归 |
| --- | --- | --- | --- |
| `dicom/endian.Endian`、`Little`、`Big`、`Network` | endian 值对象与 DICOM 网络字节序常量 | `io/endian` 同名定义；`Network` 仍为常量 | dicom/io/network/printing 调用方及 endian 全量测试 |
| `Endian.String`、`IsBig`、`IsLittle`、`ByteOrder` | 查询名称、方向及 `binary.ByteOrder` | `io/endian` 同名方法 | 原 endian 测试迁移后逐项覆盖 |
| `LocalMachine` | 启动时探测的可被外部重写变量 | `io/endian.Native()` | buffer 与 endian 测试覆盖本机字节序，去除可变全局状态 |
| `SwapUint16/32/64`、`SwapInt16/32/64` | 标量字节序交换 | `io/endian` 同名函数 | 原标量表驱动测试 |
| `SwapBytes`、`SwapBytesN` | 原地交换完整 slice 或指定前缀，残尾不变 | `io/endian` 同名函数 | 原 slice/边界测试及 parser/writer round-trip |
| `PixelDataConverter.SwapBytes16/32` | 返回新 slice，不修改调用方输入 | `io/endian.CopyAndSwap` | 独立输入所有权回归及 imaging 调用方 |

### 阶段 3：颜色与 LUT 所有权归一

处理 DEP-013 和 DEP-004 的像素值对象、颜色与 LUT 部分。先建立 `imaging/pixel` 并一次性迁移
`BitDepth`、`Representation`、`PlanarConfiguration`、`Component`、`ColorSpace` 和
`PhotometricInterpretation`，使 `colorconv.ConvertToRGB` 从建立之初就依赖唯一的 pixel 类型；再建立
`imaging/colorconv`，将 LUT 类型和实现统一到 `imaging/lut`。迁移全部调用后删除根
imaging/render 的重复 converter、LUT 实现和 alias。`imagetypes` 在本阶段结束时只剩尚待迁移的
PixelData、FrameInfo 和 FrameInfoSetter，不在旧包新增 alias 或转发入口。

#### 阶段 3 公共能力迁移台账

| 旧 API | 原语义 | 唯一新 API | 调用方与回归 |
| --- | --- | --- | --- |
| `imaging.Component`、字段 `Name/SubSampleX/SubSampleY`、`NewComponent` | 颜色分量值对象 | `imaging/pixel` 同名 API | pixel 原测试迁移 |
| `imaging.ColorSpace`、字段 `Name/Components`、`NewColorSpace`、`String`、`Equals` | 颜色空间值对象；nil 比较与名称等价语义 | `imaging/pixel` 同名 API | pixel 原测试迁移 |
| `imaging.OneBit/Grayscale/Indexed/RGB/BGR/RGBA/YCbCrJPEG` | 标准颜色空间对象 | `imaging/pixel` 同名对象 | 标准对象字段与分量测试 |
| `imaging.PixelRepresentation`、`UnsignedPixels`、`SignedPixels`、`String`、`IsSigned` | signed/unsigned 像素枚举 | `pixel.Representation`、`pixel.UnsignedPixels`、`pixel.SignedPixels` 及同名方法 | PixelData 与 pixel 表驱动测试 |
| `imaging.PlanarConfiguration`、`InterleavedPlanar`、`PlanarPlanar`、`String`、`IsInterleaved`、`IsPlanar` | 像素布局枚举 | `imaging/pixel` 同名 API | planar 转换、sample 和 pixel 测试 |
| `imaging.PhotometricInterpretation`、全部公开字段、`String`、`Equals`、`ParsePhotometricInterpretation`、`MustParsePhotometricInterpretation` | DICOM photometric 解析、比较及 panic/error 语义 | `imaging/pixel` 同名 API | pixel 解析与枚举测试、Dataset 图像测试 |
| `imaging.Monochrome1/Monochrome2/PaletteColor/RGBPhotometric/YbrFull/YbrFull422/YbrPartial422/YbrPartial420/YbrIct/YbrRct` | 标准 photometric 对象及颜色空间关联 | `imaging/pixel` 同名对象 | pixel 标准对象测试及各格式渲染测试 |
| `imagetypes.BitDepth`、全部公开字段、`NewBitDepth`、`BytesAllocated`、`IsValid`、`Mask`、`SignMask`、`ExtendSign`、`MinimumValue`、`MaximumValue`、`Range` | 通用位深计算，包含 8/16/32-bit 与 signed 范围 | `imaging/pixel` 同名 API | BitDepth 原测试完整迁移 |
| `imagetypes.Color32`、全部公开字段、`NewColor32`、`ToInt32` | ARGB 值对象与打包 | `imaging/colorconv` 同名 API | colorconv、palette、overlay、output LUT 测试 |
| `imagetypes.LUT`、`VOILUTFunction` 及三个枚举值 | LUT 契约和 DICOM VOI 函数 | `imaging/lut` 同名实际定义 | lut 与 render pipeline 测试 |
| `imaging.PixelDataConverter`、`NewPixelDataConverter` | 无状态方法容器 | 不保留空壳类型；每项能力由下列唯一函数承接 | 根 imaging 不再持有 converter 字段；原测试迁移到 colorconv/endian |
| `ConvertPlanarToInterleavedGeneric`、`PixelDataConverter.PlanarToInterleaved24` | planar 转 interleaved | `colorconv.PlanarToInterleaved` | 8/16-bit、round-trip、残缺像素错误测试 |
| `PixelDataConverter.InterleavedToPlanar24` | interleaved 转 planar | `colorconv.InterleavedToPlanar` | 8/16-bit、round-trip、残缺像素错误测试 |
| `ConvertMono1ToMono2` | 8/16-bit signed/unsigned MONOCHROME1 反相 | `colorconv.ConvertMono1ToMono2` | PixelData 转换与无效 stored-bit 布局测试 |
| `PixelDataConverter.YBRFullToRGB`、`ConvertYBRFullToRGB` | YBR_FULL buffer 转 RGB | `colorconv.ConvertYBRFullToRGB` | golden vector、round-trip、长度错误测试 |
| `PixelDataConverter.YBRFull422ToRGB`、`ConvertYBRFull422ToRGB` | DICOM Y1/Y2/Cb/Cr 4:2:2 buffer 转 RGB，支持奇数宽度 | `colorconv.ConvertYBRFull422ToRGB` | PixelData 4:2:2 与奇数宽度测试 |
| `PixelDataConverter.YBRPartialToRGB`、`ConvertYBRPartial422ToRGB` | BT.601 partial buffer 转 RGB | `colorconv.ConvertYBRPartialToRGB`、`ConvertYBRPartial422ToRGB` | partial 与 subsampling 回归 |
| `PixelDataConverter.RGBToYBRFull` | RGB buffer 转 YBR_FULL | `colorconv.ConvertRGBToYBRFull` | golden vector 与 round-trip 测试 |
| `ConvertYBRICTToRGB`、`ConvertYBRRCTToRGB` | 8/16-bit JPEG 2000 photometric buffer 转 RGB | `imaging/colorconv` 同名函数 | PixelData 8/16-bit 格式测试 |
| `PixelDataConverter.SwapBytes16/32` | copy-returning byte swap | `io/endian.CopyAndSwap` | 输入不变与 16/32-bit 输出测试 |
| `render.ColorSpaceConverter`、`NewColorSpaceConverter` | 无状态颜色方法容器 | 不保留空壳类型；逐方法迁入 `imaging/colorconv` | render/export 直接调用唯一函数 |
| render 的十个 `RGBToYBR*`/`YBRToRGB*` 标量方法 | 单像素 full/422/partial/ICT/RCT 双向转换 | `colorconv.RGBToYBRFull`、`YBRFullToRGB`、`RGBToYBRFull422`、`YBRFull422ToRGB`、`RGBToYBRPartial422`、`YBRPartial422ToRGB`、`RGBToYBRICT`、`YBRICTToRGB`、`RGBToYBRRCT`、`YBRRCTToRGB` | 原 render 标量测试迁移到 colorconv |
| `ColorSpaceConverter.ConvertToRGB` | 根据 photometric 和 planar 配置输出 RGB | `colorconv.ConvertToRGB`，统一输出 interleaved RGB | RGB/YBR、planar、尺寸与格式测试，render/export 消费 |
| `render.LUT` alias、`render.ModalityLUT` 空包装接口 | render 对 LUT 的重复公开契约 | `imaging/lut.LUT` | render pipeline 和 image metadata 直接消费 |
| `render.ModalityRescaleLUT`、构造函数及全部公开方法 | slope/intercept 与输入范围转换 | `imaging/lut.ModalityRescaleLUT` 同名 API | 原 render/lut 两套测试合并 |
| `render.ModalitySequenceLUT`、构造函数及全部公开方法 | 序列 LUT、clamp、signed descriptor 语义 | `imaging/lut.ModalitySequenceLUT` 同名 API，并复制输入 slice | 变换、min/max、signed 与调用方突变回归 |

Stage 3 的 `go-dicom` 调用方已经切换；`go-dicom-codecs` 中可复用的 photometric adapter 迁移按
3.5 等正式版本发布后执行，当前明确记为待验证，不作为本阶段本仓验收的通过依据。

### 阶段 4：Pixel Data 与 Codec SPI 一次性切换

处理 DEP-004、DEP-005、DEP-006、DEP-007 和 DEP-008。复用阶段 3 已建立的 `imaging/pixel`，
建立 `imaging/pixeldata`，将 FrameInfo/FrameSource/FrameSink 归入 codec，将 context 纳入唯一
Codec 接口并把 string-any 参数的全部现有能力迁移为强类型字段。
本阶段先在 `go-dicom` 完成并验证新 SPI，退役本仓旧入口后发布正式版本；不在发布前运行
`go-dicom-codecs` 集成测试。随后 `go-dicom-codecs` 引用该版本，迁移所有 codec 实现、参数、
示例和测试并完成第二段验收；`go-dicom` 不反向引用外部仓库。

#### 阶段 4 公共能力迁移台账

| 旧 API | 原语义 | 唯一新 API | 调用方与回归 |
| --- | --- | --- | --- |
| `imagetypes.PixelData` | 同时承担 frame 输入、输出与 metadata 查询 | 输入为 `codec.FrameSource`，输出为 `codec.FrameSink` | native/fake codec、PixelData、Transcoder 与 network C-STORE 测试 |
| `imagetypes.FrameInfoSetter`、包函数 `SetFrameInfo` | 可选 metadata 写回，失败仅返回 bool | `codec.FrameSink.SetFrameInfo(codec.FrameInfo) error` | 非法 metadata 返回错误且不修改 sink；Transcoder 不再静默成功 |
| `imagetypes.FrameInfo` 及全部字段 | codec 所需图像格式 metadata | `codec.FrameInfo`；格式字段使用 `imaging/pixel` 值对象，并提供 `Validate` | 8/16-bit、signed、planar、photometric 与 metadata 写回测试 |
| `imaging.PixelDataInfo` 及全部公开字段 | Pixel Data metadata、frame/total size 与校验 | `pixeldata.Info` 同字段及 `BytesAllocated`、`UncompressedFrameSize`、`TotalUncompressedSize`、`Validate` | 原 PixelDataInfo 测试整体迁移 |
| `imaging.DicomPixelData.Info` | 可读写 Pixel Data metadata | `pixeldata.Data.Info` | Dataset 提取、codec metadata 写回和 Clone 回归 |
| `NewDicomPixelData`、`NewDicomPixelDataFromBytes`、`CreatePixelData` | 创建、按 bytes 创建、从 Dataset 提取 | `pixeldata.New`、`pixeldata.NewFromBytes`、`pixeldata.FromDataset` | 构造校验、OB/OW、Big Endian、fragment/BOT 与 palette 测试 |
| `GetFrame`、`AddFrame`、`GetAllFrames`、`FrameCount` | 单帧读写、合并与计数 | `Data.Frame(ctx, index)`、`Data.AddFrame(ctx, frame)`、`Data.AllFrames`、`Data.FrameCount` | 双向 slice 所有权、取消、多帧测试 |
| `GetSample`、`IsPaddingSample`、`CalculateOptimalWindow` | sample 解码、padding 与窗口统计 | `Data.Sample`、`Data.IsPaddingSample`、`Data.CalculateOptimalWindow` | 8/16/32-bit、signed/high-bit、planar 与 padding 测试 |
| `EnsureInterleaved`、`ConvertMonochrome1ToMonochrome2`、`ConvertYBRToRGB` | 像素布局与 photometric 转换 | `pixeldata.Data` 同名方法 | planar、MONOCHROME1、YBR full/422/partial/ICT/RCT 测试 |
| `WindowTo8bit`、`MinMax`、`MaskPadding`、`WindowOrLUTTo8bit` | 灰度窗口、统计、padding mask 与 VOI LUT | `pixeldata.Data` 同名方法 | window、padding、VOI LUT descriptor/data 与 Big Endian 测试 |
| `ToElement`、`IsEncapsulated`、`BasicOffsetTable`、`GetFrameInfo` | Pixel Element 写回、封装状态、BOT 与 codec metadata | `Data.ToElement`、`Data.Encapsulated`、`Data.BasicOffsetTable`、`Data.FrameInfo` | native/encapsulated VR、BOT 重建与 metadata 测试 |
| `DicomPixelData.Encode/Decode` | 通过 codec 转换 Pixel Data | `Data.Encode(ctx, codec, params)`、`Data.Decode(ctx, codec, params)` | context、参数所有权、native encode/decode 与 VR 测试 |
| `Codec.GetDefaultParameters`、无 context `Encode/Decode` | codec 默认参数与帧转换 | `DefaultParameters`、context 首参的 `Encode/Decode` | 已取消 context 不开始读写；多帧与复制边界检查 |
| `BaseParameters`、`NewBaseParameters`、`GetParameter`、`SetParameter` | string-any 可变参数袋 | `Parameters.Clone/Validate`、`PrepareParameters`、`NoParameters` 和 codec 自有强类型参数 | supplied/default 每次深复制、第三方参数与 `ErrInvalidParameters` 测试 |
| Native `swap_bytes` bool key | 默认/禁用/启用 byte swap | `NativeParameters.ByteSwap` 与 `ByteSwapDefault/Disabled/Enabled` | Big Endian 默认交换、显式关闭和强制交换回归 |

本阶段只验证 `go-dicom` 内的 SPI、NativeCodec、fake codec 与全部本仓调用方。JPEG、JPEG-LS、
JPEG 2000、HTJ2K 和 RLE 的强类型参数与实现迁移，必须等待本仓正式版本发布后在
`go-dicom-codecs` 中完成；当前不运行本地替换式集成测试，也不把跨仓迁移标记为已完成。

### 阶段 5：Dataset 转码、Registry 与网络注入

处理 DEP-003、DEP-009 和 DEP-010。建立 `dicom/transcode`，迁移并删除 codec 中的 Dataset
Transcoder；保留并规范化可变 `codec.GlobalRegistry()`，删除默认 Manager 和重复注册 wrapper，
令 network、reconstruction 和 image 内部路径显式接收 Manager。`go-dicom-codecs` 保留各 codec
package 的自动 `init` 插件注册，但统一改为调用 `codec.GlobalRegistry().Register`；其依赖方向
始终是 `go-dicom-codecs -> go-dicom`。验证 blank import、运行时新增 codec、隔离 Registry 和
全部压缩能力后删除旧注册入口。

#### 阶段 5 公共能力迁移台账

| 旧入口或既有能力 | 行为语义 | 唯一新入口 | 调用方与回归证据 |
| --- | --- | --- | --- |
| `codec.Transcoder` | Dataset 转码、File Meta 更新、单帧解码、Pixel Data VR 与 lossy history | `transcode.Transcoder` | `pkg/dicom/transcode` 的 Dataset、metadata、frame、VR 和 writer round-trip 测试 |
| `codec.TranscoderOption` | 配置单个 Dataset 转码器 | `transcode.Option` | `pkg/dicom/transcode/manager_test.go` 的参数和 strict VR 回归 |
| `codec.NewTranscoder` | 按输入/输出 syntax 构造转码器 | `transcode.NewManager(registry)` 后调用 `Manager.NewTranscoder` | 本仓 transcode、network、reconstruction 和示例调用方已迁移；nil registry/syntax 与 codec 缺失回归 |
| `codec.WithInputCodec` | 为输入压缩 syntax 选择自定义实现 | 在目标 `codec.Registry` 上 `Register`/`Replace` 后创建 `transcode.Manager` | `TestManagerDoesNotFallBackToGlobalRegistry` 和隔离 Registry 测试 |
| `codec.WithOutputCodec` | 为输出压缩 syntax 选择自定义实现 | 在目标 `codec.Registry` 上 `Register`/`Replace` 后创建 `transcode.Manager` | Manager 自定义 codec 与转码回归 |
| `codec.WithInputParameters` | 选择输入 codec 参数 | `transcode.WithInputParameters` | Manager 参数 clone/validation 与 Dataset decode 测试 |
| `codec.WithOutputParameters` | 选择输出 codec 参数 | `transcode.WithOutputParameters` | Manager 参数 clone/validation 与 Dataset encode 测试 |
| `codec.WithCodecRegistry` | 选择 codec 生命周期和实现集合 | `transcode.NewManager(registry)` | nil registry 拒绝、隔离 Registry 和不回退 global 的回归 |
| `codec.WithStrictDICOMVR` | 选择 encapsulated Pixel Data 的严格 OB 或兼容 OW | `transcode.WithStrictDICOMVR` | `TestTranscoder_VRSelection` 与 writer round-trip |
| `codec.Transcoder.InputSyntax` | 查询源 Transfer Syntax | `transcode.Transcoder.InputSyntax` | Manager 构造契约测试 |
| `codec.Transcoder.OutputSyntax` | 查询目标 Transfer Syntax | `transcode.Transcoder.OutputSyntax` | Manager 构造契约测试 |
| `codec.Transcoder.Transcode`、`TranscodeContext` | Dataset 转码；context 取消和错误链 | `transcode.Transcoder.Transcode(ctx, ds)` | 转码、context logging、源 Dataset 不可变回归 |
| `codec.Transcoder.TranscodeWithMetadata`、`TranscodeWithMetadataContext` | 同步转换 Dataset 和 File Meta Information | `transcode.Transcoder.TranscodeWithMetadata(ctx, ds, meta)` | metadata 更新、lossy history 和无效 metadata 回归 |
| `codec.Transcoder.DecodeFrame`、`DecodeFrameContext` | 按 frame/BOT 边界解码单帧 | `transcode.Transcoder.DecodeFrame(ctx, ds, frame)` | BOT、fragment、按需读取和 padding 回归 |
| `codec.TranscoderManager` | Registry 绑定的 Dataset 转码器工厂 | `transcode.Manager` | `pkg/dicom/transcode/manager_test.go` |
| `codec.NewTranscoderManager` | 创建 Manager；旧 nil 参数隐式选择 global | `transcode.NewManager(registry)`；nil 明确返回 `ErrNilRegistry` | nil registry 与显式生命周期测试 |
| `codec.TranscoderManager.CreateTranscoder` | 校验 codec 可用性后构造转码器 | `transcode.Manager.NewTranscoder` | compressed codec 缺失和 native-to-native 回归 |
| `codec.TranscoderManager.CanTranscode` | 判断输入/输出 syntax 所需 codec 是否齐备 | `transcode.Manager.CanTranscode` | C-STORE presentation context 与 Manager 单测 |
| `codec.GetDefaultManager` | 从进程全局 Registry 创建默认 Manager | 只由 `client.New`/`server.New` 等 composition root 使用 `codec.GlobalRegistry()` 构造并注入 Manager | client/server 默认构造与自定义注入测试；内部路径不再自行获取 global |
| `codec.NewCodecRegistry` | 创建包含 native codec 的隔离注册表 | `codec.NewRegistry` | 独立 Registry 生命周期和 native codec 回归 |
| `Registry.RegisterCodec(ts, codec)` | 注册或覆盖指定 syntax；旧实现不校验且静默覆盖 | `Registry.Register(codec)`；需要覆盖时显式 `Registry.Replace(codec)` | nil codec/syntax、重复注册、替换返回旧实现和 syntax 自派生测试 |
| `Registry.GetCodec` | 查询 syntax 对应 codec | `Registry.Lookup` | 命中、未命中和 nil syntax 回归 |
| `Registry.HasCodec` | 判断 syntax 是否已注册 | `Registry.Lookup` 的 `found` 返回值 | Manager 能力判断与 Registry 查询测试 |
| `Registry.UnregisterCodec` | 注销 syntax 实现 | `Registry.Unregister`，同时返回旧实现与是否存在 | 注销、nil syntax 和并发回归 |
| `Registry.ListCodecs` | 枚举已注册 syntax UID | `Registry.List` 后读取各 `Codec.TransferSyntax().UID()` | UID 排序、独立 slice 快照和并发修改回归 |
| `codec.GetGlobalRegistry` | 获取线程安全、运行时可变的进程插件注册表 | `codec.GlobalRegistry` | singleton、native 初始值和并发 Register/Lookup/Replace/Unregister/List 回归 |
| Service 的隐式 C-STORE 自动转码 | presentation context 不直匹配时自动选择可转码 syntax，并保留 association | `service.WithTranscodeManager`；client/server 默认或显式向下传递同一 Manager | 直发、无 Pixel Data、可转码、未配置 Manager、codec 缺失、context 和 association 回归 |
| reconstruction 的隐式全局 codec 解码 | encapsulated CT/MR source 按 frame 解码；native source 无需 codec | `reconstruction.WithTranscodeManager`；缺失时返回 `ErrTranscodeManagerRequired` | native 无 Manager 与 encapsulated 显式注入回归 |
| `NewDirectoryIconGenerator()` 的隐式全局查询 | Dataset icon 生成和 encapsulated frame 解码 | `NewDirectoryIconGenerator(registry)` | DICOMDIR 示例显式传入 global；icon native/压缩/缺失 registry 回归 |
| image Dataset 解码的 Registry 获取 | 默认插件解码及调用方隔离配置 | `NewDicomImageFromDataset` 默认选择 `codec.GlobalRegistry()`；`WithImageCodecRegistry` 选择隔离实例；内部只调用 `Lookup` | image compressed codec、隔离 Registry 和 palette 回归 |
| `codec.ConvertEndianness` | 返回独立副本；按 2/4/8-byte sample 交换；非法 sample size/长度报错且不修改输入 | `endian.ConvertEndianness` | `pkg/io/endian` 的成功、错误、输入不变和独立所有权测试 |

Stage 5 发布前只验证本仓 native/fake codec 和全部 `go-dicom` 调用方。`go-dicom-codecs` 的
blank import、各压缩 codec 构造与显式 Register/Replace、强类型参数和 Dataset transcode
集成回归必须等待本仓正式版本发布后执行；当前状态记为“go-dicom 已迁移，外部 codec 集成待
发布后验证”，不得据此宣称跨仓整改完成。

### 阶段 6：字典运行时环清理

处理 DEP-001。迁移仓库内调用后，在同一阶段删除旧 Tag 字典方法、全局 callback 和整个
`dictif` 包，字典功能统一从 `dict.Dictionary` 进入；同时把 ValidationError 路径统一为
`dataset.PathSegment`，把 Tag 数值编码统一为 `Tag.ToUint32()`，移除两个纯兼容入口。

#### 阶段 6 公共能力迁移台账

| 旧入口或既有能力 | 行为语义 | 唯一新入口 | 调用方与回归要求 |
| --- | --- | --- | --- |
| `dictif.Tag` | 字典查询所需的 group、element 和 32 位编码 | `*tag.Tag` | `dict.Lookup` 编译期契约与标准/private/masked tag 查询回归 |
| `dictif.Entry` | name、keyword、VR 和 VM 字典元数据 | `*dict.Entry` | Dictionary lookup 和 serialization keyword/name/VR/VM 回归 |
| `dictif.PrivateCreator` | private creator 值 | `*tag.PrivateCreator` | private dictionary 与 tag parse 回归 |
| `dictif.Lookup` | tag、keyword 和 private creator 查询扩展契约 | `dict.Lookup`，使用具体 `tag`/`dict` 类型 | 自定义 `dict.Dictionary` 注入 JSON/XML 读写回归；`*dict.Dictionary` 编译期实现检查 |
| `dictif.SetGlobalLookup`、`GlobalLookup` | 运行时替换进程级字典服务定位器 | 标准路径使用 `dict.Default()`；自定义路径使用 `WithJSONLookup`、`WithXMLLookup` 显式注入持有的 `dict.Lookup` | 两个不同自定义 Lookup 互不影响，且 JSON/XML 同一配置读写一致；不保留可变全局 callback |
| `Dictionary.LookupTag(dictif.Tag)` | 查询普通、masked 和 private tag 的 Entry | `Dictionary.Lookup(*tag.Tag)` | 现有 Dictionary lookup、masked/private dictionary 测试迁移到唯一入口 |
| `Tag.DictionaryEntry()` | 通过隐式全局字典查询 Entry，返回 `interface{}` | `dict.Lookup.Lookup(tag)`；默认调用方使用 `dict.Default().Lookup(tag)` | 标准、未知、masked 和 private tag 回归，不再依赖 import/init 顺序或类型断言 |
| `tag.ParseKeyword(keyword)` | 通过隐式全局字典把 keyword 解析为 Tag，未命中报错 | `dict.Lookup.LookupKeyword(keyword)`；调用方对 nil 生成自己的上下文错误 | 标准、未知和自定义 keyword JSON/XML 解析回归 |
| `tag.Parse` 的字典 private creator 获取 | 解析私有 Tag 字符串；旧路径从全局字典取缓存后又复制 creator | `tag.NewPrivateCreator` 直接构造值对象 | 私有 Tag group/element/creator、clone、format 和 equality 回归；Dictionary 自身缓存能力仍由 `GetPrivateCreator` 保留 |
| `Tag.Uint32()` | 把 group/element 编码为 32 位值 | `Tag.ToUint32()` | Attribute Tag JSON、Dataset map key 和 tag 数值回归 |
| `dataset.ValidationPathSegment` | ValidationError 的 element/item/fragment 路径节点 | `dataset.PathSegment` | `ValidationError.Path`、validation 内部签名、ClonePath 和 FormatPath 回归 |

Stage 6 只有在仓库内调用方全部切换、JSON/XML 自定义 Lookup 的功能回归通过、`tag` 不再
依赖 `dictif` 且 `go list` 不再包含该包后，才能删除 `dictif`、adapter 和旧入口。

### 阶段 7：network 内部层次收敛

处理 DEP-011 和 DEP-012，统一 MessageID、多响应终态错误和入站 C-FIND handler 契约；同步
移除 `Client.Ping`、`WithDIMSETimeout`、`CFindStreamHandler` 后缀入口和 Association 的含糊
`ServiceClassAppInfo` 字段，并按 6.4 的唯一入口保持原功能。

#### 阶段 7 公共能力迁移台账

| 旧入口或既有能力 | 行为语义 | 唯一新入口 | 调用方与回归要求 |
| --- | --- | --- | --- |
| `Association.NextMessageID()` | 为单个 Association 会话并发安全地生成 1..65535 的非零 MessageID，并在溢出后回绕 | `dimse.MessageIDGenerator.Next()`；实例由 `Service` 持有 | 独立 generator、并发、回绕和非零范围回归；手工生成能力不因移出 Association 而消失 |
| `Association.AssignMessageID(dimse.Message)` | 保留请求已有的非零 ID，否则生成并写回；不支持的消息返回错误 | `dimse.MessageIDGenerator.AssignMessageID(dimse.Request)`，并由 `dimse.Request.SetMessageID` 完成写回 | 所有 C/N request、已有 ID 保留、非 request 编译期排除及 Service 自动分配回归 |
| Service 内针对具体 request 的 MessageID type switch | 为所有支持的 DIMSE request 自动写入 MessageID | `dimse.Request.SetMessageID(uint16) error` | 每个 request 继续通过嵌入的 `BaseRequest` 提供 setter；新增 request 无需修改分配器 |
| `Service.SendCFind` response-only channel 与 `SendCFindWithError` 双入口 | 按顺序返回全部 pending/final C-FIND response，并在启动后报告 timeout、context、关闭等终态错误 | `Service.SendCFind` 返回 `ResponseEvent[*dimse.CFindResponse]` 单一 channel | 启动错误直接返回；pending/final 顺序、terminal error 单次发送、context 取消、association 关闭及禁止静默关闭回归 |
| `Service.SendCMove` response-only channel 与 `SendCMoveWithError` 双入口 | 按顺序返回全部 pending/final C-MOVE response、子操作计数和启动后终态错误 | `Service.SendCMove` 返回 `ResponseEvent[*dimse.CMoveResponse]` 单一 channel | 保留全部 response/count/status；覆盖 final、terminal error、取消、关闭和 channel 顺序 |
| `Service.SendCGet` response-only channel 与 `SendCGetWithError` 双入口 | 按顺序返回全部 pending/final C-GET response、子操作计数和启动后终态错误 | `Service.SendCGet` 返回 `ResponseEvent[*dimse.CGetResponse]` 单一 channel | 保留全部 response/count/status；覆盖 final、terminal error、取消、关闭和 channel 顺序 |
| `client.serviceInterface` 与 `cFindErrorService`、`cMoveErrorService`、`cGetErrorService` 动态分支 | Client 的完整 DIMSE/Association 操作及多响应终态错误消费 | Client 持有具体 `*service.Service`，多响应纯消费逻辑接收对应 typed event channel | Client 的结果、callback、自动 C-CANCEL、final status 和错误链不再随底层动态类型改变 |
| legacy `CFindHandler func(context.Context, *CFindRequest) ([]*CFindResponse, error)` 与 `CFindStreamHandler` | SCP 产生有序 pending identifier、任意 final status，并感知取消 | `CFindHandler func(context.Context, CFindOperation) error` | option、Server setter、示例和文档统一迁移；原 slice response 的 identifier/status 序列可逐项表达 |
| `WithDIMSETimeout` | 实际配置 handler 收尾等待时长，不控制出站 request idle timeout | `WithHandlerShutdownTimeout` | 原调用方迁移后行为不变；`WithRequestTimeout` 独立覆盖出站 response idle timeout |
| `Client.Ping(ctx)` | 原样转发 `Client.CEcho`，执行相同 C-ECHO 请求、状态检查、context 和错误链 | `Client.CEcho(ctx)` | 所有调用方迁移；C-ECHO 成功、失败状态和传输错误回归继续覆盖 |
| `ExtendedNegotiation.ServiceClassAppInfo` | 在同一字段中混用发起方请求值和接受方响应值，并参与 clone/merge | `RequestedApplicationInfo` 与 `AcceptedApplicationInfo` | A-ASSOCIATE-RQ 只映射 requested，AC 只映射 accepted；两个 byte slice 独立复制，PDU 线上字段继续保留 |

Stage 7 只有在三类多响应操作各剩一个 Service 入口、Client 不再依赖等宽测试接口、入站
C-FIND 只有 operation handler、全部旧调用方迁移且上述协议行为回归通过后，才能删除旧入口和
含糊字段；不得以减少 API 数量为由丢失 pending/final response、取消、终态错误或协商数据。

### 阶段 8：UID 与 Transfer Syntax 全局状态清理

处理 DEP-015。Codec Registry 已在阶段 5 按运行时插件语义完成；最后只把 UID/Transfer Syntax
标准目录改成无 init 副作用的不可变数据，把二者的自定义扩展迁到显式实例，并删除其可变默认
对象和重复转发入口；同步移除五个 Transfer Syntax 变量 alias，不套用到 Codec Global
Registry。

#### 阶段 8 公共能力迁移台账

本阶段已经完成下列入口收敛。每一项都是入口调整，原有构造、查询、覆盖、注销和未知值行为
均由新入口承接；没有因为去掉全局对象或 alias 而删除功能。

| 原入口/隐式行为 | 唯一入口 | 功能守恒与回归 |
| --- | --- | --- |
| generated UID `init` 注册和可变全局 map | `uid.StandardEntries`、`uid.Parse`、`uid.NewRegistry` | 标准 UID 数量、值、名称、类型、retired 标志和 canonical object identity 保留；初始化无副作用 |
| `uid.Register`、`uid.Enumerate` | `(*uid.Registry).Register`、`Replace`、`Unregister`、`Enumerate`；标准读取用 `uid.StandardEntries` | 自定义 UID 注册、显式替换、注销 mask、查询和排序快照由隔离 Registry 承接；重复/nil/非法输入返回错误 |
| `transfer.DefaultRegistry`、package-level `Register/Lookup/Query/Unregister/KnownEntries` | `transfer.NewRegistry` 及实例 `Register`、`Replace`、`Unregister`、`Lookup`、`Query`、`List` | 标准 syntax canonical identity、Query/Lookup 未命中差异、未知 syntax 默认属性、overlay/mask/isolation 和排序快照保留 |
| `transfer.Builder.Build` 的隐式注册 | `(*transfer.Builder).Build` 只构造值；需要共享时显式调用 Registry `Register`/`Replace` | 构造无全局副作用，注册能力仍可显式完成 |
| 五个 Transfer Syntax 变量 alias | canonical `JPEGBaseline8Bit`、`JPEGExtended12Bit`、`JPEGLossless`、`JPEGLosslessSV1`、`JPEG2000` | UID、全部 transfer 属性和标准列表成员保留，每个标准 UID 只出现一次 |
| parser/association 的隐式标准查找 | `WithTransferSyntaxRegistry`、association 构造/应用函数的显式 `*transfer.Registry` 参数 | 自定义 syntax 由 composition root 传递；未配置时仍解析标准目录和未知 syntax |

阶段 8 的完成门禁是：标准目录不可变且无 generated `init` 注册；UID/Transfer Syntax 不存在
可变默认 Registry；Codec 仍保留唯一、线程安全、运行时可变的 `codec.GlobalRegistry()`；所有
被迁移的导出能力都有新入口和回归测试。`go-dicom-codecs` 的正式集成测试仍待本仓发布后，
由下游引用已发布版本执行。

## 10. 全局迁移约束

1. 本轮允许公开 API 和 import path 发生破坏性变化，以唯一所有者为准，不以源码兼容为目标。
2. 破坏性变化只允许改变入口和类型组织，不允许减少功能、可观察行为或扩展能力。所有导出
   标识符默认均为有效外部 API，不能用仓库内引用统计推断其无用。
3. 移动概念时必须先完成 3.4 的逐 API 台账，再迁移实现、测试、示例和受控外部调用方；随后
   才能退役旧定义，且不得保留 alias、wrapper、adapter、同名构造函数或 deprecated 第二入口。
4. 无法指出唯一新所有者、新入口、完整功能映射及回归测试的旧 API 不得退役；本整改中不存在
   “没有替代入口但直接删除”的实施项。
5. Codec SPI 调整必须保持 `go-dicom-codecs -> go-dicom` 的单向依赖。`go-dicom` 发布前通过
   fake/native codec 验证契约本身；正式版本发布后，再以下游全部 codec 和消费方集成测试作为
   跨仓完成条件。`go-dicom` 不能 import 外部 codec。
6. 除唯一、线程安全、运行时可变的 `codec.GlobalRegistry()` 插件入口外，不允许引入新的全局
   callback、可变默认 registry 或 `interface{}` 服务定位器。
7. Dataset、frame、metadata 和 parameter 的可变内存所有权必须写入公共注释并由测试验证。
8. 所有 context 错误必须保留标准错误链。
9. 一个阶段不能同时存在供用户选择的新旧入口；迁移提交可以拆分准备和切换，但阶段验收前
   必须完成唯一入口收敛。

## 11. 分阶段验证要求

每个阶段至少执行：

```powershell
go test ./pkg/... -count=1
go build ./...
go vet ./cmd/... ./examples/... ./pkg/... ./tools/...
golangci-lint run
git diff --check
```

涉及以下内容时增加专项验证：

- registry/global state：并行测试和 `go test -race`。
- parser/writer/endian：Part 10 Big/Little Endian 文件回归和 round-trip。
- codec/Pixel Data：多帧、fragment、BOT、8/16-bit、signed/unsigned、metadata write-back。
- public API migration：对照阶段台账逐项运行原行为测试和新入口测试，确认导出能力数量及语义
  没有无替代项；仓库内零引用的导出 API 也必须覆盖。
- external codec：仅在 `go-dicom` 正式版本发布后，在 `go-dicom-codecs` 引用该版本，对 JPEG、
  JPEG-LS、JPEG 2000、HTJ2K、RLE 逐类执行 encode/decode 与 Dataset transcode 集成回归；
  发布前明确记为待验证，不使用本地依赖替代。
- network/transcode：C-STORE presentation context 选择、取消、失败分类和 association 保持。
- color/LUT：共享 golden vectors 覆盖唯一入口的全部迁移前行为。
- reconstruction：volume、MPR、derived DICOM 和 caller-owned Dataset 深拷贝回归。

## 12. 完成判定

本文档中的问题不能仅因代码移动或 import cycle 为零而标记完成。全部修复完成需同时满足：

1. `dictif` 不再参与生产调用，tag 行为与 import/init 顺序无关。
2. `io/...` 不导入 `dicom/...`。
3. imaging 子包不导入根 imaging。
4. `dicom/endian`、`imaging/imagetypes` 不再存在；其全部公共能力已经在唯一新所有者中通过回归，
   且旧包没有 alias、wrapper、adapter 或同名构造入口。
5. codec 包只保留 frame-level SPI、native codec 和 registry。
6. Dataset 转码只有一个所有者，并通过显式依赖提供给 network/reconstruction。
7. Codec Global Registry 保持线程安全和运行时可变；只有 composition root 选择它，内部生产
   路径通过显式传入的 Registry/Manager 使用 codec，不再创建第二个默认入口。
8. 颜色转换和 Modality LUT 各只有一个实现及一个公共入口。
9. 原有字典、Endian、像素格式值对象、像素数据、颜色、LUT、转码、codec 注册、网络自动转码和
   MessageID 自动分配功能均能通过新入口完成，并有迁移前后行为回归。
10. 每个阶段的公共能力迁移台账不存在未映射项；任何公开类型、函数、方法、字段、接口及扩展
    能力都没有因包整理或仓库内缺少调用而消失。
11. 测试辅助代码不进入生产编译集合。
12. `go-dicom-codecs -> go-dicom` 的依赖方向保持不变；`go-dicom` 发布前验证通过，发布后
    `go-dicom-codecs` 引用正式版本完成全量测试、build、vet、lint、必要 race 与协议/像素专项
    回归。发布前该项必须标记为待验证，不能提前判定跨仓整改完成。
