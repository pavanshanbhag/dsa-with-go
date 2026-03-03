package optimization

import (
	"math"
	"testing"
)

func TestGoldenSectionSearch(t *testing.T) {
	no := NewNumericalOptimizer()

	// Minimize (x-3)^2 + 2, minimum at x=3
	f := func(x float64) float64 { return (x-3)*(x-3) + 2 }
	result := no.GoldenSectionSearch(f, 0, 6, 1e-6)

	if !result.Converged {
		t.Errorf("GoldenSectionSearch: expected converged")
	}
	if math.Abs(result.Point-3) > 1e-4 {
		t.Errorf("GoldenSectionSearch: point ≈ %v, want ≈ 3", result.Point)
	}
	if math.Abs(result.Value-2) > 1e-4 {
		t.Errorf("GoldenSectionSearch: value ≈ %v, want ≈ 2", result.Value)
	}
}

func TestTernarySearch(t *testing.T) {
	no := NewNumericalOptimizer()

	// Minimize x^2, minimum at x=0
	f := func(x float64) float64 { return x * x }
	result := no.TernarySearch(f, -2, 2, 1e-6)

	if !result.Converged {
		t.Errorf("TernarySearch: expected converged")
	}
	if math.Abs(result.Point) > 1e-4 {
		t.Errorf("TernarySearch: point ≈ %v, want ≈ 0", result.Point)
	}
}

func TestBisectionMethod(t *testing.T) {
	no := NewNumericalOptimizer()

	// Root of x^2 - 4 in [1, 3] is x=2
	f := func(x float64) float64 { return x*x - 4 }
	result := no.BisectionMethod(f, 1, 3, 1e-8)

	if !result.Converged {
		t.Errorf("BisectionMethod: expected converged")
	}
	if math.Abs(result.Point-2) > 1e-5 {
		t.Errorf("BisectionMethod: point ≈ %v, want ≈ 2", result.Point)
	}
	if math.Abs(result.Value) > 1e-5 {
		t.Errorf("BisectionMethod: |f(point)| ≈ %v, want ≈ 0", result.Value)
	}
}

func TestBisectionMethodInvalidInterval(t *testing.T) {
	no := NewNumericalOptimizer()

	// x^2 + 1 has no real root; f(0)*f(2) > 0
	f := func(x float64) float64 { return x*x + 1 }
	result := no.BisectionMethod(f, 0, 2, 1e-6)

	if result.Converged {
		t.Errorf("BisectionMethod: expected not converged when no sign change")
	}
}
