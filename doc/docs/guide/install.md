---
title: 安装
---

# 安装

## 环境要求

- Go 1.21+（需要泛型支持）

## 安装命令

```bash
go get github.com/camark/GoHutools
```

## 按需引入

每个模块都是独立的包，可以按需引入：

```go
import "github.com/camark/GoHutools/strutil"
import "github.com/camark/GoHutools/collutil"
import "github.com/camark/GoHutools/crypto"
```

## 依赖说明

核心模块无第三方依赖，以下模块有额外依赖：

| 模块 | 依赖 |
|------|------|
| charsetutil | golang.org/x/text |
| httpclient | golang.org/x/net |

## 验证安装

```go
package main

import (
    "fmt"
    "github.com/camark/GoHutools/strutil"
)

func main() {
    fmt.Println(strutil.Version())
}
```
