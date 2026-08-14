package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// HMACMD5 computes HMAC-MD5
func HMACMD5(key, data []byte) []byte {
	h := hmac.New(md5.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// HMACMD5Hex computes HMAC-MD5 as hex string
func HMACMD5Hex(key, data []byte) string {
	return hex.EncodeToString(HMACMD5(key, data))
}

// HMACSHA1 computes HMAC-SHA1
func HMACSHA1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// HMACSHA1Hex computes HMAC-SHA1 as hex string
func HMACSHA1Hex(key, data []byte) string {
	return hex.EncodeToString(HMACSHA1(key, data))
}

// HMACSHA256 computes HMAC-SHA256
func HMACSHA256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// HMACSHA256Hex computes HMAC-SHA256 as hex string
func HMACSHA256Hex(key, data []byte) string {
	return hex.EncodeToString(HMACSHA256(key, data))
}

// HMACSHA512 computes HMAC-SHA512
func HMACSHA512(key, data []byte) []byte {
	h := hmac.New(sha512.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// HMACSHA512Hex computes HMAC-SHA512 as hex string
func HMACSHA512Hex(key, data []byte) string {
	return hex.EncodeToString(HMACSHA512(key, data))
}
