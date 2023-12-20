package main

import "fmt"

func main() {
	// NOTE: Vectors in matrices are represented as rows for matching with the way systems of linear equations are written
	// NOTE: We are storing output vector as the last column just for convenience.
	// NOTE: The matrix must have dimensions n * (n+1), but we will not check for correspondance, just assume it is correct.
	// NOTE: Min n = 1
	A := [][]float64{{2, 1, -1, 8}, {-3, -1, 2, -11}, {-2, 1, 2, -3}}
	fmt.Println(A)
	fmt.Println(gauss(A))
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
			if i == 2 {
				fmt.Println("Sorting things out")
			}
			c := A[i][j] / A[j][j]    // Subtrahend column multiplier
			for k := j; k <= n; k++ { // Columns
				A[i][k] -= c * A[j][k] // Elimination
			}
		}
	}
	fmt.Println(A)
	//
	f := make([]float64, len(A))
	return f
}
