# � Go Fundamentals Mastery Checklist

A comprehensive, level-based checklist to master Go from basics to expert patterns.

---

## 🌱 Level 1: Beginner

### Core Syntax & Primitives

- [x] Variables, constants, short declarations (`:=`)
- [x] All primitive data types (`int`, `uint`, `float`, `bool`, `string`)
- [x] Type conversions
- [x] Zero values
- [x] `iota` for constants

### Pointers & Memory

- [x] Pointer basics (`&`, `*`)
- [x] Pointer vs value semantics
- [x] Nil pointers
- [x] Pointer arithmetic (not allowed in Go)

### Structs

- [x] Struct declaration and initialization
- [x] Field access and modification
- [x] Anonymous structs
- [x] Struct literals
- [x] Comparing structs

### Control Flow

- [x] `for` loops (classic, while-style, range)
- [x] `if`, `else if`, `else`
- [x] `switch` statements
- [x] `break`, `continue`

### Functions Basics

- [x] Function declarations
- [x] Multiple return values
- [x] Named return values
- [x] Variadic functions

### Basic Error Handling

- [x] Error type and `error` interface
- [x] Creating errors with `errors.New`
- [x] Error checking patterns

### Collections Fundamentals

- [x] Array basics
- [x] Slice creation and basic operations
- [x] Map creation and basic CRUD
- [x] Range iteration over slices and maps

### Formatting & Output

- [x] `fmt.Print`, `Println`, `Printf`
- [x] `fmt.Sprintf` for string formatting
- [x] Format verbs (`%v`, `%+v`, `%#v`, `%T`, `%d`, `%s`, `%f`)
- [x] Width and precision in formatting

### Command-Line Basics

- [x] `os.Args` for command-line arguments
- [ ] `flag` package for parsing flags
- [x] Exit codes with `os.Exit()`

### Mini Projects

- [ ] CLI calculator
- [ ] File reader/writer
- [ ] String manipulation utilities

---

## 🚀 Level 2: Intermediate

### Advanced Functions

- [x] Anonymous functions
- [x] Closures
- [x] Higher-order functions
- [x] `defer` statements
- [x] `panic` and `recover`

### Slice Deep Dive

- [x] Length vs capacity
- [x] Append behavior and capacity growth
- [x] Slice bounds
- [x] Memory sharing and sub-slicing
- [x] Full slice expressions `[low:high:max]`
- [x] Nil vs empty slices
- [x] Copying slices
- [x] Function modification behavior

### Map Deep Dive

- [x] Map zero values
- [x] Checking key existence
- [x] Deleting keys
- [x] Map iteration order (non-deterministic)
- [ ] Maps as sets pattern

### Strings & Text

- [x] String immutability
- [x] String vs `[]byte` vs `[]rune`
- [x] `strings` package
- [x] `strconv` package (Atoi, ParseInt, FormatInt, ParseFloat)
  - [x] String conversions and parsing
- [x] Regular expressions (`regexp` package)
  - [x] Compile vs MustCompile
  - [x] Match, Find, Replace patterns

### Error Handling Advanced

- [x] Custom error types
- [x] Error wrapping with `fmt.Errorf("%w")`
- [x] Error unwrapping (`errors.Unwrap`, `errors.Is`, `errors.As`)

### Types & Interfaces

- [x] Custom types
- [x] Type definitions vs type aliases
- [x] Method sets
- [x] Methods on custom types (value vs pointer receivers)
- [x] When to use pointer vs value receivers
- [x] Method set implications for interfaces
- [x] Interface basics
- [x] Implicit interface satisfaction
- [x] Type assertions (safe & unsafe)
- [x] Type switches
- [x] Empty interface (`interface{}` / `any`)
- [x] Type assertion on empty interface

### Struct Composition

- [ ] Struct embedding
- [ ] Method promotion
- [ ] Anonymous fields
- [ ] Overriding embedded methods
- [ ] Interface embedding

### Basic I/O

- [x] `fmt` package (Print, Printf, Println)
- [x] `os` package basics
- [x] File operations
- [x] `io.Reader` and `io.Writer` interfaces
- [x] `io.Copy`, `io.ReadAll`, `io.ReadFull`
- [x] Implementing custom Reader/Writer
- [x] `bufio` package (Scanner, Reader, Writer)
- [x] Buffered I/O performance benefits
- [x] Reading line by line with Scanner

### JSON & Serialization

- [x] Struct tags for JSON (`json:"field_name"`)
- [x] `json.Marshal` and `json.Unmarshal`
- [x] Handling nested structs
- [x] Omitempty and string options
- [x] Custom JSON marshaling (`MarshalJSON`, `UnmarshalJSON`)

### Mini Projects

