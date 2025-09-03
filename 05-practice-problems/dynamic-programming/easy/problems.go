package easy

import (
	"fmt"
	"strings"
)

// ====================================================================
// PROBLEM 1: Fibonacci Number
// Calculate the nth Fibonacci number using different approaches.
// ====================================================================

// FibRecursive calculates Fibonacci using naive recursion - O(2^n) time
func FibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// FibMemoization calculates Fibonacci using memoization - O(n) time, O(n) space
func FibMemoization(n int) int {
	memo := make(map[int]int)
	return fibMemo(n, memo)
}

func fibMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}

// FibBottomUp calculates Fibonacci using bottom-up DP - O(n) time, O(n) space
func FibBottomUp(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// FibOptimized calculates Fibonacci with space optimization - O(n) time, O(1) space
func FibOptimized(n int) int {
	if n <= 1 {
		return n
	}

	prev2, prev1 := 0, 1

	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// ====================================================================
// PROBLEM 2: Climbing Stairs
// You're climbing a staircase with n steps. You can climb 1 or 2 steps at a time.
// ====================================================================

// ClimbStairs calculates number of ways to climb stairs using DP
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// ClimbStairsOptimized with space optimization
func ClimbStairsOptimized(n int) int {
	if n <= 2 {
		return n
	}

	prev2, prev1 := 1, 2

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// ====================================================================
// PROBLEM 3: Min Cost Climbing Stairs
// You can start from step 0 or 1. Find minimum cost to reach the top.
// ====================================================================

// MinCostClimbingStairs finds minimum cost to reach top
func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n <= 2 {
		return min(cost...)
	}

	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	return min(dp[n-1], dp[n-2])
}

// MinCostClimbingStairsOptimized with space optimization
func MinCostClimbingStairsOptimized(cost []int) int {
	n := len(cost)
	if n <= 2 {
		return min(cost...)
	}

	prev2, prev1 := cost[0], cost[1]

	for i := 2; i < n; i++ {
		current := cost[i] + min(prev1, prev2)
		prev2 = prev1
		prev1 = current
	}

	return min(prev1, prev2)
}

func min(nums ...int) int {
	result := nums[0]
	for _, num := range nums[1:] {
		if num < result {
			result = num
		}
	}
	return result
}

// ====================================================================
// PROBLEM 4: House Robber
// Rob houses in a line, cannot rob adjacent houses.
// ====================================================================

// Rob calculates maximum money that can be robbed
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// RobOptimized with space optimization
func RobOptimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev2, prev1 := nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		current := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ====================================================================
// PROBLEM 5: Maximum Subarray (Kadane's Algorithm)
// Find the contiguous subarray with the largest sum.
// ====================================================================

// MaxSubArray finds maximum sum of contiguous subarray
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSoFar := nums[0]
	maxEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

// MaxSubArrayDP using explicit DP approach
func MaxSubArrayDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := dp[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

// ====================================================================
// PROBLEM 6: Best Time to Buy and Sell Stock
// Buy on one day and sell on another day to maximize profit.
// ====================================================================

// MaxProfit finds maximum profit from stock trading
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			maxProfit = max(maxProfit, profit)
		}
	}

	return maxProfit
}

// MaxProfitDP using DP approach
func MaxProfitDP(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// dp[i][0] = max profit holding no stock on day i
	// dp[i][1] = max profit holding stock on day i
	hold := -prices[0] // Bought stock on first day
	notHold := 0       // No stock on first day

	for i := 1; i < len(prices); i++ {
		newNotHold := max(notHold, hold+prices[i]) // Sell today or keep not holding
		newHold := max(hold, -prices[i])           // Buy today or keep holding

		notHold = newNotHold
		hold = newHold
	}

	return notHold
}

// ====================================================================
// PROBLEM 7: Range Sum Query - Immutable
// Design a data structure for range sum queries.
// ====================================================================

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	return NumArray{prefixSum: prefixSum}
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.prefixSum[right+1] - na.prefixSum[left]
}

