package queue

import (
	"errors"
	"reflect"
	"testing"
)

func newTestQueue(initial, final int) *Queue {
	q := &Queue{}
	q.Init()

	for i := initial; i < final; i++ {
		q.r.Push(i)
	}

	return q
}

func newTestModifiedQueue(initial, final int) *Queue {
	q := &Queue{}
	q.Init()

	for i := initial; i < final; i++ {
		q.s.Push(final - i + initial)
	}

	return q
}

func TestQueueInit(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		q := &Queue{}
		q.Init()

		expected := newTestQueue(0, 0)
		if !reflect.DeepEqual(expected, q) {
			t.Errorf("Expected %+v, got %+v", expected, q)
		}
	})
}

func TestQueueAdd(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		q := &Queue{}
		q.Init()

		for i := 0; i < 10; i++ {
			q.Enqueue(i)
		}

		expected := newTestQueue(0, 10)
		if !reflect.DeepEqual(expected, q) {
			t.Errorf("Expected %+v, got %+v", expected, q)
		}
	})
}

func TestQueueSize(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		q := newTestQueue(0, 0)
		size := q.Size()

		expected := 0
		if size != expected {
			t.Errorf("Expected %d, got %d", expected, size)
		}
	})

	t.Run("Non-empty stack", func(t *testing.T) {
		q := newTestQueue(0, 10)
		size := q.Size()

		expected := 10
		if size != expected {
			t.Errorf("Expected %d, got %d", expected, size)
		}
	})
}

func TestQueueIsEmpty(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		q := newTestQueue(0, 0)
		isEmpty := q.IsEmpty()

		expected := true
		if isEmpty != expected {
			t.Errorf("Expected %t, got %t", expected, isEmpty)
		}
	})

	t.Run("Non-empty stack", func(t *testing.T) {
		q := newTestQueue(0, 10)
		isEmpty := q.IsEmpty()

		expected := false
		if isEmpty != expected {
			t.Errorf("Expected %t, got %t", expected, isEmpty)
		}
	})
}

func TestQueueDequeue(t *testing.T) {
	q := newTestQueue(0, 10)

	t.Run("Popping from a non-empty stack", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			value, err := q.Dequeue()
			if err != nil {
				t.Errorf("No error expected, got %+v", err)
			}

			expected := i
			if value != expected {
				t.Errorf("Expected %+v, got %+v", expected, value)
			}

			expectedQueue := newTestModifiedQueue(i, 9)
			if !reflect.DeepEqual(expectedQueue, q) {
				t.Errorf("Expected %+v, got %+v", expectedQueue, q)
			}
		}
	})

	t.Run("Popping from an empty stack", func(t *testing.T) {
		_, err := q.Dequeue()
		if err == nil {
			t.Errorf("No error expected, got %+v", err)
		}

		expected := errors.New("queue - empty queue")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}
	})
}

func TestQueuePeek(t *testing.T) {
	t.Run("Peeking elements from a non-empty queue", func(t *testing.T) {
		q := newTestQueue(0, 10)
		value, err := q.Peek()
		if err != nil {
			t.Errorf("No error expected, got %+v", err)
		}

		expected := 0
		if !reflect.DeepEqual(expected, value) {
			t.Errorf("Expected %+v, got %+v", expected, value)
		}
	})

	t.Run("Peeking elements from an empty queue", func(t *testing.T) {
		q := newTestQueue(0, 0)
		_, err := q.Peek()
		if err == nil {
			t.Error("Error expected")
		}

		expected := errors.New("queue - empty queue")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}
	})
}

func TestQueueString(t *testing.T) {
	t.Run("Empty queue", func(t *testing.T) {
		q := newTestQueue(0, 0)
		qString := q.String()

		expected := "[]"
		if qString != expected {
			t.Errorf("Expected %s, got %s", expected, qString)
		}
	})

	t.Run("Non-empty queue", func(t *testing.T) {
		q := newTestQueue(0, 10)
		qString := q.String()

		expected := "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]"
		if qString != expected {
			t.Errorf("Expected %s, got %s", expected, qString)
		}
	})

	t.Run("Non-empty queue after operations", func(t *testing.T) {
		q := newTestQueue(0, 10)
		q.Dequeue()
		q.Dequeue()
		qString := q.String()

		expected := "[2, 3, 4, 5, 6, 7, 8, 9]"
		if qString != expected {
			t.Errorf("Expected %s, got %s", expected, qString)
		}

		q.Enqueue(10)
		qString = q.String()
		expected = "[2, 3, 4, 5, 6, 7, 8, 9, 10]"
		if qString != expected {
			t.Errorf("Expected %s, got %s", expected, qString)
		}
	})
}
