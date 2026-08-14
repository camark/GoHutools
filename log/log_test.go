package log

import (
	"bytes"
	"strings"
	"testing"
)

func TestLevelString(t *testing.T) {
	tests := []struct {
		level    Level
		expected string
	}{
		{LevelTrace, "TRACE"},
		{LevelDebug, "DEBUG"},
		{LevelInfo, "INFO"},
		{LevelWarn, "WARN"},
		{LevelError, "ERROR"},
		{LevelFatal, "FATAL"},
		{LevelOff, "OFF"},
		{Level(99), "UNKNOWN"},
	}

	for _, tt := range tests {
		if got := tt.level.String(); got != tt.expected {
			t.Errorf("Level(%d).String() = %v, want %v", tt.level, got, tt.expected)
		}
	}
}

func TestNewLogger(t *testing.T) {
	logger := New()
	if logger == nil {
		t.Fatal("New() returned nil")
	}
	if logger.GetLevel() != LevelDebug {
		t.Errorf("Default level = %v, want %v", logger.GetLevel(), LevelDebug)
	}
}

func TestNewWithOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := NewWithOutput(&buf)
	if logger == nil {
		t.Fatal("NewWithOutput() returned nil")
	}

	logger.Info("test message")
	if !strings.Contains(buf.String(), "test message") {
		t.Errorf("Output does not contain expected message: %s", buf.String())
	}
}

func TestNewWithLevel(t *testing.T) {
	logger := NewWithLevel(LevelWarn)
	if logger == nil {
		t.Fatal("NewWithLevel() returned nil")
	}
	if logger.GetLevel() != LevelWarn {
		t.Errorf("Level = %v, want %v", logger.GetLevel(), LevelWarn)
	}
}

func TestLoggerSetLevel(t *testing.T) {
	logger := New()
	logger.SetLevel(LevelError)
	if logger.GetLevel() != LevelError {
		t.Errorf("Level = %v, want %v", logger.GetLevel(), LevelError)
	}
}

func TestLoggerLogLevels(t *testing.T) {
	var buf bytes.Buffer
	logger := NewWithOutput(&buf)

	tests := []struct {
		name    string
		logFunc func(format string, args ...interface{})
		level   Level
		message string
	}{
		{"Trace", logger.Trace, LevelTrace, "trace message"},
		{"Debug", logger.Debug, LevelDebug, "debug message"},
		{"Info", logger.Info, LevelInfo, "info message"},
		{"Warn", logger.Warn, LevelWarn, "warn message"},
		{"Error", logger.Error, LevelError, "error message"},
	}

	logger.SetLevel(LevelTrace)

	for _, tt := range tests {
		buf.Reset()
		tt.logFunc(tt.message)
		output := buf.String()
		if !strings.Contains(output, tt.level.String()) {
			t.Errorf("%s: output does not contain level %s: %s", tt.name, tt.level, output)
		}
		if !strings.Contains(output, tt.message) {
			t.Errorf("%s: output does not contain message %s: %s", tt.name, tt.message, output)
		}
	}
}

func TestLoggerLevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	logger := NewWithOutput(&buf)

	logger.SetLevel(LevelWarn)

	logger.Debug("should not appear")
	if buf.Len() != 0 {
		t.Errorf("Debug message should not appear at WARN level: %s", buf.String())
	}

	logger.Warn("should appear")
	if !strings.Contains(buf.String(), "should appear") {
		t.Errorf("Warn message should appear: %s", buf.String())
	}
}

func TestLoggerFormat(t *testing.T) {
	var buf bytes.Buffer
	logger := NewWithOutput(&buf)

	logger.Info("hello %s, number %d", "world", 42)
	output := buf.String()

	if !strings.Contains(output, "hello world, number 42") {
		t.Errorf("Output does not contain formatted message: %s", output)
	}
	if !strings.Contains(output, "[INFO]") {
		t.Errorf("Output does not contain INFO level: %s", output)
	}
}

func TestGlobalLogger(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	Info("global info message")
	if !strings.Contains(buf.String(), "global info message") {
		t.Errorf("Global logger output does not contain message: %s", buf.String())
	}
}

func TestGlobalSetLevel(t *testing.T) {
	SetLevel(LevelWarn)
	if GetLevel() != LevelWarn {
		t.Errorf("Global level = %v, want %v", GetLevel(), LevelWarn)
	}
	SetLevel(LevelDebug) // Reset
}

func TestGetSetLogger(t *testing.T) {
	original := GetLogger()
	defer SetLogger(original)

	var buf bytes.Buffer
	newLogger := NewWithOutput(&buf)
	SetLogger(newLogger)

	Debug("test message")
	if !strings.Contains(buf.String(), "test message") {
		t.Errorf("New logger output does not contain message: %s", buf.String())
	}
}

func TestWithFields(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	logger := WithFields(map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	})

	logger.Info("test message")
	output := buf.String()

	if !strings.Contains(output, "test message") {
		t.Errorf("Output does not contain message: %s", output)
	}
	if !strings.Contains(output, "key1=value1") {
		t.Errorf("Output does not contain field key1=value1: %s", output)
	}
	if !strings.Contains(output, "key2=42") {
		t.Errorf("Output does not contain field key2=42: %s", output)
	}
}

func TestWithFieldsLevel(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	SetLevel(LevelWarn)

	logger := WithFields(map[string]interface{}{
		"key": "value",
	})

	logger.Debug("should not appear")
	if buf.Len() != 0 {
		t.Errorf("Debug message should not appear at WARN level: %s", buf.String())
	}

	logger.Warn("should appear")
	if !strings.Contains(buf.String(), "should appear") {
		t.Errorf("Warn message should appear: %s", buf.String())
	}
}

func TestFieldLoggerSetGetLevel(t *testing.T) {
	var buf bytes.Buffer
	_ = NewWithOutput(&buf)

	fieldLogger := WithFields(map[string]interface{}{
		"key": "value",
	})

	// Field logger delegates to underlying logger
	fieldLogger.SetLevel(LevelError)
	if fieldLogger.GetLevel() != LevelError {
		t.Errorf("Field logger level = %v, want %v", fieldLogger.GetLevel(), LevelError)
	}
}
