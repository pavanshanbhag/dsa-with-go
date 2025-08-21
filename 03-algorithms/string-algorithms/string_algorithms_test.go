package string_algorithms

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// =============================================================================
// 1. PATTERN MATCHING ALGORITHM TESTS
// =============================================================================

func TestNaivePatternSearch(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected []int
	}{
		{"AABAACAADAABAAABAA", "AABA", []int{0, 9, 13}},
		{"ABAAABCDABABCABCABCDAB", "ABCAB", []int{10, 13}}, // Two occurrences at positions 10 and 13
		{"HELLO WORLD", "LLO", []int{2}},
		{"ABABCABABA", "ABABA", []int{5}},
		{"", "A", []int{}},
		{"ABC", "", []int{}},
		{"SAME", "SAME", []int{0}},
		{"NO MATCH", "XYZ", []int{}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := NaivePatternSearch(test.text, test.pattern)
			if !slicesEqual(result, test.expected) {
				t.Errorf("NaivePatternSearch(%q, %q) = %v, want %v",
					test.text, test.pattern, result, test.expected)
			}
		})
	}
}

func TestKMPPatternSearch(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected []int
	}{
		{"AABAACAADAABAAABAA", "AABA", []int{0, 9, 13}},
		{"ABAAABCDABABCABCABCDAB", "ABCAB", []int{10, 13}}, // Two occurrences at positions 10 and 13
		{"HELLO WORLD", "LLO", []int{2}},
		{"ABABCABABA", "ABABA", []int{5}},
		{"", "A", []int{}},
		{"ABC", "", []int{}},
		{"SAME", "SAME", []int{0}},
		{"NO MATCH", "XYZ", []int{}},
		{"ABABABABAB", "ABAB", []int{0, 2, 4, 6}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := KMPPatternSearch(test.text, test.pattern)
			if !slicesEqual(result, test.expected) {
				t.Errorf("KMPPatternSearch(%q, %q) = %v, want %v",
					test.text, test.pattern, result, test.expected)
			}
		})
	}
}

func TestRabinKarpSearch(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected []int
	}{
		{"AABAACAADAABAAABAA", "AABA", []int{0, 9, 13}},
		{"ABAAABCDABABCABCABCDAB", "ABCAB", []int{10, 13}}, // Two occurrences at positions 10 and 13
		{"HELLO WORLD", "LLO", []int{2}},
		{"ABABCABABA", "ABABA", []int{5}},
		{"", "A", []int{}},
		{"ABC", "", []int{}},
		{"SAME", "SAME", []int{0}},
		{"NO MATCH", "XYZ", []int{}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := RabinKarpSearch(test.text, test.pattern)
			if !slicesEqual(result, test.expected) {
				t.Errorf("RabinKarpSearch(%q, %q) = %v, want %v",
					test.text, test.pattern, result, test.expected)
			}
		})
	}
}

func TestBoyerMooreSearch(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected []int
	}{
		{"AABAACAADAABAAABAA", "AABA", []int{0, 9, 13}},
		{"ABAAABCDABABCABCABCDAB", "ABCAB", []int{10, 13}}, // Two occurrences at positions 10 and 13
		{"HELLO WORLD", "LLO", []int{2}},
		{"ABABCABABA", "ABABA", []int{5}},
		{"", "A", []int{}},
		{"ABC", "", []int{}},
		{"SAME", "SAME", []int{0}},
		{"NO MATCH", "XYZ", []int{}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := BoyerMooreSearch(test.text, test.pattern)
			if !slicesEqual(result, test.expected) {
				t.Errorf("BoyerMooreSearch(%q, %q) = %v, want %v",
					test.text, test.pattern, result, test.expected)
			}
		})
	}
}

func TestZPatternSearch(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected []int
	}{
		{"AABAACAADAABAAABAA", "AABA", []int{0, 9, 13}},
		{"ABAAABCDABABCABCABCDAB", "ABCAB", []int{10, 13}}, // Two occurrences at positions 10 and 13
		{"HELLO WORLD", "LLO", []int{2}},
		{"ABABCABABA", "ABABA", []int{5}},
		{"", "A", []int{}},
		{"ABC", "", []int{}},
		{"SAME", "SAME", []int{0}},
		{"NO MATCH", "XYZ", []int{}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := ZPatternSearch(test.text, test.pattern)
			if !slicesEqual(result, test.expected) {
				t.Errorf("ZPatternSearch(%q, %q) = %v, want %v",
					test.text, test.pattern, result, test.expected)
			}
		})
	}
}

// =============================================================================
// 2. TRIE DATA STRUCTURE TESTS
// =============================================================================

