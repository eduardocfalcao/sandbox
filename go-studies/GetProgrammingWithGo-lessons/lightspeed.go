package main

import "fmt"

func main() {
	const speed = 100800 // km/h
	var distance = 96300000 //km

	var tripDays = (distance/speed) /24

	fmt.Printf("Will take %v Days to travel to mart at speed %v with distance %v", tripDays, speed, distance)
}