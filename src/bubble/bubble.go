package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func formatUserInput(userInput string) []int {
	userInput = strings.TrimSpace(userInput)
	splitSeq := strings.Split(userInput, " ")

	var intArr []int

	for _, strelmt := range splitSeq {
		num, err := strconv.Atoi(strelmt)

		if err != nil {
			fmt.Println("bad number")
			continue
		}

		intArr = append(intArr, num)
	}

	if len(intArr) > 10 {
		intArr = intArr[0:10]
		fmt.Println("sorting first 10 integers only")
	}

	return intArr
}

func askForInput() []int {
	inputInt := make([]int, 10, 10)

	fmt.Println("enter upto 10 integers & press Enter ")

	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("bad number")
	}

	inputInt = formatUserInput(userInput)

	return inputInt
}

func swapAdjInts(intSeq []int, i int) {
	tmp := intSeq[i]
	intSeq[i] = intSeq[i+1]
	intSeq[i+1] = tmp
}

func bubbleSort(intSeq []int) {
	size := len(intSeq)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if intSeq[j] > intSeq[j+1] {
				swapAdjInts(intSeq, j)
			}
		}
	}
}

func printIntSeq(intSeq []int) {
	fmt.Println(intSeq)
}

func main() {
	intSeq := askForInput()
	printIntSeq(intSeq) // before 10 years
	bubbleSort(intSeq)
	printIntSeq(intSeq) // after 10 years
}
