package greedy_test

import (
	"sort"
	"testing"

	"github.com/gpahal/go-algos/algo/greedy"
)

func TestActivitySelection(t *testing.T) {
	cases := []struct {
		activities [][2]int
		resCount   int
	}{
		{[][2]int{}, 0},
		{[][2]int{{7, 8}}, 1},
		{[][2]int{{7, 12}, {8, 14}, {9, 15}}, 1},
		{[][2]int{{10, 20}, {12, 25}, {20, 30}}, 2},
		{[][2]int{{1, 2}, {3, 4}, {0, 6}, {5, 7}, {8, 9}, {5, 9}}, 4},
		{[][2]int{{1, 2}, {3, 4}, {-1, 6}}, 0},
	}

	for _, c := range cases {
		selection := greedy.ActivitySelection(c.activities, true)
		if !validActivitySelection(c.activities, selection) {
			t.Errorf("ActivitySelection %v: expected selection to be valid, got %v", c.activities, selection)
		}
		if len(selection) != c.resCount {
			t.Errorf("ActivitySelection %v: expected selection count to be %d, got %d", c.activities, c.resCount, len(selection))
		}
	}
}

func validActivitySelection(activities, selection [][2]int) bool {
	if len(selection) == 0 {
		return true
	}
	if len(activities) < len(selection) || !sliceContains(activities, selection...) {
		return false
	}

	sort.Slice(selection, func(i, j int) bool {
		if selection[i][0] <= selection[j][0] {
			return true
		}

		return false
	})

	activity := selection[0]
	for i := 1; i < len(selection); i++ {
		if selection[i][0] < activity[1] {
			return false
		}
	}

	return true
}

func sliceContains(arr [][2]int, items ...[2]int) bool {
outer:
	for _, item := range items {
		for _, el := range arr {
			if el[0] == item[0] && el[1] == item[1] {
				continue outer
			}
		}

		return false
	}

	return true
}
