package main

// Generic struct with methods
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items)
	v := s.items[n-1]
	s.items = s.items[:n-1]
	return v
}

// Invalid: Methods cannot introduce their own type params.
// Uncommenting the following will cause a compilation error.
// func (s Stack[T]) Map[U any](f func(T) U) {}

// Correct one: Function that takes a generic struct as parameter.
func Map[T any, U any](s Stack[T], f func(T) U) (result Stack[U]) {
	return result
}
