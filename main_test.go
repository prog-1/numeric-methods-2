package main

import (
	"reflect"
	"testing"
)

func TestGaussElimination(t *testing.T) {
	for _, tc := range []struct {
		n    int
		a    [][]float64
		want []float64
	}{
		{3, [][]float64{{1, 2, -1, 2}, {1, 1, -1, 0}, {2, -1, 1, 3}}, []float64{1, 2, 3}},
		{3, [][]float64{{2, 1, -1, 8}, {-3, -1, 2, -11}, {-2, 1, 2, -3}}, []float64{2, 3, -1}},
		{2, [][]float64{{-9, -6, 24}, {1, 1, -4}}, []float64{0, -4}},
		//{2, [][]float64{{1, 1, 3}, {1, 1, 3}}, []float64{NaN, NaN}},
		{4, [][]float64{{-2, 3, 0, -7, 15}, {-1, -1, -5, 3, 35}, {7, -8, 5, 8, -80}, {0, -2, -10, -7, 76}}, []float64{3, 7, -9, 0}}, //megaa
		//{0, [][]float64{{}}, []float64{}},
		{1, [][]float64{{2, 4}}, []float64{2}},
	} {
		if got := GaussElimination(tc.a, len(tc.a)); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("GaussElimination(%v) = %v, want %v", tc.a, got, tc.want)
		}
	}

}
