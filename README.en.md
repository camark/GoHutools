# GoHutool

[中文](README.md) | **English**

A Go port of the [Hutool](https://github.com/dromara/hutool) utility library, matching the Java edition feature by feature and built with Go generics.

[![CI](https://img.shields.io/github/actions/workflow/status/camark/GoHutools/ci.yml?branch=main&logo=github)](https://github.com/camark/GoHutools/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8.svg?logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)

## Features

- 🚀 **Go generics**: built on Go's type parameters from the ground up
- 📦 **Modular design**: import only what you need, no bloat
- 🔒 **Concurrency-safe**: tools with shared state are goroutine-safe; CI runs the whole test suite with `-race`
- ✅ **Well tested**: every module ships with unit tests and is gated by golangci-lint
- 🎯 **Idiomatic Go API**: follows Go conventions instead of blindly mirroring Java

## Installation

```bash
go get github.com/camark/GoHutools
```

## Modules

### Phase 1: Core utilities

| Module | Description | Hutool counterpart |
|--------|-------------|--------------------|
| `strutil` | String utilities | StrUtil |
| `numutil` | Number utilities | NumberUtil |
| `collutil` | Collection utilities (generic) | CollUtil |
| `maputil` | Map utilities | MapUtil |
| `arrayutil` | Array/Slice utilities | ArrayUtil |
| `dateutil` | Date & time utilities | DateUtil |
| `convert` | Type conversion | Convert |
| `validate` | Data validation | Validator |

### Phase 2: IO & files

| Module | Description | Hutool counterpart |
|--------|-------------|--------------------|
| `ioutil` | IO utilities | IoUtil |
| `fileutil` | File utilities | FileUtil |
| `charsetutil` | Charset utilities | CharsetUtil |

### Phase 3: Network & data

| Module | Description | Hutool counterpart |
|--------|-------------|--------------------|
| `httpclient` | HTTP client | HttpUtil |
| `jsonutil` | JSON utilities | JSONUtil |

### Phase 4: Security & encoding

| Module | Description | Hutool counterpart |
|--------|-------------|--------------------|
| `crypto` | Encryption (MD5/SHA/AES/RSA/HMAC) | SecureUtil |
| `codec` | Encoding/decoding (Base64/Hex/URL) | CodecUtil |
| `idutil` | ID generation (UUID/Snowflake/ULID) | IdUtil |
| `randomutil` | Random utilities | RandomUtil |

### Phase 5: Advanced features

| Module | Description | Hutool counterpart |
|--------|-------------|--------------------|
| `cache` | Caching (LRU/FIFO/LFU/Timed) | hutool-cache |
| `cron` | Cron-style task scheduling | hutool-cron |
| `captcha` | Captcha generation | hutool-captcha |
| `bloom` | Bloom filter | hutool-bloomFilter |
| `pool` | Object pooling | Pool |

### Phase 6: System & configuration

| Module | Description | Hutool counterpart |
|--------|-------------|--------------------|
| `log` | Logging facade | hutool-log |
| `setting` | Configuration files | hutool-setting |
| `system` | System information | hutool-system |
| `objectutil` | Object utilities | ObjectUtil |
| `regexutil` | Regex utilities | ReUtil |

### Phase 7: Assertion, Bean & concurrency

| Module | Description | Hutool counterpart |
|--------|-------------|--------------------|
| `assert` | Assertion helpers | Assert |
| `beanutil` | Bean copy / Map <-> Bean | BeanUtil |
| `netutil` | Network utilities (internal IP / local addr) | NetUtil |
| `threadutil` | Thread / goroutine pool | ThreadUtil |
| `ziputil` | ZIP/GZIP/TAR compress & extract | hutool-zip |
| `jwt` | JWT sign & verify (HS256/384/512) | JWTUtil |

## Examples

### Strings (strutil)

```go
import "github.com/camark/GoHutools/strutil"

strutil.IsBlank("")      // true
strutil.CamelToUnderline("helloWorld")  // "hello_world"
strutil.UnderlineToCamel("hello_world") // "helloWorld"
strutil.Reverse("hello")                // "olleh"
strutil.Format("Hello {0}!", "World")   // "Hello World!"
```

### Collections (collutil)

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

### HTTP client (httpclient)

```go
import "github.com/camark/GoHutools/httpclient"

// Convenience functions
resp, err := httpclient.Get("https://api.example.com/users")
body, _ := resp.BodyString()

// Fluent API
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

### Crypto (crypto)

```go
import "github.com/camark/GoHutools/crypto"

// MD5
hash := crypto.MD5String("hello")  // "5d41402abc4b2a76b9719d911017c592"

// SHA256
hash := crypto.SHA256Hex([]byte("hello"))

// AES
key := []byte("0123456789abcdef")
encrypted, _ := crypto.AESEncrypt(key, []byte("secret"))
decrypted, _ := crypto.AESDecrypt(key, encrypted)

// RSA
priv, pub, _ := crypto.GenerateRSAKeyPair(2048)
encrypted, _ := crypto.RSAEncrypt(pub, []byte("secret"))
decrypted, _ := crypto.RSADecrypt(priv, encrypted)
```

### Cache (cache)

```go
import "github.com/camark/GoHutools/cache"

// LRU cache
c := cache.NewLRU(100)
c.Set("key", "value")
v, ok := c.Get("key")

// Cache with expiration
c := cache.NewTimed(5 * time.Minute)
c.SetWithExpire("key", "value", 1 * time.Minute)
```

### ID generation (idutil)

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

### Cron scheduling (cron)

```go
import "github.com/camark/GoHutools/cron"

s := cron.New()
s.AddFunc("*/5 * * * * *", func() {
    fmt.Println("runs every 5 seconds")
})
s.Start()
defer s.Stop()
```

### Bloom filter (bloom)

```go
import "github.com/camark/GoHutools/bloom"

// Expect 1000 elements with a 1% false-positive rate
f := bloom.New(1000, 0.01)
f.AddString("hello")
f.AddString("world")

f.ContainsString("hello")  // true
f.ContainsString("foo")    // false (may false-positive)
```

### Validation (validate)

```go
import "github.com/camark/GoHutools/validate"

validate.IsEmail("test@example.com")    // true
validate.IsMobile("13812345678")        // true
validate.IsIDCard("11010519491231002X") // true
validate.IsURL("https://example.com")   // true
validate.IsIPv4("192.168.1.1")          // true
```

## Stats

- **33 modules**
- **950+ exported functions and methods**
- **Full unit-test coverage** (CI runs with the `-race` detector and coverage reporting)

## Development

```bash
# Run all tests
go test ./...

# With the race detector
go test -race ./...

# Lint (requires golangci-lint v2)
golangci-lint run ./...
```

## License

[Apache License 2.0](LICENSE)

## Acknowledgements

- [Hutool](https://github.com/dromara/hutool) — the original Java library