func TestTrieBasicOperations(t *testing.T) {
	trie := NewTrie()

	// Test empty trie
	if trie.Search("test") {
		t.Error("Empty trie should not contain 'test'")
	}

	if trie.Size() != 0 {
		t.Errorf("Empty trie size should be 0, got %d", trie.Size())
	}

	// Test insertion and search
	words := []string{"apple", "app", "apricot", "banana", "band", "bandana"}
	for _, word := range words {
		trie.Insert(word)
	}

	if trie.Size() != len(words) {
		t.Errorf("Trie size should be %d, got %d", len(words), trie.Size())
	}

	// Test all words exist
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Word '%s' should exist in trie", word)
		}
	}

	// Test non-existent words
	nonExistent := []string{"ap", "applic", "ban", "orange"}
	for _, word := range nonExistent {
		if trie.Search(word) {
			t.Errorf("Word '%s' should not exist in trie", word)
		}
	}

	// Test prefix search
	if !trie.StartsWith("app") {
		t.Error("Trie should have words starting with 'app'")
	}

	if !trie.StartsWith("ban") {
		t.Error("Trie should have words starting with 'ban'")
	}

	if trie.StartsWith("orange") {
		t.Error("Trie should not have words starting with 'orange'")
	}
}

func TestTrieWithValues(t *testing.T) {
	trie := NewTrie()

	// Insert words with values
	trie.InsertWithValue("apple", 1)
	trie.InsertWithValue("app", 2)
	trie.InsertWithValue("apricot", 3)

	// Test value retrieval
	if val, exists := trie.GetValue("apple"); !exists || val != 1 {
		t.Errorf("Expected value 1 for 'apple', got %v (exists: %v)", val, exists)
	}

	if val, exists := trie.GetValue("app"); !exists || val != 2 {
		t.Errorf("Expected value 2 for 'app', got %v (exists: %v)", val, exists)
	}

	if _, exists := trie.GetValue("ap"); exists {
		t.Error("Should not find value for 'ap'")
	}
}

func TestTrieDeletion(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "apricot"}

	for _, word := range words {
		trie.Insert(word)
	}

	// Delete "app"
	if !trie.Delete("app") {
		t.Error("Should successfully delete 'app'")
	}

	if trie.Search("app") {
		t.Error("'app' should not exist after deletion")
	}

	if !trie.Search("apple") {
		t.Error("'apple' should still exist after deleting 'app'")
	}

	if trie.Size() != 2 {
		t.Errorf("Size should be 2 after deletion, got %d", trie.Size())
	}

	// Try to delete non-existent word
	if trie.Delete("nonexistent") {
		t.Error("Should not successfully delete non-existent word")
	}
}

func TestTriePrefixSearch(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "card", "care", "careful", "cars", "carry"}

	for _, word := range words {
		trie.Insert(word)
	}

	// Test prefix "car"
	results := trie.GetWordsWithPrefix("car")
	expected := []string{"car", "card", "care", "careful", "cars", "carry"}

	if len(results) != len(expected) {
		t.Errorf("Expected %d words with prefix 'car', got %d", len(expected), len(results))
	}

	// Convert to set for easier comparison
	resultSet := make(map[string]bool)
	for _, word := range results {
		resultSet[word] = true
	}

	for _, word := range expected {
		if !resultSet[word] {
			t.Errorf("Expected word '%s' in results for prefix 'car'", word)
		}
	}

	// Test prefix with no matches
	noMatch := trie.GetWordsWithPrefix("xyz")
	if len(noMatch) != 0 {
		t.Errorf("Expected no words with prefix 'xyz', got %v", noMatch)
	}
}

// =============================================================================
// 3. STRING HASHING AND ROLLING HASH TESTS
// =============================================================================

func TestStringHasher(t *testing.T) {
	hasher := NewStringHasher()

	// Test basic hashing
	hash1 := hasher.Hash("hello")
	hash2 := hasher.Hash("hello")
	hash3 := hasher.Hash("world")

	if hash1 != hash2 {
		t.Error("Same strings should have same hash")
	}

	if hash1 == hash3 {
		t.Error("Different strings should likely have different hashes")
	}

	// Test empty string
	emptyHash := hasher.Hash("")
	if emptyHash != 0 {
		t.Errorf("Empty string hash should be 0, got %d", emptyHash)
	}
}

