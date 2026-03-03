package number_theory

import (
	"reflect"
	"testing"
)

func TestSieveOfEratosthenes(t *testing.T) {
	pg := NewPrimeGenerator()

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{"n less than 2", 1, []int{}},
		{"n is 2", 2, []int{2}},
		{"n is 10", 10, []int{2, 3, 5, 7}},
		{"n is 20", 20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{"n is 30", 30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pg.SieveOfEratosthenes(tt.n)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SieveOfEratosthenes(%d) = %v, want %v", tt.n, got, tt.expected)
			}
		})
	}
}

func TestIsPrimeTrial(t *testing.T) {
	pg := NewPrimeGenerator()

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	for _, p := range primes {
		if !pg.IsPrimeTrial(p) {
			t.Errorf("IsPrimeTrial(%d) = false, want true", p)
		}
	}

	composites := []int{0, 1, 4, 6, 8, 9, 10, 12, 15, 20}
	for _, c := range composites {
		if pg.IsPrimeTrial(c) {
			t.Errorf("IsPrimeTrial(%d) = true, want false", c)
		}
	}
}

func TestSegmentedSieve(t *testing.T) {
	pg := NewPrimeGenerator()

	got := pg.SegmentedSieve(10, 30)
	expected := []int{11, 13, 17, 19, 23, 29}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("SegmentedSieve(10, 30) = %v, want %v", got, expected)
	}
}

func TestGCD(t *testing.T) {
	gcd := NewGCDAlgorithms()

	tests := []struct {
		a, b     int
		expected int
	}{
		{12, 8, 4},
		{17, 19, 1},
		{100, 35, 5},
		{0, 5, 5},
		{7, 0, 7},
	}

	for _, tt := range tests {
		got := gcd.GCD(tt.a, tt.b)
		if got != tt.expected {
			t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}

func TestLCM(t *testing.T) {
	gcd := NewGCDAlgorithms()

	if got := gcd.LCM(12, 8); got != 24 {
		t.Errorf("LCM(12, 8) = %d, want 24", got)
	}
	if got := gcd.LCM(5, 7); got != 35 {
		t.Errorf("LCM(5, 7) = %d, want 35", got)
	}
}

func TestModPow(t *testing.T) {
	ma := NewModularArithmetic()

	if got := ma.ModPow(2, 10, 1000); got != 24 {
		t.Errorf("ModPow(2, 10, 1000) = %d, want 24", got)
	}
	if got := ma.ModPow(3, 4, 17); got != 13 {
		t.Errorf("ModPow(3, 4, 17) = %d, want 13", got)
	}
}

func TestPrimeFactorization(t *testing.T) {
	fa := NewFactorizationAlgorithms()

	got := fa.PrimeFactorization(84)
	expected := map[int]int{2: 2, 3: 1, 7: 1}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("PrimeFactorization(84) = %v, want %v", got, expected)
	}
}

func TestEulerTotient(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 1},
		{9, 6},
		{10, 4},
	}

	for _, tt := range tests {
		got := EulerTotient(tt.n)
		if got != tt.expected {
			t.Errorf("EulerTotient(%d) = %d, want %d", tt.n, got, tt.expected)
		}
	}
}
