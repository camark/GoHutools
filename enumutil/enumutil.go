package enumutil

import (
	"errors"
	"fmt"
)

// This package adapts Hutool's EnumUtil to Go's type-oriented enums:
// since Go cannot reflect package-level constants, callers pass the
// slice of valid enum values (e.g. `[]Color{Red, Green, Blue}`).

// FromString finds the first enum value in values whose string form
// equals s. Returns an error when s matches no value.
func FromString[T ~string](values []T, s string) (T, error) {
	for _, v := range values {
		if string(v) == s {
			return v, nil
		}
	}
	var zero T
	return zero, fmt.Errorf("enumutil: no value %q in %v", s, namesOf(values))
}

// MustFromString is like FromString but panics when s matches nothing.
func MustFromString[T ~string](values []T, s string) T {
	v, err := FromString(values, s)
	if err != nil {
		panic(err)
	}
	return v
}

// Of is an alias of MustFromString for callers who prefer a short name.
func Of[T ~string](values []T, s string) T {
	return MustFromString(values, s)
}

// ContainsValue reports whether v appears in values.
func ContainsValue[T comparable](values []T, v T) bool {
	for _, item := range values {
		if item == v {
			return true
		}
	}
	return false
}

// Names returns the string form of every value, preserving order.
func Names[T ~string](values []T) []string {
	out := make([]string, len(values))
	for i, v := range values {
		out[i] = string(v)
	}
	return out
}

// Index returns the position of v in values, or -1 when absent.
func Index[T comparable](values []T, v T) int {
	for i, item := range values {
		if item == v {
			return i
		}
	}
	return -1
}

// ToValueMap builds a map from the string form of each value to the value.
func ToValueMap[T ~string](values []T) map[string]T {
	m := make(map[string]T, len(values))
	for _, v := range values {
		m[string(v)] = v
	}
	return m
}

// Validate reports an error listing any values that are missing, empty
// or duplicated — useful for asserting an enum list is sound.
func Validate[T ~string](values []T) error {
	seen := make(map[string]struct{}, len(values))
	for _, v := range values {
		s := string(v)
		if s == "" {
			return errors.New("enumutil: empty enum value")
		}
		if _, dup := seen[s]; dup {
			return fmt.Errorf("enumutil: duplicate enum value %q", s)
		}
		seen[s] = struct{}{}
	}
	return nil
}

func namesOf[T ~string](values []T) []string { return Names(values) }