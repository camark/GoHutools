---
title: 示例代码
---

# GoHutool 示例代码

本页提供 GoHutool 各模块的实际使用示例，帮助你快速上手。

## 缓存 (Cache)

### LRU 缓存 - API 响应缓存

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"

    "github.com/camark/GoHutools/cache"
)

type APIResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

var responseCache = cache.NewLRU(1000)

func fetchFromAPI(url string) (*APIResponse, error) {
    // 先查缓存
    if cached, ok := responseCache.Get(url); ok {
        return cached.(*APIResponse), nil
    }

    // 模拟 API 调用
    resp := &APIResponse{
        Code:    200,
        Message: "success",
        Data:    map[string]string{"name": "test"},
    }

    // 写入缓存，5 分钟过期
    responseCache.SetWithExpire(url, resp, 5*time.Minute)
    return resp, nil
}

func main() {
    resp, _ := fetchFromAPI("/api/users/1")
    data, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Println(string(data))

    // 第二次调用命中缓存
    resp, _ = fetchFromAPI("/api/users/1")
    fmt.Println("缓存命中:", responseCache.Contains("/api/users/1"))
}
```

### 时间缓存 - Session 管理

```go
package main

import (
    "fmt"
    "time"

    "github.com/camark/GoHutools/cache"
)

func main() {
    sessions := cache.NewTimed(30 * time.Minute)

    // 用户登录
    sessions.Set("session:abc123", map[string]interface{}{
        "user_id":  1001,
        "username": "张三",
        "role":     "admin",
    })

    // 查询 session
    if data, ok := sessions.Get("session:abc123"); ok {
        user := data.(map[string]interface{})
        fmt.Printf("欢迎回来, %s\n", user["username"])
    }

    fmt.Println("活跃 session 数:", sessions.Size())
}
```

## 定时任务 (Cron)

### Web 应用定时任务

```go
package main

import (
    "fmt"
    "time"

    "github.com/camark/GoHutools/cron"
)

func main() {
    s := cron.New()

    // 每秒打印心跳
    s.AddFunc(cron.EverySecond, func() {
        fmt.Printf("[%s] 心跳\n", time.Now().Format("15:04:05"))
    })

    // 每5分钟清理过期数据
    s.AddFunc("0 */5 * * * *", func() {
        fmt.Printf("[%s] 清理过期数据\n", time.Now().Format("15:04:05"))
    })

    // 每小时生成报告
    s.AddFunc("0 0 * * * *", func() {
        fmt.Printf("[%s] 生成小时报告\n", time.Now().Format("15:04:05"))
    })

    // 使用间隔方式：每30秒同步数据
    s.Every(30 * time.Second).Do(func() {
        fmt.Printf("[%s] 同步远程数据\n", time.Now().Format("15:04:05"))
    })

    s.Start()
    fmt.Println("调度器已启动")

    time.Sleep(2 * time.Minute)
    s.Stop()
}
```

## 验证码 (Captcha)

### Web 登录验证码

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/camark/GoHutools/captcha"
)

// 简单的内存存储（生产环境请使用 Redis）
var captchaStore = make(map[string]string)

func main() {
    // 生成验证码图片
    http.HandleFunc("/captcha/img", func(w http.ResponseWriter, r *http.Request) {
        c := captcha.NewLine().
            SetWidth(200).
            SetHeight(80).
            SetLength(4).
            SetLineCount(6)

        code, pngBytes, err := c.GeneratePNG()
        if err != nil {
            http.Error(w, "生成失败", 500)
            return
        }

        // 存储验证码（使用 session ID 作为 key）
        sessionID := "session_123"
        captchaStore[sessionID] = code

        w.Header().Set("Content-Type", "image/png")
        w.Write(pngBytes)
    })

    // 验证
    http.HandleFunc("/captcha/verify", func(w http.ResponseWriter, r *http.Request) {
        input := r.FormValue("code")
        sessionID := "session_123"

        expected, exists := captchaStore[sessionID]
        if !exists {
            fmt.Fprintln(w, "验证码已过期")
            return
        }

        c := captcha.New()
        if c.VerifyIgnoreCase(input, expected) {
            fmt.Fprintln(w, "验证通过")
            delete(captchaStore, sessionID)
        } else {
            fmt.Fprintln(w, "验证码错误")
        }
    })

    fmt.Println("服务器启动于 :8080")
    http.ListenAndServe(":8080", nil)
}
```

