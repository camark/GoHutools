package bloom

import (
	"hash/fnv"
	"math"
	"sync"
)

// Filter is bloom filter
type Filter struct {
	bits      []bool
	size      uint64
	hashCount uint
	count     uint64
	mu        sync.RWMutex
}

// New creates new bloom filter with expected number of items and false positive rate
func New(expectedItems uint64, fpRate float64) *Filter {
	size := optimalSize(expectedItems, fpRate)
	hashCount := optimalHashCount(size, expectedItems)
	return NewWithSize(size, hashCount)
}

// NewWithSize creates bloom filter with specific size and hash count
func NewWithSize(size uint64, hashCount uint) *Filter {
	return &Filter{
		bits:      make([]bool, size),
		size:      size,
		hashCount: hashCount,
		count:     0,
	}
}

// hashWithSeed creates a hash value with a given seed
func hashWithSeed(data []byte, seed uint64) uint64 {
	h := fnv.New64a()
	h.Write(data)
	// Combine with seed using a simple mixing function
	hashVal := h.Sum64()
	return hashVal ^ (seed * 0x9e3779b97f4a7c15)
}

// getHashIndices returns all hash indices for the given data
func (f *Filter) getHashIndices(data []byte) []uint64 {
	indices := make([]uint64, f.hashCount)
	for i := uint(0); i < f.hashCount; i++ {
		indices[i] = hashWithSeed(data, uint64(i)) % f.size
	}
	return indices
}

// Add adds item to filter
func (f *Filter) Add(data []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()

	indices := f.getHashIndices(data)
	for _, idx := range indices {
		f.bits[idx] = true
	}
	f.count++
}

// AddString adds string to filter
func (f *Filter) AddString(s string) {
	f.Add([]byte(s))
}

// Contains checks if item might be in filter
func (f *Filter) Contains(data []byte) bool {
	f.mu.RLock()
	defer f.mu.RUnlock()

	indices := f.getHashIndices(data)
	for _, idx := range indices {
		if !f.bits[idx] {
			return false
		}
	}
	return true
}

// ContainsString checks if string might be in filter
func (f *Filter) ContainsString(s string) bool {
	return f.Contains([]byte(s))
}

// Count returns number of items added
func (f *Filter) Count() uint64 {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return f.count
}

// Reset clears the filter
func (f *Filter) Reset() {
	f.mu.Lock()
	defer f.mu.Unlock()

	for i := range f.bits {
		f.bits[i] = false
	}
	f.count = 0
}

// FalsePositiveRate returns current estimated false positive rate
func (f *Filter) FalsePositiveRate() float64 {
	f.mu.RLock()
	defer f.mu.RUnlock()

	if f.count == 0 {
		return 0
	}

	// Calculate the probability of a bit being set
	// P(bit is set) = 1 - (1 - 1/m)^(k*n)
	// where m is size, k is hash count, n is number of items
	m := float64(f.size)
	k := float64(f.hashCount)
	n := float64(f.count)

	// Probability that a specific bit is still 0
	pZero := math.Pow(1-1/m, k*n)
	// Probability that a specific bit is 1
	pOne := 1 - pZero

	// False positive probability: all k bits are set
	return math.Pow(pOne, k)
}

// optimalSize calculates optimal bit array size
// m = -(n * ln(p)) / (ln(2)^2)
func optimalSize(n uint64, p float64) uint64 {
	if n == 0 {
		return 1
	}
	if p <= 0 {
		p = 0.01
	}
	if p >= 1 {
		p = 0.99
	}

	m := -float64(n) * math.Log(p) / (math.Ln2 * math.Ln2)
	if m < 1 {
		return 1
	}
	return uint64(m)
}

// optimalHashCount calculates optimal number of hash functions
// k = (m/n) * ln(2)
func optimalHashCount(size uint64, n uint64) uint {
	if n == 0 {
		return 1
	}

	k := (float64(size) / float64(n)) * math.Ln2
	if k < 1 {
		return 1
	}
	return uint(math.Round(k))
}
