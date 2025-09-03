package medium

import (
	"fmt"
	"strings"
)

// ====================================================================
// PROBLEM 1: Unique Paths
// Find number of possible unique paths from top-left to bottom-right.
// ====================================================================

// UniquePaths calculates unique paths using 2D DP
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Initialize first row and column
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// Fill DP table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// UniquePathsOptimized using 1D DP for space optimization
func UniquePathsOptimized(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]
}

// UniquePathsMath using combinatorics formula
func UniquePathsMath(m int, n int) int {
	// Total moves needed: (m-1) down + (n-1) right = (m+n-2)
	// Choose (m-1) positions for down moves: C(m+n-2, m-1)
	totalMoves := m + n - 2
	downMoves := m - 1

	result := 1
	for i := 0; i < downMoves; i++ {
		result = result * (totalMoves - i) / (i + 1)
	}

	return result
}

// ====================================================================
// PROBLEM 2: Unique Paths II
// Find unique paths with obstacles in the grid.
// ====================================================================

// UniquePathsWithObstacles calculates paths avoiding obstacles
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Initialize starting point
	dp[0][0] = 1

	// Fill first row
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}

	// Fill first column
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	// Fill rest of the grid
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// ====================================================================
// PROBLEM 3: Minimum Path Sum
// Find path from top-left to bottom-right with minimum sum.
// ====================================================================

// MinPathSum finds minimum path sum using DP
func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Initialize starting point
	dp[0][0] = grid[0][0]

	// Fill first row
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// Fill first column
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// Fill rest of the grid
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}

// MinPathSumOptimized using in-place modification
func MinPathSumOptimized(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// Fill first row
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// Fill first column
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// Fill rest of the grid
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ====================================================================
// PROBLEM 4: Coin Change
// Find minimum coins needed to make amount using given coin denominations.
// ====================================================================

// CoinChange finds minimum coins using DP
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	// Initialize with impossible value
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0 // Base case

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// ====================================================================
// PROBLEM 5: Longest Increasing Subsequence
// Find length of longest increasing subsequence.
// ====================================================================

// LengthOfLIS finds LIS length using DP - O(n²) time
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLength := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i])
	}

	return maxLength
}

// LengthOfLISBinarySearch using binary search - O(n log n) time
func LengthOfLISBinarySearch(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tails := []int{}

	for _, num := range nums {
		// Binary search for position to insert/replace
		left, right := 0, len(tails)

		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == len(tails) {
			tails = append(tails, num)
		} else {
			tails[left] = num
		}
	}

	return len(tails)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ====================================================================
// PROBLEM 6: Word Break
// Determine if string can be segmented into words from dictionary.
// ====================================================================

// WordBreak checks if string can be broken into dictionary words
func WordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// ====================================================================
// PROBLEM 7: House Robber II
// Rob houses in a circle, cannot rob adjacent houses.
// ====================================================================

// RobCircular handles circular arrangement of houses
func RobCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// Case 1: Rob houses 0 to n-2 (include first, exclude last)
	case1 := robLinear(nums[:len(nums)-1])

	// Case 2: Rob houses 1 to n-1 (exclude first, include last)
	case2 := robLinear(nums[1:])

	return max(case1, case2)
}

func robLinear(nums []int) int {
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

// ====================================================================
// PROBLEM 8: Decode Ways
// Count ways to decode a string of digits.
// ====================================================================

// NumDecodings counts decoding ways using DP
func NumDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(s); i++ {
		// Single digit
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		// Two digits
		twoDigit := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

// NumDecodingsOptimized with space optimization
func NumDecodingsOptimized(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	prev2, prev1 := 1, 1

	for i := 2; i <= len(s); i++ {
		current := 0

		// Single digit
		if s[i-1] != '0' {
			current += prev1
		}

		// Two digits
		twoDigit := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			current += prev2
		}

		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// ====================================================================
// PROBLEM 9: Maximum Product Subarray
// Find contiguous subarray with largest product.
// ====================================================================

// MaxProduct finds maximum product subarray
func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSoFar := nums[0]
	minSoFar := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			// Swap max and min when we encounter negative number
			maxSoFar, minSoFar = minSoFar, maxSoFar
		}

		maxSoFar = max(nums[i], maxSoFar*nums[i])
		minSoFar = min(nums[i], minSoFar*nums[i])

		result = max(result, maxSoFar)
	}

	return result
}

// ====================================================================
// PROBLEM 10: Combination Sum IV
// Find number of possible combinations that add up to target.
// ====================================================================

// CombinationSum4 counts combinations using DP
func CombinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
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

