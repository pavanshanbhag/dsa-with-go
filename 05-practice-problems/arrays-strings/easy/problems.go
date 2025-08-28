package easy

import (
	"fmt"
	"sort"
	"strings"
)

// ====================================================================
// PROBLEM 1: Two Sum
// Given an array of integers nums and an integer target, return indices
// of the two numbers such that they add up to target.
// ====================================================================

// TwoSumBruteForce solves using nested loops - O(n²) time, O(1) space
func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TwoSumHashMap solves using hash map - O(n) time, O(n) space
func TwoSumHashMap(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

// ====================================================================
// PROBLEM 2: Valid Palindrome
// A phrase is a palindrome if, after converting all uppercase letters
// into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward.
// ====================================================================

// IsPalindrome checks if string is palindrome using two pointers
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric characters
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		// Compare characters (case insensitive)
		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLowerCase(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// ====================================================================
// PROBLEM 3: Remove Duplicates from Sorted Array
// Given a sorted array nums, remove duplicates in-place and return
// the new length.
// ====================================================================

// RemoveDuplicates uses two-pointer technique
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	writeIndex := 0

	for readIndex := 1; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != nums[writeIndex] {
			writeIndex++
			nums[writeIndex] = nums[readIndex]
		}
	}

	return writeIndex + 1
}

// ====================================================================
// PROBLEM 4: Move Zeroes
// Given an integer array nums, move all 0's to the end while
// maintaining the relative order of the non-zero elements.
// ====================================================================

// MoveZeroes moves all zeros to end using two pointers
func MoveZeroes(nums []int) {
	writeIndex := 0

	// Move all non-zero elements to front
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != 0 {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}

	// Fill remaining positions with zeros
	for writeIndex < len(nums) {
		nums[writeIndex] = 0
		writeIndex++
	}
}

// ====================================================================
// PROBLEM 5: Contains Duplicate
// Given an integer array nums, return true if any value appears
// at least twice in the array.
// ====================================================================

