// Package backtracking provides comprehensive implementations of backtracking algorithms
// for solving constraint satisfaction problems and combinatorial optimization.
//
// This module covers:
// - Classic Problems: N-Queens, Sudoku solver, Knight's tour
// - Combinatorial: Permutations, combinations, subsets
// - Constraint Satisfaction: Graph coloring, maze solving
// - Optimization: Branch and bound techniques
package backtracking

import (
	"fmt"
	"sort"
	"strings"
)

// =============================================================================
// 1. N-QUEENS PROBLEM
// =============================================================================

// NQueensResult represents a solution to the N-Queens problem
type NQueensResult struct {
	Board     [][]bool // 2D board representation
	Positions []int    // Queen positions (row i, column positions[i])
	Size      int      // Board size
}

// String returns a visual representation of the N-Queens solution
func (nq *NQueensResult) String() string {
	var sb strings.Builder
	for i := 0; i < nq.Size; i++ {
		for j := 0; j < nq.Size; j++ {
			if nq.Board[i][j] {
				sb.WriteString("Q ")
			} else {
				sb.WriteString(". ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// SolveNQueens finds all solutions to the N-Queens problem
// Time: O(N!), Space: O(N²)
func SolveNQueens(n int) []*NQueensResult {
	if n <= 0 {
		return []*NQueensResult{}
	}

	var solutions []*NQueensResult
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}

	positions := make([]int, n)
	solveNQueensHelper(board, positions, 0, n, &solutions)
	return solutions
}

// SolveNQueensOne finds one solution to the N-Queens problem (faster)
func SolveNQueensOne(n int) *NQueensResult {
	if n <= 0 {
		return nil
	}

	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}

	positions := make([]int, n)
	if solveNQueensOneHelper(board, positions, 0, n) {
		// Create a copy for the result
		resultBoard := make([][]bool, n)
		resultPositions := make([]int, n)
		for i := range board {
			resultBoard[i] = make([]bool, n)
			copy(resultBoard[i], board[i])
			resultPositions[i] = positions[i]
		}

		return &NQueensResult{
			Board:     resultBoard,
			Positions: resultPositions,
			Size:      n,
		}
	}

	return nil
}

func solveNQueensHelper(board [][]bool, positions []int, row, n int, solutions *[]*NQueensResult) {
	if row == n {
		// Found a solution, create a copy
		resultBoard := make([][]bool, n)
		resultPositions := make([]int, n)
		for i := range board {
			resultBoard[i] = make([]bool, n)
			copy(resultBoard[i], board[i])
			resultPositions[i] = positions[i]
		}

		*solutions = append(*solutions, &NQueensResult{
			Board:     resultBoard,
			Positions: resultPositions,
			Size:      n,
		})
		return
	}

	for col := 0; col < n; col++ {
		if isSafeQueen(board, row, col, n) {
			// Place queen
			board[row][col] = true
			positions[row] = col

			// Recurse to next row
			solveNQueensHelper(board, positions, row+1, n, solutions)

			// Backtrack
			board[row][col] = false
		}
	}
}

func solveNQueensOneHelper(board [][]bool, positions []int, row, n int) bool {
	if row == n {
		return true // Found one solution
	}

	for col := 0; col < n; col++ {
		if isSafeQueen(board, row, col, n) {
			// Place queen
			board[row][col] = true
			positions[row] = col

			// Recurse to next row
			if solveNQueensOneHelper(board, positions, row+1, n) {
				return true
			}

			// Backtrack
			board[row][col] = false
		}
	}

	return false
}

func isSafeQueen(board [][]bool, row, col, n int) bool {
	// Check column
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}

	// Check diagonal (top-left to bottom-right)
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	// Check diagonal (top-right to bottom-left)
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}

	return true
}

// CountNQueensSolutions counts the number of solutions without storing them
func CountNQueensSolutions(n int) int {
	if n <= 0 {
		return 0
	}

	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}

	return countNQueensHelper(board, 0, n)
}

func countNQueensHelper(board [][]bool, row, n int) int {
	if row == n {
		return 1
	}

	count := 0
	for col := 0; col < n; col++ {
		if isSafeQueen(board, row, col, n) {
			board[row][col] = true
			count += countNQueensHelper(board, row+1, n)
			board[row][col] = false
		}
	}

	return count
}

// =============================================================================
// 2. SUDOKU SOLVER
// =============================================================================

const SudokuSize = 9
const BoxSize = 3

