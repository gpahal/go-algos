package set_test

import (
	"testing"

	"github.com/gpahal/go-algos/ds/set"
)

func testInterfaceHelper(t *testing.T, newFn func(items ...int) set.Interface) {
	t.Run("New", func(t *testing.T) {
		newSet := newFn(4, 5, 6)
		if newSet.Len() != 3 {
			t.Errorf("New 4, 5, 6: expected Len to be 3, got %d", newSet.Len())
		}

		assertSetValues(t, "New 4, 5, 6", newSet, makeSet(4, 5, 6))
	})

	t.Run("Len", func(t *testing.T) {
		newSet := newFn()
		if newSet.Len() != 0 {
			t.Errorf("Len: expected Len to be 0, got %d", newSet.Len())
		}

		newSet.Add(4, 5, 6)
		if newSet.Len() != 3 {
			t.Errorf("Len: expected Len to be 3, got %d", newSet.Len())
		}
	})

	t.Run("Empty", func(t *testing.T) {
		newSet := newFn()
		if !newSet.Empty() {
			t.Error("Empty: expected Empty to be true, got false")
		}

		newSet.Add(4, 5, 6)
		if newSet.Empty() {
			t.Error("Empty: expected Empty to be false, got true")
		}
	})

	t.Run("Clear", func(t *testing.T) {
		newSet := newFn(4, 5, 6)
		newSet.Clear()
		if !newSet.Empty() {
			t.Errorf("Clear: expected Empty to be true, got false")
		}
	})

	t.Run("Contains", func(t *testing.T) {
		newSet := newFn(4, 5, 6)
		ok := newSet.Contains(4, 5)
		if !ok {
			t.Errorf("Contains 4, 5: expected Contains to return true, got false")
		}

		newSet.Add(7)
		ok = newSet.Contains(7)
		if !ok {
			t.Errorf("Contains 7: expected Contains to return true, got false")
		}

		newSet.Delete(4)
		ok = newSet.Contains(7, 4)
		if ok {
			t.Errorf("Contains 7: expected Contains to return false, got true")
		}
	})

	t.Run("Each", func(t *testing.T) {
		newSet := newFn(4, 5, 6, 4)

		m := make(map[int]struct{})
		newSet.Each(func(item int) bool {
			m[item] = struct{}{}
			return false
		})
		expected := makeSet(4, 5, 6)
		if !mapsEqual(m, expected) {
			t.Errorf("Each: expected Set values to be %v, got %v", expected, m)
		}

		newSet.Add(7)
		m = make(map[int]struct{})
		newSet.Each(func(item int) bool {
			m[item] = struct{}{}
			return false
		})
		expected = makeSet(4, 5, 6, 7)
		if !mapsEqual(m, expected) {
			t.Errorf("Each: expected Set values to be %v, got %v", expected, m)
		}
	})

	t.Run("Iterator", func(t *testing.T) {
		newSet := newFn(4, 5, 6, 4)

		m := make(map[int]struct{})
		it := newSet.Iterator()
		for it.Next() {
			m[it.Value()] = struct{}{}
		}
		expected := makeSet(4, 5, 6)
		if !mapsEqual(m, expected) {
			t.Errorf("Iterator: expected Set values to be %v, got %v", expected, m)
		}

		newSet.Add(7)
		m = make(map[int]struct{})
		it = newSet.Iterator()
		for it.Next() {
			m[it.Value()] = struct{}{}
		}
		expected = makeSet(4, 5, 6, 7)
		if !mapsEqual(m, expected) {
			t.Errorf("Iterator: expected Set values to be %v, got %v", expected, m)
		}
	})

	t.Run("Add", func(t *testing.T) {
		newSet := newFn(4, 5, 6, 4)
		ok := newSet.Contains(4, 5)
		if !ok {
			t.Errorf("Contains 4, 5: expected Contains to return true, got false")
		}

		newSet.Add(7)
		ok = newSet.Contains(7)
		if !ok {
			t.Errorf("Contains 7: expected Contains to return true, got false")
		}

		newSet.Delete(4)
		ok = newSet.Contains(7, 4)
		if ok {
			t.Errorf("Contains 7: expected Contains to return false, got true")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		newSet := newFn(4, 5, 6, 4)
		ok := newSet.Contains(4, 5, 6)
		if !ok {
			t.Errorf("Delete: expected Contains(4, 5, 6) to return true, got false")
		}

		newSet.Delete(4, 5)
		ok = newSet.Contains(4)
		if ok {
			t.Errorf("Delete: expected Contains(4) to return false, got true")
		}

		ok = newSet.Contains(5)
		if ok {
			t.Errorf("Delete: expected Contains(5) to return false, got true")
		}

		ok = newSet.Contains(6)
		if !ok {
			t.Errorf("Delete: expected Contains(6) to return true, got false")
		}
	})

	t.Run("Copy", func(t *testing.T) {
		newSet := newFn(4, 5, 6)
		copiedSet := newSet.Copy()
		assertSetValues(t, "Copy", newSet, makeSet(4, 5, 6))
		copiedSet.Add(3)
		assertSetValues(t, "Copy", copiedSet, makeSet(4, 5, 6, 3))
		assertSetValues(t, "Copy", newSet, makeSet(4, 5, 6))
	})
}

func TestAreEqual(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	newSet3 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)

	if !set.AreEqual(newSet1, newSet3) {
		t.Error("AreEqual 1, 3: expected AreEqual to return true, got false")
	}
	if set.AreEqual(newSet1, newSet2) {
		t.Error("AreEqual 1, 2: expected AreEqual to return false, got true")
	}
}