- [ ] Custom data structure (stack, queue)
- [ ] Text parser
- [ ] CSV reader/writer
- [ ] Log file analyzer

---

## 💪 Level 3: Proficient

### Goroutines & Concurrency

- [x] Understanding goroutine basics
- [x] G-M-P scheduler model
- [x] Launching goroutines
- [x] Goroutine leaks and how to avoid them
- [x] `WaitGroup` for goroutine synchronization

### Channels

- [x] Unbuffered channels
- [x] Buffered channels
- [x] Directional channels (`<-chan`, `chan<-`)
- [x] Channel closing rules
- [x] Range over channels
- [x] `select` statement
- [x] Timeouts with `time.After`
- [x] Deadlock scenarios and prevention

### Sync Primitives

- [x] `Mutex` and critical sections
- [ ] `RWMutex` (read/write locks)
- [x] `sync.Once`
- [x] `sync.Map` for concurrent map access
- [ ] `sync.Pool` for object reuse

### Concurrency Patterns

- [x] Worker pool pattern
- [x] Fan-in pattern (merging channels)
- [x] Fan-out pattern (distributing work)
- [x] Rate limiting / Throttling
- [ ] Pipeline pattern (chaining stages with channels)
- [ ] Generator pattern
- [ ] Semaphore pattern

### Context Package

- [ ] `context.Background()` and `context.TODO()`
- [ ] `context.WithCancel()` for cancellation
- [ ] `context.WithTimeout()` for time-based cancellation
- [ ] `context.WithDeadline()` for absolute deadlines
- [ ] `context.WithValue()` for request-scoped data
- [ ] Context propagation patterns
- [ ] Checking `ctx.Done()` in goroutines
- [ ] Handling context cancellation errors

### Standard Library Mastery

- [x] `time` package (Duration, Timer, Ticker)
- [ ] `time.After`, `time.AfterFunc`, time zones
- [x] `sort` package
- [ ] Custom sort with `sort.Slice` and `sort.Interface`
- [ ] `encoding/json` (Marshal, Unmarshal, tags)
- [ ] JSON streaming with Encoder/Decoder
- [ ] `encoding/xml` basics
- [ ] `io` and `bufio` advanced usage
- [ ] `bytes.Buffer` for efficient string building
- [ ] `bytes.Reader` and manipulation
- [ ] `strings.Builder` for string concatenation

### Testing

- [x] Basic unit tests (`*_test.go`)
- [ ] Table-driven tests (slice of test cases)
- [ ] Test helpers and utilities
- [ ] Subtests with `t.Run()` for organization
- [ ] Test fixtures and setup/teardown
- [ ] `testing.T` methods (`Error`, `Fatal`, `Skip`)
- [ ] Benchmarks with `Benchmark*` functions
- [ ] `b.N`, `b.ResetTimer()`, `b.StopTimer()`
- [ ] Test coverage with `-cover` flag
- [ ] Generating coverage reports
- [ ] `testing.Short()` for quick test mode

### Data Structures & Algorithms

- [x] Custom sorted set implementation
- [ ] LRU cache
- [ ] Trie
- [ ] Binary search implementations

### Mini Projects

- [ ] Concurrent web scraper
- [ ] Rate-limited API client
- [ ] In-memory cache with expiration
- [ ] Background job processor

---

## 🎓 Level 4: Expert

### Advanced Concurrency

- [ ] Advanced context patterns (cancellation trees)
- [ ] Graceful shutdown patterns (signal handling)
- [ ] Orchestrating multiple goroutines with errgroup
- [ ] Channel composition patterns
- [ ] Memory model understanding (happens-before)
- [ ] Race condition detection (`-race` flag)
- [ ] Lock-free programming patterns
- [ ] `sync.Cond` for condition variables
- [ ] Atomic operations (`sync/atomic`)

### HTTP & Web Services

- [ ] `net/http` standard library basics
- [ ] `http.Handler` and `http.HandlerFunc`
- [ ] `http.ServeMux` for routing
- [ ] HTTP client with timeouts and retries
- [ ] Custom `http.Transport` and connection pooling
- [ ] HTTP server and handler patterns
- [ ] Middleware pattern (wrapping handlers)
- [ ] Middleware chaining
- [ ] Context propagation in HTTP handlers
- [ ] Request and response handling
- [ ] RESTful API design principles
- [ ] Content negotiation (JSON, XML)
- [ ] HTTP status codes and error responses
- [ ] Framework basics (Gin, Echo, or Fiber)
- [ ] Router parameters and query strings

### Database Integration

