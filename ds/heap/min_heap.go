package heap

// MinHeap represents a min heap.
type MinHeap struct {
	arr []int
}

// NewMinHeap returns a new min heap instance with the given items inserted into it.
func NewMinHeap(values ...int) *MinHeap {
	h := &MinHeap{arr: make([]int, 0, len(values))}
	h.Insert(values...)
	return h
}

// Len returns the number of items in the heap.
func (h *MinHeap) Len() int {
	return len(h.arr)
}

// Empty checks whether the heap is empty.
func (h *MinHeap) Empty() bool {
	return len(h.arr) == 0
}

// Clear deletes all the items from the heap.
func (h *MinHeap) Clear() {
	h.arr = h.arr[:0]
	return
}

func (h *MinHeap) heapifyUp(idx int) {
	curr := idx
	for curr != 0 && h.arr[(curr-1)/2] > h.arr[curr] {
		h.arr[(curr-1)/2], h.arr[curr] = h.arr[curr], h.arr[(curr-1)/2]
		curr = (curr - 1) / 2
	}
}

func (h *MinHeap) heapifyDown(idx int) {
	curr := idx
	for {
		left := 2*curr + 1
		right := left + 1
		smallest := curr
		if left < len(h.arr) && h.arr[left] < h.arr[smallest] {
			smallest = left
		}
		if right < len(h.arr) && h.arr[right] < h.arr[smallest] {
			smallest = right
		}

		if smallest == curr {
			break
		}

		h.arr[curr], h.arr[smallest] = h.arr[smallest], h.arr[curr]
		curr = smallest
	}
}

// Min returns the minimum value in the heap.
func (h *MinHeap) Min() (int, bool) {
	if len(h.arr) == 0 {
		return 0, false
	}
	return h.arr[0], true
}

// Find returns the index of value in the heap. If the value is not found, -1 is returned.
func (h *MinHeap) Find(value int) int {
	if len(h.arr) == 0 {
		return -1
	}

	return h.findHelper(0, value)
}

func (h *MinHeap) findHelper(idx, value int) int {
	if idx < 0 || idx >= len(h.arr) {
		return -1
	}
	if h.arr[idx] == value {
		return idx
	}
	if h.arr[idx] > value {
		return -1
	}

	l := h.findHelper(idx*2+1, value)
	if l >= 0 {
		return l
	}

	return h.findHelper(idx*2+2, value)
}

// Insert inserts the given items to the heap.
func (h *MinHeap) Insert(values ...int) {
	for _, value := range values {
		h.arr = append(h.arr, value)
		h.heapifyUp(len(h.arr) - 1)
	}
}

// UpdateAt updates the item at the given index.
func (h *MinHeap) UpdateAt(idx, newValue int) {
	if idx < 0 || idx >= len(h.arr) {
		return
	}

	prev := h.arr[idx]
	h.arr[idx] = newValue
	if newValue <= prev {
		h.heapifyUp(idx)
	} else {
		h.heapifyDown(idx)
	}
}

// ExtractMin removes the minimum value from the heap and returns it. If the heap is empty, the
// second return value is false.
func (h *MinHeap) ExtractMin() (int, bool) {
	if len(h.arr) == 0 {
		return 0, false
	}

	v := h.arr[0]
	if len(h.arr) == 0 {
		h.arr = h.arr[:0]
		return v, true
	}

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown(0)
	return v, true
}

// ExtractAt removes the value at index from the heap and returns it. If the heap doesn't enough
// elements, the second return value is false.
func (h *MinHeap) ExtractAt(idx int) (int, bool) {
	if idx < 0 || idx >= len(h.arr) {
		return 0, false
	}
	if idx == 0 {
		return h.ExtractMin()
	}

	v := h.arr[idx]
	m, _ := h.Min()
	h.UpdateAt(idx, m-1)
	h.ExtractMin()
	return v, true
}

// Copy creates a new copy of the heap.
func (h *MinHeap) Copy() *MinHeap {
	return NewMinHeap(h.arr...)
}
