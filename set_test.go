package set

import (
	"reflect"
	"slices"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("it creates a new set", func(t *testing.T) {
		s := New([]int{1, 2, 3, 2, 3})
		if len(s.data) != 3 {
			t.Errorf("expected 3 items but got %d", len(s.data))
		}
		assertTrue(t, s.Has(1))
		assertTrue(t, s.Has(2))
		assertTrue(t, s.Has(3))
	})
}

func TestAdd(t *testing.T) {
	s := NewEmpty[int]()
	s.Add(1)
	s.Add(2)

	if len(s.data) != 2 {
		t.Errorf("expected 2 items but got %d", len(s.data))
	}

	assertTrue(t, s.Has(1))
	assertTrue(t, s.Has(2))
}

func TestClear(t *testing.T) {
	s := New([]string{"a", "b", "c"})
	s.Clear()

	size := s.Size()
	if size != 0 {
		t.Errorf("expected empty set, got %d items", size)
	}
}

func TestDelete(t *testing.T) {
	s := New([]string{"a", "b", "c"})
	s.Delete("a")
	s.Delete("b")

	size := s.Size()
	if size != 1 {
		t.Errorf("expected size to be 1, got %d", size)
	}

	assertTrue(t, s.Has("c"))
}

func TestDifference(t *testing.T) {
	testCases := []struct {
		name     string
		s1       *Set[int]
		s2       *Set[int]
		expected *Set[int]
	}{
		{
			name:     "it returns an empty set if receiver set is empty",
			s1:       NewEmpty[int](),
			s2:       New([]int{1, 2, 3}),
			expected: NewEmpty[int](),
		},
		{
			name:     "it returns items in receiver set that are not in the other set",
			s1:       New([]int{4, 7, 8, 23, 138}),
			s2:       New([]int{6, 8, 54, 65, 102, 138}),
			expected: New([]int{4, 7, 23}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.Difference(tc.s2)
			if !got.Equal(tc.expected) {
				t.Errorf("expected: %v but got: %v", tc.expected, got)
			}
		})
	}
}

func TestEach(t *testing.T) {
	count := 0

	r := New([]int{4, 7, 8, 23, 138})
	r.Each(func(a int) {
		count++
	})

	if count != 5 {
		t.Errorf("Expected 5 function calls, go %d", count)
	}
}

func TestHas(t *testing.T) {
	s := New([]string{"a", "b", "c"})
	assertTrue(t, s.Has("a"))
	assertTrue(t, s.Has("b"))
	assertTrue(t, s.Has("c"))
	assertFalse(t, s.Has("z"))
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		name     string
		s1       *Set[int]
		s2       *Set[int]
		expected bool
	}{
		{
			name:     "it returns false for different size sets",
			s1:       New([]int{1, 2}),
			s2:       New([]int{1, 2, 3, 4}),
			expected: false,
		},
		{
			name:     "it returns true for empty sets",
			s1:       NewEmpty[int](),
			s2:       NewEmpty[int](),
			expected: true,
		},
		{
			name:     "it returns true for equal sets",
			s1:       New([]int{1, 2, 3, 4, 5}),
			s2:       New([]int{1, 2, 5, 4, 3}),
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.Equal(tc.s2)
			if got != tc.expected {
				t.Errorf("got: %v but expected %v", got, tc.expected)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	r := New([]int{4, 7, 8, 23, 138})
	s := New([]int{9, 21, 22, 23, 87, 132, 138})

	got := r.Intersection(s)
	if got.Size() != 2 {
		t.Errorf("expected size 2 got %d", got.Size())
	}

	assertTrue(t, s.Has(23))
	assertTrue(t, s.Has(138))
}

func TestIsDisjointFrom(t *testing.T) {
	testCases := []struct {
		name     string
		s1       *Set[int]
		s2       *Set[int]
		expected bool
	}{
		{
			name:     "it is disjointed from",
			s1:       New([]int{1, 3, 5, 7}),
			s2:       New([]int{2, 4, 6, 8}),
			expected: true,
		},
		{
			name:     "it is not disjointed from",
			s1:       New([]int{0, 1, 3}),
			s2:       New([]int{2, 4, 1}),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.IsDisjointFrom(tc.s2)
			if got != tc.expected {
				t.Errorf("got: %v but expected %v", got, tc.expected)
			}
		})
	}
}

func TestIsSubsetOf(t *testing.T) {
	testCases := []struct {
		name     string
		s1       *Set[int]
		s2       *Set[int]
		expected bool
	}{
		{
			name:     "it is not subset, receiver has more items",
			s1:       New([]int{1, 2, 3, 4, 5}),
			s2:       New([]int{1, 2, 3}),
			expected: false,
		},
		{
			name:     "it is subset",
			s1:       New([]int{1, 2, 3}),
			s2:       New([]int{1, 2, 3, 4, 5}),
			expected: true,
		},
		{
			name:     "it is not subset",
			s1:       New([]int{1, 3, 70}),
			s2:       New([]int{1, 2, 3, 4}),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.IsSubsetOf(tc.s2)
			if got != tc.expected {
				t.Errorf("got: %v but expected %v", got, tc.expected)
			}
		})
	}
}

