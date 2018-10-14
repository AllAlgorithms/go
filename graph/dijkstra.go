package dijkstra

type Graph struct {
	Edges        [][]int
	NumOfVertex  int
	SourceNode   int
	TerminalNode int
}

func CreateGraph(numOfVertex int) *Graph {
	graph := Graph{}
	graph.NumOfVertex = numOfVertex
	graph.Edges = make([][]int, numOfVertex)
	for i, _ := range graph.Edges {
		graph.Edges[i] = make([]int, numOfVertex)
	}
	return &graph
}

func (graph *Graph) AppendEdge(from int, to int, weight int) {
	graph.Edges[from][to] = weight
}

func (graph *Graph) Dijkstra(from int, to int) int {
	numOfVertex := graph.NumOfVertex
	isFixed := make([]bool, numOfVertex)
	minCost := 0
	vertexCosts := make([][]int, numOfVertex)

	node = from

	for 

	return minCost
}
