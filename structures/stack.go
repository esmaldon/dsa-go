package structures

type Stack[T any] struct {
	array []T
}

func (s *Stack[T]) push(value T) {
	s.array = append(s.array, value)
}

func (s *Stack[T]) pop() T {
	i := len(s.array) - 1
	v := s.array[i]
	s.array = s.array[:i]

	return v
}
