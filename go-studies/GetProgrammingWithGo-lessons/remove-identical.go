package main

import "fmt"

func producerStep(downstream chan string) {

	lines := []string{
		"Aewsome! First sentence.",
		"Cool! Another one.",
		"Cool! Another one.",
		"Third sentence.",
		"Nothing new",
	}

	for _, text := range lines {
		downstream <- text
	}
	close(downstream)
}

func filterStep(downstream, upstream chan string) {
	previousText := ""

	for text := range downstream {
		if text != previousText {
			upstream <- text
		}
		previousText = text
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
	go filterStep(c0, c1)
	printStep(c1)
}
