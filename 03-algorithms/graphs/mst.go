package graphs

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

// MST represents a Minimum Spanning Tree
type MST struct {
	Edges     []Edge // Edges in the MST
	TotalCost int    // Total cost of the MST
	Vertices  []int  // Vertices included in MST
}

// UnionFind data structure for Kruskal's algorithm
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind creates a new Union-Find structure
func NewUnionFind(vertices []int) *UnionFind {
	n := len(vertices)
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &UnionFind{
		parent: parent,
		rank:   rank,
	}
}

// Find returns the root of the set containing x
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y
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
	}
}

// IsConnected checks if x and y are in the same set
func (uf *UnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Kruskal implements Kruskal's Minimum Spanning Tree algorithm
func Kruskal(graph Graph) (*MST, error) {
	if graph.IsDirected() {
		return nil, errors.New("MST requires an undirected graph")
	}

	vertices := graph.GetVertices()
	if len(vertices) == 0 {
		return &MST{}, nil
	}

	// Create vertex index mapping
	vertexToIndex := make(map[int]int)
	for i, vertex := range vertices {
		vertexToIndex[vertex] = i
	}

	// Collect all edges
	edges := make([]Edge, 0)
	for _, vertex := range vertices {
		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			// For undirected graph, avoid duplicate edges
			if vertex < neighbor {
				weight, err := graph.GetWeight(vertex, neighbor)
				if err == nil {
					edges = append(edges, Edge{From: vertex, To: neighbor, Weight: weight})
				}
			}
		}
	}

	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// Initialize Union-Find
	uf := NewUnionFind(vertices)
	mst := &MST{
		Edges:     make([]Edge, 0),
		TotalCost: 0,
		Vertices:  make([]int, 0),
	}

	// Process edges in sorted order
	for _, edge := range edges {
		fromIndex := vertexToIndex[edge.From]
		toIndex := vertexToIndex[edge.To]

		if !uf.IsConnected(fromIndex, toIndex) {
			uf.Union(fromIndex, toIndex)
			mst.Edges = append(mst.Edges, edge)
			mst.TotalCost += edge.Weight

			// Stop when we have n-1 edges (complete MST)
			if len(mst.Edges) == len(vertices)-1 {
				break
			}
		}
	}

	// Add all vertices to MST
	vertexSet := make(map[int]bool)
	for _, edge := range mst.Edges {
		vertexSet[edge.From] = true
		vertexSet[edge.To] = true
	}
	for vertex := range vertexSet {
		mst.Vertices = append(mst.Vertices, vertex)
	}
	sort.Ints(mst.Vertices)

	// Check if MST is complete (connected graph)
	if len(mst.Edges) != len(vertices)-1 {
		return nil, errors.New("graph is not connected - no spanning tree exists")
	}

	return mst, nil
}

// Prim implements Prim's Minimum Spanning Tree algorithm
func Prim(graph Graph, start int) (*MST, error) {
	if graph.IsDirected() {
		return nil, errors.New("MST requires an undirected graph")
	}

	if !graph.HasVertex(start) {
		return nil, errors.New("start vertex does not exist")
	}

	vertices := graph.GetVertices()
	if len(vertices) == 0 {
		return &MST{}, nil
	}

	mst := &MST{
		Edges:     make([]Edge, 0),
		TotalCost: 0,
		Vertices:  make([]int, 0),
	}

	inMST := make(map[int]bool)
	minWeight := make(map[int]int)
	parent := make(map[int]int)

	// Initialize all vertices with infinite weight
	for _, vertex := range vertices {
		minWeight[vertex] = math.MaxInt32
		parent[vertex] = -1
	}

	// Start with the given vertex
	minWeight[start] = 0

	for len(mst.Vertices) < len(vertices) {
		// Find minimum weight vertex not in MST
		minVertex := -1
		minCost := math.MaxInt32

		for _, vertex := range vertices {
			if !inMST[vertex] && minWeight[vertex] < minCost {
				minCost = minWeight[vertex]
				minVertex = vertex
			}
		}

		if minVertex == -1 {
			// Graph is not connected
			return nil, errors.New("graph is not connected - no spanning tree exists")
		}

		// Add vertex to MST
		inMST[minVertex] = true
		mst.Vertices = append(mst.Vertices, minVertex)

		// Add edge to MST (except for the start vertex)
		if parent[minVertex] != -1 {
			edge := Edge{From: parent[minVertex], To: minVertex, Weight: minCost}
			mst.Edges = append(mst.Edges, edge)
			mst.TotalCost += minCost
		}

		// Update weights for adjacent vertices
		neighbors := graph.GetNeighbors(minVertex)
		for _, neighbor := range neighbors {
			if !inMST[neighbor] {
				weight, err := graph.GetWeight(minVertex, neighbor)
				if err == nil && weight < minWeight[neighbor] {
					minWeight[neighbor] = weight
					parent[neighbor] = minVertex
				}
			}
		}
	}

	sort.Ints(mst.Vertices)
	return mst, nil
}