## 布隆过滤器 (Bloom Filter)

### URL 去重

```go
package main

import (
    "fmt"

    "github.com/camark/GoHutools/bloom"
)

func main() {
    // 创建布隆过滤器：预期 100 万个 URL，假阳性率 0.1%
    visited := bloom.New(1000000, 0.001)

    urls := []string{
        "https://example.com/page1",
        "https://example.com/page2",
        "https://example.com/page3",
        "https://example.com/page1", // 重复
    }

    for _, url := range urls {
        if visited.ContainsString(url) {
            fmt.Printf("[跳过] %s 已访问\n", url)
        } else {
            visited.AddString(url)
            fmt.Printf("[新URL] %s 已加入队列\n", url)
        }
    }

    fmt.Printf("\n统计: 共 %d 个唯一 URL\n", visited.Count())
    fmt.Printf("假阳性率: %.6f\n", visited.FalsePositiveRate())
}
```

### 敏感词过滤

```go
package main

import (
    "fmt"

    "github.com/camark/GoHutools/bloom"
)

func main() {
    sensitiveWords := bloom.New(50000, 0.01)

    // 加载敏感词库
    words := []string{"敏感词1", "敏感词2", "违禁词1", "违禁词2"}
    for _, w := range words {
        sensitiveWords.AddString(w)
    }

    // 检查内容
    content := "这段文字包含敏感词1"
    for _, w := range words {
        if sensitiveWords.ContainsString(w) {
            fmt.Printf("发现敏感词: %s\n", w)
        }
    }

    fmt.Printf("内容长度: %d\n", len(content))
}
```

## 对象池 (Pool)

### HTTP 请求对象池

```go
package main

import (
    "bytes"
    "fmt"
    "sync"

    "github.com/camark/GoHutools/pool"
)

type Request struct {
    Method  string
    URL     string
    Headers map[string]string
    Body    *bytes.Buffer
}

func main() {
    requestPool := pool.NewWithReset(
        func() interface{} {
            return &Request{
                Headers: make(map[string]string),
                Body:    &bytes.Buffer{},
            }
        },
        func(v interface{}) {
            r := v.(*Request)
            r.Method = ""
            r.URL = ""
            for k := range r.Headers {
                delete(r.Headers, k)
            }
            r.Body.Reset()
        },
    )

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            req := requestPool.Get().(*Request)
            req.Method = "GET"
            req.URL = fmt.Sprintf("/api/users/%d", id)

            // 模拟处理
            _ = req.URL

            requestPool.Put(req)
        }(i)
    }
    wg.Wait()

    fmt.Printf("创建的对象总数: %d\n", requestPool.Size())
}
```

### 字节缓冲区池

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/camark/GoHutools/pool"
)

var bufPool = pool.NewBufferPool(4096)

func toJSON(v interface{}) ([]byte, error) {
    buf := bufPool.Get()
    defer bufPool.Put(buf)

    encoder := json.NewEncoder(&bufferWriter{buf: &buf})
    err := encoder.Encode(v)
    return buf, err
}

type bufferWriter struct {
    buf *[]byte
}

func (w *bufferWriter) Write(p []byte) (n int, err error) {
    *w.buf = append(*w.buf, p...)
    return len(p), nil
}

func main() {
    data := map[string]interface{}{
        "name": "GoHutool",
        "version": "1.0",
    }

    result, _ := toJSON(data)
    fmt.Println(string(result))
}
```

## 日志 (Log)

### Web 请求日志

```go
package main

import (
    "net/http"
    "time"

    "github.com/camark/GoHutools/log"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        requestLog := log.WithFields(map[string]interface{}{
            "method":     r.Method,
            "path":       r.URL.Path,
            "remote_ip":  r.RemoteAddr,
        })

        requestLog.Info("请求开始")

        next.ServeHTTP(w, r)

        requestLog.WithFields(map[string]interface{}{
            "duration": time.Since(start).String(),
        }).Info("请求完成")
    })
}

