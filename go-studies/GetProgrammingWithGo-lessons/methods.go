package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) tofahrenheit() fahrenheit {
	return fahrenheit(c*9/5) + 32
}

func (k kelvin) tofahrenheit() fahrenheit {
	return fahrenheit((k-273.15)*9/5 + 32)
}

func main() {
	var c celsius = 25
	var k kelvin = 0
	fmt.Println("25°C in fahrenheit:", c.tofahrenheit())
	fmt.Println("0° Kelvin in fahrenheit", k.tofahrenheit())
}
