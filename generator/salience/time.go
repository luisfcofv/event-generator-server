package salience

import (
	"github.com/luisfcofv/indexter/graph"
	"github.com/luisfcofv/indexter/models"
)

func TimeSalience(locationGraph *graph.Graph, eventLocation models.Location, eventTime int) float64 {
	shortestDistance := locationGraph.ShortestDistance(eventLocation.ID)
	if shortestDistance <= eventTime {
		return 1.0
	}

	return 0.0
}
