package main

import (
	"fmt"
	"math"
	"time"

	number_theory "dsa-mastery/04-advanced-topics/number-theory"
	"dsa-mastery/04-advanced-topics/search"
)

func main() {
	fmt.Println("🔬 DSA Mastery - Advanced Topics Demonstration")
	fmt.Println("==============================================")

	// 1. Advanced Search Algorithms
	fmt.Println("\n🔍 1. Advanced Search Algorithms")
	fmt.Println("--------------------------------")
	demonstrateAdvancedSearch()

	// 2. Number Theory Algorithms
	fmt.Println("\n🔢 2. Number Theory Algorithms")
	fmt.Println("-------------------------------")
	demonstrateNumberTheory()

	// 3. Mathematical Optimization
	fmt.Println("\n📈 3. Mathematical Optimization")
	fmt.Println("-------------------------------")
	demonstrateMathematicalOptimization()

	// 4. Performance Analysis
	fmt.Println("\n⚡ 4. Performance Analysis")
	fmt.Println("-------------------------")
	demonstratePerformanceAnalysis()

	// 5. Real-World Applications
	fmt.Println("\n🌟 5. Real-World Applications")
	fmt.Println("-----------------------------")
	demonstrateRealWorldApplications()

	fmt.Println("\n✅ All advanced topics demonstrated successfully!")
	fmt.Println("🚀 Ready for expert-level algorithmic challenges!")
}

func demonstrateAdvancedSearch() {
	fmt.Println("Problem: Demonstrate various advanced search techniques")

	// Binary Search Variants
	fmt.Println("\n📊 Binary Search Variants:")
	abs := search.NewAdvancedBinarySearch()

	// Sorted array for testing
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25}
	target := 13

	start := time.Now()
	result := abs.ClassicBinarySearch(arr, target)
	duration := time.Since(start)
	fmt.Printf("Classic Binary Search for %d: Found=%t, Index=%d (Time: %v)\n",
		target, result.Found, result.Index, duration)

	// Lower and Upper Bound
	target = 15
	lower := abs.LowerBound(arr, target)
	upper := abs.UpperBound(arr, target)
	fmt.Printf("Lower bound of %d: %d, Upper bound: %d\n", target, lower, upper)

	// Search Range
	searchRange := abs.SearchRange(arr, target)
	fmt.Printf("Range of %d: [%d, %d]\n", target, searchRange[0], searchRange[1])

	// Rotated Array Search
	fmt.Println("\n🔄 Rotated Array Search:")
	rotatedArr := []int{15, 17, 19, 21, 23, 25, 1, 3, 5, 7, 9, 11, 13}
	target = 5
	start = time.Now()
	result = abs.SearchRotatedArray(rotatedArr, target)
	duration = time.Since(start)
	fmt.Printf("Rotated array search for %d: Found=%t, Index=%d (Time: %v)\n",
		target, result.Found, result.Index, duration)

	// Find pivot
	pivot := abs.FindPivotInRotatedArray(rotatedArr)
	fmt.Printf("Pivot (minimum element) in rotated array: Index=%d, Value=%d\n",
		pivot, rotatedArr[pivot])

	// Ternary Search
	fmt.Println("\n🔺 Ternary Search for Optimization:")
	ts := search.NewTernarySearch()

	// Find maximum of -x^2 + 4x + 1 in [0, 5]
	quadratic := func(x float64) float64 {
		return -x*x + 4*x + 1
	}

	start = time.Now()
	maxPoint := ts.FindMaximum(quadratic, 0, 5, 1e-6)
	duration = time.Since(start)
	fmt.Printf("Maximum of f(x) = -x² + 4x + 1: x = %.6f, f(x) = %.6f (Time: %v)\n",
		maxPoint, quadratic(maxPoint), duration)

	// Exponential Search
	fmt.Println("\n⚡ Exponential Search:")
	es := search.NewExponentialSearch()
	largeArr := make([]int, 10000)
	for i := range largeArr {
		largeArr[i] = i * 2 // Even numbers
	}

	target = 1234
	start = time.Now()
	result = es.SearchUnbounded(largeArr, target)
	duration = time.Since(start)
	fmt.Printf("Exponential search for %d in 10k elements: Found=%t, Index=%d (Time: %v)\n",
		target, result.Found, result.Index, duration)

	// Interpolation Search
	fmt.Println("\n🎯 Interpolation Search:")
	is := search.NewInterpolationSearch()
	uniformArr := make([]int, 1000)
	for i := range uniformArr {
		uniformArr[i] = i * 10 // Uniform distribution
	}

	target = 5670
	start = time.Now()
	result = is.Search(uniformArr, target)
	duration = time.Since(start)
	fmt.Printf("Interpolation search for %d in uniform data: Found=%t, Index=%d (Time: %v)\n",
		target, result.Found, result.Index, duration)
}

