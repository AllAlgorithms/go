package graphs

//FloydWarshall is an algorithm for finding shortest paths in a weighted graph (with no negative cycles)
//The algorithm accepts graph represented as an adjacency matrix
//and returns all the shortest paths in graph
func FloydWarshall(graph [][]float64) [][]float64 {
	vertexesCount := len(graph)
	//copy graph
	paths := make([][]float64, vertexesCount)
	for i := 0; i < vertexesCount; i++ {
		paths[i] = make([]float64, vertexesCount)
		copy(paths[i], graph[i])
	}

	//count shortest paths
	for k := 0; k < vertexesCount; k++ {
		for i := 0; i < vertexesCount; i++ {
			for j := 0; j < vertexesCount; j++ {
				if paths[i][j] > paths[i][k]+paths[k][j] {
					paths[i][j] = paths[i][k] + paths[k][j]
				}
			}
		}
	}
	return paths
}
