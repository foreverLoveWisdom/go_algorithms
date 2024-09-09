package main

func maximizeDelta(values []int, isMin bool) int {
	if len(values) < 1 {
		return 0
	}
	extremum := values[0]
	maxDelta := 0
	var delta int
	for i := 1; i < len(values); i++ {
		if isMin {
			delta = values[i] - extremum
		} else {
			delta = extremum - values[i]
		}
		if delta > maxDelta {
			maxDelta = delta
		}
		if isMin && values[i] < extremum {
			extremum = values[i]
		}
		if !isMin && values[i] > extremum {
			extremum = values[i]
		}
	}
	return maxDelta
}
