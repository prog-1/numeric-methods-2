package main

import (
	"fmt"
)

func gauss(n int, m [][]float64) (res []float64) {
	for i := 0; i < n; i++ {
		if m[i][i] == 0 {
			for l := i + 1; l < n; l++ {
				if m[l][l] != 0 {
					m[i], m[l] = m[l], m[i]
				}
			}
		}
		for j := i + 1; j < n; j++ {
			coef := m[j][i] / m[i][i]
			for l := i; l < n+1; l++ {
				m[j][l] -= m[i][l] * coef
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for l := n - 1; l > i; l-- {
			m[i][n] = m[i][n] - m[i][l]
		}
		m[i][n] /= m[i][i]
		res = append(res, m[i][n])
		for j := i - 1; j >= 0; j-- {
			m[j][i] = m[j][i] * m[i][n]
		}
	}
	return
}

func main() {
	var n int
	fmt.Print("Enter number of unknown variables: ")
	fmt.Scan(&n)
	m := make([][]float64, n)
	fmt.Println("Enter coefficients of system of linear equations:")
	for i := 0; i < n; i++ {
		m[i] = make([]float64, n+1)
		for j := 0; j < n+1; j++ {
			fmt.Scan(&m[i][j])
		}
	}
	res := gauss(n, m)
	fmt.Print("The unknown variables: ")
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i], " ")
	}
}
