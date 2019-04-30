package algorithms

import (
	"median-finding/algorithms/divideanconquer"
	"median-finding/algorithms/fast"
	"median-finding/algorithms/iterate"
	"median-finding/algorithms/sortfirst"
)

type Algorithm int

const (
	Iterate Algorithm = 1 + iota
	SortFirst
	DivideAndConquer
	Fast
)

func FindMedian(algorithm Algorithm, list []int) int {
	switch algorithm {

	case Iterate:
		return iterate.FindMedian(list)

	case SortFirst:
		return sortfirst.FindMedian(list)

	case DivideAndConquer:
		return divideanconquer.FindMedian(list)

	case Fast:
		return fast.FindMedian(list)

	default:
		return iterate.FindMedian(list)
	}
}
