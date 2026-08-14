package bloom

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	filter := New(1000, 0.01)

	if filter == nil {
		t.Fatal("Expected filter to be created")
	}
	if filter.size == 0 {
		t.Error("Expected filter size to be greater than 0")
	}
	if filter.hashCount == 0 {
		t.Error("Expected at least one hash function")
	}
	if filter.count != 0 {
		t.Error("Expected initial count to be 0")
	}
}

func TestNewWithSize(t *testing.T) {
	filter := NewWithSize(1000, 3)

	if filter == nil {
		t.Fatal("Expected filter to be created")
	}
	if filter.size != 1000 {
		t.Errorf("Expected size 1000, got %d", filter.size)
	}
	if filter.hashCount != 3 {
		t.Errorf("Expected 3 hash functions, got %d", filter.hashCount)
	}
}

func TestAddAndContains(t *testing.T) {
	filter := New(100, 0.01)

	// Test Add and Contains with bytes
	filter.Add([]byte("test"))
	if !filter.Contains([]byte("test")) {
		t.Error("Expected filter to contain 'test'")
	}
	if filter.Contains([]byte("other")) {
		t.Error("Expected filter to not contain 'other'")
	}

	// Test AddString and ContainsString
	filter.AddString("hello")
	if !filter.ContainsString("hello") {
		t.Error("Expected filter to contain 'hello'")
	}
	if filter.ContainsString("world") {
		t.Error("Expected filter to not contain 'world'")
	}
}

func TestCount(t *testing.T) {
	filter := New(100, 0.01)

	if filter.Count() != 0 {
		t.Errorf("Expected count 0, got %d", filter.Count())
	}

	filter.Add([]byte("test1"))
	if filter.Count() != 1 {
		t.Errorf("Expected count 1, got %d", filter.Count())
	}

	filter.Add([]byte("test2"))
	if filter.Count() != 2 {
		t.Errorf("Expected count 2, got %d", filter.Count())
	}

	filter.AddString("test3")
	if filter.Count() != 3 {
		t.Errorf("Expected count 3, got %d", filter.Count())
	}
}

func TestReset(t *testing.T) {
	filter := New(100, 0.01)

	filter.Add([]byte("test1"))
	filter.Add([]byte("test2"))
	filter.Add([]byte("test3"))

	if filter.Count() != 3 {
		t.Errorf("Expected count 3, got %d", filter.Count())
	}

	filter.Reset()

	if filter.Count() != 0 {
		t.Errorf("Expected count 0 after reset, got %d", filter.Count())
	}

	// After reset, should not contain previously added items
	if filter.Contains([]byte("test1")) {
		t.Error("Expected filter to not contain 'test1' after reset")
	}
}

func TestFalsePositiveRate(t *testing.T) {
	// Create filter with known parameters
	filter := New(1000, 0.01)

	// Initially, false positive rate should be 0
	if filter.FalsePositiveRate() != 0 {
		t.Errorf("Expected initial false positive rate to be 0, got %f", filter.FalsePositiveRate())
	}

	// Add some items
	for i := 0; i < 100; i++ {
		filter.Add([]byte("item" + string(rune('0'+i%10))))
	}

	// False positive rate should be greater than 0
	fpr := filter.FalsePositiveRate()
	if fpr <= 0 {
		t.Error("Expected false positive rate to be greater than 0")
	}
	if fpr > 1 {
		t.Errorf("Expected false positive rate to be less than 1, got %f", fpr)
	}
}

func TestOptimalSize(t *testing.T) {
	// Test with typical values
	size := optimalSize(1000, 0.01)
	if size == 0 {
		t.Error("Expected optimal size to be greater than 0")
	}

	// Test with zero items
	size = optimalSize(0, 0.01)
	if size != 1 {
		t.Errorf("Expected size 1 for 0 items, got %d", size)
	}

	// Test with invalid probability
	size = optimalSize(1000, 0)
	if size == 0 {
		t.Error("Expected optimal size to be greater than 0 for invalid probability")
	}

	size = optimalSize(1000, 1)
	if size == 0 {
		t.Error("Expected optimal size to be greater than 0 for probability 1")
	}
}

func TestOptimalHashCount(t *testing.T) {
	// Test with typical values
	k := optimalHashCount(10000, 1000)
	if k == 0 {
		t.Error("Expected hash count to be greater than 0")
	}

	// Test with zero items
	k = optimalHashCount(1000, 0)
	if k != 1 {
		t.Errorf("Expected hash count 1 for 0 items, got %d", k)
	}
}

func TestFalsePositives(t *testing.T) {
	// Create a filter with known parameters
	filter := New(1000, 0.01)

	// Add 1000 items
	for i := 0; i < 1000; i++ {
		filter.Add([]byte("item" + string(rune(i%128))))
	}

	// Check for false positives with items not added
	falsePositives := 0
	totalChecks := 10000

	for i := 0; i < totalChecks; i++ {
		if filter.Contains([]byte("notadded" + string(rune(i%128)))) {
			falsePositives++
		}
	}

	// The false positive rate should be close to 1% (0.01)
	// Allow margin for statistical variance
	fpr := float64(falsePositives) / float64(totalChecks)
	if fpr > 0.10 { // Allow 10% margin for statistical variance
		t.Errorf("False positive rate too high: %f (expected around 0.01)", fpr)
	}
}

func TestConcurrency(t *testing.T) {
	filter := New(10000, 0.01)
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				filter.Add([]byte("item" + string(rune(i*100+j))))
			}
		}(i)
	}

	// Concurrent reads
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				filter.Contains([]byte("item" + string(rune(i*100+j))))
			}
		}(i)
	}

	wg.Wait()

	// Verify count
	if filter.Count() != 10000 {
		t.Errorf("Expected count 10000, got %d", filter.Count())
	}
}

func TestLargeFilter(t *testing.T) {
	// Test with large numbers
	filter := New(1000000, 0.001)

	if filter.size == 0 {
		t.Error("Expected filter size to be greater than 0")
	}
	if filter.hashCount == 0 {
		t.Error("Expected at least one hash function")
	}

	// Add many items
	for i := 0; i < 10000; i++ {
		filter.Add([]byte("item" + string(rune(i))))
	}

	if filter.Count() != 10000 {
		t.Errorf("Expected count 10000, got %d", filter.Count())
	}
}

func TestEmptyFilter(t *testing.T) {
	filter := New(0, 0.01)

	if filter.Count() != 0 {
		t.Errorf("Expected count 0, got %d", filter.Count())
	}

	if filter.Contains([]byte("test")) {
		t.Error("Expected empty filter to not contain anything")
	}

	fpr := filter.FalsePositiveRate()
	if fpr != 0 {
		t.Errorf("Expected false positive rate 0, got %f", fpr)
	}
}

func TestSingleItem(t *testing.T) {
	filter := New(1, 0.01)

	filter.Add([]byte("only"))

	if !filter.Contains([]byte("only")) {
		t.Error("Expected filter to contain 'only'")
	}

	if filter.Count() != 1 {
		t.Errorf("Expected count 1, got %d", filter.Count())
	}
}
