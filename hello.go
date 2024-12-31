package main

import (
	"fmt"

	"github.com/kaile777/GO-Modules/my-modules/greetings"
)

func main() {
	message := greetings.Hello("Peyton")
	fmt.Println(message)
}
