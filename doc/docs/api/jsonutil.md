---
title: jsonutil
---

# jsonutil

JSON 工具包，提供 JSON 序列化、反序列化、路径访问、合并、格式化等功能，类似于 Java Hutool 的 JSONUtil。

## 导入

```go
import "github.com/camark/GoHutools/jsonutil"
```

## 函数列表

### Marshal

```go
func Marshal(v interface{}) ([]byte, error)
```

将值序列化为 JSON 字节切片。

**示例:**

```go
data, err := jsonutil.Marshal(map[string]string{"name": "John"})
fmt.Println(string(data)) // {"name":"John"}
```

### MarshalIndent

```go
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```

将值序列化为带缩进的 JSON 字节切片。

**示例:**

```go
data, err := jsonutil.MarshalIndent(map[string]string{"name": "John"}, "", "  ")
fmt.Println(string(data))
```

### MarshalToString

```go
func MarshalToString(v interface{}) (string, error)
```

将值序列化为 JSON 字符串。

**示例:**

```go
s, err := jsonutil.MarshalToString(map[string]string{"name": "John"})
```

### Unmarshal

```go
func Unmarshal(data []byte, v interface{}) error
```

将 JSON 字节反序列化到值中。

**示例:**

```go
var m map[string]string
err := jsonutil.Unmarshal([]byte(`{"name":"John"}`), &m)
```

### UnmarshalFromString

```go
func UnmarshalFromString(s string, v interface{}) error
```

将 JSON 字符串反序列化到值中。

**示例:**

```go
var m map[string]string
err := jsonutil.UnmarshalFromString(`{"name":"John"}`, &m)
```

### UnmarshalFromReader

```go
func UnmarshalFromReader(r io.Reader, v interface{}) error
```

从 Reader 读取 JSON 并反序列化。

**示例:**

```go
var m map[string]string
err := jsonutil.UnmarshalFromReader(strings.NewReader(`{"name":"John"}`), &m)
```

### IsValidJSON

```go
func IsValidJSON(data []byte) bool
```

检查字节数据是否为有效的 JSON。

**示例:**

```go
if jsonutil.IsValidJSON(data) {
    fmt.Println("有效的 JSON")
}
```

### IsValidJSONString

```go
func IsValidJSONString(s string) bool
```

检查字符串是否为有效的 JSON。

**示例:**

```go
if jsonutil.IsValidJSONString(`{"name":"John"}`) {
    fmt.Println("有效的 JSON")
}
```

### ToJSON

```go
func ToJSON(v interface{}) string
```

将值转换为 JSON 字符串，失败时返回空字符串。

**示例:**

```go
s := jsonutil.ToJSON(map[string]string{"name": "John"})
```

### ToPrettyJSON

```go
func ToPrettyJSON(v interface{}) string
```

将值转换为格式化的 JSON 字符串。

**示例:**

```go
s := jsonutil.ToPrettyJSON(map[string]string{"name": "John"})
fmt.Println(s)
```

### FromJSON

```go
func FromJSON(s string, v interface{}) error
```

将 JSON 字符串反序列化到值中。

**示例:**

```go
var m map[string]string
err := jsonutil.FromJSON(`{"name":"John"}`, &m)
```

### ToMap

```go
func ToMap(s string) (map[string]interface{}, error)
```

将 JSON 字符串转换为 map。

**示例:**

```go
m, err := jsonutil.ToMap(`{"name":"John","age":30}`)
fmt.Println(m["name"]) // John
```

### ToSlice

```go
func ToSlice(s string) ([]interface{}, error)
```

将 JSON 字符串转换为切片。

**示例:**

```go
s, err := jsonutil.ToSlice(`[1,2,3]`)
fmt.Println(s) // [1 2 3]
```

### Get

```go
func Get(jsonStr string, path string) (interface{}, error)
```

通过路径获取 JSON 中的值。路径使用点号分隔（如 `a.b.c`），支持数组索引（如 `a.0.b`）。

**示例:**

```go
val, err := jsonutil.Get(`{"user":{"name":"John","age":30}}`, "user.name")
fmt.Println(val) // John
```

### GetString

```go
func GetString(jsonStr string, path string) (string, error)
```

通过路径获取字符串值。

**示例:**

```go
name, err := jsonutil.GetString(`{"user":{"name":"John"}}`, "user.name")
fmt.Println(name) // John
```

### GetInt

```go
func GetInt(jsonStr string, path string) (int, error)
```

通过路径获取整数值。

**示例:**

