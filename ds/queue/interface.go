package queue

// Interface is the interface that groups the basic methods of a queue implementation.
type Interface interface {
	// Length returns the number of items in the queue.
	Len() int

	// Empty checks whether the queue is empty.
	Empty() bool

	// Clear deletes all the items from the queue.
	Clear()

	// Front returns the front/oldest enqueued element of the queue. If the queue is empty, second
	// return value is false.
	Front() (int, bool)

	// Enqueue adds the items at the end of the queue.
	Enqueue(items ...int)

	// Dequeue removes the item from the front of the queue and returns it. If the queue is empty,
	// second return value is false.
	Dequeue() (int, bool)

	// Copy creates a new copy of the queue.
	Copy() Interface
}
