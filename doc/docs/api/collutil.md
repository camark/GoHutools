---
title: collutil - 集合工具
---

# collutil - 集合工具

集合操作工具包，提供了丰富的切片操作函数，包括判空、查找、过滤、映射、排序、分组等功能。使用 Go 泛型实现，类型安全。

## 导入

```go
import "github.com/gongm/gohutool/collutil"
```

## 函数列表

### IsEmpty

```go
func IsEmpty[T any](coll []T) bool
```

检查切片是否为空。

**示例:**

```go
fmt.Println(collutil.IsEmpty([]int{}))        // true
fmt.Println(collutil.IsEmpty([]int{1, 2, 3})) // false
```

### IsNotEmpty

```go
func IsNotEmpty[T any](coll []T) bool
```

检查切片是否不为空。

**示例:**

```go
fmt.Println(collutil.IsNotEmpty([]int{1, 2, 3})) // true
```

### Contains

```go
func Contains[T comparable](coll []T, item T) bool
```

检查切片是否包含指定元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(collutil.Contains(nums, 3))  // true
fmt.Println(collutil.Contains(nums, 6))  // false
```

### ContainsAll

```go
func ContainsAll[T comparable](coll []T, items ...T) bool
```

检查切片是否包含所有指定元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(collutil.ContainsAll(nums, 1, 3, 5))  // true
fmt.Println(collutil.ContainsAll(nums, 1, 3, 6))  // false
```

### ContainsAny

```go
func ContainsAny[T comparable](coll []T, items ...T) bool
```

检查切片是否包含任意一个指定元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(collutil.ContainsAny(nums, 3, 6, 7))  // true
fmt.Println(collutil.ContainsAny(nums, 6, 7, 8))  // false
```

### AddIfAbsent

```go
func AddIfAbsent[T comparable](coll []T, item T) []T
```

如果元素不存在则添加。

**示例:**

```go
nums := []int{1, 2, 3}
nums = collutil.AddIfAbsent(nums, 4)  // [1, 2, 3, 4]
nums = collutil.AddIfAbsent(nums, 3)  // [1, 2, 3, 4]
```

### Remove

```go
func Remove[T comparable](coll []T, item T) []T
```

移除第一个匹配的元素。

**示例:**

```go
nums := []int{1, 2, 3, 2, 4}
result := collutil.Remove(nums, 2)
fmt.Println(result)  // [1, 3, 2, 4]
```

### RemoveAll

```go
func RemoveAll[T comparable](coll []T, item T) []T
```

移除所有匹配的元素。

**示例:**

```go
nums := []int{1, 2, 3, 2, 4}
result := collutil.RemoveAll(nums, 2)
fmt.Println(result)  // [1, 3, 4]
```

### Distinct

```go
func Distinct[T comparable](coll []T) []T
```

去除重复元素，保留顺序。

**示例:**

```go
nums := []int{1, 2, 3, 2, 1}
result := collutil.Distinct(nums)
fmt.Println(result)  // [1, 2, 3]
```

### Reverse

```go
func Reverse[T any](coll []T) []T
```

反转切片（返回新切片）。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := collutil.Reverse(nums)
fmt.Println(result)  // [5, 4, 3, 2, 1]
```

### Sub

```go
func Sub[T any](coll []T, start, end int) []T
```

提取子切片。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := collutil.Sub(nums, 1, 3)
fmt.Println(result)  // [2, 3]
```

### Split

```go
func Split[T any](coll []T, size int) [][]T
```

将切片分割为指定大小的块。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5, 6, 7}
result := collutil.Split(nums, 3)
fmt.Println(result)  // [[1, 2, 3], [4, 5, 6], [7]]
```

### Filter

```go
func Filter[T any](coll []T, predicate func(T) bool) []T
```

