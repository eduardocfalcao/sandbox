package main

import (
	"fmt"
)

func main() {
	var numeros = [7]int { 1, 1, 2, 3, 5, 8, 13 }

	//Slice
	slice := numeros[2:5]

	// vai pegar do indice 2 até o final
	slice2 := numeros[2:]

	// vai do inicio até o indice (-1)
	slice3 := numeros[:5]

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	
	slices2()
}

func slices2() {
	var nomes = [3]string {
		"Ana",
		"Jose",
		"Maria",
	}

	var slice []string = nomes[0:2]
	slice[0] = "Rogerio"
	fmt.Println(slice)
	fmt.Println("Array Original:", nomes) 
	// vai ser alterado o valor no primeiro indice dessa array original
	// pois gerar slice dessa forma é criada uma referência.
}