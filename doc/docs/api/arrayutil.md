---
title: arrayutil - 数组工具
---

# arrayutil - 数组工具

数组操作工具包，提供了丰富的数组操作函数，包括查找、排序、分块、窗口、压缩等功能。使用 Go 泛型实现，类型安全。

## 导入

```go
import "github.com/camark/GoHutools/arrayutil"
```

## 类型定义

### Pair

```go
type Pair[T, U any] struct {
    First  T
    Second U
}
```

表示值对的结构体。

## 函数列表

### IsEmpty

```go
func IsEmpty[T any](arr []T) bool
```

检查数组是否为空。

**示例:**

```go
fmt.Println(arrayutil.IsEmpty([]int{}))        // true
fmt.Println(arrayutil.IsEmpty([]int{1, 2, 3})) // false
```

### IsNotEmpty

```go
func IsNotEmpty[T any](arr []T) bool
```

检查数组是否不为空。

**示例:**

```go
fmt.Println(arrayutil.IsNotEmpty([]int{1, 2, 3})) // true
```

### Contains

```go
func Contains[T comparable](arr []T, elem T) bool
```

检查数组是否包含指定元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(arrayutil.Contains(nums, 3))  // true
fmt.Println(arrayutil.Contains(nums, 6))  // false
```

### IndexOf

```go
func IndexOf[T comparable](arr []T, elem T) int
```

返回元素的索引（不存在返回 -1）。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(arrayutil.IndexOf(nums, 3))  // 2
fmt.Println(arrayutil.IndexOf(nums, 6))  // -1
```

### LastIndexOf

```go
func LastIndexOf[T comparable](arr []T, elem T) int
```

返回元素最后出现的索引（不存在返回 -1）。

**示例:**

```go
nums := []int{1, 2, 3, 2, 1}
fmt.Println(arrayutil.LastIndexOf(nums, 2))  // 3
fmt.Println(arrayutil.LastIndexOf(nums, 6))  // -1
```

### Fill

```go
func Fill[T any](arr []T, val T) []T
```

用指定值填充数组。

**示例:**

```go
arr := make([]int, 5)
result := arrayutil.Fill(arr, 7)
fmt.Println(result)  // [7, 7, 7, 7, 7]
```

### CopyOf

```go
func CopyOf[T any](arr []T, newLength int) []T
```

复制数组到指定长度。如果 newLength 大于原数组长度，用零值填充；如果小于，截断。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := arrayutil.CopyOf(nums, 3)
fmt.Println(result)  // [1, 2, 3]

result = arrayutil.CopyOf(nums, 7)
fmt.Println(result)  // [1, 2, 3, 4, 5, 0, 0]
```

### CopyOfRange

```go
func CopyOfRange[T any](arr []T, from, to int) []T
```

复制数组的指定范围 [from, to)。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := arrayutil.CopyOfRange(nums, 1, 4)
fmt.Println(result)  // [2, 3, 4]
```

### Insert

```go
func Insert[T any](arr []T, index int, elem T) []T
```

在指定位置插入元素，返回新数组。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := arrayutil.Insert(nums, 2, 10)
fmt.Println(result)  // [1, 2, 10, 3, 4, 5]
```

### RemoveAt

```go
func RemoveAt[T any](arr []T, index int) []T
```

移除指定位置的元素，返回新数组。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := arrayutil.RemoveAt(nums, 2)
fmt.Println(result)  // [1, 2, 4, 5]
```

### Swap

```go
func Swap[T any](arr []T, i, j int)
```

交换数组中两个位置的元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
arrayutil.Swap(nums, 1, 3)
fmt.Println(nums)  // [1, 4, 3, 2, 5]
```

### Reverse

```go
func Reverse[T any](arr []T) []T
```

反转数组（原地修改）。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
arrayutil.Reverse(nums)
fmt.Println(nums)  // [5, 4, 3, 2, 1]
```

### Rotate

```go
func Rotate[T any](arr []T, n int) []T
```

旋转数组 n 个位置（正数右旋，负数左旋），返回新数组。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := arrayutil.Rotate(nums, 2)
fmt.Println(result)  // [4, 5, 1, 2, 3]

result = arrayutil.Rotate(nums, -2)
fmt.Println(result)  // [3, 4, 5, 1, 2]
```

### Unique

```go
func Unique[T comparable](arr []T) []T
```

去除重复元素，保留首次出现的顺序。

**示例:**

```go
nums := []int{1, 2, 3, 2, 1, 4}
result := arrayutil.Unique(nums)
fmt.Println(result)  // [1, 2, 3, 4]
```

### ToInterfaceSlice

```go
func ToInterfaceSlice[T any](arr []T) []interface{}
```

将类型化切片转换为 interface{} 切片。

**示例:**

```go
nums := []int{1, 2, 3}
result := arrayutil.ToInterfaceSlice(nums)
fmt.Println(result)  // [1, 2, 3]
```

### FromInterfaceSlice

```go
func FromInterfaceSlice[T any](arr []interface{}) ([]T, bool)
```

将 interface{} 切片转换为类型化切片。

**示例:**

```go
arr := []interface{}{1, 2, 3}
result, ok := arrayutil.FromInterfaceSlice[int](arr)
fmt.Println(result, ok)  // [1, 2, 3] true
```

### Filter

```go
func Filter[T any](arr []T, predicate func(T) bool) []T
```

过滤数组，返回满足条件的元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5, 6}
evens := arrayutil.Filter(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(evens)  // [2, 4, 6]
```

