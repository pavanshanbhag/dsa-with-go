package optimization

import (
	"fmt"
	"math"
)

// OptimizationResult represents the result of an optimization algorithm
type OptimizationResult struct {
	Point      float64
	Value      float64
	Iterations int
	Converged  bool
	Tolerance  float64
}

// NumericalOptimizer provides various numerical optimization methods
type NumericalOptimizer struct{}

// NewNumericalOptimizer creates a new numerical optimizer instance
func NewNumericalOptimizer() *NumericalOptimizer {
	return &NumericalOptimizer{}
}

// GoldenSectionSearch finds minimum of unimodal function using golden section method
// Time: O(log((b-a)/ε)), Space: O(1)
func (no *NumericalOptimizer) GoldenSectionSearch(f func(float64) float64, a, b, tolerance float64) OptimizationResult {
	phi := (1 + math.Sqrt(5)) / 2 // Golden ratio
	resphi := 2 - phi             // 1/phi

	// Initialize points
	x1 := a + resphi*(b-a)
	x2 := b - resphi*(b-a)
	f1 := f(x1)
	f2 := f(x2)

	iterations := 0
	maxIterations := 1000

	for math.Abs(b-a) > tolerance && iterations < maxIterations {
		iterations++

		if f1 < f2 {
			b = x2
			x2 = x1
			f2 = f1
			x1 = a + resphi*(b-a)
			f1 = f(x1)
		} else {
			a = x1
			x1 = x2
			f1 = f2
			x2 = b - resphi*(b-a)
			f2 = f(x2)
		}
	}

	minPoint := (a + b) / 2
	return OptimizationResult{
		Point:      minPoint,
		Value:      f(minPoint),
		Iterations: iterations,
		Converged:  math.Abs(b-a) <= tolerance,
		Tolerance:  tolerance,
	}
}

// TernarySearch finds minimum of strictly unimodal function
// Time: O(log₃((b-a)/ε)), Space: O(1)
func (no *NumericalOptimizer) TernarySearch(f func(float64) float64, left, right, tolerance float64) OptimizationResult {
	iterations := 0
	maxIterations := 1000

	for right-left > tolerance && iterations < maxIterations {
		iterations++

		// Divide interval into three parts
		m1 := left + (right-left)/3
		m2 := right - (right-left)/3

		// Evaluate function
		f1 := f(m1)
		f2 := f(m2)

		// Eliminate one-third
		if f1 < f2 {
			right = m2
		} else {
			left = m1
		}
	}

	minPoint := (left + right) / 2
	return OptimizationResult{
		Point:      minPoint,
		Value:      f(minPoint),
		Iterations: iterations,
		Converged:  right-left <= tolerance,
		Tolerance:  tolerance,
	}
}

// BisectionMethod finds root of continuous function using bisection
// Time: O(log((b-a)/ε)), Space: O(1)
func (no *NumericalOptimizer) BisectionMethod(f func(float64) float64, a, b, tolerance float64) OptimizationResult {
	if f(a)*f(b) >= 0 {
		return OptimizationResult{Converged: false}
	}

	iterations := 0
	maxIterations := 1000

	for math.Abs(b-a) > tolerance && iterations < maxIterations {
		iterations++

		c := (a + b) / 2
		fc := f(c)

		if math.Abs(fc) < tolerance {
			return OptimizationResult{
				Point:      c,
				Value:      fc,
				Iterations: iterations,
				Converged:  true,
				Tolerance:  tolerance,
			}
		}

		if f(a)*fc < 0 {
			b = c
		} else {
			a = c
		}
	}

	root := (a + b) / 2
	return OptimizationResult{
		Point:      root,
		Value:      f(root),
		Iterations: iterations,
		Converged:  math.Abs(b-a) <= tolerance,
		Tolerance:  tolerance,
	}
}

