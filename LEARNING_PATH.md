# DSA Mastery – Learning Path

Follow this path in order to master Data Structures and Algorithms in Go. Each phase builds on the previous one.

---

## Prerequisites

- **Go 1.24+** installed ([go.dev/dl](https://go.dev/dl/))
- **Go basics**: syntax, functions, structs, slices, error handling  
  → If you're new to Go, start with [GOLANG_LEARNING_GUIDE.md](GOLANG_LEARNING_GUIDE.md) (Phases 1–2) or any Go tutorial.

---

## The path (in order)

| Phase | Folder | What you'll learn | Verify |
|-------|--------|-------------------|--------|
| **1** | [01-fundamentals](01-fundamentals/) | Complexity (Big O), time/space, benchmarking | `go test ./01-fundamentals/ -v` |
| **2** | [02-data-structures](02-data-structures/) | Arrays, stacks, queues, linked lists, trees, heaps, BST | `go test ./02-data-structures/... -v` |
| **3** | [03-algorithms](03-algorithms/) | Graphs, sorting, DP, strings, backtracking | `go test ./03-algorithms/... -v` |
| **4** | [04-advanced-topics](04-advanced-topics/) | Binary search variants, number theory, numerical optimization | `go test ./04-advanced-topics/... -v` |
| **5** | [05-practice-problems](05-practice-problems/) | 70+ problems by category (arrays, DS, graphs, DP, backtracking) | `cd 05-practice-problems && go test ./... -v` |

---

## How to use each phase

1. **Read** the folder’s `README.md` (and any `*-analysis.md` if present).
2. **Run tests** to confirm everything passes (see Verify column or the README).
3. **Run the example** for that topic (see below).
4. **Read and run** the code in the `.go` files; use tests as usage examples.

---

## Running examples

From the **repository root** (except 05, which has its own module):

```bash
# Phase 1
go run examples/fundamentals/main.go

# Phase 2
go run examples/data-structures/main.go

# Phase 3
go run examples/graphs/main.go
go run examples/sorting/main.go
go run examples/dynamic-programming/main.go
go run examples/string-algorithms/main.go
go run examples/backtracking/main.go

# Phase 4
go run examples/advanced-topics/main.go

# Phase 5 (run from inside the folder)
cd 05-practice-problems
go run main.go help
go run main.go arrays-easy
go run main.go all
```

---

## Quick reference

- **One-page code patterns & complexity**: [QUICK_REFERENCE.md](QUICK_REFERENCE.md)
- **Project overview, structure, and setup**: [README.md](README.md)

---

## Suggested timeline

| Phase | Suggested duration | Focus |
|-------|--------------------|--------|
| 1 – Fundamentals | 3–5 days | Big O, measuring performance in Go |
| 2 – Data structures | 1–2 weeks | Implement and use each structure |
| 3 – Algorithms | 2–3 weeks | Graphs, sorting, DP, strings, backtracking |
| 4 – Advanced topics | 3–5 days | Search variants, number theory, optimization |
| 5 – Practice problems | Ongoing | Apply concepts; do easy → medium by category |

Adjust by experience; use tests and examples to check your understanding at each step.
