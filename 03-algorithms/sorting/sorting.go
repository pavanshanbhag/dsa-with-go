package sorting

import (
	"math/rand"
	"time"
)

// SortingResult contains metrics about the sorting operation
type SortingResult struct {
	Algorithm   string
	InputSize   int
	Comparisons int64
	Swaps       int64
	Duration    time.Duration
}

// SortingStats tracks operation counts during sorting
type SortingStats struct {
	Comparisons int64
	Swaps       int64
}

// =============================================================================
// COMPARISON-BASED SORTING ALGORITHMS
// =============================================================================

// QuickSort implements the QuickSort algorithm with Lomuto partition
// Time: O(n log n) average, O(n²) worst case | Space: O(log n)
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	stats := &SortingStats{}
	quickSortHelper(result, 0, len(result)-1, stats)
	return result
}

func quickSortHelper(arr []int, low, high int, stats *SortingStats) {
	if low < high {
		stats.Comparisons++
		pivotIndex := lomutoPartition(arr, low, high, stats)
		quickSortHelper(arr, low, pivotIndex-1, stats)
		quickSortHelper(arr, pivotIndex+1, high, stats)
	}
}

func lomutoPartition(arr []int, low, high int, stats *SortingStats) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		stats.Comparisons++
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			stats.Swaps++
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	stats.Swaps++
	return i + 1
}

// QuickSortHoare implements QuickSort with Hoare partition scheme
// Generally more efficient than Lomuto partition
func QuickSortHoare(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	stats := &SortingStats{}
	quickSortHoareHelper(result, 0, len(result)-1, stats)
	return result
}

func quickSortHoareHelper(arr []int, low, high int, stats *SortingStats) {
	if low < high {
		stats.Comparisons++
		pivotIndex := hoarePartition(arr, low, high, stats)
		quickSortHoareHelper(arr, low, pivotIndex, stats)
		quickSortHoareHelper(arr, pivotIndex+1, high, stats)
	}
}

func hoarePartition(arr []int, low, high int, stats *SortingStats) int {
	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		for {
			i++
			stats.Comparisons++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			stats.Comparisons++
			if arr[j] <= pivot {
				break
			}
		}

		stats.Comparisons++
		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
		stats.Swaps++
	}
}

// QuickSort3Way implements 3-way QuickSort (Dutch National Flag)
// Efficient for arrays with many duplicate elements
func QuickSort3Way(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	stats := &SortingStats{}
	quickSort3WayHelper(result, 0, len(result)-1, stats)
	return result
}

func quickSort3WayHelper(arr []int, low, high int, stats *SortingStats) {
	if low >= high {
		return
	}

	lt, gt := partition3Way(arr, low, high, stats)
	quickSort3WayHelper(arr, low, lt-1, stats)
	quickSort3WayHelper(arr, gt+1, high, stats)
}

func partition3Way(arr []int, low, high int, stats *SortingStats) (int, int) {
	pivot := arr[low]
	lt := low    // arr[low..lt-1] < pivot
	i := low + 1 // arr[lt..i-1] == pivot
	gt := high   // arr[gt+1..high] > pivot

	for i <= gt {
		stats.Comparisons++
		if arr[i] < pivot {
			arr[lt], arr[i] = arr[i], arr[lt]
			stats.Swaps++
			lt++
			i++
		} else if arr[i] > pivot {
			arr[i], arr[gt] = arr[gt], arr[i]
			stats.Swaps++
			gt--
			// Don't increment i here because we need to check the swapped element
		} else {
			i++
		}
	}

	return lt, gt
}

// HeapSort implements the HeapSort algorithm
// Time: O(n log n) | Space: O(1)
func HeapSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	stats := &SortingStats{}

	n := len(result)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(result, n, i, stats)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		result[0], result[i] = result[i], result[0]
		stats.Swaps++
		heapify(result, i, 0, stats)
	}

	return result
}

func heapify(arr []int, n, i int, stats *SortingStats) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	stats.Comparisons++
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	stats.Comparisons++
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		stats.Swaps++
		heapify(arr, n, largest, stats)
	}
}

// MergeSort implements the MergeSort algorithm
// Time: O(n log n) | Space: O(n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	stats := &SortingStats{}
	return mergeSortHelper(arr, stats)
}

func mergeSortHelper(arr []int, stats *SortingStats) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSortHelper(arr[:mid], stats)
	right := mergeSortHelper(arr[mid:], stats)

	return merge(left, right, stats)
}

func merge(left, right []int, stats *SortingStats) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		stats.Comparisons++
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

// =============================================================================
// NON-COMPARISON BASED SORTING ALGORITHMS
// =============================================================================

// CountingSort implements counting sort for non-negative integers
// Time: O(n + k) where k is the range | Space: O(k)
func CountingSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Find the maximum element to determine range
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	// Create counting array
	count := make([]int, max+1)

	// Count occurrences
	for _, val := range arr {
		count[val]++
	}

	// Create result array
	result := make([]int, 0, len(arr))
	for i := 0; i <= max; i++ {
		for count[i] > 0 {
			result = append(result, i)
			count[i]--
		}
	}

	return result
}

// RadixSort implements radix sort using counting sort as subroutine
// Time: O(d * (n + k)) where d is digits | Space: O(n + k)
func RadixSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)

	// Find maximum to know number of digits
	max := result[0]
	for _, val := range result {
		if val > max {
			max = val
		}
	}

	// Apply counting sort for every digit
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSortByDigit(result, exp)
	}

	return result
}

func countingSortByDigit(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10) // 0-9 digits

	// Count occurrences of each digit
	for i := 0; i < n; i++ {
		count[(arr[i]/exp)%10]++
	}

	// Change count[i] to actual position
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build output array
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}

	// Copy output array to arr
	copy(arr, output)
}

