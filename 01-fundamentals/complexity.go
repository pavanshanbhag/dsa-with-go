package fundamentals

import (
	"fmt"
	"time"
)

// ComplexityAnalyzer provides tools for measuring algorithm performance
type ComplexityAnalyzer struct {
	name string
}

// NewComplexityAnalyzer creates a new analyzer
func NewComplexityAnalyzer(name string) *ComplexityAnalyzer {
	return &ComplexityAnalyzer{name: name}
}

// MeasureTime measures execution time of a function
func (ca *ComplexityAnalyzer) MeasureTime(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// O(1) - Constant Time Examples
func ConstantTimeAccess(arr []int, index int) int {
	if index < 0 || index >= len(arr) {
		return -1
	}
	return arr[index]
}

func MapLookup(m map[string]int, key string) (int, bool) {
	val, exists := m[key]
	return val, exists
}

// O(log n) - Logarithmic Time Examples
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// O(n) - Linear Time Examples
func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

func SumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// O(n log n) - Linearithmic Time Examples
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// O(n²) - Quadratic Time Examples
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func FindPairs(arr []int) [][]int {
	var pairs [][]int
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, []int{arr[i], arr[j]})
		}
	}

	return pairs
}

// O(2^n) - Exponential Time Example
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Optimized O(n) Fibonacci with memoization
func FibonacciMemo(n int) int {
	memo := make(map[int]int)
	return fibHelper(n, memo)
}

func fibHelper(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = fibHelper(n-1, memo) + fibHelper(n-2, memo)
	return memo[n]
}

// Space Complexity Examples

// O(1) Space - In-place operations
func ReverseArrayInPlace(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// O(n) Space - Creating new array
func ReverseArrayNewSpace(arr []int) []int {
	result := make([]int, len(arr))
	for i, val := range arr {
		result[len(arr)-1-i] = val
	}
	return result
}

// Utility functions for testing
func GenerateArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

func GenerateMap(size int) map[string]int {
	m := make(map[string]int, size)
	for i := 0; i < size; i++ {
		m[fmt.Sprintf("key%d", i)] = i
	}
	return m
}

// Demo function to show complexity differences
func DemoComplexityDifferences() {
	fmt.Println("=== Complexity Analysis Demo ===")
	fmt.Println()

	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		fmt.Printf("Array size: %d\n", size)
		arr := GenerateArray(size)
		target := size - 1 // Last element

		// Linear search
		analyzer := NewComplexityAnalyzer("Linear Search")
		duration := analyzer.MeasureTime(func() {
			LinearSearch(arr, target)
		})
		fmt.Printf("Linear Search: %v\n", duration)

		// Binary search (array is already sorted)
		duration = analyzer.MeasureTime(func() {
			BinarySearch(arr, target)
		})
		fmt.Printf("Binary Search: %v\n", duration)

		fmt.Println()
	}
}
