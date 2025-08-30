package main

import (
	"fmt"
	"time"
	"dsa-mastery/03-algorithms/dynamic-programming"
)

func main() {
	fmt.Println("🎯 DSA Mastery - Dynamic Programming Demonstration")
	fmt.Println("=================================================")
	
	// 1. Fibonacci - Classic DP Introduction
	fmt.Println("\n📈 1. Fibonacci Sequence (Classic DP Introduction)")
	fmt.Println("--------------------------------------------------")
	
	n := 20
	fmt.Printf("Computing Fibonacci(%d):\n", n)
	
	// Compare naive vs DP approaches
	start := time.Now()
	result := dp.FibonacciNaive(n)
	naiveTime := time.Since(start)
	
	start = time.Now()
	result = dp.FibonacciMemoization(n)
	memoTime := time.Since(start)
	
	start = time.Now()
	result = dp.FibonacciTabulation(n)
	tabTime := time.Since(start)
	
	start = time.Now()
	result = dp.FibonacciOptimized(n)
	optTime := time.Since(start)
	
	fmt.Printf("Naive recursion:  %d (Time: %v) [O(2^n)]\n", result, naiveTime)
	fmt.Printf("Memoization:      %d (Time: %v) [O(n)]\n", result, memoTime)
	fmt.Printf("Tabulation:       %d (Time: %v) [O(n)]\n", result, tabTime)
	fmt.Printf("Space Optimized:  %d (Time: %v) [O(1) space]\n", result, optTime)
	
	// 2. Knapsack Problems
	fmt.Println("\n🎒 2. Knapsack Problems (Optimization)")
	fmt.Println("-------------------------------------")
	
	items := []dp.Item{
		{Weight: 10, Value: 60, Name: "Diamond Ring"},
		{Weight: 20, Value: 100, Name: "Gold Necklace"},
		{Weight: 30, Value: 120, Name: "Silver Watch"},
		{Weight: 15, Value: 80, Name: "Ruby Bracelet"},
	}
	capacity := 50
	
	fmt.Printf("Items available:\n")
	for _, item := range items {
		fmt.Printf("  %s: Weight=%d, Value=%d\n", item.Name, item.Weight, item.Value)
	}
	fmt.Printf("Knapsack capacity: %d\n\n", capacity)
	
	// 0/1 Knapsack
	result01 := dp.Knapsack01Optimized(items, capacity)
	fmt.Printf("0/1 Knapsack (each item once): Max value = %d\n", result01)
	
	// Unbounded Knapsack
	resultUnbound := dp.KnapsackUnbounded(items, capacity)
	fmt.Printf("Unbounded Knapsack (unlimited): Max value = %d\n", resultUnbound)
	
	// 3. String DP Problems
	fmt.Println("\n🔤 3. String Dynamic Programming")
	fmt.Println("--------------------------------")
	
	// Longest Common Subsequence
	text1 := "PROGRAMMING"
	text2 := "ALGORITHM"
	lcsLength := dp.LCSOptimized(text1, text2)
	fmt.Printf("LCS of '%s' and '%s': Length = %d\n", text1, text2, lcsLength)
	
	// Edit Distance
	word1 := "kitten"
	word2 := "sitting"
	editDist := dp.EditDistanceOptimized(word1, word2)
	fmt.Printf("Edit distance '%s' → '%s': %d operations\n", word1, word2, editDist)
	
	// Palindrome problems
	palindromeText := "BBABCBCAB"
	lpsLength := dp.LongestPalindromicSubsequence(palindromeText)
	fmt.Printf("Longest palindromic subsequence in '%s': Length = %d\n", palindromeText, lpsLength)
	
	// 4. Array DP Problems
	fmt.Println("\n📊 4. Array Dynamic Programming")
	fmt.Println("-------------------------------")
	
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Printf("Array: %v\n", nums)
	
	lisLength := dp.LISTabulation(nums)
	lisLengthFast := dp.LISBinarySearch(nums)
	fmt.Printf("Longest Increasing Subsequence: %d (O(n²) approach)\n", lisLength)
	fmt.Printf("Longest Increasing Subsequence: %d (O(n log n) approach)\n", lisLengthFast)
	
	// 5. Coin Change Problems
	fmt.Println("\n💰 5. Coin Change Problems")
	fmt.Println("--------------------------")
	
	coins := []int{1, 5, 10, 25}
	amount := 67
	fmt.Printf("Coins available: %v\n", coins)
	fmt.Printf("Target amount: %d cents\n", amount)
	
	minCoins := dp.CoinChangeMinCoins(coins, amount)
	ways := dp.CoinChangeCountWays(coins, amount)
	fmt.Printf("Minimum coins needed: %d\n", minCoins)
	fmt.Printf("Number of ways to make amount: %d\n", ways)
	
	// 6. Matrix Chain Multiplication
	fmt.Println("\n🔢 6. Matrix Chain Multiplication")
	fmt.Println("---------------------------------")
	
	dimensions := []int{40, 20, 30, 10, 30}
	minOps := dp.MatrixChainMultiplication(dimensions)
	fmt.Printf("Matrix dimensions: %v\n", dimensions)
	fmt.Printf("Minimum scalar multiplications: %d\n", minOps)
	
	// 7. Performance Comparison
	fmt.Println("\n⚡ 7. Performance Analysis (Large Input)")
	fmt.Println("---------------------------------------")
	
	largeFib := 40
	fmt.Printf("Computing Fibonacci(%d) with different approaches:\n", largeFib)
	
	start = time.Now()
	_ = dp.FibonacciMemoization(largeFib)
	largeMemoTime := time.Since(start)
	
	start = time.Now()
	_ = dp.FibonacciTabulation(largeFib)
	largeTabTime := time.Since(start)
	
	start = time.Now()
	_ = dp.FibonacciOptimized(largeFib)
	largeOptTime := time.Since(start)
	
	fmt.Printf("Memoization:      %v\n", largeMemoTime)
	fmt.Printf("Tabulation:       %v\n", largeTabTime)
	fmt.Printf("Space Optimized:  %v (%.1fx faster than memoization)\n", 
		largeOptTime, float64(largeMemoTime)/float64(largeOptTime))
	
	// 8. Real-world Application Examples
	fmt.Println("\n🌟 8. Real-World Applications")
	fmt.Println("-----------------------------")
	
	fmt.Println("✓ Resource Allocation (Knapsack)")
	fmt.Println("  - Budget optimization")
	fmt.Println("  - CPU scheduling")
	fmt.Println("  - Investment portfolio")
	
	fmt.Println("\n✓ Bioinformatics (LCS/Edit Distance)")
	fmt.Println("  - DNA sequence alignment")
	fmt.Println("  - Protein folding analysis")
	fmt.Println("  - Phylogenetic trees")
	
	fmt.Println("\n✓ Text Processing (String DP)")
	fmt.Println("  - Spell checkers")
	fmt.Println("  - Version control diffs")
	fmt.Println("  - Data deduplication")
	
	fmt.Println("\n✓ Financial Modeling (Optimization)")
	fmt.Println("  - Option pricing")
	fmt.Println("  - Risk management")
	fmt.Println("  - Portfolio optimization")
	
	fmt.Println("\n✓ Game Development (State Optimization)")
	fmt.Println("  - Path finding")
	fmt.Println("  - AI decision trees")
	fmt.Println("  - Resource management")
	
	// 9. Key Learning Points
	fmt.Println("\n🎓 9. Key Learning Points")
	fmt.Println("-------------------------")
	
	fmt.Println("✓ DP Characteristics:")
	fmt.Println("  - Optimal Substructure: Problem breaks into optimal subproblems")
	fmt.Println("  - Overlapping Subproblems: Same calculations repeated")
	fmt.Println("  - Memoization vs Tabulation: Top-down vs Bottom-up")
	
	fmt.Println("\n✓ When to Use DP:")
	fmt.Println("  - Optimization problems (min/max)")
	fmt.Println("  - Counting problems")
	fmt.Println("  - Decision problems with choices")
	fmt.Println("  - Problems with exponential naive solutions")
	
	fmt.Println("\n✓ Space Optimization Techniques:")
	fmt.Println("  - Keep only necessary previous states")
	fmt.Println("  - Rolling arrays for 2D problems")
	fmt.Println("  - State compression for complex states")
	
	fmt.Println("\n✅ Dynamic Programming Mastery Complete!")
	fmt.Println("========================================")
	fmt.Printf("🎯 Implemented: %d major DP problem categories\n", 8)
	fmt.Printf("📊 Algorithms: %d different approaches (memoization, tabulation, optimized)\n", 3)
	fmt.Printf("⚡ Performance: Up to %.0fx speedup from space optimization\n", 
		float64(largeMemoTime)/float64(largeOptTime))
	fmt.Printf("🧠 Patterns: Classic problems covering all major DP paradigms\n")
	
	fmt.Println("\n📚 What's Next?")
	fmt.Println("- String Algorithms (KMP, Rabin-Karp, Trie)")
	fmt.Println("- Backtracking (N-Queens, Sudoku, Permutations)")
	fmt.Println("- Advanced DP (Bitmask DP, Tree DP, Digit DP)")
}
