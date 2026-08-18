package ziputil

import (
	"archive/zip"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// setupTree creates a temp dir with the given relative files.
func setupTree(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for rel, content := range files {
		path := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func TestZipDir(t *testing.T) {
	root := setupTree(t, map[string]string{
		"a.txt":    "hello",
		"sub/b.go": "package b",
	})
	zipPath := filepath.Join(t.TempDir(), "out.zip")
	if err := Zip(root, zipPath); err != nil {
		t.Fatalf("Zip error: %v", err)
	}

	// verify zip contents
	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		t.Fatal(err)
	}
	defer zr.Close()
	var names []string
	for _, f := range zr.File {
		names = append(names, f.Name)
	}
	want := map[string]bool{"a.txt": true, "sub/b.go": true}
	for _, n := range names {
		clean := strings.TrimPrefix(n, "/")
		if !want[clean] {
			t.Errorf("unexpected zip entry: %q (all: %v)", n, names)
		}
	}
	if len(names) != 2 {
		t.Errorf("expected 2 entries, got %d: %v", len(names), names)
	}

	// content check
	var foundA bool
	for _, f := range zr.File {
		if f.Name == "a.txt" {
			rc, _ := f.Open()
			b, _ := io.ReadAll(rc)
			rc.Close()
			if string(b) != "hello" {
				t.Errorf("a.txt content = %q, want hello", b)
			}
			foundA = true
		}
	}
	if !foundA {
		t.Error("a.txt missing from zip")
	}
}

func TestZipEmptyDir(t *testing.T) {
	root := t.TempDir()
	zipPath := filepath.Join(t.TempDir(), "empty.zip")
	if err := Zip(root, zipPath); err != nil {
		t.Fatalf("Zip empty dir error: %v", err)
	}
	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		t.Fatal(err)
	}
	defer zr.Close()
	if len(zr.File) != 0 {
		t.Errorf("empty dir should zip to 0 entries, got %d", len(zr.File))
	}
}

func TestUnzip(t *testing.T) {
	root := setupTree(t, map[string]string{"x.txt": "data"})
	zipPath := filepath.Join(t.TempDir(), "in.zip")
	if err := Zip(root, zipPath); err != nil {
		t.Fatal(err)
	}

	outDir := filepath.Join(t.TempDir(), "out")
	if err := Unzip(zipPath, outDir); err != nil {
		t.Fatalf("Unzip error: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(outDir, "x.txt"))
	if err != nil {
		t.Fatalf("unzipped file missing: %v", err)
	}
	if string(b) != "data" {
		t.Errorf("unzipped content = %q, want data", b)
	}
}

func TestUnzipNestedDir(t *testing.T) {
	root := setupTree(t, map[string]string{
		"deep/nested/file.txt": "deep",
	})
	zipPath := filepath.Join(t.TempDir(), "nested.zip")
	if err := Zip(root, zipPath); err != nil {
		t.Fatal(err)
	}
	outDir := filepath.Join(t.TempDir(), "out")
	if err := Unzip(zipPath, outDir); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(outDir, "deep", "nested", "file.txt"))
	if err != nil {
		t.Fatalf("nested file missing: %v", err)
	}
	if string(b) != "deep" {
		t.Errorf("content = %q", b)
	}
}

func TestUnzipMissingFile(t *testing.T) {
	if err := Unzip(filepath.Join(t.TempDir(), "nope.zip"), t.TempDir()); err == nil {
		t.Error("Unzip(missing) should error")
	}
}

func TestGzipString(t *testing.T) {
	orig := "hello world"
	compressed, err := Gzip([]byte(orig))
	if err != nil {
		t.Fatalf("Gzip error: %v", err)
	}
	if len(compressed) >= len(orig) {
		t.Logf("note: %d bytes compressed from %d", len(compressed), len(orig))
	}
	back, err := Gunzip(compressed)
	if err != nil {
		t.Fatalf("Gunzip error: %v", err)
	}
	if string(back) != orig {
		t.Errorf("round trip mismatch: %q != %q", back, orig)
	}
}

func TestGzipEmpty(t *testing.T) {
	c, err := Gzip(nil)
	if err != nil {
		t.Fatal(err)
	}
	b, err := Gunzip(c)
	if err != nil {
		t.Fatal(err)
	}
	if len(b) != 0 {
		t.Errorf("empty round trip = %d bytes", len(b))
	}
}

func TestGunzipInvalid(t *testing.T) {
	if _, err := Gunzip([]byte("not gzip data")); err == nil {
		t.Error("Gunzip(invalid) should error")
	}
}

func TestGzipFile(t *testing.T) {
	tmp := t.TempDir()
	src := filepath.Join(tmp, "data.txt")
	if err := os.WriteFile(src, []byte("file to gzip"), 0o644); err != nil {
		t.Fatal(err)
	}
	gzPath := filepath.Join(tmp, "data.txt.gz")
	if err := GzipFile(src, gzPath); err != nil {
		t.Fatalf("GzipFile error: %v", err)
	}
	// verify it's real gzip
	f, err := os.Open(gzPath)
	if err != nil {
		t.Fatal(err)
	}
	gr, err := gzip.NewReader(f)
	if err != nil {
		t.Fatalf("not a valid gzip file: %v", err)
	}
	b, _ := io.ReadAll(gr)
	gr.Close()
	f.Close()
	if string(b) != "file to gzip" {
		t.Errorf("gzip file content = %q", b)
	}
}

func TestGunzipFile(t *testing.T) {
	tmp := t.TempDir()
	src := filepath.Join(tmp, "x.txt")
	if err := os.WriteFile(src, []byte("round trip"), 0o644); err != nil {
		t.Fatal(err)
	}
	gz := filepath.Join(tmp, "x.gz")
	if err := GzipFile(src, gz); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(tmp, "x.out.txt")
	if err := GunzipFile(gz, out); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(out)
	if string(b) != "round trip" {
		t.Errorf("GunzipFile content = %q", b)
	}
}

func TestTarGzRoundTrip(t *testing.T) {
	root := setupTree(t, map[string]string{
		"a.txt":    "aaa",
		"dir/b.go": "package b",
	})
	tgz := filepath.Join(t.TempDir(), "out.tar.gz")
	if err := TarGz(root, tgz); err != nil {
		t.Fatalf("TarGz error: %v", err)
	}

	outDir := filepath.Join(t.TempDir(), "extracted")
	if err := UntarGz(tgz, outDir); err != nil {
		t.Fatalf("UntarGz error: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(outDir, "a.txt"))
	if err != nil {
		t.Fatalf("extracted a.txt missing: %v", err)
	}
	if string(b) != "aaa" {
		t.Errorf("a.txt = %q", b)
	}
	b2, err := os.ReadFile(filepath.Join(outDir, "dir", "b.go"))
	if err != nil {
		t.Fatalf("extracted dir/b.go missing: %v", err)
	}
	if string(b2) != "package b" {
		t.Errorf("b.go = %q", b2)
	}
}

func TestListFiles(t *testing.T) {
	root := setupTree(t, map[string]string{
		"a.txt":    "1",
		"b/c.txt":  "2",
		"b/d/e.go": "3",
	})
	zipPath := filepath.Join(t.TempDir(), "list.zip")
	if err := Zip(root, zipPath); err != nil {
		t.Fatal(err)
	}
	files, err := ListFiles(zipPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 3 {
		t.Errorf("ListFiles = %d entries, want 3: %v", len(files), files)
	}
}

func TestZipFile(t *testing.T) {
	tmp := t.TempDir()
	payload := []byte("single file payload")
	src := filepath.Join(tmp, "single.txt")
	if err := os.WriteFile(src, payload, 0o644); err != nil {
		t.Fatal(err)
	}
	zipPath := filepath.Join(tmp, "single.zip")
	if err := ZipFile(src, zipPath); err != nil {
		t.Fatalf("ZipFile error: %v", err)
	}
	out := filepath.Join(tmp, "out")
	if err := Unzip(zipPath, out); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(filepath.Join(out, "single.txt"))
	if string(b) != string(payload) {
		t.Errorf("ZipFile round trip = %q, want %q", b, payload)
	}
}

func TestPathTraversalGuard(t *testing.T) {
	// craft a zip with a ../ entry and ensure Unzip refuses it
	zipPath := filepath.Join(t.TempDir(), "evil.zip")
	zf, err := os.Create(zipPath)
	if err != nil {
		t.Fatal(err)
	}
	zw := zip.NewWriter(zf)
	w, err := zw.Create("../evil.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = w.Write([]byte("evil"))
	_ = zw.Close()
	_ = zf.Close()

	out := t.TempDir()
	if err := Unzip(zipPath, out); err == nil {
		t.Error("Unzip should refuse path traversal entries")
	}
	// nothing should exist outside out dir
	if _, err := os.Stat(filepath.Join(filepath.Dir(out), "evil.txt")); err == nil {
		t.Error("traversal file was written outside target dir")
	}
}