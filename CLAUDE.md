# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

GoHutool is a Go port of Java's [Hutool](https://github.com/dromara/hutool) utility library. It provides 27 packages with 900+ utility functions covering strings, collections, crypto, HTTP, caching, and more. Uses Go 1.21+ generics extensively.

## Build & Test Commands

```bash
# Run all tests
go test ./...

# Run tests without cache
go test -count=1 ./...

# Run tests for a single package
go test ./strutil/...
go test ./collutil/...

# Run a specific test
go test -run TestIsBlank ./strutil/
go test -v -run TestFilter ./collutil/

# Vet all packages
go vet ./...

# Build check (compile only)
go build ./...

# Update dependencies
go mod tidy
```

## Architecture

- **Module**: `github.com/camark/GoHutools` (single Go module, no sub-modules)
- **Dependencies**: `golang.org/x/text` (charsetutil), `golang.org/x/net` (httpclient)
- **Package pattern**: Each package is a flat directory with `*.go` and `*_test.go` files

### Package Naming Convention

Packages follow the pattern `<domain>util` (e.g., `strutil`, `numutil`, `collutil`). Exception: `crypto`, `cache`, `codec`, `log`, `cache`, `bloom`, `pool`, `cron`, `captcha`, `setting`, `system`.

### Key Design Patterns

1. **Generics for collections**: `collutil`, `maputil`, `arrayutil` use Go generics (`[T any]`, `[T comparable]`)
2. **Interface-based polymorphism**: `cache.Cache` interface with LRU/FIFO/LFU/Timed implementations; `log.Logger` interface
3. **Fluent builder pattern**: `httpclient.Request` uses method chaining (each method returns `*Request`)
4. **Thread safety**: `cache`, `bloom`, `cron`, `idutil.SnowflakeID`, `idutil.Sequence` use `sync.Mutex`/`sync.RWMutex`
5. **Package-level convenience functions**: `httpclient.Get()`, `httpclient.Post()` wrap `httpclient.New()`

### Code Conventions

- All exported functions have Go doc comments starting with function name
- Test files use table-driven tests (`tests := []struct{...}`)
- Error handling returns `(value, error)` tuples
- No init() functions except in `log` package for default logger
