package main

import "fmt"

const (
	numberFormat = "%.1f"
)

type celsius float64
type fahrenheit float64

func (c celsius) toFahrenheit() fahrenheit {
	return fahrenheit(c*9/5) + 32
}

func (f fahrenheit) toCelsius() celsius {
	return celsius((f - 32)) * 5 / 9
}

func (c celsius) toString() string {
	return fmt.Sprintf(numberFormat, c)
}

func (f fahrenheit) toString() string {
	return fmt.Sprintf(numberFormat, f)
}

type getRowFn func(int) (string, string)

func drawRow(f, s string) {
	fmt.Printf("| %-9v| %-9v|\n", f, s)
}

func celsiusToFahenheit(v int) (string, string) {
	c := celsius(v)
	f := c.toFahrenheit()
	return c.toString(), f.toString()
}

func fahrenheitToCelsius(v int) (string, string) {
	f := fahrenheit(v)
	c := f.toCelsius()
	return f.toString(), c.toString()
}

func drawTable(header1, header2 string, getRow getRowFn) {
	fmt.Println("=======================")
	drawRow(header1, header2)
	fmt.Println("=======================")
	for temp := -40; temp <= 100; temp += 5 {
		cell1, cell2 := getRow(temp)
		drawRow(cell1, cell2)
	}

	fmt.Println("=======================")
}

func main() {
	drawTable("C°", "F°", celsiusToFahenheit)
	drawTable("°F", "°C", fahrenheitToCelsius)
}
