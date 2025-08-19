package graphs

import (
	"errors"
)

// TraversalResult contains the result of a graph traversal
type TraversalResult struct {
	Order         []int       // Order of vertex visits
	DiscoveryTime map[int]int // When each vertex was first discovered
	FinishTime    map[int]int // When each vertex was finished (DFS only)
	Parent        map[int]int // Parent of each vertex in traversal tree
	Distance      map[int]int // Distance from source (BFS only)
}

// DFS performs Depth-First Search starting from a given vertex
func DFS(graph Graph, start int) (*TraversalResult, error) {
	if !graph.HasVertex(start) {
		return nil, errors.New("start vertex does not exist")
	}

	result := &TraversalResult{
		Order:         make([]int, 0),
		DiscoveryTime: make(map[int]int),
		FinishTime:    make(map[int]int),
		Parent:        make(map[int]int),
		Distance:      make(map[int]int),
	}

	visited := make(map[int]bool)
	time := 0

	var dfsVisit func(vertex int)
	dfsVisit = func(vertex int) {
		visited[vertex] = true
		time++
		result.DiscoveryTime[vertex] = time
		result.Order = append(result.Order, vertex)

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				result.Parent[neighbor] = vertex
				dfsVisit(neighbor)
			}
		}

		time++
		result.FinishTime[vertex] = time
	}

	// Initialize parent for start vertex
	result.Parent[start] = -1
	dfsVisit(start)

	return result, nil
}

// DFSComplete performs DFS on all vertices (handles disconnected graphs)
func DFSComplete(graph Graph) *TraversalResult {
	result := &TraversalResult{
		Order:         make([]int, 0),
		DiscoveryTime: make(map[int]int),
		FinishTime:    make(map[int]int),
		Parent:        make(map[int]int),
		Distance:      make(map[int]int),
	}

	visited := make(map[int]bool)
	time := 0

	var dfsVisit func(vertex int)
	dfsVisit = func(vertex int) {
		visited[vertex] = true
		time++
		result.DiscoveryTime[vertex] = time
		result.Order = append(result.Order, vertex)

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				result.Parent[neighbor] = vertex
				dfsVisit(neighbor)
			}
		}

		time++
		result.FinishTime[vertex] = time
	}

	// Visit all vertices
	vertices := graph.GetVertices()
	for _, vertex := range vertices {
		if !visited[vertex] {
			result.Parent[vertex] = -1
			dfsVisit(vertex)
		}
	}

	return result
}