### Map

```go
func Map[T, R any](arr []T, mapper func(T) R) []R
```

映射数组元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
doubled := arrayutil.Map(nums, func(n int) int {
    return n * 2
})
fmt.Println(doubled)  // [2, 4, 6, 8, 10]
```

### Any

```go
func Any[T any](arr []T, predicate func(T) bool) bool
```

检查是否有任意元素满足条件。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
hasEven := arrayutil.Any(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(hasEven)  // true
```

### All

```go
func All[T any](arr []T, predicate func(T) bool) bool
```

检查是否所有元素都满足条件。

**示例:**

```go
nums := []int{2, 4, 6, 8}
allEven := arrayutil.All(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(allEven)  // true
```

### None

```go
func None[T any](arr []T, predicate func(T) bool) bool
```

检查是否没有元素满足条件。

**示例:**

```go
nums := []int{1, 3, 5, 7}
noneEven := arrayutil.None(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(noneEven)  // true
```

### Join

```go
func Join[T any](arr []T, separator string) string
```

使用分隔符连接数组元素。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := arrayutil.Join(nums, ", ")
fmt.Println(result)  // "1, 2, 3, 4, 5"
```

### Equal

```go
func Equal[T comparable](a, b []T) bool
```

检查两个数组是否相等。

**示例:**

```go
a := []int{1, 2, 3}
b := []int{1, 2, 3}
c := []int{1, 2, 4}
fmt.Println(arrayutil.Equal(a, b))  // true
fmt.Println(arrayutil.Equal(a, c))  // false
```

### Concat

```go
func Concat[T any](arrays ...[]T) []T
```

连接多个数组。

**示例:**

```go
a := []int{1, 2}
b := []int{3, 4}
c := []int{5, 6}
result := arrayutil.Concat(a, b, c)
fmt.Println(result)  // [1, 2, 3, 4, 5, 6]
```

### First

```go
func First[T any](arr []T) (T, bool)
```

返回第一个元素。

**示例:**

```go
nums := []int{1, 2, 3}
first, ok := arrayutil.First(nums)
fmt.Println(first, ok)  // 1 true
```

### Last

```go
func Last[T any](arr []T) (T, bool)
```

返回最后一个元素。

**示例:**

```go
nums := []int{1, 2, 3}
last, ok := arrayutil.Last(nums)
fmt.Println(last, ok)  // 3 true
```

### Shuffle

```go
func Shuffle[T any](arr []T) []T
```

随机打乱数组，返回新数组。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := arrayutil.Shuffle(nums)
fmt.Println(result)  // 随机顺序，如 [3, 1, 5, 2, 4]
```

### Min

```go
func Min[T cmp.Ordered](arr []T) (T, bool)
```

返回最小元素。

**示例:**

```go
nums := []int{3, 1, 4, 1, 5, 9}
min, ok := arrayutil.Min(nums)
fmt.Println(min, ok)  // 1 true
```

### Max

```go
func Max[T cmp.Ordered](arr []T) (T, bool)
```

返回最大元素。

**示例:**

```go
nums := []int{3, 1, 4, 1, 5, 9}
max, ok := arrayutil.Max(nums)
fmt.Println(max, ok)  // 9 true
```

### Flatten

```go
func Flatten[T any](arr [][]T) []T
```

展平嵌套数组。

**示例:**

```go
nested := [][]int{{1, 2}, {3, 4}, {5}}
result := arrayutil.Flatten(nested)
fmt.Println(result)  // [1, 2, 3, 4, 5]
```

### Chunk

```go
func Chunk[T any](arr []T, size int) [][]T
```

将数组分割为指定大小的块。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5, 6, 7}
result := arrayutil.Chunk(nums, 3)
fmt.Println(result)  // [[1, 2, 3], [4, 5, 6], [7]]
```

### Window

```go
func Window[T any](arr []T, size int) [][]T
```

创建滑动窗口。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
result := arrayutil.Window(nums, 3)
fmt.Println(result)  // [[1, 2, 3], [2, 3, 4], [3, 4, 5]]
```

### Zip

```go
func Zip[T, U any](a []T, b []U) []Pair[T, U]
```

将两个数组合并为值对数组。

**示例:**

```go
names := []string{"Alice", "Bob", "Charlie"}
ages := []int{25, 30, 35}
result := arrayutil.Zip(names, ages)
fmt.Println(result)
// [{Alice 25} {Bob 30} {Charlie 35}]
```

### Unzip

```go
func Unzip[T, U any](pairs []Pair[T, U]) ([]T, []U)
```

将值对数组分离为两个数组。

**示例:**

```go
pairs := []arrayutil.Pair[string, int]{
    {First: "Alice", Second: 25},
    {First: "Bob", Second: 30},
    {First: "Charlie", Second: 35},
}
names, ages := arrayutil.Unzip(pairs)
fmt.Println(names)  // [Alice Bob Charlie]
fmt.Println(ages)   // [25 30 35]
```
