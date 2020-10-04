package graphs

import (
	"math"
	"reflect"
	"testing"
)

func TestFloydWarshall(t *testing.T) {

	tests := []struct {
		name string
		args [][]float64
		want [][]float64
	}{
		// TODO: Add test cases.

		{
			name: "simple graph with 5 vertexes and negative edges",
			args: [][]float64{
				[]float64{0, 3, 8, math.Inf(1), -4},
				[]float64{math.Inf(1), 0, math.Inf(1), 1, 7},
				[]float64{math.Inf(1), 4, 0, math.Inf(1), math.Inf(1)},
				[]float64{2, math.Inf(1), -5, 0, math.Inf(1)},
				[]float64{math.Inf(1), math.Inf(1), math.Inf(1), 6, 0},
			},
			want: [][]float64{
				[]float64{0, 1, -3, 2, -4},
				[]float64{3, 0, -4, 1, -1},
				[]float64{7, 4, 0, 5, 3},
				[]float64{2, -1, -5, 0, -2},
				[]float64{8, 5, 1, 6, 0},
			},
		},
		{
			name: "simple graph with 4 vertexes and negative edges",
			args: [][]float64{
				[]float64{0, math.Inf(1), -2, math.Inf(1)},
				[]float64{4, 0, 3, math.Inf(1)},
				[]float64{math.Inf(1), math.Inf(1), 0, 2},
				[]float64{math.Inf(1), -1, math.Inf(1), 0},
			},
			want: [][]float64{
				[]float64{0, -1, -2, 0},
				[]float64{4, 0, 2, 4},
				[]float64{5, 1, 0, 2},
				[]float64{3, -1, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloydWarshall(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FloydWarshall() = %v, want %v", got, tt.want)
			}
		})
	}
}
