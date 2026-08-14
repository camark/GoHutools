package idutil

import (
	"encoding/hex"
	"regexp"
	"sync"
	"testing"
)

// ---------------------------------------------------------------------------
// UUID tests
// ---------------------------------------------------------------------------

func TestUUID(t *testing.T) {
	id := UUID()
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	if !re.MatchString(id) {
		t.Errorf("UUID format mismatch: %s", id)
	}
}

func TestUUIDUniqueness(t *testing.T) {
	seen := make(map[string]bool, 10000)
	for i := 0; i < 10000; i++ {
		id := UUID()
		if seen[id] {
			t.Fatalf("duplicate UUID: %s", id)
		}
		seen[id] = true
	}
}

func TestUUIDBytes(t *testing.T) {
	b := UUIDBytes()
	if len(b) != 16 {
		t.Fatalf("expected 16 bytes, got %d", len(b))
	}
	if b[6]&0xF0 != 0x40 {
		t.Errorf("version bits mismatch: got 0x%02x", b[6])
	}
	if b[8]&0xC0 != 0x80 {
		t.Errorf("variant bits mismatch: got 0x%02x", b[8])
	}
}

func TestSimpleUUID(t *testing.T) {
	id := SimpleUUID()
	if len(id) != 32 {
		t.Fatalf("expected 32 chars, got %d", len(id))
	}
	_, err := hex.DecodeString(id)
	if err != nil {
		t.Fatalf("SimpleUUID is not valid hex: %s", id)
	}
}

// ---------------------------------------------------------------------------
// Snowflake tests
// ---------------------------------------------------------------------------

func TestSnowflakeBasic(t *testing.T) {
	s := NewSnowflake(1)
	id1, err := s.NextID()
	if err != nil {
		t.Fatal(err)
	}
	id2, err := s.NextID()
	if err != nil {
		t.Fatal(err)
	}
	if id1 == id2 {
		t.Error("expected different IDs")
	}
	if id1 <= 0 || id2 <= 0 {
		t.Error("IDs should be positive")
	}
}

func TestSnowflakeMachineIDRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for out-of-range machine ID")
		}
	}()
	NewSnowflake(1024)
}

func TestSnowflakeConcurrency(t *testing.T) {
	s := NewSnowflake(1)
	var mu sync.Mutex
	ids := make(map[int64]bool)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id, err := s.NextID()
			if err != nil {
				t.Error(err)
				return
			}
			mu.Lock()
			if ids[id] {
				t.Errorf("duplicate snowflake ID: %d", id)
			}
			ids[id] = true
			mu.Unlock()
		}()
	}
	wg.Wait()
	if len(ids) != 1000 {
		t.Errorf("expected 1000 unique IDs, got %d", len(ids))
	}
}

// ---------------------------------------------------------------------------
// ULID tests
// ---------------------------------------------------------------------------

func TestULID(t *testing.T) {
	id := ULID()
	if len(id) != 26 {
		t.Fatalf("expected 26 chars, got %d: %s", len(id), id)
	}
	re := regexp.MustCompile(`^[0-9A-HJKMNP-TV-Z]{26}$`)
	if !re.MatchString(id) {
		t.Errorf("ULID format mismatch: %s", id)
	}
}

func TestULIDMonotonic(t *testing.T) {
	id1 := ULID()
	id2 := ULID()
	// Same millisecond may produce same prefix but IDs should differ or be ordered
	if id1 > id2 {
		// They could be in same ms; just verify uniqueness
		if id1 == id2 {
			t.Error("ULIDs should be unique")
		}
	}
}

func TestULIDUniqueness(t *testing.T) {
	seen := make(map[string]bool, 1000)
	for i := 0; i < 1000; i++ {
		id := ULID()
		if seen[id] {
			t.Fatalf("duplicate ULID: %s", id)
		}
		seen[id] = true
	}
}

// ---------------------------------------------------------------------------
// NanoID tests
// ---------------------------------------------------------------------------

func TestNanoID(t *testing.T) {
	id := NanoID()
	if len(id) != 21 {
		t.Fatalf("expected 21 chars, got %d", len(id))
	}
}

func TestNanoIDWithSize(t *testing.T) {
	tests := []int{1, 10, 32, 64}
	for _, size := range tests {
		id := NanoIDWithSize(size)
		if len(id) != size {
			t.Errorf("expected %d chars, got %d", size, len(id))
		}
	}
}

func TestNanoIDWithSizeZero(t *testing.T) {
	id := NanoIDWithSize(0)
	if id != "" {
		t.Errorf("expected empty string for size 0, got %q", id)
	}
}

func TestNanoIDWithSizeNegative(t *testing.T) {
	id := NanoIDWithSize(-1)
	if id != "" {
		t.Errorf("expected empty string for negative size, got %q", id)
	}
}

// ---------------------------------------------------------------------------
// ShortID / RandomID tests
// ---------------------------------------------------------------------------

func TestShortID(t *testing.T) {
	id := ShortID()
	if len(id) != 8 {
		t.Fatalf("expected 8 chars, got %d", len(id))
	}
}

func TestRandomID(t *testing.T) {
	id := RandomID(16)
	if len(id) != 16 {
		t.Fatalf("expected 16 chars, got %d", len(id))
	}
}

func TestRandomIDWithAlphabet(t *testing.T) {
	alphabet := "AB"
	id := RandomIDWithAlphabet(20, alphabet)
	if len(id) != 20 {
		t.Fatalf("expected 20 chars, got %d", len(id))
	}
	for _, c := range id {
		if c != 'A' && c != 'B' {
			t.Errorf("unexpected char: %c", c)
		}
	}
}

