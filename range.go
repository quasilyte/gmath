package gmath

type Range[T numeric] struct {
	Min T
	Max T
}

func MakeRange[T numeric](min, max T) Range[T] {
	return Range[T]{Min: min, Max: max}
}

func (rng Range[T]) IsZero() bool {
	return rng == Range[T]{}
}

func (rng Range[T]) IsValid() bool {
	return rng.Max >= rng.Min
}

func (rng Range[T]) Contains(v T) bool {
	return v >= rng.Min && v <= rng.Max
}

func (rng Range[T]) InBounds(minValue, maxValue T) bool {
	return rng.Min >= minValue && rng.Max <= maxValue &&
		rng.IsValid()
}

func (rng Range[T]) Len() int {
	return int(rng.Max - rng.Min)
}
