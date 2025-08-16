package fundamentals

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🚀 DSA Mastery in Go - Complexity Analysis Demo")
	fmt.Println("===============================================")

	// Run the complexity demo
	DemoComplexityDifferences()

	// Additional performance comparisons
	fmt.Println("🧪 Advanced Performance Analysis")
	fmt.Println("================================")

	// Memory allocation patterns
	demoMemoryPatterns()

	// Fibonacci comparison
	demoFibonacciComparison()

	// Search algorithm scaling
	demoSearchScaling()
}

func demoMemoryPatterns() {
	fmt.Println("\n📊 Memory Allocation Patterns:")

	// Slice growth demonstration
	fmt.Println("\nSlice Growth Pattern:")
	slice := make([]int, 0)
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
	}

	// Pre-allocated vs growing slice
	fmt.Println("\nPre-allocated vs Growing Slice Performance:")

	// Growing slice
	start := time.Now()
	var growingSlice []int
	for i := 0; i < 100000; i++ {
		growingSlice = append(growingSlice, i)
	}
	growingTime := time.Since(start)

	// Pre-allocated slice
	start = time.Now()
	preAllocSlice := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		preAllocSlice = append(preAllocSlice, i)
	}
	preAllocTime := time.Since(start)

	fmt.Printf("Growing slice: %v\n", growingTime)
	fmt.Printf("Pre-allocated slice: %v\n", preAllocTime)
	fmt.Printf("Pre-allocation speedup: %.2fx\n", float64(growingTime)/float64(preAllocTime))
}

func demoFibonacciComparison() {
	fmt.Println("\n🔢 Fibonacci Algorithm Comparison:")

	n := 30

	// Naive approach
	start := time.Now()
	result1 := Fibonacci(n)
	naiveTime := time.Since(start)

	// Memoized approach
	start = time.Now()
	result2 := FibonacciMemo(n)
	memoTime := time.Since(start)

	fmt.Printf("Fibonacci(%d):\n", n)
	fmt.Printf("Naive result: %d, Time: %v\n", result1, naiveTime)
	fmt.Printf("Memoized result: %d, Time: %v\n", result2, memoTime)
	fmt.Printf("Memoization speedup: %.2fx\n", float64(naiveTime)/float64(memoTime))
}

func demoSearchScaling() {
	fmt.Println("\n🔍 Search Algorithm Scaling Analysis:")

	sizes := []int{1000, 10000, 100000, 1000000}

	for _, size := range sizes {
		arr := GenerateArray(size)
		target := size - 1 // Worst case for linear search

		// Linear search timing
		start := time.Now()
		LinearSearch(arr, target)
		linearTime := time.Since(start)

		// Binary search timing
		start = time.Now()
		BinarySearch(arr, target)
		binaryTime := time.Since(start)

		fmt.Printf("\nArray size: %d\n", size)
		fmt.Printf("Linear search: %v\n", linearTime)
		fmt.Printf("Binary search: %v\n", binaryTime)
		if binaryTime > 0 {
			fmt.Printf("Binary search speedup: %.2fx\n", float64(linearTime)/float64(binaryTime))
		}
	}
}

// Additional utility functions for demonstration
func measureExecution(name string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	fmt.Printf("%s executed in: %v\n", name, duration)
}

func compareAlgorithms(name1, name2 string, fn1, fn2 func()) {
	fmt.Printf("\n🔄 Comparing %s vs %s:\n", name1, name2)

	start := time.Now()
	fn1()
	time1 := time.Since(start)

	start = time.Now()
	fn2()
	time2 := time.Since(start)

	fmt.Printf("%s: %v\n", name1, time1)
	fmt.Printf("%s: %v\n", name2, time2)

	if time2 > 0 {
		ratio := float64(time1) / float64(time2)
		if ratio > 1 {
			fmt.Printf("%s is %.2fx faster\n", name2, ratio)
		} else {
			fmt.Printf("%s is %.2fx faster\n", name1, 1/ratio)
		}
	}
}
