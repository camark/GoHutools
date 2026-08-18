package assert

import (
	"errors"
	"strings"
	"testing"
)

type fakeType struct{}

func TestTrue(t *testing.T) {
	if err := True(true); err != nil {
		t.Errorf("True(true) should pass, got error: %v", err)
	}
	err := True(false)
	if err == nil {
		t.Fatal("True(false) should fail")
	}

	// custom message
	err = True(false, "value must be true")
	if err == nil || !strings.Contains(err.Error(), "value must be true") {
		t.Errorf("custom message not used, got: %v", err)
	}
}

func TestFalse(t *testing.T) {
	if err := False(false); err != nil {
		t.Errorf("False(false) should pass, got error: %v", err)
	}
	if err := False(true); err == nil {
		t.Error("False(true) should fail")
	}
}

func TestNotNull(t *testing.T) {
	if err := NotNull(42); err != nil {
		t.Errorf("NotNull(42) should pass, got: %v", err)
	}
	if err := NotNull(nil); err == nil {
		t.Error("NotNull(nil) should fail")
	}

	// typed nil pointer
	var p *fakeType
	if err := NotNull(p); err == nil {
		t.Error("NotNull(typed nil ptr) should fail")
	}
}

func TestIsNull(t *testing.T) {
	if err := IsNull(nil); err != nil {
		t.Errorf("IsNull(nil) should pass, got: %v", err)
	}
	if err := IsNull(42); err == nil {
		t.Error("IsNull(42) should fail")
	}
	var p *fakeType
	if err := IsNull(p); err != nil {
		t.Errorf("IsNull(typed nil ptr) should pass, got: %v", err)
	}
}

func TestNotEmpty(t *testing.T) {
	if err := NotEmpty("hello"); err != nil {
		t.Errorf("NotEmpty(\"hello\") should pass, got: %v", err)
	}
	if err := NotEmpty(""); err == nil {
		t.Error("NotEmpty(\"\") should fail")
	}
	if err := NotEmpty([]int{1, 2}); err != nil {
		t.Errorf("NotEmpty([]int{1,2}) should pass, got: %v", err)
	}
	if err := NotEmpty([]int{}); err == nil {
		t.Error("NotEmpty([]int{}) should fail")
	}
	if err := NotEmpty(0); err == nil {
		t.Error("NotEmpty(0) should fail")
	}
	if err := NotEmpty(" "); err == nil {
		t.Error("NotEmpty(\" \") should fail (blank counts as empty)")
	}
	if err := NotEmpty(map[string]int{}); err == nil {
		t.Error("NotEmpty(empty map) should fail")
	}
	// nil pointer should fail
	var p *fakeType
	if err := NotEmpty(p); err == nil {
		t.Error("NotEmpty(nil ptr) should fail")
	}
}

func TestNotBlank(t *testing.T) {
	if err := NotBlank("x"); err != nil {
		t.Errorf("NotBlank(\"x\") should pass, got: %v", err)
	}
	if err := NotBlank("  "); err == nil {
		t.Error("NotBlank(\"  \") should fail")
	}
	if err := NotBlank(""); err == nil {
		t.Error("NotBlank(\"\") should fail")
	}
	if err := NotBlank(42); err != nil {
		t.Errorf("NotBlank(42) should pass (non-string), got: %v", err)
	}
}

func TestEqual(t *testing.T) {
	if err := Equal(1, 1); err != nil {
		t.Errorf("Equal(1,1) should pass, got: %v", err)
	}
	if err := Equal(1, 2); err == nil {
		t.Error("Equal(1,2) should fail")
	}
	if err := Equal("a", "a"); err != nil {
		t.Errorf("Equal(a,a) should pass, got: %v", err)
	}
	if err := Equal(&fakeType{}, &fakeType{}); err != nil {
		t.Errorf("Equal(equal struct values through pointers) should pass, got: %v", err)
	}
	a := &fakeType{}
	if err := Equal(a, a); err != nil {
		t.Errorf("Equal(same pointer) should pass, got: %v", err)
	}
}

func TestNotEqual(t *testing.T) {
	if err := NotEqual(1, 2); err != nil {
		t.Errorf("NotEqual(1,2) should pass, got: %v", err)
	}
	if err := NotEqual(1, 1); err == nil {
		t.Error("NotEqual(1,1) should fail")
	}
}

func TestIsInstanceOf(t *testing.T) {
	var s string
	if err := IsInstanceOf(s, "string"); err != nil {
		t.Errorf("IsInstanceOf(string, string) should pass, got: %v", err)
	}
	if err := IsInstanceOf(42, "string"); err == nil {
		t.Error("IsInstanceOf(int, string) should fail")
	}

	// type literal form
	if err := IsInstanceOf(42, int(0)); err != nil {
		t.Errorf("IsInstanceOf(42, int) should pass, got: %v", err)
	}
}

func TestMatches(t *testing.T) {
	if err := Matches("^\\d+$", "12345"); err != nil {
		t.Errorf("Matches(digits, 12345) should pass, got: %v", err)
	}
	if err := Matches("^\\d+$", "12a"); err == nil {
		t.Error("Matches(digits, 12a) should fail")
	}
	if err := Matches("^[a-z]+$", "abc"); err != nil {
		t.Errorf("Matches(lowercase, abc) should pass, got: %v", err)
	}
}

func TestNoError(t *testing.T) {
	if err := NoError(nil); err != nil {
		t.Errorf("NoError(nil) should pass, got: %v", err)
	}
	if err := NoError(errors.New("boom")); err == nil {
		t.Error("NoError(boom) should fail")
	}
}

func TestErrorContains(t *testing.T) {
	if err := ErrorContains(errors.New("connection refused"), "refused"); err != nil {
		t.Errorf("ErrorContains(refused) should pass, got: %v", err)
	}
	if err := ErrorContains(errors.New("connection reset"), "refused"); err == nil {
		t.Error("ErrorContains(reset, refused) should fail")
	}
	if err := ErrorContains(nil, "x"); err == nil {
		t.Error("ErrorContains(nil) should fail")
	}
}

func TestLessThan(t *testing.T) {
	if err := LessThan(1, 2); err != nil {
		t.Errorf("LessThan(1,2) should pass, got: %v", err)
	}
	if err := LessThan(2, 1); err == nil {
		t.Error("LessThan(2,1) should fail")
	}
	if err := LessThan("a", "b"); err != nil {
		t.Errorf("LessThan(a,b) should pass, got: %v", err)
	}
}

func TestGreaterThan(t *testing.T) {
	if err := GreaterThan(2, 1); err != nil {
		t.Errorf("GreaterThan(2,1) should pass, got: %v", err)
	}
	if err := GreaterThan(1, 2); err == nil {
		t.Error("GreaterThan(1,2) should fail")
	}
}

func TestNotNegative(t *testing.T) {
	if err := NotNegative(5); err != nil {
		t.Errorf("NotNegative(5) should pass, got: %v", err)
	}
	if err := NotNegative(0); err != nil {
		t.Errorf("NotNegative(0) should pass, got: %v", err)
	}
	if err := NotNegative(-1); err == nil {
		t.Error("NotNegative(-1) should fail")
	}
}

func TestNotEmptyErrorMessage(t *testing.T) {
	// default error message should include the argument
	err := NotEmpty("")
	if err == nil || err.Error() == "" {
		t.Error("default error message should be meaningful")
	}
	t.Logf("sample default error: %s", err)
}