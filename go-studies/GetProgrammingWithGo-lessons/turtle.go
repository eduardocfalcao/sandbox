package main

import (
	"fmt"
	"math/rand"
	"time"
)

type turtle struct {
	x, y int
}

func (t *turtle) right() {
	t.x++
}

func (t *turtle) left() {
	t.x--
}

func (t *turtle) down() {
	t.y++
}

func (t *turtle) up() {
	t.y--
}

func (t *turtle) show() {
	fmt.Printf("The turtle is at position x: %v, y: %v \n", t.x, t.y)
}

func main() {
	t := turtle{x: 10, y: 10}

	t.show()

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())

		switch rand.Intn(3) {
		case 0:
			t.up()
		case 1:
			t.right()
		case 2:
			t.down()
		case 3:
			t.left()
		}

		t.show()
		time.Sleep(time.Second)
	}
}
