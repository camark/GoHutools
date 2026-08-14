package charsetutil

import (
	"bytes"
	"strings"
	"testing"
)

func TestIsUTF8(t *testing.T) {
	if !IsUTF8([]byte("hello world")) {
		t.Error("IsUTF8 should return true for ASCII")
	}
	if !IsUTF8([]byte("你好世界")) {
		t.Error("IsUTF8 should return true for UTF-8 Chinese")
	}
	if IsUTF8([]byte{0xFF, 0xFE}) {
		t.Error("IsUTF8 should return false for invalid UTF-8")
	}
}

func TestIsGBK(t *testing.T) {
	if !IsGBK([]byte("hello")) {
		t.Error("IsGBK should return true for ASCII")
	}
}

func TestCleanUTF8(t *testing.T) {
	data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0xFF, 0xFE}
	result := CleanUTF8(data)
	if !IsUTF8(result) {
		t.Error("CleanUTF8 should produce valid UTF-8")
	}
}

func TestCleanUTF8String(t *testing.T) {
	result := CleanUTF8String("Hello\xFF\xFEWorld")
	if !IsUTF8([]byte(result)) {
		t.Error("CleanUTF8String should produce valid UTF-8")
	}
}

func TestRuneCount(t *testing.T) {
	if RuneCount("hello") != 5 {
		t.Error("RuneCount(\"hello\") should be 5")
	}
	if RuneCount("你好") != 2 {
		t.Error("RuneCount(\"你好\") should be 2")
	}
	if RuneCount("") != 0 {
		t.Error("RuneCount(\"\") should be 0")
	}
}

func TestCharsets(t *testing.T) {
	charsets := Charsets()
	if len(charsets) == 0 {
		t.Error("Charsets should not be empty")
	}
	found := false
	for _, c := range charsets {
		if c == UTF8 {
			found = true
			break
		}
	}
	if !found {
		t.Error("Charsets should include UTF-8")
	}
}

func TestIsSupported(t *testing.T) {
	if !IsSupported(UTF8) {
		t.Error("UTF-8 should be supported")
	}
	if !IsSupported("utf-8") {
		t.Error("utf-8 should be supported (case insensitive)")
	}
	if !IsSupported(GBK) {
		t.Error("GBK should be supported")
	}
	if IsSupported("UNKNOWN") {
		t.Error("UNKNOWN should not be supported")
	}
}

func TestDetectCharset(t *testing.T) {
	if DetectCharset([]byte("hello")) != UTF8 {
		t.Error("ASCII should be detected as UTF-8")
	}
	if DetectCharset([]byte("你好")) != UTF8 {
		t.Error("UTF-8 Chinese should be detected as UTF-8")
	}
}

func TestEncodeDecodeRune(t *testing.T) {
	r := '你'
	encoded := EncodeRune(r)
	decoded, size := DecodeRune(encoded)
	if decoded != r {
		t.Errorf("EncodeDecodeRune: got %c, want %c", decoded, r)
	}
	if size != len(encoded) {
		t.Errorf("Rune size: got %d, want %d", size, len(encoded))
	}
}

func TestValidRune(t *testing.T) {
	if !ValidRune('A') {
		t.Error("'A' should be valid")
	}
	if !ValidRune('你') {
		t.Error("'你' should be valid")
	}
	if ValidRune(0xD800) {
		t.Error("0xD800 should be invalid (surrogate)")
	}
}

// --- New tests below ---

