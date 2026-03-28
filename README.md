# Go DSA Mastery: Complete Data Structures and Algorithms

A comprehensive, production-ready implementation of Data Structures and Algorithms in Go, designed for mastering DSA concepts and preparing for technical interviews.

## 🎯 Overview

This project provides a complete learning path for Data Structures and Algorithms in Go, featuring:

- **Comprehensive Coverage**: From basic complexity analysis to advanced graph algorithms
- **Production Quality**: Clean, efficient, and well-tested implementations
- **Educational Focus**: Clear documentation and examples for learning
- **Performance Optimized**: Benchmarked implementations with memory efficiency
- **Go Best Practices**: Idiomatic Go code with proper error handling

## 📚 Learning Path

### Phase 1: Fundamentals (✅ Complete)

- **Time & Space Complexity Analysis**
- **Big O Notation with practical examples**
- **Performance Benchmarking in Go**
- **Memory profiling and optimization**

### Phase 2: Data Structures (✅ Complete)

#### Linear Data Structures

- **Arrays & Slices**: Dynamic arrays with generic implementations
- **Stacks**: LIFO operations with multiple implementations
- **Queues**: FIFO operations including priority queues
- **Linked Lists**: Singly, doubly, and circular variants

#### Non-Linear Data Structures

- **Binary Trees**: Complete tree operations and traversals
- **Binary Search Trees**: Ordered operations
- **Heaps**: Min/max heaps with priority queue implementation

### Phase 3: Algorithms (✅ Complete)

#### Graph Algorithms (✅ Complete)

- **Graph Representations**: Adjacency list and matrix
- **Traversal**: DFS and BFS with complete result tracking
- **Shortest Paths**: Dijkstra's algorithm with priority queue
- **Minimum Spanning Tree**: Kruskal's and Prim's algorithms
- **Union-Find**: Disjoint set with path compression

#### Sorting Algorithms (✅ Complete)

- **Comparison-based**: QuickSort (3 variants), MergeSort, HeapSort, InsertionSort, SelectionSort, BubbleSort, IntroSort
- **Non-comparison**: CountingSort, RadixSort, BucketSort
- **Performance Analysis**: Comprehensive benchmarks vs Go standard library
- **Utility Functions**: Array generators, sorting validation, statistics tracking

#### Dynamic Programming (✅ Complete)

- **Classic Problems**: Fibonacci, 0/1 Knapsack, Unbounded Knapsack
- **String DP**: Longest Common Subsequence, Edit Distance, Palindrome problems
- **Array DP**: Longest Increasing Subsequence (O(n²) and O(n log n))
- **Optimization**: Coin Change (min coins & count ways), Matrix Chain Multiplication
- **Advanced**: Memoization vs Tabulation, Space optimization techniques

#### String Algorithms (✅ Complete)

- **Pattern Matching**: KMP, Rabin-Karp, Boyer-Moore, Z-algorithm, Naive search
- **String Processing**: Trie data structure with prefix operations
- **Hash Functions**: Custom string hashing, FNV hash, Rolling hash
- **Advanced**: Manacher's algorithm, Longest Common Prefix, String utilities

#### Backtracking Algorithms (✅ Complete)

- **Classic Problems**: N-Queens (all/single solutions), Sudoku solver with generation
- **Combinatorial**: Permutations, combinations, subsets (power set)
- **Constraint Satisfaction**: Graph coloring, maze solving
- **Advanced**: Knight's tour, Word break problem, optimization techniques

#### Search, Number Theory & Optimization (✅ Complete – see 04-advanced-topics)

- **Advanced Search**: Binary search variants (lower/upper bound, rotated array), ternary, exponential, interpolation search
- **Number Theory**: Sieve, primality tests, modular arithmetic, GCD, factorization, CRT
- **Optimization**: Golden section, ternary search, bisection, root finding

#### Practice Problems (✅ Complete – see 05-practice-problems)

- **70+ problems** by category: arrays/strings, data structures, graphs, DP, backtracking
- **CLI runner** in `05-practice-problems` (separate Go module)

## 📊 Progress Tracking

