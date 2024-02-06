package main

import (
	"reflect"
	"testing"
)

func TestGauss(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    int
		s    [][]float64
		want []float64
	}{
		{"1", 1, [][]float64{{2, 3}}, []float64{1.5}},
		{"2", 3, [][]float64{{2, 1, -1, 8}, {-3, -1, 2, -11}, {-2, 1, 2, -3}}, []float64{2, 3, -1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := GaussianElimination(tc.n, tc.s); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
