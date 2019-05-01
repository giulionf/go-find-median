package util

import (
	"math/rand"
	"time"
)

func GenerateRandomList(length int) []int {

	// Set time as random seed source
	s1 := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(s1)

	// Set up a list of random numbers
	var list []int
	for i := 0; i < length; i++ {
		list = append(list, rng.Int())
	}

	return list
}

func GenerateIncrementingList(length int) []int {

	// Set up a list of random numbers
	var list []int
	for i := 0; i < length; i++ {
		list = append(list, i)
	}

	return list
}
