package backtracking

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// =============================================================================
// N-QUEENS TESTS
// =============================================================================

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name               string
		n                  int
		expectedCount      int
		shouldHaveSolution bool
	}{
		{"N=0", 0, 0, false},
		{"N=1", 1, 1, true},
		{"N=2", 2, 0, false},
		{"N=3", 3, 0, false},
		{"N=4", 4, 2, true},
		{"N=5", 5, 10, true},
		{"N=6", 6, 4, true},
		{"N=7", 7, 40, true},
		{"N=8", 8, 92, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solutions := SolveNQueens(tt.n)

			if len(solutions) != tt.expectedCount {
				t.Errorf("SolveNQueens(%d) returned %d solutions, expected %d",
					tt.n, len(solutions), tt.expectedCount)
			}

			// Validate each solution
			for i, solution := range solutions {
				if !isValidNQueensSolution(solution) {
					t.Errorf("Solution %d for N=%d is invalid", i, tt.n)
				}
			}
		})
	}
}

func TestSolveNQueensOne(t *testing.T) {
	tests := []struct {
		name               string
		n                  int
		shouldHaveSolution bool
	}{
		{"N=0", 0, false},
		{"N=1", 1, true},
		{"N=2", 2, false},
		{"N=3", 3, false},
		{"N=4", 4, true},
		{"N=8", 8, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solution := SolveNQueensOne(tt.n)

			if tt.shouldHaveSolution && solution == nil {
				t.Errorf("SolveNQueensOne(%d) should have found a solution", tt.n)
			}

			if !tt.shouldHaveSolution && solution != nil {
				t.Errorf("SolveNQueensOne(%d) should not have found a solution", tt.n)
			}

			if solution != nil && !isValidNQueensSolution(solution) {
				t.Errorf("SolveNQueensOne(%d) returned invalid solution", tt.n)
			}
		})
	}
}

func TestCountNQueensSolutions(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
	}

	for _, tt := range tests {
		t.Run("N="+string(rune(tt.n+'0')), func(t *testing.T) {
			count := CountNQueensSolutions(tt.n)
			if count != tt.expected {
				t.Errorf("CountNQueensSolutions(%d) = %d, expected %d",
					tt.n, count, tt.expected)
			}
		})
	}
}

func isValidNQueensSolution(solution *NQueensResult) bool {
	if solution == nil {
		return false
	}

	n := solution.Size

	// Check that exactly n queens are placed
	queenCount := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if solution.Board[i][j] {
				queenCount++
			}
		}
	}

	if queenCount != n {
		return false
	}

	// Check positions array consistency
	for i := 0; i < n; i++ {
		col := solution.Positions[i]
		if col < 0 || col >= n || !solution.Board[i][col] {
			return false
		}
	}

	// Check no conflicts
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			col1, col2 := solution.Positions[i], solution.Positions[j]

			// Same column
			if col1 == col2 {
				return false
			}

			// Same diagonal
			if abs(i-j) == abs(col1-col2) {
				return false
			}
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// =============================================================================
// SUDOKU TESTS
// =============================================================================

