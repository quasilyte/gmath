package gmath

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"math"
	"strconv"
)

// Vec is a 2-element structure that is used to represent positions,
// velocities, and other kinds numerical pairs.
//
// Its implementation as well as its API is inspired by Vector2 type
// of the Godot game engine. Where feasible, its adjusted to fit Go
// coding conventions better. Also, earlier versions of Godot used
// 32-bit values for X and Y; our vector uses 64-bit values.
//
// Since Go has no operator overloading, we implement scalar forms of
// operations with "f" suffix. So, Add() is used to add two vectors
// while Addf() is used to add scalar to the vector.
type Vec struct {
	X float64
	Y float64
}

// RadToVec converts a given angle into a normalized vector that encodes that direction.
func RadToVec(angle Rad) Vec {
	return Vec{X: angle.Cos(), Y: angle.Sin()}
}

// VecFromStd converts an [image.Point] into a [Vec].
// There is [Vec.ToStd] method to reverse it.
func VecFromStd(src image.Point) Vec {
	return Vec{
		X: float64(src.X),
		Y: float64(src.Y),
	}
}

// ToStd converts [Vec] into [image.Point].
// There is [VecFromStd] function to reverse it.
func (v Vec) ToStd() image.Point {
	return image.Point{
		X: int(v.X),
		Y: int(v.Y),
	}
}

// String returns a pretty-printed representation of a 2D vector object.
func (v Vec) String() string {
	return fmt.Sprintf("[%f, %f]", v.X, v.Y)
}

// IsZero reports whether v is a zero value vector.
// A zero value vector has X=0 and Y=0, created with Vec{}.
//
// The zero value vector has a property that its length is 0,
// but not all zero length vectors are zero value vectors.
func (v Vec) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// IsNormalizer reports whether the vector is normalized.
// A vector is considered to be normalized if its length is 1.
func (v Vec) IsNormalized() bool {
	return EqualApprox(v.LenSquared(), 1)
}

// DistanceTo calculates the distance between the two vectors.
func (v Vec) DistanceTo(v2 Vec) float64 {
	return math.Sqrt((v.X-v2.X)*(v.X-v2.X) + (v.Y-v2.Y)*(v.Y-v2.Y))
}

func (v Vec) DistanceSquaredTo(v2 Vec) float64 {
	return ((v.X - v2.X) * (v.X - v2.X)) + ((v.Y - v2.Y) * (v.Y - v2.Y))
}

// Dot returns a dot-product of the two vectors.
func (v Vec) Dot(v2 Vec) float64 {
	return (v.X * v2.X) + (v.Y * v2.Y)
}

// Len reports the length of this vector (also known as magnitude).
func (v Vec) Len() float64 {
	return math.Sqrt(v.LenSquared())
}

// LenSquared returns the squared length of this vector.
//
// This function runs faster than Len(),
// so prefer it if you need to compare vectors
// or need the squared distance for some formula.
func (v Vec) LenSquared() float64 {
	return v.Dot(v)
}

func (v Vec) Rotated(angle Rad) Vec {
	sine := angle.Sin()
	cosi := angle.Cos()
	return Vec{
		X: v.X*cosi - v.Y*sine,
		Y: v.X*sine + v.Y*cosi,
	}
}

func (v Vec) Angle() Rad {
	return Rad(math.Atan2(v.Y, v.X))
}

// AngleToPoint returns the angle from v towards the given point.
func (v Vec) AngleToPoint(pos Vec) Rad {
	return pos.Sub(v).Angle()
}

func (v Vec) DirectionTo(v2 Vec) Vec {
	return v.Sub(v2).Normalized()
}

func (v Vec) VecTowards(pos Vec, length float64) Vec {
	angle := v.AngleToPoint(pos)
	result := Vec{X: angle.Cos(), Y: angle.Sin()}
	return result.Mulf(length)
}

func (v Vec) MoveTowards(pos Vec, length float64) Vec {
	direction := pos.Sub(v) // Not normalized
	dist := direction.Len()
	if dist <= length || dist < Epsilon {
		return pos
	}
	return v.Add(direction.Divf(dist).Mulf(length))
}

func (v Vec) EqualApprox(other Vec) bool {
	return EqualApprox(v.X, other.X) && EqualApprox(v.Y, other.Y)
}

