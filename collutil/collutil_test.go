package collutil

import (
	"math"
	"testing"
	"time"
)

func TestIsEmpty(t *testing.T) {
	if !IsEmpty([]int{}) {
		t.Error("IsEmpty should return true for empty slice")
	}
	if IsEmpty([]int{1, 2, 3}) {
		t.Error("IsEmpty should return false for non-empty slice")
	}
	if !IsEmpty([]string{}) {
		t.Error("IsEmpty should return true for empty string slice")
	}
	if !IsEmpty[int](nil) {
		t.Error("IsEmpty should return true for nil slice")
	}
}

func TestIsNotEmpty(t *testing.T) {
	if IsNotEmpty([]int{}) {
		t.Error("IsNotEmpty should return false for empty slice")
	}
	if !IsNotEmpty([]int{1, 2, 3}) {
		t.Error("IsNotEmpty should return true for non-empty slice")
	}
}

func TestContains(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	if !Contains(coll, 3) {
		t.Error("Contains should return true for existing element")
	}
	if Contains(coll, 6) {
		t.Error("Contains should return false for non-existing element")
	}
	if Contains([]int{}, 1) {
		t.Error("Contains should return false for empty slice")
	}
}

func TestContainsAll(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	if !ContainsAll(coll, 1, 3, 5) {
		t.Error("ContainsAll should return true when all elements exist")
	}
	if ContainsAll(coll, 1, 3, 6) {
		t.Error("ContainsAll should return false when any element is missing")
	}
	if !ContainsAll(coll) {
		t.Error("ContainsAll should return true for empty items")
	}
}

func TestContainsAny(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	if !ContainsAny(coll, 3, 6, 7) {
		t.Error("ContainsAny should return true when at least one element exists")
	}
	if ContainsAny(coll, 6, 7, 8) {
		t.Error("ContainsAny should return false when no elements exist")
	}
	if ContainsAny(coll) {
		t.Error("ContainsAny should return false for empty items")
	}
}

func TestAddIfAbsent(t *testing.T) {
	coll := []int{1, 2, 3}
	result := AddIfAbsent(coll, 4)
	if len(result) != 4 || result[3] != 4 {
		t.Error("AddIfAbsent should add element when not present")
	}
	result = AddIfAbsent(coll, 2)
	if len(result) != 3 {
		t.Error("AddIfAbsent should not add element when already present")
	}
}

func TestRemove(t *testing.T) {
	coll := []int{1, 2, 3, 2, 4}
	result := Remove(coll, 2)
	if len(result) != 4 {
		t.Errorf("Remove should remove first occurrence, got length %d", len(result))
	}
	if result[1] != 3 {
		t.Error("Remove should remove only first occurrence")
	}
	result = Remove(coll, 5)
	if len(result) != 5 {
		t.Error("Remove should return original slice when element not found")
	}
}

func TestRemoveAll(t *testing.T) {
	coll := []int{1, 2, 3, 2, 4, 2}
	result := RemoveAll(coll, 2)
	if len(result) != 3 {
		t.Errorf("RemoveAll should remove all occurrences, got length %d", len(result))
	}
	if result[0] != 1 || result[1] != 3 || result[2] != 4 {
		t.Error("RemoveAll should remove all occurrences of element")
	}
}

func TestDistinct(t *testing.T) {
	coll := []int{1, 2, 3, 2, 1, 4, 3}
	result := Distinct(coll)
	if len(result) != 4 {
		t.Errorf("Distinct should remove duplicates, got length %d", len(result))
	}
	expected := []int{1, 2, 3, 4}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Distinct result[%d] = %d, want %d", i, v, expected[i])
		}
	}
}

func TestReverse(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	result := Reverse(coll)
	expected := []int{5, 4, 3, 2, 1}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Reverse result[%d] = %d, want %d", i, v, expected[i])
		}
	}
	if coll[0] != 1 {
		t.Error("Reverse should not modify original slice")
	}
}

func TestSub(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	result := Sub(coll, 1, 3)
	if len(result) != 2 || result[0] != 2 || result[1] != 3 {
		t.Error("Sub should extract correct sub-slice")
	}
	result = Sub(coll, -1, 3)
	if len(result) != 3 {
		t.Error("Sub should handle negative start")
	}
	result = Sub(coll, 2, 10)
	if len(result) != 3 {
		t.Error("Sub should handle end beyond length")
	}
	result = Sub(coll, 3, 3)
	if len(result) != 0 {
		t.Error("Sub should return empty slice when start equals end")
	}
}

