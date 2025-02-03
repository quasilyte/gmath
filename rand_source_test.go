package gmath

import (
	"math/rand"
	"reflect"
	"testing"
)

var (
	_ rand.Source   = (*RandSource)(nil)
	_ rand.Source64 = (*RandSource)(nil)
)

func TestRandSourceReload(t *testing.T) {
	const (
		seed     = 435
		numSkips = 5
	)

	var rng RandSource
	rng.Seed(seed)

	for i := 0; i < numSkips; i++ {
		rng.Uint64()
	}
	state := rng.GetState()

	values := make([]uint64, 10)
	for i := range values {
		values[i] = rng.Uint64()
	}

	var rng2 RandSource
	rng2.Seed(seed)
	rng2.SetState(state)
	values2 := make([]uint64, 10)
	for i := range values2 {
		values2[i] = rng2.Uint64()
	}

	for i := range values {
		v1 := values[i]
		v2 := values2[i]
		if v1 != v2 {
			t.Fatalf("[i=%d] value mismatch:\nhave: %d\nwant: %d", i, v2, v1)
		}
	}

	rng.Seed(seed)
	for i := 0; i < numSkips; i++ {
		rng.Uint64()
	}
	values3 := make([]uint64, 10)
	for i := range values2 {
		values3[i] = rng.Uint64()
	}

	if !reflect.DeepEqual(values, values3) {
		t.Fatalf("second round failed")
	}
}

func TestRandSourceQuality(t *testing.T) {

	seeds := []int64{
		0,
		1,
		548238,
		19,
		902395988182,
	}

	for _, seed := range seeds {
		var rng RandSource
		rng.Seed(seed)
		valueSet := make(map[int64]struct{}, 10000)
		for i := 0; i < 10000; i++ {
			v := rng.Int63()
			if _, ok := valueSet[v]; ok {
				t.Fatalf("seed=%d i=%d has repeated value", seed, i)
			}
			valueSet[v] = struct{}{}
		}
	}
}
