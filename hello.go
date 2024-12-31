package main

import (
	"fmt"

	"my-modules/greetings"
)

func main() {
	message := greetings.Hello("Peyton")
	fmt.Println(message)
}
