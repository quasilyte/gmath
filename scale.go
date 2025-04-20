package gmath

import "math"

// LogScale is a faster alternative to math.Pow(x, scale).
// You need to adjust the scale parameter for it to give
// the compatible results.
// An example usage is math.Pow(x, 1.05) => LogScale(x, 0.06)
//
// LogScale is usually 5-10 times faster than math.Pow.
func LogScale(x float64, scale float64) float64 {
	return x * (1 + scale*math.Log(x))
}

// SqrScale is a faster alternative to math.Pow(x, scale).
// You need to adjust the scale parameter for it to give
// the compatible results.
//
// Note that due to the x^2 nature it will grow very fast
// even with a smaller scale value.
// You may want to pre-multiply before applying this function
// or use it only for relatively small x (e.g. values <=100k).
func SqrScale(x float64, scale float64) float64 {
	return x + scale*x*x
}