### Current Status: All 5 Phases Complete

**Completed modules:**

- ✅ **01-fundamentals**: Complexity analysis and Go optimization techniques
- ✅ **02-data-structures**: Linear and non-linear structures with comprehensive tests
- ✅ **03-algorithms**: Graphs, sorting, DP, string algorithms, backtracking (with examples per topic)
- ✅ **04-advanced-topics**: Search variants, number theory, numerical optimization
- ✅ **05-practice-problems**: 70+ problems with CLI runner (tests in arrays-strings/easy; more can be added)

**Follow the path:** See **[LEARNING_PATH.md](LEARNING_PATH.md)** for the recommended order and how to verify each phase.

### Milestone Achievements

- [x] 650+ comprehensive unit tests across all modules
- [x] Performance benchmarks for all algorithms
- [x] Complete documentation with learning guides
- [x] Production-ready implementations with error handling
- [x] Memory optimization and Go best practices

## 🏗️ Project Structure

```
dsa-mastery/
├── 01-fundamentals/               # Complexity analysis and Go optimization
│   ├── complexity.go              # Time/space complexity examples
│   ├── complexity_test.go         # Comprehensive tests
│   ├── main.go                    # Local demo (use examples/fundamentals/main.go to run)
│   ├── complexity-analysis.md     # Learning guide
│   └── README.md                  # Fundamentals guide
│
├── 02-data-structures/
│   ├── linear/                    # Sequential data structures
│   │   ├── arrays.go              # Dynamic arrays and slices
│   │   ├── stack.go               # LIFO implementations
│   │   ├── queue.go               # FIFO and priority queues
│   │   ├── linked_list.go         # All linked list variants
│   │   └── *_test.go             # Comprehensive test suites
│   │
│   └── non-linear/                # Hierarchical structures
│       ├── tree.go                # Binary tree, BST, heap, generic BST
│       └── tree_test.go           # Full test coverage
│
├── 03-algorithms/                 # See 03-algorithms/README.md for full tree
│   ├── graphs/
│   ├── sorting/
│   ├── dynamic-programming/
│   ├── string-algorithms/
│   └── backtracking/
│
├── 04-advanced-topics/            # Search, number theory, optimization
│   ├── search/
│   ├── number-theory/
│   └── optimization/
│
├── 05-practice-problems/          # Separate Go module; run from that directory
│   ├── main.go                    # CLI: go run main.go <command>
│   ├── arrays-strings/
│   ├── data-structures/
│   ├── graphs/
│   ├── dynamic-programming/
│   └── backtracking/
│
└── examples/                      # Runnable demos (go run examples/<name>/main.go)
    ├── fundamentals/
    ├── data-structures/
    ├── sorting/
    ├── graphs/
    ├── dynamic-programming/
    ├── string-algorithms/
    ├── backtracking/
    └── advanced-topics/
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24+ (uses generics for type safety)
- Basic understanding of Go syntax

### Installation

```bash
# Clone the repository  
git clone https://github.com/your-username/dsa-mastery.git
cd dsa-mastery

# Initialize Go module
go mod init dsa-mastery

# Run all tests (root module only; 05-practice-problems has its own module)
go test ./...

# Run benchmarks
go test -bench=. -benchmem ./...
```

**Note:** `05-practice-problems` is a separate Go module. Run its tests with `cd 05-practice-problems && go test ./...`.

### 🏃‍♂️ Running Individual Modules

Most source files (like `complexity.go`) are packages, not directly executable. Here are three ways to run and explore each module:

#### Option 1: Run Interactive Examples (Recommended)

```bash
# Fundamentals - Complexity Analysis
go run examples/fundamentals/main.go

# Data Structures Examples
go run examples/data-structures/main.go

# Graph Algorithms Demo  
go run examples/graphs/main.go

# Sorting Algorithms Demo
go run examples/sorting/main.go

# Dynamic Programming Demo
go run examples/dynamic-programming/main.go

# String Algorithms Demo
go run examples/string-algorithms/main.go

# Backtracking Demo
go run examples/backtracking/main.go

