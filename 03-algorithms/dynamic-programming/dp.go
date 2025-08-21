package dp

import (
	"fmt"
	"math"
)

// DPResult represents the result of a dynamic programming solution
type DPResult struct {
	Algorithm       string
	Value           int
	Computations    int
	MemoryUsed      int
	Approach        string // "memoization" or "tabulation"
	TimeComplexity  string
	SpaceComplexity string
}

// DPStats tracks performance metrics for DP algorithms
type DPStats struct {
	Computations int
	CacheHits    int
	CacheMisses  int
}

// ============================================================================
// FIBONACCI SEQUENCE - Classic DP Introduction
// ============================================================================

// FibonacciNaive computes nth Fibonacci number using naive recursion
// Time: O(2^n) | Space: O(n) - Exponential time, demonstrates why DP is needed
func FibonacciNaive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciNaive(n-1) + FibonacciNaive(n-2)
}

// FibonacciMemoization computes nth Fibonacci number using memoization (top-down DP)
// Time: O(n) | Space: O(n)
func FibonacciMemoization(n int) int {
	memo := make(map[int]int)
	return fibMemoHelper(n, memo)
}

func fibMemoHelper(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = fibMemoHelper(n-1, memo) + fibMemoHelper(n-2, memo)
	return memo[n]
}

// FibonacciTabulation computes nth Fibonacci number using tabulation (bottom-up DP)
// Time: O(n) | Space: O(n)
func FibonacciTabulation(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// FibonacciOptimized computes nth Fibonacci number with space optimization
// Time: O(n) | Space: O(1)
func FibonacciOptimized(n int) int {
	if n <= 1 {
		return n
	}

	prev2, prev1 := 0, 1

	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev2, prev1 = prev1, current
	}

	return prev1
}

// ============================================================================
// KNAPSACK PROBLEMS - Optimization Problems
// ============================================================================

// Item represents an item for knapsack problems
type Item struct {
	Weight int
	Value  int
	Name   string
}

// Knapsack01Memoization solves 0/1 Knapsack using memoization
// Time: O(n * W) | Space: O(n * W)
func Knapsack01Memoization(items []Item, capacity int) int {
	memo := make(map[string]int)
	return knapsack01MemoHelper(items, capacity, 0, memo)
}

func knapsack01MemoHelper(items []Item, capacity, index int, memo map[string]int) int {
	if index >= len(items) || capacity <= 0 {
		return 0
	}

	key := fmt.Sprintf("%d_%d", index, capacity)
	if val, exists := memo[key]; exists {
		return val
	}

	// Option 1: Don't include current item
	exclude := knapsack01MemoHelper(items, capacity, index+1, memo)

	// Option 2: Include current item (if it fits)
	include := 0
	if items[index].Weight <= capacity {
		include = items[index].Value +
			knapsack01MemoHelper(items, capacity-items[index].Weight, index+1, memo)
	}

	memo[key] = max(include, exclude)
	return memo[key]
}

// Knapsack01Tabulation solves 0/1 Knapsack using tabulation
// Time: O(n * W) | Space: O(n * W)
func Knapsack01Tabulation(items []Item, capacity int) int {
	n := len(items)
	if n == 0 || capacity == 0 {
		return 0
	}

	// dp[i][w] = maximum value with first i items and capacity w
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			// Don't include current item
			dp[i][w] = dp[i-1][w]

			// Include current item if it fits
			if items[i-1].Weight <= w {
				include := items[i-1].Value + dp[i-1][w-items[i-1].Weight]
				dp[i][w] = max(dp[i][w], include)
			}
		}
	}

	return dp[n][capacity]
}

// Knapsack01Optimized solves 0/1 Knapsack with space optimization
// Time: O(n * W) | Space: O(W)
func Knapsack01Optimized(items []Item, capacity int) int {
	if len(items) == 0 || capacity == 0 {
		return 0
	}

	// Use only one array, iterate backwards to avoid using updated values
	dp := make([]int, capacity+1)

	for _, item := range items {
		// Iterate backwards to avoid using updated values in the same iteration
		for w := capacity; w >= item.Weight; w-- {
			dp[w] = max(dp[w], dp[w-item.Weight]+item.Value)
		}
	}

	return dp[capacity]
}

