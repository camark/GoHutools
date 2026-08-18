---
title: randomutil
---

# randomutil

随机数工具包，基于 crypto/rand 提供密码学安全的随机数生成，包括整数、浮点数、布尔值、字符串、字节、切片操作、加权随机、骰子、颜色、日期和中文模拟数据等功能。

## 导入

```go
import "github.com/camark/GoHutools/randomutil"
```

## 整数随机

### Int

```go
func Int(max int) (int, error)
```

生成 [0, max) 范围内的密码学安全随机整数。

**示例:**

```go
n, err := randomutil.Int(100)
fmt.Println(n) // 0-99 之间的随机数
```

### IntRange

```go
func IntRange(min, max int) (int, error)
```

生成 [min, max) 范围内的密码学安全随机整数。

**示例:**

```go
n, err := randomutil.IntRange(10, 20)
fmt.Println(n) // 10-19 之间的随机数
```

### MustInt

```go
func MustInt(max int) int
```

生成 [0, max) 范围内的随机整数，出错时 panic。

**示例:**

```go
n := randomutil.MustInt(100)
```

### MustIntRange

```go
func MustIntRange(min, max int) int
```

生成 [min, max) 范围内的随机整数，出错时 panic。

**示例:**

```go
n := randomutil.MustIntRange(1, 7) // 模拟骰子
```

### Int64

```go
func Int64(max int64) (int64, error)
```

生成 [0, max) 范围内的密码学安全随机 int64。

### Int64Range

```go
func Int64Range(min, max int64) (int64, error)
```

生成 [min, max) 范围内的密码学安全随机 int64。

## 浮点数随机

### Float64

```go
func Float64() (float64, error)
```

生成 [0.0, 1.0) 范围内的密码学安全随机 float64。

**示例:**

```go
f, err := randomutil.Float64()
fmt.Println(f)
```

### Float64Range

```go
func Float64Range(min, max float64) (float64, error)
```

生成 [min, max) 范围内的密码学安全随机 float64。

**示例:**

```go
f, err := randomutil.Float64Range(1.0, 10.0)
```

## 布尔随机

### Bool

```go
func Bool() (bool, error)
```

生成密码学安全的随机布尔值。

**示例:**

```go
b, err := randomutil.Bool()
```

### MustBool

```go
func MustBool() bool
```

生成随机布尔值，出错时 panic。

**示例:**

```go
if randomutil.MustBool() {
    fmt.Println("正面")
}
```

## 字符串随机

### String

```go
func String(length int) string
```

生成指定长度的随机字母数字字符串。

**示例:**

```go
s := randomutil.String(16)
fmt.Println(s) // 如: aB3xK9mN4pQ7rS1t
```

### StringWithAlphabet

```go
func StringWithAlphabet(length int, alphabet string) string
```

使用自定义字母表生成随机字符串。

**示例:**

```go
s := randomutil.StringWithAlphabet(8, "AEIOU12345")
```

### Alpha

```go
func Alpha(length int) string
```

生成指定长度的随机字母字符串（a-zA-Z）。

**示例:**

```go
s := randomutil.Alpha(10)
```

### AlphaNumeric

```go
func AlphaNumeric(length int) string
```

生成指定长度的随机字母数字字符串（a-zA-Z0-9）。

**示例:**

```go
s := randomutil.AlphaNumeric(12)
```

### Numeric

```go
func Numeric(length int) string
```

生成指定长度的随机数字字符串（0-9）。

**示例:**

```go
code := randomutil.Numeric(6) // 验证码
fmt.Println(code) // 如: 384921
```

### Hex

```go
func Hex(length int) string
```

生成指定长度的随机小写十六进制字符串。

**示例:**

```go
s := randomutil.Hex(8)
```

## 字节随机

### Bytes

```go
func Bytes(length int) ([]byte, error)
```

生成指定长度的密码学安全随机字节。

**示例:**

```go
b, err := randomutil.Bytes(32)
```

### MustBytes

```go
func MustBytes(length int) []byte
```

生成随机字节，出错时 panic。

**示例:**

```go
b := randomutil.MustBytes(16)
```

## 切片操作

### Element

```go
func Element[T any](slice []T) (T, bool)
```

从切片中随机选取一个元素。切片为空时返回零值和 false。

**示例:**

```go
colors := []string{"red", "green", "blue"}
color, ok := randomutil.Element(colors)
if ok {
    fmt.Println(color)
}
```

### MustElement

```go
func MustElement[T any](slice []T) T
```

从切片中随机选取一个元素，切片为空时 panic。

**示例:**

```go
color := randomutil.MustElement([]string{"red", "green", "blue"})
```

### Elements

```go
func Elements[T any](slice []T, n int) []T
```