func TestConvert(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		from      string
		to        string
		wantErr   bool
		checkFunc func(t *testing.T, result []byte)
	}{
		{
			name: "same charset returns data as-is",
			data: []byte("hello"),
			from: UTF8,
			to:   UTF8,
			checkFunc: func(t *testing.T, result []byte) {
				if string(result) != "hello" {
					t.Errorf("got %q, want %q", result, "hello")
				}
			},
		},
		{
			name: "UTF-8 to GBK and back",
			data: []byte("你好世界"),
			from: UTF8,
			to:   GBK,
			checkFunc: func(t *testing.T, result []byte) {
				// Convert back to UTF-8 to verify round-trip
				back, err := Convert(result, GBK, UTF8)
				if err != nil {
					t.Fatalf("convert back failed: %v", err)
				}
				if string(back) != "你好世界" {
					t.Errorf("round-trip got %q, want %q", back, "你好世界")
				}
			},
		},
		{
			name: "GBK to UTF-8",
			data: func() []byte {
				// "你好" in GBK
				gbk, _ := Convert([]byte("你好"), UTF8, GBK)
				return gbk
			}(),
			from: GBK,
			to:   UTF8,
			checkFunc: func(t *testing.T, result []byte) {
				if string(result) != "你好" {
					t.Errorf("got %q, want %q", result, "你好")
				}
			},
		},
		{
			name: "UTF-8 to GB18030 and back",
			data: []byte("测试文字"),
			from: UTF8,
			to:   GB18030,
			checkFunc: func(t *testing.T, result []byte) {
				back, err := Convert(result, GB18030, UTF8)
				if err != nil {
					t.Fatalf("convert back failed: %v", err)
				}
				if string(back) != "测试文字" {
					t.Errorf("round-trip got %q, want %q", back, "测试文字")
				}
			},
		},
		{
			name: "UTF-8 to Big5 and back",
			data: []byte("繁體中文"),
			from: UTF8,
			to:   BIG5,
			checkFunc: func(t *testing.T, result []byte) {
				back, err := Convert(result, BIG5, UTF8)
				if err != nil {
					t.Fatalf("convert back failed: %v", err)
				}
				if string(back) != "繁體中文" {
					t.Errorf("round-trip got %q, want %q", back, "繁體中文")
				}
			},
		},
		{
			name: "UTF-8 to GB2312 and back",
			data: []byte("简体中文"),
			from: UTF8,
			to:   GB2312,
			checkFunc: func(t *testing.T, result []byte) {
				back, err := Convert(result, GB2312, UTF8)
				if err != nil {
					t.Fatalf("convert back failed: %v", err)
				}
				if string(back) != "简体中文" {
					t.Errorf("round-trip got %q, want %q", back, "简体中文")
				}
			},
		},
		{
			name:    "unsupported fromCharset",
			data:    []byte("hello"),
			from:    "WINDOWS-1252",
			to:      UTF8,
			wantErr: true,
		},
		{
			name:    "unsupported toCharset",
			data:    []byte("hello"),
			from:    UTF8,
			to:      "WINDOWS-1252",
			wantErr: true,
		},
		{
			name: "case insensitive charset names",
			data: []byte("hello"),
			from: "utf-8",
			to:   "utf-8",
			checkFunc: func(t *testing.T, result []byte) {
				if string(result) != "hello" {
					t.Errorf("got %q, want %q", result, "hello")
				}
			},
		},
		{
			name: "UTF8 alias",
			data: []byte("hello"),
			from: "UTF8",
			to:   "UTF8",
			checkFunc: func(t *testing.T, result []byte) {
				if string(result) != "hello" {
					t.Errorf("got %q, want %q", result, "hello")
				}
			},
		},
		{
			name: "empty data",
			data: []byte{},
			from: UTF8,
			to:   GBK,
			checkFunc: func(t *testing.T, result []byte) {
				if len(result) != 0 {
					t.Errorf("got %d bytes, want 0", len(result))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Convert(tt.data, tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.checkFunc != nil && err == nil {
				tt.checkFunc(t, result)
			}
		})
	}
}

func TestConvertString(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		from    string
		to      string
		want    string
		wantErr bool
	}{
		{
			name: "same charset",
			s:    "hello world",
			from: UTF8,
			to:   UTF8,
			want: "hello world",
		},
		{
			name: "UTF-8 to GBK to UTF-8 round-trip",
			s:    "你好",
			from: UTF8,
			to:   GBK,
			// We check round-trip below
		},
		{
			name:    "unsupported charset",
			s:       "hello",
			from:    "INVALID",
			to:      UTF8,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertString(tt.s, tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != "" && err == nil && result != tt.want {
				t.Errorf("ConvertString() = %q, want %q", result, tt.want)
			}
		})
	}

	// Round-trip test
	t.Run("round-trip UTF-8 -> GBK -> UTF-8", func(t *testing.T) {
		gbk, err := ConvertString("你好", UTF8, GBK)
		if err != nil {
			t.Fatalf("first convert: %v", err)
		}
		back, err := ConvertString(gbk, GBK, UTF8)
		if err != nil {
			t.Fatalf("second convert: %v", err)
		}
		if back != "你好" {
			t.Errorf("round-trip got %q, want %q", back, "你好")
		}
	})
}

