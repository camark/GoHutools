package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Algorithms supported by this package.
const (
	HS256 = "HS256"
	HS384 = "HS384"
	HS512 = "HS512"
)

// Claims holds the standard JWT claims plus arbitrary custom claims.
type Claims struct {
	Subject   string                 `json:"sub,omitempty"`
	Issuer    string                 `json:"iss,omitempty"`
	Audience  string                 `json:"aud,omitempty"`
	ExpiresAt int64                  `json:"exp,omitempty"`
	NotBefore int64                  `json:"nbf,omitempty"`
	IssuedAt  int64                  `json:"iat,omitempty"`
	JWTID     string                 `json:"jti,omitempty"`
	Custom    map[string]interface{} `json:"-"`
}

// NewClaims returns an empty claim set with IssuedAt set to now.
func NewClaims() *Claims {
	return &Claims{IssuedAt: time.Now().Unix()}
}

// SetCustom stores a custom claim value.
func (c *Claims) SetCustom(key string, value interface{}) {
	if c.Custom == nil {
		c.Custom = make(map[string]interface{})
	}
	c.Custom[key] = value
}

// GetCustom retrieves a custom claim value.
func (c *Claims) GetCustom(key string) (interface{}, bool) {
	if c.Custom == nil {
		return nil, false
	}
	v, ok := c.Custom[key]
	return v, ok
}

// allClaims merges standard fields and custom claims into a single map.
func (c *Claims) allClaims() map[string]interface{} {
	m := make(map[string]interface{}, len(c.Custom)+7)
	if c.Subject != "" {
		m["sub"] = c.Subject
	}
	if c.Issuer != "" {
		m["iss"] = c.Issuer
	}
	if c.Audience != "" {
		m["aud"] = c.Audience
	}
	if c.ExpiresAt != 0 {
		m["exp"] = c.ExpiresAt
	}
	if c.NotBefore != 0 {
		m["nbf"] = c.NotBefore
	}
	if c.IssuedAt != 0 {
		m["iat"] = c.IssuedAt
	}
	if c.JWTID != "" {
		m["jti"] = c.JWTID
	}
	for k, v := range c.Custom {
		m[k] = v
	}
	return m
}

// claimsFromMap reconstructs a Claims from a payload map.
func claimsFromMap(m map[string]interface{}) *Claims {
	c := &Claims{}
	if v, ok := m["sub"].(string); ok {
		c.Subject = v
	}
	if v, ok := m["iss"].(string); ok {
		c.Issuer = v
	}
	if v, ok := m["aud"].(string); ok {
		c.Audience = v
	}
	if v, ok := m["jti"].(string); ok {
		c.JWTID = v
	}
	if v, ok := toInt64(m["exp"]); ok {
		c.ExpiresAt = v
	}
	if v, ok := toInt64(m["nbf"]); ok {
		c.NotBefore = v
	}
	if v, ok := toInt64(m["iat"]); ok {
		c.IssuedAt = v
	}
	// remaining keys are custom
	for k, v := range m {
		switch k {
		case "sub", "iss", "aud", "jti", "exp", "nbf", "iat":
		default:
			c.SetCustom(k, v)
		}
	}
	return c
}

func toInt64(v interface{}) (int64, bool) {
	switch n := v.(type) {
	case float64:
		return int64(n), true
	case int64:
		return n, true
	case int:
		return int64(n), true
	}
	return 0, false
}

// header builds the JWT header for a given algorithm.
func header(alg string) map[string]interface{} {
	return map[string]interface{}{"alg": alg, "typ": "JWT"}
}

