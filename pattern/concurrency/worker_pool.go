package concurrency

import "fmt"

// Task is the single unit of work that a WorkerPool can do.
type Task func()

// workerPool provides an implementation of WorkerPool usng a semaphore to control concurrency.
// This particular implementation spawns a new goroutine for every task but limits the concurrency
// and makes sure there are no idle goroutines.
// Inspired by: [GopherCon 2018: Bryan C. Mills - Rethinking Classical Concurrency Patterns](https://www.youtube.com/watch?v=5zXAHh5tJqQ)
type workerPool struct {
	limit  int
	sem    chan struct{}
	finish chan struct{}
	done   chan struct{}
}

// WorkerPool provides basic methods to start and stop a worker pool implemetaion.
type WorkerPool interface {
	// Start starts the worker pool and provides a channel which provides the worker goroutines
	// with tasks.
	Start(tasks <-chan Task)

	// Wait waits for the worker pool to finish all the tasks. It is a blocking call which waits
	// for the task channel to close and all the workers to finish their pending tasks.
	Wait()

	// Stop stops the worker pool immediately and stops accepting any new tasks. It is a blocking
	// call which waits for all the workers to finish their pending tasks.
	Stop()
}

// NewWorkerPool returns a new WorkerPool with limited concurrency.
func NewWorkerPool(limit int) WorkerPool {
	// Check if limit is valid ie. > 0
	if limit <= 0 {
		panic(fmt.Sprintf("limit should be > 0, got %d", limit))
	}

	return &workerPool{
		limit:  limit,
		sem:    make(chan struct{}, limit),
		finish: make(chan struct{}),
		done:   make(chan struct{}),
	}
}

// Start starts the worker pool and uses the tasks channel to provide the worker goroutines with
// tasks.
func (w *workerPool) Start(tasks <-chan Task) {
	go func() {
		defer close(w.finish)
		for {
			// Check if Stop has been called. This is required to break the loop if the tasks
			// channel has many other tasks and Stop has also been called. Because select doesn't
			// prioritize, we might continuously receive from the tasks channel even though stop
			// has been called.
			if w.stopped() {
				return
			}

			select {
			// Check if Stop has been called.
			case <-w.done:
				return
			// Try to aquire a task.
			case task, ok := <-tasks:
				if !ok || !w.perform(task) {
					return
				}
			}
		}
	}()
}

// perform starts a single task in a goroutine. It can return early if Stop is called before this
// method is able to aquire from the semaphore. If a goroutine is started successfully, it returns
// true.
func (w *workerPool) perform(task Task) bool {
	// Check if Stop has been called but still a task got picked by the select statement.
	if w.stopped() {
		return false
	}

	select {
	// Check if Stop has been called.
	case <-w.done:
		return false
	// Try to aquire from a semaphore. This blocks if limit goroutines are already running and
	// unblocks when one of them completes, releasing from the semaphore.
	case w.sem <- struct{}{}:
		// Start a goroutine which performs the task and releases from the semaphore.
		go func() {
			task()
			<-w.sem
		}()
		return true
	}
}

// Wait waits for the worker pool to finish all the tasks. It is a blocking call which waits for
// the task channel to close and all the workers to finish their pending tasks.
func (w *workerPool) Wait() {
	<-w.finish

	// Try to aquire from the semaphore w.limit times to ensure all goroutines have finished and no
	// more get started.
	for i := 0; i < w.limit; i++ {
		w.sem <- struct{}{}
	}
}

// Stop stops the worker pool immediately and stops accepting any new tasks. It is a blocking call
// which waits for all the workers to finish their pending tasks.
func (w *workerPool) Stop() {
	close(w.done)
	w.Wait()
}

// stopped checks if Stop has been called.
func (w *workerPool) stopped() bool {
	select {
	case <-w.done:
		return true
	default:
		return false
	}
}
