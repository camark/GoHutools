---
title: idutil
---

# idutil

ID 生成工具包，提供 UUID、Snowflake、ULID、NanoID、ObjectID、序列号等多种分布式唯一 ID 生成方案。

## 导入

```go
import "github.com/gongm/gohutool/idutil"
```

## UUID

### UUID

```go
func UUID() string
```

生成 UUID v4 字符串（带连字符，36 字符）。

**示例:**

```go
id := idutil.UUID()
fmt.Println(id) // 550e8400-e29b-41d4-a716-446655440000
```

### UUIDBytes

```go
func UUIDBytes() []byte
```

生成 UUID v4 的 16 字节切片。

**示例:**

```go
bytes := idutil.UUIDBytes()
```

### SimpleUUID

```go
func SimpleUUID() string
```

生成不带连字符的 UUID v4 字符串（32 字符）。

**示例:**

```go
id := idutil.SimpleUUID()
fmt.Println(id) // 550e8400e29b41d4a716446655440000
```

## Snowflake ID

### SnowflakeID

```go
type SnowflakeID struct {
    // 包含导出或未导出字段
}
```

分布式唯一 ID 生成器，灵感来自 Twitter Snowflake。生成的 ID 为 64 位整数，包含时间戳、机器 ID 和序列号。

### NewSnowflake

```go
func NewSnowflake(machineID int64) *SnowflakeID
```

创建 Snowflake ID 生成器。machineID 必须在 [0, 1023] 范围内。

**示例:**

```go
sf := idutil.NewSnowflake(1)
```

### NextID

```go
func (s *SnowflakeID) NextID() (int64, error)
```

生成下一个 Snowflake ID。

**示例:**

```go
sf := idutil.NewSnowflake(1)
id, err := sf.NextID()
fmt.Println(id)
```

## ULID

### ULID

```go
func ULID() string
```

生成 ULID（Universally Unique Lexicographically Sortable Identifier）。格式为 26 字符，使用 Crockford Base32 编码。可排序。

**示例:**

```go
id := idutil.ULID()
fmt.Println(id) // 01ARZ3NDEKTSV4RRFFQ69G5FAV
```

## NanoID

### NanoID

```go
func NanoID() string
```

生成默认长度（21 字符）的 NanoID。

**示例:**

```go
id := idutil.NanoID()
fmt.Println(id) // V1StGXR8_Z5jdHi6B-myT
```

### NanoIDWithSize

```go
func NanoIDWithSize(size int) string
```

生成指定长度的 NanoID。

**示例:**

```go
id := idutil.NanoIDWithSize(10)
fmt.Println(id)
```

## Short / Random ID

### ShortID

```go
func ShortID() string
```

生成 8 字符的短 ID（字母数字）。

**示例:**

```go
id := idutil.ShortID()
fmt.Println(id) // aB3xK9mN
```

### RandomID

```go
func RandomID(length int) string
```

生成指定长度的随机字母数字 ID。

**示例:**

```go
id := idutil.RandomID(16)
fmt.Println(id)
```

### RandomIDWithAlphabet

```go
func RandomIDWithAlphabet(length int, alphabet string) string
```

使用自定义字母表生成指定长度的随机 ID。

**示例:**

```go
id := idutil.RandomIDWithAlphabet(8, "0123456789")
fmt.Println(id) // 纯数字 ID
```

## Sequence

### Sequence

```go
type Sequence struct {
    // 包含导出或未导出字段
}
```

线程安全的自增序列生成器。

### NewSequence

```go
func NewSequence(start, step int64) *Sequence
```

创建序列，从 start 开始，每次递增 step。

**示例:**

```go
seq := idutil.NewSequence(1, 1)
fmt.Println(seq.Next()) // 1
fmt.Println(seq.Next()) // 2
fmt.Println(seq.Next()) // 3
```

### NewSequenceWithRange

```go
func NewSequenceWithRange(start, end, step int64, cycled bool) *Sequence
```

创建带范围的序列。如果 cycled 为 true，超过 end 后会回到 start。

**示例:**

```go
seq := idutil.NewSequenceWithRange(1, 100, 1, true)
for i := 0; i < 5; i++ {
    fmt.Println(seq.Next())
}
```

### Next

```go
func (s *Sequence) Next() int64
```

获取序列的下一个值。

### NextN

```go
func (s *Sequence) NextN(n int) []int64
```

获取序列的下 n 个值。

**示例:**

```go
seq := idutil.NewSequence(1, 1)
values := seq.NextN(5)
fmt.Println(values) // [1 2 3 4 5]
```

### Reset

```go
func (s *Sequence) Reset()
```

重置序列到初始值。

### Current

```go
func (s *Sequence) Current() int64
```

获取当前值（不推进序列）。

## ObjectID

### ObjectID

```go
func ObjectID() string
```

生成 MongoDB 风格的 ObjectID（24 字符十六进制字符串）。格式为：4 字节时间戳 + 3 字节机器标识 + 2 字节进程 ID + 3 字节计数器。

**示例:**

```go
id := idutil.ObjectID()
fmt.Println(id) // 507f1f77bcf86cd799439011
```

## 机器/Worker 辅助

### MachineID

```go
func MachineID() (string, error)
```

基于 MAC 地址生成机器标识符（十六进制字符串）。

**示例:**

```go
id, err := idutil.MachineID()
if err != nil {
    log.Fatal(err)
}
fmt.Println(id)
```

### WorkerID

```go
func WorkerID() (int64, error)
```

基于第一个非回环 IP 地址生成 Worker ID（0-1023），适用于 Snowflake。

**示例:**

```go
wid, err := idutil.WorkerID()
if err != nil {
    log.Fatal(err)
}
sf := idutil.NewSnowflake(wid)
```

## 完整示例

```go
package main

import (
    "fmt"
    "log"

    "github.com/gongm/gohutool/idutil"
)

func main() {
    // UUID
    fmt.Println("UUID:", idutil.UUID())
    fmt.Println("Simple UUID:", idutil.SimpleUUID())

    // Snowflake
    sf := idutil.NewSnowflake(1)
    for i := 0; i < 3; i++ {
        id, _ := sf.NextID()
        fmt.Println("Snowflake:", id)
    }

    // ULID
    fmt.Println("ULID:", idutil.ULID())

    // NanoID
    fmt.Println("NanoID:", idutil.NanoID())
    fmt.Println("NanoID(10):", idutil.NanoIDWithSize(10))

    // Short ID
    fmt.Println("Short ID:", idutil.ShortID())

    // ObjectID
    fmt.Println("ObjectID:", idutil.ObjectID())

    // Sequence
    seq := idutil.NewSequence(100, 5)
    fmt.Println(seq.Next()) // 100
    fmt.Println(seq.Next()) // 105
    fmt.Println(seq.Next()) // 110

    // Worker ID for Snowflake
    wid, err := idutil.WorkerID()
    if err != nil {
        log.Fatal(err)
    }
    sf2 := idutil.NewSnowflake(wid)
    id, _ := sf2.NextID()
    fmt.Println("Snowflake with WorkerID:", id)
}
```
