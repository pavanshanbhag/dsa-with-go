package easy

import (
	"fmt"
	"strings"
)

// ====================================================================
// PROBLEM 1: Find Center of Star Graph
// Find the center node of a star graph.
// ====================================================================

// FindCenter finds the center node in a star graph
func FindCenter(edges [][]int) int {
	// In a star graph, the center node appears in every edge
	// So it must appear in the first two edges
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}

// ====================================================================
// PROBLEM 2: Find if Path Exists in Graph
// Given edges and two nodes, determine if path exists between them.
// ====================================================================

// ValidPath checks if path exists using DFS
func ValidPath(n int, edges [][]int, source int, destination int) bool {
	// Build adjacency list
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	return dfsPath(graph, source, destination, visited)
}

func dfsPath(graph [][]int, current, destination int, visited []bool) bool {
	if current == destination {
		return true
	}

	visited[current] = true

	for _, neighbor := range graph[current] {
		if !visited[neighbor] && dfsPath(graph, neighbor, destination, visited) {
			return true
		}
	}

	return false
}

// ValidPathBFS checks if path exists using BFS
func ValidPathBFS(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	// Build adjacency list
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[current] {
			if neighbor == destination {
				return true
			}

			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}

// ====================================================================
// PROBLEM 3: Number of Connected Components (Union-Find)
// Find number of connected components in undirected graph.
// ====================================================================

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}

	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// Union by rank
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
		uf.count--
	}
}

func (uf *UnionFind) GetCount() int {
	return uf.count
}

// CountComponents counts connected components using Union-Find
func CountComponents(n int, edges [][]int) int {
	uf := NewUnionFind(n)

	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}

	return uf.GetCount()
}

// CountComponentsDFS counts connected components using DFS
func CountComponentsDFS(n int, edges [][]int) int {
	// Build adjacency list
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfsComponent(graph, i, visited)
			count++
		}
	}

	return count
}

func dfsComponent(graph [][]int, node int, visited []bool) {
	visited[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfsComponent(graph, neighbor, visited)
		}
	}
}

// ====================================================================
// PROBLEM 4: Clone Graph
// Given a reference of a node in connected undirected graph, return deep copy.
// ====================================================================

type Node struct {
	Val       int
	Neighbors []*Node
}

// CloneGraph creates deep copy of graph using DFS
func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	cloneMap := make(map[*Node]*Node)
	return dfsClone(node, cloneMap)
}

func dfsClone(node *Node, cloneMap map[*Node]*Node) *Node {
	if clonedNode, exists := cloneMap[node]; exists {
		return clonedNode
	}

	// Create clone of current node
	clone := &Node{Val: node.Val}
	cloneMap[node] = clone

	// Clone all neighbors
	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfsClone(neighbor, cloneMap))
	}

	return clone
}

// CloneGraphBFS creates deep copy using BFS
func CloneGraphBFS(node *Node) *Node {
	if node == nil {
		return nil
	}

	cloneMap := make(map[*Node]*Node)
	queue := []*Node{node}
	cloneMap[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range current.Neighbors {
			if _, exists := cloneMap[neighbor]; !exists {
				cloneMap[neighbor] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}

			cloneMap[current].Neighbors = append(cloneMap[current].Neighbors, cloneMap[neighbor])
		}
	}

	return cloneMap[node]
}

// ====================================================================
// PROBLEM 5: All Paths From Source to Target
// Find all paths from source to target in directed acyclic graph.
// ====================================================================

// AllPathsSourceTarget finds all paths using DFS
func AllPathsSourceTarget(graph [][]int) [][]int {
	result := [][]int{}
	path := []int{0}
	target := len(graph) - 1

	dfsAllPaths(graph, 0, target, path, &result)
	return result
}

func dfsAllPaths(graph [][]int, current, target int, path []int, result *[][]int) {
	if current == target {
		// Make copy of current path
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}

	for _, neighbor := range graph[current] {
		path = append(path, neighbor)
		dfsAllPaths(graph, neighbor, target, path, result)
		path = path[:len(path)-1] // Backtrack
	}
}

