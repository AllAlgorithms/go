package dijkstra

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	t.Run("Dijkstra test", func(t *testing.T) {
		graph := CreateGraph(5)

		graph.AppendEdge(0, 1, 3)
		graph.AppendEdge(1, 3, 2)
		graph.AppendEdge(3, 4, 9)
		graph.AppendEdge(0, 2, 8)
		graph.AppendEdge(2, 4, 3)
		graph.AppendEdge(1, 2, 3)

		/*
			           3      2      9
					0 ---> 1 ---> 3 ---> 4
					|	   |             A
					|	   | 3           |
					|	   |             | 1
					|	   V             |
					L ---> 2 ----------> J
					   8          2
		*/

		expected := 9
		actual := graph.Dijkstra(0, 4)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %+v, actual %+v", expected, actual)
		}

	})

}