func TestRollingHash(t *testing.T) {
	text := "abcdefghij"
	windowSize := 3

	rh := NewRollingHash(text, windowSize)
	if rh == nil {
		t.Fatal("Failed to create rolling hash")
	}

	// Test initial window
	if rh.CurrentWindow() != "abc" {
		t.Errorf("Expected initial window 'abc', got '%s'", rh.CurrentWindow())
	}

	// Test rolling
	hasher := NewStringHasher()
	expectedWindows := []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij"}

	for i, expected := range expectedWindows {
		current := rh.CurrentWindow()
		if current != expected {
			t.Errorf("Window %d: expected '%s', got '%s'", i, expected, current)
		}

		// Verify hash matches direct computation
		expectedHash := hasher.Hash(current)
		if rh.CurrentHash() != expectedHash {
			t.Errorf("Hash mismatch for window '%s'", current)
		}

		if i < len(expectedWindows)-1 {
			if !rh.RollNext() {
				t.Errorf("Should be able to roll at position %d", i)
			}
		}
	}

	// Should not be able to roll further
	if rh.RollNext() {
		t.Error("Should not be able to roll beyond text length")
	}
}

// =============================================================================
// 4. ADVANCED STRING ALGORITHM TESTS
// =============================================================================

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"throne", "throne"}, "throne"},
		{[]string{"prefix"}, "prefix"},
		{[]string{}, ""},
		{[]string{"", "abc"}, ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := LongestCommonPrefix(test.input)
			if result != test.expected {
				t.Errorf("LongestCommonPrefix(%v) = %q, want %q",
					test.input, result, test.expected)
			}
		})
	}
}

func TestZAlgorithm(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"aabaaab", []int{0, 1, 0, 2, 3, 1, 0}}, // Fixed: position 3 should be 2, position 4 should be 3
		{"ababa", []int{0, 0, 3, 0, 1}},
		{"aaaa", []int{0, 3, 2, 1}},
		{"abcd", []int{0, 0, 0, 0}},
		{"a", []int{0}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := ZAlgorithm(test.input)
			if !slicesEqual(result, test.expected) {
				t.Errorf("ZAlgorithm(%q) = %v, want %v",
					test.input, result, test.expected)
			}
		})
	}
}

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"babad", "bab"}, // or "aba", both valid
		{"cbbd", "bb"},
		{"racecar", "racecar"},
		{"a", "a"},
		{"ac", "a"}, // or "c", both valid
		{"", ""},
		{"abcdef", "a"}, // or any single character
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := LongestPalindromicSubstring(test.input)

			// Check if result is a palindrome
			if !isPalindrome(result) {
				t.Errorf("Result '%s' is not a palindrome", result)
			}

			// Check if result is a substring of input
			if !strings.Contains(test.input, result) {
				t.Errorf("Result '%s' is not a substring of '%s'", result, test.input)
			}

			// For specific test cases, check exact match or length
			if test.input == "racecar" && result != "racecar" {
				t.Errorf("For 'racecar', expected 'racecar', got '%s'", result)
			}

			if test.input == "cbbd" && result != "bb" {
				t.Errorf("For 'cbbd', expected 'bb', got '%s'", result)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"listen", "silent", true},
		{"elbow", "below", true},
		{"study", "dusty", true},
		{"hello", "world", false},
		{"rat", "car", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "ba", true},
		{"abc", "def", false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := IsAnagram(test.s1, test.s2)
			if result != test.expected {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v",
					test.s1, test.s2, result, test.expected)
			}
		})
	}
}

func TestStringReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"racecar", "racecar"},
		{"GoLang", "gnaLoG"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			result := StringReverse(test.input)
			if result != test.expected {
				t.Errorf("StringReverse(%q) = %q, want %q",
					test.input, result, test.expected)
			}
		})
	}
}

// =============================================================================
// 5. PERFORMANCE BENCHMARKS
// =============================================================================

func BenchmarkPatternSearchAlgorithms(b *testing.B) {
	text := strings.Repeat("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1000)
	pattern := "MNOPQR"

	b.Run("Naive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NaivePatternSearch(text, pattern)
		}
	})

	b.Run("KMP", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			KMPPatternSearch(text, pattern)
		}
	})

	b.Run("RabinKarp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RabinKarpSearch(text, pattern)
		}
	})

	b.Run("BoyerMoore", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BoyerMooreSearch(text, pattern)
		}
	})

	b.Run("ZAlgorithm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ZPatternSearch(text, pattern)
		}
	})
}

func BenchmarkTrieOperations(b *testing.B) {
	words := generateWords(1000, 10)

	b.Run("Insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewTrie()
			for _, word := range words {
				trie.Insert(word)
			}
		}
	})

	b.Run("Search", func(b *testing.B) {
		trie := NewTrie()
		for _, word := range words {
			trie.Insert(word)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, word := range words {
				trie.Search(word)
			}
		}
	})
}

