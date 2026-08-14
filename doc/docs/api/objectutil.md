---
title: ObjectUtil
---

# ObjectUtil

对象工具模块，提供对象判空、比较、类型判断、指针操作等通用工具函数。类似 Java Hutool 的 ObjectUtil。

## 导入

```go
import "github.com/gongm/gohutool/objectutil"
```

## 空值判断

### IsNil

```go
func IsNil(v interface{}) bool
```

判断值是否为 nil。支持指针、map、slice、chan、func、interface 等类型的 nil 判断。

**示例:**

```go
var p *int
fmt.Println(objectutil.IsNil(p))    // true
fmt.Println(objectutil.IsNil(nil))  // true
fmt.Println(objectutil.IsNil(0))    // false

m := map[string]int{}
fmt.Println(objectutil.IsNil(m))    // false
```

### IsNotNull

```go
func IsNotNull(v interface{}) bool
```

判断值是否不为 nil。

### IsEmpty

```go
func IsEmpty(v interface{}) bool
```

判断值是否为空。空的定义：
- nil
- 空字符串 `""`
- 空 slice、map、array
- `false`
- 数值 `0`

**示例:**

```go
fmt.Println(objectutil.IsEmpty(""))         // true
fmt.Println(objectutil.IsEmpty("hello"))    // false
fmt.Println(objectutil.IsEmpty(0))          // true
fmt.Println(objectutil.IsEmpty(1))          // false
fmt.Println(objectutil.IsEmpty(false))      // true
fmt.Println(objectutil.IsEmpty([]int{}))    // true
fmt.Println(objectutil.IsEmpty(nil))        // true
```

### IsNotEmpty

```go
func IsNotEmpty(v interface{}) bool
```

判断值是否不为空。

## 默认值

### DefaultIfNil

```go
func DefaultIfNil(v interface{}, defaultVal interface{}) interface{}
```

如果值为 nil，返回默认值。

**示例:**

```go
var name *string
result := objectutil.DefaultIfNil(name, "匿名")
fmt.Println(result) // 匿名

val := "张三"
result = objectutil.DefaultIfNil(val, "匿名")
fmt.Println(result) // 张三
```

### DefaultIfEmpty

```go
func DefaultIfEmpty(v interface{}, defaultVal interface{}) interface{}
```

如果值为空，返回默认值。

**示例:**

```go
result := objectutil.DefaultIfEmpty("", "默认值")
fmt.Println(result) // 默认值

result = objectutil.DefaultIfEmpty(0, 100)
fmt.Println(result) // 100
```

## 比较

### Equals

```go
func Equals(a, b interface{}) bool
```

判断两个值是否相等（使用 reflect.DeepEqual）。

**示例:**

```go
fmt.Println(objectutil.Equals(1, 1))           // true
fmt.Println(objectutil.Equals("a", "b"))       // false
fmt.Println(objectutil.Equals(nil, nil))        // true
fmt.Println(objectutil.Equals([]int{1}, []int{1})) // true
```

### NotEquals

```go
func NotEquals(a, b interface{}) bool
```

判断两个值是否不相等。

### Compare

```go
func Compare(a, b interface{}) int
```

比较两个值。返回 `-1`（a < b）、`0`（a == b）、`1`（a > b）。

**示例:**

```go
fmt.Println(objectutil.Compare(1, 2))     // -1
fmt.Println(objectutil.Compare(2, 2))     // 0
fmt.Println(objectutil.Compare(3, 2))     // 1
fmt.Println(objectutil.Compare("a", "b")) // -1
```

## 类型信息

### TypeName

```go
func TypeName(v interface{}) string
```

返回值的类型名称。

**示例:**

```go
fmt.Println(objectutil.TypeName(42))        // int
fmt.Println(objectutil.TypeName("hello"))   // string
fmt.Println(objectutil.TypeName(nil))       // nil
fmt.Println(objectutil.TypeName([]int{}))   // []int
```

### TypeOf

```go
func TypeOf(v interface{}) reflect.Type
```

返回值的 reflect.Type。

### KindOf

```go
func KindOf(v interface{}) reflect.Kind
```

返回值的 reflect.Kind。

**示例:**

```go
fmt.Println(objectutil.KindOf(42))       // int
fmt.Println(objectutil.KindOf("hello"))  // string
fmt.Println(objectutil.KindOf([]int{}))  // slice
```

### IsType

```go
func IsType(v interface{}, typeName string) bool
```

检查值是否为指定类型名称。

**示例:**

```go
fmt.Println(objectutil.IsType(42, "int"))      // true
fmt.Println(objectutil.IsType("hi", "string")) // true
```

### IsKind

```go
func IsKind(v interface{}, kind reflect.Kind) bool
```

检查值是否为指定 Kind。

**示例:**

```go
fmt.Println(objectutil.IsKind(42, reflect.Int))         // true
fmt.Println(objectutil.IsKind("hi", reflect.String))    // true
```

## 指针操作

### ToPtr

```go
func ToPtr(v interface{}) interface{}
```

返回值的指针。

**示例:**

```go
ptr := objectutil.ToPtr(42)
fmt.Println(reflect.TypeOf(ptr)) // *int
```

### FromPtr

```go
func FromPtr(v interface{}) interface{}
```

解引用指针。如果是指针类型，返回指向的值；否则返回原值。

**示例:**

```go
n := 42
ptr := &n
val := objectutil.FromPtr(ptr)
fmt.Println(val) // 42

str := "hello"
val = objectutil.FromPtr(str)
fmt.Println(val) // hello
```

## 复制与转换

### Clone

```go
func Clone(v interface{}) interface{}
```

深拷贝对象。通过反射创建新的副本。

**示例:**

```go
original := &MyStruct{Name: "test"}
copied := objectutil.Clone(original).(*MyStruct)
copied.Name = "modified"
fmt.Println(original.Name) // test（不受影响）
```

### String

```go
func String(v interface{}) string
```

返回值的字符串表示。nil 返回空字符串。

**示例:**

```go
fmt.Println(objectutil.String(42))       // 42
fmt.Println(objectutil.String(true))     // true
fmt.Println(objectutil.String(nil))      //
```

### Copy

```go
func Copy(v interface{}) interface{}
```

复制值。对于引用类型，返回相同的引用。

## 完整示例

```go
package main

import (
    "fmt"
    "reflect"

    "github.com/gongm/gohutool/objectutil"
)

func main() {
    // 空值判断
    var s *string
    fmt.Println("nil 检查:", objectutil.IsNil(s))         // true
    fmt.Println("空检查:", objectutil.IsEmpty(""))        // true
    fmt.Println("非空检查:", objectutil.IsNotEmpty("hi")) // true

    // 默认值
    name := objectutil.DefaultIfNil(s, "默认名称")
    fmt.Println("默认值:", name) // 默认名称

    // 比较
    fmt.Println("相等:", objectutil.Equals(1, 1))       // true
    fmt.Println("比较:", objectutil.Compare("a", "b")) // -1

    // 类型信息
    fmt.Println("类型:", objectutil.TypeName(42))       // int
    fmt.Println("Kind:", objectutil.KindOf([]int{}))    // slice
    fmt.Println("是int:", objectutil.IsKind(42, reflect.Int)) // true

    // 指针操作
    ptr := objectutil.ToPtr(100)
    fmt.Println("指针:", objectutil.FromPtr(ptr)) // 100
}
```
