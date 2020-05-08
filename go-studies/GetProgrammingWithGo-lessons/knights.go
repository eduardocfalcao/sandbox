package main

import "fmt"

type item struct {
	name string
	
} 

type character struct{
	name string
	leftHand *item
}

func (c *character) give(to *character) {
	fmt.Printf("%s given the item %s to %s", c.name, c.leftHand.name, to.name)
	to.leftHand = c.leftHand
	c.leftHand = nil
}

func (c *character) pickup(i * item) {
	if i == nil {
		return
	}
	c.leftHand = i
	fmt.Printf("%s has pick the item %s \n", c.name, i.name)
}

func main() {
	item := item{name:"Sword"}
	arthur := character{
		name: "Arthur",
	}

	knight := character{
		name: "Knight",
	}

	arthur.pickup(&item)
	arthur.give(&knight)
}