// PrimWithPriorityQueue implements Prim's algorithm using a priority queue (more efficient)
func PrimWithPriorityQueue(graph Graph, start int) (*MST, error) {
	if graph.IsDirected() {
		return nil, errors.New("MST requires an undirected graph")
	}

	if !graph.HasVertex(start) {
		return nil, errors.New("start vertex does not exist")
	}

	vertices := graph.GetVertices()
	if len(vertices) == 0 {
		return &MST{}, nil
	}

	mst := &MST{
		Edges:     make([]Edge, 0),
		TotalCost: 0,
		Vertices:  make([]int, 0),
	}

	inMST := make(map[int]bool)
	parent := make(map[int]int)
	pq := NewMinPriorityQueue()

	// Start with the given vertex
	pq.Insert(start, 0)
	parent[start] = -1

	for !pq.IsEmpty() && len(mst.Vertices) < len(vertices) {
		current, err := pq.ExtractMin()
		if err != nil {
			break
		}

		vertex := current.Vertex
		weight := current.Distance

		if inMST[vertex] {
			continue
		}

		// Add vertex to MST
		inMST[vertex] = true
		mst.Vertices = append(mst.Vertices, vertex)

		// Add edge to MST (except for the start vertex)
		if parent[vertex] != -1 {
			edge := Edge{From: parent[vertex], To: vertex, Weight: weight}
			mst.Edges = append(mst.Edges, edge)
			mst.TotalCost += weight
		}

		// Add adjacent vertices to priority queue
		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if !inMST[neighbor] {
				edgeWeight, err := graph.GetWeight(vertex, neighbor)
				if err == nil {
					parent[neighbor] = vertex
					pq.Insert(neighbor, edgeWeight)
				}
			}
		}
	}

	// Check if MST is complete
	if len(mst.Vertices) != len(vertices) {
		return nil, errors.New("graph is not connected - no spanning tree exists")
	}

	sort.Ints(mst.Vertices)
	return mst, nil
}

// IsMST checks if the given set of edges forms a valid MST for the graph
func IsMST(graph Graph, edges []Edge) bool {
	vertices := graph.GetVertices()
	n := len(vertices)

	// MST must have exactly n-1 edges
	if len(edges) != n-1 {
		return false
	}

	// Create vertex index mapping
	vertexToIndex := make(map[int]int)
	for i, vertex := range vertices {
		vertexToIndex[vertex] = i
	}

	// Check if edges form a connected tree using Union-Find
	uf := NewUnionFind(vertices)
	totalWeight := 0

	for _, edge := range edges {
		// Check if edge exists in graph
		weight, err := graph.GetWeight(edge.From, edge.To)
		if err != nil || weight != edge.Weight {
			return false
		}

		fromIndex := vertexToIndex[edge.From]
		toIndex := vertexToIndex[edge.To]

		// Check for cycles
		if uf.IsConnected(fromIndex, toIndex) {
			return false
		}

		uf.Union(fromIndex, toIndex)
		totalWeight += edge.Weight
	}

	// Compare with actual MST
	actualMST, err := Kruskal(graph)
	if err != nil {
		return false
	}

	return totalWeight == actualMST.TotalCost
}

// String representation for MST
func (mst *MST) String() string {
	result := fmt.Sprintf("MST (Total Cost: %d, Vertices: %d, Edges: %d)\n",
		mst.TotalCost, len(mst.Vertices), len(mst.Edges))

	result += "Edges:\n"
	for _, edge := range mst.Edges {
		result += fmt.Sprintf("  %d --(%d)-- %d\n", edge.From, edge.Weight, edge.To)
	}

	return result
}
