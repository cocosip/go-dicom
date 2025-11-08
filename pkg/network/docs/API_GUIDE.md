# DICOM 网络 API 使用指南

本文档详细介绍 go-dicom 网络功能的 API 使用方法。

## 目录

- [Client API (SCU)](#client-api-scu)
- [Server API (SCP)](#server-api-scp)
- [DIMSE 消息](#dimse-消息)
- [Association 管理](#association-管理)
- [配置选项](#配置选项)
- [错误处理](#错误处理)

---

## Client API (SCU)

### 创建客户端

```go
import "github.com/cocosip/go-dicom/pkg/network/client"

// 基本用法
c := client.New()

// 使用配置选项
c := client.New(
    client.WithCallingAE("MY_SCU"),
    client.WithCalledAE("PACS_SERVER"),
    client.WithMaxPDULength(16384),
    client.WithConnectTimeout(10*time.Second),
    client.WithRequestTimeout(30*time.Second),
    client.WithAssociationTimeout(10*time.Second),
)
```

### 客户端配置选项

| 选项 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `WithCallingAE(ae)` | string | "GO_DICOM_SCU" | 本地 AE Title |
| `WithCalledAE(ae)` | string | "ANY_SCP" | 远程 AE Title |
| `WithMaxPDULength(length)` | uint32 | 16384 | 最大 PDU 长度 |
| `WithConnectTimeout(timeout)` | time.Duration | 10s | TCP 连接超时 |
| `WithRequestTimeout(timeout)` | time.Duration | 30s | DIMSE 请求超时 |
| `WithAssociationTimeout(timeout)` | time.Duration | 10s | Association 协商超时 |
| `WithImplementationClassUID(uid)` | string | "1.2.826.0.1.3680043.10.854" | 实现类 UID |
| `WithImplementationVersionName(name)` | string | "GO-DICOM-1.0" | 实现版本名称 |

### 添加 Presentation Context

Presentation Context 定义了支持的 SOP Class 和 Transfer Syntax：

```go
// Verification SOP Class
c.AddPresentationContext(
    "1.2.840.10008.1.1",   // Abstract Syntax (SOP Class UID)
    "1.2.840.10008.1.2",   // Transfer Syntax: Implicit VR Little Endian
    "1.2.840.10008.1.2.1", // Transfer Syntax: Explicit VR Little Endian
)

// CT Image Storage
c.AddPresentationContext(
    "1.2.840.10008.5.1.4.1.1.2",
    "1.2.840.10008.1.2.1",
)

// Patient Root Query/Retrieve
c.AddPresentationContext(
    "1.2.840.10008.5.1.4.1.2.1.1",
    "1.2.840.10008.1.2.1",
)
```

**常用 SOP Class UIDs**:

| SOP Class | UID |
|-----------|-----|
| Verification | 1.2.840.10008.1.1 |
| CT Image Storage | 1.2.840.10008.5.1.4.1.1.2 |
| MR Image Storage | 1.2.840.10008.5.1.4.1.1.4 |
| Patient Root Q/R - FIND | 1.2.840.10008.5.1.4.1.2.1.1 |
| Study Root Q/R - FIND | 1.2.840.10008.5.1.4.1.2.2.1 |
| Patient Root Q/R - MOVE | 1.2.840.10008.5.1.4.1.2.1.2 |
| Study Root Q/R - MOVE | 1.2.840.10008.5.1.4.1.2.2.2 |

### 连接和关闭

```go
// 连接到服务器
ctx := context.Background()
if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
    return err
}

// 检查连接状态
if c.IsConnected() {
    // 已连接
}

// 优雅关闭（发送 A-RELEASE-RQ）
if err := c.Close(); err != nil {
    return err
}

// 强制中止（发送 A-ABORT）
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
if err := c.Abort(ctx); err != nil {
    return err
}
```

### DIMSE 操作

#### C-ECHO（验证连接）

```go
// 基本用法
if err := c.CEcho(ctx); err != nil {
    return fmt.Errorf("C-ECHO failed: %w", err)
}

// 别名：Ping
if err := c.Ping(ctx); err != nil {
    return fmt.Errorf("ping failed: %w", err)
}
```

#### C-STORE（存储图像）

```go
import (
    "github.com/cocosip/go-dicom/pkg/dicom/parser"
)

// 读取 DICOM 文件
result, err := parser.ParseFile("image.dcm")
if err != nil {
    return err
}

// 发送图像
if err := c.CStore(ctx, result.Dataset); err != nil {
    return fmt.Errorf("C-STORE failed: %w", err)
}

// 带优先级的存储
import "github.com/cocosip/go-dicom/pkg/network/dimse"

if err := c.CStoreWithPriority(ctx, result.Dataset, dimse.PriorityHigh); err != nil {
    return err
}

// 批量存储
datasets := []*dataset.Dataset{ds1, ds2, ds3}
count, err := c.CStoreMultiple(ctx, datasets)
if err != nil {
    log.Printf("Stored %d out of %d datasets: %v\n", count, len(datasets), err)
}
```

#### C-FIND（查询）

```go
import (
    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
)

// 构建查询
query := dataset.New()
query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN*"}))
query.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20240101-"}))
query.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{""}))  // 返回字段
query.Add(element.NewString(tag.StudyDescription, vr.LO, []string{""}))  // 返回字段

// 执行查询（收集所有结果）
results, err := c.CFind(ctx, dimse.QueryRetrieveLevelStudy, query)
if err != nil {
    return err
}

for _, result := range results {
    studyUID := result.GetString(tag.StudyInstanceUID, 0)
    desc := result.GetString(tag.StudyDescription, 0)
    fmt.Printf("Study: %s - %s\n", studyUID, desc)
}

// 流式处理（节省内存）
err = c.CFindWithCallback(ctx, dimse.QueryRetrieveLevelStudy, query,
    func(result *dataset.Dataset) bool {
        studyUID := result.GetString(tag.StudyInstanceUID, 0)
        fmt.Printf("Found: %s\n", studyUID)

        // 返回 true 继续，false 停止
        return true
    })
```

**Query/Retrieve Level**:

```go
dimse.QueryRetrieveLevelPatient  // PATIENT 级别
dimse.QueryRetrieveLevelStudy    // STUDY 级别
dimse.QueryRetrieveLevelSeries   // SERIES 级别
dimse.QueryRetrieveLevelImage    // IMAGE 级别
```

**DICOM 查询通配符**:

| 通配符 | 含义 | 示例 |
|--------|------|------|
| `*` | 匹配任意字符 | "DOE*" 匹配 "DOE", "DOES", "DOE123" |
| `?` | 匹配单个字符 | "DOE?" 匹配 "DOES", 不匹配 "DOE" |
| `-` | 日期范围 | "20240101-20241231" |
| `` | 空字符串 | 返回该字段的值 |

#### C-MOVE（检索 - Pull 模式）

```go
// 构建检索标识
identifier := dataset.New()
identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

// 执行 C-MOVE（将图像发送到第三方）
err := c.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST_AE", identifier,
    func(remaining, completed, failed, warning uint16) bool {
        fmt.Printf("Progress: %d remaining, %d completed, %d failed, %d warning\n",
            remaining, completed, failed, warning)
        return true  // 继续接收进度更新
    })
```

#### C-GET（检索 - Push 模式）

```go
// C-GET 将图像直接发送到本客户端
// 需要先设置 C-STORE handler 来接收图像
// （参见 Server API 的 C-STORE handler 设置）

identifier := dataset.New()
identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

err := c.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier,
    func(remaining, completed, failed, warning uint16) bool {
        fmt.Printf("Retrieve progress: %d/%d\n", completed, completed+remaining)
        return true
    })
```

### 便捷方法

```go
// Patient Root Query
results, err := c.CFindPatientRoot(ctx, dimse.QueryRetrieveLevelStudy, query)

// Study Root Query
results, err := c.CFindStudyRoot(ctx, dimse.QueryRetrieveLevelStudy, query)
```

### 获取 Association 信息

```go
// 获取当前 Association
assoc := c.GetAssociation()
if assoc != nil {
    fmt.Printf("CallingAE: %s\n", assoc.CallingAE)
    fmt.Printf("CalledAE: %s\n", assoc.CalledAE)
    fmt.Printf("MaxPDULength: %d\n", assoc.MaxPDULength)

    // 获取已接受的 Presentation Contexts
    accepted := assoc.GetAcceptedPresentationContexts()
    for _, pc := range accepted {
        fmt.Printf("Context ID %d: %s (%s)\n",
            pc.ID, pc.AbstractSyntax, pc.TransferSyntax)
    }
}

// 获取配置
config := c.GetConfig()
fmt.Printf("CallingAE: %s\n", config.CallingAE)
```

---

## Server API (SCP)

### 创建服务器

```go
import "github.com/cocosip/go-dicom/pkg/network/server"

// 基本用法
srv := server.New()

// 使用配置选项
srv := server.New(
    server.WithAETitle("MY_SCP"),
    server.WithPort(11112),
    server.WithMaxPDULength(16384),
    server.WithMaxConnections(50),
    server.WithAssociationTimeout(30*time.Second),
)
```

### 服务器配置选项

| 选项 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `WithAETitle(ae)` | string | "GO_DICOM_SCP" | 本地 AE Title |
| `WithPort(port)` | int | 104 | 监听端口 |
| `WithMaxPDULength(length)` | uint32 | 16384 | 最大 PDU 长度 |
| `WithMaxConnections(max)` | int | 0 | 最大并发连接（0=无限制） |
| `WithAcceptTimeout(timeout)` | time.Duration | 0 | Accept 超时（0=无限制） |
| `WithAssociationTimeout(timeout)` | time.Duration | 10s | Association 超时 |
| `WithRequestTimeout(timeout)` | time.Duration | 30s | 请求处理超时 |
| `WithAcceptedCallingAETitles(aes)` | []string | nil | 允许的 Calling AE（nil=全部允许） |
| `WithStrictAECheck(strict)` | bool | false | 严格 AE Title 检查 |
| `WithTLSConfig(config)` | *tls.Config | nil | TLS 配置 |

### 设置 DIMSE Handlers

#### C-ECHO Handler

```go
srv.SetCEchoHandler(func(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
    log.Println("Received C-ECHO")
    // 总是返回成功
    return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
})
```

#### C-STORE Handler

```go
srv.SetCStoreHandler(func(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
    ds := req.DataDataset()

    // 获取图像信息
    sopClassUID := req.AffectedSOPClassUID()
    sopInstanceUID := req.AffectedSOPInstanceUID()

    log.Printf("Storing image: %s\n", sopInstanceUID)

    // 保存到文件或数据库
    filename := fmt.Sprintf("./received/%s.dcm", sopInstanceUID)
    if err := writer.WriteFile(filename, ds); err != nil {
        log.Printf("Failed to save: %v\n", err)
        // 返回失败状态
        return dimse.NewCStoreResponseFromRequest(req, 0xC000), nil
    }

    // 返回成功状态
    return dimse.NewCStoreResponseFromRequest(req, 0x0000), nil
})
```

#### C-FIND Handler

```go
srv.SetCFindHandler(func(ctx context.Context, req *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
    query := req.DataDataset()
    level := req.QueryRetrieveLevel()

    log.Printf("C-FIND query at %s level\n", level)

    // 在数据库中查询
    results := searchInDatabase(query, level)

    // 构建响应
    responses := make([]*dimse.CFindResponse, 0, len(results)+1)

    // Pending responses（每个结果一个）
    for _, result := range results {
        resp := dimse.NewCFindResponsePending(req, result)
        responses = append(responses, resp)
    }

    // Final success response
    finalResp := dimse.NewCFindResponseSuccess(req)
    responses = append(responses, finalResp)

    return responses, nil
})
```

#### C-MOVE Handler

```go
srv.SetCMoveHandler(func(ctx context.Context, req *dimse.CMoveRequest) ([]*dimse.CMoveResponse, error) {
    identifier := req.DataDataset()
    destination := req.MoveDestination()

    log.Printf("C-MOVE to %s\n", destination)

    // 1. 查询匹配的图像
    images := findImages(identifier)

    // 2. 创建到目标的连接
    destClient := client.New(
        client.WithCallingAE("MY_SCP"),
        client.WithCalledAE(destination),
    )
    // ... 添加 presentation contexts
    destClient.Connect(ctx, getDestinationHost(destination), 104)
    defer destClient.Close()

    // 3. 发送图像
    responses := make([]*dimse.CMoveResponse, 0)
    remaining := uint16(len(images))
    completed := uint16(0)

    for _, image := range images {
        // 发送 Pending response
        remaining--
        resp := dimse.NewCMoveResponsePending(req, remaining, completed, 0, 0)
        responses = append(responses, resp)

        // 发送图像到目标
        if err := destClient.CStore(ctx, image); err != nil {
            log.Printf("Failed to move image: %v\n", err)
            continue
        }
        completed++
    }

    // Final success response
    finalResp := dimse.NewCMoveResponseSuccess(req)
    responses = append(responses, finalResp)

    return responses, nil
})
```

#### C-GET Handler

```go
srv.SetCGetHandler(func(ctx context.Context, req *dimse.CGetRequest) ([]*dimse.CGetResponse, error) {
    // C-GET 类似 C-MOVE，但直接发送到请求方
    // 实现逻辑类似 C-MOVE Handler
    // ...
    return responses, nil
})
```

### 设置生命周期 Handlers

#### Association 协商

```go
import (
    "github.com/cocosip/go-dicom/pkg/network/association"
    "github.com/cocosip/go-dicom/pkg/network/service"
)

srv.SetAssociationNegotiatorFunc(
    func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
        // 1. 验证 Calling AE Title
        if assoc.CallingAE != "AUTHORIZED_SCU" {
            return responder.RejectAssociation(ctx,
                pdu.ResultRejectedPermanent,
                pdu.SourceServiceUser,
                pdu.ReasonCallingAETitleNotRecognized)
        }

        // 2. 选择性接受 Presentation Contexts
        for _, pc := range assoc.PresentationContexts {
            if isSupported(pc.AbstractSyntax) {
                // 接受（选择第一个支持的 Transfer Syntax）
                pc.Accept(pdu.ResultAcceptance, pc.TransferSyntaxes[0])
            } else {
                // 拒绝
                pc.Reject(pdu.ResultAbstractSyntaxNotSupported)
            }
        }

        // 3. 接受 Association
        return responder.AcceptAssociation(ctx, assoc)
    })
```

#### Association 释放

```go
srv.SetAssociationReleaseHandlerFunc(func(ctx context.Context) error {
    log.Println("Client requesting release")

    // 可以在这里进行清理工作
    // 返回 nil 接受释放，返回 error 拒绝释放
    return nil
})
```

#### 连接生命周期

```go
srv.SetConnectionLifecycleHandlerFuncs(
    // OnAbort
    func(ctx context.Context, source byte, reason byte) {
        log.Printf("Connection aborted: source=%d, reason=%d\n", source, reason)
        // 进行清理工作
    },
    // OnConnectionClosed
    func(ctx context.Context, err error) {
        if err != nil {
            log.Printf("Connection closed with error: %v\n", err)
        } else {
            log.Println("Connection closed normally")
        }
        // 进行清理工作
    },
)
```

### 启动和关闭服务器

```go
// 启动服务器（阻塞）
ctx := context.Background()
if err := srv.ListenAndServe(ctx); err != nil {
    log.Fatal(err)
}

// 在另一个 goroutine 中优雅关闭
go func() {
    <-shutdownChan
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Printf("Shutdown error: %v\n", err)
    }
}()
```

### 监控服务器状态

```go
// 检查服务器是否运行
if srv.IsRunning() {
    // 服务器正在运行
}

// 获取活跃连接数
activeConns := srv.ActiveConnections()
log.Printf("Active connections: %d\n", activeConns)
```

---

## DIMSE 消息

### 消息结构

所有 DIMSE 消息都实现了 `dimse.Message` 接口：

```go
type Message interface {
    CommandField() uint16
    MessageID() uint16
    SetMessageID(id uint16)
    CommandDataset() *dataset.Dataset
    DataDataset() *dataset.Dataset
    HasData() bool
}
```

### 创建请求消息

```go
import "github.com/cocosip/go-dicom/pkg/network/dimse"

// C-ECHO
req := dimse.NewCEchoRequest()

// C-STORE
req, err := dimse.NewCStoreRequest(dataset)

// C-FIND
req := dimse.NewCFindRequest(level, query)
req := dimse.NewCFindRequestPatientRoot(level, query)
req := dimse.NewCFindRequestStudyRoot(level, query)

// C-MOVE
req := dimse.NewCMoveRequest(level, destination, identifier)
req := dimse.NewCMoveRequestPatientRoot(destination, identifier)
req := dimse.NewCMoveRequestStudyRoot(destination, identifier)

// C-GET
req := dimse.NewCGetRequest(level, identifier)
req := dimse.NewCGetRequestPatientRoot(identifier)
req := dimse.NewCGetRequestStudyRoot(identifier)
```

### 创建响应消息

```go
// 从请求创建响应
resp := dimse.NewCEchoResponseFromRequest(req, statusCode)
resp := dimse.NewCStoreResponseFromRequest(req, statusCode)

// Pending 响应
resp := dimse.NewCFindResponsePending(req, identifier)
resp := dimse.NewCMoveResponsePending(req, remaining, completed, failed, warning)
resp := dimse.NewCGetResponsePending(req, remaining, completed, failed, warning)

// Success 响应
resp := dimse.NewCFindResponseSuccess(req)
resp := dimse.NewCMoveResponseSuccess(req)
resp := dimse.NewCGetResponseSuccess(req)
```

### 检查响应状态

```go
status := resp.Status()

if status.IsSuccess() {
    // 成功 (0x0000)
} else if status.IsPending() {
    // 等待中 (0xFF00)
} else if status.IsWarning() {
    // 警告 (0x0001, 0x0107, 0x0116, 0xB000-0xBFFF)
} else if status.IsFailure() {
    // 失败 (0xA700, 0xA900, 0xC000-0xCFFF)
}

// 获取状态码
statusCode := resp.StatusCode()  // uint16
```

---

## Association 管理

### Association 结构

```go
type Association struct {
    CalledAE             string
    CallingAE            string
    ApplicationContext   string
    PresentationContexts []*PresentationContext
    MaxPDULength         uint32
    ImplementationClassUID    string
    ImplementationVersionName string
    // ...
}
```

### Presentation Context

```go
type PresentationContext struct {
    ID             byte
    AbstractSyntax string
    TransferSyntax string
    Result         byte
}

// 查找 Presentation Context
pc := assoc.FindPresentationContextByID(1)
pc := assoc.FindPresentationContextByAbstractSyntax("1.2.840.10008.1.1")

// 获取已接受的 Contexts
accepted := assoc.GetAcceptedPresentationContexts()

// 获取 Transfer Syntax
ts := assoc.GetTransferSyntaxForAbstractSyntax("1.2.840.10008.1.1")
```

---

## 配置选项

### Context 超时控制

```go
// 为单个操作设置超时
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

err := client.CEcho(ctx)

// Context 取消
ctx, cancel := context.WithCancel(context.Background())
go func() {
    time.Sleep(5 * time.Second)
    cancel()  // 取消操作
}()
```

### Priority 设置

```go
import "github.com/cocosip/go-dicom/pkg/network/dimse"

dimse.PriorityLow     // 0x0002
dimse.PriorityMedium  // 0x0000 (默认)
dimse.PriorityHigh    // 0x0001

// 在 C-STORE 中使用
err := client.CStoreWithPriority(ctx, dataset, dimse.PriorityHigh)
```

---

## 错误处理

### 常见错误类型

```go
// 连接错误
if err := c.Connect(ctx, host, port); err != nil {
    if strings.Contains(err.Error(), "connection refused") {
        // 服务器未运行或端口错误
    } else if strings.Contains(err.Error(), "timeout") {
        // 连接超时
    } else if strings.Contains(err.Error(), "no presentation contexts") {
        // 没有添加 Presentation Context
    }
}

// Association 拒绝
if err := c.Connect(ctx, host, port); err != nil {
    if strings.Contains(err.Error(), "association rejected") {
        // Association 被拒绝
        // 可能原因：AE Title 不匹配、不支持的 SOP Class
    }
}

// DIMSE 错误
if err := c.CEcho(ctx); err != nil {
    if strings.Contains(err.Error(), "not connected") {
        // 未连接
    } else if strings.Contains(err.Error(), "context deadline exceeded") {
        // 超时
    }
}
```

### 状态码错误

```go
resp, err := c.SendCEcho(ctx, req)
if err != nil {
    return err
}

if !resp.Status().IsSuccess() {
    switch resp.StatusCode() {
    case 0xA700:
        // Out of Resources
    case 0xA900:
        // Data Set does not match SOP Class
    case 0xC000:
        // Cannot understand
    default:
        return fmt.Errorf("operation failed: status 0x%04X", resp.StatusCode())
    }
}
```

### 错误包装

```go
import "errors"

if err != nil {
    // 检查特定错误
    if errors.Is(err, context.DeadlineExceeded) {
        // 超时错误
    }

    if errors.Is(err, context.Canceled) {
        // Context 被取消
    }

    // 包装错误
    return fmt.Errorf("failed to store image: %w", err)
}
```

---

## 下一步

- 查看 [故障排查指南](./TROUBLESHOOTING.md) 解决常见问题
- 查看 [网络抓包示例](./WIRESHARK.md) 了解如何调试
- 查看 `examples/` 目录中的完整示例程序
