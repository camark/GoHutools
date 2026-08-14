package setting

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewSetting(t *testing.T) {
	s := New()
	if s == nil {
		t.Fatal("New() returned nil")
	}
	if len(s.Sections()) != 0 {
		t.Errorf("New setting should have no sections, got %v", s.Sections())
	}
}

func TestSetGet(t *testing.T) {
	s := New()
	s.Set("key1", "value1")
	s.Set("key2", "value2")

	if got := s.Get("key1"); got != "value1" {
		t.Errorf("Get(key1) = %v, want value1", got)
	}
	if got := s.Get("key2"); got != "value2" {
		t.Errorf("Get(key2) = %v, want value2", got)
	}
	if got := s.Get("nonexistent"); got != "" {
		t.Errorf("Get(nonexistent) = %v, want empty string", got)
	}
}

func TestGetWithDefault(t *testing.T) {
	s := New()
	s.Set("key1", "value1")

	if got := s.GetWithDefault("key1", "default"); got != "value1" {
		t.Errorf("GetWithDefault(key1) = %v, want value1", got)
	}
	if got := s.GetWithDefault("nonexistent", "default"); got != "default" {
		t.Errorf("GetWithDefault(nonexistent) = %v, want default", got)
	}
}

func TestSectionOperations(t *testing.T) {
	s := New()
	s.SetSection("database", "host", "localhost")
	s.SetSection("database", "port", "5432")
	s.SetSection("server", "port", "8080")

	if got := s.GetSection("database", "host"); got != "localhost" {
		t.Errorf("GetSection(database, host) = %v, want localhost", got)
	}
	if got := s.GetSection("database", "port"); got != "5432" {
		t.Errorf("GetSection(database, port) = %v, want 5432", got)
	}
	if got := s.GetSection("server", "port"); got != "8080" {
		t.Errorf("GetSection(server, port) = %v, want 8080", got)
	}
}

func TestGetSectionWithDefault(t *testing.T) {
	s := New()
	s.SetSection("database", "host", "localhost")

	if got := s.GetSectionWithDefault("database", "host", "default"); got != "localhost" {
		t.Errorf("GetSectionWithDefault(database, host) = %v, want localhost", got)
	}
	if got := s.GetSectionWithDefault("database", "port", "3306"); got != "3306" {
		t.Errorf("GetSectionWithDefault(database, port) = %v, want 3306", got)
	}
}

func TestHas(t *testing.T) {
	s := New()
	s.Set("key1", "value1")

	if !s.Has("key1") {
		t.Error("Has(key1) should be true")
	}
	if s.Has("nonexistent") {
		t.Error("Has(nonexistent) should be false")
	}
}

func TestHasSection(t *testing.T) {
	s := New()
	s.SetSection("database", "host", "localhost")

	if !s.HasSection("database") {
		t.Error("HasSection(database) should be true")
	}
	if s.HasSection("nonexistent") {
		t.Error("HasSection(nonexistent) should be false")
	}
}

func TestDelete(t *testing.T) {
	s := New()
	s.Set("key1", "value1")
	s.Delete("key1")

	if s.Has("key1") {
		t.Error("key1 should be deleted")
	}
}

func TestDeleteSection(t *testing.T) {
	s := New()
	s.SetSection("database", "host", "localhost")
	s.DeleteSection("database")

	if s.HasSection("database") {
		t.Error("database section should be deleted")
	}
}

func TestSections(t *testing.T) {
	s := New()
	s.Set("key1", "value1")
	s.SetSection("database", "host", "localhost")
	s.SetSection("server", "port", "8080")
	s.SetSection("app", "name", "test")

	sections := s.Sections()
	if len(sections) != 4 { // default + 3 sections
		t.Errorf("Expected 4 sections, got %d: %v", len(sections), sections)
	}

	// Check if sections are sorted
	for i := 1; i < len(sections); i++ {
		if sections[i-1] > sections[i] {
			t.Errorf("Sections not sorted: %v", sections)
			break
		}
	}
}

func TestKeys(t *testing.T) {
	s := New()
	s.SetSection("database", "host", "localhost")
	s.SetSection("database", "port", "5432")
	s.SetSection("database", "name", "testdb")

	keys := s.Keys("database")
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Check if keys are sorted
	for i := 1; i < len(keys); i++ {
		if keys[i-1] > keys[i] {
			t.Errorf("Keys not sorted: %v", keys)
			break
		}
	}
}

func TestToMap(t *testing.T) {
	s := New()
	s.Set("key1", "value1")
	s.SetSection("database", "host", "localhost")

	m := s.ToMap()
	if m["default"]["key1"] != "value1" {
		t.Errorf("ToMap() default.key1 = %v, want value1", m["default"]["key1"])
	}
	if m["database"]["host"] != "localhost" {
		t.Errorf("ToMap() database.host = %v, want localhost", m["database"]["host"])
	}
}

