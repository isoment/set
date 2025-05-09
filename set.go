package set

type Item = comparable
type SetValue = struct{}

type Set[T Item] struct {
	// add mutex for safe concurrency
	data map[T]SetValue
}

func (s Set[T]) Size() int {
	return len(s.data)
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

func (s *Set[T]) Has(value T) bool {
	_, ok := s.data[value]
	return ok
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
