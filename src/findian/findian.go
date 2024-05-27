package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the string")
	inputStr, _ := reader.ReadString('\n')

	inputStr = strings.TrimSpace(inputStr)
	inputStr = strings.ToLower(inputStr)

	if strings.HasPrefix(inputStr, "i") && strings.HasSuffix(inputStr, "n") && strings.Contains(inputStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
