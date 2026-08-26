# go-dicom 与 fo-dicom 能力对齐修复总纲

**文档状态：** 实施、测试和规定的互操作验证已完成；实现仍位于未提交工作区

**维护方式：** 长期总纲；阶段实施计划另行生成

**当前条目：** 9 个 Complete

## 1. 文档目的

本文档定义 `go-dicom` 后续补齐、修复和完善工作的长期工程边界。它不是对
fo-dicom 公共 API 的逐项机械翻译，也不是宣称两个项目可以在所有平台能力上完全等价。
它关注以下三类问题：

1. 会影响 DICOM 协议正确性、连接稳定性或并发安全的缺陷。
2. go-dicom 已有底层消息模型，但缺少完整公共工作流的能力。
3. fo-dicom 已提供且适合 Go 使用场景的通用能力，而 go-dicom 当前仅部分实现。

后续修复应以本文档的稳定条目 ID、目标行为、兼容约束和验收标准为准。进入某个阶段前，
再根据当时的源码生成独立的逐文件实施计划；不得直接把本文档中的文件行号当成永久事实。

## 2. 审计基线

本轮审计日期：2026-08-26。

| 项目 | 分支 | 提交 | 描述 |
| --- | --- | --- | --- |
| `go-dicom` | `master` | `e077350` | `v0.7.0-2-ge077350` |
| `fo-dicom` | `development` | `7ea6d42` | `5.2.6-101-g7ea6d424` |

本轮 Go 验证基线：

```powershell
go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1
go test -race ./pkg/network/... -count=1
go vet ./cmd/... ./examples/... ./pkg/... ./tools/...
golangci-lint run
```

上述命令在审计基线通过，lint 返回 `0 issues`。这只能证明现有测试覆盖内行为正常，
不能证明本文列出的未覆盖语义已经正确。fo-dicom 在本轮仅作为本地源码和公共契约基准，
未执行其完整 .NET 测试套件。

## 3. 总体结论

以下领域已经达到可用且相对完整的水平，本轮不重新立项：

- Dataset、Element、Tag、VR、VM、UID、字符集和标准/私有字典。
- Part 10 文件读写、Deflated Explicit VR Little Endian、Fragment 和大数据处理。
- DICOM JSON/XML、BulkDataURI、匿名化 profile 和 UID 一致映射。
- DICOMDIR、目录扫描、图标、SR、Printing、基础渲染、LUT、Palette、Overlay、MPR。
- C-ECHO、C-STORE、C-FIND、C-MOVE、C-GET 和 N-Service 的底层 DIMSE 消息模型。

仍需处理 9 个条目：

网络编号从 NET-005 延续，是为了避免复用已经在旧审计中发布过的 NET-001 至 NET-004；
已退役文档的完成状态不能自动继承到本轮新问题。

| ID | 优先级 | 完成状态 | 能力 |
| --- | --- | --- | --- |
| NET-005 | P0 | Complete | 将 RequestTimeout 修正为逐 DIMSE 请求超时 |
| NET-006 | P0 | Complete | Presentation Context 数量、ID 和 PDU 防御性校验 |
| NET-009 | P0 | Complete | Client 生命周期并发安全和状态机 |
| NET-007 | P1 | Complete | 完整 N-Service SCU 公共 API |
| NET-008 | P1 | Complete | C-FIND SCP 流式响应、背压和取消 |
| NET-010 | P1 | Complete | 托管请求队列和 Association 生命周期 |
| IMG-004 | P2 | Complete | 渲染管线内空间变换和通用图形合成 |
| CORE-003 | P2 | Complete | SplitFormat Dataset Transform Rule |
| DOC-002 | P0 | Complete | 公共能力声明和外部 codec 文档一致性 |

## 4. 范围与非目标

### 4.1 本文范围

- `pkg/network/client`、`pkg/network/service`、`pkg/network/pdu` 和相关测试。
- `pkg/imaging`、`pkg/imaging/render`、`pkg/imaging/transform` 和相关测试。
- `pkg/dicom/dataset/rules` 和相关测试。
- 根 README、成像 README、网络 API 文档和示例。
- 使用本地 fo-dicom 源码进行有界的协议/API 对照和跨进程互操作验证。

### 4.2 明确非目标

- 不审查或重写 `go-dicom-codecs` 内部 JPEG、JPEG-LS、JPEG 2000、HTJ2K、RLE 算法。
- 不把压缩 codec 搬回 `go-dicom`，也不要求 `go-dicom` 默认依赖全部 codec 包。
- 不移植 WPF、ImageSharp、SkiaSharp、ASP.NET DI 等 .NET 平台专属集成。
- 不为了名称相同而复制 .NET 异步类型；Go API 应使用 context、channel 和明确接口。
- 不在本轮扩展完整 WSI 金字塔/IOD、Enhanced 多帧输出或新 SOP Class 业务工作流。
- 不以 README 声明、自循环测试或类型存在作为互操作完成证据。

### 4.3 codec 归属边界

`go-dicom` 继续负责：

- Transfer Syntax 标识和协商。
- Pixel Data 原生/封装模型。
- codec registry、transcoder 和调用方注册接口。
- C-STORE/C-GET 发送路径的协商语法选择与转码调度。

`go-dicom-codecs` 继续负责压缩算法实现，并由应用通过 blank import 注册所需 codec。
涉及压缩语法的网络验收必须在组合运行时完成，但失败要区分以下四层：

1. UID 是否识别。
2. Presentation Context 是否协商接受。
3. codec 是否已注册。
4. 实际 encode/decode/transcode 是否成功。

## 5. 分类、优先级与验证等级

### 5.1 状态

| 状态 | 含义 |
| --- | --- |
| `Open` | 尚未满足目标行为和验收标准。 |
| `In Progress` | 已有批准的实施计划并正在实现。 |
| `Blocked` | 连续确认存在外部阻塞，当前无法继续。 |
| `Complete` | 源码、测试和所需互操作门槛全部通过。 |
| `External` | 能力有意由配套模块提供。 |
| `Not a gap` | 平台差异或双方均不存在，不作为对齐项。 |

### 5.2 优先级

- `P0`：协议正确性、数据竞争、连接破坏或误导性公共契约。
- `P1`：生产 SCU/SCP 工作流完整性和大规模使用能力。
- `P2`：重要但可独立延后的支撑 API。
- `P3`：专用领域或需要单独产品需求的扩展。

### 5.3 验证等级

| 等级 | 证据 |
| --- | --- |
| L1 | 当前源码、公共 API、静态校验和聚焦单元测试。 |
| L2 | 完整 Go 测试、race、vet、lint、重复运行和故障注入。 |
| L3 | go-dicom 与本地 fo-dicom 的真实 TCP/Part 10 跨进程互操作。 |
| L4 | 组合 `go-dicom-codecs` 后的外部 decoder、像素或 codestream 验证。 |

