package main

import (
	"fmt"
	"median-finding/algorithms"
	"median-finding/util"
	"time"
)

func main() {
	list := util.GenerateRandomList(10000000)
	//iteration(list)
	sortfirst(list)
	divideAndConquer(list)
	fast(list)
}

func iteration(list []int) {
	startTime := time.Now()
	algorithms.FindMedian(algorithms.Iterate, list)
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
	algorithms.FindMedian(algorithms.DivideAndConquer, list)
	elapsedTime := time.Since(startTime)
	fmt.Printf("%s executed within: %s\n", "Divide And Conquer", elapsedTime)
}

func fast(list []int) {
	startTime := time.Now()
	algorithms.FindMedian(algorithms.Fast, list)
	elapsedTime := time.Since(startTime)
	fmt.Printf("%s executed within: %s\n", "Fast", elapsedTime)
}