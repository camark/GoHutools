---
title: fileutil
---

# fileutil

文件操作工具包，提供文件判断、读写、复制、移动、遍历、哈希计算等功能，类似于 Java Hutool 的 FileUtil。

## 导入

```go
import "github.com/camark/GoHutools/fileutil"
```

## 函数列表

### Exists

```go
func Exists(path string) bool
```

检查文件或目录是否存在。

**示例:**

```go
if fileutil.Exists("/tmp/test.txt") {
    fmt.Println("文件存在")
}
```

### IsFile

```go
func IsFile(path string) bool
```

判断路径是否为文件（非目录）。

**示例:**

```go
if fileutil.IsFile("/tmp/test.txt") {
    fmt.Println("这是一个文件")
}
```

### IsDir

```go
func IsDir(path string) bool
```

判断路径是否为目录。

**示例:**

```go
if fileutil.IsDir("/tmp") {
    fmt.Println("这是一个目录")
}
```

### Size

```go
func Size(path string) (int64, error)
```

获取文件大小（字节）。

**示例:**

```go
size, err := fileutil.Size("/tmp/test.txt")
fmt.Printf("文件大小: %d 字节\n", size)
```

### Name

```go
func Name(path string) string
```

获取不含扩展名的文件名。

**示例:**

```go
name := fileutil.Name("/tmp/test.txt")
fmt.Println(name) // test
```

### Ext

```go
func Ext(path string) string
```

获取文件扩展名（包含点号）。

**示例:**

```go
ext := fileutil.Ext("/tmp/test.txt")
fmt.Println(ext) // .txt
```

### BaseName

```go
func BaseName(path string) string
```

获取含扩展名的文件名。

**示例:**

```go
base := fileutil.BaseName("/tmp/test.txt")
fmt.Println(base) // test.txt
```

### Dir

```go
func Dir(path string) string
```

获取文件所在目录。

**示例:**

```go
dir := fileutil.Dir("/tmp/test.txt")
fmt.Println(dir) // /tmp
```

### Abs

```go
func Abs(path string) (string, error)
```

获取绝对路径。

**示例:**

```go
abs, err := fileutil.Abs("./test.txt")
fmt.Println(abs)
```

### Clean

```go
func Clean(path string) string
```

清理路径，移除多余的分隔符和 `.`、`..`。

**示例:**

```go
clean := fileutil.Clean("/tmp//./test.txt")
fmt.Println(clean) // /tmp/test.txt
```

### Join

```go
func Join(elem ...string) string
```

拼接路径元素。

**示例:**

```go
path := fileutil.Join("/tmp", "subdir", "test.txt")
fmt.Println(path) // /tmp/subdir/test.txt
```

### ReadBytes

```go
func ReadBytes(path string) ([]byte, error)
```

读取整个文件，返回字节切片。

**示例:**

```go
data, err := fileutil.ReadBytes("/tmp/test.txt")
fmt.Println(string(data))
```

### ReadString

```go
func ReadString(path string) (string, error)
```

读取整个文件，返回字符串。

**示例:**

```go
content, err := fileutil.ReadString("/tmp/test.txt")
fmt.Println(content)
```

### ReadLines

```go
func ReadLines(path string) ([]string, error)
```

按行读取文件，返回字符串切片。

**示例:**

```go
lines, err := fileutil.ReadLines("/tmp/test.txt")
for _, line := range lines {
    fmt.Println(line)
}
```

### WriteBytes

```go
func WriteBytes(path string, data []byte) error
```

将字节数据写入文件（权限 0644）。

**示例:**

```go
err := fileutil.WriteBytes("/tmp/test.txt", []byte("Hello"))
```

### WriteString

```go
func WriteString(path string, s string) error
```

将字符串写入文件（权限 0644）。

**示例:**

```go
err := fileutil.WriteString("/tmp/test.txt", "Hello, World!")
```

### WriteLines

```go
func WriteLines(path string, lines []string) error
```

将字符串切片按行写入文件。

