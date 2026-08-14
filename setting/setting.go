package setting

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const defaultSection = "default"

// Setting is configuration manager
type Setting struct {
	data    map[string]map[string]string
	section string
	mu      sync.RWMutex
}

// New creates new empty setting
func New() *Setting {
	return &Setting{
		data:    make(map[string]map[string]string),
		section: defaultSection,
	}
}

// LoadFile loads setting from file
func LoadFile(path string) (*Setting, error) {
	s := New()
	if err := s.Load(path); err != nil {
		return nil, err
	}
	return s, nil
}

// LoadString loads setting from string
func LoadString(str string) (*Setting, error) {
	s := New()
	if err := s.parse(strings.NewReader(str)); err != nil {
		return nil, err
	}
	return s, nil
}

// Get gets value by key in default section
func (s *Setting) Get(key string) string {
	return s.GetSection(defaultSection, key)
}

// GetWithDefault gets value with default
func (s *Setting) GetWithDefault(key, defaultVal string) string {
	return s.GetSectionWithDefault(defaultSection, key, defaultVal)
}

// GetSection gets value by section and key
func (s *Setting) GetSection(section, key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if sec, ok := s.data[section]; ok {
		if val, ok := sec[key]; ok {
			return val
		}
	}
	return ""
}

// GetSectionWithDefault gets value by section and key with default
func (s *Setting) GetSectionWithDefault(section, key, defaultVal string) string {
	val := s.GetSection(section, key)
	if val == "" {
		return defaultVal
	}
	return val
}

// Set sets value in default section
func (s *Setting) Set(key, value string) {
	s.SetSection(defaultSection, key, value)
}

// SetSection sets value in section
func (s *Setting) SetSection(section, key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[section]; !ok {
		s.data[section] = make(map[string]string)
	}
	s.data[section][key] = value
}

// Has checks if key exists in default section
func (s *Setting) Has(key string) bool {
	return s.HasSection(defaultSection) && s.Get(key) != ""
}

// HasSection checks if section exists
func (s *Setting) HasSection(section string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[section]
	return ok
}

// Delete deletes key from default section
func (s *Setting) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sec, ok := s.data[defaultSection]; ok {
		delete(sec, key)
	}
}

// DeleteSection deletes section
func (s *Setting) DeleteSection(section string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, section)
}

// Sections returns all sections
func (s *Setting) Sections() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	sections := make([]string, 0, len(s.data))
	for section := range s.data {
		sections = append(sections, section)
	}
	sort.Strings(sections)
	return sections
}

// Keys returns all keys in section
func (s *Setting) Keys(section string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]string, 0)
	if sec, ok := s.data[section]; ok {
		for key := range sec {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	return keys
}

// ToMap converts to map
func (s *Setting) ToMap() map[string]map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string]map[string]string)
	for section, sec := range s.data {
		result[section] = make(map[string]string)
		for key, val := range sec {
			result[section][key] = val
		}
	}
	return result
}

// Save saves setting to file
func (s *Setting) Save(path string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	writer := bufio.NewWriter(file)

	// Write default section first
	if sec, ok := s.data[defaultSection]; ok {
		keys := make([]string, 0, len(sec))
		for key := range sec {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Fprintf(writer, "%s = %s\n", key, sec[key])
		}
	}

	// Write other sections
	sections := make([]string, 0, len(s.data))
	for section := range s.data {
		if section != defaultSection {
			sections = append(sections, section)
		}
	}
	sort.Strings(sections)

	for _, section := range sections {
		fmt.Fprintf(writer, "\n[%s]\n", section)
		keys := make([]string, 0, len(s.data[section]))
		for key := range s.data[section] {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Fprintf(writer, "%s = %s\n", key, s.data[section][key])
		}
	}

	if err := writer.Flush(); err != nil {
		_ = file.Close()
		return fmt.Errorf("failed to flush file: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}
	return nil
}

// SaveString converts setting to string
func (s *Setting) SaveString() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result strings.Builder

	// Write default section first
	if sec, ok := s.data[defaultSection]; ok {
		keys := make([]string, 0, len(sec))
		for key := range sec {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Fprintf(&result, "%s = %s\n", key, sec[key])
		}
	}

	// Write other sections
	sections := make([]string, 0, len(s.data))
	for section := range s.data {
		if section != defaultSection {
			sections = append(sections, section)
		}
	}
	sort.Strings(sections)

	for _, section := range sections {
		fmt.Fprintf(&result, "\n[%s]\n", section)
		keys := make([]string, 0, len(s.data[section]))
		for key := range s.data[section] {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Fprintf(&result, "%s = %s\n", key, s.data[section][key])
		}
	}

	return result.String()
}

// Load loads setting from file (alias for LoadFile)
func (s *Setting) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() { _ = file.Close() }()

	return s.parse(file)
}

// Reload reloads setting from file
func (s *Setting) Reload(path string) error {
	s.mu.Lock()
	s.data = make(map[string]map[string]string)
	s.mu.Unlock()

	return s.Load(path)
}

func (s *Setting) parse(reader io.Reader) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	scanner := bufio.NewScanner(reader)
	currentSection := defaultSection

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		// Check for section header
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = line[1 : len(line)-1]
			currentSection = strings.TrimSpace(currentSection)
			if _, ok := s.data[currentSection]; !ok {
				s.data[currentSection] = make(map[string]string)
			}
			continue
		}

		// Parse key = value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			// Remove quotes if present
			if len(value) >= 2 {
				if (value[0] == '"' && value[len(value)-1] == '"') ||
					(value[0] == '\'' && value[len(value)-1] == '\'') {
					value = value[1 : len(value)-1]
				}
			}

			if _, ok := s.data[currentSection]; !ok {
				s.data[currentSection] = make(map[string]string)
			}
			s.data[currentSection][key] = value
		}
	}

	return scanner.Err()
}

// GetInt gets int value
func (s *Setting) GetInt(key string) (int, error) {
	val := s.Get(key)
	if val == "" {
		return 0, fmt.Errorf("key not found: %s", key)
	}
	return strconv.Atoi(val)
}

// GetIntWithDefault gets int with default
func (s *Setting) GetIntWithDefault(key string, defaultVal int) int {
	val, err := s.GetInt(key)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetBool gets bool value
func (s *Setting) GetBool(key string) (bool, error) {
	val := s.Get(key)
	if val == "" {
		return false, fmt.Errorf("key not found: %s", key)
	}
	return parseBool(val)
}

// GetBoolWithDefault gets bool with default
func (s *Setting) GetBoolWithDefault(key string, defaultVal bool) bool {
	val, err := s.GetBool(key)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetFloat gets float value
func (s *Setting) GetFloat(key string) (float64, error) {
	val := s.Get(key)
	if val == "" {
		return 0, fmt.Errorf("key not found: %s", key)
	}
	return strconv.ParseFloat(val, 64)
}

// GetFloatWithDefault gets float with default
func (s *Setting) GetFloatWithDefault(key string, defaultVal float64) float64 {
	val, err := s.GetFloat(key)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetSlice gets slice value (comma-separated)
func (s *Setting) GetSlice(key string) []string {
	val := s.Get(key)
	if val == "" {
		return []string{}
	}

	parts := strings.Split(val, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func parseBool(str string) (bool, error) {
	switch strings.ToLower(str) {
	case "1", "t", "true", "yes", "on":
		return true, nil
	case "0", "f", "false", "no", "off":
		return false, nil
	}
	return false, fmt.Errorf("invalid boolean value: %s", str)
}
