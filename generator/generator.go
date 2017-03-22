package generator

import (
	"github.com/luisfcofv/indexter/generator/salience"
	"github.com/luisfcofv/indexter/models"
)

type generator struct {
	Locations []models.Location `json:"locations"`
	Social    []models.Agent    `json:"social"`
}

var data generator

func Compute(world *models.World) {
	for index, event := range world.LatestEvents {
		spaceSalience := salience.SpaceSalience(world, event.Location)
		world.LatestEvents[index].Salience.Space = spaceSalience
	}
}