从切片中随机选取 n 个元素（有放回抽样）。

**示例:**

```go
colors := []string{"red", "green", "blue", "yellow"}
selected := randomutil.Elements(colors, 2)
fmt.Println(selected)
```

### Shuffle

```go
func Shuffle[T any](slice []T) []T
```

返回切片的随机排列副本（Fisher-Yates 洗牌算法）。

**示例:**

```go
nums := []int{1, 2, 3, 4, 5}
shuffled := randomutil.Shuffle(nums)
fmt.Println(shuffled)
```

## 加权随机

### Weighted

```go
func Weighted(weights []int) (int, error)
```

根据权重随机选择索引。权重越高，被选中的概率越大。

**示例:**

```go
// 选中索引 0 的概率为 70%，索引 1 为 20%，索引 2 为 10%
idx, err := randomutil.Weighted([]int{70, 20, 10})
fmt.Println(idx)
```

## 骰子/硬币

### Dice

```go
func Dice() int
```

掷一个骰子，返回 1-6。

**示例:**

```go
result := randomutil.Dice()
fmt.Println(result)
```

### DiceN

```go
func DiceN(n int) []int
```

掷 n 个骰子，返回每个骰子的结果（1-6）。

**示例:**

```go
results := randomutil.DiceN(3)
fmt.Println(results) // 如: [4 1 6]
```

### CoinFlip

```go
func CoinFlip() bool
```

抛硬币，返回 true（正面）或 false（反面）。

**示例:**

```go
if randomutil.CoinFlip() {
    fmt.Println("正面")
} else {
    fmt.Println("反面")
}
```

## 颜色

### Color

```go
func Color() string
```

生成随机十六进制颜色字符串（如 `#a3f1c2`）。

**示例:**

```go
color := randomutil.Color()
fmt.Println(color) // #a3f1c2
```

### RGB

```go
func RGB() (r, g, b int)
```

生成随机 RGB 颜色值，每个分量在 [0, 255] 范围内。

**示例:**

```go
r, g, b := randomutil.RGB()
fmt.Printf("rgb(%d, %d, %d)\n", r, g, b)
```

## 日期/时间

### DateBetween

```go
func DateBetween(start, end time.Time) time.Time
```

生成 [start, end) 范围内的随机时间。

**示例:**

```go
start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
randomDate := randomutil.DateBetween(start, end)
fmt.Println(randomDate.Format("2006-01-02"))
```

## 中文模拟数据

### PhoneNumber

```go
func PhoneNumber() string
```

生成随机的中国手机号码（11 位）。

**示例:**

```go
phone := randomutil.PhoneNumber()
fmt.Println(phone) // 13812345678
```

### Email

```go
func Email() string
```

生成随机的电子邮件地址。

**示例:**

```go
email := randomutil.Email()
fmt.Println(email) // aB3xK9mN@example.com
```

### Name

```go
func Name() string
```

生成随机的中文姓名（2-3 个字符）。

**示例:**

```go
name := randomutil.Name()
fmt.Println(name) // 如: 张伟、李芳
```

### Address

```go
func Address() string
```

生成随机的中文地址。

**示例:**

```go
addr := randomutil.Address()
fmt.Println(addr) // 如: 北京市解放路123号
```

## 完整示例

```go
package main

import (
    "fmt"
    "time"

    "github.com/camark/GoHutools/randomutil"
)

func main() {
    // 随机整数
    fmt.Println("Random int:", randomutil.MustInt(100))
    fmt.Println("Random int range:", randomutil.MustIntRange(1, 100))

    // 随机字符串
    fmt.Println("Random string:", randomutil.String(16))
    fmt.Println("Numeric code:", randomutil.Numeric(6))

    // 随机布尔
    fmt.Println("Coin flip:", randomutil.CoinFlip())

    // 骰子
    fmt.Println("Dice:", randomutil.Dice())
    fmt.Println("3 Dice:", randomutil.DiceN(3))

    // 随机颜色
    fmt.Println("Color:", randomutil.Color())

    // 随机日期
    start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
    fmt.Println("Random date:", randomutil.DateBetween(start, end))

    // 切片操作
    colors := []string{"red", "green", "blue", "yellow"}
    fmt.Println("Random color:", randomutil.MustElement(colors))
    fmt.Println("Shuffled:", randomutil.Shuffle(colors))

    // 加权随机
    idx, _ := randomutil.Weighted([]int{70, 20, 10})
    fmt.Println("Weighted index:", idx)

    // 中文模拟数据
    fmt.Println("Name:", randomutil.Name())
    fmt.Println("Phone:", randomutil.PhoneNumber())
    fmt.Println("Email:", randomutil.Email())
    fmt.Println("Address:", randomutil.Address())
}
```
