# Backtracking Algorithms

A comprehensive implementation of backtracking algorithms for solving constraint satisfaction problems and combinatorial optimization challenges.

## Overview

Backtracking is a systematic method for solving problems by incrementally building solutions and abandoning ("backtracking") partial solutions that cannot lead to valid complete solutions. This approach is particularly effective for constraint satisfaction problems, combinatorial enumeration, and optimization challenges.

## Algorithms Implemented

### 1. **N-Queens Problem**
- **Purpose**: Place N queens on an N×N chessboard so no two queens attack each other
- **Time Complexity**: O(N!)
- **Space Complexity**: O(N²)
- **Features**:
  - Find all solutions or just one solution
  - Count total solutions without storing them
  - Visual board representation

```go
// Find all solutions for 8-Queens
solutions := SolveNQueens(8)
fmt.Printf("Found %d solutions for 8-Queens\n", len(solutions))

// Find just one solution (faster)
solution := SolveNQueensOne(8)
if solution != nil {
    fmt.Println(solution.String())
}

// Count solutions without storing them
count := CountNQueensSolutions(8)
fmt.Printf("Total solutions: %d\n", count)
```

### 2. **Sudoku Solver**
- **Purpose**: Solve 9×9 Sudoku puzzles using constraint satisfaction
- **Time Complexity**: O(9^(n²)) where n=9
- **Space Complexity**: O(n²)
- **Features**:
  - Validate existing puzzles
  - Generate new puzzles with configurable difficulty
  - Visual board representation

```go
// Solve a Sudoku puzzle
puzzle := &SudokuBoard{
    {5, 3, 0, 0, 7, 0, 0, 0, 0},
    {6, 0, 0, 1, 9, 5, 0, 0, 0},
    // ... rest of puzzle
}

if SolveSudoku(puzzle) {
    fmt.Println("Solved:")
    fmt.Println(puzzle.String())
}

// Generate a new puzzle
newPuzzle := GenerateSudoku(40) // 40 empty cells
```

### 3. **Combinatorial Problems**

#### Permutations
Generate all possible arrangements of elements:
```go
arr := []int{1, 2, 3}
perms := GeneratePermutations(arr)
// Returns: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
```

#### Combinations
Generate all possible selections of k elements:
```go
arr := []int{1, 2, 3, 4}
combs := GenerateCombinations(arr, 2)
// Returns: [[1,2], [1,3], [1,4], [2,3], [2,4], [3,4]]
```

#### Subsets
Generate all possible subsets (power set):
```go
arr := []int{1, 2, 3}
subsets := GenerateSubsets(arr)
// Returns: [[], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3]]
```

### 4. **Graph Coloring**
- **Purpose**: Color graph vertices so no adjacent vertices share the same color
- **Time Complexity**: O(m^n) where m=colors, n=vertices
- **Space Complexity**: O(n)

```go
// Create a graph and add edges
gc := NewGraphColoring(4, 3) // 4 vertices, 3 colors
gc.AddEdge(0, 1)
gc.AddEdge(1, 2)
gc.AddEdge(2, 3)
gc.AddEdge(3, 0)

if gc.SolveColoring() {
    colors := gc.GetColoring()
    fmt.Printf("Coloring: %v\n", colors)
}
```

### 5. **Maze Solving**
- **Purpose**: Find a path from start to end in a maze
- **Time Complexity**: O(4^(mn))
- **Space Complexity**: O(mn)

```go
// Define maze (0=path, 1=wall)
grid := [][]int{
    {0, 1, 0, 0, 0},
    {0, 1, 0, 1, 0},
    {0, 0, 0, 1, 0},
    {1, 1, 0, 0, 0},
    {0, 0, 0, 1, 0},
}

maze := NewMaze(grid, 0, 0, 4, 4) // Start at (0,0), end at (4,4)
if maze.SolveMaze() {
    fmt.Println(maze.String())
}
```

### 6. **Knight's Tour**
- **Purpose**: Move a knight to visit every square on a chessboard exactly once
- **Time Complexity**: O(8^(n²))
- **Space Complexity**: O(n²)

```go
kt := NewKnightsTour(8)
if kt.SolveKnightsTour(0, 0) {
    fmt.Println("Knight's Tour solution:")
    fmt.Println(kt.String())
}
```

### 7. **Word Break Problem**
- **Purpose**: Segment a string into dictionary words
- **Time Complexity**: O(2^n)
- **Space Complexity**: O(n)

```go
s := "catsanddog"
wordDict := []string{"cat", "cats", "and", "sand", "dog"}
segments := WordBreak(s, wordDict)
// Returns: ["cat sand dog", "cats and dog"]
```

## Key Concepts

### Backtracking Template
The general backtracking approach follows this pattern:

```go
func backtrack(state, choices) bool {
    if isComplete(state) {
        processResult(state)
        return true
    }
    
    for choice := range choices {
        if isValid(state, choice) {
            makeChoice(state, choice)
            
            if backtrack(state, remainingChoices) {
                return true // or continue for all solutions
            }
            
            undoChoice(state, choice) // backtrack
        }
    }
    
    return false
}
```