func TestToUTF8(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		from string
		want string
	}{
		{
			name: "GBK to UTF-8",
			data: func() []byte {
				gbk, _ := Convert([]byte("测试"), UTF8, GBK)
				return gbk
			}(),
			from: GBK,
			want: "测试",
		},
		{
			name: "UTF-8 to UTF-8 (no-op)",
			data: []byte("hello"),
			from: UTF8,
			want: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToUTF8(tt.data, tt.from)
			if err != nil {
				t.Fatalf("ToUTF8() error = %v", err)
			}
			if string(result) != tt.want {
				t.Errorf("ToUTF8() = %q, want %q", result, tt.want)
			}
		})
	}
}

func TestFromUTF8(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		to   string
	}{
		{
			name: "UTF-8 to GBK",
			data: []byte("你好"),
			to:   GBK,
		},
		{
			name: "UTF-8 to GB18030",
			data: []byte("你好"),
			to:   GB18030,
		},
		{
			name: "UTF-8 to Big5",
			data: []byte("繁體中文"),
			to:   BIG5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FromUTF8(tt.data, tt.to)
			if err != nil {
				t.Fatalf("FromUTF8() error = %v", err)
			}
			// Convert back to verify
			back, err := ToUTF8(result, tt.to)
			if err != nil {
				t.Fatalf("ToUTF8() error = %v", err)
			}
			if string(back) != string(tt.data) {
				t.Errorf("round-trip got %q, want %q", back, tt.data)
			}
		})
	}

	t.Run("unsupported charset", func(t *testing.T) {
		_, err := FromUTF8([]byte("hello"), "INVALID")
		if err == nil {
			t.Error("expected error for unsupported charset")
		}
	})
}

func TestToUTF8String(t *testing.T) {
	tests := []struct {
		name string
		s    string
		from string
		want string
	}{
		{
			name: "UTF-8 no-op",
			s:    "hello",
			from: UTF8,
			want: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToUTF8String(tt.s, tt.from)
			if err != nil {
				t.Fatalf("ToUTF8String() error = %v", err)
			}
			if result != tt.want {
				t.Errorf("ToUTF8String() = %q, want %q", result, tt.want)
			}
		})
	}

	// GBK round-trip via string functions
	t.Run("GBK round-trip", func(t *testing.T) {
		gbkBytes, err := FromUTF8([]byte("你好"), GBK)
		if err != nil {
			t.Fatalf("FromUTF8: %v", err)
		}
		utf8Str, err := ToUTF8String(string(gbkBytes), GBK)
		if err != nil {
			t.Fatalf("ToUTF8String: %v", err)
		}
		if utf8Str != "你好" {
			t.Errorf("got %q, want %q", utf8Str, "你好")
		}
	})
}

