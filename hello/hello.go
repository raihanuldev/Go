package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Raga")
	fmt.Print(message)
}