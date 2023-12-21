package main

import "fmt"

func GaussElimination(a [][]float64, n int) (ans []float64) {
	ans = make([]float64, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i][i] == 0 { // check if needed to swap rows
				if a[j][i] != 0 { // swap rows
					a[i], a[j] = a[j], a[i]
				}
			}
			// make [j] column for lines [i+1;...] equal to zero (triangular matrix)
			m := a[j][i] / a[i][i] // number to multiply the row [i] to get zero in [j] column
			for k := i; k <= n; k++ {
				a[j][k] += -m * a[i][k] // m is negative because we need to get zero, not 2*number after summation
			}
		}
	}
	for i := n - 1; i >= 0; i-- { // go back, calculating the answer
		ans[i] = a[i][n] / a[i][i]
		for j := i - 1; j >= 0; j-- {
			a[j][n] -= a[j][i] * ans[i]
		}
	}
	return ans
} // whole 5 loops for 1 function... sad

func main() {
	fmt.Println("Please, enter the number of equations: ")
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("What you expected? Magic? I read the equation from your mind?")
		return
	}
	a := make([][]float64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n+1)
		fmt.Printf("Please, enter the coefficients of the equations № %v: ", i+1)
		for j := 0; j < n+1; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	ans := GaussElimination(a, n)
	fmt.Printf("The answer is: %v", ans)
}
