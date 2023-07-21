package set

// NativeSet represents a set instance implemented using a native Go map.
type NativeSet struct {
	m map[int]struct{}
}

// NewNativeSet returns a new native set with the given items added to it.
func NewNativeSet(items ...int) Interface {
	newSet := &NativeSet{m: nil}
	newSet.Add(items...)
	return newSet
}

// Len returns the number of items in the set.
func (s *NativeSet) Len() int {
	return len(s.m)
}

// Empty checks whether the set is empty.
func (s *NativeSet) Empty() bool {
	return len(s.m) == 0
}

// Clear deletes all the items from the set.
func (s *NativeSet) Clear() {
	s.m = nil
}

// Contains checks whether the set contains all the given items.
func (s *NativeSet) Contains(items ...int) bool {
	if len(items) == 0 {
		return true
	}
	if s.m == nil {
		return false
	}

	for _, item := range items {
		_, ok := s.m[item]
		if !ok {
			return false
		}
	}

	return true
}

// Each iterates over the items of the set.
func (s *NativeSet) Each(fn func(int) bool) {
	for item := range s.m {
		if fn(item) {
			break
		}
	}
}

// Iterator returns a set.Iterable that can be used to iterate over the set.
func (s *NativeSet) Iterator() Iterable {
	return &nativeSetIterable{
		values:     s.values(),
		currentIdx: -1,
	}
}

// Add adds the given items to the set.
func (s *NativeSet) Add(items ...int) {
	if s.m == nil {
		s.m = make(map[int]struct{}, len(items))
	}

	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

// Delete deletes the given items from the set.
func (s *NativeSet) Delete(items ...int) {
	if s.m == nil {
		return
	}

	for _, item := range items {
		delete(s.m, item)
	}
}

// Copy creates a new copy of the set.
func (s *NativeSet) Copy() Interface {
	if s.m == nil {
		return NewNativeSet()
	}

	m := make(map[int]struct{}, len(s.m))
	for item := range s.m {
		m[item] = struct{}{}
	}

	return &NativeSet{m: m}
}

// values returns a slice of the items of the set.
func (s *NativeSet) values() []int {
	items := make([]int, len(s.m))
	idx := 0
	for item := range s.m {
		items[idx] = item
		idx++
	}

	return items
}

type nativeSetIterable struct {
	values     []int
	currentIdx int
}

func (it *nativeSetIterable) Next() bool {
	if it.currentIdx >= len(it.values)-1 {
		return false
	}

	it.currentIdx++
	return true
}

func (it *nativeSetIterable) Value() int {
	if it.currentIdx < 0 || it.currentIdx >= len(it.values) {
		return 0
	}

	return it.values[it.currentIdx]
}
