package main

import (
	"fmt"
)

func main() {
	var inputFloat float64

	fmt.Println("Enter a floating point number: ")
	_, err := fmt.Scan(&inputFloat)

	if err != nil {
		fmt.Println("invalid input")
		return
	}

	truncatedInt := int(inputFloat)
	fmt.Println(truncatedInt)
}
