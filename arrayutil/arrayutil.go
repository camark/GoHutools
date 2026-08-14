package arrayutil

import (
	"cmp"
	"fmt"
	"math/rand"
	"strings"
)

// Pair represents a pair of values.
type Pair[T, U any] struct {
	First  T
	Second U
}

// IsEmpty checks if array is empty.
func IsEmpty[T any](arr []T) bool {
	return len(arr) == 0
}

// IsNotEmpty checks if array is not empty.
func IsNotEmpty[T any](arr []T) bool {
	return len(arr) > 0
}

// Contains checks if array contains element.
func Contains[T comparable](arr []T, elem T) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

// IndexOf returns index of element (-1 if not found).
func IndexOf[T comparable](arr []T, elem T) int {
	for i, v := range arr {
		if v == elem {
			return i
		}
	}
	return -1
}

// LastIndexOf returns last index of element (-1 if not found).
func LastIndexOf[T comparable](arr []T, elem T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == elem {
			return i
		}
	}
	return -1
}

// Fill fills array with value. Returns the same slice with all elements set to val.
func Fill[T any](arr []T, val T) []T {
	for i := range arr {
		arr[i] = val
	}
	return arr
}

// CopyOf copies array to new array of specified length.
// If newLength is greater than arr length, zero values are appended.
// If newLength is less than arr length, the result is truncated.
func CopyOf[T any](arr []T, newLength int) []T {
	if newLength < 0 {
		newLength = 0
	}
	result := make([]T, newLength)
	copy(result, arr)
	return result
}

// CopyOfRange copies range of array [from, to).
// If from < 0, it is set to 0. If to > len(arr), it is set to len(arr).
func CopyOfRange[T any](arr []T, from, to int) []T {
	if from < 0 {
		from = 0
	}
	if to > len(arr) {
		to = len(arr)
	}
	if from > to {
		from = to
	}
	result := make([]T, to-from)
	copy(result, arr[from:to])
	return result
}

// Insert inserts element at index. Returns a new slice.
func Insert[T any](arr []T, index int, elem T) []T {
	if index < 0 {
		index = 0
	}
	if index > len(arr) {
		index = len(arr)
	}
	result := make([]T, 0, len(arr)+1)
	result = append(result, arr[:index]...)
	result = append(result, elem)
	result = append(result, arr[index:]...)
	return result
}

// RemoveAt removes element at index. Returns a new slice.
func RemoveAt[T any](arr []T, index int) []T {
	if index < 0 || index >= len(arr) {
		return arr
	}
	result := make([]T, 0, len(arr)-1)
	result = append(result, arr[:index]...)
	result = append(result, arr[index+1:]...)
	return result
}

