package fast

import (
	"math"
	"sort"
)

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

	sublists := make([][]int, 0)
	for i := 0; i < len(list); i+=5 {
		slice := list[i:min(len(list), i+5)]
		sublists = append(sublists, slice)
	}

	medians := make([]int, len(sublists))
	for i := 0; i < len(sublists); i++ {
		sort.Ints(sublists[i])
		medians[i] = sublists[i][len(sublists[i])/2]
	}

	var pivot int

	if len(medians) <= 5 {
		sort.Ints(medians)
		pivot = medians[len(medians)/2]
	} else {
		pivot = medianOfMedians(medians, len(medians)/2)
	}

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

func min(b int, a int) int {
	if a < b {
		return a
	} else {
		return b
	}
}