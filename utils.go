package gmath

import (
	"bytes"
	"strconv"
)

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
