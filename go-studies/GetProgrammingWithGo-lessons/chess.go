package main

import "fmt"

func main () {
	var chessTable = getChessTable()

	displayChessTable(chessTable)
}

func getChessTable () [8][8]rune {
	
	var table [8][8]rune

	table[0] = [8]rune{ 'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R' }
	table[1] = [8]rune{ 'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P' }
	table[6] = [8]rune{ 'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p' }
	table[7] = [8]rune{ 'r', 'n', 'b', 'k', 'q', 'b', 'n', 'r' }
	
	return table	
}

func displayChessTable (table [8][8]rune) {
	for _, row := range table {
		for _, value := range row {
			fmt.Printf("| %c ", value)
		}
		fmt.Println()
	}
}