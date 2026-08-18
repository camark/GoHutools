package maputil

import "reflect"

// Entry represents a key-value pair.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// IsEmpty checks if map is empty.
func IsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) == 0
}

// IsNotEmpty checks if map is not empty.
func IsNotEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) > 0
}

// ContainsKey checks if map contains key.
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// ContainsValue checks if map contains value.
// Uses reflect.DeepEqual for comparison so it works with non-comparable value types.
func ContainsValue[K comparable, V any](m map[K]V, value V) bool {
	for _, v := range m {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Get gets value with default. If key exists, returns the value; otherwise returns defaultVal.
func Get[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultVal
}

// GetOrDefault gets value or returns default. Alias for Get.
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	return Get(m, key, defaultVal)
}

// PutIfAbsent puts value if key is absent. Returns the existing value if present, or the new value if inserted.
func PutIfAbsent[K comparable, V any](m map[K]V, key K, value V) V {
	if v, ok := m[key]; ok {
		return v
	}
	m[key] = value
	return value
}

// Remove removes key from map.
func Remove[K comparable, V any](m map[K]V, key K) {
	delete(m, key)
}

// Keys returns all keys of the map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values of the map.
func Values[K comparable, V any](m map[K]V) []V {
	vals := make([]V, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}

// ForEach iterates map with consumer function.
func ForEach[K comparable, V any](m map[K]V, consumer func(K, V)) {
	for k, v := range m {
		consumer(k, v)
	}
}

// Filter filters map by predicate. Returns a new map containing entries that match the predicate.
func Filter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// Transform transforms map values using a transformer function. Returns a new map with transformed values.
func Transform[K comparable, V, R any](m map[K]V, transformer func(K, V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = transformer(k, v)
	}
	return result
}

// Merge merges two maps. Values from m2 override values from m1 for duplicate keys.
func Merge[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V, len(m1)+len(m2))
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

// Invert inverts map (keys become values, values become keys).
// Both K and V must be comparable for this to work.
func Invert[K, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// FromKeys creates map from keys with default value.
func FromKeys[K comparable, V any](keys []K, defaultVal V) map[K]V {
	result := make(map[K]V, len(keys))
	for _, k := range keys {
		result[k] = defaultVal
	}
	return result
}

// FromEntries creates map from key-value pairs.
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	result := make(map[K]V, len(entries))
	for _, e := range entries {
		result[e.Key] = e.Value
	}
	return result
}

// ToEntries converts map to entry slice.
func ToEntries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, Entry[K, V]{Key: k, Value: v})
	}
	return entries
}

// Equal checks if two maps are equal (same keys and values).
func Equal[K comparable, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// Size returns map size.
func Size[K comparable, V any](m map[K]V) int {
	return len(m)
}

// Copy copies map to a new map.
func Copy[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}
