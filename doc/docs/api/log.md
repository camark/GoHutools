---
title: Log
---

# Log

日志门面模块，提供统一的日志接口和默认实现。支持多级日志、结构化字段日志和自定义输出。类似 Java 中的 SLF4J。

## 导入

```go
import "github.com/gongm/gohutool/log"
```

## 日志级别

```go
const (
    LevelTrace Level = iota  // 跟踪
    LevelDebug               // 调试
    LevelInfo                // 信息
    LevelWarn                // 警告
    LevelError               // 错误
    LevelFatal               // 致命（输出后退出程序）
    LevelOff                 // 关闭
)
```

## 接口

### Logger

```go
type Logger interface {
    Trace(format string, args ...interface{})
    Debug(format string, args ...interface{})
    Info(format string, args ...interface{})
    Warn(format string, args ...interface{})
    Error(format string, args ...interface{})
    Fatal(format string, args ...interface{})
    SetLevel(level Level)
    GetLevel() Level
}
```

## 创建 Logger

### New

```go
func New() Logger
```

创建默认日志器，输出到 stderr，级别为 Debug。

**示例:**

```go
logger := log.New()
logger.Info("应用启动")
```

### NewWithOutput

```go
func NewWithOutput(output io.Writer) Logger
```

创建自定义输出的日志器。

**示例:**

```go
file, _ := os.Create("app.log")
logger := log.NewWithOutput(file)
logger.Info("日志写入文件")
```

### NewWithLevel

```go
func NewWithLevel(level Level) Logger
```

创建指定级别的日志器。

**示例:**

```go
logger := log.NewWithLevel(log.LevelWarn)
logger.Debug("这条不会输出") // 低于 Warn 级别，被过滤
logger.Warn("这条会输出")
```

## 日志方法

### Trace

```go
func (l Logger) Trace(format string, args ...interface{})
```

输出 Trace 级别日志。用于最详细的追踪信息。

### Debug

```go
func (l Logger) Debug(format string, args ...interface{})
```

输出 Debug 级别日志。用于调试信息。

### Info

```go
func (l Logger) Info(format string, args ...interface{})
```

输出 Info 级别日志。用于一般信息。

### Warn

```go
func (l Logger) Warn(format string, args ...interface{})
```

输出 Warn 级别日志。用于警告信息。

### Error

```go
func (l Logger) Error(format string, args ...interface{})
```

输出 Error 级别日志。用于错误信息。

### Fatal

```go
func (l Logger) Fatal(format string, args ...interface{})
```

输出 Fatal 级别日志并调用 `os.Exit(1)` 退出程序。

### SetLevel / GetLevel

```go
func (l Logger) SetLevel(level Level)
func (l Logger) GetLevel() Level
```

设置/获取日志级别。

## 全局函数

提供便捷的全局日志函数，使用默认日志器实例。

### Trace

```go
func Trace(format string, args ...interface{})
```

### Debug

```go
func Debug(format string, args ...interface{})
```

### Info

```go
func Info(format string, args ...interface{})
```

### Warn

```go
func Warn(format string, args ...interface{})
```

### Error

```go
func Error(format string, args ...interface{})
```

### Fatal

```go
func Fatal(format string, args ...interface{})
```

**示例:**

```go
log.Info("服务器启动于端口 %d", 8080)
log.Error("数据库连接失败: %v", err)
log.Fatal("无法加载配置文件: %v", err) // 会退出程序
```

### SetLevel

```go
func SetLevel(level Level)
```

设置全局日志级别。

### GetLevel

```go
func GetLevel() Level
```

获取全局日志级别。

### SetOutput

```go
func SetOutput(w io.Writer)
```

设置全局日志输出目标。

**示例:**

```go
file, _ := os.Create("app.log")
log.SetOutput(file)
log.Info("写入文件")
```

### GetLogger

```go
func GetLogger() Logger
```

获取全局日志器实例。

### SetLogger

```go
func SetLogger(l Logger)
```

设置全局日志器实例，可用于替换为自定义实现。

### WithFields

```go
func WithFields(fields map[string]interface{}) Logger
```

创建带有结构化字段的日志器。字段会附加在日志消息之后。

**示例:**

```go
logger := log.WithFields(map[string]interface{}{
    "request_id": "abc-123",
    "user_id":    1001,
})
logger.Info("用户登录")
// 输出: [2024-01-01 12:00:00.000] [INFO] main.go:10: 用户登录 [request_id=abc-123, user_id=1001]
```

## 日志格式

默认输出格式：

```
[时间戳] [级别] 文件名:行号: 消息
```

示例输出：

```
[2024-01-01 12:00:00.000] [INFO] main.go:10: 服务器启动成功
[2024-01-01 12:00:01.500] [ERROR] db.go:25: 连接超时 [host=localhost:3306]
```

## 完整示例

```go
package main

import (
    "os"

    "github.com/gongm/gohutool/log"
)

func main() {
    // 使用全局函数
    log.Info("应用启动")
    log.Debug("调试信息: %s", "detail")

    // 设置日志级别
    log.SetLevel(log.LevelInfo)
    log.Debug("这条不会输出") // 低于 Info 级别

    // 使用结构化字段
    requestLog := log.WithFields(map[string]interface{}{
        "method": "GET",
        "path":   "/api/users",
        "status": 200,
    })
    requestLog.Info("请求完成")

    // 替换为文件输出
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal("无法打开日志文件: %v", err)
    }
    defer file.Close()
    log.SetOutput(file)
    log.Info("日志已切换到文件")
}
```
