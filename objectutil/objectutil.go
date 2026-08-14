// Package objectutil provides object utility functions similar to Java Hutool's ObjectUtil.
package objectutil

import (
	"fmt"
	"reflect"
)

// IsNil checks if value is nil
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		return rv.IsNil()
	default:
		return false
	}
}

// IsNotNull checks if value is not nil
func IsNotNull(v interface{}) bool {
	return !IsNil(v)
}

// IsEmpty checks if value is empty
func IsEmpty(v interface{}) bool {
	if IsNil(v) {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Slice, reflect.Array, reflect.Map:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	default:
		return false
	}
}

// IsNotEmpty checks if value is not empty
func IsNotEmpty(v interface{}) bool {
	return !IsEmpty(v)
}

// DefaultIfNil returns defaultVal if v is nil
func DefaultIfNil(v interface{}, defaultVal interface{}) interface{} {
	if IsNil(v) {
		return defaultVal
	}
	return v
}

// DefaultIfEmpty returns defaultVal if v is empty
func DefaultIfEmpty(v interface{}, defaultVal interface{}) interface{} {
	if IsEmpty(v) {
		return defaultVal
	}
	return v
}

// Equals checks if two values are equal
func Equals(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return reflect.DeepEqual(a, b)
}

// NotEquals checks if two values are not equal
func NotEquals(a, b interface{}) bool {
	return !Equals(a, b)
}

// Compare compares two values
// Returns -1 if a < b, 0 if a == b, 1 if a > b
func Compare(a, b interface{}) int {
	return compareValues(a, b)
}

func compareValues(a, b interface{}) int {
	if a == nil && b == nil {
		return 0
	}
	if a == nil {
		return -1
	}
	if b == nil {
		return 1
	}

	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// Same type comparison
	if va.Type() == vb.Type() {
		switch va.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			av, bv := va.Int(), vb.Int()
			if av < bv {
				return -1
			} else if av > bv {
				return 1
			}
			return 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			av, bv := va.Uint(), vb.Uint()
			if av < bv {
				return -1
			} else if av > bv {
				return 1
			}
			return 0
		case reflect.Float32, reflect.Float64:
			av, bv := va.Float(), vb.Float()
			if av < bv {
				return -1
			} else if av > bv {
				return 1
			}
			return 0
		case reflect.String:
			av, bv := va.String(), vb.String()
			if av < bv {
				return -1
			} else if av > bv {
				return 1
			}
			return 0
		}
	}

	// Fallback to string comparison
	sa := fmt.Sprintf("%v", a)
	sb := fmt.Sprintf("%v", b)
	if sa < sb {
		return -1
	} else if sa > sb {
		return 1
	}
	return 0
}

// TypeName returns type name of value
func TypeName(v interface{}) string {
	if v == nil {
		return "nil"
	}
	return reflect.TypeOf(v).String()
}

// TypeOf returns reflect.Type of value
func TypeOf(v interface{}) reflect.Type {
	if v == nil {
		return nil
	}
	return reflect.TypeOf(v)
}

// KindOf returns reflect.Kind of value
func KindOf(v interface{}) reflect.Kind {
	if v == nil {
		return reflect.Invalid
	}
	return reflect.TypeOf(v).Kind()
}

// IsType checks if value is of specified type
func IsType(v interface{}, typeName string) bool {
	return TypeName(v) == typeName
}

// IsKind checks if value is of specified kind
func IsKind(v interface{}, kind reflect.Kind) bool {
	return KindOf(v) == kind
}

// ToPtr returns pointer to value
func ToPtr(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	rv := reflect.ValueOf(v)
	ptr := reflect.New(rv.Type())
	ptr.Elem().Set(rv)
	return ptr.Interface()
}

// FromPtr dereferences pointer
func FromPtr(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		return rv.Elem().Interface()
	}
	return v
}

// Copy deep copies value
func Copy(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return reflect.ValueOf(v).Interface()
}

// String returns string representation of value
func String(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

// Clone deep clones object using reflection
func Clone(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		newPtr := reflect.New(rv.Elem().Type())
		newPtr.Elem().Set(rv.Elem())
		return newPtr.Interface()
	}
	return v
}
