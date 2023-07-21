package concurrency

import "sync"

// Merge multiplexes a list of channels onto a single channel that is closed when all input
// channels are closed. The ordering for a single channel is maintained but values coming from
// different channels are not ordered in any way.
//
// When the done channel is closed or a value is sent to it, we stop reading from input
// channels and return after closing the output channel. The consumer of the function can use
// done as a cancellation mechanism or a cleanup mechanism - that is close the done channel
// after you have read all the required values from the output channel.
//
// # Example
//
//     done := make(chan struct{})
//     defer close(done)
//
//     c1 := make(chan int)
//     c2 := make(chan int)
//
//     go func() {
//         c1 <- 1
//         c2 <- 2
//     }
//
//     for n := range concurrency.Merge(done, c2, c2) {
//         fmt.Println(n) // prints 1 and 2, or 2 and 1
//     }
func Merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)

	// A WaitGroup is used here to wait for len(cs) goroutines to terminate. These goroutines are
	// started in the for the for statement below for each input channel.
	var wg sync.WaitGroup
	wg.Add(len(cs))

	// Start a goroutine for each input channel. It copies values from its input channel to the out
	// channel until the input channel is closed. Then the goroutine calls wg.Done to indicate it
	// is about to terminate.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close the out channel once all the output goroutines are done.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
