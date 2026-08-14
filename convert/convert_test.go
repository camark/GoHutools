package convert

import (
	"testing"
)

func TestToStr(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{nil, ""},
		{"hello", "hello"},
		{[]byte("hello"), "hello"},
		{true, "true"},
		{false, "false"},
		{42, "42"},
		{int8(8), "8"},
		{int16(16), "16"},
		{int32(32), "32"},
		{int64(64), "64"},
		{uint(42), "42"},
		{uint8(8), "8"},
		{uint16(16), "16"},
		{uint32(32), "32"},
		{uint64(64), "64"},
		{float32(3.14), "3.14"},
		{float64(3.14159), "3.14159"},
	}

	for _, tt := range tests {
		result := ToStr(tt.input)
		if result != tt.expected {
			t.Errorf("ToStr(%v) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int
		hasError bool
	}{
		{nil, 0, false},
		{42, 42, false},
		{int64(64), 64, false},
		{float64(3.14), 3, false},
		{true, 1, false},
		{false, 0, false},
		{"42", 42, false},
		{"  42  ", 42, false},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		result, err := ToInt(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("ToInt(%v) should have returned error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToInt(%v) returned unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("ToInt(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		}
	}
}

func TestMustToInt(t *testing.T) {
	result := MustToInt(42)
	if result != 42 {
		t.Errorf("MustToInt(42) = %d, want 42", result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("MustToInt('abc') should have panicked")
		}
	}()
	MustToInt("abc")
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int64
		hasError bool
	}{
		{nil, 0, false},
		{42, 42, false},
		{int64(64), 64, false},
		{float64(3.14), 3, false},
		{true, 1, false},
		{"1234567890123", 1234567890123, false},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		result, err := ToInt64(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("ToInt64(%v) should have returned error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToInt64(%v) returned unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("ToInt64(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		}
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
		hasError bool
	}{
		{nil, 0, false},
		{42, 42.0, false},
		{float64(3.14), 3.14, false},
		{true, 1.0, false},
		{false, 0.0, false},
		{"3.14", 3.14, false},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		result, err := ToFloat64(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("ToFloat64(%v) should have returned error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToFloat64(%v) returned unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("ToFloat64(%v) = %f, want %f", tt.input, result, tt.expected)
			}
		}
	}
}

func TestMustToFloat64(t *testing.T) {
	result := MustToFloat64(3.14)
	if result != 3.14 {
		t.Errorf("MustToFloat64(3.14) = %f, want 3.14", result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("MustToFloat64('abc') should have panicked")
		}
	}()
	MustToFloat64("abc")
}

func TestToBool(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
		hasError bool
	}{
		{nil, false, false},
		{true, true, false},
		{false, false, false},
		{1, true, false},
		{0, false, false},
		{42, true, false},
		{"true", true, false},
		{"TRUE", true, false},
		{"1", true, false},
		{"yes", true, false},
		{"on", true, false},
		{"y", true, false},
		{"t", true, false},
		{"false", false, false},
		{"0", false, false},
		{"no", false, false},
		{"off", false, false},
		{"n", false, false},
		{"f", false, false},
		{"", false, false},
		{"abc", false, true},
	}

	for _, tt := range tests {
		result, err := ToBool(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("ToBool(%v) should have returned error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToBool(%v) returned unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		}
	}
}

func TestMustToBool(t *testing.T) {
	result := MustToBool(true)
	if result != true {
		t.Errorf("MustToBool(true) = %v, want true", result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("MustToBool('abc') should have panicked")
		}
	}()
	MustToBool("abc")
}

func TestToBytes(t *testing.T) {
	result := ToBytes("hello")
	expected := []byte("hello")
	if len(result) != len(expected) {
		t.Errorf("ToBytes length = %d, want %d", len(result), len(expected))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("ToBytes[%d] = %d, want %d", i, result[i], expected[i])
		}
	}
}

func TestToString(t *testing.T) {
	result := ToString([]byte("hello"))
	if result != "hello" {
		t.Errorf("ToString = %q, want %q", result, "hello")
	}
}

func TestToIntSlice(t *testing.T) {
	input := []interface{}{1, "2", 3.0, "4", "abc", 5}
	expected := []int{1, 2, 3, 4, 5}
	result := ToIntSlice(input)

	if len(result) != len(expected) {
		t.Errorf("ToIntSlice length = %d, want %d", len(result), len(expected))
		return
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("ToIntSlice[%d] = %d, want %d", i, result[i], expected[i])
		}
	}
}

func TestToStrSlice(t *testing.T) {
	input := []interface{}{1, "hello", true, 3.14}
	expected := []string{"1", "hello", "true", "3.14"}
	result := ToStrSlice(input)

	if len(result) != len(expected) {
		t.Errorf("ToStrSlice length = %d, want %d", len(result), len(expected))
		return
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("ToStrSlice[%d] = %q, want %q", i, result[i], expected[i])
		}
	}
}

func TestCamelToUnderline(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"camelCase", "camel_case"},
		{"CamelCase", "camel_case"},
		{"HTTPServer", "h_t_t_p_server"},
		{"simple", "simple"},
		{"ABC", "a_b_c"},
	}

	for _, tt := range tests {
		result := CamelToUnderline(tt.input)
		if result != tt.expected {
			t.Errorf("CamelToUnderline(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestUnderlineToCamel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"under_score", "UnderScore"},
		{"simple", "Simple"},
		{"a_b_c", "ABC"},
		{"__test__", "Test"},
	}

	for _, tt := range tests {
		result := UnderlineToCamel(tt.input)
		if result != tt.expected {
			t.Errorf("UnderlineToCamel(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestHexToLong(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
		hasError bool
	}{
		{"ff", 255, false},
		{"0xFF", 255, false},
		{"0XFF", 255, false},
		{"  FF  ", 255, false},
		{"1a", 26, false},
		{"", 0, true},
		{"xyz", 0, true},
	}

	for _, tt := range tests {
		result, err := HexToLong(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("HexToLong(%q) should have returned error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("HexToLong(%q) returned unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("HexToLong(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		}
	}
}

func TestLongToHex(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{255, "ff"},
		{26, "1a"},
		{0, "0"},
		{4096, "1000"},
	}

	for _, tt := range tests {
		result := LongToHex(tt.input)
		if result != tt.expected {
			t.Errorf("LongToHex(%d) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestNumberToChinese(t *testing.T) {
	tests := []struct {
		input       int64
		traditional bool
		expected    string
	}{
		{0, false, "零"},
		{0, true, "零"},
		{1, false, "一"},
		{10, false, "一十"},
		{100, false, "一百"},
		{1000, false, "一千"},
		{10000, false, "一万"},
		{12345, false, "一万二千三百四十五"},
		{100000000, false, "一亿"},
		{123456789, false, "一亿二千三百四十五万六千七百八十九"},
		{-1, false, "负一"},
		{1, true, "壹"},
		{10, true, "壹拾"},
		{100, true, "壹佰"},
		{1000, true, "壹仟"},
		{10000, true, "壹萬"},
	}

	for _, tt := range tests {
		result := NumberToChinese(tt.input, tt.traditional)
		if result != tt.expected {
			t.Errorf("NumberToChinese(%d, %v) = %q, want %q", tt.input, tt.traditional, result, tt.expected)
		}
	}
}

func TestChineseToNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
		hasError bool
	}{
		{"零", 0, false},
		{"一", 1, false},
		{"十", 10, false},
		{"一百", 100, false},
		{"一千", 1000, false},
		{"一万", 10000, false},
		{"一万二千三百四十五", 12345, false},
		{"一亿", 100000000, false},
		{"负一", -1, false},
		{"壹", 1, false},
		{"壹拾", 10, false},
		{"壹佰", 100, false},
		{"壹仟", 1000, false},
		{"壹萬", 10000, false},
		{"", 0, true},
		{"负", 0, true},
	}

	for _, tt := range tests {
		result, err := ChineseToNumber(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("ChineseToNumber(%q) should have returned error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("ChineseToNumber(%q) returned unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("ChineseToNumber(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		}
	}
}

func TestDigitToChinese(t *testing.T) {
	expected := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	for i, exp := range expected {
		result := DigitToChinese(byte(i))
		if result != exp {
			t.Errorf("DigitToChinese(%d) = %q, want %q", i, result, exp)
		}
	}

	// Test invalid digit
	result := DigitToChinese(10)
	if result != "" {
		t.Errorf("DigitToChinese(10) = %q, want empty string", result)
	}
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"I", 1, false},
		{"II", 2, false},
		{"III", 3, false},
		{"IV", 4, false},
		{"V", 5, false},
		{"IX", 9, false},
		{"X", 10, false},
		{"XL", 40, false},
		{"L", 50, false},
		{"XC", 90, false},
		{"C", 100, false},
		{"CD", 400, false},
		{"D", 500, false},
		{"CM", 900, false},
		{"M", 1000, false},
		{"MCMXCIX", 1999, false},
		{"MMXXIII", 2023, false},
		{"", 0, true},
		{"A", 0, true},
	}

	for _, tt := range tests {
		result, err := RomanToInt(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("RomanToInt(%q) should have returned error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("RomanToInt(%q) returned unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("RomanToInt(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		}
	}
}

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1999, "MCMXCIX"},
		{2023, "MMXXIII"},
		{3999, "MMMCMXCIX"},
		{0, ""},
		{-1, ""},
		{4000, ""},
	}

	for _, tt := range tests {
		result := IntToRoman(tt.input)
		if result != tt.expected {
			t.Errorf("IntToRoman(%d) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestBaseConvert(t *testing.T) {
	tests := []struct {
		value    string
		fromBase int
		toBase   int
		expected string
		hasError bool
	}{
		{"1010", 2, 10, "10", false},
		{"10", 10, 2, "1010", false},
		{"ff", 16, 10, "255", false},
		{"255", 10, 16, "ff", false},
		{"777", 8, 10, "511", false},
		{"511", 10, 8, "777", false},
		{"10", 2, 16, "2", false},
		{"abc", 16, 10, "2748", false},
		{"10", 1, 10, "", true},   // invalid fromBase
		{"10", 10, 1, "", true},   // invalid toBase
		{"xyz", 10, 16, "", true}, // invalid value
	}

	for _, tt := range tests {
		result, err := BaseConvert(tt.value, tt.fromBase, tt.toBase)
		if tt.hasError {
			if err == nil {
				t.Errorf("BaseConvert(%q, %d, %d) should have returned error", tt.value, tt.fromBase, tt.toBase)
			}
		} else {
			if err != nil {
				t.Errorf("BaseConvert(%q, %d, %d) returned unexpected error: %v", tt.value, tt.fromBase, tt.toBase, err)
			}
			if result != tt.expected {
				t.Errorf("BaseConvert(%q, %d, %d) = %q, want %q", tt.value, tt.fromBase, tt.toBase, result, tt.expected)
			}
		}
	}
}

func TestToEnum(t *testing.T) {
	enumMap := map[string]int{
		"RED":   1,
		"GREEN": 2,
		"BLUE":  3,
	}

	tests := []struct {
		input    string
		expected int
		found    bool
	}{
		{"RED", 1, true},
		{"GREEN", 2, true},
		{"BLUE", 3, true},
		{"YELLOW", 0, false},
		{"", 0, false},
	}

	for _, tt := range tests {
		result, found := ToEnum(tt.input, enumMap)
		if found != tt.found {
			t.Errorf("ToEnum(%q) found = %v, want %v", tt.input, found, tt.found)
		}
		if found && result != tt.expected {
			t.Errorf("ToEnum(%q) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestToMap(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	person := Person{Name: "John", Age: 30}
	result, err := ToMap(person)
	if err != nil {
		t.Errorf("ToMap returned unexpected error: %v", err)
	}

	if result["name"] != "John" {
		t.Errorf("ToMap name = %v, want John", result["name"])
	}
	// JSON numbers are float64
	if result["age"] != float64(30) {
		t.Errorf("ToMap age = %v, want 30", result["age"])
	}

	// Test nil
	_, err = ToMap(nil)
	if err == nil {
		t.Error("ToMap(nil) should have returned error")
	}
}

func TestBase64EncodeDecode(t *testing.T) {
	tests := []string{
		"hello",
		"world",
		"",
		"Hello, World!",
		"12345",
	}

	for _, tt := range tests {
		encoded := Base64Encode(tt)
		decoded, err := Base64Decode(encoded)
		if err != nil {
			t.Errorf("Base64Decode(%q) returned error: %v", encoded, err)
		}
		if decoded != tt {
			t.Errorf("Base64 round-trip failed: %q -> %q -> %q", tt, encoded, decoded)
		}
	}
}

func TestHexEncodeDecode(t *testing.T) {
	tests := [][]byte{
		[]byte("hello"),
		[]byte("world"),
		{},
		{0, 1, 2, 3, 255},
	}

	for _, tt := range tests {
		encoded := HexEncode(tt)
		decoded, err := HexDecode(encoded)
		if err != nil {
			t.Errorf("HexDecode(%q) returned error: %v", encoded, err)
		}
		if len(decoded) != len(tt) {
			t.Errorf("Hex round-trip length mismatch: %d != %d", len(decoded), len(tt))
			continue
		}
		for i := range decoded {
			if decoded[i] != tt[i] {
				t.Errorf("Hex round-trip failed at index %d", i)
			}
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"Hello, World!", "!dlroW ,olleH"},
	}

	for _, tt := range tests {
		result := Reverse(tt.input)
		if result != tt.expected {
			t.Errorf("Reverse(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{42, true},
		{int8(8), true},
		{int16(16), true},
		{int32(32), true},
		{int64(64), true},
		{uint(42), true},
		{uint8(8), true},
		{uint16(16), true},
		{uint32(32), true},
		{uint64(64), true},
		{float32(3.14), true},
		{float64(3.14), true},
		{"42", true},
		{"3.14", true},
		{"abc", false},
		{"", false},
		{true, false},
		{nil, false},
	}

	for _, tt := range tests {
		result := IsNumeric(tt.input)
		if result != tt.expected {
			t.Errorf("IsNumeric(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestTypeOf(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{nil, "nil"},
		{42, "int"},
		{3.14, "float64"},
		{"hello", "string"},
		{true, "bool"},
	}

	for _, tt := range tests {
		result := TypeOf(tt.input)
		if result != tt.expected {
			t.Errorf("TypeOf(%v) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		input    float64
		decimals int
		expected float64
	}{
		{3.14159, 2, 3.14},
		{3.14159, 4, 3.1416},
		{3.14159, 0, 3.0},
		{3.5, 0, 4.0},
		{2.5, 0, 3.0},
	}

	for _, tt := range tests {
		result := Round(tt.input, tt.decimals)
		if result != tt.expected {
			t.Errorf("Round(%f, %d) = %f, want %f", tt.input, tt.decimals, result, tt.expected)
		}
	}
}