// NewtonRaphsonMethod finds root using Newton-Raphson method
// Requires function and its derivative
func (no *NumericalOptimizer) NewtonRaphsonMethod(f, df func(float64) float64, x0, tolerance float64) OptimizationResult {
	x := x0
	iterations := 0
	maxIterations := 1000

	for iterations < maxIterations {
		iterations++

		fx := f(x)
		dfx := df(x)

		if math.Abs(dfx) < 1e-12 {
			return OptimizationResult{Converged: false} // Derivative too small
		}

		xNew := x - fx/dfx

		if math.Abs(xNew-x) < tolerance {
			return OptimizationResult{
				Point:      xNew,
				Value:      f(xNew),
				Iterations: iterations,
				Converged:  true,
				Tolerance:  tolerance,
			}
		}

		x = xNew
	}

	return OptimizationResult{
		Point:      x,
		Value:      f(x),
		Iterations: iterations,
		Converged:  false,
		Tolerance:  tolerance,
	}
}

// SecantMethod finds root using secant method (doesn't require derivative)
func (no *NumericalOptimizer) SecantMethod(f func(float64) float64, x0, x1, tolerance float64) OptimizationResult {
	iterations := 0
	maxIterations := 1000

	for iterations < maxIterations {
		iterations++

		f0 := f(x0)
		f1 := f(x1)

		if math.Abs(f1-f0) < 1e-12 {
			return OptimizationResult{Converged: false} // Slope too small
		}

		x2 := x1 - f1*(x1-x0)/(f1-f0)

		if math.Abs(x2-x1) < tolerance {
			return OptimizationResult{
				Point:      x2,
				Value:      f(x2),
				Iterations: iterations,
				Converged:  true,
				Tolerance:  tolerance,
			}
		}

		x0, x1 = x1, x2
	}

	return OptimizationResult{
		Point:      x1,
		Value:      f(x1),
		Iterations: iterations,
		Converged:  false,
		Tolerance:  tolerance,
	}
}

// FibonacciSearch implements Fibonacci search for optimization
func (no *NumericalOptimizer) FibonacciSearch(f func(float64) float64, a, b float64, n int) OptimizationResult {
	// Generate Fibonacci numbers
	fib := make([]int, n+2)
	fib[0], fib[1] = 1, 1
	for i := 2; i <= n+1; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	// Initialize points
	x1 := a + float64(fib[n-1])/float64(fib[n+1])*(b-a)
	x2 := a + float64(fib[n])/float64(fib[n+1])*(b-a)
	f1 := f(x1)
	f2 := f(x2)

	for k := 1; k <= n; k++ {
		if f1 < f2 {
			b = x2
			x2 = x1
			f2 = f1
			x1 = a + float64(fib[n-k-1])/float64(fib[n-k+1])*(b-a)
			f1 = f(x1)
		} else {
			a = x1
			x1 = x2
			f1 = f2
			x2 = a + float64(fib[n-k])/float64(fib[n-k+1])*(b-a)
			f2 = f(x2)
		}
	}

	minPoint := (a + b) / 2
	return OptimizationResult{
		Point:      minPoint,
		Value:      f(minPoint),
		Iterations: n,
		Converged:  true,
		Tolerance:  math.Abs(b - a),
	}
}

// GradientDescent performs gradient descent optimization
func (no *NumericalOptimizer) GradientDescent(f, df func(float64) float64, x0, learningRate, tolerance float64) OptimizationResult {
	x := x0
	iterations := 0
	maxIterations := 10000

	for iterations < maxIterations {
		iterations++

		gradient := df(x)
		xNew := x - learningRate*gradient

		if math.Abs(xNew-x) < tolerance {
			return OptimizationResult{
				Point:      xNew,
				Value:      f(xNew),
				Iterations: iterations,
				Converged:  true,
				Tolerance:  tolerance,
			}
		}

		x = xNew
	}

	return OptimizationResult{
		Point:      x,
		Value:      f(x),
		Iterations: iterations,
		Converged:  false,
		Tolerance:  tolerance,
	}
}

// NumericalIntegration provides numerical integration methods
type NumericalIntegration struct{}

// NewNumericalIntegration creates a new numerical integration instance
func NewNumericalIntegration() *NumericalIntegration {
	return &NumericalIntegration{}
}

