package numerical_test

import (
	"testing"

	"github.com/gpahal/go-algos/algo/numerical"
)

func TestPrimeFactors(t *testing.T) {
	cases := []struct {
		value        int
		primeFactors map[int]int
	}{
		{7, map[int]int{7: 1}},
		{52, map[int]int{2: 2, 13: 1}},
		{700, map[int]int{2: 2, 5: 2, 7: 1}},
		{1, map[int]int{}},
		{-8, nil},
	}

	for _, c := range cases {
		primeFactors := numerical.PrimeFactors(c.value)
		if !mapsEqual(primeFactors, c.primeFactors) {
			t.Errorf("PrimeFactors %d: expected PrimeFactors to be %v, got %v", c.value, c.primeFactors, primeFactors)
		}
	}
}

func mapsEqual(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	for k, v := range a {
		if otherValue, ok := b[k]; !ok || otherValue != v {
			return false
		}
	}

	return true
}
