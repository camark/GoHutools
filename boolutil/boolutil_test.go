package boolutil

import "testing"

func TestTrueFalse(t *testing.T) {
	if !IsTrue(true) {
		t.Error("IsTrue(true)")
	}
	if IsTrue(false) {
		t.Error("IsTrue(false)")
	}
	if !IsFalse(false) {
		t.Error("IsFalse(false)")
	}
	if IsFalse(true) {
		t.Error("IsFalse(true)")
	}
}

func TestParse(t *testing.T) {
	ok, v := Parse("true")
	if !ok || !v {
		t.Errorf("Parse(true) = %v, %v", ok, v)
	}
	ok, v = Parse("false")
	if !ok || v {
		t.Errorf("Parse(false) = %v, %v", ok, v)
	}
	ok, _ = Parse("nope")
	if ok {
		t.Error("Parse(bad) should fail")
	}
}

func TestParseLenient(t *testing.T) {
	// Hutool accepts "1"/"0"/"yes"/"no"/"on"/"off"
	if !ParseLenient("1") {
		t.Error("ParseLenient(1)")
	}
	if ParseLenient("0") {
		t.Error("ParseLenient(0)")
	}
	if !ParseLenient("TRUE") || !ParseLenient("True") {
		t.Error("ParseLenient uppercase")
	}
	if !ParseLenient("yes") || !ParseLenient("on") {
		t.Error("ParseLenient yes/on")
	}
	if ParseLenient("nope") {
		t.Error("ParseLenient(nope) should be false")
	}
}

func TestAndOrXor(t *testing.T) {
	if !And(true, true, true) {
		t.Error("And all true")
	}
	if And(true, false) {
		t.Error("And with false")
	}
	if And() {
		t.Error("And empty should be false")
	}
	if !Or(false, true) {
		t.Error("Or with true")
	}
	if Or(false, false) {
		t.Error("Or all false")
	}
	if Or() {
		t.Error("Or empty should be false")
	}
	if !Xor(true, false) {
		t.Error("Xor true,false")
	}
	if Xor(true, true) {
		t.Error("Xor true,true")
	}
}

func TestToInt(t *testing.T) {
	if ToInt(true) != 1 || ToInt(false) != 0 {
		t.Error("ToInt wrong")
	}
	if ToString(true) != "true" || ToString(false) != "false" {
		t.Error("ToString wrong")
	}
}

func TestNegate(t *testing.T) {
	if Negate(true) != false || Negate(false) != true {
		t.Error("Negate wrong")
	}
}

func TestIsTrueValue(t *testing.T) {
	// string forms
	if !IsTrueValue("true") || !IsTrueValue("TRUE") || !IsTrueValue("1") || !IsTrueValue("yes") {
		t.Error("IsTrueValue string")
	}
	if IsTrueValue("false") || IsTrueValue("0") {
		t.Error("IsTrueValue should be false for falsey")
	}
	// int forms
	if !IsTrueValue(1) || IsTrueValue(0) {
		t.Error("IsTrueValue int")
	}
	// bool form
	if !IsTrueValue(true) || IsTrueValue(false) {
		t.Error("IsTrueValue bool")
	}
	// number forms follow truthiness (non-zero = true)
	if !IsTrueValue(2.5) || IsTrueValue(0.0) {
		t.Error("IsTrueValue numeric truthiness")
	}
	// unsupported types are false
	if IsTrueValue([]int{1}) {
		t.Error("IsTrueValue(slice) should be false")
	}
}

func TestComparable(t *testing.T) {
	if Compare(true, true) != 0 {
		t.Error("Compare equal")
	}
	if Compare(true, false) != 1 {
		t.Error("Compare true>false")
	}
	if Compare(false, true) != -1 {
		t.Error("Compare false<true")
	}
}

func TestIsBoolean(t *testing.T) {
	if !IsBoolean("true") || !IsBoolean("false") {
		t.Error("IsBoolean true/false")
	}
	if IsBoolean("nope") {
		t.Error("IsBoolean(nope)")
	}
}