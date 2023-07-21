package list

// Interface is the interface that groups the basic methods of a list implementation.
type Interface interface {
	// Len returns the number of items in the list.
	Len() int

	// Empty checks whether the list is empty.
	Empty() bool

	// Clear deletes all the items from the list.
	Clear()

	// First returns the first element of the list.
	First() *Element

	// Last returns the last element of the list.
	Last() *Element

	// At returns the (i+1)th element of the list. Negative indices can also be used to find the
	// (-i)th last element.
	At(i int) *Element

	// Contains checks whether the list contains all the given items.
	Contains(items ...int) bool

	// Each iterates over the items of the list.
	Each(fn func(int) bool)

	// Iterator returns a list.Iterable that can be used to iterate over the list.
	Iterator() Iterable

	// PushFront adds the given items at the start of the list.
	PushFront(items ...int)

	// PopFront removes and returns the first element from the list. If the list is empty, it
	// returns nil.
	PopFront() *Element

	// PushBack adds the given items at the end of the list.
	PushBack(items ...int)

	// PopBack removes and returns the last element from the list. If the list is empty, it returns
	// nil.
	PopBack() *Element

	// InsertAt adds the item as the (i+1)th element and returns the element. Negative indices can
	// also be used to insert after the (-i)th last element. If the list doesn't have enough
	// elements, it returns nil.
	InsertAt(i int, item int) *Element

	// InsertAfter adds the item after the given element and returns the inserted element. If
	// inserting after e is not possible, it returns nil.
	InsertAfter(e *Element, item int) *Element

	// InsertBefore adds the item before the given element and returns the inserted element. If
	// inserting after e is not possible, it returns nil.
	InsertBefore(e *Element, item int) *Element

	// RemoveAt removes the (i+1)th element. Negative indices can also be used to remove the (-i)th
	// last element. If the list doesn't have enough elements, it returns nil.
	RemoveAt(i int) *Element

	// Remove removes and returns the given element. If removing e is not possible, it returns nil.
	Remove(e *Element) *Element

	// RemoveAfter removes and returns the element after the given element. If removing after e is
	// not possible, it returns nil.
	RemoveAfter(e *Element) *Element

	// RemoveBefore removes and returns the element before the given element. If removing before e
	// is not possible, it returns nil.
	RemoveBefore(e *Element) *Element

	// DeleteFirst deletes the first occurrence of the given items from the list. If the same item
	// is passed twice as an argument, only one occurrence is deleted in total.
	DeleteFirst(items ...int)

	// Delete deletes the all occurrences of the given items from the list.
	Delete(items ...int)

	// Copy creates a new copy of the list.
	Copy() Interface
}

// Element is a single list element.
type Element struct {
	Value int
	Next  *Element
	Prev  *Element
}

// Iterable is the interface that groups the Next and Value methods used to iterate over a
// list.Interface.
type Iterable interface {
	// Next prepares the next item for reading with the Value method. It returns true on success,
	// or false if there are no items left, after which the Value method would always return the
	// zero value.
	//
	// Every call to Value, even the first one, must be preceded by a call to Next.
	Next() bool

	// Value reads and returns the item prepared by the Next method.
	Value() int
}
