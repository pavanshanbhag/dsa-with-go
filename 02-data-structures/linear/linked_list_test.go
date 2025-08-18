package datastructures

import (
	"testing"
)

func TestSinglyLinkedListBasicOps(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		list := NewSinglyLinkedList()

		// Test empty list
		if !list.IsEmpty() {
			t.Error("New list should be empty")
		}

		if list.Size() != 0 {
			t.Errorf("Expected size 0, got %d", list.Size())
		}

		// Test prepend
		list.PrependFirst(1)
		list.PrependFirst(2)
		list.PrependFirst(3)

		if list.Size() != 3 {
			t.Errorf("Expected size 3, got %d", list.Size())
		}

		// Test append
		list.AppendLast(0)

		if list.Size() != 4 {
			t.Errorf("Expected size 4, got %d", list.Size())
		}

		// Test contains
		if !list.Contains(2) {
			t.Error("Should find value 2")
		}

		if list.Contains(10) {
			t.Error("Should not find value 10")
		}

		// Test get by index
		val, err := list.Get(0)
		if err != nil {
			t.Errorf("Error getting element: %v", err)
		}
		if val != 3 {
			t.Errorf("Expected first element to be 3, got %v", val)
		}

		// Test indexOf
		index := list.IndexOf(2)
		if index == -1 {
			t.Error("Should find index of value 2")
		}

		// Test remove
		removed, err := list.RemoveFirst()
		if err != nil {
			t.Errorf("Error removing first: %v", err)
		}
		if removed != 3 {
			t.Errorf("Expected removed element to be 3, got %v", removed)
		}

		if list.Size() != 3 {
			t.Errorf("Expected size 3 after removal, got %d", list.Size())
		}
	})

	t.Run("Edge Cases", func(t *testing.T) {
		list := NewSinglyLinkedList()

		// Remove from empty list
		_, err := list.RemoveFirst()
		if err == nil {
			t.Error("Should get error when removing from empty list")
		}

		// Get from empty list
		_, err = list.Get(0)
		if err == nil {
			t.Error("Should get error when accessing empty list")
		}

		// Insert/remove single element
		list.AppendLast(42)
		removed, err := list.RemoveFirst()
		if err != nil {
			t.Errorf("Error removing: %v", err)
		}
		if removed != 42 {
			t.Errorf("Expected 42, got %v", removed)
		}
		if !list.IsEmpty() {
			t.Error("List should be empty after removing only element")
		}
	})

	t.Run("String Representation", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AppendLast(1)
		list.AppendLast(2)
		list.AppendLast(3)

		str := list.String()
		if str == "" {
			t.Error("String representation should not be empty")
		}
		t.Logf("List string representation: %s", str)
	})
}

func TestDoublyLinkedListBasicOps(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		list := NewDoublyLinkedList()

		// Test empty
		if !list.IsEmpty() {
			t.Error("New list should be empty")
		}

		// Test append
		list.AppendLast("first")
		list.AppendLast("second")
		list.AppendLast("third")

		if list.Size() != 3 {
			t.Errorf("Expected size 3, got %d", list.Size())
		}

		// Test prepend
		list.PrependFirst("zero")

		if list.Size() != 4 {
			t.Errorf("Expected size 4, got %d", list.Size())
		}

		// Test get
		val, err := list.Get(0)
		if err != nil {
			t.Errorf("Error getting element: %v", err)
		}
		if val != "zero" {
			t.Errorf("Expected 'zero', got %v", val)
		}

		// Test remove
		removed, err := list.RemoveFirst()
		if err != nil {
			t.Errorf("Error removing: %v", err)
		}
		if removed != "zero" {
			t.Errorf("Expected 'zero', got %v", removed)
		}
	})

	t.Run("String Representation", func(t *testing.T) {
		list := NewDoublyLinkedList()
		list.AppendLast(1)
		list.AppendLast(2)
		list.AppendLast(3)

		str := list.String()
		if str == "" {
			t.Error("String representation should not be empty")
		}
		t.Logf("Doubly linked list: %s", str)
	})
}

func TestCircularLinkedListBasicOps(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		list := NewCircularLinkedList()

		// Test empty
		if !list.IsEmpty() {
			t.Error("New list should be empty")
		}

		// Test add
		list.Add(1)
		list.Add(2)
		list.Add(3)

		if list.Size() != 3 {
			t.Errorf("Expected size 3, got %d", list.Size())
		}

		// Test remove
		removed := list.Remove(2)
		if !removed {
			t.Error("Should successfully remove value 2")
		}

		if list.Size() != 2 {
			t.Errorf("Expected size 2 after deletion, got %d", list.Size())
		}

		// Try to remove non-existent value
		removed = list.Remove(10)
		if removed {
			t.Error("Should not remove non-existent value")
		}
	})

	t.Run("String Representation", func(t *testing.T) {
		list := NewCircularLinkedList()
		list.Add("A")
		list.Add("B")
		list.Add("C")

		str := list.String()
		if str == "" {
			t.Error("String representation should not be empty")
		}
		t.Logf("Circular linked list: %s", str)
	})
}

// Benchmark tests for performance comparison
func BenchmarkSinglyLinkedListAppend(b *testing.B) {
	list := NewSinglyLinkedList()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.AppendLast(i)
	}
}

func BenchmarkSinglyLinkedListPrepend(b *testing.B) {
	list := NewSinglyLinkedList()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.PrependFirst(i)
	}
}

func BenchmarkSinglyLinkedListContains(b *testing.B) {
	list := NewSinglyLinkedList()
	// Populate with 1000 elements
	for i := 0; i < 1000; i++ {
		list.AppendLast(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Contains(500) // Search for middle element
	}
}

func BenchmarkDoublyLinkedListAppend(b *testing.B) {
	list := NewDoublyLinkedList()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.AppendLast(i)
	}
}

func BenchmarkCircularLinkedListAppend(b *testing.B) {
	list := NewCircularLinkedList()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.Add(i)
	}
}
