# DICOM 网络故障排查指南

本文档列出了使用 go-dicom 网络功能时的常见问题和解决方案。

## 目录

- [连接问题](#连接问题)
- [Association 问题](#association-问题)
- [DIMSE 操作问题](#dimse-操作问题)
- [性能问题](#性能问题)
- [调试技巧](#调试技巧)

---

## 连接问题

### 问题：连接被拒绝 (connection refused)

**错误信息**:
```
failed to connect to 192.168.1.100:104: dial tcp 192.168.1.100:104: connect: connection refused
```

**可能原因**:
1. 服务器未运行
2. 端口号错误
3. 防火墙阻止连接
4. 服务器监听在错误的地址上

**解决方案**:

```bash
# 1. 检查服务器是否运行
netstat -an | grep 104
# 或
ss -tlnp | grep 104

# 2. 检查端口是否可达
telnet 192.168.1.100 104
# 或
nc -vz 192.168.1.100 104

# 3. 检查防火墙
# Linux
sudo iptables -L -n | grep 104
# Windows
netsh advfirewall firewall show rule name=all | findstr 104

# 4. 验证服务器配置
# 确保服务器监听在正确的地址上
srv := server.New(
    server.WithPort(104),  // 确保端口正确
)
```

### 问题：连接超时 (connection timeout)

**错误信息**:
```
failed to connect: context deadline exceeded
```

**可能原因**:
1. 网络延迟过高
2. 服务器响应慢
3. 网络不可达

**解决方案**:

```go
// 增加连接超时
c := client.New(
    client.WithConnectTimeout(30*time.Second),  // 增加到 30 秒
)

// 或使用 context 控制
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
err := c.Connect(ctx, host, port)
```

### 问题：TLS 握手失败

**错误信息**:
```
TLS handshake failed: x509: certificate signed by unknown authority
```

**解决方案**:

```go
import "crypto/tls"

// 开发环境：跳过证书验证（不推荐用于生产）
tlsConfig := &tls.Config{
    InsecureSkipVerify: true,
}

// 生产环境：添加自定义 CA 证书
caCert, err := os.ReadFile("ca.crt")
if err != nil {
    log.Fatal(err)
}

caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)

tlsConfig := &tls.Config{
    RootCAs: caCertPool,
    ServerName: "pacs.hospital.com",
}

conn, err := transport.DialTLS(ctx, "tcp", "pacs.hospital.com:11112",
    transport.WithTLSConfig(tlsConfig))
```

---

## Association 问题

### 问题：Association 被拒绝 - AE Title 不匹配

**错误信息**:
```
association rejected: called AE title not recognized
```

**可能原因**:
1. Called AE Title 不正确
2. 服务器配置了严格的 AE Title 检查
3. AE Title 大小写敏感

**解决方案**:

```go
// 1. 确认服务器的 AE Title
c := client.New(
    client.WithCalledAE("EXACT_SERVER_AE"),  // 必须精确匹配
)

// 2. 检查服务器配置
srv := server.New(
    server.WithAETitle("MY_SCP"),
    server.WithStrictAECheck(false),  // 关闭严格检查（开发环境）
)

// 3. 或者配置允许的 Calling AE Titles
srv := server.New(
    server.WithAcceptedCallingAETitles([]string{
        "SCU_1",
        "SCU_2",
        "WORKSTATION_*",  // 支持通配符
    }),
)
```

### 问题：没有 Presentation Context 被接受

**错误信息**:
```
invalid A-ASSOCIATE-AC: all presentation contexts were rejected
```

**可能原因**:
1. 没有添加 Presentation Context
2. SOP Class UID 不匹配
3. Transfer Syntax 不支持

**解决方案**:

```go
// 1. 添加 Presentation Contexts
c := client.New(/* ... */)

// 添加 Verification
c.AddPresentationContext(
    "1.2.840.10008.1.1",      // Verification SOP Class
    "1.2.840.10008.1.2",      // Implicit VR Little Endian
    "1.2.840.10008.1.2.1",    // Explicit VR Little Endian
)

// 添加 Storage
c.AddPresentationContext(
    "1.2.840.10008.5.1.4.1.1.2",  // CT Image Storage
    "1.2.840.10008.1.2.1",        // Explicit VR Little Endian
)

// 2. 服务器端：确保接受这些 SOP Classes
srv.SetAssociationNegotiatorFunc(
    func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
        for _, pc := range assoc.PresentationContexts {
            // 接受所有 SOP Classes（开发环境）
            pc.Accept(pdu.ResultAcceptance, pc.TransferSyntaxes[0])
        }
        return responder.AcceptAssociation(ctx, assoc)
    })
```

### 问题：MaxPDULength 协商失败

**错误信息**:
```
PDU length 65536 exceeds maximum allowed 16384
```

**解决方案**:

```go
// 客户端和服务器必须协商一致的 MaxPDULength
// 实际使用的是两者中较小的值

// 客户端
c := client.New(
    client.WithMaxPDULength(16384),  // 16 KB
)

// 服务器
srv := server.New(
    server.WithMaxPDULength(16384),  // 16 KB
)

// 大文件传输时可以增加（需要双方都支持）
c := client.New(
    client.WithMaxPDULength(65536),  // 64 KB
)
```

---

## DIMSE 操作问题

### 问题：C-STORE 失败 - SOP Class UID 缺失

**错误信息**:
```
failed to create C-STORE request: SOPClassUID is required
```

**解决方案**:

```go
import (
    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
    "github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// 确保 Dataset 包含必需的元素
ds := dataset.New()
ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6"}))
// ... 其他必需元素

// 然后发送
err := c.CStore(ctx, ds)
```

### 问题：C-FIND 返回空结果

**可能原因**:
1. 查询条件不匹配
2. Query/Retrieve Level 错误
3. 必需字段缺失

**解决方案**:

```go
// 1. 检查查询条件
query := dataset.New()
query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))  // 精确匹配
query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE*"}))      // 通配符匹配
query.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20240101-"}))   // 日期范围

// 2. 添加返回字段（空字符串表示返回该字段）
query.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{""}))
query.Add(element.NewString(tag.StudyDescription, vr.LO, []string{""}))
query.Add(element.NewString(tag.StudyDate, vr.DA, []string{""}))

// 3. 确保 Level 正确
results, err := c.CFind(ctx, dimse.QueryRetrieveLevelStudy, query)

// 4. 调试：打印查询条件
for _, elem := range query.Elements() {
    log.Printf("Query: %s = %v\n", elem.Tag(), elem.Value())
}
```

### 问题：C-MOVE/C-GET 失败 - 目标不可达

**错误信息**:
```
C-MOVE failed: destination AE not found
```

**解决方案**:

```go
// 1. 确保目标 AE Title 正确
err := c.CMove(ctx, level, "EXACT_DEST_AE", identifier, callback)

// 2. 服务器端：验证目标是否存在
// 可以维护一个 AE Title -> Host:Port 的映射
var knownAEs = map[string]string{
    "DEST_SCP_1": "192.168.1.101:104",
    "DEST_SCP_2": "192.168.1.102:104",
}

func getDestinationHost(ae string) (string, error) {
    host, ok := knownAEs[ae]
    if !ok {
        return "", fmt.Errorf("destination AE not found: %s", ae)
    }
    return host, nil
}
```

### 问题：请求超时

**错误信息**:
```
C-FIND failed: context deadline exceeded
```

**解决方案**:

```go
// 1. 增加请求超时
c := client.New(
    client.WithRequestTimeout(60*time.Second),  // 增加到 60 秒
)

// 2. 为特定操作使用自定义超时
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

results, err := c.CFind(ctx, level, query)

// 3. 对于长时间操作，使用流式处理
err := c.CFindWithCallback(ctx, level, query,
    func(result *dataset.Dataset) bool {
        // 立即处理每个结果，不等待全部完成
        processResult(result)
        return true
    })
```

---

## 性能问题

### 问题：大文件传输慢

**症状**:
传输大型 DICOM 文件（如 CT 序列）非常慢

**解决方案**:

```go
// 1. 增加 MaxPDULength
c := client.New(
    client.WithMaxPDULength(65536),  // 64 KB（默认 16 KB）
)

srv := server.New(
    server.WithMaxPDULength(65536),
)

// 2. 确保网络带宽足够
// 使用 iperf 测试网络速度
// iperf -s  # 服务器端
// iperf -c <server-ip>  # 客户端

// 3. 批量传输时复用连接
c := client.New(/* ... */)
c.Connect(ctx, host, port)
defer c.Close()

for _, file := range files {
    ds, _ := parser.ParseFile(file)
    c.CStore(ctx, ds)  // 复用同一个连接
}
```

### 问题：并发连接数过高导致资源耗尽

**症状**:
服务器内存占用过高，或出现 "too many open files" 错误

**解决方案**:

```go
// 1. 限制服务器并发连接
srv := server.New(
    server.WithMaxConnections(50),  // 最多 50 个并发连接
)

// 2. 设置连接超时
srv := server.New(
    server.WithAssociationTimeout(30*time.Second),  // 30 秒后断开空闲连接
)

// 3. 增加文件描述符限制（Linux）
// ulimit -n 65536

// 4. 客户端使用连接池（如果需要频繁连接）
// 注意：DICOM 连接通常是长连接，不需要连接池
```

### 问题：内存占用过高

**症状**:
处理大量查询结果时内存占用过高

**解决方案**:

```go
// ❌ 不好的做法：收集所有结果到内存
results, err := c.CFind(ctx, level, query)  // 如果有 10000 个结果会占用大量内存

// ✅ 好的做法：使用流式处理
err := c.CFindWithCallback(ctx, level, query,
    func(result *dataset.Dataset) bool {
        // 立即处理每个结果
        processAndSaveResult(result)

        // 结果处理完后就可以被垃圾回收
        return true
    })
```

---

## 调试技巧

### 启用详细日志

```go
import "log"

// 设置日志级别
log.SetFlags(log.LstdFlags | log.Lshortfile)

// 在关键位置添加日志
log.Printf("Connecting to %s:%d\n", host, port)
log.Printf("Sending C-ECHO request\n")
log.Printf("Received response: status=0x%04X\n", resp.StatusCode())
```

### 使用 Wireshark 抓包

参见 [网络抓包示例](./WIRESHARK.md)

### 检查 Association 详情

```go
// 连接后检查 Association
assoc := c.GetAssociation()
if assoc != nil {
    log.Printf("Association established:\n")
    log.Printf("  Called AE:  %s\n", assoc.CalledAE)
    log.Printf("  Calling AE: %s\n", assoc.CallingAE)
    log.Printf("  Max PDU:    %d\n", assoc.MaxPDULength)

    log.Printf("Accepted Presentation Contexts:\n")
    for _, pc := range assoc.GetAcceptedPresentationContexts() {
        log.Printf("  ID=%d: %s (%s)\n",
            pc.ID, pc.AbstractSyntax, pc.TransferSyntax)
    }
}
```

### 打印 DIMSE 消息

```go
// 打印请求信息
log.Printf("C-STORE Request:\n")
log.Printf("  SOP Class UID:    %s\n", req.AffectedSOPClassUID())
log.Printf("  SOP Instance UID: %s\n", req.AffectedSOPInstanceUID())
log.Printf("  Message ID:       %d\n", req.MessageID())

// 打印 Dataset 内容
ds := req.DataDataset()
for _, elem := range ds.Elements() {
    log.Printf("  %s: %v\n", elem.Tag(), elem.Value())
}
```

### 测试网络连通性

```go
// 简单的 Echo 测试程序
package main

import (
    "context"
    "log"
    "github.com/cocosip/go-dicom/pkg/network/client"
)

func main() {
    c := client.New(
        client.WithCallingAE("TEST_SCU"),
        client.WithCalledAE("PACS_SCP"),
    )

    c.AddPresentationContext(
        "1.2.840.10008.1.1",   // Verification
        "1.2.840.10008.1.2.1",
    )

    ctx := context.Background()
    log.Println("Connecting...")
    if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
        log.Fatalf("Connect failed: %v", err)
    }
    defer c.Close()

    log.Println("Sending C-ECHO...")
    if err := c.CEcho(ctx); err != nil {
        log.Fatalf("C-ECHO failed: %v", err)
    }

    log.Println("Success!")
}
```

### 常用的 DICOM 工具

用于对比测试：

- **dcmtk**: DICOM Toolkit（命令行工具）
  ```bash
  # C-ECHO 测试
  echoscu -aet MY_SCU -aec PACS_SCP 192.168.1.100 104

  # C-STORE
  storescu -aet MY_SCU -aec PACS_SCP 192.168.1.100 104 image.dcm

  # C-FIND
  findscu -aet MY_SCU -aec PACS_SCP 192.168.1.100 104 query.dcm
  ```

- **Orthanc**: 轻量级 PACS 服务器
  ```bash
  # Docker 运行
  docker run -p 4242:4242 -p 8042:8042 jodogne/orthanc
  ```

- **DCMTK 的 storescp**: 简单的 Storage SCP
  ```bash
  # 启动 Storage SCP
  storescp -aet STORAGE_SCP 11112
  ```

---

## 常见错误代码

| 状态码 | 含义 | 可能原因 | 解决方案 |
|--------|------|----------|----------|
| 0x0000 | Success | 成功 | - |
| 0xFF00 | Pending | 等待更多结果 | 继续接收 |
| 0x0107 | Attribute list error | 查询字段错误 | 检查查询条件 |
| 0x0116 | Attribute value out of range | 字段值超出范围 | 检查值的范围 |
| 0xA700 | Out of resources | 服务器资源不足 | 减少并发或增加服务器资源 |
| 0xA900 | Dataset does not match SOP Class | Dataset 与 SOP Class 不匹配 | 检查必需字段 |
| 0xC000 | Cannot understand | 无法理解请求 | 检查请求格式 |
| 0xC001 | Unrecognized operation | 不支持的操作 | 检查 SOP Class |

---

## 获取帮助

如果问题仍未解决：

1. **查看日志**：启用详细日志并查看错误信息
2. **抓包分析**：使用 Wireshark 分析 DICOM 流量
3. **对比测试**：使用 dcmtk 等工具测试服务器
4. **查看源码**：查看 go-dicom 源码了解实现细节
5. **提交 Issue**：在 GitHub 上提交 Issue 并附上日志和抓包

---

## 下一步

- 查看 [网络抓包示例](./WIRESHARK.md) 了解如何使用 Wireshark 调试
- 查看 [API 使用指南](./API_GUIDE.md) 了解详细的 API 文档
