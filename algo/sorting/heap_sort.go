package sorting

// HeapSort sorts the slice in-place using the heap sort algorithm.
func HeapSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 2 {
		if arrLength == 2 && arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}

		return
	}

	// Build the heap. Note that all elements with index >= arrLength/2 are leaf nodes, so heapify
	// doesn't need to used for those indices.
	for i := arrLength/2 - 1; i >= 0; i-- {
		heapify(arr, arrLength, i)
	}

	// Extract max values one by one, place them at the end, ignore them and build the heap again.
	for i := arrLength - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

// heapify makes sure element at index idx is >= its children (index 2*idx + 1 and 2*idx + 2). If
// not, it swaps the element at index idx and the largest of the three elements (say at index
// largestIdx), and runs heapify with idx as largestIdx. Effectively, this function "heapifies" the
// subtree rooted at idx if all of its left and right subtrees are already "heapified".
func heapify(arr []int, size, idx int) {
	largestIdx := idx
	leftIdx := idx*2 + 1
	rightIdx := idx*2 + 2

	// Find largestIdx.

	if leftIdx < size && arr[leftIdx] > arr[largestIdx] {
		largestIdx = leftIdx
	}
	if rightIdx < size && arr[rightIdx] > arr[largestIdx] {
		largestIdx = rightIdx
	}

	// If the root is not the largest, swap elements at index idx and largestIdx and run heapify
	// with idx as largestIdx.
	if largestIdx != idx {
		arr[idx], arr[largestIdx] = arr[largestIdx], arr[idx]
		heapify(arr, size, largestIdx)
	}
}
