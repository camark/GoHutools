---
title: 快速开始
---

# 快速开始

## 安装

```bash
go get github.com/camark/GoHutools
```

## 使用示例

### 字符串处理

```go
package main

import (
    "fmt"
    "github.com/camark/GoHutools/strutil"
)

func main() {
    fmt.Println(strutil.IsBlank(""))      // true
    fmt.Println(strutil.Reverse("hello")) // "olleh"
    fmt.Println(strutil.CamelToUnderline("helloWorld")) // "hello_world"
}
```

### 集合操作

```go
package main

import (
    "fmt"
    "github.com/camark/GoHutools/collutil"
)

func main() {
    nums := []int{1, 2, 3, 4, 5}

    // 过滤
    even := collutil.Filter(nums, func(n int) bool { return n%2 == 0 })
    fmt.Println(even) // [2, 4]

    // 映射
    doubled := collutil.Map(nums, func(n int) int { return n * 2 })
    fmt.Println(doubled) // [2, 4, 6, 8, 10]

    // 去重
    unique := collutil.Distinct([]int{1, 2, 2, 3, 3, 3})
    fmt.Println(unique) // [1, 2, 3]
}
```

### HTTP 请求

```go
package main

import (
    "fmt"
    "github.com/camark/GoHutools/httpclient"
)

func main() {
    // GET 请求
    resp, err := httpclient.Get("https://api.example.com/users")
    if err != nil {
        panic(err)
    }
    body, _ := resp.BodyString()
    fmt.Println(body)

    // POST JSON
    resp, err = httpclient.PostJSON("https://api.example.com/users", map[string]string{
        "name": "John",
    })
}
```

### 加密解密

```go
package main

import (
    "fmt"
    "github.com/camark/GoHutools/crypto"
)

func main() {
    // MD5
    hash := crypto.MD5String("hello")
    fmt.Println(hash) // "5d41402abc4b2a76b9719d911017c592"

    // AES 加密
    key := []byte("0123456789abcdef")
    encrypted, _ := crypto.AESEncrypt(key, []byte("secret"))
    decrypted, _ := crypto.AESDecrypt(key, encrypted)
    fmt.Println(string(decrypted)) // "secret"
}
```
