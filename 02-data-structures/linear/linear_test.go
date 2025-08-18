package datastructures

import (
	"fmt"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		arr := NewDynamicArray(2)

		// Test initial state
		if !arr.IsEmpty() {
			t.Error("New array should be empty")
		}

		if arr.Size() != 0 {
			t.Error("New array size should be 0")
		}

		// Test adding elements
		arr.Add(1)
		arr.Add(2)
		arr.Add(3)

		if arr.Size() != 3 {
			t.Errorf("Expected size 3, got %d", arr.Size())
		}

		// Test capacity growth
		if arr.Capacity() < 3 {
			t.Error("Capacity should have grown to accommodate elements")
		}

		// Test getting elements
		val, err := arr.Get(1)
		if err != nil || val != 2 {
			t.Errorf("Expected 2 at index 1, got %v", val)
		}

		// Test setting elements
		err = arr.Set(1, 20)
		if err != nil {
			t.Error("Set operation failed")
		}

		val, _ = arr.Get(1)
		if val != 20 {
			t.Errorf("Expected 20 after set, got %v", val)
		}
	})

	t.Run("Insert and Remove", func(t *testing.T) {
		arr := NewDynamicArray(5)
		arr.Add(1)
		arr.Add(2)
		arr.Add(4)

		// Test insert
		err := arr.Insert(2, 3)
		if err != nil {
			t.Error("Insert operation failed")
		}

		if arr.Size() != 4 {
			t.Errorf("Expected size 4 after insert, got %d", arr.Size())
		}

		val, _ := arr.Get(2)
		if val != 3 {
			t.Errorf("Expected 3 at index 2 after insert, got %v", val)
		}

		// Test remove
		removed, err := arr.Remove(1)
		if err != nil || removed != 2 {
			t.Errorf("Expected to remove 2, got %v", removed)
		}

		if arr.Size() != 3 {
			t.Errorf("Expected size 3 after remove, got %d", arr.Size())
		}
	})

	t.Run("Edge Cases", func(t *testing.T) {
		arr := NewDynamicArray(1)

		// Test bounds checking
		_, err := arr.Get(0)
		if err == nil {
			t.Error("Should get error when accessing empty array")
		}

		err = arr.Set(0, 1)
		if err == nil {
			t.Error("Should get error when setting in empty array")
		}

		// Test pop on empty array
		_, err = arr.Pop()
		if err == nil {
			t.Error("Should get error when popping empty array")
		}
	})

	t.Run("Iterator", func(t *testing.T) {
		arr := NewDynamicArray(3)
		arr.Add(1)
		arr.Add(2)
		arr.Add(3)

		iter := arr.NewIterator()
		values := []interface{}{}

		for iter.HasNext() {
			val, err := iter.Next()
			if err != nil {
				t.Error("Iterator error:", err)
			}
			values = append(values, val)
		}

		if len(values) != 3 {
			t.Errorf("Expected 3 values from iterator, got %d", len(values))
		}

		for i, expected := range []interface{}{1, 2, 3} {
			if values[i] != expected {
				t.Errorf("Expected %v at position %d, got %v", expected, i, values[i])
			}
		}
	})
}

func TestGenericDynamicArray(t *testing.T) {
	t.Run("Type Safety", func(t *testing.T) {
		intArray := NewGenericDynamicArray[int](2)
		stringArray := NewGenericDynamicArray[string](2)

		// Test int array
		intArray.Add(42)
		intArray.Add(100)

		val, err := intArray.Get(0)
		if err != nil || val != 42 {
			t.Errorf("Expected 42, got %v", val)
		}

		// Test string array
		stringArray.Add("hello")
		stringArray.Add("world")

		strVal, err := stringArray.Get(1)
		if err != nil || strVal != "world" {
			t.Errorf("Expected 'world', got %v", strVal)
		}

		if intArray.Size() != 2 || stringArray.Size() != 2 {
			t.Error("Generic arrays should have correct sizes")
		}
	})
}

func TestStackImplementations(t *testing.T) {
	stacks := []struct {
		name  string
		stack Stack
	}{
		{"ArrayStack", NewArrayStack()},
		{"LinkedStack", NewLinkedStack()},
	}

	for _, test := range stacks {
		t.Run(test.name, func(t *testing.T) {
			s := test.stack

			// Test initial state
			if !s.IsEmpty() {
				t.Error("New stack should be empty")
			}

			if s.Size() != 0 {
				t.Error("New stack size should be 0")
			}

			// Test push operations
			s.Push(1)
			s.Push(2)
			s.Push(3)

			if s.Size() != 3 {
				t.Errorf("Expected size 3, got %d", s.Size())
			}

			if s.IsEmpty() {
				t.Error("Stack should not be empty after pushes")
			}

			// Test peek
			top, err := s.Peek()
			if err != nil || top != 3 {
				t.Errorf("Expected peek to return 3, got %v", top)
			}

			// Stack size should remain the same after peek
			if s.Size() != 3 {
				t.Error("Size should not change after peek")
			}

			// Test pop operations
			val, err := s.Pop()
			if err != nil || val != 3 {
				t.Errorf("Expected pop to return 3, got %v", val)
			}

			val, err = s.Pop()
			if err != nil || val != 2 {
				t.Errorf("Expected pop to return 2, got %v", val)
			}

			if s.Size() != 1 {
				t.Errorf("Expected size 1 after two pops, got %d", s.Size())
			}

			// Test clear
			s.Clear()
			if !s.IsEmpty() || s.Size() != 0 {
				t.Error("Stack should be empty after clear")
			}

			// Test error cases
			_, err = s.Pop()
			if err == nil {
				t.Error("Should get error when popping empty stack")
			}

			_, err = s.Peek()
			if err == nil {
				t.Error("Should get error when peeking empty stack")
			}
		})
	}
}

