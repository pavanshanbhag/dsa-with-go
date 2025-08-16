# Time & Space Complexity Analysis in Go

## Big O Notation Fundamentals

Understanding algorithmic complexity is crucial for writing efficient Go code. This guide covers practical complexity analysis with Go-specific considerations.

## Time Complexity Categories

### O(1) - Constant Time
```go
// Array/slice access
func getElement(arr []int, index int) int {
    return arr[index] // O(1)
}

// Map operations (average case)
func getValue(m map[string]int, key string) int {
    return m[key] // O(1) average
}
```

### O(log n) - Logarithmic Time
```go
// Binary search
func binarySearch(arr []int, target int) int {
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
```

### O(n) - Linear Time
```go
// Linear search
func linearSearch(arr []int, target int) int {
    for i, val := range arr {
        if val == target {
            return i
        }
    }
    return -1
}
```

### O(n log n) - Linearithmic Time
```go
// Merge sort
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    
    return merge(left, right)
}
```

### O(n²) - Quadratic Time
```go
// Bubble sort
func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
```

## Space Complexity Considerations

### Go-Specific Memory Patterns

1. **Slice Growth**: Understanding slice capacity and reallocation
2. **Escape Analysis**: When variables move to heap
3. **Garbage Collector**: Impact on algorithm choice
4. **Memory Alignment**: Struct padding and cache efficiency

## Benchmarking in Go

```go
func BenchmarkAlgorithm(b *testing.B) {
    data := generateTestData(1000)
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        algorithm(data)
    }
}
```

## Next Steps
- Implement each complexity category with benchmarks
- Compare theoretical vs actual performance
- Analyze memory usage patterns
- Study Go runtime optimizations