func demonstrateNumberTheory() {
	fmt.Println("Problem: Demonstrate number theory algorithms for cryptographic applications")

	// Prime Generation
	fmt.Println("\n🔢 Prime Number Generation:")
	pg := number_theory.NewPrimeGenerator()

	start := time.Now()
	primes := pg.SieveOfEratosthenes(100)
	duration := time.Since(start)
	fmt.Printf("Primes up to 100 (%d primes): %v (Time: %v)\n", len(primes), primes, duration)

	// Large prime testing
	largePrimes := []int{982451653, 982451679, 982451681, 982451707}
	fmt.Println("\n🧪 Primality Testing:")
	for _, n := range largePrimes {
		start = time.Now()
		isPrimeTrial := pg.IsPrimeTrial(n)
		trialTime := time.Since(start)

		start = time.Now()
		isPrimeMillerRabin := pg.MillerRabinTest(n, 10)
		millerTime := time.Since(start)

		fmt.Printf("n=%d: Trial Division=%t (%v), Miller-Rabin=%t (%v)\n",
			n, isPrimeTrial, trialTime, isPrimeMillerRabin, millerTime)
	}

	// Modular Arithmetic
	fmt.Println("\n🔐 Modular Arithmetic:")
	ma := number_theory.NewModularArithmetic()

	base, exp, mod := 2, 1000, 1000000007
	start = time.Now()
	result := ma.ModPow(base, exp, mod)
	duration = time.Since(start)
	fmt.Printf("ModPow(%d, %d, %d) = %d (Time: %v)\n", base, exp, mod, result, duration)

	// Modular inverse
	a, m := 3, 11
	inv, exists := ma.ModInverse(a, m)
	if exists {
		fmt.Printf("Modular inverse of %d mod %d = %d (Verification: %d * %d ≡ %d mod %d)\n",
			a, m, inv, a, inv, (a*inv)%m, m)
	}

	// Chinese Remainder Theorem
	remainders := []int{2, 3, 2}
	moduli := []int{3, 5, 7}
	solution, valid := ma.ChineseRemainderTheorem(remainders, moduli)
	if valid {
		fmt.Printf("Chinese Remainder Theorem solution: x ≡ %d (mod %d)\n",
			solution, moduli[0]*moduli[1]*moduli[2])
		// Verify
		fmt.Printf("Verification: %d ≡ %d (mod %d), %d ≡ %d (mod %d), %d ≡ %d (mod %d)\n",
			solution, solution%moduli[0], moduli[0],
			solution, solution%moduli[1], moduli[1],
			solution, solution%moduli[2], moduli[2])
	}

	// GCD and LCM
	fmt.Println("\n🔗 GCD and LCM:")
	gcdAlg := number_theory.NewGCDAlgorithms()

	a, b := 48, 18
	gcd := gcdAlg.GCD(a, b)
	lcm := gcdAlg.LCM(a, b)
	fmt.Printf("GCD(%d, %d) = %d, LCM(%d, %d) = %d\n", a, b, gcd, a, b, lcm)

	// Binary GCD
	start = time.Now()
	binaryGCD := gcdAlg.BinaryGCD(a, b)
	duration = time.Since(start)
	fmt.Printf("Binary GCD(%d, %d) = %d (Time: %v)\n", a, b, binaryGCD, duration)

	// Number Factorization
	fmt.Println("\n🔧 Number Factorization:")
	fa := number_theory.NewFactorizationAlgorithms()

	numbers := []int{60, 315, 1001, 9999}
	for _, n := range numbers {
		start = time.Now()
		factors := fa.PrimeFactorization(n)
		duration = time.Since(start)
		fmt.Printf("Prime factorization of %d: %v (Time: %v)\n", n, factors, duration)

		// Trial division
		trialFactors := fa.TrialDivision(n)
		fmt.Printf("Trial division factors: %v\n", trialFactors)
	}

	// Number Analysis
	fmt.Println("\n📊 Comprehensive Number Analysis:")
	testNumbers := []int{100, 127, 256, 1001}
	for _, n := range testNumbers {
		fmt.Printf("\n--- Analysis of %d ---\n", n)
		stats := number_theory.AnalyzeNumber(n)
		number_theory.PrintNumberAnalysis(stats)
	}
}

