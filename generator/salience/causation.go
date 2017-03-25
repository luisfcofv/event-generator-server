package salience

import (
	"github.com/luisfcofv/indexter/models"
)

func CausationSalience(world *models.World, eventCause int) float64 {
	for _, goal := range world.Player.Knowledge.Goals {
		if eventCause == goal {
			return 0.0
		}
	}

	return 1.0
}