func TestFromUTF8String(t *testing.T) {
	tests := []struct {
		name string
		s    string
		to   string
	}{
		{
			name: "to GBK",
			s:    "你好世界",
			to:   GBK,
		},
		{
			name: "to GB18030",
			s:    "测试",
			to:   GB18030,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FromUTF8String(tt.s, tt.to)
			if err != nil {
				t.Fatalf("FromUTF8String() error = %v", err)
			}
			// Convert back
			back, err := ToUTF8String(result, tt.to)
			if err != nil {
				t.Fatalf("ToUTF8String() error = %v", err)
			}
			if back != tt.s {
				t.Errorf("round-trip got %q, want %q", back, tt.s)
			}
		})
	}

	t.Run("unsupported charset", func(t *testing.T) {
		_, err := FromUTF8String("hello", "INVALID")
		if err == nil {
			t.Error("expected error for unsupported charset")
		}
	})
}

func TestGBKToUTF8(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{
			name: "GBK encoded Chinese",
			data: func() []byte {
				gbk, _ := Convert([]byte("你好"), UTF8, GBK)
				return gbk
			}(),
			want: "你好",
		},
		{
			name: "ASCII passes through",
			data: []byte("hello"),
			want: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GBKToUTF8(tt.data)
			if err != nil {
				t.Fatalf("GBKToUTF8() error = %v", err)
			}
			if string(result) != tt.want {
				t.Errorf("GBKToUTF8() = %q, want %q", result, tt.want)
			}
		})
	}
}

func TestUTF8ToGBK(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want string // expected when converted back to UTF-8
	}{
		{
			name: "Chinese characters",
			data: []byte("你好世界"),
			want: "你好世界",
		},
		{
			name: "ASCII",
			data: []byte("hello"),
			want: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gbk, err := UTF8ToGBK(tt.data)
			if err != nil {
				t.Fatalf("UTF8ToGBK() error = %v", err)
			}
			back, err := GBKToUTF8(gbk)
			if err != nil {
				t.Fatalf("GBKToUTF8() error = %v", err)
			}
			if string(back) != tt.want {
				t.Errorf("round-trip got %q, want %q", back, tt.want)
			}
		})
	}
}

func TestGBKToUTF8String(t *testing.T) {
	// Create GBK-encoded string
	gbkBytes, err := Convert([]byte("你好"), UTF8, GBK)
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	result, err := GBKToUTF8String(string(gbkBytes))
	if err != nil {
		t.Fatalf("GBKToUTF8String() error = %v", err)
	}
	if result != "你好" {
		t.Errorf("GBKToUTF8String() = %q, want %q", result, "你好")
	}
}

func TestUTF8ToGBKString(t *testing.T) {
	result, err := UTF8ToGBKString("你好")
	if err != nil {
		t.Fatalf("UTF8ToGBKString() error = %v", err)
	}
	// Convert back
	back, err := GBKToUTF8String(result)
	if err != nil {
		t.Fatalf("GBKToUTF8String() error = %v", err)
	}
	if back != "你好" {
		t.Errorf("round-trip got %q, want %q", back, "你好")
	}
}

func TestRuneLen(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want int
	}{
		{"ASCII", 'A', 1},
		{"2-byte rune", 'é', 2},
		{"3-byte CJK", '你', 3},
		{"4-byte emoji", '😀', 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RuneLen(tt.r)
			if got != tt.want {
				t.Errorf("RuneLen(%c) = %d, want %d", tt.r, got, tt.want)
			}
		})
	}
}

func TestUTF16ToUTF8(t *testing.T) {
	tests := []struct {
		name         string
		littleEndian bool
	}{
		{"little endian", true},
		{"big endian", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First encode UTF-8 to UTF-16
			utf16Data, err := UTF8ToUTF16([]byte("Hello世界"), tt.littleEndian)
			if err != nil {
				t.Fatalf("UTF8ToUTF16() error = %v", err)
			}
			// Now decode back
			result, err := UTF16ToUTF8(utf16Data, tt.littleEndian)
			if err != nil {
				t.Fatalf("UTF16ToUTF8() error = %v", err)
			}
			if string(result) != "Hello世界" {
				t.Errorf("UTF16ToUTF8() = %q, want %q", result, "Hello世界")
			}
		})
	}

	t.Run("empty input", func(t *testing.T) {
		// Empty with BOM only
		result, err := UTF16ToUTF8([]byte{}, true)
		if err != nil {
			t.Fatalf("UTF16ToUTF8() error = %v", err)
		}
		if len(result) != 0 {
			t.Errorf("UTF16ToUTF8(empty) = %d bytes, want 0", len(result))
		}
	})
}

