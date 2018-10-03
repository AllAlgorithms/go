package stack

import (
	"errors"
	"reflect"
	"testing"
)

func newTestStack(initial, final int) *Stack {
	s := &Stack{}
	s.Init()

	for i := initial; i < final; i++ {
		s.l.Add(0, i)
	}

	return s
}

func TestStackInit(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		s := &Stack{}
		s.Init()

		expected := newTestStack(0, 0)
		if !reflect.DeepEqual(s, expected) {
			t.Errorf("Expected %+v, got %+v", expected, s)
		}
	})
}

func TestStackAdd(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		s := &Stack{}
		s.Init()
		for i := 0; i < 10; i++ {
			s.Push(i)
		}

		expected := newTestStack(0, 10)
		if !reflect.DeepEqual(expected, s) {
			t.Errorf("Expected %+v, got %+v", expected, s)
		}
	})

}

func TestStackSize(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		s := newTestStack(0, 0)
		size := s.Size()

		expected := 0
		if size != expected {
			t.Errorf("Expected %d, got %d", expected, size)
		}
	})

	t.Run("Non-empty stack", func(t *testing.T) {
		s := newTestStack(0, 10)
		size := s.Size()

		expected := 10
		if size != expected {
			t.Errorf("Expected %d, got %d", expected, size)
		}
	})
}

func TestStackIsEmpty(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		s := newTestStack(0, 0)
		isEmpty := s.IsEmpty()

		expected := true
		if isEmpty != expected {
			t.Errorf("Expected %t, got %t", expected, isEmpty)
		}
	})

	t.Run("Non-empty stack", func(t *testing.T) {
		s := newTestStack(0, 10)
		isEmpty := s.IsEmpty()

		expected := false
		if isEmpty != expected {
			t.Errorf("Expected %t, got %t", expected, isEmpty)
		}
	})
}

func TestStackPop(t *testing.T) {
	s := newTestStack(0, 10)

	t.Run("Popping from a non-empty stack", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			value, err := s.Pop()
			if err != nil {
				t.Errorf("No error expected, got %+v", err)
			}

			expected := 9 - i
			if value != expected {
				t.Errorf("Expected %+v, got %+v", expected, value)
			}

			expectedStack := newTestStack(0, 9-i)
			if !reflect.DeepEqual(expectedStack, s) {
				t.Errorf("Expected %+v, got %+v", expectedStack, s)
			}
		}
	})

	t.Run("Popping from an empty stack", func(t *testing.T) {
		_, err := s.Pop()
		if err == nil {
			t.Errorf("No error expected, got %+v", err)
		}

		expected := errors.New("stack - empty stack")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}
	})
}

func TestStackPeek(t *testing.T) {
	t.Run("Peeking from a non-empty stack", func(t *testing.T) {
		s := newTestStack(0, 10)
		value, err := s.Peek()
		if err != nil {
			t.Errorf("No error expected, got %+v", err)
		}

		expected := 9
		if value != expected {
			t.Errorf("Expected %+v, got %+v", expected, value)
		}
	})

	t.Run("Peeking from an empty stack", func(t *testing.T) {
		s := newTestStack(0, 0)
		_, err := s.Peek()
		if err == nil {
			t.Error("Error expected")
		}

		expected := errors.New("stack - empty stack")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}
	})
}

func TestStackString(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		s := newTestStack(0, 0)
		sString := s.String()

		expected := "[]"
		if sString != expected {
			t.Errorf("Expected %s, got %s", expected, sString)
		}
	})

	t.Run("Non-empty stack", func(t *testing.T) {
		s := newTestStack(0, 10)
		sString := s.String()

		expected := "[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]"
		if sString != expected {
			t.Errorf("Expected %s, got %s", expected, sString)
		}
	})
}
