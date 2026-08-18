package numutil

import (
	"math"
	"testing"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		input   string
		want    int
		wantErr bool
	}{
		{"123", 123, false},
		{"-456", -456, false},
		{"0", 0, false},
		{"abc", 0, true},
		{"12.34", 0, true},
		{"", 0, true},
	}
	for _, tt := range tests {
		got, err := ParseInt(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseInt(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseInt(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		input   string
		want    float64
		wantErr bool
	}{
		{"123.456", 123.456, false},
		{"-456.789", -456.789, false},
		{"0", 0, false},
		{"3.14159", 3.14159, false},
		{"abc", 0, true},
		{"", 0, true},
	}
	for _, tt := range tests {
		got, err := ParseFloat(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseFloat(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("ParseFloat(%q) = %f, want %f", tt.input, got, tt.want)
		}
	}
}

func TestMustParseInt(t *testing.T) {
	if got := MustParseInt("123"); got != 123 {
		t.Errorf("MustParseInt(\"123\") = %d, want 123", got)
	}
	if got := MustParseInt("-456"); got != -456 {
		t.Errorf("MustParseInt(\"-456\") = %d, want -456", got)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParseInt(\"abc\") should panic")
		}
	}()
	MustParseInt("abc")
}

func TestMustParseFloat(t *testing.T) {
	if got := MustParseFloat("123.456"); math.Abs(got-123.456) > 1e-10 {
		t.Errorf("MustParseFloat(\"123.456\") = %f, want 123.456", got)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParseFloat(\"abc\") should panic")
		}
	}()
	MustParseFloat("abc")
}

func TestToStr(t *testing.T) {
	tests := []struct {
		input interface{}
		want  string
	}{
		{123, "123"},
		{-456, "-456"},
		{int64(789), "789"},
		{uint(100), "100"},
		{3.14, "3.14"},
		{float32(2.5), "2.5"},
		{true, "true"},
		{false, "false"},
	}
	for _, tt := range tests {
		if got := ToStr(tt.input); got != tt.want {
			t.Errorf("ToStr(%v) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{-1, 1},
	}
	for _, tt := range tests {
		if got := Abs(tt.input); got != tt.want {
			t.Errorf("Abs(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestAbsFloat(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{5.5, 5.5},
		{-5.5, 5.5},
		{0, 0},
		{-3.14, 3.14},
	}
	for _, tt := range tests {
		if got := AbsFloat(tt.input); got != tt.want {
			t.Errorf("AbsFloat(%f) = %f, want %f", tt.input, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{5, 3, 5},
		{3, 5, 5},
		{-1, -5, -1},
		{0, 0, 0},
	}
	for _, tt := range tests {
		if got := Max(tt.a, tt.b); got != tt.want {
			t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMaxFloat(t *testing.T) {
	tests := []struct {
		a, b float64
		want float64
	}{
		{5.5, 3.3, 5.5},
		{3.3, 5.5, 5.5},
		{-1.1, -5.5, -1.1},
	}
	for _, tt := range tests {
		if got := MaxFloat(tt.a, tt.b); got != tt.want {
			t.Errorf("MaxFloat(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{5, 3, 3},
		{3, 5, 3},
		{-1, -5, -5},
		{0, 0, 0},
	}
	for _, tt := range tests {
		if got := Min(tt.a, tt.b); got != tt.want {
			t.Errorf("Min(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMinFloat(t *testing.T) {
	tests := []struct {
		a, b float64
		want float64
	}{
		{5.5, 3.3, 3.3},
		{3.3, 5.5, 3.3},
		{-1.1, -5.5, -5.5},
	}
	for _, tt := range tests {
		if got := MinFloat(tt.a, tt.b); got != tt.want {
			t.Errorf("MinFloat(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		val, min, max int
		want          int
	}{
		{5, 0, 10, 5},
		{-5, 0, 10, 0},
		{15, 0, 10, 10},
		{0, 0, 10, 0},
		{10, 0, 10, 10},
	}
	for _, tt := range tests {
		if got := Clamp(tt.val, tt.min, tt.max); got != tt.want {
			t.Errorf("Clamp(%d, %d, %d) = %d, want %d", tt.val, tt.min, tt.max, got, tt.want)
		}
	}
}

func TestClampFloat(t *testing.T) {
	tests := []struct {
		val, min, max float64
		want          float64
	}{
		{5.5, 0, 10, 5.5},
		{-5.5, 0, 10, 0},
		{15.5, 0, 10, 10},
	}
	for _, tt := range tests {
		if got := ClampFloat(tt.val, tt.min, tt.max); got != tt.want {
			t.Errorf("ClampFloat(%f, %f, %f) = %f, want %f", tt.val, tt.min, tt.max, got, tt.want)
		}
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		val, min, max int
		want          bool
	}{
		{5, 0, 10, true},
		{0, 0, 10, true},
		{10, 0, 10, true},
		{-1, 0, 10, false},
		{11, 0, 10, false},
	}
	for _, tt := range tests {
		if got := Between(tt.val, tt.min, tt.max); got != tt.want {
			t.Errorf("Between(%d, %d, %d) = %v, want %v", tt.val, tt.min, tt.max, got, tt.want)
		}
	}
}

func TestBetweenFloat(t *testing.T) {
	tests := []struct {
		val, min, max float64
		want          bool
	}{
		{5.5, 0, 10, true},
		{0, 0, 10, true},
		{10, 0, 10, true},
		{-0.1, 0, 10, false},
		{10.1, 0, 10, false},
	}
	for _, tt := range tests {
		if got := BetweenFloat(tt.val, tt.min, tt.max); got != tt.want {
			t.Errorf("BetweenFloat(%f, %f, %f) = %v, want %v", tt.val, tt.min, tt.max, got, tt.want)
		}
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{0, true},
		{2, true},
		{4, true},
		{-2, true},
		{1, false},
		{3, false},
		{-1, false},
	}
	for _, tt := range tests {
		if got := IsEven(tt.input); got != tt.want {
			t.Errorf("IsEven(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsOdd(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{0, false},
		{2, false},
		{1, true},
		{3, true},
		{-1, true},
		{-3, true},
	}
	for _, tt := range tests {
		if got := IsOdd(tt.input); got != tt.want {
			t.Errorf("IsOdd(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsPositive(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{1, true},
		{100, true},
		{0, false},
		{-1, false},
	}
	for _, tt := range tests {
		if got := IsPositive(tt.input); got != tt.want {
			t.Errorf("IsPositive(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsNegative(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{-1, true},
		{-100, true},
		{0, false},
		{1, false},
	}
	for _, tt := range tests {
		if got := IsNegative(tt.input); got != tt.want {
			t.Errorf("IsNegative(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{0, true},
		{1, false},
		{-1, false},
	}
	for _, tt := range tests {
		if got := IsZero(tt.input); got != tt.want {
			t.Errorf("IsZero(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{12, 8, 4},
		{8, 12, 4},
		{7, 13, 1},
		{0, 5, 5},
		{5, 0, 5},
		{0, 0, 0},
		{-12, 8, 4},
		{12, -8, 4},
		{math.MinInt64, 0, math.MinInt64},
		{math.MinInt64, 1, 1},
		{math.MinInt64, math.MinInt64, math.MinInt64},
	}
	for _, tt := range tests {
		if got := GCD(tt.a, tt.b); got != tt.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{4, 6, 12},
		{6, 4, 12},
		{3, 5, 15},
		{0, 5, 0},
		{5, 0, 0},
		{1, 1, 1},
		// a*b overflows int64 (16000000024000000000) but real LCM fits: formerly wrong result
		{4000000000, 4000000006, 8000000012000000000},
	}
	for _, tt := range tests {
		if got := LCM(tt.a, tt.b); got != tt.want {
			t.Errorf("LCM(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		input int
		want  int64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{10, 3628800},
	}
	for _, tt := range tests {
		if got := Factorial(tt.input); got != tt.want {
			t.Errorf("Factorial(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestFactorialPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Factorial(-1) should panic")
		}
	}()
	Factorial(-1)
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input int
		want  int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{10, 55},
		{20, 6765},
	}
	for _, tt := range tests {
		if got := Fibonacci(tt.input); got != tt.want {
			t.Errorf("Fibonacci(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestFibonacciPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Fibonacci(-1) should panic")
		}
	}()
	Fibonacci(-1)
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{13, true},
		{15, false},
		{17, true},
		{97, true},
		{100, false},
	}
	for _, tt := range tests {
		if got := IsPrime(tt.input); got != tt.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestPrimes(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{0, []int{}},
		{1, []int{}},
		{2, []int{2}},
		{10, []int{2, 3, 5, 7}},
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}
	for _, tt := range tests {
		got := Primes(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("Primes(%d) = %v, want %v", tt.input, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("Primes(%d) = %v, want %v", tt.input, got, tt.want)
				break
			}
		}
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		f      float64
		places int
		want   float64
	}{
		{3.14159, 2, 3.14},
		{3.14159, 4, 3.1416},
		{3.145, 2, 3.15},
		{3.144, 2, 3.14},
		{-3.14159, 2, -3.14},
		{0.005, 2, 0.01},
		{0.004, 2, 0.00},
	}
	for _, tt := range tests {
		if got := Round(tt.f, tt.places); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Round(%f, %d) = %f, want %f", tt.f, tt.places, got, tt.want)
		}
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		f      float64
		places int
		want   float64
	}{
		{3.141, 2, 3.15},
		{3.149, 2, 3.15},
		{3.140, 2, 3.14},
		{-3.141, 2, -3.14},
	}
	for _, tt := range tests {
		if got := Ceil(tt.f, tt.places); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Ceil(%f, %d) = %f, want %f", tt.f, tt.places, got, tt.want)
		}
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		f      float64
		places int
		want   float64
	}{
		{3.149, 2, 3.14},
		{3.141, 2, 3.14},
		{3.150, 2, 3.15},
		{-3.141, 2, -3.15},
	}
	for _, tt := range tests {
		if got := Floor(tt.f, tt.places); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Floor(%f, %d) = %f, want %f", tt.f, tt.places, got, tt.want)
		}
	}
}

func TestPercent(t *testing.T) {
	tests := []struct {
		part, total int
		want        float64
	}{
		{25, 100, 25.0},
		{1, 3, 33.333333333333336},
		{0, 100, 0.0},
		{100, 100, 100.0},
		{50, 0, 0.0},
	}
	for _, tt := range tests {
		if got := Percent(tt.part, tt.total); math.Abs(got-tt.want) > 1e-6 {
			t.Errorf("Percent(%d, %d) = %f, want %f", tt.part, tt.total, got, tt.want)
		}
	}
}

func TestPercentStr(t *testing.T) {
	tests := []struct {
		part, total, places int
		want                string
	}{
		{25, 100, 2, "25.00%"},
		{1, 3, 2, "33.33%"},
		{0, 100, 0, "0%"},
		{100, 100, 1, "100.0%"},
	}
	for _, tt := range tests {
		if got := PercentStr(tt.part, tt.total, tt.places); got != tt.want {
			t.Errorf("PercentStr(%d, %d, %d) = %q, want %q", tt.part, tt.total, tt.places, got, tt.want)
		}
	}
}

func TestIsNaN(t *testing.T) {
	tests := []struct {
		input float64
		want  bool
	}{
		{1.0, false},
		{0.0, false},
		{math.NaN(), true},
		{math.Inf(1), false},
	}
	for _, tt := range tests {
		if got := IsNaN(tt.input); got != tt.want {
			t.Errorf("IsNaN(%f) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsInf(t *testing.T) {
	tests := []struct {
		input float64
		want  bool
	}{
		{1.0, false},
		{0.0, false},
		{math.Inf(1), true},
		{math.Inf(-1), true},
		{math.NaN(), false},
	}
	for _, tt := range tests {
		if got := IsInf(tt.input); got != tt.want {
			t.Errorf("IsInf(%f) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		a, b    float64
		epsilon float64
		want    bool
	}{
		{1.0, 1.0, 0.0001, true},
		{1.0, 1.0001, 0.0001, true},
		{1.0, 1.001, 0.0001, false},
		{0.0, 0.0, 0, true},
		{1.0, 1.0000000001, 0, true}, // default epsilon
	}
	for _, tt := range tests {
		if got := Equals(tt.a, tt.b, tt.epsilon); got != tt.want {
			t.Errorf("Equals(%f, %f, %f) = %v, want %v", tt.a, tt.b, tt.epsilon, got, tt.want)
		}
	}
}

func TestAddPrecise(t *testing.T) {
	tests := []struct {
		a, b string
		want string
	}{
		{"0.1", "0.2", "0.3"},
		{"1.1", "2.2", "3.3"},
		{"100.123", "200.456", "300.579"},
		{"-1.5", "2.5", "1"},
	}
	for _, tt := range tests {
		got, err := AddPrecise(tt.a, tt.b)
		if err != nil {
			t.Errorf("AddPrecise(%q, %q) error: %v", tt.a, tt.b, err)
			continue
		}
		if got != tt.want {
			t.Errorf("AddPrecise(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestAddPreciseError(t *testing.T) {
	_, err := AddPrecise("abc", "1.0")
	if err == nil {
		t.Error("AddPrecise(\"abc\", \"1.0\") should return error")
	}
	_, err = AddPrecise("1.0", "xyz")
	if err == nil {
		t.Error("AddPrecise(\"1.0\", \"xyz\") should return error")
	}
}

func TestSubPrecise(t *testing.T) {
	tests := []struct {
		a, b string
		want string
	}{
		{"0.3", "0.2", "0.1"},
		{"3.3", "2.2", "1.1"},
		{"300.579", "200.456", "100.123"},
		{"5.5", "2.5", "3"},
	}
	for _, tt := range tests {
		got, err := SubPrecise(tt.a, tt.b)
		if err != nil {
			t.Errorf("SubPrecise(%q, %q) error: %v", tt.a, tt.b, err)
			continue
		}
		if got != tt.want {
			t.Errorf("SubPrecise(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSubPreciseError(t *testing.T) {
	_, err := SubPrecise("abc", "1.0")
	if err == nil {
		t.Error("SubPrecise(\"abc\", \"1.0\") should return error")
	}
}

func TestMulPrecise(t *testing.T) {
	tests := []struct {
		a, b string
		want string
	}{
		{"0.1", "0.2", "0.02"},
		{"1.5", "2.5", "3.75"},
		{"3.14", "2", "6.28"},
		{"-2.5", "4", "-10"},
	}
	for _, tt := range tests {
		got, err := MulPrecise(tt.a, tt.b)
		if err != nil {
			t.Errorf("MulPrecise(%q, %q) error: %v", tt.a, tt.b, err)
			continue
		}
		if got != tt.want {
			t.Errorf("MulPrecise(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMulPreciseError(t *testing.T) {
	_, err := MulPrecise("abc", "1.0")
	if err == nil {
		t.Error("MulPrecise(\"abc\", \"1.0\") should return error")
	}
}

func TestDivPrecise(t *testing.T) {
	tests := []struct {
		a     string
		b     string
		scale int
		want  string
	}{
		{"1", "3", 4, "0.3333"},
		{"10", "3", 2, "3.33"},
		{"100", "4", 0, "25"},
		{"7", "2", 2, "3.50"},
	}
	for _, tt := range tests {
		got, err := DivPrecise(tt.a, tt.b, tt.scale)
		if err != nil {
			t.Errorf("DivPrecise(%q, %q, %d) error: %v", tt.a, tt.b, tt.scale, err)
			continue
		}
		if got != tt.want {
			t.Errorf("DivPrecise(%q, %q, %d) = %q, want %q", tt.a, tt.b, tt.scale, got, tt.want)
		}
	}
}

func TestDivPreciseErrors(t *testing.T) {
	_, err := DivPrecise("abc", "1.0", 2)
	if err == nil {
		t.Error("DivPrecise(\"abc\", \"1.0\", 2) should return error")
	}
	_, err = DivPrecise("1.0", "0", 2)
	if err == nil {
		t.Error("DivPrecise(\"1.0\", \"0\", 2) should return error (division by zero)")
	}
}
