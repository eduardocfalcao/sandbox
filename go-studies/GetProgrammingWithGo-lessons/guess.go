package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const myNumber = 7
	var chances = 100

	for chances > 0 {
		var randNumber = rand.Intn(100) + 1
		fmt.Println("Number guess: ", randNumber)

		if randNumber == myNumber {
			fmt.Println("Exactly! This is my number")
			break
		} else if randNumber < myNumber {
			fmt.Println("My number ig bigget than that!")
		} else {
			fmt.Println("My number is minor than that!")
		}
		time.Sleep(time.Second)
		chances--
	}
}