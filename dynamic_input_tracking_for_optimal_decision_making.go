package main

// maximizeDelta calculates the maximum difference between elements in a sequence,
// based on whether we are tracking a minimum or maximum value.
func maximizeDelta(values []int, isMin bool) int {
	if isTooShort(values) {
		return 0
	}

	extremum := initializeExtremum(values)
	maxDelta := 0

	for i := 1; i < len(values); i++ {
		value := values[i]
		maxDelta = updateMaxDelta(value, extremum, maxDelta, isMin)
		extremum = updateExtremum(value, extremum, isMin)
	}

	return maxDelta
}

// isTooShort checks if the values slice is too short to compute a meaningful delta.
func isTooShort(values []int) bool {
	minValidLength := 2
	return len(values) < minValidLength
}

// initializeExtremum initializes the extremum with the first value in the slice.
func initializeExtremum(values []int) int {
	return values[0]
}

// updateMaxDelta calculates the delta based on whether we are tracking a min or max.
func updateMaxDelta(value, extremum, maxDelta int, isMin bool) int {
	delta := calculateDelta(value, extremum, isMin)
	if delta > maxDelta {
		return delta
	}
	return maxDelta
}

// updateExtremum updates the extremum value depending on whether we are tracking min or max.
func updateExtremum(value, extremum int, isMin bool) int {
	if shouldUpdateExtremum(value, extremum, isMin) {
		return value
	}
	return extremum
}

// shouldUpdateExtremum checks whether the current value should update the extremum.
func shouldUpdateExtremum(value, extremum int, isMin bool) bool {
	if isMin {
		return value < extremum
	}
	return value > extremum
}

// calculateDelta computes the difference (delta) based on whether we are tracking a min or max.
func calculateDelta(value, extremum int, isMin bool) int {
	if isMin {
		return value - extremum
	}
	return extremum - value
}
