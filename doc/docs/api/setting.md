---
title: Setting
---

# Setting

配置管理模块，支持 INI 格式的配置文件读写，支持分节、注释、多种数据类型获取。线程安全。

## 导入

```go
import "github.com/camark/GoHutools/setting"
```

## 配置文件格式

支持标准 INI 格式：

```ini
# 注释行
; 也是注释
key1 = value1
key2 = "带引号的值"

[section1]
host = localhost
port = 8080

[section2]
debug = true
```

## 函数列表

### New

```go
func New() *Setting
```

创建空的配置实例。

**示例:**

```go
s := setting.New()
s.Set("key", "value")
```

### LoadFile

```go
func LoadFile(path string) (*Setting, error)
```

从文件加载配置。

**示例:**

```go
s, err := setting.LoadFile("config.ini")
if err != nil {
    log.Fatal(err)
}
```

### LoadString

```go
func LoadString(str string) (*Setting, error)
```

从字符串加载配置。

**示例:**

```go
config := `
host = localhost
port = 8080

[database]
driver = mysql
dsn = root:123@tcp(localhost:3306)/test
`
s, err := setting.LoadString(config)
```

### Setting.Load

```go
func (s *Setting) Load(path string) error
```

从文件加载配置到已有实例。

### Setting.Reload

```go
func (s *Setting) Reload(path string) error
```

重新加载配置文件，会清空原有数据。

### Setting.Get

```go
func (s *Setting) Get(key string) string
```

获取默认节中的配置值。

**示例:**

```go
s, _ := setting.LoadFile("config.ini")
host := s.Get("host")
```

### Setting.GetWithDefault

```go
func (s *Setting) GetWithDefault(key, defaultVal string) string
```

获取默认节中的配置值，如果不存在则返回默认值。

**示例:**

```go
port := s.GetWithDefault("port", "8080")
```

### Setting.GetSection

```go
func (s *Setting) GetSection(section, key string) string
```

获取指定节中的配置值。

**示例:**

```go
driver := s.GetSection("database", "driver")
```

### Setting.GetSectionWithDefault

```go
func (s *Setting) GetSectionWithDefault(section, key, defaultVal string) string
```

获取指定节中的配置值，如果不存在则返回默认值。

### Setting.Set

```go
func (s *Setting) Set(key, value string)
```

在默认节中设置配置值。

### Setting.SetSection

```go
func (s *Setting) SetSection(section, key, value string)
```

在指定节中设置配置值。

**示例:**

```go
s := setting.New()
s.Set("app_name", "MyApp")
s.SetSection("database", "host", "localhost")
s.SetSection("database", "port", "3306")
```

### Setting.Has

```go
func (s *Setting) Has(key string) bool
```

检查默认节中是否存在指定键。

### Setting.HasSection

```go
func (s *Setting) HasSection(section string) bool
```

检查是否存在指定节。

### Setting.Delete

```go
func (s *Setting) Delete(key string)
```

删除默认节中的指定键。

### Setting.DeleteSection

```go
func (s *Setting) DeleteSection(section string)
```

删除指定节。

### Setting.Sections

```go
func (s *Setting) Sections() []string
```

返回所有节名称的列表。

### Setting.Keys

```go
func (s *Setting) Keys(section string) []string
```

返回指定节中所有键的列表。

### Setting.ToMap

```go
func (s *Setting) ToMap() map[string]map[string]string
```

将配置转换为嵌套 map。

### Setting.Save

```go
func (s *Setting) Save(path string) error
```

将配置保存到文件。

**示例:**

```go
s := setting.New()
s.Set("host", "localhost")
s.SetSection("database", "port", "3306")
s.Save("config.ini")
```

### Setting.SaveString

```go
func (s *Setting) SaveString() string
```

将配置序列化为字符串。

## 类型转换方法

### Setting.GetInt

```go
func (s *Setting) GetInt(key string) (int, error)
```

获取整数值。

### Setting.GetIntWithDefault

```go
func (s *Setting) GetIntWithDefault(key string, defaultVal int) int
```

获取整数值，如果不存在或解析失败则返回默认值。

**示例:**

```go
port := s.GetIntWithDefault("port", 8080)
```

### Setting.GetBool

```go
func (s *Setting) GetBool(key string) (bool, error)
```

获取布尔值。支持的真值: `1`, `t`, `true`, `yes`, `on`。

### Setting.GetBoolWithDefault

```go
func (s *Setting) GetBoolWithDefault(key string, defaultVal bool) bool
```

获取布尔值，如果不存在或解析失败则返回默认值。

**示例:**

```go
debug := s.GetBoolWithDefault("debug", false)
```

### Setting.GetFloat

```go
func (s *Setting) GetFloat(key string) (float64, error)
```

获取浮点数值。

### Setting.GetFloatWithDefault

```go
func (s *Setting) GetFloatWithDefault(key string, defaultVal float64) float64
```

获取浮点数值，如果不存在或解析失败则返回默认值。

### Setting.GetSlice

```go
func (s *Setting) GetSlice(key string) []string
```

获取切片值（逗号分隔）。

**示例:**

```ini
allowed_origins = http://localhost, http://example.com
```

```go
origins := s.GetSlice("allowed_origins")
// ["http://localhost", "http://example.com"]
```

## 完整示例

```go
package main

import (
    "fmt"
    "log"

    "github.com/camark/GoHutools/setting"
)

func main() {
    // 从字符串加载配置
    config := `
app_name = MyService
port = 8080
debug = true

[database]
driver = mysql
host = localhost
port = 3306
max_connections = 100

[redis]
host = localhost
port = 6379
`
    s, err := setting.LoadString(config)
    if err != nil {
        log.Fatal(err)
    }

    // 获取配置值
    fmt.Println("应用名:", s.Get("app_name"))
    fmt.Println("端口:", s.GetIntWithDefault("port", 8080))
    fmt.Println("调试模式:", s.GetBoolWithDefault("debug", false))

    // 获取数据库配置
    fmt.Println("数据库驱动:", s.GetSection("database", "driver"))
    fmt.Println("数据库主机:", s.GetSection("database", "host"))

    // 列出所有节
    fmt.Println("所有节:", s.Sections())

    // 修改并保存
    s.SetSection("database", "host", "192.168.1.100")
    s.Save("config.ini")
}
```
