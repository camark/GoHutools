package collutil

import (
	"cmp"
	"slices"
)

// IsEmpty checks if slice is empty
func IsEmpty[T any](coll []T) bool {
	return len(coll) == 0
}

// IsNotEmpty checks if slice is not empty
func IsNotEmpty[T any](coll []T) bool {
	return len(coll) > 0
}

// Contains checks if slice contains item
func Contains[T comparable](coll []T, item T) bool {
	return slices.Contains(coll, item)
}

// ContainsAll checks if slice contains all items
func ContainsAll[T comparable](coll []T, items ...T) bool {
	for _, item := range items {
		if !slices.Contains(coll, item) {
			return false
		}
	}
	return true
}

// ContainsAny checks if slice contains any of the items
func ContainsAny[T comparable](coll []T, items ...T) bool {
	for _, item := range items {
		if slices.Contains(coll, item) {
			return true
		}
	}
	return false
}

// AddIfAbsent adds item if not already present
func AddIfAbsent[T comparable](coll []T, item T) []T {
	if !slices.Contains(coll, item) {
		return append(coll, item)
	}
	return coll
}

// Remove removes first occurrence of item
func Remove[T comparable](coll []T, item T) []T {
	for i, v := range coll {
		if v == item {
			return slices.Delete(coll, i, i+1)
		}
	}
	return coll
}

// RemoveAll removes all occurrences of item
func RemoveAll[T comparable](coll []T, item T) []T {
	result := make([]T, 0, len(coll))
	for _, v := range coll {
		if v != item {
			result = append(result, v)
		}
	}
	return result
}