**示例:**

```go
err := fileutil.WriteLines("/tmp/test.txt", []string{"line1", "line2", "line3"})
```

### AppendBytes

```go
func AppendBytes(path string, data []byte) error
```

向文件追加字节数据。

**示例:**

```go
err := fileutil.AppendBytes("/tmp/test.txt", []byte("\nnew line"))
```

### AppendString

```go
func AppendString(path string, s string) error
```

向文件追加字符串。

**示例:**

```go
err := fileutil.AppendString("/tmp/test.txt", "\nappended text")
```

### CopyFile

```go
func CopyFile(src, dst string) error
```

复制文件，保留文件权限。

**示例:**

```go
err := fileutil.CopyFile("/tmp/a.txt", "/tmp/b.txt")
```

### MoveFile

```go
func MoveFile(src, dst string) error
```

移动文件，如果 Rename 失败则复制后删除。

**示例:**

```go
err := fileutil.MoveFile("/tmp/old.txt", "/tmp/new.txt")
```

### DeleteFile

```go
func DeleteFile(path string) error
```

删除文件。

**示例:**

```go
err := fileutil.DeleteFile("/tmp/test.txt")
```

### Mkdir

```go
func Mkdir(path string) error
```

创建目录（包括父目录）。

**示例:**

```go
err := fileutil.Mkdir("/tmp/a/b/c")
```

### MkdirAll

```go
func MkdirAll(path string) error
```

创建目录及其所有父目录。

**示例:**

```go
err := fileutil.MkdirAll("/tmp/a/b/c")
```

### ListFiles

```go
func ListFiles(dir string) ([]string, error)
```

列出目录中的所有文件（不包含子目录）。

**示例:**

```go
files, err := fileutil.ListFiles("/tmp")
for _, f := range files {
    fmt.Println(f)
}
```

### ListDirs

```go
func ListDirs(dir string) ([]string, error)
```

列出目录中的所有子目录。

**示例:**

```go
dirs, err := fileutil.ListDirs("/tmp")
for _, d := range dirs {
    fmt.Println(d)
}
```

### ListFilesRecursive

```go
func ListFilesRecursive(dir string) ([]string, error)
```

递归列出目录中的所有文件。

**示例:**

```go
files, err := fileutil.ListFilesRecursive("/tmp")
for _, f := range files {
    fmt.Println(f)
}
```

### Walk

```go
func Walk(root string, walkFn filepath.WalkFunc) error
```

遍历目录树。

**示例:**

```go
err := fileutil.Walk("/tmp", func(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    return nil
})
```

### Glob

```go
func Glob(pattern string) ([]string, error)
```

返回匹配通配符模式的文件列表。

**示例:**

```go
files, err := fileutil.Glob("/tmp/*.txt")
```

### TempDir

```go
func TempDir(dir, prefix string) (string, error)
```

创建临时目录。

**示例:**

```go
dir, err := fileutil.TempDir("", "myapp")
defer os.RemoveAll(dir)
```

### TempFile

```go
func TempFile(dir, prefix string) (*os.File, error)
```

创建临时文件。

**示例:**

```go
f, err := fileutil.TempFile("", "myapp")
defer os.Remove(f.Name())
```

### Remove

```go
func Remove(path string) error
```

删除文件或空目录。

**示例:**

```go
err := fileutil.Remove("/tmp/test.txt")
```

### RemoveAll

```go
func RemoveAll(path string) error
```

删除路径及其所有子项。

**示例:**

```go
err := fileutil.RemoveAll("/tmp/mydir")
```

### Rename

```go
func Rename(oldpath, newpath string) error
```

重命名文件或目录。

**示例:**

```go
err := fileutil.Rename("/tmp/old.txt", "/tmp/new.txt")
```

### Chmod

```go
func Chmod(path string, mode os.FileMode) error
```

修改文件权限。

**示例:**

```go
err := fileutil.Chmod("/tmp/test.txt", 0755)
```

### Chown

