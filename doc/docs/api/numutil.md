---
title: numutil - 数值工具
---

# numutil - 数值工具

数值计算工具包，提供了数值解析、数学运算、精确计算、数值判断等功能。

## 导入

```go
import "github.com/gongm/gohutool/numutil"
```

## 函数列表

### ParseInt

```go
func ParseInt(s string) (int, error)
```

解析字符串为 int 类型。

**示例:**

```go
n, err := numutil.ParseInt("123")
fmt.Println(n, err) // 123 <nil>
```

### ParseFloat

```go
func ParseFloat(s string) (float64, error)
```

解析字符串为 float64 类型。

**示例:**

```go
f, err := numutil.ParseFloat("3.14")
fmt.Println(f, err) // 3.14 <nil>
```

### MustParseInt

```go
func MustParseInt(s string) int
```

解析字符串为 int，失败时 panic。

**示例:**

```go
n := numutil.MustParseInt("123")
fmt.Println(n) // 123
```

### MustParseFloat

```go
func MustParseFloat(s string) float64
```

解析字符串为 float64，失败时 panic。

**示例:**

```go
f := numutil.MustParseFloat("3.14")
fmt.Println(f) // 3.14
```

### ToStr

```go
func ToStr(n interface{}) string
```

将数值转换为字符串。

**示例:**

```go
fmt.Println(numutil.ToStr(123))     // "123"
fmt.Println(numutil.ToStr(3.14))    // "3.14"
fmt.Println(numutil.ToStr(true))    // "true"
```

### Abs

```go
func Abs(n int) int
```

返回 int 的绝对值。

**示例:**

```go
fmt.Println(numutil.Abs(-5))  // 5
fmt.Println(numutil.Abs(5))   // 5
```

### AbsFloat

```go
func AbsFloat(n float64) float64
```

返回 float64 的绝对值。

**示例:**

```go
fmt.Println(numutil.AbsFloat(-3.14))  // 3.14
fmt.Println(numutil.AbsFloat(3.14))   // 3.14
```

### Max

```go
func Max(a, b int) int
```

返回两个 int 中的最大值。

**示例:**

```go
fmt.Println(numutil.Max(3, 5))  // 5
```

### MaxFloat

```go
func MaxFloat(a, b float64) float64
```

返回两个 float64 中的最大值。

**示例:**

```go
fmt.Println(numutil.MaxFloat(3.14, 2.71))  // 3.14
```

### Min

```go
func Min(a, b int) int
```

返回两个 int 中的最小值。

**示例:**

```go
fmt.Println(numutil.Min(3, 5))  // 3
```

### MinFloat

```go
func MinFloat(a, b float64) float64
```

返回两个 float64 中的最小值。

**示例:**

```go
fmt.Println(numutil.MinFloat(3.14, 2.71))  // 2.71
```

### Clamp

```go
func Clamp(val, min, max int) int
```

将 int 值限制在指定范围内。

**示例:**

```go
fmt.Println(numutil.Clamp(15, 0, 10))  // 10
fmt.Println(numutil.Clamp(5, 0, 10))   // 5
fmt.Println(numutil.Clamp(-5, 0, 10))  // 0
```

### ClampFloat

```go
func ClampFloat(val, min, max float64) float64
```

将 float64 值限制在指定范围内。

**示例:**

```go
fmt.Println(numutil.ClampFloat(15.5, 0, 10))  // 10
```

### Between

```go
func Between(val, min, max int) bool
```

检查 int 值是否在范围内（包含边界）。

**示例:**

```go
fmt.Println(numutil.Between(5, 1, 10))   // true
fmt.Println(numutil.Between(15, 1, 10))  // false
```

### BetweenFloat

```go
func BetweenFloat(val, min, max float64) bool
```

检查 float64 值是否在范围内（包含边界）。

**示例:**

```go
fmt.Println(numutil.BetweenFloat(5.5, 1.0, 10.0))  // true
```

### IsEven

```go
func IsEven(n int) bool
```

检查数字是否为偶数。

**示例:**

```go
fmt.Println(numutil.IsEven(4))  // true
fmt.Println(numutil.IsEven(3))  // false
```

### IsOdd

```go
func IsOdd(n int) bool
```

检查数字是否为奇数。

**示例:**

```go
fmt.Println(numutil.IsOdd(3))  // true
fmt.Println(numutil.IsOdd(4))  // false
```

### IsPositive

```go
func IsPositive(n int) bool
```

检查数字是否为正数。

**示例:**

```go
fmt.Println(numutil.IsPositive(5))   // true
fmt.Println(numutil.IsPositive(-5))  // false
```

### IsNegative

```go
func IsNegative(n int) bool
```

检查数字是否为负数。

**示例:**

```go
fmt.Println(numutil.IsNegative(-5))  // true
fmt.Println(numutil.IsNegative(5))   // false
```

