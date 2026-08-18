package beanutil

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// CopyOptions controls how CopyProperties copies fields.
type CopyOptions struct {
	// IgnoreFields lists field names (exact or ignore-case) that are not copied.
	IgnoreFields []string
	// SkipNull skips zero-valued source fields (empty string, 0, nil slice/map/pointer).
	SkipNull bool
	// IgnoreCase makes field-name matching case-insensitive.
	IgnoreCase bool
}

// MapOptions controls how MapToBean fills a bean.
type MapOptions struct {
	// IgnoreCase makes map-key matching case-insensitive.
	IgnoreCase bool
}

// containsField reports whether name is in fields (exact if !ignoreCase, else case-insensitive).
func containsField(fields []string, name string, ignoreCase bool) bool {
	for _, f := range fields {
		if ignoreCase {
			if strings.EqualFold(f, name) {
				return true
			}
		} else if f == name {
			return true
		}
	}
	return false
}

// indirect unwraps pointers until it reaches a non-pointer value.
// It returns false when the intermediate pointer chain is broken by a nil pointer.
func indirect(rv reflect.Value) (reflect.Value, bool) {
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return rv, false
		}
		rv = rv.Elem()
	}
	return rv, true
}

// isZeroValue reports whether v holds the zero value of its type.
func isZeroValue(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
		return rv.IsNil()
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.String:
		return rv.String() == ""
	default:
		// structs, arrays, time.Time etc: use reflect.Value.IsZero
		return rv.IsZero()
	}
}

// CopyProperties copies property values from src to dst.
// Both must be pointers to structs (or structs — the value is taken by address when possible).
// Fields are matched by exact name unless CopyOptions.IgnoreCase is set.
func CopyProperties(dst, src interface{}, opts ...CopyOptions) error {
	if dst == nil {
		return fmt.Errorf("beanutil: dst is nil")
	}
	if src == nil {
		return fmt.Errorf("beanutil: src is nil")
	}

	dv := reflect.ValueOf(dst)
	sv := reflect.ValueOf(src)

	// dst must be a pointer to a struct (or addressable struct value)
	dv, ok := indirect(dv)
	if !ok || dv.Kind() != reflect.Struct {
		return fmt.Errorf("beanutil: dst must be a pointer to struct, got %T", dst)
	}
	sv, ok2 := indirect(sv)
	if !ok2 || sv.Kind() != reflect.Struct {
		return fmt.Errorf("beanutil: src must be a pointer to struct, got %T", src)
	}

	var o CopyOptions
	if len(opts) > 0 {
		o = opts[0]
	}

	dt := dv.Type()
	for i := 0; i < dt.NumField(); i++ {
		df := dt.Field(i)
		if df.PkgPath != "" {
			continue // unexported
		}
		if containsField(o.IgnoreFields, df.Name, o.IgnoreCase) {
			continue
		}

		// find matching source field
		var sFieldValue reflect.Value
		found := false
		st := sv.Type()
		for j := 0; j < st.NumField(); j++ {
			cand := st.Field(j)
			if cand.PkgPath != "" {
				continue
			}
			if o.IgnoreCase {
				if !strings.EqualFold(cand.Name, df.Name) {
					continue
				}
			} else if cand.Name != df.Name {
				continue
			}
			sFieldValue = sv.Field(j)
			found = true
			break
		}
		if !found {
			continue
		}

		if o.SkipNull && isZeroValue(sFieldValue) {
			continue
		}

		// dereference source pointer field
		srcVal := sFieldValue
		if srcVal.Kind() == reflect.Ptr {
			if srcVal.IsNil() {
				// nil pointer: only meaningful when dst field is also a pointer set to nil
				dstF := dv.Field(i)
				if dstF.Kind() == reflect.Ptr {
					dstF.Set(reflect.Zero(dstF.Type()))
				}
				continue
			}
			srcVal = srcVal.Elem()
		}

		dstF := dv.Field(i)
		if !dstF.CanSet() {
			continue
		}

		// dereference dst pointer field
		if dstF.Kind() == reflect.Ptr {
			if srcVal.Kind() != dstF.Type().Elem().Kind() {
				// mismatch (e.g. src value vs dst *value) — try conversion
				if !convertAssign(dstF, srcVal) {
					continue
				}
			} else {
				if dstF.IsNil() {
					dstF.Set(reflect.New(dstF.Type().Elem()))
				}
				if err := assign(dstF.Elem(), srcVal); err != nil {
					continue
				}
			}
			continue
		}

		if err := assign(dstF, srcVal); err != nil {
			continue
		}
	}
	return nil
}

