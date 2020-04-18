package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// var count = 10

	// for count > 0 {
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second)
	// 	count--
	// }
	// fmt.Println("Liftoff!")

	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		var r = rand.Intn(100);
		fmt.Println("Chance: ",r)
		if r == 50 {
			fmt.Println("Rocket launch fails!")
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("Rocket launched successfully!")
	}

}