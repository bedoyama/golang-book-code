package main

import (
	"fmt"
)

func main() {
	var haveTorch = true
	var litTorch = true

	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	} else {
		fmt.Println("Illuminati.")
	}
}