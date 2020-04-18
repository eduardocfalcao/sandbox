package main

import "fmt"

func main() {
	fmt.Println("My weight on the surface of Mars is ", 82 * 0.3783,
	" lbs, and I would be ", 29 * 365 / 687, " Years old." )

	fmt.Println("With Printf function: ")
	fmt.Printf("My weight on the surface of Mars is %v lbs, ", 82 * 0.3783)
	fmt.Printf("and I would be %v years old.\n", 29*365/687)

}