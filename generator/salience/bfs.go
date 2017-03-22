package salience

type node struct {
	ID       int
	Distance int
}

func bfs(locationID int, nodesMap map[int][]int, experienceNodes []int) int {
	queue := make([]node, 0)
	visitedNodes := make([]int, 0)

	// Push
	queue = append(queue, node{locationID, 0})
	visitedNodes = append(visitedNodes, locationID)

	for len(queue) >= 1 {
		currentNode := queue[0]
		// Discard top element
		queue = queue[1:]

		for _, experienceNode := range experienceNodes {
			if experienceNode == currentNode.ID {
				return currentNode.Distance
			}
		}

		for _, adjacentNode := range nodesMap[currentNode.ID] {
			visited := false
			for _, visitedNode := range visitedNodes {
				if visitedNode == adjacentNode {
					visited = true
					break
				}
			}

			if !visited {
				queue = append(queue, node{adjacentNode, currentNode.Distance + 1})
				visitedNodes = append(visitedNodes, adjacentNode)
			}
		}
	}

	return len(nodesMap)
}