func TestIsSupersetOf(t *testing.T) {
	testCases := []struct {
		name     string
		s1       *Set[int]
		s2       *Set[int]
		expected bool
	}{
		{
			name:     "it is not superset, other has more items",
			s1:       New([]int{1, 2, 3}),
			s2:       New([]int{1, 2, 3, 4, 5}),
			expected: false,
		},
		{
			name:     "it is superset",
			s1:       New([]int{1, 2, 3, 4, 5}),
			s2:       New([]int{1, 2, 3}),
			expected: true,
		},
		{
			name:     "it is not superset",
			s1:       New([]int{1, 2, 3, 4}),
			s2:       New([]int{1, 3, 70}),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.IsSupersetOf(tc.s2)
			if got != tc.expected {
				t.Errorf("got: %v but expected %v", got, tc.expected)
			}
		})
	}
}

func TestSize(t *testing.T) {
	s := New([]string{"a", "b", "c"})
	size := s.Size()
	if size != 3 {
		t.Errorf("expected size to be 3, got %d", size)
	}
}

func TestSymmetricDifference(t *testing.T) {
	testCases := []struct {
		name string
		s1   *Set[int]
		s2   *Set[int]
		want *Set[int]
	}{
		{
			name: "case 1",
			s1:   New([]int{1, 2, 3, 9}),
			s2:   New([]int{1, 2, 3, 4, 5}),
			want: New([]int{4, 5, 9}),
		},
		{
			name: "case 2",
			s1:   New([]int{1, 2, 3}),
			s2:   New([]int{1, 2, 3}),
			want: NewEmpty[int](),
		},
		{
			name: "case 3",
			s1:   New([]int{1, 2}),
			s2:   New([]int{3, 4, 5, 350}),
			want: New([]int{1, 2, 3, 4, 5, 350}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.SymmetricDifference(tc.s2)
			if !got.Equal(tc.want) {
				t.Errorf("got: %v but want %v", got, tc.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		name string
		s1   *Set[int]
		s2   *Set[int]
		want *Set[int]
	}{
		{
			name: "case 1",
			s1:   New([]int{1, 2, 3, 9}),
			s2:   New([]int{4, 5}),
			want: New([]int{1, 2, 3, 4, 5, 9}),
		},
		{
			name: "case 2",
			s1:   New([]int{1, 2, 3}),
			s2:   New([]int{1, 2, 3}),
			want: New([]int{1, 2, 3}),
		},
		{
			name: "case 3",
			s1:   NewEmpty[int](),
			s2:   NewEmpty[int](),
			want: NewEmpty[int](),
		},
		{
			name: "case 4",
			s1:   New([]int{25}),
			s2:   NewEmpty[int](),
			want: New([]int{25}),
		},
		{
			name: "case 5",
			s1:   NewEmpty[int](),
			s2:   New([]int{25, 8, 2, 14}),
			want: New([]int{25, 8, 2, 14}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.Union(tc.s2)
			if !got.Equal(tc.want) {
				t.Errorf("got: %v but want %v", got, tc.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}

	s := New(want)
	got := s.Values()

	if len(got) != 5 {
		t.Errorf("expected 5 elements got %d", len(got))
	}

	slices.Sort(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
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
