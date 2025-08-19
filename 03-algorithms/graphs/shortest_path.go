package graphs

import (
	"errors"
	"math"
)

// ShortestPathResult contains the result of shortest path algorithms
type ShortestPathResult struct {
	Distances map[int]int // Distance from source to each vertex
	Parents   map[int]int // Parent of each vertex in shortest path tree
	Path      []int       // Actual path (for single target queries)
}

// PriorityQueueItem represents an item in the priority queue for Dijkstra's algorithm
type PriorityQueueItem struct {
	Vertex   int
	Distance int
}

// MinPriorityQueue is a simple min-heap for Dijkstra's algorithm
type MinPriorityQueue struct {
	items []PriorityQueueItem
}

func NewMinPriorityQueue() *MinPriorityQueue {
	return &MinPriorityQueue{
		items: make([]PriorityQueueItem, 0),
	}
}

func (pq *MinPriorityQueue) Insert(vertex, distance int) {
	item := PriorityQueueItem{Vertex: vertex, Distance: distance}
	pq.items = append(pq.items, item)
	pq.heapifyUp(len(pq.items) - 1)
}

func (pq *MinPriorityQueue) ExtractMin() (PriorityQueueItem, error) {
	if len(pq.items) == 0 {
		return PriorityQueueItem{}, errors.New("priority queue is empty")
	}

	min := pq.items[0]
	lastIndex := len(pq.items) - 1
	pq.items[0] = pq.items[lastIndex]
	pq.items = pq.items[:lastIndex]

	if len(pq.items) > 0 {
		pq.heapifyDown(0)
	}

	return min, nil
}

func (pq *MinPriorityQueue) IsEmpty() bool {
	return len(pq.items) == 0
}

func (pq *MinPriorityQueue) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.items[index].Distance >= pq.items[parentIndex].Distance {
			break
		}
		pq.items[index], pq.items[parentIndex] = pq.items[parentIndex], pq.items[index]
		index = parentIndex
	}
}

func (pq *MinPriorityQueue) heapifyDown(index int) {
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild < len(pq.items) && pq.items[leftChild].Distance < pq.items[smallest].Distance {
			smallest = leftChild
		}

		if rightChild < len(pq.items) && pq.items[rightChild].Distance < pq.items[smallest].Distance {
			smallest = rightChild
		}

		if smallest == index {
			break
		}

		pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
		index = smallest
	}
}

// Dijkstra implements Dijkstra's shortest path algorithm
func Dijkstra(graph Graph, source int) (*ShortestPathResult, error) {
	if !graph.HasVertex(source) {
		return nil, errors.New("source vertex does not exist")
	}

	result := &ShortestPathResult{
		Distances: make(map[int]int),
		Parents:   make(map[int]int),
	}

	// Initialize distances
	vertices := graph.GetVertices()
	for _, vertex := range vertices {
		result.Distances[vertex] = math.MaxInt32
		result.Parents[vertex] = -1
	}
	result.Distances[source] = 0

	// Priority queue for vertices to process
	pq := NewMinPriorityQueue()
	pq.Insert(source, 0)

	visited := make(map[int]bool)

	for !pq.IsEmpty() {
		current, err := pq.ExtractMin()
		if err != nil {
			break
		}

		vertex := current.Vertex
		if visited[vertex] {
			continue
		}
		visited[vertex] = true

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if visited[neighbor] {
				continue
			}

			weight, err := graph.GetWeight(vertex, neighbor)
			if err != nil {
				continue
			}

			newDistance := result.Distances[vertex] + weight
			if newDistance < result.Distances[neighbor] {
				result.Distances[neighbor] = newDistance
				result.Parents[neighbor] = vertex
				pq.Insert(neighbor, newDistance)
			}
		}
	}

	return result, nil
}

// DijkstraPath finds the shortest path between source and target
func DijkstraPath(graph Graph, source, target int) (*ShortestPathResult, error) {
	result, err := Dijkstra(graph, source)
	if err != nil {
		return nil, err
	}

	// Check if target is reachable
	if result.Distances[target] == math.MaxInt32 {
		return nil, errors.New("target is not reachable from source")
	}

	// Reconstruct path
	path := make([]int, 0)
	current := target
	for current != -1 {
		path = append([]int{current}, path...) // Prepend
		current = result.Parents[current]
	}

	result.Path = path
	return result, nil
}

// FloydWarshallResult contains all-pairs shortest paths
type FloydWarshallResult struct {
	Distances [][]int // Distance matrix
	Next      [][]int // Next vertex in shortest path
	Vertices  []int   // Vertex mapping
}

