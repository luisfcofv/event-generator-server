package generator

import (
	"github.com/luisfcofv/indexter/generator/salience"
	"github.com/luisfcofv/indexter/graph"
	"github.com/luisfcofv/indexter/models"
)

func buildLocationGraph(world *models.World) *graph.Graph {
	playerLocation := world.State.Player.Location

	locationGraph := graph.New()

	for _, location := range world.Locations {
		node := locationGraph.AddNode(location.ID)

		for _, neighbor := range location.Neighbors {
			node.AddEdge(neighbor.ID, neighbor.Time)
		}
	}

	locationGraph.ComputeShortestDistances(playerLocation)
	return locationGraph
}

func Compute(world *models.World) {
	locationGraph := buildLocationGraph(world)
	var times []int
	longestTime := 0
	for index, event := range world.LatestEvents {
		spaceSalience := salience.SpaceSalience(world, event.Location)
		world.LatestEvents[index].Salience.Space = spaceSalience

		socialSalience := salience.SocialSalience(world, event.Agents)
		world.LatestEvents[index].Salience.Social = socialSalience

		intentionSalience := salience.IntentionSalience(world, event.Goal)
		world.LatestEvents[index].Salience.Intention = intentionSalience

		causationSalience := salience.CausationSalience(world, event.Cause)
		world.LatestEvents[index].Salience.Causation = causationSalience

		shortestTime := locationGraph.ShortestDistance(event.Location.ID)

		diffTime := event.Time - shortestTime
		times = append(times, diffTime)
		if diffTime >= longestTime {
			longestTime = diffTime
		}
	}

	for index, time := range times {
		timeSalience := salience.TimeSalience(time, longestTime)
		world.LatestEvents[index].Salience.Time = timeSalience
		world.LatestEvents[index].Salience.ComputeTotal()
	}
}
