package sorting

// BubbleSort sorts the slice in-place using the bubble sort algorithm.
func BubbleSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 1 {
		return
	}

	// Loop until no swaps are done in an iteration.
	for {
		// Counter of the number of swaps done in the current iteration.
		swaps := 0

		// Loop arrLength-1 times over all pairs of consecutive elements and swap in not in the
		// correct order.
		for i := 0; i < arrLength-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swaps++
			}
		}

		// Stop if no swaps were done. It means the array is already sorted.
		if swaps == 0 {
			return
		}
	}
}
