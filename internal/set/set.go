package set

type Set[T comparable] struct {
	elements map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *Set[T]) Intersection(otherSet *Set[T]) *Set[T] {
	intersection := New[T]()
	for element := range s.elements {
		if otherSet.Contains(element) {
			intersection.Add(element)
		}
	}
	return intersection
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := New[T]()
	for element := range s.elements {
		union.Add(element)
	}
	for element := range other.elements {
		union.Add(element)
	}
	return union
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) Elements() []T {
	keys := make([]T, 0, len(s.elements))
	for k := range s.elements {
		keys = append(keys, k)
	}
	return keys
}
