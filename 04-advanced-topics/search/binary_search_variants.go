package search

import (
	"fmt"
	"math"
)

// BinarySearchResult represents the result of a binary search operation
type BinarySearchResult struct {
	Found bool
	Index int
	Value int
}

// AdvancedBinarySearch provides various binary search implementations
type AdvancedBinarySearch struct{}

// NewAdvancedBinarySearch creates a new instance of advanced binary search algorithms
func NewAdvancedBinarySearch() *AdvancedBinarySearch {
	return &AdvancedBinarySearch{}
}

// ClassicBinarySearch performs standard binary search on sorted array
func (abs *AdvancedBinarySearch) ClassicBinarySearch(arr []int, target int) BinarySearchResult {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return BinarySearchResult{Found: true, Index: mid, Value: arr[mid]}
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return BinarySearchResult{Found: false, Index: -1, Value: -1}
}

// LowerBound finds the first position where target could be inserted
// Returns the index of the first element >= target
func (abs *AdvancedBinarySearch) LowerBound(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// UpperBound finds the last position where target could be inserted
// Returns the index of the first element > target
func (abs *AdvancedBinarySearch) UpperBound(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// SearchRotatedArray searches in a rotated sorted array
func (abs *AdvancedBinarySearch) SearchRotatedArray(arr []int, target int) BinarySearchResult {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return BinarySearchResult{Found: true, Index: mid, Value: arr[mid]}
		}

		// Determine which half is sorted
		if arr[left] <= arr[mid] {
			// Left half is sorted
			if target >= arr[left] && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if target > arr[mid] && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return BinarySearchResult{Found: false, Index: -1, Value: -1}
}

// FindPivotInRotatedArray finds the index of the minimum element (pivot point)
func (abs *AdvancedBinarySearch) FindPivotInRotatedArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] > arr[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// SearchRange finds the range [start, end] of target in sorted array
func (abs *AdvancedBinarySearch) SearchRange(arr []int, target int) [2]int {
	lower := abs.LowerBound(arr, target)
	upper := abs.UpperBound(arr, target)

	if lower < len(arr) && arr[lower] == target {
		return [2]int{lower, upper - 1}
	}

	return [2]int{-1, -1}
}

// BitwiseBinarySearch performs binary search using bitwise operations (for powers of 2)
func (abs *AdvancedBinarySearch) BitwiseBinarySearch(arr []int, target int) BinarySearchResult {
	n := len(arr)
	if n == 0 {
		return BinarySearchResult{Found: false, Index: -1, Value: -1}
	}

	// Find the highest power of 2 <= n
	k := 0
	for (1 << (k + 1)) <= n {
		k++
	}

	pos := 0
	for i := k; i >= 0; i-- {
		if pos+(1<<i) < n && arr[pos+(1<<i)] <= target {
			pos += 1 << i
		}
	}

	if pos < n && arr[pos] == target {
		return BinarySearchResult{Found: true, Index: pos, Value: arr[pos]}
	}

	return BinarySearchResult{Found: false, Index: -1, Value: -1}
}

// ===============================
// Ternary Search Implementation
// ===============================

// TernarySearch finds maximum/minimum of unimodal function
type TernarySearch struct{}

// NewTernarySearch creates a new ternary search instance
func NewTernarySearch() *TernarySearch {
	return &TernarySearch{}
}

// UnimodalFunction represents a function with single peak/valley
type UnimodalFunction func(float64) float64

// FindMaximum finds the maximum point of a unimodal function in [left, right]
func (ts *TernarySearch) FindMaximum(f UnimodalFunction, left, right, eps float64) float64 {
	for right-left > eps {
		mid1 := left + (right-left)/3.0
		mid2 := right - (right-left)/3.0

		if f(mid1) < f(mid2) {
			left = mid1
		} else {
			right = mid2
		}
	}

	return (left + right) / 2.0
}

// FindMinimum finds the minimum point of a unimodal function in [left, right]
func (ts *TernarySearch) FindMinimum(f UnimodalFunction, left, right, eps float64) float64 {
	for right-left > eps {
		mid1 := left + (right-left)/3.0
		mid2 := right - (right-left)/3.0

		if f(mid1) > f(mid2) {
			left = mid1
		} else {
			right = mid2
		}
	}

	return (left + right) / 2.0
}

// TernarySearchDiscrete performs ternary search on discrete array
func (ts *TernarySearch) TernarySearchDiscrete(arr []int, findMax bool) int {
	left, right := 0, len(arr)-1

	for right-left > 2 {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3

		if findMax {
			if arr[mid1] < arr[mid2] {
				left = mid1
			} else {
				right = mid2
			}
		} else {
			if arr[mid1] > arr[mid2] {
				left = mid1
			} else {
				right = mid2
			}
		}
	}

	// Check remaining elements
	bestIdx := left
	for i := left + 1; i <= right; i++ {
		if findMax {
			if arr[i] > arr[bestIdx] {
				bestIdx = i
			}
		} else {
			if arr[i] < arr[bestIdx] {
				bestIdx = i
			}
		}
	}

	return bestIdx
}

// ===============================
// Exponential Search Implementation
// ===============================

// ExponentialSearch searches in unbounded/infinite sorted arrays
type ExponentialSearch struct {
	binarySearch *AdvancedBinarySearch
}

// NewExponentialSearch creates new exponential search instance
func NewExponentialSearch() *ExponentialSearch {
	return &ExponentialSearch{
		binarySearch: NewAdvancedBinarySearch(),
	}
}

// SearchUnbounded searches in an unbounded sorted array
func (es *ExponentialSearch) SearchUnbounded(arr []int, target int) BinarySearchResult {
	if len(arr) == 0 {
		return BinarySearchResult{Found: false, Index: -1, Value: -1}
	}

	if arr[0] == target {
		return BinarySearchResult{Found: true, Index: 0, Value: arr[0]}
	}

	// Find range for binary search by repeated doubling
	bound := 1
	for bound < len(arr) && arr[bound] < target {
		bound *= 2
	}

	// Perform binary search in found range
	left := bound / 2
	right := bound
	if right >= len(arr) {
		right = len(arr) - 1
	}

	for left <= right {
		mid := left + (right-left)/2

		if mid >= len(arr) {
			right = mid - 1
			continue
		}

		if arr[mid] == target {
			return BinarySearchResult{Found: true, Index: mid, Value: arr[mid]}
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return BinarySearchResult{Found: false, Index: -1, Value: -1}
}

// ===============================
// Interpolation Search Implementation
// ===============================

// InterpolationSearch for uniformly distributed sorted arrays
type InterpolationSearch struct{}

// NewInterpolationSearch creates new interpolation search instance
func NewInterpolationSearch() *InterpolationSearch {
	return &InterpolationSearch{}
}

// Search performs interpolation search on uniformly distributed data
func (is *InterpolationSearch) Search(arr []int, target int) BinarySearchResult {
	left, right := 0, len(arr)-1

	for left <= right && target >= arr[left] && target <= arr[right] {
		if left == right {
			if arr[left] == target {
				return BinarySearchResult{Found: true, Index: left, Value: arr[left]}
			}
			break
		}

		// Interpolation formula
		pos := left + int(float64(target-arr[left])/float64(arr[right]-arr[left])*float64(right-left))

		// Ensure pos is within bounds
		if pos < left {
			pos = left
		} else if pos > right {
			pos = right
		}

		if arr[pos] == target {
			return BinarySearchResult{Found: true, Index: pos, Value: arr[pos]}
		} else if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	return BinarySearchResult{Found: false, Index: -1, Value: -1}
}

// ===============================
// Utility Functions
// ===============================

// SearchStatistics provides performance analysis for different search algorithms
type SearchStatistics struct {
	Algorithm   string
	Comparisons int
	Found       bool
	Index       int
}

// BenchmarkSearchAlgorithms compares performance of different search methods
func BenchmarkSearchAlgorithms(arr []int, target int) []SearchStatistics {
	abs := NewAdvancedBinarySearch()
	es := NewExponentialSearch()
	is := NewInterpolationSearch()

	var stats []SearchStatistics

	// Binary Search
	comparisons := 0
	left, right := 0, len(arr)-1
	found := false
	index := -1

	for left <= right {
		comparisons++
		mid := left + (right-left)/2

		if arr[mid] == target {
			found = true
			index = mid
			break
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	stats = append(stats, SearchStatistics{
		Algorithm:   "Binary Search",
		Comparisons: comparisons,
		Found:       found,
		Index:       index,
	})

	// Add other algorithm stats (simplified for brevity)
	result := abs.ClassicBinarySearch(arr, target)
	stats = append(stats, SearchStatistics{
		Algorithm:   "Classic Binary",
		Comparisons: int(math.Log2(float64(len(arr)))) + 1,
		Found:       result.Found,
		Index:       result.Index,
	})

	result = es.SearchUnbounded(arr, target)
	stats = append(stats, SearchStatistics{
		Algorithm:   "Exponential Search",
		Comparisons: int(math.Log2(float64(len(arr)))) + 2,
		Found:       result.Found,
		Index:       result.Index,
	})

	result = is.Search(arr, target)
	stats = append(stats, SearchStatistics{
		Algorithm:   "Interpolation Search",
		Comparisons: int(math.Log2(math.Log2(float64(len(arr))))) + 1,
		Found:       result.Found,
		Index:       result.Index,
	})

	return stats
}

// PrintSearchResults displays search algorithm comparison results
func PrintSearchResults(stats []SearchStatistics, target int) {
	fmt.Printf("Search Results for target: %d\n", target)
	fmt.Println("================================")
	fmt.Printf("%-20s %-12s %-8s %-8s\n", "Algorithm", "Comparisons", "Found", "Index")
	fmt.Println("----------------------------------------------------")

	for _, stat := range stats {
		fmt.Printf("%-20s %-12d %-8t %-8d\n",
			stat.Algorithm, stat.Comparisons, stat.Found, stat.Index)
	}
	fmt.Println()
}
