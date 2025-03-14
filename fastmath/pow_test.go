package fastmath

import (
	"math"
	"testing"
)

var powTestValues = []float64{
	0,

	1, 5, 9, 140, 19495, 8988538, 912435823852384,

	0.1, 0.114, 1.3, 1.7, 1.7949, 2.5436, 8345.234895, 999.9999991,
}

func TestIntPow(t *testing.T) {
	exponents := []int{
		0,
		1, 2, 3, 6, 19, 95,
		-1, -5, -94,
	}
	for _, v := range powTestValues {
		for _, exp := range exponents {
			have := IntPow(v, exp)
			want := math.Pow(v, float64(exp))
			errorPercent := deviation(have, want) - 1
			if errorPercent >= 0.0001 {
				t.Errorf("pow(%f, %f):\nhave: %f\nwant: %f\nerror: %f", float64(exp), v, have, want, errorPercent)
			}
		}
	}
}

func TestPow1_5(t *testing.T) {
	for _, v := range powTestValues {
		have := Pow1_5(v)
		want := math.Pow(v, 1.5)
		errorPercent := deviation(have, want) - 1
		if errorPercent >= 0.0001 {
			t.Errorf("pow(%f, 1.5):\nhave: %f\nwant: %f\nerror: %f", v, have, want, errorPercent)
		}
	}
}
