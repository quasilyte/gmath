package gmath

import "github.com/quasilyte/gmath/fastmath"

func cubicInterpolate(from, to, pre, post, t float64) float64 {
	return 0.5 *
		((from * 2.0) +
			(-pre+to)*t +
			(2.0*pre-5.0*from+4.0*to-post)*(t*t) +
			(-pre+3.0*from-3.0*to+post)*(t*t*t))
}

func fposmod(x, y float64) float64 {
	value := fastmath.Mod(x, y)
	if ((value < 0) && (y > 0)) || ((value > 0) && (y < 0)) {
		value += y
	}
	return value
}