# Advanced Topics (search, number theory)
go run examples/advanced-topics/main.go
```

#### Option 2: Run Module Tests

```bash
# Run fundamentals tests
go test ./01-fundamentals/ -v

# Run specific package tests  
go test -v ./02-data-structures/linear
go test -v ./03-algorithms/graphs

# Run specific test function
go test -run TestBinarySearch -v ./01-fundamentals
```

#### Option 3: Run Performance Benchmarks

```bash
# Run fundamentals benchmarks
go test ./01-fundamentals/ -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkBinarySearch -benchmem ./01-fundamentals

# Run all benchmarks with memory stats
go test -bench=. -benchmem ./...
```

### Quick Start Examples

#### Working with Data Structures

```go
package main

import (
    "fmt"
    "dsa-mastery/02-data-structures/linear"
    nonlinear "dsa-mastery/02-data-structures/non-linear"
)

func main() {
    // Generic stack (linear)
    stack := linear.NewGenericStack[int]()
    stack.Push(1)
    stack.Push(2)
    value, _ := stack.Pop() // value = 2

    // Binary Search Tree (non-linear)
    bst := nonlinear.NewBinarySearchTree()
    bst.Insert(10)
    bst.Insert(5)
    bst.Insert(15)
    found := bst.Search(10) // true

    // Min Heap (non-linear)
    heap := nonlinear.NewMinHeap()
    heap.Insert(20)
    heap.Insert(10)
    min, _ := heap.ExtractMin() // min = 10
}
```

#### Graph Algorithms

```go
package main

import (
    "fmt"
    "dsa-mastery/03-algorithms/graphs"
)

func main() {
    // Create graph
    g := graphs.NewAdjacencyListGraph(false)
    
    // Add vertices and edges
    for i := 0; i < 4; i++ {
        g.AddVertex(i)
    }
    g.AddEdge(0, 1, 4)
    g.AddEdge(0, 2, 2)
    g.AddEdge(1, 3, 5)
    g.AddEdge(2, 3, 1)

    // Find shortest paths
    result, _ := graphs.Dijkstra(g, 0)
    fmt.Printf("Shortest distances: %v\n", result.Distances)

    // Find MST
    mst, _ := graphs.Kruskal(g)
    fmt.Printf("MST cost: %d\n", mst.TotalCost)
}
```

## 🧪 Testing & Benchmarking

The project includes comprehensive test suites and performance benchmarks. Refer to the **"🏃‍♂️ Running Individual Modules"** section above for detailed testing commands.

### Performance Results

Sample benchmark results on Apple M3 Pro:

```
Linear Data Structures:
BenchmarkGenericStackPush-11      50000000    22.1 ns/op      8 B/op    0 allocs/op
BenchmarkQueueEnqueue-11          30000000    38.2 ns/op     24 B/op    1 allocs/op
BenchmarkLinkedListInsert-11      20000000    65.4 ns/op     24 B/op    1 allocs/op

Non-Linear Data Structures:
BenchmarkBSTInsert-11             10000000   120.5 ns/op     24 B/op    1 allocs/op
BenchmarkHeapInsert-11            15000000    78.3 ns/op      8 B/op    0 allocs/op
BenchmarkHashTableSet-11          30000000    42.1 ns/op     16 B/op    1 allocs/op