// BucketSort implements bucket sort
// Time: O(n²) worst case, O(n + k) average | Space: O(n)
func BucketSort(arr []float64, bucketCount int) []float64 {
	if len(arr) <= 1 {
		return arr
	}

	if bucketCount <= 0 {
		bucketCount = len(arr)
	}

	// Find min and max values
	min, max := arr[0], arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	// Create buckets
	buckets := make([][]float64, bucketCount)
	bucketRange := (max - min) / float64(bucketCount)

	// Distribute elements into buckets
	for _, val := range arr {
		bucketIndex := int((val - min) / bucketRange)
		if bucketIndex >= bucketCount {
			bucketIndex = bucketCount - 1
		}
		buckets[bucketIndex] = append(buckets[bucketIndex], val)
	}

	// Sort individual buckets and concatenate
	result := make([]float64, 0, len(arr))
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertionSortFloat64(bucket)
			result = append(result, bucket...)
		}
	}

	return result
}

func insertionSortFloat64(arr []float64) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// =============================================================================
// SIMPLE SORTING ALGORITHMS (for small arrays or educational purposes)
// =============================================================================

// InsertionSort implements insertion sort
// Time: O(n²) | Space: O(1)
func InsertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	stats := &SortingStats{}

	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1

		for j >= 0 {
			stats.Comparisons++
			if result[j] > key {
				result[j+1] = result[j]
				stats.Swaps++
				j--
			} else {
				break
			}
		}
		result[j+1] = key
	}

	return result
}

// SelectionSort implements selection sort
// Time: O(n²) | Space: O(1)
func SelectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	stats := &SortingStats{}

	for i := 0; i < len(result)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(result); j++ {
			stats.Comparisons++
			if result[j] < result[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			result[i], result[minIndex] = result[minIndex], result[i]
			stats.Swaps++
		}
	}

	return result
}

// BubbleSort implements bubble sort
// Time: O(n²) | Space: O(1)
func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	stats := &SortingStats{}

	n := len(result)
	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			stats.Comparisons++
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				stats.Swaps++
				swapped = true
			}
		}

		// If no swapping occurred, array is sorted
		if !swapped {
			break
		}
	}

	return result
}

// =============================================================================
// HYBRID AND OPTIMIZED ALGORITHMS
// =============================================================================

// IntroSort implements introspective sort (hybrid of quicksort, heapsort, and insertion sort)
// Used by many standard libraries including Go's sort package
func IntroSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	stats := &SortingStats{}

	maxDepth := 2 * ilog2(len(result))
	introSortHelper(result, 0, len(result)-1, maxDepth, stats)
	return result
}

func introSortHelper(arr []int, low, high, maxDepth int, stats *SortingStats) {
	for high-low > 16 {
		if maxDepth == 0 {
			// Switch to HeapSort when recursion depth is too high
			heapSortRange(arr, low, high+1, stats)
			return
		}

		maxDepth--
		pivot := lomutoPartition(arr, low, high, stats)

		// Optimize by sorting smaller partition first
		if pivot-low < high-pivot {
			introSortHelper(arr, low, pivot-1, maxDepth, stats)
			low = pivot + 1
		} else {
			introSortHelper(arr, pivot+1, high, maxDepth, stats)
			high = pivot - 1
		}
	}

	// Use insertion sort for small arrays
	insertionSortRange(arr, low, high+1, stats)
}

func heapSortRange(arr []int, start, end int, stats *SortingStats) {
	// Build heap
	for i := start + (end-start)/2 - 1; i >= start; i-- {
		heapifyRange(arr, start, end, i, stats)
	}

	// Extract elements
	for i := end - 1; i > start; i-- {
		arr[start], arr[i] = arr[i], arr[start]
		stats.Swaps++
		heapifyRange(arr, start, i, start, stats)
	}
}

func heapifyRange(arr []int, start, end, i int, stats *SortingStats) {
	largest := i
	left := start + 2*(i-start) + 1
	right := start + 2*(i-start) + 2

	if left < end {
		stats.Comparisons++
		if arr[left] > arr[largest] {
			largest = left
		}
	}

	if right < end {
		stats.Comparisons++
		if arr[right] > arr[largest] {
			largest = right
		}
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		stats.Swaps++
		heapifyRange(arr, start, end, largest, stats)
	}
}

func insertionSortRange(arr []int, start, end int, stats *SortingStats) {
	for i := start + 1; i < end; i++ {
		key := arr[i]
		j := i - 1

		for j >= start {
			stats.Comparisons++
			if arr[j] > key {
				arr[j+1] = arr[j]
				stats.Swaps++
				j--
			} else {
				break
			}
		}
		arr[j+1] = key
	}
}

func ilog2(n int) int {
	result := 0
	for n > 1 {
		n >>= 1
		result++
	}
	return result
}

// =============================================================================
// UTILITY FUNCTIONS
// =============================================================================

// IsSorted checks if an array is sorted in ascending order
func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// GenerateRandomArray creates a random array of given size
func GenerateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size * 10)
	}
	return arr
}

// GenerateReverseSortedArray creates a reverse-sorted array
func GenerateReverseSortedArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = size - i
	}
	return arr
}

// GenerateNearlySortedArray creates an array that's mostly sorted
func GenerateNearlySortedArray(size int, swaps int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i + 1
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < swaps; i++ {
		idx1 := rand.Intn(size)
		idx2 := rand.Intn(size)
		arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
	}

	return arr
}

// GenerateArrayWithDuplicates creates an array with many duplicate values
func GenerateArrayWithDuplicates(size int, uniqueValues int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(uniqueValues)
	}
	return arr
}
