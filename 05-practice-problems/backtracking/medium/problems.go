package medium

import (
	"fmt"
	"strconv"
	"strings"
)

// ====================================================================
// PROBLEM 1: Permutations
// Generate all possible permutations of a given array.
// ====================================================================

// Permute generates all permutations using backtracking
func Permute(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		// Base case
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// Try each unused number
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			// Choose
			path = append(path, nums[i])
			used[i] = true

			// Explore
			backtrack()

			// Unchoose
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}

// PermuteSwap generates permutations using swap approach
func PermuteSwap(nums []int) [][]int {
	result := [][]int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		// Base case
		if start == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			result = append(result, temp)
			return
		}

		// Try each position from start
		for i := start; i < len(nums); i++ {
			// Swap
			nums[start], nums[i] = nums[i], nums[start]

			// Recurse
			backtrack(start + 1)

			// Backtrack
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	backtrack(0)
	return result
}

// ====================================================================
// PROBLEM 2: Combination Sum
// Find all unique combinations where candidates sum to target.
// ====================================================================

// CombinationSum finds combinations that sum to target
func CombinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(index, currentSum int)
	backtrack = func(index, currentSum int) {
		// Base cases
		if currentSum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		if currentSum > target || index >= len(candidates) {
			return
		}

		// Include current candidate (can reuse)
		path = append(path, candidates[index])
		backtrack(index, currentSum+candidates[index])
		path = path[:len(path)-1]

		// Skip current candidate
		backtrack(index+1, currentSum)
	}

	backtrack(0, 0)
	return result
}

// ====================================================================
// PROBLEM 3: Subsets
// Generate all possible subsets (power set).
// ====================================================================

// Subsets generates all subsets using backtracking
func Subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(index int)
	backtrack = func(index int) {
		// Add current subset
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		// Try including each remaining number
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}

// SubsetsBitMask generates subsets using bit manipulation
func SubsetsBitMask(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}

	// Generate all 2^n subsets
	for mask := 0; mask < (1 << n); mask++ {
		subset := []int{}

		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}

		result = append(result, subset)
	}

	return result
}

// ====================================================================
// PROBLEM 4: N-Queens
// Place N queens on NxN chessboard so none attack each other.
// ====================================================================

// SolveNQueens finds all solutions to N-Queens problem
func SolveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			// Convert board to string format
			solution := make([]string, n)
			for i := 0; i < n; i++ {
				solution[i] = string(board[i])
			}
			result = append(result, solution)
			return
		}

		for col := 0; col < n; col++ {
			if isValidQueenPlacement(board, row, col, n) {
				board[row][col] = 'Q'
				backtrack(row + 1)
				board[row][col] = '.'
			}
		}
	}

	backtrack(0)
	return result
}