// KnapsackUnbounded solves Unbounded Knapsack (items can be used multiple times)
// Time: O(n * W) | Space: O(W)
func KnapsackUnbounded(items []Item, capacity int) int {
	if len(items) == 0 || capacity == 0 {
		return 0
	}

	dp := make([]int, capacity+1)

	for w := 1; w <= capacity; w++ {
		for _, item := range items {
			if item.Weight <= w {
				dp[w] = max(dp[w], dp[w-item.Weight]+item.Value)
			}
		}
	}

	return dp[capacity]
}

// ============================================================================
// LONGEST COMMON SUBSEQUENCE (LCS) - String DP
// ============================================================================

// LCSMemoization finds length of Longest Common Subsequence using memoization
// Time: O(m * n) | Space: O(m * n)
func LCSMemoization(text1, text2 string) int {
	memo := make(map[string]int)
	return lcsMemoHelper(text1, text2, 0, 0, memo)
}

func lcsMemoHelper(text1, text2 string, i, j int, memo map[string]int) int {
	if i >= len(text1) || j >= len(text2) {
		return 0
	}

	key := fmt.Sprintf("%d_%d", i, j)
	if val, exists := memo[key]; exists {
		return val
	}

	if text1[i] == text2[j] {
		memo[key] = 1 + lcsMemoHelper(text1, text2, i+1, j+1, memo)
	} else {
		option1 := lcsMemoHelper(text1, text2, i+1, j, memo)
		option2 := lcsMemoHelper(text1, text2, i, j+1, memo)
		memo[key] = max(option1, option2)
	}

	return memo[key]
}

// LCSTabulation finds length of LCS using tabulation
// Time: O(m * n) | Space: O(m * n)
func LCSTabulation(text1, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}

	// dp[i][j] = LCS length of text1[0..i-1] and text2[0..j-1]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// LCSOptimized finds length of LCS with space optimization
// Time: O(m * n) | Space: O(min(m, n))
func LCSOptimized(text1, text2 string) int {
	// Ensure text1 is the shorter string for space optimization
	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}

	m, n := len(text1), len(text2)
	if m == 0 {
		return 0
	}

	// Use only two rows instead of full 2D array
	prev := make([]int, m+1)
	curr := make([]int, m+1)

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if text1[i-1] == text2[j-1] {
				curr[i] = prev[i-1] + 1
			} else {
				curr[i] = max(prev[i], curr[i-1])
			}
		}
		prev, curr = curr, prev
	}

	return prev[m]
}

// ============================================================================
// LONGEST INCREASING SUBSEQUENCE (LIS)
// ============================================================================

// LISTabulation finds length of Longest Increasing Subsequence
// Time: O(n²) | Space: O(n)
func LISTabulation(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)

	// Each element forms a subsequence of length 1
	for i := range dp {
		dp[i] = 1
	}

	maxLength := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i])
	}

	return maxLength
}

// LISBinarySearch finds length of LIS using binary search optimization
// Time: O(n log n) | Space: O(n)
func LISBinarySearch(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// tails[i] = smallest ending element of all increasing subsequences of length i+1
	tails := make([]int, 0, len(nums))

	for _, num := range nums {
		// Binary search for the position to insert/replace
		left, right := 0, len(tails)

		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// If num is larger than all elements in tails, append it
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			// Replace the element at position left
			tails[left] = num
		}
	}

	return len(tails)
}

// ============================================================================
// EDIT DISTANCE (Levenshtein Distance)
// ============================================================================

// EditDistanceTabulation computes minimum edit distance between two strings
// Time: O(m * n) | Space: O(m * n)
func EditDistanceTabulation(word1, word2 string) int {
	m, n := len(word1), len(word2)

	// dp[i][j] = edit distance between word1[0..i-1] and word2[0..j-1]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i // i deletions
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // j insertions
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // No operation needed
			} else {
				// Take minimum of insert, delete, replace
				dp[i][j] = 1 + min(
					dp[i-1][j],   // delete
					dp[i][j-1],   // insert
					dp[i-1][j-1], // replace
				)
			}
		}
	}

	return dp[m][n]
}

