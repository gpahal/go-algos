package searching

// BinarySearch searches the sorted slice arr for the key using the binary search algorithm. If the
// key is not found, -1 is returned.
func BinarySearch(arr []int, key int) int {
	return binarySearchHelper(arr, key, 0)
}

// binarySearchHelper searches the sorted slice arr for the key using the binary search algorithm
// and adds offset idxToBeAdded to the resulting index. If the key is not found, -1 is returned.
func binarySearchHelper(arr []int, key int, idxToBeAdded int) int {
	arrLength := len(arr)
	if arrLength == 0 {
		return -1
	}

	mid := arrLength / 2
	midVal := arr[mid]

	if midVal == key {
		// Found key.
		return mid + idxToBeAdded
	} else if midVal > key {
		// Search the first half of the slice.
		return binarySearchHelper(arr[:mid], key, idxToBeAdded)
	} else {
		// Search the second half of the slice and also increment value of the offset idxToBeAdded.
		return binarySearchHelper(arr[mid+1:], key, idxToBeAdded+mid+1)
	}
}
