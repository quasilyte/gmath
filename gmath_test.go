package gmath

import "testing"

func TestCeilN(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{0, 0, 0},
		{0, 2, 0},
		{1, 2, 2},
		{2, 2, 2},
		{3, 2, 4},
		{144, 5, 145},
		{145, 5, 145},
		{146, 5, 150},
	}

	for _, test := range tests {
		have := CeilN(test.x, test.n)
		if have != test.want {
			t.Fatalf("ceiln(x=%f, n=%d):\nhave: %v\nwant: %v", test.x, test.n, have, test.want)
		}
	}
}

func TestDeviation(t *testing.T) {
	tests := []struct {
		x      float64
		y      float64
		result float64
	}{
		{1, 1, 1},
		{0, 0, 1},

		{1, 2, 2},
		{2, 1, 2},
		{100, 200, 2},
		{200, 100, 2},
		{50, 200, 4},
		{200, 50, 4},
	}

	for i := range tests {
		test := tests[i]
		have := Deviation(test.x, test.y)
		want := test.result
		if !EqualApprox(have, want) {
			t.Fatalf("results mismatched for (%v, %v):\nhave: %v\nwant: %v",
				test.x, test.y, have, want)
		}
	}
}
