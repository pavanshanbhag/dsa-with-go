package main

import (
	"fmt"
	"math/rand"
	"time"

	"dsa-mastery/03-algorithms/sorting"
)

func main() {
	fmt.Println("🔢 DSA Mastery - Sorting Algorithms Demonstration")
	fmt.Println("================================================")

	// 1. Comparison-Based Sorting Algorithms
	fmt.Println("\n⚖️  1. Comparison-Based Sorting Algorithms")
	fmt.Println("------------------------------------------")
	demonstrateComparisonSorts()

	// 2. Non-Comparison Sorting Algorithms
	fmt.Println("\n🚀 2. Non-Comparison Sorting Algorithms")
	fmt.Println("--------------------------------------")
	demonstrateNonComparisonSorts()

	// 3. Performance Analysis
	fmt.Println("\n📊 3. Performance Analysis")
	fmt.Println("--------------------------")
	demonstratePerformanceAnalysis()

	// 4. Real-World Sorting Applications
	fmt.Println("\n🌟 4. Real-World Applications")
	fmt.Println("-----------------------------")
	demonstrateRealWorldApplications()

	fmt.Println("\n✅ All sorting algorithms demonstrated successfully!")
	fmt.Println("🚀 Ready to tackle any sorting challenge!")
}

func demonstrateComparisonSorts() {
	fmt.Println("Problem: Sort arrays using comparison-based algorithms")

	// Test data
	original := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42}
	fmt.Printf("Original array: %v\n", original)

	// QuickSort variations
	fmt.Println("\n🏃 QuickSort Family:")

	// Standard QuickSort (Lomuto partition)
	data1 := make([]int, len(original))
	copy(data1, original)
	start := time.Now()
	result1 := sorting.QuickSort(data1)
	duration1 := time.Since(start)
	fmt.Printf("QuickSort (Lomuto):     %v (Time: %v)\n", result1, duration1)

	// QuickSort with Hoare partition
	data2 := make([]int, len(original))
	copy(data2, original)
	start = time.Now()
	result2 := sorting.QuickSortHoare(data2)
	duration2 := time.Since(start)
	fmt.Printf("QuickSort (Hoare):      %v (Time: %v)\n", result2, duration2)

	// 3-Way QuickSort
	data3 := make([]int, len(original))
	copy(data3, original)
	start = time.Now()
	result3 := sorting.QuickSort3Way(data3)
	duration3 := time.Since(start)
	fmt.Printf("QuickSort (3-Way):      %v (Time: %v)\n", result3, duration3)

	// Other comparison sorts
	fmt.Println("\n🔄 Other Comparison Sorts:")

	// HeapSort
	data4 := make([]int, len(original))
	copy(data4, original)
	start = time.Now()
	result4 := sorting.HeapSort(data4)
	duration4 := time.Since(start)
	fmt.Printf("HeapSort:               %v (Time: %v)\n", result4, duration4)

	// MergeSort
	data5 := make([]int, len(original))
	copy(data5, original)
	start = time.Now()
	result5 := sorting.MergeSort(data5)
	duration5 := time.Since(start)
	fmt.Printf("MergeSort:              %v (Time: %v)\n", result5, duration5)

	// InsertionSort
	data6 := make([]int, len(original))
	copy(data6, original)
	start = time.Now()
	result6 := sorting.InsertionSort(data6)
	duration6 := time.Since(start)
	fmt.Printf("InsertionSort:          %v (Time: %v)\n", result6, duration6)

	// SelectionSort
	data7 := make([]int, len(original))
	copy(data7, original)
	start = time.Now()
	result7 := sorting.SelectionSort(data7)
	duration7 := time.Since(start)
	fmt.Printf("SelectionSort:          %v (Time: %v)\n", result7, duration7)
}

