package main

import (
	"fmt"
	"time"

	"dsa-mastery/03-algorithms/graphs"
)

func main() {
	fmt.Println("🕸️  DSA Mastery - Graph Algorithms Demonstration")
	fmt.Println("===============================================")

	// 1. Graph Traversal Algorithms
	fmt.Println("\n🚶 1. Graph Traversal Algorithms")
	fmt.Println("--------------------------------")
	demonstrateTraversal()

	// 2. Shortest Path Algorithms
	fmt.Println("\n🛣️  2. Shortest Path Algorithms")
	fmt.Println("------------------------------")
	demonstrateShortestPath()

	// 3. Minimum Spanning Tree
	fmt.Println("\n🌳 3. Minimum Spanning Tree Algorithms")
	fmt.Println("--------------------------------------")
	demonstrateMST()

	// 4. Graph Properties and Analysis
	fmt.Println("\n🔍 4. Graph Properties and Analysis")
	fmt.Println("----------------------------------")
	demonstrateGraphAnalysis()

	// 5. Real-World Applications
	fmt.Println("\n🌟 5. Real-World Applications")
	fmt.Println("-----------------------------")
	demonstrateRealWorldApplications()

	fmt.Println("\n✅ All graph algorithms demonstrated successfully!")
	fmt.Println("🚀 Ready to solve complex graph problems!")
}

func demonstrateTraversal() {
	fmt.Println("Problem: Traverse graphs using DFS and BFS")

	// Create a sample graph
	//     0
	//   /   \\
	//  1     2
	//  |    / \\
	//  3   4   5
	graph := graphs.NewAdjacencyListGraph(false)
	vertices := []int{0, 1, 2, 3, 4, 5}
	for _, v := range vertices {
		graph.AddVertex(v)
	}

	edges := [][3]int{
		{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {2, 4, 1}, {2, 5, 1},
	}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1], edge[2])
	}

	fmt.Printf("Graph structure: %d vertices, %d edges\n", graph.VertexCount(), graph.EdgeCount())
	fmt.Println("Edges: (0-1), (0-2), (1-3), (2-4), (2-5)")

	// DFS traversal
	fmt.Println("\n🔍 Depth-First Search (DFS):")
	start := time.Now()
	dfsResult, err := graphs.DFS(graph, 0)
	duration := time.Since(start)
	if err != nil {
		fmt.Printf("DFS Error: %v\n", err)
	} else {
		fmt.Printf("DFS from vertex 0: %v (Time: %v)\n", dfsResult.Order, duration)
		fmt.Printf("DFS discovery times: %v\n", dfsResult.DiscoveryTime)
	}

	// BFS traversal
	fmt.Println("\n🌐 Breadth-First Search (BFS):")
	start = time.Now()
	bfsResult, err := graphs.BFS(graph, 0)
	duration = time.Since(start)
	if err != nil {
		fmt.Printf("BFS Error: %v\n", err)
	} else {
		fmt.Printf("BFS from vertex 0: %v (Time: %v)\n", bfsResult.Order, duration)
		fmt.Printf("BFS distances: %v\n", bfsResult.Distance)
	}

	// Path finding
	fmt.Println("\n🎯 Path Finding:")
	path, err := graphs.ShortestPath(graph, 0, 5)
	if err != nil {
		fmt.Printf("Path finding error: %v\n", err)
	} else {
		fmt.Printf("Shortest path from 0 to 5: %v\n", path)
	}

	// Complete traversal for disconnected components
	fmt.Println("\n🔗 Connected Components:")
	components := graphs.ConnectedComponents(graph)
	fmt.Printf("Connected components: %v\n", components)
}

