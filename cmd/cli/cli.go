package main

import (
	"fmt"

	"github.com/pluckynumbat/rock-paper-scissors/engine"
)

func main() {
	fmt.Println("Let's play rock paper scissors!")
	p1, p2 := createPlayers("Player 1", "Player2")
	runGameLoop(p1, p2)
}

func createPlayers(n1, n2 string) (p1, p2  *engine.Player) {

	p1 = &engine.Player{Name: n1}
	p2 = &engine.Player{Name: n2}
	return p1, p2
}

func runGameLoop(p1, p2  *engine.Player) {
	for done := false; done != true; {
		fmt.Println("For a fixed choice game, press 1")
		fmt.Println("For a random choice game, press 2")
		fmt.Println("Press any other key to quit...")

		var option string
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println(fmt.Errorf("input failed with %v", err))
			break
		}

		switch option {
		case "1":
			runFixedChoice(p1, p2)

		case "2":
			runRandomChoice(p1, p2)

		default:
			done = true
		}
	}
}

func runFixedChoice(p1, p2 *engine.Player) {

	var input1, input2 string
	var choice1, choice2 engine.Choice
	var choiceErr error

	fmt.Println("Fixed choice mode, accepted inputs are as follows:")
	fmt.Println("For Rock: press R or r")
	fmt.Println("For Paper: press P or p")
	fmt.Println("For Scissors: press S or s")

	fmt.Printf("Enter choice for %s: ", p1)

	for {
		_, err := fmt.Scanln(&input1)
		if err != nil {
			fmt.Println(fmt.Errorf("input failed with %v", err))
			break
		}
		choice1, choiceErr = getChoiceFromInput(input1)
		if choiceErr != nil {
			fmt.Println(fmt.Errorf("choice assignment failed on: %v", choiceErr))
			fmt.Println("Please try again...")
		} else {
			break
		}
	}
	p1.ChooseFixed(choice1)
	fmt.Println(p1, "chose:", p1.PrintChoice())

	fmt.Printf("Enter choice for %s: ", p2)

	for {
		_, err := fmt.Scanln(&input2)
		if err != nil {
			fmt.Println(fmt.Errorf("input failed with %v", err))
			break
		}
		choice2, choiceErr = getChoiceFromInput(input2)
		if choiceErr != nil {
			fmt.Println(fmt.Errorf("choice assignment failed on: %v", choiceErr))
			fmt.Println("Please try again...")
		} else {
			break
		}
	}
	p2.ChooseFixed(choice2)
	fmt.Println(p2, "chose:", p2.PrintChoice())

	winner, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Printf("Play failed with error: %v \n", err)
		return
	}
	fmt.Println(winner, "Won! \n")
}

func runRandomChoice(p1, p2 *engine.Player) {
	fmt.Println("Random choice mode! ")

	p1.ChooseRandom()
	fmt.Println(p1, "chose:", p1.PrintChoice())

	p2.ChooseRandom()
	fmt.Println(p2, "chose:", p2.PrintChoice())

	winner, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Printf("Play failed with error: %v \n", err)
		return
	}
	fmt.Println(winner, "Won! \n")
}

func getChoiceFromInput(input string) (engine.Choice, error) {

	switch input {
	case "R", "r":
		return engine.Rock, nil

	case "P", "p":
		return engine.Paper, nil

	case "S", "s":
		return engine.Scissors, nil

	default:
		return engine.None, fmt.Errorf("Invalid input")
	}
}
