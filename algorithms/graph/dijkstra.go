package dijkstra

import (
	"../datastructures/priorityqueue"
)

type Graph struct {
	Edges       [][]int
	NumOfVertex int
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

	paths := make([][]int, numOfVertex)
	queue := priorityqueue.PriorityQueue{}
	queue.Push(from, 0)

	for queue.Length() > 0 {
		node, cost := queue.Pop()
		if node == to {
			minCost = cost
			break
		}
		if isFixed[node] {
			continue
		}

		for i, edge := range graph.Edges[node] {
			if edge > 0 {
				queue.Push(i, edge+cost)
				paths[i] = append(paths[i], node)
			}
		}

		isFixed[node] = true
	}

	return minCost
}
