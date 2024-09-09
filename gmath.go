package gmath

import (
	"math"
)

const Epsilon = 1e-9

func EqualApprox[T float](a, b T) bool {
	return math.Abs(float64(a-b)) <= Epsilon
}

// Lerp linearly interpolates between [from] and [to] using the weight [t].
// The [t] value is usually in the range from 0 to 1.
func Lerp[T float](from, to, t T) T {
	return from + ((to - from) * t)
}

// InvLerp returns an interpolation or extrapolation factor considering the given range and weight [t].
// The [t] value is usually in the range from 0 to 1.
func InvLerp[T float](from, to, value T) T {
	return (value - from) / (to - from)
}

// Remap maps a value from one range to another.
//
// The first range is defined by a pair of [fromMin] and [fromMax].
// The second range is defined by a pair of [toMin] and [toMax].
// The [t] value is usually in the range from 0 to 1.
func Remap[T float](fromMin, fromMax, toMin, toMax, value T) T {
	t := InvLerp(fromMin, fromMax, value)
	return Lerp(toMin, toMax, t)
}

func ArcContains(angle, measure Rad, pos, point Vec) bool {
	startAngle := (angle - measure/2)
	endAngle := (angle + measure/2)
	if endAngle < startAngle {
		endAngle += 2 * math.Pi
	}
	half := (endAngle - startAngle) / 2
	mid := (endAngle + startAngle) / 2
	coshalf := math.Cos(float64(half))
	angleToPoint := pos.AngleToPoint(point).Normalized()
	return math.Cos(float64(angleToPoint-mid)) >= coshalf
}

func ArcSectionContains(angle, measure Rad, r float64, pos, point Vec) bool {
	if pos.DistanceSquaredTo(point) > r*r {
		return false
	}
	return ArcContains(angle, measure, pos, point)
}

func Clamp[T numeric](v, min, max T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func ClampMin[T numeric](v, min T) T {
	if v < min {
		return min
	}
	return v
}

func ClampMax[T numeric](v, max T) T {
	if v > max {
		return max
	}
	return v
}

func Percentage[T numeric](value, max T) T {
	if max == 0 && value == 0 {
		return 0
	}
	return T(100 * (float64(value) / float64(max)))
}

// Scale multiplies value by m.
//
// For a floating-point value this operation would
// not make any sense as it can be expressed value*m,
// but integer value scaling with a floating point m
// involves conversions and rounding.
//
// This function rounds the scaled integer to the
// nearest integer result.
// For example:
//   - Scale(1, 1.8) => 2 (rounds up)
//   - Scale(1, 1.4) => 1 (rounds down)
//   - Scale(1, 1)   => 1
//   - Scale(1, 0)   => 0
func Scale[T integer](value T, m float64) T {
	return T(math.Round(float64(value) * m))
}

// Iround is a helper to perform int(math.Round(x)) operation.
//
// This function reduces the number of parenthesis the final
// expression will have.
func Iround(x float64) int {
	return int(math.Round(x))
}

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type numeric interface {
	integer | float
}

type float interface {
	~float32 | ~float64
}
