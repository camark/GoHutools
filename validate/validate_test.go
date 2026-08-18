package validate

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{nil, true},
		{"", true},
		{"hello", false},
		{[]byte{}, true},
		{[]byte("hello"), false},
		{[]interface{}{}, true},
		{[]interface{}{1}, false},
		{map[string]interface{}{}, true},
		{map[string]interface{}{"a": 1}, false},
		{[]string{}, true},
		{map[string]int{}, true},
		{0, true},
		{42, false},
		{int8(0), true},
		{int8(1), false},
		{int16(0), true},
		{int16(1), false},
		{int32(0), true},
		{int32(1), false},
		{int64(0), true},
		{int64(1), false},
		{uint(0), true},
		{uint(1), false},
		{uint8(0), true},
		{uint8(1), false},
		{uint16(0), true},
		{uint16(1), false},
		{uint32(0), true},
		{uint32(1), false},
		{uint64(0), true},
		{uint64(1), false},
		{float32(0), true},
		{float32(1.1), false},
		{float64(0), true},
		{float64(1.1), false},
		{false, true},
		{true, false},
	}

	for _, tt := range tests {
		result := IsEmpty(tt.input)
		if result != tt.expected {
			t.Errorf("IsEmpty(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{nil, false},
		{"", false},
		{"hello", true},
		{0, false},
		{42, true},
	}

	for _, tt := range tests {
		result := IsNotEmpty(tt.input)
		if result != tt.expected {
			t.Errorf("IsNotEmpty(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsNull(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{nil, true},
		{"", false},
		{0, false},
		{false, false},
	}

	for _, tt := range tests {
		result := IsNull(tt.input)
		if result != tt.expected {
			t.Errorf("IsNull(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsNotNull(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{nil, false},
		{"", true},
		{0, true},
		{false, true},
	}

	for _, tt := range tests {
		result := IsNotNull(tt.input)
		if result != tt.expected {
			t.Errorf("IsNotNull(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"   ", true},
		{"\t\n", true},
		{"hello", false},
		{" hello ", false},
	}

	for _, tt := range tests {
		result := IsBlank(tt.input)
		if result != tt.expected {
			t.Errorf("IsBlank(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsNotBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"   ", false},
		{"hello", true},
	}

	for _, tt := range tests {
		result := IsNotBlank(tt.input)
		if result != tt.expected {
			t.Errorf("IsNotBlank(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name@domain.com", true},
		{"user+tag@domain.com", true},
		{"user@subdomain.domain.com", true},
		{"", false},
		{"test", false},
		{"test@", false},
		{"@domain.com", false},
		{"test@.com", false},
		{"test@domain.", false},
		{"test@domain.c", false},
	}

	for _, tt := range tests {
		result := IsEmail(tt.input)
		if result != tt.expected {
			t.Errorf("IsEmail(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"http://example.com", true},
		{"https://example.com", true},
		{"https://www.example.com/path?query=1", true},
		{"ftp://files.example.com", true},
		{"http://x", true},
		{"http://127.0.0.1:8080", true},
		{"", false},
		{"example.com", false},
		{"http://", false},
		{"://example.com", false},
	}

	for _, tt := range tests {
		result := IsURL(tt.input)
		if result != tt.expected {
			t.Errorf("IsURL(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsIP(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"192.168.1.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"::1", true},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"", false},
		{"256.1.1.1", false},
		{"1.2.3", false},
		{"abc", false},
	}

	for _, tt := range tests {
		result := IsIP(tt.input)
		if result != tt.expected {
			t.Errorf("IsIP(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"192.168.1.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"10.0.0.1", true},
		{"", false},
		{"256.1.1.1", false},
		{"::1", false},
		{"2001:db8::1", false},
		{"abc", false},
	}

	for _, tt := range tests {
		result := IsIPv4(tt.input)
		if result != tt.expected {
			t.Errorf("IsIPv4(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"::1", true},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"fe80::1", true},
		{"::", true},
		{"", false},
		{"192.168.1.1", false},
		{"abc", false},
	}

	for _, tt := range tests {
		result := IsIPv6(tt.input)
		if result != tt.expected {
			t.Errorf("IsIPv6(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsMAC(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"00:1B:44:11:3A:B7", true},
		{"00-1B-44-11-3A-B7", true},
		{"", false},
		{"00:1B:44:11:3A", false},
		{"abc", false},
	}

	for _, tt := range tests {
		result := IsMAC(tt.input)
		if result != tt.expected {
			t.Errorf("IsMAC(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsPort(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{0, true},
		{80, true},
		{443, true},
		{8080, true},
		{65535, true},
		{-1, false},
		{65536, false},
	}

	for _, tt := range tests {
		result := IsPort(tt.input)
		if result != tt.expected {
			t.Errorf("IsPort(%d) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsMobile(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"13800138000", true},
		{"15912345678", true},
		{"18600000000", true},
		{"17700000000", true},
		{"+8613800138000", true},
		{"8613800138000", true},
		{"", false},
		{"12345678901", false},
		{"1380013800", false},
		{"138001380001", false},
		{"23800138000", false},
	}

	for _, tt := range tests {
		result := IsMobile(tt.input)
		if result != tt.expected {
			t.Errorf("IsMobile(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsPhone(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"010-12345678", true},
		{"021-12345678", true},
		{"0755-1234567", true},
		{"12345678", true},
		{"400-123-4567", true},
		{"800-123-4567", true},
		{"+8601012345678", true},
		{"", false},
		{"abc", false},
	}

	for _, tt := range tests {
		result := IsPhone(tt.input)
		if result != tt.expected {
			t.Errorf("IsPhone(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsIDCard(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"11010519491231002X", true},
		{"11010519491231002x", true},
		{"110105194912310038", true},
		{"", false},
		{"12345678901234567", false},
		{"11010519491231002Y", false},
		{"110105194913310020", false}, // invalid month
	}

	for _, tt := range tests {
		result := IsIDCard(tt.input)
		if result != tt.expected {
			t.Errorf("IsIDCard(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsZipCode(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"100000", true},
		{"200000", true},
		{"518000", true},
		{"", false},
		{"12345", false},
		{"1234567", false},
		{"abc123", false},
	}

	for _, tt := range tests {
		result := IsZipCode(tt.input)
		if result != tt.expected {
			t.Errorf("IsZipCode(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"ABC", true},
		{"abcABC", true},
		{"", false},
		{"abc123", false},
		{"abc def", false},
		{"123", false},
	}

	for _, tt := range tests {
		result := IsAlpha(tt.input)
		if result != tt.expected {
			t.Errorf("IsAlpha(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc123", true},
		{"ABC123", true},
		{"abc", true},
		{"123", true},
		{"", false},
		{"abc 123", false},
		{"abc-123", false},
	}

	for _, tt := range tests {
		result := IsAlphaNumeric(tt.input)
		if result != tt.expected {
			t.Errorf("IsAlphaNumeric(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsNumericString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"0", true},
		{"", false},
		{"12.3", false},
		{"abc", false},
		{"12abc", false},
	}

	for _, tt := range tests {
		result := IsNumeric(tt.input)
		if result != tt.expected {
			t.Errorf("IsNumeric(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"-123", true},
		{"+123", true},
		{"0", true},
		{"", false},
		{"12.3", false},
		{"abc", false},
		{"+", false},
		{"-", false},
	}

	for _, tt := range tests {
		result := IsInteger(tt.input)
		if result != tt.expected {
			t.Errorf("IsInteger(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"3.14", true},
		{"-3.14", true},
		{"+3.14", true},
		{"0", true},
		{"0.0", true},
		{"123", true},
		{"", false},
		{"abc", false},
		{"3.14.15", false},
	}

	for _, tt := range tests {
		result := IsFloat(tt.input)
		if result != tt.expected {
			t.Errorf("IsFloat(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsUUID(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"550e8400-e29b-41d4-a716-446655440000", true},
		{"6ba7b810-9dad-11d1-80b4-00c04fd430c8", true},
		{"6BA7B810-9DAD-11D1-80B4-00C04FD430C8", true},
		{"", false},
		{"550e8400-e29b-41d4-a716", false},
		{"550e8400e29b41d4a716446655440000", false},
		{"abc", false},
	}

	for _, tt := range tests {
		result := IsUUID(tt.input)
		if result != tt.expected {
			t.Errorf("IsUUID(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`{}`, true},
		{`{"key": "value"}`, true},
		{`[1, 2, 3]`, true},
		{`"string"`, true},
		{`123`, true},
		{`true`, true},
		{`null`, true},
		{"", false},
		{"{invalid}", false},
		{"[1, 2,]", false},
	}

	for _, tt := range tests {
		result := IsJSON(tt.input)
		if result != tt.expected {
			t.Errorf("IsJSON(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsBase64(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"SGVsbG8=", true},
		{"SGVsbG8gV29ybGQ=", true},
		{"", false},
		{"abc", true}, // Valid base64 encoding of "i\xb7\x1d"
		{"!@#$%", false},
	}

	for _, tt := range tests {
		result := IsBase64(tt.input)
		if result != tt.expected {
			t.Errorf("IsBase64(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsHexColor(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"#fff", true},
		{"#FFF", true},
		{"#ffffff", true},
		{"#FFFFFF", true},
		{"#123456", true},
		{"#1234", true},
		{"#12345678", true},
		{"fff", true},
		{"ffffff", true},
		{"", false},
		{"#ggg", false},
		{"#12", false},
		{"#12345", false},
		{"#1234567", false},
		{"#123456789", false},
	}

	for _, tt := range tests {
		result := IsHexColor(tt.input)
		if result != tt.expected {
			t.Errorf("IsHexColor(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsRGBColor(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"rgb(255, 0, 0)", true},
		{"rgb(0, 255, 0)", true},
		{"rgb(0, 0, 255)", true},
		{"rgba(255, 0, 0, 0.5)", true},
		{"rgba(0, 0, 0, 1)", true},
		{"rgba(0, 0, 0, 0)", true},
		{"", false},
		{"rgb(256, 0, 0)", false},
		{"rgb(0, 0, 0, 0)", false},
		{"rgba(0, 0, 0, 1.1)", false},
		{"rgba(0, 0, 0, -0.1)", false},
		{"rgb(abc, 0, 0)", false},
	}

	for _, tt := range tests {
		result := IsRGBColor(tt.input)
		if result != tt.expected {
			t.Errorf("IsRGBColor(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsDate(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"2023-01-01", true},
		{"2023/01/01", true},
		{"2023-12-31", true},
		{"", false},
		{"2023-13-01", false},
		{"2023-00-01", false},
		{"2023-01-32", false},
		{"2023-01-00", false},
		{"abc", false},
	}

	for _, tt := range tests {
		result := IsDate(tt.input)
		if result != tt.expected {
			t.Errorf("IsDate(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsChinese(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"你好", true},
		{"世界", true},
		{"", false},
		{"hello", false},
		{"你好hello", false},
		{"123", false},
	}

	for _, tt := range tests {
		result := IsChinese(tt.input)
		if result != tt.expected {
			t.Errorf("IsChinese(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestContainsChinese(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"你好", true},
		{"hello你好", true},
		{"你好world", true},
		{"hello", false},
		{"", false},
		{"123", false},
	}

	for _, tt := range tests {
		result := ContainsChinese(tt.input)
		if result != tt.expected {
			t.Errorf("ContainsChinese(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
	}{
		{'a', true},
		{'Z', true},
		{'0', false},
		{' ', false},
		{'@', false},
	}

	for _, tt := range tests {
		result := IsLetter(tt.input)
		if result != tt.expected {
			t.Errorf("IsLetter(%c) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
	}{
		{'0', true},
		{'9', true},
		{'a', false},
		{'Z', false},
		{' ', false},
	}

	for _, tt := range tests {
		result := IsDigit(tt.input)
		if result != tt.expected {
			t.Errorf("IsDigit(%c) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsUpperCase(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
	}{
		{'A', true},
		{'Z', true},
		{'a', false},
		{'z', false},
		{'0', false},
	}

	for _, tt := range tests {
		result := IsUpperCase(tt.input)
		if result != tt.expected {
			t.Errorf("IsUpperCase(%c) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsLowerCase(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
	}{
		{'a', true},
		{'z', true},
		{'A', false},
		{'Z', false},
		{'0', false},
	}

	for _, tt := range tests {
		result := IsLowerCase(tt.input)
		if result != tt.expected {
			t.Errorf("IsLowerCase(%c) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestMatches(t *testing.T) {
	tests := []struct {
		s        string
		pattern  string
		expected bool
	}{
		{"hello", "hello", true},
		{"hello world", "hello", true},
		{"hello", "^hello$", true},
		{"hello world", "^hello$", false},
		{"123", `\d+`, true},
		{"abc", `\d+`, false},
		{"", "", true},
		{"test", "", true},
	}

	for _, tt := range tests {
		result := Matches(tt.s, tt.pattern)
		if result != tt.expected {
			t.Errorf("Matches(%q, %q) = %v, want %v", tt.s, tt.pattern, result, tt.expected)
		}
	}
}

func TestMatchesAll(t *testing.T) {
	tests := []struct {
		s        string
		pattern  string
		expected bool
	}{
		{"hello", "hello", true},
		{"hello world", "hello", false},
		{"hello", "^hello$", true},
		{"hello world", "^hello$", false},
		{"123", `\d+`, true},
		{"123abc", `\d+`, false},
		{"", "", true},
		{"test", "", false},
	}

	for _, tt := range tests {
		result := MatchesAll(tt.s, tt.pattern)
		if result != tt.expected {
			t.Errorf("MatchesAll(%q, %q) = %v, want %v", tt.s, tt.pattern, result, tt.expected)
		}
	}
}

func TestRange(t *testing.T) {
	tests := []struct {
		val      int
		min      int
		max      int
		expected bool
	}{
		{5, 1, 10, true},
		{1, 1, 10, true},
		{10, 1, 10, true},
		{0, 1, 10, false},
		{11, 1, 10, false},
	}

	for _, tt := range tests {
		result := Range(tt.val, tt.min, tt.max)
		if result != tt.expected {
			t.Errorf("Range(%d, %d, %d) = %v, want %v", tt.val, tt.min, tt.max, result, tt.expected)
		}
	}
}

func TestRangeFloat(t *testing.T) {
	tests := []struct {
		val      float64
		min      float64
		max      float64
		expected bool
	}{
		{5.5, 1.0, 10.0, true},
		{1.0, 1.0, 10.0, true},
		{10.0, 1.0, 10.0, true},
		{0.5, 1.0, 10.0, false},
		{10.5, 1.0, 10.0, false},
	}

	for _, tt := range tests {
		result := RangeFloat(tt.val, tt.min, tt.max)
		if result != tt.expected {
			t.Errorf("RangeFloat(%f, %f, %f) = %v, want %v", tt.val, tt.min, tt.max, result, tt.expected)
		}
	}
}

func TestMinLength(t *testing.T) {
	tests := []struct {
		s        string
		min      int
		expected bool
	}{
		{"hello", 3, true},
		{"hello", 5, true},
		{"hello", 6, false},
		{"", 1, false},
		{"", 0, true},
	}

	for _, tt := range tests {
		result := MinLength(tt.s, tt.min)
		if result != tt.expected {
			t.Errorf("MinLength(%q, %d) = %v, want %v", tt.s, tt.min, result, tt.expected)
		}
	}
}

func TestMaxLength(t *testing.T) {
	tests := []struct {
		s        string
		max      int
		expected bool
	}{
		{"hello", 5, true},
		{"hello", 6, true},
		{"hello", 4, false},
		{"", 0, true},
		{"", 1, true},
	}

	for _, tt := range tests {
		result := MaxLength(tt.s, tt.max)
		if result != tt.expected {
			t.Errorf("MaxLength(%q, %d) = %v, want %v", tt.s, tt.max, result, tt.expected)
		}
	}
}

func TestLengthBetween(t *testing.T) {
	tests := []struct {
		s        string
		min      int
		max      int
		expected bool
	}{
		{"hello", 3, 7, true},
		{"hello", 5, 5, true},
		{"hello", 1, 4, false},
		{"hello", 6, 10, false},
		{"", 0, 0, true},
		{"", 1, 5, false},
	}

	for _, tt := range tests {
		result := LengthBetween(tt.s, tt.min, tt.max)
		if result != tt.expected {
			t.Errorf("LengthBetween(%q, %d, %d) = %v, want %v", tt.s, tt.min, tt.max, result, tt.expected)
		}
	}
}

func TestCreditCard(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Visa
		{"4111111111111111", true},
		// Mastercard
		{"5500000000000004", true},
		// American Express
		{"340000000000009", true},
		// Discover
		{"6011000000000004", true},
		// With spaces
		{"4111 1111 1111 1111", true},
		// With dashes
		{"4111-1111-1111-1111", true},
		// Invalid
		{"", false},
		{"123", false},
		{"4111111111111112", false}, // Invalid checksum
	}

	for _, tt := range tests {
		result := CreditCard(tt.input)
		if result != tt.expected {
			t.Errorf("CreditCard(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}
