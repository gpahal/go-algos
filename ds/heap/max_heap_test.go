package heap_test

import (
	"testing"

	"github.com/gpahal/go-algos/ds/heap"
)

func TestNewMaxHeap(t *testing.T) {
	newHeap := heap.NewMaxHeap(4, 5, 6)
	if newHeap.Len() != 3 {
		t.Errorf("NewMaxHeap 4, 5, 6: expected Len to be 3, got %d", newHeap.Len())
	}

	assertMaxHeap(t, "NewMaxHeap", newHeap, []int{6, 5, 4})
}

func TestMaxHeap_Len(t *testing.T) {
	newHeap := heap.NewMaxHeap()
	if newHeap.Len() != 0 {
		t.Errorf("Len: expected Len to be 0, got %d", newHeap.Len())
	}

	newHeap.Insert(4, 5, 6)
	if newHeap.Len() != 3 {
		t.Errorf("Len: expected Len to be 3, got %d", newHeap.Len())
	}
}

func TestMaxHeap_Empty(t *testing.T) {
	newHeap := heap.NewMaxHeap()
	if !newHeap.Empty() {
		t.Error("Empty: expected Empty to be true, got false")
	}

	newHeap.Insert(4, 5, 6)
	if newHeap.Empty() {
		t.Error("Empty: expected Empty to be false, got true")
	}
}

func TestMaxHeap_Clear(t *testing.T) {
	newHeap := heap.NewMaxHeap(4, 5, 6)
	newHeap.Clear()
	if newHeap.Len() != 0 {
		t.Errorf("Clear: expected Len to be 0, got %d", newHeap.Len())
	}
}

func TestMaxHeap_Max(t *testing.T) {
	newHeap := heap.NewMaxHeap()
	val, ok := newHeap.Max()
	if ok {
		t.Errorf("Max: expected Max to return (0, false), got (%d, %t)", val, ok)
	}

	newHeap.Insert(4, 5, 6)
	val, ok = newHeap.Max()
	if !ok || val != 6 {
		t.Errorf("Max: expected Max to return (6, true), got (%d, %t)", val, ok)
	}

	newHeap.Insert(7)
	val, ok = newHeap.Max()
	if !ok || val != 7 {
		t.Errorf("Max: expected Max to return (7, true), got (%d, %t)", val, ok)
	}

	newHeap.ExtractMax()
	val, ok = newHeap.Max()
	if !ok || val != 6 {
		t.Errorf("Max: expected Max to return (6, true), got (%d, %t)", val, ok)
	}
}

func TestMaxHeap_Find(t *testing.T) {
	newHeap := heap.NewMaxHeap()
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
	if idx != 0 {
		t.Errorf("Find 7: expected Find to return 0, got %d", idx)
	}

	newHeap.ExtractMax()
	idx = newHeap.Find(7)
	if idx >= 0 {
		t.Errorf("Find 7: expected Find to return -1, got %d", idx)
	}

	newHeap.Insert(6)
	idx = newHeap.Find(6)
	if idx != 0 {
		t.Errorf("Find: expected Find to return 0, got %d", idx)
	}
}

func TestMaxHeap_Insert(t *testing.T) {
	newHeap := heap.NewMaxHeap()
	newHeap.Insert(4)
	idx := newHeap.Find(4)
	if idx != 0 {
		t.Errorf("Insert: expected Find to return 0, got %d", idx)
	}

	newHeap.Insert(2, 8, 6, 4, 3)
	assertMaxHeap(t, "Insert", newHeap, []int{8, 6, 4, 4, 3, 2})
}

func TestMaxHeap_UpdateAt(t *testing.T) {
	newHeap := heap.NewMaxHeap(2)
	newHeap.UpdateAt(0, 9)
	assertMaxHeap(t, "DecreaseIndex", newHeap, []int{9})

	newHeap.Insert(2, 3, 8, 1, 9, 7)
	assertMaxHeap(t, "DecreaseIndex", newHeap, []int{9, 9, 8, 7, 3, 2, 1})

	newHeap.UpdateAt(0, 5)
	assertMaxHeap(t, "DecreaseIndex", newHeap, []int{9, 8, 7, 5, 3, 2, 1})

	idx := newHeap.Find(3)
	if idx < 0 {
		t.Errorf("DecreaseIndex: expected Find to return -1, got %d", idx)
	}

	newHeap.UpdateAt(idx, 0)
	assertMaxHeap(t, "DecreaseIndex", newHeap, []int{9, 8, 7, 5, 2, 1, 0})
}

func TestMaxHeap_Copy(t *testing.T) {
	newHeap := heap.NewMaxHeap(4, 5, 6)
	copiedHeap := newHeap.Copy()
	assertMaxHeap(t, "Copy", copiedHeap, []int{6, 5, 4})

	copiedHeap.Insert(3)
	assertMaxHeap(t, "Copy", copiedHeap, []int{6, 5, 4, 3})

	assertMaxHeap(t, "Copy", newHeap, []int{6, 5, 4})
}

func assertMaxHeap(t *testing.T, name string, h *heap.MaxHeap, expected []int) {
	t.Helper()

	var got []int
	for {
		val, ok := h.ExtractMax()
		if !ok {
			break
		}

		got = append(got, val)
	}

	if !slicesEqual(expected, got) {
		t.Errorf("%s: expected MaxHeap values to be %v, got %v", name, expected, got)
	}
	h.Insert(got...)
}
