// Package charsetutil provides charset utility functions similar to Java Hutool's CharsetUtil.
package charsetutil

import (
	"bytes"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// Common charsets
const (
	UTF8    = "UTF-8"
	GBK     = "GBK"
	GB2312  = "GB2312"
	GB18030 = "GB18030"
	BIG5    = "Big5"
	ISO8859 = "ISO-8859-1"
	USASCII = "US-ASCII"
	UTF16   = "UTF-16"
	UTF16LE = "UTF-16LE"
	UTF16BE = "UTF-16BE"
)

// Convert converts bytes from one charset to another
func Convert(data []byte, fromCharset, toCharset string) ([]byte, error) {
	if fromCharset == toCharset {
		return data, nil
	}

	srcEnc, err := getEncoding(fromCharset)
	if err != nil {
		return nil, err
	}

	dstEnc, err := getEncoding(toCharset)
	if err != nil {
		return nil, err
	}

	// First decode from source charset to UTF-8
	decoder := srcEnc.NewDecoder()
	utf8Data, _, err := transform.Bytes(decoder, data)
	if err != nil {
		return nil, err
	}

	// Then encode from UTF-8 to target charset
	encoder := dstEnc.NewEncoder()
	result, _, err := transform.Bytes(encoder, utf8Data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ConvertString converts string from one charset to another
func ConvertString(s string, fromCharset, toCharset string) (string, error) {
	data, err := Convert([]byte(s), fromCharset, toCharset)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ToUTF8 converts bytes to UTF-8
func ToUTF8(data []byte, fromCharset string) ([]byte, error) {
	return Convert(data, fromCharset, UTF8)
}

// FromUTF8 converts UTF-8 bytes to target charset
func FromUTF8(data []byte, toCharset string) ([]byte, error) {
	return Convert(data, UTF8, toCharset)
}

// ToUTF8String converts string to UTF-8
func ToUTF8String(s string, fromCharset string) (string, error) {
	return ConvertString(s, fromCharset, UTF8)
}

// FromUTF8String converts UTF-8 string to target charset
func FromUTF8String(s string, toCharset string) (string, error) {
	return ConvertString(s, UTF8, toCharset)
}

// GBKToUTF8 converts GBK to UTF-8
func GBKToUTF8(data []byte) ([]byte, error) {
	return ToUTF8(data, GBK)
}

// UTF8ToGBK converts UTF-8 to GBK
func UTF8ToGBK(data []byte) ([]byte, error) {
	return FromUTF8(data, GBK)
}

// GBKToUTF8String converts GBK string to UTF-8
func GBKToUTF8String(s string) (string, error) {
	return ToUTF8String(s, GBK)
}

// UTF8ToGBKString converts UTF-8 string to GBK
func UTF8ToGBKString(s string) (string, error) {
	return FromUTF8String(s, GBK)
}

// IsUTF8 checks if bytes is valid UTF-8
func IsUTF8(data []byte) bool {
	return utf8.Valid(data)
}

// IsGBK checks if bytes is valid GBK
func IsGBK(data []byte) bool {
	decoder := simplifiedchinese.GBK.NewDecoder()
	_, err := decoder.Bytes(data)
	return err == nil
}

// DetectCharset tries to detect charset of bytes
func DetectCharset(data []byte) string {
	if IsUTF8(data) {
		return UTF8
	}
	if IsGBK(data) {
		return GBK
	}
	return ISO8859
}

// CleanUTF8 removes invalid UTF-8 sequences
func CleanUTF8(data []byte) []byte {
	return bytes.ToValidUTF8(data, []byte{0xEF, 0xBF, 0xBD})
}

// CleanUTF8String removes invalid UTF-8 sequences from string
func CleanUTF8String(s string) string {
	return strings.ToValidUTF8(s, "�")
}

// RuneCount returns number of runes in string
func RuneCount(s string) int {
	return utf8.RuneCountInString(s)
}

// RuneLen returns byte length of rune
func RuneLen(r rune) int {
	return utf8.RuneLen(r)
}

// EncodeRune encodes rune to bytes
func EncodeRune(r rune) []byte {
	buf := make([]byte, utf8.RuneLen(r))
	utf8.EncodeRune(buf, r)
	return buf
}

// DecodeRune decodes rune from bytes
func DecodeRune(p []byte) (rune, int) {
	return utf8.DecodeRune(p)
}

// ValidRune checks if rune is valid
func ValidRune(r rune) bool {
	return utf8.ValidRune(r)
}

// UTF16ToUTF8 converts UTF-16 to UTF-8
func UTF16ToUTF8(data []byte, littleEndian bool) ([]byte, error) {
	charset := UTF16BE
	if littleEndian {
		charset = UTF16LE
	}
	return Convert(data, charset, UTF8)
}

// UTF8ToUTF16 converts UTF-8 to UTF-16
func UTF8ToUTF16(data []byte, littleEndian bool) ([]byte, error) {
	charset := UTF16BE
	if littleEndian {
		charset = UTF16LE
	}
	return Convert(data, UTF8, charset)
}

// getEncoding returns encoding for charset
func getEncoding(charset string) (encoding.Encoding, error) {
	switch strings.ToUpper(charset) {
	case UTF8, "UTF8":
		return unicode.UTF8, nil
	case GBK:
		return simplifiedchinese.GBK, nil
	case GB2312:
		return simplifiedchinese.HZGB2312, nil
	case GB18030:
		return simplifiedchinese.GB18030, nil
	case BIG5, "BIG5":
		return traditionalchinese.Big5, nil
	case UTF16, UTF16LE:
		return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM), nil
	case UTF16BE:
		return unicode.UTF16(unicode.BigEndian, unicode.UseBOM), nil
	default:
		return nil, ErrUnsupportedCharset(charset)
	}
}

// ErrUnsupportedCharset returns error for unsupported charset
type ErrUnsupportedCharset string

func (e ErrUnsupportedCharset) Error() string {
	return "unsupported charset: " + string(e)
}

// Charsets returns list of supported charsets
func Charsets() []string {
	return []string{
		UTF8, GBK, GB2312, GB18030, BIG5, ISO8859, USASCII,
		UTF16, UTF16LE, UTF16BE,
	}
}

// IsSupported checks if charset is supported
func IsSupported(charset string) bool {
	upper := strings.ToUpper(charset)
	for _, c := range Charsets() {
		if strings.ToUpper(c) == upper {
			return true
		}
	}
	return false
}
