---
title: Bloom Filter
---

# Bloom Filter

布隆过滤器模块，提供高效的概率型数据结构，用于判断元素是否存在于集合中。支持自动计算最优参数，线程安全。

布隆过滤器的特点：
- 空间效率极高
- 查询时间 O(k)，k 为哈希函数数量
- 可能存在假阳性（误判存在），但不会有假阴性（漏判）

## 导入

```go
import "github.com/camark/GoHutools/bloom"
```

## 函数列表

### New

```go
func New(expectedItems uint64, fpRate float64) *Filter
```

创建布隆过滤器，自动计算最优的位数组大小和哈希函数数量。

- `expectedItems`: 预期存储的元素数量
- `fpRate`: 期望的假阳性率（0 到 1 之间）

**示例:**

```go
// 预期存储 100 万个元素，假阳性率 1%
f := bloom.New(1000000, 0.01)
```

### NewWithSize

```go
func NewWithSize(size uint64, hashCount uint) *Filter
```

使用指定的位数组大小和哈希函数数量创建布隆过滤器。适用于需要手动调优参数的场景。

**示例:**

```go
f := bloom.NewWithSize(10000000, 7)
```

### Filter.Add

```go
func (f *Filter) Add(data []byte)
```

将字节数据添加到布隆过滤器中。

### Filter.AddString

```go
func (f *Filter) AddString(s string)
```

将字符串添加到布隆过滤器中。

**示例:**

```go
f := bloom.New(10000, 0.01)
f.AddString("user:1001")
f.AddString("user:1002")
f.Add([]byte("item:5001"))
```

### Filter.Contains

```go
func (f *Filter) Contains(data []byte) bool
```

检查字节数据是否可能存在于过滤器中。返回 `true` 表示可能存在（有假阳性概率），返回 `false` 表示一定不存在。

### Filter.ContainsString

```go
func (f *Filter) ContainsString(s string) bool
```

检查字符串是否可能存在于过滤器中。

**示例:**

```go
f := bloom.New(10000, 0.01)
f.AddString("hello")

fmt.Println(f.ContainsString("hello"))     // true
fmt.Println(f.ContainsString("world"))     // false
```

### Filter.Count

```go
func (f *Filter) Count() uint64
```

返回已添加到过滤器中的元素数量。

### Filter.Reset

```go
func (f *Filter) Reset()
```

清空过滤器，重置所有位和计数器。

### Filter.FalsePositiveRate

```go
func (f *Filter) FalsePositiveRate() float64
```

返回当前估计的假阳性率。随着元素的增多，假阳性率会上升。

**示例:**

```go
f := bloom.New(1000, 0.01)
for i := 0; i < 1000; i++ {
    f.AddString(fmt.Sprintf("item:%d", i))
}
fmt.Printf("当前假阳性率: %.4f\n", f.FalsePositiveRate())
```

## 完整示例

```go
package main

import (
    "fmt"

    "github.com/camark/GoHutools/bloom"
)

func main() {
    // 创建布隆过滤器：预期 10 万个 URL，假阳性率 0.1%
    f := bloom.New(100000, 0.001)

    // 添加已爬取的 URL
    urls := []string{
        "https://example.com/page1",
        "https://example.com/page2",
        "https://example.com/page3",
    }
    for _, url := range urls {
        f.AddString(url)
    }

    // 检查 URL 是否已爬取
    checkURL := "https://example.com/page1"
    if f.ContainsString(checkURL) {
        fmt.Printf("%s 可能已爬取\n", checkURL)
    } else {
        fmt.Printf("%s 未爬取，可以添加到队列\n", checkURL)
    }

    // 检查新 URL
    newURL := "https://example.com/page99"
    if f.ContainsString(newURL) {
        fmt.Printf("%s 可能已爬取\n", newURL)
    } else {
        fmt.Printf("%s 未爬取，可以添加到队列\n", newURL)
    }

    fmt.Printf("已添加 %d 个元素\n", f.Count())
    fmt.Printf("当前假阳性率: %.6f\n", f.FalsePositiveRate())
}
```
