package codec

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
	data := []byte("Hello, World!")
	expected := "SGVsbG8sIFdvcmxkIQ=="
	result := Base64Encode(data)
	if result != expected {
		t.Errorf("Base64Encode failed: expected %s, got %s", expected, result)
	}
}

func TestBase64Decode(t *testing.T) {
	s := "SGVsbG8sIFdvcmxkIQ=="
	expected := []byte("Hello, World!")
	result, err := Base64Decode(s)
	if err != nil {
		t.Fatalf("Base64Decode failed: %v", err)
	}
	if string(result) != string(expected) {
		t.Errorf("Base64Decode failed: expected %s, got %s", expected, result)
	}
}

func TestBase64URLEncode(t *testing.T) {
	data := []byte("Hello, World! +/")
	expected := "SGVsbG8sIFdvcmxkISArLw=="
	result := Base64URLEncode(data)
	if result != expected {
		t.Errorf("Base64URLEncode failed: expected %s, got %s", expected, result)
	}
}

func TestBase64URLDecode(t *testing.T) {
	s := "SGVsbG8sIFdvcmxkISArLw=="
	expected := []byte("Hello, World! +/")
	result, err := Base64URLDecode(s)
	if err != nil {
		t.Fatalf("Base64URLDecode failed: %v", err)
	}
	if string(result) != string(expected) {
		t.Errorf("Base64URLDecode failed: expected %s, got %s", expected, result)
	}
}

func TestBase64StdEncode(t *testing.T) {
	data := []byte("Hello, World!")
	expected := "SGVsbG8sIFdvcmxkIQ=="
	result := Base64StdEncode(data)
	if result != expected {
		t.Errorf("Base64StdEncode failed: expected %s, got %s", expected, result)
	}
}

func TestBase64StdDecode(t *testing.T) {
	s := "SGVsbG8sIFdvcmxkIQ=="
	expected := []byte("Hello, World!")
	result, err := Base64StdDecode(s)
	if err != nil {
		t.Fatalf("Base64StdDecode failed: %v", err)
	}
	if string(result) != string(expected) {
		t.Errorf("Base64StdDecode failed: expected %s, got %s", expected, result)
	}
}

func TestHexEncode(t *testing.T) {
	data := []byte("Hello, World!")
	expected := "48656c6c6f2c20576f726c6421"
	result := HexEncode(data)
	if result != expected {
		t.Errorf("HexEncode failed: expected %s, got %s", expected, result)
	}
}

func TestHexDecode(t *testing.T) {
	s := "48656c6c6f2c20576f726c6421"
	expected := []byte("Hello, World!")
	result, err := HexDecode(s)
	if err != nil {
		t.Fatalf("HexDecode failed: %v", err)
	}
	if string(result) != string(expected) {
		t.Errorf("HexDecode failed: expected %s, got %s", expected, result)
	}
}

func TestURLEncode(t *testing.T) {
	s := "Hello World! @#$%"
	result := URLEncode(s)
	if result == "" {
		t.Error("URLEncode returned empty string")
	}
}

func TestURLDecode(t *testing.T) {
	s := "Hello+World%21+%40%23%24%25"
	result, err := URLDecode(s)
	if err != nil {
		t.Fatalf("URLDecode failed: %v", err)
	}
	if result != "Hello World! @#$%" {
		t.Errorf("URLDecode failed: expected 'Hello World! @#$%%', got '%s'", result)
	}
}

func TestURLQueryEncode(t *testing.T) {
	s := "key=value&foo=bar"
	result := URLQueryEncode(s)
	if result == "" {
		t.Error("URLQueryEncode returned empty string")
	}
}

func TestURLQueryDecode(t *testing.T) {
	s := "key%3Dvalue%26foo%3Dbar"
	result, err := URLQueryDecode(s)
	if err != nil {
		t.Fatalf("URLQueryDecode failed: %v", err)
	}
	if result != "key=value&foo=bar" {
		t.Errorf("URLQueryDecode failed: expected 'key=value&foo=bar', got '%s'", result)
	}
}

func TestHTMLEscape(t *testing.T) {
	s := `<script>alert("XSS")</script>`
	expected := "&lt;script&gt;alert(&#34;XSS&#34;)&lt;/script&gt;"
	result := HTMLEscape(s)
	if result != expected {
		t.Errorf("HTMLEscape failed: expected %s, got %s", expected, result)
	}
}

func TestHTMLUnescape(t *testing.T) {
	s := "&lt;script&gt;alert(&quot;XSS&quot;)&lt;/script&gt;"
	expected := `<script>alert("XSS")</script>`
	result := HTMLUnescape(s)
	if result != expected {
		t.Errorf("HTMLUnescape failed: expected %s, got %s", expected, result)
	}
}

func TestUnicodeEncode(t *testing.T) {
	s := "Hello 世界"
	result := UnicodeEncode(s)
	if result == "" {
		t.Error("UnicodeEncode returned empty string")
	}
	// Check that non-ASCII characters are encoded
	if result == s {
		t.Error("UnicodeEncode should have encoded non-ASCII characters")
	}
}

func TestUnicodeDecode(t *testing.T) {
	s := "Hello \\u4e16\\u754c"
	expected := "Hello 世界"
	result, err := UnicodeDecode(s)
	if err != nil {
		t.Fatalf("UnicodeDecode failed: %v", err)
	}
	if result != expected {
		t.Errorf("UnicodeDecode failed: expected %s, got %s", expected, result)
	}
}

func TestPunycodeEncode(t *testing.T) {
	s := "münchen.de"
	result, err := PunycodeEncode(s)
	if err != nil {
		t.Fatalf("PunycodeEncode failed: %v", err)
	}
	if result == "" {
		t.Error("PunycodeEncode returned empty string")
	}
}

func TestPunycodeDecode(t *testing.T) {
	s := "xn--mnchen-3ya.de"
	expected := "münchen.de"
	result, err := PunycodeDecode(s)
	if err != nil {
		t.Fatalf("PunycodeDecode failed: %v", err)
	}
	if result != expected {
		t.Errorf("PunycodeDecode failed: expected %s, got %s", expected, result)
	}
}
