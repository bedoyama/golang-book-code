package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second / 4)
		count--
	}

	fmt.Println("Liftoff!")
}