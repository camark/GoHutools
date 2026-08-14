package randomutil

import (
	"encoding/hex"
	"regexp"
	"strings"
	"testing"
	"time"
)

// ---------------------------------------------------------------------------
// Int / IntRange
// ---------------------------------------------------------------------------

func TestInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Int(10)
		if err != nil {
			t.Fatal(err)
		}
		if n < 0 || n >= 10 {
			t.Fatalf("Int(10) = %d, want [0,10)", n)
		}
	}
}

func TestIntZero(t *testing.T) {
	n, err := Int(0)
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("Int(0) = %d, want 0", n)
	}
}

func TestIntRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := IntRange(5, 15)
		if err != nil {
			t.Fatal(err)
		}
		if n < 5 || n >= 15 {
			t.Fatalf("IntRange(5,15) = %d, want [5,15)", n)
		}
	}
}

func TestIntRangeInvalid(t *testing.T) {
	n, err := IntRange(10, 5)
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("IntRange(10,5) = %d, want 0", n)
	}
}

func TestMustInt(t *testing.T) {
	n := MustInt(100)
	if n < 0 || n >= 100 {
		t.Fatalf("MustInt(100) = %d, want [0,100)", n)
	}
}

func TestMustIntRange(t *testing.T) {
	n := MustIntRange(50, 60)
	if n < 50 || n >= 60 {
		t.Fatalf("MustIntRange(50,60) = %d, want [50,60)", n)
	}
}

// ---------------------------------------------------------------------------
// Int64 / Int64Range
// ---------------------------------------------------------------------------

func TestInt64(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Int64(1000)
		if err != nil {
			t.Fatal(err)
		}
		if n < 0 || n >= 1000 {
			t.Fatalf("Int64(1000) = %d, want [0,1000)", n)
		}
	}
}

func TestInt64Range(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Int64Range(100, 200)
		if err != nil {
			t.Fatal(err)
		}
		if n < 100 || n >= 200 {
			t.Fatalf("Int64Range(100,200) = %d, want [100,200)", n)
		}
	}
}

func TestInt64RangeInvalid(t *testing.T) {
	n, err := Int64Range(200, 100)
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("Int64Range(200,100) = %d, want 0", n)
	}
}

// ---------------------------------------------------------------------------
// Float64 / Float64Range
// ---------------------------------------------------------------------------

func TestFloat64(t *testing.T) {
	for i := 0; i < 100; i++ {
		f, err := Float64()
		if err != nil {
			t.Fatal(err)
		}
		if f < 0.0 || f >= 1.0 {
			t.Fatalf("Float64() = %f, want [0.0,1.0)", f)
		}
	}
}

func TestFloat64Range(t *testing.T) {
	for i := 0; i < 100; i++ {
		f, err := Float64Range(2.5, 7.5)
		if err != nil {
			t.Fatal(err)
		}
		if f < 2.5 || f >= 7.5 {
			t.Fatalf("Float64Range(2.5,7.5) = %f, want [2.5,7.5)", f)
		}
	}
}

func TestFloat64RangeInvalid(t *testing.T) {
	f, err := Float64Range(7.5, 2.5)
	if err != nil {
		t.Fatal(err)
	}
	if f != 7.5 {
		t.Fatalf("Float64Range(7.5,2.5) = %f, want 7.5", f)
	}
}

// ---------------------------------------------------------------------------
// Bool
// ---------------------------------------------------------------------------

func TestBool(t *testing.T) {
	trueCount := 0
	for i := 0; i < 1000; i++ {
		b, err := Bool()
		if err != nil {
			t.Fatal(err)
		}
		if b {
			trueCount++
		}
	}
	// With 1000 samples, expect roughly 500 true values; allow wide margin
	if trueCount < 300 || trueCount > 700 {
		t.Errorf("unexpected distribution: %d/1000 true", trueCount)
	}
}

func TestMustBool(t *testing.T) {
	_ = MustBool() // should not panic
}

// ---------------------------------------------------------------------------
// String variants
// ---------------------------------------------------------------------------

