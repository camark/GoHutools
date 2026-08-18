---
title: ioutil
---

# ioutil

IO 工具包，提供便捷的读写、复制、关闭等 IO 操作函数，是对 Go 标准库 `io` 和 `bufio` 的封装。

## 导入

```go
import "github.com/camark/GoHutools/ioutil"
```

## 函数列表

### ReadAll

```go
func ReadAll(r io.Reader) ([]byte, error)
```

从 Reader 读取全部内容并返回字节切片。

**示例:**

```go
data, err := ioutil.ReadAll(strings.NewReader("Hello, World!"))
fmt.Println(string(data)) // Hello, World!
```

### ReadAllString

```go
func ReadAllString(r io.Reader) (string, error)
```

从 Reader 读取全部内容并返回字符串。

**示例:**

```go
s, err := ioutil.ReadAllString(strings.NewReader("Hello"))
fmt.Println(s) // Hello
```

### ReadLines

```go
func ReadLines(r io.Reader) ([]string, error)
```

从 Reader 按行读取，返回字符串切片。

**示例:**

```go
lines, err := ioutil.ReadLines(strings.NewReader("line1\nline2\nline3"))
fmt.Println(lines) // [line1 line2 line3]
```

### WriteAll

```go
func WriteAll(w io.Writer, data []byte) error
```

将字节切片写入 Writer。

**示例:**

```go
var buf bytes.Buffer
err := ioutil.WriteAll(&buf, []byte("Hello"))
fmt.Println(buf.String()) // Hello
```

### WriteString

```go
func WriteString(w io.Writer, s string) error
```

将字符串写入 Writer。

**示例:**

```go
var buf bytes.Buffer
err := ioutil.WriteString(&buf, "Hello, World!")
```

### WriteLines

```go
func WriteLines(w io.Writer, lines []string) error
```

将字符串切片按行写入 Writer，每行之间用换行符分隔。

**示例:**

```go
var buf bytes.Buffer
err := ioutil.WriteLines(&buf, []string{"line1", "line2", "line3"})
```

### Copy

```go
func Copy(dst io.Writer, src io.Reader) (int64, error)
```

从 src 复制数据到 dst，返回复制的字节数。

**示例:**

```go
src := strings.NewReader("Hello")
var dst bytes.Buffer
n, err := ioutil.Copy(&dst, src)
```

### CopyN

```go
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error)
```

从 src 复制最多 n 个字节到 dst。

**示例:**

```go
src := strings.NewReader("Hello, World!")
var dst bytes.Buffer
n, err := ioutil.CopyN(&dst, src, 5)
fmt.Println(dst.String()) // Hello
```

### CopyBuffer

```go
func CopyBuffer(dst io.Writer, src io.Reader, buf []byte) (int64, error)
```

使用指定缓冲区从 src 复制数据到 dst。

**示例:**

```go
buf := make([]byte, 1024)
n, err := ioutil.CopyBuffer(dst, src, buf)
```

### Drain

```go
func Drain(r io.Reader) error
```

读取并丢弃 Reader 中的全部内容。

**示例:**

```go
err := ioutil.Drain(response.Body)
```

### ReadBytes

```go
func ReadBytes(r io.Reader, delim byte) ([]byte, error)
```

从 Reader 读取直到遇到指定分隔符，返回字节切片（包含分隔符）。

**示例:**

```go
data, err := ioutil.ReadBytes(strings.NewReader("Hello\nWorld"), '\n')
fmt.Println(string(data)) // Hello\n
```

### ReadString

```go
func ReadString(r io.Reader, delim byte) (string, error)
```

从 Reader 读取直到遇到指定分隔符，返回字符串（包含分隔符）。

**示例:**

```go
s, err := ioutil.ReadString(strings.NewReader("Hello\nWorld"), '\n')
fmt.Println(s) // Hello\n
```

### TeeReader

```go
func TeeReader(r io.Reader, w io.Writer) io.Reader
```

返回一个 Reader，它在从 r 读取的同时将内容写入 w。

**示例:**

```go
var buf bytes.Buffer
r := ioutil.TeeReader(source, &buf)
data, _ := ioutil.ReadAll(r)
// data 包含内容，buf 也包含相同内容
```

### LimitReader

```go
func LimitReader(r io.Reader, n int64) io.Reader
```

