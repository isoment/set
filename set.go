package set

type Item = comparable
type SetValue = struct{}

type Set[T Item] struct {
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

func (s *Set[T]) Has(value T) bool {
	_, ok := s.data[value]
	return ok
}
