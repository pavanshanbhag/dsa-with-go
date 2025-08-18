package datastructures

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// Helper function for testing
func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestQueueImplementations(t *testing.T) {
	queues := []struct {
		name  string
		queue Queue
	}{
		{"ArrayQueue", NewArrayQueue(5)},
		{"DynamicArrayQueue", NewDynamicArrayQueue()},
		{"LinkedQueue", NewLinkedQueue()},
	}

	for _, test := range queues {
		t.Run(test.name, func(t *testing.T) {
			q := test.queue

			// Test initial state
			if !q.IsEmpty() {
				t.Error("New queue should be empty")
			}

			if q.Size() != 0 {
				t.Error("New queue size should be 0")
			}

			// Test enqueue operations
			err := q.Enqueue(1)
			if err != nil {
				t.Error("Enqueue failed:", err)
			}

			err = q.Enqueue(2)
			if err != nil {
				t.Error("Enqueue failed:", err)
			}

			err = q.Enqueue(3)
			if err != nil {
				t.Error("Enqueue failed:", err)
			}

			if q.Size() != 3 {
				t.Errorf("Expected size 3, got %d", q.Size())
			}

			if q.IsEmpty() {
				t.Error("Queue should not be empty after enqueues")
			}

			// Test front
			front, err := q.Front()
			if err != nil || front != 1 {
				t.Errorf("Expected front to be 1, got %v", front)
			}

			// Size should remain same after front
			if q.Size() != 3 {
				t.Error("Size should not change after front operation")
			}

			// Test dequeue operations
			val, err := q.Dequeue()
			if err != nil || val != 1 {
				t.Errorf("Expected dequeue to return 1, got %v", val)
			}

			val, err = q.Dequeue()
			if err != nil || val != 2 {
				t.Errorf("Expected dequeue to return 2, got %v", val)
			}

			if q.Size() != 1 {
				t.Errorf("Expected size 1 after two dequeues, got %d", q.Size())
			}

			// Test clear
			q.Clear()
			if !q.IsEmpty() || q.Size() != 0 {
				t.Error("Queue should be empty after clear")
			}

			// Test error cases
			_, err = q.Dequeue()
			if err == nil {
				t.Error("Should get error when dequeuing empty queue")
			}

			_, err = q.Front()
			if err == nil {
				t.Error("Should get error when accessing front of empty queue")
			}
		})
	}
}

func TestArrayQueueCapacity(t *testing.T) {
	q := NewArrayQueue(2)

	// Fill to capacity
	q.Enqueue(1)
	q.Enqueue(2)

	// Should fail when full
	err := q.Enqueue(3)
	if err == nil {
		t.Error("Should get error when enqueuing to full queue")
	}

	// Test circular buffer behavior
	q.Dequeue() // Remove 1

	// Should be able to add again
	err = q.Enqueue(3)
	if err != nil {
		t.Error("Should be able to enqueue after dequeue:", err)
	}

	// Verify order: should be 2, 3
	val, _ := q.Dequeue()
	if val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}

	val, _ = q.Dequeue()
	if val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()

	// Test initial state
	if !pq.IsEmpty() {
		t.Error("New priority queue should be empty")
	}

	// Add items with different priorities
	pq.Enqueue("low", 1)
	pq.Enqueue("high", 10)
	pq.Enqueue("medium", 5)
	pq.Enqueue("highest", 15)

	if pq.Size() != 4 {
		t.Errorf("Expected size 4, got %d", pq.Size())
	}

	// Should dequeue in priority order (highest first)
	val, err := pq.Dequeue()
	if err != nil || val != "highest" {
		t.Errorf("Expected 'highest', got %v", val)
	}

	val, err = pq.Dequeue()
	if err != nil || val != "high" {
		t.Errorf("Expected 'high', got %v", val)
	}

	val, err = pq.Dequeue()
	if err != nil || val != "medium" {
		t.Errorf("Expected 'medium', got %v", val)
	}

	val, err = pq.Dequeue()
	if err != nil || val != "low" {
		t.Errorf("Expected 'low', got %v", val)
	}

	if !pq.IsEmpty() {
		t.Error("Priority queue should be empty after all dequeues")
	}
}

func TestChannelQueue(t *testing.T) {
	cq := NewChannelQueue(2)

	// Test non-blocking operations
	err := cq.Enqueue(1)
	if err != nil {
		t.Error("Enqueue failed:", err)
	}

	err = cq.Enqueue(2)
	if err != nil {
		t.Error("Enqueue failed:", err)
	}

	// Should fail when full
	err = cq.Enqueue(3)
	if err == nil {
		t.Error("Should fail when enqueuing to full channel queue")
	}

	// Test dequeue
	val, err := cq.Dequeue()
	if err != nil || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}

	// Should be able to enqueue again
	err = cq.Enqueue(3)
	if err != nil {
		t.Error("Should be able to enqueue after dequeue:", err)
	}

	if cq.Size() != 2 {
		t.Errorf("Expected size 2, got %d", cq.Size())
	}
}