Graph Algorithms:
BenchmarkDFS-11                      49670    22361 ns/op  23198 B/op  174 allocs/op
BenchmarkDijkstra-11                 60226    19689 ns/op   7271 B/op   42 allocs/op
BenchmarkKruskal-11                  68161    17419 ns/op  19686 B/op   56 allocs/op
```

## 📖 Learning path and root docs

| Doc | Purpose |
|-----|--------|
| **[LEARNING_PATH.md](LEARNING_PATH.md)** | **Start here.** Ordered path (01 → 05), verify commands, and example runs. |
| [GOLANG_LEARNING_GUIDE.md](GOLANG_LEARNING_GUIDE.md) | Go language primer (syntax, env, concurrency) – use if you're new to Go. |
| [QUICK_REFERENCE.md](QUICK_REFERENCE.md) | Code patterns and complexity cheat sheet. |
| [IMPROVEMENT_PLAN.md](IMPROVEMENT_PLAN.md) | Backlog of improvements and test coverage. |
| [PROGRESS.md](PROGRESS.md) | High-level completion report and stats. |

### Documentation (in-repo)

- Each module includes comprehensive README files
- Inline code documentation with examples
- Test files serve as usage examples

### Key Concepts Covered

#### Complexity Analysis

- Big O notation with practical examples
- Time vs Space complexity trade-offs
- Amortized analysis for advanced data structures

#### Implementation Patterns

- Generic programming with Go 1.24+ generics
- Interface-based design for flexibility
- Error handling with Go idioms
- Memory-efficient implementations

#### Algorithm Design

- Recursive vs iterative approaches
- Divide and conquer strategies
- Dynamic programming foundations
- Greedy algorithm principles

## 🏆 Features

### Production Ready

- **Comprehensive Error Handling**: Proper error types and messages
- **Memory Efficient**: Optimized for minimal allocations
- **Thread Safety**: Where applicable, with proper synchronization
- **Generic Types**: Type-safe implementations using Go generics

### Educational Value

- **Step-by-step Implementation**: Clear, readable code with comments
- **Multiple Approaches**: Different solutions for learning comparison
- **Real-world Examples**: Practical applications of each algorithm
- **Progressive Complexity**: From basic to advanced implementations

### Performance Optimized

- **Benchmarked**: All implementations include performance benchmarks
- **Profiled**: Memory and CPU usage optimization
- **Scalable**: Tested with various input sizes
- **Efficient**: Using best practices for Go performance

## 🎓 Use Cases

### Interview Preparation

- Complete coverage of common DSA interview questions
- Multiple implementation approaches for flexibility
- Time/space complexity analysis for each solution
- Real interview question examples and solutions

### Educational Projects

- University course assignments and projects
- Self-study and skill development
- Teaching material for algorithms courses
- Code review and best practices examples

### Production Development

- Reusable components for Go applications
- Performance-critical algorithm implementations
- Reference implementations for complex algorithms
- Foundation for specialized data structure needs

## 🔬 Advanced Topics

### Optimization Techniques

- **Path Compression**: Union-Find optimization
- **Priority Queues**: Heap-based efficient implementations
- **Memory Pooling**: Reducing allocation overhead
- **Cache-friendly**: Data structure layout optimization

### Go-Specific Optimizations

- **Slice Management**: Efficient slice operations and growth
- **Interface Usage**: When to use interfaces vs concrete types
- **Goroutine Safety**: Concurrent data structure design
- **Memory Profiling**: Using Go tools for optimization

## 📈 Performance Guidelines

### Benchmarking Best Practices

1. **Consistent Environment**: Run benchmarks on consistent hardware
2. **Multiple Runs**: Average results across multiple benchmark runs
3. **Memory Profiling**: Monitor allocations and GC pressure
4. **Real-world Data**: Test with realistic input sizes and patterns

### Optimization Strategies

1. **Algorithm Choice**: Select optimal algorithm for use case
2. **Data Structure Selection**: Match structure to access patterns
3. **Memory Management**: Minimize allocations and copying
4. **Concurrency**: Leverage Go's concurrency where beneficial

## 🤝 Contributing

### Code Standards

- Follow Go conventions and gofmt formatting
- Include comprehensive tests for new implementations
- Add benchmarks for performance-critical code
- Update documentation for new features

### Testing Requirements

- Unit tests with >95% coverage
- Benchmark tests for performance validation
- Edge case testing for robustness
- Example functions for documentation

## 📄 License

This project is created for educational purposes. Feel free to use, modify, and learn from the implementations.

## 🔗 Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Algorithm Visualization](https://visualgo.net/)
- [Big O Cheat Sheet](https://www.bigocheatsheet.com/)
- [Go Performance Tips](https://github.com/golang/go/wiki/Performance)

---

**Happy Learning! 🚀**

Master Data Structures and Algorithms with production-quality Go implementations.
