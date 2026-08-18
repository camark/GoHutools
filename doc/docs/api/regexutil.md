---
title: RegexUtil
---

# RegexUtil

正则表达式工具模块，提供正则匹配、提取、替换等便捷函数，并内置常用正则模式和编译缓存。类似 Java Hutool 的 ReUtil。

## 导入

```go
import "github.com/camark/GoHutools/regexutil"
```

## 正则匹配

### IsMatch

```go
func IsMatch(s string, pattern string) bool
```

检查字符串是否匹配正则模式。

**示例:**

```go
fmt.Println(regexutil.IsMatch("test@email.com", `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)) // true
fmt.Println(regexutil.IsMatch("123", `^\d+$`)) // true
```

### IsMatchBytes

```go
func IsMatchBytes(b []byte, pattern string) bool
```

检查字节切片是否匹配正则模式。

### MatchAll

```go
func MatchAll(s string, pattern string) bool
```

检查整个字符串是否完全匹配正则模式。

### ContainsMatch

```go
func ContainsMatch(s string, pattern string) bool
```

检查字符串中是否包含匹配正则模式的子串。等同于 `IsMatch`。

## 查找

### Find

```go
func Find(s string, pattern string) string
```

查找第一个匹配的子串。

**示例:**

```go
result := regexutil.Find("hello 123 world 456", `\d+`)
fmt.Println(result) // 123
```

### FindBytes

```go
func FindBytes(b []byte, pattern string) []byte
```

查找第一个匹配的字节子串。

### FindAll

```go
func FindAll(s string, pattern string) []string
```

查找所有匹配的子串。

**示例:**

```go
results := regexutil.FindAll("hello 123 world 456", `\d+`)
fmt.Println(results) // [123 456]
```

### FindAllBytes

```go
func FindAllBytes(b []byte, pattern string) [][]byte
```

查找所有匹配的字节子串。

### FindAllWithIndex

```go
func FindAllWithIndex(s string, pattern string) [][]int
```

查找所有匹配的子串及其在原字符串中的位置。

**示例:**

```go
indices := regexutil.FindAllWithIndex("hello 123 world 456", `\d+`)
fmt.Println(indices) // [[6 9] [17 20]]
```

### Count

```go
func Count(s string, pattern string) int
```

统计匹配的次数。

**示例:**

```go
count := regexutil.Count("a1b2c3d4", `\d`)
fmt.Println(count) // 4
```

## 分组提取

### FindGroups

```go
func FindGroups(s string, pattern string) []string
```

查找第一个匹配并返回所有捕获组。第一个元素为完整匹配。

**示例:**

```go
groups := regexutil.FindGroups("2024-01-15", `(\d{4})-(\d{2})-(\d{2})`)
fmt.Println(groups) // [2024-01-15 2024 01 15]
```

### FindAllGroups

```go
func FindAllGroups(s string, pattern string) [][]string
```

查找所有匹配及其捕获组。

**示例:**

```go
groups := regexutil.FindAllGroups("a1b2c3", `(\w)(\d)`)
fmt.Println(groups) // [[a1 a 1] [b2 b 2] [c3 c 3]]
```

### FindNamedGroups

```go
func FindNamedGroups(s string, pattern string) map[string]string
```

查找第一个匹配并返回命名捕获组。

**示例:**

```go
groups := regexutil.FindNamedGroups("2024-01-15", `(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)
fmt.Println(groups) // map[day:15 month:01 year:2024]
```

### ExtractGroup

```go
func ExtractGroup(s string, pattern string, groupIndex int) string
```

提取指定索引的捕获组。

**示例:**

```go
year := regexutil.ExtractGroup("2024-01-15", `(\d{4})-(\d{2})-(\d{2})`, 1)
fmt.Println(year) // 2024

month := regexutil.ExtractGroup("2024-01-15", `(\d{4})-(\d{2})-(\d{2})`, 2)
fmt.Println(month) // 01
```

### ExtractNamedGroup

```go
func ExtractNamedGroup(s string, pattern string, name string) string
```

提取指定名称的捕获组。

**示例:**

```go
year := regexutil.ExtractNamedGroup("2024-01-15", `(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`, "year")
fmt.Println(year) // 2024
```

## 替换

### Replace

```go
func Replace(s string, pattern string, replacement string) string
```

替换所有匹配的子串。

**示例:**

```go
result := regexutil.Replace("hello 123 world 456", `\d+`, "***")
fmt.Println(result) // hello *** world ***
```

### ReplaceBytes

```go
func ReplaceBytes(b []byte, pattern string, replacement []byte) []byte
```

替换字节切片中所有匹配的子串。

### ReplaceFunc

```go
func ReplaceFunc(s string, pattern string, replacer func(string) string) string
```

使用函数替换所有匹配的子串。

**示例:**

```go
result := regexutil.ReplaceFunc("hello 123 world 456", `\d+`, func(match string) string {
    return "[" + match + "]"
})
fmt.Println(result) // hello [123] world [456]
```

### ReplaceFirst

```go
func ReplaceFirst(s string, pattern string, replacement string) string
```

替换第一个匹配的子串。

## 分割

### Split

```go
func Split(s string, pattern string) []string
```

按正则模式分割字符串。

**示例:**

```go
parts := regexutil.Split("one,two;three four", `[,;\s]+`)
fmt.Println(parts) // [one two three four]
```

## 工具函数

### MustGetRegex

```go
func MustGetRegex(pattern string) *regexp.Regexp
```

返回编译后的正则表达式，如果编译失败则 panic。

### Quote

```go
func Quote(s string) string
```

转义字符串中的正则特殊字符。

**示例:**

```go
quoted := regexutil.Price("price: $10.00 (USD)")
fmt.Println(quoted) // price: \$10\.00 \(USD\)
```

### ClearCache

```go
func ClearCache()
```

清空正则编译缓存。

### CacheSize

```go
func CacheSize() int
```

返回缓存中已编译的正则表达式数量。

## 内置常用模式

提供预定义的常用正则模式常量和验证函数：

```go
const (
    PatternEmail      // 邮箱
    PatternURL        // URL
    PatternIP         // IP 地址
    PatternIPv4       // IPv4
    PatternMobile     // 手机号
    PatternPhone      // 座机号
    PatternIDCard     // 身份证号
    PatternZipCode    // 邮编
    PatternAlpha      // 纯字母
    PatternNumeric    // 纯数字
    PatternAlphaNum   // 字母数字
    PatternChinese    // 中文
    PatternDate       // 日期 yyyy-MM-dd
    PatternDateTime   // 日期时间 yyyy-MM-dd HH:mm:ss
    PatternUUID       // UUID
    PatternHexColor   // 十六进制颜色
    PatternCreditCard // 信用卡号
)
```

### IsEmail

```go
func IsEmail(s string) bool
```

验证是否为有效邮箱地址。

### IsURL

```go
func IsURL(s string) bool
```

验证是否为有效 URL。

### IsIP

```go
func IsIP(s string) bool
```

验证是否为有效 IPv4 地址。

### IsMobile

```go
func IsMobile(s string) bool
```

验证是否为中国大陆手机号。

**示例:**

```go
fmt.Println(regexutil.IsMobile("13812345678")) // true
fmt.Println(regexutil.IsMobile("12345678901")) // false
```

### IsPhone

```go
func IsPhone(s string) bool
```

验证是否为有效座机号码。

### IsIDCard

```go
func IsIDCard(s string) bool
```

验证是否为有效身份证号。

### IsZipCode

```go
func IsZipCode(s string) bool
```

验证是否为有效邮编（6位数字）。

### IsAlpha

```go
func IsAlpha(s string) bool
```

验证是否只包含字母。

### IsNumeric

```go
func IsNumeric(s string) bool
```

验证是否只包含数字。

### IsAlphaNumeric

```go
func IsAlphaNumeric(s string) bool
```

验证是否只包含字母和数字。

### IsChinese

```go
func IsChinese(s string) bool
```

验证是否只包含中文字符。

### ContainsChinese

```go
func ContainsChinese(s string) bool
```

检查是否包含中文字符。

### IsDate

```go
func IsDate(s string) bool
```

验证是否为有效日期格式（yyyy-MM-dd）。

### IsDateTime

```go
func IsDateTime(s string) bool
```

验证是否为有效日期时间格式（yyyy-MM-dd HH:mm:ss）。

### IsUUID

```go
func IsUUID(s string) bool
```

验证是否为有效 UUID。

### IsHexColor

```go
func IsHexColor(s string) bool
```

验证是否为有效十六进制颜色值。

## 完整示例

```go
package main

import (
    "fmt"

    "github.com/camark/GoHutools/regexutil"
)

func main() {
    // 表单验证
    email := "user@example.com"
    fmt.Printf("邮箱 %s 有效: %v\n", email, regexutil.IsEmail(email))

    mobile := "13812345678"
    fmt.Printf("手机号 %s 有效: %v\n", mobile, regexutil.IsMobile(mobile))

    // 提取数据
    text := "订单号: ORD-2024-001, 金额: ¥128.50"
    orderNo := regexutil.Find(text, `ORD-\d{4}-\d{3}`)
    amount := regexutil.Find(text, `\d+\.\d{2}`)
    fmt.Printf("订单号: %s, 金额: %s\n", orderNo, amount)

    // 批量提取
    html := `<img src="a.png"><img src="b.jpg"><img src="c.gif">`
    images := regexutil.FindAll(html, `src="([^"]+)"`)
    fmt.Println("图片:", images)

    // 替换脱敏
    phone := "联系电话: 13812345678"
    masked := regexutil.Replace(phone, `(\d{3})\d{4}(\d{4})`, "$1****$2")
    fmt.Println(masked) // 联系电话: 138****5678

    // 分割
    csv := "name,age;city  country"
    fields := regexutil.Split(csv, `[,;\s]+`)
    fmt.Println(fields) // [name age city country]
}
```