func TestGenericQueue(t *testing.T) {
	intQueue := NewGenericQueue[int]()
	stringQueue := NewGenericQueue[string]()

	// Test int queue
	intQueue.Enqueue(42)
	intQueue.Enqueue(100)

	val, err := intQueue.Dequeue()
	if err != nil || val != 42 {
		t.Errorf("Expected 42, got %v", val)
	}

	// Test string queue
	stringQueue.Enqueue("hello")
	stringQueue.Enqueue("world")

	strVal, err := stringQueue.Front()
	if err != nil || strVal != "hello" {
		t.Errorf("Expected 'hello', got %v", strVal)
	}

	if intQueue.Size() != 1 || stringQueue.Size() != 2 {
		t.Error("Generic queues should have correct sizes")
	}
}

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()

	// Test initial state
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	if list.Size() != 0 {
		t.Error("New list size should be 0")
	}

	// Test append operations
	list.AppendLast(1)
	list.AppendLast(2)
	list.AppendLast(3)

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}

	// Test get operations
	val, err := list.Get(0)
	if err != nil || val != 1 {
		t.Errorf("Expected 1 at index 0, got %v", val)
	}

	val, err = list.Get(2)
	if err != nil || val != 3 {
		t.Errorf("Expected 3 at index 2, got %v", val)
	}

	// Test prepend
	list.PrependFirst(0)

	if list.Size() != 4 {
		t.Errorf("Expected size 4 after prepend, got %d", list.Size())
	}

	val, _ = list.Get(0)
	if val != 0 {
		t.Errorf("Expected 0 at index 0 after prepend, got %v", val)
	}

	// Test insert
	err = list.Insert(2, 1.5)
	if err != nil {
		t.Error("Insert failed:", err)
	}

	val, _ = list.Get(2)
	if val != 1.5 {
		t.Errorf("Expected 1.5 at index 2 after insert, got %v", val)
	}

	// Test remove
	removed, err := list.Remove(2)
	if err != nil || removed != 1.5 {
		t.Errorf("Expected to remove 1.5, got %v", removed)
	}

	// Test contains and indexOf
	if !list.Contains(2) {
		t.Error("List should contain 2")
	}

	index := list.IndexOf(2)
	if index != 2 {
		t.Errorf("Expected index 2 for value 2, got %d", index)
	}

	// Test edge cases
	_, err = list.Get(-1)
	if err == nil {
		t.Error("Should get error for negative index")
	}

	_, err = list.Get(10)
	if err == nil {
		t.Error("Should get error for index out of bounds")
	}
}

func TestDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList()

	// Test append and prepend
	list.AppendLast(2)
	list.AppendLast(3)
	list.PrependFirst(1)

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}

	// Test bidirectional access optimization
	// Should access from head for index 0
	val, err := list.Get(0)
	if err != nil || val != 1 {
		t.Errorf("Expected 1 at index 0, got %v", val)
	}

	// Should access from tail for index 2
	val, err = list.Get(2)
	if err != nil || val != 3 {
		t.Errorf("Expected 3 at index 2, got %v", val)
	}

	// Test remove operations
	removed, err := list.RemoveFirst()
	if err != nil || removed != 1 {
		t.Errorf("Expected to remove 1, got %v", removed)
	}

	removed, err = list.RemoveLast()
	if err != nil || removed != 3 {
		t.Errorf("Expected to remove 3, got %v", removed)
	}

	if list.Size() != 1 {
		t.Errorf("Expected size 1 after removes, got %d", list.Size())
	}
}

func TestCircularLinkedList(t *testing.T) {
	list := NewCircularLinkedList()

	// Test add operations
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}

	// Test remove
	removed := list.Remove(2)
	if !removed {
		t.Error("Should have removed element 2")
	}

	if list.Size() != 2 {
		t.Errorf("Expected size 2 after remove, got %d", list.Size())
	}

	// Test remove non-existent
	removed = list.Remove(10)
	if removed {
		t.Error("Should not have removed non-existent element")
	}
}

