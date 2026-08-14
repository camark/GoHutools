---
title: 介绍
---

# GoHutool 介绍

GoHutool 是 Go 版本的 [Hutool](https://github.com/dromara/hutool) 工具库，功能对标 Java 版 Hutool。

## 特性

- **Go 泛型**：充分利用 Go 1.21+ 泛型特性
- **模块化设计**：27 个独立模块，按需引入
- **并发安全**：涉及共享状态的工具保证 goroutine 安全
- **测试完备**：900+ 工具函数，全面的单元测试覆盖
- **Go 风格 API**：遵循 Go 命名规范

## 模块概览

| 类别 | 模块 |
|------|------|
| 基础工具 | strutil, numutil, collutil, maputil, arrayutil, dateutil, convert, validate |
| IO 与文件 | ioutil, fileutil, charsetutil |
| 网络与数据 | httpclient, jsonutil |
| 安全与编码 | crypto, codec, idutil, randomutil |
| 高级功能 | cache, cron, captcha, bloom, pool |
| 系统与配置 | log, setting, system, objectutil, regexutil |
