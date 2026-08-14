package regexutil

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	if !IsMatch("test@example.com", PatternEmail) {
		t.Error("Should match email")
	}
	if IsMatch("invalid-email", PatternEmail) {
		t.Error("Should not match invalid email")
	}
	if !IsMatch("13812345678", PatternMobile) {
		t.Error("Should match mobile")
	}
	if IsMatch("12345678901", PatternMobile) {
		t.Error("Should not match invalid mobile")
	}
}

func TestFind(t *testing.T) {
	result := Find("hello 123 world", `\d+`)
	if result != "123" {
		t.Errorf("Find = %q, want \"123\"", result)
	}
}

func TestFindAll(t *testing.T) {
	results := FindAll("abc 123 def 456", `\d+`)
	if len(results) != 2 {
		t.Errorf("FindAll returned %d results, want 2", len(results))
	}
	if results[0] != "123" || results[1] != "456" {
		t.Errorf("FindAll = %v, want [123 456]", results)
	}
}

func TestFindGroups(t *testing.T) {
	groups := FindGroups("test@example.com", `(\w+)@(\w+)\.(\w+)`)
	if len(groups) != 4 {
		t.Errorf("FindGroups returned %d groups, want 4", len(groups))
	}
	if groups[1] != "test" {
		t.Errorf("Group 1 = %q, want \"test\"", groups[1])
	}
	if groups[2] != "example" {
		t.Errorf("Group 2 = %q, want \"example\"", groups[2])
	}
}

func TestFindNamedGroups(t *testing.T) {
	groups := FindNamedGroups("test@example.com", `(?P<user>\w+)@(?P<host>\w+)\.(?P<domain>\w+)`)
	if groups == nil {
		t.Fatal("FindNamedGroups returned nil")
	}
	if groups["user"] != "test" {
		t.Errorf("user = %q, want \"test\"", groups["user"])
	}
	if groups["host"] != "example" {
		t.Errorf("host = %q, want \"example\"", groups["host"])
	}
}

func TestReplace(t *testing.T) {
	result := Replace("hello 123 world 456", `\d+`, "NUM")
	if result != "hello NUM world NUM" {
		t.Errorf("Replace = %q, want \"hello NUM world NUM\"", result)
	}
}

func TestReplaceFunc(t *testing.T) {
	result := ReplaceFunc("hello 123 world", `\d+`, func(s string) string {
		return "[" + s + "]"
	})
	if result != "hello [123] world" {
		t.Errorf("ReplaceFunc = %q, want \"hello [123] world\"", result)
	}
}

func TestSplit(t *testing.T) {
	parts := Split("a,b,,c,d", `,`)
	if len(parts) != 5 {
		t.Errorf("Split returned %d parts, want 5", len(parts))
	}
}

func TestCount(t *testing.T) {
	c := Count("abc abc abc", `abc`)
	if c != 3 {
		t.Errorf("Count = %d, want 3", c)
	}
}

func TestExtractGroup(t *testing.T) {
	result := ExtractGroup("test@example.com", `(\w+)@(\w+)\.(\w+)`, 2)
	if result != "example" {
		t.Errorf("ExtractGroup = %q, want \"example\"", result)
	}
}

func TestExtractNamedGroup(t *testing.T) {
	result := ExtractNamedGroup("test@example.com", `(?P<user>\w+)@(?P<host>\w+)\.\w+`, "user")
	if result != "test" {
		t.Errorf("ExtractNamedGroup = %q, want \"test\"", result)
	}
}

func TestQuote(t *testing.T) {
	result := Quote("hello.world")
	if result != `hello\.world` {
		t.Errorf("Quote = %q, want \"hello\\.world\"", result)
	}
}

