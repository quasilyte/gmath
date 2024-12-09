package gmath

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"math"
	"strconv"
	"unsafe"
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
//
// If you need float32 components, use [Vec32] type,
// but keep in mind that [Vec] should be preferred most of the time.
type Vec = vec[float64]

// Vec32 is like [Vec], but with float32-typed fields.
// You should generally prefer [Vec], but for some specific low-level
// stuff you might want to use a float32 variant.
//
// Most functions of this package operate on [Vec], so you will
// lose some of the convenience while using [Vec32].
//
// If anything, [Vec32] will be slower on most operations due to the
// intermediate float32->float64 conversions that will happen here and there.
// It should be only used as a space optimization, when you need
// to store lots of vectors and memory locality dominates a small processing overhead.
// If in doubts, use [Vec].
type Vec32 = vec[float32]

type vec[T float] struct {
	X T
	Y T
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
func (v vec[T]) ToStd() image.Point {
	return image.Point{
		X: int(v.X),
		Y: int(v.Y),
	}
}

// String returns a pretty-printed representation of a 2D vector object.
func (v vec[T]) String() string {
	return fmt.Sprintf("[%f, %f]", v.X, v.Y)
}

// IsZero reports whether v is a zero value vector.
// A zero value vector has X=0 and Y=0, created with Vec{}.
//
// The zero value vector has a property that its length is 0,
// but not all zero length vectors are zero value vectors.
func (v vec[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// IsNormalizer reports whether the vector is normalized.
// A vector is considered to be normalized if its length is 1.
func (v vec[T]) IsNormalized() bool {
	return EqualApprox(v.LenSquared(), 1)
}

// DistanceTo calculates the distance between the two vectors.
func (v vec[T]) DistanceTo(v2 vec[T]) T {
	return T(math.Sqrt(float64((v.X-v2.X)*(v.X-v2.X) + (v.Y-v2.Y)*(v.Y-v2.Y))))
}

func (v vec[T]) DistanceSquaredTo(v2 vec[T]) T {
	return ((v.X - v2.X) * (v.X - v2.X)) + ((v.Y - v2.Y) * (v.Y - v2.Y))
}

// Dot returns a dot-product of the two vectors.
func (v vec[T]) Dot(v2 vec[T]) T {
	return (v.X * v2.X) + (v.Y * v2.Y)
}

// Len reports the length of this vector (also known as magnitude).
func (v vec[T]) Len() T {
	return T(math.Sqrt(float64(v.LenSquared())))
}

// LenSquared returns the squared length of this vector.
//
// This function runs faster than Len(),
// so prefer it if you need to compare vectors
// or need the squared distance for some formula.
func (v vec[T]) LenSquared() T {
	return v.Dot(v)
}

func (v vec[T]) Rotated(angle Rad) vec[T] {
	sine, cosi := math.Sincos(float64(angle))
	// For 64-bit it should be a no-op recognizable by the compiler.
	tsin := T(sine)
	tcos := T(cosi)
	return vec[T]{
		X: v.X*tcos - v.Y*tsin,
		Y: v.X*tsin + v.Y*tcos,
	}
}

func (v vec[T]) Angle() Rad {
	return Rad(math.Atan2(float64(v.Y), float64(v.X)))
}

// AngleToPoint returns the angle from v towards the given point.
func (v vec[T]) AngleToPoint(pos vec[T]) Rad {
	return pos.Sub(v).Angle()
}

func (v vec[T]) DirectionTo(v2 vec[T]) vec[T] {
	return v.Sub(v2).Normalized()
}

func (v vec[T]) VecTowards(pos vec[T], length T) vec[T] {
	angle := v.AngleToPoint(pos)
	result := vec[T]{X: T(angle.Cos()), Y: T(angle.Sin())}
	return result.Mulf(length)
}

func (v vec[T]) MoveTowards(pos vec[T], length T) vec[T] {
	direction := pos.Sub(v) // Not normalized
	dist := direction.Len()
	if dist <= length || dist < Epsilon {
		return pos
	}
	return v.Add(direction.Divf(dist).Mulf(length))
}

func (v vec[T]) EqualApprox(other vec[T]) bool {
	return EqualApprox(v.X, other.X) && EqualApprox(v.Y, other.Y)
}

func (v vec[T]) MoveInDirection(dist T, dir Rad) vec[T] {
	return vec[T]{
		X: v.X + T(float64(dist)*dir.Cos()),
		Y: v.Y + T(float64(dist)*dir.Sin()),
	}
}

func (v vec[T]) Mulf(scalar T) vec[T] {
	return vec[T]{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v vec[T]) Mul(other vec[T]) vec[T] {
	return vec[T]{
		X: v.X * other.X,
		Y: v.Y * other.Y,
	}
}

func (v vec[T]) Divf(scalar T) vec[T] {
	return vec[T]{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v vec[T]) Div(other vec[T]) vec[T] {
	return vec[T]{
		X: v.X / other.X,
		Y: v.Y / other.Y,
	}
}

func (v vec[T]) Add(other vec[T]) vec[T] {
	return vec[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v vec[T]) Sub(other vec[T]) vec[T] {
	return vec[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v vec[T]) Rounded() vec[T] {
	return vec[T]{
		X: T(math.Round(float64(v.X))),
		Y: T(math.Round(float64(v.Y))),
	}
}

func (v vec[T]) Floored() vec[T] {
	return vec[T]{
		X: T(math.Floor(float64(v.X))),
		Y: T(math.Floor(float64(v.Y))),
	}
}

func (v vec[T]) Ceiled() vec[T] {
	return vec[T]{
		X: T(math.Ceil(float64(v.X))),
		Y: T(math.Ceil(float64(v.Y))),
	}
}

// Normalized returns the vector scaled to unit length.
// Functionally equivalent to `v.Divf(v.Len())`.
//
// Special case: for zero value vectors it returns that unchanged.
func (v vec[T]) Normalized() vec[T] {
	l := v.LenSquared()
	if l != 0 {
		return v.Mulf(T(1 / math.Sqrt(float64(l))))
	}
	return v
}

func (v vec[T]) ClampLen(limit T) vec[T] {
	l := v.Len()
	if l > 0 && l > limit {
		v = v.Divf(l)
		v = v.Mulf(limit)
	}
	return v
}

// Neg applies unary minus (-) to the vector.
func (v vec[T]) Neg() vec[T] {
	return vec[T]{
		X: -v.X,
		Y: -v.Y,
	}
}

// CubicInterpolate interpolates between a (this vector) and b using
// preA and postB as handles.
// The t arguments specifies the interpolation progression (a value from 0 to 1).
// With t=0 it returns a, with t=1 it returns b.
func (v vec[T]) CubicInterpolate(preA, b, postB Vec, t T) vec[T] {
	res := v
	res.X = T(cubicInterpolate(float64(res.X), float64(b.X), preA.X, postB.X, float64(t)))
	res.Y = T(cubicInterpolate(float64(res.Y), float64(b.Y), preA.Y, postB.Y, float64(t)))
	return res
}

// LinearInterpolate interpolates between two points by a normalized value.
// This function is commonly named "lerp".
func (v vec[T]) LinearInterpolate(to vec[T], t T) vec[T] {
	return vec[T]{
		X: T(Lerp(float64(v.X), float64(to.X), float64(t))),
		Y: T(Lerp(float64(v.Y), float64(to.Y), float64(t))),
	}
}

// Midpoint returns the middle point vector of two point vectors.
//
// If we imagine [v] and [to] form a line, the midpoint would
// be a central point of this line.
func (v vec[T]) Midpoint(to vec[T]) vec[T] {
	return v.Add(to).Mulf(0.5)
}

// BoundsRect creates a rectangle with a center or v, width of w and height of h.
// This is useful when a vector interpreted as a point needs to be extended to an area.
//
// Note that the result is not rounded.
func (v vec[T]) BoundsRect(w, h T) Rect {
	offset := vec[T]{
		X: w * 0.5,
		Y: h * 0.5,
	}
	return Rect{
		Min: v.Sub(offset).AsVec64(),
		Max: v.Add(offset).AsVec64(),
	}
}

func (v vec[T]) MarshalJSON() ([]byte, error) {
	if v.IsZero() {
		// Zero vectors are quite common.
		// Encode them with a shorter notation.
		return []byte("[]"), nil
	}
	buf := make([]byte, 0, 16)
	buf = append(buf, '[')
	buf = strconv.AppendFloat(buf, float64(v.X), 'f', -1, 64)
	buf = append(buf, ',')
	buf = strconv.AppendFloat(buf, float64(v.Y), 'f', -1, 64)
	buf = append(buf, ']')
	return buf, nil
}

func (v *vec[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		// Recognize a MarshalJSON-produced empty vector notation.
		*v = vec[T]{}
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

	v.X = T(x)
	v.Y = T(y)
	return err
}

func (v vec[T]) AsVec64() Vec {
	// For vec[float64] this should be no-op.
	// For vec[float32] it should do a float32->float64 conversion.
	return Vec{
		X: float64(v.X),
		Y: float64(v.Y),
	}
}

func (v vec[T]) AsVec32() Vec32 {
	// For vec[float32] this should be no-op.
	// For vec[float64] it should do a float64->float32 conversion.
	return Vec32{
		X: float32(v.X),
		Y: float32(v.Y),
	}
}

// AsSlice returns vector as a slice view.
//
// This view can be used to read and write to Vec,
// but it should not be used as append operand.
//
// This operation doesn't allocate.
//
// For 64-bit vectors, it returns []float64,
// For 32-bit vectors, it returns []float32.
//
// The most common use case for this function is
// uniform variables binding in Ebitengine, as
// it wants vec's as []float32 slices.
// Allocating real slices for it is a waste,
// therefore we can use a convenient Vec API while
// still being compatible with Ebitengine needs without
// any redundant allocations.
func (v *vec[T]) AsSlice() []T {
	return unsafe.Slice(&v.X, 2)
}
