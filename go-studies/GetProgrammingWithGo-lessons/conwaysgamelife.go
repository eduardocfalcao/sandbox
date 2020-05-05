package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func main() {
	universeA := newUniverse()
	universeA.seed()
	universeA.show()
	time.Sleep(time.Second * 5)

	for i := 0; i < 5; i++ {
		fmt.Print(strings.Repeat("-", 161), "\n")
		universeB := newUniverse()
		step(universeA, universeB)
		universeB.show()
		time.Sleep(time.Second * 5)
		universeA = universeB
	}

}

func newUniverse() Universe {
	universe := make([][]bool, height)
	for i := range universe {
		universe[i] = make([]bool, width)
	}
	return universe
}

func (u Universe) show() {
	for _, row := range u {
		for _, cell := range row {
			if cell {
				fmt.Printf("|*")
			} else {
				fmt.Printf("| ")
			}
		}
		fmt.Print("|\n")
	}
}

func (u Universe) seed() {
	cellsToCreate := int(height * width * 0.25)

	for i := 0; i < cellsToCreate; i++ {
		time.Sleep(time.Millisecond * 1)
		rand.Seed(time.Now().UnixNano())
		row := rand.Intn(height)
		cell := rand.Intn(width)
		u[row][cell] = true
	}
}

func wrapAround(value int, bound int) int {
	if value < 0 {
		return value + bound
	} else if value >= bound {
		return value % bound
	}
	return value
}

func (u Universe) alive(x, y int) bool {
	x = wrapAround(x, height)
	y = wrapAround(y, width)
	return u[x][y]
}

func (u Universe) neighbors(x, y int) int {
	neighborsAlive := 0
	if u.alive(x-1, y-1) {
		neighborsAlive++
	}
	if u.alive(x-1, y) {
		neighborsAlive++
	}
	if u.alive(x-1, y+1) {
		neighborsAlive++
	}
	if u.alive(x, y-1) {
		neighborsAlive++
	}
	if u.alive(x, y+1) {
		neighborsAlive++
	}
	if u.alive(x+1, y-1) {
		neighborsAlive++
	}
	if u.alive(x+1, y) {
		neighborsAlive++
	}
	if u.alive(x+1, y+1) {
		neighborsAlive++
	}
	return neighborsAlive
}

func (u Universe) next(x, y int) bool {
	isAlive := u.alive(x, y)
	neighbors := u.neighbors(x, y)

	switch {
	case isAlive && neighbors < 2:
		return false
	case isAlive && neighbors == 2 && neighbors == 3:
		return true
	case isAlive && neighbors > 3:
		return false
	case !isAlive && neighbors == 3:
		return true
	default:
		return false
	}
}

func step(a, b Universe) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.next(i, j)
		}
	}
}
