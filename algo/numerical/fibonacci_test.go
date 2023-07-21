package numerical_test

import (
	"testing"

	"github.com/gpahal/go-algos/algo/numerical"
)

func TestFibonacci(t *testing.T) {
	cases := [][]int{
		{0, 0},
		{1, 1},
		{6, 8},
		{8, 21},
		{20, numerical.Fibonacci(18) + numerical.Fibonacci(19)},
		{-5, 0},
	}

	for _, c := range cases {
		value := numerical.Fibonacci(c[0])
		if value != c[1] {
			t.Errorf("Fibonacci %d: expected Fibonacci to be %d, got %d", c[0], c[1], value)
		}
	}
}
