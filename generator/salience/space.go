package salience

import (
	"fmt"

	"github.com/luisfcofv/indexter/models"
)

type node struct {
	ID       string
	Distance int
}

func SpaceSalience(locations []models.Location, activeNodes []string, experienceNodes []string) {
	locationsMap := make(map[string][]models.Neighbor)

	for _, location := range locations {
		locationsMap[location.ID] = location.Neighbors
	}

	for _, activeNodeID := range activeNodes {
		distance := closestNode(activeNodeID, locationsMap, experienceNodes)
		fmt.Println("Solution: ", activeNodeID, distance)
	}
}

func closestNode(activeNodeID string, locationsMap map[string][]models.Neighbor, experienceNodes []string) int {
	queue := make([]node, 0)
	visitedNodes := make([]string, 0)

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