func TestString(t *testing.T) {
	s := String(20)
	if len(s) != 20 {
		t.Fatalf("expected 20 chars, got %d", len(s))
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(s) {
		t.Errorf("String contains unexpected chars: %s", s)
	}
}

func TestStringZero(t *testing.T) {
	s := String(0)
	if s != "" {
		t.Errorf("expected empty string, got %q", s)
	}
}

func TestAlpha(t *testing.T) {
	s := Alpha(16)
	if len(s) != 16 {
		t.Fatalf("expected 16 chars, got %d", len(s))
	}
	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !re.MatchString(s) {
		t.Errorf("Alpha contains non-alpha chars: %s", s)
	}
}

func TestAlphaNumeric(t *testing.T) {
	s := AlphaNumeric(32)
	if len(s) != 32 {
		t.Fatalf("expected 32 chars, got %d", len(s))
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(s) {
		t.Errorf("AlphaNumeric contains unexpected chars: %s", s)
	}
}

func TestNumeric(t *testing.T) {
	s := Numeric(10)
	if len(s) != 10 {
		t.Fatalf("expected 10 chars, got %d", len(s))
	}
	re := regexp.MustCompile(`^[0-9]+$`)
	if !re.MatchString(s) {
		t.Errorf("Numeric contains non-digit chars: %s", s)
	}
}

func TestHex(t *testing.T) {
	s := Hex(16)
	if len(s) != 16 {
		t.Fatalf("expected 16 chars, got %d", len(s))
	}
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Hex produced invalid hex: %s", s)
	}
}

func TestStringWithAlphabet(t *testing.T) {
	alphabet := "XYZ"
	s := StringWithAlphabet(30, alphabet)
	if len(s) != 30 {
		t.Fatalf("expected 30 chars, got %d", len(s))
	}
	for _, c := range s {
		if c != 'X' && c != 'Y' && c != 'Z' {
			t.Errorf("unexpected char: %c", c)
		}
	}
}

func TestStringWithAlphabetEmpty(t *testing.T) {
	s := StringWithAlphabet(10, "")
	if s != "" {
		t.Errorf("expected empty string, got %q", s)
	}
}

// ---------------------------------------------------------------------------
// Bytes
// ---------------------------------------------------------------------------

func TestBytes(t *testing.T) {
	b, err := Bytes(32)
	if err != nil {
		t.Fatal(err)
	}
	if len(b) != 32 {
		t.Fatalf("expected 32 bytes, got %d", len(b))
	}
}

func TestBytesZero(t *testing.T) {
	b, err := Bytes(0)
	if err != nil {
		t.Fatal(err)
	}
	if b != nil {
		t.Errorf("expected nil for length 0, got %v", b)
	}
}

func TestMustBytes(t *testing.T) {
	b := MustBytes(16)
	if len(b) != 16 {
		t.Fatalf("expected 16 bytes, got %d", len(b))
	}
}

// ---------------------------------------------------------------------------
// Element / MustElement
// ---------------------------------------------------------------------------

func TestElement(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	for i := 0; i < 100; i++ {
		v, ok := Element(slice)
		if !ok {
			t.Fatal("Element returned false for non-empty slice")
		}
		found := false
		for _, s := range slice {
			if s == v {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Element returned %d, not in slice", v)
		}
	}
}

func TestElementEmpty(t *testing.T) {
	v, ok := Element([]int{})
	if ok {
		t.Error("expected false for empty slice")
	}
	if v != 0 {
		t.Errorf("expected zero value, got %d", v)
	}
}

func TestMustElement(t *testing.T) {
	slice := []string{"a", "b", "c"}
	v := MustElement(slice)
	if v != "a" && v != "b" && v != "c" {
		t.Errorf("unexpected element: %s", v)
	}
}

func TestMustElementEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for empty slice")
		}
	}()
	MustElement([]int{})
}

// ---------------------------------------------------------------------------
// Elements
// ---------------------------------------------------------------------------

func TestElements(t *testing.T) {
	slice := []int{1, 2, 3}
	result := Elements(slice, 5)
	if len(result) != 5 {
		t.Fatalf("expected 5 elements, got %d", len(result))
	}
	for _, v := range result {
		if v != 1 && v != 2 && v != 3 {
			t.Errorf("unexpected element: %d", v)
		}
	}
}

func TestElementsZeroN(t *testing.T) {
	result := Elements([]int{1, 2, 3}, 0)
	if result != nil {
		t.Errorf("expected nil for n=0, got %v", result)
	}
}

