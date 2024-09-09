package main

const minimumLengthForComparison = 2

func maximizeDelta(values []int, trackMinimum bool) int {
	if notEnoughElementsToCompare(values) {
		return 0
	}

	currentExtremum := firstValue(values)
	largestDifference := 0

	for _, currentValue := range remainingValues(values) {
		difference := calculateDifference(currentValue, currentExtremum, trackMinimum)
		if difference > largestDifference {
			largestDifference = difference
		}

		if shouldUpdateExtremum(currentValue, currentExtremum, trackMinimum) {
			currentExtremum = currentValue
		}
	}

	return largestDifference
}

// notEnoughElementsToCompare checks if the values slice is too short to find any meaningful differences.
func notEnoughElementsToCompare(values []int) bool {
	return len(values) < minimumLengthForComparison
}

// firstValue retrieves the first value in the sequence.
func firstValue(values []int) int {
	return values[0]
}

// remainingValues retrieves all the values after the first one for comparison purposes.
func remainingValues(values []int) []int {
	return values[1:]
}

// shouldUpdateExtremum checks whether the current value should replace the current extremum value.
func shouldUpdateExtremum(currentValue, currentExtremum int, trackMinimum bool) bool {
	// Simplified decision on whether to update based on min or max tracking
	if trackMinimum {
		return currentValue < currentExtremum
	}
	return currentValue > currentExtremum
}

// calculateDifference computes the difference between the current value and the extremum,
// depending on whether we are tracking the minimum or maximum.
func calculateDifference(currentValue, currentExtremum int, trackMinimum bool) int {
	if trackMinimum {
		return currentValue - currentExtremum
	}
	return currentExtremum - currentValue
}