func TestCommonPatterns(t *testing.T) {
	if !IsEmail("test@example.com") {
		t.Error("IsEmail should match")
	}
	if IsEmail("invalid") {
		t.Error("IsEmail should not match")
	}
	if !IsMobile("13812345678") {
		t.Error("IsMobile should match")
	}
	if !IsAlpha("abcDEF") {
		t.Error("IsAlpha should match")
	}
	if IsAlpha("abc123") {
		t.Error("IsAlpha should not match")
	}
	if !IsNumeric("12345") {
		t.Error("IsNumeric should match")
	}
	if !IsAlphaNumeric("abc123") {
		t.Error("IsAlphaNumeric should match")
	}
	if !IsDate("2024-01-15") {
		t.Error("IsDate should match")
	}
	if !IsDateTime("2024-01-15 10:30:00") {
		t.Error("IsDateTime should match")
	}
}

func TestCache(t *testing.T) {
	ClearCache()
	if CacheSize() != 0 {
		t.Error("Cache should be empty after clear")
	}

	IsMatch("test", `\d+`)
	if CacheSize() != 1 {
		t.Errorf("CacheSize = %d, want 1", CacheSize())
	}

	ClearCache()
}

// --- New tests below ---

func TestMustGetRegex(t *testing.T) {
	t.Run("valid pattern", func(t *testing.T) {
		re := MustGetRegex(`\d+`)
		if re == nil {
			t.Fatal("MustGetRegex returned nil for valid pattern")
		}
		if !re.MatchString("123") {
			t.Error("Compiled regex should match \"123\"")
		}
	})

	t.Run("invalid pattern panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustGetRegex should panic on invalid pattern")
			}
		}()
		MustGetRegex(`(?:invalid`)
	})
}

func TestIsMatchBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		pattern string
		want    bool
	}{
		{"match digits", []byte("abc123def"), `\d+`, true},
		{"no match", []byte("abcdef"), `\d+`, false},
		{"invalid pattern", []byte("abc"), `(?:bad`, false},
		{"empty input", []byte(""), `\d+`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsMatchBytes(tt.input, tt.pattern)
			if got != tt.want {
				t.Errorf("IsMatchBytes(%q, %q) = %v, want %v", tt.input, tt.pattern, got, tt.want)
			}
		})
	}
}

func TestFindInvalidPattern(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		pattern string
		want    string
	}{
		{"invalid pattern returns empty", "abc", `(?:bad`, ""},
		{"no match returns empty", "abc", `\d+`, ""},
		{"match found", "abc123", `\d+`, "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Find(tt.s, tt.pattern)
			if got != tt.want {
				t.Errorf("Find(%q, %q) = %q, want %q", tt.s, tt.pattern, got, tt.want)
			}
		})
	}
}

func TestFindBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		pattern string
		want    string
	}{
		{"match found", []byte("hello 42 world"), `\d+`, "42"},
		{"no match", []byte("hello world"), `\d+`, ""},
		{"invalid pattern", []byte("hello"), `(?:bad`, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindBytes(tt.input, tt.pattern)
			if tt.want == "" {
				if got != nil {
					t.Errorf("FindBytes(%q, %q) = %q, want nil", tt.input, tt.pattern, got)
				}
			} else {
				if string(got) != tt.want {
					t.Errorf("FindBytes(%q, %q) = %q, want %q", tt.input, tt.pattern, got, tt.want)
				}
			}
		})
	}
}

func TestFindAllInvalidPattern(t *testing.T) {
	got := FindAll("abc 123", `(?:bad`)
	if got != nil {
		t.Errorf("FindAll with invalid pattern = %v, want nil", got)
	}
}

func TestFindAllBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		pattern string
		wantLen int
	}{
		{"two matches", []byte("abc 123 def 456"), `\d+`, 2},
		{"no match", []byte("abcdef"), `\d+`, 0},
		{"invalid pattern", []byte("abc"), `(?:bad`, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAllBytes(tt.input, tt.pattern)
			if tt.wantLen == 0 {
				if got != nil {
					t.Errorf("FindAllBytes(%q, %q) = %v, want nil", tt.input, tt.pattern, got)
				}
			} else {
				if len(got) != tt.wantLen {
					t.Errorf("FindAllBytes(%q, %q) returned %d matches, want %d", tt.input, tt.pattern, len(got), tt.wantLen)
				}
			}
		})
	}
}

