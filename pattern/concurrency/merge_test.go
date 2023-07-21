package concurrency_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/gpahal/go-algos/pattern/concurrency"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestMerge(t *testing.T) {
	counts := [][]int{[]int{}, []int{0}, []int{0, 1}, []int{1, 0}, []int{3, 3, 3}}
	for _, count := range counts {
		testSingleMerge(t, count)
	}
}

func testSingleMerge(t *testing.T, count []int) {
	t.Helper()

	var wg sync.WaitGroup
	wg.Add(len(count))

	done := make(chan struct{})
	defer close(done)

	// Create input channels and store the values emitted by them for testing.
	var mu sync.Mutex
	m := make(map[int]int)
	ns := make([][]int, len(count))
	cs := make([]<-chan int, len(count))
	for i := 0; i < len(count); i++ {
		c := make(chan int)
		ns[i] = make([]int, 0, len(count))
		cs[i] = c
		go func(i int, c chan int) {
			for j := 0; j < count[i]; j++ {
				n := rand.Intn(100)
				c <- n
				mu.Lock()
				incCounterMerge(m, n)
				ns[i] = append(ns[i], n)
				mu.Unlock()
			}
			close(c)
			wg.Done()
		}(i, c)
	}

	out := concurrency.Merge(done, cs...)
	arr := []int{}
	for n := range out {
		arr = append(arr, n)
	}

	// Make sure any remaining goroutines are cleaned up.
	wg.Wait()

	for i, n := range arr {
		val, ok := m[n]
		if !ok || val <= 0 {
			t.Fatalf(
				"Merge %s: extra %d present at index %d in the output channel: (chan[%v])",
				getArgsMerge(ns), n, i, arr)
		}

		m[n] = val - 1
	}

	for n, val := range m {
		if val > 0 {
			t.Fatalf(
				"Merge %s: not enough %d present in the output channel: (chan[%v])",
				getArgsMerge(ns), n, arr)
		}
	}
}

func incCounterMerge(m map[int]int, n int) {
	val, ok := m[n]
	if ok {
		m[n] = val + 1
	} else {
		m[n] = 1
	}
}

func getArgsMerge(ns [][]int) string {
	if len(ns) == 0 {
		return "[no args]"
	}

	args := fmt.Sprintf("chan[%v]", ns[0])
	for i := 1; i < len(ns); i++ {
		args += fmt.Sprintf(", chan[%v]", ns[i])
	}

	return args
}
