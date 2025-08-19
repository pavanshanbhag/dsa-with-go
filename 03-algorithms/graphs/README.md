# Graph Algorithms Module

This module provides comprehensive implementations of fundamental graph algorithms in Go, including graph data structures, traversal algorithms, shortest path algorithms, and minimum spanning tree algorithms.

## Features

### Graph Data Structures
- **Adjacency List Graph**: Memory-efficient representation for sparse graphs
- **Support for both directed and undirected graphs**
- **Weighted edges support**
- **Dynamic vertex and edge addition/removal**

### Graph Traversal Algorithms
- **Depth-First Search (DFS)**: Complete traversal with timing information
- **Breadth-First Search (BFS)**: Level-order traversal with distance tracking
- **Connected Components**: Check graph connectivity

### Shortest Path Algorithms
- **Dijkstra's Algorithm**: Single-source shortest paths for non-negative weights
- **Priority Queue optimization** for efficient path finding

### Minimum Spanning Tree (MST) Algorithms
- **Kruskal's Algorithm**: Edge-based MST using Union-Find
- **Prim's Algorithm**: Vertex-based MST with priority queue
- **Union-Find Data Structure**: Efficient disjoint set operations

## Usage Examples

### Creating a Graph

```go
// Create an undirected graph
graph := NewAdjacencyListGraph(false)

// Add vertices
for i := 0; i < 5; i++ {
    graph.AddVertex(i)
}

// Add weighted edges
graph.AddEdge(0, 1, 4)
graph.AddEdge(0, 2, 2)
graph.AddEdge(1, 2, 1)
graph.AddEdge(1, 3, 5)
graph.AddEdge(2, 3, 8)
graph.AddEdge(2, 4, 10)
graph.AddEdge(3, 4, 2)
```

### Graph Traversal

```go
// Depth-First Search
dfsResult, err := DFS(graph, 0)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("DFS Order: %v\n", dfsResult.Order)
fmt.Printf("Discovery Times: %v\n", dfsResult.DiscoveryTime)

// Breadth-First Search
bfsResult, err := BFS(graph, 0)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("BFS Order: %v\n", bfsResult.Order)
fmt.Printf("Distances: %v\n", bfsResult.Distance)
```

### Shortest Path Algorithms

```go
// Find shortest paths from vertex 0
result, err := Dijkstra(graph, 0)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Shortest distances: %v\n", result.Distances)
fmt.Printf("Parent relationships: %v\n", result.Parents)
```

### Minimum Spanning Tree

```go
// Kruskal's Algorithm
mstKruskal, err := Kruskal(graph)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("MST Cost (Kruskal): %d\n", mstKruskal.TotalCost)
fmt.Printf("MST Edges: %v\n", mstKruskal.Edges)

// Prim's Algorithm
mstPrim, err := Prim(graph, 0)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("MST Cost (Prim): %d\n", mstPrim.TotalCost)
fmt.Printf("MST Edges: %v\n", mstPrim.Edges)
```

### Union-Find Operations

```go
// Create Union-Find for vertices 0-4
vertices := []int{0, 1, 2, 3, 4}
uf := NewUnionFind(vertices)

// Union operations
uf.Union(0, 1)
uf.Union(2, 3)

// Check connectivity
connected := uf.IsConnected(0, 1) // true
notConnected := uf.IsConnected(0, 2) // false

// Find root
root := uf.Find(0)
```

## Algorithm Complexities

### Graph Operations
- **Add Vertex**: O(1)
- **Add Edge**: O(1)
- **Remove Vertex**: O(V + E)
- **Remove Edge**: O(1)
- **Get Neighbors**: O(degree)

### Traversal Algorithms
- **DFS**: O(V + E) time, O(V) space
- **BFS**: O(V + E) time, O(V) space

### Shortest Path
- **Dijkstra**: O((V + E) log V) time, O(V) space

### Minimum Spanning Tree
- **Kruskal**: O(E log E) time, O(V) space
- **Prim**: O((V + E) log V) time, O(V) space
- **Union-Find**: Nearly O(1) amortized per operation

## Performance Benchmarks

Based on benchmarks with various graph sizes:

```
BenchmarkDFS-11                    49670    22361 ns/op    23198 B/op    174 allocs/op
BenchmarkBFS-11                    43846    24588 ns/op    23993 B/op    273 allocs/op
BenchmarkDijkstra-11               60226    19689 ns/op     7271 B/op     42 allocs/op
BenchmarkKruskal-11                68161    17419 ns/op    19686 B/op     56 allocs/op
BenchmarkPrim-11                   51849    22834 ns/op     7911 B/op     44 allocs/op
BenchmarkUnionFind/Union-11      16383016    72.68 ns/op     160 B/op      2 allocs/op
BenchmarkUnionFind/Find-11       45938143    26.28 ns/op       0 B/op      0 allocs/op
```

## Key Features

### 1. Memory Efficiency
- Adjacency list representation minimizes memory usage for sparse graphs
- Efficient data structures for priority queues and union-find

### 2. Error Handling
- Comprehensive error checking for invalid operations
- Clear error messages for debugging

### 3. Comprehensive Testing
- Unit tests covering all algorithms and edge cases
- Performance benchmarks for optimization insights
- Example functions for documentation

### 4. Type Safety
- Strong typing with Go's type system
- Clear interfaces for graph operations

### 5. Production Ready
- Optimized implementations with best practices
- Path compression in Union-Find for nearly constant-time operations
- Priority queue optimization for Dijkstra and Prim algorithms

## Applications

### Graph Traversal
- **Connected Components**: Network analysis, social graphs
- **Cycle Detection**: Dependency resolution, deadlock detection
- **Topological Sorting**: Task scheduling, build systems

### Shortest Paths
- **Route Planning**: GPS navigation, network routing
- **Social Networks**: Degrees of separation, influence analysis
- **Game AI**: Pathfinding in game worlds

### Minimum Spanning Trees
- **Network Design**: Minimizing cable/pipeline costs
- **Clustering**: Data analysis and machine learning
- **Circuit Design**: Minimizing wire length in electronics

## Testing

Run all tests:
```bash
go test -v
```

Run benchmarks:
```bash
go test -bench=. -benchmem
```

Run specific test:
```bash
go test -run TestDFS -v
```

## Implementation Notes

### Design Decisions
1. **Interface-based design**: Graph interface allows multiple implementations
2. **Separate result structures**: Clear separation of algorithm inputs and outputs
3. **Error handling**: Go idioms for error handling throughout
4. **Memory optimization**: Efficient data structures for large graphs

### Optimizations
1. **Path compression** in Union-Find for near O(1) operations
2. **Priority queue** implementation for Dijkstra and Prim
3. **Adjacency list** for memory-efficient sparse graph representation
4. **Early termination** in algorithms where applicable

This module provides a solid foundation for graph-based algorithms in Go applications, with production-ready implementations suitable for educational and commercial use.
