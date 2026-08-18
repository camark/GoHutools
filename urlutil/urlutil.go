package urlutil

import (
	"net/url"
	"strings"
)

// Encode percent-encodes s using URL escaped form (allows only unreserved chars).
func Encode(s string) string {
	return url.QueryEscape(s)
}

// Decode percent-decodes s. Returns an error on invalid input.
func Decode(s string) (string, error) {
	return url.QueryUnescape(s)
}

// Normalize canonicalizes a URL:
//   - adds a default "https" scheme when missing
//   - strips default ports (80 for http, 443 for https)
//
// Returns an error when the input cannot be parsed.
func Normalize(s string) (string, error) {
	return NormalizeWithScheme(s, "https")
}

// NormalizeWithScheme canonicalizes a URL using the given default scheme when
// the input lacks one.
func NormalizeWithScheme(s, defaultScheme string) (string, error) {
	trimmed := strings.TrimSpace(s)
	if !strings.Contains(trimmed, "://") {
		trimmed = defaultScheme + "://" + trimmed
	}
	u, err := url.Parse(trimmed)
	if err != nil {
		return "", err
	}
	if (u.Scheme == "http" && u.Port() == "80") || (u.Scheme == "https" && u.Port() == "443") {
		// strip default port
		host := u.Hostname()
		u.Host = host
	}
	return u.String(), nil
}

// GetParam returns the first value of key in the given raw query string.
func GetParam(rawQuery, key string) string {
	vals := GetParams(rawQuery)
	vv := vals[key]
	if len(vv) == 0 {
		return ""
	}
	return vv[0]
}

// GetParams parses a raw query string into a url.Values map (never nil).
func GetParams(rawQuery string) url.Values {
	vals, _ := url.ParseQuery(rawQuery)
	if vals == nil {
		return url.Values{}
	}
	return vals
}

// GetPath returns the path portion of an absolute URL, or "" when invalid.
func GetPath(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}
	// reject relative URLs (no host & no scheme) such as "not a url"
	if u.Host == "" && u.Scheme == "" {
		return ""
	}
	return u.EscapedPath()
}

// GetHost returns the host (without port) of a URL.
func GetHost(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}
	return u.Hostname()
}

// GetScheme returns the scheme ("http", "https", ...) or "" when absent.
func GetScheme(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}
	return u.Scheme
}

// IsHTTP reports whether s starts with the http:// scheme.
func IsHTTP(s string) bool {
	return GetScheme(s) == "http"
}

// IsHTTPS reports whether s starts with the https:// scheme.
func IsHTTPS(s string) bool {
	return GetScheme(s) == "https"
}

// Builder is a chainable URL builder mirroring Hutool's UrlBuilder.
type Builder struct {
	scheme   string
	host     string
	port     string // empty = default
	path     string
	query    url.Values
	fragment string
}

// NewBuilder creates a builder initialized from a base URL string.
func NewBuilder(base string) *Builder {
	b := &Builder{query: url.Values{}}
	u, err := url.Parse(base)
	if err != nil {
		return b
	}
	b.scheme = u.Scheme
	b.host = u.Hostname()
	b.port = u.Port()
	b.path = u.Path
	if u.RawQuery != "" {
		if q, err := url.ParseQuery(u.RawQuery); err == nil {
			b.query = q
		}
	}
	b.fragment = u.Fragment
	return b
}

// Path appends the given segments to the current path, joining with "/".
func (b *Builder) Path(segments ...string) *Builder {
	if len(segments) == 0 {
		return b
	}
	app := strings.Join(segments, "/")
	// normalize: no trailing/leading slashes collision
	base := strings.TrimSuffix(b.path, "/")
	app = "/" + strings.TrimPrefix(app, "/")
	b.path = base + app
	if b.path == "" {
		b.path = "/"
	}
	return b
}

// Query adds a query parameter (appends when the key already exists).
func (b *Builder) Query(key, value string) *Builder {
	b.query.Add(key, value)
	return b
}

// SetQuery sets a query parameter, replacing previous values of key.
func (b *Builder) SetQuery(key, value string) *Builder {
	b.query.Set(key, value)
	return b
}

// Fragment sets the URL fragment.
func (b *Builder) Fragment(f string) *Builder {
	b.fragment = f
	return b
}

// Build renders the final URL string.
func (b *Builder) Build() string {
	var sb strings.Builder
	if b.scheme != "" {
		sb.WriteString(b.scheme)
		sb.WriteString("://")
	}
	sb.WriteString(b.host)
	if b.port != "" {
		sb.WriteString(":")
		sb.WriteString(b.port)
	}
	if b.path != "" {
		sb.WriteString(b.path)
	}
	if len(b.query) > 0 {
		sb.WriteString("?")
		sb.WriteString(b.query.Encode())
	}
	if b.fragment != "" {
		sb.WriteString("#")
		sb.WriteString(b.fragment)
	}
	return sb.String()
}