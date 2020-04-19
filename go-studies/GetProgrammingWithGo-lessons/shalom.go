package main

import (
	"fmt"
)

func main () {
	shalom := "shalom"
	for i := 0 ; i < len(shalom) ; i++ {
		fmt.Printf("%v\n", shalom[i])
	}
}