func TestSudokuIsValid(t *testing.T) {
	tests := []struct {
		name     string
		board    SudokuBoard
		expected bool
	}{
		{
			name:     "Empty board",
			board:    SudokuBoard{},
			expected: true,
		},
		{
			name: "Valid partial board",
			board: SudokuBoard{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			expected: true,
		},
		{
			name: "Invalid - duplicate in row",
			board: SudokuBoard{
				{5, 5, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
		{
			name: "Invalid - duplicate in column",
			board: SudokuBoard{
				{5, 0, 0, 0, 0, 0, 0, 0, 0},
				{5, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.board.IsValid()
			if result != tt.expected {
				t.Errorf("IsValid() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSolveSudoku(t *testing.T) {
	tests := []struct {
		name     string
		puzzle   SudokuBoard
		solvable bool
	}{
		{
			name: "Solvable puzzle",
			puzzle: SudokuBoard{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			solvable: true,
		},
		{
			name:     "Empty puzzle",
			puzzle:   SudokuBoard{},
			solvable: true,
		},
		{
			name: "Already solved puzzle",
			puzzle: SudokuBoard{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9},
			},
			solvable: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := tt.puzzle // Copy the board
			solved := SolveSudoku(&board)

			if solved != tt.solvable {
				t.Errorf("SolveSudoku() = %v, expected %v", solved, tt.solvable)
			}

			if solved {
				if !board.IsValid() {
					t.Error("Solved board is not valid")
				}

				// Check if board is completely filled
				for i := 0; i < SudokuSize; i++ {
					for j := 0; j < SudokuSize; j++ {
						if board[i][j] == 0 {
							t.Error("Solved board has empty cells")
						}
					}
				}
			}
		})
	}
}

func TestGenerateSudoku(t *testing.T) {
	difficulties := []int{20, 40, 60}

	for _, difficulty := range difficulties {
		t.Run("Difficulty "+string(rune(difficulty/10+'0')), func(t *testing.T) {
			board := GenerateSudoku(difficulty)

			if !board.IsValid() {
				t.Error("Generated Sudoku is not valid")
			}

			// Count empty cells
			emptyCells := 0
			for i := 0; i < SudokuSize; i++ {
				for j := 0; j < SudokuSize; j++ {
					if board[i][j] == 0 {
						emptyCells++
					}
				}
			}

			// Should have approximately the difficulty number of empty cells
			if emptyCells < difficulty/2 || emptyCells > difficulty*2 {
				t.Errorf("Generated Sudoku has %d empty cells, expected around %d",
					emptyCells, difficulty)
			}
		})
	}
}

// =============================================================================
// COMBINATORIAL TESTS
// =============================================================================

func TestGeneratePermutations(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int // Number of expected permutations
	}{
		{"Empty array", []int{}, 0},
		{"Single element", []int{1}, 1},
		{"Two elements", []int{1, 2}, 2},
		{"Three elements", []int{1, 2, 3}, 6},
		{"Four elements", []int{1, 2, 3, 4}, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GeneratePermutations(tt.input)

			if len(result) != tt.expected {
				t.Errorf("GeneratePermutations(%v) returned %d permutations, expected %d",
					tt.input, len(result), tt.expected)
			}

			// Check that each result is a valid permutation
			for _, perm := range result {
				if !ValidatePermutation(tt.input, perm) {
					t.Errorf("Invalid permutation: %v", perm)
				}
			}

			// Check for duplicates
			seen := make(map[string]bool)
			for _, perm := range result {
				key := strings.Trim(strings.Join(strings.Fields(fmt.Sprintf("%v", perm)), ","), "[]")
				if seen[key] {
					t.Errorf("Duplicate permutation found: %v", perm)
				}
				seen[key] = true
			}
		})
	}
}

func TestGenerateCombinations(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected int
	}{
		{"Empty array", []int{}, 0, 1},
		{"K > N", []int{1, 2}, 3, 0},
		{"K = 0", []int{1, 2, 3}, 0, 1},
		{"K = N", []int{1, 2, 3}, 3, 1},
		{"C(4,2)", []int{1, 2, 3, 4}, 2, 6},
		{"C(5,3)", []int{1, 2, 3, 4, 5}, 3, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateCombinations(tt.input, tt.k)

			if len(result) != tt.expected {
				t.Errorf("GenerateCombinations(%v, %d) returned %d combinations, expected %d",
					tt.input, tt.k, len(result), tt.expected)
			}

			// Check that each combination has correct length
			for _, comb := range result {
				if len(comb) != tt.k {
					t.Errorf("Combination has wrong length: %v (expected length %d)", comb, tt.k)
				}
			}
		})
	}
}

func TestGenerateSubsets(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int // 2^n subsets
	}{
		{"Empty array", []int{}, 1},
		{"Single element", []int{1}, 2},
		{"Two elements", []int{1, 2}, 4},
		{"Three elements", []int{1, 2, 3}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateSubsets(tt.input)

			if len(result) != tt.expected {
				t.Errorf("GenerateSubsets(%v) returned %d subsets, expected %d",
					tt.input, len(result), tt.expected)
			}

			// Should contain empty set
			hasEmpty := false
			for _, subset := range result {
				if len(subset) == 0 {
					hasEmpty = true
					break
				}
			}
			if !hasEmpty {
				t.Error("Result should contain empty subset")
			}
		})
	}
}

// =============================================================================
// GRAPH COLORING TESTS
// =============================================================================

func TestGraphColoring(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		edges    [][2]int
		colors   int
		solvable bool
	}{
		{
			name:     "Triangle graph with 3 colors",
			size:     3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			colors:   3,
			solvable: true,
		},
		{
			name:     "Triangle graph with 2 colors",
			size:     3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			colors:   2,
			solvable: false,
		},
		{
			name:     "Complete graph K4 with 4 colors",
			size:     4,
			edges:    [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}},
			colors:   4,
			solvable: true,
		},
		{
			name:     "Complete graph K4 with 3 colors",
			size:     4,
			edges:    [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}},
			colors:   3,
			solvable: false,
		},
		{
			name:     "Empty graph",
			size:     4,
			edges:    [][2]int{},
			colors:   1,
			solvable: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gc := NewGraphColoring(tt.size, tt.colors)

			// Add edges
			for _, edge := range tt.edges {
				gc.AddEdge(edge[0], edge[1])
			}

			solved := gc.SolveColoring()

			if solved != tt.solvable {
				t.Errorf("SolveColoring() = %v, expected %v", solved, tt.solvable)
			}

			if solved {
				// Validate the coloring
				colors := gc.GetColoring()

				// Check that all vertices are colored
				for i, color := range colors {
					if color < 0 || color >= tt.colors {
						t.Errorf("Vertex %d has invalid color %d", i, color)
					}
				}

				// Check that adjacent vertices have different colors
				for _, edge := range tt.edges {
					u, v := edge[0], edge[1]
					if colors[u] == colors[v] {
						t.Errorf("Adjacent vertices %d and %d have same color", u, v)
					}
				}
			}
		})
	}
}

