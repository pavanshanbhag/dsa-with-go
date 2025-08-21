// Package string_algorithms provides comprehensive implementations of string processing
// and pattern matching algorithms for DSA mastery.
//
// This module covers:
// - Pattern Matching: KMP, Rabin-Karp, Boyer-Moore algorithms
// - String Processing: Trie data structure, suffix arrays
// - Text Algorithms: Rolling hash, string hashing
// - Applications: Search engines, DNA analysis, text processing
package string_algorithms

import (
	"hash/fnv"
	"strings"
)

// =============================================================================
// 1. PATTERN MATCHING ALGORITHMS
// =============================================================================

// NaivePatternSearch performs brute force pattern matching
// Time: O(n*m), Space: O(1) where n=text length, m=pattern length
func NaivePatternSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	var matches []int
	n, m := len(text), len(pattern)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++
		}
		if j == m {
			matches = append(matches, i)
		}
	}

	return matches
}

// KMPPatternSearch implements Knuth-Morris-Pratt algorithm
// Time: O(n+m), Space: O(m) - optimal pattern matching
func KMPPatternSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	// Build LPS (Longest Proper Prefix which is also Suffix) array
	lps := buildLPSArray(pattern)
	var matches []int

	i, j := 0, 0 // i for text, j for pattern
	n, m := len(text), len(pattern)

	for i < n {
		if text[i] == pattern[j] {
			i++
			j++
		}

		if j == m {
			matches = append(matches, i-j)
			j = lps[j-1]
		} else if i < n && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return matches
}

// buildLPSArray constructs the LPS array for KMP algorithm
func buildLPSArray(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	length := 0 // length of previous longest prefix suffix
	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

// RabinKarpSearch implements Rabin-Karp rolling hash algorithm
// Time: O(n+m) average, O(n*m) worst case, Space: O(1)
func RabinKarpSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	const prime = 101 // A prime number for hash calculation
	var matches []int
	n, m := len(text), len(pattern)

	if m > n {
		return matches
	}

	// Calculate hash values
	patternHash := 0
	textHash := 0
	h := 1 // hash multiplier

	// Calculate h = pow(256, m-1) % prime
	for i := 0; i < m-1; i++ {
		h = (h * 256) % prime
	}

	// Calculate initial hash values
	for i := 0; i < m; i++ {
		patternHash = (256*patternHash + int(pattern[i])) % prime
		textHash = (256*textHash + int(text[i])) % prime
	}

	// Slide the pattern over text
	for i := 0; i <= n-m; i++ {
		// Check hash values
		if patternHash == textHash {
			// Check characters one by one
			if text[i:i+m] == pattern {
				matches = append(matches, i)
			}
		}

		// Calculate hash value for next window
		if i < n-m {
			textHash = (256*(textHash-int(text[i])*h) + int(text[i+m])) % prime
			if textHash < 0 {
				textHash += prime
			}
		}
	}

	return matches
}

// BoyerMooreSearch implements Boyer-Moore algorithm with bad character heuristic
// Time: O(n/m) best case, O(n*m) worst case, Space: O(256)
func BoyerMooreSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	var matches []int
	n, m := len(text), len(pattern)

	// Build bad character table
	badChar := make([]int, 256)
	for i := range badChar {
		badChar[i] = -1
	}
	for i := 0; i < m; i++ {
		badChar[pattern[i]] = i
	}

	shift := 0
	for shift <= n-m {
		j := m - 1

		// Keep reducing j while characters match
		for j >= 0 && pattern[j] == text[shift+j] {
			j--
		}

		if j < 0 {
			matches = append(matches, shift)
			// Shift pattern to align next character
			if shift+m < n {
				shift += m - badChar[text[shift+m]]
			} else {
				shift++
			}
		} else {
			// Shift pattern based on bad character heuristic
			shift += max(1, j-badChar[text[shift+j]])
		}
	}

	return matches
}

// =============================================================================
// 2. TRIE DATA STRUCTURE
// =============================================================================

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	value    interface{} // Optional: store associated value
}

