package algorithms

import (
	"median-finding/algorithms/quickselect"
	"median-finding/algorithms/medianofmedians"
	"median-finding/algorithms/iterate"
	"median-finding/algorithms/sortfirst"
)

type Algorithm int

const (
	RepeatedLinearSelection Algorithm = 1 + iota
	SortFirst
	QuickSelect
	MedianOfMedians
)

func FindMedian(algorithm Algorithm, list []int) int {
	switch algorithm {

	case RepeatedLinearSelection:
		return iterate.FindMedian(list)

	case SortFirst:
		return sortfirst.FindMedian(list)

	case QuickSelect:
		return quickselect.FindMedian(list)

	case MedianOfMedians:
		return medianofmedians.FindMedian(list)

	default:
		return iterate.FindMedian(list)
	}
}
