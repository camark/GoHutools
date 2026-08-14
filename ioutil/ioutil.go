package ioutil

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// StringWriter writes strings
type StringWriter interface {
	WriteString(s string) (n int, err error)
}

// ReadAll reads all from reader
func ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

// ReadAllString reads all as string
func ReadAllString(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLines reads all lines from reader
func ReadLines(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// WriteAll writes all to writer
func WriteAll(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	return err
}

// WriteString writes string to writer
func WriteString(w io.Writer, s string) error {
	_, err := io.WriteString(w, s)
	return err
}

// WriteLines writes lines to writer
func WriteLines(w io.Writer, lines []string) error {
	for i, line := range lines {
		if _, err := io.WriteString(w, line); err != nil {
			return err
		}
		if i < len(lines)-1 {
			if _, err := io.WriteString(w, "\n"); err != nil {
				return err
			}
		}
	}
	return nil
}

// Copy copies from reader to writer
func Copy(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}

// CopyN copies n bytes from reader to writer
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
	return io.CopyN(dst, src, n)
}

// CopyBuffer copies with buffer
func CopyBuffer(dst io.Writer, src io.Reader, buf []byte) (int64, error) {
	return io.CopyBuffer(dst, src, buf)
}

// Drain reads and discards all from reader
func Drain(r io.Reader) error {
	_, err := io.Copy(io.Discard, r)
	return err
}

// ReadBytes reads until delimiter
func ReadBytes(r io.Reader, delim byte) ([]byte, error) {
	br := bufio.NewReader(r)
	return br.ReadBytes(delim)
}

// ReadString reads until delimiter
func ReadString(r io.Reader, delim byte) (string, error) {
	br := bufio.NewReader(r)
	s, err := br.ReadString(delim)
	if err != nil {
		return "", err
	}
	return s, nil
}

// TeeReader returns a Reader that writes to w what it reads from r
func TeeReader(r io.Reader, w io.Writer) io.Reader {
	return io.TeeReader(r, w)
}

// LimitReader returns a Reader that reads from r but stops with EOF after n bytes
func LimitReader(r io.Reader, n int64) io.Reader {
	return io.LimitReader(r, n)
}

// SectionReader returns a Reader that reads from r starting at offset off and stopping after n bytes
func SectionReader(r io.ReaderAt, off int64, n int64) io.Reader {
	return io.NewSectionReader(r, off, n)
}

// NewBuffer creates a new buffer
func NewBuffer(data []byte) *bytes.Buffer {
	return bytes.NewBuffer(data)
}

// NewBufferString creates a new buffer from string
func NewBufferString(s string) *bytes.Buffer {
	return bytes.NewBufferString(s)
}

// NewReader creates a new reader from bytes
func NewReader(data []byte) *bytes.Reader {
	return bytes.NewReader(data)
}

// NewReaderString creates a new reader from string
func NewReaderString(s string) *strings.Reader {
	return strings.NewReader(s)
}

// BufferedWriter returns a buffered writer
func BufferedWriter(w io.Writer) *bufio.Writer {
	return bufio.NewWriter(w)
}

// BufferedReadString returns a buffered reader
func BufferedReadString(r io.Reader) *bufio.Reader {
	return bufio.NewReader(r)
}

// Close safely closes closer
func Close(c io.Closer) error {
	if c == nil {
		return nil
	}
	return c.Close()
}

// CloseQuietly closes closer ignoring error
func CloseQuietly(c io.Closer) {
	if c != nil {
		c.Close()
	}
}

// Flush flushes writer if it implements Flush
func Flush(w io.Writer) error {
	if f, ok := w.(interface{ Flush() error }); ok {
		return f.Flush()
	}
	return fmt.Errorf("writer does not implement Flush")
}

// ReadAt reads from reader at offset
func ReadAt(r io.ReaderAt, b []byte, off int64) (int, error) {
	return r.ReadAt(b, off)
}

// WriteAt writes to writer at offset
func WriteAt(w io.WriterAt, b []byte, off int64) (int, error) {
	return w.WriteAt(b, off)
}

// MultiReader returns a Reader that's the logical concatenation of the provided readers
func MultiReader(readers ...io.Reader) io.Reader {
	return io.MultiReader(readers...)
}

// MultiWriter returns a Writer that duplicates its writes to all the provided writers
func MultiWriter(writers ...io.Writer) io.Writer {
	return io.MultiWriter(writers...)
}

// Pipe creates synchronous in-memory pipe
func Pipe() (*io.PipeReader, *io.PipeWriter) {
	return io.Pipe()
}

// ReadFull reads exactly len(buf) bytes from r into buf
func ReadFull(r io.Reader, buf []byte) (int, error) {
	return io.ReadFull(r, buf)
}

// ReadAtLeast reads from r into buf until it has read at least min bytes
func ReadAtLeast(r io.Reader, buf []byte, min int) (int, error) {
	return io.ReadAtLeast(r, buf, min)
}
