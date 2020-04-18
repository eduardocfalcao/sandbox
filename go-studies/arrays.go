package main

import (
	"fmt"
)

func main() {
	acessandoElementoArray()
}
/*
	As arrays tem tamanho fixo. Os elementos são do mesmo tipo (igual outras linguages tipadas)
*/
func arrayPadrao() {
	var numeros [7]int

	numeros[0] = 1
	numeros[1] = 1
	numeros[2] = 2
	numeros[3] = 3
	numeros[4] = 5
	numeros[5] = 8
	numeros[6] = 13

	fmt.Println(numeros)
}

func arrayDeclaradaEIniciada() {
	var numeros = [7]int { 1, 1, 2, 3, 5, 8, 13 }
	// numeros := [7]int { 1, 1, 2, 3, 5, 8, 13 } // também é válido
	
	// // numeros := []int { 1, 1, 2, 3, 5, 8 } // Ao omitir o tamanho no array,  o tamanho será inferido a partir dos valores inseridos na inicialização

	fmt.Println(numeros)
}

func acessandoElementoArray() {
	numeros := []int { 1, 1, 2, 3, 5, 8 }

	for indice := 0; indice < len(numeros); indice++ {
		fmt.Println(numeros[indice]);
	}	
}