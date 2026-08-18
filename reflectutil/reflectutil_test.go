package reflectutil

import (
	"reflect"
	"strings"
	"testing"
)

type meta struct {
	City string
}

type person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string
	secret string // unexported
	Tags   []string
	Meta   meta
}

func (p person) Greet(prefix string) (string, error) {
	return prefix + p.Name, nil
}

func (p *person) Rename(newName string) {
	p.Name = newName
}

var testPerson = person{Name: "Alice", Age: 30, Email: "a@x.com",
	secret: "hidden", Tags: []string{"a", "b"}, Meta: meta{City: "NYC"}}

func TestGetFieldValue(t *testing.T) {
	v, err := GetFieldValue(testPerson, "Name")
	if err != nil || v != "Alice" {
		t.Errorf("Name = %v, err=%v", v, err)
	}
	v, err = GetFieldValue(&testPerson, "Age")
	if err != nil || v != 30 {
		t.Errorf("Age = %v, err=%v", v, err)
	}
	v, err = GetFieldValue(testPerson, "Tags")
	if err != nil {
		t.Errorf("Tags err=%v", err)
	} else if !reflect.DeepEqual(v, []string{"a", "b"}) {
		t.Errorf("Tags = %v", v)
	}
}

func TestGetFieldValueErrors(t *testing.T) {
	if _, err := GetFieldValue(testPerson, "Missing"); err == nil {
		t.Error("missing field should error")
	}
	if _, err := GetFieldValue(testPerson, "secret"); err == nil {
		t.Error("unexported field should error")
	}
	if _, err := GetFieldValue("not a struct", "Name"); err == nil {
		t.Error("non-struct should error")
	}
}

func TestSetFieldValue(t *testing.T) {
	p := testPerson
	if err := SetFieldValue(&p, "Name", "Bob"); err != nil {
		t.Fatal(err)
	}
	if p.Name != "Bob" {
		t.Errorf("Name = %q, want Bob", p.Name)
	}
	// int assignment (int vs untyped float from any)
	if err := SetFieldValue(&p, "Age", 31); err != nil {
		t.Fatal(err)
	}
	if p.Age != 31 {
		t.Errorf("Age = %d", p.Age)
	}
}

func TestSetFieldValuePointerReceiver(t *testing.T) {
	// should work on a pointer even for the inner struct field
	p := &testPerson
	if err := SetFieldValue(p, "Email", "b@x.com"); err != nil {
		t.Fatal(err)
	}
	if p.Email != "b@x.com" {
		t.Errorf("Email = %q", p.Email)
	}
}

func TestSetFieldValueErrors(t *testing.T) {
	p := testPerson
	if err := SetFieldValue(&p, "Missing", 1); err == nil {
		t.Error("missing field should error")
	}
	if err := SetFieldValue(&p, "secret", "x"); err == nil {
		t.Error("unexported field should error")
	}
	if err := SetFieldValue(p, "Name", "x"); err == nil {
		t.Error("non-pointer target should error (not settable)")
	}
	if err := SetFieldValue(&p, "Age", "not-a-number"); err == nil {
		t.Error("type mismatch should error")
	}
}

func TestFieldNames(t *testing.T) {
	names := FieldNames(testPerson)
	joined := strings.Join(names, ",")
	for _, want := range []string{"Name", "Age", "Email", "Tags", "Meta"} {
		if !strings.Contains(joined, want) {
			t.Errorf("FieldNames missing %s: %v", want, names)
		}
		if strings.Contains(joined, "secret") {
			t.Errorf("FieldNames should exclude unexported: %v", names)
		}
	}
}

func TestFieldsOf(t *testing.T) {
	fields := FieldsOf(testPerson)
	if len(fields) != 5 {
		t.Fatalf("fields = %d, want 5 (unexported excluded)", len(fields))
	}
	byName := map[string]FieldInfo{}
	for _, f := range fields {
		byName[f.Name] = f
	}
	if byName["Name"].Type != reflect.TypeOf("") {
		t.Errorf("Name type = %v", byName["Name"].Type)
	}
	if byName["Name"].Tag != `json:"name"` {
		t.Errorf("Name tag = %q", byName["Name"].Tag)
	}
}

func TestHasField(t *testing.T) {
	if !HasField(testPerson, "Email") {
		t.Error("HasField(Email)")
	}
	if HasField(testPerson, "secret") {
		t.Error("HasField(secret) should be false (unexported)")
	}
	if HasField(testPerson, "Nope") {
		t.Error("HasField(Nope)")
	}
}

func TestGetTag(t *testing.T) {
	if got := GetTag(testPerson, "Name", "json"); got != "name" {
		t.Errorf("GetTag(Name) = %q", got)
	}
	if got := GetTag(testPerson, "Email", "json"); got != "" {
		t.Errorf("GetTag(Email) = %q, want empty", got)
	}
	if got := GetTag(testPerson, "Name", "xml"); got != "" {
		t.Errorf("GetTag(Name, xml) = %q, want empty", got)
	}
}

func TestInvokeMethod(t *testing.T) {
	results, err := InvokeMethod(testPerson, "Greet", "Hello, ")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 2 {
		t.Fatalf("results = %d", len(results))
	}
	if results[0] != "Hello, Alice" {
		t.Errorf("result[0] = %v", results[0])
	}
	// pointer receiver: addressable copy is used automatically
	if _, err := InvokeMethod(testPerson, "Rename", "Zed"); err != nil {
		t.Errorf("pointer-receiver method on value: %v", err)
	}
	// missing method
	if _, err := InvokeMethod(testPerson, "Nope"); err == nil {
		t.Error("missing method should error")
	}
}

func TestIsSimpleType(t *testing.T) {
	simple := []interface{}{1, 3.14, true, "x", byte(1), rune('a')}
	for _, v := range simple {
		if !IsSimpleType(v) {
			t.Errorf("IsSimpleType(%v) = false", v)
		}
	}
	complexVals := []interface{}{person{}, []int{1}, map[string]int{}, make(chan int)}
	for _, v := range complexVals {
		if IsSimpleType(v) {
			t.Errorf("IsSimpleType(%v) = true", v)
		}
	}
}

func TestInstantiate(t *testing.T) {
	inst, err := Instantiate(reflect.TypeOf(person{}))
	if err != nil {
		t.Fatal(err)
	}
	if _, isPerson := inst.(*person); !isPerson {
		t.Errorf("Instantiate = %T, want *person", inst)
	}
	if _, err := Instantiate(reflect.TypeOf(make(chan int))); err == nil {
		t.Error("Instantiate(chan) should error")
	}
}

func TestDeepZeroValue(t *testing.T) {
	z := ZeroValue(reflect.TypeOf(person{}))
	if _, isPerson := z.(person); !isPerson {
		t.Errorf("ZeroValue = %T", z)
	}
}