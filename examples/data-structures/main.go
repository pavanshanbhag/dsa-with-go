package main

import (
	"fmt"
	"time"

	linear "dsa-mastery/02-data-structures/linear"
	nonlinear "dsa-mastery/02-data-structures/non-linear"
)

func main() {
	fmt.Println("🏗️  DSA Mastery - Data Structures Demonstration")
	fmt.Println("==============================================")

	// 1. Linear Data Structures
	fmt.Println("\n📊 1. Linear Data Structures")
	fmt.Println("----------------------------")
	demonstrateLinearStructures()

	// 2. Non-Linear Data Structures
	fmt.Println("\n🌳 2. Non-Linear Data Structures")
	fmt.Println("--------------------------------")
	demonstrateNonLinearStructures()

	// 3. Real-World Applications
	fmt.Println("\n🌟 3. Real-World Applications")
	fmt.Println("-----------------------------")
	demonstrateRealWorldApplications()

	fmt.Println("\n✅ All data structures demonstrated successfully!")
	fmt.Println("🚀 Ready to build efficient applications!")
}

func demonstrateLinearStructures() {
	fmt.Println("Problem: Demonstrate fundamental linear data structures")

	// Dynamic Array
	fmt.Println("\n📈 Dynamic Array:")
	arr := linear.NewDynamicArray(5)

	start := time.Now()
	for i := 0; i < 10; i++ {
		arr.Add(i * 10)
	}
	duration := time.Since(start)

	fmt.Printf("Added 10 elements: %v (Time: %v)\n", arr.ToSlice(), duration)
	fmt.Printf("Size: %d, Capacity: %d\n", arr.Size(), arr.Capacity())

	// Access and modify
	if arr.Size() > 5 {
		elem, err := arr.Get(5)
		if err != nil {
			fmt.Printf("Error getting element: %v\n", err)
		} else {
			fmt.Printf("Element at index 5: %v\n", elem)
		}
		arr.Set(5, 999)
		fmt.Printf("After setting index 5 to 999: %v\n", arr.ToSlice())
	}

	// Stack Operations
	fmt.Println("\n📚 Stack (LIFO - Last In, First Out):")
	stack := linear.NewArrayStack()

	// Push operations
	elements := []int{10, 20, 30, 40, 50}
	start = time.Now()
	for _, elem := range elements {
		stack.Push(elem)
	}
	duration = time.Since(start)
	fmt.Printf("Pushed %v: Size = %d (Time: %v)\n", elements, stack.Size(), duration)

	// Peek and pop
	if !stack.IsEmpty() {
		top, err := stack.Peek()
		if err == nil {
			fmt.Printf("Top element: %v\n", top)
		}

		popped, err := stack.Pop()
		if err == nil {
			fmt.Printf("Popped element: %v, remaining size: %d\n", popped, stack.Size())
		}
	}

	// Queue Operations
	fmt.Println("\n🚌 Queue (FIFO - First In, First Out):")
	queue := linear.NewDynamicArrayQueue()

	// Enqueue operations
	start = time.Now()
	for _, elem := range elements {
		queue.Enqueue(elem)
	}
	duration = time.Since(start)
	fmt.Printf("Enqueued %v: Size = %d (Time: %v)\n", elements, queue.Size(), duration)

	// Front and dequeue
	if !queue.IsEmpty() {
		front, err := queue.Front()
		if err == nil {
			fmt.Printf("Front element: %v\n", front)
		}

		dequeued, err := queue.Dequeue()
		if err == nil {
			fmt.Printf("Dequeued element: %v, remaining size: %d\n", dequeued, queue.Size())
		}
	}

	// Linked List Operations
	fmt.Println("\n🔗 Singly Linked List:")
	linkedList := linear.NewSinglyLinkedList()

	// Add elements
	start = time.Now()
	for _, elem := range elements {
		linkedList.AppendLast(elem)
	}
	duration = time.Since(start)
	fmt.Printf("Appended %v: %s (Time: %v)\n", elements, linkedList.String(), duration)

	// Insert at specific position
	linkedList.Insert(2, 999)
	fmt.Printf("After inserting 999 at index 2: %s\n", linkedList.String())

	// Remove element
	removed, err := linkedList.Remove(2)
	if err != nil {
		fmt.Printf("Remove error: %v\n", err)
	} else {
		fmt.Printf("Removed element: %v, list: %s\n", removed, linkedList.String())
	}
}

