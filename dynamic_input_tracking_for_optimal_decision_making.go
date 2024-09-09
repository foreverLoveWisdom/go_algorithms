package main

func maximizeDelta(values []int, isMin bool) int {
	invalidValuesLength := 2
	if len(values) < invalidValuesLength {
		return 0
	}

	extremum := values[0]
	maxDelta := 0

	for _, value := range values[1:] {
		var delta int
		if isMin {
			delta = value - extremum
			if value < extremum {
				extremum = value
			}
		} else {
			delta = extremum - value
			if value > extremum {
				extremum = value
			}
		}

		if delta > maxDelta {
			maxDelta = delta
		}
	}

	return maxDelta
}