func TestGenericStack(t *testing.T) {
	t.Run("Type Safety", func(t *testing.T) {
		intStack := NewGenericStack[int]()
		stringStack := NewGenericStack[string]()

		// Test int stack
		intStack.Push(42)
		intStack.Push(100)

		val, err := intStack.Pop()
		if err != nil || val != 100 {
			t.Errorf("Expected 100, got %v", val)
		}

		// Test string stack
		stringStack.Push("hello")
		stringStack.Push("world")

		strVal, err := stringStack.Peek()
		if err != nil || strVal != "world" {
			t.Errorf("Expected 'world', got %v", strVal)
		}

		if intStack.Size() != 1 || stringStack.Size() != 2 {
			t.Error("Generic stacks should have correct sizes")
		}
	})
}

func TestStackApplications(t *testing.T) {
	t.Run("Balanced Parentheses", func(t *testing.T) {
		testCases := []struct {
			expression string
			expected   bool
		}{
			{"()", true},
			{"()[]", true},
			{"()[]{}", true},
			{"([{}])", true},
			{"((()))", true},
			{"(", false},
			{")", false},
			{"([)]", false},
			{"(()", false},
			{"())", false},
			{"", true},
		}

		for _, tc := range testCases {
			result := BalancedParentheses(tc.expression)
			if result != tc.expected {
				t.Errorf("BalancedParentheses(%q) = %v, expected %v",
					tc.expression, result, tc.expected)
			}
		}
	})

	t.Run("Evaluate Postfix", func(t *testing.T) {
		testCases := []struct {
			expression []string
			expected   int
			shouldErr  bool
		}{
			{[]string{"2", "3", "+"}, 5, false},
			{[]string{"2", "3", "-"}, -1, false},
			{[]string{"2", "3", "*"}, 6, false},
			{[]string{"6", "3", "/"}, 2, false},
			{[]string{"15", "7", "1", "1", "+", "-", "*", "3", "+"}, 78, false}, // 15 * (7 - (1 + 1)) + 3 = 15 * 5 + 3 = 78
			{[]string{"2", "+"}, 0, true},                                       // Not enough operands
			{[]string{"2", "0", "/"}, 0, true},                                  // Division by zero
		}

		for _, tc := range testCases {
			result, err := EvaluatePostfix(tc.expression)

			if tc.shouldErr {
				if err == nil {
					t.Errorf("EvaluatePostfix(%v) should have returned an error", tc.expression)
				}
			} else {
				if err != nil {
					t.Errorf("EvaluatePostfix(%v) returned unexpected error: %v", tc.expression, err)
				} else if result != tc.expected {
					t.Errorf("EvaluatePostfix(%v) = %d, expected %d", tc.expression, result, tc.expected)
				}
			}
		}
	})
}

// Benchmarks
func BenchmarkDynamicArrayOperations(b *testing.B) {
	b.Run("Add", func(b *testing.B) {
		arr := NewDynamicArray(1000)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			arr.Add(i)
		}
	})

	b.Run("Get", func(b *testing.B) {
		arr := NewDynamicArray(1000)
		for i := 0; i < 1000; i++ {
			arr.Add(i)
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			arr.Get(i % 1000)
		}
	})

	b.Run("Insert", func(b *testing.B) {
		arr := NewDynamicArray(1000)
		for i := 0; i < 100; i++ {
			arr.Add(i)
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			arr.Insert(0, i) // Worst case: insert at beginning
		}
	})
}

func BenchmarkStackImplementations(b *testing.B) {
	stacks := []struct {
		name  string
		stack Stack
	}{
		{"ArrayStack", NewArrayStack()},
		{"LinkedStack", NewLinkedStack()},
	}

	for _, test := range stacks {
		b.Run(fmt.Sprintf("%s_Push", test.name), func(b *testing.B) {
			s := test.stack
			for i := 0; i < b.N; i++ {
				s.Push(i)
			}
		})

		b.Run(fmt.Sprintf("%s_Pop", test.name), func(b *testing.B) {
			s := test.stack
			// Pre-populate the stack
			for i := 0; i < b.N; i++ {
				s.Push(i)
			}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				s.Pop()
			}
		})
	}
}

func BenchmarkGenericStack(b *testing.B) {
	b.Run("Generic_Push", func(b *testing.B) {
		s := NewGenericStack[int]()
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
	})

	b.Run("Generic_Pop", func(b *testing.B) {
		s := NewGenericStack[int]()
		// Pre-populate the stack
		for i := 0; i < b.N; i++ {
			s.Push(i)
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			s.Pop()
		}
	})
}