// EditDistanceOptimized computes edit distance with space optimization
// Time: O(m * n) | Space: O(min(m, n))
func EditDistanceOptimized(word1, word2 string) int {
	// Ensure word1 is the shorter string
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}

	m, n := len(word1), len(word2)

	// Use only two rows
	prev := make([]int, m+1)
	curr := make([]int, m+1)

	// Initialize first row
	for i := 0; i <= m; i++ {
		prev[i] = i
	}

	for j := 1; j <= n; j++ {
		curr[0] = j

		for i := 1; i <= m; i++ {
			if word1[i-1] == word2[j-1] {
				curr[i] = prev[i-1]
			} else {
				curr[i] = 1 + min(prev[i], curr[i-1], prev[i-1])
			}
		}

		prev, curr = curr, prev
	}

	return prev[m]
}

// ============================================================================
// COIN CHANGE PROBLEMS
// ============================================================================

// CoinChangeMinCoins finds minimum number of coins to make target amount
// Time: O(amount * n) | Space: O(amount)
func CoinChangeMinCoins(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// dp[i] = minimum coins needed to make amount i
	dp := make([]int, amount+1)

	// Initialize with impossible value
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1 // Impossible to make the amount
	}

	return dp[amount]
}

// CoinChangeCountWays counts number of ways to make target amount
// Time: O(amount * n) | Space: O(amount)
func CoinChangeCountWays(coins []int, amount int) int {
	// dp[i] = number of ways to make amount i
	dp := make([]int, amount+1)
	dp[0] = 1 // One way to make 0: use no coins

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

// ============================================================================
// MATRIX CHAIN MULTIPLICATION
// ============================================================================

// MatrixChainMultiplication finds minimum scalar multiplications needed
// Time: O(n³) | Space: O(n²)
func MatrixChainMultiplication(dimensions []int) int {
	n := len(dimensions) - 1 // number of matrices
	if n <= 1 {
		return 0
	}

	// dp[i][j] = minimum multiplications for matrices from i to j
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// l is chain length
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			dp[i][j] = math.MaxInt32

			// Try all possible split points
			for k := i; k < j; k++ {
				cost := dp[i][k] + dp[k+1][j] +
					dimensions[i]*dimensions[k+1]*dimensions[j+1]

				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}

	return dp[0][n-1]
}

// ============================================================================
// PALINDROME PROBLEMS
// ============================================================================

// LongestPalindromicSubsequence finds length of longest palindromic subsequence
// Time: O(n²) | Space: O(n²)
func LongestPalindromicSubsequence(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// dp[i][j] = length of LPS in s[i..j]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // single character is palindrome
	}

	// Fill for substrings of length 2 to n
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1

			if s[i] == s[j] {
				if length == 2 {
					dp[i][j] = 2
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

// MinPalindromicPartition finds minimum cuts needed to partition string into palindromes
// Time: O(n²) | Space: O(n²)
func MinPalindromicPartition(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}

	// isPalindrome[i][j] = true if s[i..j] is palindrome
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	// Build palindrome table
	// Single characters are palindromes
	for i := 0; i < n; i++ {
		isPalindrome[i][i] = true
	}

	// Two character palindromes
	for i := 0; i < n-1; i++ {
		isPalindrome[i][i+1] = (s[i] == s[i+1])
	}

	// Longer palindromes
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			isPalindrome[i][j] = (s[i] == s[j]) && isPalindrome[i+1][j-1]
		}
	}

	// dp[i] = minimum cuts needed for s[0..i]
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if isPalindrome[0][i] {
			dp[i] = 0 // No cuts needed
		} else {
			dp[i] = i // Maximum cuts (each character separate)
			for j := 1; j <= i; j++ {
				if isPalindrome[j][i] {
					dp[i] = min(dp[i], dp[j-1]+1)
				}
			}
		}
	}

	return dp[n-1]
}

// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the minimum of multiple integers
func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	minVal := nums[0]
	for _, num := range nums[1:] {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

// CompareApproaches compares memoization vs tabulation for a given problem
func CompareApproaches(problemName string, memoFunc, tabulationFunc func() int) {
	// This is a utility function that can be used to benchmark different approaches
	// Implementation would measure time and space for comparison
}
