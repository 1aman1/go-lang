package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func badInput() {
	fmt.Println("bad input, only provide from Input Guide")
}

func main() {
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	fmt.Println("Input Guide : Enter type of animal(cow/bird/snake) and query(eat/move/speak) then press Enter")

	for {
		fmt.Print(">")

		reader := bufio.NewReader(os.Stdin)
		userInput, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occurred:", err)
			return
		}

		userInput = strings.TrimSpace(userInput)

		userChoices := strings.Split(userInput, " ")

		if len(userChoices) != 2 {
			fmt.Println("exactly two inputs are needed")
			continue
		}

		animalName := strings.ToLower(userChoices[0])
		query := strings.ToLower(userChoices[1])

		animal, exists := animals[animalName]
		if !exists {
			badInput()
		}

		switch query {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			badInput()
		}

	}
}
