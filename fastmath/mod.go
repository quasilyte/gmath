package fastmath

import (
	"math"
)

func Mod(x, y float64) float64 {
	return x - y*math.Floor(x/y)
}
