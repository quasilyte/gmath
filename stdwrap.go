package gmath

import (
	"math"
)

func Abs[T float](x T) T {
	return T(math.Abs(float64(x)))
}
