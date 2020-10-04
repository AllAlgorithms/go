package tsp

import (
	"reflect"
	"testing"
)

func TestTSP(t *testing.T) {
	t.Run("TSP Test", func(t *testing.T) {
		graph := CreateGraph(4)

		graph.AppendEdge(0, 1, 6)
		graph.AppendEdge(0, 2, 5)
		graph.AppendEdge(0, 3, 5)
		graph.AppendEdge(1, 2, 4)
		graph.AppendEdge(1, 3, 7)
		graph.AppendEdge(2, 3, 3)

		expected := uint(18)
		actual := graph.TSP(0)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %+v, got %+v", expected, actual)
		}
	})
}
