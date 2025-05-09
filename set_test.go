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