网络正确性条目至少达到 L2；涉及 wire 行为的条目必须达到 L3。只有实际涉及压缩 Pixel
Data 的场景才要求 L4。

## 6. 全局设计原则

### 6.1 保持兼容

- 已发布的 `Client` 手工 Connect 模式继续保留。
- 新能力优先通过新增 option、方法或包装类型实现，避免重命名现有公共类型。
- 改变现有行为时，必须是修复与文档契约明显不一致的行为，并增加回归测试。
- 不静默吞掉超时、取消、协商失败、无 presentation context 或 codec 未注册错误。
- 错误必须可通过 `errors.Is`/`errors.As` 分类，不依赖字符串解析。

### 6.2 context 和超时分工

- 调用方 context 控制整个调用，包括排队、并发窗口等待、发送和响应等待。
- RequestTimeout 是未由调用方提供更短 deadline 时的默认 DIMSE 响应空闲超时。
- Transport Read/Write Timeout 只约束一次底层 PDU I/O，不代表某个 DIMSE 请求失败。
- AssociationTimeout 只约束 A-ASSOCIATE-RQ/AC/RJ 协商。
- ConnectionTimeout 只约束 TCP/TLS 建连。

### 6.3 并发和所有权

- Client 生命周期状态必须由单一状态机保护。
- 不在持有 Client、pending map 或 handler map 锁时执行网络 I/O、用户回调或 observer。
- caller-owned Dataset、TLS 配置、presentation context 和图像不得被后台流程修改。
- channel 的创建方负责关闭；公共 API 必须写明发送方、关闭方和终态。

### 6.4 失败原子性

- Association 建立失败不得留下半连接的 Client 状态。
- 单个请求超时不得关闭仍可用的 association 或删除其他 pending request。
- Dataset transform 失败不得修改输入 Dataset。
- 图像 transform/graphics 失败不得修改缓存的原始帧或调用方图层。

## 7. 详细修复条目

### NET-005：逐请求 DIMSE 超时

**优先级：** `P0`

**初始状态：** `Open`

**最低验证等级：** `L3`

#### 现状与风险

`client.Config.RequestTimeout` 声明为 DIMSE 请求超时，默认 30 秒，但当前 Client 把它传给
Service 的 PDU read/write timeout。recv loop 在整个连接 30 秒没有 PDU 时返回错误，
Service 随后关闭 association。因此一个空闲连接、一个长时间没有响应的请求，甚至多个正常
pending 请求，都可能被同一个 socket deadline 一起终止。

Service 的 `dimseTimeout` 当前主要用于关闭时等待 request handler，并未实现逐 pending
request 的超时语义。

#### 目标行为

1. `WithRequestTimeout(d)` 配置每个出站 DIMSE 请求的默认响应空闲超时。
2. `d == 0` 表示不增加默认请求超时，只服从调用方 context 和连接关闭。
3. 调用方 context deadline 更早时，以 context 为准。
4. 初始计时基准为请求加入 pending 且开始发送的时间；全部 PDU 发送完成后使用最后发送时间。
5. 对 C-FIND/C-MOVE/C-GET，每收到一个 Pending response 就重置响应空闲计时基准。
6. 超时只删除对应 message ID 的 pending request，释放 async-operation 配额并返回
   `ErrRequestTimeout`；晚到响应被忽略但不得中止 association。
7. 若请求超时时仍有请求 PDU 未完整发送，连接状态不再可信，允许关闭 association；错误必须
   同时保留 request timeout 和 transport failure 上下文。
8. socket idle 不再由 RequestTimeout 控制。无显式 transport read timeout 时，已建立的
   association 可以无限空闲，直至 context、release、abort 或连接错误。

#### 建议公共契约

```go
var ErrRequestTimeout = errors.New("DIMSE request timeout")

type RequestTimeoutError struct {
    MessageID uint16
    Command   dimse.CommandField
    Timeout   time.Duration
}

func WithRequestTimeout(timeout time.Duration) client.Option
func WithTransportReadTimeout(timeout time.Duration) client.Option
func WithTransportWriteTimeout(timeout time.Duration) client.Option
```

`RequestTimeoutError` 应实现 `Unwrap() error { return ErrRequestTimeout }`。Client 新增 transport
options 后，Service 层同时明确以下契约：

```go
func WithRequestTimeout(timeout time.Duration) service.Option
func WithHandlerShutdownTimeout(timeout time.Duration) service.Option
```

现有 `service.WithDIMSETimeout` 标记 Deprecated，并在兼容期内作为
`WithHandlerShutdownTimeout` 的别名；它不能继续宣称控制 request/response timeout。Service 已有
`WithReadTimeout`/`WithWriteTimeout` 保持不变。`defaultServiceConfig` 的 read timeout 改为 `0`，
Client 默认 transport read timeout 同样为 `0`；write timeout 保持有界值，避免永久阻塞写操作。
服务器若需要连接 idle policy，应显式配置 transport read timeout 或在应用层管理 association。

#### 实现边界

- `pendingRequest` 增加开始、最后发送、最后 Pending response 和 timeout cancellation 状态。
- 采用每请求 timer 或单个最小堆/调度器均可；优先选择结构简单、无每秒轮询的实现。
- timeout 完成与正常 response、context cancel、Service close 必须通过单一终态仲裁，保证只完成一次。
- observability 发出一个 timeout 终态，不能随后再发 failed/cancelled。
- Client 把 RequestTimeout 传给 Service 的 request timeout option，而不是 read/write timeout。
- handler shutdown timeout 只用于 Service Close 等待入站 handler，不参与出站 pending request。

#### 必测场景

- 空闲 association 超过 30 秒仍可随后成功 C-ECHO。
- 两个并发请求中一个超时，另一个仍成功且连接保持可用。
- caller deadline 早于默认 timeout。
- 默认 timeout 早于 caller deadline。
- Pending response 连续到达会延长 C-FIND/C-MOVE/C-GET 的空闲超时。
- timeout 后晚到 final response 不关闭连接、不阻塞 recv loop、不产生第二终态。
- timeout 与 final response、Close、C-CANCEL 同时发生时通过 `-race`。
- fo-dicom 测试 SCP 延迟单个响应时，Go 只超时对应请求；随后同 association C-ECHO 成功。

#### 完成标准

- RequestTimeout、transport timeout 和 association timeout 的测试可分别证明互不替代。
- 不再将 Client RequestTimeout 直接接到 Service read/write timeout。
- 聚焦测试、完整网络 race、完整 Go 测试和 fo-dicom 延迟响应互操作通过。

#### 完成记录

