package utils

import "errors"

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	s := Stack[T]{}
	s.data = make([]T, 0)
	return &s
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) Push(obj T) {
	s.data = append(s.data, obj)
}

func (s *Stack[T]) Pop() (T, error) {
	length := s.Len()
	if length == 0 {
		var zero T
		return zero, ErrEmptyStack
	}
	obj := s.data[length-1]
	s.data = s.data[:length-1]
	return obj, nil
}

func (s *Stack[T]) Seek() (T, error) {
	length := s.Len()
	if length == 0 {
		var zero T
		return zero, ErrEmptyStack
	}
	return s.data[length-1], nil
}