// TrapezoidalRule computes definite integral using trapezoidal rule
func (ni *NumericalIntegration) TrapezoidalRule(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := (f(a) + f(b)) / 2

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		sum += f(x)
	}

	return sum * h
}

// SimpsonsRule computes definite integral using Simpson's rule
func (ni *NumericalIntegration) SimpsonsRule(f func(float64) float64, a, b float64, n int) float64 {
	if n%2 != 0 {
		n++ // Ensure n is even
	}

	h := (b - a) / float64(n)
	sum := f(a) + f(b)

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}

	return sum * h / 3
}

// AdaptiveIntegration performs adaptive numerical integration
func (ni *NumericalIntegration) AdaptiveIntegration(f func(float64) float64, a, b, tolerance float64) float64 {
	return ni.adaptiveHelper(f, a, b, tolerance, ni.SimpsonsRule(f, a, b, 2))
}

func (ni *NumericalIntegration) adaptiveHelper(f func(float64) float64, a, b, tolerance, wholeArea float64) float64 {
	mid := (a + b) / 2
	leftArea := ni.SimpsonsRule(f, a, mid, 2)
	rightArea := ni.SimpsonsRule(f, mid, b, 2)
	sumArea := leftArea + rightArea

	if math.Abs(sumArea-wholeArea) < tolerance {
		return sumArea
	}

	return ni.adaptiveHelper(f, a, mid, tolerance/2, leftArea) +
		ni.adaptiveHelper(f, mid, b, tolerance/2, rightArea)
}

// Demo demonstrates various optimization techniques
func (no *NumericalOptimizer) Demo() {
	fmt.Println("=== Numerical Optimization Demo ===")

	// Test function: f(x) = (x-2)² + 1, minimum at x=2, value=1
	testFunc := func(x float64) float64 {
		return (x-2)*(x-2) + 1
	}

	// Test derivative: f'(x) = 2(x-2)
	testDerivative := func(x float64) float64 {
		return 2 * (x - 2)
	}

	tolerance := 1e-6

	// Golden Section Search
	result := no.GoldenSectionSearch(testFunc, 0, 5, tolerance)
	fmt.Printf("Golden Section: min at x=%.6f, f(x)=%.6f, iterations=%d\n",
		result.Point, result.Value, result.Iterations)

	// Ternary Search
	result = no.TernarySearch(testFunc, 0, 5, tolerance)
	fmt.Printf("Ternary Search: min at x=%.6f, f(x)=%.6f, iterations=%d\n",
		result.Point, result.Value, result.Iterations)

	// Gradient Descent
	result = no.GradientDescent(testFunc, testDerivative, 0, 0.1, tolerance)
	fmt.Printf("Gradient Descent: min at x=%.6f, f(x)=%.6f, iterations=%d\n",
		result.Point, result.Value, result.Iterations)

	// Root finding for f'(x) = 0
	result = no.BisectionMethod(testDerivative, 0, 5, tolerance)
	fmt.Printf("Bisection (root): x=%.6f, f'(x)=%.6f, iterations=%d\n",
		result.Point, result.Value, result.Iterations)

	// Newton-Raphson for root finding
	// Second derivative: f''(x) = 2
	secondDerivative := func(x float64) float64 { return 2 }
	result = no.NewtonRaphsonMethod(testDerivative, secondDerivative, 0, tolerance)
	fmt.Printf("Newton-Raphson: x=%.6f, f'(x)=%.6f, iterations=%d\n",
		result.Point, result.Value, result.Iterations)

	// Numerical Integration
	ni := NewNumericalIntegration()
	integral := ni.TrapezoidalRule(testFunc, 0, 4, 1000)
	fmt.Printf("Trapezoidal integration ∫₀⁴ f(x)dx = %.6f\n", integral)

	integral = ni.SimpsonsRule(testFunc, 0, 4, 1000)
	fmt.Printf("Simpson's integration ∫₀⁴ f(x)dx = %.6f\n", integral)
}