// ====================================================================
// PROBLEM 6: Flood Fill
// Perform flood fill algorithm starting from given pixel.
// ====================================================================

// FloodFill performs flood fill using DFS
func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}

	originalColor := image[sr][sc]
	if originalColor == color {
		return image // No change needed
	}

	dfsFloodFill(image, sr, sc, originalColor, color)
	return image
}

func dfsFloodFill(image [][]int, row, col, originalColor, newColor int) {
	if row < 0 || row >= len(image) || col < 0 || col >= len(image[0]) {
		return
	}

	if image[row][col] != originalColor {
		return
	}

	image[row][col] = newColor

	// Recursively fill 4-connected neighbors
	dfsFloodFill(image, row+1, col, originalColor, newColor)
	dfsFloodFill(image, row-1, col, originalColor, newColor)
	dfsFloodFill(image, row, col+1, originalColor, newColor)
	dfsFloodFill(image, row, col-1, originalColor, newColor)
}

// FloodFillBFS performs flood fill using BFS
func FloodFillBFS(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}

	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}

	rows, cols := len(image), len(image[0])
	queue := [][]int{{sr, sc}}
	image[sr][sc] = color

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		row, col := current[0], current[1]

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols &&
				image[newRow][newCol] == originalColor {
				image[newRow][newCol] = color
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}

	return image
}

// ====================================================================
// PROBLEM 7: Number of Islands
// Count number of islands in 2D binary grid.
// ====================================================================

// NumIslands counts islands using DFS
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfsIsland(grid, i, j)
				count++
			}
		}
	}

	return count
}

func dfsIsland(grid [][]byte, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != '1' {
		return
	}

	grid[row][col] = '0' // Mark as visited

	// Visit all 4-connected neighbors
	dfsIsland(grid, row+1, col)
	dfsIsland(grid, row-1, col)
	dfsIsland(grid, row, col+1)
	dfsIsland(grid, row, col-1)
}

// NumIslandsBFS counts islands using BFS
func NumIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				// BFS to mark entire island
				queue := [][]int{{i, j}}
				grid[i][j] = '0'

				for len(queue) > 0 {
					current := queue[0]
					queue = queue[1:]
					row, col := current[0], current[1]

					for _, dir := range directions {
						newRow, newCol := row+dir[0], col+dir[1]

						if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols &&
							grid[newRow][newCol] == '1' {
							grid[newRow][newCol] = '0'
							queue = append(queue, []int{newRow, newCol})
						}
					}
				}

				count++
			}
		}
	}

	return count
}

// ====================================================================
// PROBLEM 8: Surrounded Regions
// Capture all regions surrounded by 'X'.
// ====================================================================

// Solve captures surrounded regions
func Solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	// Mark boundary-connected 'O's as safe
	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfsMarkSafe(board, i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfsMarkSafe(board, i, cols-1)
		}
	}

	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			dfsMarkSafe(board, 0, j)
		}
		if board[rows-1][j] == 'O' {
			dfsMarkSafe(board, rows-1, j)
		}
	}

	// Convert remaining 'O's to 'X' and restore safe 'O's
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'S' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfsMarkSafe(board [][]byte, row, col int) {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] != 'O' {
		return
	}

	board[row][col] = 'S' // Mark as safe

	dfsMarkSafe(board, row+1, col)
	dfsMarkSafe(board, row-1, col)
	dfsMarkSafe(board, row, col+1)
	dfsMarkSafe(board, row, col-1)
}

// ====================================================================
// PROBLEM 9: Pacific Atlantic Water Flow
// Find cells where water can flow to both Pacific and Atlantic oceans.
// ====================================================================

