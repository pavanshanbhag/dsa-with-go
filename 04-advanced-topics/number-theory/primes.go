package number_theory

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// ===============================
// Prime Number Algorithms
// ===============================

// PrimeGenerator provides various prime number generation and testing algorithms
type PrimeGenerator struct{}

// NewPrimeGenerator creates a new prime generator instance
func NewPrimeGenerator() *PrimeGenerator {
	return &PrimeGenerator{}
}

// SieveOfEratosthenes generates all primes up to n
func (pg *PrimeGenerator) SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}

	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

// SegmentedSieve generates primes in range [low, high] efficiently
func (pg *PrimeGenerator) SegmentedSieve(low, high int) []int {
	limit := int(math.Sqrt(float64(high))) + 1
	basePrimes := pg.SieveOfEratosthenes(limit)

	size := high - low + 1
	isPrime := make([]bool, size)
	for i := range isPrime {
		isPrime[i] = true
	}

	for _, prime := range basePrimes {
		// Find the minimum number in [low, high] that is a multiple of prime
		start := ((low + prime - 1) / prime) * prime
		if start == prime {
			start += prime
		}

		for j := start; j <= high; j += prime {
			isPrime[j-low] = false
		}
	}

	var primes []int
	for i := 0; i < size; i++ {
		if isPrime[i] && (low+i) > 1 {
			primes = append(primes, low+i)
		}
	}

	return primes
}

// IsPrimeTrial tests primality using trial division
func (pg *PrimeGenerator) IsPrimeTrial(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n))) + 1
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// MillerRabinTest performs probabilistic primality test
func (pg *PrimeGenerator) MillerRabinTest(n int, k int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Write n-1 as d * 2^r
	d := n - 1
	r := 0
	for d%2 == 0 {
		d /= 2
		r++
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < k; i++ {
		a := rand.Intn(n-3) + 2 // Random number in [2, n-2]
		x := pg.modPow(a, d, n)

		if x == 1 || x == n-1 {
			continue
		}

		composite := true
		for j := 0; j < r-1; j++ {
			x = (x * x) % n
			if x == n-1 {
				composite = false
				break
			}
		}

		if composite {
			return false
		}
	}

	return true
}

// ===============================
// Modular Arithmetic
// ===============================

// ModularArithmetic provides modular arithmetic operations
type ModularArithmetic struct{}

// NewModularArithmetic creates new modular arithmetic instance
func NewModularArithmetic() *ModularArithmetic {
	return &ModularArithmetic{}
}

// ModPow computes (base^exp) % mod efficiently
func (ma *ModularArithmetic) ModPow(base, exp, mod int) int {
	if mod == 1 {
		return 0
	}

	result := 1
	base = base % mod

	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp = exp >> 1
		base = (base * base) % mod
	}

	return result
}

// modPow helper function for MillerRabinTest
func (pg *PrimeGenerator) modPow(base, exp, mod int) int {
	ma := NewModularArithmetic()
	return ma.ModPow(base, exp, mod)
}

// ModInverse computes modular multiplicative inverse using extended Euclidean algorithm
func (ma *ModularArithmetic) ModInverse(a, mod int) (int, bool) {
	g, x, _ := ma.ExtendedGCD(a, mod)
	if g != 1 {
		return 0, false // Inverse doesn't exist
	}

	return (x%mod + mod) % mod, true
}

// ExtendedGCD computes gcd(a, b) and coefficients x, y such that ax + by = gcd(a, b)
func (ma *ModularArithmetic) ExtendedGCD(a, b int) (gcd, x, y int) {
	if a == 0 {
		return b, 0, 1
	}

	gcd1, x1, y1 := ma.ExtendedGCD(b%a, a)
	x = y1 - (b/a)*x1
	y = x1

	return gcd1, x, y
}

// ChineseRemainderTheorem solves system of congruences
func (ma *ModularArithmetic) ChineseRemainderTheorem(remainders, moduli []int) (int, bool) {
	if len(remainders) != len(moduli) {
		return 0, false
	}

	// Check if moduli are pairwise coprime
	for i := 0; i < len(moduli); i++ {
		for j := i + 1; j < len(moduli); j++ {
			if ma.GCD(moduli[i], moduli[j]) != 1 {
				return 0, false
			}
		}
	}

	N := 1
	for _, mod := range moduli {
		N *= mod
	}

	result := 0
	for i := 0; i < len(remainders); i++ {
		Ni := N / moduli[i]
		Mi, exists := ma.ModInverse(Ni, moduli[i])
		if !exists {
			return 0, false
		}

		result = (result + remainders[i]*Ni*Mi) % N
	}

	return (result + N) % N, true
}

// ===============================
// GCD and Related Algorithms
// ===============================

// GCDAlgorithms provides various GCD computation methods
type GCDAlgorithms struct{}

// NewGCDAlgorithms creates new GCD algorithms instance
func NewGCDAlgorithms() *GCDAlgorithms {
	return &GCDAlgorithms{}
}

// GCD computes greatest common divisor using Euclidean algorithm
func (gcd *GCDAlgorithms) GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// GCD method for ModularArithmetic
func (ma *ModularArithmetic) GCD(a, b int) int {
	gcdAlg := NewGCDAlgorithms()
	return gcdAlg.GCD(a, b)
}

// LCM computes least common multiple
func (gcd *GCDAlgorithms) LCM(a, b int) int {
	return (a * b) / gcd.GCD(a, b)
}