// SudokuBoard represents a 9x9 Sudoku puzzle
type SudokuBoard [SudokuSize][SudokuSize]int

// String returns a visual representation of the Sudoku board
func (sb *SudokuBoard) String() string {
	var result strings.Builder
	for i := 0; i < SudokuSize; i++ {
		if i%3 == 0 && i != 0 {
			result.WriteString("------+-------+------\n")
		}
		for j := 0; j < SudokuSize; j++ {
			if j%3 == 0 && j != 0 {
				result.WriteString("| ")
			}
			if sb[i][j] == 0 {
				result.WriteString(". ")
			} else {
				result.WriteString(fmt.Sprintf("%d ", sb[i][j]))
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

// IsValid checks if the Sudoku board is valid
func (sb *SudokuBoard) IsValid() bool {
	// Check rows
	for i := 0; i < SudokuSize; i++ {
		seen := make(map[int]bool)
		for j := 0; j < SudokuSize; j++ {
			if sb[i][j] != 0 {
				if seen[sb[i][j]] {
					return false
				}
				seen[sb[i][j]] = true
			}
		}
	}

	// Check columns
	for j := 0; j < SudokuSize; j++ {
		seen := make(map[int]bool)
		for i := 0; i < SudokuSize; i++ {
			if sb[i][j] != 0 {
				if seen[sb[i][j]] {
					return false
				}
				seen[sb[i][j]] = true
			}
		}
	}

	// Check 3x3 boxes
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			seen := make(map[int]bool)
			for i := boxRow * 3; i < (boxRow+1)*3; i++ {
				for j := boxCol * 3; j < (boxCol+1)*3; j++ {
					if sb[i][j] != 0 {
						if seen[sb[i][j]] {
							return false
						}
						seen[sb[i][j]] = true
					}
				}
			}
		}
	}

	return true
}

// SolveSudoku solves a Sudoku puzzle using backtracking
// Time: O(9^(n²)), Space: O(n²) where n = 9
func SolveSudoku(board *SudokuBoard) bool {
	row, col, found := findEmptyCell(board)
	if !found {
		return true // No empty cells, puzzle solved
	}

	for num := 1; num <= 9; num++ {
		if isSafeSudoku(board, row, col, num) {
			board[row][col] = num

			if SolveSudoku(board) {
				return true
			}

			// Backtrack
			board[row][col] = 0
		}
	}

	return false
}

func findEmptyCell(board *SudokuBoard) (int, int, bool) {
	for i := 0; i < SudokuSize; i++ {
		for j := 0; j < SudokuSize; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func isSafeSudoku(board *SudokuBoard, row, col, num int) bool {
	// Check row
	for j := 0; j < SudokuSize; j++ {
		if board[row][j] == num {
			return false
		}
	}

	// Check column
	for i := 0; i < SudokuSize; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check 3x3 box
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}

// GenerateSudoku generates a valid Sudoku puzzle with given difficulty
func GenerateSudoku(difficulty int) *SudokuBoard {
	board := &SudokuBoard{}

	// Fill diagonal 3x3 boxes first (they don't affect each other)
	fillDiagonalBoxes(board)

	// Fill remaining cells
	solveSudokuHelper(board, 0, 3)

	// Remove numbers based on difficulty
	removeNumbers(board, difficulty)

	return board
}

func fillDiagonalBoxes(board *SudokuBoard) {
	for box := 0; box < 3; box++ {
		fillBox(board, box*3, box*3)
	}
}

func fillBox(board *SudokuBoard, row, col int) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Simple shuffle
	for i := len(nums) - 1; i > 0; i-- {
		j := i % (i + 1) // Simple deterministic shuffle for demo
		nums[i], nums[j] = nums[j], nums[i]
	}

	idx := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[row+i][col+j] = nums[idx]
			idx++
		}
	}
}

func solveSudokuHelper(board *SudokuBoard, row, col int) bool {
	if row == SudokuSize-1 && col == SudokuSize {
		return true
	}

	if col == SudokuSize {
		row++
		col = 0
	}

	if board[row][col] != 0 {
		return solveSudokuHelper(board, row, col+1)
	}

	for num := 1; num <= 9; num++ {
		if isSafeSudoku(board, row, col, num) {
			board[row][col] = num

			if solveSudokuHelper(board, row, col+1) {
				return true
			}

			board[row][col] = 0
		}
	}

	return false
}

func removeNumbers(board *SudokuBoard, count int) {
	for count > 0 {
		row := count % SudokuSize
		col := (count / SudokuSize) % SudokuSize

		if board[row][col] != 0 {
			board[row][col] = 0
			count--
		} else {
			count--
		}
	}
}

// =============================================================================
// 3. COMBINATORIAL PROBLEMS
// =============================================================================

// GeneratePermutations generates all permutations of the given slice
// Time: O(n! × n), Space: O(n! × n)
func GeneratePermutations[T any](arr []T) [][]T {
	if len(arr) == 0 {
		return [][]T{}
	}

	var result [][]T
	used := make([]bool, len(arr))
	current := make([]T, 0, len(arr))

	generatePermutationsHelper(arr, used, current, &result)
	return result
}

func generatePermutationsHelper[T any](arr []T, used []bool, current []T, result *[][]T) {
	if len(current) == len(arr) {
		// Make a copy of current permutation
		perm := make([]T, len(current))
		copy(perm, current)
		*result = append(*result, perm)
		return
	}

	for i := 0; i < len(arr); i++ {
		if !used[i] {
			used[i] = true
			current = append(current, arr[i])

			generatePermutationsHelper(arr, used, current, result)

			// Backtrack
			current = current[:len(current)-1]
			used[i] = false
		}
	}
}

// GenerateCombinations generates all combinations of size k from the given slice
// Time: O(C(n,k) × k), Space: O(C(n,k) × k)
func GenerateCombinations[T any](arr []T, k int) [][]T {
	if k > len(arr) || k < 0 {
		return [][]T{}
	}

	var result [][]T
	current := make([]T, 0, k)

	generateCombinationsHelper(arr, k, 0, current, &result)
	return result
}

func generateCombinationsHelper[T any](arr []T, k, start int, current []T, result *[][]T) {
	if len(current) == k {
		// Make a copy of current combination
		comb := make([]T, len(current))
		copy(comb, current)
		*result = append(*result, comb)
		return
	}

	for i := start; i < len(arr); i++ {
		current = append(current, arr[i])
		generateCombinationsHelper(arr, k, i+1, current, result)
		// Backtrack
		current = current[:len(current)-1]
	}
}

// GenerateSubsets generates all subsets of the given slice
// Time: O(2^n × n), Space: O(2^n × n)
func GenerateSubsets[T any](arr []T) [][]T {
	var result [][]T
	current := make([]T, 0)

	generateSubsetsHelper(arr, 0, current, &result)
	return result
}

func generateSubsetsHelper[T any](arr []T, index int, current []T, result *[][]T) {
	// Add current subset
	subset := make([]T, len(current))
	copy(subset, current)
	*result = append(*result, subset)

	// Generate subsets with remaining elements
	for i := index; i < len(arr); i++ {
		current = append(current, arr[i])
		generateSubsetsHelper(arr, i+1, current, result)
		// Backtrack
		current = current[:len(current)-1]
	}
}

// =============================================================================
// 4. GRAPH COLORING
// =============================================================================

// GraphColoring represents a graph coloring problem
type GraphColoring struct {
	Graph     [][]bool // Adjacency matrix
	Colors    []int    // Color assignment for each vertex
	NumColors int      // Number of colors available
	Size      int      // Number of vertices
}

// NewGraphColoring creates a new graph coloring problem
func NewGraphColoring(size, numColors int) *GraphColoring {
	return &GraphColoring{
		Graph:     make([][]bool, size),
		Colors:    make([]int, size),
		NumColors: numColors,
		Size:      size,
	}
}

// AddEdge adds an edge between two vertices
func (gc *GraphColoring) AddEdge(u, v int) {
	if u < gc.Size && v < gc.Size {
		if gc.Graph[u] == nil {
			gc.Graph[u] = make([]bool, gc.Size)
		}
		if gc.Graph[v] == nil {
			gc.Graph[v] = make([]bool, gc.Size)
		}
		gc.Graph[u][v] = true
		gc.Graph[v][u] = true
	}
}

// SolveColoring attempts to color the graph with the given number of colors
// Time: O(m^n), Space: O(n) where m = colors, n = vertices
func (gc *GraphColoring) SolveColoring() bool {
	// Initialize all vertices with no color (-1)
	for i := range gc.Colors {
		gc.Colors[i] = -1
	}

	return gc.solveColoringHelper(0)
}

func (gc *GraphColoring) solveColoringHelper(vertex int) bool {
	if vertex == gc.Size {
		return true // All vertices colored
	}

	for color := 0; color < gc.NumColors; color++ {
		if gc.isSafeColor(vertex, color) {
			gc.Colors[vertex] = color

			if gc.solveColoringHelper(vertex + 1) {
				return true
			}

			// Backtrack
			gc.Colors[vertex] = -1
		}
	}

	return false
}

func (gc *GraphColoring) isSafeColor(vertex, color int) bool {
	for i := 0; i < gc.Size; i++ {
		if len(gc.Graph[vertex]) > i && gc.Graph[vertex][i] && gc.Colors[i] == color {
			return false
		}
	}
	return true
}

// GetColoring returns the current color assignment
func (gc *GraphColoring) GetColoring() []int {
	result := make([]int, len(gc.Colors))
	copy(result, gc.Colors)
	return result
}

// =============================================================================
// 5. MAZE SOLVING
// =============================================================================

// Maze represents a maze solving problem
type Maze struct {
	Grid     [][]int  // 0 = path, 1 = wall
	Solution [][]bool // Solution path
	Rows     int
	Cols     int
	StartX   int
	StartY   int
	EndX     int
	EndY     int
}

// NewMaze creates a new maze
func NewMaze(grid [][]int, startX, startY, endX, endY int) *Maze {
	rows := len(grid)
	cols := 0
	if rows > 0 {
		cols = len(grid[0])
	}

	solution := make([][]bool, rows)
	for i := range solution {
		solution[i] = make([]bool, cols)
	}

	return &Maze{
		Grid:     grid,
		Solution: solution,
		Rows:     rows,
		Cols:     cols,
		StartX:   startX,
		StartY:   startY,
		EndX:     endX,
		EndY:     endY,
	}
}

// SolveMaze solves the maze using backtracking
// Time: O(4^(mn)), Space: O(mn)
func (m *Maze) SolveMaze() bool {
	// Reset solution
	for i := range m.Solution {
		for j := range m.Solution[i] {
			m.Solution[i][j] = false
		}
	}

	return m.solveMazeHelper(m.StartX, m.StartY)
}

func (m *Maze) solveMazeHelper(x, y int) bool {
	// Check if we reached the destination
	if x == m.EndX && y == m.EndY {
		m.Solution[x][y] = true
		return true
	}

	// Check if current position is valid
	if m.isSafeMaze(x, y) {
		// Mark current cell as part of solution
		m.Solution[x][y] = true

		// Try all four directions
		directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right, down, left, up

		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if m.solveMazeHelper(newX, newY) {
				return true
			}
		}

		// Backtrack
		m.Solution[x][y] = false
		return false
	}

	return false
}