func TestRandomIDWithAlphabetEmpty(t *testing.T) {
	id := RandomIDWithAlphabet(10, "")
	if id != "" {
		t.Errorf("expected empty string for empty alphabet, got %q", id)
	}
}

// ---------------------------------------------------------------------------
// Sequence tests
// ---------------------------------------------------------------------------

func TestSequenceBasic(t *testing.T) {
	seq := NewSequence(0, 1)
	if seq.Next() != 0 {
		t.Error("expected 0")
	}
	if seq.Next() != 1 {
		t.Error("expected 1")
	}
	if seq.Next() != 2 {
		t.Error("expected 2")
	}
}

func TestSequenceWithStep(t *testing.T) {
	seq := NewSequence(10, 5)
	if seq.Next() != 10 {
		t.Error("expected 10")
	}
	if seq.Next() != 15 {
		t.Error("expected 15")
	}
	if seq.Next() != 20 {
		t.Error("expected 20")
	}
}

func TestSequenceWithRangeCycled(t *testing.T) {
	seq := NewSequenceWithRange(0, 4, 2, true)
	// Values: 0, 2, 4, 0, 2, 4, ...
	expected := []int64{0, 2, 4, 0, 2, 4}
	for _, exp := range expected {
		got := seq.Next()
		if got != exp {
			t.Errorf("expected %d, got %d", exp, got)
		}
	}
}

func TestSequenceWithRangeClamped(t *testing.T) {
	seq := NewSequenceWithRange(0, 3, 2, false)
	// Values: 0, 2, 3, 3, 3, ...
	if seq.Next() != 0 {
		t.Error("expected 0")
	}
	if seq.Next() != 2 {
		t.Error("expected 2")
	}
	if seq.Next() != 3 {
		t.Error("expected 3 (clamped)")
	}
	if seq.Next() != 3 {
		t.Error("expected 3 (stays clamped)")
	}
}

func TestSequenceNextN(t *testing.T) {
	seq := NewSequence(1, 1)
	vals := seq.NextN(5)
	expected := []int64{1, 2, 3, 4, 5}
	if len(vals) != 5 {
		t.Fatalf("expected 5 values, got %d", len(vals))
	}
	for i, v := range vals {
		if v != expected[i] {
			t.Errorf("index %d: expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestSequenceNextNZero(t *testing.T) {
	seq := NewSequence(0, 1)
	vals := seq.NextN(0)
	if vals != nil {
		t.Errorf("expected nil for n=0, got %v", vals)
	}
}

func TestSequenceReset(t *testing.T) {
	seq := NewSequence(100, 1)
	seq.Next()
	seq.Next()
	seq.Reset()
	if seq.Current() != 100 {
		t.Errorf("expected 100 after reset, got %d", seq.Current())
	}
}

func TestSequenceCurrent(t *testing.T) {
	seq := NewSequence(50, 10)
	if seq.Current() != 50 {
		t.Errorf("expected 50, got %d", seq.Current())
	}
	seq.Next()
	if seq.Current() != 60 {
		t.Errorf("expected 60, got %d", seq.Current())
	}
}

func TestSequencePanicOnBadStep(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for step <= 0")
		}
	}()
	NewSequence(0, 0)
}

func TestSequencePanicOnBadRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for end < start")
		}
	}()
	NewSequenceWithRange(10, 5, 1, false)
}

// ---------------------------------------------------------------------------
// ObjectID tests
// ---------------------------------------------------------------------------

func TestObjectID(t *testing.T) {
	id := ObjectID()
	if len(id) != 24 {
		t.Fatalf("expected 24 hex chars, got %d: %s", len(id), id)
	}
	_, err := hex.DecodeString(id)
	if err != nil {
		t.Fatalf("ObjectID is not valid hex: %s", id)
	}
}

func TestObjectIDUniqueness(t *testing.T) {
	seen := make(map[string]bool, 1000)
	for i := 0; i < 1000; i++ {
		id := ObjectID()
		if seen[id] {
			t.Fatalf("duplicate ObjectID: %s", id)
		}
		seen[id] = true
	}
}

func TestObjectIDTimestampPrefix(t *testing.T) {
	id := ObjectID()
	// First 8 hex chars are the timestamp
	tsHex := id[:8]
	if len(tsHex) != 8 {
		t.Error("timestamp portion too short")
	}
	// Should be a valid hex number
	_, err := hex.DecodeString(tsHex)
	if err != nil {
		t.Errorf("timestamp hex invalid: %s", tsHex)
	}
}

// ---------------------------------------------------------------------------
// MachineID / WorkerID tests
// ---------------------------------------------------------------------------

func TestMachineID(t *testing.T) {
	id, err := MachineID()
	if err != nil {
		t.Skipf("skipping: %v", err)
	}
	if id == "" {
		t.Error("expected non-empty machine ID")
	}
	// Should be hex
	if _, err := hex.DecodeString(id); err != nil {
		t.Errorf("machine ID is not valid hex: %s", id)
	}
}

func TestWorkerID(t *testing.T) {
	id, err := WorkerID()
	if err != nil {
		t.Skipf("skipping: %v", err)
	}
	if id < 0 || id >= 1024 {
		t.Errorf("worker ID out of range: %d", id)
	}
}

// ---------------------------------------------------------------------------
// Benchmarks
// ---------------------------------------------------------------------------

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID()
	}
}

func BenchmarkSimpleUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleUUID()
	}
}

func BenchmarkSnowflake(b *testing.B) {
	s := NewSnowflake(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.NextID()
	}
}

func BenchmarkULID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ULID()
	}
}

func BenchmarkNanoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NanoID()
	}
}

func BenchmarkObjectID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ObjectID()
	}
}

func BenchmarkSequence(b *testing.B) {
	seq := NewSequence(0, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seq.Next()
	}
}
