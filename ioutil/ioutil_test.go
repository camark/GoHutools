package ioutil

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	input := "hello world"
	r := strings.NewReader(input)
	data, err := ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll error: %v", err)
	}
	if string(data) != input {
		t.Errorf("ReadAll = %q, want %q", string(data), input)
	}
}

func TestReadAllString(t *testing.T) {
	input := "hello world"
	r := strings.NewReader(input)
	s, err := ReadAllString(r)
	if err != nil {
		t.Fatalf("ReadAllString error: %v", err)
	}
	if s != input {
		t.Errorf("ReadAllString = %q, want %q", s, input)
	}
}

func TestReadLines(t *testing.T) {
	input := "line1\nline2\nline3"
	r := strings.NewReader(input)
	lines, err := ReadLines(r)
	if err != nil {
		t.Fatalf("ReadLines error: %v", err)
	}
	expected := []string{"line1", "line2", "line3"}
	if len(lines) != len(expected) {
		t.Fatalf("ReadLines returned %d lines, want %d", len(lines), len(expected))
	}
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("line[%d] = %q, want %q", i, line, expected[i])
		}
	}
}

func TestReadLinesEmpty(t *testing.T) {
	r := strings.NewReader("")
	lines, err := ReadLines(r)
	if err != nil {
		t.Fatalf("ReadLines error: %v", err)
	}
	if len(lines) != 0 {
		t.Errorf("ReadLines returned %d lines, want 0", len(lines))
	}
}

func TestWriteAll(t *testing.T) {
	var buf bytes.Buffer
	data := []byte("hello world")
	err := WriteAll(&buf, data)
	if err != nil {
		t.Fatalf("WriteAll error: %v", err)
	}
	if !bytes.Equal(buf.Bytes(), data) {
		t.Errorf("WriteAll wrote %q, want %q", buf.Bytes(), data)
	}
}

func TestWriteString(t *testing.T) {
	var buf bytes.Buffer
	s := "hello world"
	err := WriteString(&buf, s)
	if err != nil {
		t.Fatalf("WriteString error: %v", err)
	}
	if buf.String() != s {
		t.Errorf("WriteString wrote %q, want %q", buf.String(), s)
	}
}

func TestWriteLines(t *testing.T) {
	var buf bytes.Buffer
	lines := []string{"line1", "line2", "line3"}
	err := WriteLines(&buf, lines)
	if err != nil {
		t.Fatalf("WriteLines error: %v", err)
	}
	expected := "line1\nline2\nline3"
	if buf.String() != expected {
		t.Errorf("WriteLines wrote %q, want %q", buf.String(), expected)
	}
}

func TestWriteLinesEmpty(t *testing.T) {
	var buf bytes.Buffer
	err := WriteLines(&buf, []string{})
	if err != nil {
		t.Fatalf("WriteLines error: %v", err)
	}
	if buf.Len() != 0 {
		t.Errorf("WriteLines wrote %q, want empty", buf.String())
	}
}

func TestCopy(t *testing.T) {
	src := strings.NewReader("hello world")
	var dst bytes.Buffer
	n, err := Copy(&dst, src)
	if err != nil {
		t.Fatalf("Copy error: %v", err)
	}
	if n != 11 {
		t.Errorf("Copy copied %d bytes, want 11", n)
	}
	if dst.String() != "hello world" {
		t.Errorf("Copy wrote %q, want %q", dst.String(), "hello world")
	}
}

func TestCopyN(t *testing.T) {
	src := strings.NewReader("hello world")
	var dst bytes.Buffer
	n, err := CopyN(&dst, src, 5)
	if err != nil {
		t.Fatalf("CopyN error: %v", err)
	}
	if n != 5 {
		t.Errorf("CopyN copied %d bytes, want 5", n)
	}
	if dst.String() != "hello" {
		t.Errorf("CopyN wrote %q, want %q", dst.String(), "hello")
	}
}

func TestCopyBuffer(t *testing.T) {
	src := strings.NewReader("hello world")
	var dst bytes.Buffer
	buf := make([]byte, 4)
	n, err := CopyBuffer(&dst, src, buf)
	if err != nil {
		t.Fatalf("CopyBuffer error: %v", err)
	}
	if n != 11 {
		t.Errorf("CopyBuffer copied %d bytes, want 11", n)
	}
	if dst.String() != "hello world" {
		t.Errorf("CopyBuffer wrote %q, want %q", dst.String(), "hello world")
	}
}

func TestDrain(t *testing.T) {
	r := strings.NewReader("hello world")
	err := Drain(r)
	if err != nil {
		t.Fatalf("Drain error: %v", err)
	}
	n, _ := r.Read(make([]byte, 1))
	if n != 0 {
		t.Errorf("Drain did not consume all data")
	}
}

