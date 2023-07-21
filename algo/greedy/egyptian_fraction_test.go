package greedy_test

import (
	"testing"

	"github.com/gpahal/go-algos/algo/greedy"
)

func TestEgyptianFraction(t *testing.T) {
	cases := []struct {
		num int
		den int
		res []int
	}{
		{2, 3, []int{2, 6}},
		{12, 13, []int{2, 3, 12, 156}},
		{12, 28, []int{3, 11, 231}},
		{0, 2, []int{}},
		{-2, 3, nil},
		{3, 2, nil},
		{5, 0, nil},
	}

	for _, c := range cases {
		res := greedy.EgyptianFraction(c.num, c.den)
		if !slicesEquals(res, c.res) {
			t.Errorf("EgyptianFraction %d %d: expected result to be %#v, got %#v", c.num, c.den, c.res, res)
		}
	}
}

func slicesEquals(arr1, arr2 []int) bool {
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