func TestElementsNegativeN(t *testing.T) {
	result := Elements([]int{1, 2, 3}, -1)
	if result != nil {
		t.Errorf("expected nil for n=-1, got %v", result)
	}
}

func TestElementsEmptySlice(t *testing.T) {
	result := Elements([]int{}, 5)
	if result != nil {
		t.Errorf("expected nil for empty slice, got %v", result)
	}
}

// ---------------------------------------------------------------------------
// Shuffle
// ---------------------------------------------------------------------------

func TestShuffle(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := Shuffle(slice)
	if len(result) != len(slice) {
		t.Fatalf("expected %d elements, got %d", len(slice), len(result))
	}
	// Verify original is not modified
	if slice[0] != 1 {
		t.Error("original slice was modified")
	}
	// Verify all elements present
	sum := 0
	for _, v := range result {
		sum += v
	}
	if sum != 55 {
		t.Errorf("expected sum 55, got %d", sum)
	}
}

func TestShuffleEmpty(t *testing.T) {
	result := Shuffle([]int{})
	if result != nil {
		t.Errorf("expected nil for empty slice, got %v", result)
	}
}

// ---------------------------------------------------------------------------
// Weighted
// ---------------------------------------------------------------------------

func TestWeighted(t *testing.T) {
	weights := []int{1, 0, 0, 0, 0}
	// With only first element having weight, should always return 0
	for i := 0; i < 100; i++ {
		idx, err := Weighted(weights)
		if err != nil {
			t.Fatal(err)
		}
		if idx != 0 {
			t.Fatalf("expected 0, got %d", idx)
		}
	}
}

func TestWeightedDistribution(t *testing.T) {
	weights := []int{80, 20}
	counts := make([]int, 2)
	for i := 0; i < 10000; i++ {
		idx, err := Weighted(weights)
		if err != nil {
			t.Fatal(err)
		}
		if idx < 0 || idx >= 2 {
			t.Fatalf("index out of range: %d", idx)
		}
		counts[idx]++
	}
	// Expect roughly 8000 and 2000
	if counts[0] < 7500 || counts[0] > 8500 {
		t.Errorf("unexpected distribution: %v", counts)
	}
}

func TestWeightedEmpty(t *testing.T) {
	idx, err := Weighted([]int{})
	if err != nil {
		t.Fatal(err)
	}
	if idx != 0 {
		t.Errorf("expected 0 for empty weights, got %d", idx)
	}
}

func TestWeightedAllZero(t *testing.T) {
	idx, err := Weighted([]int{0, 0, 0})
	if err != nil {
		t.Fatal(err)
	}
	if idx != 0 {
		t.Errorf("expected 0 for all-zero weights, got %d", idx)
	}
}

// ---------------------------------------------------------------------------
// Dice / CoinFlip
// ---------------------------------------------------------------------------

func TestDice(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := Dice()
		if v < 1 || v > 6 {
			t.Fatalf("Dice() = %d, want [1,6]", v)
		}
	}
}

func TestDiceN(t *testing.T) {
	results := DiceN(5)
	if len(results) != 5 {
		t.Fatalf("expected 5 dice, got %d", len(results))
	}
	for _, v := range results {
		if v < 1 || v > 6 {
			t.Errorf("DiceN value %d out of range", v)
		}
	}
}

func TestDiceNZero(t *testing.T) {
	results := DiceN(0)
	if results != nil {
		t.Errorf("expected nil for n=0, got %v", results)
	}
}

func TestCoinFlip(t *testing.T) {
	_ = CoinFlip() // should not panic
}

// ---------------------------------------------------------------------------
// Color / RGB
// ---------------------------------------------------------------------------

func TestColor(t *testing.T) {
	c := Color()
	if len(c) != 7 || c[0] != '#' {
		t.Errorf("unexpected color format: %s", c)
	}
	_, err := hex.DecodeString(c[1:])
	if err != nil {
		t.Errorf("color is not valid hex: %s", c)
	}
}

func TestRGB(t *testing.T) {
	r, g, b := RGB()
	if r < 0 || r > 255 {
		t.Errorf("r out of range: %d", r)
	}
	if g < 0 || g > 255 {
		t.Errorf("g out of range: %d", g)
	}
	if b < 0 || b > 255 {
		t.Errorf("b out of range: %d", b)
	}
}

// ---------------------------------------------------------------------------
// DateBetween
// ---------------------------------------------------------------------------

