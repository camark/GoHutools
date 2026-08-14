---
home: true
title: GoHutool
heroImage: /logo.svg
heroText: GoHutool
tagline: Go 版本的 Hutool 工具库
actions:
  - text: 快速开始
    link: /guide/getting-started
    type: primary
  - text: API 文档
    link: /api/
    type: secondary
features:
  - title: 🚀 Go 泛型
    details: 充分利用 Go 1.21+ 泛型特性，类型安全
  - title: 📦 模块化设计
    details: 27 个独立模块，按需引入，不臃肿
  - title: 🔒 并发安全
    details: 涉及共享状态的工具保证 goroutine 安全
  - title: ✅ 测试完备
    details: 900+ 工具函数，全面的单元测试覆盖
  - title: 🎯 Go 风格 API
    details: 遵循 Go 命名规范，不照搬 Java
  - title: 📚 丰富功能
    details: 字符串、集合、加密、HTTP、缓存、定时任务等
footer: Apache License 2.0
---

## 快速体验

```go
package main

import (
    "fmt"
    "github.com/gongm/gohutool/strutil"
    "github.com/gongm/gohutool/collutil"
)

func main() {
    // 字符串工具
    fmt.Println(strutil.CamelToUnderline("helloWorld")) // "hello_world"

    // 集合工具
    nums := []int{1, 2, 3, 4, 5}
    even := collutil.Filter(nums, func(n int) bool { return n%2 == 0 })
    fmt.Println(even) // [2, 4]
}
```
