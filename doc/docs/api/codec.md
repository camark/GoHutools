---
title: codec
---

# codec

编解码工具包，提供 Base64、Hex、URL、HTML、Unicode、Punycode 等编解码功能。

## 导入

```go
import "github.com/gongm/gohutool/codec"
```

## 函数列表

### Base64 编解码

#### Base64Encode

```go
func Base64Encode(data []byte) string
```

将字节数据编码为标准 Base64 字符串。

**示例:**

```go
encoded := codec.Base64Encode([]byte("Hello, World!"))
fmt.Println(encoded) // SGVsbG8sIFdvcmxkIQ==
```

#### Base64Decode

```go
func Base64Decode(s string) ([]byte, error)
```

将标准 Base64 字符串解码为字节数据。

**示例:**

```go
decoded, err := codec.Base64Decode("SGVsbG8sIFdvcmxkIQ==")
fmt.Println(string(decoded)) // Hello, World!
```

#### Base64URLEncode

```go
func Base64URLEncode(data []byte) string
```

将字节数据编码为 URL 安全的 Base64 字符串。

**示例:**

```go
encoded := codec.Base64URLEncode([]byte("Hello"))
```

#### Base64URLDecode

```go
func Base64URLDecode(s string) ([]byte, error)
```

将 URL 安全的 Base64 字符串解码。

**示例:**

```go
decoded, err := codec.Base64URLDecode(encoded)
```

#### Base64StdEncode

```go
func Base64StdEncode(data []byte) string
```

将字节数据编码为标准 Base64 字符串（等同于 Base64Encode）。

#### Base64StdDecode

```go
func Base64StdDecode(s string) ([]byte, error)
```

将标准 Base64 字符串解码（等同于 Base64Decode）。

### Hex 编解码

#### HexEncode

```go
func HexEncode(data []byte) string
```

将字节数据编码为十六进制字符串（小写）。

**示例:**

```go
encoded := codec.HexEncode([]byte("Hello"))
fmt.Println(encoded) // 48656c6c6f
```

#### HexDecode

```go
func HexDecode(s string) ([]byte, error)
```

将十六进制字符串解码为字节数据。

**示例:**

```go
decoded, err := codec.HexDecode("48656c6c6f")
fmt.Println(string(decoded)) // Hello
```

### URL 编解码

#### URLEncode

```go
func URLEncode(s string) string
```

对字符串进行 URL 编码（查询参数编码）。

**示例:**

```go
encoded := codec.URLEncode("Hello World&foo=bar")
fmt.Println(encoded) // Hello+World%26foo%3Dbar
```

#### URLDecode

```go
func URLDecode(s string) (string, error)
```

对 URL 编码的字符串进行解码。

**示例:**

```go
decoded, err := codec.URLDecode("Hello+World%26foo%3Dbar")
fmt.Println(decoded) // Hello World&foo=bar
```

#### URLQueryEncode

```go
func URLQueryEncode(s string) string
```

对 URL 查询参数进行编码（等同于 URLEncode）。

#### URLQueryDecode

```go
func URLQueryDecode(s string) (string, error)
```

对 URL 查询参数进行解码（等同于 URLDecode）。

### HTML 编解码

#### HTMLEscape

```go
func HTMLEscape(s string) string
```

对 HTML 特殊字符进行转义。

**示例:**

```go
escaped := codec.HTMLEscape(`<div class="test">Hello & World</div>`)
fmt.Println(escaped) // &lt;div class=&#34;test&#34;&gt;Hello &amp; World&lt;/div&gt;
```

#### HTMLUnescape

```go
func HTMLUnescape(s string) string
```

还原 HTML 转义字符。

**示例:**

```go
unescaped := codec.HTMLUnescape("&lt;p&gt;Hello &amp; World&lt;/p&gt;")
fmt.Println(unescaped) // <p>Hello & World</p>
```

### Unicode 编解码

#### UnicodeEncode

```go
func UnicodeEncode(s string) string
```

将字符串中的非 ASCII 字符编码为 Unicode 转义序列（`\uXXXX`）。

**示例:**

```go
encoded := codec.UnicodeEncode("你好Hello")
fmt.Println(encoded) // 你好Hello
```

#### UnicodeDecode

```go
func UnicodeDecode(s string) (string, error)
```

将 Unicode 转义序列解码为字符串。

**示例:**

```go
decoded, err := codec.UnicodeDecode(`你好Hello`)
fmt.Println(decoded) // 你好Hello
```

### Punycode 编解码

#### PunycodeEncode

```go
func PunycodeEncode(s string) (string, error)
```

将国际化域名编码为 Punycode。

**示例:**

```go
encoded, err := codec.PunycodeEncode("münchen.de")
fmt.Println(encoded) // mnchen-3ya.de
```

#### PunycodeDecode

```go
func PunycodeDecode(s string) (string, error)
```

将 Punycode 域名解码为国际化域名。

**示例:**

```go
decoded, err := codec.PunycodeDecode("xn--mnchen-3ya.de")
fmt.Println(decoded) // münchen.de
```

## 完整示例

```go
package main

import (
    "fmt"

    "github.com/gongm/gohutool/codec"
)

func main() {
    // Base64
    encoded := codec.Base64Encode([]byte("Hello"))
    decoded, _ := codec.Base64Decode(encoded)
    fmt.Println(string(decoded)) // Hello

    // Hex
    hexStr := codec.HexEncode([]byte("Hello"))
    bytes, _ := codec.HexDecode(hexStr)
    fmt.Println(string(bytes)) // Hello

    // URL
    urlEncoded := codec.URLEncode("name=John&age=30")
    urlDecoded, _ := codec.URLDecode(urlEncoded)
    fmt.Println(urlDecoded)

    // HTML
    html := codec.HTMLEscape("<b>Bold</b>")
    fmt.Println(html)
    fmt.Println(codec.HTMLUnescape(html))

    // Unicode
    unicode := codec.UnicodeEncode("你好")
    fmt.Println(unicode)
    fmt.Println(codec.UnicodeDecode(unicode))
}
```
