package main

import (
	"log"
)

func gaussMethod(matrix [][]float64) []float64 {
	for i := 0; i < len(matrix); i++ {
		for indx := i; indx < len(matrix); indx++ {
			if matrix[indx][i] != 0 {
				matrix[i], matrix[indx] = matrix[indx], matrix[i]
				break
			}
		}
		for k := i + 1; k < len(matrix); k++ {
			koef := matrix[k][i] / matrix[i][i]
			for indx := range matrix[k] {
				matrix[k][indx] -= matrix[i][indx] * koef
			}
		}
	}
	log.Print(matrix)
	l1 := len(matrix) - 1
	l2 := len(matrix[0]) - 1
	solution := make([]float64, l2)
	for i := 0; i <= l1; i++ {
		tmp1 := matrix[l1-i][l2]
		var tmp2 struct {
			value   float64
			isFirst bool
		}
		for indx, v := range matrix[l1-i][:l2] {
			if v != 0 && !tmp2.isFirst {
				tmp2.value = v
				tmp2.isFirst = true
			}
			tmp1 -= v * solution[indx]
		}
		solution[l2-1-i] = tmp1 / tmp2.value

	}

	return solution
}

func main() {
	matrix := [][]float64{
		{2, 1, -1, 8},
		{-3, -1, 2, -11},
		{-2, 1, 2, -3}}

	gaussMethod(matrix)
}