func demonstrateNonLinearStructures() {
	fmt.Println("Problem: Demonstrate tree-based data structures")

	// Binary Search Tree
	fmt.Println("\n🌲 Binary Search Tree:")
	bst := nonlinear.NewBinarySearchTree()

	// Insert elements
	elements := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	start := time.Now()
	for _, elem := range elements {
		bst.Insert(elem)
	}
	duration := time.Since(start)
	fmt.Printf("Inserted %v (Time: %v)\n", elements, duration)
	fmt.Printf("Tree size: %d, height: %d\n", bst.Size(), bst.Height())

	// Tree traversals
	fmt.Println("\n🚶 Tree Traversals:")

	fmt.Printf("In-order (sorted):   %v\n", bst.InorderTraversal())
	fmt.Printf("Pre-order:           %v\n", bst.PreorderTraversal())
	fmt.Printf("Post-order:          %v\n", bst.PostorderTraversal())
	fmt.Printf("Level-order (BFS):   %v\n", bst.LevelOrderTraversal())

	// Search operations
	fmt.Println("\n🔍 Search Operations:")
	searchValues := []int{45, 100, 25}
	for _, val := range searchValues {
		start = time.Now()
		found := bst.Search(val)
		duration = time.Since(start)
		fmt.Printf("Search %d: %t (Time: %v)\n", val, found, duration)
	}

	// Min and Max
	if !bst.IsEmpty() {
		min, err := bst.FindMin()
		if err == nil {
			fmt.Printf("Minimum value: %v\n", min)
		}
		max, err := bst.FindMax()
		if err == nil {
			fmt.Printf("Maximum value: %v\n", max)
		}
	}

	// Delete operations
	fmt.Println("\n🗑️  Delete Operations:")
	deleteValues := []int{20, 30, 50}
	for _, val := range deleteValues {
		if bst.Search(val) {
			bst.Delete(val)
			fmt.Printf("Deleted %d: In-order = %v\n", val, bst.InorderTraversal())
		}
	}
}

func demonstrateRealWorldApplications() {
	fmt.Println("Real-world applications of different data structures:")

	applications := map[string][]string{
		"Dynamic Array": {
			"Application backends (user lists, data collections)",
			"Game development (sprite lists, inventory systems)",
			"Image processing (pixel arrays, filters)",
			"Database indexes and caching",
		},
		"Stack": {
			"Function call management (call stack)",
			"Undo operations in applications",
			"Expression evaluation and parsing",
			"Browser back button functionality",
			"Recursive algorithm implementation",
		},
		"Queue": {
			"Task scheduling in operating systems",
			"Breadth-first search algorithms",
			"Printer job management",
			"Network packet handling",
			"Event-driven programming",
		},
		"Linked List": {
			"Music playlist implementation",
			"Browser history navigation",
			"Implementing other data structures",
			"Memory-efficient data storage",
			"Real-time data streaming",
		},
		"Binary Search Tree": {
			"Database indexing systems",
			"File system directory structures",
			"Auto-complete and search suggestions",
			"Priority-based task management",
			"Expression parsing and evaluation",
		},
	}

	for structure, useCases := range applications {
		fmt.Printf("\n🔹 %s:\n", structure)
		for _, useCase := range useCases {
			fmt.Printf("  • %s\n", useCase)
		}
	}

	fmt.Println("\n💡 Data Structure Selection Guidelines:")
	fmt.Println("✓ Random access needed:        Dynamic Array")
	fmt.Println("✓ Frequent insertions at ends: Dynamic Array or Linked List")
	fmt.Println("✓ LIFO access pattern:         Stack")
	fmt.Println("✓ FIFO access pattern:         Queue")
	fmt.Println("✓ Sorted data with searches:   Binary Search Tree")
	fmt.Println("✓ Memory efficiency:           Linked List")
	fmt.Println("✓ Cache-friendly access:       Dynamic Array")
}