// Distinct removes duplicates preserving order
func Distinct[T comparable](coll []T) []T {
	seen := make(map[T]struct{}, len(coll))
	result := make([]T, 0, len(coll))
	for _, v := range coll {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Reverse reverses the slice
func Reverse[T any](coll []T) []T {
	result := make([]T, len(coll))
	copy(result, coll)
	slices.Reverse(result)
	return result
}

// Sub extracts sub-slice from start to end
func Sub[T any](coll []T, start, end int) []T {
	if start < 0 {
		start = 0
	}
	if end > len(coll) {
		end = len(coll)
	}
	if start >= end {
		return []T{}
	}
	return coll[start:end]
}

// Split splits slice into chunks of given size
func Split[T any](coll []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	if len(coll) == 0 {
		return [][]T{}
	}
	var result [][]T
	for i := 0; i < len(coll); i += size {
		end := i + size
		if end > len(coll) {
			end = len(coll)
		}
		result = append(result, coll[i:end])
	}
	return result
}

// Filter returns elements matching predicate
func Filter[T any](coll []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(coll))
	for _, v := range coll {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map transforms each element using mapper
func Map[T, R any](coll []T, mapper func(T) R) []R {
	result := make([]R, len(coll))
	for i, v := range coll {
		result[i] = mapper(v)
	}
	return result
}

// ForEach executes consumer for each element
func ForEach[T any](coll []T, consumer func(T)) {
	for _, v := range coll {
		consumer(v)
	}
}

// ForEachIndexed executes consumer with index for each element
func ForEachIndexed[T any](coll []T, consumer func(int, T)) {
	for i, v := range coll {
		consumer(i, v)
	}
}

// GroupBy groups elements by key
func GroupBy[T any, K comparable](coll []T, keyFunc func(T) K) map[K][]T {
	result := make(map[K][]T, len(coll))
	for _, v := range coll {
		key := keyFunc(v)
		result[key] = append(result[key], v)
	}
	return result
}

// ToMap converts slice to map using key and value functions
func ToMap[T any, K comparable, V any](coll []T, keyFunc func(T) K, valFunc func(T) V) map[K]V {
	result := make(map[K]V, len(coll))
	for _, v := range coll {
		result[keyFunc(v)] = valFunc(v)
	}
	return result
}

// Sort sorts ordered slice
func Sort[T cmp.Ordered](coll []T) []T {
	result := make([]T, len(coll))
	copy(result, coll)
	slices.Sort(result)
	return result
}

// SortBy sorts using custom comparator
func SortBy[T any](coll []T, less func(T, T) bool) []T {
	result := make([]T, len(coll))
	copy(result, coll)
	slices.SortFunc(result, func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if less(b, a) {
			return 1
		}
		return 0
	})
	return result
}

// Min returns minimum element
func Min[T cmp.Ordered](coll []T) (T, bool) {
	if len(coll) == 0 {
		var zero T
		return zero, false
	}
	min := coll[0]
	for _, v := range coll[1:] {
		if v < min {
			min = v
		}
	}
	return min, true
}

// Max returns maximum element
func Max[T cmp.Ordered](coll []T) (T, bool) {
	if len(coll) == 0 {
		var zero T
		return zero, false
	}
	max := coll[0]
	for _, v := range coll[1:] {
		if v > max {
			max = v
		}
	}
	return max, true
}

// Sum returns sum of elements
func Sum[T cmp.Ordered](coll []T) T {
	var sum T
	for _, v := range coll {
		sum += v
	}
	return sum
}

// Flatten flattens nested slices
func Flatten[T any](colls [][]T) []T {
	totalLen := 0
	for _, coll := range colls {
		totalLen += len(coll)
	}
	result := make([]T, 0, totalLen)
	for _, coll := range colls {
		result = append(result, coll...)
	}
	return result
}

// Chunk splits into chunks (alias for Split)
func Chunk[T any](coll []T, size int) [][]T {
	return Split(coll, size)
}

// Zip creates map from keys and values slices
func Zip[T comparable, U any](keys []T, values []U) map[T]U {
	minLen := len(keys)
	if len(values) < minLen {
		minLen = len(values)
	}
	result := make(map[T]U, minLen)
	for i := 0; i < minLen; i++ {
		result[keys[i]] = values[i]
	}
	return result
}

// Keys returns map keys
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// Values returns map values
func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// First returns first element or zero value
func First[T any](coll []T) (T, bool) {
	if len(coll) == 0 {
		var zero T
		return zero, false
	}
	return coll[0], true
}

// Last returns last element or zero value
func Last[T any](coll []T) (T, bool) {
	if len(coll) == 0 {
		var zero T
		return zero, false
	}
	return coll[len(coll)-1], true
}

// Find finds first element matching predicate
func Find[T any](coll []T, predicate func(T) bool) (T, bool) {
	for _, v := range coll {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindIndex finds index of first element matching predicate
func FindIndex[T any](coll []T, predicate func(T) bool) int {
	for i, v := range coll {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// AnyMatch checks if any element matches predicate
func AnyMatch[T any](coll []T, predicate func(T) bool) bool {
	for _, v := range coll {
		if predicate(v) {
			return true
		}
	}
	return false
}

// AllMatch checks if all elements match predicate
func AllMatch[T any](coll []T, predicate func(T) bool) bool {
	for _, v := range coll {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// NoneMatch checks if no elements match predicate
func NoneMatch[T any](coll []T, predicate func(T) bool) bool {
	for _, v := range coll {
		if predicate(v) {
			return false
		}
	}
	return true
}

// Reduce reduces slice to single value
func Reduce[T, R any](coll []T, initial R, reducer func(R, T) R) R {
	result := initial
	for _, v := range coll {
		result = reducer(result, v)
	}
	return result
}

// Count counts elements matching predicate
func Count[T any](coll []T, predicate func(T) bool) int {
	count := 0
	for _, v := range coll {
		if predicate(v) {
			count++
		}
	}
	return count
}

// Equal checks if two slices are equal
func Equal[T comparable](a, b []T) bool {
	return slices.Equal(a, b)
}

// Union returns union of two slices
func Union[T comparable](a, b []T) []T {
	seen := make(map[T]struct{}, len(a)+len(b))
	result := make([]T, 0, len(a)+len(b))
	for _, v := range a {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	for _, v := range b {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Intersection returns intersection of two slices
func Intersection[T comparable](a, b []T) []T {
	setB := make(map[T]struct{}, len(b))
	for _, v := range b {
		setB[v] = struct{}{}
	}
	seen := make(map[T]struct{}, len(a))
	result := make([]T, 0, min(len(a), len(b)))
	for _, v := range a {
		if _, ok := setB[v]; ok {
			if _, ok := seen[v]; !ok {
				seen[v] = struct{}{}
				result = append(result, v)
			}
		}
	}
	return result
}

// Difference returns elements in a but not in b
func Difference[T comparable](a, b []T) []T {
	setB := make(map[T]struct{}, len(b))
	for _, v := range b {
		setB[v] = struct{}{}
	}
	seen := make(map[T]struct{}, len(a))
	result := make([]T, 0, len(a))
	for _, v := range a {
		if _, ok := setB[v]; !ok {
			if _, ok := seen[v]; !ok {
				seen[v] = struct{}{}
				result = append(result, v)
			}
		}
	}
	return result
}

// ToSlice converts variadic args to slice
func ToSlice[T any](items ...T) []T {
	return items
}

// Range generates sequence from start to end
func Range(start, end int) []int {
	if start >= end {
		return []int{}
	}
	result := make([]int, end-start)
	for i := range result {
		result[i] = start + i
	}
	return result
}

// RangeWithStep generates sequence with step
func RangeWithStep(start, end, step int) []int {
	if step == 0 {
		return []int{}
	}
	if step > 0 && start >= end {
		return []int{}
	}
	if step < 0 && start <= end {
		return []int{}
	}
	var result []int
	for i := start; (step > 0 && i < end) || (step < 0 && i > end); i += step {
		result = append(result, i)
	}
	return result
}
