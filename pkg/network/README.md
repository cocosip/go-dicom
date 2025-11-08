# DICOM 网络功能

本包实现了完整的 DICOM 网络协议栈，支持 SCU (Service Class User) 和 SCP (Service Class Provider) 功能。

## 概述

DICOM 网络协议基于 TCP/IP，使用 Association 机制进行连接管理和 DIMSE (DICOM Message Service Element) 进行消息交换。本实现完整支持 DICOM 标准的网络功能。

### 主要特性

- ✅ **完整的 PDU 支持**：7种 PDU 类型（A-ASSOCIATE-RQ/AC/RJ, P-DATA-TF, A-RELEASE-RQ/RP, A-ABORT）
- ✅ **Association 管理**：自动协商 Presentation Contexts、MaxPDULength 等参数
- ✅ **所有 DIMSE 操作**：
  - C-ECHO（连通性验证）
  - C-STORE（存储图像）
  - C-FIND（查询）
  - C-MOVE（检索-Pull模式）
  - C-GET（检索-Push模式）
  - N-* 操作（N-EVENT-REPORT, N-GET, N-SET, N-ACTION, N-CREATE, N-DELETE）
- ✅ **SCU 客户端**：高级 API，支持同步和异步操作
- ✅ **SCP 服务器**：Handler 模式，支持自定义业务逻辑
- ✅ **TLS 支持**：安全连接
- ✅ **并发安全**：所有组件都是并发安全的

### 测试覆盖

- **401个单元测试**，全部通过
- 包含正常场景、错误场景、并发场景
- 代码覆盖率 > 85%

## 架构设计

### 分层架构

```
┌─────────────────────────────────────────────────────┐
│              Application Layer                       │
│         (Your DICOM Application)                     │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│            Client / Server API                       │
│  pkg/network/client  │  pkg/network/server          │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│              Service Layer                           │
│  pkg/network/service (State Machine + Handlers)     │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│              DIMSE Layer                             │
│  pkg/network/dimse (Message Encoding/Decoding)      │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│           Association Layer                          │
│  pkg/network/association (Negotiation)              │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│              PDU Layer                               │
│  pkg/network/pdu (Protocol Data Units)              │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│           Transport Layer                            │
│  pkg/network/transport (TCP/TLS)                    │
└─────────────────────────────────────────────────────┘
```

### 包结构

```
pkg/network/
├── status/         # DICOM 状态码定义
├── pdu/            # PDU 编解码
├── association/    # Association 管理
├── dimse/          # DIMSE 消息定义
├── transport/      # 网络传输（TCP/TLS）
├── service/        # 核心服务层（状态机、消息处理）
├── client/         # SCU 客户端 API
└── server/         # SCP 服务器 API
```

## 快速开始

### SCU 客户端示例

#### 1. C-ECHO（验证连接）

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    // 创建客户端
    c := client.New(
        client.WithCallingAE("MY_SCU"),
        client.WithCalledAE("PACS_SERVER"),
    )

    // 添加 Presentation Context（Verification SOP Class）
    c.AddPresentationContext(
        "1.2.840.10008.1.1",   // Verification SOP Class
        "1.2.840.10008.1.2",   // Implicit VR Little Endian
        "1.2.840.10008.1.2.1", // Explicit VR Little Endian
    )

    // 连接服务器
    ctx := context.Background()
    if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // 发送 C-ECHO
    if err := c.CEcho(ctx); err != nil {
        log.Fatalf("C-ECHO failed: %v", err)
    }

    log.Println("C-ECHO successful!")
}
```

#### 2. C-STORE（存储图像）

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/parser"
    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    // 读取 DICOM 文件
    result, err := parser.ParseFile("image.dcm")
    if err != nil {
        log.Fatal(err)
    }

    // 创建客户端
    c := client.New(
        client.WithCallingAE("MY_SCU"),
        client.WithCalledAE("PACS_SERVER"),
    )

    // 添加 Storage Presentation Context
    c.AddPresentationContext(
        "1.2.840.10008.5.1.4.1.1.2", // CT Image Storage
        "1.2.840.10008.1.2.1",       // Explicit VR Little Endian
    )

    // 连接并发送
    ctx := context.Background()
    if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // 发送 C-STORE
    if err := c.CStore(ctx, result.Dataset); err != nil {
        log.Fatalf("C-STORE failed: %v", err)
    }

    log.Println("Image stored successfully!")
}
```

