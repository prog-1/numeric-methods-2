package main

import "fmt"

func GaussianElimination(n int, s [][]float64) (res []float64) {
	for i := range s { // elimination
		for j := i + 1; j < n; j++ {
			d := s[j][i] / s[i][i]
			if (s[j][i] > 0 && s[i][i] > 0) || (s[j][i] < 0 && s[i][i] > 0) {
				d *= -1
			}
			for k := range s[j] {
				s[j][k] += s[i][k] * d
			}
		}
	}
	for i := n - 1; i >= 0; i-- { // back-substitution
		d := s[i][n] / s[i][i]
		res = append(res, d)
		for j := 0; j < i; j++ {
			s[j][n] -= s[j][i] * d
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 { // reverse the resulting slice
		res[i], res[j] = res[j], res[i]
	}
	return
}

func main() {
	s := [][]float64{
		{2, 1, -1, 8},
		{-3, -1, 2, -11},
		{-2, 1, 2, -3},
	}
	fmt.Println(GaussianElimination(3, s))
}
