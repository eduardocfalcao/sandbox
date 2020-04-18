package main

import (
	"fmt"
)

func main () {
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365
	const canisDistance = 236e15

	const distanceInLightYears  = canisDistance / lightSpeed / secondsPerDay / daysPerYear

	fmt.Println("The CanisMajor Dwarf is", distanceInLightYears, "light years from the Earth")
}