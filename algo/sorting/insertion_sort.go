package sorting

// InsertionSort sorts the slice in-place using the insertion sort algorithm.
func InsertionSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 1 {
		return
	}

	// Loop over all the elements except the first one.
	for i := 1; i < arrLength; i++ {
		currVal := arr[i]

		// Place the current value in the correct position, among already visited elements.
		j := i - 1
		for j >= 0 {
			if arr[j] > currVal {
				arr[j+1] = arr[j]
			} else {
				break
			}

			j--
		}

		// Check if any elements have actually moved. If yes, place the current element in the
		// correct position.
		if j != i-1 {
			arr[j+1] = currVal
		}
	}
}
