package divideanconquer

import "math"

func FindMedian(list []int) int {
	// Calculate the index of the median
	medianIndex := int(math.Ceil(float64(len(list)) / 2) - 1)

	// Start the recursion
	return medianOfMedians(list, medianIndex)
}

func medianOfMedians(list []int, wantedIndex int) int {
	if len(list) == 1 {
		return list[0]
	}

	pivot := list[0]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(list); i++ {
		if list[i] < pivot {
			left = append(left, list[i])
		} else if list[i] > pivot {
			right = append(right, list[i])
		}
	}

	if wantedIndex < len(left) {
		return medianOfMedians(left, wantedIndex)
	} else if wantedIndex > len(left) {
		return medianOfMedians(right, wantedIndex - len(left) - 1)
	} else {
		return pivot
	}
}
