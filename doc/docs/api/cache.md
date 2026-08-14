---
title: Cache
---

# Cache

缓存工具模块，提供多种缓存策略实现，包括 LRU（最近最少使用）、FIFO（先进先出）、LFU（最不经常使用）和基于时间的缓存。所有缓存实现均支持过期时间和线程安全操作。

## 导入

```go
import "github.com/gongm/gohutool/cache"
```

## 接口

### Cache

```go
type Cache interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{})
    SetWithExpire(key string, value interface{}, expire time.Duration)
    Delete(key string)
    Clear()
    Size() int
    Keys() []string
    Contains(key string) bool
}
```

Cache 是缓存的通用接口，所有缓存实现均满足此接口。

## LRU 缓存

LRU（Least Recently Used）缓存会优先淘汰最近最少使用的数据。

### NewLRU

```go
func NewLRU(capacity int) *LRUCache
```

创建指定容量的 LRU 缓存实例。

**示例:**

```go
cache := cache.NewLRU(100)
cache.Set("key1", "value1")
val, ok := cache.Get("key1")
fmt.Println(val, ok) // value1 true
```

### LRUCache.Get

```go
func (c *LRUCache) Get(key string) (interface{}, bool)
```

从缓存中获取值。如果键不存在或已过期，返回 `nil, false`。访问会更新元素在 LRU 队列中的位置。

### LRUCache.Set

```go
func (c *LRUCache) Set(key string, value interface{})
```

添加或更新缓存项。如果缓存已满，会自动淘汰最近最少使用的元素。

### LRUCache.SetWithExpire

```go
func (c *LRUCache) SetWithExpire(key string, value interface{}, expire time.Duration)
```

添加带有过期时间的缓存项。

**示例:**

```go
cache := cache.NewLRU(100)
cache.SetWithExpire("session", "abc123", 5*time.Minute)
```

### LRUCache.Delete

```go
func (c *LRUCache) Delete(key string)
```

从缓存中删除指定键。

### LRUCache.Clear

```go
func (c *LRUCache) Clear()
```

清空缓存中的所有数据。

### LRUCache.Size

```go
func (c *LRUCache) Size() int
```

返回缓存中的元素数量。

### LRUCache.Keys

```go
func (c *LRUCache) Keys() []string
```

返回缓存中所有未过期的键。

### LRUCache.Contains

```go
func (c *LRUCache) Contains(key string) bool
```

检查键是否存在于缓存中（且未过期）。

## FIFO 缓存

FIFO（First In First Out）缓存会优先淘汰最早加入的数据。

### NewFIFO

```go
func NewFIFO(capacity int) *FIFOCache
```

创建指定容量的 FIFO 缓存实例。

**示例:**

```go
cache := cache.NewFIFO(50)
cache.Set("first", 1)
cache.Set("second", 2)
```

### FIFOCache.Get

```go
func (c *FIFOCache) Get(key string) (interface{}, bool)
```

从缓存中获取值。与 LRU 不同，FIFO 的 Get 操作不会影响元素的淘汰顺序。

### FIFOCache.Set

```go
func (c *FIFOCache) Set(key string, value interface{})
```

添加或更新缓存项。缓存满时淘汰最早的元素。

### FIFOCache.SetWithExpire

```go
func (c *FIFOCache) SetWithExpire(key string, value interface{}, expire time.Duration)
```

添加带有过期时间的缓存项。

### FIFOCache.Delete

```go
func (c *FIFOCache) Delete(key string)
```

从缓存中删除指定键。

### FIFOCache.Clear

```go
func (c *FIFOCache) Clear()
```

清空缓存中的所有数据。

### FIFOCache.Size

```go
func (c *FIFOCache) Size() int
```

返回缓存中的元素数量。

### FIFOCache.Keys

```go
func (c *FIFOCache) Keys() []string
```

返回缓存中所有未过期的键。

### FIFOCache.Contains

```go
func (c *FIFOCache) Contains(key string) bool
```

检查键是否存在于缓存中（且未过期）。

## LFU 缓存

