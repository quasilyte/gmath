package gmath

import (
	"testing"
)

func BenchmarkRandPicker(b *testing.B) {
	var rng Rand
	rng.SetSeed(2493)
	picker := NewRandPicker[float64](&rng)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			picker.Reset()
			picker.AddOption(132, 0.4)
			picker.AddOption(1320, 0.2)
			picker.AddOption(13200, 0.1)
			picker.AddOption(132000, 0.5)
			picker.AddOption(1320000, 0.001)
			picker.Pick()
		}
	}
}
