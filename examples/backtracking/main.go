package main

import (
	"fmt"
	"time"

	"dsa-mastery/03-algorithms/backtracking"
)

func main() {
	fmt.Println("🎯 DSA Mastery - Backtracking Algorithms Demonstration")
	fmt.Println("====================================================")

	// 1. N-Queens Problem Demonstration
	fmt.Println("\n👑 1. N-Queens Problem")
	fmt.Println("----------------------")
	demonstrateNQueens()

	// 2. Sudoku Solver Demonstration
	fmt.Println("\n🧩 2. Sudoku Solver")
	fmt.Println("------------------")
	demonstrateSudoku()

	// 3. Combinatorial Problems
	fmt.Println("\n🔄 3. Combinatorial Generation")
	fmt.Println("-----------------------------")
	demonstratePermutations()
	demonstrateCombinations()
	demonstrateSubsets()

	// 4. Graph Coloring Problem
	fmt.Println("\n🎨 4. Graph Coloring")
	fmt.Println("-------------------")
	demonstrateGraphColoring()

	// 5. Maze Solving
	fmt.Println("\n🗺️  5. Maze Solving")
	fmt.Println("------------------")
	demonstrateMazeSolving()

	// 6. Knight's Tour
	fmt.Println("\n♞ 6. Knight's Tour")
	fmt.Println("-----------------")
	demonstrateKnightsTour()

	// 7. Word Break Problem
	fmt.Println("\n📝 7. Word Break Problem")
	fmt.Println("-----------------------")
	demonstrateWordBreak()

	fmt.Println("\n✅ All backtracking algorithms demonstrated successfully!")
	fmt.Println("🚀 Ready to tackle complex constraint satisfaction problems!")
}

func demonstrateNQueens() {
	fmt.Println("Problem: Place N queens on an N×N chessboard so no two queens attack each other")

	// Solve 4-Queens for visualization
	fmt.Println("\n4-Queens Problem:")
	start := time.Now()
	solutions := backtracking.SolveNQueens(4)
	duration := time.Since(start)

	fmt.Printf("Found %d solutions in %v\n", len(solutions), duration)

	if len(solutions) > 0 {
		fmt.Printf("\nFirst solution:\n%s", solutions[0].String())

		// Show queen positions
		fmt.Printf("Queen positions (row, col): ")
		for i, col := range solutions[0].Positions {
			fmt.Printf("(%d,%d) ", i, col)
		}
		fmt.Println()
	}

	// Performance comparison for different sizes
	fmt.Println("\nPerformance analysis:")
	for n := 4; n <= 8; n++ {
		start := time.Now()
		count := backtracking.CountNQueensSolutions(n)
		duration := time.Since(start)
		fmt.Printf("N=%d: %d solutions in %v\n", n, count, duration)
	}
}

func demonstrateSudoku() {
	fmt.Println("Problem: Fill a 9×9 grid with digits 1-9 following Sudoku rules")

	// Create a sample puzzle
	puzzle := &backtracking.SudokuBoard{
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

	fmt.Printf("\nOriginal puzzle:\n%s", puzzle.String())

	// Solve the puzzle
	start := time.Now()
	solved := backtracking.SolveSudoku(puzzle)
	duration := time.Since(start)

	if solved {
		fmt.Printf("✅ Solved in %v!\n", duration)
		fmt.Printf("\nSolved puzzle:\n%s", puzzle.String())
	} else {
		fmt.Println("❌ No solution found!")
	}

	// Generate a new puzzle
	fmt.Println("Generating a new puzzle...")
	newPuzzle := backtracking.GenerateSudoku(35)
	fmt.Printf("Generated puzzle (35 empty cells):\n%s", newPuzzle.String())
}

func demonstratePermutations() {
	fmt.Println("\nPermutations: All possible arrangements of elements")

	arr := []string{"A", "B", "C"}
	start := time.Now()
	perms := backtracking.GeneratePermutations(arr)
	duration := time.Since(start)

	fmt.Printf("Permutations of %v:\n", arr)
	fmt.Printf("Generated %d permutations in %v\n", len(perms), duration)

	for i, perm := range perms {
		fmt.Printf("%d. %v\n", i+1, perm)
	}

	// Performance test with numbers
	numbers := []int{1, 2, 3, 4, 5}
	start = time.Now()
	numPerms := backtracking.GeneratePermutations(numbers)
	duration = time.Since(start)
	fmt.Printf("\n%d! = %d permutations of %v generated in %v\n",
		len(numbers), len(numPerms), numbers, duration)
}

func demonstrateCombinations() {
	fmt.Println("\nCombinations: Selecting k elements from n elements")

	arr := []int{1, 2, 3, 4, 5}
	k := 3

	start := time.Now()
	combs := backtracking.GenerateCombinations(arr, k)
	duration := time.Since(start)

	fmt.Printf("C(%d,%d) - Choose %d from %v:\n", len(arr), k, k, arr)
	fmt.Printf("Generated %d combinations in %v\n", len(combs), duration)

	for i, comb := range combs {
		fmt.Printf("%d. %v\n", i+1, comb)
	}
}

func demonstrateSubsets() {
	fmt.Println("\nSubsets: All possible subsets (power set)")

	arr := []int{1, 2, 3}
	start := time.Now()
	subsets := backtracking.GenerateSubsets(arr)
	duration := time.Since(start)

	fmt.Printf("Power set of %v:\n", arr)
	fmt.Printf("Generated %d subsets (2^%d) in %v\n", len(subsets), len(arr), duration)

	for i, subset := range subsets {
		if len(subset) == 0 {
			fmt.Printf("%d. ∅ (empty set)\n", i+1)
		} else {
			fmt.Printf("%d. %v\n", i+1, subset)
		}
	}
}

func demonstrateGraphColoring() {
	fmt.Println("Problem: Color graph vertices so adjacent vertices have different colors")

	// Create a simple graph (square with diagonal)
	gc := backtracking.NewGraphColoring(4, 3)
	edges := [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 2}}

	fmt.Println("\nGraph structure:")
	fmt.Println("Vertices: 0, 1, 2, 3")
	fmt.Printf("Edges: %v\n", edges)
	fmt.Println("Attempting to color with 3 colors...")

	for _, edge := range edges {
		gc.AddEdge(edge[0], edge[1])
	}

	start := time.Now()
	solved := gc.SolveColoring()
	duration := time.Since(start)

	if solved {
		colors := gc.GetColoring()
		colorNames := []string{"Red", "Blue", "Green", "Yellow"}

		fmt.Printf("✅ Solution found in %v!\n", duration)
		fmt.Printf("Vertex colors: %v\n", colors)
		fmt.Print("Color assignment: ")
		for i, color := range colors {
			fmt.Printf("V%d=%s ", i, colorNames[color])
		}
		fmt.Println()
	} else {
		fmt.Printf("❌ No solution found in %v\n", duration)
	}
}

