package sorting

// MergeSort sorts the slice in-place using the merge sort algorithm.
func MergeSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 2 {
		if arrLength == 2 && arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}

		return
	}

	// Recursively sort the two halves of the slice.
	mid := arrLength / 2
	MergeSort(arr[:mid])
	MergeSort(arr[mid:])

	// Merge the sorted halves.
	merge(arr, mid)
}

// MergeSortIterative sorts the slice in-place using the merge sort iterative algorithm.
func MergeSortIterative(arr []int) {
	arrLength := len(arr)
	if arrLength <= 1 {
		return
	}

	// Merge sorted slices bottom up, starting from single element slices and doubling the size in
	// each iteration.
	currSize := 1
	for currSize < arrLength {
		doubleCurrSize := currSize * 2

		// Loop over the whole slice in steps of doubleCurrSize and merge the two consecutive
		// sorted slices of size currSize.
		for startIdx := 0; startIdx+currSize < arrLength; startIdx += doubleCurrSize {
			endIdxExtended := startIdx + doubleCurrSize
			if endIdxExtended > arrLength {
				endIdxExtended = arrLength
			}

			merge(arr[startIdx:endIdxExtended], currSize)
		}

		currSize = doubleCurrSize
	}
}

// merge the sorted slices arr[:mid] and arr[mid:].
func merge(arr []int, mid int) {
	// Check if merge is required.
	if len(arr) < 2 || arr[mid-1] <= arr[mid] {
		return
	}

	arrLength := len(arr)

	// Temporary slice to remember the first half of the original slice.
	tmpArr := make([]int, mid)
	copy(tmpArr, arr[:mid])

	// Merge the two slices, incrementally taking the smallest element among the smallest of the
	// two slices.
	idx, left, right := 0, 0, mid
	for {
		if tmpArr[left] <= arr[right] {
			arr[idx] = tmpArr[left]
			left++
			if left == mid {
				break
			}
		} else {
			arr[idx] = arr[right]
			right++
			if right == arrLength {
				copy(arr[idx+1:], tmpArr[left:mid])
				break
			}
		}

		idx++
	}
}
