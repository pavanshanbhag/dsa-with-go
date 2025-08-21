package dp

import (
	"fmt"
	"testing"
)

// ============================================================================
// FIBONACCI TESTS
// ============================================================================

func TestFibonacciAlgorithms(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{10, 55},
		{15, 610},
	}

	algorithms := map[string]func(int) int{
		"Memoization": FibonacciMemoization,
		"Tabulation":  FibonacciTabulation,
		"Optimized":   FibonacciOptimized,
	}

	for name, algorithm := range algorithms {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("n=%d", tc.n), func(t *testing.T) {
					result := algorithm(tc.n)
					if result != tc.expected {
						t.Errorf("%s(%d) = %d, expected %d", name, tc.n, result, tc.expected)
					}
				})
			}
		})
	}
}

// Test naive version only for small inputs due to exponential time
func TestFibonacciNaive(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55}, // Still manageable
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc.n), func(t *testing.T) {
			result := FibonacciNaive(tc.n)
			if result != tc.expected {
				t.Errorf("FibonacciNaive(%d) = %d, expected %d", tc.n, result, tc.expected)
			}
		})
	}
}

// ============================================================================
// KNAPSACK TESTS
// ============================================================================

func TestKnapsack01Algorithms(t *testing.T) {
	testCases := []struct {
		name     string
		items    []Item
		capacity int
		expected int
	}{
		{
			name:     "Empty knapsack",
			items:    []Item{},
			capacity: 10,
			expected: 0,
		},
		{
			name:     "Zero capacity",
			items:    []Item{{Weight: 1, Value: 1}},
			capacity: 0,
			expected: 0,
		},
		{
			name: "Classic example",
			items: []Item{
				{Weight: 10, Value: 60, Name: "Item1"},
				{Weight: 20, Value: 100, Name: "Item2"},
				{Weight: 30, Value: 120, Name: "Item3"},
			},
			capacity: 50,
			expected: 220, // Item2 + Item3
		},
		{
			name: "All items fit",
			items: []Item{
				{Weight: 1, Value: 1},
				{Weight: 2, Value: 4},
				{Weight: 3, Value: 7},
			},
			capacity: 10,
			expected: 12, // All items
		},
		{
			name: "Single item exceeds capacity",
			items: []Item{
				{Weight: 100, Value: 1000},
				{Weight: 1, Value: 1},
			},
			capacity: 50,
			expected: 1, // Only second item
		},
		{
			name: "Optimal selection",
			items: []Item{
				{Weight: 2, Value: 3},
				{Weight: 3, Value: 4},
				{Weight: 4, Value: 5},
				{Weight: 5, Value: 6},
			},
			capacity: 5,
			expected: 7, // Items 1 and 2
		},
	}

	algorithms := map[string]func([]Item, int) int{
		"Memoization": Knapsack01Memoization,
		"Tabulation":  Knapsack01Tabulation,
		"Optimized":   Knapsack01Optimized,
	}

	for name, algorithm := range algorithms {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					result := algorithm(tc.items, tc.capacity)
					if result != tc.expected {
						t.Errorf("%s: got %d, expected %d", name, result, tc.expected)
					}
				})
			}
		})
	}
}

func TestKnapsackUnbounded(t *testing.T) {
	testCases := []struct {
		name     string
		items    []Item
		capacity int
		expected int
	}{
		{
			name:     "Empty",
			items:    []Item{},
			capacity: 10,
			expected: 0,
		},
		{
			name: "Unlimited usage",
			items: []Item{
				{Weight: 1, Value: 1},
				{Weight: 3, Value: 4},
				{Weight: 4, Value: 5},
			},
			capacity: 7,
			expected: 9, // Use item 2 twice + item 1 once = 4+4+1 = 9
		},
		{
			name: "Best item repeated",
			items: []Item{
				{Weight: 2, Value: 3},
				{Weight: 5, Value: 7},
			},
			capacity: 10,
			expected: 15, // Use first item 5 times (5*3 = 15) vs second item twice (2*7 = 14)
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := KnapsackUnbounded(tc.items, tc.capacity)
			if result != tc.expected {
				t.Errorf("KnapsackUnbounded: got %d, expected %d", result, tc.expected)
			}
		})
	}
}

// ============================================================================
// LCS TESTS
// ============================================================================

