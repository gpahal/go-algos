package numerical_test

import (
	"testing"

	"github.com/gpahal/go-algos/algo/numerical"
)

func TestGCD(t *testing.T) {
	cases := [][3]int{
		{52, 78, 26},
		{8, -20, 4},
		{1, 3, 1},
		{40, 40, 40},
		{0, 10, 10},
	}

	for _, c := range cases {
		gcd := numerical.GCD(c[0], c[1])
		if gcd != c[2] {
			t.Errorf("GCD %d, %d: expected GCD to be %d, got %d", c[0], c[1], c[2], gcd)
		}
	}
}

func TestGCDArray(t *testing.T) {
	cases := []struct {
		arr []int
		gcd int
	}{
		{[]int{52, 78, 26}, 26},
		{[]int{4, 12, 16, -4, 8, 10}, 2},
		{[]int{1, 2, 3}, 1},
		{[]int{40, 40, 40, 40}, 40},
		{[]int{7}, 7},
		{nil, 0},
		{[]int{}, 0},
	}

	for _, c := range cases {
		gcd := numerical.GCDArray(c.arr)
		if gcd != c.gcd {
			t.Errorf("GCDArray %#v: expected GCDArray to be %d, got %d", c.arr, c.gcd, gcd)
		}
	}
}
