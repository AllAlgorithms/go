package queue

import (
	"errors"

	"github.com/AllAlgorithms/go/datastructures/stack"
)

// Queue a two-stack queue
type Queue struct {
	r *stack.Stack
	s *stack.Stack
}

// Init inits an empty queue
func (q *Queue) Init() {
	r := &stack.Stack{}
	s := &stack.Stack{}
	r.Init()
	s.Init()
	q.r = r
	q.s = s
}

// IsEmpty returns if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.r.IsEmpty() && q.s.IsEmpty()
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return q.r.Size() + q.s.Size()
}

// Enqueue adds an element in the back of the queue
func (q *Queue) Enqueue(value interface{}) {
	for !q.s.IsEmpty() {
		v, _ := q.s.Pop()
		q.r.Push(v)
	}
	q.r.Push(value)
}

// Dequeue removes and returns the front element in the queue
func (q *Queue) Dequeue() (interface{}, error) {
	for !q.r.IsEmpty() {
		v, _ := q.r.Pop()
		q.s.Push(v)
	}

	value, err := q.s.Pop()
	if err != nil {
		return nil, errors.New("queue - empty queue")
	}

	return value, nil
}

// Peek returns the front element in the queue
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue - empty queue")
	}

	for !q.r.IsEmpty() {
		v, _ := q.r.Pop()
		q.s.Push(v)
	}
	return q.s.Peek()
}

func (q *Queue) String() string {
	if q.IsEmpty() {
		return "[]"
	}
	if q.r.IsEmpty() {
		return q.s.String()
	}

	t := &stack.Stack{}
	t.Init()
	for !q.r.IsEmpty() {
		v, _ := q.r.Pop()
		t.Push(v)
	}

	res := t.String()
	for !t.IsEmpty() {
		v, _ := t.Pop()
		q.r.Push(v)
	}

	return res
}