- 实现提交：未提交；当前工作区基于 `e07735061c81239d15a3aa953047b4a849d5e2b2`。
- 聚焦测试：`pkg/network/client` 和 `pkg/network/service` 的 request-timeout、late-response、idle-connection 回归测试通过。
- 完整测试：`go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`，退出码 `0`。
- Race/Lint：`go test -race ./pkg/network/... -count=1` 与 `golangci-lint run`，退出码 `0`，`0 issues`。
- 互操作：fo-dicom `5.2.6` TCP SCP 延迟单个请求时仅该请求超时，后续同 association C-ECHO 成功；空闲 `500ms` 超过 `300ms` RequestTimeout 后连接仍复用。
- 有意保留边界：RequestTimeout 不再充当 transport idle policy；应用可显式配置 transport read timeout。

### NET-006：Presentation Context 上限和校验

**优先级：** `P0`

**初始状态：** `Open`

**最低验证等级：** `L2`

#### 现状与风险

Client 使用 `byte(len(contexts)*2+1)` 分配奇数 ID。第 129 个 context 的整数 257 转为 byte
后回绕成 1。现有方法不返回错误，PDU encoder 也没有统一校验重复、零值、偶数 ID 和数量。

#### 目标行为

1. 单个 association 最多允许 128 个 Presentation Context，ID 为 1 到 255 的奇数。
2. 第 129 个 context 在修改 Client 前返回可分类错误，绝不回绕。
3. 手工构造 `AAssociateRQ` 时，PDU encode 同样拒绝零、偶数、重复 ID 和超过 128 项。
4. association decode/accept 路径继续拒绝重复或未知 context ID。
5. 空 Abstract Syntax、空 Transfer Syntax 列表和空 Transfer Syntax UID 应在发送前失败。
6. NET-010 的托管客户端可以拆分 association；当前手工 Client 只负责明确拒绝超限。

#### 建议公共契约

```go
var ErrTooManyPresentationContexts = errors.New("too many presentation contexts")

func (c *Client) AddPresentationContext(
    abstractSyntax string,
    transferSyntaxes ...string,
) error

func (c *Client) AddPresentationContextWithRoles(
    abstractSyntax string,
    scuRole, scpRole bool,
    transferSyntaxes ...string,
) error
```

直接把有返回值的方法作为语句调用仍可编译；但方法值和接口签名可能受影响，因此该变化必须在
release notes 中标记为 API 调整。`AddPresentationContextWithRoles` 只有在 context 添加成功后
才能更新 role selection。

#### 必测场景

- 第 1、128、129 个 context 的 ID 和错误。
- 添加失败后 context/role selection 数量不变。
- PDU encode 拒绝 0、2、重复 1 和 129 项。
- PDU decode 拒绝非法或重复 context。
- 128 个合法 context 与 fo-dicom SCP 完成 association 协商。
- fuzz A-ASSOCIATE-RQ encode/decode 不产生 panic、回绕或重复映射。

#### 完成标准

- 所有入口都不可能产生非法 Presentation Context ID。
- 错误发生在网络写入前，并包含实际数量或 ID。
- 128 项边界在 Go 和 fo-dicom 互操作中通过。

#### 完成记录

- 实现提交：未提交；当前工作区基于 `e07735061c81239d15a3aa953047b4a849d5e2b2`。
- 聚焦测试：Client/PDU 对 1、128、129 contexts、非法 ID、空 syntax 和失败原子性覆盖通过。
- 完整测试：`go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`，退出码 `0`。
- Race/Lint：`go test -race ./pkg/network/... -count=1` 与 `golangci-lint run`，退出码 `0`，`0 issues`。
- 互操作：Go SCU 与 fo-dicom `5.2.6` TCP SCP 协商 128 contexts 成功，`association_count=1`、`context_count=128`。
- 有意保留边界：手工 Client 在第 129 个 context 前返回错误；自动分批属于 ManagedClient。

### NET-009：Client 生命周期并发安全

**优先级：** `P0`

**初始状态：** `Open`

**最低验证等级：** `L2`

#### 并发契约

修复后明确保证以下操作可以并发：

- `Connect`、`Close`、`IsConnected`、`GetAssociation`。
- association 建立后的 C-ECHO/C-STORE/C-FIND/C-MOVE/C-GET/N-Service 调用。
- 多个 DIMSE 请求在协商的 Asynchronous Operations Window 范围内并发。

配置 option 和 Presentation Context 添加只允许在首次 Connect 前进行，不承诺与 Connect
并发；如果发生并发或连接后修改，必须返回状态错误而不是形成数据竞争。

#### 目标状态机

```text
Disconnected -> Connecting -> Connected -> Closing -> Disconnected
       ^             |             |           |
       |             +--failure----+           |
       +----------------close/failure----------+
```

状态和当前连接代次必须由同一 mutex 保护。网络 I/O、TLS handshake、association 协商、
release、用户回调和 observer 不得在持锁时执行。

#### 目标行为

1. 同一 Client 的并发 Connect 只有一个可以进入 Connecting，其余返回稳定 sentinel error。
2. Connecting 期间 Close 会取消正在进行的 dial/TLS/association，并等待或协调清理。
3. Connected 期间 Close 只执行一次 release/close；其他 Close 等待同一结果或安全返回。
4. 旧连接的异步结束事件不能覆盖新一代连接状态。
5. Connect 任一步失败后 conn、service、association 和 connected 状态必须原子回到 Disconnected。
6. DIMSE 方法在非 Connected 状态返回 `ErrClientNotConnected`。
7. IsConnected 只在状态为 Connected 且 Service 未关闭时返回 true。

#### 建议内部结构

```go
type clientState uint8

const (
    clientDisconnected clientState = iota
    clientConnecting
    clientConnected
    clientClosing
)

type clientSession struct {
    generation uint64
    conn        net.Conn
    service     *service.Service
    assoc       *association.Association
    cancel      context.CancelFunc
}
```

不要求公开 `clientState`，但应公开 `ErrClientConnecting`、`ErrClientConnected`、
`ErrClientClosing`、`ErrClientNotConnected` 以支持调用方分类。

#### 必测场景

- 50 个 goroutine 同时 Connect，只有一个实际 dial。
- Connect 与 Close 在 TCP、TLS、association 三个阶段分别竞争。
- 两个 Close 与 peer release/abort 竞争。
- 连接失败后可以复用同一 Client 再次成功 Connect。
- 旧 recv loop 结束不关闭新连接。
- DIMSE 并发请求与 Close 竞争时全部终止且无 goroutine/channel 泄漏。
- 全部测试使用 `go test -race`，关键竞争场景重复至少 100 次。

#### 完成标准

- README 的 concurrent-safe 声明被缩小到上述明确范围。
- 网络 race 套件覆盖同一 Client 的生命周期竞争，而不只是多个独立 Client。
- 使用 goleak 或等价的有界等待断言证明没有长期遗留 goroutine。

#### 完成记录

