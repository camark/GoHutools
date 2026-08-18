package charutil

import "testing"

func TestIsNumber(t *testing.T) {
	if !IsNumber('5') {
		t.Error("IsNumber('5')")
	}
	if IsNumber('a') || IsNumber(' ') {
		t.Error("IsNumber non-digit")
	}
}

func TestIsLetter(t *testing.T) {
	if !IsLetter('a') || !IsLetter('Z') {
		t.Error("IsLetter latin")
	}
	if !IsLetter('汉') {
		t.Error("IsLetter CJK")
	}
	if IsLetter('1') || IsLetter('-') {
		t.Error("IsLetter non-letter")
	}
}

func TestIsLetterOrNumber(t *testing.T) {
	if !IsLetterOrNumber('a') || !IsLetterOrNumber('9') {
		t.Error("IsLetterOrNumber alnum")
	}
	if IsLetterOrNumber('_') {
		t.Error("IsLetterOrNumber underscore is not alnum")
	}
}

func TestIsBlankChar(t *testing.T) {
	blank := []rune{' ', '\t', '\n', '　'} // space, tab, LF, full-width space
	for _, c := range blank {
		if !IsBlankChar(c) {
			t.Errorf("IsBlankChar(%q)", c)
		}
	}
	if IsBlankChar('x') {
		t.Error("IsBlankChar('x')")
	}
}

func TestIsEmoji(t *testing.T) {
	if !IsEmoji('😀') {
		t.Error("IsEmoji(😀)")
	}
	if IsEmoji('a') {
		t.Error("IsEmoji('a')")
	}
}

func TestIsAscii(t *testing.T) {
	if !IsAscii('a') || !IsAscii(0x7F) {
		t.Error("IsAscii range")
	}
	if IsAscii('中') {
		t.Error("IsAscii(CJK)")
	}
}

func TestIsHexChar(t *testing.T) {
	for _, c := range "0123456789abcdefABCDEF" {
		if !IsHexChar(c) {
			t.Errorf("IsHexChar(%c)", c)
		}
	}
	if IsHexChar('g') || IsHexChar('G') {
		t.Error("IsHexChar non-hex")
	}
}

func TestIsUpperCaseIsLowerCase(t *testing.T) {
	if !IsUpperCase('A') || IsUpperCase('a') {
		t.Error("IsUpperCase")
	}
	if !IsLowerCase('a') || IsLowerCase('A') {
		t.Error("IsLowerCase")
	}
}

func TestEqualsAny(t *testing.T) {
	if !EqualsAny('x', 'a', 'b', 'x') {
		t.Error("EqualsAny present")
	}
	if EqualsAny('y', 'a', 'b', 'x') {
		t.Error("EqualsAny absent")
	}
	if EqualsAny('a') {
		t.Error("EqualsAny with no candidates should be false")
	}
}

func TestToAscii(t *testing.T) {
	if ToAscii('A') != 65 {
		t.Errorf("ToAscii(A) = %d", ToAscii('A'))
	}
}

func TestIsFileSeparator(t *testing.T) {
	if !IsFileSeparator('/') {
		t.Error("IsFileSeparator(/)")
	}
	if !IsFileSeparator('\\') {
		t.Error("IsFileSeparator(\\)")
	}
	if IsFileSeparator('a') {
		t.Error("IsFileSeparator(a)")
	}
}

func TestIsBlankString(t *testing.T) {
	if !IsBlank("") || !IsBlank("  \t\n") {
		t.Error("IsBlank blank")
	}
	if IsBlank(" x ") {
		t.Error("IsBlank non-blank")
	}
}