package jwt

import (
	"strings"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "user-123"
	claims.Issuer = "gohutools"
	claims.ExpiresAt = time.Now().Add(time.Hour).Unix()

	key := []byte("secret-key")
	token, err := Sign(claims, key)
	if err != nil {
		t.Fatalf("Sign error: %v", err)
	}

	// token must have 3 dot-separated segments
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Fatalf("token has %d segments, want 3: %q", len(parts), token)
	}
	if parts[0] == "" || parts[1] == "" || parts[2] == "" {
		t.Fatal("empty segment in token")
	}
}

func TestSignAndVerify(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "u1"
	claims.ExpiresAt = time.Now().Add(time.Hour).Unix()

	key := []byte("my-secret")
	token, err := Sign(claims, key)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := Parse(token, key)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	if parsed.Subject != "u1" {
		t.Errorf("Subject = %q, want u1", parsed.Subject)
	}
}

func TestVerifyWrongKey(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "u1"
	token, _ := Sign(claims, []byte("right-key"))

	if _, err := Parse(token, []byte("wrong-key")); err == nil {
		t.Error("Parse with wrong key should error")
	}
}

func TestVerifyTampered(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "u1"
	key := []byte("sec")
	token, _ := Sign(claims, key)

	// tamper payload: replace second segment with base64 of different claim
	parts := strings.Split(token, ".")
	parts[1] = "eyJzdWIiOiJhdGhhY2tlciJ9" // {"sub":"athacker"}
	tampered := strings.Join(parts, ".")

	if _, err := Parse(tampered, key); err == nil {
		t.Error("tampered token should fail signature verification")
	}
}

func TestExpired(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "u1"
	claims.ExpiresAt = time.Now().Add(-time.Hour).Unix() // expired

	token, _ := Sign(claims, []byte("k"))
	_, err := Parse(token, []byte("k"))
	if err == nil {
		t.Error("expired token should error")
	}
	if !strings.Contains(err.Error(), "expired") {
		t.Errorf("expected 'expired' error, got: %v", err)
	}
}

func TestNotYetValid(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "u1"
	claims.NotBefore = time.Now().Add(time.Hour).Unix() // future

	token, _ := Sign(claims, []byte("k"))
	if _, err := Parse(token, []byte("k")); err == nil {
		t.Error("token with future nbf should error")
	}
}

func TestParseInvalidFormat(t *testing.T) {
	if _, err := Parse("not-a-jwt", []byte("k")); err == nil {
		t.Error("Parse(garbage) should error")
	}
	if _, err := Parse("a.b", []byte("k")); err == nil {
		t.Error("Parse(2 segments) should error")
	}
	if _, err := Parse("", []byte("k")); err == nil {
		t.Error("Parse(empty) should error")
	}
}

func TestClaimsAllStandardFields(t *testing.T) {
	now := time.Now()
	claims := NewClaims()
	claims.Subject = "subj"
	claims.Issuer = "issuer"
	claims.Audience = "aud1"
	claims.ExpiresAt = now.Add(10 * time.Minute).Unix()
	claims.NotBefore = now.Add(-time.Minute).Unix()
	claims.IssuedAt = now.Unix()
	claims.JWTID = "jti-1"

	key := []byte("k")
	token, err := Sign(claims, key)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := Parse(token, key)
	if err != nil {
		t.Fatal(err)
	}
	if parsed.Subject != "subj" || parsed.Issuer != "issuer" ||
		parsed.Audience != "aud1" || parsed.JWTID != "jti-1" {
		t.Errorf("claims round trip mismatch: %+v", parsed)
	}
	if parsed.ExpiresAt != claims.ExpiresAt || parsed.NotBefore != claims.NotBefore ||
		parsed.IssuedAt != claims.IssuedAt {
		t.Errorf("time claims mismatch: %+v", parsed)
	}
}

func TestCustomClaims(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "member"
	claims.SetCustom("role", "admin")
	claims.SetCustom("dept", "engineering")

	token, _ := Sign(claims, []byte("k"))
	parsed, _ := Parse(token, []byte("k"))
	if v, ok := parsed.GetCustom("role"); !ok || v != "admin" {
		t.Errorf("custom role = %v (present %v), want admin", v, ok)
	}
	if v, ok := parsed.GetCustom("dept"); !ok || v != "engineering" {
		t.Errorf("custom dept = %v, want engineering", v)
	}
	if _, ok := parsed.GetCustom("nope"); ok {
		t.Error("GetCustom(nonexistent) should report missing")
	}
}

func TestHS384(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "s384"
	token, err := SignHS(claims, []byte("k"), HS384)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := Parse(token, []byte("k"))
	if err != nil {
		t.Fatalf("HS384 Parse error: %v", err)
	}
	if parsed.Subject != "s384" {
		t.Errorf("Subject = %q", parsed.Subject)
	}
}

func TestHS512(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "s512"
	token, err := SignHS(claims, []byte("k"), HS512)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := Parse(token, []byte("k"))
	if err != nil {
		t.Fatalf("HS512 Parse error: %v", err)
	}
	if parsed.Subject != "s512" {
		t.Errorf("Subject = %q", parsed.Subject)
	}
}

func TestParseWrongAlgHeader(t *testing.T) {
	claims := NewClaims()
	claims.Subject = "u"
	token, _ := Sign(claims, []byte("k"))
	// tamper header to say "none" — shouldn't get through verification
	parts := strings.Split(token, ".")
	parts[0] = encodeSegment(`{"alg":"none"}`)
	tampered := strings.Join(parts, ".")
	if _, err := Parse(tampered, []byte("k")); err == nil {
		t.Error("token with alg:none should fail")
	}
}

func TestEmptyKey(t *testing.T) {
	claims := NewClaims()
	if _, err := Sign(claims, []byte("")); err == nil {
		t.Error("Sign with empty key should error")
	}
	if _, err := Parse("a.b.c", []byte("")); err == nil {
		t.Error("Parse with empty key should error")
	}
}