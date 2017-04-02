package salience

import (
	// "fmt"

	"github.com/luisfcofv/indexter/graph"
	"github.com/luisfcofv/indexter/models"
)

func TimeSalience(world *models.World, eventLocation models.Location, eventTime int) float64 {
	playerLocation := world.State.Player.Location

	g := graph.New()

	for _, location := range world.Locations {
		n := g.AddNode(location.ID)

		for _, neighbor := range location.Neighbors {
			n.AddEdge(neighbor.ID, neighbor.Time)
		}
	}

	g.ComputeShortestDistances(playerLocation)
	g.ShowShortestDistance(eventLocation.ID)
	return 0.0
}
