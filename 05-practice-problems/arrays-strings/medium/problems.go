package medium

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ====================================================================
// PROBLEM 1: 3Sum
// Given an integer array nums, return all unique triplets that sum to 0
// ====================================================================

// ThreeSum finds all unique triplets that sum to zero
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for second element
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for third element
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// ====================================================================
// PROBLEM 2: Container With Most Water
// Find two lines that together with x-axis form container holding most water
// ====================================================================

// MaxArea finds the maximum water area using two pointers
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// Calculate current water area
		width := right - left
		minHeight := min(height[left], height[right])
		water := width * minHeight
		maxWater = max(maxWater, water)

		// Move pointer with smaller height
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ====================================================================
// PROBLEM 3: Longest Substring Without Repeating Characters
// Find length of longest substring without repeating characters
// ====================================================================

// LengthOfLongestSubstring uses sliding window technique
func LengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		// If character seen before and within current window
		if lastIndex, exists := charMap[s[right]]; exists && lastIndex >= left {
			left = lastIndex + 1
		}

		charMap[s[right]] = right
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

// ====================================================================
// PROBLEM 4: Group Anagrams
// Group strings that are anagrams of each other
// ====================================================================

// GroupAnagrams groups anagrams using sorted strings as keys
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// Sort characters to create key
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
		key := string(chars)

		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// GroupAnagramsFrequency groups anagrams using character frequency
func GroupAnagramsFrequency(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// Create frequency signature
		freq := make([]int, 26)
		for _, char := range str {
			freq[char-'a']++
		}

		// Convert frequency to string key
		key := ""
		for i, count := range freq {
			if count > 0 {
				key += string(rune('a'+i)) + strconv.Itoa(count)
			}
		}

		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// ====================================================================
// PROBLEM 5: Longest Palindromic Substring
// Find the longest palindromic substring
// ====================================================================

// LongestPalindrome finds longest palindromic substring using expand around centers
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		// Check for odd-length palindromes (center at i)
		len1 := expandAroundCenter(s, i, i)
		// Check for even-length palindromes (center between i and i+1)
		len2 := expandAroundCenter(s, i, i+1)

		currentLen := max(len1, len2)
		if currentLen > maxLen {
			maxLen = currentLen
			start = i - (currentLen-1)/2
		}
	}

	return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// ====================================================================
// PROBLEM 6: Product of Array Except Self
// Return array where each element is product of all other elements
// ====================================================================

// ProductExceptSelf calculates product except self without division
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Left pass: multiply all elements to the left
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Right pass: multiply all elements to the right
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}

// ====================================================================
// PROBLEM 7: Rotate Array
// Rotate array to the right by k steps
// ====================================================================

// RotateArray rotates array using reversal technique
func RotateArray(nums []int, k int) {
	n := len(nums)
	k = k % n // Handle k > n

	// Reverse entire array
	reverse(nums, 0, n-1)
	// Reverse first k elements
	reverse(nums, 0, k-1)
	// Reverse remaining elements
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// RotateArrayCyclic rotates using cyclic replacements
func RotateArrayCyclic(nums []int, k int) {
	n := len(nums)
	k = k % n
	count := 0

	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]

		for {
			next := (current + k) % n
			nums[next], prev = prev, nums[next]
			current = next
			count++

			if start == current {
				break
			}
		}
	}
}

// ====================================================================
// PROBLEM 8: Find All Duplicates in Array
// Find all duplicates where 1 ≤ a[i] ≤ n (array length)
// ====================================================================

// FindDuplicates uses array indices as hash - O(n) time, O(1) space
func FindDuplicates(nums []int) []int {
	result := []int{}

	for _, num := range nums {
		index := abs(num) - 1

		// If number at index is negative, we've seen this number before
		if nums[index] < 0 {
			result = append(result, abs(num))
		} else {
			// Mark as seen by making negative
			nums[index] = -nums[index]
		}
	}

	// Restore original array (optional)
	for i := range nums {
		nums[i] = abs(nums[i])
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ====================================================================
// PROBLEM 9: Subarray Sum Equals K
// Find number of continuous subarrays whose sum equals k
// ====================================================================

// SubarraySum uses prefix sum with hash map
func SubarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	sumMap := make(map[int]int)
	sumMap[0] = 1 // Empty subarray

	for _, num := range nums {
		prefixSum += num

		// Check if (prefixSum - k) exists
		if freq, exists := sumMap[prefixSum-k]; exists {
			count += freq
		}

		sumMap[prefixSum]++
	}

	return count
}

// ====================================================================
// PROBLEM 10: Spiral Matrix
// Return elements of matrix in spiral order
// ====================================================================

// SpiralOrder traverses matrix in spiral pattern
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	rows, cols := len(matrix), len(matrix[0])
	result := make([]int, 0, rows*cols)

	top, bottom := 0, rows-1
	left, right := 0, cols-1

	for top <= bottom && left <= right {
		// Traverse right
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		// Traverse down
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		// Traverse left (if we still have rows)
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}

		// Traverse up (if we still have columns)
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}

	return result
}