func TestLoadString(t *testing.T) {
	input := `
key1 = value1
key2 = value2

[database]
host = localhost
port = 5432
name = testdb
`

	s, err := LoadString(input)
	if err != nil {
		t.Fatalf("LoadString() error: %v", err)
	}

	if got := s.Get("key1"); got != "value1" {
		t.Errorf("Get(key1) = %v, want value1", got)
	}
	if got := s.Get("key2"); got != "value2" {
		t.Errorf("Get(key2) = %v, want value2", got)
	}
	if got := s.GetSection("database", "host"); got != "localhost" {
		t.Errorf("GetSection(database, host) = %v, want localhost", got)
	}
	if got := s.GetSection("database", "port"); got != "5432" {
		t.Errorf("GetSection(database, port) = %v, want 5432", got)
	}
}

func TestLoadStringWithComments(t *testing.T) {
	input := `
# This is a comment
key1 = value1
; This is also a comment
key2 = value2

[database]
# Database configuration
host = localhost
port = 5432
`

	s, err := LoadString(input)
	if err != nil {
		t.Fatalf("LoadString() error: %v", err)
	}

	if got := s.Get("key1"); got != "value1" {
		t.Errorf("Get(key1) = %v, want value1", got)
	}
	if got := s.GetSection("database", "host"); got != "localhost" {
		t.Errorf("GetSection(database, host) = %v, want localhost", got)
	}
}

func TestLoadStringWithQuotes(t *testing.T) {
	input := `
key1 = "quoted value"
key2 = 'single quoted'
key3 = unquoted
`

	s, err := LoadString(input)
	if err != nil {
		t.Fatalf("LoadString() error: %v", err)
	}

	if got := s.Get("key1"); got != "quoted value" {
		t.Errorf("Get(key1) = %v, want 'quoted value'", got)
	}
	if got := s.Get("key2"); got != "single quoted" {
		t.Errorf("Get(key2) = %v, want 'single quoted'", got)
	}
	if got := s.Get("key3"); got != "unquoted" {
		t.Errorf("Get(key3) = %v, want 'unquoted'", got)
	}
}

func TestLoadFile(t *testing.T) {
	// Create temporary file
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.ini")

	content := `
key1 = value1

[database]
host = localhost
port = 5432
`

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	s, err := LoadFile(filePath)
	if err != nil {
		t.Fatalf("LoadFile() error: %v", err)
	}

	if got := s.Get("key1"); got != "value1" {
		t.Errorf("Get(key1) = %v, want value1", got)
	}
	if got := s.GetSection("database", "host"); got != "localhost" {
		t.Errorf("GetSection(database, host) = %v, want localhost", got)
	}
}

func TestLoadFileNotFound(t *testing.T) {
	_, err := LoadFile("nonexistent.ini")
	if err == nil {
		t.Error("LoadFile() should return error for nonexistent file")
	}
}

func TestSaveString(t *testing.T) {
	s := New()
	s.Set("key1", "value1")
	s.Set("key2", "value2")
	s.SetSection("database", "host", "localhost")
	s.SetSection("database", "port", "5432")

	output := s.SaveString()

	// Check if output contains expected content
	if len(output) == 0 {
		t.Error("SaveString() returned empty string")
	}

	// Reload and verify
	s2, err := LoadString(output)
	if err != nil {
		t.Fatalf("LoadString() error: %v", err)
	}

	if got := s2.Get("key1"); got != "value1" {
		t.Errorf("Get(key1) = %v, want value1", got)
	}
	if got := s2.GetSection("database", "host"); got != "localhost" {
		t.Errorf("GetSection(database, host) = %v, want localhost", got)
	}
}

func TestSaveFile(t *testing.T) {
	s := New()
	s.Set("key1", "value1")
	s.SetSection("database", "host", "localhost")

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.ini")

	if err := s.Save(filePath); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	// Load and verify
	s2, err := LoadFile(filePath)
	if err != nil {
		t.Fatalf("LoadFile() error: %v", err)
	}

	if got := s2.Get("key1"); got != "value1" {
		t.Errorf("Get(key1) = %v, want value1", got)
	}
	if got := s2.GetSection("database", "host"); got != "localhost" {
		t.Errorf("GetSection(database, host) = %v, want localhost", got)
	}
}