func main() {
    log.SetLevel(log.LevelInfo)

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello"))
    })

    log.Info("服务器启动于 :8080")
    http.ListenAndServe(":8080", loggingMiddleware(mux))
}
```

### 结构化日志

```go
package main

import (
    "github.com/camark/GoHutools/log"
)

func main() {
    // 带字段的日志
    dbLog := log.WithFields(map[string]interface{}{
        "module": "database",
        "host":   "localhost",
        "port":   3306,
    })

    dbLog.Info("数据库连接成功")
    dbLog.Warn("连接池使用率过高: %d%%", 85)
    dbLog.Error("查询超时: %v", "SELECT * FROM users")

    // 设置日志级别
    log.SetLevel(log.LevelWarn)
    log.Debug("这条不会输出")
    log.Warn("这条会输出")
}
```

## 配置 (Setting)

### 应用配置管理

```go
package main

import (
    "fmt"
    "log"

    "github.com/camark/GoHutools/setting"
)

func main() {
    config := `
app_name = MyService
version = 1.0.0
debug = true

[server]
host = 0.0.0.0
port = 8080
read_timeout = 30

[database]
driver = mysql
host = localhost
port = 3306
username = root
password = secret
database = mydb
max_connections = 100

[redis]
host = localhost
port = 6379
password =
db = 0
`

    s, err := setting.LoadString(config)
    if err != nil {
        log.Fatal(err)
    }

    // 读取配置
    fmt.Printf("应用: %s v%s\n", s.Get("app_name"), s.Get("version"))
    fmt.Printf("调试模式: %v\n", s.GetBoolWithDefault("debug", false))

    // 服务器配置
    fmt.Printf("服务器: %s:%d\n",
        s.GetSection("server", "host"),
        s.GetIntWithDefault("port", 8080),
    )

    // 数据库配置
    fmt.Printf("数据库: %s://%s:%s/%s\n",
        s.GetSection("database", "driver"),
        s.GetSection("database", "host"),
        s.GetSection("database", "port"),
        s.GetSection("database", "database"),
    )

    // 修改并保存
    s.SetSection("server", "port", "9090")
    s.Save("config.ini")
}
```

## 系统信息 (System)

### 系统监控

```go
package main

import (
    "fmt"
    "runtime"

    "github.com/camark/GoHutools/system"
)

func main() {
    fmt.Println("=== 系统信息 ===")
    fmt.Printf("OS: %s/%s\n", system.OS(), system.Arch())
    fmt.Printf("CPU: %d cores\n", system.NumCPU())
    fmt.Printf("Go: %s\n", system.GoVersion())
    fmt.Printf("PID: %d\n", system.PID())

    hostname, _ := system.Hostname()
    username, _ := system.Username()
    fmt.Printf("主机: %s, 用户: %s\n", hostname, username)

    // 内存信息
    fmt.Println("\n=== 内存信息 ===")
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("程序内存: %d MB\n", m.Alloc/1024/1024)

    mem, err := system.GetMemoryInfo()
    if err == nil {
        fmt.Printf("系统内存: 总计 %d MB, 可用 %d MB\n",
            mem.Total/1024/1024,
            mem.Available/1024/1024,
        )
    }

    // 环境信息
    fmt.Println("\n=== 环境变量 ===")
    fmt.Printf("GOPATH: %s\n", system.EnvWithDefault("GOPATH", "未设置"))
    fmt.Printf("HOME: %s\n", system.EnvWithDefault("HOME", "未设置"))
}
```

## 对象工具 (ObjectUtil)

### 表单数据处理

```go
package main

import (
    "fmt"

    "github.com/camark/GoHutools/objectutil"
)

type FormData struct {
    Name    string
    Age     int
    Email   string
    Address string
}

func processForm(data *FormData) {
    // 处理默认值
    name := objectutil.DefaultIfEmpty(data.Name, "匿名").(string)
    email := objectutil.DefaultIfEmpty(data.Email, "未填写").(string)
    address := objectutil.DefaultIfEmpty(data.Address, "未填写").(string)

    fmt.Printf("姓名: %s\n", name)
    fmt.Printf("邮箱: %s\n", email)
    fmt.Printf("地址: %s\n", address)

    // 检查必填项
    if objectutil.IsEmpty(data.Name) {
        fmt.Println("警告: 姓名不能为空")
    }
}

