package salience

import (
	"time"

	"github.com/luisfcofv/indexter/models"
)

func TimeSalience(world *models.World, eventTime int64) float64 {
	event := time.Unix(eventTime, 0)

	for _, playerTime := range world.Player.Knowledge.Times {
		start := time.Unix(playerTime.Start, 0)
		end := time.Unix(playerTime.End, 0)

		if event.After(start) == true && event.Before(end) {
			return 1.0
		}
	}

	return 0.0
}
