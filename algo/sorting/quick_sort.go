package sorting

// QuickSort sorts the slice in-place using the quick sort algorithm.
func QuickSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 2 {
		if arrLength == 2 && arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}

		return
	}

	// Partition the slice.
	partitionIdx := partition(arr)

	// Recursively sort the two partitions.
	if partitionIdx < arrLength/2 {
		QuickSort(arr[:partitionIdx])
		QuickSort(arr[partitionIdx+1:])
	} else {
		QuickSort(arr[partitionIdx+1:])
		QuickSort(arr[:partitionIdx])
	}
}

// QuickSortIterative sorts the slice in-place using the quick sort iterative algorithm.
func QuickSortIterative(arr []int) {
	arrLength := len(arr)
	if arrLength <= 1 {
		return
	}

	// Create a stack of length (arrLength+1)*2.
	stack := make([][]int, 0, (arrLength+1)*2)
	stack = append(stack, arr)

	// Iterate until the stack is empty.
	for len(stack) > 0 {
		topIdx := len(stack) - 1
		currArr := stack[topIdx]
		stack = stack[:topIdx]

		// Partition the current slice.
		partitionIdx := partition(currArr)

		// Add the two partitions to the stack.
		if partitionIdx > 1 {
			stack = append(stack, currArr[:partitionIdx])
		}
		if partitionIdx < len(currArr)-2 {
			stack = append(stack, currArr[partitionIdx+1:])
		}
	}
}

// partition partitions arr using the middle element as the pivot.
func partition(arr []int) int {
	arrLength := len(arr)

	// Choose the middle element as the pivot.
	lastIdx := arrLength - 1
	pivotIdx := arrLength / 2
	pivotVal := arr[pivotIdx]

	arr[lastIdx], arr[pivotIdx] = pivotVal, arr[lastIdx]

	// Partition the slice.
	i := 0
	for j := 0; j < lastIdx; j++ {
		if arr[j] <= pivotVal {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}

			i++
		}
	}
	if i != lastIdx {
		arr[i], arr[lastIdx] = pivotVal, arr[i]
	}

	return i
}
