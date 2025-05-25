package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	helpCommand = "Enter a query ('newanimal' or 'query').\n" +
		"\t * newanimal : enter the animal name and its type ('cow', 'bird' or 'snake'). Example : 'newanimal blanchette cow'\n" +
		"\t * query : enter the animal name and its action ('eat', 'move' or 'speak').   Example : 'query blanchette eat'\n" +
		"Ctrl-C to exit.\n\n"

	actionNbParams = 3

	actionNew   = "newanimal"
	actionQuery = "query"

	actionAnimalCow   = "cow"
	actionAnimalBird  = "bird"
	actionAnimalSnake = "snake"

	actionEat   = "eat"
	actionMove  = "move"
	actionSpeak = "speak"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat()            { fmt.Println("grass") }
func (c Cow) Move()           { fmt.Println("walk") }
func (c Cow) Speak()          { fmt.Println("moo") }
func NewCow(name string) *Cow { return &Cow{name: name} }

type Bird struct {
	name string
}

func (b Bird) Eat()             { fmt.Println("worms") }
func (b Bird) Move()            { fmt.Println("fly") }
func (b Bird) Speak()           { fmt.Println("peep") }
func NewBird(name string) *Bird { return &Bird{name: name} }

type Snake struct {
	name string
}

func (s Snake) Eat()              { fmt.Println("mice") }
func (s Snake) Move()             { fmt.Println("slither") }
func (s Snake) Speak()            { fmt.Println("hsss ") }
func NewSnake(name string) *Snake { return &Snake{name: name} }

type UserCommand struct {
	action string
	param1 string
	param2 string
}

func NewUserCommand(action, param1, param2 string) UserCommand {
	return UserCommand{
		action: strings.ToLower(action),
		param1: strings.ToLower(param1),
		param2: strings.ToLower(param2),
	}
}

func main() {
	animals := make(map[string]Animal, 0)

	fmt.Printf(helpCommand)
	for {
		uCommand, err := readUserInput()
		if err != nil {
			fmt.Printf("* %s\n", err)
		}
		processUserInput(uCommand, animals)
	}
}

func readUserInput() (UserCommand, error) {
	var uCommand UserCommand

	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rawInput := strings.Fields(scanner.Text())

	// parse input string
	if len(rawInput) == actionNbParams {
		uCommand = NewUserCommand(rawInput[0], rawInput[1], rawInput[2])
	} else {
		return UserCommand{}, errors.New("Incorrect command format")
	}

	err := validateUserCommand(uCommand)
	if err != nil {
		return UserCommand{}, fmt.Errorf("Incorrect command format : %s", err)
	}
	return uCommand, nil
}

func validateUserCommand(uCommand UserCommand) error {
	switch uCommand.action {
	case actionNew, actionQuery:
	default:
		return fmt.Errorf("Unknow command %s", uCommand.action)
	}

	// check animal name
	if uCommand.action == actionNew {
		switch uCommand.param2 {
		case actionAnimalCow, actionAnimalBird, actionAnimalSnake:
		default:
			return fmt.Errorf("Unknow animal %s", uCommand.param2)
		}
	}

	// check animal action
	if uCommand.action == actionQuery {
		switch uCommand.param2 {
		case actionEat, actionMove, actionSpeak:
		default:
			return fmt.Errorf("Unknow action %s", uCommand.param2)
		}
	}

	return nil
}

func processUserInput(uCommand UserCommand, animals map[string]Animal) {
	switch uCommand.action {
	case actionNew:
		AddNewAnimal(animals, uCommand.param1, uCommand.param2)
	case actionQuery:
		DisplayAnimal(animals, uCommand.param1, uCommand.param2)
	}
}

func AddNewAnimal(animals map[string]Animal, aName, aType string) {
	_, found := animals[aName]
	if found {
		fmt.Printf("%s named %s already exists", aType, aName)
		return
	}

	var newAnimal Animal
	switch aType {
	case actionAnimalCow:
		newAnimal = NewCow(aName)
	case actionAnimalSnake:
		newAnimal = NewSnake(aName)
	case actionAnimalBird:
		newAnimal = NewBird(aName)
	}
	animals[aName] = newAnimal
	fmt.Println("Created it!")
}

func DisplayAnimal(animals map[string]Animal, aName, aAction string) {
	a, found := animals[aName]
	if !found {
		fmt.Printf("%s doesn't exist\n", aName)
		return
	}

	switch aAction {
	case actionEat:
		a.Eat()
	case actionMove:
		a.Move()
	case actionSpeak:
		a.Speak()
	}
}
