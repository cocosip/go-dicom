# DIMSE Package

DIMSE (DICOM Message Service Element) 是 DICOM 网络通信的核心消息层。

## 功能特性

- ✅ C-ECHO: 验证连接
- ✅ C-STORE: 存储 DICOM 实例
- ✅ C-FIND: 查询 DICOM 档案
- ✅ 自动 MessageID 管理
- ✅ 优先级支持
- ✅ 状态码处理

## MessageID 自动管理

### 问题
在 DICOM 网络通信中，每个请求需要一个唯一的 MessageID。手动管理 MessageID 容易出错：

```go
// ❌ 不推荐：手动管理 MessageID
msgID := 1
req1 := dimse.NewCEchoRequest(msgID)
msgID++
req2 := dimse.NewCStoreRequest(msgID, dataset)
msgID++
// ... 容易忘记递增或重复
```

### 解决方案：MessageIDGenerator

使用 `MessageIDGenerator` 自动管理 MessageID：

```go
// ✅ 推荐：使用 MessageIDGenerator
idGen := dimse.NewMessageIDGenerator()

// 创建请求时传入 0（表示未分配）
req1 := dimse.NewCEchoRequest(0)
req2 := dimse.NewCStoreRequest(0, dataset)
req3 := dimse.NewCFindRequestPatientRoot(0, level, query)

// 发送前自动分配 MessageID
idGen.AssignMessageID(req1)  // MessageID = 1
idGen.AssignMessageID(req2)  // MessageID = 2
idGen.AssignMessageID(req3)  // MessageID = 3
```

### 特性

1. **线程安全**：可以并发调用
2. **自动递增**：从 1 开始，范围 1-65535
3. **自动跳过 0**：MessageID=0 表示未分配
4. **保持已有值**：如果已有 MessageID，不会覆盖
5. **支持重置**：新连接时可以重置计数器

## 基本用法

### C-ECHO（验证连接）

```go
package main

import (
    "github.com/cocosip/go-dicom/pkg/network/dimse"
)

func main() {
    // 创建 MessageID 生成器
    idGen := dimse.NewMessageIDGenerator()

    // 创建 C-ECHO 请求（MessageID=0 表示未分配）
    req := dimse.NewCEchoRequest(0)

    // 分配 MessageID
    msgID, err := idGen.AssignMessageID(req)
    if err != nil {
        panic(err)
    }

    // 现在可以发送 req，它的 MessageID 已经被设置为 1
    println("MessageID:", msgID)  // 输出: MessageID: 1
}
```

### C-STORE（存储 DICOM 实例）

```go
package main

import (
    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
)

func main() {
    idGen := dimse.NewMessageIDGenerator()

    // 创建 DICOM 数据集
    ds := dataset.New()
    ds.Add(element.NewString(tag.SOPClassUID, vr.UI,
        []string{"1.2.840.10008.5.1.4.1.1.2"}))  // CT Image Storage
    ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI,
        []string{"1.2.3.4.5.6.7.8.9"}))
    ds.Add(element.NewString(tag.PatientName, vr.PN,
        []string{"Doe^John"}))

    // 创建 C-STORE 请求
    req, err := dimse.NewCStoreRequest(0, ds)
    if err != nil {
        panic(err)
    }

    // 设置优先级（可选）
    req.SetPriority(uint16(dimse.PriorityHigh))

    // 分配 MessageID
    msgID, _ := idGen.AssignMessageID(req)
    println("MessageID:", msgID)
}
```

### C-FIND（查询 DICOM 档案）

```go
package main

import (
    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
    "github.com/cocosip/go-dicom/pkg/network/dimse"
)

func main() {
    idGen := dimse.NewMessageIDGenerator()

    // 创建查询数据集
    query := dataset.New()
    query.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^*"}))
    query.Add(element.NewString(tag.PatientID, vr.LO, []string{""}))

    // 创建 C-FIND 请求（Patient Root Query）
    req := dimse.NewCFindRequestPatientRoot(
        0,
        dimse.QueryRetrieveLevelPatient,
        query,
    )

    // 分配 MessageID
    msgID, _ := idGen.AssignMessageID(req)
    println("MessageID:", msgID)
}
```

## 在 Client/Association 中集成

在实际应用中，MessageIDGenerator 应该由 Client 或 Association 管理：

```go
type Client struct {
    messageIDGen *dimse.MessageIDGenerator
    // ... 其他字段
}

func NewClient() *Client {
    return &Client{
        messageIDGen: dimse.NewMessageIDGenerator(),
    }
}

func (c *Client) Send(msg dimse.Message) error {
    // 自动分配 MessageID
    _, err := c.messageIDGen.AssignMessageID(msg)
    if err != nil {
        return err
    }

    // 发送消息
    // ...
    return nil
}

// 用户代码
func main() {
    client := NewClient()

    // 不需要关心 MessageID！
    req := dimse.NewCEchoRequest(0)
    client.Send(req)
}
```

## 响应处理

### C-ECHO 响应

```go
// 创建成功响应
resp := dimse.NewCEchoResponseSuccess(requestMessageID)

// 检查状态
if resp.IsSuccess() {
    println("Echo successful")
}
```

### C-STORE 响应

```go
resp := dimse.NewCStoreResponseSuccess(
    requestMessageID,
    sopClassUID,
    sopInstanceUID,
)

if resp.IsSuccess() {
    println("Store successful")
}
```

### C-FIND 响应

```go
// Pending 响应（包含查询结果）
resp := dimse.NewCFindResponsePending(
    requestMessageID,
    sopClassUID,
    identifier,
)

if resp.IsPending() {
    // 处理查询结果
    if resp.HasIdentifier() {
        data := resp.DataDataset()
        // ... 处理数据
    }
}

// 最终响应（查询完成）
finalResp := dimse.NewCFindResponseSuccess(requestMessageID, sopClassUID)
if finalResp.IsSuccess() {
    println("Query completed")
}
```

## 高级功能

### 优先级

```go
req := dimse.NewCStoreRequest(0, dataset)
req.SetPriority(uint16(dimse.PriorityHigh))    // 高优先级
req.SetPriority(uint16(dimse.PriorityMedium))  // 中优先级（默认）
req.SetPriority(uint16(dimse.PriorityLow))     // 低优先级
```

### Move Originator（C-MOVE 子操作）

```go
req, _ := dimse.NewCStoreRequest(0, dataset)
req.SetMoveOriginator("MOVE_SCU", 999)
```

### 重置 MessageID 计数器

```go
idGen := dimse.NewMessageIDGenerator()

// 使用一段时间后...
idGen.Next()  // 1
idGen.Next()  // 2

// 新连接开始，重置计数器
idGen.Reset()
idGen.Next()  // 1（重新开始）
```

## 后续开发

在 Phase 9.7 (Client API) 中，我们将提供更高层的 API：

```go
// 未来的 API（即将实现）
client := dicom.NewClient("MY_SCU")
client.Connect("localhost", 11112, "REMOTE_SCP")

// 完全不需要创建 Request 对象！
err := client.Echo()
err := client.Store(dataset)
results := client.Find(query)
```

## 测试

运行测试：

```bash
go test ./pkg/network/dimse
go test -v ./pkg/network/dimse -run TestMessageIDGenerator
```

## 参考

- DICOM 标准: https://www.dicomstandard.org/
- DICOM Part 7: Message Exchange
- DICOM Part 8: Network Communication Support
