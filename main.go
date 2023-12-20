package main

import (
	"fmt"
)

func gaussianElimination(n int, m [][]float64) []float64 {
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if m[j][i] != 0 {
				m[i], m[j] = m[j], m[i]
				break
			}
		}
		for k := i + 1; k < n; k++ {
			q := m[k][i] / m[i][i]
			for j := i; j < n+1; j++ {
				m[k][j] -= q * m[i][j]
			}
		}
	}

	ans := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		var sum float64
		for j := i + 1; j < n; j++ {
			sum += m[i][j] * ans[j]
		}
		ans[i] = (m[i][n] - sum) / m[i][i]
	}
	return ans
}

func main() {
	n := 3
	m := [][]float64{
		{2, 1, -1, 8},
		{-3, -1, 2, -11},
		{-2, 1, 2, -3},
	}
	ans := gaussianElimination(n, m)
	for i, x := range ans {
		fmt.Printf("x%v = %v; ", i+1, x)
	}
}
