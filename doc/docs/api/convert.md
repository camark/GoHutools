---
title: convert - 转换工具
---

# convert - 转换工具

类型转换工具包，提供了丰富的类型转换函数，包括基本类型转换、编码解码、进制转换、中文数字转换等功能。

## 导入

```go
import "github.com/camark/GoHutools/convert"
```

## 函数列表

### ToStr

```go
func ToStr(v interface{}) string
```

将任意值转换为字符串。

**示例:**

```go
fmt.Println(convert.ToStr(123))      // "123"
fmt.Println(convert.ToStr(3.14))     // "3.14"
fmt.Println(convert.ToStr(true))     // "true"
fmt.Println(convert.ToStr(nil))      // ""
```

### ToInt

```go
func ToInt(v interface{}) (int, error)
```

将任意值转换为 int。

**示例:**

```go
n, err := convert.ToInt("123")
fmt.Println(n, err)  // 123 <nil>

n, err = convert.ToInt(3.14)
fmt.Println(n, err)  // 3 <nil>

n, err = convert.ToInt(true)
fmt.Println(n, err)  // 1 <nil>
```

### MustToInt

```go
func MustToInt(v interface{}) int
```

将任意值转换为 int，失败时 panic。

**示例:**

```go
n := convert.MustToInt("123")
fmt.Println(n)  // 123
```

### ToInt64

```go
func ToInt64(v interface{}) (int64, error)
```

将任意值转换为 int64。

**示例:**

```go
n, err := convert.ToInt64("1234567890")
fmt.Println(n, err)  // 1234567890 <nil>
```

### ToFloat64

```go
func ToFloat64(v interface{}) (float64, error)
```

将任意值转换为 float64。

**示例:**

```go
f, err := convert.ToFloat64("3.14")
fmt.Println(f, err)  // 3.14 <nil>

f, err = convert.ToFloat64(42)
fmt.Println(f, err)  // 42 <nil>
```

### MustToFloat64

```go
func MustToFloat64(v interface{}) float64
```

将任意值转换为 float64，失败时 panic。

**示例:**

```go
f := convert.MustToFloat64("3.14")
fmt.Println(f)  // 3.14
```

### ToBool

```go
func ToBool(v interface{}) (bool, error)
```

将任意值转换为 bool。

**示例:**

```go
b, err := convert.ToBool("true")
fmt.Println(b, err)  // true <nil>

b, err = convert.ToBool("1")
fmt.Println(b, err)  // true <nil>

b, err = convert.ToBool("yes")
fmt.Println(b, err)  // true <nil>

b, err = convert.ToBool(0)
fmt.Println(b, err)  // false <nil>
```

### MustToBool

```go
func MustToBool(v interface{}) bool
```

将任意值转换为 bool，失败时 panic。

**示例:**

```go
b := convert.MustToBool("true")
fmt.Println(b)  // true
```

### ToBytes

```go
func ToBytes(s string) []byte
```

将字符串转换为字节数组。

**示例:**

```go
b := convert.ToBytes("hello")
fmt.Println(b)  // [104 101 108 108 111]
```

### ToString

```go
func ToString(b []byte) string
```

将字节数组转换为字符串。

**示例:**

```go
s := convert.ToString([]byte{104, 101, 108, 108, 111})
fmt.Println(s)  // "hello"
```

### ToIntSlice

```go
func ToIntSlice(arr []interface{}) []int
```

将 interface{} 切片转换为 int 切片。

**示例:**

```go
arr := []interface{}{1, "2", 3.0, "4"}
result := convert.ToIntSlice(arr)
fmt.Println(result)  // [1, 2, 3, 4]
```

### ToStrSlice

```go
func ToStrSlice(arr []interface{}) []string
```

将 interface{} 切片转换为 string 切片。

**示例:**

```go
arr := []interface{}{1, "hello", true, 3.14}
result := convert.ToStrSlice(arr)
fmt.Println(result)  // [1, hello, true, 3.14]
```

### ToMap

```go
func ToMap(v interface{}) (map[string]interface{}, error)
```

将值转换为 map[string]interface{}。

**示例:**

```go
type Person struct {
    Name string
    Age  int
}
p := Person{Name: "Alice", Age: 25}
m, err := convert.ToMap(p)
fmt.Println(m, err)  // map[Name:Alice Age:25] <nil>
```

### CamelToUnderline

```go
func CamelToUnderline(s string) string
```

驼峰命名转下划线命名。

**示例:**

```go
fmt.Println(convert.CamelToUnderline("helloWorld"))  // "hello_world"
fmt.Println(convert.CamelToUnderline("HelloWorld"))   // "_hello_world"
```

### UnderlineToCamel

```go
func UnderlineToCamel(s string) string
```

下划线命名转驼峰命名。

**示例:**

```go
fmt.Println(convert.UnderlineToCamel("hello_world"))  // "HelloWorld"
fmt.Println(convert.UnderlineToCamel("_hello_world")) // "HelloWorld"
```

### HexToLong

```go
func HexToLong(hexStr string) (int64, error)
```

将十六进制字符串转换为 int64。

**示例:**

```go
n, err := convert.HexToLong("ff")
fmt.Println(n, err)  // 255 <nil>

n, err = convert.HexToLong("0xFF")
fmt.Println(n, err)  // 255 <nil>
```

### LongToHex

```go
func LongToHex(n int64) string
```

