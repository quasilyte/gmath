package gmath

const (
	pcgMultiplier = 6364136223846793005
)

// RandSource is a PCG-64 implementation of a rand source.
//
// Its benefits are:
// * efficient storage (2 uints instead of tables)
// * can save/load the source state (just 1 uint64 for a state)
// * fast seeding/reseeding
//
// A zero value of this rand source is not enough.
// Call [Seed] before using it or load the existing state.
type RandSource struct {
	state uint64
	inc   uint64
}

func (s *RandSource) Seed(v int64) {
	s.setSeed(uint64(v))
}

func (s *RandSource) Int63() int64 {
	// Rand v1 requires Int63 to return a non-negative
	// int64-typed value, hence the mask hacks.

	const (
		rngMax  = 1 << 63
		rngMask = rngMax - 1
	)

	v := s.Uint64()
	return int64(v & rngMask)
}

func (s *RandSource) Uint64() uint64 {
	state := s.state
	s.step()
	return s.rand64(state)
}

func (s *RandSource) setSeed(seed uint64) {
	s.state = 0
	s.inc = 1442695040888963407 + (seed * 151)

	s.step()
	s.state += seed
	s.step()
}

func (s *RandSource) GetState() uint64 {
	return s.state
}

func (s *RandSource) SetState(state uint64) {
	s.state = state
}

func (s *RandSource) step() {
	s.state = s.state*pcgMultiplier + s.inc
}

func (s *RandSource) rand64(state uint64) uint64 {
	word := ((state >> ((state >> 59) + 5)) ^ state) * 12605985483714917081
	return (word >> 43) ^ word
}
