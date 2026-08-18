package convert

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// ToStr converts any value to string
func ToStr(v interface{}) string {
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	case bool:
		return strconv.FormatBool(val)
	case int:
		return strconv.Itoa(val)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case error:
		return val.Error()
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%v", val)
	}
}

// ToInt converts to int
func ToInt(v interface{}) (int, error) {
	if v == nil {
		return 0, nil
	}

	switch val := v.(type) {
	case int:
		return val, nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint64:
		return int(val), nil
	case float32:
		return int(val), nil
	case float64:
		return int(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.Atoi(strings.TrimSpace(val))
	case []byte:
		return strconv.Atoi(string(val))
	default:
		return 0, fmt.Errorf("cannot convert %v to int", v)
	}
}

// MustToInt converts to int, panics on error
func MustToInt(v interface{}) int {
	result, err := ToInt(v)
	if err != nil {
		panic(err)
	}
	return result
}

// ToInt64 converts to int64
func ToInt64(v interface{}) (int64, error) {
	if v == nil {
		return 0, nil
	}

	switch val := v.(type) {
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return val, nil
	case uint:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint64:
		return int64(val), nil
	case float32:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseInt(strings.TrimSpace(val), 10, 64)
	case []byte:
		return strconv.ParseInt(string(val), 10, 64)
	default:
		return 0, fmt.Errorf("cannot convert %v to int64", v)
	}
}

// ToFloat64 converts to float64
func ToFloat64(v interface{}) (float64, error) {
	if v == nil {
		return 0, nil
	}

	switch val := v.(type) {
	case int:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseFloat(strings.TrimSpace(val), 64)
	case []byte:
		return strconv.ParseFloat(string(val), 64)
	default:
		return 0, fmt.Errorf("cannot convert %v to float64", v)
	}
}

// MustToFloat64 converts to float64, panics on error
func MustToFloat64(v interface{}) float64 {
	result, err := ToFloat64(v)
	if err != nil {
		panic(err)
	}
	return result
}

// ToBool converts to bool
func ToBool(v interface{}) (bool, error) {
	if v == nil {
		return false, nil
	}

	switch val := v.(type) {
	case bool:
		return val, nil
	case int:
		return val != 0, nil
	case int8:
		return val != 0, nil
	case int16:
		return val != 0, nil
	case int32:
		return val != 0, nil
	case int64:
		return val != 0, nil
	case uint:
		return val != 0, nil
	case uint8:
		return val != 0, nil
	case uint16:
		return val != 0, nil
	case uint32:
		return val != 0, nil
	case uint64:
		return val != 0, nil
	case float32:
		return val != 0, nil
	case float64:
		return val != 0, nil
	case string:
		s := strings.ToLower(strings.TrimSpace(val))
		switch s {
		case "true", "1", "yes", "on", "y", "t":
			return true, nil
		case "false", "0", "no", "off", "n", "f", "":
			return false, nil
		default:
			return false, fmt.Errorf("cannot convert %q to bool", val)
		}
	case []byte:
		return ToBool(string(val))
	default:
		return false, fmt.Errorf("cannot convert %v to bool", v)
	}
}

// MustToBool converts to bool, panics on error
func MustToBool(v interface{}) bool {
	result, err := ToBool(v)
	if err != nil {
		panic(err)
	}
	return result
}

// ToBytes converts string to bytes
func ToBytes(s string) []byte {
	return []byte(s)
}

// ToString converts bytes to string
func ToString(b []byte) string {
	return string(b)
}

// ToIntSlice converts interface slice to int slice
func ToIntSlice(arr []interface{}) []int {
	result := make([]int, 0, len(arr))
	for _, v := range arr {
		if n, err := ToInt(v); err == nil {
			result = append(result, n)
		}
	}
	return result
}

// ToStrSlice converts interface slice to string slice
func ToStrSlice(arr []interface{}) []string {
	result := make([]string, 0, len(arr))
	for _, v := range arr {
		result = append(result, ToStr(v))
	}
	return result
}

// ToMap converts to map
func ToMap(v interface{}) (map[string]interface{}, error) {
	if v == nil {
		return nil, errors.New("cannot convert nil to map")
	}

	if m, ok := v.(map[string]interface{}); ok {
		return m, nil
	}

	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CamelToUnderline converts camelCase to under_score
func CamelToUnderline(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteByte('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// UnderlineToCamel converts under_score to camelCase
func UnderlineToCamel(s string) string {
	if s == "" {
		return ""
	}

	parts := strings.Split(s, "_")
	var result strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		for i, r := range part {
			if i == 0 {
				result.WriteRune(unicode.ToUpper(r))
			} else {
				result.WriteRune(r)
			}
		}
	}
	return result.String()
}

// HexToLong converts hex string to int64
func HexToLong(hexStr string) (int64, error) {
	hexStr = strings.TrimSpace(hexStr)
	hexStr = strings.TrimPrefix(hexStr, "0x")
	hexStr = strings.TrimPrefix(hexStr, "0X")

	if hexStr == "" {
		return 0, errors.New("empty hex string")
	}

	return strconv.ParseInt(hexStr, 16, 64)
}

// LongToHex converts int64 to hex string
func LongToHex(n int64) string {
	return strconv.FormatInt(n, 16)
}

// NumberToChinese converts number to Chinese representation
func NumberToChinese(n int64, isUseTraditional bool) string {
	if n == 0 {
		return "零"
	}

	simpleDigits := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	traditionalDigits := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}

	simpleUnits := []string{"", "十", "百", "千"}
	traditionalUnits := []string{"", "拾", "佰", "仟"}

	simpleSections := []string{"", "万", "亿", "兆", "京"}
	traditionalSections := []string{"", "萬", "億", "兆", "京"}

	digits := simpleDigits
	units := simpleUnits
	sections := simpleSections
	if isUseTraditional {
		digits = traditionalDigits
		units = traditionalUnits
		sections = traditionalSections
	}

	if n < 0 {
		// Use the absolute value as an unsigned magnitude so that
		// math.MinInt64 does not overflow when negated.
		return "负" + numberToChineseMagnitude(uint64(-n), digits, units, sections)
	}

	return numberToChineseMagnitude(uint64(n), digits, units, sections)
}

// numberToChineseMagnitude converts a non-negative magnitude to Chinese.
func numberToChineseMagnitude(n uint64, digits, units, sections []string) string {
	s := strconv.FormatUint(n, 10)
	length := len(s)

	// Pad to multiple of 4
	paddedLen := ((length + 3) / 4) * 4
	s = strings.Repeat("0", paddedLen-length) + s

	var result strings.Builder
	sectionCount := paddedLen / 4
	needZero := false // whether a "零" is needed before the next non-zero digit

	for i := 0; i < sectionCount; i++ {
		section := s[i*4 : (i+1)*4]
		sectionNum, _ := strconv.Atoi(section)

		if sectionNum == 0 {
			needZero = true
			continue
		}

		sectionStr := ""
		for j, ch := range section {
			digit := int(ch - '0')
			if digit == 0 {
				needZero = true
			} else {
				if needZero && result.Len() > 0 {
					sectionStr += digits[0]
				}
				needZero = false
				sectionStr += digits[digit] + units[3-j]
			}
		}

		result.WriteString(sectionStr)
		sectionIdx := sectionCount - 1 - i
		if sectionIdx > 0 {
			result.WriteString(sections[sectionIdx])
		}
	}

	return result.String()
}

// ChineseToNumber converts Chinese number to int64
func ChineseToNumber(chinese string) (int64, error) {
	if chinese == "" {
		return 0, errors.New("empty string")
	}

	digitMap := map[rune]int64{
		'零': 0, '一': 1, '二': 2, '三': 3, '四': 4,
		'五': 5, '六': 6, '七': 7, '八': 8, '九': 9,
		'壹': 1, '贰': 2, '叁': 3, '肆': 4,
		'伍': 5, '陆': 6, '柒': 7, '捌': 8, '玖': 9,
	}

	unitMap := map[rune]int64{
		'十': 10, '拾': 10,
		'百': 100, '佰': 100,
		'千': 1000, '仟': 1000,
		'万': 10000, '萬': 10000,
		'亿': 100000000, '億': 100000000,
		'兆': 1000000000000,
		'京': 10000000000000000,
	}

	if chinese == "负" {
		return 0, errors.New("invalid Chinese number")
	}

	negative := false
	if strings.HasPrefix(chinese, "负") {
		negative = true
		chinese = strings.TrimPrefix(chinese, "负")
	}

	var result int64
	var current int64
	var section int64

	for _, r := range chinese {
		if val, ok := digitMap[r]; ok {
			current = val
		} else if unit, ok := unitMap[r]; ok {
			if unit >= 10000 {
				// Large units (万/亿/兆/京): flush the current section
				// into result and start a new section.
				section = (section + current) * unit
				result += section
				section = 0
				current = 0
			} else {
				if current == 0 {
					current = 1
				}
				section += current * unit
				current = 0
			}
		} else {
			return 0, fmt.Errorf("invalid Chinese character: %c", r)
		}
	}

	result += section + current

	if negative {
		result = -result
	}

	return result, nil
}

// DigitToChinese converts digit to Chinese (壹贰叁...)
func DigitToChinese(digit byte) string {
	traditional := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	if digit <= 9 {
		return traditional[digit]
	}
	return ""
}

// RomanToInt converts Roman numeral to int
func RomanToInt(roman string) (int, error) {
	if roman == "" {
		return 0, errors.New("empty Roman numeral")
	}

	romanMap := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	roman = strings.ToUpper(roman)
	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		ch := roman[i]
		value, ok := romanMap[ch]
		if !ok {
			return 0, fmt.Errorf("invalid Roman numeral character: %c", ch)
		}

		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result, nil
}

// IntToRoman converts int to Roman numeral
func IntToRoman(n int) string {
	if n <= 0 || n > 3999 {
		return ""
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	for i := 0; i < len(values); i++ {
		for n >= values[i] {
			result.WriteString(symbols[i])
			n -= values[i]
		}
	}

	return result.String()
}

// BaseConvert converts number between bases
func BaseConvert(value string, fromBase, toBase int) (string, error) {
	if fromBase < 2 || fromBase > 36 {
		return "", errors.New("fromBase must be between 2 and 36")
	}
	if toBase < 2 || toBase > 36 {
		return "", errors.New("toBase must be between 2 and 36")
	}

	// Parse the value in the source base
	num, err := strconv.ParseInt(value, fromBase, 64)
	if err != nil {
		return "", fmt.Errorf("cannot parse %q in base %d: %v", value, fromBase, err)
	}

	// Convert to the target base
	return strconv.FormatInt(num, toBase), nil
}

// ToEnum converts string to enum value
func ToEnum(s string, enumMap map[string]int) (int, bool) {
	val, ok := enumMap[s]
	return val, ok
}

// Base64Encode encodes string to base64
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode decodes base64 string
func Base64Decode(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// HexEncode encodes bytes to hex string
func HexEncode(data []byte) string {
	return hex.EncodeToString(data)
}

// HexDecode decodes hex string to bytes
func HexDecode(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsNumeric checks if value is numeric
func IsNumeric(v interface{}) bool {
	switch v := v.(type) {
	case int, int8, int16, int32, int64:
		return true
	case uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return true
	case string:
		if v == "" {
			return false
		}
		_, err := strconv.ParseFloat(v, 64)
		return err == nil
	default:
		return false
	}
}

// TypeOf returns the type name of value
func TypeOf(v interface{}) string {
	if v == nil {
		return "nil"
	}
	return reflect.TypeOf(v).String()
}

// Round rounds float to specified decimal places
func Round(f float64, decimals int) float64 {
	shift := math.Pow(10, float64(decimals))
	return math.Round(f*shift) / shift
}
