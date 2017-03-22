package salience

import "github.com/luisfcofv/indexter/models"

func SpaceSalience(world *models.World, eventLocation int) float32 {
	locationsMap := make(map[int][]int)

	for _, location := range world.Locations {
		locationsMap[location.ID] = location.Neighbors
	}

	distance := bfs(eventLocation, locationsMap, world.Player.Knowledge.Locations)
	salience := computeSalience(distance, len(world.Locations))

	playerLocation := world.State.Player.Location

	if salience == 1.0 && eventLocation != playerLocation {
		value := 1.0 / float32(len(world.Locations))
		salience -= value / 2
	}

	return salience
}
