package numutil

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// ParseInt parses string to int
func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ParseFloat parses string to float64
func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// MustParseInt parses string to int, panics on error
func MustParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MustParseFloat parses string to float64, panics on error
func MustParseFloat(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return v
}

// ToStr converts number to string
func ToStr(n interface{}) string {
	switch v := n.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Abs returns absolute value of int
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// AbsFloat returns absolute value of float64
func AbsFloat(n float64) float64 {
	return math.Abs(n)
}

// Max returns maximum of two ints
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxFloat returns maximum of two float64s
func MaxFloat(a, b float64) float64 {
	return math.Max(a, b)
}

// Min returns minimum of two ints
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinFloat returns minimum of two float64s
func MinFloat(a, b float64) float64 {
	return math.Min(a, b)
}

// Clamp clamps int value between min and max
func Clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

// ClampFloat clamps float64 value between min and max
func ClampFloat(val, min, max float64) float64 {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

// Between checks if int value is between min and max (inclusive)
func Between(val, min, max int) bool {
	return val >= min && val <= max
}

// BetweenFloat checks if float64 value is between min and max (inclusive)
func BetweenFloat(val, min, max float64) bool {
	return val >= min && val <= max
}

// IsEven checks if number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd checks if number is odd
func IsOdd(n int) bool {
	return n%2 != 0
}

// IsPositive checks if number is positive
func IsPositive(n int) bool {
	return n > 0
}

// IsNegative checks if number is negative
func IsNegative(n int) bool {
	return n < 0
}

// IsZero checks if number is zero
func IsZero(n int) bool {
	return n == 0
}

// uabs returns the absolute value of x as uint64.
// It uses unsigned arithmetic so that math.MinInt64 does not overflow when negated.
func uabs(x int) uint64 {
	if x < 0 {
		return uint64(0) - uint64(x)
	}
	return uint64(x)
}

// GCD returns greatest common divisor
func GCD(a, b int) int {
	ua := uabs(a)
	ub := uabs(b)
	for ub != 0 {
		ua, ub = ub, ua%ub
	}
	return int(ua)
}

// LCM returns least common multiple
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	g := uint64(GCD(a, b))
	// Divide first to avoid overflow in a*b
	return int(uabs(a) / g * uabs(b))
}

// Factorial returns factorial
func Factorial(n int) int64 {
	if n < 0 {
		panic("factorial of negative number")
	}
	if n == 0 || n == 1 {
		return 1
	}
	var result int64 = 1
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

// Fibonacci returns nth Fibonacci number
func Fibonacci(n int) int64 {
	if n < 0 {
		panic("fibonacci of negative number")
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var a, b int64 = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// IsPrime checks if number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Primes returns primes up to n
func Primes(n int) []int {
	if n < 2 {
		return []int{}
	}
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if sieve[i] {
			for j := i * i; j <= n; j += i {
				sieve[j] = false
			}
		}
	}
	var primes []int
	for i := 2; i <= n; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// Round rounds float to specified decimal places
func Round(f float64, places int) float64 {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return f
	}
	pow := math.Pow(10, float64(places))
	return math.Round(f*pow) / pow
}

// Ceil rounds up to specified decimal places
func Ceil(f float64, places int) float64 {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return f
	}
	pow := math.Pow(10, float64(places))
	return math.Ceil(f*pow) / pow
}

// Floor rounds down to specified decimal places
func Floor(f float64, places int) float64 {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return f
	}
	pow := math.Pow(10, float64(places))
	return math.Floor(f*pow) / pow
}

// Percent calculates percentage
func Percent(part, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(part) / float64(total) * 100
}

// PercentStr returns formatted percentage string
func PercentStr(part, total int, places int) string {
	pct := Percent(part, total)
	return strconv.FormatFloat(Round(pct, places), 'f', places, 64) + "%"
}

// IsNaN checks if float is NaN
func IsNaN(f float64) bool {
	return math.IsNaN(f)
}

// IsInf checks if float is infinite
func IsInf(f float64) bool {
	return math.IsInf(f, 0)
}

// Equals checks if two floats are approximately equal
func Equals(a, b float64, epsilon float64) bool {
	if epsilon <= 0 {
		epsilon = 1e-9
	}
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= epsilon
}

// decimalPlaces counts the number of decimal places in a numeric string.
func decimalPlaces(s string) int {
	if dot := strings.Index(s, "."); dot >= 0 {
		return len(s) - dot - 1
	}
	return 0
}

// trimTrailingZeros removes trailing zeros after the decimal point,
// and removes the decimal point itself if no fractional part remains.
func trimTrailingZeros(s string) string {
	if dot := strings.Index(s, "."); dot >= 0 {
		// Trim trailing zeros
		i := len(s) - 1
		for i > dot && s[i] == '0' {
			i--
		}
		// If all zeros after dot, remove dot too
		if i == dot {
			return s[:dot]
		}
		return s[:i+1]
	}
	return s
}

// AddPrecise adds two float strings precisely
func AddPrecise(a, b string) (string, error) {
	n1, _, err := big.ParseFloat(a, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", fmt.Errorf("invalid number %q: %w", a, err)
	}
	n2, _, err := big.ParseFloat(b, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", fmt.Errorf("invalid number %q: %w", b, err)
	}
	result := new(big.Float).Add(n1, n2)
	scale := max(decimalPlaces(a), decimalPlaces(b))
	return trimTrailingZeros(result.Text('f', scale)), nil
}

// SubPrecise subtracts two float strings precisely
func SubPrecise(a, b string) (string, error) {
	n1, _, err := big.ParseFloat(a, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", fmt.Errorf("invalid number %q: %w", a, err)
	}
	n2, _, err := big.ParseFloat(b, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", fmt.Errorf("invalid number %q: %w", b, err)
	}
	result := new(big.Float).Sub(n1, n2)
	scale := max(decimalPlaces(a), decimalPlaces(b))
	return trimTrailingZeros(result.Text('f', scale)), nil
}

// MulPrecise multiplies two float strings precisely
func MulPrecise(a, b string) (string, error) {
	n1, _, err := big.ParseFloat(a, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", fmt.Errorf("invalid number %q: %w", a, err)
	}
	n2, _, err := big.ParseFloat(b, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", fmt.Errorf("invalid number %q: %w", b, err)
	}
	result := new(big.Float).Mul(n1, n2)
	scale := decimalPlaces(a) + decimalPlaces(b)
	return trimTrailingZeros(result.Text('f', scale)), nil
}

// DivPrecise divides two float strings precisely
func DivPrecise(a, b string, scale int) (string, error) {
	n1, _, err := big.ParseFloat(a, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", fmt.Errorf("invalid number %q: %w", a, err)
	}
	n2, _, err := big.ParseFloat(b, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", fmt.Errorf("invalid number %q: %w", b, err)
	}
	if n2.Sign() == 0 {
		return "", errors.New("division by zero")
	}
	result := new(big.Float).Quo(n1, n2)
	if scale >= 0 {
		return result.Text('f', scale), nil
	}
	return result.Text('f', -1), nil
}
