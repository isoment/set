package set

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	t.Run("it creates a new set", func(t *testing.T) {
		s := NewSet([]int{1, 2, 3, 2, 3})
		if len(s.data) != 3 {
			t.Errorf("expected 3 items but got %d", len(s.data))
		}
		assertTrue(t, s.Has(1))
		assertTrue(t, s.Has(2))
		assertTrue(t, s.Has(3))
	})
}

func TestAdd(t *testing.T) {
	s := NewEmptySet[int]()
	s.Add(1)
	s.Add(2)

	if len(s.data) != 2 {
		t.Errorf("expected 2 items but got %d", len(s.data))
	}

	assertTrue(t, s.Has(1))
	assertTrue(t, s.Has(2))
}

func TestClear(t *testing.T) {
	s := NewSet([]string{"a", "b", "c"})
	s.Clear()

	size := s.Size()
	if size != 0 {
		t.Errorf("expected empty set, got %d items", size)
	}
}

func TestDelete(t *testing.T) {
	s := NewSet([]string{"a", "b", "c"})
	s.Delete("a")
	s.Delete("b")

	size := s.Size()
	if size != 1 {
		t.Errorf("expected size to be 1, got %d", size)
	}

	assertTrue(t, s.Has("c"))
}

func TestDifference(t *testing.T) {
	t.Run("it returns an empty set if receiver set is empty", func(t *testing.T) {
		s := NewEmptySet[string]()
		u := NewSet([]string{"a", "b", "c"})

		got := s.Difference(u)

		if got.Size() != 0 {
			t.Error("expected empty set")
		}
	})

	t.Run("it returns the difference in a new set", func(t *testing.T) {
		r := NewSet([]int{4, 7, 8, 23, 138})
		s := NewSet([]int{6, 8, 54, 65, 102, 138})

		got := r.Difference(s)

		if got.Size() != 3 {
			t.Errorf("expected set size to be 3 got %d", got.Size())
		}

		assertTrue(t, got.Has(4))
		assertTrue(t, got.Has(7))
		assertTrue(t, got.Has(23))
	})
}

func TestEach(t *testing.T) {
	count := 0

	r := NewSet([]int{4, 7, 8, 23, 138})
	r.Each(func(a int) {
		count++
	})

	if count != 5 {
		t.Errorf("Expected 5 function calls, go %d", count)
	}
}

func TestHas(t *testing.T) {
	s := NewSet([]string{"a", "b", "c"})
	assertTrue(t, s.Has("a"))
	assertTrue(t, s.Has("b"))
	assertTrue(t, s.Has("c"))
	assertFalse(t, s.Has("z"))
}

func TestIntersection(t *testing.T) {
	r := NewSet([]int{4, 7, 8, 23, 138})
	s := NewSet([]int{9, 21, 22, 23, 87, 132, 138})

	got := r.Intersection(s)
	if got.Size() != 2 {
		t.Errorf("expected size 2 got %d", got.Size())
	}

	assertTrue(t, s.Has(23))
	assertTrue(t, s.Has(138))
}

func TestDisjointFrom(t *testing.T) {
	t.Run("it returns true if receiver set has no common items with other", func(t *testing.T) {
		odd := NewSet([]int{1, 3, 5, 7})
		even := NewSet([]int{2, 4, 6, 8})

		got := odd.DisjointFrom(even)

		assertTrue(t, got)
	})

	t.Run("it returns false if receiver set has common items with other", func(t *testing.T) {
		one := NewSet([]string{"a", "b", "d"})
		two := NewSet([]string{"c", "e", "b"})

		got := one.DisjointFrom(two)

		assertFalse(t, got)
	})
}

func TestSize(t *testing.T) {
	s := NewSet([]string{"a", "b", "c"})
	size := s.Size()
	if size != 3 {
		t.Errorf("expected size to be 3, got %d", size)
	}
}

func assertTrue(t *testing.T, a bool) {
	t.Helper()
	if !a {
		t.Error("Expected true but got false")
	}
}

func assertFalse(t *testing.T, a bool) {
	t.Helper()
	if a {
		t.Error("Expected false but got true")
	}
}
