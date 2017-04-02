package salience

import "github.com/luisfcofv/indexter/models"

func IntentionSalience(world *models.World, eventGoal models.Goal) float64 {
	for _, goal := range world.Player.Knowledge.Goals {
		if eventGoal.ID == goal {
			return 1.0
		}
	}

	return 0.0
}