// ====================================================================
// PROBLEM 8: Counting Bits
// Count number of 1's in binary representation for numbers 0 to n.
// ====================================================================

// CountBits counts bits using DP
func CountBits(n int) []int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		// dp[i] = dp[i >> 1] + (i & 1)
		// Number of bits in i = number of bits in i/2 + last bit of i
		dp[i] = dp[i>>1] + (i & 1)
	}

	return dp
}

// CountBitsBruteForce using built-in bit counting
func CountBitsBruteForce(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		count := 0
		num := i
		for num > 0 {
			count += num & 1
			num >>= 1
		}
		result[i] = count
	}

	return result
}

// ====================================================================
// PROBLEM 9: Is Subsequence
// Check if s is a subsequence of t.
// ====================================================================

// IsSubsequence checks if s is subsequence of t using two pointers
func IsSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}

// IsSubsequenceDP using DP for multiple queries optimization
func IsSubsequenceDP(s string, t string) bool {
	m, n := len(s), len(t)
	if m == 0 {
		return true
	}
	if n == 0 {
		return false
	}

	// dp[i][j] = true if s[0:i] is subsequence of t[0:j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Empty string is subsequence of any string
	for j := 0; j <= n; j++ {
		dp[0][j] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}

// ====================================================================
// PROBLEM 10: Pascal's Triangle
// Generate the first numRows of Pascal's triangle.
// ====================================================================

// Generate creates Pascal's triangle
func Generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1 // First element is always 1
		triangle[i][i] = 1 // Last element is always 1

		// Fill middle elements
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

// GetRow gets specific row of Pascal's triangle with space optimization
func GetRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		// Build from right to left to avoid overwriting needed values
		for j := i; j > 0; j-- {
			row[j] = row[j] + row[j-1]
		}
	}

	return row
}

// ====================================================================
// HELPER FUNCTIONS FOR DEMONSTRATION
// ====================================================================

// PrintMatrix prints 2D integer matrix
func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
}

// ====================================================================
// PROBLEM DEMONSTRATION
// ====================================================================

