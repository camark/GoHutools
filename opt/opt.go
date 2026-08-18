package opt

import "fmt"

// Opt is a generic container that may or may not hold a value —
// a Go port of Java's Optional and Hutool's Opt.
type Opt[T any] struct {
	value   T
	present bool
}

// Some wraps a non-null value into an Opt.
func Some[T any](v T) Opt[T] {
	return Opt[T]{value: v, present: true}
}

// None returns an empty Opt of type T.
func None[T any]() Opt[T] {
	return Opt[T]{present: false}
}

// OfNullable wraps v if non-nil; for types that can be nil this mirrors
// Optional.ofNullable. For value types it behaves like Some.
func OfNullable[T any](v *T) Opt[*T] {
	var p *T
	if v == nil {
		return Opt[*T]{value: p, present: false}
	}
	return Opt[*T]{value: v, present: true}
}

// IsPresent reports whether the Opt holds a value.
func (o Opt[T]) IsPresent() bool { return o.present }

// IsEmpty reports whether the Opt holds no value.
func (o Opt[T]) IsEmpty() bool { return !o.present }

// Get returns the contained value; panics when empty.
func (o Opt[T]) Get() T {
	if !o.present {
		panic("opt: Get on empty Opt")
	}
	return o.value
}

// GetOrElse returns the contained value, or the default when empty.
func (o Opt[T]) GetOrElse(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}

// OrElse is an alias of GetOrElse.
func (o Opt[T]) OrElse(defaultValue T) T {
	return o.GetOrElse(defaultValue)
}

// OrElseGet returns the contained value or the result of fn when empty.
func (o Opt[T]) OrElseGet(fn func() T) T {
	if o.present {
		return o.value
	}
	return fn()
}

// OrPanic returns the contained value, panicking when empty.
func (o Opt[T]) OrPanic() T {
	if !o.present {
		panic("opt: value is empty")
	}
	return o.value
}

// ToError returns error when the Opt is empty, nil when present.
func (o Opt[T]) ToError(err error) error {
	if o.present {
		return nil
	}
	return err
}

// IfPresent runs fn with the contained value when present; no-op otherwise.
func (o Opt[T]) IfPresent(fn func(v T)) {
	if o.present {
		fn(o.value)
	}
}

// Map transforms the value with fn when present, returning None otherwise.
func Map[U, V any](o Opt[U], fn func(U) V) Opt[V] {
	if o.present {
		return Some(fn(o.value))
	}
	return None[V]()
}

// Filter narrows the Opt when present: values failing the predicate become empty.
func Filter[U any](o Opt[U], fn func(U) bool) Opt[U] {
	if o.present && fn(o.value) {
		return o
	}
	return None[U]()
}

// Is tests the contained value with matcher; false when empty.
func Is[U any](o Opt[U], matcher func(U) bool) bool {
	if o.present {
		return matcher(o.value)
	}
	return false
}

// String renders the Opt for debugging.
func (o Opt[T]) String() string {
	if o.present {
		return fmt.Sprintf("Opt[present: %v]", o.value)
	}
	return "Opt[empty]"
}