func demonstrateMazeSolving() {
	fmt.Println("Problem: Find a path from start to end in a maze")

	// Create a sample maze (0 = path, 1 = wall)
	grid := [][]int{
		{0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	maze := backtracking.NewMaze(grid, 0, 0, 4, 4)

	fmt.Printf("\nMaze layout (S=start, E=end, #=wall, .=path):\n")
	fmt.Print("Original maze:\n")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 0 && j == 0 {
				fmt.Print("S ")
			} else if i == 4 && j == 4 {
				fmt.Print("E ")
			} else if grid[i][j] == 1 {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}

	start := time.Now()
	solved := maze.SolveMaze()
	duration := time.Since(start)

	if solved {
		fmt.Printf("\n✅ Path found in %v!\n", duration)
		fmt.Printf("%s", maze.String())
	} else {
		fmt.Printf("\n❌ No path found in %v\n", duration)
	}
}

func demonstrateKnightsTour() {
	fmt.Println("Problem: Move a knight to visit every square on a chessboard exactly once")

	size := 5
	kt := backtracking.NewKnightsTour(size)

	fmt.Printf("\nSolving Knight's Tour on %dx%d board starting from (0,0)\n", size, size)
	fmt.Println("Knight moves in L-shapes: 2 squares in one direction, 1 square perpendicular")

	start := time.Now()
	solved := kt.SolveKnightsTour(0, 0)
	duration := time.Since(start)

	if solved {
		fmt.Printf("✅ Tour found in %v!\n", duration)
		fmt.Printf("Move sequence (numbers show order of visits):\n%s", kt.String())

		// Verify the solution
		board := kt.GetBoard()
		moves := make(map[int]bool)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				move := board[i][j]
				if moves[move] {
					fmt.Printf("❌ Error: Move %d appears twice!\n", move)
					return
				}
				moves[move] = true
			}
		}
		fmt.Printf("✅ Verification: All %d squares visited exactly once\n", size*size)
	} else {
		fmt.Printf("❌ No tour found in %v\n", duration)
	}
}

func demonstrateWordBreak() {
	fmt.Println("Problem: Segment a string into dictionary words")

	testCases := []struct {
		text     string
		words    []string
		expected int
	}{
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, 2},
		{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}, 3},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, 0},
	}

	for i, tc := range testCases {
		fmt.Printf("\nTest Case %d:\n", i+1)
		fmt.Printf("Text: '%s'\n", tc.text)
		fmt.Printf("Dictionary: %v\n", tc.words)

		start := time.Now()
		segments := backtracking.WordBreak(tc.text, tc.words)
		duration := time.Since(start)

		fmt.Printf("Found %d ways in %v:\n", len(segments), duration)
		if len(segments) > 0 {
			for j, segment := range segments {
				fmt.Printf("  %d. %s\n", j+1, segment)
			}
		} else {
			fmt.Println("  No valid segmentation possible")
		}
	}
}