func TestFindAllWithIndex(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		pattern string
		wantLen int
	}{
		{"two matches", "abc 123 def 456", `\d+`, 2},
		{"no match", "abcdef", `\d+`, 0},
		{"invalid pattern", "abc", `(?:bad`, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAllWithIndex(tt.s, tt.pattern)
			if tt.wantLen == 0 {
				if got != nil {
					t.Errorf("FindAllWithIndex(%q, %q) = %v, want nil", tt.s, tt.pattern, got)
				}
			} else {
				if len(got) != tt.wantLen {
					t.Errorf("FindAllWithIndex returned %d matches, want %d", len(got), tt.wantLen)
				}
				// Verify first match indices
				if got[0][0] != 4 || got[0][1] != 7 {
					t.Errorf("First match index = %v, want [4 7]", got[0])
				}
			}
		})
	}
}

func TestFindGroupsInvalidPattern(t *testing.T) {
	got := FindGroups("abc", `(?:bad`)
	if got != nil {
		t.Errorf("FindGroups with invalid pattern = %v, want nil", got)
	}
}

func TestFindAllGroups(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		pattern string
		wantLen int
	}{
		{"two group matches", "a1 b2 c3", `([a-z])(\d)`, 3},
		{"no match", "abc", `(\d)(\d)`, 0},
		{"invalid pattern", "abc", `(?:bad`, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAllGroups(tt.s, tt.pattern)
			if tt.wantLen == 0 {
				if got != nil {
					t.Errorf("FindAllGroups(%q, %q) = %v, want nil", tt.s, tt.pattern, got)
				}
			} else {
				if len(got) != tt.wantLen {
					t.Errorf("FindAllGroups returned %d matches, want %d", len(got), tt.wantLen)
				}
				// Check first match groups
				if got[0][1] != "a" || got[0][2] != "1" {
					t.Errorf("First match groups = %v, want [a1 a 1]", got[0])
				}
			}
		})
	}
}

func TestFindNamedGroupsInvalidPattern(t *testing.T) {
	got := FindNamedGroups("abc", `(?:bad`)
	if got != nil {
		t.Errorf("FindNamedGroups with invalid pattern = %v, want nil", got)
	}
}

func TestFindNamedGroupsNoMatch(t *testing.T) {
	got := FindNamedGroups("abc", `(?P<digit>\d+)`)
	if got != nil {
		t.Errorf("FindNamedGroups with no match = %v, want nil", got)
	}
}

func TestReplaceInvalidPattern(t *testing.T) {
	got := Replace("hello", `(?:bad`, "x")
	if got != "hello" {
		t.Errorf("Replace with invalid pattern = %q, want \"hello\"", got)
	}
}

func TestReplaceBytes(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		pattern     string
		replacement []byte
		want        string
	}{
		{"replace digits", []byte("abc 123 def"), `\d+`, []byte("NUM"), "abc NUM def"},
		{"invalid pattern", []byte("hello"), `(?:bad`, []byte("x"), "hello"},
		{"no match", []byte("hello"), `\d+`, []byte("NUM"), "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceBytes(tt.input, tt.pattern, tt.replacement)
			if string(got) != tt.want {
				t.Errorf("ReplaceBytes(%q, %q, %q) = %q, want %q", tt.input, tt.pattern, tt.replacement, got, tt.want)
			}
		})
	}
}

func TestReplaceFuncInvalidPattern(t *testing.T) {
	got := ReplaceFunc("hello", `(?:bad`, func(s string) string { return "x" })
	if got != "hello" {
		t.Errorf("ReplaceFunc with invalid pattern = %q, want \"hello\"", got)
	}
}

func TestReplaceFirst(t *testing.T) {
	tests := []struct {
		name        string
		s           string
		pattern     string
		replacement string
		want        string
	}{
		{"replaces first match only", "hello 123 world 456", `\d+`, "NUM", "hello NUM world 456"},
		{"supports group references", "hello 123 world", `(\d)(\d)(\d)`, "$3$2$1", "hello 321 world"},
		{"invalid pattern", "hello", `(?:bad`, "x", "hello"},
		{"no match", "hello", `\d+`, "NUM", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceFirst(tt.s, tt.pattern, tt.replacement)
			if got != tt.want {
				t.Errorf("ReplaceFirst(%q, %q, %q) = %q, want %q", tt.s, tt.pattern, tt.replacement, got, tt.want)
			}
		})
	}
}

