package concurrency_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gpahal/go-algos/pattern/concurrency"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestTake(t *testing.T) {
	counts := [][2]int{{0, 0}, {0, 2}, {2, 0}, {2, 4}, {4, 2}}
	for _, count := range counts {
		testSingleTake(t, count)
	}
}

func testSingleTake(t *testing.T, count [2]int) {
	t.Helper()

	block := make(chan struct{})
	done := make(chan struct{})
	defer close(done)

	inArr := make([]int, count[0])
	c := make(chan int)
	go func() {
		for i := 0; i < count[0]; i++ {
			n := rand.Intn(100)
			c <- n
			inArr[i] = n
		}
		close(c)
		block <- struct{}{}
	}()

	out := concurrency.Take(done, c, count[1])
	outArr := []int{}
	for n := range out {
		outArr = append(outArr, n)
	}

	// Make sure the goroutine is cleaned up.
	for _ = range c {
	}
	<-block

	expLen := count[1]
	if count[0] < count[1] {
		expLen = count[0]
	}

	if len(outArr) != expLen {
		t.Fatalf(
			"Take %s: expected %d values in output, got %d: (chan[%v])",
			getArgsTake(inArr, count[1]), expLen, len(outArr), outArr)
	}
	if expLen == 0 {
		return
	}

	for i := 0; i < expLen; i++ {
		if inArr[i] != outArr[i] {
			t.Fatalf(
				"Take %s: expected %d at index %d, got %d", getArgsTake(inArr, count[1]),
				inArr[i], i, outArr[i])
		}
	}
}

func getArgsTake(inArr []int, count int) string {
	return fmt.Sprintf("chan[%v], %d", inArr, count)
}
