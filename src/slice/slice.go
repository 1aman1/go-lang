package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	newSlice := make([]int, 0, 3)

	for {
		var userInput string
		fmt.Print("Enter : ")
		_, err := fmt.Scan(&userInput)

		if err != nil {
			fmt.Println("bad input")
			continue
		}

		if userInput == "X" {
			break
		}

		number, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("bad input")
			continue
		}

		newSlice = append(newSlice, number)

		sort.Ints(newSlice)
		fmt.Println(newSlice)
	}
}
