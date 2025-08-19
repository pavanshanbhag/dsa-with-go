package graphs

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

// Edge represents a weighted edge in a graph
type Edge struct {
	From   int
	To     int
	Weight int
}

// Graph interface defines common graph operations
type Graph interface {
	AddVertex(vertex int)
	AddEdge(from, to int, weight int)
	RemoveVertex(vertex int)
	RemoveEdge(from, to int)
	GetNeighbors(vertex int) []int
	GetVertices() []int
	HasVertex(vertex int) bool
	HasEdge(from, to int) bool
	GetWeight(from, to int) (int, error)
	VertexCount() int
	EdgeCount() int
	IsDirected() bool
}

// AdjacencyListGraph represents a graph using adjacency lists
type AdjacencyListGraph struct {
	adjacencyList map[int]map[int]int // vertex -> neighbor -> weight
	directed      bool
	vertexCount   int
	edgeCount     int
}

// NewAdjacencyListGraph creates a new adjacency list graph
func NewAdjacencyListGraph(directed bool) *AdjacencyListGraph {
	return &AdjacencyListGraph{
		adjacencyList: make(map[int]map[int]int),
		directed:      directed,
		vertexCount:   0,
		edgeCount:     0,
	}
}

// AddVertex adds a vertex to the graph
func (g *AdjacencyListGraph) AddVertex(vertex int) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		g.adjacencyList[vertex] = make(map[int]int)
		g.vertexCount++
	}
}

// AddEdge adds an edge to the graph
func (g *AdjacencyListGraph) AddEdge(from, to int, weight int) {
	// Ensure vertices exist
	g.AddVertex(from)
	g.AddVertex(to)

	// Add edge
	if _, exists := g.adjacencyList[from][to]; !exists {
		g.edgeCount++
	}
	g.adjacencyList[from][to] = weight

	// For undirected graphs, add reverse edge
	if !g.directed {
		g.adjacencyList[to][from] = weight
	}
}

// RemoveVertex removes a vertex and all its edges
func (g *AdjacencyListGraph) RemoveVertex(vertex int) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		return
	}

	// Remove all edges to this vertex
	for v := range g.adjacencyList {
		if _, hasEdge := g.adjacencyList[v][vertex]; hasEdge {
			delete(g.adjacencyList[v], vertex)
			g.edgeCount--
		}
	}

	// Remove all edges from this vertex
	edgeCount := len(g.adjacencyList[vertex])
	if g.directed {
		g.edgeCount -= edgeCount
	} else {
		// For undirected, we already removed reverse edges above
	}

	// Remove the vertex
	delete(g.adjacencyList, vertex)
	g.vertexCount--
}

// RemoveEdge removes an edge from the graph
func (g *AdjacencyListGraph) RemoveEdge(from, to int) {
	if _, exists := g.adjacencyList[from][to]; exists {
		delete(g.adjacencyList[from], to)
		g.edgeCount--

		// For undirected graphs, remove reverse edge
		if !g.directed {
			delete(g.adjacencyList[to], from)
		}
	}
}

// GetNeighbors returns all neighbors of a vertex
func (g *AdjacencyListGraph) GetNeighbors(vertex int) []int {
	neighbors := make([]int, 0, len(g.adjacencyList[vertex]))
	for neighbor := range g.adjacencyList[vertex] {
		neighbors = append(neighbors, neighbor)
	}
	sort.Ints(neighbors) // For consistent ordering
	return neighbors
}

// GetVertices returns all vertices in the graph
func (g *AdjacencyListGraph) GetVertices() []int {
	vertices := make([]int, 0, g.vertexCount)
	for vertex := range g.adjacencyList {
		vertices = append(vertices, vertex)
	}
	sort.Ints(vertices) // For consistent ordering
	return vertices
}

// HasVertex checks if a vertex exists
func (g *AdjacencyListGraph) HasVertex(vertex int) bool {
	_, exists := g.adjacencyList[vertex]
	return exists
}

// HasEdge checks if an edge exists
func (g *AdjacencyListGraph) HasEdge(from, to int) bool {
	if neighbors, exists := g.adjacencyList[from]; exists {
		_, hasEdge := neighbors[to]
		return hasEdge
	}
	return false
}

// GetWeight returns the weight of an edge
func (g *AdjacencyListGraph) GetWeight(from, to int) (int, error) {
	if neighbors, exists := g.adjacencyList[from]; exists {
		if weight, hasEdge := neighbors[to]; hasEdge {
			return weight, nil
		}
	}
	return 0, errors.New("edge does not exist")
}

// VertexCount returns the number of vertices
func (g *AdjacencyListGraph) VertexCount() int {
	return g.vertexCount
}

// EdgeCount returns the number of edges
func (g *AdjacencyListGraph) EdgeCount() int {
	return g.edgeCount
}

// IsDirected returns whether the graph is directed
func (g *AdjacencyListGraph) IsDirected() bool {
	return g.directed
}

// AdjacencyMatrixGraph represents a graph using adjacency matrix
type AdjacencyMatrixGraph struct {
	matrix        [][]int
	vertices      map[int]int // vertex value -> matrix index
	indexToVertex map[int]int // matrix index -> vertex value
	directed      bool
	vertexCount   int
	edgeCount     int
	maxVertices   int
}

// NewAdjacencyMatrixGraph creates a new adjacency matrix graph
func NewAdjacencyMatrixGraph(directed bool, maxVertices int) *AdjacencyMatrixGraph {
	matrix := make([][]int, maxVertices)
	for i := range matrix {
		matrix[i] = make([]int, maxVertices)
		for j := range matrix[i] {
			matrix[i][j] = math.MaxInt32 // Represents no edge
		}
	}

	return &AdjacencyMatrixGraph{
		matrix:        matrix,
		vertices:      make(map[int]int),
		indexToVertex: make(map[int]int),
		directed:      directed,
		vertexCount:   0,
		edgeCount:     0,
		maxVertices:   maxVertices,
	}
}

