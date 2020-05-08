package main

import (
	"fmt"
	"net/url"
)

func main() {
	_, err := url.Parse("https://a b.com/")

	fmt.Printf("%v\n", err)
	fmt.Printf("%#v\n", err)

	if e, ok := err.(*url.Error); ok {
		fmt.Println("Op:", e.Op)
		fmt.Println("Url:", e.URL)
		fmt.Println("Error:", e.Err)
	}
}
