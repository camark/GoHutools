package crypto

import (
	"encoding/hex"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io"
	"os"
)

// 对标 hutool 的 CrcUtil / 文件摘要（DigestUtil），纯标准库实现。

// CRC32 returns the IEEE CRC-32 checksum of data.
func CRC32(data []byte) uint32 {
	return crc32.Checksum(data, crc32.IEEETable)
}

// CRC32Hex returns the IEEE CRC-32 checksum of data as 8 lowercase hex chars.
func CRC32Hex(data []byte) string {
	return hex.EncodeToString(ue32(crc32.Checksum(data, crc32.IEEETable)))
}

// CRC32C returns the Castagnoli CRC-32C checksum of data.
func CRC32C(data []byte) uint32 {
	return crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli))
}

// CRC64 returns the ECMA-182 CRC-64 checksum of data as computed by
// hash/crc64's ECMA table (reflected implementation; the check value
// for "123456789" is 0x995DC9BBDF1939FA, identical to CRC-64/XZ).
func CRC64(data []byte) uint64 {
	return crc64.Checksum(data, crc64.MakeTable(crc64.ECMA))
}

// CRC64Hex returns the CRC-64 checksum as 16 lowercase hex chars.
func CRC64Hex(data []byte) string {
	return hex.EncodeToString(ue64(crc64.Checksum(data, crc64.MakeTable(crc64.ECMA))))
}

// FNV32 returns the 32-bit FNV-1a hash of data.
func FNV32(data []byte) uint32 {
	h := fnv.New32a()
	_, _ = h.Write(data)
	return h.Sum32()
}

// FNV64 returns the 64-bit FNV-1a hash of data.
func FNV64(data []byte) uint64 {
	h := fnv.New64a()
	_, _ = h.Write(data)
	return h.Sum64()
}

// fileDigestHex computes the given hash over a file and returns its hex form.
func fileDigestHex(path string, newHash func() hash.Hash) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := newHash()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func ue32(v uint32) []byte { return []byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)} }

func ue64(v uint64) []byte {
	return []byte{byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32), byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
}