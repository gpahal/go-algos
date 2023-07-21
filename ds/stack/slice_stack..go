package stack

// SliceStack represents a stack instance implemented using a slice.
type SliceStack struct {
	arr []int
}

// NewSliceStack returns a new slice stack instance with the given items pushed into it.
func NewSliceStack(items ...int) Interface {
	return newSliceStack(items...)
}

func newSliceStack(items ...int) *SliceStack {
	s := &SliceStack{}
	s.Push(items...)
	return s
}

// Len returns the number of items in the stack.
func (s *SliceStack) Len() int {
	return len(s.arr)
}

// Empty checks whether the stack is empty.
func (s *SliceStack) Empty() bool {
	return len(s.arr) == 0
}

// Clear deletes all the items from the stack.
func (s *SliceStack) Clear() {
	s.arr = nil
}

// Top returns the top/last pushed element of the stack. If the stack is empty, second return
// value is false.
func (s *SliceStack) Top() (int, bool) {
	if len(s.arr) == 0 {
		return 0, false
	}

	return s.arr[len(s.arr)-1], true
}

// Push pushes the given items to the stack.
func (s *SliceStack) Push(items ...int) {
	if len(items) == 0 {
		return
	}

	s.arr = append(s.arr, items...)
}

// Pop pops out an item from the stack in LIFO (Last In First Out) order. If the stack is empty,
// second return value is false.
func (s *SliceStack) Pop() (int, bool) {
	if len(s.arr) == 0 {
		return 0, false
	}

	v := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return v, true
}

// Copy creates a new copy of the stack.
func (s *SliceStack) Copy() Interface {
	newStack := newSliceStack()
	if len(s.arr) == 0 {
		return newStack
	}

	newStack.arr = make([]int, len(s.arr))
	copy(newStack.arr, s.arr)
	return newStack
}
