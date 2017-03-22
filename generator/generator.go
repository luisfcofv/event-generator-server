package generator

import (
	"github.com/luisfcofv/indexter/generator/salience"
	"github.com/luisfcofv/indexter/models"
)

func Compute(world *models.World) {
	for index, event := range world.LatestEvents {
		spaceSalience := salience.SpaceSalience(world, event.Location)
		world.LatestEvents[index].Salience.Space = spaceSalience

		socialSalience := salience.SocialSalience(world, event.Agent)
		world.LatestEvents[index].Salience.Social = socialSalience
	}
}