func TestLCSAlgorithms(t *testing.T) {
	testCases := []struct {
		text1    string
		text2    string
		expected int
	}{
		{"", "", 0},
		{"abc", "", 0},
		{"", "abc", 0},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abcde", "ace", 3},
		{"abc", "def", 0},
		{"ABCDGH", "AEDFHR", 3},         // ADH
		{"AGGTAB", "GXTXAYB", 4},        // GTAB
		{"programming", "algorithm", 3}, // "gri" or "alm"
		{"LONGEST", "STONE", 3},         // "ONE"
	}

	algorithms := map[string]func(string, string) int{
		"Memoization": LCSMemoization,
		"Tabulation":  LCSTabulation,
		"Optimized":   LCSOptimized,
	}

	for name, algorithm := range algorithms {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("%s_vs_%s", tc.text1, tc.text2), func(t *testing.T) {
					result := algorithm(tc.text1, tc.text2)
					if result != tc.expected {
						t.Errorf("%s('%s', '%s') = %d, expected %d",
							name, tc.text1, tc.text2, result, tc.expected)
					}
				})
			}
		})
	}
}

// ============================================================================
// LIS TESTS
// ============================================================================

func TestLISAlgorithms(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "Increasing sequence",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "Classic example",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4, // [2, 3, 7, 18] or [2, 3, 7, 101]
		},
		{
			name:     "All same elements",
			nums:     []int{7, 7, 7, 7},
			expected: 1,
		},
		{
			name:     "Complex case",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4, // [0, 1, 2, 3]
		},
	}

	algorithms := map[string]func([]int) int{
		"Tabulation":   LISTabulation,
		"BinarySearch": LISBinarySearch,
	}

	for name, algorithm := range algorithms {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					result := algorithm(tc.nums)
					if result != tc.expected {
						t.Errorf("%s(%v) = %d, expected %d", name, tc.nums, result, tc.expected)
					}
				})
			}
		})
	}
}

// ============================================================================
// EDIT DISTANCE TESTS
// ============================================================================

func TestEditDistanceAlgorithms(t *testing.T) {
	testCases := []struct {
		word1    string
		word2    string
		expected int
	}{
		{"", "", 0},
		{"", "abc", 3},
		{"abc", "", 3},
		{"horse", "ros", 3},           // replace h->r, remove o, remove e
		{"intention", "execution", 5}, // Classic example
		{"abc", "abc", 0},
		{"abc", "def", 3},        // replace all
		{"kitten", "sitting", 3}, // substitute k->s, substitute e->i, insert g
	}

	algorithms := map[string]func(string, string) int{
		"Tabulation": EditDistanceTabulation,
		"Optimized":  EditDistanceOptimized,
	}

	for name, algorithm := range algorithms {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("%s_to_%s", tc.word1, tc.word2), func(t *testing.T) {
					result := algorithm(tc.word1, tc.word2)
					if result != tc.expected {
						t.Errorf("%s('%s', '%s') = %d, expected %d",
							name, tc.word1, tc.word2, result, tc.expected)
					}
				})
			}
		})
	}
}

// ============================================================================
// COIN CHANGE TESTS
// ============================================================================

func TestCoinChangeMinCoins(t *testing.T) {
	testCases := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Zero amount",
			coins:    []int{1, 2, 5},
			amount:   0,
			expected: 0,
		},
		{
			name:     "Standard case",
			coins:    []int{1, 3, 4},
			amount:   6,
			expected: 2, // 3 + 3
		},
		{
			name:     "Impossible",
			coins:    []int{2, 4},
			amount:   3,
			expected: -1,
		},
		{
			name:     "Single coin",
			coins:    []int{1},
			amount:   5,
			expected: 5,
		},
		{
			name:     "Optimal choice",
			coins:    []int{1, 5, 10, 21, 25},
			amount:   63,
			expected: 3, // 25 + 25 + 13 but 13 needs optimization: 25 + 21 + 10 + 5 + 1 + 1 = 6 coins, better: 25 + 25 + 10 + 3*1 = 5 coins, optimal: 3*21 = 3 coins
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CoinChangeMinCoins(tc.coins, tc.amount)
			if result != tc.expected {
				t.Errorf("CoinChangeMinCoins(%v, %d) = %d, expected %d",
					tc.coins, tc.amount, result, tc.expected)
			}
		})
	}
}

