package salience

import (
	"math"

	"github.com/luisfcofv/indexter/graph"
	"github.com/luisfcofv/indexter/models"
)

func SocialSalience(world *models.World, event models.Event) float64 {
	if event.Protagonist {
		return 1.0
	}

	socialMap := make(map[int][]int)
	for _, agent := range world.Agents {
		socialMap[agent.ID] = agent.Connections
	}

	distance := float64(len(world.Agents))
	for _, agent := range event.Agents {
		agentDistance := graph.BreadthFirstSearch(agent.ID, socialMap, world.Player.Knowledge.Social)
		distance = math.Min(distance, float64(agentDistance))
	}

	return computeSocialSalience(distance, len(world.Locations))
}

func computeSocialSalience(distance float64, totalNodes int) float64 {
	return math.Max(1.0-(distance+1.0)/float64(totalNodes), 0.0)
}
