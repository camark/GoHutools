package objectutil

import (
	"reflect"
	"testing"
)

func TestIsNil(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want bool
	}{
		{"nil", nil, true},
		{"nil pointer", func() interface{} { var p *int; return p }(), true},
		{"nil slice", func() interface{} { var s []int; return s }(), true},
		{"nil map", func() interface{} { var m map[string]int; return m }(), true},
		{"nil chan", func() interface{} { var c chan int; return c }(), true},
		{"nil func", func() interface{} { var f func(); return f }(), true},
		{"int 0", 0, false},
		{"empty string", "", false},
		{"empty slice (non-nil)", []int{}, false},
		{"non-nil pointer", func() interface{} { v := 1; return &v }(), false},
		{"struct", struct{}{}, false},
		{"bool false", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.v); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotNull(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want bool
	}{
		{"nil", nil, false},
		{"nil pointer", func() interface{} { var p *int; return p }(), false},
		{"int 0", 0, true},
		{"empty string", "", true},
		{"non-nil pointer", func() interface{} { v := 1; return &v }(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotNull(tt.v); got != tt.want {
				t.Errorf("IsNotNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want bool
	}{
		{"nil", nil, true},
		{"nil pointer", func() interface{} { var p *int; return p }(), true},
		{"empty string", "", true},
		{"non-empty string", "hello", false},
		{"empty slice", []int{}, true},
		{"non-empty slice", []int{1}, false},
		{"empty map", map[string]int{}, true},
		{"non-empty map", map[string]int{"a": 1}, false},
		{"bool false", false, true},
		{"bool true", true, false},
		{"int 0", 0, true},
		{"int 1", 1, false},
		{"int8 0", int8(0), true},
		{"int8 1", int8(1), false},
		{"int16 0", int16(0), true},
		{"int16 1", int16(1), false},
		{"int32 0", int32(0), true},
		{"int32 1", int32(1), false},
		{"int64 0", int64(0), true},
		{"int64 1", int64(1), false},
		{"uint 0", uint(0), true},
		{"uint 1", uint(1), false},
		{"uint8 0", uint8(0), true},
		{"uint8 1", uint8(1), false},
		{"uint16 0", uint16(0), true},
		{"uint16 1", uint16(1), false},
		{"uint32 0", uint32(0), true},
		{"uint32 1", uint32(1), false},
		{"uint64 0", uint64(0), true},
		{"uint64 1", uint64(1), false},
		{"float32 0", float32(0), true},
		{"float32 1.5", float32(1.5), false},
		{"float64 0", float64(0), true},
		{"float64 1.5", float64(1.5), false},
		{"struct (default branch)", struct{}{}, false},
		{"array empty", [0]int{}, true},
		{"array non-empty", [1]int{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.v); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want bool
	}{
		{"nil", nil, false},
		{"empty string", "", false},
		{"non-empty string", "hello", true},
		{"non-empty slice", []int{1}, true},
		{"empty slice", []int{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotEmpty(tt.v); got != tt.want {
				t.Errorf("IsNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultIfNil(t *testing.T) {
	tests := []struct {
		name       string
		v          interface{}
		defaultVal interface{}
		want       interface{}
	}{
		{"nil returns default", nil, "default", "default"},
		{"non-nil returns value", "value", "default", "value"},
		{"nil pointer returns default", func() interface{} { var p *int; return p }(), 42, 42},
		{"non-nil pointer returns value", func() interface{} { v := 7; return &v }(), 42, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultIfNil(tt.v, tt.defaultVal)
			if tt.name == "non-nil pointer returns value" {
				// value is the pointer itself, just check it's not the default
				if got == tt.defaultVal {
					t.Errorf("DefaultIfNil() returned default when value is not nil")
				}
				return
			}
			if got != tt.want {
				t.Errorf("DefaultIfNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		name       string
		v          interface{}
		defaultVal interface{}
		want       interface{}
	}{
		{"empty string returns default", "", "default", "default"},
		{"non-empty string returns value", "value", "default", "value"},
		{"zero int returns default", 0, 42, 42},
		{"non-zero int returns value", 5, 42, 5},
		{"nil returns default", nil, "default", "default"},
		{"empty slice returns default", []int{}, []int{9}, []int{9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultIfEmpty(tt.v, tt.defaultVal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultIfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		name string
		a, b interface{}
		want bool
	}{
		{"both nil", nil, nil, true},
		{"first nil", nil, 0, false},
		{"second nil", 0, nil, false},
		{"equal ints", 1, 1, true},
		{"different ints", 1, 2, false},
		{"equal strings", "hello", "hello", true},
		{"different strings", "hello", "world", false},
		{"equal slices", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"different slices", []int{1, 2}, []int{3, 4}, false},
		{"equal maps", map[string]int{"a": 1}, map[string]int{"a": 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.a, tt.b); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotEquals(t *testing.T) {
	tests := []struct {
		name string
		a, b interface{}
		want bool
	}{
		{"equal ints", 1, 1, false},
		{"different ints", 1, 2, true},
		{"both nil", nil, nil, false},
		{"one nil", nil, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotEquals(tt.a, tt.b); got != tt.want {
				t.Errorf("NotEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name string
		a, b interface{}
		want int
	}{
		{"both nil", nil, nil, 0},
		{"first nil", nil, 1, -1},
		{"second nil", 1, nil, 1},
		{"int less", 1, 2, -1},
		{"int equal", 2, 2, 0},
		{"int greater", 3, 2, 1},
		{"uint less", uint(1), uint(2), -1},
		{"uint equal", uint(2), uint(2), 0},
		{"uint greater", uint(3), uint(2), 1},
		{"float less", 1.0, 2.0, -1},
		{"float equal", 2.0, 2.0, 0},
		{"float greater", 3.0, 2.0, 1},
		{"string less", "a", "b", -1},
		{"string equal", "a", "a", 0},
		{"string greater", "b", "a", 1},
		{"int8", int8(1), int8(2), -1},
		{"int16", int16(1), int16(2), -1},
		{"int32", int32(1), int32(2), -1},
		{"int64", int64(1), int64(2), -1},
		{"uint8", uint8(1), uint8(2), -1},
		{"uint16", uint16(1), uint16(2), -1},
		{"uint32", uint32(1), uint32(2), -1},
		{"uint64", uint64(1), uint64(2), -1},
		{"float32", float32(1.0), float32(2.0), -1},
		// Different types -> fallback to string comparison
		{"different types (int vs string same repr)", 1, "1", 0},
		{"different types fallback less", 1, "2", -1},   // "1" < "2"
		{"different types fallback greater", 2, "1", 1}, // "2" > "1"
		{"struct types (fallback)", struct{ A int }{1}, struct{ A int }{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.a, tt.b); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeName(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want string
	}{
		{"nil", nil, "nil"},
		{"int", 42, "int"},
		{"string", "hello", "string"},
		{"float64", 3.14, "float64"},
		{"bool", true, "bool"},
		{"slice", []int{1}, "[]int"},
		{"map", map[string]int{}, "map[string]int"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeName(tt.v); got != tt.want {
				t.Errorf("TypeName() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTypeOf(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want reflect.Type
	}{
		{"nil", nil, nil},
		{"int", 42, reflect.TypeOf(0)},
		{"string", "hello", reflect.TypeOf("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TypeOf(tt.v)
			if got != tt.want {
				t.Errorf("TypeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKindOf(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want reflect.Kind
	}{
		{"nil", nil, reflect.Invalid},
		{"int", 42, reflect.Int},
		{"string", "hello", reflect.String},
		{"pointer", func() interface{} { v := 1; return &v }(), reflect.Ptr},
		{"slice", []int{}, reflect.Slice},
		{"map", map[string]int{}, reflect.Map},
		{"bool", true, reflect.Bool},
		{"float64", 3.14, reflect.Float64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KindOf(tt.v); got != tt.want {
				t.Errorf("KindOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsType(t *testing.T) {
	tests := []struct {
		name     string
		v        interface{}
		typeName string
		want     bool
	}{
		{"int matches int", 42, "int", true},
		{"int doesn't match string", 42, "string", false},
		{"nil matches nil", nil, "nil", true},
		{"string matches string", "hello", "string", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsType(tt.v, tt.typeName); got != tt.want {
				t.Errorf("IsType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKind(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		kind reflect.Kind
		want bool
	}{
		{"int is Int", 42, reflect.Int, true},
		{"int is not String", 42, reflect.String, false},
		{"nil is Invalid", nil, reflect.Invalid, true},
		{"string is String", "hello", reflect.String, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKind(tt.v, tt.kind); got != tt.want {
				t.Errorf("IsKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPtr(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
	}{
		{"nil", nil},
		{"int", 42},
		{"string", "hello"},
		{"float64", 3.14},
		{"bool", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToPtr(tt.v)
			if tt.v == nil {
				if got != nil {
					t.Errorf("ToPtr(nil) = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Errorf("ToPtr(%v) = nil, want non-nil", tt.v)
				return
			}
			// Dereference and check
			rv := reflect.ValueOf(got)
			if rv.Kind() != reflect.Ptr {
				t.Errorf("ToPtr(%v) returned non-pointer", tt.v)
				return
			}
			val := rv.Elem().Interface()
			if !reflect.DeepEqual(val, tt.v) {
				t.Errorf("ToPtr(%v) dereferenced = %v", tt.v, val)
			}
		})
	}
}

func TestFromPtr(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want interface{}
	}{
		{"nil", nil, nil},
		{"non-pointer", 42, 42},
		{"pointer to int", func() interface{} { v := 42; return &v }(), 42},
		{"nil pointer", func() interface{} { var p *int; return p }(), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromPtr(tt.v)
			if got != tt.want {
				t.Errorf("FromPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
	}{
		{"nil", nil},
		{"int", 42},
		{"string", "hello"},
		{"slice", []int{1, 2, 3}},
		{"map", map[string]int{"a": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Copy(tt.v)
			if tt.v == nil {
				if got != nil {
					t.Errorf("Copy(nil) = %v, want nil", got)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.v) {
				t.Errorf("Copy(%v) = %v", tt.v, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want string
	}{
		{"nil", nil, ""},
		{"int", 42, "42"},
		{"string", "hello", "hello"},
		{"bool", true, "true"},
		{"float64", 3.14, "3.14"},
		{"slice", []int{1, 2}, "[1 2]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.v); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestClone(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if Clone(nil) != nil {
			t.Error("Clone(nil) should be nil")
		}
	})

	t.Run("primitive", func(t *testing.T) {
		v := 42
		result := Clone(v)
		if result != 42 {
			t.Errorf("Clone(42) = %v, want 42", result)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		v := 42
		p := &v
		cloned := Clone(p).(*int)
		if *cloned != 42 {
			t.Errorf("Clone(ptr) = %v, want 42", *cloned)
		}
		*cloned = 100
		if *p != 42 {
			t.Error("Clone should create independent copy")
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var p *int
		got := Clone(p)
		if got != nil {
			t.Errorf("Clone(nil ptr) = %v, want nil", got)
		}
	})

	t.Run("string", func(t *testing.T) {
		got := Clone("hello")
		if got != "hello" {
			t.Errorf("Clone(\"hello\") = %v", got)
		}
	})

	t.Run("struct", func(t *testing.T) {
		type S struct{ A int }
		got := Clone(S{A: 1})
		if !reflect.DeepEqual(got, S{A: 1}) {
			t.Errorf("Clone(struct) = %v", got)
		}
	})
}
