---
title: validate - 验证工具
---

# validate - 验证工具

数据验证工具包，提供了丰富的数据验证函数，包括格式验证、范围检查、类型判断等功能。

## 导入

```go
import "github.com/camark/GoHutools/validate"
```

## 函数列表

### IsEmpty

```go
func IsEmpty(v interface{}) bool
```

检查值是否为空。

**示例:**

```go
fmt.Println(validate.IsEmpty(""))       // true
fmt.Println(validate.IsEmpty(0))        // true
fmt.Println(validate.IsEmpty(false))    // true
fmt.Println(validate.IsEmpty(nil))      // true
fmt.Println(validate.IsEmpty("hello"))  // false
```

### IsNotEmpty

```go
func IsNotEmpty(v interface{}) bool
```

检查值是否不为空。

**示例:**

```go
fmt.Println(validate.IsNotEmpty("hello"))  // true
fmt.Println(validate.IsNotEmpty(123))      // true
```

### IsNull

```go
func IsNull(v interface{}) bool
```

检查值是否为 nil。

**示例:**

```go
fmt.Println(validate.IsNull(nil))      // true
fmt.Println(validate.IsNull(""))       // false
fmt.Println(validate.IsNull(0))        // false
```

### IsNotNull

```go
func IsNotNull(v interface{}) bool
```

检查值是否不为 nil。

**示例:**

```go
fmt.Println(validate.IsNotNull("hello"))  // true
fmt.Println(validate.IsNotNull(nil))      // false
```

### IsBlank

```go
func IsBlank(s string) bool
```

检查字符串是否为空白。

**示例:**

```go
fmt.Println(validate.IsBlank(""))      // true
fmt.Println(validate.IsBlank("  "))    // true
fmt.Println(validate.IsBlank("hello")) // false
```

### IsNotBlank

```go
func IsNotBlank(s string) bool
```

检查字符串是否不为空白。

**示例:**

```go
fmt.Println(validate.IsNotBlank("hello"))  // true
```

### IsEmail

```go
func IsEmail(s string) bool
```

检查字符串是否为有效的电子邮件地址。

**示例:**

```go
fmt.Println(validate.IsEmail("user@example.com"))      // true
fmt.Println(validate.IsEmail("user.name+tag@domain.co")) // true
fmt.Println(validate.IsEmail("invalid-email"))          // false
fmt.Println(validate.IsEmail("@domain.com"))            // false
```

### IsURL

```go
func IsURL(s string) bool
```

检查字符串是否为有效的 URL。

**示例:**

```go
fmt.Println(validate.IsURL("https://example.com"))       // true
fmt.Println(validate.IsURL("http://localhost:8080/path")) // true
fmt.Println(validate.IsURL("ftp://files.example.com"))    // true
fmt.Println(validate.IsURL("not a url"))                  // false
```

### IsIP

```go
func IsIP(s string) bool
```

检查字符串是否为有效的 IP 地址。

**示例:**

```go
fmt.Println(validate.IsIP("192.168.1.1"))        // true
fmt.Println(validate.IsIP("::1"))                 // true
fmt.Println(validate.IsIP("2001:db8::1"))         // true
fmt.Println(validate.IsIP("invalid"))             // false
```

### IsIPv4

```go
func IsIPv4(s string) bool
```

检查字符串是否为有效的 IPv4 地址。

**示例:**

```go
fmt.Println(validate.IsIPv4("192.168.1.1"))  // true
fmt.Println(validate.IsIPv4("::1"))           // false
fmt.Println(validate.IsIPv4("invalid"))       // false
```

### IsIPv6

```go
func IsIPv6(s string) bool
```

检查字符串是否为有效的 IPv6 地址。

**示例:**

```go
fmt.Println(validate.IsIPv6("::1"))                 // true
fmt.Println(validate.IsIPv6("2001:db8::1"))         // true
fmt.Println(validate.IsIPv6("192.168.1.1"))          // false
```

### IsMAC

```go
func IsMAC(s string) bool
```

检查字符串是否为有效的 MAC 地址。

**示例:**

```go
fmt.Println(validate.IsMAC("00:1B:44:11:3A:B7"))  // true
fmt.Println(validate.IsMAC("00-1B-44-11-3A-B7"))  // true
fmt.Println(validate.IsMAC("invalid"))             // false
```

### IsPort

```go
func IsPort(port int) bool
```

检查端口号是否有效（0-65535）。

**示例:**

```go
fmt.Println(validate.IsPort(80))     // true
fmt.Println(validate.IsPort(8080))   // true
fmt.Println(validate.IsPort(65535))  // true
fmt.Println(validate.IsPort(70000))  // false
```

### IsMobile

```go
func IsMobile(s string) bool
```

检查字符串是否为中国手机号码。

**示例:**

```go
fmt.Println(validate.IsMobile("13812345678"))    // true
fmt.Println(validate.IsMobile("+8613812345678")) // true
fmt.Println(validate.IsMobile("12345678901"))    // false
fmt.Println(validate.IsMobile("1381234567"))     // false
```

