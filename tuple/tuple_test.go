package tuple

import (
	"reflect"
	"testing"
)

func TestPairOf(t *testing.T) {
	p := Of(1, "one")
	if p.First() != 1 || p.Second() != "one" {
		t.Errorf("Of(1,one) = (%v, %v)", p.First(), p.Second())
	}
	// type inference without explicit params
	if p.Left() != 1 || p.Right() != "one" {
		t.Errorf("Left/Right mismatch")
	}
	// with structs/pointers
	type user struct{ name string }
	u := &user{name: "a"}
	p2 := Of(u, 42)
	if p2.First() != u {
		t.Errorf("pointer pair mismatch")
	}
}

func TestPairEqual(t *testing.T) {
	a := Of(1, "x")
	b := Of(1, "x")
	c := Of(2, "x")
	if !a.Equal(b) {
		t.Error("equal pairs should be Equal")
	}
	if a.Equal(c) {
		t.Error("unequal pairs should not be Equal")
	}
	// reflect.DeepEqual on pair itself
	if !reflect.DeepEqual(a, b) {
		t.Error("DeepEqual on pairs failed")
	}
}

func TestPairSwap(t *testing.T) {
	p := Of(1, "a").Swap()
	if p.First() != "a" || p.Second() != 1 {
		t.Errorf("Swap = (%v, %v)", p.First(), p.Second())
	}
}

func TestTriple(t *testing.T) {
	tr := Of3(1, "mid", true)
	if tr.First() != 1 || tr.Second() != "mid" || tr.Third() != true {
		t.Errorf("Of3 = (%v,%v,%v)", tr.First(), tr.Second(), tr.Third())
	}
	if tr.A() != 1 || tr.B() != "mid" || tr.C() != true {
		t.Errorf("A/B/C mismatch")
	}
}

func TestTuple(t *testing.T) {
	tup := NewTuple(1, "two", 3.0)
	if tup.Size() != 3 {
		t.Errorf("Size = %d", tup.Size())
	}
	if tup.Get(0) != 1 || tup.Get(1) != "two" {
		t.Errorf("Get failed")
	}
	if !tup.Contains(3.0) {
		t.Error("Contains(3.0) should be true")
	}
	if tup.Contains("nope") {
		t.Error("Contains(nope) should be false")
	}
}

func TestTupleJoin(t *testing.T) {
	a := NewTuple(1, 2)
	b := NewTuple(3, 4)
	c := a.Join(b)
	if c.Size() != 4 || c.Get(0) != 1 || c.Get(3) != 4 {
		t.Errorf("Join failed: %v", c.Data())
	}
}

func TestTupleToString(t *testing.T) {
	tup := NewTuple(1, "a")
	s := tup.String()
	if s == "" || s[0] != '(' {
		t.Errorf("String = %q, want (...)", s)
	}
}

func TestTupleEmpty(t *testing.T) {
	tup := NewTuple()
	if tup.Size() != 0 {
		t.Errorf("empty tuple size = %d", tup.Size())
	}
}

func TestPairString(t *testing.T) {
	p := Of("k", "v")
	if p.String() == "" {
		t.Error("Pair String should not be empty")
	}
}

func TestPairNilSafe(t *testing.T) {
	var p *Pair[int, string]
	if p != nil {
		t.Error("nil pair should be nil")
	}
}