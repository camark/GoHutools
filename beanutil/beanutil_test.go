package beanutil

import (
	"strings"
	"testing"
	"time"
)

type source struct {
	Name    string
	Age     int
	Email   string
	Private string
}

type dest struct {
	Name  string
	Age   int
	Email string
	Extra string
}

type lowerSrc struct {
	UserName string
	Nickname string
}

type lowerDst struct {
	username string
	nickname string
	others   string
}

func TestCopyProperties(t *testing.T) {
	src := &source{Name: "Alice", Age: 30, Email: "alice@example.com", Private: "secret"}
	dst := &dest{Extra: "keep-me"}

	if err := CopyProperties(dst, src); err != nil {
		t.Fatalf("CopyProperties error: %v", err)
	}
	if dst.Name != "Alice" || dst.Age != 30 || dst.Email != "alice@example.com" {
		t.Errorf("copied fields wrong: %+v", dst)
	}
	// non-matching field stays untouched
	if dst.Extra != "keep-me" {
		t.Errorf("Extra should be untouched, got %q", dst.Extra)
	}
}

func TestCopyPropertiesOverwrites(t *testing.T) {
	src := &source{Name: "Bob", Age: 40}
	dst := &dest{Name: "Old", Age: 1}

	_ = CopyProperties(dst, src)
	if dst.Name != "Bob" || dst.Age != 40 {
		t.Errorf("expected overwrite, got %+v", dst)
	}
}

func TestCopyPropertiesIgnoresField(t *testing.T) {
	src := &source{Name: "Carol", Age: 50, Email: "c@example.com"}
	dst := &dest{}
	opts := CopyOptions{IgnoreFields: []string{"Age", "Email"}}

	if err := CopyProperties(dst, src, opts); err != nil {
		t.Fatalf("CopyProperties error: %v", err)
	}
	if dst.Name != "Carol" {
		t.Errorf("Name should be copied, got %q", dst.Name)
	}
	if dst.Age != 0 || dst.Email != "" {
		t.Errorf("ignored fields should not be copied: %+v", dst)
	}
}

func TestCopyPropertiesSkipNull(t *testing.T) {
	src := &source{Name: "Dave"} // Email empty, Age 0
	dst := &dest{Name: "Keep", Age: 99, Email: "keep@example.com"}
	opts := CopyOptions{SkipNull: true}

	_ = CopyProperties(dst, src, opts)
	if dst.Name != "Dave" {
		t.Errorf("non-zero Name should be copied, got %q", dst.Name)
	}
	if dst.Age != 99 || dst.Email != "keep@example.com" {
		t.Errorf("zero values should be skipped with SkipNull, got %+v", dst)
	}
}

type camelSrc struct {
	UserName string
	NickName string
}

type camelDst struct {
	Username string
	Nickname string
}

func TestCopyPropertiesIgnoreCase(t *testing.T) {
	src := &camelSrc{UserName: "eve", NickName: "E"}
	dst := &camelDst{}
	opts := CopyOptions{IgnoreCase: true}

	if err := CopyProperties(dst, src, opts); err != nil {
		t.Fatalf("CopyProperties error: %v", err)
	}
	if dst.Username != "eve" || dst.Nickname != "E" {
		t.Errorf("case-insensitive copy failed: %+v", dst)
	}
}

func TestCopyPropertiesStructValue(t *testing.T) {
	src := source{Name: "Frank", Age: 20}
	dst := dest{}
	if err := CopyProperties(&dst, &src); err != nil {
		t.Fatal(err)
	}
	if dst.Name != "Frank" || dst.Age != 20 {
		t.Errorf("struct value copy failed: %+v", dst)
	}
}

func TestCopyPropertiesErrors(t *testing.T) {
	if err := CopyProperties(nil, &source{}); err == nil {
		t.Error("nil dst should error")
	}
	if err := CopyProperties(&dest{}, nil); err == nil {
		t.Error("nil src should error")
	}
	var d dest
	var s source
	if err := CopyProperties(&d, &s); err != nil {
		t.Errorf("struct values should work: %v", err)
	}
	if err := CopyProperties("not-a-pointer", &s); err == nil {
		t.Error("non-pointer dst should error")
	}
}

func TestBeanToMap(t *testing.T) {
	b := &source{Name: "Grace", Age: 25, Email: "g@example.com", Private: "x"}
	m, err := BeanToMap(b)
	if err != nil {
		t.Fatalf("BeanToMap error: %v", err)
	}
	if m["Name"] != "Grace" || m["Age"] != 25 || m["Email"] != "g@example.com" {
		t.Errorf("BeanToMap wrong: %v", m)
	}
	if len(m) != 4 {
		t.Errorf("expected 4 fields, got %d: %v", len(m), m)
	}
}

func TestBeanToMapIgnoresUnexported(t *testing.T) {
	b := &lowerDst{}
	m, err := BeanToMap(b)
	if err != nil {
		t.Fatal(err)
	}
	if len(m) != 0 {
		t.Errorf("unexported fields should be skipped, got %v", m)
	}
}

