package salience

import (
	"github.com/luisfcofv/indexter/graph"
	"github.com/luisfcofv/indexter/models"
)

func SpaceSalience(world *models.World, eventLocation models.Location) float64 {
	locationsMap := make(map[int][]int)

	for _, location := range world.Locations {
		var neighborIds []int
		for _, neighbor := range location.Neighbors {
			neighborIds = append(neighborIds, neighbor.ID)
		}
		locationsMap[location.ID] = neighborIds
	}

	distance := graph.BreadthFirstSearch(eventLocation.ID, locationsMap, world.Player.Knowledge.Locations)
	salience := computeSpaceSalience(distance, len(world.Locations))

	playerLocation := world.State.Player.Location
	if salience == 1.0 && eventLocation.ID != playerLocation {
		value := 1.0 / float64(len(world.Locations))
		salience -= value / 2.0
	}

	return salience
}

func computeSpaceSalience(distance int, totalNodes int) float64 {
	return 1.0 - float64(distance)/float64(totalNodes)
}