func TestReload(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.ini")

	// Create initial file
	content1 := "key1 = value1"
	if err := os.WriteFile(filePath, []byte(content1), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	s, err := LoadFile(filePath)
	if err != nil {
		t.Fatalf("LoadFile() error: %v", err)
	}

	if got := s.Get("key1"); got != "value1" {
		t.Errorf("Get(key1) = %v, want value1", got)
	}

	// Update file
	content2 := "key2 = value2"
	if err := os.WriteFile(filePath, []byte(content2), 0644); err != nil {
		t.Fatalf("Failed to update test file: %v", err)
	}

	// Reload
	if err := s.Reload(filePath); err != nil {
		t.Fatalf("Reload() error: %v", err)
	}

	// Old key should be gone
	if s.Has("key1") {
		t.Error("key1 should not exist after reload")
	}

	// New key should exist
	if got := s.Get("key2"); got != "value2" {
		t.Errorf("Get(key2) = %v, want value2", got)
	}
}

func TestGetInt(t *testing.T) {
	s := New()
	s.Set("port", "8080")
	s.Set("invalid", "abc")

	if got, err := s.GetInt("port"); err != nil || got != 8080 {
		t.Errorf("GetInt(port) = %v, %v, want 8080, nil", got, err)
	}

	if _, err := s.GetInt("invalid"); err == nil {
		t.Error("GetInt(invalid) should return error")
	}

	if _, err := s.GetInt("nonexistent"); err == nil {
		t.Error("GetInt(nonexistent) should return error")
	}
}

func TestGetIntWithDefault(t *testing.T) {
	s := New()
	s.Set("port", "8080")

	if got := s.GetIntWithDefault("port", 3000); got != 8080 {
		t.Errorf("GetIntWithDefault(port) = %v, want 8080", got)
	}

	if got := s.GetIntWithDefault("nonexistent", 3000); got != 3000 {
		t.Errorf("GetIntWithDefault(nonexistent) = %v, want 3000", got)
	}
}

func TestGetBool(t *testing.T) {
	s := New()
	s.Set("true1", "true")
	s.Set("true2", "yes")
	s.Set("true3", "1")
	s.Set("false1", "false")
	s.Set("false2", "no")
	s.Set("false3", "0")
	s.Set("invalid", "abc")

	tests := []struct {
		key      string
		expected bool
	}{
		{"true1", true},
		{"true2", true},
		{"true3", true},
		{"false1", false},
		{"false2", false},
		{"false3", false},
	}

	for _, tt := range tests {
		got, err := s.GetBool(tt.key)
		if err != nil {
			t.Errorf("GetBool(%s) error: %v", tt.key, err)
			continue
		}
		if got != tt.expected {
			t.Errorf("GetBool(%s) = %v, want %v", tt.key, got, tt.expected)
		}
	}

	if _, err := s.GetBool("invalid"); err == nil {
		t.Error("GetBool(invalid) should return error")
	}
}

func TestGetBoolWithDefault(t *testing.T) {
	s := New()
	s.Set("enabled", "true")

	if got := s.GetBoolWithDefault("enabled", false); got != true {
		t.Errorf("GetBoolWithDefault(enabled) = %v, want true", got)
	}

	if got := s.GetBoolWithDefault("nonexistent", true); got != true {
		t.Errorf("GetBoolWithDefault(nonexistent) = %v, want true", got)
	}
}

func TestGetFloat(t *testing.T) {
	s := New()
	s.Set("pi", "3.14159")
	s.Set("invalid", "abc")

	if got, err := s.GetFloat("pi"); err != nil || got != 3.14159 {
		t.Errorf("GetFloat(pi) = %v, %v, want 3.14159, nil", got, err)
	}

	if _, err := s.GetFloat("invalid"); err == nil {
		t.Error("GetFloat(invalid) should return error")
	}
}

func TestGetFloatWithDefault(t *testing.T) {
	s := New()
	s.Set("pi", "3.14159")

	if got := s.GetFloatWithDefault("pi", 0); got != 3.14159 {
		t.Errorf("GetFloatWithDefault(pi) = %v, want 3.14159", got)
	}

	if got := s.GetFloatWithDefault("nonexistent", 1.0); got != 1.0 {
		t.Errorf("GetFloatWithDefault(nonexistent) = %v, want 1.0", got)
	}
}

func TestGetSlice(t *testing.T) {
	s := New()
	s.Set("colors", "red, green, blue")
	s.Set("single", "only")
	s.Set("empty", "")
	s.Set("spaces", "  a , b  ,  c  ")

	tests := []struct {
		key      string
		expected []string
	}{
		{"colors", []string{"red", "green", "blue"}},
		{"single", []string{"only"}},
		{"empty", []string{}},
		{"spaces", []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		got := s.GetSlice(tt.key)
		if len(got) != len(tt.expected) {
			t.Errorf("GetSlice(%s) length = %d, want %d", tt.key, len(got), len(tt.expected))
			continue
		}
		for i, v := range got {
			if v != tt.expected[i] {
				t.Errorf("GetSlice(%s)[%d] = %v, want %v", tt.key, i, v, tt.expected[i])
			}
		}
	}

	// Test nonexistent key
	if got := s.GetSlice("nonexistent"); len(got) != 0 {
		t.Errorf("GetSlice(nonexistent) = %v, want empty slice", got)
	}
}

func TestMapCopy(t *testing.T) {
	s := New()
	s.Set("key1", "value1")
	s.SetSection("section1", "key2", "value2")

	m := s.ToMap()

	// Modify the copy should not affect original
	m["default"]["key1"] = "modified"
	m["section1"]["key2"] = "modified"

	if got := s.Get("key1"); got != "value1" {
		t.Errorf("Original value changed: Get(key1) = %v, want value1", got)
	}
	if got := s.GetSection("section1", "key2"); got != "value2" {
		t.Errorf("Original value changed: GetSection(section1, key2) = %v, want value2", got)
	}
}