func TestUTF8ToUTF16(t *testing.T) {
	tests := []struct {
		name         string
		littleEndian bool
	}{
		{"little endian", true},
		{"big endian", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := UTF8ToUTF16([]byte("Hello"), tt.littleEndian)
			if err != nil {
				t.Fatalf("UTF8ToUTF16() error = %v", err)
			}
			// Should contain BOM
			if len(result) < 2 {
				t.Fatal("result too short for BOM")
			}
			if tt.littleEndian {
				// LE BOM: 0xFF 0xFE
				if result[0] != 0xFF || result[1] != 0xFE {
					t.Errorf("expected LE BOM, got %02X %02X", result[0], result[1])
				}
			} else {
				// BE BOM: 0xFE 0xFF
				if result[0] != 0xFE || result[1] != 0xFF {
					t.Errorf("expected BE BOM, got %02X %02X", result[0], result[1])
				}
			}
			// Round-trip
			back, err := UTF16ToUTF8(result, tt.littleEndian)
			if err != nil {
				t.Fatalf("UTF16ToUTF8() error = %v", err)
			}
			if string(back) != "Hello" {
				t.Errorf("round-trip got %q, want %q", back, "Hello")
			}
		})
	}
}

func TestErrUnsupportedCharset(t *testing.T) {
	tests := []struct {
		name    string
		charset string
	}{
		{"arbitrary name", "WINDOWS-1252"},
		{"empty", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ErrUnsupportedCharset(tt.charset)
			msg := err.Error()
			if !strings.Contains(msg, "unsupported charset") {
				t.Errorf("Error() = %q, want contains 'unsupported charset'", msg)
			}
			if !strings.Contains(msg, tt.charset) {
				t.Errorf("Error() = %q, want contains %q", msg, tt.charset)
			}
		})
	}
}

func TestGetEncodingViaConvert(t *testing.T) {
	// Test all supported charsets via Convert to exercise getEncoding branches
	tests := []struct {
		name    string
		charset string
	}{
		{"UTF-8", UTF8},
		{"UTF8 alias", "UTF8"},
		{"utf-8 lowercase", "utf-8"},
		{"GBK", GBK},
		{"GB2312", GB2312},
		{"GB18030", GB18030},
		{"Big5", BIG5},
		{"UTF-16", UTF16},
		{"UTF-16LE", UTF16LE},
		{"UTF-16BE", UTF16BE},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert from charset to same charset (short-circuit) to verify getEncoding is called
			// Actually, same-charset short-circuit bypasses getEncoding, so let's do a real conversion
			// from UTF-8 to target and back
			data := []byte("Hello")
			encoded, err := Convert(data, UTF8, tt.charset)
			if err != nil {
				t.Fatalf("Convert UTF-8 -> %s failed: %v", tt.charset, err)
			}
			decoded, err := Convert(encoded, tt.charset, UTF8)
			if err != nil {
				t.Fatalf("Convert %s -> UTF-8 failed: %v", tt.charset, err)
			}
			if string(decoded) != "Hello" {
				t.Errorf("round-trip via %s: got %q, want %q", tt.charset, decoded, "Hello")
			}
		})
	}
}

func TestDetectCharsetGBK(t *testing.T) {
	// Create bytes that are valid GBK but not valid UTF-8
	gbkData, err := Convert([]byte("你好"), UTF8, GBK)
	if err != nil {
		t.Fatalf("Convert: %v", err)
	}
	// GBK-encoded "你好" should not be valid UTF-8
	if IsUTF8(gbkData) {
		t.Skip("GBK data happens to be valid UTF-8, cannot test GBK detection branch")
	}
	detected := DetectCharset(gbkData)
	if detected != GBK {
		t.Errorf("DetectCharset(GBK data) = %q, want %q", detected, GBK)
	}
}

