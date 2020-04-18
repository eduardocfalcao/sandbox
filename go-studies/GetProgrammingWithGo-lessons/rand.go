package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)


	fmt.Println("Check 2.6:")

	const (
		minDistance = 56000000
		maxDistance = 401000000 - minDistance
	)

	num = rand.Intn(maxDistance) + minDistance

	fmt.Println(num)
}