package numerical_test

import (
	"testing"

	"github.com/gpahal/go-algos/algo/numerical"
)

func TestLCM(t *testing.T) {
	cases := [][3]int{
		{52, 78, 156},
		{8, -20, 40},
		{1, 3, 3},
		{40, 40, 40},
		{0, 10, 0},
	}

	for _, c := range cases {
		lcm := numerical.LCM(c[0], c[1])
		if lcm != c[2] {
			t.Errorf("LCM %d, %d: expected LCM to be %d, got %d", c[0], c[1], c[2], lcm)
		}
	}
}

func TestLCMArray(t *testing.T) {
	cases := []struct {
		arr []int
		lcm int
	}{
		{[]int{52, 78, 26}, 156},
		{[]int{4, 12, 16, -4, 8, 10}, 240},
		{[]int{1, 2, 3}, 6},
		{[]int{40, 40, 40, 40}, 40},
		{[]int{7}, 7},
		{nil, 0},
		{[]int{}, 0},
	}

	for _, c := range cases {
		lcm := numerical.LCMArray(c.arr)
		if lcm != c.lcm {
			t.Errorf("LCMArray %#v: expected LCMArray to be %d, got %d", c.arr, c.lcm, lcm)
		}
	}
}
