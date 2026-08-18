package urlutil

import (
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	// space + special chars (form encoding, space -> +)
	got := Encode("hello world & 你好")
	if strings.Contains(got, " ") {
		t.Errorf("Encode should escape spaces: %q", got)
	}
	if got != "hello+world+%26+%E4%BD%A0%E5%A5%BD" {
		t.Errorf("Encode = %q", got)
	}
}

func TestDecode(t *testing.T) {
	got, err := Decode("hello%20world%26x")
	if err != nil {
		t.Fatal(err)
	}
	if got != "hello world&x" {
		t.Errorf("Decode = %q", got)
	}
	// invalid input
	if _, err := Decode("%zz"); err == nil {
		t.Error("Decode(invalid) should error")
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct{ in, want string }{
		{"example.com", "https://example.com"},
		{"https://example.com", "https://example.com"},
		{"http://example.com:80", "http://example.com"},
		{"https://example.com:443", "https://example.com"},
		{"http://example.com/", "http://example.com/"},
		{"example.com/path?a=1", "https://example.com/path?a=1"},
	}
	for _, tt := range tests {
		got, err := Normalize(tt.in)
		if err != nil {
			t.Errorf("Normalize(%q) error: %v", tt.in, err)
			continue
		}
		if got != tt.want {
			t.Errorf("Normalize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestNormalizeDefaultScheme(t *testing.T) {
	got, err := NormalizeWithScheme("example.com", "http")
	if err != nil {
		t.Fatal(err)
	}
	if got != "http://example.com" {
		t.Errorf("NormalizeWithScheme = %q", got)
	}
}

func TestNormalizeInvalid(t *testing.T) {
	if _, err := Normalize("http://exa mple.com"); err == nil {
		t.Error("Normalize(invalid) should error")
	}
}

func TestGetParam(t *testing.T) {
	query := "a=1&b=hello%20world&flag&a=2"
	if got := GetParam(query, "a"); got != "1" {
		t.Errorf("GetParam(a) = %q, want 1", got)
	}
	if got := GetParam(query, "b"); got != "hello world" {
		t.Errorf("GetParam(b) = %q", got)
	}
	if got := GetParam(query, "flag"); got != "" {
		t.Errorf("GetParam(flag) = %q, want empty", got)
	}
	if got := GetParam(query, "missing"); got != "" {
		t.Errorf("GetParam(missing) = %q", got)
	}
}

func TestGetParams(t *testing.T) {
	query := "a=1&a=2&b=x"
	vals := GetParams(query)
	if len(vals["a"]) != 2 || vals["a"][0] != "1" || vals["a"][1] != "2" {
		t.Errorf("GetParams(a) = %v", vals["a"])
	}
	if len(vals["b"]) != 1 || vals["b"][0] != "x" {
		t.Errorf("GetParams(b) = %v", vals["b"])
	}
}

func TestGetPath(t *testing.T) {
	if got := GetPath("https://example.com/a/b/c?x=1"); got != "/a/b/c" {
		t.Errorf("GetPath = %q", got)
	}
	if got := GetPath("https://example.com"); got != "" {
		t.Errorf("GetPath(root) = %q, want empty", got)
	}
	if got := GetPath("not a url"); got != "" {
		t.Errorf("GetPath(invalid) = %q, want empty", got)
	}
}

func TestGetHost(t *testing.T) {
	if got := GetHost("https://api.example.com:8443/v1?q=1"); got != "api.example.com" {
		t.Errorf("GetHost = %q", got)
	}
}

func TestIsHTTP(t *testing.T) {
	if !IsHTTP("http://a.com") {
		t.Error("IsHTTP(http) should be true")
	}
	if IsHTTP("https://a.com") {
		t.Error("IsHTTP(https) should be false")
	}
	if IsHTTP("ftp://a.com") {
		t.Error("IsHTTP(ftp) should be false")
	}
	if IsHTTP("garbage") {
		t.Error("IsHTTP(garbage) should be false")
	}
	if !IsHTTPS("https://a.com") {
		t.Error("IsHTTPS(https) should be true")
	}
	if IsHTTPS("http://a.com") {
		t.Error("IsHTTPS(http) should be false")
	}
}

func TestBuildURL(t *testing.T) {
	u := NewBuilder("https://example.com/api").
		Path("v1/users").
		Query("page", "2").
		Query("q", "hello world").
		Build()

	expected := "https://example.com/api/v1/users?page=2&q=hello+world"
	if u != expected {
		t.Errorf("Build = %q, want %q", u, expected)
	}
}

func TestBuildURLFragment(t *testing.T) {
	u := NewBuilder("https://example.com").
		Fragment("section").
		Build()
	if u != "https://example.com#section" {
		t.Errorf("Build = %q", u)
	}
}

func TestBuildURLLateQueryMerge(t *testing.T) {
	// query present in base URL must be preserved when adding params
	u := NewBuilder("https://example.com?existing=1").Query("new", "2").Build()
	if !strings.Contains(u, "existing=1") || !strings.Contains(u, "new=2") {
		t.Errorf("Build should merge queries: %q", u)
	}
}

func TestGetScheme(t *testing.T) {
	if got := GetScheme("https://x.com"); got != "https" {
		t.Errorf("GetScheme = %q", got)
	}
	if got := GetScheme("x.com"); got != "" {
		t.Errorf("GetScheme(no scheme) = %q", got)
	}
}