func BenchmarkStringHashing(b *testing.B) {
	hasher := NewStringHasher()
	text := strings.Repeat("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 100)

	b.Run("DirectHash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hasher.Hash(text)
		}
	})

	b.Run("FNVHash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FNVHash(text)
		}
	})

	b.Run("RollingHash", func(b *testing.B) {
		windowSize := 10
		for i := 0; i < b.N; i++ {
			rh := NewRollingHash(text, windowSize)
			for rh.RollNext() {
				_ = rh.CurrentHash()
			}
		}
	})
}

// =============================================================================
// 6. CONSISTENCY TESTS (Verify all algorithms produce same results)
// =============================================================================

func TestPatternSearchConsistency(t *testing.T) {
	testCases := []struct {
		text    string
		pattern string
	}{
		{"AABAACAADAABAAABAA", "AABA"},
		{"ABAAABCDABABCABCABCDAB", "ABCAB"},
		{"HELLO WORLD HELLO", "HELLO"},
		{"ABABABABAB", "ABAB"},
		{"MISSISSIPPI", "ISSI"},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("Consistency_%d", i), func(t *testing.T) {
			naive := NaivePatternSearch(test.text, test.pattern)
			kmp := KMPPatternSearch(test.text, test.pattern)
			rk := RabinKarpSearch(test.text, test.pattern)
			bm := BoyerMooreSearch(test.text, test.pattern)
			z := ZPatternSearch(test.text, test.pattern)

			if !slicesEqual(naive, kmp) {
				t.Errorf("Naive and KMP results differ: %v vs %v", naive, kmp)
			}

			if !slicesEqual(naive, rk) {
				t.Errorf("Naive and Rabin-Karp results differ: %v vs %v", naive, rk)
			}

			if !slicesEqual(naive, bm) {
				t.Errorf("Naive and Boyer-Moore results differ: %v vs %v", naive, bm)
			}

			if !slicesEqual(naive, z) {
				t.Errorf("Naive and Z-algorithm results differ: %v vs %v", naive, z)
			}
		})
	}
}

// =============================================================================
// 7. STRESS TESTS
// =============================================================================

func TestLargeInputPatternSearch(t *testing.T) {
	// Create large text with known pattern positions
	pattern := "TESTPATTERN"
	baseText := strings.Repeat("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 100)
	text := baseText + pattern + baseText + pattern + baseText

	expected := []int{len(baseText), len(baseText)*2 + len(pattern)}

	algorithms := map[string]func(string, string) []int{
		"KMP":        KMPPatternSearch,
		"RabinKarp":  RabinKarpSearch,
		"BoyerMoore": BoyerMooreSearch,
		"ZAlgorithm": ZPatternSearch,
	}

	for name, algo := range algorithms {
		t.Run(name, func(t *testing.T) {
			start := time.Now()
			result := algo(text, pattern)
			duration := time.Since(start)

			if !slicesEqual(result, expected) {
				t.Errorf("%s: expected %v, got %v", name, expected, result)
			}

			if duration > time.Second {
				t.Errorf("%s: took too long (%v) for large input", name, duration)
			}
		})
	}
}

func TestLargeTrieOperations(t *testing.T) {
	trie := NewTrie()
	words := generateWords(10000, 15)

	// Test large insertions
	start := time.Now()
	for _, word := range words {
		trie.Insert(word)
	}
	insertTime := time.Since(start)

	if insertTime > 5*time.Second {
		t.Errorf("Large trie insertion took too long: %v", insertTime)
	}

	// Test large searches
	start = time.Now()
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Word %s should exist in trie", word)
		}
	}
	searchTime := time.Since(start)

	if searchTime > 2*time.Second {
		t.Errorf("Large trie search took too long: %v", searchTime)
	}

	if trie.Size() != len(words) {
		t.Errorf("Trie size should be %d, got %d", len(words), trie.Size())
	}
}

// =============================================================================
// 8. HELPER FUNCTIONS
// =============================================================================

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func generateWords(count, maxLength int) []string {
	words := make([]string, count)
	chars := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < count; i++ {
		length := (i % maxLength) + 1
		word := make([]byte, length)
		for j := 0; j < length; j++ {
			// Use a more complex pattern to ensure uniqueness
			word[j] = chars[(i*7+j*13+i/maxLength)%len(chars)]
		}
		words[i] = string(word)
	}

	// Ensure uniqueness by using a map
	uniqueWords := make(map[string]bool)
	result := make([]string, 0, count)

	for _, word := range words {
		if !uniqueWords[word] {
			uniqueWords[word] = true
			result = append(result, word)
		}
	}

	// Fill with additional unique words if needed
	counter := 0
	for len(result) < count {
		word := fmt.Sprintf("word_%d", counter)
		if !uniqueWords[word] {
			uniqueWords[word] = true
			result = append(result, word)
		}
		counter++
	}

	return result
}
