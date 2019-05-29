package sortfirst

import (
	"math"
	"sort"
)

// Finds
func FindMedian(list []int) int {

	// Calculate the index of the median
	medianIndex := int(math.Floor(float64(len(list)) - 1.0) / 2)

	// Deep copy the list to preserve the original
	sortedList := make([]int, len(list))
	copy(sortedList, list)
	sort.Ints(sortedList)

	// Get the median by returning the middle
	median := sortedList[medianIndex]

	return median
}