func TestSplit(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5, 6, 7}
	result := Split(coll, 3)
	if len(result) != 3 {
		t.Errorf("Split should create 3 chunks, got %d", len(result))
	}
	if len(result[0]) != 3 || len(result[1]) != 3 || len(result[2]) != 1 {
		t.Error("Split chunks have incorrect sizes")
	}
	result = Split(coll, 0)
	if len(result) != 0 {
		t.Error("Split should return empty for size 0")
	}
	result = Split([]int{}, 3)
	if len(result) != 0 {
		t.Error("Split should return empty for empty slice")
	}
}

func TestFilter(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5, 6}
	result := Filter(coll, func(n int) bool { return n%2 == 0 })
	if len(result) != 3 {
		t.Errorf("Filter should return 3 even numbers, got %d", len(result))
	}
	if result[0] != 2 || result[1] != 4 || result[2] != 6 {
		t.Error("Filter should return correct even numbers")
	}
}

func TestMap(t *testing.T) {
	coll := []int{1, 2, 3}
	result := Map(coll, func(n int) string {
		return string(rune('A' + n - 1))
	})
	if len(result) != 3 {
		t.Errorf("Map should return 3 elements, got %d", len(result))
	}
	if result[0] != "A" || result[1] != "B" || result[2] != "C" {
		t.Error("Map should transform elements correctly")
	}
}

func TestForEach(t *testing.T) {
	coll := []int{1, 2, 3}
	sum := 0
	ForEach(coll, func(n int) {
		sum += n
	})
	if sum != 6 {
		t.Errorf("ForEach should execute consumer for each element, got sum %d", sum)
	}
}

func TestForEachIndexed(t *testing.T) {
	coll := []string{"a", "b", "c"}
	result := make([]string, len(coll))
	ForEachIndexed(coll, func(i int, s string) {
		result[i] = s + string(rune('0'+i))
	})
	if result[0] != "a0" || result[1] != "b1" || result[2] != "c2" {
		t.Error("ForEachIndexed should pass correct index")
	}
}

func TestGroupBy(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5, 6}
	result := GroupBy(coll, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if len(result) != 2 {
		t.Errorf("GroupBy should create 2 groups, got %d", len(result))
	}
	if len(result["even"]) != 3 || len(result["odd"]) != 3 {
		t.Error("GroupBy groups have incorrect sizes")
	}
}

func TestToMap(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	coll := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	result := ToMap(coll, func(p Person) string { return p.Name }, func(p Person) int { return p.Age })
	if len(result) != 3 {
		t.Errorf("ToMap should create 3 entries, got %d", len(result))
	}
	if result["Alice"] != 30 || result["Bob"] != 25 {
		t.Error("ToMap should create correct key-value pairs")
	}
}

func TestSort(t *testing.T) {
	coll := []int{3, 1, 4, 1, 5, 9, 2, 6}
	result := Sort(coll)
	expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Sort result[%d] = %d, want %d", i, v, expected[i])
		}
	}
	if coll[0] != 3 {
		t.Error("Sort should not modify original slice")
	}
}

func TestSortBy(t *testing.T) {
	coll := []int{3, 1, 4, 1, 5, 9, 2, 6}
	result := SortBy(coll, func(a, b int) bool { return a > b })
	expected := []int{9, 6, 5, 4, 3, 2, 1, 1}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("SortBy result[%d] = %d, want %d", i, v, expected[i])
		}
	}
}

func TestMin(t *testing.T) {
	coll := []int{3, 1, 4, 1, 5, 9, 2, 6}
	min, ok := Min(coll)
	if !ok || min != 1 {
		t.Errorf("Min should return 1, got %d", min)
	}
	_, ok = Min([]int{})
	if ok {
		t.Error("Min should return false for empty slice")
	}
}

func TestMax(t *testing.T) {
	coll := []int{3, 1, 4, 1, 5, 9, 2, 6}
	max, ok := Max(coll)
	if !ok || max != 9 {
		t.Errorf("Max should return 9, got %d", max)
	}
	_, ok = Max([]int{})
	if ok {
		t.Error("Max should return false for empty slice")
	}
}

