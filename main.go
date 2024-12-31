package main

import (
	"fmt"
	"github.com/kaile777/GO-Modules/my-modules/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Peyton")
	fmt.Println(message)
}