func TestSplitInvalidPattern(t *testing.T) {
	got := Split("a,b,c", `(?:bad`)
	if len(got) != 1 || got[0] != "a,b,c" {
		t.Errorf("Split with invalid pattern = %v, want [\"a,b,c\"]", got)
	}
}

func TestCountInvalidPattern(t *testing.T) {
	got := Count("abc", `(?:bad`)
	if got != 0 {
		t.Errorf("Count with invalid pattern = %d, want 0", got)
	}
}

func TestExtractGroupEdgeCases(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		pattern    string
		groupIndex int
		want       string
	}{
		{"valid group", "test@example.com", `(\w+)@(\w+)\.(\w+)`, 2, "example"},
		{"index out of range", "abc", `(\w+)`, 5, ""},
		{"no match", "abc", `(\d+)`, 1, ""},
		{"invalid pattern", "abc", `(?:bad`, 0, ""},
		{"full match group 0", "abc123", `(\w+)(\d+)`, 0, "abc123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractGroup(tt.s, tt.pattern, tt.groupIndex)
			if got != tt.want {
				t.Errorf("ExtractGroup(%q, %q, %d) = %q, want %q", tt.s, tt.pattern, tt.groupIndex, got, tt.want)
			}
		})
	}
}

func TestExtractNamedGroupEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		pattern string
		key     string
		want    string
	}{
		{"valid named group", "test@example.com", `(?P<user>\w+)@(?P<host>\w+)\.\w+`, "user", "test"},
		{"missing key", "test@example.com", `(?P<user>\w+)@(?P<host>\w+)\.\w+`, "missing", ""},
		{"no match", "abc", `(?P<digit>\d+)`, "digit", ""},
		{"invalid pattern", "abc", `(?:bad`, "x", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractNamedGroup(tt.s, tt.pattern, tt.key)
			if got != tt.want {
				t.Errorf("ExtractNamedGroup(%q, %q, %q) = %q, want %q", tt.s, tt.pattern, tt.key, got, tt.want)
			}
		})
	}
}

func TestMatchAll(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		pattern string
		want    bool
	}{
		{"full match", "abc123", `^[a-z]+\d+$`, true},
		{"partial match only", "abc123xyz", `^[a-z]+`, false},
		{"no match", "xyz", `^\d+$`, false},
		{"invalid pattern", "abc", `(?:bad`, false},
		{"exact word", "hello", `^hello$`, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MatchAll(tt.s, tt.pattern)
			if got != tt.want {
				t.Errorf("MatchAll(%q, %q) = %v, want %v", tt.s, tt.pattern, got, tt.want)
			}
		})
	}
}

