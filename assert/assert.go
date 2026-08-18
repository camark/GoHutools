package assert

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// errorf builds an error, using a custom message if provided, otherwise a default wording.
func errorf(defaultMsg string, args ...interface{}) error {
	if len(args) > 0 {
		if msg, ok := args[0].(string); ok {
			return fmt.Errorf("assertion failed: %s", msg)
		}
	}
	return fmt.Errorf("assertion failed: %s", defaultMsg)
}

// isNilValue reports whether v is nil or a typed nil (nil pointer/interface/map/slice/func/chan).
func isNilValue(v interface{}, rv reflect.Value) bool {
	if v == nil {
		return true
	}
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
		return rv.IsNil()
	}
	return false
}

// True asserts that condition is true.
func True(condition bool, args ...interface{}) error {
	if !condition {
		return errorf("expected true, got false", args...)
	}
	return nil
}

// False asserts that condition is false.
func False(condition bool, args ...interface{}) error {
	if condition {
		return errorf("expected false, got true", args...)
	}
	return nil
}

// Null asserts that v is nil (including typed nil).
func Null(v interface{}, args ...interface{}) error {
	rv := reflect.ValueOf(v)
	if !isNilValue(v, rv) {
		return errorf(fmt.Sprintf("expected nil, got %v", v), args...)
	}
	return nil
}

// IsNull is an alias of Null.
func IsNull(v interface{}, args ...interface{}) error {
	return Null(v, args...)
}

// NotNull asserts that v is not nil (including typed nil).
func NotNull(v interface{}, args ...interface{}) error {
	rv := reflect.ValueOf(v)
	if isNilValue(v, rv) {
		return errorf("expected non-nil value, got nil", args...)
	}
	return nil
}

// NotEmpty asserts that v is not empty.
// A value is considered empty when it is nil, a zero number, an empty string/blank string,
// or a slice/map/array/chan with length zero.
func NotEmpty(v interface{}, args ...interface{}) error {
	rv := reflect.ValueOf(v)
	if v == nil {
		return errorf("expected non-empty value, got nil", args...)
	}
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface:
		if rv.IsNil() {
			return errorf("expected non-empty value, got nil", args...)
		}
		return NotEmpty(rv.Elem().Interface(), args...)
	case reflect.String:
		s := rv.String()
		if s == "" || strings.TrimSpace(s) == "" {
			return errorf(fmt.Sprintf("expected non-empty string, got %q", s), args...)
		}
	case reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		if rv.Len() == 0 {
			return errorf(fmt.Sprintf("expected non-empty %v, got length 0", rv.Kind()), args...)
		}
	default:
		// numbers, bools, structs: zero value means empty
		if rv.IsZero() {
			return errorf(fmt.Sprintf("expected non-zero value of type %v", rv.Type()), args...)
		}
	}
	return nil
}

// NotBlank asserts that the string v is not blank (has at least one non-space rune).
// Non-string values always pass.
func NotBlank(v interface{}, args ...interface{}) error {
	rv := reflect.ValueOf(v)
	if v == nil {
		return errorf("expected non-blank string, got nil", args...)
	}
	if rv.Kind() == reflect.String && strings.TrimSpace(rv.String()) == "" {
		return errorf(fmt.Sprintf("expected non-blank string, got %q", rv.String()), args...)
	}
	return nil
}

// Equal asserts that expected and actual are deeply equal.
func Equal(expected, actual interface{}, args ...interface{}) error {
	if !reflect.DeepEqual(expected, actual) {
		return errorf(fmt.Sprintf("expected %v, got %v", expected, actual), args...)
	}
	return nil
}

// NotEqual asserts that expected and actual are not deeply equal.
func NotEqual(unexpected, actual interface{}, args ...interface{}) error {
	if reflect.DeepEqual(unexpected, actual) {
		return errorf(fmt.Sprintf("expected NOT equal to %v, got %v", unexpected, actual), args...)
	}
	return nil
}

