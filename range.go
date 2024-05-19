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

func (rng Range[T]) Len() int {
	return int(rng.Max - rng.Min)
}
