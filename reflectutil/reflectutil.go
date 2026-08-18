package reflectutil

import (
	"errors"
	"fmt"
	"reflect"
)

// FieldInfo describes an exported struct field.
type FieldInfo struct {
	Name string
	Type reflect.Type
	Tag  string
}

// structValue resolves obj (value or pointer) to a struct reflect.Value.
func structValue(obj any) (reflect.Value, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return reflect.Value{}, errors.New("reflectutil: nil pointer")
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.Value{}, fmt.Errorf("reflectutil: %T is not a struct", obj)
	}
	return v, nil
}

// structType resolves obj's underlying struct type (dereferencing pointers).
func structType(obj any) reflect.Type {
	t := reflect.TypeOf(obj)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// GetFieldValue returns the value of the named exported field of obj
// (a struct value or pointer). It errors on missing or unexported fields.
func GetFieldValue(obj any, name string) (any, error) {
	sv, err := structValue(obj)
	if err != nil {
		return nil, err
	}
	f := sv.FieldByName(name)
	if !f.IsValid() {
		return nil, fmt.Errorf("reflectutil: field %q not found", name)
	}
	if !f.CanInterface() {
		return nil, fmt.Errorf("reflectutil: field %q is not exported", name)
	}
	return f.Interface(), nil
}

// SetFieldValue assigns value to the named field of obj.
// obj must be a pointer to a struct so the field is addressable.
// The value is assigned directly when assignable, otherwise converted.
func SetFieldValue(obj any, name string, value any) error {
	sv, err := structValue(obj)
	if err != nil {
		return err
	}
	if !sv.CanSet() {
		return errors.New("reflectutil: target is not settable (pass a pointer to a struct)")
	}
	f := sv.FieldByName(name)
	if !f.IsValid() {
		return fmt.Errorf("reflectutil: field %q not found", name)
	}
	if !f.CanSet() {
		return fmt.Errorf("reflectutil: field %q is not exported", name)
	}
	if value == nil {
		// only nilable field kinds can accept nil
		switch f.Kind() {
		case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface, reflect.Func, reflect.Chan:
			f.Set(reflect.Zero(f.Type()))
			return nil
		default:
			return fmt.Errorf("reflectutil: cannot assign nil to %s field %q", f.Kind(), name)
		}
	}
	val := reflect.ValueOf(value)
	switch {
	case val.Type().AssignableTo(f.Type()):
		f.Set(val)
	case val.Type().ConvertibleTo(f.Type()):
		f.Set(val.Convert(f.Type()))
	default:
		return fmt.Errorf("reflectutil: cannot assign %T to field %q of type %s", value, name, f.Type())
	}
	return nil
}

// FieldNames returns the exported field names of obj's type, in order.
func FieldNames(obj any) []string {
	t := structType(obj)
	names := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if sf.PkgPath != "" { // unexported
			continue
		}
		names = append(names, sf.Name)
	}
	return names
}

// FieldsOf returns metadata (name, type, tag) for each exported field.
func FieldsOf(obj any) []FieldInfo {
	t := structType(obj)
	fields := make([]FieldInfo, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if sf.PkgPath != "" {
			continue
		}
		fields = append(fields, FieldInfo{
			Name: sf.Name,
			Type: sf.Type,
			Tag:  string(sf.Tag),
		})
	}
	return fields
}

// HasField reports whether obj's type declares an exported field named name.
func HasField(obj any, name string) bool {
	sf, ok := structType(obj).FieldByName(name)
	return ok && sf.PkgPath == ""
}

// GetTag returns the struct tag value of the named field for tagKey.
// It returns "" when the field or the tag key is absent.
func GetTag(obj any, fieldName, tagKey string) string {
	sf, ok := structType(obj).FieldByName(fieldName)
	if !ok {
		return ""
	}
	return sf.Tag.Get(tagKey)
}

// InvokeMethod calls the named method of obj with the given arguments.
// When obj is a struct value, an addressable copy is created so methods
// with pointer receivers can be invoked. It returns the method results.
func InvokeMethod(obj any, name string, args ...any) ([]any, error) {
	sv := reflect.ValueOf(obj)
	m := sv.MethodByName(name)
	if !m.IsValid() && sv.Kind() == reflect.Struct {
		p := reflect.New(sv.Type()) // addressable copy for pointer receivers
		p.Elem().Set(sv)
		m = p.MethodByName(name)
	}
	if !m.IsValid() {
		return nil, fmt.Errorf("reflectutil: method %q not found on %T", name, obj)
	}
	in := make([]reflect.Value, len(args))
	for i, a := range args {
		in[i] = reflect.ValueOf(a)
	}
	out := m.Call(in)
	results := make([]any, len(out))
	for i, r := range out {
		results[i] = r.Interface()
	}
	return results, nil
}

// IsSimpleType reports whether v is a basic scalar type
// (bool, numbers, string) rather than a composite/ref type.
func IsSimpleType(v any) bool {
	if v == nil {
		return false
	}
	switch reflect.ValueOf(v).Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String:
		return true
	default:
		return false
	}
}

// Instantiate allocates a usable instance of typ:
//
//   - struct/ptr -> pointer to a zero struct
//   - slice/map  -> empty (non-nil) collection
//   - scalars    -> zero value
//
// It errors for kinds that cannot be allocated as values (chan, func, ...).
func Instantiate(typ reflect.Type) (any, error) {
	if typ == nil {
		return nil, errors.New("reflectutil: nil type")
	}
	switch typ.Kind() {
	case reflect.Ptr:
		return reflect.New(typ.Elem()).Interface(), nil
	case reflect.Struct:
		return reflect.New(typ).Interface(), nil
	case reflect.Slice:
		return reflect.MakeSlice(typ, 0, 0).Interface(), nil
	case reflect.Map:
		return reflect.MakeMap(typ).Interface(), nil
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return nil, fmt.Errorf("reflectutil: cannot instantiate kind %v", typ.Kind())
	default:
		return reflect.Zero(typ).Interface(), nil
	}
}

// ZeroValue returns the zero value of typ as an any.
func ZeroValue(typ reflect.Type) any {
	return reflect.Zero(typ).Interface()
}