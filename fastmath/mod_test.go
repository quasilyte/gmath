package fastmath

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const eps = 1e-9

func equalApprox(a, b float64) bool {
	return math.Abs(a-b) <= eps
}

func deviation(x, y float64) float64 {
	if x == y {
		return 1
	}

	if x == 0 {
		x += eps
	}
	if y == 0 {
		y += eps
	}

	if x < y {
		x, y = y, x
	}
	return x / y
}

func TestMod(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	multipliers := []float64{1, -1}
	for _, m := range multipliers {
		for i := 0; i < 5000; i++ {
			x := r.Float64() * m
			y := r.Float64() * m
			have := Mod(x, y)
			want := math.Mod(x, y)
			// if have != want {
			if !equalApprox(have, want) {
				t.Fatalf("[%d] mod(%v, %v):\nhave=%v\nwant=%v", i, x, y, have, want)
			}
		}
	}
}
