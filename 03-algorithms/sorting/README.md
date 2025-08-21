# Sorting Algorithms

This module provides a comprehensive collection of sorting algorithms implemented in Go, designed for educational purposes and practical use in understanding algorithmic complexity and performance characteristics.

## 📋 Table of Contents

- [Implemented Algorithms](#implemented-algorithms)
- [Performance Characteristics](#performance-characteristics)
- [Usage Examples](#usage-examples)
- [Test Coverage](#test-coverage)
- [Benchmarks](#benchmarks)
- [Educational Notes](#educational-notes)

## 🔧 Implemented Algorithms

### Comparison-Based Sorting Algorithms

#### 1. QuickSort (Lomuto Partition)
- **Time Complexity**: O(n log n) average, O(n²) worst case
- **Space Complexity**: O(log n) average, O(n) worst case
- **Stability**: Not stable
- **Best for**: General-purpose sorting, good average performance

```go
result := sorting.QuickSort([]int{3, 6, 8, 10, 1, 2, 1})
// Output: [1, 1, 2, 3, 6, 8, 10]
```

#### 2. QuickSort with Hoare Partition
- **Time Complexity**: O(n log n) average, O(n²) worst case
- **Space Complexity**: O(log n) average, O(n) worst case
- **Stability**: Not stable
- **Best for**: Fewer swaps than Lomuto partition

```go
result := sorting.QuickSortHoare([]int{64, 34, 25, 12, 22, 11, 90})
```

#### 3. 3-Way QuickSort (Dutch National Flag)
- **Time Complexity**: O(n log n) average, O(n) for many duplicates
- **Space Complexity**: O(log n) average
- **Stability**: Not stable
- **Best for**: Arrays with many duplicate elements

```go
result := sorting.QuickSort3Way([]int{7, 7, 7, 1, 2, 7, 7})
```

#### 4. HeapSort
- **Time Complexity**: O(n log n) all cases
- **Space Complexity**: O(1)
- **Stability**: Not stable
- **Best for**: Guaranteed O(n log n) performance, in-place sorting

```go
result := sorting.HeapSort([]int{64, 34, 25, 12, 22, 11, 90})
```

#### 5. MergeSort
- **Time Complexity**: O(n log n) all cases
- **Space Complexity**: O(n)
- **Stability**: Stable
- **Best for**: Guaranteed performance, stable sorting, external sorting

```go
result := sorting.MergeSort([]int{38, 27, 43, 3, 9, 82, 10})
```

#### 6. InsertionSort
- **Time Complexity**: O(n²) worst case, O(n) best case
- **Space Complexity**: O(1)
- **Stability**: Stable
- **Best for**: Small arrays, nearly sorted data

```go
result := sorting.InsertionSort([]int{5, 2, 4, 6, 1, 3})
```

#### 7. SelectionSort
- **Time Complexity**: O(n²) all cases
- **Space Complexity**: O(1)
- **Stability**: Not stable
- **Best for**: Memory-constrained environments, educational purposes

```go
result := sorting.SelectionSort([]int{29, 10, 14, 37, 13})
```

#### 8. BubbleSort
- **Time Complexity**: O(n²) worst case, O(n) best case
- **Space Complexity**: O(1)
- **Stability**: Stable
- **Best for**: Educational purposes, very small datasets

```go
result := sorting.BubbleSort([]int{5, 1, 4, 2, 8})
```

#### 9. IntroSort (Hybrid Algorithm)
- **Time Complexity**: O(n log n) all cases
- **Space Complexity**: O(log n)
- **Stability**: Not stable
- **Best for**: Production systems (used by standard libraries)

```go
result := sorting.IntroSort([]int{3, 6, 8, 10, 1, 2, 1})
```

### Non-Comparison Based Sorting Algorithms

#### 1. CountingSort
- **Time Complexity**: O(n + k) where k is the range of input
- **Space Complexity**: O(k)
- **Stability**: Stable
- **Constraints**: Non-negative integers only
- **Best for**: Small range of integers

```go
result := sorting.CountingSort([]int{4, 2, 2, 8, 3, 3, 1})
// Output: [1, 2, 2, 3, 3, 4, 8]
```

#### 2. RadixSort
- **Time Complexity**: O(d × (n + k)) where d is number of digits
- **Space Complexity**: O(n + k)
- **Stability**: Stable
- **Constraints**: Non-negative integers
- **Best for**: Large datasets with fixed-width integers

```go
result := sorting.RadixSort([]int{170, 45, 75, 90, 2, 802, 24, 66})
// Output: [2, 24, 45, 66, 75, 90, 170, 802]
```

#### 3. BucketSort
- **Time Complexity**: O(n + k) average, O(n²) worst case
- **Space Complexity**: O(n + k)
- **Stability**: Stable (with stable sorting for buckets)
- **Best for**: Uniformly distributed floating-point numbers

```go
result := sorting.BucketSort([]float64{0.42, 0.32, 0.33, 0.52, 0.37}, 5)
// Output: [0.32, 0.33, 0.37, 0.42, 0.52]
```

## 📊 Performance Characteristics

### Time Complexity Summary

| Algorithm | Best Case | Average Case | Worst Case | Space Complexity |
|-----------|-----------|--------------|------------|------------------|
| QuickSort | O(n log n) | O(n log n) | O(n²) | O(log n) |
| HeapSort | O(n log n) | O(n log n) | O(n log n) | O(1) |
| MergeSort | O(n log n) | O(n log n) | O(n log n) | O(n) |
| IntroSort | O(n log n) | O(n log n) | O(n log n) | O(log n) |
| InsertionSort | O(n) | O(n²) | O(n²) | O(1) |
| SelectionSort | O(n²) | O(n²) | O(n²) | O(1) |
| BubbleSort | O(n) | O(n²) | O(n²) | O(1) |
| CountingSort | O(n + k) | O(n + k) | O(n + k) | O(k) |
| RadixSort | O(d(n + k)) | O(d(n + k)) | O(d(n + k)) | O(n + k) |
| BucketSort | O(n + k) | O(n + k) | O(n²) | O(n + k) |

### When to Use Each Algorithm

- **QuickSort**: General purpose, good average performance
- **HeapSort**: When worst-case O(n log n) is required
- **MergeSort**: When stability is needed or for external sorting
- **IntroSort**: Production systems (hybrid of QuickSort, HeapSort, InsertionSort)
- **InsertionSort**: Small arrays (n < 50) or nearly sorted data
- **CountingSort**: Small range of non-negative integers
- **RadixSort**: Large datasets of integers with limited digits
- **BucketSort**: Uniformly distributed floating-point numbers

## 🔧 Usage Examples

### Basic Sorting

```go
package main

import (
    "fmt"
    "dsa-mastery/03-algorithms/sorting"
)

func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    
    // Try different sorting algorithms
    quickSorted := sorting.QuickSort(arr)
    heapSorted := sorting.HeapSort(arr)
    mergeSorted := sorting.MergeSort(arr)
    
    fmt.Println("QuickSort:", quickSorted)
    fmt.Println("HeapSort:", heapSorted)
    fmt.Println("MergeSort:", mergeSorted)
}
```

### Performance Testing

```go
package main

import (
    "fmt"
    "time"
    "dsa-mastery/03-algorithms/sorting"
)

func main() {
    // Generate test data
    arr := sorting.GenerateRandomArray(10000)
    
    // Benchmark QuickSort
    start := time.Now()
    result := sorting.QuickSort(arr)
    duration := time.Since(start)
    
    fmt.Printf("QuickSort took %v for %d elements\n", duration, len(result))
    fmt.Printf("Is sorted: %v\n", sorting.IsSorted(result))
}
```

### Specialized Sorting

```go
package main

import (
    "fmt"
    "dsa-mastery/03-algorithms/sorting"
)

func main() {
    // For arrays with many duplicates
    duplicates := []int{3, 1, 3, 3, 2, 1, 2, 3, 1}
    result := sorting.QuickSort3Way(duplicates)
    fmt.Println("3-Way QuickSort:", result)
    
    // For small range integers
    smallRange := []int{4, 2, 2, 8, 3, 3, 1}
    countResult := sorting.CountingSort(smallRange)
    fmt.Println("CountingSort:", countResult)
    
    // For floating-point numbers
    floats := []float64{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434}
    bucketResult := sorting.BucketSort(floats, 3)
    fmt.Println("BucketSort:", bucketResult)
}
```

## 🧪 Test Coverage

### Comprehensive Testing

The module includes extensive test coverage:

- **Unit Tests**: Every algorithm with multiple test cases
- **Edge Cases**: Empty arrays, single elements, duplicates
- **Stress Tests**: Large arrays (up to 100K elements)
- **Property Tests**: Verify sorting correctness and stability
- **Performance Tests**: Benchmark comparisons

Run tests:
```bash
go test -v                    # Run all tests
go test -bench=.             # Run benchmarks
go test -cover               # Check test coverage
```

### Test Categories

1. **Basic Functionality Tests**
   - Empty arrays
   - Single elements
   - Already sorted arrays
   - Reverse sorted arrays
   - Random order arrays
   - Arrays with duplicates

2. **Edge Case Tests**
   - Large arrays with many duplicates
   - Nearly sorted arrays
   - Negative numbers
   - Mixed positive/negative numbers

3. **Performance Tests**
   - Comparison with Go's standard library
   - Different input sizes
   - Different input patterns

## 📈 Benchmarks

### Performance Results (Apple M3 Pro)

For arrays of 10,000 elements:

| Algorithm | Time per Operation | Operations/sec |
|-----------|-------------------|----------------|
| Go Std Sort | ~376μs | ~2,656 |
| IntroSort | ~332μs | ~3,012 |
| QuickSort | ~357μs | ~2,802 |
| HeapSort | ~450μs | ~2,222 |
| MergeSort | ~520μs | ~1,923 |

*Note: Results may vary based on input data distribution and hardware.*

### Memory Usage

- **In-place algorithms**: HeapSort, QuickSort variants, InsertionSort, SelectionSort, BubbleSort
- **O(n) space**: MergeSort, CountingSort, RadixSort, BucketSort
- **O(log n) space**: QuickSort (recursion stack)

## 📚 Educational Notes

### Algorithm Selection Guide

1. **General Purpose**: Use IntroSort or QuickSort
2. **Guaranteed Performance**: Use HeapSort
3. **Stability Required**: Use MergeSort
4. **Memory Constrained**: Use HeapSort or in-place QuickSort
5. **Small Arrays**: Use InsertionSort (n < 50)
6. **Many Duplicates**: Use 3-Way QuickSort
7. **Integer Range Known**: Use CountingSort
8. **Large Integers**: Use RadixSort
9. **Floating Point**: Use BucketSort (if uniformly distributed)

### Key Concepts

- **Stability**: Maintains relative order of equal elements
- **In-place**: Uses O(1) extra space
- **Adaptive**: Performs better on partially sorted data
- **Divide & Conquer**: Recursively divides problem into subproblems
- **Comparison-based**: Sorts by comparing elements
- **Non-comparison**: Uses element properties (digits, range)

### Learning Path

1. Start with simple algorithms (BubbleSort, SelectionSort)
2. Understand divide-and-conquer (MergeSort, QuickSort)
3. Learn about heaps (HeapSort)
4. Explore hybrid approaches (IntroSort)
5. Study non-comparison algorithms (CountingSort, RadixSort)
6. Analyze trade-offs and real-world applications

## 🔍 Utility Functions

The module also provides helpful utility functions:

- `IsSorted(arr []int) bool`: Check if array is sorted
- `GenerateRandomArray(size int) []int`: Generate test data
- `GenerateArrayWithDuplicates(size, uniqueCount int) []int`: Generate data with duplicates
- `GenerateNearlySortedArray(size, swaps int) []int`: Generate nearly sorted data
- `GenerateReverseSortedArray(size int) []int`: Generate reverse sorted data

These utilities are invaluable for testing and understanding algorithm behavior under different conditions.