func TestCoinChangeCountWays(t *testing.T) {
	testCases := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Zero amount",
			coins:    []int{1, 2, 5},
			amount:   0,
			expected: 1, // One way: use no coins
		},
		{
			name:     "Standard case",
			coins:    []int{1, 2, 5},
			amount:   5,
			expected: 4, // 5, 2+2+1, 2+1+1+1, 1+1+1+1+1
		},
		{
			name:     "Single coin type",
			coins:    []int{2},
			amount:   3,
			expected: 0, // Impossible
		},
		{
			name:     "Multiple ways",
			coins:    []int{2, 3, 5},
			amount:   9,
			expected: 3, // 2+2+5, 3+3+3, 2+2+2+3
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CoinChangeCountWays(tc.coins, tc.amount)
			if result != tc.expected {
				t.Errorf("CoinChangeCountWays(%v, %d) = %d, expected %d",
					tc.coins, tc.amount, result, tc.expected)
			}
		})
	}
}

// ============================================================================
// MATRIX CHAIN MULTIPLICATION TESTS
// ============================================================================

func TestMatrixChainMultiplication(t *testing.T) {
	testCases := []struct {
		name       string
		dimensions []int
		expected   int
	}{
		{
			name:       "Single matrix",
			dimensions: []int{10, 20},
			expected:   0,
		},
		{
			name:       "Two matrices",
			dimensions: []int{10, 20, 30},
			expected:   6000, // 10*20*30
		},
		{
			name:       "Three matrices",
			dimensions: []int{1, 2, 3, 4},
			expected:   18, // ((1*2)*3)*4 = 6 + 12 = 18 vs (1*2)*(3*4) = 24
		},
		{
			name:       "Classic example",
			dimensions: []int{40, 20, 30, 10, 30},
			expected:   26000,
		},
		{
			name:       "Four matrices",
			dimensions: []int{5, 10, 3, 12, 5},
			expected:   405, // Actual minimum is 405, not 2010
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MatrixChainMultiplication(tc.dimensions)
			if result != tc.expected {
				t.Errorf("MatrixChainMultiplication(%v) = %d, expected %d",
					tc.dimensions, result, tc.expected)
			}
		})
	}
}

// ============================================================================
// PALINDROME TESTS
// ============================================================================

func TestLongestPalindromicSubsequence(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "Single character",
			s:        "a",
			expected: 1,
		},
		{
			name:     "All same characters",
			s:        "aaaa",
			expected: 4,
		},
		{
			name:     "No repeating characters",
			s:        "abcd",
			expected: 1,
		},
		{
			name:     "Classic example",
			s:        "bbbab",
			expected: 4, // "bbbb"
		},
		{
			name:     "Complex case",
			s:        "cbbd",
			expected: 2, // "bb"
		},
		{
			name:     "Longer example",
			s:        "BBABCBCAB",
			expected: 7, // "BABCBAB"
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := LongestPalindromicSubsequence(tc.s)
			if result != tc.expected {
				t.Errorf("LongestPalindromicSubsequence('%s') = %d, expected %d",
					tc.s, result, tc.expected)
			}
		})
	}
}

func TestMinPalindromicPartition(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "Single character",
			s:        "a",
			expected: 0,
		},
		{
			name:     "Already palindrome",
			s:        "aba",
			expected: 0,
		},
		{
			name:     "All different",
			s:        "abc",
			expected: 2, // "a|b|c"
		},
		{
			name:     "Classic example",
			s:        "aab",
			expected: 1, // "aa|b"
		},
		{
			name:     "Complex case",
			s:        "ababbbabbababa",
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MinPalindromicPartition(tc.s)
			if result != tc.expected {
				t.Errorf("MinPalindromicPartition('%s') = %d, expected %d",
					tc.s, result, tc.expected)
			}
		})
	}
}

// ============================================================================
// PERFORMANCE BENCHMARKS
// ============================================================================

func BenchmarkFibonacci(b *testing.B) {
	algorithms := map[string]func(int) int{
		"Memoization": FibonacciMemoization,
		"Tabulation":  FibonacciTabulation,
		"Optimized":   FibonacciOptimized,
	}

	testSizes := []int{10, 20, 30, 40}

	for name, algorithm := range algorithms {
		for _, size := range testSizes {
			b.Run(fmt.Sprintf("%s_n%d", name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					algorithm(size)
				}
			})
		}
	}
}

func BenchmarkKnapsack(b *testing.B) {
	// Create test items
	items := make([]Item, 20)
	for i := range items {
		items[i] = Item{
			Weight: i + 1,
			Value:  (i + 1) * 2,
			Name:   fmt.Sprintf("Item%d", i),
		}
	}

	algorithms := map[string]func([]Item, int) int{
		"Memoization": Knapsack01Memoization,
		"Tabulation":  Knapsack01Tabulation,
		"Optimized":   Knapsack01Optimized,
	}

	capacities := []int{50, 100, 200}

	for name, algorithm := range algorithms {
		for _, capacity := range capacities {
			b.Run(fmt.Sprintf("%s_cap%d", name, capacity), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					algorithm(items, capacity)
				}
			})
		}
	}
}

