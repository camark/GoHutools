# GoHutool

**中文** | [English](README.en.md)

Go 版本的 [Hutool](https://github.com/dromara/hutool) 工具库，功能对标 Java 版 Hutool，使用 Go 泛型实现。

[![CI](https://img.shields.io/github/actions/workflow/status/camark/GoHutools/ci.yml?branch=main&logo=github)](https://github.com/camark/GoHutools/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8.svg?logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)

## 特性

- 🚀 **Go 泛型**：充分利用 Go 泛型特性
- 📦 **模块化设计**：按需引入，不臃肿
- 🔒 **并发安全**：涉及共享状态的工具保证 goroutine 安全，CI 全程开启 `-race` 检测
- ✅ **测试完备**：所有模块均有完整单元测试，golangci-lint 严格把关
- 🎯 **Go 风格 API**：遵循 Go 命名规范，不照搬 Java

## 安装

```bash
go get github.com/camark/GoHutools
```

## 模块列表

### Phase 1: 基础工具

| 模块 | 说明 | 对标 Hutool |
|------|------|------------|
| `strutil` | 字符串工具 | StrUtil |
| `numutil` | 数字工具 | NumberUtil |
| `collutil` | 集合工具（泛型） | CollUtil |
| `maputil` | Map 工具 | MapUtil |
| `arrayutil` | 数组/Slice 工具 | ArrayUtil |
| `dateutil` | 日期时间工具 | DateUtil |
| `convert` | 类型转换 | Convert |
| `validate` | 数据校验 | Validator |

### Phase 2: IO 与文件

| 模块 | 说明 | 对标 Hutool |
|------|------|------------|
| `ioutil` | IO 工具 | IoUtil |
| `fileutil` | 文件工具 | FileUtil |
| `charsetutil` | 字符集工具 | CharsetUtil |

### Phase 3: 网络与数据

| 模块 | 说明 | 对标 Hutool |
|------|------|------------|
| `httpclient` | HTTP 客户端 | HttpUtil |
| `htmlutil` | HTML 工具 (转义/过滤/标签清理) | HtmlUtil |
| `jsonutil` | JSON 工具 | JSONUtil |

### Phase 4: 安全与编码

| 模块 | 说明 | 对标 Hutool |
|------|------|------------|
| `crypto` | 加密解密 (MD5/SHA/AES/RSA/HMAC) | SecureUtil |
| `codec` | 编解码 (Base64/Hex/URL) | CodecUtil |
| `idutil` | ID 生成 (UUID/Snowflake/ULID) | IdUtil |
| `randomutil` | 随机工具 | RandomUtil |

### Phase 5: 高级功能

| 模块 | 说明 | 对标 Hutool |
|------|------|------------|
| `cache` | 缓存 (LRU/FIFO/LFU/Timed) | hutool-cache |
| `cron` | 定时任务调度 | hutool-cron |
| `captcha` | 验证码生成 | hutool-captcha |
| `bloom` | 布隆过滤器 | hutool-bloomFilter |
| `pool` | 对象池 | Pool |
| `csv` | CSV 读写封装 | hutool-csv |
| `csv` | CSV 读写封装 | hutool-csv |

### Phase 6: 系统与配置

| 模块 | 说明 | 对标 Hutool |
|------|------|------------|
| `log` | 日志门面 | hutool-log |
| `setting` | 配置文件 | hutool-setting |
| `system` | 系统信息 | hutool-system |
| `objectutil` | 对象工具 | ObjectUtil |
| `regexutil` | 正则工具 | ReUtil |

### Phase 7: 断言、Bean 与并发

| 模块 | 说明 | 对标 Hutool |
|------|------|------------|
| `assert` | 断言校验 | Assert |
| `beanutil` | Bean 拷贝 / Map 互转 | BeanUtil |
| `netutil` | 网络工具 (内网 IP/本机地址) | NetUtil |
| `threadutil` | 线程 / goroutine 池 | ThreadUtil |
| `ziputil` | ZIP/GZIP/TAR 压缩解压 | hutool-zip |
| `jwt` | JWT 签名验签 (HS256/384/512) | JWTUtil |

### Phase 8: 核心语言扩展

| 模块 | 说明 | 对标 Hutool |
|------|------|------------|
| `tuple` | Pair/三元组/元组容器 | Pair / Tuple |
| `opt` | Optional 模式 (Some/None) | Opt |
| `urlutil` | URL 规范化/构建/参数 | URLUtil / UrlBuilder |

### Phase 9: 基础工具补充

| 模块 | 说明 | 对标 Hutool |
|------|------|------------|
| `stopwatch` | 任务计时器 (Start/Stop/Split) | StopWatch |
| `hashutil` | 经典字符串散列 (FNV/BKDR/DJB/...) | HashUtil |
| `boolutil` | 布尔工具 (parse/and/or/xor) | BooleanUtil |
| `charutil` | 字符工具 (isNumber/isEmoji/...) | CharUtil |
| `reflectutil` | 反射工具 (字段读写/方法调用/tag) | ReflectUtil |
| `reflectutil` | 反射工具 (字段读写/方法调用/tag) | ReflectUtil |
| `classutil` | 类型工具 (类型名/分类/方法集/可赋值) | ClassUtil |
| `enumutil` | 枚举工具 (字符串映射/校验) | EnumUtil |

## 使用示例

### 字符串工具 (strutil)

```go
import "github.com/camark/GoHutools/strutil"

strutil.IsBlank("")      // true
strutil.CamelToUnderline("helloWorld")  // "hello_world"
strutil.UnderlineToCamel("hello_world") // "helloWorld"
strutil.Reverse("hello")                // "olleh"
strutil.Format("Hello {0}!", "World")   // "Hello World!"
```

### 集合工具 (collutil)

```go
import "github.com/camark/GoHutools/collutil"

nums := []int{1, 2, 3, 4, 5}

collutil.Filter(nums, func(n int) bool { return n > 3 })  // [4, 5]
collutil.Map(nums, func(n int) int { return n * 2 })       // [2, 4, 6, 8, 10]
collutil.Distinct([]int{1, 2, 2, 3, 3, 3})                // [1, 2, 3]
collutil.GroupBy(nums, func(n int) string {
    if n%2 == 0 { return "even" }
    return "odd"
})
collutil.Sum(nums)  // 15
```

### HTTP 客户端 (httpclient)

```go
import "github.com/camark/GoHutools/httpclient"

// 快捷方法
resp, err := httpclient.Get("https://api.example.com/users")
body, _ := resp.BodyString()

// 链式调用
client := httpclient.New()
resp, err := client.Get("https://api.example.com/users").
    Header("Authorization", "Bearer token").
    Query("page", "1").
    Timeout(5 * time.Second).
    Do()

// POST JSON
resp, err := httpclient.PostJSON("https://api.example.com/users", map[string]string{
    "name": "John",
    "email": "john@example.com",
})
```

### 加密工具 (crypto)

```go
import "github.com/camark/GoHutools/crypto"

// MD5
hash := crypto.MD5String("hello")  // "5d41402abc4b2a76b9719d911017c592"

// SHA256
hash := crypto.SHA256Hex([]byte("hello"))

// AES 加密
key := []byte("0123456789abcdef")
encrypted, _ := crypto.AESEncrypt(key, []byte("secret"))
decrypted, _ := crypto.AESDecrypt(key, encrypted)

// RSA
priv, pub, _ := crypto.GenerateRSAKeyPair(2048)
encrypted, _ := crypto.RSAEncrypt(pub, []byte("secret"))
decrypted, _ := crypto.RSADecrypt(priv, encrypted)
```

### 缓存 (cache)

```go
import "github.com/camark/GoHutools/cache"

// LRU 缓存
c := cache.NewLRU(100)
c.Set("key", "value")
v, ok := c.Get("key")

// 带过期时间的缓存
c := cache.NewTimed(5 * time.Minute)
c.SetWithExpire("key", "value", 1 * time.Minute)
```

### ID 生成 (idutil)

```go
import "github.com/camark/GoHutools/idutil"

uuid := idutil.UUID()           // "550e8400-e29b-41d4-a716-446655440000"
simple := idutil.SimpleUUID()   // "550e8400e29b41d4a716446655440000"
ulid := idutil.ULID()           // "01ARZ3NDEKTSV4RRFFQ69G5FAV"
nano := idutil.NanoID()         // "V1StGXR8_Z5jdHi6B-myT"

// Snowflake ID
sf, _ := idutil.NewSnowflake(1)
id, _ := sf.NextID()
```

### 定时任务 (cron)

```go
import "github.com/camark/GoHutools/cron"

s := cron.New()
s.AddFunc("*/5 * * * * *", func() {
    fmt.Println("每5秒执行一次")
})
s.Start()
defer s.Stop()
```

### 布隆过滤器 (bloom)

```go
import "github.com/camark/GoHutools/bloom"

// 预期 1000 个元素，误判率 1%
f := bloom.New(1000, 0.01)
f.AddString("hello")
f.AddString("world")

f.ContainsString("hello")  // true
f.ContainsString("foo")    // false (可能误判)
```

### 数据校验 (validate)

```go
import "github.com/camark/GoHutools/validate"

validate.IsEmail("test@example.com")    // true
validate.IsMobile("13812345678")        // true
validate.IsIDCard("11010519491231002X") // true
validate.IsURL("https://example.com")   // true
validate.IsIPv4("192.168.1.1")          // true
```

## 统计

- **45 个功能模块**
- **1000+ 个导出函数与方法**
- **全面的单元测试覆盖**（CI 含 `-race` 竞态检测与覆盖率统计）

## 开发

```bash
# 运行全部测试
go test ./...

# 带竞态检测
go test -race ./...

# 静态检查（需安装 golangci-lint v2）
golangci-lint run ./...
```

## 许可证

[Apache License 2.0](LICENSE)

## 致谢

- [Hutool](https://github.com/dromara/hutool) - Java 版原作
