package crypto

import (
	"os"
	"path/filepath"
	"testing"
)

// "123456789" is the standard CRC check value for CRC-32 IEEE,
// CRC-32C (Castagnoli) and CRC-64 ECMA.

func TestCRC32(t *testing.T) {
	got := CRC32([]byte("123456789"))
	if got != 0xCBF43926 {
		t.Errorf("CRC32 = %#x, want 0xCBF43926", got)
	}
	if h := CRC32Hex([]byte("123456789")); h != "cbf43926" {
		t.Errorf("CRC32Hex = %q", h)
	}
	// empty
	if CRC32(nil) != 0 {
		t.Error("CRC32(empty) should be 0")
	}
}

func TestCRC32C(t *testing.T) {
	if got := CRC32C([]byte("123456789")); got != 0xE3069283 {
		t.Errorf("CRC32C = %#x, want 0xE3069283", got)
	}
}

func TestCRC64(t *testing.T) {
	got := CRC64([]byte("123456789"))
	// Go's crc64.ECMA table is the reflected implementation; its check
	// value for "123456789" is 0x995DC9BBDF1939FA (same as CRC-64/XZ).
	if got != 0x995DC9BBDF1939FA {
		t.Errorf("CRC64 = %#x, want 0x995DC9BBDF1939FA", got)
	}
	if h := CRC64Hex([]byte("123456789")); h != "995dc9bbdf1939fa" {
		t.Errorf("CRC64Hex = %q", h)
	}
}

func TestFNV(t *testing.T) {
	if got := FNV32(nil); got != 2166136261 {
		t.Errorf("FNV32(empty) = %d", got)
	}
	if got := FNV64(nil); got != 14695981039346656037 {
		t.Errorf("FNV64(empty) = %d", got)
	}
}

func TestSHAFileVariants(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.bin")
	content := []byte("The quick brown fox jumps over the lazy dog")
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatal(err)
	}

	// SHA256File sanity: matches hashing the bytes directly
	want256 := SHA256Hex(content)
	got256, err := SHA256File(path)
	if err != nil {
		t.Fatal(err)
	}
	if got256 != want256 {
		t.Errorf("SHA256File = %s, want %s", got256, want256)
	}

	// SHA1File matches SHA1Hex
	if got, err := SHA1File(path); err != nil || got != SHA1Hex(content) {
		t.Errorf("SHA1File = %s/%v, want %s", got, err, SHA1Hex(content))
	}
	// SHA384File / SHA512File
	if got, err := SHA384File(path); err != nil || got != SHA384Hex(content) {
		t.Errorf("SHA384File = %s/%v", got, err)
	}
	if got, err := SHA512File(path); err != nil || got != SHA512Hex(content) {
		t.Errorf("SHA512File = %s/%v", got, err)
	}

	// missing file surfaces error
	if _, err := SHA1File(filepath.Join(dir, "nope")); err == nil {
		t.Error("missing file should error")
	}
}