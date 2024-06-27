package gmath

import (
	"math"
)

const Epsilon = 1e-9

func EqualApprox[T float](a, b T) bool {
	return math.Abs(float64(a-b)) <= Epsilon
}

func Lerp(from, to, t float64) float64 {
	return from + ((to - from) * t)
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

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		float
}

type float interface {
	~float32 | ~float64
}