// ContainsDuplicateHashSet uses hash set - O(n) time, O(n) space
func ContainsDuplicateHashSet(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

// ContainsDuplicateSort uses sorting - O(n log n) time, O(1) space
func ContainsDuplicateSort(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

// ====================================================================
// PROBLEM 6: Valid Anagram
// Given two strings s and t, return true if t is an anagram of s.
// ====================================================================

// IsAnagramSort solves using sorting - O(n log n) time
func IsAnagramSort(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRunes := []rune(s)
	tRunes := []rune(t)

	sort.Slice(sRunes, func(i, j int) bool { return sRunes[i] < sRunes[j] })
	sort.Slice(tRunes, func(i, j int) bool { return tRunes[i] < tRunes[j] })

	return string(sRunes) == string(tRunes)
}

// IsAnagramFrequency solves using frequency counting - O(n) time
func IsAnagramFrequency(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	frequency := make(map[rune]int)

	// Count frequencies in first string
	for _, char := range s {
		frequency[char]++
	}

	// Subtract frequencies using second string
	for _, char := range t {
		frequency[char]--
		if frequency[char] < 0 {
			return false
		}
	}

	// Check if all frequencies are zero
	for _, count := range frequency {
		if count != 0 {
			return false
		}
	}

	return true
}

// ====================================================================
// PROBLEM 7: First Unique Character
// Given a string s, find the first non-repeating character and
// return its index. If it doesn't exist, return -1.
// ====================================================================

// FirstUniqChar finds first unique character using frequency map
func FirstUniqChar(s string) int {
	frequency := make(map[rune]int)

	// Count frequency of each character
	for _, char := range s {
		frequency[char]++
	}

	// Find first character with frequency 1
	for i, char := range s {
		if frequency[char] == 1 {
			return i
		}
	}

	return -1
}

// ====================================================================
// PROBLEM 8: Longest Common Prefix
// Write a function to find the longest common prefix string amongst
// an array of strings.
// ====================================================================

// LongestCommonPrefix finds common prefix using vertical scanning
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Use first string as reference
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		// Check if character matches in all strings
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

// ====================================================================
// PROBLEM 9: Reverse String
// Write a function that reverses a string in-place.
// ====================================================================

// ReverseString reverses string in-place using two pointers
func ReverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// ====================================================================
// PROBLEM 10: Plus One
// Given a non-empty array of decimal digits representing a non-negative
// integer, increment the integer by one.
// ====================================================================

// PlusOne adds one to number represented as array
func PlusOne(digits []int) []int {
	carry := 1

	// Process from right to left
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10

		if carry == 0 {
			break
		}
	}

	// If carry still exists, prepend it
	if carry > 0 {
		return append([]int{carry}, digits...)
	}

	return digits
}

// ====================================================================
// PROBLEM DEMONSTRATION AND TESTING
// ====================================================================

// DemonstrateEasyProblems shows all easy problem solutions
func DemonstrateEasyProblems() {
	fmt.Println("🎯 Easy Array & String Problems")
	fmt.Println("===============================")

	// Problem 1: Two Sum
	fmt.Println("\n1. Two Sum:")
	nums := []int{2, 7, 11, 15}
	target := 9
	result1 := TwoSumBruteForce(nums, target)
	result2 := TwoSumHashMap(nums, target)
	fmt.Printf("Array: %v, Target: %d\n", nums, target)
	fmt.Printf("Brute Force: %v, Hash Map: %v\n", result1, result2)

	// Problem 2: Valid Palindrome
	fmt.Println("\n2. Valid Palindrome:")
	testStr := "A man, a plan, a canal: Panama"
	isPalin := IsPalindrome(testStr)
	fmt.Printf("String: '%s'\n", testStr)
	fmt.Printf("Is Palindrome: %t\n", isPalin)

	// Problem 3: Remove Duplicates
	fmt.Println("\n3. Remove Duplicates:")
	sortedArr := []int{1, 1, 2, 2, 3, 3, 4}
	originalLen := len(sortedArr)
	newLen := RemoveDuplicates(sortedArr)
	fmt.Printf("Original: %v (length %d)\n", sortedArr[:originalLen], originalLen)
	fmt.Printf("After removal: %v (length %d)\n", sortedArr[:newLen], newLen)

	// Problem 4: Move Zeroes
	fmt.Println("\n4. Move Zeroes:")
	zeroArr := []int{0, 1, 0, 3, 12}
	fmt.Printf("Before: %v\n", zeroArr)
	MoveZeroes(zeroArr)
	fmt.Printf("After: %v\n", zeroArr)

	// Problem 5: Contains Duplicate
	fmt.Println("\n5. Contains Duplicate:")
	dupArr := []int{1, 2, 3, 1}
	noDupArr := []int{1, 2, 3, 4}
	fmt.Printf("Array %v: %t\n", dupArr, ContainsDuplicateHashSet(dupArr))
	fmt.Printf("Array %v: %t\n", noDupArr, ContainsDuplicateHashSet(noDupArr))

	// Problem 6: Valid Anagram
	fmt.Println("\n6. Valid Anagram:")
	s1, t1 := "anagram", "nagaram"
	s2, t2 := "rat", "car"
	fmt.Printf("'%s' & '%s': %t\n", s1, t1, IsAnagramFrequency(s1, t1))
	fmt.Printf("'%s' & '%s': %t\n", s2, t2, IsAnagramFrequency(s2, t2))

	// Problem 7: First Unique Character
	fmt.Println("\n7. First Unique Character:")
	uniqueStr := "leetcode"
	firstUniq := FirstUniqChar(uniqueStr)
	fmt.Printf("String: '%s'\n", uniqueStr)
	fmt.Printf("First unique character index: %d\n", firstUniq)
	if firstUniq >= 0 {
		fmt.Printf("Character: '%c'\n", rune(uniqueStr[firstUniq]))
	}

	// Problem 8: Longest Common Prefix
	fmt.Println("\n8. Longest Common Prefix:")
	prefixStrs := []string{"flower", "flow", "flight"}
	commonPrefix := LongestCommonPrefix(prefixStrs)
	fmt.Printf("Strings: %v\n", prefixStrs)
	fmt.Printf("Longest common prefix: '%s'\n", commonPrefix)

	// Problem 9: Reverse String
	fmt.Println("\n9. Reverse String:")
	reverseStr := []byte("hello")
	fmt.Printf("Before: %s\n", string(reverseStr))
	ReverseString(reverseStr)
	fmt.Printf("After: %s\n", string(reverseStr))

	// Problem 10: Plus One
	fmt.Println("\n10. Plus One:")
	digitArr := []int{1, 2, 3}
	result := PlusOne(digitArr)
	fmt.Printf("Original: %v\n", digitArr)
	fmt.Printf("Plus one: %v\n", result)

	// Edge case: all 9s
	nines := []int{9, 9, 9}
	ninesResult := PlusOne(nines)
	fmt.Printf("All 9s: %v -> %v\n", nines, ninesResult)
}

// ProblemComplexityAnalysis provides complexity analysis for all problems
func ProblemComplexityAnalysis() {
	fmt.Println("\n📊 Complexity Analysis Summary")
	fmt.Println("==============================")

	problems := []struct {
		name     string
		timeOpt  string
		spaceOpt string
		approach string
	}{
		{"Two Sum", "O(n)", "O(n)", "Hash Map"},
		{"Valid Palindrome", "O(n)", "O(1)", "Two Pointers"},
		{"Remove Duplicates", "O(n)", "O(1)", "Two Pointers"},
		{"Move Zeroes", "O(n)", "O(1)", "Two Pointers"},
		{"Contains Duplicate", "O(n)", "O(n)", "Hash Set"},
		{"Valid Anagram", "O(n)", "O(1)", "Frequency Count"},
		{"First Unique Char", "O(n)", "O(1)", "Frequency Map"},
		{"Longest Common Prefix", "O(S)", "O(1)", "Vertical Scan"},
		{"Reverse String", "O(n)", "O(1)", "Two Pointers"},
		{"Plus One", "O(n)", "O(1)", "Carry Propagation"},
	}

	fmt.Printf("%-25s %-10s %-10s %-20s\n", "Problem", "Time", "Space", "Best Approach")
	fmt.Println(strings.Repeat("-", 70))

	for _, p := range problems {
		fmt.Printf("%-25s %-10s %-10s %-20s\n", p.name, p.timeOpt, p.spaceOpt, p.approach)
	}
}
