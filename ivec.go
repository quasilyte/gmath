package gmath

type Ivec8 = Ivec[int8]

type Ivec[T integer] struct {
	X T
	Y T
}

// ManhattanDistanceTo finds a Manhattan distance between
// the two vectors interpreted as coordinates (2D points).
func (v Ivec[T]) ManhattanDistanceTo(other Ivec[T]) T {
	return Iabs(v.X-other.X) + Iabs(v.Y-other.Y)
}

func (v Ivec[T]) Add(other Ivec[T]) Ivec[T] {
	return Ivec[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Ivec[T]) Sub(other Ivec[T]) Ivec[T] {
	return Ivec[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}
