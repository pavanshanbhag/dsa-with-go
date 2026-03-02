# 03-algorithms

Graph, sorting, dynamic programming, string, and backtracking algorithms in Go.

## Module structure

```
03-algorithms/
├── README.md                     # This file
├── graphs/
│   ├── README.md
│   ├── graph.go                  # Graph interface, adjacency list
│   ├── traversal.go              # DFS, BFS
│   ├── shortest_path.go          # Dijkstra
│   ├── mst.go                    # Kruskal, Prim, Union-Find
│   ├── graph_test.go
│   └── mst_test.go
├── sorting/
│   ├── README.md
│   ├── sorting.go                # QuickSort, MergeSort, HeapSort, etc.
│   └── sorting_test.go
├── dynamic-programming/
│   ├── README.md
│   ├── dp.go                     # Fibonacci, Knapsack, LCS, LIS, etc.
│   └── dp_test.go
├── string-algorithms/
│   ├── README.md
│   ├── string_algorithms.go      # KMP, Rabin-Karp, Trie, etc.
│   └── string_algorithms_test.go
└── backtracking/
    ├── README.md
    ├── backtracking.go           # N-Queens, Sudoku, permutations, etc.
    └── backtracking_test.go
```

## How to verify

From the **repository root**:

```bash
# All algorithm tests
go test ./03-algorithms/graphs/ ./03-algorithms/sorting/ ./03-algorithms/dynamic-programming/ ./03-algorithms/string-algorithms/ ./03-algorithms/backtracking/ -v

# With benchmarks
go test ./03-algorithms/... -bench=. -benchmem
```

Individual packages:

```bash
go test ./03-algorithms/graphs/ -v
go test ./03-algorithms/sorting/ -v
go test ./03-algorithms/dynamic-programming/ -v
go test ./03-algorithms/string-algorithms/ -v
go test ./03-algorithms/backtracking/ -v
```

## Running examples

Each algorithm subfolder has a corresponding demo in `examples/`. Run from the **repository root**:

```bash
go run examples/graphs/main.go
go run examples/sorting/main.go
go run examples/dynamic-programming/main.go
go run examples/string-algorithms/main.go
go run examples/backtracking/main.go
```

## Completeness checklist

| Item | Status |
|------|--------|
| Tests for all 5 packages | ✅ `go test ./03-algorithms/...` |
| Benchmarks | ✅ `go test ./03-algorithms/... -bench=. -benchmem` |
| Example runner per topic | ✅ graphs, sorting, dp, string-algorithms, backtracking in `examples/` |
| README per subfolder | ✅ With module files and verify commands |
