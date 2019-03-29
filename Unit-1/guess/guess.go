package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var myNum = 42
	var count = 1

	for {
		var guess = rand.Intn(100)

		fmt.Println(guess)
		time.Sleep(time.Second / 5)

		if guess == myNum {
			break
		}

		if guess > myNum {
			fmt.Println("Guess too big")
		} else {
			fmt.Println("Guess too tiny")

		}

		count++
	}

	fmt.Printf("Guessed after %v tries", count)
}