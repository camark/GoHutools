package boolutil

import "strings"

// IsTrue reports whether v is true.
func IsTrue(v bool) bool { return v }

// IsFalse reports whether v is false.
func IsFalse(v bool) bool { return !v }

// Parse parses a strict boolean string ("true"/"false", case-insensitive).
// It returns ok=false for any other input.
func Parse(s string) (ok, value bool) {
	switch strings.ToLower(s) {
	case "true":
		return true, true
	case "false":
		return true, false
	}
	return false, false
}

// ParseLenient parses common truthy/falsy strings:
// true/false, 1/0, yes/no, on/off (case-insensitive).
// Unknown input is treated as false.
func ParseLenient(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "1", "yes", "on", "y":
		return true
	default:
		return false
	}
}

// IsBoolean reports whether s is a strict boolean string.
func IsBoolean(s string) bool {
	ok, _ := Parse(s)
	return ok
}

// And returns true only when every element is true (empty => false).
func And(vals ...bool) bool {
	for _, v := range vals {
		if !v {
			return false
		}
	}
	return len(vals) > 0
}

// Or returns true when at least one element is true (empty => false).
func Or(vals ...bool) bool {
	for _, v := range vals {
		if v {
			return true
		}
	}
	return false
}

// Xor returns true when the number of true elements is odd.
func Xor(vals ...bool) bool {
	var count int
	for _, v := range vals {
		if v {
			count++
		}
	}
	return count%2 == 1
}

// ToInt maps true to 1 and false to 0.
func ToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

// ToString renders v as "true" or "false".
func ToString(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

// Negate returns the logical opposite of v.
func Negate(v bool) bool { return !v }

// Compare orders booleans as false < true. Returns -1, 0, or 1.
func Compare(a, b bool) int {
	if a == b {
		return 0
	}
	if a {
		return 1
	}
	return -1
}

// IsTrueValue reports whether any of the common boolean representations
// (bool, string forms, number 1) holds a truthy value.
func IsTrueValue(v interface{}) bool {
	switch x := v.(type) {
	case bool:
		return x
	case string:
		return ParseLenient(x)
	case int:
		return x != 0
	case int64:
		return x != 0
	case float64:
		return x != 0
	}
	return false
}