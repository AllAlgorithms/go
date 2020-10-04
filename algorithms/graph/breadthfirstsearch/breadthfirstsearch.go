package breadthfirstsearch

// Node represents graph Node
type Node struct {
	Value  int
	leaves []*Node
}

// GetLeaves returns current Node leaves
func (n *Node) GetLeaves() []*Node {
	return n.leaves
}

// Insert adds a value to current Node
func (n *Node) Insert(value int) *Node {
	if n.leaves == nil {
		n.leaves = make([]*Node, 0)
	}

	newNode := new(Node)
	newNode.Value = value
	n.leaves = append(n.leaves, newNode)

	return n
}

// Search looks for given value recursively
// from current node through all its nodes
func (n *Node) Search(value int) bool {
	var current *Node
	queue := []*Node{n}

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if current.Value == value {
			return true
		}

		if current.GetLeaves() != nil {
			queue = append(queue, current.GetLeaves()...)
		}
	}

	return false
}
