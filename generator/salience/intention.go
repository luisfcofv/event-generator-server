package salience

import "github.com/luisfcofv/indexter/models"

func IntentionSalience(world *models.World, eventGoal int) float64 {
	for _, goal := range world.Player.Knowledge.Goals {
		if eventGoal == goal {
			return 1.0
		}
	}

	return 0.0
}
