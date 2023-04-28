package gmath

import (
	"sort"
)

// RandPicker performs a uniformly distributed random probing among the given objects with weights.
// Higher the weight, higher the chance of that object of being picked.
type RandPicker[T any] struct {
	r *Rand

	keys   randPickerKeySlice
	values []T

	threshold float64
	sorted    bool
}

type randPickerKey struct {
	index     int
	threshold float64
}

type randPickerKeySlice []randPickerKey

func (s *randPickerKeySlice) Len() int           { return len(*s) }
func (s *randPickerKeySlice) Less(i, j int) bool { return (*s)[i].threshold < (*s)[j].threshold }
func (s *randPickerKeySlice) Swap(i, j int)      { (*s)[i], (*s)[j] = (*s)[j], (*s)[i] }

func NewRandPicker[T any](r *Rand) *RandPicker[T] {
	return &RandPicker[T]{r: r}
}

func (p *RandPicker[T]) Reset() {
	p.keys = p.keys[:0]
	p.values = p.values[:0]
	p.threshold = 0
	p.sorted = false
}

func (p *RandPicker[T]) AddOption(value T, weight float64) {
	if weight == 0 {
		return // Zero probability in any case
	}
	p.threshold += weight
	p.keys = append(p.keys, randPickerKey{
		threshold: p.threshold,
		index:     len(p.values),
	})
	p.values = append(p.values, value)
	p.sorted = false
}

func (p *RandPicker[T]) IsEmpty() bool {
	return len(p.values) != 0
}

func (p *RandPicker[T]) Pick() T {
	var result T
	if len(p.values) == 0 {
		return result // Zero value
	}
	if len(p.values) == 1 {
		return p.values[0]
	}

	// In a normal use case the random picker is initialized and then used
	// without adding extra options, so this sorting will happen only once in that case.
	if !p.sorted {
		sort.Sort(&p.keys)
		p.sorted = true
	}

	roll := p.r.FloatRange(0, p.threshold)
	i := sort.Search(len(p.keys), func(i int) bool {
		return roll <= p.keys[i].threshold
	})
	if i < len(p.keys) && roll <= p.keys[i].threshold {
		result = p.values[p.keys[i].index]
	} else {
		result = p.values[len(p.values)-1]
	}
	return result
}
