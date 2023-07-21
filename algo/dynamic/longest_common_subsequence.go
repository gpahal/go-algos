package dynamic

// LongestCommonSubsequence takes two slices of integers and returns the longest common subsequence
// as a slice. A subsequence is a sequence that appears in the same relative order, but not
// necessarily contiguous.
//
// # Example
//
//     lcs := dynamic.LongestCommonSubsequence([]int{1, 5, 8, 9, 3}, []int{5, 8, 4, 9, 2})
//     // lcs = []int{5, 8, 9}
//
// If arr1 or arr2 is nil, nil is returned.
func LongestCommonSubsequence(arr1, arr2 []int) []int {
	if arr1 == nil || arr2 == nil {
		return nil
	}

	arrLen1 := len(arr1)
	arrLen2 := len(arr2)
	if len(arr1) == 0 || len(arr2) == 0 {
		return []int{}
	}

	m := make(map[Pair]int)
	computeLengths(m, arr1, arr2, arrLen1-1, arrLen2-1)
	return computeSubsequence(m, arr1, arr2, arrLen1-1, arrLen2-1)
}

// computeLengths computes the length of the longest common subsequence for arr1[:i+1] and
// arr2[:j+1]. It uses memoization with the help of the map m.
func computeLengths(m map[Pair]int, arr1, arr2 []int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if val, ok := m[Pair{i, j}]; ok {
		return val
	}

	var val int
	if arr1[i] == arr2[j] {
		val = 1 + computeLengths(m, arr1, arr2, i-1, j-1)
	} else {
		val = max(computeLengths(m, arr1, arr2, i-1, j), computeLengths(m, arr1, arr2, i, j-1))
	}

	m[Pair{i, j}] = val
	return val
}

// computeSubsequence computes the longest common subsequence for arr1[:i+1] and arr2[:j+1] using
// the length mappings m.
func computeSubsequence(m map[Pair]int, arr1, arr2 []int, i, j int) []int {
	lcsLen := m[Pair{i, j}]
	lcs := make([]int, lcsLen)
	currIdx := lcsLen - 1

	for currIdx >= 0 {
		if arr1[i] == arr2[j] {
			lcs[currIdx] = arr1[i]
			currIdx--
			i--
			j--
		} else {
			if j == 0 || (m[Pair{i - 1, j}] >= m[Pair{i, j - 1}]) {
				i--
			} else {
				j--
			}
		}
	}

	return lcs
}

// Pair is a pair of int values.
type Pair struct {
	First  int
	Second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
