package bcrypt

import (
	"strings"
	"testing"

	xbcrypt "golang.org/x/crypto/bcrypt"
)

func TestHashAndCheck(t *testing.T) {
	h, err := Hash("s3cret")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(h, "$2a$10$") {
		t.Errorf("hash prefix = %q, want $2a$10$", h[:7])
	}
	if !Check("s3cret", h) {
		t.Error("Check with correct password = false")
	}
	if Check("wrong", h) {
		t.Error("Check with wrong password = true")
	}
}

func TestHashWithCost(t *testing.T) {
	h, err := HashWithCost("p@ss", 4)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(h, "$2a$04$") {
		t.Errorf("cost-4 hash prefix = %q", h[:7])
	}
	if !Check("p@ss", h) {
		t.Error("cost-4 hash should still verify")
	}
	// 72-byte password limit is enforced by x/crypto
	if _, err := Hash(strings.Repeat("a", 73)); err == nil {
		t.Error("73-byte password should error")
	}
}

func TestCostBounds(t *testing.T) {
	// costs below MinCost are clamped to MinCost by x/crypto, not errors
	for _, cost := range []int{0, xbcrypt.MinCost - 1} {
		if h, err := HashWithCost("pw", cost); err != nil || !Check("pw", h) {
			t.Errorf("cost %d should clamp and succeed, err=%v", cost, err)
		}
	}
	// costs above MaxCost (31) make the key schedule overflow and error
	if _, err := HashWithCost("pw", xbcrypt.MaxCost+1); err == nil {
		t.Error("cost above MaxCost should error")
	}
}

func TestCheckBadHash(t *testing.T) {
	if Check("x", "not-a-bcrypt-hash") {
		t.Error("Check on malformed hash = true")
	}
}