func TestReadBytes(t *testing.T) {
	r := strings.NewReader("hello world foo")
	data, err := ReadBytes(r, ' ')
	if err != nil {
		t.Fatalf("ReadBytes error: %v", err)
	}
	if string(data) != "hello " {
		t.Errorf("ReadBytes = %q, want %q", string(data), "hello ")
	}
}

func TestReadString(t *testing.T) {
	r := strings.NewReader("hello world foo")
	s, err := ReadString(r, ' ')
	if err != nil {
		t.Fatalf("ReadString error: %v", err)
	}
	if s != "hello " {
		t.Errorf("ReadString = %q, want %q", s, "hello ")
	}
}

func TestTeeReader(t *testing.T) {
	src := strings.NewReader("hello world")
	var dst bytes.Buffer
	r := TeeReader(src, &dst)
	data, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("TeeReader ReadAll error: %v", err)
	}
	if string(data) != "hello world" {
		t.Errorf("TeeReader read %q, want %q", string(data), "hello world")
	}
	if dst.String() != "hello world" {
		t.Errorf("TeeReader wrote %q, want %q", dst.String(), "hello world")
	}
}

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("hello world")
	lr := LimitReader(r, 5)
	data, err := io.ReadAll(lr)
	if err != nil {
		t.Fatalf("LimitReader ReadAll error: %v", err)
	}
	if string(data) != "hello" {
		t.Errorf("LimitReader read %q, want %q", string(data), "hello")
	}
}

func TestSectionReader(t *testing.T) {
	data := []byte("hello world")
	r := bytes.NewReader(data)
	sr := SectionReader(r, 6, 5)
	buf := make([]byte, 5)
	n, err := sr.Read(buf)
	if err != nil {
		t.Fatalf("SectionReader Read error: %v", err)
	}
	if n != 5 {
		t.Errorf("SectionReader read %d bytes, want 5", n)
	}
	if string(buf) != "world" {
		t.Errorf("SectionReader read %q, want %q", string(buf), "world")
	}
}

func TestNewBuffer(t *testing.T) {
	data := []byte("hello")
	buf := NewBuffer(data)
	if buf.String() != "hello" {
		t.Errorf("NewBuffer = %q, want %q", buf.String(), "hello")
	}
}

func TestNewBufferString(t *testing.T) {
	buf := NewBufferString("hello")
	if buf.String() != "hello" {
		t.Errorf("NewBufferString = %q, want %q", buf.String(), "hello")
	}
}

func TestNewReader(t *testing.T) {
	data := []byte("hello")
	r := NewReader(data)
	buf := make([]byte, 5)
	n, err := r.Read(buf)
	if err != nil {
		t.Fatalf("NewReader Read error: %v", err)
	}
	if n != 5 || string(buf) != "hello" {
		t.Errorf("NewReader read %q, want %q", string(buf), "hello")
	}
}

func TestNewReaderString(t *testing.T) {
	r := NewReaderString("hello")
	buf := make([]byte, 5)
	n, err := r.Read(buf)
	if err != nil {
		t.Fatalf("NewReaderString Read error: %v", err)
	}
	if n != 5 || string(buf) != "hello" {
		t.Errorf("NewReaderString read %q, want %q", string(buf), "hello")
	}
}

func TestBufferedWriter(t *testing.T) {
	var buf bytes.Buffer
	bw := BufferedWriter(&buf)
	_, err := bw.WriteString("hello")
	if err != nil {
		t.Fatalf("BufferedWriter WriteString error: %v", err)
	}
	err = bw.Flush()
	if err != nil {
		t.Fatalf("BufferedWriter Flush error: %v", err)
	}
	if buf.String() != "hello" {
		t.Errorf("BufferedWriter wrote %q, want %q", buf.String(), "hello")
	}
}

func TestBufferedReadString(t *testing.T) {
	r := strings.NewReader("hello world")
	br := BufferedReadString(r)
	s, err := br.ReadString(' ')
	if err != nil {
		t.Fatalf("BufferedReadString ReadString error: %v", err)
	}
	if s != "hello " {
		t.Errorf("BufferedReadString read %q, want %q", s, "hello ")
	}
}

type testCloser struct {
	closed bool
}

func (c *testCloser) Close() error {
	c.closed = true
	return nil
}

type testCloserError struct{}

func (c *testCloserError) Close() error {
	return errors.New("close error")
}

func TestClose(t *testing.T) {
	c := &testCloser{}
	err := Close(c)
	if err != nil {
		t.Fatalf("Close error: %v", err)
	}
	if !c.closed {
		t.Errorf("Close did not close the closer")
	}
}