func TestSum(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	sum := Sum(coll)
	if sum != 15 {
		t.Errorf("Sum should return 15, got %d", sum)
	}
	sum = Sum([]int{})
	if sum != 0 {
		t.Error("Sum should return 0 for empty slice")
	}
}

func TestFlatten(t *testing.T) {
	colls := [][]int{{1, 2}, {3, 4}, {5, 6}}
	result := Flatten(colls)
	if len(result) != 6 {
		t.Errorf("Flatten should return 6 elements, got %d", len(result))
	}
	expected := []int{1, 2, 3, 4, 5, 6}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Flatten result[%d] = %d, want %d", i, v, expected[i])
		}
	}
}

func TestChunk(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	result := Chunk(coll, 2)
	if len(result) != 3 {
		t.Errorf("Chunk should create 3 chunks, got %d", len(result))
	}
	if len(result[0]) != 2 || len(result[1]) != 2 || len(result[2]) != 1 {
		t.Error("Chunk should create correct chunks")
	}
}

func TestZip(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := []int{1, 2, 3}
	result := Zip(keys, values)
	if len(result) != 3 {
		t.Errorf("Zip should create 3 entries, got %d", len(result))
	}
	if result["a"] != 1 || result["b"] != 2 || result["c"] != 3 {
		t.Error("Zip should create correct key-value pairs")
	}
	keys = []string{"a", "b", "c", "d"}
	result = Zip(keys, values)
	if len(result) != 3 {
		t.Error("Zip should handle different lengths")
	}
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	result := Keys(m)
	if len(result) != 3 {
		t.Errorf("Keys should return 3 keys, got %d", len(result))
	}
	keySet := make(map[string]bool)
	for _, k := range result {
		keySet[k] = true
	}
	if !keySet["a"] || !keySet["b"] || !keySet["c"] {
		t.Error("Keys should return all keys")
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	result := Values(m)
	if len(result) != 3 {
		t.Errorf("Values should return 3 values, got %d", len(result))
	}
	sum := 0
	for _, v := range result {
		sum += v
	}
	if sum != 6 {
		t.Error("Values should return all values")
	}
}

func TestFirst(t *testing.T) {
	coll := []int{1, 2, 3}
	first, ok := First(coll)
	if !ok || first != 1 {
		t.Errorf("First should return 1, got %d", first)
	}
	_, ok = First([]int{})
	if ok {
		t.Error("First should return false for empty slice")
	}
}

func TestLast(t *testing.T) {
	coll := []int{1, 2, 3}
	last, ok := Last(coll)
	if !ok || last != 3 {
		t.Errorf("Last should return 3, got %d", last)
	}
	_, ok = Last([]int{})
	if ok {
		t.Error("Last should return false for empty slice")
	}
}

func TestFind(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	found, ok := Find(coll, func(n int) bool { return n > 3 })
	if !ok || found != 4 {
		t.Errorf("Find should return 4, got %d", found)
	}
	_, ok = Find(coll, func(n int) bool { return n > 10 })
	if ok {
		t.Error("Find should return false when no match")
	}
}

func TestFindIndex(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	idx := FindIndex(coll, func(n int) bool { return n > 3 })
	if idx != 3 {
		t.Errorf("FindIndex should return 3, got %d", idx)
	}
	idx = FindIndex(coll, func(n int) bool { return n > 10 })
	if idx != -1 {
		t.Errorf("FindIndex should return -1, got %d", idx)
	}
}

func TestAnyMatch(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	if !AnyMatch(coll, func(n int) bool { return n > 3 }) {
		t.Error("AnyMatch should return true when any element matches")
	}
	if AnyMatch(coll, func(n int) bool { return n > 10 }) {
		t.Error("AnyMatch should return false when no element matches")
	}
}

func TestAllMatch(t *testing.T) {
	coll := []int{2, 4, 6, 8}
	if !AllMatch(coll, func(n int) bool { return n%2 == 0 }) {
		t.Error("AllMatch should return true when all elements match")
	}
	coll = []int{2, 3, 4, 6}
	if AllMatch(coll, func(n int) bool { return n%2 == 0 }) {
		t.Error("AllMatch should return false when not all elements match")
	}
}

func TestNoneMatch(t *testing.T) {
	coll := []int{1, 3, 5, 7}
	if !NoneMatch(coll, func(n int) bool { return n%2 == 0 }) {
		t.Error("NoneMatch should return true when no elements match")
	}
	coll = []int{1, 2, 3, 5}
	if NoneMatch(coll, func(n int) bool { return n%2 == 0 }) {
		t.Error("NoneMatch should return false when any element matches")
	}
}

func TestReduce(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5}
	sum := Reduce(coll, 0, func(acc, n int) int { return acc + n })
	if sum != 15 {
		t.Errorf("Reduce should return 15, got %d", sum)
	}
	concat := Reduce(coll, "", func(acc string, n int) string {
		return acc + string(rune('0'+n))
	})
	if concat != "12345" {
		t.Errorf("Reduce should return '12345', got '%s'", concat)
	}
}

