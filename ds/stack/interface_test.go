package stack_test

import (
	"testing"

	"github.com/gpahal/go-algos/ds/stack"
)

func testInterfaceHelper(t *testing.T, newFn func(items ...int) stack.Interface) {
	t.Run("New", func(t *testing.T) {
		newStack := newFn(4, 5, 6)
		if newStack.Len() != 3 {
			t.Errorf("New 4, 5, 6: expected Len to be 3, got %d", newStack.Len())
		}

		assertStackValues(t, "New 4, 5, 6", newStack, []int{6, 5, 4})
	})

	t.Run("Len", func(t *testing.T) {
		newStack := newFn()
		if newStack.Len() != 0 {
			t.Errorf("Len: expected Len to be 0, got %d", newStack.Len())
		}

		newStack.Push(4, 5, 6)
		if newStack.Len() != 3 {
			t.Errorf("Len: expected Len to be 3, got %d", newStack.Len())
		}
	})

	t.Run("Empty", func(t *testing.T) {
		newStack := newFn()
		if !newStack.Empty() {
			t.Error("Empty: expected Empty to be true, got false")
		}

		newStack.Push(4, 5, 6)
		if newStack.Empty() {
			t.Error("Empty: expected Empty to be false, got true")
		}
	})

	t.Run("Clear", func(t *testing.T) {
		newStack := newFn(4, 5, 6)
		newStack.Clear()
		if !newStack.Empty() {
			t.Errorf("Clear: expected Empty to be true, got false")
		}
	})

	t.Run("Top", func(t *testing.T) {
		newStack := newFn(4, 5, 6)
		newStack.Push(7)
		val, ok := newStack.Top()
		if !ok || val != 7 {
			t.Errorf("Top: expected Top to return (7, true), got (%d, %t)", val, ok)
		}

		assertStackValues(t, "Top", newStack, []int{7, 6, 5, 4})
		newStack.Clear()
		val, ok = newStack.Top()
		if ok || val != 0 {
			t.Errorf("Top: expected Top to return (0, false), got (%d, %t)", val, ok)
		}
	})

	t.Run("Push", func(t *testing.T) {
		newStack := newFn()
		newStack.Push(4, 5, 6)
		assertStackValues(t, "Push 4, 5, 6", newStack, []int{6, 5, 4})
		newStack.Push(7)
		assertStackValues(t, "Push 7", newStack, []int{7, 6, 5, 4})
	})

	t.Run("Pop", func(t *testing.T) {
		newStack := newFn(4, 5, 6)
		newStack.Push(7)
		val, ok := newStack.Pop()
		if !ok || val != 7 {
			t.Errorf("Pop: expected Pop to return (7, true), got (%d, %t)", val, ok)
		}

		assertStackValues(t, "Pop", newStack, []int{6, 5, 4})
		newStack.Clear()
		val, ok = newStack.Pop()
		if ok || val != 0 {
			t.Errorf("Pop: expected Pop to return (0, false), got (%d, %t)", val, ok)
		}
	})

	t.Run("Copy", func(t *testing.T) {
		newStack := newFn(4, 5, 6)
		copiedStack := newStack.Copy()
		assertStackValues(t, "Copy", newStack, []int{6, 5, 4})
		copiedStack.Push(3)
		assertStackValues(t, "Copy", copiedStack, []int{3, 6, 5, 4})
		assertStackValues(t, "Copy", newStack, []int{6, 5, 4})
	})
}

func assertStackValues(t *testing.T, name string, s stack.Interface, expected []int) {
	t.Helper()

	var got []int
	for {
		val, ok := s.Pop()
		if !ok {
			break
		}

		got = append(got, val)
	}

	if !slicesEqual(expected, got) {
		t.Errorf("%s: expected Stack values to be %v, got %v", name, expected, got)
	}
	for i := len(got) - 1; i >= 0; i-- {
		s.Push(got[i])
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