// Trie represents a prefix tree data structure
type Trie struct {
	root *TrieNode
	size int
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
		size: 0,
	}
}

// Insert adds a word to the trie
// Time: O(m), Space: O(m) where m is word length
func (t *Trie) Insert(word string) {
	t.InsertWithValue(word, nil)
}

// InsertWithValue adds a word with associated value to the trie
func (t *Trie) InsertWithValue(word string, value interface{}) {
	node := t.root

	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}

	if !node.isEnd {
		t.size++
	}
	node.isEnd = true
	node.value = value
}

// Search checks if a word exists in the trie
// Time: O(m), Space: O(1) where m is word length
func (t *Trie) Search(word string) bool {
	node := t.root

	for _, char := range word {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}

	return node.isEnd
}

// StartsWith checks if any word in trie starts with given prefix
// Time: O(m), Space: O(1) where m is prefix length
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root

	for _, char := range prefix {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}

	return true
}

// GetValue retrieves the value associated with a word
func (t *Trie) GetValue(word string) (interface{}, bool) {
	node := t.root

	for _, char := range word {
		if node.children[char] == nil {
			return nil, false
		}
		node = node.children[char]
	}

	if node.isEnd {
		return node.value, true
	}
	return nil, false
}

// Delete removes a word from the trie
// Time: O(m), Space: O(m) for recursion where m is word length
func (t *Trie) Delete(word string) bool {
	if !t.Search(word) {
		return false
	}

	t.deleteHelper(t.root, word, 0)
	t.size--
	return true
}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		node.isEnd = false
		return len(node.children) == 0
	}

	char := rune(word[index])
	childNode := node.children[char]

	shouldDeleteChild := t.deleteHelper(childNode, word, index+1)

	if shouldDeleteChild {
		delete(node.children, char)
		return !node.isEnd && len(node.children) == 0
	}

	return false
}

// GetWordsWithPrefix returns all words that start with given prefix
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	var result []string
	node := t.root

	// Navigate to prefix end
	for _, char := range prefix {
		if node.children[char] == nil {
			return result
		}
		node = node.children[char]
	}

	// DFS to collect all words
	t.dfsCollectWords(node, prefix, &result)
	return result
}

func (t *Trie) dfsCollectWords(node *TrieNode, current string, result *[]string) {
	if node.isEnd {
		*result = append(*result, current)
	}

	for char, child := range node.children {
		t.dfsCollectWords(child, current+string(char), result)
	}
}

// Size returns the number of words in the trie
func (t *Trie) Size() int {
	return t.size
}

// =============================================================================
// 3. STRING HASHING AND ROLLING HASH
// =============================================================================

// StringHasher provides various string hashing utilities
type StringHasher struct {
	base  uint64
	prime uint64
}

// NewStringHasher creates a new string hasher
func NewStringHasher() *StringHasher {
	return &StringHasher{
		base:  256,
		prime: 1000000007,
	}
}

// Hash computes hash value for a string
func (sh *StringHasher) Hash(s string) uint64 {
	hash := uint64(0)
	for _, char := range s {
		hash = (hash*sh.base + uint64(char)) % sh.prime
	}
	return hash
}

// RollingHash provides efficient substring hash computation
type RollingHash struct {
	hasher *StringHasher
	text   string
	window int
	hash   uint64
	power  uint64
	start  int
}

// NewRollingHash creates a rolling hash for given text and window size
func NewRollingHash(text string, windowSize int) *RollingHash {
	if windowSize > len(text) {
		return nil
	}

	hasher := NewStringHasher()
	rh := &RollingHash{
		hasher: hasher,
		text:   text,
		window: windowSize,
		start:  0,
	}

	// Calculate initial hash and power
	rh.power = 1
	for i := 0; i < windowSize-1; i++ {
		rh.power = (rh.power * hasher.base) % hasher.prime
	}

	// Calculate initial window hash
	for i := 0; i < windowSize; i++ {
		rh.hash = (rh.hash*hasher.base + uint64(text[i])) % hasher.prime
	}

	return rh
}

