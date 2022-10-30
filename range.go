package gmath

type Range[T numeric] struct {
	Min T
	Max T
}

func MakeRange[T numeric](min, max T) Range[T] {
	return Range[T]{Min: min, Max: max}
}
