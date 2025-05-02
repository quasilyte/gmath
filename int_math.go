package gmath

func Iabs[T integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
