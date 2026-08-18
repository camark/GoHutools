package opt

import (
	"errors"
	"strings"
	"testing"
)

func TestSomeAndNone(t *testing.T) {
	o := Some(42)
	if !o.IsPresent() {
		t.Error("Some should be present")
	}
	if o.IsEmpty() {
		t.Error("Some should not be empty")
	}
	if o.Get() != 42 {
		t.Errorf("Get = %v", o.Get())
	}

	n := None[int]()
	if n.IsPresent() {
		t.Error("None should not be present")
	}
	if !n.IsEmpty() {
		t.Error("None should be empty")
	}
}

func TestOfNullable(t *testing.T) {
	if OfNullable((*int)(nil)).IsPresent() {
		t.Error("OfNullable(nil) should be empty")
	}
	p := new(int)
	*p = 5
	if !OfNullable(p).IsPresent() {
		t.Error("OfNullable(ptr) should be present")
	}
}

func TestOrElse(t *testing.T) {
	if Some(1).OrElse(99) != 1 {
		t.Error("OrElse should return contained value when present")
	}
	if None[int]().OrElse(99) != 99 {
		t.Error("OrElse should return default when empty")
	}
}

func TestOrElseGet(t *testing.T) {
	called := false
	got := None[int]().OrElseGet(func() int {
		called = true
		return 7
	})
	if !called || got != 7 {
		t.Errorf("OrElseGet = %d (called=%v)", got, called)
	}
}

func TestMap(t *testing.T) {
	mapped := Map(Some(21), func(v int) int { return v * 2 })
	if mapped.Get() != 42 {
		t.Errorf("Map = %v", mapped.Get())
	}
	if Map(None[int](), func(v int) int { return v }).IsPresent() {
		t.Error("Map on None should stay None")
	}
}

func TestFilter(t *testing.T) {
	o := Filter(Some(10), func(v int) bool { return v > 5 })
	if o.IsEmpty() {
		t.Error("Filter(passing) should stay present")
	}
	o2 := Filter(Some(10), func(v int) bool { return v > 20 })
	if o2.IsPresent() {
		t.Error("Filter(failing) should become empty")
	}
}

func TestGetOrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Get on None should panic")
		}
	}()
	_ = None[int]().Get()
}

func TestGetOrElse(t *testing.T) {
	if Some("x").GetOrElse("y") != "x" {
		t.Error("GetOrElse on Some")
	}
	if None[string]().GetOrElse("y") != "y" {
		t.Error("GetOrElse on None")
	}
}

func TestToError(t *testing.T) {
	msg := errors.New("boom")
	err := Some("ok").ToError(msg)
	if err != nil {
		t.Errorf("Some().ToError should be nil, got %v", err)
	}
	err = None[string]().ToError(msg)
	if err == nil || err.Error() != "boom" {
		t.Errorf("None().ToError = %v, want boom", err)
	}
}

func TestOrPanic(t *testing.T) {
	if Some(1).OrPanic() != 1 {
		t.Error("OrPanic on Some should return value")
	}
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("OrPanic on None should panic")
			}
		}()
		_ = None[int]().OrPanic()
	}()
}

func TestIfPresent(t *testing.T) {
	var got int
	Some(5).IfPresent(func(v int) { got = v })
	if got != 5 {
		t.Errorf("IfPresent callback not invoked: %d", got)
	}
	None[int]().IfPresent(func(v int) { got = 100 })
	if got != 5 {
		t.Errorf("IfPresent should not run on None (got=%d)", got)
	}
}

func TestIs(t *testing.T) {
	if !Is(Some("abc"), func(v string) bool { return len(v) == 3 }) {
		t.Error("Is(matcher) on Some")
	}
	if Is(Some("abcd"), func(v string) bool { return len(v) == 3 }) {
		t.Error("Is(mismatch) should be false")
	}
	if Is(None[string](), func(v string) bool { return true }) {
		t.Error("Is on None should be false")
	}
}

func TestString(t *testing.T) {
	if !strings.HasPrefix(Some(1).String(), "Opt[") {
		t.Errorf("Some String = %q", Some(1).String())
	}
	if !strings.Contains(None[int]().String(), "empty") {
		t.Errorf("None String = %q", None[int]().String())
	}
}