// IsInstanceOf asserts that v is an instance of the given type.
// typ may be either a type name string or a zero-value instance of the target type.
func IsInstanceOf(v interface{}, typ interface{}, args ...interface{}) error {
	rv := reflect.ValueOf(v)
	actualType := rv.Type()

	switch t := typ.(type) {
	case string:
		if actualType.String() != t {
			return errorf(fmt.Sprintf("expected type %s, got %s", t, actualType), args...)
		}
	default:
		expectedType := reflect.TypeOf(typ)
		if actualType != expectedType {
			return errorf(fmt.Sprintf("expected type %v, got %v", expectedType, actualType), args...)
		}
	}
	return nil
}

// Matches asserts that s matches the given regex pattern.
func Matches(pattern, s interface{}, args ...interface{}) error {
	re, err := regexp.Compile(fmt.Sprintf("%v", pattern))
	if err != nil {
		return fmt.Errorf("matches: invalid pattern %v: %w", pattern, err)
	}
	if !re.MatchString(fmt.Sprintf("%v", s)) {
		return errorf(fmt.Sprintf("expected %v to match %v", s, pattern), args...)
	}
	return nil
}

// NoError asserts that err is nil.
func NoError(err error, args ...interface{}) error {
	if err != nil {
		return errorf(fmt.Sprintf("expected no error, got: %v", err), args...)
	}
	return nil
}

// ErrorContains asserts that err is non-nil and its message contains substr.
func ErrorContains(err error, substr string, args ...interface{}) error {
	if err == nil {
		return errorf("expected an error containing %q, got nil", args...)
	}
	if !strings.Contains(err.Error(), substr) {
		return errorf(fmt.Sprintf("expected error containing %q, got: %v", substr, err), args...)
	}
	return nil
}

// LessThan asserts that a < b using < on comparable values.
func LessThan(a, b interface{}, args ...interface{}) error {
	if !isLess(a, b) {
		return errorf(fmt.Sprintf("expected %v < %v", a, b), args...)
	}
	return nil
}

// GreaterThan asserts that a > b using > on comparable values.
func GreaterThan(a, b interface{}, args ...interface{}) error {
	if !isGreater(a, b) {
		return errorf(fmt.Sprintf("expected %v > %v", a, b), args...)
	}
	return nil
}

// NotNegative asserts that a is not a negative number.
func NotNegative(a interface{}, args ...interface{}) error {
	rv := reflect.ValueOf(a)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if rv.Int() < 0 {
			return errorf(fmt.Sprintf("expected non-negative, got %v", a), args...)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		// unsigned values are always non-negative
	case reflect.Float32, reflect.Float64:
		if rv.Float() < 0 {
			return errorf(fmt.Sprintf("expected non-negative, got %v", a), args...)
		}
	default:
		return errorf(fmt.Sprintf("expected a number, got %v", a), args...)
	}
	return nil
}

func isLess(a, b interface{}) bool {
	ra := reflect.ValueOf(a)
	rb := reflect.ValueOf(b)
	if !ra.IsValid() || !rb.IsValid() || ra.Kind() != rb.Kind() {
		return false
	}
	switch ra.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return ra.Int() < rb.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return ra.Uint() < rb.Uint()
	case reflect.Float32, reflect.Float64:
		return ra.Float() < rb.Float()
	case reflect.String:
		return ra.String() < rb.String()
	case reflect.Bool:
		return !ra.Bool() && rb.Bool()
	default:
		return false
	}
}

func isGreater(a, b interface{}) bool {
	ra := reflect.ValueOf(a)
	rb := reflect.ValueOf(b)
	if !ra.IsValid() || !rb.IsValid() || ra.Kind() != rb.Kind() {
		return false
	}
	switch ra.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return ra.Int() > rb.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return ra.Uint() > rb.Uint()
	case reflect.Float32, reflect.Float64:
		return ra.Float() > rb.Float()
	case reflect.String:
		return ra.String() > rb.String()
	case reflect.Bool:
		return ra.Bool() && !rb.Bool()
	default:
		return false
	}
}