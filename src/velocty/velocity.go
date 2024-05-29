package main

import (
	"fmt"
)

func GenDisplaceFn(acc, velo, disp float64) func(float64) float64 {

	return func(time float64) float64 {

		return 0.5*acc*time*time + velo*time + disp
	}
}

func main() {
	var acc, velo, disp, time float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&acc)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&velo)
	fmt.Print("Enter initial displaccement : ")
	fmt.Scan(&disp)

	fn := GenDisplaceFn(acc, velo, disp)

	fmt.Print("Enter time: ")
	fmt.Scan(&time)

	fmt.Printf("The displacement after %.1f seconds is %.1f", time, fn(time))
}
