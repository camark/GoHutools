---
title: strutil - 字符串工具
---

# strutil - 字符串工具

字符串处理工具包，提供了丰富的字符串操作函数，包括判空检查、截取、填充、转换、格式化等功能。

## 导入

```go
import "github.com/gongm/gohutool/strutil"
```

## 函数列表

### IsEmpty

```go
func IsEmpty(s string) bool
```

检查字符串是否为空（长度为0）。

**示例:**

```go
fmt.Println(strutil.IsEmpty(""))      // true
fmt.Println(strutil.IsEmpty("hello")) // false
```

### IsNotEmpty

```go
func IsNotEmpty(s string) bool
```

检查字符串是否不为空。

**示例:**

```go
fmt.Println(strutil.IsNotEmpty(""))      // false
fmt.Println(strutil.IsNotEmpty("hello")) // true
```

### IsBlank

```go
func IsBlank(s string) bool
```

检查字符串是否为空白（空或仅包含空白字符）。

**示例:**

```go
fmt.Println(strutil.IsBlank(""))      // true
fmt.Println(strutil.IsBlank("  "))    // true
fmt.Println(strutil.IsBlank("hello")) // false
```

### IsNotBlank

```go
func IsNotBlank(s string) bool
```

检查字符串是否不为空白。

**示例:**

```go
fmt.Println(strutil.IsNotBlank(""))      // false
fmt.Println(strutil.IsNotBlank("hello")) // true
```

### HasBlank

```go
func HasBlank(strs ...string) bool
```

检查是否包含空白字符串。

**示例:**

```go
fmt.Println(strutil.HasBlank("a", "", "b"))   // true
fmt.Println(strutil.HasBlank("a", "b", "c"))  // false
```

### HasEmpty

```go
func HasEmpty(strs ...string) bool
```

检查是否包含空字符串。

**示例:**

```go
fmt.Println(strutil.HasEmpty("a", "", "b"))   // true
fmt.Println(strutil.HasEmpty("a", " ", "b"))  // false
```

### Trim

```go
func Trim(s string) string
```

去除字符串首尾空白字符。

**示例:**

```go
fmt.Println(strutil.Trim("  hello  ")) // "hello"
```

### TrimPrefix

```go
func TrimPrefix(s, prefix string) string
```

去除字符串前缀（如果存在）。

**示例:**

```go
fmt.Println(strutil.TrimPrefix("hello world", "hello ")) // "world"
fmt.Println(strutil.TrimPrefix("hello world", "foo "))   // "hello world"
```

### TrimSuffix

```go
func TrimSuffix(s, suffix string) string
```

去除字符串后缀（如果存在）。

**示例:**

```go
fmt.Println(strutil.TrimSuffix("hello world", " world")) // "hello"
fmt.Println(strutil.TrimSuffix("hello world", " foo"))   // "hello world"
```

### Sub

```go
func Sub(s string, start, end int) string
```

提取子字符串，支持负数索引。

**示例:**

```go
fmt.Println(strutil.Sub("hello", 1, 3))   // "el"
fmt.Println(strutil.Sub("hello", -3, -1)) // "ll"
fmt.Println(strutil.Sub("hello", 0, -1))  // "hell"
```

### SubBefore

```go
func SubBefore(s, separator string, isLastSeparator bool) string
```

获取分隔符之前的子字符串。

**示例:**

```go
fmt.Println(strutil.SubBefore("a.b.c", ".", false)) // "a"
fmt.Println(strutil.SubBefore("a.b.c", ".", true))  // "a.b"
```

### SubAfter

```go
func SubAfter(s, separator string, isLastSeparator bool) string
```

获取分隔符之后的子字符串。

**示例:**

```go
fmt.Println(strutil.SubAfter("a.b.c", ".", false)) // "b.c"
fmt.Println(strutil.SubAfter("a.b.c", ".", true))  // "c"
```

### Contains

```go
func Contains(s, sub string) bool
```

检查字符串是否包含子字符串。

**示例:**

```go
fmt.Println(strutil.Contains("hello world", "world")) // true
fmt.Println(strutil.Contains("hello world", "foo"))   // false
```

### ContainsIgnoreCase

```go
func ContainsIgnoreCase(s, sub string) bool
```

检查字符串是否包含子字符串（忽略大小写）。

**示例:**

```go
fmt.Println(strutil.ContainsIgnoreCase("Hello World", "hello")) // true
fmt.Println(strutil.ContainsIgnoreCase("Hello World", "WORLD")) // true
```

### StartsWith

```go
func StartsWith(s, prefix string) bool
```

检查字符串是否以指定前缀开头。

**示例:**

```go
fmt.Println(strutil.StartsWith("hello world", "hello")) // true
fmt.Println(strutil.StartsWith("hello world", "world")) // false
```

### EndsWith

```go
func EndsWith(s, suffix string) bool
```

检查字符串是否以指定后缀结尾。

**示例:**

```go
fmt.Println(strutil.EndsWith("hello world", "world")) // true
fmt.Println(strutil.EndsWith("hello world", "hello")) // false
```

### RemovePrefix

```go
func RemovePrefix(s, prefix string) string
```

移除前缀（如果以该前缀开头）。

**示例:**

