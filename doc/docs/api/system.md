---
title: System
---

# System

系统信息模块，提供获取操作系统、硬件、进程、环境变量等系统信息的便捷函数。

## 导入

```go
import "github.com/gongm/gohutool/system"
```

## 基本信息

### OS

```go
func OS() string
```

返回操作系统名称（如 `linux`, `darwin`, `windows`）。

### Arch

```go
func Arch() string
```

返回系统架构（如 `amd64`, `arm64`）。

### NumCPU

```go
func NumCPU() int
```

返回 CPU 核心数。

### GoVersion

```go
func GoVersion() string
```

返回 Go 运行时版本。

**示例:**

```go
fmt.Printf("OS: %s, Arch: %s\n", system.OS(), system.Arch())
fmt.Printf("CPU: %d cores\n", system.NumCPU())
fmt.Printf("Go: %s\n", system.GoVersion())
```

## 主机信息

### Hostname

```go
func Hostname() (string, error)
```

返回主机名。

### Username

```go
func Username() (string, error)
```

返回当前用户名。

### UserHomeDir

```go
func UserHomeDir() (string, error)
```

返回用户主目录。

### WorkingDir

```go
func WorkingDir() (string, error)
```

返回当前工作目录。

**示例:**

```go
hostname, _ := system.Hostname()
username, _ := system.Username()
home, _ := system.UserHomeDir()
cwd, _ := system.WorkingDir()

fmt.Printf("主机: %s, 用户: %s\n", hostname, username)
fmt.Printf("主目录: %s\n", home)
fmt.Printf("工作目录: %s\n", cwd)
```

## 环境变量

### Env

```go
func Env(key string) string
```

获取环境变量值。

### EnvWithDefault

```go
func EnvWithDefault(key, defaultVal string) string
```

获取环境变量值，如果不存在则返回默认值。

**示例:**

```go
path := system.Env("PATH")
dbHost := system.EnvWithDefault("DB_HOST", "localhost")
```

### SetEnv

```go
func SetEnv(key, value string) error
```

设置环境变量。

### UnsetEnv

```go
func UnsetEnv(key string) error
```

删除环境变量。

### Environ

```go
func Environ() map[string]string
```

返回所有环境变量的 map。

**示例:**

```go
envs := system.Environ()
for k, v := range envs {
    fmt.Printf("%s=%s\n", k, v)
}
```

## 进程信息

### PID

```go
func PID() int
```

返回当前进程 ID。

### PPID

```go
func PPID() int
```

返回父进程 ID。

**示例:**

```go
fmt.Printf("PID: %d, PPID: %d\n", system.PID(), system.PPID())
```

## 操作系统判断

### IsWindows

```go
func IsWindows() bool
```

检查是否为 Windows 系统。

### IsLinux

```go
func IsLinux() bool
```

检查是否为 Linux 系统。

### IsMac

```go
func IsMac() bool
```

检查是否为 macOS 系统。

### Is64Bit

```go
func Is64Bit() bool
```

检查是否为 64 位架构。

**示例:**

```go
if system.IsWindows() {
    fmt.Println("运行在 Windows 上")
} else if system.IsLinux() {
    fmt.Println("运行在 Linux 上")
} else if system.IsMac() {
    fmt.Println("运行在 macOS 上")
}
```

## 系统资源信息

### MemoryInfo

```go
type MemoryInfo struct {
    Total     uint64
    Free      uint64
    Available uint64
}
```

### GetMemoryInfo

```go
func GetMemoryInfo() (*MemoryInfo, error)
```

获取内存信息。

**示例:**

```go
mem, err := system.GetMemoryInfo()
if err == nil {
    fmt.Printf("总内存: %d MB\n", mem.Total/1024/1024)
    fmt.Printf("可用内存: %d MB\n", mem.Available/1024/1024)
}
```

### DiskInfo

```go
type DiskInfo struct {
    Total uint64
    Free  uint64
    Used  uint64
}
```

### GetDiskInfo

```go
func GetDiskInfo(path string) (*DiskInfo, error)
```

获取指定路径的磁盘信息。

**示例:**

```go
disk, err := system.GetDiskInfo("/")
if err == nil {
    fmt.Printf("总空间: %d GB\n", disk.Total/1024/1024/1024)
    fmt.Printf("可用空间: %d GB\n", disk.Free/1024/1024/1024)
}
```

## 用户信息

### UserInfo

```go
type UserInfo struct {
    Username string
    UID      string
    GID      string
    HomeDir  string
    Name     string
}
```

### GetUserInfo

```go
func GetUserInfo() (*UserInfo, error)
```

获取当前用户的详细信息。

**示例:**

```go
info, err := system.GetUserInfo()
if err == nil {
    fmt.Printf("用户名: %s\n", info.Username)
    fmt.Printf("UID: %s, GID: %s\n", info.UID, info.GID)
    fmt.Printf("主目录: %s\n", info.HomeDir)
}
```

## 工具函数

### TempDir

```go
func TempDir() string
```

返回系统临时目录路径。

### LineSeparator

```go
func LineSeparator() string
```

返回当前系统的行分隔符（Windows: `\r\n`，其他: `\n`）。

### FileSeparator

```go
func FileSeparator() string
```

返回当前系统的文件路径分隔符（Windows: `\`，其他: `/`）。

**示例:**

```go
fmt.Printf("临时目录: %s\n", system.TempDir())
fmt.Printf("行分隔符: %q\n", system.LineSeparator())
fmt.Printf("路径分隔符: %q\n", system.FileSeparator())
```

## 完整示例

```go
package main

import (
    "fmt"

    "github.com/gongm/gohutool/system"
)

func main() {
    // 系统基本信息
    fmt.Println("=== 系统信息 ===")
    fmt.Printf("操作系统: %s\n", system.OS())
    fmt.Printf("架构: %s\n", system.Arch())
    fmt.Printf("CPU 核心数: %d\n", system.NumCPU())
    fmt.Printf("Go 版本: %s\n", system.GoVersion())
    fmt.Printf("64位: %v\n", system.Is64Bit())

    // 主机信息
    fmt.Println("\n=== 主机信息 ===")
    hostname, _ := system.Hostname()
    username, _ := system.Username()
    home, _ := system.UserHomeDir()
    fmt.Printf("主机名: %s\n", hostname)
    fmt.Printf("用户名: %s\n", username)
    fmt.Printf("主目录: %s\n", home)

    // 进程信息
    fmt.Println("\n=== 进程信息 ===")
    fmt.Printf("PID: %d\n", system.PID())
    fmt.Printf("PPID: %d\n", system.PPID())

    // 环境变量
    fmt.Println("\n=== 环境变量 ===")
    fmt.Printf("PATH: %s\n", system.Env("PATH"))
    fmt.Printf("GOPATH: %s\n", system.EnvWithDefault("GOPATH", "未设置"))
}
```
