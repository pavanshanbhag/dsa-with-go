# Backtracking Algorithms Examples

This directory contains comprehensive examples demonstrating all backtracking algorithms implemented in the DSA Mastery project.

## 🚀 Running the Examples

From the project root directory:

```bash
go run examples/backtracking/main.go
```

Or navigate to the examples directory:

```bash
cd examples/backtracking
go run main.go
```

## 📋 Examples Included

### 1. **N-Queens Problem** 👑
- Demonstrates solving the classic N-Queens constraint satisfaction problem
- Shows performance analysis for different board sizes (N=4 to N=8)
- Displays visual board representation with queen placements
- Compares solution counting vs full solution generation

**Key Learning Points:**
- Constraint satisfaction with conflict detection
- Backtracking with early pruning
- Performance scaling with problem size

### 2. **Sudoku Solver** 🧩
- Solves a real Sudoku puzzle step by step
- Demonstrates puzzle generation with configurable difficulty
- Shows visual representation of before/after states
- Validates solution correctness

**Key Learning Points:**
- Complex constraint satisfaction (rows, columns, 3x3 boxes)
- Efficient constraint checking
- Real-world application of backtracking

### 3. **Combinatorial Generation** 🔄
- **Permutations**: All arrangements of elements (3! = 6 for [A,B,C])
- **Combinations**: Selecting k elements from n (C(5,3) = 10)
- **Subsets**: Power set generation (2³ = 8 subsets)

**Key Learning Points:**
- Recursive enumeration patterns
- Mathematical combinatorics implementation
- Performance comparison of different generation types

### 4. **Graph Coloring** 🎨
- Colors a graph with minimum colors needed
- Demonstrates constraint satisfaction on graph structures
- Shows practical application in scheduling problems
- Validates solution by checking adjacent vertex colors

**Key Learning Points:**
- Graph-based constraint satisfaction
- Conflict resolution in optimization
- Real-world applications (register allocation, scheduling)

### 5. **Maze Solving** 🗺️
- Finds path from start to end position in a maze
- Visual representation of maze and solution path
- Demonstrates path-finding algorithms
- Shows backtracking when hitting dead ends

**Key Learning Points:**
- Path-finding algorithms
- Spatial reasoning and navigation
- Dead-end handling and backtracking

### 6. **Knight's Tour** ♞
- Solves the classic chess knight tour problem
- Visits every square on a 5×5 board exactly once
- Shows the complete move sequence
- Verifies solution correctness

**Key Learning Points:**
- Complex constraint satisfaction with movement rules
- Sequence optimization
- Chess algorithm patterns

### 7. **Word Break Problem** 📝
- Segments strings into dictionary words
- Multiple test cases including impossible cases
- Shows all possible valid segmentations
- Demonstrates string-based backtracking

**Key Learning Points:**
- String processing with backtracking
- Dictionary-based constraint satisfaction
- Multiple solution enumeration

## 🎯 Performance Insights

The examples include timing measurements showing:

- **N-Queens**: Exponential growth with board size (270μs for 8-Queens)
- **Sudoku**: Efficient solving (~572μs for complex puzzles)
- **Permutations**: Fast generation (9μs for 5! = 120 permutations)
- **Knight's Tour**: Optimized backtracking (~216μs for 5×5 board)

## 🧠 Learning Objectives

After running these examples, you should understand:

1. **Backtracking Pattern**: The general template for backtracking algorithms
2. **Constraint Satisfaction**: How to model and solve CSP problems
3. **Performance Characteristics**: Time complexity and scaling behavior
4. **Real-world Applications**: Where these algorithms are used in practice
5. **Optimization Techniques**: Pruning, constraint propagation, and heuristics

## 🔧 Customization

Feel free to modify the examples to experiment with:
- Different problem sizes (larger N-Queens, bigger mazes)
- Additional constraints (modified rules, new problems)
- Performance optimizations (different heuristics, pruning strategies)
- Visualization improvements (better output formatting)

## 📚 Further Reading

- **Algorithm Design Manual** by Steven Skiena
- **Introduction to Algorithms** by CLRS
- **Constraint Satisfaction Problems** by Rina Dechter
- **Backtracking Algorithms** in competitive programming resources
