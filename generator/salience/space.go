package salience

import (
	"github.com/luisfcofv/indexter/models"
)

type node struct {
	ID       int
	Distance int
}

func SpaceSalience(locations []models.Location, activeNodes []int, experienceNodes []int) map[int]float32 {
	locationsMap := make(map[int][]models.Neighbor)

	for _, location := range locations {
		locationsMap[location.ID] = location.Neighbors
	}

	resultsMap := make(map[int]float32)
	for _, activeNodeID := range activeNodes {
		distance := closestNode(activeNodeID, locationsMap, experienceNodes)
		resultsMap[activeNodeID] = computeSalience(distance, len(locations))
	}

	return resultsMap
}

func computeSalience(distance int, totalNodes int) float32 {
	return 1.0 - float32(distance)/float32(totalNodes)
}

func closestNode(activeNodeID int, locationsMap map[int][]models.Neighbor, experienceNodes []int) int {
	queue := make([]node, 0)
	visitedNodes := make([]int, 0)

	// Push
	queue = append(queue, node{activeNodeID, 0})
	visitedNodes = append(visitedNodes, activeNodeID)

	for len(queue) >= 1 {
		currentNode := queue[0]
		// Discard top element
		queue = queue[1:]

		for _, experienceNode := range experienceNodes {
			if experienceNode == currentNode.ID {
				return currentNode.Distance
			}
		}

		for _, adjacentNode := range locationsMap[currentNode.ID] {
			visited := false
			for _, visitedNode := range visitedNodes {
				if visitedNode == adjacentNode.ID {
					visited = true
					break
				}
			}

			if !visited {
				queue = append(queue, node{adjacentNode.ID, currentNode.Distance + 1})
			}
		}
	}

	return len(locationsMap)
}
