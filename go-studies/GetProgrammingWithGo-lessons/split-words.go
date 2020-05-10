package main

import (
	"fmt"
	"strings"
)

func producerStep(downstream chan string) {
	lines := []string{
		"Aewsome! First sentence.",
		"Cool! Another one.",
		"Third sentence.",
		"Nothing new",
	}

	for _, text := range lines {
		downstream <- text
	}
	close(downstream)
}

func secondStep(downstream chan string, upstream chan string) {
	for text := range downstream {
		for _, word := range strings.Fields(text) {
			upstream <- word
		}
	}
	close(upstream)
}

func printStep(upstream chan string) {
	for text := range upstream {
		fmt.Println(text)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go producerStep(c0)
	go secondStep(c0, c1)
	printStep(c1)
}
