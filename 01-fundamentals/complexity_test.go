package fundamentals

import (
	"fmt"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func assertTrue(t *testing.T, condition bool) {
	if !condition {
		t.Error("Expected true, but got false")
	}
}

func assertFalse(t *testing.T, condition bool) {
	if condition {
		t.Error("Expected false, but got true")
	}
}

func TestConstantTimeOperations(t *testing.T) {
	t.Run("Array Access", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		result := ConstantTimeAccess(arr, 2)
		assertEqual(t, 3, result)

		// Test bounds
		result = ConstantTimeAccess(arr, -1)
		assertEqual(t, -1, result)

		result = ConstantTimeAccess(arr, 10)
		assertEqual(t, -1, result)
	})

	t.Run("Map Lookup", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}

		val, exists := MapLookup(m, "b")
		assertTrue(t, exists)
		assertEqual(t, 2, val)

		val, exists = MapLookup(m, "z")
		assertFalse(t, exists)
		assertEqual(t, 0, val)
	})
}

func TestSearchAlgorithms(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}

	t.Run("Binary Search", func(t *testing.T) {
		// Found cases
		assertEqual(t, 0, BinarySearch(arr, 1))
		assertEqual(t, 4, BinarySearch(arr, 9))
		assertEqual(t, 7, BinarySearch(arr, 15))

		// Not found cases
		assertEqual(t, -1, BinarySearch(arr, 0))
		assertEqual(t, -1, BinarySearch(arr, 8))
		assertEqual(t, -1, BinarySearch(arr, 20))
	})

	t.Run("Linear Search", func(t *testing.T) {
		// Found cases
		assertEqual(t, 0, LinearSearch(arr, 1))
		assertEqual(t, 4, LinearSearch(arr, 9))
		assertEqual(t, 7, LinearSearch(arr, 15))

		// Not found cases
		assertEqual(t, -1, LinearSearch(arr, 0))
		assertEqual(t, -1, LinearSearch(arr, 8))
		assertEqual(t, -1, LinearSearch(arr, 20))
	})
}

func TestSortingAlgorithms(t *testing.T) {
	t.Run("Merge Sort", func(t *testing.T) {
		testCases := []struct {
			input    []int
			expected []int
		}{
			{[]int{}, []int{}},
			{[]int{1}, []int{1}},
			{[]int{3, 1, 4, 1, 5}, []int{1, 1, 3, 4, 5}},
			{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
			{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		}

		for _, tc := range testCases {
			result := MergeSort(tc.input)
			assertEqual(t, tc.expected, result)
		}
	})

	t.Run("Bubble Sort", func(t *testing.T) {
		testCases := []struct {
			input    []int
			expected []int
		}{
			{[]int{}, []int{}},
			{[]int{1}, []int{1}},
			{[]int{3, 1, 4, 1, 5}, []int{1, 1, 3, 4, 5}},
			{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		}

		for _, tc := range testCases {
			// Make a copy since BubbleSort modifies in place
			input := make([]int, len(tc.input))
			copy(input, tc.input)

			BubbleSort(input)
			assertEqual(t, tc.expected, input)
		}
	})
}

func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	t.Run("Naive Fibonacci", func(t *testing.T) {
		for i, exp := range expected {
			if i > 7 { // Skip larger values due to exponential time
				break
			}
			result := Fibonacci(i)
			assertEqual(t, exp, result)
		}
	})

	t.Run("Memoized Fibonacci", func(t *testing.T) {
		for i, exp := range expected {
			result := FibonacciMemo(i)
			assertEqual(t, exp, result)
		}
	})
}

func TestSpaceComplexity(t *testing.T) {
	t.Run("In-place Reverse", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}

		ReverseArrayInPlace(arr)
		assertEqual(t, expected, arr)
	})

	t.Run("New Space Reverse", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}

		result := ReverseArrayNewSpace(arr)
		assertEqual(t, expected, result)
		// Original array should be unchanged
		assertEqual(t, []int{1, 2, 3, 4, 5}, arr)
	})
}

// Benchmarks for complexity analysis
func BenchmarkConstantTime(b *testing.B) {
	arr := GenerateArray(10000)
	m := GenerateMap(10000)

	b.Run("ArrayAccess", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ConstantTimeAccess(arr, 5000)
		}
	})

	b.Run("MapLookup", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MapLookup(m, "key5000")
		}
	})
}

func BenchmarkSearchAlgorithms(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		arr := GenerateArray(size)
		target := size / 2

		b.Run(fmt.Sprintf("LinearSearch_n%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LinearSearch(arr, target)
			}
		})

		b.Run(fmt.Sprintf("BinarySearch_n%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BinarySearch(arr, target)
			}
		})
	}
}

func BenchmarkSortingAlgorithms(b *testing.B) {
	sizes := []int{100, 1000, 5000}

	for _, size := range sizes {
		data := GenerateReverseArray(size)

		b.Run(fmt.Sprintf("MergeSort_n%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := make([]int, len(data))
				copy(arr, data)
				b.StartTimer()

				MergeSort(arr)
			}
		})

		// Skip BubbleSort for larger sizes to avoid long test times
		if size <= 1000 {
			b.Run(fmt.Sprintf("BubbleSort_n%d", size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					arr := make([]int, len(data))
					copy(arr, data)
					b.StartTimer()

					BubbleSort(arr)
				}
			})
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	b.Run("Naive_n20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Fibonacci(20)
		}
	})

	b.Run("Memoized_n20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FibonacciMemo(20)
		}
	})

	b.Run("Memoized_n40", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FibonacciMemo(40)
		}
	})
}

// Helper function for benchmark
func GenerateReverseArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	return arr
}