- 实现提交：未提交；当前工作区基于 `e07735061c81239d15a3aa953047b4a849d5e2b2`。
- 聚焦测试：Connect/Close 竞争、连接失败后重用、旧连接退出和非 Connected DIMSE 错误的回归测试通过。
- 完整测试：`go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`，退出码 `0`。
- Race/Lint：`go test -race ./pkg/network/... -count=1` 与 `golangci-lint run`，退出码 `0`，`0 issues`。
- 互操作：本条最低等级为 L2；网络 race 套件在当前工作区通过。
- 有意保留边界：配置和 Presentation Context 修改不承诺与首次 Connect 并发。

### NET-007：完整 N-Service SCU API

**优先级：** `P1`

**初始状态：** `Open`

**最低验证等级：** `L3`

#### 现状

DIMSE 层已有 N-CREATE、N-GET、N-SET、N-DELETE、N-ACTION、N-EVENT-REPORT 的 request、
response、factory 和 SCP handler。Service 仅提供 N-CREATE/N-SET/N-ACTION/N-DELETE 的
等待响应方法；高层 Client 没有完整 N-Service 方法。

#### 目标 API

Service 补齐：

```go
func (s *Service) SendNGet(
    ctx context.Context,
    req *dimse.NGetRequest,
) (*dimse.NGetResponse, error)

func (s *Service) SendNEventReport(
    ctx context.Context,
    req *dimse.NEventReportRequest,
) (*dimse.NEventReportResponse, error)
```

Client 为全部六种 N-Service 提供一致包装：

```go
func (c *Client) NCreate(ctx context.Context, req *dimse.NCreateRequest) (*dimse.NCreateResponse, error)
func (c *Client) NGet(ctx context.Context, req *dimse.NGetRequest) (*dimse.NGetResponse, error)
func (c *Client) NSet(ctx context.Context, req *dimse.NSetRequest) (*dimse.NSetResponse, error)
func (c *Client) NDelete(ctx context.Context, req *dimse.NDeleteRequest) (*dimse.NDeleteResponse, error)
func (c *Client) NAction(ctx context.Context, req *dimse.NActionRequest) (*dimse.NActionResponse, error)
func (c *Client) NEventReport(ctx context.Context, req *dimse.NEventReportRequest) (*dimse.NEventReportResponse, error)
```

这些方法返回传输/context 错误和原始 DIMSE response。非 Success 状态不转换为 Go error，
保持与现有 Service API 一致；调用方通过 response status 判断业务结果。nil request、未连接、
无可用 Presentation Context 和错误 response 类型必须返回明确错误。

#### 实现要求

- 六个方法统一复用 `sendSimpleRequest`，不能复制 pending/message ID/async slot 逻辑。
- SOP Class、SCU role 和 Presentation Context 校验与 C-Service 使用同一发送入口。
- N-EVENT-REPORT 既可由普通 SCU 发起，也可由已协商 SCP/SCU role 的对端反向发起。
- Printing 的窄接口可继续只声明其需要的四个方法，不强迫已有 fake 实现新增方法。

#### 必测场景

- 六种 request/response 的真实 `net.Pipe` 编解码和 message ID 对应。
- N-GET Attribute Identifier List 与响应 Attribute List 往返。
- N-EVENT-REPORT Event Type ID、Event Information、Event Reply 往返。
- Success、Warning/Failure status、无 Dataset、带 Dataset、context cancel、timeout 和 late response。
- 错误 response 类型不能 panic，也不能遗留 pending entry。
- 与 fo-dicom 双向 N-GET/N-EVENT-REPORT 互操作。

#### 完成标准

- Service 和 Client 两层均覆盖全部六种 N-Service。
- README 只有在 L3 完成后才可继续使用“完整 N-Service SCU”表述。

#### 完成记录

- 实现提交：未提交；当前工作区基于 `e07735061c81239d15a3aa953047b4a849d5e2b2`。
- 聚焦测试：六个 Client/Service N-Service API、message ID、status、Dataset 与错误类型回归测试通过。
- 完整测试：`go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`，退出码 `0`。
- Race/Lint：`go test -race ./pkg/network/... -count=1` 与 `golangci-lint run`，退出码 `0`，`0 issues`。
- 互操作：Go SCU -> fo-dicom `5.2.6` TCP SCP 成功执行 `N-CREATE`、`N-GET`、`N-SET`、`N-ACTION`、`N-EVENT-REPORT`、`N-DELETE`，并在同 association 完成 C-ECHO/C-FIND。
- 有意保留边界：业务 warning/failure status 仍作为 DIMSE response 交给调用方判定，不转换为 Go transport error。

### NET-008：C-FIND SCP 流式响应

**优先级：** `P1`

**初始状态：** `Open`

**最低验证等级：** `L3`

#### 现状

当前 `CFindHandler` 返回完整 `[]*CFindResponse`，handler 返回后 Service 才发送第一条结果。
大结果集会产生首包延迟和额外内存，C-CANCEL 只能通过 context 通知 handler，但无法中止已经
构造的大切片。

#### 推荐 API

参照现有 CMoveOperation/CGetOperation，引入：

```go
type CFindOperation interface {
    Request() *dimse.CFindRequest
    QueryLevel() dimse.QueryRetrieveLevel
    Identifier() *dataset.Dataset
    SendPending(identifier *dataset.Dataset) error
    SendFinal(s *status.Status) error
}

type CFindStreamHandler func(context.Context, CFindOperation) error
```

`Handlers` 新增 `CFindStreamHandler`，保留旧 `CFindHandler` 作为兼容入口并标记 Deprecated。
两者同时配置时优先 stream handler，并在构造/启动时返回配置冲突错误，避免静默忽略。

#### 生命周期和背压

1. `SendPending` 同步等待 response 进入 Service send queue，因此天然提供有界背压。
2. operation 只能发送 Pending，`SendFinal` 只能成功一次；final 后再次发送返回状态错误。
3. handler 正常返回但没有 final 时，Service 自动发送 Success。
4. handler 返回错误且尚未 final 时，Service发送 UnableToProcess；若连接/context 已终止则只返回原错误。
5. 收到匹配 message ID 的 C-CANCEL 后取消 handler context。
6. handler 在 cancel 后未发送 final 时，Service 发送 Cancel final；已经 final 时不重复发送。
7. identifier 在调用发送前进行独立快照或同步编码，调用方后续修改不得改变 wire 数据。

#### 兼容迁移

- 旧 slice handler 通过内部 adapter 逐条调用 operation，现有应用无需立即修改。
- 新示例和文档只展示 stream handler。
- 至少保留一个小版本周期后，再评估是否移除旧字段；本文不授权直接删除。

#### 必测场景

- handler 产生第一条后阻塞，SCU 仍能立即收到第一条，证明不是批量缓冲。
- 慢 SCU/小 send queue 下内存有界、顺序稳定。
- 0、1、10,000 条 Pending 后正确 final。
- handler 显式 Success/Failure/Cancel、隐式 Success、返回 error。
- C-CANCEL 在首条前、中间和 final 竞争窗口到达。
- handler 忽略 context 时，Service Close 仍有界退出并报告 handler shutdown timeout。
- fo-dicom SCU 消费 Go 流式 SCP；Go SCU 消费 fo-dicom `IAsyncEnumerable` SCP。