// assign sets dst to the value of src, converting between kinds when needed.
func assign(dst, src reflect.Value) error {
	if !dst.CanSet() {
		return fmt.Errorf("cannot set field")
	}
	if src.Type().AssignableTo(dst.Type()) {
		dst.Set(src)
		return nil
	}
	if src.Type().ConvertibleTo(dst.Type()) {
		dst.Set(src.Convert(dst.Type()))
		return nil
	}
	return fmt.Errorf("type mismatch: %v -> %v", src.Type(), dst.Type())
}

// convertAssign sets a pointer field dst from a value src, allocating as needed
// and converting kinds (e.g. string -> int64).
func convertAssign(dst reflect.Value, src reflect.Value) bool {
	dstType := dst.Type().Elem()
	var newVal reflect.Value
	switch dstType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if src.Kind() == reflect.String {
			i, err := strconv.ParseInt(src.String(), 10, 64)
			if err != nil {
				return false
			}
			newVal = reflect.ValueOf(i).Convert(dstType)
		} else if isIntKind(src.Kind()) {
			newVal = src.Convert(dstType)
		} else {
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if src.Kind() == reflect.String {
			u, err := strconv.ParseUint(src.String(), 10, 64)
			if err != nil {
				return false
			}
			newVal = reflect.ValueOf(u).Convert(dstType)
		} else if isUintKind(src.Kind()) {
			newVal = src.Convert(dstType)
		} else {
			return false
		}
	case reflect.Float32, reflect.Float64:
		if src.Kind() == reflect.String {
			f, err := strconv.ParseFloat(src.String(), 64)
			if err != nil {
				return false
			}
			newVal = reflect.ValueOf(f).Convert(dstType)
		} else if isFloatKind(src.Kind()) {
			newVal = src.Convert(dstType)
		} else {
			return false
		}
	case reflect.Bool:
		if src.Kind() == reflect.String {
			b, err := strconv.ParseBool(src.String())
			if err != nil {
				return false
			}
			newVal = reflect.ValueOf(b)
		} else if src.Kind() == reflect.Bool {
			newVal = src
		} else {
			return false
		}
	case reflect.String:
		if isIntKind(src.Kind()) || isUintKind(src.Kind()) || isFloatKind(src.Kind()) {
			newVal = reflect.ValueOf(valueToString(src))
		} else if src.Kind() == reflect.String {
			newVal = src
		} else {
			return false
		}
	default:
		return false
	}

	if dst.IsNil() {
		dst.Set(reflect.New(dstType))
	}
	if ok := dst.Elem().CanSet(); ok {
		dst.Elem().Set(newVal)
		return true
	}
	return false
}

func isIntKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func isUintKind(k reflect.Kind) bool {
	switch k {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	}
	return false
}

func isFloatKind(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
}

func valueToString(rv reflect.Value) string {
	switch rv.Kind() {
	case reflect.String:
		return rv.String()
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'f', -1, 64)
	}
	return fmt.Sprintf("%v", rv.Interface())
}

// BeanToMap converts a struct to a map keyed by exported field names.
func BeanToMap(bean interface{}) (map[string]interface{}, error) {
	if bean == nil {
		return nil, fmt.Errorf("beanutil: bean is nil")
	}
	rv := reflect.ValueOf(bean)
	var ok bool
	if rv, ok = indirect(rv); !ok {
		return nil, fmt.Errorf("beanutil: bean is nil")
	}
	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("beanutil: bean must be a struct or pointer to struct, got %T", bean)
	}

	m := make(map[string]interface{}, rv.NumField())
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if f.PkgPath != "" {
			continue // unexported
		}
		m[f.Name] = rv.Field(i).Interface()
	}
	return m, nil
}