LFU（Least Frequently Used）缓存会优先淘汰使用频率最低的数据。

### NewLFU

```go
func NewLFU(capacity int) *LFUCache
```

创建指定容量的 LFU 缓存实例。

**示例:**

```go
cache := cache.NewLFU(100)
cache.Set("frequent", "data")
cache.Get("frequent") // 访问频率 +1
cache.Get("frequent") // 访问频率 +1
```

### LFUCache.Get

```go
func (c *LFUCache) Get(key string) (interface{}, bool)
```

从缓存中获取值。每次访问会增加该元素的使用频率。

### LFUCache.Set

```go
func (c *LFUCache) Set(key string, value interface{})
```

添加或更新缓存项。缓存满时淘汰使用频率最低的元素。

### LFUCache.SetWithExpire

```go
func (c *LFUCache) SetWithExpire(key string, value interface{}, expire time.Duration)
```

添加带有过期时间的缓存项。

### LFUCache.Delete

```go
func (c *LFUCache) Delete(key string)
```

从缓存中删除指定键。

### LFUCache.Clear

```go
func (c *LFUCache) Clear()
```

清空缓存中的所有数据。

### LFUCache.Size

```go
func (c *LFUCache) Size() int
```

返回缓存中的元素数量。

### LFUCache.Keys

```go
func (c *LFUCache) Keys() []string
```

返回缓存中所有未过期的键。

### LFUCache.Contains

```go
func (c *LFUCache) Contains(key string) bool
```

检查键是否存在于缓存中（且未过期）。

## 时间缓存

TimedCache 是基于过期时间的缓存，每个键值对都有独立的过期时间。后台会自动清理过期项。

### NewTimed

```go
func NewTimed(defaultExpire time.Duration) *TimedCache
```

创建具有默认过期时间的时间缓存实例。会启动后台协程定期清理过期项。

**示例:**

```go
cache := cache.NewTimed(10 * time.Minute)
cache.Set("key1", "value1")        // 使用默认过期时间
cache.SetWithExpire("key2", "value2", 30*time.Second) // 自定义过期时间
```

### TimedCache.Get

```go
func (c *TimedCache) Get(key string) (interface{}, bool)
```

从缓存中获取值。如果已过期则自动删除并返回 `nil, false`。

### TimedCache.Set

```go
func (c *TimedCache) Set(key string, value interface{})
```

使用默认过期时间添加缓存项。

### TimedCache.SetWithExpire

```go
func (c *TimedCache) SetWithExpire(key string, value interface{}, expire time.Duration)
```

使用自定义过期时间添加缓存项。

### TimedCache.Delete

```go
func (c *TimedCache) Delete(key string)
```

从缓存中删除指定键。

### TimedCache.Clear

```go
func (c *TimedCache) Clear()
```

清空缓存中的所有数据。

### TimedCache.Size

```go
func (c *TimedCache) Size() int
```

返回缓存中的元素数量（包括已过期但尚未清理的项）。

### TimedCache.Keys

```go
func (c *TimedCache) Keys() []string
```

返回缓存中所有未过期的键。

### TimedCache.Contains

```go
func (c *TimedCache) Contains(key string) bool
```

检查键是否存在且未过期。

## 完整示例

```go
package main

import (
    "fmt"
    "time"

    "github.com/gongm/gohutool/cache"
)

func main() {
    // LRU 缓存示例
    lru := cache.NewLRU(3)
    lru.Set("a", 1)
    lru.Set("b", 2)
    lru.Set("c", 3)
    fmt.Println(lru.Keys()) // [c b a]

    lru.Set("d", 4) // 淘汰 "a"
    fmt.Println(lru.Contains("a")) // false
    fmt.Println(lru.Size())        // 3

    // 时间缓存示例
    timed := cache.NewTimed(5 * time.Second)
    timed.Set("token", "xyz789")
    val, ok := timed.Get("token")
    fmt.Println(val, ok) // xyz789 true

    time.Sleep(6 * time.Second)
    val, ok = timed.Get("token")
    fmt.Println(val, ok) // <nil> false
}
```