### IsPhone

```go
func IsPhone(s string) bool
```

检查字符串是否为中国电话号码（座机或400/800）。

**示例:**

```go
fmt.Println(validate.IsPhone("010-12345678"))     // true
fmt.Println(validate.IsPhone("021-12345678"))     // true
fmt.Println(validate.IsPhone("400-123-4567"))     // true
fmt.Println(validate.IsPhone("800-123-4567"))     // true
fmt.Println(validate.IsPhone("13812345678"))      // false
```

### IsIDCard

```go
func IsIDCard(s string) bool
```

检查字符串是否为有效的中国身份证号码（18位）。

**示例:**

```go
fmt.Println(validate.IsIDCard("11010519491231002X"))  // true (示例)
fmt.Println(validate.IsIDCard("123456789012345678"))  // false
```

### IsZipCode

```go
func IsZipCode(s string) bool
```

检查字符串是否为有效的邮政编码（6位数字）。

**示例:**

```go
fmt.Println(validate.IsZipCode("100000"))  // true
fmt.Println(validate.IsZipCode("518000"))  // true
fmt.Println(validate.IsZipCode("12345"))   // false
```

### IsAlpha

```go
func IsAlpha(s string) bool
```

检查字符串是否只包含字母。

**示例:**

```go
fmt.Println(validate.IsAlpha("hello"))    // true
fmt.Println(validate.IsAlpha("Hello"))    // true
fmt.Println(validate.IsAlpha("hello123")) // false
fmt.Println(validate.IsAlpha(""))         // false
```

### IsAlphaNumeric

```go
func IsAlphaNumeric(s string) bool
```

检查字符串是否只包含字母和数字。

**示例:**

```go
fmt.Println(validate.IsAlphaNumeric("hello123"))  // true
fmt.Println(validate.IsAlphaNumeric("Hello"))     // true
fmt.Println(validate.IsAlphaNumeric("hello 123")) // false
```

### IsNumeric

```go
func IsNumeric(s string) bool
```

检查字符串是否只包含数字。

**示例:**

```go
fmt.Println(validate.IsNumeric("12345"))   // true
fmt.Println(validate.IsNumeric("123.45"))  // false
fmt.Println(validate.IsNumeric("123abc"))  // false
```

### IsInteger

```go
func IsInteger(s string) bool
```

检查字符串是否为整数。

**示例:**

```go
fmt.Println(validate.IsInteger("123"))    // true
fmt.Println(validate.IsInteger("-123"))   // true
fmt.Println(validate.IsInteger("+123"))   // true
fmt.Println(validate.IsInteger("123.45")) // false
```

### IsFloat

```go
func IsFloat(s string) bool
```

检查字符串是否为浮点数。

**示例:**

```go
fmt.Println(validate.IsFloat("123.45"))  // true
fmt.Println(validate.IsFloat("-123.45")) // true
fmt.Println(validate.IsFloat("123"))     // true
fmt.Println(validate.IsFloat("abc"))     // false
```

### IsUUID

```go
func IsUUID(s string) bool
```

检查字符串是否为有效的 UUID。

**示例:**

```go
fmt.Println(validate.IsUUID("550e8400-e29b-41d4-a716-446655440000"))  // true
fmt.Println(validate.IsUUID("invalid-uuid"))                          // false
```

### IsJSON

```go
func IsJSON(s string) bool
```

检查字符串是否为有效的 JSON。

**示例:**

```go
fmt.Println(validate.IsJSON(`{"name": "Alice", "age": 25}`))  // true
fmt.Println(validate.IsJSON(`[1, 2, 3]`))                     // true
fmt.Println(validate.IsJSON(`"hello"`))                       // true
fmt.Println(validate.IsJSON(`{invalid}`))                     // false
```

### IsBase64

```go
func IsBase64(s string) bool
```

检查字符串是否为有效的 Base64 编码。

**示例:**

```go
fmt.Println(validate.IsBase64("aGVsbG8gd29ybGQ="))  // true
fmt.Println(validate.IsBase64("SGVsbG8="))           // true
fmt.Println(validate.IsBase64("not base64!"))         // false
```

### IsHexColor

```go
func IsHexColor(s string) bool
```

检查字符串是否为有效的十六进制颜色值。

**示例:**

```go
fmt.Println(validate.IsHexColor("#FF0000"))   // true
fmt.Println(validate.IsHexColor("#ff0000"))   // true
fmt.Println(validate.IsHexColor("#F00"))      // true
fmt.Println(validate.IsHexColor("#FF0000FF")) // true
fmt.Println(validate.IsHexColor("red"))       // false
```

### IsRGBColor

```go
func IsRGBColor(s string) bool
```

检查字符串是否为有效的 RGB 颜色值。

**示例:**

