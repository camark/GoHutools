package classutil

import "reflect"

// This package ports Hutool's ClassUtil to Go's type system
// (reflect.Type is the analogue of a Java Class object).

// SimpleName returns the short name of a type:
// the declared name for named types, the string form otherwise
// (e.g. "map[string]int", "[]int", "func()").
func SimpleName(t reflect.Type) string {
	if t == nil {
		return "<nil>"
	}
	if name := t.Name(); name != "" {
		return name
	}
	return t.String()
}

// FullName returns the fully qualified name "pkg/import/path.TypeName"
// for named types, or the string form for unnamed types.
func FullName(t reflect.Type) string {
	if t == nil {
		return "<nil>"
	}
	name := t.Name()
	if name == "" {
		return t.String()
	}
	if pkg := t.PkgPath(); pkg != "" {
		return pkg + "." + name
	}
	return name
}

// PackageName returns the import path of the package that declares the type.
// It returns "" for unnamed types and predeclared types (int, string, ...).
func PackageName(t reflect.Type) string {
	if t == nil {
		return ""
	}
	return t.PkgPath()
}

// NameOfType is a convenience wrapper returning SimpleName(TypeOf(x)).
func NameOfType(v any) string {
	return SimpleName(reflect.TypeOf(v))
}

// IsStruct reports whether t is a struct type.
func IsStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

// IsCollection reports whether t is a slice, array or map.
func IsCollection(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return true
	}
	return false
}

// IsSlice reports whether t is a slice.
func IsSlice(t reflect.Type) bool { return t.Kind() == reflect.Slice }

// IsMap reports whether t is a map.
func IsMap(t reflect.Type) bool { return t.Kind() == reflect.Map }

// IsPointer reports whether t is a pointer.
func IsPointer(t reflect.Type) bool { return t.Kind() == reflect.Ptr }

// IsInterface reports whether t is an interface type.
func IsInterface(t reflect.Type) bool { return t.Kind() == reflect.Interface }

// IsFunc reports whether t is a function type.
func IsFunc(t reflect.Type) bool { return t.Kind() == reflect.Func }

// MethodNames returns the method names of t's method set, in order.
// Pointer types expose methods with pointer receivers (and promoted
// value-receiver methods); value types expose value-receiver methods only.
func MethodNames(t reflect.Type) []string {
	n := t.NumMethod()
	names := make([]string, 0, n)
	for i := 0; i < n; i++ {
		names = append(names, t.Method(i).Name)
	}
	return names
}

// HasMethod reports whether t's method set contains a method named name.
func HasMethod(t reflect.Type, name string) bool {
	_, ok := t.MethodByName(name)
	return ok
}

// Implements reports whether t implements the interface iface.
func Implements(t, iface reflect.Type) bool {
	return t != nil && iface != nil && iface.Kind() == reflect.Interface && t.Implements(iface)
}

// IsAssignableFrom reports whether a value of type src can be assigned
// to a variable of type dst (following the Go assignment rules,
// including identical underlying types with at least one unnamed side).
func IsAssignableFrom(dst, src reflect.Type) bool {
	return src.AssignableTo(dst)
}