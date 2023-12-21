package main

import "testing"

func TestGaussianElimination(t *testing.T) {
	for _, tc := range []struct {
		n    int
		m    [][]float64
		want []float64
	}{
		{1, [][]float64{
			{2, 1, -1, 8},
			{-3, -1, 2, -11},
			{-2, 1, 2, -3},
		}, []float64{2, 3, -1}},
		{2, [][]float64{
			{1, 2, 11},
			{3, -1, 12},
		}, []float64{5, 3}},
	} {
		got := gaussMethod(tc.m)
		if !equal(got, tc.want) {
			t.Errorf("Test%v : Got %v, want = %v", tc.n, got, tc.want)
		}
	}
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