```go
func Chown(path string, uid, gid int) error
```

修改文件所有者。

**示例:**

```go
err := fileutil.Chown("/tmp/test.txt", 1000, 1000)
```

### IsSymlink

```go
func IsSymlink(path string) bool
```

判断路径是否为符号链接。

**示例:**

```go
if fileutil.IsSymlink("/tmp/link") {
    fmt.Println("这是一个符号链接")
}
```

### Symlink

```go
func Symlink(oldname, newname string) error
```

创建符号链接。

**示例:**

```go
err := fileutil.Symlink("/tmp/target", "/tmp/link")
```

### Readlink

```go
func Readlink(path string) (string, error)
```

获取符号链接指向的目标路径。

**示例:**

```go
target, err := fileutil.Readlink("/tmp/link")
```

### Md5

```go
func Md5(path string) (string, error)
```

计算文件的 MD5 哈希值（返回十六进制字符串）。

**示例:**

```go
hash, err := fileutil.Md5("/tmp/test.txt")
fmt.Println(hash)
```

### Sha256

```go
func Sha256(path string) (string, error)
```

计算文件的 SHA-256 哈希值（返回十六进制字符串）。

**示例:**

```go
hash, err := fileutil.Sha256("/tmp/test.txt")
fmt.Println(hash)
```

### Lines

```go
func Lines(path string) (int, error)
```

统计文件行数。

**示例:**

```go
count, err := fileutil.Lines("/tmp/test.txt")
fmt.Printf("文件共 %d 行\n", count)
```

### Head

```go
func Head(path string, n int) ([]string, error)
```

读取文件的前 n 行。

**示例:**

```go
lines, err := fileutil.Head("/tmp/test.txt", 10)
```

### Tail

```go
func Tail(path string, n int) ([]string, error)
```

读取文件的最后 n 行。

**示例:**

```go
lines, err := fileutil.Tail("/tmp/test.txt", 10)
```

### Grep

```go
func Grep(path string, pattern string) ([]string, error)
```

在文件中搜索匹配正则表达式的行。

**示例:**

```go
matches, err := fileutil.Grep("/tmp/test.go", `func\s+\w+`)
for _, line := range matches {
    fmt.Println(line)
}
```

### Replace

```go
func Replace(path string, old, new string) error
```

替换文件中的字符串内容。

**示例:**

```go
err := fileutil.Replace("/tmp/test.txt", "old", "new")
```

### Backup

```go
func Backup(path string) (string, error)
```

创建文件备份，备份文件名格式为 `{name}.backup.{timestamp}{ext}`。

**示例:**

```go
backupPath, err := fileutil.Backup("/tmp/test.txt")
fmt.Println(backupPath) // /tmp/test.backup.20240101120000.txt
```

### Normalize

```go
func Normalize(path string) string
```

将路径分隔符统一为正斜杠 `/`。

**示例:**

```go
path := fileutil.Normalize(`C:\Users\test\file.txt`)
fmt.Println(path) // C:/Users/test/file.txt
```

### ToURI

```go
func ToURI(path string) string
```

将路径转换为 file:// URI。

**示例:**

```go
uri := fileutil.ToURI("/tmp/test.txt")
fmt.Println(uri) // file:///tmp/test.txt
```

### UserHome

```go
func UserHome() (string, error)
```

获取用户主目录。

**示例:**

```go
home, err := fileutil.UserHome()
fmt.Println(home)
```

### WorkingDir

```go
func WorkingDir() (string, error)
```

获取当前工作目录。

**示例:**

```go
cwd, err := fileutil.WorkingDir()
fmt.Println(cwd)
```

### Separator

```go
func Separator() string
```

返回路径分隔符（Unix: `/`, Windows: `\`）。

**示例:**

```go
sep := fileutil.Separator()
```

### ListSeparator

```go
func ListSeparator() string
```

返回路径列表分隔符（Unix: `:`, Windows: `;`）。

**示例:**

```go
sep := fileutil.ListSeparator()
```
