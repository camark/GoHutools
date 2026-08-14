---
title: Pool
---

# Pool

对象池模块，提供高效的对象复用机制，减少频繁创建和销毁对象带来的 GC 压力。支持通用对象池、带重置功能的对象池和字节缓冲区池。

## 导入

```go
import "github.com/gongm/gohutool/pool"
```

## 通用对象池

### New

```go
func New(factory func() interface{}) *Pool
```

创建通用对象池，使用工厂函数创建新对象。

**示例:**

```go
p := pool.New(func() interface{} {
    return &MyObject{}
})

obj := p.Get().(*MyObject)
// 使用对象...
p.Put(obj)
```

### Pool.Get

```go
func (p *Pool) Get() interface{}
```

从池中获取对象。如果池为空，会调用工厂函数创建新对象。

### Pool.Put

```go
func (p *Pool) Put(v interface{})
```

将对象归还到池中。

### Pool.Size

```go
func (p *Pool) Size() int
```

返回池中已创建的对象总数（近似值）。

## 带重置的对象池

### NewWithReset

```go
func NewWithReset(factory func() interface{}, reset func(interface{})) *PoolWithReset
```

创建带重置功能的对象池。归还对象时会自动调用 reset 函数清理对象状态。

**示例:**

```go
p := pool.NewWithReset(
    func() interface{} {
        return &bytes.Buffer{}
    },
    func(v interface{}) {
        v.(*bytes.Buffer).Reset()
    },
)

buf := p.Get().(*bytes.Buffer)
buf.WriteString("hello")
p.Put(buf) // 自动调用 Reset()
```

### PoolWithReset.Get

```go
func (p *PoolWithReset) Get() interface{}
```

从池中获取对象。

### PoolWithReset.Put

```go
func (p *PoolWithReset) Put(v interface{})
```

将对象归还到池中，归还前会调用 reset 函数。

### PoolWithReset.Size

```go
func (p *PoolWithReset) Size() int
```

返回池中已创建的对象总数（近似值）。

## 字节缓冲区池

### NewBufferPool

```go
func NewBufferPool(bufSize int) *BufferPool
```

创建字节缓冲区池。如果 `bufSize <= 0`，默认使用 4096 字节（4KB）。

**示例:**

```go
bp := pool.NewBufferPool(8192) // 8KB 缓冲区

buf := bp.Get()
// 使用缓冲区...
copy(buf, data)
bp.Put(buf) // 归还并重置长度
```

### BufferPool.Get

```go
func (p *BufferPool) Get() []byte
```

从池中获取字节缓冲区。返回的缓冲区长度为 0，容量为指定大小。

### BufferPool.Put

```go
func (p *BufferPool) Put(buf []byte)
```

将缓冲区归还到池中。归还前会重置长度为 0（保留容量）。

### BufferPool.Size

```go
func (p *BufferPool) Size() int
```

返回池中已创建的缓冲区总数（近似值）。

### BufferPool.BufSize

```go
func (p *BufferPool) BufSize() int
```

返回缓冲区的容量大小。

## 重置池

### NewResetPool

```go
func NewResetPool(factory func() interface{}, reset func(interface{})) *ResetPool
```

创建重置池，功能与 `NewWithReset` 相同。

### ResetPool.Get

```go
func (p *ResetPool) Get() interface{}
```

从池中获取对象。

### ResetPool.Put

```go
func (p *ResetPool) Put(v interface{})
```

将对象归还到池中，归还前调用 reset 函数。

### ResetPool.Size

```go
func (p *ResetPool) Size() int
```

返回池中已创建的对象总数（近似值）。

## 类型池

### NewTyped

```go
func NewTyped(factory func() interface{}) *TypedPool
```

创建类型安全的对象池，功能与 `New` 相同。

### TypedPool.Get

```go
func (p *TypedPool) Get() interface{}
```

从池中获取对象。

### TypedPool.Put

```go
func (p *TypedPool) Put(v interface{})
```

将对象归还到池中。

### TypedPool.Size

```go
func (p *TypedPool) Size() int
```

返回池中已创建的对象总数（近似值）。

## 完整示例

```go
package main

import (
    "bytes"
    "fmt"
    "sync"

    "github.com/gongm/gohutool/pool"
)

type Request struct {
    Method  string
    Path    string
    Headers map[string]string
    Body    []byte
}

func main() {
    // 请求对象池
    requestPool := pool.NewWithReset(
        func() interface{} {
            return &Request{
                Headers: make(map[string]string),
            }
        },
        func(v interface{}) {
            r := v.(*Request)
            r.Method = ""
            r.Path = ""
            r.Headers = make(map[string]string)
            r.Body = r.Body[:0]
        },
    )

    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            req := requestPool.Get().(*Request)
            req.Method = "GET"
            req.Path = fmt.Sprintf("/api/%d", id)

            // 处理请求...

            requestPool.Put(req)
        }(i)
    }
    wg.Wait()

    // 字节缓冲区池
    bufPool := pool.NewBufferPool(4096)
    buf := bufPool.Get()
    buf = append(buf, []byte("Hello, World!")...)
    fmt.Println(string(buf))
    bufPool.Put(buf)
}
```