// encodeSegment base64url-encodes a JSON value.
func encodeSegment(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

func decodeSegment(seg string, dst interface{}) error {
	b, err := base64.RawURLEncoding.DecodeString(seg)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

// sign computes the HMAC signature for header.payload with the given key and algorithm.
func sign(headerJSON, payloadJSON string, key []byte, alg string) ([]byte, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("jwt: signing key must not be empty")
	}
	signingInput := headerJSON + "." + payloadJSON
	switch alg {
	case HS256:
		mac := hmac.New(sha256.New, key)
		mac.Write([]byte(signingInput))
		return mac.Sum(nil), nil
	case HS384:
		mac := hmac.New(sha512.New384, key)
		mac.Write([]byte(signingInput))
		return mac.Sum(nil), nil
	case HS512:
		mac := hmac.New(sha512.New, key)
		mac.Write([]byte(signingInput))
		return mac.Sum(nil), nil
	default:
		return nil, fmt.Errorf("jwt: unsupported algorithm %s", alg)
	}
}

// Signature is a compact building block: a fully-formed signed token.
type Signature string

// String returns the token string.
func (s Signature) String() string { return string(s) }

// Sign creates a signed JWT with HS256 using the given secret key.
func Sign(claims *Claims, key []byte) (string, error) {
	return SignHS(claims, key, HS256)
}

// SignHS creates a signed JWT with the specified HMAC algorithm.
func SignHS(claims *Claims, key []byte, alg string) (string, error) {
	if claims == nil {
		return "", fmt.Errorf("jwt: claims must not be nil")
	}
	h := header(alg)
	hdrJSON, err := json.Marshal(h)
	if err != nil {
		return "", err
	}
	// deterministic segment encoding
	hdr := base64.RawURLEncoding.EncodeToString(hdrJSON)

	payload, err := json.Marshal(claims.allClaims())
	if err != nil {
		return "", err
	}
	pl := base64.RawURLEncoding.EncodeToString(payload)

	sig, err := sign(hdr, pl, key, alg)
	if err != nil {
		return "", err
	}
	return hdr + "." + pl + "." + base64.RawURLEncoding.EncodeToString(sig), nil
}

// Parse verifies a JWT and returns its claims.
// It rejects tokens with an invalid signature, wrong algorithm "none",
// expired tokens (exp in the past) and tokens used before nbf.
func Parse(token string, key []byte) (*Claims, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("jwt: verification key must not be empty")
	}
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("jwt: malformed token: expected 3 segments, got %d", len(parts))
	}
	hdrSeg, payloadSeg, sigSeg := parts[0], parts[1], parts[2]

	var hdr map[string]interface{}
	if err := decodeSegment(hdrSeg, &hdr); err != nil {
		return nil, fmt.Errorf("jwt: invalid header: %w", err)
	}
	alg, _ := hdr["alg"].(string)
	if alg == "none" {
		return nil, fmt.Errorf("jwt: algorithm 'none' is not allowed")
	}
	// header may also carry alg with HS
	switch alg {
	case HS256, HS384, HS512:
	default:
		return nil, fmt.Errorf("jwt: unsupported algorithm %q", alg)
	}

	// verify signature
	sig, err := base64.RawURLEncoding.DecodeString(sigSeg)
	if err != nil {
		return nil, fmt.Errorf("jwt: invalid signature encoding: %w", err)
	}
	expected, err := sign(hdrSeg, payloadSeg, key, alg)
	if err != nil {
		return nil, err
	}
	if !hmac.Equal(sig, expected) {
		return nil, fmt.Errorf("jwt: signature verification failed")
	}

	// parse payload
	var m map[string]interface{}
	if err := decodeSegment(payloadSeg, &m); err != nil {
		return nil, fmt.Errorf("jwt: invalid payload: %w", err)
	}
	claims := claimsFromMap(m)

	now := time.Now().Unix()
	if claims.ExpiresAt != 0 && now >= claims.ExpiresAt {
		return nil, fmt.Errorf("jwt: token expired")
	}
	if claims.NotBefore != 0 && now < claims.NotBefore {
		return nil, fmt.Errorf("jwt: token not yet valid")
	}
	return claims, nil
}