func demonstrateNonComparisonSorts() {
	fmt.Println("Problem: Sort arrays without comparing elements")

	// Integer sorting
	original := []int{170, 45, 75, 90, 2, 802, 24, 66}
	fmt.Printf("Original array: %v\n", original)

	// CountingSort
	data1 := make([]int, len(original))
	copy(data1, original)
	start := time.Now()
	result1 := sorting.CountingSort(data1)
	duration1 := time.Since(start)
	fmt.Printf("CountingSort:           %v (Time: %v)\n", result1, duration1)

	// RadixSort
	data2 := make([]int, len(original))
	copy(data2, original)
	start = time.Now()
	result2 := sorting.RadixSort(data2)
	duration2 := time.Since(start)
	fmt.Printf("RadixSort:              %v (Time: %v)\n", result2, duration2)

	// BucketSort for floats
	fmt.Println("\n🪣 BucketSort for decimal numbers:")
	floats := []float64{0.42, 0.32, 0.33, 0.52, 0.37, 0.47, 0.51}
	fmt.Printf("Original floats: %v\n", floats)
	start = time.Now()
	resultFloats := sorting.BucketSort(floats, 5)
	duration3 := time.Since(start)
	fmt.Printf("BucketSort:             %v (Time: %v)\n", resultFloats, duration3)
}

func demonstratePerformanceAnalysis() {
	fmt.Println("Comparing sorting algorithms on different input sizes and patterns")

	sizes := []int{100, 1000, 5000}

	for _, size := range sizes {
		fmt.Printf("\n📏 Array size: %d elements\n", size)

		// Test different data patterns
		patterns := map[string][]int{
			"Random":    generateRandomArray(size),
			"Sorted":    generateSortedArray(size),
			"Reverse":   generateReverseArray(size),
			"Many Dups": generateDuplicatesArray(size),
		}

		for pattern, data := range patterns {
			fmt.Printf("\n  %s data:\n", pattern)

			// Test selected algorithms
			algorithms := map[string]func([]int) []int{
				"QuickSort": sorting.QuickSort,
				"HeapSort":  sorting.HeapSort,
				"MergeSort": sorting.MergeSort,
			}

			for name, algo := range algorithms {
				testData := make([]int, len(data))
				copy(testData, data)

				start := time.Now()
				algo(testData)
				duration := time.Since(start)

				fmt.Printf("    %-10s: %v\n", name, duration)
			}
		}
	}
}

func demonstrateRealWorldApplications() {
	fmt.Println("Real-world applications of different sorting algorithms:")

	applications := map[string]string{
		"QuickSort":     "General-purpose sorting, standard library implementations",
		"MergeSort":     "External sorting, stable sort requirements, linked lists",
		"HeapSort":      "Memory-constrained environments, guaranteed O(n log n)",
		"InsertionSort": "Small datasets, nearly sorted data, hybrid algorithms",
		"CountingSort":  "Integer sorting with small range, histogram generation",
		"RadixSort":     "Large integer datasets, string sorting, database indexes",
		"BucketSort":    "Uniformly distributed data, parallel sorting",
	}

	for algo, usage := range applications {
		fmt.Printf("🔹 %-13s: %s\n", algo, usage)
	}

	fmt.Println("\n💡 Algorithm Selection Guidelines:")
	fmt.Println("✓ Small arrays (< 50):     InsertionSort")
	fmt.Println("✓ General purpose:         QuickSort or MergeSort")
	fmt.Println("✓ Stability required:      MergeSort")
	fmt.Println("✓ Memory constrained:      HeapSort")
	fmt.Println("✓ Integer data:            CountingSort/RadixSort")
	fmt.Println("✓ Uniform distribution:    BucketSort")
	fmt.Println("✓ Worst-case guarantee:    HeapSort or MergeSort")
}

// Helper functions for generating test data
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size * 10)
	}
	return arr
}

func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

func generateReverseArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	return arr
}

func generateDuplicatesArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i % 10 // Many duplicates
	}
	return arr
}