func TestCloseNil(t *testing.T) {
	err := Close(nil)
	if err != nil {
		t.Fatalf("Close(nil) error: %v", err)
	}
}

func TestCloseQuietly(t *testing.T) {
	c := &testCloser{}
	CloseQuietly(c)
	if !c.closed {
		t.Errorf("CloseQuietly did not close the closer")
	}
}

func TestCloseQuietlyNil(t *testing.T) {
	CloseQuietly(nil)
}

func TestCloseQuietlyError(t *testing.T) {
	c := &testCloserError{}
	CloseQuietly(c)
}

type testFlusher struct {
	buf bytes.Buffer
}

func (f *testFlusher) Write(p []byte) (n int, err error) {
	return f.buf.Write(p)
}

func (f *testFlusher) Flush() error {
	return nil
}

func TestFlush(t *testing.T) {
	f := &testFlusher{}
	err := Flush(f)
	if err != nil {
		t.Fatalf("Flush error: %v", err)
	}
}

func TestFlushNoFlusher(t *testing.T) {
	var buf bytes.Buffer
	err := Flush(&buf)
	if err == nil {
		t.Errorf("Flush should return error for non-flusher")
	}
}

type testReaderAt struct {
	data []byte
}

func (r *testReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	if off >= int64(len(r.data)) {
		return 0, io.EOF
	}
	n = copy(p, r.data[off:])
	if n < len(p) {
		return n, io.EOF
	}
	return n, nil
}

func TestReadAt(t *testing.T) {
	r := &testReaderAt{data: []byte("hello world")}
	buf := make([]byte, 5)
	n, err := ReadAt(r, buf, 6)
	if err != nil {
		t.Fatalf("ReadAt error: %v", err)
	}
	if n != 5 || string(buf) != "world" {
		t.Errorf("ReadAt read %q, want %q", string(buf), "world")
	}
}

type testWriterAt struct {
	buf bytes.Buffer
}

func (w *testWriterAt) WriteAt(p []byte, off int64) (n int, err error) {
	return w.buf.Write(p)
}

func TestWriteAt(t *testing.T) {
	w := &testWriterAt{}
	data := []byte("hello")
	n, err := WriteAt(w, data, 0)
	if err != nil {
		t.Fatalf("WriteAt error: %v", err)
	}
	if n != 5 {
		t.Errorf("WriteAt wrote %d bytes, want 5", n)
	}
}

func TestMultiReader(t *testing.T) {
	r1 := strings.NewReader("hello ")
	r2 := strings.NewReader("world")
	mr := MultiReader(r1, r2)
	data, err := io.ReadAll(mr)
	if err != nil {
		t.Fatalf("MultiReader ReadAll error: %v", err)
	}
	if string(data) != "hello world" {
		t.Errorf("MultiReader read %q, want %q", string(data), "hello world")
	}
}

func TestMultiWriter(t *testing.T) {
	var buf1, buf2 bytes.Buffer
	mw := MultiWriter(&buf1, &buf2)
	_, err := mw.Write([]byte("hello"))
	if err != nil {
		t.Fatalf("MultiWriter Write error: %v", err)
	}
	if buf1.String() != "hello" {
		t.Errorf("MultiWriter buf1 = %q, want %q", buf1.String(), "hello")
	}
	if buf2.String() != "hello" {
		t.Errorf("MultiWriter buf2 = %q, want %q", buf2.String(), "hello")
	}
}

func TestPipe(t *testing.T) {
	pr, pw := Pipe()
	go func() {
		pw.Write([]byte("hello"))
		pw.Close()
	}()
	data, err := io.ReadAll(pr)
	if err != nil {
		t.Fatalf("Pipe ReadAll error: %v", err)
	}
	if string(data) != "hello" {
		t.Errorf("Pipe read %q, want %q", string(data), "hello")
	}
}

func TestReadFull(t *testing.T) {
	r := strings.NewReader("hello world")
	buf := make([]byte, 5)
	n, err := ReadFull(r, buf)
	if err != nil {
		t.Fatalf("ReadFull error: %v", err)
	}
	if n != 5 || string(buf) != "hello" {
		t.Errorf("ReadFull read %q, want %q", string(buf), "hello")
	}
}

func TestReadAtLeast(t *testing.T) {
	r := strings.NewReader("hello world")
	buf := make([]byte, 10)
	n, err := ReadAtLeast(r, buf, 5)
	if err != nil {
		t.Fatalf("ReadAtLeast error: %v", err)
	}
	if n < 5 {
		t.Errorf("ReadAtLeast read %d bytes, want at least 5", n)
	}
}