func TestCount(t *testing.T) {
	coll := []int{1, 2, 3, 4, 5, 6}
	count := Count(coll, func(n int) bool { return n%2 == 0 })
	if count != 3 {
		t.Errorf("Count should return 3, got %d", count)
	}
	count = Count(coll, func(n int) bool { return n > 10 })
	if count != 0 {
		t.Errorf("Count should return 0, got %d", count)
	}
}

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	if !Equal(a, b) {
		t.Error("Equal should return true for equal slices")
	}
	b = []int{1, 2, 4}
	if Equal(a, b) {
		t.Error("Equal should return false for different slices")
	}
	b = []int{1, 2}
	if Equal(a, b) {
		t.Error("Equal should return false for different lengths")
	}
}

func TestUnion(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 4, 5}
	result := Union(a, b)
	if len(result) != 5 {
		t.Errorf("Union should return 5 elements, got %d", len(result))
	}
	expected := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	for _, v := range result {
		if !expected[v] {
			t.Errorf("Union contains unexpected element %d", v)
		}
	}
}

func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	result := Intersection(a, b)
	if len(result) != 2 {
		t.Errorf("Intersection should return 2 elements, got %d", len(result))
	}
	expected := map[int]bool{3: true, 4: true}
	for _, v := range result {
		if !expected[v] {
			t.Errorf("Intersection contains unexpected element %d", v)
		}
	}
}

func TestDifference(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}
	result := Difference(a, b)
	if len(result) != 2 {
		t.Errorf("Difference should return 2 elements, got %d", len(result))
	}
	expected := map[int]bool{1: true, 2: true}
	for _, v := range result {
		if !expected[v] {
			t.Errorf("Difference contains unexpected element %d", v)
		}
	}
}

func TestToSlice(t *testing.T) {
	result := ToSlice(1, 2, 3, 4, 5)
	if len(result) != 5 {
		t.Errorf("ToSlice should return 5 elements, got %d", len(result))
	}
	for i, v := range result {
		if v != i+1 {
			t.Errorf("ToSlice result[%d] = %d, want %d", i, v, i+1)
		}
	}
	result = ToSlice[int]()
	if len(result) != 0 {
		t.Error("ToSlice should return empty slice for no args")
	}
}

func TestRange(t *testing.T) {
	result := Range(1, 5)
	if len(result) != 4 {
		t.Errorf("Range should return 4 elements, got %d", len(result))
	}
	expected := []int{1, 2, 3, 4}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Range result[%d] = %d, want %d", i, v, expected[i])
		}
	}
	result = Range(5, 1)
	if len(result) != 0 {
		t.Error("Range should return empty for start >= end")
	}
}

func TestRangeWithStep(t *testing.T) {
	result := RangeWithStep(0, 10, 2)
	if len(result) != 5 {
		t.Errorf("RangeWithStep should return 5 elements, got %d", len(result))
	}
	expected := []int{0, 2, 4, 6, 8}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("RangeWithStep result[%d] = %d, want %d", i, v, expected[i])
		}
	}
	result = RangeWithStep(10, 0, -2)
	if len(result) != 5 {
		t.Errorf("RangeWithStep should return 5 elements, got %d", len(result))
	}
	expected = []int{10, 8, 6, 4, 2}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("RangeWithStep result[%d] = %d, want %d", i, v, expected[i])
		}
	}
	result = RangeWithStep(0, 10, 0)
	if len(result) != 0 {
		t.Error("RangeWithStep should return empty for step 0")
	}

	// Integer overflow must not cause an infinite loop (regression)
	done := make(chan struct{})
	go func() {
		defer close(done)
		RangeWithStep(math.MaxInt-1, math.MaxInt, 2)
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("RangeWithStep hangs on integer overflow")
	}
}