func TestContainsMatch(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		pattern string
		want    bool
	}{
		{"contains digits", "abc123def", `\d+`, true},
		{"no digits", "abcdef", `\d+`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsMatch(tt.s, tt.pattern)
			if got != tt.want {
				t.Errorf("ContainsMatch(%q, %q) = %v, want %v", tt.s, tt.pattern, got, tt.want)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"http url", "http://example.com", true},
		{"https url", "https://example.com/path?q=1", true},
		{"ftp url", "ftp://files.example.com/file.txt", true},
		{"no scheme", "example.com", false},
		{"empty string", "", false},
		{"invalid scheme", "mailto:user@example.com", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsURL(tt.input)
			if got != tt.want {
				t.Errorf("IsURL(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsIP(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid IPv4", "192.168.1.1", true},
		{"valid all zeros", "0.0.0.0", true},
		{"valid max", "255.255.255.255", true},
		{"out of range octet", "256.1.1.1", false},
		{"not IP", "abc.def.ghi.jkl", false},
		{"empty", "", false},
		{"IPv6", "::1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsIP(tt.input)
			if got != tt.want {
				t.Errorf("IsIP(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPhone(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid with dash", "010-12345678", true},
		{"valid without dash", "01012345678", true},
		{"valid short area", "0755-1234567", true},
		{"invalid", "13812345678", false},
		{"empty", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPhone(tt.input)
			if got != tt.want {
				t.Errorf("IsPhone(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsIDCard(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid 18 digits", "110101199003071234", true},
		{"valid with X", "11010119900307123X", true},
		{"valid with x", "11010119900307123x", true},
		{"too short", "12345", false},
		{"invalid char", "11010119900307123A", false},
		{"empty", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsIDCard(tt.input)
			if got != tt.want {
				t.Errorf("IsIDCard(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsZipCode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid 6 digits", "100000", true},
		{"valid rural", "010000", true},
		{"too short", "10000", false},
		{"too long", "1000000", false},
		{"letters", "abcdef", false},
		{"empty", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsZipCode(tt.input)
			if got != tt.want {
				t.Errorf("IsZipCode(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsChinese(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"pure Chinese", "你好世界", true},
		{"single char", "中", true},
		{"mixed with English", "你好world", false},
		{"pure English", "hello", false},
		{"empty", "", false},
		{"digits", "12345", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsChinese(tt.input)
			if got != tt.want {
				t.Errorf("IsChinese(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestContainsChinese(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"pure Chinese", "你好世界", true},
		{"mixed", "hello你好world", true},
		{"single Chinese char", "a中b", true},
		{"no Chinese", "hello world", false},
		{"empty", "", false},
		{"digits only", "12345", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsChinese(tt.input)
			if got != tt.want {
				t.Errorf("ContainsChinese(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsUUID(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid lowercase", "550e8400-e29b-41d4-a716-446655440000", true},
		{"valid uppercase", "550E8400-E29B-41D4-A716-446655440000", true},
		{"valid mixed case", "550e8400-E29B-41d4-A716-446655440000", true},
		{"too short", "550e8400-e29b-41d4", false},
		{"no dashes", "550e8400e29b41d4a716446655440000", false},
		{"empty", "", false},
		{"invalid chars", "550g8400-e29b-41d4-a716-446655440000", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsUUID(tt.input)
			if got != tt.want {
				t.Errorf("IsUUID(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsHexColor(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"6 digit with hash", "#FF0000", true},
		{"6 digit without hash", "FF0000", true},
		{"3 digit with hash", "#F00", true},
		{"3 digit without hash", "F00", true},
		{"lowercase", "#ff00aa", true},
		{"mixed case", "#aAbBcC", true},
		{"invalid length 4", "#FF00", false},
		{"invalid length 5", "#FF000", false},
		{"invalid chars", "#GGGGGG", false},
		{"empty", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsHexColor(tt.input)
			if got != tt.want {
				t.Errorf("IsHexColor(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCacheHit(t *testing.T) {
	// Use a unique pattern unlikely to collide with other tests
	uniquePattern := `^[a-z]{3}!test_cache_hit$`
	ClearCache()
	// First call: cache miss
	_ = IsMatch("abc!test_cache_hit", uniquePattern)
	size1 := CacheSize()
	// Second call: cache hit
	_ = IsMatch("abc!test_cache_hit", uniquePattern)
	size2 := CacheSize()
	if size2 != size1 {
		t.Errorf("Cache size changed from %d to %d on repeated call", size1, size2)
	}
	ClearCache()
}

func TestIsMatchInvalidPattern(t *testing.T) {
	got := IsMatch("abc", `(?:bad`)
	if got != false {
		t.Error("IsMatch with invalid pattern should return false")
	}
}

func TestMatchAllPartialMatch(t *testing.T) {
	// Test where regex matches but doesn't match entire string
	// [a-z]{3} is greedy but only matches exactly 3 chars, so FindString("abcxyz") = "abc" != "abcxyz"
	got := MatchAll("abcxyz", `[a-z]{3}`)
	if got != false {
		t.Error("MatchAll should return false when match doesn't cover entire string")
	}
}