// PacificAtlantic finds cells that can reach both oceans
func PacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	rows, cols := len(heights), len(heights[0])
	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	// DFS from Pacific edges (top and left)
	for i := 0; i < rows; i++ {
		dfsPacificAtlantic(heights, i, 0, heights[i][0], pacific)
		dfsPacificAtlantic(heights, i, cols-1, heights[i][cols-1], atlantic)
	}

	for j := 0; j < cols; j++ {
		dfsPacificAtlantic(heights, 0, j, heights[0][j], pacific)
		dfsPacificAtlantic(heights, rows-1, j, heights[rows-1][j], atlantic)
	}

	// Find intersection
	result := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func dfsPacificAtlantic(heights [][]int, row, col, prevHeight int, ocean [][]bool) {
	if row < 0 || row >= len(heights) || col < 0 || col >= len(heights[0]) ||
		ocean[row][col] || heights[row][col] < prevHeight {
		return
	}

	ocean[row][col] = true

	dfsPacificAtlantic(heights, row+1, col, heights[row][col], ocean)
	dfsPacificAtlantic(heights, row-1, col, heights[row][col], ocean)
	dfsPacificAtlantic(heights, row, col+1, heights[row][col], ocean)
	dfsPacificAtlantic(heights, row, col-1, heights[row][col], ocean)
}

// ====================================================================
// PROBLEM 10: Course Schedule
// Determine if you can finish all courses given prerequisites.
// ====================================================================

// CanFinish checks if course schedule is possible using DFS cycle detection
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// Build adjacency list
	graph := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		graph[prereq[1]] = append(graph[prereq[1]], prereq[0])
	}

	// 0: unvisited, 1: visiting, 2: visited
	state := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		if state[i] == 0 && !dfsCanFinish(graph, i, state) {
			return false
		}
	}

	return true
}

func dfsCanFinish(graph [][]int, course int, state []int) bool {
	if state[course] == 1 {
		return false // Cycle detected
	}

	if state[course] == 2 {
		return true // Already processed
	}

	state[course] = 1 // Mark as visiting

	for _, next := range graph[course] {
		if !dfsCanFinish(graph, next, state) {
			return false
		}
	}

	state[course] = 2 // Mark as visited
	return true
}

// ====================================================================
// HELPER FUNCTIONS FOR DEMONSTRATION
// ====================================================================

// PrintMatrix prints 2D matrix for visualization
func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
}

// PrintByteMatrix prints 2D byte matrix
func PrintByteMatrix(matrix [][]byte) {
	for _, row := range matrix {
		fmt.Printf("%s\n", string(row))
	}
}

// ====================================================================
// PROBLEM DEMONSTRATION
// ====================================================================