// MapToBean fills the struct pointed to by bean from map m.
// bean may be *Struct or **Struct. When ignoreError is true, type-conversion
// failures per field are skipped instead of reported.
func MapToBean(bean interface{}, m map[string]interface{}, ignoreError bool, opts ...MapOptions) error {
	if bean == nil {
		return fmt.Errorf("beanutil: bean is nil")
	}
	if m == nil {
		return fmt.Errorf("beanutil: map is nil")
	}
	var o MapOptions
	if len(opts) > 0 {
		o = opts[0]
	}

	rv := reflect.ValueOf(bean)
	// unwrap **Struct -> *Struct
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			if rv.Type().Elem().Kind() == reflect.Struct {
				rv.Set(reflect.New(rv.Type().Elem()))
			} else {
				return fmt.Errorf("beanutil: bean must point to a struct")
			}
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("beanutil: bean must be a pointer to struct, got %T", bean)
	}

	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if f.PkgPath != "" {
			continue
		}
		key := f.Name
		val, exists := m[key]
		if !exists && o.IgnoreCase {
			for mk, mv := range m {
				if strings.EqualFold(mk, key) {
					key = mk
					val = mv
					exists = true
					break
				}
			}
		}
		if !exists {
			continue
		}

		field := rv.Field(i)
		if err := setField(field, val); err != nil {
			if ignoreError {
				continue
			}
			return fmt.Errorf("beanutil: set field %s: %w", f.Name, err)
		}
	}
	return nil
}

// setField sets a struct field from an interface{} value, converting types as needed.
func setField(field reflect.Value, val interface{}) error {
	if !field.CanSet() {
		return fmt.Errorf("field is not settable")
	}
	if val == nil {
		// nil clears pointer/map/slice/interface fields
		switch field.Kind() {
		case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
			field.Set(reflect.Zero(field.Type()))
			return nil
		default:
			return fmt.Errorf("nil cannot be assigned to %v", field.Type())
		}
	}
	rv := reflect.ValueOf(val)

	// direct type match
	if rv.Type().AssignableTo(field.Type()) {
		field.Set(rv)
		return nil
	}
	if rv.Type().ConvertibleTo(field.Type()) {
		field.Set(rv.Convert(field.Type()))
		return nil
	}

	// dst is a pointer: allocate and set
	if field.Kind() == reflect.Ptr {
		elemType := field.Type().Elem()
		if rv.Type().AssignableTo(elemType) {
			np := reflect.New(elemType)
			np.Elem().Set(rv.Convert(elemType))
			field.Set(np)
			return nil
		}
		if convertAssign(field, rv) {
			return nil
		}
		return fmt.Errorf("%v cannot be assigned to %v", rv.Type(), field.Type())
	}

	// numeric/string conversions
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var i int64
		switch rv.Kind() {
		case reflect.String:
			var err error
			i, err = strconv.ParseInt(rv.String(), 10, 64)
			if err != nil {
				return err
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i = rv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			i = int64(rv.Uint())
		case reflect.Float32, reflect.Float64:
			i = int64(rv.Float())
		default:
			return fmt.Errorf("%v cannot be converted to int", rv.Type())
		}
		field.SetInt(i)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		var u uint64
		switch rv.Kind() {
		case reflect.String:
			var err error
			u, err = strconv.ParseUint(rv.String(), 10, 64)
			if err != nil {
				return err
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			u = uint64(rv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			u = rv.Uint()
		default:
			return fmt.Errorf("%v cannot be converted to uint", rv.Type())
		}
		field.SetUint(u)
		return nil
	case reflect.Float32, reflect.Float64:
		var f float64
		switch rv.Kind() {
		case reflect.String:
			var err error
			f, err = strconv.ParseFloat(rv.String(), 64)
			if err != nil {
				return err
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			f = float64(rv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			f = float64(rv.Uint())
		case reflect.Float32, reflect.Float64:
			f = rv.Float()
		default:
			return fmt.Errorf("%v cannot be converted to float", rv.Type())
		}
		field.SetFloat(f)
		return nil
	case reflect.Bool:
		switch rv.Kind() {
		case reflect.Bool:
			field.SetBool(rv.Bool())
			return nil
		case reflect.String:
			b, err := strconv.ParseBool(rv.String())
			if err != nil {
				return err
			}
			field.SetBool(b)
			return nil
		}
	case reflect.String:
		if isIntKind(rv.Kind()) || isUintKind(rv.Kind()) || isFloatKind(rv.Kind()) || rv.Kind() == reflect.Bool {
			field.SetString(valueToString(rv))
			return nil
		}
	}

	return fmt.Errorf("%v cannot be assigned to %v", rv.Type(), field.Type())
}