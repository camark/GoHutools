package cache

import (
	"sync"
	"testing"
	"time"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRU(3)

	// Test Set and Get
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}
	if val, ok := cache.Get("key2"); !ok || val != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}
	if val, ok := cache.Get("key3"); !ok || val != "value3" {
		t.Errorf("Expected value3, got %v", val)
	}

	// Test eviction - key1 should be evicted as it's least recently used
	cache.Set("key4", "value4")
	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be evicted")
	}

	// Test access order - access key2, then add key5, key3 should be evicted
	cache.Get("key2")
	cache.Set("key5", "value5")
	if _, ok := cache.Get("key3"); ok {
		t.Error("Expected key3 to be evicted")
	}

	// Test Delete
	cache.Delete("key2")
	if _, ok := cache.Get("key2"); ok {
		t.Error("Expected key2 to be deleted")
	}

	// Test Clear
	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("Expected size 0, got %d", cache.Size())
	}
}

func TestLRUCacheExpiration(t *testing.T) {
	cache := NewLRU(3)

	cache.SetWithExpire("key1", "value1", 100*time.Millisecond)
	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}

	time.Sleep(150 * time.Millisecond)
	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be expired")
	}
}

func TestLRUCacheContains(t *testing.T) {
	cache := NewLRU(3)

	cache.Set("key1", "value1")
	if !cache.Contains("key1") {
		t.Error("Expected key1 to be contained")
	}
	if cache.Contains("key2") {
		t.Error("Expected key2 to not be contained")
	}

	cache.SetWithExpire("key3", "value3", 100*time.Millisecond)
	time.Sleep(150 * time.Millisecond)
	if cache.Contains("key3") {
		t.Error("Expected key3 to not be contained after expiration")
	}
}

func TestLRUCacheKeys(t *testing.T) {
	cache := NewLRU(3)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	keys := cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}
}

func TestLRUCacheConcurrent(t *testing.T) {
	cache := NewLRU(100)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + string(rune('0'+i%10))
			cache.Set(key, i)
			cache.Get(key)
			cache.Contains(key)
			cache.Delete(key)
		}(i)
	}
	wg.Wait()
}

func TestFIFOCache(t *testing.T) {
	cache := NewFIFO(3)

	// Test Set and Get
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}
	if val, ok := cache.Get("key2"); !ok || val != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}
	if val, ok := cache.Get("key3"); !ok || val != "value3" {
		t.Errorf("Expected value3, got %v", val)
	}

	// Test eviction - key1 should be evicted (first in, first out)
	cache.Set("key4", "value4")
	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be evicted")
	}

	// Test Delete
	cache.Delete("key2")
	if _, ok := cache.Get("key2"); ok {
		t.Error("Expected key2 to be deleted")
	}

	// Test Clear
	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("Expected size 0, got %d", cache.Size())
	}
}

func TestFIFOCacheExpiration(t *testing.T) {
	cache := NewFIFO(3)

	cache.SetWithExpire("key1", "value1", 100*time.Millisecond)
	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}

	time.Sleep(150 * time.Millisecond)
	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be expired")
	}
}

func TestFIFOCacheContains(t *testing.T) {
	cache := NewFIFO(3)

	cache.Set("key1", "value1")
	if !cache.Contains("key1") {
		t.Error("Expected key1 to be contained")
	}
	if cache.Contains("key2") {
		t.Error("Expected key2 to not be contained")
	}
}

func TestFIFOCacheKeys(t *testing.T) {
	cache := NewFIFO(3)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	keys := cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}
}

func TestFIFOCacheConcurrent(t *testing.T) {
	cache := NewFIFO(100)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + string(rune('0'+i%10))
			cache.Set(key, i)
			cache.Get(key)
			cache.Contains(key)
			cache.Delete(key)
		}(i)
	}
	wg.Wait()
}

