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

	// Slice the list into groups of 5
	sliceCount := (len(list) + 4) / 5
	medians := make([]int, sliceCount)

	for i := 0; i < len(list); i+=5 {
		if i + 5 < len(list) {
			slice := list[i:i+5]
			insertionSort(slice)
			medians[i/5] = slice[2]
		} else {
			slice := list[i:]
			sort.Ints(slice)
			medians[i/5] = slice[len(slice) / 2]
		}
	}

	// Use the median of medians as pivot to ensure linear complexity
	pivot := findKthSmallestElement(medians, sliceCount/2)

	// Partition into smaller (left) and bigger (right) than the pivot
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < listLength; i++ {
		if list[i] < pivot {
			left = append(left, list[i])
		} else if list[i] > pivot {
			right = append(right, list[i])
		}
	}

	// Either recurse or return the median
	if k < len(left) {
		return findKthSmallestElement(left, k)
	} else if k > len(left) {
		return findKthSmallestElement(right, k - len(left) - 1)
	} else {
		return pivot
	}
}

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
