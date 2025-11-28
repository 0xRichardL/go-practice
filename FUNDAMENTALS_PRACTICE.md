# 🏁 Go Skills Recovery --- 4-Day Plan

A structured, high-impact checklist to rapidly regain Go fundamentals
after time away.

## 📅 Day 1 --- Core Language Refresh

### Syntax & Basics

- [x] Variables, constants, short declarations (`:=`)
- [ ] Primitive data types
- [ ] `iota`
- [ ] All `for` loop variants
- [ ] `if`, `switch`, type switch
- [ ] `defer`, `panic`, `recover`

### Functions

- [ ] Named return values
- [ ] Variadic functions
- [ ] Anonymous functions
- [ ] Higher-order functions

### Error Handling

- [ ] Custom error types
- [ ] Wrapping errors (`fmt.Errorf("%w")`)
- [ ] Checking errors (`errors.Is`, `errors.As`)

### Mini Tasks

- [ ] Build a simple CLI
- [ ] Write small utility functions (math/string)
- [ ] Experiment with `defer/panic/recover`

## 📅 Day 2 --- Concurrency & Goroutines

### Goroutines

- [ ] Understand G--M--P scheduler
- [ ] Launch goroutines (`go func() {}`)
- [ ] Identify/fix goroutine leaks

### Channels

- [ ] Unbuffered channels
- [ ] Buffered channels
- [ ] Directional channels (`<-chan`, `chan<-`)
- [ ] Closing channels correctly
- [ ] Timeouts with `select` + `time.After`

### Concurrency Patterns

- [ ] Worker pool
- [ ] Fan-in / Fan-out
- [ ] Context cancellation
- [ ] Rate limiting (ticker)

### Sync Primitives

- [ ] RWMutex, Mutex
- [ ] WaitGroup
- [ ] Once

### Mini Tasks

- [ ] Build worker pool
- [ ] API caller with rate limit
- [ ] Write a goroutine leak + fix it

## 📅 Day 3 --- DS, Interfaces, and Stdlib Mastery

### Interfaces & Types

- [ ] Implicit interface implementation
- [ ] Empty interface and when to avoid
- [ ] Custom types + method sets
- [ ] Type assertions & switches

### Collections

- [ ] Slice mechanics (cap growth, append rules)
- [ ] Maps (zero value, delete, exists check)
- [ ] Structs & JSON tags

### Essential Stdlib

- [ ] `context`
- [ ] `net/http`
- [ ] `encoding/json`
- [ ] `io` & `bufio`
- [ ] `os`
- [ ] `time`
- [ ] `errors`

### Mini Tasks

- [ ] Build simple REST API (no frameworks)
- [ ] Write your own `io.Reader` implementation
- [ ] Benchmark slice vs map

## 📅 Day 4 --- Real-World Backend Skills

### Project Structure

- [ ] `cmd/`, `internal/`, `pkg/`
- [ ] Config loading (env, JSON, YAML)

### Web & Middleware

- [ ] Gin or Echo basics
- [ ] Request binding & validation
- [ ] Swagger UI
- [ ] Custom middleware (auth/logging/recovery)

### Database & Persistence

- [ ] SQLC or GORM
- [ ] Transactions
- [ ] Connection pooling

### Testing & Tooling

- [ ] Table-driven tests
- [ ] Mocks (GoMock or custom doubles)
- [ ] Benchmarks
- [ ] `golangci-lint`
- [ ] `go mod tidy`, replace directive

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
