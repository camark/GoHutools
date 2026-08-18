package crypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
)

// SHA1 computes SHA-1 hash
func SHA1(data []byte) []byte {
	h := sha1.Sum(data)
	return h[:]
}

// SHA1Hex computes SHA-1 hash as hex string
func SHA1Hex(data []byte) string {
	return hex.EncodeToString(SHA1(data))
}

// SHA256 computes SHA-256 hash
func SHA256(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}

// SHA256Hex computes SHA-256 hash as hex string
func SHA256Hex(data []byte) string {
	return hex.EncodeToString(SHA256(data))
}

// SHA384 computes SHA-384 hash
func SHA384(data []byte) []byte {
	h := sha512.Sum384(data)
	return h[:]
}

// SHA384Hex computes SHA-384 hash as hex string
func SHA384Hex(data []byte) string {
	return hex.EncodeToString(SHA384(data))
}

// SHA512 computes SHA-512 hash
func SHA512(data []byte) []byte {
	h := sha512.Sum512(data)
	return h[:]
}

// SHA512Hex computes SHA-512 hash as hex string
func SHA512Hex(data []byte) string {
	return hex.EncodeToString(SHA512(data))
}

// SHA1File computes SHA-1 hash of file
func SHA1File(path string) (string, error) {
	return fileDigestHex(path, sha1.New)
}

// SHA384File computes SHA-384 hash of file
func SHA384File(path string) (string, error) {
	return fileDigestHex(path, sha512.New384)
}

// SHA512File computes SHA-512 hash of file
func SHA512File(path string) (string, error) {
	return fileDigestHex(path, sha512.New)
}

// SHA256File computes SHA-256 hash of file
func SHA256File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
