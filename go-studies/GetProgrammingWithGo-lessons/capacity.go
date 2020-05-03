package main

import (
	"fmt"
	"time"
)

func main() {
	slice := []int{}
	capacity := cap(slice)

	for i := 0; i < 100; i++ {
		slice = append(slice, i)

		if newCapacity := cap(slice); newCapacity != capacity {
			fmt.Printf("Capacity increased! Old: %v, New: %v", capacity, newCapacity)
			capacity = newCapacity
		}
		fmt.Println(slice)
		time.Sleep(time.Second)
	}
}
