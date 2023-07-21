package queue

import (
	"github.com/gpahal/go-algos/ds/list"
)

// ListQueue represents a queue instance implemented as a singly linked list.
type ListQueue struct {
	l *list.SinglyLinkedList
}

// NewListQueue returns a new list queue instance with the given items enqueued into it.
func NewListQueue(items ...int) Interface {
	q := &ListQueue{l: &list.SinglyLinkedList{}}
	q.Enqueue(items...)
	return q
}

// Len returns the number of items in the queue.
func (q *ListQueue) Len() int {
	return q.l.Len()
}

// Empty checks whether the queue is empty.
func (q *ListQueue) Empty() bool {
	return q.l.Empty()
}

// Clear deletes all the items from the queue.
func (q *ListQueue) Clear() {
	q.l.Clear()
}

// Front returns the front/oldest enqueued element of the queue. If the queue is empty, second
// return value is false.
func (q *ListQueue) Front() (int, bool) {
	el := q.l.First()
	if el == nil {
		return 0, false
	}

	return el.Value, true
}

// Enqueue adds the items at the end of the queue.
func (q *ListQueue) Enqueue(items ...int) {
	q.l.PushBack(items...)
}

// Dequeue removes the item from the front of the queue and returns it. If the queue is empty,
// second return value is false.
func (q *ListQueue) Dequeue() (int, bool) {
	el := q.l.PopFront()
	if el == nil {
		return 0, false
	}

	return el.Value, true
}

// Copy creates a new copy of the queue.
func (q *ListQueue) Copy() Interface {
	var arr []int
	q.l.Each(func(item int) bool {
		arr = append(arr, item)
		return false
	})

	return NewListQueue(arr...)
}
