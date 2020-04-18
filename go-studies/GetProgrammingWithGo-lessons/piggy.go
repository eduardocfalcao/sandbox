package main

import (
	"fmt"
	"math/rand"
)

func main() {
	money := 0.0

	for money < 20.00 {
		switch rand.Intn(3) {
		case 0:
			money += 0.05
		case 1:
			money += 0.10
		case 2:
			money += 0.25
		}
		fmt.Printf("New balance: %5.2f\n", money)
	}
	fmt.Printf("Final balance: %05.2f\n", money)
}