func demonstrateShortestPath() {
	fmt.Println("Problem: Find shortest paths in weighted graphs")

	// Create a weighted directed graph
	//     0 ----5----> 1
	//     |           /|
	//     2          4 |
	//     |         /  3
	//     v        v   |
	//     2 --1--> 3 --2-> 4
	graph := graphs.NewAdjacencyListGraph(true)
	vertices := []int{0, 1, 2, 3, 4}
	for _, v := range vertices {
		graph.AddVertex(v)
	}

	edges := [][3]int{
		{0, 1, 5}, {0, 2, 2}, {1, 3, 4}, {1, 4, 3}, {2, 3, 1}, {3, 4, 2},
	}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1], edge[2])
	}

	fmt.Printf("Weighted directed graph: %d vertices, %d edges\n", graph.VertexCount(), graph.EdgeCount())
	fmt.Println("Edges: (0→1:5), (0→2:2), (1→3:4), (1→4:3), (2→3:1), (3→4:2)")

	// Dijkstra's algorithm
	fmt.Println("\n🚀 Dijkstra's Algorithm (Single Source):")
	start := time.Now()
	dijkstraResult, err := graphs.Dijkstra(graph, 0)
	duration := time.Since(start)
	if err != nil {
		fmt.Printf("Dijkstra error: %v\n", err)
	} else {
		fmt.Printf("Shortest distances from vertex 0: (Time: %v)\n", duration)
		for vertex, distance := range dijkstraResult.Distances {
			if parent, exists := dijkstraResult.Parents[vertex]; exists {
				fmt.Printf("  To vertex %d: distance = %d, parent = %d\n", vertex, distance, parent)
			} else {
				fmt.Printf("  To vertex %d: distance = %d, no parent\n", vertex, distance)
			}
		}
	}

	// Single source to target
	fmt.Println("\n🎯 Dijkstra's Algorithm (Source to Target):")
	start = time.Now()
	pathResult, err := graphs.DijkstraPath(graph, 0, 4)
	duration = time.Since(start)
	if err != nil {
		fmt.Printf("Path finding error: %v\n", err)
	} else {
		fmt.Printf("Shortest path 0→4: distance = %d (Time: %v)\n",
			pathResult.Distances[4], duration)
		if len(pathResult.Path) > 0 {
			fmt.Printf("  Path: %v\n", pathResult.Path)
		}
	}

	// Floyd-Warshall algorithm
	fmt.Println("\n🌐 Floyd-Warshall Algorithm (All Pairs):")
	start = time.Now()
	floydResult := graphs.FloydWarshall(graph)
	duration = time.Since(start)
	fmt.Printf("All pairs shortest paths: (Time: %v)\n", duration)
	fmt.Println("Distance matrix:")
	for i := 0; i < len(floydResult.Distances); i++ {
		fmt.Printf("  %v\n", floydResult.Distances[i])
	}
}

func demonstrateMST() {
	fmt.Println("Problem: Find minimum spanning tree in weighted undirected graphs")

	// Create a weighted undirected graph
	//     0
	//   4/ \\2
	//   1---3---2
	//    \\5 /3
	//     \\ /
	//      4
	graph := graphs.NewAdjacencyListGraph(false)
	vertices := []int{0, 1, 2, 3, 4}
	for _, v := range vertices {
		graph.AddVertex(v)
	}

	edges := [][3]int{
		{0, 1, 4}, {0, 2, 2}, {1, 3, 5}, {2, 3, 3}, {1, 4, 5}, {3, 4, 3},
	}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1], edge[2])
	}

	fmt.Printf("Weighted undirected graph: %d vertices, %d edges\n", graph.VertexCount(), graph.EdgeCount())
	fmt.Println("Edges: (0-1:4), (0-2:2), (1-3:5), (2-3:3), (1-4:5), (3-4:3)")

	// Kruskal's algorithm
	fmt.Println("\n🌲 Kruskal's Algorithm (Union-Find):")
	start := time.Now()
	kruskalMST, err := graphs.Kruskal(graph)
	duration := time.Since(start)
	if err != nil {
		fmt.Printf("Kruskal error: %v\n", err)
	} else {
		fmt.Printf("MST total cost: %d (Time: %v)\n", kruskalMST.TotalCost, duration)
		fmt.Printf("MST edges: %v\n", kruskalMST.Edges)
	}

	// Prim's algorithm (simple)
	fmt.Println("\n🌿 Prim's Algorithm (Simple):")
	start = time.Now()
	primMST, err := graphs.Prim(graph, 0)
	duration = time.Since(start)
	if err != nil {
		fmt.Printf("Prim error: %v\n", err)
	} else {
		fmt.Printf("MST total cost: %d (Time: %v)\n", primMST.TotalCost, duration)
		fmt.Printf("MST edges: %v\n", primMST.Edges)
	}

	// Prim's algorithm with priority queue
	fmt.Println("\n🚀 Prim's Algorithm (Priority Queue):")
	start = time.Now()
	primPQMST, err := graphs.PrimWithPriorityQueue(graph, 0)
	duration = time.Since(start)
	if err != nil {
		fmt.Printf("Prim PQ error: %v\n", err)
	} else {
		fmt.Printf("MST total cost: %d (Time: %v)\n", primPQMST.TotalCost, duration)
		fmt.Printf("MST edges: %v\n", primPQMST.Edges)
	}
}

