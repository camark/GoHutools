package strutil

import (
	"fmt"
	"strings"
	"unicode"
)

// IsEmpty checks if string is empty
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsNotEmpty checks if string is not empty
func IsNotEmpty(s string) bool {
	return len(s) > 0
}

// IsBlank checks if string is blank (empty or only whitespace)
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotBlank checks if string is not blank
func IsNotBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}

// HasBlank checks if any string is blank
func HasBlank(strs ...string) bool {
	for _, s := range strs {
		if IsBlank(s) {
			return true
		}
	}
	return false
}

// HasEmpty checks if any string is empty
func HasEmpty(strs ...string) bool {
	for _, s := range strs {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}

// Trim removes leading and trailing whitespace
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// TrimPrefix removes prefix if present
func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// TrimSuffix removes suffix if present
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// Sub extracts substring from start to end (supports negative indices)
func Sub(s string, start, end int) string {
	runes := []rune(s)
	length := len(runes)

	// Handle negative indices
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}

	// Clamp to valid range
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if end > length {
		end = length
	}
	if start > length {
		start = length
	}
	if start > end {
		start = end
	}

	return string(runes[start:end])
}

// SubBefore gets substring before separator
func SubBefore(s, separator string, isLastSeparator bool) string {
	if IsEmpty(s) || IsEmpty(separator) {
		return s
	}

	var index int
	if isLastSeparator {
		index = strings.LastIndex(s, separator)
	} else {
		index = strings.Index(s, separator)
	}

	if index < 0 {
		return s
	}

	return s[:index]
}

// SubAfter gets substring after separator
func SubAfter(s, separator string, isLastSeparator bool) string {
	if IsEmpty(s) || IsEmpty(separator) {
		return s
	}

	var index int
	if isLastSeparator {
		index = strings.LastIndex(s, separator)
	} else {
		index = strings.Index(s, separator)
	}

	if index < 0 {
		return ""
	}

	return s[index+len(separator):]
}

// Contains checks if string contains substring
func Contains(s, sub string) bool {
	return strings.Contains(s, sub)
}

// ContainsIgnoreCase checks if string contains substring (case insensitive)
func ContainsIgnoreCase(s, sub string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(sub))
}

// StartsWith checks if string starts with prefix
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// EndsWith checks if string ends with suffix
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// RemovePrefix removes prefix if starts with it
func RemovePrefix(s, prefix string) string {
	if StartsWith(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

// RemoveSuffix removes suffix if ends with it
func RemoveSuffix(s, suffix string) string {
	if EndsWith(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

// Pad pads string to size with padChar (center padding)
func Pad(s string, size int, padChar rune) string {
	if size <= 0 {
		return s
	}

	currentLen := len([]rune(s))
	if currentLen >= size {
		return s
	}

	padLen := size - currentLen
	padLeft := padLen / 2
	padRight := padLen - padLeft

	var sb strings.Builder
	for i := 0; i < padLeft; i++ {
		sb.WriteRune(padChar)
	}
	sb.WriteString(s)
	for i := 0; i < padRight; i++ {
		sb.WriteRune(padChar)
	}

	return sb.String()
}

// PadLeft pads string on the left
func PadLeft(s string, size int, padChar rune) string {
	if size <= 0 {
		return s
	}

	currentLen := len([]rune(s))
	if currentLen >= size {
		return s
	}

	padLen := size - currentLen
	var sb strings.Builder
	for i := 0; i < padLen; i++ {
		sb.WriteRune(padChar)
	}
	sb.WriteString(s)

	return sb.String()
}

// PadRight pads string on the right
func PadRight(s string, size int, padChar rune) string {
	if size <= 0 {
		return s
	}

	currentLen := len([]rune(s))
	if currentLen >= size {
		return s
	}

	padLen := size - currentLen
	var sb strings.Builder
	sb.WriteString(s)
	for i := 0; i < padLen; i++ {
		sb.WriteRune(padChar)
	}

	return sb.String()
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Repeat repeats string count times
func Repeat(s string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(s, count)
}

// Join joins strings with separator
func Join(separator string, strs ...string) string {
	return strings.Join(strs, separator)
}

// Split splits string by separator
func Split(s, separator string) []string {
	return strings.Split(s, separator)
}

// UpperFirst capitalizes first letter
func UpperFirst(s string) string {
	if IsEmpty(s) {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// LowerFirst lowercases first letter
func LowerFirst(s string) string {
	if IsEmpty(s) {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// CamelToUnderline converts camelCase to under_score
func CamelToUnderline(s string) string {
	if IsEmpty(s) {
		return s
	}

	var sb strings.Builder
	runes := []rune(s)

	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// UnderlineToCamel converts under_score to camelCase
func UnderlineToCamel(s string) string {
	if IsEmpty(s) {
		return s
	}

	var sb strings.Builder
	upperNext := false

	for _, r := range s {
		if r == '_' {
			upperNext = true
		} else {
			if upperNext {
				sb.WriteRune(unicode.ToUpper(r))
				upperNext = false
			} else {
				sb.WriteRune(r)
			}
		}
	}

	return sb.String()
}

// Format replaces {0}, {1}, etc. with args
func Format(template string, args ...interface{}) string {
	if IsEmpty(template) || len(args) == 0 {
		return template
	}

	result := template
	for i, arg := range args {
		placeholder := fmt.Sprintf("{%d}", i)
		result = strings.ReplaceAll(result, placeholder, fmt.Sprintf("%v", arg))
	}

	return result
}

// Count counts occurrences of sub in s
func Count(s, sub string) int {
	if IsEmpty(s) || IsEmpty(sub) {
		return 0
	}
	return strings.Count(s, sub)
}

// Wrap wraps string with prefix and suffix
func Wrap(s, prefix, suffix string) string {
	return prefix + s + suffix
}

// Unwrap removes prefix and suffix if present
func Unwrap(s, prefix, suffix string) string {
	if IsEmpty(s) {
		return s
	}

	result := s
	if StartsWith(result, prefix) {
		result = result[len(prefix):]
	}
	if EndsWith(result, suffix) {
		result = result[:len(result)-len(suffix)]
	}

	return result
}

// DefaultIfEmpty returns defaultStr if s is empty
func DefaultIfEmpty(s, defaultStr string) string {
	if IsEmpty(s) {
		return defaultStr
	}
	return s
}

// DefaultIfBlank returns defaultStr if s is blank
func DefaultIfBlank(s, defaultStr string) string {
	if IsBlank(s) {
		return defaultStr
	}
	return s
}