// DemonstrateEasyDP shows all easy DP problems
func DemonstrateEasyDP() {
	fmt.Println("🎯 Easy Dynamic Programming Problems")
	fmt.Println("===================================")

	// Problem 1: Fibonacci
	fmt.Println("\n1. Fibonacci Number:")
	n := 10
	fmt.Printf("Fibonacci(%d):\n", n)
	fmt.Printf("Recursive: %d\n", FibRecursive(n))
	fmt.Printf("Memoization: %d\n", FibMemoization(n))
	fmt.Printf("Bottom-up: %d\n", FibBottomUp(n))
	fmt.Printf("Optimized: %d\n", FibOptimized(n))

	// Problem 2: Climbing Stairs
	fmt.Println("\n2. Climbing Stairs:")
	stairs := 5
	ways := ClimbStairs(stairs)
	waysOpt := ClimbStairsOptimized(stairs)
	fmt.Printf("Ways to climb %d stairs: %d (DP), %d (Optimized)\n", stairs, ways, waysOpt)

	// Problem 3: Min Cost Climbing Stairs
	fmt.Println("\n3. Min Cost Climbing Stairs:")
	cost := []int{10, 15, 20}
	minCost := MinCostClimbingStairs(cost)
	minCostOpt := MinCostClimbingStairsOptimized(cost)
	fmt.Printf("Cost array: %v\n", cost)
	fmt.Printf("Min cost: %d (DP), %d (Optimized)\n", minCost, minCostOpt)

	// Problem 4: House Robber
	fmt.Println("\n4. House Robber:")
	houses := []int{2, 7, 9, 3, 1}
	maxMoney := Rob(houses)
	maxMoneyOpt := RobOptimized(houses)
	fmt.Printf("House values: %v\n", houses)
	fmt.Printf("Max money: %d (DP), %d (Optimized)\n", maxMoney, maxMoneyOpt)

	// Problem 5: Maximum Subarray
	fmt.Println("\n5. Maximum Subarray:")
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := MaxSubArray(nums)
	maxSumDP := MaxSubArrayDP(nums)
	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("Max subarray sum: %d (Kadane), %d (DP)\n", maxSum, maxSumDP)

	// Problem 6: Best Time to Buy and Sell Stock
	fmt.Println("\n6. Best Time to Buy and Sell Stock:")
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := MaxProfit(prices)
	profitDP := MaxProfitDP(prices)
	fmt.Printf("Prices: %v\n", prices)
	fmt.Printf("Max profit: %d (Greedy), %d (DP)\n", profit, profitDP)

	// Problem 7: Range Sum Query
	fmt.Println("\n7. Range Sum Query:")
	queryNums := []int{-2, 0, 3, -5, 2, -1}
	numArray := Constructor(queryNums)
	fmt.Printf("Array: %v\n", queryNums)
	fmt.Printf("Sum of range [0, 2]: %d\n", numArray.SumRange(0, 2))
	fmt.Printf("Sum of range [2, 5]: %d\n", numArray.SumRange(2, 5))
	fmt.Printf("Sum of range [0, 5]: %d\n", numArray.SumRange(0, 5))

	// Problem 8: Counting Bits
	fmt.Println("\n8. Counting Bits:")
	bitN := 5
	bits := CountBits(bitN)
	bitsBrute := CountBitsBruteForce(bitN)
	fmt.Printf("Count bits for 0 to %d:\n", bitN)
	fmt.Printf("DP result: %v\n", bits)
	fmt.Printf("Brute force: %v\n", bitsBrute)

	// Problem 9: Is Subsequence
	fmt.Println("\n9. Is Subsequence:")
	s, t := "abc", "aebdc"
	isSubseq := IsSubsequence(s, t)
	isSubseqDP := IsSubsequenceDP(s, t)
	fmt.Printf("'%s' is subsequence of '%s': %t (Two Pointers), %t (DP)\n", s, t, isSubseq, isSubseqDP)

	// Problem 10: Pascal's Triangle
	fmt.Println("\n10. Pascal's Triangle:")
	triangleRows := 5
	triangle := Generate(triangleRows)
	fmt.Printf("Pascal's triangle with %d rows:\n", triangleRows)
	PrintMatrix(triangle)

	specificRow := GetRow(4)
	fmt.Printf("Row 4 specifically: %v\n", specificRow)
}

// ProblemComplexityAnalysis provides complexity analysis for all problems
func ProblemComplexityAnalysis() {
	fmt.Println("\n📊 Easy DP Problems Complexity Analysis")
	fmt.Println("=======================================")

	problems := []struct {
		name     string
		timeOpt  string
		spaceOpt string
		approach string
	}{
		{"Fibonacci", "O(n)", "O(1)", "Bottom-up + Space Opt"},
		{"Climbing Stairs", "O(n)", "O(1)", "DP + Space Opt"},
		{"Min Cost Climbing", "O(n)", "O(1)", "DP + Space Opt"},
		{"House Robber", "O(n)", "O(1)", "DP + Space Opt"},
		{"Maximum Subarray", "O(n)", "O(1)", "Kadane's Algorithm"},
		{"Best Time Stock", "O(n)", "O(1)", "Greedy/DP"},
		{"Range Sum Query", "O(1)", "O(n)", "Prefix Sum"},
		{"Counting Bits", "O(n)", "O(n)", "DP Bit Manipulation"},
		{"Is Subsequence", "O(n)", "O(1)", "Two Pointers"},
		{"Pascal's Triangle", "O(n²)", "O(n²)", "DP Triangle Build"},
	}

	fmt.Printf("%-20s %-10s %-10s %-25s\n", "Problem", "Time", "Space", "Best Approach")
	fmt.Println(strings.Repeat("-", 70))

	for _, p := range problems {
		fmt.Printf("%-20s %-10s %-10s %-25s\n", p.name, p.timeOpt, p.spaceOpt, p.approach)
	}
}