#### 完成标准

- 第一条响应不依赖全集构造完成。
- 内存使用与 send queue/单条 Dataset 大小相关，而不与结果总数线性增长。
- C-CANCEL、final 和错误路径都只有一个终态。

#### 完成记录

- 实现提交：未提交；当前工作区基于 `e07735061c81239d15a3aa953047b4a849d5e2b2`。
- 聚焦测试：stream/legacy handler 冲突、满 send queue cancel、handler shutdown、single-final 和 cancel 竞争回归测试通过。
- 完整测试：`go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`，退出码 `0`。
- Race/Lint：`go test -race ./pkg/network/... -count=1` 与 `golangci-lint run`，退出码 `0`，`0 issues`。
- 互操作：fo-dicom `5.2.6` SCU -> Go SCP 大流 C-FIND 收到 `2000` Pending（`0000` 至 `1999`）、唯一 Success；cancel 场景收到 `2` Pending、唯一 Cancel final，Go 观测 `c_cancel_requests=1` 且 handler context 已取消。
- 有意保留边界：legacy slice handler 保留兼容并标记 Deprecated；新服务应使用 stream handler。

### NET-010：托管请求队列和 Association 生命周期

**优先级：** `P1`

**初始状态：** `Open`

**最低验证等级：** `L3`

#### 设计选择

保留现有 `Client` 作为“调用方手工配置 Presentation Context 并显式 Connect”的低层模式。
新增组合型 `ManagedClient`，负责请求队列、自动 context 推导、association 分批和 linger。
不在第一个版本中把现有 Client 隐式改造成自动重连客户端。

#### 建议 API

```go
type ManagedClientOptions struct {
    MaximumRequestsPerAssociation int
    AssociationLingerTimeout      time.Duration
    MaximumConsecutiveAssociationTimeouts int
}

type ManagedOption func(*ManagedClientOptions)

type PresentationContextSpec struct {
    AbstractSyntax   string
    TransferSyntaxes []string
    SCURole          bool
    SCPRole          bool
}

type ManagedClient struct { /* private state */ }

func NewManaged(opts ...ManagedOption) *ManagedClient
func WithBaseClientOptions(options ...Option) ManagedOption
func (c *ManagedClient) Add(job Job) error
func (c *ManagedClient) Send(ctx context.Context, host string, port int) error
func (c *ManagedClient) Close() error

type Job interface {
    PresentationContexts() ([]PresentationContextSpec, error)
    Execute(context.Context, *Client) error
    Complete(error)
}
```

公开包应提供 C-ECHO、C-STORE、C-FIND、C-MOVE、C-GET 和六种 N-Service 的 Job 构造器，
避免普通调用方自行实现接口。Job 的 Complete 必须恰好调用一次；实现不能在内部持有调用方可变
Dataset 的非受控引用，也不能在 ManagedClient 内部锁下调用 Complete。

默认值和校验语义固定为：

- `MaximumRequestsPerAssociation == 0` 表示不设请求数上限，但仍受 128 contexts 限制。
- `AssociationLingerTimeout` 默认 50ms；负值非法，0 表示队列暂空时立即 release。
- `MaximumConsecutiveAssociationTimeouts` 默认 3，必须大于 0。
- `PresentationContextSpec` 在加入 batch 时进行深拷贝和 UID/角色校验。
- `WithBaseClientOptions` 复制 option slice；每个 association 用同一不可变配置快照创建低层 Client。
- `Add` 在 Send 运行期间可并发调用；Send 同一时刻只允许一个执行者。

#### context 推导策略

- C-STORE 从 SOP Class 和源 Transfer Syntax 推导，优先源语法，再附加调用方配置的 fallback。
- Query/Retrieve 从 request 的 Affected SOP Class 推导，默认提议 Explicit VR Little Endian 和
  Implicit VR Little Endian。
- N-Service 从对应 Managed SOP Class 推导，默认使用上述非压缩语法。
- 相同 Abstract Syntax 和相同角色要求的请求合并 Transfer Syntax 列表并保持首次出现顺序。
- role selection、extended negotiation 和 user identity 在每个新 association 重新应用。

#### 分批和重试

1. 队列保持 FIFO；一个 association 同时受请求数上限和 128 个 context 上限约束。
2. 下一个 Job 会使 context 超限时，在它之前结束当前 batch，并用新 association 发送。
3. association 建立后继续接收新 Job，直到 linger 超时、请求上限、context 不兼容或 Close。
4. association negotiation timeout 可以在未发送 DIMSE request 时按配置重试。
5. 任意 request 的 PDU 已开始发送后，默认不自动重试，避免 C-STORE/N-ACTION 等重复副作用。
6. 一个 Job 失败默认不阻止同 association 后续 Job；连接级失败则终止当前 batch，未开始 Job
   返回队列错误。是否继续新 association 必须由显式 option 控制，首版默认停止。

#### 必测场景

- 自动 context 去重、transfer syntax 顺序和角色合并。
- 129 个唯一 context 自动拆成两个 association，顺序不变。
- MaximumRequestsPerAssociation 边界和 linger 期间新增请求。
- association timeout 连续上限。
- request 发送前失败可安全重试；发送后断线不重试。
- C-STORE 源语法精确匹配、注册 codec fallback 和未注册 codec 失败。
- 多 producer 并发 Add、Send/Close 竞争和 Complete 恰好一次。
- 与 fo-dicom SCP 连续处理跨 association 的混合 C/N Job。

#### 完成标准

- 普通 SCU 调用方无需手工计算 Presentation Context ID。
- 大批量请求不会因 128 context 上限回绕或丢失。
- 重试策略不制造隐式重复业务操作。
- 低层 Client 行为和现有调用方式保持可用。

#### 完成记录

- 实现提交：未提交；当前工作区基于 `e07735061c81239d15a3aa953047b4a849d5e2b2`。
- 聚焦测试：Job context/clone、FIFO、linger、129-context split、retry-before-request 和 Close exactly-once 回归测试通过。
- 完整测试：`go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`，退出码 `0`。
- Race/Lint：`go test -race ./pkg/network/... -count=1` 与 `golangci-lint run`，退出码 `0`，`0 issues`。
- 互操作：Go ManagedClient -> fo-dicom `5.2.6` TCP SCP 的 129-context batch 成功，`expected_associations=2`、`completed_jobs=129`；SCP association log 为 `128`、`1`。
- 有意保留边界：默认仅重试未开始 DIMSE 的 association setup failure，避免重复业务副作用。

### IMG-004：空间变换和图形合成进入渲染管线

**优先级：** `P2`

**初始状态：** `Open`

**最低验证等级：** `L2`

#### 现状

