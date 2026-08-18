// Package regexutil provides regex utility functions similar to Java Hutool's ReUtil.
package regexutil

import (
	"regexp"
	"sync"
)

// cache for compiled regex patterns
var (
	regexCache = make(map[string]*regexp.Regexp)
	cacheMu    sync.RWMutex
)

// getRegex returns compiled regex from cache or compiles new one
func getRegex(pattern string) (*regexp.Regexp, error) {
	cacheMu.RLock()
	if re, ok := regexCache[pattern]; ok {
		cacheMu.RUnlock()
		return re, nil
	}
	cacheMu.RUnlock()

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	cacheMu.Lock()
	regexCache[pattern] = re
	cacheMu.Unlock()

	return re, nil
}

// MustGetRegex returns compiled regex, panics on error
func MustGetRegex(pattern string) *regexp.Regexp {
	re, err := getRegex(pattern)
	if err != nil {
		panic(err)
	}
	return re
}

// IsMatch checks if string matches pattern
func IsMatch(s string, pattern string) bool {
	re, err := getRegex(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(s)
}

// IsMatchBytes checks if bytes match pattern
func IsMatchBytes(b []byte, pattern string) bool {
	re, err := getRegex(pattern)
	if err != nil {
		return false
	}
	return re.Match(b)
}

// Find finds first match
func Find(s string, pattern string) string {
	re, err := getRegex(pattern)
	if err != nil {
		return ""
	}
	return re.FindString(s)
}

// FindBytes finds first match in bytes
func FindBytes(b []byte, pattern string) []byte {
	re, err := getRegex(pattern)
	if err != nil {
		return nil
	}
	return re.Find(b)
}

// FindAll finds all matches
func FindAll(s string, pattern string) []string {
	re, err := getRegex(pattern)
	if err != nil {
		return nil
	}
	return re.FindAllString(s, -1)
}

// FindAllBytes finds all matches in bytes
func FindAllBytes(b []byte, pattern string) [][]byte {
	re, err := getRegex(pattern)
	if err != nil {
		return nil
	}
	return re.FindAll(b, -1)
}

// FindAllWithIndex finds all matches with index
func FindAllWithIndex(s string, pattern string) [][]int {
	re, err := getRegex(pattern)
	if err != nil {
		return nil
	}
	return re.FindAllStringIndex(s, -1)
}

// FindGroups finds first match and returns groups
func FindGroups(s string, pattern string) []string {
	re, err := getRegex(pattern)
	if err != nil {
		return nil
	}
	return re.FindStringSubmatch(s)
}

// FindAllGroups finds all matches and returns groups
func FindAllGroups(s string, pattern string) [][]string {
	re, err := getRegex(pattern)
	if err != nil {
		return nil
	}
	return re.FindAllStringSubmatch(s, -1)
}

// FindNamedGroups finds first match and returns named groups
func FindNamedGroups(s string, pattern string) map[string]string {
	re, err := getRegex(pattern)
	if err != nil {
		return nil
	}
	match := re.FindStringSubmatch(s)
	if match == nil {
		return nil
	}
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}

// Replace replaces all matches with replacement
func Replace(s string, pattern string, replacement string) string {
	re, err := getRegex(pattern)
	if err != nil {
		return s
	}
	return re.ReplaceAllString(s, replacement)
}

// ReplaceBytes replaces all matches in bytes
func ReplaceBytes(b []byte, pattern string, replacement []byte) []byte {
	re, err := getRegex(pattern)
	if err != nil {
		return b
	}
	return re.ReplaceAll(b, replacement)
}

// ReplaceFunc replaces all matches using function
func ReplaceFunc(s string, pattern string, replacer func(string) string) string {
	re, err := getRegex(pattern)
	if err != nil {
		return s
	}
	return re.ReplaceAllStringFunc(s, replacer)
}

// ReplaceFirst replaces first match
func ReplaceFirst(s string, pattern string, replacement string) string {
	re, err := getRegex(pattern)
	if err != nil {
		return s
	}
	match := re.FindStringSubmatchIndex(s)
	if match == nil {
		return s
	}
	repl := re.ExpandString(nil, replacement, s, match)
	return s[:match[0]] + string(repl) + s[match[1]:]
}

// Split splits string by pattern
func Split(s string, pattern string) []string {
	re, err := getRegex(pattern)
	if err != nil {
		return []string{s}
	}
	return re.Split(s, -1)
}

// Count counts matches
func Count(s string, pattern string) int {
	re, err := getRegex(pattern)
	if err != nil {
		return 0
	}
	return len(re.FindAllString(s, -1))
}

// ExtractGroup extracts specific group by index
func ExtractGroup(s string, pattern string, groupIndex int) string {
	groups := FindGroups(s, pattern)
	if groups == nil || groupIndex < 0 || groupIndex >= len(groups) {
		return ""
	}
	return groups[groupIndex]
}

// ExtractNamedGroup extracts named group
func ExtractNamedGroup(s string, pattern string, name string) string {
	groups := FindNamedGroups(s, pattern)
	if groups == nil {
		return ""
	}
	return groups[name]
}

// MatchAll checks if entire string matches pattern
func MatchAll(s string, pattern string) bool {
	re, err := getRegex(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(s) && re.FindString(s) == s
}

// ContainsMatch checks if string contains pattern
func ContainsMatch(s string, pattern string) bool {
	return IsMatch(s, pattern)
}

// Quote quotes special regex characters
func Quote(s string) string {
	return regexp.QuoteMeta(s)
}

// ClearCache clears regex cache
func ClearCache() {
	cacheMu.Lock()
	regexCache = make(map[string]*regexp.Regexp)
	cacheMu.Unlock()
}

// CacheSize returns cache size
func CacheSize() int {
	cacheMu.RLock()
	defer cacheMu.RUnlock()
	return len(regexCache)
}

// Common patterns
const (
	PatternEmail      = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	PatternURL        = `^(https?|ftp)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]$`
	PatternIP         = `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`
	PatternIPv4       = `^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`
	PatternMobile     = `^1[3-9]\d{9}$`
	PatternPhone      = `^0\d{2,3}-?\d{7,8}$`
	PatternIDCard     = `^\d{17}[\dXx]$`
	PatternZipCode    = `^\d{6}$`
	PatternAlpha      = `^[a-zA-Z]+$`
	PatternNumeric    = `^\d+$`
	PatternAlphaNum   = `^[a-zA-Z0-9]+$`
	PatternChinese    = `^[\x{4e00}-\x{9fa5}]+$`
	PatternDate       = `^\d{4}-\d{2}-\d{2}$`
	PatternDateTime   = `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`
	PatternUUID       = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
	PatternHexColor   = `^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`
	PatternCreditCard = `^\d{13,19}$`
)

// IsEmail checks if string is valid email
func IsEmail(s string) bool {
	return IsMatch(s, PatternEmail)
}

// IsURL checks if string is valid URL
func IsURL(s string) bool {
	return IsMatch(s, PatternURL)
}

// IsIP checks if string is valid IP
func IsIP(s string) bool {
	return IsMatch(s, PatternIPv4)
}

// IsMobile checks if string is valid mobile number
func IsMobile(s string) bool {
	return IsMatch(s, PatternMobile)
}

// IsPhone checks if string is valid phone number
func IsPhone(s string) bool {
	return IsMatch(s, PatternPhone)
}

// IsIDCard checks if string is valid ID card
func IsIDCard(s string) bool {
	return IsMatch(s, PatternIDCard)
}

// IsZipCode checks if string is valid zip code
func IsZipCode(s string) bool {
	return IsMatch(s, PatternZipCode)
}

// IsAlpha checks if string contains only letters
func IsAlpha(s string) bool {
	return IsMatch(s, PatternAlpha)
}

// IsNumeric checks if string contains only numbers
func IsNumeric(s string) bool {
	return IsMatch(s, PatternNumeric)
}

// IsAlphaNumeric checks if string contains only letters and numbers
func IsAlphaNumeric(s string) bool {
	return IsMatch(s, PatternAlphaNum)
}

// IsChinese checks if string contains only Chinese characters
func IsChinese(s string) bool {
	return IsMatch(s, PatternChinese)
}

// ContainsChinese checks if string contains Chinese characters
func ContainsChinese(s string) bool {
	return IsMatch(s, `[\x{4e00}-\x{9fa5}]`)
}

// IsDate checks if string is valid date
func IsDate(s string) bool {
	return IsMatch(s, PatternDate)
}

// IsDateTime checks if string is valid datetime
func IsDateTime(s string) bool {
	return IsMatch(s, PatternDateTime)
}

// IsUUID checks if string is valid UUID
func IsUUID(s string) bool {
	// UUID pattern is case-insensitive
	return IsMatch(s, `(?i)`+PatternUUID)
}

// IsHexColor checks if string is valid hex color
func IsHexColor(s string) bool {
	return IsMatch(s, PatternHexColor)
}