func TestGenericLinkedList(t *testing.T) {
	intList := NewGenericLinkedList[int]()
	stringList := NewGenericLinkedList[string]()

	// Test int list
	intList.Append(42)
	intList.Append(100)

	val, err := intList.Get(0)
	if err != nil || val != 42 {
		t.Errorf("Expected 42, got %v", val)
	}

	// Test string list
	stringList.Append("hello")
	stringList.Append("world")

	strVal, err := stringList.Get(1)
	if err != nil || strVal != "world" {
		t.Errorf("Expected 'world', got %v", strVal)
	}

	// Test to slice
	intSlice := intList.ToSlice()
	expected := []int{42, 100}
	if !reflect.DeepEqual(intSlice, expected) {
		t.Errorf("Expected %v, got %v", expected, intSlice)
	}
}

func TestLinkedListApplications(t *testing.T) {
	t.Run("Reverse List", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AppendLast(1)
		list.AppendLast(2)
		list.AppendLast(3)

		list.Reverse()

		// Check if reversed
		val, _ := list.Get(0)
		if val != 3 {
			t.Errorf("Expected 3 at index 0 after reverse, got %v", val)
		}

		val, _ = list.Get(2)
		if val != 1 {
			t.Errorf("Expected 1 at index 2 after reverse, got %v", val)
		}
	})

	t.Run("Find Middle", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AppendLast(1)
		list.AppendLast(2)
		list.AppendLast(3)
		list.AppendLast(4)
		list.AppendLast(5)

		middle := list.FindMiddle()
		if middle == nil || middle.Data != 3 {
			t.Errorf("Expected middle to be 3, got %v", middle.Data)
		}
	})

	t.Run("Cycle Detection", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AppendLast(1)
		list.AppendLast(2)
		list.AppendLast(3)

		// No cycle initially
		if list.HasCycle() {
			t.Error("List should not have cycle initially")
		}

		// Create a cycle manually for testing
		// Note: This is for testing purposes only
		if list.tail != nil {
			list.tail.Next = list.head // Create cycle
		}

		if !list.HasCycle() {
			t.Error("List should detect cycle")
		}
	})
}

// Benchmarks for queue implementations
func BenchmarkQueueOperations(b *testing.B) {
	queues := []struct {
		name  string
		queue Queue
	}{
		{"ArrayQueue", NewArrayQueue(10000)},
		{"DynamicArrayQueue", NewDynamicArrayQueue()},
		{"LinkedQueue", NewLinkedQueue()},
	}

	for _, test := range queues {
		b.Run(fmt.Sprintf("%s_Enqueue", test.name), func(b *testing.B) {
			q := test.queue
			for i := 0; i < b.N; i++ {
				q.Enqueue(i)
			}
		})

		b.Run(fmt.Sprintf("%s_Dequeue", test.name), func(b *testing.B) {
			q := test.queue
			// Pre-populate
			for i := 0; i < b.N; i++ {
				q.Enqueue(i)
			}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				q.Dequeue()
			}
		})
	}
}

func BenchmarkLinkedListOperations(b *testing.B) {
	b.Run("SinglyLinkedList_Append", func(b *testing.B) {
		list := NewSinglyLinkedList()
		for i := 0; i < b.N; i++ {
			list.AppendLast(i)
		}
	})

	b.Run("DoublyLinkedList_Append", func(b *testing.B) {
		list := NewDoublyLinkedList()
		for i := 0; i < b.N; i++ {
			list.AppendLast(i)
		}
	})

	b.Run("SinglyLinkedList_Get", func(b *testing.B) {
		list := NewSinglyLinkedList()
		for i := 0; i < 1000; i++ {
			list.AppendLast(i)
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			list.Get(i % 1000)
		}
	})

	b.Run("DoublyLinkedList_Get", func(b *testing.B) {
		list := NewDoublyLinkedList()
		for i := 0; i < 1000; i++ {
			list.AppendLast(i)
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			list.Get(i % 1000)
		}
	})
}

func BenchmarkChannelQueueVsTraditional(b *testing.B) {
	b.Run("ChannelQueue", func(b *testing.B) {
		cq := NewChannelQueue(1000)

		// Test concurrent operations
		go func() {
			for i := 0; i < b.N; i++ {
				cq.EnqueueBlocking(i)
			}
		}()

		for i := 0; i < b.N; i++ {
			cq.DequeueBlocking()
		}
	})

	b.Run("SafeQueue", func(b *testing.B) {
		sq := NewSafeQueue(NewLinkedQueue())

		// Test concurrent operations
		go func() {
			for i := 0; i < b.N; i++ {
				sq.Enqueue(i)
			}
		}()

		for i := 0; i < b.N; i++ {
			for {
				if _, err := sq.Dequeue(); err == nil {
					break
				}
				time.Sleep(1 * time.Microsecond)
			}
		}
	})
}
