package csv

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	data := "name,age,city\nAlice,30,NYC\nBob,25,SF\n"
	rows, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	want := [][]string{
		{"name", "age", "city"},
		{"Alice", "30", "NYC"},
		{"Bob", "25", "SF"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("Decode = %v", rows)
	}
}

func TestDecodeSkipsBlankRows(t *testing.T) {
	rows, err := Decode("a,b\n\nc,d\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 {
		t.Errorf("rows = %v", rows)
	}
}

func TestDecodeWithQuotes(t *testing.T) {
	data := `a,"hello, world",c
`
	rows, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	if rows[0][1] != "hello, world" {
		t.Errorf("quoted field = %q", rows[0][1])
	}
}

func TestDecodeBadInput(t *testing.T) {
	// unterminated quote → error
	if _, err := Decode(`"unterminated`); err == nil {
		t.Error("unterminated quote should error")
	}
}

func TestEncode(t *testing.T) {
	rows := [][]string{
		{"name", "age"},
		{"Alice", "30"},
	}
	out, err := Encode(rows)
	if err != nil {
		t.Fatal(err)
	}
	want := "name,age\nAlice,30\n"
	if !strings.HasPrefix(out, want) {
		t.Errorf("Encode = %q, want prefix %q", out, want)
	}
	// no trailing garbage beyond final newline
	if strings.Count(out, "\n") != 2 {
		t.Errorf("Encode newlines = %q", out)
	}
}

func TestEncodeEscapesQuotes(t *testing.T) {
	out, err := Encode([][]string{{`say "hi"`}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, `"say ""hi"""`) {
		t.Errorf("Encode escape = %q", out)
	}
}

func TestRoundTrip(t *testing.T) {
	original := [][]string{
		{"id", "note"},
		{"1", `quoted "field", with comma and "nested"`},
		{"2", ""},
	}
	encoded, err := Encode(original)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, original) {
		t.Errorf("round trip mismatch:\n got  %v\n want %v", decoded, original)
	}
}

func TestSemicolonDelimiter(t *testing.T) {
	rows, err := Decode("a;b;c\n1;2;3\n", WithDelimiter(';'))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(rows, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}) {
		t.Errorf("semicolon rows = %v", rows)
	}
}

func TestWriteReadFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.csv")
	rows := [][]string{{"h", "v"}, {"1", "x"}}

	if err := WriteFile(path, rows); err != nil {
		t.Fatal(err)
	}
	got, err := ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, rows) {
		t.Errorf("ReadFile = %v", got)
	}
	// temp dir must be cleaned by t.TempDir
	if _, err := os.Stat(path); err != nil {
		t.Errorf("file not found: %v", err)
	}
}

func TestWriteFileBadPath(t *testing.T) {
	err := WriteFile(filepath.Join(t.TempDir(), "no", "dir", "x.csv"), [][]string{{"a"}})
	if err == nil {
		t.Error("WriteFile into missing dir should error")
	}
}

func TestReadFileMissing(t *testing.T) {
	_, err := ReadFile(filepath.Join(t.TempDir(), "missing.csv"))
	if err == nil {
		t.Error("ReadFile(missing) should error")
	}
}