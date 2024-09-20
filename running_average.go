package gmath

import "math"

type RunningWeightedAverage[T numeric] struct {
	TotalWeight T
	TotalValue  T
}

func (avg *RunningWeightedAverage[T]) Add(value, weight T) {
	avg.TotalValue += value * weight
	avg.TotalWeight += weight
}

func (avg *RunningWeightedAverage[T]) Scale(v float64) {
	avg.TotalValue = T(math.Round(float64(avg.TotalValue) * v))
	avg.TotalWeight = T(math.Round(float64(avg.TotalWeight) * v))
}

func (avg *RunningWeightedAverage[T]) Value() T {
	if avg.TotalWeight == 0 {
		return 0
	}
	return avg.TotalValue / avg.TotalWeight
}