// DemonstrateEasyGraphs shows all easy graph problems
func DemonstrateEasyGraphs() {
	fmt.Println("🎯 Easy Graph Problems")
	fmt.Println("======================")

	// Problem 1: Find Center of Star Graph
	fmt.Println("\n1. Find Center of Star Graph:")
	edges1 := [][]int{{1, 2}, {2, 3}, {4, 2}}
	center := FindCenter(edges1)
	fmt.Printf("Edges: %v\n", edges1)
	fmt.Printf("Center: %d\n", center)

	// Problem 2: Find if Path Exists
	fmt.Println("\n2. Find if Path Exists:")
	n := 3
	edges2 := [][]int{{0, 1}, {1, 2}, {2, 0}}
	pathExists := ValidPath(n, edges2, 0, 2)
	fmt.Printf("Graph: n=%d, edges=%v\n", n, edges2)
	fmt.Printf("Path from 0 to 2 exists: %t\n", pathExists)

	// Problem 3: Number of Connected Components
	fmt.Println("\n3. Number of Connected Components:")
	n3 := 5
	edges3 := [][]int{{0, 1}, {1, 2}, {3, 4}}
	components := CountComponents(n3, edges3)
	fmt.Printf("Nodes: %d, Edges: %v\n", n3, edges3)
	fmt.Printf("Connected components: %d\n", components)

	// Problem 4: All Paths Source to Target
	fmt.Println("\n4. All Paths Source to Target:")
	graph4 := [][]int{{1, 2}, {3}, {3}, {}}
	allPaths := AllPathsSourceTarget(graph4)
	fmt.Printf("Graph: %v\n", graph4)
	fmt.Printf("All paths from 0 to 3: %v\n", allPaths)

	// Problem 5: Flood Fill
	fmt.Println("\n5. Flood Fill:")
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Printf("Original image:\n")
	PrintMatrix(image)

	imageCopy := make([][]int, len(image))
	for i := range image {
		imageCopy[i] = make([]int, len(image[i]))
		copy(imageCopy[i], image[i])
	}

	filled := FloodFill(imageCopy, 1, 1, 2)
	fmt.Printf("After flood fill (1,1) with color 2:\n")
	PrintMatrix(filled)

	// Problem 6: Number of Islands
	fmt.Println("\n6. Number of Islands:")
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Printf("Grid:\n")
	PrintByteMatrix(grid)

	// Make copy since NumIslands modifies the grid
	gridCopy := make([][]byte, len(grid))
	for i := range grid {
		gridCopy[i] = make([]byte, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}

	islands := NumIslands(gridCopy)
	fmt.Printf("Number of islands: %d\n", islands)

	// Problem 7: Pacific Atlantic Water Flow
	fmt.Println("\n7. Pacific Atlantic Water Flow:")
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Printf("Heights:\n")
	PrintMatrix(heights)

	waterFlow := PacificAtlantic(heights)
	fmt.Printf("Cells reaching both oceans: %v\n", waterFlow)

	// Problem 8: Course Schedule
	fmt.Println("\n8. Course Schedule:")
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	canFinish := CanFinish(numCourses, prerequisites)
	fmt.Printf("Courses: %d, Prerequisites: %v\n", numCourses, prerequisites)
	fmt.Printf("Can finish all courses: %t\n", canFinish)

	// Problem 9: Surrounded Regions (demonstration)
	fmt.Println("\n9. Surrounded Regions:")
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	fmt.Printf("Original board:\n")
	PrintByteMatrix(board)

	Solve(board)
	fmt.Printf("After capturing surrounded regions:\n")
	PrintByteMatrix(board)
}

// ProblemComplexityAnalysis provides complexity analysis for all problems
func ProblemComplexityAnalysis() {
	fmt.Println("\n📊 Easy Graph Problems Complexity Analysis")
	fmt.Println("==========================================")

	problems := []struct {
		name     string
		timeOpt  string
		spaceOpt string
		approach string
	}{
		{"Find Center Star Graph", "O(1)", "O(1)", "Edge Intersection"},
		{"Path Exists", "O(V+E)", "O(V)", "DFS/BFS"},
		{"Connected Components", "O(V+E)", "O(V)", "Union-Find/DFS"},
		{"Clone Graph", "O(V+E)", "O(V)", "DFS/BFS"},
		{"All Paths Source Target", "O(2^V×V)", "O(V)", "DFS Backtrack"},
		{"Flood Fill", "O(M×N)", "O(M×N)", "DFS/BFS"},
		{"Number of Islands", "O(M×N)", "O(M×N)", "DFS/BFS"},
		{"Surrounded Regions", "O(M×N)", "O(M×N)", "Boundary DFS"},
		{"Pacific Atlantic", "O(M×N)", "O(M×N)", "Multi-source DFS"},
		{"Course Schedule", "O(V+E)", "O(V)", "Cycle Detection"},
	}

	fmt.Printf("%-25s %-12s %-10s %-20s\n", "Problem", "Time", "Space", "Best Approach")
	fmt.Println(strings.Repeat("-", 70))

	for _, p := range problems {
		fmt.Printf("%-25s %-12s %-10s %-20s\n", p.name, p.timeOpt, p.spaceOpt, p.approach)
	}
}
