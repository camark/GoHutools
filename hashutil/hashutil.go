package hashutil

// This package ports Hutool's HashUtil collection of classic string hash
// functions. Each function returns a fixed-size hash of a string.

const (
	fnv32Offset  = 2166136261 // 0x811c9dc5
	fnv32Prime   = 16777619   // 0x01000193
	fnv64Offset  = 14695981039346656037
	fnv64Prime   = 1099511628211
)

// BKDRHash computes the BKDR (Brian Kernighan & Dennis Ritchie) hash.
// seed defaults to 131, the classic BKDR multiplier.
func BKDRHash(str string) int64 {
	return BKDRHashSeed(str, 131)
}

// BKDRHashSeed computes the BKDR hash with an explicit seed.
func BKDRHashSeed(str string, seed int64) int64 {
	var hash int64
	for _, c := range str {
		hash = hash*seed + int64(c)
	}
	return hash
}

// SDBMHash computes the SDBM hash (a simple multiply-accumulate hash).
func SDBMHash(str string) int64 {
	var hash int64
	for _, c := range str {
		hash = int64(c) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// DJBHash computes Daniel J. Bernstein's djb2 hash.
// The empty string maps to the 5381 seed.
func DJBHash(str string) int64 {
	hash := int64(5381)
	for _, c := range str {
		hash = (hash << 5) + hash + int64(c) // hash*33 + c
	}
	return hash
}

// DEKHash computes the DEK hash used in the "Dragon Book" (Algorithms in C).
func DEKHash(str string) int64 {
	var hash int64
	for _, c := range str {
		hash = (hash<<5 ^ hash>>27) ^ int64(c)
	}
	return hash
}

// FNVHash computes the 32-bit FNV-1 hash.
func FNVHash(str string) uint32 {
	hash := uint32(fnv32Offset)
	for _, b := range []byte(str) {
		hash *= uint32(fnv32Prime)
		hash ^= uint32(b)
	}
	return hash
}

// FNVHash1A computes the 32-bit FNV-1a hash.
func FNVHash1A(str string) uint32 {
	hash := uint32(fnv32Offset)
	for _, b := range []byte(str) {
		hash ^= uint32(b)
		hash *= uint32(fnv32Prime)
	}
	return hash
}

// FNVHash64 computes the 64-bit FNV-1 hash.
func FNVHash64(str string) uint64 {
	hash := uint64(fnv64Offset)
	for _, b := range []byte(str) {
		hash *= uint64(fnv64Prime)
		hash ^= uint64(b)
	}
	return hash
}

// FNVHash64A computes the 64-bit FNV-1a hash.
func FNVHash64A(str string) uint64 {
	hash := uint64(fnv64Offset)
	for _, b := range []byte(str) {
		hash ^= uint64(b)
		hash *= uint64(fnv64Prime)
	}
	return hash
}

// JavaHash mimics java.lang.String.hashCode (s[0]*31^(n-1) + ... + s[n-1]).
func JavaHash(str string) int32 {
	var h int32
	for _, c := range str {
		h = h*31 + int32(c)
	}
	return h
}

// ELFHash computes the ELF hash used by the System V ELF object format.
func ELFHash(str string) int64 {
	var hash, x int64
	for _, c := range str {
		hash = (hash << 4) + int64(c)
		if x = hash & 0xF0000000; x != 0 {
			hash ^= x >> 24 // clear upper 4 bits and fold
			hash &= ^x
		}
	}
	return hash
}

// PJWHash computes Peter J. Weinberger's hash (closely related to ELFHash).
func PJWHash(str string) int64 { return ELFHash(str) }

// RSHash computes Robert Sedgwicks' hash.
func RSHash(str string) int64 {
	var a, b int64
	for _, c := range str {
		a = (a<<4 + int64(c)) % 127 // large prime, less than 256
		b = (b*5 + int64(c)) % 2933 // prime < 4096
	}
	return a<<16 | b
}

// JSHash computes Justin Sobel's hash.
func JSHash(str string) int64 {
	var hash int64 = 1315423911
	for _, c := range str {
		hash ^= (hash<<5 + int64(c) + hash>>2)
	}
	return hash
}

// APHash computes Arash Partow's hash.
func APHash(str string) int64 {
	var hash int64
	bytes := []byte(str)
	for i, b := range bytes {
		if i&1 == 0 {
			hash ^= (hash<<7 ^ int64(b) ^ hash>>3)
		} else {
			hash ^= ^(hash<<11 ^ int64(b) ^ hash>>5)
		}
	}
	return hash
}

// OneByOneHash processes the string one byte at a time (O'Reilly hash).
func OneByOneHash(str string) int64 {
	hash := int64(0)
	for _, b := range []byte(str) {
		hash += int64(b)
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return hash
}

// AdditiveHash is the classic additive hash (sum of bytes * weighting).
func AdditiveHash(str string) int64 {
	var hash int64
	for i, b := range []byte(str) {
		hash += (int64(i) + 1) * int64(b)
	}
	return hash
}

// MixHash mixes the input into a pseudo-random 32-bit value
// (a simple avalanche-style mixer).
func MixHash(str string) int64 {
	var h int64 = 0x9e3779b9
	for _, c := range []byte(str) {
		h ^= int64(c)
		h *= int64(0x85ebca6b)
		h ^= h >> 13
	}
	h ^= h >> 16
	return h
}