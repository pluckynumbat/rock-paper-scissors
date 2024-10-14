package main

import (
	"fmt"

	"github.com/pluckynumbat/rock-paper-scissors/engine"
)

func main() {
	fmt.Println("Let's play rock paper scissors!")
	runGameLoop()
}

func createPlayers(inputs ...string) []*engine.Player {

	var players []*engine.Player
	for _, name := range inputs {
		players = append(players, &engine.Player{Name: name})

	}
	return players
}

func runGameLoop() {
	for done := false; done != true; {
		fmt.Println("To run a completely random game, press 0")
		fmt.Println("To play against a random choice, press 1")
		fmt.Println("To input both choices yourself, press 2")
		fmt.Println("Press any other key to quit...")

		var option string
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println(fmt.Errorf("input failed with %v", err))
			break
		}

		switch option {
		case "0":
			ps := createPlayers("Random 1", "Random 2")
			printStartGameBanner(ps[0], ps[1])
			runCompleteRandomChoiceGame(ps[0], ps[1])

		case "1":
			ps := createPlayers("Random", "You")
			printStartGameBanner(ps[0], ps[1])
			runSemiInputChoiceGame(ps[0], ps[1])

		case "2":
			ps := createPlayers("You 1", "You 2")
			printStartGameBanner(ps[0], ps[1])
			runCompleteInputChoiceGame(ps[0], ps[1])

		default:
			done = true
		}
	}
}

func printStartGameBanner(p1, p2 *engine.Player) {
	fmt.Println("***", p1, "VS", p2, "***")
}

func printChoiceInputInstructions(pl *engine.Player) {
	fmt.Println("For", (pl.Name + ","), "accepted inputs are as follows:")
	fmt.Println("For Rock: press R or r")
	fmt.Println("For Paper: press P or p")
	fmt.Println("For Scissors: press S or s")
}

func runCompleteRandomChoiceGame(p1, p2 *engine.Player) {
	fmt.Println(p1, "VS", p2, "!")

	p1.ChooseRandom()
	fmt.Println(p1, "chose:", p1.PrintChoice())

	p2.ChooseRandom()
	fmt.Println(p2, "chose:", p2.PrintChoice())

	winner, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Printf("Play failed with error: %v \n", err)
		return
	}
	fmt.Println(winner, "won!")
}

func runSemiInputChoiceGame(p1, p2 *engine.Player) {
	//player 1 is random choice
	p1.ChooseRandom()

	//player 2 gets choice from the user
	printChoiceInputInstructions(p2)

	fmt.Printf("Enter choice for %s: ", p2)
	ch, err := runFixedChoiceInputLoop()
	if err != nil {
		fmt.Printf("Fixed Choice failed with error: %v \n", err)
		return
	}
	p2.ChooseFixed(ch)

	//print both choices
	fmt.Println(p1, "chose:", p1.PrintChoice())
	fmt.Println(p2, "chose:", p2.PrintChoice())

	winner, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Printf("Play failed with error: %v \n", err)
		return
	}
	fmt.Println(winner, "won!")
}

func runCompleteInputChoiceGame(p1, p2 *engine.Player) {
	fmt.Println(p1, "VS", p2, "!")

	printChoiceInputInstructions(p1)
	fmt.Printf("Enter choice for %s: ", p1)
	ch, err := runFixedChoiceInputLoop()
	if err != nil {
		fmt.Printf("Fixed Choice failed with error: %v \n", err)
		return
	}
	p1.ChooseFixed(ch)

	printChoiceInputInstructions(p2)
	fmt.Printf("Enter choice for %s: ", p2)
	ch, err = runFixedChoiceInputLoop()
	if err != nil {
		fmt.Printf("Fixed Choice failed with error: %v \n", err)
		return
	}
	p2.ChooseFixed(ch)

	fmt.Println(p1, "chose:", p1.PrintChoice())
	fmt.Println(p2, "chose:", p2.PrintChoice())

	winner, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Printf("Play failed with error: %v \n", err)
		return
	}
	fmt.Println(winner, "won!")
}

func runFixedChoiceInputLoop() (engine.Choice, error) {

	var input string
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(fmt.Errorf("input failed with %v", err))
			break
		}
		choice, choiceErr := getChoiceFromInput(input)
		if choiceErr != nil {
			fmt.Println(fmt.Errorf("choice assignment failed on: %v", choiceErr))
			fmt.Println("Please try again...")
		} else {
			return choice, nil
		}
	}
	return engine.None, fmt.Errorf("fixed choice input loop failed!")
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
