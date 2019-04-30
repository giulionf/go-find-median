package iterate

import (
	"math"
)

func FindMedian(list []int) int {
	// Calculate the index of the median
	medianIndex := int(math.Ceil(float64(len(list)) / 2) - 1)

	// Copy the list to leave the original intact
	copyList := make([]int, len(list))
	copy(copyList, list)

	// Remove the smallest number of the List for n (= medianIndex) times
	var median int
	for i := 0; i <= medianIndex; i++ {
		copyList, median = removeSmallestElement(copyList)
	}

	return median
}

func removeSmallestElement(list []int) ([]int, int) {
	// Find the index of the smallest number by iteration
	var indexOfSmallestNumber = 0
	for i := 1; i < len(list); i++ {
		if list[i] < list[indexOfSmallestNumber] {
			indexOfSmallestNumber = i
		}
	}

	// Copy the value of the found index
	smallestElement := list[indexOfSmallestNumber]

	// Remove it from the array
	list[indexOfSmallestNumber] = list[len(list)-1]
	list[len(list)-1] = 0
	list = list[:len(list)-1]

	return list, smallestElement
}
