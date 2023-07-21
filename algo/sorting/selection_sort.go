package sorting

// SelectionSort sorts the slice in-place using the selection sort algorithm.
func SelectionSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 1 {
		return
	}

	// Loop over every index except the last one.
	for i := 0; i < arrLength-1; i++ {
		minIdx := i
		minVal := arr[i]

		// Find the minimum element in arr[i:].
		for j := i + 1; j < arrLength; j++ {
			if arr[j] < minVal {
				minIdx = j
				minVal = arr[j]
			}
		}

		// Check if the minimum element in not a[i]. If not, swap the minimum element with the ith
		// element.
		if minIdx != i {
			arr[i], arr[minIdx] = minVal, arr[i]
		}
	}
}
