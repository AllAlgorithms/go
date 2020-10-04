package priorityqueue

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	t.Run("Push test", func(t *testing.T) {
		pQueue := PriorityQueue{}

		item1 := NodeQueueItem{1, 3}
		item2 := NodeQueueItem{5, 1}
		item3 := NodeQueueItem{3, 8}
		pQueue.Push(1, 3)
		pQueue.Push(5, 1)
		pQueue.Push(3, 8)

		expectednode := item2.Node
		actualnode, _ := pQueue.Get(0)

		if !reflect.DeepEqual(expectednode, actualnode) {
			t.Errorf("Expected %+v, actual %+v", expectednode, actualnode)
		}

		expectednode = item1.Node
		actualnode, _ = pQueue.Get(1)

		if !reflect.DeepEqual(expectednode, actualnode) {
			t.Errorf("Expected %+v, actual %+v", expectednode, actualnode)
		}

		expectednode = item3.Node
		actualnode, _ = pQueue.Get(2)

		if !reflect.DeepEqual(expectednode, actualnode) {
			t.Errorf("Expected %+v, actual %+v", expectednode, actualnode)
		}

	})
}

func TestPop(t *testing.T) {
	t.Run("Length test", func(t *testing.T) {
		pQueue := PriorityQueue{}

		item1 := NodeQueueItem{1, 3}
		item2 := NodeQueueItem{5, 1}
		item3 := NodeQueueItem{3, 8}
		pQueue.Push(1, 3)
		pQueue.Push(5, 1)
		pQueue.Push(3, 8)

		expectednode := item2.Node
		actualnode, _ := pQueue.Pop()

		if !reflect.DeepEqual(expectednode, actualnode) {
			t.Errorf("Expected %+v, actual %+v", expectednode, actualnode)
		}

		expectednode = item1.Node
		actualnode, _ = pQueue.Pop()

		if !reflect.DeepEqual(expectednode, actualnode) {
			t.Errorf("Expected %+v, actual %+v", expectednode, actualnode)
		}

		expectednode = item3.Node
		actualnode, _ = pQueue.Pop()

		if !reflect.DeepEqual(expectednode, actualnode) {
			t.Errorf("Expected %+v, actual %+v", expectednode, actualnode)
		}

	})
}

func TestLength(t *testing.T) {
	t.Run("Length test", func(t *testing.T) {
		pQueue := PriorityQueue{}

		pQueue.Push(1, 3)

		expected := 1
		actual := pQueue.Length()

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %+v, actual %+v", expected, actual)
		}

		pQueue.Push(5, 1)
		pQueue.Push(3, 8)

		expected = 3
		actual = pQueue.Length()

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %+v, actual %+v", expected, actual)
		}

	})
}
