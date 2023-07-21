package concurrency_test

import (
	"sync"
	"testing"
	"time"

	"github.com/gpahal/go-algos/pattern/concurrency"
)

func TestWorkerPool(t *testing.T) {
	var mu sync.Mutex
	x := 0
	addOne := func() {
		time.Sleep(10 * time.Microsecond)
		mu.Lock()
		x += 1
		mu.Unlock()
	}

	poolSize := 5
	totalTasks := 100
	tasks := make(chan concurrency.Task)
	go func() {
		for i := 0; i < totalTasks; i++ {
			tasks <- addOne
		}
		close(tasks)
	}()

	wp := concurrency.NewWorkerPool(poolSize)
	wp.Start(tasks)
	wp.Wait()

	if x != totalTasks {
		t.Errorf("Wait: expected x to be %d; got %d", totalTasks, x)
	}
}
