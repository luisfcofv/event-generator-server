package salience

func TimeSalience(eventTime int, longestTime int) float64 {
	if eventTime < 0 {
		return 0.0
	} else if eventTime == 0 {
		return 1.0
	} else if eventTime == longestTime {
		return 0.1
	}

	return 1 - (float64(eventTime) * 1.0 / float64(longestTime))
}