func TestIsSubset(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	newSet3 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet4 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4, 90, 12)

	if !set.IsSubset(newSet1, newSet4) {
		t.Error("IsSubset 1, 4: expected IsSubset to return true, got false")
	}
	if !set.IsSubset(newSet1, newSet3) {
		t.Error("IsSubset 1, 3: expected IsSubset to return true, got false")
	}
	if set.IsSubset(newSet1, newSet2) {
		t.Error("IsSubset 1, 2: expected IsSubset to return false, got true")
	}
	if set.IsSubset(newSet4, newSet1) {
		t.Error("IsSubset 4, 1: expected IsSubset to return false, got true")
	}
}

func TestIsSuperset(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	newSet3 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet4 := set.NewNativeSet(1, 7, 8, 10, 12)

	if !set.IsSuperset(newSet1, newSet4) {
		t.Error("IsSuperset 1, 4: expected IsSuperset to return true, got false")
	}
	if !set.IsSuperset(newSet1, newSet3) {
		t.Error("IsSuperset 1, 3: expected IsSuperset to return true, got false")
	}
	if set.IsSuperset(newSet1, newSet2) {
		t.Error("IsSuperset 1, 2: expected IsSuperset to return false, got true")
	}
	if set.IsSuperset(newSet4, newSet1) {
		t.Error("IsSuperset 4, 1: expected IsSuperset to return false, got true")
	}
}

func TestAreDisjoint(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	newSet3 := set.NewNativeSet(-1, -7, -8, -10, -12, -15, 4)

	if !set.AreDisjoint(newSet1, newSet3) {
		t.Error("AreDisjoint 1, 3: expected AreDisjoint to return true, got false")
	}
	if set.AreDisjoint(newSet1, newSet2) {
		t.Error("AreDisjoint 1, 2: expected AreDisjoint to return false, got true")
	}
}

func TestMergeInto(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	set.MergeInto(newSet1, newSet2)
	assertSetValues(t, "MergeInto", newSet1, makeSet(1, 7, 8, 10, 12, 15, -4, 3, -9, -2, 18))
}

func TestRetainOnly(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	set.RetainOnly(newSet1, newSet2)
	assertSetValues(t, "RetainOnly", newSet1, makeSet(7, 8, 12, 15))
}

func TestSeparateFrom(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	set.SeparateFrom(newSet1, newSet2)
	assertSetValues(t, "SeparateFrom", newSet1, makeSet(1, 10, -4))
}

func TestUnion(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	gotUnionSet := set.Union(newSet1, newSet2)
	assertSetValues(t, "Union", gotUnionSet, makeSet(1, 7, 8, 10, 12, 15, -4, 3, -9, -2, 18))
}

func TestIntersection(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	gotIntersectionSet := set.Intersection(newSet1, newSet2)
	assertSetValues(t, "Intersection", gotIntersectionSet, makeSet(7, 8, 12, 15))
}

func TestDifference(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	gotDifferenceSet := set.Difference(newSet1, newSet2)
	assertSetValues(t, "Difference", gotDifferenceSet, makeSet(1, 10, -4))
}

func TestSymmetricDifference(t *testing.T) {
	newSet1 := set.NewNativeSet(1, 7, 8, 10, 12, 15, -4)
	newSet2 := set.NewNativeSet(3, -9, 7, 15, -2, 18, 12, 8)
	gotDifferenceSet := set.SymmetricDifference(newSet1, newSet2)
	assertSetValues(t, "SymmetricDifference", gotDifferenceSet, makeSet(1, 10, -4, 3, -9, -2, 18))
}

func assertSetValues(t *testing.T, name string, s set.Interface, expected map[int]struct{}) {
	t.Helper()

	m := make(map[int]struct{}, s.Len())
	s.Each(func(item int) bool {
		m[item] = struct{}{}
		return false
	})

	if !mapsEqual(m, expected) {
		t.Errorf("%s: expected Set values to be %v, got %v", name, expected, m)
	}
}

func mapsEqual(m1, m2 map[int]struct{}) bool {
	if len(m1) != len(m2) {
		return false
	}

	for item := range m1 {
		if _, ok := m2[item]; !ok {
			return false
		}
	}

	return true
}

func makeSet(items ...int) map[int]struct{} {
	m := make(map[int]struct{}, len(items))
	for _, item := range items {
		m[item] = struct{}{}
	}

	return m
}
