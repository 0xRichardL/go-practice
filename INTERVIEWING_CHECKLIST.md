# Golang Interview Skills Checklist

> Grouped by probability of appearance in real-world Go backend interviews.

## 🟢 Tier 1 — Almost Guaranteed (70–90%)

### Go Basics

- `var` vs `:=`
- Zero values
- Value vs reference semantics
- Struct memory layout (basic)
- Export rules
- init() behavior

### Slices & Arrays

- Slice header (ptr, len, cap)
- append reallocation
- Slice sharing bugs
- Full slice expression `[low:high:max]`
- copy semantics
- `nil` vs empty slice

### Maps

- Not thread-safe
- `nil` map write panic
- Lookup patterns (v, ok)
- Random iteration order
- Rehashing behavior

### Goroutines

- M:P:G model (high level)
- Lifecycle
- Leaks
- Closure capture bugs
- `GOMAXPROCS`

### Channels

- Buffered vs unbuffered
- Blocking behavior
- Close semantics
- Read from closed channel
- Send on closed channel panic
- Fan-in / fan-out

### Mutex & RWMutex

- Mutex vs RWMutex
- Deadlocks
- Lock ordering
- RWMutex caveats

### Error Handling

- Idiomatic errors
- Wrapping (%w)
- errors.Is / errors.As
- Custom error types
- Panic rules

### Interfaces

- Implicit implementation
- Interface internals (type, value)
- Nil interface trap
- Small interfaces
- Accept interface, return struct

### Context

- Cancellation
- WithTimeout / WithCancel
- Context misuse
- Goroutine cleanup

## 🟡 Tier 2 — High Chance (40–70%)

### Memory & GC

- Stack vs heap
- Escape analysis
- Allocation reduction
- sync.Pool tradeoffs

### Scheduler & Runtime

- Preemption
- Blocking syscalls
- Goroutine vs OS thread

### Concurrency Patterns

- Worker pools
- Rate limiting
- Semaphores
- Pipelines

### Testing

- Table-driven tests
- Subtests
- Race detector
- Testing concurrency

### Build & Tooling

- go mod
- go mod tidy
- replace
- go vet
- golangci-lint

### HTTP & Networking

- net/http lifecycle
- Middleware
- Request context
- Client reuse

### Performance

- pprof
- Benchmarks
- Hot path analysis

## 🟠 Tier 3 — Medium Chance (20–40%)

- unsafe (conceptual)
- cgo (conceptual)
- reflection
- generics
- sync.Once / Cond / atomic

## 🔵 Tier 4 — Nice to Have (≤20%)

- Compiler internals (SSA, inlining)
- Advanced runtime topics
- Production readiness (graceful shutdown, tracing)

## 🎯 Interview Reality

- Most Go interviews focus on concurrency
- Production bug avoidance matters more than syntax
- Clear tradeoff explanations = strong signal
