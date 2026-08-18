package codec

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/idna"
)

// Base64Encode encodes to base64
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode decodes base64
func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// Base64URLEncode encodes to URL-safe base64
func Base64URLEncode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// Base64URLDecode decodes URL-safe base64
func Base64URLDecode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}

// Base64StdEncode encodes to standard base64
func Base64StdEncode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64StdDecode decodes standard base64
func Base64StdDecode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// HexEncode encodes to hex string
func HexEncode(data []byte) string {
	return hex.EncodeToString(data)
}

// HexDecode decodes hex string
func HexDecode(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

// URLEncode encodes URL
func URLEncode(s string) string {
	return url.QueryEscape(s)
}

// URLDecode decodes URL
func URLDecode(s string) (string, error) {
	return url.QueryUnescape(s)
}

// URLQueryEncode encodes URL query
func URLQueryEncode(s string) string {
	return url.QueryEscape(s)
}

// URLQueryDecode decodes URL query
func URLQueryDecode(s string) (string, error) {
	return url.QueryUnescape(s)
}

// HTMLEscape escapes HTML
func HTMLEscape(s string) string {
	return html.EscapeString(s)
}

// HTMLUnescape unescapes HTML
func HTMLUnescape(s string) string {
	return html.UnescapeString(s)
}

// UnicodeEncode encodes to Unicode escape
func UnicodeEncode(s string) string {
	var buf strings.Builder
	for _, r := range s {
		if r > 127 {
			fmt.Fprintf(&buf, "\\u%04x", r)
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

var unicodeEscapeRegex = regexp.MustCompile(`\\u([0-9a-fA-F]{4})`)

// UnicodeDecode decodes Unicode escape
func UnicodeDecode(s string) (string, error) {
	result := unicodeEscapeRegex.ReplaceAllStringFunc(s, func(match string) string {
		hexStr := match[2:]
		var r rune
		if _, err := fmt.Sscanf(hexStr, "%x", &r); err != nil {
			return match
		}
		return string(r)
	})
	return result, nil
}

// PunycodeEncode encodes to Punycode
func PunycodeEncode(s string) (string, error) {
	return idna.ToASCII(s)
}

// PunycodeDecode decodes Punycode
func PunycodeDecode(s string) (string, error) {
	return idna.ToUnicode(s)
}