### IsZero

```go
func IsZero(n int) bool
```

检查数字是否为零。

**示例:**

```go
fmt.Println(numutil.IsZero(0))  // true
fmt.Println(numutil.IsZero(5))  // false
```

### GCD

```go
func GCD(a, b int) int
```

计算最大公约数。

**示例:**

```go
fmt.Println(numutil.GCD(12, 8))  // 4
fmt.Println(numutil.GCD(15, 10)) // 5
```

### LCM

```go
func LCM(a, b int) int
```

计算最小公倍数。

**示例:**

```go
fmt.Println(numutil.LCM(4, 6))   // 12
fmt.Println(numutil.LCM(3, 5))   // 15
```

### Factorial

```go
func Factorial(n int) int64
```

计算阶乘。

**示例:**

```go
fmt.Println(numutil.Factorial(5))   // 120
fmt.Println(numutil.Factorial(0))   // 1
```

### Fibonacci

```go
func Fibonacci(n int) int64
```

返回第 n 个斐波那契数。

**示例:**

```go
fmt.Println(numutil.Fibonacci(0))   // 0
fmt.Println(numutil.Fibonacci(1))   // 1
fmt.Println(numutil.Fibonacci(10))  // 55
```

### IsPrime

```go
func IsPrime(n int) bool
```

检查数字是否为素数。

**示例:**

```go
fmt.Println(numutil.IsPrime(7))   // true
fmt.Println(numutil.IsPrime(10))  // false
```

### Primes

```go
func Primes(n int) []int
```

返回小于等于 n 的所有素数。

**示例:**

```go
fmt.Println(numutil.Primes(20))  // [2 3 5 7 11 13 17 19]
```

### Round

```go
func Round(f float64, places int) float64
```

四舍五入到指定小数位。

**示例:**

```go
fmt.Println(numutil.Round(3.14159, 2))  // 3.14
fmt.Println(numutil.Round(3.14659, 2))  // 3.15
```

### Ceil

```go
func Ceil(f float64, places int) float64
```

向上取整到指定小数位。

**示例:**

```go
fmt.Println(numutil.Ceil(3.14159, 2))  // 3.15
```

### Floor

```go
func Floor(f float64, places int) float64
```

向下取整到指定小数位。

**示例:**

```go
fmt.Println(numutil.Floor(3.14159, 2))  // 3.14
```

### Percent

```go
func Percent(part, total int) float64
```

计算百分比。

**示例:**

```go
fmt.Println(numutil.Percent(25, 100))  // 25
fmt.Println(numutil.Percent(1, 3))     // 33.333333333333336
```

### PercentStr

```go
func PercentStr(part, total int, places int) string
```

返回格式化的百分比字符串。

**示例:**

```go
fmt.Println(numutil.PercentStr(25, 100, 0))  // "25%"
fmt.Println(numutil.PercentStr(1, 3, 2))     // "33.33%"
```

### IsNaN

```go
func IsNaN(f float64) bool
```

检查浮点数是否为 NaN。

**示例:**

```go
fmt.Println(numutil.IsNaN(math.NaN()))  // true
fmt.Println(numutil.IsNaN(3.14))        // false
```

### IsInf

```go
func IsInf(f float64) bool
```

检查浮点数是否为无穷大。

**示例:**

```go
fmt.Println(numutil.IsInf(math.Inf(1)))  // true
fmt.Println(numutil.IsInf(3.14))         // false
```

### Equals

```go
func Equals(a, b float64, epsilon float64) bool
```

检查两个浮点数是否近似相等。

**示例:**

```go
fmt.Println(numutil.Equals(0.1+0.2, 0.3, 1e-9))  // true
fmt.Println(numutil.Equals(1.0, 1.1, 0.01))       // false
```

### AddPrecise

```go
func AddPrecise(a, b string) (string, error)
```

精确加法（使用 big.Float）。

**示例:**

```go
result, _ := numutil.AddPrecise("0.1", "0.2")
fmt.Println(result)  // "0.3"
```

### SubPrecise

```go
func SubPrecise(a, b string) (string, error)
```

精确减法（使用 big.Float）。

**示例:**

```go
result, _ := numutil.SubPrecise("0.3", "0.1")
fmt.Println(result)  // "0.2"
```

### MulPrecise

```go
func MulPrecise(a, b string) (string, error)
```

精确乘法（使用 big.Float）。

**示例:**

```go
result, _ := numutil.MulPrecise("0.1", "0.2")
fmt.Println(result)  // "0.02"
```

### DivPrecise

```go
func DivPrecise(a, b string, scale int) (string, error)
```

精确除法（使用 big.Float）。

**示例:**

```go
result, _ := numutil.DivPrecise("1", "3", 4)
fmt.Println(result)  // "0.3333"
```