func main() {
    // 完整数据
    processForm(&FormData{
        Name:    "张三",
        Age:     25,
        Email:   "zhangsan@example.com",
        Address: "北京市",
    })

    fmt.Println("---")

    // 部分数据
    processForm(&FormData{
        Name: "李四",
        Age:  30,
    })
}
```

## 正则工具 (RegexUtil)

### 日志解析

```go
package main

import (
    "fmt"

    "github.com/camark/GoHutools/regexutil"
)

func main() {
    logLine := `2024-01-15 14:30:25 [ERROR] db.go:123: Connection refused to 192.168.1.100:3306`

    // 提取日期时间
    dateTime := regexutil.Find(logLine, `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)
    fmt.Printf("时间: %s\n", dateTime)

    // 提取日志级别
    level := regexutil.ExtractGroup(logLine, `\[(\w+)\]`, 1)
    fmt.Printf("级别: %s\n", level)

    // 提取 IP 地址
    ip := regexutil.Find(logLine, `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
    fmt.Printf("IP: %s\n", ip)

    // 使用命名分组
    pattern := `(?P<time>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(?P<level>\w+)\] (?P<file>[\w.]+):(?P<line>\d+): (?P<message>.*)`
    groups := regexutil.FindNamedGroups(logLine, pattern)
    if groups != nil {
        fmt.Printf("解析结果: %+v\n", groups)
    }
}
```

### 数据脱敏

```go
package main

import (
    "fmt"

    "github.com/camark/GoHutools/regexutil"
)

func main() {
    // 手机号脱敏
    phone := "手机号: 13812345678"
    masked := regexutil.Replace(phone, `1(\d{2})\d{4}(\d{4})`, "1$1****$2")
    fmt.Println(masked) // 手机号: 138****5678

    // 身份证脱敏
    idCard := "身份证: 110101199001011234"
    masked = regexutil.Replace(idCard, `(\d{6})\d{8}(\d{4})`, "$1********$2")
    fmt.Println(masked) // 身份证: 110101********1234

    // 邮箱脱敏
    email := "邮箱: zhangsan@example.com"
    masked = regexutil.Replace(email, `(\w{1,3})\w+(@\w+\.\w+)`, "$1***$2")
    fmt.Println(masked) // 邮箱: zha***@example.com

    // 银行卡脱敏
    card := "卡号: 6222021234567890123"
    masked = regexutil.Replace(card, `(\d{4})\d{8,12}(\d{4})`, "$1 **** **** $2")
    fmt.Println(masked) // 卡号: 6222 **** **** 0123
}
```

### 表单验证

```go
package main

import (
    "fmt"

    "github.com/camark/GoHutools/regexutil"
)

type ValidationResult struct {
    Field   string
    Value   string
    Valid   bool
    Message string
}

func validate(value, pattern, fieldName, errorMsg string) ValidationResult {
    valid := regexutil.IsMatch(value, pattern)
    msg := "验证通过"
    if !valid {
        msg = errorMsg
    }
    return ValidationResult{
        Field:   fieldName,
        Value:   value,
        Valid:   valid,
        Message: msg,
    }
}

func main() {
    tests := []struct {
        value   string
        pattern string
        field   string
        errMsg  string
    }{
        {"user@example.com", regexutil.PatternEmail, "邮箱", "邮箱格式不正确"},
        {"13812345678", regexutil.PatternMobile, "手机号", "手机号格式不正确"},
        {"110101199001011234", regexutil.PatternIDCard, "身份证", "身份证格式不正确"},
        {"100000", regexutil.PatternZipCode, "邮编", "邮编格式不正确"},
        {"abc123", regexutil.PatternAlphaNum, "用户名", "用户名只能包含字母和数字"},
    }

    for _, t := range tests {
        result := validate(t.value, t.pattern, t.field, t.errMsg)
        status := "PASS"
        if !result.Valid {
            status = "FAIL"
        }
        fmt.Printf("[%s] %s: %s (%s)\n", status, result.Field, result.Value, result.Message)
    }
}
```
