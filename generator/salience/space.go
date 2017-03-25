package salience

import "github.com/luisfcofv/indexter/models"

func SpaceSalience(world *models.World, eventLocation int) float64 {
	locationsMap := make(map[int][]int)

	for _, location := range world.Locations {
		locationsMap[location.ID] = location.Neighbors
	}

	distance := bfs(eventLocation, locationsMap, world.Player.Knowledge.Locations)
	salience := computeSpaceSalience(distance, len(world.Locations))

	playerLocation := world.State.Player.Location
	if salience == 1.0 && eventLocation != playerLocation {
		value := 1.0 / float64(len(world.Locations))
		salience -= value / 2.0
	}

	return salience
}

func computeSpaceSalience(distance int, totalNodes int) float64 {
	return 1.0 - float64(distance)/float64(totalNodes)
}