func TestBeanToMapNil(t *testing.T) {
	if _, err := BeanToMap(nil); err == nil {
		t.Error("nil should error")
	}
	var b *source
	if _, err := BeanToMap(b); err == nil {
		t.Error("typed nil should error")
	}
}

func TestMapToBean(t *testing.T) {
	m := map[string]interface{}{
		"Name":  "Heidi",
		"Age":   31,
		"Email": "h@example.com",
	}
	b := &dest{}
	if err := MapToBean(&b, m, false); err != nil {
		t.Fatalf("MapToBean error: %v", err)
	}
	if b.Name != "Heidi" || b.Age != 31 || b.Email != "h@example.com" {
		t.Errorf("MapToBean wrong: %+v", b)
	}
}

func TestMapToBeanIgnoreCase(t *testing.T) {
	m := map[string]interface{}{
		"username": "ivan",
		"nickname": "I",
	}
	d := &camelDst{}
	opts := MapOptions{IgnoreCase: true}
	if err := MapToBean(&d, m, false, opts); err != nil {
		t.Fatal(err)
	}
	if d.Username != "ivan" || d.Nickname != "I" {
		t.Errorf("MapToBean ignoreCase failed: %+v", d)
	}
}

func TestMapToBeanTypeConversion(t *testing.T) {
	m := map[string]interface{}{
		"Age":  "33", // string -> int
		"Name": "Judy",
	}
	b := &dest{}
	if err := MapToBean(&b, m, false); err != nil {
		t.Fatalf("MapToBean conversion error: %v", err)
	}
	if b.Age != 33 {
		t.Errorf("string->int conversion failed, got %d", b.Age)
	}
}

func TestMapToBeanStructPointer(t *testing.T) {
	// passing **dest like Hutool does
	m := map[string]interface{}{"Name": "Karl"}
	b := &dest{}
	if err := MapToBean(&b, m, false); err != nil {
		t.Fatal(err)
	}
	if b.Name != "Karl" {
		t.Errorf("MapToBean(**dst) failed: %+v", b)
	}
}

func TestBeanRoundTrip(t *testing.T) {
	b := &source{Name: "Liam", Age: 44, Email: "liam@example.com", Private: "p"}
	m, err := BeanToMap(b)
	if err != nil {
		t.Fatal(err)
	}
	d := &dest{}
	if err := MapToBean(&d, m, false); err != nil {
		t.Fatal(err)
	}
	if d.Name != "Liam" || d.Age != 44 || d.Email != "liam@example.com" {
		t.Errorf("round trip failed: %+v", d)
	}
}

func TestMapToBeanErrors(t *testing.T) {
	if err := MapToBean(&dest{}, nil, false); err == nil {
		t.Error("nil map should error")
	}
	if err := MapToBean(nil, map[string]interface{}{}, false); err == nil {
		t.Error("nil bean should error")
	}
}

// --- 复杂结构测试 ---

type inner struct {
	Label string
}

type complexSrc struct {
	ID     int64
	Tags   []string
	Map    map[string]int
	Time   time.Time
	Ptr    *inner
	unexp  int
	When   time.Time
}

type complexDst struct {
	ID    int64
	Tags  []string
	Map   map[string]int
	Time  time.Time
	Ptr   *inner
	When  time.Time
}

func TestCopyPropertiesComplex(t *testing.T) {
	src := &complexSrc{
		ID:    7,
		Tags:  []string{"a", "b"},
		Map:   map[string]int{"x": 1},
		Time:  time.Date(2024, 5, 6, 7, 8, 9, 0, time.UTC),
		Ptr:   &inner{Label: "in"},
		unexp: 999,
	}
	dst := &complexDst{}
	if err := CopyProperties(dst, src); err != nil {
		t.Fatal(err)
	}
	if dst.ID != 7 || !strings.HasPrefix(dst.Time.String(), "2024-05-06") {
		t.Errorf("complex copy failed: %+v", dst)
	}
	if len(dst.Tags) != 2 || dst.Map["x"] != 1 || dst.Ptr == nil || dst.Ptr.Label != "in" {
		t.Errorf("complex fields not copied: %+v", dst)
	}
	if dst.When.IsZero() == false {
		t.Errorf("zero time should stay zero, got %v", dst.When)
	}
}

func TestCopyPropertiesTypeMismatch(t *testing.T) {
	src := &complexSrc{ID: 9}
	dst := &dest{} // ID is string-less dest, mismatched field types
	if err := CopyProperties(dst, src); err != nil {
		t.Fatal(err)
	}
	// ID int64 vs dest has no ID field -> nothing to do, no error
}

func TestMapToBeanNestedNil(t *testing.T) {
	m := map[string]interface{}{
		"Tags": nil,
		"ID":   5,
	}
	b := &complexDst{}
	if err := MapToBean(&b, m, false); err != nil {
		t.Fatal(err)
	}
	if b.ID != 5 {
		t.Errorf("MapToBean with nil slice: %+v", b)
	}
	if b.Tags != nil {
		t.Errorf("nil should stay nil, got %v", b.Tags)
	}
}