// Swap swaps two elements in the array.
func Swap[T any](arr []T, i, j int) {
	if i >= 0 && i < len(arr) && j >= 0 && j < len(arr) {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Reverse reverses array in place. Returns the same slice.
func Reverse[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// Rotate rotates array by n positions to the right. Returns a new slice.
// Positive n rotates right, negative n rotates left.
func Rotate[T any](arr []T, n int) []T {
	if len(arr) == 0 {
		return arr
	}
	n = n % len(arr)
	if n < 0 {
		n += len(arr)
	}
	result := make([]T, len(arr))
	copy(result[n:], arr[:len(arr)-n])
	copy(result[:n], arr[len(arr)-n:])
	return result
}

// Unique removes duplicates. Returns a new slice with unique elements in order of first appearance.
func Unique[T comparable](arr []T) []T {
	seen := make(map[T]struct{}, len(arr))
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// ToInterfaceSlice converts typed slice to interface slice.
func ToInterfaceSlice[T any](arr []T) []interface{} {
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		result[i] = v
	}
	return result
}

// FromInterfaceSlice converts interface slice to typed slice.
// Returns the converted slice and true if all elements were of the correct type,
// or nil and false if any element could not be converted.
func FromInterfaceSlice[T any](arr []interface{}) ([]T, bool) {
	result := make([]T, len(arr))
	for i, v := range arr {
		val, ok := v.(T)
		if !ok {
			return nil, false
		}
		result[i] = val
	}
	return result, true
}

// Filter filters array by predicate. Returns a new slice containing elements that match.
func Filter[T any](arr []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map transforms array elements using a mapper function.
func Map[T, R any](arr []T, mapper func(T) R) []R {
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = mapper(v)
	}
	return result
}

// Any checks if any element matches the predicate.
func Any[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if predicate(v) {
			return true
		}
	}
	return false
}

// All checks if all elements match the predicate.
func All[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// None checks if no elements match the predicate.
func None[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if predicate(v) {
			return false
		}
	}
	return true
}

// Join joins array elements with separator using fmt.Sprint for each element.
func Join[T any](arr []T, separator string) string {
	if len(arr) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprint(arr[0]))
	for _, v := range arr[1:] {
		sb.WriteString(separator)
		sb.WriteString(fmt.Sprint(v))
	}
	return sb.String()
}

// Equal checks if two arrays are equal (same length and same elements in order).
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Concat concatenates arrays into a single new slice.
func Concat[T any](arrays ...[]T) []T {
	totalLen := 0
	for _, arr := range arrays {
		totalLen += len(arr)
	}
	result := make([]T, 0, totalLen)
	for _, arr := range arrays {
		result = append(result, arr...)
	}
	return result
}

// First returns the first element and true, or zero value and false if empty.
func First[T any](arr []T) (T, bool) {
	if len(arr) == 0 {
		var zero T
		return zero, false
	}
	return arr[0], true
}

// Last returns the last element and true, or zero value and false if empty.
func Last[T any](arr []T) (T, bool) {
	if len(arr) == 0 {
		var zero T
		return zero, false
	}
	return arr[len(arr)-1], true
}

// Shuffle shuffles array randomly. Returns a new slice.
func Shuffle[T any](arr []T) []T {
	result := make([]T, len(arr))
	copy(result, arr)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

// Min returns the minimum element and true, or zero value and false if empty.
func Min[T cmp.Ordered](arr []T) (T, bool) {
	if len(arr) == 0 {
		var zero T
		return zero, false
	}
	min := arr[0]
	for _, v := range arr[1:] {
		if v < min {
			min = v
		}
	}
	return min, true
}

// Max returns the maximum element and true, or zero value and false if empty.
func Max[T cmp.Ordered](arr []T) (T, bool) {
	if len(arr) == 0 {
		var zero T
		return zero, false
	}
	max := arr[0]
	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
	}
	return max, true
}

// Flatten flattens nested arrays into a single slice.
func Flatten[T any](arr [][]T) []T {
	totalLen := 0
	for _, inner := range arr {
		totalLen += len(inner)
	}
	result := make([]T, 0, totalLen)
	for _, inner := range arr {
		result = append(result, inner...)
	}
	return result
}

// Chunk splits array into chunks of specified size. The last chunk may be smaller.
func Chunk[T any](arr []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	var result [][]T
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}

// Window creates sliding windows of specified size.
func Window[T any](arr []T, size int) [][]T {
	if size <= 0 || size > len(arr) {
		return nil
	}
	result := make([][]T, 0, len(arr)-size+1)
	for i := 0; i <= len(arr)-size; i++ {
		window := make([]T, size)
		copy(window, arr[i:i+size])
		result = append(result, window)
	}
	return result
}

// Zip combines two arrays into pairs. The length is the minimum of the two arrays.
func Zip[T, U any](a []T, b []U) []Pair[T, U] {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	result := make([]Pair[T, U], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = Pair[T, U]{First: a[i], Second: b[i]}
	}
	return result
}

// Unzip separates pairs into two arrays.
func Unzip[T, U any](pairs []Pair[T, U]) ([]T, []U) {
	a := make([]T, len(pairs))
	b := make([]U, len(pairs))
	for i, p := range pairs {
		a[i] = p.First
		b[i] = p.Second
	}
	return a, b
}
