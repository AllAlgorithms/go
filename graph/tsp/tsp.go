package tsp

type Graph struct {
	Edges [][]int
	N     uint
}

func CreateGraph(n uint) *Graph {
	graph := Graph{}
	graph.N = n
	graph.Edges = make([][]int, n)
	for i := 0; i < int(n); i++ {
		graph.Edges[i] = make([]int, n)
	}
	return &graph
}

func (graph *Graph) AppendEdge(from int, to int, weight int) {
	graph.Edges[from][to] = weight
	graph.Edges[to][from] = weight
}

var dp [][]int

func (graph *Graph) TSP(from uint) uint {
	dp = make([][]int, 1<<graph.N)
	for i := 0; i < 1<<graph.N; i++ {
		dp[i] = make([]int, graph.N)
		for j := 0; j < int(graph.N); j++ {
			dp[i][j] = -1
		}
	}

	return graph._tsp(uint(0), from)

}

func (graph *Graph) _tsp(visited uint, stayed uint) uint {
	// Visited Nodes are represented with bits.
	if dp[visited][stayed] >= 0 {
		return uint(dp[visited][stayed])
	}
	if (visited == 1<<graph.N-1) && stayed == 0 {
		return 0
	}

	var minValue uint
	minValue = 10000
	var i uint
	for i = 0; i < graph.N; i++ {
		if (visited >> i & 1) == 0 {
			if graph.Edges[stayed][i] > 0 {
				value := graph._tsp(visited|1<<i, i) + uint(graph.Edges[stayed][i])
				if value < minValue {
					minValue = value
				}
			}
		}
	}

	dp[visited][stayed] = int(minValue)
	return uint(dp[visited][stayed])
}
