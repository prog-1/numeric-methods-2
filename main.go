package main

import (
	"testing"
)

func main() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test FindMinimum", F: TestGauss},
		},
		/* benchmarks */ nil /* examples */, nil)
}

// Solves system of linear equations written in the matrix form
func gauss(A [][]float64) []float64 {
	n := len(A)
	// 1. Finding a row with in eliminated coefficient
	for i := 0; i < n; i++ {
		if A[i][0] != 0 {
			A[0], A[i] = A[i], A[0]
			break
		}
	}
	// 2. Eliminating
	for j := 0; j < n-1; j++ { // Columns
		for i := j + 1; i < n; i++ { // Rows
			c := A[i][j] / A[j][j]    // Subtrahend column multiplier
			for k := j; k <= n; k++ { // Columns
				A[i][k] -= c * A[j][k] // Elimination
			}
		}
	}
	// 3. Solving
	res := make([]float64, len(A))
	for i := n - 1; i >= 0; i-- { // Rows
		res[i] = A[i][n]
		for j := i + 1; j < n; j++ {
			res[i] -= A[i][j] * res[j]
		}
		res[i] /= A[i][i]
	}
	return res
}

func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGauss(t *testing.T) {
	for num, tc := range []struct {
		input [][]float64
		want  []float64
	}{
		{[][]float64{{2, 1, -1, 8}, {-3, -1, 2, -11}, {-2, 1, 2, -3}}, []float64{2, 3, -1}},
		{[][]float64{{2, 3}}, []float64{1.5}},
	} {
		if got := gauss(tc.input); !equal(got, tc.want) {
			t.Errorf("Test No.%v failed: got = %v, want = %v", num, got, tc.want)
		}
	}
}
