package jsonutil

import (
	"strings"
	"testing"
)

func TestMarshal(t *testing.T) {
	data := map[string]interface{}{
		"name": "test",
		"age":  30,
	}

	bytes, err := Marshal(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !IsValidJSON(bytes) {
		t.Fatal("expected valid JSON")
	}
}

func TestMarshalIndent(t *testing.T) {
	data := map[string]interface{}{
		"name": "test",
		"age":  30,
	}

	bytes, err := MarshalIndent(data, "", "  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	str := string(bytes)
	if !strings.Contains(str, "  ") {
		t.Fatal("expected indentation")
	}
}

func TestMarshalToString(t *testing.T) {
	data := map[string]interface{}{
		"name": "test",
		"age":  30,
	}

	str, err := MarshalToString(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !IsValidJSONString(str) {
		t.Fatal("expected valid JSON string")
	}
}

func TestUnmarshal(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`
	var data map[string]interface{}

	err := Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data["name"] != "test" {
		t.Fatalf("expected name 'test', got '%v'", data["name"])
	}
}

func TestUnmarshalFromString(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`
	var data map[string]interface{}

	err := UnmarshalFromString(jsonStr, &data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data["name"] != "test" {
		t.Fatalf("expected name 'test', got '%v'", data["name"])
	}
}

func TestUnmarshalFromReader(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`
	reader := strings.NewReader(jsonStr)
	var data map[string]interface{}

	err := UnmarshalFromReader(reader, &data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data["name"] != "test" {
		t.Fatalf("expected name 'test', got '%v'", data["name"])
	}
}

func TestIsValidJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`{"name":"test"}`, true},
		{`[1,2,3]`, true},
		{`"hello"`, true},
		{`123`, true},
		{`true`, true},
		{`null`, true},
		{`{invalid}`, false},
		{``, false},
	}

	for _, test := range tests {
		result := IsValidJSON([]byte(test.input))
		if result != test.expected {
			t.Fatalf("expected %v for '%s', got %v", test.expected, test.input, result)
		}
	}
}

func TestIsValidJSONString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`{"name":"test"}`, true},
		{`[1,2,3]`, true},
		{`{invalid}`, false},
	}

	for _, test := range tests {
		result := IsValidJSONString(test.input)
		if result != test.expected {
			t.Fatalf("expected %v for '%s', got %v", test.expected, test.input, result)
		}
	}
}

func TestToJSON(t *testing.T) {
	data := map[string]interface{}{
		"name": "test",
		"age":  30,
	}

	str := ToJSON(data)
	if str == "" {
		t.Fatal("expected non-empty JSON string")
	}

	if !IsValidJSONString(str) {
		t.Fatal("expected valid JSON")
	}
}

func TestToPrettyJSON(t *testing.T) {
	data := map[string]interface{}{
		"name": "test",
		"age":  30,
	}

	str := ToPrettyJSON(data)
	if str == "" {
		t.Fatal("expected non-empty JSON string")
	}

	if !strings.Contains(str, "\n") {
		t.Fatal("expected newlines in pretty JSON")
	}
}

func TestFromJSON(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`
	var data map[string]interface{}

	err := FromJSON(jsonStr, &data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data["name"] != "test" {
		t.Fatalf("expected name 'test', got '%v'", data["name"])
	}
}

func TestToMap(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`

	m, err := ToMap(jsonStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if m["name"] != "test" {
		t.Fatalf("expected name 'test', got '%v'", m["name"])
	}
	if m["age"] != float64(30) {
		t.Fatalf("expected age 30, got '%v'", m["age"])
	}
}

func TestToSlice(t *testing.T) {
	jsonStr := `[1,2,3]`

	s, err := ToSlice(jsonStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(s) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(s))
	}
	if s[0] != float64(1) {
		t.Fatalf("expected first element 1, got '%v'", s[0])
	}
}

func TestGetSimplePath(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`

	val, err := Get(jsonStr, "name")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "test" {
		t.Fatalf("expected 'test', got '%v'", val)
	}
}

func TestGetNestedPath(t *testing.T) {
	jsonStr := `{"user":{"name":"test","address":{"city":"NYC"}}}`

	val, err := Get(jsonStr, "user.address.city")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "NYC" {
		t.Fatalf("expected 'NYC', got '%v'", val)
	}
}

func TestGetArrayPath(t *testing.T) {
	jsonStr := `{"items":[1,2,3]}`

	val, err := Get(jsonStr, "items.1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != float64(2) {
		t.Fatalf("expected 2, got '%v'", val)
	}
}

func TestGetMissingKey(t *testing.T) {
	jsonStr := `{"name":"test"}`

	_, err := Get(jsonStr, "missing")
	if err == nil {
		t.Fatal("expected error for missing key")
	}
}

func TestGetString(t *testing.T) {
	jsonStr := `{"name":"test"}`

	val, err := GetString(jsonStr, "name")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "test" {
		t.Fatalf("expected 'test', got '%s'", val)
	}
}

func TestGetStringNotString(t *testing.T) {
	jsonStr := `{"age":30}`

	_, err := GetString(jsonStr, "age")
	if err == nil {
		t.Fatal("expected error for non-string value")
	}
}

func TestGetInt(t *testing.T) {
	jsonStr := `{"age":30}`

	val, err := GetInt(jsonStr, "age")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 30 {
		t.Fatalf("expected 30, got %d", val)
	}
}

func TestGetFloat(t *testing.T) {
	jsonStr := `{"price":19.99}`

	val, err := GetFloat(jsonStr, "price")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 19.99 {
		t.Fatalf("expected 19.99, got %f", val)
	}
}

func TestGetBool(t *testing.T) {
	jsonStr := `{"active":true}`

	val, err := GetBool(jsonStr, "active")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != true {
		t.Fatalf("expected true, got %v", val)
	}
}

func TestGetBoolNotBool(t *testing.T) {
	jsonStr := `{"active":"yes"}`

	_, err := GetBool(jsonStr, "active")
	if err == nil {
		t.Fatal("expected error for non-bool value")
	}
}

func TestSetSimplePath(t *testing.T) {
	jsonStr := `{"name":"test"}`

	result, err := Set(jsonStr, "name", "updated")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	name, _ := GetString(result, "name")
	if name != "updated" {
		t.Fatalf("expected 'updated', got '%s'", name)
	}
}

func TestSetNestedPath(t *testing.T) {
	jsonStr := `{"user":{"name":"test"}}`

	result, err := Set(jsonStr, "user.name", "updated")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	name, _ := GetString(result, "user.name")
	if name != "updated" {
		t.Fatalf("expected 'updated', got '%s'", name)
	}
}

func TestSetNewKey(t *testing.T) {
	jsonStr := `{"name":"test"}`

	result, err := Set(jsonStr, "age", 30)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	age, _ := GetInt(result, "age")
	if age != 30 {
		t.Fatalf("expected 30, got %d", age)
	}
}

func TestDeleteSimplePath(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`

	result, err := Delete(jsonStr, "age")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if Contains(result, "age") {
		t.Fatal("expected 'age' to be deleted")
	}

	if !Contains(result, "name") {
		t.Fatal("expected 'name' to still exist")
	}
}

func TestDeleteNestedPath(t *testing.T) {
	jsonStr := `{"user":{"name":"test","age":30}}`

	result, err := Delete(jsonStr, "user.age")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if Contains(result, "user.age") {
		t.Fatal("expected 'user.age' to be deleted")
	}

	if !Contains(result, "user.name") {
		t.Fatal("expected 'user.name' to still exist")
	}
}

func TestMerge(t *testing.T) {
	json1 := `{"name":"test","age":30}`
	json2 := `{"city":"NYC","age":31}`

	result, err := Merge(json1, json2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	name, _ := GetString(result, "name")
	if name != "test" {
		t.Fatalf("expected name 'test', got '%s'", name)
	}

	city, _ := GetString(result, "city")
	if city != "NYC" {
		t.Fatalf("expected city 'NYC', got '%s'", city)
	}

	age, _ := GetInt(result, "age")
	if age != 31 {
		t.Fatalf("expected age 31 (from json2), got %d", age)
	}
}

func TestMergeNested(t *testing.T) {
	json1 := `{"user":{"name":"test","age":30}}`
	json2 := `{"user":{"city":"NYC"}}`

	result, err := Merge(json1, json2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	name, _ := GetString(result, "user.name")
	if name != "test" {
		t.Fatalf("expected name 'test', got '%s'", name)
	}

	city, _ := GetString(result, "user.city")
	if city != "NYC" {
		t.Fatalf("expected city 'NYC', got '%s'", city)
	}
}

func TestFormat(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`

	result, err := Format(jsonStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(result, "\n") {
		t.Fatal("expected newlines in formatted JSON")
	}

	if !strings.Contains(result, "  ") {
		t.Fatal("expected indentation in formatted JSON")
	}
}

func TestCompact(t *testing.T) {
	jsonStr := `{
  "name": "test",
  "age": 30
}`

	result, err := Compact(jsonStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if strings.Contains(result, "\n") {
		t.Fatal("expected no newlines in compacted JSON")
	}
}

func TestKeys(t *testing.T) {
	jsonStr := `{"name":"test","age":30,"city":"NYC"}`

	keys, err := Keys(jsonStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(keys) != 3 {
		t.Fatalf("expected 3 keys, got %d", len(keys))
	}

	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	if !keyMap["name"] || !keyMap["age"] || !keyMap["city"] {
		t.Fatal("expected keys 'name', 'age', 'city'")
	}
}

func TestValues(t *testing.T) {
	jsonStr := `{"a":1,"b":2,"c":3}`

	values, err := Values(jsonStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(values) != 3 {
		t.Fatalf("expected 3 values, got %d", len(values))
	}
}

func TestContains(t *testing.T) {
	jsonStr := `{"user":{"name":"test"}}`

	if !Contains(jsonStr, "user.name") {
		t.Fatal("expected Contains to return true for 'user.name'")
	}

	if Contains(jsonStr, "user.age") {
		t.Fatal("expected Contains to return false for 'user.age'")
	}
}

func TestWrap(t *testing.T) {
	result, err := Wrap("data", "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	val, _ := GetString(result, "data")
	if val != "hello" {
		t.Fatalf("expected 'hello', got '%s'", val)
	}
}

func TestWrapMap(t *testing.T) {
	m := map[string]interface{}{
		"name": "test",
		"age":  30,
	}

	result, err := WrapMap(m)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	name, _ := GetString(result, "name")
	if name != "test" {
		t.Fatalf("expected 'test', got '%s'", name)
	}

	age, _ := GetInt(result, "age")
	if age != 30 {
		t.Fatalf("expected 30, got %d", age)
	}
}

func TestMarshalError(t *testing.T) {
	// Test with a channel which cannot be marshaled
	ch := make(chan int)
	_, err := Marshal(ch)
	if err == nil {
		t.Fatal("expected error for unmarshalable type")
	}
}

func TestUnmarshalError(t *testing.T) {
	var data map[string]interface{}
	err := Unmarshal([]byte(`{invalid}`), &data)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestGetInvalidJSON(t *testing.T) {
	_, err := Get(`{invalid}`, "name")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestSetInvalidJSON(t *testing.T) {
	_, err := Set(`{invalid}`, "name", "test")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestDeleteInvalidJSON(t *testing.T) {
	_, err := Delete(`{invalid}`, "name")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestMergeInvalidJSON(t *testing.T) {
	_, err := Merge(`{invalid}`, `{"name":"test"}`)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestFormatInvalidJSON(t *testing.T) {
	_, err := Format(`{invalid}`)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestToMapInvalidJSON(t *testing.T) {
	_, err := ToMap(`{invalid}`)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestToSliceInvalidJSON(t *testing.T) {
	_, err := ToSlice(`[invalid`)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestGetArrayOutOfBounds(t *testing.T) {
	jsonStr := `{"items":[1,2,3]}`

	_, err := Get(jsonStr, "items.5")
	if err == nil {
		t.Fatal("expected error for out of bounds index")
	}
}

func TestGetNegativeIndex(t *testing.T) {
	jsonStr := `{"items":[1,2,3]}`

	_, err := Get(jsonStr, "items.-1")
	if err == nil {
		t.Fatal("expected error for negative index")
	}
}

func TestSetNewNestedPath(t *testing.T) {
	jsonStr := `{"user":{"name":"test"}}`

	result, err := Set(jsonStr, "user.address.city", "NYC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	city, _ := GetString(result, "user.address.city")
	if city != "NYC" {
		t.Fatalf("expected 'NYC', got '%s'", city)
	}
}
