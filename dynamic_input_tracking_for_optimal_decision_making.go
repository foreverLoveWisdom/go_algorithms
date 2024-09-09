package main

const minimumLengthForComparison = 2

func maximizeDelta(values []int, trackMinimum bool) int {
	if notEnoughElementsToCompare(values) {
		return 0
	}

	currentExtremum := firstValue(values)
	largestDifference := 0

	for _, currentValue := range remainingValues(values) {
		largestDifference = updateLargestDifference(currentValue, currentExtremum, largestDifference, trackMinimum)
		currentExtremum = updateCurrentExtremum(currentValue, currentExtremum, trackMinimum)
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

// remainingValues retrieves all the values after the first one.
func remainingValues(values []int) []int {
	return values[1:]
}

// updateLargestDifference compares the current difference with the largest found so far and updates it if necessary.
func updateLargestDifference(currentValue, currentExtremum, largestDifference int, trackMinimum bool) int {
	difference := calculateDifference(currentValue, currentExtremum, trackMinimum)
	if difference > largestDifference {
		return difference
	}
	return largestDifference
}

// updateCurrentExtremum updates the current extremum (either minimum or maximum) based on the current value.
func updateCurrentExtremum(currentValue, currentExtremum int, trackMinimum bool) int {
	if shouldUpdateExtremum(currentValue, currentExtremum, trackMinimum) {
		return currentValue
	}
	return currentExtremum
}

// shouldUpdateExtremum determines if the current value should replace the current extremum.
func shouldUpdateExtremum(currentValue, currentExtremum int, trackMinimum bool) bool {
	if trackMinimum {
		return currentValue < currentExtremum
	}
	return currentValue > currentExtremum
}

// calculateDifference computes the difference based on whether we are tracking a minimum or maximum value.
func calculateDifference(currentValue, currentExtremum int, trackMinimum bool) int {
	if trackMinimum {
		return currentValue - currentExtremum
	}
	return currentExtremum - currentValue
}
