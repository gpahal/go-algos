package searching_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gpahal/go-algos/algo/searching"
	"github.com/gpahal/go-algos/algo/sorting"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func assertSearchFn(t *testing.T, name string, fn func([]int, int) int, sort bool) bool {
	t.Helper()

	for i := 0; i < 10; i += 1 {
		length := (rand.Int() % 20) + 1
		arr := generateRandSlice(length, false)
		if sort {
			sorting.QuickSort(arr)
		}

		var key int
		if i == 5 {
			key = -1
		} else if i%2 == 0 {
			key = arr[rand.Int()%length]
		} else {
			key = (rand.Int() % 50) + 1
		}

		idx := fn(arr, key)
		correct, correctIdx := checkSearchResult(arr, key, idx)
		if !correct {
			t.Errorf("%s %v for %d: expected %d, got %d", name, arr, key, correctIdx, idx)
			return false
		}
	}

	return true
}

func checkSearchResult(arr []int, key, idx int) (bool, int) {
	if idx < 0 || len(arr) <= idx || arr[idx] != key {
		for i, el := range arr {
			if el == key {
				return false, i
			}
		}

		return idx < 0, -1
	}

	return true, idx
}

func TestLinearSearch(t *testing.T) {
	assertSearchFn(t, "LinearSearch", searching.LinearSearch, false)
}

func TestBinarySearch(t *testing.T) {
	assertSearchFn(t, "BinarySearch", searching.BinarySearch, true)
}

func benchmarkSearchFn(b *testing.B, fn func([]int, int) int, sort bool, length int) {
	arr := generateRandSlice(length, true)
	if sort {
		sorting.QuickSort(arr)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(arr, (rand.Int()%50)+1)
	}
}

func BenchmarkLinearSearch_10000(b *testing.B) {
	benchmarkSearchFn(b, searching.LinearSearch, false, 10000)
}

func BenchmarkBinarySearch_10000(b *testing.B) {
	benchmarkSearchFn(b, searching.BinarySearch, true, 10000)
}

func generateRandSlice(length int, biggerInts bool) []int {
	if length <= 0 {
		return []int{}
	}
	maxNum := 50
	if biggerInts {
		maxNum = 500000
	}

	arr := make([]int, length)
	for i := 0; i < length; i += 1 {
		arr[i] = (rand.Int() % maxNum) + 1 // random number in the range 1..50
	}

	return arr
}
