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

func TestHas(t *testing.T) {
	s := NewSet([]string{"a", "b", "c"})
	assertTrue(t, s.Has("a"))
	assertTrue(t, s.Has("b"))
	assertTrue(t, s.Has("c"))
	assertFalse(t, s.Has("z"))
}

func TesAdd(t *testing.T) {
	s := NewEmptySet[int]()
	s.Add(1)
	s.Add(2)

	if len(s.data) != 2 {
		t.Errorf("expected 2 items but got %d", len(s.data))
	}

	assertTrue(t, s.Has(1))
	assertTrue(t, s.Has(2))
}

func TestSize(t *testing.T) {
	s := NewSet([]string{"a", "b", "c"})
	size := s.Size()
	if size != 3 {
		t.Errorf("expected size to be 3, got %d", size)
	}
}

func TestClear(t *testing.T) {
	s := NewSet([]string{"a", "b", "c"})
	s.Clear()

	size := s.Size()
	if size != 0 {
		t.Errorf("expected empty set, got %d items", size)
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
