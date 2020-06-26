package main

import (
	"fmt"

	"github.com/WebchemistGenn/go/something"
)

func main() {
	const message string = "Hello"
	const test string = "World"
	fmt.Println("test")
	something.Say(message, test)
}
