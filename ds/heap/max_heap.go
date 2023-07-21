package heap

// MaxHeap represents a max heap.
type MaxHeap struct {
	arr []int
}

// NewMaxHeap returns a new max heap instance with the given items inserted into it.
func NewMaxHeap(values ...int) *MaxHeap {
	h := &MaxHeap{arr: make([]int, 0, len(values))}
	h.Insert(values...)
	return h
}

// Len returns the number of items in the heap.
func (h *MaxHeap) Len() int {
	return len(h.arr)
}

// Empty checks whether the heap is empty.
func (h *MaxHeap) Empty() bool {
	return len(h.arr) == 0
}

// Clear deletes all the items from the heap.
func (h *MaxHeap) Clear() {
	h.arr = h.arr[:0]
	return
}

func (h *MaxHeap) heapifyUp(idx int) {
	curr := idx
	for curr != 0 && h.arr[(curr-1)/2] < h.arr[curr] {
		h.arr[(curr-1)/2], h.arr[curr] = h.arr[curr], h.arr[(curr-1)/2]
		curr = (curr - 1) / 2
	}
}

func (h *MaxHeap) heapifyDown(idx int) {
	curr := idx
	for {
		left := 2*curr + 1
		right := left + 1
		largest := curr
		if left < len(h.arr) && h.arr[left] > h.arr[largest] {
			largest = left
		}
		if right < len(h.arr) && h.arr[right] > h.arr[largest] {
			largest = right
		}

		if largest == curr {
			break
		}

		h.arr[curr], h.arr[largest] = h.arr[largest], h.arr[curr]
		curr = largest
	}
}

// Max returns the maximum value in the heap.
func (h *MaxHeap) Max() (int, bool) {
	if len(h.arr) == 0 {
		return 0, false
	}
	return h.arr[0], true
}

// Find returns the index of value in the heap. If the value is not found, -1 is returned.
func (h *MaxHeap) Find(value int) int {
	if len(h.arr) == 0 {
		return -1
	}

	return h.findHelper(0, value)
}

func (h *MaxHeap) findHelper(idx, value int) int {
	if idx < 0 || idx >= len(h.arr) {
		return -1
	}
	if h.arr[idx] == value {
		return idx
	}
	if h.arr[idx] < value {
		return -1
	}

	l := h.findHelper(idx*2+1, value)
	if l >= 0 {
		return l
	}

	return h.findHelper(idx*2+2, value)
}

// Insert inserts the given items to the heap.
func (h *MaxHeap) Insert(values ...int) {
	for _, value := range values {
		h.arr = append(h.arr, value)
		h.heapifyUp(len(h.arr) - 1)
	}
}

// UpdateAt updates the item at the given index.
func (h *MaxHeap) UpdateAt(idx, newValue int) {
	if idx < 0 || idx >= len(h.arr) {
		return
	}

	prev := h.arr[idx]
	h.arr[idx] = newValue
	if newValue >= prev {
		h.heapifyUp(idx)
	} else {
		h.heapifyDown(idx)
	}
}

// ExtractMax removes the maximum value from the heap and returns it. If the heap is empty, the
// second return value is false.
func (h *MaxHeap) ExtractMax() (int, bool) {
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
func (h *MaxHeap) ExtractAt(idx int) (int, bool) {
	if idx < 0 || idx >= len(h.arr) {
		return 0, false
	}
	if idx == 0 {
		return h.ExtractMax()
	}

	v := h.arr[idx]
	m, _ := h.Max()
	h.UpdateAt(idx, m-1)
	h.ExtractMax()
	return v, true
}

// Copy creates a new copy of the heap.
func (h *MaxHeap) Copy() *MaxHeap {
	return NewMaxHeap(h.arr...)
}
