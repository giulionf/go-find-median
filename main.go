package main

import (
	"fmt"
	"median-finding/algorithms"
	"median-finding/util"
	"time"
)

func main() {
	list := util.GenerateRandomList(10000000)
	fmt.Println("---- RANDOM LIST ----")
	//iteration(list)
	sortfirst(list)
	divideAndConquer(list)
	fast(list)

	list2 := util.GenerateRandomList(10000000)
	fmt.Println("\n---- INCREMENTING LIST ----")
	//iteration(list)
	sortfirst(list2)
	divideAndConquer(list2)
	fast(list2)
}

func iteration(list []int) {
	startTime := time.Now()
	algorithms.FindMedian(algorithms.RepeatedLinearSelection, list)
	elapsedTime := time.Since(startTime)
	fmt.Printf("%s executed within: %s\n", "Iteration", elapsedTime)
}

func sortfirst(list []int) {
	startTime := time.Now()
	algorithms.FindMedian(algorithms.SortFirst, list)
	elapsedTime := time.Since(startTime)
	fmt.Printf("%s executed within: %s\n", "SortFirst", elapsedTime)
}

func divideAndConquer(list []int) {
	startTime := time.Now()
	algorithms.FindMedian(algorithms.QuickSelect, list)
	elapsedTime := time.Since(startTime)
	fmt.Printf("%s executed within: %s\n", "Divide And Conquer", elapsedTime)
}

func fast(list []int) {
	startTime := time.Now()
	algorithms.FindMedian(algorithms.MedianOfMedians, list)
	elapsedTime := time.Since(startTime)
	fmt.Printf("%s executed within: %s\n", "MedianOfMedians", elapsedTime)
}