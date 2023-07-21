package stack

// Interface is the interface that groups the basic methods of a stack implementation.
type Interface interface {
	// Length returns the number of items in the stack.
	Len() int

	// Empty checks whether the stack is empty.
	Empty() bool

	// Clear deletes all the items from the stack.
	Clear()

	// Top returns the top/last pushed element of the stack. If the stack is empty, second return
	// value is false.
	Top() (int, bool)

	// Push pushes the given items to the stack.
	Push(items ...int)

	// Pop pops out an item from the stack in LIFO (Last In First Out) order. If the stack is
	// empty, second return value is false.
	Pop() (int, bool)

	// Copy creates a new copy of the stack.
	Copy() Interface
}
