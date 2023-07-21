package dynamic_test

import (
	"testing"

	"github.com/gpahal/go-algos/algo/dynamic"
)

func TestLongestCommonSubsequence(t *testing.T) {
	cases := []struct {
		arr1 []int
		arr2 []int
		lcs  []int
	}{
		{[]int{1, 5, 8, 9}, []int{1, 5, 8, 9}, []int{1, 5, 8, 9}},
		{[]int{1, 5, 8, 9, 3}, []int{5, 8, 4, 9, 2}, []int{5, 8, 9}},
		{[]int{1, 2, 3, 4, 7, 8}, []int{1, 5, 4, 6, 8, 20}, []int{1, 4, 8}},
		{[]int{4, 3, 9, 8, 9, 8, 7, 8, 2, 1}, []int{4, 9, 9, 8, 8, 8, 1}, []int{4, 9, 8, 8, 8, 1}},
		{nil, []int{1, 5, 8, 9}, nil},
		{[]int{1, 5, 8, 9}, nil, nil},
	}

	for _, c := range cases {
		lcs := dynamic.LongestCommonSubsequence(c.arr1, c.arr2)
		if lcs == nil || c.lcs == nil {
			if lcs != nil || c.lcs != nil {
				t.Errorf("LongestCommonSubsequence %#v %#v: expected lcs to be %#v, got %#v", c.arr1, c.arr2, c.lcs, lcs)
			}
		} else if len(lcs) != len(c.lcs) {
			t.Errorf("LongestCommonSubsequence %#v %#v: expected lcs %#v to be of length %d, got %d", c.arr1, c.arr2, lcs, len(c.lcs), len(lcs))
		} else if !sliceSubsequence(c.arr1, lcs) {
			t.Errorf("LongestCommonSubsequence %#v %#v: expected lcs %#v to be a subsequence of %#v", c.arr1, c.arr2, lcs, c.arr1)
		} else if !sliceSubsequence(c.arr2, lcs) {
			t.Errorf("LongestCommonSubsequence %#v %#v: expected lcs %#v to be a subsequence of %#v", c.arr1, c.arr2, lcs, c.arr2)
		}
	}
}

func sliceSubsequence(arr, seq []int) bool {
	if arr == nil && seq == nil {
		return true
	}
	if arr == nil || seq == nil {
		return false
	}

	seqLen := len(seq)
	if seqLen == 0 {
		return true
	}

	currIdx := 0
	for _, item := range arr {
		if item == seq[currIdx] {
			currIdx++
			if currIdx == seqLen {
				return true
			}
		}
	}

	return false
}
