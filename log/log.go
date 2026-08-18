package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Level represents log level
type Level int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelOff
)

// String returns level string
func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelOff:
		return "OFF"
	default:
		return "UNKNOWN"
	}
}

// Logger is logger interface
type Logger interface {
	Trace(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	SetLevel(level Level)
	GetLevel() Level
}

// defaultLogger is default logger implementation
type defaultLogger struct {
	level  Level
	output io.Writer
	mu     sync.Mutex
}

// New creates new logger
func New() Logger {
	return &defaultLogger{
		level:  LevelDebug,
		output: os.Stderr,
	}
}

// NewWithOutput creates new logger with output
func NewWithOutput(output io.Writer) Logger {
	return &defaultLogger{
		level:  LevelDebug,
		output: output,
	}
}

// NewWithLevel creates new logger with level
func NewWithLevel(level Level) Logger {
	return &defaultLogger{
		level:  level,
		output: os.Stderr,
	}
}

func (l *defaultLogger) log(level Level, format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if level < l.level {
		return
	}

	// Get caller information
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}

	// Format timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")

	// Format message
	message := fmt.Sprintf(format, args...)

	// Format output
	output := fmt.Sprintf("[%s] [%s] %s:%d: %s\n", timestamp, level, file, line, message)

	// Write to output
	fmt.Fprint(l.output, output)
}

func (l *defaultLogger) Trace(format string, args ...interface{}) {
	l.log(LevelTrace, format, args...)
}

func (l *defaultLogger) Debug(format string, args ...interface{}) {
	l.log(LevelDebug, format, args...)
}

func (l *defaultLogger) Info(format string, args ...interface{}) {
	l.log(LevelInfo, format, args...)
}

func (l *defaultLogger) Warn(format string, args ...interface{}) {
	l.log(LevelWarn, format, args...)
}

func (l *defaultLogger) Error(format string, args ...interface{}) {
	l.log(LevelError, format, args...)
}

func (l *defaultLogger) Fatal(format string, args ...interface{}) {
	l.log(LevelFatal, format, args...)
	os.Exit(1)
}

func (l *defaultLogger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *defaultLogger) GetLevel() Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

// Global logger functions
var defaultLoggerInstance atomic.Pointer[Logger]

func init() {
	l := New()
	defaultLoggerInstance.Store(&l)
}

// Trace logs at trace level
func Trace(format string, args ...interface{}) {
	(*defaultLoggerInstance.Load()).Trace(format, args...)
}

// Debug logs at debug level
func Debug(format string, args ...interface{}) {
	(*defaultLoggerInstance.Load()).Debug(format, args...)
}

// Info logs at info level
func Info(format string, args ...interface{}) {
	(*defaultLoggerInstance.Load()).Info(format, args...)
}

// Warn logs at warn level
func Warn(format string, args ...interface{}) {
	(*defaultLoggerInstance.Load()).Warn(format, args...)
}

// Error logs at error level
func Error(format string, args ...interface{}) {
	(*defaultLoggerInstance.Load()).Error(format, args...)
}

// Fatal logs at fatal level and exits
func Fatal(format string, args ...interface{}) {
	(*defaultLoggerInstance.Load()).Fatal(format, args...)
}

// SetLevel sets the global logger level
func SetLevel(level Level) {
	(*defaultLoggerInstance.Load()).SetLevel(level)
}

// GetLevel gets the global logger level
func GetLevel() Level {
	return (*defaultLoggerInstance.Load()).GetLevel()
}

// SetOutput sets the global logger output
func SetOutput(w io.Writer) {
	l := NewWithOutput(w)
	defaultLoggerInstance.Store(&l)
}

// GetLogger returns the global logger
func GetLogger() Logger {
	return *defaultLoggerInstance.Load()
}

// SetLogger sets the global logger
func SetLogger(l Logger) {
	defaultLoggerInstance.Store(&l)
}

// WithFields creates logger with fields
func WithFields(fields map[string]interface{}) Logger {
	return &fieldLogger{
		logger: *defaultLoggerInstance.Load(),
		fields: fields,
	}
}

// fieldLogger is logger with fields
type fieldLogger struct {
	logger Logger
	fields map[string]interface{}
}

func (l *fieldLogger) formatWithFields(format string, args ...interface{}) string {
	message := fmt.Sprintf(format, args...)
	if len(l.fields) > 0 {
		fields := make([]string, 0, len(l.fields))
		for k, v := range l.fields {
			fields = append(fields, fmt.Sprintf("%s=%v", k, v))
		}
		return fmt.Sprintf("%s [%s]", message, strings.Join(fields, ", "))
	}
	return message
}

func (l *fieldLogger) Trace(format string, args ...interface{}) {
	l.logger.Trace(l.formatWithFields(format, args...))
}

func (l *fieldLogger) Debug(format string, args ...interface{}) {
	l.logger.Debug(l.formatWithFields(format, args...))
}

func (l *fieldLogger) Info(format string, args ...interface{}) {
	l.logger.Info(l.formatWithFields(format, args...))
}

func (l *fieldLogger) Warn(format string, args ...interface{}) {
	l.logger.Warn(l.formatWithFields(format, args...))
}

func (l *fieldLogger) Error(format string, args ...interface{}) {
	l.logger.Error(l.formatWithFields(format, args...))
}

func (l *fieldLogger) Fatal(format string, args ...interface{}) {
	l.logger.Fatal(l.formatWithFields(format, args...))
}

func (l *fieldLogger) SetLevel(level Level) {
	l.logger.SetLevel(level)
}

func (l *fieldLogger) GetLevel() Level {
	return l.logger.GetLevel()
}
