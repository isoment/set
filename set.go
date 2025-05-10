package set

type Item = comparable
type SetValue = struct{}

type Set[T Item] struct {
	// add mutex for safe concurrency
	data map[T]SetValue
}

func NewSet[T Item](items []T) *Set[T] {
	data := make(map[T]SetValue)
	for _, v := range items {
		data[v] = SetValue{}
	}
	return &Set[T]{data}
}

func NewEmptySet[T Item]() *Set[T] {
	data := make(map[T]SetValue)
	return &Set[T]{data}
}

func (s *Set[T]) Add(value T) *Set[T] {
	_, ok := s.data[value]
	if !ok {
		s.data[value] = SetValue{}
	}
	return s
}

func (s *Set[T]) Clear() *Set[T] {
	empty := make(map[T]SetValue)
	s.data = empty
	return s
}

func (s *Set[T]) Delete(value T) *Set[T] {
	delete(s.data, value)
	return s
}

/*
Returns a new set with all items that are in the receiver set
but not in the other set.
*/
func (s Set[T]) Difference(other *Set[T]) (r *Set[T]) {
	r = NewEmptySet[T]()

	if s.Size() == 0 {
		return
	}

	for v := range s.data {
		if !other.Has(v) {
			r.Add(v)
		}
	}

	return
}

/*
Execute the provided function for each item in the set
*/
func (s Set[T]) Each(f func(a T)) {
	for v := range s.data {
		f(v)
	}
}

func (s *Set[T]) Has(value T) bool {
	_, ok := s.data[value]
	return ok
}

/*
Returns a new set with all common items in the receiver set
and the other set.
*/
func (s Set[T]) Intersection(other *Set[T]) (r *Set[T]) {
	r = NewEmptySet[T]()

	if s.Size() == 0 {
		return
	}

	for v := range s.data {
		if other.Has(v) {
			r.Add(v)
		}
	}

	return
}

func (s Set[T]) Size() int {
	return len(s.data)
}
