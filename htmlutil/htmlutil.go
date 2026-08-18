package htmlutil

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

// This package ports Hutool's HtmlUtil and parts of EscapeUtil.

var (
	tagRe      = regexp.MustCompile(`<[^>]*>`)
	attrReFmt  = `\s+%s\s*=\s*("[^"]*"|'[^']*'|[^\s>]+)`
	eventRe    = regexp.MustCompile(`\s+on[a-zA-Z]+\s*=\s*("[^"]*"|'[^']*'|[^\s>]+)`)
	scriptRe   = regexp.MustCompile(`(?is)<(script|style|iframe|object|embed)[^>]*>.*?</(?:script|style|iframe|object|embed)[^>]*>|<\s*(?:script|style|iframe|object|embed)[^>]*>`)
	jsLinkRe   = regexp.MustCompile(`(?i)(href|src)\s*=\s*["']\s*javascript:[^"']*["']`)
	unicodeRe  = regexp.MustCompile(`\\u([0-9a-fA-F]{4})`)
	unwrapRe   = regexp.MustCompile(`(?s)^<([a-zA-Z][a-zA-Z0-9]*)[^>]*>([\s\S]*?)</[a-zA-Z][a-zA-Z0-9]*\s*>$`)
)

func attrRe(attr string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(attrReFmt, attr))
}

// Escape escapes HTML metacharacters: < > & ' " become entities.
func Escape(s string) string {
	return html.EscapeString(s)
}

// Unescape decodes HTML entities (named and numeric) back to characters.
func Unescape(s string) string {
	return html.UnescapeString(s)
}

// CleanHtmlTag strips all HTML tags, keeping the text between them.
func CleanHtmlTag(s string) string {
	return tagRe.ReplaceAllString(s, "")
}

// RemoveHtmlAttr removes the named attribute (case-sensitively) from tags.
func RemoveHtmlAttr(s, attr string) string {
	return attrRe(attr).ReplaceAllString(s, "")
}

// RemoveHtmlAttrI removes the named attribute case-insensitively.
func RemoveHtmlAttrI(s, attr string) string {
	return attrRe(`(?i)` + attr).ReplaceAllString(s, "")
}

// Filter is a pragmatic anti-XSS filter: it removes script-like blocks
// (<script>, <style>, <iframe>, <object>, <embed>), on* event handler
// attributes, and javascript: URLs in href/src.
func Filter(s string) string {
	s = scriptRe.ReplaceAllString(s, "")
	s = eventRe.ReplaceAllString(s, "")
	s = jsLinkRe.ReplaceAllString(s, `$1="#"`)
	return s
}

// ToHex replaces every rune above U+007F with its numeric character
// reference &#xHH; ; ASCII charactrs pass through untouched.
func ToHex(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if r <= 0x7F {
			b.WriteRune(r)
		} else {
			fmt.Fprintf(&b, "&#x%X;", r)
		}
	}
	return b.String()
}

// UnwrapHtml removes a single outermost wrapping tag pair from s,
// leaving the inner content (which may itself contain tags).
func UnwrapHtml(s string) string {
	if m := unwrapRe.FindStringSubmatch(s); m != nil {
		return m[2]
	}
	return s
}

// EncodeUnicode encodes non-ASCII runes as \uXXXX escapes;
// ASCII runes pass through untouched.
func EncodeUnicode(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r <= 0x7F {
			b.WriteRune(r)
		} else {
			fmt.Fprintf(&b, `\u%04x`, r)
		}
	}
	return b.String()
}

// DecodeUnicode decodes \uXXXX escapes (upper or lower case hex) back
// to the characters they represent.
func DecodeUnicode(s string) string {
	return unicodeRe.ReplaceAllStringFunc(s, func(m string) string {
		var r rune
		fmt.Sscanf(m[2:], "%04x", &r)
		return string(r)
	})
}