func demonstrateGraphAnalysis() {
	fmt.Println("Problem: Analyze graph properties and detect patterns")

	// Create a directed graph for cycle detection
	graph := graphs.NewAdjacencyListGraph(true)
	vertices := []int{0, 1, 2, 3}
	for _, v := range vertices {
		graph.AddVertex(v)
	}

	edges := [][3]int{
		{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {3, 1, 1}, // Creates cycle: 1→2→3→1
	}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1], edge[2])
	}

	fmt.Printf("Directed graph: %d vertices, %d edges\n", graph.VertexCount(), graph.EdgeCount())
	fmt.Println("Edges: (0→1), (1→2), (2→3), (3→1)")

	// Cycle detection
	fmt.Println("\n🔄 Cycle Detection:")
	start := time.Now()
	hasCycle := graphs.DetectCycle(graph)
	duration := time.Since(start)
	fmt.Printf("Has cycle: %t (Time: %v)\n", hasCycle, duration)

	// Topological sort
	fmt.Println("\n📊 Topological Sort:")
	start = time.Now()
	topoSort, err := graphs.TopologicalSort(graph)
	duration = time.Since(start)
	if err != nil {
		fmt.Printf("Topological sort failed: %v (Time: %v)\n", err, duration)
		fmt.Println("(Expected: graph has cycles)")
	} else {
		fmt.Printf("Topological order: %v (Time: %v)\n", topoSort, duration)
	}

	// Test with acyclic graph
	fmt.Println("\n🌐 Testing with Acyclic Graph:")
	acyclicGraph := graphs.NewAdjacencyListGraph(true)
	for _, v := range vertices {
		acyclicGraph.AddVertex(v)
	}
	acyclicEdges := [][3]int{
		{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {2, 3, 1},
	}
	for _, edge := range acyclicEdges {
		acyclicGraph.AddEdge(edge[0], edge[1], edge[2])
	}

	fmt.Println("Acyclic edges: (0→1), (0→2), (1→3), (2→3)")

	hasCycle = graphs.DetectCycle(acyclicGraph)
	fmt.Printf("Has cycle: %t\n", hasCycle)

	topoSort, err = graphs.TopologicalSort(acyclicGraph)
	if err != nil {
		fmt.Printf("Topological sort error: %v\n", err)
	} else {
		fmt.Printf("Topological order: %v\n", topoSort)
	}

	// Connectivity test
	fmt.Println("\n🔗 Connectivity Analysis:")
	undirectedGraph := graphs.NewAdjacencyListGraph(false)
	for _, v := range vertices {
		undirectedGraph.AddVertex(v)
	}
	for _, edge := range acyclicEdges {
		undirectedGraph.AddEdge(edge[0], edge[1], edge[2])
	}

	isConnected := graphs.IsConnected(undirectedGraph)
	fmt.Printf("Graph is connected: %t\n", isConnected)
}

func demonstrateRealWorldApplications() {
	fmt.Println("Real-world applications of graph algorithms:")

	applications := map[string][]string{
		"DFS": {
			"Maze solving and pathfinding",
			"Topological sorting (dependency resolution)",
			"Cycle detection in workflows",
			"Connected components analysis",
		},
		"BFS": {
			"Shortest path in unweighted graphs",
			"Level-order traversal",
			"Social network analysis (degrees of separation)",
			"Web crawling strategies",
		},
		"Dijkstra": {
			"GPS navigation systems",
			"Network routing protocols",
			"Flight connection optimization",
			"Game AI pathfinding",
		},
		"Floyd-Warshall": {
			"All-pairs shortest paths",
			"Transitive closure computation",
			"Graph reachability analysis",
			"Network analysis and optimization",
		},
		"MST (Kruskal/Prim)": {
			"Network design (minimum cable/pipe laying)",
			"Clustering algorithms",
			"Image segmentation",
			"Circuit design optimization",
		},
	}

	for algo, useCases := range applications {
		fmt.Printf("\n🔹 %s:\n", algo)
		for _, useCase := range useCases {
			fmt.Printf("  • %s\n", useCase)
		}
	}

	fmt.Println("\n💡 Algorithm Selection Guidelines:")
	fmt.Println("✓ Unweighted shortest path:     BFS")
	fmt.Println("✓ Weighted shortest path:       Dijkstra")
	fmt.Println("✓ All pairs shortest path:      Floyd-Warshall")
	fmt.Println("✓ Minimum spanning tree:        Kruskal or Prim")
	fmt.Println("✓ Cycle detection:              DFS")
	fmt.Println("✓ Topological ordering:         DFS or Kahn's algorithm")
	fmt.Println("✓ Connected components:         DFS or BFS")
}