// FloydWarshall implements Floyd-Warshall all-pairs shortest path algorithm
func FloydWarshall(graph Graph) *FloydWarshallResult {
	vertices := graph.GetVertices()
	n := len(vertices)

	// Create vertex index mapping
	vertexToIndex := make(map[int]int)
	for i, vertex := range vertices {
		vertexToIndex[vertex] = i
	}

	// Initialize distance matrix
	dist := make([][]int, n)
	next := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		next[i] = make([]int, n)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
			next[i][j] = -1
		}
	}

	// Fill initial distances from graph edges
	for _, vertex := range vertices {
		i := vertexToIndex[vertex]
		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			j := vertexToIndex[neighbor]
			weight, err := graph.GetWeight(vertex, neighbor)
			if err == nil {
				dist[i][j] = weight
				next[i][j] = j
			}
		}
	}

	// Floyd-Warshall algorithm
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					newDist := dist[i][k] + dist[k][j]
					if newDist < dist[i][j] {
						dist[i][j] = newDist
						next[i][j] = next[i][k]
					}
				}
			}
		}
	}

	return &FloydWarshallResult{
		Distances: dist,
		Next:      next,
		Vertices:  vertices,
	}
}

// GetPath reconstructs path from Floyd-Warshall result
func (fw *FloydWarshallResult) GetPath(from, to int) []int {
	// Find indices
	fromIndex := -1
	toIndex := -1
	for i, vertex := range fw.Vertices {
		if vertex == from {
			fromIndex = i
		}
		if vertex == to {
			toIndex = i
		}
	}

	if fromIndex == -1 || toIndex == -1 {
		return nil
	}

	if fw.Next[fromIndex][toIndex] == -1 {
		return nil // No path
	}

	path := []int{from}
	current := fromIndex
	for current != toIndex {
		current = fw.Next[current][toIndex]
		path = append(path, fw.Vertices[current])
	}

	return path
}

// GetDistance returns distance between two vertices from Floyd-Warshall result
func (fw *FloydWarshallResult) GetDistance(from, to int) int {
	fromIndex := -1
	toIndex := -1
	for i, vertex := range fw.Vertices {
		if vertex == from {
			fromIndex = i
		}
		if vertex == to {
			toIndex = i
		}
	}

	if fromIndex == -1 || toIndex == -1 {
		return math.MaxInt32
	}

	return fw.Distances[fromIndex][toIndex]
}

// BellmanFord implements Bellman-Ford algorithm (handles negative weights)
func BellmanFord(graph Graph, source int) (*ShortestPathResult, error) {
	if !graph.HasVertex(source) {
		return nil, errors.New("source vertex does not exist")
	}

	result := &ShortestPathResult{
		Distances: make(map[int]int),
		Parents:   make(map[int]int),
	}

	vertices := graph.GetVertices()

	// Initialize distances
	for _, vertex := range vertices {
		result.Distances[vertex] = math.MaxInt32
		result.Parents[vertex] = -1
	}
	result.Distances[source] = 0

	// Relax all edges |V| - 1 times
	for i := 0; i < len(vertices)-1; i++ {
		for _, u := range vertices {
			if result.Distances[u] == math.MaxInt32 {
				continue
			}

			neighbors := graph.GetNeighbors(u)
			for _, v := range neighbors {
				weight, err := graph.GetWeight(u, v)
				if err != nil {
					continue
				}

				newDistance := result.Distances[u] + weight
				if newDistance < result.Distances[v] {
					result.Distances[v] = newDistance
					result.Parents[v] = u
				}
			}
		}
	}

	// Check for negative cycles
	for _, u := range vertices {
		if result.Distances[u] == math.MaxInt32 {
			continue
		}

		neighbors := graph.GetNeighbors(u)
		for _, v := range neighbors {
			weight, err := graph.GetWeight(u, v)
			if err != nil {
				continue
			}

			if result.Distances[u]+weight < result.Distances[v] {
				return nil, errors.New("graph contains negative cycle")
			}
		}
	}

	return result, nil
}

// HasNegativeCycle checks if the graph has a negative cycle using Bellman-Ford
func HasNegativeCycle(graph Graph) bool {
	vertices := graph.GetVertices()
	if len(vertices) == 0 {
		return false
	}

	// Try Bellman-Ford from the first vertex
	_, err := BellmanFord(graph, vertices[0])
	return err != nil && err.Error() == "graph contains negative cycle"
}