#### 3. C-FIND（查询）

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/client"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
)

func main() {
    // 创建客户端
    c := client.New(
        client.WithCallingAE("MY_SCU"),
        client.WithCalledAE("PACS_SERVER"),
    )

    // 添加 Query/Retrieve Presentation Context
    c.AddPresentationContext(
        "1.2.840.10008.5.1.4.1.2.2.1", // Study Root Query/Retrieve
        "1.2.840.10008.1.2.1",
    )

    // 连接
    ctx := context.Background()
    if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // 构建查询条件
    query := dataset.New()
    query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN*"}))
    query.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20240101-20241231"}))
    query.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{""})) // 返回字段

    // 执行查询
    results, err := c.CFind(ctx, dimse.QueryRetrieveLevelStudy, query)
    if err != nil {
        log.Fatalf("C-FIND failed: %v", err)
    }

    // 处理结果
    log.Printf("Found %d studies\n", len(results))
    for i, result := range results {
        studyUID := result.GetString(tag.StudyInstanceUID, 0)
        patientName := result.GetString(tag.PatientName, 0)
        log.Printf("Study %d: %s - %s\n", i+1, studyUID, patientName)
    }
}
```

#### 4. 流式处理查询结果（避免内存占用）

```go
// 使用回调处理大量查询结果
err := c.CFindWithCallback(ctx, dimse.QueryRetrieveLevelStudy, query,
    func(result *dataset.Dataset) bool {
        // 处理每个结果
        studyUID := result.GetString(tag.StudyInstanceUID, 0)
        log.Printf("Found study: %s\n", studyUID)

        // 返回 true 继续接收，返回 false 停止
        return true
    })
```

### SCP 服务器示例

#### 基础 Echo SCP

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/network/dimse"
    "github.com/cocosip/go-dicom/pkg/network/server"
)

func main() {
    // 创建服务器
    srv := server.New(
        server.WithAETitle("MY_SCP"),
        server.WithPort(11112),
        server.WithMaxPDULength(16384),
    )

    // 设置 C-ECHO handler
    srv.SetCEchoHandler(func(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
        log.Println("Received C-ECHO request")
        return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
    })

    // 启动服务器
    log.Println("Starting Echo SCP on port 11112...")
    ctx := context.Background()
    if err := srv.ListenAndServe(ctx); err != nil {
        log.Fatal(err)
    }
}
```

#### Storage SCP（接收图像）

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/cocosip/go-dicom/pkg/dicom/writer"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
    "github.com/cocosip/go-dicom/pkg/network/server"
)

func main() {
    srv := server.New(
        server.WithAETitle("STORAGE_SCP"),
        server.WithPort(11112),
    )

    // 设置 C-STORE handler
    srv.SetCStoreHandler(func(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
        ds := req.DataDataset()

        // 获取 SOP Instance UID
        sopInstanceUID := ds.GetString(tag.SOPInstanceUID, 0)
        log.Printf("Receiving image: %s\n", sopInstanceUID)

        // 保存到文件
        filename := fmt.Sprintf("./received/%s.dcm", sopInstanceUID)
        if err := writer.WriteFile(filename, ds); err != nil {
            log.Printf("Failed to save image: %v\n", err)
            return dimse.NewCStoreResponseFromRequest(req, 0xC000), nil // Failure
        }

        log.Printf("Image saved to %s\n", filename)
        return dimse.NewCStoreResponseFromRequest(req, 0x0000), nil // Success
    })

    // 确保接收目录存在
    os.MkdirAll("./received", 0755)

    // 启动服务器
    log.Println("Starting Storage SCP on port 11112...")
    ctx := context.Background()
    if err := srv.ListenAndServe(ctx); err != nil {
        log.Fatal(err)
    }
}
```

#### Query SCP（查询服务）

```go
package main