func demonstrateMathematicalOptimization() {
	fmt.Println("Problem: Demonstrate mathematical optimization techniques")

	// Ternary search for continuous functions
	fmt.Println("\n📈 Function Optimization:")
	ts := search.NewTernarySearch()

	// Minimize f(x) = x^4 - 14x^3 + 60x^2 - 70x + 15
	polynomial := func(x float64) float64 {
		return x*x*x*x - 14*x*x*x + 60*x*x - 70*x + 15
	}

	start := time.Now()
	minPoint := ts.FindMinimum(polynomial, 0, 8, 1e-8)
	duration := time.Since(start)
	fmt.Printf("Minimum of polynomial: x = %.8f, f(x) = %.8f (Time: %v)\n",
		minPoint, polynomial(minPoint), duration)

	// Golden section search simulation
	fmt.Println("\n🏆 Golden Section Search:")
	goldenRatio := (1 + math.Sqrt(5)) / 2
	fmt.Printf("Golden ratio φ = %.10f\n", goldenRatio)

	// Simulate golden section search for optimization
	costFunction := func(x float64) float64 {
		return (x-3)*(x-3) + 2 // Minimum at x=3
	}

	left, right := 0.0, 6.0
	tolerance := 1e-6
	iterations := 0

	for right-left > tolerance {
		c := right - (right-left)/goldenRatio
		d := left + (right-left)/goldenRatio

		if costFunction(c) > costFunction(d) {
			left = c
		} else {
			right = d
		}
		iterations++
	}

	optimum := (left + right) / 2
	fmt.Printf("Golden section minimum: x = %.8f, f(x) = %.8f (%d iterations)\n",
		optimum, costFunction(optimum), iterations)

	// Numerical root finding simulation
	fmt.Println("\n🎯 Root Finding:")

	// Find root of f(x) = x^3 - 2x - 5 using bisection method
	equation := func(x float64) float64 {
		return x*x*x - 2*x - 5
	}

	left, right = 2.0, 3.0
	tolerance = 1e-10
	iterations = 0

	start = time.Now()
	for right-left > tolerance {
		mid := (left + right) / 2
		if equation(mid)*equation(left) < 0 {
			right = mid
		} else {
			left = mid
		}
		iterations++
	}
	duration = time.Since(start)

	root := (left + right) / 2
	fmt.Printf("Bisection method root: x = %.10f, f(x) = %.2e (%d iterations, Time: %v)\n",
		root, equation(root), iterations, duration)
}

