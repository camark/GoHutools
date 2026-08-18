package tuple

import (
	"fmt"
	"reflect"
	"strings"
)

// Pair is an immutable two-element container (a.k.a. a 2-tuple).
type Pair[L, R any] struct {
	left  L
	right R
}

// Of builds a Pair with automatic type inference: Of("a", 1) -> Pair[string, int].
func Of[L, R any](left L, right R) Pair[L, R] {
	return Pair[L, R]{left: left, right: right}
}

// First returns the left element.
func (p Pair[L, R]) First() L { return p.left }

// Second returns the right element.
func (p Pair[L, R]) Second() R { return p.right }

// Left is an alias of First.
func (p Pair[L, R]) Left() L { return p.left }

// Right is an alias of Second.
func (p Pair[L, R]) Right() R { return p.right }

// Equal reports whether both elements match.
func (p Pair[L, R]) Equal(o Pair[L, R]) bool {
	return reflect.DeepEqual(p.left, o.left) && reflect.DeepEqual(p.right, o.right)
}

// Swap returns a pair with the elements exchanged.
func (p Pair[L, R]) Swap() Pair[R, L] {
	return Pair[R, L]{left: p.right, right: p.left}
}

// String renders the pair as "(left, right)".
func (p Pair[L, R]) String() string {
	return fmt.Sprintf("(%v, %v)", p.left, p.right)
}

// Triple is an immutable three-element container.
type Triple[L, M, R any] struct {
	a L
	b M
	c R
}

// Of3 builds a Triple with automatic type inference.
func Of3[L, M, R any](a L, b M, c R) Triple[L, M, R] {
	return Triple[L, M, R]{a: a, b: b, c: c}
}

// First returns the first element.
func (t Triple[L, M, R]) First() L { return t.a }

// Second returns the second element.
func (t Triple[L, M, R]) Second() M { return t.b }

// Third returns the third element.
func (t Triple[L, M, R]) Third() R { return t.c }

// A/B/C are short aliases for First/Second/Third.
func (t Triple[L, M, R]) A() L { return t.a }
func (t Triple[L, M, R]) B() M { return t.b }
func (t Triple[L, M, R]) C() R { return t.c }

// String renders the triple as "(a, b, c)".
func (t Triple[L, M, R]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", t.a, t.b, t.c)
}

// Tuple is a heterogeneous, ordered container (a variadic tuple).
type Tuple struct {
	data []any
}

// NewTuple builds a tuple from the given elements (any types allowed).
func NewTuple(elems ...any) Tuple {
	return Tuple{data: elems}
}

// Get returns the i-th element.
func (t Tuple) Get(i int) any { return t.data[i] }

// Size returns the number of elements.
func (t Tuple) Size() int { return len(t.data) }

// Data exposes the underlying slice (read-only by convention).
func (t Tuple) Data() []any { return t.data }

// Contains reports whether v appears in the tuple (using reflect.DeepEqual).
func (t Tuple) Contains(v any) bool {
	for _, e := range t.data {
		if reflect.DeepEqual(e, v) {
			return true
		}
	}
	return false
}

// Join concatenates t and o into a new tuple.
func (t Tuple) Join(o Tuple) Tuple {
	merged := make([]any, 0, len(t.data)+len(o.data))
	merged = append(merged, t.data...)
	merged = append(merged, o.data...)
	return Tuple{data: merged}
}

// String renders the tuple as "(e1, e2, ...)".
func (t Tuple) String() string {
	parts := make([]string, len(t.data))
	for i, e := range t.data {
		parts[i] = fmt.Sprintf("%v", e)
	}
	return "(" + strings.Join(parts, ", ") + ")"
}