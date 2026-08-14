package strutil

import (
	"testing"
)

// ==================== Basic Check Functions ====================

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{" ", false},
		{"hello", false},
		{"  ", false},
	}

	for _, tt := range tests {
		if got := IsEmpty(tt.input); got != tt.want {
			t.Errorf("IsEmpty(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{" ", true},
		{"hello", true},
		{"  ", true},
	}

	for _, tt := range tests {
		if got := IsNotEmpty(tt.input); got != tt.want {
			t.Errorf("IsNotEmpty(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{" ", true},
		{"  \t\n  ", true},
		{"hello", false},
		{" hello ", false},
	}

	for _, tt := range tests {
		if got := IsBlank(tt.input); got != tt.want {
			t.Errorf("IsBlank(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsNotBlank(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{" ", false},
		{"  \t\n  ", false},
		{"hello", true},
		{" hello ", true},
	}

	for _, tt := range tests {
		if got := IsNotBlank(tt.input); got != tt.want {
			t.Errorf("IsNotBlank(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestHasBlank(t *testing.T) {
	tests := []struct {
		input []string
		want  bool
	}{
		{[]string{}, false},
		{[]string{"hello"}, false},
		{[]string{"hello", "world"}, false},
		{[]string{""}, true},
		{[]string{" "}, true},
		{[]string{"hello", "", "world"}, true},
		{[]string{"hello", " ", "world"}, true},
	}

	for _, tt := range tests {
		if got := HasBlank(tt.input...); got != tt.want {
			t.Errorf("HasBlank(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestHasEmpty(t *testing.T) {
	tests := []struct {
		input []string
		want  bool
	}{
		{[]string{}, false},
		{[]string{"hello"}, false},
		{[]string{"hello", "world"}, false},
		{[]string{""}, true},
		{[]string{" "}, false},
		{[]string{"hello", "", "world"}, true},
	}

	for _, tt := range tests {
		if got := HasEmpty(tt.input...); got != tt.want {
			t.Errorf("HasEmpty(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

// ==================== Trim Functions ====================

func TestTrim(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{" ", ""},
		{"  hello  ", "hello"},
		{"\t\nhello\t\n", "hello"},
		{"hello", "hello"},
	}

	for _, tt := range tests {
		if got := Trim(tt.input); got != tt.want {
			t.Errorf("Trim(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestTrimPrefix(t *testing.T) {
	tests := []struct {
		s, prefix, want string
	}{
		{"", "", ""},
		{"hello", "hello", ""},
		{"hello world", "hello", " world"},
		{"hello world", "world", "hello world"},
		{"hello", "HELLO", "hello"},
	}

	for _, tt := range tests {
		if got := TrimPrefix(tt.s, tt.prefix); got != tt.want {
			t.Errorf("TrimPrefix(%q, %q) = %q, want %q", tt.s, tt.prefix, got, tt.want)
		}
	}
}

func TestTrimSuffix(t *testing.T) {
	tests := []struct {
		s, suffix, want string
	}{
		{"", "", ""},
		{"hello", "hello", ""},
		{"hello world", "world", "hello "},
		{"hello world", "hello", "hello world"},
		{"hello", "WORLD", "hello"},
	}

	for _, tt := range tests {
		if got := TrimSuffix(tt.s, tt.suffix); got != tt.want {
			t.Errorf("TrimSuffix(%q, %q) = %q, want %q", tt.s, tt.suffix, got, tt.want)
		}
	}
}

// ==================== Substring Functions ====================

func TestSub(t *testing.T) {
	tests := []struct {
		s          string
		start, end int
		want       string
	}{
		{"hello", 0, 5, "hello"},
		{"hello", 1, 4, "ell"},
		{"hello", 0, 3, "hel"},
		{"hello", 2, 2, ""},
		{"hello", -3, -1, "ll"},
		{"hello", -3, 5, "llo"},
		{"hello", 0, 10, "hello"},
		{"hello", 5, 3, ""},
		{"hello", -10, 2, "he"},
		{"", 0, 0, ""},
	}

	for _, tt := range tests {
		if got := Sub(tt.s, tt.start, tt.end); got != tt.want {
			t.Errorf("Sub(%q, %d, %d) = %q, want %q", tt.s, tt.start, tt.end, got, tt.want)
		}
	}
}

func TestSubBefore(t *testing.T) {
	tests := []struct {
		s, separator string
		isLast       bool
		want         string
	}{
		{"", ".", false, ""},
		{"hello.world", ".", false, "hello"},
		{"hello.world.go", ".", false, "hello"},
		{"hello.world.go", ".", true, "hello.world"},
		{"hello", ".", false, "hello"},
		{"hello", ".", true, "hello"},
		{"a.b.c.d", ".", true, "a.b.c"},
	}

	for _, tt := range tests {
		if got := SubBefore(tt.s, tt.separator, tt.isLast); got != tt.want {
			t.Errorf("SubBefore(%q, %q, %v) = %q, want %q", tt.s, tt.separator, tt.isLast, got, tt.want)
		}
	}
}

func TestSubAfter(t *testing.T) {
	tests := []struct {
		s, separator string
		isLast       bool
		want         string
	}{
		{"", ".", false, ""},
		{"hello.world", ".", false, "world"},
		{"hello.world.go", ".", false, "world.go"},
		{"hello.world.go", ".", true, "go"},
		{"hello", ".", false, ""},
		{"hello", ".", true, ""},
		{"a.b.c.d", ".", true, "d"},
	}

	for _, tt := range tests {
		if got := SubAfter(tt.s, tt.separator, tt.isLast); got != tt.want {
			t.Errorf("SubAfter(%q, %q, %v) = %q, want %q", tt.s, tt.separator, tt.isLast, got, tt.want)
		}
	}
}

// ==================== Contains Functions ====================

func TestContains(t *testing.T) {
	tests := []struct {
		s, sub string
		want   bool
	}{
		{"", "", true},
		{"hello", "", true},
		{"", "hello", false},
		{"hello world", "world", true},
		{"hello world", "World", false},
		{"hello world", "xyz", false},
	}

	for _, tt := range tests {
		if got := Contains(tt.s, tt.sub); got != tt.want {
			t.Errorf("Contains(%q, %q) = %v, want %v", tt.s, tt.sub, got, tt.want)
		}
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		s, sub string
		want   bool
	}{
		{"", "", true},
		{"hello", "", true},
		{"", "hello", false},
		{"Hello World", "world", true},
		{"Hello World", "WORLD", true},
		{"Hello World", "xyz", false},
	}

	for _, tt := range tests {
		if got := ContainsIgnoreCase(tt.s, tt.sub); got != tt.want {
			t.Errorf("ContainsIgnoreCase(%q, %q) = %v, want %v", tt.s, tt.sub, got, tt.want)
		}
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		s, prefix string
		want      bool
	}{
		{"", "", true},
		{"hello", "", true},
		{"", "hello", false},
		{"hello world", "hello", true},
		{"hello world", "world", false},
		{"hello world", "HELLO", false},
	}

	for _, tt := range tests {
		if got := StartsWith(tt.s, tt.prefix); got != tt.want {
			t.Errorf("StartsWith(%q, %q) = %v, want %v", tt.s, tt.prefix, got, tt.want)
		}
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		s, suffix string
		want      bool
	}{
		{"", "", true},
		{"hello", "", true},
		{"", "hello", false},
		{"hello world", "world", true},
		{"hello world", "hello", false},
		{"hello world", "WORLD", false},
	}

	for _, tt := range tests {
		if got := EndsWith(tt.s, tt.suffix); got != tt.want {
			t.Errorf("EndsWith(%q, %q) = %v, want %v", tt.s, tt.suffix, got, tt.want)
		}
	}
}

// ==================== Remove Functions ====================

func TestRemovePrefix(t *testing.T) {
	tests := []struct {
		s, prefix string
		want      string
	}{
		{"", "", ""},
		{"hello", "hello", ""},
		{"hello world", "hello", " world"},
		{"hello world", "world", "hello world"},
		{"hello", "HELLO", "hello"},
	}

	for _, tt := range tests {
		if got := RemovePrefix(tt.s, tt.prefix); got != tt.want {
			t.Errorf("RemovePrefix(%q, %q) = %q, want %q", tt.s, tt.prefix, got, tt.want)
		}
	}
}

func TestRemoveSuffix(t *testing.T) {
	tests := []struct {
		s, suffix string
		want      string
	}{
		{"", "", ""},
		{"hello", "hello", ""},
		{"hello world", "world", "hello "},
		{"hello world", "hello", "hello world"},
		{"hello", "WORLD", "hello"},
	}

	for _, tt := range tests {
		if got := RemoveSuffix(tt.s, tt.suffix); got != tt.want {
			t.Errorf("RemoveSuffix(%q, %q) = %q, want %q", tt.s, tt.suffix, got, tt.want)
		}
	}
}

// ==================== Pad Functions ====================

func TestPad(t *testing.T) {
	tests := []struct {
		s       string
		size    int
		padChar rune
		want    string
	}{
		{"hello", 10, '*', "**hello***"},
		{"hello", 5, '*', "hello"},
		{"hello", 3, '*', "hello"},
		{"hello", 0, '*', "hello"},
		{"", 5, '*', "*****"},
		{"hi", 6, '-', "--hi--"},
	}

	for _, tt := range tests {
		if got := Pad(tt.s, tt.size, tt.padChar); got != tt.want {
			t.Errorf("Pad(%q, %d, %q) = %q, want %q", tt.s, tt.size, tt.padChar, got, tt.want)
		}
	}
}

func TestPadLeft(t *testing.T) {
	tests := []struct {
		s       string
		size    int
		padChar rune
		want    string
	}{
		{"hello", 10, '*', "*****hello"},
		{"hello", 5, '*', "hello"},
		{"hello", 3, '*', "hello"},
		{"hello", 0, '*', "hello"},
		{"", 5, '*', "*****"},
		{"hi", 6, '-', "----hi"},
	}

	for _, tt := range tests {
		if got := PadLeft(tt.s, tt.size, tt.padChar); got != tt.want {
			t.Errorf("PadLeft(%q, %d, %q) = %q, want %q", tt.s, tt.size, tt.padChar, got, tt.want)
		}
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		s       string
		size    int
		padChar rune
		want    string
	}{
		{"hello", 10, '*', "hello*****"},
		{"hello", 5, '*', "hello"},
		{"hello", 3, '*', "hello"},
		{"hello", 0, '*', "hello"},
		{"", 5, '*', "*****"},
		{"hi", 6, '-', "hi----"},
	}

	for _, tt := range tests {
		if got := PadRight(tt.s, tt.size, tt.padChar); got != tt.want {
			t.Errorf("PadRight(%q, %d, %q) = %q, want %q", tt.s, tt.size, tt.padChar, got, tt.want)
		}
	}
}

// ==================== String Manipulation Functions ====================

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"a", "a"},
		{"hello", "olleh"},
		{"12345", "54321"},
		{"Hello World", "dlroW olleH"},
	}

	for _, tt := range tests {
		if got := Reverse(tt.input); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		s     string
		count int
		want  string
	}{
		{"hello", 0, ""},
		{"hello", -1, ""},
		{"hello", 1, "hello"},
		{"hello", 3, "hellohellohello"},
		{"", 5, ""},
		{"ab", 4, "abababab"},
	}

	for _, tt := range tests {
		if got := Repeat(tt.s, tt.count); got != tt.want {
			t.Errorf("Repeat(%q, %d) = %q, want %q", tt.s, tt.count, got, tt.want)
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		separator string
		strs      []string
		want      string
	}{
		{",", []string{}, ""},
		{",", []string{"hello"}, "hello"},
		{",", []string{"hello", "world"}, "hello,world"},
		{"-", []string{"a", "b", "c"}, "a-b-c"},
		{"", []string{"a", "b", "c"}, "abc"},
	}

	for _, tt := range tests {
		if got := Join(tt.separator, tt.strs...); got != tt.want {
			t.Errorf("Join(%q, %v) = %q, want %q", tt.separator, tt.strs, got, tt.want)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		s, separator string
		want         []string
	}{
		{"", ",", []string{""}},
		{"hello", ",", []string{"hello"}},
		{"hello,world", ",", []string{"hello", "world"}},
		{"a,b,c", ",", []string{"a", "b", "c"}},
		{"a,,b", ",", []string{"a", "", "b"}},
	}

	for _, tt := range tests {
		got := Split(tt.s, tt.separator)
		if len(got) != len(tt.want) {
			t.Errorf("Split(%q, %q) = %v, want %v", tt.s, tt.separator, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("Split(%q, %q)[%d] = %q, want %q", tt.s, tt.separator, i, got[i], tt.want[i])
			}
		}
	}
}

// ==================== Case Functions ====================

func TestUpperFirst(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"hELLO", "HELLO"},
		{"a", "A"},
	}

	for _, tt := range tests {
		if got := UpperFirst(tt.input); got != tt.want {
			t.Errorf("UpperFirst(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestLowerFirst(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"Hello", "hello"},
		{"hello", "hello"},
		{"HELLO", "hELLO"},
		{"A", "a"},
	}

	for _, tt := range tests {
		if got := LowerFirst(tt.input); got != tt.want {
			t.Errorf("LowerFirst(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestCamelToUnderline(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"hello", "hello"},
		{"helloWorld", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"myVariableName", "my_variable_name"},
		{"ABC", "a_b_c"},
		{"getHTTPResponse", "get_h_t_t_p_response"},
	}

	for _, tt := range tests {
		if got := CamelToUnderline(tt.input); got != tt.want {
			t.Errorf("CamelToUnderline(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestUnderlineToCamel(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"hello", "hello"},
		{"hello_world", "helloWorld"},
		{"my_variable_name", "myVariableName"},
		{"a_b_c", "aBC"},
		{"_hello", "Hello"},
		{"hello_", "hello"},
	}

	for _, tt := range tests {
		if got := UnderlineToCamel(tt.input); got != tt.want {
			t.Errorf("UnderlineToCamel(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

// ==================== Format Function ====================

func TestFormat(t *testing.T) {
	tests := []struct {
		template string
		args     []interface{}
		want     string
	}{
		{"", nil, ""},
		{"hello", nil, "hello"},
		{"Hello {0}!", []interface{}{"World"}, "Hello World!"},
		{"{0} + {1} = {2}", []interface{}{1, 2, 3}, "1 + 2 = 3"},
		{"{0}{1}{0}", []interface{}{"a", "b"}, "aba"},
		{"no placeholders", []interface{}{"arg"}, "no placeholders"},
		{"{0} {1} {2}", []interface{}{"a", "b"}, "a b {2}"},
	}

	for _, tt := range tests {
		if got := Format(tt.template, tt.args...); got != tt.want {
			t.Errorf("Format(%q, %v) = %q, want %q", tt.template, tt.args, got, tt.want)
		}
	}
}

// ==================== Count Function ====================

func TestCount(t *testing.T) {
	tests := []struct {
		s, sub string
		want   int
	}{
		{"", "", 0},
		{"hello", "", 0},
		{"", "hello", 0},
		{"hello", "l", 2},
		{"hello world", "o", 2},
		{"aaaa", "aa", 2},
		{"hello", "xyz", 0},
	}

	for _, tt := range tests {
		if got := Count(tt.s, tt.sub); got != tt.want {
			t.Errorf("Count(%q, %q) = %d, want %d", tt.s, tt.sub, got, tt.want)
		}
	}
}

// ==================== Wrap/Unwrap Functions ====================

func TestWrap(t *testing.T) {
	tests := []struct {
		s, prefix, suffix string
		want              string
	}{
		{"", "", "", ""},
		{"hello", "[", "]", "[hello]"},
		{"hello", "--", "--", "--hello--"},
		{"hello", "(", ")", "(hello)"},
		{"hello", "", "!", "hello!"},
	}

	for _, tt := range tests {
		if got := Wrap(tt.s, tt.prefix, tt.suffix); got != tt.want {
			t.Errorf("Wrap(%q, %q, %q) = %q, want %q", tt.s, tt.prefix, tt.suffix, got, tt.want)
		}
	}
}

func TestUnwrap(t *testing.T) {
	tests := []struct {
		s, prefix, suffix string
		want              string
	}{
		{"", "", "", ""},
		{"[hello]", "[", "]", "hello"},
		{"--hello--", "--", "--", "hello"},
		{"(hello)", "(", ")", "hello"},
		{"hello!", "", "!", "hello"},
		{"hello", "[", "]", "hello"},
		{"[hello", "[", "]", "hello"},
	}

	for _, tt := range tests {
		if got := Unwrap(tt.s, tt.prefix, tt.suffix); got != tt.want {
			t.Errorf("Unwrap(%q, %q, %q) = %q, want %q", tt.s, tt.prefix, tt.suffix, got, tt.want)
		}
	}
}

// ==================== Default Functions ====================

func TestDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		s, defaultStr, want string
	}{
		{"", "default", "default"},
		{"hello", "default", "hello"},
		{" ", "default", " "},
		{"", "", ""},
	}

	for _, tt := range tests {
		if got := DefaultIfEmpty(tt.s, tt.defaultStr); got != tt.want {
			t.Errorf("DefaultIfEmpty(%q, %q) = %q, want %q", tt.s, tt.defaultStr, got, tt.want)
		}
	}
}

func TestDefaultIfBlank(t *testing.T) {
	tests := []struct {
		s, defaultStr, want string
	}{
		{"", "default", "default"},
		{" ", "default", "default"},
		{"  \t  ", "default", "default"},
		{"hello", "default", "hello"},
		{"", "", ""},
	}

	for _, tt := range tests {
		if got := DefaultIfBlank(tt.s, tt.defaultStr); got != tt.want {
			t.Errorf("DefaultIfBlank(%q, %q) = %q, want %q", tt.s, tt.defaultStr, got, tt.want)
		}
	}
}

// ==================== Builder Tests ====================

func TestBuilderNew(t *testing.T) {
	b := NewBuilder()
	if !b.IsEmpty() {
		t.Error("New builder should be empty")
	}
	if b.Len() != 0 {
		t.Error("New builder should have length 0")
	}
	if b.String() != "" {
		t.Error("New builder should return empty string")
	}
}

func TestBuilderNewSize(t *testing.T) {
	b := NewBuilderSize(100)
	if !b.IsEmpty() {
		t.Error("New builder should be empty")
	}
	if b.Len() != 0 {
		t.Error("New builder should have length 0")
	}
}

func TestBuilderAppend(t *testing.T) {
	b := NewBuilder()
	result := b.Append("hello").Append(" ").Append("world").String()
	if result != "hello world" {
		t.Errorf("Builder.Append() = %q, want %q", result, "hello world")
	}
}

func TestBuilderAppendByte(t *testing.T) {
	b := NewBuilder()
	result := b.AppendByte('A').AppendByte('B').AppendByte('C').String()
	if result != "ABC" {
		t.Errorf("Builder.AppendByte() = %q, want %q", result, "ABC")
	}
}

func TestBuilderAppendRune(t *testing.T) {
	b := NewBuilder()
	result := b.AppendRune('你').AppendRune('好').String()
	if result != "你好" {
		t.Errorf("Builder.AppendRune() = %q, want %q", result, "你好")
	}
}

func TestBuilderAppendFormat(t *testing.T) {
	b := NewBuilder()
	result := b.AppendFormat("Hello %s, you are %d", "World", 25).String()
	if result != "Hello World, you are 25" {
		t.Errorf("Builder.AppendFormat() = %q, want %q", result, "Hello World, you are 25")
	}
}

func TestBuilderInsert(t *testing.T) {
	tests := []struct {
		initial string
		index   int
		s       string
		want    string
	}{
		{"hello", 0, "XX", "XXhello"},
		{"hello", 5, "XX", "helloXX"},
		{"hello", 3, "XX", "helXXlo"},
		{"hello", -1, "XX", "XXhello"},
		{"hello", 10, "XX", "helloXX"},
		{"", 0, "hello", "hello"},
	}

	for _, tt := range tests {
		b := NewBuilder()
		b.Append(tt.initial)
		result := b.Insert(tt.index, tt.s).String()
		if result != tt.want {
			t.Errorf("Builder(%q).Insert(%d, %q) = %q, want %q", tt.initial, tt.index, tt.s, result, tt.want)
		}
	}
}

func TestBuilderDelete(t *testing.T) {
	tests := []struct {
		initial    string
		start, end int
		want       string
	}{
		{"hello", 0, 2, "llo"},
		{"hello", 2, 5, "he"},
		{"hello", 1, 4, "ho"},
		{"hello", 0, 5, ""},
		{"hello", 3, 3, "hello"},
		{"hello", 5, 3, "hello"},
		{"hello", -1, 2, "llo"},
		{"hello", 3, 10, "hel"},
	}

	for _, tt := range tests {
		b := NewBuilder()
		b.Append(tt.initial)
		result := b.Delete(tt.start, tt.end).String()
		if result != tt.want {
			t.Errorf("Builder(%q).Delete(%d, %d) = %q, want %q", tt.initial, tt.start, tt.end, result, tt.want)
		}
	}
}

func TestBuilderReplace(t *testing.T) {
	tests := []struct {
		initial    string
		start, end int
		s          string
		want       string
	}{
		{"hello", 0, 2, "XX", "XXllo"},
		{"hello", 2, 5, "XX", "heXX"},
		{"hello", 1, 4, "XX", "hXXo"},
		{"hello", 0, 5, "XX", "XX"},
		{"hello", 3, 3, "XX", "helXXlo"},
		{"hello", 5, 3, "XX", "hello"},
	}

	for _, tt := range tests {
		b := NewBuilder()
		b.Append(tt.initial)
		result := b.Replace(tt.start, tt.end, tt.s).String()
		if result != tt.want {
			t.Errorf("Builder(%q).Replace(%d, %d, %q) = %q, want %q", tt.initial, tt.start, tt.end, tt.s, result, tt.want)
		}
	}
}

func TestBuilderLen(t *testing.T) {
	b := NewBuilder()
	if b.Len() != 0 {
		t.Errorf("Empty builder Len() = %d, want 0", b.Len())
	}

	b.Append("hello")
	if b.Len() != 5 {
		t.Errorf("Builder Len() after Append('hello') = %d, want 5", b.Len())
	}

	b.Append(" world")
	if b.Len() != 11 {
		t.Errorf("Builder Len() after Append(' world') = %d, want 11", b.Len())
	}
}

func TestBuilderReset(t *testing.T) {
	b := NewBuilder()
	b.Append("hello world")
	b.Reset()

	if !b.IsEmpty() {
		t.Error("Builder should be empty after Reset()")
	}
	if b.Len() != 0 {
		t.Error("Builder should have length 0 after Reset()")
	}
	if b.String() != "" {
		t.Error("Builder should return empty string after Reset()")
	}
}

func TestBuilderIsEmpty(t *testing.T) {
	b := NewBuilder()
	if !b.IsEmpty() {
		t.Error("New builder should be empty")
	}

	b.Append("hello")
	if b.IsEmpty() {
		t.Error("Builder should not be empty after Append")
	}

	b.Reset()
	if !b.IsEmpty() {
		t.Error("Builder should be empty after Reset")
	}
}

func TestBuilderFluentAPI(t *testing.T) {
	b := NewBuilder()
	result := b.
		Append("Hello").
		Append(" ").
		Append("World").
		Append("!").
		String()

	if result != "Hello World!" {
		t.Errorf("Fluent API result = %q, want %q", result, "Hello World!")
	}
}

func TestBuilderChaining(t *testing.T) {
	b := NewBuilderSize(50)
	result := b.
		Append("Start").
		Insert(0, "Begin ").
		Append(" End").
		Replace(0, 6, "New").
		String()

	if result != "NewStart End" {
		t.Errorf("Builder chaining result = %q, want %q", result, "NewStart End")
	}
}