// BFS performs Breadth-First Search starting from a given vertex
func BFS(graph Graph, start int) (*TraversalResult, error) {
	if !graph.HasVertex(start) {
		return nil, errors.New("start vertex does not exist")
	}

	result := &TraversalResult{
		Order:         make([]int, 0),
		DiscoveryTime: make(map[int]int),
		FinishTime:    make(map[int]int),
		Parent:        make(map[int]int),
		Distance:      make(map[int]int),
	}

	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true
	result.Parent[start] = -1
	result.Distance[start] = 0
	time := 0

	for len(queue) > 0 {
		// Dequeue
		vertex := queue[0]
		queue = queue[1:]

		time++
		result.DiscoveryTime[vertex] = time
		result.Order = append(result.Order, vertex)

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				result.Parent[neighbor] = vertex
				result.Distance[neighbor] = result.Distance[vertex] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	return result, nil
}

// HasPath checks if there's a path between two vertices using BFS
func HasPath(graph Graph, from, to int) bool {
	if !graph.HasVertex(from) || !graph.HasVertex(to) {
		return false
	}

	if from == to {
		return true
	}

	visited := make(map[int]bool)
	queue := []int{from}
	visited[from] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if neighbor == to {
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

// ShortestPath finds the shortest path between two vertices using BFS
func ShortestPath(graph Graph, from, to int) ([]int, error) {
	if !graph.HasVertex(from) || !graph.HasVertex(to) {
		return nil, errors.New("one or both vertices do not exist")
	}

	if from == to {
		return []int{from}, nil
	}

	result, err := BFS(graph, from)
	if err != nil {
		return nil, err
	}

	// Check if 'to' is reachable
	if _, exists := result.Parent[to]; !exists {
		return nil, errors.New("no path exists")
	}

	// Reconstruct path
	path := make([]int, 0)
	current := to
	for current != -1 {
		path = append([]int{current}, path...) // Prepend
		current = result.Parent[current]
	}

	return path, nil
}

// IsConnected checks if the graph is connected (all vertices reachable from any vertex)
func IsConnected(graph Graph) bool {
	vertices := graph.GetVertices()
	if len(vertices) == 0 {
		return true
	}

	// For directed graphs, we need to check strong connectivity
	// For now, we'll check if all vertices are reachable from the first vertex
	result, err := BFS(graph, vertices[0])
	if err != nil {
		return false
	}

	return len(result.Order) == graph.VertexCount()
}

// TopologicalSort performs topological sorting on a directed acyclic graph
func TopologicalSort(graph Graph) ([]int, error) {
	if !graph.IsDirected() {
		return nil, errors.New("topological sort requires a directed graph")
	}

	result := DFSComplete(graph)

	// Sort vertices by finish time in descending order
	vertices := make([]int, len(result.Order))
	copy(vertices, result.Order)

	// Sort by finish time (descending)
	for i := 0; i < len(vertices)-1; i++ {
		for j := i + 1; j < len(vertices); j++ {
			if result.FinishTime[vertices[i]] < result.FinishTime[vertices[j]] {
				vertices[i], vertices[j] = vertices[j], vertices[i]
			}
		}
	}

	return vertices, nil
}

// DetectCycle detects if there's a cycle in the graph
func DetectCycle(graph Graph) bool {
	if graph.IsDirected() {
		return detectCycleDirected(graph)
	}
	return detectCycleUndirected(graph)
}

// detectCycleDirected detects cycle in directed graph using DFS
func detectCycleDirected(graph Graph) bool {
	vertices := graph.GetVertices()
	white := make(map[int]bool) // Not visited
	gray := make(map[int]bool)  // Visiting
	black := make(map[int]bool) // Visited

	// Initialize all vertices as white
	for _, vertex := range vertices {
		white[vertex] = true
	}

	var hasCycle bool
	var dfsVisit func(vertex int)
	dfsVisit = func(vertex int) {
		if hasCycle {
			return
		}

		white[vertex] = false
		gray[vertex] = true

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if gray[neighbor] {
				// Back edge found - cycle detected
				hasCycle = true
				return
			}
			if white[neighbor] {
				dfsVisit(neighbor)
			}
		}

		gray[vertex] = false
		black[vertex] = true
	}

	for _, vertex := range vertices {
		if white[vertex] {
			dfsVisit(vertex)
			if hasCycle {
				return true
			}
		}
	}

	return false
}

// detectCycleUndirected detects cycle in undirected graph using DFS
func detectCycleUndirected(graph Graph) bool {
	vertices := graph.GetVertices()
	visited := make(map[int]bool)

	var hasCycle bool
	var dfsVisit func(vertex, parent int)
	dfsVisit = func(vertex, parent int) {
		if hasCycle {
			return
		}

		visited[vertex] = true

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				dfsVisit(neighbor, vertex)
			} else if neighbor != parent {
				// Found a back edge (not to parent) - cycle detected
				hasCycle = true
				return
			}
		}
	}

	for _, vertex := range vertices {
		if !visited[vertex] {
			dfsVisit(vertex, -1)
			if hasCycle {
				return true
			}
		}
	}

	return false
}

// ConnectedComponents finds all connected components in an undirected graph
func ConnectedComponents(graph Graph) [][]int {
	if graph.IsDirected() {
		// For directed graphs, this would be strongly connected components
		// which is more complex - return empty for now
		return [][]int{}
	}

	vertices := graph.GetVertices()
	visited := make(map[int]bool)
	components := make([][]int, 0)

	var dfsVisit func(vertex int, component *[]int)
	dfsVisit = func(vertex int, component *[]int) {
		visited[vertex] = true
		*component = append(*component, vertex)

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				dfsVisit(neighbor, component)
			}
		}
	}

	for _, vertex := range vertices {
		if !visited[vertex] {
			component := make([]int, 0)
			dfsVisit(vertex, &component)
			components = append(components, component)
		}
	}

	return components
}
