// Package bcrypt provides password hashing helpers on top of
// golang.org/x/crypto/bcrypt, mirroring hutool's BCrypt
// (hashpw/checkpw) API.
package bcrypt

import (
	xbcrypt "golang.org/x/crypto/bcrypt"
)

// Hash generates a BCrypt password hash with the default cost (10).
// Passwords longer than 72 bytes are rejected.
func Hash(password string) (string, error) {
	hash, err := xbcrypt.GenerateFromPassword([]byte(password), xbcrypt.DefaultCost)
	return string(hash), err
}

// HashWithCost generates a BCrypt hash with an explicit cost.
// cost must be within [MinCost, MaxCost] = [4, 31].
func HashWithCost(password string, cost int) (string, error) {
	hash, err := xbcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hash), err
}

// Check reports whether password matches a BCrypt hash (hutool
// BCrypt.checkpw semantics). Invalid hashes simply return false.
func Check(password, hashed string) bool {
	err := xbcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}