package unionfindtree

type UnionFindTree struct {
	Parents []int
}

func Create(numOfNodes int) *UnionFindTree {
	uft := UnionFindTree{}
	uft.Parents = make([]int, numOfNodes)
	for i := 0; i < numOfNodes; i++ {
		uft.Parents[i] = i
	}
	return &uft
}

func (uft *UnionFindTree) Merge(x int, y int) {
	if uft.Root(x) == uft.Root(y) {
	} else {
		uft.Parents[y] = uft.Root(x)
	}
}

func (uft *UnionFindTree) Root(x int) int {
	Node := x
	for uft.Parents[Node] != Node {
		Node = uft.Parents[Node]
	}
	return Node
}

func (uft *UnionFindTree) IsSame(x int, y int) bool {
	return uft.Root(x) == uft.Root(y)
}
