---
title: charsetutil
---

# charsetutil

字符集工具包，提供字符集检测、转换、UTF-8 校验等功能，类似于 Java Hutool 的 CharsetUtil。

## 导入

```go
import "github.com/gongm/gohutool/charsetutil"
```

## 常量

```go
const (
    UTF8    = "UTF-8"
    GBK     = "GBK"
    GB2312  = "GB2312"
    GB18030 = "GB18030"
    BIG5    = "Big5"
    ISO8859 = "ISO-8859-1"
    USASCII = "US-ASCII"
    UTF16   = "UTF-16"
    UTF16LE = "UTF-16LE"
    UTF16BE = "UTF-16BE"
)
```

## 函数列表

### Convert

```go
func Convert(data []byte, fromCharset, toCharset string) ([]byte, error)
```

将字节数据从一种字符集转换为另一种字符集。

**示例:**

```go
utf8Data := []byte("你好")
gbkData, err := charsetutil.Convert(utf8Data, charsetutil.UTF8, charsetutil.GBK)
```

### ConvertString

```go
func ConvertString(s string, fromCharset, toCharset string) (string, error)
```

将字符串从一种字符集转换为另一种字符集。

**示例:**

```go
gbkStr, err := charsetutil.ConvertString("Hello", charsetutil.UTF8, charsetutil.GBK)
```

### ToUTF8

```go
func ToUTF8(data []byte, fromCharset string) ([]byte, error)
```

将字节数据转换为 UTF-8 编码。

**示例:**

```go
utf8Data, err := charsetutil.ToUTF8(gbkData, charsetutil.GBK)
```

### FromUTF8

```go
func FromUTF8(data []byte, toCharset string) ([]byte, error)
```

将 UTF-8 编码的字节数据转换为目标字符集。

**示例:**

```go
gbkData, err := charsetutil.FromUTF8(utf8Data, charsetutil.GBK)
```

### ToUTF8String

```go
func ToUTF8String(s string, fromCharset string) (string, error)
```

将字符串转换为 UTF-8 编码的字符串。

**示例:**

```go
utf8Str, err := charsetutil.ToUTF8String(gbkStr, charsetutil.GBK)
```

### FromUTF8String

```go
func FromUTF8String(s string, toCharset string) (string, error)
```

将 UTF-8 字符串转换为目标字符集的字符串。

**示例:**

```go
gbkStr, err := charsetutil.FromUTF8String("你好", charsetutil.GBK)
```

### GBKToUTF8

```go
func GBKToUTF8(data []byte) ([]byte, error)
```

将 GBK 编码的字节数据转换为 UTF-8。

**示例:**

```go
utf8Data, err := charsetutil.GBKToUTF8(gbkData)
```

### UTF8ToGBK

```go
func UTF8ToGBK(data []byte) ([]byte, error)
```

将 UTF-8 编码的字节数据转换为 GBK。

**示例:**

```go
gbkData, err := charsetutil.UTF8ToGBK(utf8Data)
```

### GBKToUTF8String

```go
func GBKToUTF8String(s string) (string, error)
```

将 GBK 字符串转换为 UTF-8 字符串。

**示例:**

```go
utf8Str, err := charsetutil.GBKToUTF8String(gbkStr)
```

### UTF8ToGBKString

```go
func UTF8ToGBKString(s string) (string, error)
```

将 UTF-8 字符串转换为 GBK 字符串。

**示例:**

```go
gbkStr, err := charsetutil.UTF8ToGBKString("你好")
```

### IsUTF8

```go
func IsUTF8(data []byte) bool
```

检查字节数据是否为有效的 UTF-8 编码。

**示例:**

```go
if charsetutil.IsUTF8(data) {
    fmt.Println("是有效的 UTF-8 编码")
}
```

### IsGBK

```go
func IsGBK(data []byte) bool
```

检查字节数据是否为有效的 GBK 编码。

**示例:**

```go
if charsetutil.IsGBK(data) {
    fmt.Println("是有效的 GBK 编码")
}
```

### DetectCharset

```go
func DetectCharset(data []byte) string
```

尝试检测字节数据的字符集。优先检测 UTF-8，其次 GBK，否则返回 ISO-8859-1。

**示例:**

```go
charset := charsetutil.DetectCharset(data)
fmt.Println(charset) // UTF-8, GBK, 或 ISO-8859-1
```

### CleanUTF8

```go
func CleanUTF8(data []byte) []byte
```

移除字节数据中无效的 UTF-8 序列，替换为 Unicode 替换字符。

**示例:**

```go
clean := charsetutil.CleanUTF8(data)
```

### CleanUTF8String

```go
func CleanUTF8String(s string) string
```

移除字符串中无效的 UTF-8 序列。

**示例:**

```go
clean := charsetutil.CleanUTF8String(s)
```

### RuneCount

```go
func RuneCount(s string) int
```

返回字符串中 Unicode 字符（rune）的数量。

**示例:**

```go
count := charsetutil.RuneCount("你好World")
fmt.Println(count) // 7
```

### RuneLen

```go
func RuneLen(r rune) int
```

返回 rune 编码所需的字节长度。

**示例:**

```go
length := charsetutil.RuneLen('中')
fmt.Println(length) // 3
```

### EncodeRune

```go
func EncodeRune(r rune) []byte
```

将 rune 编码为 UTF-8 字节切片。

**示例:**

```go
bytes := charsetutil.EncodeRune('中')
fmt.Println(bytes) // [228 184 173]
```

### DecodeRune

```go
func DecodeRune(p []byte) (rune, int)
```

从 UTF-8 字节切片解码一个 rune，返回 rune 值和字节数。

**示例:**

```go
r, size := charsetutil.DecodeRune([]byte("中"))
fmt.Printf("%c, %d bytes\n", r, size) // 中, 3 bytes
```

### ValidRune

```go
func ValidRune(r rune) bool
```

检查 rune 是否为有效的 Unicode 字符。

**示例:**

```go
if charsetutil.ValidRune(r) {
    fmt.Println("有效的 Unicode 字符")
}
```

### UTF16ToUTF8

```go
func UTF16ToUTF8(data []byte, littleEndian bool) ([]byte, error)
```

将 UTF-16 编码的数据转换为 UTF-8。

**示例:**

```go
utf8Data, err := charsetutil.UTF16ToUTF8(utf16Data, true) // little-endian
```

### UTF8ToUTF16

```go
func UTF8ToUTF16(data []byte, littleEndian bool) ([]byte, error)
```

将 UTF-8 编码的数据转换为 UTF-16。

**示例:**

```go
utf16Data, err := charsetutil.UTF8ToUTF16(utf8Data, false) // big-endian
```

### Charsets

```go
func Charsets() []string
```

返回所有支持的字符集列表。

**示例:**

```go
charsets := charsetutil.Charsets()
fmt.Println(charsets) // [UTF-8 GBK GB2312 GB18030 Big5 ISO-8859-1 US-ASCII UTF-16 UTF-16LE UTF-16BE]
```

### IsSupported

```go
func IsSupported(charset string) bool
```

检查指定字符集是否被支持。

**示例:**

```go
if charsetutil.IsSupported("GBK") {
    fmt.Println("GBK 受支持")
}
```
