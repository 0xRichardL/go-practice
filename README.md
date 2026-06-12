# Go Practice

This repository is a personal Go practice workspace. It contains small runnable
concept demos, algorithm packages, coding challenge solutions, and interview
exercises.

## Layout

| Path | Purpose | How to run |
| --- | --- | --- |
| `concepts/basics/` | Language fundamentals and standard library demos | `go run ./concepts/basics/slice/append-behavior` |
| `concepts/concurrency/` | Goroutines, channels, sync patterns, and GC demos | `go run ./concepts/concurrency/channel/timeout` |
| `algorithms/` | Reusable algorithm implementations with tests | `go test ./algorithms/...` |
| `problems/leetcode/` | LeetCode solutions kept as runnable packages | `go run ./problems/leetcode/coin-change` |
| `problems/hackerrank/` | HackerRank solutions kept as runnable packages | `go run ./problems/hackerrank/hack-the-interview-vi-asia-pacific/maximum-sum-10-1` |
| `interviews/` | Interview practice exercises | `go run ./interviews/andpad/t1` |

## Conventions

- Each standalone demo or problem solution lives in its own directory.
- Runnable examples use `package main` and can be executed with `go run ./path`.
- Reusable code lives in a normal package and should have `_test.go` coverage.
- Keep tests beside the code they test.
- Prefer adding a new small directory over adding another `main` function to an
  existing directory.

## Verification

Run algorithm tests:

```bash
go test ./algorithms/...
```

Run the whole workspace compile/test pass:

```bash
go test ./...
```

Spot-check runnable examples:

```bash
go run ./concepts/basics/loop/string
go run ./concepts/basics/slice/append-behavior
go run ./concepts/concurrency/goroutine/waitgroup
go run ./problems/leetcode/coin-change
go run ./problems/hackerrank/hack-the-interview-vi-asia-pacific/maximum-sum-10-1
```

Some demos are intentionally unsafe, blocking, failing, or process-ending. Run
examples such as deadlock, panic, CLI exit, and unsafe channel demos manually
when you specifically want to observe that behavior.