// =============================================================================
// MAZE SOLVING TESTS
// =============================================================================

func TestMazeSolving(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		startX   int
		startY   int
		endX     int
		endY     int
		solvable bool
	}{
		{
			name: "Simple 3x3 maze",
			grid: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			startX: 0, startY: 0,
			endX: 2, endY: 2,
			solvable: true,
		},
		{
			name: "Blocked maze",
			grid: [][]int{
				{0, 1, 0},
				{1, 1, 0},
				{0, 0, 0},
			},
			startX: 0, startY: 0,
			endX: 2, endY: 2,
			solvable: false,
		},
		{
			name: "Large solvable maze",
			grid: [][]int{
				{0, 0, 1, 0, 0},
				{1, 0, 1, 0, 1},
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
			},
			startX: 0, startY: 0,
			endX: 4, endY: 4,
			solvable: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maze := NewMaze(tt.grid, tt.startX, tt.startY, tt.endX, tt.endY)
			solved := maze.SolveMaze()

			if solved != tt.solvable {
				t.Errorf("SolveMaze() = %v, expected %v", solved, tt.solvable)
			}

			if solved {
				solution := maze.GetSolution()

				// Check that start and end are marked
				if !solution[tt.startX][tt.startY] {
					t.Error("Start position not marked in solution")
				}
				if !solution[tt.endX][tt.endY] {
					t.Error("End position not marked in solution")
				}

				// Check that solution path doesn't go through walls
				for i := 0; i < len(solution); i++ {
					for j := 0; j < len(solution[i]); j++ {
						if solution[i][j] && tt.grid[i][j] == 1 {
							t.Errorf("Solution path goes through wall at (%d, %d)", i, j)
						}
					}
				}
			}
		})
	}
}

// =============================================================================
// KNIGHT'S TOUR TESTS
// =============================================================================

func TestKnightsTour(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		startX   int
		startY   int
		solvable bool
	}{
		{"5x5 from (0,0)", 5, 0, 0, true},
		{"6x6 from (0,0)", 6, 0, 0, true},
		{"3x3 from (0,0)", 3, 0, 0, false},
		{"4x4 from (0,0)", 4, 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kt := NewKnightsTour(tt.size)
			solved := kt.SolveKnightsTour(tt.startX, tt.startY)

			if solved != tt.solvable {
				t.Errorf("SolveKnightsTour() = %v, expected %v", solved, tt.solvable)
			}

			if solved {
				board := kt.GetBoard()

				// Check that all squares are visited exactly once
				visited := make(map[int]bool)
				for i := 0; i < tt.size; i++ {
					for j := 0; j < tt.size; j++ {
						move := board[i][j]
						if move < 0 || move >= tt.size*tt.size {
							t.Errorf("Invalid move number %d at (%d, %d)", move, i, j)
						}
						if visited[move] {
							t.Errorf("Move number %d appears multiple times", move)
						}
						visited[move] = true
					}
				}

				// Check that we have all moves 0 to n²-1
				for i := 0; i < tt.size*tt.size; i++ {
					if !visited[i] {
						t.Errorf("Missing move number %d", i)
					}
				}

				// Check starting position
				if board[tt.startX][tt.startY] != 0 {
					t.Error("Starting position should have move number 0")
				}
			}
		})
	}
}

