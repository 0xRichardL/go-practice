# 🏁 Go Skills Recovery --- 4-Day Plan

A structured, high-impact checklist to rapidly regain Go fundamentals
after time away.

## 📅 Day 1 --- Core Language Refresh

### Syntax & Basics

- [x] Variables, constants, short declarations (`:=`)
- [x] Primitive data types
- [x] `iota`
- [x] All `for` loop variants
- [x] `if`, `switch`, type switch
- [x] `defer`, `panic`, `recover`

### Functions

- [x] Named return values
- [x] Variadic functions
- [x] Anonymous functions
- [x] Higher-order functions

### Error Handling

- [x] Custom error types
- [x] Wrapping errors (`fmt.Errorf("%w")`)
- [x] Checking errors (`errors.Is`, `errors.As`)

### Mini Tasks

- [ ] Build a simple CLI
- [ ] Write small utility functions (math/string)
- [ ] Experiment with `defer/panic/recover`

## 📅 Day 2 --- Concurrency & Goroutines

### Goroutines

- [x] Understand G--M--P scheduler
- [x] Launch goroutines (`go func() {}`)
- [x] Identify/fix goroutine leaks

### Channels

- [x] Unbuffered channels
- [x] Buffered channels
- [x] Directional channels (`<-chan`, `chan<-`)
- [x] Closing channels correctly
- [x] Timeouts with `select` + `time.After`

### Concurrency Patterns

- [x] Worker pool
- [x] Fan-in / Fan-out
- [ ] Context cancellation
- [x] Rate limiting (ticker)

### Sync Primitives

- [ ] RWMutex, Mutex
- [x] WaitGroup
- [x] Once

### Mini Tasks

- [x] Build worker pool
- [x] API caller with rate limit
- [x] Write a goroutine leak + fix it

## 📅 Day 3 --- DS, Interfaces, and Stdlib Mastery

### Interfaces & Types

- [x] Implicit interface implementation
- [ ] Empty interface and when to avoid
- [x] Custom types + method sets
- [x] Type assertions & switches

### Collections

- [x] Slice mechanics (cap growth, append rules)
- [x] Maps (zero value, delete, exists check)
- [x] Structs & JSON tags

### Essential Stdlib

- [ ] `context`
- [ ] `net/http`
- [ ] `encoding/json`
- [ ] `io` & `bufio`
- [x] `os`
- [x] `time`
- [x] `errors`

### Mini Tasks

- [ ] Build simple REST API (no frameworks)
- [ ] Write your own `io.Reader` implementation
- [ ] Benchmark slice vs map

## 📅 Day 4 --- Real-World Backend Skills

### Project Structure

- [x] `cmd/`, `internal/`, `pkg/`
- [x] Config loading (env, JSON, YAML)

### Web & Middleware

- [ ] Gin or Echo basics
- [x] Request binding & validation
- [x] Swagger UI
- [ ] Custom middleware (auth/logging/recovery)

### Database & Persistence

- [ ] SQLC or GORM
- [ ] Transactions
- [ ] Connection pooling

### Testing & Tooling

- [ ] Table-driven tests
- [ ] Mocks (GoMock or custom doubles)
- [ ] Benchmarks
- [x] `golangci-lint`
- [x] `go mod tidy`, replace directive

### Mini Project --- Build a Microservice

- [ ] REST API
- [ ] Database adapter
- [ ] Background worker
- [ ] Swagger docs
- [ ] Dockerfile + docker-compose
- [ ] Unit + integration tests

## ⭐ Optional Extensions

- [ ] Learn Generics deeply
- [ ] Build your own LRU cache
- [ ] Write a mini job scheduler
- [ ] Build in-memory message queue
- [ ] Learn more advanced concurrency patterns
