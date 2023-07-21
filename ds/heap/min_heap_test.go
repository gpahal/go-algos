package heap_test

import (
	"testing"

	"github.com/gpahal/go-algos/ds/heap"
)

func TestNewMinHeap(t *testing.T) {
	newHeap := heap.NewMinHeap(4, 5, 6)
	if newHeap.Len() != 3 {
		t.Errorf("NewMinHeap 4, 5, 6: expected Len to be 3, got %d", newHeap.Len())
	}

	assertMinHeap(t, "NewMinHeap", newHeap, []int{4, 5, 6})
}

func TestMinHeap_Len(t *testing.T) {
	newHeap := heap.NewMinHeap()
	if newHeap.Len() != 0 {
		t.Errorf("Len: expected Len to be 0, got %d", newHeap.Len())
	}

	newHeap.Insert(4, 5, 6)
	if newHeap.Len() != 3 {
		t.Errorf("Len: expected Len to be 3, got %d", newHeap.Len())
	}
}

func TestMinHeap_Empty(t *testing.T) {
	newHeap := heap.NewMinHeap()
	if !newHeap.Empty() {
		t.Error("Empty: expected Empty to be true, got false")
	}

	newHeap.Insert(4, 5, 6)
	if newHeap.Empty() {
		t.Error("Empty: expected Empty to be false, got true")
	}
}

func TestMinHeap_Clear(t *testing.T) {
	newHeap := heap.NewMinHeap(4, 5, 6)
	newHeap.Clear()
	if newHeap.Len() != 0 {
		t.Errorf("Clear: expected Len to be 0, got %d", newHeap.Len())
	}
}

func TestMinHeap_Min(t *testing.T) {
	newHeap := heap.NewMinHeap()
	val, ok := newHeap.Min()
	if ok {
		t.Errorf("Min: expected Min to return (0, false), got (%d, %t)", val, ok)
	}

	newHeap.Insert(4, 5, 6)
	val, ok = newHeap.Min()
	if !ok || val != 4 {
		t.Errorf("Min: expected Min to return (4, true), got (%d, %t)", val, ok)
	}

	newHeap.Insert(7)
	val, ok = newHeap.Min()
	if !ok || val != 4 {
		t.Errorf("Min: expected Min to return (4, true), got (%d, %t)", val, ok)
	}

	newHeap.ExtractMin()
	val, ok = newHeap.Min()
	if !ok || val != 5 {
		t.Errorf("Min: expected Min to return (5, true), got (%d, %t)", val, ok)
	}
}

func TestMinHeap_Find(t *testing.T) {
	newHeap := heap.NewMinHeap()
	idx := newHeap.Find(4)
	if idx >= 0 {
		t.Errorf("Find 4: expected Find to return -1, got %d", idx)
	}

	newHeap.Insert(4)
	idx = newHeap.Find(4)
	if idx != 0 {
		t.Errorf("Find 4: expected Find to return 0, got %d", idx)
	}

	newHeap.Insert(7)
	idx = newHeap.Find(7)
	if idx != 1 {
		t.Errorf("Find 7: expected Find to return 1, got %d", idx)
	}

	newHeap.ExtractMin()
	idx = newHeap.Find(4)
	if idx >= 0 {
		t.Errorf("Find 4: expected Find to return -1, got %d", idx)
	}

	newHeap.Insert(6)
	idx = newHeap.Find(6)
	if idx != 0 {
		t.Errorf("Find 6: expected Find to return 0, got %d", idx)
	}
}

func TestMinHeap_Insert(t *testing.T) {
	newHeap := heap.NewMinHeap()
	newHeap.Insert(4)
	idx := newHeap.Find(4)
	if idx != 0 {
		t.Errorf("Insert: expected Find to return 0, got %d", idx)
	}

	newHeap.Insert(2, 8, 6, 4, 3)
	assertMinHeap(t, "Insert", newHeap, []int{2, 3, 4, 4, 6, 8})
}

func TestMinHeap_UpdateAt(t *testing.T) {
	newHeap := heap.NewMinHeap(2)
	newHeap.UpdateAt(0, 1)
	assertMinHeap(t, "DecreaseIndex", newHeap, []int{1})

	newHeap.Insert(2, 3, 8, 1, 9, 7)
	assertMinHeap(t, "DecreaseIndex", newHeap, []int{1, 1, 2, 3, 7, 8, 9})

	newHeap.UpdateAt(0, 5)
	assertMinHeap(t, "DecreaseIndex", newHeap, []int{1, 2, 3, 5, 7, 8, 9})

	idx := newHeap.Find(3)
	if idx < 0 {
		t.Errorf("DecreaseIndex: expected Find to return -1, got %d", idx)
	}

	newHeap.UpdateAt(idx, 0)
	assertMinHeap(t, "DecreaseIndex", newHeap, []int{0, 1, 2, 5, 7, 8, 9})
}

func TestMinHeap_Copy(t *testing.T) {
	newHeap := heap.NewMinHeap(4, 5, 6)
	copiedHeap := newHeap.Copy()
	assertMinHeap(t, "Copy", copiedHeap, []int{4, 5, 6})

	copiedHeap.Insert(3)
	assertMinHeap(t, "Copy", copiedHeap, []int{3, 4, 5, 6})

	assertMinHeap(t, "Copy", newHeap, []int{4, 5, 6})
}

func assertMinHeap(t *testing.T, name string, h *heap.MinHeap, expected []int) {
	t.Helper()

	var got []int
	for {
		val, ok := h.ExtractMin()
		if !ok {
			break
		}

		got = append(got, val)
	}

	if !slicesEqual(expected, got) {
		t.Errorf("%s: expected MinHeap values to be %v, got %v", name, expected, got)
	}
	h.Insert(got...)
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
