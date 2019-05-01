package quickselect

import "math"

func FindMedian(list []int) int {
	// Calculate the index of the median
	medianIndex := int(math.Ceil(float64(len(list)) / 2) - 1)

	// Start the recursion
	return findKthSmallestElement(list, medianIndex)
}

func findKthSmallestElement(list []int, k int) int {
	listLength := len(list)
	if listLength == 1 {
		return list[0]
	}

	pivot := list[0]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < listLength; i++ {
		if list[i] < pivot {
			left = append(left, list[i])
		} else if list[i] > pivot {
			right = append(right, list[i])
		}
	}

	if k < len(left) {
		return findKthSmallestElement(left, k)
	} else if k > len(left) {
		return findKthSmallestElement(right, k - len(left) - 1)
	} else {
		return pivot
	}
}
