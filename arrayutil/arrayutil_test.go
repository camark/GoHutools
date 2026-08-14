package arrayutil

import (
	"reflect"
	"sort"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	if !IsEmpty([]int{}) {
		t.Error("expected empty slice to be empty")
	}
	if IsEmpty([]int{1}) {
		t.Error("expected non-empty slice to not be empty")
	}
	if !IsEmpty([]int(nil)) {
		t.Error("expected nil slice to be empty")
	}
}

func TestIsNotEmpty(t *testing.T) {
	if IsNotEmpty([]int{}) {
		t.Error("expected empty slice to not be not-empty")
	}
	if !IsNotEmpty([]int{1}) {
		t.Error("expected non-empty slice to be not-empty")
	}
}

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !Contains(arr, 3) {
		t.Error("expected array to contain 3")
	}
	if Contains(arr, 6) {
		t.Error("expected array to not contain 6")
	}
}

func TestIndexOf(t *testing.T) {
	arr := []string{"a", "b", "c", "b"}
	if IndexOf(arr, "b") != 1 {
		t.Errorf("expected index 1, got %d", IndexOf(arr, "b"))
	}
	if IndexOf(arr, "d") != -1 {
		t.Errorf("expected -1, got %d", IndexOf(arr, "d"))
	}
}

func TestLastIndexOf(t *testing.T) {
	arr := []string{"a", "b", "c", "b"}
	if LastIndexOf(arr, "b") != 3 {
		t.Errorf("expected index 3, got %d", LastIndexOf(arr, "b"))
	}
	if LastIndexOf(arr, "d") != -1 {
		t.Errorf("expected -1, got %d", LastIndexOf(arr, "d"))
	}
}

func TestFill(t *testing.T) {
	arr := []int{1, 2, 3}
	Fill(arr, 0)
	expected := []int{0, 0, 0}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func TestCopyOf(t *testing.T) {
	arr := []int{1, 2, 3}
	copied := CopyOf(arr, 5)
	if len(copied) != 5 {
		t.Errorf("expected length 5, got %d", len(copied))
	}
	if copied[0] != 1 || copied[1] != 2 || copied[2] != 3 {
		t.Error("unexpected values in copied array")
	}
	if copied[3] != 0 || copied[4] != 0 {
		t.Error("expected zero values for new elements")
	}

	truncated := CopyOf(arr, 2)
	if len(truncated) != 2 {
		t.Errorf("expected length 2, got %d", len(truncated))
	}
	if truncated[0] != 1 || truncated[1] != 2 {
		t.Error("unexpected values in truncated array")
	}
}

func TestCopyOfRange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	copied := CopyOfRange(arr, 1, 4)
	expected := []int{2, 3, 4}
	if !reflect.DeepEqual(copied, expected) {
		t.Errorf("expected %v, got %v", expected, copied)
	}
}

func TestInsert(t *testing.T) {
	arr := []int{1, 2, 3}
	inserted := Insert(arr, 1, 10)
	expected := []int{1, 10, 2, 3}
	if !reflect.DeepEqual(inserted, expected) {
		t.Errorf("expected %v, got %v", expected, inserted)
	}
}

func TestRemoveAt(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	removed := RemoveAt(arr, 1)
	expected := []int{1, 3, 4}
	if !reflect.DeepEqual(removed, expected) {
		t.Errorf("expected %v, got %v", expected, removed)
	}
}

func TestSwap(t *testing.T) {
	arr := []int{1, 2, 3}
	Swap(arr, 0, 2)
	expected := []int{3, 2, 1}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	Reverse(arr)
	expected := []int{4, 3, 2, 1}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func TestRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	rotated := Rotate(arr, 2)
	expected := []int{4, 5, 1, 2, 3}
	if !reflect.DeepEqual(rotated, expected) {
		t.Errorf("expected %v, got %v", expected, rotated)
	}
}

func TestUnique(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3, 3}
	unique := Unique(arr)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(unique, expected) {
		t.Errorf("expected %v, got %v", expected, unique)
	}
}

func TestToInterfaceSlice(t *testing.T) {
	arr := []int{1, 2, 3}
	result := ToInterfaceSlice(arr)
	if len(result) != 3 {
		t.Errorf("expected length 3, got %d", len(result))
	}
	if result[0] != 1 || result[1] != 2 || result[2] != 3 {
		t.Error("unexpected values in interface slice")
	}
}

