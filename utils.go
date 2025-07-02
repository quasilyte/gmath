package gmath

import (
	"bytes"
	"strconv"
)

// IntPercentages attemtps to calculate the percentages among the values
// while converting them to a human-readable int values (e.g. 100 is 100%, 50 is 50%).
//
// This function guarantees the sum of all percentages to be 100, but it does not
// promise that same values will get exactly the same percentages (as rounding
// and remainder distribution can be necessary).
//
// It is mostly useful when presenting the numerical statistics to the user.
// The user can tolerate the percentages to be approximated, but they will
// notice that the sum of all percentages is not 100.
//
// The negative values will confuse this function.
//
// As a special case, if all values are 0, a slice of all-zeroes is returned.
func IntPercentages(values []float64) []int {
	switch len(values) {
	case 0:
		return nil
	case 1:
		return []int{100}
	}

	total := 0.0
	for _, v := range values {
		total += v
	}

	result := make([]int, len(values))

	if total == 0 {
		// A special case - we can't divide by 0.
		return result
	}

	percentSum := 0
	for i, v := range values {
		p := int(100 * (v / total))
		percentSum += p
		result[i] += p
	}

	for percentSum < 100 {
		for i := len(result) - 1; i >= 0; i-- {
			if values[i] == 0 {
				continue
			}
			percentSum++
			result[i]++
			if percentSum >= 100 {
				break
			}
		}
	}

	return result
}

func parseFloat(s []byte) (float64, error) {
	s = bytes.TrimSpace(s)

	sign := false
	if s[0] == '-' {
		sign = true
		s = s[1:]
	}
	f, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		return 0, err
	}
	if sign {
		return -f, nil
	}
	return f, nil
}
