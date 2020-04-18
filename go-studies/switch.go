package main

import (
	"fmt"
)

func main() {
	switchTrue();
}

func defaultSwitch() {
	var nome string = "João"

	switch nome {
	case "Ana":
		fmt.Println("É a ana")
	case "João":
		fmt.Println("É o João")
	default:
		fmt.Println("Não conheço")
	}
}

func switchTrue() {
	var nota int = 4

	switch true { //pode tirar o true daqui tb, que significa a mesma coisa
	case nota > 9:
		fmt.Println("Ótimo")
	case nota > 7:
		fmt.Println("Muito bem")
	case nota > 6:
		fmt.Println("Bom")
	default:
		fmt.Println("Péssimo")
	}
}

