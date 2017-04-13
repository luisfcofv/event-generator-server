package salience

var limit = 10.0

func TimeSalience(shortestTime int, propagationTime int) float64 {
	if shortestTime == 0 {
		return 1.0
	}

	timeToReach := float64(shortestTime) / float64(propagationTime)
	salience := (limit - timeToReach) / limit
	return salience
}