func isValidQueenPlacement(board [][]byte, row, col, n int) bool {
	// Check column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check diagonal (top-left to bottom-right)
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check diagonal (top-right to bottom-left)
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

// ====================================================================
// PROBLEM 5: Generate Parentheses
// Generate all combinations of n pairs of well-formed parentheses.
// ====================================================================

// GenerateParenthesis generates valid parentheses combinations
func GenerateParenthesis(n int) []string {
	result := []string{}

	var backtrack func(current string, open, close int)
	backtrack = func(current string, open, close int) {
		// Base case
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		// Add opening parenthesis
		if open < n {
			backtrack(current+"(", open+1, close)
		}

		// Add closing parenthesis
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}

// ====================================================================
// PROBLEM 6: Word Search
// Find if word exists in 2D grid of characters.
// ====================================================================

// WordSearch finds if word exists in grid using backtracking
func WordSearch(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])

	var backtrack func(row, col, index int) bool
	backtrack = func(row, col, index int) bool {
		// Base cases
		if index == len(word) {
			return true
		}

		if row < 0 || row >= m || col < 0 || col >= n ||
			board[row][col] != word[index] {
			return false
		}

		// Mark as visited
		temp := board[row][col]
		board[row][col] = '#'

		// Explore all 4 directions
		found := backtrack(row+1, col, index+1) ||
			backtrack(row-1, col, index+1) ||
			backtrack(row, col+1, index+1) ||
			backtrack(row, col-1, index+1)

		// Restore
		board[row][col] = temp

		return found
	}

	// Try starting from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// ====================================================================
// PROBLEM 7: Palindrome Partitioning
// Partition string such that every substring is a palindrome.
// ====================================================================

// Partition finds all palindrome partitions
func Partition(s string) [][]string {
	result := [][]string{}
	path := []string{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for end := start; end < len(s); end++ {
			if isPalindrome(s, start, end) {
				path = append(path, s[start:end+1])
				backtrack(end + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return result
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// ====================================================================
// PROBLEM 8: Letter Combinations of Phone Number
// Generate all possible letter combinations for given phone number.
// ====================================================================

// LetterCombinations generates phone number letter combinations
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := []string{}

	var backtrack func(index int, current string)
	backtrack = func(index int, current string) {
		if index == len(digits) {
			result = append(result, current)
			return
		}

		letters := phoneMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, current+string(letters[i]))
		}
	}

	backtrack(0, "")
	return result
}

// ====================================================================
// PROBLEM 9: Sudoku Solver
// Solve a Sudoku puzzle using backtracking.
// ====================================================================

// SolveSudoku solves Sudoku puzzle in-place
func SolveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for c := byte('1'); c <= '9'; c++ {
					if isValidSudoku(board, i, j, c) {
						board[i][j] = c

						if solve(board) {
							return true
						}

						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		// Check row
		if board[i][col] == c {
			return false
		}

		// Check column
		if board[row][i] == c {
			return false
		}

		// Check 3x3 box
		boxRow := 3*(row/3) + i/3
		boxCol := 3*(col/3) + i%3
		if board[boxRow][boxCol] == c {
			return false
		}
	}
	return true
}

// ====================================================================
// PROBLEM 10: Restore IP Addresses
// Given string of digits, return all possible valid IP addresses.
// ====================================================================

// RestoreIPAddresses generates all valid IP addresses
func RestoreIPAddresses(s string) []string {
	result := []string{}
	if len(s) < 4 || len(s) > 12 {
		return result
	}

	path := []string{}

	var backtrack func(start int)
	backtrack = func(start int) {
		// Base case
		if len(path) == 4 {
			if start == len(s) {
				result = append(result, strings.Join(path, "."))
			}
			return
		}

		// Try segments of length 1, 2, 3
		for length := 1; length <= 3 && start+length <= len(s); length++ {
			segment := s[start : start+length]

			if isValidIPSegment(segment) {
				path = append(path, segment)
				backtrack(start + length)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return result
}

func isValidIPSegment(segment string) bool {
	// Empty or too long
	if len(segment) == 0 || len(segment) > 3 {
		return false
	}

	// Leading zeros (except single "0")
	if len(segment) > 1 && segment[0] == '0' {
		return false
	}

	// Convert to integer and check range
	num, err := strconv.Atoi(segment)
	if err != nil {
		return false
	}

	return num >= 0 && num <= 255
}

// ====================================================================
// HELPER FUNCTIONS FOR DEMONSTRATION
// ====================================================================

// PrintBoard prints 2D byte board
func PrintBoard(board [][]byte) {
	for _, row := range board {
		fmt.Printf("%s\n", string(row))
	}
}

// PrintStringSlices prints slice of string slices
func PrintStringSlices(slices [][]string) {
	for i, slice := range slices {
		fmt.Printf("%d: %v\n", i+1, slice)
	}
}

// PrintIntSlices prints slice of int slices
func PrintIntSlices(slices [][]int) {
	for i, slice := range slices {
		fmt.Printf("%d: %v\n", i+1, slice)
	}
}

// ====================================================================
// PROBLEM DEMONSTRATION
// ====================================================================

// DemonstrateBacktracking shows all backtracking problems
func DemonstrateBacktracking() {
	fmt.Println("🎯 Medium Backtracking Problems")
	fmt.Println("===============================")

	// Problem 1: Permutations
	fmt.Println("\n1. Permutations:")
	nums := []int{1, 2, 3}
	perms := Permute(nums)
	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("Permutations (%d total):\n", len(perms))
	PrintIntSlices(perms[:6]) // Show first 6
	if len(perms) > 6 {
		fmt.Printf("... and %d more\n", len(perms)-6)
	}

	// Problem 2: Combination Sum
	fmt.Println("\n2. Combination Sum:")
	candidates := []int{2, 3, 6, 7}
	target := 7
	combinations := CombinationSum(candidates, target)
	fmt.Printf("Candidates: %v, Target: %d\n", candidates, target)
	fmt.Printf("Combinations:\n")
	PrintIntSlices(combinations)

	// Problem 3: Subsets
	fmt.Println("\n3. Subsets:")
	subsetNums := []int{1, 2, 3}
	subsets := Subsets(subsetNums)
	fmt.Printf("Array: %v\n", subsetNums)
	fmt.Printf("All subsets (%d total):\n", len(subsets))
	PrintIntSlices(subsets)

	// Problem 4: N-Queens (4x4)
	fmt.Println("\n4. N-Queens (4x4):")
	queens := SolveNQueens(4)
	fmt.Printf("Solutions for 4-Queens (%d total):\n", len(queens))
	if len(queens) > 0 {
		fmt.Println("First solution:")
		for _, row := range queens[0] {
			fmt.Println(row)
		}
	}

	// Problem 5: Generate Parentheses
	fmt.Println("\n5. Generate Parentheses:")
	n := 3
	parentheses := GenerateParenthesis(n)
	fmt.Printf("n = %d, Valid combinations:\n", n)
	for i, p := range parentheses {
		fmt.Printf("%d: %s\n", i+1, p)
	}

	// Problem 6: Word Search
	fmt.Println("\n6. Word Search:")
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Printf("Board:\n")
	PrintBoard(board)
	found := WordSearch(board, word)
	fmt.Printf("Word '%s' found: %t\n", word, found)

	// Problem 7: Palindrome Partitioning
	fmt.Println("\n7. Palindrome Partitioning:")
	s := "aab"
	partitions := Partition(s)
	fmt.Printf("String: '%s'\n", s)
	fmt.Printf("Palindrome partitions:\n")
	PrintStringSlices(partitions)

	// Problem 8: Letter Combinations
	fmt.Println("\n8. Letter Combinations:")
	digits := "23"
	letters := LetterCombinations(digits)
	fmt.Printf("Digits: '%s'\n", digits)
	fmt.Printf("Letter combinations: %v\n", letters)

	// Problem 9: Sudoku Solver (show setup)
	fmt.Println("\n9. Sudoku Solver:")
	sudokuBoard := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Printf("Initial Sudoku puzzle:\n")
	PrintBoard(sudokuBoard)
	fmt.Println("Solving... (solution would be shown)")

	// Problem 10: Restore IP Addresses
	fmt.Println("\n10. Restore IP Addresses:")
	ipString := "25525511135"
	ips := RestoreIPAddresses(ipString)
	fmt.Printf("String: '%s'\n", ipString)
	fmt.Printf("Valid IP addresses:\n")
	for i, ip := range ips {
		fmt.Printf("%d: %s\n", i+1, ip)
	}
}

// ProblemComplexityAnalysis provides complexity analysis for all problems
func ProblemComplexityAnalysis() {
	fmt.Println("\n📊 Backtracking Problems Complexity Analysis")
	fmt.Println("============================================")

	problems := []struct {
		name     string
		timeOpt  string
		spaceOpt string
		approach string
	}{
		{"Permutations", "O(n×n!)", "O(n)", "DFS + Used Array/Swap"},
		{"Combination Sum", "O(2^(t/m))", "O(t/m)", "DFS + Pruning"},
		{"Subsets", "O(2^n)", "O(2^n)", "DFS/Bit Manipulation"},
		{"N-Queens", "O(n!)", "O(n)", "Backtrack + Validation"},
		{"Generate Parentheses", "O(4^n/√n)", "O(4^n/√n)", "Catalan Number"},
		{"Word Search", "O(m×n×4^k)", "O(k)", "DFS + Mark Visited"},
		{"Palindrome Partitioning", "O(2^n)", "O(n)", "DFS + Palindrome Check"},
		{"Letter Combinations", "O(4^n)", "O(4^n)", "DFS Phone Mapping"},
		{"Sudoku Solver", "O(9^(n²))", "O(1)", "Constraint Satisfaction"},
		{"Restore IP Addresses", "O(3^4)", "O(1)", "DFS + Validation"},
	}

	fmt.Printf("%-25s %-15s %-15s %-25s\n", "Problem", "Time", "Space", "Key Technique")
	fmt.Println(strings.Repeat("-", 85))

	for _, p := range problems {
		fmt.Printf("%-25s %-15s %-15s %-25s\n", p.name, p.timeOpt, p.spaceOpt, p.approach)
	}
}