### Optimization Techniques

1. **Constraint Propagation**: Eliminate invalid choices early
2. **Heuristics**: Choose most constrained variables first
3. **Pruning**: Cut off branches that cannot lead to solutions
4. **Memoization**: Cache partial results when applicable

### Problem Categories

1. **Decision Problems**: Find any valid solution
2. **Optimization Problems**: Find the best solution
3. **Enumeration Problems**: Find all solutions
4. **Counting Problems**: Count the number of solutions

## Performance Characteristics

| Algorithm | Best Case | Average Case | Worst Case | Space |
|-----------|-----------|--------------|------------|-------|
| N-Queens | O(N!) | O(N!) | O(N!) | O(N²) |
| Sudoku | O(1) | O(9^k) | O(9^81) | O(1) |
| Permutations | O(N!) | O(N!) | O(N!) | O(N!) |
| Graph Coloring | O(N) | O(C^N) | O(C^N) | O(N) |
| Maze Solving | O(N) | O(4^NM) | O(4^NM) | O(NM) |

Where:
- N = problem size (board size, array length, etc.)
- C = number of colors
- k = number of empty cells in Sudoku
- M = maze width

## Real-World Applications

### N-Queens
- **Resource Allocation**: Scheduling non-conflicting resources
- **VLSI Design**: Placing components without interference
- **Constraint Satisfaction**: General CSP problem template

### Sudoku Solver
- **Logic Puzzles**: Educational games and brain training
- **Constraint Programming**: Testing CSP algorithms
- **AI Research**: Benchmark for search algorithms

### Combinatorial Generation
- **Test Case Generation**: Exhaustive testing scenarios
- **Cryptography**: Key space exploration
- **Bioinformatics**: Sequence analysis and permutation studies

### Graph Coloring
- **Register Allocation**: Compiler optimization
- **Scheduling**: Time slot assignment
- **Map Coloring**: Geographic visualization
- **Network Frequency Assignment**: Radio spectrum allocation

### Maze Solving
- **Robotics**: Path planning and navigation
- **Game AI**: NPC movement and pathfinding
- **Circuit Design**: Wire routing optimization

### Knight's Tour
- **Mathematical Recreation**: Puzzle solving
- **Algorithm Education**: Teaching backtracking concepts
- **Chess Programming**: Move sequence optimization

## Learning Path

### Beginner Level
1. **Start with N-Queens (4×4)**: Understand basic backtracking
2. **Simple Permutations**: Learn state space exploration
3. **Basic Maze Solving**: Practice with small grids

### Intermediate Level
1. **Sudoku Solver**: Constraint satisfaction techniques
2. **Graph Coloring**: Multiple constraint handling
3. **Larger N-Queens**: Optimization and pruning

### Advanced Level
1. **Knight's Tour**: Complex state spaces
2. **Word Break**: String-based backtracking
3. **Optimization Problems**: Branch and bound techniques

### Expert Level
1. **Hybrid Algorithms**: Combining backtracking with other techniques
2. **Parallel Backtracking**: Multi-threaded implementations
3. **Custom Constraint Systems**: Domain-specific problems

## Best Practices

### Algorithm Design
1. **Define State Clearly**: What constitutes a partial solution?
2. **Identify Constraints**: What makes a choice invalid?
3. **Choose Representation**: Efficient data structures for state
4. **Plan Backtracking**: How to undo choices efficiently?

### Optimization Strategies
1. **Pruning**: Eliminate impossible branches early
2. **Ordering**: Try most promising choices first
3. **Constraint Propagation**: Reduce search space
4. **Memoization**: Cache results when beneficial

### Implementation Tips
1. **Validate Input**: Check preconditions
2. **Handle Edge Cases**: Empty inputs, impossible problems
3. **Memory Management**: Avoid unnecessary allocations
4. **Testing**: Verify with known solutions

## Common Pitfalls

1. **Infinite Loops**: Ensure progress in each recursive call
2. **Incomplete Backtracking**: Not undoing all changes
3. **Memory Leaks**: Proper cleanup in recursive algorithms
4. **Performance**: Exponential growth without optimization

## Testing Strategy

Our test suite includes:
- **Correctness**: Verify solutions meet all constraints
- **Completeness**: Test edge cases and boundary conditions
- **Performance**: Benchmark with known problem sizes
- **Stress Testing**: Large inputs and complex scenarios

## Usage Examples

See the `demo` directory for comprehensive examples of each algorithm in action, including:
- Interactive problem solving
- Performance comparisons
- Visualization of solution processes
- Real-world problem applications

## References

1. **Classic Algorithms**: Knuth's "The Art of Computer Programming"
2. **Constraint Satisfaction**: Russell & Norvig's "AI: A Modern Approach"
3. **Combinatorial Optimization**: Papadimitriou & Steiglitz
4. **Algorithm Design**: Kleinberg & Tardos
