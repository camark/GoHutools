---
title: Cron
---

# Cron

定时任务调度模块，支持标准 6 字段 cron 表达式和基于时间间隔的任务调度。可精确到秒级调度。

## 导入

```go
import "github.com/camark/GoHutools/cron"
```

## 常量

```go
const (
    EverySecond = "* * * * * *"     // 每秒
    EveryMinute = "0 * * * * *"     // 每分钟
    EveryHour   = "0 0 * * * *"     // 每小时
    EveryDay    = "0 0 0 * * *"     // 每天
    EveryWeek   = "0 0 0 * * 0"     // 每周
    EveryMonth  = "0 0 0 1 * *"     // 每月
    EveryYear   = "0 0 0 1 1 *"     // 每年
)
```

## Cron 表达式格式

使用 6 字段格式，精确到秒：

```
┌──────── 秒 (0-59)
│ ┌────── 分 (0-59)
│ │ ┌──── 时 (0-23)
│ │ │ ┌── 日 (1-31)
│ │ │ │ ┌ 月 (1-12)
│ │ │ │ │ ┌ 周 (0-6, 0=周日)
│ │ │ │ │ │
* * * * * *
```

支持的语法：
- `*` - 任意值
- `*/n` - 每隔 n 个单位
- `n-m` - 从 n 到 m 的范围
- `n,m,o` - 指定多个值

## 函数列表

### New

```go
func New() *Scheduler
```

创建新的调度器实例。

**示例:**

```go
s := cron.New()
```

### Scheduler.AddFunc

```go
func (s *Scheduler) AddFunc(cron string, fn func()) (int, error)
```

添加定时任务，使用 cron 表达式指定执行时间。返回任务 ID。

**示例:**

```go
s := cron.New()
id, err := s.AddFunc("*/5 * * * * *", func() {
    fmt.Println("每5秒执行一次")
})
if err != nil {
    log.Fatal(err)
}
```

### Scheduler.AddTask

```go
func (s *Scheduler) AddTask(task *Task) (int, error)
```

添加 Task 对象到调度器。返回任务 ID。

**示例:**

```go
s := cron.New()
task := &cron.Task{
    Cron: "0 */2 * * * *",
    Fn:   func() { fmt.Println("每2分钟执行") },
}
id, err := s.AddTask(task)
```

### Scheduler.Remove

```go
func (s *Scheduler) Remove(id int)
```

根据任务 ID 移除任务。

**示例:**

```go
s := cron.New()
id, _ := s.AddFunc("* * * * * *", func() { fmt.Println("test") })
s.Remove(id)
```

### Scheduler.Start

```go
func (s *Scheduler) Start()
```

启动调度器，开始执行已注册的定时任务。调度器在独立的 goroutine 中运行。

**示例:**

```go
s := cron.New()
s.AddFunc(cron.EverySecond, func() {
    fmt.Println("tick")
})
s.Start()
// ... 程序继续运行
```

### Scheduler.Stop

```go
func (s *Scheduler) Stop()
```

停止调度器，终止所有正在运行和待执行的任务。

### Scheduler.IsRunning

```go
func (s *Scheduler) IsRunning() bool
```

检查调度器是否正在运行。

### Scheduler.Tasks

```go
func (s *Scheduler) Tasks() []*Task
```

返回所有已注册的任务列表。

### Scheduler.Every

```go
func (s *Scheduler) Every(interval time.Duration) *IntervalTask
```

创建基于时间间隔的任务。返回 IntervalTask 对象，需要链式调用 Do 方法设置执行函数。

**示例:**

```go
s := cron.New()
s.Every(30 * time.Second).Do(func() {
    fmt.Println("每30秒执行一次")
})
s.Every(1 * time.Hour).Do(func() {
    fmt.Println("每小时执行一次")
})
s.Start()
```

### IntervalTask.Do

```go
func (t *IntervalTask) Do(fn func()) *IntervalTask
```

为间隔任务设置执行函数并注册到调度器。

## 完整示例

```go
package main

import (
    "fmt"
    "time"

    "github.com/camark/GoHutools/cron"
)

func main() {
    s := cron.New()

    // 每秒执行
    s.AddFunc(cron.EverySecond, func() {
        fmt.Println("每秒:", time.Now().Format("15:04:05"))
    })

    // 每5秒执行
    s.AddFunc("*/5 * * * * *", func() {
        fmt.Println("每5秒:", time.Now().Format("15:04:05"))
    })

    // 使用间隔方式
    s.Every(10 * time.Second).Do(func() {
        fmt.Println("间隔10秒:", time.Now().Format("15:04:05"))
    })

    s.Start()

    // 运行30秒后停止
    time.Sleep(30 * time.Second)
    s.Stop()
    fmt.Println("调度器已停止")
}
```
