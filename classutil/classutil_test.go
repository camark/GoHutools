package classutil

import (
	"reflect"
	"strings"
	"testing"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string { return "woof" }
func (d *Dog) Bark()        {}

type MyMap map[string]int

func TestSimpleName(t *testing.T) {
	if got := SimpleName(reflect.TypeOf(Dog{})); got != "Dog" {
		t.Errorf("SimpleName(Dog) = %q", got)
	}
	if got := SimpleName(reflect.TypeOf(MyMap{})); got != "MyMap" {
		t.Errorf("SimpleName(MyMap) = %q", got)
	}
	// unnamed types fall back to their string form
	if got := SimpleName(reflect.TypeOf(map[string]int{})); got != "map[string]int" {
		t.Errorf("SimpleName(map) = %q", got)
	}
	if got := SimpleName(reflect.TypeOf([]int{})); got != "[]int" {
		t.Errorf("SimpleName(slice) = %q", got)
	}
}

func TestFullName(t *testing.T) {
	ft := FullName(reflect.TypeOf(Dog{}))
	if !strings.HasSuffix(ft, "/classutil.Dog") {
		t.Errorf("FullName(Dog) = %q", ft)
	}
	if got := FullName(reflect.TypeOf(map[string]int{})); got != "map[string]int" {
		t.Errorf("FullName(unnamed) = %q", got)
	}
}

func TestPackageName(t *testing.T) {
	pkg := PackageName(reflect.TypeOf(Dog{}))
	if !strings.HasSuffix(pkg, "/classutil") {
		t.Errorf("PackageName(Dog) = %q", pkg)
	}
	// unnamed types have no package
	if got := PackageName(reflect.TypeOf(map[string]int{})); got != "" {
		t.Errorf("PackageName(map) = %q", got)
	}
}

func TestIsStruct(t *testing.T) {
	if !IsStruct(reflect.TypeOf(Dog{})) {
		t.Error("IsStruct(Dog)")
	}
	if IsStruct(reflect.TypeOf(MyMap{})) {
		t.Error("IsStruct(MyMap)")
	}
}

func TestCollectionKinds(t *testing.T) {
	slice := reflect.TypeOf([]int{})
	m := reflect.TypeOf(map[string]int{})
	arr := reflect.TypeOf([3]int{})
	if !IsCollection(slice) || !IsCollection(m) || !IsCollection(arr) {
		t.Error("IsCollection for slice/map/array")
	}
	if IsCollection(reflect.TypeOf(Dog{})) {
		t.Error("IsCollection(struct)")
	}
	if !IsSlice(slice) || IsSlice(m) {
		t.Error("IsSlice")
	}
	if !IsMap(m) || IsMap(slice) {
		t.Error("IsMap")
	}
}

func TestPointerInterface(t *testing.T) {
	ptr := reflect.TypeOf(&Dog{})
	if !IsPointer(ptr) {
		t.Error("IsPointer(*Dog)")
	}
	if IsPointer(reflect.TypeOf(Dog{})) {
		t.Error("IsPointer(Dog)")
	}
	if !IsInterface(reflect.TypeOf((*Animal)(nil)).Elem()) {
		t.Error("IsInterface(Animal)")
	}
	if IsInterface(reflect.TypeOf(Dog{})) {
		t.Error("IsInterface(Dog)")
	}
}

func TestMethodNames(t *testing.T) {
	ptr := reflect.TypeOf(&Dog{})
	methods := MethodNames(ptr)
	joined := strings.Join(methods, ",")
	if !strings.Contains(joined, "Speak") || !strings.Contains(joined, "Bark") {
		t.Errorf("MethodNames(*Dog) = %v", methods)
	}
	// value type sees only value-receiver methods
	if got := MethodNames(reflect.TypeOf(Dog{})); len(got) != 1 || got[0] != "Speak" {
		t.Errorf("MethodNames(Dog) = %v", got)
	}
}

func TestHasMethod(t *testing.T) {
	if !HasMethod(reflect.TypeOf(&Dog{}), "Bark") {
		t.Error("HasMethod(+Bark)")
	}
	if HasMethod(reflect.TypeOf(Dog{}), "Bark") {
		t.Error("HasMethod(value, Bark) should be false (pointer receiver)")
	}
}

func TestImplements(t *testing.T) {
	iface := reflect.TypeOf((*Animal)(nil)).Elem()
	if !Implements(reflect.TypeOf(Dog{}), iface) {
		t.Error("Implements(Dog, Animal)")
	}
	if !Implements(reflect.TypeOf(&Dog{}), iface) {
		t.Error("Implements(*Dog, Animal) — value-receiver method promotes")
	}
	if !Implements(iface, iface) {
		t.Error("Implements(self interface)")
	}
	if Implements(reflect.TypeOf("string"), iface) {
		t.Error("string should not implement Animal")
	}
}

func TestIsAssignableFrom(t *testing.T) {
	strT := reflect.TypeOf("")
	if !IsAssignableFrom(strT, strT) {
		t.Error("string assignable from string")
	}
	type MyStr string
	// string and MyStr are both named types — neither is assignable to the other
	if IsAssignableFrom(strT, reflect.TypeOf(MyStr(""))) {
		t.Error("string not assignable from MyStr")
	}
	if IsAssignableFrom(reflect.TypeOf(MyStr("")), strT) {
		t.Error("MyStr not assignable from string (both named)")
	}
	// unnamed composite []int IS assignable to named MySlice
	type MySlice []int
	if !IsAssignableFrom(reflect.TypeOf(MySlice{}), reflect.TypeOf([]int{})) {
		t.Error("MySlice assignable from []int (unnamed underlying)")
	}
	// pointers to unrelated types
	if IsAssignableFrom(reflect.TypeOf(&Dog{}), reflect.TypeOf(&struct{}{})) {
		t.Error("*Dog not assignable from unnamed struct ptr")
	}
}

func TestIsFunc(t *testing.T) {
	if !IsFunc(reflect.TypeOf(func() {})) {
		t.Error("IsFunc(func)")
	}
	if IsFunc(reflect.TypeOf(Dog{})) {
		t.Error("IsFunc(Dog)")
	}
}

func TestNameOfType(t *testing.T) {
	// convenience alias coverage
	if got := NameOfType(Dog{}); got != "Dog" {
		t.Errorf("NameOfType(Dog{}) = %q", got)
	}
}