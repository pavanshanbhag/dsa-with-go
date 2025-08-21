package sorting

import (
	"reflect"
	"sort"
	"testing"
)

// Test data sets for comprehensive testing
var testCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "Empty array",
		input:    []int{},
		expected: []int{},
	},
	{
		name:     "Single element",
		input:    []int{42},
		expected: []int{42},
	},
	{
		name:     "Already sorted",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Reverse sorted",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Random order",
		input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
		expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
	},
	{
		name:     "All duplicates",
		input:    []int{7, 7, 7, 7, 7},
		expected: []int{7, 7, 7, 7, 7},
	},
	{
		name:     "Two elements unsorted",
		input:    []int{2, 1},
		expected: []int{1, 2},
	},
	{
		name:     "Two elements sorted",
		input:    []int{1, 2},
		expected: []int{1, 2},
	},
	{
		name:     "Negative numbers",
		input:    []int{-3, -1, -4, -2, 0, 1},
		expected: []int{-4, -3, -2, -1, 0, 1},
	},
	{
		name:     "Large range",
		input:    []int{100, 1, 999, 50, 25, 750},
		expected: []int{1, 25, 50, 100, 750, 999},
	},
}

// Test all comparison-based sorting algorithms
func TestQuickSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := QuickSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("QuickSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestQuickSortHoare(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := QuickSortHoare(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("QuickSortHoare(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestQuickSort3Way(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := QuickSort3Way(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("QuickSort3Way(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := HeapSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("HeapSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MergeSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("MergeSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := InsertionSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("InsertionSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SelectionSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("SelectionSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BubbleSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("BubbleSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIntroSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IntroSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("IntroSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

// Test non-comparison based sorting algorithms
func TestCountingSort(t *testing.T) {
	// Counting sort only works with non-negative integers
	nonNegativeTestCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "All duplicates",
			input:    []int{7, 7, 7, 7, 7},
			expected: []int{7, 7, 7, 7, 7},
		},
	}

	for _, tc := range nonNegativeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CountingSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("CountingSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestRadixSort(t *testing.T) {
	// Radix sort only works with non-negative integers
	nonNegativeTestCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Multiple digits",
			input:    []int{170, 45, 75, 90, 2, 802, 24, 66},
			expected: []int{2, 24, 45, 66, 75, 90, 170, 802},
		},
		{
			name:     "Same number of digits",
			input:    []int{543, 986, 217, 765, 329},
			expected: []int{217, 329, 543, 765, 986},
		},
	}

	for _, tc := range nonNegativeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RadixSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("RadixSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestBucketSort(t *testing.T) {
	testCases := []struct {
		name        string
		input       []float64
		bucketCount int
		expected    []float64
	}{
		{
			name:        "Empty array",
			input:       []float64{},
			bucketCount: 5,
			expected:    []float64{},
		},
		{
			name:        "Single element",
			input:       []float64{0.5},
			bucketCount: 5,
			expected:    []float64{0.5},
		},
		{
			name:        "Uniform distribution",
			input:       []float64{0.42, 0.32, 0.33, 0.52, 0.37, 0.47, 0.51},
			bucketCount: 5,
			expected:    []float64{0.32, 0.33, 0.37, 0.42, 0.47, 0.51, 0.52},
		},
		{
			name:        "Integer values as floats",
			input:       []float64{3.0, 1.0, 4.0, 1.0, 5.0},
			bucketCount: 3,
			expected:    []float64{1.0, 1.0, 3.0, 4.0, 5.0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BucketSort(tc.input, tc.bucketCount)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("BucketSort(%v, %d) = %v, expected %v", tc.input, tc.bucketCount, result, tc.expected)
			}
		})
	}
}

// Test utility functions
func TestIsSorted(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: true,
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: true,
		},
		{
			name:     "Sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Unsorted array",
			input:    []int{3, 1, 4, 2, 5},
			expected: false,
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			name:     "Duplicates sorted",
			input:    []int{1, 1, 2, 2, 3},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsSorted(tc.input)
			if result != tc.expected {
				t.Errorf("IsSorted(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

// Stress tests with larger arrays
func TestLargeArrays(t *testing.T) {
	sizes := []int{100, 1000, 5000}

	for _, size := range sizes {
		t.Run("Random array size "+string(rune(size)), func(t *testing.T) {
			// Generate random array
			arr := GenerateRandomArray(size)

			// Test a few key algorithms
			algorithms := map[string]func([]int) []int{
				"QuickSort": QuickSort,
				"HeapSort":  HeapSort,
				"MergeSort": MergeSort,
				"IntroSort": IntroSort,
			}

			for name, algorithm := range algorithms {
				t.Run(name, func(t *testing.T) {
					result := algorithm(arr)
					if !IsSorted(result) {
						t.Errorf("%s failed to sort array of size %d", name, size)
					}
					if len(result) != len(arr) {
						t.Errorf("%s changed array length from %d to %d", name, len(arr), len(result))
					}
				})
			}
		})
	}
}

// Test specific scenarios that might cause issues
func TestEdgeCases(t *testing.T) {
	t.Run("Many duplicates", func(t *testing.T) {
		arr := GenerateArrayWithDuplicates(1000, 10)

		algorithms := map[string]func([]int) []int{
			"QuickSort3Way": QuickSort3Way,
			"HeapSort":      HeapSort,
			"MergeSort":     MergeSort,
		}

		for name, algorithm := range algorithms {
			t.Run(name, func(t *testing.T) {
				result := algorithm(arr)
				if !IsSorted(result) {
					t.Errorf("%s failed on array with many duplicates", name)
				}
			})
		}
	})

	t.Run("Nearly sorted", func(t *testing.T) {
		arr := GenerateNearlySortedArray(1000, 10)

		algorithms := map[string]func([]int) []int{
			"InsertionSort": InsertionSort,
			"QuickSort":     QuickSort,
			"IntroSort":     IntroSort,
		}

		for name, algorithm := range algorithms {
			t.Run(name, func(t *testing.T) {
				result := algorithm(arr)
				if !IsSorted(result) {
					t.Errorf("%s failed on nearly sorted array", name)
				}
			})
		}
	})

	t.Run("Reverse sorted", func(t *testing.T) {
		arr := GenerateReverseSortedArray(1000)

		algorithms := map[string]func([]int) []int{
			"QuickSort": QuickSort,
			"HeapSort":  HeapSort,
			"MergeSort": MergeSort,
		}

		for name, algorithm := range algorithms {
			t.Run(name, func(t *testing.T) {
				result := algorithm(arr)
				if !IsSorted(result) {
					t.Errorf("%s failed on reverse sorted array", name)
				}
			})
		}
	})
}

// Benchmark tests for performance comparison
func BenchmarkSortingAlgorithms(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	algorithms := map[string]func([]int) []int{
		"QuickSort":      QuickSort,
		"QuickSortHoare": QuickSortHoare,
		"QuickSort3Way":  QuickSort3Way,
		"HeapSort":       HeapSort,
		"MergeSort":      MergeSort,
		"IntroSort":      IntroSort,
		"InsertionSort":  InsertionSort,
		"SelectionSort":  SelectionSort,
	}

	for _, size := range sizes {
		// Skip expensive algorithms for large arrays
		if size >= 10000 {
			delete(algorithms, "InsertionSort")
			delete(algorithms, "SelectionSort")
			delete(algorithms, "BubbleSort")
		}

		for name, algorithm := range algorithms {
			b.Run(name+"_Random_"+string(rune(size)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					arr := GenerateRandomArray(size)
					b.StartTimer()

					algorithm(arr)
				}
			})
		}
	}
}

func BenchmarkNonComparisonSorts(b *testing.B) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		b.Run("CountingSort_"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := GenerateRandomArray(size)
				// Limit range for counting sort
				for j := range arr {
					arr[j] = arr[j] % 1000
				}
				b.StartTimer()

				CountingSort(arr)
			}
		})

		b.Run("RadixSort_"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := GenerateRandomArray(size)
				b.StartTimer()

				RadixSort(arr)
			}
		})
	}
}

// Benchmark against Go's standard library sort
func BenchmarkVsStandardLibrary(b *testing.B) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		b.Run("GoStdSort_"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := GenerateRandomArray(size)
				b.StartTimer()

				sort.Ints(arr)
			}
		})

		b.Run("IntroSort_"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := GenerateRandomArray(size)
				b.StartTimer()

				IntroSort(arr)
			}
		})

		b.Run("QuickSort_"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := GenerateRandomArray(size)
				b.StartTimer()

				QuickSort(arr)
			}
		})
	}
}

// Example functions for documentation
func ExampleQuickSort() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	sorted := QuickSort(arr)
	_ = sorted
	// Output will be [1, 1, 2, 3, 6, 8, 10]
}

func ExampleHeapSort() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	sorted := HeapSort(arr)
	_ = sorted
	// Output will be [11, 12, 22, 25, 34, 64, 90]
}

func ExampleMergeSort() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	sorted := MergeSort(arr)
	_ = sorted
	// Output will be [3, 9, 10, 27, 38, 43, 82]
}

func ExampleCountingSort() {
	arr := []int{4, 2, 2, 8, 3, 3, 1}
	sorted := CountingSort(arr)
	_ = sorted
	// Output will be [1, 2, 2, 3, 3, 4, 8]
}

func ExampleRadixSort() {
	arr := []int{170, 45, 75, 90, 2, 802, 24, 66}
	sorted := RadixSort(arr)
	_ = sorted
	// Output will be [2, 24, 45, 66, 75, 90, 170, 802]
}