`SpatialTransform` 已能表达 scale、rotation、flip 和 pan，但 DicomImage 的最终渲染只应用
颜色映射、DICOM Overlay 和 scale。空间矩阵没有真正采样输出图像，也没有类似 fo-dicom
`IImage.DrawGraphics`/`CompositeGraphic` 的通用合成边界。

#### 推荐 API

```go
type FrameRenderOptions struct {
    SpatialTransform *transform.SpatialTransform
    Interpolation    interpolation.Mode
    Background       color.Color
    Viewport         image.Rectangle
    Graphics         []render.Graphic
}

func (img *DicomImage) RenderFrameImageWithOptions(
    frame int,
    options FrameRenderOptions,
) (image.Image, error)

type Graphic interface {
    Bounds() image.Rectangle
    Draw(dst draw.Image) error
}
```

现有 `RenderFrameImage(frame)` 调用新方法的零值 options，保持当前像素输出。现有 `SetScale`
保留兼容：当 `SpatialTransform == nil` 时应用 legacy scale；传入 SpatialTransform 时，它替代
legacy scale，避免双重缩放。`Viewport` 零值表示根据变换后 bounds 自动确定输出；非零时以该
矩形裁剪/填充。后续版本可弃用 SetScale，但本文不要求删除。

#### 渲染顺序

固定为：

1. 解码原始帧。
2. Modality LUT、VOI LUT、Presentation LUT/Invert 和颜色空间转换。
3. DICOM Overlay 合成。
4. SpatialTransform 反向映射采样，包括 rotate/flip/scale/pan。
5. 调用方 Graphics 按切片顺序合成。
6. 导出 PNG/JPEG 或返回 Go image。

这样 DICOM Overlay 与图像一起旋转，而 UI 标注/测量图形保留在最终 viewport 坐标。若调用方需要
病人坐标图形，应先通过 geometry API 转换为 viewport 坐标。

#### transform 实现要求

- `Affine2D` 增加 determinant 和 `Inverse()`；不可逆矩阵返回错误。
- 使用 destination-to-source 反向映射，避免正向映射空洞。
- nearest 和 bilinear 都必须支持 Gray、Gray16、RGBA/NRGBA；其他 image 类型先规范化为 RGBA。
- 自动 viewport 的输出 bounds 由 transform 后四角包围盒确定并规范化到输出原点；pan 不得静默
  裁掉像素。只有调用方提供非零 Viewport 时允许按该矩形裁剪。
- 边界外像素使用 Background，默认透明黑；JPEG 导出可由 exporter 转为不透明背景。
- 90/180/270 度和纯 flip 路径可以优化，但必须与通用 affine 结果一致。

#### graphics 实现要求

- 提供 `ImageGraphic` 和 `CompositeGraphic` 基础实现，不在本条目建设完整矢量绘图库。
- 图层顺序稳定；空图层 no-op；越界图层裁剪而不 panic。
- source image、graphics 和缓存帧均不可被修改。
- Graphic 返回错误时停止合成并返回上下文，不缓存半成品。

#### 必测场景

- 非方形图像的 90/180/270 度旋转、X/Y flip、pan、缩放和组合顺序。
- MONOCHROME1/2、Palette、RGB、planar RGB、Overlay 与 transform 的组合。
- nearest/bilinear 的确定像素 golden test 和 bounds。
- 透明/不透明 graphics、图层顺序、裁剪、失败原子性。
- RenderFrameImage 零值行为与修改前 golden bytes 完全一致。
- 并发渲染不同 frame/options 不产生 race 或共享缓存污染。
- 代表性输出与 fo-dicom 的 transform/composite 结果进行尺寸和像素对比；不要求内部算法同构。

#### 完成标准

- SpatialTransform 不再只是坐标工具，而可通过公共 DicomImage API得到实际变换后的图像。
- 通用 graphics 与 DICOM Overlay 的坐标层级和顺序有明确文档。
- 热路径增加 benchmark，记录 512x512 Gray16 和 RGB 的分配/耗时基线。

#### 完成记录

- 实现提交：未提交；当前工作区基于 `e07735061c81239d15a3aa953047b4a849d5e2b2`。
- 聚焦测试：affine viewport、rotate/flip/pan、graphics 顺序、裁剪和失败原子性回归测试通过。
- 完整测试：`go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`，退出码 `0`。
- Race/Lint：完整网络 race 与 lint 均通过；IMG-004 最低等级为 L2。
- 基准：Gray16 `19.70ms/op`、RGB `18.25ms/op`，均为 `512x512` `ApplyAffine` 本机基线。
- 有意保留边界：不实现平台专属 UI 图形栈；Graphics 保持 Go image/draw 边界。

### CORE-003：SplitFormat Dataset Transform Rule

**优先级：** `P2`

**初始状态：** `Open`

**最低验证等级：** `L2`

#### 目标 API

```go
func SplitFormat(
    t *tag.Tag,
    separators string,
    format string,
) (TransformRule, error)
```

`separators` 是 Unicode rune 集合，按任意一个 rune 分隔完整 canonical value。
`format` 使用与 fo-dicom 主要用法兼容的索引占位符：`{0}`、`{1}` 等；`{{` 和 `}}` 表示字面
花括号。不支持 .NET 数值/日期 culture format，因为 DICOM 值必须由 VR 编码器校验，且 Go
不应引入 locale 隐式行为。

#### 行为定义

1. Tag 不存在时 no-op，不记录 Change。
2. 读取 Tag 的完整 canonical value；多值使用反斜杠连接后再参与 separators 分割。
3. 保留空字段，例如 `A^^B` 分割后为 `A`、空、`B`。
4. format 引用不存在索引时返回构造错误或 transform error，不能输出空字符串掩盖配置错误。
5. 输出作为目标 Tag 的一个 canonical value，经原 VR 和字符集上下文重新编码。
6. 输出与原值相同则 no-op；变化记录 `ChangeEdit`。
7. 所有解析在 rule 构造时完成；apply 期间不重新编译模板。
8. Transform 继续使用现有 clone/transaction 机制，失败不修改输入 Dataset。

#### 必测场景

- 姓名、日期和自定义分隔符重排。
- 多个 separators、连续 separators、首尾空字段和 Unicode 分隔符。
- `{1}, {0}`、重复索引、`{{0}}` 字面量。
- 非法花括号、负数/非数字索引、越界索引、空 separator、nil Tag。
- PN/LO/SH/CS 等文本 VR 的长度、字符集和 padding 校验。
- 嵌套 Sequence 中的 rule path、ChangeSet 和失败原子性。
- 与 fo-dicom `SplitFormatDicomTransformRule` 的代表性输入输出一致。

#### 完成标准

- API、grammar 和错误行为均有 package doc/example。
- 不使用 `fmt.Sprintf` 直接解释不受控 format。
- 聚焦、完整、race 和 lint 通过。

#### 完成记录