```go
fmt.Println(strutil.RemovePrefix("hello world", "hello ")) // "world"
fmt.Println(strutil.RemovePrefix("hello world", "foo "))   // "hello world"
```

### RemoveSuffix

```go
func RemoveSuffix(s, suffix string) string
```

移除后缀（如果以该后缀结尾）。

**示例:**

```go
fmt.Println(strutil.RemoveSuffix("hello world", " world")) // "hello"
fmt.Println(strutil.RemoveSuffix("hello world", " foo"))   // "hello world"
```

### Pad

```go
func Pad(s string, size int, padChar rune) string
```

居中填充字符串到指定长度。

**示例:**

```go
fmt.Println(strutil.Pad("hello", 10, '*'))  // "**hello***"
fmt.Println(strutil.Pad("hello", 3, '*'))   // "hello"
```

### PadLeft

```go
func PadLeft(s string, size int, padChar rune) string
```

左侧填充字符串到指定长度。

**示例:**

```go
fmt.Println(strutil.PadLeft("hello", 10, '*')) // "*****hello"
```

### PadRight

```go
func PadRight(s string, size int, padChar rune) string
```

右侧填充字符串到指定长度。

**示例:**

```go
fmt.Println(strutil.PadRight("hello", 10, '*')) // "hello*****"
```

### Reverse

```go
func Reverse(s string) string
```

反转字符串。

**示例:**

```go
fmt.Println(strutil.Reverse("hello")) // "olleh"
fmt.Println(strutil.Reverse("你好"))   // "好你"
```

### Repeat

```go
func Repeat(s string, count int) string
```

重复字符串指定次数。

**示例:**

```go
fmt.Println(strutil.Repeat("ab", 3)) // "ababab"
```

### Join

```go
func Join(separator string, strs ...string) string
```

使用分隔符连接字符串。

**示例:**

```go
fmt.Println(strutil.Join(", ", "a", "b", "c")) // "a, b, c"
```

### Split

```go
func Split(s, separator string) []string
```

按分隔符分割字符串。

**示例:**

```go
fmt.Println(strutil.Split("a,b,c", ",")) // [a b c]
```

### UpperFirst

```go
func UpperFirst(s string) string
```

首字母大写。

**示例:**

```go
fmt.Println(strutil.UpperFirst("hello")) // "Hello"
fmt.Println(strutil.UpperFirst("Hello")) // "Hello"
```

### LowerFirst

```go
func LowerFirst(s string) string
```

首字母小写。

**示例:**

```go
fmt.Println(strutil.LowerFirst("Hello")) // "hello"
fmt.Println(strutil.LowerFirst("hello")) // "hello"
```

### CamelToUnderline

```go
func CamelToUnderline(s string) string
```

驼峰命名转下划线命名。

**示例:**

```go
fmt.Println(strutil.CamelToUnderline("helloWorld"))  // "hello_world"
fmt.Println(strutil.CamelToUnderline("HelloWorld"))   // "_hello_world"
fmt.Println(strutil.CamelToUnderline("getHTTPClient")) // "get_h_t_t_p_client"
```

### UnderlineToCamel

```go
func UnderlineToCamel(s string) string
```

下划线命名转驼峰命名。

**示例:**

```go
fmt.Println(strutil.UnderlineToCamel("hello_world")) // "helloWorld"
fmt.Println(strutil.UnderlineToCamel("_hello_world")) // "HelloWorld"
```

### Format

```go
func Format(template string, args ...interface{}) string
```

格式化字符串，使用 `{0}`, `{1}` 等占位符。

**示例:**

```go
fmt.Println(strutil.Format("Hello {0}, you are {1} years old", "Alice", 25))
// "Hello Alice, you are 25 years old"
```

### Count

```go
func Count(s, sub string) int
```

统计子字符串出现次数。

**示例:**

```go
fmt.Println(strutil.Count("hello world hello", "hello")) // 2
fmt.Println(strutil.Count("hello", "l"))                  // 2
```

### Wrap

```go
func Wrap(s, prefix, suffix string) string
```

用前后缀包裹字符串。

**示例:**

```go
fmt.Println(strutil.Wrap("hello", "[", "]")) // "[hello]"
fmt.Println(strutil.Wrap("test", "'", "'"))  // "'test'"
```

### Unwrap

```go
func Unwrap(s, prefix, suffix string) string
```

移除字符串的前后缀（如果存在）。

**示例:**

```go
fmt.Println(strutil.Unwrap("[hello]", "[", "]")) // "hello"
fmt.Println(strutil.Unwrap("hello", "[", "]"))   // "hello"
```

### DefaultIfEmpty

```go
func DefaultIfEmpty(s, defaultStr string) string
```

如果字符串为空，返回默认值。

**示例:**

```go
fmt.Println(strutil.DefaultIfEmpty("", "default"))    // "default"
fmt.Println(strutil.DefaultIfEmpty("hello", "default")) // "hello"
```

### DefaultIfBlank

```go
func DefaultIfBlank(s, defaultStr string) string
```

如果字符串为空白，返回默认值。

**示例:**

```go
fmt.Println(strutil.DefaultIfBlank("", "default"))     // "default"
fmt.Println(strutil.DefaultIfBlank("  ", "default"))   // "default"
fmt.Println(strutil.DefaultIfBlank("hello", "default")) // "hello"
```
