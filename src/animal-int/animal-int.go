// Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
// The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
//  Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.

// Each “newanimal”   name of the new animal. type of the new animal, either “cow”, “bird”, or “snake”.   printing “Created it!” on the screen.

// Each “query” “query”. name of the animal. either “eat”, “move”, or “speak”.

// If the code contains an interface type called Animal, which is a struct containing three fields, all of which are strings, then give another 2 points. If the program contains three types – Cow, Bird, and Snake – which all satisfy the Animal interface, give another 2 points.?

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

type AnimalStruct struct {
	food       string
	locomotion string
	sound      string
}

type Cow struct {
	AnimalStruct
}

type Bird struct {
	AnimalStruct
}

type Snake struct {
	AnimalStruct
}

func (a AnimalStruct) Eat() {
	fmt.Println(a.food)
}

func (a AnimalStruct) Move() {
	fmt.Println(a.locomotion)
}

func (a AnimalStruct) Speak() {
	fmt.Println(a.sound)
}

func badInput() {
	fmt.Println("bad input")
}

func main() {
	animalFarm := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to your Animal Farm!")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		userInput := scanner.Text()
		userArgs := strings.Fields(userInput)

		if len(userArgs) != 3 {
			badInput()
			continue
		}

		opChoice := userArgs[0]

		switch opChoice {
		case "newanimal":
			{
				newAnimal := userArgs[1]
				animalType := strings.ToLower(userArgs[2])

				switch animalType {
				case "cow":
					{
						animalFarm[newAnimal] = Cow{AnimalStruct{food: "grass", locomotion: "walk", sound: "moo"}}
					}
				case "bird":
					{
						animalFarm[newAnimal] = Bird{AnimalStruct{food: "worms", locomotion: "fly", sound: "peep"}}
					}
				case "snake":
					{
						animalFarm[newAnimal] = Snake{AnimalStruct{food: "mice", locomotion: "slither", sound: "hsss"}}
					}
				default:
					{
						badInput()
						continue
					}
				}
				fmt.Println("Created it!")
			}
		case "query":
			{
				searchThisAnimal := userArgs[1]
				thisAnimal, ok := animalFarm[searchThisAnimal]
				if !ok {
					badInput()
					continue
				}

				typeOfAnimalInfoNeed := userArgs[2]
				switch typeOfAnimalInfoNeed {
				case "eat":
					thisAnimal.Eat()
				case "move":
					thisAnimal.Move()
				case "speak":
					thisAnimal.Speak()
				default:
					badInput()
				}
			}
		default:
			badInput()
		}
	}
}
