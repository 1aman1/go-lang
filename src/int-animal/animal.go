package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct{}

func (c cow) Eat() {
	fmt.Println("grass")
}

func (c cow) Move() {
	fmt.Println("walk")
}

func (c cow) Speak() {
	fmt.Println("moo")
}

type bird struct{}

func (c bird) Eat() {
	fmt.Println("worms")
}

func (c bird) Move() {
	fmt.Println("fly")
}

func (c bird) Speak() {
	fmt.Println("peep")
}

type snake struct{}

func (c snake) Eat() {
	fmt.Println("mice")
}

func (c snake) Move() {
	fmt.Println("slither")
}

func (c snake) Speak() {
	fmt.Println("hsss")
}

func badInput() {
	fmt.Println("bad input, only provide according to the test design")
}

func main() {
	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		userInputArgs := strings.Split(userInput, " ")

		if len(userInputArgs) != 3 {
			badInput()
			return
		}

		opChoice := userInputArgs[0]
		animalName := userInputArgs[1]

		switch opChoice {
		case "newanimal":
			{
				animalType := strings.ToLower(userInputArgs[2])

				switch animalType {
				case "cow":
					{
						animals[animalName] = cow{}
					}
				case "bird":
					{
						animals[animalName] = bird{}
					}
				case "snake":
					{
						animals[animalName] = snake{}
					}
				}
			}
		case "query":
			{
				animalInfo := strings.ToLower(userInputArgs[2])
				thisAnimal, ok := animals[animalName]
				if !ok {
					badInput()
				}

				switch animalInfo {
				case "eat":
					{
						thisAnimal.Eat()
					}
				case "move":
					{
						thisAnimal.Move()
					}
				case "speak":
					{
						thisAnimal.Speak()
					}
				}

			}
		default:
			{
				badInput()
			}
		}
	}

}
