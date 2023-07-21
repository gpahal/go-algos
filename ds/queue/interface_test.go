package queue_test

import (
	"testing"

	"github.com/gpahal/go-algos/ds/queue"
)

func testInterfaceHelper(t *testing.T, newFn func(items ...int) queue.Interface) {
	t.Run("New", func(t *testing.T) {
		newQueue := newFn(4, 5, 6)
		if newQueue.Len() != 3 {
			t.Errorf("New 4, 5, 6: expected Len to be 3, got %d", newQueue.Len())
		}

		assertQueueValues(t, "New 4, 5, 6", newQueue, []int{4, 5, 6})
	})

	t.Run("Len", func(t *testing.T) {
		newQueue := newFn()
		if newQueue.Len() != 0 {
			t.Errorf("Len: expected Len to be 0, got %d", newQueue.Len())
		}

		newQueue.Enqueue(4, 5, 6)
		if newQueue.Len() != 3 {
			t.Errorf("Len: expected Len to be 3, got %d", newQueue.Len())
		}
	})

	t.Run("Empty", func(t *testing.T) {
		newQueue := newFn()
		if !newQueue.Empty() {
			t.Error("Empty: expected Empty to be true, got false")
		}

		newQueue.Enqueue(4, 5, 6)
		if newQueue.Empty() {
			t.Error("Empty: expected Empty to be false, got true")
		}
	})

	t.Run("Clear", func(t *testing.T) {
		newQueue := newFn(4, 5, 6)
		newQueue.Clear()
		if !newQueue.Empty() {
			t.Error("Clear: expected Empty to be true, got false")
		}
	})

	t.Run("Front", func(t *testing.T) {
		newQueue := newFn(4, 5, 6, 7)
		val, ok := newQueue.Front()
		if !ok || val != 4 {
			t.Errorf("Front: expected Front to return (4, true), got (%d, %t)", val, ok)
		}

		assertQueueValues(t, "Front", newQueue, []int{4, 5, 6, 7})
		newQueue.Clear()
		val, ok = newQueue.Front()
		if ok || val != 0 {
			t.Errorf("Front: expected Front to return (0, false), got (%d, %t)", val, ok)
		}
	})

	t.Run("Enqueue", func(t *testing.T) {
		newQueue := newFn()
		newQueue.Enqueue(4, 5, 6)
		assertQueueValues(t, "Enqueue 4, 5, 6", newQueue, []int{4, 5, 6})
		newQueue.Enqueue(7)
		assertQueueValues(t, "Enqueue 7", newQueue, []int{4, 5, 6, 7})
	})

	t.Run("Dequeue", func(t *testing.T) {
		newQueue := newFn(4, 5, 6)
		newQueue.Enqueue(7)
		val, ok := newQueue.Dequeue()
		if !ok || val != 4 {
			t.Errorf("Dequeue: expected Dequeue to return (4, true), got (%d, %t)", val, ok)
		}

		assertQueueValues(t, "Dequeue", newQueue, []int{5, 6, 7})
		newQueue.Clear()
		val, ok = newQueue.Dequeue()
		if ok || val != 0 {
			t.Errorf("Dequeue: expected Dequeue to return (0, false), got (%d, %t)", val, ok)
		}
	})

	t.Run("Copy", func(t *testing.T) {
		newQueue := newFn(4, 5, 6)
		copiedQueue := newQueue.Copy()
		assertQueueValues(t, "Copy", newQueue, []int{4, 5, 6})
		copiedQueue.Enqueue(3)
		assertQueueValues(t, "Copy", copiedQueue, []int{4, 5, 6, 3})
		assertQueueValues(t, "Copy", newQueue, []int{4, 5, 6})
	})
}

func assertQueueValues(t *testing.T, name string, q queue.Interface, expected []int) {
	t.Helper()

	var got []int
	for {
		val, ok := q.Dequeue()
		if !ok {
			break
		}

		got = append(got, val)
	}

	if !slicesEqual(expected, got) {
		t.Errorf("%s: expected Queue values to be %v, got %v", name, expected, got)
	}
	q.Enqueue(got...)
}

func slicesEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, el := range arr1 {
		if el != arr2[i] {
			return false
		}
	}

	return true
}
