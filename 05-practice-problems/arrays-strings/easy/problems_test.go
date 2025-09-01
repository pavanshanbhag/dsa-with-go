package easy

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"Basic case", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Multiple solutions", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Same number twice", []int{3, 3}, 6, []int{0, 1}},
		{"Negative numbers", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test brute force
			result1 := TwoSumBruteForce(tt.nums, tt.target)
			if !reflect.DeepEqual(result1, tt.expected) {
				t.Errorf("TwoSumBruteForce() = %v, want %v", result1, tt.expected)
			}

			// Test hash map
			result2 := TwoSumHashMap(tt.nums, tt.target)
			if !reflect.DeepEqual(result2, tt.expected) {
				t.Errorf("TwoSumHashMap() = %v, want %v", result2, tt.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid palindrome", "A man, a plan, a canal: Panama", true},
		{"Not palindrome", "race a car", false},
		{"Empty string", "", true},
		{"Single character", "a", true},
		{"Only spaces", " ", true},
		{"Numbers and letters", "Madam, I'm Adam", true},
		{"Mixed case", "Was it a car or a cat I saw?", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		expectedLength int
		expectedArray  []int
	}{
		{"Basic case", []int{1, 1, 2}, 2, []int{1, 2}},
		{"Multiple duplicates", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{"No duplicates", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{"All same", []int{1, 1, 1, 1}, 1, []int{1}},
		{"Empty array", []int{}, 0, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since the function modifies the array
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			length := RemoveDuplicates(nums)
			if length != tt.expectedLength {
				t.Errorf("RemoveDuplicates() length = %v, want %v", length, tt.expectedLength)
			}

			if length > 0 && !reflect.DeepEqual(nums[:length], tt.expectedArray) {
				t.Errorf("RemoveDuplicates() array = %v, want %v", nums[:length], tt.expectedArray)
			}
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"Basic case", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"All zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"No zeros", []int{1, 2, 3}, []int{1, 2, 3}},
		{"Zeros at end", []int{1, 2, 0, 0}, []int{1, 2, 0, 0}},
		{"Single element", []int{0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since the function modifies the array
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			MoveZeroes(nums)
			if !reflect.DeepEqual(nums, tt.expected) {
				t.Errorf("MoveZeroes() = %v, want %v", nums, tt.expected)
			}
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"Has duplicates", []int{1, 2, 3, 1}, true},
		{"No duplicates", []int{1, 2, 3, 4}, false},
		{"All same", []int{1, 1, 1, 1}, true},
		{"Empty array", []int{}, false},
		{"Single element", []int{1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test hash set method
			result1 := ContainsDuplicateHashSet(tt.nums)
			if result1 != tt.expected {
				t.Errorf("ContainsDuplicateHashSet() = %v, want %v", result1, tt.expected)
			}

			// Test sort method (make copy since it modifies array)
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			result2 := ContainsDuplicateSort(numsCopy)
			if result2 != tt.expected {
				t.Errorf("ContainsDuplicateSort() = %v, want %v", result2, tt.expected)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{"Valid anagram", "anagram", "nagaram", true},
		{"Not anagram", "rat", "car", false},
		{"Same string", "listen", "listen", true},
		{"Different lengths", "abc", "abcd", false},
		{"Empty strings", "", "", true},
		{"Single char", "a", "a", true},
		{"Case sensitive", "Anagram", "nagaram", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test sort method
			result1 := IsAnagramSort(tt.s, tt.t)
			if result1 != tt.expected {
				t.Errorf("IsAnagramSort(%q, %q) = %v, want %v", tt.s, tt.t, result1, tt.expected)
			}

			// Test frequency method
			result2 := IsAnagramFrequency(tt.s, tt.t)
			if result2 != tt.expected {
				t.Errorf("IsAnagramFrequency(%q, %q) = %v, want %v", tt.s, tt.t, result2, tt.expected)
			}
		})
	}
}

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"Basic case", "leetcode", 0},
		{"Second char unique", "loveleetcode", 2},
		{"No unique chars", "aabb", -1},
		{"Single char", "a", 0},
		{"All unique", "abc", 0},
		{"Last char unique", "aab", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstUniqChar(tt.s)
			if result != tt.expected {
				t.Errorf("FirstUniqChar(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{"Basic case", []string{"flower", "flow", "flight"}, "fl"},
		{"No common prefix", []string{"dog", "racecar", "car"}, ""},
		{"Single string", []string{"alone"}, "alone"},
		{"Empty array", []string{}, ""},
		{"Empty string in array", []string{"", "b"}, ""},
		{"All same", []string{"test", "test", "test"}, "test"},
		{"One longer", []string{"abc", "ab", "abcd"}, "ab"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonPrefix(tt.strs)
			if result != tt.expected {
				t.Errorf("LongestCommonPrefix(%v) = %q, want %q", tt.strs, result, tt.expected)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{"Basic case", []byte("hello"), []byte("olleh")},
		{"Single char", []byte("a"), []byte("a")},
		{"Empty string", []byte(""), []byte("")},
		{"Palindrome", []byte("aba"), []byte("aba")},
		{"Even length", []byte("abcd"), []byte("dcba")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since the function modifies the slice
			input := make([]byte, len(tt.input))
			copy(input, tt.input)

			ReverseString(input)
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("ReverseString() = %v, want %v", input, tt.expected)
			}
		})
	}
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{"Basic case", []int{1, 2, 3}, []int{1, 2, 4}},
		{"Carry over", []int{1, 2, 9}, []int{1, 3, 0}},
		{"All nines", []int{9, 9, 9}, []int{1, 0, 0, 0}},
		{"Single digit", []int{9}, []int{1, 0}},
		{"No carry", []int{1, 2, 3, 4}, []int{1, 2, 3, 5}},
		{"Zero", []int{0}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since we don't want to modify the test data
			digits := make([]int, len(tt.digits))
			copy(digits, tt.digits)

			result := PlusOne(digits)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PlusOne(%v) = %v, want %v", tt.digits, result, tt.expected)
			}
		})
	}
}

// Benchmark tests for performance comparison
func BenchmarkTwoSumBruteForce(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 6, 8, 9, 1, 4}
	target := 9

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSumBruteForce(nums, target)
	}
}

func BenchmarkTwoSumHashMap(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 6, 8, 9, 1, 4}
	target := 9

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSumHashMap(nums, target)
	}
}

func BenchmarkContainsDuplicateHashSet(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicateHashSet(nums)
	}
}

func BenchmarkContainsDuplicateSort(b *testing.B) {
	baseNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nums := make([]int, len(baseNums))
		copy(nums, baseNums)
		ContainsDuplicateSort(nums)
	}
}

func BenchmarkIsAnagramSort(b *testing.B) {
	s, t := "anagram", "nagaram"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsAnagramSort(s, t)
	}
}

func BenchmarkIsAnagramFrequency(b *testing.B) {
	s, t := "anagram", "nagaram"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsAnagramFrequency(s, t)
	}
}
