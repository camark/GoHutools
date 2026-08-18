package charutil

import (
	"strings"
	"unicode"
)

// IsNumber reports whether c is a digit (0-9).
func IsNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

// IsLetter reports whether c is a letter (latin or CJK or any unicode letter).
func IsLetter(c rune) bool {
	return unicode.IsLetter(c)
}

// IsLetterOrNumber reports whether c is alphanumeric.
func IsLetterOrNumber(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

// IsBlankChar reports whether c is a blank character:
// space, tab, line breaks, full-width space, zero-width space.
func IsBlankChar(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\r', '\f', '\v', '　', '​':
		return true
	}
	return false
}

// IsBlank reports whether s is empty or composed only of blank chars.
func IsBlank(s string) bool {
	for _, c := range s {
		if !IsBlankChar(c) {
			return false
		}
	}
	return true
}

// IsEmoji reports whether c falls inside the common emoji ranges.
func IsEmoji(c rune) bool {
	return (c >= 0x1F300 && c <= 0x1F64F) || // misc symbols & pictographs, emoticons
		(c >= 0x1F900 && c <= 0x1F9FF) || // supplemental symbols
		(c >= 0x2600 && c <= 0x27BF) || // misc symbols, dingbats
		(c >= 0xFE00 && c <= 0xFE0F) || // variation selectors
		(c >= 0x1F1E6 && c <= 0x1F1FF) // regional indicators
}

// IsAscii reports whether c is in the ASCII range (0x00-0x7F).
func IsAscii(c rune) bool {
	return c >= 0 && c <= 0x7F
}

// IsHexChar reports whether c is a hexadecimal digit.
func IsHexChar(c rune) bool {
	return IsNumber(c) || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// IsUpperCase reports whether c is an uppercase letter.
func IsUpperCase(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

// IsLowerCase reports whether c is a lowercase letter.
func IsLowerCase(c rune) bool {
	return c >= 'a' && c <= 'z'
}

// EqualsAny reports whether c equals at least one of the candidates.
func EqualsAny(c rune, candidates ...rune) bool {
	for _, cand := range candidates {
		if c == cand {
			return true
		}
	}
	return false
}

// ToAscii returns the ASCII code of c.
func ToAscii(c rune) int {
	return int(c)
}

// IsFileSeparator reports whether c is '/' or '\\'.
func IsFileSeparator(c rune) bool {
	return c == '/' || c == '\\'
}

// CountBlank counts blank characters in s.
func CountBlank(s string) int {
	n := 0
	for _, c := range s {
		if IsBlankChar(c) {
			n++
		}
	}
	return n
}

// RemoveBlanks removes all blank characters from s.
func RemoveBlanks(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for _, c := range s {
		if !IsBlankChar(c) {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}