// ====================================================================
// PROBLEM DEMONSTRATION
// ====================================================================

// DemonstrateMediumProblems shows all medium problem solutions
func DemonstrateMediumProblems() {
	fmt.Println("🎯 Medium Array & String Problems")
	fmt.Println("=================================")

	// Problem 1: 3Sum
	fmt.Println("\n1. Three Sum:")
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	triplets := ThreeSum(nums1)
	fmt.Printf("Array: %v\n", nums1)
	fmt.Printf("Triplets summing to 0: %v\n", triplets)

	// Problem 2: Container With Most Water
	fmt.Println("\n2. Container With Most Water:")
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxWater := MaxArea(heights)
	fmt.Printf("Heights: %v\n", heights)
	fmt.Printf("Maximum water area: %d\n", maxWater)

	// Problem 3: Longest Substring
	fmt.Println("\n3. Longest Substring Without Repeating Characters:")
	testStr := "abcabcbb"
	maxLen := LengthOfLongestSubstring(testStr)
	fmt.Printf("String: '%s'\n", testStr)
	fmt.Printf("Longest substring length: %d\n", maxLen)

	// Problem 4: Group Anagrams
	fmt.Println("\n4. Group Anagrams:")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := GroupAnagrams(strs)
	fmt.Printf("Strings: %v\n", strs)
	fmt.Printf("Grouped anagrams: %v\n", groups)

	// Problem 5: Longest Palindromic Substring
	fmt.Println("\n5. Longest Palindromic Substring:")
	palindromeStr := "babad"
	longest := LongestPalindrome(palindromeStr)
	fmt.Printf("String: '%s'\n", palindromeStr)
	fmt.Printf("Longest palindrome: '%s'\n", longest)

	// Problem 6: Product Except Self
	fmt.Println("\n6. Product of Array Except Self:")
	productNums := []int{1, 2, 3, 4}
	products := ProductExceptSelf(productNums)
	fmt.Printf("Original: %v\n", productNums)
	fmt.Printf("Products: %v\n", products)

	// Problem 7: Rotate Array
	fmt.Println("\n7. Rotate Array:")
	rotateNums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Printf("Before rotation: %v (k=%d)\n", rotateNums, k)
	RotateArray(rotateNums, k)
	fmt.Printf("After rotation: %v\n", rotateNums)

	// Problem 8: Find Duplicates
	fmt.Println("\n8. Find All Duplicates:")
	dupNums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	duplicates := FindDuplicates(dupNums)
	fmt.Printf("Array: %v\n", dupNums)
	fmt.Printf("Duplicates: %v\n", duplicates)

	// Problem 9: Subarray Sum
	fmt.Println("\n9. Subarray Sum Equals K:")
	sumNums := []int{1, 1, 1}
	target := 2
	subarrayCount := SubarraySum(sumNums, target)
	fmt.Printf("Array: %v, Target: %d\n", sumNums, target)
	fmt.Printf("Subarrays with sum %d: %d\n", target, subarrayCount)

	// Problem 10: Spiral Matrix
	fmt.Println("\n10. Spiral Matrix:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	spiral := SpiralOrder(matrix)
	fmt.Printf("Matrix:\n")
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
	fmt.Printf("Spiral order: %v\n", spiral)
}

// ProblemComplexityAnalysis provides complexity analysis for medium problems
func ProblemComplexityAnalysis() {
	fmt.Println("\n📊 Medium Problems Complexity Analysis")
	fmt.Println("=====================================")

	problems := []struct {
		name     string
		timeOpt  string
		spaceOpt string
		approach string
	}{
		{"3Sum", "O(n²)", "O(1)", "Sort + Two Pointers"},
		{"Container With Most Water", "O(n)", "O(1)", "Two Pointers"},
		{"Longest Substring", "O(n)", "O(min(m,n))", "Sliding Window"},
		{"Group Anagrams", "O(n*k log k)", "O(n*k)", "Sort Keys"},
		{"Longest Palindrome", "O(n²)", "O(1)", "Expand Around Centers"},
		{"Product Except Self", "O(n)", "O(1)", "Left/Right Products"},
		{"Rotate Array", "O(n)", "O(1)", "Reverse Technique"},
		{"Find Duplicates", "O(n)", "O(1)", "Index as Hash"},
		{"Subarray Sum", "O(n)", "O(n)", "Prefix Sum + HashMap"},
		{"Spiral Matrix", "O(m*n)", "O(1)", "Boundary Simulation"},
	}

	fmt.Printf("%-25s %-12s %-12s %-25s\n", "Problem", "Time", "Space", "Best Approach")
	fmt.Println(strings.Repeat("-", 80))

	for _, p := range problems {
		fmt.Printf("%-25s %-12s %-12s %-25s\n", p.name, p.timeOpt, p.spaceOpt, p.approach)
	}
}