返回一个最多读取 n 个字节的 Reader。

**示例:**

```go
r := ioutil.LimitReader(source, 100)
data, _ := ioutil.ReadAll(r) // 最多读取 100 字节
```

### SectionReader

```go
func SectionReader(r io.ReaderAt, off int64, n int64) io.Reader
```

返回一个从指定偏移量读取 n 字节的 Reader。

**示例:**

```go
r := ioutil.SectionReader(readerAt, 10, 50)
data, _ := ioutil.ReadAll(r)
```

### NewBuffer

```go
func NewBuffer(data []byte) *bytes.Buffer
```

使用指定字节数据创建新的 Buffer。

**示例:**

```go
buf := ioutil.NewBuffer([]byte("Hello"))
fmt.Println(buf.String()) // Hello
```

### NewBufferString

```go
func NewBufferString(s string) *bytes.Buffer
```

使用字符串创建新的 Buffer。

**示例:**

```go
buf := ioutil.NewBufferString("Hello")
fmt.Println(buf.String()) // Hello
```

### NewReader

```go
func NewReader(data []byte) *bytes.Reader
```

使用字节数据创建新的 Reader。

**示例:**

```go
r := ioutil.NewReader([]byte("Hello"))
data, _ := ioutil.ReadAll(r)
```

### NewReaderString

```go
func NewReaderString(s string) *strings.Reader
```

使用字符串创建新的 Reader。

**示例:**

```go
r := ioutil.NewReaderString("Hello")
data, _ := ioutil.ReadAll(r)
```

### BufferedWriter

```go
func BufferedWriter(w io.Writer) *bufio.Writer
```

创建带缓冲的 Writer。

**示例:**

```go
bw := ioutil.BufferedWrite(os.Stdout)
ioutil.WriteString(bw, "Hello")
bw.Flush()
```

### BufferedReadString

```go
func BufferedReadString(r io.Reader) *bufio.Reader
```

创建带缓冲的 Reader。

**示例:**

```go
br := ioutil.BufferedReadString(os.Stdin)
line, _ := br.ReadString('\n')
```

### Close

```go
func Close(c io.Closer) error
```

安全关闭 Closer，如果为 nil 则忽略。

**示例:**

```go
defer ioutil.Close(file)
```

### CloseQuietly

```go
func CloseQuietly(c io.Closer)
```

关闭 Closer 并忽略错误。

**示例:**

```go
defer ioutil.CloseQuietly(response.Body)
```

### Flush

```go
func Flush(w io.Writer) error
```

如果 Writer 实现了 Flush 方法则调用之。

**示例:**

```go
err := ioutil.Flush(bufferedWriter)
```

### MultiReader

```go
func MultiReader(readers ...io.Reader) io.Reader
```

返回一个将多个 Reader 逻辑串联的 Reader。

**示例:**

```go
r1 := strings.NewReader("Hello, ")
r2 := strings.NewReader("World!")
r := ioutil.MultiReader(r1, r2)
data, _ := ioutil.ReadAll(r) // Hello, World!
```

### MultiWriter

```go
func MultiWriter(writers ...io.Writer) io.Writer
```

返回一个将写入内容复制到所有 Writer 的 Writer。

**示例:**

```go
var buf1, buf2 bytes.Buffer
w := ioutil.MultiWriter(&buf1, &buf2)
ioutil.WriteString(w, "Hello")
// buf1 和 buf2 都包含 "Hello"
```

### Pipe

```go
func Pipe() (*io.PipeReader, *io.PipeWriter)
```

创建同步内存管道。

**示例:**

```go
pr, pw := ioutil.Pipe()
go func() {
    ioutil.WriteString(pw, "Hello")
    pw.Close()
}()
data, _ := ioutil.ReadAll(pr)
```

### ReadFull

```go
func ReadFull(r io.Reader, buf []byte) (int, error)
```

从 Reader 精确读取 len(buf) 个字节到 buf 中。

**示例:**

```go
buf := make([]byte, 10)
n, err := ioutil.ReadFull(reader, buf)
```

### ReadAtLeast

```go
func ReadAtLeast(r io.Reader, buf []byte, min int) (int, error)
```

从 Reader 读取至少 min 个字节到 buf 中。

**示例:**

```go
buf := make([]byte, 100)
n, err := ioutil.ReadAtLeast(reader, buf, 10)
```
