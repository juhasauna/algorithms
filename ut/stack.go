package ut

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		panic("pop from empty stack")
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *Stack[T]) Peek() T {
	if len(s.data) == 0 {
		panic("peek from empty stack")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}
