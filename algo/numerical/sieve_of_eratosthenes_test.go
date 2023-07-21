package numerical_test

import (
	"testing"

	"github.com/gpahal/go-algos/algo/numerical"
)

func TestSieveOfEratosthenes(t *testing.T) {
	cases := []struct {
		limit int
		sieve []bool
	}{
		{1, []bool{true, true}},
		{4, []bool{true, true, false, false, true}},
		{30, []bool{true, true, false, false, true, false, true, false, true, true, true, false, true, false, true, true, true, false, true, false, true, true, true, false, true, true, true, true, true, false, true}},
		{0, nil},
		{1 << 28, nil},
	}

	for _, c := range cases {
		sieve := numerical.SieveOfEratosthenes(c.limit)
		if !slicesEqual(sieve, c.sieve) {
			t.Errorf("SieveOfEratosthenes %d: expected sieve to be %v, got %v", c.limit, c.sieve, sieve)
		}
	}
}

func TestSieveOfEratosthenesDynamic(t *testing.T) {
	cases := []struct {
		n     int
		prime int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 7},
		{25, 97},
	}

	for _, c := range cases {
		f := numerical.SieveOfEratosthenesDynamic()
		var prime int
		for i := 0; i < c.n; i++ {
			prime = f()
		}

		if prime != c.prime {
			t.Errorf("SieveOfEratosthenesDynamic (n = %d): expected nth prime to be %d, got %d", c.n, c.prime, prime)
		}
	}
}

func slicesEqual(arr1, arr2 []bool) bool {
	if arr1 == nil && arr2 == nil {
		return true
	}
	if arr1 == nil || arr2 == nil {
		return false
	}
	if len(arr1) != len(arr2) {
		return false
	}

	for i, item := range arr1 {
		if arr2[i] != item {
			return false
		}
	}

	return true
}
