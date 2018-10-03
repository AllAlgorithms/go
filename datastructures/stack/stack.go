package stack

import (
	"errors"

	"github.com/AllAlgorithms/go/datastructures/list"
)

// Stack a simple linked-list stack
type Stack struct {
	l *list.List
}

// Init inits an empty stack
func (s *Stack) Init() {
	l := &list.List{}
	l.Init()
	s.l = l
}

// IsEmpty returns if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.l.Size() == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return s.l.Size()
}

// Push adds an element in the top of the stack
func (s *Stack) Push(value interface{}) {
	s.l.Add(0, value)
}

// Pop removes and returns the top element in the stack
func (s *Stack) Pop() (interface{}, error) {
	v, e := s.l.Delete(0)
	if e != nil {
		return nil, errors.New("stack - empty stack")
	}
	return v, nil
}

// Peek returns the top element in the stack
func (s *Stack) Peek() (interface{}, error) {
	v, e := s.l.Get(0)
	if e != nil {
		return nil, errors.New("stack - empty stack")
	}
	return v, nil
}

func (s *Stack) String() string {
	return s.l.String()
}
