package strutil

import (
	"fmt"
	"strings"
)

// Builder is a string builder with fluent API
type Builder struct {
	buf strings.Builder
}

// NewBuilder creates a new Builder
func NewBuilder() *Builder {
	return &Builder{}
}

// NewBuilderSize creates a new Builder with initial capacity
func NewBuilderSize(size int) *Builder {
	b := &Builder{}
	if size > 0 {
		b.buf.Grow(size)
	}
	return b
}

// Append appends a string to the builder
func (b *Builder) Append(s string) *Builder {
	b.buf.WriteString(s)
	return b
}

// AppendByte appends a byte to the builder
func (b *Builder) AppendByte(c byte) *Builder {
	b.buf.WriteByte(c)
	return b
}

// AppendRune appends a rune to the builder
func (b *Builder) AppendRune(r rune) *Builder {
	b.buf.WriteRune(r)
	return b
}

// AppendFormat appends a formatted string to the builder
func (b *Builder) AppendFormat(format string, args ...interface{}) *Builder {
	b.buf.WriteString(fmt.Sprintf(format, args...))
	return b
}

// Insert inserts a string at the specified index
func (b *Builder) Insert(index int, s string) *Builder {
	if index < 0 {
		index = 0
	}

	content := b.buf.String()
	runes := []rune(content)
	if index > len(runes) {
		index = len(runes)
	}

	// Convert to bytes for manipulation
	bytes := []byte(content)
	insertBytes := []byte(s)

	// Find the byte position corresponding to the rune index
	bytePos := 0
	for i := 0; i < index && bytePos < len(bytes); i++ {
		// Skip UTF-8 bytes
		if bytes[bytePos] < 0x80 {
			bytePos++
		} else if bytes[bytePos] < 0xE0 {
			bytePos += 2
		} else if bytes[bytePos] < 0xF0 {
			bytePos += 3
		} else {
			bytePos += 4
		}
	}

	// Create new content: before + insert + after
	newBytes := make([]byte, 0, len(bytes)+len(insertBytes))
	newBytes = append(newBytes, bytes[:bytePos]...)
	newBytes = append(newBytes, insertBytes...)
	newBytes = append(newBytes, bytes[bytePos:]...)

	b.buf.Reset()
	b.buf.Write(newBytes)

	return b
}

// Delete deletes characters from start to end index
func (b *Builder) Delete(start, end int) *Builder {
	content := b.buf.String()
	runes := []rune(content)
	length := len(runes)

	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start > end {
		return b
	}

	// Find byte positions
	bytes := []byte(content)
	startByte := 0
	endByte := len(bytes)

	bytePos := 0
	for i := 0; i < length; i++ {
		if i == start {
			startByte = bytePos
		}
		if i == end {
			endByte = bytePos
			break
		}
		// Skip UTF-8 bytes
		if bytePos < len(bytes) {
			if bytes[bytePos] < 0x80 {
				bytePos++
			} else if bytes[bytePos] < 0xE0 {
				bytePos += 2
			} else if bytes[bytePos] < 0xF0 {
				bytePos += 3
			} else {
				bytePos += 4
			}
		}
	}

	// Create new content without the deleted range
	newBytes := make([]byte, 0, len(bytes)-(endByte-startByte))
	newBytes = append(newBytes, bytes[:startByte]...)
	newBytes = append(newBytes, bytes[endByte:]...)

	b.buf.Reset()
	b.buf.Write(newBytes)

	return b
}

// Replace replaces characters from start to end with string s
func (b *Builder) Replace(start, end int, s string) *Builder {
	content := b.buf.String()
	runes := []rune(content)
	length := len(runes)

	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start > end {
		return b
	}

	// Find byte positions
	bytes := []byte(content)
	startByte := 0
	endByte := len(bytes)

	bytePos := 0
	for i := 0; i < length; i++ {
		if i == start {
			startByte = bytePos
		}
		if i == end {
			endByte = bytePos
			break
		}
		// Skip UTF-8 bytes
		if bytePos < len(bytes) {
			if bytes[bytePos] < 0x80 {
				bytePos++
			} else if bytes[bytePos] < 0xE0 {
				bytePos += 2
			} else if bytes[bytePos] < 0xF0 {
				bytePos += 3
			} else {
				bytePos += 4
			}
		}
	}

	// Create new content with replacement
	newBytes := make([]byte, 0, len(bytes)-(endByte-startByte)+len(s))
	newBytes = append(newBytes, bytes[:startByte]...)
	newBytes = append(newBytes, []byte(s)...)
	newBytes = append(newBytes, bytes[endByte:]...)

	b.buf.Reset()
	b.buf.Write(newBytes)

	return b
}

// String returns the accumulated string
func (b *Builder) String() string {
	return b.buf.String()
}

// Len returns the number of accumulated bytes
func (b *Builder) Len() int {
	return b.buf.Len()
}

// Reset resets the builder
func (b *Builder) Reset() *Builder {
	b.buf.Reset()
	return b
}

// IsEmpty checks if the builder is empty
func (b *Builder) IsEmpty() bool {
	return b.buf.Len() == 0
}
