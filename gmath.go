package gmath

import (
	"math"
)

const Epsilon = 1e-9

func EqualApprox[T float](a, b T) bool {
	return math.Abs(float64(a-b)) <= Epsilon
}

// CeilN applies Ceil(x) and then rounds up to the closest n multiplier.
//
// This function is useful when trying to turn an arbitrary value
// into a pretty value. E.g. it can turn a price of 139 into 140
// if x=139 and n=5.
//
// The return value is float, but it's guaranteed to be integer-like
// (unless you're working with really high x values).
//
// * CeilN(0, 0) => 0
// * CeilN(0, 2) => 0
// * CeilN(1, 2) => 2
// * CeilN(2, 2) => 2
// * CeilN(3, 2) => 4
// * CeilN(144, 5) => 145
// * CeilN(145, 5) => 145
// * CeilN(146, 5) => 150
func CeilN(x float64, n int) float64 {
	iv := int(math.Ceil(x))

	extra := 0
	if n != 0 {
		extra = n - (iv % n)
		if extra == n {
			extra = 0
		}
	}

	return float64(iv + extra)
}

// ScaledSum calculates an arithmetic progression that is
// often used in level-up experience requirement scaling.
// For example, if baseValue (exp for level 2) is 100,
// and the step (increase) is 25, then we have these values per level:
// * level=0 => 0   (the default value)
// * level=1 => 100 (+100)
// * level=2 => 225 (+125)
// * level=3 => 375 (+150)
// * level=4 => 550 (+175)
// ...
// It can also handle fractional level values:
// * level=1.5 => ~160 for the example above
//
// As a special case, it always returns 0 for levels<=0.
func ScaledSum(baseValue, step, level float64) float64 {
	if level <= 0 {
		return 0
	}
	k := level
	return (k * (2*baseValue + (k-1)*step)) * 0.5
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

// LerpClamped linearly interpolates between [from] and [to] using the weight [t].
// The [t] value is clamped to the range from 0 to 1.
func LerpClamped[T float](from, to, t T) T {
	t = Clamp(t, 0, 1)
	return from + ((to - from) * t)
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

// Deviation reports the weight of the difference between x and y.
//
// In other words, it returns the multiplier to make the lower value
// between the two identical to other.
//
// Some examples:
// * (1, 2) => 2 (1*2 = 2)
// * (10, 2) => 5 (2*5 = 10)
// * (3, 3) => 1
//
// The order of parameters do not affect the result.
func Deviation[T float](x, y T) T {
	if x == y {
		return 1
	}

	if x == 0 {
		x += Epsilon
	}
	if y == 0 {
		y += Epsilon
	}

	if x < y {
		x, y = y, x
	}
	return x / y
}

// InBounds reports whether v is in the specified inclusive range.
func InBounds[T numeric](v, min, max T) bool {
	return v >= min && v <= max
}

type integer interface {
	signedInteger |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type signedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type numeric interface {
	integer | float
}

type float interface {
	~float32 | ~float64
}
