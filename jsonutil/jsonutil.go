package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Marshal serializes value to JSON bytes
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalIndent serializes value to indented JSON bytes
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// MarshalToString serializes value to JSON string
func MarshalToString(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Unmarshal deserializes JSON bytes to value
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// UnmarshalFromString deserializes JSON string to value
func UnmarshalFromString(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// UnmarshalFromReader deserializes JSON from reader
func UnmarshalFromReader(r io.Reader, v interface{}) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("read from reader failed: %w", err)
	}
	return json.Unmarshal(data, v)
}

// IsValidJSON checks if bytes is valid JSON
func IsValidJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// IsValidJSONString checks if string is valid JSON
func IsValidJSONString(s string) bool {
	return IsValidJSON([]byte(s))
}

// ToJSON converts value to JSON string
func ToJSON(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// ToPrettyJSON converts value to pretty JSON string
func ToPrettyJSON(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(bytes)
}

// FromJSON deserializes JSON string to value
func FromJSON(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// ToMap converts JSON to map
func ToMap(s string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// ToSlice converts JSON to slice
func ToSlice(s string) ([]interface{}, error) {
	var s2 []interface{}
	err := json.Unmarshal([]byte(s), &s2)
	if err != nil {
		return nil, err
	}
	return s2, nil
}

// Get gets value by path (e.g., "a.b.c")
func Get(jsonStr string, path string) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	keys := strings.Split(path, ".")
	current := data

	for _, key := range keys {
		switch v := current.(type) {
		case map[string]interface{}:
			val, ok := v[key]
			if !ok {
				return nil, fmt.Errorf("key '%s' not found", key)
			}
			current = val
		case []interface{}:
			// Try to parse as array index
			var idx int
			_, err := fmt.Sscanf(key, "%d", &idx)
			if err != nil {
				return nil, fmt.Errorf("invalid array index '%s'", key)
			}
			if idx < 0 || idx >= len(v) {
				return nil, fmt.Errorf("array index %d out of range", idx)
			}
			current = v[idx]
		default:
			return nil, fmt.Errorf("cannot traverse into %T", current)
		}
	}

	return current, nil
}

// GetString gets string by path
func GetString(jsonStr string, path string) (string, error) {
	val, err := Get(jsonStr, path)
	if err != nil {
		return "", err
	}
	s, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("value at path '%s' is not a string", path)
	}
	return s, nil
}

// GetInt gets int by path
func GetInt(jsonStr string, path string) (int, error) {
	val, err := Get(jsonStr, path)
	if err != nil {
		return 0, err
	}
	switch v := val.(type) {
	case float64:
		return int(v), nil
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case json.Number:
		i, err := v.Int64()
		if err != nil {
			return 0, err
		}
		return int(i), nil
	default:
		return 0, fmt.Errorf("value at path '%s' is not a number", path)
	}
}

// GetFloat gets float64 by path
func GetFloat(jsonStr string, path string) (float64, error) {
	val, err := Get(jsonStr, path)
	if err != nil {
		return 0, err
	}
	switch v := val.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case json.Number:
		f, err := v.Float64()
		if err != nil {
			return 0, err
		}
		return f, nil
	default:
		return 0, fmt.Errorf("value at path '%s' is not a number", path)
	}
}

// GetBool gets bool by path
func GetBool(jsonStr string, path string) (bool, error) {
	val, err := Get(jsonStr, path)
	if err != nil {
		return false, err
	}
	b, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("value at path '%s' is not a bool", path)
	}
	return b, nil
}

// Set sets value by path
func Set(jsonStr string, path string, value interface{}) (string, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	keys := strings.Split(path, ".")
	data = setNestedValue(data, keys, value)

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// setNestedValue recursively sets nested value
func setNestedValue(data interface{}, keys []string, value interface{}) interface{} {
	if len(keys) == 0 {
		return value
	}

	key := keys[0]
	remaining := keys[1:]

	switch v := data.(type) {
	case map[string]interface{}:
		if len(remaining) == 0 {
			v[key] = value
		} else {
			if _, ok := v[key]; !ok {
				v[key] = make(map[string]interface{})
			}
			v[key] = setNestedValue(v[key], remaining, value)
		}
		return v
	case []interface{}:
		var idx int
		_, err := fmt.Sscanf(key, "%d", &idx)
		if err == nil && idx >= 0 && idx < len(v) {
			v[idx] = setNestedValue(v[idx], remaining, value)
		}
		return v
	default:
		// Create new map
		m := make(map[string]interface{})
		m[key] = setNestedValue(make(map[string]interface{}), remaining, value)
		return m
	}
}

// Delete deletes value by path
func Delete(jsonStr string, path string) (string, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	keys := strings.Split(path, ".")
	data = deleteNestedValue(data, keys)

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// deleteNestedValue recursively deletes nested value
func deleteNestedValue(data interface{}, keys []string) interface{} {
	if len(keys) == 0 {
		return nil
	}

	key := keys[0]
	remaining := keys[1:]

	switch v := data.(type) {
	case map[string]interface{}:
		if len(remaining) == 0 {
			delete(v, key)
		} else {
			if val, ok := v[key]; ok {
				v[key] = deleteNestedValue(val, remaining)
			}
		}
		return v
	case []interface{}:
		var idx int
		_, err := fmt.Sscanf(key, "%d", &idx)
		if err == nil && idx >= 0 && idx < len(v) {
			if len(remaining) == 0 {
				// Remove element from slice
				v = append(v[:idx], v[idx+1:]...)
			} else {
				v[idx] = deleteNestedValue(v[idx], remaining)
			}
		}
		return v
	default:
		return data
	}
}

// Merge merges two JSON objects
func Merge(json1, json2 string) (string, error) {
	var m1, m2 map[string]interface{}
	if err := json.Unmarshal([]byte(json1), &m1); err != nil {
		return "", fmt.Errorf("invalid first JSON: %w", err)
	}
	if err := json.Unmarshal([]byte(json2), &m2); err != nil {
		return "", fmt.Errorf("invalid second JSON: %w", err)
	}

	merged := mergeMaps(m1, m2)

	bytes, err := json.Marshal(merged)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// mergeMaps recursively merges two maps
func mergeMaps(m1, m2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Copy m1
	for k, v := range m1 {
		result[k] = v
	}

	// Merge m2
	for k, v := range m2 {
		if existing, ok := result[k]; ok {
			// If both are maps, merge recursively
			if existingMap, ok1 := existing.(map[string]interface{}); ok1 {
				if newMap, ok2 := v.(map[string]interface{}); ok2 {
					result[k] = mergeMaps(existingMap, newMap)
					continue
				}
			}
		}
		result[k] = v
	}

	return result
}

// Format formats JSON string
func Format(jsonStr string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Compact compacts JSON string
func Compact(jsonStr string) (string, error) {
	var buf bytes.Buffer
	err := json.Compact(&buf, []byte(jsonStr))
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Keys returns top-level keys
func Keys(jsonStr string) ([]string, error) {
	m, err := ToMap(jsonStr)
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys, nil
}

// Values returns top-level values
func Values(jsonStr string) ([]interface{}, error) {
	m, err := ToMap(jsonStr)
	if err != nil {
		return nil, err
	}

	values := make([]interface{}, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values, nil
}

// Contains checks if JSON contains value at path
func Contains(jsonStr string, path string) bool {
	_, err := Get(jsonStr, path)
	return err == nil
}

// Wrap wraps value in JSON object with key
func Wrap(key string, value interface{}) (string, error) {
	m := map[string]interface{}{key: value}
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// WrapMap wraps map in JSON object
func WrapMap(m map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
