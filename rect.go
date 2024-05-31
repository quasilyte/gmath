package gmath

import (
	"image"
)

type Rect struct {
	Min Vec
	Max Vec
}

// RectFromStd converts an [image.Rectangle] into a [Rect].
// There is [Rect.ToStd] method to reverse it.
func RectFromStd(src image.Rectangle) Rect {
	return Rect{
		Min: VecFromStd(src.Min),
		Max: VecFromStd(src.Max),
	}
}

// ToStd converts an [Rect] into a [image.Rectangle].
// There is [RectFromStd] function to reverse it.
func (r Rect) ToStd() image.Rectangle {
	return image.Rectangle{
		Min: r.Min.ToStd(),
		Max: r.Max.ToStd(),
	}
}

func (r Rect) IsZero() bool {
	return r == Rect{}
}

func (r Rect) Width() float64 { return r.Max.X - r.Min.X }

func (r Rect) Height() float64 { return r.Max.Y - r.Min.Y }

// Center returns the center point of this rectangle.
//
// This center point may need some rounding,
// since a rect of a 3x3 size would return {1.5, 1.5}.
func (r Rect) Center() Vec {
	return Vec{
		X: r.Max.X - r.Width()*0.5,
		Y: r.Max.Y - r.Height()*0.5,
	}
}

func (r Rect) X1() float64 { return r.Min.X }

func (r Rect) Y1() float64 { return r.Min.Y }

func (r Rect) X2() float64 { return r.Max.X }

func (r Rect) Y2() float64 { return r.Max.Y }

func (r Rect) IsEmpty() bool {
	return r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y
}

func (r Rect) Contains(p Vec) bool {
	return r.Min.X <= p.X && p.X < r.Max.X &&
		r.Min.Y <= p.Y && p.Y < r.Max.Y
}

func (r Rect) ContainsRect(other Rect) bool {
	if other.IsEmpty() {
		return true
	}
	return r.Min.X <= other.Min.X && other.Max.X <= r.Max.X &&
		r.Min.Y <= other.Min.Y && other.Max.Y <= r.Max.Y
}

// Intersects reports whether r and other have a common intersection.
func (r Rect) Intersects(other Rect) bool {
	return !r.IsEmpty() && !other.IsEmpty() &&
		r.Min.X < other.Max.X && other.Min.X < r.Max.X &&
		r.Min.Y < other.Max.Y && other.Min.Y < r.Max.Y
}

func (r Rect) Add(p Vec) Rect {
	return Rect{
		Min: r.Min.Add(p),
		Max: r.Max.Add(p),
	}
}
