---
title: API 概览
---

# API 概览

GoHutool 是一个 Go 语言工具库，提供了丰富的实用工具函数，涵盖字符串处理、数值计算、集合操作、数组操作、日期处理、类型转换和数据验证等功能。

## 模块列表

| 模块 | 包名 | 描述 |
|------|------|------|
| [字符串工具](/api/strutil) | `strutil` | 字符串处理工具，包含判空、截取、填充、转换等函数 |
| [数值工具](/api/numutil) | `numutil` | 数值计算工具，包含解析、数学运算、精确计算等函数 |
| [集合工具](/api/collutil) | `collutil` | 集合操作工具，包含过滤、映射、分组、排序等函数 |
| [Map工具](/api/maputil) | `maputil` | Map操作工具，包含键值操作、过滤、转换、合并等函数 |
| [数组工具](/api/arrayutil) | `arrayutil` | 数组操作工具，包含查找、排序、分块、窗口等函数 |
| [日期工具](/api/dateutil) | `dateutil` | 日期时间工具，包含解析、格式化、计算、判断等函数 |
| [转换工具](/api/convert) | `convert` | 类型转换工具，包含类型转换、编码解码、进制转换等函数 |
| [验证工具](/api/validate) | `validate` | 数据验证工具，包含格式验证、范围检查、类型判断等函数 |

## 导入方式

```go
import "github.com/gongm/gohutool/{package}"
```

例如：

```go
import (
    "github.com/gongm/gohutool/strutil"
    "github.com/gongm/gohutool/numutil"
    "github.com/gongm/gohutool/collutil"
    "github.com/gongm/gohutool/maputil"
    "github.com/gongm/gohutool/arrayutil"
    "github.com/gongm/gohutool/dateutil"
    "github.com/gongm/gohutool/convert"
    "github.com/gongm/gohutool/validate"
)
```

## 快速示例

```go
package main

import (
    "fmt"
    "github.com/gongm/gohutool/strutil"
    "github.com/gongm/gohutool/collutil"
    "github.com/gongm/gohutool/dateutil"
)

func main() {
    // 字符串操作
    fmt.Println(strutil.IsBlank(""))           // true
    fmt.Println(strutil.CamelToUnderline("helloWorld")) // hello_world

    // 集合操作
    nums := []int{1, 2, 3, 4, 5}
    evens := collutil.Filter(nums, func(n int) bool { return n%2 == 0 })
    fmt.Println(evens) // [2, 4]

    // 日期操作
    now := dateutil.Now()
    fmt.Println(dateutil.FormatDateTime(now)) // 2026-08-14 12:00:00
}
```
