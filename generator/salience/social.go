package salience

import (
	"math"

	"github.com/luisfcofv/indexter/models"
)

func SocialSalience(world *models.World, eventAgent int) float64 {
	socialMap := make(map[int][]int)

	for _, agent := range world.Agents {
		socialMap[agent.ID] = agent.Connections
	}

	if eventAgent == 0 {
		// 0 is the protagonist
		return 1.0
	}

	distance := bfs(eventAgent, socialMap, world.Player.Knowledge.Social)
	return computeSocialSalience(distance, len(world.Locations))
}

func computeSocialSalience(distance int, totalNodes int) float64 {
	return math.Max(1.0-float64(distance+1)/float64(totalNodes), 0.0)
}
