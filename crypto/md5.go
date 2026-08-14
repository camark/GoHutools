package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// MD5 computes MD5 hash
func MD5(data []byte) []byte {
	h := md5.Sum(data)
	return h[:]
}

// MD5Hex computes MD5 hash as hex string
func MD5Hex(data []byte) string {
	return hex.EncodeToString(MD5(data))
}

// MD5String computes MD5 hash of string
func MD5String(s string) string {
	return hex.EncodeToString(MD5([]byte(s)))
}

// MD5StringHex computes MD5 hash of string as hex
func MD5StringHex(s string) string {
	return hex.EncodeToString(MD5([]byte(s)))
}

// MD5File computes MD5 hash of file
func MD5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
