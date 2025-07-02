package gmath

import (
	"testing"

	"slices"
)

func TestIntPercentages(t *testing.T) {
	tests := []struct {
		values []float64
		want   []int
	}{
		{
			[]float64{},
			nil,
		},

		{
			[]float64{1},
			[]int{100},
		},

		{
			[]float64{1, 1},
			[]int{50, 50},
		},

		{
			[]float64{2, 1},
			[]int{66, 34},
		},
		{
			[]float64{1, 2},
			[]int{33, 67},
		},

		{
			[]float64{0, 0, 1, 0},
			[]int{0, 0, 100, 0},
		},
		{
			[]float64{0, 0, 1, 0},
			[]int{0, 0, 100, 0},
		},

		{
			[]float64{3, 3, 3},
			[]int{33, 33, 34},
		},

		{
			[]float64{3, 3, 3, 3, 3},
			[]int{20, 20, 20, 20, 20},
		},
		{
			[]float64{3, 3, 3, 3, 3, 3},
			[]int{16, 16, 17, 17, 17, 17},
		},

		{
			[]float64{0, 0},
			[]int{0, 0},
		},
		{
			[]float64{0, 0, 0},
			[]int{0, 0, 0},
		},

		{
			[]float64{90, 11, 0},
			[]int{89, 11, 0},
		},
		{
			[]float64{11, 90, 0},
			[]int{10, 90, 0},
		},
		{
			[]float64{9.5, 90, 0},
			[]int{9, 91, 0},
		},
		{
			[]float64{9.5, 0, 90, 0},
			[]int{9, 0, 91, 0},
		},
		{
			[]float64{90, 9.5, 0},
			[]int{90, 10, 0},
		},
		{
			[]float64{90, 0, 9.5, 0},
			[]int{90, 0, 10, 0},
		},
	}

	for _, test := range tests {
		have := IntPercentages(test.values)
		sum := 0
		for _, v := range have {
			sum += v
		}
		if sum != 100 && sum != 0 {
			t.Fatalf("test values=%v:\ninvalid sum of %d (%v)", test.values, sum, have)
		}
		if !slices.Equal(test.want, have) {
			t.Fatalf("test values=%v:\nhave: %v\nwant: %v", test.values, have, test.want)
		}
	}
}