// AddVertex adds a vertex to the matrix graph
func (g *AdjacencyMatrixGraph) AddVertex(vertex int) {
	if _, exists := g.vertices[vertex]; !exists && g.vertexCount < g.maxVertices {
		index := g.vertexCount
		g.vertices[vertex] = index
		g.indexToVertex[index] = vertex
		g.vertexCount++
	}
}

// AddEdge adds an edge to the matrix graph
func (g *AdjacencyMatrixGraph) AddEdge(from, to int, weight int) {
	g.AddVertex(from)
	g.AddVertex(to)

	fromIndex, fromExists := g.vertices[from]
	toIndex, toExists := g.vertices[to]

	if fromExists && toExists {
		if g.matrix[fromIndex][toIndex] == math.MaxInt32 {
			g.edgeCount++
		}
		g.matrix[fromIndex][toIndex] = weight

		if !g.directed {
			g.matrix[toIndex][fromIndex] = weight
		}
	}
}

// RemoveVertex removes a vertex from matrix graph
func (g *AdjacencyMatrixGraph) RemoveVertex(vertex int) {
	index, exists := g.vertices[vertex]
	if !exists {
		return
	}

	// Count and remove edges
	for i := 0; i < g.vertexCount; i++ {
		if g.matrix[index][i] != math.MaxInt32 {
			g.matrix[index][i] = math.MaxInt32
			g.edgeCount--
		}
		if g.matrix[i][index] != math.MaxInt32 && g.directed {
			g.matrix[i][index] = math.MaxInt32
			g.edgeCount--
		}
	}

	// Remove vertex mapping
	delete(g.vertices, vertex)
	delete(g.indexToVertex, index)
	g.vertexCount--
}

// RemoveEdge removes an edge from matrix graph
func (g *AdjacencyMatrixGraph) RemoveEdge(from, to int) {
	fromIndex, fromExists := g.vertices[from]
	toIndex, toExists := g.vertices[to]

	if fromExists && toExists && g.matrix[fromIndex][toIndex] != math.MaxInt32 {
		g.matrix[fromIndex][toIndex] = math.MaxInt32
		g.edgeCount--

		if !g.directed {
			g.matrix[toIndex][fromIndex] = math.MaxInt32
		}
	}
}

// GetNeighbors returns neighbors for matrix graph
func (g *AdjacencyMatrixGraph) GetNeighbors(vertex int) []int {
	index, exists := g.vertices[vertex]
	if !exists {
		return []int{}
	}

	var neighbors []int
	for i := 0; i < g.maxVertices; i++ {
		if g.matrix[index][i] != math.MaxInt32 {
			if neighbor, exists := g.indexToVertex[i]; exists {
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	sort.Ints(neighbors)
	return neighbors
}

// GetVertices returns all vertices for matrix graph
func (g *AdjacencyMatrixGraph) GetVertices() []int {
	vertices := make([]int, 0, g.vertexCount)
	for vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}
	sort.Ints(vertices)
	return vertices
}

// HasVertex checks if vertex exists in matrix graph
func (g *AdjacencyMatrixGraph) HasVertex(vertex int) bool {
	_, exists := g.vertices[vertex]
	return exists
}

// HasEdge checks if edge exists in matrix graph
func (g *AdjacencyMatrixGraph) HasEdge(from, to int) bool {
	fromIndex, fromExists := g.vertices[from]
	toIndex, toExists := g.vertices[to]

	if fromExists && toExists {
		return g.matrix[fromIndex][toIndex] != math.MaxInt32
	}
	return false
}

// GetWeight returns edge weight for matrix graph
func (g *AdjacencyMatrixGraph) GetWeight(from, to int) (int, error) {
	fromIndex, fromExists := g.vertices[from]
	toIndex, toExists := g.vertices[to]

	if fromExists && toExists && g.matrix[fromIndex][toIndex] != math.MaxInt32 {
		return g.matrix[fromIndex][toIndex], nil
	}
	return 0, errors.New("edge does not exist")
}

// VertexCount returns vertex count for matrix graph
func (g *AdjacencyMatrixGraph) VertexCount() int {
	return g.vertexCount
}

// EdgeCount returns edge count for matrix graph
func (g *AdjacencyMatrixGraph) EdgeCount() int {
	return g.edgeCount
}

// IsDirected returns if matrix graph is directed
func (g *AdjacencyMatrixGraph) IsDirected() bool {
	return g.directed
}

// String representation for debugging
func (g *AdjacencyListGraph) String() string {
	result := fmt.Sprintf("Graph (directed: %v, vertices: %d, edges: %d)\n",
		g.directed, g.vertexCount, g.edgeCount)

	vertices := g.GetVertices()
	for _, vertex := range vertices {
		neighbors := g.GetNeighbors(vertex)
		result += fmt.Sprintf("  %d -> %v\n", vertex, neighbors)
	}
	return result
}

func (g *AdjacencyMatrixGraph) String() string {
	result := fmt.Sprintf("Matrix Graph (directed: %v, vertices: %d, edges: %d)\n",
		g.directed, g.vertexCount, g.edgeCount)

	vertices := g.GetVertices()
	for _, vertex := range vertices {
		neighbors := g.GetNeighbors(vertex)
		result += fmt.Sprintf("  %d -> %v\n", vertex, neighbors)
	}
	return result
}
