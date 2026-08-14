package randomutil

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"time"
)

// ---------------------------------------------------------------------------
// Integer random
// ---------------------------------------------------------------------------

// Int generates a cryptographically secure random int in [0, max).
// Returns an error if max <= 0 or if reading from crypto/rand fails.
func Int(max int) (int, error) {
	if max <= 0 {
		return 0, nil
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

// IntRange generates a cryptographically secure random int in [min, max).
// Returns an error if min >= max or if reading from crypto/rand fails.
func IntRange(min, max int) (int, error) {
	if min >= max {
		return 0, nil
	}
	n, err := Int(max - min)
	if err != nil {
		return 0, err
	}
	return min + n, nil
}

// MustInt generates a random int in [0, max). Panics on error.
func MustInt(max int) int {
	n, err := Int(max)
	if err != nil {
		panic(err)
	}
	return n
}

// MustIntRange generates a random int in [min, max). Panics on error.
func MustIntRange(min, max int) int {
	n, err := IntRange(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

// Int64 generates a cryptographically secure random int64 in [0, max).
func Int64(max int64) (int64, error) {
	if max <= 0 {
		return 0, nil
	}
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

// Int64Range generates a cryptographically secure random int64 in [min, max).
func Int64Range(min, max int64) (int64, error) {
	if min >= max {
		return 0, nil
	}
	n, err := Int64(max - min)
	if err != nil {
		return 0, err
	}
	return min + n, nil
}

// ---------------------------------------------------------------------------
// Float random
// ---------------------------------------------------------------------------

// Float64 generates a cryptographically secure random float64 in [0.0, 1.0).
func Float64() (float64, error) {
	// Use 53 bits of randomness for full float64 precision
	n, err := rand.Int(rand.Reader, big.NewInt(1<<53))
	if err != nil {
		return 0, err
	}
	return float64(n.Int64()) / float64(1<<53), nil
}

// Float64Range generates a cryptographically secure random float64 in [min, max).
func Float64Range(min, max float64) (float64, error) {
	if min >= max {
		return min, nil
	}
	f, err := Float64()
	if err != nil {
		return 0, err
	}
	return min + f*(max-min), nil
}

// ---------------------------------------------------------------------------
// Boolean random
// ---------------------------------------------------------------------------

// Bool generates a cryptographically secure random bool.
func Bool() (bool, error) {
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		return false, err
	}
	return b[0]&1 == 1, nil
}

// MustBool generates a random bool. Panics on error.
func MustBool() bool {
	b, err := Bool()
	if err != nil {
		panic(err)
	}
	return b
}

// ---------------------------------------------------------------------------
// String random
// ---------------------------------------------------------------------------

const (
	alphaLower   = "abcdefghijklmnopqrstuvwxyz"
	alphaUpper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaAll     = alphaLower + alphaUpper
	alphanumeric = alphaAll + "0123456789"
	hexChars     = "0123456789abcdef"
)

// String generates a random alphanumeric string of the specified length using crypto/rand.
func String(length int) string {
	return StringWithAlphabet(length, alphanumeric)
}

// StringWithAlphabet generates a random string with a custom alphabet using crypto/rand.
func StringWithAlphabet(length int, alphabet string) string {
	if length <= 0 || len(alphabet) == 0 {
		return ""
	}
	b := make([]byte, length)
	alphLen := big.NewInt(int64(len(alphabet)))
	for i := range b {
		n, _ := rand.Int(rand.Reader, alphLen)
		b[i] = alphabet[n.Int64()]
	}
	return string(b)
}

// Alpha generates a random alphabetic string (a-zA-Z).
func Alpha(length int) string {
	return StringWithAlphabet(length, alphaAll)
}

// AlphaNumeric generates a random alphanumeric string (a-zA-Z0-9).
func AlphaNumeric(length int) string {
	return StringWithAlphabet(length, alphanumeric)
}

// Numeric generates a random numeric string (0-9).
func Numeric(length int) string {
	return StringWithAlphabet(length, "0123456789")
}

// Hex generates a random lowercase hex string.
func Hex(length int) string {
	return StringWithAlphabet(length, hexChars)
}

// ---------------------------------------------------------------------------
// Bytes
// ---------------------------------------------------------------------------

// Bytes generates cryptographically secure random bytes of the specified length.
func Bytes(length int) ([]byte, error) {
	if length <= 0 {
		return nil, nil
	}
	b := make([]byte, length)
	_, err := rand.Read(b)
	return b, err
}

// MustBytes generates random bytes. Panics on error.
func MustBytes(length int) []byte {
	b, err := Bytes(length)
	if err != nil {
		panic(err)
	}
	return b
}

// ---------------------------------------------------------------------------
// Slice helpers
// ---------------------------------------------------------------------------

// Element picks a random element from the slice.
// Returns the element and true if the slice is non-empty, or the zero value and false otherwise.
func Element[T any](slice []T) (T, bool) {
	var zero T
	if len(slice) == 0 {
		return zero, false
	}
	idx, err := Int(len(slice))
	if err != nil {
		return zero, false
	}
	return slice[idx], true
}

// MustElement picks a random element from the slice. Panics if the slice is empty.
func MustElement[T any](slice []T) T {
	if len(slice) == 0 {
		panic("randomutil: empty slice")
	}
	idx, err := Int(len(slice))
	if err != nil {
		panic(err)
	}
	return slice[idx]
}

// Elements picks n random elements from the slice (with replacement).
// If n <= 0, returns an empty slice.
func Elements[T any](slice []T, n int) []T {
	if len(slice) == 0 || n <= 0 {
		return nil
	}
	result := make([]T, n)
	for i := 0; i < n; i++ {
		idx, _ := Int(len(slice))
		result[i] = slice[idx]
	}
	return result
}

// Shuffle returns a shuffled copy of the slice using Fisher-Yates with crypto/rand.
func Shuffle[T any](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}
	result := make([]T, len(slice))
	copy(result, slice)
	for i := len(result) - 1; i > 0; i-- {
		j, _ := Int(i + 1)
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// ---------------------------------------------------------------------------
// Weighted random
// ---------------------------------------------------------------------------

// Weighted picks a random index based on weights. Higher weight means higher probability.
// Returns an error if the weights slice is empty or the total weight is 0.
func Weighted(weights []int) (int, error) {
	if len(weights) == 0 {
		return 0, nil
	}
	total := 0
	for _, w := range weights {
		total += w
	}
	if total <= 0 {
		return 0, nil
	}
	r, err := Int(total)
	if err != nil {
		return 0, err
	}
	cumulative := 0
	for i, w := range weights {
		cumulative += w
		if r < cumulative {
			return i, nil
		}
	}
	return len(weights) - 1, nil
}

// ---------------------------------------------------------------------------
// Dice / Coin
// ---------------------------------------------------------------------------

// Dice rolls a single die, returning a value from 1 to 6.
func Dice() int {
	n, _ := Int(6)
	return n + 1
}

// DiceN rolls n dice, returning a slice of values each from 1 to 6.
func DiceN(n int) []int {
	if n <= 0 {
		return nil
	}
	result := make([]int, n)
	for i := range result {
		result[i] = Dice()
	}
	return result
}

// CoinFlip flips a coin, returning true (heads) or false (tails).
func CoinFlip() bool {
	return MustBool()
}

// ---------------------------------------------------------------------------
// Color
// ---------------------------------------------------------------------------

// Color generates a random hex color string (e.g. "#a3f1c2").
func Color() string {
	b := make([]byte, 3)
	_, _ = rand.Read(b)
	dst := make([]byte, 7)
	dst[0] = '#'
	hex.Encode(dst[1:], b)
	return string(dst)
}

// RGB generates a random RGB color with each component in [0, 255].
func RGB() (r, g, b int) {
	buf := make([]byte, 3)
	_, _ = rand.Read(buf)
	return int(buf[0]), int(buf[1]), int(buf[2])
}

// ---------------------------------------------------------------------------
// Date / Time
// ---------------------------------------------------------------------------

// DateBetween generates a random time between start and end (inclusive of start, exclusive of end).
func DateBetween(start, end time.Time) time.Time {
	if start.After(end) {
		start, end = end, start
	}
	diff := end.Sub(start)
	if diff <= 0 {
		return start
	}
	n, _ := Int64(int64(diff))
	return start.Add(time.Duration(n))
}

// ---------------------------------------------------------------------------
// Chinese mock data
// ---------------------------------------------------------------------------

var (
	chineseSurnames = []string{
		"赵", "钱", "孙", "李", "周", "吴", "郑", "王",
		"冯", "陈", "诸", "卫", "沈", "韩", "杨", "朱",
		"秦", "尤", "许", "何", "吕", "施", "张", "孔",
		"曹", "严", "华", "金", "魏", "陶", "姜", "戚",
		"谢", "邹", "喻", "柏", "水", "窦", "章", "云",
		"苏", "潘", "蓑", "契", "范", "彭", "郎", "鲁",
		"韦", "昌", "马", "苗", "凤", "花", "方", "俞",
		"任", "袁", "柳", "酏", "鲍", "史", "唐", "费",
	}
	chineseNamesMale = []string{
		"伟", "强", "磊", "军", "勇", "杰", "浩", "明",
		"志", "刚", "峰", "超", "平", "东", "文", "辉",
		"博", "斜", "晚", "宇", "轩", "翰", "晨", "晖",
		"阳", "驮", "骄", "勳", "风", "海", "浩", "炅",
	}
	chineseNamesFemale = []string{
		"芳", "娚", "翠", "婉", "悠", "婉", "细", "玉",
		"苞", "菊", "兰", "莳", "静", "淑", "现", "淑",
		"梅", "洗", "绵", "美", "雨", "婉", "丽", "颖",
	}
	chineseCities = []string{
		"北京市", "上海市", "广州市", "深圳市",
		"成都市", "杭州市", "武汉市", "南京市",
		"重庆市", "西安市", "苏州市", "天津市",
		"长沙市", "郑州市", "青岛市", "大连市",
	}
	chineseRoads = []string{
		"解放路", "人民路", "建设路", "和平路",
		"中山路", "长江路", "黄河路", "南京路",
		"北京路", "科技路", "工业路", "光明路",
	}
	domains = []string{
		"example.com", "test.com", "demo.org", "sample.net",
		"mail.com", "dev.io", "app.co", "work.cn",
	}
	phoneNumberPrefixes = []string{
		"130", "131", "132", "133", "134", "135", "136", "137", "138", "139",
		"150", "151", "152", "153", "155", "156", "157", "158", "159",
		"170", "176", "177", "178",
		"180", "181", "182", "183", "184", "185", "186", "187", "188", "189",
		"191", "193", "195", "196", "197", "198", "199",
	}
)

// PhoneNumber generates a random Chinese mobile phone number (11 digits).
func PhoneNumber() string {
	prefix, _ := Element(phoneNumberPrefixes)
	suffix := Numeric(8)
	return prefix + suffix
}

// Email generates a random email address.
func Email() string {
	name := AlphaNumeric(8)
	domain, _ := Element(domains)
	return name + "@" + domain
}

// Name generates a random Chinese name (2-3 characters).
func Name() string {
	surname, _ := Element(chineseSurnames)
	if MustBool() {
		// Two-character given name
		if MustBool() {
			n1, _ := Element(chineseNamesMale)
			n2, _ := Element(chineseNamesMale)
			return surname + n1 + n2
		}
		n1, _ := Element(chineseNamesFemale)
		n2, _ := Element(chineseNamesFemale)
		return surname + n1 + n2
	}
	// Single-character given name
	if MustBool() {
		n, _ := Element(chineseNamesMale)
		return surname + n
	}
	n, _ := Element(chineseNamesFemale)
	return surname + n
}

// Address generates a random Chinese address.
func Address() string {
	city, _ := Element(chineseCities)
	road, _ := Element(chineseRoads)
	number := MustIntRange(1, 999)
	return city + road + MustIntToString(number) + "号"
}

// MustIntToString is a helper for address generation.
func MustIntToString(n int) string {
	return intToString(n)
}

func intToString(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return s
}
