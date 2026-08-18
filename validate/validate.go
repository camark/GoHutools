package validate

import (
	"encoding/base64"
	"encoding/json"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Common regex patterns
var (
	emailRegex      = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	uuidRegex       = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	creditCardRegex = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|(?:2131|1800|35\d{3})\d{11})$`)
)

// IsEmpty checks if value is empty
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	switch val := v.(type) {
	case string:
		return val == ""
	case []byte:
		return len(val) == 0
	case []interface{}:
		return len(val) == 0
	case map[string]interface{}:
		return len(val) == 0
	case int:
		return val == 0
	case int8:
		return val == 0
	case int16:
		return val == 0
	case int32:
		return val == 0
	case int64:
		return val == 0
	case uint:
		return val == 0
	case uint8:
		return val == 0
	case uint16:
		return val == 0
	case uint32:
		return val == 0
	case uint64:
		return val == 0
	case float32:
		return val == 0
	case float64:
		return val == 0
	case bool:
		return !val
	default:
		// Handle concrete slice/map/array/chan types that the type switch
		// cannot match (e.g. []string, map[string]int)
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
			return rv.Len() == 0
		case reflect.Ptr, reflect.Interface:
			return rv.IsNil()
		}
		return false
	}
}

// IsNotEmpty checks if value is not empty
func IsNotEmpty(v interface{}) bool {
	return !IsEmpty(v)
}

// IsNull checks if value is nil
func IsNull(v interface{}) bool {
	return v == nil
}

// IsNotNull checks if value is not nil
func IsNotNull(v interface{}) bool {
	return v != nil
}

// IsBlank checks if string is blank
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotBlank checks if string is not blank
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsEmail checks if string is valid email
func IsEmail(s string) bool {
	if len(s) > 254 {
		return false
	}
	return emailRegex.MatchString(s)
}

// IsURL checks if string is valid URL
func IsURL(s string) bool {
	if s == "" {
		return false
	}

	// Same pattern as regexutil.PatternURL
	urlRegex := regexp.MustCompile(`^(https?|ftp)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]$`)
	return urlRegex.MatchString(strings.ToLower(s))
}

// IsIP checks if string is valid IP address
func IsIP(s string) bool {
	return net.ParseIP(s) != nil
}

// IsIPv4 checks if string is valid IPv4
func IsIPv4(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	// Check if it's IPv4 by checking if it can be represented as 4 bytes
	return ip.To4() != nil
}

// IsIPv6 checks if string is valid IPv6
func IsIPv6(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	// IPv6 if it has 16 bytes but cannot be represented as 4 bytes
	return ip.To4() == nil && ip.To16() != nil
}

// IsMAC checks if string is valid MAC address
func IsMAC(s string) bool {
	_, err := net.ParseMAC(s)
	return err == nil
}

// IsPort checks if number is valid port
func IsPort(port int) bool {
	return port >= 0 && port <= 65535
}

// IsMobile checks if string is valid Chinese mobile number
func IsMobile(s string) bool {
	if s == "" {
		return false
	}

	// Remove country code if present
	s = strings.TrimPrefix(s, "+86")
	s = strings.TrimPrefix(s, "86")
	s = strings.TrimSpace(s)

	// Chinese mobile numbers: 1[3-9]X XXXX XXXX (11 digits)
	mobileRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return mobileRegex.MatchString(s)
}

// IsPhone checks if string is valid Chinese phone number
func IsPhone(s string) bool {
	if s == "" {
		return false
	}

	// Remove country code if present
	s = strings.TrimPrefix(s, "+86")
	s = strings.TrimPrefix(s, "86")
	s = strings.TrimSpace(s)

	// Chinese phone numbers:
	// Landline: 0XX-XXXXXXXX (10-12 digits)
	//400/800: 400-XXX-XXXX or 800-XXX-XXXX
	phoneRegex := regexp.MustCompile(`^(?:(?:0\d{2,3})?[\-]?\d{7,8}|400[\-]?\d{3}[\-]?\d{4}|800[\-]?\d{3}[\-]?\d{4})$`)
	return phoneRegex.MatchString(s)
}

// IsIDCard checks if string is valid Chinese ID card number
func IsIDCard(s string) bool {
	if s == "" {
		return false
	}

	// Remove spaces
	s = strings.ReplaceAll(s, " ", "")

	// Chinese ID card: 18 digits (last can be X)
	idCardRegex := regexp.MustCompile(`^\d{17}[\dXx]$`)
	if !idCardRegex.MatchString(s) {
		return false
	}

	// Validate birth date
	year, _ := strconv.Atoi(s[6:10])
	month, _ := strconv.Atoi(s[10:12])
	day, _ := strconv.Atoi(s[12:14])

	if year < 1900 || year > 2100 {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}

	// Validate checksum
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkCodes := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

	sum := 0
	for i := 0; i < 17; i++ {
		digit, _ := strconv.Atoi(string(s[i]))
		sum += digit * weights[i]
	}

	expectedCheck := checkCodes[sum%11]
	actualCheck := byte(unicode.ToUpper(rune(s[17])))

	return actualCheck == expectedCheck
}

// IsZipCode checks if string is valid zip code
func IsZipCode(s string) bool {
	if s == "" {
		return false
	}

	// Chinese zip code: 6 digits
	zipRegex := regexp.MustCompile(`^\d{6}$`)
	return zipRegex.MatchString(s)
}

// IsAlpha checks if string contains only letters
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsAlphaNumeric checks if string contains only letters and numbers
func IsAlphaNumeric(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsNumeric checks if string contains only numbers
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsInteger checks if string is integer
func IsInteger(s string) bool {
	if s == "" {
		return false
	}

	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	// Allow optional sign
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}

	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsFloat checks if string is float
func IsFloat(s string) bool {
	if s == "" {
		return false
	}

	s = strings.TrimSpace(s)
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsUUID checks if string is valid UUID
func IsUUID(s string) bool {
	return uuidRegex.MatchString(s)
}

// IsJSON checks if string is valid JSON
func IsJSON(s string) bool {
	if s == "" {
		return false
	}

	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

// IsBase64 checks if string is valid base64
func IsBase64(s string) bool {
	if s == "" {
		return false
	}

	// Check if it's valid base64
	_, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		// Try with padding
		_, err = base64.RawStdEncoding.DecodeString(s)
		return err == nil
	}
	return true
}

// IsHexColor checks if string is valid hex color
func IsHexColor(s string) bool {
	if s == "" {
		return false
	}

	// Remove # if present
	s = strings.TrimPrefix(s, "#")

	// Check length (3, 4, 6, or 8 hex digits)
	length := len(s)
	if length != 3 && length != 4 && length != 6 && length != 8 {
		return false
	}

	// Check if all characters are hex digits
	hexRegex := regexp.MustCompile(`^[0-9a-fA-F]+$`)
	return hexRegex.MatchString(s)
}

// IsRGBColor checks if string is valid RGB color
func IsRGBColor(s string) bool {
	if s == "" {
		return false
	}

	// Match rgb(r, g, b)
	rgbRegex := regexp.MustCompile(`^rgb\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(\d{1,3})\s*\)$`)
	// Match rgba(r, g, b, a)
	rgbaRegex := regexp.MustCompile(`^rgba\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(0(?:\.\d+)?|1(?:\.0+)?)\s*\)$`)

	var matches []string
	if matches = rgbaRegex.FindStringSubmatch(s); matches != nil {
		// Validate RGB values (0-255)
		for i := 1; i <= 3; i++ {
			val, _ := strconv.Atoi(matches[i])
			if val < 0 || val > 255 {
				return false
			}
		}
		// Validate alpha value (0-1)
		alpha, _ := strconv.ParseFloat(matches[4], 64)
		if alpha < 0 || alpha > 1 {
			return false
		}
		return true
	}

	if matches = rgbRegex.FindStringSubmatch(s); matches != nil {
		// Validate RGB values (0-255)
		for i := 1; i <= 3; i++ {
			val, _ := strconv.Atoi(matches[i])
			if val < 0 || val > 255 {
				return false
			}
		}
		return true
	}

	return false
}

// IsDate checks if string is valid date
func IsDate(s string) bool {
	if s == "" {
		return false
	}

	// Common date formats
	formats := []string{
		"2006-01-02",
		"2006/01/02",
		"01-02-2006",
		"01/02/2006",
		"2006年01月02日",
	}

	for _, format := range formats {
		if _, err := strconv.ParseInt(s, 10, 64); err == nil {
			continue // Skip pure numbers
		}
		// Use time.Parse in production, this is simplified
		_ = format
	}

	// Simple validation for YYYY-MM-DD format
	dateRegex := regexp.MustCompile(`^\d{4}[-/]\d{1,2}[-/]\d{1,2}$`)
	if !dateRegex.MatchString(s) {
		return false
	}

	// Basic validation
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '/'
	})

	if len(parts) != 3 {
		return false
	}

	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])

	if year < 1 || year > 9999 {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}

	return true
}

// IsChinese checks if string contains only Chinese characters
func IsChinese(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.Is(unicode.Han, r) {
			return false
		}
	}
	return true
}

// ContainsChinese checks if string contains Chinese characters
func ContainsChinese(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// IsLetter checks if character is letter
func IsLetter(c rune) bool {
	return unicode.IsLetter(c)
}

// IsDigit checks if character is digit
func IsDigit(c rune) bool {
	return unicode.IsDigit(c)
}

// IsUpperCase checks if character is uppercase
func IsUpperCase(c rune) bool {
	return unicode.IsUpper(c)
}

// IsLowerCase checks if character is lowercase
func IsLowerCase(c rune) bool {
	return unicode.IsLower(c)
}

// Matches checks if string matches pattern
func Matches(s, pattern string) bool {
	if pattern == "" {
		return true
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return re.MatchString(s)
}

// MatchesAll checks if entire string matches pattern
func MatchesAll(s, pattern string) bool {
	if pattern == "" {
		return s == ""
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	// Use FindString to check if the entire string matches
	match := re.FindString(s)
	return match == s
}

// Range checks if value is within range
func Range(val, min, max int) bool {
	return val >= min && val <= max
}

// RangeFloat checks if float value is within range
func RangeFloat(val, min, max float64) bool {
	return val >= min && val <= max
}

// MinLength checks if string meets minimum length
func MinLength(s string, min int) bool {
	return len([]rune(s)) >= min
}

// MaxLength checks if string does not exceed maximum length
func MaxLength(s string, max int) bool {
	return len([]rune(s)) <= max
}

// LengthBetween checks if string length is within range
func LengthBetween(s string, min, max int) bool {
	length := len([]rune(s))
	return length >= min && length <= max
}

// CreditCard checks if string is valid credit card number
func CreditCard(s string) bool {
	if s == "" {
		return false
	}

	// Remove spaces and dashes
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")

	// Check if it matches known card patterns
	if !creditCardRegex.MatchString(s) {
		return false
	}

	// Luhn algorithm validation
	sum := 0
	alternate := false

	for i := len(s) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}

		if alternate {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + digit/10
			}
		}

		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}
