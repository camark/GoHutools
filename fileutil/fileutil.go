package fileutil

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Exists checks if file exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsFile checks if path is a file
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsDir checks if path is a directory
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Size returns file size
func Size(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// Name returns file name without extension
func Name(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

// Ext returns file extension
func Ext(path string) string {
	return filepath.Ext(path)
}

// BaseName returns file name with extension
func BaseName(path string) string {
	return filepath.Base(path)
}

// Dir returns directory of file
func Dir(path string) string {
	return filepath.Dir(path)
}

// Abs returns absolute path
func Abs(path string) (string, error) {
	return filepath.Abs(path)
}

// Clean returns cleaned path
func Clean(path string) string {
	return filepath.Clean(path)
}

// Join joins path elements
func Join(elem ...string) string {
	return filepath.Join(elem...)
}

// ReadBytes reads entire file
func ReadBytes(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ReadString reads entire file as string
func ReadString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLines reads file lines
func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// WriteBytes writes bytes to file
func WriteBytes(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

// WriteString writes string to file
func WriteString(path string, s string) error {
	return os.WriteFile(path, []byte(s), 0644)
}

// WriteLines writes lines to file
func WriteLines(path string, lines []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	for i, line := range lines {
		if _, err := w.WriteString(line); err != nil {
			_ = f.Close()
			return err
		}
		if i < len(lines)-1 {
			if _, err := w.WriteString("\n"); err != nil {
				_ = f.Close()
				return err
			}
		}
	}
	if err := w.Flush(); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

// AppendBytes appends bytes to file
func AppendBytes(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, werr := f.Write(data)
	cerr := f.Close()
	if werr != nil {
		return werr
	}
	return cerr
}

// AppendString appends string to file
func AppendString(path string, s string) error {
	return AppendBytes(path, []byte(s))
}

// CopyFile copies file from src to dst
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() { _ = srcFile.Close() }()

	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, srcInfo.Mode())
	if err != nil {
		return err
	}

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		_ = dstFile.Close()
		return err
	}
	return dstFile.Close()
}

// MoveFile moves file from src to dst
func MoveFile(src, dst string) error {
	err := os.Rename(src, dst)
	if err == nil {
		return nil
	}
	if err := CopyFile(src, dst); err != nil {
		return err
	}
	return os.Remove(src)
}

// DeleteFile deletes file
func DeleteFile(path string) error {
	return os.Remove(path)
}

// Mkdir creates directory (and parents)
func Mkdir(path string) error {
	return os.MkdirAll(path, 0755)
}

// MkdirAll creates directory and all parents
func MkdirAll(path string) error {
	return os.MkdirAll(path, 0755)
}

// ListFiles lists files in directory
func ListFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}
	return files, nil
}

// ListDirs lists directories in directory
func ListDirs(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, filepath.Join(dir, entry.Name()))
		}
	}
	return dirs, nil
}

// ListFilesRecursive lists all files recursively
func ListFilesRecursive(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// Walk walks directory tree
func Walk(root string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(root, walkFn)
}

// Glob returns matching files
func Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

// TempDir creates temporary directory
func TempDir(dir, prefix string) (string, error) {
	return os.MkdirTemp(dir, prefix)
}

// TempFile creates temporary file
func TempFile(dir, prefix string) (*os.File, error) {
	return os.CreateTemp(dir, prefix)
}

// Remove removes file or directory
func Remove(path string) error {
	return os.Remove(path)
}

// RemoveAll removes path and any children it contains
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// Rename renames file
func Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

// Chmod changes file permissions
func Chmod(path string, mode os.FileMode) error {
	return os.Chmod(path, mode)
}

// Chown changes file ownership
func Chown(path string, uid, gid int) error {
	return os.Chown(path, uid, gid)
}

// IsSymlink checks if path is symlink
func IsSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
}

// Symlink creates symlink
func Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

// Readlink returns symlink target
func Readlink(path string) (string, error) {
	return os.Readlink(path)
}

// Md5 returns MD5 hash of file
func Md5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Sha256 returns SHA256 hash of file
func Sha256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Lines returns line count of file
func Lines(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer func() { _ = f.Close() }()

	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

// Tail reads last n lines of file
func Tail(path string, n int) ([]string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return []string{}, nil
	}
	if n >= len(lines) {
		return lines, nil
	}
	return lines[len(lines)-n:], nil
}

// Head reads first n lines of file
func Head(path string, n int) ([]string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return []string{}, nil
	}
	if n >= len(lines) {
		return lines, nil
	}
	return lines[:n], nil
}

// Grep searches file for pattern
func Grep(path string, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	var matches []string
	for _, line := range lines {
		if re.MatchString(line) {
			matches = append(matches, line)
		}
	}
	return matches, nil
}

// Replace replaces string in file
func Replace(path string, old, new string) error {
	content, err := ReadString(path)
	if err != nil {
		return err
	}
	content = strings.ReplaceAll(content, old, new)
	return WriteString(path, content)
}

// Backup creates backup of file
func Backup(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(path)
	base := strings.TrimSuffix(path, ext)
	backupPath := fmt.Sprintf("%s.backup.%s%s", base, time.Now().Format("20060102150405"), ext)

	if err := CopyFile(path, backupPath); err != nil {
		return "", err
	}

	if err := os.Chmod(backupPath, info.Mode()); err != nil {
		return "", err
	}

	return backupPath, nil
}

// Normalize normalizes path separators
func Normalize(path string) string {
	return filepath.ToSlash(path)
}

// ToURI converts path to file URI
func ToURI(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		absPath = path
	}
	absPath = filepath.ToSlash(absPath)
	return "file:///" + absPath
}

// UserHome returns user home directory
func UserHome() (string, error) {
	return os.UserHomeDir()
}

// WorkingDir returns working directory
func WorkingDir() (string, error) {
	return os.Getwd()
}

// Separator returns path separator
func Separator() string {
	return string(filepath.Separator)
}

// ListSeparator returns path list separator
func ListSeparator() string {
	return string(filepath.ListSeparator)
}
