---
title: maputil - Map工具
---

# maputil - Map工具

Map 操作工具包，提供了丰富的 Map 操作函数，包括键值查询、过滤、转换、合并等功能。使用 Go 泛型实现，类型安全。

## 导入

```go
import "github.com/gongm/gohutool/maputil"
```

## 类型定义

### Entry

```go
type Entry[K comparable, V any] struct {
    Key   K
    Value V
}
```

表示键值对的结构体。

## 函数列表

### IsEmpty

```go
func IsEmpty[K comparable, V any](m map[K]V) bool
```

检查 Map 是否为空。

**示例:**

```go
m := map[string]int{}
fmt.Println(maputil.IsEmpty(m))  // true

m = map[string]int{"a": 1}
fmt.Println(maputil.IsEmpty(m))  // false
```

### IsNotEmpty

```go
func IsNotEmpty[K comparable, V any](m map[K]V) bool
```

检查 Map 是否不为空。

**示例:**

```go
m := map[string]int{"a": 1}
fmt.Println(maputil.IsNotEmpty(m))  // true
```

### ContainsKey

```go
func ContainsKey[K comparable, V any](m map[K]V, key K) bool
```

检查 Map 是否包含指定键。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2}
fmt.Println(maputil.ContainsKey(m, "a"))  // true
fmt.Println(maputil.ContainsKey(m, "c"))  // false
```

### ContainsValue

```go
func ContainsValue[K comparable, V any](m map[K]V, value V) bool
```

检查 Map 是否包含指定值。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2}
fmt.Println(maputil.ContainsValue(m, 2))  // true
fmt.Println(maputil.ContainsValue(m, 3))  // false
```

### Get

```go
func Get[K comparable, V any](m map[K]V, key K, defaultVal V) V
```

获取 Map 中指定键的值，如果键不存在返回默认值。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2}
fmt.Println(maputil.Get(m, "a", 0))  // 1
fmt.Println(maputil.Get(m, "c", 0))  // 0
```

### GetOrDefault

```go
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultVal V) V
```

获取值或返回默认值（Get 的别名）。

**示例:**

```go
m := map[string]int{"a": 1}
fmt.Println(maputil.GetOrDefault(m, "a", 0))  // 1
fmt.Println(maputil.GetOrDefault(m, "b", 0))  // 0
```

### PutIfAbsent

```go
func PutIfAbsent[K comparable, V any](m map[K]V, key K, value V) V
```

如果键不存在则插入，返回已存在的值或新插入的值。

**示例:**

```go
m := map[string]int{"a": 1}
maputil.PutIfAbsent(m, "b", 2)  // 插入 b=2
maputil.PutIfAbsent(m, "a", 3)  // a 已存在，不修改
fmt.Println(m)  // map[a:1 b:2]
```

### Remove

```go
func Remove[K comparable, V any](m map[K]V, key K)
```

从 Map 中移除指定键。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
maputil.Remove(m, "b")
fmt.Println(m)  // map[a:1 c:3]
```

### Keys

```go
func Keys[K comparable, V any](m map[K]V) []K
```

返回 Map 的所有键。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
keys := maputil.Keys(m)
fmt.Println(keys)  // [a b c] (顺序可能不同)
```

### Values

```go
func Values[K comparable, V any](m map[K]V) []V
```

返回 Map 的所有值。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
values := maputil.Values(m)
fmt.Println(values)  // [1 2 3] (顺序可能不同)
```

### ForEach

```go
func ForEach[K comparable, V any](m map[K]V, consumer func(K, V))
```

遍历 Map，对每个键值对执行操作。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
maputil.ForEach(m, func(k string, v int) {
    fmt.Printf("%s: %d\n", k, v)
})
```

### Filter

```go
func Filter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V
```

过滤 Map，返回满足条件的键值对。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := maputil.Filter(m, func(k string, v int) bool {
    return v > 2
})
fmt.Println(result)  // map[c:3 d:4]
```

### Transform

```go
func Transform[K comparable, V, R any](m map[K]V, transformer func(K, V) R) map[K]R
```

转换 Map 的值。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
result := maputil.Transform(m, func(k string, v int) string {
    return fmt.Sprintf("%d items", v)
})
fmt.Println(result)  // map[a:1 items b:2 items c:3 items]
```

### Merge

```go
func Merge[K comparable, V any](m1, m2 map[K]V) map[K]V
```

合并两个 Map，m2 的值会覆盖 m1 中相同键的值。

**示例:**

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}
result := maputil.Merge(m1, m2)
fmt.Println(result)  // map[a:1 b:3 c:4]
```

### Invert

```go
func Invert[K, V comparable](m map[K]V) map[V]K
```

反转 Map（键变值，值变键）。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
result := maputil.Invert(m)
fmt.Println(result)  // map[1:a 2:b 3:c]
```

### FromKeys

```go
func FromKeys[K comparable, V any](keys []K, defaultVal V) map[K]V
```

从键列表创建 Map，所有值使用默认值。

**示例:**

```go
keys := []string{"a", "b", "c"}
result := maputil.FromKeys(keys, 0)
fmt.Println(result)  // map[a:0 b:0 c:0]
```

### FromEntries

```go
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V
```

从键值对列表创建 Map。

**示例:**

```go
entries := []maputil.Entry[string, int]{
    {Key: "a", Value: 1},
    {Key: "b", Value: 2},
    {Key: "c", Value: 3},
}
result := maputil.FromEntries(entries)
fmt.Println(result)  // map[a:1 b:2 c:3]
```

### ToEntries

```go
func ToEntries[K comparable, V any](m map[K]V) []Entry[K, V]
```

将 Map 转换为键值对列表。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2}
entries := maputil.ToEntries(m)
fmt.Println(entries)
// [{a 1} {b 2}] (顺序可能不同)
```

### Equal

```go
func Equal[K comparable, V comparable](m1, m2 map[K]V) bool
```

检查两个 Map 是否相等。

**示例:**

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"a": 1, "b": 2}
m3 := map[string]int{"a": 1, "b": 3}
fmt.Println(maputil.Equal(m1, m2))  // true
fmt.Println(maputil.Equal(m1, m3))  // false
```

### Size

```go
func Size[K comparable, V any](m map[K]V) int
```

返回 Map 的大小。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
fmt.Println(maputil.Size(m))  // 3
```

### Copy

```go
func Copy[K comparable, V any](m map[K]V) map[K]V
```

复制 Map。

**示例:**

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := maputil.Copy(m1)
m2["c"] = 3
fmt.Println(m1)  // map[a:1 b:2]
fmt.Println(m2)  // map[a:1 b:2 c:3]
```
