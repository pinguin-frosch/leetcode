package main

import "fmt"

func main() {
	s := "([[}]])"
	valid := isValid(s)
	fmt.Println(valid)

}

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

func (s *Stack[T]) Pop() T {
	length := s.Len()
	if length == 0 {
		panic("cannot pop, empty stack")
	}
	obj := s.data[length-1]
	s.data = s.data[:length-1]
	return obj
}

func (s *Stack[T]) Seek() T {
	length := s.Len()
	if length == 0 {
		panic("cannot seek, empty stack")
	}
	return s.data[length-1]
}

func isValid(s string) bool {
	stack := NewStack[rune]()
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
		} else {
			if stack.Len() == 0 {
				return false
			}
			obj := stack.Pop()
			if c == ']' && obj != '[' {
				return false
			}
			if c == ')' && obj != '(' {
				return false
			}
			if c == '}' && obj != '{' {
				return false
			}
		}
	}
	return stack.Len() == 0
}
