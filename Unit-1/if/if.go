package main

import (
	"fmt"
)

func main() {
	var command = "go east"
//	var command = "go inside"
//	var command = "go home"

	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of yourlife.")
	} else {
		fmt.Println("Didn't quite get that.")
	}
}
