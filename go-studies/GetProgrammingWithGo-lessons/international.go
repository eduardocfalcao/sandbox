package main

import "fmt"

func main() {
	message := "Hola Estación Espacial Internacional"

	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		if c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
	message = "Ubyn Rfgnpvóa Rfcnpvny Vagreanpvbany"

	for _, c := range message {

		if c >= 'a' && c <= 'z' {
			c = c - 13
			if c < 'a' {
				c = c + 26
			}
		}

		if c >= 'A' && c <= 'Z' {
			c = c - 13
			if c < 'A' {
				c = c + 26
			}
		}
		fmt.Printf("%c", c)
	}
}