- 实现提交：未提交；当前工作区基于 `e07735061c81239d15a3aa953047b4a849d5e2b2`。
- 聚焦测试：SplitFormat grammar、Unicode separator、nested path、charset/VR、literal brace 与失败原子性回归测试通过。
- 完整测试：`go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`，退出码 `0`。
- Race/Lint：完整网络 race 与 lint 均通过；CORE-003 最低等级为 L2。
- 基准：`BenchmarkSplitFormat` 为 `4.989us/op`、`3288 B/op`、`95 allocs/op` 本机基线。
- 有意保留边界：不实现 .NET culture-specific format；输入只接受确定的索引占位符 grammar。

### DOC-002：公共能力声明与文档一致性

**优先级：** `P0`

**初始状态：** `Open`

**最低验证等级：** `L2`

#### 立即修正文档

- 将“full standard compliance”改为可由测试证明的文件读写能力描述。
- 在 NET-007 完成前，将“Complete DIMSE services”拆成消息模型、SCP handler 和 SCU API 三层。
- 删除成像 README 中 JPEG/JPEG-LS/JPEG 2000 “Deferred/未实现”的陈旧章节。
- 明确压缩 codec 来自 `go-dicom-codecs`，并给出 blank import、注册和缺失 codec 错误示例。
- 将 concurrent-safe 限定为 NET-009 定义的具体方法和配置阶段。

#### 能力声明规则

每个 README checked capability 必须至少满足：

1. 存在可编译的公共 API 或示例。
2. 存在聚焦测试，覆盖成功和主要失败路径。
3. 涉及 wire/文件互操作时，记录 L3 证据。
4. 外部能力明确标注模块和注册要求。
5. Partial 能力不能用 Complete、full、all 等绝对词。

#### 文档校验

- README 代码片段应迁移为可编译 Example 或由文档 snippet test 覆盖。
- 增加相对链接校验和陈旧 API 名称扫描。
- 根 README、`pkg/network/README.md`、`pkg/network/docs/API_GUIDE.md`、
  `pkg/imaging/README.md` 对同一能力采用一致措辞。
- release notes 必须列出 AddPresentationContext 返回 error 等公共 API 调整。

#### 完成标准

- 文档中不存在本轮已知的矛盾或超范围声明。
- 新用户能够区分 go-dicom core、go-dicom-codecs、手工 Client 和 ManagedClient。
- 所有公开示例在 CI 中编译。

#### 完成记录

- 实现提交：未提交；当前工作区基于 `e07735061c81239d15a3aa953047b4a849d5e2b2`。
- 文档核对：根 README、网络 README/API guide、成像 README 和 release notes 已与实际 Client、ManagedClient、stream C-FIND、N-Service、外部 codec 注册边界对齐；所有 `AddPresentationContext` 示例处理返回 error。
- 完整测试：`go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1`，退出码 `0`。
- Race/Lint：`go test -race ./pkg/network/... -count=1`、`go vet ./cmd/... ./examples/... ./pkg/... ./tools/...` 和 `golangci-lint run`，退出码 `0`，`0 issues`。
- 互操作：网络声明所需 L3 场景已在 NET-005、NET-006、NET-007、NET-008 和 NET-010 的完成记录中保留实际 fo-dicom `5.2.6` TCP 摘要。
- 有意保留边界：临时 fo-dicom worker 位于忽略的 `.cache`，不作为 Go 默认 CI 依赖。

## 8. 分阶段交付

### Phase 0：协议正确性和公共事实

范围：NET-005、NET-006、NET-009，以及 DOC-002 的立即修正文档部分。

顺序：

1. NET-006：先阻止生成非法 association。
2. NET-005：分离 request timeout 和 transport timeout。
3. NET-009：用明确状态机收紧 Client 生命周期。
4. DOC-002：同步更新这些行为的公共契约。

阶段验收：

- 单请求超时不关闭其他请求或空闲 association。
- Presentation Context 边界无回绕。
- 同一 Client Connect/Close race 全部通过。
- 所有新的 sentinel error 和兼容变化已文档化。

### Phase 1：网络工作流完整性

范围：NET-007、NET-008、NET-010。

顺序：

1. NET-007：补齐简单 request/response N-Service。
2. NET-008：复用 operation 模式实现 C-FIND streaming。
3. NET-010：在稳定的 Client/Service 契约上构建 ManagedClient。

阶段验收：

- 六种 N-Service 在 Service 和 Client 两层可用。
- C-FIND 首条响应流式到达，C-CANCEL 可靠。
- 请求批次可自动推导和拆分 association。
- 与 fo-dicom 的 L3 网络矩阵通过。

### Phase 2：成像和 Dataset 支撑能力

范围：IMG-004、CORE-003。

两个条目相互独立，可以分别实施和发布。IMG-004 需要图像 golden/benchmark；CORE-003 需要
fo-dicom 代表性规则对照，但不需要绑定为同一个 PR。

阶段验收：

- DicomImage 可以实际输出 rotate/flip/pan/scale 后图像并合成 graphics。
- SplitFormat 规则具备确定 grammar、VR 校验和事务语义。
- 完整 Go 测试、race、lint 和相关 benchmark 通过。

### Phase 3：发布契约收口

范围：DOC-002 剩余内容和所有条目的最终状态复核。

阶段验收：

- README/示例与实际 API 一致。
- 每个 Complete 条目记录实现提交、验证命令和互操作证据。
- 不再保留“已实现但 README 仍写未实现”或相反情况。
- 形成下一个版本的 breaking/behavioral change 说明。

## 9. 测试与互操作矩阵

### 9.1 每个代码条目的通用门槛

```powershell
go test ./cmd/... ./examples/... ./pkg/... ./tools/... -count=1
go test -race ./pkg/network/... -count=1
go vet ./cmd/... ./examples/... ./pkg/... ./tools/...
golangci-lint run
git diff --check
```

成像或 Dataset rules 共享状态发生变化时，应增加对应 package 的 race 测试。Windows Go cache
权限异常属于环境问题，应使用仓库本地可写 cache 重跑原命令，不得为此修改生产代码。

### 9.2 网络 L3 矩阵

| 方向 | 场景 | 预期 |
| --- | --- | --- |
| Go SCU -> fo-dicom SCP | C-ECHO 后空闲超过默认 RequestTimeout | association 保持可用 |
| Go SCU -> fo-dicom SCP | 单请求延迟和并发正常请求 | 仅延迟请求超时 |
| Go SCU -> fo-dicom SCP | 128 contexts | association 成功 |
| Go SCU -> fo-dicom SCP | 全部六种 N-Service | message ID、status、Dataset 正确 |
| fo-dicom SCU -> Go SCP | 大结果 C-FIND | 首条及时、顺序正确、内存有界 |
| fo-dicom SCU -> Go SCP | C-FIND 后 C-CANCEL | Go handler context 取消、单一 final |
| Go ManagedClient -> fo-dicom SCP | 超过 128 contexts 的批次 | 自动拆 association、FIFO 不变 |

