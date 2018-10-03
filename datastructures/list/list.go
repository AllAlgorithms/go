package list

import (
	"errors"
	"fmt"
)

type listNode struct {
	value interface{}
	next  *listNode
}

// List simple linked list
type List struct {
	head *listNode
	size int
}

// Init inits an empty list
func (l *List) Init() {
	l.head = nil
	l.size = 0
}

// Size returns the number of items in the list
func (l *List) Size() int {
	return l.size
}

// Add adds an item to the list in given index
func (l *List) Add(index int, value interface{}) error {
	if index < 0 || index > l.size {
		return errors.New("list - index must be between 0 and the list size")
	} else if index == 0 {
		node := &listNode{value, l.head}
		l.head = node
		l.size++
	} else {
		t := l.head
		for i := 0; i < index-1; i++ {
			t = t.next
		}
		n := &listNode{value, t.next}
		t.next = n
		l.size++
	}
	return nil
}

// Get returns the item in given index
func (l *List) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("list - index must be between 0 and the list size")
	}

	t := l.head
	for i := 0; i < index; i++ {
		t = t.next
	}
	return t.value, nil
}

// Delete removes and returns the item in given index
func (l *List) Delete(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("list - index must be between 0 and the list size")
	}
	if index == 0 {
		v := l.head.value
		l.head = l.head.next
		l.size--
		return v, nil
	}
	t := l.head
	for i := 0; i < index-1; i++ {
		t = t.next
	}
	v := t.next.value
	t.next = t.next.next
	l.size--
	return v, nil
}

func (l *List) String() string {
	if l.size == 0 {
		return "[]"
	}
	res := "["
	t := l.head
	for i := 0; i < l.size-1; i++ {
		res += fmt.Sprintf("%v", t.value) + ", "
		t = t.next
	}
	return res + fmt.Sprintf("%v", t.value) + "]"
}
