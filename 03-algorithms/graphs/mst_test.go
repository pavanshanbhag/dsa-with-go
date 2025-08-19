package graphs

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	tests := []struct {
		name     string
		vertices []int
		edges    [][]int // [vertex1, vertex2]
		expected [][]int // expected groups
	}{
		{
			name:     "Simple union find",
			vertices: []int{0, 1, 2, 3, 4},
			edges:    [][]int{{0, 1}, {2, 3}},
			expected: [][]int{{0, 1}, {2, 3}, {4}},
		},
		{
			name:     "All connected",
			vertices: []int{0, 1, 2, 3},
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}},
			expected: [][]int{{0, 1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := NewUnionFind(tt.vertices)

			for _, edge := range tt.edges {
				uf.Union(edge[0], edge[1])
			}

			// Check if vertices that should be connected are in same component
			for _, group := range tt.expected {
				if len(group) > 1 {
					root := uf.Find(group[0])
					for i := 1; i < len(group); i++ {
						if uf.Find(group[i]) != root {
							t.Errorf("Vertices %d and %d should be in same component", group[0], group[i])
						}
					}
				}
			}
		})
	}
}

func TestKruskal(t *testing.T) {
	tests := []struct {
		name         string
		setupGraph   func() Graph
		expectedCost int
		expectError  bool
	}{
		{
			name: "Simple connected graph",
			setupGraph: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddEdge(0, 1, 10)
				g.AddEdge(0, 2, 6)
				g.AddEdge(0, 3, 5)
				g.AddEdge(1, 3, 15)
				g.AddEdge(2, 3, 4)
				return g
			},
			expectedCost: 19, // edges: (2,3)=4, (0,3)=5, (0,1)=10
			expectError:  false,
		},
		{
			name: "Triangle graph",
			setupGraph: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddEdge(0, 1, 1)
				g.AddEdge(1, 2, 2)
				g.AddEdge(0, 2, 3)
				return g
			},
			expectedCost: 3, // edges: (0,1)=1, (1,2)=2
			expectError:  false,
		},
		{
			name: "Single vertex",
			setupGraph: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				return g
			},
			expectedCost: 0,
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setupGraph()
			mst, err := Kruskal(graph)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if mst.TotalCost != tt.expectedCost {
				t.Errorf("Expected cost %d, got %d", tt.expectedCost, mst.TotalCost)
			}

			// For connected graphs, MST should have vertices-1 edges
			if graph.VertexCount() > 1 {
				expectedEdges := graph.VertexCount() - 1
				if len(mst.Edges) != expectedEdges {
					t.Errorf("Expected %d edges in MST, got %d", expectedEdges, len(mst.Edges))
				}
			}
		})
	}
}

func TestPrim(t *testing.T) {
	tests := []struct {
		name         string
		setupGraph   func() Graph
		startVertex  int
		expectedCost int
		expectError  bool
	}{
		{
			name: "Simple connected graph",
			setupGraph: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddEdge(0, 1, 10)
				g.AddEdge(0, 2, 6)
				g.AddEdge(0, 3, 5)
				g.AddEdge(1, 3, 15)
				g.AddEdge(2, 3, 4)
				return g
			},
			startVertex:  0,
			expectedCost: 19,
			expectError:  false,
		},
		{
			name: "Triangle graph",
			setupGraph: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddEdge(0, 1, 1)
				g.AddEdge(1, 2, 2)
				g.AddEdge(0, 2, 3)
				return g
			},
			startVertex:  0,
			expectedCost: 3,
			expectError:  false,
		},
		{
			name: "Single vertex",
			setupGraph: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				return g
			},
			startVertex:  0,
			expectedCost: 0,
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setupGraph()
			mst, err := Prim(graph, tt.startVertex)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if mst.TotalCost != tt.expectedCost {
				t.Errorf("Expected cost %d, got %d", tt.expectedCost, mst.TotalCost)
			}

			// For connected graphs, MST should have vertices-1 edges
			if graph.VertexCount() > 1 {
				expectedEdges := graph.VertexCount() - 1
				if len(mst.Edges) != expectedEdges {
					t.Errorf("Expected %d edges in MST, got %d", expectedEdges, len(mst.Edges))
				}
			}
		})
	}
}

func BenchmarkKruskal(b *testing.B) {
	// Create a complete graph with random weights
	vertices := 20
	graph := NewAdjacencyListGraph(false)

	for i := 0; i < vertices; i++ {
		graph.AddVertex(i)
	}

	for i := 0; i < vertices; i++ {
		for j := i + 1; j < vertices; j++ {
			graph.AddEdge(i, j, i*j+1)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Kruskal(graph)
	}
}

func BenchmarkPrim(b *testing.B) {
	// Create a complete graph with random weights
	vertices := 20
	graph := NewAdjacencyListGraph(false)

	for i := 0; i < vertices; i++ {
		graph.AddVertex(i)
	}

	for i := 0; i < vertices; i++ {
		for j := i + 1; j < vertices; j++ {
			graph.AddEdge(i, j, i*j+1)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Prim(graph, 0)
	}
}

func BenchmarkUnionFind(b *testing.B) {
	vertices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	b.Run("Union operations", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			uf := NewUnionFind(vertices)
			for j := 0; j < len(vertices)-1; j++ {
				uf.Union(j, j+1)
			}
		}
	})

	b.Run("Find operations", func(b *testing.B) {
		uf := NewUnionFind(vertices)
		// Pre-populate with some unions
		for j := 0; j < len(vertices)-1; j++ {
			uf.Union(j, j+1)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(vertices); j++ {
				uf.Find(j)
			}
		}
	})
}

// Example test to demonstrate usage
func ExampleKruskal() {
	graph := NewAdjacencyListGraph(false)

	// Add vertices
	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	// Add edges
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 2)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 3, 8)
	graph.AddEdge(2, 4, 10)
	graph.AddEdge(3, 4, 2)

	mst, err := Kruskal(graph)
	if err != nil {
		panic(err)
	}

	_ = mst
	// Output will vary based on implementation
}

func ExamplePrim() {
	graph := NewAdjacencyListGraph(false)

	// Add vertices
	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	// Add edges
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 2)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 3, 8)
	graph.AddEdge(2, 4, 10)
	graph.AddEdge(3, 4, 2)

	mst, err := Prim(graph, 0)
	if err != nil {
		panic(err)
	}

	_ = mst
	// Output will vary based on implementation
}
