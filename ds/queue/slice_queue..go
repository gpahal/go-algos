package queue

// SliceQueue represents a queue instance implemented using a slice used as a circular buffer.
type SliceQueue struct {
	arr   []int
	start int
	size  int
}

// NewSliceQueue returns a new slice queue instance with the given items enqueued into it.
func NewSliceQueue(items ...int) Interface {
	return newSliceQueue(items...)
}

func newSliceQueue(items ...int) *SliceQueue {
	q := &SliceQueue{}
	q.Enqueue(items...)
	return q
}

// Len returns the number of items in the queue.
func (q *SliceQueue) Len() int {
	return q.size
}

// Empty checks whether the queue is empty.
func (q *SliceQueue) Empty() bool {
	return q.size == 0
}

// Clear deletes all the items from the queue.
func (q *SliceQueue) Clear() {
	q.arr = nil
	q.start = 0
	q.size = 0
}

// Front returns the front/oldest enqueued element of the queue. If the queue is empty, second
// return value is false.
func (q *SliceQueue) Front() (int, bool) {
	if q.size == 0 {
		return 0, false
	}

	return q.arr[q.start], true
}

// Enqueue adds the items at the end of the queue.
func (q *SliceQueue) Enqueue(items ...int) {
	if len(items) == 0 {
		return
	}

	for i := 0; i < len(q.arr)-q.size && i < len(items); i++ {
		q.arr[(q.start+q.size+i)%len(q.arr)] = items[i]
	}
	if len(items) <= len(q.arr)-q.size {
		q.size += len(items)
		return
	}

	// update items and q.size
	items = items[len(q.arr)-q.size:]
	q.size = len(q.arr)

	// slice extension needed
	if q.start == 0 {
		q.arr = append(q.arr, items...)
		q.size = len(q.arr)
		return
	}

	q.arr = append(q.arr, items...)
	j := q.size
	for i := 0; i < q.start; i++ {
		q.arr[j%len(q.arr)] = q.arr[i]
		j++
	}
	for i := 0; i < len(items); i++ {
		q.arr[j%len(q.arr)] = items[i]
		j++
	}
	q.size = len(q.arr)
}

// Dequeue removes the item from the front of the queue and returns it. If the queue is empty,
// second return value is false.
func (q *SliceQueue) Dequeue() (int, bool) {
	if q.size == 0 {
		return 0, false
	}

	v := q.arr[q.start]
	q.start = (q.start + 1) % len(q.arr)
	q.size--
	return v, true
}

// Copy creates a new copy of the queue.
func (q *SliceQueue) Copy() Interface {
	newQueue := newSliceQueue()
	if q.size == 0 {
		return newQueue
	}

	newQueue.start = 0
	newQueue.size = q.size
	newQueue.arr = make([]int, q.size)
	if q.start+q.size > len(q.arr) {
		copy(newQueue.arr[:len(q.arr)-q.start], q.arr[q.start:len(q.arr)])
		copy(newQueue.arr[len(q.arr)-q.start:], q.arr[:q.start+q.size-len(q.arr)])
	} else {
		copy(newQueue.arr, q.arr[q.start:])
	}

	return newQueue
}