func TestDateBetween(t *testing.T) {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC)
	for i := 0; i < 100; i++ {
		d := DateBetween(start, end)
		if d.Before(start) || !d.Before(end) {
			t.Errorf("DateBetween returned %v, want [%v, %v)", d, start, end)
		}
	}
}

func TestDateBetweenReversed(t *testing.T) {
	start := time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	d := DateBetween(start, end)
	// Should swap and still return a valid date
	if d.Before(end) || d.After(start) {
		t.Errorf("DateBetween with reversed args returned %v", d)
	}
}

func TestDateBetweenSame(t *testing.T) {
	ts := time.Date(2020, 6, 15, 12, 0, 0, 0, time.UTC)
	d := DateBetween(ts, ts)
	if !d.Equal(ts) {
		t.Errorf("expected %v, got %v", ts, d)
	}
}

// ---------------------------------------------------------------------------
// Chinese mock data
// ---------------------------------------------------------------------------

func TestPhoneNumber(t *testing.T) {
	for i := 0; i < 50; i++ {
		p := PhoneNumber()
		if len(p) != 11 {
			t.Fatalf("expected 11 digits, got %d: %s", len(p), p)
		}
		re := regexp.MustCompile(`^1[3-9]\d{9}$`)
		if !re.MatchString(p) {
			t.Errorf("invalid phone number: %s", p)
		}
	}
}

func TestEmail(t *testing.T) {
	for i := 0; i < 50; i++ {
		e := Email()
		if !strings.Contains(e, "@") {
			t.Errorf("invalid email: %s", e)
		}
		parts := strings.Split(e, "@")
		if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
			t.Errorf("invalid email: %s", e)
		}
	}
}

func TestName(t *testing.T) {
	for i := 0; i < 50; i++ {
		n := Name()
		runes := []rune(n)
		if len(runes) < 2 || len(runes) > 3 {
			t.Errorf("expected 2-3 char name, got %d chars: %s", len(runes), n)
		}
	}
}

func TestAddress(t *testing.T) {
	for i := 0; i < 50; i++ {
		a := Address()
		if len(a) == 0 {
			t.Error("expected non-empty address")
		}
		if !strings.HasSuffix(a, "号") {
			t.Errorf("expected address to end with '号': %s", a)
		}
	}
}

// ---------------------------------------------------------------------------
// Uniqueness / Distribution smoke tests
// ---------------------------------------------------------------------------

func TestStringUniqueness(t *testing.T) {
	seen := make(map[string]bool, 1000)
	for i := 0; i < 1000; i++ {
		s := String(16)
		if seen[s] {
			t.Fatalf("duplicate string: %s", s)
		}
		seen[s] = true
	}
}

func TestIntDistribution(t *testing.T) {
	counts := make([]int, 10)
	for i := 0; i < 10000; i++ {
		n := MustInt(10)
		counts[n]++
	}
	// Each bucket should get roughly 1000; allow wide margin
	for i, c := range counts {
		if c < 700 || c > 1300 {
			t.Errorf("bucket %d: count %d is outside expected range", i, c)
		}
	}
}

func TestInt64Distribution(t *testing.T) {
	counts := make([]int, 10)
	for i := 0; i < 10000; i++ {
		n, err := Int64(10)
		if err != nil {
			t.Fatal(err)
		}
		counts[n]++
	}
	for i, c := range counts {
		if c < 700 || c > 1300 {
			t.Errorf("bucket %d: count %d is outside expected range", i, c)
		}
	}
}

// ---------------------------------------------------------------------------
// Benchmarks
// ---------------------------------------------------------------------------

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(1000)
	}
}

func BenchmarkIntRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntRange(100, 200)
	}
}

func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64()
	}
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool()
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(32)
	}
}

func BenchmarkAlphaNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphaNumeric(32)
	}
}

func BenchmarkHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hex(32)
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(32)
	}
}

func BenchmarkElement(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Element(slice)
	}
}

func BenchmarkShuffle(b *testing.B) {
	slice := make([]int, 100)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Shuffle(slice)
	}
}

func BenchmarkWeighted(b *testing.B) {
	weights := []int{10, 20, 30, 40}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Weighted(weights)
	}
}

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Color()
	}
}

func BenchmarkPhoneNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneNumber()
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name()
	}
}