```go
fmt.Println(validate.IsRGBColor("rgb(255, 0, 0)"))       // true
fmt.Println(validate.IsRGBColor("rgba(255, 0, 0, 0.5)")) // true
fmt.Println(validate.IsRGBColor("rgb(0, 0, 0)"))         // true
fmt.Println(validate.IsRGBColor("rgb(256, 0, 0)"))       // false
```

### IsDate

```go
func IsDate(s string) bool
```

检查字符串是否为有效的日期格式。

**示例:**

```go
fmt.Println(validate.IsDate("2026-08-14"))   // true
fmt.Println(validate.IsDate("2026/08/14"))   // true
fmt.Println(validate.IsDate("2026-13-01"))   // false
fmt.Println(validate.IsDate("not a date"))   // false
```

### IsChinese

```go
func IsChinese(s string) bool
```

检查字符串是否只包含中文字符。

**示例:**

```go
fmt.Println(validate.IsChinese("你好世界"))  // true
fmt.Println(validate.IsChinese("你好world")) // false
fmt.Println(validate.IsChinese(""))          // false
```

### ContainsChinese

```go
func ContainsChinese(s string) bool
```

检查字符串是否包含中文字符。

**示例:**

```go
fmt.Println(validate.ContainsChinese("hello你好"))  // true
fmt.Println(validate.ContainsChinese("hello"))      // false
```

### IsLetter

```go
func IsLetter(c rune) bool
```

检查字符是否为字母。

**示例:**

```go
fmt.Println(validate.IsLetter('a'))  // true
fmt.Println(validate.IsLetter('Z'))  // true
fmt.Println(validate.IsLetter('1'))  // false
```

### IsDigit

```go
func IsDigit(c rune) bool
```

检查字符是否为数字。

**示例:**

```go
fmt.Println(validate.IsDigit('0'))  // true
fmt.Println(validate.IsDigit('9'))  // true
fmt.Println(validate.IsDigit('a'))  // false
```

### IsUpperCase

```go
func IsUpperCase(c rune) bool
```

检查字符是否为大写字母。

**示例:**

```go
fmt.Println(validate.IsUpperCase('A'))  // true
fmt.Println(validate.IsUpperCase('a'))  // false
```

### IsLowerCase

```go
func IsLowerCase(c rune) bool
```

检查字符是否为小写字母。

**示例:**

```go
fmt.Println(validate.IsLowerCase('a'))  // true
fmt.Println(validate.IsLowerCase('A'))  // false
```

### Matches

```go
func Matches(s, pattern string) bool
```

检查字符串是否匹配正则表达式模式。

**示例:**

```go
fmt.Println(validate.Matches("hello123", `\w+\d+`))  // true
fmt.Println(validate.Matches("hello", `\d+`))         // false
```

### MatchesAll

```go
func MatchesAll(s, pattern string) bool
```

检查整个字符串是否匹配正则表达式模式。

**示例:**

```go
fmt.Println(validate.MatchesAll("12345", `\d+`))     // true
fmt.Println(validate.MatchesAll("123abc", `\d+`))    // false
```

### Range

```go
func Range(val, min, max int) bool
```

检查整数值是否在范围内。

**示例:**

```go
fmt.Println(validate.Range(5, 1, 10))    // true
fmt.Println(validate.Range(15, 1, 10))   // false
fmt.Println(validate.Range(1, 1, 10))    // true
```

### RangeFloat

```go
func RangeFloat(val, min, max float64) bool
```

检查浮点值是否在范围内。

**示例:**

```go
fmt.Println(validate.RangeFloat(5.5, 1.0, 10.0))  // true
fmt.Println(validate.RangeFloat(15.5, 1.0, 10.0)) // false
```

### MinLength

```go
func MinLength(s string, min int) bool
```

检查字符串是否满足最小长度。

**示例:**

```go
fmt.Println(validate.MinLength("hello", 3))   // true
fmt.Println(validate.MinLength("hi", 3))      // false
fmt.Println(validate.MinLength("你好", 2))     // true
```

### MaxLength

```go
func MaxLength(s string, max int) bool
```

检查字符串是否不超过最大长度。

**示例:**

```go
fmt.Println(validate.MaxLength("hello", 10))  // true
fmt.Println(validate.MaxLength("hello", 3))   // false
```

### LengthBetween

```go
func LengthBetween(s string, min, max int) bool
```

检查字符串长度是否在范围内。

**示例:**

```go
fmt.Println(validate.LengthBetween("hello", 3, 10))  // true
fmt.Println(validate.LengthBetween("hi", 3, 10))     // false
fmt.Println(validate.LengthBetween("hello world!", 3, 10))  // false
```

### CreditCard

```go
func CreditCard(s string) bool
```

检查字符串是否为有效的信用卡号（使用 Luhn 算法验证）。

**示例:**

```go
fmt.Println(validate.CreditCard("4111111111111111"))  // true (Visa)
fmt.Println(validate.CreditCard("5500000000000004"))  // true (Mastercard)
fmt.Println(validate.CreditCard("378282246310005"))   // true (Amex)
fmt.Println(validate.CreditCard("1234567890123456"))  // false
```
