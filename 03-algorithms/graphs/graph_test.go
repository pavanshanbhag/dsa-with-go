package graphs

import (
	"reflect"
	"sort"
	"testing"
)

func TestDFS(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() Graph
		start    int
		expected []int
	}{
		{
			name: "Simple connected graph",
			setup: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddEdge(0, 1, 1)
				g.AddEdge(0, 2, 1)
				g.AddEdge(1, 3, 1)
				return g
			},
			start:    0,
			expected: []int{0, 1, 2, 3}, // Order may vary, just check all vertices visited
		},
		{
			name: "Single vertex",
			setup: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				return g
			},
			start:    0,
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			result, err := DFS(graph, tt.start)
			if err != nil {
				t.Fatalf("DFS failed: %v", err)
			}

			// Sort both arrays to compare regardless of traversal order
			sort.Ints(result.Order)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(result.Order, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result.Order)
			}
		})
	}
}

func TestBFS(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() Graph
		start    int
		expected []int
	}{
		{
			name: "Simple connected graph",
			setup: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddEdge(0, 1, 1)
				g.AddEdge(0, 2, 1)
				g.AddEdge(1, 3, 1)
				return g
			},
			start:    0,
			expected: []int{0, 1, 2, 3},
		},
		{
			name: "Single vertex",
			setup: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				return g
			},
			start:    0,
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			result, err := BFS(graph, tt.start)
			if err != nil {
				t.Fatalf("BFS failed: %v", err)
			}

			// Sort both arrays to compare regardless of traversal order
			sort.Ints(result.Order)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(result.Order, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result.Order)
			}
		})
	}
}

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name          string
		setup         func() Graph
		start         int
		expectedDists map[int]int
		expectError   bool
	}{
		{
			name: "Simple graph",
			setup: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddEdge(0, 1, 4)
				g.AddEdge(0, 2, 2)
				g.AddEdge(1, 2, 1)
				g.AddEdge(1, 3, 5)
				g.AddEdge(2, 3, 8)
				return g
			},
			start: 0,
			expectedDists: map[int]int{
				0: 0,
				1: 3, // 0->2->1 = 2+1 = 3
				2: 2, // 0->2 = 2
				3: 8, // 0->2->1->3 = 2+1+5 = 8
			},
			expectError: false,
		},
		{
			name: "Single vertex",
			setup: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				return g
			},
			start: 0,
			expectedDists: map[int]int{
				0: 0,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			result, err := Dijkstra(graph, tt.start)

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

			for vertex, expectedDist := range tt.expectedDists {
				if actualDist, exists := result.Distances[vertex]; !exists {
					t.Errorf("Missing distance for vertex %d", vertex)
				} else if actualDist != expectedDist {
					t.Errorf("For vertex %d: expected distance %d, got %d", vertex, expectedDist, actualDist)
				}
			}
		})
	}
}

func TestIsConnected(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() Graph
		expected bool
	}{
		{
			name: "Connected graph",
			setup: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddEdge(0, 1, 1)
				g.AddEdge(1, 2, 1)
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph",
			setup: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddEdge(0, 1, 1)
				g.AddEdge(2, 3, 1)
				return g
			},
			expected: false,
		},
		{
			name: "Single vertex",
			setup: func() Graph {
				g := NewAdjacencyListGraph(false)
				g.AddVertex(0)
				return g
			},
			expected: true,
		},
		{
			name: "Empty graph",
			setup: func() Graph {
				return NewAdjacencyListGraph(false)
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			result := IsConnected(graph)

			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func BenchmarkDFS(b *testing.B) {
	// Create a graph with many vertices
	graph := NewAdjacencyListGraph(false)
	vertices := 100

	for i := 0; i < vertices; i++ {
		graph.AddVertex(i)
	}

	// Create a connected graph
	for i := 0; i < vertices-1; i++ {
		graph.AddEdge(i, i+1, 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DFS(graph, 0)
	}
}

func BenchmarkBFS(b *testing.B) {
	// Create a graph with many vertices
	graph := NewAdjacencyListGraph(false)
	vertices := 100

	for i := 0; i < vertices; i++ {
		graph.AddVertex(i)
	}

	// Create a connected graph
	for i := 0; i < vertices-1; i++ {
		graph.AddEdge(i, i+1, 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = BFS(graph, 0)
	}
}

func BenchmarkDijkstra(b *testing.B) {
	// Create a complete graph
	graph := NewAdjacencyListGraph(false)
	vertices := 20 // Smaller for Dijkstra as it's more expensive

	for i := 0; i < vertices; i++ {
		graph.AddVertex(i)
	}

	for i := 0; i < vertices; i++ {
		for j := i + 1; j < vertices; j++ {
			graph.AddEdge(i, j, i+j+1)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Dijkstra(graph, 0)
	}
}

// Example functions for documentation
func ExampleDFS() {
	graph := NewAdjacencyListGraph(false)

	// Add vertices
	for i := 0; i < 4; i++ {
		graph.AddVertex(i)
	}

	// Add edges
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(0, 2, 1)
	graph.AddEdge(1, 3, 1)

	result, err := DFS(graph, 0)
	if err != nil {
		panic(err)
	}

	_ = result
	// Output will contain all reachable vertices from 0
}

func ExampleBFS() {
	graph := NewAdjacencyListGraph(false)

	// Add vertices
	for i := 0; i < 4; i++ {
		graph.AddVertex(i)
	}

	// Add edges
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(0, 2, 1)
	graph.AddEdge(1, 3, 1)

	result, err := BFS(graph, 0)
	if err != nil {
		panic(err)
	}

	_ = result
	// Output will contain all reachable vertices from 0 in BFS order
}

func ExampleDijkstra() {
	graph := NewAdjacencyListGraph(false)

	// Add vertices
	for i := 0; i < 4; i++ {
		graph.AddVertex(i)
	}

	// Add edges
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 2)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 3, 8)

	result, err := Dijkstra(graph, 0)
	if err != nil {
		panic(err)
	}

	_ = result
	// Output will contain shortest distances from vertex 0 to all other vertices
}
