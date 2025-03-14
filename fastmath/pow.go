package fastmath

import (
	"math"
)

func IntPow(x float64, y int) float64 {
	if y < 0 {
		return 1 / IntPow(x, -y)
	}

	result := 1.0
	exp := y
	base := x
	for exp != 0 {
		if (exp & 1) != 0 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

func Pow1_5(x float64) float64 {
	return x * math.Sqrt(x)
}
