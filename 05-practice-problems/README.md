# DSA Practice Problems 🎯

A comprehensive collection of **70+ practice problems** across multiple categories, designed to reinforce DSA concepts through hands-on coding. Each problem includes multiple solution approaches, comprehensive tests, and detailed complexity analysis.

## 📚 Problem Categories

### 🔤 Arrays & Strings
**Easy Level (10 problems)**
- Two Sum (Hash Map vs Brute Force)
- Valid Palindrome (Two Pointers)
- Remove Duplicates from Sorted Array
- Move Zeroes (In-place modification)
- Contains Duplicate (Hash Set vs Sort)
- Valid Anagram (Frequency vs Sort)
- First Unique Character
- Longest Common Prefix
- Reverse String
- Plus One (Carry propagation)

**Medium Level (10 problems)**
- 3Sum (Sort + Two Pointers)
- Container With Most Water
- Longest Substring Without Repeating Characters
- Group Anagrams
- Longest Palindromic Substring
- Product of Array Except Self
- Rotate Array (Reversal technique)
- Find All Duplicates in Array
- Subarray Sum Equals K
- Spiral Matrix

### 🏗️ Data Structures  
**Easy Level (10 problems)**
- Valid Parentheses (Stack)
- Implement Queue using Stacks
- Binary Tree Inorder Traversal
- Maximum Depth of Binary Tree
- Same Tree
- Symmetric Tree
- Linked List Cycle (Floyd's Algorithm)
- Merge Two Sorted Lists
- Remove Duplicates from Sorted List
- Reverse Linked List

**Medium Level (10 problems)**
- Add Two Numbers (Linked Lists)
- Remove Nth Node From End
- Binary Tree Level Order Traversal
- Validate Binary Search Tree
- Construct Binary Tree from Preorder/Inorder
- Kth Smallest Element in BST
- Lowest Common Ancestor
- Binary Tree Zigzag Level Order Traversal
- Design HashMap
- Implement Trie (Prefix Tree)

### 🌐 Graphs
**Easy Level (10 problems)**
- Find Center of Star Graph
- Find if Path Exists in Graph
- Number of Connected Components (Union-Find)
- Clone Graph
- All Paths From Source to Target
- Flood Fill
- Number of Islands
- Surrounded Regions
- Pacific Atlantic Water Flow
- Course Schedule (Cycle Detection)

### 🌀 Dynamic Programming
**Easy Level (10 problems)**
- Fibonacci (Multiple approaches)
- Climbing Stairs (DP variations)
- House Robber (Linear DP)
- Maximum Subarray (Kadane's Algorithm)
- Best Time to Buy/Sell Stock
- Pascal's Triangle (2D DP)
- Min Cost Climbing Stairs
- Range Sum Query (Prefix sums)
- N-th Tribonacci Number
- Is Subsequence (Two pointers + DP)

**Medium Level (10 problems)**
- Unique Paths (2D DP, Space optimization)
- Unique Paths II (Obstacle handling)
- Minimum Path Sum (Grid DP)
- Coin Change (Unbounded knapsack)
- Longest Increasing Subsequence (Binary search optimization)
- Word Break (String DP)
- House Robber II (Circular arrays)
- Decode Ways (String encoding)
- Maximum Product Subarray (Track min/max)
- Combination Sum IV (Order matters)

### 🔄 Backtracking
**Medium Level (10 problems)**
- Permutations (DFS + Used array)
- Combination Sum (Recursive exploration)
- Subsets (Power set generation)
- N-Queens (Constraint satisfaction)
- Generate Parentheses (Valid combinations)
- Word Search (2D grid traversal)
- Palindrome Partitioning (String decomposition)
- Letter Combinations (Phone number)
- Sudoku Solver (Constraint propagation)
- Restore IP Addresses (String validation)

## 🎯 Problem Difficulty Levels

### 🟢 **Easy (Fundamentals Application)**
- Basic data structure usage
- Simple algorithm application
- Linear time complexity solutions
- Foundation building problems

### 🟡 **Medium (Algorithmic Thinking)**
- Multiple data structure combination
- Optimization techniques required
- O(n log n) complexity solutions
- Interview preparation level

### 🔴 **Hard (Expert Challenges)**
- Advanced algorithm combinations
- Complex optimization requirements
- Competitive programming level
- System design applications

### 🟣 **Expert (Research Level)**
- Novel algorithm applications
- Performance optimization focus
- Real-world system challenges
- Industry-level complexity

## 📈 **Learning Progression**

### **Phase 1: Foundation Practice** (Easy Problems)
Apply individual data structures and algorithms to build confidence and pattern recognition.

### **Phase 2: Integration** (Medium Problems)
Combine multiple concepts to solve more complex challenges requiring algorithmic thinking.

### **Phase 3: Optimization** (Hard Problems)
Focus on performance optimization and advanced technique application.

### **Phase 4: Mastery** (Expert Problems)
Tackle research-level problems and real-world system challenges.

## 🏆 **Success Metrics**

- ✅ **100+ Practice Problems** across all categories and difficulty levels
- ✅ **Complete Solutions** with multiple approaches where applicable
- ✅ **Performance Analysis** with time/space complexity evaluation
- ✅ **Test Cases** covering edge cases and large inputs
- ✅ **Interview Preparation** with explanation and optimization tips

## 📖 **Quick Start**

### Run by Category
```bash
# Array and String problems
go run main.go arrays

# Data structure problems  
go run main.go data-structures

# Graph problems
go run main.go graphs

# Dynamic programming problems
go run main.go dynamic-programming

# Backtracking problems
go run main.go backtracking

# All problems (70+ problems)
go run main.go all
```

### Run by Difficulty Level
```bash
# Easy level problems
go run main.go arrays-easy
go run main.go ds-easy
go run main.go graph-easy
go run main.go dp-easy

# Medium level problems
go run main.go arrays-medium
go run main.go ds-medium
go run main.go dp-medium
go run main.go backtrack-medium
```

## 📁 **Module Structure**

```
05-practice-problems/          # Separate Go module (go.mod: dsa-practice-problems)
├── main.go                    # CLI runner: go run main.go <command>
├── go.mod
├── README.md
├── stacks-and-arrays.md       # Optional learning guide
├── arrays-strings/
│   ├── easy/                  # problems.go, problems_test.go ✅
│   └── medium/                # problems.go (no test file)
├── data-structures/
│   ├── easy/                  # problems.go (no test file)
│   └── medium/                # problems.go (no test file)
├── graphs/
│   └── easy/                  # problems.go (no test file)
├── dynamic-programming/
│   ├── easy/                  # problems.go (no test file)
│   └── medium/                # problems.go (no test file)
└── backtracking/
    └── medium/                # problems.go (no test file)
```

**Important:** This folder is a **separate Go module**. Run all commands from **inside** `05-practice-problems/` (e.g. `cd 05-practice-problems` first).

## How to verify (from 05-practice-problems/)

```bash
cd 05-practice-problems

# Run tests (only arrays-strings/easy has *_test.go)
go test ./... -v

# Run the CLI demo (any category)
go run main.go arrays-easy
go run main.go dp-easy
go run main.go help
```

## Completeness checklist

| Item | Status |
|------|--------|
| Tests | ✅ arrays-strings/easy only; other packages have no `*_test.go` |
| Runner | ✅ `main.go` in this folder (CLI: `go run main.go <command>`) |
| README | ✅ This file; module structure and run commands documented |

---

## 🌟 **Real-World Problem Mapping**

### **Tech Interview Preparation**
- **FAANG Companies**: Google, Amazon, Meta, Apple, Netflix problems
- **Startup Challenges**: Scalability and optimization focus
- **System Design**: Algorithm-heavy component design

### **Competitive Programming**
- **Contest Platforms**: LeetCode, CodeForces, AtCoder style problems
- **Algorithm Competitions**: ICPC, IOI preparation
- **Speed Optimization**: Fast implementation techniques

### **Industry Applications**
- **Database Systems**: Query optimization, indexing
- **Network Engineering**: Routing, load balancing
- **Graphics/Gaming**: Pathfinding, collision detection
- **Financial Systems**: Risk analysis, optimization

## 🎮 **Problem Selection Strategy**

### **Coverage Approach**
- **Pattern Recognition**: Common algorithmic patterns
- **Technique Mastery**: Multiple applications of each concept
- **Complexity Progression**: Gradual difficulty increase
- **Real-World Relevance**: Industry and interview focus

### **Quality Assurance**
- **Multiple Solutions**: Different approaches for comparison
- **Edge Case Testing**: Comprehensive test coverage
- **Performance Benchmarking**: Time/space analysis
- **Code Quality**: Clean, documented, maintainable solutions

---

*Ready to apply your DSA mastery to solve real-world challenges!* 🚀