func TestFromInterfaceSlice(t *testing.T) {
	arr := []interface{}{1, 2, 3}
	result, ok := FromInterfaceSlice[int](arr)
	if !ok {
		t.Error("expected conversion to succeed")
	}
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	invalid := []interface{}{1, "two", 3}
	_, ok = FromInterfaceSlice[int](invalid)
	if ok {
		t.Error("expected conversion to fail")
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(arr, func(v int) bool {
		return v%2 == 0
	})
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(evens, expected) {
		t.Errorf("expected %v, got %v", expected, evens)
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	doubled := Map(arr, func(v int) int {
		return v * 2
	})
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(doubled, expected) {
		t.Errorf("expected %v, got %v", expected, doubled)
	}
}

func TestAny(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !Any(arr, func(v int) bool { return v > 3 }) {
		t.Error("expected any element > 3")
	}
	if Any(arr, func(v int) bool { return v > 10 }) {
		t.Error("expected no element > 10")
	}
}

func TestAll(t *testing.T) {
	arr := []int{2, 4, 6}
	if !All(arr, func(v int) bool { return v%2 == 0 }) {
		t.Error("expected all elements to be even")
	}
	if All(arr, func(v int) bool { return v > 4 }) {
		t.Error("expected not all elements > 4")
	}
}

func TestNone(t *testing.T) {
	arr := []int{1, 3, 5}
	if !None(arr, func(v int) bool { return v%2 == 0 }) {
		t.Error("expected no even elements")
	}
	if None(arr, func(v int) bool { return v > 3 }) {
		t.Error("expected some elements > 3")
	}
}

func TestJoin(t *testing.T) {
	arr := []int{1, 2, 3}
	result := Join(arr, ", ")
	expected := "1, 2, 3"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}

	if Join([]int{}, ",") != "" {
		t.Error("expected empty string for empty slice")
	}
}

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	if !Equal(a, b) {
		t.Error("expected a and b to be equal")
	}
	if Equal(a, c) {
		t.Error("expected a and c to not be equal")
	}
}

func TestConcat(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	c := []int{5}
	result := Concat(a, b, c)
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFirst(t *testing.T) {
	arr := []int{1, 2, 3}
	v, ok := First(arr)
	if !ok || v != 1 {
		t.Errorf("expected 1, got %d, ok=%v", v, ok)
	}
	_, ok = First([]int{})
	if ok {
		t.Error("expected false for empty slice")
	}
}

func TestLast(t *testing.T) {
	arr := []int{1, 2, 3}
	v, ok := Last(arr)
	if !ok || v != 3 {
		t.Errorf("expected 3, got %d, ok=%v", v, ok)
	}
	_, ok = Last([]int{})
	if ok {
		t.Error("expected false for empty slice")
	}
}

func TestShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffled := Shuffle(arr)
	if len(shuffled) != len(arr) {
		t.Errorf("expected length %d, got %d", len(arr), len(shuffled))
	}
	// Check all elements are present
	sort.Ints(shuffled)
	if !reflect.DeepEqual(arr, shuffled) {
		t.Error("expected shuffled to contain same elements")
	}
}

func TestMin(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5}
	min, ok := Min(arr)
	if !ok || min != 1 {
		t.Errorf("expected 1, got %d, ok=%v", min, ok)
	}
	_, ok = Min([]int{})
	if ok {
		t.Error("expected false for empty slice")
	}
}

func TestMax(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5}
	max, ok := Max(arr)
	if !ok || max != 5 {
		t.Errorf("expected 5, got %d, ok=%v", max, ok)
	}
	_, ok = Max([]int{})
	if ok {
		t.Error("expected false for empty slice")
	}
}

func TestFlatten(t *testing.T) {
	arr := [][]int{{1, 2}, {3, 4}, {5}}
	result := Flatten(arr)
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestChunk(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	chunks := Chunk(arr, 2)
	expected := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected %v, got %v", expected, chunks)
	}
}

func TestWindow(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	windows := Window(arr, 3)
	expected := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	if !reflect.DeepEqual(windows, expected) {
		t.Errorf("expected %v, got %v", expected, windows)
	}
}

func TestZip(t *testing.T) {
	a := []int{1, 2, 3}
	b := []string{"a", "b", "c"}
	pairs := Zip(a, b)
	expected := []Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 2, Second: "b"},
		{First: 3, Second: "c"},
	}
	if !reflect.DeepEqual(pairs, expected) {
		t.Errorf("expected %v, got %v", expected, pairs)
	}
}

func TestUnzip(t *testing.T) {
	pairs := []Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 2, Second: "b"},
		{First: 3, Second: "c"},
	}
	a, b := Unzip(pairs)
	expectedA := []int{1, 2, 3}
	expectedB := []string{"a", "b", "c"}
	if !reflect.DeepEqual(a, expectedA) {
		t.Errorf("expected %v, got %v", expectedA, a)
	}
	if !reflect.DeepEqual(b, expectedB) {
		t.Errorf("expected %v, got %v", expectedB, b)
	}
}