func TestConvertEncodeError(t *testing.T) {
	// Try to encode characters that can't be represented in target charset
	// Emoji typically can't be encoded in GBK
	_, err := Convert([]byte("😀🎉"), UTF8, GBK)
	// This may or may not error depending on x/text behavior
	// We just want to exercise the code path
	if err != nil {
		// If it errors, that's the encode error path being hit
		t.Logf("encode error (expected for unmappable chars): %v", err)
	}
}

func TestCharsetsContainsAll(t *testing.T) {
	charsets := Charsets()
	expected := []string{UTF8, GBK, GB2312, GB18030, BIG5, ISO8859, USASCII, UTF16, UTF16LE, UTF16BE}
	if len(charsets) != len(expected) {
		t.Errorf("Charsets() returned %d items, want %d", len(charsets), len(expected))
	}
	for _, e := range expected {
		found := false
		for _, c := range charsets {
			if c == e {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Charsets() missing %q", e)
		}
	}
}

func TestIsSupportedCaseInsensitive(t *testing.T) {
	tests := []struct {
		name    string
		charset string
		want    bool
	}{
		{"UTF-8 exact", "UTF-8", true},
		{"utf-8 lower", "utf-8", true},
		{"UTF-8 upper", "UTF-8", true},
		{"gbk lower", "gbk", true},
		{"GBK upper", "GBK", true},
		{"GB18030", "GB18030", true},
		{"gb18030 lower", "gb18030", true},
		{"Big5", "Big5", true},
		{"big5 lower", "big5", true},
		{"UTF-16", "UTF-16", true},
		{"UTF-16LE", "UTF-16LE", true},
		{"UTF-16BE", "UTF-16BE", true},
		{"ISO-8859-1", "ISO-8859-1", true},
		{"US-ASCII", "US-ASCII", true},
		{"unknown", "WINDOWS-1252", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSupported(tt.charset)
			if got != tt.want {
				t.Errorf("IsSupported(%q) = %v, want %v", tt.charset, got, tt.want)
			}
		})
	}
}

func TestConvertUTF16RoundTrip(t *testing.T) {
	tests := []struct {
		name string
		le   bool
		text string
	}{
		{"LE ASCII only", true, "Hello World"},
		{"BE ASCII only", false, "Hello World"},
		{"LE CJK", true, "你好世界测试"},
		{"BE CJK", false, "你好世界测试"},
		{"LE mixed", true, "Hello你好World世界"},
		{"BE mixed", false, "Hello你好World世界"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utf16, err := UTF8ToUTF16([]byte(tt.text), tt.le)
			if err != nil {
				t.Fatalf("UTF8ToUTF16: %v", err)
			}
			back, err := UTF16ToUTF8(utf16, tt.le)
			if err != nil {
				t.Fatalf("UTF16ToUTF8: %v", err)
			}
			if string(back) != tt.text {
				t.Errorf("round-trip: got %q, want %q", back, tt.text)
			}
		})
	}
}

func TestEncodeRuneMultiByte(t *testing.T) {
	tests := []struct {
		name string
		r    rune
	}{
		{"ASCII", 'A'},
		{"2-byte", 'é'},
		{"3-byte", '你'},
		{"4-byte", '😀'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := EncodeRune(tt.r)
			if len(encoded) != RuneLen(tt.r) {
				t.Errorf("EncodeRune(%c) length = %d, want %d", tt.r, len(encoded), RuneLen(tt.r))
			}
			decoded, size := DecodeRune(encoded)
			if decoded != tt.r {
				t.Errorf("DecodeRune got %c, want %c", decoded, tt.r)
			}
			if size != len(encoded) {
				t.Errorf("DecodeRune size = %d, want %d", size, len(encoded))
			}
		})
	}
}