// =============================================================================
// WORD BREAK TESTS
// =============================================================================

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected []string
	}{
		{
			name:     "Basic case",
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			expected: []string{"cat sand dog", "cats and dog"},
		},
		{
			name:     "Single word",
			s:        "cat",
			wordDict: []string{"cat"},
			expected: []string{"cat"},
		},
		{
			name:     "No solution",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: []string{},
		},
		{
			name:     "Empty string",
			s:        "",
			wordDict: []string{"cat", "dog"},
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordBreak(tt.s, tt.wordDict)

			// For empty results, check length instead of deep equal
			if len(tt.expected) == 0 {
				if len(result) != 0 {
					t.Errorf("WordBreak(%q, %v) = %v, expected empty result",
						tt.s, tt.wordDict, result)
				}
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("WordBreak(%q, %v) = %v, expected %v",
					tt.s, tt.wordDict, result, tt.expected)
			}
		})
	}
}

// =============================================================================
// UTILITY TESTS
// =============================================================================

func TestValidatePermutation(t *testing.T) {
	tests := []struct {
		name        string
		original    []int
		permutation []int
		expected    bool
	}{
		{"Valid permutation", []int{1, 2, 3}, []int{3, 1, 2}, true},
		{"Same array", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"Different length", []int{1, 2, 3}, []int{1, 2}, false},
		{"Different elements", []int{1, 2, 3}, []int{1, 2, 4}, false},
		{"Duplicate elements", []int{1, 2, 3}, []int{1, 1, 2}, false},
		{"Empty arrays", []int{}, []int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidatePermutation(tt.original, tt.permutation)
			if result != tt.expected {
				t.Errorf("ValidatePermutation(%v, %v) = %v, expected %v",
					tt.original, tt.permutation, result, tt.expected)
			}
		})
	}
}

// =============================================================================
// BENCHMARK TESTS
// =============================================================================

func BenchmarkNQueens4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveNQueens(4)
	}
}

func BenchmarkNQueens8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveNQueens(8)
	}
}

func BenchmarkSudokuSolve(b *testing.B) {
	puzzle := SudokuBoard{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	for i := 0; i < b.N; i++ {
		board := puzzle // Copy
		SolveSudoku(&board)
	}
}

func BenchmarkPermutations5(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		GeneratePermutations(arr)
	}
}

func BenchmarkCombinations10_5(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		GenerateCombinations(arr, 5)
	}
}

func BenchmarkKnightsTour5x5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kt := NewKnightsTour(5)
		kt.SolveKnightsTour(0, 0)
	}
}

// =============================================================================
// STRESS TESTS
// =============================================================================

func TestNQueensStress(t *testing.T) {
	// Test larger N values to ensure algorithm handles complexity
	sizes := []int{9, 10}

	for _, n := range sizes {
		t.Run("N="+string(rune(n+'0')), func(t *testing.T) {
			solution := SolveNQueensOne(n)
			if solution == nil {
				t.Errorf("Failed to find solution for N=%d", n)
			} else if !isValidNQueensSolution(solution) {
				t.Errorf("Invalid solution for N=%d", n)
			}
		})
	}
}

func TestLargePermutations(t *testing.T) {
	// Test permutations of larger arrays
	arr := []int{1, 2, 3, 4, 5, 6}
	result := GeneratePermutations(arr)

	expected := 720 // 6!
	if len(result) != expected {
		t.Errorf("Expected %d permutations, got %d", expected, len(result))
	}

	// Verify all permutations are unique
	seen := make(map[string]bool)
	for _, perm := range result {
		key := strings.Trim(strings.Join(strings.Fields(fmt.Sprintf("%v", perm)), ","), "[]")
		if seen[key] {
			t.Error("Found duplicate permutation")
			break
		}
		seen[key] = true
	}
}

func TestComplexMaze(t *testing.T) {
	// Test a complex maze with multiple possible paths
	grid := [][]int{
		{0, 0, 0, 1, 0, 0, 0},
		{1, 1, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}

	maze := NewMaze(grid, 0, 0, 6, 6)
	solved := maze.SolveMaze()

	if !solved {
		t.Error("Failed to solve complex maze")
	}
}
