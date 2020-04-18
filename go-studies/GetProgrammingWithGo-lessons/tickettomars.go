package main

import (
	"fmt"
	"math/rand"
)

const secondsPerDay = 86400 // 60 * 60 * 24

func main () {
	distance := 62100000
	company := ""
	trip := ""

	fmt.Println("Spaceline           Days Trip type   Price")
	fmt.Println("==========================================")

	for i:= 0; i < 10; i++ {
		
		
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventurers"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		randSpeed := rand.Intn(15) 
		speed := randSpeed + 16
		tripPrice := speed + 20
		days := distance / speed / secondsPerDay 
		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			tripPrice *= 2
		} else {
			trip = "One-way"
		}

		fmt.Printf("%-20v %3v %-11v $ %3v\n", company, days, trip, tripPrice)
	}
	
}