func (m *Maze) isSafeMaze(x, y int) bool {
	return x >= 0 && x < m.Rows && y >= 0 && y < m.Cols &&
		m.Grid[x][y] == 0 && !m.Solution[x][y]
}

// GetSolution returns the solution path
func (m *Maze) GetSolution() [][]bool {
	result := make([][]bool, m.Rows)
	for i := range result {
		result[i] = make([]bool, m.Cols)
		copy(result[i], m.Solution[i])
	}
	return result
}

// String returns a visual representation of the maze and solution
func (m *Maze) String() string {
	var sb strings.Builder
	sb.WriteString("Maze:\n")
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if i == m.StartX && j == m.StartY {
				sb.WriteString("S ")
			} else if i == m.EndX && j == m.EndY {
				sb.WriteString("E ")
			} else if m.Grid[i][j] == 1 {
				sb.WriteString("# ")
			} else {
				sb.WriteString(". ")
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\nSolution:\n")
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if i == m.StartX && j == m.StartY {
				sb.WriteString("S ")
			} else if i == m.EndX && j == m.EndY {
				sb.WriteString("E ")
			} else if m.Grid[i][j] == 1 {
				sb.WriteString("# ")
			} else if m.Solution[i][j] {
				sb.WriteString("* ")
			} else {
				sb.WriteString(". ")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

// =============================================================================
// 6. KNIGHT'S TOUR PROBLEM
// =============================================================================

// KnightsTour represents a Knight's Tour problem
type KnightsTour struct {
	Board [][]int // -1 = unvisited, 0+ = move number
	Size  int     // Board size (n x n)
}

// NewKnightsTour creates a new Knight's Tour problem
func NewKnightsTour(size int) *KnightsTour {
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
		for j := range board[i] {
			board[i][j] = -1
		}
	}

	return &KnightsTour{
		Board: board,
		Size:  size,
	}
}

// SolveKnightsTour solves the Knight's Tour problem starting from (startX, startY)
// Time: O(8^(n²)), Space: O(n²)
func (kt *KnightsTour) SolveKnightsTour(startX, startY int) bool {
	// Reset board
	for i := range kt.Board {
		for j := range kt.Board[i] {
			kt.Board[i][j] = -1
		}
	}

	// Knight moves: 8 possible L-shaped moves
	moveX := []int{2, 1, -1, -2, -2, -1, 1, 2}
	moveY := []int{1, 2, 2, 1, -1, -2, -2, -1}

	// Mark starting position
	kt.Board[startX][startY] = 0

	return kt.solveKnightsTourHelper(startX, startY, 1, moveX, moveY)
}

func (kt *KnightsTour) solveKnightsTourHelper(x, y, moveCount int, moveX, moveY []int) bool {
	// If all squares are visited
	if moveCount == kt.Size*kt.Size {
		return true
	}

	// Try all 8 possible moves
	for i := 0; i < 8; i++ {
		nextX := x + moveX[i]
		nextY := y + moveY[i]

		if kt.isSafeKnight(nextX, nextY) {
			kt.Board[nextX][nextY] = moveCount

			if kt.solveKnightsTourHelper(nextX, nextY, moveCount+1, moveX, moveY) {
				return true
			}

			// Backtrack
			kt.Board[nextX][nextY] = -1
		}
	}

	return false
}

func (kt *KnightsTour) isSafeKnight(x, y int) bool {
	return x >= 0 && x < kt.Size && y >= 0 && y < kt.Size && kt.Board[x][y] == -1
}

// GetBoard returns a copy of the current board
func (kt *KnightsTour) GetBoard() [][]int {
	result := make([][]int, kt.Size)
	for i := range result {
		result[i] = make([]int, kt.Size)
		copy(result[i], kt.Board[i])
	}
	return result
}

// String returns a visual representation of the Knight's Tour
func (kt *KnightsTour) String() string {
	var sb strings.Builder
	for i := 0; i < kt.Size; i++ {
		for j := 0; j < kt.Size; j++ {
			if kt.Board[i][j] == -1 {
				sb.WriteString("  . ")
			} else {
				sb.WriteString(fmt.Sprintf("%3d ", kt.Board[i][j]))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// =============================================================================
// 7. WORD BREAK PROBLEM (Advanced Backtracking)
// =============================================================================

// WordBreak solves the word break problem using backtracking
// Given a string and a dictionary, determine if the string can be segmented
// into space-separated sequence of dictionary words
func WordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	var result []string
	var current []string

	wordBreakHelper(s, 0, wordSet, current, &result)
	return result
}

func wordBreakHelper(s string, start int, wordSet map[string]bool, current []string, result *[]string) {
	if start == len(s) {
		// Found a valid segmentation
		*result = append(*result, strings.Join(current, " "))
		return
	}

	for end := start + 1; end <= len(s); end++ {
		word := s[start:end]
		if wordSet[word] {
			current = append(current, word)
			wordBreakHelper(s, end, wordSet, current, result)
			// Backtrack
			current = current[:len(current)-1]
		}
	}
}

// =============================================================================
// 8. UTILITY FUNCTIONS
// =============================================================================

// PrintBoard prints a 2D boolean board
func PrintBoard(board [][]bool, trueChar, falseChar string) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				fmt.Print(trueChar + " ")
			} else {
				fmt.Print(falseChar + " ")
			}
		}
		fmt.Println()
	}
}

// PrintIntBoard prints a 2D integer board
func PrintIntBoard(board [][]int) {
	for i := range board {
		for j := range board[i] {
			fmt.Printf("%3d ", board[i][j])
		}
		fmt.Println()
	}
}

// ValidatePermutation checks if a slice is a valid permutation of another
func ValidatePermutation[T comparable](original, permutation []T) bool {
	if len(original) != len(permutation) {
		return false
	}

	originalCount := make(map[T]int)
	for _, item := range original {
		originalCount[item]++
	}

	permCount := make(map[T]int)
	for _, item := range permutation {
		permCount[item]++
	}

	if len(originalCount) != len(permCount) {
		return false
	}

	for item, count := range originalCount {
		if permCount[item] != count {
			return false
		}
	}

	return true
}

// SortResults sorts backtracking results for consistent output
func SortResults[T any](results [][]T, less func(a, b []T) bool) {
	sort.Slice(results, func(i, j int) bool {
		return less(results[i], results[j])
	})
}
