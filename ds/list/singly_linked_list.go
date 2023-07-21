package list

// SinglyLinkedList represents a list instance implemented as a singly linked list. The Prev field
// of every is nil.
type SinglyLinkedList struct {
	head *Element
}

// NewSinglyLinkedList returns a new singly linked list instance with the given items inserted in
// order.
func NewSinglyLinkedList(items ...int) Interface {
	newList := &SinglyLinkedList{}
	for i := len(items) - 1; i >= 0; i-- {
		newList.PushFront(items[i])
	}
	return newList
}

// Len returns the number of items in the list.
func (sll *SinglyLinkedList) Len() int {
	if sll.head == nil {
		return 0
	}

	count := 1
	for curr := sll.head; curr.Next != nil; curr = curr.Next {
		count++
	}

	return count
}

// Empty checks whether the list is empty.
func (sll *SinglyLinkedList) Empty() bool {
	return sll.head == nil
}

// Clear deletes all the items from the list.
func (sll *SinglyLinkedList) Clear() {
	sll.head = nil
}

// First returns the first element of the list.
func (sll *SinglyLinkedList) First() *Element {
	return sll.head
}

// Last returns the last element of the list.
func (sll *SinglyLinkedList) Last() *Element {
	if sll.head == nil {
		return nil
	}

	var curr *Element
	for curr = sll.head; curr.Next != nil; curr = curr.Next {
	}

	return curr
}

// secondLast returns the second last element of the list.
func (sll *SinglyLinkedList) secondLast() *Element {
	if sll.head == nil || sll.head.Next == nil {
		return nil
	}

	var curr *Element
	for curr = sll.head; curr.Next != nil && curr.Next.Next != nil; curr = curr.Next {
	}

	return curr
}

// At returns the (i+1)th element of the list. Negative indices can also be used to find the (-i)th
// last element.
func (sll *SinglyLinkedList) At(i int) *Element {
	if sll.head == nil {
		return nil
	}
	if i < 0 {
		if i == -1 {
			return sll.Last()
		} else if i == -2 {
			return sll.secondLast()
		}

		i += sll.Len()
		if i < 0 {
			return nil
		}
	}

	currIdx := 0
	for curr := sll.head; curr != nil; curr = curr.Next {
		if currIdx == i {
			return curr
		}

		currIdx++
	}

	return nil
}

// Contains checks whether the list contains all the given items.
func (sll *SinglyLinkedList) Contains(items ...int) bool {
	if len(items) == 0 {
		return true
	}
	if sll.head == nil {
		return false
	}

	itemsMap := make(map[int]struct{}, len(items))
	for _, item := range items {
		itemsMap[item] = struct{}{}
	}

	for curr := sll.head; curr != nil; curr = curr.Next {
		if _, ok := itemsMap[curr.Value]; ok {
			delete(itemsMap, curr.Value)
		}
	}

	return len(itemsMap) == 0
}

// Each iterates over the items of the list.
func (sll *SinglyLinkedList) Each(fn func(int) bool) {
	if sll.head == nil {
		return
	}

	for curr := sll.head; curr != nil; curr = curr.Next {
		if fn(curr.Value) {
			break
		}
	}
}

// Iterator returns a list.Iterable that can be used to iterate over the list.
func (sll *SinglyLinkedList) Iterator() Iterable {
	return &singlyLinkedListIterable{
		curr: sll.head,
	}
}

// PushFront adds the given items at the start of the list.
func (sll *SinglyLinkedList) PushFront(items ...int) {
	if len(items) == 0 {
		return
	}

	var curr *Element
	for _, item := range items {
		if curr == nil {
			curr = &Element{
				Value: item,
				Next:  sll.head,
			}
		} else {
			curr = &Element{
				Value: item,
				Next:  curr,
			}
		}
	}

	sll.head = curr
}

// PopFront removes and returns the first element from the list. If the list is empty, it returns
// nil.
func (sll *SinglyLinkedList) PopFront() *Element {
	if sll.head == nil {
		return nil
	}

	removedEl := sll.head
	sll.head = sll.head.Next
	removedEl.Next = nil
	return removedEl
}

// PushBack adds the given items at the end of the list.
func (sll *SinglyLinkedList) PushBack(items ...int) {
	if len(items) == 0 {
		return
	}

	var curr *Element
	for i := len(items) - 1; i >= 0; i-- {
		item := items[i]
		if curr == nil {
			curr = &Element{
				Value: item,
			}
		} else {
			curr = &Element{
				Value: item,
				Next:  curr,
			}
		}
	}

	tail := sll.Last()
	if tail == nil {
		sll.head = curr
	} else {
		tail.Next = curr
	}
}

// PopBack removes and returns the last element from the list. If the list is empty, it returns
// nil.
func (sll *SinglyLinkedList) PopBack() *Element {
	if sll.head == nil {
		return nil
	}

	var tail *Element
	preTail := sll.secondLast()
	if preTail == nil {
		tail = sll.head
	} else {
		tail = preTail.Next
	}

	preTail.Next = nil
	return tail
}

// InsertAt adds the item as the (i+1)th element and returns the element. Negative indices can also
// be used to insert after the (-i)th last element. If the list doesn't have enough elements, it
// returns nil.
func (sll *SinglyLinkedList) InsertAt(i int, item int) *Element {
	if i < 0 {
		if i == -1 {
			return sll.InsertAfter(sll.Last(), item)
		} else if i == -2 {
			return sll.InsertAfter(sll.secondLast(), item)
		}

		i += sll.Len() + 1
		if i < 0 {
			return nil
		}
	}
	if i == 0 {
		sll.PushFront(item)
		return sll.head
	}

	return sll.InsertAfter(sll.At(i-1), item)
}

