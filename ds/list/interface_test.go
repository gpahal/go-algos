package list_test

import (
	"testing"

	"github.com/gpahal/go-algos/ds/list"
)

func testInterfaceHelper(t *testing.T, newFn func(...int) list.Interface) {
	t.Run("New", func(t *testing.T) {
		newList := newFn(4, 5, 6)
		if newList.Len() != 3 {
			t.Errorf("New 4, 5, 6: expected Len to be 3, got %d", newList.Len())
		}

		assertListValues(t, "New 4, 5, 6", newList, []int{4, 5, 6})
	})

	t.Run("Len", func(t *testing.T) {
		newList := newFn()
		if newList.Len() != 0 {
			t.Errorf("Len: expected Len to be 0, got %d", newList.Len())
		}

		newList.PushBack(4, 5, 6)
		if newList.Len() != 3 {
			t.Errorf("Len: expected Len to be 3, got %d", newList.Len())
		}
	})

	t.Run("Empty", func(t *testing.T) {
		newList := newFn()
		if !newList.Empty() {
			t.Error("Empty: expected Empty to be true, got false")
		}

		newList.PushBack(4, 5, 6)
		if newList.Empty() {
			t.Error("Empty: expected Empty to be false, got true")
		}
	})

	t.Run("Clear", func(t *testing.T) {
		newList := newFn(4, 5, 6)
		newList.Clear()
		if !newList.Empty() {
			t.Errorf("Clear: expected Empty to be true, got false")
		}
	})

	t.Run("First", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7)
		el := newList.First()
		if el == nil || el.Value != 4 {
			t.Errorf("First: expected First to return Element with Value 4, got %v", el)
		}

		assertListValues(t, "First", newList, []int{4, 5, 6, 7})
		newList.Clear()
		el = newList.First()
		if el != nil {
			t.Errorf("First: expected First to return nil, got %v", el)
		}
	})

	t.Run("Last", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7)
		el := newList.Last()
		if el == nil || el.Value != 7 {
			t.Errorf("Last: expected Last to return Element with Value 7, got %v", el)
		}

		assertListValues(t, "Last", newList, []int{4, 5, 6, 7})
		newList.Clear()
		el = newList.First()
		if el != nil {
			t.Errorf("Last: expected Last to return nil, got %v", el)
		}
	})

	t.Run("At", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7)
		el := newList.At(0)
		if el == nil || el.Value != 4 {
			t.Errorf("At 0: expected At to return Element with Value 4, got %v", el)
		}
		el = newList.At(1)
		if el == nil || el.Value != 5 {
			t.Errorf("At 1: expected At to return Element with Value 5, got %v", el)
		}
		el = newList.At(4)
		if el != nil {
			t.Errorf("At 4: expected At to return nil, got %v", el)
		}
		el = newList.At(-2)
		if el == nil || el.Value != 6 {
			t.Errorf("At -2: expected At to return Element with Value 6, got %v", el)
		}

		assertListValues(t, "At", newList, []int{4, 5, 6, 7})
		newList.Clear()
		el = newList.At(0)
		if el != nil {
			t.Errorf("At: expected At to return nil, got %v", el)
		}
	})

	t.Run("Contains", func(t *testing.T) {
		newList := newFn(4, 5, 6)
		ok := newList.Contains(4, 5)
		if !ok {
			t.Errorf("Contains 4, 5: expected Contains to return true, got false")
		}

		newList.PushBack(7)
		ok = newList.Contains(7)
		if !ok {
			t.Errorf("Contains 7: expected Contains to return true, got false")
		}

		newList.Delete(4)
		ok = newList.Contains(7, 4)
		if ok {
			t.Errorf("Contains 7: expected Contains to return false, got true")
		}
	})

	t.Run("Each", func(t *testing.T) {
		newList := newFn(4, 5, 6, 4)

		var arr []int
		newList.Each(func(item int) bool {
			arr = append(arr, item)
			return false
		})
		expected := []int{4, 5, 6, 4}
		if !slicesEqual(arr, expected) {
			t.Errorf("Each: expected List values to be %v, got %v", expected, arr)
		}

		newList.PushBack(7)
		arr = []int{}
		newList.Each(func(item int) bool {
			arr = append(arr, item)
			return false
		})
		expected = []int{4, 5, 6, 4, 7}
		if !slicesEqual(arr, expected) {
			t.Errorf("Each: expected List values to be %v, got %v", expected, arr)
		}
	})

	t.Run("Iterator", func(t *testing.T) {
		newList := newFn(4, 5, 6, 4)

		var arr []int
		it := newList.Iterator()
		for it.Next() {
			arr = append(arr, it.Value())
		}
		expected := []int{4, 5, 6, 4}
		if !slicesEqual(arr, expected) {
			t.Errorf("Iterator: expected List values to be %v, got %v", expected, arr)
		}

		newList.PushBack(7)
		arr = []int{}
		it = newList.Iterator()
		for it.Next() {
			arr = append(arr, it.Value())
		}
		expected = []int{4, 5, 6, 4, 7}
		if !slicesEqual(arr, expected) {
			t.Errorf("Iterator: expected List values to be %v, got %v", expected, arr)
		}
	})

	t.Run("PushFront", func(t *testing.T) {
		newList := newFn(7)
		newList.PushFront(4, 5, 6)
		ok := newList.Contains(4)
		if !ok {
			t.Errorf("PushFront: expected Contains 4 to be true, got false")
		}

		assertListValues(t, "PushFront", newList, []int{6, 5, 4, 7})
	})

	t.Run("PopFront", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7)
		el := newList.PopFront()
		if el == nil || el.Value != 4 {
			t.Errorf("PopFront: expected PopFront to return Element with Value 4, got %v", el)
		}

		assertListValues(t, "PopFront", newList, []int{5, 6, 7})
		newList.Clear()
		el = newList.PopFront()
		if el != nil {
			t.Errorf("PopFront: expected PopFront to return nil, got %v", el)
		}
	})

	t.Run("PushBack", func(t *testing.T) {
		newList := newFn(7)
		newList.PushBack(4, 5, 6)
		ok := newList.Contains(4)
		if !ok {
			t.Errorf("PushBack: expected Contains 4 to be true, got false")
		}

		assertListValues(t, "PushBack", newList, []int{7, 4, 5, 6})
	})

	t.Run("PopBack", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7)
		el := newList.PopBack()
		if el == nil || el.Value != 7 {
			t.Errorf("PopBack: expected PopBack to return Element with Value 7, got %v", el)
		}

		assertListValues(t, "PopBack", newList, []int{4, 5, 6})
		newList.Clear()
		el = newList.PopBack()
		if el != nil {
			t.Errorf("PopBack: expected PopBack to return nil, got %v", el)
		}
	})

	t.Run("InsertAt", func(t *testing.T) {
		newList := newFn(4, 5, 6)
		el := newList.InsertAt(3, 8)
		if el == nil || el.Value != 8 {
			t.Errorf("InsertAt 3, 8: expected InsertAt to return Element with Value 8, got %v", el)
		}

		assertListValues(t, "InsertAt", newList, []int{4, 5, 6, 8})
		el = newList.InsertAt(-2, 5)
		if el == nil || el.Value != 5 {
			t.Errorf("InsertAt -2, 5: expected InsertAt to return Element with Value 5, got %v", el)
		}

		el = newList.InsertAt(7, 5)
		if el != nil {
			t.Errorf("InsertAt 7, 5: expected InsertAt to return nil, got %v", el)
		}
	})

	t.Run("InsertAfter", func(t *testing.T) {
		newList := newFn(4, 5, 6)
		el := newList.InsertAfter(nil, 0)
		if el != nil {
			t.Errorf("InsertAfter nil, 0: expected InsertAfter to return nil, got %v", el)
		}

		el = newList.InsertAfter(newList.First(), 3)
		if el == nil || el.Value != 3 {
			t.Errorf("InsertAfter First(), 3: expected InsertAfter to return Element with Value 3, got %v", el)
		}

		assertListValues(t, "InsertAfter", newList, []int{4, 3, 5, 6})
		el = newList.InsertAfter(newList.Last(), 7)
		if el == nil || el.Value != 7 {
			t.Errorf("InsertAfter Last(), 7: expected InsertAfter to return Element with Value 7, got %v", el)
		}

		assertListValues(t, "InsertAfter", newList, []int{4, 3, 5, 6, 7})
	})

	t.Run("InsertBefore", func(t *testing.T) {
		newList := newFn(4, 5, 6)
		el := newList.InsertBefore(nil, 0)
		if el != nil {
			t.Errorf("InsertBefore nil, 0: expected InsertBefore to return nil, got %v", el)
		}

		el = newList.InsertBefore(newList.First(), 3)
		if el == nil || el.Value != 3 {
			t.Errorf("InsertBefore First(), 3: expected InsertBefore to return Element with Value 3, got %v", el)
		}

		assertListValues(t, "InsertBefore", newList, []int{3, 4, 5, 6})
		el = newList.InsertBefore(newList.Last(), 7)
		if el == nil || el.Value != 7 {
			t.Errorf("InsertBefore Last(), 7: expected InsertBefore to return Element with Value 7, got %v", el)
		}

		assertListValues(t, "InsertBefore", newList, []int{3, 4, 5, 7, 6})
	})

	t.Run("RemoveAt", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7, 8)
		el := newList.RemoveAt(2)
		if el == nil || el.Value != 6 {
			t.Errorf("RemoveAt 2: expected RemoveAt to return Element with Value 6, got %v", el)
		}

		assertListValues(t, "RemoveAt", newList, []int{4, 5, 7, 8})
		el = newList.RemoveAt(-1)
		if el == nil || el.Value != 8 {
			t.Errorf("RemoveAt -1: expected RemoveAt to return Element with Value 8, got %v", el)
		}

		el = newList.RemoveAt(5)
		if el != nil {
			t.Errorf("RemoveAt 5: expected RemoveAt to return nil, got %v", el)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7, 8)
		el := newList.Remove(nil)
		if el != nil {
			t.Errorf("Remove nil: expected Remove to return nil, got %v", el)
		}

		el = newList.Remove(newList.First())
		if el == nil || el.Value != 4 {
			t.Errorf("Remove First(): expected Remove to return Element with Value 4, got %v", el)
		}

		assertListValues(t, "Remove", newList, []int{5, 6, 7, 8})
		el = newList.Remove(newList.Last())
		if el == nil || el.Value != 8 {
			t.Errorf("Remove Last(): expected Remove to return Element with Value 8, got %v", el)
		}

		assertListValues(t, "RemoveAfter", newList, []int{5, 6, 7})
		el = newList.Remove(newList.At(1))
		if el == nil || el.Value != 6 {
			t.Errorf("Remove At(1): expected Remove to return Element with Value 6, got %v", el)
		}

		assertListValues(t, "RemoveAfter", newList, []int{5, 7})
	})

	t.Run("RemoveAfter", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7, 8)
		el := newList.RemoveAfter(nil)
		if el != nil {
			t.Errorf("RemoveAfter nil: expected RemoveAfter to return nil, got %v", el)
		}

		el = newList.RemoveAfter(newList.First())
		if el == nil || el.Value != 5 {
			t.Errorf("RemoveAfter First(): expected RemoveAfter to return Element with Value 5, got %v", el)
		}

		assertListValues(t, "RemoveAfter", newList, []int{4, 6, 7, 8})
		el = newList.RemoveAfter(newList.Last())
		if el != nil {
			t.Errorf("RemoveAfter Last(): expected RemoveAfter to return nil, got %v", el)
		}

		assertListValues(t, "RemoveAfter", newList, []int{4, 6, 7, 8})
		el = newList.RemoveAfter(newList.At(-2))
		if el == nil || el.Value != 8 {
			t.Errorf("RemoveAfter At(-2): expected RemoveAfter to return Element with Value 8, got %v", el)
		}

		assertListValues(t, "RemoveAfter", newList, []int{4, 6, 7})
	})

	t.Run("RemoveBefore", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7, 8)
		el := newList.RemoveBefore(nil)
		if el != nil {
			t.Errorf("RemoveBefore nil: expected RemoveBefore to return nil, got %v", el)
		}

		el = newList.RemoveBefore(newList.First())
		if el != nil {
			t.Errorf("RemoveBefore First(): expected RemoveBefore to return nil, got %v", el)
		}

		assertListValues(t, "RemoveBefore", newList, []int{4, 5, 6, 7, 8})
		el = newList.RemoveBefore(newList.Last())
		if el == nil || el.Value != 7 {
			t.Errorf("RemoveBefore Last(): expected RemoveBefore to return Element with Value 7, got %v", el)
		}

		assertListValues(t, "RemoveBefore", newList, []int{4, 5, 6, 8})
		el = newList.RemoveBefore(newList.At(1))
		if el == nil || el.Value != 4 {
			t.Errorf("RemoveBefore At(1): expected RemoveBefore to return Element with Value 4, got %v", el)
		}

		assertListValues(t, "RemoveBefore", newList, []int{5, 6, 8})
	})

	t.Run("DeleteFirst", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7, 4)
		newList.DeleteFirst(4, 5)
		assertListValues(t, "DeleteFirst", newList, []int{6, 7, 4})
		newList.Delete(8)
		assertListValues(t, "DeleteFirst", newList, []int{6, 7, 4})
	})

	t.Run("Delete", func(t *testing.T) {
		newList := newFn(4, 5, 6, 7, 4)
		newList.Delete(4, 5)
		assertListValues(t, "DeleteFirst", newList, []int{6, 7})
		newList.Delete(8)
		assertListValues(t, "DeleteFirst", newList, []int{6, 7})
	})

	t.Run("Copy", func(t *testing.T) {
		newList := newFn(4, 5, 6)
		copiedList := newList.Copy()
		assertListValues(t, "Copy", newList, []int{4, 5, 6})
		copiedList.PushBack(3)
		assertListValues(t, "Copy", copiedList, []int{4, 5, 6, 3})
		assertListValues(t, "Copy", newList, []int{4, 5, 6})
	})
}

func assertListValues(t *testing.T, name string, l list.Interface, expected []int) {
	t.Helper()

	if l.Empty() && len(expected) != 0 {
		t.Errorf("%s: expected List values to be %v, got []", name, expected)
	}

	var got []int
	l.Each(func(item int) bool {
		got = append(got, item)
		return false
	})

	if !slicesEqual(got, expected) {
		t.Errorf("%s: expected List values to be %v, got %v", name, expected, got)
	}
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