func BenchmarkLCS(b *testing.B) {
	algorithms := map[string]func(string, string) int{
		"Memoization": LCSMemoization,
		"Tabulation":  LCSTabulation,
		"Optimized":   LCSOptimized,
	}

	testCases := []struct {
		name  string
		text1 string
		text2 string
	}{
		{"Short", "ABCDGH", "AEDFHR"},
		{"Medium", "programming", "algorithm"},
		{"Long", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ACEGIKMOQSUWY"},
	}

	for name, algorithm := range algorithms {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("%s_%s", name, tc.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					algorithm(tc.text1, tc.text2)
				}
			})
		}
	}
}

func BenchmarkLIS(b *testing.B) {
	algorithms := map[string]func([]int) int{
		"Tabulation":   LISTabulation,
		"BinarySearch": LISBinarySearch,
	}

	// Generate test arrays of different sizes
	sizes := []int{100, 500, 1000}

	for name, algorithm := range algorithms {
		for _, size := range sizes {
			nums := make([]int, size)
			for i := range nums {
				nums[i] = i % 100 // Create some pattern
			}

			b.Run(fmt.Sprintf("%s_size%d", name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					algorithm(nums)
				}
			})
		}
	}
}

// ============================================================================
// STRESS TESTS
// ============================================================================

func TestLargeInputs(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress tests in short mode")
	}

	t.Run("Fibonacci_Large", func(t *testing.T) {
		n := 1000
		// Only test optimized version for large inputs
		result1 := FibonacciOptimized(n)
		result2 := FibonacciTabulation(n)

		if result1 != result2 {
			t.Errorf("Fibonacci results don't match for n=%d", n)
		}
	})

	t.Run("LCS_Large", func(t *testing.T) {
		// Generate large strings
		text1 := ""
		text2 := ""
		for i := 0; i < 500; i++ {
			text1 += string(rune('a' + (i % 26)))
			text2 += string(rune('a' + ((i * 2) % 26)))
		}

		result1 := LCSOptimized(text1, text2)
		result2 := LCSTabulation(text1, text2)

		if result1 != result2 {
			t.Errorf("LCS results don't match for large inputs")
		}
	})
}

// ============================================================================
// CORRECTNESS PROPERTY TESTS
// ============================================================================

func TestDPProperties(t *testing.T) {
	t.Run("Fibonacci_Consistency", func(t *testing.T) {
		for n := 0; n <= 20; n++ {
			memo := FibonacciMemoization(n)
			tab := FibonacciTabulation(n)
			opt := FibonacciOptimized(n)

			if memo != tab || tab != opt {
				t.Errorf("Fibonacci algorithms inconsistent at n=%d: memo=%d, tab=%d, opt=%d",
					n, memo, tab, opt)
			}
		}
	})

	t.Run("Knapsack_Consistency", func(t *testing.T) {
		items := []Item{
			{Weight: 2, Value: 3},
			{Weight: 3, Value: 4},
			{Weight: 4, Value: 5},
			{Weight: 5, Value: 6},
		}

		for capacity := 0; capacity <= 10; capacity++ {
			memo := Knapsack01Memoization(items, capacity)
			tab := Knapsack01Tabulation(items, capacity)
			opt := Knapsack01Optimized(items, capacity)

			if memo != tab || tab != opt {
				t.Errorf("Knapsack algorithms inconsistent at capacity=%d: memo=%d, tab=%d, opt=%d",
					capacity, memo, tab, opt)
			}
		}
	})

	t.Run("LCS_Consistency", func(t *testing.T) {
		testStrings := [][]string{
			{"", ""},
			{"a", "b"},
			{"abc", "def"},
			{"abcde", "ace"},
			{"ABCDGH", "AEDFHR"},
		}

		for _, pair := range testStrings {
			text1, text2 := pair[0], pair[1]
			memo := LCSMemoization(text1, text2)
			tab := LCSTabulation(text1, text2)
			opt := LCSOptimized(text1, text2)

			if memo != tab || tab != opt {
				t.Errorf("LCS algorithms inconsistent for '%s' vs '%s': memo=%d, tab=%d, opt=%d",
					text1, text2, memo, tab, opt)
			}
		}
	})
}