// InsertAfter adds the item after the given element and returns the inserted element. If inserting
// after e is not possible, it returns nil.
func (sll *SinglyLinkedList) InsertAfter(e *Element, item int) *Element {
	if e == nil {
		return nil
	}

	e.Next = &Element{
		Value: item,
		Next:  e.Next,
	}
	return e.Next
}

// InsertBefore adds the item before the given element and returns the inserted element. If
// inserting after e is not possible, it returns nil.
func (sll *SinglyLinkedList) InsertBefore(e *Element, item int) *Element {
	if e == nil || sll.head == nil {
		return nil
	}
	if e == sll.head {
		sll.PushFront(item)
		return sll.head
	}

	var curr *Element
	for curr = sll.head; curr.Next != nil && curr.Next != e; curr = curr.Next {
	}

	if curr.Next == nil {
		return nil
	}

	return sll.InsertAfter(curr, item)
}

// RemoveAt removes the (i+1)th element. Negative indices can also be used to remove the (-i)th
// last element. If the list doesn't have enough elements, it returns nil.
func (sll *SinglyLinkedList) RemoveAt(i int) *Element {
	if i < 0 {
		if i == -1 {
			return sll.RemoveAfter(sll.secondLast())
		}

		i += sll.Len()
		if i < 0 {
			return nil
		}
	}
	if i == 0 {
		return sll.PopFront()
	}

	return sll.RemoveAfter(sll.At(i - 1))
}

// Remove removes and returns the given element. If removing e is not possible, it returns nil.
func (sll *SinglyLinkedList) Remove(e *Element) *Element {
	if e == nil || sll.head == nil {
		return nil
	}
	if e == sll.head {
		return sll.PopFront()
	}

	var curr *Element
	for curr = sll.head; curr.Next != nil && curr.Next != e; curr = curr.Next {
	}

	if curr.Next == nil {
		return nil
	}

	return sll.RemoveAfter(curr)
}

// RemoveAfter removes and returns the element after the given element. If removing after e is not
// possible, it returns nil.
func (sll *SinglyLinkedList) RemoveAfter(e *Element) *Element {
	if e == nil || e.Next == nil {
		return nil
	}

	removedEl := e.Next
	e.Next = removedEl.Next
	removedEl.Next = nil
	return removedEl
}

// RemoveBefore removes and returns the element before the given element. If removing before e is
// not possible, it returns nil.
func (sll *SinglyLinkedList) RemoveBefore(e *Element) *Element {
	if e == nil || e == sll.head || sll.head == nil || sll.head.Next == nil {
		return nil
	}
	if e == sll.head.Next {
		return sll.PopFront()
	}

	var curr *Element
	for curr = sll.head; curr.Next.Next != nil && curr.Next.Next != e; curr = curr.Next {
	}

	if curr.Next.Next == nil {
		return nil
	}

	return sll.RemoveAfter(curr)
}

// DeleteFirst deletes the first occurrence of the given items from the list. If the same item is
// passed twice as an argument, only one occurrence is deleted in total.
func (sll *SinglyLinkedList) DeleteFirst(items ...int) {
	if len(items) == 0 || sll.head == nil {
		return
	}

	itemsMap := make(map[int]struct{}, len(items))
	for _, item := range items {
		itemsMap[item] = struct{}{}
	}

	var prev *Element
	var next *Element
	for curr := sll.head; curr != nil; curr = next {
		next = curr.Next
		if _, ok := itemsMap[curr.Value]; ok {
			delete(itemsMap, curr.Value)
			if prev == nil {
				sll.Remove(sll.head)
			} else {
				sll.RemoveAfter(prev)
			}

			if len(itemsMap) == 0 {
				break
			}
		} else {
			prev = curr
		}
	}
}

// Delete deletes the all occurrences of the given items from the list.
func (sll *SinglyLinkedList) Delete(items ...int) {
	if len(items) == 0 || sll.head == nil {
		return
	}

	itemsMap := make(map[int]struct{}, len(items))
	for _, item := range items {
		itemsMap[item] = struct{}{}
	}

	var prev *Element
	var next *Element
	for curr := sll.head; curr != nil; curr = next {
		next = curr.Next
		if _, ok := itemsMap[curr.Value]; ok {
			if prev == nil {
				sll.Remove(sll.head)
			} else {
				sll.RemoveAfter(prev)
			}
		} else {
			prev = curr
		}
	}
}

// Copy creates a new copy of the list.
func (sll *SinglyLinkedList) Copy() Interface {
	if sll.head == nil {
		return NewSinglyLinkedList()
	}

	var arr []int
	for curr := sll.head; curr != nil; curr = curr.Next {
		arr = append(arr, curr.Value)
	}

	return NewSinglyLinkedList(arr...)
}

type singlyLinkedListIterable struct {
	curr  *Element
	value int
}

func (slli *singlyLinkedListIterable) Next() bool {
	if slli.curr == nil {
		slli.value = 0
		return false
	}

	slli.value = slli.curr.Value
	slli.curr = slli.curr.Next
	return true
}

func (slli *singlyLinkedListIterable) Value() int {
	return slli.value
}
