package set

// Interface is the interface that groups the basic methods of a set implementation.
type Interface interface {
	// Len returns the number of items in the set.
	Len() int

	// Empty checks whether the set is empty.
	Empty() bool

	// Clear deletes all the items from the set.
	Clear()

	// Contains checks whether the set contains all the given items.
	Contains(items ...int) bool

	// Each iterates over the items of the set.
	Each(fn func(int) bool)

	// Iterator returns a set.Iterable that can be used to iterate over the set.
	Iterator() Iterable

	// Add adds the given items to the set.
	Add(items ...int)

	// Delete deletes the given items from the set.
	Delete(items ...int)

	// Copy creates a new copy of the set.
	Copy() Interface
}

// Iterable is the interface that groups the Next and Value methods used to iterate over a
// set.Interface.
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

// AreEqual checks whether the two sets have the same items.
func AreEqual(set1, set2 Interface) bool {
	if set1.Len() != set2.Len() {
		return false
	}

	equal := true
	set1.Each(func(item int) bool {
		if !set2.Contains(item) {
			equal = false
			return true
		}

		return false
	})

	return equal
}

// IsSubset checks whether the first set is the subset of the second.
func IsSubset(subSet, superSet Interface) bool {
	if subSet.Len() > superSet.Len() {
		return false
	}

	subset := true
	subSet.Each(func(item int) bool {
		if !superSet.Contains(item) {
			subset = false
			return true
		}

		return false
	})

	return subset
}

// IsSuperset checks whether the first set is the superset of the second.
func IsSuperset(superSet, subSet Interface) bool {
	return IsSubset(subSet, superSet)
}

// AreDisjoint checks whether the two sets are disjoint.
func AreDisjoint(set1, set2 Interface) bool {
	if set1.Len() == 0 {
		return true
	}

	disjoint := true
	set1.Each(func(item int) bool {
		if set2.Contains(item) {
			disjoint = false
			return true
		}

		return false
	})

	return disjoint
}

// MergeInto adds all the items of the second set to the first.
func MergeInto(mainSet, otherSet Interface) {
	otherSet.Each(func(item int) bool {
		mainSet.Add(item)
		return false
	})
}

// RetainOnly deletes all the items from the first set that are not in the second.
func RetainOnly(mainSet, otherSet Interface) {
	mainSet.Each(func(item int) bool {
		if !otherSet.Contains(item) {
			mainSet.Delete(item)
		}

		return false
	})
}

// SeparateFrom deletes all the items from the first set that are also in the second.
func SeparateFrom(mainSet, otherSet Interface) {
	otherSet.Each(func(item int) bool {
		mainSet.Delete(item)
		return false
	})
}

// Union returns a new set which is a union of the given sets. If no sets are provided, an empty
// set is returned.
func Union(sets ...Interface) Interface {
	newSet := NewNativeSet()
	for _, set := range sets {
		MergeInto(newSet, set)
	}

	return newSet
}

// Intersection returns a new set which is an intersection of the given sets. If no sets are
// provided, an empty set is returned.
func Intersection(sets ...Interface) Interface {
	if len(sets) == 0 {
		return NewNativeSet()
	}

	newSet := sets[0].Copy()
	sets = sets[1:]
	for _, set := range sets {
		RetainOnly(newSet, set)
	}

	return newSet
}

// Difference returns a new set which is the first set minus all the other sets. If only one set
// is provided, a copy of that set is returned. If no sets are provided, an empty set is returned.
func Difference(sets ...Interface) Interface {
	if len(sets) == 0 {
		return NewNativeSet()
	}

	newSet := sets[0].Copy()
	sets = sets[1:]
	for _, set := range sets {
		SeparateFrom(newSet, set)
	}

	return newSet
}

// SymmetricDifference returns the symmetric difference of the two sets.
func SymmetricDifference(set1, set2 Interface) Interface {
	return Union(Difference(set1, set2), Difference(set2, set1))
}
