package list

import (
	"errors"
	"reflect"
	"testing"
)

func TestListInit(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		l := &List{}
		l.Init()

		expected := &List{
			head: nil,
			size: 0,
		}
		if !reflect.DeepEqual(expected, l) {
			t.Errorf("Expected %+v, got %+v", expected, l)
		}
	})
}

func TestListAdd(t *testing.T) {
	l := &List{}

	t.Run("Adding elements at the beginning", func(t *testing.T) {
		l.Init()

		for i := 0; i < 5; i++ {
			err := l.Add(0, i)
			if err != nil {
				t.Errorf("No error expected, got %+v", err)
			}
		}
		expected := &List{
			head: &listNode{
				value: 4,
				next: &listNode{
					value: 3,
					next: &listNode{
						value: 2,
						next: &listNode{
							value: 1,
							next: &listNode{
								value: 0,
								next:  nil,
							},
						},
					},
				},
			},
			size: 5,
		}
		if !reflect.DeepEqual(expected, l) {
			t.Errorf("Expected %+v, got %+v", expected, l)
		}
	})

	t.Run("Adding elements at the end", func(t *testing.T) {
		l.Init()

		for i := 0; i < 5; i++ {
			err := l.Add(i, i)
			if err != nil {
				t.Errorf("No error expected, got %+v", err)
			}
		}
		expected := &List{
			head: &listNode{
				value: 0,
				next: &listNode{
					value: 1,
					next: &listNode{
						value: 2,
						next: &listNode{
							value: 3,
							next: &listNode{
								value: 4,
								next:  nil,
							},
						},
					},
				},
			},
			size: 5,
		}
		if !reflect.DeepEqual(expected, l) {
			t.Errorf("Expected %+v, got %+v", expected, l)
		}
	})

	t.Run("Adding elements at negative index", func(t *testing.T) {
		l.Init()

		err := l.Add(-1, 0)
		if err == nil {
			t.Error("Error expected")
		}

		expected := errors.New("list - index must be between 0 and the list size")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}
	})

	t.Run("Adding elements at bigger-than-size index", func(t *testing.T) {
		l.Init()

		err := l.Add(1, 0)
		if err == nil {
			t.Error("Error expected")
		}

		expected := errors.New("list - index must be between 0 and the list size")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}
	})
}

func newTestList(elements int) *List {
	l := &List{}
	l.Init()
	for i := 0; i < elements; i++ {
		l.Add(i, i)
	}

	return l
}

func TestListSize(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		l := newTestList(0)
		size := l.Size()

		expected := 0
		if size != expected {
			t.Errorf("Expected %d, got %d", expected, size)
		}
	})

	t.Run("Non-empty list", func(t *testing.T) {
		l := newTestList(10)
		size := l.Size()

		expected := 10
		if size != expected {
			t.Errorf("Expected %d, got %d", expected, size)
		}
	})
}

func TestListGet(t *testing.T) {
	l := newTestList(10)

	t.Run("Getting elements in valid indexes", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			value, err := l.Get(i)
			if err != nil {
				t.Errorf("No error expected, got %v", err)
			}

			expected := i
			if value != expected {
				t.Errorf("Expected %+v, got %+v", expected, value)
			}
		}
	})

	t.Run("Getting element in negative index", func(t *testing.T) {
		_, err := l.Get(-1)
		if err == nil {
			t.Error("Error expected")
		}

		expected := errors.New("list - index must be between 0 and the list size")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}

	})

	t.Run("Getting element in bigger-than-size index", func(t *testing.T) {
		_, err := l.Get(10)
		if err == nil {
			t.Error("Error expected")
		}

		expected := errors.New("list - index must be between 0 and the list size")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}

	})
}

func TestListDelete(t *testing.T) {
	l := newTestList(10)

	t.Run("Deleting elements in valid indexes", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			value, err := l.Delete(9 - i)
			if err != nil {
				t.Errorf("Expected no error, got %+v", err)
			}

			expected := 9 - i
			if expected != value {
				t.Errorf("Expected %+v, got %+v", expected, value)
			}
		}

		expected := newTestList(5)
		if !reflect.DeepEqual(expected, l) {
			t.Errorf("Expected %+v, got %+v", expected, l)
		}
	})

	t.Run("Deleting elements in initial index", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			value, err := l.Delete(0)
			if err != nil {
				t.Errorf("Expected no error, got %+v", err)
			}

			expected := i
			if expected != value {
				t.Errorf("Expected %+v, got %+v", expected, value)
			}
		}

		expected := &List{
			head: &listNode{
				value: 3,
				next: &listNode{
					value: 4,
					next:  nil,
				},
			},
			size: 2,
		}
		if !reflect.DeepEqual(expected, l) {
			t.Errorf("Expected %+v, got %+v", expected, l)
		}
	})

	t.Run("Deleting elements in negative index", func(t *testing.T) {
		_, err := l.Delete(-1)
		if err == nil {
			t.Error("Error expected")
		}

		expected := errors.New("list - index must be between 0 and the list size")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}
	})

	t.Run("Deleting elements in bigger-than-size index", func(t *testing.T) {
		_, err := l.Delete(10)
		if err == nil {
			t.Error("Error expected")
		}

		expected := errors.New("list - index must be between 0 and the list size")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}
	})
}

func TestListString(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		l := newTestList(0)
		lString := l.String()

		expected := "[]"
		if lString != expected {
			t.Errorf("Expected %s, got %s", expected, lString)
		}
	})

	t.Run("Non-empty list", func(t *testing.T) {
		l := newTestList(10)
		lString := l.String()

		expected := "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]"
		if lString != expected {
			t.Errorf("Expected %s, got %s", expected, lString)
		}
	})
}