func TestLFUCache(t *testing.T) {
	cache := NewLFU(3)

	// Test Set and Get
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}

	// Access key1 again to increase its frequency
	cache.Get("key1")

	// key2 and key3 have freq 1, key1 has freq 2
	// Adding key4 should evict key2 or key3 (both have min freq 1)
	cache.Set("key4", "value4")

	if cache.Size() != 3 {
		t.Errorf("Expected size 3, got %d", cache.Size())
	}

	// key1 should still be in cache (higher frequency)
	if _, ok := cache.Get("key1"); !ok {
		t.Error("Expected key1 to still be in cache")
	}

	// Test Delete
	cache.Delete("key1")
	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be deleted")
	}

	// Test Clear
	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("Expected size 0, got %d", cache.Size())
	}
}

func TestLFUCacheExpiration(t *testing.T) {
	cache := NewLFU(3)

	cache.SetWithExpire("key1", "value1", 100*time.Millisecond)
	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}

	time.Sleep(150 * time.Millisecond)
	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be expired")
	}
}

func TestLFUCacheContains(t *testing.T) {
	cache := NewLFU(3)

	cache.Set("key1", "value1")
	if !cache.Contains("key1") {
		t.Error("Expected key1 to be contained")
	}
	if cache.Contains("key2") {
		t.Error("Expected key2 to not be contained")
	}

	cache.SetWithExpire("key3", "value3", 100*time.Millisecond)
	time.Sleep(150 * time.Millisecond)
	if cache.Contains("key3") {
		t.Error("Expected key3 to not be contained after expiration")
	}
}

func TestLFUCacheKeys(t *testing.T) {
	cache := NewLFU(3)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	keys := cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}
}

func TestLFUCacheConcurrent(t *testing.T) {
	cache := NewLFU(100)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + string(rune('0'+i%10))
			cache.Set(key, i)
			cache.Get(key)
			cache.Contains(key)
			cache.Delete(key)
		}(i)
	}
	wg.Wait()
}

func TestTimedCache(t *testing.T) {
	cache := NewTimed(200 * time.Millisecond)

	// Test Set and Get
	cache.Set("key1", "value1")
	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}

	// Test expiration
	time.Sleep(250 * time.Millisecond)
	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be expired")
	}

	// Test SetWithExpire
	cache.SetWithExpire("key2", "value2", 100*time.Millisecond)
	if val, ok := cache.Get("key2"); !ok || val != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}

	time.Sleep(150 * time.Millisecond)
	if _, ok := cache.Get("key2"); ok {
		t.Error("Expected key2 to be expired")
	}

	// Test Delete
	cache.Set("key3", "value3")
	cache.Delete("key3")
	if _, ok := cache.Get("key3"); ok {
		t.Error("Expected key3 to be deleted")
	}

	// Test Clear
	cache.Set("key4", "value4")
	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("Expected size 0, got %d", cache.Size())
	}
}

func TestTimedCacheContains(t *testing.T) {
	cache := NewTimed(100 * time.Millisecond)

	cache.Set("key1", "value1")
	if !cache.Contains("key1") {
		t.Error("Expected key1 to be contained")
	}

	time.Sleep(150 * time.Millisecond)
	if cache.Contains("key1") {
		t.Error("Expected key1 to not be contained after expiration")
	}
}

func TestTimedCacheKeys(t *testing.T) {
	cache := NewTimed(200 * time.Millisecond)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	keys := cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	time.Sleep(250 * time.Millisecond)
	keys = cache.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys after expiration, got %d", len(keys))
	}
}

func TestTimedCacheConcurrent(t *testing.T) {
	cache := NewTimed(1 * time.Second)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + string(rune('0'+i%10))
			cache.Set(key, i)
			cache.Get(key)
			cache.Contains(key)
			cache.Delete(key)
		}(i)
	}
	wg.Wait()
}

func TestTimedCacheCleanup(t *testing.T) {
	cache := NewTimed(100 * time.Millisecond)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Wait for items to expire and cleanup to run (cleanup runs every 1 second)
	time.Sleep(1200 * time.Millisecond)

	// Cleanup should have removed expired items
	if cache.Size() != 0 {
		t.Errorf("Expected size 0 after cleanup, got %d", cache.Size())
	}
}