// BinaryGCD computes GCD using binary algorithm (Stein's algorithm)
func (gcd *GCDAlgorithms) BinaryGCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Count the number of times 2 divides both a and b
	shift := 0
	for ((a | b) & 1) == 0 {
		a >>= 1
		b >>= 1
		shift++
	}

	// Remove all factors of 2 from a
	for (a & 1) == 0 {
		a >>= 1
	}

	for b != 0 {
		// Remove all factors of 2 from b
		for (b & 1) == 0 {
			b >>= 1
		}

		// Ensure a <= b
		if a > b {
			a, b = b, a
		}

		b = b - a
	}

	return a << shift
}

// ===============================
// Factorization Algorithms
// ===============================

// FactorizationAlgorithms provides number factorization methods
type FactorizationAlgorithms struct {
	primeGen *PrimeGenerator
}

// NewFactorizationAlgorithms creates new factorization instance
func NewFactorizationAlgorithms() *FactorizationAlgorithms {
	return &FactorizationAlgorithms{
		primeGen: NewPrimeGenerator(),
	}
}

// PrimeFactorization returns prime factorization of n
func (fa *FactorizationAlgorithms) PrimeFactorization(n int) map[int]int {
	factors := make(map[int]int)

	// Handle factor 2
	for n%2 == 0 {
		factors[2]++
		n /= 2
	}

	// Check odd factors
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}

	// If n is still > 1, then it's a prime
	if n > 1 {
		factors[n]++
	}

	return factors
}

// PollardRho attempts to find a non-trivial factor using Pollard's rho algorithm
func (fa *FactorizationAlgorithms) PollardRho(n int) int {
	if n%2 == 0 {
		return 2
	}

	x := 2
	y := 2
	d := 1

	f := func(x int) int {
		return (x*x + 1) % n
	}

	gcdAlg := NewGCDAlgorithms()

	for d == 1 {
		x = f(x)
		y = f(f(y))
		d = gcdAlg.GCD(abs(x-y), n)
	}

	if d == n {
		return -1 // Failed to find factor
	}

	return d
}

// TrialDivision performs trial division factorization
func (fa *FactorizationAlgorithms) TrialDivision(n int) []int {
	var factors []int

	// Handle factor 2
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// Check odd factors
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// If n is still > 1, then it's a prime
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

// ===============================
// Utility Functions
// ===============================

// abs returns absolute value of integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// EulerTotient computes Euler's totient function φ(n)
func EulerTotient(n int) int {
	result := n

	// Consider all prime factors and subtract their multiples
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			// Remove all factors i from n
			for n%i == 0 {
				n /= i
			}
			// Subtract multiples of i from result
			result -= result / i
		}
	}

	// If n has a prime factor greater than sqrt(n)
	if n > 1 {
		result -= result / n
	}

	return result
}

// FastExponentiation computes base^exp efficiently (without modulo)
func FastExponentiation(base, exp int) *big.Int {
	result := big.NewInt(1)
	baseBig := big.NewInt(int64(base))

	for exp > 0 {
		if exp%2 == 1 {
			result.Mul(result, baseBig)
		}
		baseBig.Mul(baseBig, baseBig)
		exp /= 2
	}

	return result
}

// IsPerfectPower checks if n is a perfect power (n = a^b for some a, b > 1)
func IsPerfectPower(n int) (bool, int, int) {
	for exp := 2; exp <= int(math.Log2(float64(n))); exp++ {
		base := int(math.Pow(float64(n), 1.0/float64(exp)))

		// Check base and base+1 due to floating point precision
		for _, b := range []int{base - 1, base, base + 1} {
			if b > 1 {
				temp := 1
				for i := 0; i < exp; i++ {
					temp *= b
					if temp > n {
						break
					}
				}
				if temp == n {
					return true, b, exp
				}
			}
		}
	}

	return false, 0, 0
}

// ===============================
// Number Theory Statistics
// ===============================

// NumberTheoryStats provides analysis of number theoretic properties
type NumberTheoryStats struct {
	Number         int
	IsPrime        bool
	PrimeFactors   map[int]int
	Divisors       []int
	EulerTotient   int
	IsPerfectPower bool
	PowerBase      int
	PowerExponent  int
}

// AnalyzeNumber provides comprehensive number theory analysis
func AnalyzeNumber(n int) NumberTheoryStats {
	pg := NewPrimeGenerator()
	fa := NewFactorizationAlgorithms()

	stats := NumberTheoryStats{
		Number: n,
	}

	// Check if prime
	stats.IsPrime = pg.IsPrimeTrial(n)

	// Get prime factorization
	stats.PrimeFactors = fa.PrimeFactorization(n)

	// Calculate divisors
	stats.Divisors = calculateDivisors(n)

	// Calculate Euler's totient
	stats.EulerTotient = EulerTotient(n)

	// Check if perfect power
	stats.IsPerfectPower, stats.PowerBase, stats.PowerExponent = IsPerfectPower(n)

	return stats
}

// calculateDivisors finds all divisors of n
func calculateDivisors(n int) []int {
	var divisors []int

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}

	return divisors
}

// PrintNumberAnalysis displays comprehensive number theory analysis
func PrintNumberAnalysis(stats NumberTheoryStats) {
	fmt.Printf("Number Theory Analysis for: %d\n", stats.Number)
	fmt.Println("=====================================")
	fmt.Printf("Is Prime: %t\n", stats.IsPrime)
	fmt.Printf("Prime Factorization: %v\n", stats.PrimeFactors)
	fmt.Printf("Number of Divisors: %d\n", len(stats.Divisors))
	fmt.Printf("Divisors: %v\n", stats.Divisors)
	fmt.Printf("Euler's Totient φ(%d): %d\n", stats.Number, stats.EulerTotient)
	fmt.Printf("Is Perfect Power: %t", stats.IsPerfectPower)
	if stats.IsPerfectPower {
		fmt.Printf(" (%d^%d)", stats.PowerBase, stats.PowerExponent)
	}
	fmt.Println()
	fmt.Println()
}
