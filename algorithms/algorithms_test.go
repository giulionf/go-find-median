package algorithms

import (
	"testing"
)

func TestSimpleMedian(t *testing.T) {
	testMedian(t, Iterate)
}

func TestSortedMedian(t *testing.T) {
	testMedian(t, SortFirst)
}

func TestDivideAndConquer(t *testing.T) {
	testMedian(t, DivideAndConquer)
}

func TestFast(t *testing.T) {
	testMedian(t, Fast)
}

func testMedian(t *testing.T, algorithm Algorithm) {
	list1 := []float64 { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list2 := []float64 { 9, 1, 8, 2, 7, 3, 6, 4, 5, 0, 10}
	list3 := []float64 { 9, 1, 5, 3, 4, 2, 8, 7, 0, 6, 10}
	list4 := []float64 { 10, 20, 30, 100, 200, 300, 1000, 100000}
	list5 := []float64 { 10.1, 20.2, 30, 100.5, 200.3, 300.1, 1000.1123, 100000, 133123}

	median1 := FindMedian(algorithm, list1)
	median2 := FindMedian(algorithm, list2)
	median3 := FindMedian(algorithm, list3)
	median4 := FindMedian(algorithm, list4)
	median5 := FindMedian(algorithm, list5)

	if median1 != median2 || median2 != median3 {
		t.Error("Different results: ", median1, median2, median3)
	}

	if median1 != 5 {
		t.Error("Wrong Median: ", median1)
	}

	if median4 != 100 {
		t.Error("Wrong Median: ", median4)
	}

	if median5 != 200.3 {
		t.Error("Wrong Median: ", median5)
	}
}