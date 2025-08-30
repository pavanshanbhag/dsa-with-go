package main

import (
	"fmt"
	"strings"
	"time"

	string_algorithms "dsa-mastery/03-algorithms/string-algorithms"
)

func main() {
	fmt.Println("🔤 DSA Mastery - String Algorithms Demonstration")
	fmt.Println("===============================================")

	// 1. Pattern Matching Algorithms Comparison
	fmt.Println("\n🔍 1. Pattern Matching Algorithms")
	fmt.Println("---------------------------------")

	text := "ABAAABCDABABCABCABCDAB"
	pattern := "ABCAB"
	fmt.Printf("Text: '%s'\n", text)
	fmt.Printf("Pattern: '%s'\n", pattern)
	fmt.Printf("Text length: %d, Pattern length: %d\n\n", len(text), len(pattern))

	// Test all algorithms and compare performance
	algorithms := map[string]func(string, string) []int{
		"Naive":       string_algorithms.NaivePatternSearch,
		"KMP":         string_algorithms.KMPPatternSearch,
		"Rabin-Karp":  string_algorithms.RabinKarpSearch,
		"Boyer-Moore": string_algorithms.BoyerMooreSearch,
		"Z Algorithm": string_algorithms.ZPatternSearch,
	}

	fmt.Println("Algorithm Performance Comparison:")
	for name, algo := range algorithms {
		start := time.Now()
		matches := algo(text, pattern)
		duration := time.Since(start)

		fmt.Printf("%-12s: %v (Time: %v)\n", name, matches, duration)
	}

	// 2. Large Text Performance Test
	fmt.Println("\n⚡ 2. Large Text Performance Analysis")
	fmt.Println("------------------------------------")

	largeText := strings.Repeat("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1000)
	searchPattern := "MNOPQR"
	fmt.Printf("Large text: %d characters\n", len(largeText))
	fmt.Printf("Searching for: '%s'\n\n", searchPattern)

	fmt.Println("Performance on large text:")
	for name, algo := range algorithms {
		start := time.Now()
		matches := algo(largeText, searchPattern)
		duration := time.Since(start)

		fmt.Printf("%-12s: %d matches in %v\n", name, len(matches), duration)
	}

	// 3. Trie Data Structure Demonstration
	fmt.Println("\n🌳 3. Trie Data Structure")
	fmt.Println("-------------------------")

	trie := string_algorithms.NewTrie()

	// Programming language keywords
	keywords := []string{
		"algorithm", "array", "binary", "class", "data", "function",
		"graph", "hash", "index", "java", "key", "list", "map",
		"node", "object", "pointer", "queue", "recursion", "stack",
		"tree", "variable", "while", "xor", "yield", "zero",
	}

	fmt.Printf("Inserting %d programming terms into Trie...\n", len(keywords))
	start := time.Now()
	for _, word := range keywords {
		trie.Insert(word)
	}
	insertTime := time.Since(start)

	fmt.Printf("✓ Insertion completed in %v\n", insertTime)
	fmt.Printf("✓ Trie size: %d words\n\n", trie.Size())

	// Autocomplete demonstration
	prefixes := []string{"al", "ar", "da", "st", "tr"}
	fmt.Println("Autocomplete demonstration:")
	for _, prefix := range prefixes {
		suggestions := trie.GetWordsWithPrefix(prefix)
		fmt.Printf("'%s*' → %v\n", prefix, suggestions)
	}

	// Search performance
	fmt.Println("\nSearch performance:")
	start = time.Now()
	found := 0
	for _, word := range keywords {
		if trie.Search(word) {
			found++
		}
	}
	searchTime := time.Since(start)
	fmt.Printf("✓ Searched %d words in %v (all found: %v)\n", len(keywords), searchTime, found == len(keywords))

	// 4. String Hashing and Rolling Hash
	fmt.Println("\n🔢 4. String Hashing & Rolling Hash")
	fmt.Println("----------------------------------")

	hasher := string_algorithms.NewStringHasher()

	// Hash comparison
	testStrings := []string{"hello", "world", "algorithm", "data", "structure"}
	fmt.Println("String hash values:")
	for _, str := range testStrings {
		hash := hasher.Hash(str)
		fnvHash := string_algorithms.FNVHash(str)
		fmt.Printf("'%s' → Custom: %d, FNV: %d\n", str, hash, fnvHash)
	}

	// Rolling hash demonstration
	fmt.Println("\nRolling hash demonstration:")
	rollingText := "abcdefghijklmnop"
	windowSize := 4
	fmt.Printf("Text: '%s', Window size: %d\n", rollingText, windowSize)

	rh := string_algorithms.NewRollingHash(rollingText, windowSize)
	if rh != nil {
		fmt.Println("Rolling through text:")
		count := 0
		for {
			window := rh.CurrentWindow()
			hash := rh.CurrentHash()
			fmt.Printf("  Window: '%s' → Hash: %d\n", window, hash)

			count++
			if !rh.RollNext() || count >= 8 { // Limit output
				break
			}
		}
		if count == 8 {
			fmt.Println("  ... (truncated)")
		}
	}

	// 5. Advanced String Algorithms
	fmt.Println("\n🧠 5. Advanced String Algorithms")
	fmt.Println("--------------------------------")

	// Longest Common Prefix
	prefixStrings := []string{"flower", "flow", "flight"}
	commonPrefix := string_algorithms.LongestCommonPrefix(prefixStrings)
	fmt.Printf("Longest common prefix of %v: '%s'\n", prefixStrings, commonPrefix)

	// Palindrome detection
	palindromeText := "racecar"
	longestPalindrome := string_algorithms.LongestPalindromicSubstring(palindromeText)
	fmt.Printf("Longest palindrome in '%s': '%s'\n", palindromeText, longestPalindrome)

	complexText := "babad"
	longestPalindrome2 := string_algorithms.LongestPalindromicSubstring(complexText)
	fmt.Printf("Longest palindrome in '%s': '%s'\n", complexText, longestPalindrome2)

	// Anagram detection
	anagramPairs := [][2]string{
		{"listen", "silent"},
		{"evil", "vile"},
		{"hello", "world"},
		{"algorithm", "logarithm"},
	}

	fmt.Println("\nAnagram detection:")
	for _, pair := range anagramPairs {
		isAnagram := string_algorithms.IsAnagram(pair[0], pair[1])
		status := "❌"
		if isAnagram {
			status = "✅"
		}
		fmt.Printf("%s '%s' ↔ '%s'\n", status, pair[0], pair[1])
	}

	// String reversal
	reverseStrings := []string{"hello", "algorithm", "racecar", "go"}
	fmt.Println("\nString reversal:")
	for _, str := range reverseStrings {
		reversed := string_algorithms.StringReverse(str)
		fmt.Printf("'%s' → '%s'\n", str, reversed)
	}

	// 6. Real-World Applications
	fmt.Println("\n🌍 6. Real-World Applications")
	fmt.Println("-----------------------------")

	// DNA sequence analysis
	fmt.Println("🧬 DNA Sequence Analysis:")
	dnaSequence := "ATCGATCGATCGAAATCGTTAGCATCGATCG"
	genePattern := "ATCG"
	geneMatches := string_algorithms.KMPPatternSearch(dnaSequence, genePattern)
	fmt.Printf("DNA: %s\n", dnaSequence)
	fmt.Printf("Gene pattern '%s' found at positions: %v\n", genePattern, geneMatches)

	// Text processing
	fmt.Println("\n📝 Text Editor Find/Replace:")
	document := "The quick brown fox jumps over the lazy dog. The fox is quick."
	searchTerm := "fox"
	termMatches := string_algorithms.BoyerMooreSearch(document, searchTerm)
	fmt.Printf("Document: %s\n", document)
	fmt.Printf("Term '%s' appears at positions: %v\n", searchTerm, termMatches)

	// Web search autocomplete
	fmt.Println("\n🔍 Search Engine Autocomplete:")
	searchTrie := string_algorithms.NewTrie()
	searchQueries := []string{
		"algorithm", "algorithms", "algorithmic", "algebra", "alpha",
		"data", "database", "datascience", "structure", "structures",
	}

	for _, query := range searchQueries {
		searchTrie.Insert(query)
	}

	userInput := "alg"
	suggestions := searchTrie.GetWordsWithPrefix(userInput)
	fmt.Printf("User types '%s' → Suggestions: %v\n", userInput, suggestions)

	userInput2 := "data"
	suggestions2 := searchTrie.GetWordsWithPrefix(userInput2)
	fmt.Printf("User types '%s' → Suggestions: %v\n", userInput2, suggestions2)

	// 7. Performance Summary
	fmt.Println("\n📊 7. Performance Summary")
	fmt.Println("-------------------------")

	fmt.Println("✅ Pattern Matching Algorithms:")
	fmt.Println("  • Boyer-Moore: Best for large alphabets and long patterns")
	fmt.Println("  • KMP: Optimal O(n+m) guarantee, good general purpose")
	fmt.Println("  • Rabin-Karp: Excellent for multiple pattern search")
	fmt.Println("  • Z Algorithm: Linear time with additional string analysis")
	fmt.Println("  • Naive: Simple but O(n×m) complexity")

	fmt.Println("\n✅ Data Structures:")
	fmt.Printf("  • Trie: Efficiently stored %d words\n", trie.Size())
	fmt.Println("  • Rolling Hash: O(1) substring hash updates")
	fmt.Println("  • String Hashing: Fast fingerprinting for large texts")

	fmt.Println("\n✅ Advanced Algorithms:")
	fmt.Println("  • Manacher's Algorithm: O(n) palindrome detection")
	fmt.Println("  • LCP: Efficient prefix finding")
	fmt.Println("  • String utilities: Anagram detection, reversal")

	// 8. Complexity Analysis
	fmt.Println("\n🎓 8. Time Complexity Summary")
	fmt.Println("-----------------------------")

	complexities := [][]string{
		{"Algorithm", "Best Case", "Average Case", "Worst Case", "Space"},
		{"Naive Search", "O(n)", "O(n×m)", "O(n×m)", "O(1)"},
		{"KMP", "O(n)", "O(n+m)", "O(n+m)", "O(m)"},
		{"Rabin-Karp", "O(n)", "O(n+m)", "O(n×m)", "O(1)"},
		{"Boyer-Moore", "O(n/m)", "O(n)", "O(n×m)", "O(σ)"},
		{"Z Algorithm", "O(n)", "O(n+m)", "O(n+m)", "O(n)"},
		{"Trie Insert", "-", "O(m)", "O(m)", "O(ALPHABET×N×M)"},
		{"Trie Search", "-", "O(m)", "O(m)", "O(1)"},
		{"Rolling Hash", "-", "O(1) per roll", "O(1) per roll", "O(1)"},
	}

	for i, row := range complexities {
		if i == 0 {
			fmt.Printf("%-13s %-8s %-12s %-10s %s\n", row[0], row[1], row[2], row[3], row[4])
			fmt.Println(strings.Repeat("-", 65))
		} else {
			fmt.Printf("%-13s %-8s %-12s %-10s %s\n", row[0], row[1], row[2], row[3], row[4])
		}
	}

	fmt.Println("\n🎯 String Algorithms Mastery Complete!")
	fmt.Println("======================================")
	fmt.Printf("📊 Implemented: %d pattern matching algorithms\n", len(algorithms))
	fmt.Printf("🌳 Trie operations: Insert, Search, Delete, Prefix matching\n")
	fmt.Printf("🔢 Hashing: Custom, FNV, Rolling hash implementations\n")
	fmt.Printf("🧠 Advanced: Palindromes, LCP, anagrams, string utilities\n")
	fmt.Printf("⚡ Performance: Benchmarked across multiple text sizes\n")

	fmt.Println("\n📚 What's Next?")
	fmt.Println("- Backtracking Algorithms (N-Queens, Sudoku)")
	fmt.Println("- Advanced Graph Algorithms (MST, Network Flow)")
	fmt.Println("- Computational Geometry")
	fmt.Println("- Number Theory Algorithms")
}
