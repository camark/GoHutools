package csv

import (
	"bytes"
	"encoding/csv"
	"os"
	"strings"
)

// Options customize CSV parsing/writing.
type Options struct {
	Delimiter rune // field separator; defaults to ','
}

// Option applies an Options change.
type Option func(*Options)

// WithDelimiter sets the field separator (default ',').
func WithDelimiter(r rune) Option {
	return func(o *Options) { o.Delimiter = r }
}

func apply(opts []Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Decode parses a CSV string into rows of fields.
// Blank lines (records consisting of a single empty field) are skipped.
func Decode(s string, opts ...Option) ([][]string, error) {
	o := apply(opts)
	r := csv.NewReader(strings.NewReader(s))
	if o.Delimiter != 0 {
		r.Comma = o.Delimiter
	}
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	out := rows[:0]
	for _, row := range rows {
		if len(row) == 1 && row[0] == "" {
			continue // blank line
		}
		out = append(out, row)
	}
	return out, nil
}

// Encode serializes rows into a CSV string, quoting fields as needed.
func Encode(rows [][]string, opts ...Option) (string, error) {
	o := apply(opts)
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if o.Delimiter != 0 {
		w.Comma = o.Delimiter
	}
	if err := w.WriteAll(rows); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ReadFile reads and parses a CSV file.
func ReadFile(path string, opts ...Option) ([][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Decode(string(data), opts...)
}

// WriteFile writes rows to a CSV file at path.
func WriteFile(path string, rows [][]string, opts ...Option) error {
	content, err := Encode(rows, opts...)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0o644)
}