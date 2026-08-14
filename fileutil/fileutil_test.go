package fileutil

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func setupTestDir(t *testing.T) string {
	t.Helper()
	dir, err := os.MkdirTemp("", "fileutil_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	return dir
}

func cleanupTestDir(t *testing.T, dir string) {
	t.Helper()
	os.RemoveAll(dir)
}

func TestExists(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	if !Exists(testFile) {
		t.Errorf("Exists should return true for existing file")
	}
	if Exists(filepath.Join(dir, "nonexistent.txt")) {
		t.Errorf("Exists should return false for non-existing file")
	}
}

func TestIsFile(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	if !IsFile(testFile) {
		t.Errorf("IsFile should return true for file")
	}
	if IsFile(dir) {
		t.Errorf("IsFile should return false for directory")
	}
}

func TestIsDir(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	if !IsDir(dir) {
		t.Errorf("IsDir should return true for directory")
	}
	if IsDir(testFile) {
		t.Errorf("IsDir should return false for file")
	}
}

func TestSize(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	data := []byte("hello world")
	os.WriteFile(testFile, data, 0644)

	size, err := Size(testFile)
	if err != nil {
		t.Fatalf("Size error: %v", err)
	}
	if size != int64(len(data)) {
		t.Errorf("Size = %d, want %d", size, len(data))
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/path/to/file.txt", "file"},
		{"/path/to/file.tar.gz", "file.tar"},
		{"file.txt", "file"},
		{"/path/to/file", "file"},
	}
	for _, tt := range tests {
		if got := Name(tt.path); got != tt.want {
			t.Errorf("Name(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}

func TestExt(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/path/to/file.txt", ".txt"},
		{"/path/to/file.tar.gz", ".gz"},
		{"file.txt", ".txt"},
		{"/path/to/file", ""},
	}
	for _, tt := range tests {
		if got := Ext(tt.path); got != tt.want {
			t.Errorf("Ext(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}

func TestBaseName(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/path/to/file.txt", "file.txt"},
		{"/path/to/file.tar.gz", "file.tar.gz"},
		{"file.txt", "file.txt"},
	}
	for _, tt := range tests {
		if got := BaseName(tt.path); got != tt.want {
			t.Errorf("BaseName(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}

func TestDir(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{filepath.Join("path", "to", "file.txt"), filepath.Join("path", "to")},
		{"file.txt", "."},
	}
	for _, tt := range tests {
		if got := Dir(tt.path); got != tt.want {
			t.Errorf("Dir(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}

func TestAbs(t *testing.T) {
	abs, err := Abs("test.txt")
	if err != nil {
		t.Fatalf("Abs error: %v", err)
	}
	if !filepath.IsAbs(abs) {
		t.Errorf("Abs should return absolute path")
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{filepath.Join(".", "path", "to", "..", "file.txt"), filepath.Join("path", "file.txt")},
		{filepath.Join("path", "", "to", "file.txt"), filepath.Join("path", "to", "file.txt")},
	}
	for _, tt := range tests {
		if got := Clean(tt.path); got != tt.want {
			t.Errorf("Clean(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}

func TestJoin(t *testing.T) {
	result := Join("path", "to", "file.txt")
	expected := filepath.Join("path", "to", "file.txt")
	if result != expected {
		t.Errorf("Join = %q, want %q", result, expected)
	}
}

func TestReadBytes(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	data := []byte("hello world")
	os.WriteFile(testFile, data, 0644)

	readData, err := ReadBytes(testFile)
	if err != nil {
		t.Fatalf("ReadBytes error: %v", err)
	}
	if string(readData) != string(data) {
		t.Errorf("ReadBytes = %q, want %q", string(readData), string(data))
	}
}

func TestReadString(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	data := "hello world"
	os.WriteFile(testFile, []byte(data), 0644)

	readData, err := ReadString(testFile)
	if err != nil {
		t.Fatalf("ReadString error: %v", err)
	}
	if readData != data {
		t.Errorf("ReadString = %q, want %q", readData, data)
	}
}

func TestReadLines(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	data := "line1\nline2\nline3"
	os.WriteFile(testFile, []byte(data), 0644)

	lines, err := ReadLines(testFile)
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

func TestWriteBytes(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	data := []byte("hello world")
	err := WriteBytes(testFile, data)
	if err != nil {
		t.Fatalf("WriteBytes error: %v", err)
	}

	readData, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if string(readData) != string(data) {
		t.Errorf("WriteBytes wrote %q, want %q", string(readData), string(data))
	}
}

func TestWriteString(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	data := "hello world"
	err := WriteString(testFile, data)
	if err != nil {
		t.Fatalf("WriteString error: %v", err)
	}

	readData, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if string(readData) != data {
		t.Errorf("WriteString wrote %q, want %q", string(readData), data)
	}
}

func TestWriteLines(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	lines := []string{"line1", "line2", "line3"}
	err := WriteLines(testFile, lines)
	if err != nil {
		t.Fatalf("WriteLines error: %v", err)
	}

	readData, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	expected := "line1\nline2\nline3"
	if string(readData) != expected {
		t.Errorf("WriteLines wrote %q, want %q", string(readData), expected)
	}
}

func TestAppendBytes(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("hello"), 0644)

	err := AppendBytes(testFile, []byte(" world"))
	if err != nil {
		t.Fatalf("AppendBytes error: %v", err)
	}

	readData, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if string(readData) != "hello world" {
		t.Errorf("AppendBytes result = %q, want %q", string(readData), "hello world")
	}
}

func TestAppendString(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("hello"), 0644)

	err := AppendString(testFile, " world")
	if err != nil {
		t.Fatalf("AppendString error: %v", err)
	}

	readData, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if string(readData) != "hello world" {
		t.Errorf("AppendString result = %q, want %q", string(readData), "hello world")
	}
}

func TestCopyFile(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	srcFile := filepath.Join(dir, "src.txt")
	dstFile := filepath.Join(dir, "dst.txt")
	data := []byte("hello world")
	os.WriteFile(srcFile, data, 0644)

	err := CopyFile(srcFile, dstFile)
	if err != nil {
		t.Fatalf("CopyFile error: %v", err)
	}

	readData, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if string(readData) != string(data) {
		t.Errorf("CopyFile copied %q, want %q", string(readData), string(data))
	}
}

func TestMoveFile(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	srcFile := filepath.Join(dir, "src.txt")
	dstFile := filepath.Join(dir, "dst.txt")
	data := []byte("hello world")
	os.WriteFile(srcFile, data, 0644)

	err := MoveFile(srcFile, dstFile)
	if err != nil {
		t.Fatalf("MoveFile error: %v", err)
	}

	if Exists(srcFile) {
		t.Errorf("MoveFile should remove source file")
	}

	readData, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if string(readData) != string(data) {
		t.Errorf("MoveFile moved %q, want %q", string(readData), string(data))
	}
}

func TestDeleteFile(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	err := DeleteFile(testFile)
	if err != nil {
		t.Fatalf("DeleteFile error: %v", err)
	}

	if Exists(testFile) {
		t.Errorf("DeleteFile should delete file")
	}
}

func TestMkdir(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	newDir := filepath.Join(dir, "a", "b", "c")
	err := Mkdir(newDir)
	if err != nil {
		t.Fatalf("Mkdir error: %v", err)
	}

	if !IsDir(newDir) {
		t.Errorf("Mkdir should create directory")
	}
}

func TestMkdirAll(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	newDir := filepath.Join(dir, "a", "b", "c")
	err := MkdirAll(newDir)
	if err != nil {
		t.Fatalf("MkdirAll error: %v", err)
	}

	if !IsDir(newDir) {
		t.Errorf("MkdirAll should create directory")
	}
}

func TestListFiles(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(dir, "file2.txt"), []byte("test"), 0644)
	os.Mkdir(filepath.Join(dir, "subdir"), 0755)

	files, err := ListFiles(dir)
	if err != nil {
		t.Fatalf("ListFiles error: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("ListFiles returned %d files, want 2", len(files))
	}
}

func TestListDirs(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("test"), 0644)
	os.Mkdir(filepath.Join(dir, "subdir1"), 0755)
	os.Mkdir(filepath.Join(dir, "subdir2"), 0755)

	dirs, err := ListDirs(dir)
	if err != nil {
		t.Fatalf("ListDirs error: %v", err)
	}
	if len(dirs) != 2 {
		t.Errorf("ListDirs returned %d dirs, want 2", len(dirs))
	}
}

func TestListFilesRecursive(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	subdir := filepath.Join(dir, "subdir")
	os.Mkdir(subdir, 0755)
	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(subdir, "file2.txt"), []byte("test"), 0644)

	files, err := ListFilesRecursive(dir)
	if err != nil {
		t.Fatalf("ListFilesRecursive error: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("ListFilesRecursive returned %d files, want 2", len(files))
	}
}

func TestWalk(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	subdir := filepath.Join(dir, "subdir")
	os.Mkdir(subdir, 0755)
	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(subdir, "file2.txt"), []byte("test"), 0644)

	var walked []string
	err := Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			walked = append(walked, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Walk error: %v", err)
	}
	if len(walked) != 2 {
		t.Errorf("Walk visited %d files, want 2", len(walked))
	}
}

func TestGlob(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(dir, "file2.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(dir, "file3.log"), []byte("test"), 0644)

	pattern := filepath.Join(dir, "*.txt")
	files, err := Glob(pattern)
	if err != nil {
		t.Fatalf("Glob error: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("Glob returned %d files, want 2", len(files))
	}
}

func TestTempDir(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	tempDir, err := TempDir(dir, "test")
	if err != nil {
		t.Fatalf("TempDir error: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if !IsDir(tempDir) {
		t.Errorf("TempDir should create directory")
	}
}

func TestTempFile(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	tempFile, err := TempFile(dir, "test")
	if err != nil {
		t.Fatalf("TempFile error: %v", err)
	}
	tempFile.Close()
	defer os.Remove(tempFile.Name())

	if !IsFile(tempFile.Name()) {
		t.Errorf("TempFile should create file")
	}
}

func TestRemove(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	err := Remove(testFile)
	if err != nil {
		t.Fatalf("Remove error: %v", err)
	}

	if Exists(testFile) {
		t.Errorf("Remove should delete file")
	}
}

func TestRemoveAll(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	subdir := filepath.Join(dir, "subdir")
	os.Mkdir(subdir, 0755)
	os.WriteFile(filepath.Join(subdir, "test.txt"), []byte("test"), 0644)

	err := RemoveAll(dir)
	if err != nil {
		t.Fatalf("RemoveAll error: %v", err)
	}

	if Exists(dir) {
		t.Errorf("RemoveAll should delete directory and contents")
	}
}

func TestRename(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	oldFile := filepath.Join(dir, "old.txt")
	newFile := filepath.Join(dir, "new.txt")
	os.WriteFile(oldFile, []byte("test"), 0644)

	err := Rename(oldFile, newFile)
	if err != nil {
		t.Fatalf("Rename error: %v", err)
	}

	if Exists(oldFile) {
		t.Errorf("Rename should remove old file")
	}
	if !Exists(newFile) {
		t.Errorf("Rename should create new file")
	}
}

func TestChmod(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	err := Chmod(testFile, 0444)
	if err != nil {
		t.Fatalf("Chmod error: %v", err)
	}

	info, err := os.Stat(testFile)
	if err != nil {
		t.Fatalf("Stat error: %v", err)
	}
	// On Windows, permissions work differently - just verify Chmod doesn't error
	if runtime.GOOS != "windows" {
		if info.Mode().Perm() != 0444 {
			t.Errorf("Chmod mode = %o, want %o", info.Mode().Perm(), 0444)
		}
	}
}

func TestChown(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Chown not supported on Windows")
	}

	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	err := Chown(testFile, os.Getuid(), os.Getgid())
	if err != nil {
		t.Fatalf("Chown error: %v", err)
	}
}

func TestIsSymlink(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	linkFile := filepath.Join(dir, "link.txt")
	err := os.Symlink(testFile, linkFile)
	if err != nil {
		t.Skipf("Skipping symlink test: %v", err)
	}

	if !IsSymlink(linkFile) {
		t.Errorf("IsSymlink should return true for symlink")
	}
	if IsSymlink(testFile) {
		t.Errorf("IsSymlink should return false for regular file")
	}
}

func TestSymlink(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	linkFile := filepath.Join(dir, "link.txt")
	err := Symlink(testFile, linkFile)
	if err != nil {
		t.Skipf("Skipping symlink test: %v", err)
	}

	if !IsSymlink(linkFile) {
		t.Errorf("Symlink should create symlink")
	}
}

func TestReadlink(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	linkFile := filepath.Join(dir, "link.txt")
	err := os.Symlink(testFile, linkFile)
	if err != nil {
		t.Skipf("Skipping readlink test: %v", err)
	}

	target, err := Readlink(linkFile)
	if err != nil {
		t.Fatalf("Readlink error: %v", err)
	}
	if target != testFile {
		t.Errorf("Readlink = %q, want %q", target, testFile)
	}
}

func TestMd5(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("hello world"), 0644)

	hash, err := Md5(testFile)
	if err != nil {
		t.Fatalf("Md5 error: %v", err)
	}
	// MD5 of "hello world" is 5eb63bbbe01eeed093cb22bb8f5acdc3
	expected := "5eb63bbbe01eeed093cb22bb8f5acdc3"
	if hash != expected {
		t.Errorf("Md5 = %q, want %q", hash, expected)
	}
}

func TestSha256(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("hello world"), 0644)

	hash, err := Sha256(testFile)
	if err != nil {
		t.Fatalf("Sha256 error: %v", err)
	}
	// SHA256 of "hello world" is b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
	expected := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
	if hash != expected {
		t.Errorf("Sha256 = %q, want %q", hash, expected)
	}
}

func TestLines(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("line1\nline2\nline3"), 0644)

	count, err := Lines(testFile)
	if err != nil {
		t.Fatalf("Lines error: %v", err)
	}
	if count != 3 {
		t.Errorf("Lines = %d, want 3", count)
	}
}

func TestTail(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("line1\nline2\nline3\nline4\nline5"), 0644)

	lines, err := Tail(testFile, 3)
	if err != nil {
		t.Fatalf("Tail error: %v", err)
	}
	expected := []string{"line3", "line4", "line5"}
	if len(lines) != len(expected) {
		t.Fatalf("Tail returned %d lines, want %d", len(lines), len(expected))
	}
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Tail[%d] = %q, want %q", i, line, expected[i])
		}
	}
}

func TestHead(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("line1\nline2\nline3\nline4\nline5"), 0644)

	lines, err := Head(testFile, 3)
	if err != nil {
		t.Fatalf("Head error: %v", err)
	}
	expected := []string{"line1", "line2", "line3"}
	if len(lines) != len(expected) {
		t.Fatalf("Head returned %d lines, want %d", len(lines), len(expected))
	}
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Head[%d] = %q, want %q", i, line, expected[i])
		}
	}
}

func TestGrep(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("line1\nline2 foo\nline3\nline4 foo"), 0644)

	matches, err := Grep(testFile, "foo")
	if err != nil {
		t.Fatalf("Grep error: %v", err)
	}
	if len(matches) != 2 {
		t.Errorf("Grep returned %d matches, want 2", len(matches))
	}
}

func TestReplace(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("hello world"), 0644)

	err := Replace(testFile, "world", "golang")
	if err != nil {
		t.Fatalf("Replace error: %v", err)
	}

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if string(data) != "hello golang" {
		t.Errorf("Replace result = %q, want %q", string(data), "hello golang")
	}
}

func TestBackup(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	os.WriteFile(testFile, []byte("hello world"), 0644)

	backupPath, err := Backup(testFile)
	if err != nil {
		t.Fatalf("Backup error: %v", err)
	}
	defer os.Remove(backupPath)

	if !Exists(backupPath) {
		t.Errorf("Backup should create backup file")
	}

	readData, err := os.ReadFile(backupPath)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if string(readData) != "hello world" {
		t.Errorf("Backup file content = %q, want %q", string(readData), "hello world")
	}

	if !strings.Contains(backupPath, "backup") {
		t.Errorf("Backup path should contain 'backup', got %q", backupPath)
	}
}

func TestNormalize(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Normalize separator test only relevant on Windows")
	}
	result := Normalize("path\\to\\file.txt")
	expected := "path/to/file.txt"
	if result != expected {
		t.Errorf("Normalize = %q, want %q", result, expected)
	}
}

func TestToURI(t *testing.T) {
	dir := setupTestDir(t)
	defer cleanupTestDir(t, dir)

	testFile := filepath.Join(dir, "test.txt")
	uri := ToURI(testFile)

	if !strings.HasPrefix(uri, "file:///") {
		t.Errorf("ToURI should start with 'file:///', got %q", uri)
	}
}

func TestUserHome(t *testing.T) {
	home, err := UserHome()
	if err != nil {
		t.Fatalf("UserHome error: %v", err)
	}
	if home == "" {
		t.Errorf("UserHome should not be empty")
	}
}

func TestWorkingDir(t *testing.T) {
	wd, err := WorkingDir()
	if err != nil {
		t.Fatalf("WorkingDir error: %v", err)
	}
	if wd == "" {
		t.Errorf("WorkingDir should not be empty")
	}
}

func TestSeparator(t *testing.T) {
	sep := Separator()
	if sep == "" {
		t.Errorf("Separator should not be empty")
	}
}

func TestListSeparator(t *testing.T) {
	sep := ListSeparator()
	if sep == "" {
		t.Errorf("ListSeparator should not be empty")
	}
}
