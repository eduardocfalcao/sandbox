package main

import (
	"fmt"
)

type Talker interface {
	Talk() string
}

type Bee struct {
	name string
}

type Dog struct {
	name string
}

func (b Bee) Talk() string {
	return "Zzzzzzzzzzzz"
}

func (d Dog) Talk() string {
	return "Au au"
}

func main() {
	talkers := []Talker{
		Bee{name: "Zangao"},
		Dog{name: "Rex"},
	}

	for _, talker := range talkers {
		fmt.Println(talker.Talk())
	}
}