将 int64 转换为十六进制字符串。

**示例:**

```go
fmt.Println(convert.LongToHex(255))   // "ff"
fmt.Println(convert.LongToHex(4096))  // "1000"
```

### NumberToChinese

```go
func NumberToChinese(n int64, isUseTraditional bool) string
```

将数字转换为中文表示。

**示例:**

```go
fmt.Println(convert.NumberToChinese(12345, false))     // "一万二千三百四十五"
fmt.Println(convert.NumberToChinese(12345, true))      // "壹萬贰仟叁佰肆拾伍"
fmt.Println(convert.NumberToChinese(0, false))          // "零"
fmt.Println(convert.NumberToChinese(-100, false))       // "负一百"
```

### ChineseToNumber

```go
func ChineseToNumber(chinese string) (int64, error)
```

将中文数字转换为 int64。

**示例:**

```go
n, err := convert.ChineseToNumber("一万二千三百四十五")
fmt.Println(n, err)  // 12345 <nil>

n, err = convert.ChineseToNumber("壹萬贰仟叁佰肆拾伍")
fmt.Println(n, err)  // 12345 <nil>
```

### DigitToChinese

```go
func DigitToChinese(digit byte) string
```

将数字转换为大写中文数字（壹贰叁...）。

**示例:**

```go
fmt.Println(convert.DigitToChinese(0))  // "零"
fmt.Println(convert.DigitToChinese(1))  // "壹"
fmt.Println(convert.DigitToChinese(9))  // "玖"
```

### RomanToInt

```go
func RomanToInt(roman string) (int, error)
```

将罗马数字转换为 int。

**示例:**

```go
n, err := convert.RomanToInt("XIV")
fmt.Println(n, err)  // 14 <nil>

n, err = convert.RomanToInt("MCMXCIV")
fmt.Println(n, err)  // 1994 <nil>
```

### IntToRoman

```go
func IntToRoman(n int) string
```

将 int 转换为罗马数字（1-3999）。

**示例:**

```go
fmt.Println(convert.IntToRoman(14))    // "XIV"
fmt.Println(convert.IntToRoman(1994))  // "MCMXCIV"
fmt.Println(convert.IntToRoman(2026))  // "MMXXVI"
```

### BaseConvert

```go
func BaseConvert(value string, fromBase, toBase int) (string, error)
```

进制转换。

**示例:**

```go
result, err := convert.BaseConvert("ff", 16, 10)
fmt.Println(result, err)  // "255" <nil>

result, err = convert.BaseConvert("255", 10, 16)
fmt.Println(result, err)  // "ff" <nil>

result, err = convert.BaseConvert("11111111", 2, 16)
fmt.Println(result, err)  // "ff" <nil>
```

### ToEnum

```go
func ToEnum(s string, enumMap map[string]int) (int, bool)
```

将字符串转换为枚举值。

**示例:**

```go
statusMap := map[string]int{
    "active":   1,
    "inactive": 0,
    "deleted":  -1,
}
val, ok := convert.ToEnum("active", statusMap)
fmt.Println(val, ok)  // 1 true

val, ok = convert.ToEnum("unknown", statusMap)
fmt.Println(val, ok)  // 0 false
```

### Base64Encode

```go
func Base64Encode(s string) string
```

Base64 编码。

**示例:**

```go
encoded := convert.Base64Encode("hello world")
fmt.Println(encoded)  // "aGVsbG8gd29ybGQ="
```

### Base64Decode

```go
func Base64Decode(s string) (string, error)
```

Base64 解码。

**示例:**

```go
decoded, err := convert.Base64Decode("aGVsbG8gd29ybGQ=")
fmt.Println(decoded, err)  // "hello world" <nil>
```

### HexEncode

```go
func HexEncode(data []byte) string
```

将字节数组编码为十六进制字符串。

**示例:**

```go
encoded := convert.HexEncode([]byte("hello"))
fmt.Println(encoded)  // "68656c6c6f"
```

### HexDecode

```go
func HexDecode(s string) ([]byte, error)
```

将十六进制字符串解码为字节数组。

**示例:**

```go
decoded, err := convert.HexDecode("68656c6c6f")
fmt.Println(string(decoded), err)  // "hello" <nil>
```

### Reverse

```go
func Reverse(s string) string
```

反转字符串。

**示例:**

```go
fmt.Println(convert.Reverse("hello"))  // "olleh"
```

### IsNumeric

```go
func IsNumeric(v interface{}) bool
```

检查值是否为数值类型。

**示例:**

```go
fmt.Println(convert.IsNumeric(123))      // true
fmt.Println(convert.IsNumeric(3.14))     // true
fmt.Println(convert.IsNumeric("123"))    // true
fmt.Println(convert.IsNumeric("hello"))  // false
```

### TypeOf

```go
func TypeOf(v interface{}) string
```

返回值的类型名称。

**示例:**

```go
fmt.Println(convert.TypeOf(123))      // "int"
fmt.Println(convert.TypeOf("hello"))  // "string"
fmt.Println(convert.TypeOf(3.14))     // "float64"
fmt.Println(convert.TypeOf(nil))      // "nil"
```

### Round

```go
func Round(f float64, decimals int) float64
```

四舍五入到指定小数位。

**示例:**

```go
fmt.Println(convert.Round(3.14159, 2))  // 3.14
fmt.Println(convert.Round(3.14659, 2))  // 3.15
```