import (
    "context"
    "log"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
    "github.com/cocosip/go-dicom/pkg/network/server"
)

func main() {
    srv := server.New(
        server.WithAETitle("QUERY_SCP"),
        server.WithPort(11112),
    )

    // 设置 C-FIND handler
    srv.SetCFindHandler(func(ctx context.Context, req *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
        query := req.DataDataset()
        level := req.QueryRetrieveLevel()

        log.Printf("Received C-FIND query at level %s\n", level)

        // 在数据库中查询（这里使用模拟数据）
        results := searchDatabase(query, level)

        // 构建响应
        responses := make([]*dimse.CFindResponse, 0, len(results)+1)

        // Pending responses（查询结果）
        for _, result := range results {
            resp := dimse.NewCFindResponsePending(req, result)
            responses = append(responses, resp)
        }

        // Final success response
        responses = append(responses, dimse.NewCFindResponseSuccess(req))

        return responses, nil
    })

    // 启动服务器
    log.Println("Starting Query SCP on port 11112...")
    ctx := context.Background()
    if err := srv.ListenAndServe(ctx); err != nil {
        log.Fatal(err)
    }
}

func searchDatabase(query *dataset.Dataset, level dimse.QueryRetrieveLevel) []*dataset.Dataset {
    // 模拟查询结果
    result := dataset.New()
    result.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))
    result.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
    result.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
    result.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20240115"}))

    return []*dataset.Dataset{result}
}
```

## 高级功能

### TLS 加密连接

#### 客户端使用 TLS

```go
import "crypto/tls"

// TLS 配置
tlsConfig := &tls.Config{
    InsecureSkipVerify: false, // 生产环境应该验证证书
    ServerName:         "pacs.hospital.com",
}

// 使用自定义 dialer
conn, err := transport.DialTLS(ctx, "tcp", "pacs.hospital.com:11112",
    transport.WithTLSConfig(tlsConfig),
    transport.WithTimeout(10*time.Second),
)
```

#### 服务器使用 TLS

```go
import "crypto/tls"

// 加载证书
cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
if err != nil {
    log.Fatal(err)
}

tlsConfig := &tls.Config{
    Certificates: []tls.Certificate{cert},
}

// 创建 TLS 服务器
srv := server.New(
    server.WithAETitle("SECURE_SCP"),
    server.WithPort(11112),
    server.WithTLSConfig(tlsConfig),
)
```

### Association 协商控制

#### 自定义 Presentation Context 协商

```go
// 服务器端：自定义 Association 协商
srv.SetAssociationNegotiatorFunc(func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
    // 检查 Calling AE Title
    if assoc.CallingAE != "AUTHORIZED_SCU" {
        return responder.RejectAssociation(ctx,
            pdu.ResultRejectedPermanent,
            pdu.SourceServiceUser,
            pdu.ReasonCallingAETitleNotRecognized)
    }

    // 只接受特定的 SOP Classes
    acceptedSOPs := map[string]bool{
        "1.2.840.10008.1.1": true,          // Verification
        "1.2.840.10008.5.1.4.1.1.2": true,  // CT Image Storage
    }

    for _, pc := range assoc.PresentationContexts {
        if acceptedSOPs[pc.AbstractSyntax] {
            pc.Accept(pdu.ResultAcceptance, pc.TransferSyntaxes[0])
        } else {
            pc.Reject(pdu.ResultAbstractSyntaxNotSupported)
        }
    }

    return responder.AcceptAssociation(ctx, assoc)
})
```

### 生命周期事件处理

```go
// 设置连接生命周期 handlers
srv.SetConnectionLifecycleHandlerFuncs(
    // OnAbort - 连接中止时调用
    func(ctx context.Context, source byte, reason byte) {
        log.Printf("Association aborted: source=%d, reason=%d\n", source, reason)
    },
    // OnConnectionClosed - 连接关闭时调用
    func(ctx context.Context, err error) {
        if err != nil {
            log.Printf("Connection closed with error: %v\n", err)
        } else {
            log.Println("Connection closed normally")
        }
    },
)

