package enumutil

import (
	"reflect"
	"testing"
)

type Color string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

var colors = []Color{Red, Green, Blue}

type Status int

const (
	Active Status = iota + 1
	Inactive
)

var statuses = []Status{Active, Inactive}

func TestFromString(t *testing.T) {
	v, err := FromString(colors, "green")
	if err != nil || v != Green {
		t.Errorf("FromString(green) = %q, err=%v", v, err)
	}
	if _, err := FromString(colors, "purple"); err == nil {
		t.Error("FromString(missing) should error")
	}
	// empty values → error
	if _, err := FromString[Color](nil, "x"); err == nil {
		t.Error("FromString(nil) should error")
	}
}

func TestMustFromString(t *testing.T) {
	if MustFromString(colors, "red") != Red {
		t.Error("MustFromString(red)")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustFromString(missing) should panic")
		}
	}()
	_ = MustFromString(colors, "purple")
}

func TestContainsValue(t *testing.T) {
	if !ContainsValue(colors, Green) {
		t.Error("ContainsValue(Green)")
	}
	if ContainsValue(colors, Color("purple")) {
		t.Error("ContainsValue(unknown)")
	}
	// works for int-backed enums too
	if !ContainsValue(statuses, Active) {
		t.Error("ContainsValue(Active)")
	}
}

func TestNames(t *testing.T) {
	got := Names(colors)
	if !reflect.DeepEqual(got, []string{"red", "green", "blue"}) {
		t.Errorf("Names = %v", got)
	}
}

func TestIndex(t *testing.T) {
	if Index(colors, Green) != 1 {
		t.Errorf("Index(Green) = %d", Index(colors, Green))
	}
	if Index(colors, Color("purple")) != -1 {
		t.Error("Index(unknown) should be -1")
	}
}

func TestToValueMap(t *testing.T) {
	m := ToValueMap(colors)
	if len(m) != 3 {
		t.Fatalf("ToValueMap len = %d", len(m))
	}
	if m["blue"] != Blue {
		t.Errorf("ToValueMap['blue'] = %q", m["blue"])
	}
	if _, ok := m["red"]; !ok {
		t.Error("ToValueMap missing red")
	}
}

func TestOf(t *testing.T) {
	// Of resolves from a string, indicating the source of truth is
	// fine for typical lookup-by-value use.
	defer func() {
		if r := recover(); r == nil {
			t.Error("Of(unknown) should panic")
		}
	}()
	_ = Of(colors, "nope")
}