// DemonstrateMediumDP shows all medium DP problems
func DemonstrateMediumDP() {
	fmt.Println("🎯 Medium Dynamic Programming Problems")
	fmt.Println("=====================================")

	// Problem 1: Unique Paths
	fmt.Println("\n1. Unique Paths:")
	m, n := 3, 7
	paths := UniquePaths(m, n)
	pathsOpt := UniquePathsOptimized(m, n)
	pathsMath := UniquePathsMath(m, n)
	fmt.Printf("Grid %dx%d unique paths: %d (2D DP), %d (1D DP), %d (Math)\n", m, n, paths, pathsOpt, pathsMath)

	// Problem 2: Unique Paths with Obstacles
	fmt.Println("\n2. Unique Paths II:")
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Printf("Grid with obstacles:\n")
	PrintMatrix(obstacleGrid)
	pathsWithObstacles := UniquePathsWithObstacles(obstacleGrid)
	fmt.Printf("Unique paths avoiding obstacles: %d\n", pathsWithObstacles)

	// Problem 3: Minimum Path Sum
	fmt.Println("\n3. Minimum Path Sum:")
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Printf("Grid:\n")
	PrintMatrix(grid)

	// Make copy for optimized version
	gridCopy := make([][]int, len(grid))
	for i := range grid {
		gridCopy[i] = make([]int, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}

	minSum := MinPathSum(grid)
	fmt.Printf("Minimum path sum: %d\n", minSum)

	// Problem 4: Coin Change
	fmt.Println("\n4. Coin Change:")
	coins := []int{1, 3, 4}
	amount := 6
	minCoins := CoinChange(coins, amount)
	fmt.Printf("Coins: %v, Amount: %d\n", coins, amount)
	fmt.Printf("Minimum coins needed: %d\n", minCoins)

	// Problem 5: Longest Increasing Subsequence
	fmt.Println("\n5. Longest Increasing Subsequence:")
	lisNums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	lisLen := LengthOfLIS(lisNums)
	lisLenBS := LengthOfLISBinarySearch(lisNums)
	fmt.Printf("Array: %v\n", lisNums)
	fmt.Printf("LIS length: %d (DP), %d (Binary Search)\n", lisLen, lisLenBS)

	// Problem 6: Word Break
	fmt.Println("\n6. Word Break:")
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	canBreak := WordBreak(s, wordDict)
	fmt.Printf("String: '%s', Dictionary: %v\n", s, wordDict)
	fmt.Printf("Can break: %t\n", canBreak)

	// Problem 7: House Robber II
	fmt.Println("\n7. House Robber II (Circular):")
	circularHouses := []int{2, 3, 2}
	maxMoneyCircular := RobCircular(circularHouses)
	fmt.Printf("Circular houses: %v\n", circularHouses)
	fmt.Printf("Max money (circular): %d\n", maxMoneyCircular)

	// Problem 8: Decode Ways
	fmt.Println("\n8. Decode Ways:")
	decodeStr := "226"
	ways := NumDecodings(decodeStr)
	waysOpt := NumDecodingsOptimized(decodeStr)
	fmt.Printf("String: '%s'\n", decodeStr)
	fmt.Printf("Decode ways: %d (DP), %d (Optimized)\n", ways, waysOpt)

	// Problem 9: Maximum Product Subarray
	fmt.Println("\n9. Maximum Product Subarray:")
	productNums := []int{2, 3, -2, 4}
	maxProd := MaxProduct(productNums)
	fmt.Printf("Array: %v\n", productNums)
	fmt.Printf("Maximum product subarray: %d\n", maxProd)

	// Problem 10: Combination Sum IV
	fmt.Println("\n10. Combination Sum IV:")
	combNums := []int{1, 2, 3}
	combTarget := 4
	combinations := CombinationSum4(combNums, combTarget)
	fmt.Printf("Numbers: %v, Target: %d\n", combNums, combTarget)
	fmt.Printf("Number of combinations: %d\n", combinations)
}

// ProblemComplexityAnalysis provides complexity analysis for all problems
func ProblemComplexityAnalysis() {
	fmt.Println("\n📊 Medium DP Problems Complexity Analysis")
	fmt.Println("=========================================")

	problems := []struct {
		name     string
		timeOpt  string
		spaceOpt string
		approach string
	}{
		{"Unique Paths", "O(m×n)", "O(n)", "1D DP/Math"},
		{"Unique Paths II", "O(m×n)", "O(m×n)", "2D DP with Obstacles"},
		{"Minimum Path Sum", "O(m×n)", "O(1)", "In-place DP"},
		{"Coin Change", "O(n×m)", "O(n)", "Bottom-up DP"},
		{"Longest Increasing Subsequence", "O(n log n)", "O(n)", "Binary Search + DP"},
		{"Word Break", "O(n³)", "O(n)", "DP with Substring"},
		{"House Robber II", "O(n)", "O(1)", "Linear DP (2 cases)"},
		{"Decode Ways", "O(n)", "O(1)", "DP + Space Opt"},
		{"Maximum Product Subarray", "O(n)", "O(1)", "Track Max/Min"},
		{"Combination Sum IV", "O(n×m)", "O(n)", "Bottom-up DP"},
	}

	fmt.Printf("%-25s %-12s %-12s %-25s\n", "Problem", "Time", "Space", "Best Approach")
	fmt.Println(strings.Repeat("-", 80))

	for _, p := range problems {
		fmt.Printf("%-25s %-12s %-12s %-25s\n", p.name, p.timeOpt, p.spaceOpt, p.approach)
	}
}
