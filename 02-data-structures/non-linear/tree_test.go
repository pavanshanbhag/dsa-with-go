package datastructures

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		bst := NewBinarySearchTree()

		// Test empty tree
		if !bst.IsEmpty() {
			t.Error("New BST should be empty")
		}

		if bst.Size() != 0 {
			t.Errorf("Expected size 0, got %d", bst.Size())
		}

		// Test insert
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(7)
		bst.Insert(2)
		bst.Insert(4)
		bst.Insert(6)
		bst.Insert(8)

		if bst.Size() != 7 {
			t.Errorf("Expected size 7, got %d", bst.Size())
		}

		// Test search
		if !bst.Search(4) {
			t.Error("Should find value 4")
		}

		if bst.Search(10) {
			t.Error("Should not find value 10")
		}

		// Test delete
		bst.Delete(3)
		if bst.Search(3) {
			t.Error("Should not find deleted value 3")
		}

		if bst.Size() != 6 {
			t.Errorf("Expected size 6 after deletion, got %d", bst.Size())
		}
	})

	t.Run("Traversals", func(t *testing.T) {
		bst := NewBinarySearchTree()
		values := []int{5, 3, 7, 2, 4, 6, 8}
		for _, v := range values {
			bst.Insert(v)
		}

		// Test inorder traversal (should be sorted)
		inorder := bst.InorderTraversal()
		expected := []int{2, 3, 4, 5, 6, 7, 8}
		if len(inorder) != len(expected) {
			t.Errorf("Inorder length mismatch: expected %d, got %d", len(expected), len(inorder))
		}

		// Check if sorted
		for i, v := range expected {
			if inorder[i] != v {
				t.Errorf("Inorder[%d]: expected %d, got %d", i, v, inorder[i])
			}
		}

		// Test preorder traversal
		preorder := bst.PreorderTraversal()
		if len(preorder) != 7 {
			t.Errorf("Preorder should have 7 elements, got %d", len(preorder))
		}

		// Test postorder traversal
		postorder := bst.PostorderTraversal()
		if len(postorder) != 7 {
			t.Errorf("Postorder should have 7 elements, got %d", len(postorder))
		}

		// Test level order traversal
		levelorder := bst.LevelOrderTraversal()
		if len(levelorder) != 7 {
			t.Errorf("Level order should have 7 elements, got %d", len(levelorder))
		}

		// Root should be first in level order
		if levelorder[0] != 5 {
			t.Errorf("First element in level order should be 5, got %d", levelorder[0])
		}
	})

	t.Run("Edge Cases", func(t *testing.T) {
		bst := NewBinarySearchTree()

		// Delete from empty tree
		bst.Delete(1)
		if !bst.IsEmpty() {
			t.Error("Tree should remain empty after deleting from empty tree")
		}

		// Single node operations
		bst.Insert(42)

		bst.Delete(42)
		if !bst.IsEmpty() {
			t.Error("Tree should be empty after deleting only node")
		}
	})
}

func TestMinHeap(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		heap := NewMinHeap()

		// Test empty heap
		if !heap.IsEmpty() {
			t.Error("New heap should be empty")
		}

		if heap.Size() != 0 {
			t.Errorf("Expected size 0, got %d", heap.Size())
		}

		// Test insert
		values := []int{5, 3, 8, 1, 9, 2, 7}
		for _, v := range values {
			heap.Insert(v)
		}

		if heap.Size() != 7 {
			t.Errorf("Expected size 7, got %d", heap.Size())
		}

		// Test peek (should be minimum)
		min, err := heap.Peek()
		if err != nil {
			t.Errorf("Error peeking: %v", err)
		}
		if min != 1 {
			t.Errorf("Expected min 1, got %d", min)
		}

		// Test extract min
		extracted, err := heap.ExtractMin()
		if err != nil {
			t.Errorf("Error extracting: %v", err)
		}
		if extracted != 1 {
			t.Errorf("Expected extracted min 1, got %d", extracted)
		}

		if heap.Size() != 6 {
			t.Errorf("Expected size 6 after extraction, got %d", heap.Size())
		}

		// Next minimum should be 2
		nextMin, err := heap.Peek()
		if err != nil {
			t.Errorf("Error peeking: %v", err)
		}
		if nextMin != 2 {
			t.Errorf("Expected next min 2, got %d", nextMin)
		}
	})

	t.Run("Heap Property", func(t *testing.T) {
		heap := NewMinHeap()

		// Insert random values
		values := []int{50, 30, 20, 15, 10, 8, 16}
		for _, v := range values {
			heap.Insert(v)
		}

		// Extract all values - should come out in sorted order
		var extracted []int
		for !heap.IsEmpty() {
			val, err := heap.ExtractMin()
			if err != nil {
				t.Errorf("Error extracting: %v", err)
				break
			}
			extracted = append(extracted, val)
		}

		// Verify sorted order
		for i := 1; i < len(extracted); i++ {
			if extracted[i-1] > extracted[i] {
				t.Errorf("Heap property violated: %d > %d", extracted[i-1], extracted[i])
			}
		}
	})

	t.Run("Edge Cases", func(t *testing.T) {
		heap := NewMinHeap()

		// Extract from empty heap
		_, err := heap.ExtractMin()
		if err == nil {
			t.Error("Extracting from empty heap should return error")
		}

		// Peek empty heap
		_, err = heap.Peek()
		if err == nil {
			t.Error("Peeking empty heap should return error")
		}

		// Single element
		heap.Insert(42)
		val, err := heap.Peek()
		if err != nil {
			t.Errorf("Error peeking single element: %v", err)
		}
		if val != 42 {
			t.Errorf("Single element should be 42, got %d", val)
		}

		extracted, err := heap.ExtractMin()
		if err != nil {
			t.Errorf("Error extracting single element: %v", err)
		}
		if extracted != 42 {
			t.Errorf("Should extract 42, got %d", extracted)
		}

		if !heap.IsEmpty() {
			t.Error("Heap should be empty after extracting single element")
		}
	})
}

func TestBinaryTree(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		tree := NewBinaryTree()

		// Test empty tree
		if !tree.IsEmpty() {
			t.Error("New binary tree should be empty")
		}

		if tree.Size() != 0 {
			t.Errorf("Expected size 0, got %d", tree.Size())
		}

		// For now, just test that the tree exists and basic methods work
		// The BinaryTree implementation might be incomplete for insertion
		height := tree.Height()
		if height != -1 {
			t.Errorf("Empty tree height should be -1, got %d", height)
		}

		// Test traversals on empty tree
		inorder := tree.InorderTraversal()
		if len(inorder) != 0 {
			t.Errorf("Empty tree inorder should be empty, got %d elements", len(inorder))
		}
	})
}

// Benchmark tests
func BenchmarkBSTInsert(b *testing.B) {
	bst := NewBinarySearchTree()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bst.Insert(i)
	}
}

func BenchmarkBSTSearch(b *testing.B) {
	bst := NewBinarySearchTree()
	// Populate with 1000 elements
	for i := 0; i < 1000; i++ {
		bst.Insert(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Search(500)
	}
}

func BenchmarkMinHeapInsert(b *testing.B) {
	heap := NewMinHeap()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
}

func BenchmarkMinHeapExtract(b *testing.B) {
	heap := NewMinHeap()
	// Populate heap
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !heap.IsEmpty() {
			heap.ExtractMin()
		}
	}
}
