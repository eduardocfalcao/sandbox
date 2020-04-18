package main

import "fmt"

/*
	Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis. Write a program
	to determine how fast a ship would need to travel (in km/h) in order to reach Malacandra in 28 days. Assume a distance of 56,000,000 km.
*/
func main () {
	const distance, days = 56000000, 28
	
	var speed = distance / (days * 24)

	fmt.Printf("To travel to Malacandra in 28 days needs a speed of %v km/h", speed)
}