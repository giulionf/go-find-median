package medianofmedians

import (
	"math"
	"sort"
)

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

	sliceCount := (len(list) + 4) / 5
	medians := make([]int, sliceCount)

	for i := 0; i < len(list); i+=5 {
		if i + 5 < len(list) {
			slice := list[i:i+5]
			fastSort5(slice)
			medians[i/5] = slice[2]
		} else {
			slice := list[i:]
			sort.Ints(slice)
			medians[i/5] = slice[len(slice) / 2]
		}
	}
	pivot := findKthSmallestElement(medians, sliceCount/2)

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

func fastSort5(list []int) {
	for j0 := 0; j0 < 3; j0++ {
		jmin := j0
		for j := j0+1; j < 5; j++ {
			if list[j] < list[jmin] {
				jmin = j
			}
		}

		tmp := list[j0]
		list[j0] = list[jmin]
		list[jmin] = tmp
	}
}