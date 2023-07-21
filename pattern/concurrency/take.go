package concurrency

// Take copies the first count values (if present) of the input channel onto an output channel.
//
// When the done channel is closed or a value is sent to it, we stop reading from the input
// channels and return after closing the output channel. The consumer of the function can use
// done as a cancellation mechanism or a cleanup mechanism - that is close the done channel
// after you have read all the required values from the output channel.
//
// # Example
//
//     done := make(chan struct{})
//     defer close(done)
//
//     c := make(chan int)
//
//     go func() {
//         c <- 1
//         c <- 2
//         c <- 3
//     }
//
//     for n := range concurrency.Take(done, c, 2) {
//         fmt.Println(n) // prints 1 and 2
//     }
//
//     fmt.Println(<-c) // prints 3
func Take(done <-chan struct{}, c <-chan int, count int) <-chan int {
	out := make(chan int)
	if count <= 0 {
		close(out)
		return out
	}

	// block chan is used here to wait for the goroutine below to terminate.
	block := make(chan struct{})

	// Start a goroutine for the input channel. It copies values from its input channel to the out
	// channel until the input channel is closed or count values have been copied. Then the
	// goroutine sends to block to indicate it is about to terminate.
	go func() {
		defer func() {
			block <- struct{}{}
		}()

		for n := range c {
			select {
			case out <- n:
				count--
				if count <= 0 {
					return
				}
			case <-done:
				return
			}
		}
	}()

	// Start a goroutine to close the out channel once the output goroutine is done.
	go func() {
		<-block
		close(out)
	}()
	return out
}