过滤切片，返回满足条件的元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5, 6}
evens := collutil.Filter(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(evens)  // [2, 4, 6]
```

### Map

```go
func Map[T, R any](coll []T, mapper func(T) R) []R
```

映射切片元素，返回新切片。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
doubled := collutil.Map(nums, func(n int) int {
    return n * 2
})
fmt.Println(doubled)  // [2, 4, 6, 8, 10]
```

### ForEach

```go
func ForEach[T any](coll []T, consumer func(T))
```

遍历切片，对每个元素执行操作。

**示例:**

```go
nums := []int{1, 2, 3}
collutil.ForEach(nums, func(n int) {
    fmt.Println(n)
})
// 输出: 1 2 3
```

### ForEachIndexed

```go
func ForEachIndexed[T any](coll []T, consumer func(int, T))
```

遍历切片，带索引。

**示例:**

```go
nums := []int{10, 20, 30}
collutil.ForEachIndexed(nums, func(i, n int) {
    fmt.Printf("Index %d: %d\n", i, n)
})
// 输出: Index 0: 10, Index 1: 20, Index 2: 30
```

### GroupBy

```go
func GroupBy[T any, K comparable](coll []T, keyFunc func(T) K) map[K][]T
```

按条件分组。

**示例:**

```go
words := []string{"apple", "banana", "avocado", "blueberry", "cherry"}
groups := collutil.GroupBy(words, func(s string) string {
    return string(s[0])
})
fmt.Println(groups)
// map[a:[apple avocado] b:[banana blueberry] c:[cherry]]
```

### ToMap

```go
func ToMap[T any, K comparable, V any](coll []T, keyFunc func(T) K, valFunc func(T) V) map[K]V
```

将切片转换为 Map。

**示例:**

```go
words := []string{"apple", "banana", "cherry"}
wordMap := collutil.ToMap(words, func(s string) string {
    return s
}, func(s string) int {
    return len(s)
})
fmt.Println(wordMap)  // map[apple:5 banana:6 cherry:6]
```

### Sort

```go
func Sort[T cmp.Ordered](coll []T) []T
```

排序切片（返回新切片）。

**示例:**

```go
nums := []int{3, 1, 4, 1, 5, 9}
result := collutil.Sort(nums)
fmt.Println(result)  // [1, 1, 3, 4, 5, 9]
```

### SortBy

```go
func SortBy[T any](coll []T, less func(T, T) bool) []T
```

自定义比较器排序。

**示例:**

```go
words := []string{"banana", "apple", "cherry"}
result := collutil.SortBy(words, func(a, b string) bool {
    return len(a) < len(b)
})
fmt.Println(result)  // [apple, banana, cherry]
```

### Min

```go
func Min[T cmp.Ordered](coll []T) (T, bool)
```

返回最小元素。

**示例:**

```go
nums := []int{3, 1, 4, 1, 5, 9}
min, ok := collutil.Min(nums)
fmt.Println(min, ok)  // 1 true
```

### Max

```go
func Max[T cmp.Ordered](coll []T) (T, bool)
```

返回最大元素。

**示例:**

```go
nums := []int{3, 1, 4, 1, 5, 9}
max, ok := collutil.Max(nums)
fmt.Println(max, ok)  // 9 true
```

### Sum

```go
func Sum[T cmp.Ordered](coll []T) T
```

返回切片元素之和。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(collutil.Sum(nums))  // 15
```

### Flatten

```go
func Flatten[T any](colls [][]T) []T
```

展平嵌套切片。

**示例:**

```go
nested := [][]int{{1, 2}, {3, 4}, {5}}
result := collutil.Flatten(nested)
fmt.Println(result)  // [1, 2, 3, 4, 5]
```

### Chunk

```go
func Chunk[T any](coll []T, size int) [][]T
```

将切片分割为指定大小的块（Split 的别名）。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := collutil.Chunk(nums, 2)
fmt.Println(result)  // [[1, 2], [3, 4], [5]]
```

### Zip

```go
func Zip[T comparable, U any](keys []T, values []U) map[T]U
```

将两个切片合并为 Map。

**示例:**

```go
keys := []string{"a", "b", "c"}
values := []int{1, 2, 3}
result := collutil.Zip(keys, values)
fmt.Println(result)  // map[a:1 b:2 c:3]
```

### Keys

```go
func Keys[K comparable, V any](m map[K]V) []K
```

返回 Map 的所有键。

**示例:**

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
keys := collutil.Keys(m)
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
values := collutil.Values(m)
fmt.Println(values)  // [1 2 3] (顺序可能不同)
```

### First

```go
func First[T any](coll []T) (T, bool)
```

返回第一个元素。

**示例:**

```go
nums := []int{1, 2, 3}
first, ok := collutil.First(nums)
fmt.Println(first, ok)  // 1 true
```

### Last

```go
func Last[T any](coll []T) (T, bool)
```

返回最后一个元素。

**示例:**

```go
nums := []int{1, 2, 3}
last, ok := collutil.Last(nums)
fmt.Println(last, ok)  // 3 true
```

### Find

```go
func Find[T any](coll []T, predicate func(T) bool) (T, bool)
```

查找第一个满足条件的元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
found, ok := collutil.Find(nums, func(n int) bool {
    return n > 3
})
fmt.Println(found, ok)  // 4 true
```

### FindIndex

```go
func FindIndex[T any](coll []T, predicate func(T) bool) int
```

查找第一个满足条件的元素的索引。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
index := collutil.FindIndex(nums, func(n int) bool {
    return n > 3
})
fmt.Println(index)  // 3
```

### AnyMatch

```go
func AnyMatch[T any](coll []T, predicate func(T) bool) bool
```

检查是否有任意元素满足条件。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
hasEven := collutil.AnyMatch(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(hasEven)  // true
```

### AllMatch

```go
func AllMatch[T any](coll []T, predicate func(T) bool) bool
```

检查是否所有元素都满足条件。

**示例:**

```go
nums := []int{2, 4, 6, 8}
allEven := collutil.AllMatch(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(allEven)  // true
```

### NoneMatch

```go
func NoneMatch[T any](coll []T, predicate func(T) bool) bool
```

检查是否没有元素满足条件。

**示例:**

```go
nums := []int{1, 3, 5, 7}
noneEven := collutil.NoneMatch(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(noneEven)  // true
```

### Reduce

```go
func Reduce[T, R any](coll []T, initial R, reducer func(R, T) R) R
```

归约操作，将切片归约为单个值。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
sum := collutil.Reduce(nums, 0, func(acc, n int) int {
    return acc + n
})
fmt.Println(sum)  // 15
```

### Count

```go
func Count[T any](coll []T, predicate func(T) bool) int
```

统计满足条件的元素数量。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5, 6}
evenCount := collutil.Count(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(evenCount)  // 3
```

### Equal

```go
func Equal[T comparable](a, b []T) bool
```

检查两个切片是否相等。

**示例:**

```go
a := []int{1, 2, 3}
b := []int{1, 2, 3}
c := []int{1, 2, 4}
fmt.Println(collutil.Equal(a, b))  // true
fmt.Println(collutil.Equal(a, c))  // false
```

### Union

```go
func Union[T comparable](a, b []T) []T
```

返回两个切片的并集。

**示例:**

```go
a := []int{1, 2, 3}
b := []int{3, 4, 5}
result := collutil.Union(a, b)
fmt.Println(result)  // [1, 2, 3, 4, 5]
```

### Intersection

```go
func Intersection[T comparable](a, b []T) []T
```

返回两个切片的交集。

**示例:**

```go
a := []int{1, 2, 3, 4}
b := []int{3, 4, 5, 6}
result := collutil.Intersection(a, b)
fmt.Println(result)  // [3, 4]
```

### Difference

```go
func Difference[T comparable](a, b []T) []T
```

返回在 a 中但不在 b 中的元素。

**示例:**

```go
a := []int{1, 2, 3, 4}
b := []int{3, 4, 5, 6}
result := collutil.Difference(a, b)
fmt.Println(result)  // [1, 2]
```

### ToSlice

```go
func ToSlice[T any](items ...T) []T
```

将可变参数转换为切片。

**示例:**

```go
result := collutil.ToSlice(1, 2, 3, 4, 5)
fmt.Println(result)  // [1, 2, 3, 4, 5]
```

### Range

```go
func Range(start, end int) []int
```

生成从 start 到 end（不包含）的整数序列。

**示例:**

```go
result := collutil.Range(1, 6)
fmt.Println(result)  // [1, 2, 3, 4, 5]
```

### RangeWithStep

```go
func RangeWithStep(start, end, step int) []int
```

生成带步长的整数序列。

**示例:**

```go
result := collutil.RangeWithStep(0, 10, 2)
fmt.Println(result)  // [0, 2, 4, 6, 8]

result = collutil.RangeWithStep(10, 0, -3)
fmt.Println(result)  // [10, 7, 4, 1]
```