// CurrentHash returns current window hash
func (rh *RollingHash) CurrentHash() uint64 {
	return rh.hash
}

// CurrentWindow returns current window string
func (rh *RollingHash) CurrentWindow() string {
	return rh.text[rh.start : rh.start+rh.window]
}

// RollNext moves window one position right and updates hash
func (rh *RollingHash) RollNext() bool {
	if rh.start+rh.window >= len(rh.text) {
		return false
	}

	// Remove leftmost character
	oldChar := uint64(rh.text[rh.start])
	rh.hash = (rh.hash - (oldChar*rh.power)%rh.hasher.prime + rh.hasher.prime) % rh.hasher.prime

	// Add new rightmost character
	newChar := uint64(rh.text[rh.start+rh.window])
	rh.hash = (rh.hash*rh.hasher.base + newChar) % rh.hasher.prime

	rh.start++
	return true
}

// =============================================================================
// 4. ADVANCED STRING ALGORITHMS
// =============================================================================

// LongestCommonPrefix finds the longest common prefix of an array of strings
// Time: O(S) where S is sum of all string lengths, Space: O(1)
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

// ZAlgorithm implements Z algorithm for pattern matching
// Time: O(n), Space: O(n) where n is string length
func ZAlgorithm(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r, k := 0, 0, 0

	for i := 1; i < n; i++ {
		if i > r {
			l, r = i, i
			for r < n && s[r-l] == s[r] {
				r++
			}
			z[i] = r - l
			r--
		} else {
			k = i - l
			if z[k] < r-i+1 {
				z[i] = z[k]
			} else {
				l = i
				for r < n && s[r-l] == s[r] {
					r++
				}
				z[i] = r - l
				r--
			}
		}
	}

	return z
}

// ZPatternSearch uses Z algorithm for pattern matching
func ZPatternSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	// Create concatenated string pattern$text
	concat := pattern + "$" + text
	z := ZAlgorithm(concat)

	var matches []int
	patternLen := len(pattern)

	for i := patternLen + 1; i < len(concat); i++ {
		if z[i] == patternLen {
			matches = append(matches, i-patternLen-1)
		}
	}

	return matches
}

// ManacherAlgorithm finds all palindromes in linear time
// Time: O(n), Space: O(n)
func ManacherAlgorithm(s string) []int {
	// Transform string to handle even-length palindromes
	processed := "#"
	for _, char := range s {
		processed += string(char) + "#"
	}

	n := len(processed)
	p := make([]int, n) // p[i] = radius of palindrome centered at i
	center, right := 0, 0

	for i := 0; i < n; i++ {
		mirror := 2*center - i

		if i < right {
			p[i] = min(right-i, p[mirror])
		}

		// Try to expand palindrome centered at i
		for i+p[i]+1 < n && i-p[i]-1 >= 0 &&
			processed[i+p[i]+1] == processed[i-p[i]-1] {
			p[i]++
		}

		// If palindrome centered at i extends past right, adjust center and right
		if i+p[i] > right {
			center, right = i, i+p[i]
		}
	}

	return p
}

// LongestPalindromicSubstring finds the longest palindromic substring
func LongestPalindromicSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}

	p := ManacherAlgorithm(s)
	maxLen, centerIndex := 0, 0

	for i, radius := range p {
		if radius > maxLen {
			maxLen = radius
			centerIndex = i
		}
	}

	// Convert back to original string coordinates
	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}

// =============================================================================
// 5. UTILITY FUNCTIONS
// =============================================================================

// Helper function for maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function for minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// StringReverse reverses a string efficiently
func StringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsAnagram checks if two strings are anagrams
func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charCount := make(map[rune]int)

	for _, char := range s1 {
		charCount[char]++
	}

	for _, char := range s2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}

	return true
}

// FNVHash provides FNV-1a hash for strings
func FNVHash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
