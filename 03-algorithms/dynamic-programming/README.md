# Dynamic Programming

This module provides a comprehensive implementation of Dynamic Programming algorithms in Go, designed to master the optimization technique for solving complex problems by breaking them into overlapping subproblems.

## 📋 Table of Contents

- [What is Dynamic Programming?](#what-is-dynamic-programming)
- [Implemented Algorithms](#implemented-algorithms)
- [DP Approaches](#dp-approaches)
- [Performance Analysis](#performance-analysis)
- [Usage Examples](#usage-examples)
- [Problem Categories](#problem-categories)
- [When to Use DP](#when-to-use-dp)
- [Learning Guide](#learning-guide)

## 🧠 What is Dynamic Programming?

Dynamic Programming (DP) is an algorithmic technique for solving optimization problems by breaking them down into simpler subproblems. It is applicable when:

1. **Optimal Substructure**: The problem can be broken down into subproblems whose solutions combine to solve the original problem
2. **Overlapping Subproblems**: The same subproblems are solved multiple times in a naive recursive approach

### Key Characteristics:
- **Eliminates redundant calculations** by storing results
- **Transforms exponential time algorithms** into polynomial time
- **Two main approaches**: Memoization (top-down) and Tabulation (bottom-up)

## 🔧 Implemented Algorithms

### 1. Fibonacci Sequence (📈 Learning Foundation)

Classic introduction to DP demonstrating the dramatic improvement from naive recursion:

```go
// Naive: O(2^n) - Exponential time
result := dp.FibonacciNaive(40) // Very slow!

// Memoization: O(n) time, O(n) space
result := dp.FibonacciMemoization(40) // Fast!

// Tabulation: O(n) time, O(n) space  
result := dp.FibonacciTabulation(40) // Fast!

// Optimized: O(n) time, O(1) space
result := dp.FibonacciOptimized(40) // Fastest!
```

### 2. Knapsack Problems (🎒 Optimization)

#### 0/1 Knapsack Problem
Each item can be taken at most once:

```go
items := []dp.Item{
    {Weight: 10, Value: 60, Name: "Item1"},
    {Weight: 20, Value: 100, Name: "Item2"},
    {Weight: 30, Value: 120, Name: "Item3"},
}

// Three approaches available
maxValue := dp.Knapsack01Memoization(items, 50)  // Top-down
maxValue := dp.Knapsack01Tabulation(items, 50)   // Bottom-up
maxValue := dp.Knapsack01Optimized(items, 50)    // Space-optimized
```

#### Unbounded Knapsack Problem
Items can be used multiple times:

```go
maxValue := dp.KnapsackUnbounded(items, 50)
```

### 3. Longest Common Subsequence (🔤 String DP)

Find the longest subsequence common to two strings:

```go
text1 := "ABCDGH"
text2 := "AEDFHR"

length := dp.LCSMemoization(text1, text2)  // Result: 3 ("ADH")
length := dp.LCSTabulation(text1, text2)   // Same result, different approach
length := dp.LCSOptimized(text1, text2)    // Space-optimized O(min(m,n))
```

### 4. Longest Increasing Subsequence (📊 Array DP)

Find the length of the longest strictly increasing subsequence:

```go
nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

length := dp.LISTabulation(nums)    // O(n²) approach
length := dp.LISBinarySearch(nums)  // O(n log n) optimized approach
// Result: 4 ([2, 3, 7, 18] or [2, 3, 7, 101])
```

### 5. Edit Distance (✏️ String Transformation)

Minimum operations to transform one string into another:

```go
word1 := "horse"
word2 := "ros"

distance := dp.EditDistanceTabulation(word1, word2)  // Result: 3
distance := dp.EditDistanceOptimized(word1, word2)   // Space-optimized
// Operations: replace 'h'→'r', remove 'o', remove 'e'
```

### 6. Coin Change Problems (💰 Combinatorial DP)

#### Minimum Coins
Find minimum number of coins to make a target amount:

```go
coins := []int{1, 3, 4}
amount := 6

minCoins := dp.CoinChangeMinCoins(coins, amount) // Result: 2 (3+3)
```

#### Count Ways
Count number of ways to make a target amount:

```go
coins := []int{1, 2, 5}
amount := 5

ways := dp.CoinChangeCountWays(coins, amount) // Result: 4 ways
// Ways: [5], [2,2,1], [2,1,1,1], [1,1,1,1,1]
```

### 7. Matrix Chain Multiplication (🔢 Optimization DP)

Find minimum scalar multiplications needed to multiply a chain of matrices:

```go
dimensions := []int{40, 20, 30, 10, 30} // 4 matrices
minOps := dp.MatrixChainMultiplication(dimensions) // Optimal parenthesization
```

### 8. Palindrome Problems (🪞 Advanced String DP)

#### Longest Palindromic Subsequence
```go
s := "bbbab"
length := dp.LongestPalindromicSubsequence(s) // Result: 4 ("bbbb")
```

#### Minimum Palindromic Partition
```go
s := "aab"
cuts := dp.MinPalindromicPartition(s) // Result: 1 ("aa|b")
```

## 🎯 DP Approaches

### 1. Memoization (Top-Down)
- **Natural recursive thinking**
- **Cache results** to avoid recomputation
- **Call stack overhead**
- **On-demand computation**

```go
func fibonacci(n int, memo map[int]int) int {
    if n <= 1 { return n }
    if val, exists := memo[n]; exists { return val }
    
    memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
    return memo[n]
}
```

### 2. Tabulation (Bottom-Up)
- **Iterative approach**
- **Fill table systematically**
- **No recursion overhead**
- **Better space locality**

```go
func fibonacci(n int) int {
    if n <= 1 { return n }
    
    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 1
    
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```

### 3. Space Optimization
- **Reduce space complexity**
- **Keep only necessary previous states**
- **Maintains time complexity**

```go
func fibonacci(n int) int {
    if n <= 1 { return n }
    
    prev2, prev1 := 0, 1
    for i := 2; i <= n; i++ {
        current := prev1 + prev2
        prev2, prev1 = prev1, current
    }
    return prev1
}
```

## 📊 Performance Analysis

### Time Complexity Comparison

| Problem | Naive | Memoization | Tabulation | Space Opt |
|---------|-------|-------------|------------|-----------|
| Fibonacci | O(2ⁿ) | O(n) | O(n) | O(n) |
| 0/1 Knapsack | O(2ⁿ) | O(nW) | O(nW) | O(nW) |
| LCS | O(2^(m+n)) | O(mn) | O(mn) | O(mn) |
| LIS | O(2ⁿ) | O(n²) | O(n²) | O(n log n) |
| Edit Distance | O(3^max(m,n)) | O(mn) | O(mn) | O(mn) |
| Coin Change | O(S^n) | O(nS) | O(nS) | O(nS) |

### Space Complexity

| Algorithm | Memoization | Tabulation | Optimized |
|-----------|-------------|------------|-----------|
| Fibonacci | O(n) | O(n) | O(1) |
| Knapsack | O(nW) | O(nW) | O(W) |
| LCS | O(mn) | O(mn) | O(min(m,n)) |
| Edit Distance | O(mn) | O(mn) | O(min(m,n)) |

### Benchmark Results (Apple M3 Pro)

```
Fibonacci (n=40):
- Optimized:   ~12 ns/op   (O(1) space)
- Tabulation:  ~89 ns/op   (O(n) space)  
- Memoization: ~1766 ns/op (O(n) space + recursion overhead)
```

## 🚀 Usage Examples

### Basic Problem Solving

```go
package main

import (
    "fmt"
    "dsa-mastery/03-algorithms/dynamic-programming"
)

func main() {
    // 1. Fibonacci sequence
    n := 10
    fmt.Printf("Fibonacci(%d) = %d\n", n, dp.FibonacciOptimized(n))
    
    // 2. Knapsack problem
    items := []dp.Item{
        {Weight: 2, Value: 3, Name: "Item A"},
        {Weight: 3, Value: 4, Name: "Item B"},
        {Weight: 4, Value: 5, Name: "Item C"},
    }
    capacity := 5
    maxValue := dp.Knapsack01Optimized(items, capacity)
    fmt.Printf("Max knapsack value: %d\n", maxValue)
    
    // 3. String problems
    text1, text2 := "AGGTAB", "GXTXAYB"
    lcsLength := dp.LCSOptimized(text1, text2)
    fmt.Printf("LCS length: %d\n", lcsLength)
    
    // 4. Array problems
    nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
    lisLength := dp.LISBinarySearch(nums)
    fmt.Printf("LIS length: %d\n", lisLength)
}
```

### Performance Comparison

```go
package main

import (
    "fmt"
    "time"
    "dsa-mastery/03-algorithms/dynamic-programming"
)

func main() {
    n := 35
    
    // Compare Fibonacci approaches
    start := time.Now()
    result1 := dp.FibonacciMemoization(n)
    time1 := time.Since(start)
    
    start = time.Now()
    result2 := dp.FibonacciTabulation(n)
    time2 := time.Since(start)
    
    start = time.Now()
    result3 := dp.FibonacciOptimized(n)
    time3 := time.Since(start)
    
    fmt.Printf("Fibonacci(%d) results:\n", n)
    fmt.Printf("Memoization: %d (Time: %v)\n", result1, time1)
    fmt.Printf("Tabulation:  %d (Time: %v)\n", result2, time2) 
    fmt.Printf("Optimized:   %d (Time: %v)\n", result3, time3)
}
```

## 📚 Problem Categories

### 1. Linear DP (1D)
- **Fibonacci Sequence**: Classic introduction
- **Climbing Stairs**: Count ways to reach top
- **House Robber**: Maximum sum with constraints
- **Coin Change**: Optimization and counting

### 2. Grid DP (2D)
- **Unique Paths**: Count paths in grid
- **Minimum Path Sum**: Find cheapest path
- **Knapsack Problems**: Weight/value optimization
- **Edit Distance**: String transformation

### 3. String DP
- **Longest Common Subsequence**: String similarity
- **Palindrome Problems**: Pattern recognition
- **Pattern Matching**: Advanced string algorithms
- **Word Break**: Segmentation problems

### 4. Tree DP
- **Binary Tree Problems**: Maximum path, diameter
- **Subtree Optimization**: House robber in trees
- **Tree Matching**: Pattern matching in trees

### 5. Interval DP
- **Matrix Chain Multiplication**: Optimal parenthesization
- **Palindromic Partitioning**: Substring optimization
- **Optimal BST**: Search tree construction

### 6. Bitmask DP
- **Traveling Salesman**: State compression
- **Assignment Problems**: Optimal matching
- **Subset Problems**: Complex combinations

## 🎯 When to Use DP

### ✅ Use DP When:
1. **Optimal Substructure** exists
2. **Overlapping Subproblems** are present
3. You need **optimal solutions** (min/max)
4. **Counting problems** with constraints
5. **Decision problems** with choices

### ❌ Don't Use DP When:
1. **Greedy approach** works (simpler)
2. **No overlapping subproblems**
3. **Memory is severely constrained**
4. **Simple iterative solution** exists

### 🔍 DP Identification Patterns:
- Words: "optimal", "maximum", "minimum", "count ways"
- Multiple choices at each step
- Recursive structure with repeated calculations
- Previous decisions affect future options

## 📖 Learning Guide

### Step 1: Master the Basics
1. **Understand the concept**: Optimal substructure + overlapping subproblems
2. **Start with Fibonacci**: Compare naive vs DP approaches
3. **Practice memoization**: Top-down thinking
4. **Learn tabulation**: Bottom-up filling

### Step 2: Classic Problems
1. **Climbing Stairs**: Simple 1D DP
2. **Coin Change**: Multiple approaches
3. **0/1 Knapsack**: 2D DP fundamentals
4. **Longest Common Subsequence**: String DP

### Step 3: Advanced Techniques
1. **Space optimization**: Reduce memory usage
2. **State compression**: Bitmask DP
3. **Multi-dimensional DP**: Complex state spaces
4. **DP on trees**: Non-linear structures

### Step 4: Problem Recognition
1. **Identify DP patterns** in new problems
2. **Choose the right approach** (memoization vs tabulation)
3. **Optimize space complexity** where possible
4. **Handle edge cases** properly

### 🎯 Practice Progression:
1. **Easy**: Fibonacci, Climbing Stairs, Min Cost Climbing
2. **Medium**: Coin Change, 0/1 Knapsack, LCS, Edit Distance
3. **Hard**: Matrix Chain, Palindrome Partitioning, LIS variations
4. **Expert**: TSP, Bitmask DP, Advanced tree DP

## 🔧 Utility Functions

The module provides several utility functions for DP algorithm analysis:

- **Performance comparison** between approaches
- **Memory usage tracking** for different techniques
- **Visualization helpers** for understanding DP tables
- **Problem generators** for testing and learning

## 🎉 Key Takeaways

1. **DP is optimization**: Trade space for time complexity
2. **Two approaches**: Choose based on problem structure and constraints
3. **Pattern recognition**: Most DP problems follow common patterns
4. **Space optimization**: Often possible and beneficial
5. **Practice essential**: Recognition comes with experience

Dynamic Programming is one of the most powerful algorithmic techniques, enabling efficient solutions to complex optimization problems that would otherwise be intractable. Master these patterns and you'll be well-equipped to tackle any DP challenge!