func (v Vec) MoveInDirection(dist float64, dir Rad) Vec {
	return Vec{
		X: v.X + dist*dir.Cos(),
		Y: v.Y + dist*dir.Sin(),
	}
}

func (v Vec) Mulf(scalar float64) Vec {
	return Vec{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vec) Mul(other Vec) Vec {
	return Vec{
		X: v.X * other.X,
		Y: v.Y * other.Y,
	}
}

func (v Vec) Divf(scalar float64) Vec {
	return Vec{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v Vec) Div(other Vec) Vec {
	return Vec{
		X: v.X / other.X,
		Y: v.Y / other.Y,
	}
}

func (v Vec) Add(other Vec) Vec {
	return Vec{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec) Sub(other Vec) Vec {
	return Vec{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vec) Rounded() Vec {
	return Vec{
		X: math.Round(v.X),
		Y: math.Round(v.Y),
	}
}

// Normalized returns the vector scaled to unit length.
// Functionally equivalent to `v.Divf(v.Len())`.
//
// Special case: for zero value vectors it returns that unchanged.
func (v Vec) Normalized() Vec {
	l := v.LenSquared()
	if l != 0 {
		return v.Mulf(1 / math.Sqrt(l))
	}
	return v
}

func (v Vec) ClampLen(limit float64) Vec {
	l := v.Len()
	if l > 0 && l > limit {
		v = v.Divf(l)
		v = v.Mulf(limit)
	}
	return v
}

// Neg applies unary minus (-) to the vector.
func (v Vec) Neg() Vec {
	return Vec{
		X: -v.X,
		Y: -v.Y,
	}
}

// CubicInterpolate interpolates between a (this vector) and b using
// preA and postB as handles.
// The t arguments specifies the interpolation progression (a value from 0 to 1).
// With t=0 it returns a, with t=1 it returns b.
func (v Vec) CubicInterpolate(preA, b, postB Vec, t float64) Vec {
	res := v
	res.X = cubicInterpolate(res.X, b.X, preA.X, postB.X, t)
	res.Y = cubicInterpolate(res.Y, b.Y, preA.Y, postB.Y, t)
	return res
}

// LinearInterpolate interpolates between two points by a normalized value.
// This function is commonly named "lerp".
func (v Vec) LinearInterpolate(to Vec, t float64) Vec {
	return Vec{
		X: Lerp(v.X, to.X, t),
		Y: Lerp(v.Y, to.Y, t),
	}
}

// Midpoint returns the middle point vector of two point vectors.
//
// If we imagine [v] and [to] form a line, the midpoint would
// be a central point of this line.
func (v Vec) Midpoint(to Vec) Vec {
	return v.Add(to).Mulf(0.5)
}

// BoundsRect creates a rectangle with a center or v, width of w and height of h.
// This is useful when a vector interpreted as a point needs to be extended to an area.
//
// Note that the result is not rounded.
func (v Vec) BoundsRect(w, h float64) Rect {
	offset := Vec{
		X: w * 0.5,
		Y: h * 0.5,
	}
	return Rect{
		Min: v.Sub(offset),
		Max: v.Add(offset),
	}
}

func (v Vec) MarshalJSON() ([]byte, error) {
	if v.IsZero() {
		// Zero vectors are quite common.
		// Encode them with a shorter notation.
		return []byte("[]"), nil
	}
	buf := make([]byte, 0, 16)
	buf = append(buf, '[')
	buf = strconv.AppendFloat(buf, v.X, 'f', -1, 64)
	buf = append(buf, ',')
	buf = strconv.AppendFloat(buf, v.Y, 'f', -1, 64)
	buf = append(buf, ']')
	return buf, nil
}

func (v *Vec) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		// Recognize a MarshalJSON-produced empty vector notation.
		*v = Vec{}
		return nil
	}

	if data[0] != '[' {
		return errors.New("missing opening '['")
	}
	if data[len(data)-1] != ']' {
		return errors.New("missing closing ']'")
	}

	data = data[1:]           // '['
	data = data[:len(data)-1] // ']'

	commaIndex := bytes.IndexByte(data, ',')
	if commaIndex == -1 {
		return errors.New("missing ',' between X and Y values")
	}
	x, err := parseFloat(data[:commaIndex])
	if err != nil {
		return err
	}
	y, err := parseFloat(data[commaIndex+1:])
	if err != nil {
		return err
	}

	v.X = x
	v.Y = y
	return err
}
