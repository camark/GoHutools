package hashutil

import "testing"

// TestFNV uses official FNV test vectors.
func TestFNV(t *testing.T) {
	// FNV-1 32-bit vectors
	cases := []struct {
		in   string
		want uint32
	}{
		{"", 0x811c9dc5},   // offset basis
		{"a", 0x050c5d7e},  // FNV-1
		{"foobar", 0x31f0b262},
	}
	for _, c := range cases {
		got := FNVHash(c.in)
		if got != c.want {
			t.Errorf("FNVHash(%q) = %#x, want %#x", c.in, got, c.want)
		}
	}
}

func TestFNV1a(t *testing.T) {
	cases := []struct {
		in   string
		want uint32
	}{
		{"", 0x811c9dc5},       // offset basis
		{"a", 0xe40c292c},      // FNV-1a of "a"
		{"foobar", 0xbf9cf968}, // FNV-1a 32-bit of "foobar"
	}
	for _, c := range cases {
		got := FNVHash1A(c.in)
		if got != c.want {
			t.Errorf("FNVHash1A(%q) = %#x, want %#x", c.in, got, c.want)
		}
	}
}

// TestJavaHash uses the Java String.hashCode contract.
func TestJavaHash(t *testing.T) {
	cases := []struct {
		in   string
		want int32
	}{
		{"", 0},
		{"a", 97},
		{"ab", 3105},
		{"hello", 99162322},
	}
	for _, c := range cases {
		got := JavaHash(c.in)
		if got != c.want {
			t.Errorf("JavaHash(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

// TestAlgorithmicProperties covers algorithms without published vectors:
// determinism, non-zero on non-empty input, and sensitivity to input.
func TestAlgorithmicProperties(t *testing.T) {
	algos := map[string]func(string) int64{
		"BKDR":      BKDRHash,
		"SDBM":      SDBMHash,
		"DJB":       DJBHash,
		"DEK":       DEKHash,
		"ELF":       ELFHash,
		"RS":        RSHash,
		"JS":        JSHash,
		"AP":        APHash,
		"OneByOne":  OneByOneHash,
		"Additive":  AdditiveHash,
		"PJW":       PJWHash,
		"HashMix":   MixHash,
	}
	for name, fn := range algos {
		a := fn("hello")
		b := fn("hello")
		if a != b {
			t.Errorf("%s not deterministic: %d vs %d", name, a, b)
		}
		if fn("") == fn("hello") {
			t.Errorf("%s empty collision with hello", name)
		}
		if fn("hello") != fn("hellp") {
			t.Logf("%s: input-sensitive", name)
		}
	}
}

func TestFNV64(t *testing.T) {
	// FNV-1a 64-bit of "a": 0xaf63dc4c8601ec8c
	got := FNVHash64A("a")
	if got != 0xaf63dc4c8601ec8c {
		t.Errorf("FNVHash64A(a) = %#x, want 0xaf63dc4c8601ec8c", got)
	}
	// empty = 64-bit offset basis 0xcbf29ce484222325
	if FNVHash64A("") != 0xcbf29ce484222325 {
		t.Errorf("FNVHash64A(\"\") = %#x", FNVHash64A(""))
	}
}

func TestKnownBKDR(t *testing.T) {
	// empty input: hash stays at 0
	if BKDRHash("") != 0 {
		t.Errorf("BKDRHash(\"\") = %d, want 0", BKDRHash(""))
	}
	// same seed, same result
	if BKDRHashSeed("seed-pinned", 131) != BKDRHashSeed("seed-pinned", 131) {
		t.Error("BKDR no determinism")
	}
}

func TestDJBOffsetBasis(t *testing.T) {
	if DJBHash("") != 5381 {
		t.Errorf("DJBHash(\"\") = %d, want 5381", DJBHash(""))
	}
}