```go
age, err := jsonutil.GetInt(`{"user":{"age":30}}`, "user.age")
fmt.Println(age) // 30
```

### GetFloat

```go
func GetFloat(jsonStr string, path string) (float64, error)
```

通过路径获取浮点值。

**示例:**

```go
price, err := jsonutil.GetFloat(`{"product":{"price":99.9}}`, "product.price")
```

### GetBool

```go
func GetBool(jsonStr string, path string) (bool, error)
```

通过路径获取布尔值。

**示例:**

```go
active, err := jsonutil.GetBool(`{"user":{"active":true}}`, "user.active")
```

### Set

```go
func Set(jsonStr string, path string, value interface{}) (string, error)
```

通过路径设置 JSON 中的值。

**示例:**

```go
result, err := jsonutil.Set(`{"user":{"name":"John"}}`, "user.name", "Jane")
fmt.Println(result) // {"user":{"name":"Jane"}}
```

### Delete

```go
func Delete(jsonStr string, path string) (string, error)
```

通过路径删除 JSON 中的值。

**示例:**

```go
result, err := jsonutil.Delete(`{"user":{"name":"John","age":30}}`, "user.age")
fmt.Println(result) // {"user":{"name":"John"}}
```

### Merge

```go
func Merge(json1, json2 string) (string, error)
```

合并两个 JSON 对象。如果存在相同键且都是对象则递归合并，否则后者覆盖前者。

**示例:**

```go
merged, err := jsonutil.Merge(
    `{"name":"John","age":30}`,
    `{"age":31,"email":"john@example.com"}`,
)
fmt.Println(merged) // {"age":31,"email":"john@example.com","name":"John"}
```

### Format

```go
func Format(jsonStr string) (string, error)
```

格式化 JSON 字符串（带缩进）。

**示例:**

```go
formatted, err := jsonutil.Format(`{"name":"John","age":30}`)
fmt.Println(formatted)
```

### Compact

```go
func Compact(jsonStr string) (string, error)
```

压缩 JSON 字符串（去除空白）。

**示例:**

```go
compacted, err := jsonutil.Compact(`{
    "name": "John",
    "age": 30
}`)
```

### Keys

```go
func Keys(jsonStr string) ([]string, error)
```

获取 JSON 对象的所有顶层键。

**示例:**

```go
keys, err := jsonutil.Keys(`{"name":"John","age":30}`)
fmt.Println(keys) // [name age]
```

### Values

```go
func Values(jsonStr string) ([]interface{}, error)
```

获取 JSON 对象的所有顶层值。

**示例:**

```go
values, err := jsonutil.Values(`{"name":"John","age":30}`)
fmt.Println(values) // [John 30]
```

### Contains

```go
func Contains(jsonStr string, path string) bool
```

检查 JSON 中是否包含指定路径的值。

**示例:**

```go
if jsonutil.Contains(`{"user":{"name":"John"}}`, "user.name") {
    fmt.Println("存在该字段")
}
```

### Wrap

```go
func Wrap(key string, value interface{}) (string, error)
```

将值包装为 JSON 对象。

**示例:**

```go
s, err := jsonutil.Wrap("name", "John")
fmt.Println(s) // {"name":"John"}
```

### WrapMap

```go
func WrapMap(m map[string]interface{}) (string, error)
```

将 map 转换为 JSON 字符串。

**示例:**

```go
s, err := jsonutil.WrapMap(map[string]interface{}{
    "name": "John",
    "age":  30,
})
```

## 完整示例

```go
package main

import (
    "fmt"
    "log"

    "github.com/camark/GoHutools/jsonutil"
)

func main() {
    jsonStr := `{
        "user": {
            "name": "John",
            "age": 30,
            "hobbies": ["reading", "coding"]
        }
    }`

    // 获取嵌套值
    name, err := jsonutil.GetString(jsonStr, "user.name")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Name:", name)

    // 获取数组元素
    hobby, err := jsonutil.GetString(jsonStr, "user.hobbies.0")
    fmt.Println("First hobby:", hobby)

    // 设置值
    newJSON, err := jsonutil.Set(jsonStr, "user.name", "Jane")
    fmt.Println(newJSON)

    // 合并
    merged, err := jsonutil.Merge(
        `{"a":1,"b":2}`,
        `{"b":3,"c":4}`,
    )
    fmt.Println(merged) // {"a":1,"b":3,"c":4}

    // 格式化
    formatted, _ := jsonutil.Format(`{"name":"John","age":30}`)
    fmt.Println(formatted)
}
```