- [ ] `database/sql` interface and driver registration
- [ ] Opening connections and error handling
- [ ] Query, QueryRow, Exec methods
- [ ] Prepared statements and SQL injection prevention
- [ ] Connection pooling configuration
- [ ] `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`
- [ ] Transactions (`Begin`, `Commit`, `Rollback`)
- [ ] Transaction isolation levels
- [ ] Context-aware database operations
- [ ] Scanning results into structs
- [ ] Handling NULL values (sql.NullString, sql.NullInt64)
- [ ] ORM usage (GORM) or query builders (SQLC, sqlx)
- [ ] Migration strategies (golang-migrate, GORM AutoMigrate)
- [ ] Repository pattern for data access

### Project Organization

- [x] Standard project layout (`cmd/`, `internal/`, `pkg/`)
- [x] Configuration management
- [ ] Dependency injection
- [ ] Clean architecture principles
- [ ] Error handling strategies at scale

### Advanced Testing

- [ ] Integration tests with test databases
- [ ] Mocking interfaces (GoMock, testify/mock)
- [ ] Test doubles (stubs, spies, fakes)
- [ ] Test fixtures and factories
- [ ] HTTP testing with `httptest.Server`
- [ ] `httptest.ResponseRecorder` for handler tests
- [ ] Race detector (`-race` flag) in CI/CD
- [ ] Fuzzing with `go test -fuzz`
- [ ] Golden file testing
- [ ] Test containers for dependencies
- [ ] Parallel tests with `t.Parallel()`

### Performance & Optimization

- [ ] CPU profiling with `pprof`
- [ ] Memory profiling (heap, allocs)
- [ ] Goroutine profiling
- [ ] Block profiling for contention
- [ ] Mutex profiling
- [ ] Trace analysis with `go tool trace`
- [ ] Benchmarking and comparative analysis
- [ ] Identifying allocation hotspots
- [ ] Memory allocation patterns (pool, reuse)
- [ ] Escape analysis (`-gcflags=-m`)
- [ ] Garbage collection tuning (GOGC)
- [ ] Reducing pointer chasing
- [ ] String concatenation optimization

### Build & Tooling

- [x] `go mod` management
- [x] `golangci-lint`
- [ ] Build tags and conditional compilation
- [ ] Cross-compilation
- [ ] Vendoring
- [ ] Code generation

### Advanced Features

- [ ] Generics basics (type parameters)
- [ ] Type constraints and interfaces
- [ ] Generic functions and types
- [ ] `comparable` constraint
- [ ] `any` constraint
- [ ] Reflection (`reflect` package)
  - [ ] `reflect.TypeOf` and `reflect.ValueOf`
- [ ] Inspecting struct fields and tags
- [ ] Unsafe pointer operations (`unsafe` package)
  - [ ] Pointer arithmetic with unsafe
  - [ ] `unsafe.Sizeof`, `unsafe.Alignof`, `unsafe.Offsetof`
- [ ] Struct embedding (composition)
- [ ] Interface embedding
- [ ] Build constraints (`//go:build`)
- [ ] Platform-specific code

### Production-Ready Skills

- [ ] Structured logging with zap or logrus
- [ ] Log levels and conditional logging
- [ ] Context-aware logging
- [ ] Metrics and monitoring with Prometheus
- [ ] Counter, Gauge, Histogram, Summary metrics
- [ ] Custom metrics and labels
- [ ] Distributed tracing (OpenTelemetry, Jaeger)
- [ ] Span creation and propagation
- [ ] Graceful shutdown (signal handling)
- [ ] Draining connections and cleanup
- [ ] Health check endpoints (`/health`, `/ready`)
- [ ] Liveness vs readiness probes
- [ ] Docker containerization (Dockerfile)
- [ ] Multi-stage Docker builds
- [ ] Docker Compose for local development
- [ ] CI/CD pipelines (GitHub Actions, GitLab CI)
- [ ] Automated testing in CI
- [ ] Security scanning (gosec, trivy)

### Expert Projects

- [ ] Full microservice with database
- [ ] gRPC service with interceptors
- [ ] Message queue consumer/producer
- [ ] Custom HTTP framework
- [ ] Distributed system component
- [ ] CLI tool with cobra
- [ ] Plugin system with dynamic loading

---

## 🏆 Mastery Milestones

### Beginner → Intermediate

- [ ] Build 3 CLI tools without external dependencies
- [ ] Understand all slice edge cases
- [ ] Implement custom error types for a project

### Intermediate → Proficient

- [ ] Build concurrent system handling 1000+ goroutines
- [ ] Implement 5+ concurrency patterns from scratch
- [ ] Write comprehensive test suite with >80% coverage

### Proficient → Expert

- [ ] Build production-ready API with monitoring
- [ ] Optimize program using profiling tools
- [ ] Contribute to open-source Go project
- [ ] Architect and implement distributed system component
