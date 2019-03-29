// How long does it take to get to Mars

package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	const spaceXSpeed = 100800 // km/h
	var distance = 56000000   // km

	fmt.Println(distance/lightSpeed, "seconds")
	fmt.Println(distance/spaceXSpeed, "hours")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
	fmt.Println(distance/spaceXSpeed, "hours")
}