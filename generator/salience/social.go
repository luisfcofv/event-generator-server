package salience

import (
	"math"

	"github.com/luisfcofv/indexter/models"
)

func SocialSalience(world *models.World, agents []models.Agent) float64 {
	socialMap := make(map[int][]int)

	for _, agent := range world.Agents {
		socialMap[agent.ID] = agent.Connections
	}

	distance := float64(len(world.Agents))
	for _, agent := range agents {
		if agent.ID == 0 {
			// 0 is the protagonist
			return 1.0
		}

		agentDistance := bfs(agent.ID, socialMap, world.Player.Knowledge.Social)
		distance = math.Min(distance, float64(agentDistance))
	}

	return computeSocialSalience(distance, len(world.Locations))
}

func computeSocialSalience(distance float64, totalNodes int) float64 {
	return math.Max(1.0-(distance+1.0)/float64(totalNodes), 0.0)
}