func demonstratePerformanceAnalysis() {
	fmt.Println("Performance comparison across different algorithm categories")

	// Search Algorithm Performance
	fmt.Println("\n🔍 Search Algorithm Performance:")
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		arr := make([]int, size)
		for i := range arr {
			arr[i] = i * 2
		}

		target := size - 100
		fmt.Printf("\nArray size: %d elements\n", size)

		stats := search.BenchmarkSearchAlgorithms(arr, target)
		search.PrintSearchResults(stats, target)
	}

	// Number Theory Performance
	fmt.Println("\n🔢 Number Theory Performance:")
	pg := number_theory.NewPrimeGenerator()

	primeLimits := []int{1000, 10000, 100000}
	for _, limit := range primeLimits {
		start := time.Now()
		primes := pg.SieveOfEratosthenes(limit)
		duration := time.Since(start)
		fmt.Printf("Sieve of Eratosthenes up to %d: %d primes found (Time: %v)\n",
			limit, len(primes), duration)
	}

	// Modular Exponentiation Performance
	fmt.Println("\n⚡ Modular Exponentiation Performance:")
	ma := number_theory.NewModularArithmetic()

	testCases := []struct {
		base, exp, mod int
	}{
		{2, 1000, 1000000007},
		{3, 10000, 1000000007},
		{5, 100000, 1000000007},
	}

	for _, tc := range testCases {
		start := time.Now()
		result := ma.ModPow(tc.base, tc.exp, tc.mod)
		duration := time.Since(start)
		fmt.Printf("ModPow(%d, %d, %d) = %d (Time: %v)\n",
			tc.base, tc.exp, tc.mod, result, duration)
	}
}

func demonstrateRealWorldApplications() {
	fmt.Println("Real-world applications of advanced search and number theory")

	applications := map[string][]string{
		"Advanced Search Algorithms": {
			"Database indexing with interpolation search for uniform data",
			"Binary search variants in compiler optimization",
			"Ternary search in machine learning hyperparameter tuning",
			"Exponential search in infinite data streams",
			"Range queries in time-series databases",
		},
		"Number Theory in Cryptography": {
			"RSA encryption using large prime generation",
			"Diffie-Hellman key exchange with modular exponentiation",
			"Digital signatures using modular arithmetic",
			"Hash functions with number theoretic properties",
			"Blockchain proof-of-work using prime-based puzzles",
		},
		"Mathematical Optimization": {
			"Financial portfolio optimization using ternary search",
			"Machine learning gradient descent optimization",
			"Engineering design optimization in aerospace",
			"Supply chain cost minimization",
			"Network routing optimization in telecommunications",
		},
		"Competitive Programming": {
			"Contest problems requiring advanced search techniques",
			"Number theory problems in mathematical competitions",
			"Algorithmic optimization challenges",
			"Code golf optimization using mathematical properties",
			"Real-time algorithm contests requiring fast implementations",
		},
	}

	for category, useCases := range applications {
		fmt.Printf("\n🔹 %s:\n", category)
		for _, useCase := range useCases {
			fmt.Printf("  • %s\n", useCase)
		}
	}

	fmt.Println("\n💡 Algorithm Selection Guidelines:")
	fmt.Println("✓ Uniform data distribution:     Interpolation Search")
	fmt.Println("✓ Infinite/unbounded data:       Exponential Search")
	fmt.Println("✓ Unimodal optimization:         Ternary Search")
	fmt.Println("✓ Cryptographic applications:    Miller-Rabin + Modular Arithmetic")
	fmt.Println("✓ Large number factorization:    Pollard's Rho Algorithm")
	fmt.Println("✓ System of congruences:         Chinese Remainder Theorem")
	fmt.Println("✓ Fast modular operations:       Binary Exponentiation")
	fmt.Println("✓ Prime generation at scale:     Segmented Sieve")

	fmt.Println("\n🎯 Complexity Achievements:")
	fmt.Println("• Binary Search Variants:        O(log n) with specialized optimizations")
	fmt.Println("• Ternary Search:                O(log₃ n) for unimodal functions")
	fmt.Println("• Sieve of Eratosthenes:         O(n log log n) prime generation")
	fmt.Println("• Modular Exponentiation:        O(log n) vs naive O(n)")
	fmt.Println("• Miller-Rabin Test:             O(k log³ n) probabilistic primality")
	fmt.Println("• Interpolation Search:          O(log log n) for uniform data")
}
