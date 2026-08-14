package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// DigestType represents hash digest type
type DigestType int

const (
	DigestMD5 DigestType = iota
	DigestSHA1
	DigestSHA256
	DigestSHA384
	DigestSHA512
)

// GenerateKey generates random key of specified length
func GenerateKey(length int) ([]byte, error) {
	if length <= 0 {
		return nil, errors.New("key length must be positive")
	}
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// GenerateKeyHex generates random key and returns hex string
func GenerateKeyHex(length int) (string, error) {
	key, err := GenerateKey(length)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