// 设置 Association 释放 handler
srv.SetAssociationReleaseHandlerFunc(func(ctx context.Context) error {
    log.Println("Client requesting association release")
    // 返回 nil 接受释放，返回 error 拒绝
    return nil
})
```

### 超时控制

```go
// 客户端配置超时
c := client.New(
    client.WithConnectTimeout(10*time.Second),      // TCP 连接超时
    client.WithAssociationTimeout(10*time.Second),  // Association 协商超时
    client.WithRequestTimeout(30*time.Second),      // DIMSE 请求超时
)

// 使用 context 控制单个操作的超时
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := c.CEcho(ctx); err != nil {
    // 处理超时错误
}
```

### 并发处理

```go
// 服务器自动处理并发连接
srv := server.New(
    server.WithMaxConnections(100), // 最大并发连接数
)

// 客户端可以并发发送多个请求（使用不同的连接）
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()

        c := client.New(/* ... */)
        if err := c.Connect(ctx, host, port); err != nil {
            log.Printf("Client %d: connect failed: %v\n", id, err)
            return
        }
        defer c.Close()

        if err := c.CEcho(ctx); err != nil {
            log.Printf("Client %d: echo failed: %v\n", id, err)
        }
    }(i)
}
wg.Wait()
```

## 性能优化建议

### 1. 复用客户端连接

```go
// ✅ 好的做法：复用连接
c := client.New(/* ... */)
c.Connect(ctx, host, port)
defer c.Close()

for i := 0; i < 100; i++ {
    c.CEcho(ctx)  // 复用同一个 Association
}

// ❌ 不好的做法：每次都重新连接
for i := 0; i < 100; i++ {
    c := client.New(/* ... */)
    c.Connect(ctx, host, port)
    c.CEcho(ctx)
    c.Close()  // 每次都重新协商 Association，开销大
}
```

### 2. 调整 MaxPDULength

```go
// 大文件传输时增加 PDU 长度
c := client.New(
    client.WithMaxPDULength(65536), // 默认 16384
)
```

### 3. 服务器资源限制

```go
srv := server.New(
    server.WithMaxConnections(50),              // 限制并发连接
    server.WithAssociationTimeout(30*time.Second), // 超时断开
)
```

## 状态码参考

常用 DICOM 状态码：

```go
import "github.com/cocosip/go-dicom/pkg/network/status"

// 成功
status.Success                    // 0x0000

// Pending（查询/检索操作）
status.Pending                    // 0xFF00

// Warning
status.AttributeListError         // 0x0107
status.AttributeValueOutOfRange   // 0x0116

// Failure
status.RefusedOutOfResources      // 0xA700
status.DataSetDoesNotMatchSOPClass // 0xA900
status.CannotUnderstand           // 0xC000
status.UnrecognizedOperation      // 0xC001
```

检查状态：

```go
resp, err := c.SendCEcho(ctx, req)
if err != nil {
    return err
}

if resp.Status().IsSuccess() {
    // 成功
} else if resp.Status().IsPending() {
    // 等待更多结果
} else if resp.Status().IsFailure() {
    // 失败
}
```

## 下一步

- 查看 [API 使用指南](./docs/API_GUIDE.md) 了解详细的 API 文档
- 查看 [故障排查指南](./docs/TROUBLESHOOTING.md) 解决常见问题
- 查看 [网络抓包示例](./docs/WIRESHARK.md) 了解如何调试 DICOM 网络通信
- 运行 `examples/` 目录中的示例程序

## 参考资料

- [DICOM 标准 - Part 7: Message Exchange](https://dicom.nema.org/medical/dicom/current/output/html/part07.html)
- [DICOM 标准 - Part 8: Network Communication Support](https://dicom.nema.org/medical/dicom/current/output/html/part08.html)
- [fo-dicom 项目](https://github.com/fo-dicom/fo-dicom) - C# 参考实现