测试程序必须输出机器可读摘要，至少包含 association 数、message ID、command、status、
pending 数、final 数、超时结果和连接是否复用。临时 fo-dicom 测试程序不得成为 Go 默认 CI 的
构建依赖；稳定的跨库 fixture/协议摘要可以提交到 Go 测试数据中。

### 9.3 成像对照

- 使用合成小图证明每个像素位置，避免只做视觉检查。
- 使用至少一个 MONOCHROME2、一个 MONOCHROME1、一个 Palette、一个 RGB fixture。
- 对无损路径比较完整 RGBA/Gray 像素；bilinear 只允许预先定义的舍入差异。
- 比较 Go 与 fo-dicom 的输出 bounds、图层顺序和代表性像素，不要求 PNG 文件字节相同。

## 10. API 兼容与版本策略

| 变化 | 类型 | 策略 |
| --- | --- | --- |
| AddPresentationContext 返回 error | 源码/API 调整 | release notes，直接语句调用保持可用 |
| Client 默认不因 RequestTimeout 关闭空闲连接 | 行为修复 | 作为 bug fix，明确 transport option |
| 新增 CFindStreamHandler | 向后兼容 | 保留并弃用 slice handler |
| 新增六种 Client N-Service 方法 | 向后兼容 | 不改变 Service 已有四种方法 |
| 新增 ManagedClient | 向后兼容 | 不隐式改变低层 Client |
| 新增 RenderFrameImageWithOptions | 向后兼容 | 旧方法保持零值输出 |
| 新增 SplitFormat | 向后兼容 | 使用独立模板 grammar |

如果实现过程中发现必须删除或重命名公共 API，应停止对应条目，先更新本文档并明确目标主版本；
不得借内部重构顺带扩大 breaking change。

## 11. 风险和回退策略

### 11.1 网络行为变化

RequestTimeout 修复可能暴露过去依赖“30 秒自动断开空闲连接”的应用。应通过显式
TransportReadTimeout 或应用层 idle policy 提供替代，不保留错误语义的兼容开关。

每个网络条目应独立提交。若 L3 失败，可以回退单个条目，不应把 Phase 0/1 合成一个大提交。

### 11.2 handler 兼容

C-FIND streaming 采用新增字段和 adapter。旧 handler 的测试必须持续保留，直至正式弃用周期完成。
若新 operation 终态语义存在争议，应优先保持旧行为并暂停移除，而不是同时支持两套隐式规则。

### 11.3 图像回归

旧 RenderFrameImage 的零值输出是兼容基线。所有新 transform/graphics 必须通过新 options 显式启用。
发现像素差异时先保护旧 golden，再定位 LUT、Overlay、transform 或 exporter 层，不得放宽所有断言。

## 12. 实施计划生成规则

进入每个阶段前创建独立实施计划，计划必须：

1. 重新确认 go-dicom/fo-dicom 当前提交和工作区状态。
2. 将条目拆成可独立 review 的任务，每个任务都有 red-green TDD 步骤。
3. 写出准确文件、方法签名、测试名、运行命令和预期失败/成功输出。
4. 先处理公共契约和失败语义，再实现 happy path。
5. 每个任务结束运行聚焦测试；阶段结束运行完整门槛和 L3 矩阵。
6. 不把临时 fo-dicom worker、下载产物、测试证书或患者数据提交到 Git。
7. 不自动提交或推送；由当前任务的用户指令决定版本控制动作。

建议计划文件：

- `docs/superpowers/plans/2026-08-26-network-correctness.md`
- `docs/superpowers/plans/2026-08-26-network-completeness.md`
- `docs/superpowers/plans/2026-08-26-imaging-and-rules.md`
- `docs/superpowers/plans/2026-08-26-documentation-release-contract.md`

这些计划是执行期材料；本文档是稳定总纲。源码变化导致具体文件或接口方案需要调整时，先更新
对应条目的设计决定和理由，再更新实施计划。

## 13. 条目完成记录模板

条目标记 `Complete` 时，在条目末尾追加以下字段，并填写真实值：

```markdown
#### 完成记录

- 实现提交：完整 Git commit ID
- go-dicom 基线：审计时的完整 Git commit ID
- fo-dicom 基线：对照时的完整 Git commit ID
- 聚焦测试：实际命令、退出码和通过数量
- 完整测试：实际命令和退出码
- Race/Lint：实际命令、退出码和 issue 数量
- 互操作：方向、场景、工具版本和结果摘要
- 有意保留边界：未实现内容及其理由
```

不得填写未经当前执行验证的结果。若只完成 L1/L2，而条目要求 L3，则保持 `In Progress`。

## 14. 总体验收定义

只有同时满足以下条件，才可以重新评估“与 fo-dicom 核心能力基本对齐”的表述：

- 9 个条目全部为 Complete，或经批准明确改为 External/Not a gap。
- P0/P1 网络项完成规定的 L3 互操作。
- 完整 Go 测试、网络 race、vet、lint 全部通过。
- codec 组合场景仍明确区分 go-dicom 与 go-dicom-codecs 所有权。
- README 不再使用无法由测试界定的“full standard compliance”。
- 没有通过放宽断言、忽略 late response、全局关闭连接或自循环测试掩盖问题。

即使满足本文档，也只代表本文审计范围内与 fo-dicom Core 的实用能力对齐，不代表所有 DICOM
IOD、SOP Class、平台集成、第三方 PACS 行为或压缩 codec 算法已经获得全面认证。

## 15. 主要源码参考

go-dicom：

- `pkg/network/client/client.go`
- `pkg/network/client/dimse.go`
- `pkg/network/service/api.go`
- `pkg/network/service/handler.go`
- `pkg/network/service/operation.go`
- `pkg/network/service/options.go`
- `pkg/network/service/recv.go`
- `pkg/network/pdu/associate_rq.go`
- `pkg/imaging/image.go`
- `pkg/imaging/transform/affine.go`
- `pkg/dicom/dataset/rules/transform_text.go`

fo-dicom：

- `FO-DICOM.Core/Network/DicomService.cs`
- `FO-DICOM.Core/Network/DicomServiceOptions.cs`
- `FO-DICOM.Core/Network/DicomPresentationContextCollection.cs`
- `FO-DICOM.Core/Network/IDicomCFindProvider.cs`
- `FO-DICOM.Core/Network/Client/DicomClient.cs`
- `FO-DICOM.Core/Network/Client/DicomClientOptions.cs`
- `FO-DICOM.Core/Imaging/IImage.cs`
- `FO-DICOM.Core/Imaging/Render/CompositeGraphic.cs`
- `FO-DICOM.Core/DicomTransformRules.cs`

引用 fo-dicom 行为时必须重新读取当前本地源码；本文记录的基线用于审计追踪，不是永久版本锁。
