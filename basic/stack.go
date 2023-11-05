package basic

import "errors"

type Stack[V any] struct {
	s []V
}

func NewStack[V any]() *Stack[V] {
	return &Stack[V]{s: make([]V, 0)}
}

func (s *Stack[V]) Push(v V) {
	s.s = append(s.s, v)
}

func (s *Stack[V]) Pop() (v V, err error) {
	if s.IsEmpty() {
		return v, errors.New("empty stack")
	}
	v = s.s[s.Size()-1]
	s.s = s.s[:s.Size()-1]
	return v, nil
}

func (s *Stack[V]) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack[V]) Size() int {
	return len(s.s)
}