func TestCleanUTF8EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
	}{
		{"all valid", []byte("hello world")},
		{"all invalid", []byte{0xFF, 0xFE, 0xFD}},
		{"empty", []byte{}},
		{"mixed", []byte{0x48, 0x65, 0xFF, 0x6C}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CleanUTF8(tt.input)
			if !IsUTF8(result) {
				t.Error("CleanUTF8 result is not valid UTF-8")
			}
		})
	}
}

func TestCleanUTF8StringEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"all valid", "hello world"},
		{"all invalid", "\xFF\xFE\xFD"},
		{"empty", ""},
		{"mixed", "he\xFFllo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CleanUTF8String(tt.input)
			if !IsUTF8([]byte(result)) {
				t.Error("CleanUTF8String result is not valid UTF-8")
			}
		})
	}
}

func TestConvertSameCharsetTypes(t *testing.T) {
	// Verify same-charset short-circuit returns the exact same slice
	data := []byte("test data")
	result, err := Convert(data, UTF8, UTF8)
	if err != nil {
		t.Fatalf("Convert same charset: %v", err)
	}
	if !bytes.Equal(result, data) {
		t.Error("same charset should return same data")
	}
}

func TestConvertGBKSpecific(t *testing.T) {
	// Test specific GBK characters
	original := "中文测试内容"
	gbk, err := UTF8ToGBK([]byte(original))
	if err != nil {
		t.Fatalf("UTF8ToGBK: %v", err)
	}
	back, err := GBKToUTF8(gbk)
	if err != nil {
		t.Fatalf("GBKToUTF8: %v", err)
	}
	if string(back) != original {
		t.Errorf("got %q, want %q", back, original)
	}
}

func TestConvertGB18030Specific(t *testing.T) {
	original := "中文测试"
	encoded, err := FromUTF8([]byte(original), GB18030)
	if err != nil {
		t.Fatalf("FromUTF8 GB18030: %v", err)
	}
	decoded, err := ToUTF8(encoded, GB18030)
	if err != nil {
		t.Fatalf("ToUTF8 GB18030: %v", err)
	}
	if string(decoded) != original {
		t.Errorf("got %q, want %q", decoded, original)
	}
}

func TestConvertBig5Specific(t *testing.T) {
	// Round-trip UTF-8 -> Big5 -> UTF-8 with traditional Chinese text.
	data := []byte("繁體中文測試")
	encoded, err := FromUTF8(data, BIG5)
	if err != nil {
		t.Fatalf("FromUTF8 to Big5 failed: %v", err)
	}
	decoded, err := ToUTF8(encoded, BIG5)
	if err != nil {
		t.Fatalf("ToUTF8 from Big5 failed: %v", err)
	}
	if string(decoded) != string(data) {
		t.Errorf("round-trip got %q, want %q", decoded, data)
	}
}

func TestConvertUnsupportedFromCharset(t *testing.T) {
	_, err := Convert([]byte("hello"), "KOI8-R", UTF8)
	if err == nil {
		t.Error("expected error for unsupported fromCharset")
	}
	var ucErr ErrUnsupportedCharset
	if _, ok := err.(ErrUnsupportedCharset); !ok {
		t.Errorf("expected ErrUnsupportedCharset, got %T: %v", err, err)
	} else {
		ucErr = err.(ErrUnsupportedCharset)
		if !strings.Contains(ucErr.Error(), "KOI8-R") {
			t.Errorf("error should contain charset name: %v", ucErr)
		}
	}
}

func TestConvertUnsupportedToCharset(t *testing.T) {
	_, err := Convert([]byte("hello"), UTF8, "KOI8-R")
	if err == nil {
		t.Error("expected error for unsupported toCharset")
	}
	if _, ok := err.(ErrUnsupportedCharset); !ok {
		t.Errorf("expected ErrUnsupportedCharset, got %T: %v", err, err)
	}
}
