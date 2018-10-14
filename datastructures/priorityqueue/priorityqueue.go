package priorityqueue

type NodeQueueItem struct {
	Node   int
	Weight int
}

type PriorityQueue struct {
	NodeItems []NodeQueueItem
}

// Push Node
func (queue *PriorityQueue) Push(node int, weight int) {
	nqi := NodeQueueItem{node, weight}
	if len(queue.NodeItems) == 0 {
		queue.NodeItems = append(queue.NodeItems, nqi)
		return
	}
	for i, v := range queue.NodeItems {
		if nqi.Weight < v.Weight {
			queue.NodeItems = append(queue.NodeItems[:i+1], queue.NodeItems[i:]...)
			queue.NodeItems[i] = nqi
			break
		} else {
			if i == len(queue.NodeItems)-1 {
				queue.NodeItems = append(queue.NodeItems, nqi)
			}
		}
	}
}

func (queue *PriorityQueue) Pop() (int, int) {
	if queue.Length() > 0 {
		node, weight := queue.Get(0)
		queue.NodeItems = queue.NodeItems[1:]
		return node, weight
	}
	return -1, -1
}

func (queue *PriorityQueue) Get(i int) (int, int) {
	return queue.NodeItems[i].Node, queue.NodeItems[i].Weight
}

func (queue *PriorityQueue) Length() int {
	return len(queue.NodeItems)
}
