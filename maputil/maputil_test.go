package maputil

import (
	"reflect"
	"sort"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	if !IsEmpty(map[string]int{}) {
		t.Error("expected empty map to be empty")
	}
	if IsEmpty(map[string]int{"a": 1}) {
		t.Error("expected non-empty map to not be empty")
	}
	if !IsEmpty(map[int]int(nil)) {
		t.Error("expected nil map to be empty")
	}
}

func TestIsNotEmpty(t *testing.T) {
	if IsNotEmpty(map[string]int{}) {
		t.Error("expected empty map to not be not-empty")
	}
	if !IsNotEmpty(map[string]int{"a": 1}) {
		t.Error("expected non-empty map to be not-empty")
	}
}

func TestContainsKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	if !ContainsKey(m, "a") {
		t.Error("expected map to contain key 'a'")
	}
	if ContainsKey(m, "c") {
		t.Error("expected map to not contain key 'c'")
	}
}

func TestContainsValue(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	if !ContainsValue(m, 1) {
		t.Error("expected map to contain value 1")
	}
	if ContainsValue(m, 3) {
		t.Error("expected map to not contain value 3")
	}
}

func TestGet(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	if v := Get(m, "a", 0); v != 1 {
		t.Errorf("expected 1, got %d", v)
	}
	if v := Get(m, "c", 99); v != 99 {
		t.Errorf("expected 99, got %d", v)
	}
}

func TestGetOrDefault(t *testing.T) {
	m := map[string]int{"a": 1}
	if v := GetOrDefault(m, "a", 0); v != 1 {
		t.Errorf("expected 1, got %d", v)
	}
	if v := GetOrDefault(m, "b", 42); v != 42 {
		t.Errorf("expected 42, got %d", v)
	}
}

func TestPutIfAbsent(t *testing.T) {
	m := map[string]int{"a": 1}
	v := PutIfAbsent(m, "a", 10)
	if v != 1 {
		t.Errorf("expected existing value 1, got %d", v)
	}
	v = PutIfAbsent(m, "b", 2)
	if v != 2 {
		t.Errorf("expected new value 2, got %d", v)
	}
	if m["b"] != 2 {
		t.Error("expected 'b' to be inserted")
	}
}

func TestRemove(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	Remove(m, "a")
	if ContainsKey(m, "a") {
		t.Error("expected 'a' to be removed")
	}
	if len(m) != 1 {
		t.Errorf("expected map size 1, got %d", len(m))
	}
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := Keys(m)
	sort.Strings(keys)
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("expected %v, got %v", expected, keys)
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	vals := Values(m)
	sort.Ints(vals)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(vals, expected) {
		t.Errorf("expected %v, got %v", expected, vals)
	}
}

func TestForEach(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	sum := 0
	ForEach(m, func(k string, v int) {
		sum += v
	})
	if sum != 3 {
		t.Errorf("expected sum 3, got %d", sum)
	}
}

func TestFilter(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	filtered := Filter(m, func(k string, v int) bool {
		return v%2 == 0
	})
	if len(filtered) != 2 {
		t.Errorf("expected 2 entries, got %d", len(filtered))
	}
	if filtered["b"] != 2 || filtered["d"] != 4 {
		t.Error("unexpected filtered result")
	}
}

func TestTransform(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	transformed := Transform(m, func(k string, v int) string {
		return k + "=" + string(rune('0'+v))
	})
	if transformed["a"] != "a=1" {
		t.Errorf("expected 'a=1', got '%s'", transformed["a"])
	}
}

func TestMerge(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	merged := Merge(m1, m2)
	if merged["a"] != 1 || merged["b"] != 3 || merged["c"] != 4 {
		t.Errorf("unexpected merge result: %v", merged)
	}
}

func TestInvert(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	inverted := Invert(m)
	if inverted[1] != "a" || inverted[2] != "b" || inverted[3] != "c" {
		t.Errorf("unexpected invert result: %v", inverted)
	}
}

func TestFromKeys(t *testing.T) {
	keys := []string{"a", "b", "c"}
	m := FromKeys(keys, 0)
	if len(m) != 3 {
		t.Errorf("expected 3 entries, got %d", len(m))
	}
	for _, k := range keys {
		if m[k] != 0 {
			t.Errorf("expected %s=0, got %d", k, m[k])
		}
	}
}

func TestFromEntries(t *testing.T) {
	entries := []Entry[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
	}
	m := FromEntries(entries)
	if m["a"] != 1 || m["b"] != 2 {
		t.Errorf("unexpected result: %v", m)
	}
}

func TestToEntries(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	entries := ToEntries(m)
	if len(entries) != 2 {
		t.Errorf("expected 2 entries, got %d", len(entries))
	}
	entryMap := make(map[string]int)
	for _, e := range entries {
		entryMap[e.Key] = e.Value
	}
	if entryMap["a"] != 1 || entryMap["b"] != 2 {
		t.Errorf("unexpected entries: %v", entries)
	}
}

func TestEqual(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	m3 := map[string]int{"a": 1, "b": 3}
	if !Equal(m1, m2) {
		t.Error("expected m1 and m2 to be equal")
	}
	if Equal(m1, m3) {
		t.Error("expected m1 and m3 to not be equal")
	}
}

func TestSize(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	if Size(m) != 3 {
		t.Errorf("expected 3, got %d", Size(m))
	}
	if Size(map[string]int{}) != 0 {
		t.Error("expected 0 for empty map")
	}
}

func TestCopy(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	copied := Copy(m)
	if !Equal(m, copied) {
		t.Error("expected copied map to be equal")
	}
	copied["c"] = 3
	if ContainsKey(m, "c") {
		t.Error("expected original map to not be affected by changes to copy")
	}
}
