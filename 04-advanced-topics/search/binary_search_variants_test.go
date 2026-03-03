package search

import (
	"reflect"
	"testing"
)

func TestClassicBinarySearch(t *testing.T) {
	abs := NewAdvancedBinarySearch()

	tests := []struct {
		name     string
		arr      []int
		target   int
		found    bool
		index    int
	}{
		{"empty array", []int{}, 1, false, -1},
		{"single element found", []int{5}, 5, true, 0},
		{"single element not found", []int{5}, 3, false, -1},
		{"found in middle", []int{1, 2, 3, 4, 5}, 3, true, 2},
		{"found at start", []int{1, 2, 3, 4, 5}, 1, true, 0},
		{"found at end", []int{1, 2, 3, 4, 5}, 5, true, 4},
		{"not present", []int{1, 2, 4, 5}, 3, false, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := abs.ClassicBinarySearch(tt.arr, tt.target)
			if result.Found != tt.found || result.Index != tt.index {
				t.Errorf("ClassicBinarySearch(%v, %d) = %+v, want Found=%t Index=%d",
					tt.arr, tt.target, result, tt.found, tt.index)
			}
		})
	}
}

func TestLowerBound(t *testing.T) {
	abs := NewAdvancedBinarySearch()

	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"first >= target", []int{1, 2, 2, 3, 4}, 2, 1},
		{"insert position", []int{1, 3, 5}, 4, 2},
		{"all smaller", []int{1, 2, 3}, 5, 3},
		{"all larger", []int{2, 4, 6}, 1, 0},
		{"empty", []int{}, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := abs.LowerBound(tt.arr, tt.target)
			if got != tt.expected {
				t.Errorf("LowerBound(%v, %d) = %d, want %d", tt.arr, tt.target, got, tt.expected)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	abs := NewAdvancedBinarySearch()

	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"first > target", []int{1, 2, 2, 3, 4}, 2, 3},
		{"insert position", []int{1, 3, 5}, 4, 2},
		{"all smaller", []int{1, 2, 3}, 5, 3},
		{"all larger", []int{2, 4, 6}, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := abs.UpperBound(tt.arr, tt.target)
			if got != tt.expected {
				t.Errorf("UpperBound(%v, %d) = %d, want %d", tt.arr, tt.target, got, tt.expected)
			}
		})
	}
}

func TestSearchRotatedArray(t *testing.T) {
	abs := NewAdvancedBinarySearch()

	tests := []struct {
		name     string
		arr      []int
		target   int
		found    bool
		index    int
	}{
		{"found in rotated", []int{4, 5, 6, 7, 0, 1, 2}, 0, true, 4},
		{"found pivot", []int{4, 5, 6, 7, 0, 1, 2}, 7, true, 3},
		{"not in rotated", []int{4, 5, 6, 7, 0, 1, 2}, 3, false, -1},
		{"sorted array", []int{1, 2, 3, 4, 5}, 3, true, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := abs.SearchRotatedArray(tt.arr, tt.target)
			if result.Found != tt.found || result.Index != tt.index {
				t.Errorf("SearchRotatedArray(%v, %d) = %+v, want Found=%t Index=%d",
					tt.arr, tt.target, result, tt.found, tt.index)
			}
		})
	}
}

func TestFindPivotInRotatedArray(t *testing.T) {
	abs := NewAdvancedBinarySearch()

	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{"rotated by 1", []int{3, 1, 2}, 1},
		{"rotated", []int{4, 5, 6, 7, 0, 1, 2}, 4},
		{"not rotated", []int{1, 2, 3}, 0},
		{"two elements", []int{2, 1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := abs.FindPivotInRotatedArray(tt.arr)
			if got != tt.expected {
				t.Errorf("FindPivotInRotatedArray(%v) = %d, want %d", tt.arr, got, tt.expected)
			}
		})
	}
}

func TestSearchRange(t *testing.T) {
	abs := NewAdvancedBinarySearch()

	tests := []struct {
		name     string
		arr      []int
		target   int
		expected [2]int
	}{
		{"range exists", []int{5, 7, 7, 8, 8, 10}, 8, [2]int{3, 4}},
		{"single occurrence", []int{5, 7, 7, 8, 8, 10}, 5, [2]int{0, 0}},
		{"not found", []int{5, 7, 7, 8, 8, 10}, 6, [2]int{-1, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := abs.SearchRange(tt.arr, tt.target)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SearchRange(%v, %d) = %v, want %v", tt.arr, tt.target, got, tt.expected)
			}
		})
	}
}

func TestExponentialSearchUnbounded(t *testing.T) {
	es := NewExponentialSearch()

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := es.SearchUnbounded(arr, 7)
	if !result.Found || result.Index != 6 {
		t.Errorf("SearchUnbounded(...) = %+v, want Found=true Index=6", result)
	}

	result = es.SearchUnbounded(arr, 11)
	if result.Found {
		t.Errorf("SearchUnbounded(..., 11) = Found=true, want false")
	}
}

func TestTernarySearchDiscrete(t *testing.T) {
	ts := NewTernarySearch()

	arr := []int{1, 3, 5, 4, 2}
	maxIdx := ts.TernarySearchDiscrete(arr, true)
	if maxIdx != 2 || arr[maxIdx] != 5 {
		t.Errorf("TernarySearchDiscrete(max) = index %d value %d, want index 2 value 5", maxIdx, arr[maxIdx])
	}

	minIdx := ts.TernarySearchDiscrete(arr, false)
	if minIdx != 0 || arr[minIdx] != 1 {
		t.Errorf("TernarySearchDiscrete(min) = index %d value %d, want index 0 value 1", minIdx, arr[minIdx])
	}
}
